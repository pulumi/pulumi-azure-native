package versioning

import (
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultVersion(t *testing.T) {
	t.Run("Single Version", func(t *testing.T) {
		actual := BuildDefaultConfig(map[openapi.ProviderName]VersionResources{
			"Provider": {
				"2020-01-01": []openapi.ResourceName{
					"Resource A",
				},
				"2021-02-02": []openapi.ResourceName{
					"Resource A",
					"Resource B",
				},
			},
		})
		v2021 := "2021-02-02"
		expected := DefaultConfig{
			"Provider": {
				Tracking: &v2021,
			},
		}
		assert.Equal(t, expected, actual)
	})
}

func TestFindMinimalVersionSet(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		actual := findMinimalVersionSet(map[openapi.ApiVersion][]openapi.ResourceName{})
		expected := []openapi.ApiVersion{}
		assert.Equal(t, expected, actual)
	})
	t.Run("latest superset", func(t *testing.T) {
		actual := findMinimalVersionSet(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01": {
				"Resource A",
			},
			"2021-02-02": {
				"Resource A",
				"Resource B",
			},
		})
		expected := []openapi.ApiVersion{
			"2021-02-02",
		}
		assert.Equal(t, expected, actual)
	})
	t.Run("rollup", func(t *testing.T) {
		actual := findMinimalVersionSet(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01": {
				"Resource A",
				"Resource B",
			},
			"2021-02-02": {
				"Resource A",
			},
		})
		expected := []openapi.ApiVersion{
			"2020-01-01", "2021-02-02",
		}
		assert.Equal(t, expected, actual)
	})
}

func TestFilterCandidateVersions(t *testing.T) {
	t.Run("empty spec", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{})
		expected := []string{}
		assert.Equal(t, expected, actual)
	})
	t.Run("skips last preview", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01":         {},
			"2020-01-02-preview": {},
		})
		expected := []string{"2020-01-01"}
		assert.Equal(t, expected, actual)
	})
	t.Run("skips preview which is now stable", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01":         {},
			"2020-01-01-preview": {},
			"2022-02-02":         {},
		})
		expected := []string{"2020-01-01", "2022-02-02"}
		assert.Equal(t, expected, actual)
	})
	t.Run("skips multiple previews recently after a stable", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01":         {},
			"2020-01-01-preview": {},
			"2020-06-01-preview": {},
			"2022-02-02":         {},
		})
		expected := []string{"2020-01-01", "2022-02-02"}
		assert.Equal(t, expected, actual)
	})
	t.Run("includes previews a long time after a stable", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01":         {},
			"2021-01-01-preview": {},
			"2022-01-01-preview": {},
		})
		expected := []string{"2020-01-01", "2022-01-01-preview"}
		assert.Equal(t, expected, actual)
	})
	t.Run("only preview", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01-preview": {},
		})
		expected := []string{"2020-01-01-preview"}
		assert.Equal(t, expected, actual)
	})
}
