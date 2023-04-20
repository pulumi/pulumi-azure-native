package versioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqueezeSimple(t *testing.T) {
	specVersions := SpecVersions{
		"providerA": {
			"versionA": []string{"resourceA", "resourceB"},
			"versionB": []string{"resourceA", "resourceB"},
		},
	}

	squeeze := Squeeze{
		"providerA/versionA:resourceA": "providerA/versionB:resourceA",
	}

	filteredSpec := SqueezeSpec(specVersions, squeeze)
	expected := SpecVersions{
		"providerA": {
			"versionA": []string{"resourceB"},
			"versionB": []string{"resourceA", "resourceB"},
		},
	}
	assert.Equal(t, expected, filteredSpec)
}
