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

func TestPortalDashboardGen(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	// We take a snapshot of the schema as we're focusing on testing how we parse and modify the spec
	// To update the specifications run in the repo root:
	// rm -rf provider/pkg/gen/test-data/dashboard/azure-rest-api-specs/specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/
	// rm -rf provider/pkg/gen/test-data/dashboard/azure-rest-api-specs/specification/common-types/resource-management/v5/
	// cp -r azure-rest-api-specs/specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/ provider/pkg/gen/test-data/dashboard/azure-rest-api-specs/specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/
	// cp -r azure-rest-api-specs/specification/common-types/resource-management/v5/ provider/pkg/gen/test-data/dashboard/azure-rest-api-specs/specification/common-types/resource-management/v5/
	// rm -rf provider/pkg/gen/test-data/dashboard/azure-rest-api-specs/specification/portal/resource-manager/Microsoft.Portal/preview/2020-09-01-preview/examples/
	rootDir := path.Join(wd, "test-data", "dashboard")

	modules, _, err := openapi.ReadAzureModules(path.Join(rootDir, "azure-rest-api-specs"), "Portal", "2020-09-01-preview")
	if err != nil {
		t.Fatal(err)
	}
	modules = openapi.ApplyTransformations(modules, openapi.DefaultVersions{
		"Portal": {
			"Dashboard": {
				ApiVersion: "2020-09-01-preview",
			},
		},
	}, openapi.DefaultVersions{}, nil, nil)
	versioning := versioningStub{
		// These are extracted from versions/v2-token-paths.json
		previousTokenPaths: map[string]string{
			"azure-native:portal/v20190101preview:Dashboard": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}",
			"azure-native:portal/v20200901preview:Dashboard": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}",
			"azure-native:portal/v20221201preview:Dashboard": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}",
			"azure-native:portal:Dashboard":                  "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Portal/dashboards/{dashboardName}",
		},
	}
	generationResult, err := PulumiSchema(rootDir, modules, versioning, semver.MustParse("3.0.0"), false /* onlyExplicitVersions */)
	if err != nil {
		t.Fatal(err)
	}

	// Ensure the dashboard resource is present in schema and metadata
	// and snapshot the generation result so we can see the impact of future refactors.
	assert.NotNil(t, generationResult.Schema.Resources["azure-native:portal:Dashboard"])
	snaps.MatchJSON(t, generationResult.Schema.Resources["azure-native:portal:Dashboard"])

	assert.NotNil(t, generationResult.Metadata.Resources["azure-native:portal:Dashboard"])
	snaps.MatchJSON(t, generationResult.Metadata.Resources["azure-native:portal:Dashboard"])

	snaps.MatchJSON(t, generationResult.Schema.Types)
	snaps.MatchJSON(t, generationResult.Metadata.Types)
}
