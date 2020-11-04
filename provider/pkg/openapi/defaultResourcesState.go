package openapi

// defaultResourcesStateRaw is a map non-normalized paths. It's handy to have paths as they are in the Open API spec's
// latest version for easy search. This map shouldn't be used for lookups: use 'defaultResourcesStateMap' instead.
var defaultResourcesStateRaw = map[string]map[string]interface{}{
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMariaDB/servers/{serverName}/configurations/{configurationName}": {
		"source": "system-default",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySQL/servers/{serverName}/configurations/{configurationName}": {
		"source": "system-default",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSQL/servers/{serverName}/configurations/{configurationName}": {
		"source": "system-default",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkRuleSets/default": {
		"defaultAction": "Deny",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/networkRuleSets/default": {
		"defaultAction": "Deny",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{transparentDataEncryptionName}": {
		"status": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}":  {},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}":  {},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/{queueServiceName}": {},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/{tableServiceName}": {},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/transparentDataEncryption/{transparentDataEncryptionName}": {
		"status": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames": {
		"appSettingNames": []string{},
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork": {
		"swiftSupported": true,
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork": {},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web":                        {},
}

// defaultResourcesStateNormalized maps normalized paths of resources to default state of a resource. The default state is
// the configuration of that resource after its parent resource is created and before any explicit PUT operations on
// its endpoint. These resources are created automatically by Azure with their parents, therefore, the provider can
// expect to see this state at the time of the Create operation.
var defaultResourcesStateNormalized map[string]map[string]interface{}

func init() {
	defaultResourcesStateNormalized = map[string]map[string]interface{}{}
	for key, value := range defaultResourcesStateRaw {
		defaultResourcesStateNormalized[normalizePath(key)] = value
	}
}
