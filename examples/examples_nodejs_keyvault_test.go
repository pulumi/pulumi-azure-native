// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all

package examples

import (
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"os"
	"path/filepath"
	"testing"
)

func TestAccKeyVaultTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "keyvault"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccKeyVaultTs_OICD(t *testing.T) {
	skipIfShort(t)

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
