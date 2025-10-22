// Copyright 2016-2020, Pulumi Corporation.

package gen

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
)

func testVersion(input, expected string) func(t *testing.T) {
	return func(t *testing.T) {
		version, err := semver.Parse(input)
		if err != nil {
			assert.Error(t, err)
			return
		}
		actual := GoModVersion(&version)
		assert.Equal(t, expected, actual)
	}
}

func TestGoModVersion(t *testing.T) {
	t.Run("Major", testVersion("1.0.0", "v1.0.0"))
	t.Run("Build", testVersion("1.1.0-alpha.1662577914+9fa804e8", "v1.1.0-alpha.9fa804e8"))
	t.Run("Local Dirty", testVersion("1.1.0-alpha.0+9fa804e8.dirty", "v1.1.0-alpha.9fa804e8.dirty"))
	t.Run("No tags", testVersion("0.0.1-alpha.0+68804cfa", "v0.0.1-alpha.68804cfa"))
	t.Run("Beta", testVersion("2.0.0-beta.3", "v2.0.0-beta.3"))
	t.Run("Build with leading zero", testVersion("3.9.0-alpha.1761105334+0701526", "v3.9.0-alpha.g0701526"))
	t.Run("Build with leading zero and dirty", testVersion("3.9.0-alpha.0+0701526.dirty", "v3.9.0-alpha.g0701526.dirty"))
}
