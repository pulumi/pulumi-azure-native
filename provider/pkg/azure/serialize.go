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
			appServicePlanID := extractAppServicePlanID(webAppProps)
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
		return extractAppServicePlanID(bodyProps)
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

	// Get mutex for serialization
	return c.getSerializationMutex(serializationKey)
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

	// Get mutex for serialization
	return c.getSerializationMutex(serializationKey)
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

// extractAppServicePlanID extracts the App Service Plan ID from a request body.
// For Web Apps (sites), it extracts the App Service Plan ID from the request body if available.
func extractAppServicePlanID(bodyProps map[string]any) string {
	// For Web Apps (sites), try to extract App Service Plan ID from body
	if bodyProps != nil {
		// Try common property names for App Service Plan ID
		if appServicePlanId, ok := bodyProps["appServicePlanId"].(string); ok && appServicePlanId != "" {
			return appServicePlanId
		}
	}
	return ""
}

// getSerializationMutex returns a mutex for the given serialization key, creating it if needed.
// This mutex is used to serialize operations for resources that require serialization (e.g., due to "exclusive lock" 429 errors).
// The serializationKey can be an App Service Plan ID, resource group, or any other identifier that groups resources requiring serialization.
func (c *azCoreClient) getSerializationMutex(serializationKey string) *sync.Mutex {
	if serializationKey == "" {
		return nil
	}

	c.aspMutexesMu.Lock()
	defer c.aspMutexesMu.Unlock()

	if mutex, ok := c.aspMutexes[serializationKey]; ok {
		return mutex
	}

	mutex := &sync.Mutex{}
	c.aspMutexes[serializationKey] = mutex
	return mutex
}
