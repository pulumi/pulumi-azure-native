// Copyright 2016-2020, Pulumi Corporation.

package version

import "github.com/blang/semver"

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

func GetVersion() semver.Version {
	return semver.MustParse(Version)
}
