// Copyright 2016-2020, Pulumi Corporation.

package resources

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceName(t *testing.T) {
	testCases := map[string]string{
		"GetUserSettings":                         "UserSettings",
		"Mediaservices_Get":                       "MediaService",
		"Redis_Get":                               "Redis",
		"AssessmentsMetadata_GetInSubscription":   "AssessmentMetadataInSubscription",
		"Caches_Get":                              "Cache",
		"CognitiveServicesAccounts_GetProperties": "CognitiveServicesAccount",
		"WorkspaceCollections_getByName":          "WorkspaceCollection",
		"getUserSettingsWithLocation":             "UserSettingsWithLocation",
		"SupportPlanTypes_ListInfo":               "SupportPlanTypeInfo",
		"WebSite_GetWebSite":                      "WebSite",
		"ManagedCluster_ClusterUserGet":           "ManagedClusterUser",
		"BlobService_ServicePropertiesGet":        "BlobServiceProperties",
		"SaasResource-listAccessToken":            "SaasResourceAccessToken",
		"WebApps_ListApplicationSettings":         "WebAppApplicationSettings",
		"Products_GetProducts":                    "Products",
	}
	for operationID, expected := range testCases {
		actual := ResourceName(operationID)
		assert.Equal(t, expected, actual)
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
		{"/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.CostManagement/costAllocationRules/{ruleName}","ruleName", "random"},
	}
	for _, values := range testCases {
		path := values[0]
		name := values[1]
		expected := values[2]
		actual, ok := AutoName(name, path)
		if !ok && expected != "" {
			assert.Fail(t, "expected auto-name for %q %q", path, name)
		}
		assert.Equal(t, expected, string(actual))
	}
}
