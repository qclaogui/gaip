// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-kit/log/level"
	"github.com/googleapis/gapic-showcase/util/genrest/resttools"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/internal/repository"
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

	g.RegisterRoute("/library/healthz", g.healthzHandler(), false, http.MethodGet)
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

	// Register Services to the GRPCServer.
	projectpb.RegisterProjectServiceServer(g.Server.GRPCServer, srv)
	projectpb.RegisterIdentityServiceServer(g.Server.GRPCServer, srv)
	projectpb.RegisterEchoServiceServer(g.Server.GRPCServer, srv)

	// Register routes
	g.RegisterRoute("/v1/echo:echo", g.HandleEcho(srv), false, http.MethodPost)
	return nil
}

// HandleEcho translates REST requests/responses on the wire to internal proto messages for Echo
//
//	HTTP binding pattern: POST "/v1/echo:echo"
func (g *Gaip) HandleEcho(srv *project.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = level.Info(g.Server.Log).Log("msg", "[HandleEcho] received request")

		urlPathParams := mux.Vars(r)
		numURLPathParams := len(urlPathParams)
		_ = level.Info(g.Server.Log).Log("msg", fmt.Sprintf("urlPathParams (expect 0, have %d): %q", numURLPathParams, urlPathParams))

		if numURLPathParams != 0 {
			return
		}

		systemParameters, queryParams, err := resttools.GetSystemParameters(r)
		if err != nil {
			return
		}

		request := &projectpb.EchoRequest{}
		// Intentional: Field values in the URL path override those set in the body.
		var jsonReader bytes.Buffer
		bodyReader := io.TeeReader(r.Body, &jsonReader)
		rBytes := make([]byte, r.ContentLength)
		_, err = bodyReader.Read(rBytes)
		if err != nil && !errors.Is(err, io.EOF) {
			_ = level.Error(g.Server.Log).Log("msg", "error reading body content", "error", err)
			return
		}

		if err = resttools.FromJSON().Unmarshal(rBytes, request); err != nil {
			_ = level.Error(g.Server.Log).Log("msg", "error reading body params", "error", err)
			return
		}

		if err = resttools.CheckRequestFormat(&jsonReader, r, request.ProtoReflect()); err != nil {
			_ = level.Error(g.Server.Log).Log("msg", "REST request failed format check", "error", err)
			return
		}

		if len(queryParams) > 0 {
			_ = level.Error(g.Server.Log).Log("msg", "encountered unexpected query params", "params", queryParams)
			return
		}
		if err = resttools.PopulateSingularFields(request, urlPathParams); err != nil {
			_ = level.Error(g.Server.Log).Log("msg", "error reading URL path params", "error", err)
			return
		}

		marshaler := resttools.ToJSON()
		marshaler.UseEnumNumbers = systemParameters.EnumEncodingAsInt
		requestJSON, _ := marshaler.Marshal(request)
		_ = level.Info(g.Server.Log).Log("msg", fmt.Sprintf("request: %s", requestJSON))

		ctx := context.WithValue(r.Context(), resttools.BindingURIKey, "/v1/echo:echo")

		response, err := srv.Echo(ctx, request)
		if err != nil {
			//backend.ReportGRPCError(w, err)
			return
		}

		json, err := marshaler.Marshal(response)
		if err != nil {
			_ = level.Info(g.Server.Log).Log("msg", "error json-encoding response", "error", err)
			return
		}

		_, _ = w.Write(json)
	}
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

	// Register Services to the GRPCServer.
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

	// Register Services to the GRPCServer.
	todopb.RegisterToDoServiceServer(g.Server.GRPCServer, srv)

	// Register HTTP/REST gateway
	var opts []grpc.DialOption
	//opts = interceptors.RegisterGRPCDailOption()
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	gwmux := runtime.NewServeMux()

	// Register the gRPC server's handler with the Router gwmux
	ctx := context.Background()
	//err = todopb.RegisterToDoServiceHandlerServer(ctx, gwmux, srv)
	err = todopb.RegisterToDoServiceHandlerFromEndpoint(ctx, gwmux, g.Server.GRPCListenAddr().String(), opts)
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
