package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type resourceIDCase struct {
	id   string
	path string
}

func TestParseResourceID(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		cases := []resourceIDCase{
			// ID shorter than Path
			{"/resourceGroup/myrg", "/resourceGroup/{resourceGroup}/subResource"},
			// ID longer than Path
			{"/resourceGroup/myrg/cdn/mycdn", "/resourceGroup/{resourceGroup}/cdn"},
			// Segment names don't match
			{"/resourceGroup/myrg/foo/mycdn", "/resourceGroup/{resourceGroup}/bar/{cdn}"},
		}
		for _, testCase := range cases {
			_, err := ParseResourceID(testCase.id, testCase.path)
			assert.Error(t, err)
		}
	})

	t.Run("full", func(t *testing.T) {
		id := "/subscriptions/0282681f-7a9e-123b-40b2-96babd57a8a1/resourcegroups/pulumi-name/providers/Microsoft.Network/networkInterfaces/pulumi-nic/ipConfigurations/ipconfig1"
		path := "/subscriptions/{subscriptionID}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}"
		actual, err := ParseResourceID(id, path)
		assert.NoError(t, err)
		expected := map[string]string{
			"subscriptionID":       "0282681f-7a9e-123b-40b2-96babd57a8a1",
			"resourceGroupName":    "pulumi-name",
			"networkInterfaceName": "pulumi-nic",
			"ipConfigurationName":  "ipconfig1",
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("scoped", func(t *testing.T) {
		id := "/subscriptions/1200b1c8-3c58-42db-b33a-304a75913333/resourceGroups/devops-dev/providers/Microsoft.Authorization/roleAssignments/2a88abc7-f599-0eba-a21f-a1817e597115"
		path := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
		actual, err := ParseResourceID(id, path)
		assert.NoError(t, err)
		expected := map[string]string{
			"scope":              "subscriptions/1200b1c8-3c58-42db-b33a-304a75913333/resourceGroups/devops-dev",
			"roleAssignmentName": "2a88abc7-f599-0eba-a21f-a1817e597115",
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("nested identifier", func(t *testing.T) {
		id := "/subscriptions/1200b1c8-3c58-42db-b33a-304a75913333/resourceGroups/resource-group/providers/Microsoft.KeyVault/vaults/vault-name/accessPolicy/policy-object-id"
		path := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/accessPolicy/{policy.objectId}"
		actual, err := ParseResourceID(id, path)
		assert.NoError(t, err)
		expected := map[string]string{
			"subscriptionId":    "1200b1c8-3c58-42db-b33a-304a75913333",
			"resourceGroupName": "resource-group",
			"vaultName":         "vault-name",
			"policy.objectId":   "policy-object-id",
		}
		assert.Equal(t, expected, actual)
	})
}
