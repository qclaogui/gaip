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

package project

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	projectpb "github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var newIdentityClientHook clientHook

// IdentityCallOptions contains the retry settings for each method of IdentityClient.
type IdentityCallOptions struct {
	CreateUser []gax.CallOption
	GetUser    []gax.CallOption
	ListUsers  []gax.CallOption
	UpdateUser []gax.CallOption
	DeleteUser []gax.CallOption
}

func defaultIdentityGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("localhost:9095"),
		internaloption.WithDefaultMTLSEndpoint("localhost:9095"),
		internaloption.WithDefaultAudience("https://localhost/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultIdentityCallOptions() *IdentityCallOptions {
	return &IdentityCallOptions{
		CreateUser: []gax.CallOption{},
		GetUser:    []gax.CallOption{},
		ListUsers:  []gax.CallOption{},
		UpdateUser: []gax.CallOption{},
		DeleteUser: []gax.CallOption{},
	}
}

func defaultIdentityRESTCallOptions() *IdentityCallOptions {
	return &IdentityCallOptions{
		CreateUser: []gax.CallOption{},
		GetUser:    []gax.CallOption{},
		ListUsers:  []gax.CallOption{},
		UpdateUser: []gax.CallOption{},
		DeleteUser: []gax.CallOption{},
	}
}

// internalIdentityClient is an interface that defines the methods available from .
type internalIdentityClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateUser(context.Context, *projectpb.CreateUserRequest, ...gax.CallOption) (*projectpb.User, error)
	GetUser(context.Context, *projectpb.GetUserRequest, ...gax.CallOption) (*projectpb.User, error)
	ListUsers(context.Context, *projectpb.ListUsersRequest, ...gax.CallOption) *UserIterator
	UpdateUser(context.Context, *projectpb.UpdateUserRequest, ...gax.CallOption) (*projectpb.User, error)
	DeleteUser(context.Context, *projectpb.DeleteUserRequest, ...gax.CallOption) error
}

// IdentityClient is a client for interacting with .
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A simple identity service.
type IdentityClient struct {
	// The internal transport-dependent client.
	internalClient internalIdentityClient

	// The call options for this service.
	CallOptions *IdentityCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IdentityClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IdentityClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *IdentityClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateUser creates a user.
func (c *IdentityClient) CreateUser(ctx context.Context, req *projectpb.CreateUserRequest, opts ...gax.CallOption) (*projectpb.User, error) {
	return c.internalClient.CreateUser(ctx, req, opts...)
}

// GetUser retrieves the User with the given uri.
func (c *IdentityClient) GetUser(ctx context.Context, req *projectpb.GetUserRequest, opts ...gax.CallOption) (*projectpb.User, error) {
	return c.internalClient.GetUser(ctx, req, opts...)
}

// ListUsers lists all users.
func (c *IdentityClient) ListUsers(ctx context.Context, req *projectpb.ListUsersRequest, opts ...gax.CallOption) *UserIterator {
	return c.internalClient.ListUsers(ctx, req, opts...)
}

// UpdateUser updates a user.
func (c *IdentityClient) UpdateUser(ctx context.Context, req *projectpb.UpdateUserRequest, opts ...gax.CallOption) (*projectpb.User, error) {
	return c.internalClient.UpdateUser(ctx, req, opts...)
}

// DeleteUser deletes a user, their profile, and all of their authored messages.
func (c *IdentityClient) DeleteUser(ctx context.Context, req *projectpb.DeleteUserRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteUser(ctx, req, opts...)
}

// identityGRPCClient is a client for interacting with  over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type identityGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing IdentityClient
	CallOptions **IdentityCallOptions

	// The gRPC API client.
	identityClient projectpb.IdentityServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewIdentityClient creates a new identity service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A simple identity service.
func NewIdentityClient(ctx context.Context, opts ...option.ClientOption) (*IdentityClient, error) {
	clientOpts := defaultIdentityGRPCClientOptions()
	if newIdentityClientHook != nil {
		hookOpts, err := newIdentityClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := IdentityClient{CallOptions: defaultIdentityCallOptions()}

	c := &identityGRPCClient{
		connPool:       connPool,
		identityClient: projectpb.NewIdentityServiceClient(connPool),
		CallOptions:    &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *identityGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *identityGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *identityGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type identityRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing IdentityClient
	CallOptions **IdentityCallOptions
}

// NewIdentityRESTClient creates a new identity service rest client.
//
// A simple identity service.
func NewIdentityRESTClient(ctx context.Context, opts ...option.ClientOption) (*IdentityClient, error) {
	clientOpts := append(defaultIdentityRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultIdentityRESTCallOptions()
	c := &identityRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &IdentityClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultIdentityRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://localhost:9095"),
		internaloption.WithDefaultMTLSEndpoint("https://localhost:9095"),
		internaloption.WithDefaultAudience("https://localhost/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *identityRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *identityRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *identityRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *identityGRPCClient) CreateUser(ctx context.Context, req *projectpb.CreateUserRequest, opts ...gax.CallOption) (*projectpb.User, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).CreateUser[0:len((*c.CallOptions).CreateUser):len((*c.CallOptions).CreateUser)], opts...)
	var resp *projectpb.User
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityClient.CreateUser(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *identityGRPCClient) GetUser(ctx context.Context, req *projectpb.GetUserRequest, opts ...gax.CallOption) (*projectpb.User, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetUser[0:len((*c.CallOptions).GetUser):len((*c.CallOptions).GetUser)], opts...)
	var resp *projectpb.User
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityClient.GetUser(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *identityGRPCClient) ListUsers(ctx context.Context, req *projectpb.ListUsersRequest, opts ...gax.CallOption) *UserIterator {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListUsers[0:len((*c.CallOptions).ListUsers):len((*c.CallOptions).ListUsers)], opts...)
	it := &UserIterator{}
	req = proto.Clone(req).(*projectpb.ListUsersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*projectpb.User, string, error) {
		resp := &projectpb.ListUsersResponse{}
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
			resp, err = c.identityClient.ListUsers(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetUsers(), resp.GetNextPageToken(), nil
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

func (c *identityGRPCClient) UpdateUser(ctx context.Context, req *projectpb.UpdateUserRequest, opts ...gax.CallOption) (*projectpb.User, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "user.name", url.QueryEscape(req.GetUser().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateUser[0:len((*c.CallOptions).UpdateUser):len((*c.CallOptions).UpdateUser)], opts...)
	var resp *projectpb.User
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.identityClient.UpdateUser(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *identityGRPCClient) DeleteUser(ctx context.Context, req *projectpb.DeleteUserRequest, opts ...gax.CallOption) error {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteUser[0:len((*c.CallOptions).DeleteUser):len((*c.CallOptions).DeleteUser)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.identityClient.DeleteUser(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// CreateUser creates a user.
func (c *identityRESTClient) CreateUser(ctx context.Context, req *projectpb.CreateUserRequest, opts ...gax.CallOption) (*projectpb.User, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetUser()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/users")

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateUser[0:len((*c.CallOptions).CreateUser):len((*c.CallOptions).CreateUser)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &projectpb.User{}
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

// GetUser retrieves the User with the given uri.
func (c *identityRESTClient) GetUser(ctx context.Context, req *projectpb.GetUserRequest, opts ...gax.CallOption) (*projectpb.User, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetUser[0:len((*c.CallOptions).GetUser):len((*c.CallOptions).GetUser)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &projectpb.User{}
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

// ListUsers lists all users.
func (c *identityRESTClient) ListUsers(ctx context.Context, req *projectpb.ListUsersRequest, opts ...gax.CallOption) *UserIterator {
	it := &UserIterator{}
	req = proto.Clone(req).(*projectpb.ListUsersRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*projectpb.User, string, error) {
		resp := &projectpb.ListUsersResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/users")

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
		return resp.GetUsers(), resp.GetNextPageToken(), nil
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

// UpdateUser updates a user.
func (c *identityRESTClient) UpdateUser(ctx context.Context, req *projectpb.UpdateUserRequest, opts ...gax.CallOption) (*projectpb.User, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetUser()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetUser().GetName())

	params := url.Values{}
	if req.GetUpdateMask() != nil {
		updateMask, err := protojson.Marshal(req.GetUpdateMask())
		if err != nil {
			return nil, err
		}
		params.Add("updateMask", string(updateMask[1:len(updateMask)-1]))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "user.name", url.QueryEscape(req.GetUser().GetName()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateUser[0:len((*c.CallOptions).UpdateUser):len((*c.CallOptions).UpdateUser)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &projectpb.User{}
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

// DeleteUser deletes a user, their profile, and all of their authored messages.
func (c *identityRESTClient) DeleteUser(ctx context.Context, req *projectpb.DeleteUserRequest, opts ...gax.CallOption) error {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1/%v", req.GetName())

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
