// Copyright 2023, Pulumi Corporation.  All rights reserved.

// Disable running if we've specifically selected unit tests to run as this is an integration test.
// This is easier than having to remember to explicitly tag every unit test with `//go:build unit || all`.
//go:build !unit

package provider

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"

	"github.com/pulumi/providertest/providers"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertrefresh"
	"github.com/pulumi/providertest/pulumitest/opttest"

	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

func TestStorageBlob(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "storage-blob")
	pt.Preview(t)
	pt.Up(t)
	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}

func TestApi(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "api")
	pt.Preview(t)
	pt.Up(t)
	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}

func TestRequiredContainers(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "required-containers")
	pt.Preview(t)
	pt.Up(t)
	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}

func TestParallelSubnetCreation(t *testing.T) {
	t.Parallel()
	if !util.EnableAzcoreBackend() {
		t.Skip("Skipping test because it requires the AZCore backend")
	}
	pt := newPulumiTest(t, "parallel-subnet-creation")
	pt.Preview(t)
	pt.Up(t)
	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}

func newPulumiTest(t *testing.T, testProgramDir string) *pulumitest.PulumiTest {
	t.Helper()
	if testing.Short() {
		t.Skipf("Skipping in testing.Short() mode, assuming this is a CI run without cloud credentials")
		return nil
	}
	dir := filepath.Join("test-programs", testProgramDir)
	azureLocation := getLocation()
	rpFactory := providers.ResourceProviderFactory(providerServer)
	attachOpt := opttest.AttachProvider("azure", rpFactory)
	pt := pulumitest.NewPulumiTest(t, dir, attachOpt)
	pt.SetConfig(t, "azure-native:location", azureLocation)
	return pt
}

func providerServer(_ providers.PulumiTest) (pulumirpc.ResourceProviderServer, error) {
	version.Version = "0.0.1"

	schemaBytes, err := os.ReadFile(filepath.Join("..", "..", "..", "bin", "schema-full.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to read schema file, run `make schema` before running tests: %v", err)
	}

	azureAPIResourcesBytes, err := os.ReadFile(filepath.Join("..", "..", "..", "bin", "metadata-compact.json"))
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata file, run `make schema` before running tests: %v", err)
	}

	return makeProvider(nil, "azure-native", "2.0.0", schemaBytes, azureAPIResourcesBytes)
}

func getLocation() string {
	azureLocation := os.Getenv("ARM_LOCATION")
	if azureLocation == "" {
		azureLocation = "westus2"
		fmt.Println("Defaulting location to 'westus2'. You can override using the ARM_LOCATION variable.")
	}

	return azureLocation
}

func skipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
}
