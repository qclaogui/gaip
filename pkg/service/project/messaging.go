// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package project

import (
	"context"

	pb "github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
)

func (srv *Server) CreateRoom(ctx context.Context, req *pb.CreateRoomRequest) (*pb.Room, error) {
	return srv.repoMessaging.CreateRoom(ctx, req)
}
func (srv *Server) GetRoom(ctx context.Context, req *pb.GetRoomRequest) (*pb.Room, error) {
	return srv.repoMessaging.GetRoom(ctx, req)
}
