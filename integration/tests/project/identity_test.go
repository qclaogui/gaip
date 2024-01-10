// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.
//go:build requires_docker

package project_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	project "github.com/qclaogui/gaip/genproto/project/apiv1"
	pb "github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"google.golang.org/api/iterator"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Clients are initialized in main_test.go.
var (
	identityGRPC *project.IdentityClient
	identityREST *project.IdentityClient
)

func TestUserCRUD(t *testing.T) {
	ctx := context.Background()

	for _, client := range map[string]*project.IdentityClient{
		"grpc": identityGRPC,
		"rest": identityREST,
	} {
		// Create UserRequest
		create := &pb.CreateUserRequest{
			User: &pb.User{
				DisplayName: "Jane Doe",
				Email:       "janedoe@example.com",
				Nickname:    proto.String("Doe"),
				HeightFeet:  proto.Float64(6.2),
			},
		}
		usr, err := client.CreateUser(ctx, create)
		if err != nil {
			t.Fatal(err)
		}

		want := create.GetUser()
		if usr.GetName() == "" {
			t.Errorf("CreateUser().Name was unexpectedly empty")
		}

		if usr.GetDisplayName() != want.GetDisplayName() {
			t.Errorf("CreateUser().DisplayName = %q, want = %q", usr.GetDisplayName(), want.GetDisplayName())
		}
		if usr.GetEmail() != want.GetEmail() {
			t.Errorf("CreateUser().Email = %q, want = %q", usr.GetEmail(), want.GetEmail())
		}

		if usr.GetCreateTime() == nil {
			t.Errorf("CreateUser().CreateTime was unexpectedly empty")
		}
		if usr.GetUpdateTime() == nil {
			t.Errorf("CreateUser().UpdateTime was unexpectedly empty")
		}
		if usr.GetNickname() != want.GetNickname() {
			t.Errorf("CreateUser().Nickname = %q, want = %q", usr.GetNickname(), want.GetNickname())
		}
		if usr.GetHeightFeet() != want.GetHeightFeet() {
			t.Errorf("CreateUser().HeightFeet = %f, want = %f", usr.GetHeightFeet(), want.GetHeightFeet())
		}

		if usr.Age != nil {
			t.Errorf("CreateUser().Age was unexpectedly set to: %d", usr.GetAge())
		}
		if usr.EnableNotifications != nil {
			t.Errorf("CreateUser().EnableNotifications was unexpectedly set to: %v", usr.GetEnableNotifications())
		}

		// List UsersRequest
		list := &pb.ListUsersRequest{
			PageSize: 5,
		}
		it := client.ListUsers(context.Background(), list)
		if mSize := it.PageInfo().MaxSize; mSize != int(list.PageSize) {
			t.Errorf("PageInfo().MaxSize = %d, want %d", mSize, list.PageSize)
		}

		listed, err := it.Next()
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(listed, usr, cmp.Comparer(proto.Equal)); diff != "" {
			t.Errorf("ListUsers() got=-, want=+:%s", diff)
		}

		//	Get UserRequest
		get := &pb.GetUserRequest{
			Name: usr.GetName(),
		}

		got, err := client.GetUser(ctx, get)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(got, usr, cmp.Comparer(proto.Equal)); diff != "" {
			t.Errorf("GetUser() got=-, want=+:%s", diff)
		}

		//	Update UserRequest
		update := &pb.UpdateUserRequest{
			User: &pb.User{
				Name:                got.GetName(),
				DisplayName:         got.GetDisplayName(),
				Email:               "janedoe@jane.com",
				HeightFeet:          proto.Float64(6.0),
				EnableNotifications: proto.Bool(true),
			},

			UpdateMask: &fieldmaskpb.FieldMask{
				Paths: []string{
					"email",
					"height_feet",
					"enable_notifications",
				},
			},
		}
		updated, err := client.UpdateUser(ctx, update)
		if err != nil {
			t.Fatal(err)
		}

		if diff := cmp.Diff(updated, usr, cmp.Comparer(proto.Equal)); diff == "" {
			t.Errorf("UpdateUser() users were the same, update failed")
		}

		if updated.GetEmail() == usr.GetEmail() {
			t.Errorf("UpdateUser().Email was not updated as expected")
		}
		if updated.GetNickname() != usr.GetNickname() {
			t.Errorf("UpdateUser().Nickname = %q, want = %q", updated.GetNickname(), usr.GetNickname())
		}

		if updated.GetHeightFeet() == usr.GetHeightFeet() {
			t.Errorf("UpdateUser().HeightFeet was not updated as expected")
		}
		if updated.EnableNotifications == nil || !updated.GetEnableNotifications() {
			t.Errorf("UpdateUser().EnableNotifications was not updated as expected")
		}
		if updated.Age != nil {
			t.Errorf("UpdateUser().Age was unexpectedly updated")
		}

		//	Delete UserRequest
		del := &pb.DeleteUserRequest{
			Name: usr.GetName(),
		}
		if err = client.DeleteUser(ctx, del); err != nil {
			t.Fatal(err)
		}

		it = client.ListUsers(context.Background(), &pb.ListUsersRequest{})
		_, err = it.Next()
		if !errors.Is(err, iterator.Done) {
			t.Errorf("ListUsers() = %q, want %q", err, iterator.Done)
		}
	}

}
