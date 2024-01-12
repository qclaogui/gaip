// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

//go:build requires_docker

package project_test

import (
	"context"
	"testing"

	project "github.com/qclaogui/gaip/genproto/project/apiv1"
	pb "github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
)

// Clients are initialized in main_test.go.
var (
	messagingGRPC *project.MessagingClient
	messagingREST *project.MessagingClient
)

func TestRoomCRUD(t *testing.T) {
	ctx := context.Background()

	for _, client := range map[string]*project.MessagingClient{
		//"grpc": messagingGRPC,
		"rest": messagingREST,
	} {
		// Create RoomRequest
		create := &pb.CreateRoomRequest{
			Room: &pb.Room{
				DisplayName: "Codelab",
				Description: "it is take time to get good at coding",
			},
		}

		room, err := client.CreateRoom(ctx, create)
		if err != nil {
			t.Fatal(err)
		}

		want := create.GetRoom()
		if room.GetName() == "" {
			t.Errorf("CreateRoom().Name was unexpectedly empty")
		}

		if room.GetDisplayName() != want.GetDisplayName() {
			t.Errorf("CreateRoom().DisplayName = %q, want = %q", room.GetDisplayName(), want.GetDisplayName())
		}

		if room.GetCreateTime() == nil {
			t.Errorf("CreateRoom().CreateTime was unexpectedly empty")
		}
		if room.GetUpdateTime() == nil {
			t.Errorf("CreateRoom().UpdateTime was unexpectedly empty")
		}

		// List UsersRequest

	}
}
