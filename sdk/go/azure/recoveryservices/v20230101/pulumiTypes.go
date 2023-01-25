


package v20230101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureBackupServerContainer struct {
	BackupManagementType  *string                   `pulumi:"backupManagementType"`
	CanReRegister         *bool                     `pulumi:"canReRegister"`
	ContainerId           *string                   `pulumi:"containerId"`
	ContainerType         string                    `pulumi:"containerType"`
	DpmAgentVersion       *string                   `pulumi:"dpmAgentVersion"`
	DpmServers            []string                  `pulumi:"dpmServers"`
	ExtendedInfo          *DPMContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName          *string                   `pulumi:"friendlyName"`
	HealthStatus          *string                   `pulumi:"healthStatus"`
	ProtectableObjectType *string                   `pulumi:"protectableObjectType"`
	ProtectedItemCount    *float64                  `pulumi:"protectedItemCount"`
	ProtectionStatus      *string                   `pulumi:"protectionStatus"`
	RegistrationStatus    *string                   `pulumi:"registrationStatus"`
	UpgradeAvailable      *bool                     `pulumi:"upgradeAvailable"`
}

type AzureBackupServerContainerResponse struct {
	BackupManagementType  *string                           `pulumi:"backupManagementType"`
	CanReRegister         *bool                             `pulumi:"canReRegister"`
	ContainerId           *string                           `pulumi:"containerId"`
	ContainerType         string                            `pulumi:"containerType"`
	DpmAgentVersion       *string                           `pulumi:"dpmAgentVersion"`
	DpmServers            []string                          `pulumi:"dpmServers"`
	ExtendedInfo          *DPMContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName          *string                           `pulumi:"friendlyName"`
	HealthStatus          *string                           `pulumi:"healthStatus"`
	ProtectableObjectType *string                           `pulumi:"protectableObjectType"`
	ProtectedItemCount    *float64                          `pulumi:"protectedItemCount"`
	ProtectionStatus      *string                           `pulumi:"protectionStatus"`
	RegistrationStatus    *string                           `pulumi:"registrationStatus"`
	UpgradeAvailable      *bool                             `pulumi:"upgradeAvailable"`
}

type AzureFileShareProtectionPolicy struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{} `pulumi:"schedulePolicy"`
	TimeZone                       *string     `pulumi:"timeZone"`
	WorkLoadType                   *string     `pulumi:"workLoadType"`
}

type AzureFileShareProtectionPolicyResponse struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{} `pulumi:"schedulePolicy"`
	TimeZone                       *string     `pulumi:"timeZone"`
	WorkLoadType                   *string     `pulumi:"workLoadType"`
}

type AzureFileshareProtectedItem struct {
	BackupSetName                    *string                                  `pulumi:"backupSetName"`
	ContainerName                    *string                                  `pulumi:"containerName"`
	CreateMode                       *string                                  `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                  `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                  `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureFileshareProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                  `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                                    `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                    `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                    `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                    `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails      `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                  `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                  `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                  `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                  `pulumi:"policyId"`
	PolicyName                       *string                                  `pulumi:"policyName"`
	ProtectedItemType                string                                   `pulumi:"protectedItemType"`
	ProtectionState                  *string                                  `pulumi:"protectionState"`
	ProtectionStatus                 *string                                  `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                 `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                     `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                  `pulumi:"sourceResourceId"`
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
	BackupManagementType             string                                           `pulumi:"backupManagementType"`
	BackupSetName                    *string                                          `pulumi:"backupSetName"`
	ContainerName                    *string                                          `pulumi:"containerName"`
	CreateMode                       *string                                          `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                          `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                          `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureFileshareProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                          `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                                            `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                            `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                            `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                            `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse      `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                          `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                          `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                          `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                          `pulumi:"policyId"`
	PolicyName                       *string                                          `pulumi:"policyName"`
	ProtectedItemType                string                                           `pulumi:"protectedItemType"`
	ProtectionState                  *string                                          `pulumi:"protectionState"`
	ProtectionStatus                 *string                                          `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                         `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                             `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                          `pulumi:"sourceResourceId"`
	WorkloadType                     string                                           `pulumi:"workloadType"`
}

type AzureIaaSClassicComputeVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	ProtectableObjectType *string `pulumi:"protectableObjectType"`
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
	ProtectableObjectType *string `pulumi:"protectableObjectType"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type AzureIaaSClassicComputeVMProtectedItem struct {
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedProperties                   `pulumi:"extendedProperties"`
	IsArchiveEnabled                 *bool                                 `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                               `pulumi:"lastBackupStatus"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	PolicyName                       *string                               `pulumi:"policyName"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ProtectionStatus                 *string                               `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                              `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                  `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
}

type AzureIaaSClassicComputeVMProtectedItemResponse struct {
	BackupManagementType             string                                        `pulumi:"backupManagementType"`
	BackupSetName                    *string                                       `pulumi:"backupSetName"`
	ContainerName                    *string                                       `pulumi:"containerName"`
	CreateMode                       *string                                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedPropertiesResponse                   `pulumi:"extendedProperties"`
	FriendlyName                     string                                        `pulumi:"friendlyName"`
	HealthDetails                    []AzureIaaSVMHealthDetailsResponse            `pulumi:"healthDetails"`
	HealthStatus                     string                                        `pulumi:"healthStatus"`
	IsArchiveEnabled                 *bool                                         `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                         `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                       `pulumi:"lastBackupStatus"`
	LastBackupTime                   string                                        `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                       `pulumi:"policyId"`
	PolicyName                       *string                                       `pulumi:"policyName"`
	ProtectedItemDataId              string                                        `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                                       `pulumi:"protectionState"`
	ProtectionStatus                 *string                                       `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                      `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                          `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                       `pulumi:"sourceResourceId"`
	VirtualMachineId                 string                                        `pulumi:"virtualMachineId"`
	WorkloadType                     string                                        `pulumi:"workloadType"`
}

type AzureIaaSComputeVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	ProtectableObjectType *string `pulumi:"protectableObjectType"`
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
	ProtectableObjectType *string `pulumi:"protectableObjectType"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
	ResourceGroup         *string `pulumi:"resourceGroup"`
	VirtualMachineId      *string `pulumi:"virtualMachineId"`
	VirtualMachineVersion *string `pulumi:"virtualMachineVersion"`
}

type AzureIaaSComputeVMProtectedItem struct {
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedProperties                   `pulumi:"extendedProperties"`
	IsArchiveEnabled                 *bool                                 `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                               `pulumi:"lastBackupStatus"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	PolicyName                       *string                               `pulumi:"policyName"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ProtectionStatus                 *string                               `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                              `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                  `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
}

type AzureIaaSComputeVMProtectedItemResponse struct {
	BackupManagementType             string                                        `pulumi:"backupManagementType"`
	BackupSetName                    *string                                       `pulumi:"backupSetName"`
	ContainerName                    *string                                       `pulumi:"containerName"`
	CreateMode                       *string                                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedPropertiesResponse                   `pulumi:"extendedProperties"`
	FriendlyName                     string                                        `pulumi:"friendlyName"`
	HealthDetails                    []AzureIaaSVMHealthDetailsResponse            `pulumi:"healthDetails"`
	HealthStatus                     string                                        `pulumi:"healthStatus"`
	IsArchiveEnabled                 *bool                                         `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                         `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                       `pulumi:"lastBackupStatus"`
	LastBackupTime                   string                                        `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                       `pulumi:"policyId"`
	PolicyName                       *string                                       `pulumi:"policyName"`
	ProtectedItemDataId              string                                        `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                                       `pulumi:"protectionState"`
	ProtectionStatus                 *string                                       `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                      `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                          `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                       `pulumi:"sourceResourceId"`
	VirtualMachineId                 string                                        `pulumi:"virtualMachineId"`
	WorkloadType                     string                                        `pulumi:"workloadType"`
}

type AzureIaaSVMHealthDetailsResponse struct {
	Code            int      `pulumi:"code"`
	Message         string   `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
	Title           string   `pulumi:"title"`
}

type AzureIaaSVMProtectedItem struct {
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedProperties                   `pulumi:"extendedProperties"`
	IsArchiveEnabled                 *bool                                 `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetails   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                               `pulumi:"lastBackupStatus"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	PolicyName                       *string                               `pulumi:"policyName"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ProtectionStatus                 *string                               `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                              `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                  `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
}

type AzureIaaSVMProtectedItemExtendedInfo struct {
	NewestRecoveryPointInArchive *string `pulumi:"newestRecoveryPointInArchive"`
	OldestRecoveryPoint          *string `pulumi:"oldestRecoveryPoint"`
	OldestRecoveryPointInArchive *string `pulumi:"oldestRecoveryPointInArchive"`
	OldestRecoveryPointInVault   *string `pulumi:"oldestRecoveryPointInVault"`
	PolicyInconsistent           *bool   `pulumi:"policyInconsistent"`
	RecoveryPointCount           *int    `pulumi:"recoveryPointCount"`
}

type AzureIaaSVMProtectedItemExtendedInfoResponse struct {
	NewestRecoveryPointInArchive *string `pulumi:"newestRecoveryPointInArchive"`
	OldestRecoveryPoint          *string `pulumi:"oldestRecoveryPoint"`
	OldestRecoveryPointInArchive *string `pulumi:"oldestRecoveryPointInArchive"`
	OldestRecoveryPointInVault   *string `pulumi:"oldestRecoveryPointInVault"`
	PolicyInconsistent           *bool   `pulumi:"policyInconsistent"`
	RecoveryPointCount           *int    `pulumi:"recoveryPointCount"`
}

type AzureIaaSVMProtectedItemResponse struct {
	BackupManagementType             string                                        `pulumi:"backupManagementType"`
	BackupSetName                    *string                                       `pulumi:"backupSetName"`
	ContainerName                    *string                                       `pulumi:"containerName"`
	CreateMode                       *string                                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureIaaSVMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	ExtendedProperties               *ExtendedPropertiesResponse                   `pulumi:"extendedProperties"`
	FriendlyName                     string                                        `pulumi:"friendlyName"`
	HealthDetails                    []AzureIaaSVMHealthDetailsResponse            `pulumi:"healthDetails"`
	HealthStatus                     string                                        `pulumi:"healthStatus"`
	IsArchiveEnabled                 *bool                                         `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                         `pulumi:"isScheduledForDeferredDelete"`
	KpisHealths                      map[string]KPIResourceHealthDetailsResponse   `pulumi:"kpisHealths"`
	LastBackupStatus                 *string                                       `pulumi:"lastBackupStatus"`
	LastBackupTime                   string                                        `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                       `pulumi:"policyId"`
	PolicyName                       *string                                       `pulumi:"policyName"`
	ProtectedItemDataId              string                                        `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                                       `pulumi:"protectionState"`
	ProtectionStatus                 *string                                       `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                      `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                          `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                       `pulumi:"sourceResourceId"`
	VirtualMachineId                 string                                        `pulumi:"virtualMachineId"`
	WorkloadType                     string                                        `pulumi:"workloadType"`
}

type AzureIaaSVMProtectionPolicy struct {
	BackupManagementType           string                      `pulumi:"backupManagementType"`
	InstantRPDetails               *InstantRPAdditionalDetails `pulumi:"instantRPDetails"`
	InstantRpRetentionRangeInDays  *int                        `pulumi:"instantRpRetentionRangeInDays"`
	PolicyType                     *string                     `pulumi:"policyType"`
	ProtectedItemsCount            *int                        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string                    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{}                 `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{}                 `pulumi:"schedulePolicy"`
	TieringPolicy                  map[string]TieringPolicy    `pulumi:"tieringPolicy"`
	TimeZone                       *string                     `pulumi:"timeZone"`
}

type AzureIaaSVMProtectionPolicyResponse struct {
	BackupManagementType           string                              `pulumi:"backupManagementType"`
	InstantRPDetails               *InstantRPAdditionalDetailsResponse `pulumi:"instantRPDetails"`
	InstantRpRetentionRangeInDays  *int                                `pulumi:"instantRpRetentionRangeInDays"`
	PolicyType                     *string                             `pulumi:"policyType"`
	ProtectedItemsCount            *int                                `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string                            `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{}                         `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{}                         `pulumi:"schedulePolicy"`
	TieringPolicy                  map[string]TieringPolicyResponse    `pulumi:"tieringPolicy"`
	TimeZone                       *string                             `pulumi:"timeZone"`
}

type AzureMonitorAlertSettings struct {
	AlertsForAllJobFailures *string `pulumi:"alertsForAllJobFailures"`
}





type AzureMonitorAlertSettingsInput interface {
	pulumi.Input

	ToAzureMonitorAlertSettingsOutput() AzureMonitorAlertSettingsOutput
	ToAzureMonitorAlertSettingsOutputWithContext(context.Context) AzureMonitorAlertSettingsOutput
}

type AzureMonitorAlertSettingsArgs struct {
	AlertsForAllJobFailures pulumi.StringPtrInput `pulumi:"alertsForAllJobFailures"`
}

func (AzureMonitorAlertSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorAlertSettings)(nil)).Elem()
}

func (i AzureMonitorAlertSettingsArgs) ToAzureMonitorAlertSettingsOutput() AzureMonitorAlertSettingsOutput {
	return i.ToAzureMonitorAlertSettingsOutputWithContext(context.Background())
}

func (i AzureMonitorAlertSettingsArgs) ToAzureMonitorAlertSettingsOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorAlertSettingsOutput)
}

func (i AzureMonitorAlertSettingsArgs) ToAzureMonitorAlertSettingsPtrOutput() AzureMonitorAlertSettingsPtrOutput {
	return i.ToAzureMonitorAlertSettingsPtrOutputWithContext(context.Background())
}

func (i AzureMonitorAlertSettingsArgs) ToAzureMonitorAlertSettingsPtrOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorAlertSettingsOutput).ToAzureMonitorAlertSettingsPtrOutputWithContext(ctx)
}









type AzureMonitorAlertSettingsPtrInput interface {
	pulumi.Input

	ToAzureMonitorAlertSettingsPtrOutput() AzureMonitorAlertSettingsPtrOutput
	ToAzureMonitorAlertSettingsPtrOutputWithContext(context.Context) AzureMonitorAlertSettingsPtrOutput
}

type azureMonitorAlertSettingsPtrType AzureMonitorAlertSettingsArgs

func AzureMonitorAlertSettingsPtr(v *AzureMonitorAlertSettingsArgs) AzureMonitorAlertSettingsPtrInput {
	return (*azureMonitorAlertSettingsPtrType)(v)
}

func (*azureMonitorAlertSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorAlertSettings)(nil)).Elem()
}

func (i *azureMonitorAlertSettingsPtrType) ToAzureMonitorAlertSettingsPtrOutput() AzureMonitorAlertSettingsPtrOutput {
	return i.ToAzureMonitorAlertSettingsPtrOutputWithContext(context.Background())
}

func (i *azureMonitorAlertSettingsPtrType) ToAzureMonitorAlertSettingsPtrOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureMonitorAlertSettingsPtrOutput)
}

type AzureMonitorAlertSettingsOutput struct{ *pulumi.OutputState }

func (AzureMonitorAlertSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorAlertSettings)(nil)).Elem()
}

func (o AzureMonitorAlertSettingsOutput) ToAzureMonitorAlertSettingsOutput() AzureMonitorAlertSettingsOutput {
	return o
}

func (o AzureMonitorAlertSettingsOutput) ToAzureMonitorAlertSettingsOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsOutput {
	return o
}

func (o AzureMonitorAlertSettingsOutput) ToAzureMonitorAlertSettingsPtrOutput() AzureMonitorAlertSettingsPtrOutput {
	return o.ToAzureMonitorAlertSettingsPtrOutputWithContext(context.Background())
}

func (o AzureMonitorAlertSettingsOutput) ToAzureMonitorAlertSettingsPtrOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureMonitorAlertSettings) *AzureMonitorAlertSettings {
		return &v
	}).(AzureMonitorAlertSettingsPtrOutput)
}

func (o AzureMonitorAlertSettingsOutput) AlertsForAllJobFailures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorAlertSettings) *string { return v.AlertsForAllJobFailures }).(pulumi.StringPtrOutput)
}

type AzureMonitorAlertSettingsPtrOutput struct{ *pulumi.OutputState }

func (AzureMonitorAlertSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorAlertSettings)(nil)).Elem()
}

func (o AzureMonitorAlertSettingsPtrOutput) ToAzureMonitorAlertSettingsPtrOutput() AzureMonitorAlertSettingsPtrOutput {
	return o
}

func (o AzureMonitorAlertSettingsPtrOutput) ToAzureMonitorAlertSettingsPtrOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsPtrOutput {
	return o
}

func (o AzureMonitorAlertSettingsPtrOutput) Elem() AzureMonitorAlertSettingsOutput {
	return o.ApplyT(func(v *AzureMonitorAlertSettings) AzureMonitorAlertSettings {
		if v != nil {
			return *v
		}
		var ret AzureMonitorAlertSettings
		return ret
	}).(AzureMonitorAlertSettingsOutput)
}

func (o AzureMonitorAlertSettingsPtrOutput) AlertsForAllJobFailures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMonitorAlertSettings) *string {
		if v == nil {
			return nil
		}
		return v.AlertsForAllJobFailures
	}).(pulumi.StringPtrOutput)
}

type AzureMonitorAlertSettingsResponse struct {
	AlertsForAllJobFailures *string `pulumi:"alertsForAllJobFailures"`
}

type AzureMonitorAlertSettingsResponseOutput struct{ *pulumi.OutputState }

func (AzureMonitorAlertSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureMonitorAlertSettingsResponse)(nil)).Elem()
}

func (o AzureMonitorAlertSettingsResponseOutput) ToAzureMonitorAlertSettingsResponseOutput() AzureMonitorAlertSettingsResponseOutput {
	return o
}

func (o AzureMonitorAlertSettingsResponseOutput) ToAzureMonitorAlertSettingsResponseOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsResponseOutput {
	return o
}

func (o AzureMonitorAlertSettingsResponseOutput) AlertsForAllJobFailures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureMonitorAlertSettingsResponse) *string { return v.AlertsForAllJobFailures }).(pulumi.StringPtrOutput)
}

type AzureMonitorAlertSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureMonitorAlertSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureMonitorAlertSettingsResponse)(nil)).Elem()
}

func (o AzureMonitorAlertSettingsResponsePtrOutput) ToAzureMonitorAlertSettingsResponsePtrOutput() AzureMonitorAlertSettingsResponsePtrOutput {
	return o
}

func (o AzureMonitorAlertSettingsResponsePtrOutput) ToAzureMonitorAlertSettingsResponsePtrOutputWithContext(ctx context.Context) AzureMonitorAlertSettingsResponsePtrOutput {
	return o
}

func (o AzureMonitorAlertSettingsResponsePtrOutput) Elem() AzureMonitorAlertSettingsResponseOutput {
	return o.ApplyT(func(v *AzureMonitorAlertSettingsResponse) AzureMonitorAlertSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AzureMonitorAlertSettingsResponse
		return ret
	}).(AzureMonitorAlertSettingsResponseOutput)
}

func (o AzureMonitorAlertSettingsResponsePtrOutput) AlertsForAllJobFailures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureMonitorAlertSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlertsForAllJobFailures
	}).(pulumi.StringPtrOutput)
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
	BackupManagementType  *string                             `pulumi:"backupManagementType"`
	ContainerType         string                              `pulumi:"containerType"`
	ExtendedInfo          *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName          *string                             `pulumi:"friendlyName"`
	HealthStatus          *string                             `pulumi:"healthStatus"`
	LastUpdatedTime       *string                             `pulumi:"lastUpdatedTime"`
	OperationType         *string                             `pulumi:"operationType"`
	ProtectableObjectType *string                             `pulumi:"protectableObjectType"`
	RegistrationStatus    *string                             `pulumi:"registrationStatus"`
	SourceResourceId      *string                             `pulumi:"sourceResourceId"`
	WorkloadType          *string                             `pulumi:"workloadType"`
}

type AzureSQLAGWorkloadContainerProtectionContainerResponse struct {
	BackupManagementType  *string                                     `pulumi:"backupManagementType"`
	ContainerType         string                                      `pulumi:"containerType"`
	ExtendedInfo          *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName          *string                                     `pulumi:"friendlyName"`
	HealthStatus          *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime       *string                                     `pulumi:"lastUpdatedTime"`
	OperationType         *string                                     `pulumi:"operationType"`
	ProtectableObjectType *string                                     `pulumi:"protectableObjectType"`
	RegistrationStatus    *string                                     `pulumi:"registrationStatus"`
	SourceResourceId      *string                                     `pulumi:"sourceResourceId"`
	WorkloadType          *string                                     `pulumi:"workloadType"`
}

type AzureSqlContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	ProtectableObjectType *string `pulumi:"protectableObjectType"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
}

type AzureSqlContainerResponse struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	ProtectableObjectType *string `pulumi:"protectableObjectType"`
	RegistrationStatus    *string `pulumi:"registrationStatus"`
}

type AzureSqlProtectedItem struct {
	BackupSetName                    *string                            `pulumi:"backupSetName"`
	ContainerName                    *string                            `pulumi:"containerName"`
	CreateMode                       *string                            `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                            `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                            `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureSqlProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	IsArchiveEnabled                 *bool                              `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                              `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                              `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                              `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                            `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                            `pulumi:"policyId"`
	PolicyName                       *string                            `pulumi:"policyName"`
	ProtectedItemDataId              *string                            `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                             `pulumi:"protectedItemType"`
	ProtectionState                  *string                            `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                           `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                               `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                            `pulumi:"sourceResourceId"`
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
	BackupManagementType             string                                     `pulumi:"backupManagementType"`
	BackupSetName                    *string                                    `pulumi:"backupSetName"`
	ContainerName                    *string                                    `pulumi:"containerName"`
	CreateMode                       *string                                    `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                    `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                    `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureSqlProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	IsArchiveEnabled                 *bool                                      `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                      `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                      `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                      `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                                    `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                    `pulumi:"policyId"`
	PolicyName                       *string                                    `pulumi:"policyName"`
	ProtectedItemDataId              *string                                    `pulumi:"protectedItemDataId"`
	ProtectedItemType                string                                     `pulumi:"protectedItemType"`
	ProtectionState                  *string                                    `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                   `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                       `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                    `pulumi:"sourceResourceId"`
	WorkloadType                     string                                     `pulumi:"workloadType"`
}

type AzureSqlProtectionPolicy struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
}

type AzureSqlProtectionPolicyResponse struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
}

type AzureStorageContainer struct {
	AcquireStorageAccountLock *string  `pulumi:"acquireStorageAccountLock"`
	BackupManagementType      *string  `pulumi:"backupManagementType"`
	ContainerType             string   `pulumi:"containerType"`
	FriendlyName              *string  `pulumi:"friendlyName"`
	HealthStatus              *string  `pulumi:"healthStatus"`
	ProtectableObjectType     *string  `pulumi:"protectableObjectType"`
	ProtectedItemCount        *float64 `pulumi:"protectedItemCount"`
	RegistrationStatus        *string  `pulumi:"registrationStatus"`
	ResourceGroup             *string  `pulumi:"resourceGroup"`
	SourceResourceId          *string  `pulumi:"sourceResourceId"`
	StorageAccountVersion     *string  `pulumi:"storageAccountVersion"`
}

type AzureStorageContainerResponse struct {
	AcquireStorageAccountLock *string  `pulumi:"acquireStorageAccountLock"`
	BackupManagementType      *string  `pulumi:"backupManagementType"`
	ContainerType             string   `pulumi:"containerType"`
	FriendlyName              *string  `pulumi:"friendlyName"`
	HealthStatus              *string  `pulumi:"healthStatus"`
	ProtectableObjectType     *string  `pulumi:"protectableObjectType"`
	ProtectedItemCount        *float64 `pulumi:"protectedItemCount"`
	RegistrationStatus        *string  `pulumi:"registrationStatus"`
	ResourceGroup             *string  `pulumi:"resourceGroup"`
	SourceResourceId          *string  `pulumi:"sourceResourceId"`
	StorageAccountVersion     *string  `pulumi:"storageAccountVersion"`
}

type AzureVMAppContainerProtectionContainer struct {
	BackupManagementType  *string                             `pulumi:"backupManagementType"`
	ContainerType         string                              `pulumi:"containerType"`
	ExtendedInfo          *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName          *string                             `pulumi:"friendlyName"`
	HealthStatus          *string                             `pulumi:"healthStatus"`
	LastUpdatedTime       *string                             `pulumi:"lastUpdatedTime"`
	OperationType         *string                             `pulumi:"operationType"`
	ProtectableObjectType *string                             `pulumi:"protectableObjectType"`
	RegistrationStatus    *string                             `pulumi:"registrationStatus"`
	SourceResourceId      *string                             `pulumi:"sourceResourceId"`
	WorkloadType          *string                             `pulumi:"workloadType"`
}

type AzureVMAppContainerProtectionContainerResponse struct {
	BackupManagementType  *string                                     `pulumi:"backupManagementType"`
	ContainerType         string                                      `pulumi:"containerType"`
	ExtendedInfo          *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName          *string                                     `pulumi:"friendlyName"`
	HealthStatus          *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime       *string                                     `pulumi:"lastUpdatedTime"`
	OperationType         *string                                     `pulumi:"operationType"`
	ProtectableObjectType *string                                     `pulumi:"protectableObjectType"`
	RegistrationStatus    *string                                     `pulumi:"registrationStatus"`
	SourceResourceId      *string                                     `pulumi:"sourceResourceId"`
	WorkloadType          *string                                     `pulumi:"workloadType"`
}

type AzureVmWorkloadProtectedItem struct {
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	IsArchiveEnabled                 *bool                                     `pulumi:"isArchiveEnabled"`
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
	PolicyName                       *string                                   `pulumi:"policyName"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SoftDeleteRetentionPeriod        *int                                      `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
}

type AzureVmWorkloadProtectedItemExtendedInfo struct {
	NewestRecoveryPointInArchive *string `pulumi:"newestRecoveryPointInArchive"`
	OldestRecoveryPoint          *string `pulumi:"oldestRecoveryPoint"`
	OldestRecoveryPointInArchive *string `pulumi:"oldestRecoveryPointInArchive"`
	OldestRecoveryPointInVault   *string `pulumi:"oldestRecoveryPointInVault"`
	PolicyState                  *string `pulumi:"policyState"`
	RecoveryModel                *string `pulumi:"recoveryModel"`
	RecoveryPointCount           *int    `pulumi:"recoveryPointCount"`
}

type AzureVmWorkloadProtectedItemExtendedInfoResponse struct {
	NewestRecoveryPointInArchive *string `pulumi:"newestRecoveryPointInArchive"`
	OldestRecoveryPoint          *string `pulumi:"oldestRecoveryPoint"`
	OldestRecoveryPointInArchive *string `pulumi:"oldestRecoveryPointInArchive"`
	OldestRecoveryPointInVault   *string `pulumi:"oldestRecoveryPointInVault"`
	PolicyState                  *string `pulumi:"policyState"`
	RecoveryModel                *string `pulumi:"recoveryModel"`
	RecoveryPointCount           *int    `pulumi:"recoveryPointCount"`
}

type AzureVmWorkloadProtectedItemResponse struct {
	BackupManagementType             string                                            `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     string                                            `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                                             `pulumi:"isArchiveEnabled"`
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
	PolicyName                       *string                                           `pulumi:"policyName"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 string                                            `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SoftDeleteRetentionPeriod        *int                                              `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     string                                            `pulumi:"workloadType"`
}

type AzureVmWorkloadProtectionPolicy struct {
	BackupManagementType           string                `pulumi:"backupManagementType"`
	MakePolicyConsistent           *bool                 `pulumi:"makePolicyConsistent"`
	ProtectedItemsCount            *int                  `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string              `pulumi:"resourceGuardOperationRequests"`
	Settings                       *Settings             `pulumi:"settings"`
	SubProtectionPolicy            []SubProtectionPolicy `pulumi:"subProtectionPolicy"`
	WorkLoadType                   *string               `pulumi:"workLoadType"`
}

type AzureVmWorkloadProtectionPolicyResponse struct {
	BackupManagementType           string                        `pulumi:"backupManagementType"`
	MakePolicyConsistent           *bool                         `pulumi:"makePolicyConsistent"`
	ProtectedItemsCount            *int                          `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string                      `pulumi:"resourceGuardOperationRequests"`
	Settings                       *SettingsResponse             `pulumi:"settings"`
	SubProtectionPolicy            []SubProtectionPolicyResponse `pulumi:"subProtectionPolicy"`
	WorkLoadType                   *string                       `pulumi:"workLoadType"`
}

type AzureVmWorkloadSAPAseDatabaseProtectedItem struct {
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	IsArchiveEnabled                 *bool                                     `pulumi:"isArchiveEnabled"`
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
	PolicyName                       *string                                   `pulumi:"policyName"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SoftDeleteRetentionPeriod        *int                                      `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
}

type AzureVmWorkloadSAPAseDatabaseProtectedItemResponse struct {
	BackupManagementType             string                                            `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     string                                            `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                                             `pulumi:"isArchiveEnabled"`
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
	PolicyName                       *string                                           `pulumi:"policyName"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 string                                            `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SoftDeleteRetentionPeriod        *int                                              `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     string                                            `pulumi:"workloadType"`
}

type AzureVmWorkloadSAPHanaDBInstanceProtectedItem struct {
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	IsArchiveEnabled                 *bool                                     `pulumi:"isArchiveEnabled"`
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
	PolicyName                       *string                                   `pulumi:"policyName"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SoftDeleteRetentionPeriod        *int                                      `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
}

type AzureVmWorkloadSAPHanaDBInstanceProtectedItemResponse struct {
	BackupManagementType             string                                            `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     string                                            `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                                             `pulumi:"isArchiveEnabled"`
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
	PolicyName                       *string                                           `pulumi:"policyName"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 string                                            `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SoftDeleteRetentionPeriod        *int                                              `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     string                                            `pulumi:"workloadType"`
}

type AzureVmWorkloadSAPHanaDatabaseProtectedItem struct {
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	IsArchiveEnabled                 *bool                                     `pulumi:"isArchiveEnabled"`
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
	PolicyName                       *string                                   `pulumi:"policyName"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SoftDeleteRetentionPeriod        *int                                      `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
}

type AzureVmWorkloadSAPHanaDatabaseProtectedItemResponse struct {
	BackupManagementType             string                                            `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     string                                            `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                                             `pulumi:"isArchiveEnabled"`
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
	PolicyName                       *string                                           `pulumi:"policyName"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 string                                            `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SoftDeleteRetentionPeriod        *int                                              `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     string                                            `pulumi:"workloadType"`
}

type AzureVmWorkloadSQLDatabaseProtectedItem struct {
	BackupSetName                    *string                                   `pulumi:"backupSetName"`
	ContainerName                    *string                                   `pulumi:"containerName"`
	CreateMode                       *string                                   `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                   `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                   `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	IsArchiveEnabled                 *bool                                     `pulumi:"isArchiveEnabled"`
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
	PolicyName                       *string                                   `pulumi:"policyName"`
	ProtectedItemDataSourceId        *string                                   `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                   `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                    `pulumi:"protectedItemType"`
	ProtectionState                  *string                                   `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                  `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                   `pulumi:"serverName"`
	SoftDeleteRetentionPeriod        *int                                      `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                   `pulumi:"sourceResourceId"`
}

type AzureVmWorkloadSQLDatabaseProtectedItemResponse struct {
	BackupManagementType             string                                            `pulumi:"backupManagementType"`
	BackupSetName                    *string                                           `pulumi:"backupSetName"`
	ContainerName                    *string                                           `pulumi:"containerName"`
	CreateMode                       *string                                           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                                           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                           `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *AzureVmWorkloadProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     string                                            `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                                             `pulumi:"isArchiveEnabled"`
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
	PolicyName                       *string                                           `pulumi:"policyName"`
	ProtectedItemDataSourceId        *string                                           `pulumi:"protectedItemDataSourceId"`
	ProtectedItemHealthStatus        *string                                           `pulumi:"protectedItemHealthStatus"`
	ProtectedItemType                string                                            `pulumi:"protectedItemType"`
	ProtectionState                  *string                                           `pulumi:"protectionState"`
	ProtectionStatus                 string                                            `pulumi:"protectionStatus"`
	ResourceGuardOperationRequests   []string                                          `pulumi:"resourceGuardOperationRequests"`
	ServerName                       *string                                           `pulumi:"serverName"`
	SoftDeleteRetentionPeriod        *int                                              `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                           `pulumi:"sourceResourceId"`
	WorkloadType                     string                                            `pulumi:"workloadType"`
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
	BackupManagementType  *string                             `pulumi:"backupManagementType"`
	ContainerType         string                              `pulumi:"containerType"`
	ExtendedInfo          *AzureWorkloadContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName          *string                             `pulumi:"friendlyName"`
	HealthStatus          *string                             `pulumi:"healthStatus"`
	LastUpdatedTime       *string                             `pulumi:"lastUpdatedTime"`
	OperationType         *string                             `pulumi:"operationType"`
	ProtectableObjectType *string                             `pulumi:"protectableObjectType"`
	RegistrationStatus    *string                             `pulumi:"registrationStatus"`
	SourceResourceId      *string                             `pulumi:"sourceResourceId"`
	WorkloadType          *string                             `pulumi:"workloadType"`
}

type AzureWorkloadContainerAutoProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}

type AzureWorkloadContainerAutoProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType string  `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
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
	BackupManagementType  *string                                     `pulumi:"backupManagementType"`
	ContainerType         string                                      `pulumi:"containerType"`
	ExtendedInfo          *AzureWorkloadContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName          *string                                     `pulumi:"friendlyName"`
	HealthStatus          *string                                     `pulumi:"healthStatus"`
	LastUpdatedTime       *string                                     `pulumi:"lastUpdatedTime"`
	OperationType         *string                                     `pulumi:"operationType"`
	ProtectableObjectType *string                                     `pulumi:"protectableObjectType"`
	RegistrationStatus    *string                                     `pulumi:"registrationStatus"`
	SourceResourceId      *string                                     `pulumi:"sourceResourceId"`
	WorkloadType          *string                                     `pulumi:"workloadType"`
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

type ClassicAlertSettings struct {
	AlertsForCriticalOperations *string `pulumi:"alertsForCriticalOperations"`
}





type ClassicAlertSettingsInput interface {
	pulumi.Input

	ToClassicAlertSettingsOutput() ClassicAlertSettingsOutput
	ToClassicAlertSettingsOutputWithContext(context.Context) ClassicAlertSettingsOutput
}

type ClassicAlertSettingsArgs struct {
	AlertsForCriticalOperations pulumi.StringPtrInput `pulumi:"alertsForCriticalOperations"`
}

func (ClassicAlertSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassicAlertSettings)(nil)).Elem()
}

func (i ClassicAlertSettingsArgs) ToClassicAlertSettingsOutput() ClassicAlertSettingsOutput {
	return i.ToClassicAlertSettingsOutputWithContext(context.Background())
}

func (i ClassicAlertSettingsArgs) ToClassicAlertSettingsOutputWithContext(ctx context.Context) ClassicAlertSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassicAlertSettingsOutput)
}

func (i ClassicAlertSettingsArgs) ToClassicAlertSettingsPtrOutput() ClassicAlertSettingsPtrOutput {
	return i.ToClassicAlertSettingsPtrOutputWithContext(context.Background())
}

func (i ClassicAlertSettingsArgs) ToClassicAlertSettingsPtrOutputWithContext(ctx context.Context) ClassicAlertSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassicAlertSettingsOutput).ToClassicAlertSettingsPtrOutputWithContext(ctx)
}









type ClassicAlertSettingsPtrInput interface {
	pulumi.Input

	ToClassicAlertSettingsPtrOutput() ClassicAlertSettingsPtrOutput
	ToClassicAlertSettingsPtrOutputWithContext(context.Context) ClassicAlertSettingsPtrOutput
}

type classicAlertSettingsPtrType ClassicAlertSettingsArgs

func ClassicAlertSettingsPtr(v *ClassicAlertSettingsArgs) ClassicAlertSettingsPtrInput {
	return (*classicAlertSettingsPtrType)(v)
}

func (*classicAlertSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClassicAlertSettings)(nil)).Elem()
}

func (i *classicAlertSettingsPtrType) ToClassicAlertSettingsPtrOutput() ClassicAlertSettingsPtrOutput {
	return i.ToClassicAlertSettingsPtrOutputWithContext(context.Background())
}

func (i *classicAlertSettingsPtrType) ToClassicAlertSettingsPtrOutputWithContext(ctx context.Context) ClassicAlertSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClassicAlertSettingsPtrOutput)
}

type ClassicAlertSettingsOutput struct{ *pulumi.OutputState }

func (ClassicAlertSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassicAlertSettings)(nil)).Elem()
}

func (o ClassicAlertSettingsOutput) ToClassicAlertSettingsOutput() ClassicAlertSettingsOutput {
	return o
}

func (o ClassicAlertSettingsOutput) ToClassicAlertSettingsOutputWithContext(ctx context.Context) ClassicAlertSettingsOutput {
	return o
}

func (o ClassicAlertSettingsOutput) ToClassicAlertSettingsPtrOutput() ClassicAlertSettingsPtrOutput {
	return o.ToClassicAlertSettingsPtrOutputWithContext(context.Background())
}

func (o ClassicAlertSettingsOutput) ToClassicAlertSettingsPtrOutputWithContext(ctx context.Context) ClassicAlertSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClassicAlertSettings) *ClassicAlertSettings {
		return &v
	}).(ClassicAlertSettingsPtrOutput)
}

func (o ClassicAlertSettingsOutput) AlertsForCriticalOperations() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClassicAlertSettings) *string { return v.AlertsForCriticalOperations }).(pulumi.StringPtrOutput)
}

type ClassicAlertSettingsPtrOutput struct{ *pulumi.OutputState }

func (ClassicAlertSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClassicAlertSettings)(nil)).Elem()
}

func (o ClassicAlertSettingsPtrOutput) ToClassicAlertSettingsPtrOutput() ClassicAlertSettingsPtrOutput {
	return o
}

func (o ClassicAlertSettingsPtrOutput) ToClassicAlertSettingsPtrOutputWithContext(ctx context.Context) ClassicAlertSettingsPtrOutput {
	return o
}

func (o ClassicAlertSettingsPtrOutput) Elem() ClassicAlertSettingsOutput {
	return o.ApplyT(func(v *ClassicAlertSettings) ClassicAlertSettings {
		if v != nil {
			return *v
		}
		var ret ClassicAlertSettings
		return ret
	}).(ClassicAlertSettingsOutput)
}

func (o ClassicAlertSettingsPtrOutput) AlertsForCriticalOperations() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClassicAlertSettings) *string {
		if v == nil {
			return nil
		}
		return v.AlertsForCriticalOperations
	}).(pulumi.StringPtrOutput)
}

type ClassicAlertSettingsResponse struct {
	AlertsForCriticalOperations *string `pulumi:"alertsForCriticalOperations"`
}

type ClassicAlertSettingsResponseOutput struct{ *pulumi.OutputState }

func (ClassicAlertSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClassicAlertSettingsResponse)(nil)).Elem()
}

func (o ClassicAlertSettingsResponseOutput) ToClassicAlertSettingsResponseOutput() ClassicAlertSettingsResponseOutput {
	return o
}

func (o ClassicAlertSettingsResponseOutput) ToClassicAlertSettingsResponseOutputWithContext(ctx context.Context) ClassicAlertSettingsResponseOutput {
	return o
}

func (o ClassicAlertSettingsResponseOutput) AlertsForCriticalOperations() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClassicAlertSettingsResponse) *string { return v.AlertsForCriticalOperations }).(pulumi.StringPtrOutput)
}

type ClassicAlertSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ClassicAlertSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClassicAlertSettingsResponse)(nil)).Elem()
}

func (o ClassicAlertSettingsResponsePtrOutput) ToClassicAlertSettingsResponsePtrOutput() ClassicAlertSettingsResponsePtrOutput {
	return o
}

func (o ClassicAlertSettingsResponsePtrOutput) ToClassicAlertSettingsResponsePtrOutputWithContext(ctx context.Context) ClassicAlertSettingsResponsePtrOutput {
	return o
}

func (o ClassicAlertSettingsResponsePtrOutput) Elem() ClassicAlertSettingsResponseOutput {
	return o.ApplyT(func(v *ClassicAlertSettingsResponse) ClassicAlertSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ClassicAlertSettingsResponse
		return ret
	}).(ClassicAlertSettingsResponseOutput)
}

func (o ClassicAlertSettingsResponsePtrOutput) AlertsForCriticalOperations() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClassicAlertSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlertsForCriticalOperations
	}).(pulumi.StringPtrOutput)
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
	BackupSetName                    *string                       `pulumi:"backupSetName"`
	ContainerName                    *string                       `pulumi:"containerName"`
	CreateMode                       *string                       `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                       `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                       `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *DPMProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                       `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                         `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                         `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                         `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                         `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                       `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                       `pulumi:"policyId"`
	PolicyName                       *string                       `pulumi:"policyName"`
	ProtectedItemType                string                        `pulumi:"protectedItemType"`
	ProtectionState                  *string                       `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                      `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                          `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                       `pulumi:"sourceResourceId"`
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
	BackupManagementType             string                                `pulumi:"backupManagementType"`
	BackupSetName                    *string                               `pulumi:"backupSetName"`
	ContainerName                    *string                               `pulumi:"containerName"`
	CreateMode                       *string                               `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string                               `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                               `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *DPMProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                               `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                                 `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                 `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                 `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                 `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string                               `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                               `pulumi:"policyId"`
	PolicyName                       *string                               `pulumi:"policyName"`
	ProtectedItemType                string                                `pulumi:"protectedItemType"`
	ProtectionState                  *string                               `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                              `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                  `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                               `pulumi:"sourceResourceId"`
	WorkloadType                     string                                `pulumi:"workloadType"`
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

type DailySchedule struct {
	ScheduleRunTimes []string `pulumi:"scheduleRunTimes"`
}

type DailyScheduleResponse struct {
	ScheduleRunTimes []string `pulumi:"scheduleRunTimes"`
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
	BackupManagementType  *string                   `pulumi:"backupManagementType"`
	CanReRegister         *bool                     `pulumi:"canReRegister"`
	ContainerId           *string                   `pulumi:"containerId"`
	ContainerType         string                    `pulumi:"containerType"`
	DpmAgentVersion       *string                   `pulumi:"dpmAgentVersion"`
	DpmServers            []string                  `pulumi:"dpmServers"`
	ExtendedInfo          *DPMContainerExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName          *string                   `pulumi:"friendlyName"`
	HealthStatus          *string                   `pulumi:"healthStatus"`
	ProtectableObjectType *string                   `pulumi:"protectableObjectType"`
	ProtectedItemCount    *float64                  `pulumi:"protectedItemCount"`
	ProtectionStatus      *string                   `pulumi:"protectionStatus"`
	RegistrationStatus    *string                   `pulumi:"registrationStatus"`
	UpgradeAvailable      *bool                     `pulumi:"upgradeAvailable"`
}

type DpmContainerResponse struct {
	BackupManagementType  *string                           `pulumi:"backupManagementType"`
	CanReRegister         *bool                             `pulumi:"canReRegister"`
	ContainerId           *string                           `pulumi:"containerId"`
	ContainerType         string                            `pulumi:"containerType"`
	DpmAgentVersion       *string                           `pulumi:"dpmAgentVersion"`
	DpmServers            []string                          `pulumi:"dpmServers"`
	ExtendedInfo          *DPMContainerExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName          *string                           `pulumi:"friendlyName"`
	HealthStatus          *string                           `pulumi:"healthStatus"`
	ProtectableObjectType *string                           `pulumi:"protectableObjectType"`
	ProtectedItemCount    *float64                          `pulumi:"protectedItemCount"`
	ProtectionStatus      *string                           `pulumi:"protectionStatus"`
	RegistrationStatus    *string                           `pulumi:"registrationStatus"`
	UpgradeAvailable      *bool                             `pulumi:"upgradeAvailable"`
}

type ErrorDetailResponse struct {
	Code            string   `pulumi:"code"`
	Message         string   `pulumi:"message"`
	Recommendations []string `pulumi:"recommendations"`
}

type ExtendedProperties struct {
	DiskExclusionProperties *DiskExclusionProperties `pulumi:"diskExclusionProperties"`
	LinuxVmApplicationName  *string                  `pulumi:"linuxVmApplicationName"`
}

type ExtendedPropertiesResponse struct {
	DiskExclusionProperties *DiskExclusionPropertiesResponse `pulumi:"diskExclusionProperties"`
	LinuxVmApplicationName  *string                          `pulumi:"linuxVmApplicationName"`
}

type GenericContainer struct {
	BackupManagementType  *string                       `pulumi:"backupManagementType"`
	ContainerType         string                        `pulumi:"containerType"`
	ExtendedInformation   *GenericContainerExtendedInfo `pulumi:"extendedInformation"`
	FabricName            *string                       `pulumi:"fabricName"`
	FriendlyName          *string                       `pulumi:"friendlyName"`
	HealthStatus          *string                       `pulumi:"healthStatus"`
	ProtectableObjectType *string                       `pulumi:"protectableObjectType"`
	RegistrationStatus    *string                       `pulumi:"registrationStatus"`
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
	BackupManagementType  *string                               `pulumi:"backupManagementType"`
	ContainerType         string                                `pulumi:"containerType"`
	ExtendedInformation   *GenericContainerExtendedInfoResponse `pulumi:"extendedInformation"`
	FabricName            *string                               `pulumi:"fabricName"`
	FriendlyName          *string                               `pulumi:"friendlyName"`
	HealthStatus          *string                               `pulumi:"healthStatus"`
	ProtectableObjectType *string                               `pulumi:"protectableObjectType"`
	RegistrationStatus    *string                               `pulumi:"registrationStatus"`
}

type GenericProtectedItem struct {
	BackupSetName                    *string           `pulumi:"backupSetName"`
	ContainerName                    *string           `pulumi:"containerName"`
	CreateMode                       *string           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string           `pulumi:"deferredDeleteTimeRemaining"`
	FabricName                       *string           `pulumi:"fabricName"`
	FriendlyName                     *string           `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool             `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool             `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string           `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string           `pulumi:"policyId"`
	PolicyName                       *string           `pulumi:"policyName"`
	PolicyState                      *string           `pulumi:"policyState"`
	ProtectedItemId                  *float64          `pulumi:"protectedItemId"`
	ProtectedItemType                string            `pulumi:"protectedItemType"`
	ProtectionState                  *string           `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string          `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int              `pulumi:"softDeleteRetentionPeriod"`
	SourceAssociations               map[string]string `pulumi:"sourceAssociations"`
	SourceResourceId                 *string           `pulumi:"sourceResourceId"`
}

type GenericProtectedItemResponse struct {
	BackupManagementType             string            `pulumi:"backupManagementType"`
	BackupSetName                    *string           `pulumi:"backupSetName"`
	ContainerName                    *string           `pulumi:"containerName"`
	CreateMode                       *string           `pulumi:"createMode"`
	DeferredDeleteTimeInUTC          *string           `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string           `pulumi:"deferredDeleteTimeRemaining"`
	FabricName                       *string           `pulumi:"fabricName"`
	FriendlyName                     *string           `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool             `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool             `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool             `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool             `pulumi:"isScheduledForDeferredDelete"`
	LastRecoveryPoint                *string           `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string           `pulumi:"policyId"`
	PolicyName                       *string           `pulumi:"policyName"`
	PolicyState                      *string           `pulumi:"policyState"`
	ProtectedItemId                  *float64          `pulumi:"protectedItemId"`
	ProtectedItemType                string            `pulumi:"protectedItemType"`
	ProtectionState                  *string           `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string          `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int              `pulumi:"softDeleteRetentionPeriod"`
	SourceAssociations               map[string]string `pulumi:"sourceAssociations"`
	SourceResourceId                 *string           `pulumi:"sourceResourceId"`
	WorkloadType                     string            `pulumi:"workloadType"`
}

type GenericProtectionPolicy struct {
	BackupManagementType           string                `pulumi:"backupManagementType"`
	FabricName                     *string               `pulumi:"fabricName"`
	ProtectedItemsCount            *int                  `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string              `pulumi:"resourceGuardOperationRequests"`
	SubProtectionPolicy            []SubProtectionPolicy `pulumi:"subProtectionPolicy"`
	TimeZone                       *string               `pulumi:"timeZone"`
}

type GenericProtectionPolicyResponse struct {
	BackupManagementType           string                        `pulumi:"backupManagementType"`
	FabricName                     *string                       `pulumi:"fabricName"`
	ProtectedItemsCount            *int                          `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string                      `pulumi:"resourceGuardOperationRequests"`
	SubProtectionPolicy            []SubProtectionPolicyResponse `pulumi:"subProtectionPolicy"`
	TimeZone                       *string                       `pulumi:"timeZone"`
}

type HourlySchedule struct {
	Interval                *int    `pulumi:"interval"`
	ScheduleWindowDuration  *int    `pulumi:"scheduleWindowDuration"`
	ScheduleWindowStartTime *string `pulumi:"scheduleWindowStartTime"`
}

type HourlyScheduleResponse struct {
	Interval                *int    `pulumi:"interval"`
	ScheduleWindowDuration  *int    `pulumi:"scheduleWindowDuration"`
	ScheduleWindowStartTime *string `pulumi:"scheduleWindowStartTime"`
}

type IaaSVMContainer struct {
	BackupManagementType  *string `pulumi:"backupManagementType"`
	ContainerType         string  `pulumi:"containerType"`
	FriendlyName          *string `pulumi:"friendlyName"`
	HealthStatus          *string `pulumi:"healthStatus"`
	ProtectableObjectType *string `pulumi:"protectableObjectType"`
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
	ProtectableObjectType *string `pulumi:"protectableObjectType"`
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

type ImmutabilitySettings struct {
	State *string `pulumi:"state"`
}





type ImmutabilitySettingsInput interface {
	pulumi.Input

	ToImmutabilitySettingsOutput() ImmutabilitySettingsOutput
	ToImmutabilitySettingsOutputWithContext(context.Context) ImmutabilitySettingsOutput
}

type ImmutabilitySettingsArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (ImmutabilitySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilitySettings)(nil)).Elem()
}

func (i ImmutabilitySettingsArgs) ToImmutabilitySettingsOutput() ImmutabilitySettingsOutput {
	return i.ToImmutabilitySettingsOutputWithContext(context.Background())
}

func (i ImmutabilitySettingsArgs) ToImmutabilitySettingsOutputWithContext(ctx context.Context) ImmutabilitySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilitySettingsOutput)
}

func (i ImmutabilitySettingsArgs) ToImmutabilitySettingsPtrOutput() ImmutabilitySettingsPtrOutput {
	return i.ToImmutabilitySettingsPtrOutputWithContext(context.Background())
}

func (i ImmutabilitySettingsArgs) ToImmutabilitySettingsPtrOutputWithContext(ctx context.Context) ImmutabilitySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilitySettingsOutput).ToImmutabilitySettingsPtrOutputWithContext(ctx)
}









type ImmutabilitySettingsPtrInput interface {
	pulumi.Input

	ToImmutabilitySettingsPtrOutput() ImmutabilitySettingsPtrOutput
	ToImmutabilitySettingsPtrOutputWithContext(context.Context) ImmutabilitySettingsPtrOutput
}

type immutabilitySettingsPtrType ImmutabilitySettingsArgs

func ImmutabilitySettingsPtr(v *ImmutabilitySettingsArgs) ImmutabilitySettingsPtrInput {
	return (*immutabilitySettingsPtrType)(v)
}

func (*immutabilitySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutabilitySettings)(nil)).Elem()
}

func (i *immutabilitySettingsPtrType) ToImmutabilitySettingsPtrOutput() ImmutabilitySettingsPtrOutput {
	return i.ToImmutabilitySettingsPtrOutputWithContext(context.Background())
}

func (i *immutabilitySettingsPtrType) ToImmutabilitySettingsPtrOutputWithContext(ctx context.Context) ImmutabilitySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImmutabilitySettingsPtrOutput)
}

type ImmutabilitySettingsOutput struct{ *pulumi.OutputState }

func (ImmutabilitySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilitySettings)(nil)).Elem()
}

func (o ImmutabilitySettingsOutput) ToImmutabilitySettingsOutput() ImmutabilitySettingsOutput {
	return o
}

func (o ImmutabilitySettingsOutput) ToImmutabilitySettingsOutputWithContext(ctx context.Context) ImmutabilitySettingsOutput {
	return o
}

func (o ImmutabilitySettingsOutput) ToImmutabilitySettingsPtrOutput() ImmutabilitySettingsPtrOutput {
	return o.ToImmutabilitySettingsPtrOutputWithContext(context.Background())
}

func (o ImmutabilitySettingsOutput) ToImmutabilitySettingsPtrOutputWithContext(ctx context.Context) ImmutabilitySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImmutabilitySettings) *ImmutabilitySettings {
		return &v
	}).(ImmutabilitySettingsPtrOutput)
}

func (o ImmutabilitySettingsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImmutabilitySettings) *string { return v.State }).(pulumi.StringPtrOutput)
}

type ImmutabilitySettingsPtrOutput struct{ *pulumi.OutputState }

func (ImmutabilitySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutabilitySettings)(nil)).Elem()
}

func (o ImmutabilitySettingsPtrOutput) ToImmutabilitySettingsPtrOutput() ImmutabilitySettingsPtrOutput {
	return o
}

func (o ImmutabilitySettingsPtrOutput) ToImmutabilitySettingsPtrOutputWithContext(ctx context.Context) ImmutabilitySettingsPtrOutput {
	return o
}

func (o ImmutabilitySettingsPtrOutput) Elem() ImmutabilitySettingsOutput {
	return o.ApplyT(func(v *ImmutabilitySettings) ImmutabilitySettings {
		if v != nil {
			return *v
		}
		var ret ImmutabilitySettings
		return ret
	}).(ImmutabilitySettingsOutput)
}

func (o ImmutabilitySettingsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutabilitySettings) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type ImmutabilitySettingsResponse struct {
	State *string `pulumi:"state"`
}

type ImmutabilitySettingsResponseOutput struct{ *pulumi.OutputState }

func (ImmutabilitySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImmutabilitySettingsResponse)(nil)).Elem()
}

func (o ImmutabilitySettingsResponseOutput) ToImmutabilitySettingsResponseOutput() ImmutabilitySettingsResponseOutput {
	return o
}

func (o ImmutabilitySettingsResponseOutput) ToImmutabilitySettingsResponseOutputWithContext(ctx context.Context) ImmutabilitySettingsResponseOutput {
	return o
}

func (o ImmutabilitySettingsResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImmutabilitySettingsResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type ImmutabilitySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ImmutabilitySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImmutabilitySettingsResponse)(nil)).Elem()
}

func (o ImmutabilitySettingsResponsePtrOutput) ToImmutabilitySettingsResponsePtrOutput() ImmutabilitySettingsResponsePtrOutput {
	return o
}

func (o ImmutabilitySettingsResponsePtrOutput) ToImmutabilitySettingsResponsePtrOutputWithContext(ctx context.Context) ImmutabilitySettingsResponsePtrOutput {
	return o
}

func (o ImmutabilitySettingsResponsePtrOutput) Elem() ImmutabilitySettingsResponseOutput {
	return o.ApplyT(func(v *ImmutabilitySettingsResponse) ImmutabilitySettingsResponse {
		if v != nil {
			return *v
		}
		var ret ImmutabilitySettingsResponse
		return ret
	}).(ImmutabilitySettingsResponseOutput)
}

func (o ImmutabilitySettingsResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImmutabilitySettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
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
	ProtectableObjectType     *string                     `pulumi:"protectableObjectType"`
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
	ProtectableObjectType     *string                             `pulumi:"protectableObjectType"`
	ProtectedItemCount        *float64                            `pulumi:"protectedItemCount"`
	RegistrationStatus        *string                             `pulumi:"registrationStatus"`
}

type MabFileFolderProtectedItem struct {
	BackupSetName                    *string                                 `pulumi:"backupSetName"`
	ComputerName                     *string                                 `pulumi:"computerName"`
	ContainerName                    *string                                 `pulumi:"containerName"`
	CreateMode                       *string                                 `pulumi:"createMode"`
	DeferredDeleteSyncTimeInUTC      *float64                                `pulumi:"deferredDeleteSyncTimeInUTC"`
	DeferredDeleteTimeInUTC          *string                                 `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                 `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *MabFileFolderProtectedItemExtendedInfo `pulumi:"extendedInfo"`
	FriendlyName                     *string                                 `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                                   `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                   `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                   `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                   `pulumi:"isScheduledForDeferredDelete"`
	LastBackupStatus                 *string                                 `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                 `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                 `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                 `pulumi:"policyId"`
	PolicyName                       *string                                 `pulumi:"policyName"`
	ProtectedItemType                string                                  `pulumi:"protectedItemType"`
	ProtectionState                  *string                                 `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                    `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                 `pulumi:"sourceResourceId"`
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
	BackupManagementType             string                                          `pulumi:"backupManagementType"`
	BackupSetName                    *string                                         `pulumi:"backupSetName"`
	ComputerName                     *string                                         `pulumi:"computerName"`
	ContainerName                    *string                                         `pulumi:"containerName"`
	CreateMode                       *string                                         `pulumi:"createMode"`
	DeferredDeleteSyncTimeInUTC      *float64                                        `pulumi:"deferredDeleteSyncTimeInUTC"`
	DeferredDeleteTimeInUTC          *string                                         `pulumi:"deferredDeleteTimeInUTC"`
	DeferredDeleteTimeRemaining      *string                                         `pulumi:"deferredDeleteTimeRemaining"`
	ExtendedInfo                     *MabFileFolderProtectedItemExtendedInfoResponse `pulumi:"extendedInfo"`
	FriendlyName                     *string                                         `pulumi:"friendlyName"`
	IsArchiveEnabled                 *bool                                           `pulumi:"isArchiveEnabled"`
	IsDeferredDeleteScheduleUpcoming *bool                                           `pulumi:"isDeferredDeleteScheduleUpcoming"`
	IsRehydrate                      *bool                                           `pulumi:"isRehydrate"`
	IsScheduledForDeferredDelete     *bool                                           `pulumi:"isScheduledForDeferredDelete"`
	LastBackupStatus                 *string                                         `pulumi:"lastBackupStatus"`
	LastBackupTime                   *string                                         `pulumi:"lastBackupTime"`
	LastRecoveryPoint                *string                                         `pulumi:"lastRecoveryPoint"`
	PolicyId                         *string                                         `pulumi:"policyId"`
	PolicyName                       *string                                         `pulumi:"policyName"`
	ProtectedItemType                string                                          `pulumi:"protectedItemType"`
	ProtectionState                  *string                                         `pulumi:"protectionState"`
	ResourceGuardOperationRequests   []string                                        `pulumi:"resourceGuardOperationRequests"`
	SoftDeleteRetentionPeriod        *int                                            `pulumi:"softDeleteRetentionPeriod"`
	SourceResourceId                 *string                                         `pulumi:"sourceResourceId"`
	WorkloadType                     string                                          `pulumi:"workloadType"`
}

type MabProtectionPolicy struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{} `pulumi:"schedulePolicy"`
}

type MabProtectionPolicyResponse struct {
	BackupManagementType           string      `pulumi:"backupManagementType"`
	ProtectedItemsCount            *int        `pulumi:"protectedItemsCount"`
	ResourceGuardOperationRequests []string    `pulumi:"resourceGuardOperationRequests"`
	RetentionPolicy                interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy                 interface{} `pulumi:"schedulePolicy"`
}

type MonitoringSettings struct {
	AzureMonitorAlertSettings *AzureMonitorAlertSettings `pulumi:"azureMonitorAlertSettings"`
	ClassicAlertSettings      *ClassicAlertSettings      `pulumi:"classicAlertSettings"`
}





type MonitoringSettingsInput interface {
	pulumi.Input

	ToMonitoringSettingsOutput() MonitoringSettingsOutput
	ToMonitoringSettingsOutputWithContext(context.Context) MonitoringSettingsOutput
}

type MonitoringSettingsArgs struct {
	AzureMonitorAlertSettings AzureMonitorAlertSettingsPtrInput `pulumi:"azureMonitorAlertSettings"`
	ClassicAlertSettings      ClassicAlertSettingsPtrInput      `pulumi:"classicAlertSettings"`
}

func (MonitoringSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringSettings)(nil)).Elem()
}

func (i MonitoringSettingsArgs) ToMonitoringSettingsOutput() MonitoringSettingsOutput {
	return i.ToMonitoringSettingsOutputWithContext(context.Background())
}

func (i MonitoringSettingsArgs) ToMonitoringSettingsOutputWithContext(ctx context.Context) MonitoringSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSettingsOutput)
}

func (i MonitoringSettingsArgs) ToMonitoringSettingsPtrOutput() MonitoringSettingsPtrOutput {
	return i.ToMonitoringSettingsPtrOutputWithContext(context.Background())
}

func (i MonitoringSettingsArgs) ToMonitoringSettingsPtrOutputWithContext(ctx context.Context) MonitoringSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSettingsOutput).ToMonitoringSettingsPtrOutputWithContext(ctx)
}









type MonitoringSettingsPtrInput interface {
	pulumi.Input

	ToMonitoringSettingsPtrOutput() MonitoringSettingsPtrOutput
	ToMonitoringSettingsPtrOutputWithContext(context.Context) MonitoringSettingsPtrOutput
}

type monitoringSettingsPtrType MonitoringSettingsArgs

func MonitoringSettingsPtr(v *MonitoringSettingsArgs) MonitoringSettingsPtrInput {
	return (*monitoringSettingsPtrType)(v)
}

func (*monitoringSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSettings)(nil)).Elem()
}

func (i *monitoringSettingsPtrType) ToMonitoringSettingsPtrOutput() MonitoringSettingsPtrOutput {
	return i.ToMonitoringSettingsPtrOutputWithContext(context.Background())
}

func (i *monitoringSettingsPtrType) ToMonitoringSettingsPtrOutputWithContext(ctx context.Context) MonitoringSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringSettingsPtrOutput)
}

type MonitoringSettingsOutput struct{ *pulumi.OutputState }

func (MonitoringSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringSettings)(nil)).Elem()
}

func (o MonitoringSettingsOutput) ToMonitoringSettingsOutput() MonitoringSettingsOutput {
	return o
}

func (o MonitoringSettingsOutput) ToMonitoringSettingsOutputWithContext(ctx context.Context) MonitoringSettingsOutput {
	return o
}

func (o MonitoringSettingsOutput) ToMonitoringSettingsPtrOutput() MonitoringSettingsPtrOutput {
	return o.ToMonitoringSettingsPtrOutputWithContext(context.Background())
}

func (o MonitoringSettingsOutput) ToMonitoringSettingsPtrOutputWithContext(ctx context.Context) MonitoringSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonitoringSettings) *MonitoringSettings {
		return &v
	}).(MonitoringSettingsPtrOutput)
}

func (o MonitoringSettingsOutput) AzureMonitorAlertSettings() AzureMonitorAlertSettingsPtrOutput {
	return o.ApplyT(func(v MonitoringSettings) *AzureMonitorAlertSettings { return v.AzureMonitorAlertSettings }).(AzureMonitorAlertSettingsPtrOutput)
}

func (o MonitoringSettingsOutput) ClassicAlertSettings() ClassicAlertSettingsPtrOutput {
	return o.ApplyT(func(v MonitoringSettings) *ClassicAlertSettings { return v.ClassicAlertSettings }).(ClassicAlertSettingsPtrOutput)
}

type MonitoringSettingsPtrOutput struct{ *pulumi.OutputState }

func (MonitoringSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSettings)(nil)).Elem()
}

func (o MonitoringSettingsPtrOutput) ToMonitoringSettingsPtrOutput() MonitoringSettingsPtrOutput {
	return o
}

func (o MonitoringSettingsPtrOutput) ToMonitoringSettingsPtrOutputWithContext(ctx context.Context) MonitoringSettingsPtrOutput {
	return o
}

func (o MonitoringSettingsPtrOutput) Elem() MonitoringSettingsOutput {
	return o.ApplyT(func(v *MonitoringSettings) MonitoringSettings {
		if v != nil {
			return *v
		}
		var ret MonitoringSettings
		return ret
	}).(MonitoringSettingsOutput)
}

func (o MonitoringSettingsPtrOutput) AzureMonitorAlertSettings() AzureMonitorAlertSettingsPtrOutput {
	return o.ApplyT(func(v *MonitoringSettings) *AzureMonitorAlertSettings {
		if v == nil {
			return nil
		}
		return v.AzureMonitorAlertSettings
	}).(AzureMonitorAlertSettingsPtrOutput)
}

func (o MonitoringSettingsPtrOutput) ClassicAlertSettings() ClassicAlertSettingsPtrOutput {
	return o.ApplyT(func(v *MonitoringSettings) *ClassicAlertSettings {
		if v == nil {
			return nil
		}
		return v.ClassicAlertSettings
	}).(ClassicAlertSettingsPtrOutput)
}

type MonitoringSettingsResponse struct {
	AzureMonitorAlertSettings *AzureMonitorAlertSettingsResponse `pulumi:"azureMonitorAlertSettings"`
	ClassicAlertSettings      *ClassicAlertSettingsResponse      `pulumi:"classicAlertSettings"`
}

type MonitoringSettingsResponseOutput struct{ *pulumi.OutputState }

func (MonitoringSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringSettingsResponse)(nil)).Elem()
}

func (o MonitoringSettingsResponseOutput) ToMonitoringSettingsResponseOutput() MonitoringSettingsResponseOutput {
	return o
}

func (o MonitoringSettingsResponseOutput) ToMonitoringSettingsResponseOutputWithContext(ctx context.Context) MonitoringSettingsResponseOutput {
	return o
}

func (o MonitoringSettingsResponseOutput) AzureMonitorAlertSettings() AzureMonitorAlertSettingsResponsePtrOutput {
	return o.ApplyT(func(v MonitoringSettingsResponse) *AzureMonitorAlertSettingsResponse {
		return v.AzureMonitorAlertSettings
	}).(AzureMonitorAlertSettingsResponsePtrOutput)
}

func (o MonitoringSettingsResponseOutput) ClassicAlertSettings() ClassicAlertSettingsResponsePtrOutput {
	return o.ApplyT(func(v MonitoringSettingsResponse) *ClassicAlertSettingsResponse { return v.ClassicAlertSettings }).(ClassicAlertSettingsResponsePtrOutput)
}

type MonitoringSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (MonitoringSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringSettingsResponse)(nil)).Elem()
}

func (o MonitoringSettingsResponsePtrOutput) ToMonitoringSettingsResponsePtrOutput() MonitoringSettingsResponsePtrOutput {
	return o
}

func (o MonitoringSettingsResponsePtrOutput) ToMonitoringSettingsResponsePtrOutputWithContext(ctx context.Context) MonitoringSettingsResponsePtrOutput {
	return o
}

func (o MonitoringSettingsResponsePtrOutput) Elem() MonitoringSettingsResponseOutput {
	return o.ApplyT(func(v *MonitoringSettingsResponse) MonitoringSettingsResponse {
		if v != nil {
			return *v
		}
		var ret MonitoringSettingsResponse
		return ret
	}).(MonitoringSettingsResponseOutput)
}

func (o MonitoringSettingsResponsePtrOutput) AzureMonitorAlertSettings() AzureMonitorAlertSettingsResponsePtrOutput {
	return o.ApplyT(func(v *MonitoringSettingsResponse) *AzureMonitorAlertSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AzureMonitorAlertSettings
	}).(AzureMonitorAlertSettingsResponsePtrOutput)
}

func (o MonitoringSettingsResponsePtrOutput) ClassicAlertSettings() ClassicAlertSettingsResponsePtrOutput {
	return o.ApplyT(func(v *MonitoringSettingsResponse) *ClassicAlertSettingsResponse {
		if v == nil {
			return nil
		}
		return v.ClassicAlertSettings
	}).(ClassicAlertSettingsResponsePtrOutput)
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
	Location   string                                 `pulumi:"location"`
	Name       string                                 `pulumi:"name"`
	Properties VaultPrivateEndpointConnectionResponse `pulumi:"properties"`
	Type       string                                 `pulumi:"type"`
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

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Properties() VaultPrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) VaultPrivateEndpointConnectionResponse {
		return v.Properties
	}).(VaultPrivateEndpointConnectionResponseOutput)
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
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

type ResourceGuardOperationDetail struct {
	DefaultResourceRequest *string `pulumi:"defaultResourceRequest"`
	VaultCriticalOperation *string `pulumi:"vaultCriticalOperation"`
}





type ResourceGuardOperationDetailInput interface {
	pulumi.Input

	ToResourceGuardOperationDetailOutput() ResourceGuardOperationDetailOutput
	ToResourceGuardOperationDetailOutputWithContext(context.Context) ResourceGuardOperationDetailOutput
}

type ResourceGuardOperationDetailArgs struct {
	DefaultResourceRequest pulumi.StringPtrInput `pulumi:"defaultResourceRequest"`
	VaultCriticalOperation pulumi.StringPtrInput `pulumi:"vaultCriticalOperation"`
}

func (ResourceGuardOperationDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuardOperationDetail)(nil)).Elem()
}

func (i ResourceGuardOperationDetailArgs) ToResourceGuardOperationDetailOutput() ResourceGuardOperationDetailOutput {
	return i.ToResourceGuardOperationDetailOutputWithContext(context.Background())
}

func (i ResourceGuardOperationDetailArgs) ToResourceGuardOperationDetailOutputWithContext(ctx context.Context) ResourceGuardOperationDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardOperationDetailOutput)
}





type ResourceGuardOperationDetailArrayInput interface {
	pulumi.Input

	ToResourceGuardOperationDetailArrayOutput() ResourceGuardOperationDetailArrayOutput
	ToResourceGuardOperationDetailArrayOutputWithContext(context.Context) ResourceGuardOperationDetailArrayOutput
}

type ResourceGuardOperationDetailArray []ResourceGuardOperationDetailInput

func (ResourceGuardOperationDetailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceGuardOperationDetail)(nil)).Elem()
}

func (i ResourceGuardOperationDetailArray) ToResourceGuardOperationDetailArrayOutput() ResourceGuardOperationDetailArrayOutput {
	return i.ToResourceGuardOperationDetailArrayOutputWithContext(context.Background())
}

func (i ResourceGuardOperationDetailArray) ToResourceGuardOperationDetailArrayOutputWithContext(ctx context.Context) ResourceGuardOperationDetailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardOperationDetailArrayOutput)
}

type ResourceGuardOperationDetailOutput struct{ *pulumi.OutputState }

func (ResourceGuardOperationDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuardOperationDetail)(nil)).Elem()
}

func (o ResourceGuardOperationDetailOutput) ToResourceGuardOperationDetailOutput() ResourceGuardOperationDetailOutput {
	return o
}

func (o ResourceGuardOperationDetailOutput) ToResourceGuardOperationDetailOutputWithContext(ctx context.Context) ResourceGuardOperationDetailOutput {
	return o
}

func (o ResourceGuardOperationDetailOutput) DefaultResourceRequest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardOperationDetail) *string { return v.DefaultResourceRequest }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardOperationDetailOutput) VaultCriticalOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardOperationDetail) *string { return v.VaultCriticalOperation }).(pulumi.StringPtrOutput)
}

type ResourceGuardOperationDetailArrayOutput struct{ *pulumi.OutputState }

func (ResourceGuardOperationDetailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResourceGuardOperationDetail)(nil)).Elem()
}

func (o ResourceGuardOperationDetailArrayOutput) ToResourceGuardOperationDetailArrayOutput() ResourceGuardOperationDetailArrayOutput {
	return o
}

func (o ResourceGuardOperationDetailArrayOutput) ToResourceGuardOperationDetailArrayOutputWithContext(ctx context.Context) ResourceGuardOperationDetailArrayOutput {
	return o
}

func (o ResourceGuardOperationDetailArrayOutput) Index(i pulumi.IntInput) ResourceGuardOperationDetailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResourceGuardOperationDetail {
		return vs[0].([]ResourceGuardOperationDetail)[vs[1].(int)]
	}).(ResourceGuardOperationDetailOutput)
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

type ResourceGuardProxyBase struct {
	Description                   *string                        `pulumi:"description"`
	LastUpdatedTime               *string                        `pulumi:"lastUpdatedTime"`
	ResourceGuardOperationDetails []ResourceGuardOperationDetail `pulumi:"resourceGuardOperationDetails"`
	ResourceGuardResourceId       *string                        `pulumi:"resourceGuardResourceId"`
}





type ResourceGuardProxyBaseInput interface {
	pulumi.Input

	ToResourceGuardProxyBaseOutput() ResourceGuardProxyBaseOutput
	ToResourceGuardProxyBaseOutputWithContext(context.Context) ResourceGuardProxyBaseOutput
}

type ResourceGuardProxyBaseArgs struct {
	Description                   pulumi.StringPtrInput                  `pulumi:"description"`
	LastUpdatedTime               pulumi.StringPtrInput                  `pulumi:"lastUpdatedTime"`
	ResourceGuardOperationDetails ResourceGuardOperationDetailArrayInput `pulumi:"resourceGuardOperationDetails"`
	ResourceGuardResourceId       pulumi.StringPtrInput                  `pulumi:"resourceGuardResourceId"`
}

func (ResourceGuardProxyBaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuardProxyBase)(nil)).Elem()
}

func (i ResourceGuardProxyBaseArgs) ToResourceGuardProxyBaseOutput() ResourceGuardProxyBaseOutput {
	return i.ToResourceGuardProxyBaseOutputWithContext(context.Background())
}

func (i ResourceGuardProxyBaseArgs) ToResourceGuardProxyBaseOutputWithContext(ctx context.Context) ResourceGuardProxyBaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardProxyBaseOutput)
}

func (i ResourceGuardProxyBaseArgs) ToResourceGuardProxyBasePtrOutput() ResourceGuardProxyBasePtrOutput {
	return i.ToResourceGuardProxyBasePtrOutputWithContext(context.Background())
}

func (i ResourceGuardProxyBaseArgs) ToResourceGuardProxyBasePtrOutputWithContext(ctx context.Context) ResourceGuardProxyBasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardProxyBaseOutput).ToResourceGuardProxyBasePtrOutputWithContext(ctx)
}









type ResourceGuardProxyBasePtrInput interface {
	pulumi.Input

	ToResourceGuardProxyBasePtrOutput() ResourceGuardProxyBasePtrOutput
	ToResourceGuardProxyBasePtrOutputWithContext(context.Context) ResourceGuardProxyBasePtrOutput
}

type resourceGuardProxyBasePtrType ResourceGuardProxyBaseArgs

func ResourceGuardProxyBasePtr(v *ResourceGuardProxyBaseArgs) ResourceGuardProxyBasePtrInput {
	return (*resourceGuardProxyBasePtrType)(v)
}

func (*resourceGuardProxyBasePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuardProxyBase)(nil)).Elem()
}

func (i *resourceGuardProxyBasePtrType) ToResourceGuardProxyBasePtrOutput() ResourceGuardProxyBasePtrOutput {
	return i.ToResourceGuardProxyBasePtrOutputWithContext(context.Background())
}

func (i *resourceGuardProxyBasePtrType) ToResourceGuardProxyBasePtrOutputWithContext(ctx context.Context) ResourceGuardProxyBasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceGuardProxyBasePtrOutput)
}

type ResourceGuardProxyBaseOutput struct{ *pulumi.OutputState }

func (ResourceGuardProxyBaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceGuardProxyBase)(nil)).Elem()
}

func (o ResourceGuardProxyBaseOutput) ToResourceGuardProxyBaseOutput() ResourceGuardProxyBaseOutput {
	return o
}

func (o ResourceGuardProxyBaseOutput) ToResourceGuardProxyBaseOutputWithContext(ctx context.Context) ResourceGuardProxyBaseOutput {
	return o
}

func (o ResourceGuardProxyBaseOutput) ToResourceGuardProxyBasePtrOutput() ResourceGuardProxyBasePtrOutput {
	return o.ToResourceGuardProxyBasePtrOutputWithContext(context.Background())
}

func (o ResourceGuardProxyBaseOutput) ToResourceGuardProxyBasePtrOutputWithContext(ctx context.Context) ResourceGuardProxyBasePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceGuardProxyBase) *ResourceGuardProxyBase {
		return &v
	}).(ResourceGuardProxyBasePtrOutput)
}

func (o ResourceGuardProxyBaseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardProxyBase) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyBaseOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardProxyBase) *string { return v.LastUpdatedTime }).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyBaseOutput) ResourceGuardOperationDetails() ResourceGuardOperationDetailArrayOutput {
	return o.ApplyT(func(v ResourceGuardProxyBase) []ResourceGuardOperationDetail { return v.ResourceGuardOperationDetails }).(ResourceGuardOperationDetailArrayOutput)
}

func (o ResourceGuardProxyBaseOutput) ResourceGuardResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardProxyBase) *string { return v.ResourceGuardResourceId }).(pulumi.StringPtrOutput)
}

type ResourceGuardProxyBasePtrOutput struct{ *pulumi.OutputState }

func (ResourceGuardProxyBasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceGuardProxyBase)(nil)).Elem()
}

func (o ResourceGuardProxyBasePtrOutput) ToResourceGuardProxyBasePtrOutput() ResourceGuardProxyBasePtrOutput {
	return o
}

func (o ResourceGuardProxyBasePtrOutput) ToResourceGuardProxyBasePtrOutputWithContext(ctx context.Context) ResourceGuardProxyBasePtrOutput {
	return o
}

func (o ResourceGuardProxyBasePtrOutput) Elem() ResourceGuardProxyBaseOutput {
	return o.ApplyT(func(v *ResourceGuardProxyBase) ResourceGuardProxyBase {
		if v != nil {
			return *v
		}
		var ret ResourceGuardProxyBase
		return ret
	}).(ResourceGuardProxyBaseOutput)
}

func (o ResourceGuardProxyBasePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuardProxyBase) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyBasePtrOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuardProxyBase) *string {
		if v == nil {
			return nil
		}
		return v.LastUpdatedTime
	}).(pulumi.StringPtrOutput)
}

func (o ResourceGuardProxyBasePtrOutput) ResourceGuardOperationDetails() ResourceGuardOperationDetailArrayOutput {
	return o.ApplyT(func(v *ResourceGuardProxyBase) []ResourceGuardOperationDetail {
		if v == nil {
			return nil
		}
		return v.ResourceGuardOperationDetails
	}).(ResourceGuardOperationDetailArrayOutput)
}

func (o ResourceGuardProxyBasePtrOutput) ResourceGuardResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceGuardProxyBase) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGuardResourceId
	}).(pulumi.StringPtrOutput)
}

type ResourceGuardProxyBaseResponse struct {
	Description                   *string                                `pulumi:"description"`
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

func (o ResourceGuardProxyBaseResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceGuardProxyBaseResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
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

type SecuritySettings struct {
	ImmutabilitySettings *ImmutabilitySettings `pulumi:"immutabilitySettings"`
}





type SecuritySettingsInput interface {
	pulumi.Input

	ToSecuritySettingsOutput() SecuritySettingsOutput
	ToSecuritySettingsOutputWithContext(context.Context) SecuritySettingsOutput
}

type SecuritySettingsArgs struct {
	ImmutabilitySettings ImmutabilitySettingsPtrInput `pulumi:"immutabilitySettings"`
}

func (SecuritySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecuritySettings)(nil)).Elem()
}

func (i SecuritySettingsArgs) ToSecuritySettingsOutput() SecuritySettingsOutput {
	return i.ToSecuritySettingsOutputWithContext(context.Background())
}

func (i SecuritySettingsArgs) ToSecuritySettingsOutputWithContext(ctx context.Context) SecuritySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecuritySettingsOutput)
}

func (i SecuritySettingsArgs) ToSecuritySettingsPtrOutput() SecuritySettingsPtrOutput {
	return i.ToSecuritySettingsPtrOutputWithContext(context.Background())
}

func (i SecuritySettingsArgs) ToSecuritySettingsPtrOutputWithContext(ctx context.Context) SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecuritySettingsOutput).ToSecuritySettingsPtrOutputWithContext(ctx)
}









type SecuritySettingsPtrInput interface {
	pulumi.Input

	ToSecuritySettingsPtrOutput() SecuritySettingsPtrOutput
	ToSecuritySettingsPtrOutputWithContext(context.Context) SecuritySettingsPtrOutput
}

type securitySettingsPtrType SecuritySettingsArgs

func SecuritySettingsPtr(v *SecuritySettingsArgs) SecuritySettingsPtrInput {
	return (*securitySettingsPtrType)(v)
}

func (*securitySettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecuritySettings)(nil)).Elem()
}

func (i *securitySettingsPtrType) ToSecuritySettingsPtrOutput() SecuritySettingsPtrOutput {
	return i.ToSecuritySettingsPtrOutputWithContext(context.Background())
}

func (i *securitySettingsPtrType) ToSecuritySettingsPtrOutputWithContext(ctx context.Context) SecuritySettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecuritySettingsPtrOutput)
}

type SecuritySettingsOutput struct{ *pulumi.OutputState }

func (SecuritySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecuritySettings)(nil)).Elem()
}

func (o SecuritySettingsOutput) ToSecuritySettingsOutput() SecuritySettingsOutput {
	return o
}

func (o SecuritySettingsOutput) ToSecuritySettingsOutputWithContext(ctx context.Context) SecuritySettingsOutput {
	return o
}

func (o SecuritySettingsOutput) ToSecuritySettingsPtrOutput() SecuritySettingsPtrOutput {
	return o.ToSecuritySettingsPtrOutputWithContext(context.Background())
}

func (o SecuritySettingsOutput) ToSecuritySettingsPtrOutputWithContext(ctx context.Context) SecuritySettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecuritySettings) *SecuritySettings {
		return &v
	}).(SecuritySettingsPtrOutput)
}

func (o SecuritySettingsOutput) ImmutabilitySettings() ImmutabilitySettingsPtrOutput {
	return o.ApplyT(func(v SecuritySettings) *ImmutabilitySettings { return v.ImmutabilitySettings }).(ImmutabilitySettingsPtrOutput)
}

type SecuritySettingsPtrOutput struct{ *pulumi.OutputState }

func (SecuritySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecuritySettings)(nil)).Elem()
}

func (o SecuritySettingsPtrOutput) ToSecuritySettingsPtrOutput() SecuritySettingsPtrOutput {
	return o
}

func (o SecuritySettingsPtrOutput) ToSecuritySettingsPtrOutputWithContext(ctx context.Context) SecuritySettingsPtrOutput {
	return o
}

func (o SecuritySettingsPtrOutput) Elem() SecuritySettingsOutput {
	return o.ApplyT(func(v *SecuritySettings) SecuritySettings {
		if v != nil {
			return *v
		}
		var ret SecuritySettings
		return ret
	}).(SecuritySettingsOutput)
}

func (o SecuritySettingsPtrOutput) ImmutabilitySettings() ImmutabilitySettingsPtrOutput {
	return o.ApplyT(func(v *SecuritySettings) *ImmutabilitySettings {
		if v == nil {
			return nil
		}
		return v.ImmutabilitySettings
	}).(ImmutabilitySettingsPtrOutput)
}

type SecuritySettingsResponse struct {
	ImmutabilitySettings *ImmutabilitySettingsResponse `pulumi:"immutabilitySettings"`
}

type SecuritySettingsResponseOutput struct{ *pulumi.OutputState }

func (SecuritySettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecuritySettingsResponse)(nil)).Elem()
}

func (o SecuritySettingsResponseOutput) ToSecuritySettingsResponseOutput() SecuritySettingsResponseOutput {
	return o
}

func (o SecuritySettingsResponseOutput) ToSecuritySettingsResponseOutputWithContext(ctx context.Context) SecuritySettingsResponseOutput {
	return o
}

func (o SecuritySettingsResponseOutput) ImmutabilitySettings() ImmutabilitySettingsResponsePtrOutput {
	return o.ApplyT(func(v SecuritySettingsResponse) *ImmutabilitySettingsResponse { return v.ImmutabilitySettings }).(ImmutabilitySettingsResponsePtrOutput)
}

type SecuritySettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SecuritySettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecuritySettingsResponse)(nil)).Elem()
}

func (o SecuritySettingsResponsePtrOutput) ToSecuritySettingsResponsePtrOutput() SecuritySettingsResponsePtrOutput {
	return o
}

func (o SecuritySettingsResponsePtrOutput) ToSecuritySettingsResponsePtrOutputWithContext(ctx context.Context) SecuritySettingsResponsePtrOutput {
	return o
}

func (o SecuritySettingsResponsePtrOutput) Elem() SecuritySettingsResponseOutput {
	return o.ApplyT(func(v *SecuritySettingsResponse) SecuritySettingsResponse {
		if v != nil {
			return *v
		}
		var ret SecuritySettingsResponse
		return ret
	}).(SecuritySettingsResponseOutput)
}

func (o SecuritySettingsResponsePtrOutput) ImmutabilitySettings() ImmutabilitySettingsResponsePtrOutput {
	return o.ApplyT(func(v *SecuritySettingsResponse) *ImmutabilitySettingsResponse {
		if v == nil {
			return nil
		}
		return v.ImmutabilitySettings
	}).(ImmutabilitySettingsResponsePtrOutput)
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
	HourlySchedule          *HourlySchedule `pulumi:"hourlySchedule"`
	SchedulePolicyType      string          `pulumi:"schedulePolicyType"`
	ScheduleRunDays         []DayOfWeek     `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    *string         `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        []string        `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency *int            `pulumi:"scheduleWeeklyFrequency"`
}

type SimpleSchedulePolicyResponse struct {
	HourlySchedule          *HourlyScheduleResponse `pulumi:"hourlySchedule"`
	SchedulePolicyType      string                  `pulumi:"schedulePolicyType"`
	ScheduleRunDays         []string                `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    *string                 `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        []string                `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency *int                    `pulumi:"scheduleWeeklyFrequency"`
}

type SimpleSchedulePolicyV2 struct {
	DailySchedule        *DailySchedule  `pulumi:"dailySchedule"`
	HourlySchedule       *HourlySchedule `pulumi:"hourlySchedule"`
	SchedulePolicyType   string          `pulumi:"schedulePolicyType"`
	ScheduleRunFrequency *string         `pulumi:"scheduleRunFrequency"`
	WeeklySchedule       *WeeklySchedule `pulumi:"weeklySchedule"`
}

type SimpleSchedulePolicyV2Response struct {
	DailySchedule        *DailyScheduleResponse  `pulumi:"dailySchedule"`
	HourlySchedule       *HourlyScheduleResponse `pulumi:"hourlySchedule"`
	SchedulePolicyType   string                  `pulumi:"schedulePolicyType"`
	ScheduleRunFrequency *string                 `pulumi:"scheduleRunFrequency"`
	WeeklySchedule       *WeeklyScheduleResponse `pulumi:"weeklySchedule"`
}

type Sku struct {
	Capacity *string `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.StringPtrInput `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Capacity }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
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
	Capacity *string `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Capacity }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
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
	PolicyType      *string                  `pulumi:"policyType"`
	RetentionPolicy interface{}              `pulumi:"retentionPolicy"`
	SchedulePolicy  interface{}              `pulumi:"schedulePolicy"`
	TieringPolicy   map[string]TieringPolicy `pulumi:"tieringPolicy"`
}

type SubProtectionPolicyResponse struct {
	PolicyType      *string                          `pulumi:"policyType"`
	RetentionPolicy interface{}                      `pulumi:"retentionPolicy"`
	SchedulePolicy  interface{}                      `pulumi:"schedulePolicy"`
	TieringPolicy   map[string]TieringPolicyResponse `pulumi:"tieringPolicy"`
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

type TieringPolicy struct {
	Duration     *int    `pulumi:"duration"`
	DurationType *string `pulumi:"durationType"`
	TieringMode  *string `pulumi:"tieringMode"`
}

type TieringPolicyResponse struct {
	Duration     *int    `pulumi:"duration"`
	DurationType *string `pulumi:"durationType"`
	TieringMode  *string `pulumi:"tieringMode"`
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
	GroupIds                          []string                                       `pulumi:"groupIds"`
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

func (o VaultPrivateEndpointConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
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
	Encryption          *VaultPropertiesEncryption `pulumi:"encryption"`
	MonitoringSettings  *MonitoringSettings        `pulumi:"monitoringSettings"`
	PublicNetworkAccess *string                    `pulumi:"publicNetworkAccess"`
	SecuritySettings    *SecuritySettings          `pulumi:"securitySettings"`
}





type VaultPropertiesInput interface {
	pulumi.Input

	ToVaultPropertiesOutput() VaultPropertiesOutput
	ToVaultPropertiesOutputWithContext(context.Context) VaultPropertiesOutput
}

type VaultPropertiesArgs struct {
	Encryption          VaultPropertiesEncryptionPtrInput `pulumi:"encryption"`
	MonitoringSettings  MonitoringSettingsPtrInput        `pulumi:"monitoringSettings"`
	PublicNetworkAccess pulumi.StringPtrInput             `pulumi:"publicNetworkAccess"`
	SecuritySettings    SecuritySettingsPtrInput          `pulumi:"securitySettings"`
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

func (o VaultPropertiesOutput) MonitoringSettings() MonitoringSettingsPtrOutput {
	return o.ApplyT(func(v VaultProperties) *MonitoringSettings { return v.MonitoringSettings }).(MonitoringSettingsPtrOutput)
}

func (o VaultPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesOutput) SecuritySettings() SecuritySettingsPtrOutput {
	return o.ApplyT(func(v VaultProperties) *SecuritySettings { return v.SecuritySettings }).(SecuritySettingsPtrOutput)
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

func (o VaultPropertiesPtrOutput) MonitoringSettings() MonitoringSettingsPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *MonitoringSettings {
		if v == nil {
			return nil
		}
		return v.MonitoringSettings
	}).(MonitoringSettingsPtrOutput)
}

func (o VaultPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesPtrOutput) SecuritySettings() SecuritySettingsPtrOutput {
	return o.ApplyT(func(v *VaultProperties) *SecuritySettings {
		if v == nil {
			return nil
		}
		return v.SecuritySettings
	}).(SecuritySettingsPtrOutput)
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
	BackupStorageVersion                string                                             `pulumi:"backupStorageVersion"`
	Encryption                          *VaultPropertiesResponseEncryption                 `pulumi:"encryption"`
	MonitoringSettings                  *MonitoringSettingsResponse                        `pulumi:"monitoringSettings"`
	MoveDetails                         *VaultPropertiesResponseMoveDetails                `pulumi:"moveDetails"`
	MoveState                           string                                             `pulumi:"moveState"`
	PrivateEndpointConnections          []PrivateEndpointConnectionVaultPropertiesResponse `pulumi:"privateEndpointConnections"`
	PrivateEndpointStateForBackup       string                                             `pulumi:"privateEndpointStateForBackup"`
	PrivateEndpointStateForSiteRecovery string                                             `pulumi:"privateEndpointStateForSiteRecovery"`
	ProvisioningState                   string                                             `pulumi:"provisioningState"`
	PublicNetworkAccess                 *string                                            `pulumi:"publicNetworkAccess"`
	RedundancySettings                  *VaultPropertiesResponseRedundancySettings         `pulumi:"redundancySettings"`
	SecuritySettings                    *SecuritySettingsResponse                          `pulumi:"securitySettings"`
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

func (o VaultPropertiesResponseOutput) BackupStorageVersion() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.BackupStorageVersion }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) Encryption() VaultPropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *VaultPropertiesResponseEncryption { return v.Encryption }).(VaultPropertiesResponseEncryptionPtrOutput)
}

func (o VaultPropertiesResponseOutput) MonitoringSettings() MonitoringSettingsResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *MonitoringSettingsResponse { return v.MonitoringSettings }).(MonitoringSettingsResponsePtrOutput)
}

func (o VaultPropertiesResponseOutput) MoveDetails() VaultPropertiesResponseMoveDetailsPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *VaultPropertiesResponseMoveDetails { return v.MoveDetails }).(VaultPropertiesResponseMoveDetailsPtrOutput)
}

func (o VaultPropertiesResponseOutput) MoveState() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.MoveState }).(pulumi.StringOutput)
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

func (o VaultPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseOutput) RedundancySettings() VaultPropertiesResponseRedundancySettingsPtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *VaultPropertiesResponseRedundancySettings {
		return v.RedundancySettings
	}).(VaultPropertiesResponseRedundancySettingsPtrOutput)
}

func (o VaultPropertiesResponseOutput) SecuritySettings() SecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *SecuritySettingsResponse { return v.SecuritySettings }).(SecuritySettingsResponsePtrOutput)
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

type VaultPropertiesResponseMoveDetails struct {
	CompletionTimeUtc string `pulumi:"completionTimeUtc"`
	OperationId       string `pulumi:"operationId"`
	SourceResourceId  string `pulumi:"sourceResourceId"`
	StartTimeUtc      string `pulumi:"startTimeUtc"`
	TargetResourceId  string `pulumi:"targetResourceId"`
}

type VaultPropertiesResponseMoveDetailsOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseMoveDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponseMoveDetails)(nil)).Elem()
}

func (o VaultPropertiesResponseMoveDetailsOutput) ToVaultPropertiesResponseMoveDetailsOutput() VaultPropertiesResponseMoveDetailsOutput {
	return o
}

func (o VaultPropertiesResponseMoveDetailsOutput) ToVaultPropertiesResponseMoveDetailsOutputWithContext(ctx context.Context) VaultPropertiesResponseMoveDetailsOutput {
	return o
}

func (o VaultPropertiesResponseMoveDetailsOutput) CompletionTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseMoveDetails) string { return v.CompletionTimeUtc }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseMoveDetailsOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseMoveDetails) string { return v.OperationId }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseMoveDetailsOutput) SourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseMoveDetails) string { return v.SourceResourceId }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseMoveDetailsOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseMoveDetails) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseMoveDetailsOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseMoveDetails) string { return v.TargetResourceId }).(pulumi.StringOutput)
}

type VaultPropertiesResponseMoveDetailsPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseMoveDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponseMoveDetails)(nil)).Elem()
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) ToVaultPropertiesResponseMoveDetailsPtrOutput() VaultPropertiesResponseMoveDetailsPtrOutput {
	return o
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) ToVaultPropertiesResponseMoveDetailsPtrOutputWithContext(ctx context.Context) VaultPropertiesResponseMoveDetailsPtrOutput {
	return o
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) Elem() VaultPropertiesResponseMoveDetailsOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) VaultPropertiesResponseMoveDetails {
		if v != nil {
			return *v
		}
		var ret VaultPropertiesResponseMoveDetails
		return ret
	}).(VaultPropertiesResponseMoveDetailsOutput)
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) CompletionTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) *string {
		if v == nil {
			return nil
		}
		return &v.CompletionTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) *string {
		if v == nil {
			return nil
		}
		return &v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) *string {
		if v == nil {
			return nil
		}
		return &v.SourceResourceId
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) *string {
		if v == nil {
			return nil
		}
		return &v.StartTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseMoveDetailsPtrOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseMoveDetails) *string {
		if v == nil {
			return nil
		}
		return &v.TargetResourceId
	}).(pulumi.StringPtrOutput)
}

type VaultPropertiesResponseRedundancySettings struct {
	CrossRegionRestore            string `pulumi:"crossRegionRestore"`
	StandardTierStorageRedundancy string `pulumi:"standardTierStorageRedundancy"`
}

type VaultPropertiesResponseRedundancySettingsOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseRedundancySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponseRedundancySettings)(nil)).Elem()
}

func (o VaultPropertiesResponseRedundancySettingsOutput) ToVaultPropertiesResponseRedundancySettingsOutput() VaultPropertiesResponseRedundancySettingsOutput {
	return o
}

func (o VaultPropertiesResponseRedundancySettingsOutput) ToVaultPropertiesResponseRedundancySettingsOutputWithContext(ctx context.Context) VaultPropertiesResponseRedundancySettingsOutput {
	return o
}

func (o VaultPropertiesResponseRedundancySettingsOutput) CrossRegionRestore() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseRedundancySettings) string { return v.CrossRegionRestore }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseRedundancySettingsOutput) StandardTierStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponseRedundancySettings) string { return v.StandardTierStorageRedundancy }).(pulumi.StringOutput)
}

type VaultPropertiesResponseRedundancySettingsPtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseRedundancySettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponseRedundancySettings)(nil)).Elem()
}

func (o VaultPropertiesResponseRedundancySettingsPtrOutput) ToVaultPropertiesResponseRedundancySettingsPtrOutput() VaultPropertiesResponseRedundancySettingsPtrOutput {
	return o
}

func (o VaultPropertiesResponseRedundancySettingsPtrOutput) ToVaultPropertiesResponseRedundancySettingsPtrOutputWithContext(ctx context.Context) VaultPropertiesResponseRedundancySettingsPtrOutput {
	return o
}

func (o VaultPropertiesResponseRedundancySettingsPtrOutput) Elem() VaultPropertiesResponseRedundancySettingsOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseRedundancySettings) VaultPropertiesResponseRedundancySettings {
		if v != nil {
			return *v
		}
		var ret VaultPropertiesResponseRedundancySettings
		return ret
	}).(VaultPropertiesResponseRedundancySettingsOutput)
}

func (o VaultPropertiesResponseRedundancySettingsPtrOutput) CrossRegionRestore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseRedundancySettings) *string {
		if v == nil {
			return nil
		}
		return &v.CrossRegionRestore
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponseRedundancySettingsPtrOutput) StandardTierStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponseRedundancySettings) *string {
		if v == nil {
			return nil
		}
		return &v.StandardTierStorageRedundancy
	}).(pulumi.StringPtrOutput)
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

type WeeklySchedule struct {
	ScheduleRunDays  []DayOfWeek `pulumi:"scheduleRunDays"`
	ScheduleRunTimes []string    `pulumi:"scheduleRunTimes"`
}

type WeeklyScheduleResponse struct {
	ScheduleRunDays  []string `pulumi:"scheduleRunDays"`
	ScheduleRunTimes []string `pulumi:"scheduleRunTimes"`
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
	pulumi.RegisterOutputType(AzureMonitorAlertSettingsOutput{})
	pulumi.RegisterOutputType(AzureMonitorAlertSettingsPtrOutput{})
	pulumi.RegisterOutputType(AzureMonitorAlertSettingsResponseOutput{})
	pulumi.RegisterOutputType(AzureMonitorAlertSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ClassicAlertSettingsOutput{})
	pulumi.RegisterOutputType(ClassicAlertSettingsPtrOutput{})
	pulumi.RegisterOutputType(ClassicAlertSettingsResponseOutput{})
	pulumi.RegisterOutputType(ClassicAlertSettingsResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ImmutabilitySettingsOutput{})
	pulumi.RegisterOutputType(ImmutabilitySettingsPtrOutput{})
	pulumi.RegisterOutputType(ImmutabilitySettingsResponseOutput{})
	pulumi.RegisterOutputType(ImmutabilitySettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(MonitoringSettingsOutput{})
	pulumi.RegisterOutputType(MonitoringSettingsPtrOutput{})
	pulumi.RegisterOutputType(MonitoringSettingsResponseOutput{})
	pulumi.RegisterOutputType(MonitoringSettingsResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ResourceGuardOperationDetailOutput{})
	pulumi.RegisterOutputType(ResourceGuardOperationDetailArrayOutput{})
	pulumi.RegisterOutputType(ResourceGuardOperationDetailResponseOutput{})
	pulumi.RegisterOutputType(ResourceGuardOperationDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceGuardProxyBaseOutput{})
	pulumi.RegisterOutputType(ResourceGuardProxyBasePtrOutput{})
	pulumi.RegisterOutputType(ResourceGuardProxyBaseResponseOutput{})
	pulumi.RegisterOutputType(SecuritySettingsOutput{})
	pulumi.RegisterOutputType(SecuritySettingsPtrOutput{})
	pulumi.RegisterOutputType(SecuritySettingsResponseOutput{})
	pulumi.RegisterOutputType(SecuritySettingsResponsePtrOutput{})
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
	pulumi.RegisterOutputType(VaultPropertiesResponseMoveDetailsOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseMoveDetailsPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseRedundancySettingsOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseRedundancySettingsPtrOutput{})
}
