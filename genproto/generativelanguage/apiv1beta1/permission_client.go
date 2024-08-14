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

package generativelanguage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	generativelanguagepb "github.com/qclaogui/gaip/genproto/generativelanguage/apiv1beta1/generativelanguagepb"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newPermissionClientHook clientHook

// PermissionCallOptions contains the retry settings for each method of PermissionClient.
type PermissionCallOptions struct {
	CreatePermission  []gax.CallOption
	GetPermission     []gax.CallOption
	ListPermissions   []gax.CallOption
	UpdatePermission  []gax.CallOption
	DeletePermission  []gax.CallOption
	TransferOwnership []gax.CallOption
}

func defaultPermissionGRPCClientOptions() []option.ClientOption {
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

func defaultPermissionCallOptions() *PermissionCallOptions {
	return &PermissionCallOptions{
		CreatePermission: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetPermission: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListPermissions: []gax.CallOption{},
		UpdatePermission: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeletePermission: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		TransferOwnership: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultPermissionRESTCallOptions() *PermissionCallOptions {
	return &PermissionCallOptions{
		CreatePermission: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		GetPermission: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		ListPermissions: []gax.CallOption{},
		UpdatePermission: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		DeletePermission: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		TransferOwnership: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalPermissionClient is an interface that defines the methods available from Generative Language API.
type internalPermissionClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreatePermission(context.Context, *generativelanguagepb.CreatePermissionRequest, ...gax.CallOption) (*generativelanguagepb.Permission, error)
	GetPermission(context.Context, *generativelanguagepb.GetPermissionRequest, ...gax.CallOption) (*generativelanguagepb.Permission, error)
	ListPermissions(context.Context, *generativelanguagepb.ListPermissionsRequest, ...gax.CallOption) *PermissionIterator
	UpdatePermission(context.Context, *generativelanguagepb.UpdatePermissionRequest, ...gax.CallOption) (*generativelanguagepb.Permission, error)
	DeletePermission(context.Context, *generativelanguagepb.DeletePermissionRequest, ...gax.CallOption) error
	TransferOwnership(context.Context, *generativelanguagepb.TransferOwnershipRequest, ...gax.CallOption) (*generativelanguagepb.TransferOwnershipResponse, error)
}

// PermissionClient is a client for interacting with Generative Language API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides methods for managing permissions to PaLM API resources.
type PermissionClient struct {
	// The internal transport-dependent client.
	internalClient internalPermissionClient

	// The call options for this service.
	CallOptions *PermissionCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PermissionClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PermissionClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *PermissionClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreatePermission create a permission to a specific resource.
func (c *PermissionClient) CreatePermission(ctx context.Context, req *generativelanguagepb.CreatePermissionRequest, opts ...gax.CallOption) (*generativelanguagepb.Permission, error) {
	return c.internalClient.CreatePermission(ctx, req, opts...)
}

// GetPermission gets information about a specific Permission.
func (c *PermissionClient) GetPermission(ctx context.Context, req *generativelanguagepb.GetPermissionRequest, opts ...gax.CallOption) (*generativelanguagepb.Permission, error) {
	return c.internalClient.GetPermission(ctx, req, opts...)
}

// ListPermissions lists permissions for the specific resource.
func (c *PermissionClient) ListPermissions(ctx context.Context, req *generativelanguagepb.ListPermissionsRequest, opts ...gax.CallOption) *PermissionIterator {
	return c.internalClient.ListPermissions(ctx, req, opts...)
}

// UpdatePermission updates the permission.
func (c *PermissionClient) UpdatePermission(ctx context.Context, req *generativelanguagepb.UpdatePermissionRequest, opts ...gax.CallOption) (*generativelanguagepb.Permission, error) {
	return c.internalClient.UpdatePermission(ctx, req, opts...)
}

// DeletePermission deletes the permission.
func (c *PermissionClient) DeletePermission(ctx context.Context, req *generativelanguagepb.DeletePermissionRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeletePermission(ctx, req, opts...)
}

// TransferOwnership transfers ownership of the tuned model.
// This is the only way to change ownership of the tuned model.
// The current owner will be downgraded to writer role.
func (c *PermissionClient) TransferOwnership(ctx context.Context, req *generativelanguagepb.TransferOwnershipRequest, opts ...gax.CallOption) (*generativelanguagepb.TransferOwnershipResponse, error) {
	return c.internalClient.TransferOwnership(ctx, req, opts...)
}

// permissionGRPCClient is a client for interacting with Generative Language API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type permissionGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing PermissionClient
	CallOptions **PermissionCallOptions

	// The gRPC API client.
	permissionClient generativelanguagepb.PermissionServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewPermissionClient creates a new permission service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides methods for managing permissions to PaLM API resources.
func NewPermissionClient(ctx context.Context, opts ...option.ClientOption) (*PermissionClient, error) {
	clientOpts := defaultPermissionGRPCClientOptions()
	if newPermissionClientHook != nil {
		hookOpts, err := newPermissionClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PermissionClient{CallOptions: defaultPermissionCallOptions()}

	c := &permissionGRPCClient{
		connPool:         connPool,
		permissionClient: generativelanguagepb.NewPermissionServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *permissionGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *permissionGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *permissionGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type permissionRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing PermissionClient
	CallOptions **PermissionCallOptions
}

// NewPermissionRESTClient creates a new permission service rest client.
//
// Provides methods for managing permissions to PaLM API resources.
func NewPermissionRESTClient(ctx context.Context, opts ...option.ClientOption) (*PermissionClient, error) {
	clientOpts := append(defaultPermissionRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultPermissionRESTCallOptions()
	c := &permissionRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &PermissionClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultPermissionRESTClientOptions() []option.ClientOption {
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
func (c *permissionRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *permissionRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *permissionRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *permissionGRPCClient) CreatePermission(ctx context.Context, req *generativelanguagepb.CreatePermissionRequest, opts ...gax.CallOption) (*generativelanguagepb.Permission, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreatePermission[0:len((*c.CallOptions).CreatePermission):len((*c.CallOptions).CreatePermission)], opts...)
	var resp *generativelanguagepb.Permission
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.permissionClient.CreatePermission(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *permissionGRPCClient) GetPermission(ctx context.Context, req *generativelanguagepb.GetPermissionRequest, opts ...gax.CallOption) (*generativelanguagepb.Permission, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetPermission[0:len((*c.CallOptions).GetPermission):len((*c.CallOptions).GetPermission)], opts...)
	var resp *generativelanguagepb.Permission
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.permissionClient.GetPermission(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *permissionGRPCClient) ListPermissions(ctx context.Context, req *generativelanguagepb.ListPermissionsRequest, opts ...gax.CallOption) *PermissionIterator {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).ListPermissions[0:len((*c.CallOptions).ListPermissions):len((*c.CallOptions).ListPermissions)], opts...)
	it := &PermissionIterator{}
	req = proto.Clone(req).(*generativelanguagepb.ListPermissionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*generativelanguagepb.Permission, string, error) {
		resp := &generativelanguagepb.ListPermissionsResponse{}
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
			resp, err = c.permissionClient.ListPermissions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPermissions(), resp.GetNextPageToken(), nil
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

func (c *permissionGRPCClient) UpdatePermission(ctx context.Context, req *generativelanguagepb.UpdatePermissionRequest, opts ...gax.CallOption) (*generativelanguagepb.Permission, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "permission.name", url.QueryEscape(req.GetPermission().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdatePermission[0:len((*c.CallOptions).UpdatePermission):len((*c.CallOptions).UpdatePermission)], opts...)
	var resp *generativelanguagepb.Permission
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.permissionClient.UpdatePermission(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *permissionGRPCClient) DeletePermission(ctx context.Context, req *generativelanguagepb.DeletePermissionRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeletePermission[0:len((*c.CallOptions).DeletePermission):len((*c.CallOptions).DeletePermission)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.permissionClient.DeletePermission(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *permissionGRPCClient) TransferOwnership(ctx context.Context, req *generativelanguagepb.TransferOwnershipRequest, opts ...gax.CallOption) (*generativelanguagepb.TransferOwnershipResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).TransferOwnership[0:len((*c.CallOptions).TransferOwnership):len((*c.CallOptions).TransferOwnership)], opts...)
	var resp *generativelanguagepb.TransferOwnershipResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.permissionClient.TransferOwnership(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreatePermission create a permission to a specific resource.
func (c *permissionRESTClient) CreatePermission(ctx context.Context, req *generativelanguagepb.CreatePermissionRequest, opts ...gax.CallOption) (*generativelanguagepb.Permission, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetPermission()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta/%v/permissions", req.GetParent())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreatePermission[0:len((*c.CallOptions).CreatePermission):len((*c.CallOptions).CreatePermission)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.Permission{}
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

// GetPermission gets information about a specific Permission.
func (c *permissionRESTClient) GetPermission(ctx context.Context, req *generativelanguagepb.GetPermissionRequest, opts ...gax.CallOption) (*generativelanguagepb.Permission, error) {
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
	opts = append((*c.CallOptions).GetPermission[0:len((*c.CallOptions).GetPermission):len((*c.CallOptions).GetPermission)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.Permission{}
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

// ListPermissions lists permissions for the specific resource.
func (c *permissionRESTClient) ListPermissions(ctx context.Context, req *generativelanguagepb.ListPermissionsRequest, opts ...gax.CallOption) *PermissionIterator {
	it := &PermissionIterator{}
	req = proto.Clone(req).(*generativelanguagepb.ListPermissionsRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*generativelanguagepb.Permission, string, error) {
		resp := &generativelanguagepb.ListPermissionsResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1beta/%v/permissions", req.GetParent())

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
			return nil, "", e
		}
		it.Response = resp
		return resp.GetPermissions(), resp.GetNextPageToken(), nil
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

// UpdatePermission updates the permission.
func (c *permissionRESTClient) UpdatePermission(ctx context.Context, req *generativelanguagepb.UpdatePermissionRequest, opts ...gax.CallOption) (*generativelanguagepb.Permission, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetPermission()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta/%v", req.GetPermission().GetName())

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
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "permission.name", url.QueryEscape(req.GetPermission().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdatePermission[0:len((*c.CallOptions).UpdatePermission):len((*c.CallOptions).UpdatePermission)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.Permission{}
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

// DeletePermission deletes the permission.
func (c *permissionRESTClient) DeletePermission(ctx context.Context, req *generativelanguagepb.DeletePermissionRequest, opts ...gax.CallOption) error {
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

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// TransferOwnership transfers ownership of the tuned model.
// This is the only way to change ownership of the tuned model.
// The current owner will be downgraded to writer role.
func (c *permissionRESTClient) TransferOwnership(ctx context.Context, req *generativelanguagepb.TransferOwnershipRequest, opts ...gax.CallOption) (*generativelanguagepb.TransferOwnershipResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta/%v:transferOwnership", req.GetName())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).TransferOwnership[0:len((*c.CallOptions).TransferOwnership):len((*c.CallOptions).TransferOwnership)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &generativelanguagepb.TransferOwnershipResponse{}
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
