// Copyright 2016-2020, Pulumi Corporation.

package version

import "github.com/blang/semver"

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

func GetVersion() semver.Version {
	v := Version
	if v == "" {
		// fallback to a default version e.g. for unit tests (see: PROVIDER_VERSION in Makefile)
		v = "2.0.0-alpha.0+dev"
	}
	return semver.MustParse(v)
}
