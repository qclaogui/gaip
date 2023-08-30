// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package middleware

import (
	"context"
	"log/slog"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

// AddLogging returns grpc.Server config option that turn on logging.
func AddLogging(opts []grpc.ServerOption) []grpc.ServerOption {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	lo := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}

	// Add unary interceptor
	opts = append(opts, grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(InterceptorLogger(logger), lo...),
	))

	// Add stream interceptor
	opts = append(opts, grpc.ChainStreamInterceptor(
		logging.StreamServerInterceptor(InterceptorLogger(logger), lo...),
	))

	return opts
}
