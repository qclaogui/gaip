// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package service

import "google.golang.org/grpc"

// Backend abstracts a registrable GRPC service.
type Backend interface {
	RegisterGRPC(*grpc.Server)
}
