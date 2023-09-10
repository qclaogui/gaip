// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package app

import (
	"flag"
	"strconv"

	"github.com/go-kit/log"
	"github.com/grafana/dskit/server"
	"github.com/pkg/errors"
	bookstorev1alpha1 "github.com/qclaogui/golang-api-server/pkg/service/bookstore/v1alpha1"
	"github.com/qclaogui/golang-api-server/pkg/vault"
)

// Config is configuration for Server
type Config struct {
	PrintConfig bool `yaml:"-"`

	Server server.Config `yaml:"server"`

	Bookstore bookstorev1alpha1.Config `yaml:"bookstore"`

	Vault vault.Config `yaml:"vault"`

	//// DB Datastore parameters section
	//// Host is host of database
	//Host string
	//// User is username to connect to database
	//User string
	//// Password password to connect to database
	//Password string
	//// Schema is schema of database
	//Schema string
}

// RegisterFlags registers flag.
func (c *Config) RegisterFlags(fs *flag.FlagSet, _ log.Logger) {
	c.Server.MetricsNamespace = "qclaogui"

	// Enable native histograms for enabled scrapers with 10% bucket growth.
	c.Server.MetricsNativeHistogramFactor = 1.1
	c.Server.ExcludeRequestInLog = true
	c.Server.DisableRequestSuccessLog = true

	fs.BoolVar(&c.PrintConfig, "print.config", false, "Print the config and exit.")

	// Register Server Config
	c.registerServerFlagsWithChangedDefaultValues(fs)

	// Register Bookstore Config
	c.Bookstore.RegisterFlags(fs)

	// Register Vault Config
	c.Vault.RegisterFlags(fs)
}

// Validate the app config and return an error if the validation doesn't pass
func (c *Config) Validate(_ log.Logger) error {

	// Validate Bookstore Config
	if err := c.Bookstore.Validate(); err != nil {
		return errors.Wrap(err, "invalid Bookstore config")
	}

	// Validate Vault Config
	if err := c.Vault.Validate(); err != nil {
		return errors.Wrap(err, "invalid vault config")
	}
	return nil
}

func (c *Config) registerServerFlagsWithChangedDefaultValues(fs *flag.FlagSet) {
	throwaway := flag.NewFlagSet("throwaway", flag.PanicOnError)
	// Register to throwaway flags first. Default values are remembered during registration and cannot be changed,
	// but we can take values from throwaway flag set and re-register into supplied flag set with new default values.
	c.Server.RegisterFlags(throwaway)

	defaultsOverrides := map[string]string{
		"server.http-write-timeout":                         "2m",
		"server.grpc.keepalive.min-time-between-pings":      "10s",
		"server.grpc.keepalive.ping-without-stream-allowed": "true",
		"server.http-listen-port":                           "8080",
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
