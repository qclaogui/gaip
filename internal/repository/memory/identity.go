// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"context"
	"fmt"
	"sync"

	"github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"github.com/qclaogui/gaip/internal/pagination"
	"github.com/qclaogui/gaip/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type userEntry struct {
	user    *showcasepb.User
	deleted bool
}

func NewIdentity() (showcasepb.IdentityServiceServer, error) {
	s := &identityImpl{
		token: service.NewTokenGenerator(),
		keys:  map[string]int{},
	}
	return s, nil
}

// The identityImpl type implements a projectpb.IdentityServiceServer.
type identityImpl struct {
	showcasepb.UnimplementedIdentityServiceServer

	uid   service.UniqID
	token service.TokenGenerator

	mu    sync.Mutex
	keys  map[string]int
	users []userEntry
}

// validate validate
func (s *identityImpl) validate(u *showcasepb.User) error {
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
func (s *identityImpl) CreateUser(_ context.Context, req *showcasepb.CreateUserRequest) (*showcasepb.User, error) {
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
func (s *identityImpl) GetUser(_ context.Context, req *showcasepb.GetUserRequest) (*showcasepb.User, error) {
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
func (s *identityImpl) ListUsers(_ context.Context, req *showcasepb.ListUsersRequest) (*showcasepb.ListUsersResponse, error) {
	pageToken, err := pagination.ParsePageToken(req)
	if err != nil {
		return nil, err
	}

	startPos := pageToken.Offset
	// endPos := startPos + int64(req.GetPageSize())

	start, err := s.token.GetIndex(req.GetPageToken())
	if err != nil {
		return nil, err
	}
	var users []*showcasepb.User
	offset := 0
	for _, entry := range s.users[startPos:] {
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

	return &showcasepb.ListUsersResponse{Users: users, NextPageToken: nextToken}, nil
}

// UpdateUser Update User
func (s *identityImpl) UpdateUser(_ context.Context, req *showcasepb.UpdateUserRequest) (*showcasepb.User, error) {
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
	updated := proto.Clone(existing).(*showcasepb.User)
	applyFieldMask(u.ProtoReflect(), updated.ProtoReflect(), mask.GetPaths())
	updated.CreateTime = existing.GetCreateTime()
	updated.UpdateTime = timestamppb.Now()

	s.users[i] = userEntry{user: updated}
	return updated, nil
}

// DeleteUser Delete User
func (s *identityImpl) DeleteUser(_ context.Context, req *showcasepb.DeleteUserRequest) (*emptypb.Empty, error) {
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
