// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func dir(t *testing.T, testCaseDir string) string {
	return filepath.Join(getCwd(t), "azure-native-sdk", testCaseDir)
}

func TestAccSimpleGoSdk(t *testing.T) {
	test := getGoBaseOptionsSdk(t).
		With(integration.ProgramTestOptions{
			Dir: dir(t, "go-simple"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccClientConfigGoSdk(t *testing.T) {
	test := getGoBaseOptionsSdk(t).
		With(integration.ProgramTestOptions{
			Dir: dir(t, "go-clientconfig"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccAksGoSdk(t *testing.T) {
	t.Skip("Disabled due to https://github.com/pulumi/pulumi-azure-native/issues/304")
	test := getGoBaseOptionsSdk(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "go-aks"),
		})

	integration.ProgramTest(t, &test)
}

func TestServicebusRecreateSdk(t *testing.T) {
	test := getGoBaseOptionsSdk(t).
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

func getGoBaseOptionsSdk(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			genReplace(t, "authorization"),
			genReplace(t, "containerservices"),
			genReplace(t, "resources"),
			genReplace(t, "servicebus"),
			genReplace(t, "storage"),
		},
	})

	return baseGo
}

// For context: https://github.com/pulumi/pulumi/pull/11055
func genReplace(t *testing.T, subPkg string) string {
	moduleDir, err := filepath.Abs(fmt.Sprintf("../sdk/pulumi-azure-native-sdk/%s", subPkg))
	require.NoError(t, err)
	return fmt.Sprintf("github.com/pulumi/pulumi-azure-native-sdk/%s=%s", subPkg, moduleDir)
}
