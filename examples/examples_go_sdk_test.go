// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build go || all
// +build go all

package examples

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

const sdkPath = "../sdk/pulumi-azure-native-sdk"

func testDir(t *testing.T, testCaseDirs ...string) string {
	segments := []string{getCwd(t), "azure-native-sdk"}
	segments = append(segments, testCaseDirs...)
	return filepath.Join(segments...)
}

func TestAccSimpleGoSdk(t *testing.T) {
	test := getGoBaseOptionsSdk(t, testDir(t, "go-simple"))

	integration.ProgramTest(t, &test)
}

func TestAccClientConfigGoSdk(t *testing.T) {
	test := getGoBaseOptionsSdk(t, testDir(t, "go-clientconfig"))

	integration.ProgramTest(t, &test)
}

func TestAccUserAssignedIdentitySdk(t *testing.T) {
	test := getGoBaseOptionsSdk(t, testDir(t, "go-user-assigned-identity")).
		With(integration.ProgramTestOptions{
			RunUpdateTest: false,
		})

	integration.ProgramTest(t, &test)
}

func TestAccAksGoSdk(t *testing.T) {
	t.Skip("Disabled due to https://github.com/pulumi/pulumi-azure-native/issues/304")
	test := getGoBaseOptionsSdk(t, testDir(t, "go-aks"))

	integration.ProgramTest(t, &test)
}

func TestServicebusRecreateSdk(t *testing.T) {
	skipIfShort(t)
	test := getGoBaseOptionsSdk(t, testDir(t, "go-servicebus-recreate", "step1")).
		With(integration.ProgramTestOptions{
			EditDirs: []integration.EditDir{
				{
					Dir:      testDir(t, "go-servicebus-recreate", "step2"),
					Additive: true,
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func getGoBaseOptionsSdk(t *testing.T, dir string) integration.ProgramTestOptions {
	base := getBaseOptions(t)

	rootSdkPath, err := filepath.Abs(sdkPath)
	require.NoError(t, err)

	replacements, err := getSdkReplacements(dir, rootSdkPath)
	require.NoError(t, err)

	goDepRoot := os.Getenv("PULUMI_GO_DEP_ROOT")
	if goDepRoot == "" {
		goDepRoot, err = filepath.Abs("../..")
		if err != nil {
			t.Fatal(err)
		}
	}

	baseGo := base.With(integration.ProgramTestOptions{
		Dir:          dir,
		Dependencies: replacements,
		Env: []string{
			fmt.Sprintf("PULUMI_GO_DEP_ROOT=%s", goDepRoot),
		},
	})

	return baseGo
}

func getSdkReplacements(dir, rootSdkPath string) ([]string, error) {
	required, err := getRequiredPathsFromGoMod(dir)
	if err != nil {
		return nil, err
	}

	// Find pulumi-azure-native-sdk packages - ignoring the /vX suffix
	// Match repo root or sub-packages with optional version suffix
	matchAzureNativePackage := regexp.MustCompile(`^github\.com\/pulumi\/pulumi-azure-native-sdk\/?([^\/]*)(\/v\d+)$`)
	var replacements []string
	for _, pkg := range required {
		matches := matchAzureNativePackage.FindStringSubmatch(pkg)
		if len(matches) < 2 {
			continue
		}
		module := matches[1]
		modulePath := path.Join(rootSdkPath, module)
		replacement := fmt.Sprintf("%s=%s", pkg, modulePath)
		replacements = append(replacements, replacement)
	}
	return replacements, nil
}

func getRequiredPathsFromGoMod(dir string) ([]string, error) {
	// Run the command `go mod edit --json` and parse the output to get the list of required packages
	cmd := exec.Command("go", "mod", "edit", "--json")
	cmd.Dir = dir
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	// Deserialize output into a GoMod struct
	var goMod GoMod
	err = json.Unmarshal(out, &goMod)
	if err != nil {
		return nil, err
	}
	// Get the list of required packages
	var required []string
	for _, req := range goMod.Require {
		required = append(required, req.Path)
	}
	return required, nil
}

type Module struct {
	Path    string
	Version string
}

type GoMod struct {
	Module  ModPath
	Go      string
	Require []Require
	Exclude []Module
	Replace []Replace
	Retract []Retract
}

type ModPath struct {
	Path       string
	Deprecated string
}

type Require struct {
	Path     string
	Version  string
	Indirect bool
}

type Replace struct {
	Old Module
	New Module
}

type Retract struct {
	Low       string
	High      string
	Rationale string
}
