// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/golang-api-server/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EchoService interface {
	service.Backend

	projectpb.EchoServiceServer
}

type echoServiceImpl struct {
	projectpb.UnimplementedEchoServiceServer

	nowF func() time.Time
}

func NewEchoService() (EchoService, error) {
	s := &echoServiceImpl{}
	return s, nil
}

func (s *echoServiceImpl) RegisterGRPC(grpcServer *grpc.Server) {
	grpcServer.RegisterService(&projectpb.EchoService_ServiceDesc, s)
}

func (s *echoServiceImpl) Wait(_ context.Context, req *projectpb.WaitRequest) (*longrunningpb.Operation, error) {
	endTime := time.Unix(0, 0).UTC()
	if ttl := req.GetTtl(); ttl != nil {
		endTime = s.nowF().Add(ttl.AsDuration())
	}
	if end := req.GetEndTime(); end != nil {
		endTime = end.AsTime()
	}
	endTimeProto := timestamppb.New(endTime)
	req.End = &projectpb.WaitRequest_EndTime{
		EndTime: endTimeProto,
	}

	done := s.nowF().After(endTime)
	reqBytes, _ := proto.Marshal(req)

	name := fmt.Sprintf(
		"operations/qclaogui.project.v1.EchoService/Wait/%s",
		base64.StdEncoding.EncodeToString(reqBytes))

	answer := &longrunningpb.Operation{
		Name: name,
		Done: done,
	}

	if done && (req.GetError() != nil) {
		answer.Result = &longrunningpb.Operation_Error{Error: req.GetError()}
	}

	if done && (req.GetSuccess() != nil) {
		resp, _ := anypb.New(req.GetSuccess())
		answer.Result = &longrunningpb.Operation_Response{Response: resp}
	}

	if !done {
		meta, _ := anypb.New(&projectpb.WaitMetadata{EndTime: endTimeProto})
		answer.Metadata = meta
	}

	return answer, nil

}
