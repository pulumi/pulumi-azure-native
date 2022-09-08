// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccSimpleGo(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "go-simple"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccClientConfigGo(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "go-clientconfig"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccAksGo(t *testing.T) {
	t.Skip("Disabled due to https://github.com/pulumi/pulumi-azure-native/issues/304")
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "go-aks"),
		})

	integration.ProgramTest(t, &test)
}

func TestServicebusRecreate(t *testing.T) {
	test := getGoBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "go-servicebus-recreate", "step1"),
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join(getCwd(t), "go-servicebus-recreate", "step2"),
					Additive: true,
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func getGoBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"github.com/pulumi/pulumi-azure-native/sdk/go",
		},
	})

	return baseGo
}
