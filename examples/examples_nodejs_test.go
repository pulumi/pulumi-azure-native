// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.
//go:build nodejs || all

package examples

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccAppServiceTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "appservice"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccCosmosDBTs(t *testing.T) {
	t.Skip("Skipping due to CosmosDB failing with ServiceUnavailable due to high demand")
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "cosmosdb"),
		})

	integration.ProgramTest(t, &test)
}

func TestAccSimpleTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "simple"),
		})

	integration.ProgramTest(t, &test)
}

func TestStaticWebsiteDestroyTs(t *testing.T) {
	// Tests eventually-consistent deletion handling.
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "static-website"),
		})

	integration.ProgramTest(t, &test)
}

func TestImportTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "import"),
		})

	integration.ProgramTest(t, &test)
}

func TestPostgresTs(t *testing.T) {
	t.Skip("disabled due to a failure, see https://github.com/pulumi/pulumi-azure-native/issues/898")
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "postgres"),
		})

	integration.ProgramTest(t, &test)
}

func TestMySqlTs(t *testing.T) {
	t.Skip("looks unreliable, getting errors with Code=ResourceNotFound Code=InternalServerError")
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "mysql"),
		})

	integration.ProgramTest(t, &test)
}

func TestMessagingTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "messaging"),
		})

	integration.ProgramTest(t, &test)
}

func TestSecretsTs(t *testing.T) {
	secretMessage := "secret message for testing"

	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "secrets"),
			Config: map[string]string{
				"message": secretMessage,
			},
			ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
				assert.NotNil(t, stackInfo.Deployment)
				state, err := json.Marshal(stackInfo.Deployment)
				assert.NoError(t, err)

				assert.NotContains(t, string(state), secretMessage)
			},
		})

	integration.ProgramTest(t, &test)
}

func TestTimeSeriesTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "timeseries"),
		})

	integration.ProgramTest(t, &test)
}

func TestPublicIpUpdateTs(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:         filepath.Join(getCwd(t), "public-ip-update"),
			SkipRefresh: true,
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join("public-ip-update", "step2"),
					Additive: true,
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestVnetSubnetsResolution(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "vnet-subnets-resolution"),
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join("vnet-subnets-resolution", "step2"),
					Additive: true,
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestStorageAccountNetworkRule(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "storageaccount-networkrule"),
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join("storageaccount-networkrule", "step2"),
					Additive: true,
					ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
						assert.NotNil(t, stackInfo.Deployment)
						assert.NotNil(t, stackInfo.Deployment.Resources)

						found := false
						for _, resource := range stackInfo.Deployment.Resources {
							if resource.Type == "azure-native:storage:StorageAccount" {
								found = true

								networkRuleSet, ok := resource.Outputs["networkRuleSet"]
								assert.True(t, ok)
								networkRuleSetMap, ok := networkRuleSet.(map[string]interface{})
								assert.True(t, ok)

								assert.Equal(t, "Allow", networkRuleSetMap["defaultAction"])
								assert.Empty(t, networkRuleSetMap["ipRules"])

								break
							}
						}
						assert.True(t, found, "no storage account found in deployed resources")
					},
				},
			},
			Verbose: true,
		})

	integration.ProgramTest(t, &test)
}

func TestAccKeyVaultAccessPoliciesTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			ExpectRefreshChanges: false,
			Dir:                  filepath.Join(getCwd(t), "keyvault-accesspolicies"),
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join("keyvault-accesspolicies", "2-update-keyvault"),
					Additive: true,
					// Check that the stand-alone AccessPolicies are still there, not deleted by the Vault update.
					ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
						require.NotNil(t, stackInfo.Deployment)
						require.NotNil(t, stackInfo.Deployment.Resources)

						accessPolicies := 0
						for _, resource := range stackInfo.Deployment.Resources {
							if resource.Type == "azure-native:keyvault:AccessPolicy" {
								accessPolicies++
							}
						}
						assert.Equal(t, 2, accessPolicies)

						// check the number of policies as returned by Azure directly via invoke
						numberOfAPs, ok := stackInfo.Outputs["numberOfAPs"].(float64)
						require.True(t, ok)
						assert.Equal(t, 2.0, numberOfAPs)
					},
				},
				{
					Dir:      filepath.Join("keyvault-accesspolicies", "3-update-accesspolicies"),
					Additive: true,
					// Check that the stand-alone AccessPolicies were updated resp. deleted.
					ExtraRuntimeValidation: func(t *testing.T, stackInfo integration.RuntimeValidationStackInfo) {
						require.NotNil(t, stackInfo.Deployment)
						require.NotNil(t, stackInfo.Deployment.Resources)

						ap1Found := false
						for _, resource := range stackInfo.Deployment.Resources {
							urn := string(resource.URN)
							if strings.HasSuffix(urn, "keyvault:AccessPolicy::ap1") {
								ap1Found = true
								accessPolicy, ok := resource.Outputs["policy"]
								require.True(t, ok, "Property 'policy' not found")
								accessPolicyObj, ok := accessPolicy.(map[string]interface{})
								require.True(t, ok, "Property 'policy' is not an object")

								permissions, ok := accessPolicyObj["permissions"]
								require.True(t, ok, "Property 'policy.permissions' not found")
								permissionsObj, ok := permissions.(map[string]interface{})
								require.True(t, ok, "Property 'policy.permissions' is not an object")

								keyPermissions, ok := permissionsObj["keys"]
								require.True(t, ok, "Property 'policy.permissions.keys' not found")
								keyPermissionsArray, ok := keyPermissions.([]any)
								require.True(t, ok, "Property 'policy.permissions.keys' is not an array")

								require.Equal(t, 1, len(keyPermissionsArray))
								assert.Equal(t, "get", keyPermissionsArray[0].(string))
							} else if strings.HasSuffix(urn, "keyvault:AccessPolicy::ap2") {
								t.Errorf("AccessPolicy ap2 should have been deleted")
							}
						}
						assert.True(t, ap1Found, "AccessPolicy ap1 not found")

						// Check the number of policies as returned by Azure directly via invoke.
						// This doesn't work here because we have no way of waiting for the deletion of ap2.
						// numberOfAPs, ok := stackInfo.Outputs["numberOfAPs"].(float64)
						// assert.True(t, ok)
						// assert.Equal(t, 1.0, numberOfAPs)
					},
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestAccBlobContainerLegalHold(t *testing.T) {
	skipIfShort(t)
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir:                  filepath.Join(getCwd(t), "blobcontainer-legalhold"),
			ExpectRefreshChanges: false,
			EditDirs: []integration.EditDir{
				{
					Dir:      filepath.Join("blobcontainer-legalhold", "2-update-legalhold"),
					Additive: true,
				},
			},
		})

	integration.ProgramTest(t, &test)
}

func TestPortalDashboardTs(t *testing.T) {
	test := getJSBaseOptions(t).
		With(integration.ProgramTestOptions{
			Dir: filepath.Join(getCwd(t), "dashboard"),
		})

	integration.ProgramTest(t, &test)
}

func getJSBaseOptions(t *testing.T) integration.ProgramTestOptions {
	base := getBaseOptions(t)
	baseJS := base.With(integration.ProgramTestOptions{
		Dependencies: []string{
			"@pulumi/azure-native",
		},
	})

	return baseJS
}
