// Copyright 2016-2025, Pulumi Corporation.

package provider

import (
	"context"
	"fmt"
	"strings"

	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	azpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	auth "github.com/microsoftgraph/msgraph-sdk-go-core/authentication"
	"github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// ClientConfig represents the provider's Azure client configuration, including the Azure environment,
// client identity, and target subscription.
type ClientConfig struct {
	Cloud azcloud.Configuration

	ClientID       string
	ObjectID       string
	SubscriptionID string
	TenantID       string

	AuthenticatedAsAServicePrincipal bool
}

// GetClientConfig resolves the provider's identity given the auth configuration and a TokenCredential.
// It returns a ClientConfig which contains the client ID, object ID, subscription ID, tenant.
func GetClientConfig(ctx context.Context, config *authConfiguration, cred *azAccount) (*ClientConfig, error) {

	// The original specification for what constitutes the "client configuration" is from here:
	// https://github.com/hashicorp/terraform-provider-azurerm/blob/572bb4f37d73f4f0d914737eaca4e5a1fd084c86/internal/clients/auth.go#L33

	endpoint := getGraphEndpoint(cred.Cloud)
	logging.V(9).Infof("MSGraph endpoint: %s", endpoint)

	// Acquire an access token so we can inspect the claims
	scope := fmt.Sprintf("https://%s/.default", endpoint)
	token, err := cred.GetToken(ctx, azpolicy.TokenRequestOptions{
		Scopes: []string{scope},
	})
	if err != nil {
		return nil, fmt.Errorf("could not acquire access token to parse claims: %+v", err)
	}

	tokenClaims, err := azure.ParseClaims(token)
	if err != nil {
		return nil, fmt.Errorf("parsing claims from access token: %+v", err)
	}

	authenticatedAsServicePrincipal := true
	if strings.Contains(strings.ToLower(tokenClaims.Scopes), "openid") {
		authenticatedAsServicePrincipal = false
	}

	clientId := tokenClaims.AppId
	if clientId == "" {
		logging.V(9).Infof("Using user-supplied ClientID because the `appid` claim was missing from the access token")
		clientId = config.clientId
	}

	objectId := tokenClaims.ObjectId
	if objectId == "" {
		// Initialize an MS Graph client for the target cloud
		authProvider, err := auth.NewAzureIdentityAuthenticationProviderWithScopes(cred, []string{scope})
		if err != nil {
			return nil, err
		}
		adapter, err := msgraphsdk.NewGraphRequestAdapter(authProvider)
		if err != nil {
			return nil, err
		}
		adapter.SetBaseUrl(fmt.Sprintf("https://%s/v1.0", endpoint))
		client := msgraphsdk.NewGraphServiceClient(adapter)

		// Lookup the object ID
		if authenticatedAsServicePrincipal {
			logging.V(9).Infof("Querying Microsoft Graph to discover authenticated service principal object ID because the `oid` claim was missing from the access token")
			id, err := servicePrincipalObjectID(ctx, client, clientId)
			if err != nil {
				return nil, fmt.Errorf("attempting to discover object ID for authenticated service principal with client ID %q: %w", clientId, err)
			}

			objectId = *id
		} else {
			logging.V(9).Infof("Querying Microsoft Graph to discover authenticated user principal object ID because the `oid` claim was missing from the access token")
			id, err := userPrincipalObjectID(ctx, client)
			if err != nil {
				return nil, fmt.Errorf("attempting to discover object ID for authenticated user principal: %w", err)
			}
			objectId = *id
		}
	}

	tenantId := tokenClaims.TenantId
	if tenantId == "" {
		logging.V(9).Infof("Using user-supplied TenantID because the `tid` claim was missing from the access token")
		tenantId = config.tenantId
	}

	account := &ClientConfig{
		Cloud: cred.Cloud,

		ClientID:       clientId,
		ObjectID:       objectId,
		SubscriptionID: cred.SubscriptionId,
		TenantID:       tenantId,

		AuthenticatedAsAServicePrincipal: authenticatedAsServicePrincipal,
	}

	return account, nil
}

func servicePrincipalObjectID(ctx context.Context, client *msgraphsdk.GraphServiceClient, clientId string) (*string, error) {
	filter := fmt.Sprintf("appId eq '%s'", clientId)
	response, err := client.ServicePrincipals().Get(ctx, &serviceprincipals.ServicePrincipalsRequestBuilderGetRequestConfiguration{
		QueryParameters: &serviceprincipals.ServicePrincipalsRequestBuilderGetQueryParameters{
			Filter: &filter,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	principals := response.GetValue()
	if len(principals) != 1 {
		return nil, fmt.Errorf("unexpected number of results, expected 1, received %d", len(principals))
	}

	id := principals[0].GetId()
	if id == nil {
		return nil, errors.New("returned object ID was nil")
	}

	return id, nil
}

func userPrincipalObjectID(ctx context.Context, client *msgraphsdk.GraphServiceClient) (*string, error) {
	me, err := client.Me().Get(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	if me.GetId() == nil {
		return nil, fmt.Errorf("returned object ID was nil")
	}

	return me.GetId(), nil
}

// from: https://github.com/Azure/go-autorest/blob/autorest/v0.11.29/autorest/azure/environments.go
func getGraphEndpoint(cloud azcloud.Configuration) string {
	suffix := "graph.microsoft.com"
	if cloud.ActiveDirectoryAuthorityHost == azcloud.AzureChina.ActiveDirectoryAuthorityHost {
		suffix = "microsoftgraph.chinacloudapi.cn"
	} else if cloud.ActiveDirectoryAuthorityHost == azcloud.AzureGovernment.ActiveDirectoryAuthorityHost {
		suffix = "graph.microsoft.us"
	}
	return suffix
}
