// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package main

import (
	"fmt"
	"os"

	"github.com/qclaogui/golang-api-server/pkg/cmd"
)

func main() {
	if err := cmd.Bootstrap(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Oops! %v\n", err)
		os.Exit(1)
	}
}
