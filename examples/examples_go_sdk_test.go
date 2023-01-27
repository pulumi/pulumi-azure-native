// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

const sdkPath = "../sdk/pulumi-azure-native-sdk"

func testDir(t *testing.T, testCaseDirs ...string) string {
	segments := []string{getCwd(t), "azure-native-sdk"}
	segments = append(segments, testCaseDirs...)
	return filepath.Join(segments...)
}

func TestAccSimpleGoSdk(t *testing.T) {
	test := getGoBaseOptionsSdk(t).
		With(integration.ProgramTestOptions{
			Dir: testDir(t, "go-simple"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccClientConfigGoSdk(t *testing.T) {
	test := getGoBaseOptionsSdk(t).
		With(integration.ProgramTestOptions{
			Dir: testDir(t, "go-clientconfig"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccAksGoSdk(t *testing.T) {
	t.Skip("Disabled due to https://github.com/pulumi/pulumi-azure-native/issues/304")
	test := getGoBaseOptionsSdk(t).
		With(integration.ProgramTestOptions{
			Dir: testDir(t, "go-aks"),
		})

	integration.ProgramTest(t, &test)
}

func TestServicebusRecreateSdk(t *testing.T) {
	skipIfShort(t)
	test := getGoBaseOptionsSdk(t).
		With(integration.ProgramTestOptions{
			Dir: testDir(t, "go-servicebus-recreate", "step1"),
			EditDirs: []integration.EditDir{
				{
					Dir:      testDir(t, "go-servicebus-recreate", "step2"),
					Additive: true,
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func getGoBaseOptionsSdk(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)

	required, err := getRequiredGoPackages()
	if err != nil {
		t.Fatal(err)
	}

	rootSdkPath, err := filepath.Abs(sdkPath)
	require.NoError(t, err)

	dependencies := make([]string, len(required))
	for i, pkg := range required {
		dependencies[i] = genReplace(t, rootSdkPath, pkg)
	}

	dependencies = append(dependencies, "github.com/pulumi/pulumi-azure-native-sdk="+rootSdkPath)

	baseGo := base.With(integration.ProgramTestOptions{
		Dependencies: dependencies,
	})

	return baseGo
}

// We don't know which packages the tests need, so we just get all of them.
func getRequiredGoPackages() ([]string, error) {
	dirs, err := os.ReadDir(sdkPath)
	if err != nil {
		return nil, err
	}

	packages := make([]string, 0)
	for _, entry := range dirs {
		if entry.IsDir() {
			packages = append(packages, entry.Name())
		}
	}

	return packages, nil
}

// For context: https://github.com/pulumi/pulumi/pull/11055
func genReplace(t *testing.T, sdkPath string, subPkg string) string {
	moduleDir := fmt.Sprintf("%s/%s", sdkPath, subPkg)
	return fmt.Sprintf("github.com/pulumi/pulumi-azure-native-sdk/%s=%s", subPkg, moduleDir)
}
