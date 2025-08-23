// Copyright 2022, Pulumi Corporation.

package provider

import (
	"os"
	"testing"

	azcloud "github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/hashicorp/go-azure-helpers/authentication"
	goversion "github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptedAzVersions(t *testing.T) {
	goodVersions := []string{
		"2.37", "2.37.1", "2.38", "2.40", "2.99.99",
	}
	badVersions := []string{
		"1",
		"2", "2.0.80", "2.0.81", "2.1.81", "2.20", "2.33", "2.36.6",
		"3", "3.0.1", "3.5",
		"4",
	}

	for _, good := range goodVersions {
		v := goversion.Must(goversion.NewVersion(good))
		assert.Nil(t, assertAzVersion(v))
	}

	for _, bad := range badVersions {
		v := goversion.Must(goversion.NewVersion(bad))
		assert.NotNil(t, assertAzVersion(v))
	}
}

func TestOidcConfigFromGithub(t *testing.T) {
	p := azureNativeProvider{}

	// GH pre-set env
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "t1")
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", "url1")

	conf, err := p.determineOidcConfig()
	assert.Nil(t, err)
	assert.Equal(t, "t1", conf.oidcRequestToken)
	assert.Equal(t, "url1", conf.oidcRequestUrl)
	assert.Empty(t, conf.oidcToken)
}

func TestOidcConfigWithRequestUrlFromEnv(t *testing.T) {
	p := azureNativeProvider{}

	t.Setenv("ARM_OIDC_REQUEST_TOKEN", "t1")
	t.Setenv("ARM_OIDC_REQUEST_URL", "url1")

	conf, err := p.determineOidcConfig()
	assert.Nil(t, err)
	assert.Equal(t, "t1", conf.oidcRequestToken)
	assert.Equal(t, "url1", conf.oidcRequestUrl)
	assert.Empty(t, conf.oidcToken)
}

func TestOidcConfigWithRequestUrlFromConfig(t *testing.T) {
	p := azureNativeProvider{
		config: map[string]string{
			"oidcRequestToken": "t1",
			"oidcRequestUrl":   "url1",
		},
	}

	conf, err := p.determineOidcConfig()
	assert.Nil(t, err)
	assert.Equal(t, "t1", conf.oidcRequestToken)
	assert.Equal(t, "url1", conf.oidcRequestUrl)
	assert.Empty(t, conf.oidcToken)
}

func TestOidcConfigWithTokenFromEnv(t *testing.T) {
	p := azureNativeProvider{}

	t.Setenv("ARM_OIDC_TOKEN", "t2")

	conf, err := p.determineOidcConfig()
	assert.Nil(t, err)
	assert.Empty(t, conf.oidcRequestToken)
	assert.Empty(t, conf.oidcRequestUrl)
	assert.Equal(t, "t2", conf.oidcToken)
}

func TestOidcConfigWithTokenFromConfig(t *testing.T) {
	p := azureNativeProvider{
		config: map[string]string{
			"oidcToken": "t3",
		},
	}

	conf, err := p.determineOidcConfig()
	assert.Nil(t, err)
	assert.Empty(t, conf.oidcRequestToken)
	assert.Empty(t, conf.oidcRequestUrl)
	assert.Equal(t, "t3", conf.oidcToken)
}

func TestOidcWithTokenFileFromEnv(t *testing.T) {
	assertConf := func(t *testing.T, p azureNativeProvider) {
		conf, err := p.determineOidcConfig()
		assert.Nil(t, err)
		assert.Empty(t, conf.oidcRequestToken)
		assert.Empty(t, conf.oidcRequestUrl)
		assert.Empty(t, conf.oidcToken)
		assert.Equal(t, "file1", conf.oidcTokenFilePath)
	}

	t.Run("FromEnv", func(t *testing.T) {
		p := azureNativeProvider{}
		t.Setenv("ARM_OIDC_TOKEN_FILE_PATH", "file1")
		assertConf(t, p)
	})

	t.Run("FromConf", func(t *testing.T) {
		p := azureNativeProvider{
			config: map[string]string{
				"oidcTokenFilePath": "file1",
			},
		}
		assertConf(t, p)
	})
}

func TestOidcEmptyConfig(t *testing.T) {
	p := azureNativeProvider{}

	t.Setenv("ARM_OIDC_TOKEN", "")
	t.Setenv("ARM_OIDC_TOKEN_FILE_PATH", "")
	t.Setenv("ARM_OIDC_REQUEST_TOKEN", "")
	t.Setenv("ARM_OIDC_REQUEST_URL", "")
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "")
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", "")

	_, err := p.determineOidcConfig()
	assert.Error(t, err)
}

func TestOidcUrlTokenPairValidation(t *testing.T) {
	p := azureNativeProvider{}

	// With a request token we also need a request URL.
	t.Setenv("ARM_OIDC_REQUEST_TOKEN", "t1")
	t.Setenv("ARM_OIDC_REQUEST_URL", "")
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "")
	t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", "")

	_, err := p.determineOidcConfig()
	assert.Error(t, err)
}

func TestOidcPrefersToken(t *testing.T) {
	p := azureNativeProvider{
		config: map[string]string{
			"oidcRequestToken": "t1",
			"oidcRequestUrl":   "url1",
			"oidcToken":        "t2",
		},
	}

	config, err := p.determineOidcConfig()
	assert.Nil(t, err)
	assert.NotEmpty(t, config.oidcToken)
}

// For the CLI authentication method, getAuthConfig() will shell out to the `az` CLI program to determine the
// cloud and subscription. This makes the test slower and will fail if `az` is not installed. In addition,
// go-azure-helpers only allows using `az` with user authentication, not service principal, which azure/login
// doesn't support. Therefore, we don't test the CLI authentication method by default. To test it locally, set
// PULUMI_RUN_CLI_AUTH_TESTS="true".
func TestAuthConfigs(t *testing.T) {
	// Unset the service principal env vars for this test so each subtest starts with a clean slate.
	t.Setenv("ARM_TENANT_ID", "")
	t.Setenv("ARM_CLIENT_ID", "")
	t.Setenv("ARM_CLIENT_SECRET", "")

	type providerFactory func() azureNativeProvider

	authMethods := map[string]providerFactory{
		"oidc": func() azureNativeProvider {
			return azureNativeProvider{
				config: map[string]string{
					"useOidc":   "true",
					"oidcToken": "t3",
					"tenantId":  "123",
					"clientId":  "123",
				},
			}
		},
		"oidcTokenFile": func() azureNativeProvider {
			return azureNativeProvider{
				config: map[string]string{
					"useOidc":           "true",
					"oidcTokenFilePath": "file",
					"tenantId":          "123",
					"clientId":          "123",
				},
			}
		},
		"clientSecret": func() azureNativeProvider {
			return azureNativeProvider{
				config: map[string]string{
					"tenantId":       "123",
					"clientId":       "123",
					"clientSecret":   "123",
					"subscriptionId": "123",
				},
			}
		},
		"msi": func() azureNativeProvider {
			return azureNativeProvider{
				config: map[string]string{
					"useMsi": "true",
				},
			}
		},
	}

	if os.Getenv("PULUMI_RUN_CLI_AUTH_TESTS") == "true" {
		authMethods["cli"] = func() azureNativeProvider {
			return azureNativeProvider{
				config: map[string]string{}, // CLI is fallback so doesn't need configuration
			}
		}
	}

	for authMethod, f := range authMethods {
		t.Run(authMethod+" is valid", func(t *testing.T) {
			p := f()

			conf, err := p.getAuthConfig()
			require.NoError(t, err)
			require.NotNil(t, conf)

			assert.Equal(t, authMethod == "cli", conf.useCli)

			assert.Equal(t, "public", conf.Environment)
		})

		t.Run(authMethod+" public env", func(t *testing.T) {
			p := f()
			p.config["environment"] = "public"

			conf, err := p.getAuthConfig()
			require.NoError(t, err)
			assert.Equal(t, "public", conf.Environment)
		})

		t.Run(authMethod+" usgov env", func(t *testing.T) {
			p := f()
			p.config["environment"] = "usgovernment"

			// The CLI auth method requires that the user switches `az` to the desired cloud/environment manually.
			// It will fail here and we assert an appropriate error.
			conf, err := p.getAuthConfig()
			if authMethod == "cli" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), "configured Azure environment 'usgovernment' does not match the determined environment 'public'")
			} else {
				require.NoError(t, err)
				assert.Equal(t, "usgovernment", conf.Environment)
			}
		})
	}
}

func TestGetAuthConfigCloud(t *testing.T) {
	t.Run("Default is public", func(t *testing.T) {
		p := azureNativeProvider{
			config: map[string]string{
				"useMsi": "true",
			},
		}
		conf, err := p.getAuthConfig()
		require.NoError(t, err)
		assert.Equal(t, azcloud.AzurePublic, conf.cloud())
	})

	t.Run("Public", func(t *testing.T) {
		p := azureNativeProvider{
			config: map[string]string{
				"environment": "public",
				"useMsi":      "true",
			},
		}
		conf, err := p.getAuthConfig()
		require.NoError(t, err)
		assert.Equal(t, azcloud.AzurePublic, conf.cloud())
	})

	t.Run("China", func(t *testing.T) {
		p := azureNativeProvider{
			config: map[string]string{
				"environment": "china",
				"useMsi":      "true",
			},
		}
		conf, err := p.getAuthConfig()
		require.NoError(t, err)
		assert.Equal(t, azcloud.AzureChina, conf.cloud())
	})

	t.Run("US Government", func(t *testing.T) {
		p := azureNativeProvider{
			config: map[string]string{
				"environment": "usgovernment",
				"useMsi":      "true",
			},
		}
		conf, err := p.getAuthConfig()
		require.NoError(t, err)
		assert.Equal(t, azcloud.AzureGovernment, conf.cloud())
	})

	t.Run("Public from environment variable", func(t *testing.T) {
		p := azureNativeProvider{}
		t.Setenv("ARM_ENVIRONMENT", "public")
		t.Setenv("ARM_USE_MSI", "true")
		conf, err := p.getAuthConfig()
		require.NoError(t, err)
		assert.Equal(t, azcloud.AzurePublic, conf.cloud())
	})

	t.Run("China from environment variable", func(t *testing.T) {
		p := azureNativeProvider{}
		t.Setenv("ARM_ENVIRONMENT", "china")
		t.Setenv("ARM_USE_MSI", "true")
		conf, err := p.getAuthConfig()
		require.NoError(t, err)
		assert.Equal(t, azcloud.AzureChina, conf.cloud())
	})

	t.Run("US Government from environment variable", func(t *testing.T) {
		p := azureNativeProvider{}
		t.Setenv("ARM_ENVIRONMENT", "usgovernment")
		t.Setenv("ARM_USE_MSI", "true")
		conf, err := p.getAuthConfig()
		require.NoError(t, err)
		assert.Equal(t, azcloud.AzureGovernment, conf.cloud())
	})
}

func TestGetCloud(t *testing.T) {
	t.Run("Public", func(t *testing.T) {
		a := authConfig{
			Config: &authentication.Config{
				Environment: "public",
			},
		}
		assert.Equal(t, azcloud.AzurePublic, a.cloud())
	})

	t.Run("China", func(t *testing.T) {
		a := authConfig{
			Config: &authentication.Config{
				Environment: "china",
			},
		}
		assert.Equal(t, azcloud.AzureChina, a.cloud())
	})

	t.Run("US Government", func(t *testing.T) {
		a := authConfig{
			Config: &authentication.Config{
				Environment: "usgov",
			},
		}
		assert.Equal(t, azcloud.AzureGovernment, a.cloud())
	})

	t.Run("Unknown", func(t *testing.T) {
		a := authConfig{}
		assert.Equal(t, azcloud.AzurePublic, a.cloud())
	})
}

func TestReadAzureEnvironmentFromConfig(t *testing.T) {
	// Reset environment
	t.Setenv("ARM_ENVIRONMENT", "")
	t.Setenv("AZURE_ENVIRONMENT", "")

	t.Run("Default to Public", func(t *testing.T) {
		p := azureNativeProvider{config: map[string]string{}}
		assert.Equal(t, "public", readAzureEnvironmentFromConfig(&p))
	})

	t.Run("From config", func(t *testing.T) {
		p := azureNativeProvider{config: map[string]string{"environment": "azureusgovernment"}}
		assert.Equal(t, "azureusgovernment", readAzureEnvironmentFromConfig(&p))
	})

	t.Run("From AZURE_ENVIRONMENT", func(t *testing.T) {
		p := azureNativeProvider{config: map[string]string{}}
		t.Setenv("AZURE_ENVIRONMENT", "azureusgovernment")
		assert.Equal(t, "azureusgovernment", readAzureEnvironmentFromConfig(&p))
	})

	t.Run("From ARM_ENVIRONMENT", func(t *testing.T) {
		p := azureNativeProvider{config: map[string]string{}}
		t.Setenv("ARM_ENVIRONMENT", "azureusgovernment")
		assert.Equal(t, "azureusgovernment", readAzureEnvironmentFromConfig(&p))
	})

	t.Run("ARM_ENVIRONMENT over AZURE_ENVIRONMENT", func(t *testing.T) {
		p := azureNativeProvider{config: map[string]string{}}
		t.Setenv("ARM_ENVIRONMENT", "azureusgovernment")
		t.Setenv("AZURE_ENVIRONMENT", "azurechina")
		assert.Equal(t, "azureusgovernment", readAzureEnvironmentFromConfig(&p))
	})
}
