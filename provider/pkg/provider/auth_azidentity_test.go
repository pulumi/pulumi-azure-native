// Copyright 2016-2024, Pulumi Corporation.

package provider

import (
	"context"
	_ "embed"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed test.pfx
var testPfxCert []byte

func TestGetAuthConfig(t *testing.T) {
	setAuthEnvVariables := func(value, boolValue, cloudValue string) {
		if value != "" {
			t.Setenv("ARM_AUXILIARY_TENANT_IDS", `["`+value+`"]`)
		}
		t.Setenv("ARM_CLIENT_CERTIFICATE_PASSWORD", value)
		t.Setenv("ARM_CLIENT_CERTIFICATE_PATH", value)
		t.Setenv("ARM_CLIENT_ID", value)
		t.Setenv("ARM_CLIENT_SECRET", value)
		t.Setenv("ARM_ENVIRONMENT", cloudValue)
		t.Setenv("ARM_OIDC_TOKEN", value)
		t.Setenv("ARM_OIDC_TOKEN_FILE_PATH", value)
		t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", value)
		t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", value)
		t.Setenv("ARM_SUBSCRIPTION_ID", value)
		t.Setenv("ARM_TENANT_ID", value)
		t.Setenv("ARM_USE_MSI", boolValue)
		t.Setenv("ARM_USE_OIDC", boolValue)
	}

	t.Run("empty", func(t *testing.T) {
		setAuthEnvVariables("", "", "")
		p := azureNativeProvider{}
		c, err := p.readAuthConfig()
		require.NoError(t, err)
		require.NotNil(t, c)
		require.Empty(t, c.auxTenants)
		require.Empty(t, c.clientCertPassword)
		require.Empty(t, c.clientCertPath)
		require.Empty(t, c.clientId)
		require.Empty(t, c.clientSecret)
		require.Equal(t, cloud.AzurePublic, c.cloud)
		require.Empty(t, c.oidcToken)
		require.Empty(t, c.oidcTokenFilePath)
		require.Empty(t, c.oidcTokenRequestToken)
		require.Empty(t, c.oidcTokenRequestUrl)
		require.Empty(t, c.subscriptionId)
		require.Empty(t, c.tenantId)
		require.False(t, c.useOidc)
		require.False(t, c.useMsi)
	})

	t.Run("values from config take precedence", func(t *testing.T) {
		setAuthEnvVariables("env", "false", "china")

		p := azureNativeProvider{
			config: map[string]string{
				"auxiliaryTenantIds":        `["conf"]`,
				"clientCertificatePassword": "conf",
				"clientCertificatePath":     "conf",
				"clientId":                  "conf",
				"clientSecret":              "conf",
				"environment":               "usgov",
				"oidcToken":                 "conf",
				"oidcTokenFilePath":         "conf",
				"oidcRequestToken":          "conf",
				"oidcRequestUrl":            "conf",
				"subscriptionId":            "conf",
				"tenantId":                  "conf",
				"useMsi":                    "true",
				"useOidc":                   "true",
			},
		}

		c, err := p.readAuthConfig()
		require.NoError(t, err)
		require.NotNil(t, c)
		require.Equal(t, []string{"conf"}, c.auxTenants)
		require.Equal(t, "conf", c.clientCertPassword)
		require.Equal(t, "conf", c.clientCertPath)
		require.Equal(t, "conf", c.clientId)
		require.Equal(t, "conf", c.clientSecret)
		require.Equal(t, cloud.AzureGovernment, c.cloud)
		require.Equal(t, "conf", c.oidcToken)
		require.Equal(t, "conf", c.oidcTokenFilePath)
		require.Equal(t, "conf", c.oidcTokenRequestToken)
		require.Equal(t, "conf", c.oidcTokenRequestUrl)
		require.Equal(t, "conf", c.subscriptionId)
		require.Equal(t, "conf", c.tenantId)
		require.True(t, c.useOidc)
		require.True(t, c.useMsi)
	})

	t.Run("values from env", func(t *testing.T) {
		p := azureNativeProvider{}
		setAuthEnvVariables("env", "true", "china")

		c, err := p.readAuthConfig()
		require.NoError(t, err)
		require.NotNil(t, c)
		require.Equal(t, []string{"env"}, c.auxTenants)
		require.Equal(t, "env", c.clientCertPassword)
		require.Equal(t, "env", c.clientCertPath)
		require.Equal(t, "env", c.clientId)
		require.Equal(t, "env", c.clientSecret)
		require.Equal(t, cloud.AzureChina, c.cloud)
		require.Equal(t, "env", c.oidcToken)
		require.Equal(t, "env", c.oidcTokenFilePath)
		require.Equal(t, "env", c.oidcTokenRequestToken)
		require.Equal(t, "env", c.oidcTokenRequestUrl)
		require.Equal(t, "env", c.subscriptionId)
		require.Equal(t, "env", c.tenantId)
		require.True(t, c.useOidc)
		require.True(t, c.useMsi)
	})
}

func TestNewCredential(t *testing.T) {
	t.Run("SP with client secret", func(t *testing.T) {
		conf := &authConfiguration{
			clientId:     "client-id",
			clientSecret: "client-secret",
			tenantId:     "tenant-id",
		}
		cred, err := newSingleMethodAuthCredential(conf)
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientSecretCredential{}, cred)
	})

	t.Run("Incomplete SP with client secret conf missing tenant id", func(t *testing.T) {
		conf := &authConfiguration{
			clientId:     "client-id",
			clientSecret: "client-secret",
		}
		_, err := newSingleMethodAuthCredential(conf)
		require.Error(t, err)
		require.Contains(t, err.Error(), "tenant")
	})

	t.Run("SP with client cert", func(t *testing.T) {
		certPath := filepath.Join(t.TempDir(), "cert.pfx")
		require.NoError(t, os.WriteFile(certPath, testPfxCert, 0644))

		conf := &authConfiguration{
			clientId:           "client-id",
			clientCertPath:     certPath,
			clientCertPassword: "pulumi",
			tenantId:           "tenant-id",
		}
		cred, err := newSingleMethodAuthCredential(conf)
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientCertificateCredential{}, cred)
	})

	t.Run("SP with invalid client cert", func(t *testing.T) {
		certPath := filepath.Join(t.TempDir(), "cert.pem")
		require.NoError(t, os.WriteFile(certPath, []byte("cert"), 0644))

		conf := &authConfiguration{
			clientId:       "client-id",
			clientCertPath: certPath,
			tenantId:       "tenant-id",
		}
		_, err := newSingleMethodAuthCredential(conf)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to parse certificate")
	})

	t.Run("SP with client cert and wrong password", func(t *testing.T) {
		certPath := filepath.Join(t.TempDir(), "cert.pfx")
		require.NoError(t, os.WriteFile(certPath, testPfxCert, 0644))

		conf := &authConfiguration{
			clientId:           "client-id",
			clientCertPath:     certPath,
			clientCertPassword: "wrong",
			tenantId:           "tenant-id",
		}
		_, err := newSingleMethodAuthCredential(conf)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to parse certificate")
		require.Contains(t, err.Error(), "password incorrect")
	})

	t.Run("OIDC with token", func(t *testing.T) {
		conf := &authConfiguration{
			useOidc:   true,
			oidcToken: "oidc-token",
			clientId:  "client-id",
			tenantId:  "tenant-id",
		}
		cred, err := newSingleMethodAuthCredential(conf)
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientAssertionCredential{}, cred)
	})

	t.Run("OIDC with token file", func(t *testing.T) {
		tokenPath := filepath.Join(t.TempDir(), "my.token")
		require.NoError(t, os.WriteFile(tokenPath, []byte("token"), 0644))

		conf := &authConfiguration{
			useOidc:           true,
			oidcTokenFilePath: tokenPath,
			clientId:          "client-id",
			tenantId:          "tenant-id",
		}
		cred, err := newSingleMethodAuthCredential(conf)
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientAssertionCredential{}, cred)
	})

	t.Run("OIDC with wrong token file", func(t *testing.T) {
		conf := &authConfiguration{
			useOidc:           true,
			oidcTokenFilePath: filepath.Join(t.TempDir(), "foo"),
			clientId:          "client-id",
			tenantId:          "tenant-id",
		}
		_, err := newSingleMethodAuthCredential(conf)
		require.Error(t, err)
	})

	t.Run("OIDC with token exchange URL", func(t *testing.T) {
		conf := &authConfiguration{
			useOidc:               true,
			oidcTokenRequestToken: "oidc-token",
			oidcTokenRequestUrl:   "oidc-token-url",
			clientId:              "client-id",
			tenantId:              "tenant-id",
		}
		cred, err := newSingleMethodAuthCredential(conf)
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientAssertionCredential{}, cred)
	})

	t.Run("Incomplete OIDC conf", func(t *testing.T) {
		for _, conf := range []*authConfiguration{
			{
				useOidc:   true,
				oidcToken: "oidc-token",
				clientId:  "client-id",
			},
			{
				useOidc:             true,
				oidcTokenRequestUrl: "oidc-token-url",
				clientId:            "client-id",
				tenantId:            "tenant-id",
			},
		} {
			_, err := newSingleMethodAuthCredential(conf)
			require.Error(t, err)
		}
	})

	t.Run("MSI", func(t *testing.T) {
		conf := &authConfiguration{
			useMsi: true,
		}
		cred, err := newSingleMethodAuthCredential(conf)
		require.NoError(t, err)
		require.IsType(t, &azidentity.ManagedIdentityCredential{}, cred)
	})

	// Used for user-assigned managed identity
	t.Run("MSI with client id", func(t *testing.T) {
		conf := &authConfiguration{
			clientId: "123",
			useMsi:   true,
		}
		cred, err := newSingleMethodAuthCredential(conf)
		require.NoError(t, err)
		require.IsType(t, &azidentity.ManagedIdentityCredential{}, cred)
	})

	t.Run("CLI", func(t *testing.T) {
		conf := &authConfiguration{}
		cred, err := newSingleMethodAuthCredential(conf)
		require.NoError(t, err)
		require.IsType(t, &azidentity.AzureCLICredential{}, cred)
	})

	t.Run("CLI with tenant ids", func(t *testing.T) {
		conf := &authConfiguration{
			tenantId:   "tenant-id",
			auxTenants: []string{"123", "456"},
		}
		cred, err := newSingleMethodAuthCredential(conf)
		require.NoError(t, err)
		require.IsType(t, &azidentity.AzureCLICredential{}, cred)
	})
}

func TestOidcTokenExchangeAssertion(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		assert.Equal(t, "application/x-www-form-urlencoded", r.Header.Get("Content-Type"))
		assert.Equal(t, "application/json", r.Header.Get("Accept"))
		assert.Equal(t, "api://AzureADTokenExchange", r.URL.Query().Get("audience"))
		assert.Equal(t, "Bearer oidc-token", r.Header.Get("Authorization"))

		w.Write([]byte(`{"Value": "new-oidc-token"}`))
	}))
	defer ts.Close()

	conf := &authConfiguration{
		oidcTokenRequestToken: "oidc-token",
		oidcTokenRequestUrl:   ts.URL,
	}

	assertion := getOidcTokenExchangeAssertion(conf)

	oidcToken, err := assertion(context.Background())
	require.NoError(t, err)
	require.Equal(t, "new-oidc-token", oidcToken)
}
