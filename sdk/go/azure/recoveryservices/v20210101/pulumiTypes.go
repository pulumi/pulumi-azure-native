


package v20210101

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

type CmkKekIdentity struct {
	UseSystemAssignedIdentity *bool   `pulumi:"useSystemAssignedIdentity"`
	UserAssignedIdentity      *string `pulumi:"userAssignedIdentity"`
}





type CmkKekIdentityInput interface {
	pulumi.Input

	ToCmkKekIdentityOutput() CmkKekIdentityOutput
	ToCmkKekIdentityOutputWithContext(context.Context) CmkKekIdentityOutput
}

type CmkKekIdentityArgs struct {
	UseSystemAssignedIdentity pulumi.BoolPtrInput   `pulumi:"useSystemAssignedIdentity"`
	UserAssignedIdentity      pulumi.StringPtrInput `pulumi:"userAssignedIdentity"`
}

func (CmkKekIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKekIdentity)(nil)).Elem()
}

func (i CmkKekIdentityArgs) ToCmkKekIdentityOutput() CmkKekIdentityOutput {
	return i.ToCmkKekIdentityOutputWithContext(context.Background())
}

func (i CmkKekIdentityArgs) ToCmkKekIdentityOutputWithContext(ctx context.Context) CmkKekIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKekIdentityOutput)
}

func (i CmkKekIdentityArgs) ToCmkKekIdentityPtrOutput() CmkKekIdentityPtrOutput {
	return i.ToCmkKekIdentityPtrOutputWithContext(context.Background())
}

func (i CmkKekIdentityArgs) ToCmkKekIdentityPtrOutputWithContext(ctx context.Context) CmkKekIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKekIdentityOutput).ToCmkKekIdentityPtrOutputWithContext(ctx)
}









type CmkKekIdentityPtrInput interface {
	pulumi.Input

	ToCmkKekIdentityPtrOutput() CmkKekIdentityPtrOutput
	ToCmkKekIdentityPtrOutputWithContext(context.Context) CmkKekIdentityPtrOutput
}

type cmkKekIdentityPtrType CmkKekIdentityArgs

func CmkKekIdentityPtr(v *CmkKekIdentityArgs) CmkKekIdentityPtrInput {
	return (*cmkKekIdentityPtrType)(v)
}

func (*cmkKekIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKekIdentity)(nil)).Elem()
}

func (i *cmkKekIdentityPtrType) ToCmkKekIdentityPtrOutput() CmkKekIdentityPtrOutput {
	return i.ToCmkKekIdentityPtrOutputWithContext(context.Background())
}

func (i *cmkKekIdentityPtrType) ToCmkKekIdentityPtrOutputWithContext(ctx context.Context) CmkKekIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKekIdentityPtrOutput)
}

type CmkKekIdentityOutput struct{ *pulumi.OutputState }

func (CmkKekIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKekIdentity)(nil)).Elem()
}

func (o CmkKekIdentityOutput) ToCmkKekIdentityOutput() CmkKekIdentityOutput {
	return o
}

func (o CmkKekIdentityOutput) ToCmkKekIdentityOutputWithContext(ctx context.Context) CmkKekIdentityOutput {
	return o
}

func (o CmkKekIdentityOutput) ToCmkKekIdentityPtrOutput() CmkKekIdentityPtrOutput {
	return o.ToCmkKekIdentityPtrOutputWithContext(context.Background())
}

func (o CmkKekIdentityOutput) ToCmkKekIdentityPtrOutputWithContext(ctx context.Context) CmkKekIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CmkKekIdentity) *CmkKekIdentity {
		return &v
	}).(CmkKekIdentityPtrOutput)
}

func (o CmkKekIdentityOutput) UseSystemAssignedIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CmkKekIdentity) *bool { return v.UseSystemAssignedIdentity }).(pulumi.BoolPtrOutput)
}

func (o CmkKekIdentityOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CmkKekIdentity) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type CmkKekIdentityPtrOutput struct{ *pulumi.OutputState }

func (CmkKekIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKekIdentity)(nil)).Elem()
}

func (o CmkKekIdentityPtrOutput) ToCmkKekIdentityPtrOutput() CmkKekIdentityPtrOutput {
	return o
}

func (o CmkKekIdentityPtrOutput) ToCmkKekIdentityPtrOutputWithContext(ctx context.Context) CmkKekIdentityPtrOutput {
	return o
}

func (o CmkKekIdentityPtrOutput) Elem() CmkKekIdentityOutput {
	return o.ApplyT(func(v *CmkKekIdentity) CmkKekIdentity {
		if v != nil {
			return *v
		}
		var ret CmkKekIdentity
		return ret
	}).(CmkKekIdentityOutput)
}

func (o CmkKekIdentityPtrOutput) UseSystemAssignedIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CmkKekIdentity) *bool {
		if v == nil {
			return nil
		}
		return v.UseSystemAssignedIdentity
	}).(pulumi.BoolPtrOutput)
}

func (o CmkKekIdentityPtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CmkKekIdentity) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type CmkKekIdentityResponse struct {
	UseSystemAssignedIdentity *bool   `pulumi:"useSystemAssignedIdentity"`
	UserAssignedIdentity      *string `pulumi:"userAssignedIdentity"`
}

type CmkKekIdentityResponseOutput struct{ *pulumi.OutputState }

func (CmkKekIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKekIdentityResponse)(nil)).Elem()
}

func (o CmkKekIdentityResponseOutput) ToCmkKekIdentityResponseOutput() CmkKekIdentityResponseOutput {
	return o
}

func (o CmkKekIdentityResponseOutput) ToCmkKekIdentityResponseOutputWithContext(ctx context.Context) CmkKekIdentityResponseOutput {
	return o
}

func (o CmkKekIdentityResponseOutput) UseSystemAssignedIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CmkKekIdentityResponse) *bool { return v.UseSystemAssignedIdentity }).(pulumi.BoolPtrOutput)
}

func (o CmkKekIdentityResponseOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CmkKekIdentityResponse) *string { return v.UserAssignedIdentity }).(pulumi.StringPtrOutput)
}

type CmkKekIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (CmkKekIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKekIdentityResponse)(nil)).Elem()
}

func (o CmkKekIdentityResponsePtrOutput) ToCmkKekIdentityResponsePtrOutput() CmkKekIdentityResponsePtrOutput {
	return o
}

func (o CmkKekIdentityResponsePtrOutput) ToCmkKekIdentityResponsePtrOutputWithContext(ctx context.Context) CmkKekIdentityResponsePtrOutput {
	return o
}

func (o CmkKekIdentityResponsePtrOutput) Elem() CmkKekIdentityResponseOutput {
	return o.ApplyT(func(v *CmkKekIdentityResponse) CmkKekIdentityResponse {
		if v != nil {
			return *v
		}
		var ret CmkKekIdentityResponse
		return ret
	}).(CmkKekIdentityResponseOutput)
}

func (o CmkKekIdentityResponsePtrOutput) UseSystemAssignedIdentity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CmkKekIdentityResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseSystemAssignedIdentity
	}).(pulumi.BoolPtrOutput)
}

func (o CmkKekIdentityResponsePtrOutput) UserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CmkKekIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.StringPtrOutput)
}

type CmkKeyVaultProperties struct {
	KeyUri *string `pulumi:"keyUri"`
}





type CmkKeyVaultPropertiesInput interface {
	pulumi.Input

	ToCmkKeyVaultPropertiesOutput() CmkKeyVaultPropertiesOutput
	ToCmkKeyVaultPropertiesOutputWithContext(context.Context) CmkKeyVaultPropertiesOutput
}

type CmkKeyVaultPropertiesArgs struct {
	KeyUri pulumi.StringPtrInput `pulumi:"keyUri"`
}

func (CmkKeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKeyVaultProperties)(nil)).Elem()
}

func (i CmkKeyVaultPropertiesArgs) ToCmkKeyVaultPropertiesOutput() CmkKeyVaultPropertiesOutput {
	return i.ToCmkKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i CmkKeyVaultPropertiesArgs) ToCmkKeyVaultPropertiesOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKeyVaultPropertiesOutput)
}

func (i CmkKeyVaultPropertiesArgs) ToCmkKeyVaultPropertiesPtrOutput() CmkKeyVaultPropertiesPtrOutput {
	return i.ToCmkKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i CmkKeyVaultPropertiesArgs) ToCmkKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKeyVaultPropertiesOutput).ToCmkKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type CmkKeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToCmkKeyVaultPropertiesPtrOutput() CmkKeyVaultPropertiesPtrOutput
	ToCmkKeyVaultPropertiesPtrOutputWithContext(context.Context) CmkKeyVaultPropertiesPtrOutput
}

type cmkKeyVaultPropertiesPtrType CmkKeyVaultPropertiesArgs

func CmkKeyVaultPropertiesPtr(v *CmkKeyVaultPropertiesArgs) CmkKeyVaultPropertiesPtrInput {
	return (*cmkKeyVaultPropertiesPtrType)(v)
}

func (*cmkKeyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKeyVaultProperties)(nil)).Elem()
}

func (i *cmkKeyVaultPropertiesPtrType) ToCmkKeyVaultPropertiesPtrOutput() CmkKeyVaultPropertiesPtrOutput {
	return i.ToCmkKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *cmkKeyVaultPropertiesPtrType) ToCmkKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CmkKeyVaultPropertiesPtrOutput)
}

type CmkKeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (CmkKeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKeyVaultProperties)(nil)).Elem()
}

func (o CmkKeyVaultPropertiesOutput) ToCmkKeyVaultPropertiesOutput() CmkKeyVaultPropertiesOutput {
	return o
}

func (o CmkKeyVaultPropertiesOutput) ToCmkKeyVaultPropertiesOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesOutput {
	return o
}

func (o CmkKeyVaultPropertiesOutput) ToCmkKeyVaultPropertiesPtrOutput() CmkKeyVaultPropertiesPtrOutput {
	return o.ToCmkKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o CmkKeyVaultPropertiesOutput) ToCmkKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CmkKeyVaultProperties) *CmkKeyVaultProperties {
		return &v
	}).(CmkKeyVaultPropertiesPtrOutput)
}

func (o CmkKeyVaultPropertiesOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CmkKeyVaultProperties) *string { return v.KeyUri }).(pulumi.StringPtrOutput)
}

type CmkKeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CmkKeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKeyVaultProperties)(nil)).Elem()
}

func (o CmkKeyVaultPropertiesPtrOutput) ToCmkKeyVaultPropertiesPtrOutput() CmkKeyVaultPropertiesPtrOutput {
	return o
}

func (o CmkKeyVaultPropertiesPtrOutput) ToCmkKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesPtrOutput {
	return o
}

func (o CmkKeyVaultPropertiesPtrOutput) Elem() CmkKeyVaultPropertiesOutput {
	return o.ApplyT(func(v *CmkKeyVaultProperties) CmkKeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret CmkKeyVaultProperties
		return ret
	}).(CmkKeyVaultPropertiesOutput)
}

func (o CmkKeyVaultPropertiesPtrOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CmkKeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyUri
	}).(pulumi.StringPtrOutput)
}

type CmkKeyVaultPropertiesResponse struct {
	KeyUri *string `pulumi:"keyUri"`
}

type CmkKeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CmkKeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CmkKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o CmkKeyVaultPropertiesResponseOutput) ToCmkKeyVaultPropertiesResponseOutput() CmkKeyVaultPropertiesResponseOutput {
	return o
}

func (o CmkKeyVaultPropertiesResponseOutput) ToCmkKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesResponseOutput {
	return o
}

func (o CmkKeyVaultPropertiesResponseOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CmkKeyVaultPropertiesResponse) *string { return v.KeyUri }).(pulumi.StringPtrOutput)
}

type CmkKeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CmkKeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CmkKeyVaultPropertiesResponse)(nil)).Elem()
}

func (o CmkKeyVaultPropertiesResponsePtrOutput) ToCmkKeyVaultPropertiesResponsePtrOutput() CmkKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o CmkKeyVaultPropertiesResponsePtrOutput) ToCmkKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) CmkKeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o CmkKeyVaultPropertiesResponsePtrOutput) Elem() CmkKeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *CmkKeyVaultPropertiesResponse) CmkKeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CmkKeyVaultPropertiesResponse
		return ret
	}).(CmkKeyVaultPropertiesResponseOutput)
}

func (o CmkKeyVaultPropertiesResponsePtrOutput) KeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CmkKeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyUri
	}).(pulumi.StringPtrOutput)
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

type IdentityData struct {
	Type                   string                 `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityDataInput interface {
	pulumi.Input

	ToIdentityDataOutput() IdentityDataOutput
	ToIdentityDataOutputWithContext(context.Context) IdentityDataOutput
}

type IdentityDataArgs struct {
	Type                   pulumi.StringInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput    `pulumi:"userAssignedIdentities"`
}

func (IdentityDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityData)(nil)).Elem()
}

func (i IdentityDataArgs) ToIdentityDataOutput() IdentityDataOutput {
	return i.ToIdentityDataOutputWithContext(context.Background())
}

func (i IdentityDataArgs) ToIdentityDataOutputWithContext(ctx context.Context) IdentityDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataOutput)
}

func (i IdentityDataArgs) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return i.ToIdentityDataPtrOutputWithContext(context.Background())
}

func (i IdentityDataArgs) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataOutput).ToIdentityDataPtrOutputWithContext(ctx)
}









type IdentityDataPtrInput interface {
	pulumi.Input

	ToIdentityDataPtrOutput() IdentityDataPtrOutput
	ToIdentityDataPtrOutputWithContext(context.Context) IdentityDataPtrOutput
}

type identityDataPtrType IdentityDataArgs

func IdentityDataPtr(v *IdentityDataArgs) IdentityDataPtrInput {
	return (*identityDataPtrType)(v)
}

func (*identityDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityData)(nil)).Elem()
}

func (i *identityDataPtrType) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return i.ToIdentityDataPtrOutputWithContext(context.Background())
}

func (i *identityDataPtrType) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataPtrOutput)
}

type IdentityDataOutput struct{ *pulumi.OutputState }

func (IdentityDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityData)(nil)).Elem()
}

func (o IdentityDataOutput) ToIdentityDataOutput() IdentityDataOutput {
	return o
}

func (o IdentityDataOutput) ToIdentityDataOutputWithContext(ctx context.Context) IdentityDataOutput {
	return o
}

func (o IdentityDataOutput) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return o.ToIdentityDataPtrOutputWithContext(context.Background())
}

func (o IdentityDataOutput) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityData) *IdentityData {
		return &v
	}).(IdentityDataPtrOutput)
}

func (o IdentityDataOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityData) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentityDataOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v IdentityData) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type IdentityDataPtrOutput struct{ *pulumi.OutputState }

func (IdentityDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityData)(nil)).Elem()
}

func (o IdentityDataPtrOutput) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return o
}

func (o IdentityDataPtrOutput) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return o
}

func (o IdentityDataPtrOutput) Elem() IdentityDataOutput {
	return o.ApplyT(func(v *IdentityData) IdentityData {
		if v != nil {
			return *v
		}
		var ret IdentityData
		return ret
	}).(IdentityDataOutput)
}

func (o IdentityDataPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityData) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDataPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *IdentityData) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityDataResponse struct {
	PrincipalId            string                          `pulumi:"principalId"`
	TenantId               string                          `pulumi:"tenantId"`
	Type                   string                          `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityResponse `pulumi:"userAssignedIdentities"`
}

type IdentityDataResponseOutput struct{ *pulumi.OutputState }

func (IdentityDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityDataResponse)(nil)).Elem()
}

func (o IdentityDataResponseOutput) ToIdentityDataResponseOutput() IdentityDataResponseOutput {
	return o
}

func (o IdentityDataResponseOutput) ToIdentityDataResponseOutputWithContext(ctx context.Context) IdentityDataResponseOutput {
	return o
}

func (o IdentityDataResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDataResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityDataResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDataResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityDataResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDataResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentityDataResponseOutput) UserAssignedIdentities() UserIdentityResponseMapOutput {
	return o.ApplyT(func(v IdentityDataResponse) map[string]UserIdentityResponse { return v.UserAssignedIdentities }).(UserIdentityResponseMapOutput)
}

type IdentityDataResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityDataResponse)(nil)).Elem()
}

func (o IdentityDataResponsePtrOutput) ToIdentityDataResponsePtrOutput() IdentityDataResponsePtrOutput {
	return o
}

func (o IdentityDataResponsePtrOutput) ToIdentityDataResponsePtrOutputWithContext(ctx context.Context) IdentityDataResponsePtrOutput {
	return o
}

func (o IdentityDataResponsePtrOutput) Elem() IdentityDataResponseOutput {
	return o.ApplyT(func(v *IdentityDataResponse) IdentityDataResponse {
		if v != nil {
			return *v
		}
		var ret IdentityDataResponse
		return ret
	}).(IdentityDataResponseOutput)
}

func (o IdentityDataResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDataResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDataResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDataResponsePtrOutput) UserAssignedIdentities() UserIdentityResponseMapOutput {
	return o.ApplyT(func(v *IdentityDataResponse) map[string]UserIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityResponseMapOutput)
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

type PrivateEndpointConnectionVaultPropertiesResponse struct {
	Id         string                                 `pulumi:"id"`
	Properties VaultPrivateEndpointConnectionResponse `pulumi:"properties"`
}

type PrivateEndpointConnectionVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionVaultPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) ToPrivateEndpointConnectionVaultPropertiesResponseOutput() PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) ToPrivateEndpointConnectionVaultPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Properties() VaultPrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) VaultPrivateEndpointConnectionResponse {
		return v.Properties
	}).(VaultPrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointConnectionVaultPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionVaultPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutput() PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionVaultPropertiesResponse {
		return vs[0].([]PrivateEndpointConnectionVaultPropertiesResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionVaultPropertiesResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
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

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
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
		return &v.Id
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

type Sku struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
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

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type UpgradeDetailsResponse struct {
	EndTimeUtc         string `pulumi:"endTimeUtc"`
	LastUpdatedTimeUtc string `pulumi:"lastUpdatedTimeUtc"`
	Message            string `pulumi:"message"`
	OperationId        string `pulumi:"operationId"`
	PreviousResourceId string `pulumi:"previousResourceId"`
	StartTimeUtc       string `pulumi:"startTimeUtc"`
	Status             string `pulumi:"status"`
	TriggerType        string `pulumi:"triggerType"`
	UpgradedResourceId string `pulumi:"upgradedResourceId"`
}

type UpgradeDetailsResponseOutput struct{ *pulumi.OutputState }

func (UpgradeDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradeDetailsResponse)(nil)).Elem()
}

func (o UpgradeDetailsResponseOutput) ToUpgradeDetailsResponseOutput() UpgradeDetailsResponseOutput {
	return o
}

func (o UpgradeDetailsResponseOutput) ToUpgradeDetailsResponseOutputWithContext(ctx context.Context) UpgradeDetailsResponseOutput {
	return o
}

func (o UpgradeDetailsResponseOutput) EndTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.EndTimeUtc }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) LastUpdatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.LastUpdatedTimeUtc }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.OperationId }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) PreviousResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.PreviousResourceId }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.TriggerType }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) UpgradedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.UpgradedResourceId }).(pulumi.StringOutput)
}

type UpgradeDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (UpgradeDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeDetailsResponse)(nil)).Elem()
}

func (o UpgradeDetailsResponsePtrOutput) ToUpgradeDetailsResponsePtrOutput() UpgradeDetailsResponsePtrOutput {
	return o
}

func (o UpgradeDetailsResponsePtrOutput) ToUpgradeDetailsResponsePtrOutputWithContext(ctx context.Context) UpgradeDetailsResponsePtrOutput {
	return o
}

func (o UpgradeDetailsResponsePtrOutput) Elem() UpgradeDetailsResponseOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) UpgradeDetailsResponse {
		if v != nil {
			return *v
		}
		var ret UpgradeDetailsResponse
		return ret
	}).(UpgradeDetailsResponseOutput)
}

func (o UpgradeDetailsResponsePtrOutput) EndTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) LastUpdatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) PreviousResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreviousResourceId
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TriggerType
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) UpgradedResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradedResourceId
	}).(pulumi.StringPtrOutput)
}

type UserIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutput() UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) ToUserIdentityResponseOutputWithContext(ctx context.Context) UserIdentityResponseOutput {
	return o
}

func (o UserIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityResponse)(nil)).Elem()
}

func (o UserIdentityResponseMapOutput) ToUserIdentityResponseMapOutput() UserIdentityResponseMapOutput {
	return o
}

func (o UserIdentityResponseMapOutput) ToUserIdentityResponseMapOutputWithContext(ctx context.Context) UserIdentityResponseMapOutput {
	return o
}

func (o UserIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityResponse {
		return vs[0].(map[string]UserIdentityResponse)[vs[1].(string)]
	}).(UserIdentityResponseOutput)
}

type VaultPrivateEndpointConnectionResponse struct {
	PrivateEndpoint                   PrivateEndpointResponse                        `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState VaultPrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                         `pulumi:"provisioningState"`
}

type VaultPrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (VaultPrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o VaultPrivateEndpointConnectionResponseOutput) ToVaultPrivateEndpointConnectionResponseOutput() VaultPrivateEndpointConnectionResponseOutput {
	return o
}

func (o VaultPrivateEndpointConnectionResponseOutput) ToVaultPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) VaultPrivateEndpointConnectionResponseOutput {
	return o
}

func (o VaultPrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponseOutput)
}

func (o VaultPrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() VaultPrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) VaultPrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(VaultPrivateLinkServiceConnectionStateResponseOutput)
}

func (o VaultPrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type VaultPrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}

type VaultPrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (VaultPrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) ToVaultPrivateLinkServiceConnectionStateResponseOutput() VaultPrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) ToVaultPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) VaultPrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateLinkServiceConnectionStateResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateLinkServiceConnectionStateResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateLinkServiceConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
}

type VaultProperties struct {
	Encryption *VaultPropertiesEncryption `pulumi:"encryption"`
}





type VaultPropertiesInput interface {
	pulumi.Input

	ToVaultPropertiesOutput() VaultPropertiesOutput
	ToVaultPropertiesOutputWithContext(context.Context) VaultPropertiesOutput
}

type VaultPropertiesArgs struct {
	Encryption VaultPropertiesEncryptionPtrInput `pulumi:"encryption"`
}

func (VaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultProperties)(nil)).Elem()
}

func (i VaultPropertiesArgs) ToVaultPropertiesOutput() VaultPropertiesOutput {
	return i.ToVaultPropertiesOutputWithContext(context.Background())
}

func (i VaultPropertiesArgs) ToVaultPropertiesOutputWithContext(ctx context.Context) VaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesOutput)
}

func (i VaultPropertiesArgs) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return i.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i VaultPropertiesArgs) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesOutput).ToVaultPropertiesPtrOutputWithContext(ctx)
}









type VaultPropertiesPtrInput interface {
	pulumi.Input

	ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput
	ToVaultPropertiesPtrOutputWithContext(context.Context) VaultPropertiesPtrOutput
}

type vaultPropertiesPtrType VaultPropertiesArgs

func VaultPropertiesPtr(v *VaultPropertiesArgs) VaultPropertiesPtrInput {
	return (*vaultPropertiesPtrType)(v)
}

func (*vaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultProperties)(nil)).Elem()
}

func (i *vaultPropertiesPtrType) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return i.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *vaultPropertiesPtrType) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesPtrOutput)
}

type VaultPropertiesOutput struct{ *pulumi.OutputState }

func (VaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultProperties)(nil)).Elem()
}

func (o VaultPropertiesOutput) ToVaultPropertiesOutput() VaultPropertiesOutput {
	return o
}

func (o VaultPropertiesOutput) ToVaultPropertiesOutputWithContext(ctx context.Context) VaultPropertiesOutput {
	return o
}

func (o VaultPropertiesOutput) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return o.ToVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o VaultPropertiesOutput) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VaultProperties) *VaultProperties {
		return &v
	}).(VaultPropertiesPtrOutput)
}

func (o VaultPropertiesOutput) Encryption() VaultPropertiesEncryptionPtrOutput {
	return o.ApplyT(func(v VaultProperties) *VaultPropertiesEncryption { return v.Encryption }).(VaultPropertiesEncryptionPtrOutput)
}

type VaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultProperties)(nil)).Elem()
}

func (o VaultPropertiesPtrOutput) ToVaultPropertiesPtrOutput() VaultPropertiesPtrOutput {
	return o
}

func (o VaultPropertiesPtrOutput) ToVaultPropertiesPtrOutputWithContext(ctx context.Context) VaultPropertiesPtrOutput {
	return o
}

func (o VaultPropertiesPtrOutput) Elem() VaultPropertiesOutput {
	return o.ApplyT(func(v *VaultProperties) VaultProperties {
		if v != nil {
			return *v
		}
		var ret VaultProperties
		return ret
	}).(VaultPropertiesOutput)
}

func (o VaultPropertiesPtrOutput) Encryption() VaultPropertiesEncryptionPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *VaultPropertiesEncryption {
		if v == nil {
			return nil
		}
		return v.Encryption
	}).(VaultPropertiesEncryptionPtrOutput)
}

type VaultPropertiesEncryption struct {
	InfrastructureEncryption *string                `pulumi:"infrastructureEncryption"`
	KekIdentity              *CmkKekIdentity        `pulumi:"kekIdentity"`
	KeyVaultProperties       *CmkKeyVaultProperties `pulumi:"keyVaultProperties"`
}





type VaultPropertiesEncryptionInput interface {
	pulumi.Input

	ToVaultPropertiesEncryptionOutput() VaultPropertiesEncryptionOutput
	ToVaultPropertiesEncryptionOutputWithContext(context.Context) VaultPropertiesEncryptionOutput
}

type VaultPropertiesEncryptionArgs struct {
	InfrastructureEncryption pulumi.StringPtrInput         `pulumi:"infrastructureEncryption"`
	KekIdentity              CmkKekIdentityPtrInput        `pulumi:"kekIdentity"`
	KeyVaultProperties       CmkKeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
}

func (VaultPropertiesEncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesEncryption)(nil)).Elem()
}

func (i VaultPropertiesEncryptionArgs) ToVaultPropertiesEncryptionOutput() VaultPropertiesEncryptionOutput {
	return i.ToVaultPropertiesEncryptionOutputWithContext(context.Background())
}

func (i VaultPropertiesEncryptionArgs) ToVaultPropertiesEncryptionOutputWithContext(ctx context.Context) VaultPropertiesEncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesEncryptionOutput)
}

func (i VaultPropertiesEncryptionArgs) ToVaultPropertiesEncryptionPtrOutput() VaultPropertiesEncryptionPtrOutput {
	return i.ToVaultPropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (i VaultPropertiesEncryptionArgs) ToVaultPropertiesEncryptionPtrOutputWithContext(ctx context.Context) VaultPropertiesEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesEncryptionOutput).ToVaultPropertiesEncryptionPtrOutputWithContext(ctx)
}









type VaultPropertiesEncryptionPtrInput interface {
	pulumi.Input

	ToVaultPropertiesEncryptionPtrOutput() VaultPropertiesEncryptionPtrOutput
	ToVaultPropertiesEncryptionPtrOutputWithContext(context.Context) VaultPropertiesEncryptionPtrOutput
}

type vaultPropertiesEncryptionPtrType VaultPropertiesEncryptionArgs

func VaultPropertiesEncryptionPtr(v *VaultPropertiesEncryptionArgs) VaultPropertiesEncryptionPtrInput {
	return (*vaultPropertiesEncryptionPtrType)(v)
}

func (*vaultPropertiesEncryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesEncryption)(nil)).Elem()
}

func (i *vaultPropertiesEncryptionPtrType) ToVaultPropertiesEncryptionPtrOutput() VaultPropertiesEncryptionPtrOutput {
	return i.ToVaultPropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (i *vaultPropertiesEncryptionPtrType) ToVaultPropertiesEncryptionPtrOutputWithContext(ctx context.Context) VaultPropertiesEncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesEncryptionPtrOutput)
}

type VaultPropertiesEncryptionOutput struct{ *pulumi.OutputState }

func (VaultPropertiesEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesEncryption)(nil)).Elem()
}

func (o VaultPropertiesEncryptionOutput) ToVaultPropertiesEncryptionOutput() VaultPropertiesEncryptionOutput {
	return o
}

func (o VaultPropertiesEncryptionOutput) ToVaultPropertiesEncryptionOutputWithContext(ctx context.Context) VaultPropertiesEncryptionOutput {
	return o
}

func (o VaultPropertiesEncryptionOutput) ToVaultPropertiesEncryptionPtrOutput() VaultPropertiesEncryptionPtrOutput {
	return o.ToVaultPropertiesEncryptionPtrOutputWithContext(context.Background())
}

func (o VaultPropertiesEncryptionOutput) ToVaultPropertiesEncryptionPtrOutputWithContext(ctx context.Context) VaultPropertiesEncryptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VaultPropertiesEncryption) *VaultPropertiesEncryption {
		return &v
	}).(VaultPropertiesEncryptionPtrOutput)
}

func (o VaultPropertiesEncryptionOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesEncryption) *string { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesEncryptionOutput) KekIdentity() CmkKekIdentityPtrOutput {
	return o.ApplyT(func(v VaultPropertiesEncryption) *CmkKekIdentity { return v.KekIdentity }).(CmkKekIdentityPtrOutput)
}

func (o VaultPropertiesEncryptionOutput) KeyVaultProperties() CmkKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v VaultPropertiesEncryption) *CmkKeyVaultProperties { return v.KeyVaultProperties }).(CmkKeyVaultPropertiesPtrOutput)
}

type VaultPropertiesEncryptionPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesEncryption)(nil)).Elem()
}

func (o VaultPropertiesEncryptionPtrOutput) ToVaultPropertiesEncryptionPtrOutput() VaultPropertiesEncryptionPtrOutput {
	return o
}

func (o VaultPropertiesEncryptionPtrOutput) ToVaultPropertiesEncryptionPtrOutputWithContext(ctx context.Context) VaultPropertiesEncryptionPtrOutput {
	return o
}

func (o VaultPropertiesEncryptionPtrOutput) Elem() VaultPropertiesEncryptionOutput {
	return o.ApplyT(func(v *VaultPropertiesEncryption) VaultPropertiesEncryption {
		if v != nil {
			return *v
		}
		var ret VaultPropertiesEncryption
		return ret
	}).(VaultPropertiesEncryptionOutput)
}

func (o VaultPropertiesEncryptionPtrOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesEncryption) *string {
		if v == nil {
			return nil
		}
		return v.InfrastructureEncryption
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesEncryptionPtrOutput) KekIdentity() CmkKekIdentityPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesEncryption) *CmkKekIdentity {
		if v == nil {
			return nil
		}
		return v.KekIdentity
	}).(CmkKekIdentityPtrOutput)
}

func (o VaultPropertiesEncryptionPtrOutput) KeyVaultProperties() CmkKeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesEncryption) *CmkKeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(CmkKeyVaultPropertiesPtrOutput)
}

type VaultPropertiesResponse struct {
	Encryption                          *VaultPropertiesResponseEncryption                 `pulumi:"encryption"`
	PrivateEndpointConnections          []PrivateEndpointConnectionVaultPropertiesResponse `pulumi:"privateEndpointConnections"`
	PrivateEndpointStateForBackup       string                                             `pulumi:"privateEndpointStateForBackup"`
	PrivateEndpointStateForSiteRecovery string                                             `pulumi:"privateEndpointStateForSiteRecovery"`
	ProvisioningState                   string                                             `pulumi:"provisioningState"`
	UpgradeDetails                      *UpgradeDetailsResponse                            `pulumi:"upgradeDetails"`
}

type VaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponse)(nil)).Elem()
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponseOutput() VaultPropertiesResponseOutput {
	return o
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponseOutputWithContext(ctx context.Context) VaultPropertiesResponseOutput {
	return o
}

func (o VaultPropertiesResponseOutput) Encryption() VaultPropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *VaultPropertiesResponseEncryption { return v.Encryption }).(VaultPropertiesResponseEncryptionPtrOutput)
}

func (o VaultPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) []PrivateEndpointConnectionVaultPropertiesResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionVaultPropertiesResponseArrayOutput)
}

func (o VaultPropertiesResponseOutput) PrivateEndpointStateForBackup() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.PrivateEndpointStateForBackup }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) PrivateEndpointStateForSiteRecovery() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.PrivateEndpointStateForSiteRecovery }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) UpgradeDetails() UpgradeDetailsResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *UpgradeDetailsResponse { return v.UpgradeDetails }).(UpgradeDetailsResponsePtrOutput)
}

type VaultPropertiesResponseEncryption struct {
	InfrastructureEncryption *string                        `pulumi:"infrastructureEncryption"`
	KekIdentity              *CmkKekIdentityResponse        `pulumi:"kekIdentity"`
	KeyVaultProperties       *CmkKeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
}

type VaultPropertiesResponseEncryptionOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseEncryptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponseEncryption)(nil)).Elem()
}

func (o VaultPropertiesResponseEncryptionOutput) ToVaultPropertiesResponseEncryptionOutput() VaultPropertiesResponseEncryptionOutput {
	return o
}

func (o VaultPropertiesResponseEncryptionOutput) ToVaultPropertiesResponseEncryptionOutputWithContext(ctx context.Context) VaultPropertiesResponseEncryptionOutput {
	return o
}

func (o VaultPropertiesResponseEncryptionOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponseEncryption) *string { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseEncryptionOutput) KekIdentity() CmkKekIdentityResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponseEncryption) *CmkKekIdentityResponse { return v.KekIdentity }).(CmkKekIdentityResponsePtrOutput)
}

func (o VaultPropertiesResponseEncryptionOutput) KeyVaultProperties() CmkKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponseEncryption) *CmkKeyVaultPropertiesResponse { return v.KeyVaultProperties }).(CmkKeyVaultPropertiesResponsePtrOutput)
}

type VaultPropertiesResponseEncryptionPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseEncryptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponseEncryption)(nil)).Elem()
}

func (o VaultPropertiesResponseEncryptionPtrOutput) ToVaultPropertiesResponseEncryptionPtrOutput() VaultPropertiesResponseEncryptionPtrOutput {
	return o
}

func (o VaultPropertiesResponseEncryptionPtrOutput) ToVaultPropertiesResponseEncryptionPtrOutputWithContext(ctx context.Context) VaultPropertiesResponseEncryptionPtrOutput {
	return o
}

func (o VaultPropertiesResponseEncryptionPtrOutput) Elem() VaultPropertiesResponseEncryptionOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseEncryption) VaultPropertiesResponseEncryption {
		if v != nil {
			return *v
		}
		var ret VaultPropertiesResponseEncryption
		return ret
	}).(VaultPropertiesResponseEncryptionOutput)
}

func (o VaultPropertiesResponseEncryptionPtrOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseEncryption) *string {
		if v == nil {
			return nil
		}
		return v.InfrastructureEncryption
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseEncryptionPtrOutput) KekIdentity() CmkKekIdentityResponsePtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseEncryption) *CmkKekIdentityResponse {
		if v == nil {
			return nil
		}
		return v.KekIdentity
	}).(CmkKekIdentityResponsePtrOutput)
}

func (o VaultPropertiesResponseEncryptionPtrOutput) KeyVaultProperties() CmkKeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseEncryption) *CmkKeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(CmkKeyVaultPropertiesResponsePtrOutput)
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
	pulumi.RegisterOutputType(CmkKekIdentityOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityPtrOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityResponseOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityDataOutput{})
	pulumi.RegisterOutputType(IdentityDataPtrOutput{})
	pulumi.RegisterOutputType(IdentityDataResponseOutput{})
	pulumi.RegisterOutputType(IdentityDataResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionVaultPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UpgradeDetailsResponseOutput{})
	pulumi.RegisterOutputType(UpgradeDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VaultPrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(VaultPrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesOutput{})
	pulumi.RegisterOutputType(VaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesEncryptionOutput{})
	pulumi.RegisterOutputType(VaultPropertiesEncryptionPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseEncryptionOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseEncryptionPtrOutput{})
}
