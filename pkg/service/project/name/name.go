// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

// Package name deals with parsing and formatting resource names used in the Library API
package name

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ParseProject parses the project ID from a project resource name.
func ParseProject(name string) (string, error) {
	params := strings.Split(name, "/")
	if len(params) != 2 {
		return "", status.Errorf(codes.InvalidArgument, `project name must be in the form "projects/[PROJECT_ID]", got %q`, name)
	}
	if params[0] != "projects" {
		return "", status.Errorf(codes.InvalidArgument, `project name must be in the form "projects/[PROJECT_ID]", got %q`, name)
	}
	if params[1] == "" {
		return "", status.Errorf(codes.InvalidArgument, `project name must be in the form "projects/[PROJECT_ID]", got %q`, name)
	}

	return params[1], nil
}
