


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureBackupServerContainer struct {
	BackupManagementType *string                   `pulumi:"backupManagementType"`
	CanReRegister        *bool                     `pulumi:"canReRegister"`
	ContainerId          *string                   `pulumi:"containerId"`
	ContainerType        string                    `pulumi:"containerType"`
	DpmAgentVersion      *string                   `pulumi:"dpmAgentVersion"`
	DpmServers           []string                  `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                   `pulumi:"friendlyName"`
	HealthStatus         *string                   `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                  `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                   `pulumi:"protectionStatus"`
	RegistrationStatus   *string                   `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                     `pulumi:"upgradeAvailable"`
}

type AzureBackupServerContainerResponse struct {
	BackupManagementType *string                           `pulumi:"backupManagementType"`
	CanReRegister        *bool                             `pulumi:"canReRegister"`
	ContainerId          *string                           `pulumi:"containerId"`
	ContainerType        string                            `pulumi:"containerType"`
	DpmAgentVersion      *string                           `pulumi:"dpmAgentVersion"`
	DpmServers           []string                          `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                           `pulumi:"friendlyName"`
	HealthStatus         *string                           `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                          `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                           `pulumi:"protectionStatus"`
	RegistrationStatus   *string                           `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                             `pulumi:"upgradeAvailable"`
}

type AzureFileShareProtectionPolicy struct {
	BackupManagementType string      `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy       interface{} `pulumi:"schedulePolicy"`
	TimeZone             *string     `pulumi:"timeZone"`
	WorkLoadType         *string     `pulumi:"workLoadType"`
}

type AzureFileShareProtectionPolicyResponse struct {
	BackupManagementType string      `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy       interface{} `pulumi:"schedulePolicy"`
	TimeZone             *string     `pulumi:"timeZone"`
	WorkLoadType         *string     `pulumi:"workLoadType"`
}

type AzureFileshareProtectedItem struct {
	BackupManagementType             *string                                  `pulumi:"backupManagementType"`
	BackupSetName                    *string                                  `pulumi:"backupSetName"`
	ContainerName                    *string                                  `pulumi:"containerName"`
	CreateMode                       *string                                  `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                  `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                  `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureFileshareProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                  `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                    `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                    `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                    `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails      `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                  `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                  `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                  `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                  `pulumi:"policyId"`
	ProtectedItemType                string                                   `pulumi:"protectedItemType"`
	ProtectionState                  *string                                  `pulumi:"protectionState"`
	ProtectionStatus                 *string                                  `pulumi:"protectionStatus"`
	SourceResourceId                 *string                                  `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                  `pulumi:"workloadType"`
}

type AzureFileshareProtectedItemExtendedInfo struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyState         *string `pulumi:"policyState"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}

type AzureFileshareProtectedItemExtendedInfoResponse struct {
	OldestRecoveryPoint   *string `pulumi:"oldestRecoveryPoint"`
	PolicyState           *string `pulumi:"policyState"`
	RecoveryPointCount    *int    `pulumi:"recoveryPointCount"`
	ResourceState         string  `pulumi:"resourceState"`
	ResourceStateSyncTime string  `pulumi:"resourceStateSyncTime"`
}

type AzureFileshareProtectedItemResponse struct {
	BackupManagementType             *string                                          `pulumi:"backupManagementType"`
	BackupSetName                    *string                                          `pulumi:"backupSetName"`
	ContainerName                    *string                                          `pulumi:"containerName"`
	CreateMode                       *string                                          `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                          `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                          `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureFileshareProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                          `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                            `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                            `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                            `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse      `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                          `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                          `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                          `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                          `pulumi:"policyId"`
	ProtectedItemType                string                                           `pulumi:"protectedItemType"`
	ProtectionState                  *string                                          `pulumi:"protectionState"`
	ProtectionStatus                 *string                                          `pulumi:"protectionStatus"`
	SourceResourceId                 *string                                          `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                          `pulumi:"workloadType"`
}

type AzureIaaSClassicComputeVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type AzureIaaSClassicComputeVMContainerResponse struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type AzureIaaSClassicComputeVMProtectedItem struct {
	BackupManagementType             *string                               `pulumi:"backupManagementType"`
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedProperties                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                               `pulumi:"friendlyName"`
	HealthStatus                     *string                               `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                               `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                               `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	ProtectedItemDataId              *string                               `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ProtectionStatus                 *string                               `pulumi:"protectionStatus"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                               `pulumi:"virtualMachineId"`
	WorkloadType                     *string                               `pulumi:"workloadType"`
}

type AzureIaaSClassicComputeVMProtectedItemResponse struct {
	BackupManagementType             *string                                       `pulumi:"backupManagementType"`
	BackupSetName                    *string                                       `pulumi:"backupSetName"`
	ContainerName                    *string                                       `pulumi:"containerName"`
	CreateMode                       *string                                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedPropertiesResponse                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                                       `pulumi:"friendlyName"`
	HealthDetails                    []AzureIaaSVMHealthDetailsResponse            `pulumi:"healthDetails"`
	HealthStatus                     *string                                       `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                         `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                       `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                       `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                       `pulumi:"policyId"`
	ProtectedItemDataId              *string                                       `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                                       `pulumi:"protectionState"`
	ProtectionStatus                 *string                                       `pulumi:"protectionStatus"`
	SourceResourceId                 *string                                       `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                                       `pulumi:"virtualMachineId"`
	WorkloadType                     *string                                       `pulumi:"workloadType"`
}

type AzureIaaSComputeVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type AzureIaaSComputeVMContainerResponse struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type AzureIaaSComputeVMProtectedItem struct {
	BackupManagementType             *string                               `pulumi:"backupManagementType"`
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedProperties                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                               `pulumi:"friendlyName"`
	HealthStatus                     *string                               `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                               `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                               `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	ProtectedItemDataId              *string                               `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ProtectionStatus                 *string                               `pulumi:"protectionStatus"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                               `pulumi:"virtualMachineId"`
	WorkloadType                     *string                               `pulumi:"workloadType"`
}

type AzureIaaSComputeVMProtectedItemResponse struct {
	BackupManagementType             *string                                       `pulumi:"backupManagementType"`
	BackupSetName                    *string                                       `pulumi:"backupSetName"`
	ContainerName                    *string                                       `pulumi:"containerName"`
	CreateMode                       *string                                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedPropertiesResponse                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                                       `pulumi:"friendlyName"`
	HealthDetails                    []AzureIaaSVMHealthDetailsResponse            `pulumi:"healthDetails"`
	HealthStatus                     *string                                       `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                         `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                       `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                       `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                       `pulumi:"policyId"`
	ProtectedItemDataId              *string                                       `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                                       `pulumi:"protectionState"`
	ProtectionStatus                 *string                                       `pulumi:"protectionStatus"`
	SourceResourceId                 *string                                       `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                                       `pulumi:"virtualMachineId"`
	WorkloadType                     *string                                       `pulumi:"workloadType"`
}

type AzureIaaSVMHealthDetailsResponse struct {
	Code            int      `pulumi:"code"`
	Message         string   `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           string   `pulumi:"title"`
}

type AzureIaaSVMProtectedItem struct {
	BackupManagementType             *string                               `pulumi:"backupManagementType"`
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedProperties                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                               `pulumi:"friendlyName"`
	HealthStatus                     *string                               `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                               `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                               `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	ProtectedItemDataId              *string                               `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ProtectionStatus                 *string                               `pulumi:"protectionStatus"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                               `pulumi:"virtualMachineId"`
	WorkloadType                     *string                               `pulumi:"workloadType"`
}

type AzureIaaSVMProtectedItemExtendedInfo struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyInconsistent  *bool   `pulumi:"policyInconsistent"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}

type AzureIaaSVMProtectedItemExtendedInfoResponse struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyInconsistent  *bool   `pulumi:"policyInconsistent"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}

type AzureIaaSVMProtectedItemResponse struct {
	BackupManagementType             *string                                       `pulumi:"backupManagementType"`
	BackupSetName                    *string                                       `pulumi:"backupSetName"`
	ContainerName                    *string                                       `pulumi:"containerName"`
	CreateMode                       *string                                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedPropertiesResponse                   `pulumi:"extendedProperties"`
	FriendlyName                     *string                                       `pulumi:"friendlyName"`
	HealthDetails                    []AzureIaaSVMHealthDetailsResponse            `pulumi:"healthDetails"`
	HealthStatus                     *string                                       `pulumi:"healthStatus"`
	IsDeferredDeleteScheduleUpcoming *bool                                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                         `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                       `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                       `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                       `pulumi:"policyId"`
	ProtectedItemDataId              *string                                       `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                                       `pulumi:"protectionState"`
	ProtectionStatus                 *string                                       `pulumi:"protectionStatus"`
	SourceResourceId                 *string                                       `pulumi:"sourceResourceId"`
	VirtualMachineId                 *string                                       `pulumi:"virtualMachineId"`
	WorkloadType                     *string                                       `pulumi:"workloadType"`
}

type AzureIaaSVMProtectionPolicy struct {
	BackupManagementType          string                      `pulumi:"backupManagementType"`
	InstantRPDetails              *InstantRPAdditionalDetails `pulumi:"instantRPDetails"`
	InstantRpRetentionRangeInDays *int                        `pulumi:"instantRpRetentionRangeInDays"`
	ProtectedItemsCount           *int                        `pulumi:"protectedItemsCount"`
	RetentionPolicy               interface{}                 `pulumi:"retentionPolicy"`
	SchedulePolicy                interface{}                 `pulumi:"schedulePolicy"`
	TimeZone                      *string                     `pulumi:"timeZone"`
}

type AzureIaaSVMProtectionPolicyResponse struct {
	BackupManagementType          string                              `pulumi:"backupManagementType"`
	InstantRPDetails              *InstantRPAdditionalDetailsResponse `pulumi:"instantRPDetails"`
	InstantRpRetentionRangeInDays *int                                `pulumi:"instantRpRetentionRangeInDays"`
	ProtectedItemsCount           *int                                `pulumi:"protectedItemsCount"`
	RetentionPolicy               interface{}                         `pulumi:"retentionPolicy"`
	SchedulePolicy                interface{}                         `pulumi:"schedulePolicy"`
	TimeZone                      *string                             `pulumi:"timeZone"`
}

type AzureRecoveryServiceVaultProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}

type AzureRecoveryServiceVaultProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}

type AzureResourceProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	FriendlyName             *string `pulumi:"friendlyName"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}

type AzureResourceProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	FriendlyName             *string `pulumi:"friendlyName"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}

type AzureSQLAGWorkloadContainerProtectionContainer struct {
	BackupManagementType *string                             `pulumi:"backupManagementType"`
	ContainerType        string                              `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                             `pulumi:"friendlyName"`
	HealthStatus         *string                             `pulumi:"healthStatus"`
	LastUpdatedTime      *string                             `pulumi:"lastUpdatedTime"`
	OperationType        *string                             `pulumi:"operationType"`
	RegistrationStatus   *string                             `pulumi:"registrationStatus"`
	SourceResourceId     *string                             `pulumi:"sourceResourceId"`
	WorkloadType         *string                             `pulumi:"workloadType"`
}

type AzureSQLAGWorkloadContainerProtectionContainerResponse struct {
	BackupManagementType *string                                     `pulumi:"backupManagementType"`
	ContainerType        string                                      `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                                     `pulumi:"friendlyName"`
	HealthStatus         *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime      *string                                     `pulumi:"lastUpdatedTime"`
	OperationType        *string                                     `pulumi:"operationType"`
	RegistrationStatus   *string                                     `pulumi:"registrationStatus"`
	SourceResourceId     *string                                     `pulumi:"sourceResourceId"`
	WorkloadType         *string                                     `pulumi:"workloadType"`
}

type AzureSqlContainer struct {
	BackupManagementType *string `pulumi:"backupManagementType"`
	ContainerType        string  `pulumi:"containerType"`
	FriendlyName         *string `pulumi:"friendlyName"`
	HealthStatus         *string `pulumi:"healthStatus"`
	RegistrationStatus   *string `pulumi:"registrationStatus"`
}

type AzureSqlContainerResponse struct {
	BackupManagementType *string `pulumi:"backupManagementType"`
	ContainerType        string  `pulumi:"containerType"`
	FriendlyName         *string `pulumi:"friendlyName"`
	HealthStatus         *string `pulumi:"healthStatus"`
	RegistrationStatus   *string `pulumi:"registrationStatus"`
}

type AzureSqlProtectedItem struct {
	BackupManagementType             *string                            `pulumi:"backupManagementType"`
	BackupSetName                    *string                            `pulumi:"backupSetName"`
	ContainerName                    *string                            `pulumi:"containerName"`
	CreateMode                       *string                            `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                            `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                            `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureSqlProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	IsDeferredDeleteScheduleUpcoming *bool                              `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                              `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                              `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                            `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                            `pulumi:"policyId"`
	ProtectedItemDataId              *string                            `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                             `pulumi:"protectedItemType"`
	ProtectionState                  *string                            `pulumi:"protectionState"`
	SourceResourceId                 *string                            `pulumi:"sourceResourceId"`
	WorkloadType                     *string                            `pulumi:"workloadType"`
}

type AzureSqlProtectedItemExtendedInfo struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyState         *string `pulumi:"policyState"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}

type AzureSqlProtectedItemExtendedInfoResponse struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyState         *string `pulumi:"policyState"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}

type AzureSqlProtectedItemResponse struct {
	BackupManagementType             *string                                    `pulumi:"backupManagementType"`
	BackupSetName                    *string                                    `pulumi:"backupSetName"`
	ContainerName                    *string                                    `pulumi:"containerName"`
	CreateMode                       *string                                    `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                    `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                    `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureSqlProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	IsDeferredDeleteScheduleUpcoming *bool                                      `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                      `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                      `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                                    `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                    `pulumi:"policyId"`
	ProtectedItemDataId              *string                                    `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                     `pulumi:"protectedItemType"`
	ProtectionState                  *string                                    `pulumi:"protectionState"`
	SourceResourceId                 *string                                    `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                    `pulumi:"workloadType"`
}

type AzureSqlProtectionPolicy struct {
	BackupManagementType string      `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
}

type AzureSqlProtectionPolicyResponse struct {
	BackupManagementType string      `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
}

type AzureStorageContainer struct {
	BackupManagementType  *string  `pulumi:"backupManagementType"`
	ContainerType         string   `pulumi:"containerType"`
	FriendlyName          *string  `pulumi:"friendlyName"`
	HealthStatus          *string  `pulumi:"healthStatus"`
	ProtectedItemCount    *float64 `pulumi:"protectedItemCount"`
	RegistrationStatus    *string  `pulumi:"registrationStatus"`
	ResourceGroup         *string  `pulumi:"resourceGroup"`
	SourceResourceId      *string  `pulumi:"sourceResourceId"`
	StorageAccountVersion *string  `pulumi:"storageAccountVersion"`
}

type AzureStorageContainerResponse struct {
	BackupManagementType  *string  `pulumi:"backupManagementType"`
	ContainerType         string   `pulumi:"containerType"`
	FriendlyName          *string  `pulumi:"friendlyName"`
	HealthStatus          *string  `pulumi:"healthStatus"`
	ProtectedItemCount    *float64 `pulumi:"protectedItemCount"`
	RegistrationStatus    *string  `pulumi:"registrationStatus"`
	ResourceGroup         *string  `pulumi:"resourceGroup"`
	SourceResourceId      *string  `pulumi:"sourceResourceId"`
	StorageAccountVersion *string  `pulumi:"storageAccountVersion"`
}

type AzureVMAppContainerProtectionContainer struct {
	BackupManagementType *string                             `pulumi:"backupManagementType"`
	ContainerType        string                              `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                             `pulumi:"friendlyName"`
	HealthStatus         *string                             `pulumi:"healthStatus"`
	LastUpdatedTime      *string                             `pulumi:"lastUpdatedTime"`
	OperationType        *string                             `pulumi:"operationType"`
	RegistrationStatus   *string                             `pulumi:"registrationStatus"`
	SourceResourceId     *string                             `pulumi:"sourceResourceId"`
	WorkloadType         *string                             `pulumi:"workloadType"`
}

type AzureVMAppContainerProtectionContainerResponse struct {
	BackupManagementType *string                                     `pulumi:"backupManagementType"`
	ContainerType        string                                      `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                                     `pulumi:"friendlyName"`
	HealthStatus         *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime      *string                                     `pulumi:"lastUpdatedTime"`
	OperationType        *string                                     `pulumi:"operationType"`
	RegistrationStatus   *string                                     `pulumi:"registrationStatus"`
	SourceResourceId     *string                                     `pulumi:"sourceResourceId"`
	WorkloadType         *string                                     `pulumi:"workloadType"`
}

type AzureVmWorkloadProtectedItem struct {
	BackupManagementType             *string                                   `pulumi:"backupManagementType"`
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                     `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails       `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                   `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                   `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                   `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                   `pulumi:"parentName"`
	ParentType                       *string                                   `pulumi:"parentType"`
	PolicyId                         *string                                   `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ProtectionStatus                 *string                                   `pulumi:"protectionStatus"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                   `pulumi:"workloadType"`
}

type AzureVmWorkloadProtectedItemExtendedInfo struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyState         *string `pulumi:"policyState"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}

type AzureVmWorkloadProtectedItemExtendedInfoResponse struct {
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	PolicyState         *string `pulumi:"policyState"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}

type AzureVmWorkloadProtectedItemResponse struct {
	BackupManagementType             *string                                           `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                             `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse       `pulumi:"kpisHealths"`
	LastBackupErrorDetail            *ErrorDetailResponse                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 *string                                           `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                           `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                           `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                           `pulumi:"parentName"`
	ParentType                       *string                                           `pulumi:"parentType"`
	PolicyId                         *string                                           `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 *string                                           `pulumi:"protectionStatus"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                           `pulumi:"workloadType"`
}

type AzureVmWorkloadProtectionPolicy struct {
	BackupManagementType string                `pulumi:"backupManagementType"`
	MakePolicyConsistent *bool                 `pulumi:"makePolicyConsistent"`
	ProtectedItemsCount  *int                  `pulumi:"protectedItemsCount"`
	Settings             *Settings             `pulumi:"settings"`
	SubProtectionPolicy  []SubProtectionPolicy `pulumi:"subProtectionPolicy"`
	WorkLoadType         *string               `pulumi:"workLoadType"`
}

type AzureVmWorkloadProtectionPolicyResponse struct {
	BackupManagementType string                        `pulumi:"backupManagementType"`
	MakePolicyConsistent *bool                         `pulumi:"makePolicyConsistent"`
	ProtectedItemsCount  *int                          `pulumi:"protectedItemsCount"`
	Settings             *SettingsResponse             `pulumi:"settings"`
	SubProtectionPolicy  []SubProtectionPolicyResponse `pulumi:"subProtectionPolicy"`
	WorkLoadType         *string                       `pulumi:"workLoadType"`
}

type AzureVmWorkloadSAPAseDatabaseProtectedItem struct {
	BackupManagementType             *string                                   `pulumi:"backupManagementType"`
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                     `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails       `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                   `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                   `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                   `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                   `pulumi:"parentName"`
	ParentType                       *string                                   `pulumi:"parentType"`
	PolicyId                         *string                                   `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ProtectionStatus                 *string                                   `pulumi:"protectionStatus"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                   `pulumi:"workloadType"`
}

type AzureVmWorkloadSAPAseDatabaseProtectedItemResponse struct {
	BackupManagementType             *string                                           `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                             `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse       `pulumi:"kpisHealths"`
	LastBackupErrorDetail            *ErrorDetailResponse                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 *string                                           `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                           `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                           `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                           `pulumi:"parentName"`
	ParentType                       *string                                           `pulumi:"parentType"`
	PolicyId                         *string                                           `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 *string                                           `pulumi:"protectionStatus"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                           `pulumi:"workloadType"`
}

type AzureVmWorkloadSAPHanaDatabaseProtectedItem struct {
	BackupManagementType             *string                                   `pulumi:"backupManagementType"`
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                     `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails       `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                   `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                   `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                   `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                   `pulumi:"parentName"`
	ParentType                       *string                                   `pulumi:"parentType"`
	PolicyId                         *string                                   `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ProtectionStatus                 *string                                   `pulumi:"protectionStatus"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                   `pulumi:"workloadType"`
}

type AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse struct {
	BackupManagementType             *string                                           `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                             `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse       `pulumi:"kpisHealths"`
	LastBackupErrorDetail            *ErrorDetailResponse                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 *string                                           `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                           `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                           `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                           `pulumi:"parentName"`
	ParentType                       *string                                           `pulumi:"parentType"`
	PolicyId                         *string                                           `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 *string                                           `pulumi:"protectionStatus"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                           `pulumi:"workloadType"`
}

type AzureVmWorkloadSQLDatabaseProtectedItem struct {
	BackupManagementType             *string                                   `pulumi:"backupManagementType"`
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                   `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                     `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                     `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                     `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails       `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                   `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                   `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                   `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                   `pulumi:"parentName"`
	ParentType                       *string                                   `pulumi:"parentType"`
	PolicyId                         *string                                   `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ProtectionStatus                 *string                                   `pulumi:"protectionStatus"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                   `pulumi:"workloadType"`
}

type AzureVmWorkloadSQLDatabaseProtectedItemResponse struct {
	BackupManagementType             *string                                           `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                             `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse       `pulumi:"kpisHealths"`
	LastBackupErrorDetail            *ErrorDetailResponse                              `pulumi:"lastBackupErrorDetail"`
	LastBackupStatus                 *string                                           `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                           `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                           `pulumi:"lastRecoveryPoint"`
	ParentName                       *string                                           `pulumi:"parentName"`
	ParentType                       *string                                           `pulumi:"parentType"`
	PolicyId                         *string                                           `pulumi:"policyId"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 *string                                           `pulumi:"protectionStatus"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                           `pulumi:"workloadType"`
}

type AzureWorkloadAutoProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}

type AzureWorkloadAutoProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}

type AzureWorkloadContainer struct {
	BackupManagementType *string                             `pulumi:"backupManagementType"`
	ContainerType        string                              `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                             `pulumi:"friendlyName"`
	HealthStatus         *string                             `pulumi:"healthStatus"`
	LastUpdatedTime      *string                             `pulumi:"lastUpdatedTime"`
	OperationType        *string                             `pulumi:"operationType"`
	RegistrationStatus   *string                             `pulumi:"registrationStatus"`
	SourceResourceId     *string                             `pulumi:"sourceResourceId"`
	WorkloadType         *string                             `pulumi:"workloadType"`
}

type AzureWorkloadContainerExtendedInfo struct {
	HostServerName *string                `pulumi:"hostServerName"`
	InquiryInfo    *InquiryInfo           `pulumi:"inquiryInfo"`
	NodesList      []DistributedNodesInfo `pulumi:"nodesList"`
}

type AzureWorkloadContainerExtendedInfoResponse struct {
	HostServerName *string                        `pulumi:"hostServerName"`
	InquiryInfo    *InquiryInfoResponse           `pulumi:"inquiryInfo"`
	NodesList      []DistributedNodesInfoResponse `pulumi:"nodesList"`
}

type AzureWorkloadContainerResponse struct {
	BackupManagementType *string                                     `pulumi:"backupManagementType"`
	ContainerType        string                                      `pulumi:"containerType"`
	ExtendedInfo         *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                                     `pulumi:"friendlyName"`
	HealthStatus         *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime      *string                                     `pulumi:"lastUpdatedTime"`
	OperationType        *string                                     `pulumi:"operationType"`
	RegistrationStatus   *string                                     `pulumi:"registrationStatus"`
	SourceResourceId     *string                                     `pulumi:"sourceResourceId"`
	WorkloadType         *string                                     `pulumi:"workloadType"`
}

type AzureWorkloadSQLAutoProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
	WorkloadItemType         *string `pulumi:"workloadItemType"`
}

type AzureWorkloadSQLAutoProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
	WorkloadItemType         *string `pulumi:"workloadItemType"`
}

type ContainerIdentityInfo struct {
	AadTenantId              *string `pulumi:"aadTenantId"`
	Audience                 *string `pulumi:"audience"`
	ServicePrincipalClientId *string `pulumi:"servicePrincipalClientId"`
	UniqueName               *string `pulumi:"uniqueName"`
}

type ContainerIdentityInfoResponse struct {
	AadTenantId              *string `pulumi:"aadTenantId"`
	Audience                 *string `pulumi:"audience"`
	ServicePrincipalClientId *string `pulumi:"servicePrincipalClientId"`
	UniqueName               *string `pulumi:"uniqueName"`
}

type DPMContainerExtendedInfo struct {
	LastRefreshedAt *string `pulumi:"lastRefreshedAt"`
}

type DPMContainerExtendedInfoResponse struct {
	LastRefreshedAt *string `pulumi:"lastRefreshedAt"`
}

type DPMProtectedItem struct {
	BackupEngineName                 *string                       `pulumi:"backupEngineName"`
	BackupManagementType             *string                       `pulumi:"backupManagementType"`
	BackupSetName                    *string                       `pulumi:"backupSetName"`
	ContainerName                    *string                       `pulumi:"containerName"`
	CreateMode                       *string                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *DPMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                       `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                         `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                       `pulumi:"policyId"`
	ProtectedItemType                string                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                       `pulumi:"protectionState"`
	SourceResourceId                 *string                       `pulumi:"sourceResourceId"`
	WorkloadType                     *string                       `pulumi:"workloadType"`
}

type DPMProtectedItemExtendedInfo struct {
	DiskStorageUsedInBytes       *string           `pulumi:"diskStorageUsedInBytes"`
	IsCollocated                 *bool             `pulumi:"isCollocated"`
	IsPresentOnCloud             *bool             `pulumi:"isPresentOnCloud"`
	LastBackupStatus             *string           `pulumi:"lastBackupStatus"`
	LastRefreshedAt              *string           `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint          *string           `pulumi:"oldestRecoveryPoint"`
	OnPremiseLatestRecoveryPoint *string           `pulumi:"onPremiseLatestRecoveryPoint"`
	OnPremiseOldestRecoveryPoint *string           `pulumi:"onPremiseOldestRecoveryPoint"`
	OnPremiseRecoveryPointCount  *int              `pulumi:"onPremiseRecoveryPointCount"`
	ProtectableObjectLoadPath    map[string]string `pulumi:"protectableObjectLoadPath"`
	Protected                    *bool             `pulumi:"protected"`
	ProtectionGroupName          *string           `pulumi:"protectionGroupName"`
	RecoveryPointCount           *int              `pulumi:"recoveryPointCount"`
	TotalDiskStorageSizeInBytes  *string           `pulumi:"totalDiskStorageSizeInBytes"`
}

type DPMProtectedItemExtendedInfoResponse struct {
	DiskStorageUsedInBytes       *string           `pulumi:"diskStorageUsedInBytes"`
	IsCollocated                 *bool             `pulumi:"isCollocated"`
	IsPresentOnCloud             *bool             `pulumi:"isPresentOnCloud"`
	LastBackupStatus             *string           `pulumi:"lastBackupStatus"`
	LastRefreshedAt              *string           `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint          *string           `pulumi:"oldestRecoveryPoint"`
	OnPremiseLatestRecoveryPoint *string           `pulumi:"onPremiseLatestRecoveryPoint"`
	OnPremiseOldestRecoveryPoint *string           `pulumi:"onPremiseOldestRecoveryPoint"`
	OnPremiseRecoveryPointCount  *int              `pulumi:"onPremiseRecoveryPointCount"`
	ProtectableObjectLoadPath    map[string]string `pulumi:"protectableObjectLoadPath"`
	Protected                    *bool             `pulumi:"protected"`
	ProtectionGroupName          *string           `pulumi:"protectionGroupName"`
	RecoveryPointCount           *int              `pulumi:"recoveryPointCount"`
	TotalDiskStorageSizeInBytes  *string           `pulumi:"totalDiskStorageSizeInBytes"`
}

type DPMProtectedItemResponse struct {
	BackupEngineName                 *string                               `pulumi:"backupEngineName"`
	BackupManagementType             *string                               `pulumi:"backupManagementType"`
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *DPMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                               `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
	WorkloadType                     *string                               `pulumi:"workloadType"`
}

type DailyRetentionFormat struct {
	DaysOfTheMonth []Day `pulumi:"daysOfTheMonth"`
}

type DailyRetentionFormatResponse struct {
	DaysOfTheMonth []DayResponse `pulumi:"daysOfTheMonth"`
}

type DailyRetentionSchedule struct {
	RetentionDuration *RetentionDuration `pulumi:"retentionDuration"`
	RetentionTimes    []string           `pulumi:"retentionTimes"`
}

type DailyRetentionScheduleResponse struct {
	RetentionDuration *RetentionDurationResponse `pulumi:"retentionDuration"`
	RetentionTimes    []string                   `pulumi:"retentionTimes"`
}

type Day struct {
	Date   *int  `pulumi:"date"`
	IsLast *bool `pulumi:"isLast"`
}

type DayResponse struct {
	Date   *int  `pulumi:"date"`
	IsLast *bool `pulumi:"isLast"`
}

type DiskExclusionProperties struct {
	DiskLunList     []int `pulumi:"diskLunList"`
	IsInclusionList *bool `pulumi:"isInclusionList"`
}

type DiskExclusionPropertiesResponse struct {
	DiskLunList     []int `pulumi:"diskLunList"`
	IsInclusionList *bool `pulumi:"isInclusionList"`
}

type DistributedNodesInfo struct {
	NodeName *string `pulumi:"nodeName"`
	Status   *string `pulumi:"status"`
}

type DistributedNodesInfoResponse struct {
	ErrorDetail *ErrorDetailResponse `pulumi:"errorDetail"`
	NodeName    *string              `pulumi:"nodeName"`
	Status      *string              `pulumi:"status"`
}

type DpmContainer struct {
	BackupManagementType *string                   `pulumi:"backupManagementType"`
	CanReRegister        *bool                     `pulumi:"canReRegister"`
	ContainerId          *string                   `pulumi:"containerId"`
	ContainerType        string                    `pulumi:"containerType"`
	DpmAgentVersion      *string                   `pulumi:"dpmAgentVersion"`
	DpmServers           []string                  `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName         *string                   `pulumi:"friendlyName"`
	HealthStatus         *string                   `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                  `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                   `pulumi:"protectionStatus"`
	RegistrationStatus   *string                   `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                     `pulumi:"upgradeAvailable"`
}

type DpmContainerResponse struct {
	BackupManagementType *string                           `pulumi:"backupManagementType"`
	CanReRegister        *bool                             `pulumi:"canReRegister"`
	ContainerId          *string                           `pulumi:"containerId"`
	ContainerType        string                            `pulumi:"containerType"`
	DpmAgentVersion      *string                           `pulumi:"dpmAgentVersion"`
	DpmServers           []string                          `pulumi:"dpmServers"`
	ExtendedInfo         *DPMContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName         *string                           `pulumi:"friendlyName"`
	HealthStatus         *string                           `pulumi:"healthStatus"`
	ProtectedItemCount   *float64                          `pulumi:"protectedItemCount"`
	ProtectionStatus     *string                           `pulumi:"protectionStatus"`
	RegistrationStatus   *string                           `pulumi:"registrationStatus"`
	UpgradeAvailable     *bool                             `pulumi:"upgradeAvailable"`
}

type ErrorDetailResponse struct {
	Code            string   `pulumi:"code"`
	Message         string   `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
}

type ExtendedProperties struct {
	DiskExclusionProperties *DiskExclusionProperties `pulumi:"diskExclusionProperties"`
}

type ExtendedPropertiesResponse struct {
	DiskExclusionProperties *DiskExclusionPropertiesResponse `pulumi:"diskExclusionProperties"`
}

type GenericContainer struct {
	BackupManagementType *string                       `pulumi:"backupManagementType"`
	ContainerType        string                        `pulumi:"containerType"`
	ExtendedInformation  *GenericContainerExtendedInfo `pulumi:"extendedInformation"`
	FabricName           *string                       `pulumi:"fabricName"`
	FriendlyName         *string                       `pulumi:"friendlyName"`
	HealthStatus         *string                       `pulumi:"healthStatus"`
	RegistrationStatus   *string                       `pulumi:"registrationStatus"`
}

type GenericContainerExtendedInfo struct {
	ContainerIdentityInfo *ContainerIdentityInfo `pulumi:"containerIdentityInfo"`
	RawCertData           *string                `pulumi:"rawCertData"`
	ServiceEndpoints      map[string]string      `pulumi:"serviceEndpoints"`
}

type GenericContainerExtendedInfoResponse struct {
	ContainerIdentityInfo *ContainerIdentityInfoResponse `pulumi:"containerIdentityInfo"`
	RawCertData           *string                        `pulumi:"rawCertData"`
	ServiceEndpoints      map[string]string              `pulumi:"serviceEndpoints"`
}

type GenericContainerResponse struct {
	BackupManagementType *string                               `pulumi:"backupManagementType"`
	ContainerType        string                                `pulumi:"containerType"`
	ExtendedInformation  *GenericContainerExtendedInfoResponse `pulumi:"extendedInformation"`
	FabricName           *string                               `pulumi:"fabricName"`
	FriendlyName         *string                               `pulumi:"friendlyName"`
	HealthStatus         *string                               `pulumi:"healthStatus"`
	RegistrationStatus   *string                               `pulumi:"registrationStatus"`
}

type GenericProtectedItem struct {
	BackupManagementType             *string           `pulumi:"backupManagementType"`
	BackupSetName                    *string           `pulumi:"backupSetName"`
	ContainerName                    *string           `pulumi:"containerName"`
	CreateMode                       *string           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string           `pulumi:"deferredDeleteTimeRemaining"`
	FabricName                       *string           `pulumi:"fabricName"`
	FriendlyName                     *string           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool             `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string           `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string           `pulumi:"policyId"`
	PolicyState                      *string           `pulumi:"policyState"`
	ProtectedItemId                  *float64          `pulumi:"protectedItemId"`
	ProtectedItemType                string            `pulumi:"protectedItemType"`
	ProtectionState                  *string           `pulumi:"protectionState"`
	SourceAssociations               map[string]string `pulumi:"sourceAssociations"`
	SourceResourceId                 *string           `pulumi:"sourceResourceId"`
	WorkloadType                     *string           `pulumi:"workloadType"`
}

type GenericProtectedItemResponse struct {
	BackupManagementType             *string           `pulumi:"backupManagementType"`
	BackupSetName                    *string           `pulumi:"backupSetName"`
	ContainerName                    *string           `pulumi:"containerName"`
	CreateMode                       *string           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string           `pulumi:"deferredDeleteTimeRemaining"`
	FabricName                       *string           `pulumi:"fabricName"`
	FriendlyName                     *string           `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool             `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string           `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string           `pulumi:"policyId"`
	PolicyState                      *string           `pulumi:"policyState"`
	ProtectedItemId                  *float64          `pulumi:"protectedItemId"`
	ProtectedItemType                string            `pulumi:"protectedItemType"`
	ProtectionState                  *string           `pulumi:"protectionState"`
	SourceAssociations               map[string]string `pulumi:"sourceAssociations"`
	SourceResourceId                 *string           `pulumi:"sourceResourceId"`
	WorkloadType                     *string           `pulumi:"workloadType"`
}

type GenericProtectionPolicy struct {
	BackupManagementType string                `pulumi:"backupManagementType"`
	FabricName           *string               `pulumi:"fabricName"`
	ProtectedItemsCount  *int                  `pulumi:"protectedItemsCount"`
	SubProtectionPolicy  []SubProtectionPolicy `pulumi:"subProtectionPolicy"`
	TimeZone             *string               `pulumi:"timeZone"`
}

type GenericProtectionPolicyResponse struct {
	BackupManagementType string                        `pulumi:"backupManagementType"`
	FabricName           *string                       `pulumi:"fabricName"`
	ProtectedItemsCount  *int                          `pulumi:"protectedItemsCount"`
	SubProtectionPolicy  []SubProtectionPolicyResponse `pulumi:"subProtectionPolicy"`
	TimeZone             *string                       `pulumi:"timeZone"`
}

type IaaSVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type IaaSVMContainerResponse struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type InquiryInfo struct {
	InquiryDetails []WorkloadInquiryDetails `pulumi:"inquiryDetails"`
	Status         *string                  `pulumi:"status"`
}

type InquiryInfoResponse struct {
	ErrorDetail    *ErrorDetailResponse             `pulumi:"errorDetail"`
	InquiryDetails []WorkloadInquiryDetailsResponse `pulumi:"inquiryDetails"`
	Status         *string                          `pulumi:"status"`
}

type InquiryValidation struct {
	Status *string `pulumi:"status"`
}

type InquiryValidationResponse struct {
	AdditionalDetail string               `pulumi:"additionalDetail"`
	ErrorDetail      *ErrorDetailResponse `pulumi:"errorDetail"`
	Status           *string              `pulumi:"status"`
}

type InstantRPAdditionalDetails struct {
	AzureBackupRGNamePrefix *string `pulumi:"azureBackupRGNamePrefix"`
	AzureBackupRGNameSuffix *string `pulumi:"azureBackupRGNameSuffix"`
}

type InstantRPAdditionalDetailsResponse struct {
	AzureBackupRGNamePrefix *string `pulumi:"azureBackupRGNamePrefix"`
	AzureBackupRGNameSuffix *string `pulumi:"azureBackupRGNameSuffix"`
}

type KPIResourceHealthDetails struct {
	ResourceHealthStatus *string `pulumi:"resourceHealthStatus"`
}

type KPIResourceHealthDetailsResponse struct {
	ResourceHealthDetails []ResourceHealthDetailsResponse `pulumi:"resourceHealthDetails"`
	ResourceHealthStatus  *string                         `pulumi:"resourceHealthStatus"`
}

type LogSchedulePolicy struct {
	ScheduleFrequencyInMins *int   `pulumi:"scheduleFrequencyInMins"`
	SchedulePolicyType      string `pulumi:"schedulePolicyType"`
}

type LogSchedulePolicyResponse struct {
	ScheduleFrequencyInMins *int   `pulumi:"scheduleFrequencyInMins"`
	SchedulePolicyType      string `pulumi:"schedulePolicyType"`
}

type LongTermRetentionPolicy struct {
	DailySchedule       *DailyRetentionSchedule   `pulumi:"dailySchedule"`
	MonthlySchedule     *MonthlyRetentionSchedule `pulumi:"monthlySchedule"`
	RetentionPolicyType string                    `pulumi:"retentionPolicyType"`
	WeeklySchedule      *WeeklyRetentionSchedule  `pulumi:"weeklySchedule"`
	YearlySchedule      *YearlyRetentionSchedule  `pulumi:"yearlySchedule"`
}

type LongTermRetentionPolicyResponse struct {
	DailySchedule       *DailyRetentionScheduleResponse   `pulumi:"dailySchedule"`
	MonthlySchedule     *MonthlyRetentionScheduleResponse `pulumi:"monthlySchedule"`
	RetentionPolicyType string                            `pulumi:"retentionPolicyType"`
	WeeklySchedule      *WeeklyRetentionScheduleResponse  `pulumi:"weeklySchedule"`
	YearlySchedule      *YearlyRetentionScheduleResponse  `pulumi:"yearlySchedule"`
}

type LongTermSchedulePolicy struct {
	SchedulePolicyType string `pulumi:"schedulePolicyType"`
}

type LongTermSchedulePolicyResponse struct {
	SchedulePolicyType string `pulumi:"schedulePolicyType"`
}

type MABContainerHealthDetails struct {
	Code            *int     `pulumi:"code"`
	Message         *string  `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           *string  `pulumi:"title"`
}

type MABContainerHealthDetailsResponse struct {
	Code            *int     `pulumi:"code"`
	Message         *string  `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           *string  `pulumi:"title"`
}

type MabContainer struct {
	AgentVersion              *string                     `pulumi:"agentVersion"`
	BackupManagementType      *string                     `pulumi:"backupManagementType"`
	CanReRegister             *bool                       `pulumi:"canReRegister"`
	ContainerHealthState      *string                     `pulumi:"containerHealthState"`
	ContainerId               *float64                    `pulumi:"containerId"`
	ContainerType             string                      `pulumi:"containerType"`
	ExtendedInfo              *MabContainerExtendedInfo   `pulumi:"extendedInfo"`
	FriendlyName              *string                     `pulumi:"friendlyName"`
	HealthStatus              *string                     `pulumi:"healthStatus"`
	MabContainerHealthDetails []MABContainerHealthDetails `pulumi:"mabContainerHealthDetails"`
	ProtectedItemCount        *float64                    `pulumi:"protectedItemCount"`
	RegistrationStatus        *string                     `pulumi:"registrationStatus"`
}

type MabContainerExtendedInfo struct {
	BackupItemType   *string  `pulumi:"backupItemType"`
	BackupItems      []string `pulumi:"backupItems"`
	LastBackupStatus *string  `pulumi:"lastBackupStatus"`
	LastRefreshedAt  *string  `pulumi:"lastRefreshedAt"`
	PolicyName       *string  `pulumi:"policyName"`
}

type MabContainerExtendedInfoResponse struct {
	BackupItemType   *string  `pulumi:"backupItemType"`
	BackupItems      []string `pulumi:"backupItems"`
	LastBackupStatus *string  `pulumi:"lastBackupStatus"`
	LastRefreshedAt  *string  `pulumi:"lastRefreshedAt"`
	PolicyName       *string  `pulumi:"policyName"`
}

type MabContainerResponse struct {
	AgentVersion              *string                             `pulumi:"agentVersion"`
	BackupManagementType      *string                             `pulumi:"backupManagementType"`
	CanReRegister             *bool                               `pulumi:"canReRegister"`
	ContainerHealthState      *string                             `pulumi:"containerHealthState"`
	ContainerId               *float64                            `pulumi:"containerId"`
	ContainerType             string                              `pulumi:"containerType"`
	ExtendedInfo              *MabContainerExtendedInfoResponse   `pulumi:"extendedInfo"`
	FriendlyName              *string                             `pulumi:"friendlyName"`
	HealthStatus              *string                             `pulumi:"healthStatus"`
	MabContainerHealthDetails []MABContainerHealthDetailsResponse `pulumi:"mabContainerHealthDetails"`
	ProtectedItemCount        *float64                            `pulumi:"protectedItemCount"`
	RegistrationStatus        *string                             `pulumi:"registrationStatus"`
}

type MabFileFolderProtectedItem struct {
	BackupManagementType             *string                                 `pulumi:"backupManagementType"`
	BackupSetName                    *string                                 `pulumi:"backupSetName"`
	ComputerName                     *string                                 `pulumi:"computerName"`
	ContainerName                    *string                                 `pulumi:"containerName"`
	CreateMode                       *string                                 `pulumi:"createMode"`
	DeferredDeleteSyncTimeInUTC      *float64                                `pulumi:"deferredDeleteSyncTimeInUTC"`
	DeferredDeleteTimeInUTC          *string                                 `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                 `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *MabFileFolderProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                 `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                   `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                   `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                   `pulumi:"isScheduledForDeferredDelete"`
	LastBackupStatus                 *string                                 `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                 `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                 `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                 `pulumi:"policyId"`
	ProtectedItemType                string                                  `pulumi:"protectedItemType"`
	ProtectionState                  *string                                 `pulumi:"protectionState"`
	SourceResourceId                 *string                                 `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                 `pulumi:"workloadType"`
}

type MabFileFolderProtectedItemExtendedInfo struct {
	LastRefreshedAt     *string `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}

type MabFileFolderProtectedItemExtendedInfoResponse struct {
	LastRefreshedAt     *string `pulumi:"lastRefreshedAt"`
	OldestRecoveryPoint *string `pulumi:"oldestRecoveryPoint"`
	RecoveryPointCount  *int    `pulumi:"recoveryPointCount"`
}

type MabFileFolderProtectedItemResponse struct {
	BackupManagementType             *string                                         `pulumi:"backupManagementType"`
	BackupSetName                    *string                                         `pulumi:"backupSetName"`
	ComputerName                     *string                                         `pulumi:"computerName"`
	ContainerName                    *string                                         `pulumi:"containerName"`
	CreateMode                       *string                                         `pulumi:"createMode"`
	DeferredDeleteSyncTimeInUTC      *float64                                        `pulumi:"deferredDeleteSyncTimeInUTC"`
	DeferredDeleteTimeInUTC          *string                                         `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                         `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *MabFileFolderProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                         `pulumi:"friendlyName"`
	IsDeferredDeleteScheduleUpcoming *bool                                           `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                           `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                           `pulumi:"isScheduledForDeferredDelete"`
	LastBackupStatus                 *string                                         `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                         `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                         `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                         `pulumi:"policyId"`
	ProtectedItemType                string                                          `pulumi:"protectedItemType"`
	ProtectionState                  *string                                         `pulumi:"protectionState"`
	SourceResourceId                 *string                                         `pulumi:"sourceResourceId"`
	WorkloadType                     *string                                         `pulumi:"workloadType"`
}

type MabProtectionPolicy struct {
	BackupManagementType string      `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy       interface{} `pulumi:"schedulePolicy"`
}

type MabProtectionPolicyResponse struct {
	BackupManagementType string      `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy       interface{} `pulumi:"schedulePolicy"`
}

type MonthlyRetentionSchedule struct {
	RetentionDuration           *RetentionDuration     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormat  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *string                `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormat `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string               `pulumi:"retentionTimes"`
}

type MonthlyRetentionScheduleResponse struct {
	RetentionDuration           *RetentionDurationResponse     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormatResponse  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *string                        `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormatResponse `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string                       `pulumi:"retentionTimes"`
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionType struct {
	PrivateEndpoint                   *PrivateEndpoint                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string                            `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	PrivateEndpoint                   PrivateEndpointPtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringPtrInput                     `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypePtrOutput() PrivateEndpointConnectionTypePtrOutput {
	return i.ToPrivateEndpointConnectionTypePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput).ToPrivateEndpointConnectionTypePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionTypePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypePtrOutput() PrivateEndpointConnectionTypePtrOutput
	ToPrivateEndpointConnectionTypePtrOutputWithContext(context.Context) PrivateEndpointConnectionTypePtrOutput
}

type privateEndpointConnectionTypePtrType PrivateEndpointConnectionTypeArgs

func PrivateEndpointConnectionTypePtr(v *PrivateEndpointConnectionTypeArgs) PrivateEndpointConnectionTypePtrInput {
	return (*privateEndpointConnectionTypePtrType)(v)
}

func (*privateEndpointConnectionTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionType)(nil)).Elem()
}

func (i *privateEndpointConnectionTypePtrType) ToPrivateEndpointConnectionTypePtrOutput() PrivateEndpointConnectionTypePtrOutput {
	return i.ToPrivateEndpointConnectionTypePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionTypePtrType) ToPrivateEndpointConnectionTypePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypePtrOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypePtrOutput() PrivateEndpointConnectionTypePtrOutput {
	return o.ToPrivateEndpointConnectionTypePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionType) *PrivateEndpointConnectionType {
		return &v
	}).(PrivateEndpointConnectionTypePtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *PrivateEndpoint { return v.PrivateEndpoint }).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionTypeOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionTypePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypePtrOutput) ToPrivateEndpointConnectionTypePtrOutput() PrivateEndpointConnectionTypePtrOutput {
	return o
}

func (o PrivateEndpointConnectionTypePtrOutput) ToPrivateEndpointConnectionTypePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypePtrOutput {
	return o
}

func (o PrivateEndpointConnectionTypePtrOutput) Elem() PrivateEndpointConnectionTypeOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionType) PrivateEndpointConnectionType {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionType
		return ret
	}).(PrivateEndpointConnectionTypeOutput)
}

func (o PrivateEndpointConnectionTypePtrOutput) PrivateEndpoint() PrivateEndpointPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionType) *PrivateEndpoint {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionTypePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionType) *PrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionTypePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionType) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string                                    `pulumi:"provisioningState"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponse struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionRequired pulumi.StringPtrInput `pulumi:"actionRequired"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	Status         pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionRequired *string `pulumi:"actionRequired"`
	Description    *string `pulumi:"description"`
	Status         *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ResourceGuardOperationDetailResponse struct {
	DefaultResourceRequest *string `pulumi:"defaultResourceRequest"`
	VaultCriticalOperation *string `pulumi:"vaultCriticalOperation"`
}

type ResourceGuardOperationDetailResponseOutput struct{ *pulumi.OutputState }

func (ResourceGuardOperationDetailResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuardOperationDetailResponse)(nil)).Elem()
}

func (o ResourceGuardOperationDetailResponseOutput) ToResourceGuardOperationDetailResponseOutput() ResourceGuardOperationDetailResponseOutput {
	return o
}

func (o ResourceGuardOperationDetailResponseOutput) ToResourceGuardOperationDetailResponseOutputWithContext(ctx context.Context) ResourceGuardOperationDetailResponseOutput {
	return o
}

func (o ResourceGuardOperationDetailResponseOutput) DefaultResourceRequest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardOperationDetailResponse) *string { return v.DefaultResourceRequest }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardOperationDetailResponseOutput) VaultCriticalOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardOperationDetailResponse) *string { return v.VaultCriticalOperation }).(pulumi.StringPtrOutput)
}

type ResourceGuardOperationDetailResponseArrayOutput struct{ *pulumi.OutputState }

func (ResourceGuardOperationDetailResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceGuardOperationDetailResponse)(nil)).Elem()
}

func (o ResourceGuardOperationDetailResponseArrayOutput) ToResourceGuardOperationDetailResponseArrayOutput() ResourceGuardOperationDetailResponseArrayOutput {
	return o
}

func (o ResourceGuardOperationDetailResponseArrayOutput) ToResourceGuardOperationDetailResponseArrayOutputWithContext(ctx context.Context) ResourceGuardOperationDetailResponseArrayOutput {
	return o
}

func (o ResourceGuardOperationDetailResponseArrayOutput) Index(i pulumi.IntInput) ResourceGuardOperationDetailResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceGuardOperationDetailResponse {
		return vs[0].([]ResourceGuardOperationDetailResponse)[vs[1].(int)]
	}).(ResourceGuardOperationDetailResponseOutput)
}

type ResourceGuardProxyBaseResponse struct {
	LastUpdatedTime               *string                                `pulumi:"lastUpdatedTime"`
	ResourceGuardOperationDetails []ResourceGuardOperationDetailResponse `pulumi:"resourceGuardOperationDetails"`
	ResourceGuardResourceId       *string                                `pulumi:"resourceGuardResourceId"`
}

type ResourceGuardProxyBaseResponseOutput struct{ *pulumi.OutputState }

func (ResourceGuardProxyBaseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuardProxyBaseResponse)(nil)).Elem()
}

func (o ResourceGuardProxyBaseResponseOutput) ToResourceGuardProxyBaseResponseOutput() ResourceGuardProxyBaseResponseOutput {
	return o
}

func (o ResourceGuardProxyBaseResponseOutput) ToResourceGuardProxyBaseResponseOutputWithContext(ctx context.Context) ResourceGuardProxyBaseResponseOutput {
	return o
}

func (o ResourceGuardProxyBaseResponseOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardProxyBaseResponse) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyBaseResponseOutput) ResourceGuardOperationDetails() ResourceGuardOperationDetailResponseArrayOutput {
	return o.ApplyT(func(v ResourceGuardProxyBaseResponse) []ResourceGuardOperationDetailResponse {
		return v.ResourceGuardOperationDetails
	}).(ResourceGuardOperationDetailResponseArrayOutput)
}

func (o ResourceGuardProxyBaseResponseOutput) ResourceGuardResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardProxyBaseResponse) *string { return v.ResourceGuardResourceId }).(pulumi.StringPtrOutput)
}

type ResourceHealthDetailsResponse struct {
	Code            int      `pulumi:"code"`
	Message         string   `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           string   `pulumi:"title"`
}

type RetentionDuration struct {
	Count        *int    `pulumi:"count"`
	DurationType *string `pulumi:"durationType"`
}

type RetentionDurationResponse struct {
	Count        *int    `pulumi:"count"`
	DurationType *string `pulumi:"durationType"`
}

type Settings struct {
	IsCompression    *bool   `pulumi:"isCompression"`
	Issqlcompression *bool   `pulumi:"issqlcompression"`
	TimeZone         *string `pulumi:"timeZone"`
}

type SettingsResponse struct {
	IsCompression    *bool   `pulumi:"isCompression"`
	Issqlcompression *bool   `pulumi:"issqlcompression"`
	TimeZone         *string `pulumi:"timeZone"`
}

type SimpleRetentionPolicy struct {
	RetentionDuration   *RetentionDuration `pulumi:"retentionDuration"`
	RetentionPolicyType string             `pulumi:"retentionPolicyType"`
}

type SimpleRetentionPolicyResponse struct {
	RetentionDuration   *RetentionDurationResponse `pulumi:"retentionDuration"`
	RetentionPolicyType string                     `pulumi:"retentionPolicyType"`
}

type SimpleSchedulePolicy struct {
	SchedulePolicyType      string      `pulumi:"schedulePolicyType"`
	ScheduleRunDays         []DayOfWeek `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    *string     `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        []string    `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency *int        `pulumi:"scheduleWeeklyFrequency"`
}

type SimpleSchedulePolicyResponse struct {
	SchedulePolicyType      string   `pulumi:"schedulePolicyType"`
	ScheduleRunDays         []string `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    *string  `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        []string `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency *int     `pulumi:"scheduleWeeklyFrequency"`
}

type SubProtectionPolicy struct {
	PolicyType      *string     `pulumi:"policyType"`
	RetentionPolicy interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy  interface{} `pulumi:"schedulePolicy"`
}

type SubProtectionPolicyResponse struct {
	PolicyType      *string     `pulumi:"policyType"`
	RetentionPolicy interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy  interface{} `pulumi:"schedulePolicy"`
}

type WeeklyRetentionFormat struct {
	DaysOfTheWeek   []DayOfWeek   `pulumi:"daysOfTheWeek"`
	WeeksOfTheMonth []WeekOfMonth `pulumi:"weeksOfTheMonth"`
}

type WeeklyRetentionFormatResponse struct {
	DaysOfTheWeek   []string `pulumi:"daysOfTheWeek"`
	WeeksOfTheMonth []string `pulumi:"weeksOfTheMonth"`
}

type WeeklyRetentionSchedule struct {
	DaysOfTheWeek     []DayOfWeek        `pulumi:"daysOfTheWeek"`
	RetentionDuration *RetentionDuration `pulumi:"retentionDuration"`
	RetentionTimes    []string           `pulumi:"retentionTimes"`
}

type WeeklyRetentionScheduleResponse struct {
	DaysOfTheWeek     []string                   `pulumi:"daysOfTheWeek"`
	RetentionDuration *RetentionDurationResponse `pulumi:"retentionDuration"`
	RetentionTimes    []string                   `pulumi:"retentionTimes"`
}

type WorkloadInquiryDetails struct {
	InquiryValidation *InquiryValidation `pulumi:"inquiryValidation"`
	ItemCount         *float64           `pulumi:"itemCount"`
	Type              *string            `pulumi:"type"`
}

type WorkloadInquiryDetailsResponse struct {
	InquiryValidation *InquiryValidationResponse `pulumi:"inquiryValidation"`
	ItemCount         *float64                   `pulumi:"itemCount"`
	Type              *string                    `pulumi:"type"`
}

type YearlyRetentionSchedule struct {
	MonthsOfYear                []MonthOfYear          `pulumi:"monthsOfYear"`
	RetentionDuration           *RetentionDuration     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormat  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *string                `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormat `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string               `pulumi:"retentionTimes"`
}

type YearlyRetentionScheduleResponse struct {
	MonthsOfYear                []string                       `pulumi:"monthsOfYear"`
	RetentionDuration           *RetentionDurationResponse     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormatResponse  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *string                        `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormatResponse `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string                       `pulumi:"retentionTimes"`
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceGuardOperationDetailResponseOutput{})
	pulumi.RegisterOutputType(ResourceGuardOperationDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceGuardProxyBaseResponseOutput{})
}
