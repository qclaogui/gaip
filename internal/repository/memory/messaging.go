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
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
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
func (m *messagingImpl) CreateRoom(_ context.Context, in *showcasepb.CreateRoomRequest) (*showcasepb.Room, error) {
	m.roomMu.Lock()
	defer m.roomMu.Unlock()

	r := in.GetRoom()
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
func (m *messagingImpl) GetRoom(_ context.Context, in *showcasepb.GetRoomRequest) (*showcasepb.Room, error) {
	m.roomMu.Lock()
	defer m.roomMu.Unlock()

	if i, ok := m.roomKeys[in.GetName()]; ok {
		entry := m.rooms[i]
		if !entry.deleted {
			return entry.room, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "A room with name %s not found.", in.GetName())
}

func (m *messagingImpl) UpdateRoom(_ context.Context, in *showcasepb.UpdateRoomRequest) (*showcasepb.Room, error) {
	r := in.GetRoom()
	mask := in.GetUpdateMask()

	m.roomMu.Lock()
	defer m.roomMu.Unlock()

	err := validateRoom(r)
	if err != nil {
		return nil, err
	}

	i, ok := m.roomKeys[r.GetName()]
	if !ok || m.rooms[i].deleted {
		return nil, status.Errorf(codes.NotFound, "A room with name %s not found.", r.GetName())
	}

	existing := m.rooms[i].room
	// Validate Unique Fields.
	uniqName := func(x *showcasepb.Room) bool {
		return x != existing && (r.GetDisplayName() == x.GetDisplayName())
	}

	if m.anyRoom(uniqName) {
		return nil, status.Errorf(codes.AlreadyExists, "A room with either display_name, %s already exists.", r.GetDisplayName())
	}

	// Update store.
	updated := proto.Clone(existing).(*showcasepb.Room)
	applyFieldMask(r.ProtoReflect(), updated.ProtoReflect(), mask.GetPaths())
	updated.CreateTime = existing.GetCreateTime()
	updated.UpdateTime = timestamppb.Now()

	m.rooms[i] = roomEntry{room: updated}
	return updated, nil
}

func (m *messagingImpl) DeleteRoom(_ context.Context, in *showcasepb.DeleteRoomRequest) (*emptypb.Empty, error) {
	m.roomMu.Lock()
	defer m.roomMu.Unlock()

	i, ok := m.roomKeys[in.GetName()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "A room with name %s not found.", in.GetName())
	}

	entry := m.rooms[i]
	m.rooms[i] = roomEntry{room: entry.room, deleted: true}

	return &emptypb.Empty{}, nil
}

// ListRooms Lists all chat rooms.
func (m *messagingImpl) ListRooms(_ context.Context, in *showcasepb.ListRoomsRequest) (*showcasepb.ListRoomsResponse, error) {
	start, err := m.token.GetIndex(in.GetPageToken())
	if err != nil {
		return nil, err
	}

	offset := 0
	var rooms []*showcasepb.Room
	for _, entry := range m.rooms[start:] {
		offset++
		if entry.deleted {
			continue
		}
		rooms = append(rooms, entry.room)
		if len(rooms) >= int(in.GetPageSize()) {
			break
		}
	}

	var nextToken string
	if start+offset < len(m.rooms) {
		nextToken = m.token.ForIndex(start + offset)
	}

	return &showcasepb.ListRoomsResponse{Rooms: rooms, NextPageToken: nextToken}, nil
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
		return status.Errorf(codes.InvalidArgument, "The field `display_name` is required.")
	}
	return nil
}
