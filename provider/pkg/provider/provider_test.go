// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/stretchr/testify/assert"
)

func TestMapsAutorestCloudToAzureSdkCloud(t *testing.T) {
	p := azureNativeProvider{}

	// default cloud
	assert.Equal(t, p.cloud(), cloud.AzurePublic)

	p.environment = azure.PublicCloud
	assert.Equal(t, p.cloud(), cloud.AzurePublic)

	p.environment = azure.ChinaCloud
	assert.Equal(t, p.cloud(), cloud.AzureChina)

	p.environment = azure.USGovernmentCloud
	assert.Equal(t, p.cloud(), cloud.AzureGovernment)
}

func TestUseMSALByDefault(t *testing.T) {
	p, usedADAL, usedMSAL := setUpProviderWithMockTokenGetters()

	callGetToken(t, p)
	assert.False(t, *usedADAL)
	assert.True(t, *usedMSAL)
}

func TestUseMSALForManagedIdentity(t *testing.T) {
	p, usedADAL, usedMSAL := setUpProviderWithMockTokenGetters()
	p.config["useMsi"] = "true"

	callGetToken(t, p)
	assert.False(t, *usedADAL)
	assert.True(t, *usedMSAL)
}

func TestUseMSALForServicePrincipalSecret(t *testing.T) {
	p, usedADAL, usedMSAL := setUpProviderWithMockTokenGetters()
	p.config["clientSecret"] = "verysecret"
	p.config["subscriptionId"] = "123"
	p.config["clientId"] = "456"
	p.config["tenantId"] = "789"

	callGetToken(t, p)
	assert.False(t, *usedADAL)
	assert.True(t, *usedMSAL)
}

func TestUseADALForMultitenantAuth(t *testing.T) {
	p, usedADAL, usedMSAL := setUpProviderWithMockTokenGetters()

	p.config["auxiliaryTenantIds"] = "[\"1\"]"
	callGetToken(t, p)
	assert.True(t, *usedADAL)
	assert.False(t, *usedMSAL)
}

func TestUseADALWhenRequested(t *testing.T) {
	p, usedADAL, usedMSAL := setUpProviderWithMockTokenGetters()

	p.config["useLegacyADALAuth"] = "true"
	p.config["auxiliaryTenantIds"] = "[]" // even without multitenant auth
	callGetToken(t, p)
	assert.True(t, *usedADAL)
	assert.False(t, *usedMSAL)
}

func callGetToken(t *testing.T, p *azureNativeProvider) {
	auth, err := p.getAuthConfig()
	assert.NoError(t, err)

	p.getOAuthToken(context.Background(), auth, "")
}

// create a new provider where the func's to obtain auth tokens instead just update the bool return values if they were called
func setUpProviderWithMockTokenGetters() (p *azureNativeProvider, usedADAL *bool, usedMSAL *bool) {
	p = &azureNativeProvider{}
	p.config = map[string]string{}

	var adal bool
	p.adalTokenClient = func(ctx context.Context, k *azureNativeProvider, auth *authentication.Config, endpoint string) (string, error) {
		adal = true
		return "", nil
	}

	var msal bool
	p.msalTokenClient = func(ctx context.Context, k *azureNativeProvider, auth *authentication.Config, endpoint string) (string, error) {
		msal = true
		return "", nil
	}

	usedADAL = &adal
	usedMSAL = &msal
	return
}
