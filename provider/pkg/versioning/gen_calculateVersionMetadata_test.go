package versioning

import (
	"testing"
	"time"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/providerlist"
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
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2019-01-01": {"Resource1": {}, "Resource2": {}},
					"2020-01-01": {"Resource1": {}, "Resource2": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(openapi.ApiVersion("2020-01-01")),
			},
		}), result.Spec)
	})

	t.Run("New module tracks latest of two preview versions", func(t *testing.T) {
		sources := VersionSources{
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2019-01-01-preview": {"Resource1": {}},
					"2020-01-01-preview": {"Resource1": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(openapi.ApiVersion("2020-01-01-preview")),
			},
		}), result.Spec)
	})

	t.Run("New module tracks preview if more than 1 year newer than last stable", func(t *testing.T) {
		sources := VersionSources{
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2019-01-01":         {"Resource1": {}},
					"2020-02-01-preview": {"Resource1": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(openapi.ApiVersion("2020-02-01-preview")),
			},
		}), result.Spec)
	})

	t.Run("New module tracks older stable if not yet 1 year old", func(t *testing.T) {
		today := time.Now().Format("2006-01-02")
		elevenMonthsAgo := time.Now().Add(-time.Hour * 24 * 30 * 11).Format("2006-01-02")
		olderStable := openapi.ApiVersion(elevenMonthsAgo)
		newerPreview := openapi.ApiVersion(today + "-preview")
		sources := VersionSources{
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					olderStable:  {"Resource1": {}},
					newerPreview: {"Resource1": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(olderStable),
			},
		}), result.Spec)
	})

	t.Run("New module prefers live version", func(t *testing.T) {
		sources := VersionSources{
			ProviderList: providerlist.ProviderList{
				Providers: []providerlist.Provider{
					{
						Namespace: "Microsoft.Module",
						ResourceTypes: []providerlist.ResourceType{
							{
								ResourceType: "Resource1",
								ApiVersions:  []string{"2019-01-01"},
							},
							{
								ResourceType: "Resource2",
								ApiVersions:  []string{"2019-01-01"},
							},
						},
					},
				},
			},
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2019-01-01": {"Resource1": {}, "Resource2": {}},
					"2020-01-01": {"Resource1": {}, "Resource2": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(openapi.ApiVersion("2019-01-01")),
			},
		}), result.Spec)
	})

	t.Run("New module with partial latest version includes additions if live", func(t *testing.T) {
		sources := VersionSources{
			ProviderList: providerlist.ProviderList{
				Providers: []providerlist.Provider{
					{
						Namespace: "Microsoft.Module",
						ResourceTypes: []providerlist.ResourceType{
							{
								ResourceType: "Resource1",
								ApiVersions:  []string{"2020-01-01", "2021-01-01"},
							},
							{
								ResourceType: "Resource2",
								ApiVersions:  []string{"2020-01-01"},
							},
						},
					},
				},
			},
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2020-01-01": {"Resource1": {}, "Resource2": {}},
					"2021-01-01": {"Resource1": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				// Tracks latest
				Tracking: ptr(openapi.ApiVersion("2021-01-01")),
				// Adds specific resources missing from latest
				Additions: &map[string]openapi.ApiVersion{
					"Resource2": "2020-01-01",
				},
			},
		}), result.Spec)
	})

	t.Run("New spec from config excluding all future versions of resource", func(t *testing.T) {
		sources := VersionSources{
			Config: map[openapi.ModuleName]moduleCuration{
				"Module": {
					Exclusions: map[openapi.ResourceName]openapi.ApiVersion{
						"Resource2": "*",
					},
				},
			},
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2019-01-01": {"Resource1": {}, "Resource2": {}},
					"2020-01-01": {"Resource1": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(openapi.ApiVersion("2020-01-01")),
			},
		}), result.Spec)
	})

	t.Run("New spec from config excluding specific recent resource version", func(t *testing.T) {
		sources := VersionSources{
			Config: map[openapi.ModuleName]moduleCuration{
				"Module": {
					Exclusions: map[openapi.ResourceName]openapi.ApiVersion{
						"Resource2": "2019-01-01",
					},
				},
			},
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2019-01-01": {"Resource1": {}, "Resource2": {}},
					"2020-01-01": {"Resource1": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(openapi.ApiVersion("2020-01-01")),
			},
		}), result.Spec)
	})

	t.Run("New spec with preference for explicit version selection", func(t *testing.T) {
		sources := VersionSources{
			Config: map[openapi.ModuleName]moduleCuration{
				"Module": {
					Explicit: true,
				},
			},
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2019-01-01": {"Resource1": {}, "Resource2": {}},
					"2020-01-01": {"Resource1": {}},
				},
			},
			ProviderList: providerlist.ProviderList{
				Providers: []providerlist.Provider{
					{
						Namespace: "Microsoft.Module",
						ResourceTypes: []providerlist.ResourceType{
							{
								ResourceType: "Resource1",
								ApiVersions:  []string{"2019-01-01", "2020-01-01"},
							},
							{
								ResourceType: "Resource2",
								ApiVersions:  []string{"2019-01-01"},
							},
						},
					},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Additions: &map[string]openapi.ApiVersion{
					"Resource1": "2020-01-01",
					"Resource2": "2019-01-01",
				},
			},
		}), result.Spec)
	})

	t.Run("New resource ignored if not live", func(t *testing.T) {
		sources := VersionSources{
			ProviderList: providerlist.ProviderList{
				Providers: []providerlist.Provider{
					{
						Namespace: "Microsoft.Module",
						ResourceTypes: []providerlist.ResourceType{
							{
								ResourceType: "Resource1",
								ApiVersions:  []string{"2020-01-01"},
							},
							{
								ResourceType: "Resource2",
								ApiVersions:  []string{"2020-01-01"},
							},
						},
					},
				},
			},
			Spec: map[openapi.ModuleName]ModuleSpec{
				"Module": {
					Tracking: ptr(openapi.ApiVersion("2020-01-01")),
				},
			},
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2020-01-01": {"Resource1": {}},
					"2021-01-01": {"Resource1": {}, "Resource2": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(openapi.ApiVersion("2020-01-01")),
			},
		}), result.Spec)
	})

	t.Run("New version ignored if already tracking during upgrade", func(t *testing.T) {
		sources := VersionSources{
			Spec: map[openapi.ModuleName]ModuleSpec{
				"Module": {
					Tracking: ptr(openapi.ApiVersion("2020-01-01")),
				},
			},
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2020-01-01": {"Resource1": {}, "Resource2": {}},
					"2021-01-01": {"Resource1": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(openapi.ApiVersion("2020-01-01")),
			},
		}), result.Spec)
	})

	t.Run("New resource in new version added during upgrade if live", func(t *testing.T) {
		sources := VersionSources{
			ProviderList: providerlist.ProviderList{
				Providers: []providerlist.Provider{
					{
						Namespace: "Microsoft.Module",
						ResourceTypes: []providerlist.ResourceType{
							{
								ResourceType: "Resource1",
								ApiVersions:  []string{"2020-01-01", "2021-01-01"},
							},
							{
								ResourceType: "Resource2",
								ApiVersions:  []string{"2021-01-01"},
							},
						},
					},
				},
			},
			Spec: map[openapi.ModuleName]ModuleSpec{
				"Module": {
					Tracking: ptr(openapi.ApiVersion("2020-01-01")),
				},
			},
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2020-01-01": {"Resource1": {}},
					"2021-01-01": {"Resource1": {}, "Resource2": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(openapi.ApiVersion("2020-01-01")),
				Additions: &map[string]openapi.ApiVersion{
					"Resource2": "2021-01-01",
				},
			},
		}), result.Spec)
	})

	t.Run("Remove previous addition via exclusion", func(t *testing.T) {
		sources := VersionSources{
			Spec: map[openapi.ModuleName]ModuleSpec{
				"Module": {
					Tracking: ptr(openapi.ApiVersion("2020-01-01")),
					Additions: &map[string]openapi.ApiVersion{
						"Resource2": "2021-01-01",
					},
				},
			},
			Config: map[openapi.ModuleName]moduleCuration{
				"Module": {
					Exclusions: map[openapi.ResourceName]openapi.ApiVersion{
						"Resource2": "2021-01-01",
					},
				},
			},
			AllResourcesByVersion: ModuleVersionResources{
				"Module": {
					"2020-01-01": {"Resource1": {}},
					"2021-01-01": {"Resource1": {}, "Resource2": {}},
				},
			},
		}

		result, err := calculateVersionMetadata(sources)

		assert.NoError(t, err)
		assert.Equal(t, Spec(map[openapi.ModuleName]ModuleSpec{
			"Module": {
				Tracking: ptr(openapi.ApiVersion("2020-01-01")),
			},
		}), result.Spec)
	})
}

func ptr[T any](s T) *T {
	return &s
}
