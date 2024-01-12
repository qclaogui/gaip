// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type roomEntry struct {
	room    *projectpb.Room
	deleted bool
}

func NewMessaging() (projectpb.MessagingServiceServer, error) {
	s := &messagingMemImpl{
		nowF:     time.Now,
		token:    service.NewTokenGenerator(),
		roomKeys: map[string]int{},
	}
	return s, nil
}

// The messagingMemImpl type implements a projectpb.MessagingServiceServer.
type messagingMemImpl struct {
	nowF  func() time.Time
	token service.TokenGenerator

	roomUID  service.UniqID
	roomMu   sync.Mutex
	roomKeys map[string]int
	rooms    []roomEntry
}

// CreateRoom Creates a room.
func (m *messagingMemImpl) CreateRoom(_ context.Context, req *projectpb.CreateRoomRequest) (*projectpb.Room, error) {
	m.roomMu.Lock()
	defer m.roomMu.Unlock()

	r := req.GetRoom()
	err := validateRoom(r)
	if err != nil {
		return nil, err
	}

	// Validate Unique Fields.
	uniqName := func(x *projectpb.Room) bool {
		return r.GetDisplayName() == x.GetDisplayName()
	}
	if m.anyRoom(uniqName) {
		return nil, status.Errorf(codes.AlreadyExists, "A room with display_name=%s already exists.", r.GetDisplayName())
	}

	// Assign info.
	id := m.roomUID.Next()
	name := fmt.Sprintf("rooms/%d", id)
	now := timestamppb.Now()

	r.Name = name
	r.CreateTime = now
	r.UpdateTime = now

	// Insert.
	index := len(m.rooms)
	m.rooms = append(m.rooms, roomEntry{room: r})
	m.roomKeys[name] = index

	return r, nil
}

// GetRoom Retrieves the Room with the given resource name.
func (m *messagingMemImpl) GetRoom(context.Context, *projectpb.GetRoomRequest) (*projectpb.Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}

func (m *messagingMemImpl) anyRoom(f func(*projectpb.Room) bool) bool {
	for _, entry := range m.rooms {
		if !entry.deleted && f(entry.room) {
			return true
		}
	}
	return false
}

func validateRoom(r *projectpb.Room) error {
	// Validate Required Fields.
	if r.GetDisplayName() == "" {
		return status.Errorf(
			codes.InvalidArgument,
			"The field `display_name` is required.")
	}
	return nil
}
