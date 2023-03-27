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
	if os.Getenv("RUN_OIDC_TESTS") != "true" {
		t.Skip("Skipping OIDC test outside of CI")
	}
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "keyvault"),
			Env: []string{
				"ARM_USE_OIDC=true",
				// not strictly necessary but making sure we test the OIDC path
				"ARM_CLIENT_SECRET=",
				// AD app 'oidc-test', owned by tkappler - temporary until our main app is OIDC-enabled
				"ARM_CLIENT_ID=89380e12-5be6-486a-89ef-eea107af2f47",
			},
		})

	integration.ProgramTest(t, &test)
}
