package azure

import (
	"context"
	"strings"
	"sync"
)

// Resource type patterns that require serialization due to exclusive lock constraints
var (
	// Web Apps require serialization due to "webspace affinity"
	webAppResourcePattern = "/providers/Microsoft.Web/sites/"
	// Add more resource type patterns here as needed
	// otherResourcePattern = "/providers/Microsoft.Other/..."
)

// Global map to track resource groups that need serialization due to exclusive lock errors.
// This is shared across all azCoreClient instances.
var (
	serializedAppServicePlans   = make(map[string]bool)
	serializedAppServicePlansMu sync.RWMutex
)

// needsSerialization checks if a resource type requires serialization.
// Returns true if the resource type is known to require serialization.
func needsSerialization(resourceID string) bool {
	switch {
	case strings.Contains(resourceID, webAppResourcePattern):
		// Web Apps require serialization due to "webspace affinity"
		return true
	// Add more resource types here as needed
	// case strings.Contains(resourceID, otherResourcePattern):
	//     return true
	default:
		return false
	}
}

// extractSerializationKeyForDelete extracts the serialization key for DELETE operations.
// Fetches resource properties via GET if needed.
// Returns the serialization key (e.g., App Service Plan ID) or empty string.
func (c *azCoreClient) extractSerializationKeyForDelete(ctx context.Context, id, apiVersion string, queryParams map[string]any) string {
	switch {
	case strings.Contains(id, webAppResourcePattern):
		// Web Apps: fetch resource to get App Service Plan ID
		webAppProps, err := c.Get(ctx, id, apiVersion, queryParams)
		if err == nil && webAppProps != nil {
			appServicePlanID := extractAppServicePlanID(id, webAppProps)
			if appServicePlanID != "" {
				return appServicePlanID
			}
		}
		// Fall back to resource group if we can't get App Service Plan ID
		return extractResourceGroupFromID(id)
	// Add more resource types here as needed
	// case strings.Contains(id, "/providers/Microsoft.Other/..."):
	//     return extractOtherResourceKeyForDelete(ctx, id, apiVersion, queryParams)
	default:
		return ""
	}
}

// extractSerializationKeyForPutOrPatch extracts the serialization key for PUT/PATCH operations.
// Uses bodyProps if available.
// Returns the serialization key (e.g., App Service Plan ID) or empty string.
func extractSerializationKeyForPutOrPatch(id string, bodyProps map[string]any) string {
	switch {
	case strings.Contains(id, webAppResourcePattern):
		// Web Apps: extract App Service Plan ID from bodyProps
		return extractAppServicePlanID(id, bodyProps)
	// Add more resource types here as needed
	// case strings.Contains(id, "/providers/Microsoft.Other/..."):
	//     return extractOtherResourceKeyForPutOrPatch(id, bodyProps)
	default:
		return ""
	}
}

// checkForSerialization returns a mutex if the resource requires serialization.
// Used for DELETE operations where we may need to fetch resource properties.
func (c *azCoreClient) checkForSerialization(ctx context.Context, id, apiVersion string, queryParams map[string]any) *sync.Mutex {
	if !needsSerialization(id) {
		return nil
	}

	serializationKey := c.extractSerializationKeyForDelete(ctx, id, apiVersion, queryParams)
	if serializationKey == "" {
		return nil
	}

	// Mark for serialization and get mutex
	serializedAppServicePlansMu.Lock()
	serializedAppServicePlans[serializationKey] = true
	serializedAppServicePlansMu.Unlock()
	return c.getAppServicePlanMutex(serializationKey)
}

// checkForSerializationPutOrPatch returns a mutex if the resource requires serialization.
// Used for PUT/PATCH operations where bodyProps are available.
func (c *azCoreClient) checkForSerializationPutOrPatch(ctx context.Context, id string, bodyProps map[string]any) *sync.Mutex {
	if !needsSerialization(id) {
		return nil
	}

	serializationKey := extractSerializationKeyForPutOrPatch(id, bodyProps)
	if serializationKey == "" {
		return nil
	}

	// Mark for serialization and get mutex
	serializedAppServicePlansMu.Lock()
	serializedAppServicePlans[serializationKey] = true
	serializedAppServicePlansMu.Unlock()
	return c.getAppServicePlanMutex(serializationKey)
}

// extractResourceGroupFromID extracts the resource group name from an Azure resource ID.
// Azure resource IDs follow the pattern:
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/...
func extractResourceGroupFromID(resourceID string) string {
	parts := strings.Split(resourceID, "/")
	for i, part := range parts {
		if part == "resourceGroups" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}

// extractAppServicePlanID extracts the App Service Plan ID from a resource ID or request body.
// For App Service Plans (serverfarms), it extracts the full resource ID.
// For Web Apps (sites), it extracts the App Service Plan ID from the request body if available.
func extractAppServicePlanID(resourceID string, bodyProps map[string]any) string {
	// Check if this is an App Service Plan (serverfarm) resource
	if strings.Contains(resourceID, "/providers/Microsoft.Web/serverfarms/") {
		// Extract the full App Service Plan resource ID
		// Format: /subscriptions/{sub}/resourceGroups/{rg}/providers/Microsoft.Web/serverfarms/{name}
		return resourceID
	}

	// For Web Apps (sites), try to extract App Service Plan ID from body
	if strings.Contains(resourceID, "/providers/Microsoft.Web/sites/") && bodyProps != nil {
		// Try common property names for App Service Plan ID
		if serverFarmId, ok := bodyProps["serverFarmId"].(string); ok && serverFarmId != "" {
			return serverFarmId
		}
		if appServicePlanId, ok := bodyProps["appServicePlanId"].(string); ok && appServicePlanId != "" {
			return appServicePlanId
		}
		// Also check nested properties
		if siteConfig, ok := bodyProps["siteConfig"].(map[string]any); ok {
			if serverFarmId, ok := siteConfig["serverFarmId"].(string); ok && serverFarmId != "" {
				return serverFarmId
			}
		}
	}

	return ""
}

// getAppServicePlanMutex returns a mutex for the given App Service Plan ID, creating it if needed.
// This mutex is used to serialize operations for App Service Plans that have hit "exclusive lock" 429 errors.
func (c *azCoreClient) getAppServicePlanMutex(appServicePlanID string) *sync.Mutex {
	if appServicePlanID == "" {
		return nil
	}

	c.aspMutexesMu.Lock()
	defer c.aspMutexesMu.Unlock()

	if mutex, ok := c.aspMutexes[appServicePlanID]; ok {
		return mutex
	}

	mutex := &sync.Mutex{}
	c.aspMutexes[appServicePlanID] = mutex
	return mutex
}
