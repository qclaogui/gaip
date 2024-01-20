// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package generativeai_test

import (
	"context"
	"strings"
	"testing"

	generativelanguage "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1"
	pb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1/generativelanguagepb"
)

// Clients are initialized in main_test.go.
var (
	generativelanguageGRPC *generativelanguage.GenerativeClient
	generativelanguageREST *generativelanguage.GenerativeClient
)

func TestCountTokens(t *testing.T) {
	// only when genai.enabled=true
	t.Skip()
	text := &pb.Part{
		Data: &pb.Part_Text{Text: "The rain in Spain falls mainly on the plain."},
	}
	req := newCountTokensRequest(newUserContent(text))

	for typ, client := range map[string]*generativelanguage.GenerativeClient{
		"grpc": generativelanguageGRPC,
		//"rest": generativelanguageREST,
	} {
		resp, err := client.CountTokens(context.Background(), req)
		if err != nil {
			t.Fatalf("%s client.CountTokens() failed: %v", typ, err)
		}
		if g, w := resp.TotalTokens, int32(11); g != w {
			t.Errorf("%s client.CountTokens() got resp.TotalTokens %d, want %d", typ, g, w)
		}
	}
}

func newUserContent(parts ...*pb.Part) *pb.Content {
	return &pb.Content{Role: "user", Parts: parts}
}

func newCountTokensRequest(contents ...*pb.Content) *pb.CountTokensRequest {
	return &pb.CountTokensRequest{
		Model:    fullModelName("gemini-pro"),
		Contents: contents,
	}
}

func fullModelName(name string) string {
	if strings.HasPrefix(name, "models/") {
		return name
	}
	return "models/" + name
}
