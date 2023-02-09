// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build nodejs || (nodejs && oidc) || all

package examples

import (
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"path/filepath"
	"testing"
)

func TestAccKeyVaultTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions2(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "keyvault"),
		})

	integration.ProgramTest(t, &test)
}

// A copy of getJSBaseOptions from examples_nodejs_test.go because that one
// isn't built when testing with tag=oidc
func getJSBaseOptions2(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/azure-native",
		},
	})

	return baseJS
}
