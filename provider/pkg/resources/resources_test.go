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
