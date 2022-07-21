// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.1
// source: grpc_hook.proto

package grpc_hook

import (
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

type ImportSetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adapter  string `protobuf:"bytes,1,opt,name=adapter,proto3" json:"adapter,omitempty"`
	Page     int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Pagesize int32  `protobuf:"varint,3,opt,name=pagesize,proto3" json:"pagesize,omitempty"`
}

func (x *ImportSetRequest) Reset() {
	*x = ImportSetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_hook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportSetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSetRequest) ProtoMessage() {}

func (x *ImportSetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_hook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportSetRequest.ProtoReflect.Descriptor instead.
func (*ImportSetRequest) Descriptor() ([]byte, []int) {
	return file_grpc_hook_proto_rawDescGZIP(), []int{0}
}

func (x *ImportSetRequest) GetAdapter() string {
	if x != nil {
		return x.Adapter
	}
	return ""
}

func (x *ImportSetRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ImportSetRequest) GetPagesize() int32 {
	if x != nil {
		return x.Pagesize
	}
	return 0
}

type Row struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Row) Reset() {
	*x = Row{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_hook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Row) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Row) ProtoMessage() {}

func (x *Row) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_hook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Row.ProtoReflect.Descriptor instead.
func (*Row) Descriptor() ([]byte, []int) {
	return file_grpc_hook_proto_rawDescGZIP(), []int{1}
}

func (x *Row) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type Field struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fieldname string `protobuf:"bytes,1,opt,name=fieldname,proto3" json:"fieldname,omitempty"`
	Fieldtype string `protobuf:"bytes,2,opt,name=fieldtype,proto3" json:"fieldtype,omitempty"`
}

func (x *Field) Reset() {
	*x = Field{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_hook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Field) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Field) ProtoMessage() {}

func (x *Field) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_hook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Field.ProtoReflect.Descriptor instead.
func (*Field) Descriptor() ([]byte, []int) {
	return file_grpc_hook_proto_rawDescGZIP(), []int{2}
}

func (x *Field) GetFieldname() string {
	if x != nil {
		return x.Fieldname
	}
	return ""
}

func (x *Field) GetFieldtype() string {
	if x != nil {
		return x.Fieldtype
	}
	return ""
}

type ImportSetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table  string   `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Fields []*Field `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	Rows   []*Row   `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *ImportSetResponse) Reset() {
	*x = ImportSetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_hook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportSetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSetResponse) ProtoMessage() {}

func (x *ImportSetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_hook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportSetResponse.ProtoReflect.Descriptor instead.
func (*ImportSetResponse) Descriptor() ([]byte, []int) {
	return file_grpc_hook_proto_rawDescGZIP(), []int{3}
}

func (x *ImportSetResponse) GetTable() string {
	if x != nil {
		return x.Table
	}
	return ""
}

func (x *ImportSetResponse) GetFields() []*Field {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *ImportSetResponse) GetRows() []*Row {
	if x != nil {
		return x.Rows
	}
	return nil
}

type ScriptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ScriptPath string            `protobuf:"bytes,1,opt,name=scriptPath,proto3" json:"scriptPath,omitempty"`
	Params     map[string]string `protobuf:"bytes,2,rep,name=params,proto3" json:"params,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ScriptRequest) Reset() {
	*x = ScriptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_hook_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScriptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScriptRequest) ProtoMessage() {}

func (x *ScriptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_hook_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScriptRequest.ProtoReflect.Descriptor instead.
func (*ScriptRequest) Descriptor() ([]byte, []int) {
	return file_grpc_hook_proto_rawDescGZIP(), []int{4}
}

func (x *ScriptRequest) GetScriptPath() string {
	if x != nil {
		return x.ScriptPath
	}
	return ""
}

func (x *ScriptRequest) GetParams() map[string]string {
	if x != nil {
		return x.Params
	}
	return nil
}

type ScriptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer map[string]string `protobuf:"bytes,1,rep,name=answer,proto3" json:"answer,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ScriptResponse) Reset() {
	*x = ScriptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_hook_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScriptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScriptResponse) ProtoMessage() {}

func (x *ScriptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_hook_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScriptResponse.ProtoReflect.Descriptor instead.
func (*ScriptResponse) Descriptor() ([]byte, []int) {
	return file_grpc_hook_proto_rawDescGZIP(), []int{5}
}

func (x *ScriptResponse) GetAnswer() map[string]string {
	if x != nil {
		return x.Answer
	}
	return nil
}

var File_grpc_hook_proto protoreflect.FileDescriptor

var file_grpc_hook_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x22, 0x5c, 0x0a, 0x10,
	0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x1d, 0x0a, 0x03, 0x52, 0x6f,
	0x77, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x05, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x74, 0x79, 0x70, 0x65, 0x22, 0x77,
	0x0a, 0x11, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x52, 0x06, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x52, 0x6f,
	0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0xa8, 0x01, 0x0a, 0x0d, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3c, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x8a, 0x01, 0x0a, 0x0e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x6f, 0x6f,
	0x6b, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0xa0, 0x01, 0x0a, 0x0b, 0x48, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x45, 0x0a, 0x0c, 0x4f, 0x6e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x12,
	0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x0b, 0x4f, 0x6e, 0x49, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x6f, 0x6f,
	0x6b, 0x2e, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x6d, 0x6f, 0x72, 0x74, 0x61, 0x7a, 0x61, 0x2f, 0x61, 0x63, 0x65, 0x71, 0x6c, 0x2f,
	0x62, 0x73, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x68, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_hook_proto_rawDescOnce sync.Once
	file_grpc_hook_proto_rawDescData = file_grpc_hook_proto_rawDesc
)

func file_grpc_hook_proto_rawDescGZIP() []byte {
	file_grpc_hook_proto_rawDescOnce.Do(func() {
		file_grpc_hook_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_hook_proto_rawDescData)
	})
	return file_grpc_hook_proto_rawDescData
}

var file_grpc_hook_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_grpc_hook_proto_goTypes = []interface{}{
	(*ImportSetRequest)(nil),  // 0: grpc_hook.ImportSetRequest
	(*Row)(nil),               // 1: grpc_hook.Row
	(*Field)(nil),             // 2: grpc_hook.Field
	(*ImportSetResponse)(nil), // 3: grpc_hook.ImportSetResponse
	(*ScriptRequest)(nil),     // 4: grpc_hook.ScriptRequest
	(*ScriptResponse)(nil),    // 5: grpc_hook.ScriptResponse
	nil,                       // 6: grpc_hook.ScriptRequest.ParamsEntry
	nil,                       // 7: grpc_hook.ScriptResponse.AnswerEntry
}
var file_grpc_hook_proto_depIdxs = []int32{
	2, // 0: grpc_hook.ImportSetResponse.fields:type_name -> grpc_hook.Field
	1, // 1: grpc_hook.ImportSetResponse.rows:type_name -> grpc_hook.Row
	6, // 2: grpc_hook.ScriptRequest.params:type_name -> grpc_hook.ScriptRequest.ParamsEntry
	7, // 3: grpc_hook.ScriptResponse.answer:type_name -> grpc_hook.ScriptResponse.AnswerEntry
	4, // 4: grpc_hook.HookService.OnScriptCall:input_type -> grpc_hook.ScriptRequest
	0, // 5: grpc_hook.HookService.OnImportSet:input_type -> grpc_hook.ImportSetRequest
	5, // 6: grpc_hook.HookService.OnScriptCall:output_type -> grpc_hook.ScriptResponse
	3, // 7: grpc_hook.HookService.OnImportSet:output_type -> grpc_hook.ImportSetResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_grpc_hook_proto_init() }
func file_grpc_hook_proto_init() {
	if File_grpc_hook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_hook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportSetRequest); i {
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
		file_grpc_hook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Row); i {
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
		file_grpc_hook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Field); i {
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
		file_grpc_hook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportSetResponse); i {
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
		file_grpc_hook_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScriptRequest); i {
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
		file_grpc_hook_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScriptResponse); i {
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
			RawDescriptor: file_grpc_hook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_hook_proto_goTypes,
		DependencyIndexes: file_grpc_hook_proto_depIdxs,
		MessageInfos:      file_grpc_hook_proto_msgTypes,
	}.Build()
	File_grpc_hook_proto = out.File
	file_grpc_hook_proto_rawDesc = nil
	file_grpc_hook_proto_goTypes = nil
	file_grpc_hook_proto_depIdxs = nil
}
