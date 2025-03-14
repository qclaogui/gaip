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
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: qclaogui/generativelanguage/v1beta1/discuss_service.proto

package generativelanguagepb

import (
	reflect "reflect"
	sync "sync"
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

// Request to generate a message response from the model.
type GenerateMessageRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The name of the model to use.
	//
	// Format: `name=models/{model}`.
	Model string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	// Required. The structured textual input given to the model as a prompt.
	//
	// Given a
	// prompt, the model will return what it predicts is the next message in the
	// discussion.
	Prompt *MessagePrompt `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	// Optional. Controls the randomness of the output.
	//
	// Values can range over `[0.0,1.0]`,
	// inclusive. A value closer to `1.0` will produce responses that are more
	// varied, while a value closer to `0.0` will typically result in
	// less surprising responses from the model.
	Temperature *float32 `protobuf:"fixed32,3,opt,name=temperature,proto3,oneof" json:"temperature,omitempty"`
	// Optional. The number of generated response messages to return.
	//
	// This value must be between
	// `[1, 8]`, inclusive. If unset, this will default to `1`.
	CandidateCount *int32 `protobuf:"varint,4,opt,name=candidate_count,json=candidateCount,proto3,oneof" json:"candidate_count,omitempty"`
	// Optional. The maximum cumulative probability of tokens to consider when
	// sampling.
	//
	// The model uses combined Top-k and nucleus sampling.
	//
	// Nucleus sampling considers the smallest set of tokens whose probability
	// sum is at least `top_p`.
	TopP *float32 `protobuf:"fixed32,5,opt,name=top_p,json=topP,proto3,oneof" json:"top_p,omitempty"`
	// Optional. The maximum number of tokens to consider when sampling.
	//
	// The model uses combined Top-k and nucleus sampling.
	//
	// Top-k sampling considers the set of `top_k` most probable tokens.
	TopK          *int32 `protobuf:"varint,6,opt,name=top_k,json=topK,proto3,oneof" json:"top_k,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateMessageRequest) Reset() {
	*x = GenerateMessageRequest{}
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateMessageRequest) ProtoMessage() {}

func (x *GenerateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateMessageRequest.ProtoReflect.Descriptor instead.
func (*GenerateMessageRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateMessageRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *GenerateMessageRequest) GetPrompt() *MessagePrompt {
	if x != nil {
		return x.Prompt
	}
	return nil
}

func (x *GenerateMessageRequest) GetTemperature() float32 {
	if x != nil && x.Temperature != nil {
		return *x.Temperature
	}
	return 0
}

func (x *GenerateMessageRequest) GetCandidateCount() int32 {
	if x != nil && x.CandidateCount != nil {
		return *x.CandidateCount
	}
	return 0
}

func (x *GenerateMessageRequest) GetTopP() float32 {
	if x != nil && x.TopP != nil {
		return *x.TopP
	}
	return 0
}

func (x *GenerateMessageRequest) GetTopK() int32 {
	if x != nil && x.TopK != nil {
		return *x.TopK
	}
	return 0
}

// The response from the model.
//
// This includes candidate messages and
// conversation history in the form of chronologically-ordered messages.
type GenerateMessageResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Candidate response messages from the model.
	Candidates []*Message `protobuf:"bytes,1,rep,name=candidates,proto3" json:"candidates,omitempty"`
	// The conversation history used by the model.
	Messages []*Message `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	// A set of content filtering metadata for the prompt and response
	// text.
	//
	// This indicates which `SafetyCategory`(s) blocked a
	// candidate from this response, the lowest `HarmProbability`
	// that triggered a block, and the HarmThreshold setting for that category.
	Filters       []*ContentFilter `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateMessageResponse) Reset() {
	*x = GenerateMessageResponse{}
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateMessageResponse) ProtoMessage() {}

func (x *GenerateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateMessageResponse.ProtoReflect.Descriptor instead.
func (*GenerateMessageResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateMessageResponse) GetCandidates() []*Message {
	if x != nil {
		return x.Candidates
	}
	return nil
}

func (x *GenerateMessageResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *GenerateMessageResponse) GetFilters() []*ContentFilter {
	if x != nil {
		return x.Filters
	}
	return nil
}

// The base unit of structured text.
//
// A `Message` includes an `author` and the `content` of
// the `Message`.
//
// The `author` is used to tag messages when they are fed to the
// model as text.
type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. The author of this Message.
	//
	// This serves as a key for tagging
	// the content of this Message when it is fed to the model as text.
	//
	// The author can be any alphanumeric string.
	Author string `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	// Required. The text content of the structured `Message`.
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	// Output only. Citation information for model-generated `content` in this
	// `Message`.
	//
	// If this `Message` was generated as output from the model, this field may be
	// populated with attribution information for any text included in the
	// `content`. This field is used only on output.
	CitationMetadata *CitationMetadata `protobuf:"bytes,3,opt,name=citation_metadata,json=citationMetadata,proto3,oneof" json:"citation_metadata,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[2]
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
	return file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetCitationMetadata() *CitationMetadata {
	if x != nil {
		return x.CitationMetadata
	}
	return nil
}

// All of the structured input text passed to the model as a prompt.
//
// A `MessagePrompt` contains a structured set of fields that provide context
// for the conversation, examples of user input/model output message pairs that
// prime the model to respond in different ways, and the conversation history
// or list of messages representing the alternating turns of the conversation
// between the user and the model.
type MessagePrompt struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Optional. Text that should be provided to the model first to ground the
	// response.
	//
	// If not empty, this `context` will be given to the model first before the
	// `examples` and `messages`. When using a `context` be sure to provide it
	// with every request to maintain continuity.
	//
	// This field can be a description of your prompt to the model to help provide
	// context and guide the responses. Examples: "Translate the phrase from
	// English to French." or "Given a statement, classify the sentiment as happy,
	// sad or neutral."
	//
	// Anything included in this field will take precedence over message history
	// if the total input size exceeds the model's `input_token_limit` and the
	// input request is truncated.
	Context string `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	// Optional. Examples of what the model should generate.
	//
	// This includes both user input and the response that the model should
	// emulate.
	//
	// These `examples` are treated identically to conversation messages except
	// that they take precedence over the history in `messages`:
	// If the total input size exceeds the model's `input_token_limit` the input
	// will be truncated. Items will be dropped from `messages` before `examples`.
	Examples []*Example `protobuf:"bytes,2,rep,name=examples,proto3" json:"examples,omitempty"`
	// Required. A snapshot of the recent conversation history sorted
	// chronologically.
	//
	// Turns alternate between two authors.
	//
	// If the total input size exceeds the model's `input_token_limit` the input
	// will be truncated: The oldest items will be dropped from `messages`.
	Messages      []*Message `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MessagePrompt) Reset() {
	*x = MessagePrompt{}
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MessagePrompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessagePrompt) ProtoMessage() {}

func (x *MessagePrompt) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessagePrompt.ProtoReflect.Descriptor instead.
func (*MessagePrompt) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescGZIP(), []int{3}
}

func (x *MessagePrompt) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *MessagePrompt) GetExamples() []*Example {
	if x != nil {
		return x.Examples
	}
	return nil
}

func (x *MessagePrompt) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

// An input/output example used to instruct the Model.
//
// It demonstrates how the model should respond or format its response.
type Example struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. An example of an input `Message` from the user.
	Input *Message `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	// Required. An example of what the model should output given the input.
	Output        *Message `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Example) Reset() {
	*x = Example{}
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Example) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Example) ProtoMessage() {}

func (x *Example) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Example.ProtoReflect.Descriptor instead.
func (*Example) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescGZIP(), []int{4}
}

func (x *Example) GetInput() *Message {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Example) GetOutput() *Message {
	if x != nil {
		return x.Output
	}
	return nil
}

// Counts the number of tokens in the `prompt` sent to a model.
//
// Models may tokenize text differently, so each model may return a different
// `token_count`.
type CountMessageTokensRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. The model's resource name. This serves as an ID for the Model to
	// use.
	//
	// This name should match a model name returned by the `ListModels` method.
	//
	// Format: `models/{model}`
	Model string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	// Required. The prompt, whose token count is to be returned.
	Prompt        *MessagePrompt `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CountMessageTokensRequest) Reset() {
	*x = CountMessageTokensRequest{}
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountMessageTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountMessageTokensRequest) ProtoMessage() {}

func (x *CountMessageTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountMessageTokensRequest.ProtoReflect.Descriptor instead.
func (*CountMessageTokensRequest) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescGZIP(), []int{5}
}

func (x *CountMessageTokensRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CountMessageTokensRequest) GetPrompt() *MessagePrompt {
	if x != nil {
		return x.Prompt
	}
	return nil
}

// A response from `CountMessageTokens`.
//
// It returns the model's `token_count` for the `prompt`.
type CountMessageTokensResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The number of tokens that the `model` tokenizes the `prompt` into.
	//
	// Always non-negative.
	TokenCount    int32 `protobuf:"varint,1,opt,name=token_count,json=tokenCount,proto3" json:"token_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CountMessageTokensResponse) Reset() {
	*x = CountMessageTokensResponse{}
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CountMessageTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountMessageTokensResponse) ProtoMessage() {}

func (x *CountMessageTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountMessageTokensResponse.ProtoReflect.Descriptor instead.
func (*CountMessageTokensResponse) Descriptor() ([]byte, []int) {
	return file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescGZIP(), []int{6}
}

func (x *CountMessageTokensResponse) GetTokenCount() int32 {
	if x != nil {
		return x.TokenCount
	}
	return 0
}

var File_qclaogui_generativelanguage_v1beta1_discuss_service_proto protoreflect.FileDescriptor

var file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDesc = string([]byte{
	0x0a, 0x39, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x61, 0x66,
	0x65, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x03, 0x0a, 0x16, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2e, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x27, 0x0a, 0x25, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x50, 0x0a, 0x06, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x42, 0x04,
	0xe2, 0x41, 0x01, 0x02, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x2b, 0x0a, 0x0b,
	0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x0f, 0x63, 0x61, 0x6e,
	0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x48, 0x01, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x64,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x01, 0x48, 0x02, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x50, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x05, 0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x01, 0x48, 0x03, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x4b, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a,
	0x0c, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x12, 0x0a,
	0x10, 0x5f, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x22, 0xff, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x0a, 0x63, 0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x48, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2c, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x4c, 0x0a, 0x07, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x12, 0x1e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x6d, 0x0a, 0x11, 0x63, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x71,
	0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x43, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x03, 0x48, 0x00, 0x52, 0x10, 0x63, 0x69, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01,
	0x42, 0x14, 0x0a, 0x12, 0x5f, 0x63, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22, 0xcf, 0x01, 0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x1e, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x4e, 0x0a, 0x08, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x08,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x71, 0x63, 0x6c,
	0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x08,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x4a,
	0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x04, 0xe2, 0x41,
	0x01, 0x02, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0xb3, 0x01, 0x0a, 0x19, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xe2, 0x41, 0x01, 0x02, 0xfa, 0x41, 0x27,
	0x0a, 0x25, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x50,
	0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32,
	0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x02, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x22, 0x3d, 0x0a, 0x1a, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32,
	0x8e, 0x04, 0x0a, 0x0e, 0x44, 0x69, 0x73, 0x63, 0x75, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xf8, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3b, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75,
	0x69, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x6a, 0xda, 0x41, 0x34, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2c, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x2c, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2c, 0x63,
	0x61, 0x6e, 0x64, 0x69, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2c, 0x74,
	0x6f, 0x70, 0x5f, 0x70, 0x2c, 0x74, 0x6f, 0x70, 0x5f, 0x6b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d,
	0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x3d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0xdc, 0x01,
	0x0a, 0x12, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x73, 0x12, 0x3e, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3f, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x45, 0xda, 0x41, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2c,
	0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x30, 0x3a, 0x01, 0x2a, 0x22,
	0x2b, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x7b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x3d,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x1a, 0x22, 0xca, 0x41,
	0x1f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x2e, 0x71, 0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x42, 0x56, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71,
	0x63, 0x6c, 0x61, 0x6f, 0x67, 0x75, 0x69, 0x2f, 0x67, 0x61, 0x69, 0x70, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x76, 0x65, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescOnce sync.Once
	file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescData []byte
)

func file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescGZIP() []byte {
	file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescOnce.Do(func() {
		file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDesc), len(file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDesc)))
	})
	return file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDescData
}

var (
	file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
	file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_goTypes  = []any{
		(*GenerateMessageRequest)(nil),     // 0: qclaogui.generativelanguage.v1beta1.GenerateMessageRequest
		(*GenerateMessageResponse)(nil),    // 1: qclaogui.generativelanguage.v1beta1.GenerateMessageResponse
		(*Message)(nil),                    // 2: qclaogui.generativelanguage.v1beta1.Message
		(*MessagePrompt)(nil),              // 3: qclaogui.generativelanguage.v1beta1.MessagePrompt
		(*Example)(nil),                    // 4: qclaogui.generativelanguage.v1beta1.Example
		(*CountMessageTokensRequest)(nil),  // 5: qclaogui.generativelanguage.v1beta1.CountMessageTokensRequest
		(*CountMessageTokensResponse)(nil), // 6: qclaogui.generativelanguage.v1beta1.CountMessageTokensResponse
		(*ContentFilter)(nil),              // 7: qclaogui.generativelanguage.v1beta1.ContentFilter
		(*CitationMetadata)(nil),           // 8: qclaogui.generativelanguage.v1beta1.CitationMetadata
	}
)

var file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_depIdxs = []int32{
	3,  // 0: qclaogui.generativelanguage.v1beta1.GenerateMessageRequest.prompt:type_name -> qclaogui.generativelanguage.v1beta1.MessagePrompt
	2,  // 1: qclaogui.generativelanguage.v1beta1.GenerateMessageResponse.candidates:type_name -> qclaogui.generativelanguage.v1beta1.Message
	2,  // 2: qclaogui.generativelanguage.v1beta1.GenerateMessageResponse.messages:type_name -> qclaogui.generativelanguage.v1beta1.Message
	7,  // 3: qclaogui.generativelanguage.v1beta1.GenerateMessageResponse.filters:type_name -> qclaogui.generativelanguage.v1beta1.ContentFilter
	8,  // 4: qclaogui.generativelanguage.v1beta1.Message.citation_metadata:type_name -> qclaogui.generativelanguage.v1beta1.CitationMetadata
	4,  // 5: qclaogui.generativelanguage.v1beta1.MessagePrompt.examples:type_name -> qclaogui.generativelanguage.v1beta1.Example
	2,  // 6: qclaogui.generativelanguage.v1beta1.MessagePrompt.messages:type_name -> qclaogui.generativelanguage.v1beta1.Message
	2,  // 7: qclaogui.generativelanguage.v1beta1.Example.input:type_name -> qclaogui.generativelanguage.v1beta1.Message
	2,  // 8: qclaogui.generativelanguage.v1beta1.Example.output:type_name -> qclaogui.generativelanguage.v1beta1.Message
	3,  // 9: qclaogui.generativelanguage.v1beta1.CountMessageTokensRequest.prompt:type_name -> qclaogui.generativelanguage.v1beta1.MessagePrompt
	0,  // 10: qclaogui.generativelanguage.v1beta1.DiscussService.GenerateMessage:input_type -> qclaogui.generativelanguage.v1beta1.GenerateMessageRequest
	5,  // 11: qclaogui.generativelanguage.v1beta1.DiscussService.CountMessageTokens:input_type -> qclaogui.generativelanguage.v1beta1.CountMessageTokensRequest
	1,  // 12: qclaogui.generativelanguage.v1beta1.DiscussService.GenerateMessage:output_type -> qclaogui.generativelanguage.v1beta1.GenerateMessageResponse
	6,  // 13: qclaogui.generativelanguage.v1beta1.DiscussService.CountMessageTokens:output_type -> qclaogui.generativelanguage.v1beta1.CountMessageTokensResponse
	12, // [12:14] is the sub-list for method output_type
	10, // [10:12] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_init() }
func file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_init() {
	if File_qclaogui_generativelanguage_v1beta1_discuss_service_proto != nil {
		return
	}
	file_qclaogui_generativelanguage_v1beta1_citation_proto_init()
	file_qclaogui_generativelanguage_v1beta1_safety_proto_init()
	file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[0].OneofWrappers = []any{}
	file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDesc), len(file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_goTypes,
		DependencyIndexes: file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_depIdxs,
		MessageInfos:      file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_msgTypes,
	}.Build()
	File_qclaogui_generativelanguage_v1beta1_discuss_service_proto = out.File
	file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_goTypes = nil
	file_qclaogui_generativelanguage_v1beta1_discuss_service_proto_depIdxs = nil
}
