// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"github.com/go-kit/log/level"
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service/bookstore"
)

func (g *Gaip) initBookstore() error {
	if !g.Cfg.BookstoreCfg.Enabled {
		_ = level.Warn(g.Server.Log).Log("msg", "bookstore.enabled=false")
		return nil
	}

	g.Cfg.BookstoreCfg.Log = g.Server.Log
	g.Cfg.BookstoreCfg.Registerer = g.Registerer

	repo, err := repository.NewBookstore(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	g.Cfg.BookstoreCfg.Repo = repo

	srv, err := bookstore.NewServer(g.Cfg.BookstoreCfg)
	if err != nil {
		return err
	}

	// Register Service Server
	bookstorepb.RegisterBookstoreServiceServer(g.Server.GRPCServer, srv)

	return nil
}
