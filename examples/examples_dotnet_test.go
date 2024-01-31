// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build dotnet || all
// +build dotnet all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccAppServiceDockerDotnet(t *testing.T) {
	t.Skip("temporarily skipping until a compatible docker release for 3.0 has been made")
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "cs-appservice-docker"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccSimpleDotnet(t *testing.T) {
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "cs-simple"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccSql(t *testing.T) {
	skipIfShort(t)
	test := getCsharpBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "cs-sql"),
		})

	test.RequireEmptyPreviewAfterRefresh = false
	integration.ProgramTest(t, &test)
}

func getCsharpBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseCsharp := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"Pulumi.AzureNative",
		},
	})

	return baseCsharp
}
