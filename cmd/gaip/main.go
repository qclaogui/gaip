// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"bytes"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/go-kit/log/level"
	"github.com/grafana/dskit/flagext"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/qclaogui/gaip/pkg/gaip"
	"github.com/qclaogui/gaip/pkg/version"
	lg "github.com/qclaogui/gaip/tools/log"
	"github.com/qclaogui/gaip/tools/usage"
	"gopkg.in/yaml.v3"
)

// configHash exposes information about the loaded config
var configHash = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "gaip_config_hash",
		Help: "Hash of the currently active config file.",
	},
	[]string{"sha256"},
)

const configFileOption = "config.file"

func init() {
	prometheus.MustRegister(configHash)
}

type mainFlags struct {
	rateLimitedLogsEnabled        bool
	rateLimitedLogsPerSecond      float64
	rateLimitedLogsPerSecondBurst int
	printVersion                  bool
	printHelp                     bool
	printHelpAll                  bool
	dumpYaml                      bool
}

func (mf *mainFlags) registerFlags(f *flag.FlagSet) {
	f.BoolVar(&mf.rateLimitedLogsEnabled, "log.rate-limit-enabled", false, "Use rate limited logger to reduce the number of logged messages per second.")
	f.Float64Var(&mf.rateLimitedLogsPerSecond, "log.rate-limit-logs-per-second", 10000, "Maximum number of messages per second to be logged.")
	f.IntVar(&mf.rateLimitedLogsPerSecondBurst, "log.rate-limit-logs-per-second-burst", 25000, "Burst size, i.e., maximum number of messages that can be logged in a second, temporarily exceeding the configured maximum logs per second.")
	f.BoolVar(&mf.printVersion, "version", false, "Print application version and exit.")
	f.BoolVar(&mf.printHelp, "help", false, "Print basic help.")
	f.BoolVar(&mf.printHelp, "h", false, "Print basic help.")
	f.BoolVar(&mf.printHelpAll, "help-all", false, "Print help, also including advanced and experimental parameters.")
	f.BoolVar(&mf.dumpYaml, "dump-yaml", false, "Print full config yaml.")
}

func main() {
	// Cleanup all flags registered via init() methods of 3rd-party libraries.
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	// This sets default values from flags to the config.
	// It needs to be called before parsing the config file!
	var cfg gaip.Config
	cfg.RegisterFlags(flag.CommandLine, lg.Logger)

	if configFile := parseConfigFileParameter(os.Args[1:]); configFile != "" {
		if err := LoadConfig(configFile, &cfg); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error loading config from %s: %v\n", configFile, err)
			exit(1)
		}
	}

	// Ignore -config.file here, since it is parsed separately, but still present on the command line.
	flagext.IgnoredFlag(flag.CommandLine, configFileOption, "Configuration file to load.")

	var mf mainFlags
	mf.registerFlags(flag.CommandLine)

	flag.CommandLine.Usage = func() { /* don't do anything by default, we will print usage ourselves. */ }
	flag.CommandLine.Init(flag.CommandLine.Name(), flag.ContinueOnError)

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		_, _ = fmt.Fprintln(flag.CommandLine.Output(), "Run with -help to get a list of available parameters")
	}

	if mf.printHelp || mf.printHelpAll {
		// Print available parameters to stdout, so that users can grep/less them easily.
		flag.CommandLine.SetOutput(os.Stdout)
		if err := usage.Usage(mf.printHelpAll, &mf, &cfg); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "error printing usage: %s\n", err)
			exit(1)
		}
		return
	}

	if mf.printVersion {
		_, _ = fmt.Fprintln(os.Stdout, version.String())
		return
	}

	// Validate the config once both the config file has been loaded
	// and CLI flags parsed.
	if err := cfg.Validate(lg.Logger); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error validating config: %v\n", err)
	}

	if mf.dumpYaml {
		DumpYaml(&cfg)
		return
	}

	reg := prometheus.DefaultRegisterer
	cfg.Server.Log = lg.InitLogger(cfg.Server.LogFormat, cfg.Server.LogLevel, lg.LoggerConfig{
		Enabled:            mf.rateLimitedLogsEnabled,
		LogsPerSecond:      mf.rateLimitedLogsPerSecond,
		LogsPerSecondBurst: mf.rateLimitedLogsPerSecondBurst,
		Registry:           reg,
	})
	_ = level.Info(lg.Logger).Log("msg", "Starting application", "version", version.GetVersion())

	g, err := gaip.New(cfg, reg)
	lg.CheckFatal("initializing application", err)

	err = g.Bootstrap()
	lg.CheckFatal("running application", err)
}

func exit(code int) {
	if err := lg.Flush(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "Could not flush logger", err)
	}
	os.Exit(code)
}

// Parse -config.file option
func parseConfigFileParameter(args []string) (configFile string) {
	// ignore errors and any output here. Any flag errors will be reported by main flag.Parse() call.
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	// usage not used in these functions.
	fs.StringVar(&configFile, configFileOption, "", "")

	// Try to find -config.file option in the flags. As Parsing stops on the first error, eg. unknown flag, we simply
	// try remaining parameters until we find config flag, or there are no params left.
	// (ContinueOnError just means that flag.Parse doesn't call panic or os.Exit, but it returns error, which we ignore)
	for len(args) > 0 {
		_ = fs.Parse(args)
		args = args[1:]
	}
	return
}

// LoadConfig read YAML-formatted config from filename into cfg.
func LoadConfig(filename string, cfg *gaip.Config) error {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "Error reading config file")
	}
	// create a sha256 hash of the config before expansion and expose it via
	// the config_info metric
	hash := sha256.Sum256(buf)
	configHash.Reset()
	configHash.WithLabelValues(fmt.Sprintf("%x", hash)).Set(1)

	dec := yaml.NewDecoder(bytes.NewReader(expandEnvironmentVariables(buf)))
	dec.KnownFields(true)

	// Unmarshal with common config unmarshaler.
	if err = dec.Decode(cfg); err != nil {
		return errors.Wrap(err, "Error parsing config file")
	}
	return nil
}

// expandEnvironmentVariables replaces ${var} or $var in config according to the values of the current environment variables.
// The replacement is case-sensitive. References to undefined variables are replaced by the empty string.
// A default value can be given by using the form ${var:default value}.
func expandEnvironmentVariables(config []byte) []byte {
	return []byte(os.Expand(string(config), func(key string) string {
		keyAndDefault := strings.SplitN(key, ":", 2)
		key = keyAndDefault[0]
		v := os.Getenv(key)

		if v == "" && len(keyAndDefault) == 2 {
			// Set value to the default.
			v = keyAndDefault[1]
		}
		return v
	}))
}

func DumpYaml(cfg *gaip.Config) {
	out, err := yaml.Marshal(cfg)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf("%s\n", out)
	}
}
