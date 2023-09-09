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
	"github.com/qclaogui/golang-api-server/pkg/protocol/grpc/interceptors"
	"github.com/qclaogui/golang-api-server/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunGRPCServer runs gRPC service to publish service
func RunGRPCServer(ctx context.Context, cfg server.Config, backends ...service.Backend) error {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPCListenAddress, cfg.GRPCListenPort))
	if err != nil {
		return err
	}

	// gRPC server startup options
	grpcServer := grpc.NewServer(interceptors.RegisterGRPCServerOption()...)
	//	register backend service
	for _, backend := range backends {
		backend.RegisterGRPC(grpcServer)
	}

	// Register reflection service on gRPC server.
	// Enable reflection to allow clients to query the server's services
	reflection.Register(grpcServer)

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		for range quit {
			// sig is a ^C, handle it
			slog.Warn("Graceful shutting down gRPC server...")

			grpcServer.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gRPC server
	slog.Warn("starting gRPC server...", "grpc_port", cfg.GRPCListenPort)
	return grpcServer.Serve(listen)
}
