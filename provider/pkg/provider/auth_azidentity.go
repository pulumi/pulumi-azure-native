// Copyright 2016-2022, Pulumi Corporation.

package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/pkg/errors"

	hamiltonAuth "github.com/manicminer/hamilton-autorest/auth"
)

func (k *azureNativeProvider) getAuthConfig2() (*authConfig, error) {
	auxTenantsString := k.getConfig("auxiliaryTenantIds", "ARM_AUXILIARY_TENANT_IDS")
	var auxTenants []string
	if auxTenantsString != "" {
		err := json.Unmarshal([]byte(auxTenantsString), &auxTenants)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal '%s' as Auxiliary Tenants", auxTenantsString)
		}
	}
	envName := k.getConfig("environment", "ARM_ENVIRONMENT")
	if envName == "" {
		envName = "public"
	}

	useOIDC := k.getConfig("useOidc", "ARM_USE_OIDC") == "true"
	oidcConf, err := k.determineOidcConfig()
	if useOIDC && err != nil {
		return nil, err
	}

	builder := &authentication.Builder{
		SubscriptionID:       k.getConfig("subscriptionId", "ARM_SUBSCRIPTION_ID"),
		ClientID:             k.getConfig("clientId", "ARM_CLIENT_ID"),
		TenantID:             k.getConfig("tenantId", "ARM_TENANT_ID"),
		Environment:          envName,
		MsiEndpoint:          k.getConfig("msiEndpoint", "ARM_MSI_ENDPOINT"),
		AuxiliaryTenantIDs:   auxTenants,
		ClientSecretDocsLink: "https://www.pulumi.com/docs/intro/cloud-providers/azure/setup/#service-principal-authentication",
		MetadataHost:         k.getConfig("metadataHost", "ARM_METADATA_HOSTNAME"),

		// Service Principal section.
		ClientCertPath:     k.getConfig("clientCertificatePath", "ARM_CLIENT_CERTIFICATE_PATH"),
		ClientCertPassword: k.getConfig("clientCertificatePassword", "ARM_CLIENT_CERTIFICATE_PASSWORD"),
		ClientSecret:       k.getConfig("clientSecret", "ARM_CLIENT_SECRET"),

		// OIDC section.
		IDTokenRequestToken: oidcConf.oidcRequestToken,
		IDTokenRequestURL:   oidcConf.oidcRequestUrl,
		IDToken:             oidcConf.oidcToken,
		IDTokenFilePath:     oidcConf.oidcTokenFilePath,

		// Feature Toggles
		SupportsClientCertAuth:         true,
		SupportsClientSecretAuth:       true,
		SupportsOIDCAuth:               useOIDC,
		UseMicrosoftGraph:              useOIDC,
		SupportsManagedServiceIdentity: k.getConfig("useMsi", "ARM_USE_MSI") == "true",
		SupportsAzureCliToken:          !useOIDC,
		SupportsAuxiliaryTenants:       len(auxTenants) > 0,
	}

	needCli := needAzCli(builder)
	if needCli {
		requireCliHint := "MSI, OIDC, client secret and client certificate authentication are not configured so we require the AZ CLI for authentication"
		// Check that we have a good version of the cli (#1565). The check needs to happen before
		// builder.Build(), or we return the less fitting error message from go-azure-helpers.
		v, err := getAzVersion()
		if err != nil {
			return nil, fmt.Errorf("checking az cli version: %w: %s", err, requireCliHint)
		}
		if err = assertAzVersion(v); err != nil {
			return nil, err
		}
	}

	c, err := builder.Build()
	if err != nil {
		return nil, err
	}

	// CLI authentication can only use the environment that the CLI is configured for.
	// Fail early if that's not the one from config, to avoid obscure errors about missing subscriptions or endpoints later.
	if c.Environment != envName {
		return nil, errors.Errorf(`The configured Azure environment '%s' does not match the determined environment '%s'.
When authenticating using the Azure CLI, the configured environment needs to match the one shown by 'az account show'.
You can change environments using 'az cloud set --name <environment>'.`,
			envName, c.Environment)
	}

	return &authConfig{c, needCli}, nil
}

// type AuthorizerFactory func(api environments.Api) (autorest.Authorizer, error)

func (k *azureNativeProvider) makeAuthorizerFactories2(ctx context.Context,
	authConfig *authConfig) (AuthorizerFactory, AuthorizerFactory, error) {

	buildSender := sender.BuildSender("AzureNative")

	oauthConfig, err := k.buildOAuthConfig(authConfig.Config)
	if err != nil {
		return nil, nil, err
	}

	endpoint := k.environment.TokenAudience

	if authConfig.useCli {
		return k.makeADALAuthorizerFactories(ctx, authConfig.Config, oauthConfig, endpoint, buildSender)
	}
	return k.makeMSALAuthorizerFactories(ctx, authConfig.Config, oauthConfig, endpoint, buildSender)
}

func (k *azureNativeProvider) getOAuthToken2(ctx context.Context, auth *authConfig, endpoint string) (string, error) {
	buildSender := sender.BuildSender("AzureNative")
	oauthConfig, err := k.buildOAuthConfig(auth.Config)
	if err != nil {
		return "", err
	}

	api := k.autorestEnvToHamiltonEnv().ResourceManager

	var authorizer autorest.Authorizer
	if auth.useCli {
		authorizer, err = auth.GetADALToken(ctx, buildSender, oauthConfig, endpoint)
	} else {
		authorizer, err = auth.GetMSALToken(ctx, api, buildSender, oauthConfig, endpoint)
	}
	if err != nil {
		return "", fmt.Errorf("getting authorization token: %w", err)
	}

	// go-azure-helpers returns different kinds of Authorizer from different auth methods so we
	// need to check to choose the right method to get a token.
	var token string
	ba, ok := authorizer.(*autorest.BearerAuthorizer)
	if ok {
		tokenProvider := ba.TokenProvider()
		token = tokenProvider.OAuthToken()
	} else {
		if outer, ok := authorizer.(*hamiltonAuth.Authorizer); ok {
			t, err := outer.Token()
			if err != nil {
				return "", err
			}
			token = t.AccessToken
		}
	}

	if token == "" {
		return "", fmt.Errorf("empty token from %T", authorizer)
	}
	return token, nil
}
