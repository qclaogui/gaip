// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package routeguide

import (
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
)

type Repository interface {
	routeguidepb.RouteGuideServiceServer
}
