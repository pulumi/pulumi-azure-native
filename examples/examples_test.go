// Copyright 2021, Pulumi Corporation.  All rights reserved.

package examples

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/apitype"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

const pulumiExamplesPath = "../p-examples"

func getLocation(t *testing.T) string {
	azureLocation := os.Getenv("ARM_LOCATION")
	if azureLocation == "" {
		azureLocation = "westus2"
		fmt.Println("Defaulting location to 'westus2'. You can override using the ARM_LOCATION variable.")
	}

	return azureLocation
}

func azureNativeBinaryDir(t *testing.T) string {
	binPath, err := filepath.Abs("../bin")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Using azure-native binary from %s", binPath)
	files, err := os.ReadDir(binPath)
	if err != nil {
		t.Fatalf("failed to read directory %s: %v", binPath, err)
	}

	found := false
	for _, file := range files {
		if file.Name() == "pulumi-resource-azure-native" || file.Name() == "pulumi-resource-azure-native.exe" {
			found = true
			break
		}
	}

	if !found {
		t.Fatalf("could not find pulumi-resource-azure-native binary in %s, make sure to build the provider before running the tests", binPath)
	}

	return binPath
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	azureLocation := getLocation(t)
	binPath, err := filepath.Abs("../bin")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Using binPath %s\n", binPath)
	return integration.ProgramTestOptions{
		ExpectRefreshChanges:            true,
		RequireEmptyPreviewAfterRefresh: true,
		Config: map[string]string{
			"azure-native:location": azureLocation,
		},
		LocalProviders: []integration.LocalDependency{
			{
				Package: "azure-native",
				Path:    binPath,
			},
		},
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func skipIfShort(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}
}

func createTest(t *testing.T, source string, options ...opttest.Option) *pulumitest.PulumiTest {
	opts := []opttest.Option{
		opttest.LocalProviderPath("azure-native", azureNativeBinaryDir(t)),
	}
	opts = append(opts, options...)
	pt := pulumitest.NewPulumiTest(t,
		source,
		opts...)

	pt.SetConfig(t, "azure-native:location", getLocation(t))
	return pt
}

func azureCredentials(t *testing.T) azcore.TokenCredential {
	if os.Getenv("CI") == "" {
		// local test, use CLI credentials which assumers the user is logged in via `az login`
		cred, err := azidentity.NewAzureCLICredential(nil)
		require.NoError(t, err)
		return cred
	}
	cred, err := azidentity.NewClientSecretCredential(
		os.Getenv("ARM_TENANT_ID"),
		os.Getenv("ARM_CLIENT_ID"),
		os.Getenv("ARM_CLIENT_SECRET"),
		nil)
	require.NoError(t, err)
	return cred
}

func serviceBusNamespaceClient(t *testing.T) *armservicebus.NamespacesClient {
	tokenCred := azureCredentials(t)
	sub := os.Getenv("ARM_SUBSCRIPTION_ID")
	client, err := armservicebus.NewNamespacesClient(sub, tokenCred, nil)
	require.NoError(t, err)
	return client
}

func TestServiceBusRefreshAfterNamespaceDeletion(t *testing.T) {
	source := filepath.Join(getCwd(t), "servicebus-refresh-after-namespace-deletion")
	pt := createTest(t, source)
	upResult := pt.Up(t)

	defer func() {
		// clean up left over resources which should only be the resource group
		pt.Destroy(t)
	}()

	serviceBusNamespace := ""
	resourceGroupName := ""
	for key, value := range upResult.Outputs {
		if key == "namespaceName" {
			serviceBusNamespace = value.Value.(string)
		}
		if key == "resourceGroupName" {
			resourceGroupName = value.Value.(string)
		}
	}

	assert.NotEmpty(t, serviceBusNamespace, "namespace should not be empty")
	assert.NotEmpty(t, resourceGroupName, "resource group name should not be empty")

	t.Logf("resource group name: %s", resourceGroupName)
	t.Logf("service bus namespace: %s", serviceBusNamespace)

	// delete the service bus namespace via the API
	// then assert that a pulumi refresh works (resources deleted from the stack)

	client := serviceBusNamespaceClient(t)
	ctx := context.Background()
	poller, err := client.BeginDelete(ctx, resourceGroupName, serviceBusNamespace, nil)
	assert.NoError(t, err, "begin deleting namespace should not error")
	_, err = poller.PollUntilDone(ctx, nil)
	assert.NoError(t, err, "polling deletion until done should not error")

	t.Log("Finished deleting the service bus")
	t.Log("Attempting to refresh...")
	// Deleting the namespace is done, now refresh the stack
	pt.Refresh(t)
	t.Log("Refresh done, attempting to destroy leftover resources (resource group)")
}

func keyVaultClient(t *testing.T) *armkeyvault.VaultsClient {
	tokenCred := azureCredentials(t)
	sub := os.Getenv("ARM_SUBSCRIPTION_ID")
	client, err := armkeyvault.NewVaultsClient(sub, tokenCred, nil)
	require.NoError(t, err)
	return client
}

func tagsClient(t *testing.T) *armresources.TagsClient {
	sub := os.Getenv("ARM_SUBSCRIPTION_ID")
	client, err := armresources.NewTagsClient(sub, azureCredentials(t), nil)
	require.NoError(t, err)
	return client
}

// Assert that a RoleAssignment is deleted when performing a pulumi refresh
// after the KeyVault that contains the RoleAssignment is deleted outside of pulumi.
func TestRoleAssignmentRefreshAfterVaultDeletion(t *testing.T) {
	source := filepath.Join(getCwd(t), "keyvault-role-assignment-refresh-after-deletion")
	pt := createTest(t, source)
	upResult := pt.Up(t)

	defer func() {
		// clean up left over resources which should only be the resource group
		pt.Destroy(t)
	}()

	vaultName := ""
	resourceGroupName := ""
	for key, value := range upResult.Outputs {
		if key == "vaultName" {
			vaultName = value.Value.(string)
		}
		if key == "resourceGroupName" {
			resourceGroupName = value.Value.(string)
		}
	}

	assert.NotEmpty(t, vaultName, "key vault name should not be empty")
	assert.NotEmpty(t, resourceGroupName, "resource group name should not be empty")

	// delete the key vault via the API
	client := keyVaultClient(t)
	ctx := context.Background()
	_, err := client.Delete(ctx, resourceGroupName, vaultName, nil)
	assert.NoError(t, err, "begin deleting vault should not error")

	t.Log("Finished deleting the vault")
	t.Log("Attempting to refresh...")
	// Deleting the vault is done, now refresh the stack
	result, err := pt.CurrentStack().Refresh(ctx)
	assert.NoError(t, err, "refresh should not error")
	assert.Empty(t, result.StdErr, "refresh should not have any errors in stderr")
}

func updatePulumiYAML(t *testing.T, workDir string, contents map[string]any) {
	contentsInYAML, err := yaml.Marshal(contents)
	require.NoError(t, err, "marshalling contents to YAML should not error")
	yamlPath := filepath.Join(workDir, "Pulumi.yaml")
	err = os.WriteFile(yamlPath, contentsInYAML, 0644)
	require.NoError(t, err, "writing updated Pulumi.yaml should not error")
}

type tempProj struct {
	dir  string
	name string
}

func tempProject(t *testing.T) tempProj {
	tempDir := t.TempDir()
	projectName := fmt.Sprintf("empty-program-%s", filepath.Base(tempDir))
	updatePulumiYAML(t, tempDir, map[string]any{
		"name":    projectName,
		"runtime": "yaml",
	})
	return tempProj{
		dir:  tempDir,
		name: projectName,
	}
}

// When in CI, only run YAML tests when RUNNING_YAML_TESTS is set to true, since YAML tests are long-running and we don't want them to run on every PR by default.
// otherwise when testing locally, run the test as is.
func skipIfNotYamlInCI(t *testing.T) {
	runningInCI := os.Getenv("CI") != ""
	runningYamlTests := os.Getenv("RUNNING_YAML_TESTS")
	if runningInCI && runningYamlTests != "true" {
		t.Skip("Skipping YAML tests in CI since RUNNING_YAML_TESTS is not set to true")
	}
}

// This is a test to verify that adding the isHnsEnabled property to a storage account
// does not cause replacements, since it defaults to false.
// in fact, adding a property that defaults to false should not cause any changes at all since
// the default value is the same as what we are setting it to.
func TestAddingHnsEnabledToStorageAccountDoesNotCauseReplacements_YAML(t *testing.T) {
	skipIfNotYamlInCI(t)
	proj := tempProject(t)
	azureBinaryDir := azureNativeBinaryDir(t)
	test := createTest(t, proj.dir)

	plugins := map[string]any{
		"providers": []interface{}{
			map[string]any{
				"name": "azure-native",
				"path": azureBinaryDir,
			},
		},
	}

	storageAccountProperties := map[string]any{
		"resourceGroupName": "${resourcegroup.name}",
		"location":          "${resourcegroup.location}",
		"kind":              "StorageV2",
		"sku": map[string]any{
			"name": "Standard_LRS",
		},
	}

	storageAccount := map[string]any{
		"type":       "azure-native:storage:StorageAccount",
		"properties": storageAccountProperties,
	}

	program := map[string]any{
		"name":    proj.name,
		"runtime": "yaml",
		"resources": map[string]any{
			"resourcegroup": map[string]any{
				"type": "azure-native:resources:ResourceGroup",
			},
			"storageaccount": storageAccount,
		},
		"plugins": plugins,
	}

	// initial program
	updatePulumiYAML(t, test.WorkingDir(), program)

	// Deploy the storage account
	upResult := test.Up(t)
	assert.Empty(t, upResult.StdErr, "up should not have any errors")
	defer test.Destroy(t)

	// Update the storage account to add the IsHnsEnabled property
	storageAccountProperties["isHnsEnabled"] = false

	// Update the program with the new property
	updatePulumiYAML(t, test.WorkingDir(), program)

	preview := test.Preview(t)
	t.Logf("Preview STDOUT: \n%s", preview.StdOut)
	assert.Equal(t, map[apitype.OpType]int{
		// nothing has changed
		apitype.OpSame: 3,
	}, preview.ChangeSummary)
}

// TestListSubscriptionsInvoke verifies that the listSubscriptions invoke works
// by calling it from a YAML program and putting the results in an output.
func TestListSubscriptionsInvoke_YAML(t *testing.T) {
	skipIfNotYamlInCI(t)
	proj := tempProject(t)
	azureBinaryDir := azureNativeBinaryDir(t)
	test := createTest(t, proj.dir)

	plugins := map[string]any{
		"providers": []interface{}{
			map[string]any{
				"name": "azure-native",
				"path": azureBinaryDir,
			},
		},
	}

	program := map[string]any{
		"name":    proj.name,
		"runtime": "yaml",
		"variables": map[string]any{
			"subscriptions": map[string]any{
				"fn::invoke": map[string]any{
					"function": "azure-native:authorization:listSubscriptions",
				},
			},
		},
		"outputs": map[string]any{
			"firstSubscriptionId": "${subscriptions.value[0].subscriptionId}",
		},
		"plugins": plugins,
	}

	updatePulumiYAML(t, test.WorkingDir(), program)

	upResult := test.Up(t)
	t.Logf("Up STDOUT: \n%s", upResult.StdOut)

	firstSubscriptionId := ""
	for key, value := range upResult.Outputs {
		if key == "firstSubscriptionId" {
			firstSubscriptionId = value.Value.(string)
		}
	}

	assert.NotEmpty(t, firstSubscriptionId, "firstSubscriptionId should not be empty")
	t.Logf("First subscription ID: %s", firstSubscriptionId)
}

func TestSqlServerWithPrivateEndpoint_YAML(t *testing.T) {
	skipIfNotYamlInCI(t)
	proj := tempProject(t)
	azureBinaryDir := azureNativeBinaryDir(t)
	test := createTest(t, proj.dir)

	plugins := map[string]any{
		"providers": []interface{}{
			map[string]any{
				"name": "azure-native",
				"path": azureBinaryDir,
			},
		},
	}

	program := map[string]any{
		"name":    proj.name,
		"runtime": "yaml",
		"variables": map[string]any{
			"clientConfig": map[string]any{
				"fn::invoke": map[string]any{
					"function": "azure-native:authorization:getClientConfig",
				},
			},
			"storageAccountKeys": map[string]any{
				"fn::invoke": map[string]any{
					"function": "azure-native:storage:listStorageAccountKeys",
					"arguments": map[string]any{
						"resourceGroupName": "${resourcegroup.name}",
						"accountName":       "${sa.name}",
					},
				},
			},
		},
		"resources": map[string]any{
			"resourcegroup": map[string]any{
				"type": "azure-native:resources:ResourceGroup",
			},
			"server": map[string]any{
				"type": "azure-native:sql:Server",
				"properties": map[string]any{
					"resourceGroupName":          "${resourcegroup.name}",
					"location":                   "${resourcegroup.location}",
					"administratorLogin":         "dummylogin",
					"administratorLoginPassword": "Un53cuRE!",
					"version":                    "12.0",
				},
				"options": map[string]any{
					"ignoreChanges": []string{"administrators"},
				},
			},
			"enableADS": map[string]any{
				"type": "azure-native:sql:ServerSecurityAlertPolicy",
				"properties": map[string]any{
					"resourceGroupName":       "${resourcegroup.name}",
					"serverName":              "${server.name}",
					"securityAlertPolicyName": "Default",
					"state":                   "Enabled",
				},
			},
			"sa": map[string]any{
				"type": "azure-native:storage:StorageAccount",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"sku": map[string]any{
						"name": "Standard_LRS",
					},
					"kind": "StorageV2",
				},
			},
			"blobs": map[string]any{
				"type": "azure-native:storage:BlobContainer",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"accountName":       "${sa.name}",
				},
			},
			"serverVulnAssessment": map[string]any{
				"type": "azure-native:sql:ServerVulnerabilityAssessment",
				"properties": map[string]any{
					"resourceGroupName":           "${resourcegroup.name}",
					"serverName":                  "${server.name}",
					"vulnerabilityAssessmentName": "default",
					"recurringScans": map[string]any{
						"emailSubscriptionAdmins": false,
						"emails":                  []string{"hi@example.com"},
						"isEnabled":               true,
					},
					"storageContainerPath":    "https://${sa.name}.blob.core.windows.net/${blobs.name}",
					"storageAccountAccessKey": "${storageAccountKeys.keys[0].value}",
				},
				"options": map[string]any{
					"dependsOn": []string{"${enableADS}"},
				},
			},
			"sqlFwRule": map[string]any{
				"type": "azure-native:sql:FirewallRule",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"serverName":        "${server.name}",
					"firewallRuleName":  "ClientIPAddress",
					"startIpAddress":    "222.222.222.222",
					"endIpAddress":      "222.222.222.222",
				},
			},
			"vnet": map[string]any{
				"type": "azure-native:network:VirtualNetwork",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"location":          "${resourcegroup.location}",
					"addressSpace": map[string]any{
						"addressPrefixes": []string{"10.1.0.0/16"},
					},
					"subnets": []map[string]any{
						{
							"name":                           "default",
							"addressPrefix":                  "10.1.0.0/24",
							"privateEndpointNetworkPolicies": "Disabled",
						},
					},
				},
			},
			"endpoint": map[string]any{
				"type": "azure-native:network:PrivateEndpoint",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"location":          "${resourcegroup.location}",
					"privateLinkServiceConnections": []map[string]any{
						{
							"groupIds":             []string{"sqlServer"},
							"privateLinkServiceId": "${server.id}",
							"name":                 "conn-sql",
						},
					},
					"subnet": map[string]any{
						"id": "${vnet.subnets[0].id}",
					},
				},
			},
			"zoneGroup": map[string]any{
				"type": "azure-native:network:PrivateDnsZoneGroup",
				"properties": map[string]any{
					"resourceGroupName":   "${resourcegroup.name}",
					"privateEndpointName": "${endpoint.name}",
				},
			},
			"adAdmin": map[string]any{
				"type": "azure-native:sql:ServerAzureADAdministrator",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"serverName":        "${server.name}",
					"administratorName": "ActiveDirectory",
					"administratorType": "ActiveDirectory",
					"login":             "foo@example.com",
					"sid":               "c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c",
					"tenantId":          "${clientConfig.tenantId}",
				},
			},
			"adOnlyAuth": map[string]any{
				"type": "azure-native:sql:ServerAzureADOnlyAuthentication",
				"properties": map[string]any{
					"resourceGroupName":         "${resourcegroup.name}",
					"serverName":                "${server.name}",
					"authenticationName":        "Default",
					"azureADOnlyAuthentication": false,
				},
				"options": map[string]any{
					"dependsOn": []string{"${adAdmin}"},
				},
			},
		},
		"outputs": map[string]any{
			"serverName": "${server.name}",
		},
		"plugins": plugins,
	}

	updatePulumiYAML(t, test.WorkingDir(), program)

	upResult := test.Up(t)
	assert.Empty(t, upResult.StdErr, "up should not have any errors")
	defer test.Destroy(t)

	output, ok := upResult.Outputs["serverName"]
	require.True(t, ok, "serverName output should be present")
	serverName, ok := output.Value.(string)
	require.True(t, ok, "serverName should be a string")
	assert.NotEmpty(t, serverName, "serverName should not be empty")

	// Assert that preview doesn't have changes
	preview := test.Preview(t)
	t.Logf("Preview STDOUT: \n%s", preview.StdOut)
	assert.Equal(t, map[apitype.OpType]int{
		apitype.OpSame: 13,
	}, preview.ChangeSummary)
}

func TestManagedClusterHasPopulatedKubeletIdentity_YAML(t *testing.T) {
	skipIfNotYamlInCI(t)
	proj := tempProject(t)
	azureBinaryDir := azureNativeBinaryDir(t)
	test := createTest(t, proj.dir)

	plugins := map[string]any{
		"providers": []interface{}{
			map[string]any{
				"name": "azure-native",
				"path": azureBinaryDir,
			},
		},
	}

	program := map[string]any{
		"name":    proj.name,
		"runtime": "yaml",
		"variables": map[string]any{
			"clientConfig": map[string]any{
				"fn::invoke": map[string]any{
					"function": "azure-native:authorization:getClientConfig",
				},
			},
		},
		"resources": map[string]any{
			"resourcegroup": map[string]any{
				"type": "azure-native:resources:ResourceGroup",
			},
			"managedCluster": map[string]any{
				"type": "azure-native:containerservice:ManagedCluster",
				"properties": map[string]any{
					"resourceGroupName": "${resourcegroup.name}",
					"location":          "${resourcegroup.location}",
					"dnsPrefix":         "dns-prefix",
					"enableRBAC":        true,
					"identity": map[string]any{
						"type": "SystemAssigned",
					},
					"aadProfile": map[string]any{
						"enableAzureRBAC": true,
						"managed":         true,
						"adminGroupObjectIDs": []string{
							"${clientConfig.objectId}",
						},
					},
					"agentPoolProfiles": []map[string]any{
						{
							"name":                "agentpool",
							"count":               1,
							"vmSize":              "Standard_DS2_v2",
							"osType":              "Linux",
							"mode":                "System",
							"maxPods":             110,
							"type":                "VirtualMachineScaleSets",
							"orchestratorVersion": "1.21.2",
						},
					},
					"linuxProfile": map[string]any{
						"adminUsername": "azureuser",
						"ssh": map[string]any{
							"publicKeys": []map[string]any{
								{
									// generated via ssh-keygen -t rsa -b 2048 -f /tmp/test_key -N "" -C "test@example.com" 2>/dev/null && cat /tmp/test_key.pub
									"keyData": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDPNoOivGx7lpsvSGx6603NLFhXCu9sl1c7Dcao29+p8m4kqElT6IxlfEZombjIGG/23vGRQr1/cuHhnOPieNf0UCEjmvkflm7iXjJTOVC04f96P9DiEck1+Vts+H9kBWYvik3nzmCayrX8lbjQKV9+UpmIpBC1OboETzsdhVqeXzaYAXE6wYA5jzjFaVAK9ORMiAQyTiTkEt6eGQoDLTdBbVPkFS2JNvjGTydhIakQ4bVdY8nUrj7hHGmsCu84ydKE36s5Am3pwU1cGNhfjGWDz8NViDQJ27eWzYNdhfwQ5F5/gu64XvY00XPh6F2wQ1jMMq3yzi9RNg1usedOVJ29 test@example.com",
								},
							},
						},
					},
				},
			},
		},
		"outputs": map[string]any{
			"clientId":   "${managedCluster.identityProfile.kubeletidentity.clientId}",
			"objectId":   "${managedCluster.identityProfile.kubeletidentity.objectId}",
			"resourceId": "${managedCluster.identityProfile.kubeletidentity.resourceId}",
		},
		"plugins": plugins,
	}

	// initial program
	updatePulumiYAML(t, test.WorkingDir(), program)

	upResult := test.Up(t)
	assert.Empty(t, upResult.StdErr, "up should not have any errors")

	nonEmpty := func(outputKey string) {
		output, ok := upResult.Outputs[outputKey]
		require.True(t, ok, "output %s should be present", outputKey)
		value, ok := output.Value.(string)
		require.True(t, ok, "output %s should be a string, instead it was %T", outputKey, output.Value)
		assert.NotEmpty(t, value, "value of '%s' should not be empty", outputKey)
	}

	defer test.Destroy(t)

	nonEmpty("clientId")
	nonEmpty("objectId")
	nonEmpty("resourceId")
}

// TestTagAtScopeAddedOnSecondDeploy_YAML verifies that a TagAtScope resource can be added
// to a resource group after it has already been created, using the PATCH/Merge path.
func TestTagAtScopeAddedOnSecondDeploy_YAML(t *testing.T) {
	skipIfNotYamlInCI(t)
	proj := tempProject(t)
	azureBinaryDir := azureNativeBinaryDir(t)
	test := createTest(t, proj.dir)

	plugins := map[string]any{
		"providers": []interface{}{
			map[string]any{
				"name": "azure-native",
				"path": azureBinaryDir,
			},
		},
	}

	program := map[string]any{
		"name":    proj.name,
		"runtime": "yaml",
		"resources": map[string]any{
			"resourcegroup": map[string]any{
				"type": "azure-native:resources:ResourceGroup",
			},
			"tags": map[string]any{
				"type": "azure-native:resources:TagAtScope",
				"properties": map[string]any{
					"scope": "${resourcegroup.id}",
					"properties": map[string]any{
						"tags": map[string]any{
							"environment": "test",
							"managedBy":   "pulumi",
						},
					},
				},
			},
		},
		"outputs": map[string]any{
			"rg": "${resourcegroup.id}",
		},
		"plugins": plugins,
	}

	// Deploy only the resource group.
	updatePulumiYAML(t, test.WorkingDir(), program)

	upResult := test.Up(t)
	assert.Empty(t, upResult.StdErr, "first up should not have any errors")
	defer test.Destroy(t)

	// Add a TagAtScope resource on the second deploy.
	program["resources"].(map[string]any)["tagsV2"] = map[string]any{
		"type": "azure-native:resources:TagAtScope",
		"properties": map[string]any{
			"scope": "${resourcegroup.id}",
			"properties": map[string]any{
				"tags": map[string]any{
					"another": "tag",
				},
			},
		},
	}

	updatePulumiYAML(t, test.WorkingDir(), program)

	upResult = test.Up(t)
	assert.Empty(t, upResult.StdErr, "second up should not have any errors")

	// Retrieve the resource group ID from the stack output and verify that both TagAtScope
	// resources' tags are present on the scope via the Azure API.
	rgID, ok := upResult.Outputs["rg"]
	require.True(t, ok, "rg output should be present")
	rgIDStr, ok := rgID.Value.(string)
	require.True(t, ok, "rg output should be a string")

	ctx := context.Background()
	tagsResp, err := tagsClient(t).GetAtScope(ctx, rgIDStr, nil)
	require.NoError(t, err, "getting tags at scope should not error")

	actualTags := map[string]string{}
	if tagsResp.Properties != nil {
		for k, v := range tagsResp.Properties.Tags {
			if v != nil {
				actualTags[k] = *v
			}
		}
	}
	t.Logf("Tags on scope: %v", actualTags)

	// Tags from the "tags" resource.
	assert.Equal(t, "test", actualTags["environment"])
	assert.Equal(t, "pulumi", actualTags["managedBy"])
	// Tags from the "tagsV2" resource.
	assert.Equal(t, "tag", actualTags["another"])

	// A subsequent preview should show no changes, confirming the operation is idempotent.
	preview := test.Preview(t)
	t.Logf("Preview STDOUT: \n%s", preview.StdOut)
	assert.Equal(t, map[apitype.OpType]int{
		apitype.OpSame: 4, // stack + resourcegroup + tags + tagsV2
	}, preview.ChangeSummary)
}
