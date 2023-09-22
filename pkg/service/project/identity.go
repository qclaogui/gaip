// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"
	"fmt"
	"sync"

	"github.com/qclaogui/golang-api-server/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/golang-api-server/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// IdentityServer Identity service server
type IdentityServer interface {
	service.Backend

	projectpb.IdentityServiceServer
}

type userEntry struct {
	user    *projectpb.User
	deleted bool
}

func NewIdentityServer() (IdentityServer, error) {
	s := &identityServerImpl{
		token: service.NewTokenGenerator(),
		keys:  map[string]int{},
	}
	return s, nil
}

// The identityServerImpl type implements a project server.
type identityServerImpl struct {
	uid   service.UniqID
	token service.TokenGenerator

	mu    sync.Mutex
	keys  map[string]int
	users []userEntry
}

func (s *identityServerImpl) RegisterGRPC(grpcServer *grpc.Server) {
	grpcServer.RegisterService(&projectpb.IdentityService_ServiceDesc, s)
}

// validate validate
func (s *identityServerImpl) validate(u *projectpb.User) error {
	// Validate Required Fields.
	if u.GetDisplayName() == "" {
		return status.Errorf(codes.InvalidArgument, "The field `display_name` is required.")
	}
	if u.GetEmail() == "" {
		return status.Errorf(codes.InvalidArgument, "The field `email` is required.")
	}

	// Validate Unique Fields.
	for _, x := range s.users {
		if x.deleted {
			continue
		}

		if (u.GetEmail() == x.user.GetEmail()) && (u.GetName() != x.user.GetName()) {
			return status.Errorf(codes.AlreadyExists, "A user with email %s already exists.", u.GetEmail())
		}
	}

	return nil
}

// CreateUser Create User
func (s *identityServerImpl) CreateUser(_ context.Context, req *projectpb.CreateUserRequest) (*projectpb.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	u := req.GetUser()

	// Ignore passed in name.
	u.Name = ""

	err := s.validate(u)
	if err != nil {
		return nil, err
	}

	// Assign info.
	id := s.uid.Next()
	name := fmt.Sprintf("users/%d", id)

	now := timestamppb.Now()
	u.Name = name
	u.CreateTime = now
	u.UpdateTime = now

	// Insert.
	index := len(s.users)
	s.users = append(s.users, userEntry{user: u})
	s.keys[name] = index

	return u, nil
}

// GetUser Get User
func (s *identityServerImpl) GetUser(_ context.Context, req *projectpb.GetUserRequest) (*projectpb.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	name := req.GetName()
	if i, ok := s.keys[name]; ok {
		entry := s.users[i]
		if !entry.deleted {
			return entry.user, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "A user with name %s not found.", name)
}

// ListUsers List Users
func (s *identityServerImpl) ListUsers(_ context.Context, req *projectpb.ListUsersRequest) (*projectpb.ListUsersResponse, error) {
	start, err := s.token.GetIndex(req.GetPageToken())
	if err != nil {
		return nil, err
	}
	var users []*projectpb.User
	offset := 0
	for _, entry := range s.users[start:] {
		offset++
		if entry.deleted {
			continue
		}

		users = append(users, entry.user)
		if len(users) >= int(req.GetPageSize()) {
			break
		}
	}

	nextToken := ""
	if start+offset < len(users) {
		nextToken = s.token.ForIndex(start + offset)
	}

	return &projectpb.ListUsersResponse{Users: users, NextPageToken: nextToken}, nil
}

// UpdateUser Update User
func (s *identityServerImpl) UpdateUser(_ context.Context, req *projectpb.UpdateUserRequest) (*projectpb.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	mask := req.GetUpdateMask()
	u := req.GetUser()

	i, ok := s.keys[u.GetName()]
	if !ok || s.users[i].deleted {
		return nil, status.Errorf(codes.NotFound, "A user with name %s not found.", u.GetName())
	}

	err := s.validate(u)
	if err != nil {
		return nil, err
	}

	// Update store.
	existing := s.users[i].user
	updated := proto.Clone(existing).(*projectpb.User)
	applyFieldMask(u.ProtoReflect(), updated.ProtoReflect(), mask.GetPaths())
	updated.CreateTime = existing.GetCreateTime()
	updated.UpdateTime = timestamppb.Now()

	s.users[i] = userEntry{user: updated}
	return updated, nil
}

// DeleteUser Delete User
func (s *identityServerImpl) DeleteUser(_ context.Context, req *projectpb.DeleteUserRequest) (*emptypb.Empty, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	i, ok := s.keys[req.GetName()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "A user with name %s not found.", req.GetName())
	}

	entry := s.users[i]
	s.users[i] = userEntry{user: entry.user, deleted: true}

	return &emptypb.Empty{}, nil
}

func strContains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

// applyFieldMask applies the values from the src message to the values of the
// dst message according to the contents of the given field mask paths.
// If paths is empty/nil, or contains *, it is considered a full update.
//
// TODO: Does not support nested message paths. Currently only used with flat
// resource messages.
func applyFieldMask(src, dst protoreflect.Message, paths []string) {
	fullUpdate := len(paths) == 0 || strContains(paths, "*")

	fields := dst.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		isOneof := field.ContainingOneof() != nil && !field.ContainingOneof().IsSynthetic()

		// Set field in dst with value from src, skipping true oneofs, while
		// handling proto3_optional fields represented as synthetic oneofs.
		if (fullUpdate || strContains(paths, string(field.Name()))) && !isOneof {
			dst.Set(field, src.Get(field))
		}
	}

	oneofs := dst.Descriptor().Oneofs()
	for i := 0; i < oneofs.Len(); i++ {
		oneof := oneofs.Get(i)
		// Skip proto3_optional synthetic oneofs.
		if oneof.IsSynthetic() {
			continue
		}

		setOneof := src.WhichOneof(oneof)
		if setOneof == nil && fullUpdate {
			// Full update with no field set in this oneof of
			// src means clear all fields for this oneof in dst.
			fields := oneof.Fields()
			for j := 0; j < fields.Len(); j++ {
				dst.Clear(fields.Get(j))
			}
		} else if setOneof != nil && (fullUpdate || strContains(paths, string(setOneof.Name()))) {
			// Full update or targeted updated with a field set in this oneof of
			// src means set that field for the same oneof in dst, which implicitly
			// clears any previously set field for this oneof.
			dst.Set(setOneof, src.Get(setOneof))
		}
	}

}
