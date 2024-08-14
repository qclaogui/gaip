// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"github.com/go-kit/log/level"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service/routeguide"
)

func (g *Gaip) initRouteGuide() error {
	if !g.Cfg.RouteGuideCfg.Enabled {
		_ = level.Warn(g.Server.Log).Log("msg", "routeguide.enabled=false")
		return nil
	}

	g.Cfg.RouteGuideCfg.Log = g.Server.Log
	g.Cfg.RouteGuideCfg.Registerer = g.Registerer

	repo, err := repository.NewRouteGuide(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}
	g.Cfg.RouteGuideCfg.Repo = repo

	srv, err := routeguide.NewServer(g.Cfg.RouteGuideCfg)
	if err != nil {
		return err
	}

	// Register Services to the GRPCServer.
	routeguidepb.RegisterRouteGuideServiceServer(g.Server.GRPCServer, srv)

	return nil
}
