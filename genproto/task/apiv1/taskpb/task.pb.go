// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: qclaogui/task/v1/task.proto

package taskpb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Deadline      *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`
	CompletedAt   *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=completed_at,json=completedAt,proto3" json:"completed_at,omitempty"`
	CreateTime    *timestamppb.Timestamp `protobuf:"bytes,1000,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateTime    *timestamppb.Timestamp `protobuf:"bytes,1001,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_qclaogui_task_v1_task_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetDeadline() *timestamppb.Timestamp {
	if x != nil {
		return x.Deadline
	}
	return nil
}

func (x *Task) GetCompletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CompletedAt
	}
	return nil
}

func (x *Task) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Task) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

type CreateTaskRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Task is the the task to create.
	Task          *Task `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTaskRequest) Reset() {
	*x = CreateTaskRequest{}
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRequest) ProtoMessage() {}

func (x *CreateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_task_v1_task_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type UpdateTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateTaskRequest) Reset() {
	*x = UpdateTaskRequest{}
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRequest) ProtoMessage() {}

func (x *UpdateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_task_v1_task_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateTaskRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *UpdateTaskRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTaskRequest) Reset() {
	*x = DeleteTaskRequest{}
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRequest) ProtoMessage() {}

func (x *DeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_task_v1_task_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteTaskRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UndeleteTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UndeleteTaskRequest) Reset() {
	*x = UndeleteTaskRequest{}
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UndeleteTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UndeleteTaskRequest) ProtoMessage() {}

func (x *UndeleteTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UndeleteTaskRequest.ProtoReflect.Descriptor instead.
func (*UndeleteTaskRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_task_v1_task_proto_rawDescGZIP(), []int{4}
}

func (x *UndeleteTaskRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetTaskRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTaskRequest) Reset() {
	*x = GetTaskRequest{}
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRequest) ProtoMessage() {}

func (x *GetTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_task_v1_task_proto_rawDescGZIP(), []int{5}
}

func (x *GetTaskRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListTasksRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PageSize      int32                  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken     string                 `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksRequest) Reset() {
	*x = ListTasksRequest{}
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksRequest) ProtoMessage() {}

func (x *ListTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksRequest.ProtoReflect.Descriptor instead.
func (*ListTasksRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_task_v1_task_proto_rawDescGZIP(), []int{6}
}

func (x *ListTasksRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTasksRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListTasksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tasks         []*Task                `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	NextPageToken string                 `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTasksResponse) Reset() {
	*x = ListTasksResponse{}
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTasksResponse) ProtoMessage() {}

func (x *ListTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_task_v1_task_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTasksResponse.ProtoReflect.Descriptor instead.
func (*ListTasksResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_task_v1_task_proto_rawDescGZIP(), []int{7}
}

func (x *ListTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

func (x *ListTasksResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_qclaogui_task_v1_task_proto protoreflect.FileDescriptor

const file_qclaogui_task_v1_task_proto_rawDesc = "" +
	"\n" +
	"\x1bqclaogui/task/v1/task.proto\x12\x10qclaogui.task.v1\x1a\x1bbuf/validate/validate.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a google/protobuf/field_mask.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x96\x03\n" +
	"\x04Task\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1d\n" +
	"\x05title\x18\x02 \x01(\tB\a\xbaH\x04r\x02\x10\x03R\x05title\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x126\n" +
	"\bdeadline\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\bdeadline\x12=\n" +
	"\fcompleted_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\vcompletedAt\x12J\n" +
	"\vcreate_time\x18\xe8\a \x01(\v2\x1a.google.protobuf.TimestampB\f\xe2A\x01\x03\xbaH\x05\xb2\x01\x028\x01R\n" +
	"createTime\x12B\n" +
	"\vupdate_time\x18\xe9\a \x01(\v2\x1a.google.protobuf.TimestampB\x04\xe2A\x01\x03R\n" +
	"updateTime:6\xeaA3\n" +
	"\x16task.qclaogui.com/Task\x12\ftasks/{task}*\x05tasks2\x04task\"S\n" +
	"\x11CreateTaskRequest\x120\n" +
	"\x04task\x18\x02 \x01(\v2\x16.qclaogui.task.v1.TaskB\x04\xe2A\x01\x02R\x04taskJ\x04\b\x01\x10\x02R\x06parent\"\x9d\x01\n" +
	"\x11UpdateTaskRequest\x12K\n" +
	"\x04task\x18\x01 \x01(\v2\x16.qclaogui.task.v1.TaskB\x1f\xe2A\x01\x02\xfaA\x18\n" +
	"\x16task.qclaogui.com/TaskR\x04task\x12;\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\"K\n" +
	"\x11DeleteTaskRequest\x126\n" +
	"\x02id\x18\x01 \x01(\x03B&\xe2A\x01\x02\xfaA\x18\n" +
	"\x16task.qclaogui.com/Task\xbaH\x04\"\x02 \x00R\x02id\"M\n" +
	"\x13UndeleteTaskRequest\x126\n" +
	"\x02id\x18\x01 \x01(\x03B&\xe2A\x01\x02\xfaA\x18\n" +
	"\x16task.qclaogui.com/Task\xbaH\x04\"\x02 \x00R\x02id\"H\n" +
	"\x0eGetTaskRequest\x126\n" +
	"\x02id\x18\x01 \x01(\x03B&\xe2A\x01\x02\xfaA\x18\n" +
	"\x16task.qclaogui.com/Task\xbaH\x04\"\x02 \x00R\x02id\"h\n" +
	"\x10ListTasksRequest\x12!\n" +
	"\tpage_size\x18\x02 \x01(\x05B\x04\xe2A\x01\x01R\bpageSize\x12#\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tB\x04\xe2A\x01\x01R\tpageTokenJ\x04\b\x01\x10\x02R\x06parent\"i\n" +
	"\x11ListTasksResponse\x12,\n" +
	"\x05tasks\x18\x01 \x03(\v2\x16.qclaogui.task.v1.TaskR\x05tasks\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageTokenB5Z3github.com/qclaogui/gaip/genproto/task/apiv1/taskpbb\x06proto3"

var (
	file_qclaogui_task_v1_task_proto_rawDescOnce sync.Once
	file_qclaogui_task_v1_task_proto_rawDescData []byte
)

func file_qclaogui_task_v1_task_proto_rawDescGZIP() []byte {
	file_qclaogui_task_v1_task_proto_rawDescOnce.Do(func() {
		file_qclaogui_task_v1_task_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_task_v1_task_proto_rawDesc), len(file_qclaogui_task_v1_task_proto_rawDesc)))
	})
	return file_qclaogui_task_v1_task_proto_rawDescData
}

var (
	file_qclaogui_task_v1_task_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
	file_qclaogui_task_v1_task_proto_goTypes  = []any{
		(*Task)(nil),                  // 0: qclaogui.task.v1.Task
		(*CreateTaskRequest)(nil),     // 1: qclaogui.task.v1.CreateTaskRequest
		(*UpdateTaskRequest)(nil),     // 2: qclaogui.task.v1.UpdateTaskRequest
		(*DeleteTaskRequest)(nil),     // 3: qclaogui.task.v1.DeleteTaskRequest
		(*UndeleteTaskRequest)(nil),   // 4: qclaogui.task.v1.UndeleteTaskRequest
		(*GetTaskRequest)(nil),        // 5: qclaogui.task.v1.GetTaskRequest
		(*ListTasksRequest)(nil),      // 6: qclaogui.task.v1.ListTasksRequest
		(*ListTasksResponse)(nil),     // 7: qclaogui.task.v1.ListTasksResponse
		(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
		(*fieldmaskpb.FieldMask)(nil), // 9: google.protobuf.FieldMask
	}
)

var file_qclaogui_task_v1_task_proto_depIdxs = []int32{
	8, // 0: qclaogui.task.v1.Task.deadline:type_name -> google.protobuf.Timestamp
	8, // 1: qclaogui.task.v1.Task.completed_at:type_name -> google.protobuf.Timestamp
	8, // 2: qclaogui.task.v1.Task.create_time:type_name -> google.protobuf.Timestamp
	8, // 3: qclaogui.task.v1.Task.update_time:type_name -> google.protobuf.Timestamp
	0, // 4: qclaogui.task.v1.CreateTaskRequest.task:type_name -> qclaogui.task.v1.Task
	0, // 5: qclaogui.task.v1.UpdateTaskRequest.task:type_name -> qclaogui.task.v1.Task
	9, // 6: qclaogui.task.v1.UpdateTaskRequest.update_mask:type_name -> google.protobuf.FieldMask
	0, // 7: qclaogui.task.v1.ListTasksResponse.tasks:type_name -> qclaogui.task.v1.Task
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_qclaogui_task_v1_task_proto_init() }
func file_qclaogui_task_v1_task_proto_init() {
	if File_qclaogui_task_v1_task_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_task_v1_task_proto_rawDesc), len(file_qclaogui_task_v1_task_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qclaogui_task_v1_task_proto_goTypes,
		DependencyIndexes: file_qclaogui_task_v1_task_proto_depIdxs,
		MessageInfos:      file_qclaogui_task_v1_task_proto_msgTypes,
	}.Build()
	File_qclaogui_task_v1_task_proto = out.File
	file_qclaogui_task_v1_task_proto_goTypes = nil
	file_qclaogui_task_v1_task_proto_depIdxs = nil
}
