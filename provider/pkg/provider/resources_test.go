// Copyright 2016-2020, Pulumi Corporation.

package provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceName(t *testing.T) {
	testCases := map[string]string{
		"GetUserSettings": "UserSettings",
		"Mediaservices_Get": "MediaService",
		"Redis_Get": "Redis",
		"AssessmentsMetadata_GetInSubscription": "AssessmentMetadataInSubscription",
		"Caches_Get": "Cache",
		"CognitiveServicesAccounts_GetProperties": "CognitiveServicesAccount",
		"WorkspaceCollections_getByName": "WorkspaceCollection",
		"getUserSettingsWithLocation": "UserSettingsWithLocation",
		"SupportPlanTypes_ListInfo": "SupportPlanTypeInfo",
		"WebSite_GetWebSite": "WebSite",
		"ManagedCluster_ClusterUserGet": "ManagedClusterUser",
		"BlobService_ServicePropertiesGet": "BlobServiceProperties",
	}
	for operationId, expected := range testCases {
		actual := ResourceName(operationId)
		assert.Equal(t, expected, actual)
	}
}
