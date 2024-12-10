// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package showcase

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	pb "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (srv *Server) CreateRoom(ctx context.Context, in *pb.CreateRoomRequest) (*pb.Room, error) {
	return srv.repoMessaging.CreateRoom(ctx, in)
}

func (srv *Server) GetRoom(ctx context.Context, in *pb.GetRoomRequest) (*pb.Room, error) {
	return srv.repoMessaging.GetRoom(ctx, in)
}

func (srv *Server) UpdateRoom(ctx context.Context, in *pb.UpdateRoomRequest) (*pb.Room, error) {
	return srv.repoMessaging.UpdateRoom(ctx, in)
}

func (srv *Server) DeleteRoom(ctx context.Context, in *pb.DeleteRoomRequest) (*emptypb.Empty, error) {
	return srv.repoMessaging.DeleteRoom(ctx, in)
}

func (srv *Server) ListRooms(ctx context.Context, in *pb.ListRoomsRequest) (*pb.ListRoomsResponse, error) {
	return srv.repoMessaging.ListRooms(ctx, in)
}

func (srv *Server) CreateBlurb(ctx context.Context, in *pb.CreateBlurbRequest) (*pb.Blurb, error) {
	parent := in.GetParent()
	if err := srv.validateParent(ctx, parent); err != nil {
		return nil, err
	}

	b := in.GetBlurb()
	if err := validateBlurb(b); err != nil {
		return nil, err
	}

	return srv.repoMessaging.CreateBlurb(ctx, in)
}

func (srv *Server) GetBlurb(ctx context.Context, in *pb.GetBlurbRequest) (*pb.Blurb, error) {
	return srv.repoMessaging.GetBlurb(ctx, in)
}

func (srv *Server) UpdateBlurb(ctx context.Context, in *pb.UpdateBlurbRequest) (*pb.Blurb, error) {
	return srv.repoMessaging.UpdateBlurb(ctx, in)
}

func (srv *Server) DeleteBlurb(ctx context.Context, in *pb.DeleteBlurbRequest) (*emptypb.Empty, error) {
	return srv.repoMessaging.DeleteBlurb(ctx, in)
}

func (srv *Server) ListBlurbs(ctx context.Context, in *pb.ListBlurbsRequest) (*pb.ListBlurbsResponse, error) {
	if err := srv.validateParent(ctx, in.GetParent()); err != nil {
		return nil, err
	}
	return srv.repoMessaging.ListBlurbs(ctx, in)
}

// SearchBlurbs This method searches through all blurbs across all rooms and profiles
// for blurbs containing to words found in the query. Only posts that
// contain an exact match of a queried word will be returned.
func (srv *Server) SearchBlurbs(ctx context.Context, in *pb.SearchBlurbsRequest) (*longrunningpb.Operation, error) {
	if err := srv.validateParent(ctx, in.GetParent()); err != nil {
		return nil, err
	}

	reqBytes, _ := proto.Marshal(in)
	name := fmt.Sprintf(
		"operations/qclaogui.showcase.v1beta1.MessagingService/SearchBlurbs/%s",
		base64.StdEncoding.EncodeToString(reqBytes))

	meta, _ := anypb.New(&pb.SearchBlurbsMetadata{
		RetryInfo: &errdetails.RetryInfo{
			RetryDelay: durationpb.New(1 * time.Second),
		},
	})

	return &longrunningpb.Operation{Name: name, Done: false, Metadata: meta}, nil
}

// StreamBlurbs This returns a stream that emits the blurbs that are created for a
// particular chat room or user profile.
func (srv *Server) StreamBlurbs(in *pb.StreamBlurbsRequest, stream pb.MessagingService_StreamBlurbsServer) error {
	parent := in.GetName()
	if err := srv.validateParent(stream.Context(), parent); err != nil {
		return err
	}

	expireTime := in.GetExpireTime()
	if err := expireTime.CheckValid(); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	for {
		if srv.nowF().After(expireTime.AsTime()) {
			break
		}
		if err := srv.validateParent(stream.Context(), parent); err != nil {
			return err
		}
	}

	return nil
}

// SendBlurbs This is a stream to create multiple blurbs. If an invalid blurb is
// requested to be created, the stream will close with an error.
func (srv *Server) SendBlurbs(stream pb.MessagingService_SendBlurbsServer) error {
	var names []string

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return stream.SendAndClose(&pb.SendBlurbsResponse{Names: names})
		}

		if err != nil {
			return withCreatedNames(err, names)
		}

		ctx := stream.Context()
		parent := req.GetParent()
		if err = srv.validateParent(ctx, parent); err != nil {
			return withCreatedNames(err, names)
		}

		blurb, err := srv.CreateBlurb(
			ctx,
			&pb.CreateBlurbRequest{Parent: parent, Blurb: req.GetBlurb()},
		)
		if err != nil {
			return withCreatedNames(err, names)
		}

		names = append(names, blurb.GetName())
	}
}

// Connect This method starts a bidirectional stream that receives all blurbs that
// are being created after the stream has started and sends requests to create
// blurbs. If an invalid blurb is requested to be created, the stream will
// close with an error.
func (srv *Server) Connect(stream pb.MessagingService_ConnectServer) error {
	configured := false
	parent := ""

	for {
		req, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}

		// Setup Configuration
		if !configured && req != nil {
			if req.GetConfig() == nil {
				return status.Error(codes.InvalidArgument, "The first request to Connect, must contain a config field")
			}

			parent = req.GetConfig().GetParent()
			if err = srv.validateParent(stream.Context(), parent); err != nil {
				return err
			}

			configured = true
			continue
		}

		// Check if the parent still exists.
		if err = srv.validateParent(stream.Context(), parent); err != nil {
			return err
		}

		// Create the blurb
		if req == nil || req.GetBlurb() == nil {
			continue
		}

		_, err = srv.CreateBlurb(
			stream.Context(),
			&pb.CreateBlurbRequest{Parent: parent, Blurb: req.GetBlurb()},
		)
		if err != nil {
			return err
		}
	}
}

func (srv *Server) validateParent(ctx context.Context, p string) error {
	_, uErr := srv.GetUser(ctx, &pb.GetUserRequest{Name: strings.TrimSuffix(p, "/profile")})
	_, rErr := srv.GetRoom(ctx, &pb.GetRoomRequest{Name: p})
	if uErr != nil && rErr != nil {
		return status.Errorf(codes.NotFound, "Parent %s not found.", p)
	}
	return nil
}

func withCreatedNames(err error, names []string) error {
	s, _ := status.FromError(err)
	spb := s.Proto()

	details, err := anypb.New(&pb.SendBlurbsResponse{Names: names})
	if err == nil {
		spb.Details = append(spb.Details, details)
	}

	return status.ErrorProto(spb)
}

func validateBlurb(b *pb.Blurb) error {
	// Validate Required Fields.
	if b.GetUser() == "" {
		return status.Errorf(codes.InvalidArgument, "The field `user` is required.")
	}
	return nil
}
