// Copyright 2016-2024, Pulumi Corporation.

package azure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractingNetworkPrivateEndpointName(t *testing.T) {
	t.Parallel()

	exampleID := "/subscriptions/subNAME/resourceGroups/myResourceGroup/providers/Microsoft.Network/privateEndpoints/myPrivateEndpoint"
	expectedName := "myPrivateEndpoint"

	assert.True(t, needsSerialization(exampleID), "Network Private Endpoint resources should require serialization.")
	actualName := extractPrivateEndpointNameFromID(exampleID)
	assert.Equal(t, expectedName, actualName, "The extracted private endpoint name should match the expected name.")
}
