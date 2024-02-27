// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"net/http"

	"github.com/go-kit/log/level"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb/genrest"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service/library"
)

func (g *Gaip) initLibrary() error {
	if !g.Cfg.LibraryCfg.Enabled {
		_ = level.Warn(g.Server.Log).Log("msg", "library.enabled=false")
		return nil
	}

	g.Cfg.LibraryCfg.Log = g.Server.Log
	g.Cfg.LibraryCfg.Registerer = g.Registerer

	repo, err := repository.NewLibrary(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}
	g.Cfg.LibraryCfg.Repo = repo

	srv, err := library.NewServer(g.Cfg.LibraryCfg)
	if err != nil {
		return err
	}

	// Register Service Server
	librarypb.RegisterLibraryServiceServer(g.Server.GRPCServer, srv)
	genrest.RegisterHandlersLibraryService(g.Server.Router, srv, g.Server.Log)

	g.RegisterRoute("/library/healthz", g.healthzHandler(), false, http.MethodGet)
	return nil
}
