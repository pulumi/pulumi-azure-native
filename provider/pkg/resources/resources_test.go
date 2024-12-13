// Copyright 2016-2020, Pulumi Corporation.

package resources

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type resourceNameTestCase struct {
	operationID string
	path        string
	expected    string
}

func TestResourceName(t *testing.T) {
	// Empty path means "doesn't matter", not that these resources don't have a path.
	testCases2 := []resourceNameTestCase{
		{"GetUserSettings", "", "UserSettings"},
		{"Mediaservices_Get", "", "MediaService"},
		{"Redis_Get", "", "Redis"},
		{"AssessmentsMetadata_GetInSubscription", "", "AssessmentMetadataInSubscription"},
		{"Caches_Get", "", "Cache"},
		{"CognitiveServicesAccounts_GetProperties", "", "CognitiveServicesAccount"},
		{"WorkspaceCollections_getByName", "", "WorkspaceCollection"},
		{"getUserSettingsWithLocation", "", "UserSettingsWithLocation"},
		{"SupportPlanTypes_ListInfo", "", "SupportPlanTypeInfo"},
		{"WebSite_GetWebSite", "", "WebSite"},
		{"ManagedCluster_ClusterUserGet", "", "ManagedClusterUser"},
		{"BlobService_ServicePropertiesGet", "", "BlobServiceProperties"},
		{"SaasResource-listAccessToken", "", "SaasResourceAccessToken"},
		{"WebApps_ListApplicationSettings", "", "WebAppApplicationSettings"},
		{"Products_GetProducts", "", "Products"},
		{"PowerBIResources_ListByResourceName", "", "PowerBIResource"},

		// An exception for https://github.com/pulumi/pulumi-azure-native/issues/583, disambiguated by path.
		{
			"RecordSets_Get",
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateZoneName}/{recordType}/{relativeRecordSetName}",
			"PrivateRecordSet",
		},
		{
			"RecordSets_Get",
			"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsZones/{zoneName}/{recordType}/{relativeRecordSetName}",
			"RecordSet",
		},
	}

	for _, tc := range testCases2 {
		actual, _ := ResourceName(tc.operationID, tc.path)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestAutoName(t *testing.T) {
	testCases := [][]string{
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}", "resourceGroupName", "random"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}", "serverName", "random"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}", "databaseName", "copy"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}", "serverName", ""},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "name", "random"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}", "siteName", ""},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings", "name", ""},
		{"/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/costAllocationRules/{ruleName}", "ruleName", "random"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "endpointName", "random"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/customDomains/{customDomainName}", "customDomainName", "random"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}", "endpointName", "random"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", "customDomainName", "random"},
		{"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/jobSchedules/{jobScheduleId}", "jobScheduleId", "uuid", "uuid"},
		{"/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}", "roleAssignmentName", "uuid"},
		{"/{scope}/providers/Microsoft.Authorization/roleDefinitions/{roleDefinitionId}", "roleDefinitionId", "uuid"},
	}
	for _, values := range testCases {
		path := values[0]
		name := values[1]
		expected := values[2]
		format := ""
		if len(values) > 3 {
			format = values[2]
		}
		namer := NewAutoNamer(path)
		actual, ok := namer.AutoName(name, format)
		if !ok && expected != "" {
			assert.Fail(t, "expected auto-name for %q %q", path, name)
		}
		assert.Equal(t, expected, string(actual))
	}
}

func TestTraverseProperties(t *testing.T) {
	properties := map[string]AzureAPIProperty{
		"properties": {
			Type: "object",
			Ref:  "#/types/azure-native:keyvault:VaultProperties",
		},
		"location": {
			Type: "string",
		},
	}

	res := AzureAPIResource{
		PutParameters: []AzureAPIParameter{
			{
				Location: "body",
				Name:     "bodyProperties",
				Body: &AzureAPIType{
					Properties: properties,
				},
			},
		},
	}

	// Mock the type lookup to only return the type referenced in the resource above
	lookupType := func(ref string) (*AzureAPIType, bool, error) {
		if ref == "#/types/azure-native:keyvault:VaultProperties" {
			return &AzureAPIType{
				Properties: map[string]AzureAPIProperty{
					"accessPolicies": {
						Type: "array",
						Items: &AzureAPIProperty{
							Type: "object",
							Ref:  "#/types/azure-native:keyvault:AccessPolicyEntry",
						},
						Containers: []string{"container"}, // not the case in the real KV spec but we want to test this
					},
				},
			}, true, nil
		}
		if ref == "#/types/azure-native:keyvault:AccessPolicyEntry" {
			return &AzureAPIType{
				Properties: map[string]AzureAPIProperty{
					"permissions": {
						Type: "array",
						Items: &AzureAPIProperty{
							Type: "string", // not true in the real KV spec but good enough
						},
						Containers:                 []string{"container2", "container3"},
						MaintainSubResourceIfUnset: true,
					},
					"other_array": {
						Type: "array",
						Items: &AzureAPIProperty{
							Type: "string",
						},
					},
				},
			}, true, nil
		}
		return nil, false, nil
	}

	t.Run("including containers", func(t *testing.T) {
		visited := map[string][]string{}
		visitor := func(name string, property AzureAPIProperty, path []string) {
			visited[name] = path
		}

		TraverseProperties(properties, lookupType, true, visitor)

		expected := map[string][]string{
			"properties":     {},
			"accessPolicies": {"properties", "container"},
			"permissions":    {"properties", "container", "accessPolicies", "container2", "container3"},
			"other_array":    {"properties", "container", "accessPolicies"},
			"location":       {},
		}
		assert.Equal(t, expected, visited)
	})

	t.Run("without containers", func(t *testing.T) {
		visited := map[string][]string{}
		visitor := func(name string, property AzureAPIProperty, path []string) {
			visited[name] = path
		}

		TraverseProperties(properties, lookupType, false, visitor)

		expected := map[string][]string{
			"properties":     {},
			"accessPolicies": {"properties"},
			"permissions":    {"properties", "accessPolicies"},
			"other_array":    {"properties", "accessPolicies"},
			"location":       {},
		}
		assert.Equal(t, expected, visited)
	})

	t.Run("collect subresource properties with containers", func(t *testing.T) {
		paths := res.PathsToSubResourcePropertiesToMaintain(true, lookupType)
		assert.Equal(t, 1, len(paths))
		assert.Equal(t, []string{"properties", "container", "accessPolicies", "container2", "container3", "permissions"}, paths[0])
	})

	t.Run("collect subresource properties without containers", func(t *testing.T) {
		paths := res.PathsToSubResourcePropertiesToMaintain(false, lookupType)
		assert.Equal(t, 1, len(paths))
		assert.Equal(t, []string{"properties", "accessPolicies", "permissions"}, paths[0])
	})
}

func TestResourceProviderNaming(t *testing.T) {
	t.Run("Standard case", func(t *testing.T) {
		actual := ResourceProvider("/go/pulumi-azure-native/azure-rest-api-specs/specification/EnterpriseKnowledgeGraph/resource-manager/Microsoft.EnterpriseKnowledgeGraph/preview/2018-12-03/EnterpriseKnowledgeGraphSwagger.json", "/providers/Microsoft.EnterpriseKnowledgeGraph/operations")
		assert.Equal(t, "EnterpriseKnowledgeGraph", actual)
	})
	t.Run("PaloAltoNetworks namespace", func(t *testing.T) {
		actual := ResourceProvider("/go/pulumi-azure-native/azure-rest-api-specs/specification/paloaltonetworks/resource-manager/PaloAltoNetworks.Cloudngfw/preview/2022-08-29-preview/PaloAltoNetworks.Cloudngfw.json", "/providers/PaloAltoNetworks.Cloudngfw/globalRulestacks")
		assert.Equal(t, "Cloudngfw", actual)
	})
	t.Run("When the provider names of file path and URI don't match, return empty", func(t *testing.T) {
		actual := ResourceProvider("/go/pulumi-azure-native/azure-rest-api-specs/specification/EnterpriseKnowledgeGraph/resource-manager/Microsoft.One/preview/2018-12-03/EnterpriseKnowledgeGraphSwagger.json", "/providers/Microsoft.Two/operations")
		assert.Equal(t, "", actual)
	})
	t.Run("Change lower case to title case", func(t *testing.T) {
		actual := ResourceProvider("/go/pulumi-azure-native/azure-rest-api-specs/specification/EnterpriseKnowledgeGraph/resource-manager/microsoft.fooBar/preview/2018-12-03/EnterpriseKnowledgeGraphSwagger.json", "/providers/microsoft.fooBar/operations")
		assert.Equal(t, "FooBar", actual)
	})
}
