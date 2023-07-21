//go:build release

/*
Copyright © 2023 Weifeng Wang <qclaogui@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/blang/semver"
	"github.com/dave/jennifer/jen"

	"github.com/qclaogui/golang-api-server/pkg/version"
)

const (
	versionFilename         = "pkg/version/release.go"
	defaultPreReleaseID     = "dev"
	defaultReleaseCandidate = "rc.0"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("missing argument")
	}

	command := os.Args[1]

	switch command {
	case "development":
		newVersion, newPreRelease := nextDevelopmentIteration()
		if err := writeVersionToFile(newVersion, newPreRelease, versionFilename); err != nil {
			log.Fatalf("unable to write file: %v", err)
		}

		version.Version = newVersion
		version.PreReleaseID = newPreRelease
		fmt.Println(version.GetVersion())
	case "next-pre-release-id":
		switch len(os.Args) {
		case 2:
			fmt.Println(defaultReleaseCandidate)
		case 3:
			next, err := nextPreReleaseID(os.Args[2])
			if err != nil {
				log.Fatalf("error generating next pre-release ID: %v", err)
			}
			fmt.Println(next)
		default:
			log.Fatalf("usage: release_generate %s <latest-rc-tag>", command)
		}

	case "full-version":
		fmt.Println(version.GetVersion())
	case "print-version":
		// Print simplified version X.Y.Z
		fmt.Println(version.Version)
	case "print-major-minor-version":
		fmt.Println(printMajorMinor())
	default:
		log.Fatalf("unknown option %q. Expected one of %v", command, strings.Join([]string{"development", "next-pre-release-id", "full-version", "print-version", "print-major-minor-version"}, ", "))
	}

}

func nextPreReleaseID(latestPreReleaseVersion string) (string, error) {
	if latestPreReleaseVersion == "" {
		return defaultReleaseCandidate, nil
	}

	latestPreReleaseVersion = strings.TrimPrefix(latestPreReleaseVersion, "v")
	ver, err := semver.Parse(latestPreReleaseVersion)
	if err != nil {
		return "", fmt.Errorf("invalid pre-release version: %w", err)
	}
	currentVersion, err := semver.Parse(version.Version)
	if err != nil {
		return "", fmt.Errorf("unexpected error parsing current version: %s: %w", version.Version, err)
	}

	verWithoutPre := ver
	verWithoutPre.Pre = nil
	if verWithoutPre.LT(currentVersion) || len(ver.Pre) == 0 {
		return defaultReleaseCandidate, nil
	}

	if len(ver.Pre) != 2 {
		return "", errors.New("unexpected format for PR version")
	}
	id := ver.Pre[1]
	if !id.IsNumeric() {
		return "", fmt.Errorf("expected PR version to be numeric; got %q", id.String())
	}

	return fmt.Sprintf("rc.%d", id.VersionNum+1), nil

}

func printMajorMinor() string {
	ver := semver.MustParse(version.Version)
	return fmt.Sprintf("%v.%v", ver.Major, ver.Minor)
}

func nextDevelopmentIteration() (string, string) {
	ver := semver.MustParse(version.Version)
	ver.Minor++
	return ver.String(), defaultPreReleaseID
}

func writeVersionToFile(version, preReleaseID, fileName string) error {
	f := jen.NewFilePath("pkg/version")
	f.PackageComment(`/*
	Copyright © 2023 Weifeng Wang <qclaogui@gmail.com>
	
	Permission is hereby granted, free of charge, to any person obtaining a copy
	of this software and associated documentation files (the "Software"), to deal
	in the Software without restriction, including without limitation the rights
	to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
	copies of the Software, and to permit persons to whom the Software is
	furnished to do so, subject to the following conditions:
	
	The above copyright notice and this permission notice shall be included in
	all copies or substantial portions of the Software.
	
	THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
	IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
	AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
	LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
	OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
	THE SOFTWARE.
	*/`)

	f.Comment("This file was generated by release_generate.go; DO NOT EDIT.")
	f.Line()

	f.Comment("Version is the version number in semver format X.Y.Z")
	f.Var().Id("Version").Op("=").Lit(version)

	f.Comment("PreReleaseID can be empty for releases, \"rc.X\" for release candidates and \"dev\" for snapshots")
	f.Var().Id("PreReleaseID").Op("=").Lit(preReleaseID)

	f.Comment("GitCommit is the short commit hash. It will be set by the linker.")
	f.Var().Id("GitCommit").Op("=").Lit("")

	f.Comment("BuildDate is the time of the build with format yyyy-mm-ddThh:mm:ssZ. It will be set by the linker.")
	f.Var().Id("BuildDate").Op("=").Lit("")

	f.Comment("GoVersion returns the Go version string.")
	f.Var().Id("GoVersion").Op("=").Lit(runtime.Version())

	return f.Save(fileName)
}
