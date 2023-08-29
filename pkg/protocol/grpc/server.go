package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"

	"log/slog"

	pbrouteguide "github.com/qclaogui/golang-api-server/pkg/api/routeguidepb"
	pbtodov1 "github.com/qclaogui/golang-api-server/pkg/api/todopb/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunServer runs gRPC service to publish service
func RunServer(
	ctx context.Context,
	toDoSrv pbtodov1.ToDoServiceServer,
	routeGuideSrv pbrouteguide.RouteGuideServer,
	port string,
) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server startup options
	var opts []grpc.ServerOption
	// add middleware
	// opts = middleware.AddLogging(logger.Log, opts)

	// register service
	srv := grpc.NewServer(opts...)

	//	register service
	pbtodov1.RegisterToDoServiceServer(srv, toDoSrv)
	pbrouteguide.RegisterRouteGuideServer(srv, routeGuideSrv)

	// Register reflection service on gRPC server.
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
	slog.Info("starting gRPC server...", "grpc_port", port)
	return srv.Serve(listen)
}
