// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package showcase_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/googleapis/gax-go/v2"
	showcase "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1"
	pb "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Clients are initialized in main_test.go.
var (
	echoGRPC *showcase.EchoClient
	echoREST *showcase.EchoClient
)

func TestEcho(t *testing.T) {
	content := "hello world!"
	req := &pb.EchoRequest{
		Response: &pb.EchoRequest_Content{
			Content: content,
		},
	}

	for typ, client := range map[string]*showcase.EchoClient{
		"grpc": echoGRPC,
		"rest": echoREST,
	} {
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.GetContent() != req.GetContent() {
			t.Errorf("%s Echo() = %q, want %q", typ, resp.GetContent(), content)
		}
	}

}

func TestEcho_error(t *testing.T) {
	val := codes.Canceled
	req := &pb.EchoRequest{
		Response: &pb.EchoRequest_Error{
			Error: &spb.Status{Code: int32(val)},
		},
	}

	for typ, client := range map[string]*showcase.EchoClient{
		"grpc": echoGRPC,
		"rest": echoREST,
	} {
		resp, err := client.Echo(context.Background(), req)
		if resp != nil || err == nil {
			t.Errorf("%s Echo() = %v, wanted error %d", typ, resp, val)
		}

		if typ == "grpc" {
			s, _ := status.FromError(err)
			if s.Code() != val {
				t.Errorf("%s Echo() errors with %d, want %d", typ, s.Code(), val)
			}
		} else {
			want := 499
			gerr := &googleapi.Error{}
			if !errors.As(err, &gerr) {
				t.Errorf("%s Echo() returned unexpected error type: %v", typ, err)
			} else if gerr.Code != want {
				t.Errorf("%s Echo() errors with %d, want %d", typ, gerr.Code, want)
			}
		}
	}

}

// Test dynamic routing header generation. We cannot guarantee the order that headers are sent, so we check that the header sent contains the correct elements as opposed to checking
// the header itself.
func TestEchoHeader(t *testing.T) {
	var testCases = []struct {
		req  *pb.EchoRequest
		want []string
	}{
		{
			req:  &pb.EchoRequest{OtherHeader: "projects/123/instances/456"},
			want: []string{"baz=projects%2F123%2Finstances%2F456", "qux=projects%2F123"},
		},
		{
			req:  &pb.EchoRequest{OtherHeader: "instances/456"},
			want: []string{"baz=instances%2F456"},
		},
		{
			req:  &pb.EchoRequest{Header: "potato"},
			want: []string{"header=potato", "routing_id=potato"},
		},
		{
			req:  &pb.EchoRequest{Header: "projects/123/instances/456"},
			want: []string{"header=projects%2F123%2Finstances%2F456", "routing_id=projects%2F123%2Finstances%2F456", "super_id=projects%2F123", "table_name=projects%2F123%2Finstances%2F456", "instance_id=instances%2F456"},
		},
		{
			req: &pb.EchoRequest{
				Header:      "regions/123/zones/456",
				OtherHeader: "projects/123/instances/456",
			},
			want: []string{"baz=projects%2F123%2Finstances%2F456", "qux=projects%2F123", "table_name=regions%2F123%2Fzones%2F456", "header=regions%2F123%2Fzones%2F456", "routing_id=regions%2F123%2Fzones%2F456"},
		},
	}

	for _, tc := range testCases {
		mdForHeaders := metadata.New(map[string]string{})
		_, _ = echoGRPC.Echo(context.Background(), tc.req, gax.WithGRPCOptions(grpc.Header(&mdForHeaders)))
		got := mdForHeaders.Get("x-goog-request-params")
		got = strings.Split(got[0], "&")
		sort.Strings(got)
		sort.Strings(tc.want)

		if diff := cmp.Diff(got, tc.want); diff != "" {
			t.Errorf("got(-),want(+):\n%s", diff)
		}
	}
}

func TestEchoHeaderREST(t *testing.T) {
	var testCases = []struct {
		req  *pb.EchoRequest
		want []string
	}{
		{
			req:  &pb.EchoRequest{OtherHeader: "projects/123/instances/456"},
			want: []string{"baz=projects%2F123%2Finstances%2F456", "qux=projects%2F123"},
		},
		{
			req:  &pb.EchoRequest{OtherHeader: "instances/456"},
			want: []string{"baz=instances%2F456"},
		},
		{
			req:  &pb.EchoRequest{Header: "potato"},
			want: []string{"header=potato", "routing_id=potato"},
		},
		{
			req:  &pb.EchoRequest{Header: "projects/123/instances/456"},
			want: []string{"header=projects%2F123%2Finstances%2F456", "routing_id=projects%2F123%2Finstances%2F456", "super_id=projects%2F123", "table_name=projects%2F123%2Finstances%2F456", "instance_id=instances%2F456"},
		},
		{
			req: &pb.EchoRequest{
				Header:      "regions/123/zones/456",
				OtherHeader: "projects/123/instances/456",
			},
			want: []string{"baz=projects%2F123%2Finstances%2F456", "qux=projects%2F123", "table_name=regions%2F123%2Fzones%2F456", "header=regions%2F123%2Fzones%2F456", "routing_id=regions%2F123%2Fzones%2F456"},
		},
	}

	for _, tc := range testCases {
		// Wrap the default RoundTripper with our own that asserts on the response
		// headers expected by the test.
		wrapped := &http.Client{}
		wrapped.Transport = headerChecker{rt: wrapped.Transport, want: tc.want, t: t}
		echoWrapped, err := showcase.NewEchoRESTClient(
			context.Background(),
			option.WithEndpoint("http://localhost:7469"),
			option.WithoutAuthentication(),
		)
		if err != nil {
			t.Fatal(err)
		}

		_, _ = echoWrapped.Echo(context.Background(), tc.req)
		_ = echoWrapped.Close()
	}
}

func TestXGoogHeaders(t *testing.T) {
	// Inspect the private property `xGoogHeaders` of the transport-specific
	// client implementation that is populated on creation of the client.
	w := reflect.ValueOf(*echoGRPC)
	x := w.FieldByName("internalClient")
	y := x.Elem().Elem()
	info := y.FieldByName("xGoogHeaders")

	var goVersion string
	vals := make([]string, 0)
	for i := 0; i < info.Len(); i++ {
		key := info.Index(i)
		// Only check for the client info set by the generated setGoogleClientInfo()
		if key.String() != "x-goog-api-client" {
			continue
		}

		vals = append(vals, info.Index(i+1).String())
	}

	for i := 0; goVersion == "" || i < len(vals); i++ {
		v := vals[i]
		split := strings.Split(v, " ")
		for _, s := range split {
			// For now, we only want to check that the Go version is being
			// properly populated.
			if strings.HasPrefix(s, "gl-go/") {
				goVersion = s
				break
			}
		}
	}

	if goVersion == "" {
		t.Errorf("expected Go version pair to be populated, but wasn't: %v", info)
	} else if strings.Contains(goVersion, "UNKNOWN") {
		t.Errorf("expected Go version pair to not be UNKNOWN: %q", goVersion)
	}
}

type headerChecker struct {
	rt   http.RoundTripper
	want []string
	t    *testing.T
}

func (hc headerChecker) RoundTrip(r *http.Request) (*http.Response, error) {
	resp, err := hc.rt.RoundTrip(r)

	header := resp.Header
	got := header[http.CanonicalHeaderKey("x-goog-request-params")]
	got = strings.Split(got[0], "&")
	sort.Strings(got)
	sort.Strings(hc.want)
	if diff := cmp.Diff(got, hc.want); diff != "" {
		hc.t.Errorf("got(-),want(+):\n%s", diff)
	}
	return resp, err
}

// Chat, Collect, and Expand are streaming methods and don't have interesting REST semantics
func TestExpand(t *testing.T) {
	content := "The rain in Spain stays mainly on the plain!"
	req := &pb.ExpandRequest{
		Content: content,
	}

	for typ, client := range map[string]*showcase.EchoClient{
		"grpc": echoGRPC,
		"rest": echoREST,
	} {
		s, err := client.Expand(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}

		var rests []string
		for {
			resp, err2 := s.Recv()
			if errors.Is(err2, io.EOF) {
				break
			}
			if err2 != nil {
				t.Fatal(err2)
			}
			rests = append(rests, resp.GetContent())
		}

		got := strings.Join(rests, " ")
		if content != got {
			t.Errorf("%s Expand() = %q, want %q", typ, got, content)
		}
	}
}

// Chat, Collect, and Expand are streaming methods and don't have interesting REST semantics
func TestCollect(t *testing.T) {
	content := "The rain in Spain stays mainly on the plain!"
	s, err := echoGRPC.Collect(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	for _, str := range strings.Split(content, " ") {
		req := &pb.EchoRequest{
			Response: &pb.EchoRequest_Content{
				Content: str,
			},
		}
		if err = s.Send(req); err != nil {
			t.Fatal(err)
		}
	}

	resp, err := s.CloseAndRecv()
	if err != nil {
		t.Fatal(err)
	}

	if content != resp.GetContent() {
		t.Errorf("Collect() = %q, want %q", resp.GetContent(), content)
	}
}

// Chat, Collect, and Expand are streaming methods and don't have interesting REST semantics
func TestChat(t *testing.T) {
	content := "The rain in Spain stays mainly on the plain!"
	s, err := echoGRPC.Chat(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	for _, str := range strings.Split(content, " ") {
		req := &pb.EchoRequest{
			Response: &pb.EchoRequest_Content{
				Content: str,
			},
		}
		if err = s.Send(req); err != nil {
			t.Fatal(err)
		}
	}

	if err = s.CloseSend(); err != nil {
		t.Fatal(err)
	}

	var rests []string

	for {
		resp, err2 := s.Recv()
		if errors.Is(err2, io.EOF) {
			break
		}
		if err2 != nil {
			t.Fatal(err2)
		}
		rests = append(rests, resp.GetContent())
	}
	got := strings.Join(rests, " ")
	if content != got {
		t.Errorf("Chat() = %q, want %q", got, content)
	}
}

func TestWait(t *testing.T) {
	content := "hello world!"
	req := &pb.WaitRequest{
		End: &pb.WaitRequest_Ttl{
			Ttl: &durationpb.Duration{Seconds: 2},
		},
		Response: &pb.WaitRequest_Success{
			Success: &pb.WaitResponse{
				Content: content,
			},
		},
	}

	for typ, client := range map[string]*showcase.EchoClient{
		"grpc": echoGRPC,
		"rest": echoREST,
	} {
		op, err := client.Wait(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}
		resp, err := op.Wait(context.Background())
		if err != nil {
			t.Fatal(err)
		}

		if resp.GetContent() != content {
			t.Errorf("%s Wait() = %q, want %q", typ, resp.GetContent(), content)
		}
	}

}

func TestPagination(t *testing.T) {
	content := "foo bar biz baz"
	expected := strings.Split(content, " ")
	req := &pb.PagedExpandRequest{
		Content:  content,
		PageSize: 2,
	}

	for typ, client := range map[string]*showcase.EchoClient{
		"grpc": echoGRPC,
		"rest": echoREST,
	} {
		it := client.PagedExpand(context.Background(), req)

		ndx := 0
		for {
			resp, err := it.Next()
			if errors.Is(err, iterator.Done) {
				break
			}

			if err != nil {
				t.Fatal(err)
			}

			if resp.GetContent() != expected[ndx] {
				t.Errorf("%s Chat() = %s, want %s", typ, resp.GetContent(), expected[ndx])
			}
			ndx++
		}
	}
}

func TestPaginationWithToken(t *testing.T) {
	content := "ab cd ef gh ij kl"
	expected := strings.Split(content, " ")[1:]
	req := &pb.PagedExpandRequest{
		Content:   content,
		PageSize:  2,
		PageToken: "1",
	}

	for typ, client := range map[string]*showcase.EchoClient{
		"grpc": echoGRPC,
		"rest": echoREST,
	} {
		it := client.PagedExpand(context.Background(), req)

		ndx := 0
		for {
			resp, err := it.Next()
			if errors.Is(err, iterator.Done) {
				break
			}

			if err != nil {
				t.Fatal(err)
			}

			if ndx >= len(expected) {
				t.Errorf("%s PagedExpand() received more items than expected", typ)
			} else if resp.GetContent() != expected[ndx] {
				t.Errorf("%s PagedExpand() = %s, want %s", typ, resp.GetContent(), expected[ndx])
			}
			ndx++
		}

	}

}

func TestBlock(t *testing.T) {
	content := "hello world!"
	req := &pb.BlockRequest{
		ResponseDelay: &durationpb.Duration{Nanos: 1000},
		Response: &pb.BlockRequest_Success{
			Success: &pb.BlockResponse{
				Content: content,
			},
		},
	}
	resp, err := echoGRPC.Block(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.GetContent() != content {
		t.Errorf("Block() = %q, want %q", resp.GetContent(), content)
	}
}

func TestBlock_timeout(t *testing.T) {
	content := "hello world!"
	req := &pb.BlockRequest{
		ResponseDelay: &durationpb.Duration{Seconds: 1},
		Response: &pb.BlockRequest_Success{
			Success: &pb.BlockResponse{
				Content: content,
			},
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	want := status.New(codes.DeadlineExceeded, "context deadline exceeded")
	resp, err := echoGRPC.Block(ctx, req)
	if err == nil {
		t.Errorf("Block() got %+v, want %+v", resp, want)
	} else if got, ok := status.FromError(err); !ok || got.Code() != want.Code() {
		t.Errorf("Block() got %+v, want %+v", err, want)
	}
}

func TestBlock_default_timeout(t *testing.T) {
	content := "hello world!"
	req := &pb.BlockRequest{
		ResponseDelay: &durationpb.Duration{Seconds: 6},
		Response: &pb.BlockRequest_Success{
			Success: &pb.BlockResponse{
				Content: content,
			},
		},
	}

	want := status.New(codes.DeadlineExceeded, "context deadline exceeded")
	resp, err := echoGRPC.Block(context.Background(), req)
	if err == nil {
		t.Errorf("Block() got %+v, want %+v", resp, want)
	} else if got, ok := status.FromError(err); !ok || got.Code() != want.Code() {
		t.Errorf("Block() got %+v, want %+v", err, want)
	}
}

func TestBlock_override_default_timeout(t *testing.T) {
	content := "hello world!"
	req := &pb.BlockRequest{
		ResponseDelay: &durationpb.Duration{Seconds: 6},
		Response: &pb.BlockRequest_Success{
			Success: &pb.BlockResponse{
				Content: content,
			},
		},
	}

	resp, err := echoGRPC.Block(context.Background(), req, gax.WithTimeout(10*time.Second))
	if err != nil {
		t.Error(err)
	}
	if resp.GetContent() != content {
		t.Errorf("Block() = %q, want %q", resp.GetContent(), content)
	}
}
