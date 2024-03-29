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

package bookstore_test

import (
	"context"

	bookstore "github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1"
	bookstorepb "github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func ExampleNewClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := bookstore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateBook() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := bookstore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &bookstorepb.CreateBookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb#CreateBookRequest.
	}
	resp, err := c.CreateBook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateShelf() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := bookstore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &bookstorepb.CreateShelfRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb#CreateShelfRequest.
	}
	resp, err := c.CreateShelf(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteBook() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := bookstore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &bookstorepb.DeleteBookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb#DeleteBookRequest.
	}
	err = c.DeleteBook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_DeleteShelf() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := bookstore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &bookstorepb.DeleteShelfRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb#DeleteShelfRequest.
	}
	err = c.DeleteShelf(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetBook() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := bookstore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &bookstorepb.GetBookRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb#GetBookRequest.
	}
	resp, err := c.GetBook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetShelf() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := bookstore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &bookstorepb.GetShelfRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb#GetShelfRequest.
	}
	resp, err := c.GetShelf(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListBooks() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := bookstore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &bookstorepb.ListBooksRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb#ListBooksRequest.
	}
	resp, err := c.ListBooks(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListShelves() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := bookstore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &emptypb.Empty{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/protobuf/types/known/emptypb#Empty.
	}
	resp, err := c.ListShelves(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
