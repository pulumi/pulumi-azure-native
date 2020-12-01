// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build nodejs all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func TestAccApiTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "api"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccAppServiceTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "appservice"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccCosmosDBTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "cosmosdb"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccKeyVaultTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "keyvault"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccSimpleTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "simple"),
		})

	integration.ProgramTest(t, &test)
}

func TestDestroyTs(t *testing.T) {
	// Tests eventually-consistent deletion handling.
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "static-website"),
		})

	integration.ProgramTest(t, &test)
}

func TestImportTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "import"),
		})

	integration.ProgramTest(t, &test)
}

func TestPostgresTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "postgres"),
		})

	integration.ProgramTest(t, &test)
}

func TestMessagingTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/azure-nextgen",
		},
	})

	return baseJS
}
