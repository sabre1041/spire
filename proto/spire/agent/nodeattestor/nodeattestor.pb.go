//* Responsible for attesting the physical nodes identity.  The plugin
//will be responsible to retrieve an identity document or data associated
//with the physical node.  This data will be used when attesting to the server.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: spire/agent/nodeattestor/nodeattestor.proto

package nodeattestor

import (
	common "github.com/spiffe/spire/proto/spire/common"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//* Represents an empty request
type FetchAttestationDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge []byte `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
}

func (x *FetchAttestationDataRequest) Reset() {
	*x = FetchAttestationDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_agent_nodeattestor_nodeattestor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAttestationDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAttestationDataRequest) ProtoMessage() {}

func (x *FetchAttestationDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_agent_nodeattestor_nodeattestor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAttestationDataRequest.ProtoReflect.Descriptor instead.
func (*FetchAttestationDataRequest) Descriptor() ([]byte, []int) {
	return file_spire_agent_nodeattestor_nodeattestor_proto_rawDescGZIP(), []int{0}
}

func (x *FetchAttestationDataRequest) GetChallenge() []byte {
	if x != nil {
		return x.Challenge
	}
	return nil
}

//* Represents the attested data and base SPIFFE ID
type FetchAttestationDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* A type which contains attestation data for specific platform
	AttestationData *common.AttestationData `protobuf:"bytes,1,opt,name=attestation_data,json=attestationData,proto3" json:"attestation_data,omitempty"`
	//* response to the challenge (if challenge was present) *
	Response []byte `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *FetchAttestationDataResponse) Reset() {
	*x = FetchAttestationDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_agent_nodeattestor_nodeattestor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAttestationDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAttestationDataResponse) ProtoMessage() {}

func (x *FetchAttestationDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_agent_nodeattestor_nodeattestor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAttestationDataResponse.ProtoReflect.Descriptor instead.
func (*FetchAttestationDataResponse) Descriptor() ([]byte, []int) {
	return file_spire_agent_nodeattestor_nodeattestor_proto_rawDescGZIP(), []int{1}
}

func (x *FetchAttestationDataResponse) GetAttestationData() *common.AttestationData {
	if x != nil {
		return x.AttestationData
	}
	return nil
}

func (x *FetchAttestationDataResponse) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_spire_agent_nodeattestor_nodeattestor_proto protoreflect.FileDescriptor

var file_spire_agent_nodeattestor_nodeattestor_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x6e, 0x6f,
	0x64, 0x65, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61,
	0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x1a, 0x19, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x1b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x22, 0x8a, 0x01, 0x0a, 0x1c, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x48, 0x0a, 0x10, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0f, 0x61, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x32, 0xde,
	0x02, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x12,
	0x89, 0x01, 0x0a, 0x14, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x35, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x61, 0x74, 0x74, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x36, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x5a, 0x0a, 0x09, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12, 0x25, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70,
	0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_spire_agent_nodeattestor_nodeattestor_proto_rawDescOnce sync.Once
	file_spire_agent_nodeattestor_nodeattestor_proto_rawDescData = file_spire_agent_nodeattestor_nodeattestor_proto_rawDesc
)

func file_spire_agent_nodeattestor_nodeattestor_proto_rawDescGZIP() []byte {
	file_spire_agent_nodeattestor_nodeattestor_proto_rawDescOnce.Do(func() {
		file_spire_agent_nodeattestor_nodeattestor_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_agent_nodeattestor_nodeattestor_proto_rawDescData)
	})
	return file_spire_agent_nodeattestor_nodeattestor_proto_rawDescData
}

var file_spire_agent_nodeattestor_nodeattestor_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_spire_agent_nodeattestor_nodeattestor_proto_goTypes = []interface{}{
	(*FetchAttestationDataRequest)(nil),  // 0: spire.agent.nodeattestor.FetchAttestationDataRequest
	(*FetchAttestationDataResponse)(nil), // 1: spire.agent.nodeattestor.FetchAttestationDataResponse
	(*common.AttestationData)(nil),       // 2: spire.common.AttestationData
	(*plugin.ConfigureRequest)(nil),      // 3: spire.common.plugin.ConfigureRequest
	(*plugin.GetPluginInfoRequest)(nil),  // 4: spire.common.plugin.GetPluginInfoRequest
	(*plugin.ConfigureResponse)(nil),     // 5: spire.common.plugin.ConfigureResponse
	(*plugin.GetPluginInfoResponse)(nil), // 6: spire.common.plugin.GetPluginInfoResponse
}
var file_spire_agent_nodeattestor_nodeattestor_proto_depIdxs = []int32{
	2, // 0: spire.agent.nodeattestor.FetchAttestationDataResponse.attestation_data:type_name -> spire.common.AttestationData
	0, // 1: spire.agent.nodeattestor.NodeAttestor.FetchAttestationData:input_type -> spire.agent.nodeattestor.FetchAttestationDataRequest
	3, // 2: spire.agent.nodeattestor.NodeAttestor.Configure:input_type -> spire.common.plugin.ConfigureRequest
	4, // 3: spire.agent.nodeattestor.NodeAttestor.GetPluginInfo:input_type -> spire.common.plugin.GetPluginInfoRequest
	1, // 4: spire.agent.nodeattestor.NodeAttestor.FetchAttestationData:output_type -> spire.agent.nodeattestor.FetchAttestationDataResponse
	5, // 5: spire.agent.nodeattestor.NodeAttestor.Configure:output_type -> spire.common.plugin.ConfigureResponse
	6, // 6: spire.agent.nodeattestor.NodeAttestor.GetPluginInfo:output_type -> spire.common.plugin.GetPluginInfoResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spire_agent_nodeattestor_nodeattestor_proto_init() }
func file_spire_agent_nodeattestor_nodeattestor_proto_init() {
	if File_spire_agent_nodeattestor_nodeattestor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_agent_nodeattestor_nodeattestor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAttestationDataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_spire_agent_nodeattestor_nodeattestor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAttestationDataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spire_agent_nodeattestor_nodeattestor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_agent_nodeattestor_nodeattestor_proto_goTypes,
		DependencyIndexes: file_spire_agent_nodeattestor_nodeattestor_proto_depIdxs,
		MessageInfos:      file_spire_agent_nodeattestor_nodeattestor_proto_msgTypes,
	}.Build()
	File_spire_agent_nodeattestor_nodeattestor_proto = out.File
	file_spire_agent_nodeattestor_nodeattestor_proto_rawDesc = nil
	file_spire_agent_nodeattestor_nodeattestor_proto_goTypes = nil
	file_spire_agent_nodeattestor_nodeattestor_proto_depIdxs = nil
}
