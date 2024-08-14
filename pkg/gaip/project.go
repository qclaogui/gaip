// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"github.com/go-kit/log/level"
	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service/project"
)

func (g *Gaip) initProject() error {
	if !g.Cfg.ProjectCfg.Enabled {
		_ = level.Warn(g.Server.Log).Log("msg", "project.enabled=false")
		return nil
	}

	g.Cfg.ProjectCfg.Log = g.Server.Log
	g.Cfg.ProjectCfg.Registerer = g.Registerer

	repoProject, err := repository.NewProject(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}
	g.Cfg.ProjectCfg.RepoProject = repoProject

	srv, err := project.NewServer(g.Cfg.ProjectCfg)
	if err != nil {
		return err
	}

	// Register ProjectServiceServer
	projectpb.RegisterProjectServiceServer(g.Server.GRPCServer, srv)

	return nil
}
