// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"github.com/go-kit/log/level"
	"github.com/qclaogui/gaip/genproto/task/apiv1/taskpb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service/task"
)

func (g *Gaip) initTask() error {
	if !g.Cfg.TaskCfg.Enabled {
		_ = level.Warn(g.Server.Log).Log("msg", "task.enabled=false")
		return nil
	}

	g.Cfg.TaskCfg.Log = g.Server.Log
	g.Cfg.TaskCfg.Registerer = g.Registerer

	repo, err := repository.NewTasks(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}
	g.Cfg.TaskCfg.Repo = repo

	srv, err := task.NewServer(g.Cfg.TaskCfg)
	if err != nil {
		return err
	}

	taskpb.RegisterTasksServiceServer(g.Server.GRPCServer, srv)

	return nil
}
