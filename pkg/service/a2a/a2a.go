// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package a2a

import (
	"context"

	pb "github.com/qclaogui/gaip/genproto/a2a/apiv1/a2apb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	return s.repo.SendMessage(ctx, req)
}

func (s *Server) SendStreamingMessage(req *pb.SendMessageRequest, steam pb.A2AService_SendStreamingMessageServer) error {
	return s.repo.SendStreamingMessage(req, steam)
}

func (s *Server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.Task, error) {
	return s.repo.GetTask(ctx, req)
}

func (s *Server) CancelTask(ctx context.Context, req *pb.CancelTaskRequest) (*pb.Task, error) {
	return s.repo.CancelTask(ctx, req)
}

func (s *Server) TaskSubscription(req *pb.TaskSubscriptionRequest, steam pb.A2AService_SendStreamingMessageServer) error {
	return s.repo.TaskSubscription(req, steam)
}

func (s *Server) CreateTaskPushNotificationConfig(ctx context.Context, req *pb.CreateTaskPushNotificationConfigRequest) (*pb.TaskPushNotificationConfig, error) {
	return s.repo.CreateTaskPushNotificationConfig(ctx, req)
}

func (s *Server) GetTaskPushNotificationConfig(ctx context.Context, req *pb.GetTaskPushNotificationConfigRequest) (*pb.TaskPushNotificationConfig, error) {
	return s.repo.GetTaskPushNotificationConfig(ctx, req)
}

func (s *Server) ListTaskPushNotificationConfig(ctx context.Context, req *pb.ListTaskPushNotificationConfigRequest) (*pb.ListTaskPushNotificationConfigResponse, error) {
	return s.repo.ListTaskPushNotificationConfig(ctx, req)
}

func (s *Server) DeleteTaskPushNotificationConfig(ctx context.Context, req *pb.DeleteTaskPushNotificationConfigRequest) (*emptypb.Empty, error) {
	return s.repo.DeleteTaskPushNotificationConfig(ctx, req)
}

func (s *Server) GetAgentCard(ctx context.Context, req *pb.GetAgentCardRequest) (*pb.AgentCard, error) {
	return s.repo.GetAgentCard(ctx, req)
}
