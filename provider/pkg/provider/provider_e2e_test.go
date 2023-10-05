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

	"github.com/pulumi/pulumi/pkg/v3/codegen"
	rp "github.com/pulumi/pulumi/pkg/v3/resource/provider"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/deepcopy"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/rpcutil"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

func TestStorageBlob(t *testing.T) {
	newAzureTest("storage-blob").Run(t)
}

func TestApi(t *testing.T) {
	newAzureTest("api").
		SkipMatrixLanguages("csharp"). // Cannot implicitly convert type 'string[]' to 'Pulumi.InputList<Pulumi.Union<string, Pulumi.AzureNative.ApiManagement.Protocol>>'
		Run(t)
}

func TestServiceBusRecreate(t *testing.T) {
	newAzureTest("servicebus-recreate-1").
		AddEditDir("../servicebus-recreate-2").
		Run(t)
}

func TestFunctionScmFtpDeletion(t *testing.T) {
	newAzureTest("function-scm-ftp-deletion").
		SkipMatrixLanguages("csharp"). // 'ListStorageAccountKeysResult' does not contain a definition for 'Apply' and no accessible extension method 'Apply' accepting a first argument of type 'ListStorageAccountKeysResult' could be found
		SkipMatrixLanguages("python"). // <lambda>() missing 1 required positional argument: 'storage_account_keys'
		Run(t)
}

func newAzureTest(dir string) *testContext {
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

	return NewE2eTest(integration.ProgramTestOptions{
		Dir: filepath.Join(cwd, "test-programs", dir),
		Config: map[string]string{
			"azure-native:location": azureLocation,
		},
		ExpectRefreshChanges: true,
	}).WithProvider(startProvider).
		SkipMatrixLanguages("go") // Conversion fails due to split SDK
}

// startProvider starts the provider in a goProc and returns the port it's listening on.
// To shut down the provider, cancel the context.
func startProvider(ctx context.Context) (*RunningProvider, error) {
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

	return &RunningProvider{Name: "azure-native", Port: handle.Port}, nil
}

///////////// Generic matrix test code below ///////////////

func NewE2eTest(opts integration.ProgramTestOptions) *testContext {
	matrixTestLanguagesEnv, ok := os.LookupEnv("PULUMI_MATRIX_TEST_LANGUAGES")
	var matrixTestLanguages []string
	if ok {
		matrixTestLanguages = strings.Split(matrixTestLanguagesEnv, ",")
	} else {
		matrixTestLanguages = []string{"python", "csharp", "typescript", "go"}
	}
	return &testContext{
		matrixTestLanguages: matrixTestLanguages,
		skippedLanguages:    codegen.StringSet{},
		Options:             &opts,
	}
}

type testContext struct {
	matrixTestLanguages []string
	skippedLanguages    codegen.StringSet
	startProvider       *StartProvider // This should be pat of ProgramTestOptions
	Options             *integration.ProgramTestOptions
}

type StartProvider func(ctx context.Context) (*RunningProvider, error)

type RunningProvider struct {
	Name string
	Port int
}

func (tc *testContext) AddEditDir(editDirs ...string) *testContext {
	for _, editDir := range editDirs {
		// Resolve path relative to the initial program
		resolvedEditDir := filepath.Join(tc.Options.Dir, editDir)
		tc.Options.EditDirs = append(tc.Options.EditDirs, integration.EditDir{
			Dir:      resolvedEditDir,
			Additive: true,
		})
	}
	return tc
}

func (tc *testContext) WithProvider(startProvider StartProvider) *testContext {
	tc.startProvider = &startProvider
	return tc
}

func (tc *testContext) SkipMatrixLanguages(languages ...string) *testContext {
	for _, language := range languages {
		tc.skippedLanguages.Add(language)
	}
	return tc
}

func (tc *testContext) Run(t *testing.T) {
	if tc.startProvider == nil {
		start := *tc.startProvider
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		runningProvider, err := start(ctx)
		if err != nil {
			t.Fatal(err)
		}

		tc.Options.Env = []string{
			fmt.Sprintf("PULUMI_DEBUG_PROVIDERS=%s:%d", runningProvider.Name, runningProvider.Port),
		}
	}

	t.Run("yaml", func(t *testing.T) {
		integration.ProgramTest(t, tc.Options)
	})
	for _, language := range tc.matrixTestLanguages {
		currentLanguage := language
		t.Run(currentLanguage, func(t *testing.T) {
			if tc.skippedLanguages.Has(language) {
				t.Skip("skipping language")
				return
			}

			opts := CopyOptions(tc.Options)
			convertedDir := t.TempDir()
			opts.Dir = convertedDir

			cmd := exec.Command("pulumi", "convert", "--language", currentLanguage, "--generate-only", "--out", convertedDir)
			cmd.Dir = tc.Options.Dir
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("failed to convert to nodejs: %s\n%s", err, out)
			}

			for i, editDir := range opts.EditDirs {
				nodejsEditDir := t.TempDir()
				opts.EditDirs[i].Dir = nodejsEditDir

				cmd := exec.Command("pulumi", "convert", "--language", currentLanguage, "--generate-only", "--out", convertedDir)
				cmd.Dir = editDir.Dir
				out, err := cmd.CombinedOutput()
				if err != nil {
					t.Fatalf("failed to convert to nodejs: %s\n%s", err, out)
				}
			}

			integration.ProgramTest(t, opts)
		})
	}
}

func CopyOptions(opts *integration.ProgramTestOptions) *integration.ProgramTestOptions {
	copy := deepcopy.Copy(*opts)
	newOpts, ok := copy.(integration.ProgramTestOptions)
	if !ok {
		panic("failed to copy options")
	}
	return &newOpts
}

func skipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
}
