// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"github.com/go-kit/log/level"
	"github.com/qclaogui/gaip/genproto/generativelanguage/apiv1/generativelanguagepb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service/generativeai"
)

func (g *Gaip) initGenerativeAI() error {
	if !g.Cfg.GenaiCfg.Enabled {
		_ = level.Warn(g.Server.Log).Log("msg", "genai.enabled=false")
		return nil
	}

	g.Cfg.GenaiCfg.Log = g.Server.Log
	g.Cfg.GenaiCfg.Registerer = g.Registerer

	repo, err := repository.NewGenerativeAI(g.Cfg.GenaiCfg)
	if err != nil {
		return err
	}
	g.Cfg.GenaiCfg.RepoGenerative = repo

	srv, err := generativeai.NewServer(g.Cfg.GenaiCfg)
	if err != nil {
		return err
	}

	// Register Service Server
	generativelanguagepb.RegisterGenerativeServiceServer(g.Server.GRPCServer, srv)

	return nil
}
