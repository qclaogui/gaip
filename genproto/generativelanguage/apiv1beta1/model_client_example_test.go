// Copyright 2025 Google LLC
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

package generativelanguage_test

import (
	"context"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	generativelanguage "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1"
	generativelanguagepb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb"
	"google.golang.org/api/iterator"
)

func ExampleNewModelClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewModelRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleModelClient_CreateTunedModel() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.CreateTunedModelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#CreateTunedModelRequest.
	}
	op, err := c.CreateTunedModel(ctx, req)
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

func ExampleModelClient_DeleteTunedModel() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.DeleteTunedModelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#DeleteTunedModelRequest.
	}
	err = c.DeleteTunedModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleModelClient_GetModel() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.GetModelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#GetModelRequest.
	}
	resp, err := c.GetModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelClient_GetTunedModel() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.GetTunedModelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#GetTunedModelRequest.
	}
	resp, err := c.GetTunedModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelClient_ListModels() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.ListModelsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#ListModelsRequest.
	}
	it := c.ListModels(ctx, req)
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

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*generativelanguagepb.ListModelsResponse)
	}
}

func ExampleModelClient_ListTunedModels() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.ListTunedModelsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#ListTunedModelsRequest.
	}
	it := c.ListTunedModels(ctx, req)
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

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*generativelanguagepb.ListTunedModelsResponse)
	}
}

func ExampleModelClient_UpdateTunedModel() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.UpdateTunedModelRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#UpdateTunedModelRequest.
	}
	resp, err := c.UpdateTunedModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelClient_GetOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleModelClient_ListOperations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewModelClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/longrunning/autogen/longrunningpb#ListOperationsRequest.
	}
	it := c.ListOperations(ctx, req)
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

		// If you need to access the underlying RPC response,
		// you can do so by casting the `Response` as below.
		// Otherwise, remove this line. Only populated after
		// first call to Next(). Not safe for concurrent access.
		_ = it.Response.(*longrunningpb.ListOperationsResponse)
	}
}
