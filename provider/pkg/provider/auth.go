// Copyright 2016-2022, Pulumi Corporation.

package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/manicminer/hamilton/environments"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"

	goversion "github.com/hashicorp/go-version"
	hamiltonAuth "github.com/manicminer/hamilton-autorest/auth"
)

type authConfig struct {
	*authentication.Config
	useCli bool
}

type oidcConfig struct {
	oidcToken         string
	oidcTokenFilePath string
	oidcRequestToken  string
	oidcRequestUrl    string
}

// Assumes that OIDC authentication is requested. Accordingly, returns an error if OIDC is not
// properly configured. If ARM_OIDC_TOKEN/azure-native:oidcToken is set, it is preferred.
func (k *azureNativeProvider) determineOidcConfig() (oidcConfig, error) {
	oidcToken := k.getConfig("oidcToken", "ARM_OIDC_TOKEN")
	if oidcToken != "" {
		return oidcConfig{
			oidcToken: oidcToken,
		}, nil
	}

	oidcTokenFilePath := k.getConfig("oidcTokenFilePath", "ARM_OIDC_TOKEN_FILE_PATH")
	if oidcTokenFilePath != "" {
		return oidcConfig{
			oidcTokenFilePath: oidcTokenFilePath,
		}, nil
	}

	oidcRequestToken := k.getConfig("oidcRequestToken", "ARM_OIDC_REQUEST_TOKEN")
	oidcRequestUrl := k.getConfig("oidcRequestUrl", "ARM_OIDC_REQUEST_URL")
	if oidcRequestToken != "" && oidcRequestUrl != "" {
		return oidcConfig{
			oidcRequestToken: oidcRequestToken,
			oidcRequestUrl:   oidcRequestUrl,
		}, nil
	}

	// GitHub Actions pre-set env
	oidcRequestToken = k.getConfig("oidcRequestToken", "ACTIONS_ID_TOKEN_REQUEST_TOKEN")
	oidcRequestUrl = k.getConfig("oidcRequestUrl", "ACTIONS_ID_TOKEN_REQUEST_URL")
	if oidcRequestToken != "" && oidcRequestUrl != "" {
		return oidcConfig{
			oidcRequestToken: oidcRequestToken,
			oidcRequestUrl:   oidcRequestUrl,
		}, nil
	}

	return oidcConfig{}, fmt.Errorf(`OIDC authentication was requested via useOidc/ARM_USE_OIDC but no token or request URL were 
configured. See https://www.pulumi.com/registry/packages/azure-native/installation-configuration/#credentials for more information.`)
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

// Without any of the other methods above, we fall back to the `az` CLI to authenticate.
func needAzCli(builder *authentication.Builder) bool {
	return !builder.SupportsManagedServiceIdentity &&
		!builder.SupportsOIDCAuth &&
		builder.ClientSecret == "" &&
		builder.ClientCertPath == ""
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

// Run `az` with the given args and unmarshal the output into the given target struct.
// The `az` output is always in JSON, requested via `--output json`.
func runAzCmd(target interface{}, arg ...string) error {
	var stderr bytes.Buffer
	var stdout bytes.Buffer

	arg = append(arg, "--output", "json")

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
		logging.V(9).Infof("Getting ADAL token for %s", endpoint)
		return authConfig.GetADALToken(ctx, autorestSender, oauthConfig, endpoint)
	}

	bearerAuthFactory := func(api environments.Api) (autorest.Authorizer, error) {
		logging.V(9).Infof("Getting ADAL bearer auth callback for %s", endpoint)
		return authConfig.ADALBearerAuthorizerCallback(ctx, autorestSender, oauthConfig), nil
	}

	return tokenFactory, bearerAuthFactory, nil
}

func (k *azureNativeProvider) makeMSALAuthorizerFactories(ctx context.Context, authConfig *authentication.Config,
	oauthConfig *authentication.OAuthConfig, endpoint string, autorestSender autorest.Sender) (AuthorizerFactory, AuthorizerFactory, error) {

	tokenFactory := func(api environments.Api) (autorest.Authorizer, error) {
		logging.V(9).Infof("Getting MSAL token for %s", endpoint)
		return authConfig.GetMSALToken(ctx, api, autorestSender, oauthConfig, endpoint)
	}

	bearerAuthFactory := func(api environments.Api) (autorest.Authorizer, error) {
		logging.V(9).Infof("Getting MSAL bearer auth callback for %s", endpoint)
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

// Implements the `azidentity.TokenCredential` interface.
type azCoreTokenCredential struct {
	p        *azureNativeProvider
	endpoint string
}

func (cred azCoreTokenCredential) GetToken(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
	authConfig, err := cred.p.getAuthConfig()
	if err != nil {
		return azcore.AccessToken{}, err
	}
	endpoint := cred.endpoint
	if endpoint == "" {
		endpoint = cred.p.environment.ResourceManagerEndpoint
	}
	token, err := cred.p.getOAuthToken(ctx, authConfig, endpoint)
	if err != nil {
		return azcore.AccessToken{}, err
	}
	return azcore.AccessToken{
		Token: token,
		// This hard-coded expiry is not ideal but we don't know the lifetime of the token at this
		// point because the Azure response containing it is down the call stack in go-azure-helpers.
		ExpiresOn: time.Now().Add(2 * time.Hour),
	}, nil
}

var _ azcore.TokenCredential = (*azCoreTokenCredential)(nil)

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
