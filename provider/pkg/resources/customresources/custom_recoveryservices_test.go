// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/stretchr/testify/assert"
)

func TestUpdateInputWithFileShareSystemName(t *testing.T) {
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
		input := azureFileShareInput()

		reader := &mockSystemNameReader{systemName: "systemName"}
		err := updateInputWithFileShareSystemName(context.Background(), input, reader)
		assert.NoError(t, err)
		assert.Equal(t, "systemName", input["protectedItemName"].StringValue())
		assert.Equal(t, "fileshare", input["__friendlyProtectedItemName"].StringValue())
	})

	t.Run("leaves original name unchanged if no system name is found", func(t *testing.T) {
		input := azureFileShareInput()
		origInput := input.Copy()

		reader := &mockSystemNameReader{systemName: ""}
		err := updateInputWithFileShareSystemName(context.Background(), input, reader)
		assert.NoError(t, err)
		assert.Equal(t, origInput, input)
	})

	t.Run("should not modify protected item of type other than AzureFileShareProtectedItem", func(t *testing.T) {
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
