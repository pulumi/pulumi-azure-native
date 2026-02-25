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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
	"github.com/pulumi/providertest/pulumitest"
	"github.com/pulumi/providertest/pulumitest/opttest"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func azureNativeBinary(t *testing.T) opttest.Option {
	binPath, err := filepath.Abs("../bin")
	if err != nil {
		t.Fatal(err)
	}
	return opttest.LocalProviderPath("azure-native", binPath)
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

func createTest(t *testing.T, source string) *pulumitest.PulumiTest {
	pt := pulumitest.NewPulumiTest(t,
		source,
		azureNativeBinary(t))

	pt.SetConfig(t, "azure-native:location", getLocation(t))
	return pt
}

func azureCredentials(t *testing.T) azcore.TokenCredential {
	if os.Getenv("CI") == "" {
		// local test, use CLI credentials
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
