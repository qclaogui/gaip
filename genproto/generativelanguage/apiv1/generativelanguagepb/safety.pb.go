// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: qclaogui/generativelanguage/v1/safety.proto

package generativelanguagepb

import (
	reflect "reflect"
	sync "sync"

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

// The category of a rating.
//
// These categories cover various kinds of harms that developers
// may wish to adjust.
type HarmCategory int32

const (
	// Category is unspecified.
	HarmCategory_HARM_CATEGORY_UNSPECIFIED HarmCategory = 0
	// Negative or harmful comments targeting identity and/or protected attribute.
	HarmCategory_HARM_CATEGORY_DEROGATORY HarmCategory = 1
	// Content that is rude, disrepspectful, or profane.
	HarmCategory_HARM_CATEGORY_TOXICITY HarmCategory = 2
	// Describes scenarios depictng violence against an individual or group, or
	// general descriptions of gore.
	HarmCategory_HARM_CATEGORY_VIOLENCE HarmCategory = 3
	// Contains references to sexual acts or other lewd content.
	HarmCategory_HARM_CATEGORY_SEXUAL HarmCategory = 4
	// Promotes unchecked medical advice.
	HarmCategory_HARM_CATEGORY_MEDICAL HarmCategory = 5
	// Dangerous content that promotes, facilitates, or encourages harmful acts.
	HarmCategory_HARM_CATEGORY_DANGEROUS HarmCategory = 6
	// Harasment content.
	HarmCategory_HARM_CATEGORY_HARASSMENT HarmCategory = 7
	// Hate speech and content.
	HarmCategory_HARM_CATEGORY_HATE_SPEECH HarmCategory = 8
	// Sexually explicit content.
	HarmCategory_HARM_CATEGORY_SEXUALLY_EXPLICIT HarmCategory = 9
	// Dangerous content.
	HarmCategory_HARM_CATEGORY_DANGEROUS_CONTENT HarmCategory = 10
)

// Enum value maps for HarmCategory.
var (
	HarmCategory_name = map[int32]string{
		0:  "HARM_CATEGORY_UNSPECIFIED",
		1:  "HARM_CATEGORY_DEROGATORY",
		2:  "HARM_CATEGORY_TOXICITY",
		3:  "HARM_CATEGORY_VIOLENCE",
		4:  "HARM_CATEGORY_SEXUAL",
		5:  "HARM_CATEGORY_MEDICAL",
		6:  "HARM_CATEGORY_DANGEROUS",
		7:  "HARM_CATEGORY_HARASSMENT",
		8:  "HARM_CATEGORY_HATE_SPEECH",
		9:  "HARM_CATEGORY_SEXUALLY_EXPLICIT",
		10: "HARM_CATEGORY_DANGEROUS_CONTENT",
	}
	HarmCategory_value = map[string]int32{
		"HARM_CATEGORY_UNSPECIFIED":       0,
		"HARM_CATEGORY_DEROGATORY":        1,
		"HARM_CATEGORY_TOXICITY":          2,
		"HARM_CATEGORY_VIOLENCE":          3,
		"HARM_CATEGORY_SEXUAL":            4,
		"HARM_CATEGORY_MEDICAL":           5,
		"HARM_CATEGORY_DANGEROUS":         6,
		"HARM_CATEGORY_HARASSMENT":        7,
		"HARM_CATEGORY_HATE_SPEECH":       8,
		"HARM_CATEGORY_SEXUALLY_EXPLICIT": 9,
		"HARM_CATEGORY_DANGEROUS_CONTENT": 10,
	}
)

func (x HarmCategory) Enum() *HarmCategory {
	p := new(HarmCategory)
	*p = x
	return p
}

func (x HarmCategory) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HarmCategory) Descriptor() protoreflect.EnumDescriptor {
	return file_qclaogui_generativelanguage_v1_safety_proto_enumTypes[0].Descriptor()
}

func (HarmCategory) Type() protoreflect.EnumType {
	return &file_qclaogui_generativelanguage_v1_safety_proto_enumTypes[0]
}

func (x HarmCategory) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HarmCategory.Descriptor instead.
func (HarmCategory) EnumDescriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1_safety_proto_rawDescGZIP(), []int{0}
}

// The probability that a piece of content is harmful.
//
// The classification system gives the probability of the content being
// unsafe. This does not indicate the severity of harm for a piece of content.
type SafetyRating_HarmProbability int32

const (
	// Probability is unspecified.
	SafetyRating_HARM_PROBABILITY_UNSPECIFIED SafetyRating_HarmProbability = 0
	// Content has a negligible chance of being unsafe.
	SafetyRating_HARM_PROBABILITY_NEGLIGIBLE SafetyRating_HarmProbability = 1
	// Content has a low chance of being unsafe.
	SafetyRating_HARM_PROBABILITY_LOW SafetyRating_HarmProbability = 2
	// Content has a medium chance of being unsafe.
	SafetyRating_HARM_PROBABILITY_MEDIUM SafetyRating_HarmProbability = 3
	// Content has a high chance of being unsafe.
	SafetyRating_HARM_PROBABILITY_HIGH SafetyRating_HarmProbability = 4
)

// Enum value maps for SafetyRating_HarmProbability.
var (
	SafetyRating_HarmProbability_name = map[int32]string{
		0: "HARM_PROBABILITY_UNSPECIFIED",
		1: "HARM_PROBABILITY_NEGLIGIBLE",
		2: "HARM_PROBABILITY_LOW",
		3: "HARM_PROBABILITY_MEDIUM",
		4: "HARM_PROBABILITY_HIGH",
	}
	SafetyRating_HarmProbability_value = map[string]int32{
		"HARM_PROBABILITY_UNSPECIFIED": 0,
		"HARM_PROBABILITY_NEGLIGIBLE":  1,
		"HARM_PROBABILITY_LOW":         2,
		"HARM_PROBABILITY_MEDIUM":      3,
		"HARM_PROBABILITY_HIGH":        4,
	}
)

func (x SafetyRating_HarmProbability) Enum() *SafetyRating_HarmProbability {
	p := new(SafetyRating_HarmProbability)
	*p = x
	return p
}

func (x SafetyRating_HarmProbability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SafetyRating_HarmProbability) Descriptor() protoreflect.EnumDescriptor {
	return file_qclaogui_generativelanguage_v1_safety_proto_enumTypes[1].Descriptor()
}

func (SafetyRating_HarmProbability) Type() protoreflect.EnumType {
	return &file_qclaogui_generativelanguage_v1_safety_proto_enumTypes[1]
}

func (x SafetyRating_HarmProbability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SafetyRating_HarmProbability.Descriptor instead.
func (SafetyRating_HarmProbability) EnumDescriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1_safety_proto_rawDescGZIP(), []int{0, 0}
}

// Block at and beyond a specified harm probability.
type SafetySetting_HarmBlockThreshold int32

const (
	// Threshold is unspecified.
	SafetySetting_HARM_BLOCK_THRESHOLD_UNSPECIFIED SafetySetting_HarmBlockThreshold = 0
	// Content with NEGLIGIBLE will be allowed.
	SafetySetting_HARM_BLOCK_THRESHOLD_BLOCK_LOW_AND_ABOVE SafetySetting_HarmBlockThreshold = 1
	// Content with NEGLIGIBLE and LOW will be allowed.
	SafetySetting_HARM_BLOCK_THRESHOLD_BLOCK_MEDIUM_AND_ABOVE SafetySetting_HarmBlockThreshold = 2
	// Content with NEGLIGIBLE, LOW, and MEDIUM will be allowed.
	SafetySetting_HARM_BLOCK_THRESHOLD_BLOCK_ONLY_HIGH SafetySetting_HarmBlockThreshold = 3
	// All content will be allowed.
	SafetySetting_HARM_BLOCK_THRESHOLD_BLOCK_NONE SafetySetting_HarmBlockThreshold = 4
)

// Enum value maps for SafetySetting_HarmBlockThreshold.
var (
	SafetySetting_HarmBlockThreshold_name = map[int32]string{
		0: "HARM_BLOCK_THRESHOLD_UNSPECIFIED",
		1: "HARM_BLOCK_THRESHOLD_BLOCK_LOW_AND_ABOVE",
		2: "HARM_BLOCK_THRESHOLD_BLOCK_MEDIUM_AND_ABOVE",
		3: "HARM_BLOCK_THRESHOLD_BLOCK_ONLY_HIGH",
		4: "HARM_BLOCK_THRESHOLD_BLOCK_NONE",
	}
	SafetySetting_HarmBlockThreshold_value = map[string]int32{
		"HARM_BLOCK_THRESHOLD_UNSPECIFIED":            0,
		"HARM_BLOCK_THRESHOLD_BLOCK_LOW_AND_ABOVE":    1,
		"HARM_BLOCK_THRESHOLD_BLOCK_MEDIUM_AND_ABOVE": 2,
		"HARM_BLOCK_THRESHOLD_BLOCK_ONLY_HIGH":        3,
		"HARM_BLOCK_THRESHOLD_BLOCK_NONE":             4,
	}
)

func (x SafetySetting_HarmBlockThreshold) Enum() *SafetySetting_HarmBlockThreshold {
	p := new(SafetySetting_HarmBlockThreshold)
	*p = x
	return p
}

func (x SafetySetting_HarmBlockThreshold) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SafetySetting_HarmBlockThreshold) Descriptor() protoreflect.EnumDescriptor {
	return file_qclaogui_generativelanguage_v1_safety_proto_enumTypes[2].Descriptor()
}

func (SafetySetting_HarmBlockThreshold) Type() protoreflect.EnumType {
	return &file_qclaogui_generativelanguage_v1_safety_proto_enumTypes[2]
}

func (x SafetySetting_HarmBlockThreshold) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SafetySetting_HarmBlockThreshold.Descriptor instead.
func (SafetySetting_HarmBlockThreshold) EnumDescriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1_safety_proto_rawDescGZIP(), []int{1, 0}
}

// Safety rating for a piece of content.
//
// The safety rating contains the category of harm and the
// harm probability level in that category for a piece of content.
// Content is classified for safety across a number of
// harm categories and the probability of the harm classification is included
// here.
type SafetyRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The category for this rating.
	Category HarmCategory `protobuf:"varint,3,opt,name=category,proto3,enum=qclaogui.generativelanguage.v1.HarmCategory" json:"category,omitempty"`
	// Required. The probability of harm for this content.
	Probability SafetyRating_HarmProbability `protobuf:"varint,4,opt,name=probability,proto3,enum=qclaogui.generativelanguage.v1.SafetyRating_HarmProbability" json:"probability,omitempty"`
	// Was this content blocked because of this rating?
	Blocked bool `protobuf:"varint,5,opt,name=blocked,proto3" json:"blocked,omitempty"`
}

func (x *SafetyRating) Reset() {
	*x = SafetyRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_generativelanguage_v1_safety_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SafetyRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SafetyRating) ProtoMessage() {}

func (x *SafetyRating) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1_safety_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SafetyRating.ProtoReflect.Descriptor instead.
func (*SafetyRating) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1_safety_proto_rawDescGZIP(), []int{0}
}

func (x *SafetyRating) GetCategory() HarmCategory {
	if x != nil {
		return x.Category
	}
	return HarmCategory_HARM_CATEGORY_UNSPECIFIED
}

func (x *SafetyRating) GetProbability() SafetyRating_HarmProbability {
	if x != nil {
		return x.Probability
	}
	return SafetyRating_HARM_PROBABILITY_UNSPECIFIED
}

func (x *SafetyRating) GetBlocked() bool {
	if x != nil {
		return x.Blocked
	}
	return false
}

// Safety setting, affecting the safety-blocking behavior.
//
// Passing a safety setting for a category changes the allowed proability that
// content is blocked.
type SafetySetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The category for this setting.
	Category HarmCategory `protobuf:"varint,3,opt,name=category,proto3,enum=qclaogui.generativelanguage.v1.HarmCategory" json:"category,omitempty"`
	// Required. Controls the probability threshold at which harm is blocked.
	Threshold SafetySetting_HarmBlockThreshold `protobuf:"varint,4,opt,name=threshold,proto3,enum=qclaogui.generativelanguage.v1.SafetySetting_HarmBlockThreshold" json:"threshold,omitempty"`
}

func (x *SafetySetting) Reset() {
	*x = SafetySetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_qclaogui_generativelanguage_v1_safety_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SafetySetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SafetySetting) ProtoMessage() {}

func (x *SafetySetting) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1_safety_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SafetySetting.ProtoReflect.Descriptor instead.
func (*SafetySetting) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1_safety_proto_rawDescGZIP(), []int{1}
}

func (x *SafetySetting) GetCategory() HarmCategory {
	if x != nil {
		return x.Category
	}
	return HarmCategory_HARM_CATEGORY_UNSPECIFIED
}

func (x *SafetySetting) GetThreshold() SafetySetting_HarmBlockThreshold {
	if x != nil {
		return x.Threshold
	}
	return SafetySetting_HARM_BLOCK_THRESHOLD_UNSPECIFIED
}

var File_qclaogui_generativelanguage_v1_safety_proto protoreflect.FileDescriptor

var file_qclaogui_generativelanguage_v1_safety_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x61, 0x66, 0x65, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x71,
	0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f,
	0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87,
	0x03, 0x0a, 0x0c, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x4e, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2c, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x48, 0x61, 0x72, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x64, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x66, 0x65, 0x74, 0x79, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x48, 0x61, 0x72, 0x6d, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22,
	0xa6, 0x01, 0x0a, 0x0f, 0x48, 0x61, 0x72, 0x6d, 0x50, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x12, 0x20, 0x0a, 0x1c, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x50, 0x52, 0x4f, 0x42,
	0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x50, 0x52,
	0x4f, 0x42, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x45, 0x47, 0x4c, 0x49, 0x47,
	0x49, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x50,
	0x52, 0x4f, 0x42, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x02,
	0x12, 0x1b, 0x0a, 0x17, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x41, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x03, 0x12, 0x19, 0x0a,
	0x15, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x50, 0x52, 0x4f, 0x42, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x04, 0x22, 0xb0, 0x03, 0x0a, 0x0d, 0x53, 0x61, 0x66,
	0x65, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x4e, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x71,
	0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x61,
	0x72, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02,
	0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x64, 0x0a, 0x09, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x40, 0x2e,
	0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x66, 0x65, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x48, 0x61, 0x72,
	0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x42,
	0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x22, 0xe8, 0x01, 0x0a, 0x12, 0x48, 0x61, 0x72, 0x6d, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x24, 0x0a, 0x20, 0x48, 0x41, 0x52, 0x4d, 0x5f,
	0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x2c, 0x0a,
	0x28, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x54, 0x48, 0x52, 0x45,
	0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x4c, 0x4f, 0x57, 0x5f,
	0x41, 0x4e, 0x44, 0x5f, 0x41, 0x42, 0x4f, 0x56, 0x45, 0x10, 0x01, 0x12, 0x2f, 0x0a, 0x2b, 0x48,
	0x41, 0x52, 0x4d, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48,
	0x4f, 0x4c, 0x44, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d,
	0x5f, 0x41, 0x4e, 0x44, 0x5f, 0x41, 0x42, 0x4f, 0x56, 0x45, 0x10, 0x02, 0x12, 0x28, 0x0a, 0x24,
	0x48, 0x41, 0x52, 0x4d, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x53,
	0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x5f,
	0x48, 0x49, 0x47, 0x48, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x42,
	0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x54, 0x48, 0x52, 0x45, 0x53, 0x48, 0x4f, 0x4c, 0x44, 0x5f, 0x42,
	0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x04, 0x2a, 0xdc, 0x02, 0x0a, 0x0c,
	0x48, 0x61, 0x72, 0x6d, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x19,
	0x48, 0x41, 0x52, 0x4d, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x48,
	0x41, 0x52, 0x4d, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x44, 0x45, 0x52,
	0x4f, 0x47, 0x41, 0x54, 0x4f, 0x52, 0x59, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x41, 0x52,
	0x4d, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x54, 0x4f, 0x58, 0x49, 0x43,
	0x49, 0x54, 0x59, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x56, 0x49, 0x4f, 0x4c, 0x45, 0x4e, 0x43, 0x45, 0x10,
	0x03, 0x12, 0x18, 0x0a, 0x14, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x53, 0x45, 0x58, 0x55, 0x41, 0x4c, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x48,
	0x41, 0x52, 0x4d, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x4d, 0x45, 0x44,
	0x49, 0x43, 0x41, 0x4c, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x43,
	0x41, 0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x44, 0x41, 0x4e, 0x47, 0x45, 0x52, 0x4f, 0x55,
	0x53, 0x10, 0x06, 0x12, 0x1c, 0x0a, 0x18, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x43, 0x41, 0x54, 0x45,
	0x47, 0x4f, 0x52, 0x59, 0x5f, 0x48, 0x41, 0x52, 0x41, 0x53, 0x53, 0x4d, 0x45, 0x4e, 0x54, 0x10,
	0x07, 0x12, 0x1d, 0x0a, 0x19, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f,
	0x52, 0x59, 0x5f, 0x48, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x50, 0x45, 0x45, 0x43, 0x48, 0x10, 0x08,
	0x12, 0x23, 0x0a, 0x1f, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x43, 0x41, 0x54, 0x45, 0x47, 0x4f, 0x52,
	0x59, 0x5f, 0x53, 0x45, 0x58, 0x55, 0x41, 0x4c, 0x4c, 0x59, 0x5f, 0x45, 0x58, 0x50, 0x4c, 0x49,
	0x43, 0x49, 0x54, 0x10, 0x09, 0x12, 0x23, 0x0a, 0x1f, 0x48, 0x41, 0x52, 0x4d, 0x5f, 0x43, 0x41,
	0x54, 0x45, 0x47, 0x4f, 0x52, 0x59, 0x5f, 0x44, 0x41, 0x4e, 0x47, 0x45, 0x52, 0x4f, 0x55, 0x53,
	0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x10, 0x0a, 0x42, 0x51, 0x5a, 0x4f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2f, 0x67, 0x61, 0x69, 0x70, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_qclaogui_generativelanguage_v1_safety_proto_rawDescOnce sync.Once
	file_qclaogui_generativelanguage_v1_safety_proto_rawDescData = file_qclaogui_generativelanguage_v1_safety_proto_rawDesc
)

func file_qclaogui_generativelanguage_v1_safety_proto_rawDescGZIP() []byte {
	file_qclaogui_generativelanguage_v1_safety_proto_rawDescOnce.Do(func() {
		file_qclaogui_generativelanguage_v1_safety_proto_rawDescData = protoimpl.X.CompressGZIP(file_qclaogui_generativelanguage_v1_safety_proto_rawDescData)
	})
	return file_qclaogui_generativelanguage_v1_safety_proto_rawDescData
}

var file_qclaogui_generativelanguage_v1_safety_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_qclaogui_generativelanguage_v1_safety_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_qclaogui_generativelanguage_v1_safety_proto_goTypes = []interface{}{
	(HarmCategory)(0),                     // 0: qclaogui.generativelanguage.v1.HarmCategory
	(SafetyRating_HarmProbability)(0),     // 1: qclaogui.generativelanguage.v1.SafetyRating.HarmProbability
	(SafetySetting_HarmBlockThreshold)(0), // 2: qclaogui.generativelanguage.v1.SafetySetting.HarmBlockThreshold
	(*SafetyRating)(nil),                  // 3: qclaogui.generativelanguage.v1.SafetyRating
	(*SafetySetting)(nil),                 // 4: qclaogui.generativelanguage.v1.SafetySetting
}
var file_qclaogui_generativelanguage_v1_safety_proto_depIdxs = []int32{
	0, // 0: qclaogui.generativelanguage.v1.SafetyRating.category:type_name -> qclaogui.generativelanguage.v1.HarmCategory
	1, // 1: qclaogui.generativelanguage.v1.SafetyRating.probability:type_name -> qclaogui.generativelanguage.v1.SafetyRating.HarmProbability
	0, // 2: qclaogui.generativelanguage.v1.SafetySetting.category:type_name -> qclaogui.generativelanguage.v1.HarmCategory
	2, // 3: qclaogui.generativelanguage.v1.SafetySetting.threshold:type_name -> qclaogui.generativelanguage.v1.SafetySetting.HarmBlockThreshold
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_qclaogui_generativelanguage_v1_safety_proto_init() }
func file_qclaogui_generativelanguage_v1_safety_proto_init() {
	if File_qclaogui_generativelanguage_v1_safety_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_qclaogui_generativelanguage_v1_safety_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SafetyRating); i {
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
		file_qclaogui_generativelanguage_v1_safety_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SafetySetting); i {
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
			RawDescriptor: file_qclaogui_generativelanguage_v1_safety_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_qclaogui_generativelanguage_v1_safety_proto_goTypes,
		DependencyIndexes: file_qclaogui_generativelanguage_v1_safety_proto_depIdxs,
		EnumInfos:         file_qclaogui_generativelanguage_v1_safety_proto_enumTypes,
		MessageInfos:      file_qclaogui_generativelanguage_v1_safety_proto_msgTypes,
	}.Build()
	File_qclaogui_generativelanguage_v1_safety_proto = out.File
	file_qclaogui_generativelanguage_v1_safety_proto_rawDesc = nil
	file_qclaogui_generativelanguage_v1_safety_proto_goTypes = nil
	file_qclaogui_generativelanguage_v1_safety_proto_depIdxs = nil
}
