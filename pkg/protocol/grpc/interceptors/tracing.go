// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package interceptors

import (
	"context"
	"log"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc"
)

// ServerOptionTracing Use the OpenTelemetry gRPC server interceptor for tracing
func ServerOptionTracing(serverOpts []grpc.ServerOption) []grpc.ServerOption {

	serverOpts = append(serverOpts, grpc.StatsHandler(otelgrpc.NewServerHandler()))

	return serverOpts
}

// WithDailOptionTracing Use the OpenTelemetry gRPC client interceptor for tracing
func WithDailOptionTracing(dialOpts []grpc.DialOption) []grpc.DialOption {

	dialOpts = append(dialOpts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))

	return dialOpts
}

// InitTracing OpenTelemetry tracing and return a function to stop the tracer provider
func InitTracing() func() {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("failed to create stdout exporter: %v", err)
	}

	// Create a simple span processor that writes to the exporter
	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(bsp))
	otel.SetTracerProvider(tp)

	// Set the global propagator to use W3C Trace Context
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Return a function to stop the tracer provider
	return func() {
		if err = tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shut down tracer provider: %v", err)
		}
	}
}
