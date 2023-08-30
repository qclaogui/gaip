// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package middleware

import (
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AddRecovery(opts []grpc.ServerOption) []grpc.ServerOption {
	// Define customFn to handle panic
	customFn := func(p any) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	// Shared options for the logger, with a custom gRPC code to log level function.
	ro := []recovery.Option{
		recovery.WithRecoveryHandler(customFn),
	}

	// Add unary interceptor
	opts = append(opts, grpc.ChainUnaryInterceptor(
		recovery.UnaryServerInterceptor(ro...),
	))

	// Add stream interceptor
	opts = append(opts, grpc.ChainStreamInterceptor(
		recovery.StreamServerInterceptor(ro...),
	))

	return opts
}
