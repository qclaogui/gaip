// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package interceptors

import "google.golang.org/grpc"

// RegisterGRPCServerOption returns grpc.Server config option that turn on.
func RegisterGRPCServerOption(serverOpts ...grpc.ServerOption) []grpc.ServerOption {
	serverOpts = ServerOptionLogging(serverOpts)
	serverOpts = ServerOptionRecovery(serverOpts)

	return serverOpts
}

// RegisterGRPCDailOption returns grpc.Dial config option that turn on.
func RegisterGRPCDailOption(dailOpts ...grpc.DialOption) []grpc.DialOption {
	dailOpts = WithDailOptionLogging(dailOpts)
	dailOpts = WithDailOptionTracing(dailOpts)
	dailOpts = WithDailOptionCredentials(dailOpts)

	return dailOpts
}
