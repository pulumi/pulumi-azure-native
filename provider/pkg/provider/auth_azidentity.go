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
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

func newOidcCredential(authConf *authConfiguration) (azcore.TokenCredential, error) {
	if authConf.oidcToken != "" {
		return azidentity.NewClientAssertionCredential(
			authConf.tenantId,
			authConf.clientId,
			func(ctx context.Context) (string, error) {
				return authConf.oidcToken, nil
			},
			nil)
	}

	if authConf.oidcTokenRequestUrl != "" && authConf.oidcTokenRequestToken != "" {
		return azidentity.NewClientAssertionCredential(
			authConf.tenantId,
			authConf.clientId,
			func(ctx context.Context) (string, error) {
				req, err := http.NewRequestWithContext(ctx, http.MethodGet, authConf.oidcTokenRequestUrl, http.NoBody)
				if err != nil {
					return "", fmt.Errorf("GitHub OIDC: failed to build request to %s: %w", authConf.oidcTokenRequestUrl, err)
				}

				query, err := url.ParseQuery(req.URL.RawQuery)
				if err != nil {
					return "", fmt.Errorf("githubAssertion: cannot parse URL query")
				}
				query.Set("audience", "api://AzureADTokenExchange")
				req.URL.RawQuery = query.Encode()

				req.Header.Set("Accept", "application/json")
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authConf.oidcTokenRequestToken))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					return "", fmt.Errorf("GitHub OIDC: couldn't request token: %w", err)
				}

				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return "", fmt.Errorf("GitHub OIDC: cannot parse response: %w", err)
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
			},
			nil)
	}

	return nil, errors.New("OIDC token or request URL and token are not provided")
}

func (k *azureNativeProvider) newChainedAuthCredential() (*azidentity.ChainedTokenCredential, error) {
	authConf, err := k.getAuthConfig3()
	if err != nil {
		return nil, err
	}

	var credentials []azcore.TokenCredential
	var errs []error
	add := func(cred azcore.TokenCredential, name string, err error) ([]azcore.TokenCredential, []error) {
		if err != nil {
			errs = append(errs, err)
			logging.V(9).Infof("%s credential cannot be used: %v", name, err)
		} else {
			credentials = append(credentials, cred)
			logging.V(9).Infof("%s credential is valid and will be considered", name)
		}
		return credentials, errs
	}

	spCert, err := azidentity.NewClientSecretCredential(authConf.tenantId, authConf.clientId, authConf.clientSecret, nil)
	credentials, errs = add(spCert, "SP with client secret", err)

	oidc, err := newOidcCredential(authConf)
	credentials, errs = add(oidc, "OIDC", err)

	// TODO,tkappler using it as-is in the chain causes a hard failure when the MSI endpoint is not available
	// Need to wrap with a timeout handler as shown here:
	// https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity#example-NewChainedTokenCredential-ManagedIdentityTimeout
	// managedIdentity, err := azidentity.NewManagedIdentityCredential(nil)
	// credentials, errs = add(managedIdentity, "Managed Identity", err)

	cli, err := azidentity.NewAzureCLICredential(nil)
	credentials, errs = add(cli, "Azure CLI", err)

	if len(credentials) == 0 {
		return nil, errors.Errorf("Failed to find valid credentials: %v", errs)
	}

	return azidentity.NewChainedTokenCredential(credentials, nil)
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

func (k *azureNativeProvider) newSingleMethodAuthCredential() (azcore.TokenCredential, error) {
	authConf, err := k.getAuthConfig3()
	if err != nil {
		return nil, err
	}

	if authConf.clientSecret != "" {
		spCert, err := azidentity.NewClientSecretCredential(authConf.tenantId, authConf.clientId, authConf.clientSecret, nil)
		if err == nil {
			return spCert, nil
		}
		logging.V(9).Infof("SP with client secret credential cannot be used: %v", err)
	} else {
		logging.V(9).Infof("SP with client secret credential is not enabled, skipping")
	}

	if authConf.clientCertPath != "" {
		certs, key, err := readCertificate(authConf.clientCertPath, authConf.clientCertPassword)
		if err != nil {
			return nil, err
		}

		spCert, err := azidentity.NewClientCertificateCredential(authConf.tenantId, authConf.clientId, certs, key, nil)
		if err == nil {
			return spCert, nil
		}
		logging.V(9).Infof("SP with client certificate credential cannot be used: %v", err)
	} else {
		logging.V(9).Infof("SP with client certificate credential is not enabled, skipping")
	}

	if k.getConfig("useOidc", "ARM_USE_OIDC") == "true" {
		oidc, err := newOidcCredential(authConf)
		if err == nil {
			return oidc, nil
		}
		logging.V(9).Infof("OIDC credential cannot be used: %v", err)
	} else {
		logging.V(9).Infof("OIDC credential is not enabled, skipping")
	}

	if k.getConfig("useMsi", "ARM_USE_MSI") == "true" {
		managedIdentity, err := azidentity.NewManagedIdentityCredential(nil)
		if err == nil {
			return managedIdentity, nil
		}
		logging.V(9).Infof("Managed Identity (MSI) credential cannot be used: %v", err)
	} else {
		logging.V(9).Infof("Managed Identity (MSI) credential is not enabled, skipping")
	}

	cli, err := azidentity.NewAzureCLICredential(nil)
	if err == nil {
		return cli, nil
	}
	logging.V(9).Infof("Azure CLI credential cannot be used: %v", err)

	return nil, errors.Errorf("Failed to find any valid credentials")
}

type authConfiguration struct {
	subscriptionId string
	cloud          string

	clientId   string
	tenantId   string
	auxTenants []string

	oidcToken             string
	oidcTokenRequestToken string
	oidcTokenRequestUrl   string

	clientSecret       string
	clientCertPath     string
	clientCertPassword string
}

func (k *azureNativeProvider) getAuthConfig3() (*authConfiguration, error) {
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

	return &authConfiguration{
		subscriptionId:     k.getConfig("subscriptionId", "ARM_SUBSCRIPTION_ID"),
		clientId:           k.getConfig("clientId", "ARM_CLIENT_ID"),
		tenantId:           k.getConfig("tenantId", "ARM_TENANT_ID"),
		auxTenants:         auxTenants,
		cloud:              envName,
		clientSecret:       k.getConfig("clientSecret", "ARM_CLIENT_SECRET"),
		clientCertPath:     k.getConfig("clientCertificatePath", "ARM_CLIENT_CERTIFICATE_PATH"),
		clientCertPassword: k.getConfig("clientCertificatePassword", "ARM_CLIENT_CERTIFICATE_PASSWORD"),

		oidcToken:             k.getConfig("oidcToken", "ARM_OIDC_TOKEN"),
		oidcTokenRequestToken: k.getConfig("oidcRequestToken", "ACTIONS_ID_TOKEN_REQUEST_TOKEN"),
		oidcTokenRequestUrl:   k.getConfig("oidcRequestUrl", "ACTIONS_ID_TOKEN_REQUEST_URL"),
	}, nil
}
