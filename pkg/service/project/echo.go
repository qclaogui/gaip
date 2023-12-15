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
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EchoService interface {
	pb.EchoServiceServer
}

type echoServiceImpl struct {
	pb.UnimplementedEchoServiceServer

	nowF func() time.Time
}

func NewEchoService() (EchoService, error) {
	s := &echoServiceImpl{}
	return s, nil
}

// Echo This method simply echoes the request.
//
// This method showcases unary RPCs.
func (s *echoServiceImpl) Echo(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	err := status.ErrorProto(req.GetError())
	if err != nil {
		return nil, err
	}

	echoHeaders(ctx)
	echoTrailers(ctx)

	return &pb.EchoResponse{Content: req.GetContent(), Severity: req.GetSeverity()}, nil
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
func (s *echoServiceImpl) Expand(req *pb.ExpandRequest, stream pb.EchoService_ExpandServer) error {
	for _, word := range strings.Fields(req.GetContent()) {
		err := stream.Send(&pb.EchoResponse{Content: word})
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
func (s *echoServiceImpl) Collect(stream pb.EchoService_CollectServer) error {
	var resp []string

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			echoStreamingHeaders(stream)
			echoStreamingTrailers(stream)
			return stream.SendAndClose(&pb.EchoResponse{Content: strings.Join(resp, " ")})
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
func (s *echoServiceImpl) Chat(stream pb.EchoService_ChatServer) error {
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
		_ = stream.Send(&pb.EchoResponse{Content: req.GetContent()})
	}
}

// PagedExpand This is similar to the Expand method but instead of returning a stream of
// expanded words, this method returns a paged list of expanded words.
func (s *echoServiceImpl) PagedExpand(ctx context.Context, req *pb.PagedExpandRequest) (*pb.PagedExpandResponse, error) {
	words := strings.Fields(req.GetContent())

	start, end, nextToken, err := processPageTokens(len(words), req.GetPageSize(), req.GetPageToken())
	if err != nil {
		return nil, err
	}

	var responses []*pb.EchoResponse
	for _, word := range words[start:end] {
		responses = append(responses, &pb.EchoResponse{Content: word})
	}

	echoHeaders(ctx)
	echoTrailers(ctx)

	return &pb.PagedExpandResponse{
		Responses:     responses,
		NextPageToken: nextToken,
	}, nil
}

func processPageTokens(numElements int, pageSize int32, pageToken string) (start, end int32, nextToken string, err error) {
	if pageSize < 0 {
		return 0, 0, "", status.Error(codes.InvalidArgument, "the page size provided must not be negative.")
	}

	if pageToken != "" {
		token, err := strconv.Atoi(pageToken)
		token32 := int32(token)

		if err != nil || token32 < 0 || token32 >= int32(numElements) {
			return 0, 0, "", status.Errorf(
				codes.InvalidArgument,
				"invalid page token: %s. Token must be within the range [0, %d)",
				pageToken,
				numElements)
		}
		start = token32
	}

	if pageSize == 0 {
		pageSize = int32(numElements)
	}

	end = min(start+pageSize, int32(numElements))

	if end < int32(numElements) {
		nextToken = strconv.Itoa(int(end))
	}

	return start, end, nextToken, nil
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

// Block This method will block (wait) for the requested amount of time
// and then return the response or error.
//
// This method showcases how a client handles delays or retries.
func (s *echoServiceImpl) Block(ctx context.Context, req *pb.BlockRequest) (*pb.BlockResponse, error) {
	d := req.GetResponseDelay().AsDuration()
	time.Sleep(d)
	if req.GetError() != nil {
		return nil, status.ErrorProto(req.GetError())
	}

	echoHeaders(ctx)
	echoTrailers(ctx)
	return req.GetSuccess(), nil
}

// Wait This method will wait for the requested amount of time and then return.
//
// This method showcases how a client handles a request timeout.
func (s *echoServiceImpl) Wait(_ context.Context, req *pb.WaitRequest) (*longrunningpb.Operation, error) {
	endTime := time.Unix(0, 0).UTC()
	if ttl := req.GetTtl(); ttl != nil {
		endTime = s.nowF().Add(ttl.AsDuration())
	}
	if end := req.GetEndTime(); end != nil {
		endTime = end.AsTime()
	}
	endTimeProto := timestamppb.New(endTime)
	req.End = &pb.WaitRequest_EndTime{
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
		meta, _ := anypb.New(&pb.WaitMetadata{EndTime: endTimeProto})
		answer.Metadata = meta
	}

	return answer, nil

}
