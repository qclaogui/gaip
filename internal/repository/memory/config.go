// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package memory

import (
	"flag"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	defaultPageSize = 20
	maxPageSize     = 10000
	maxBatchSize    = 1000
)

type Config struct {
	FilePath string `yaml:"file_path"`
}

func (cfg *Config) RegisterFlagsWithPrefix(prefix string, fs *flag.FlagSet) {
	fs.StringVar(&cfg.FilePath, prefix+"memory.file-path", "", "Path of JSON file for loads data")
}

func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	cfg.RegisterFlagsWithPrefix("", fs)
}

func (cfg *Config) Validate() error {
	return nil
}

// validatePageSize returns the default page size if the specified page size is 0, otherwise it
// validates the specified page size.
func validatePageSize(ps int32) (int32, error) {
	switch {
	case ps == 0:
		return defaultPageSize, nil
	case ps > maxPageSize:
		return 0, status.Errorf(codes.InvalidArgument, "page size %d cannot be large than max page size %d", ps, maxPageSize)
	case ps < 0:
		return 0, status.Errorf(codes.InvalidArgument, "page size %d cannot be negative", ps)
	}

	return ps, nil
}

// Parses the page token to an int. Returns defaultValue if parsing fails
func parsePageToken(pageToken string, defaultValue int) int {
	if pageToken == "" {
		return defaultValue
	}
	parsed, err := strconv.Atoi(pageToken)
	if err != nil {
		return defaultValue
	}
	return parsed
}

// nextPageToken returns the next page token (the next item index or empty if not more items are left)
func nextPageToken(lastPage, total int) string {
	if lastPage == total {
		return ""
	}
	return strconv.Itoa(lastPage)
}

// Returns the smallest of a and b
func minx(a, b int) int {
	if a < b {
		return a
	}
	return b
}
