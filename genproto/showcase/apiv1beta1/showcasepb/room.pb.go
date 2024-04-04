// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: qclaogui/showcase/v1beta1/room.proto

package showcasepb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A chat room.
type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the chat room.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The human readable name of the chat room.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The description of the chat room.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// The timestamp at which the user was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The latest timestamp at which the user was updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_room_proto_rawDescGZIP(), []int{0}
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Room) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Room) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Room) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// The request message for CreateRoom method.
type CreateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The room to create.
	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
}

func (x *CreateRoomRequest) Reset() {
	*x = CreateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRoomRequest) ProtoMessage() {}

func (x *CreateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRoomRequest.ProtoReflect.Descriptor instead.
func (*CreateRoomRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_room_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRoomRequest) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

// The request message for GetRoom method.
type GetRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the requested room.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetRoomRequest) Reset() {
	*x = GetRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRoomRequest) ProtoMessage() {}

func (x *GetRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRoomRequest.ProtoReflect.Descriptor instead.
func (*GetRoomRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_room_proto_rawDescGZIP(), []int{2}
}

func (x *GetRoomRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The room to update.
	Room *Room `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	// The field mask to determine which fields are to be updated. If empty, the
	// server will assume all fields are to be updated.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
}

func (x *UpdateRoomRequest) Reset() {
	*x = UpdateRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRoomRequest) ProtoMessage() {}

func (x *UpdateRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRoomRequest.ProtoReflect.Descriptor instead.
func (*UpdateRoomRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_room_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRoomRequest) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *UpdateRoomRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the requested room.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteRoomRequest) Reset() {
	*x = DeleteRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoomRequest) ProtoMessage() {}

func (x *DeleteRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoomRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoomRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_room_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRoomRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListRoomsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of rooms return. Server may return fewer rooms
	// than requested. If unspecified, server will pick an appropriate default.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// The value of google.showcase.v1beta1.ListRoomsResponse.next_page_token
	// returned from the previous call to
	// `qclaogui.showcase.v1beta1.Messaging\ListRooms` method.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListRoomsRequest) Reset() {
	*x = ListRoomsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoomsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomsRequest) ProtoMessage() {}

func (x *ListRoomsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomsRequest.ProtoReflect.Descriptor instead.
func (*ListRoomsRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_room_proto_rawDescGZIP(), []int{5}
}

func (x *ListRoomsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListRoomsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListRoomsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of rooms.
	Rooms []*Room `protobuf:"bytes,1,rep,name=rooms,proto3" json:"rooms,omitempty"`
	// A token to retrieve next page of results.
	// Pass this value in ListRoomsRequest.page_token field in the subsequent
	// call to `qclaogui.showcase.v1beta1.Messaging\ListRooms` method to retrieve
	// the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListRoomsResponse) Reset() {
	*x = ListRoomsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRoomsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRoomsResponse) ProtoMessage() {}

func (x *ListRoomsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRoomsResponse.ProtoReflect.Descriptor instead.
func (*ListRoomsResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_room_proto_rawDescGZIP(), []int{6}
}

func (x *ListRoomsResponse) GetRooms() []*Room {
	if x != nil {
		return x.Rooms
	}
	return nil
}

func (x *ListRoomsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_qclaogui_showcase_v1beta1_room_proto protoreflect.FileDescriptor

var file_qclaogui_showcase_v1beta1_room_proto_rawDesc = []byte{
	0x0a, 0x24, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x77, 0x63,
	0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x72, 0x6f, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69,
	0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x97, 0x02, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x2d, 0xea, 0x41, 0x2a,
	0x0a, 0x1a, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f,
	0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0c, 0x72, 0x6f,
	0x6f, 0x6d, 0x73, 0x2f, 0x7b, 0x72, 0x6f, 0x6f, 0x6d, 0x7d, 0x22, 0x48, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x33, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x22, 0x48, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1c, 0x0a, 0x1a, 0x73, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x85,
	0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x73, 0x68,
	0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x22, 0x4b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41, 0x02, 0xfa, 0x41,
	0x1c, 0x0a, 0x1a, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x4e, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x72, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x72, 0x6f, 0x6f, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67,
	0x75, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x05, 0x72, 0x6f, 0x6f, 0x6d, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67,
	0x61, 0x69, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x6f,
	0x77, 0x63, 0x61, 0x73, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x73, 0x68, 0x6f, 0x77, 0x63, 0x61, 0x73, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_qclaogui_showcase_v1beta1_room_proto_rawDescOnce sync.Once
	file_qclaogui_showcase_v1beta1_room_proto_rawDescData = file_qclaogui_showcase_v1beta1_room_proto_rawDesc
)

func file_qclaogui_showcase_v1beta1_room_proto_rawDescGZIP() []byte {
	file_qclaogui_showcase_v1beta1_room_proto_rawDescOnce.Do(func() {
		file_qclaogui_showcase_v1beta1_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_qclaogui_showcase_v1beta1_room_proto_rawDescData)
	})
	return file_qclaogui_showcase_v1beta1_room_proto_rawDescData
}

var file_qclaogui_showcase_v1beta1_room_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_qclaogui_showcase_v1beta1_room_proto_goTypes = []interface{}{
	(*Room)(nil),                  // 0: qclaogui.showcase.v1beta1.Room
	(*CreateRoomRequest)(nil),     // 1: qclaogui.showcase.v1beta1.CreateRoomRequest
	(*GetRoomRequest)(nil),        // 2: qclaogui.showcase.v1beta1.GetRoomRequest
	(*UpdateRoomRequest)(nil),     // 3: qclaogui.showcase.v1beta1.UpdateRoomRequest
	(*DeleteRoomRequest)(nil),     // 4: qclaogui.showcase.v1beta1.DeleteRoomRequest
	(*ListRoomsRequest)(nil),      // 5: qclaogui.showcase.v1beta1.ListRoomsRequest
	(*ListRoomsResponse)(nil),     // 6: qclaogui.showcase.v1beta1.ListRoomsResponse
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil), // 8: google.protobuf.FieldMask
}
var file_qclaogui_showcase_v1beta1_room_proto_depIdxs = []int32{
	7, // 0: qclaogui.showcase.v1beta1.Room.create_time:type_name -> google.protobuf.Timestamp
	7, // 1: qclaogui.showcase.v1beta1.Room.update_time:type_name -> google.protobuf.Timestamp
	0, // 2: qclaogui.showcase.v1beta1.CreateRoomRequest.room:type_name -> qclaogui.showcase.v1beta1.Room
	0, // 3: qclaogui.showcase.v1beta1.UpdateRoomRequest.room:type_name -> qclaogui.showcase.v1beta1.Room
	8, // 4: qclaogui.showcase.v1beta1.UpdateRoomRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 5: qclaogui.showcase.v1beta1.ListRoomsResponse.rooms:type_name -> qclaogui.showcase.v1beta1.Room
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_qclaogui_showcase_v1beta1_room_proto_init() }
func file_qclaogui_showcase_v1beta1_room_proto_init() {
	if File_qclaogui_showcase_v1beta1_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qclaogui_showcase_v1beta1_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_qclaogui_showcase_v1beta1_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRoomRequest); i {
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
		file_qclaogui_showcase_v1beta1_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRoomRequest); i {
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
		file_qclaogui_showcase_v1beta1_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRoomRequest); i {
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
		file_qclaogui_showcase_v1beta1_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRoomRequest); i {
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
		file_qclaogui_showcase_v1beta1_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoomsRequest); i {
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
		file_qclaogui_showcase_v1beta1_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRoomsResponse); i {
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
			RawDescriptor: file_qclaogui_showcase_v1beta1_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qclaogui_showcase_v1beta1_room_proto_goTypes,
		DependencyIndexes: file_qclaogui_showcase_v1beta1_room_proto_depIdxs,
		MessageInfos:      file_qclaogui_showcase_v1beta1_room_proto_msgTypes,
	}.Build()
	File_qclaogui_showcase_v1beta1_room_proto = out.File
	file_qclaogui_showcase_v1beta1_room_proto_rawDesc = nil
	file_qclaogui_showcase_v1beta1_room_proto_goTypes = nil
	file_qclaogui_showcase_v1beta1_room_proto_depIdxs = nil
}
