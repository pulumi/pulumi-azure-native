// Copyright 2021, Pulumi Corporation.  All rights reserved.

package examples

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getLocation(t *testing.T) string {
	azureLocation := os.Getenv("ARM_LOCATION")
	if azureLocation == "" {
		azureLocation = "westus2"
		fmt.Println("Defaulting location to 'westus2'. You can override using the ARM_LOCATION variable.")
	}

	return azureLocation
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	azureLocation := getLocation(t)
	binPath, err := filepath.Abs("../bin")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Using binPath %s\n", binPath)
	return integration.ProgramTestOptions{
		ExpectRefreshChanges:            true,
		RequireEmptyPreviewAfterRefresh: true,
		Config: map[string]string{
			"azure-native:location": azureLocation,
		},
		LocalProviders: []integration.LocalDependency{
			{
				Package: "azure-native",
				Path:    binPath,
			},
		},
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func skipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
}
