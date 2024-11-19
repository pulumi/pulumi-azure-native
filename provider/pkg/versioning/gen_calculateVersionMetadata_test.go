package versioning

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/stretchr/testify/assert"
)

func TestCalculateVersionMetadata(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		sources := VersionSources{}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("New module tracks latest version", func(t *testing.T) {
		sources := VersionSources{
			AllResourcesByVersion: map[string]map[string][]string{
				"Module": {
					"2019-01-01": {"Resource1", "Resource2"},
					"2020-01-01": {"Resource1", "Resource2"},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ProviderName]ProviderSpec{
			"Module": {
				Tracking: strptr("2020-01-01"),
			},
		}), result.Spec)
	})

	t.Run("New module with partial latest version includes additions", func(t *testing.T) {
		sources := VersionSources{
			AllResourcesByVersion: map[string]map[string][]string{
				"Module": {
					"2020-01-01": {"Resource1", "Resource2"},
					"2021-01-01": {"Resource1"},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ProviderName]ProviderSpec{
			"Module": {
				// Tracks latest
				Tracking: strptr("2021-01-01"),
				// Adds specific resources missing from latest
				Additions: &map[string]string{
					"Resource2": "2020-01-01",
				},
			},
		}), result.Spec)
	})

	t.Run("New spec from config excluding all future versions of resource", func(t *testing.T) {
		sources := VersionSources{
			Config: map[openapi.ProviderName]providerCuration{
				"Module": {
					Exclusions: map[openapi.ResourceName]openapi.ApiVersion{
						"Resource2": "*",
					},
				},
			},
			AllResourcesByVersion: map[string]map[string][]string{
				"Module": {
					"2019-01-01": {"Resource1", "Resource2"},
					"2020-01-01": {"Resource1"},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ProviderName]ProviderSpec{
			"Module": {
				Tracking: strptr("2020-01-01"),
			},
		}), result.Spec)
	})

	t.Run("New spec from config excluding specific recent resource version", func(t *testing.T) {
		sources := VersionSources{
			Config: map[openapi.ProviderName]providerCuration{
				"Module": {
					Exclusions: map[openapi.ResourceName]openapi.ApiVersion{
						"Resource2": "2019-01-01",
					},
				},
			},
			AllResourcesByVersion: map[string]map[string][]string{
				"Module": {
					"2019-01-01": {"Resource1", "Resource2"},
					"2020-01-01": {"Resource1"},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ProviderName]ProviderSpec{
			"Module": {
				Tracking: strptr("2020-01-01"),
			},
		}), result.Spec)
	})

	t.Run("New version ignored if already tracking during upgrade", func(t *testing.T) {
		sources := VersionSources{
			Spec: map[openapi.ProviderName]ProviderSpec{
				"Module": {
					Tracking: strptr("2020-01-01"),
				},
			},
			AllResourcesByVersion: map[string]map[string][]string{
				"Module": {
					"2020-01-01": {"Resource1", "Resource2"},
					"2021-01-01": {"Resource1"},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ProviderName]ProviderSpec{
			"Module": {
				Tracking: strptr("2020-01-01"),
			},
		}), result.Spec)
	})

	t.Run("New resource in new version added during upgrade", func(t *testing.T) {
		sources := VersionSources{
			Spec: map[openapi.ProviderName]ProviderSpec{
				"Module": {
					Tracking: strptr("2020-01-01"),
				},
			},
			AllResourcesByVersion: map[string]map[string][]string{
				"Module": {
					"2020-01-01": {"Resource1"},
					"2021-01-01": {"Resource1", "Resource2"},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ProviderName]ProviderSpec{
			"Module": {
				Tracking: strptr("2020-01-01"),
				Additions: &map[string]string{
					"Resource2": "2021-01-01",
				},
			},
		}), result.Spec)
	})
}

func strptr(s string) *string {
	return &s
}
