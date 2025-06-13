// Copyright 2016-2022, Pulumi Corporation.

package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/go-autorest/autorest"
	azureEnv "github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure/cloud"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"

	goversion "github.com/hashicorp/go-version"
)

type authConfig struct {
	*authentication.Config
	useCli bool
}

func (a *authConfig) autorestEnvironment() (azureEnv.Environment, error) {
	envName := a.Environment
	env, err := azureEnv.EnvironmentFromName(envName)
	if err != nil {
		env, err = azureEnv.EnvironmentFromName(fmt.Sprintf("AZURE%sCLOUD", envName))
		if err != nil {
			return azureEnv.Environment{}, errors.Wrapf(err, "environment %q was not found", envName)
		}
	}
	return env, nil
}

func (a *authConfig) cloud() cloud.Configuration {
	cloudName := "public"
	if a.Config != nil && a.Config.Environment != "" {
		cloudName = a.Config.Environment
	}
	wellknown, _ := cloud.FromName(cloudName) // defaults to public
	return wellknown
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
	envName := readAzureEnvironmentFromConfig(k)

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

func readAzureEnvironmentFromConfig(k *azureNativeProvider) string {
	envName := k.getConfig("environment", "ARM_ENVIRONMENT")
	if envName == "" {
		envName = k.getConfig("environment", "AZURE_ENVIRONMENT")
	}
	if envName == "" {
		envName = "public"
	}
	return envName
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
	const versionHint = `Please make sure that the Azure CLI is installed in a version of at least 2.37 but less than 
3.x; or configure another authentication method. See 
https://www.pulumi.com/registry/packages/azure-native/installation-configuration/#credentials 
for more information.`

	// We need this version because it doesn't print the error of #1565
	versionRange := goversion.MustConstraints(goversion.NewConstraint(">=2.37.0, <3"))

	if !versionRange.Check(version) {
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
