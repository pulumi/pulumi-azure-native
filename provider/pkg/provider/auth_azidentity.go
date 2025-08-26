// Copyright 2016-2022, Pulumi Corporation.

package provider

import (
	"context"
	"crypto"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	_ "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime" // initialize the well-known cloud configurations
	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const (
	cliCloudMismatchMessage = `
The configured Azure cloud '%s' does not match the active cloud '%s'.
When authenticating using the Azure CLI, the configured cloud needs to match the one shown by 'az account show'.
You can change clouds using 'az cloud set --name <cloud>'.`
)

// azAccount is the provider's resolved Azure account for authentication and for resource management.
type azAccount struct {
	azcore.TokenCredential

	// The Azure cloud configuration to use for resource management operations.
	Cloud azcloud.Configuration

	// The subscription ID to use for resource management operations; empty if not configured or auto-detected.
	SubscriptionId string
}

// NewAzCoreIdentity is the main entry to the new azcore/azidentity-based authentication stack.
// It determines the provider's Azure account such as the cloud and subscription ID, and a
// TokenCredential which can be passed into various Azure Go SDKs.
func NewAzCoreIdentity(ctx context.Context, authConf *authConfiguration, baseClientOpts policy.ClientOptions) (*azAccount, error) {
	account := &azAccount{}

	if authConf.cloud != nil {
		baseClientOpts.Cloud = *authConf.cloud
	}

	// Create the azcore.TokenCredential implementation based on the auth configuration.
	// This routine evaluates the auth configuration and other environment variables,
	// and ultimately resolves the Azure cloud and subscription ID.
	cred, err := newSingleMethodAuthCredential(authConf, baseClientOpts)
	if err != nil {
		return nil, err
	}
	account.TokenCredential = cred

	// Based on the credential type, we may be able to resolve the Azure cloud and subscription ID.
	if _, ok := cred.(*azidentity.AzureCLICredential); ok {
		// get the given (or default) account from the Azure CLI
		activeSubscription, err := authConf.showSubscription(ctx, authConf.subscriptionId)
		if err != nil {
			return nil, err
		}
		logging.V(6).Infof("Using Az account %q", activeSubscription.Name)

		// automatically configure the environment and/or subscription ID based on the Azure CLI account.
		sc := azure.GetCloudByName(activeSubscription.EnvironmentName)
		if authConf.cloud != nil && sc.ActiveDirectoryAuthorityHost != authConf.cloud.ActiveDirectoryAuthorityHost {
			return nil, fmt.Errorf(cliCloudMismatchMessage, azure.GetCloudName(*authConf.cloud), activeSubscription.EnvironmentName)
		}
		account.Cloud = sc
		account.SubscriptionId = activeSubscription.ID
	} else {
		// use the configured values and/or defaults
		if authConf.cloud == nil {
			// if the cloud is not set, use the default Azure cloud as does newSingleMethodAuthCredential.
			account.Cloud = azcloud.AzurePublic
		} else {
			account.Cloud = *authConf.cloud
		}
		account.SubscriptionId = authConf.subscriptionId
	}

	logging.V(6).Infof("Using Az cloud %q with subscription ID %q", azure.GetCloudName(account.Cloud), account.SubscriptionId)
	return account, nil
}

// newSingleMethodAuthCredential creates an azcore.TokenCredential. Depending on the given authConfiguration, it is
// backed by one of several authentication methods such as Service Principal, OIDC, Managed Identity, or Azure CLI.
//
// Note: this function's behavior is written to match the behavior of
// "github.com/hashicorp/go-azure-helpers/authentication".Builder.Build() in some ways, to minimize changes in provider
// behavior when upgrading authentication dependencies from go-azure-helpers to azidentity. Namely:
//   - The order in which the the different authentication methods are attempted is the same:
//     1. Default Credential
//     2. Service Principal with client certificate
//     3. Service Principal with client secret
//     4. OIDC
//     5. Managed Identity
//     6. Azure CLI
//   - When a method is configured but instantiating the credential fails, we return an error and do not fall through to
//     the next method.
//   - Auxiliary or additional tenants are supported for SP with client secret and CLI authentication, not for others.
func newSingleMethodAuthCredential(authConf *authConfiguration, baseClientOpts azcore.ClientOptions) (azcore.TokenCredential, error) {
	if authConf.useDefault {
		logging.V(9).Infof("[auth] Using default Azure credential")
		fmtErrorMessage := "A %s must be configured when authenticating using the Default Azure Credential."
		if authConf.subscriptionId == "" {
			return nil, fmt.Errorf(fmtErrorMessage, "Subscription ID")
		}
		options := &azidentity.DefaultAzureCredentialOptions{
			ClientOptions:              baseClientOpts,
			AdditionallyAllowedTenants: authConf.auxTenants, // usually empty which is fine
		}
		return azidentity.NewDefaultAzureCredential(options)
	} else {
		logging.V(11).Infof("Default Azure Credential is not enabled, skipping")
	}

	if authConf.clientCertPath != "" {
		logging.V(9).Infof("[auth] Using SP with client certificate credential")
		fmtErrorMessage := "A %s must be configured when authenticating as a Service Principal using a Client Certificate."
		if authConf.subscriptionId == "" {
			return nil, fmt.Errorf(fmtErrorMessage, "Subscription ID")
		}
		if authConf.tenantId == "" {
			return nil, fmt.Errorf(fmtErrorMessage, "Tenant ID")
		}
		if authConf.clientId == "" {
			return nil, fmt.Errorf(fmtErrorMessage, "Client ID")
		}
		certs, key, err := readCertificate(authConf.clientCertPath, authConf.clientCertPassword)
		if err != nil {
			return nil, fmt.Errorf("the Client Certificate Path is not a valid pfx file: %w", err)
		}
		options := &azidentity.ClientCertificateCredentialOptions{
			AdditionallyAllowedTenants: authConf.auxTenants, // usually empty which is fine
			ClientOptions:              baseClientOpts,
			DisableInstanceDiscovery:   authConf.disableInstanceDiscovery,
		}
		return azidentity.NewClientCertificateCredential(authConf.tenantId, authConf.clientId, certs, key, options)
	} else {
		logging.V(11).Infof("SP with client certificate credential is not enabled, skipping")
	}

	if authConf.clientSecret != "" {
		logging.V(9).Infof("[auth] Using SP with client secret credential")
		fmtErrorMessage := "A %s must be configured when authenticating as a Service Principal using a Client Secret."
		if authConf.subscriptionId == "" {
			return nil, fmt.Errorf(fmtErrorMessage, "Subscription ID")
		}
		if authConf.tenantId == "" {
			return nil, fmt.Errorf(fmtErrorMessage, "Tenant ID")
		}
		if authConf.clientId == "" {
			return nil, fmt.Errorf(fmtErrorMessage, "Client ID")
		}
		options := &azidentity.ClientSecretCredentialOptions{
			AdditionallyAllowedTenants: authConf.auxTenants, // usually empty which is fine
			ClientOptions:              baseClientOpts,
			DisableInstanceDiscovery:   authConf.disableInstanceDiscovery,
		}
		return azidentity.NewClientSecretCredential(authConf.tenantId, authConf.clientId, authConf.clientSecret, options)
	} else {
		logging.V(11).Infof("SP with client secret credential is not enabled, skipping")
	}

	if authConf.useOidc {
		logging.V(9).Infof("[auth] Using OIDC credential")
		return newOidcCredential(authConf, baseClientOpts)
	} else {
		logging.V(11).Infof("OIDC credential is not enabled, skipping")
	}

	if authConf.useMsi {
		logging.V(9).Infof("[auth] Using Managed Identity (MSI) credential")
		fmtErrorMessage := "A %s must be configured when authenticating as a Managed Identity using MSI."
		if authConf.subscriptionId == "" {
			return nil, fmt.Errorf(fmtErrorMessage, "Subscription ID")
		}
		msiOpts := azidentity.ManagedIdentityCredentialOptions{
			ClientOptions: baseClientOpts,
		}
		if authConf.clientId != "" {
			msiOpts.ID = azidentity.ClientID(authConf.clientId)
		}
		return azidentity.NewManagedIdentityCredential(&msiOpts)
	} else {
		logging.V(11).Infof("Managed Identity (MSI) credential is not enabled, skipping")
	}

	logging.V(9).Infof("[auth] Using Azure CLI credential")
	options := &azidentity.AzureCLICredentialOptions{
		AdditionallyAllowedTenants: authConf.auxTenants, // usually empty which is fine
	}
	// note that the subscription ID is discoverable when using the Azure CLI credential and hence optional.
	if authConf.subscriptionId != "" {
		options.Subscription = authConf.subscriptionId
	}
	return azidentity.NewAzureCLICredential(options)
}

// newOidcCredential creates a TokenCredential for OpenID Connect (OIDC) authentication.
// An OIDC credential is an azidentity.ClientAssertionCredential that authenticates with a token
// obtained from a callback function. The token itself can be provided in various ways:
// - directly via config/environment variable
// - from a file
// - through a token exchange by making a request to a configured endpoint
// This function configures the client assertion callback according to the above cases.
func newOidcCredential(authConf *authConfiguration, clientOpts azcore.ClientOptions) (azcore.TokenCredential, error) {
	fmtErrorMessage := "A %s must be configured when authenticating with OIDC."
	if authConf.subscriptionId == "" {
		return nil, fmt.Errorf(fmtErrorMessage, "Subscription ID")
	}
	if authConf.tenantId == "" {
		return nil, fmt.Errorf(fmtErrorMessage, "Tenant ID")
	}
	if authConf.clientId == "" {
		return nil, fmt.Errorf(fmtErrorMessage, "Client ID")
	}

	// The generic client assertion that simply returns the token it was created with.
	oidcTokenCredentialCallback := func(token string) (azcore.TokenCredential, error) {
		return azidentity.NewClientAssertionCredential(
			authConf.tenantId,
			authConf.clientId,
			func(ctx context.Context) (string, error) {
				return token, nil
			},
			nil)
	}

	if authConf.oidcToken != "" {
		return oidcTokenCredentialCallback(authConf.oidcToken)
	}

	if authConf.oidcTokenFilePath != "" {
		token, err := os.ReadFile(authConf.oidcTokenFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read OIDC token from %s: %w", authConf.oidcTokenFilePath, err)
		}
		return oidcTokenCredentialCallback(string(token))
	}

	// In this case, we need to obtain the OIDC token first from the configured endpoint.
	if authConf.oidcTokenRequestUrl != "" && authConf.oidcTokenRequestToken != "" {
		return azidentity.NewClientAssertionCredential(
			authConf.tenantId,
			authConf.clientId,
			getOidcTokenExchangeAssertion(authConf),
			&azidentity.ClientAssertionCredentialOptions{
				ClientOptions: clientOpts,
			})
	}

	return nil, fmt.Errorf(fmtErrorMessage, "OIDC token or request URL and token")
}

// getOidcTokenExchangeAssertion returns a callback function that implements the OIDC token
// exchange flow on GitHub. The function makes a request to the configured endpoint with the
// configured bearer token and returns the OIDC token from the response. It's intended to be used
// in an azidentity.ClientAssertionCredential.
func getOidcTokenExchangeAssertion(authConf *authConfiguration) func(ctx context.Context) (string, error) {
	return func(ctx context.Context) (string, error) {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, authConf.oidcTokenRequestUrl, http.NoBody)
		if err != nil {
			return "", fmt.Errorf("GitHub OIDC: failed to build request to %s: %w", authConf.oidcTokenRequestUrl, err)
		}

		query, err := url.ParseQuery(req.URL.RawQuery)
		if err != nil {
			return "", fmt.Errorf("githubAssertion: cannot parse URL query")
		}

		// see https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/configuring-openid-connect-in-azure#adding-the-federated-credentials-to-azure
		audience := "api://AzureADTokenExchange"
		if authConf.oidcAudience != "" {
			audience = authConf.oidcAudience
		}
		query.Set("audience", audience)
		req.URL.RawQuery = query.Encode()

		req.Header.Set("Accept", "application/json")
		// see https://docs.github.com/en/actions/security-for-github-actions/security-hardening-your-deployments/about-security-hardening-with-openid-connect#updating-your-actions-for-oidc
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authConf.oidcTokenRequestToken))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", fmt.Errorf("GitHub OIDC: couldn't request token: %w", err)
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("GitHub OIDC: cannot read response: %w", err)
		}

		if c := resp.StatusCode; c < 200 || c > 299 {
			return "", fmt.Errorf("GitHub OIDC: %d with response: %s", resp.StatusCode, body)
		}

		var tokenResponse struct {
			Value string `json:"value"`
		}
		if err := json.Unmarshal(body, &tokenResponse); err != nil {
			return "", fmt.Errorf("GitHub OIDC: cannot unmarshal response: %w", err)
		}

		return tokenResponse.Value, nil
	}
}

func readCertificate(certPath, certPassword string) ([]*x509.Certificate, crypto.PrivateKey, error) {
	cert, err := os.ReadFile(certPath)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to read certificate from %s", certPath)
	}

	var pwBytes []byte
	if certPassword != "" {
		pwBytes = []byte(certPassword)
	}

	certs, key, err := azidentity.ParseCertificates(cert, pwBytes)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to parse certificate from %s", certPath)
	}
	if len(certs) == 0 {
		return nil, nil, errors.Errorf("no certificates found in %s", certPath)
	}

	return certs, key, nil
}

type authConfiguration struct {
	cloud *azcloud.Configuration

	subscriptionId string

	clientId   string
	tenantId   string
	auxTenants []string

	useOidc               bool
	oidcAudience          string
	oidcToken             string
	oidcTokenFilePath     string
	oidcTokenRequestToken string
	oidcTokenRequestUrl   string

	clientSecret       string
	clientCertPath     string
	clientCertPassword string

	// Note: for MSI, there used to be a msiEndpoint/ARM_MSI_ENDPOINT and a metadataHost/
	// ARM_METADATA_HOSTNAME configuration. The newer azidentity package handles the MSI endpoint
	// automatically:
	// https://github.com/Azure/azure-sdk-for-go/blob/sdk/azidentity/v1.8.0/sdk/azidentity/managed_identity_client.go#L143
	useMsi bool

	// Enables the use of azcore's DefaultAzureCredential strategy.
	// DefaultAzureCredential simplifies authentication while developing applications that deploy to Azure by
	// combining credentials used in Azure hosting environments and credentials used in local development.
	useDefault bool

	// When true, disables instance discovery for credentials that support it.
	disableInstanceDiscovery bool

	// showSubscription invokes `az account show` and is overridable by tests to fake invoking the az CLI.
	showSubscription azSubscriptionProvider
}

type configGetter func(configName, envName string) string

// readAuthConfig collects auth-related configuration from Pulumi config and environment variables
func readAuthConfig(getConfig configGetter) (*authConfiguration, error) {
	cloud := getCloud(getConfig)

	auxTenantsString := getConfig("auxiliaryTenantIds", "ARM_AUXILIARY_TENANT_IDS")
	var auxTenants []string
	if auxTenantsString != "" {
		err := json.Unmarshal([]byte(auxTenantsString), &auxTenants)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal '%s' as Auxiliary Tenants", auxTenantsString)
		}
	}

	authConf := &authConfiguration{
		clientId:   getConfig("clientId", "ARM_CLIENT_ID"),
		tenantId:   getConfig("tenantId", "ARM_TENANT_ID"),
		auxTenants: auxTenants,
		cloud:      cloud,

		subscriptionId: getConfig("subscriptionId", "ARM_SUBSCRIPTION_ID"),

		clientSecret:       getConfig("clientSecret", "ARM_CLIENT_SECRET"),
		clientCertPath:     getConfig("clientCertificatePath", "ARM_CLIENT_CERTIFICATE_PATH"),
		clientCertPassword: getConfig("clientCertificatePassword", "ARM_CLIENT_CERTIFICATE_PASSWORD"),

		useMsi: getConfig("useMsi", "ARM_USE_MSI") == "true",

		useOidc:               getConfig("useOidc", "ARM_USE_OIDC") == "true",
		oidcAudience:          getConfig("oidcAudience", "ARM_OIDC_AUDIENCE"),
		oidcToken:             getConfig("oidcToken", "ARM_OIDC_TOKEN"),
		oidcTokenFilePath:     getConfig("oidcTokenFilePath", "ARM_OIDC_TOKEN_FILE_PATH"),
		oidcTokenRequestToken: getConfig("oidcRequestToken", "ACTIONS_ID_TOKEN_REQUEST_TOKEN"),
		oidcTokenRequestUrl:   getConfig("oidcRequestUrl", "ACTIONS_ID_TOKEN_REQUEST_URL"),

		useDefault: getConfig("useDefaultAzureCredential", "ARM_USE_DEFAULT_AZURE_CREDENTIAL") == "true",

		disableInstanceDiscovery: getConfig("disableInstanceDiscovery", "ARM_DISABLE_INSTANCE_DISCOVERY") == "true",

		showSubscription: defaultAzSubscriptionProvider,
	}

	return authConf, nil
}

// getCloud returns the configured Azure cloud (environment).
// Returns nil if not configured, to allow for other detection methods before defaulting to the public cloud.
func getCloud(getConfig configGetter) *azcloud.Configuration {
	activeDirectoryAuthorityHost := getConfig("activeDirectoryAuthorityHost", "ARM_ACTIVE_DIRECTORY_AUTHORITY_HOST")
	resourceManagerAudience := getConfig("resourceManagerAudience", "ARM_RESOURCE_MANAGER_AUDIENCE")
	resourceManagerEndpoint := getConfig("resourceManagerEndpoint", "ARM_RESOURCE_MANAGER_ENDPOINT")
	if activeDirectoryAuthorityHost != "" && resourceManagerAudience != "" && resourceManagerEndpoint != "" {
		return &azcloud.Configuration{
			ActiveDirectoryAuthorityHost: activeDirectoryAuthorityHost,
			Services: map[azcloud.ServiceName]azcloud.ServiceConfiguration{
				azcloud.ResourceManager: {
					Audience: resourceManagerAudience,
					Endpoint: resourceManagerEndpoint,
				},
			},
		}
	}

	// Otherwise fall back to using the Environment Name from the fixed list
	envName := getConfig("environment", "ARM_ENVIRONMENT")
	if envName == "" {
		envName = getConfig("environment", "AZURE_ENVIRONMENT")
	}
	if envName != "" {
		cloud := azure.GetCloudByName(envName)
		return &cloud
	}
	return nil
}
