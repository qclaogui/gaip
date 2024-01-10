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
	"sync"
)

const (
	binaryPath = "../../../bin/gaip_darwin_amd64"
)

type TestLog struct {
	TestDir    string
	TestLog    string
	TestOutput string
}

var logChan chan TestLog

func executeCommand(command string, args []string, taskDescription string) {
	fmt.Printf("%s...\n", taskDescription)
	cmd := exec.Command(command, args...)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("Error: %s\n", stderr.String())
	}
}

func buildBinary() {
	executeCommand("make", []string{"-C", "..", "build"}, "Building")
}

func runSingleTest(testDir string) {
	info, err := os.Stat(testDir)
	if err != nil {
		panic(err)
	}
	if !info.IsDir() {
		return
	}

	dirName := filepath.Base(testDir)
	var logBuffer bytes.Buffer
	cmd := exec.Command(binaryPath, "-log.level=debug")
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
		panic(err)
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

	var wg sync.WaitGroup
	for _, testDir := range testDirs {
		fmt.Println("Running", testDir)
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			runSingleTest(dir)
		}(testDir)
	}
	wg.Wait()
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
