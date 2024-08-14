// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"context"

	"github.com/go-kit/log/level"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service/todo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (g *Gaip) initTodo() error {
	if !g.Cfg.TodoCfg.Enabled {
		_ = level.Warn(g.Server.Log).Log("msg", "todo.enabled=false")
		return nil
	}

	g.Cfg.TodoCfg.Log = g.Server.Log
	g.Cfg.TodoCfg.Registerer = g.Registerer

	repo, err := repository.NewTodo(g.Cfg.RepoCfg)
	if err != nil {
		return err
	}

	g.Cfg.TodoCfg.Repo = repo

	srv, err := todo.NewServer(g.Cfg.TodoCfg)
	if err != nil {
		return err
	}

	// Register Services to the GRPCServer.
	todopb.RegisterToDoServiceServer(g.Server.GRPCServer, srv)

	// Register HTTP/REST gateway
	var opts []grpc.DialOption
	//opts = interceptors.RegisterGRPCDailOption()
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	ctx := context.Background()

	// Register the gRPC server's handler with the Router gwmux
	gwmux := runtime.NewServeMux()
	err = todopb.RegisterToDoServiceHandlerServer(ctx, gwmux, srv)

	_ = opts
	// Note: Make sure the gRPC server is running properly and accessible
	//err = todopb.RegisterToDoServiceHandlerFromEndpoint(ctx, gwmux, g.Server.GRPCListenAddr().String(), opts)
	if err != nil {
		return err
	}

	// Set up the REST server and handle requests by proxying them to the gRPC server
	//g.Server.Router.PathPrefix("/v1/todos").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	_ = level.Info(g.Server.Log).Log("msg", "Incoming OPTIONS for request", "request_uri", r.RequestURI)
	//	w.Header().Add("access-control-allow-credentials", "true")
	//	w.Header().Add("access-control-allow-headers", "*")
	//	w.Header().Add("access-control-allow-methods", http.MethodPost)
	//	w.Header().Add("access-control-allow-origin", "*")
	//	w.Header().Add("access-control-max-age", "3600")
	//	w.WriteHeader(http.StatusOK)
	//}).Methods(http.MethodOptions)

	g.Server.Router.PathPrefix("/v1/todos").Handler(gwmux)

	return nil
}
