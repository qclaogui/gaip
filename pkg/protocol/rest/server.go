// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package rest

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pbtodov1 "github.com/qclaogui/golang-api-server/api/gen/proto/todo/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, grpcPort, port string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gwmux := runtime.NewServeMux()

	// gRPC client options
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// Use the OpenTelemetry gRPC client interceptor for tracing
	// opts = append(opts, grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))

	// Register the gRPC server's handler with the HTTP gwmux
	err := pbtodov1.RegisterToDoServiceHandlerFromEndpoint(ctx, gwmux, "localhost:"+grpcPort, opts)
	if err != nil {
		slog.Error("failed to start HTTP gateway", "error", err)
		return err
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: gwmux,
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
	return srv.ListenAndServe()
}
