package azure

import (
	"context"
	"strings"
	"sync"
)

// Resource type patterns that require serialization due to exclusive lock constraints
var (
	// Web Apps require serialization due to limitations on handling hardware infrastucture in the AZ Core SDK.
	webAppResourcePattern = "/providers/Microsoft.Web/sites/"
)

func needsSerialization(resourceID string) bool {
	switch {
	case strings.Contains(resourceID, webAppResourcePattern):
		return true
	default:
		return false
	}
}

// extractSerializationKeyForDelete extracts the serialization key for DELETE operations.
func (c *azCoreClient) extractSerializationKeyForDelete(ctx context.Context, id, apiVersion string, queryParams map[string]any) string {
	switch {
	case strings.Contains(id, webAppResourcePattern):
		// Web Apps: fetch resource to get App Service Plan ID
		webAppProps, err := c.Get(ctx, id, apiVersion, queryParams)
		if err == nil && webAppProps != nil {
			appServicePlanID := extractAppServicePlanID(webAppProps)
			if appServicePlanID != "" {
				return appServicePlanID
			}
		}
		// Fall back to resource group if we can't get App Service Plan ID
		return extractResourceGroupFromID(id)
	default:
		return ""
	}
}

// extractSerializationKeyForPutOrPatch extracts the serialization key for PUT/PATCH operations using bodyProps if available.
func extractSerializationKeyForPutOrPatch(id string, bodyProps map[string]any) string {
	switch {
	case strings.Contains(id, webAppResourcePattern):
		// Web Apps: extract App Service Plan ID from bodyProps
		return extractAppServicePlanID(bodyProps)
	default:
		return ""
	}
}

func extractResourceGroupFromID(resourceID string) string {
	// Azure resource IDs follow the pattern:
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/...
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if part == "resourceGroups" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

// extractAppServicePlanID extracts the App Service Plan ID from a request body or resource properties.
// For Web Apps, the App Service Plan ID is stored at properties.serverFarmId.
func extractAppServicePlanID(bodyProps map[string]any) string {
	if bodyProps == nil {
		return ""
	}
	// properties.serverFarmId is the App Service Plan.
	if properties, ok := bodyProps["properties"].(map[string]any); ok {
		if serverFarmId, ok := properties["serverFarmId"].(string); ok && serverFarmId != "" {
			return serverFarmId
		}
	}
	return ""
}

// getSerializationMutex returns a mutex for the given serialization key, creating it if needed.
// This mutex is used to serialize operations for resources that require serialization.
func (c *azCoreClient) getSerializationMutex(serializationKey string) *sync.Mutex {
	if serializationKey == "" {
		return nil
	}

	c.serializationsMapMutex.Lock()
	defer c.serializationsMapMutex.Unlock()

	if mutex, ok := c.serializationsMap[serializationKey]; ok {
		return mutex
	}

	mutex := &sync.Mutex{}
	c.serializationsMap[serializationKey] = mutex
	return mutex
}

// checkForSerializationDelete returns a mutex if the resource requires serialization.
func (c *azCoreClient) checkForSerializationDelete(ctx context.Context, id, apiVersion string, queryParams map[string]any) *sync.Mutex {
	if !needsSerialization(id) {
		return nil
	}

	serializationKey := c.extractSerializationKeyForDelete(ctx, id, apiVersion, queryParams)
	if serializationKey == "" {
		return nil
	}

	// Get mutex for this specific serialization key (e.g., App Service Plan ID)
	return c.getSerializationMutex(serializationKey)
}

// checkForSerializationPutOrPatch returns a mutex if the resource requires serialization.
func (c *azCoreClient) checkForSerializationPutOrPatch(id string, bodyProps map[string]any) *sync.Mutex {
	if !needsSerialization(id) {
		return nil
	}

	serializationKey := extractSerializationKeyForPutOrPatch(id, bodyProps)
	if serializationKey == "" {
		return nil
	}

	// Get mutex for this specific serialization key (e.g., App Service Plan ID)
	return c.getSerializationMutex(serializationKey)
}
