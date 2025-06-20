// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: a2a/v1/message.proto

package a2apb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	Role_ROLE_UNSPECIFIED Role = 0
	// USER role refers to communication from the client to the server.
	Role_ROLE_USER Role = 1
	// AGENT role refers to communication from the server to the client.
	Role_ROLE_AGENT Role = 2
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "ROLE_USER",
		2: "ROLE_AGENT",
	}
	Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"ROLE_USER":        1,
		"ROLE_AGENT":       2,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_a2a_v1_message_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_a2a_v1_message_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_a2a_v1_message_proto_rawDescGZIP(), []int{0}
}

// Message is one unit of communication between client and server. It is
// associated with a context and optionally a task. Since the server is
// responsible for the context definition, it must always provide a context_id
// in its messages. The client can optionally provide the context_id if it
// knows the context to associate the message to. Similarly for task_id,
// except the server decides if a task is created and whether to include the
// task_id.
type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The message id of the message. This is required and created by the
	// message creator.
	MessageId string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	// The context id of the message. This is optional and if set, the message
	// will be associated with the given context.
	ContextId string `protobuf:"bytes,2,opt,name=context_id,json=contextId,proto3" json:"context_id,omitempty"`
	// The task id of the message. This is optional and if set, the message
	// will be associated with the given task.
	TaskId string `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// A role for the message.
	Role Role `protobuf:"varint,4,opt,name=role,proto3,enum=a2a.v1.Role" json:"role,omitempty"`
	// Content is the container of the message content.
	Content []*Part `protobuf:"bytes,5,rep,name=content,proto3" json:"content,omitempty"`
	// Any optional metadata to provide along with the message.
	Metadata *structpb.Struct `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The URIs of extensions that are present or contributed to this Message.
	Extensions    []string `protobuf:"bytes,7,rep,name=extensions,proto3" json:"extensions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_a2a_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_a2a_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_a2a_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

func (x *Message) GetContextId() string {
	if x != nil {
		return x.ContextId
	}
	return ""
}

func (x *Message) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *Message) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *Message) GetContent() []*Part {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Message) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Message) GetExtensions() []string {
	if x != nil {
		return x.Extensions
	}
	return nil
}

// Part represents a container for a section of communication content.
// Parts can be purely textual, some sort of file (image, video, etc) or
// a structured data blob (i.e. JSON).
type Part struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Part:
	//
	//	*Part_Text
	//	*Part_File
	//	*Part_Data
	Part          isPart_Part `protobuf_oneof:"part"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Part) Reset() {
	*x = Part{}
	mi := &file_a2a_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Part) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Part) ProtoMessage() {}

func (x *Part) ProtoReflect() protoreflect.Message {
	mi := &file_a2a_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Part.ProtoReflect.Descriptor instead.
func (*Part) Descriptor() ([]byte, []int) {
	return file_a2a_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *Part) GetPart() isPart_Part {
	if x != nil {
		return x.Part
	}
	return nil
}

func (x *Part) GetText() string {
	if x != nil {
		if x, ok := x.Part.(*Part_Text); ok {
			return x.Text
		}
	}
	return ""
}

func (x *Part) GetFile() *FilePart {
	if x != nil {
		if x, ok := x.Part.(*Part_File); ok {
			return x.File
		}
	}
	return nil
}

func (x *Part) GetData() *DataPart {
	if x != nil {
		if x, ok := x.Part.(*Part_Data); ok {
			return x.Data
		}
	}
	return nil
}

type isPart_Part interface {
	isPart_Part()
}

type Part_Text struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

type Part_File struct {
	File *FilePart `protobuf:"bytes,2,opt,name=file,proto3,oneof"`
}

type Part_Data struct {
	Data *DataPart `protobuf:"bytes,3,opt,name=data,proto3,oneof"`
}

func (*Part_Text) isPart_Part() {}

func (*Part_File) isPart_Part() {}

func (*Part_Data) isPart_Part() {}

// FilePart represents the different ways files can be provided. If files are
// small, directly feeding the bytes is supported via file_with_bytes. If the
// file is large, the agent should read the content as appropriate directly
// from the file_with_uri source.
type FilePart struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to File:
	//
	//	*FilePart_FileWithUri
	//	*FilePart_FileWithBytes
	File          isFilePart_File `protobuf_oneof:"file"`
	MimeType      string          `protobuf:"bytes,3,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FilePart) Reset() {
	*x = FilePart{}
	mi := &file_a2a_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FilePart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilePart) ProtoMessage() {}

func (x *FilePart) ProtoReflect() protoreflect.Message {
	mi := &file_a2a_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilePart.ProtoReflect.Descriptor instead.
func (*FilePart) Descriptor() ([]byte, []int) {
	return file_a2a_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *FilePart) GetFile() isFilePart_File {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *FilePart) GetFileWithUri() string {
	if x != nil {
		if x, ok := x.File.(*FilePart_FileWithUri); ok {
			return x.FileWithUri
		}
	}
	return ""
}

func (x *FilePart) GetFileWithBytes() []byte {
	if x != nil {
		if x, ok := x.File.(*FilePart_FileWithBytes); ok {
			return x.FileWithBytes
		}
	}
	return nil
}

func (x *FilePart) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

type isFilePart_File interface {
	isFilePart_File()
}

type FilePart_FileWithUri struct {
	FileWithUri string `protobuf:"bytes,1,opt,name=file_with_uri,json=fileWithUri,proto3,oneof"`
}

type FilePart_FileWithBytes struct {
	FileWithBytes []byte `protobuf:"bytes,2,opt,name=file_with_bytes,json=fileWithBytes,proto3,oneof"`
}

func (*FilePart_FileWithUri) isFilePart_File() {}

func (*FilePart_FileWithBytes) isFilePart_File() {}

// DataPart represents a structured blob. This is most commonly a JSON payload.
type DataPart struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Data          *structpb.Struct       `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataPart) Reset() {
	*x = DataPart{}
	mi := &file_a2a_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataPart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataPart) ProtoMessage() {}

func (x *DataPart) ProtoReflect() protoreflect.Message {
	mi := &file_a2a_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPart.ProtoReflect.Descriptor instead.
func (*DataPart) Descriptor() ([]byte, []int) {
	return file_a2a_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *DataPart) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_a2a_v1_message_proto protoreflect.FileDescriptor

const file_a2a_v1_message_proto_rawDesc = "" +
	"\n" +
	"\x14a2a/v1/message.proto\x12\x06a2a.v1\x1a\x1cgoogle/protobuf/struct.proto\"\xff\x01\n" +
	"\aMessage\x12\x1d\n" +
	"\n" +
	"message_id\x18\x01 \x01(\tR\tmessageId\x12\x1d\n" +
	"\n" +
	"context_id\x18\x02 \x01(\tR\tcontextId\x12\x17\n" +
	"\atask_id\x18\x03 \x01(\tR\x06taskId\x12 \n" +
	"\x04role\x18\x04 \x01(\x0e2\f.a2a.v1.RoleR\x04role\x12&\n" +
	"\acontent\x18\x05 \x03(\v2\f.a2a.v1.PartR\acontent\x123\n" +
	"\bmetadata\x18\x06 \x01(\v2\x17.google.protobuf.StructR\bmetadata\x12\x1e\n" +
	"\n" +
	"extensions\x18\a \x03(\tR\n" +
	"extensions\"t\n" +
	"\x04Part\x12\x14\n" +
	"\x04text\x18\x01 \x01(\tH\x00R\x04text\x12&\n" +
	"\x04file\x18\x02 \x01(\v2\x10.a2a.v1.FilePartH\x00R\x04file\x12&\n" +
	"\x04data\x18\x03 \x01(\v2\x10.a2a.v1.DataPartH\x00R\x04dataB\x06\n" +
	"\x04part\"\x7f\n" +
	"\bFilePart\x12$\n" +
	"\rfile_with_uri\x18\x01 \x01(\tH\x00R\vfileWithUri\x12(\n" +
	"\x0ffile_with_bytes\x18\x02 \x01(\fH\x00R\rfileWithBytes\x12\x1b\n" +
	"\tmime_type\x18\x03 \x01(\tR\bmimeTypeB\x06\n" +
	"\x04file\"7\n" +
	"\bDataPart\x12+\n" +
	"\x04data\x18\x01 \x01(\v2\x17.google.protobuf.StructR\x04data*;\n" +
	"\x04Role\x12\x14\n" +
	"\x10ROLE_UNSPECIFIED\x10\x00\x12\r\n" +
	"\tROLE_USER\x10\x01\x12\x0e\n" +
	"\n" +
	"ROLE_AGENT\x10\x02B3Z1github.com/qclaogui/gaip/genproto/a2a/apiv1/a2apbb\x06proto3"

var (
	file_a2a_v1_message_proto_rawDescOnce sync.Once
	file_a2a_v1_message_proto_rawDescData []byte
)

func file_a2a_v1_message_proto_rawDescGZIP() []byte {
	file_a2a_v1_message_proto_rawDescOnce.Do(func() {
		file_a2a_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_a2a_v1_message_proto_rawDesc), len(file_a2a_v1_message_proto_rawDesc)))
	})
	return file_a2a_v1_message_proto_rawDescData
}

var (
	file_a2a_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_a2a_v1_message_proto_msgTypes  = make([]protoimpl.MessageInfo, 4)
	file_a2a_v1_message_proto_goTypes   = []any{
		(Role)(0),               // 0: a2a.v1.Role
		(*Message)(nil),         // 1: a2a.v1.Message
		(*Part)(nil),            // 2: a2a.v1.Part
		(*FilePart)(nil),        // 3: a2a.v1.FilePart
		(*DataPart)(nil),        // 4: a2a.v1.DataPart
		(*structpb.Struct)(nil), // 5: google.protobuf.Struct
	}
)

var file_a2a_v1_message_proto_depIdxs = []int32{
	0, // 0: a2a.v1.Message.role:type_name -> a2a.v1.Role
	2, // 1: a2a.v1.Message.content:type_name -> a2a.v1.Part
	5, // 2: a2a.v1.Message.metadata:type_name -> google.protobuf.Struct
	3, // 3: a2a.v1.Part.file:type_name -> a2a.v1.FilePart
	4, // 4: a2a.v1.Part.data:type_name -> a2a.v1.DataPart
	5, // 5: a2a.v1.DataPart.data:type_name -> google.protobuf.Struct
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_a2a_v1_message_proto_init() }
func file_a2a_v1_message_proto_init() {
	if File_a2a_v1_message_proto != nil {
		return
	}
	file_a2a_v1_message_proto_msgTypes[1].OneofWrappers = []any{
		(*Part_Text)(nil),
		(*Part_File)(nil),
		(*Part_Data)(nil),
	}
	file_a2a_v1_message_proto_msgTypes[2].OneofWrappers = []any{
		(*FilePart_FileWithUri)(nil),
		(*FilePart_FileWithBytes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_a2a_v1_message_proto_rawDesc), len(file_a2a_v1_message_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_a2a_v1_message_proto_goTypes,
		DependencyIndexes: file_a2a_v1_message_proto_depIdxs,
		EnumInfos:         file_a2a_v1_message_proto_enumTypes,
		MessageInfos:      file_a2a_v1_message_proto_msgTypes,
	}.Build()
	File_a2a_v1_message_proto = out.File
	file_a2a_v1_message_proto_goTypes = nil
	file_a2a_v1_message_proto_depIdxs = nil
}
