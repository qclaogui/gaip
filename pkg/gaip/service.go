// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"context"
	"log/slog"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/protocol/grpc/interceptors"
	"github.com/qclaogui/gaip/pkg/service/bookstore"
	"github.com/qclaogui/gaip/pkg/service/library"
	"github.com/qclaogui/gaip/pkg/service/project"
	"github.com/qclaogui/gaip/pkg/service/routeguide"
	"github.com/qclaogui/gaip/pkg/service/todo"
	"github.com/qclaogui/gaip/pkg/vault"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (g *Gaip) initBookstore() error {
	if !g.Cfg.BookstoreCfg.Enabled {
		return nil
	}

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

func (g *Gaip) initLibrary() error {
	if !g.Cfg.LibraryCfg.Enabled {
		return nil
	}

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

	//g.RegisterRoute("/library/healthz", g.healthzHandler(), false, http.MethodGet)
	return nil
}

func (g *Gaip) initProject() error {
	if !g.Cfg.ProjectCfg.Enabled {
		return nil
	}

	g.Cfg.ProjectCfg.Registerer = g.Registerer

	repoProject, err := repository.NewProject(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}
	g.Cfg.ProjectCfg.RepoProject = repoProject

	repoIdentity, err := repository.NewIdentity(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}
	g.Cfg.ProjectCfg.RepoIdentity = repoIdentity

	srv, err := project.NewServer(g.Cfg.ProjectCfg)
	if err != nil {
		return err
	}

	// Register Service Server
	projectpb.RegisterProjectServiceServer(g.Server.GRPCServer, srv)
	projectpb.RegisterIdentityServiceServer(g.Server.GRPCServer, srv)
	projectpb.RegisterEchoServiceServer(g.Server.GRPCServer, srv)

	// Register routes
	//g.RegisterRoute("/v1/echo:echo", rest.HandleEcho(), false, true, http.MethodPost)
	return nil
}

func (g *Gaip) initRouteGuide() error {
	if !g.Cfg.RouteGuideCfg.Enabled {
		return nil
	}

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

	// Register Service Server
	routeguidepb.RegisterRouteGuideServiceServer(g.Server.GRPCServer, srv)

	return nil
}

func (g *Gaip) initTodo() error {
	if !g.Cfg.TodoCfg.Enabled {
		return nil
	}

	repo, err := repository.NewTodo(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	// Update the config.
	g.Cfg.TodoCfg.Repo = repo
	g.Cfg.TodoCfg.Registerer = g.Registerer

	srv, err := todo.NewServer(g.Cfg.TodoCfg)
	if err != nil {
		return err
	}

	// Register GRPC Server
	todopb.RegisterToDoServiceServer(g.Server.GRPCServer, srv)

	// Register HTTP/REST gateway
	opts := interceptors.RegisterGRPCDailOption()
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	gwmux := runtime.NewServeMux()

	// Register the gRPC server's handler with the Router gwmux
	ctx := context.Background()
	err = todopb.RegisterToDoServiceHandlerFromEndpoint(ctx, gwmux, g.Server.GRPCListenAddr().String(), opts)
	if err != nil {
		slog.Error("failed to start Router gateway", "error", err)
		return err
	}

	// https://grpc-ecosystem.github.io/grpc-gateway/docs/mapping/grpc_api_configuration/#generate_unbound_methods
	// URI path is built from the service’s name and method: /<fully qualified service name>/<method name> (e.g.: /my.package.EchoService/Echo)

	// https://cloud.google.com/endpoints/docs/grpc/transcoding?hl=zh-cn#where_to_configure_transcoding

	// Set up the REST server and handle requests by proxying them to the gRPC server
	g.Server.Router.PathPrefix("/v1/todos").Handler(gwmux)

	//slog.Warn("starting Router/REST gateway...", "http_listen_addr", server.HTTPListenAddr().String())

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
