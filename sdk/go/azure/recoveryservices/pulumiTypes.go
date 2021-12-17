


package recoveryservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type A2AContainerMappingInput struct {
	AgentAutoUpdateStatus  *string `pulumi:"agentAutoUpdateStatus"`
	AutomationAccountArmId *string `pulumi:"automationAccountArmId"`
	InstanceType           *string `pulumi:"instanceType"`
}

type A2AEnableProtectionInput struct {
	DiskEncryptionInfo                *DiskEncryptionInfo            `pulumi:"diskEncryptionInfo"`
	FabricObjectId                    *string                        `pulumi:"fabricObjectId"`
	InstanceType                      *string                        `pulumi:"instanceType"`
	MultiVmGroupName                  *string                        `pulumi:"multiVmGroupName"`
	RecoveryAvailabilitySetId         *string                        `pulumi:"recoveryAvailabilitySetId"`
	RecoveryAvailabilityZone          *string                        `pulumi:"recoveryAvailabilityZone"`
	RecoveryAzureNetworkId            *string                        `pulumi:"recoveryAzureNetworkId"`
	RecoveryBootDiagStorageAccountId  *string                        `pulumi:"recoveryBootDiagStorageAccountId"`
	RecoveryCloudServiceId            *string                        `pulumi:"recoveryCloudServiceId"`
	RecoveryContainerId               *string                        `pulumi:"recoveryContainerId"`
	RecoveryProximityPlacementGroupId *string                        `pulumi:"recoveryProximityPlacementGroupId"`
	RecoveryResourceGroupId           *string                        `pulumi:"recoveryResourceGroupId"`
	RecoverySubnetName                *string                        `pulumi:"recoverySubnetName"`
	VmDisks                           []A2AVmDiskInputDetails        `pulumi:"vmDisks"`
	VmManagedDisks                    []A2AVmManagedDiskInputDetails `pulumi:"vmManagedDisks"`
}

type A2APolicyCreationInput struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      *string `pulumi:"instanceType"`
	MultiVmSyncStatus                 string  `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
}

type A2APolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string  `pulumi:"instanceType"`
	MultiVmSyncStatus                 *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int    `pulumi:"recoveryPointThresholdInMinutes"`
}

type A2AProtectedDiskDetailsResponse struct {
	AllowedDiskLevelOperation              []string `pulumi:"allowedDiskLevelOperation"`
	DataPendingAtSourceAgentInMB           *float64 `pulumi:"dataPendingAtSourceAgentInMB"`
	DataPendingInStagingStorageAccountInMB *float64 `pulumi:"dataPendingInStagingStorageAccountInMB"`
	DekKeyVaultArmId                       *string  `pulumi:"dekKeyVaultArmId"`
	DiskCapacityInBytes                    *float64 `pulumi:"diskCapacityInBytes"`
	DiskName                               *string  `pulumi:"diskName"`
	DiskState                              *string  `pulumi:"diskState"`
	DiskType                               *string  `pulumi:"diskType"`
	DiskUri                                *string  `pulumi:"diskUri"`
	IsDiskEncrypted                        *bool    `pulumi:"isDiskEncrypted"`
	IsDiskKeyEncrypted                     *bool    `pulumi:"isDiskKeyEncrypted"`
	KekKeyVaultArmId                       *string  `pulumi:"kekKeyVaultArmId"`
	KeyIdentifier                          *string  `pulumi:"keyIdentifier"`
	MonitoringJobType                      *string  `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion         *int     `pulumi:"monitoringPercentageCompletion"`
	PrimaryDiskAzureStorageAccountId       *string  `pulumi:"primaryDiskAzureStorageAccountId"`
	PrimaryStagingAzureStorageAccountId    *string  `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryAzureStorageAccountId          *string  `pulumi:"recoveryAzureStorageAccountId"`
	RecoveryDiskUri                        *string  `pulumi:"recoveryDiskUri"`
	ResyncRequired                         *bool    `pulumi:"resyncRequired"`
	SecretIdentifier                       *string  `pulumi:"secretIdentifier"`
}

type A2AProtectedManagedDiskDetailsResponse struct {
	AllowedDiskLevelOperation              []string `pulumi:"allowedDiskLevelOperation"`
	DataPendingAtSourceAgentInMB           *float64 `pulumi:"dataPendingAtSourceAgentInMB"`
	DataPendingInStagingStorageAccountInMB *float64 `pulumi:"dataPendingInStagingStorageAccountInMB"`
	DekKeyVaultArmId                       *string  `pulumi:"dekKeyVaultArmId"`
	DiskCapacityInBytes                    *float64 `pulumi:"diskCapacityInBytes"`
	DiskId                                 *string  `pulumi:"diskId"`
	DiskName                               *string  `pulumi:"diskName"`
	DiskState                              *string  `pulumi:"diskState"`
	DiskType                               *string  `pulumi:"diskType"`
	FailoverDiskName                       *string  `pulumi:"failoverDiskName"`
	IsDiskEncrypted                        *bool    `pulumi:"isDiskEncrypted"`
	IsDiskKeyEncrypted                     *bool    `pulumi:"isDiskKeyEncrypted"`
	KekKeyVaultArmId                       *string  `pulumi:"kekKeyVaultArmId"`
	KeyIdentifier                          *string  `pulumi:"keyIdentifier"`
	MonitoringJobType                      *string  `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion         *int     `pulumi:"monitoringPercentageCompletion"`
	PrimaryStagingAzureStorageAccountId    *string  `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryDiskEncryptionSetId            *string  `pulumi:"recoveryDiskEncryptionSetId"`
	RecoveryReplicaDiskAccountType         *string  `pulumi:"recoveryReplicaDiskAccountType"`
	RecoveryReplicaDiskId                  *string  `pulumi:"recoveryReplicaDiskId"`
	RecoveryResourceGroupId                *string  `pulumi:"recoveryResourceGroupId"`
	RecoveryTargetDiskAccountType          *string  `pulumi:"recoveryTargetDiskAccountType"`
	RecoveryTargetDiskId                   *string  `pulumi:"recoveryTargetDiskId"`
	ResyncRequired                         *bool    `pulumi:"resyncRequired"`
	SecretIdentifier                       *string  `pulumi:"secretIdentifier"`
	TfoDiskName                            *string  `pulumi:"tfoDiskName"`
}

type A2AProtectionContainerMappingDetailsResponse struct {
	AgentAutoUpdateStatus  *string `pulumi:"agentAutoUpdateStatus"`
	AutomationAccountArmId *string `pulumi:"automationAccountArmId"`
	InstanceType           string  `pulumi:"instanceType"`
	JobScheduleName        *string `pulumi:"jobScheduleName"`
	ScheduleName           *string `pulumi:"scheduleName"`
}

type A2AReplicationDetailsResponse struct {
	AgentVersion                       *string                                    `pulumi:"agentVersion"`
	FabricObjectId                     *string                                    `pulumi:"fabricObjectId"`
	InitialPrimaryFabricLocation       string                                     `pulumi:"initialPrimaryFabricLocation"`
	InitialPrimaryZone                 string                                     `pulumi:"initialPrimaryZone"`
	InitialRecoveryFabricLocation      string                                     `pulumi:"initialRecoveryFabricLocation"`
	InitialRecoveryZone                string                                     `pulumi:"initialRecoveryZone"`
	InstanceType                       string                                     `pulumi:"instanceType"`
	IsReplicationAgentUpdateRequired   *bool                                      `pulumi:"isReplicationAgentUpdateRequired"`
	LastHeartbeat                      *string                                    `pulumi:"lastHeartbeat"`
	LastRpoCalculatedTime              *string                                    `pulumi:"lastRpoCalculatedTime"`
	LifecycleId                        *string                                    `pulumi:"lifecycleId"`
	ManagementId                       *string                                    `pulumi:"managementId"`
	MonitoringJobType                  *string                                    `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion     *int                                       `pulumi:"monitoringPercentageCompletion"`
	MultiVmGroupCreateOption           *string                                    `pulumi:"multiVmGroupCreateOption"`
	MultiVmGroupId                     *string                                    `pulumi:"multiVmGroupId"`
	MultiVmGroupName                   *string                                    `pulumi:"multiVmGroupName"`
	OsType                             *string                                    `pulumi:"osType"`
	PrimaryFabricLocation              *string                                    `pulumi:"primaryFabricLocation"`
	ProtectedDisks                     []A2AProtectedDiskDetailsResponse          `pulumi:"protectedDisks"`
	ProtectedManagedDisks              []A2AProtectedManagedDiskDetailsResponse   `pulumi:"protectedManagedDisks"`
	RecoveryAvailabilitySet            *string                                    `pulumi:"recoveryAvailabilitySet"`
	RecoveryAvailabilityZone           *string                                    `pulumi:"recoveryAvailabilityZone"`
	RecoveryAzureResourceGroupId       *string                                    `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureVMName                *string                                    `pulumi:"recoveryAzureVMName"`
	RecoveryAzureVMSize                *string                                    `pulumi:"recoveryAzureVMSize"`
	RecoveryBootDiagStorageAccountId   *string                                    `pulumi:"recoveryBootDiagStorageAccountId"`
	RecoveryCloudService               *string                                    `pulumi:"recoveryCloudService"`
	RecoveryFabricLocation             *string                                    `pulumi:"recoveryFabricLocation"`
	RecoveryFabricObjectId             *string                                    `pulumi:"recoveryFabricObjectId"`
	RecoveryProximityPlacementGroupId  *string                                    `pulumi:"recoveryProximityPlacementGroupId"`
	RpoInSeconds                       *float64                                   `pulumi:"rpoInSeconds"`
	SelectedRecoveryAzureNetworkId     *string                                    `pulumi:"selectedRecoveryAzureNetworkId"`
	SelectedTfoAzureNetworkId          *string                                    `pulumi:"selectedTfoAzureNetworkId"`
	TestFailoverRecoveryFabricObjectId *string                                    `pulumi:"testFailoverRecoveryFabricObjectId"`
	TfoAzureVMName                     *string                                    `pulumi:"tfoAzureVMName"`
	UnprotectedDisks                   []A2AUnprotectedDiskDetailsResponse        `pulumi:"unprotectedDisks"`
	VmEncryptionType                   string                                     `pulumi:"vmEncryptionType"`
	VmNics                             []VMNicDetailsResponse                     `pulumi:"vmNics"`
	VmProtectionState                  *string                                    `pulumi:"vmProtectionState"`
	VmProtectionStateDescription       *string                                    `pulumi:"vmProtectionStateDescription"`
	VmSyncedConfigDetails              *AzureToAzureVmSyncedConfigDetailsResponse `pulumi:"vmSyncedConfigDetails"`
}

type A2AUnprotectedDiskDetailsResponse struct {
	DiskLunId *int `pulumi:"diskLunId"`
}

type A2AVmDiskInputDetails struct {
	DiskUri                             *string `pulumi:"diskUri"`
	PrimaryStagingAzureStorageAccountId *string `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryAzureStorageAccountId       *string `pulumi:"recoveryAzureStorageAccountId"`
}

type A2AVmManagedDiskInputDetails struct {
	DiskEncryptionInfo                  *DiskEncryptionInfo `pulumi:"diskEncryptionInfo"`
	DiskId                              *string             `pulumi:"diskId"`
	PrimaryStagingAzureStorageAccountId *string             `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryDiskEncryptionSetId         *string             `pulumi:"recoveryDiskEncryptionSetId"`
	RecoveryReplicaDiskAccountType      *string             `pulumi:"recoveryReplicaDiskAccountType"`
	RecoveryResourceGroupId             *string             `pulumi:"recoveryResourceGroupId"`
	RecoveryTargetDiskAccountType       *string             `pulumi:"recoveryTargetDiskAccountType"`
}

type AADProperties struct {
	Audience                 *string `pulumi:"audience"`
	Authority                *string `pulumi:"authority"`
	ServicePrincipalClientId *string `pulumi:"servicePrincipalClientId"`
	ServicePrincipalObjectId *string `pulumi:"servicePrincipalObjectId"`
	TenantId                 *string `pulumi:"tenantId"`
}

type AddRecoveryServicesProviderInputProperties struct {
	AuthenticationIdentityInput          IdentityProviderInput  `pulumi:"authenticationIdentityInput"`
	DataPlaneAuthenticationIdentityInput *IdentityProviderInput `pulumi:"dataPlaneAuthenticationIdentityInput"`
	MachineId                            *string                `pulumi:"machineId"`
	MachineName                          string                 `pulumi:"machineName"`
	ResourceAccessIdentityInput          IdentityProviderInput  `pulumi:"resourceAccessIdentityInput"`
}





type AddRecoveryServicesProviderInputPropertiesInput interface {
	pulumi.Input

	ToAddRecoveryServicesProviderInputPropertiesOutput() AddRecoveryServicesProviderInputPropertiesOutput
	ToAddRecoveryServicesProviderInputPropertiesOutputWithContext(context.Context) AddRecoveryServicesProviderInputPropertiesOutput
}

type AddRecoveryServicesProviderInputPropertiesArgs struct {
	AuthenticationIdentityInput          IdentityProviderInputInput    `pulumi:"authenticationIdentityInput"`
	DataPlaneAuthenticationIdentityInput IdentityProviderInputPtrInput `pulumi:"dataPlaneAuthenticationIdentityInput"`
	MachineId                            pulumi.StringPtrInput         `pulumi:"machineId"`
	MachineName                          pulumi.StringInput            `pulumi:"machineName"`
	ResourceAccessIdentityInput          IdentityProviderInputInput    `pulumi:"resourceAccessIdentityInput"`
}

func (AddRecoveryServicesProviderInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddRecoveryServicesProviderInputProperties)(nil)).Elem()
}

func (i AddRecoveryServicesProviderInputPropertiesArgs) ToAddRecoveryServicesProviderInputPropertiesOutput() AddRecoveryServicesProviderInputPropertiesOutput {
	return i.ToAddRecoveryServicesProviderInputPropertiesOutputWithContext(context.Background())
}

func (i AddRecoveryServicesProviderInputPropertiesArgs) ToAddRecoveryServicesProviderInputPropertiesOutputWithContext(ctx context.Context) AddRecoveryServicesProviderInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddRecoveryServicesProviderInputPropertiesOutput)
}

type AddRecoveryServicesProviderInputPropertiesOutput struct{ *pulumi.OutputState }

func (AddRecoveryServicesProviderInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddRecoveryServicesProviderInputProperties)(nil)).Elem()
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) ToAddRecoveryServicesProviderInputPropertiesOutput() AddRecoveryServicesProviderInputPropertiesOutput {
	return o
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) ToAddRecoveryServicesProviderInputPropertiesOutputWithContext(ctx context.Context) AddRecoveryServicesProviderInputPropertiesOutput {
	return o
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) AuthenticationIdentityInput() IdentityProviderInputOutput {
	return o.ApplyT(func(v AddRecoveryServicesProviderInputProperties) IdentityProviderInput {
		return v.AuthenticationIdentityInput
	}).(IdentityProviderInputOutput)
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) DataPlaneAuthenticationIdentityInput() IdentityProviderInputPtrOutput {
	return o.ApplyT(func(v AddRecoveryServicesProviderInputProperties) *IdentityProviderInput {
		return v.DataPlaneAuthenticationIdentityInput
	}).(IdentityProviderInputPtrOutput)
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) MachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddRecoveryServicesProviderInputProperties) *string { return v.MachineId }).(pulumi.StringPtrOutput)
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v AddRecoveryServicesProviderInputProperties) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) ResourceAccessIdentityInput() IdentityProviderInputOutput {
	return o.ApplyT(func(v AddRecoveryServicesProviderInputProperties) IdentityProviderInput {
		return v.ResourceAccessIdentityInput
	}).(IdentityProviderInputOutput)
}

type AddVCenterRequestProperties struct {
	FriendlyName    *string `pulumi:"friendlyName"`
	IpAddress       *string `pulumi:"ipAddress"`
	Port            *string `pulumi:"port"`
	ProcessServerId *string `pulumi:"processServerId"`
	RunAsAccountId  *string `pulumi:"runAsAccountId"`
}





type AddVCenterRequestPropertiesInput interface {
	pulumi.Input

	ToAddVCenterRequestPropertiesOutput() AddVCenterRequestPropertiesOutput
	ToAddVCenterRequestPropertiesOutputWithContext(context.Context) AddVCenterRequestPropertiesOutput
}

type AddVCenterRequestPropertiesArgs struct {
	FriendlyName    pulumi.StringPtrInput `pulumi:"friendlyName"`
	IpAddress       pulumi.StringPtrInput `pulumi:"ipAddress"`
	Port            pulumi.StringPtrInput `pulumi:"port"`
	ProcessServerId pulumi.StringPtrInput `pulumi:"processServerId"`
	RunAsAccountId  pulumi.StringPtrInput `pulumi:"runAsAccountId"`
}

func (AddVCenterRequestPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddVCenterRequestProperties)(nil)).Elem()
}

func (i AddVCenterRequestPropertiesArgs) ToAddVCenterRequestPropertiesOutput() AddVCenterRequestPropertiesOutput {
	return i.ToAddVCenterRequestPropertiesOutputWithContext(context.Background())
}

func (i AddVCenterRequestPropertiesArgs) ToAddVCenterRequestPropertiesOutputWithContext(ctx context.Context) AddVCenterRequestPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddVCenterRequestPropertiesOutput)
}

func (i AddVCenterRequestPropertiesArgs) ToAddVCenterRequestPropertiesPtrOutput() AddVCenterRequestPropertiesPtrOutput {
	return i.ToAddVCenterRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i AddVCenterRequestPropertiesArgs) ToAddVCenterRequestPropertiesPtrOutputWithContext(ctx context.Context) AddVCenterRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddVCenterRequestPropertiesOutput).ToAddVCenterRequestPropertiesPtrOutputWithContext(ctx)
}









type AddVCenterRequestPropertiesPtrInput interface {
	pulumi.Input

	ToAddVCenterRequestPropertiesPtrOutput() AddVCenterRequestPropertiesPtrOutput
	ToAddVCenterRequestPropertiesPtrOutputWithContext(context.Context) AddVCenterRequestPropertiesPtrOutput
}

type addVCenterRequestPropertiesPtrType AddVCenterRequestPropertiesArgs

func AddVCenterRequestPropertiesPtr(v *AddVCenterRequestPropertiesArgs) AddVCenterRequestPropertiesPtrInput {
	return (*addVCenterRequestPropertiesPtrType)(v)
}

func (*addVCenterRequestPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddVCenterRequestProperties)(nil)).Elem()
}

func (i *addVCenterRequestPropertiesPtrType) ToAddVCenterRequestPropertiesPtrOutput() AddVCenterRequestPropertiesPtrOutput {
	return i.ToAddVCenterRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i *addVCenterRequestPropertiesPtrType) ToAddVCenterRequestPropertiesPtrOutputWithContext(ctx context.Context) AddVCenterRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddVCenterRequestPropertiesPtrOutput)
}

type AddVCenterRequestPropertiesOutput struct{ *pulumi.OutputState }

func (AddVCenterRequestPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddVCenterRequestProperties)(nil)).Elem()
}

func (o AddVCenterRequestPropertiesOutput) ToAddVCenterRequestPropertiesOutput() AddVCenterRequestPropertiesOutput {
	return o
}

func (o AddVCenterRequestPropertiesOutput) ToAddVCenterRequestPropertiesOutputWithContext(ctx context.Context) AddVCenterRequestPropertiesOutput {
	return o
}

func (o AddVCenterRequestPropertiesOutput) ToAddVCenterRequestPropertiesPtrOutput() AddVCenterRequestPropertiesPtrOutput {
	return o.ToAddVCenterRequestPropertiesPtrOutputWithContext(context.Background())
}

func (o AddVCenterRequestPropertiesOutput) ToAddVCenterRequestPropertiesPtrOutputWithContext(ctx context.Context) AddVCenterRequestPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddVCenterRequestProperties) *AddVCenterRequestProperties {
		return &v
	}).(AddVCenterRequestPropertiesPtrOutput)
}

func (o AddVCenterRequestPropertiesOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddVCenterRequestProperties) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AddVCenterRequestPropertiesOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddVCenterRequestProperties) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o AddVCenterRequestPropertiesOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddVCenterRequestProperties) *string { return v.Port }).(pulumi.StringPtrOutput)
}

func (o AddVCenterRequestPropertiesOutput) ProcessServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddVCenterRequestProperties) *string { return v.ProcessServerId }).(pulumi.StringPtrOutput)
}

func (o AddVCenterRequestPropertiesOutput) RunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddVCenterRequestProperties) *string { return v.RunAsAccountId }).(pulumi.StringPtrOutput)
}

type AddVCenterRequestPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AddVCenterRequestPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddVCenterRequestProperties)(nil)).Elem()
}

func (o AddVCenterRequestPropertiesPtrOutput) ToAddVCenterRequestPropertiesPtrOutput() AddVCenterRequestPropertiesPtrOutput {
	return o
}

func (o AddVCenterRequestPropertiesPtrOutput) ToAddVCenterRequestPropertiesPtrOutputWithContext(ctx context.Context) AddVCenterRequestPropertiesPtrOutput {
	return o
}

func (o AddVCenterRequestPropertiesPtrOutput) Elem() AddVCenterRequestPropertiesOutput {
	return o.ApplyT(func(v *AddVCenterRequestProperties) AddVCenterRequestProperties {
		if v != nil {
			return *v
		}
		var ret AddVCenterRequestProperties
		return ret
	}).(AddVCenterRequestPropertiesOutput)
}

func (o AddVCenterRequestPropertiesPtrOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddVCenterRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.FriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o AddVCenterRequestPropertiesPtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddVCenterRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o AddVCenterRequestPropertiesPtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddVCenterRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.StringPtrOutput)
}

func (o AddVCenterRequestPropertiesPtrOutput) ProcessServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddVCenterRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProcessServerId
	}).(pulumi.StringPtrOutput)
}

func (o AddVCenterRequestPropertiesPtrOutput) RunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddVCenterRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.RunAsAccountId
	}).(pulumi.StringPtrOutput)
}

type AgentDetailsResponse struct {
	AgentId   string                     `pulumi:"agentId"`
	BiosId    string                     `pulumi:"biosId"`
	Disks     []AgentDiskDetailsResponse `pulumi:"disks"`
	Fqdn      string                     `pulumi:"fqdn"`
	MachineId string                     `pulumi:"machineId"`
}

type AgentDiskDetailsResponse struct {
	CapacityInBytes float64 `pulumi:"capacityInBytes"`
	DiskId          string  `pulumi:"diskId"`
	DiskName        string  `pulumi:"diskName"`
	IsOSDisk        string  `pulumi:"isOSDisk"`
	LunId           int     `pulumi:"lunId"`
}

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

type AzureFabricCreationInput struct {
	InstanceType *string `pulumi:"instanceType"`
	Location     *string `pulumi:"location"`
}

type AzureFabricSpecificDetailsResponse struct {
	ContainerIds []string `pulumi:"containerIds"`
	InstanceType string   `pulumi:"instanceType"`
	Location     *string  `pulumi:"location"`
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

type AzureToAzureCreateNetworkMappingInput struct {
	InstanceType     *string `pulumi:"instanceType"`
	PrimaryNetworkId *string `pulumi:"primaryNetworkId"`
}

type AzureToAzureNetworkMappingSettingsResponse struct {
	InstanceType           string  `pulumi:"instanceType"`
	PrimaryFabricLocation  *string `pulumi:"primaryFabricLocation"`
	RecoveryFabricLocation *string `pulumi:"recoveryFabricLocation"`
}

type AzureToAzureVmSyncedConfigDetailsResponse struct {
	InputEndpoints []InputEndpointResponse `pulumi:"inputEndpoints"`
	Tags           map[string]string       `pulumi:"tags"`
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

type AzureVmDiskDetailsResponse struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskId              *string `pulumi:"diskId"`
	LunId               *string `pulumi:"lunId"`
	MaxSizeMB           *string `pulumi:"maxSizeMB"`
	TargetDiskLocation  *string `pulumi:"targetDiskLocation"`
	TargetDiskName      *string `pulumi:"targetDiskName"`
	VhdId               *string `pulumi:"vhdId"`
	VhdName             *string `pulumi:"vhdName"`
	VhdType             *string `pulumi:"vhdType"`
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

type CreateNetworkMappingInputProperties struct {
	FabricSpecificDetails interface{} `pulumi:"fabricSpecificDetails"`
	RecoveryFabricName    *string     `pulumi:"recoveryFabricName"`
	RecoveryNetworkId     *string     `pulumi:"recoveryNetworkId"`
}





type CreateNetworkMappingInputPropertiesInput interface {
	pulumi.Input

	ToCreateNetworkMappingInputPropertiesOutput() CreateNetworkMappingInputPropertiesOutput
	ToCreateNetworkMappingInputPropertiesOutputWithContext(context.Context) CreateNetworkMappingInputPropertiesOutput
}

type CreateNetworkMappingInputPropertiesArgs struct {
	FabricSpecificDetails pulumi.Input          `pulumi:"fabricSpecificDetails"`
	RecoveryFabricName    pulumi.StringPtrInput `pulumi:"recoveryFabricName"`
	RecoveryNetworkId     pulumi.StringPtrInput `pulumi:"recoveryNetworkId"`
}

func (CreateNetworkMappingInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateNetworkMappingInputProperties)(nil)).Elem()
}

func (i CreateNetworkMappingInputPropertiesArgs) ToCreateNetworkMappingInputPropertiesOutput() CreateNetworkMappingInputPropertiesOutput {
	return i.ToCreateNetworkMappingInputPropertiesOutputWithContext(context.Background())
}

func (i CreateNetworkMappingInputPropertiesArgs) ToCreateNetworkMappingInputPropertiesOutputWithContext(ctx context.Context) CreateNetworkMappingInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateNetworkMappingInputPropertiesOutput)
}

func (i CreateNetworkMappingInputPropertiesArgs) ToCreateNetworkMappingInputPropertiesPtrOutput() CreateNetworkMappingInputPropertiesPtrOutput {
	return i.ToCreateNetworkMappingInputPropertiesPtrOutputWithContext(context.Background())
}

func (i CreateNetworkMappingInputPropertiesArgs) ToCreateNetworkMappingInputPropertiesPtrOutputWithContext(ctx context.Context) CreateNetworkMappingInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateNetworkMappingInputPropertiesOutput).ToCreateNetworkMappingInputPropertiesPtrOutputWithContext(ctx)
}









type CreateNetworkMappingInputPropertiesPtrInput interface {
	pulumi.Input

	ToCreateNetworkMappingInputPropertiesPtrOutput() CreateNetworkMappingInputPropertiesPtrOutput
	ToCreateNetworkMappingInputPropertiesPtrOutputWithContext(context.Context) CreateNetworkMappingInputPropertiesPtrOutput
}

type createNetworkMappingInputPropertiesPtrType CreateNetworkMappingInputPropertiesArgs

func CreateNetworkMappingInputPropertiesPtr(v *CreateNetworkMappingInputPropertiesArgs) CreateNetworkMappingInputPropertiesPtrInput {
	return (*createNetworkMappingInputPropertiesPtrType)(v)
}

func (*createNetworkMappingInputPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateNetworkMappingInputProperties)(nil)).Elem()
}

func (i *createNetworkMappingInputPropertiesPtrType) ToCreateNetworkMappingInputPropertiesPtrOutput() CreateNetworkMappingInputPropertiesPtrOutput {
	return i.ToCreateNetworkMappingInputPropertiesPtrOutputWithContext(context.Background())
}

func (i *createNetworkMappingInputPropertiesPtrType) ToCreateNetworkMappingInputPropertiesPtrOutputWithContext(ctx context.Context) CreateNetworkMappingInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateNetworkMappingInputPropertiesPtrOutput)
}

type CreateNetworkMappingInputPropertiesOutput struct{ *pulumi.OutputState }

func (CreateNetworkMappingInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateNetworkMappingInputProperties)(nil)).Elem()
}

func (o CreateNetworkMappingInputPropertiesOutput) ToCreateNetworkMappingInputPropertiesOutput() CreateNetworkMappingInputPropertiesOutput {
	return o
}

func (o CreateNetworkMappingInputPropertiesOutput) ToCreateNetworkMappingInputPropertiesOutputWithContext(ctx context.Context) CreateNetworkMappingInputPropertiesOutput {
	return o
}

func (o CreateNetworkMappingInputPropertiesOutput) ToCreateNetworkMappingInputPropertiesPtrOutput() CreateNetworkMappingInputPropertiesPtrOutput {
	return o.ToCreateNetworkMappingInputPropertiesPtrOutputWithContext(context.Background())
}

func (o CreateNetworkMappingInputPropertiesOutput) ToCreateNetworkMappingInputPropertiesPtrOutputWithContext(ctx context.Context) CreateNetworkMappingInputPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateNetworkMappingInputProperties) *CreateNetworkMappingInputProperties {
		return &v
	}).(CreateNetworkMappingInputPropertiesPtrOutput)
}

func (o CreateNetworkMappingInputPropertiesOutput) FabricSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v CreateNetworkMappingInputProperties) interface{} { return v.FabricSpecificDetails }).(pulumi.AnyOutput)
}

func (o CreateNetworkMappingInputPropertiesOutput) RecoveryFabricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateNetworkMappingInputProperties) *string { return v.RecoveryFabricName }).(pulumi.StringPtrOutput)
}

func (o CreateNetworkMappingInputPropertiesOutput) RecoveryNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateNetworkMappingInputProperties) *string { return v.RecoveryNetworkId }).(pulumi.StringPtrOutput)
}

type CreateNetworkMappingInputPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CreateNetworkMappingInputPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateNetworkMappingInputProperties)(nil)).Elem()
}

func (o CreateNetworkMappingInputPropertiesPtrOutput) ToCreateNetworkMappingInputPropertiesPtrOutput() CreateNetworkMappingInputPropertiesPtrOutput {
	return o
}

func (o CreateNetworkMappingInputPropertiesPtrOutput) ToCreateNetworkMappingInputPropertiesPtrOutputWithContext(ctx context.Context) CreateNetworkMappingInputPropertiesPtrOutput {
	return o
}

func (o CreateNetworkMappingInputPropertiesPtrOutput) Elem() CreateNetworkMappingInputPropertiesOutput {
	return o.ApplyT(func(v *CreateNetworkMappingInputProperties) CreateNetworkMappingInputProperties {
		if v != nil {
			return *v
		}
		var ret CreateNetworkMappingInputProperties
		return ret
	}).(CreateNetworkMappingInputPropertiesOutput)
}

func (o CreateNetworkMappingInputPropertiesPtrOutput) FabricSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *CreateNetworkMappingInputProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.FabricSpecificDetails
	}).(pulumi.AnyOutput)
}

func (o CreateNetworkMappingInputPropertiesPtrOutput) RecoveryFabricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateNetworkMappingInputProperties) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryFabricName
	}).(pulumi.StringPtrOutput)
}

func (o CreateNetworkMappingInputPropertiesPtrOutput) RecoveryNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateNetworkMappingInputProperties) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryNetworkId
	}).(pulumi.StringPtrOutput)
}

type CreatePolicyInputProperties struct {
	ProviderSpecificInput interface{} `pulumi:"providerSpecificInput"`
}





type CreatePolicyInputPropertiesInput interface {
	pulumi.Input

	ToCreatePolicyInputPropertiesOutput() CreatePolicyInputPropertiesOutput
	ToCreatePolicyInputPropertiesOutputWithContext(context.Context) CreatePolicyInputPropertiesOutput
}

type CreatePolicyInputPropertiesArgs struct {
	ProviderSpecificInput pulumi.Input `pulumi:"providerSpecificInput"`
}

func (CreatePolicyInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatePolicyInputProperties)(nil)).Elem()
}

func (i CreatePolicyInputPropertiesArgs) ToCreatePolicyInputPropertiesOutput() CreatePolicyInputPropertiesOutput {
	return i.ToCreatePolicyInputPropertiesOutputWithContext(context.Background())
}

func (i CreatePolicyInputPropertiesArgs) ToCreatePolicyInputPropertiesOutputWithContext(ctx context.Context) CreatePolicyInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatePolicyInputPropertiesOutput)
}

func (i CreatePolicyInputPropertiesArgs) ToCreatePolicyInputPropertiesPtrOutput() CreatePolicyInputPropertiesPtrOutput {
	return i.ToCreatePolicyInputPropertiesPtrOutputWithContext(context.Background())
}

func (i CreatePolicyInputPropertiesArgs) ToCreatePolicyInputPropertiesPtrOutputWithContext(ctx context.Context) CreatePolicyInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatePolicyInputPropertiesOutput).ToCreatePolicyInputPropertiesPtrOutputWithContext(ctx)
}









type CreatePolicyInputPropertiesPtrInput interface {
	pulumi.Input

	ToCreatePolicyInputPropertiesPtrOutput() CreatePolicyInputPropertiesPtrOutput
	ToCreatePolicyInputPropertiesPtrOutputWithContext(context.Context) CreatePolicyInputPropertiesPtrOutput
}

type createPolicyInputPropertiesPtrType CreatePolicyInputPropertiesArgs

func CreatePolicyInputPropertiesPtr(v *CreatePolicyInputPropertiesArgs) CreatePolicyInputPropertiesPtrInput {
	return (*createPolicyInputPropertiesPtrType)(v)
}

func (*createPolicyInputPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreatePolicyInputProperties)(nil)).Elem()
}

func (i *createPolicyInputPropertiesPtrType) ToCreatePolicyInputPropertiesPtrOutput() CreatePolicyInputPropertiesPtrOutput {
	return i.ToCreatePolicyInputPropertiesPtrOutputWithContext(context.Background())
}

func (i *createPolicyInputPropertiesPtrType) ToCreatePolicyInputPropertiesPtrOutputWithContext(ctx context.Context) CreatePolicyInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreatePolicyInputPropertiesPtrOutput)
}

type CreatePolicyInputPropertiesOutput struct{ *pulumi.OutputState }

func (CreatePolicyInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatePolicyInputProperties)(nil)).Elem()
}

func (o CreatePolicyInputPropertiesOutput) ToCreatePolicyInputPropertiesOutput() CreatePolicyInputPropertiesOutput {
	return o
}

func (o CreatePolicyInputPropertiesOutput) ToCreatePolicyInputPropertiesOutputWithContext(ctx context.Context) CreatePolicyInputPropertiesOutput {
	return o
}

func (o CreatePolicyInputPropertiesOutput) ToCreatePolicyInputPropertiesPtrOutput() CreatePolicyInputPropertiesPtrOutput {
	return o.ToCreatePolicyInputPropertiesPtrOutputWithContext(context.Background())
}

func (o CreatePolicyInputPropertiesOutput) ToCreatePolicyInputPropertiesPtrOutputWithContext(ctx context.Context) CreatePolicyInputPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreatePolicyInputProperties) *CreatePolicyInputProperties {
		return &v
	}).(CreatePolicyInputPropertiesPtrOutput)
}

func (o CreatePolicyInputPropertiesOutput) ProviderSpecificInput() pulumi.AnyOutput {
	return o.ApplyT(func(v CreatePolicyInputProperties) interface{} { return v.ProviderSpecificInput }).(pulumi.AnyOutput)
}

type CreatePolicyInputPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CreatePolicyInputPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreatePolicyInputProperties)(nil)).Elem()
}

func (o CreatePolicyInputPropertiesPtrOutput) ToCreatePolicyInputPropertiesPtrOutput() CreatePolicyInputPropertiesPtrOutput {
	return o
}

func (o CreatePolicyInputPropertiesPtrOutput) ToCreatePolicyInputPropertiesPtrOutputWithContext(ctx context.Context) CreatePolicyInputPropertiesPtrOutput {
	return o
}

func (o CreatePolicyInputPropertiesPtrOutput) Elem() CreatePolicyInputPropertiesOutput {
	return o.ApplyT(func(v *CreatePolicyInputProperties) CreatePolicyInputProperties {
		if v != nil {
			return *v
		}
		var ret CreatePolicyInputProperties
		return ret
	}).(CreatePolicyInputPropertiesOutput)
}

func (o CreatePolicyInputPropertiesPtrOutput) ProviderSpecificInput() pulumi.AnyOutput {
	return o.ApplyT(func(v *CreatePolicyInputProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.ProviderSpecificInput
	}).(pulumi.AnyOutput)
}

type CreateProtectionContainerMappingInputProperties struct {
	PolicyId                    *string     `pulumi:"policyId"`
	ProviderSpecificInput       interface{} `pulumi:"providerSpecificInput"`
	TargetProtectionContainerId *string     `pulumi:"targetProtectionContainerId"`
}





type CreateProtectionContainerMappingInputPropertiesInput interface {
	pulumi.Input

	ToCreateProtectionContainerMappingInputPropertiesOutput() CreateProtectionContainerMappingInputPropertiesOutput
	ToCreateProtectionContainerMappingInputPropertiesOutputWithContext(context.Context) CreateProtectionContainerMappingInputPropertiesOutput
}

type CreateProtectionContainerMappingInputPropertiesArgs struct {
	PolicyId                    pulumi.StringPtrInput `pulumi:"policyId"`
	ProviderSpecificInput       pulumi.Input          `pulumi:"providerSpecificInput"`
	TargetProtectionContainerId pulumi.StringPtrInput `pulumi:"targetProtectionContainerId"`
}

func (CreateProtectionContainerMappingInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateProtectionContainerMappingInputProperties)(nil)).Elem()
}

func (i CreateProtectionContainerMappingInputPropertiesArgs) ToCreateProtectionContainerMappingInputPropertiesOutput() CreateProtectionContainerMappingInputPropertiesOutput {
	return i.ToCreateProtectionContainerMappingInputPropertiesOutputWithContext(context.Background())
}

func (i CreateProtectionContainerMappingInputPropertiesArgs) ToCreateProtectionContainerMappingInputPropertiesOutputWithContext(ctx context.Context) CreateProtectionContainerMappingInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateProtectionContainerMappingInputPropertiesOutput)
}

func (i CreateProtectionContainerMappingInputPropertiesArgs) ToCreateProtectionContainerMappingInputPropertiesPtrOutput() CreateProtectionContainerMappingInputPropertiesPtrOutput {
	return i.ToCreateProtectionContainerMappingInputPropertiesPtrOutputWithContext(context.Background())
}

func (i CreateProtectionContainerMappingInputPropertiesArgs) ToCreateProtectionContainerMappingInputPropertiesPtrOutputWithContext(ctx context.Context) CreateProtectionContainerMappingInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateProtectionContainerMappingInputPropertiesOutput).ToCreateProtectionContainerMappingInputPropertiesPtrOutputWithContext(ctx)
}









type CreateProtectionContainerMappingInputPropertiesPtrInput interface {
	pulumi.Input

	ToCreateProtectionContainerMappingInputPropertiesPtrOutput() CreateProtectionContainerMappingInputPropertiesPtrOutput
	ToCreateProtectionContainerMappingInputPropertiesPtrOutputWithContext(context.Context) CreateProtectionContainerMappingInputPropertiesPtrOutput
}

type createProtectionContainerMappingInputPropertiesPtrType CreateProtectionContainerMappingInputPropertiesArgs

func CreateProtectionContainerMappingInputPropertiesPtr(v *CreateProtectionContainerMappingInputPropertiesArgs) CreateProtectionContainerMappingInputPropertiesPtrInput {
	return (*createProtectionContainerMappingInputPropertiesPtrType)(v)
}

func (*createProtectionContainerMappingInputPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateProtectionContainerMappingInputProperties)(nil)).Elem()
}

func (i *createProtectionContainerMappingInputPropertiesPtrType) ToCreateProtectionContainerMappingInputPropertiesPtrOutput() CreateProtectionContainerMappingInputPropertiesPtrOutput {
	return i.ToCreateProtectionContainerMappingInputPropertiesPtrOutputWithContext(context.Background())
}

func (i *createProtectionContainerMappingInputPropertiesPtrType) ToCreateProtectionContainerMappingInputPropertiesPtrOutputWithContext(ctx context.Context) CreateProtectionContainerMappingInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateProtectionContainerMappingInputPropertiesPtrOutput)
}

type CreateProtectionContainerMappingInputPropertiesOutput struct{ *pulumi.OutputState }

func (CreateProtectionContainerMappingInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateProtectionContainerMappingInputProperties)(nil)).Elem()
}

func (o CreateProtectionContainerMappingInputPropertiesOutput) ToCreateProtectionContainerMappingInputPropertiesOutput() CreateProtectionContainerMappingInputPropertiesOutput {
	return o
}

func (o CreateProtectionContainerMappingInputPropertiesOutput) ToCreateProtectionContainerMappingInputPropertiesOutputWithContext(ctx context.Context) CreateProtectionContainerMappingInputPropertiesOutput {
	return o
}

func (o CreateProtectionContainerMappingInputPropertiesOutput) ToCreateProtectionContainerMappingInputPropertiesPtrOutput() CreateProtectionContainerMappingInputPropertiesPtrOutput {
	return o.ToCreateProtectionContainerMappingInputPropertiesPtrOutputWithContext(context.Background())
}

func (o CreateProtectionContainerMappingInputPropertiesOutput) ToCreateProtectionContainerMappingInputPropertiesPtrOutputWithContext(ctx context.Context) CreateProtectionContainerMappingInputPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateProtectionContainerMappingInputProperties) *CreateProtectionContainerMappingInputProperties {
		return &v
	}).(CreateProtectionContainerMappingInputPropertiesPtrOutput)
}

func (o CreateProtectionContainerMappingInputPropertiesOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateProtectionContainerMappingInputProperties) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o CreateProtectionContainerMappingInputPropertiesOutput) ProviderSpecificInput() pulumi.AnyOutput {
	return o.ApplyT(func(v CreateProtectionContainerMappingInputProperties) interface{} { return v.ProviderSpecificInput }).(pulumi.AnyOutput)
}

func (o CreateProtectionContainerMappingInputPropertiesOutput) TargetProtectionContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateProtectionContainerMappingInputProperties) *string { return v.TargetProtectionContainerId }).(pulumi.StringPtrOutput)
}

type CreateProtectionContainerMappingInputPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CreateProtectionContainerMappingInputPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateProtectionContainerMappingInputProperties)(nil)).Elem()
}

func (o CreateProtectionContainerMappingInputPropertiesPtrOutput) ToCreateProtectionContainerMappingInputPropertiesPtrOutput() CreateProtectionContainerMappingInputPropertiesPtrOutput {
	return o
}

func (o CreateProtectionContainerMappingInputPropertiesPtrOutput) ToCreateProtectionContainerMappingInputPropertiesPtrOutputWithContext(ctx context.Context) CreateProtectionContainerMappingInputPropertiesPtrOutput {
	return o
}

func (o CreateProtectionContainerMappingInputPropertiesPtrOutput) Elem() CreateProtectionContainerMappingInputPropertiesOutput {
	return o.ApplyT(func(v *CreateProtectionContainerMappingInputProperties) CreateProtectionContainerMappingInputProperties {
		if v != nil {
			return *v
		}
		var ret CreateProtectionContainerMappingInputProperties
		return ret
	}).(CreateProtectionContainerMappingInputPropertiesOutput)
}

func (o CreateProtectionContainerMappingInputPropertiesPtrOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateProtectionContainerMappingInputProperties) *string {
		if v == nil {
			return nil
		}
		return v.PolicyId
	}).(pulumi.StringPtrOutput)
}

func (o CreateProtectionContainerMappingInputPropertiesPtrOutput) ProviderSpecificInput() pulumi.AnyOutput {
	return o.ApplyT(func(v *CreateProtectionContainerMappingInputProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.ProviderSpecificInput
	}).(pulumi.AnyOutput)
}

func (o CreateProtectionContainerMappingInputPropertiesPtrOutput) TargetProtectionContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateProtectionContainerMappingInputProperties) *string {
		if v == nil {
			return nil
		}
		return v.TargetProtectionContainerId
	}).(pulumi.StringPtrOutput)
}

type CreateRecoveryPlanInputProperties struct {
	FailoverDeploymentModel *string                `pulumi:"failoverDeploymentModel"`
	Groups                  []RecoveryPlanGroup    `pulumi:"groups"`
	PrimaryFabricId         string                 `pulumi:"primaryFabricId"`
	ProviderSpecificInput   []RecoveryPlanA2AInput `pulumi:"providerSpecificInput"`
	RecoveryFabricId        string                 `pulumi:"recoveryFabricId"`
}





type CreateRecoveryPlanInputPropertiesInput interface {
	pulumi.Input

	ToCreateRecoveryPlanInputPropertiesOutput() CreateRecoveryPlanInputPropertiesOutput
	ToCreateRecoveryPlanInputPropertiesOutputWithContext(context.Context) CreateRecoveryPlanInputPropertiesOutput
}

type CreateRecoveryPlanInputPropertiesArgs struct {
	FailoverDeploymentModel pulumi.StringPtrInput          `pulumi:"failoverDeploymentModel"`
	Groups                  RecoveryPlanGroupArrayInput    `pulumi:"groups"`
	PrimaryFabricId         pulumi.StringInput             `pulumi:"primaryFabricId"`
	ProviderSpecificInput   RecoveryPlanA2AInputArrayInput `pulumi:"providerSpecificInput"`
	RecoveryFabricId        pulumi.StringInput             `pulumi:"recoveryFabricId"`
}

func (CreateRecoveryPlanInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateRecoveryPlanInputProperties)(nil)).Elem()
}

func (i CreateRecoveryPlanInputPropertiesArgs) ToCreateRecoveryPlanInputPropertiesOutput() CreateRecoveryPlanInputPropertiesOutput {
	return i.ToCreateRecoveryPlanInputPropertiesOutputWithContext(context.Background())
}

func (i CreateRecoveryPlanInputPropertiesArgs) ToCreateRecoveryPlanInputPropertiesOutputWithContext(ctx context.Context) CreateRecoveryPlanInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateRecoveryPlanInputPropertiesOutput)
}

type CreateRecoveryPlanInputPropertiesOutput struct{ *pulumi.OutputState }

func (CreateRecoveryPlanInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateRecoveryPlanInputProperties)(nil)).Elem()
}

func (o CreateRecoveryPlanInputPropertiesOutput) ToCreateRecoveryPlanInputPropertiesOutput() CreateRecoveryPlanInputPropertiesOutput {
	return o
}

func (o CreateRecoveryPlanInputPropertiesOutput) ToCreateRecoveryPlanInputPropertiesOutputWithContext(ctx context.Context) CreateRecoveryPlanInputPropertiesOutput {
	return o
}

func (o CreateRecoveryPlanInputPropertiesOutput) FailoverDeploymentModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateRecoveryPlanInputProperties) *string { return v.FailoverDeploymentModel }).(pulumi.StringPtrOutput)
}

func (o CreateRecoveryPlanInputPropertiesOutput) Groups() RecoveryPlanGroupArrayOutput {
	return o.ApplyT(func(v CreateRecoveryPlanInputProperties) []RecoveryPlanGroup { return v.Groups }).(RecoveryPlanGroupArrayOutput)
}

func (o CreateRecoveryPlanInputPropertiesOutput) PrimaryFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v CreateRecoveryPlanInputProperties) string { return v.PrimaryFabricId }).(pulumi.StringOutput)
}

func (o CreateRecoveryPlanInputPropertiesOutput) ProviderSpecificInput() RecoveryPlanA2AInputArrayOutput {
	return o.ApplyT(func(v CreateRecoveryPlanInputProperties) []RecoveryPlanA2AInput { return v.ProviderSpecificInput }).(RecoveryPlanA2AInputArrayOutput)
}

func (o CreateRecoveryPlanInputPropertiesOutput) RecoveryFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v CreateRecoveryPlanInputProperties) string { return v.RecoveryFabricId }).(pulumi.StringOutput)
}

type CurrentJobDetailsResponse struct {
	JobId     *string `pulumi:"jobId"`
	JobName   *string `pulumi:"jobName"`
	StartTime *string `pulumi:"startTime"`
}

type CurrentJobDetailsResponseOutput struct{ *pulumi.OutputState }

func (CurrentJobDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CurrentJobDetailsResponse)(nil)).Elem()
}

func (o CurrentJobDetailsResponseOutput) ToCurrentJobDetailsResponseOutput() CurrentJobDetailsResponseOutput {
	return o
}

func (o CurrentJobDetailsResponseOutput) ToCurrentJobDetailsResponseOutputWithContext(ctx context.Context) CurrentJobDetailsResponseOutput {
	return o
}

func (o CurrentJobDetailsResponseOutput) JobId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CurrentJobDetailsResponse) *string { return v.JobId }).(pulumi.StringPtrOutput)
}

func (o CurrentJobDetailsResponseOutput) JobName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CurrentJobDetailsResponse) *string { return v.JobName }).(pulumi.StringPtrOutput)
}

func (o CurrentJobDetailsResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CurrentJobDetailsResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type CurrentScenarioDetailsResponse struct {
	JobId        *string `pulumi:"jobId"`
	ScenarioName *string `pulumi:"scenarioName"`
	StartTime    *string `pulumi:"startTime"`
}

type CurrentScenarioDetailsResponseOutput struct{ *pulumi.OutputState }

func (CurrentScenarioDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CurrentScenarioDetailsResponse)(nil)).Elem()
}

func (o CurrentScenarioDetailsResponseOutput) ToCurrentScenarioDetailsResponseOutput() CurrentScenarioDetailsResponseOutput {
	return o
}

func (o CurrentScenarioDetailsResponseOutput) ToCurrentScenarioDetailsResponseOutputWithContext(ctx context.Context) CurrentScenarioDetailsResponseOutput {
	return o
}

func (o CurrentScenarioDetailsResponseOutput) JobId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CurrentScenarioDetailsResponse) *string { return v.JobId }).(pulumi.StringPtrOutput)
}

func (o CurrentScenarioDetailsResponseOutput) ScenarioName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CurrentScenarioDetailsResponse) *string { return v.ScenarioName }).(pulumi.StringPtrOutput)
}

func (o CurrentScenarioDetailsResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CurrentScenarioDetailsResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type CurrentScenarioDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (CurrentScenarioDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CurrentScenarioDetailsResponse)(nil)).Elem()
}

func (o CurrentScenarioDetailsResponsePtrOutput) ToCurrentScenarioDetailsResponsePtrOutput() CurrentScenarioDetailsResponsePtrOutput {
	return o
}

func (o CurrentScenarioDetailsResponsePtrOutput) ToCurrentScenarioDetailsResponsePtrOutputWithContext(ctx context.Context) CurrentScenarioDetailsResponsePtrOutput {
	return o
}

func (o CurrentScenarioDetailsResponsePtrOutput) Elem() CurrentScenarioDetailsResponseOutput {
	return o.ApplyT(func(v *CurrentScenarioDetailsResponse) CurrentScenarioDetailsResponse {
		if v != nil {
			return *v
		}
		var ret CurrentScenarioDetailsResponse
		return ret
	}).(CurrentScenarioDetailsResponseOutput)
}

func (o CurrentScenarioDetailsResponsePtrOutput) JobId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CurrentScenarioDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.JobId
	}).(pulumi.StringPtrOutput)
}

func (o CurrentScenarioDetailsResponsePtrOutput) ScenarioName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CurrentScenarioDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScenarioName
	}).(pulumi.StringPtrOutput)
}

func (o CurrentScenarioDetailsResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CurrentScenarioDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
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

type DataStoreResponse struct {
	Capacity     *string `pulumi:"capacity"`
	FreeSpace    *string `pulumi:"freeSpace"`
	SymbolicName *string `pulumi:"symbolicName"`
	Type         *string `pulumi:"type"`
	Uuid         *string `pulumi:"uuid"`
}

type Day struct {
	Date   *int  `pulumi:"date"`
	IsLast *bool `pulumi:"isLast"`
}

type DayResponse struct {
	Date   *int  `pulumi:"date"`
	IsLast *bool `pulumi:"isLast"`
}

type DiskDetailsResponse struct {
	MaxSizeMB *float64 `pulumi:"maxSizeMB"`
	VhdId     *string  `pulumi:"vhdId"`
	VhdName   *string  `pulumi:"vhdName"`
	VhdType   *string  `pulumi:"vhdType"`
}

type DiskEncryptionInfo struct {
	DiskEncryptionKeyInfo *DiskEncryptionKeyInfo `pulumi:"diskEncryptionKeyInfo"`
	KeyEncryptionKeyInfo  *KeyEncryptionKeyInfo  `pulumi:"keyEncryptionKeyInfo"`
}

type DiskEncryptionKeyInfo struct {
	KeyVaultResourceArmId *string `pulumi:"keyVaultResourceArmId"`
	SecretIdentifier      *string `pulumi:"secretIdentifier"`
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

type DraDetailsResponse struct {
	Health           string                `pulumi:"health"`
	HealthErrors     []HealthErrorResponse `pulumi:"healthErrors"`
	Id               string                `pulumi:"id"`
	LastHeartbeatUtc string                `pulumi:"lastHeartbeatUtc"`
	Name             string                `pulumi:"name"`
	Version          string                `pulumi:"version"`
}

type EnableMigrationInputProperties struct {
	PolicyId                string                        `pulumi:"policyId"`
	ProviderSpecificDetails VMwareCbtEnableMigrationInput `pulumi:"providerSpecificDetails"`
}





type EnableMigrationInputPropertiesInput interface {
	pulumi.Input

	ToEnableMigrationInputPropertiesOutput() EnableMigrationInputPropertiesOutput
	ToEnableMigrationInputPropertiesOutputWithContext(context.Context) EnableMigrationInputPropertiesOutput
}

type EnableMigrationInputPropertiesArgs struct {
	PolicyId                pulumi.StringInput                 `pulumi:"policyId"`
	ProviderSpecificDetails VMwareCbtEnableMigrationInputInput `pulumi:"providerSpecificDetails"`
}

func (EnableMigrationInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnableMigrationInputProperties)(nil)).Elem()
}

func (i EnableMigrationInputPropertiesArgs) ToEnableMigrationInputPropertiesOutput() EnableMigrationInputPropertiesOutput {
	return i.ToEnableMigrationInputPropertiesOutputWithContext(context.Background())
}

func (i EnableMigrationInputPropertiesArgs) ToEnableMigrationInputPropertiesOutputWithContext(ctx context.Context) EnableMigrationInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnableMigrationInputPropertiesOutput)
}

type EnableMigrationInputPropertiesOutput struct{ *pulumi.OutputState }

func (EnableMigrationInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnableMigrationInputProperties)(nil)).Elem()
}

func (o EnableMigrationInputPropertiesOutput) ToEnableMigrationInputPropertiesOutput() EnableMigrationInputPropertiesOutput {
	return o
}

func (o EnableMigrationInputPropertiesOutput) ToEnableMigrationInputPropertiesOutputWithContext(ctx context.Context) EnableMigrationInputPropertiesOutput {
	return o
}

func (o EnableMigrationInputPropertiesOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v EnableMigrationInputProperties) string { return v.PolicyId }).(pulumi.StringOutput)
}

func (o EnableMigrationInputPropertiesOutput) ProviderSpecificDetails() VMwareCbtEnableMigrationInputOutput {
	return o.ApplyT(func(v EnableMigrationInputProperties) VMwareCbtEnableMigrationInput { return v.ProviderSpecificDetails }).(VMwareCbtEnableMigrationInputOutput)
}

type EnableProtectionInputProperties struct {
	PolicyId                *string     `pulumi:"policyId"`
	ProtectableItemId       *string     `pulumi:"protectableItemId"`
	ProviderSpecificDetails interface{} `pulumi:"providerSpecificDetails"`
}





type EnableProtectionInputPropertiesInput interface {
	pulumi.Input

	ToEnableProtectionInputPropertiesOutput() EnableProtectionInputPropertiesOutput
	ToEnableProtectionInputPropertiesOutputWithContext(context.Context) EnableProtectionInputPropertiesOutput
}

type EnableProtectionInputPropertiesArgs struct {
	PolicyId                pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectableItemId       pulumi.StringPtrInput `pulumi:"protectableItemId"`
	ProviderSpecificDetails pulumi.Input          `pulumi:"providerSpecificDetails"`
}

func (EnableProtectionInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnableProtectionInputProperties)(nil)).Elem()
}

func (i EnableProtectionInputPropertiesArgs) ToEnableProtectionInputPropertiesOutput() EnableProtectionInputPropertiesOutput {
	return i.ToEnableProtectionInputPropertiesOutputWithContext(context.Background())
}

func (i EnableProtectionInputPropertiesArgs) ToEnableProtectionInputPropertiesOutputWithContext(ctx context.Context) EnableProtectionInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnableProtectionInputPropertiesOutput)
}

func (i EnableProtectionInputPropertiesArgs) ToEnableProtectionInputPropertiesPtrOutput() EnableProtectionInputPropertiesPtrOutput {
	return i.ToEnableProtectionInputPropertiesPtrOutputWithContext(context.Background())
}

func (i EnableProtectionInputPropertiesArgs) ToEnableProtectionInputPropertiesPtrOutputWithContext(ctx context.Context) EnableProtectionInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnableProtectionInputPropertiesOutput).ToEnableProtectionInputPropertiesPtrOutputWithContext(ctx)
}









type EnableProtectionInputPropertiesPtrInput interface {
	pulumi.Input

	ToEnableProtectionInputPropertiesPtrOutput() EnableProtectionInputPropertiesPtrOutput
	ToEnableProtectionInputPropertiesPtrOutputWithContext(context.Context) EnableProtectionInputPropertiesPtrOutput
}

type enableProtectionInputPropertiesPtrType EnableProtectionInputPropertiesArgs

func EnableProtectionInputPropertiesPtr(v *EnableProtectionInputPropertiesArgs) EnableProtectionInputPropertiesPtrInput {
	return (*enableProtectionInputPropertiesPtrType)(v)
}

func (*enableProtectionInputPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnableProtectionInputProperties)(nil)).Elem()
}

func (i *enableProtectionInputPropertiesPtrType) ToEnableProtectionInputPropertiesPtrOutput() EnableProtectionInputPropertiesPtrOutput {
	return i.ToEnableProtectionInputPropertiesPtrOutputWithContext(context.Background())
}

func (i *enableProtectionInputPropertiesPtrType) ToEnableProtectionInputPropertiesPtrOutputWithContext(ctx context.Context) EnableProtectionInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnableProtectionInputPropertiesPtrOutput)
}

type EnableProtectionInputPropertiesOutput struct{ *pulumi.OutputState }

func (EnableProtectionInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnableProtectionInputProperties)(nil)).Elem()
}

func (o EnableProtectionInputPropertiesOutput) ToEnableProtectionInputPropertiesOutput() EnableProtectionInputPropertiesOutput {
	return o
}

func (o EnableProtectionInputPropertiesOutput) ToEnableProtectionInputPropertiesOutputWithContext(ctx context.Context) EnableProtectionInputPropertiesOutput {
	return o
}

func (o EnableProtectionInputPropertiesOutput) ToEnableProtectionInputPropertiesPtrOutput() EnableProtectionInputPropertiesPtrOutput {
	return o.ToEnableProtectionInputPropertiesPtrOutputWithContext(context.Background())
}

func (o EnableProtectionInputPropertiesOutput) ToEnableProtectionInputPropertiesPtrOutputWithContext(ctx context.Context) EnableProtectionInputPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnableProtectionInputProperties) *EnableProtectionInputProperties {
		return &v
	}).(EnableProtectionInputPropertiesPtrOutput)
}

func (o EnableProtectionInputPropertiesOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnableProtectionInputProperties) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o EnableProtectionInputPropertiesOutput) ProtectableItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnableProtectionInputProperties) *string { return v.ProtectableItemId }).(pulumi.StringPtrOutput)
}

func (o EnableProtectionInputPropertiesOutput) ProviderSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v EnableProtectionInputProperties) interface{} { return v.ProviderSpecificDetails }).(pulumi.AnyOutput)
}

type EnableProtectionInputPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EnableProtectionInputPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnableProtectionInputProperties)(nil)).Elem()
}

func (o EnableProtectionInputPropertiesPtrOutput) ToEnableProtectionInputPropertiesPtrOutput() EnableProtectionInputPropertiesPtrOutput {
	return o
}

func (o EnableProtectionInputPropertiesPtrOutput) ToEnableProtectionInputPropertiesPtrOutputWithContext(ctx context.Context) EnableProtectionInputPropertiesPtrOutput {
	return o
}

func (o EnableProtectionInputPropertiesPtrOutput) Elem() EnableProtectionInputPropertiesOutput {
	return o.ApplyT(func(v *EnableProtectionInputProperties) EnableProtectionInputProperties {
		if v != nil {
			return *v
		}
		var ret EnableProtectionInputProperties
		return ret
	}).(EnableProtectionInputPropertiesOutput)
}

func (o EnableProtectionInputPropertiesPtrOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnableProtectionInputProperties) *string {
		if v == nil {
			return nil
		}
		return v.PolicyId
	}).(pulumi.StringPtrOutput)
}

func (o EnableProtectionInputPropertiesPtrOutput) ProtectableItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnableProtectionInputProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProtectableItemId
	}).(pulumi.StringPtrOutput)
}

func (o EnableProtectionInputPropertiesPtrOutput) ProviderSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *EnableProtectionInputProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.ProviderSpecificDetails
	}).(pulumi.AnyOutput)
}

type EncryptionDetailsResponse struct {
	KekCertExpiryDate *string `pulumi:"kekCertExpiryDate"`
	KekCertThumbprint *string `pulumi:"kekCertThumbprint"`
	KekState          *string `pulumi:"kekState"`
}

type EncryptionDetailsResponseOutput struct{ *pulumi.OutputState }

func (EncryptionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionDetailsResponse)(nil)).Elem()
}

func (o EncryptionDetailsResponseOutput) ToEncryptionDetailsResponseOutput() EncryptionDetailsResponseOutput {
	return o
}

func (o EncryptionDetailsResponseOutput) ToEncryptionDetailsResponseOutputWithContext(ctx context.Context) EncryptionDetailsResponseOutput {
	return o
}

func (o EncryptionDetailsResponseOutput) KekCertExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionDetailsResponse) *string { return v.KekCertExpiryDate }).(pulumi.StringPtrOutput)
}

func (o EncryptionDetailsResponseOutput) KekCertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionDetailsResponse) *string { return v.KekCertThumbprint }).(pulumi.StringPtrOutput)
}

func (o EncryptionDetailsResponseOutput) KekState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionDetailsResponse) *string { return v.KekState }).(pulumi.StringPtrOutput)
}

type EncryptionDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionDetailsResponse)(nil)).Elem()
}

func (o EncryptionDetailsResponsePtrOutput) ToEncryptionDetailsResponsePtrOutput() EncryptionDetailsResponsePtrOutput {
	return o
}

func (o EncryptionDetailsResponsePtrOutput) ToEncryptionDetailsResponsePtrOutputWithContext(ctx context.Context) EncryptionDetailsResponsePtrOutput {
	return o
}

func (o EncryptionDetailsResponsePtrOutput) Elem() EncryptionDetailsResponseOutput {
	return o.ApplyT(func(v *EncryptionDetailsResponse) EncryptionDetailsResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionDetailsResponse
		return ret
	}).(EncryptionDetailsResponseOutput)
}

func (o EncryptionDetailsResponsePtrOutput) KekCertExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.KekCertExpiryDate
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionDetailsResponsePtrOutput) KekCertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.KekCertThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionDetailsResponsePtrOutput) KekState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.KekState
	}).(pulumi.StringPtrOutput)
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

type FabricCreationInputProperties struct {
	CustomDetails interface{} `pulumi:"customDetails"`
}





type FabricCreationInputPropertiesInput interface {
	pulumi.Input

	ToFabricCreationInputPropertiesOutput() FabricCreationInputPropertiesOutput
	ToFabricCreationInputPropertiesOutputWithContext(context.Context) FabricCreationInputPropertiesOutput
}

type FabricCreationInputPropertiesArgs struct {
	CustomDetails pulumi.Input `pulumi:"customDetails"`
}

func (FabricCreationInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FabricCreationInputProperties)(nil)).Elem()
}

func (i FabricCreationInputPropertiesArgs) ToFabricCreationInputPropertiesOutput() FabricCreationInputPropertiesOutput {
	return i.ToFabricCreationInputPropertiesOutputWithContext(context.Background())
}

func (i FabricCreationInputPropertiesArgs) ToFabricCreationInputPropertiesOutputWithContext(ctx context.Context) FabricCreationInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricCreationInputPropertiesOutput)
}

func (i FabricCreationInputPropertiesArgs) ToFabricCreationInputPropertiesPtrOutput() FabricCreationInputPropertiesPtrOutput {
	return i.ToFabricCreationInputPropertiesPtrOutputWithContext(context.Background())
}

func (i FabricCreationInputPropertiesArgs) ToFabricCreationInputPropertiesPtrOutputWithContext(ctx context.Context) FabricCreationInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricCreationInputPropertiesOutput).ToFabricCreationInputPropertiesPtrOutputWithContext(ctx)
}









type FabricCreationInputPropertiesPtrInput interface {
	pulumi.Input

	ToFabricCreationInputPropertiesPtrOutput() FabricCreationInputPropertiesPtrOutput
	ToFabricCreationInputPropertiesPtrOutputWithContext(context.Context) FabricCreationInputPropertiesPtrOutput
}

type fabricCreationInputPropertiesPtrType FabricCreationInputPropertiesArgs

func FabricCreationInputPropertiesPtr(v *FabricCreationInputPropertiesArgs) FabricCreationInputPropertiesPtrInput {
	return (*fabricCreationInputPropertiesPtrType)(v)
}

func (*fabricCreationInputPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FabricCreationInputProperties)(nil)).Elem()
}

func (i *fabricCreationInputPropertiesPtrType) ToFabricCreationInputPropertiesPtrOutput() FabricCreationInputPropertiesPtrOutput {
	return i.ToFabricCreationInputPropertiesPtrOutputWithContext(context.Background())
}

func (i *fabricCreationInputPropertiesPtrType) ToFabricCreationInputPropertiesPtrOutputWithContext(ctx context.Context) FabricCreationInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricCreationInputPropertiesPtrOutput)
}

type FabricCreationInputPropertiesOutput struct{ *pulumi.OutputState }

func (FabricCreationInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FabricCreationInputProperties)(nil)).Elem()
}

func (o FabricCreationInputPropertiesOutput) ToFabricCreationInputPropertiesOutput() FabricCreationInputPropertiesOutput {
	return o
}

func (o FabricCreationInputPropertiesOutput) ToFabricCreationInputPropertiesOutputWithContext(ctx context.Context) FabricCreationInputPropertiesOutput {
	return o
}

func (o FabricCreationInputPropertiesOutput) ToFabricCreationInputPropertiesPtrOutput() FabricCreationInputPropertiesPtrOutput {
	return o.ToFabricCreationInputPropertiesPtrOutputWithContext(context.Background())
}

func (o FabricCreationInputPropertiesOutput) ToFabricCreationInputPropertiesPtrOutputWithContext(ctx context.Context) FabricCreationInputPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FabricCreationInputProperties) *FabricCreationInputProperties {
		return &v
	}).(FabricCreationInputPropertiesPtrOutput)
}

func (o FabricCreationInputPropertiesOutput) CustomDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v FabricCreationInputProperties) interface{} { return v.CustomDetails }).(pulumi.AnyOutput)
}

type FabricCreationInputPropertiesPtrOutput struct{ *pulumi.OutputState }

func (FabricCreationInputPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FabricCreationInputProperties)(nil)).Elem()
}

func (o FabricCreationInputPropertiesPtrOutput) ToFabricCreationInputPropertiesPtrOutput() FabricCreationInputPropertiesPtrOutput {
	return o
}

func (o FabricCreationInputPropertiesPtrOutput) ToFabricCreationInputPropertiesPtrOutputWithContext(ctx context.Context) FabricCreationInputPropertiesPtrOutput {
	return o
}

func (o FabricCreationInputPropertiesPtrOutput) Elem() FabricCreationInputPropertiesOutput {
	return o.ApplyT(func(v *FabricCreationInputProperties) FabricCreationInputProperties {
		if v != nil {
			return *v
		}
		var ret FabricCreationInputProperties
		return ret
	}).(FabricCreationInputPropertiesOutput)
}

func (o FabricCreationInputPropertiesPtrOutput) CustomDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *FabricCreationInputProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.CustomDetails
	}).(pulumi.AnyOutput)
}

type FabricPropertiesResponse struct {
	BcdrState                 *string                    `pulumi:"bcdrState"`
	CustomDetails             interface{}                `pulumi:"customDetails"`
	EncryptionDetails         *EncryptionDetailsResponse `pulumi:"encryptionDetails"`
	FriendlyName              *string                    `pulumi:"friendlyName"`
	Health                    *string                    `pulumi:"health"`
	HealthErrorDetails        []HealthErrorResponse      `pulumi:"healthErrorDetails"`
	InternalIdentifier        *string                    `pulumi:"internalIdentifier"`
	RolloverEncryptionDetails *EncryptionDetailsResponse `pulumi:"rolloverEncryptionDetails"`
}

type FabricPropertiesResponseOutput struct{ *pulumi.OutputState }

func (FabricPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FabricPropertiesResponse)(nil)).Elem()
}

func (o FabricPropertiesResponseOutput) ToFabricPropertiesResponseOutput() FabricPropertiesResponseOutput {
	return o
}

func (o FabricPropertiesResponseOutput) ToFabricPropertiesResponseOutputWithContext(ctx context.Context) FabricPropertiesResponseOutput {
	return o
}

func (o FabricPropertiesResponseOutput) BcdrState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FabricPropertiesResponse) *string { return v.BcdrState }).(pulumi.StringPtrOutput)
}

func (o FabricPropertiesResponseOutput) CustomDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v FabricPropertiesResponse) interface{} { return v.CustomDetails }).(pulumi.AnyOutput)
}

func (o FabricPropertiesResponseOutput) EncryptionDetails() EncryptionDetailsResponsePtrOutput {
	return o.ApplyT(func(v FabricPropertiesResponse) *EncryptionDetailsResponse { return v.EncryptionDetails }).(EncryptionDetailsResponsePtrOutput)
}

func (o FabricPropertiesResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FabricPropertiesResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o FabricPropertiesResponseOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FabricPropertiesResponse) *string { return v.Health }).(pulumi.StringPtrOutput)
}

func (o FabricPropertiesResponseOutput) HealthErrorDetails() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v FabricPropertiesResponse) []HealthErrorResponse { return v.HealthErrorDetails }).(HealthErrorResponseArrayOutput)
}

func (o FabricPropertiesResponseOutput) InternalIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FabricPropertiesResponse) *string { return v.InternalIdentifier }).(pulumi.StringPtrOutput)
}

func (o FabricPropertiesResponseOutput) RolloverEncryptionDetails() EncryptionDetailsResponsePtrOutput {
	return o.ApplyT(func(v FabricPropertiesResponse) *EncryptionDetailsResponse { return v.RolloverEncryptionDetails }).(EncryptionDetailsResponsePtrOutput)
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

type HealthErrorResponse struct {
	CreationTimeUtc              *string                    `pulumi:"creationTimeUtc"`
	CustomerResolvability        *string                    `pulumi:"customerResolvability"`
	EntityId                     *string                    `pulumi:"entityId"`
	ErrorCategory                *string                    `pulumi:"errorCategory"`
	ErrorCode                    *string                    `pulumi:"errorCode"`
	ErrorId                      *string                    `pulumi:"errorId"`
	ErrorLevel                   *string                    `pulumi:"errorLevel"`
	ErrorMessage                 *string                    `pulumi:"errorMessage"`
	ErrorSource                  *string                    `pulumi:"errorSource"`
	ErrorType                    *string                    `pulumi:"errorType"`
	InnerHealthErrors            []InnerHealthErrorResponse `pulumi:"innerHealthErrors"`
	PossibleCauses               *string                    `pulumi:"possibleCauses"`
	RecommendedAction            *string                    `pulumi:"recommendedAction"`
	RecoveryProviderErrorMessage *string                    `pulumi:"recoveryProviderErrorMessage"`
	SummaryMessage               *string                    `pulumi:"summaryMessage"`
}

type HealthErrorResponseOutput struct{ *pulumi.OutputState }

func (HealthErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthErrorResponse)(nil)).Elem()
}

func (o HealthErrorResponseOutput) ToHealthErrorResponseOutput() HealthErrorResponseOutput {
	return o
}

func (o HealthErrorResponseOutput) ToHealthErrorResponseOutputWithContext(ctx context.Context) HealthErrorResponseOutput {
	return o
}

func (o HealthErrorResponseOutput) CreationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.CreationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) CustomerResolvability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.CustomerResolvability }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) EntityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.EntityId }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) ErrorCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.ErrorCategory }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) ErrorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.ErrorId }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) ErrorLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.ErrorLevel }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) ErrorSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.ErrorSource }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) ErrorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.ErrorType }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) InnerHealthErrors() InnerHealthErrorResponseArrayOutput {
	return o.ApplyT(func(v HealthErrorResponse) []InnerHealthErrorResponse { return v.InnerHealthErrors }).(InnerHealthErrorResponseArrayOutput)
}

func (o HealthErrorResponseOutput) PossibleCauses() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.PossibleCauses }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) RecommendedAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.RecommendedAction }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) RecoveryProviderErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.RecoveryProviderErrorMessage }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) SummaryMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.SummaryMessage }).(pulumi.StringPtrOutput)
}

type HealthErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthErrorResponse)(nil)).Elem()
}

func (o HealthErrorResponseArrayOutput) ToHealthErrorResponseArrayOutput() HealthErrorResponseArrayOutput {
	return o
}

func (o HealthErrorResponseArrayOutput) ToHealthErrorResponseArrayOutputWithContext(ctx context.Context) HealthErrorResponseArrayOutput {
	return o
}

func (o HealthErrorResponseArrayOutput) Index(i pulumi.IntInput) HealthErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthErrorResponse {
		return vs[0].([]HealthErrorResponse)[vs[1].(int)]
	}).(HealthErrorResponseOutput)
}

type HyperVReplicaAzureEnableProtectionInput struct {
	DisksToInclude                  []string `pulumi:"disksToInclude"`
	EnableRdpOnTargetOption         *string  `pulumi:"enableRdpOnTargetOption"`
	HvHostVmId                      *string  `pulumi:"hvHostVmId"`
	InstanceType                    *string  `pulumi:"instanceType"`
	LogStorageAccountId             *string  `pulumi:"logStorageAccountId"`
	OsType                          *string  `pulumi:"osType"`
	TargetAvailabilitySetId         *string  `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone          *string  `pulumi:"targetAvailabilityZone"`
	TargetAzureNetworkId            *string  `pulumi:"targetAzureNetworkId"`
	TargetAzureSubnetId             *string  `pulumi:"targetAzureSubnetId"`
	TargetAzureV1ResourceGroupId    *string  `pulumi:"targetAzureV1ResourceGroupId"`
	TargetAzureV2ResourceGroupId    *string  `pulumi:"targetAzureV2ResourceGroupId"`
	TargetAzureVmName               *string  `pulumi:"targetAzureVmName"`
	TargetProximityPlacementGroupId *string  `pulumi:"targetProximityPlacementGroupId"`
	TargetStorageAccountId          *string  `pulumi:"targetStorageAccountId"`
	TargetVmSize                    *string  `pulumi:"targetVmSize"`
	UseManagedDisks                 *string  `pulumi:"useManagedDisks"`
	VhdId                           *string  `pulumi:"vhdId"`
	VmName                          *string  `pulumi:"vmName"`
}

type HyperVReplicaAzurePolicyDetailsResponse struct {
	ActiveStorageAccountId                        *string `pulumi:"activeStorageAccountId"`
	ApplicationConsistentSnapshotFrequencyInHours *int    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Encryption                                    *string `pulumi:"encryption"`
	InstanceType                                  string  `pulumi:"instanceType"`
	OnlineReplicationStartTime                    *string `pulumi:"onlineReplicationStartTime"`
	RecoveryPointHistoryDurationInHours           *int    `pulumi:"recoveryPointHistoryDurationInHours"`
	ReplicationInterval                           *int    `pulumi:"replicationInterval"`
}

type HyperVReplicaAzurePolicyInput struct {
	ApplicationConsistentSnapshotFrequencyInHours *int     `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	InstanceType                                  *string  `pulumi:"instanceType"`
	OnlineReplicationStartTime                    *string  `pulumi:"onlineReplicationStartTime"`
	RecoveryPointHistoryDuration                  *int     `pulumi:"recoveryPointHistoryDuration"`
	ReplicationInterval                           *int     `pulumi:"replicationInterval"`
	StorageAccounts                               []string `pulumi:"storageAccounts"`
}

type HyperVReplicaAzureReplicationDetailsResponse struct {
	AzureVmDiskDetails               []AzureVmDiskDetailsResponse       `pulumi:"azureVmDiskDetails"`
	EnableRdpOnTargetOption          *string                            `pulumi:"enableRdpOnTargetOption"`
	Encryption                       *string                            `pulumi:"encryption"`
	InitialReplicationDetails        *InitialReplicationDetailsResponse `pulumi:"initialReplicationDetails"`
	InstanceType                     string                             `pulumi:"instanceType"`
	LastReplicatedTime               *string                            `pulumi:"lastReplicatedTime"`
	LastRpoCalculatedTime            *string                            `pulumi:"lastRpoCalculatedTime"`
	LicenseType                      *string                            `pulumi:"licenseType"`
	OSDetails                        *OSDetailsResponse                 `pulumi:"oSDetails"`
	RecoveryAvailabilitySetId        *string                            `pulumi:"recoveryAvailabilitySetId"`
	RecoveryAzureLogStorageAccountId *string                            `pulumi:"recoveryAzureLogStorageAccountId"`
	RecoveryAzureResourceGroupId     *string                            `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureStorageAccount      *string                            `pulumi:"recoveryAzureStorageAccount"`
	RecoveryAzureVMSize              *string                            `pulumi:"recoveryAzureVMSize"`
	RecoveryAzureVmName              *string                            `pulumi:"recoveryAzureVmName"`
	RpoInSeconds                     *float64                           `pulumi:"rpoInSeconds"`
	SelectedRecoveryAzureNetworkId   *string                            `pulumi:"selectedRecoveryAzureNetworkId"`
	SelectedSourceNicId              *string                            `pulumi:"selectedSourceNicId"`
	SourceVmCpuCount                 *int                               `pulumi:"sourceVmCpuCount"`
	SourceVmRamSizeInMB              *int                               `pulumi:"sourceVmRamSizeInMB"`
	TargetAvailabilityZone           *string                            `pulumi:"targetAvailabilityZone"`
	TargetProximityPlacementGroupId  *string                            `pulumi:"targetProximityPlacementGroupId"`
	UseManagedDisks                  *string                            `pulumi:"useManagedDisks"`
	VmId                             *string                            `pulumi:"vmId"`
	VmNics                           []VMNicDetailsResponse             `pulumi:"vmNics"`
	VmProtectionState                *string                            `pulumi:"vmProtectionState"`
	VmProtectionStateDescription     *string                            `pulumi:"vmProtectionStateDescription"`
}

type HyperVReplicaBasePolicyDetailsResponse struct {
	AllowedAuthenticationType                     *int    `pulumi:"allowedAuthenticationType"`
	ApplicationConsistentSnapshotFrequencyInHours *int    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Compression                                   *string `pulumi:"compression"`
	InitialReplicationMethod                      *string `pulumi:"initialReplicationMethod"`
	InstanceType                                  string  `pulumi:"instanceType"`
	OfflineReplicationExportPath                  *string `pulumi:"offlineReplicationExportPath"`
	OfflineReplicationImportPath                  *string `pulumi:"offlineReplicationImportPath"`
	OnlineReplicationStartTime                    *string `pulumi:"onlineReplicationStartTime"`
	RecoveryPoints                                *int    `pulumi:"recoveryPoints"`
	ReplicaDeletionOption                         *string `pulumi:"replicaDeletionOption"`
	ReplicationPort                               *int    `pulumi:"replicationPort"`
}

type HyperVReplicaBaseReplicationDetailsResponse struct {
	InitialReplicationDetails    *InitialReplicationDetailsResponse `pulumi:"initialReplicationDetails"`
	InstanceType                 string                             `pulumi:"instanceType"`
	LastReplicatedTime           *string                            `pulumi:"lastReplicatedTime"`
	VMDiskDetails                []DiskDetailsResponse              `pulumi:"vMDiskDetails"`
	VmId                         *string                            `pulumi:"vmId"`
	VmNics                       []VMNicDetailsResponse             `pulumi:"vmNics"`
	VmProtectionState            *string                            `pulumi:"vmProtectionState"`
	VmProtectionStateDescription *string                            `pulumi:"vmProtectionStateDescription"`
}

type HyperVReplicaBluePolicyDetailsResponse struct {
	AllowedAuthenticationType                     *int    `pulumi:"allowedAuthenticationType"`
	ApplicationConsistentSnapshotFrequencyInHours *int    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Compression                                   *string `pulumi:"compression"`
	InitialReplicationMethod                      *string `pulumi:"initialReplicationMethod"`
	InstanceType                                  string  `pulumi:"instanceType"`
	OfflineReplicationExportPath                  *string `pulumi:"offlineReplicationExportPath"`
	OfflineReplicationImportPath                  *string `pulumi:"offlineReplicationImportPath"`
	OnlineReplicationStartTime                    *string `pulumi:"onlineReplicationStartTime"`
	RecoveryPoints                                *int    `pulumi:"recoveryPoints"`
	ReplicaDeletionOption                         *string `pulumi:"replicaDeletionOption"`
	ReplicationFrequencyInSeconds                 *int    `pulumi:"replicationFrequencyInSeconds"`
	ReplicationPort                               *int    `pulumi:"replicationPort"`
}

type HyperVReplicaBluePolicyInput struct {
	AllowedAuthenticationType                     *int    `pulumi:"allowedAuthenticationType"`
	ApplicationConsistentSnapshotFrequencyInHours *int    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Compression                                   *string `pulumi:"compression"`
	InitialReplicationMethod                      *string `pulumi:"initialReplicationMethod"`
	InstanceType                                  *string `pulumi:"instanceType"`
	OfflineReplicationExportPath                  *string `pulumi:"offlineReplicationExportPath"`
	OfflineReplicationImportPath                  *string `pulumi:"offlineReplicationImportPath"`
	OnlineReplicationStartTime                    *string `pulumi:"onlineReplicationStartTime"`
	RecoveryPoints                                *int    `pulumi:"recoveryPoints"`
	ReplicaDeletion                               *string `pulumi:"replicaDeletion"`
	ReplicationFrequencyInSeconds                 *int    `pulumi:"replicationFrequencyInSeconds"`
	ReplicationPort                               *int    `pulumi:"replicationPort"`
}

type HyperVReplicaBlueReplicationDetailsResponse struct {
	InitialReplicationDetails    *InitialReplicationDetailsResponse `pulumi:"initialReplicationDetails"`
	InstanceType                 string                             `pulumi:"instanceType"`
	LastReplicatedTime           *string                            `pulumi:"lastReplicatedTime"`
	VMDiskDetails                []DiskDetailsResponse              `pulumi:"vMDiskDetails"`
	VmId                         *string                            `pulumi:"vmId"`
	VmNics                       []VMNicDetailsResponse             `pulumi:"vmNics"`
	VmProtectionState            *string                            `pulumi:"vmProtectionState"`
	VmProtectionStateDescription *string                            `pulumi:"vmProtectionStateDescription"`
}

type HyperVReplicaPolicyDetailsResponse struct {
	AllowedAuthenticationType                     *int    `pulumi:"allowedAuthenticationType"`
	ApplicationConsistentSnapshotFrequencyInHours *int    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Compression                                   *string `pulumi:"compression"`
	InitialReplicationMethod                      *string `pulumi:"initialReplicationMethod"`
	InstanceType                                  string  `pulumi:"instanceType"`
	OfflineReplicationExportPath                  *string `pulumi:"offlineReplicationExportPath"`
	OfflineReplicationImportPath                  *string `pulumi:"offlineReplicationImportPath"`
	OnlineReplicationStartTime                    *string `pulumi:"onlineReplicationStartTime"`
	RecoveryPoints                                *int    `pulumi:"recoveryPoints"`
	ReplicaDeletionOption                         *string `pulumi:"replicaDeletionOption"`
	ReplicationPort                               *int    `pulumi:"replicationPort"`
}

type HyperVReplicaPolicyInput struct {
	AllowedAuthenticationType                     *int    `pulumi:"allowedAuthenticationType"`
	ApplicationConsistentSnapshotFrequencyInHours *int    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Compression                                   *string `pulumi:"compression"`
	InitialReplicationMethod                      *string `pulumi:"initialReplicationMethod"`
	InstanceType                                  *string `pulumi:"instanceType"`
	OfflineReplicationExportPath                  *string `pulumi:"offlineReplicationExportPath"`
	OfflineReplicationImportPath                  *string `pulumi:"offlineReplicationImportPath"`
	OnlineReplicationStartTime                    *string `pulumi:"onlineReplicationStartTime"`
	RecoveryPoints                                *int    `pulumi:"recoveryPoints"`
	ReplicaDeletion                               *string `pulumi:"replicaDeletion"`
	ReplicationPort                               *int    `pulumi:"replicationPort"`
}

type HyperVReplicaReplicationDetailsResponse struct {
	InitialReplicationDetails    *InitialReplicationDetailsResponse `pulumi:"initialReplicationDetails"`
	InstanceType                 string                             `pulumi:"instanceType"`
	LastReplicatedTime           *string                            `pulumi:"lastReplicatedTime"`
	VMDiskDetails                []DiskDetailsResponse              `pulumi:"vMDiskDetails"`
	VmId                         *string                            `pulumi:"vmId"`
	VmNics                       []VMNicDetailsResponse             `pulumi:"vmNics"`
	VmProtectionState            *string                            `pulumi:"vmProtectionState"`
	VmProtectionStateDescription *string                            `pulumi:"vmProtectionStateDescription"`
}

type HyperVSiteDetailsResponse struct {
	InstanceType string `pulumi:"instanceType"`
}

type IPConfigResponse struct {
	LBBackendAddressPoolIds []string `pulumi:"lBBackendAddressPoolIds"`
	PublicIpAddressId       *string  `pulumi:"publicIpAddressId"`
	StaticIPAddress         *string  `pulumi:"staticIPAddress"`
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

type IdentityProviderDetailsResponse struct {
	AadAuthority  *string `pulumi:"aadAuthority"`
	ApplicationId *string `pulumi:"applicationId"`
	Audience      *string `pulumi:"audience"`
	ObjectId      *string `pulumi:"objectId"`
	TenantId      *string `pulumi:"tenantId"`
}

type IdentityProviderDetailsResponseOutput struct{ *pulumi.OutputState }

func (IdentityProviderDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderDetailsResponse)(nil)).Elem()
}

func (o IdentityProviderDetailsResponseOutput) ToIdentityProviderDetailsResponseOutput() IdentityProviderDetailsResponseOutput {
	return o
}

func (o IdentityProviderDetailsResponseOutput) ToIdentityProviderDetailsResponseOutputWithContext(ctx context.Context) IdentityProviderDetailsResponseOutput {
	return o
}

func (o IdentityProviderDetailsResponseOutput) AadAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProviderDetailsResponse) *string { return v.AadAuthority }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderDetailsResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProviderDetailsResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderDetailsResponseOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProviderDetailsResponse) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderDetailsResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProviderDetailsResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o IdentityProviderDetailsResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProviderDetailsResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type IdentityProviderDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityProviderDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderDetailsResponse)(nil)).Elem()
}

func (o IdentityProviderDetailsResponsePtrOutput) ToIdentityProviderDetailsResponsePtrOutput() IdentityProviderDetailsResponsePtrOutput {
	return o
}

func (o IdentityProviderDetailsResponsePtrOutput) ToIdentityProviderDetailsResponsePtrOutputWithContext(ctx context.Context) IdentityProviderDetailsResponsePtrOutput {
	return o
}

func (o IdentityProviderDetailsResponsePtrOutput) Elem() IdentityProviderDetailsResponseOutput {
	return o.ApplyT(func(v *IdentityProviderDetailsResponse) IdentityProviderDetailsResponse {
		if v != nil {
			return *v
		}
		var ret IdentityProviderDetailsResponse
		return ret
	}).(IdentityProviderDetailsResponseOutput)
}

func (o IdentityProviderDetailsResponsePtrOutput) AadAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProviderDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AadAuthority
	}).(pulumi.StringPtrOutput)
}

func (o IdentityProviderDetailsResponsePtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProviderDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityProviderDetailsResponsePtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProviderDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o IdentityProviderDetailsResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProviderDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityProviderDetailsResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProviderDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type IdentityProviderInput struct {
	AadAuthority  string `pulumi:"aadAuthority"`
	ApplicationId string `pulumi:"applicationId"`
	Audience      string `pulumi:"audience"`
	ObjectId      string `pulumi:"objectId"`
	TenantId      string `pulumi:"tenantId"`
}





type IdentityProviderInputInput interface {
	pulumi.Input

	ToIdentityProviderInputOutput() IdentityProviderInputOutput
	ToIdentityProviderInputOutputWithContext(context.Context) IdentityProviderInputOutput
}

type IdentityProviderInputArgs struct {
	AadAuthority  pulumi.StringInput `pulumi:"aadAuthority"`
	ApplicationId pulumi.StringInput `pulumi:"applicationId"`
	Audience      pulumi.StringInput `pulumi:"audience"`
	ObjectId      pulumi.StringInput `pulumi:"objectId"`
	TenantId      pulumi.StringInput `pulumi:"tenantId"`
}

func (IdentityProviderInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderInput)(nil)).Elem()
}

func (i IdentityProviderInputArgs) ToIdentityProviderInputOutput() IdentityProviderInputOutput {
	return i.ToIdentityProviderInputOutputWithContext(context.Background())
}

func (i IdentityProviderInputArgs) ToIdentityProviderInputOutputWithContext(ctx context.Context) IdentityProviderInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderInputOutput)
}

func (i IdentityProviderInputArgs) ToIdentityProviderInputPtrOutput() IdentityProviderInputPtrOutput {
	return i.ToIdentityProviderInputPtrOutputWithContext(context.Background())
}

func (i IdentityProviderInputArgs) ToIdentityProviderInputPtrOutputWithContext(ctx context.Context) IdentityProviderInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderInputOutput).ToIdentityProviderInputPtrOutputWithContext(ctx)
}









type IdentityProviderInputPtrInput interface {
	pulumi.Input

	ToIdentityProviderInputPtrOutput() IdentityProviderInputPtrOutput
	ToIdentityProviderInputPtrOutputWithContext(context.Context) IdentityProviderInputPtrOutput
}

type identityProviderInputPtrType IdentityProviderInputArgs

func IdentityProviderInputPtr(v *IdentityProviderInputArgs) IdentityProviderInputPtrInput {
	return (*identityProviderInputPtrType)(v)
}

func (*identityProviderInputPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderInput)(nil)).Elem()
}

func (i *identityProviderInputPtrType) ToIdentityProviderInputPtrOutput() IdentityProviderInputPtrOutput {
	return i.ToIdentityProviderInputPtrOutputWithContext(context.Background())
}

func (i *identityProviderInputPtrType) ToIdentityProviderInputPtrOutputWithContext(ctx context.Context) IdentityProviderInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderInputPtrOutput)
}

type IdentityProviderInputOutput struct{ *pulumi.OutputState }

func (IdentityProviderInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderInput)(nil)).Elem()
}

func (o IdentityProviderInputOutput) ToIdentityProviderInputOutput() IdentityProviderInputOutput {
	return o
}

func (o IdentityProviderInputOutput) ToIdentityProviderInputOutputWithContext(ctx context.Context) IdentityProviderInputOutput {
	return o
}

func (o IdentityProviderInputOutput) ToIdentityProviderInputPtrOutput() IdentityProviderInputPtrOutput {
	return o.ToIdentityProviderInputPtrOutputWithContext(context.Background())
}

func (o IdentityProviderInputOutput) ToIdentityProviderInputPtrOutputWithContext(ctx context.Context) IdentityProviderInputPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProviderInput) *IdentityProviderInput {
		return &v
	}).(IdentityProviderInputPtrOutput)
}

func (o IdentityProviderInputOutput) AadAuthority() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityProviderInput) string { return v.AadAuthority }).(pulumi.StringOutput)
}

func (o IdentityProviderInputOutput) ApplicationId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityProviderInput) string { return v.ApplicationId }).(pulumi.StringOutput)
}

func (o IdentityProviderInputOutput) Audience() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityProviderInput) string { return v.Audience }).(pulumi.StringOutput)
}

func (o IdentityProviderInputOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityProviderInput) string { return v.ObjectId }).(pulumi.StringOutput)
}

func (o IdentityProviderInputOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityProviderInput) string { return v.TenantId }).(pulumi.StringOutput)
}

type IdentityProviderInputPtrOutput struct{ *pulumi.OutputState }

func (IdentityProviderInputPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderInput)(nil)).Elem()
}

func (o IdentityProviderInputPtrOutput) ToIdentityProviderInputPtrOutput() IdentityProviderInputPtrOutput {
	return o
}

func (o IdentityProviderInputPtrOutput) ToIdentityProviderInputPtrOutputWithContext(ctx context.Context) IdentityProviderInputPtrOutput {
	return o
}

func (o IdentityProviderInputPtrOutput) Elem() IdentityProviderInputOutput {
	return o.ApplyT(func(v *IdentityProviderInput) IdentityProviderInput {
		if v != nil {
			return *v
		}
		var ret IdentityProviderInput
		return ret
	}).(IdentityProviderInputOutput)
}

func (o IdentityProviderInputPtrOutput) AadAuthority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProviderInput) *string {
		if v == nil {
			return nil
		}
		return &v.AadAuthority
	}).(pulumi.StringPtrOutput)
}

func (o IdentityProviderInputPtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProviderInput) *string {
		if v == nil {
			return nil
		}
		return &v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityProviderInputPtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProviderInput) *string {
		if v == nil {
			return nil
		}
		return &v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o IdentityProviderInputPtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProviderInput) *string {
		if v == nil {
			return nil
		}
		return &v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityProviderInputPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProviderInput) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

type InMageAgentDetailsResponse struct {
	AgentExpiryDate        *string `pulumi:"agentExpiryDate"`
	AgentUpdateStatus      *string `pulumi:"agentUpdateStatus"`
	AgentVersion           *string `pulumi:"agentVersion"`
	PostUpdateRebootStatus *string `pulumi:"postUpdateRebootStatus"`
}

type InMageAzureV2DiskInputDetails struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskId              *string `pulumi:"diskId"`
	DiskType            *string `pulumi:"diskType"`
	LogStorageAccountId *string `pulumi:"logStorageAccountId"`
}

type InMageAzureV2EnableProtectionInput struct {
	DiskEncryptionSetId             *string                         `pulumi:"diskEncryptionSetId"`
	DiskType                        *string                         `pulumi:"diskType"`
	DisksToInclude                  []InMageAzureV2DiskInputDetails `pulumi:"disksToInclude"`
	EnableRdpOnTargetOption         *string                         `pulumi:"enableRdpOnTargetOption"`
	InstanceType                    *string                         `pulumi:"instanceType"`
	LogStorageAccountId             *string                         `pulumi:"logStorageAccountId"`
	MasterTargetId                  *string                         `pulumi:"masterTargetId"`
	MultiVmGroupId                  *string                         `pulumi:"multiVmGroupId"`
	MultiVmGroupName                *string                         `pulumi:"multiVmGroupName"`
	ProcessServerId                 *string                         `pulumi:"processServerId"`
	RunAsAccountId                  *string                         `pulumi:"runAsAccountId"`
	StorageAccountId                *string                         `pulumi:"storageAccountId"`
	TargetAvailabilitySetId         *string                         `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone          *string                         `pulumi:"targetAvailabilityZone"`
	TargetAzureNetworkId            *string                         `pulumi:"targetAzureNetworkId"`
	TargetAzureSubnetId             *string                         `pulumi:"targetAzureSubnetId"`
	TargetAzureV1ResourceGroupId    *string                         `pulumi:"targetAzureV1ResourceGroupId"`
	TargetAzureV2ResourceGroupId    *string                         `pulumi:"targetAzureV2ResourceGroupId"`
	TargetAzureVmName               *string                         `pulumi:"targetAzureVmName"`
	TargetProximityPlacementGroupId *string                         `pulumi:"targetProximityPlacementGroupId"`
	TargetVmSize                    *string                         `pulumi:"targetVmSize"`
}

type InMageAzureV2ManagedDiskDetailsResponse struct {
	DiskId            *string `pulumi:"diskId"`
	ReplicaDiskType   *string `pulumi:"replicaDiskType"`
	SeedManagedDiskId *string `pulumi:"seedManagedDiskId"`
}

type InMageAzureV2PolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string  `pulumi:"instanceType"`
	MultiVmSyncStatus                 *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int    `pulumi:"recoveryPointThresholdInMinutes"`
}

type InMageAzureV2PolicyInput struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      *string `pulumi:"instanceType"`
	MultiVmSyncStatus                 string  `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int    `pulumi:"recoveryPointThresholdInMinutes"`
}

type InMageAzureV2ProtectedDiskDetailsResponse struct {
	DiskCapacityInBytes       *float64 `pulumi:"diskCapacityInBytes"`
	DiskId                    *string  `pulumi:"diskId"`
	DiskName                  *string  `pulumi:"diskName"`
	DiskResized               *string  `pulumi:"diskResized"`
	FileSystemCapacityInBytes *float64 `pulumi:"fileSystemCapacityInBytes"`
	HealthErrorCode           *string  `pulumi:"healthErrorCode"`
	LastRpoCalculatedTime     *string  `pulumi:"lastRpoCalculatedTime"`
	ProtectionStage           *string  `pulumi:"protectionStage"`
	PsDataInMegaBytes         *float64 `pulumi:"psDataInMegaBytes"`
	ResyncDurationInSeconds   *float64 `pulumi:"resyncDurationInSeconds"`
	ResyncProgressPercentage  *int     `pulumi:"resyncProgressPercentage"`
	ResyncRequired            *string  `pulumi:"resyncRequired"`
	RpoInSeconds              *float64 `pulumi:"rpoInSeconds"`
	SourceDataInMegaBytes     *float64 `pulumi:"sourceDataInMegaBytes"`
	TargetDataInMegaBytes     *float64 `pulumi:"targetDataInMegaBytes"`
}

type InMageAzureV2ReplicationDetailsResponse struct {
	AgentExpiryDate                  *string                                     `pulumi:"agentExpiryDate"`
	AgentVersion                     *string                                     `pulumi:"agentVersion"`
	AzureVMDiskDetails               []AzureVmDiskDetailsResponse                `pulumi:"azureVMDiskDetails"`
	CompressedDataRateInMB           *float64                                    `pulumi:"compressedDataRateInMB"`
	Datastores                       []string                                    `pulumi:"datastores"`
	DiscoveryType                    *string                                     `pulumi:"discoveryType"`
	DiskResized                      *string                                     `pulumi:"diskResized"`
	EnableRdpOnTargetOption          *string                                     `pulumi:"enableRdpOnTargetOption"`
	InfrastructureVmId               *string                                     `pulumi:"infrastructureVmId"`
	InstanceType                     string                                      `pulumi:"instanceType"`
	IpAddress                        *string                                     `pulumi:"ipAddress"`
	IsAgentUpdateRequired            *string                                     `pulumi:"isAgentUpdateRequired"`
	IsRebootAfterUpdateRequired      *string                                     `pulumi:"isRebootAfterUpdateRequired"`
	LastHeartbeat                    *string                                     `pulumi:"lastHeartbeat"`
	LastRpoCalculatedTime            *string                                     `pulumi:"lastRpoCalculatedTime"`
	LastUpdateReceivedTime           *string                                     `pulumi:"lastUpdateReceivedTime"`
	LicenseType                      *string                                     `pulumi:"licenseType"`
	MasterTargetId                   *string                                     `pulumi:"masterTargetId"`
	MultiVmGroupId                   *string                                     `pulumi:"multiVmGroupId"`
	MultiVmGroupName                 *string                                     `pulumi:"multiVmGroupName"`
	MultiVmSyncStatus                *string                                     `pulumi:"multiVmSyncStatus"`
	OsDiskId                         *string                                     `pulumi:"osDiskId"`
	OsType                           *string                                     `pulumi:"osType"`
	OsVersion                        *string                                     `pulumi:"osVersion"`
	ProcessServerId                  *string                                     `pulumi:"processServerId"`
	ProcessServerName                string                                      `pulumi:"processServerName"`
	ProtectedDisks                   []InMageAzureV2ProtectedDiskDetailsResponse `pulumi:"protectedDisks"`
	ProtectedManagedDisks            []InMageAzureV2ManagedDiskDetailsResponse   `pulumi:"protectedManagedDisks"`
	ProtectionStage                  *string                                     `pulumi:"protectionStage"`
	RecoveryAvailabilitySetId        *string                                     `pulumi:"recoveryAvailabilitySetId"`
	RecoveryAzureLogStorageAccountId *string                                     `pulumi:"recoveryAzureLogStorageAccountId"`
	RecoveryAzureResourceGroupId     *string                                     `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureStorageAccount      *string                                     `pulumi:"recoveryAzureStorageAccount"`
	RecoveryAzureVMName              *string                                     `pulumi:"recoveryAzureVMName"`
	RecoveryAzureVMSize              *string                                     `pulumi:"recoveryAzureVMSize"`
	ReplicaId                        *string                                     `pulumi:"replicaId"`
	ResyncProgressPercentage         *int                                        `pulumi:"resyncProgressPercentage"`
	RpoInSeconds                     *float64                                    `pulumi:"rpoInSeconds"`
	SelectedRecoveryAzureNetworkId   *string                                     `pulumi:"selectedRecoveryAzureNetworkId"`
	SelectedSourceNicId              *string                                     `pulumi:"selectedSourceNicId"`
	SelectedTfoAzureNetworkId        *string                                     `pulumi:"selectedTfoAzureNetworkId"`
	SourceVmCpuCount                 *int                                        `pulumi:"sourceVmCpuCount"`
	SourceVmRamSizeInMB              *int                                        `pulumi:"sourceVmRamSizeInMB"`
	TargetAvailabilityZone           *string                                     `pulumi:"targetAvailabilityZone"`
	TargetProximityPlacementGroupId  *string                                     `pulumi:"targetProximityPlacementGroupId"`
	TargetVmId                       *string                                     `pulumi:"targetVmId"`
	UncompressedDataRateInMB         *float64                                    `pulumi:"uncompressedDataRateInMB"`
	UseManagedDisks                  *string                                     `pulumi:"useManagedDisks"`
	VCenterInfrastructureId          *string                                     `pulumi:"vCenterInfrastructureId"`
	ValidationErrors                 []HealthErrorResponse                       `pulumi:"validationErrors"`
	VhdName                          *string                                     `pulumi:"vhdName"`
	VmId                             *string                                     `pulumi:"vmId"`
	VmNics                           []VMNicDetailsResponse                      `pulumi:"vmNics"`
	VmProtectionState                *string                                     `pulumi:"vmProtectionState"`
	VmProtectionStateDescription     *string                                     `pulumi:"vmProtectionStateDescription"`
}

type InMageBasePolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes *int    `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    string  `pulumi:"instanceType"`
	MultiVmSyncStatus               *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes *int    `pulumi:"recoveryPointThresholdInMinutes"`
}

type InMageDiskExclusionInput struct {
	DiskSignatureOptions []InMageDiskSignatureExclusionOptions `pulumi:"diskSignatureOptions"`
	VolumeOptions        []InMageVolumeExclusionOptions        `pulumi:"volumeOptions"`
}

type InMageDiskSignatureExclusionOptions struct {
	DiskSignature *string `pulumi:"diskSignature"`
}

type InMageEnableProtectionInput struct {
	DatastoreName      *string                   `pulumi:"datastoreName"`
	DiskExclusionInput *InMageDiskExclusionInput `pulumi:"diskExclusionInput"`
	DisksToInclude     []string                  `pulumi:"disksToInclude"`
	InstanceType       *string                   `pulumi:"instanceType"`
	MasterTargetId     string                    `pulumi:"masterTargetId"`
	MultiVmGroupId     string                    `pulumi:"multiVmGroupId"`
	MultiVmGroupName   string                    `pulumi:"multiVmGroupName"`
	ProcessServerId    string                    `pulumi:"processServerId"`
	RetentionDrive     string                    `pulumi:"retentionDrive"`
	RunAsAccountId     *string                   `pulumi:"runAsAccountId"`
	VmFriendlyName     *string                   `pulumi:"vmFriendlyName"`
}

type InMagePolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes *int    `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    string  `pulumi:"instanceType"`
	MultiVmSyncStatus               *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes *int    `pulumi:"recoveryPointThresholdInMinutes"`
}

type InMagePolicyInput struct {
	AppConsistentFrequencyInMinutes *int    `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    *string `pulumi:"instanceType"`
	MultiVmSyncStatus               string  `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes *int    `pulumi:"recoveryPointThresholdInMinutes"`
}

type InMageProtectedDiskDetailsResponse struct {
	DiskCapacityInBytes       *float64 `pulumi:"diskCapacityInBytes"`
	DiskId                    *string  `pulumi:"diskId"`
	DiskName                  *string  `pulumi:"diskName"`
	DiskResized               *string  `pulumi:"diskResized"`
	FileSystemCapacityInBytes *float64 `pulumi:"fileSystemCapacityInBytes"`
	HealthErrorCode           *string  `pulumi:"healthErrorCode"`
	LastRpoCalculatedTime     *string  `pulumi:"lastRpoCalculatedTime"`
	ProtectionStage           *string  `pulumi:"protectionStage"`
	PsDataInMB                *float64 `pulumi:"psDataInMB"`
	ResyncDurationInSeconds   *float64 `pulumi:"resyncDurationInSeconds"`
	ResyncProgressPercentage  *int     `pulumi:"resyncProgressPercentage"`
	ResyncRequired            *string  `pulumi:"resyncRequired"`
	RpoInSeconds              *float64 `pulumi:"rpoInSeconds"`
	SourceDataInMB            *float64 `pulumi:"sourceDataInMB"`
	TargetDataInMB            *float64 `pulumi:"targetDataInMB"`
}

type InMageRcmAgentUpgradeBlockingErrorDetailsResponse struct {
	ErrorCode              string            `pulumi:"errorCode"`
	ErrorMessage           string            `pulumi:"errorMessage"`
	ErrorMessageParameters map[string]string `pulumi:"errorMessageParameters"`
	ErrorTags              map[string]string `pulumi:"errorTags"`
	PossibleCauses         string            `pulumi:"possibleCauses"`
	RecommendedAction      string            `pulumi:"recommendedAction"`
}

type InMageRcmDiskInput struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskId              *string `pulumi:"diskId"`
	DiskType            *string `pulumi:"diskType"`
	LogStorageAccountId *string `pulumi:"logStorageAccountId"`
}

type InMageRcmDisksDefaultInput struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskType            *string `pulumi:"diskType"`
	LogStorageAccountId *string `pulumi:"logStorageAccountId"`
}

type InMageRcmEnableProtectionInput struct {
	DisksDefault                          *InMageRcmDisksDefaultInput `pulumi:"disksDefault"`
	DisksToInclude                        []InMageRcmDiskInput        `pulumi:"disksToInclude"`
	FabricDiscoveryMachineId              *string                     `pulumi:"fabricDiscoveryMachineId"`
	InstanceType                          *string                     `pulumi:"instanceType"`
	LicenseType                           *string                     `pulumi:"licenseType"`
	MultiVmGroupName                      *string                     `pulumi:"multiVmGroupName"`
	ProcessServerId                       *string                     `pulumi:"processServerId"`
	RunAsAccountId                        *string                     `pulumi:"runAsAccountId"`
	TargetAvailabilitySetId               *string                     `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                *string                     `pulumi:"targetAvailabilityZone"`
	TargetBootDiagnosticsStorageAccountId *string                     `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetNetworkId                       *string                     `pulumi:"targetNetworkId"`
	TargetProximityPlacementGroupId       *string                     `pulumi:"targetProximityPlacementGroupId"`
	TargetResourceGroupId                 *string                     `pulumi:"targetResourceGroupId"`
	TargetSubnetName                      *string                     `pulumi:"targetSubnetName"`
	TargetVmName                          *string                     `pulumi:"targetVmName"`
	TargetVmSize                          *string                     `pulumi:"targetVmSize"`
	TestNetworkId                         *string                     `pulumi:"testNetworkId"`
	TestSubnetName                        *string                     `pulumi:"testSubnetName"`
}

type InMageRcmFabricCreationInput struct {
	AuthCertificate     *string                `pulumi:"authCertificate"`
	InstanceType        *string                `pulumi:"instanceType"`
	PhysicalSiteId      *string                `pulumi:"physicalSiteId"`
	SourceAgentIdentity *IdentityProviderInput `pulumi:"sourceAgentIdentity"`
	VmwareSiteId        *string                `pulumi:"vmwareSiteId"`
}

type InMageRcmFabricSpecificDetailsResponse struct {
	AgentDetails       []AgentDetailsResponse            `pulumi:"agentDetails"`
	ControlPlaneUri    string                            `pulumi:"controlPlaneUri"`
	DataPlaneUri       string                            `pulumi:"dataPlaneUri"`
	Dras               []DraDetailsResponse              `pulumi:"dras"`
	InstanceType       string                            `pulumi:"instanceType"`
	PhysicalSiteId     string                            `pulumi:"physicalSiteId"`
	ProcessServers     []ProcessServerDetailsResponse    `pulumi:"processServers"`
	PushInstallers     []PushInstallerDetailsResponse    `pulumi:"pushInstallers"`
	RcmProxies         []RcmProxyDetailsResponse         `pulumi:"rcmProxies"`
	ReplicationAgents  []ReplicationAgentDetailsResponse `pulumi:"replicationAgents"`
	ReprotectAgents    []ReprotectAgentDetailsResponse   `pulumi:"reprotectAgents"`
	ServiceContainerId string                            `pulumi:"serviceContainerId"`
	ServiceEndpoint    string                            `pulumi:"serviceEndpoint"`
	ServiceResourceId  string                            `pulumi:"serviceResourceId"`
	VmwareSiteId       string                            `pulumi:"vmwareSiteId"`
}

type InMageRcmLastAgentUpgradeErrorDetailsResponse struct {
	ErrorCode              string            `pulumi:"errorCode"`
	ErrorMessage           string            `pulumi:"errorMessage"`
	ErrorMessageParameters map[string]string `pulumi:"errorMessageParameters"`
	ErrorTags              map[string]string `pulumi:"errorTags"`
	PossibleCauses         string            `pulumi:"possibleCauses"`
	RecommendedAction      string            `pulumi:"recommendedAction"`
}

type InMageRcmMobilityAgentDetailsResponse struct {
	AgentVersionExpiryDate               string   `pulumi:"agentVersionExpiryDate"`
	DriverVersion                        string   `pulumi:"driverVersion"`
	DriverVersionExpiryDate              string   `pulumi:"driverVersionExpiryDate"`
	IsUpgradeable                        string   `pulumi:"isUpgradeable"`
	LastHeartbeatUtc                     string   `pulumi:"lastHeartbeatUtc"`
	LatestUpgradableVersionWithoutReboot string   `pulumi:"latestUpgradableVersionWithoutReboot"`
	LatestVersion                        string   `pulumi:"latestVersion"`
	ReasonsBlockingUpgrade               []string `pulumi:"reasonsBlockingUpgrade"`
	Version                              string   `pulumi:"version"`
}

type InMageRcmNicDetailsResponse struct {
	IsPrimaryNic          string `pulumi:"isPrimaryNic"`
	IsSelectedForFailover string `pulumi:"isSelectedForFailover"`
	NicId                 string `pulumi:"nicId"`
	SourceIPAddress       string `pulumi:"sourceIPAddress"`
	SourceIPAddressType   string `pulumi:"sourceIPAddressType"`
	SourceNetworkId       string `pulumi:"sourceNetworkId"`
	SourceSubnetName      string `pulumi:"sourceSubnetName"`
	TargetIPAddress       string `pulumi:"targetIPAddress"`
	TargetIPAddressType   string `pulumi:"targetIPAddressType"`
	TargetSubnetName      string `pulumi:"targetSubnetName"`
	TestIPAddress         string `pulumi:"testIPAddress"`
	TestIPAddressType     string `pulumi:"testIPAddressType"`
	TestSubnetName        string `pulumi:"testSubnetName"`
}

type InMageRcmPolicyCreationInput struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	EnableMultiVmSync                 *string `pulumi:"enableMultiVmSync"`
	InstanceType                      *string `pulumi:"instanceType"`
	RecoveryPointHistoryInMinutes     *int    `pulumi:"recoveryPointHistoryInMinutes"`
}

type InMageRcmPolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes int    `pulumi:"crashConsistentFrequencyInMinutes"`
	EnableMultiVmSync                 string `pulumi:"enableMultiVmSync"`
	InstanceType                      string `pulumi:"instanceType"`
	RecoveryPointHistoryInMinutes     int    `pulumi:"recoveryPointHistoryInMinutes"`
}

type InMageRcmProtectedDiskDetailsResponse struct {
	CapacityInBytes     float64 `pulumi:"capacityInBytes"`
	DiskEncryptionSetId string  `pulumi:"diskEncryptionSetId"`
	DiskId              string  `pulumi:"diskId"`
	DiskName            string  `pulumi:"diskName"`
	DiskType            string  `pulumi:"diskType"`
	IsOSDisk            string  `pulumi:"isOSDisk"`
	LogStorageAccountId string  `pulumi:"logStorageAccountId"`
	SeedManagedDiskId   string  `pulumi:"seedManagedDiskId"`
	TargetManagedDiskId string  `pulumi:"targetManagedDiskId"`
}

type InMageRcmReplicationDetailsResponse struct {
	AgentUpgradeBlockingErrorDetails      []InMageRcmAgentUpgradeBlockingErrorDetailsResponse `pulumi:"agentUpgradeBlockingErrorDetails"`
	AgentUpgradeState                     string                                              `pulumi:"agentUpgradeState"`
	AllocatedMemoryInMB                   float64                                             `pulumi:"allocatedMemoryInMB"`
	DiscoveryType                         string                                              `pulumi:"discoveryType"`
	FabricDiscoveryMachineId              string                                              `pulumi:"fabricDiscoveryMachineId"`
	FailoverRecoveryPointId               string                                              `pulumi:"failoverRecoveryPointId"`
	FirmwareType                          string                                              `pulumi:"firmwareType"`
	InitialReplicationProcessedBytes      float64                                             `pulumi:"initialReplicationProcessedBytes"`
	InitialReplicationProgressPercentage  int                                                 `pulumi:"initialReplicationProgressPercentage"`
	InitialReplicationTransferredBytes    float64                                             `pulumi:"initialReplicationTransferredBytes"`
	InstanceType                          string                                              `pulumi:"instanceType"`
	InternalIdentifier                    string                                              `pulumi:"internalIdentifier"`
	IsLastUpgradeSuccessful               string                                              `pulumi:"isLastUpgradeSuccessful"`
	LastAgentUpgradeErrorDetails          []InMageRcmLastAgentUpgradeErrorDetailsResponse     `pulumi:"lastAgentUpgradeErrorDetails"`
	LastAgentUpgradeFailedJobId           string                                              `pulumi:"lastAgentUpgradeFailedJobId"`
	LastAgentUpgradeType                  string                                              `pulumi:"lastAgentUpgradeType"`
	LastRecoveryPointId                   string                                              `pulumi:"lastRecoveryPointId"`
	LastRecoveryPointReceived             string                                              `pulumi:"lastRecoveryPointReceived"`
	LastRpoCalculatedTime                 string                                              `pulumi:"lastRpoCalculatedTime"`
	LastRpoInSeconds                      float64                                             `pulumi:"lastRpoInSeconds"`
	LicenseType                           string                                              `pulumi:"licenseType"`
	MobilityAgentDetails                  InMageRcmMobilityAgentDetailsResponse               `pulumi:"mobilityAgentDetails"`
	MultiVmGroupName                      string                                              `pulumi:"multiVmGroupName"`
	OsType                                string                                              `pulumi:"osType"`
	ProcessServerId                       string                                              `pulumi:"processServerId"`
	ProcessServerName                     string                                              `pulumi:"processServerName"`
	ProcessorCoreCount                    int                                                 `pulumi:"processorCoreCount"`
	ProtectedDisks                        []InMageRcmProtectedDiskDetailsResponse             `pulumi:"protectedDisks"`
	ResyncProcessedBytes                  float64                                             `pulumi:"resyncProcessedBytes"`
	ResyncProgressPercentage              int                                                 `pulumi:"resyncProgressPercentage"`
	ResyncRequired                        string                                              `pulumi:"resyncRequired"`
	ResyncState                           string                                              `pulumi:"resyncState"`
	ResyncTransferredBytes                float64                                             `pulumi:"resyncTransferredBytes"`
	RunAsAccountId                        string                                              `pulumi:"runAsAccountId"`
	TargetAvailabilitySetId               string                                              `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                string                                              `pulumi:"targetAvailabilityZone"`
	TargetBootDiagnosticsStorageAccountId string                                              `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetGeneration                      string                                              `pulumi:"targetGeneration"`
	TargetNetworkId                       string                                              `pulumi:"targetNetworkId"`
	TargetProximityPlacementGroupId       string                                              `pulumi:"targetProximityPlacementGroupId"`
	TargetResourceGroupId                 string                                              `pulumi:"targetResourceGroupId"`
	TargetVmName                          string                                              `pulumi:"targetVmName"`
	TargetVmSize                          string                                              `pulumi:"targetVmSize"`
	TestNetworkId                         string                                              `pulumi:"testNetworkId"`
	VmNics                                []InMageRcmNicDetailsResponse                       `pulumi:"vmNics"`
}

type InMageReplicationDetailsResponse struct {
	ActiveSiteType               *string                              `pulumi:"activeSiteType"`
	AgentDetails                 *InMageAgentDetailsResponse          `pulumi:"agentDetails"`
	AzureStorageAccountId        *string                              `pulumi:"azureStorageAccountId"`
	CompressedDataRateInMB       *float64                             `pulumi:"compressedDataRateInMB"`
	ConsistencyPoints            map[string]string                    `pulumi:"consistencyPoints"`
	Datastores                   []string                             `pulumi:"datastores"`
	DiscoveryType                *string                              `pulumi:"discoveryType"`
	DiskResized                  *string                              `pulumi:"diskResized"`
	InfrastructureVmId           *string                              `pulumi:"infrastructureVmId"`
	InstanceType                 string                               `pulumi:"instanceType"`
	IpAddress                    *string                              `pulumi:"ipAddress"`
	LastHeartbeat                *string                              `pulumi:"lastHeartbeat"`
	LastRpoCalculatedTime        *string                              `pulumi:"lastRpoCalculatedTime"`
	LastUpdateReceivedTime       *string                              `pulumi:"lastUpdateReceivedTime"`
	MasterTargetId               *string                              `pulumi:"masterTargetId"`
	MultiVmGroupId               *string                              `pulumi:"multiVmGroupId"`
	MultiVmGroupName             *string                              `pulumi:"multiVmGroupName"`
	MultiVmSyncStatus            *string                              `pulumi:"multiVmSyncStatus"`
	OsDetails                    *OSDiskDetailsResponse               `pulumi:"osDetails"`
	OsVersion                    *string                              `pulumi:"osVersion"`
	ProcessServerId              *string                              `pulumi:"processServerId"`
	ProtectedDisks               []InMageProtectedDiskDetailsResponse `pulumi:"protectedDisks"`
	ProtectionStage              *string                              `pulumi:"protectionStage"`
	RebootAfterUpdateStatus      *string                              `pulumi:"rebootAfterUpdateStatus"`
	ReplicaId                    *string                              `pulumi:"replicaId"`
	ResyncDetails                *InitialReplicationDetailsResponse   `pulumi:"resyncDetails"`
	RetentionWindowEnd           *string                              `pulumi:"retentionWindowEnd"`
	RetentionWindowStart         *string                              `pulumi:"retentionWindowStart"`
	RpoInSeconds                 *float64                             `pulumi:"rpoInSeconds"`
	SourceVmCpuCount             *int                                 `pulumi:"sourceVmCpuCount"`
	SourceVmRamSizeInMB          *int                                 `pulumi:"sourceVmRamSizeInMB"`
	UncompressedDataRateInMB     *float64                             `pulumi:"uncompressedDataRateInMB"`
	VCenterInfrastructureId      *string                              `pulumi:"vCenterInfrastructureId"`
	ValidationErrors             []HealthErrorResponse                `pulumi:"validationErrors"`
	VmId                         *string                              `pulumi:"vmId"`
	VmNics                       []VMNicDetailsResponse               `pulumi:"vmNics"`
	VmProtectionState            *string                              `pulumi:"vmProtectionState"`
	VmProtectionStateDescription *string                              `pulumi:"vmProtectionStateDescription"`
}

type InMageVolumeExclusionOptions struct {
	OnlyExcludeIfSingleVolume *string `pulumi:"onlyExcludeIfSingleVolume"`
	VolumeLabel               *string `pulumi:"volumeLabel"`
}

type InitialReplicationDetailsResponse struct {
	InitialReplicationProgressPercentage *string `pulumi:"initialReplicationProgressPercentage"`
	InitialReplicationType               *string `pulumi:"initialReplicationType"`
}

type InnerHealthErrorResponse struct {
	CreationTimeUtc              *string `pulumi:"creationTimeUtc"`
	EntityId                     *string `pulumi:"entityId"`
	ErrorCategory                *string `pulumi:"errorCategory"`
	ErrorCode                    *string `pulumi:"errorCode"`
	ErrorLevel                   *string `pulumi:"errorLevel"`
	ErrorMessage                 *string `pulumi:"errorMessage"`
	ErrorSource                  *string `pulumi:"errorSource"`
	ErrorType                    *string `pulumi:"errorType"`
	PossibleCauses               *string `pulumi:"possibleCauses"`
	RecommendedAction            *string `pulumi:"recommendedAction"`
	RecoveryProviderErrorMessage *string `pulumi:"recoveryProviderErrorMessage"`
	SummaryMessage               *string `pulumi:"summaryMessage"`
}

type InnerHealthErrorResponseOutput struct{ *pulumi.OutputState }

func (InnerHealthErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InnerHealthErrorResponse)(nil)).Elem()
}

func (o InnerHealthErrorResponseOutput) ToInnerHealthErrorResponseOutput() InnerHealthErrorResponseOutput {
	return o
}

func (o InnerHealthErrorResponseOutput) ToInnerHealthErrorResponseOutputWithContext(ctx context.Context) InnerHealthErrorResponseOutput {
	return o
}

func (o InnerHealthErrorResponseOutput) CreationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.CreationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) EntityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.EntityId }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) ErrorCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.ErrorCategory }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) ErrorLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.ErrorLevel }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) ErrorSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.ErrorSource }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) ErrorType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.ErrorType }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) PossibleCauses() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.PossibleCauses }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) RecommendedAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.RecommendedAction }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) RecoveryProviderErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.RecoveryProviderErrorMessage }).(pulumi.StringPtrOutput)
}

func (o InnerHealthErrorResponseOutput) SummaryMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.SummaryMessage }).(pulumi.StringPtrOutput)
}

type InnerHealthErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (InnerHealthErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InnerHealthErrorResponse)(nil)).Elem()
}

func (o InnerHealthErrorResponseArrayOutput) ToInnerHealthErrorResponseArrayOutput() InnerHealthErrorResponseArrayOutput {
	return o
}

func (o InnerHealthErrorResponseArrayOutput) ToInnerHealthErrorResponseArrayOutputWithContext(ctx context.Context) InnerHealthErrorResponseArrayOutput {
	return o
}

func (o InnerHealthErrorResponseArrayOutput) Index(i pulumi.IntInput) InnerHealthErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InnerHealthErrorResponse {
		return vs[0].([]InnerHealthErrorResponse)[vs[1].(int)]
	}).(InnerHealthErrorResponseOutput)
}

type InputEndpointResponse struct {
	EndpointName *string `pulumi:"endpointName"`
	PrivatePort  *int    `pulumi:"privatePort"`
	Protocol     *string `pulumi:"protocol"`
	PublicPort   *int    `pulumi:"publicPort"`
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

type KeyEncryptionKeyInfo struct {
	KeyIdentifier         *string `pulumi:"keyIdentifier"`
	KeyVaultResourceArmId *string `pulumi:"keyVaultResourceArmId"`
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

type MasterTargetServerResponse struct {
	AgentExpiryDate         *string                   `pulumi:"agentExpiryDate"`
	AgentVersion            *string                   `pulumi:"agentVersion"`
	AgentVersionDetails     *VersionDetailsResponse   `pulumi:"agentVersionDetails"`
	DataStores              []DataStoreResponse       `pulumi:"dataStores"`
	DiskCount               *int                      `pulumi:"diskCount"`
	HealthErrors            []HealthErrorResponse     `pulumi:"healthErrors"`
	Id                      *string                   `pulumi:"id"`
	IpAddress               *string                   `pulumi:"ipAddress"`
	LastHeartbeat           *string                   `pulumi:"lastHeartbeat"`
	MarsAgentExpiryDate     *string                   `pulumi:"marsAgentExpiryDate"`
	MarsAgentVersion        *string                   `pulumi:"marsAgentVersion"`
	MarsAgentVersionDetails *VersionDetailsResponse   `pulumi:"marsAgentVersionDetails"`
	Name                    *string                   `pulumi:"name"`
	OsType                  *string                   `pulumi:"osType"`
	OsVersion               *string                   `pulumi:"osVersion"`
	RetentionVolumes        []RetentionVolumeResponse `pulumi:"retentionVolumes"`
	ValidationErrors        []HealthErrorResponse     `pulumi:"validationErrors"`
	VersionStatus           *string                   `pulumi:"versionStatus"`
}

type MigrationItemPropertiesResponse struct {
	AllowedOperations           []string                           `pulumi:"allowedOperations"`
	CurrentJob                  CurrentJobDetailsResponse          `pulumi:"currentJob"`
	Health                      string                             `pulumi:"health"`
	HealthErrors                []HealthErrorResponse              `pulumi:"healthErrors"`
	MachineName                 string                             `pulumi:"machineName"`
	MigrationState              string                             `pulumi:"migrationState"`
	MigrationStateDescription   string                             `pulumi:"migrationStateDescription"`
	PolicyFriendlyName          string                             `pulumi:"policyFriendlyName"`
	PolicyId                    string                             `pulumi:"policyId"`
	ProviderSpecificDetails     *VMwareCbtMigrationDetailsResponse `pulumi:"providerSpecificDetails"`
	RecoveryServicesProviderId  string                             `pulumi:"recoveryServicesProviderId"`
	TestMigrateState            string                             `pulumi:"testMigrateState"`
	TestMigrateStateDescription string                             `pulumi:"testMigrateStateDescription"`
}

type MigrationItemPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MigrationItemPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationItemPropertiesResponse)(nil)).Elem()
}

func (o MigrationItemPropertiesResponseOutput) ToMigrationItemPropertiesResponseOutput() MigrationItemPropertiesResponseOutput {
	return o
}

func (o MigrationItemPropertiesResponseOutput) ToMigrationItemPropertiesResponseOutputWithContext(ctx context.Context) MigrationItemPropertiesResponseOutput {
	return o
}

func (o MigrationItemPropertiesResponseOutput) AllowedOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) []string { return v.AllowedOperations }).(pulumi.StringArrayOutput)
}

func (o MigrationItemPropertiesResponseOutput) CurrentJob() CurrentJobDetailsResponseOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) CurrentJobDetailsResponse { return v.CurrentJob }).(CurrentJobDetailsResponseOutput)
}

func (o MigrationItemPropertiesResponseOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.Health }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) HealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) []HealthErrorResponse { return v.HealthErrors }).(HealthErrorResponseArrayOutput)
}

func (o MigrationItemPropertiesResponseOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) MigrationState() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.MigrationState }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) MigrationStateDescription() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.MigrationStateDescription }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) PolicyFriendlyName() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.PolicyFriendlyName }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.PolicyId }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) ProviderSpecificDetails() VMwareCbtMigrationDetailsResponsePtrOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) *VMwareCbtMigrationDetailsResponse {
		return v.ProviderSpecificDetails
	}).(VMwareCbtMigrationDetailsResponsePtrOutput)
}

func (o MigrationItemPropertiesResponseOutput) RecoveryServicesProviderId() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.RecoveryServicesProviderId }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) TestMigrateState() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.TestMigrateState }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) TestMigrateStateDescription() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.TestMigrateStateDescription }).(pulumi.StringOutput)
}

type MobilityServiceUpdateResponse struct {
	OsType       *string `pulumi:"osType"`
	RebootStatus *string `pulumi:"rebootStatus"`
	Version      *string `pulumi:"version"`
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

type NetworkMappingPropertiesResponse struct {
	FabricSpecificSettings      interface{} `pulumi:"fabricSpecificSettings"`
	PrimaryFabricFriendlyName   *string     `pulumi:"primaryFabricFriendlyName"`
	PrimaryNetworkFriendlyName  *string     `pulumi:"primaryNetworkFriendlyName"`
	PrimaryNetworkId            *string     `pulumi:"primaryNetworkId"`
	RecoveryFabricArmId         *string     `pulumi:"recoveryFabricArmId"`
	RecoveryFabricFriendlyName  *string     `pulumi:"recoveryFabricFriendlyName"`
	RecoveryNetworkFriendlyName *string     `pulumi:"recoveryNetworkFriendlyName"`
	RecoveryNetworkId           *string     `pulumi:"recoveryNetworkId"`
	State                       *string     `pulumi:"state"`
}

type NetworkMappingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (NetworkMappingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkMappingPropertiesResponse)(nil)).Elem()
}

func (o NetworkMappingPropertiesResponseOutput) ToNetworkMappingPropertiesResponseOutput() NetworkMappingPropertiesResponseOutput {
	return o
}

func (o NetworkMappingPropertiesResponseOutput) ToNetworkMappingPropertiesResponseOutputWithContext(ctx context.Context) NetworkMappingPropertiesResponseOutput {
	return o
}

func (o NetworkMappingPropertiesResponseOutput) FabricSpecificSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v NetworkMappingPropertiesResponse) interface{} { return v.FabricSpecificSettings }).(pulumi.AnyOutput)
}

func (o NetworkMappingPropertiesResponseOutput) PrimaryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkMappingPropertiesResponse) *string { return v.PrimaryFabricFriendlyName }).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponseOutput) PrimaryNetworkFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkMappingPropertiesResponse) *string { return v.PrimaryNetworkFriendlyName }).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponseOutput) PrimaryNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkMappingPropertiesResponse) *string { return v.PrimaryNetworkId }).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponseOutput) RecoveryFabricArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkMappingPropertiesResponse) *string { return v.RecoveryFabricArmId }).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponseOutput) RecoveryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkMappingPropertiesResponse) *string { return v.RecoveryFabricFriendlyName }).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponseOutput) RecoveryNetworkFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkMappingPropertiesResponse) *string { return v.RecoveryNetworkFriendlyName }).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponseOutput) RecoveryNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkMappingPropertiesResponse) *string { return v.RecoveryNetworkId }).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkMappingPropertiesResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type OSDetailsResponse struct {
	OSMajorVersion *string `pulumi:"oSMajorVersion"`
	OSMinorVersion *string `pulumi:"oSMinorVersion"`
	OSVersion      *string `pulumi:"oSVersion"`
	OsEdition      *string `pulumi:"osEdition"`
	OsType         *string `pulumi:"osType"`
	ProductType    *string `pulumi:"productType"`
}

type OSDiskDetailsResponse struct {
	OsType  *string `pulumi:"osType"`
	OsVhdId *string `pulumi:"osVhdId"`
	VhdName *string `pulumi:"vhdName"`
}

type PolicyPropertiesResponse struct {
	FriendlyName            *string     `pulumi:"friendlyName"`
	ProviderSpecificDetails interface{} `pulumi:"providerSpecificDetails"`
}

type PolicyPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PolicyPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyPropertiesResponse)(nil)).Elem()
}

func (o PolicyPropertiesResponseOutput) ToPolicyPropertiesResponseOutput() PolicyPropertiesResponseOutput {
	return o
}

func (o PolicyPropertiesResponseOutput) ToPolicyPropertiesResponseOutputWithContext(ctx context.Context) PolicyPropertiesResponseOutput {
	return o
}

func (o PolicyPropertiesResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyPropertiesResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o PolicyPropertiesResponseOutput) ProviderSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v PolicyPropertiesResponse) interface{} { return v.ProviderSpecificDetails }).(pulumi.AnyOutput)
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

type ProcessServerDetailsResponse struct {
	AvailableMemoryInBytes             float64               `pulumi:"availableMemoryInBytes"`
	AvailableSpaceInBytes              float64               `pulumi:"availableSpaceInBytes"`
	FreeSpacePercentage                float64               `pulumi:"freeSpacePercentage"`
	Health                             string                `pulumi:"health"`
	HealthErrors                       []HealthErrorResponse `pulumi:"healthErrors"`
	HistoricHealth                     string                `pulumi:"historicHealth"`
	Id                                 string                `pulumi:"id"`
	LastHeartbeatUtc                   string                `pulumi:"lastHeartbeatUtc"`
	MemoryUsagePercentage              float64               `pulumi:"memoryUsagePercentage"`
	Name                               string                `pulumi:"name"`
	ProcessorUsagePercentage           float64               `pulumi:"processorUsagePercentage"`
	ThroughputInBytes                  float64               `pulumi:"throughputInBytes"`
	ThroughputUploadPendingDataInBytes float64               `pulumi:"throughputUploadPendingDataInBytes"`
	TotalMemoryInBytes                 float64               `pulumi:"totalMemoryInBytes"`
	TotalSpaceInBytes                  float64               `pulumi:"totalSpaceInBytes"`
	UsedMemoryInBytes                  float64               `pulumi:"usedMemoryInBytes"`
	UsedSpaceInBytes                   float64               `pulumi:"usedSpaceInBytes"`
	Version                            string                `pulumi:"version"`
}

type ProcessServerResponse struct {
	AgentExpiryDate                    *string                         `pulumi:"agentExpiryDate"`
	AgentVersion                       *string                         `pulumi:"agentVersion"`
	AgentVersionDetails                *VersionDetailsResponse         `pulumi:"agentVersionDetails"`
	AvailableMemoryInBytes             *float64                        `pulumi:"availableMemoryInBytes"`
	AvailableSpaceInBytes              *float64                        `pulumi:"availableSpaceInBytes"`
	CpuLoad                            *string                         `pulumi:"cpuLoad"`
	CpuLoadStatus                      *string                         `pulumi:"cpuLoadStatus"`
	FriendlyName                       *string                         `pulumi:"friendlyName"`
	Health                             string                          `pulumi:"health"`
	HealthErrors                       []HealthErrorResponse           `pulumi:"healthErrors"`
	HostId                             *string                         `pulumi:"hostId"`
	Id                                 *string                         `pulumi:"id"`
	IpAddress                          *string                         `pulumi:"ipAddress"`
	LastHeartbeat                      *string                         `pulumi:"lastHeartbeat"`
	MachineCount                       *string                         `pulumi:"machineCount"`
	MarsCommunicationStatus            string                          `pulumi:"marsCommunicationStatus"`
	MarsRegistrationStatus             string                          `pulumi:"marsRegistrationStatus"`
	MemoryUsageStatus                  *string                         `pulumi:"memoryUsageStatus"`
	MobilityServiceUpdates             []MobilityServiceUpdateResponse `pulumi:"mobilityServiceUpdates"`
	OsType                             *string                         `pulumi:"osType"`
	OsVersion                          *string                         `pulumi:"osVersion"`
	PsServiceStatus                    *string                         `pulumi:"psServiceStatus"`
	PsStatsRefreshTime                 string                          `pulumi:"psStatsRefreshTime"`
	ReplicationPairCount               *string                         `pulumi:"replicationPairCount"`
	SpaceUsageStatus                   *string                         `pulumi:"spaceUsageStatus"`
	SslCertExpiryDate                  *string                         `pulumi:"sslCertExpiryDate"`
	SslCertExpiryRemainingDays         *int                            `pulumi:"sslCertExpiryRemainingDays"`
	SystemLoad                         *string                         `pulumi:"systemLoad"`
	SystemLoadStatus                   *string                         `pulumi:"systemLoadStatus"`
	ThroughputInBytes                  float64                         `pulumi:"throughputInBytes"`
	ThroughputInMBps                   float64                         `pulumi:"throughputInMBps"`
	ThroughputStatus                   string                          `pulumi:"throughputStatus"`
	ThroughputUploadPendingDataInBytes float64                         `pulumi:"throughputUploadPendingDataInBytes"`
	TotalMemoryInBytes                 *float64                        `pulumi:"totalMemoryInBytes"`
	TotalSpaceInBytes                  *float64                        `pulumi:"totalSpaceInBytes"`
	VersionStatus                      *string                         `pulumi:"versionStatus"`
}

type ProtectionContainerMappingPropertiesResponse struct {
	Health                                *string               `pulumi:"health"`
	HealthErrorDetails                    []HealthErrorResponse `pulumi:"healthErrorDetails"`
	PolicyFriendlyName                    *string               `pulumi:"policyFriendlyName"`
	PolicyId                              *string               `pulumi:"policyId"`
	ProviderSpecificDetails               interface{}           `pulumi:"providerSpecificDetails"`
	SourceFabricFriendlyName              *string               `pulumi:"sourceFabricFriendlyName"`
	SourceProtectionContainerFriendlyName *string               `pulumi:"sourceProtectionContainerFriendlyName"`
	State                                 *string               `pulumi:"state"`
	TargetFabricFriendlyName              *string               `pulumi:"targetFabricFriendlyName"`
	TargetProtectionContainerFriendlyName *string               `pulumi:"targetProtectionContainerFriendlyName"`
	TargetProtectionContainerId           *string               `pulumi:"targetProtectionContainerId"`
}

type ProtectionContainerMappingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ProtectionContainerMappingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionContainerMappingPropertiesResponse)(nil)).Elem()
}

func (o ProtectionContainerMappingPropertiesResponseOutput) ToProtectionContainerMappingPropertiesResponseOutput() ProtectionContainerMappingPropertiesResponseOutput {
	return o
}

func (o ProtectionContainerMappingPropertiesResponseOutput) ToProtectionContainerMappingPropertiesResponseOutputWithContext(ctx context.Context) ProtectionContainerMappingPropertiesResponseOutput {
	return o
}

func (o ProtectionContainerMappingPropertiesResponseOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) *string { return v.Health }).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponseOutput) HealthErrorDetails() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) []HealthErrorResponse {
		return v.HealthErrorDetails
	}).(HealthErrorResponseArrayOutput)
}

func (o ProtectionContainerMappingPropertiesResponseOutput) PolicyFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) *string { return v.PolicyFriendlyName }).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponseOutput) ProviderSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) interface{} { return v.ProviderSpecificDetails }).(pulumi.AnyOutput)
}

func (o ProtectionContainerMappingPropertiesResponseOutput) SourceFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) *string { return v.SourceFabricFriendlyName }).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponseOutput) SourceProtectionContainerFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) *string {
		return v.SourceProtectionContainerFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponseOutput) TargetFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) *string { return v.TargetFabricFriendlyName }).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponseOutput) TargetProtectionContainerFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) *string {
		return v.TargetProtectionContainerFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponseOutput) TargetProtectionContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) *string { return v.TargetProtectionContainerId }).(pulumi.StringPtrOutput)
}

type PushInstallerDetailsResponse struct {
	Health           string                `pulumi:"health"`
	HealthErrors     []HealthErrorResponse `pulumi:"healthErrors"`
	Id               string                `pulumi:"id"`
	LastHeartbeatUtc string                `pulumi:"lastHeartbeatUtc"`
	Name             string                `pulumi:"name"`
	Version          string                `pulumi:"version"`
}

type RcmAzureMigrationPolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string  `pulumi:"instanceType"`
	MultiVmSyncStatus                 *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int    `pulumi:"recoveryPointThresholdInMinutes"`
}

type RcmProxyDetailsResponse struct {
	Health           string                `pulumi:"health"`
	HealthErrors     []HealthErrorResponse `pulumi:"healthErrors"`
	Id               string                `pulumi:"id"`
	LastHeartbeatUtc string                `pulumi:"lastHeartbeatUtc"`
	Name             string                `pulumi:"name"`
	Version          string                `pulumi:"version"`
}

type RecoveryPlanA2ADetailsResponse struct {
	InstanceType string  `pulumi:"instanceType"`
	PrimaryZone  *string `pulumi:"primaryZone"`
	RecoveryZone *string `pulumi:"recoveryZone"`
}

type RecoveryPlanA2ADetailsResponseOutput struct{ *pulumi.OutputState }

func (RecoveryPlanA2ADetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanA2ADetailsResponse)(nil)).Elem()
}

func (o RecoveryPlanA2ADetailsResponseOutput) ToRecoveryPlanA2ADetailsResponseOutput() RecoveryPlanA2ADetailsResponseOutput {
	return o
}

func (o RecoveryPlanA2ADetailsResponseOutput) ToRecoveryPlanA2ADetailsResponseOutputWithContext(ctx context.Context) RecoveryPlanA2ADetailsResponseOutput {
	return o
}

func (o RecoveryPlanA2ADetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanA2ADetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o RecoveryPlanA2ADetailsResponseOutput) PrimaryZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanA2ADetailsResponse) *string { return v.PrimaryZone }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanA2ADetailsResponseOutput) RecoveryZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanA2ADetailsResponse) *string { return v.RecoveryZone }).(pulumi.StringPtrOutput)
}

type RecoveryPlanA2ADetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (RecoveryPlanA2ADetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanA2ADetailsResponse)(nil)).Elem()
}

func (o RecoveryPlanA2ADetailsResponseArrayOutput) ToRecoveryPlanA2ADetailsResponseArrayOutput() RecoveryPlanA2ADetailsResponseArrayOutput {
	return o
}

func (o RecoveryPlanA2ADetailsResponseArrayOutput) ToRecoveryPlanA2ADetailsResponseArrayOutputWithContext(ctx context.Context) RecoveryPlanA2ADetailsResponseArrayOutput {
	return o
}

func (o RecoveryPlanA2ADetailsResponseArrayOutput) Index(i pulumi.IntInput) RecoveryPlanA2ADetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecoveryPlanA2ADetailsResponse {
		return vs[0].([]RecoveryPlanA2ADetailsResponse)[vs[1].(int)]
	}).(RecoveryPlanA2ADetailsResponseOutput)
}

type RecoveryPlanA2AInput struct {
	InstanceType *string `pulumi:"instanceType"`
	PrimaryZone  *string `pulumi:"primaryZone"`
	RecoveryZone *string `pulumi:"recoveryZone"`
}





type RecoveryPlanA2AInputInput interface {
	pulumi.Input

	ToRecoveryPlanA2AInputOutput() RecoveryPlanA2AInputOutput
	ToRecoveryPlanA2AInputOutputWithContext(context.Context) RecoveryPlanA2AInputOutput
}

type RecoveryPlanA2AInputArgs struct {
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	PrimaryZone  pulumi.StringPtrInput `pulumi:"primaryZone"`
	RecoveryZone pulumi.StringPtrInput `pulumi:"recoveryZone"`
}

func (RecoveryPlanA2AInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanA2AInput)(nil)).Elem()
}

func (i RecoveryPlanA2AInputArgs) ToRecoveryPlanA2AInputOutput() RecoveryPlanA2AInputOutput {
	return i.ToRecoveryPlanA2AInputOutputWithContext(context.Background())
}

func (i RecoveryPlanA2AInputArgs) ToRecoveryPlanA2AInputOutputWithContext(ctx context.Context) RecoveryPlanA2AInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanA2AInputOutput)
}





type RecoveryPlanA2AInputArrayInput interface {
	pulumi.Input

	ToRecoveryPlanA2AInputArrayOutput() RecoveryPlanA2AInputArrayOutput
	ToRecoveryPlanA2AInputArrayOutputWithContext(context.Context) RecoveryPlanA2AInputArrayOutput
}

type RecoveryPlanA2AInputArray []RecoveryPlanA2AInputInput

func (RecoveryPlanA2AInputArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanA2AInput)(nil)).Elem()
}

func (i RecoveryPlanA2AInputArray) ToRecoveryPlanA2AInputArrayOutput() RecoveryPlanA2AInputArrayOutput {
	return i.ToRecoveryPlanA2AInputArrayOutputWithContext(context.Background())
}

func (i RecoveryPlanA2AInputArray) ToRecoveryPlanA2AInputArrayOutputWithContext(ctx context.Context) RecoveryPlanA2AInputArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanA2AInputArrayOutput)
}

type RecoveryPlanA2AInputOutput struct{ *pulumi.OutputState }

func (RecoveryPlanA2AInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanA2AInput)(nil)).Elem()
}

func (o RecoveryPlanA2AInputOutput) ToRecoveryPlanA2AInputOutput() RecoveryPlanA2AInputOutput {
	return o
}

func (o RecoveryPlanA2AInputOutput) ToRecoveryPlanA2AInputOutputWithContext(ctx context.Context) RecoveryPlanA2AInputOutput {
	return o
}

func (o RecoveryPlanA2AInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanA2AInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanA2AInputOutput) PrimaryZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanA2AInput) *string { return v.PrimaryZone }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanA2AInputOutput) RecoveryZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanA2AInput) *string { return v.RecoveryZone }).(pulumi.StringPtrOutput)
}

type RecoveryPlanA2AInputArrayOutput struct{ *pulumi.OutputState }

func (RecoveryPlanA2AInputArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanA2AInput)(nil)).Elem()
}

func (o RecoveryPlanA2AInputArrayOutput) ToRecoveryPlanA2AInputArrayOutput() RecoveryPlanA2AInputArrayOutput {
	return o
}

func (o RecoveryPlanA2AInputArrayOutput) ToRecoveryPlanA2AInputArrayOutputWithContext(ctx context.Context) RecoveryPlanA2AInputArrayOutput {
	return o
}

func (o RecoveryPlanA2AInputArrayOutput) Index(i pulumi.IntInput) RecoveryPlanA2AInputOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecoveryPlanA2AInput {
		return vs[0].([]RecoveryPlanA2AInput)[vs[1].(int)]
	}).(RecoveryPlanA2AInputOutput)
}

type RecoveryPlanAction struct {
	ActionName         string   `pulumi:"actionName"`
	FailoverDirections []string `pulumi:"failoverDirections"`
	FailoverTypes      []string `pulumi:"failoverTypes"`
}





type RecoveryPlanActionInput interface {
	pulumi.Input

	ToRecoveryPlanActionOutput() RecoveryPlanActionOutput
	ToRecoveryPlanActionOutputWithContext(context.Context) RecoveryPlanActionOutput
}

type RecoveryPlanActionArgs struct {
	ActionName         pulumi.StringInput      `pulumi:"actionName"`
	FailoverDirections pulumi.StringArrayInput `pulumi:"failoverDirections"`
	FailoverTypes      pulumi.StringArrayInput `pulumi:"failoverTypes"`
}

func (RecoveryPlanActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanAction)(nil)).Elem()
}

func (i RecoveryPlanActionArgs) ToRecoveryPlanActionOutput() RecoveryPlanActionOutput {
	return i.ToRecoveryPlanActionOutputWithContext(context.Background())
}

func (i RecoveryPlanActionArgs) ToRecoveryPlanActionOutputWithContext(ctx context.Context) RecoveryPlanActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanActionOutput)
}





type RecoveryPlanActionArrayInput interface {
	pulumi.Input

	ToRecoveryPlanActionArrayOutput() RecoveryPlanActionArrayOutput
	ToRecoveryPlanActionArrayOutputWithContext(context.Context) RecoveryPlanActionArrayOutput
}

type RecoveryPlanActionArray []RecoveryPlanActionInput

func (RecoveryPlanActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanAction)(nil)).Elem()
}

func (i RecoveryPlanActionArray) ToRecoveryPlanActionArrayOutput() RecoveryPlanActionArrayOutput {
	return i.ToRecoveryPlanActionArrayOutputWithContext(context.Background())
}

func (i RecoveryPlanActionArray) ToRecoveryPlanActionArrayOutputWithContext(ctx context.Context) RecoveryPlanActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanActionArrayOutput)
}

type RecoveryPlanActionOutput struct{ *pulumi.OutputState }

func (RecoveryPlanActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanAction)(nil)).Elem()
}

func (o RecoveryPlanActionOutput) ToRecoveryPlanActionOutput() RecoveryPlanActionOutput {
	return o
}

func (o RecoveryPlanActionOutput) ToRecoveryPlanActionOutputWithContext(ctx context.Context) RecoveryPlanActionOutput {
	return o
}

func (o RecoveryPlanActionOutput) ActionName() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanAction) string { return v.ActionName }).(pulumi.StringOutput)
}

func (o RecoveryPlanActionOutput) FailoverDirections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecoveryPlanAction) []string { return v.FailoverDirections }).(pulumi.StringArrayOutput)
}

func (o RecoveryPlanActionOutput) FailoverTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecoveryPlanAction) []string { return v.FailoverTypes }).(pulumi.StringArrayOutput)
}

type RecoveryPlanActionArrayOutput struct{ *pulumi.OutputState }

func (RecoveryPlanActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanAction)(nil)).Elem()
}

func (o RecoveryPlanActionArrayOutput) ToRecoveryPlanActionArrayOutput() RecoveryPlanActionArrayOutput {
	return o
}

func (o RecoveryPlanActionArrayOutput) ToRecoveryPlanActionArrayOutputWithContext(ctx context.Context) RecoveryPlanActionArrayOutput {
	return o
}

func (o RecoveryPlanActionArrayOutput) Index(i pulumi.IntInput) RecoveryPlanActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecoveryPlanAction {
		return vs[0].([]RecoveryPlanAction)[vs[1].(int)]
	}).(RecoveryPlanActionOutput)
}

type RecoveryPlanActionResponse struct {
	ActionName         string      `pulumi:"actionName"`
	CustomDetails      interface{} `pulumi:"customDetails"`
	FailoverDirections []string    `pulumi:"failoverDirections"`
	FailoverTypes      []string    `pulumi:"failoverTypes"`
}

type RecoveryPlanActionResponseOutput struct{ *pulumi.OutputState }

func (RecoveryPlanActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanActionResponse)(nil)).Elem()
}

func (o RecoveryPlanActionResponseOutput) ToRecoveryPlanActionResponseOutput() RecoveryPlanActionResponseOutput {
	return o
}

func (o RecoveryPlanActionResponseOutput) ToRecoveryPlanActionResponseOutputWithContext(ctx context.Context) RecoveryPlanActionResponseOutput {
	return o
}

func (o RecoveryPlanActionResponseOutput) ActionName() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanActionResponse) string { return v.ActionName }).(pulumi.StringOutput)
}

func (o RecoveryPlanActionResponseOutput) CustomDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v RecoveryPlanActionResponse) interface{} { return v.CustomDetails }).(pulumi.AnyOutput)
}

func (o RecoveryPlanActionResponseOutput) FailoverDirections() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecoveryPlanActionResponse) []string { return v.FailoverDirections }).(pulumi.StringArrayOutput)
}

func (o RecoveryPlanActionResponseOutput) FailoverTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecoveryPlanActionResponse) []string { return v.FailoverTypes }).(pulumi.StringArrayOutput)
}

type RecoveryPlanActionResponseArrayOutput struct{ *pulumi.OutputState }

func (RecoveryPlanActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanActionResponse)(nil)).Elem()
}

func (o RecoveryPlanActionResponseArrayOutput) ToRecoveryPlanActionResponseArrayOutput() RecoveryPlanActionResponseArrayOutput {
	return o
}

func (o RecoveryPlanActionResponseArrayOutput) ToRecoveryPlanActionResponseArrayOutputWithContext(ctx context.Context) RecoveryPlanActionResponseArrayOutput {
	return o
}

func (o RecoveryPlanActionResponseArrayOutput) Index(i pulumi.IntInput) RecoveryPlanActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecoveryPlanActionResponse {
		return vs[0].([]RecoveryPlanActionResponse)[vs[1].(int)]
	}).(RecoveryPlanActionResponseOutput)
}

type RecoveryPlanAutomationRunbookActionDetailsResponse struct {
	FabricLocation string  `pulumi:"fabricLocation"`
	InstanceType   string  `pulumi:"instanceType"`
	RunbookId      *string `pulumi:"runbookId"`
	Timeout        *string `pulumi:"timeout"`
}

type RecoveryPlanGroup struct {
	EndGroupActions           []RecoveryPlanAction        `pulumi:"endGroupActions"`
	GroupType                 string                      `pulumi:"groupType"`
	ReplicationProtectedItems []RecoveryPlanProtectedItem `pulumi:"replicationProtectedItems"`
	StartGroupActions         []RecoveryPlanAction        `pulumi:"startGroupActions"`
}





type RecoveryPlanGroupInput interface {
	pulumi.Input

	ToRecoveryPlanGroupOutput() RecoveryPlanGroupOutput
	ToRecoveryPlanGroupOutputWithContext(context.Context) RecoveryPlanGroupOutput
}

type RecoveryPlanGroupArgs struct {
	EndGroupActions           RecoveryPlanActionArrayInput        `pulumi:"endGroupActions"`
	GroupType                 pulumi.StringInput                  `pulumi:"groupType"`
	ReplicationProtectedItems RecoveryPlanProtectedItemArrayInput `pulumi:"replicationProtectedItems"`
	StartGroupActions         RecoveryPlanActionArrayInput        `pulumi:"startGroupActions"`
}

func (RecoveryPlanGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanGroup)(nil)).Elem()
}

func (i RecoveryPlanGroupArgs) ToRecoveryPlanGroupOutput() RecoveryPlanGroupOutput {
	return i.ToRecoveryPlanGroupOutputWithContext(context.Background())
}

func (i RecoveryPlanGroupArgs) ToRecoveryPlanGroupOutputWithContext(ctx context.Context) RecoveryPlanGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanGroupOutput)
}





type RecoveryPlanGroupArrayInput interface {
	pulumi.Input

	ToRecoveryPlanGroupArrayOutput() RecoveryPlanGroupArrayOutput
	ToRecoveryPlanGroupArrayOutputWithContext(context.Context) RecoveryPlanGroupArrayOutput
}

type RecoveryPlanGroupArray []RecoveryPlanGroupInput

func (RecoveryPlanGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanGroup)(nil)).Elem()
}

func (i RecoveryPlanGroupArray) ToRecoveryPlanGroupArrayOutput() RecoveryPlanGroupArrayOutput {
	return i.ToRecoveryPlanGroupArrayOutputWithContext(context.Background())
}

func (i RecoveryPlanGroupArray) ToRecoveryPlanGroupArrayOutputWithContext(ctx context.Context) RecoveryPlanGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanGroupArrayOutput)
}

type RecoveryPlanGroupOutput struct{ *pulumi.OutputState }

func (RecoveryPlanGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanGroup)(nil)).Elem()
}

func (o RecoveryPlanGroupOutput) ToRecoveryPlanGroupOutput() RecoveryPlanGroupOutput {
	return o
}

func (o RecoveryPlanGroupOutput) ToRecoveryPlanGroupOutputWithContext(ctx context.Context) RecoveryPlanGroupOutput {
	return o
}

func (o RecoveryPlanGroupOutput) EndGroupActions() RecoveryPlanActionArrayOutput {
	return o.ApplyT(func(v RecoveryPlanGroup) []RecoveryPlanAction { return v.EndGroupActions }).(RecoveryPlanActionArrayOutput)
}

func (o RecoveryPlanGroupOutput) GroupType() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanGroup) string { return v.GroupType }).(pulumi.StringOutput)
}

func (o RecoveryPlanGroupOutput) ReplicationProtectedItems() RecoveryPlanProtectedItemArrayOutput {
	return o.ApplyT(func(v RecoveryPlanGroup) []RecoveryPlanProtectedItem { return v.ReplicationProtectedItems }).(RecoveryPlanProtectedItemArrayOutput)
}

func (o RecoveryPlanGroupOutput) StartGroupActions() RecoveryPlanActionArrayOutput {
	return o.ApplyT(func(v RecoveryPlanGroup) []RecoveryPlanAction { return v.StartGroupActions }).(RecoveryPlanActionArrayOutput)
}

type RecoveryPlanGroupArrayOutput struct{ *pulumi.OutputState }

func (RecoveryPlanGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanGroup)(nil)).Elem()
}

func (o RecoveryPlanGroupArrayOutput) ToRecoveryPlanGroupArrayOutput() RecoveryPlanGroupArrayOutput {
	return o
}

func (o RecoveryPlanGroupArrayOutput) ToRecoveryPlanGroupArrayOutputWithContext(ctx context.Context) RecoveryPlanGroupArrayOutput {
	return o
}

func (o RecoveryPlanGroupArrayOutput) Index(i pulumi.IntInput) RecoveryPlanGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecoveryPlanGroup {
		return vs[0].([]RecoveryPlanGroup)[vs[1].(int)]
	}).(RecoveryPlanGroupOutput)
}

type RecoveryPlanGroupResponse struct {
	EndGroupActions           []RecoveryPlanActionResponse        `pulumi:"endGroupActions"`
	GroupType                 string                              `pulumi:"groupType"`
	ReplicationProtectedItems []RecoveryPlanProtectedItemResponse `pulumi:"replicationProtectedItems"`
	StartGroupActions         []RecoveryPlanActionResponse        `pulumi:"startGroupActions"`
}

type RecoveryPlanGroupResponseOutput struct{ *pulumi.OutputState }

func (RecoveryPlanGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanGroupResponse)(nil)).Elem()
}

func (o RecoveryPlanGroupResponseOutput) ToRecoveryPlanGroupResponseOutput() RecoveryPlanGroupResponseOutput {
	return o
}

func (o RecoveryPlanGroupResponseOutput) ToRecoveryPlanGroupResponseOutputWithContext(ctx context.Context) RecoveryPlanGroupResponseOutput {
	return o
}

func (o RecoveryPlanGroupResponseOutput) EndGroupActions() RecoveryPlanActionResponseArrayOutput {
	return o.ApplyT(func(v RecoveryPlanGroupResponse) []RecoveryPlanActionResponse { return v.EndGroupActions }).(RecoveryPlanActionResponseArrayOutput)
}

func (o RecoveryPlanGroupResponseOutput) GroupType() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanGroupResponse) string { return v.GroupType }).(pulumi.StringOutput)
}

func (o RecoveryPlanGroupResponseOutput) ReplicationProtectedItems() RecoveryPlanProtectedItemResponseArrayOutput {
	return o.ApplyT(func(v RecoveryPlanGroupResponse) []RecoveryPlanProtectedItemResponse {
		return v.ReplicationProtectedItems
	}).(RecoveryPlanProtectedItemResponseArrayOutput)
}

func (o RecoveryPlanGroupResponseOutput) StartGroupActions() RecoveryPlanActionResponseArrayOutput {
	return o.ApplyT(func(v RecoveryPlanGroupResponse) []RecoveryPlanActionResponse { return v.StartGroupActions }).(RecoveryPlanActionResponseArrayOutput)
}

type RecoveryPlanGroupResponseArrayOutput struct{ *pulumi.OutputState }

func (RecoveryPlanGroupResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanGroupResponse)(nil)).Elem()
}

func (o RecoveryPlanGroupResponseArrayOutput) ToRecoveryPlanGroupResponseArrayOutput() RecoveryPlanGroupResponseArrayOutput {
	return o
}

func (o RecoveryPlanGroupResponseArrayOutput) ToRecoveryPlanGroupResponseArrayOutputWithContext(ctx context.Context) RecoveryPlanGroupResponseArrayOutput {
	return o
}

func (o RecoveryPlanGroupResponseArrayOutput) Index(i pulumi.IntInput) RecoveryPlanGroupResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecoveryPlanGroupResponse {
		return vs[0].([]RecoveryPlanGroupResponse)[vs[1].(int)]
	}).(RecoveryPlanGroupResponseOutput)
}

type RecoveryPlanManualActionDetailsResponse struct {
	Description  *string `pulumi:"description"`
	InstanceType string  `pulumi:"instanceType"`
}

type RecoveryPlanPropertiesResponse struct {
	AllowedOperations                []string                         `pulumi:"allowedOperations"`
	CurrentScenario                  *CurrentScenarioDetailsResponse  `pulumi:"currentScenario"`
	CurrentScenarioStatus            *string                          `pulumi:"currentScenarioStatus"`
	CurrentScenarioStatusDescription *string                          `pulumi:"currentScenarioStatusDescription"`
	FailoverDeploymentModel          *string                          `pulumi:"failoverDeploymentModel"`
	FriendlyName                     *string                          `pulumi:"friendlyName"`
	Groups                           []RecoveryPlanGroupResponse      `pulumi:"groups"`
	LastPlannedFailoverTime          *string                          `pulumi:"lastPlannedFailoverTime"`
	LastTestFailoverTime             *string                          `pulumi:"lastTestFailoverTime"`
	LastUnplannedFailoverTime        *string                          `pulumi:"lastUnplannedFailoverTime"`
	PrimaryFabricFriendlyName        *string                          `pulumi:"primaryFabricFriendlyName"`
	PrimaryFabricId                  *string                          `pulumi:"primaryFabricId"`
	ProviderSpecificDetails          []RecoveryPlanA2ADetailsResponse `pulumi:"providerSpecificDetails"`
	RecoveryFabricFriendlyName       *string                          `pulumi:"recoveryFabricFriendlyName"`
	RecoveryFabricId                 *string                          `pulumi:"recoveryFabricId"`
	ReplicationProviders             []string                         `pulumi:"replicationProviders"`
}

type RecoveryPlanPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RecoveryPlanPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanPropertiesResponse)(nil)).Elem()
}

func (o RecoveryPlanPropertiesResponseOutput) ToRecoveryPlanPropertiesResponseOutput() RecoveryPlanPropertiesResponseOutput {
	return o
}

func (o RecoveryPlanPropertiesResponseOutput) ToRecoveryPlanPropertiesResponseOutputWithContext(ctx context.Context) RecoveryPlanPropertiesResponseOutput {
	return o
}

func (o RecoveryPlanPropertiesResponseOutput) AllowedOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) []string { return v.AllowedOperations }).(pulumi.StringArrayOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) CurrentScenario() CurrentScenarioDetailsResponsePtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *CurrentScenarioDetailsResponse { return v.CurrentScenario }).(CurrentScenarioDetailsResponsePtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) CurrentScenarioStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.CurrentScenarioStatus }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) CurrentScenarioStatusDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.CurrentScenarioStatusDescription }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) FailoverDeploymentModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.FailoverDeploymentModel }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) Groups() RecoveryPlanGroupResponseArrayOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) []RecoveryPlanGroupResponse { return v.Groups }).(RecoveryPlanGroupResponseArrayOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) LastPlannedFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.LastPlannedFailoverTime }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) LastTestFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.LastTestFailoverTime }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) LastUnplannedFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.LastUnplannedFailoverTime }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) PrimaryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.PrimaryFabricFriendlyName }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) PrimaryFabricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.PrimaryFabricId }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) ProviderSpecificDetails() RecoveryPlanA2ADetailsResponseArrayOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) []RecoveryPlanA2ADetailsResponse {
		return v.ProviderSpecificDetails
	}).(RecoveryPlanA2ADetailsResponseArrayOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) RecoveryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.RecoveryFabricFriendlyName }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) RecoveryFabricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) *string { return v.RecoveryFabricId }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponseOutput) ReplicationProviders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecoveryPlanPropertiesResponse) []string { return v.ReplicationProviders }).(pulumi.StringArrayOutput)
}

type RecoveryPlanProtectedItem struct {
	Id               *string `pulumi:"id"`
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}





type RecoveryPlanProtectedItemInput interface {
	pulumi.Input

	ToRecoveryPlanProtectedItemOutput() RecoveryPlanProtectedItemOutput
	ToRecoveryPlanProtectedItemOutputWithContext(context.Context) RecoveryPlanProtectedItemOutput
}

type RecoveryPlanProtectedItemArgs struct {
	Id               pulumi.StringPtrInput `pulumi:"id"`
	VirtualMachineId pulumi.StringPtrInput `pulumi:"virtualMachineId"`
}

func (RecoveryPlanProtectedItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanProtectedItem)(nil)).Elem()
}

func (i RecoveryPlanProtectedItemArgs) ToRecoveryPlanProtectedItemOutput() RecoveryPlanProtectedItemOutput {
	return i.ToRecoveryPlanProtectedItemOutputWithContext(context.Background())
}

func (i RecoveryPlanProtectedItemArgs) ToRecoveryPlanProtectedItemOutputWithContext(ctx context.Context) RecoveryPlanProtectedItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanProtectedItemOutput)
}





type RecoveryPlanProtectedItemArrayInput interface {
	pulumi.Input

	ToRecoveryPlanProtectedItemArrayOutput() RecoveryPlanProtectedItemArrayOutput
	ToRecoveryPlanProtectedItemArrayOutputWithContext(context.Context) RecoveryPlanProtectedItemArrayOutput
}

type RecoveryPlanProtectedItemArray []RecoveryPlanProtectedItemInput

func (RecoveryPlanProtectedItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanProtectedItem)(nil)).Elem()
}

func (i RecoveryPlanProtectedItemArray) ToRecoveryPlanProtectedItemArrayOutput() RecoveryPlanProtectedItemArrayOutput {
	return i.ToRecoveryPlanProtectedItemArrayOutputWithContext(context.Background())
}

func (i RecoveryPlanProtectedItemArray) ToRecoveryPlanProtectedItemArrayOutputWithContext(ctx context.Context) RecoveryPlanProtectedItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanProtectedItemArrayOutput)
}

type RecoveryPlanProtectedItemOutput struct{ *pulumi.OutputState }

func (RecoveryPlanProtectedItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanProtectedItem)(nil)).Elem()
}

func (o RecoveryPlanProtectedItemOutput) ToRecoveryPlanProtectedItemOutput() RecoveryPlanProtectedItemOutput {
	return o
}

func (o RecoveryPlanProtectedItemOutput) ToRecoveryPlanProtectedItemOutputWithContext(ctx context.Context) RecoveryPlanProtectedItemOutput {
	return o
}

func (o RecoveryPlanProtectedItemOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanProtectedItem) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanProtectedItemOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanProtectedItem) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

type RecoveryPlanProtectedItemArrayOutput struct{ *pulumi.OutputState }

func (RecoveryPlanProtectedItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanProtectedItem)(nil)).Elem()
}

func (o RecoveryPlanProtectedItemArrayOutput) ToRecoveryPlanProtectedItemArrayOutput() RecoveryPlanProtectedItemArrayOutput {
	return o
}

func (o RecoveryPlanProtectedItemArrayOutput) ToRecoveryPlanProtectedItemArrayOutputWithContext(ctx context.Context) RecoveryPlanProtectedItemArrayOutput {
	return o
}

func (o RecoveryPlanProtectedItemArrayOutput) Index(i pulumi.IntInput) RecoveryPlanProtectedItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecoveryPlanProtectedItem {
		return vs[0].([]RecoveryPlanProtectedItem)[vs[1].(int)]
	}).(RecoveryPlanProtectedItemOutput)
}

type RecoveryPlanProtectedItemResponse struct {
	Id               *string `pulumi:"id"`
	VirtualMachineId *string `pulumi:"virtualMachineId"`
}

type RecoveryPlanProtectedItemResponseOutput struct{ *pulumi.OutputState }

func (RecoveryPlanProtectedItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanProtectedItemResponse)(nil)).Elem()
}

func (o RecoveryPlanProtectedItemResponseOutput) ToRecoveryPlanProtectedItemResponseOutput() RecoveryPlanProtectedItemResponseOutput {
	return o
}

func (o RecoveryPlanProtectedItemResponseOutput) ToRecoveryPlanProtectedItemResponseOutputWithContext(ctx context.Context) RecoveryPlanProtectedItemResponseOutput {
	return o
}

func (o RecoveryPlanProtectedItemResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanProtectedItemResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanProtectedItemResponseOutput) VirtualMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanProtectedItemResponse) *string { return v.VirtualMachineId }).(pulumi.StringPtrOutput)
}

type RecoveryPlanProtectedItemResponseArrayOutput struct{ *pulumi.OutputState }

func (RecoveryPlanProtectedItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanProtectedItemResponse)(nil)).Elem()
}

func (o RecoveryPlanProtectedItemResponseArrayOutput) ToRecoveryPlanProtectedItemResponseArrayOutput() RecoveryPlanProtectedItemResponseArrayOutput {
	return o
}

func (o RecoveryPlanProtectedItemResponseArrayOutput) ToRecoveryPlanProtectedItemResponseArrayOutputWithContext(ctx context.Context) RecoveryPlanProtectedItemResponseArrayOutput {
	return o
}

func (o RecoveryPlanProtectedItemResponseArrayOutput) Index(i pulumi.IntInput) RecoveryPlanProtectedItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RecoveryPlanProtectedItemResponse {
		return vs[0].([]RecoveryPlanProtectedItemResponse)[vs[1].(int)]
	}).(RecoveryPlanProtectedItemResponseOutput)
}

type RecoveryPlanScriptActionDetailsResponse struct {
	FabricLocation string  `pulumi:"fabricLocation"`
	InstanceType   string  `pulumi:"instanceType"`
	Path           string  `pulumi:"path"`
	Timeout        *string `pulumi:"timeout"`
}

type RecoveryServicesProviderPropertiesResponse struct {
	AllowedScenarios              []string                         `pulumi:"allowedScenarios"`
	AuthenticationIdentityDetails *IdentityProviderDetailsResponse `pulumi:"authenticationIdentityDetails"`
	ConnectionStatus              *string                          `pulumi:"connectionStatus"`
	DraIdentifier                 *string                          `pulumi:"draIdentifier"`
	FabricFriendlyName            *string                          `pulumi:"fabricFriendlyName"`
	FabricType                    *string                          `pulumi:"fabricType"`
	FriendlyName                  *string                          `pulumi:"friendlyName"`
	HealthErrorDetails            []HealthErrorResponse            `pulumi:"healthErrorDetails"`
	LastHeartBeat                 *string                          `pulumi:"lastHeartBeat"`
	ProtectedItemCount            *int                             `pulumi:"protectedItemCount"`
	ProviderVersion               *string                          `pulumi:"providerVersion"`
	ProviderVersionDetails        *VersionDetailsResponse          `pulumi:"providerVersionDetails"`
	ProviderVersionExpiryDate     *string                          `pulumi:"providerVersionExpiryDate"`
	ProviderVersionState          *string                          `pulumi:"providerVersionState"`
	ResourceAccessIdentityDetails *IdentityProviderDetailsResponse `pulumi:"resourceAccessIdentityDetails"`
	ServerVersion                 *string                          `pulumi:"serverVersion"`
}

type RecoveryServicesProviderPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RecoveryServicesProviderPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryServicesProviderPropertiesResponse)(nil)).Elem()
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ToRecoveryServicesProviderPropertiesResponseOutput() RecoveryServicesProviderPropertiesResponseOutput {
	return o
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ToRecoveryServicesProviderPropertiesResponseOutputWithContext(ctx context.Context) RecoveryServicesProviderPropertiesResponseOutput {
	return o
}

func (o RecoveryServicesProviderPropertiesResponseOutput) AllowedScenarios() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) []string { return v.AllowedScenarios }).(pulumi.StringArrayOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) AuthenticationIdentityDetails() IdentityProviderDetailsResponsePtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *IdentityProviderDetailsResponse {
		return v.AuthenticationIdentityDetails
	}).(IdentityProviderDetailsResponsePtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ConnectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.ConnectionStatus }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) DraIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.DraIdentifier }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) FabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.FabricFriendlyName }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) FabricType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.FabricType }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) HealthErrorDetails() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) []HealthErrorResponse { return v.HealthErrorDetails }).(HealthErrorResponseArrayOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) LastHeartBeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.LastHeartBeat }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ProtectedItemCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *int { return v.ProtectedItemCount }).(pulumi.IntPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ProviderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.ProviderVersion }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ProviderVersionDetails() VersionDetailsResponsePtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *VersionDetailsResponse {
		return v.ProviderVersionDetails
	}).(VersionDetailsResponsePtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ProviderVersionExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.ProviderVersionExpiryDate }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ProviderVersionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.ProviderVersionState }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ResourceAccessIdentityDetails() IdentityProviderDetailsResponsePtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *IdentityProviderDetailsResponse {
		return v.ResourceAccessIdentityDetails
	}).(IdentityProviderDetailsResponsePtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.ServerVersion }).(pulumi.StringPtrOutput)
}

type ReplicationAgentDetailsResponse struct {
	Health           string                `pulumi:"health"`
	HealthErrors     []HealthErrorResponse `pulumi:"healthErrors"`
	Id               string                `pulumi:"id"`
	LastHeartbeatUtc string                `pulumi:"lastHeartbeatUtc"`
	Name             string                `pulumi:"name"`
	Version          string                `pulumi:"version"`
}

type ReplicationProtectedItemPropertiesResponse struct {
	ActiveLocation                          *string                         `pulumi:"activeLocation"`
	AllowedOperations                       []string                        `pulumi:"allowedOperations"`
	CurrentScenario                         *CurrentScenarioDetailsResponse `pulumi:"currentScenario"`
	FailoverHealth                          *string                         `pulumi:"failoverHealth"`
	FailoverRecoveryPointId                 *string                         `pulumi:"failoverRecoveryPointId"`
	FriendlyName                            *string                         `pulumi:"friendlyName"`
	HealthErrors                            []HealthErrorResponse           `pulumi:"healthErrors"`
	LastSuccessfulFailoverTime              *string                         `pulumi:"lastSuccessfulFailoverTime"`
	LastSuccessfulTestFailoverTime          *string                         `pulumi:"lastSuccessfulTestFailoverTime"`
	PolicyFriendlyName                      *string                         `pulumi:"policyFriendlyName"`
	PolicyId                                *string                         `pulumi:"policyId"`
	PrimaryFabricFriendlyName               *string                         `pulumi:"primaryFabricFriendlyName"`
	PrimaryFabricProvider                   *string                         `pulumi:"primaryFabricProvider"`
	PrimaryProtectionContainerFriendlyName  *string                         `pulumi:"primaryProtectionContainerFriendlyName"`
	ProtectableItemId                       *string                         `pulumi:"protectableItemId"`
	ProtectedItemType                       *string                         `pulumi:"protectedItemType"`
	ProtectionState                         *string                         `pulumi:"protectionState"`
	ProtectionStateDescription              *string                         `pulumi:"protectionStateDescription"`
	ProviderSpecificDetails                 interface{}                     `pulumi:"providerSpecificDetails"`
	RecoveryContainerId                     *string                         `pulumi:"recoveryContainerId"`
	RecoveryFabricFriendlyName              *string                         `pulumi:"recoveryFabricFriendlyName"`
	RecoveryFabricId                        *string                         `pulumi:"recoveryFabricId"`
	RecoveryProtectionContainerFriendlyName *string                         `pulumi:"recoveryProtectionContainerFriendlyName"`
	RecoveryServicesProviderId              *string                         `pulumi:"recoveryServicesProviderId"`
	ReplicationHealth                       *string                         `pulumi:"replicationHealth"`
	TestFailoverState                       *string                         `pulumi:"testFailoverState"`
	TestFailoverStateDescription            *string                         `pulumi:"testFailoverStateDescription"`
}

type ReplicationProtectedItemPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ReplicationProtectedItemPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationProtectedItemPropertiesResponse)(nil)).Elem()
}

func (o ReplicationProtectedItemPropertiesResponseOutput) ToReplicationProtectedItemPropertiesResponseOutput() ReplicationProtectedItemPropertiesResponseOutput {
	return o
}

func (o ReplicationProtectedItemPropertiesResponseOutput) ToReplicationProtectedItemPropertiesResponseOutputWithContext(ctx context.Context) ReplicationProtectedItemPropertiesResponseOutput {
	return o
}

func (o ReplicationProtectedItemPropertiesResponseOutput) ActiveLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.ActiveLocation }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) AllowedOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) []string { return v.AllowedOperations }).(pulumi.StringArrayOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) CurrentScenario() CurrentScenarioDetailsResponsePtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *CurrentScenarioDetailsResponse {
		return v.CurrentScenario
	}).(CurrentScenarioDetailsResponsePtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) FailoverHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.FailoverHealth }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) FailoverRecoveryPointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.FailoverRecoveryPointId }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) HealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) []HealthErrorResponse { return v.HealthErrors }).(HealthErrorResponseArrayOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) LastSuccessfulFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.LastSuccessfulFailoverTime }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) LastSuccessfulTestFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.LastSuccessfulTestFailoverTime }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) PolicyFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.PolicyFriendlyName }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) PrimaryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.PrimaryFabricFriendlyName }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) PrimaryFabricProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.PrimaryFabricProvider }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) PrimaryProtectionContainerFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string {
		return v.PrimaryProtectionContainerFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) ProtectableItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.ProtectableItemId }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) ProtectedItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.ProtectedItemType }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) ProtectionStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.ProtectionStateDescription }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) ProviderSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) interface{} { return v.ProviderSpecificDetails }).(pulumi.AnyOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) RecoveryContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.RecoveryContainerId }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) RecoveryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.RecoveryFabricFriendlyName }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) RecoveryFabricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.RecoveryFabricId }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) RecoveryProtectionContainerFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string {
		return v.RecoveryProtectionContainerFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) RecoveryServicesProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.RecoveryServicesProviderId }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) ReplicationHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.ReplicationHealth }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) TestFailoverState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.TestFailoverState }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) TestFailoverStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.TestFailoverStateDescription }).(pulumi.StringPtrOutput)
}

type ReprotectAgentDetailsResponse struct {
	Health           string                `pulumi:"health"`
	HealthErrors     []HealthErrorResponse `pulumi:"healthErrors"`
	Id               string                `pulumi:"id"`
	LastHeartbeatUtc string                `pulumi:"lastHeartbeatUtc"`
	Name             string                `pulumi:"name"`
	Version          string                `pulumi:"version"`
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

type RetentionVolumeResponse struct {
	CapacityInBytes     *float64 `pulumi:"capacityInBytes"`
	FreeSpaceInBytes    *float64 `pulumi:"freeSpaceInBytes"`
	ThresholdPercentage *int     `pulumi:"thresholdPercentage"`
	VolumeName          *string  `pulumi:"volumeName"`
}

type RunAsAccountResponse struct {
	AccountId   *string `pulumi:"accountId"`
	AccountName *string `pulumi:"accountName"`
}

type SanEnableProtectionInput struct {
	InstanceType *string `pulumi:"instanceType"`
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

type StorageClassificationMappingPropertiesResponse struct {
	TargetStorageClassificationId *string `pulumi:"targetStorageClassificationId"`
}

type StorageClassificationMappingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageClassificationMappingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageClassificationMappingPropertiesResponse)(nil)).Elem()
}

func (o StorageClassificationMappingPropertiesResponseOutput) ToStorageClassificationMappingPropertiesResponseOutput() StorageClassificationMappingPropertiesResponseOutput {
	return o
}

func (o StorageClassificationMappingPropertiesResponseOutput) ToStorageClassificationMappingPropertiesResponseOutputWithContext(ctx context.Context) StorageClassificationMappingPropertiesResponseOutput {
	return o
}

func (o StorageClassificationMappingPropertiesResponseOutput) TargetStorageClassificationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageClassificationMappingPropertiesResponse) *string { return v.TargetStorageClassificationId }).(pulumi.StringPtrOutput)
}

type StorageMappingInputProperties struct {
	TargetStorageClassificationId *string `pulumi:"targetStorageClassificationId"`
}





type StorageMappingInputPropertiesInput interface {
	pulumi.Input

	ToStorageMappingInputPropertiesOutput() StorageMappingInputPropertiesOutput
	ToStorageMappingInputPropertiesOutputWithContext(context.Context) StorageMappingInputPropertiesOutput
}

type StorageMappingInputPropertiesArgs struct {
	TargetStorageClassificationId pulumi.StringPtrInput `pulumi:"targetStorageClassificationId"`
}

func (StorageMappingInputPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageMappingInputProperties)(nil)).Elem()
}

func (i StorageMappingInputPropertiesArgs) ToStorageMappingInputPropertiesOutput() StorageMappingInputPropertiesOutput {
	return i.ToStorageMappingInputPropertiesOutputWithContext(context.Background())
}

func (i StorageMappingInputPropertiesArgs) ToStorageMappingInputPropertiesOutputWithContext(ctx context.Context) StorageMappingInputPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageMappingInputPropertiesOutput)
}

func (i StorageMappingInputPropertiesArgs) ToStorageMappingInputPropertiesPtrOutput() StorageMappingInputPropertiesPtrOutput {
	return i.ToStorageMappingInputPropertiesPtrOutputWithContext(context.Background())
}

func (i StorageMappingInputPropertiesArgs) ToStorageMappingInputPropertiesPtrOutputWithContext(ctx context.Context) StorageMappingInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageMappingInputPropertiesOutput).ToStorageMappingInputPropertiesPtrOutputWithContext(ctx)
}









type StorageMappingInputPropertiesPtrInput interface {
	pulumi.Input

	ToStorageMappingInputPropertiesPtrOutput() StorageMappingInputPropertiesPtrOutput
	ToStorageMappingInputPropertiesPtrOutputWithContext(context.Context) StorageMappingInputPropertiesPtrOutput
}

type storageMappingInputPropertiesPtrType StorageMappingInputPropertiesArgs

func StorageMappingInputPropertiesPtr(v *StorageMappingInputPropertiesArgs) StorageMappingInputPropertiesPtrInput {
	return (*storageMappingInputPropertiesPtrType)(v)
}

func (*storageMappingInputPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageMappingInputProperties)(nil)).Elem()
}

func (i *storageMappingInputPropertiesPtrType) ToStorageMappingInputPropertiesPtrOutput() StorageMappingInputPropertiesPtrOutput {
	return i.ToStorageMappingInputPropertiesPtrOutputWithContext(context.Background())
}

func (i *storageMappingInputPropertiesPtrType) ToStorageMappingInputPropertiesPtrOutputWithContext(ctx context.Context) StorageMappingInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageMappingInputPropertiesPtrOutput)
}

type StorageMappingInputPropertiesOutput struct{ *pulumi.OutputState }

func (StorageMappingInputPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageMappingInputProperties)(nil)).Elem()
}

func (o StorageMappingInputPropertiesOutput) ToStorageMappingInputPropertiesOutput() StorageMappingInputPropertiesOutput {
	return o
}

func (o StorageMappingInputPropertiesOutput) ToStorageMappingInputPropertiesOutputWithContext(ctx context.Context) StorageMappingInputPropertiesOutput {
	return o
}

func (o StorageMappingInputPropertiesOutput) ToStorageMappingInputPropertiesPtrOutput() StorageMappingInputPropertiesPtrOutput {
	return o.ToStorageMappingInputPropertiesPtrOutputWithContext(context.Background())
}

func (o StorageMappingInputPropertiesOutput) ToStorageMappingInputPropertiesPtrOutputWithContext(ctx context.Context) StorageMappingInputPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageMappingInputProperties) *StorageMappingInputProperties {
		return &v
	}).(StorageMappingInputPropertiesPtrOutput)
}

func (o StorageMappingInputPropertiesOutput) TargetStorageClassificationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageMappingInputProperties) *string { return v.TargetStorageClassificationId }).(pulumi.StringPtrOutput)
}

type StorageMappingInputPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StorageMappingInputPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageMappingInputProperties)(nil)).Elem()
}

func (o StorageMappingInputPropertiesPtrOutput) ToStorageMappingInputPropertiesPtrOutput() StorageMappingInputPropertiesPtrOutput {
	return o
}

func (o StorageMappingInputPropertiesPtrOutput) ToStorageMappingInputPropertiesPtrOutputWithContext(ctx context.Context) StorageMappingInputPropertiesPtrOutput {
	return o
}

func (o StorageMappingInputPropertiesPtrOutput) Elem() StorageMappingInputPropertiesOutput {
	return o.ApplyT(func(v *StorageMappingInputProperties) StorageMappingInputProperties {
		if v != nil {
			return *v
		}
		var ret StorageMappingInputProperties
		return ret
	}).(StorageMappingInputPropertiesOutput)
}

func (o StorageMappingInputPropertiesPtrOutput) TargetStorageClassificationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageMappingInputProperties) *string {
		if v == nil {
			return nil
		}
		return v.TargetStorageClassificationId
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

type VCenterPropertiesResponse struct {
	DiscoveryStatus       *string               `pulumi:"discoveryStatus"`
	FabricArmResourceName *string               `pulumi:"fabricArmResourceName"`
	FriendlyName          *string               `pulumi:"friendlyName"`
	HealthErrors          []HealthErrorResponse `pulumi:"healthErrors"`
	InfrastructureId      *string               `pulumi:"infrastructureId"`
	InternalId            *string               `pulumi:"internalId"`
	IpAddress             *string               `pulumi:"ipAddress"`
	LastHeartbeat         *string               `pulumi:"lastHeartbeat"`
	Port                  *string               `pulumi:"port"`
	ProcessServerId       *string               `pulumi:"processServerId"`
	RunAsAccountId        *string               `pulumi:"runAsAccountId"`
}

type VCenterPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VCenterPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VCenterPropertiesResponse)(nil)).Elem()
}

func (o VCenterPropertiesResponseOutput) ToVCenterPropertiesResponseOutput() VCenterPropertiesResponseOutput {
	return o
}

func (o VCenterPropertiesResponseOutput) ToVCenterPropertiesResponseOutputWithContext(ctx context.Context) VCenterPropertiesResponseOutput {
	return o
}

func (o VCenterPropertiesResponseOutput) DiscoveryStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) *string { return v.DiscoveryStatus }).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponseOutput) FabricArmResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) *string { return v.FabricArmResourceName }).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponseOutput) HealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) []HealthErrorResponse { return v.HealthErrors }).(HealthErrorResponseArrayOutput)
}

func (o VCenterPropertiesResponseOutput) InfrastructureId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) *string { return v.InfrastructureId }).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponseOutput) InternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) *string { return v.InternalId }).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponseOutput) LastHeartbeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) *string { return v.LastHeartbeat }).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponseOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) *string { return v.Port }).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponseOutput) ProcessServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) *string { return v.ProcessServerId }).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponseOutput) RunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VCenterPropertiesResponse) *string { return v.RunAsAccountId }).(pulumi.StringPtrOutput)
}

type VMNicDetailsResponse struct {
	EnableAcceleratedNetworkingOnRecovery *bool              `pulumi:"enableAcceleratedNetworkingOnRecovery"`
	EnableAcceleratedNetworkingOnTfo      *bool              `pulumi:"enableAcceleratedNetworkingOnTfo"`
	IpAddressType                         *string            `pulumi:"ipAddressType"`
	NicId                                 *string            `pulumi:"nicId"`
	PrimaryNicStaticIPAddress             *string            `pulumi:"primaryNicStaticIPAddress"`
	RecoveryLBBackendAddressPoolIds       []string           `pulumi:"recoveryLBBackendAddressPoolIds"`
	RecoveryNetworkSecurityGroupId        *string            `pulumi:"recoveryNetworkSecurityGroupId"`
	RecoveryNicIpAddressType              *string            `pulumi:"recoveryNicIpAddressType"`
	RecoveryNicName                       *string            `pulumi:"recoveryNicName"`
	RecoveryNicResourceGroupName          *string            `pulumi:"recoveryNicResourceGroupName"`
	RecoveryPublicIpAddressId             *string            `pulumi:"recoveryPublicIpAddressId"`
	RecoveryVMNetworkId                   *string            `pulumi:"recoveryVMNetworkId"`
	RecoveryVMSubnetName                  *string            `pulumi:"recoveryVMSubnetName"`
	ReplicaNicId                          *string            `pulumi:"replicaNicId"`
	ReplicaNicStaticIPAddress             *string            `pulumi:"replicaNicStaticIPAddress"`
	ReuseExistingNic                      *bool              `pulumi:"reuseExistingNic"`
	SelectionType                         *string            `pulumi:"selectionType"`
	SourceNicArmId                        *string            `pulumi:"sourceNicArmId"`
	TfoIPConfigs                          []IPConfigResponse `pulumi:"tfoIPConfigs"`
	TfoNetworkSecurityGroupId             *string            `pulumi:"tfoNetworkSecurityGroupId"`
	TfoRecoveryNicName                    *string            `pulumi:"tfoRecoveryNicName"`
	TfoRecoveryNicResourceGroupName       *string            `pulumi:"tfoRecoveryNicResourceGroupName"`
	TfoReuseExistingNic                   *bool              `pulumi:"tfoReuseExistingNic"`
	TfoVMNetworkId                        *string            `pulumi:"tfoVMNetworkId"`
	TfoVMSubnetName                       *string            `pulumi:"tfoVMSubnetName"`
	VMNetworkName                         *string            `pulumi:"vMNetworkName"`
	VMSubnetName                          *string            `pulumi:"vMSubnetName"`
}


func (val *VMNicDetailsResponse) Defaults() *VMNicDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ReuseExistingNic) {
		reuseExistingNic_ := false
		tmp.ReuseExistingNic = &reuseExistingNic_
	}
	if isZero(tmp.TfoReuseExistingNic) {
		tfoReuseExistingNic_ := false
		tmp.TfoReuseExistingNic = &tfoReuseExistingNic_
	}
	return &tmp
}

type VMwareCbtContainerMappingInput struct {
	InstanceType                         *string `pulumi:"instanceType"`
	KeyVaultId                           string  `pulumi:"keyVaultId"`
	KeyVaultUri                          string  `pulumi:"keyVaultUri"`
	ServiceBusConnectionStringSecretName string  `pulumi:"serviceBusConnectionStringSecretName"`
	StorageAccountId                     string  `pulumi:"storageAccountId"`
	StorageAccountSasSecretName          string  `pulumi:"storageAccountSasSecretName"`
	TargetLocation                       string  `pulumi:"targetLocation"`
}

type VMwareCbtDiskInput struct {
	DiskId                         string  `pulumi:"diskId"`
	DiskType                       *string `pulumi:"diskType"`
	IsOSDisk                       string  `pulumi:"isOSDisk"`
	LogStorageAccountId            string  `pulumi:"logStorageAccountId"`
	LogStorageAccountSasSecretName string  `pulumi:"logStorageAccountSasSecretName"`
}





type VMwareCbtDiskInputInput interface {
	pulumi.Input

	ToVMwareCbtDiskInputOutput() VMwareCbtDiskInputOutput
	ToVMwareCbtDiskInputOutputWithContext(context.Context) VMwareCbtDiskInputOutput
}

type VMwareCbtDiskInputArgs struct {
	DiskId                         pulumi.StringInput    `pulumi:"diskId"`
	DiskType                       pulumi.StringPtrInput `pulumi:"diskType"`
	IsOSDisk                       pulumi.StringInput    `pulumi:"isOSDisk"`
	LogStorageAccountId            pulumi.StringInput    `pulumi:"logStorageAccountId"`
	LogStorageAccountSasSecretName pulumi.StringInput    `pulumi:"logStorageAccountSasSecretName"`
}

func (VMwareCbtDiskInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtDiskInput)(nil)).Elem()
}

func (i VMwareCbtDiskInputArgs) ToVMwareCbtDiskInputOutput() VMwareCbtDiskInputOutput {
	return i.ToVMwareCbtDiskInputOutputWithContext(context.Background())
}

func (i VMwareCbtDiskInputArgs) ToVMwareCbtDiskInputOutputWithContext(ctx context.Context) VMwareCbtDiskInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtDiskInputOutput)
}





type VMwareCbtDiskInputArrayInput interface {
	pulumi.Input

	ToVMwareCbtDiskInputArrayOutput() VMwareCbtDiskInputArrayOutput
	ToVMwareCbtDiskInputArrayOutputWithContext(context.Context) VMwareCbtDiskInputArrayOutput
}

type VMwareCbtDiskInputArray []VMwareCbtDiskInputInput

func (VMwareCbtDiskInputArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMwareCbtDiskInput)(nil)).Elem()
}

func (i VMwareCbtDiskInputArray) ToVMwareCbtDiskInputArrayOutput() VMwareCbtDiskInputArrayOutput {
	return i.ToVMwareCbtDiskInputArrayOutputWithContext(context.Background())
}

func (i VMwareCbtDiskInputArray) ToVMwareCbtDiskInputArrayOutputWithContext(ctx context.Context) VMwareCbtDiskInputArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtDiskInputArrayOutput)
}

type VMwareCbtDiskInputOutput struct{ *pulumi.OutputState }

func (VMwareCbtDiskInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtDiskInput)(nil)).Elem()
}

func (o VMwareCbtDiskInputOutput) ToVMwareCbtDiskInputOutput() VMwareCbtDiskInputOutput {
	return o
}

func (o VMwareCbtDiskInputOutput) ToVMwareCbtDiskInputOutputWithContext(ctx context.Context) VMwareCbtDiskInputOutput {
	return o
}

func (o VMwareCbtDiskInputOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtDiskInput) string { return v.DiskId }).(pulumi.StringOutput)
}

func (o VMwareCbtDiskInputOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtDiskInput) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtDiskInputOutput) IsOSDisk() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtDiskInput) string { return v.IsOSDisk }).(pulumi.StringOutput)
}

func (o VMwareCbtDiskInputOutput) LogStorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtDiskInput) string { return v.LogStorageAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtDiskInputOutput) LogStorageAccountSasSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtDiskInput) string { return v.LogStorageAccountSasSecretName }).(pulumi.StringOutput)
}

type VMwareCbtDiskInputArrayOutput struct{ *pulumi.OutputState }

func (VMwareCbtDiskInputArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMwareCbtDiskInput)(nil)).Elem()
}

func (o VMwareCbtDiskInputArrayOutput) ToVMwareCbtDiskInputArrayOutput() VMwareCbtDiskInputArrayOutput {
	return o
}

func (o VMwareCbtDiskInputArrayOutput) ToVMwareCbtDiskInputArrayOutputWithContext(ctx context.Context) VMwareCbtDiskInputArrayOutput {
	return o
}

func (o VMwareCbtDiskInputArrayOutput) Index(i pulumi.IntInput) VMwareCbtDiskInputOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMwareCbtDiskInput {
		return vs[0].([]VMwareCbtDiskInput)[vs[1].(int)]
	}).(VMwareCbtDiskInputOutput)
}

type VMwareCbtEnableMigrationInput struct {
	DataMoverRunAsAccountId               string               `pulumi:"dataMoverRunAsAccountId"`
	DisksToInclude                        []VMwareCbtDiskInput `pulumi:"disksToInclude"`
	InstanceType                          string               `pulumi:"instanceType"`
	LicenseType                           *string              `pulumi:"licenseType"`
	SnapshotRunAsAccountId                string               `pulumi:"snapshotRunAsAccountId"`
	TargetAvailabilitySetId               *string              `pulumi:"targetAvailabilitySetId"`
	TargetBootDiagnosticsStorageAccountId *string              `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetNetworkId                       string               `pulumi:"targetNetworkId"`
	TargetResourceGroupId                 string               `pulumi:"targetResourceGroupId"`
	TargetSubnetName                      *string              `pulumi:"targetSubnetName"`
	TargetVmName                          *string              `pulumi:"targetVmName"`
	TargetVmSize                          *string              `pulumi:"targetVmSize"`
	VmwareMachineId                       string               `pulumi:"vmwareMachineId"`
}





type VMwareCbtEnableMigrationInputInput interface {
	pulumi.Input

	ToVMwareCbtEnableMigrationInputOutput() VMwareCbtEnableMigrationInputOutput
	ToVMwareCbtEnableMigrationInputOutputWithContext(context.Context) VMwareCbtEnableMigrationInputOutput
}

type VMwareCbtEnableMigrationInputArgs struct {
	DataMoverRunAsAccountId               pulumi.StringInput           `pulumi:"dataMoverRunAsAccountId"`
	DisksToInclude                        VMwareCbtDiskInputArrayInput `pulumi:"disksToInclude"`
	InstanceType                          pulumi.StringInput           `pulumi:"instanceType"`
	LicenseType                           pulumi.StringPtrInput        `pulumi:"licenseType"`
	SnapshotRunAsAccountId                pulumi.StringInput           `pulumi:"snapshotRunAsAccountId"`
	TargetAvailabilitySetId               pulumi.StringPtrInput        `pulumi:"targetAvailabilitySetId"`
	TargetBootDiagnosticsStorageAccountId pulumi.StringPtrInput        `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetNetworkId                       pulumi.StringInput           `pulumi:"targetNetworkId"`
	TargetResourceGroupId                 pulumi.StringInput           `pulumi:"targetResourceGroupId"`
	TargetSubnetName                      pulumi.StringPtrInput        `pulumi:"targetSubnetName"`
	TargetVmName                          pulumi.StringPtrInput        `pulumi:"targetVmName"`
	TargetVmSize                          pulumi.StringPtrInput        `pulumi:"targetVmSize"`
	VmwareMachineId                       pulumi.StringInput           `pulumi:"vmwareMachineId"`
}

func (VMwareCbtEnableMigrationInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtEnableMigrationInput)(nil)).Elem()
}

func (i VMwareCbtEnableMigrationInputArgs) ToVMwareCbtEnableMigrationInputOutput() VMwareCbtEnableMigrationInputOutput {
	return i.ToVMwareCbtEnableMigrationInputOutputWithContext(context.Background())
}

func (i VMwareCbtEnableMigrationInputArgs) ToVMwareCbtEnableMigrationInputOutputWithContext(ctx context.Context) VMwareCbtEnableMigrationInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtEnableMigrationInputOutput)
}

type VMwareCbtEnableMigrationInputOutput struct{ *pulumi.OutputState }

func (VMwareCbtEnableMigrationInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtEnableMigrationInput)(nil)).Elem()
}

func (o VMwareCbtEnableMigrationInputOutput) ToVMwareCbtEnableMigrationInputOutput() VMwareCbtEnableMigrationInputOutput {
	return o
}

func (o VMwareCbtEnableMigrationInputOutput) ToVMwareCbtEnableMigrationInputOutputWithContext(ctx context.Context) VMwareCbtEnableMigrationInputOutput {
	return o
}

func (o VMwareCbtEnableMigrationInputOutput) DataMoverRunAsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) string { return v.DataMoverRunAsAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) DisksToInclude() VMwareCbtDiskInputArrayOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) []VMwareCbtDiskInput { return v.DisksToInclude }).(VMwareCbtDiskInputArrayOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) SnapshotRunAsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) string { return v.SnapshotRunAsAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TargetAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetBootDiagnosticsStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TargetBootDiagnosticsStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) string { return v.TargetNetworkId }).(pulumi.StringOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) string { return v.TargetResourceGroupId }).(pulumi.StringOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TargetSubnetName }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetVmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TargetVmName }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TargetVmSize }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) VmwareMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) string { return v.VmwareMachineId }).(pulumi.StringOutput)
}

type VMwareCbtMigrationDetailsResponse struct {
	DataMoverRunAsAccountId               string                                  `pulumi:"dataMoverRunAsAccountId"`
	InstanceType                          string                                  `pulumi:"instanceType"`
	LastRecoveryPointReceived             string                                  `pulumi:"lastRecoveryPointReceived"`
	LicenseType                           *string                                 `pulumi:"licenseType"`
	MigrationRecoveryPointId              string                                  `pulumi:"migrationRecoveryPointId"`
	OsType                                string                                  `pulumi:"osType"`
	ProtectedDisks                        []VMwareCbtProtectedDiskDetailsResponse `pulumi:"protectedDisks"`
	SnapshotRunAsAccountId                string                                  `pulumi:"snapshotRunAsAccountId"`
	TargetAvailabilitySetId               *string                                 `pulumi:"targetAvailabilitySetId"`
	TargetBootDiagnosticsStorageAccountId *string                                 `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetLocation                        string                                  `pulumi:"targetLocation"`
	TargetNetworkId                       *string                                 `pulumi:"targetNetworkId"`
	TargetResourceGroupId                 *string                                 `pulumi:"targetResourceGroupId"`
	TargetVmName                          *string                                 `pulumi:"targetVmName"`
	TargetVmSize                          *string                                 `pulumi:"targetVmSize"`
	VmNics                                []VMwareCbtNicDetailsResponse           `pulumi:"vmNics"`
	VmwareMachineId                       string                                  `pulumi:"vmwareMachineId"`
}

type VMwareCbtMigrationDetailsResponseOutput struct{ *pulumi.OutputState }

func (VMwareCbtMigrationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtMigrationDetailsResponse)(nil)).Elem()
}

func (o VMwareCbtMigrationDetailsResponseOutput) ToVMwareCbtMigrationDetailsResponseOutput() VMwareCbtMigrationDetailsResponseOutput {
	return o
}

func (o VMwareCbtMigrationDetailsResponseOutput) ToVMwareCbtMigrationDetailsResponseOutputWithContext(ctx context.Context) VMwareCbtMigrationDetailsResponseOutput {
	return o
}

func (o VMwareCbtMigrationDetailsResponseOutput) DataMoverRunAsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.DataMoverRunAsAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) LastRecoveryPointReceived() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.LastRecoveryPointReceived }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) MigrationRecoveryPointId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.MigrationRecoveryPointId }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) ProtectedDisks() VMwareCbtProtectedDiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) []VMwareCbtProtectedDiskDetailsResponse {
		return v.ProtectedDisks
	}).(VMwareCbtProtectedDiskDetailsResponseArrayOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) SnapshotRunAsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.SnapshotRunAsAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetBootDiagnosticsStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetBootDiagnosticsStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetLocation() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.TargetLocation }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetNetworkId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetVmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetVmName }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetVmSize }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) VmNics() VMwareCbtNicDetailsResponseArrayOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) []VMwareCbtNicDetailsResponse { return v.VmNics }).(VMwareCbtNicDetailsResponseArrayOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) VmwareMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.VmwareMachineId }).(pulumi.StringOutput)
}

type VMwareCbtMigrationDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (VMwareCbtMigrationDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMwareCbtMigrationDetailsResponse)(nil)).Elem()
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) ToVMwareCbtMigrationDetailsResponsePtrOutput() VMwareCbtMigrationDetailsResponsePtrOutput {
	return o
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) ToVMwareCbtMigrationDetailsResponsePtrOutputWithContext(ctx context.Context) VMwareCbtMigrationDetailsResponsePtrOutput {
	return o
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) Elem() VMwareCbtMigrationDetailsResponseOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) VMwareCbtMigrationDetailsResponse {
		if v != nil {
			return *v
		}
		var ret VMwareCbtMigrationDetailsResponse
		return ret
	}).(VMwareCbtMigrationDetailsResponseOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) DataMoverRunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DataMoverRunAsAccountId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InstanceType
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) LastRecoveryPointReceived() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastRecoveryPointReceived
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) MigrationRecoveryPointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MigrationRecoveryPointId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) ProtectedDisks() VMwareCbtProtectedDiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) []VMwareCbtProtectedDiskDetailsResponse {
		if v == nil {
			return nil
		}
		return v.ProtectedDisks
	}).(VMwareCbtProtectedDiskDetailsResponseArrayOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) SnapshotRunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SnapshotRunAsAccountId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetAvailabilitySetId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetBootDiagnosticsStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetBootDiagnosticsStorageAccountId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TargetLocation
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetNetworkId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetResourceGroupId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetVmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetVmName
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetVmSize
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) VmNics() VMwareCbtNicDetailsResponseArrayOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) []VMwareCbtNicDetailsResponse {
		if v == nil {
			return nil
		}
		return v.VmNics
	}).(VMwareCbtNicDetailsResponseArrayOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) VmwareMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VmwareMachineId
	}).(pulumi.StringPtrOutput)
}

type VMwareCbtNicDetailsResponse struct {
	IsPrimaryNic           *string `pulumi:"isPrimaryNic"`
	IsSelectedForMigration *string `pulumi:"isSelectedForMigration"`
	NicId                  string  `pulumi:"nicId"`
	SourceIPAddress        string  `pulumi:"sourceIPAddress"`
	SourceIPAddressType    string  `pulumi:"sourceIPAddressType"`
	SourceNetworkId        string  `pulumi:"sourceNetworkId"`
	TargetIPAddress        *string `pulumi:"targetIPAddress"`
	TargetIPAddressType    *string `pulumi:"targetIPAddressType"`
	TargetSubnetName       *string `pulumi:"targetSubnetName"`
}

type VMwareCbtNicDetailsResponseOutput struct{ *pulumi.OutputState }

func (VMwareCbtNicDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtNicDetailsResponse)(nil)).Elem()
}

func (o VMwareCbtNicDetailsResponseOutput) ToVMwareCbtNicDetailsResponseOutput() VMwareCbtNicDetailsResponseOutput {
	return o
}

func (o VMwareCbtNicDetailsResponseOutput) ToVMwareCbtNicDetailsResponseOutputWithContext(ctx context.Context) VMwareCbtNicDetailsResponseOutput {
	return o
}

func (o VMwareCbtNicDetailsResponseOutput) IsPrimaryNic() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.IsPrimaryNic }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) IsSelectedForMigration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.IsSelectedForMigration }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) NicId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) string { return v.NicId }).(pulumi.StringOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) SourceIPAddress() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) string { return v.SourceIPAddress }).(pulumi.StringOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) SourceIPAddressType() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) string { return v.SourceIPAddressType }).(pulumi.StringOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) SourceNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) string { return v.SourceNetworkId }).(pulumi.StringOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) TargetIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.TargetIPAddress }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) TargetIPAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.TargetIPAddressType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) TargetSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.TargetSubnetName }).(pulumi.StringPtrOutput)
}

type VMwareCbtNicDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (VMwareCbtNicDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMwareCbtNicDetailsResponse)(nil)).Elem()
}

func (o VMwareCbtNicDetailsResponseArrayOutput) ToVMwareCbtNicDetailsResponseArrayOutput() VMwareCbtNicDetailsResponseArrayOutput {
	return o
}

func (o VMwareCbtNicDetailsResponseArrayOutput) ToVMwareCbtNicDetailsResponseArrayOutputWithContext(ctx context.Context) VMwareCbtNicDetailsResponseArrayOutput {
	return o
}

func (o VMwareCbtNicDetailsResponseArrayOutput) Index(i pulumi.IntInput) VMwareCbtNicDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMwareCbtNicDetailsResponse {
		return vs[0].([]VMwareCbtNicDetailsResponse)[vs[1].(int)]
	}).(VMwareCbtNicDetailsResponseOutput)
}

type VMwareCbtPolicyCreationInput struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      *string `pulumi:"instanceType"`
	RecoveryPointHistoryInMinutes     *int    `pulumi:"recoveryPointHistoryInMinutes"`
}

type VMwareCbtProtectedDiskDetailsResponse struct {
	CapacityInBytes                float64 `pulumi:"capacityInBytes"`
	DiskId                         string  `pulumi:"diskId"`
	DiskName                       string  `pulumi:"diskName"`
	DiskPath                       string  `pulumi:"diskPath"`
	DiskType                       *string `pulumi:"diskType"`
	IsOSDisk                       string  `pulumi:"isOSDisk"`
	LogStorageAccountId            string  `pulumi:"logStorageAccountId"`
	LogStorageAccountSasSecretName string  `pulumi:"logStorageAccountSasSecretName"`
	SeedManagedDiskId              string  `pulumi:"seedManagedDiskId"`
	TargetManagedDiskId            string  `pulumi:"targetManagedDiskId"`
}

type VMwareCbtProtectedDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (VMwareCbtProtectedDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtProtectedDiskDetailsResponse)(nil)).Elem()
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) ToVMwareCbtProtectedDiskDetailsResponseOutput() VMwareCbtProtectedDiskDetailsResponseOutput {
	return o
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) ToVMwareCbtProtectedDiskDetailsResponseOutputWithContext(ctx context.Context) VMwareCbtProtectedDiskDetailsResponseOutput {
	return o
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) CapacityInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) float64 { return v.CapacityInBytes }).(pulumi.Float64Output)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) DiskId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.DiskId }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) DiskName() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.DiskName }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) DiskPath() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.DiskPath }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) IsOSDisk() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.IsOSDisk }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) LogStorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.LogStorageAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) LogStorageAccountSasSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.LogStorageAccountSasSecretName }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) SeedManagedDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.SeedManagedDiskId }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) TargetManagedDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.TargetManagedDiskId }).(pulumi.StringOutput)
}

type VMwareCbtProtectedDiskDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (VMwareCbtProtectedDiskDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMwareCbtProtectedDiskDetailsResponse)(nil)).Elem()
}

func (o VMwareCbtProtectedDiskDetailsResponseArrayOutput) ToVMwareCbtProtectedDiskDetailsResponseArrayOutput() VMwareCbtProtectedDiskDetailsResponseArrayOutput {
	return o
}

func (o VMwareCbtProtectedDiskDetailsResponseArrayOutput) ToVMwareCbtProtectedDiskDetailsResponseArrayOutputWithContext(ctx context.Context) VMwareCbtProtectedDiskDetailsResponseArrayOutput {
	return o
}

func (o VMwareCbtProtectedDiskDetailsResponseArrayOutput) Index(i pulumi.IntInput) VMwareCbtProtectedDiskDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMwareCbtProtectedDiskDetailsResponse {
		return vs[0].([]VMwareCbtProtectedDiskDetailsResponse)[vs[1].(int)]
	}).(VMwareCbtProtectedDiskDetailsResponseOutput)
}

type VMwareCbtProtectionContainerMappingDetailsResponse struct {
	InstanceType                         string `pulumi:"instanceType"`
	KeyVaultId                           string `pulumi:"keyVaultId"`
	KeyVaultUri                          string `pulumi:"keyVaultUri"`
	ServiceBusConnectionStringSecretName string `pulumi:"serviceBusConnectionStringSecretName"`
	StorageAccountId                     string `pulumi:"storageAccountId"`
	StorageAccountSasSecretName          string `pulumi:"storageAccountSasSecretName"`
	TargetLocation                       string `pulumi:"targetLocation"`
}

type VMwareDetailsResponse struct {
	AgentCount                 *string                      `pulumi:"agentCount"`
	AgentExpiryDate            *string                      `pulumi:"agentExpiryDate"`
	AgentVersion               *string                      `pulumi:"agentVersion"`
	AgentVersionDetails        *VersionDetailsResponse      `pulumi:"agentVersionDetails"`
	AvailableMemoryInBytes     *float64                     `pulumi:"availableMemoryInBytes"`
	AvailableSpaceInBytes      *float64                     `pulumi:"availableSpaceInBytes"`
	CpuLoad                    *string                      `pulumi:"cpuLoad"`
	CpuLoadStatus              *string                      `pulumi:"cpuLoadStatus"`
	CsServiceStatus            *string                      `pulumi:"csServiceStatus"`
	DatabaseServerLoad         *string                      `pulumi:"databaseServerLoad"`
	DatabaseServerLoadStatus   *string                      `pulumi:"databaseServerLoadStatus"`
	HostName                   *string                      `pulumi:"hostName"`
	InstanceType               string                       `pulumi:"instanceType"`
	IpAddress                  *string                      `pulumi:"ipAddress"`
	LastHeartbeat              *string                      `pulumi:"lastHeartbeat"`
	MasterTargetServers        []MasterTargetServerResponse `pulumi:"masterTargetServers"`
	MemoryUsageStatus          *string                      `pulumi:"memoryUsageStatus"`
	ProcessServerCount         *string                      `pulumi:"processServerCount"`
	ProcessServers             []ProcessServerResponse      `pulumi:"processServers"`
	ProtectedServers           *string                      `pulumi:"protectedServers"`
	PsTemplateVersion          *string                      `pulumi:"psTemplateVersion"`
	ReplicationPairCount       *string                      `pulumi:"replicationPairCount"`
	RunAsAccounts              []RunAsAccountResponse       `pulumi:"runAsAccounts"`
	SpaceUsageStatus           *string                      `pulumi:"spaceUsageStatus"`
	SslCertExpiryDate          *string                      `pulumi:"sslCertExpiryDate"`
	SslCertExpiryRemainingDays *int                         `pulumi:"sslCertExpiryRemainingDays"`
	SystemLoad                 *string                      `pulumi:"systemLoad"`
	SystemLoadStatus           *string                      `pulumi:"systemLoadStatus"`
	TotalMemoryInBytes         *float64                     `pulumi:"totalMemoryInBytes"`
	TotalSpaceInBytes          *float64                     `pulumi:"totalSpaceInBytes"`
	VersionStatus              *string                      `pulumi:"versionStatus"`
	WebLoad                    *string                      `pulumi:"webLoad"`
	WebLoadStatus              *string                      `pulumi:"webLoadStatus"`
}

type VMwareV2FabricCreationInput struct {
	InstanceType        *string `pulumi:"instanceType"`
	MigrationSolutionId string  `pulumi:"migrationSolutionId"`
	VmwareSiteId        string  `pulumi:"vmwareSiteId"`
}

type VMwareV2FabricSpecificDetailsResponse struct {
	InstanceType        string `pulumi:"instanceType"`
	MigrationSolutionId string `pulumi:"migrationSolutionId"`
	ServiceEndpoint     string `pulumi:"serviceEndpoint"`
	ServiceResourceId   string `pulumi:"serviceResourceId"`
	VmwareSiteId        string `pulumi:"vmwareSiteId"`
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

type VersionDetailsResponse struct {
	ExpiryDate *string `pulumi:"expiryDate"`
	Status     *string `pulumi:"status"`
	Version    *string `pulumi:"version"`
}

type VersionDetailsResponseOutput struct{ *pulumi.OutputState }

func (VersionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VersionDetailsResponse)(nil)).Elem()
}

func (o VersionDetailsResponseOutput) ToVersionDetailsResponseOutput() VersionDetailsResponseOutput {
	return o
}

func (o VersionDetailsResponseOutput) ToVersionDetailsResponseOutputWithContext(ctx context.Context) VersionDetailsResponseOutput {
	return o
}

func (o VersionDetailsResponseOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VersionDetailsResponse) *string { return v.ExpiryDate }).(pulumi.StringPtrOutput)
}

func (o VersionDetailsResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VersionDetailsResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o VersionDetailsResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VersionDetailsResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type VersionDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (VersionDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VersionDetailsResponse)(nil)).Elem()
}

func (o VersionDetailsResponsePtrOutput) ToVersionDetailsResponsePtrOutput() VersionDetailsResponsePtrOutput {
	return o
}

func (o VersionDetailsResponsePtrOutput) ToVersionDetailsResponsePtrOutputWithContext(ctx context.Context) VersionDetailsResponsePtrOutput {
	return o
}

func (o VersionDetailsResponsePtrOutput) Elem() VersionDetailsResponseOutput {
	return o.ApplyT(func(v *VersionDetailsResponse) VersionDetailsResponse {
		if v != nil {
			return *v
		}
		var ret VersionDetailsResponse
		return ret
	}).(VersionDetailsResponseOutput)
}

func (o VersionDetailsResponsePtrOutput) ExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VersionDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExpiryDate
	}).(pulumi.StringPtrOutput)
}

func (o VersionDetailsResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VersionDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o VersionDetailsResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VersionDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type VmmDetailsResponse struct {
	InstanceType string `pulumi:"instanceType"`
}

type VmmToAzureCreateNetworkMappingInput struct {
	InstanceType *string `pulumi:"instanceType"`
}

type VmmToAzureNetworkMappingSettingsResponse struct {
	InstanceType string `pulumi:"instanceType"`
}

type VmmToVmmCreateNetworkMappingInput struct {
	InstanceType *string `pulumi:"instanceType"`
}

type VmmToVmmNetworkMappingSettingsResponse struct {
	InstanceType string `pulumi:"instanceType"`
}

type VmwareCbtPolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   *int   `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int   `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string `pulumi:"instanceType"`
	RecoveryPointHistoryInMinutes     *int   `pulumi:"recoveryPointHistoryInMinutes"`
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

type WorkloadCrrAccessTokenResponse struct {
	AccessTokenString                           *string           `pulumi:"accessTokenString"`
	BMSActiveRegion                             *string           `pulumi:"bMSActiveRegion"`
	BackupManagementType                        *string           `pulumi:"backupManagementType"`
	ContainerId                                 *string           `pulumi:"containerId"`
	ContainerName                               *string           `pulumi:"containerName"`
	ContainerType                               *string           `pulumi:"containerType"`
	CoordinatorServiceStampId                   *string           `pulumi:"coordinatorServiceStampId"`
	CoordinatorServiceStampUri                  *string           `pulumi:"coordinatorServiceStampUri"`
	DatasourceContainerName                     *string           `pulumi:"datasourceContainerName"`
	DatasourceId                                *string           `pulumi:"datasourceId"`
	DatasourceName                              *string           `pulumi:"datasourceName"`
	DatasourceType                              *string           `pulumi:"datasourceType"`
	ObjectType                                  string            `pulumi:"objectType"`
	PolicyId                                    *string           `pulumi:"policyId"`
	PolicyName                                  *string           `pulumi:"policyName"`
	ProtectableObjectContainerHostOsName        *string           `pulumi:"protectableObjectContainerHostOsName"`
	ProtectableObjectFriendlyName               *string           `pulumi:"protectableObjectFriendlyName"`
	ProtectableObjectParentLogicalContainerName *string           `pulumi:"protectableObjectParentLogicalContainerName"`
	ProtectableObjectProtectionState            *string           `pulumi:"protectableObjectProtectionState"`
	ProtectableObjectUniqueName                 *string           `pulumi:"protectableObjectUniqueName"`
	ProtectableObjectWorkloadType               *string           `pulumi:"protectableObjectWorkloadType"`
	ProtectionContainerId                       *float64          `pulumi:"protectionContainerId"`
	ProtectionServiceStampId                    *string           `pulumi:"protectionServiceStampId"`
	ProtectionServiceStampUri                   *string           `pulumi:"protectionServiceStampUri"`
	RecoveryPointId                             *string           `pulumi:"recoveryPointId"`
	RecoveryPointTime                           *string           `pulumi:"recoveryPointTime"`
	ResourceGroupName                           *string           `pulumi:"resourceGroupName"`
	ResourceId                                  *string           `pulumi:"resourceId"`
	ResourceName                                *string           `pulumi:"resourceName"`
	RpIsManagedVirtualMachine                   *bool             `pulumi:"rpIsManagedVirtualMachine"`
	RpOriginalSAOption                          *bool             `pulumi:"rpOriginalSAOption"`
	RpTierInformation                           map[string]string `pulumi:"rpTierInformation"`
	RpVMSizeDescription                         *string           `pulumi:"rpVMSizeDescription"`
	SubscriptionId                              *string           `pulumi:"subscriptionId"`
	TokenExtendedInformation                    *string           `pulumi:"tokenExtendedInformation"`
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
	pulumi.RegisterOutputType(AddRecoveryServicesProviderInputPropertiesOutput{})
	pulumi.RegisterOutputType(AddVCenterRequestPropertiesOutput{})
	pulumi.RegisterOutputType(AddVCenterRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityPtrOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityResponseOutput{})
	pulumi.RegisterOutputType(CmkKekIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CmkKeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CreateNetworkMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreateNetworkMappingInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreatePolicyInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreatePolicyInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreateProtectionContainerMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreateProtectionContainerMappingInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreateRecoveryPlanInputPropertiesOutput{})
	pulumi.RegisterOutputType(CurrentJobDetailsResponseOutput{})
	pulumi.RegisterOutputType(CurrentScenarioDetailsResponseOutput{})
	pulumi.RegisterOutputType(CurrentScenarioDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(EnableMigrationInputPropertiesOutput{})
	pulumi.RegisterOutputType(EnableProtectionInputPropertiesOutput{})
	pulumi.RegisterOutputType(EnableProtectionInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsResponseOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(FabricCreationInputPropertiesOutput{})
	pulumi.RegisterOutputType(FabricCreationInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(FabricPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HealthErrorResponseOutput{})
	pulumi.RegisterOutputType(HealthErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityDataOutput{})
	pulumi.RegisterOutputType(IdentityDataPtrOutput{})
	pulumi.RegisterOutputType(IdentityDataResponseOutput{})
	pulumi.RegisterOutputType(IdentityDataResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderDetailsResponseOutput{})
	pulumi.RegisterOutputType(IdentityProviderDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderInputOutput{})
	pulumi.RegisterOutputType(IdentityProviderInputPtrOutput{})
	pulumi.RegisterOutputType(InnerHealthErrorResponseOutput{})
	pulumi.RegisterOutputType(InnerHealthErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(MigrationItemPropertiesResponseOutput{})
	pulumi.RegisterOutputType(NetworkMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PolicyPropertiesResponseOutput{})
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
	pulumi.RegisterOutputType(ProtectionContainerMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanA2ADetailsResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanA2ADetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanA2AInputOutput{})
	pulumi.RegisterOutputType(RecoveryPlanA2AInputArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanActionOutput{})
	pulumi.RegisterOutputType(RecoveryPlanActionArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanActionResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanActionResponseArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanGroupOutput{})
	pulumi.RegisterOutputType(RecoveryPlanGroupArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanGroupResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanProtectedItemOutput{})
	pulumi.RegisterOutputType(RecoveryPlanProtectedItemArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanProtectedItemResponseArrayOutput{})
	pulumi.RegisterOutputType(RecoveryServicesProviderPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ReplicationProtectedItemPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ResourceGuardOperationDetailResponseOutput{})
	pulumi.RegisterOutputType(ResourceGuardOperationDetailResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceGuardProxyBaseResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageClassificationMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(StorageMappingInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UpgradeDetailsResponseOutput{})
	pulumi.RegisterOutputType(UpgradeDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VCenterPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VMwareCbtDiskInputOutput{})
	pulumi.RegisterOutputType(VMwareCbtDiskInputArrayOutput{})
	pulumi.RegisterOutputType(VMwareCbtEnableMigrationInputOutput{})
	pulumi.RegisterOutputType(VMwareCbtMigrationDetailsResponseOutput{})
	pulumi.RegisterOutputType(VMwareCbtMigrationDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(VMwareCbtNicDetailsResponseOutput{})
	pulumi.RegisterOutputType(VMwareCbtNicDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(VMwareCbtProtectedDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(VMwareCbtProtectedDiskDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(VaultPrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(VaultPrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesOutput{})
	pulumi.RegisterOutputType(VaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesEncryptionOutput{})
	pulumi.RegisterOutputType(VaultPropertiesEncryptionPtrOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseEncryptionOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseEncryptionPtrOutput{})
	pulumi.RegisterOutputType(VersionDetailsResponseOutput{})
	pulumi.RegisterOutputType(VersionDetailsResponsePtrOutput{})
}
