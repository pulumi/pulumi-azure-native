// Copyright 2024, Pulumi Corporation.

package resources

import (
	"fmt"
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

func TestParseToken(t *testing.T) {
	type testCase struct {
		mod          string
		apiVersion   string
		resourceName string
		err          string
	}
	tests := []struct {
		name     string
		token    string
		expected testCase
	}{
		{
			name:  "ValidTokenDefaultVersion",
			token: "azure-native:messaging:ServiceBus",
			expected: testCase{
				mod:          "messaging",
				apiVersion:   "",
				resourceName: "ServiceBus",
			},
		},
		{
			name:  "ValidTokenExplicitVersion",
			token: "azure-native:messaging/v20240101preview:ServiceBus",
			expected: testCase{
				mod:          "messaging",
				apiVersion:   "v20240101preview",
				resourceName: "ServiceBus",
			},
		},
		{
			name:  "MalformedTokenTooFewParts",
			token: "messaging/v20240101preview:ServiceBus",
			expected: testCase{
				err: "malformed token 'messaging/v20240101preview:ServiceBus'",
			},
		},
		{
			name:  "MalformedTokenTooManyParts",
			token: "azure-native:messaging:ServiceBus:eventhub:EventHub",
			expected: testCase{
				err: "malformed token 'azure-native:messaging:ServiceBus:eventhub:EventHub'",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mod, apiVersion, resourceName, err := ParseToken(test.token)
			assert.Equal(t, test.expected.mod, mod)
			assert.Equal(t, test.expected.apiVersion, apiVersion)
			assert.Equal(t, test.expected.resourceName, resourceName)
			if err != nil {
				assert.Equal(t, test.expected.err, err.Error())
			} else {
				assert.Equal(t, test.expected.err, "")
			}
		})
	}
}
func TestBuildToken(t *testing.T) {
	tests := []struct {
		mod        string
		apiVersion string
		name       string
		expected   string
	}{
		{
			mod:        "messaging",
			apiVersion: "",
			name:       "ServiceBus",
			expected:   "azure-native:messaging:ServiceBus",
		},
		{
			mod:        "messaging",
			apiVersion: "v20240101preview",
			name:       "ServiceBus",
			expected:   "azure-native:messaging/v20240101preview:ServiceBus",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("mod=%s, apiVersion=%s, name=%s", test.mod, test.apiVersion, test.name), func(t *testing.T) {
			actual := BuildToken(test.mod, test.apiVersion, test.name)
			assert.Equal(t, test.expected, actual)
		})
	}
}
