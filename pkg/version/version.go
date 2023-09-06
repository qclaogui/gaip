// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package version

import (
	"encoding/json"
	"fmt"
	"strings"
)

//go:generate go run ./release_generate.go

// Info VersionInfo contains build information.
type Info struct {
	Version      string
	PreReleaseID string
	GitCommit    string
	BuildDate    string
	GoVersion    string
}

// GetVersionInfo returns version Info struct
func GetVersionInfo() Info {
	return Info{
		Version:      Version,
		PreReleaseID: PreReleaseID,
		GitCommit:    GitCommit,
		BuildDate:    BuildDate,
		GoVersion:    GoVersion,
	}
}

// ExtraSep separates semver version from any extra version info
const ExtraSep = "-"

// String return version info as JSON
func String() string {
	if data, err := json.Marshal(GetVersionInfo()); err == nil {
		return string(data)
	}
	return ""
}

// GetVersion return the exact version of this build
func GetVersion() string {
	if PreReleaseID == "" {
		return Version
	}

	versionWithPR := fmt.Sprintf("%s%s%s", Version, ExtraSep, PreReleaseID)

	if isReleaseCandidate(PreReleaseID) || (GitCommit == "" || BuildDate == "") {
		return versionWithPR
	}

	//  Include build metadata
	return fmt.Sprintf("%s+%s.%s",
		versionWithPR,
		GitCommit,
		BuildDate,
	)
}

func isReleaseCandidate(preReleaseID string) bool {
	return strings.HasPrefix(preReleaseID, "rc.")
}
