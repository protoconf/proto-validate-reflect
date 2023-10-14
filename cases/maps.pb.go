// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: maps.proto

package cases

import (
	_ "github.com/protoconf/proto-validate-reflect/validate"
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

type MapNone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[uint32]bool `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *MapNone) Reset() {
	*x = MapNone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapNone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapNone) ProtoMessage() {}

func (x *MapNone) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapNone.ProtoReflect.Descriptor instead.
func (*MapNone) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{0}
}

func (x *MapNone) GetVal() map[uint32]bool {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapMin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[int32]float32 `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
}

func (x *MapMin) Reset() {
	*x = MapMin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapMin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMin) ProtoMessage() {}

func (x *MapMin) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapMin.ProtoReflect.Descriptor instead.
func (*MapMin) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{1}
}

func (x *MapMin) GetVal() map[int32]float32 {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapMax struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[int64]float64 `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *MapMax) Reset() {
	*x = MapMax{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapMax) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMax) ProtoMessage() {}

func (x *MapMax) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapMax.ProtoReflect.Descriptor instead.
func (*MapMax) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{2}
}

func (x *MapMax) GetVal() map[int64]float64 {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapMinMax struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[string]bool `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *MapMinMax) Reset() {
	*x = MapMinMax{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapMinMax) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapMinMax) ProtoMessage() {}

func (x *MapMinMax) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapMinMax.ProtoReflect.Descriptor instead.
func (*MapMinMax) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{3}
}

func (x *MapMinMax) GetVal() map[string]bool {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapExact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[uint64]string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapExact) Reset() {
	*x = MapExact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapExact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapExact) ProtoMessage() {}

func (x *MapExact) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapExact.ProtoReflect.Descriptor instead.
func (*MapExact) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{4}
}

func (x *MapExact) GetVal() map[uint64]string {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapNoSparse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[uint32]*MapNoSparse_Msg `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapNoSparse) Reset() {
	*x = MapNoSparse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapNoSparse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapNoSparse) ProtoMessage() {}

func (x *MapNoSparse) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapNoSparse.ProtoReflect.Descriptor instead.
func (*MapNoSparse) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{5}
}

func (x *MapNoSparse) GetVal() map[uint32]*MapNoSparse_Msg {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapKeys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[int64]string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"zigzag64,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapKeys) Reset() {
	*x = MapKeys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapKeys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapKeys) ProtoMessage() {}

func (x *MapKeys) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapKeys.ProtoReflect.Descriptor instead.
func (*MapKeys) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{6}
}

func (x *MapKeys) GetVal() map[int64]string {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[string]string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapValues) Reset() {
	*x = MapValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapValues) ProtoMessage() {}

func (x *MapValues) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapValues.ProtoReflect.Descriptor instead.
func (*MapValues) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{7}
}

func (x *MapValues) GetVal() map[string]string {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapKeysPattern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[string]string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapKeysPattern) Reset() {
	*x = MapKeysPattern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapKeysPattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapKeysPattern) ProtoMessage() {}

func (x *MapKeysPattern) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapKeysPattern.ProtoReflect.Descriptor instead.
func (*MapKeysPattern) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{8}
}

func (x *MapKeysPattern) GetVal() map[string]string {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapValuesPattern struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[string]string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapValuesPattern) Reset() {
	*x = MapValuesPattern{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapValuesPattern) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapValuesPattern) ProtoMessage() {}

func (x *MapValuesPattern) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapValuesPattern.ProtoReflect.Descriptor instead.
func (*MapValuesPattern) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{9}
}

func (x *MapValuesPattern) GetVal() map[string]string {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapRecursive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[uint32]*MapRecursive_Msg `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapRecursive) Reset() {
	*x = MapRecursive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapRecursive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRecursive) ProtoMessage() {}

func (x *MapRecursive) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapRecursive.ProtoReflect.Descriptor instead.
func (*MapRecursive) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{10}
}

func (x *MapRecursive) GetVal() map[uint32]*MapRecursive_Msg {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapExactIgnore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val map[uint64]string `protobuf:"bytes,1,rep,name=val,proto3" json:"val,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MapExactIgnore) Reset() {
	*x = MapExactIgnore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapExactIgnore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapExactIgnore) ProtoMessage() {}

func (x *MapExactIgnore) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapExactIgnore.ProtoReflect.Descriptor instead.
func (*MapExactIgnore) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{11}
}

func (x *MapExactIgnore) GetVal() map[uint64]string {
	if x != nil {
		return x.Val
	}
	return nil
}

type MapNoSparse_Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MapNoSparse_Msg) Reset() {
	*x = MapNoSparse_Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapNoSparse_Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapNoSparse_Msg) ProtoMessage() {}

func (x *MapNoSparse_Msg) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapNoSparse_Msg.ProtoReflect.Descriptor instead.
func (*MapNoSparse_Msg) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{5, 1}
}

type MapRecursive_Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val string `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *MapRecursive_Msg) Reset() {
	*x = MapRecursive_Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_maps_proto_msgTypes[24]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapRecursive_Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapRecursive_Msg) ProtoMessage() {}

func (x *MapRecursive_Msg) ProtoReflect() protoreflect.Message {
	mi := &file_maps_proto_msgTypes[24]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapRecursive_Msg.ProtoReflect.Descriptor instead.
func (*MapRecursive_Msg) Descriptor() ([]byte, []int) {
	return file_maps_proto_rawDescGZIP(), []int{10, 1}
}

func (x *MapRecursive_Msg) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

var File_maps_proto protoreflect.FileDescriptor

var file_maps_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x07, 0x4d, 0x61,
	0x70, 0x4e, 0x6f, 0x6e, 0x65, 0x12, 0x37, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4e, 0x6f, 0x6e, 0x65,
	0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x36,
	0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x82, 0x01, 0x0a, 0x06, 0x4d, 0x61, 0x70, 0x4d, 0x69,
	0x6e, 0x12, 0x40, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4d, 0x69, 0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x9a, 0x01, 0x02, 0x08, 0x02, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x82, 0x01, 0x0a, 0x06,
	0x4d, 0x61, 0x70, 0x4d, 0x61, 0x78, 0x12, 0x40, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e,
	0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4d, 0x61, 0x78,
	0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x9a, 0x01,
	0x02, 0x10, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x8a, 0x01, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x4d, 0x69, 0x6e, 0x4d, 0x61, 0x78, 0x12, 0x45,
	0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65,
	0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4d, 0x69, 0x6e, 0x4d, 0x61, 0x78, 0x2e, 0x56, 0x61, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x9a, 0x01, 0x04, 0x08, 0x02, 0x10, 0x04,
	0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x88, 0x01,
	0x0a, 0x08, 0x4d, 0x61, 0x70, 0x45, 0x78, 0x61, 0x63, 0x74, 0x12, 0x44, 0x0a, 0x03, 0x76, 0x61,
	0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e,
	0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4d, 0x61,
	0x70, 0x45, 0x78, 0x61, 0x63, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42,
	0x0a, 0xfa, 0x42, 0x07, 0x9a, 0x01, 0x04, 0x08, 0x03, 0x10, 0x03, 0x52, 0x03, 0x76, 0x61, 0x6c,
	0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb9, 0x01, 0x0a, 0x0b, 0x4d, 0x61, 0x70,
	0x4e, 0x6f, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61,
	0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4e,
	0x6f, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0x9a, 0x01, 0x02, 0x18, 0x01, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a,
	0x5c, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3a, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x73,
	0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4e, 0x6f, 0x53, 0x70, 0x61, 0x72, 0x73, 0x65, 0x2e, 0x4d,
	0x73, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x05, 0x0a,
	0x03, 0x4d, 0x73, 0x67, 0x22, 0x88, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x73,
	0x12, 0x45, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x9a, 0x01, 0x06, 0x22, 0x04, 0x42, 0x02,
	0x10, 0x00, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x8c, 0x01, 0x0a, 0x09, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x47, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2e, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x42, 0x0c, 0xfa, 0x42, 0x09, 0x9a, 0x01, 0x06, 0x2a, 0x04, 0x72, 0x02, 0x10,
	0x03, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa5,
	0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x73, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x6e, 0x12, 0x5b, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x4b, 0x65, 0x79, 0x73, 0x50, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x1b, 0xfa, 0x42,
	0x18, 0x9a, 0x01, 0x15, 0x22, 0x13, 0x72, 0x11, 0x32, 0x0f, 0x28, 0x3f, 0x69, 0x29, 0x5e, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x2b, 0x24, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x36,
	0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xa9, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x70, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x5d, 0x0a, 0x03, 0x76,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4d,
	0x61, 0x70, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x2e,
	0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x1b, 0xfa, 0x42, 0x18, 0x9a, 0x01, 0x15,
	0x2a, 0x13, 0x72, 0x11, 0x32, 0x0f, 0x28, 0x3f, 0x69, 0x29, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x5d, 0x2b, 0x24, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61,
	0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xcd, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73,
	0x69, 0x76, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73,
	0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73,
	0x69, 0x76, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x1a, 0x5d, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x63, 0x75, 0x72, 0x73, 0x69, 0x76,
	0x65, 0x2e, 0x4d, 0x73, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x20, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x19, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x03, 0x52, 0x03, 0x76,
	0x61, 0x6c, 0x22, 0x96, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x70, 0x45, 0x78, 0x61, 0x63, 0x74, 0x49,
	0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x4c, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2e, 0x68, 0x61, 0x72, 0x6e, 0x65,
	0x73, 0x73, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x45, 0x78, 0x61, 0x63,
	0x74, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x0c, 0xfa, 0x42, 0x09, 0x9a, 0x01, 0x06, 0x08, 0x03, 0x10, 0x03, 0x30, 0x01, 0x52, 0x03,
	0x76, 0x61, 0x6c, 0x1a, 0x36, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6e, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2d, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_maps_proto_rawDescOnce sync.Once
	file_maps_proto_rawDescData = file_maps_proto_rawDesc
)

func file_maps_proto_rawDescGZIP() []byte {
	file_maps_proto_rawDescOnce.Do(func() {
		file_maps_proto_rawDescData = protoimpl.X.CompressGZIP(file_maps_proto_rawDescData)
	})
	return file_maps_proto_rawDescData
}

var file_maps_proto_msgTypes = make([]protoimpl.MessageInfo, 26)
var file_maps_proto_goTypes = []interface{}{
	(*MapNone)(nil),          // 0: tests.harness.cases.MapNone
	(*MapMin)(nil),           // 1: tests.harness.cases.MapMin
	(*MapMax)(nil),           // 2: tests.harness.cases.MapMax
	(*MapMinMax)(nil),        // 3: tests.harness.cases.MapMinMax
	(*MapExact)(nil),         // 4: tests.harness.cases.MapExact
	(*MapNoSparse)(nil),      // 5: tests.harness.cases.MapNoSparse
	(*MapKeys)(nil),          // 6: tests.harness.cases.MapKeys
	(*MapValues)(nil),        // 7: tests.harness.cases.MapValues
	(*MapKeysPattern)(nil),   // 8: tests.harness.cases.MapKeysPattern
	(*MapValuesPattern)(nil), // 9: tests.harness.cases.MapValuesPattern
	(*MapRecursive)(nil),     // 10: tests.harness.cases.MapRecursive
	(*MapExactIgnore)(nil),   // 11: tests.harness.cases.MapExactIgnore
	nil,                      // 12: tests.harness.cases.MapNone.ValEntry
	nil,                      // 13: tests.harness.cases.MapMin.ValEntry
	nil,                      // 14: tests.harness.cases.MapMax.ValEntry
	nil,                      // 15: tests.harness.cases.MapMinMax.ValEntry
	nil,                      // 16: tests.harness.cases.MapExact.ValEntry
	nil,                      // 17: tests.harness.cases.MapNoSparse.ValEntry
	(*MapNoSparse_Msg)(nil),  // 18: tests.harness.cases.MapNoSparse.Msg
	nil,                      // 19: tests.harness.cases.MapKeys.ValEntry
	nil,                      // 20: tests.harness.cases.MapValues.ValEntry
	nil,                      // 21: tests.harness.cases.MapKeysPattern.ValEntry
	nil,                      // 22: tests.harness.cases.MapValuesPattern.ValEntry
	nil,                      // 23: tests.harness.cases.MapRecursive.ValEntry
	(*MapRecursive_Msg)(nil), // 24: tests.harness.cases.MapRecursive.Msg
	nil,                      // 25: tests.harness.cases.MapExactIgnore.ValEntry
}
var file_maps_proto_depIdxs = []int32{
	12, // 0: tests.harness.cases.MapNone.val:type_name -> tests.harness.cases.MapNone.ValEntry
	13, // 1: tests.harness.cases.MapMin.val:type_name -> tests.harness.cases.MapMin.ValEntry
	14, // 2: tests.harness.cases.MapMax.val:type_name -> tests.harness.cases.MapMax.ValEntry
	15, // 3: tests.harness.cases.MapMinMax.val:type_name -> tests.harness.cases.MapMinMax.ValEntry
	16, // 4: tests.harness.cases.MapExact.val:type_name -> tests.harness.cases.MapExact.ValEntry
	17, // 5: tests.harness.cases.MapNoSparse.val:type_name -> tests.harness.cases.MapNoSparse.ValEntry
	19, // 6: tests.harness.cases.MapKeys.val:type_name -> tests.harness.cases.MapKeys.ValEntry
	20, // 7: tests.harness.cases.MapValues.val:type_name -> tests.harness.cases.MapValues.ValEntry
	21, // 8: tests.harness.cases.MapKeysPattern.val:type_name -> tests.harness.cases.MapKeysPattern.ValEntry
	22, // 9: tests.harness.cases.MapValuesPattern.val:type_name -> tests.harness.cases.MapValuesPattern.ValEntry
	23, // 10: tests.harness.cases.MapRecursive.val:type_name -> tests.harness.cases.MapRecursive.ValEntry
	25, // 11: tests.harness.cases.MapExactIgnore.val:type_name -> tests.harness.cases.MapExactIgnore.ValEntry
	18, // 12: tests.harness.cases.MapNoSparse.ValEntry.value:type_name -> tests.harness.cases.MapNoSparse.Msg
	24, // 13: tests.harness.cases.MapRecursive.ValEntry.value:type_name -> tests.harness.cases.MapRecursive.Msg
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_maps_proto_init() }
func file_maps_proto_init() {
	if File_maps_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_maps_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapNone); i {
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
		file_maps_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapMin); i {
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
		file_maps_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapMax); i {
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
		file_maps_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapMinMax); i {
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
		file_maps_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapExact); i {
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
		file_maps_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapNoSparse); i {
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
		file_maps_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapKeys); i {
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
		file_maps_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapValues); i {
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
		file_maps_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapKeysPattern); i {
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
		file_maps_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapValuesPattern); i {
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
		file_maps_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapRecursive); i {
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
		file_maps_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapExactIgnore); i {
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
		file_maps_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapNoSparse_Msg); i {
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
		file_maps_proto_msgTypes[24].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapRecursive_Msg); i {
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
			RawDescriptor: file_maps_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   26,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_maps_proto_goTypes,
		DependencyIndexes: file_maps_proto_depIdxs,
		MessageInfos:      file_maps_proto_msgTypes,
	}.Build()
	File_maps_proto = out.File
	file_maps_proto_rawDesc = nil
	file_maps_proto_goTypes = nil
	file_maps_proto_depIdxs = nil
}
