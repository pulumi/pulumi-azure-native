// Copyright 2016-2022, Pulumi Corporation.

package provider

import (
	"context"
	"crypto"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
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

// newTokenCredential is the main entry to the new azcore/azidentity-based authentication stack. It returns a
// TokenCredential which can be passed into various Azure Go SDKs.
func newTokenCredential(authConf *authConfiguration) (azcore.TokenCredential, error) {
	return newSingleMethodAuthCredential(authConf)
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
func newSingleMethodAuthCredential(authConf *authConfiguration) (azcore.TokenCredential, error) {
	baseClientOpts := azcore.ClientOptions{
		Cloud: authConf.cloud,
	}

	if authConf.clientCertPath != "" {
		logging.V(9).Infof("[auth] Using SP with client certificate credential")
		certs, key, err := readCertificate(authConf.clientCertPath, authConf.clientCertPassword)
		if err != nil {
			return nil, err
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
		return newOidcCredential(authConf)
	} else {
		logging.V(9).Infof("OIDC credential is not enabled, skipping")
	}

	if authConf.useMsi {
		logging.V(9).Infof("[auth] Using Managed Identity (MSI) credential")
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
	cli, err := azidentity.NewAzureCLICredential(options)
	if err == nil {
		return cli, nil
	}
	return nil, errors.Errorf("Failed to find any valid credentials")
}

// newOidcCredential creates a TokenCredential for OpenID Connect (OIDC) authentication.
// An OIDC credential is an azidentity.ClientAssertionCredential that authenticates with a token
// obtained from a callback function. The token itself can be provided in various ways:
// - directly via config/environment variable
// - from a file
// - through a token exchange by making a request to a configured endpoint
// This function configures the client assertion callback according to the above cases.
func newOidcCredential(authConf *authConfiguration) (azcore.TokenCredential, error) {
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
				ClientOptions: azcore.ClientOptions{
					Cloud: authConf.cloud,
				},
			})
	}

	return nil, errors.New("OIDC token or request URL and token are not provided")
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
	cloud azcloud.Configuration

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
}

// readAuthConfig collects auth-related configuration from Pulumi config and environment variables
func (k *azureNativeProvider) readAuthConfig() (*authConfiguration, error) {

	envName := readAzureEnvironmentFromConfig(k)
	cloud := azure.GetCloudByName(envName)

	auxTenantsString := k.getConfig("auxiliaryTenantIds", "ARM_AUXILIARY_TENANT_IDS")
	var auxTenants []string
	if auxTenantsString != "" {
		err := json.Unmarshal([]byte(auxTenantsString), &auxTenants)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal '%s' as Auxiliary Tenants", auxTenantsString)
		}
	}

	return &authConfiguration{
		clientId:   k.getConfig("clientId", "ARM_CLIENT_ID"),
		tenantId:   k.getConfig("tenantId", "ARM_TENANT_ID"),
		auxTenants: auxTenants,
		cloud:      cloud,

		subscriptionId: k.getConfig("subscriptionId", "ARM_SUBSCRIPTION_ID"),

		clientSecret:       k.getConfig("clientSecret", "ARM_CLIENT_SECRET"),
		clientCertPath:     k.getConfig("clientCertificatePath", "ARM_CLIENT_CERTIFICATE_PATH"),
		clientCertPassword: k.getConfig("clientCertificatePassword", "ARM_CLIENT_CERTIFICATE_PASSWORD"),

		useMsi: k.getConfig("useMsi", "ARM_USE_MSI") == "true",

		useOidc:               k.getConfig("useOidc", "ARM_USE_OIDC") == "true",
		oidcAudience:          k.getConfig("oidcAudience", "ARM_OIDC_AUDIENCE"),
		oidcToken:             k.getConfig("oidcToken", "ARM_OIDC_TOKEN"),
		oidcTokenFilePath:     k.getConfig("oidcTokenFilePath", "ARM_OIDC_TOKEN_FILE_PATH"),
		oidcTokenRequestToken: k.getConfig("oidcRequestToken", "ACTIONS_ID_TOKEN_REQUEST_TOKEN"),
		oidcTokenRequestUrl:   k.getConfig("oidcRequestUrl", "ACTIONS_ID_TOKEN_REQUEST_URL"),
	}, nil
}

// Claims is used to unmarshall the claims from a JWT issued by the Microsoft Identity Platform.
type Claims struct {
	Audience          string   `json:"aud"`
	Issuer            string   `json:"iss"`
	IdentityProvider  string   `json:"idp"`
	ObjectId          string   `json:"oid"`
	Roles             []string `json:"roles"`
	Scopes            string   `json:"scp"`
	Subject           string   `json:"sub"`
	TenantRegionScope string   `json:"tenant_region_scope"`
	TenantId          string   `json:"tid"`
	Version           string   `json:"ver"`

	AppDisplayName string `json:"app_displayname,omitempty"`
	AppId          string `json:"appid,omitempty"`
	IdType         string `json:"idtyp,omitempty"`
}

// parseClaims retrieves and parses the claims from a JWT issued by the Microsoft Identity Platform.
func parseClaims(token azcore.AccessToken) (claims Claims, err error) {
	jwt := strings.Split(token.Token, ".")
	payload, err := base64.RawURLEncoding.DecodeString(jwt[1])
	if err != nil {
		return
	}
	err = json.Unmarshal(payload, &claims)
	return
}

type ClientConfig struct {
	Cloud azcloud.Configuration

	ClientId       string
	ObjectId       string
	SubscriptionId string
	TenantId       string

	AuthenticatedAsAServicePrincipal bool
}

func GetClientConfig(ctx context.Context, config *authConfiguration, cred azcore.TokenCredential) (*ClientConfig, error) {

	// https://github.com/hashicorp/terraform-provider-azurerm/blob/572bb4f37d73f4f0d914737eaca4e5a1fd084c86/internal/clients/auth.go#L33

	subscriptionId := config.subscriptionId

	// Acquire an access token so we can inspect the claims
	// TODO: support other locations
	scope := "https://graph.microsoft.com/.default"
	token, err := cred.GetToken(ctx, policy.TokenRequestOptions{
		Scopes: []string{scope},
	})
	if err != nil {
		return nil, fmt.Errorf("could not acquire access token to parse claims: %+v", err)
	}

	tokenClaims, err := parseClaims(token)
	if err != nil {
		return nil, fmt.Errorf("parsing claims from access token: %+v", err)
	}

	authenticatedAsServicePrincipal := true
	if strings.Contains(strings.ToLower(tokenClaims.Scopes), "openid") {
		authenticatedAsServicePrincipal = false
	}

	clientId := tokenClaims.AppId
	if clientId == "" {
		logging.V(9).Infof("[DEBUG] Using user-supplied ClientID because the `appid` claim was missing from the access token")
		clientId = config.clientId
	}

	objectId := tokenClaims.ObjectId
	if objectId == "" {
		graphClient, err := msgraphsdk.NewGraphServiceClientWithCredentials(cred, []string{scope})
		if err != nil {
			return nil, fmt.Errorf("failed to create Microsoft Graph client: %w", err)
		}
		if authenticatedAsServicePrincipal {
			logging.V(9).Infof("[DEBUG] Querying Microsoft Graph to discover authenticated service principal object ID because the `oid` claim was missing from the access token")
			id, err := servicePrincipalObjectID(ctx, graphClient, clientId)
			if err != nil {
				return nil, fmt.Errorf("attempting to discover object ID for authenticated service principal with client ID %q: %w", clientId, err)
			}

			objectId = *id
		} else {
			logging.V(9).Infof("[DEBUG] Querying Microsoft Graph to discover authenticated user principal object ID because the `oid` claim was missing from the access token")
			id, err := userPrincipalObjectID(ctx, graphClient)
			if err != nil {
				return nil, fmt.Errorf("attempting to discover object ID for authenticated user principal: %w", err)
			}
			objectId = *id
		}
	}

	tenantId := tokenClaims.TenantId
	if tenantId == "" {
		log.Printf("[DEBUG] Using user-supplied TenantID because the `tid` claim was missing from the access token")
		tenantId = config.tenantId
	}

	account := &ClientConfig{
		Cloud: config.cloud,

		ClientId:       clientId,
		ObjectId:       objectId,
		SubscriptionId: subscriptionId,
		TenantId:       tenantId,

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
