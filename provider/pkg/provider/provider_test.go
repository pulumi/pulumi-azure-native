// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/stretchr/testify/assert"
)

func TestMapsAutorestCloudToAzureSdkCloud(t *testing.T) {
	k := azureNativeProvider{}

	// default cloud
	assert.Equal(t, k.cloud(), cloud.AzurePublic)

	k.environment = azure.PublicCloud
	assert.Equal(t, k.cloud(), cloud.AzurePublic)

	k.environment = azure.ChinaCloud
	assert.Equal(t, k.cloud(), cloud.AzureChina)

	k.environment = azure.USGovernmentCloud
	assert.Equal(t, k.cloud(), cloud.AzureGovernment)
}
