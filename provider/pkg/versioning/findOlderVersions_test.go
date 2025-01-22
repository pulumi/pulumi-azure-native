package versioning

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/stretchr/testify/assert"
)

func TestFindOlderVersions(t *testing.T) {
	moduleA := openapi.ModuleName("moduleA")
	resourceA := "resourceA"
	resourceB := "resourceB"
	olderVersion := fakeApiVersion(FakeApiVersion{})
	versionA := fakeApiVersion(FakeApiVersion{GreaterThan: olderVersion})
	versionB := fakeApiVersion(FakeApiVersion{GreaterThan: olderVersion})

	specVersions := ModuleVersionResources{
		moduleA: {
			olderVersion: []string{resourceA, resourceB},
			versionA:     []string{resourceA, resourceB},
			versionB:     []string{resourceA, resourceB},
		},
	}
	defaultVersions := openapi.DefaultVersions{
		moduleA: {
			resourceA: openapi.DefinitionVersion{ApiVersion: versionA},
			resourceB: openapi.DefinitionVersion{ApiVersion: versionB},
		},
	}
	olderVersions := findOlderVersions(specVersions, defaultVersions)
	expected := openapi.ModuleVersionList{
		moduleA: {
			olderVersion,
		},
	}
	assert.Equal(t, expected, olderVersions, "olderVersion=%s, versionA=%s, versionB=%s", olderVersion, versionA, versionB)
}

type FakeApiVersion struct {
	LessThan    openapi.ApiVersion
	GreaterThan openapi.ApiVersion
}

const ApiVersionLayout = "2006-01-02"

func fakeApiVersion(spec FakeApiVersion) openapi.ApiVersion {
	min, _ := time.Parse(ApiVersionLayout, "2000-01-01")
	max, _ := time.Parse(ApiVersionLayout, "2100-01-01")
	if spec.GreaterThan != "" {
		min, _ = time.Parse(ApiVersionLayout, string(spec.GreaterThan))
		min = min.AddDate(0, 0, 1)
	}
	if spec.LessThan != "" {
		max, _ = time.Parse(ApiVersionLayout, string(spec.LessThan))
		max = max.AddDate(0, 0, -1)
	}
	return openapi.ApiVersion(gofakeit.DateRange(min, max).Format(ApiVersionLayout))
}
