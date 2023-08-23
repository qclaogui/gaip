package main

import (
	"fmt"
	"os"

	"github.com/qclaogui/golang-api-server/pkg/cmd"
)

func main() {
	if err := cmd.RunToDoServer(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops! %v\n", err)
		os.Exit(1)
	}
}
