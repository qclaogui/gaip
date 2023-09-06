// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package interceptors

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// ServerOptionCredentials ServerOption Credentials
func ServerOptionCredentials(serverOpts []grpc.ServerOption) []grpc.ServerOption {
	return serverOpts
}

// WithDailOptionCredentials DialOption Credentials
func WithDailOptionCredentials(dialOpts []grpc.DialOption) []grpc.DialOption {
	dialOpts = append(dialOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	return dialOpts
}
