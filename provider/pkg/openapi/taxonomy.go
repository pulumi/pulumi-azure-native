package openapi

import "github.com/pulumi/pulumi/pkg/v2/codegen"

var ambientResources = []string {
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames",
}

var ambientResourceMap = codegen.NewStringSet()
func init() {
	for _, path := range ambientResources {
		ambientResourceMap.Add(normalizePath(path))
	}
}

// isAmbient returns true for paths of resources that can't be created or deleted explicitly. Their lifecycle matches
// the lifecycle of the parent resource, managed automatically by Azure.
func isAmbient(path string) bool {
	return ambientResourceMap.Has(normalizePath(path))
}
