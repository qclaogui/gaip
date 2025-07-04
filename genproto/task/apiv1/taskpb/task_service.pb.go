// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: qclaogui/task/v1/task_service.proto

package taskpb

import (
	reflect "reflect"
	unsafe "unsafe"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_qclaogui_task_v1_task_service_proto protoreflect.FileDescriptor

const file_qclaogui_task_v1_task_service_proto_rawDesc = "" +
	"\n" +
	"#qclaogui/task/v1/task_service.proto\x12\x10qclaogui.task.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1bqclaogui/task/v1/task.proto2\xc4\x05\n" +
	"\fTasksService\x12i\n" +
	"\n" +
	"CreateTask\x12#.qclaogui.task.v1.CreateTaskRequest\x1a\x16.qclaogui.task.v1.Task\"\x1e\xdaA\x04task\x82\xd3\xe4\x93\x02\x11:\x04task\"\t/v1/tasks\x12f\n" +
	"\n" +
	"DeleteTask\x12#.qclaogui.task.v1.DeleteTaskRequest\x1a\x16.qclaogui.task.v1.Task\"\x1b\xdaA\x02id\x82\xd3\xe4\x93\x02\x10*\x0e/v1/tasks/{id}\x12q\n" +
	"\fUndeleteTask\x12%.qclaogui.task.v1.UndeleteTaskRequest\x1a\x16.qclaogui.task.v1.Task\"\"\xdaA\x02id\x82\xd3\xe4\x93\x02\x17:\x01*\"\x12/v1/tasks:undelete\x12}\n" +
	"\n" +
	"UpdateTask\x12#.qclaogui.task.v1.UpdateTaskRequest\x1a\x16.qclaogui.task.v1.Task\"2\xdaA\x0eid,update_mask\x82\xd3\xe4\x93\x02\x1b:\x04task2\x13/v1/tasks/{task.id}\x12b\n" +
	"\aGetTask\x12 .qclaogui.task.v1.GetTaskRequest\x1a\x16.qclaogui.task.v1.Task\"\x1d\xdaA\x04task\x82\xd3\xe4\x93\x02\x10\x12\x0e/v1/tasks/{id}\x12g\n" +
	"\tListTasks\x12\".qclaogui.task.v1.ListTasksRequest\x1a#.qclaogui.task.v1.ListTasksResponse\"\x11\x82\xd3\xe4\x93\x02\v\x12\t/v1/tasks\x1a\"\xcaA\x0elocalhost:9095\x8a\xd4\xdb\xd2\x0f\vv1_20240506B5Z3github.com/qclaogui/gaip/genproto/task/apiv1/taskpbb\x06proto3"

var file_qclaogui_task_v1_task_service_proto_goTypes = []any{
	(*CreateTaskRequest)(nil),   // 0: qclaogui.task.v1.CreateTaskRequest
	(*DeleteTaskRequest)(nil),   // 1: qclaogui.task.v1.DeleteTaskRequest
	(*UndeleteTaskRequest)(nil), // 2: qclaogui.task.v1.UndeleteTaskRequest
	(*UpdateTaskRequest)(nil),   // 3: qclaogui.task.v1.UpdateTaskRequest
	(*GetTaskRequest)(nil),      // 4: qclaogui.task.v1.GetTaskRequest
	(*ListTasksRequest)(nil),    // 5: qclaogui.task.v1.ListTasksRequest
	(*Task)(nil),                // 6: qclaogui.task.v1.Task
	(*ListTasksResponse)(nil),   // 7: qclaogui.task.v1.ListTasksResponse
}

var file_qclaogui_task_v1_task_service_proto_depIdxs = []int32{
	0, // 0: qclaogui.task.v1.TasksService.CreateTask:input_type -> qclaogui.task.v1.CreateTaskRequest
	1, // 1: qclaogui.task.v1.TasksService.DeleteTask:input_type -> qclaogui.task.v1.DeleteTaskRequest
	2, // 2: qclaogui.task.v1.TasksService.UndeleteTask:input_type -> qclaogui.task.v1.UndeleteTaskRequest
	3, // 3: qclaogui.task.v1.TasksService.UpdateTask:input_type -> qclaogui.task.v1.UpdateTaskRequest
	4, // 4: qclaogui.task.v1.TasksService.GetTask:input_type -> qclaogui.task.v1.GetTaskRequest
	5, // 5: qclaogui.task.v1.TasksService.ListTasks:input_type -> qclaogui.task.v1.ListTasksRequest
	6, // 6: qclaogui.task.v1.TasksService.CreateTask:output_type -> qclaogui.task.v1.Task
	6, // 7: qclaogui.task.v1.TasksService.DeleteTask:output_type -> qclaogui.task.v1.Task
	6, // 8: qclaogui.task.v1.TasksService.UndeleteTask:output_type -> qclaogui.task.v1.Task
	6, // 9: qclaogui.task.v1.TasksService.UpdateTask:output_type -> qclaogui.task.v1.Task
	6, // 10: qclaogui.task.v1.TasksService.GetTask:output_type -> qclaogui.task.v1.Task
	7, // 11: qclaogui.task.v1.TasksService.ListTasks:output_type -> qclaogui.task.v1.ListTasksResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_qclaogui_task_v1_task_service_proto_init() }
func file_qclaogui_task_v1_task_service_proto_init() {
	if File_qclaogui_task_v1_task_service_proto != nil {
		return
	}
	file_qclaogui_task_v1_task_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_task_v1_task_service_proto_rawDesc), len(file_qclaogui_task_v1_task_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qclaogui_task_v1_task_service_proto_goTypes,
		DependencyIndexes: file_qclaogui_task_v1_task_service_proto_depIdxs,
	}.Build()
	File_qclaogui_task_v1_task_service_proto = out.File
	file_qclaogui_task_v1_task_service_proto_goTypes = nil
	file_qclaogui_task_v1_task_service_proto_depIdxs = nil
}
