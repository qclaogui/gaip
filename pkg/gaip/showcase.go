// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"net/http"

	"cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/go-kit/log/level"
	"github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb"
	"github.com/qclaogui/gaip/genproto/showcase/apiv1beta1/showcasepb/genrest"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service/showcase"
)

func (g *Gaip) initShowcase() error {
	if !g.Cfg.ShowcaseCfg.Enabled {
		_ = level.Warn(g.Server.Log).Log("msg", "showcase.enabled=false")
		return nil
	}

	g.Cfg.ShowcaseCfg.Log = g.Server.Log
	g.Cfg.ShowcaseCfg.Registerer = g.Registerer

	repoIdentity, err := repository.NewIdentity(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}
	g.Cfg.ShowcaseCfg.RepoIdentity = repoIdentity

	repoMessaging, err := repository.NewMessaging(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}
	g.Cfg.ShowcaseCfg.RepoMessaging = repoMessaging

	srv, err := showcase.NewServer(g.Cfg.ShowcaseCfg)
	if err != nil {
		return err
	}

	showcasepb.RegisterEchoServiceServer(g.Server.GRPCServer, srv)
	genrest.RegisterHandlersEchoService(g.Server.Router, srv, g.Server.Log)

	showcasepb.RegisterIdentityServiceServer(g.Server.GRPCServer, srv)
	genrest.RegisterHandlersIdentityService(g.Server.Router, srv, g.Server.Log)

	showcasepb.RegisterMessagingServiceServer(g.Server.GRPCServer, srv)
	genrest.RegisterHandlersMessagingService(g.Server.Router, srv, g.Server.Log)

	// FATAL: [core] grpc: Server.RegisterService found duplicate service registration for "google.longrunning.Operations"
	// Register OperationsServer
	longrunningpb.RegisterOperationsServer(g.Server.GRPCServer, srv)
	g.RegisterRoute("/v1beta1/operations", srv.HandleListOperations(), false, http.MethodGet)
	g.RegisterRoute("/v1beta1/{name:operations/[^:]+}", srv.HandleGetOperation(), false, http.MethodGet)
	g.RegisterRoute("/v1beta1/{name:operations/[^:]+}", srv.HandleDeleteOperation(), false, http.MethodDelete)
	g.RegisterRoute("/v1beta1/{name:operations/[^:]+}:cancel", srv.HandleCancelOperation(), false, http.MethodPost)

	return nil
}
