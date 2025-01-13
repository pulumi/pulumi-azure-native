package gen

import (
	"os"
	"path"
	"testing"

	"github.com/blang/semver"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

func TestAliasesGen(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// We take a snapshot of the schema as we're focusing on testing how we parse the spec
	// To update the specifications run in the repo root:
	// 1. rm -rf provider/pkg/gen/test-data/aliases/azure-rest-api-specs/specification/common-types/resource-management/v1/
	// 2. rm -rf provider/pkg/gen/test-data/aliases/azure-rest-api-specs/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/
	// 3. cp -r azure-rest-api-specs/specification/common-types/resource-management/v1 provider/pkg/gen/test-data/aliases/azure-rest-api-specs/specification/common-types/resource-management/
	// 4. cp -r azure-rest-api-specs/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01 provider/pkg/gen/test-data/aliases/azure-rest-api-specs/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/
	// 5. rm -rf provider/pkg/gen/test-data/aliases/azure-rest-api-specs/specification/azureactivedirectory/resource-manager/Microsoft.Aadiam/stable/2020-03-01/examples/
	rootDir := path.Join(wd, "test-data", "aliases")

	// azure-rest-api-specs/specification/common-types/resource-management/v1
	providers, _, err := openapi.ReadAzureProviders(path.Join(rootDir, "azure-rest-api-specs"), "Aadiam", "2020-03-01")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("v2", func(t *testing.T) {
		generationResult, err := PulumiSchema(rootDir, providers, versioningStub{}, semver.MustParse("2.0.0"))
		if err != nil {
			t.Fatal(err)
		}
		s := snaps.WithConfig(snaps.Filename("gen_aliases_test_v2"))
		s.MatchJSON(t, generationResult.Schema)
		s.MatchJSON(t, generationResult.Metadata)
	})

	t.Run("v3", func(t *testing.T) {
		generationResult, err := PulumiSchema(rootDir, providers, versioningStub{}, semver.MustParse("3.0.0"))
		if err != nil {
			t.Fatal(err)
		}
		s := snaps.WithConfig(snaps.Filename("gen_aliases_test_v3"))
		s.MatchJSON(t, generationResult.Schema)
		s.MatchJSON(t, generationResult.Metadata)
	})
}
