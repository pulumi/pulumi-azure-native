package versioning

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/stretchr/testify/assert"
)

func TestDefaultVersion(t *testing.T) {
	t.Run("simple latest additive", func(t *testing.T) {
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
		}, nil, DefaultConfig{})
		v2021 := "2021-02-02"
		expected := DefaultConfig{
			"Provider": {
				Tracking: &v2021,
			},
		}
		assert.Equal(t, expected, actual)
	})
	t.Run("two-previews", func(t *testing.T) {
		actual := BuildDefaultConfig(map[openapi.ProviderName]VersionResources{
			"Provider": {
				"2020-01-01-preview": []openapi.ResourceName{
					"Resource A",
				},
				"2021-02-02-preview": []openapi.ResourceName{
					"Resource A",
					"Resource B",
				},
			},
		}, nil, DefaultConfig{})
		v2021 := "2021-02-02-preview"
		expected := DefaultConfig{
			"Provider": {
				Tracking: &v2021,
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("prefer existing tracking version", func(t *testing.T) {
		v2021 := "2021-02-02"
		v2020 := "2020-01-01"

		actual := BuildDefaultConfig(map[openapi.ProviderName]VersionResources{
			"Provider": {
				v2020: []openapi.ResourceName{
					"Resource A",
				},
				v2021: []openapi.ResourceName{
					"Resource A",
					"Resource B",
				},
			},
		}, nil, DefaultConfig{
			"Provider": {
				Tracking: &v2020,
			},
		})
		expected := DefaultConfig{
			"Provider": {
				Tracking: &v2020,
				Additions: &map[openapi.ResourceName]openapi.ApiVersion{
					"Resource B": v2021,
				},
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("prefer preview versions - treat as stable", func(t *testing.T) {
		v2021 := "2021-02-02-preview"
		v2020 := "2020-01-01"

		spec := map[openapi.ProviderName]VersionResources{
			"Provider": {
				v2020: []openapi.ResourceName{
					"Resource A",
				},
				v2021: []openapi.ResourceName{
					"Resource A",
					"Resource B",
				},
			},
		}
		curations := Curations{
			"Provider": {
				Preview: PreviewPrefer,
			},
		}
		existing := DefaultConfig{}
		actual := BuildDefaultConfig(spec, curations, existing)
		expected := DefaultConfig{
			"Provider": {
				Tracking: &v2021,
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("exclude preview versions", func(t *testing.T) {
		v2021 := "2021-02-02-preview"
		v2020 := "2020-01-01"

		spec := map[openapi.ProviderName]VersionResources{
			"Provider": {
				v2020: []openapi.ResourceName{
					"Resource A",
				},
				v2021: []openapi.ResourceName{
					"Resource A",
					"Resource B",
				},
			},
		}
		curations := Curations{
			"Provider": {
				Preview: PreviewExclude,
			},
		}
		existing := DefaultConfig{}
		actual := BuildDefaultConfig(spec, curations, existing)
		expected := DefaultConfig{
			"Provider": {
				Tracking: &v2020,
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("include new services", func(t *testing.T) {
		v2020 := "2020-01-01"

		actual := BuildDefaultConfig(map[openapi.ProviderName]VersionResources{
			"Provider A": {
				v2020: []openapi.ResourceName{
					"Resource A",
				},
			},
			"Provider B": {
				v2020: []openapi.ResourceName{
					"Resource B",
				},
			},
		}, nil, DefaultConfig{
			"Provider A": {
				Tracking: &v2020,
			},
		})
		expected := DefaultConfig{
			"Provider A": {
				Tracking: &v2020,
			},
			"Provider B": {
				Tracking: &v2020,
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("prefer existing addition version", func(t *testing.T) {
		v2020 := "2020-01-01"
		v2021 := "2021-02-02"
		v2022 := "2022-03-03"

		existingConfig := DefaultConfig{
			"Provider": {
				Tracking: &v2021,
				Additions: &map[openapi.ResourceName]openapi.ApiVersion{
					"Resource B": v2020,
				},
			},
		}

		actual := BuildDefaultConfig(map[openapi.ProviderName]VersionResources{
			"Provider": {
				v2020: []openapi.ResourceName{
					"Resource B",
				},
				v2021: []openapi.ResourceName{
					"Resource A",
				},
				v2022: []openapi.ResourceName{
					"Resource B",
					"Resource C",
				},
			},
		}, nil, existingConfig)
		expected := DefaultConfig{
			"Provider": {
				Tracking: &v2021,
				Additions: &map[openapi.ResourceName]openapi.ApiVersion{
					"Resource B": v2020,
					"Resource C": v2022,
				},
			},
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("exclude existing addition", func(t *testing.T) {
		v2020 := "2020-01-01"
		v2021 := "2021-02-02"

		spec := map[openapi.ProviderName]VersionResources{
			"Provider": {
				v2020: []openapi.ResourceName{
					"Resource B",
				},
				v2021: []openapi.ResourceName{
					"Resource A",
				},
			},
		}

		curations := Curations{
			"Provider": {
				Exclusions: map[string]string{
					"Resource B": v2020,
				},
			},
		}

		existingConfig := DefaultConfig{
			"Provider": {
				Tracking: &v2021,
				Additions: &map[openapi.ResourceName]openapi.ApiVersion{
					"Resource B": v2020,
				},
			},
		}

		actual := BuildDefaultConfig(spec, curations, existingConfig)
		expected := DefaultConfig{
			"Provider": {
				Tracking: &v2021,
			},
		}
		assert.Equal(t, expected, actual)
	})
}

func TestExplicitPreference(t *testing.T) {
	v2020 := "2020-01-01"
	v2021 := "2021-02-02"

	spec := map[openapi.ProviderName]VersionResources{
		"Provider": {
			v2020: []openapi.ResourceName{
				"Resource B",
			},
			v2021: []openapi.ResourceName{
				"Resource A",
			},
		},
	}

	curations := Curations{
		"Provider": {
			Explicit: true,
		},
	}

	existingConfig := DefaultConfig{}

	actual := BuildDefaultConfig(spec, curations, existingConfig)
	expected := DefaultConfig{
		"Provider": {
			Additions: &map[openapi.ResourceName]openapi.ApiVersion{
				"Resource A": v2021,
				"Resource B": v2020,
			},
		},
	}
	assert.Equal(t, expected, actual)
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
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{}, "")
		expected := []string{}
		assert.Equal(t, expected, actual)
	})
	t.Run("skips last preview", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01":         {},
			"2020-01-02-preview": {},
		}, "")
		expected := []string{"2020-01-01"}
		assert.Equal(t, expected, actual)
	})
	t.Run("skips preview which is now stable", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01":         {},
			"2020-01-01-preview": {},
			"2022-02-02":         {},
		}, "")
		expected := []string{"2020-01-01", "2022-02-02"}
		assert.Equal(t, expected, actual)
	})
	t.Run("skips multiple previews recently after a stable", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01":         {},
			"2020-01-01-preview": {},
			"2020-06-01-preview": {},
			"2022-02-02":         {},
		}, "")
		expected := []string{"2020-01-01", "2022-02-02"}
		assert.Equal(t, expected, actual)
	})
	t.Run("includes previews a long time after a stable", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01":         {},
			"2020-12-01-preview": {}, // ignored because it's within 1 year of last stable
			"2022-01-01-preview": {},
		}, "")
		expected := []string{"2020-01-01", "2022-01-01-preview"}
		assert.Equal(t, expected, actual)
	})
	t.Run("single preview", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01-preview": {},
		}, "")
		expected := []string{"2020-01-01-preview"}
		assert.Equal(t, expected, actual)
	})
	t.Run("only previews", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2020-01-01-preview": {},
			"2021-01-01-preview": {},
		}, "")
		expected := []string{"2020-01-01-preview", "2021-01-01-preview"}
		assert.Equal(t, expected, actual)
	})
	t.Run("remove private previews", func(t *testing.T) {
		actual := filterCandidateVersions(map[openapi.ApiVersion][]openapi.ResourceName{
			"2015-01-14-preview":        {},
			"2015-01-14-privatepreview": {},
		}, "")
		expected := []string{"2015-01-14-preview"}
		assert.Equal(t, expected, actual)
	})
}
