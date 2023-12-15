// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package log

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	dslog "github.com/grafana/dskit/log"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Logger is a shared go-kit logger.
	Logger         = log.NewNopLogger()
	bufferedLogger *dslog.BufferedLogger
)

type LoggerConfig struct {
	Enabled            bool
	LogsPerSecond      float64
	LogsPerSecondBurst int
	Registry           prometheus.Registerer
}

// InitLogger initialises the global gokit logger and returns that logger.
func InitLogger(logFormat string, logLevel dslog.Level, cfg LoggerConfig) log.Logger {
	var (
		logEntries    uint32 = 256                    // buffer up to 256 log lines in memory before flushing to a write(2) syscall
		logBufferSize uint32 = 10e6                   // 10MB
		flushTimeout         = 100 * time.Millisecond // flush the buffer after 100ms regardless of how full it is, to prevent losing many logs in case of ungraceful termination
	)

	writer := dslog.NewBufferedLogger(os.Stderr, logEntries,
		dslog.WithFlushPeriod(flushTimeout),
		dslog.WithPrellocatedBuffer(logBufferSize),
	)

	logger := dslog.NewGoKitWithWriter(logFormat, writer)

	if cfg.Enabled {
		// use UTC timestamps and skip 6 stack frames if rate limited logger is needed.
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.Caller(6))
		logger = dslog.NewRateLimitedLogger(logger, cfg.LogsPerSecond, cfg.LogsPerSecondBurst, cfg.Registry)
	} else {
		// use UTC timestamps and skip 5 stack frames if no rate limited logger is needed.
		logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.Caller(5))
	}
	// Must put the level filter last for efficiency.
	logger = newFilter(logger, logLevel)

	// Set global logger.
	Logger = logger
	return logger
}

// Pass through Logger and implement the DebugEnabled interface that spanlogger looks for.
type levelFilter struct {
	log.Logger
	debugEnabled bool
}

func newFilter(logger log.Logger, lvl dslog.Level) log.Logger {
	return &levelFilter{
		Logger:       level.NewFilter(logger, lvl.Option),
		debugEnabled: lvl.String() == "debug", // Using inside knowledge about the hierarchy of possible options.
	}
}

func (f *levelFilter) DebugEnabled() bool {
	return f.debugEnabled
}

// CheckFatal prints an error and exits with error code 1 if err is non-nil
func CheckFatal(location string, err error) {
	if err != nil {
		logger := level.Error(Logger)
		if location != "" {
			logger = log.With(logger, "msg", "error "+location)
		}
		// %+v gets the stack trace from errors using github.com/pkg/errors
		_ = logger.Log("err", fmt.Sprintf("%+v", err))

		if err = Flush(); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, "Could not flush logger", err)
		}
		os.Exit(1)
	}
}

// Flush forces the buffered logger, if configured, to flush to the underlying writer
// This is typically only called when the application is shutting down.
func Flush() error {
	if bufferedLogger != nil {
		return bufferedLogger.Flush()
	}

	return nil
}

type DoNotLogError struct{ Err error }

func (i DoNotLogError) Error() string                                     { return i.Err.Error() }
func (i DoNotLogError) Unwrap() error                                     { return i.Err }
func (i DoNotLogError) ShouldLog(_ context.Context, _ time.Duration) bool { return false }
