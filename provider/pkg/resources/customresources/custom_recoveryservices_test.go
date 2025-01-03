// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	recovery "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicesbackup/v4"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestUpdateInputWithFileShareSystemName(t *testing.T) {
	t.Parallel()

	azureFileShareInput := func() resource.PropertyMap {
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

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()
		input := azureFileShareInput()

		reader := &mockSystemNameReader{systemName: "systemName"}
		err := updateInputWithFileShareSystemName(context.Background(), input, reader)
		assert.NoError(t, err)
		assert.Equal(t, "systemName", input["protectedItemName"].StringValue())
		assert.Equal(t, "fileshare", input["__friendlyProtectedItemName"].StringValue())
	})

	t.Run("leaves original name unchanged if no system name is found", func(t *testing.T) {
		t.Parallel()
		input := azureFileShareInput()
		origInput := input.Copy()

		reader := &mockSystemNameReader{systemName: ""}
		err := updateInputWithFileShareSystemName(context.Background(), input, reader)
		assert.NoError(t, err)
		assert.Equal(t, origInput, input)
	})

	t.Run("should not modify protected item of type other than AzureFileShareProtectedItem", func(t *testing.T) {
		t.Parallel()
		input := azureFileShareInput()
		input["properties"].ObjectValue()["protectedItemType"] = resource.NewStringProperty("SomeOtherProtectedItem")
		origInput := input.Copy()

		reader := &mockSystemNameReader{systemName: "test"}
		err := updateInputWithFileShareSystemName(context.Background(), input, reader)
		assert.NoError(t, err)
		assert.Equal(t, origInput, input)
	})
}

type mockSystemNameReader struct {
	systemName string
	err        error
}

func (m *mockSystemNameReader) readSystemNameFromProtectableItem(ctx context.Context, input *protectedItemProperties) (string, error) {
	return m.systemName, m.err
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
