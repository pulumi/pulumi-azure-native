package versionLookup_test

import (
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/versionLookup"
	"github.com/stretchr/testify/assert"
)

func TestLookup(t *testing.T) {
	result, ok := versionLookup.GetDefaultApiVersionForResource("Web", "WebApp")
	assert.True(t, ok, "expected to find default API version for Web/WebApp")
	// Verify we can convert it back to a date to prove it's in a valid format
	_, err := openapi.ApiVersionToDate(openapi.ApiVersion(result))
	assert.NoError(t, err, "expected to parse API version as date")
}
