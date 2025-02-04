// A Notifier plugin reacts to various server related events

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: spire/server/notifier/notifier.proto

package notifier

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

type BundleLoaded struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bundle *common.Bundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
}

func (x *BundleLoaded) Reset() {
	*x = BundleLoaded{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_server_notifier_notifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BundleLoaded) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleLoaded) ProtoMessage() {}

func (x *BundleLoaded) ProtoReflect() protoreflect.Message {
	mi := &file_spire_server_notifier_notifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleLoaded.ProtoReflect.Descriptor instead.
func (*BundleLoaded) Descriptor() ([]byte, []int) {
	return file_spire_server_notifier_notifier_proto_rawDescGZIP(), []int{0}
}

func (x *BundleLoaded) GetBundle() *common.Bundle {
	if x != nil {
		return x.Bundle
	}
	return nil
}

type BundleUpdated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bundle *common.Bundle `protobuf:"bytes,1,opt,name=bundle,proto3" json:"bundle,omitempty"`
}

func (x *BundleUpdated) Reset() {
	*x = BundleUpdated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_server_notifier_notifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BundleUpdated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BundleUpdated) ProtoMessage() {}

func (x *BundleUpdated) ProtoReflect() protoreflect.Message {
	mi := &file_spire_server_notifier_notifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BundleUpdated.ProtoReflect.Descriptor instead.
func (*BundleUpdated) Descriptor() ([]byte, []int) {
	return file_spire_server_notifier_notifier_proto_rawDescGZIP(), []int{1}
}

func (x *BundleUpdated) GetBundle() *common.Bundle {
	if x != nil {
		return x.Bundle
	}
	return nil
}

type NotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*NotifyRequest_BundleUpdated
	Event isNotifyRequest_Event `protobuf_oneof:"event"`
}

func (x *NotifyRequest) Reset() {
	*x = NotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_server_notifier_notifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyRequest) ProtoMessage() {}

func (x *NotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_server_notifier_notifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyRequest.ProtoReflect.Descriptor instead.
func (*NotifyRequest) Descriptor() ([]byte, []int) {
	return file_spire_server_notifier_notifier_proto_rawDescGZIP(), []int{2}
}

func (m *NotifyRequest) GetEvent() isNotifyRequest_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *NotifyRequest) GetBundleUpdated() *BundleUpdated {
	if x, ok := x.GetEvent().(*NotifyRequest_BundleUpdated); ok {
		return x.BundleUpdated
	}
	return nil
}

type isNotifyRequest_Event interface {
	isNotifyRequest_Event()
}

type NotifyRequest_BundleUpdated struct {
	// BundleUpdated is emitted whenever SPIRE server changes the trust
	// bundle.
	BundleUpdated *BundleUpdated `protobuf:"bytes,1,opt,name=bundle_updated,json=bundleUpdated,proto3,oneof"`
}

func (*NotifyRequest_BundleUpdated) isNotifyRequest_Event() {}

type NotifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyResponse) Reset() {
	*x = NotifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_server_notifier_notifier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyResponse) ProtoMessage() {}

func (x *NotifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_server_notifier_notifier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyResponse.ProtoReflect.Descriptor instead.
func (*NotifyResponse) Descriptor() ([]byte, []int) {
	return file_spire_server_notifier_notifier_proto_rawDescGZIP(), []int{3}
}

type NotifyAndAdviseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*NotifyAndAdviseRequest_BundleLoaded
	Event isNotifyAndAdviseRequest_Event `protobuf_oneof:"event"`
}

func (x *NotifyAndAdviseRequest) Reset() {
	*x = NotifyAndAdviseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_server_notifier_notifier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyAndAdviseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyAndAdviseRequest) ProtoMessage() {}

func (x *NotifyAndAdviseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_server_notifier_notifier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyAndAdviseRequest.ProtoReflect.Descriptor instead.
func (*NotifyAndAdviseRequest) Descriptor() ([]byte, []int) {
	return file_spire_server_notifier_notifier_proto_rawDescGZIP(), []int{4}
}

func (m *NotifyAndAdviseRequest) GetEvent() isNotifyAndAdviseRequest_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *NotifyAndAdviseRequest) GetBundleLoaded() *BundleLoaded {
	if x, ok := x.GetEvent().(*NotifyAndAdviseRequest_BundleLoaded); ok {
		return x.BundleLoaded
	}
	return nil
}

type isNotifyAndAdviseRequest_Event interface {
	isNotifyAndAdviseRequest_Event()
}

type NotifyAndAdviseRequest_BundleLoaded struct {
	// BundleLoaded is emitted on startup after SPIRE server creates/loads
	// the trust bundle. If an error is returned SPIRE server is shut down.
	BundleLoaded *BundleLoaded `protobuf:"bytes,1,opt,name=bundle_loaded,json=bundleLoaded,proto3,oneof"`
}

func (*NotifyAndAdviseRequest_BundleLoaded) isNotifyAndAdviseRequest_Event() {}

type NotifyAndAdviseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyAndAdviseResponse) Reset() {
	*x = NotifyAndAdviseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_server_notifier_notifier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyAndAdviseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyAndAdviseResponse) ProtoMessage() {}

func (x *NotifyAndAdviseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_server_notifier_notifier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyAndAdviseResponse.ProtoReflect.Descriptor instead.
func (*NotifyAndAdviseResponse) Descriptor() ([]byte, []int) {
	return file_spire_server_notifier_notifier_proto_rawDescGZIP(), []int{5}
}

var File_spire_server_notifier_notifier_proto protoreflect.FileDescriptor

var file_spire_server_notifier_notifier_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x1a, 0x19, 0x73,
	0x70, 0x69, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x0c, 0x42, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x62, 0x75,
	0x6e, 0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x52, 0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x3d, 0x0a, 0x0d, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x62, 0x75, 0x6e,
	0x64, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x52,
	0x06, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x67, 0x0a, 0x0d, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0e, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0d, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6d, 0x0a, 0x16, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x6e, 0x64, 0x41,
	0x64, 0x76, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x0d,
	0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x5f, 0x6c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x75, 0x6e, 0x64,
	0x6c, 0x65, 0x4c, 0x6f, 0x61, 0x64, 0x65, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x22, 0x19, 0x0a, 0x17, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x6e, 0x64, 0x41, 0x64,
	0x76, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x97, 0x03, 0x0a,
	0x08, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x06, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x24, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x70, 0x0a, 0x0f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x41, 0x6e, 0x64, 0x41, 0x64, 0x76,
	0x69, 0x73, 0x65, 0x12, 0x2d, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x41, 0x6e, 0x64, 0x41, 0x64, 0x76, 0x69, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x41, 0x6e, 0x64, 0x41, 0x64, 0x76, 0x69, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5a, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x12,
	0x25, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x29, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x66, 0x66, 0x65, 0x2f, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_server_notifier_notifier_proto_rawDescOnce sync.Once
	file_spire_server_notifier_notifier_proto_rawDescData = file_spire_server_notifier_notifier_proto_rawDesc
)

func file_spire_server_notifier_notifier_proto_rawDescGZIP() []byte {
	file_spire_server_notifier_notifier_proto_rawDescOnce.Do(func() {
		file_spire_server_notifier_notifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_server_notifier_notifier_proto_rawDescData)
	})
	return file_spire_server_notifier_notifier_proto_rawDescData
}

var file_spire_server_notifier_notifier_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_spire_server_notifier_notifier_proto_goTypes = []interface{}{
	(*BundleLoaded)(nil),                 // 0: spire.server.notifier.BundleLoaded
	(*BundleUpdated)(nil),                // 1: spire.server.notifier.BundleUpdated
	(*NotifyRequest)(nil),                // 2: spire.server.notifier.NotifyRequest
	(*NotifyResponse)(nil),               // 3: spire.server.notifier.NotifyResponse
	(*NotifyAndAdviseRequest)(nil),       // 4: spire.server.notifier.NotifyAndAdviseRequest
	(*NotifyAndAdviseResponse)(nil),      // 5: spire.server.notifier.NotifyAndAdviseResponse
	(*common.Bundle)(nil),                // 6: spire.common.Bundle
	(*plugin.ConfigureRequest)(nil),      // 7: spire.common.plugin.ConfigureRequest
	(*plugin.GetPluginInfoRequest)(nil),  // 8: spire.common.plugin.GetPluginInfoRequest
	(*plugin.ConfigureResponse)(nil),     // 9: spire.common.plugin.ConfigureResponse
	(*plugin.GetPluginInfoResponse)(nil), // 10: spire.common.plugin.GetPluginInfoResponse
}
var file_spire_server_notifier_notifier_proto_depIdxs = []int32{
	6,  // 0: spire.server.notifier.BundleLoaded.bundle:type_name -> spire.common.Bundle
	6,  // 1: spire.server.notifier.BundleUpdated.bundle:type_name -> spire.common.Bundle
	1,  // 2: spire.server.notifier.NotifyRequest.bundle_updated:type_name -> spire.server.notifier.BundleUpdated
	0,  // 3: spire.server.notifier.NotifyAndAdviseRequest.bundle_loaded:type_name -> spire.server.notifier.BundleLoaded
	2,  // 4: spire.server.notifier.Notifier.Notify:input_type -> spire.server.notifier.NotifyRequest
	4,  // 5: spire.server.notifier.Notifier.NotifyAndAdvise:input_type -> spire.server.notifier.NotifyAndAdviseRequest
	7,  // 6: spire.server.notifier.Notifier.Configure:input_type -> spire.common.plugin.ConfigureRequest
	8,  // 7: spire.server.notifier.Notifier.GetPluginInfo:input_type -> spire.common.plugin.GetPluginInfoRequest
	3,  // 8: spire.server.notifier.Notifier.Notify:output_type -> spire.server.notifier.NotifyResponse
	5,  // 9: spire.server.notifier.Notifier.NotifyAndAdvise:output_type -> spire.server.notifier.NotifyAndAdviseResponse
	9,  // 10: spire.server.notifier.Notifier.Configure:output_type -> spire.common.plugin.ConfigureResponse
	10, // 11: spire.server.notifier.Notifier.GetPluginInfo:output_type -> spire.common.plugin.GetPluginInfoResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_spire_server_notifier_notifier_proto_init() }
func file_spire_server_notifier_notifier_proto_init() {
	if File_spire_server_notifier_notifier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_server_notifier_notifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BundleLoaded); i {
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
		file_spire_server_notifier_notifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BundleUpdated); i {
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
		file_spire_server_notifier_notifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyRequest); i {
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
		file_spire_server_notifier_notifier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyResponse); i {
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
		file_spire_server_notifier_notifier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyAndAdviseRequest); i {
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
		file_spire_server_notifier_notifier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyAndAdviseResponse); i {
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
	file_spire_server_notifier_notifier_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*NotifyRequest_BundleUpdated)(nil),
	}
	file_spire_server_notifier_notifier_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*NotifyAndAdviseRequest_BundleLoaded)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_spire_server_notifier_notifier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_server_notifier_notifier_proto_goTypes,
		DependencyIndexes: file_spire_server_notifier_notifier_proto_depIdxs,
		MessageInfos:      file_spire_server_notifier_notifier_proto_msgTypes,
	}.Build()
	File_spire_server_notifier_notifier_proto = out.File
	file_spire_server_notifier_notifier_proto_rawDesc = nil
	file_spire_server_notifier_notifier_proto_goTypes = nil
	file_spire_server_notifier_notifier_proto_depIdxs = nil
}
