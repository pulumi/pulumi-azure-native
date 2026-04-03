package gen

import (
	"os"
	"path"
	"testing"

	"github.com/blang/semver"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
	"github.com/stretchr/testify/assert"
)

func TestManagedClusterGen(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("error while getting current working directory: %v", err)
	}

	rootDir := path.Join(wd, "test-data", "managedcluster", "azure-rest-api-specs")
	modules, _, err := openapi.ReadAzureModules(rootDir, "", "")
	if err != nil {
		t.Fatalf("error while reading azure modules: %v", err)
	}

	assert.NotEmpty(t, modules, "read modules from azure-specs are not empty")
	modules = openapi.ApplyTransformations(modules, openapi.DefaultVersions{
		"ContainerService": {
			"ManagedCluster":    {ApiVersion: "2025-10-01"},
			"NodeCustomization": {ApiVersion: "2025-09-02-preview"}, // ← adds 2-prop conflict
		},
	}, openapi.DefaultVersions{}, nil, nil)
	versioning := versioningStub{
		previousTokenPaths: map[string]string{},
	}
	generationResult, err := PulumiSchema(rootDir, modules, versioning, semver.MustParse("3.0.0"), true /* onlyExplicitVersions */)
	assert.NoError(t, err, "error while generating Pulumi schema")

	// Ensure the managed cluster resource is present in schema and metadata
	// and snapshot the generation result so we can see the impact of future refactors.
	managedClusterResource, ok := generationResult.Schema.Resources["azure-native:containerservice:ManagedCluster"]
	assert.True(t, ok, "ManagedCluster resource not found in generated schema")
	snaps.MatchJSON(t, managedClusterResource)

	managedClusterResourceMeta, ok := generationResult.Metadata.Resources["azure-native:containerservice:ManagedCluster"]
	assert.True(t, ok, "ManagedCluster resource not found in generated metadata")
	snaps.MatchJSON(t, managedClusterResourceMeta)

	userManagedIdentity, ok := generationResult.Schema.Types["azure-native:containerservice:UserAssignedIdentityResponse"]
	assert.True(t, ok, "UserAssignedIdentityResponse type not found in generated schema...")
	assert.Equal(t, 3, len(userManagedIdentity.Properties), "UserAssignedIdentityResponse should have 3 properties")
	assert.Contains(t, userManagedIdentity.Properties, "resourceId", "UserAssignedIdentityResponse should contain resourceId property")
	assert.Contains(t, userManagedIdentity.Properties, "clientId", "UserAssignedIdentityResponse should contain clientId property")
	assert.Contains(t, userManagedIdentity.Properties, "objectId", "UserAssignedIdentityResponse should contain objectId property")
}
