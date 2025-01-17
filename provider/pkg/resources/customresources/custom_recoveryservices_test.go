// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	recovery "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func not var to ensure we have separate copies for separate tests.
func standardAzureFileshareProtectedItemInput() resource.PropertyMap {
	return resource.NewPropertyMapFromMap(map[string]any{
		"containerName":     "container",
		"fabricName":        "Azure",
		"protectedItemName": "fileshare",
		resourceGroupName:   "rg",
		"vaultName":         "vault",
		"properties": map[string]any{
			"protectedItemType": "AzureFileShareProtectedItem",
			"policyId":          "policy",
			"sourceResourceId":  "azure/sa1",
		},
	})
}

func TestGetIdAndQuery(t *testing.T) {
	t.Parallel()

	// A minimal copy of the actual resource metadata, reduced to what's needed for the test.
	protectedItemResource := resources.AzureAPIResource{
		APIVersion: "2023-04-01",
		Path:       protectedItemPath,
		PutParameters: []resources.AzureAPIParameter{
			{
				Name:     "resourceGroupName",
				Location: "path",
				Value:    &resources.AzureAPIProperty{},
			},
			{
				Name:     "vaultName",
				Location: "path",
				Value:    &resources.AzureAPIProperty{},
			},
			{
				Name:     "subscriptionId",
				Location: "path",
				Value:    &resources.AzureAPIProperty{},
			},
			{
				Name:     "fabricName",
				Location: "path",
				Value:    &resources.AzureAPIProperty{},
			},
			{
				Name:     "containerName",
				Location: "path",
				Value:    &resources.AzureAPIProperty{},
			},
			{
				Name:     "protectedItemName",
				Location: "path",
				Value:    &resources.AzureAPIProperty{},
			},
			{
				Name:     "properties",
				Location: "body",
				Value:    &resources.AzureAPIProperty{},
				Body: &resources.AzureAPIType{
					Properties: map[string]resources.AzureAPIProperty{
						"protectedItemType": {Type: "string"},
						"policyId":          {Type: "string"},
						"sourceResourceId":  {Type: "string"},
					},
				},
			},
		},
	}

	reader := &mockSystemNameReader{systemNames: []string{"azurefileshare;339f9859"}}
	azureClient := azure.MockAzureClient{}
	crudClient := crud.NewResourceCrudClient(&azureClient, nil, nil, "123", &protectedItemResource)

	id, queryParams, err := getIdAndQuery(context.Background(), standardAzureFileshareProtectedItemInput(), crudClient, reader)
	require.NoError(t, err)
	assert.Equal(t, "/subscriptions/123/resourceGroups/rg/providers/Microsoft.RecoveryServices/vaults/vault/backupFabrics/Azure/protectionContainers/container/protectedItems/azurefileshare%3B339f9859", id)

	// Query params should be identical to non-custom ones.
	_, regularQueryParams, err := crudClient.PrepareAzureRESTIdAndQuery(resource.PropertyMap{})
	require.NoError(t, err)
	assert.Equal(t, regularQueryParams, queryParams)
}

func TestRetrieveSystemName(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		reader := &mockSystemNameReader{systemNames: []string{"systemName"}}
		systemName, err := retrieveSystemName(context.Background(), standardAzureFileshareProtectedItemInput(), reader, 1, 0*time.Second)
		assert.NoError(t, err)
		assert.Equal(t, "systemName", systemName)
		assert.True(t, reader.inquireCalled)
	})

	t.Run("no system name is found", func(t *testing.T) {
		t.Parallel()
		reader := &mockSystemNameReader{systemNames: []string{""}}
		systemName, err := retrieveSystemName(context.Background(), standardAzureFileshareProtectedItemInput(), reader, 1, 0*time.Second)
		assert.Error(t, err)
		assert.Empty(t, systemName)
		assert.True(t, reader.inquireCalled)
	})

	t.Run("system name is found after multiple attempts", func(t *testing.T) {
		t.Parallel()
		reader := &mockSystemNameReader{
			systemNames: []string{"", "", "systemName", ""},
			errors:      []error{nil, nil, nil, nil},
		}
		systemName, err := retrieveSystemName(context.Background(), standardAzureFileshareProtectedItemInput(), reader, 5, 0*time.Second)
		assert.NoError(t, err)
		assert.Equal(t, "systemName", systemName)
		assert.True(t, reader.inquireCalled)
	})

	t.Run("stops after max attempts", func(t *testing.T) {
		t.Parallel()
		reader := &mockSystemNameReader{
			systemNames: []string{"", "", "systemName"},
			errors:      []error{nil, nil, nil},
		}
		systemName, err := retrieveSystemName(context.Background(), standardAzureFileshareProtectedItemInput(), reader, 2, 0*time.Second)
		assert.Error(t, err)
		assert.Empty(t, systemName)
		assert.True(t, reader.inquireCalled)
	})

	t.Run("no system name for protected item of type other than AzureFileShareProtectedItem", func(t *testing.T) {
		t.Parallel()
		input := standardAzureFileshareProtectedItemInput()
		input["properties"].ObjectValue()["protectedItemType"] = resource.NewStringProperty("SomeOtherProtectedItem")

		reader := &mockSystemNameReader{systemNames: []string{"test"}, errors: []error{nil}}
		systemName, err := retrieveSystemName(context.Background(), input, reader, 1, 0*time.Second)
		assert.NoError(t, err)
		assert.Empty(t, systemName)
		assert.False(t, reader.inquireCalled)
	})
}

type mockSystemNameReader struct {
	systemNames []string
	errors      []error
	attempt     int
	// Was the /inquire API resp. the SDK's ProtectionContainersClient.Inquire called?
	inquireCalled bool
}

func (m *mockSystemNameReader) readSystemNameFromProtectableItem(ctx context.Context, input *protectedItemProperties) (string, error) {
	name := m.systemNames[m.attempt]
	var err error
	if len(m.errors) > m.attempt {
		err = m.errors[m.attempt]
	}
	m.attempt++
	return name, err
}

func (m *mockSystemNameReader) inquireContainer(ctx context.Context, input *protectedItemProperties) error {
	m.inquireCalled = true
	return nil
}

func TestFindSystemName(t *testing.T) {
	t.Parallel()

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		friendlyName := "friendlyName"
		storageAccountName := "storageAccountName"
		items := []*recovery.WorkloadProtectableItemResource{
			{
				Name: to.Ptr("different one"),
				Properties: &recovery.AzureFileShareProtectableItem{
					FriendlyName:                to.Ptr("different name"),
					ParentContainerFriendlyName: to.Ptr(storageAccountName),
				},
			},
			{
				Name: to.Ptr("different two"),
				Properties: &recovery.AzureFileShareProtectableItem{
					FriendlyName:                to.Ptr(friendlyName),
					ParentContainerFriendlyName: to.Ptr("different storage account name"),
				},
			},
			{
				Name: to.Ptr("systemName"),
				Properties: &recovery.AzureFileShareProtectableItem{
					FriendlyName:                to.Ptr(friendlyName),
					ParentContainerFriendlyName: to.Ptr(storageAccountName),
				},
			},
		}
		systemName, err := findSystemName(items, "friendlyName", "storageAccountName")
		assert.NoError(t, err)
		assert.Equal(t, "systemName", systemName)
	})

	t.Run("returns empty string if no matching item is found", func(t *testing.T) {
		t.Parallel()
		items := []*recovery.WorkloadProtectableItemResource{
			{
				Name: to.Ptr("different one"),
				Properties: &recovery.AzureFileShareProtectableItem{
					FriendlyName:                to.Ptr("different name"),
					ParentContainerFriendlyName: to.Ptr("different storageAccountName"),
				},
			},
		}
		systemName, err := findSystemName(items, "friendlyName", "storageAccountName")
		assert.NoError(t, err)
		assert.Equal(t, "", systemName)
	})

	t.Run("returns error if the item is not a file share", func(t *testing.T) {
		t.Parallel()
		items := []*recovery.WorkloadProtectableItemResource{
			{
				Name: to.Ptr("systemName"),
				Properties: &recovery.AzureVMWorkloadProtectableItem{
					FriendlyName: to.Ptr("friendlyName"),
				},
			},
		}
		systemName, err := findSystemName(items, "friendlyName", "storageAccountName")
		assert.Error(t, err)
		assert.Equal(t, "", systemName)
	})

	t.Run("no items", func(t *testing.T) {
		t.Parallel()
		systemName, err := findSystemName(nil, "friendlyName", "storageAccountName")
		assert.NoError(t, err)
		assert.Equal(t, "", systemName)
	})
}
