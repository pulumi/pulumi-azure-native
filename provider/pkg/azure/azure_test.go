// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/stretchr/testify/assert"
)

func TestGetCloudByName(t *testing.T) {
	for _, tc := range []struct {
		name     string
		expected cloud.Configuration
	}{
		{name: "", expected: cloud.AzurePublic},
		{name: "public", expected: cloud.AzurePublic},
		{name: "china", expected: cloud.AzureChina},
		{name: "usgov", expected: cloud.AzureGovernment},
		{name: "usgovernment", expected: cloud.AzureGovernment},
	} {
		assert.Equal(t, tc.expected, GetCloudByName(tc.name), tc.name)
	}
}
