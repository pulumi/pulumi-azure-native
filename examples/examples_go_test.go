// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
// +build go all

package examples

import (
	"testing"

	"github.com/pulumi/pulumi/pkg/v2/testing/integration"
)

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/pulumi/pulumi-azurerm/sdk",
		},
	})

	return baseGo
}
