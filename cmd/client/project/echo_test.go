// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project_test

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
	"time"

	project "github.com/qclaogui/gaip/genproto/project/apiv1"
	pb "github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/durationpb"
)

// Clients are initialized in main_test.go.
var (
	echoGRPC *project.EchoClient
	echoREST *project.EchoClient
)

func TestEcho(t *testing.T) {
	t.Skip()
	content := "hello world!"
	req := &pb.EchoRequest{
		Response: &pb.EchoRequest_Content{
			Content: content,
		},
	}

	for typ, client := range map[string]*project.EchoClient{"grpc": echoGRPC, "rest": echoREST} {
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}

		if resp.GetContent() != req.GetContent() {
			t.Errorf("%s Echo() = %q, want %q", typ, resp.GetContent(), content)
		}
	}

}

// Chat, Collect, and Expand are streaming methods and don't have interesting REST semantics
func TestExpand(t *testing.T) {
	t.Skip()
	content := "The rain in Spain stays mainly on the plain!"
	req := &pb.ExpandRequest{
		Content: content,
	}

	for typ, client := range map[string]*project.EchoClient{"grpc": echoGRPC, "rest": echoREST} {
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
	t.Skip()
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
	t.Skip()
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
	t.Skip()
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

	for typ, client := range map[string]*project.EchoClient{"grpc": echoGRPC, "rest": echoREST} {
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

func TestBlock(t *testing.T) {
	t.Skip()
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
	//t.Skip()
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
	if err != nil {
		t.Errorf("Block() got %+v, want %+v", resp, want)
	} else if got, ok := status.FromError(err); !ok || got.Code() != want.Code() {
		t.Errorf("Block() got %+v, want %+v", err, want)
	}
}
