// Copyright 2016-2020, Pulumi Corporation.

package versioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqueezePreserve(t *testing.T) {
	squeeze := Squeeze{
		"azure-native:provider/version1:resourceA": "azure-native:provider/version2:resourceA",
	}
	squeeze.PreserveResources([]string{
		"azure-native:provider/version1:resourceA",
	})

	assert.Equal(t, squeeze, Squeeze{})
}
