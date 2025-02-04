//* A plugin which is responsible for generating and storing a key pair,
//optionally with a hardware-backed secret store.  It is used for generating
//the key pair for the Base SPIFFE Id of the Node Agent, and persisting
//that identity across restarts/reboots

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: spire/agent/keymanager/keymanager.proto

package keymanager

import (
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
type GenerateKeyPairRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenerateKeyPairRequest) Reset() {
	*x = GenerateKeyPairRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKeyPairRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKeyPairRequest) ProtoMessage() {}

func (x *GenerateKeyPairRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKeyPairRequest.ProtoReflect.Descriptor instead.
func (*GenerateKeyPairRequest) Descriptor() ([]byte, []int) {
	return file_spire_agent_keymanager_keymanager_proto_rawDescGZIP(), []int{0}
}

//* Represents a public and private key pair
type GenerateKeyPairResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* Public key
	PublicKey []byte `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	//* Private key
	PrivateKey []byte `protobuf:"bytes,2,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
}

func (x *GenerateKeyPairResponse) Reset() {
	*x = GenerateKeyPairResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateKeyPairResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateKeyPairResponse) ProtoMessage() {}

func (x *GenerateKeyPairResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateKeyPairResponse.ProtoReflect.Descriptor instead.
func (*GenerateKeyPairResponse) Descriptor() ([]byte, []int) {
	return file_spire_agent_keymanager_keymanager_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateKeyPairResponse) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *GenerateKeyPairResponse) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

//* Represents a private key
type StorePrivateKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* Private key
	PrivateKey []byte `protobuf:"bytes,1,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
}

func (x *StorePrivateKeyRequest) Reset() {
	*x = StorePrivateKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorePrivateKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorePrivateKeyRequest) ProtoMessage() {}

func (x *StorePrivateKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorePrivateKeyRequest.ProtoReflect.Descriptor instead.
func (*StorePrivateKeyRequest) Descriptor() ([]byte, []int) {
	return file_spire_agent_keymanager_keymanager_proto_rawDescGZIP(), []int{2}
}

func (x *StorePrivateKeyRequest) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

//* Represents an empty response
type StorePrivateKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StorePrivateKeyResponse) Reset() {
	*x = StorePrivateKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorePrivateKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorePrivateKeyResponse) ProtoMessage() {}

func (x *StorePrivateKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorePrivateKeyResponse.ProtoReflect.Descriptor instead.
func (*StorePrivateKeyResponse) Descriptor() ([]byte, []int) {
	return file_spire_agent_keymanager_keymanager_proto_rawDescGZIP(), []int{3}
}

//* Represents an empty request
type FetchPrivateKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchPrivateKeyRequest) Reset() {
	*x = FetchPrivateKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchPrivateKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPrivateKeyRequest) ProtoMessage() {}

func (x *FetchPrivateKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPrivateKeyRequest.ProtoReflect.Descriptor instead.
func (*FetchPrivateKeyRequest) Descriptor() ([]byte, []int) {
	return file_spire_agent_keymanager_keymanager_proto_rawDescGZIP(), []int{4}
}

//* Represents a private key
type FetchPrivateKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//* Private key
	PrivateKey []byte `protobuf:"bytes,1,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
}

func (x *FetchPrivateKeyResponse) Reset() {
	*x = FetchPrivateKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchPrivateKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchPrivateKeyResponse) ProtoMessage() {}

func (x *FetchPrivateKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_agent_keymanager_keymanager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchPrivateKeyResponse.ProtoReflect.Descriptor instead.
func (*FetchPrivateKeyResponse) Descriptor() ([]byte, []int) {
	return file_spire_agent_keymanager_keymanager_proto_rawDescGZIP(), []int{5}
}

func (x *FetchPrivateKeyResponse) GetPrivateKey() []byte {
	if x != nil {
		return x.PrivateKey
	}
	return nil
}

var File_spire_agent_keymanager_keymanager_proto protoreflect.FileDescriptor

var file_spire_agent_keymanager_keymanager_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x6b, 0x65,
	0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x73, 0x70, 0x69, 0x72, 0x65,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x1a, 0x20, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b,
	0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x57, 0x0a,
	0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x38, 0x0a, 0x16, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50,
	0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x22, 0x19, 0x0a, 0x17, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x46,
	0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x39, 0x0a, 0x17, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x32, 0xac, 0x04, 0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x72, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61,
	0x69, 0x72, 0x12, 0x2e, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x2e, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a, 0x0f, 0x46, 0x65, 0x74, 0x63, 0x68,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x2e, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x6b, 0x65, 0x79, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x43,
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
	0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70,
	0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x6b, 0x65, 0x79,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_agent_keymanager_keymanager_proto_rawDescOnce sync.Once
	file_spire_agent_keymanager_keymanager_proto_rawDescData = file_spire_agent_keymanager_keymanager_proto_rawDesc
)

func file_spire_agent_keymanager_keymanager_proto_rawDescGZIP() []byte {
	file_spire_agent_keymanager_keymanager_proto_rawDescOnce.Do(func() {
		file_spire_agent_keymanager_keymanager_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_agent_keymanager_keymanager_proto_rawDescData)
	})
	return file_spire_agent_keymanager_keymanager_proto_rawDescData
}

var file_spire_agent_keymanager_keymanager_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_spire_agent_keymanager_keymanager_proto_goTypes = []interface{}{
	(*GenerateKeyPairRequest)(nil),       // 0: spire.agent.keymanager.GenerateKeyPairRequest
	(*GenerateKeyPairResponse)(nil),      // 1: spire.agent.keymanager.GenerateKeyPairResponse
	(*StorePrivateKeyRequest)(nil),       // 2: spire.agent.keymanager.StorePrivateKeyRequest
	(*StorePrivateKeyResponse)(nil),      // 3: spire.agent.keymanager.StorePrivateKeyResponse
	(*FetchPrivateKeyRequest)(nil),       // 4: spire.agent.keymanager.FetchPrivateKeyRequest
	(*FetchPrivateKeyResponse)(nil),      // 5: spire.agent.keymanager.FetchPrivateKeyResponse
	(*plugin.ConfigureRequest)(nil),      // 6: spire.common.plugin.ConfigureRequest
	(*plugin.GetPluginInfoRequest)(nil),  // 7: spire.common.plugin.GetPluginInfoRequest
	(*plugin.ConfigureResponse)(nil),     // 8: spire.common.plugin.ConfigureResponse
	(*plugin.GetPluginInfoResponse)(nil), // 9: spire.common.plugin.GetPluginInfoResponse
}
var file_spire_agent_keymanager_keymanager_proto_depIdxs = []int32{
	0, // 0: spire.agent.keymanager.KeyManager.GenerateKeyPair:input_type -> spire.agent.keymanager.GenerateKeyPairRequest
	2, // 1: spire.agent.keymanager.KeyManager.StorePrivateKey:input_type -> spire.agent.keymanager.StorePrivateKeyRequest
	4, // 2: spire.agent.keymanager.KeyManager.FetchPrivateKey:input_type -> spire.agent.keymanager.FetchPrivateKeyRequest
	6, // 3: spire.agent.keymanager.KeyManager.Configure:input_type -> spire.common.plugin.ConfigureRequest
	7, // 4: spire.agent.keymanager.KeyManager.GetPluginInfo:input_type -> spire.common.plugin.GetPluginInfoRequest
	1, // 5: spire.agent.keymanager.KeyManager.GenerateKeyPair:output_type -> spire.agent.keymanager.GenerateKeyPairResponse
	3, // 6: spire.agent.keymanager.KeyManager.StorePrivateKey:output_type -> spire.agent.keymanager.StorePrivateKeyResponse
	5, // 7: spire.agent.keymanager.KeyManager.FetchPrivateKey:output_type -> spire.agent.keymanager.FetchPrivateKeyResponse
	8, // 8: spire.agent.keymanager.KeyManager.Configure:output_type -> spire.common.plugin.ConfigureResponse
	9, // 9: spire.agent.keymanager.KeyManager.GetPluginInfo:output_type -> spire.common.plugin.GetPluginInfoResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_spire_agent_keymanager_keymanager_proto_init() }
func file_spire_agent_keymanager_keymanager_proto_init() {
	if File_spire_agent_keymanager_keymanager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_agent_keymanager_keymanager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKeyPairRequest); i {
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
		file_spire_agent_keymanager_keymanager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateKeyPairResponse); i {
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
		file_spire_agent_keymanager_keymanager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorePrivateKeyRequest); i {
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
		file_spire_agent_keymanager_keymanager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorePrivateKeyResponse); i {
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
		file_spire_agent_keymanager_keymanager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchPrivateKeyRequest); i {
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
		file_spire_agent_keymanager_keymanager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchPrivateKeyResponse); i {
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
			RawDescriptor: file_spire_agent_keymanager_keymanager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_agent_keymanager_keymanager_proto_goTypes,
		DependencyIndexes: file_spire_agent_keymanager_keymanager_proto_depIdxs,
		MessageInfos:      file_spire_agent_keymanager_keymanager_proto_msgTypes,
	}.Build()
	File_spire_agent_keymanager_keymanager_proto = out.File
	file_spire_agent_keymanager_keymanager_proto_rawDesc = nil
	file_spire_agent_keymanager_keymanager_proto_goTypes = nil
	file_spire_agent_keymanager_keymanager_proto_depIdxs = nil
}
