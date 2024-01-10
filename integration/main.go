// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var specificTest string

func main() {
	rootCmd := &cobra.Command{
		Use:   "integration-tests",
		Short: "Run integration tests",
		Run:   runIntegrationTests,
	}

	rootCmd.PersistentFlags().StringVar(&specificTest, "test", "", "Specific test directory to run")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runIntegrationTests(cmd *cobra.Command, args []string) {
	_, _ = cmd, args

	defer reportResults()

	buildBinary()

	if specificTest != "" {
		fmt.Println("Running", specificTest)
		if !filepath.IsAbs(specificTest) && !strings.HasPrefix(specificTest, "./tests/") {
			specificTest = "./tests/" + specificTest
		}
		logChan = make(chan TestLog, 1)
		runSingleTest(specificTest)
	} else {

		testDirs, err := filepath.Glob("./tests/*")
		if err != nil {
			panic(err)
		}
		logChan = make(chan TestLog, len(testDirs))
		runAllTests()
	}
}
