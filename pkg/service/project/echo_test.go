// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"
	"encoding/base64"
	"strings"
	"testing"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb"
	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestWait_pending(t *testing.T) {
	now := time.Unix(1, 0)
	endTime := time.Unix(2, 0)

	ttl := endTime.Sub(now)
	nowF := func() time.Time { return time.Unix(1, 0) }

	endTimeProto := timestamppb.New(endTime)

	tests := []*projectpb.WaitRequest{
		{
			End: &projectpb.WaitRequest_EndTime{
				EndTime: endTimeProto,
			},
		},
		{
			End: &projectpb.WaitRequest_Ttl{
				Ttl: durationpb.New(ttl),
			},
		},
	}

	for _, req := range tests {
		waiter := &echoServiceImpl{nowF: nowF}

		op, _ := waiter.Wait(context.Background(), req)
		if op.Done {
			t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
				"✘got: %v\n\x1b[92m"+
				"want: %v\x1b[39m", "done=false", "done=true")
		}

		checkName(t, req, op)

		if op.Metadata == nil {
			t.Errorf("Wait() for %q expected metadata, got nil", req)
		}

		meta := &projectpb.WaitMetadata{}
		_ = anypb.UnmarshalTo(op.Metadata, meta, proto.UnmarshalOptions{})
		if !proto.Equal(endTimeProto, meta.EndTime) {
			t.Errorf(
				"Wait for %q expected metadata with Endtime=%q, got %q",
				req,
				endTimeProto,
				meta.EndTime)
		}

	}
}

func checkName(t *testing.T, req *projectpb.WaitRequest, op *longrunningpb.Operation) {
	if !strings.HasPrefix(op.Name, "operations/qclaogui.project.v1.EchoService/Wait/") {
		t.Errorf("\nOops 🔥\x1b[91m Failed asserting that\x1b[39m\n"+
			"✘got: %v\n\x1b[92m"+
			"want: %v\x1b[39m", op.Name, "op.Name prefex 'operations/qclaogui.project.v1.EchoService/Wait/'")
	}

	nameProto := &projectpb.WaitRequest{}

	encodedBytes := strings.TrimPrefix(
		op.Name,
		"operations/qclaogui.project.v1.EchoService/Wait/")
	bytes, _ := base64.StdEncoding.DecodeString(encodedBytes)

	_ = proto.Unmarshal(bytes, nameProto)
	if !proto.Equal(nameProto, req) {
		t.Errorf(
			"Wait() for %q expected unmarshalled name=%q, got name=%q",
			req,
			req,
			nameProto)
	}
}

func TestWait_success(t *testing.T) {
	nowF := func() time.Time { return time.Unix(3, 0) }
	endTime := timestamppb.New(time.Unix(2, 0))
	success := &projectpb.WaitResponse{Content: "Hello World!"}
	req := &projectpb.WaitRequest{
		End:      &projectpb.WaitRequest_EndTime{EndTime: endTime},
		Response: &projectpb.WaitRequest_Success{Success: success},
	}

	waiter := &echoServiceImpl{nowF: nowF}
	op, _ := waiter.Wait(context.Background(), req)

	checkName(t, req, op)

	if !op.Done {
		t.Errorf("Wait() for %q expected done=true got done=false", req)
	}

	if op.Metadata != nil {
		t.Errorf("Wait() for %q expected nil metadata, got %q", req, op.Metadata)
	}

	if op.GetError() != nil {
		t.Errorf("Wait() expected op.Error=nil, got %q", op.GetError())
	}

	if op.GetResponse() == nil {
		t.Error("Wait() expected op.Response!=nil")
	}

	resp := &projectpb.WaitResponse{}
	_ = anypb.UnmarshalTo(op.GetResponse(), resp, proto.UnmarshalOptions{})
	if !proto.Equal(resp, success) {
		t.Errorf("Wait() expected op.GetResponse()=%q, got %q", success, resp)
	}
}

func TestWait_error(t *testing.T) {
	nowF := func() time.Time { return time.Unix(3, 0) }
	endTime := timestamppb.New(time.Unix(2, 0))
	expErr := &status.Status{Code: int32(1), Message: "Error!"}
	req := &projectpb.WaitRequest{
		End: &projectpb.WaitRequest_EndTime{
			EndTime: endTime,
		},
		Response: &projectpb.WaitRequest_Error{Error: expErr},
	}

	waiter := &echoServiceImpl{nowF: nowF}
	op, _ := waiter.Wait(context.Background(), req)

	checkName(t, req, op)

	if !op.Done {
		t.Errorf("Wait() for %q expected done=true got done=false", req)
	}

	if op.Metadata != nil {
		t.Errorf("Wait() for %q expected nil metadata, got %q", req, op.Metadata)
	}

	if op.GetResponse() != nil {
		t.Errorf("Wait() expected op.Response=nil, got %q", op.GetResponse())
	}

	if !proto.Equal(expErr, op.GetError()) {
		t.Errorf("Wait() expected op.Error=%q, got %q", expErr, op.GetError())
	}
}
