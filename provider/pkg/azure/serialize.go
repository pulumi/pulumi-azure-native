package azure

import (
	"context"
	"strings"
	"sync"
)

// Global map to track resource groups that need serialization due to exclusive lock errors.
// This is shared across all azCoreClient instances.
var (
	serializedAppServicePlans   = make(map[string]bool)
	serializedAppServicePlansMu sync.RWMutex
)

func (c *azCoreClient) checkForSerialization(ctx context.Context, id, apiVersion string, queryParams map[string]any) sync.Mutex {
	appServicePlanID := ""
	var mutex *sync.Mutex
	if strings.Contains(id, "/providers/Microsoft.Web/sites/") {
		// For Web Apps, fetch the resource first to get the App Service Plan ID
		webAppProps, err := c.Get(ctx, id, apiVersion, queryParams)
		if err == nil && webAppProps != nil {
			appServicePlanID = extractAppServicePlanID(id, webAppProps)
		}

		// If we couldn't get the App Service Plan ID, fall back to resource group
		if appServicePlanID == "" {
			resourceGroup := extractResourceGroupFromID(id)
			if resourceGroup != "" {
				appServicePlanID = resourceGroup
			}
		}

		// Proactively mark for serialization and get mutex
		if appServicePlanID != "" {
			serializedAppServicePlansMu.Lock()
			serializedAppServicePlans[appServicePlanID] = true
			serializedAppServicePlansMu.Unlock()
			mutex = c.getAppServicePlanMutex(appServicePlanID)
		}
	}
	return *mutex
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

// needsSerializationForWebResources checks if a resource ID is for a web resource that requires serialization.
// Web resources (App Service Plans, Web Apps) have "webspace affinity" and can only process one operation at a time.
func needsSerializationForWebResources(resourceID string) bool {
	// Check if this is a Microsoft.Web resource
	if strings.Contains(resourceID, "/providers/Microsoft.Web/") {
		return true
	}
	return false
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
