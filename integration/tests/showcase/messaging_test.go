// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package showcase_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	showcase "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1"
	pb "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Clients are initialized in main_test.go.
var (
	messagingGRPC *showcase.MessagingClient
	messagingREST *showcase.MessagingClient
)

func Test_Room_lifecycle(t *testing.T) {
	for _, client := range map[string]*showcase.MessagingClient{
		//"grpc": messagingGRPC,
		"rest": messagingREST,
	} {
		first, err := client.CreateRoom(
			context.Background(),
			&pb.CreateRoomRequest{Room: &pb.Room{DisplayName: "Living Room"}})
		if err != nil {
			t.Errorf("Create: unexpected err %+v", err)
		}

		deleted, err := client.CreateRoom(
			context.Background(),
			&pb.CreateRoomRequest{Room: &pb.Room{DisplayName: "Library"}})
		if err != nil {
			t.Errorf("Create: unexpected err %+v", err)
		}

		err = client.DeleteRoom(
			context.Background(),
			&pb.DeleteRoomRequest{Name: deleted.GetName()})
		if err != nil {
			t.Errorf("Delete: unexpected err %+v", err)
		}

		created, err := client.CreateRoom(
			context.Background(),
			&pb.CreateRoomRequest{Room: &pb.Room{DisplayName: "Weight Room"}})
		if err != nil {
			t.Errorf("Create: unexpected err %+v", err)
		}

		got, err := client.GetRoom(
			context.Background(),
			&pb.GetRoomRequest{Name: created.GetName()})
		if err != nil {
			t.Errorf("Get: unexpected err %+v", err)
		}

		if diff := cmp.Diff(created, got, cmp.Comparer(proto.Equal)); diff != "" {
			t.Errorf("client.GetRoom() got=-, want=+:%s", diff)
		}

		got.DisplayName = "Library"

		_, err = client.UpdateRoom(
			context.Background(),
			&pb.UpdateRoomRequest{Room: got, UpdateMask: nil})
		if err != nil {
			t.Errorf("Update: unexpected err %+v", err)
		}

		updated, err := client.GetRoom(
			context.Background(),
			&pb.GetRoomRequest{Name: got.GetName()})
		if err != nil {
			t.Errorf("Get: unexpected err %+v", err)
		}

		// ignore update time changed on updates.
		got.UpdateTime = updated.UpdateTime

		if diff := cmp.Diff(updated, got, cmp.Comparer(proto.Equal)); diff != "" {
			t.Errorf("client.UpdateRoom() got=-, want=+:%s", diff)
		}

		list := &pb.ListRoomsRequest{PageSize: 1}
		it := client.ListRooms(context.Background(), list)

		if mSize := it.PageInfo().MaxSize; mSize != int(list.PageSize) {
			t.Errorf("client.ListRooms(): it.PageInfo().MaxSize = %d, want %d", mSize, list.PageSize)
		}

		listed, err := it.Next()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(listed, first, cmp.Comparer(proto.Equal)); diff != "" {
			t.Errorf("client.ListRooms() got=-, want=+:%s", diff)
		}
	}
}

func Test_CreateRoom_invalid(t *testing.T) {
	_, err := messagingGRPC.CreateRoom(
		context.Background(),
		&pb.CreateRoomRequest{Room: &pb.Room{DisplayName: ""}},
	)

	st, _ := status.FromError(err)
	if st.Code() != codes.InvalidArgument {
		t.Errorf("client.CreateRoom() Want error code %d got %d", codes.InvalidArgument, st.Code())
	}
}
