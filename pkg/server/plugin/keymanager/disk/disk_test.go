package disk_test

import (
	"context"
	"crypto/x509"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spiffe/spire/pkg/common/plugin"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager"
	"github.com/spiffe/spire/pkg/server/plugin/keymanager/disk"
	keymanagertest "github.com/spiffe/spire/pkg/server/plugin/keymanager/test"
	spi "github.com/spiffe/spire/proto/spire/common/plugin"
	keymanagerv0 "github.com/spiffe/spire/proto/spire/server/keymanager/v0"
	"github.com/spiffe/spire/test/spiretest"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
)

func TestKeyManagerContract(t *testing.T) {
	keymanagertest.Test(t, func(t *testing.T) keymanager.KeyManager {
		dir := spiretest.TempDir(t)
		km, err := loadPlugin(t, "keys_path = %q", filepath.Join(dir, "keys.json"))
		require.NoError(t, err)
		return km
	})
}

func TestConfigure(t *testing.T) {
	t.Run("missing keys path", func(t *testing.T) {
		_, err := loadPlugin(t, "")
		spiretest.RequireGRPCStatus(t, err, codes.InvalidArgument, "keys_path is required")
	})
}

func TestGenerateKeyBeforeConfigure(t *testing.T) {
	var km keymanager.V0
	spiretest.LoadPlugin(t, disk.BuiltIn(), &km)

	_, err := km.GenerateKey(context.Background(), "id", keymanager.ECP256)
	spiretest.RequireGRPCStatus(t, err, codes.FailedPrecondition, "keymanager(disk): not configured")
}

func TestGenerateKeyPersistence(t *testing.T) {
	dir := filepath.Join(spiretest.TempDir(t), "no-such-dir")

	km, err := loadPlugin(t, "keys_path = %q", filepath.Join(dir, "keys.json"))
	require.NoError(t, err)

	// assert failure to generate key when directory is gone
	_, err = km.GenerateKey(context.Background(), "id", keymanager.ECP256)
	spiretest.RequireGRPCStatusContains(t, err, codes.Internal, "unable to write entries")

	// create the directory and generate the key
	mkdir(t, dir)
	keyIn, err := km.GenerateKey(context.Background(), "id", keymanager.ECP256)
	require.NoError(t, err)

	// reload the plugin. original key should have persisted.
	km, err = loadPlugin(t, "keys_path = %q", filepath.Join(dir, "keys.json"))
	require.NoError(t, err)
	keyOut, err := km.GetKey(context.Background(), "id")
	require.NoError(t, err)
	require.Equal(t,
		publicKeyBytes(t, keyIn),
		publicKeyBytes(t, keyOut),
	)

	// remove the directory and try to overwrite. original key should remain.
	rmdir(t, dir)
	_, err = km.GenerateKey(context.Background(), "id", keymanager.ECP256)
	spiretest.RequireGRPCStatusContains(t, err, codes.Internal, "unable to write entries")

	keyOut, err = km.GetKey(context.Background(), "id")
	require.NoError(t, err)
	require.Equal(t,
		publicKeyBytes(t, keyIn),
		publicKeyBytes(t, keyOut),
	)
}

func loadPlugin(t *testing.T, configFmt string, configArgs ...interface{}) (keymanager.KeyManager, error) {
	// This little workaround to get at the configuration interface
	// won't be required after the catalog system refactor
	km := struct {
		plugin.Facade
		keymanagerv0.Plugin
	}{}

	spiretest.LoadPlugin(t, disk.BuiltIn(), &km)

	_, err := km.Configure(context.Background(), &spi.ConfigureRequest{
		Configuration: fmt.Sprintf(configFmt, configArgs...),
	})
	return keymanager.V0{
		Facade: km.Facade,
		Plugin: km.Plugin,
	}, err
}

func mkdir(t *testing.T, dir string) {
	require.NoError(t, os.Mkdir(dir, 0755))
}

func rmdir(t *testing.T, dir string) {
	require.NoError(t, os.RemoveAll(dir))
}

func publicKeyBytes(t *testing.T, key keymanager.Key) []byte {
	b, err := x509.MarshalPKIXPublicKey(key.Public())
	require.NoError(t, err)
	return b
}
