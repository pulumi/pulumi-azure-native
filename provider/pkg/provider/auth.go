// Copyright 2016-2022, Pulumi Corporation.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/manicminer/hamilton/environments"
	"github.com/pkg/errors"
	"github.com/ryboe/q"

	goversion "github.com/hashicorp/go-version"
	hamiltonAuth "github.com/manicminer/hamilton-autorest/auth"
)

type authConfig struct {
	*authentication.Config
	useCli bool
}

func (k *azureNativeProvider) getAuthConfig() (*authConfig, error) {
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

	useMsi := k.getConfig("useMsi", "ARM_USE_MSI") == "true"
	clientSecret := k.getConfig("clientSecret", "ARM_CLIENT_SECRET")
	clientCertPath := k.getConfig("clientCertificatePath", "ARM_CLIENT_CERTIFICATE_PATH")

	useOIDC := k.getConfig("useOidc", "ARM_USE_OIDC") == "true"
	var oidcRequestToken, oidcRequestUrl string
	if useOIDC {
		oidcRequestToken = k.getConfig("oidcRequestToken", "ARM_OIDC_REQUEST_TOKEN")
		if oidcRequestToken == "" {
			// The ACTIONS_ variables are set by GitHub.
			oidcRequestToken = k.getConfig("oidcRequestToken", "ACTIONS_ID_TOKEN_REQUEST_TOKEN")
		}
		oidcRequestUrl = k.getConfig("oidcRequestUrl", "ARM_OIDC_REQUEST_URL")
		if oidcRequestUrl == "" {
			oidcRequestUrl = k.getConfig("oidcRequestUrl", "ACTIONS_ID_TOKEN_REQUEST_URL")
		}

		if oidcRequestToken == "" || oidcRequestUrl == "" {
			return nil, fmt.Errorf(`OIDC authentication is requested via useMsi/ARM_USE_MSI but no token and/or
request URL were configured. See
https://www.pulumi.com/registry/packages/azure-native/installation-configuration/#credentials for more information.`)
		}
	}

	// Without either of the methods above, we need the `az` CLI to authenticate. Check that we
	// have a good version (#1565). The check needs to happen before builder.Build(), or we return
	// the less fitting error message from go-azure-helpers.
	useCli := !useMsi && !useOIDC && clientSecret == "" && clientCertPath == ""
	if useCli {
		v, err := getAzVersion()
		if err != nil {
			return nil, err
		}
		if err = assertAzVersion(v); err != nil {
			return nil, err
		}
	}

	builder := &authentication.Builder{
		SubscriptionID: k.getConfig("subscriptionId", "ARM_SUBSCRIPTION_ID"),
		ClientID:       k.getConfig("clientId", "ARM_CLIENT_ID"),
		ClientSecret:   clientSecret,
		TenantID:       k.getConfig("tenantId", "ARM_TENANT_ID"),
		Environment:    envName,
		ClientCertPath: clientCertPath,
		ClientCertPassword: k.getConfig("clientCertificatePassword",
			"ARM_CLIENT_CERTIFICATE_PASSWORD"),
		MsiEndpoint:          k.getConfig("msiEndpoint", "ARM_MSI_ENDPOINT"),
		AuxiliaryTenantIDs:   auxTenants,
		ClientSecretDocsLink: "https://www.pulumi.com/docs/intro/cloud-providers/azure/setup/#service-principal-authentication",

		// OIDC section.
		IDTokenRequestToken: oidcRequestToken,
		IDTokenRequestURL:   oidcRequestUrl,
		IDToken:             k.getConfig("oidcToken", "ARM_OIDC_TOKEN"),
		IDTokenFilePath:     k.getConfig("oidcTokenFilePath", "ARM_OIDC_TOKEN_FILE_PATH"),
		MetadataHost:        k.getConfig("metadataHost", "ARM_METADATA_HOSTNAME"),

		// Feature Toggles
		SupportsClientCertAuth:         true,
		SupportsClientSecretAuth:       true,
		SupportsOIDCAuth:               useOIDC,
		UseMicrosoftGraph:              useOIDC,
		SupportsManagedServiceIdentity: useMsi,
		SupportsAzureCliToken:          !useOIDC,
		SupportsAuxiliaryTenants:       len(auxTenants) > 0,
	}

	c, err := builder.Build()
	if err != nil {
		return nil, err
	}

	return &authConfig{c, useCli}, nil
}

func getAzVersion() (*goversion.Version, error) {
	_, err := exec.LookPath("az")
	if err != nil {
		return nil, fmt.Errorf("could not find `az`: %w", err)
	}

	var azVersion struct {
		Cli string `json:"azure-cli"`
	}
	err = runAzCmd(&azVersion, "version")
	if err != nil {
		return nil, fmt.Errorf("could not determine az version: %w", err)
	}

	actual, err := goversion.NewVersion(azVersion.Cli)
	if err != nil {
		return nil, fmt.Errorf("could not parse az version \"%q\": %w", azVersion.Cli, err)
	}

	return actual, nil
}

func assertAzVersion(version *goversion.Version) error {
	const versionHint = `Please make sure that the Azure CLI is installed in a version either
between 2.0.81 and 2.33, or at least 2.37 but less than 3.x; or configure another authentication
method. See https://www.pulumi.com/registry/packages/azure-native/installation-configuration/#credentials
for more information.`

	// We need this version because it doesn't print the error of #1565
	lowerOkRange := goversion.MustConstraints(goversion.NewConstraint(">=2.0.81, <2.34"))
	upperOkRange := goversion.MustConstraints(goversion.NewConstraint(">=2.37.0, <3"))

	if !lowerOkRange.Check(version) && !upperOkRange.Check(version) {
		return fmt.Errorf("found incompatible az version %s. %s", version, versionHint)
	}

	return nil
}

// Run `az` with the given args and unmarshal the output (must be JSON) into the given target struct
func runAzCmd(target interface{}, arg ...string) error {
	var stderr bytes.Buffer
	var stdout bytes.Buffer

	cmd := exec.Command("az", arg...)

	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		err := fmt.Errorf("running az: %w", err)
		if stdErrStr := stderr.String(); stdErrStr != "" {
			err = fmt.Errorf("%s: %s", err, strings.TrimSpace(stdErrStr))
		}
		return err
	}

	if err := json.Unmarshal(stdout.Bytes(), target); err != nil {
		return fmt.Errorf("unmarshaling the result of Azure CLI: %w", err)
	}

	return nil
}

type AuthorizerFactory func(api environments.Api) (autorest.Authorizer, error)

func (k *azureNativeProvider) makeAuthorizerFactories(ctx context.Context,
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

func (k *azureNativeProvider) makeADALAuthorizerFactories(ctx context.Context,
	authConfig *authentication.Config, oauthConfig *authentication.OAuthConfig, endpoint string, autorestSender autorest.Sender) (AuthorizerFactory, AuthorizerFactory, error) {

	tokenFactory := func(api environments.Api) (autorest.Authorizer, error) {
		q.Q("Getting ADAL token", api.Endpoint, endpoint)
		return authConfig.GetADALToken(ctx, autorestSender, oauthConfig, endpoint)
	}

	bearerAuthFactory := func(api environments.Api) (autorest.Authorizer, error) {
		q.Q("Getting ADAL bearer auth callback", api.Endpoint, endpoint)
		return authConfig.ADALBearerAuthorizerCallback(ctx, autorestSender, oauthConfig), nil
	}

	return tokenFactory, bearerAuthFactory, nil
}

func (k *azureNativeProvider) makeMSALAuthorizerFactories(ctx context.Context, authConfig *authentication.Config,
	oauthConfig *authentication.OAuthConfig, endpoint string, autorestSender autorest.Sender) (AuthorizerFactory, AuthorizerFactory, error) {

	tokenFactory := func(api environments.Api) (autorest.Authorizer, error) {
		q.Q("Getting MSAL token", api.Endpoint, endpoint)
		return authConfig.GetMSALToken(ctx, api, autorestSender, oauthConfig, endpoint)
	}

	bearerAuthFactory := func(api environments.Api) (autorest.Authorizer, error) {
		q.Q("Getting MSAL bearer auth callback", api.Endpoint, endpoint)
		return authConfig.MSALBearerAuthorizerCallback(ctx, api, autorestSender, oauthConfig, endpoint), nil
	}

	return tokenFactory, bearerAuthFactory, nil
}

func (k *azureNativeProvider) getOAuthToken(ctx context.Context, auth *authConfig, endpoint string) (string, error) {
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

func (k *azureNativeProvider) buildOAuthConfig(authConfig *authentication.Config) (*authentication.OAuthConfig, error) {
	oauthConfig, err := authConfig.BuildOAuthConfig(k.environment.ActiveDirectoryEndpoint)
	if err != nil {
		return nil, err
	}
	if oauthConfig == nil {
		return nil, fmt.Errorf("unable to configure OAuthConfig for tenant %s", authConfig.TenantID)
	}
	return oauthConfig, nil
}
