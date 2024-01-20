// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package generativeai

import (
	"context"

	"github.com/go-kit/log/level"
	pb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1/generativelanguagepb"
)

func (s *Server) GenerateContent(ctx context.Context, req *pb.GenerateContentRequest) (*pb.GenerateContentResponse, error) {
	return s.repoGenerative.GenerateContent(ctx, req)
}

func (s *Server) StreamGenerateContent(req *pb.GenerateContentRequest, stream pb.GenerativeService_StreamGenerateContentServer) error {
	return s.repoGenerative.StreamGenerateContent(req, stream)
}

func (s *Server) EmbedContent(ctx context.Context, req *pb.EmbedContentRequest) (*pb.EmbedContentResponse, error) {
	return s.repoGenerative.EmbedContent(ctx, req)
}

func (s *Server) BatchEmbedContents(ctx context.Context, req *pb.BatchEmbedContentsRequest) (*pb.BatchEmbedContentsResponse, error) {
	return s.repoGenerative.BatchEmbedContents(ctx, req)
}

func (s *Server) CountTokens(ctx context.Context, req *pb.CountTokensRequest) (*pb.CountTokensResponse, error) {
	_ = level.Info(s.logger).Log("msg", "[CountTokens] received request")
	defer func() { _ = level.Info(s.logger).Log("msg", "[CountTokens] completed request") }()

	return s.repoGenerative.CountTokens(ctx, req)
}
