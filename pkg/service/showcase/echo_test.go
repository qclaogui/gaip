// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package showcase

import (
	"context"
	"encoding/base64"
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"github.com/stretchr/testify/require"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// mockSTS grpc.ServerTransportStream
type mockSTS struct {
	stream grpc.ServerStream
	t      *testing.T
}

func (m *mockSTS) Method() string                  { return "" }
func (m *mockSTS) SetHeader(md metadata.MD) error  { _ = m.stream.SetHeader(md); return nil }
func (m *mockSTS) SendHeader(md metadata.MD) error { return m.stream.SendHeader(md) }
func (m *mockSTS) SetTrailer(md metadata.MD) error { m.stream.SetTrailer(md); return nil }

type mockUnaryStream struct {
	head  []string
	trail []string
	t     *testing.T
	grpc.ServerStream
}

func (m *mockUnaryStream) Method() string                { return "" }
func (m *mockUnaryStream) Send(_ *pb.EchoResponse) error { return nil }
func (m *mockUnaryStream) Context() context.Context      { return nil }
func (m *mockUnaryStream) SetTrailer(md metadata.MD) {
	m.trail = append(m.trail, md.Get("showcase-trailer")...)
}

func (m *mockUnaryStream) SetHeader(md metadata.MD) error {
	m.head = append(m.head, md.Get("x-goog-request-params")...)
	return nil
}

func (m *mockUnaryStream) verify(expectHeadersAndTrailers bool) {
	if expectHeadersAndTrailers && !reflect.DeepEqual([]string{"show", "case"}, m.trail) && !reflect.DeepEqual([]string{"showcaseHeader, anotherHeader"}, m.head) {
		m.t.Errorf("Unary stream did not get all expected headers and trailers.\nGot these headers: %+v\nGot these trailers: %+v", m.head, m.trail)
	}
}

type mockExpandStream struct {
	exp   []string
	head  []string
	trail []string

	t *testing.T

	pb.EchoService_ExpandServer
}

func (m *mockExpandStream) Send(resp *pb.EchoResponse) error {
	if resp.GetContent() != m.exp[0] {
		m.t.Errorf("Expand expected to send %s but sent %s", m.exp[0], resp.GetContent())
	}

	m.exp = m.exp[1:]
	return nil
}

func (m *mockExpandStream) Context() context.Context {
	return appendTestOutgoingMetadata(context.Background(), &mockSTS{stream: m, t: m.t})
}

func (m *mockExpandStream) SetTrailer(md metadata.MD) {
	m.trail = append(m.trail, md.Get("showcase-trailer")...)
}

func (m *mockExpandStream) SetHeader(md metadata.MD) error {
	m.head = append(m.head, md.Get("x-goog-request-params")...)
	return nil
}

func (m *mockExpandStream) verify(expectHeadersAndTrailers bool) {
	if len(m.exp) > 0 {
		m.t.Errorf("Expand did not stream all expected values. %d expected values remaining.", len(m.exp))
	}

	if expectHeadersAndTrailers && !reflect.DeepEqual([]string{"show", "case"}, m.trail) && !reflect.DeepEqual([]string{"showcaseHeader", "anotherHeader"}, m.head) {
		m.t.Errorf("Expand did not get all expected headers and trailers.\nGot these headers: %+v\nGot these trailers: %+v", m.head, m.trail)
	}
}

func serverSetup(t *testing.T) pb.EchoServiceServer {
	//s, err := NewEchoService()
	//require.NoError(t, err)

	cfg := Config{}
	s, err := NewServer(cfg)
	require.NoError(t, err)
	return s
}

func TestExpand(t *testing.T) {
	contentTable := []string{"Hello World", "Hola", ""}
	errTable := []*spb.Status{
		{Code: int32(codes.OK)},
		{Code: int32(codes.InvalidArgument)},
		nil,
	}

	s := serverSetup(t)
	for _, c := range contentTable {
		for _, e := range errTable {
			stream := &mockExpandStream{exp: strings.Fields(c), t: t}
			err := s.Expand(&pb.ExpandRequest{Content: c, Error: e}, stream)
			sts, _ := status.FromError(err)
			if int32(sts.Code()) != e.GetCode() {
				t.Errorf("Expand expected stream to return status with code %d but code %d", sts.Code(), e.GetCode())
			}
			stream.verify(e == nil)
		}
	}

}

func TestExpandWithWaitTime(t *testing.T) {
	s := serverSetup(t)
	//This stream should take at least 300ms to complete because there are 7 messages, and we wait 50ms between sending each message.
	content := "This stream should take 300ms to complete"
	stream := &mockExpandStream{exp: strings.Fields(content), t: t}
	streamWaitTime := durationpb.New(time.Duration(50) * time.Millisecond)
	start := time.Now()

	err := s.Expand(&pb.ExpandRequest{Content: content, StreamWaitTime: streamWaitTime}, stream)

	actualTimeSpent := int(time.Since(start).Milliseconds())
	expectedTimeSpent := 300
	if actualTimeSpent < expectedTimeSpent {
		t.Errorf("Expand stream should take at least %d ms to complete, but it only took %d ms", expectedTimeSpent, actualTimeSpent)
	}

	stream.verify(err == nil)
}

type errorExpandStream struct {
	err error
	pb.EchoService_ExpandServer
}

func (s *errorExpandStream) Send(_ *pb.EchoResponse) error {
	return s.err
}

func TestExpand_streamErr(t *testing.T) {
	e := errors.New("test Error")

	stream := &errorExpandStream{err: e}

	s := serverSetup(t)
	err := s.Expand(&pb.ExpandRequest{Content: "Hello World"}, stream)
	if !errors.Is(err, e) {
		t.Error("Expand expected to pass through stream errors.")
	}
}

type mockCollectStream struct {
	requests []*pb.EchoRequest
	head     []string
	trail    []string
	exp      *string

	t *testing.T
	pb.EchoService_CollectServer
}

func (m *mockCollectStream) SendAndClose(resp *pb.EchoResponse) error {
	if m.exp == nil {
		m.t.Errorf("Collect Stream SendAndClose called unexpectedly")
	}

	if resp.GetContent() != *m.exp {
		m.t.Errorf("Collect expected to return '%s', but returned '%s'", *m.exp, resp.GetContent())
	}
	return nil
}

func (m *mockCollectStream) Recv() (*pb.EchoRequest, error) {
	if len(m.requests) > 0 {
		req := m.requests[0]
		m.requests = m.requests[1:]
		return req, nil
	}
	return nil, io.EOF
}
func (m *mockCollectStream) Context() context.Context {
	return appendTestOutgoingMetadata(context.Background(), &mockSTS{stream: m, t: m.t})
}

func (m *mockCollectStream) SetHeader(md metadata.MD) error {
	m.head = append(m.head, md.Get("x-goog-request-params")...)
	return nil
}

func (m *mockCollectStream) SetTrailer(md metadata.MD) {
	m.trail = append(m.trail, md.Get("showcase-trailer")...)
}

func (m *mockCollectStream) verify(expectHeadersAndTrailers bool) {
	if expectHeadersAndTrailers && !reflect.DeepEqual([]string{"show", "case"}, m.trail) && !reflect.DeepEqual([]string{"showcaseHeader", "anotherHeader"}, m.head) {
		m.t.Errorf("Collect did not get all expected trailers.\nGot these headers: %+v\nGot these trailers: %+v", m.head, m.trail)
	}
}

func TestCollect(t *testing.T) {
	strPtr := func(s string) *string { return &s }

	tests := []struct {
		contents []string
		exp      *string
		err      *spb.Status
	}{
		{[]string{"Hello", "", "World"}, strPtr("Hello World"), nil},
	}

	s := serverSetup(t)
	for _, test := range tests {
		var requests []*pb.EchoRequest
		for _, content := range test.contents {
			requests = append(requests, &pb.EchoRequest{Response: &pb.EchoRequest_Content{Content: content}})
		}

		if test.err != nil {
			requests = append(requests, &pb.EchoRequest{Response: &pb.EchoRequest_Error{Error: test.err}})
		}

		stream := &mockCollectStream{requests: requests, exp: test.exp, t: t}

		err := s.Collect(stream)
		expCode := status.FromProto(test.err).Code()
		sts, _ := status.FromError(err)
		if sts.Code() != expCode {
			t.Errorf("Collect expected to return with code %d, but returned %d", expCode, sts.Code())
		}

		stream.verify(test.err == nil)
	}
}

type errorCollectStream struct {
	err error
	pb.EchoService_CollectServer
}

func (s *errorCollectStream) Recv() (*pb.EchoRequest, error) {
	return nil, s.err
}

func TestCollect_streamErr(t *testing.T) {
	e := errors.New("test Error")
	stream := &errorCollectStream{err: e}

	s := serverSetup(t)
	err := s.Collect(stream)
	if !errors.Is(err, e) {
		t.Error("Collect expected to pass through stream errors.")
	}
}

type mockChatStream struct {
	pb.EchoService_ChatServer
	t *testing.T

	requests []*pb.EchoRequest
	head     []string
	trail    []string
	curr     *pb.EchoRequest
}

func (m *mockChatStream) Recv() (*pb.EchoRequest, error) {
	if len(m.requests) > 0 {
		m.curr = m.requests[0]
		m.requests = m.requests[1:]
		return m.curr, nil
	}
	return nil, io.EOF
}

func (m *mockChatStream) Send(resp *pb.EchoResponse) error {
	if m.curr == nil {
		m.t.Errorf("Chat unexpectedly tried to send content.")
	}

	if resp.GetContent() != m.curr.GetContent() {
		m.t.Errorf("Chat expected to send content %s, but sent %s", m.curr.GetContent(), resp.GetContent())
		m.curr = nil
	}
	return nil
}

func (m *mockChatStream) Context() context.Context {
	return appendTestOutgoingMetadata(context.Background(), &mockSTS{stream: m, t: m.t})
}

func (m *mockChatStream) SetHeader(md metadata.MD) error {
	m.head = append(m.head, md.Get("x-goog-request-params")...)
	return nil
}

func (m *mockChatStream) SetTrailer(md metadata.MD) {
	m.trail = append(m.trail, md.Get("showcase-trailer")...)
}

func (m *mockChatStream) verify(expectHeadersAndTrailers bool) {
	if expectHeadersAndTrailers && !reflect.DeepEqual([]string{"show", "case"}, m.trail) && !reflect.DeepEqual([]string{"showcaseHeader", "anotherHeader"}, m.head) {
		m.t.Errorf("Chat did not get all expected trailers.\nGot these headers: %+v\nGot these trailers: %+v", m.head, m.trail)
	}
}

func TestChat(t *testing.T) {
	tests := []struct {
		contents []string
		err      *spb.Status
	}{
		{[]string{"Hello", "World"}, nil},
		{[]string{"Hello", "World"}, &spb.Status{Code: int32(codes.InvalidArgument)}},
		{[]string{}, &spb.Status{Code: int32(codes.InvalidArgument)}},
	}

	s := serverSetup(t)

	for _, test := range tests {
		var requests []*pb.EchoRequest

		for _, content := range test.contents {
			requests = append(requests, &pb.EchoRequest{Response: &pb.EchoRequest_Content{Content: content}})
		}

		if test.err != nil {
			requests = append(requests, &pb.EchoRequest{Response: &pb.EchoRequest_Error{Error: test.err}})
		}

		stream := &mockChatStream{requests: requests, t: t}
		err := s.Chat(stream)
		sts, _ := status.FromError(err)

		expCode := status.FromProto(test.err).Code()
		if sts.Code() != expCode {
			t.Errorf("Chat expected to return status with code %d, but returned %d", expCode, sts.Code())
		}

		stream.verify(test.err == nil)
	}
}

type errorChatStream struct {
	err error
	pb.EchoService_ChatServer
}

func (s *errorChatStream) Recv() (*pb.EchoRequest, error) {
	return nil, s.err
}

func TestChat_streamErr(t *testing.T) {
	e := errors.New("test Error")
	stream := &errorChatStream{err: e}

	s := serverSetup(t)
	err := s.Chat(stream)
	if !errors.Is(err, e) {
		t.Error("Chat expected to pass through stream errors.")
	}
}

func appendTestOutgoingMetadata(ctx context.Context, stream grpc.ServerTransportStream) context.Context {
	ctx = grpc.NewContextWithServerTransportStream(ctx, stream)
	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("showcase-trailer", "show", "showcase-trailer", "case", "trailer", "trail", "x-goog-request-params", "showcaseHeader", "x-goog-request-params", "anotherHeader", "header", "head"))
	return ctx
}

func TestPagedExpand_invalidArgs(t *testing.T) {
	tests := []*pb.PagedExpandRequest{
		{PageSize: -1},
		{PageToken: "-1"},
		{PageToken: "BOGUS"},
		{Content: "one", PageToken: "1"},
		{Content: "one", PageToken: "2"},
	}

	s := serverSetup(t)
	for _, test := range tests {
		_, err := s.PagedExpand(context.Background(), test)
		sts, _ := status.FromError(err)
		if sts.Code() != codes.InvalidArgument {
			t.Errorf("PagedExpand() expected error code: %d, got error code %d",
				codes.InvalidArgument, sts.Code())
		}
	}
}
func TestPagedExpand(t *testing.T) {
	tests := []struct {
		req  *pb.PagedExpandRequest
		resp *pb.PagedExpandResponse
	}{
		{
			req: &pb.PagedExpandRequest{Content: "Hello world!"},
			resp: &pb.PagedExpandResponse{
				Responses: []*pb.EchoResponse{
					{Content: "Hello"},
					{Content: "world!"},
				},
			},
		},
		{
			req: &pb.PagedExpandRequest{PageSize: 3, Content: "Hello world!"},
			resp: &pb.PagedExpandResponse{
				Responses: []*pb.EchoResponse{
					{Content: "Hello"},
					{Content: "world!"},
				},
			},
		},
		{
			req: &pb.PagedExpandRequest{PageSize: 3, Content: "The rain in Spain falls mainly on the plain!"},
			resp: &pb.PagedExpandResponse{
				Responses: []*pb.EchoResponse{
					{Content: "The"},
					{Content: "rain"},
					{Content: "in"},
				},
				NextPageToken: "3",
			},
		}, {
			req: &pb.PagedExpandRequest{
				PageSize:  3,
				PageToken: "3",
				Content:   "The rain in Spain falls mainly on the plain!",
			},
			resp: &pb.PagedExpandResponse{
				Responses: []*pb.EchoResponse{
					{Content: "Spain"},
					{Content: "falls"},
					{Content: "mainly"},
				},
				NextPageToken: "6",
			},
		},
	}

	s := serverSetup(t)
	for _, test := range tests {
		stream := &mockUnaryStream{t: t}
		ctx := appendTestOutgoingMetadata(context.Background(), &mockSTS{t: t, stream: stream})
		resp, err := s.PagedExpand(ctx, test.req)
		if err != nil {
			t.Error(err)
		}

		if !proto.Equal(test.resp, resp) {
			t.Errorf("PagedExpand with input %q, expected: %q, got: %q",
				test.req.String(), test.resp.String(), resp.String())
		}

		stream.verify(err == nil)
	}

}

func TestBlockSuccess(t *testing.T) {
	tests := []struct {
		seconds int64
		nanos   int32
		resp    string
	}{
		{1, int32(1000), "hello"},
		{5, int32(10), "world"},
	}
	nowF := func() time.Time { return time.Unix(1, 0) }

	for _, test := range tests {
		waiter := &Server{nowF: nowF}

		request := &pb.BlockRequest{
			ResponseDelay: &durationpb.Duration{
				Seconds: test.seconds,
				Nanos:   test.nanos,
			},
			Response: &pb.BlockRequest_Success{
				Success: &pb.BlockResponse{Content: test.resp},
			},
		}

		stream := &mockUnaryStream{t: t}
		ctx := appendTestOutgoingMetadata(context.Background(), &mockSTS{t: t, stream: stream})
		out, err := waiter.Block(ctx, request)
		if err != nil {
			t.Error(err)
		}

		if out.GetContent() != test.resp {
			t.Errorf("Expected Wait test to return %s, but returned %s", test.resp, out.GetContent())
		}
		stream.verify(err == nil)
	}
}

func TestWait_pending(t *testing.T) {
	now := time.Unix(1, 0)
	endTime := time.Unix(2, 0)

	ttl := endTime.Sub(now)
	nowF := func() time.Time { return time.Unix(1, 0) }

	endTimeProto := timestamppb.New(endTime)

	tests := []*pb.WaitRequest{
		{
			End: &pb.WaitRequest_EndTime{
				EndTime: endTimeProto,
			},
		},
		{
			End: &pb.WaitRequest_Ttl{
				Ttl: durationpb.New(ttl),
			},
		},
	}

	for _, req := range tests {
		waiter := &Server{nowF: nowF}

		op, _ := waiter.Wait(context.Background(), req)
		if op.Done {
			t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
				"✘got: %v\n\x1b[92m"+
				"want: %v\x1b[39m", "done=false", "done=true")
		}

		checkName(t, req, op)

		if op.Metadata == nil {
			t.Errorf("Wait() for %q expected metadata, got nil", req)
		}

		meta := &pb.WaitMetadata{}
		_ = anypb.UnmarshalTo(op.Metadata, meta, proto.UnmarshalOptions{})
		if !proto.Equal(endTimeProto, meta.EndTime) {
			t.Errorf(
				"Wait for %q expected metadata with Endtime=%q, got %q",
				req,
				endTimeProto,
				meta.EndTime)
		}

	}
}

func checkName(t *testing.T, req *pb.WaitRequest, op *longrunningpb.Operation) {
	if !strings.HasPrefix(op.Name, "operations/qclaogui.showcase.v1beta1.EchoService/Wait/") {
		t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
			"✘got: %v\n\x1b[92m"+
			"want: %v\x1b[39m", op.Name, "op.Name prefex 'operations/qclaogui.showcase.v1beta1.EchoService/Wait/'")
	}

	nameProto := &pb.WaitRequest{}

	encodedBytes := strings.TrimPrefix(
		op.Name,
		"operations/qclaogui.showcase.v1beta1.EchoService/Wait/")
	bytes, _ := base64.StdEncoding.DecodeString(encodedBytes)

	_ = proto.Unmarshal(bytes, nameProto)
	if !proto.Equal(nameProto, req) {
		t.Errorf(
			"Wait() for %q expected unmarshalled name=%q, got name=%q",
			req,
			req,
			nameProto)
	}
}

func TestWait_success(t *testing.T) {
	nowF := func() time.Time { return time.Unix(3, 0) }
	endTime := timestamppb.New(time.Unix(2, 0))
	success := &pb.WaitResponse{Content: "Hello World!"}
	req := &pb.WaitRequest{
		End:      &pb.WaitRequest_EndTime{EndTime: endTime},
		Response: &pb.WaitRequest_Success{Success: success},
	}

	waiter := &Server{nowF: nowF}
	op, _ := waiter.Wait(context.Background(), req)

	checkName(t, req, op)

	if !op.Done {
		t.Errorf("Wait() for %q expected done=true got done=false", req)
	}

	if op.Metadata != nil {
		t.Errorf("Wait() for %q expected nil metadata, got %q", req, op.Metadata)
	}

	if op.GetError() != nil {
		t.Errorf("Wait() expected op.Error=nil, got %q", op.GetError())
	}

	if op.GetResponse() == nil {
		t.Error("Wait() expected op.Response!=nil")
	}

	resp := &pb.WaitResponse{}
	_ = anypb.UnmarshalTo(op.GetResponse(), resp, proto.UnmarshalOptions{})
	if !proto.Equal(resp, success) {
		t.Errorf("Wait() expected op.GetResponse()=%q, got %q", success, resp)
	}
}

func TestWait_error(t *testing.T) {
	nowF := func() time.Time { return time.Unix(3, 0) }
	endTime := timestamppb.New(time.Unix(2, 0))
	expErr := &spb.Status{Code: int32(1), Message: "Error!"}
	req := &pb.WaitRequest{
		End: &pb.WaitRequest_EndTime{
			EndTime: endTime,
		},
		Response: &pb.WaitRequest_Error{Error: expErr},
	}

	waiter := &Server{nowF: nowF}
	op, _ := waiter.Wait(context.Background(), req)

	checkName(t, req, op)

	if !op.Done {
		t.Errorf("Wait() for %q expected done=true got done=false", req)
	}

	if op.Metadata != nil {
		t.Errorf("Wait() for %q expected nil metadata, got %q", req, op.Metadata)
	}

	if op.GetResponse() != nil {
		t.Errorf("Wait() expected op.Response=nil, got %q", op.GetResponse())
	}

	if !proto.Equal(expErr, op.GetError()) {
		t.Errorf("Wait() expected op.Error=%q, got %q", expErr, op.GetError())
	}
}
