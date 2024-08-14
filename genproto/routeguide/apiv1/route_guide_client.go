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

package routeguide

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	routeguidepb "github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	GetFeature   []gax.CallOption
	ListFeatures []gax.CallOption
	RecordRoute  []gax.CallOption
	RouteChat    []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("localhost:9095"),
		internaloption.WithDefaultEndpointTemplate("localhost:9095"),
		internaloption.WithDefaultMTLSEndpoint("localhost:9095"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://localhost/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		internaloption.EnableNewAuthLibrary(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		GetFeature: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    10 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListFeatures: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    10 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		RecordRoute: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    10 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		RouteChat: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    10 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultRESTCallOptions() *CallOptions {
	return &CallOptions{
		GetFeature: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    10 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ListFeatures: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    10 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		RecordRoute: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    10 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		RouteChat: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    10 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalClient is an interface that defines the methods available from Client Libraries Routeguide API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetFeature(context.Context, *routeguidepb.GetFeatureRequest, ...gax.CallOption) (*routeguidepb.GetFeatureResponse, error)
	ListFeatures(context.Context, *routeguidepb.ListFeaturesRequest, ...gax.CallOption) (routeguidepb.RouteGuideService_ListFeaturesClient, error)
	RecordRoute(context.Context, ...gax.CallOption) (routeguidepb.RouteGuideService_RecordRouteClient, error)
	RouteChat(context.Context, ...gax.CallOption) (routeguidepb.RouteGuideService_RouteChatClient, error)
}

// Client is a client for interacting with Client Libraries Routeguide API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Interface exported by the server.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetFeature a simple RPC.
//
// Obtains the feature at a given position.
//
// A feature with an empty name is returned if there’s no feature at the given
// position.
func (c *Client) GetFeature(ctx context.Context, req *routeguidepb.GetFeatureRequest, opts ...gax.CallOption) (*routeguidepb.GetFeatureResponse, error) {
	return c.internalClient.GetFeature(ctx, req, opts...)
}

// ListFeatures a server-to-client streaming RPC.
//
// Obtains the Features available within the given Rectangle.  Results are
// streamed rather than returned at once (e.g. in a response message with a
// repeated field), as the rectangle may cover a large area and contain a
// huge number of features.
func (c *Client) ListFeatures(ctx context.Context, req *routeguidepb.ListFeaturesRequest, opts ...gax.CallOption) (routeguidepb.RouteGuideService_ListFeaturesClient, error) {
	return c.internalClient.ListFeatures(ctx, req, opts...)
}

// RecordRoute a client-to-server streaming RPC.
//
// Accepts a stream of Points on a route being traversed, returning a
// RouteSummary when traversal is completed.
//
// This method is not supported for the REST transport.
func (c *Client) RecordRoute(ctx context.Context, opts ...gax.CallOption) (routeguidepb.RouteGuideService_RecordRouteClient, error) {
	return c.internalClient.RecordRoute(ctx, opts...)
}

// RouteChat a Bidirectional streaming RPC.
//
// Accepts a stream of RouteChatRequest sent while a route is being traversed,
// while receiving other RouteChatResponse (e.g. from other users).
//
// This method is not supported for the REST transport.
func (c *Client) RouteChat(ctx context.Context, opts ...gax.CallOption) (routeguidepb.RouteGuideService_RouteChatClient, error) {
	return c.internalClient.RouteChat(ctx, opts...)
}

// gRPCClient is a client for interacting with Client Libraries Routeguide API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client routeguidepb.RouteGuideServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewClient creates a new route guide service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Interface exported by the server.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:    connPool,
		client:      routeguidepb.NewRouteGuideServiceClient(connPool),
		CallOptions: &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
		"x-goog-api-version", "v1_20241015",
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type restClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions
}

// NewRESTClient creates a new route guide service rest client.
//
// Interface exported by the server.
func NewRESTClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := append(defaultRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultRESTCallOptions()
	c := &restClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &Client{internalClient: c, CallOptions: callOpts}, nil
}

func defaultRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://localhost:9095"),
		internaloption.WithDefaultEndpointTemplate("https://localhost:9095"),
		internaloption.WithDefaultMTLSEndpoint("https://localhost:9095"),
		internaloption.WithDefaultUniverseDomain("googleapis.com"),
		internaloption.WithDefaultAudience("https://localhost/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableNewAuthLibrary(),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *restClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
		"x-goog-api-version", "v1_20241015",
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *restClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *restClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *gRPCClient) GetFeature(ctx context.Context, req *routeguidepb.GetFeatureRequest, opts ...gax.CallOption) (*routeguidepb.GetFeatureResponse, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).GetFeature[0:len((*c.CallOptions).GetFeature):len((*c.CallOptions).GetFeature)], opts...)
	var resp *routeguidepb.GetFeatureResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetFeature(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListFeatures(ctx context.Context, req *routeguidepb.ListFeaturesRequest, opts ...gax.CallOption) (routeguidepb.RouteGuideService_ListFeaturesClient, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListFeatures[0:len((*c.CallOptions).ListFeatures):len((*c.CallOptions).ListFeatures)], opts...)
	var resp routeguidepb.RouteGuideService_ListFeaturesClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ListFeatures(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) RecordRoute(ctx context.Context, opts ...gax.CallOption) (routeguidepb.RouteGuideService_RecordRouteClient, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	var resp routeguidepb.RouteGuideService_RecordRouteClient
	opts = append((*c.CallOptions).RecordRoute[0:len((*c.CallOptions).RecordRoute):len((*c.CallOptions).RecordRoute)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.RecordRoute(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) RouteChat(ctx context.Context, opts ...gax.CallOption) (routeguidepb.RouteGuideService_RouteChatClient, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	var resp routeguidepb.RouteGuideService_RouteChatClient
	opts = append((*c.CallOptions).RouteChat[0:len((*c.CallOptions).RouteChat):len((*c.CallOptions).RouteChat)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.RouteChat(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetFeature a simple RPC.
//
// Obtains the feature at a given position.
//
// A feature with an empty name is returned if there’s no feature at the given
// position.
func (c *restClient) GetFeature(ctx context.Context, req *routeguidepb.GetFeatureRequest, opts ...gax.CallOption) (*routeguidepb.GetFeatureResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/routeguide:get-feature")

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetFeature[0:len((*c.CallOptions).GetFeature):len((*c.CallOptions).GetFeature)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &routeguidepb.GetFeatureResponse{}
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

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
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

// ListFeatures a server-to-client streaming RPC.
//
// Obtains the Features available within the given Rectangle.  Results are
// streamed rather than returned at once (e.g. in a response message with a
// repeated field), as the rectangle may cover a large area and contain a
// huge number of features.
func (c *restClient) ListFeatures(ctx context.Context, req *routeguidepb.ListFeaturesRequest, opts ...gax.CallOption) (routeguidepb.RouteGuideService_ListFeaturesClient, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/routeguide:list-feature")

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	var streamClient *listFeaturesRESTClient
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

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		streamClient = &listFeaturesRESTClient{
			ctx:    ctx,
			md:     metadata.MD(httpRsp.Header),
			stream: gax.NewProtoJSONStreamReader(httpRsp.Body, (&routeguidepb.ListFeaturesResponse{}).ProtoReflect().Type()),
		}
		return nil
	}, opts...)

	return streamClient, e
}

// listFeaturesRESTClient is the stream client used to consume the server stream created by
// the REST implementation of ListFeatures.
type listFeaturesRESTClient struct {
	ctx    context.Context
	md     metadata.MD
	stream *gax.ProtoJSONStream
}

func (c *listFeaturesRESTClient) Recv() (*routeguidepb.ListFeaturesResponse, error) {
	if err := c.ctx.Err(); err != nil {
		defer c.stream.Close()
		return nil, err
	}
	msg, err := c.stream.Recv()
	if err != nil {
		defer c.stream.Close()
		return nil, err
	}
	res := msg.(*routeguidepb.ListFeaturesResponse)
	return res, nil
}

func (c *listFeaturesRESTClient) Header() (metadata.MD, error) {
	return c.md, nil
}

func (c *listFeaturesRESTClient) Trailer() metadata.MD {
	return c.md
}

func (c *listFeaturesRESTClient) CloseSend() error {
	// This is a no-op to fulfill the interface.
	return errors.New("this method is not implemented for a server-stream")
}

func (c *listFeaturesRESTClient) Context() context.Context {
	return c.ctx
}

func (c *listFeaturesRESTClient) SendMsg(m interface{}) error {
	// This is a no-op to fulfill the interface.
	return errors.New("this method is not implemented for a server-stream")
}

func (c *listFeaturesRESTClient) RecvMsg(m interface{}) error {
	// This is a no-op to fulfill the interface.
	return errors.New("this method is not implemented, use Recv")
}

// RecordRoute a client-to-server streaming RPC.
//
// Accepts a stream of Points on a route being traversed, returning a
// RouteSummary when traversal is completed.
//
// This method is not supported for the REST transport.
func (c *restClient) RecordRoute(ctx context.Context, opts ...gax.CallOption) (routeguidepb.RouteGuideService_RecordRouteClient, error) {
	return nil, errors.New("RecordRoute not yet supported for REST clients")
}

// RouteChat a Bidirectional streaming RPC.
//
// Accepts a stream of RouteChatRequest sent while a route is being traversed,
// while receiving other RouteChatResponse (e.g. from other users).
//
// This method is not supported for the REST transport.
func (c *restClient) RouteChat(ctx context.Context, opts ...gax.CallOption) (routeguidepb.RouteGuideService_RouteChatClient, error) {
	return nil, errors.New("RouteChat not yet supported for REST clients")
}
