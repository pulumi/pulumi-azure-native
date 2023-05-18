// Copyright 2021, Pulumi Corporation.  All rights reserved.

package defaults

import (
	"fmt"
	"reflect"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/openapi/paths"
)

// defaultResourcesStateRaw is a map non-normalized paths. It's handy to have paths as they are in the Open API spec's
// latest version for easy search. This map shouldn't be used for lookups: use 'defaultResourcesStateMap' instead.
var defaultResourcesStateRaw = map[string]map[string]interface{}{
	"/{resourceId}/providers/Microsoft.Security/advancedThreatProtectionSettings/{settingName}": {
		"isEnabled": false,
	},
	"/{scope}/providers/Microsoft.Resources/tags/default": {
		"properties": map[string]string{},
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMariaDB/servers/{serverName}/configurations/{configurationName}": {
		"source": "system-default",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySQL/flexibleServers/{serverName}/configurations/{configurationName}": {
		"source": "system-default",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForMySQL/servers/{serverName}/configurations/{configurationName}": {
		"source": "system-default",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSql/flexibleServers/{serverName}/configurations/{configurationName}": {
		"source": "system-default",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBForPostgreSQL/servers/{serverName}/configurations/{configurationName}": {
		"source": "system-default",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}": {
		"state": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkRuleSets/default": {
		"defaultAction":       "Deny",
		"publicNetworkAccess": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/currentbillingfeatures": {
		"currentBillingFeatures": []string{"Basic"},
		"dataVolumeCap":          map[string]string{},
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs/{ConfigurationId}": {
		"enabled": false,
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/networkRuleSets/default": {
		"defaultAction": "Deny",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors/{advisorName}": {
		"autoExecuteStatus": "Default",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/azureADOnlyAuthentications/{authenticationName}": {
		"azureADOnlyAuthentication": false,
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advisors/{advisorName}": {
		"autoExecuteStatus": "Default",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/auditingSettings/{blobAuditingPolicyName}": {
		"state": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupLongTermRetentionPolicies/{policyName}": {
		"weeklyRetention":  "PT0S",
		"monthlyRetention": "PT0S",
		"yearlyRetention":  "PT0S",
		"weekOfYear":       1,
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/backupShortTermRetentionPolicies/{policyName}": {
		"retentionDays": 7,
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/dataMaskingPolicies/{dataMaskingPolicyName}": {
		"dataMaskingState": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/extendedAuditingSettings/{blobAuditingPolicyName}": {
		"state": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/geoBackupPolicies/{geoBackupPolicyName}": {
		"state": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/securityAlertPolicies/{securityAlertPolicyName}": {
		"state": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/auditingSettings/{blobAuditingPolicyName}": {
		"state": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/encryptionProtector/{encryptionProtectorName}": {
		"serverKeyType": "ServiceManaged",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/extendedAuditingSettings/{blobAuditingPolicyName}": {
		"state": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}": {
		"state": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/vulnerabilityAssessments/{vulnerabilityAssessmentName}": {
		"recurringScans": map[string]interface{}{
			"isEnabled":               false,
			"emailSubscriptionAdmins": true,
		},
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/{BlobServicesName}": {},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}/immutabilityPolicies/{immutabilityPolicyName}": {
		"immutabilityPeriodSinceCreationInDays": 0,
		"state":                                 "Deleted",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/encryptionScopes/{encryptionScopeName}": {
		"state": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/fileServices/{FileServicesName}":  {},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/queueServices/{queueServiceName}": {},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/tableServices/{tableServiceName}": {},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/transparentDataEncryption/{transparentDataEncryptionName}": {
		"status": "Disabled",
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/authsettings": {
		"enabled": false,
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/logs": {
		"applicationLogs": map[string]string{},
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/config/slotConfigNames": {
		"appSettingNames": []string{},
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/networkConfig/virtualNetwork": {
		"swiftSupported": true,
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}/networkConfig/virtualNetwork": {
		"subnetResourceId": "*", // This is going to be a resource ID, so we choose accept any value here.
		"swiftSupported":   true,
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/sourcecontrols/web": {},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/monitoringSettings/default": {
		"properties": map[string]interface{}{
			"traceEnabled":                  false,
			"appInsightsInstrumentationKey": nil,
			"appInsightsSamplingRate":       10.0,
		},
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/default": {
		"properties": map[string]interface{}{
			"configServer": nil,
		},
	},
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/buildServices/{buildServiceName}/agentPools/{agentPoolName}": {
		"properties": map[string]interface{}{
			"poolSize": map[string]interface{}{
				"name": "S1",
			},
		},
	},
}

var skipDeleteResources = map[string]bool{
	// https://learn.microsoft.com/en-us/azure/azure-sql/database/transparent-data-encryption-tde-overview
	"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/transparentDataEncryption/{tdeName}": true,
}

// defaultResourcesStateNormalized maps normalized paths of resources to default state of a resource. The default state is
// the configuration of that resource after its parent resource is created and before any explicit PUT operations on
// its endpoint. These resources are created automatically by Azure with their parents, therefore, the provider can
// expect to see this state at the time of the Create operation.
var defaultResourcesStateNormalized map[string]DefaultResourceState

func init() {
	defaultResourcesStateNormalized = map[string]DefaultResourceState{}
	for key := range skipDeleteResources {
		addNormalisedState(key, DefaultResourceState{
			SkipDelete: true,
		})
	}
	for key, value := range defaultResourcesStateRaw {
		addNormalisedState(key, DefaultResourceState{
			State:                  value,
			HasNonEmptyCollections: containsNonEmptyCollections(value),
		})
	}
}

func addNormalisedState(path string, state DefaultResourceState) {
	normalizedPath := paths.NormalizePath(path)
	if _, exists := defaultResourcesStateNormalized[normalizedPath]; exists {
		panic(fmt.Errorf("FATAL: default state for %s is already set", normalizedPath))
	}
	defaultResourcesStateNormalized[normalizedPath] = state
}

func containsNonEmptyCollections(value map[string]interface{}) bool {
	for _, propValue := range value {
		if propValue == nil {
			continue
		}
		typ := reflect.TypeOf(propValue)
		switch typ.Kind() {
		case reflect.Slice, reflect.Array:
			val := reflect.ValueOf(propValue)
			if val.Len() > 0 {
				return true
			}
		}
	}
	return false
}

type DefaultResourceState struct {
	// State is the default state of a resource.
	State map[string]interface{}
	// SkipDelete is true if the PUT operation should be skipped during the delete step.
	SkipDelete bool
	// HasNonEmptyCollections is true if any property of the default state is a non-empty collection.
	HasNonEmptyCollections bool
}

func GetDefaultResourceState(path string) *DefaultResourceState {
	normalizedPath := paths.NormalizePath(path)
	defaults, ok := defaultResourcesStateNormalized[normalizedPath]
	if !ok {
		return nil
	}
	return &defaults
}

func SkipDeleteOperation(path string) bool {
	defaultState := GetDefaultResourceState(path)
	if defaultState == nil {
		return false
	}
	return defaultState.SkipDelete
}
