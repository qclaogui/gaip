// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

var binaryPath = fmt.Sprintf("../../../bin/gaip_%s_%s", runtime.GOOS, runtime.GOARCH)

var logChan chan TestLog

type TestLog struct {
	TestDir    string
	TestLog    string
	TestOutput string
}

var specificTest string

func main() {
	rootCmd := &cobra.Command{
		Use:   "integration-tests",
		Short: "Run integration tests",
		Run:   runIntegrationTests,
	}

	rootCmd.PersistentFlags().StringVarP(&specificTest, "test", "t", "", "Specific test directory to run")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runIntegrationTests(*cobra.Command, []string) {
	defer reportResults()

	// build Binary
	executeCommand("make", []string{"-C", "..", "build"}, "Building")

	if specificTest != "" {
		fmt.Println("Running", specificTest)
		if !filepath.IsAbs(specificTest) && !strings.HasPrefix(specificTest, "./tests/") {
			specificTest = "./tests/" + specificTest
		}
		logChan = make(chan TestLog, 1)
		runSingleTest(specificTest)
		close(logChan)
	} else {
		testDirs, err := filepath.Glob("./tests/*")
		if err != nil {
			panic(err)
		}
		logChan = make(chan TestLog, len(testDirs))
		runAllTests()
	}
}

func executeCommand(command string, args []string, taskDescription string) {
	fmt.Printf("%s...\n", taskDescription)
	cmd := exec.Command(command, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error: %s\n", stderr.String())
	}
}

func runSingleTest(testDir string) {
	info, err := os.Stat(testDir)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !info.IsDir() {
		return
	}

	dirName := filepath.Base(testDir)
	var logBuffer bytes.Buffer
	cmd := exec.Command(binaryPath, "-log.level=info")
	cmd.Dir = testDir
	cmd.Stdout = &logBuffer
	cmd.Stderr = &logBuffer

	if err = cmd.Start(); err != nil {
		logChan <- TestLog{
			TestDir: dirName,
			TestLog: fmt.Sprintf("Failed to start gaip: %v", err),
		}
		return
	}

	testCmd := exec.Command("go", "test", "-tags=requires_docker")
	testCmd.Dir = testDir
	testOutput, errTest := testCmd.CombinedOutput()

	if err = cmd.Process.Kill(); err != nil {
		fmt.Println(err)
		return
	}

	if errTest != nil {
		logChan <- TestLog{
			TestDir:    dirName,
			TestLog:    logBuffer.String(),
			TestOutput: string(testOutput),
		}
	}
}

func runAllTests() {
	testDirs, err := filepath.Glob("./tests/*")
	if err != nil {
		panic(err)
	}

	for _, testDir := range testDirs {
		fmt.Println("Running", testDir)
		runSingleTest(testDir)
	}

	close(logChan)
}

func reportResults() {
	testsFailed := 0
	for lc := range logChan {
		fmt.Printf("Failure detected in %s:\n", lc.TestDir)
		fmt.Println("Test output:", lc.TestOutput)
		fmt.Println("Test logs:", lc.TestLog)
		testsFailed++
	}

	if testsFailed > 0 {
		fmt.Printf("%d tests failed!\n", testsFailed)
	} else {
		fmt.Println("All integration tests passed!")
	}
}
