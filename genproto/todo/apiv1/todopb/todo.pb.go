// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0
// source: qclaogui/todo/v1/todo.proto

package todopb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Tasks wo have to do
type ToDo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique integer identifier of the todo task
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Completed     bool                   `protobuf:"varint,4,opt,name=completed,proto3" json:"completed,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{0}
}

func (x *ToDo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ToDo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ToDo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ToDo) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

func (x *ToDo) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *ToDo) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type CreateTodoRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version
	Api           string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Item          *ToDo  `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTodoRequest) Reset() {
	*x = CreateTodoRequest{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoRequest) ProtoMessage() {}

func (x *CreateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoRequest.ProtoReflect.Descriptor instead.
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTodoRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *CreateTodoRequest) GetItem() *ToDo {
	if x != nil {
		return x.Item
	}
	return nil
}

type CreateTodoResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version
	Api           string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id            string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTodoResponse) Reset() {
	*x = CreateTodoResponse{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTodoResponse) ProtoMessage() {}

func (x *CreateTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTodoResponse.ProtoReflect.Descriptor instead.
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTodoResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *CreateTodoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTodoRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version
	Api           string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id            string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTodoRequest) Reset() {
	*x = GetTodoRequest{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoRequest) ProtoMessage() {}

func (x *GetTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoRequest.ProtoReflect.Descriptor instead.
func (*GetTodoRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{3}
}

func (x *GetTodoRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *GetTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetTodoResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version
	Api           string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Item          *ToDo  `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTodoResponse) Reset() {
	*x = GetTodoResponse{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTodoResponse) ProtoMessage() {}

func (x *GetTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTodoResponse.ProtoReflect.Descriptor instead.
func (*GetTodoResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{4}
}

func (x *GetTodoResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *GetTodoResponse) GetItem() *ToDo {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpdateTodoRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version
	Api           string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Item          *ToDo  `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTodoRequest) Reset() {
	*x = UpdateTodoRequest{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoRequest) ProtoMessage() {}

func (x *UpdateTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoRequest.ProtoReflect.Descriptor instead.
func (*UpdateTodoRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTodoRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *UpdateTodoRequest) GetItem() *ToDo {
	if x != nil {
		return x.Item
	}
	return nil
}

type UpdateTodoResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version
	Api           string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Updated       int64  `protobuf:"varint,2,opt,name=updated,proto3" json:"updated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTodoResponse) Reset() {
	*x = UpdateTodoResponse{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTodoResponse) ProtoMessage() {}

func (x *UpdateTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTodoResponse.ProtoReflect.Descriptor instead.
func (*UpdateTodoResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTodoResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *UpdateTodoResponse) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type DeleteTodoRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version
	Api           string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Id            string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTodoRequest) Reset() {
	*x = DeleteTodoRequest{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoRequest) ProtoMessage() {}

func (x *DeleteTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoRequest.ProtoReflect.Descriptor instead.
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTodoRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *DeleteTodoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteTodoResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version
	Api           string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Deleted       int64  `protobuf:"varint,2,opt,name=deleted,proto3" json:"deleted,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTodoResponse) Reset() {
	*x = DeleteTodoResponse{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoResponse) ProtoMessage() {}

func (x *DeleteTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoResponse.ProtoReflect.Descriptor instead.
func (*DeleteTodoResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteTodoResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *DeleteTodoResponse) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

type ListTodoRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version
	Api           string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Limit         int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Completed     bool   `protobuf:"varint,3,opt,name=completed,proto3" json:"completed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTodoRequest) Reset() {
	*x = ListTodoRequest{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodoRequest) ProtoMessage() {}

func (x *ListTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodoRequest.ProtoReflect.Descriptor instead.
func (*ListTodoRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{9}
}

func (x *ListTodoRequest) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *ListTodoRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListTodoRequest) GetCompleted() bool {
	if x != nil {
		return x.Completed
	}
	return false
}

type ListTodoResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// API version
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// List of all
	Items         []*ToDo `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTodoResponse) Reset() {
	*x = ListTodoResponse{}
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTodoResponse) ProtoMessage() {}

func (x *ListTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_todo_v1_todo_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTodoResponse.ProtoReflect.Descriptor instead.
func (*ListTodoResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_todo_v1_todo_proto_rawDescGZIP(), []int{10}
}

func (x *ListTodoResponse) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *ListTodoResponse) GetItems() []*ToDo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_qclaogui_todo_v1_todo_proto protoreflect.FileDescriptor

var file_qclaogui_todo_v1_todo_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x71,
	0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe6, 0x01, 0x0a, 0x04, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12,
	0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x51, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69,
	0x12, 0x2a, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x36, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x32, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x2a, 0x0a,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x71, 0x63,
	0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x6f, 0x44, 0x6f, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x51, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69,
	0x12, 0x2a, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x40, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x61, 0x70, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x35,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x40, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x57, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70,
	0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x22, 0x52, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x2c, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x61, 0x69, 0x70,
	0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_qclaogui_todo_v1_todo_proto_rawDescOnce sync.Once
	file_qclaogui_todo_v1_todo_proto_rawDescData []byte
)

func file_qclaogui_todo_v1_todo_proto_rawDescGZIP() []byte {
	file_qclaogui_todo_v1_todo_proto_rawDescOnce.Do(func() {
		file_qclaogui_todo_v1_todo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_todo_v1_todo_proto_rawDesc), len(file_qclaogui_todo_v1_todo_proto_rawDesc)))
	})
	return file_qclaogui_todo_v1_todo_proto_rawDescData
}

var (
	file_qclaogui_todo_v1_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
	file_qclaogui_todo_v1_todo_proto_goTypes  = []any{
		(*ToDo)(nil),                  // 0: qclaogui.todo.v1.ToDo
		(*CreateTodoRequest)(nil),     // 1: qclaogui.todo.v1.CreateTodoRequest
		(*CreateTodoResponse)(nil),    // 2: qclaogui.todo.v1.CreateTodoResponse
		(*GetTodoRequest)(nil),        // 3: qclaogui.todo.v1.GetTodoRequest
		(*GetTodoResponse)(nil),       // 4: qclaogui.todo.v1.GetTodoResponse
		(*UpdateTodoRequest)(nil),     // 5: qclaogui.todo.v1.UpdateTodoRequest
		(*UpdateTodoResponse)(nil),    // 6: qclaogui.todo.v1.UpdateTodoResponse
		(*DeleteTodoRequest)(nil),     // 7: qclaogui.todo.v1.DeleteTodoRequest
		(*DeleteTodoResponse)(nil),    // 8: qclaogui.todo.v1.DeleteTodoResponse
		(*ListTodoRequest)(nil),       // 9: qclaogui.todo.v1.ListTodoRequest
		(*ListTodoResponse)(nil),      // 10: qclaogui.todo.v1.ListTodoResponse
		(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
	}
)

var file_qclaogui_todo_v1_todo_proto_depIdxs = []int32{
	11, // 0: qclaogui.todo.v1.ToDo.create_time:type_name -> google.protobuf.Timestamp
	11, // 1: qclaogui.todo.v1.ToDo.update_time:type_name -> google.protobuf.Timestamp
	0,  // 2: qclaogui.todo.v1.CreateTodoRequest.item:type_name -> qclaogui.todo.v1.ToDo
	0,  // 3: qclaogui.todo.v1.GetTodoResponse.item:type_name -> qclaogui.todo.v1.ToDo
	0,  // 4: qclaogui.todo.v1.UpdateTodoRequest.item:type_name -> qclaogui.todo.v1.ToDo
	0,  // 5: qclaogui.todo.v1.ListTodoResponse.items:type_name -> qclaogui.todo.v1.ToDo
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_qclaogui_todo_v1_todo_proto_init() }
func file_qclaogui_todo_v1_todo_proto_init() {
	if File_qclaogui_todo_v1_todo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_todo_v1_todo_proto_rawDesc), len(file_qclaogui_todo_v1_todo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qclaogui_todo_v1_todo_proto_goTypes,
		DependencyIndexes: file_qclaogui_todo_v1_todo_proto_depIdxs,
		MessageInfos:      file_qclaogui_todo_v1_todo_proto_msgTypes,
	}.Build()
	File_qclaogui_todo_v1_todo_proto = out.File
	file_qclaogui_todo_v1_todo_proto_goTypes = nil
	file_qclaogui_todo_v1_todo_proto_depIdxs = nil
}
