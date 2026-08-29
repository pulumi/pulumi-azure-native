// Copyright 2023, Pulumi Corporation.  All rights reserved.

// Disable running if we've specifically selected unit tests to run as this is an integration test.
// This is easier than having to remember to explicitly tag every unit test with `//go:build unit || all`.
//go:build !unit

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/providertest"
	"github.com/pulumi/providertest/optproviderupgrade"
	"github.com/pulumi/providertest/providers"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/assertpreview"
	"github.com/pulumi/providertest/pulumitest/assertrefresh"
	"github.com/pulumi/providertest/pulumitest/changesummary"
	"github.com/pulumi/providertest/pulumitest/opttest"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/debug"
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

func TestWebAppSiteExtensions(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "site-extension")
	defer func() {
		pt.Destroy(t)
	}()
	pt.Up(t)
}

func TestTagAtScope(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "tag-at-scope")
	defer func() {
		pt.Destroy(t)
	}()

	pt.Up(t)
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
	pt := newPulumiTest(t, "parallel-subnet-creation")
	pt.Preview(t)
	pt.Up(t)
	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}

// TestKeyVaultSecretAndKeySoftDeleteRecovery verifies that re-creating a KeyVault secret/key after it has been
// deleted (and therefore soft-deleted by Azure, which enforces soft-delete on all vaults) succeeds by recovering
// the soft-deleted resource instead of failing with a conflict. See custom_keyvault.go and issues #1174, #1211.
func TestKeyVaultSecretAndKeySoftDeleteRecovery(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "keyvault-soft-delete-recovery/step1")
	defer func() {
		pt.Destroy(t)
	}()

	up := pt.Up(t)
	assertNonEmptyStringOutput(t, up, "secretUri")
	assertNonEmptyStringOutput(t, up, "keyUri")

	// Remove the secret and key. The provider deletes them via the KeyVault data plane, which soft-deletes them.
	pt.UpdateSource(t, "test-programs", "keyvault-soft-delete-recovery", "step2")
	pt.Up(t)

	// Re-introduce the secret and key with the same names. Azure still has them in a soft-deleted state, so
	// creating them again requires recovering them first; a plain ARM PUT would fail with a conflict.
	pt.UpdateSource(t, "test-programs", "keyvault-soft-delete-recovery", "step1")
	up = pt.Up(t)

	upSummary := changesummary.FromStringIntMap(*up.Summary.ResourceChanges)
	assert.Equal(t, 2, upSummary[apitype.OpCreate], "expected the secret and key to be (re-)created")

	// Real ARM-issued outputs (as opposed to nil outputs from a create that never actually called ARM) confirm
	// the resources were genuinely created, not just recovered and left stale.
	assertNonEmptyStringOutput(t, up, "secretUri")
	assertNonEmptyStringOutput(t, up, "keyUri")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))
}

func assertNonEmptyStringOutput(t *testing.T, up auto.UpResult, name string) {
	t.Helper()
	value, ok := up.Outputs[name].Value.(string)
	require.True(t, ok, "expected output %q to be a string", name)
	require.NotEmpty(t, value, "expected output %q to be non-empty", name)
}

func TestGenericResourceCreatingCongitiveServicesAccount(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "generic-resource-cognitive-services")
	upResult := pt.Up(t)
	t.Logf("std out:\n%s\n", upResult.StdOut)
	errorMsg := "Failed to read resource after Create. Please report this issue."
	assert.NotContainsf(t, upResult.StdOut, errorMsg, "Expected not to see error message '%s' in stderr", errorMsg)
}

// TestWebAppBackupConfiguration guards against issue #4408: Azure's PUT response for this resource
// returns the parent site's id rather than the backup config's own id, which used to get adopted as
// the resource's canonical ID (see resources.IgnoreResponseID) and broke the subsequent read.
func TestWebAppBackupConfiguration(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "webapp-backup-configuration")
	defer func() {
		pt.Destroy(t)
	}()

	up := pt.Up(t)
	errorMsg := "Failed to read resource after Create. Please report this issue."
	assert.NotContainsf(t, up.StdOut, errorMsg, "Expected not to see error message %q in stdout", errorMsg)

	backupConfigId, ok := up.Outputs["backupConfigId"].Value.(string)
	require.True(t, ok)
	assert.Contains(t, backupConfigId, "/config/backup")

	assertrefresh.HasNoChanges(t, pt.Refresh(t))

	// A second Up must be a no-op: the bug caused refresh/preview to always see a new create.
	up2 := pt.Up(t)
	upSummary := changesummary.FromStringIntMap(*up2.Summary.ResourceChanges)
	assert.Zero(t, upSummary[apitype.OpCreate], "expected no creates on a repeat up")
	assert.Zero(t, upSummary[apitype.OpUpdate], "expected no updates on a repeat up")
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

// Tests that we are able to delete the program that has
// a NetworkRuleSet for the service bus namespace
// which in case for Pulumi and Azure, means reverting the resource
// to its default state {"defaultAction": "Allow"} when deleting
// since it is a Singleton resource
func TestServiceBusNetworkRuleset(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "servicebus-network-ruleset")
	pt.Preview(t)
	pt.Up(t)
	pt.Destroy(t)
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

func TestDefaultAzSubscriptionProvider(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping in testing.Short() mode")
		return
	}

	// AZURE_CONFIG_DIR_FOR_TEST is set by the GH workflow build-test.yml
	// to provide an isolated configuration directory for the Azure CLI.
	configDir := os.Getenv("AZURE_CONFIG_DIR_FOR_TEST")
	if configDir == "" {
		if os.Getenv("CI") != "" {
			t.Error("CLI test without AZURE_CONFIG_DIR_FOR_TEST")
		}
		t.Skip("Skipping CLI test without AZURE_CONFIG_DIR_FOR_TEST")
	}
	t.Setenv("AZURE_CONFIG_DIR", configDir)

	ctx := context.Background()
	subscription, err := defaultAzSubscriptionProvider(ctx, os.Getenv("ARM_SUBSCRIPTION_ID"))
	assert.NoError(t, err)
	assert.NotNil(t, subscription)
}

func TestAzidentity(t *testing.T) {
	if testing.Short() {
		t.Skipf("Skipping in testing.Short() mode")
		return
	}

	validate := func(t *testing.T, up auto.UpResult) (map[string]interface{}, jwt.MapClaims) {
		// validate clientConfig
		require.Contains(t, up.Outputs, "clientConfig", "expected clientConfig output")
		clientConfig, _ := up.Outputs["clientConfig"].Value.(map[string]interface{})
		clientConfigJSON, _ := json.Marshal(clientConfig)
		t.Logf("clientConfig: %s", clientConfigJSON)

		assert.Contains(t, clientConfig, "clientId")
		assert.Contains(t, clientConfig, "objectId")
		assert.Contains(t, clientConfig, "subscriptionId")
		assert.Contains(t, clientConfig, "tenantId")

		// validate clientToken
		require.Contains(t, up.Outputs, "clientToken", "expected clientToken output")
		clientToken, _ := up.Outputs["clientToken"].Value.(map[string]interface{})
		claims, err := parseJwtUnverified(clientToken["token"].(string))
		require.NoError(t, err)
		claimsJSON, _ := json.Marshal(claims)
		t.Logf("clientToken: %s", claimsJSON)

		return clientConfig, claims
	}

	t.Run("OIDC", func(t *testing.T) {
		oidcClientId := os.Getenv("OIDC_ARM_CLIENT_ID")
		if oidcClientId == "" {
			t.Skip("Skipping OIDC test without OIDC_ARM_CLIENT_ID")
		}

		t.Setenv("ARM_USE_OIDC", "true")
		t.Setenv("ARM_CLIENT_ID", oidcClientId)
		// Make sure we test the OIDC method
		t.Setenv("ARM_CLIENT_SECRET", "")
		t.Setenv("ARM_CLIENT_CERTIFICATE_PATH", "")
		t.Setenv("ARM_CLIENT_CERTIFICATE_PASSWORD", "")

		pt := newPulumiTest(t, "azidentity")

		up := pt.Up(t)
		clientConfig, clientToken := validate(t, up)
		assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientConfig["clientId"])
		assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientToken["appid"])
		assert.Equal(t, "app", clientToken["idtyp"])
	})

	t.Run("SP_clientsecret", func(t *testing.T) {
		clientSecret := os.Getenv("ARM_CLIENT_SECRET")
		if clientSecret == "" {
			if os.Getenv("CI") != "" {
				t.Error("SP test without ARM_CLIENT_SECRET")
			}
			t.Skip("Skipping SP test without ARM_CLIENT_SECRET")
		}

		t.Setenv("ARM_CLIENT_ID", os.Getenv("ARM_CLIENT_ID"))
		t.Setenv("ARM_CLIENT_SECRET", clientSecret)
		// Make sure we test the client secret method
		t.Setenv("ARM_CLIENT_CERTIFICATE_PASSWORD", "")
		t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "")
		t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", "")

		pt := newPulumiTest(t, "azidentity")

		up := pt.Up(t)
		clientConfig, clientToken := validate(t, up)
		assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientConfig["clientId"])
		assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientToken["appid"])
		assert.Equal(t, "app", clientToken["idtyp"])
	})

	t.Run("SP_clientcert", func(t *testing.T) {
		// TOD pulumi/pulumi-azure-native#4765
		t.Skip("Skipping SP test with client certificate until we can get a valid cert for CI")

		certPath := os.Getenv("ARM_CLIENT_CERTIFICATE_PATH_FOR_TEST")
		if certPath == "" {
			if os.Getenv("CI") != "" {
				t.Error("SP test without ARM_CLIENT_CERTIFICATE_PATH_FOR_TEST")
			}
			t.Skip("Skipping SP test without ARM_CLIENT_CERTIFICATE_PATH_FOR_TEST")
		}

		t.Setenv("ARM_CLIENT_ID", os.Getenv("ARM_CLIENT_ID"))
		t.Setenv("ARM_CLIENT_CERTIFICATE_PATH", certPath)
		t.Setenv("ARM_CLIENT_CERTIFICATE_PASSWORD", os.Getenv("ARM_CLIENT_CERTIFICATE_PASSWORD_FOR_TEST"))
		// Make sure we test the client certificate method
		t.Setenv("ARM_CLIENT_SECRET", "")
		t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "")
		t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", "")

		pt := newPulumiTest(t, "azidentity")

		up := pt.Up(t)
		clientConfig, clientToken := validate(t, up)
		assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientConfig["clientId"])
		assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientToken["appid"])
		assert.Equal(t, "app", clientToken["idtyp"])
	})

	t.Run("CLI", func(t *testing.T) {
		// AZURE_CONFIG_DIR_FOR_TEST is set by the GH workflow build-test.yml
		// to provide an isolated configuration directory for the Azure CLI.
		configDir := os.Getenv("AZURE_CONFIG_DIR_FOR_TEST")
		if configDir == "" {
			if os.Getenv("CI") != "" {
				t.Error("CLI test without AZURE_CONFIG_DIR_FOR_TEST")
			}
			t.Skip("Skipping CLI test without AZURE_CONFIG_DIR_FOR_TEST")
		}
		t.Setenv("AZURE_CONFIG_DIR", configDir)

		// Make sure we test the CLI method
		t.Setenv("ARM_USE_MSI", "false")
		t.Setenv("ARM_USE_OIDC", "false")
		t.Setenv("ARM_TENANT_ID", "")
		t.Setenv("ARM_CLIENT_ID", "")
		t.Setenv("ARM_CLIENT_SECRET", "")
		t.Setenv("ARM_CLIENT_CERTIFICATE_PATH", "")
		t.Setenv("ARM_CLIENT_CERTIFICATE_PASSWORD", "")
		t.Setenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN", "")
		t.Setenv("ACTIONS_ID_TOKEN_REQUEST_URL", "")

		pt := newPulumiTest(t, "azidentity")
		up := pt.Up(t)
		clientConfig, clientToken := validate(t, up)
		// When using service principal authentication, verify we got valid credentials
		assert.NotEmpty(t, clientConfig["clientId"], "clientId should be present")
		assert.NotEmpty(t, clientToken["appid"], "appid should be present")
		// Service principal tokens have idtyp "app" (not "user")
		assert.Equal(t, "app", clientToken["idtyp"])
	})

	t.Run("Default Azure Credential", func(t *testing.T) {
		t.Setenv("ARM_USE_DEFAULT_AZURE_CREDENTIAL", "true")

		if _, ok := os.LookupEnv("CI"); ok {
			// Configure the default credential chain to use variables provided in build-test.yml, per:
			// https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity#readme-environment-variables
			t.Setenv("AZURE_TOKEN_CREDENTIALS", "EnvironmentCredential")
			t.Setenv("AZURE_CLIENT_ID", os.Getenv("ARM_CLIENT_ID"))
			t.Setenv("AZURE_TENANT_ID", os.Getenv("ARM_TENANT_ID"))
			t.Setenv("AZURE_CLIENT_SECRET", os.Getenv("ARM_CLIENT_SECRET"))

			// Ensure that a subscription ID was provided, because ADC doesn't provide one.
			require.NotEmpty(t, os.Getenv("ARM_SUBSCRIPTION_ID"))
		}

		pt := newPulumiTest(t, "azidentity")
		up := pt.Up(t)
		clientConfig, clientToken := validate(t, up)

		if _, ok := os.LookupEnv("CI"); ok {
			assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientConfig["clientId"])
			assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientToken["appid"])
			assert.Equal(t, "app", clientToken["idtyp"])
		}
	})
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

func TestUpgradeAksApiVersion_2_90_0(t *testing.T) {
	upgradeTest(t, "upgrade-aks-api-version", "2.90.0",
		// v2 uses versioned type (containerservice/v20240102preview), v3 uses unversioned with alias
		optproviderupgrade.NewSourcePath(filepath.Join("test-programs", "upgrade-aks-api-version", "v3")))
}

// TestManagedClusterNetworkProfileForceNewOnlyImmutableFields is a regression test for issue #4756:
// ManagedCluster.networkProfile used to be marked ForceNew as a whole object, so changing any
// sub-property under it forced a full AKS cluster replacement, even sub-properties that AKS supports
// updating in place (e.g. advancedNetworking/ACNS via `az aks update --enable-acns`).
//
// This only previews (rather than applies) the replacement scenario, since actually replacing a real
// AKS cluster is slow and unnecessary to validate: the change under test is purely in the provider's
// Diff logic (which properties get classified as forcing a replacement), not in the mechanics of
// applying a replacement, which are unaffected.
func TestManagedClusterNetworkProfileForceNewOnlyImmutableFields(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "aks-networkprofile-forcenew")
	defer func() {
		pt.Destroy(t)
	}()

	// Baseline: dnsServiceIP=10.0.0.10, advancedNetworking.enabled=false.
	pt.Up(t)

	// Changing only advancedNetworking (freely updatable per `az aks update`) should update in place.
	pt.SetConfig(t, "advancedNetworkingEnabled", "true")
	mutablePreview := pt.Preview(t)
	assertpreview.HasNoReplacements(t, mutablePreview)
	mutableSummary := changesummary.ChangeSummary(mutablePreview.ChangeSummary)
	assert.Greater(t, mutableSummary[apitype.OpUpdate], 0,
		"expected an in-place update when only advancedNetworking changes")

	// Revert advancedNetworking so only dnsServiceIP (genuinely immutable, no `az aks update`
	// equivalent) differs from the deployed baseline. That alone must force a replacement.
	pt.SetConfig(t, "advancedNetworkingEnabled", "false")
	pt.SetConfig(t, "dnsServiceIP", "10.0.1.10")
	immutablePreview := pt.Preview(t)
	immutableSummary := changesummary.ChangeSummary(immutablePreview.ChangeSummary)
	replacementOps := immutableSummary[apitype.OpReplace] +
		immutableSummary[apitype.OpCreateReplacement] +
		immutableSummary[apitype.OpDeleteReplaced]
	assert.Greater(t, replacementOps, 0,
		"expected a replacement when changing the immutable dnsServiceIP sub-property")
}

// TestDatabricksWorkspaceComputeModeDefault is a regression test for issue #4766: Databricks
// Workspace's computeMode became a required input in API version 2026-01-01. Confirmed live
// against Azure (raw ARM REST PUT, an ARM template deployment, and reading a workspace created
// via an older API version that predates the property) that omitting computeMode on create is
// accepted and Azure backfills "Hybrid" server-side, so the provider keeps it as an optional
// input (see notRequiredInputsMap in provider/pkg/gen/replacement.go) instead of a required one.
//
// This also verifies that changing computeMode away from its backfilled value forces a
// replacement, since the API documents it as "Required on create, cannot be changed" (see
// forceNewMap in the same file). Like TestManagedClusterNetworkProfileForceNewOnlyImmutableFields
// above, the replacement step only previews (never applies) since actually replacing a
// Databricks workspace is slow and unnecessary to validate the Diff logic under test.
//
// NOTE: explicitly (re-)specifying computeMode="Hybrid" after adopting the backfilled default
// currently still shows up as a one-time input diff (the provider diffs old-vs-new *inputs*, not
// against the live server value, and this affects Read/refresh too -- not just Diff/preview).
// Suppressing that spurious diff is being tracked as separate follow-up work since it's a
// broader pattern that shows up for other optional/server-defaulted properties too, not just this
// one; this test only asserts the parts of the issue #4766 fix that are actually in scope here:
// the create succeeds without computeMode, Azure backfills "Hybrid", and re-pinning it to that
// same value never forces a replacement.
func TestDatabricksWorkspaceComputeModeDefault(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "databricks-computemode/step1")
	defer func() {
		pt.Destroy(t)
	}()

	// step1: create a workspace without setting computeMode at all.
	up := pt.Up(t)
	computeMode, ok := up.Outputs["computeMode"].Value.(string)
	require.True(t, ok, "expected computeMode output to be a string")
	assert.Equal(t, "Hybrid", computeMode, "expected Azure to backfill computeMode to Hybrid when omitted")
	// Refresh should succeed without error, but may report the computeMode input being adopted
	// into state as a one-time change -- see the NOTE above, not asserted here.
	pt.Refresh(t)

	// step2: explicitly set computeMode to "Hybrid" -- the value Azure already backfilled. This
	// may still show up as a one-time update (see NOTE above), but must never be a replacement.
	pt.UpdateSource(t, "test-programs", "databricks-computemode", "step2")
	hybridPreview := pt.Preview(t)
	assertpreview.HasNoReplacements(t, hybridPreview)

	// step3: change computeMode to "Serverless". Per the API docs this is immutable after
	// create, so this must force a replacement. Preview only -- see comment above.
	pt.UpdateSource(t, "test-programs", "databricks-computemode", "step3")
	serverlessPreview := pt.Preview(t)
	serverlessSummary := changesummary.ChangeSummary(serverlessPreview.ChangeSummary)
	replacementOpsServerless := serverlessSummary[apitype.OpReplace] +
		serverlessSummary[apitype.OpCreateReplacement] +
		serverlessSummary[apitype.OpDeleteReplaced]
	assert.Greater(t, replacementOpsServerless, 0,
		"expected a replacement when changing computeMode from Hybrid to Serverless")
}

// TestSecurityInsightsWatchlistDelete is a regression test for issue #4816: Azure's delete
// status monitor for a Watchlist never reports a terminal state, even though the resource
// itself is deleted within seconds. Without the GET-based fallback this hangs until the
// provider's delete timeout. See Azure/azure-sdk-for-go#26937 for the confirmed upstream bug.
func TestSecurityInsightsWatchlistDelete(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "watchlist-delete/step1")
	defer func() {
		pt.Destroy(t)
	}()
	pt.Up(t)

	pt.UpdateSource(t, "test-programs", "watchlist-delete", "step2")
	start := time.Now()
	up := pt.Up(t)
	elapsed := time.Since(start)

	upSummary := changesummary.FromStringIntMap(*up.Summary.ResourceChanges)
	assert.Equal(t, 1, upSummary[apitype.OpDelete], "expected the watchlist to be deleted")
	assert.Less(t, elapsed, 10*time.Minute,
		"delete took too long, the GET-based fallback may not be engaging")
}

// TestStorageAccountSingletonChildAfterReplace is a regression test for issue #4738: replacing a
// StorageAccount (e.g. changing the immutable isHnsEnabled property) cascades into a
// delete-then-create of its FileServiceProperties singleton child. That child's Create call used
// to run an existence check that could transiently fail with 400 AuthenticationFailed while
// ARM's auth context propagated for the newly recreated account. CanCreate now skips that check
// for singleton resources, since a singleton always exists once its parent does.
func TestStorageAccountSingletonChildAfterReplace(t *testing.T) {
	t.Parallel()
	pt := newPulumiTest(t, "storage-sa-replace-singleton-child/step1")
	defer func() {
		pt.Destroy(t)
	}()

	pt.Up(t)

	// Removing isHnsEnabled forces the replace; pt.Up fails the test on any error.
	pt.UpdateSource(t, "test-programs", "storage-sa-replace-singleton-child", "step2")
	pt.Up(t)
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

func parseJwtUnverified(tokenString string) (jwt.MapClaims, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims, nil
}

func debugLogging() debug.LoggingOptions {
	var level uint = 11
	return debug.LoggingOptions{
		LogLevel:      &level,
		Debug:         true,
		FlowToPlugins: true,
		LogToStdErr:   true,
	}
}
