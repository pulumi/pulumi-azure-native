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
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "keyvault"),
		})

	integration.ProgramTest(t, &test)
}
