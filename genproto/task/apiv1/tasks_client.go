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

package task

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	taskpb "github.com/qclaogui/gaip/genproto/task/apiv1/taskpb"
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

var newTasksClientHook clientHook

// TasksCallOptions contains the retry settings for each method of TasksClient.
type TasksCallOptions struct {
	CreateTask   []gax.CallOption
	DeleteTask   []gax.CallOption
	UndeleteTask []gax.CallOption
	UpdateTask   []gax.CallOption
	GetTask      []gax.CallOption
	ListTasks    []gax.CallOption
}

func defaultTasksGRPCClientOptions() []option.ClientOption {
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

func defaultTasksCallOptions() *TasksCallOptions {
	return &TasksCallOptions{
		CreateTask: []gax.CallOption{
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
		DeleteTask: []gax.CallOption{
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
		UndeleteTask: []gax.CallOption{
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
		UpdateTask: []gax.CallOption{
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
		GetTask: []gax.CallOption{
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
		ListTasks: []gax.CallOption{
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
	}
}

func defaultTasksRESTCallOptions() *TasksCallOptions {
	return &TasksCallOptions{
		CreateTask: []gax.CallOption{
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
		DeleteTask: []gax.CallOption{
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
		UndeleteTask: []gax.CallOption{
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
		UpdateTask: []gax.CallOption{
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
		GetTask: []gax.CallOption{
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
		ListTasks: []gax.CallOption{
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

// internalTasksClient is an interface that defines the methods available from .
type internalTasksClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateTask(context.Context, *taskpb.CreateTaskRequest, ...gax.CallOption) (*taskpb.Task, error)
	DeleteTask(context.Context, *taskpb.DeleteTaskRequest, ...gax.CallOption) (*taskpb.Task, error)
	UndeleteTask(context.Context, *taskpb.UndeleteTaskRequest, ...gax.CallOption) (*taskpb.Task, error)
	UpdateTask(context.Context, *taskpb.UpdateTaskRequest, ...gax.CallOption) (*taskpb.Task, error)
	GetTask(context.Context, *taskpb.GetTaskRequest, ...gax.CallOption) (*taskpb.Task, error)
	ListTasks(context.Context, *taskpb.ListTasksRequest, ...gax.CallOption) *TaskIterator
}

// TasksClient is a client for interacting with .
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// TasksService holds the methods to persist, modify and remove Tasks.
type TasksClient struct {
	// The internal transport-dependent client.
	internalClient internalTasksClient

	// The call options for this service.
	CallOptions *TasksCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TasksClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TasksClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TasksClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateTask createTask creates a Task.
func (c *TasksClient) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	return c.internalClient.CreateTask(ctx, req, opts...)
}

func (c *TasksClient) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	return c.internalClient.DeleteTask(ctx, req, opts...)
}

func (c *TasksClient) UndeleteTask(ctx context.Context, req *taskpb.UndeleteTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	return c.internalClient.UndeleteTask(ctx, req, opts...)
}

func (c *TasksClient) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	return c.internalClient.UpdateTask(ctx, req, opts...)
}

// GetTask getTask returns a Task.
func (c *TasksClient) GetTask(ctx context.Context, req *taskpb.GetTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	return c.internalClient.GetTask(ctx, req, opts...)
}

// ListTasks listTasks returns a list of Tasks.
func (c *TasksClient) ListTasks(ctx context.Context, req *taskpb.ListTasksRequest, opts ...gax.CallOption) *TaskIterator {
	return c.internalClient.ListTasks(ctx, req, opts...)
}

// tasksGRPCClient is a client for interacting with  over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type tasksGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TasksClient
	CallOptions **TasksCallOptions

	// The gRPC API client.
	tasksClient taskpb.TasksServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string

	logger *slog.Logger
}

// NewTasksClient creates a new tasks service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// TasksService holds the methods to persist, modify and remove Tasks.
func NewTasksClient(ctx context.Context, opts ...option.ClientOption) (*TasksClient, error) {
	clientOpts := defaultTasksGRPCClientOptions()
	if newTasksClientHook != nil {
		hookOpts, err := newTasksClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TasksClient{CallOptions: defaultTasksCallOptions()}

	c := &tasksGRPCClient{
		connPool:    connPool,
		tasksClient: taskpb.NewTasksServiceClient(connPool),
		CallOptions: &client.CallOptions,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *tasksGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *tasksGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
		"x-goog-api-version", "v1_20240506",
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *tasksGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type tasksRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing TasksClient
	CallOptions **TasksCallOptions

	logger *slog.Logger
}

// NewTasksRESTClient creates a new tasks service rest client.
//
// TasksService holds the methods to persist, modify and remove Tasks.
func NewTasksRESTClient(ctx context.Context, opts ...option.ClientOption) (*TasksClient, error) {
	clientOpts := append(defaultTasksRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultTasksRESTCallOptions()
	c := &tasksRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
		logger:      internaloption.GetLogger(opts),
	}
	c.setGoogleClientInfo()

	return &TasksClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultTasksRESTClientOptions() []option.ClientOption {
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
func (c *tasksRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{
		"x-goog-api-client", gax.XGoogHeader(kv...),
		"x-goog-api-version", "v1_20240506",
	}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *tasksRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *tasksRESTClient) Connection() *grpc.ClientConn {
	return nil
}

func (c *tasksGRPCClient) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).CreateTask[0:len((*c.CallOptions).CreateTask):len((*c.CallOptions).CreateTask)], opts...)
	var resp *taskpb.Task
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tasksClient.CreateTask, req, settings.GRPC, c.logger, "CreateTask")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tasksGRPCClient) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "id", req.GetId())}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).DeleteTask[0:len((*c.CallOptions).DeleteTask):len((*c.CallOptions).DeleteTask)], opts...)
	var resp *taskpb.Task
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tasksClient.DeleteTask, req, settings.GRPC, c.logger, "DeleteTask")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tasksGRPCClient) UndeleteTask(ctx context.Context, req *taskpb.UndeleteTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).UndeleteTask[0:len((*c.CallOptions).UndeleteTask):len((*c.CallOptions).UndeleteTask)], opts...)
	var resp *taskpb.Task
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tasksClient.UndeleteTask, req, settings.GRPC, c.logger, "UndeleteTask")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tasksGRPCClient) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "task.id", req.GetTask().GetId())}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).UpdateTask[0:len((*c.CallOptions).UpdateTask):len((*c.CallOptions).UpdateTask)], opts...)
	var resp *taskpb.Task
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tasksClient.UpdateTask, req, settings.GRPC, c.logger, "UpdateTask")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tasksGRPCClient) GetTask(ctx context.Context, req *taskpb.GetTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "id", req.GetId())}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetTask[0:len((*c.CallOptions).GetTask):len((*c.CallOptions).GetTask)], opts...)
	var resp *taskpb.Task
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = executeRPC(ctx, c.tasksClient.GetTask, req, settings.GRPC, c.logger, "GetTask")
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tasksGRPCClient) ListTasks(ctx context.Context, req *taskpb.ListTasksRequest, opts ...gax.CallOption) *TaskIterator {
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
	opts = append((*c.CallOptions).ListTasks[0:len((*c.CallOptions).ListTasks):len((*c.CallOptions).ListTasks)], opts...)
	it := &TaskIterator{}
	req = proto.Clone(req).(*taskpb.ListTasksRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*taskpb.Task, string, error) {
		resp := &taskpb.ListTasksResponse{}
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
			resp, err = executeRPC(ctx, c.tasksClient.ListTasks, req, settings.GRPC, c.logger, "ListTasks")
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTasks(), resp.GetNextPageToken(), nil
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

// CreateTask createTask creates a Task.
func (c *tasksRESTClient) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetTask()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/tasks")

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateTask[0:len((*c.CallOptions).CreateTask):len((*c.CallOptions).CreateTask)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &taskpb.Task{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "CreateTask")
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

func (c *tasksRESTClient) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/tasks/%v", req.GetId())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "id", req.GetId())}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).DeleteTask[0:len((*c.CallOptions).DeleteTask):len((*c.CallOptions).DeleteTask)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &taskpb.Task{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "DeleteTask")
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

func (c *tasksRESTClient) UndeleteTask(ctx context.Context, req *taskpb.UndeleteTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/tasks:undelete")

	// Build HTTP headers from client and context metadata.
	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UndeleteTask[0:len((*c.CallOptions).UndeleteTask):len((*c.CallOptions).UndeleteTask)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &taskpb.Task{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UndeleteTask")
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

func (c *tasksRESTClient) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetTask()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/tasks/%v", req.GetTask().GetId())

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
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "task.id", req.GetTask().GetId())}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).UpdateTask[0:len((*c.CallOptions).UpdateTask):len((*c.CallOptions).UpdateTask)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &taskpb.Task{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, jsonReq, "UpdateTask")
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

// GetTask getTask returns a Task.
func (c *tasksRESTClient) GetTask(ctx context.Context, req *taskpb.GetTaskRequest, opts ...gax.CallOption) (*taskpb.Task, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/tasks/%v", req.GetId())

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v", "id", req.GetId())}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetTask[0:len((*c.CallOptions).GetTask):len((*c.CallOptions).GetTask)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &taskpb.Task{}
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

		buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "GetTask")
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

// ListTasks listTasks returns a list of Tasks.
func (c *tasksRESTClient) ListTasks(ctx context.Context, req *taskpb.ListTasksRequest, opts ...gax.CallOption) *TaskIterator {
	it := &TaskIterator{}
	req = proto.Clone(req).(*taskpb.ListTasksRequest)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	it.InternalFetch = func(pageSize int, pageToken string) ([]*taskpb.Task, string, error) {
		resp := &taskpb.ListTasksResponse{}
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
		baseUrl.Path += fmt.Sprintf("/v1/tasks")

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

			buf, err := executeHTTPRequest(ctx, c.httpClient, httpReq, c.logger, nil, "ListTasks")
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
		return resp.GetTasks(), resp.GetNextPageToken(), nil
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
