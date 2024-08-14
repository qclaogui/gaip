// Copyright 2024 Google LLC
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

	generativelanguage "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1"
	generativelanguagepb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb"
	"google.golang.org/api/iterator"
)

func ExampleNewRetrieverClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewRetrieverRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleRetrieverClient_BatchCreateChunks() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.BatchCreateChunksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#BatchCreateChunksRequest.
	}
	resp, err := c.BatchCreateChunks(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_BatchDeleteChunks() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.BatchDeleteChunksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#BatchDeleteChunksRequest.
	}
	err = c.BatchDeleteChunks(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleRetrieverClient_BatchUpdateChunks() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.BatchUpdateChunksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#BatchUpdateChunksRequest.
	}
	resp, err := c.BatchUpdateChunks(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_CreateChunk() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.CreateChunkRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#CreateChunkRequest.
	}
	resp, err := c.CreateChunk(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_CreateCorpus() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.CreateCorpusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#CreateCorpusRequest.
	}
	resp, err := c.CreateCorpus(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_CreateDocument() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.CreateDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#CreateDocumentRequest.
	}
	resp, err := c.CreateDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_DeleteChunk() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.DeleteChunkRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#DeleteChunkRequest.
	}
	err = c.DeleteChunk(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleRetrieverClient_DeleteCorpus() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.DeleteCorpusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#DeleteCorpusRequest.
	}
	err = c.DeleteCorpus(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleRetrieverClient_DeleteDocument() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.DeleteDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#DeleteDocumentRequest.
	}
	err = c.DeleteDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleRetrieverClient_GetChunk() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.GetChunkRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#GetChunkRequest.
	}
	resp, err := c.GetChunk(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_GetCorpus() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.GetCorpusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#GetCorpusRequest.
	}
	resp, err := c.GetCorpus(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_GetDocument() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.GetDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#GetDocumentRequest.
	}
	resp, err := c.GetDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_ListChunks() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.ListChunksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#ListChunksRequest.
	}
	it := c.ListChunks(ctx, req)
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
		_ = it.Response.(*generativelanguagepb.ListChunksResponse)
	}
}

func ExampleRetrieverClient_ListCorpora() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.ListCorporaRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#ListCorporaRequest.
	}
	it := c.ListCorpora(ctx, req)
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
		_ = it.Response.(*generativelanguagepb.ListCorporaResponse)
	}
}

func ExampleRetrieverClient_ListDocuments() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.ListDocumentsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#ListDocumentsRequest.
	}
	it := c.ListDocuments(ctx, req)
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
		_ = it.Response.(*generativelanguagepb.ListDocumentsResponse)
	}
}

func ExampleRetrieverClient_QueryCorpus() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.QueryCorpusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#QueryCorpusRequest.
	}
	resp, err := c.QueryCorpus(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_QueryDocument() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.QueryDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#QueryDocumentRequest.
	}
	resp, err := c.QueryDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_UpdateChunk() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.UpdateChunkRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#UpdateChunkRequest.
	}
	resp, err := c.UpdateChunk(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_UpdateCorpus() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.UpdateCorpusRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#UpdateCorpusRequest.
	}
	resp, err := c.UpdateCorpus(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleRetrieverClient_UpdateDocument() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := generativelanguage.NewRetrieverClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &generativelanguagepb.UpdateDocumentRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb#UpdateDocumentRequest.
	}
	resp, err := c.UpdateDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
