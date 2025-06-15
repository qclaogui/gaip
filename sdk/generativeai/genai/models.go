// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package genai

import (
	"context"

	pb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta/generativelanguagepb"
)

type Models struct {
	apiClient *apiClient

	SafetySettings []*SafetySetting
	Tools          []*Tool
	ToolConfig     *ToolConfig

	GenerationConfig *GenerationConfig
	// SystemInstruction (also known as "system prompt") is a more forceful prompt to the model.
	// The model will adhere the instructions more strongly than if they appeared in a normal prompt.
	SystemInstruction *Content
	// The name of the CachedContent to use.
	// Must have already been created with [Client.CreateCachedContent].
	CachedContentName string
}

// GenerateContent generates content based on the provided model, contents, and configuration.
func (m Models) GenerateContent(ctx context.Context, model string, contents []*Content, config *GenerationConfig) (*GenerateContentResponse, error) {
	if config != nil {
		config.setDefaults()
	}
	return m.generateContent(ctx, model, contents, config)
}

func (m Models) generateContent(ctx context.Context, model string, contents []*Content, config *GenerationConfig) (*GenerateContentResponse, error) {
	req, err := m.newGenerateContentRequest(model, contents, config)
	if err != nil {
		return nil, err
	}
	res, err := m.apiClient.gc.GenerateContent(ctx, req)
	if err != nil {
		return nil, err
	}
	return protoToResponse(res)
}

func (m Models) newGenerateContentRequest(model string, contents []*Content, config *GenerationConfig) (*pb.GenerateContentRequest, error) {
	return pvCatchPanic(func() *pb.GenerateContentRequest {
		var cc *string
		if m.CachedContentName != "" {
			cc = &m.CachedContentName
		}

		req := &pb.GenerateContentRequest{
			Model:             model,
			Contents:          transformSlice(contents, (*Content).toProto),
			SafetySettings:    transformSlice(m.SafetySettings, (*SafetySetting).toProto),
			Tools:             transformSlice(m.Tools, (*Tool).toProto),
			ToolConfig:        m.ToolConfig.toProto(),
			GenerationConfig:  config.toProto(),
			SystemInstruction: m.SystemInstruction.toProto(),
			CachedContent:     cc,
		}
		debugPrint(req)
		return req
	})
}
