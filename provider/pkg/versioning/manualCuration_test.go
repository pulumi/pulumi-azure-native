package versioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExcluded(t *testing.T) {
	var curations = make(Curations)

	assert.False(t, curations.IsExcluded("Compute", "someResource"))

	curations["Compute"] = providerCuration{}
	assert.False(t, curations.IsExcluded("Compute", "someResource"))

	curations["Compute"] = providerCuration{Exclusions: []string{"anotherResource"}}
	assert.False(t, curations.IsExcluded("Compute", "someResource"))

	curations["Compute"] = providerCuration{Exclusions: []string{"anotherResource", "someResource"}}
	assert.True(t, curations.IsExcluded("Compute", "someResource"))
}
