package gen

import (
	"os"
	"path"
	"testing"

	"github.com/blang/semver"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi"
)

func TestVnetGen(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// We take a snapshot of the schema as we're focusing on testing how we parse the spec
	// To update the specifications run in the repo root:
	// 1. rm -rf provider/pkg/gen/test-data/vnet/azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/
	// 2. cp -r azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01 provider/pkg/gen/test-data/vnet/azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/
	// 3. rm -rf provider/pkg/gen/test-data/vnet/azure-rest-api-specs/specification/network/resource-manager/Microsoft.Network/stable/2023-02-01/examples/
	rootDir := path.Join(wd, "test-data", "vnet")

	modules, _, err := openapi.ReadAzureModules(path.Join(rootDir, "azure-rest-api-specs"), "Network", "2023-02-01")
	if err != nil {
		t.Fatal(err)
	}
	generationResult, err := PulumiSchema(rootDir, modules, versioningStub{}, semver.MustParse("2.0.0"), false /* onlyExplicitVersions */)
	if err != nil {
		t.Fatal(err)
	}
	snaps.MatchJSON(t, generationResult.Schema)
	snaps.MatchJSON(t, generationResult.Metadata)
}
