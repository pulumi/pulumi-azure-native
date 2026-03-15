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

// This is a test to verify that adding the isHnsEnabled property to a storage account
// does not cause replacements, since it defaults to false.
// in fact, adding a property that defaults to false should not cause any changes at all since
// the default value is the same as what we are setting it to.
func TestAddingHnsEnabledToStorageAccountDoesNotCauseReplacements(t *testing.T) {
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
func TestListSubscriptionsInvoke(t *testing.T) {
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

func TestManagedClusterHasPopulatedKubeletIdentity(t *testing.T) {
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
	//nonEmpty("principalId")
	nonEmpty("objectId")
	nonEmpty("resourceId")
}
