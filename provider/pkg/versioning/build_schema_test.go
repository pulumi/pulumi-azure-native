package versioning

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleStability(t *testing.T) {
	// These are examples of resources which have been unstable in the docs schema generation in the past.
	// We disable these tests to avoid breaking the build. Re-enable them to continue investigating the issues.
	t.Skip("Skipping unstable tests")
	t.Run("AppPlatform_Deployment", func(t *testing.T) {
		testExampleGeneration(t, "AppPlatform", "2023-05-01-preview", "Deployment")
	})
	t.Run("ApiManagement_WorkspaceApiRelease", func(t *testing.T) {
		testExampleGeneration(t, "ApiManagement", "2022-09-01-preview", "WorkspaceApiRelease")
	})
	t.Run("Authorization_RoleManagementPolicyAssignment", func(t *testing.T) {
		testExampleGeneration(t, "Authorization", "2020-10-01", "RoleManagementPolicyAssignment")
	})
	t.Run("AzureStackHci_UpdateRun", func(t *testing.T) {
		testExampleGeneration(t, "AzureStackHCI", "2023-03-01", "UpdateRun")
	})
	t.Run("DataMigration_ServiceTask", func(t *testing.T) {
		testExampleGeneration(t, "DataMigration", "2021-06-30", "ServiceTask")
	})
	t.Run("MachineLearningServices_LabelingJob", func(t *testing.T) {
		// Often hangs - unclear why this is.
		t.Skip("Skipping long-running test")
		testExampleGeneration(t, "MachineLearningServices", "2023-04-01-preview", "LabelingJob")
	})
	t.Run("Network_NetworkInterface", func(t *testing.T) {
		// Observed failing on 9th iteration
		testExampleGeneration(t, "Network", "2023-02-01", "NetworkInterface")
	})
}

func testExampleGeneration(t *testing.T, namespace, versionFilter, resource string) {
	if t.Skipped() {
		return
	}
	firstRun, err := generateResourceExamples(namespace, versionFilter, resource)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		secondRun, err := generateResourceExamples(namespace, versionFilter, resource)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, firstRun, secondRun, "Failed on iteration %d", i+1)
		if t.Failed() {
			return
		}
	}
}

func generateResourceExamples(namespace, versionFilter, resourceToken string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	rootDir := path.Join(wd, "..", "..", "..")

	buildSchemaResult, err := BuildSchema(BuildSchemaArgs{
		RootDir: rootDir,
		Specs: ReadSpecsArgs{
			NamespaceFilter: namespace,
			VersionsFilter:  versionFilter,
		},
		ExcludeExplicitVersions: true,
		ExampleLanguages:        []string{"nodejs", "dotnet", "python", "go", "java", "yaml"},
		Version:                 "2.0.0",
	})

	if err != nil {
		return "", err
	}
	token := fmt.Sprintf("azure-native:%s:%s", strings.ToLower(namespace), resourceToken)
	resource, ok := buildSchemaResult.PackageSpec.Resources[token]
	if !ok {
		return "", fmt.Errorf("resource %s not found", token)
	}
	return resource.Description, nil
}
