// Copyright 2022, Pulumi Corporation.

package provider

import (
	"testing"

	goversion "github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptedAzVersions(t *testing.T) {
	goodVersions := []string{
		"2.0.81", "2.0.82", "2.1.81", "2.20", "2.33", "2.33.9",
		"2.37", "2.37.1", "2.38", "2.40", "2.99.99",
	}
	badVersions := []string{
		"1",
		"2", "2.0.80", "2.34", "2.34.1", "2.35", "2.36", "2.36.6",
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

func TestOidcEmptyConfig(t *testing.T) {
	p := azureNativeProvider{}

	_, err := p.determineOidcConfig()
	assert.NotNil(t, err)
}

func TestOidcUrlTokenPairValidation(t *testing.T) {
	p := azureNativeProvider{}

	// With a request token we also need a request URL.
	t.Setenv("ARM_OIDC_REQUEST_TOKEN", "t1")

	_, err := p.determineOidcConfig()
	assert.NotNil(t, err)
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

// Note: for the CLI authentication method, getAuthConfig() will shell out to the `az` CLI program to determine the
// cloud and subscription. This makes the test slower and will fail if `az` is not installed. By default, it is
// installed on GH runners, however, so we include the test for completeness.
func TestAuthConfigs(t *testing.T) {
	type providerFactory func() azureNativeProvider

	for authMethod, f := range map[string]providerFactory{
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
		"cli": func() azureNativeProvider {
			return azureNativeProvider{
				config: map[string]string{}, // CLI is fallback
			}
		},
	} {
		t.Run(authMethod+" is valid", func(t *testing.T) {
			p := f()
			conf, err := p.getAuthConfig()
			require.NoError(t, err)
			require.NotNil(t, conf)
		})

		t.Run(authMethod+" uses CLI", func(t *testing.T) {
			p := f()
			conf, _ := p.getAuthConfig()
			assert.Equal(t, authMethod == "cli", conf.useCli)
		})

		t.Run(authMethod+" default env", func(t *testing.T) {
			p := f()
			conf, _ := p.getAuthConfig()
			assert.Equal(t, "public", conf.Environment)
		})

		t.Run(authMethod+" public env", func(t *testing.T) {
			p := f()
			p.config["environment"] = "public"

			conf, _ := p.getAuthConfig()
			assert.Equal(t, "public", conf.Environment)
		})

		// The CLI auth method requires that the user switches `az` to the desired cloud/environment manually,
		// it will fail here and we assert an appropriate error.
		t.Run(authMethod+" usgov env", func(t *testing.T) {
			p := f()
			p.config["environment"] = "usgovernment"

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
