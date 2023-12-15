// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package grpc

// RunGRPCServer runs gRPC service to publish service
func RunGRPCServer() error {

	// gRPC server startup options
	//grpcServer := grpc.NewServer(interceptors.RegisterGRPCServerOption()...)
	////	register backend service
	//for _, backend := range backends {
	//	backend.RegisterGRPC(grpcServer)
	//}

	// Register reflection service on gRPC server.
	// Enable reflection to allow clients to query the server's services
	//reflection.Register(grpcServer)

	// start gRPC server
	//slog.Warn("starting gRPC server...", "grpc_port", cfg.GRPCListenPort)
	//return grpcServer.Serve(listen)
	return nil
}
