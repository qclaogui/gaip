// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/qclaogui/golang-api-server/pkg/protocol/grpc/interceptors"
	routeguidev1 "github.com/qclaogui/golang-api-server/pkg/service/routeguide/v1"
	todov1 "github.com/qclaogui/golang-api-server/pkg/service/todo/v1"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/qclaogui/golang-api-server/pkg/protocol/grpc"
	"github.com/qclaogui/golang-api-server/pkg/protocol/rest"
)

// Config is configuration for Server
type Config struct {
	// gRPC server start parameters section
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string

	// HTTP/REST gateway start parameters section
	// HTTPPort is TCP port to listen by HTTP/REST gateway
	HTTPPort string

	// DB Datastore parameters section
	// DBHost is host of database
	DBHost string
	// DBUser is username to connect to database
	DBUser string
	// DBPassword password to connect to database
	DBPassword string
	// DBSchema is schema of database
	DBSchema string

	// Log parameters section
	// LogLevel is global log level: Debug(-1), Info(0), Warn(1), Error(2), DPanic(3), Panic(4), Fatal(5)
	LogLevel int
}

// Bootstrap bootstrap gRPC server and HTTP gateway
func Bootstrap() error {
	ctx := context.Background()
	var cfg Config
	flag.StringVar(&cfg.GRPCPort, "grpc-port", "9095", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "8080", "http port to bind")
	flag.StringVar(&cfg.DBHost, "db-host", "127.0.0.1", "Database host")
	flag.StringVar(&cfg.DBUser, "db-user", "root", "Database user")
	flag.StringVar(&cfg.DBPassword, "db-password", "", "Database password")
	flag.StringVar(&cfg.DBSchema, "db-schema", "dev", "Database schema")
	flag.IntVar(&cfg.LogLevel, "log-level", 0, "Global log level")
	flag.Parse()

	// Initialize tracing and handle the tracer provider shutdown
	stopTracing := interceptors.InitTracing()
	defer stopTracing()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("invalid TCP port for HTTP gateway: '%s'", cfg.HTTPPort)
	}

	// add MySQL driver specific parameter to parse date/time
	// Drop it for another database
	//param := "parseTime=true"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBSchema, param)
	//toDoSrv, err := todov1.NewServiceServer(todov1.WithMysqlRepository(dsn))

	toDoSrv, err := todov1.NewServiceServer(todov1.WithMemoryRepository())
	if err != nil {
		return err
	}
	routeGuideSrv, err := routeguidev1.NewServiceServer(routeguidev1.WithMemoryRepository())
	if err != nil {
		return err
	}

	// Start the REST server in a goroutine
	go func() {
		if err = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort); err != nil {
			log.Fatalf("Failed to run REST server: %v", err)
		}
	}()

	// run gRPC server
	return grpc.RunServer(
		ctx,
		toDoSrv,
		routeGuideSrv,
		cfg.GRPCPort,
	)
}
