package k8sbundle

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"sync"
	"testing"

	"github.com/spiffe/spire/pkg/common/catalog"
	"github.com/spiffe/spire/pkg/server/plugin/hostservices"
	"github.com/spiffe/spire/proto/spire/common"
	spi "github.com/spiffe/spire/proto/spire/common/plugin"
	notifierv0 "github.com/spiffe/spire/proto/spire/server/notifier/v0"
	"github.com/spiffe/spire/test/fakes/fakeidentityprovider"
	"github.com/spiffe/spire/test/spiretest"
	"google.golang.org/grpc/codes"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
)

var (
	testBundle = &common.Bundle{
		RootCas: []*common.Certificate{
			{DerBytes: []byte("FOO")},
			{DerBytes: []byte("BAR")},
		},
	}

	testBundle2 = &common.Bundle{
		RootCas: []*common.Certificate{
			{DerBytes: []byte("BAR")},
			{DerBytes: []byte("BAZ")},
		},
	}
)

const (
	// PEM encoding of the root CAs in testBundle
	testBundleData  = "-----BEGIN CERTIFICATE-----\nRk9P\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nQkFS\n-----END CERTIFICATE-----\n"
	testBundle2Data = "-----BEGIN CERTIFICATE-----\nQkFS\n-----END CERTIFICATE-----\n-----BEGIN CERTIFICATE-----\nQkFa\n-----END CERTIFICATE-----\n"
)

func Test(t *testing.T) {
	spiretest.Run(t, new(Suite))
}

type Suite struct {
	spiretest.Suite

	r *fakeidentityprovider.IdentityProvider
	k *fakeKubeClient

	raw *Plugin
	p   notifierv0.Plugin
}

func (s *Suite) SetupTest() {
	s.r = fakeidentityprovider.New()
	s.k = newFakeKubeClient()

	s.raw = New()
	s.LoadPlugin(builtIn(s.raw), &s.p,
		spiretest.HostService(hostservices.IdentityProviderHostServiceServer(s.r)))

	s.withKubeClient(s.k, "")
}

func (s *Suite) TestNotifyFailsIfNotConfigured() {
	resp, err := s.p.Notify(context.Background(), &notifierv0.NotifyRequest{})
	s.RequireGRPCStatus(err, codes.Unknown, "k8s-bundle: not configured")
	s.Nil(resp)
}

func (s *Suite) TestNotifyIgnoresUnknownEvents() {
	s.configure("")

	resp, err := s.p.Notify(context.Background(), &notifierv0.NotifyRequest{})
	s.NoError(err)
	s.AssertProtoEqual(&notifierv0.NotifyResponse{}, resp)
}

func (s *Suite) TestNotifyAndAdviseFailsIfNotConfigured() {
	resp, err := s.p.NotifyAndAdvise(context.Background(), &notifierv0.NotifyAndAdviseRequest{})
	s.RequireGRPCStatus(err, codes.Unknown, "k8s-bundle: not configured")
	s.Nil(resp)
}

func (s *Suite) TestNotifyAndAdviseIgnoresUnknownEvents() {
	s.configure("")

	resp, err := s.p.NotifyAndAdvise(context.Background(), &notifierv0.NotifyAndAdviseRequest{})
	s.NoError(err)
	s.AssertProtoEqual(&notifierv0.NotifyAndAdviseResponse{}, resp)
}

func (s *Suite) TestBundleLoadedWhenCannotCreateClient() {
	s.withKubeClient(nil, "")

	s.configure("")

	resp, err := s.p.NotifyAndAdvise(context.Background(), &notifierv0.NotifyAndAdviseRequest{
		Event: &notifierv0.NotifyAndAdviseRequest_BundleLoaded{
			BundleLoaded: &notifierv0.BundleLoaded{
				Bundle: testBundle,
			},
		},
	})
	s.RequireGRPCStatus(err, codes.Unknown, "kube client not configured")
	s.Nil(resp)
}

func (s *Suite) TestBundleLoadedConfigMapGetFailure() {
	s.configure("")

	resp, err := s.p.NotifyAndAdvise(context.Background(), &notifierv0.NotifyAndAdviseRequest{
		Event: &notifierv0.NotifyAndAdviseRequest_BundleLoaded{
			BundleLoaded: &notifierv0.BundleLoaded{
				Bundle: testBundle,
			},
		},
	})
	s.RequireGRPCStatus(err, codes.Unknown, "k8s-bundle: unable to update: unable to get list: not found")
	s.Nil(resp)
}

func (s *Suite) TestBundleLoadedConfigMapPatchFailure() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, string(debug.Stack()))
		}
	}()
	s.k.setConfigMap(newConfigMap())
	s.k.setPatchErr(errors.New("some error"))
	s.r.AppendBundle(testBundle)

	s.configure("")

	resp, err := s.p.NotifyAndAdvise(context.Background(), &notifierv0.NotifyAndAdviseRequest{
		Event: &notifierv0.NotifyAndAdviseRequest_BundleLoaded{
			BundleLoaded: &notifierv0.BundleLoaded{
				Bundle: testBundle,
			},
		},
	})
	s.RequireGRPCStatus(err, codes.Unknown, "k8s-bundle: unable to update: spire/spire-bundle: some error")
	s.Nil(resp)
}

func (s *Suite) TestBundleLoadedConfigMapUpdateConflict() {
	s.k.setConfigMap(newConfigMap())
	s.k.setPatchErr(&k8serrors.StatusError{
		ErrStatus: metav1.Status{
			Code:    http.StatusConflict,
			Message: "unexpected version",
			Reason:  "Conflict",
		},
	})

	// return a different bundle when fetched the second time
	s.r.AppendBundle(testBundle)
	s.r.AppendBundle(testBundle2)

	s.configure("")

	resp, err := s.p.NotifyAndAdvise(context.Background(), &notifierv0.NotifyAndAdviseRequest{
		Event: &notifierv0.NotifyAndAdviseRequest_BundleLoaded{
			BundleLoaded: &notifierv0.BundleLoaded{
				Bundle: testBundle,
			},
		},
	})
	s.NoError(err)
	s.AssertProtoEqual(&notifierv0.NotifyAndAdviseResponse{}, resp)

	// make sure the config map contains the second bundle data
	configMap := s.k.getConfigMap("spire", "spire-bundle")
	s.Require().NotNil(configMap)
	s.Require().NotNil(configMap.Data)
	s.Equal(testBundle2Data, configMap.Data["bundle.crt"])
}

func (s *Suite) TestBundleLoadedWithDefaultConfiguration() {
	s.configure("")
	s.k.setConfigMap(newConfigMap())
	s.r.AppendBundle(testBundle)

	resp, err := s.p.NotifyAndAdvise(context.Background(), &notifierv0.NotifyAndAdviseRequest{
		Event: &notifierv0.NotifyAndAdviseRequest_BundleLoaded{
			BundleLoaded: &notifierv0.BundleLoaded{
				Bundle: testBundle,
			},
		},
	})
	s.Require().NoError(err)
	s.RequireProtoEqual(&notifierv0.NotifyAndAdviseResponse{}, resp)

	s.Require().Equal(&corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       "spire",
			Name:            "spire-bundle",
			ResourceVersion: "2",
		},
		Data: map[string]string{
			"bundle.crt": testBundleData,
		},
	}, s.k.getConfigMap("spire", "spire-bundle"))
}

func (s *Suite) TestBundleLoadedWithConfigurationOverrides() {
	s.withKubeClient(s.k, "/some/file/path")

	s.k.setConfigMap(&corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       "NAMESPACE",
			Name:            "CONFIGMAP",
			ResourceVersion: "2",
		},
	})
	s.r.AppendBundle(testBundle)

	s.configure(`
namespace = "NAMESPACE"
config_map = "CONFIGMAP"
config_map_key = "CONFIGMAPKEY"
kube_config_file_path = "/some/file/path"
`)

	resp, err := s.p.NotifyAndAdvise(context.Background(), &notifierv0.NotifyAndAdviseRequest{
		Event: &notifierv0.NotifyAndAdviseRequest_BundleLoaded{
			BundleLoaded: &notifierv0.BundleLoaded{
				Bundle: testBundle,
			},
		},
	})
	s.Require().NoError(err)
	s.NotNil(resp)

	s.Equal(&corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       "NAMESPACE",
			Name:            "CONFIGMAP",
			ResourceVersion: "3",
		},
		Data: map[string]string{
			"CONFIGMAPKEY": testBundleData,
		},
	}, s.k.getConfigMap("NAMESPACE", "CONFIGMAP"))
}

func (s *Suite) TestBundleUpdatedWhenCannotCreateClient() {
	s.withKubeClient(nil, "")

	s.configure("")

	resp, err := s.p.Notify(context.Background(), &notifierv0.NotifyRequest{
		Event: &notifierv0.NotifyRequest_BundleUpdated{
			BundleUpdated: &notifierv0.BundleUpdated{
				Bundle: testBundle,
			},
		},
	})
	s.RequireGRPCStatus(err, codes.Unknown, "kube client not configured")
	s.Nil(resp)
}

func (s *Suite) TestBundleUpdatedConfigMapGetFailure() {
	s.configure("")

	resp, err := s.p.Notify(context.Background(), &notifierv0.NotifyRequest{
		Event: &notifierv0.NotifyRequest_BundleUpdated{
			BundleUpdated: &notifierv0.BundleUpdated{
				Bundle: testBundle,
			},
		},
	})
	s.RequireGRPCStatus(err, codes.Unknown, "k8s-bundle: unable to update: unable to get list: not found")
	s.Nil(resp)
}

func (s *Suite) TestBundleUpdatedConfigMapPatchFailure() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, string(debug.Stack()))
		}
	}()
	s.k.setConfigMap(newConfigMap())
	s.k.setPatchErr(errors.New("some error"))
	s.r.AppendBundle(testBundle)

	s.configure("")

	resp, err := s.p.Notify(context.Background(), &notifierv0.NotifyRequest{
		Event: &notifierv0.NotifyRequest_BundleUpdated{
			BundleUpdated: &notifierv0.BundleUpdated{
				Bundle: testBundle,
			},
		},
	})
	s.RequireGRPCStatus(err, codes.Unknown, "k8s-bundle: unable to update: spire/spire-bundle: some error")
	s.Nil(resp)
}

func (s *Suite) TestBundleUpdatedConfigMapUpdateConflict() {
	s.k.setConfigMap(newConfigMap())
	s.k.setPatchErr(&k8serrors.StatusError{
		ErrStatus: metav1.Status{
			Code:    http.StatusConflict,
			Message: "unexpected version",
			Reason:  "Conflict",
		},
	})

	// return a different bundle when fetched the second time
	s.r.AppendBundle(testBundle)
	s.r.AppendBundle(testBundle2)

	s.configure("")

	resp, err := s.p.Notify(context.Background(), &notifierv0.NotifyRequest{
		Event: &notifierv0.NotifyRequest_BundleUpdated{
			BundleUpdated: &notifierv0.BundleUpdated{
				Bundle: testBundle,
			},
		},
	})
	s.NoError(err)
	s.AssertProtoEqual(&notifierv0.NotifyResponse{}, resp)

	// make sure the config map contains the second bundle data
	configMap := s.k.getConfigMap("spire", "spire-bundle")
	s.Require().NotNil(configMap)
	s.Require().NotNil(configMap.Data)
	s.Equal(testBundle2Data, configMap.Data["bundle.crt"])
}

func (s *Suite) TestBundleUpdatedWithDefaultConfiguration() {
	s.configure("")
	s.k.setConfigMap(newConfigMap())
	s.r.AppendBundle(testBundle)

	resp, err := s.p.Notify(context.Background(), &notifierv0.NotifyRequest{
		Event: &notifierv0.NotifyRequest_BundleUpdated{
			BundleUpdated: &notifierv0.BundleUpdated{
				Bundle: testBundle,
			},
		},
	})
	s.Require().NoError(err)
	s.RequireProtoEqual(&notifierv0.NotifyResponse{}, resp)

	s.Equal(&corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       "spire",
			Name:            "spire-bundle",
			ResourceVersion: "2",
		},
		Data: map[string]string{
			"bundle.crt": testBundleData,
		},
	}, s.k.getConfigMap("spire", "spire-bundle"))
}

func (s *Suite) TestBundleUpdatedWithConfigurationOverrides() {
	s.withKubeClient(s.k, "/some/file/path")

	s.k.setConfigMap(&corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       "NAMESPACE",
			Name:            "CONFIGMAP",
			ResourceVersion: "2",
		},
	})
	s.r.AppendBundle(testBundle)

	s.configure(`
namespace = "NAMESPACE"
config_map = "CONFIGMAP"
config_map_key = "CONFIGMAPKEY"
kube_config_file_path = "/some/file/path"
`)

	resp, err := s.p.Notify(context.Background(), &notifierv0.NotifyRequest{
		Event: &notifierv0.NotifyRequest_BundleUpdated{
			BundleUpdated: &notifierv0.BundleUpdated{
				Bundle: testBundle,
			},
		},
	})
	s.Require().NoError(err)
	s.NotNil(resp)

	s.Equal(&corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       "NAMESPACE",
			Name:            "CONFIGMAP",
			ResourceVersion: "3",
		},
		Data: map[string]string{
			"CONFIGMAPKEY": testBundleData,
		},
	}, s.k.getConfigMap("NAMESPACE", "CONFIGMAP"))
}

func (s *Suite) TestConfigureWithMalformedConfiguration() {
	_, err := s.p.Configure(context.Background(), &spi.ConfigureRequest{
		Configuration: "blah",
	})
	s.RequireGRPCStatusContains(err, codes.Unknown, "k8s-bundle: unable to decode configuration")
}

func (s *Suite) TestGetPluginInfo() {
	resp, err := s.p.GetPluginInfo(context.Background(), &spi.GetPluginInfoRequest{})
	s.NoError(err)
	s.AssertProtoEqual(&spi.GetPluginInfoResponse{}, resp)
}

func (s *Suite) TestBundleFailsToLoadIfHostServicesUnavailabler() {
	p, err := catalog.LoadBuiltInPlugin(context.Background(), catalog.BuiltInPlugin{
		Plugin: BuiltIn(),
	})
	if !s.AssertGRPCStatusContains(err, codes.Unknown, "k8s-bundle: IdentityProvider host service is required") {
		p.Close()
	}
}

func (s *Suite) withKubeClient(client kubeClient, expectedConfigPath string) {
	s.raw.hooks.newKubeClient = func(c *pluginConfig) ([]kubeClient, error) {
		s.Equal(expectedConfigPath, c.KubeConfigFilePath)
		if client == nil {
			return nil, errors.New("kube client not configured")
		}
		return []kubeClient{client}, nil
	}
}

func (s *Suite) configure(configuration string) {
	_, err := s.p.Configure(context.Background(), &spi.ConfigureRequest{
		Configuration: configuration,
	})
	s.Require().NoError(err)
}

type fakeKubeClient struct {
	mu         sync.RWMutex
	configMaps map[string]*corev1.ConfigMap
	patchErr   error
}

func newFakeKubeClient(configMaps ...*corev1.ConfigMap) *fakeKubeClient {
	c := &fakeKubeClient{
		configMaps: make(map[string]*corev1.ConfigMap),
	}
	for _, configMap := range configMaps {
		c.setConfigMap(configMap)
	}
	return c
}

func (c *fakeKubeClient) Get(ctx context.Context, namespace, configMap string) (runtime.Object, error) {
	entry := c.getConfigMap(namespace, configMap)
	if entry == nil {
		return nil, errors.New("not found")
	}
	return entry, nil
}
func (c *fakeKubeClient) GetList(ctx context.Context, config *pluginConfig) (runtime.Object, error) {
	list := c.getConfigMapList()
	if list.Items == nil {
		return nil, errors.New("not found")
	}
	return list, nil
}

func (c *fakeKubeClient) CreatePatch(ctx context.Context, config *pluginConfig, obj runtime.Object, resp *hostservices.FetchX509IdentityResponse) (runtime.Object, error) {
	configMap, ok := obj.(*corev1.ConfigMap)
	if !ok {
		return nil, k8sErr.New("wrong type, expecting config map")
	}
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			ResourceVersion: configMap.ResourceVersion,
		},
		Data: map[string]string{
			config.ConfigMapKey: bundleData(resp.Bundle),
		},
	}, nil
}

func (c *fakeKubeClient) Patch(ctx context.Context, namespace, configMap string, patchBytes []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.configMaps[configMapKey(namespace, configMap)]
	if !ok {
		return errors.New("not found")
	}

	// if there is a patch error configured, return it and clear the patchErr state.
	patchErr := c.patchErr
	c.patchErr = nil
	if patchErr != nil {
		return patchErr
	}

	patchedMap := new(corev1.ConfigMap)
	if err := json.Unmarshal(patchBytes, patchedMap); err != nil {
		return err
	}
	resourceVersion, err := strconv.Atoi(patchedMap.ResourceVersion)
	if err != nil {
		return errors.New("patch does not have resource version")
	}
	entry.ResourceVersion = fmt.Sprint(resourceVersion + 1)
	if entry.Data == nil {
		entry.Data = map[string]string{}
	}
	for key, data := range patchedMap.Data {
		entry.Data[key] = data
	}
	return nil
}

func (c *fakeKubeClient) Watch(ctx context.Context, config *pluginConfig) (watch.Interface, error) {
	return nil, nil
}

func (c *fakeKubeClient) getConfigMap(namespace, configMap string) *corev1.ConfigMap {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.configMaps[configMapKey(namespace, configMap)]
}

func (c *fakeKubeClient) getConfigMapList() *corev1.ConfigMapList {
	c.mu.RLock()
	defer c.mu.RUnlock()
	configMapList := &corev1.ConfigMapList{}
	for _, configMap := range c.configMaps {
		configMapList.Items = append(configMapList.Items, *configMap)
	}
	return configMapList
}

func (c *fakeKubeClient) setConfigMap(configMap *corev1.ConfigMap) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.configMaps[configMapKey(configMap.Namespace, configMap.Name)] = configMap
}

func (c *fakeKubeClient) setPatchErr(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.patchErr = err
}

func configMapKey(namespace, configMap string) string {
	return fmt.Sprintf("%s|%s", namespace, configMap)
}

func newConfigMap() *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       "spire",
			Name:            "spire-bundle",
			ResourceVersion: "1",
		},
	}
}
