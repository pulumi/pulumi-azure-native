// Copyright 2022, Pulumi Corporation.

package provider

import (
	"testing"

	goversion "github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
)

func TestAcceptedAzVersions(t *testing.T) {
	goodVersions := []string{
		"2.37", "2.37.1", "2.38", "2.40", "2.99.99",
	}
	badVersions := []string{
		"1",
		"2", "2.0.80", "2.0.81", "2.1.81", "2.20", "2.33", "2.36.6",
		"3", "3.0.1", "3.5",
		"4",
	}

	for _, good := range goodVersions {
		v := goversion.Must(goversion.NewVersion(good))
		assert.Nil(t, assertAzVersion(v))
	}

	for _, bad := range badVersions {
		v := goversion.Must(goversion.NewVersion(bad))
		assert.NotNil(t, assertAzVersion(v))
	}
}
