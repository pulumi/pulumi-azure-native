// Copyright 2022, Pulumi Corporation.

package provider

import (
	"testing"

	goversion "github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
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

func TestOidcEmpyConfig(t *testing.T) {
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
