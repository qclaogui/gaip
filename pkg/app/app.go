// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package app

import (
	"context"
	"fmt"
	"os"

	"github.com/grafana/dskit/server"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/pkg/protocol/grpc/interceptors"
	"github.com/qclaogui/gaip/pkg/protocol/rest"
	"github.com/qclaogui/gaip/pkg/service/bookstore"
	"github.com/qclaogui/gaip/pkg/service/library"
	"github.com/qclaogui/gaip/pkg/service/project"
	"github.com/qclaogui/gaip/pkg/service/routeguide"
	"github.com/qclaogui/gaip/pkg/service/todo"
	"github.com/qclaogui/gaip/pkg/vault"
	lg "github.com/qclaogui/gaip/tools/log"
	"gopkg.in/yaml.v3"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/qclaogui/gaip/pkg/protocol/grpc"
)

type Application struct {
	Cfg        Config
	Registerer prometheus.Registerer

	Server *server.Server

	BookstoreServer bookstore.Server
	LibraryServer   library.Service
	ProjectServer   project.Service

	Vault *vault.Vault
}

// initVault init Vault
func (app *Application) initVault() error {
	if !app.Cfg.Vault.Enabled {
		return nil
	}

	v, err := vault.New(app.Cfg.Vault)
	if err != nil {
		return err
	}
	app.Vault = v

	// Update Configs - KVStore
	//app.Cfg.MemberlistKV.TCPTransport.TLS.Reader = app.Vault

	// Update Configs - GRPC Clients
	//app.Cfg.Worker.GRPCClientConfig.TLS.Reader = app.Vault

	return nil
}

func (app *Application) initBookstore() (bookstore.Server, error) {
	srv, err := bookstore.NewBookstoreServer(app.Cfg.Bookstore)
	if err != nil {
		return nil, err
	}

	app.BookstoreServer = srv
	return srv, nil
}

func (app *Application) initLibrary() (library.Service, error) {
	srv, err := library.NewLibraryService(app.Cfg.Library)
	if err != nil {
		return nil, err
	}

	app.LibraryServer = srv
	return srv, nil
}

func (app *Application) initProject() (project.Service, error) {
	srv, err := project.NewProjectService(app.Cfg.Project)
	if err != nil {
		return nil, err
	}

	app.ProjectServer = srv
	return srv, nil
}

// NewApplication makes a new Application.
func NewApplication(cfg Config, reg prometheus.Registerer) (*Application, error) {
	if cfg.PrintConfig {
		if err := yaml.NewEncoder(os.Stdout).Encode(&cfg); err != nil {
			fmt.Println("Error encoding config:", err)
		}
		os.Exit(0)
	}

	app := &Application{
		Cfg:        cfg,
		Registerer: reg,
	}

	return app, nil
}

// Bootstrap bootstrap gRPC server and HTTP gateway
func (app *Application) Bootstrap() error {
	ctx := context.Background()

	// Initialize tracing and handle the tracer provider shutdown
	stopTracing := interceptors.InitTracing()
	defer stopTracing()

	// Application init
	if err := app.initVault(); err != nil {
		return err
	}

	toDoSrv, err := todo.NewServiceServer(lg.Logger, todo.WithMemoryRepository())
	if err != nil {
		return err
	}
	routeGuideSrv, err := routeguide.NewServiceServer(lg.Logger, routeguide.WithMemoryRepository())
	if err != nil {
		return err
	}

	if _, err = app.initBookstore(); err != nil {
		return err
	}

	if _, err = app.initLibrary(); err != nil {
		return err
	}

	if _, err = app.initProject(); err != nil {
		return err
	}

	// Start the REST server in goroutine
	go func() {
		err = rest.RunRESTServer(ctx, app.Cfg.Server)
		lg.CheckFatal("running REST server", err)
	}()

	return grpc.RunGRPCServer(
		ctx,
		app.Cfg.Server,
		toDoSrv,
		routeGuideSrv,
		app.BookstoreServer,
		app.LibraryServer,
		app.ProjectServer,
	)
}
