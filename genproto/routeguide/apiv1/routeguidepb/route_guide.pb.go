// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: qclaogui/routeguide/v1/route_guide.proto

package routeguidepb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetFeatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point *Point `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *GetFeatureRequest) Reset() {
	*x = GetFeatureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureRequest) ProtoMessage() {}

func (x *GetFeatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureRequest.ProtoReflect.Descriptor instead.
func (*GetFeatureRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{0}
}

func (x *GetFeatureRequest) GetPoint() *Point {
	if x != nil {
		return x.Point
	}
	return nil
}

type GetFeatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature *Feature `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *GetFeatureResponse) Reset() {
	*x = GetFeatureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFeatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFeatureResponse) ProtoMessage() {}

func (x *GetFeatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFeatureResponse.ProtoReflect.Descriptor instead.
func (*GetFeatureResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{1}
}

func (x *GetFeatureResponse) GetFeature() *Feature {
	if x != nil {
		return x.Feature
	}
	return nil
}

type ListFeaturesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rectangle *Rectangle `protobuf:"bytes,1,opt,name=rectangle,proto3" json:"rectangle,omitempty"`
}

func (x *ListFeaturesRequest) Reset() {
	*x = ListFeaturesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeaturesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeaturesRequest) ProtoMessage() {}

func (x *ListFeaturesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeaturesRequest.ProtoReflect.Descriptor instead.
func (*ListFeaturesRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{2}
}

func (x *ListFeaturesRequest) GetRectangle() *Rectangle {
	if x != nil {
		return x.Rectangle
	}
	return nil
}

type ListFeaturesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature *Feature `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *ListFeaturesResponse) Reset() {
	*x = ListFeaturesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFeaturesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFeaturesResponse) ProtoMessage() {}

func (x *ListFeaturesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFeaturesResponse.ProtoReflect.Descriptor instead.
func (*ListFeaturesResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{3}
}

func (x *ListFeaturesResponse) GetFeature() *Feature {
	if x != nil {
		return x.Feature
	}
	return nil
}

type RecordRouteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point *Point `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *RecordRouteRequest) Reset() {
	*x = RecordRouteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordRouteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordRouteRequest) ProtoMessage() {}

func (x *RecordRouteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordRouteRequest.ProtoReflect.Descriptor instead.
func (*RecordRouteRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{4}
}

func (x *RecordRouteRequest) GetPoint() *Point {
	if x != nil {
		return x.Point
	}
	return nil
}

type RecordRouteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteSummary *RouteSummary `protobuf:"bytes,1,opt,name=route_summary,json=routeSummary,proto3" json:"route_summary,omitempty"`
}

func (x *RecordRouteResponse) Reset() {
	*x = RecordRouteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordRouteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordRouteResponse) ProtoMessage() {}

func (x *RecordRouteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordRouteResponse.ProtoReflect.Descriptor instead.
func (*RecordRouteResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{5}
}

func (x *RecordRouteResponse) GetRouteSummary() *RouteSummary {
	if x != nil {
		return x.RouteSummary
	}
	return nil
}

// A message sent while at a given point.
type RouteChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteNote *RouteNote `protobuf:"bytes,1,opt,name=route_note,json=routeNote,proto3" json:"route_note,omitempty"`
}

func (x *RouteChatRequest) Reset() {
	*x = RouteChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteChatRequest) ProtoMessage() {}

func (x *RouteChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteChatRequest.ProtoReflect.Descriptor instead.
func (*RouteChatRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{6}
}

func (x *RouteChatRequest) GetRouteNote() *RouteNote {
	if x != nil {
		return x.RouteNote
	}
	return nil
}

// A message response at a given point.
type RouteChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RouteNote *RouteNote `protobuf:"bytes,1,opt,name=route_note,json=routeNote,proto3" json:"route_note,omitempty"`
}

func (x *RouteChatResponse) Reset() {
	*x = RouteChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteChatResponse) ProtoMessage() {}

func (x *RouteChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteChatResponse.ProtoReflect.Descriptor instead.
func (*RouteChatResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{7}
}

func (x *RouteChatResponse) GetRouteNote() *RouteNote {
	if x != nil {
		return x.RouteNote
	}
	return nil
}

// Points are represented as latitude-longitude pairs in the E7 representation
// (degrees multiplied by 10**7 and rounded to the nearest integer).
// Latitudes should be in the range +/- 90 degrees and longitude should be in
// the range +/- 180 degrees (inclusive).
type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  int32 `protobuf:"varint,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude int32 `protobuf:"varint,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{8}
}

func (x *Point) GetLatitude() int32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Point) GetLongitude() int32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

// A latitude-longitude rectangle, represented as two diagonally opposite
// points "lo" and "hi".
type Rectangle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One corner of the rectangle.
	Lo *Point `protobuf:"bytes,1,opt,name=lo,proto3" json:"lo,omitempty"`
	// The other corner of the rectangle.
	Hi *Point `protobuf:"bytes,2,opt,name=hi,proto3" json:"hi,omitempty"`
}

func (x *Rectangle) Reset() {
	*x = Rectangle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rectangle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rectangle) ProtoMessage() {}

func (x *Rectangle) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rectangle.ProtoReflect.Descriptor instead.
func (*Rectangle) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{9}
}

func (x *Rectangle) GetLo() *Point {
	if x != nil {
		return x.Lo
	}
	return nil
}

func (x *Rectangle) GetHi() *Point {
	if x != nil {
		return x.Hi
	}
	return nil
}

// A feature names something at a given point.
//
// If a feature could not be named, the name is empty.
type Feature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the feature.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The point where the feature is detected.
	Location *Point `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *Feature) Reset() {
	*x = Feature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Feature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Feature) ProtoMessage() {}

func (x *Feature) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Feature.ProtoReflect.Descriptor instead.
func (*Feature) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{10}
}

func (x *Feature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Feature) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

// A RouteNote is a message sent while at a given point.
type RouteNote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The location from which the message is sent.
	Location *Point `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	// The message to be sent.
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RouteNote) Reset() {
	*x = RouteNote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteNote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteNote) ProtoMessage() {}

func (x *RouteNote) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteNote.ProtoReflect.Descriptor instead.
func (*RouteNote) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{11}
}

func (x *RouteNote) GetLocation() *Point {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *RouteNote) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// A RouteSummary is received in response to a RecordRoute rpc.
//
// It contains the number of individual points received, the number of
// detected features, and the total distance covered as the cumulative sum of
// the distance between each point.
type RouteSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of points received.
	PointCount int32 `protobuf:"varint,1,opt,name=point_count,json=pointCount,proto3" json:"point_count,omitempty"`
	// The number of known features passed while traversing the route.
	FeatureCount int32 `protobuf:"varint,2,opt,name=feature_count,json=featureCount,proto3" json:"feature_count,omitempty"`
	// The distance covered in metres.
	Distance int32 `protobuf:"varint,3,opt,name=distance,proto3" json:"distance,omitempty"`
	// The duration of the traversal in seconds.
	ElapsedTime int32 `protobuf:"varint,4,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"`
}

func (x *RouteSummary) Reset() {
	*x = RouteSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteSummary) ProtoMessage() {}

func (x *RouteSummary) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteSummary.ProtoReflect.Descriptor instead.
func (*RouteSummary) Descriptor() ([]byte, []int) {
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP(), []int{12}
}

func (x *RouteSummary) GetPointCount() int32 {
	if x != nil {
		return x.PointCount
	}
	return 0
}

func (x *RouteSummary) GetFeatureCount() int32 {
	if x != nil {
		return x.FeatureCount
	}
	return 0
}

func (x *RouteSummary) GetDistance() int32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *RouteSummary) GetElapsedTime() int32 {
	if x != nil {
		return x.ElapsedTime
	}
	return 0
}

var File_qclaogui_routeguide_v1_route_guide_proto protoreflect.FileDescriptor

var file_qclaogui_routeguide_v1_route_guide_proto_rawDesc = []byte{
	0x0a, 0x28, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x22, 0x48, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x22, 0x4f, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x56, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x74, 0x61, 0x6e, 0x67, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67,
	0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x63, 0x74, 0x61, 0x6e, 0x67, 0x6c, 0x65, 0x52, 0x09, 0x72, 0x65, 0x63, 0x74,
	0x61, 0x6e, 0x67, 0x6c, 0x65, 0x22, 0x51, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x49, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33,
	0x0a, 0x05, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x22, 0x60, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0d, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x75,
	0x6d, 0x6d, 0x61, 0x72, 0x79, 0x22, 0x54, 0x0a, 0x10, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0a, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x22, 0x55, 0x0a, 0x11, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x40, 0x0a, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x09, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x4e, 0x6f,
	0x74, 0x65, 0x22, 0x41, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x69, 0x0a, 0x09, 0x52, 0x65, 0x63, 0x74, 0x61, 0x6e, 0x67,
	0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x02, 0x6c, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67,
	0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x6c,
	0x6f, 0x12, 0x2d, 0x0a, 0x02, 0x68, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x68, 0x69,
	0x22, 0x58, 0x0a, 0x07, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x39, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x60, 0x0a, 0x09, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x93, 0x01, 0x0a,
	0x0c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x32, 0xbd, 0x03, 0x0a, 0x11, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x47, 0x75, 0x69, 0x64,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x6d, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12,
	0x2b, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x71,
	0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x6a,
	0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2a, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x71, 0x63, 0x6c, 0x61,
	0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x66, 0x0a, 0x09, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x28, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67,
	0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x29, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x67, 0x75, 0x69, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x61, 0x69, 0x70, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75, 0x69,
	0x64, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x67, 0x75,
	0x69, 0x64, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qclaogui_routeguide_v1_route_guide_proto_rawDescOnce sync.Once
	file_qclaogui_routeguide_v1_route_guide_proto_rawDescData = file_qclaogui_routeguide_v1_route_guide_proto_rawDesc
)

func file_qclaogui_routeguide_v1_route_guide_proto_rawDescGZIP() []byte {
	file_qclaogui_routeguide_v1_route_guide_proto_rawDescOnce.Do(func() {
		file_qclaogui_routeguide_v1_route_guide_proto_rawDescData = protoimpl.X.CompressGZIP(file_qclaogui_routeguide_v1_route_guide_proto_rawDescData)
	})
	return file_qclaogui_routeguide_v1_route_guide_proto_rawDescData
}

var file_qclaogui_routeguide_v1_route_guide_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_qclaogui_routeguide_v1_route_guide_proto_goTypes = []interface{}{
	(*GetFeatureRequest)(nil),    // 0: qclaogui.routeguide.v1.GetFeatureRequest
	(*GetFeatureResponse)(nil),   // 1: qclaogui.routeguide.v1.GetFeatureResponse
	(*ListFeaturesRequest)(nil),  // 2: qclaogui.routeguide.v1.ListFeaturesRequest
	(*ListFeaturesResponse)(nil), // 3: qclaogui.routeguide.v1.ListFeaturesResponse
	(*RecordRouteRequest)(nil),   // 4: qclaogui.routeguide.v1.RecordRouteRequest
	(*RecordRouteResponse)(nil),  // 5: qclaogui.routeguide.v1.RecordRouteResponse
	(*RouteChatRequest)(nil),     // 6: qclaogui.routeguide.v1.RouteChatRequest
	(*RouteChatResponse)(nil),    // 7: qclaogui.routeguide.v1.RouteChatResponse
	(*Point)(nil),                // 8: qclaogui.routeguide.v1.Point
	(*Rectangle)(nil),            // 9: qclaogui.routeguide.v1.Rectangle
	(*Feature)(nil),              // 10: qclaogui.routeguide.v1.Feature
	(*RouteNote)(nil),            // 11: qclaogui.routeguide.v1.RouteNote
	(*RouteSummary)(nil),         // 12: qclaogui.routeguide.v1.RouteSummary
}
var file_qclaogui_routeguide_v1_route_guide_proto_depIdxs = []int32{
	8,  // 0: qclaogui.routeguide.v1.GetFeatureRequest.point:type_name -> qclaogui.routeguide.v1.Point
	10, // 1: qclaogui.routeguide.v1.GetFeatureResponse.feature:type_name -> qclaogui.routeguide.v1.Feature
	9,  // 2: qclaogui.routeguide.v1.ListFeaturesRequest.rectangle:type_name -> qclaogui.routeguide.v1.Rectangle
	10, // 3: qclaogui.routeguide.v1.ListFeaturesResponse.feature:type_name -> qclaogui.routeguide.v1.Feature
	8,  // 4: qclaogui.routeguide.v1.RecordRouteRequest.point:type_name -> qclaogui.routeguide.v1.Point
	12, // 5: qclaogui.routeguide.v1.RecordRouteResponse.route_summary:type_name -> qclaogui.routeguide.v1.RouteSummary
	11, // 6: qclaogui.routeguide.v1.RouteChatRequest.route_note:type_name -> qclaogui.routeguide.v1.RouteNote
	11, // 7: qclaogui.routeguide.v1.RouteChatResponse.route_note:type_name -> qclaogui.routeguide.v1.RouteNote
	8,  // 8: qclaogui.routeguide.v1.Rectangle.lo:type_name -> qclaogui.routeguide.v1.Point
	8,  // 9: qclaogui.routeguide.v1.Rectangle.hi:type_name -> qclaogui.routeguide.v1.Point
	8,  // 10: qclaogui.routeguide.v1.Feature.location:type_name -> qclaogui.routeguide.v1.Point
	8,  // 11: qclaogui.routeguide.v1.RouteNote.location:type_name -> qclaogui.routeguide.v1.Point
	0,  // 12: qclaogui.routeguide.v1.RouteGuideService.GetFeature:input_type -> qclaogui.routeguide.v1.GetFeatureRequest
	2,  // 13: qclaogui.routeguide.v1.RouteGuideService.ListFeatures:input_type -> qclaogui.routeguide.v1.ListFeaturesRequest
	4,  // 14: qclaogui.routeguide.v1.RouteGuideService.RecordRoute:input_type -> qclaogui.routeguide.v1.RecordRouteRequest
	6,  // 15: qclaogui.routeguide.v1.RouteGuideService.RouteChat:input_type -> qclaogui.routeguide.v1.RouteChatRequest
	1,  // 16: qclaogui.routeguide.v1.RouteGuideService.GetFeature:output_type -> qclaogui.routeguide.v1.GetFeatureResponse
	3,  // 17: qclaogui.routeguide.v1.RouteGuideService.ListFeatures:output_type -> qclaogui.routeguide.v1.ListFeaturesResponse
	5,  // 18: qclaogui.routeguide.v1.RouteGuideService.RecordRoute:output_type -> qclaogui.routeguide.v1.RecordRouteResponse
	7,  // 19: qclaogui.routeguide.v1.RouteGuideService.RouteChat:output_type -> qclaogui.routeguide.v1.RouteChatResponse
	16, // [16:20] is the sub-list for method output_type
	12, // [12:16] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_qclaogui_routeguide_v1_route_guide_proto_init() }
func file_qclaogui_routeguide_v1_route_guide_proto_init() {
	if File_qclaogui_routeguide_v1_route_guide_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeatureRequest); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFeatureResponse); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeaturesRequest); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFeaturesResponse); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordRouteRequest); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordRouteResponse); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteChatRequest); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteChatResponse); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rectangle); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Feature); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteNote); i {
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
		file_qclaogui_routeguide_v1_route_guide_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteSummary); i {
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
			RawDescriptor: file_qclaogui_routeguide_v1_route_guide_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qclaogui_routeguide_v1_route_guide_proto_goTypes,
		DependencyIndexes: file_qclaogui_routeguide_v1_route_guide_proto_depIdxs,
		MessageInfos:      file_qclaogui_routeguide_v1_route_guide_proto_msgTypes,
	}.Build()
	File_qclaogui_routeguide_v1_route_guide_proto = out.File
	file_qclaogui_routeguide_v1_route_guide_proto_rawDesc = nil
	file_qclaogui_routeguide_v1_route_guide_proto_goTypes = nil
	file_qclaogui_routeguide_v1_route_guide_proto_depIdxs = nil
}
