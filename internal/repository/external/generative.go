// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package external

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"github.com/qclaogui/gaip/genproto/generativelanguage/apiv1/generativelanguagepb"
	"google.golang.org/api/option"
)

func NewGenerativeAI(APIKey string) (generativelanguagepb.GenerativeServiceServer, error) {
	//client, err := genai.NewClient(context.Background(), option.WithAPIKey(os.Getenv("API_KEY")))
	client, err := genai.NewClient(context.Background(), option.WithAPIKey(APIKey))
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

func partFromProto(p *generativelanguagepb.Part) genai.Part {
	switch d := p.Data.(type) {
	case *generativelanguagepb.Part_Text:
		return genai.Text(d.Text)
	case *generativelanguagepb.Part_InlineData:
		return genai.Blob{
			MIMEType: d.InlineData.MimeType,
			Data:     d.InlineData.Data,
		}
	default:
		panic(fmt.Errorf("unknown Part.Data type %T", p.Data))
	}
}

func (g *generativeImpl) CountTokens(ctx context.Context, req *generativelanguagepb.CountTokensRequest) (*generativelanguagepb.CountTokensResponse, error) {
	var parts []genai.Part
	for _, content := range req.GetContents() {
		for _, part := range content.GetParts() {
			parts = append(parts, partFromProto(part))
		}
	}

	model := g.genaiClient.GenerativeModel(req.GetModel())
	resp, err := model.CountTokens(ctx, parts...)
	if err != nil {
		return nil, err
	}

	return &generativelanguagepb.CountTokensResponse{
		TotalTokens: resp.TotalTokens,
	}, nil
}
