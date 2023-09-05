// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"context"
	"io/fs"
	"log/slog"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbtodov1 "github.com/qclaogui/golang-api-server/api/gen/proto/todo/v1"
	"github.com/qclaogui/golang-api-server/pkg/protocol/grpc/interceptors"
	"github.com/qclaogui/golang-api-server/pkg/protocol/rest/middleware"
	"github.com/qclaogui/golang-api-server/third_party"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, grpcPort, port string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials())) //TODO(qc) move to interceptors
	opts = interceptors.RegisterGRPCDailOption(opts...)

	// The address to the gRPC server, in the gRPC standard naming format.
	// See https://github.com/grpc/grpc/blob/master/doc/naming.md for more information.
	endpoint := "dns:///0.0.0.0:" + grpcPort

	gwmux := runtime.NewServeMux()

	// Register the gRPC server's handler with the HTTP gwmux
	err := pbtodov1.RegisterToDoServiceHandlerFromEndpoint(ctx, gwmux, endpoint, opts)
	if err != nil {
		slog.Error("failed to start HTTP gateway", "error", err)
		return err
	}

	// Register middleware chain
	handler := middleware.WrapperHandler(gwmux)

	openAPI := getOpenAPIHandler()
	// Set up the REST server on port cfg.HTTPPort and handle requests by proxying them to the gRPC server
	srv := &http.Server{
		Addr: ":" + port,
		// Handler: handler,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/api") {
				handler.ServeHTTP(w, r)
				return
			}
			openAPI.ServeHTTP(w, r)
		}),
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			slog.Warn("Graceful shutting down HTTP/REST gateway...")
		}
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	slog.Warn("starting HTTP/REST gateway...", "http_port", port)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return srv.ListenAndServe()
}

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	_ = mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "gen/openapiv2")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}
