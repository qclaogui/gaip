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
	"github.com/qclaogui/golang-api-server/pkg/protocol/grpc/interceptors"
	"github.com/qclaogui/golang-api-server/pkg/protocol/rest"
	routeguidev1 "github.com/qclaogui/golang-api-server/pkg/service/routeguide/v1"
	todov1 "github.com/qclaogui/golang-api-server/pkg/service/todo/v1"
	"github.com/qclaogui/golang-api-server/pkg/vault"
	util_log "github.com/qclaogui/golang-api-server/tools/log"
	"gopkg.in/yaml.v3"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/qclaogui/golang-api-server/pkg/protocol/grpc"
)

type Application struct {
	Cfg        Config
	Registerer prometheus.Registerer

	Server *server.Server

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

// New makes a new Application.
func New(cfg Config, reg prometheus.Registerer) (*Application, error) {
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

	// Application init
	if err := app.initVault(); err != nil {
		return err
	}

	// flag.StringVar(&cfg.DBHost, "db-host", "127.0.0.1", "Database host")
	// flag.StringVar(&cfg.DBUser, "db-user", "root", "Database user")
	// flag.StringVar(&cfg.DBPassword, "db-password", "", "Database password")
	// flag.StringVar(&cfg.DBSchema, "db-schema", "dev", "Database schema")
	// flag.IntVar(&cfg.LogLevel, "log-level", 0, "Global log level")

	// Initialize tracing and handle the tracer provider shutdown
	stopTracing := interceptors.InitTracing()
	defer stopTracing()

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	//param := "parseTime=true"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBSchema, param)
	//toDoSrv, err := todov1.NewServiceServer(todov1.WithMysqlRepository(dsn))

	toDoSrv, err := todov1.NewServiceServer(util_log.Logger, todov1.WithMemoryRepository())
	if err != nil {
		return err
	}
	routeGuideSrv, err := routeguidev1.NewServiceServer(routeguidev1.WithMemoryRepository())
	if err != nil {
		return err
	}

	// Start the REST server in goroutine
	go func() {
		err = rest.RunRESTServer(ctx, app.Cfg.Server)
		util_log.CheckFatal("running REST server", err)
	}()

	return grpc.RunGRPCServer(ctx, app.Cfg.Server, toDoSrv, routeGuideSrv)
}
