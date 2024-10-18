// Copyright 2023, Pulumi Corporation.  All rights reserved.

// Disable running if we've specifically selected unit tests to run as this is an integration test.
// This is easier than having to remember to explicitly tag every unit test with `//go:build unit || all`.
//go:build !unit

package provider

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"google.golang.org/grpc"

	rp "github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/rpcutil"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

func TestStorageBlob(t *testing.T) {
	runTestProgram(t, "storage-blob")
}

func TestApi(t *testing.T) {
	runTestProgram(t, "api")
}

func TestRequiredContainers(t *testing.T) {
	runTestProgram(t, "required-containers")
}

func TestParallelSubnetCreation(t *testing.T) {
	// Force this test to use the new AZCore client
	// We have to set this here rather than in the ProgramTest config because the provider is started up before starting the test.
	os.Setenv("PULUMI_ENABLE_AZCORE_BACKEND", "true")
	runTestProgram(t, "parallel-subnet-creation")
}

// runTestProgram runs an example from ./examples/<initialDir>
// Any editDirs are applied in order, and the program is run after each edit. e.g. ./examples/<editDir>
func runTestProgram(t *testing.T, initialDir string, editDirs ...string) {
	if t.Skipped() {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	port, err := startProvider(ctx)
	if err != nil {
		t.Fatal(err)
	}

	opts, err := getTestOptions(initialDir, editDirs, port)
	if err != nil {
		t.Fatal(err)
		return
	}

	integration.ProgramTest(t, opts)
}

func getTestOptions(initialDir string, editDirs []string, port int) (*integration.ProgramTestOptions, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	azureLocation := getLocation()
	test := integration.ProgramTestOptions{
		Dir:                  filepath.Join(cwd, "test-programs", initialDir),
		ExpectRefreshChanges: true,
		Config: map[string]string{
			"azure-native:location": azureLocation,
		},
		Env: []string{
			fmt.Sprintf("PULUMI_DEBUG_PROVIDERS=azure-native:%d", port),
		},
	}
	for _, editDir := range editDirs {
		test.EditDirs = append(test.EditDirs, integration.EditDir{
			Dir: filepath.Join(cwd, "test-programs", editDir),
		})
	}
	return &test, nil
}

// startProvider starts the provider in a goProc and returns the port it's listening on.
// To shut down the provider, cancel the context.
func startProvider(ctx context.Context) (int, error) {
	providerName := "azure-native"
	cancelChannel := make(chan bool)
	go func() {
		<-ctx.Done()
		close(cancelChannel)
	}()
	var host *rp.HostClient // We don't initialize this - to mimic attach mode
	version := ""
	schemaBytes, err := os.ReadFile(filepath.Join("..", "..", "..", "bin", "schema-full.json"))
	if err != nil {
		return 0, err
	}
	azureAPIResourcesBytes, err := os.ReadFile(filepath.Join("..", "..", "..", "bin", "metadata-compact.json"))
	if err != nil {
		return 0, err
	}
	handle, err := rpcutil.ServeWithOptions(rpcutil.ServeOptions{
		Cancel: cancelChannel,
		Init: func(srv *grpc.Server) error {
			prov, proverr := makeProvider(host, providerName, version, schemaBytes, azureAPIResourcesBytes)
			if proverr != nil {
				return fmt.Errorf("failed to create resource provider: %v", proverr)
			}
			pulumirpc.RegisterResourceProviderServer(srv, prov)
			return nil
		},
		Options: rpcutil.OpenTracingServerInterceptorOptions(nil),
	})
	if err != nil {
		return 0, fmt.Errorf("fatal: %v", err)
	}

	return handle.Port, nil
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
