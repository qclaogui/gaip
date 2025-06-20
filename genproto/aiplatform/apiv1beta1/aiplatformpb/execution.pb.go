// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: qclaogui/aiplatform/v1beta1/execution.proto

package aiplatformpb

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Describes the state of the Execution.
type Execution_State int32

const (
	// Unspecified Execution state
	Execution_STATE_UNSPECIFIED Execution_State = 0
	// The Execution is new
	Execution_NEW Execution_State = 1
	// The Execution is running
	Execution_RUNNING Execution_State = 2
	// The Execution has finished running
	Execution_COMPLETE Execution_State = 3
	// The Execution has failed
	Execution_FAILED Execution_State = 4
	// The Execution completed through Cache hit.
	Execution_CACHED Execution_State = 5
	// The Execution was cancelled.
	Execution_CANCELLED Execution_State = 6
)

// Enum value maps for Execution_State.
var (
	Execution_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "NEW",
		2: "RUNNING",
		3: "COMPLETE",
		4: "FAILED",
		5: "CACHED",
		6: "CANCELLED",
	}
	Execution_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"NEW":               1,
		"RUNNING":           2,
		"COMPLETE":          3,
		"FAILED":            4,
		"CACHED":            5,
		"CANCELLED":         6,
	}
)

func (x Execution_State) Enum() *Execution_State {
	p := new(Execution_State)
	*p = x
	return p
}

func (x Execution_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Execution_State) Descriptor() protoreflect.EnumDescriptor {
	return file_qclaogui_aiplatform_v1beta1_execution_proto_enumTypes[0].Descriptor()
}

func (Execution_State) Type() protoreflect.EnumType {
	return &file_qclaogui_aiplatform_v1beta1_execution_proto_enumTypes[0]
}

func (x Execution_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Execution_State.Descriptor instead.
func (Execution_State) EnumDescriptor() ([]byte, []int) {
	return file_qclaogui_aiplatform_v1beta1_execution_proto_rawDescGZIP(), []int{0, 0}
}

// Instance of a general execution.
type Execution struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The resource name of the Execution.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// User provided display name of the Execution.
	// May be up to 128 Unicode characters.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The state of this Execution. This is a property of the Execution, and does
	// not imply or capture any ongoing process. This property is managed by
	// clients (such as Vertex AI Pipelines) and the system does not prescribe
	// or check the validity of state transitions.
	State Execution_State `protobuf:"varint,6,opt,name=state,proto3,enum=qclaogui.aiplatform.v1beta1.Execution_State" json:"state,omitempty"`
	// An eTag used to perform consistent read-modify-write updates. If not set, a
	// blind "overwrite" update happens.
	Etag string `protobuf:"bytes,9,opt,name=etag,proto3" json:"etag,omitempty"`
	// The labels with user-defined metadata to organize your Executions.
	//
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// No more than 64 user labels can be associated with one Execution (System
	// labels are excluded).
	Labels map[string]string `protobuf:"bytes,10,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Output only. Timestamp when this Execution was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Output only. Timestamp when this Execution was last updated.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// The title of the schema describing the metadata.
	//
	// Schema title and version is expected to be registered in earlier Create
	// Schema calls. And both are used together as unique identifiers to identify
	// schemas within the local metadata store.
	SchemaTitle string `protobuf:"bytes,13,opt,name=schema_title,json=schemaTitle,proto3" json:"schema_title,omitempty"`
	// The version of the schema in `schema_title` to use.
	//
	// Schema title and version is expected to be registered in earlier Create
	// Schema calls. And both are used together as unique identifiers to identify
	// schemas within the local metadata store.
	SchemaVersion string `protobuf:"bytes,14,opt,name=schema_version,json=schemaVersion,proto3" json:"schema_version,omitempty"`
	// Properties of the Execution.
	// Top level metadata keys' heading and trailing spaces will be trimmed.
	// The size of this field should not exceed 200KB.
	Metadata *structpb.Struct `protobuf:"bytes,15,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Description of the Execution
	Description   string `protobuf:"bytes,16,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Execution) Reset() {
	*x = Execution{}
	mi := &file_qclaogui_aiplatform_v1beta1_execution_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Execution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Execution) ProtoMessage() {}

func (x *Execution) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_aiplatform_v1beta1_execution_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Execution.ProtoReflect.Descriptor instead.
func (*Execution) Descriptor() ([]byte, []int) {
	return file_qclaogui_aiplatform_v1beta1_execution_proto_rawDescGZIP(), []int{0}
}

func (x *Execution) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Execution) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Execution) GetState() Execution_State {
	if x != nil {
		return x.State
	}
	return Execution_STATE_UNSPECIFIED
}

func (x *Execution) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Execution) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Execution) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Execution) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *Execution) GetSchemaTitle() string {
	if x != nil {
		return x.SchemaTitle
	}
	return ""
}

func (x *Execution) GetSchemaVersion() string {
	if x != nil {
		return x.SchemaVersion
	}
	return ""
}

func (x *Execution) GetMetadata() *structpb.Struct {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Execution) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_qclaogui_aiplatform_v1beta1_execution_proto protoreflect.FileDescriptor

const file_qclaogui_aiplatform_v1beta1_execution_proto_rawDesc = "" +
	"\n" +
	"+qclaogui/aiplatform/v1beta1/execution.proto\x12\x1bqclaogui.aiplatform.v1beta1\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xc5\x06\n" +
	"\tExecution\x12\x18\n" +
	"\x04name\x18\x01 \x01(\tB\x04\xe2A\x01\x03R\x04name\x12!\n" +
	"\fdisplay_name\x18\x02 \x01(\tR\vdisplayName\x12B\n" +
	"\x05state\x18\x06 \x01(\x0e2,.qclaogui.aiplatform.v1beta1.Execution.StateR\x05state\x12\x12\n" +
	"\x04etag\x18\t \x01(\tR\x04etag\x12J\n" +
	"\x06labels\x18\n" +
	" \x03(\v22.qclaogui.aiplatform.v1beta1.Execution.LabelsEntryR\x06labels\x12A\n" +
	"\vcreate_time\x18\v \x01(\v2\x1a.google.protobuf.TimestampB\x04\xe2A\x01\x03R\n" +
	"createTime\x12A\n" +
	"\vupdate_time\x18\f \x01(\v2\x1a.google.protobuf.TimestampB\x04\xe2A\x01\x03R\n" +
	"updateTime\x12!\n" +
	"\fschema_title\x18\r \x01(\tR\vschemaTitle\x12%\n" +
	"\x0eschema_version\x18\x0e \x01(\tR\rschemaVersion\x123\n" +
	"\bmetadata\x18\x0f \x01(\v2\x17.google.protobuf.StructR\bmetadata\x12 \n" +
	"\vdescription\x18\x10 \x01(\tR\vdescription\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"i\n" +
	"\x05State\x12\x15\n" +
	"\x11STATE_UNSPECIFIED\x10\x00\x12\a\n" +
	"\x03NEW\x10\x01\x12\v\n" +
	"\aRUNNING\x10\x02\x12\f\n" +
	"\bCOMPLETE\x10\x03\x12\n" +
	"\n" +
	"\x06FAILED\x10\x04\x12\n" +
	"\n" +
	"\x06CACHED\x10\x05\x12\r\n" +
	"\tCANCELLED\x10\x06:\x89\x01\xeaA\x85\x01\n" +
	"#aiplatform.googleapis.com/Execution\x12^projects/{project}/locations/{location}/metadataStores/{metadata_store}/executions/{execution}BFZDgithub.com/qclaogui/gaip/genproto/aiplatform/apiv1beta1/aiplatformpbb\x06proto3"

var (
	file_qclaogui_aiplatform_v1beta1_execution_proto_rawDescOnce sync.Once
	file_qclaogui_aiplatform_v1beta1_execution_proto_rawDescData []byte
)

func file_qclaogui_aiplatform_v1beta1_execution_proto_rawDescGZIP() []byte {
	file_qclaogui_aiplatform_v1beta1_execution_proto_rawDescOnce.Do(func() {
		file_qclaogui_aiplatform_v1beta1_execution_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_aiplatform_v1beta1_execution_proto_rawDesc), len(file_qclaogui_aiplatform_v1beta1_execution_proto_rawDesc)))
	})
	return file_qclaogui_aiplatform_v1beta1_execution_proto_rawDescData
}

var (
	file_qclaogui_aiplatform_v1beta1_execution_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
	file_qclaogui_aiplatform_v1beta1_execution_proto_msgTypes  = make([]protoimpl.MessageInfo, 2)
	file_qclaogui_aiplatform_v1beta1_execution_proto_goTypes   = []any{
		(Execution_State)(0),          // 0: qclaogui.aiplatform.v1beta1.Execution.State
		(*Execution)(nil),             // 1: qclaogui.aiplatform.v1beta1.Execution
		nil,                           // 2: qclaogui.aiplatform.v1beta1.Execution.LabelsEntry
		(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
		(*structpb.Struct)(nil),       // 4: google.protobuf.Struct
	}
)

var file_qclaogui_aiplatform_v1beta1_execution_proto_depIdxs = []int32{
	0, // 0: qclaogui.aiplatform.v1beta1.Execution.state:type_name -> qclaogui.aiplatform.v1beta1.Execution.State
	2, // 1: qclaogui.aiplatform.v1beta1.Execution.labels:type_name -> qclaogui.aiplatform.v1beta1.Execution.LabelsEntry
	3, // 2: qclaogui.aiplatform.v1beta1.Execution.create_time:type_name -> google.protobuf.Timestamp
	3, // 3: qclaogui.aiplatform.v1beta1.Execution.update_time:type_name -> google.protobuf.Timestamp
	4, // 4: qclaogui.aiplatform.v1beta1.Execution.metadata:type_name -> google.protobuf.Struct
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_qclaogui_aiplatform_v1beta1_execution_proto_init() }
func file_qclaogui_aiplatform_v1beta1_execution_proto_init() {
	if File_qclaogui_aiplatform_v1beta1_execution_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_aiplatform_v1beta1_execution_proto_rawDesc), len(file_qclaogui_aiplatform_v1beta1_execution_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qclaogui_aiplatform_v1beta1_execution_proto_goTypes,
		DependencyIndexes: file_qclaogui_aiplatform_v1beta1_execution_proto_depIdxs,
		EnumInfos:         file_qclaogui_aiplatform_v1beta1_execution_proto_enumTypes,
		MessageInfos:      file_qclaogui_aiplatform_v1beta1_execution_proto_msgTypes,
	}.Build()
	File_qclaogui_aiplatform_v1beta1_execution_proto = out.File
	file_qclaogui_aiplatform_v1beta1_execution_proto_goTypes = nil
	file_qclaogui_aiplatform_v1beta1_execution_proto_depIdxs = nil
}
