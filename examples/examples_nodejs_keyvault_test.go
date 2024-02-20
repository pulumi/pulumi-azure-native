// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccKeyVaultTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "keyvault"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccKeyVaultTs_OICD(t *testing.T) {
	oidcClientId := os.Getenv("OIDC_ARM_CLIENT_ID")
	if oidcClientId == "" {
		t.Skip("Skipping OIDC test without OIDC_ARM_CLIENT_ID")
	}
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "keyvault"),
			Env: []string{
				"ARM_USE_OIDC=true",
				"ARM_CLIENT_ID=" + oidcClientId,
				// not strictly necessary but making sure we test the OIDC path
				"ARM_CLIENT_SECRET=",
			},
		})

	integration.ProgramTest(t, &test)
}

// This test is almost like TestAccKeyVaultTs_OICD but uses an explicit provider.
// We want to test configuring the provider via its arguments, not the environment.
func TestAccKeyVaultTs_OICDExplicit(t *testing.T) {
	skipIfShort(t)

	oidcClientId := os.Getenv("OIDC_ARM_CLIENT_ID")
	if oidcClientId == "" {
		t.Skip("Skipping OIDC test without OIDC_ARM_CLIENT_ID")
	}

	// These variables are set by GH. The provider reads them automatically, so we unset
	// them and use custom ones instead to be sure we're actually testing the explicit provider path.
	oidcToken := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
	oidcUrl := os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "keyvault-explicit-provider"),
			Env: []string{
				"OIDC_TOKEN_TEST=" + oidcToken,
				"OIDC_URL_TEST=" + oidcUrl,

				// unset to make sure we test the right code path
				"ACTIONS_ID_TOKEN_REQUEST_TOKEN=",
				"ACTIONS_ID_TOKEN_REQUEST_URL=",
				"ARM_CLIENT_SECRET="},
		})

	integration.ProgramTest(t, &test)
}
