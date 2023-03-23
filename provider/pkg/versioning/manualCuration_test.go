package versioning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsExcluded_NoProviders(t *testing.T) {
	var curations = make(Curations)

	isExcluded, err := curations.IsExcluded("Compute", "someResource", "2020-01-01")

	assert.False(t, isExcluded)
	assert.Nil(t, err)
}

func TestIsExcluded_NoProviderExclusions(t *testing.T) {
	var curations = make(Curations)
	curations["Compute"] = providerCuration{}

	isExcluded, err := curations.IsExcluded("Compute", "someResource", "2020-01-01")

	assert.False(t, isExcluded)
	assert.Nil(t, err)
}

func TestIsExcluded_OtherExclusion(t *testing.T) {
	var curations = make(Curations)
	curations["Compute"] = providerCuration{Exclusions: map[string]string{"anotherResource": "*"}}

	isExcluded, err := curations.IsExcluded("Compute", "someResource", "2020-01-01")

	assert.False(t, isExcluded)
	assert.Nil(t, err)
}

func TestIsExcluded_WildcardExclusion(t *testing.T) {
	var curations = make(Curations)
	curations["Compute"] = providerCuration{Exclusions: map[string]string{"someResource": "*"}}

	isExcluded, err := curations.IsExcluded("Compute", "someResource", "2020-01-01")

	assert.True(t, isExcluded)
	assert.Nil(t, err)
}

func TestIsExcluded_ExactExclusion(t *testing.T) {
	var curations = make(Curations)
	curations["Compute"] = providerCuration{Exclusions: map[string]string{"someResource": "2020-01-01"}}

	isExcluded, err := curations.IsExcluded("Compute", "someResource", "2020-01-01")

	assert.True(t, isExcluded)
	assert.Nil(t, err)
}

func TestIsExcluded_OverExclusion(t *testing.T) {
	var curations = make(Curations)
	curations["Compute"] = providerCuration{Exclusions: map[string]string{"someResource": "2022-12-12"}}

	isExcluded, err := curations.IsExcluded("Compute", "someResource", "2020-01-01")

	assert.True(t, isExcluded)
	assert.Nil(t, err)
}

func TestIsExcluded_UnderExclusion(t *testing.T) {
	var curations = make(Curations)
	curations["Compute"] = providerCuration{Exclusions: map[string]string{"someResource": "2000-01-01"}}

	isExcluded, err := curations.IsExcluded("Compute", "someResource", "2020-01-01")

	assert.False(t, isExcluded)
	assert.Error(t, err)
}
