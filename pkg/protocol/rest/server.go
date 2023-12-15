// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/pkg/protocol/grpc/interceptors"
	"github.com/qclaogui/gaip/pkg/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RunRESTServer runs Router/REST gateway
func RunRESTServer(ctx context.Context, server *service.Server) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())) //TODO(qc) move to interceptors
	opts = interceptors.RegisterGRPCDailOption(opts...)

	// The address to the gRPC server, in the gRPC standard naming format.
	// See https://github.com/grpc/grpc/blob/master/doc/naming.md for more information.
	//endpoint := fmt.Sprintf("dns:///0.0.0.0:%d", server.GRPCListenPort)

	gwmux := runtime.NewServeMux()

	// Register the gRPC server's handler with the Router gwmux
	err := todopb.RegisterToDoServiceHandlerFromEndpoint(ctx, gwmux, server.GRPCListenAddr().String(), opts)
	if err != nil {
		slog.Error("failed to start Router gateway", "error", err)
		return err
	}

	// Set up the REST server and handle requests by proxying them to the gRPC server
	server.Router.PathPrefix("").Handler(gwmux)

	//slog.Warn("starting Router/REST gateway...", "http_listen_addr", server.HTTPListenAddr().String())
	return nil
}
