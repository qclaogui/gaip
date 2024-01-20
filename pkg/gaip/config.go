// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"flag"
	"strconv"

	"github.com/pkg/errors"
	"github.com/qclaogui/gaip/internal/repository"
	"github.com/qclaogui/gaip/pkg/service"
	"github.com/qclaogui/gaip/pkg/service/bookstore"
	"github.com/qclaogui/gaip/pkg/service/generativeai"
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

	ServerCfg service.Config    `yaml:"server"`
	RepoCfg   repository.Config `yaml:"database"`

	BookstoreCfg  bookstore.Config    `yaml:"bookstore"`
	GenaiCfg      generativeai.Config `yaml:"genai"`
	LibraryCfg    library.Config      `yaml:"library"`
	ProjectCfg    project.Config      `yaml:"project"`
	RouteGuideCfg routeguide.Config   `yaml:"routeguide"`
	TodoCfg       todo.Config         `yaml:"todo"`

	VaultCfg vault.Config `yaml:"vault"`
}

// RegisterFlags registers flag.
func (c *Config) RegisterFlags(fs *flag.FlagSet) {
	c.ServerCfg.MetricsNamespace = "gaip"

	// Enable native histograms for enabled scrapers with 10% bucket growth.
	c.ServerCfg.MetricsNativeHistogramFactor = 1.1
	c.ServerCfg.ExcludeRequestInLog = true
	c.ServerCfg.DisableRequestSuccessLog = true

	fs.BoolVar(&c.PrintConfig, "print.config", false, "Print the config and exit.")

	c.registerServerFlagsWithChangedDefaultValues(fs)

	// Register Common Repository Config
	c.RepoCfg.RegisterFlags(fs)

	// Register services server Config
	c.BookstoreCfg.RegisterFlags(fs)
	c.GenaiCfg.RegisterFlags(fs)
	c.LibraryCfg.RegisterFlags(fs)
	c.ProjectCfg.RegisterFlags(fs)
	c.RouteGuideCfg.RegisterFlags(fs)
	c.TodoCfg.RegisterFlags(fs)

	// Register Vault Config
	c.VaultCfg.RegisterFlags(fs)
}

// Validate the app config and return an error if the validation doesn't pass
func (c *Config) Validate() error {
	if err := c.BookstoreCfg.Validate(); err != nil {
		return errors.Wrap(err, "invalid Bookstore config")
	}

	if err := c.LibraryCfg.Validate(); err != nil {
		return errors.Wrap(err, "invalid Library config")
	}

	if err := c.ProjectCfg.Validate(); err != nil {
		return errors.Wrap(err, "invalid Project config")
	}

	if err := c.RouteGuideCfg.Validate(); err != nil {
		return errors.Wrap(err, "invalid RouteGuide config")
	}

	if err := c.TodoCfg.Validate(); err != nil {
		return errors.Wrap(err, "invalid Todo config")
	}

	if err := c.VaultCfg.Validate(); err != nil {
		return errors.Wrap(err, "invalid Vault config")
	}

	return nil
}

func (c *Config) registerServerFlagsWithChangedDefaultValues(fs *flag.FlagSet) {
	throwaway := flag.NewFlagSet("throwaway", flag.PanicOnError)
	// Register to throwaway flags first. Default values are remembered during registration and cannot be changed,
	// but we can take values from throwaway flag set and re-register into supplied flag set with new default values.
	c.ServerCfg.RegisterFlags(throwaway)

	defaultsOverrides := map[string]string{
		"server.http-listen-port":                           "7469",
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
