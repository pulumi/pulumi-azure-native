package versioning

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pulumi/pulumi-azure-native/provider/pkg/openapi"
	"github.com/stretchr/testify/assert"
)

func TestFindOlderVersions(t *testing.T) {
	providerA := gofakeit.Noun()
	resourceA := gofakeit.Noun()
	resourceB := gofakeit.Noun()
	olderVersion := fakeApiVersion(FakeApiVersion{})
	versionA := fakeApiVersion(FakeApiVersion{GreaterThan: olderVersion})
	versionB := fakeApiVersion(FakeApiVersion{GreaterThan: olderVersion})

	specVersions := SpecVersions{
		providerA: {
			olderVersion: []string{resourceA, resourceB},
			versionA:     []string{resourceA, resourceB},
			versionB:     []string{resourceA, resourceB},
		},
	}
	defaultVersion := openapi.DefaultVersionLock{
		providerA: {
			resourceA: versionA,
			resourceB: versionB,
		},
	}
	olderVersions := findOlderVersions(specVersions, defaultVersion)
	expected := openapi.ProviderVersionList{
		providerA: {
			olderVersion,
		},
	}
	assert.Equal(t, expected, olderVersions)
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
		min, _ = time.Parse(ApiVersionLayout, spec.GreaterThan)
		min = min.AddDate(0, 0, 1)
	}
	if spec.LessThan != "" {
		max, _ = time.Parse(ApiVersionLayout, spec.GreaterThan)
		max = max.AddDate(0, 0, -1)
	}
	return gofakeit.DateRange(min, max).Format(ApiVersionLayout)
}
