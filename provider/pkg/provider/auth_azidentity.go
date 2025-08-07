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
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	"github.com/microsoftgraph/msgraph-sdk-go/serviceprincipals"
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
	switch cred.(type) {
	case *azidentity.AzureCLICredential:
		// get the given (or default if id is empty) account from the Azure CLI
		activeSubscription, err := authConf.showSubscription(ctx, authConf.subscriptionId)
		if err != nil {
			return nil, fmt.Errorf("failed to get account from Azure CLI: %w", err)
		}
		logging.V(6).Infof("Using Az account %q", activeSubscription.Name)

		// automatically configure the environment and/or subscription ID based on the Azure CLI account.
		sc := azure.GetCloudByName(activeSubscription.EnvironmentName)
		if authConf.cloud != nil && sc.ActiveDirectoryAuthorityHost != authConf.cloud.ActiveDirectoryAuthorityHost {
			return nil, fmt.Errorf(cliCloudMismatchMessage, azure.GetCloudName(*authConf.cloud), activeSubscription.EnvironmentName)
		}
		account.Cloud = sc
		account.SubscriptionId = activeSubscription.ID

	default:
		if authConf.cloud == nil {
			// if the cloud is not set, use the default Azure cloud
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
//     1. Service Principal with client certificate
//     2. Service Principal with client secret
//     3. OIDC
//     4. Managed Identity
//     5. Azure CLI
//   - When a method is configured but instantiating the credential fails, we return an error and do not fall through to
//     the next method.
//   - Auxiliary or additional tenants are supported for SP with client secret and CLI authentication, not for others.
func newSingleMethodAuthCredential(authConf *authConfiguration, baseClientOpts azcore.ClientOptions) (azcore.TokenCredential, error) {
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
		}
		return azidentity.NewClientCertificateCredential(authConf.tenantId, authConf.clientId, certs, key, options)
	} else {
		logging.V(9).Infof("SP with client certificate credential is not enabled, skipping")
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
		}
		return azidentity.NewClientSecretCredential(authConf.tenantId, authConf.clientId, authConf.clientSecret, options)
	} else {
		logging.V(9).Infof("SP with client secret credential is not enabled, skipping")
	}

	if authConf.useOidc {
		logging.V(9).Infof("[auth] Using OIDC credential")
		return newOidcCredential(authConf, baseClientOpts)
	} else {
		logging.V(9).Infof("OIDC credential is not enabled, skipping")
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
		logging.V(9).Infof("Managed Identity (MSI) credential is not enabled, skipping")
	}

	logging.V(9).Infof("[auth] Using Azure CLI credential")
	options := &azidentity.AzureCLICredentialOptions{
		AdditionallyAllowedTenants: authConf.auxTenants, // usually empty which is fine
	}
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

	// showSubscription invokes `az account show` and is overridable by tests to fake invoking the az CLI.
	showSubscription azSubscriptionProvider
}

type configGetter func(configName, envName string) string

// readAuthConfig collects auth-related configuration from Pulumi config and environment variables
func readAuthConfig(getConfig configGetter) (*authConfiguration, error) {
	cloud := readAzureCloudFromConfig(getConfig)

	auxTenantsString := getConfig("auxiliaryTenantIds", "ARM_AUXILIARY_TENANT_IDS")
	var auxTenants []string
	if auxTenantsString != "" {
		err := json.Unmarshal([]byte(auxTenantsString), &auxTenants)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal '%s' as Auxiliary Tenants", auxTenantsString)
		}
	}

	return &authConfiguration{
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

		showSubscription: defaultAzSubscriptionProvider,
	}, nil
}

func readAzureCloudFromConfig(getConfig configGetter) *azcloud.Configuration {
	envName := getConfig("environment", "ARM_ENVIRONMENT")
	if envName == "" {
		envName = getConfig("environment", "AZURE_ENVIRONMENT")
	}
	if envName == "" {
		// Leave the environment unset at this stage, to be resolved later in initAccount.
		return nil
	}
	cloud := azure.GetCloudByName(envName)
	return &cloud
}

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

	// https://github.com/hashicorp/terraform-provider-azurerm/blob/572bb4f37d73f4f0d914737eaca4e5a1fd084c86/internal/clients/auth.go#L33

	// Acquire an access token so we can inspect the claims
	// TODO: support other locations
	scope := "https://graph.microsoft.com/.default"
	token, err := cred.GetToken(ctx, policy.TokenRequestOptions{
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
		graphClient, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{scope})
		if err != nil {
			return nil, fmt.Errorf("failed to create Microsoft Graph client: %w", err)
		}
		if authenticatedAsServicePrincipal {
			logging.V(9).Infof("Querying Microsoft Graph to discover authenticated service principal object ID because the `oid` claim was missing from the access token")
			id, err := servicePrincipalObjectID(ctx, graphClient, clientId)
			if err != nil {
				return nil, fmt.Errorf("attempting to discover object ID for authenticated service principal with client ID %q: %w", clientId, err)
			}

			objectId = *id
		} else {
			logging.V(9).Infof("Querying Microsoft Graph to discover authenticated user principal object ID because the `oid` claim was missing from the access token")
			id, err := userPrincipalObjectID(ctx, graphClient)
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
