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

package generativelanguage

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"

	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	gax "github.com/googleapis/gax-go/v2"
	generativelanguagepb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newCacheClientHook clientHook

// CacheCallOptions contains the retry settings for each method of CacheClient.
type CacheCallOptions struct {
	ListCachedContents  []gax.CallOption
	CreateCachedContent []gax.CallOption
	GetCachedContent    []gax.CallOption
	UpdateCachedContent []gax.CallOption
	DeleteCachedContent []gax.CallOption
	GetOperation        []gax.CallOption
	ListOperations      []gax.CallOption
}

func defaultCacheGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("generativelanguage.qclaogui.com:443"),
		internaloption.WithDefaultEndpointTemplate("generativelanguage.qclaogui.com:443"),
		internaloption.WithDefaultMTLSEndpoint("generativelanguage.qclaogui.com:443"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://generativelanguage.qclaogui.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCacheCallOptions() *CacheCallOptions {
	return &CacheCallOptions{
		ListCachedContents:  []gax.CallOption{},
		CreateCachedContent: []gax.CallOption{},
		GetCachedContent:    []gax.CallOption{},
		UpdateCachedContent: []gax.CallOption{},
		DeleteCachedContent: []gax.CallOption{},
		GetOperation:        []gax.CallOption{},
		ListOperations:      []gax.CallOption{},
	}
}

func defaultCacheRESTCallOptions() *CacheCallOptions {
	return &CacheCallOptions{
		ListCachedContents:  []gax.CallOption{},
		CreateCachedContent: []gax.CallOption{},
		GetCachedContent:    []gax.CallOption{},
		UpdateCachedContent: []gax.CallOption{},
		DeleteCachedContent: []gax.CallOption{},
		GetOperation:        []gax.CallOption{},
		ListOperations:      []gax.CallOption{},
	}
}

// internalCacheClient is an interface that defines the methods available from Generative Language API.
type internalCacheClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListCachedContents(context.Context, *generativelanguagepb.ListCachedContentsRequest, ...gax.CallOption) *CachedContentIterator
	CreateCachedContent(context.Context, *generativelanguagepb.CreateCachedContentRequest, ...gax.CallOption) (*generativelanguagepb.CachedContent, error)
	GetCachedContent(context.Context, *generativelanguagepb.GetCachedContentRequest, ...gax.CallOption) (*generativelanguagepb.CachedContent, error)
	UpdateCachedContent(context.Context, *generativelanguagepb.UpdateCachedContentRequest, ...gax.CallOption) (*generativelanguagepb.CachedContent, error)
	DeleteCachedContent(context.Context, *generativelanguagepb.DeleteCachedContentRequest, ...gax.CallOption) error
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
	ListOperations(context.Context, *longrunningpb.ListOperationsRequest, ...gax.CallOption) *OperationIterator
}

// CacheClient is a client for interacting with Generative Language API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// API for managing cache of content (CachedContent resources) that can be used
// in GenerativeService requests. This way generate content requests can benefit
// from preprocessing work being done earlier, possibly lowering their
// computational cost. It is intended to be used with large contexts.
type CacheClient struct {
	// The internal transport-dependent client.
	internalClient internalCacheClient

	// The call options for this service.
	CallOptions *CacheCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CacheClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CacheClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *CacheClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListCachedContents lists CachedContents.
func (c *CacheClient) ListCachedContents(ctx context.Context, req *generativelanguagepb.ListCachedContentsRequest, opts ...gax.CallOption) *CachedContentIterator {
	return c.internalClient.ListCachedContents(ctx, req, opts...)
}

// CreateCachedContent creates CachedContent resource.
func (c *CacheClient) CreateCachedContent(ctx context.Context, req *generativelanguagepb.CreateCachedContentRequest, opts ...gax.CallOption) (*generativelanguagepb.CachedContent, error) {
	return c.internalClient.CreateCachedContent(ctx, req, opts...)
}

// GetCachedContent reads CachedContent resource.
func (c *CacheClient) GetCachedContent(ctx context.Context, req *generativelanguagepb.GetCachedContentRequest, opts ...gax.CallOption) (*generativelanguagepb.CachedContent, error) {
	return c.internalClient.GetCachedContent(ctx, req, opts...)
}

// UpdateCachedContent updates CachedContent resource (only expiration is updatable).
func (c *CacheClient) UpdateCachedContent(ctx context.Context, req *generativelanguagepb.UpdateCachedContentRequest, opts ...gax.CallOption) (*generativelanguagepb.CachedContent, error) {
	return c.internalClient.UpdateCachedContent(ctx, req, opts...)
}

// DeleteCachedContent deletes CachedContent resource.
func (c *CacheClient) DeleteCachedContent(ctx context.Context, req *generativelanguagepb.DeleteCachedContentRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteCachedContent(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *CacheClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *CacheClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	return c.internalClient.ListOperations(ctx, req, opts...)
}

// cacheGRPCClient is a client for interacting with Generative Language API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type cacheGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing CacheClient
	CallOptions **CacheCallOptions

	// The gRPC API client.
	cacheClient generativelanguagepb.CacheServiceClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewCacheClient creates a new cache service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// API for managing cache of content (CachedContent resources) that can be used
// in GenerativeService requests. This way generate content requests can benefit
// from preprocessing work being done earlier, possibly lowering their
// computational cost. It is intended to be used with large contexts.
func NewCacheClient(ctx context.Context, opts ...option.ClientOption) (*CacheClient, error) {
	clientOpts := defaultCacheGRPCClientOptions()
	if newCacheClientHook != nil {
		hookOpts, err := newCacheClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CacheClient{CallOptions: defaultCacheCallOptions()}

	c := &cacheGRPCClient{
		connPool:         connPool,
		cacheClient:      generativelanguagepb.NewCacheServiceClient(connPool),
		CallOptions:      &client.CallOptions,
		logger:           internaloption.GetLogger(opts),
		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *cacheGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *cacheGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *cacheGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type cacheRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing CacheClient
	CallOptions **CacheCallOptions

	logger *slog.Logger
}

// NewCacheRESTClient creates a new cache service rest client.
//
// API for managing cache of content (CachedContent resources) that can be used
// in GenerativeService requests. This way generate content requests can benefit
// from preprocessing work being done earlier, possibly lowering their
// computational cost. It is intended to be used with large contexts.
func NewCacheRESTClient(ctx context.Context, opts ...option.ClientOption) (*CacheClient, error) {
	clientOpts := append(defaultCacheRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultCacheRESTCallOptions()
	c := &cacheRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &CacheClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultCacheRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://generativelanguage.qclaogui.com"),
		internaloption.WithDefaultEndpointTemplate("https://generativelanguage.qclaogui.com"),
		internaloption.WithDefaultMTLSEndpoint("https://generativelanguage.qclaogui.com"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://generativelanguage.qclaogui.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *cacheRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *cacheRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *cacheRESTClient) Connection() *grpc.ClientConn {
	return nil
}

func (c *cacheGRPCClient) ListCachedContents(ctx context.Context, req *generativelanguagepb.ListCachedContentsRequest, opts ...gax.CallOption) *CachedContentIterator {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListCachedContents[0:len((*c.CallOptions).ListCachedContents):len((*c.CallOptions).ListCachedContents)], opts...)
	it := &CachedContentIterator{}
	req = proto.Clone(req).(*generativelanguagepb.ListCachedContentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*generativelanguagepb.CachedContent, string, error) {
		resp := &generativelanguagepb.ListCachedContentsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.cacheClient.ListCachedContents, req, settings.GRPC, c.logger, "ListCachedContents")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetCachedContents(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *cacheGRPCClient) CreateCachedContent(ctx context.Context, req *generativelanguagepb.CreateCachedContentRequest, opts ...gax.CallOption) (*generativelanguagepb.CachedContent, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).CreateCachedContent[0:len((*c.CallOptions).CreateCachedContent):len((*c.CallOptions).CreateCachedContent)], opts...)
	var resp *generativelanguagepb.CachedContent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.cacheClient.CreateCachedContent, req, settings.GRPC, c.logger, "CreateCachedContent")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *cacheGRPCClient) GetCachedContent(ctx context.Context, req *generativelanguagepb.GetCachedContentRequest, opts ...gax.CallOption) (*generativelanguagepb.CachedContent, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetCachedContent[0:len((*c.CallOptions).GetCachedContent):len((*c.CallOptions).GetCachedContent)], opts...)
	var resp *generativelanguagepb.CachedContent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.cacheClient.GetCachedContent, req, settings.GRPC, c.logger, "GetCachedContent")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *cacheGRPCClient) UpdateCachedContent(ctx context.Context, req *generativelanguagepb.UpdateCachedContentRequest, opts ...gax.CallOption) (*generativelanguagepb.CachedContent, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "cached_content.name", url.QueryEscape(req.GetCachedContent().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateCachedContent[0:len((*c.CallOptions).UpdateCachedContent):len((*c.CallOptions).UpdateCachedContent)], opts...)
	var resp *generativelanguagepb.CachedContent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.cacheClient.UpdateCachedContent, req, settings.GRPC, c.logger, "UpdateCachedContent")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *cacheGRPCClient) DeleteCachedContent(ctx context.Context, req *generativelanguagepb.DeleteCachedContentRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteCachedContent[0:len((*c.CallOptions).DeleteCachedContent):len((*c.CallOptions).DeleteCachedContent)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = executeRPC(ctx, c.cacheClient.DeleteCachedContent, req, settings.GRPC, c.logger, "DeleteCachedContent")
		return err
	}, opts...)
	return err
}

func (c *cacheGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.operationsClient.GetOperation, req, settings.GRPC, c.logger, "GetOperation")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *cacheGRPCClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = executeRPC(ctx, c.operationsClient.ListOperations, req, settings.GRPC, c.logger, "ListOperations")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// ListCachedContents lists CachedContents.
func (c *cacheRESTClient) ListCachedContents(ctx context.Context, req *generativelanguagepb.ListCachedContentsRequest, opts ...gax.CallOption) *CachedContentIterator {
	it := &CachedContentIterator{}
	req = proto.Clone(req).(*generativelanguagepb.ListCachedContentsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*generativelanguagepb.CachedContent, string, error) {
		resp := &generativelanguagepb.ListCachedContentsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1beta/cachedContents")

		params := url.Values{}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListCachedContents")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetCachedContents(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

// CreateCachedContent creates CachedContent resource.
func (c *cacheRESTClient) CreateCachedContent(ctx context.Context, req *generativelanguagepb.CreateCachedContentRequest, opts ...gax.CallOption) (*generativelanguagepb.CachedContent, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetCachedContent()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta/cachedContents")

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateCachedContent[0:len((*c.CallOptions).CreateCachedContent):len((*c.CallOptions).CreateCachedContent)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.CachedContent{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CreateCachedContent")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetCachedContent reads CachedContent resource.
func (c *cacheRESTClient) GetCachedContent(ctx context.Context, req *generativelanguagepb.GetCachedContentRequest, opts ...gax.CallOption) (*generativelanguagepb.CachedContent, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetCachedContent[0:len((*c.CallOptions).GetCachedContent):len((*c.CallOptions).GetCachedContent)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.CachedContent{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetCachedContent")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// UpdateCachedContent updates CachedContent resource (only expiration is updatable).
func (c *cacheRESTClient) UpdateCachedContent(ctx context.Context, req *generativelanguagepb.UpdateCachedContentRequest, opts ...gax.CallOption) (*generativelanguagepb.CachedContent, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetCachedContent()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta/%v", req.GetCachedContent().GetName())

	params := url.Values{}
	if req.GetUpdateMask() != nil {
		field, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(field[1:len(field)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "cached_content.name", url.QueryEscape(req.GetCachedContent().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateCachedContent[0:len((*c.CallOptions).UpdateCachedContent):len((*c.CallOptions).UpdateCachedContent)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.CachedContent{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateCachedContent")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// DeleteCachedContent deletes CachedContent resource.
func (c *cacheRESTClient) DeleteCachedContent(ctx context.Context, req *generativelanguagepb.DeleteCachedContentRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		_, err = executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteCachedContent")
		return err
	}, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *cacheRESTClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetOperation")
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ListOperations is a utility method from google.longrunning.Operations.
func (c *cacheRESTClient) ListOperations(ctx context.Context, req *longrunningpb.ListOperationsRequest, opts ...gax.CallOption) *OperationIterator {
	it := &OperationIterator{}
	req = proto.Clone(req).(*longrunningpb.ListOperationsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*longrunningpb.Operation, string, error) {
		resp := &longrunningpb.ListOperationsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		baseUrl, err := url.Parse(c.endpoint)
		if err != nil {
			return nil, "", err
		}
		baseUrl.Path += fmt.Sprintf("/v1beta1/%v/operations", req.GetName())

		params := url.Values{}
		if req.GetFilter() != "" {
			params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
		}
		if req.GetPageSize() != 0 {
			params.Add("pageSize", fmt.Sprintf("%v", req.GetPageSize()))
		}
		if req.GetPageToken() != "" {
			params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
		}

		baseUrl.RawQuery = params.Encode()

		// Build HTTP headers from client and context metadata.
		hds := append(c.xGoogHeaders, "Content-Type", "application/json")
		headers := gax.BuildHeaders(ctx, hds...)
		e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			if settings.Path != "" {
				baseUrl.Path = settings.Path
			}
			httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
			if err != nil {
				return err
			}
			httpReq.Header = headers

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListOperations")
			if err != nil {
				return err
			}
			if err := unm.Unmarshal(buf, resp); err != nil {
				return err
			}

			return nil
		}, opts...)
		if e != nil {
			return nil, "", e
		}
		it.Response = resp
		return resp.GetOperations(), resp.GetNextPageToken(), nil
	}

	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}
