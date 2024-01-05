// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project_test

import (
	"context"
	"testing"

	project "github.com/qclaogui/gaip/genproto/project/apiv1"
	pb "github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
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
