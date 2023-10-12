// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package project_test

import (
	"context"
	"io"

	project "github.com/qclaogui/golang-api-server/genproto/project/apiv1"
	projectpb "github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb"
	"google.golang.org/api/iterator"
)

func ExampleNewEchoClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := project.NewEchoClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewEchoRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := project.NewEchoRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleEchoClient_Echo() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := project.NewEchoClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &projectpb.EchoRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb#EchoRequest.
	}
	resp, err := c.Echo(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEchoClient_Chat() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := project.NewEchoClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()
	stream, err := c.Chat(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	go func() {
		reqs := []*projectpb.EchoRequest{
			// TODO: Create requests.
		}
		for _, req := range reqs {
			if err := stream.Send(req); err != nil {
				// TODO: Handle error.
			}
		}
		stream.CloseSend()
	}()
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			// TODO: handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleEchoClient_PagedExpand() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := project.NewEchoClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &projectpb.PagedExpandRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb#PagedExpandRequest.
	}
	it := c.PagedExpand(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleEchoClient_Wait() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := project.NewEchoClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &projectpb.WaitRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb#WaitRequest.
	}
	op, err := c.Wait(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEchoClient_Block() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := project.NewEchoClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &projectpb.BlockRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb#BlockRequest.
	}
	resp, err := c.Block(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
