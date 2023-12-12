package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type resourceIDCase struct {
	id   string
	path string
}

func TestParseInvalidResourceID(t *testing.T) {
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
}

func TestParseFullResourceID(t *testing.T) {
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
}

func TestParseScopedResourceID(t *testing.T) {
	id := "/subscriptions/1200b1c8-3c58-42db-b33a-304a75913333/resourceGroups/devops-dev/providers/Microsoft.Authorization/roleAssignments/2a88abc7-f599-0eba-a21f-a1817e597115"
	path := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	actual, err := ParseResourceID(id, path)
	assert.NoError(t, err)
	expected := map[string]string{
		"scope":              "subscriptions/1200b1c8-3c58-42db-b33a-304a75913333/resourceGroups/devops-dev",
		"roleAssignmentName": "2a88abc7-f599-0eba-a21f-a1817e597115",
	}
	assert.Equal(t, expected, actual)
}
