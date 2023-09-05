// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package interceptors

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ServerOptionRecovery returns grpc.Server config option that turn on panic recovery.
func ServerOptionRecovery(serverOpts []grpc.ServerOption) []grpc.ServerOption {
	// Define customFn to handle panic
	customFn := func(p any) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	// Shared options for the logger, with a custom gRPC code to log level function.
	opts := []recovery.Option{
		recovery.WithRecoveryHandler(customFn),
	}

	// Add unary interceptor
	serverOpts = append(serverOpts, grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(opts...),
	))
	// Add stream interceptor
	serverOpts = append(serverOpts, grpc.ChainStreamInterceptor(
		recovery.StreamServerInterceptor(opts...),
	))

	return serverOpts
}
