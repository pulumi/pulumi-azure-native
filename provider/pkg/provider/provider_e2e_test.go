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
	"github.com/stretchr/testify/assert"

	"github.com/pulumi/providertest"
	"github.com/pulumi/providertest/optproviderupgrade"
	"github.com/pulumi/providertest/providers"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/assertrefresh"
	"github.com/pulumi/providertest/pulumitest/changesummary"
	"github.com/pulumi/providertest/pulumitest/opttest"

	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

var schemaBytes []byte
var azureAPIResourcesBytes []byte

func init() {
	var err error
	schemaBytes, err = os.ReadFile(filepath.Join("..", "..", "..", "bin", "schema-full.json"))
	if err != nil {
		fmt.Printf("failed to read schema file, run `make schema` before running tests: %v", err)
	}

	azureAPIResourcesBytes, err = os.ReadFile(filepath.Join("..", "..", "..", "bin", "metadata-compact.json"))
	if err != nil {
		fmt.Printf("failed to read metadata file, run `make schema` before running tests: %v", err)
	}
}

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

func TestSubResources(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "subresources")

	// deploy an NSG with an "external" security rule, and an NSG with an inline security rule.
	up := pt.Up(t)
	assert.Len(t, up.Outputs["external-nsg-security-rules"].Value, 0)
	assert.Len(t, up.Outputs["inline-nsg-security-rules"].Value, 1)
	inlineRule := up.Outputs["inline-nsg-security-rules"].Value.([]any)[0].(map[string]any)
	assert.Equal(t, "inline", inlineRule["name"])

	// update a tag on the NSGs, and then check that the external security rules are now available as outputs.
	pt.SetConfig(t, "subresources:revision", "2")
	up = pt.Up(t)
	upSummary := changesummary.FromStringIntMap(*up.Summary.ResourceChanges)
	assert.Equal(t, 2, upSummary[apitype.OpUpdate])
	assert.Len(t, up.Outputs["inline-nsg-security-rules"].Value, 1)
	assert.Len(t, up.Outputs["external-nsg-security-rules"].Value, 1)
	externalRule := up.Outputs["external-nsg-security-rules"].Value.([]any)[0].(map[string]any)
	assert.Equal(t, "external", externalRule["name"])

	// check that the state is stable after a refresh.
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

func TestAutonaming(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "autonaming", opttest.Env("PULUMI_EXPERIMENTAL", "1"))
	pt.Preview(t)
	up := pt.Up(t)
	rgname, ok := up.Outputs["rgname"].Value.(string)
	assert.True(t, ok)
	assert.Contains(t, rgname, "autonaming-rg-") // project + name + random suffix
	saname, ok := up.Outputs["saname"].Value.(string)
	assert.True(t, ok)
	assert.Contains(t, saname, "autonamingsa") // project + name + random suffix, no dashes
}

func TestTagging(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "tagging")
	pt.Preview(t)
	up := pt.Up(t)

	// Verify that the tags were truly applied to the resource groups.
	// One must first refresh the state to see the tags applied by the TagAtScope resource.
	up = pt.Up(t, optup.Refresh())
	rg1Tags, _ := up.Outputs["rg_1_tags"].Value.(map[string]any)
	assert.Equal(t, map[string]any{"owner": "tag_1"}, rg1Tags)
	rg2Tags, _ := up.Outputs["rg_2_tags"].Value.(map[string]any)
	assert.Equal(t, map[string]any{"owner": "tag_2"}, rg2Tags)
}

func TestUpgradeKeyVault_2_76_0(t *testing.T) {
	upgradeTest(t, "upgrade-keyvault", "2.76.0")
}

func TestUpgradeNetworkedVm_2_76_0(t *testing.T) {
	upgradeTest(t, "upgrade-networked-vm", "2.76.0")
}

func TestUpgradeStorageBlob_2_76_0(t *testing.T) {
	upgradeTest(t, "upgrade-storage-blob", "2.76.0")
}

func TestUpgradeSqlDatabase_2_76_0(t *testing.T) {
	upgradeTest(t, "upgrade-sql-database", "2.76.0")
}

func TestUpgradeServiceBusMessaging_2_76_0(t *testing.T) {
	upgradeTest(t, "upgrade-servicebus-messaging", "2.76.0")
}

func TestUpgradeAppServicesWebApp_2_76_0(t *testing.T) {
	upgradeTest(t, "upgrade-appservices-webapp", "2.76.0")
}

func TestUpgradeCosmosdbNosql_2_90_0(t *testing.T) {
	upgradeTest(t, "upgrade-cosmosdb-nosql", "2.90.0",
		// DocumentDB was renamed to CosmosDB in v3
		optproviderupgrade.NewSourcePath(filepath.Join("test-programs", "upgrade-cosmosdb-nosql", "v3-cosmosdb")))
}

func TestUpgradeContainerServiceAgentPool_2_90_0(t *testing.T) {
	upgradeTest(t, "upgrade-containerservice-agentpool", "2.90.0")
}

func upgradeTest(t *testing.T, testProgramDir string, upgradeFromVersion string, opts ...optproviderupgrade.PreviewProviderUpgradeOpt) {
	t.Helper()
	if testing.Short() {
		t.Skipf("Skipping in testing.Short() mode, assuming this is a CI run without cloud credentials")
		return
	}

	dir := filepath.Join("test-programs", testProgramDir)
	azureLocation := getLocation()
	rpFactory := providers.ResourceProviderFactory(providerServer)
	cacheDir := providertest.GetUpgradeCacheDir(filepath.Base(dir), upgradeFromVersion)
	pt := pulumitest.NewPulumiTest(t, dir,
		opttest.AttachProvider("azure-native",
			rpFactory.ReplayInvokes(filepath.Join(cacheDir, "grpc.json"), false /* allowLiveFallback */)))
	pt.SetConfig(t, "azure-native:location", azureLocation)
	previewResult := providertest.PreviewProviderUpgrade(t, pt, "azure-native", upgradeFromVersion, opts...)
	assertpreview.HasNoReplacements(t, previewResult)
	assertpreview.HasNoDeletes(t, previewResult)
}

func newPulumiTest(t *testing.T, testProgramDir string, opts ...opttest.Option) *pulumitest.PulumiTest {
	t.Helper()
	if testing.Short() {
		t.Skipf("Skipping in testing.Short() mode, assuming this is a CI run without cloud credentials")
		return nil
	}
	dir := filepath.Join("test-programs", testProgramDir)
	azureLocation := getLocation()
	rpFactory := providers.ResourceProviderFactory(providerServer)
	attachOpt := opttest.AttachProvider("azure-native", rpFactory)
	pt := pulumitest.NewPulumiTest(t, dir, append(opts, attachOpt)...)
	pt.SetConfig(t, "azure-native:location", azureLocation)
	return pt
}

func providerServer(_ providers.PulumiTest) (pulumirpc.ResourceProviderServer, error) {
	version.Version = os.Getenv("PROVIDER_VERSION")
	if version.Version == "" {
		version.Version = "3.0.0"
	}
	if len(schemaBytes) == 0 {
		return nil, fmt.Errorf("schema not loaded")
	}
	if len(azureAPIResourcesBytes) == 0 {
		return nil, fmt.Errorf("azure API resources not loaded")
	}

	return makeProvider(nil, "azure-native", version.GetVersion().String(), schemaBytes, azureAPIResourcesBytes)
}

func getLocation() string {
	azureLocation := os.Getenv("ARM_LOCATION")
	if azureLocation == "" {
		azureLocation = "westus2"
		fmt.Println("Defaulting location to 'westus2'. You can override using the ARM_LOCATION variable.")
	}

	return azureLocation
}
