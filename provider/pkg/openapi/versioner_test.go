// Copyright 2016-2020, Pulumi Corporation.

package openapi

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/collections"
	"github.com/stretchr/testify/assert"
)

// Test cases:
//   - Res 1 defined across all API versions
//   - Res 2 defined in two API versions
//   - Res 3 has been renamed but its path is the same, so we consider it the same resource
//     (e.g., happened in Web module for ServerFarm -> AppServicePlan)
//   - Res 4 is named consistently, but the path has changed over time
//     (e.g., happened with several resources in ApiManagement)
var versionMap = map[SdkVersion]VersionResources{
	"v20200101": {
		Resources: map[string]*ResourceSpec{
			"Res1": makeResource("/someprefix/Microsoft.Foo/res1/{res1Name}", "Res 1 v1"),
			"Res2": makeResource("/someprefix/Microsoft.Foo/res2/{res2Name}", "Res 2 v1"),
			"Res3": makeResource("/someprefix/Microsoft.Foo/res3/{res3Name}", "Res 3 v1"),
		},
	},
	"v20200201": {
		Resources: map[string]*ResourceSpec{
			"Res1":        makeResource("/someprefix/Microsoft.Foo/res1/{res1Name}", "Res 1 v2"),
			"Res2":        makeResource("/someprefix/Microsoft.Foo/res2/{res2Name}", "Res 2 v2"),
			"Res3Renamed": makeResource("/someprefix/Microsoft.Foo/res3/{res3Name}", "Res 3 v2"),
			"Res4":        makeResource("/someprefix/Microsoft.Foo/res4/{res4Name}", "Res 4 v1"),
		},
	},
	"v20200301": {
		Resources: map[string]*ResourceSpec{
			"Res1":        makeResource("/someprefix/Microsoft.Foo/res1/{res1Name}", "Res 1 v3"),
			"Res3Renamed": makeResource("/someprefix/Microsoft.Foo/res3/{res3Name}", "Res 3 v3"),
			"Res4":        makeResource("/someprefix/Microsoft.Foo/Res-4/{res4AnotherName}", "Res 4 v2"),
		},
	},
	// The next version is "unknown" yet.
	"v20200401": {
		Resources: map[string]*ResourceSpec{
			"Res1": makeResource("/someprefix/Microsoft.Foo/res1/{res1Name}", "Res 1 v4"),
			"Res4": makeResource("/someprefix/Microsoft.Foo/Res-4/{res4AnotherName}", "Res 4 v3"),
		},
	},
}

func TestFindingPathVersions(t *testing.T) {
	expected := map[string]*collections.OrderableSet[SdkVersion]{
		"/someprefix/microsoft.foo/res1/{}": collections.NewOrderableSet[SdkVersion]("v20200101", "v20200201", "v20200301", "v20200401"),
		"/someprefix/microsoft.foo/res2/{}": collections.NewOrderableSet[SdkVersion]("v20200101", "v20200201"),
		"/someprefix/microsoft.foo/res3/{}": collections.NewOrderableSet[SdkVersion]("v20200101", "v20200201", "v20200301"),
		"/someprefix/microsoft.foo/res4/{}": collections.NewOrderableSet[SdkVersion]("v20200201", "v20200301", "v20200401"),
	}

	actual := calculatePathVersions(versionMap)
	assert.Equal(t, expected, actual)
}

func TestSqueezeSimple(t *testing.T) {
	modules := AzureModules{
		"module": {
			"version1": {
				Resources: map[ResourceName]*ResourceSpec{
					"resourceA": {
						Path: "/someprefix/Microsoft.Foo/res1/{res1Name}",
					},
					"resourceB": {
						Path: "/someprefix/Microsoft.Foo/res2/{res2Name}",
					},
				},
				Invokes: map[InvokeName]*ResourceSpec{
					"invokeA": {
						// Sits under resourceA
						Path: "/someprefix/Microsoft.Foo/res1/{res2Name}/invokeA",
					},
					"invokeB": {
						// Sits under resourceB
						Path: "/someprefix/Microsoft.Foo/res2/{res2Name}/invokeB",
					},
				},
			},
			"version2": {
				Resources: map[ResourceName]*ResourceSpec{
					"resourceA": {
						Path: "/someprefix/Microsoft.Foo/res1/{res1Name}",
					},
					"resourceB": {
						Path: "/someprefix/Microsoft.Foo/res2/{res2Name}",
					},
				},
				Invokes: map[InvokeName]*ResourceSpec{},
			},
		},
	}

	squeeze := RemovableResources{
		"azure-native:module/version1:resourceA": "azure-native:module/version2:resourceA",
	}

	filteredSpec := RemoveResources(modules, squeeze)

	expected := AzureModules{
		"module": {
			"version1": {
				Resources: map[ResourceName]*ResourceSpec{
					"resourceB": {
						Path: "/someprefix/Microsoft.Foo/res2/{res2Name}",
					},
				},
				Invokes: map[InvokeName]*ResourceSpec{
					"invokeB": {
						Path: "/someprefix/Microsoft.Foo/res2/{res2Name}/invokeB",
					},
				},
			},
			"version2": {
				Resources: map[ResourceName]*ResourceSpec{
					"resourceA": {
						Path: "/someprefix/Microsoft.Foo/res1/{res1Name}",
					},
					"resourceB": {
						Path: "/someprefix/Microsoft.Foo/res2/{res2Name}",
					},
				},
				Invokes: map[InvokeName]*ResourceSpec{},
			},
		},
	}
	assert.Equal(t, expected, filteredSpec)
}

func makeResource(path, description string) *ResourceSpec {
	return &ResourceSpec{
		Path: path,
		PathItem: &spec.PathItem{
			PathItemProps: spec.PathItemProps{
				Put: &spec.Operation{
					OperationProps: spec.OperationProps{
						Description: description,
					},
				},
			},
		},
	}
}

func TestSdkToApiVersion(t *testing.T) {
	testConvert := func(t *testing.T, input SdkVersion, expected ApiVersion) {
		t.Helper()
		actual, err := SdkToApiVersion(input)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}
	t.Run("stable", func(t *testing.T) {
		testConvert(t, "v20200101", "2020-01-01")
	})
	t.Run("preview", func(t *testing.T) {
		testConvert(t, "v20200101preview", "2020-01-01-preview")
	})
	t.Run("privatepreview", func(t *testing.T) {
		testConvert(t, "v20200101privatepreview", "2020-01-01-privatepreview")
	})
	t.Run("beta", func(t *testing.T) {
		testConvert(t, "v20200101beta", "2020-01-01-beta")
	})
	t.Run("missing leading v", func(t *testing.T) {
		_, err := SdkToApiVersion("20200101privatepreview")
		assert.Error(t, err)
	})
	t.Run("date too short", func(t *testing.T) {
		_, err := SdkToApiVersion("v2020011")
		assert.Error(t, err)
	})
	t.Run("unknown suffix", func(t *testing.T) {
		_, err := SdkToApiVersion("v20200101foo")
		assert.Error(t, err)
	})
}
