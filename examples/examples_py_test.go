// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build python || all
// +build python all

package examples

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func TestAccSimplePython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-simple"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccLoadBalancer(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-loadbalancer"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccLoadBalancerSubResource(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-loadbalancer-subresource"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccNSG(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-nsg"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccUserAssignedIdentityPython(t *testing.T) {
	test := getPythonBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "py-user-assigned-identity"),
		})

	integration.ProgramTest(t, &test)
}

func getPythonBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	sdkPath := filepath.Join("..", "sdk", "python", "bin")
	if _, err := os.Stat(sdkPath); os.IsNotExist(err) {
		t.Fatalf("python SDK not found at %s, run `make build_python` first", sdkPath)
	}
	basePy := base.With(integration.ProgramTestOptions{
		Dependencies: []string{sdkPath},
	})

	return basePy
}
