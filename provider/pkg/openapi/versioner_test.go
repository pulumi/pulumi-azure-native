// Copyright 2016-2020, Pulumi Corporation.

package openapi

import (
	"github.com/go-openapi/spec"
	"github.com/pulumi/pulumi/pkg/v2/codegen"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test cases:
// - Res 1 defined across all API versions
// - Res 2 defined in two API versions
// - Res 3 has been renamed but its path is the same, so we consider it the same resource
//   (e.g., happened in Web provider for ServerFarm -> AppServicePlan)
// - Res 4 is named consistently, but the path has changed over time
//   (e.g., happened with several resources in ApiManagement)
var versionMap = map[string]VersionResources{
	"2020-01-01": {
		Resources: map[string]*ResourceSpec{
			"Res1": makeResource("/someprefix/Microsoft.Foo/res1/{res1Name}", "Res 1 v1"),
			"Res2": makeResource("/someprefix/Microsoft.Foo/res2/{res2Name}", "Res 2 v1"),
			"Res3": makeResource("/someprefix/Microsoft.Foo/res3/{res3Name}", "Res 3 v1"),
		},
	},
	"2020-02-01": {
		Resources: map[string]*ResourceSpec{
			"Res1":        makeResource("/someprefix/Microsoft.Foo/res1/{res1Name}", "Res 1 v2"),
			"Res2":        makeResource("/someprefix/Microsoft.Foo/res2/{res2Name}", "Res 2 v2"),
			"Res3Renamed": makeResource("/someprefix/Microsoft.Foo/res3/{res3Name}", "Res 3 v2"),
			"Res4":        makeResource("/someprefix/Microsoft.Foo/res4/{res4Name}", "Res 4 v1"),
		},
	},
	"2020-03-01": {
		Resources: map[string]*ResourceSpec{
			"Res1":        makeResource("/someprefix/Microsoft.Foo/res1/{res1Name}", "Res 1 v3"),
			"Res3Renamed": makeResource("/someprefix/Microsoft.Foo/res3/{res3Name}", "Res 3 v3"),
			"Res4":        makeResource("/someprefix/Microsoft.Foo/Res-4/{res4AnotherName}", "Res 4 v2"),
		},
	},
	// The next version is "unknown" yet.
	"2020-04-01": {
		Resources: map[string]*ResourceSpec{
			"Res1":        makeResource("/someprefix/Microsoft.Foo/res1/{res1Name}", "Res 1 v4"),
			"Res4":        makeResource("/someprefix/Microsoft.Foo/Res-4/{res4AnotherName}", "Res 4 v3"),
		},
	},
}

func TestFindingKnownLatestResourceVersions(t *testing.T) {
	checker := versioner{lookup: map[string]map[string]codegen.StringSet{
		"test": {
			// Version '2020-04-01' isn't known for any resource.
			"": codegen.NewStringSet("2020-01-01", "2020-02-01", "2020-03-01"),
			// Res1 isn't known in '2020-01-01'.
			"res1":codegen.NewStringSet("2020-02-01", "2020-03-01"),
			// Res2 isn't explicitly listed in the lookup: parent lookup is used.
			// Res3 isn't known in '2020-03-01'.
			"res3": codegen.NewStringSet("2020-01-01", "2020-02-01"),
		},
	}}
	expected := map[string]*ResourceSpec{
		"Res1":        makeResource("/someprefix/Microsoft.Foo/res1/{res1Name}", "Res 1 v3"),
		"Res2":        makeResource("/someprefix/Microsoft.Foo/res2/{res2Name}", "Res 2 v2"),
		"Res3Renamed": makeResource("/someprefix/Microsoft.Foo/res3/{res3Name}", "Res 3 v2"),
		"Res4":        makeResource("/someprefix/Microsoft.Foo/Res-4/{res4AnotherName}", "Res 4 v2"),
	}

	actual := checker.calculateLatestVersions("test", versionMap, false /* invokes */)
	assert.Equal(t, expected, actual)
}

func TestFindingUnknownLatestResourceVersions(t *testing.T) {
	checker := versioner{lookup: map[string]map[string]codegen.StringSet{}}
	expected := map[string]*ResourceSpec{
		"Res1":        makeResource("/someprefix/Microsoft.Foo/res1/{res1Name}", "Res 1 v4"),
		"Res2":        makeResource("/someprefix/Microsoft.Foo/res2/{res2Name}", "Res 2 v2"),
		"Res3Renamed": makeResource("/someprefix/Microsoft.Foo/res3/{res3Name}", "Res 3 v3"),
		"Res4":        makeResource("/someprefix/Microsoft.Foo/Res-4/{res4AnotherName}", "Res 4 v3"),
	}

	actual := checker.calculateLatestVersions("test", versionMap, false /* invokes */)
	assert.Equal(t, expected, actual)
}

func TestFindingPathVersions(t *testing.T) {
	expected := map[string]codegen.StringSet{
		"/someprefix/microsoft.foo/res1/{}": codegen.NewStringSet("2020-01-01", "2020-02-01", "2020-03-01", "2020-04-01"),
		"/someprefix/microsoft.foo/res2/{}": codegen.NewStringSet("2020-01-01", "2020-02-01"),
		"/someprefix/microsoft.foo/res3/{}": codegen.NewStringSet("2020-01-01", "2020-02-01", "2020-03-01"),
		"/someprefix/microsoft.foo/res4/{}": codegen.NewStringSet("2020-02-01", "2020-03-01", "2020-04-01"),
	}

	checker := versioner{}
	actual := checker.calculatePathVersions(versionMap)
	assert.Equal(t, actual, expected)
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
