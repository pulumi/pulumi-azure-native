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

	"github.com/pulumi/providertest"
	rp "github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/rpcutil"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

func TestStorageBlob(t *testing.T) {
	newAzureTest("storage-blob").Run(t)
}

func TestApi(t *testing.T) {
	newAzureTest("api",
		providertest.WithSkipSdk("csharp", "fails with: Cannot implicitly convert type 'string[]' to 'Pulumi.InputList<Pulumi.Union<string, Pulumi.AzureNative.ApiManagement.Protocol>>'"),
	).Run(t)
}

func TestServiceBusRecreate(t *testing.T) {
	newAzureTest("servicebus-recreate-1",
		providertest.WithUpdateStep(providertest.UpdateStepDir("../servicebus-recreate-2")),
	).Run(t)
}

func TestFunctionScmFtpDeletion(t *testing.T) {
	newAzureTest("function-scm-ftp-deletion",
		providertest.WithSkipSdk("csharp", "fails with: 'ListStorageAccountKeysResult' does not contain a definition for 'Apply' and no accessible extension method 'Apply' accepting a first argument of type 'ListStorageAccountKeysResult' could be found"),
		providertest.WithSkipSdk("python", "fails with: <lambda>() missing 1 required positional argument: 'storage_account_keys'"),
	).Run(t)
}

func newAzureTest(dir string, opts ...providertest.Option) *providertest.ProviderTest {
	azureLocation := os.Getenv("ARM_LOCATION")
	if azureLocation == "" {
		azureLocation = "westus2"
		fmt.Println("Defaulting location to 'westus2'. You can override using the ARM_LOCATION variable.")
	}

	cwd, err := os.Getwd()
	if err != nil {
		// If we can't get the current working directory, we'll fail hard.
		panic(err)
	}

	opts = append(opts,
		providertest.WithProvider(startProvider),
		providertest.WithConfig("azure-native:location", azureLocation),
		providertest.WithSkipSdk("go", "generated SDK imports don't match split SDK layout"),
		providertest.WithE2eOptions(func(opts *integration.ProgramTestOptions) {
			opts.ExpectRefreshChanges = true // Azure Native causes diffs after a refresh
		}),
	)

	return providertest.NewProviderTest(
		filepath.Join(cwd, "test-programs", dir),
		opts...,
	)
}

// startProvider starts the provider in a goProc and returns the port it's listening on.
// To shut down the provider, cancel the context.
func startProvider(ctx context.Context) (*providertest.ProviderAttach, error) {
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
		return nil, err
	}
	azureAPIResourcesBytes, err := os.ReadFile(filepath.Join("..", "..", "..", "bin", "metadata-compact.json"))
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("fatal: %v", err)
	}

	return &providertest.ProviderAttach{Name: providerName, Port: handle.Port}, nil
}

func skipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
}
