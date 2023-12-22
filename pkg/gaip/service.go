// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"net/http"

	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service/bookstore"
	"github.com/qclaogui/gaip/pkg/service/library"
	"github.com/qclaogui/gaip/pkg/service/project"
	projectrepository "github.com/qclaogui/gaip/pkg/service/project/repository"
	"github.com/qclaogui/gaip/pkg/service/routeguide"
	"github.com/qclaogui/gaip/pkg/service/todo"
	"github.com/qclaogui/gaip/pkg/vault"
)

func (g *Gaip) initBookstore() error {
	if !g.Cfg.BookstoreCfg.Enabled {
		return nil
	}

	repo, err := repository.NewBookstore(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	g.Cfg.BookstoreCfg.Repo = repo
	if _, err = bookstore.New(g.Cfg.BookstoreCfg, g.Server); err != nil {
		return err
	}

	return nil
}

func (g *Gaip) initLibrary() error {
	if !g.Cfg.LibraryCfg.Enabled {
		return nil
	}

	repo, err := repository.NewLibrary(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	g.Cfg.LibraryCfg.Repo = repo
	if _, err = library.New(g.Cfg.LibraryCfg, g.Server); err != nil {
		return err
	}

	g.RegisterRoute("/library/healthz", g.healthzHandler(), false, true, http.MethodGet)
	return nil
}

func (g *Gaip) initProject() error {
	if !g.Cfg.ProjectCfg.Enabled {
		return nil
	}

	repo, err := projectrepository.NewProject(g.Cfg.ProjectCfg.RepoCfg)
	if err != nil {
		return err
	}
	g.Cfg.ProjectCfg.Repo = repo
	if _, err = project.New(g.Cfg.ProjectCfg, g.Server); err != nil {
		return err
	}

	return nil
}

func (g *Gaip) initRouteGuide() error {
	if !g.Cfg.RouteGuideCfg.Enabled {
		return nil
	}

	repo, err := repository.NewRouteGuide(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	g.Cfg.RouteGuideCfg.Repo = repo
	if _, err = routeguide.New(g.Cfg.RouteGuideCfg, g.Server); err != nil {
		return err
	}

	return nil
}

func (g *Gaip) initTodo() error {
	repo, err := repository.NewTodo(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	g.Cfg.TodoCfg.Repo = repo
	if _, err = todo.New(g.Cfg.TodoCfg, g.Server); err != nil {
		return err
	}

	return nil
}

// initVault init Vault
func (g *Gaip) initVault() error {
	if !g.Cfg.VaultCfg.Enabled {
		return nil
	}

	v, err := vault.New(g.Cfg.VaultCfg)
	if err != nil {
		return err
	}
	g.Vault = v

	// Update Configs - KVStore
	//g.Cfg.MemberlistKV.TCPTransport.TLS.Reader = g.VaultCfg

	// Update Configs - GRPCServer Clients
	//g.Cfg.Worker.GRPCClientConfig.TLS.Reader = g.VaultCfg

	return nil
}
