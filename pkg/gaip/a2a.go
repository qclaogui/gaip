// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"github.com/go-kit/log/level"
	"github.com/qclaogui/gaip/genproto/a2a/apiv1/a2apb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service/a2a"
)

func (g *Gaip) initA2A() error {
	if !g.Cfg.A2ACfg.Enabled {
		_ = level.Warn(g.Server.Log).Log("msg", "a2a.enabled=false")
		return nil
	}

	g.Cfg.A2ACfg.Log = g.Server.Log
	g.Cfg.A2ACfg.Registerer = g.Registerer

	repo, err := repository.NewA2A(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}
	g.Cfg.A2ACfg.Repo = repo

	srv, err := a2a.NewServer(g.Cfg.A2ACfg)
	if err != nil {
		return err
	}

	a2apb.RegisterA2AServiceServer(g.Server.GRPCServer, srv)
	return nil
}
