// Copyright 2016-2024, Pulumi Corporation.

package provider

import (
	"context"
	_ "embed"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure/cloud"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed test.pfx
var testPfxCert []byte

type testProvider struct {
	config map[string]string
}

func (k *testProvider) getConfig(configName, envName string) string {
	if val, ok := k.config[configName]; ok {
		return val
	}
	return os.Getenv(envName)
}

func TestGetAuthConfig(t *testing.T) {
	getMetadata := func(ctx context.Context, endpoint string) (cloud.Configuration, error) {
		return cloud.Configuration{}, nil
	}

	setAuthEnvVariables := func(value, boolValue string) {
		if value != "" {
			t.Setenv("ARM_AUXILIARY_TENANT_IDS", `["`+value+`"]`)
		}
		t.Setenv("ARM_CLIENT_CERTIFICATE_PASSWORD", value)
		t.Setenv("ARM_CLIENT_CERTIFICATE_PATH", value)
		t.Setenv("ARM_CLIENT_ID", value)
		t.Setenv("ARM_CLIENT_SECRET", value)
		t.Setenv("ARM_OIDC_AUDIENCE", value)
		t.Setenv("ARM_OIDC_TOKEN", value)
		t.Setenv("ARM_OIDC_TOKEN_FILE_PATH", value)
		t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", value)
		t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", value)
		t.Setenv("ARM_TENANT_ID", value)
		t.Setenv("ARM_USE_MSI", boolValue)
		t.Setenv("ARM_USE_OIDC", boolValue)
	}

	t.Run("empty", func(t *testing.T) {
		setAuthEnvVariables("", "")
		p := &testProvider{}
		c, err := readAuthConfig(context.Background(), p.getConfig, getMetadata)

		require.NoError(t, err)
		require.NotNil(t, c)
		require.Empty(t, c.auxTenants)
		require.Empty(t, c.clientCertPassword)
		require.Empty(t, c.clientCertPath)
		require.Empty(t, c.clientId)
		require.Empty(t, c.clientSecret)
		require.Empty(t, c.oidcAudience)
		require.Empty(t, c.oidcToken)
		require.Empty(t, c.oidcTokenFilePath)
		require.Empty(t, c.oidcTokenRequestToken)
		require.Empty(t, c.oidcTokenRequestUrl)
		require.Empty(t, c.tenantId)
		require.False(t, c.useOidc)
		require.False(t, c.useMsi)
		require.Nil(t, c.cloud)
	})

	t.Run("values from config take precedence", func(t *testing.T) {
		setAuthEnvVariables("env", "false")
		t.Setenv("ARM_ENVIRONMENT", "public")

		p := &testProvider{
			config: map[string]string{
				"auxiliaryTenantIds":        `["conf"]`,
				"clientCertificatePassword": "conf",
				"clientCertificatePath":     "conf",
				"clientId":                  "conf",
				"clientSecret":              "conf",
				"environment":               "usgov",
				"oidcAudience":              "conf",
				"oidcToken":                 "conf",
				"oidcTokenFilePath":         "conf",
				"oidcRequestToken":          "conf",
				"oidcRequestUrl":            "conf",
				"tenantId":                  "conf",
				"useMsi":                    "true",
				"useOidc":                   "true",
			},
		}

		c, err := readAuthConfig(context.Background(), p.getConfig, getMetadata)
		require.NoError(t, err)
		require.NotNil(t, c)
		require.Equal(t, []string{"conf"}, c.auxTenants)
		require.Equal(t, "conf", c.clientCertPassword)
		require.Equal(t, "conf", c.clientCertPath)
		require.Equal(t, "conf", c.clientId)
		require.Equal(t, "conf", c.clientSecret)
		require.Equal(t, "conf", c.oidcAudience)
		require.Equal(t, "conf", c.oidcToken)
		require.Equal(t, "conf", c.oidcTokenFilePath)
		require.Equal(t, "conf", c.oidcTokenRequestToken)
		require.Equal(t, "conf", c.oidcTokenRequestUrl)
		require.Equal(t, "conf", c.tenantId)
		require.NotNil(t, c.cloud)
		require.Equal(t, "https://login.microsoftonline.us/", c.cloud.ActiveDirectoryAuthorityHost)
		require.True(t, c.useOidc)
		require.True(t, c.useMsi)
	})

	t.Run("values from env", func(t *testing.T) {
		p := &testProvider{}
		setAuthEnvVariables("env", "true")
		t.Setenv("ARM_ENVIRONMENT", "usgovernment")

		c, err := readAuthConfig(context.Background(), p.getConfig, getMetadata)
		require.NoError(t, err)
		require.NotNil(t, c)
		require.Equal(t, []string{"env"}, c.auxTenants)
		require.Equal(t, "env", c.clientCertPassword)
		require.Equal(t, "env", c.clientCertPath)
		require.Equal(t, "env", c.clientId)
		require.Equal(t, "env", c.clientSecret)
		require.Equal(t, "env", c.oidcAudience)
		require.Equal(t, "env", c.oidcToken)
		require.Equal(t, "env", c.oidcTokenFilePath)
		require.Equal(t, "env", c.oidcTokenRequestToken)
		require.Equal(t, "env", c.oidcTokenRequestUrl)
		require.Equal(t, "env", c.tenantId)
		require.NotNil(t, c.cloud)
		require.Equal(t, "https://login.microsoftonline.us/", c.cloud.ActiveDirectoryAuthorityHost)
		require.True(t, c.useOidc)
		require.True(t, c.useMsi)
	})
}

func TestNewCredential(t *testing.T) {
	t.Run("SP with client secret", func(t *testing.T) {
		conf := &authConfiguration{
			clientId:       "client-id",
			clientSecret:   "client-secret",
			tenantId:       "tenant-id",
			subscriptionId: "subscription-id",
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientSecretCredential{}, cred)
		clientVal := reflect.ValueOf(cred).Elem().FieldByName("client")
		require.Equal(t, "client-id", clientVal.Elem().FieldByName("clientID").String())
		require.Equal(t, "tenant-id", clientVal.Elem().FieldByName("tenantID").String())
	})

	t.Run("Incomplete SP missing subscription id", func(t *testing.T) {
		conf := &authConfiguration{
			clientId:     "client-id",
			clientSecret: "client-secret",
			tenantId:     "tenant-id",
		}
		_, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.Error(t, err)
		require.Contains(t, err.Error(), "Subscription")
	})

	t.Run("Incomplete SP with client secret conf missing tenant id", func(t *testing.T) {
		conf := &authConfiguration{
			clientId:       "client-id",
			clientSecret:   "client-secret",
			subscriptionId: "subscription-id",
		}
		_, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.Error(t, err)
		require.Contains(t, err.Error(), "Tenant")
	})

	t.Run("SP with client cert", func(t *testing.T) {
		certPath := filepath.Join(t.TempDir(), "cert.pfx")
		require.NoError(t, os.WriteFile(certPath, testPfxCert, 0644))

		conf := &authConfiguration{
			clientId:           "client-id",
			clientCertPath:     certPath,
			clientCertPassword: "pulumi",
			tenantId:           "tenant-id",
			subscriptionId:     "subscription-id",
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientCertificateCredential{}, cred)
		clientVal := reflect.ValueOf(cred).Elem().FieldByName("client")
		require.Equal(t, "client-id", clientVal.Elem().FieldByName("clientID").String())
		require.Equal(t, "tenant-id", clientVal.Elem().FieldByName("tenantID").String())
	})

	t.Run("SP with invalid client cert", func(t *testing.T) {
		certPath := filepath.Join(t.TempDir(), "cert.pem")
		require.NoError(t, os.WriteFile(certPath, []byte("cert"), 0644))

		conf := &authConfiguration{
			clientId:       "client-id",
			clientCertPath: certPath,
			tenantId:       "tenant-id",
			subscriptionId: "subscription-id",
		}
		_, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
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
			subscriptionId:     "subscription-id",
		}
		_, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to parse certificate")
		require.Contains(t, err.Error(), "password incorrect")
	})

	t.Run("OIDC with token", func(t *testing.T) {
		conf := &authConfiguration{
			useOidc:        true,
			oidcToken:      "oidc-token",
			clientId:       "client-id",
			tenantId:       "tenant-id",
			subscriptionId: "subscription-id",
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientAssertionCredential{}, cred)
		clientVal := reflect.ValueOf(cred).Elem().FieldByName("client")
		require.Equal(t, "client-id", clientVal.Elem().FieldByName("clientID").String())
		require.Equal(t, "tenant-id", clientVal.Elem().FieldByName("tenantID").String())
	})

	t.Run("OIDC with token file", func(t *testing.T) {
		tokenPath := filepath.Join(t.TempDir(), "my.token")
		require.NoError(t, os.WriteFile(tokenPath, []byte("token"), 0644))

		conf := &authConfiguration{
			useOidc:           true,
			oidcTokenFilePath: tokenPath,
			clientId:          "client-id",
			tenantId:          "tenant-id",
			subscriptionId:    "subscription-id",
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientAssertionCredential{}, cred)
	})

	t.Run("OIDC with wrong token file", func(t *testing.T) {
		conf := &authConfiguration{
			useOidc:           true,
			oidcTokenFilePath: filepath.Join(t.TempDir(), "foo"),
			clientId:          "client-id",
			tenantId:          "tenant-id",
			subscriptionId:    "subscription-id",
		}
		_, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.Error(t, err)
		require.ErrorIs(t, err, os.ErrNotExist)
	})

	t.Run("OIDC with token exchange URL", func(t *testing.T) {
		conf := &authConfiguration{
			useOidc:               true,
			oidcTokenRequestToken: "oidc-token",
			oidcTokenRequestUrl:   "oidc-token-url",
			clientId:              "client-id",
			tenantId:              "tenant-id",
			subscriptionId:        "subscription-id",
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientAssertionCredential{}, cred)
	})

	t.Run("OIDC with token exchange URL and custom audience", func(t *testing.T) {
		conf := &authConfiguration{
			useOidc:               true,
			oidcAudience:          "api://oidc-audience",
			oidcTokenRequestToken: "oidc-token",
			oidcTokenRequestUrl:   "oidc-token-url",
			clientId:              "client-id",
			tenantId:              "tenant-id",
			subscriptionId:        "subscription-id",
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.ClientAssertionCredential{}, cred)
	})

	t.Run("Incomplete OIDC conf", func(t *testing.T) {
		for _, conf := range []*authConfiguration{
			{ // missing tenantId
				useOidc:        true,
				oidcToken:      "oidc-token",
				clientId:       "client-id",
				subscriptionId: "subscription-id",
			},
			{ // missing oidcTokenRequestToken
				useOidc:             true,
				oidcTokenRequestUrl: "oidc-token-url",
				clientId:            "client-id",
				tenantId:            "tenant-id",
				subscriptionId:      "subscription-id",
			},
			{ // missing subscriptionId
				useOidc:   true,
				oidcToken: "oidc-token",
				clientId:  "client-id",
				tenantId:  "tenant-id",
			},
		} {
			_, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
			require.Error(t, err)
		}
	})

	t.Run("MSI", func(t *testing.T) {
		conf := &authConfiguration{
			useMsi:         true,
			subscriptionId: "subscription-id",
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.ManagedIdentityCredential{}, cred)
	})

	// Used for user-assigned managed identity
	t.Run("MSI with client id", func(t *testing.T) {
		conf := &authConfiguration{
			useMsi:         true,
			clientId:       "123",
			subscriptionId: "subscription-id",
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.ManagedIdentityCredential{}, cred)
		// FUTURE: assert cred.client.id = "123"
	})

	t.Run("CLI", func(t *testing.T) {
		conf := &authConfiguration{}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.AzureCLICredential{}, cred)
	})

	t.Run("CLI with aux tenants", func(t *testing.T) {
		conf := &authConfiguration{
			auxTenants: []string{"123", "456"},
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.AzureCLICredential{}, cred)
		optsVal := reflect.ValueOf(cred).Elem().FieldByName("opts")
		require.Equal(t, "123", optsVal.FieldByName("AdditionallyAllowedTenants").Index(0).String())
		require.Equal(t, "456", optsVal.FieldByName("AdditionallyAllowedTenants").Index(1).String())
	})

	t.Run("CLI with subscription id", func(t *testing.T) {
		conf := &authConfiguration{
			subscriptionId: "subscription-id",
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.AzureCLICredential{}, cred)
		optsVal := reflect.ValueOf(cred).Elem().FieldByName("opts")
		require.Equal(t, "subscription-id", optsVal.FieldByName("Subscription").String())
	})

	t.Run("Azure Default Credential", func(t *testing.T) {
		conf := &authConfiguration{
			useDefault:     true,
			subscriptionId: "subscription-id",
		}
		cred, err := newSingleMethodAuthCredential(conf, azcore.ClientOptions{})
		require.NoError(t, err)
		require.IsType(t, &azidentity.DefaultAzureCredential{}, cred)
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

var testEnvironment = cloud.Configuration{
	Name: "Test",
	Configuration: azcloud.Configuration{
		ActiveDirectoryAuthorityHost: "https://login.test/",
		Services: map[azcloud.ServiceName]azcloud.ServiceConfiguration{
			azcloud.ResourceManager: {
				Audience: "https://management.core.test/",
				Endpoint: "https://management.test/",
			},
		},
	},
	Endpoints: cloud.ConfigurationEndpoints{
		MicrosoftGraph: "https://microsoftgraph.test/",
	},
	Suffixes: cloud.ConfigurationSuffixes{
		StorageEndpoint: "core.storage.test",
		KeyVaultDNS:     ".vault.test",
	},
}

func TestReadCloudConfiguration(t *testing.T) {
	type testCase struct {
		name      string
		config    map[string]string
		env       map[string]string
		metadata  cloud.Configuration
		expectErr bool
		expect    *cloud.Configuration
	}

	tests := []testCase{
		{
			name:      "no config, no env returns nil",
			config:    map[string]string{},
			env:       map[string]string{},
			expectErr: false,
			expect:    nil,
		},
		{
			name: "well-known cloud by config",
			config: map[string]string{
				"environment": "public",
			},
			env:       map[string]string{},
			expectErr: false,
			expect:    &cloud.AzurePublic,
		},
		{
			name:   "well-known cloud by ARM_ENVIRONMENT",
			config: map[string]string{},
			env: map[string]string{
				"ARM_ENVIRONMENT": "usgovernment",
			},
			expectErr: false,
			expect:    &cloud.AzureGovernment,
		},
		{
			name:   "well-known cloud by AZURE_ENVIRONMENT",
			config: map[string]string{},
			env: map[string]string{
				"AZURE_ENVIRONMENT": "china",
			},
			expectErr: false,
			expect:    &cloud.AzureChina,
		},
		{
			name: "from metadata server",
			env: map[string]string{
				"ARM_METADATA_HOSTNAME": "metadata.test",
			},
			metadata:  testEnvironment,
			expectErr: false,
			expect:    &testEnvironment,
		},
		{
			name: "from metadata server (overriding ARM_ENVIRONMENT)",
			env: map[string]string{
				"ARM_METADATA_HOSTNAME": "metadata.test",
				"ARM_ENVIRONMENT":       "usgovernment",
			},
			metadata:  testEnvironment,
			expectErr: false,
			expect:    &testEnvironment,
		},
		{
			name: "unknown environment with no metadata server returns error",
			config: map[string]string{
				"environment": "unknowncloud",
			},
			env:       map[string]string{},
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.env {
				t.Setenv(k, v)
			}
			getMetadata := func(ctx context.Context, endpoint string) (cloud.Configuration, error) {
				return tc.metadata, nil
			}
			getConfig := func(configName, envName string) string {
				if val, ok := tc.config[configName]; ok {
					return val
				}
				return os.Getenv(envName)
			}
			cloudConf, err := readCloudConfiguration(context.Background(), getConfig, getMetadata)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tc.expect == nil {
					require.Nil(t, cloudConf)
				} else {
					require.NotNil(t, cloudConf)
					require.Equal(t, tc.expect.Name, cloudConf.Name)
					require.Equal(t, tc.expect.Configuration.ActiveDirectoryAuthorityHost, cloudConf.Configuration.ActiveDirectoryAuthorityHost)
					require.Equal(t,
						tc.expect.Configuration.Services[azcloud.ResourceManager],
						cloudConf.Configuration.Services[azcloud.ResourceManager],
					)
					require.Equal(t, tc.expect.Endpoints.MicrosoftGraph, cloudConf.Endpoints.MicrosoftGraph)
					require.Equal(t, tc.expect.Suffixes.StorageEndpoint, cloudConf.Suffixes.StorageEndpoint)
					require.Equal(t, tc.expect.Suffixes.KeyVaultDNS, cloudConf.Suffixes.KeyVaultDNS)
				}
			}
		})
	}
}
