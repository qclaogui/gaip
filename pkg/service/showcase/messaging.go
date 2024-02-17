// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package showcase

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
)

func (srv *Server) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.Room, error) {
	return srv.repoMessaging.CreateRoom(ctx, req)
}
func (srv *Server) GetRoom(ctx context.Context, req *pb.GetRoomRequest) (*pb.Room, error) {
	return srv.repoMessaging.GetRoom(ctx, req)
}

func (srv *Server) UpdateRoom(ctx context.Context, req *pb.UpdateRoomRequest) (*pb.Room, error) {
	return srv.repoMessaging.UpdateRoom(ctx, req)
}

func (srv *Server) DeleteRoom(ctx context.Context, req *pb.DeleteRoomRequest) (*emptypb.Empty, error) {
	return srv.repoMessaging.DeleteRoom(ctx, req)
}

func (srv *Server) ListRooms(ctx context.Context, req *pb.ListRoomsRequest) (*pb.ListRoomsResponse, error) {
	return srv.repoMessaging.ListRooms(ctx, req)
}
