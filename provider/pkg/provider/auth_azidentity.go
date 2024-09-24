// Copyright 2016-2022, Pulumi Corporation.

package provider

import (
	"context"
	"crypto"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"

	hamiltonAuth "github.com/manicminer/hamilton-autorest/auth"
)

func newOidcCredential(authConf *authConfiguration) (azcore.TokenCredential, error) {
	return azidentity.NewClientAssertionCredential(
		authConf.tenantId,
		authConf.clientId,
		func(ctx context.Context) (string, error) {
			return authConf.oidcToken, nil
		},
		nil)
}

func (k *azureNativeProvider) newPulumiAuthCredential() (*azidentity.ChainedTokenCredential, error) {
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

	oidcToken string

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
	}, nil
}

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
