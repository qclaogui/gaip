// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/qclaogui/gaip/genproto/a2a/apiv1/a2apb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewA2A() (a2apb.A2AServiceServer, error) {
	m := &a2aImpl{
		taskStore:   make(map[string]*a2apb.Task),
		taskHistory: make(map[string][]*a2apb.Message),
	}
	return m, nil
}

// a2aImpl is used to implement a2apb.A2AServiceServer.
type a2aImpl struct {
	a2apb.UnimplementedA2AServiceServer
	// agentCard a2apb.AgentCard

	taskStore   map[string]*a2apb.Task
	taskHistory map[string][]*a2apb.Message
	mu          sync.RWMutex
}

func (s *a2aImpl) SendMessage(_ context.Context, req *a2apb.SendMessageRequest) (*a2apb.SendMessageResponse, error) {
	incomingMessage := req.GetRequest()
	if incomingMessage.GetMessageId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "message.messageId is required.")
	}

	// Default to blocking behavior if 'blocking' is not explicitly false.
	isBlocking := req.GetConfiguration().GetBlocking()
	taskId := incomingMessage.GetTaskId()
	_ = isBlocking
	// incomingMessage would contain taskId, if a task already exists.
	if taskId != "" {
		task := s.taskStore[taskId]

		return &a2apb.SendMessageResponse{
			Payload: &a2apb.SendMessageResponse_Task{
				Task: task,
			},
		}, nil

	}

	// Create new task
	task := &a2apb.Task{
		Id: uuid.NewString(),
		Status: &a2apb.TaskStatus{
			State: a2apb.TaskState_TASK_STATE_WORKING,
		},
	}

	// Process task
	task.Status.State = a2apb.TaskState_TASK_STATE_COMPLETED

	// Store task and history
	s.taskStore[task.Id] = task
	s.taskHistory[task.Id] = append(s.taskHistory[task.Id], req.Request)

	return &a2apb.SendMessageResponse{
		Payload: &a2apb.SendMessageResponse_Task{
			Task: task,
		},
	}, nil
}

func (s *a2aImpl) SendStreamingMessage(*a2apb.SendMessageRequest, a2apb.A2AService_SendStreamingMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendStreamingMessage not implemented")
}

func (s *a2aImpl) GetTask(_ context.Context, req *a2apb.GetTaskRequest) (*a2apb.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.taskStore[req.GetName()]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "A task with name %s not found.", req.GetName())
	}

	return task, nil
}

func (s *a2aImpl) CancelTask(_ context.Context, req *a2apb.CancelTaskRequest) (*a2apb.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.taskStore[req.GetName()]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "A task with name %s not found.", req.GetName())
	}

	// Update task status to canceled
	task.Status.State = a2apb.TaskState_TASK_STATE_CANCELLED
	s.taskStore[req.GetName()] = task

	return task, nil
}

func (s *a2aImpl) TaskSubscription(*a2apb.TaskSubscriptionRequest, a2apb.A2AService_SendStreamingMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method TaskSubscription not implemented")
}

func (s *a2aImpl) CreateTaskPushNotificationConfig(context.Context, *a2apb.CreateTaskPushNotificationConfigRequest) (*a2apb.TaskPushNotificationConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskPushNotificationConfig not implemented")
}

func (s *a2aImpl) GetTaskPushNotificationConfig(context.Context, *a2apb.GetTaskPushNotificationConfigRequest) (*a2apb.TaskPushNotificationConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskPushNotificationConfig not implemented")
}

func (s *a2aImpl) ListTaskPushNotificationConfig(context.Context, *a2apb.ListTaskPushNotificationConfigRequest) (*a2apb.ListTaskPushNotificationConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTaskPushNotificationConfig not implemented")
}

func (s *a2aImpl) DeleteTaskPushNotificationConfig(context.Context, *a2apb.DeleteTaskPushNotificationConfigRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTaskPushNotificationConfig not implemented")
}

func (s *a2aImpl) GetAgentCard(context.Context, *a2apb.GetAgentCardRequest) (*a2apb.AgentCard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgentCard not implemented")
}
