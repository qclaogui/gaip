// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package interceptors

import (
	"google.golang.org/grpc"
)

// ServerOptionTLS TLSCredentials
func ServerOptionTLS(serverOpts []grpc.ServerOption) []grpc.ServerOption {
	return serverOpts
}

func WithDailOptionTLS(dialOpts []grpc.DialOption) []grpc.DialOption {
	return dialOpts
}
