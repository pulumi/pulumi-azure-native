package paths

import "strings"

// Sometimes, Azure resources change the API paths between API versions. Most of the time, we can detect that based
// on operation names. However, in a number of cases, the operation names change at the same time.
// legacyPathMappings provides a manual map to help our codegen discover the old "aliases" of new resource paths
// and group them under the same Pulumi resource, including proper top-level resource calculation and aliasing.
var legacyPathMappings = map[string]string{
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/buildTasks/{buildTaskName}":                                       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/tasks/{taskName}",
	"/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.CostManagement/connectors/{connectorName}":                                                                    "/providers/Microsoft.CostManagement/cloudConnectors/{connectorName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}":                            "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/cassandra/keyspaces/{keyspaceName}/tables/{tableName}":         "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/cassandraKeyspaces/{keyspaceName}/tables/{tableName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}":                              "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/gremlin/databases/{databaseName}/graphs/{graphName}":           "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/gremlinDatabases/{databaseName}/graphs/{graphName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}":                              "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/mongodb/databases/{databaseName}/collections/{collectionName}": "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/collections/{collectionName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}":                                  "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/sql/databases/{databaseName}/containers/{containerName}":       "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/apis/table/tables/{tableName}":                                      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/interfaceEndpoints/{interfaceEndpointName}":                                                           "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}",
	"/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/listKeys":                                                      "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/sharedKeys",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/appliances/{applianceName}":                                                                         "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applications/{applicationName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applianceDefinitions/{applianceDefinitionName}":                                                     "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Solutions/applicationDefinitions/{applicationDefinitionName}",
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/managedHostingEnvironments/{name}":                                                                        "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/hostingEnvironments/{name}",
}

var legacyPathMappingNormalized = map[string]string{}

func init() {
	for key, value := range legacyPathMappings {
		legacyPathMappingNormalized[NormalizePath(key)] = NormalizePath(value)
	}
}

// NormalizePath converts an API path to its canonical form (lowercase, with all placeholders removed). The paths that
// convert to the same canonical path are considered to represent the same resource.
func NormalizePath(path string) string {
	lowerPath := strings.ReplaceAll(strings.ToLower(strings.TrimSuffix(path, "/")), "-", "")

	// Work around an odd version v2019-01-01-preview of SecurityInsights where they have a parameter for the provider.
	// This breaks all path matching for that version which includes quite a lot of resources. Instead of providing
	// a value per each resource, let's replace this path segment while normalizing.
	lowerPath = strings.ReplaceAll(lowerPath, "providers/{operationalinsightsresourceprovider}", "providers/microsoft.operationalinsights")

	parts := strings.Split(lowerPath, "/")
	newParts := make([]string, len(parts))
	for i, part := range parts {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			newParts[i] = "{}"
		} else {
			newParts[i] = part
		}
	}
	normalized := strings.Join(newParts, "/")
	if override, ok := legacyPathMappingNormalized[normalized]; ok {
		return override
	}
	return normalized
}
