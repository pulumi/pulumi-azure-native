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

	"github.com/golang-jwt/jwt"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
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
			if os.Getenv("CI") != "" {
				t.Error("OIDC test without OIDC_ARM_CLIENT_ID")
			}
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
		assert.Equal(t, "04b07795-8ddb-461a-bbee-02f9e1bf7b46", clientConfig["clientId"])
		assert.Equal(t, "04b07795-8ddb-461a-bbee-02f9e1bf7b46", clientToken["appid"])
		assert.Equal(t, "user", clientToken["idtyp"])
	})

	t.Run("Default Azure Credential", func(t *testing.T) {
		t.Setenv("ARM_USE_DEFAULT_AZURE_CREDENTIAL", "true")

		if _, ok := os.LookupEnv("CI"); ok {
			// Configure the default credential chain to use variables provided in build-test.yml, per:
			// https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity#readme-environment-variables
			t.Setenv("AZURE_CLIENT_ID", os.Getenv("ARM_CLIENT_ID"))
			t.Setenv("AZURE_TENANT_ID", os.Getenv("ARM_TENANT_ID"))
			t.Setenv("AZURE_CLIENT_SECRET", os.Getenv("ARM_CLIENT_SECRET"))
		}

		pt := newPulumiTest(t, "azidentity")
		up := pt.Up(t, optup.DebugLogging(debugLogging()))
		clientConfig, clientToken := validate(t, up)

		if _, ok := os.LookupEnv("CI"); ok {
			assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientConfig["clientId"])
			assert.Equal(t, os.Getenv("ARM_CLIENT_ID"), clientToken["appid"])
			assert.Equal(t, "app", clientToken["idtyp"])
		}
	})
}

func TestAutorest(t *testing.T) {
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
		t.Setenv("PULUMI_ENABLE_AZCORE_BACKEND", "false")

		oidcClientId := os.Getenv("OIDC_ARM_CLIENT_ID")
		if oidcClientId == "" {
			if os.Getenv("CI") != "" {
				t.Error("OIDC test without OIDC_ARM_CLIENT_ID")
			}
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
		t.Setenv("PULUMI_ENABLE_AZCORE_BACKEND", "false")

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
		t.Setenv("PULUMI_ENABLE_AZCORE_BACKEND", "false")

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
		t.Setenv("PULUMI_ENABLE_AZCORE_BACKEND", "false")

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
		assert.Equal(t, "04b07795-8ddb-461a-bbee-02f9e1bf7b46", clientConfig["clientId"])
		assert.Equal(t, "04b07795-8ddb-461a-bbee-02f9e1bf7b46", clientToken["appid"])
		assert.Equal(t, "user", clientToken["idtyp"])
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
