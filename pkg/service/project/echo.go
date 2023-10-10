// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/golang-api-server/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EchoService interface {
	service.Backend

	projectpb.EchoServiceServer
}

type echoServiceImpl struct {
	projectpb.UnimplementedEchoServiceServer

	nowF func() time.Time
}

func NewEchoService() (EchoService, error) {
	s := &echoServiceImpl{}
	return s, nil
}

func (s *echoServiceImpl) RegisterGRPC(grpcServer *grpc.Server) {
	grpcServer.RegisterService(&projectpb.EchoService_ServiceDesc, s)
}

// Echo This method simply echoes the request.
//
// This method showcases unary RPCs.
func (s *echoServiceImpl) Echo(ctx context.Context, req *projectpb.EchoRequest) (*projectpb.EchoResponse, error) {
	err := status.ErrorProto(req.GetError())
	if err != nil {
		return nil, err
	}

	echoHeaders(ctx)
	echoTrailers(ctx)

	return &projectpb.EchoResponse{Content: req.GetContent(), Severity: req.GetSeverity()}, nil
}

// echo any provided headers in the metadata
func echoHeaders(ctx context.Context) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}

	values := md.Get("x-goog-request-params")
	for _, value := range values {
		header := metadata.Pairs("x-goog-request-params", value)
		_ = grpc.SetHeader(ctx, header)
	}
}

// echo any provided trailing metadata
func echoTrailers(ctx context.Context) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}

	values := md.Get("showcase-trailer")
	for _, value := range values {
		trailer := metadata.Pairs("showcase-trailer", value)
		_ = grpc.SetTrailer(ctx, trailer)
	}
}

// Expand This method splits the given content into words and will pass each word back
// through the stream.
//
// This method showcases server-side streaming RPCs.
func (s *echoServiceImpl) Expand(req *projectpb.ExpandRequest, stream projectpb.EchoService_ExpandServer) error {
	for _, word := range strings.Fields(req.GetContent()) {
		err := stream.Send(&projectpb.EchoResponse{Content: word})
		if err != nil {
			return err
		}

		time.Sleep(req.GetStreamWaitTime().AsDuration())
	}

	echoStreamingHeaders(stream)

	if req.GetError() != nil {
		return status.ErrorProto(req.GetError())
	}

	echoStreamingTrailers(stream)
	return nil
}

// Collect This method will collect the words given to it. When the stream is closed
// by the client, this method will return the a concatenation of the strings
// passed to it.
//
// This method showcases client-side streaming RPCs.
func (s *echoServiceImpl) Collect(stream projectpb.EchoService_CollectServer) error {
	var resp []string

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			echoStreamingHeaders(stream)
			echoStreamingTrailers(stream)
			return stream.SendAndClose(&projectpb.EchoResponse{Content: strings.Join(resp, " ")})
		}
		if err != nil {
			return err
		}

		if err = status.ErrorProto(req.GetError()); err != nil {
			return err
		}

		if req.GetContent() != "" {
			resp = append(resp, req.GetContent())
		}
	}
}

// Chat This method, upon receiving a request on the stream, will pass the same
// content back on the stream.
//
// This method showcases bidirectional streaming RPCs.
func (s *echoServiceImpl) Chat(stream projectpb.EchoService_ChatServer) error {
	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			echoStreamingTrailers(stream)
			return nil
		}
		if err != nil {
			return err
		}

		if err = status.ErrorProto(req.GetError()); err != nil {
			return err
		}

		echoStreamingHeaders(stream)
		_ = stream.Send(&projectpb.EchoResponse{Content: req.GetContent()})
	}
}

func echoStreamingHeaders(stream grpc.ServerStream) {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return
	}

	values := md.Get("x-goog-request-params")
	for _, value := range values {
		header := metadata.Pairs("x-goog-request-params", value)
		if stream.SetHeader(header) != nil {
			return
		}
	}
}

func echoStreamingTrailers(stream grpc.ServerStream) {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return
	}

	values := md.Get("showcase-trailer")
	for _, value := range values {
		trailer := metadata.Pairs("showcase-trailer", value)
		stream.SetTrailer(trailer)
	}
}

// Wait This method will wait for the requested amount of time and then return.
//
// This method showcases how a client handles a request timeout.
func (s *echoServiceImpl) Wait(_ context.Context, req *projectpb.WaitRequest) (*longrunningpb.Operation, error) {
	endTime := time.Unix(0, 0).UTC()
	if ttl := req.GetTtl(); ttl != nil {
		endTime = s.nowF().Add(ttl.AsDuration())
	}
	if end := req.GetEndTime(); end != nil {
		endTime = end.AsTime()
	}
	endTimeProto := timestamppb.New(endTime)
	req.End = &projectpb.WaitRequest_EndTime{
		EndTime: endTimeProto,
	}

	done := s.nowF().After(endTime)
	reqBytes, _ := proto.Marshal(req)

	name := fmt.Sprintf(
		"operations/qclaogui.project.v1.EchoService/Wait/%s",
		base64.StdEncoding.EncodeToString(reqBytes))

	answer := &longrunningpb.Operation{
		Name: name,
		Done: done,
	}

	if done && (req.GetError() != nil) {
		answer.Result = &longrunningpb.Operation_Error{Error: req.GetError()}
	}

	if done && (req.GetSuccess() != nil) {
		resp, _ := anypb.New(req.GetSuccess())
		answer.Result = &longrunningpb.Operation_Response{Response: resp}
	}

	if !done {
		meta, _ := anypb.New(&projectpb.WaitMetadata{EndTime: endTimeProto})
		answer.Metadata = meta
	}

	return answer, nil

}
