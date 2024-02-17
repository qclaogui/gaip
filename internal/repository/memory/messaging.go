// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"github.com/qclaogui/gaip/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type roomEntry struct {
	room    *showcasepb.Room
	deleted bool
}

func NewMessaging() (showcasepb.MessagingServiceServer, error) {
	s := &messagingImpl{
		nowF:     time.Now,
		token:    service.NewTokenGenerator(),
		roomKeys: map[string]int{},
	}
	return s, nil
}

// The messagingImpl type implements a projectpb.MessagingServiceServer.
type messagingImpl struct {
	showcasepb.UnimplementedMessagingServiceServer

	nowF  func() time.Time
	token service.TokenGenerator

	roomUID  service.UniqID
	roomMu   sync.Mutex
	roomKeys map[string]int
	rooms    []roomEntry
}

// CreateRoom Creates a room.
func (m *messagingImpl) CreateRoom(_ context.Context, req *showcasepb.CreateRoomRequest) (*showcasepb.Room, error) {
	m.roomMu.Lock()
	defer m.roomMu.Unlock()

	r := req.GetRoom()
	err := validateRoom(r)
	if err != nil {
		return nil, err
	}

	// Validate Unique Fields.
	uniqName := func(x *showcasepb.Room) bool {
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
func (m *messagingImpl) GetRoom(context.Context, *showcasepb.GetRoomRequest) (*showcasepb.Room, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}

func (m *messagingImpl) anyRoom(f func(*showcasepb.Room) bool) bool {
	for _, entry := range m.rooms {
		if !entry.deleted && f(entry.room) {
			return true
		}
	}
	return false
}

func validateRoom(r *showcasepb.Room) error {
	// Validate Required Fields.
	if r.GetDisplayName() == "" {
		return status.Errorf(
			codes.InvalidArgument,
			"The field `display_name` is required.")
	}
	return nil
}
