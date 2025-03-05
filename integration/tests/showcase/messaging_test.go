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
	for typ, client := range map[string]*showcase.MessagingClient{
		//"grpc": messagingGRPC,
		"rest": messagingREST,
	} {
		t.Run(typ, func(t *testing.T) {
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
		})
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

func newUser(name string) (*pb.User, error) {
	return identityGRPC.CreateUser(context.Background(), &pb.CreateUserRequest{
		User: &pb.User{
			DisplayName: name,
			Email:       name + "@example.com",
			Nickname:    proto.String(name),
			HeightFeet:  proto.Float64(6.2),
		},
	})
}

func Test_Blurb_lifecycle(t *testing.T) {
	usr, errU := newUser("Test_Blurb_lifecycle")
	if errU != nil {
		t.Fatalf("identityGRPC.CreateUser() failed: %v", errU)
	}

	parent := usr.GetName() + "/profile"

	for typ, client := range map[string]*showcase.MessagingClient{
		"grpc": messagingGRPC,
		//"rest": messagingREST,
	} {
		t.Run(typ, func(t *testing.T) {
			first, err := client.CreateBlurb(
				context.Background(),
				&pb.CreateBlurbRequest{
					Parent: parent,
					Blurb: &pb.Blurb{
						User:    usr.GetName(),
						Content: &pb.Blurb_Text{Text: "hello world"},
					},
				})
			if err != nil {
				t.Errorf("Create: unexpected err %+v", err)
			}

			deleted, err := client.CreateBlurb(
				context.Background(),
				&pb.CreateBlurbRequest{
					Parent: parent,
					Blurb: &pb.Blurb{
						User:    usr.GetName(),
						Content: &pb.Blurb_Text{Text: "bark"},
					},
				})
			if err != nil {
				t.Errorf("Create: unexpected err %+v", err)
			}

			if err = client.DeleteBlurb(
				context.Background(),
				&pb.DeleteBlurbRequest{Name: deleted.GetName()}); err != nil {
				t.Errorf("Delete: unexpected err %+v", err)
			}

			created, err := client.CreateBlurb(
				context.Background(),
				&pb.CreateBlurbRequest{
					Parent: parent,
					Blurb: &pb.Blurb{
						User:    usr.GetName(),
						Content: &pb.Blurb_Text{Text: "meow"},
					},
				})
			if err != nil {
				t.Errorf("Create: unexpected err %+v", err)
			}

			got, err := client.GetBlurb(
				context.Background(),
				&pb.GetBlurbRequest{
					Name: created.GetName(),
				})
			if err != nil {
				t.Errorf("Get: unexpected err %+v", err)
			}

			if diff := cmp.Diff(created, got, cmp.Comparer(proto.Equal)); diff != "" {
				t.Errorf("client.GetBlurb() got=-, want=+:%s", diff)
			}

			got.Content = &pb.Blurb_Text{Text: "purrr"}
			_, err = client.UpdateBlurb(
				context.Background(),
				&pb.UpdateBlurbRequest{Blurb: got, UpdateMask: nil})
			if err != nil {
				t.Errorf("Update: unexpected err %+v", err)
			}

			updated, err := client.GetBlurb(
				context.Background(),
				&pb.GetBlurbRequest{Name: got.GetName()})
			if err != nil {
				t.Errorf("Get: unexpected err %+v", err)
			}

			// ignore update time changed on updates.
			got.UpdateTime = updated.UpdateTime
			if diff := cmp.Diff(updated, got, cmp.Comparer(proto.Equal)); diff != "" {
				t.Errorf("client.UpdateBlurb() got=-, want=+:%s", diff)
			}

			list := &pb.ListBlurbsRequest{
				Parent:    parent,
				PageSize:  1,
				PageToken: "",
			}
			it := client.ListBlurbs(context.Background(), list)
			if mSize := it.PageInfo().MaxSize; mSize != int(list.PageSize) {
				t.Errorf("client.ListBlurbs(): it.PageInfo().MaxSize = %d, want %d", mSize, list.PageSize)
			}

			// got first one
			listed, err := it.Next()
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(listed, first, cmp.Comparer(proto.Equal)); diff != "" {
				t.Errorf("client.ListBlurbs() got=-, want=+:%s", diff)
			}
		})
	}
}

func Test_CreateBlurb_invalid(t *testing.T) {
	usr, errU := newUser("Test_CreateBlurb_invalid")
	if errU != nil {
		t.Fatalf("identityGRPC.CreateUser() failed: %v", errU)
	}
	parent := usr.GetName() + "/profile"

	for typ, client := range map[string]*showcase.MessagingClient{
		"grpc": messagingGRPC,
		// "rest": messagingREST,
	} {
		t.Run(typ, func(t *testing.T) {
			_, err := client.CreateBlurb(context.Background(),
				&pb.CreateBlurbRequest{
					Parent: parent,
					Blurb:  &pb.Blurb{},
				})

			st, _ := status.FromError(err)
			want := "The field `user` is required."
			if st.Message() != want {
				t.Errorf("client.CreateBlurb() Want Message %s got %s",
					want, st.Message())
			}
		})
	}
}
