// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"flag"
	"strconv"

	"github.com/go-kit/log"
	"github.com/pkg/errors"
	"github.com/qclaogui/gaip/pkg/service"
	"github.com/qclaogui/gaip/pkg/service/bookstore"
	"github.com/qclaogui/gaip/pkg/service/library"
	"github.com/qclaogui/gaip/pkg/service/project"
	"github.com/qclaogui/gaip/pkg/service/routeguide"
	"github.com/qclaogui/gaip/pkg/service/todo"
	"github.com/qclaogui/gaip/pkg/vault"
)

// Config is configuration for Server
type Config struct {
	PrintConfig            bool `yaml:"-"`
	EnableGoRuntimeMetrics bool `yaml:"enable_go_runtime_metrics" category:"advanced"`

	Server service.Config `yaml:"server"`

	Todo       todo.Config       `yaml:"todo"`
	RouteGuide routeguide.Config `yaml:"routeguide"`
	Bookstore  bookstore.Config  `yaml:"bookstore"`
	Library    library.Config    `yaml:"library"`
	Project    project.Config    `yaml:"project"`

	Vault vault.Config `yaml:"vault"`
}

// RegisterFlags registers flag.
func (c *Config) RegisterFlags(fs *flag.FlagSet, _ log.Logger) {
	c.Server.MetricsNamespace = "gaip"

	// Enable native histograms for enabled scrapers with 10% bucket growth.
	c.Server.MetricsNativeHistogramFactor = 1.1
	c.Server.ExcludeRequestInLog = true
	c.Server.DisableRequestSuccessLog = true

	fs.BoolVar(&c.PrintConfig, "print.config", false, "Print the config and exit.")

	// Register projectServerImpl Config
	// Register service server Config
	c.registerServerFlagsWithChangedDefaultValues(fs)

	// Register bookstore Config
	c.Todo.RegisterFlags(fs)
	c.RouteGuide.RegisterFlags(fs)
	c.Bookstore.RegisterFlags(fs)
	c.Library.RegisterFlags(fs)
	c.Project.RegisterFlags(fs)

	// Register Vault Config
	c.Vault.RegisterFlags(fs)
}

// Validate the app config and return an error if the validation doesn't pass
func (c *Config) Validate(_ log.Logger) error {

	if err := c.Todo.Validate(); err != nil {
		return errors.Wrap(err, "invalid Todo config")
	}

	if err := c.RouteGuide.Validate(); err != nil {
		return errors.Wrap(err, "invalid RouteGuide config")
	}

	if err := c.Bookstore.Validate(); err != nil {
		return errors.Wrap(err, "invalid Bookstore config")
	}

	if err := c.Library.Validate(); err != nil {
		return errors.Wrap(err, "invalid Library config")
	}

	if err := c.Project.Validate(); err != nil {
		return errors.Wrap(err, "invalid Project config")
	}

	if err := c.Vault.Validate(); err != nil {
		return errors.Wrap(err, "invalid Vault config")
	}

	return nil
}

func (c *Config) registerServerFlagsWithChangedDefaultValues(fs *flag.FlagSet) {
	throwaway := flag.NewFlagSet("throwaway", flag.PanicOnError)
	// Register to throwaway flags first. Default values are remembered during registration and cannot be changed,
	// but we can take values from throwaway flag set and re-register into supplied flag set with new default values.
	c.Server.RegisterFlags(throwaway)

	defaultsOverrides := map[string]string{
		"server.http-listen-port":                           "8080",
		"server.http-write-timeout":                         "2m",
		"server.grpc.keepalive.min-time-between-pings":      "10s",
		"server.grpc.keepalive.ping-without-stream-allowed": "true",
		"server.grpc-max-recv-msg-size-bytes":               strconv.Itoa(100 * 1024 * 1024),
		"server.grpc-max-send-msg-size-bytes":               strconv.Itoa(100 * 1024 * 1024),
	}

	throwaway.VisitAll(func(f *flag.Flag) {
		if defaultValue, ok := defaultsOverrides[f.Name]; ok {
			_ = f.Value.Set(defaultValue)
		}
		fs.Var(f.Value, f.Name, f.Usage)
	})
}
