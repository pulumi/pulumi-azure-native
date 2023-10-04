// Copyright 2023, Pulumi Corporation.  All rights reserved.

// Disable running if we've specifically selected unit tests to run as this is an integration test.
// This is easier than having to remember to explicitly tag every unit test with `//go:build unit || all`.
//go:build !unit

package provider

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
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

func TestServiceBusRecreate(t *testing.T) {
	runTestProgram(t, "servicebus-recreate-1", "service-bus-recreate-2")
}

func TestFunctionScmFtpDeletion(t *testing.T) {
	runTestProgram(t, "function-scm-ftp-deletion")
}

// runTestProgram runs an example from ./examples/<initialDir>
// Any editDirs are applied in order, and the program is run after each edit. e.g. ./examples/<editDir>
func runTestProgram(t *testing.T, initialDir string, editDirs ...string) {
	if t.Skipped() {
		return
	}
	t.Run("yaml", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		port, err := startProvider(ctx)
		if err != nil {
			t.Fatal(err)
		}

		opts := getYamlTestOptions(initialDir, editDirs, port)
		if err != nil {
			t.Fatal(err)
			return
		}
		integration.ProgramTest(t, opts)
	})

	// Run locally using: PULUMI_MATRIX_TEST_LANGUAGES=typescript,dotnet PROVIDER_TEST_TAGS=all make test_provider
	languagesEnv, enableMatrixTest := os.LookupEnv("PULUMI_MATRIX_TEST_LANGUAGES")
	if !enableMatrixTest {
		return
	}
	languages := strings.Split(languagesEnv, ",")
	for _, language := range languages {
		currentLanguage := language
		t.Run(currentLanguage, func(t *testing.T) {
			nodeDir := t.TempDir()

			cmd := exec.Command("pulumi", "convert", "--language", currentLanguage, "--generate-only", "--out", nodeDir)
			cmd.Dir = resolveTestProgramDir(initialDir)
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("failed to convert to nodejs: %s\n%s", err, out)
			}

			nodejsEditDirs := make([]string, len(editDirs))
			for _, editDir := range editDirs {
				nodejsEditDir := t.TempDir()
				nodejsEditDirs = append(nodejsEditDirs, nodejsEditDir)
				cmd := exec.Command("pulumi", "convert", "--language", currentLanguage, "--generate-only", "--out", nodeDir)
				cmd.Dir = resolveTestProgramDir(editDir)
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("failed to convert to nodejs: %s\n%s", err, out)
				}
			}

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			port, err := startProvider(ctx)
			if err != nil {
				t.Fatal(err)
			}

			opts := getTestOptions(nodeDir, nodejsEditDirs, port)
			integration.ProgramTest(t, opts)
		})
	}
}

func getYamlTestOptions(initialDir string, editDirs []string, port int) *integration.ProgramTestOptions {
	resolvedEditDirs := make([]string, len(editDirs))
	for i, editDir := range editDirs {
		resolvedEditDirs[i] = resolveTestProgramDir(editDir)
	}
	return getTestOptions(resolveTestProgramDir(initialDir), resolvedEditDirs, port)
}

func resolveTestProgramDir(dir string) string {
	cwd, err := os.Getwd()
	if err != nil {
		// If we can't get the current working directory, we'll fail hard.
		panic(err)
	}
	return filepath.Join(cwd, "test-programs", dir)
}

func getTestOptions(initialDir string, editDirs []string, port int) *integration.ProgramTestOptions {
	azureLocation := getLocation()
	test := integration.ProgramTestOptions{
		Dir:                  initialDir,
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
			Dir:      editDir,
			Additive: true,
		})
	}
	return &test
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
