// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package routeguide

import (
	pb "github.com/qclaogui/golang-api-server/api/routeguide/v1/routeguidepb"
)

type Repository interface {
	pb.RouteGuideServiceServer
}
