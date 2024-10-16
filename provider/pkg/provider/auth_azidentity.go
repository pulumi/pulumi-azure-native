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
	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

// newTokenCredential is the main entry to the new azcore/azidentity-based authenticattion stack. It returns a
// TokenCredential which can be passed into various Azure Go SDKs.
func (k *azureNativeProvider) newTokenCredential() (azcore.TokenCredential, error) {
	authConf, err := k.readAuthConfig()
	if err != nil {
		return nil, err
	}

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
		query.Set("audience", "api://AzureADTokenExchange")
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
	if err != nil || len(certs) == 0 {
		return nil, nil, errors.Wrapf(err, "failed to parse certificate from %s", certPath)
	}

	return certs, key, nil
}

type authConfiguration struct {
	subscriptionId string
	cloud          azcloud.Configuration

	clientId   string
	tenantId   string
	auxTenants []string

	useOidc               bool
	oidcToken             string
	oidcTokenFilePath     string
	oidcTokenRequestToken string
	oidcTokenRequestUrl   string

	clientSecret       string
	clientCertPath     string
	clientCertPassword string

	// Note: for MSI, there used to be a "msiEndpoint"/"ARM_MSI_ENDPOINT" configuration. The newer
	// azidentity package handles the MSI endpoint automatically:
	// https://github.com/Azure/azure-sdk-for-go/blob/sdk/azidentity/v1.8.0/sdk/azidentity/managed_identity_client.go#L143
	useMsi bool
}

// getAuthConfig collects auth-related configuration from Pulumi config and environment variables
func (k *azureNativeProvider) readAuthConfig() (*authConfiguration, error) {
	auxTenantsString := k.getConfig("auxiliaryTenantIds", "ARM_AUXILIARY_TENANT_IDS")
	var auxTenants []string
	if auxTenantsString != "" {
		err := json.Unmarshal([]byte(auxTenantsString), &auxTenants)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to unmarshal '%s' as Auxiliary Tenants", auxTenantsString)
		}
	}

	cloudName := k.getConfig("environment", "ARM_ENVIRONMENT")
	if cloudName == "" {
		cloudName = "public"
	}
	cloud := azure.GetCloudByName(cloudName)

	return &authConfiguration{
		subscriptionId: k.getConfig("subscriptionId", "ARM_SUBSCRIPTION_ID"),
		clientId:       k.getConfig("clientId", "ARM_CLIENT_ID"),
		tenantId:       k.getConfig("tenantId", "ARM_TENANT_ID"),
		auxTenants:     auxTenants,
		cloud:          cloud,

		clientSecret:       k.getConfig("clientSecret", "ARM_CLIENT_SECRET"),
		clientCertPath:     k.getConfig("clientCertificatePath", "ARM_CLIENT_CERTIFICATE_PATH"),
		clientCertPassword: k.getConfig("clientCertificatePassword", "ARM_CLIENT_CERTIFICATE_PASSWORD"),

		useMsi: k.getConfig("useMsi", "ARM_USE_MSI") == "true",

		useOidc:               k.getConfig("useOidc", "ARM_USE_OIDC") == "true",
		oidcToken:             k.getConfig("oidcToken", "ARM_OIDC_TOKEN"),
		oidcTokenFilePath:     k.getConfig("oidcTokenFilePath", "ARM_OIDC_TOKEN_FILE_PATH"),
		oidcTokenRequestToken: k.getConfig("oidcRequestToken", "ACTIONS_ID_TOKEN_REQUEST_TOKEN"),
		oidcTokenRequestUrl:   k.getConfig("oidcRequestUrl", "ACTIONS_ID_TOKEN_REQUEST_URL"),
	}, nil
}
