// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: qclaogui/showcase/v1beta1/blurb.proto

package showcasepb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
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

// The action that triggered the blurb to be returned.
type StreamBlurbsResponse_Action int32

const (
	StreamBlurbsResponse_ACTION_UNSPECIFIED StreamBlurbsResponse_Action = 0
	// Specifies that the blurb was created.
	StreamBlurbsResponse_ACTION_CREATE StreamBlurbsResponse_Action = 1
	// Specifies that the blurb was updated.
	StreamBlurbsResponse_ACTION_UPDATE StreamBlurbsResponse_Action = 2
	// Specifies that the blurb was deleted.
	StreamBlurbsResponse_ACTION_DELETE StreamBlurbsResponse_Action = 3
)

// Enum value maps for StreamBlurbsResponse_Action.
var (
	StreamBlurbsResponse_Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "ACTION_CREATE",
		2: "ACTION_UPDATE",
		3: "ACTION_DELETE",
	}
	StreamBlurbsResponse_Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"ACTION_CREATE":      1,
		"ACTION_UPDATE":      2,
		"ACTION_DELETE":      3,
	}
)

func (x StreamBlurbsResponse_Action) Enum() *StreamBlurbsResponse_Action {
	p := new(StreamBlurbsResponse_Action)
	*p = x
	return p
}

func (x StreamBlurbsResponse_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StreamBlurbsResponse_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_qclaogui_showcase_v1beta1_blurb_proto_enumTypes[0].Descriptor()
}

func (StreamBlurbsResponse_Action) Type() protoreflect.EnumType {
	return &file_qclaogui_showcase_v1beta1_blurb_proto_enumTypes[0]
}

func (x StreamBlurbsResponse_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StreamBlurbsResponse_Action.Descriptor instead.
func (StreamBlurbsResponse_Action) EnumDescriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{11, 0}
}

// This protocol buffer message represents a blurb sent to a chat room or
// posted on a user profile.
type Blurb struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the chat room.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource name of the blurb's author.
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	// Types that are valid to be assigned to Content:
	//
	//	*Blurb_Text
	//	*Blurb_Image
	Content isBlurb_Content `protobuf_oneof:"content"`
	// The timestamp at which the user was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// The latest timestamp at which the user was updated.
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Blurb) Reset() {
	*x = Blurb{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Blurb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blurb) ProtoMessage() {}

func (x *Blurb) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blurb.ProtoReflect.Descriptor instead.
func (*Blurb) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{0}
}

func (x *Blurb) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Blurb) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Blurb) GetContent() isBlurb_Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Blurb) GetText() string {
	if x != nil {
		if x, ok := x.Content.(*Blurb_Text); ok {
			return x.Text
		}
	}
	return ""
}

func (x *Blurb) GetImage() []byte {
	if x != nil {
		if x, ok := x.Content.(*Blurb_Image); ok {
			return x.Image
		}
	}
	return nil
}

func (x *Blurb) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Blurb) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type isBlurb_Content interface {
	isBlurb_Content()
}

type Blurb_Text struct {
	// The textual content of this blurb.
	Text string `protobuf:"bytes,3,opt,name=text,proto3,oneof"`
}

type Blurb_Image struct {
	// The image content of this blurb.
	Image []byte `protobuf:"bytes,4,opt,name=image,proto3,oneof"`
}

func (*Blurb_Text) isBlurb_Content() {}

func (*Blurb_Image) isBlurb_Content() {}

type CreateBlurbRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the chat room or user profile that this blurb will
	// be tied to.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The blurb to create.
	Blurb         *Blurb `protobuf:"bytes,2,opt,name=blurb,proto3" json:"blurb,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBlurbRequest) Reset() {
	*x = CreateBlurbRequest{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBlurbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlurbRequest) ProtoMessage() {}

func (x *CreateBlurbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlurbRequest.ProtoReflect.Descriptor instead.
func (*CreateBlurbRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBlurbRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreateBlurbRequest) GetBlurb() *Blurb {
	if x != nil {
		return x.Blurb
	}
	return nil
}

type GetBlurbRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the requested blurb.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBlurbRequest) Reset() {
	*x = GetBlurbRequest{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBlurbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlurbRequest) ProtoMessage() {}

func (x *GetBlurbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlurbRequest.ProtoReflect.Descriptor instead.
func (*GetBlurbRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlurbRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateBlurbRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The blurb to update.
	Blurb *Blurb `protobuf:"bytes,1,opt,name=blurb,proto3" json:"blurb,omitempty"`
	// The field mask to determine which fields are to be updated. If empty, the
	// server will assume all fields are to be updated.
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateBlurbRequest) Reset() {
	*x = UpdateBlurbRequest{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateBlurbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlurbRequest) ProtoMessage() {}

func (x *UpdateBlurbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlurbRequest.ProtoReflect.Descriptor instead.
func (*UpdateBlurbRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateBlurbRequest) GetBlurb() *Blurb {
	if x != nil {
		return x.Blurb
	}
	return nil
}

func (x *UpdateBlurbRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteBlurbRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the requested blurb.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteBlurbRequest) Reset() {
	*x = DeleteBlurbRequest{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteBlurbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlurbRequest) ProtoMessage() {}

func (x *DeleteBlurbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlurbRequest.ProtoReflect.Descriptor instead.
func (*DeleteBlurbRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteBlurbRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListBlurbsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of the requested room or profile who blurbs to list.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of blurbs to return. Server may return fewer
	// blurbs than requested. If unspecified, server will pick an appropriate
	// default.
	PageSize      int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBlurbsRequest) Reset() {
	*x = ListBlurbsRequest{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlurbsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlurbsRequest) ProtoMessage() {}

func (x *ListBlurbsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlurbsRequest.ProtoReflect.Descriptor instead.
func (*ListBlurbsRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{5}
}

func (x *ListBlurbsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListBlurbsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListBlurbsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListBlurbsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The list of blurbs.
	Blurbs        []*Blurb `protobuf:"bytes,1,rep,name=blurbs,proto3" json:"blurbs,omitempty"`
	NextPageToken string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListBlurbsResponse) Reset() {
	*x = ListBlurbsResponse{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListBlurbsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlurbsResponse) ProtoMessage() {}

func (x *ListBlurbsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlurbsResponse.ProtoReflect.Descriptor instead.
func (*ListBlurbsResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{6}
}

func (x *ListBlurbsResponse) GetBlurbs() []*Blurb {
	if x != nil {
		return x.Blurbs
	}
	return nil
}

func (x *ListBlurbsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type SearchBlurbsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The query used to search for blurbs containing to words of this string.
	// Only posts that contain an exact match of a queried word will be returned.
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	// The rooms or profiles to search. If unset, `SearchBlurbs` will search all
	// rooms and all profiles.
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// The maximum number of blurbs return. Server may return fewer
	// blurbs than requested. If unspecified, server will pick an appropriate
	// default.
	PageSize      int32  `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchBlurbsRequest) Reset() {
	*x = SearchBlurbsRequest{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchBlurbsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBlurbsRequest) ProtoMessage() {}

func (x *SearchBlurbsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBlurbsRequest.ProtoReflect.Descriptor instead.
func (*SearchBlurbsRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{7}
}

func (x *SearchBlurbsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchBlurbsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *SearchBlurbsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchBlurbsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type SearchBlurbsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Blurbs that matched the search query.
	Blurbs []*Blurb `protobuf:"bytes,1,rep,name=blurbs,proto3" json:"blurbs,omitempty"`
	// A token to retrieve next page of results.
	// Pass this value in SearchBlurbsRequest.page_token field in the subsequent
	// call to `google.showcase.v1beta1.Blurb\SearchBlurbs` method to
	// retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchBlurbsResponse) Reset() {
	*x = SearchBlurbsResponse{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchBlurbsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBlurbsResponse) ProtoMessage() {}

func (x *SearchBlurbsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBlurbsResponse.ProtoReflect.Descriptor instead.
func (*SearchBlurbsResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{8}
}

func (x *SearchBlurbsResponse) GetBlurbs() []*Blurb {
	if x != nil {
		return x.Blurbs
	}
	return nil
}

func (x *SearchBlurbsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type SearchBlurbsMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// This signals to the client when to next poll for response.
	RetryInfo     *errdetails.RetryInfo `protobuf:"bytes,1,opt,name=retry_info,json=retryInfo,proto3" json:"retry_info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchBlurbsMetadata) Reset() {
	*x = SearchBlurbsMetadata{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchBlurbsMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchBlurbsMetadata) ProtoMessage() {}

func (x *SearchBlurbsMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchBlurbsMetadata.ProtoReflect.Descriptor instead.
func (*SearchBlurbsMetadata) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{9}
}

func (x *SearchBlurbsMetadata) GetRetryInfo() *errdetails.RetryInfo {
	if x != nil {
		return x.RetryInfo
	}
	return nil
}

type StreamBlurbsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The resource name of a chat room or user profile whose blurbs to stream.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The time at which this stream will close.
	ExpireTime    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expire_time,json=expireTime,proto3" json:"expire_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamBlurbsRequest) Reset() {
	*x = StreamBlurbsRequest{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamBlurbsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamBlurbsRequest) ProtoMessage() {}

func (x *StreamBlurbsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamBlurbsRequest.ProtoReflect.Descriptor instead.
func (*StreamBlurbsRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{10}
}

func (x *StreamBlurbsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StreamBlurbsRequest) GetExpireTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpireTime
	}
	return nil
}

type StreamBlurbsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The blurb that was either created, updated, or deleted.
	Blurb *Blurb `protobuf:"bytes,1,opt,name=blurb,proto3" json:"blurb,omitempty"`
	// The action that triggered the blurb to be returned.
	Action        StreamBlurbsResponse_Action `protobuf:"varint,2,opt,name=action,proto3,enum=qclaogui.showcase.v1beta1.StreamBlurbsResponse_Action" json:"action,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StreamBlurbsResponse) Reset() {
	*x = StreamBlurbsResponse{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StreamBlurbsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamBlurbsResponse) ProtoMessage() {}

func (x *StreamBlurbsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamBlurbsResponse.ProtoReflect.Descriptor instead.
func (*StreamBlurbsResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{11}
}

func (x *StreamBlurbsResponse) GetBlurb() *Blurb {
	if x != nil {
		return x.Blurb
	}
	return nil
}

func (x *StreamBlurbsResponse) GetAction() StreamBlurbsResponse_Action {
	if x != nil {
		return x.Action
	}
	return StreamBlurbsResponse_ACTION_UNSPECIFIED
}

type SendBlurbsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The names of successful blurb creations.
	Names         []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendBlurbsResponse) Reset() {
	*x = SendBlurbsResponse{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendBlurbsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBlurbsResponse) ProtoMessage() {}

func (x *SendBlurbsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBlurbsResponse.ProtoReflect.Descriptor instead.
func (*SendBlurbsResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{12}
}

func (x *SendBlurbsResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type ConnectRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Request:
	//
	//	*ConnectRequest_Config
	//	*ConnectRequest_Blurb
	Request       isConnectRequest_Request `protobuf_oneof:"request"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{13}
}

func (x *ConnectRequest) GetRequest() isConnectRequest_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ConnectRequest) GetConfig() *ConnectRequest_ConnectConfig {
	if x != nil {
		if x, ok := x.Request.(*ConnectRequest_Config); ok {
			return x.Config
		}
	}
	return nil
}

func (x *ConnectRequest) GetBlurb() *Blurb {
	if x != nil {
		if x, ok := x.Request.(*ConnectRequest_Blurb); ok {
			return x.Blurb
		}
	}
	return nil
}

type isConnectRequest_Request interface {
	isConnectRequest_Request()
}

type ConnectRequest_Config struct {
	// Provides information that specifies how to process subsequent requests.
	// The first `ConnectRequest` message must contain a `config`  message.
	Config *ConnectRequest_ConnectConfig `protobuf:"bytes,1,opt,name=config,proto3,oneof"`
}

type ConnectRequest_Blurb struct {
	// The blurb to be created.
	Blurb *Blurb `protobuf:"bytes,2,opt,name=blurb,proto3,oneof"`
}

func (*ConnectRequest_Config) isConnectRequest_Request() {}

func (*ConnectRequest_Blurb) isConnectRequest_Request() {}

type ConnectRequest_ConnectConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The room or profile to follow and create messages for.
	Parent        string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectRequest_ConnectConfig) Reset() {
	*x = ConnectRequest_ConnectConfig{}
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectRequest_ConnectConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest_ConnectConfig) ProtoMessage() {}

func (x *ConnectRequest_ConnectConfig) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest_ConnectConfig.ProtoReflect.Descriptor instead.
func (*ConnectRequest_ConnectConfig) Descriptor() ([]byte, []int) {
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP(), []int{13, 0}
}

func (x *ConnectRequest_ConnectConfig) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

var File_qclaogui_showcase_v1beta1_blurb_proto protoreflect.FileDescriptor

const file_qclaogui_showcase_v1beta1_blurb_proto_rawDesc = "" +
	"\n" +
	"%qclaogui/showcase/v1beta1/blurb.proto\x12\x19qclaogui.showcase.v1beta1\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a google/protobuf/field_mask.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/rpc/error_details.proto\"\xf7\x02\n" +
	"\x05Blurb\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x127\n" +
	"\x04user\x18\x02 \x01(\tB#\xe2A\x01\x02\xfaA\x1c\n" +
	"\x1ashowcase.qclaogui.com/UserR\x04user\x12\x14\n" +
	"\x04text\x18\x03 \x01(\tH\x00R\x04text\x12\x16\n" +
	"\x05image\x18\x04 \x01(\fH\x00R\x05image\x12A\n" +
	"\vcreate_time\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampB\x04\xe2A\x01\x03R\n" +
	"createTime\x12A\n" +
	"\vupdate_time\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampB\x04\xe2A\x01\x03R\n" +
	"updateTime:b\xeaA_\n" +
	"\x1bshowcase.qclaogui.com/Blurb\x12#users/{user}/profile/blurbs/{blurb}\x12\x1brooms/{room}/blurbs/{blurb}B\t\n" +
	"\acontent\"\x8a\x01\n" +
	"\x12CreateBlurbRequest\x12<\n" +
	"\x06parent\x18\x01 \x01(\tB$\xe2A\x01\x02\xfaA\x1d\x12\x1bshowcase.qclaogui.com/BlurbR\x06parent\x126\n" +
	"\x05blurb\x18\x02 \x01(\v2 .qclaogui.showcase.v1beta1.BlurbR\x05blurb\"K\n" +
	"\x0fGetBlurbRequest\x128\n" +
	"\x04name\x18\x01 \x01(\tB$\xe2A\x01\x02\xfaA\x1d\n" +
	"\x1bshowcase.qclaogui.com/BlurbR\x04name\"\x89\x01\n" +
	"\x12UpdateBlurbRequest\x126\n" +
	"\x05blurb\x18\x01 \x01(\v2 .qclaogui.showcase.v1beta1.BlurbR\x05blurb\x12;\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\"N\n" +
	"\x12DeleteBlurbRequest\x128\n" +
	"\x04name\x18\x01 \x01(\tB$\xe2A\x01\x02\xfaA\x1d\n" +
	"\x1bshowcase.qclaogui.com/BlurbR\x04name\"\x8d\x01\n" +
	"\x11ListBlurbsRequest\x12<\n" +
	"\x06parent\x18\x01 \x01(\tB$\xe2A\x01\x02\xfaA\x1d\x12\x1bshowcase.qclaogui.com/BlurbR\x06parent\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tR\tpageToken\"v\n" +
	"\x12ListBlurbsResponse\x128\n" +
	"\x06blurbs\x18\x01 \x03(\v2 .qclaogui.showcase.v1beta1.BlurbR\x06blurbs\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\xa7\x01\n" +
	"\x13SearchBlurbsRequest\x12\x1a\n" +
	"\x05query\x18\x01 \x01(\tB\x04\xe2A\x01\x02R\x05query\x128\n" +
	"\x06parent\x18\x02 \x01(\tB \xfaA\x1d\x12\x1bshowcase.qclaogui.com/BlurbR\x06parent\x12\x1b\n" +
	"\tpage_size\x18\x03 \x01(\x05R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x04 \x01(\tR\tpageToken\"x\n" +
	"\x14SearchBlurbsResponse\x128\n" +
	"\x06blurbs\x18\x01 \x03(\v2 .qclaogui.showcase.v1beta1.BlurbR\x06blurbs\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"L\n" +
	"\x14SearchBlurbsMetadata\x124\n" +
	"\n" +
	"retry_info\x18\x01 \x01(\v2\x15.google.rpc.RetryInfoR\tretryInfo\"\x92\x01\n" +
	"\x13StreamBlurbsRequest\x128\n" +
	"\x04name\x18\x01 \x01(\tB$\xe2A\x01\x02\xfaA\x1d\x12\x1bshowcase.qclaogui.com/BlurbR\x04name\x12A\n" +
	"\vexpire_time\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampB\x04\xe2A\x01\x02R\n" +
	"expireTime\"\xf9\x01\n" +
	"\x14StreamBlurbsResponse\x126\n" +
	"\x05blurb\x18\x01 \x01(\v2 .qclaogui.showcase.v1beta1.BlurbR\x05blurb\x12N\n" +
	"\x06action\x18\x02 \x01(\x0e26.qclaogui.showcase.v1beta1.StreamBlurbsResponse.ActionR\x06action\"Y\n" +
	"\x06Action\x12\x16\n" +
	"\x12ACTION_UNSPECIFIED\x10\x00\x12\x11\n" +
	"\rACTION_CREATE\x10\x01\x12\x11\n" +
	"\rACTION_UPDATE\x10\x02\x12\x11\n" +
	"\rACTION_DELETE\x10\x03\"*\n" +
	"\x12SendBlurbsResponse\x12\x14\n" +
	"\x05names\x18\x01 \x03(\tR\x05names\"\xf3\x01\n" +
	"\x0eConnectRequest\x12Q\n" +
	"\x06config\x18\x01 \x01(\v27.qclaogui.showcase.v1beta1.ConnectRequest.ConnectConfigH\x00R\x06config\x128\n" +
	"\x05blurb\x18\x02 \x01(\v2 .qclaogui.showcase.v1beta1.BlurbH\x00R\x05blurb\x1aI\n" +
	"\rConnectConfig\x128\n" +
	"\x06parent\x18\x01 \x01(\tB \xfaA\x1d\x12\x1bshowcase.qclaogui.com/BlurbR\x06parentB\t\n" +
	"\arequestBBZ@github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepbb\x06proto3"

var (
	file_qclaogui_showcase_v1beta1_blurb_proto_rawDescOnce sync.Once
	file_qclaogui_showcase_v1beta1_blurb_proto_rawDescData []byte
)

func file_qclaogui_showcase_v1beta1_blurb_proto_rawDescGZIP() []byte {
	file_qclaogui_showcase_v1beta1_blurb_proto_rawDescOnce.Do(func() {
		file_qclaogui_showcase_v1beta1_blurb_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_showcase_v1beta1_blurb_proto_rawDesc), len(file_qclaogui_showcase_v1beta1_blurb_proto_rawDesc)))
	})
	return file_qclaogui_showcase_v1beta1_blurb_proto_rawDescData
}

var (
	file_qclaogui_showcase_v1beta1_blurb_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes  = make([]protoimpl.MessageInfo, 15)
	file_qclaogui_showcase_v1beta1_blurb_proto_goTypes   = []any{
		(StreamBlurbsResponse_Action)(0),     // 0: qclaogui.showcase.v1beta1.StreamBlurbsResponse.Action
		(*Blurb)(nil),                        // 1: qclaogui.showcase.v1beta1.Blurb
		(*CreateBlurbRequest)(nil),           // 2: qclaogui.showcase.v1beta1.CreateBlurbRequest
		(*GetBlurbRequest)(nil),              // 3: qclaogui.showcase.v1beta1.GetBlurbRequest
		(*UpdateBlurbRequest)(nil),           // 4: qclaogui.showcase.v1beta1.UpdateBlurbRequest
		(*DeleteBlurbRequest)(nil),           // 5: qclaogui.showcase.v1beta1.DeleteBlurbRequest
		(*ListBlurbsRequest)(nil),            // 6: qclaogui.showcase.v1beta1.ListBlurbsRequest
		(*ListBlurbsResponse)(nil),           // 7: qclaogui.showcase.v1beta1.ListBlurbsResponse
		(*SearchBlurbsRequest)(nil),          // 8: qclaogui.showcase.v1beta1.SearchBlurbsRequest
		(*SearchBlurbsResponse)(nil),         // 9: qclaogui.showcase.v1beta1.SearchBlurbsResponse
		(*SearchBlurbsMetadata)(nil),         // 10: qclaogui.showcase.v1beta1.SearchBlurbsMetadata
		(*StreamBlurbsRequest)(nil),          // 11: qclaogui.showcase.v1beta1.StreamBlurbsRequest
		(*StreamBlurbsResponse)(nil),         // 12: qclaogui.showcase.v1beta1.StreamBlurbsResponse
		(*SendBlurbsResponse)(nil),           // 13: qclaogui.showcase.v1beta1.SendBlurbsResponse
		(*ConnectRequest)(nil),               // 14: qclaogui.showcase.v1beta1.ConnectRequest
		(*ConnectRequest_ConnectConfig)(nil), // 15: qclaogui.showcase.v1beta1.ConnectRequest.ConnectConfig
		(*timestamppb.Timestamp)(nil),        // 16: google.protobuf.Timestamp
		(*fieldmaskpb.FieldMask)(nil),        // 17: google.protobuf.FieldMask
		(*errdetails.RetryInfo)(nil),         // 18: google.rpc.RetryInfo
	}
)

var file_qclaogui_showcase_v1beta1_blurb_proto_depIdxs = []int32{
	16, // 0: qclaogui.showcase.v1beta1.Blurb.create_time:type_name -> google.protobuf.Timestamp
	16, // 1: qclaogui.showcase.v1beta1.Blurb.update_time:type_name -> google.protobuf.Timestamp
	1,  // 2: qclaogui.showcase.v1beta1.CreateBlurbRequest.blurb:type_name -> qclaogui.showcase.v1beta1.Blurb
	1,  // 3: qclaogui.showcase.v1beta1.UpdateBlurbRequest.blurb:type_name -> qclaogui.showcase.v1beta1.Blurb
	17, // 4: qclaogui.showcase.v1beta1.UpdateBlurbRequest.update_mask:type_name -> google.protobuf.FieldMask
	1,  // 5: qclaogui.showcase.v1beta1.ListBlurbsResponse.blurbs:type_name -> qclaogui.showcase.v1beta1.Blurb
	1,  // 6: qclaogui.showcase.v1beta1.SearchBlurbsResponse.blurbs:type_name -> qclaogui.showcase.v1beta1.Blurb
	18, // 7: qclaogui.showcase.v1beta1.SearchBlurbsMetadata.retry_info:type_name -> google.rpc.RetryInfo
	16, // 8: qclaogui.showcase.v1beta1.StreamBlurbsRequest.expire_time:type_name -> google.protobuf.Timestamp
	1,  // 9: qclaogui.showcase.v1beta1.StreamBlurbsResponse.blurb:type_name -> qclaogui.showcase.v1beta1.Blurb
	0,  // 10: qclaogui.showcase.v1beta1.StreamBlurbsResponse.action:type_name -> qclaogui.showcase.v1beta1.StreamBlurbsResponse.Action
	15, // 11: qclaogui.showcase.v1beta1.ConnectRequest.config:type_name -> qclaogui.showcase.v1beta1.ConnectRequest.ConnectConfig
	1,  // 12: qclaogui.showcase.v1beta1.ConnectRequest.blurb:type_name -> qclaogui.showcase.v1beta1.Blurb
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_qclaogui_showcase_v1beta1_blurb_proto_init() }
func file_qclaogui_showcase_v1beta1_blurb_proto_init() {
	if File_qclaogui_showcase_v1beta1_blurb_proto != nil {
		return
	}
	file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[0].OneofWrappers = []any{
		(*Blurb_Text)(nil),
		(*Blurb_Image)(nil),
	}
	file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes[13].OneofWrappers = []any{
		(*ConnectRequest_Config)(nil),
		(*ConnectRequest_Blurb)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_showcase_v1beta1_blurb_proto_rawDesc), len(file_qclaogui_showcase_v1beta1_blurb_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qclaogui_showcase_v1beta1_blurb_proto_goTypes,
		DependencyIndexes: file_qclaogui_showcase_v1beta1_blurb_proto_depIdxs,
		EnumInfos:         file_qclaogui_showcase_v1beta1_blurb_proto_enumTypes,
		MessageInfos:      file_qclaogui_showcase_v1beta1_blurb_proto_msgTypes,
	}.Build()
	File_qclaogui_showcase_v1beta1_blurb_proto = out.File
	file_qclaogui_showcase_v1beta1_blurb_proto_goTypes = nil
	file_qclaogui_showcase_v1beta1_blurb_proto_depIdxs = nil
}
