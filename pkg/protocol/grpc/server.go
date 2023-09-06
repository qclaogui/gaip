// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/grafana/dskit/server"
	pbrouteguidev1 "github.com/qclaogui/golang-api-server/api/gen/proto/routeguide/v1"
	pbtodov1 "github.com/qclaogui/golang-api-server/api/gen/proto/todo/v1"
	"github.com/qclaogui/golang-api-server/pkg/protocol/grpc/interceptors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunGRPCServer runs gRPC service to publish service
func RunGRPCServer(
	ctx context.Context,
	toDoSrv pbtodov1.ToDoServiceServer,
	routeGuideSrv pbrouteguidev1.RouteGuideServiceServer,
	cfg server.Config,
) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPCListenAddress, cfg.GRPCListenPort))
	if err != nil {
		return err
	}

	// gRPC server startup options
	srv := grpc.NewServer(interceptors.RegisterGRPCServerOption()...)

	//	register service
	pbtodov1.RegisterToDoServiceServer(srv, toDoSrv)
	pbrouteguidev1.RegisterRouteGuideServiceServer(srv, routeGuideSrv)

	// Register reflection service on gRPC server.
	// Enable reflection to allow clients to query the server's services
	reflection.Register(srv)

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		for range quit {
			// sig is a ^C, handle it
			slog.Warn("Graceful shutting down gRPC server...")

			srv.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gRPC server
	slog.Warn("starting gRPC server...", "grpc_port", cfg.GRPCListenPort)
	return srv.Serve(listen)
}
