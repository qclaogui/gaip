// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package external

import (
	"context"

	"github.com/qclaogui/gaip/genproto/generativelanguage/apiv1/generativelanguagepb"
	"google.golang.org/genai"
)

func NewGenerativeAI(APIKey string) (generativelanguagepb.GenerativeServiceServer, error) {
	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  APIKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, err
	}
	s := &generativeImpl{
		genaiClient: client,
	}
	return s, nil
}

// The generativeImpl type implements a generativelanguagepb.GenerativeServiceServer.
type generativeImpl struct {
	generativelanguagepb.UnimplementedGenerativeServiceServer

	genaiClient *genai.Client
}

//func (g *generativeImpl) GenerateContent(ctx context.Context, request *generativelanguagepb.GenerateContentRequest) (*generativelanguagepb.GenerateContentResponse, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (g *generativeImpl) StreamGenerateContent(request *generativelanguagepb.GenerateContentRequest, server generativelanguagepb.GenerativeService_StreamGenerateContentServer) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (g *generativeImpl) EmbedContent(ctx context.Context, request *generativelanguagepb.EmbedContentRequest) (*generativelanguagepb.EmbedContentResponse, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (g *generativeImpl) BatchEmbedContents(ctx context.Context, request *generativelanguagepb.BatchEmbedContentsRequest) (*generativelanguagepb.BatchEmbedContentsResponse, error) {
//	//TODO implement me
//	panic("implement me")
//}

func (g *generativeImpl) CountTokens(ctx context.Context, req *generativelanguagepb.CountTokensRequest) (*generativelanguagepb.CountTokensResponse, error) {
	// Collect all parts from all contents in the request.
	parts := make([]*genai.Part, 0, len(req.GetContents()))
	for _, content := range req.GetContents() {
		for _, part := range content.GetParts() {
			parts = append(parts, &genai.Part{
				Text: part.GetText(),
			})
		}
	}

	// Call CountTokens once with the aggregated parts.
	resp, err := g.genaiClient.Models.CountTokens(ctx, req.GetModel(), []*genai.Content{{Parts: parts}}, nil)
	if err != nil {
		return nil, err
	}

	return &generativelanguagepb.CountTokensResponse{
		TotalTokens: resp.TotalTokens,
	}, nil
}
