// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package interceptors

import (
	"context"
	"log/slog"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"google.golang.org/grpc"
)

// interceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func interceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

// ServerOptionLogging returns grpc.Server config option that turn on logging.
func ServerOptionLogging(serverOpts []grpc.ServerOption) []grpc.ServerOption {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	// Add unary interceptor
	serverOpts = append(serverOpts, grpc.ChainUnaryInterceptor(
		logging.UnaryServerInterceptor(interceptorLogger(logger), opts...),
	))
	// Add stream interceptor
	serverOpts = append(serverOpts, grpc.ChainStreamInterceptor(
		logging.StreamServerInterceptor(interceptorLogger(logger), opts...),
	))

	return serverOpts
}

// WithDailOptionLogging returns grpc.DialOption config option that turn on logging.
func WithDailOptionLogging(dialOpts []grpc.DialOption) []grpc.DialOption {
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	// Add unary interceptor
	dialOpts = append(dialOpts, grpc.WithChainUnaryInterceptor(
		logging.UnaryClientInterceptor(interceptorLogger(logger), opts...),
	))
	// Add stream interceptor
	dialOpts = append(dialOpts, grpc.WithChainStreamInterceptor(
		logging.StreamClientInterceptor(interceptorLogger(logger), opts...),
	))

	return dialOpts
}
