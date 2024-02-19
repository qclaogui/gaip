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

type blurbEntry struct {
	blurb   *showcasepb.Blurb
	deleted bool
}

type blurbIndex struct {
	// The parent of the blurb.
	row string
	// The index within the list of blurbs of a parent.
	col int
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

	blurbMu    sync.Mutex
	blurbKeys  map[string]blurbIndex
	blurbs     map[string][]blurbEntry
	parentUIDs map[string]*service.UniqID
}

// CreateRoom Creates a room.
func (s *messagingImpl) CreateRoom(_ context.Context, in *showcasepb.CreateRoomRequest) (*showcasepb.Room, error) {
	s.roomMu.Lock()
	defer s.roomMu.Unlock()

	r := in.GetRoom()
	err := validateRoom(r)
	if err != nil {
		return nil, err
	}

	// Validate Unique Fields.
	uniqName := func(x *showcasepb.Room) bool {
		return r.GetDisplayName() == x.GetDisplayName()
	}
	if s.anyRoom(uniqName) {
		return nil, status.Errorf(codes.AlreadyExists, "A room with display_name=%s already exists.", r.GetDisplayName())
	}

	// Assign info.
	id := s.roomUID.Next()
	name := fmt.Sprintf("rooms/%d", id)
	now := timestamppb.Now()

	r.Name = name
	r.CreateTime = now
	r.UpdateTime = now

	// Insert.
	index := len(s.rooms)
	s.rooms = append(s.rooms, roomEntry{room: r})
	s.roomKeys[name] = index

	return r, nil
}

// GetRoom Retrieves the Room with the given resource name.
func (s *messagingImpl) GetRoom(_ context.Context, in *showcasepb.GetRoomRequest) (*showcasepb.Room, error) {
	s.roomMu.Lock()
	defer s.roomMu.Unlock()

	if i, ok := s.roomKeys[in.GetName()]; ok {
		entry := s.rooms[i]
		if !entry.deleted {
			return entry.room, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "A room with name %s not found.", in.GetName())
}

func (s *messagingImpl) UpdateRoom(_ context.Context, in *showcasepb.UpdateRoomRequest) (*showcasepb.Room, error) {
	r := in.GetRoom()
	mask := in.GetUpdateMask()

	s.roomMu.Lock()
	defer s.roomMu.Unlock()

	err := validateRoom(r)
	if err != nil {
		return nil, err
	}

	i, ok := s.roomKeys[r.GetName()]
	if !ok || s.rooms[i].deleted {
		return nil, status.Errorf(codes.NotFound, "A room with name %s not found.", r.GetName())
	}

	existing := s.rooms[i].room
	// Validate Unique Fields.
	uniqName := func(x *showcasepb.Room) bool {
		return x != existing && (r.GetDisplayName() == x.GetDisplayName())
	}

	if s.anyRoom(uniqName) {
		return nil, status.Errorf(codes.AlreadyExists, "A room with either display_name, %s already exists.", r.GetDisplayName())
	}

	// Update store.
	updated := proto.Clone(existing).(*showcasepb.Room)
	applyFieldMask(r.ProtoReflect(), updated.ProtoReflect(), mask.GetPaths())
	updated.CreateTime = existing.GetCreateTime()
	updated.UpdateTime = timestamppb.Now()

	s.rooms[i] = roomEntry{room: updated}
	return updated, nil
}

func (s *messagingImpl) DeleteRoom(_ context.Context, in *showcasepb.DeleteRoomRequest) (*emptypb.Empty, error) {
	s.roomMu.Lock()
	defer s.roomMu.Unlock()

	i, ok := s.roomKeys[in.GetName()]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "A room with name %s not found.", in.GetName())
	}

	entry := s.rooms[i]
	s.rooms[i] = roomEntry{room: entry.room, deleted: true}

	return &emptypb.Empty{}, nil
}

// ListRooms Lists all chat rooms.
func (s *messagingImpl) ListRooms(_ context.Context, in *showcasepb.ListRoomsRequest) (*showcasepb.ListRoomsResponse, error) {
	start, err := s.token.GetIndex(in.GetPageToken())
	if err != nil {
		return nil, err
	}

	offset := 0
	var rooms []*showcasepb.Room
	for _, entry := range s.rooms[start:] {
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
	if start+offset < len(s.rooms) {
		nextToken = s.token.ForIndex(start + offset)
	}

	return &showcasepb.ListRoomsResponse{Rooms: rooms, NextPageToken: nextToken}, nil
}

func (s *messagingImpl) anyRoom(f func(*showcasepb.Room) bool) bool {
	for _, entry := range s.rooms {
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

// CreateBlurb Creates a blurb. If the parent is a room, the blurb is understood to be a
// message in that room. If the parent is a profile, the blurb is understood
// to be a post on the profile.
func (s *messagingImpl) CreateBlurb(_ context.Context, in *showcasepb.CreateBlurbRequest) (*showcasepb.Blurb, error) {
	parent := in.GetParent()

	s.blurbMu.Lock()
	defer s.blurbMu.Unlock()

	b := in.GetBlurb()

	// Assign info.
	parentBs, ok := s.blurbs[parent]
	if !ok {
		parentBs = []blurbEntry{}
	}

	pUID, ok := s.parentUIDs[parent]
	if !ok {
		pUID = &service.UniqID{}
		s.parentUIDs[parent] = pUID
	}

	id := pUID.Next()
	name := fmt.Sprintf("%s/blurbs/%d", parent, id)
	now := timestamppb.Now()

	b.Name = name
	b.CreateTime = now
	b.UpdateTime = now

	// Insert.
	index := len(parentBs)
	s.blurbs[parent] = append(parentBs, blurbEntry{blurb: b})
	s.blurbKeys[name] = blurbIndex{row: parent, col: index}

	return b, nil
}

// GetBlurb Retrieves the Blurb with the given resource name.
func (s *messagingImpl) GetBlurb(_ context.Context, in *showcasepb.GetBlurbRequest) (*showcasepb.Blurb, error) {
	s.blurbMu.Lock()
	defer s.blurbMu.Unlock()

	if i, ok := s.blurbKeys[in.GetName()]; ok {
		entry := s.blurbs[i.row][i.col]
		if !entry.deleted {
			return entry.blurb, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, "A blurb with name %s not found.", in.GetName())
}

// UpdateBlurb Updates a blurb.
func (s *messagingImpl) UpdateBlurb(_ context.Context, in *showcasepb.UpdateBlurbRequest) (*showcasepb.Blurb, error) {
	s.blurbMu.Lock()
	defer s.blurbMu.Unlock()

	mask := in.GetUpdateMask()
	b := in.GetBlurb()
	i, ok := s.blurbKeys[b.GetName()]
	if !ok || s.blurbs[i.row][i.col].deleted {
		return nil, status.Errorf(codes.NotFound, "A blurb with name %s not found.", b.GetName())
	}

	if b.GetUser() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "The field `user` is required.")
	}

	// Update store.
	existing := s.blurbs[i.row][i.col].blurb
	updated := proto.Clone(existing).(*showcasepb.Blurb)
	applyFieldMask(b.ProtoReflect(), updated.ProtoReflect(), mask.GetPaths())
	updated.CreateTime = existing.GetCreateTime()
	updated.UpdateTime = timestamppb.Now()
	s.blurbs[i.row][i.col] = blurbEntry{blurb: updated}

	return updated, nil
}

// DeleteBlurb Deletes a blurb.
func (s *messagingImpl) DeleteBlurb(_ context.Context, in *showcasepb.DeleteBlurbRequest) (*emptypb.Empty, error) {
	s.blurbMu.Lock()
	defer s.blurbMu.Unlock()

	i, ok := s.blurbKeys[in.GetName()]
	if ok {
		return nil, status.Errorf(codes.NotFound, "A blurb with name %s not found.", in.GetName())
	}

	entry := s.blurbs[i.row][i.col]
	s.blurbs[i.row][i.col] = blurbEntry{blurb: entry.blurb, deleted: true}
	return &emptypb.Empty{}, nil
}

// ListBlurbs Lists blurbs for a specific chat room or user profile depending on the
// parent resource name.
func (s *messagingImpl) ListBlurbs(ctx context.Context, in *showcasepb.ListBlurbsRequest) (*showcasepb.ListBlurbsResponse, error) {
	passFilter := func(_ *showcasepb.Blurb) bool { return true }
	return s.FilteredListBlurbs(ctx, in, passFilter)
}

func (s *messagingImpl) FilteredListBlurbs(_ context.Context, in *showcasepb.ListBlurbsRequest, f func(*showcasepb.Blurb) bool) (*showcasepb.ListBlurbsResponse, error) {
	bs, ok := s.blurbs[in.GetParent()]
	if !ok {
		return &showcasepb.ListBlurbsResponse{}, nil
	}

	start, err := s.token.GetIndex(in.GetPageToken())
	if err != nil {
		return nil, err
	}

	offset := 0
	var blurbs []*showcasepb.Blurb
	for _, entry := range bs[start:] {
		offset++
		if entry.deleted {
			continue
		}

		// run filter
		if f(entry.blurb) {
			blurbs = append(blurbs, entry.blurb)
		}

		if len(blurbs) >= int(in.GetPageSize()) {
			break
		}
	}

	nextToken := ""
	if start+offset < len(bs) {
		nextToken = s.token.ForIndex(start + offset)
	}

	return &showcasepb.ListBlurbsResponse{Blurbs: blurbs, NextPageToken: nextToken}, nil
}
