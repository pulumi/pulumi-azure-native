


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type A2AContainerMappingInput struct {
	AgentAutoUpdateStatus               *string `pulumi:"agentAutoUpdateStatus"`
	AutomationAccountArmId              *string `pulumi:"automationAccountArmId"`
	AutomationAccountAuthenticationType *string `pulumi:"automationAccountAuthenticationType"`
	InstanceType                        string  `pulumi:"instanceType"`
}


func (val *A2AContainerMappingInput) Defaults() *A2AContainerMappingInput {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AutomationAccountAuthenticationType) {
		automationAccountAuthenticationType_ := "RunAsAccount"
		tmp.AutomationAccountAuthenticationType = &automationAccountAuthenticationType_
	}
	return &tmp
}

type A2ACrossClusterMigrationEnableProtectionInput struct {
	FabricObjectId      *string `pulumi:"fabricObjectId"`
	InstanceType        string  `pulumi:"instanceType"`
	RecoveryContainerId *string `pulumi:"recoveryContainerId"`
}

type A2ACrossClusterMigrationPolicyCreationInput struct {
	InstanceType string `pulumi:"instanceType"`
}

type A2ACrossClusterMigrationReplicationDetailsResponse struct {
	FabricObjectId               *string `pulumi:"fabricObjectId"`
	InstanceType                 string  `pulumi:"instanceType"`
	LifecycleId                  *string `pulumi:"lifecycleId"`
	OsType                       *string `pulumi:"osType"`
	PrimaryFabricLocation        *string `pulumi:"primaryFabricLocation"`
	VmProtectionState            *string `pulumi:"vmProtectionState"`
	VmProtectionStateDescription *string `pulumi:"vmProtectionStateDescription"`
}

type A2AEnableProtectionInput struct {
	DiskEncryptionInfo                 *DiskEncryptionInfo            `pulumi:"diskEncryptionInfo"`
	FabricObjectId                     string                         `pulumi:"fabricObjectId"`
	InstanceType                       string                         `pulumi:"instanceType"`
	MultiVmGroupId                     *string                        `pulumi:"multiVmGroupId"`
	MultiVmGroupName                   *string                        `pulumi:"multiVmGroupName"`
	RecoveryAvailabilitySetId          *string                        `pulumi:"recoveryAvailabilitySetId"`
	RecoveryAvailabilityZone           *string                        `pulumi:"recoveryAvailabilityZone"`
	RecoveryAzureNetworkId             *string                        `pulumi:"recoveryAzureNetworkId"`
	RecoveryBootDiagStorageAccountId   *string                        `pulumi:"recoveryBootDiagStorageAccountId"`
	RecoveryCapacityReservationGroupId *string                        `pulumi:"recoveryCapacityReservationGroupId"`
	RecoveryCloudServiceId             *string                        `pulumi:"recoveryCloudServiceId"`
	RecoveryContainerId                *string                        `pulumi:"recoveryContainerId"`
	RecoveryExtendedLocation           *ExtendedLocation              `pulumi:"recoveryExtendedLocation"`
	RecoveryProximityPlacementGroupId  *string                        `pulumi:"recoveryProximityPlacementGroupId"`
	RecoveryResourceGroupId            *string                        `pulumi:"recoveryResourceGroupId"`
	RecoverySubnetName                 *string                        `pulumi:"recoverySubnetName"`
	RecoveryVirtualMachineScaleSetId   *string                        `pulumi:"recoveryVirtualMachineScaleSetId"`
	VmDisks                            []A2AVmDiskInputDetails        `pulumi:"vmDisks"`
	VmManagedDisks                     []A2AVmManagedDiskInputDetails `pulumi:"vmManagedDisks"`
}

type A2APolicyCreationInput struct {
	AppConsistentFrequencyInMinutes   *int   `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int   `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string `pulumi:"instanceType"`
	MultiVmSyncStatus                 string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int   `pulumi:"recoveryPointHistory"`
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
	FailoverDiskName                       *string  `pulumi:"failoverDiskName"`
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
	TfoDiskName                            *string  `pulumi:"tfoDiskName"`
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
	PrimaryDiskEncryptionSetId             *string  `pulumi:"primaryDiskEncryptionSetId"`
	PrimaryStagingAzureStorageAccountId    *string  `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryDiskEncryptionSetId            *string  `pulumi:"recoveryDiskEncryptionSetId"`
	RecoveryOrignalTargetDiskId            *string  `pulumi:"recoveryOrignalTargetDiskId"`
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
	AgentAutoUpdateStatus               *string `pulumi:"agentAutoUpdateStatus"`
	AutomationAccountArmId              *string `pulumi:"automationAccountArmId"`
	AutomationAccountAuthenticationType *string `pulumi:"automationAccountAuthenticationType"`
	InstanceType                        string  `pulumi:"instanceType"`
	JobScheduleName                     *string `pulumi:"jobScheduleName"`
	ScheduleName                        *string `pulumi:"scheduleName"`
}


func (val *A2AProtectionContainerMappingDetailsResponse) Defaults() *A2AProtectionContainerMappingDetailsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AutomationAccountAuthenticationType) {
		automationAccountAuthenticationType_ := "RunAsAccount"
		tmp.AutomationAccountAuthenticationType = &automationAccountAuthenticationType_
	}
	return &tmp
}

type A2AReplicationDetailsResponse struct {
	AgentCertificateExpiryDate                  string                                     `pulumi:"agentCertificateExpiryDate"`
	AgentExpiryDate                             *string                                    `pulumi:"agentExpiryDate"`
	AgentVersion                                *string                                    `pulumi:"agentVersion"`
	AutoProtectionOfDataDisk                    *string                                    `pulumi:"autoProtectionOfDataDisk"`
	FabricObjectId                              *string                                    `pulumi:"fabricObjectId"`
	InitialPrimaryExtendedLocation              *ExtendedLocationResponse                  `pulumi:"initialPrimaryExtendedLocation"`
	InitialPrimaryFabricLocation                string                                     `pulumi:"initialPrimaryFabricLocation"`
	InitialPrimaryZone                          string                                     `pulumi:"initialPrimaryZone"`
	InitialRecoveryExtendedLocation             *ExtendedLocationResponse                  `pulumi:"initialRecoveryExtendedLocation"`
	InitialRecoveryFabricLocation               string                                     `pulumi:"initialRecoveryFabricLocation"`
	InitialRecoveryZone                         string                                     `pulumi:"initialRecoveryZone"`
	InstanceType                                string                                     `pulumi:"instanceType"`
	IsReplicationAgentCertificateUpdateRequired *bool                                      `pulumi:"isReplicationAgentCertificateUpdateRequired"`
	IsReplicationAgentUpdateRequired            *bool                                      `pulumi:"isReplicationAgentUpdateRequired"`
	LastHeartbeat                               *string                                    `pulumi:"lastHeartbeat"`
	LastRpoCalculatedTime                       *string                                    `pulumi:"lastRpoCalculatedTime"`
	LifecycleId                                 *string                                    `pulumi:"lifecycleId"`
	ManagementId                                *string                                    `pulumi:"managementId"`
	MonitoringJobType                           *string                                    `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion              *int                                       `pulumi:"monitoringPercentageCompletion"`
	MultiVmGroupCreateOption                    *string                                    `pulumi:"multiVmGroupCreateOption"`
	MultiVmGroupId                              *string                                    `pulumi:"multiVmGroupId"`
	MultiVmGroupName                            *string                                    `pulumi:"multiVmGroupName"`
	OsType                                      *string                                    `pulumi:"osType"`
	PrimaryAvailabilityZone                     *string                                    `pulumi:"primaryAvailabilityZone"`
	PrimaryExtendedLocation                     *ExtendedLocationResponse                  `pulumi:"primaryExtendedLocation"`
	PrimaryFabricLocation                       *string                                    `pulumi:"primaryFabricLocation"`
	ProtectedDisks                              []A2AProtectedDiskDetailsResponse          `pulumi:"protectedDisks"`
	ProtectedManagedDisks                       []A2AProtectedManagedDiskDetailsResponse   `pulumi:"protectedManagedDisks"`
	RecoveryAvailabilitySet                     *string                                    `pulumi:"recoveryAvailabilitySet"`
	RecoveryAvailabilityZone                    *string                                    `pulumi:"recoveryAvailabilityZone"`
	RecoveryAzureGeneration                     string                                     `pulumi:"recoveryAzureGeneration"`
	RecoveryAzureResourceGroupId                *string                                    `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureVMName                         *string                                    `pulumi:"recoveryAzureVMName"`
	RecoveryAzureVMSize                         *string                                    `pulumi:"recoveryAzureVMSize"`
	RecoveryBootDiagStorageAccountId            *string                                    `pulumi:"recoveryBootDiagStorageAccountId"`
	RecoveryCapacityReservationGroupId          *string                                    `pulumi:"recoveryCapacityReservationGroupId"`
	RecoveryCloudService                        *string                                    `pulumi:"recoveryCloudService"`
	RecoveryExtendedLocation                    *ExtendedLocationResponse                  `pulumi:"recoveryExtendedLocation"`
	RecoveryFabricLocation                      *string                                    `pulumi:"recoveryFabricLocation"`
	RecoveryFabricObjectId                      *string                                    `pulumi:"recoveryFabricObjectId"`
	RecoveryProximityPlacementGroupId           *string                                    `pulumi:"recoveryProximityPlacementGroupId"`
	RecoveryVirtualMachineScaleSetId            *string                                    `pulumi:"recoveryVirtualMachineScaleSetId"`
	RpoInSeconds                                *float64                                   `pulumi:"rpoInSeconds"`
	SelectedRecoveryAzureNetworkId              *string                                    `pulumi:"selectedRecoveryAzureNetworkId"`
	SelectedTfoAzureNetworkId                   *string                                    `pulumi:"selectedTfoAzureNetworkId"`
	TestFailoverRecoveryFabricObjectId          *string                                    `pulumi:"testFailoverRecoveryFabricObjectId"`
	TfoAzureVMName                              *string                                    `pulumi:"tfoAzureVMName"`
	UnprotectedDisks                            []A2AUnprotectedDiskDetailsResponse        `pulumi:"unprotectedDisks"`
	VmEncryptionType                            string                                     `pulumi:"vmEncryptionType"`
	VmNics                                      []VMNicDetailsResponse                     `pulumi:"vmNics"`
	VmProtectionState                           *string                                    `pulumi:"vmProtectionState"`
	VmProtectionStateDescription                *string                                    `pulumi:"vmProtectionStateDescription"`
	VmSyncedConfigDetails                       *AzureToAzureVmSyncedConfigDetailsResponse `pulumi:"vmSyncedConfigDetails"`
}

type A2AUnprotectedDiskDetailsResponse struct {
	DiskAutoProtectionStatus *string `pulumi:"diskAutoProtectionStatus"`
	DiskLunId                *int    `pulumi:"diskLunId"`
}

type A2AVmDiskInputDetails struct {
	DiskUri                             string `pulumi:"diskUri"`
	PrimaryStagingAzureStorageAccountId string `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryAzureStorageAccountId       string `pulumi:"recoveryAzureStorageAccountId"`
}

type A2AVmManagedDiskInputDetails struct {
	DiskEncryptionInfo                  *DiskEncryptionInfo `pulumi:"diskEncryptionInfo"`
	DiskId                              string              `pulumi:"diskId"`
	PrimaryStagingAzureStorageAccountId string              `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryDiskEncryptionSetId         *string             `pulumi:"recoveryDiskEncryptionSetId"`
	RecoveryReplicaDiskAccountType      *string             `pulumi:"recoveryReplicaDiskAccountType"`
	RecoveryResourceGroupId             string              `pulumi:"recoveryResourceGroupId"`
	RecoveryTargetDiskAccountType       *string             `pulumi:"recoveryTargetDiskAccountType"`
}

type A2AZoneDetailsResponse struct {
	Source *string `pulumi:"source"`
	Target *string `pulumi:"target"`
}

type AddRecoveryServicesProviderInputProperties struct {
	AuthenticationIdentityInput          IdentityProviderInput  `pulumi:"authenticationIdentityInput"`
	BiosId                               *string                `pulumi:"biosId"`
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
	BiosId                               pulumi.StringPtrInput         `pulumi:"biosId"`
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

func (o AddRecoveryServicesProviderInputPropertiesOutput) BiosId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddRecoveryServicesProviderInputProperties) *string { return v.BiosId }).(pulumi.StringPtrOutput)
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

type AzureFabricCreationInput struct {
	InstanceType string  `pulumi:"instanceType"`
	Location     *string `pulumi:"location"`
}

type AzureFabricSpecificDetailsResponse struct {
	ContainerIds []string                 `pulumi:"containerIds"`
	InstanceType string                   `pulumi:"instanceType"`
	Location     *string                  `pulumi:"location"`
	Zones        []A2AZoneDetailsResponse `pulumi:"zones"`
}

type AzureToAzureCreateNetworkMappingInput struct {
	InstanceType     string `pulumi:"instanceType"`
	PrimaryNetworkId string `pulumi:"primaryNetworkId"`
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

type AzureVmDiskDetailsResponse struct {
	CustomTargetDiskName *string `pulumi:"customTargetDiskName"`
	DiskEncryptionSetId  *string `pulumi:"diskEncryptionSetId"`
	DiskId               *string `pulumi:"diskId"`
	LunId                *string `pulumi:"lunId"`
	MaxSizeMB            *string `pulumi:"maxSizeMB"`
	TargetDiskLocation   *string `pulumi:"targetDiskLocation"`
	TargetDiskName       *string `pulumi:"targetDiskName"`
	VhdId                *string `pulumi:"vhdId"`
	VhdName              *string `pulumi:"vhdName"`
	VhdType              *string `pulumi:"vhdType"`
}

type CreateNetworkMappingInputProperties struct {
	FabricSpecificDetails interface{} `pulumi:"fabricSpecificDetails"`
	RecoveryFabricName    *string     `pulumi:"recoveryFabricName"`
	RecoveryNetworkId     string      `pulumi:"recoveryNetworkId"`
}





type CreateNetworkMappingInputPropertiesInput interface {
	pulumi.Input

	ToCreateNetworkMappingInputPropertiesOutput() CreateNetworkMappingInputPropertiesOutput
	ToCreateNetworkMappingInputPropertiesOutputWithContext(context.Context) CreateNetworkMappingInputPropertiesOutput
}

type CreateNetworkMappingInputPropertiesArgs struct {
	FabricSpecificDetails pulumi.Input          `pulumi:"fabricSpecificDetails"`
	RecoveryFabricName    pulumi.StringPtrInput `pulumi:"recoveryFabricName"`
	RecoveryNetworkId     pulumi.StringInput    `pulumi:"recoveryNetworkId"`
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

func (o CreateNetworkMappingInputPropertiesOutput) FabricSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v CreateNetworkMappingInputProperties) interface{} { return v.FabricSpecificDetails }).(pulumi.AnyOutput)
}

func (o CreateNetworkMappingInputPropertiesOutput) RecoveryFabricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateNetworkMappingInputProperties) *string { return v.RecoveryFabricName }).(pulumi.StringPtrOutput)
}

func (o CreateNetworkMappingInputPropertiesOutput) RecoveryNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v CreateNetworkMappingInputProperties) string { return v.RecoveryNetworkId }).(pulumi.StringOutput)
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

type CriticalJobHistoryDetailsResponse struct {
	JobId     string `pulumi:"jobId"`
	JobName   string `pulumi:"jobName"`
	JobStatus string `pulumi:"jobStatus"`
	StartTime string `pulumi:"startTime"`
}

type CriticalJobHistoryDetailsResponseOutput struct{ *pulumi.OutputState }

func (CriticalJobHistoryDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CriticalJobHistoryDetailsResponse)(nil)).Elem()
}

func (o CriticalJobHistoryDetailsResponseOutput) ToCriticalJobHistoryDetailsResponseOutput() CriticalJobHistoryDetailsResponseOutput {
	return o
}

func (o CriticalJobHistoryDetailsResponseOutput) ToCriticalJobHistoryDetailsResponseOutputWithContext(ctx context.Context) CriticalJobHistoryDetailsResponseOutput {
	return o
}

func (o CriticalJobHistoryDetailsResponseOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v CriticalJobHistoryDetailsResponse) string { return v.JobId }).(pulumi.StringOutput)
}

func (o CriticalJobHistoryDetailsResponseOutput) JobName() pulumi.StringOutput {
	return o.ApplyT(func(v CriticalJobHistoryDetailsResponse) string { return v.JobName }).(pulumi.StringOutput)
}

func (o CriticalJobHistoryDetailsResponseOutput) JobStatus() pulumi.StringOutput {
	return o.ApplyT(func(v CriticalJobHistoryDetailsResponse) string { return v.JobStatus }).(pulumi.StringOutput)
}

func (o CriticalJobHistoryDetailsResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v CriticalJobHistoryDetailsResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

type CriticalJobHistoryDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (CriticalJobHistoryDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CriticalJobHistoryDetailsResponse)(nil)).Elem()
}

func (o CriticalJobHistoryDetailsResponseArrayOutput) ToCriticalJobHistoryDetailsResponseArrayOutput() CriticalJobHistoryDetailsResponseArrayOutput {
	return o
}

func (o CriticalJobHistoryDetailsResponseArrayOutput) ToCriticalJobHistoryDetailsResponseArrayOutputWithContext(ctx context.Context) CriticalJobHistoryDetailsResponseArrayOutput {
	return o
}

func (o CriticalJobHistoryDetailsResponseArrayOutput) Index(i pulumi.IntInput) CriticalJobHistoryDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CriticalJobHistoryDetailsResponse {
		return vs[0].([]CriticalJobHistoryDetailsResponse)[vs[1].(int)]
	}).(CriticalJobHistoryDetailsResponseOutput)
}

type CurrentJobDetailsResponse struct {
	JobId     string `pulumi:"jobId"`
	JobName   string `pulumi:"jobName"`
	StartTime string `pulumi:"startTime"`
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

func (o CurrentJobDetailsResponseOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v CurrentJobDetailsResponse) string { return v.JobId }).(pulumi.StringOutput)
}

func (o CurrentJobDetailsResponseOutput) JobName() pulumi.StringOutput {
	return o.ApplyT(func(v CurrentJobDetailsResponse) string { return v.JobName }).(pulumi.StringOutput)
}

func (o CurrentJobDetailsResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v CurrentJobDetailsResponse) string { return v.StartTime }).(pulumi.StringOutput)
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

type DataStoreResponse struct {
	Capacity     *string `pulumi:"capacity"`
	FreeSpace    *string `pulumi:"freeSpace"`
	SymbolicName *string `pulumi:"symbolicName"`
	Type         *string `pulumi:"type"`
	Uuid         *string `pulumi:"uuid"`
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

type DraDetailsResponse struct {
	BiosId                    string                `pulumi:"biosId"`
	ForwardProtectedItemCount int                   `pulumi:"forwardProtectedItemCount"`
	Health                    string                `pulumi:"health"`
	HealthErrors              []HealthErrorResponse `pulumi:"healthErrors"`
	Id                        string                `pulumi:"id"`
	LastHeartbeatUtc          string                `pulumi:"lastHeartbeatUtc"`
	Name                      string                `pulumi:"name"`
	ReverseProtectedItemCount int                   `pulumi:"reverseProtectedItemCount"`
	Version                   string                `pulumi:"version"`
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

type ExtendedLocation struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ExtendedLocation) string { return v.Name }).(pulumi.StringOutput)
}

func (o ExtendedLocationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ExtendedLocation) string { return v.Type }).(pulumi.StringOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
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

type HyperVHostDetailsResponse struct {
	Id               string `pulumi:"id"`
	MarsAgentVersion string `pulumi:"marsAgentVersion"`
	Name             string `pulumi:"name"`
}

type HyperVReplicaAzureDiskInputDetails struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskId              *string `pulumi:"diskId"`
	DiskType            *string `pulumi:"diskType"`
	LogStorageAccountId *string `pulumi:"logStorageAccountId"`
}

type HyperVReplicaAzureEnableProtectionInput struct {
	DiskEncryptionSetId             *string                              `pulumi:"diskEncryptionSetId"`
	DiskType                        *string                              `pulumi:"diskType"`
	DisksToInclude                  []string                             `pulumi:"disksToInclude"`
	DisksToIncludeForManagedDisks   []HyperVReplicaAzureDiskInputDetails `pulumi:"disksToIncludeForManagedDisks"`
	EnableRdpOnTargetOption         *string                              `pulumi:"enableRdpOnTargetOption"`
	HvHostVmId                      *string                              `pulumi:"hvHostVmId"`
	InstanceType                    string                               `pulumi:"instanceType"`
	LicenseType                     *string                              `pulumi:"licenseType"`
	LogStorageAccountId             *string                              `pulumi:"logStorageAccountId"`
	OsType                          *string                              `pulumi:"osType"`
	SeedManagedDiskTags             map[string]string                    `pulumi:"seedManagedDiskTags"`
	SqlServerLicenseType            *string                              `pulumi:"sqlServerLicenseType"`
	TargetAvailabilitySetId         *string                              `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone          *string                              `pulumi:"targetAvailabilityZone"`
	TargetAzureNetworkId            *string                              `pulumi:"targetAzureNetworkId"`
	TargetAzureSubnetId             *string                              `pulumi:"targetAzureSubnetId"`
	TargetAzureV1ResourceGroupId    *string                              `pulumi:"targetAzureV1ResourceGroupId"`
	TargetAzureV2ResourceGroupId    *string                              `pulumi:"targetAzureV2ResourceGroupId"`
	TargetAzureVmName               *string                              `pulumi:"targetAzureVmName"`
	TargetManagedDiskTags           map[string]string                    `pulumi:"targetManagedDiskTags"`
	TargetNicTags                   map[string]string                    `pulumi:"targetNicTags"`
	TargetProximityPlacementGroupId *string                              `pulumi:"targetProximityPlacementGroupId"`
	TargetStorageAccountId          *string                              `pulumi:"targetStorageAccountId"`
	TargetVmSize                    *string                              `pulumi:"targetVmSize"`
	TargetVmTags                    map[string]string                    `pulumi:"targetVmTags"`
	UseManagedDisks                 *string                              `pulumi:"useManagedDisks"`
	UseManagedDisksForReplication   *string                              `pulumi:"useManagedDisksForReplication"`
	VhdId                           *string                              `pulumi:"vhdId"`
	VmName                          *string                              `pulumi:"vmName"`
}

type HyperVReplicaAzureManagedDiskDetailsResponse struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskId              *string `pulumi:"diskId"`
	ReplicaDiskType     *string `pulumi:"replicaDiskType"`
	SeedManagedDiskId   *string `pulumi:"seedManagedDiskId"`
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
	InstanceType                                  string   `pulumi:"instanceType"`
	OnlineReplicationStartTime                    *string  `pulumi:"onlineReplicationStartTime"`
	RecoveryPointHistoryDuration                  *int     `pulumi:"recoveryPointHistoryDuration"`
	ReplicationInterval                           *int     `pulumi:"replicationInterval"`
	StorageAccounts                               []string `pulumi:"storageAccounts"`
}

type HyperVReplicaAzureReplicationDetailsResponse struct {
	AzureVmDiskDetails               []AzureVmDiskDetailsResponse                   `pulumi:"azureVmDiskDetails"`
	EnableRdpOnTargetOption          *string                                        `pulumi:"enableRdpOnTargetOption"`
	Encryption                       *string                                        `pulumi:"encryption"`
	InitialReplicationDetails        *InitialReplicationDetailsResponse             `pulumi:"initialReplicationDetails"`
	InstanceType                     string                                         `pulumi:"instanceType"`
	LastRecoveryPointReceived        string                                         `pulumi:"lastRecoveryPointReceived"`
	LastReplicatedTime               *string                                        `pulumi:"lastReplicatedTime"`
	LastRpoCalculatedTime            *string                                        `pulumi:"lastRpoCalculatedTime"`
	LicenseType                      *string                                        `pulumi:"licenseType"`
	OSDetails                        *OSDetailsResponse                             `pulumi:"oSDetails"`
	ProtectedManagedDisks            []HyperVReplicaAzureManagedDiskDetailsResponse `pulumi:"protectedManagedDisks"`
	RecoveryAvailabilitySetId        *string                                        `pulumi:"recoveryAvailabilitySetId"`
	RecoveryAzureLogStorageAccountId *string                                        `pulumi:"recoveryAzureLogStorageAccountId"`
	RecoveryAzureResourceGroupId     *string                                        `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureStorageAccount      *string                                        `pulumi:"recoveryAzureStorageAccount"`
	RecoveryAzureVMSize              *string                                        `pulumi:"recoveryAzureVMSize"`
	RecoveryAzureVmName              *string                                        `pulumi:"recoveryAzureVmName"`
	RpoInSeconds                     *float64                                       `pulumi:"rpoInSeconds"`
	SeedManagedDiskTags              map[string]string                              `pulumi:"seedManagedDiskTags"`
	SelectedRecoveryAzureNetworkId   *string                                        `pulumi:"selectedRecoveryAzureNetworkId"`
	SelectedSourceNicId              *string                                        `pulumi:"selectedSourceNicId"`
	SourceVmCpuCount                 *int                                           `pulumi:"sourceVmCpuCount"`
	SourceVmRamSizeInMB              *int                                           `pulumi:"sourceVmRamSizeInMB"`
	SqlServerLicenseType             *string                                        `pulumi:"sqlServerLicenseType"`
	TargetAvailabilityZone           *string                                        `pulumi:"targetAvailabilityZone"`
	TargetManagedDiskTags            map[string]string                              `pulumi:"targetManagedDiskTags"`
	TargetNicTags                    map[string]string                              `pulumi:"targetNicTags"`
	TargetProximityPlacementGroupId  *string                                        `pulumi:"targetProximityPlacementGroupId"`
	TargetVmTags                     map[string]string                              `pulumi:"targetVmTags"`
	UseManagedDisks                  *string                                        `pulumi:"useManagedDisks"`
	VmId                             *string                                        `pulumi:"vmId"`
	VmNics                           []VMNicDetailsResponse                         `pulumi:"vmNics"`
	VmProtectionState                *string                                        `pulumi:"vmProtectionState"`
	VmProtectionStateDescription     *string                                        `pulumi:"vmProtectionStateDescription"`
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
	InstanceType                                  string  `pulumi:"instanceType"`
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
	InstanceType                                  string  `pulumi:"instanceType"`
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
	HyperVHosts  []HyperVHostDetailsResponse `pulumi:"hyperVHosts"`
	InstanceType string                      `pulumi:"instanceType"`
}

type IPConfigDetailsResponse struct {
	IpAddressType                   *string  `pulumi:"ipAddressType"`
	IsPrimary                       *bool    `pulumi:"isPrimary"`
	IsSeletedForFailover            *bool    `pulumi:"isSeletedForFailover"`
	Name                            *string  `pulumi:"name"`
	RecoveryIPAddressType           *string  `pulumi:"recoveryIPAddressType"`
	RecoveryLBBackendAddressPoolIds []string `pulumi:"recoveryLBBackendAddressPoolIds"`
	RecoveryPublicIPAddressId       *string  `pulumi:"recoveryPublicIPAddressId"`
	RecoveryStaticIPAddress         *string  `pulumi:"recoveryStaticIPAddress"`
	RecoverySubnetName              *string  `pulumi:"recoverySubnetName"`
	StaticIPAddress                 *string  `pulumi:"staticIPAddress"`
	SubnetName                      *string  `pulumi:"subnetName"`
	TfoLBBackendAddressPoolIds      []string `pulumi:"tfoLBBackendAddressPoolIds"`
	TfoPublicIPAddressId            *string  `pulumi:"tfoPublicIPAddressId"`
	TfoStaticIPAddress              *string  `pulumi:"tfoStaticIPAddress"`
	TfoSubnetName                   *string  `pulumi:"tfoSubnetName"`
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
	InstanceType                    string                          `pulumi:"instanceType"`
	LicenseType                     *string                         `pulumi:"licenseType"`
	LogStorageAccountId             *string                         `pulumi:"logStorageAccountId"`
	MasterTargetId                  *string                         `pulumi:"masterTargetId"`
	MultiVmGroupId                  *string                         `pulumi:"multiVmGroupId"`
	MultiVmGroupName                *string                         `pulumi:"multiVmGroupName"`
	ProcessServerId                 *string                         `pulumi:"processServerId"`
	RunAsAccountId                  *string                         `pulumi:"runAsAccountId"`
	SeedManagedDiskTags             map[string]string               `pulumi:"seedManagedDiskTags"`
	SqlServerLicenseType            *string                         `pulumi:"sqlServerLicenseType"`
	StorageAccountId                *string                         `pulumi:"storageAccountId"`
	TargetAvailabilitySetId         *string                         `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone          *string                         `pulumi:"targetAvailabilityZone"`
	TargetAzureNetworkId            *string                         `pulumi:"targetAzureNetworkId"`
	TargetAzureSubnetId             *string                         `pulumi:"targetAzureSubnetId"`
	TargetAzureV1ResourceGroupId    *string                         `pulumi:"targetAzureV1ResourceGroupId"`
	TargetAzureV2ResourceGroupId    *string                         `pulumi:"targetAzureV2ResourceGroupId"`
	TargetAzureVmName               *string                         `pulumi:"targetAzureVmName"`
	TargetManagedDiskTags           map[string]string               `pulumi:"targetManagedDiskTags"`
	TargetNicTags                   map[string]string               `pulumi:"targetNicTags"`
	TargetProximityPlacementGroupId *string                         `pulumi:"targetProximityPlacementGroupId"`
	TargetVmSize                    *string                         `pulumi:"targetVmSize"`
	TargetVmTags                    map[string]string               `pulumi:"targetVmTags"`
}

type InMageAzureV2ManagedDiskDetailsResponse struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskId              *string `pulumi:"diskId"`
	ReplicaDiskType     *string `pulumi:"replicaDiskType"`
	SeedManagedDiskId   *string `pulumi:"seedManagedDiskId"`
	TargetDiskName      *string `pulumi:"targetDiskName"`
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
	AppConsistentFrequencyInMinutes   *int   `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int   `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string `pulumi:"instanceType"`
	MultiVmSyncStatus                 string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int   `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int   `pulumi:"recoveryPointThresholdInMinutes"`
}

type InMageAzureV2ProtectedDiskDetailsResponse struct {
	DiskCapacityInBytes                 *float64 `pulumi:"diskCapacityInBytes"`
	DiskId                              *string  `pulumi:"diskId"`
	DiskName                            *string  `pulumi:"diskName"`
	DiskResized                         *string  `pulumi:"diskResized"`
	FileSystemCapacityInBytes           *float64 `pulumi:"fileSystemCapacityInBytes"`
	HealthErrorCode                     *string  `pulumi:"healthErrorCode"`
	LastRpoCalculatedTime               *string  `pulumi:"lastRpoCalculatedTime"`
	ProgressHealth                      *string  `pulumi:"progressHealth"`
	ProgressStatus                      *string  `pulumi:"progressStatus"`
	ProtectionStage                     *string  `pulumi:"protectionStage"`
	PsDataInMegaBytes                   *float64 `pulumi:"psDataInMegaBytes"`
	ResyncDurationInSeconds             *float64 `pulumi:"resyncDurationInSeconds"`
	ResyncLast15MinutesTransferredBytes *float64 `pulumi:"resyncLast15MinutesTransferredBytes"`
	ResyncLastDataTransferTimeUTC       *string  `pulumi:"resyncLastDataTransferTimeUTC"`
	ResyncProcessedBytes                *float64 `pulumi:"resyncProcessedBytes"`
	ResyncProgressPercentage            *int     `pulumi:"resyncProgressPercentage"`
	ResyncRequired                      *string  `pulumi:"resyncRequired"`
	ResyncStartTime                     *string  `pulumi:"resyncStartTime"`
	ResyncTotalTransferredBytes         *float64 `pulumi:"resyncTotalTransferredBytes"`
	RpoInSeconds                        *float64 `pulumi:"rpoInSeconds"`
	SecondsToTakeSwitchProvider         *float64 `pulumi:"secondsToTakeSwitchProvider"`
	SourceDataInMegaBytes               *float64 `pulumi:"sourceDataInMegaBytes"`
	TargetDataInMegaBytes               *float64 `pulumi:"targetDataInMegaBytes"`
}

type InMageAzureV2ReplicationDetailsResponse struct {
	AgentExpiryDate                    *string                                                   `pulumi:"agentExpiryDate"`
	AgentVersion                       *string                                                   `pulumi:"agentVersion"`
	AzureVMDiskDetails                 []AzureVmDiskDetailsResponse                              `pulumi:"azureVMDiskDetails"`
	AzureVmGeneration                  *string                                                   `pulumi:"azureVmGeneration"`
	CompressedDataRateInMB             *float64                                                  `pulumi:"compressedDataRateInMB"`
	Datastores                         []string                                                  `pulumi:"datastores"`
	DiscoveryType                      *string                                                   `pulumi:"discoveryType"`
	DiskResized                        *string                                                   `pulumi:"diskResized"`
	EnableRdpOnTargetOption            *string                                                   `pulumi:"enableRdpOnTargetOption"`
	FirmwareType                       *string                                                   `pulumi:"firmwareType"`
	InfrastructureVmId                 *string                                                   `pulumi:"infrastructureVmId"`
	InstanceType                       string                                                    `pulumi:"instanceType"`
	IpAddress                          *string                                                   `pulumi:"ipAddress"`
	IsAdditionalStatsAvailable         *bool                                                     `pulumi:"isAdditionalStatsAvailable"`
	IsAgentUpdateRequired              *string                                                   `pulumi:"isAgentUpdateRequired"`
	IsRebootAfterUpdateRequired        *string                                                   `pulumi:"isRebootAfterUpdateRequired"`
	LastHeartbeat                      *string                                                   `pulumi:"lastHeartbeat"`
	LastRecoveryPointReceived          string                                                    `pulumi:"lastRecoveryPointReceived"`
	LastRpoCalculatedTime              *string                                                   `pulumi:"lastRpoCalculatedTime"`
	LastUpdateReceivedTime             *string                                                   `pulumi:"lastUpdateReceivedTime"`
	LicenseType                        *string                                                   `pulumi:"licenseType"`
	MasterTargetId                     *string                                                   `pulumi:"masterTargetId"`
	MultiVmGroupId                     *string                                                   `pulumi:"multiVmGroupId"`
	MultiVmGroupName                   *string                                                   `pulumi:"multiVmGroupName"`
	MultiVmSyncStatus                  *string                                                   `pulumi:"multiVmSyncStatus"`
	OsDiskId                           *string                                                   `pulumi:"osDiskId"`
	OsType                             *string                                                   `pulumi:"osType"`
	OsVersion                          *string                                                   `pulumi:"osVersion"`
	ProcessServerId                    *string                                                   `pulumi:"processServerId"`
	ProcessServerName                  *string                                                   `pulumi:"processServerName"`
	ProtectedDisks                     []InMageAzureV2ProtectedDiskDetailsResponse               `pulumi:"protectedDisks"`
	ProtectedManagedDisks              []InMageAzureV2ManagedDiskDetailsResponse                 `pulumi:"protectedManagedDisks"`
	ProtectionStage                    *string                                                   `pulumi:"protectionStage"`
	RecoveryAvailabilitySetId          *string                                                   `pulumi:"recoveryAvailabilitySetId"`
	RecoveryAzureLogStorageAccountId   *string                                                   `pulumi:"recoveryAzureLogStorageAccountId"`
	RecoveryAzureResourceGroupId       *string                                                   `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureStorageAccount        *string                                                   `pulumi:"recoveryAzureStorageAccount"`
	RecoveryAzureVMName                *string                                                   `pulumi:"recoveryAzureVMName"`
	RecoveryAzureVMSize                *string                                                   `pulumi:"recoveryAzureVMSize"`
	ReplicaId                          *string                                                   `pulumi:"replicaId"`
	ResyncProgressPercentage           *int                                                      `pulumi:"resyncProgressPercentage"`
	RpoInSeconds                       *float64                                                  `pulumi:"rpoInSeconds"`
	SeedManagedDiskTags                map[string]string                                         `pulumi:"seedManagedDiskTags"`
	SelectedRecoveryAzureNetworkId     *string                                                   `pulumi:"selectedRecoveryAzureNetworkId"`
	SelectedSourceNicId                *string                                                   `pulumi:"selectedSourceNicId"`
	SelectedTfoAzureNetworkId          *string                                                   `pulumi:"selectedTfoAzureNetworkId"`
	SourceVmCpuCount                   *int                                                      `pulumi:"sourceVmCpuCount"`
	SourceVmRamSizeInMB                *int                                                      `pulumi:"sourceVmRamSizeInMB"`
	SqlServerLicenseType               *string                                                   `pulumi:"sqlServerLicenseType"`
	SwitchProviderBlockingErrorDetails []InMageAzureV2SwitchProviderBlockingErrorDetailsResponse `pulumi:"switchProviderBlockingErrorDetails"`
	SwitchProviderDetails              *InMageAzureV2SwitchProviderDetailsResponse               `pulumi:"switchProviderDetails"`
	TargetAvailabilityZone             *string                                                   `pulumi:"targetAvailabilityZone"`
	TargetManagedDiskTags              map[string]string                                         `pulumi:"targetManagedDiskTags"`
	TargetNicTags                      map[string]string                                         `pulumi:"targetNicTags"`
	TargetProximityPlacementGroupId    *string                                                   `pulumi:"targetProximityPlacementGroupId"`
	TargetVmId                         *string                                                   `pulumi:"targetVmId"`
	TargetVmTags                       map[string]string                                         `pulumi:"targetVmTags"`
	TotalDataTransferred               *float64                                                  `pulumi:"totalDataTransferred"`
	TotalProgressHealth                *string                                                   `pulumi:"totalProgressHealth"`
	UncompressedDataRateInMB           *float64                                                  `pulumi:"uncompressedDataRateInMB"`
	UseManagedDisks                    *string                                                   `pulumi:"useManagedDisks"`
	VCenterInfrastructureId            *string                                                   `pulumi:"vCenterInfrastructureId"`
	ValidationErrors                   []HealthErrorResponse                                     `pulumi:"validationErrors"`
	VhdName                            *string                                                   `pulumi:"vhdName"`
	VmId                               *string                                                   `pulumi:"vmId"`
	VmNics                             []VMNicDetailsResponse                                    `pulumi:"vmNics"`
	VmProtectionState                  *string                                                   `pulumi:"vmProtectionState"`
	VmProtectionStateDescription       *string                                                   `pulumi:"vmProtectionStateDescription"`
}

type InMageAzureV2SwitchProviderBlockingErrorDetailsResponse struct {
	ErrorCode              string            `pulumi:"errorCode"`
	ErrorMessage           string            `pulumi:"errorMessage"`
	ErrorMessageParameters map[string]string `pulumi:"errorMessageParameters"`
	ErrorTags              map[string]string `pulumi:"errorTags"`
	PossibleCauses         string            `pulumi:"possibleCauses"`
	RecommendedAction      string            `pulumi:"recommendedAction"`
}

type InMageAzureV2SwitchProviderDetailsResponse struct {
	TargetApplianceId string `pulumi:"targetApplianceId"`
	TargetFabricId    string `pulumi:"targetFabricId"`
	TargetResourceId  string `pulumi:"targetResourceId"`
	TargetVaultId     string `pulumi:"targetVaultId"`
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
	InstanceType       string                    `pulumi:"instanceType"`
	MasterTargetId     string                    `pulumi:"masterTargetId"`
	MultiVmGroupId     string                    `pulumi:"multiVmGroupId"`
	MultiVmGroupName   string                    `pulumi:"multiVmGroupName"`
	ProcessServerId    string                    `pulumi:"processServerId"`
	RetentionDrive     string                    `pulumi:"retentionDrive"`
	RunAsAccountId     *string                   `pulumi:"runAsAccountId"`
	VmFriendlyName     *string                   `pulumi:"vmFriendlyName"`
}

type InMageFabricSwitchProviderBlockingErrorDetailsResponse struct {
	ErrorCode              string            `pulumi:"errorCode"`
	ErrorMessage           string            `pulumi:"errorMessage"`
	ErrorMessageParameters map[string]string `pulumi:"errorMessageParameters"`
	ErrorTags              map[string]string `pulumi:"errorTags"`
	PossibleCauses         string            `pulumi:"possibleCauses"`
	RecommendedAction      string            `pulumi:"recommendedAction"`
}

type InMagePolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes *int    `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    string  `pulumi:"instanceType"`
	MultiVmSyncStatus               *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes *int    `pulumi:"recoveryPointThresholdInMinutes"`
}

type InMagePolicyInput struct {
	AppConsistentFrequencyInMinutes *int   `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    string `pulumi:"instanceType"`
	MultiVmSyncStatus               string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            *int   `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes *int   `pulumi:"recoveryPointThresholdInMinutes"`
}

type InMageProtectedDiskDetailsResponse struct {
	DiskCapacityInBytes                 *float64 `pulumi:"diskCapacityInBytes"`
	DiskId                              *string  `pulumi:"diskId"`
	DiskName                            *string  `pulumi:"diskName"`
	DiskResized                         *string  `pulumi:"diskResized"`
	FileSystemCapacityInBytes           *float64 `pulumi:"fileSystemCapacityInBytes"`
	HealthErrorCode                     *string  `pulumi:"healthErrorCode"`
	LastRpoCalculatedTime               *string  `pulumi:"lastRpoCalculatedTime"`
	ProgressHealth                      *string  `pulumi:"progressHealth"`
	ProgressStatus                      *string  `pulumi:"progressStatus"`
	ProtectionStage                     *string  `pulumi:"protectionStage"`
	PsDataInMB                          *float64 `pulumi:"psDataInMB"`
	ResyncDurationInSeconds             *float64 `pulumi:"resyncDurationInSeconds"`
	ResyncLast15MinutesTransferredBytes *float64 `pulumi:"resyncLast15MinutesTransferredBytes"`
	ResyncLastDataTransferTimeUTC       *string  `pulumi:"resyncLastDataTransferTimeUTC"`
	ResyncProcessedBytes                *float64 `pulumi:"resyncProcessedBytes"`
	ResyncProgressPercentage            *int     `pulumi:"resyncProgressPercentage"`
	ResyncRequired                      *string  `pulumi:"resyncRequired"`
	ResyncStartTime                     *string  `pulumi:"resyncStartTime"`
	ResyncTotalTransferredBytes         *float64 `pulumi:"resyncTotalTransferredBytes"`
	RpoInSeconds                        *float64 `pulumi:"rpoInSeconds"`
	SourceDataInMB                      *float64 `pulumi:"sourceDataInMB"`
	TargetDataInMB                      *float64 `pulumi:"targetDataInMB"`
}

type InMageRcmAgentUpgradeBlockingErrorDetailsResponse struct {
	ErrorCode              string            `pulumi:"errorCode"`
	ErrorMessage           string            `pulumi:"errorMessage"`
	ErrorMessageParameters map[string]string `pulumi:"errorMessageParameters"`
	ErrorTags              map[string]string `pulumi:"errorTags"`
	PossibleCauses         string            `pulumi:"possibleCauses"`
	RecommendedAction      string            `pulumi:"recommendedAction"`
}

type InMageRcmDiscoveredProtectedVmDetailsResponse struct {
	CreatedTimestamp       string   `pulumi:"createdTimestamp"`
	Datastores             []string `pulumi:"datastores"`
	IpAddresses            []string `pulumi:"ipAddresses"`
	IsDeleted              bool     `pulumi:"isDeleted"`
	LastDiscoveryTimeInUtc string   `pulumi:"lastDiscoveryTimeInUtc"`
	OsName                 string   `pulumi:"osName"`
	PowerStatus            string   `pulumi:"powerStatus"`
	UpdatedTimestamp       string   `pulumi:"updatedTimestamp"`
	VCenterFqdn            string   `pulumi:"vCenterFqdn"`
	VCenterId              string   `pulumi:"vCenterId"`
	VmFqdn                 string   `pulumi:"vmFqdn"`
	VmwareToolsStatus      string   `pulumi:"vmwareToolsStatus"`
}

type InMageRcmDiskInput struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskId              string  `pulumi:"diskId"`
	DiskType            string  `pulumi:"diskType"`
	LogStorageAccountId string  `pulumi:"logStorageAccountId"`
}

type InMageRcmDisksDefaultInput struct {
	DiskEncryptionSetId *string `pulumi:"diskEncryptionSetId"`
	DiskType            string  `pulumi:"diskType"`
	LogStorageAccountId string  `pulumi:"logStorageAccountId"`
}

type InMageRcmEnableProtectionInput struct {
	DisksDefault                          *InMageRcmDisksDefaultInput `pulumi:"disksDefault"`
	DisksToInclude                        []InMageRcmDiskInput        `pulumi:"disksToInclude"`
	FabricDiscoveryMachineId              string                      `pulumi:"fabricDiscoveryMachineId"`
	InstanceType                          string                      `pulumi:"instanceType"`
	LicenseType                           *string                     `pulumi:"licenseType"`
	MultiVmGroupName                      *string                     `pulumi:"multiVmGroupName"`
	ProcessServerId                       string                      `pulumi:"processServerId"`
	RunAsAccountId                        *string                     `pulumi:"runAsAccountId"`
	TargetAvailabilitySetId               *string                     `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                *string                     `pulumi:"targetAvailabilityZone"`
	TargetBootDiagnosticsStorageAccountId *string                     `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetNetworkId                       *string                     `pulumi:"targetNetworkId"`
	TargetProximityPlacementGroupId       *string                     `pulumi:"targetProximityPlacementGroupId"`
	TargetResourceGroupId                 string                      `pulumi:"targetResourceGroupId"`
	TargetSubnetName                      *string                     `pulumi:"targetSubnetName"`
	TargetVmName                          *string                     `pulumi:"targetVmName"`
	TargetVmSize                          *string                     `pulumi:"targetVmSize"`
	TestNetworkId                         *string                     `pulumi:"testNetworkId"`
	TestSubnetName                        *string                     `pulumi:"testSubnetName"`
}

type InMageRcmFabricCreationInput struct {
	InstanceType        string                `pulumi:"instanceType"`
	PhysicalSiteId      string                `pulumi:"physicalSiteId"`
	SourceAgentIdentity IdentityProviderInput `pulumi:"sourceAgentIdentity"`
	VmwareSiteId        string                `pulumi:"vmwareSiteId"`
}

type InMageRcmFabricSpecificDetailsResponse struct {
	AgentDetails               []AgentDetailsResponse            `pulumi:"agentDetails"`
	ControlPlaneUri            string                            `pulumi:"controlPlaneUri"`
	DataPlaneUri               string                            `pulumi:"dataPlaneUri"`
	Dras                       []DraDetailsResponse              `pulumi:"dras"`
	InstanceType               string                            `pulumi:"instanceType"`
	MarsAgents                 []MarsAgentDetailsResponse        `pulumi:"marsAgents"`
	PhysicalSiteId             string                            `pulumi:"physicalSiteId"`
	ProcessServers             []ProcessServerDetailsResponse    `pulumi:"processServers"`
	PushInstallers             []PushInstallerDetailsResponse    `pulumi:"pushInstallers"`
	RcmProxies                 []RcmProxyDetailsResponse         `pulumi:"rcmProxies"`
	ReplicationAgents          []ReplicationAgentDetailsResponse `pulumi:"replicationAgents"`
	ReprotectAgents            []ReprotectAgentDetailsResponse   `pulumi:"reprotectAgents"`
	ServiceContainerId         string                            `pulumi:"serviceContainerId"`
	ServiceEndpoint            string                            `pulumi:"serviceEndpoint"`
	ServiceResourceId          string                            `pulumi:"serviceResourceId"`
	SourceAgentIdentityDetails *IdentityProviderDetailsResponse  `pulumi:"sourceAgentIdentityDetails"`
	VmwareSiteId               string                            `pulumi:"vmwareSiteId"`
}

type InMageRcmFailbackDiscoveredProtectedVmDetailsResponse struct {
	CreatedTimestamp       string   `pulumi:"createdTimestamp"`
	Datastores             []string `pulumi:"datastores"`
	IpAddresses            []string `pulumi:"ipAddresses"`
	IsDeleted              bool     `pulumi:"isDeleted"`
	LastDiscoveryTimeInUtc string   `pulumi:"lastDiscoveryTimeInUtc"`
	OsName                 string   `pulumi:"osName"`
	PowerStatus            string   `pulumi:"powerStatus"`
	UpdatedTimestamp       string   `pulumi:"updatedTimestamp"`
	VCenterFqdn            string   `pulumi:"vCenterFqdn"`
	VCenterId              string   `pulumi:"vCenterId"`
	VmFqdn                 string   `pulumi:"vmFqdn"`
	VmwareToolsStatus      string   `pulumi:"vmwareToolsStatus"`
}

type InMageRcmFailbackMobilityAgentDetailsResponse struct {
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

type InMageRcmFailbackNicDetailsResponse struct {
	AdapterType     string `pulumi:"adapterType"`
	MacAddress      string `pulumi:"macAddress"`
	NetworkName     string `pulumi:"networkName"`
	SourceIpAddress string `pulumi:"sourceIpAddress"`
}

type InMageRcmFailbackPolicyCreationInput struct {
	AppConsistentFrequencyInMinutes   *int   `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int   `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string `pulumi:"instanceType"`
}

type InMageRcmFailbackPolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   *int   `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int   `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string `pulumi:"instanceType"`
}

type InMageRcmFailbackProtectedDiskDetailsResponse struct {
	CapacityInBytes               float64                               `pulumi:"capacityInBytes"`
	DataPendingAtSourceAgentInMB  float64                               `pulumi:"dataPendingAtSourceAgentInMB"`
	DataPendingInLogDataStoreInMB float64                               `pulumi:"dataPendingInLogDataStoreInMB"`
	DiskId                        string                                `pulumi:"diskId"`
	DiskName                      string                                `pulumi:"diskName"`
	DiskUuid                      string                                `pulumi:"diskUuid"`
	IrDetails                     *InMageRcmFailbackSyncDetailsResponse `pulumi:"irDetails"`
	IsInitialReplicationComplete  string                                `pulumi:"isInitialReplicationComplete"`
	IsOSDisk                      string                                `pulumi:"isOSDisk"`
	LastSyncTime                  string                                `pulumi:"lastSyncTime"`
	ResyncDetails                 *InMageRcmFailbackSyncDetailsResponse `pulumi:"resyncDetails"`
}

type InMageRcmFailbackReplicationDetailsResponse struct {
	AzureVirtualMachineId                      string                                                 `pulumi:"azureVirtualMachineId"`
	DiscoveredVmDetails                        *InMageRcmFailbackDiscoveredProtectedVmDetailsResponse `pulumi:"discoveredVmDetails"`
	InitialReplicationProcessedBytes           float64                                                `pulumi:"initialReplicationProcessedBytes"`
	InitialReplicationProgressHealth           string                                                 `pulumi:"initialReplicationProgressHealth"`
	InitialReplicationProgressPercentage       int                                                    `pulumi:"initialReplicationProgressPercentage"`
	InitialReplicationTransferredBytes         float64                                                `pulumi:"initialReplicationTransferredBytes"`
	InstanceType                               string                                                 `pulumi:"instanceType"`
	InternalIdentifier                         string                                                 `pulumi:"internalIdentifier"`
	IsAgentRegistrationSuccessfulAfterFailover bool                                                   `pulumi:"isAgentRegistrationSuccessfulAfterFailover"`
	LastPlannedFailoverStartTime               string                                                 `pulumi:"lastPlannedFailoverStartTime"`
	LastPlannedFailoverStatus                  string                                                 `pulumi:"lastPlannedFailoverStatus"`
	LastUsedPolicyFriendlyName                 string                                                 `pulumi:"lastUsedPolicyFriendlyName"`
	LastUsedPolicyId                           string                                                 `pulumi:"lastUsedPolicyId"`
	LogStorageAccountId                        string                                                 `pulumi:"logStorageAccountId"`
	MobilityAgentDetails                       *InMageRcmFailbackMobilityAgentDetailsResponse         `pulumi:"mobilityAgentDetails"`
	MultiVmGroupName                           string                                                 `pulumi:"multiVmGroupName"`
	OsType                                     string                                                 `pulumi:"osType"`
	ProtectedDisks                             []InMageRcmFailbackProtectedDiskDetailsResponse        `pulumi:"protectedDisks"`
	ReprotectAgentId                           string                                                 `pulumi:"reprotectAgentId"`
	ReprotectAgentName                         string                                                 `pulumi:"reprotectAgentName"`
	ResyncProcessedBytes                       float64                                                `pulumi:"resyncProcessedBytes"`
	ResyncProgressHealth                       string                                                 `pulumi:"resyncProgressHealth"`
	ResyncProgressPercentage                   int                                                    `pulumi:"resyncProgressPercentage"`
	ResyncRequired                             string                                                 `pulumi:"resyncRequired"`
	ResyncState                                string                                                 `pulumi:"resyncState"`
	ResyncTransferredBytes                     float64                                                `pulumi:"resyncTransferredBytes"`
	TargetDataStoreName                        string                                                 `pulumi:"targetDataStoreName"`
	TargetVmName                               string                                                 `pulumi:"targetVmName"`
	TargetvCenterId                            string                                                 `pulumi:"targetvCenterId"`
	VmNics                                     []InMageRcmFailbackNicDetailsResponse                  `pulumi:"vmNics"`
}

type InMageRcmFailbackSyncDetailsResponse struct {
	Last15MinutesTransferredBytes float64 `pulumi:"last15MinutesTransferredBytes"`
	LastDataTransferTimeUtc       string  `pulumi:"lastDataTransferTimeUtc"`
	LastRefreshTime               string  `pulumi:"lastRefreshTime"`
	ProcessedBytes                float64 `pulumi:"processedBytes"`
	ProgressHealth                string  `pulumi:"progressHealth"`
	ProgressPercentage            int     `pulumi:"progressPercentage"`
	StartTime                     string  `pulumi:"startTime"`
	TransferredBytes              float64 `pulumi:"transferredBytes"`
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
	LatestAgentReleaseDate               string   `pulumi:"latestAgentReleaseDate"`
	LatestUpgradableVersionWithoutReboot string   `pulumi:"latestUpgradableVersionWithoutReboot"`
	LatestVersion                        string   `pulumi:"latestVersion"`
	ReasonsBlockingUpgrade               []string `pulumi:"reasonsBlockingUpgrade"`
	Version                              string   `pulumi:"version"`
}

type InMageRcmNicDetailsResponse struct {
	IsPrimaryNic          *string `pulumi:"isPrimaryNic"`
	IsSelectedForFailover *string `pulumi:"isSelectedForFailover"`
	NicId                 string  `pulumi:"nicId"`
	SourceIPAddress       string  `pulumi:"sourceIPAddress"`
	SourceIPAddressType   string  `pulumi:"sourceIPAddressType"`
	SourceNetworkId       string  `pulumi:"sourceNetworkId"`
	SourceSubnetName      string  `pulumi:"sourceSubnetName"`
	TargetIPAddress       *string `pulumi:"targetIPAddress"`
	TargetIPAddressType   *string `pulumi:"targetIPAddressType"`
	TargetSubnetName      *string `pulumi:"targetSubnetName"`
	TestIPAddress         *string `pulumi:"testIPAddress"`
	TestIPAddressType     *string `pulumi:"testIPAddressType"`
	TestSubnetName        *string `pulumi:"testSubnetName"`
}

type InMageRcmPolicyCreationInput struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	EnableMultiVmSync                 *string `pulumi:"enableMultiVmSync"`
	InstanceType                      string  `pulumi:"instanceType"`
	RecoveryPointHistoryInMinutes     *int    `pulumi:"recoveryPointHistoryInMinutes"`
}

type InMageRcmPolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	EnableMultiVmSync                 *string `pulumi:"enableMultiVmSync"`
	InstanceType                      string  `pulumi:"instanceType"`
	RecoveryPointHistoryInMinutes     *int    `pulumi:"recoveryPointHistoryInMinutes"`
}

type InMageRcmProtectedDiskDetailsResponse struct {
	CapacityInBytes               float64                       `pulumi:"capacityInBytes"`
	DataPendingAtSourceAgentInMB  float64                       `pulumi:"dataPendingAtSourceAgentInMB"`
	DataPendingInLogDataStoreInMB float64                       `pulumi:"dataPendingInLogDataStoreInMB"`
	DiskEncryptionSetId           string                        `pulumi:"diskEncryptionSetId"`
	DiskId                        string                        `pulumi:"diskId"`
	DiskName                      string                        `pulumi:"diskName"`
	DiskType                      *string                       `pulumi:"diskType"`
	IrDetails                     *InMageRcmSyncDetailsResponse `pulumi:"irDetails"`
	IsInitialReplicationComplete  string                        `pulumi:"isInitialReplicationComplete"`
	IsOSDisk                      string                        `pulumi:"isOSDisk"`
	LogStorageAccountId           string                        `pulumi:"logStorageAccountId"`
	ResyncDetails                 *InMageRcmSyncDetailsResponse `pulumi:"resyncDetails"`
	SeedBlobUri                   string                        `pulumi:"seedBlobUri"`
	SeedManagedDiskId             string                        `pulumi:"seedManagedDiskId"`
	TargetManagedDiskId           string                        `pulumi:"targetManagedDiskId"`
}

type InMageRcmProtectionContainerMappingDetailsResponse struct {
	EnableAgentAutoUpgrade string `pulumi:"enableAgentAutoUpgrade"`
	InstanceType           string `pulumi:"instanceType"`
}

type InMageRcmReplicationDetailsResponse struct {
	AgentUpgradeAttemptToVersion               string                                              `pulumi:"agentUpgradeAttemptToVersion"`
	AgentUpgradeBlockingErrorDetails           []InMageRcmAgentUpgradeBlockingErrorDetailsResponse `pulumi:"agentUpgradeBlockingErrorDetails"`
	AgentUpgradeJobId                          string                                              `pulumi:"agentUpgradeJobId"`
	AgentUpgradeState                          string                                              `pulumi:"agentUpgradeState"`
	AllocatedMemoryInMB                        float64                                             `pulumi:"allocatedMemoryInMB"`
	DiscoveredVmDetails                        *InMageRcmDiscoveredProtectedVmDetailsResponse      `pulumi:"discoveredVmDetails"`
	DiscoveryType                              string                                              `pulumi:"discoveryType"`
	FabricDiscoveryMachineId                   string                                              `pulumi:"fabricDiscoveryMachineId"`
	FailoverRecoveryPointId                    string                                              `pulumi:"failoverRecoveryPointId"`
	FirmwareType                               string                                              `pulumi:"firmwareType"`
	InitialReplicationProcessedBytes           float64                                             `pulumi:"initialReplicationProcessedBytes"`
	InitialReplicationProgressHealth           string                                              `pulumi:"initialReplicationProgressHealth"`
	InitialReplicationProgressPercentage       int                                                 `pulumi:"initialReplicationProgressPercentage"`
	InitialReplicationTransferredBytes         float64                                             `pulumi:"initialReplicationTransferredBytes"`
	InstanceType                               string                                              `pulumi:"instanceType"`
	InternalIdentifier                         string                                              `pulumi:"internalIdentifier"`
	IsAgentRegistrationSuccessfulAfterFailover bool                                                `pulumi:"isAgentRegistrationSuccessfulAfterFailover"`
	IsLastUpgradeSuccessful                    string                                              `pulumi:"isLastUpgradeSuccessful"`
	LastAgentUpgradeErrorDetails               []InMageRcmLastAgentUpgradeErrorDetailsResponse     `pulumi:"lastAgentUpgradeErrorDetails"`
	LastAgentUpgradeType                       string                                              `pulumi:"lastAgentUpgradeType"`
	LastRecoveryPointId                        string                                              `pulumi:"lastRecoveryPointId"`
	LastRecoveryPointReceived                  string                                              `pulumi:"lastRecoveryPointReceived"`
	LastRpoCalculatedTime                      string                                              `pulumi:"lastRpoCalculatedTime"`
	LastRpoInSeconds                           float64                                             `pulumi:"lastRpoInSeconds"`
	LicenseType                                *string                                             `pulumi:"licenseType"`
	MobilityAgentDetails                       *InMageRcmMobilityAgentDetailsResponse              `pulumi:"mobilityAgentDetails"`
	MultiVmGroupName                           string                                              `pulumi:"multiVmGroupName"`
	OsType                                     string                                              `pulumi:"osType"`
	PrimaryNicIpAddress                        string                                              `pulumi:"primaryNicIpAddress"`
	ProcessServerId                            string                                              `pulumi:"processServerId"`
	ProcessServerName                          string                                              `pulumi:"processServerName"`
	ProcessorCoreCount                         int                                                 `pulumi:"processorCoreCount"`
	ProtectedDisks                             []InMageRcmProtectedDiskDetailsResponse             `pulumi:"protectedDisks"`
	ResyncProcessedBytes                       float64                                             `pulumi:"resyncProcessedBytes"`
	ResyncProgressHealth                       string                                              `pulumi:"resyncProgressHealth"`
	ResyncProgressPercentage                   int                                                 `pulumi:"resyncProgressPercentage"`
	ResyncRequired                             string                                              `pulumi:"resyncRequired"`
	ResyncState                                string                                              `pulumi:"resyncState"`
	ResyncTransferredBytes                     float64                                             `pulumi:"resyncTransferredBytes"`
	RunAsAccountId                             string                                              `pulumi:"runAsAccountId"`
	StorageAccountId                           string                                              `pulumi:"storageAccountId"`
	TargetAvailabilitySetId                    *string                                             `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                     *string                                             `pulumi:"targetAvailabilityZone"`
	TargetBootDiagnosticsStorageAccountId      *string                                             `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetGeneration                           string                                              `pulumi:"targetGeneration"`
	TargetLocation                             *string                                             `pulumi:"targetLocation"`
	TargetNetworkId                            *string                                             `pulumi:"targetNetworkId"`
	TargetProximityPlacementGroupId            *string                                             `pulumi:"targetProximityPlacementGroupId"`
	TargetResourceGroupId                      *string                                             `pulumi:"targetResourceGroupId"`
	TargetVmName                               *string                                             `pulumi:"targetVmName"`
	TargetVmSize                               *string                                             `pulumi:"targetVmSize"`
	TestNetworkId                              *string                                             `pulumi:"testNetworkId"`
	VmNics                                     []InMageRcmNicDetailsResponse                       `pulumi:"vmNics"`
}

type InMageRcmSyncDetailsResponse struct {
	Last15MinutesTransferredBytes float64 `pulumi:"last15MinutesTransferredBytes"`
	LastDataTransferTimeUtc       string  `pulumi:"lastDataTransferTimeUtc"`
	LastRefreshTime               string  `pulumi:"lastRefreshTime"`
	ProcessedBytes                float64 `pulumi:"processedBytes"`
	ProgressHealth                string  `pulumi:"progressHealth"`
	ProgressPercentage            int     `pulumi:"progressPercentage"`
	StartTime                     string  `pulumi:"startTime"`
	TransferredBytes              float64 `pulumi:"transferredBytes"`
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
	IsAdditionalStatsAvailable   *bool                                `pulumi:"isAdditionalStatsAvailable"`
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
	TotalDataTransferred         *float64                             `pulumi:"totalDataTransferred"`
	TotalProgressHealth          *string                              `pulumi:"totalProgressHealth"`
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
	CustomerResolvability        *string `pulumi:"customerResolvability"`
	EntityId                     *string `pulumi:"entityId"`
	ErrorCategory                *string `pulumi:"errorCategory"`
	ErrorCode                    *string `pulumi:"errorCode"`
	ErrorId                      *string `pulumi:"errorId"`
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

func (o InnerHealthErrorResponseOutput) CustomerResolvability() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.CustomerResolvability }).(pulumi.StringPtrOutput)
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

func (o InnerHealthErrorResponseOutput) ErrorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InnerHealthErrorResponse) *string { return v.ErrorId }).(pulumi.StringPtrOutput)
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

type KeyEncryptionKeyInfo struct {
	KeyIdentifier         *string `pulumi:"keyIdentifier"`
	KeyVaultResourceArmId *string `pulumi:"keyVaultResourceArmId"`
}

type MarsAgentDetailsResponse struct {
	BiosId           string                `pulumi:"biosId"`
	FabricObjectId   string                `pulumi:"fabricObjectId"`
	Fqdn             string                `pulumi:"fqdn"`
	Health           string                `pulumi:"health"`
	HealthErrors     []HealthErrorResponse `pulumi:"healthErrors"`
	Id               string                `pulumi:"id"`
	LastHeartbeatUtc string                `pulumi:"lastHeartbeatUtc"`
	Name             string                `pulumi:"name"`
	Version          string                `pulumi:"version"`
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
	AllowedOperations           []string                            `pulumi:"allowedOperations"`
	CriticalJobHistory          []CriticalJobHistoryDetailsResponse `pulumi:"criticalJobHistory"`
	CurrentJob                  CurrentJobDetailsResponse           `pulumi:"currentJob"`
	EventCorrelationId          string                              `pulumi:"eventCorrelationId"`
	Health                      string                              `pulumi:"health"`
	HealthErrors                []HealthErrorResponse               `pulumi:"healthErrors"`
	LastMigrationStatus         string                              `pulumi:"lastMigrationStatus"`
	LastMigrationTime           string                              `pulumi:"lastMigrationTime"`
	LastTestMigrationStatus     string                              `pulumi:"lastTestMigrationStatus"`
	LastTestMigrationTime       string                              `pulumi:"lastTestMigrationTime"`
	MachineName                 string                              `pulumi:"machineName"`
	MigrationState              string                              `pulumi:"migrationState"`
	MigrationStateDescription   string                              `pulumi:"migrationStateDescription"`
	PolicyFriendlyName          string                              `pulumi:"policyFriendlyName"`
	PolicyId                    string                              `pulumi:"policyId"`
	ProviderSpecificDetails     *VMwareCbtMigrationDetailsResponse  `pulumi:"providerSpecificDetails"`
	RecoveryServicesProviderId  string                              `pulumi:"recoveryServicesProviderId"`
	ReplicationStatus           string                              `pulumi:"replicationStatus"`
	TestMigrateState            string                              `pulumi:"testMigrateState"`
	TestMigrateStateDescription string                              `pulumi:"testMigrateStateDescription"`
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

func (o MigrationItemPropertiesResponseOutput) CriticalJobHistory() CriticalJobHistoryDetailsResponseArrayOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) []CriticalJobHistoryDetailsResponse {
		return v.CriticalJobHistory
	}).(CriticalJobHistoryDetailsResponseArrayOutput)
}

func (o MigrationItemPropertiesResponseOutput) CurrentJob() CurrentJobDetailsResponseOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) CurrentJobDetailsResponse { return v.CurrentJob }).(CurrentJobDetailsResponseOutput)
}

func (o MigrationItemPropertiesResponseOutput) EventCorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.EventCorrelationId }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.Health }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) HealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) []HealthErrorResponse { return v.HealthErrors }).(HealthErrorResponseArrayOutput)
}

func (o MigrationItemPropertiesResponseOutput) LastMigrationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.LastMigrationStatus }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) LastMigrationTime() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.LastMigrationTime }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) LastTestMigrationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.LastTestMigrationStatus }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) LastTestMigrationTime() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.LastTestMigrationTime }).(pulumi.StringOutput)
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

func (o MigrationItemPropertiesResponseOutput) ReplicationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.ReplicationStatus }).(pulumi.StringOutput)
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

type ProcessServerDetailsResponse struct {
	AvailableMemoryInBytes             float64               `pulumi:"availableMemoryInBytes"`
	AvailableSpaceInBytes              float64               `pulumi:"availableSpaceInBytes"`
	BiosId                             string                `pulumi:"biosId"`
	DiskUsageStatus                    string                `pulumi:"diskUsageStatus"`
	FabricObjectId                     string                `pulumi:"fabricObjectId"`
	Fqdn                               string                `pulumi:"fqdn"`
	FreeSpacePercentage                float64               `pulumi:"freeSpacePercentage"`
	Health                             string                `pulumi:"health"`
	HealthErrors                       []HealthErrorResponse `pulumi:"healthErrors"`
	HistoricHealth                     string                `pulumi:"historicHealth"`
	Id                                 string                `pulumi:"id"`
	IpAddresses                        []string              `pulumi:"ipAddresses"`
	LastHeartbeatUtc                   string                `pulumi:"lastHeartbeatUtc"`
	MemoryUsagePercentage              float64               `pulumi:"memoryUsagePercentage"`
	MemoryUsageStatus                  string                `pulumi:"memoryUsageStatus"`
	Name                               string                `pulumi:"name"`
	ProcessorUsagePercentage           float64               `pulumi:"processorUsagePercentage"`
	ProcessorUsageStatus               string                `pulumi:"processorUsageStatus"`
	ProtectedItemCount                 int                   `pulumi:"protectedItemCount"`
	SystemLoad                         float64               `pulumi:"systemLoad"`
	SystemLoadStatus                   string                `pulumi:"systemLoadStatus"`
	ThroughputInBytes                  float64               `pulumi:"throughputInBytes"`
	ThroughputStatus                   string                `pulumi:"throughputStatus"`
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
	BiosId           string                `pulumi:"biosId"`
	FabricObjectId   string                `pulumi:"fabricObjectId"`
	Fqdn             string                `pulumi:"fqdn"`
	Health           string                `pulumi:"health"`
	HealthErrors     []HealthErrorResponse `pulumi:"healthErrors"`
	Id               string                `pulumi:"id"`
	LastHeartbeatUtc string                `pulumi:"lastHeartbeatUtc"`
	Name             string                `pulumi:"name"`
	Version          string                `pulumi:"version"`
}

type RcmProxyDetailsResponse struct {
	BiosId                   string                `pulumi:"biosId"`
	ClientAuthenticationType string                `pulumi:"clientAuthenticationType"`
	FabricObjectId           string                `pulumi:"fabricObjectId"`
	Fqdn                     string                `pulumi:"fqdn"`
	Health                   string                `pulumi:"health"`
	HealthErrors             []HealthErrorResponse `pulumi:"healthErrors"`
	Id                       string                `pulumi:"id"`
	LastHeartbeatUtc         string                `pulumi:"lastHeartbeatUtc"`
	Name                     string                `pulumi:"name"`
	Version                  string                `pulumi:"version"`
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
	InstanceType             string            `pulumi:"instanceType"`
	PrimaryExtendedLocation  *ExtendedLocation `pulumi:"primaryExtendedLocation"`
	PrimaryZone              *string           `pulumi:"primaryZone"`
	RecoveryExtendedLocation *ExtendedLocation `pulumi:"recoveryExtendedLocation"`
	RecoveryZone             *string           `pulumi:"recoveryZone"`
}





type RecoveryPlanA2AInputInput interface {
	pulumi.Input

	ToRecoveryPlanA2AInputOutput() RecoveryPlanA2AInputOutput
	ToRecoveryPlanA2AInputOutputWithContext(context.Context) RecoveryPlanA2AInputOutput
}

type RecoveryPlanA2AInputArgs struct {
	InstanceType             pulumi.StringInput       `pulumi:"instanceType"`
	PrimaryExtendedLocation  ExtendedLocationPtrInput `pulumi:"primaryExtendedLocation"`
	PrimaryZone              pulumi.StringPtrInput    `pulumi:"primaryZone"`
	RecoveryExtendedLocation ExtendedLocationPtrInput `pulumi:"recoveryExtendedLocation"`
	RecoveryZone             pulumi.StringPtrInput    `pulumi:"recoveryZone"`
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

func (o RecoveryPlanA2AInputOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanA2AInput) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o RecoveryPlanA2AInputOutput) PrimaryExtendedLocation() ExtendedLocationPtrOutput {
	return o.ApplyT(func(v RecoveryPlanA2AInput) *ExtendedLocation { return v.PrimaryExtendedLocation }).(ExtendedLocationPtrOutput)
}

func (o RecoveryPlanA2AInputOutput) PrimaryZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanA2AInput) *string { return v.PrimaryZone }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanA2AInputOutput) RecoveryExtendedLocation() ExtendedLocationPtrOutput {
	return o.ApplyT(func(v RecoveryPlanA2AInput) *ExtendedLocation { return v.RecoveryExtendedLocation }).(ExtendedLocationPtrOutput)
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
	ActionName         string      `pulumi:"actionName"`
	CustomDetails      interface{} `pulumi:"customDetails"`
	FailoverDirections []string    `pulumi:"failoverDirections"`
	FailoverTypes      []string    `pulumi:"failoverTypes"`
}





type RecoveryPlanActionInput interface {
	pulumi.Input

	ToRecoveryPlanActionOutput() RecoveryPlanActionOutput
	ToRecoveryPlanActionOutputWithContext(context.Context) RecoveryPlanActionOutput
}

type RecoveryPlanActionArgs struct {
	ActionName         pulumi.StringInput      `pulumi:"actionName"`
	CustomDetails      pulumi.Input            `pulumi:"customDetails"`
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

func (o RecoveryPlanActionOutput) CustomDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v RecoveryPlanAction) interface{} { return v.CustomDetails }).(pulumi.AnyOutput)
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

type RecoveryPlanAutomationRunbookActionDetails struct {
	FabricLocation string  `pulumi:"fabricLocation"`
	InstanceType   string  `pulumi:"instanceType"`
	RunbookId      *string `pulumi:"runbookId"`
	Timeout        *string `pulumi:"timeout"`
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

type RecoveryPlanManualActionDetails struct {
	Description  *string `pulumi:"description"`
	InstanceType string  `pulumi:"instanceType"`
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

type RecoveryPlanScriptActionDetails struct {
	FabricLocation string  `pulumi:"fabricLocation"`
	InstanceType   string  `pulumi:"instanceType"`
	Path           string  `pulumi:"path"`
	Timeout        *string `pulumi:"timeout"`
}

type RecoveryPlanScriptActionDetailsResponse struct {
	FabricLocation string  `pulumi:"fabricLocation"`
	InstanceType   string  `pulumi:"instanceType"`
	Path           string  `pulumi:"path"`
	Timeout        *string `pulumi:"timeout"`
}

type RecoveryServicesProviderPropertiesResponse struct {
	AllowedScenarios                       []string                         `pulumi:"allowedScenarios"`
	AuthenticationIdentityDetails          *IdentityProviderDetailsResponse `pulumi:"authenticationIdentityDetails"`
	BiosId                                 *string                          `pulumi:"biosId"`
	ConnectionStatus                       *string                          `pulumi:"connectionStatus"`
	DataPlaneAuthenticationIdentityDetails *IdentityProviderDetailsResponse `pulumi:"dataPlaneAuthenticationIdentityDetails"`
	DraIdentifier                          *string                          `pulumi:"draIdentifier"`
	FabricFriendlyName                     *string                          `pulumi:"fabricFriendlyName"`
	FabricType                             *string                          `pulumi:"fabricType"`
	FriendlyName                           *string                          `pulumi:"friendlyName"`
	HealthErrorDetails                     []HealthErrorResponse            `pulumi:"healthErrorDetails"`
	LastHeartBeat                          *string                          `pulumi:"lastHeartBeat"`
	MachineId                              *string                          `pulumi:"machineId"`
	MachineName                            *string                          `pulumi:"machineName"`
	ProtectedItemCount                     *int                             `pulumi:"protectedItemCount"`
	ProviderVersion                        *string                          `pulumi:"providerVersion"`
	ProviderVersionDetails                 *VersionDetailsResponse          `pulumi:"providerVersionDetails"`
	ProviderVersionExpiryDate              *string                          `pulumi:"providerVersionExpiryDate"`
	ProviderVersionState                   *string                          `pulumi:"providerVersionState"`
	ResourceAccessIdentityDetails          *IdentityProviderDetailsResponse `pulumi:"resourceAccessIdentityDetails"`
	ServerVersion                          *string                          `pulumi:"serverVersion"`
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

func (o RecoveryServicesProviderPropertiesResponseOutput) BiosId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.BiosId }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ConnectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.ConnectionStatus }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) DataPlaneAuthenticationIdentityDetails() IdentityProviderDetailsResponsePtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *IdentityProviderDetailsResponse {
		return v.DataPlaneAuthenticationIdentityDetails
	}).(IdentityProviderDetailsResponsePtrOutput)
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

func (o RecoveryServicesProviderPropertiesResponseOutput) MachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.MachineId }).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponseOutput) MachineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryServicesProviderPropertiesResponse) *string { return v.MachineName }).(pulumi.StringPtrOutput)
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
	BiosId           string                `pulumi:"biosId"`
	FabricObjectId   string                `pulumi:"fabricObjectId"`
	Fqdn             string                `pulumi:"fqdn"`
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
	EventCorrelationId                      *string                         `pulumi:"eventCorrelationId"`
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
	SwitchProviderState                     *string                         `pulumi:"switchProviderState"`
	SwitchProviderStateDescription          *string                         `pulumi:"switchProviderStateDescription"`
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

func (o ReplicationProtectedItemPropertiesResponseOutput) EventCorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.EventCorrelationId }).(pulumi.StringPtrOutput)
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

func (o ReplicationProtectedItemPropertiesResponseOutput) SwitchProviderState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.SwitchProviderState }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) SwitchProviderStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.SwitchProviderStateDescription }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) TestFailoverState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.TestFailoverState }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) TestFailoverStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.TestFailoverStateDescription }).(pulumi.StringPtrOutput)
}

type ReprotectAgentDetailsResponse struct {
	AccessibleDatastores []string              `pulumi:"accessibleDatastores"`
	BiosId               string                `pulumi:"biosId"`
	FabricObjectId       string                `pulumi:"fabricObjectId"`
	Fqdn                 string                `pulumi:"fqdn"`
	Health               string                `pulumi:"health"`
	HealthErrors         []HealthErrorResponse `pulumi:"healthErrors"`
	Id                   string                `pulumi:"id"`
	LastDiscoveryInUtc   string                `pulumi:"lastDiscoveryInUtc"`
	LastHeartbeatUtc     string                `pulumi:"lastHeartbeatUtc"`
	Name                 string                `pulumi:"name"`
	ProtectedItemCount   int                   `pulumi:"protectedItemCount"`
	VcenterId            string                `pulumi:"vcenterId"`
	Version              string                `pulumi:"version"`
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
	EnableAcceleratedNetworkingOnRecovery *bool                     `pulumi:"enableAcceleratedNetworkingOnRecovery"`
	EnableAcceleratedNetworkingOnTfo      *bool                     `pulumi:"enableAcceleratedNetworkingOnTfo"`
	IpConfigs                             []IPConfigDetailsResponse `pulumi:"ipConfigs"`
	NicId                                 *string                   `pulumi:"nicId"`
	RecoveryNetworkSecurityGroupId        *string                   `pulumi:"recoveryNetworkSecurityGroupId"`
	RecoveryNicName                       *string                   `pulumi:"recoveryNicName"`
	RecoveryNicResourceGroupName          *string                   `pulumi:"recoveryNicResourceGroupName"`
	RecoveryVMNetworkId                   *string                   `pulumi:"recoveryVMNetworkId"`
	ReplicaNicId                          *string                   `pulumi:"replicaNicId"`
	ReuseExistingNic                      *bool                     `pulumi:"reuseExistingNic"`
	SelectionType                         *string                   `pulumi:"selectionType"`
	SourceNicArmId                        *string                   `pulumi:"sourceNicArmId"`
	TargetNicName                         *string                   `pulumi:"targetNicName"`
	TfoNetworkSecurityGroupId             *string                   `pulumi:"tfoNetworkSecurityGroupId"`
	TfoRecoveryNicName                    *string                   `pulumi:"tfoRecoveryNicName"`
	TfoRecoveryNicResourceGroupName       *string                   `pulumi:"tfoRecoveryNicResourceGroupName"`
	TfoReuseExistingNic                   *bool                     `pulumi:"tfoReuseExistingNic"`
	TfoVMNetworkId                        *string                   `pulumi:"tfoVMNetworkId"`
	VMNetworkName                         *string                   `pulumi:"vMNetworkName"`
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
	InstanceType                         string  `pulumi:"instanceType"`
	KeyVaultId                           *string `pulumi:"keyVaultId"`
	KeyVaultUri                          *string `pulumi:"keyVaultUri"`
	ServiceBusConnectionStringSecretName *string `pulumi:"serviceBusConnectionStringSecretName"`
	StorageAccountId                     string  `pulumi:"storageAccountId"`
	StorageAccountSasSecretName          *string `pulumi:"storageAccountSasSecretName"`
	TargetLocation                       string  `pulumi:"targetLocation"`
}

type VMwareCbtDiskInput struct {
	DiskEncryptionSetId            *string `pulumi:"diskEncryptionSetId"`
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
	DiskEncryptionSetId            pulumi.StringPtrInput `pulumi:"diskEncryptionSetId"`
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

func (o VMwareCbtDiskInputOutput) DiskEncryptionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtDiskInput) *string { return v.DiskEncryptionSetId }).(pulumi.StringPtrOutput)
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
	PerformAutoResync                     *string              `pulumi:"performAutoResync"`
	SeedDiskTags                          map[string]string    `pulumi:"seedDiskTags"`
	SnapshotRunAsAccountId                string               `pulumi:"snapshotRunAsAccountId"`
	SqlServerLicenseType                  *string              `pulumi:"sqlServerLicenseType"`
	TargetAvailabilitySetId               *string              `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                *string              `pulumi:"targetAvailabilityZone"`
	TargetBootDiagnosticsStorageAccountId *string              `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetDiskTags                        map[string]string    `pulumi:"targetDiskTags"`
	TargetNetworkId                       string               `pulumi:"targetNetworkId"`
	TargetNicTags                         map[string]string    `pulumi:"targetNicTags"`
	TargetProximityPlacementGroupId       *string              `pulumi:"targetProximityPlacementGroupId"`
	TargetResourceGroupId                 string               `pulumi:"targetResourceGroupId"`
	TargetSubnetName                      *string              `pulumi:"targetSubnetName"`
	TargetVmName                          *string              `pulumi:"targetVmName"`
	TargetVmSize                          *string              `pulumi:"targetVmSize"`
	TargetVmTags                          map[string]string    `pulumi:"targetVmTags"`
	TestNetworkId                         *string              `pulumi:"testNetworkId"`
	TestSubnetName                        *string              `pulumi:"testSubnetName"`
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
	PerformAutoResync                     pulumi.StringPtrInput        `pulumi:"performAutoResync"`
	SeedDiskTags                          pulumi.StringMapInput        `pulumi:"seedDiskTags"`
	SnapshotRunAsAccountId                pulumi.StringInput           `pulumi:"snapshotRunAsAccountId"`
	SqlServerLicenseType                  pulumi.StringPtrInput        `pulumi:"sqlServerLicenseType"`
	TargetAvailabilitySetId               pulumi.StringPtrInput        `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                pulumi.StringPtrInput        `pulumi:"targetAvailabilityZone"`
	TargetBootDiagnosticsStorageAccountId pulumi.StringPtrInput        `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetDiskTags                        pulumi.StringMapInput        `pulumi:"targetDiskTags"`
	TargetNetworkId                       pulumi.StringInput           `pulumi:"targetNetworkId"`
	TargetNicTags                         pulumi.StringMapInput        `pulumi:"targetNicTags"`
	TargetProximityPlacementGroupId       pulumi.StringPtrInput        `pulumi:"targetProximityPlacementGroupId"`
	TargetResourceGroupId                 pulumi.StringInput           `pulumi:"targetResourceGroupId"`
	TargetSubnetName                      pulumi.StringPtrInput        `pulumi:"targetSubnetName"`
	TargetVmName                          pulumi.StringPtrInput        `pulumi:"targetVmName"`
	TargetVmSize                          pulumi.StringPtrInput        `pulumi:"targetVmSize"`
	TargetVmTags                          pulumi.StringMapInput        `pulumi:"targetVmTags"`
	TestNetworkId                         pulumi.StringPtrInput        `pulumi:"testNetworkId"`
	TestSubnetName                        pulumi.StringPtrInput        `pulumi:"testSubnetName"`
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

func (o VMwareCbtEnableMigrationInputOutput) PerformAutoResync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.PerformAutoResync }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) SeedDiskTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) map[string]string { return v.SeedDiskTags }).(pulumi.StringMapOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) SnapshotRunAsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) string { return v.SnapshotRunAsAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) SqlServerLicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.SqlServerLicenseType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TargetAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TargetAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetBootDiagnosticsStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TargetBootDiagnosticsStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetDiskTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) map[string]string { return v.TargetDiskTags }).(pulumi.StringMapOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) string { return v.TargetNetworkId }).(pulumi.StringOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetNicTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) map[string]string { return v.TargetNicTags }).(pulumi.StringMapOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TargetProximityPlacementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TargetProximityPlacementGroupId }).(pulumi.StringPtrOutput)
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

func (o VMwareCbtEnableMigrationInputOutput) TargetVmTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) map[string]string { return v.TargetVmTags }).(pulumi.StringMapOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TestNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TestNetworkId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) TestSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) *string { return v.TestSubnetName }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputOutput) VmwareMachineId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) string { return v.VmwareMachineId }).(pulumi.StringOutput)
}

type VMwareCbtMigrationDetailsResponse struct {
	DataMoverRunAsAccountId               string                                  `pulumi:"dataMoverRunAsAccountId"`
	FirmwareType                          string                                  `pulumi:"firmwareType"`
	InitialSeedingProgressPercentage      int                                     `pulumi:"initialSeedingProgressPercentage"`
	InitialSeedingRetryCount              float64                                 `pulumi:"initialSeedingRetryCount"`
	InstanceType                          string                                  `pulumi:"instanceType"`
	LastRecoveryPointId                   string                                  `pulumi:"lastRecoveryPointId"`
	LastRecoveryPointReceived             string                                  `pulumi:"lastRecoveryPointReceived"`
	LicenseType                           *string                                 `pulumi:"licenseType"`
	MigrationProgressPercentage           int                                     `pulumi:"migrationProgressPercentage"`
	MigrationRecoveryPointId              string                                  `pulumi:"migrationRecoveryPointId"`
	OsType                                string                                  `pulumi:"osType"`
	PerformAutoResync                     *string                                 `pulumi:"performAutoResync"`
	ProtectedDisks                        []VMwareCbtProtectedDiskDetailsResponse `pulumi:"protectedDisks"`
	ResumeProgressPercentage              int                                     `pulumi:"resumeProgressPercentage"`
	ResumeRetryCount                      float64                                 `pulumi:"resumeRetryCount"`
	ResyncProgressPercentage              int                                     `pulumi:"resyncProgressPercentage"`
	ResyncRequired                        string                                  `pulumi:"resyncRequired"`
	ResyncRetryCount                      float64                                 `pulumi:"resyncRetryCount"`
	ResyncState                           string                                  `pulumi:"resyncState"`
	SeedDiskTags                          map[string]string                       `pulumi:"seedDiskTags"`
	SnapshotRunAsAccountId                string                                  `pulumi:"snapshotRunAsAccountId"`
	SqlServerLicenseType                  *string                                 `pulumi:"sqlServerLicenseType"`
	StorageAccountId                      string                                  `pulumi:"storageAccountId"`
	TargetAvailabilitySetId               *string                                 `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                *string                                 `pulumi:"targetAvailabilityZone"`
	TargetBootDiagnosticsStorageAccountId *string                                 `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetDiskTags                        map[string]string                       `pulumi:"targetDiskTags"`
	TargetGeneration                      string                                  `pulumi:"targetGeneration"`
	TargetLocation                        string                                  `pulumi:"targetLocation"`
	TargetNetworkId                       *string                                 `pulumi:"targetNetworkId"`
	TargetNicTags                         map[string]string                       `pulumi:"targetNicTags"`
	TargetProximityPlacementGroupId       *string                                 `pulumi:"targetProximityPlacementGroupId"`
	TargetResourceGroupId                 *string                                 `pulumi:"targetResourceGroupId"`
	TargetVmName                          *string                                 `pulumi:"targetVmName"`
	TargetVmSize                          *string                                 `pulumi:"targetVmSize"`
	TargetVmTags                          map[string]string                       `pulumi:"targetVmTags"`
	TestNetworkId                         *string                                 `pulumi:"testNetworkId"`
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

func (o VMwareCbtMigrationDetailsResponseOutput) FirmwareType() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.FirmwareType }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) InitialSeedingProgressPercentage() pulumi.IntOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) int { return v.InitialSeedingProgressPercentage }).(pulumi.IntOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) InitialSeedingRetryCount() pulumi.Float64Output {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) float64 { return v.InitialSeedingRetryCount }).(pulumi.Float64Output)
}

func (o VMwareCbtMigrationDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) LastRecoveryPointId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.LastRecoveryPointId }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) LastRecoveryPointReceived() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.LastRecoveryPointReceived }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) MigrationProgressPercentage() pulumi.IntOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) int { return v.MigrationProgressPercentage }).(pulumi.IntOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) MigrationRecoveryPointId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.MigrationRecoveryPointId }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) PerformAutoResync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.PerformAutoResync }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) ProtectedDisks() VMwareCbtProtectedDiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) []VMwareCbtProtectedDiskDetailsResponse {
		return v.ProtectedDisks
	}).(VMwareCbtProtectedDiskDetailsResponseArrayOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) ResumeProgressPercentage() pulumi.IntOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) int { return v.ResumeProgressPercentage }).(pulumi.IntOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) ResumeRetryCount() pulumi.Float64Output {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) float64 { return v.ResumeRetryCount }).(pulumi.Float64Output)
}

func (o VMwareCbtMigrationDetailsResponseOutput) ResyncProgressPercentage() pulumi.IntOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) int { return v.ResyncProgressPercentage }).(pulumi.IntOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) ResyncRequired() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.ResyncRequired }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) ResyncRetryCount() pulumi.Float64Output {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) float64 { return v.ResyncRetryCount }).(pulumi.Float64Output)
}

func (o VMwareCbtMigrationDetailsResponseOutput) ResyncState() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.ResyncState }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) SeedDiskTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) map[string]string { return v.SeedDiskTags }).(pulumi.StringMapOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) SnapshotRunAsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.SnapshotRunAsAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) SqlServerLicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.SqlServerLicenseType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetAvailabilityZone }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetBootDiagnosticsStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetBootDiagnosticsStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetDiskTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) map[string]string { return v.TargetDiskTags }).(pulumi.StringMapOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetGeneration() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.TargetGeneration }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetLocation() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.TargetLocation }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetNetworkId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetNicTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) map[string]string { return v.TargetNicTags }).(pulumi.StringMapOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TargetProximityPlacementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TargetProximityPlacementGroupId }).(pulumi.StringPtrOutput)
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

func (o VMwareCbtMigrationDetailsResponseOutput) TargetVmTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) map[string]string { return v.TargetVmTags }).(pulumi.StringMapOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) TestNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) *string { return v.TestNetworkId }).(pulumi.StringPtrOutput)
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

func (o VMwareCbtMigrationDetailsResponsePtrOutput) FirmwareType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FirmwareType
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) InitialSeedingProgressPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.InitialSeedingProgressPercentage
	}).(pulumi.IntPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) InitialSeedingRetryCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.InitialSeedingRetryCount
	}).(pulumi.Float64PtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InstanceType
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) LastRecoveryPointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastRecoveryPointId
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

func (o VMwareCbtMigrationDetailsResponsePtrOutput) MigrationProgressPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MigrationProgressPercentage
	}).(pulumi.IntPtrOutput)
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

func (o VMwareCbtMigrationDetailsResponsePtrOutput) PerformAutoResync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PerformAutoResync
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

func (o VMwareCbtMigrationDetailsResponsePtrOutput) ResumeProgressPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ResumeProgressPercentage
	}).(pulumi.IntPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) ResumeRetryCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.ResumeRetryCount
	}).(pulumi.Float64PtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) ResyncProgressPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *int {
		if v == nil {
			return nil
		}
		return &v.ResyncProgressPercentage
	}).(pulumi.IntPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) ResyncRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResyncRequired
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) ResyncRetryCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.ResyncRetryCount
	}).(pulumi.Float64PtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) ResyncState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResyncState
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) SeedDiskTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.SeedDiskTags
	}).(pulumi.StringMapOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) SnapshotRunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SnapshotRunAsAccountId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) SqlServerLicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SqlServerLicenseType
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StorageAccountId
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

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetAvailabilityZone
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

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetDiskTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.TargetDiskTags
	}).(pulumi.StringMapOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetGeneration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TargetGeneration
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

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetNicTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.TargetNicTags
	}).(pulumi.StringMapOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetProximityPlacementGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetProximityPlacementGroupId
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

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TargetVmTags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.TargetVmTags
	}).(pulumi.StringMapOutput)
}

func (o VMwareCbtMigrationDetailsResponsePtrOutput) TestNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TestNetworkId
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
	TargetNicName          *string `pulumi:"targetNicName"`
	TargetSubnetName       *string `pulumi:"targetSubnetName"`
	TestIPAddress          *string `pulumi:"testIPAddress"`
	TestIPAddressType      *string `pulumi:"testIPAddressType"`
	TestNetworkId          *string `pulumi:"testNetworkId"`
	TestSubnetName         *string `pulumi:"testSubnetName"`
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

func (o VMwareCbtNicDetailsResponseOutput) TargetNicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.TargetNicName }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) TargetSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.TargetSubnetName }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) TestIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.TestIPAddress }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) TestIPAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.TestIPAddressType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) TestNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.TestNetworkId }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtNicDetailsResponseOutput) TestSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtNicDetailsResponse) *string { return v.TestSubnetName }).(pulumi.StringPtrOutput)
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
	AppConsistentFrequencyInMinutes   *int   `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int   `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string `pulumi:"instanceType"`
	RecoveryPointHistoryInMinutes     *int   `pulumi:"recoveryPointHistoryInMinutes"`
}

type VMwareCbtProtectedDiskDetailsResponse struct {
	CapacityInBytes                float64 `pulumi:"capacityInBytes"`
	DiskEncryptionSetId            string  `pulumi:"diskEncryptionSetId"`
	DiskId                         string  `pulumi:"diskId"`
	DiskName                       string  `pulumi:"diskName"`
	DiskPath                       string  `pulumi:"diskPath"`
	DiskType                       *string `pulumi:"diskType"`
	IsOSDisk                       string  `pulumi:"isOSDisk"`
	LogStorageAccountId            string  `pulumi:"logStorageAccountId"`
	LogStorageAccountSasSecretName string  `pulumi:"logStorageAccountSasSecretName"`
	SeedBlobUri                    string  `pulumi:"seedBlobUri"`
	SeedManagedDiskId              string  `pulumi:"seedManagedDiskId"`
	TargetBlobUri                  string  `pulumi:"targetBlobUri"`
	TargetDiskName                 *string `pulumi:"targetDiskName"`
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

func (o VMwareCbtProtectedDiskDetailsResponseOutput) DiskEncryptionSetId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.DiskEncryptionSetId }).(pulumi.StringOutput)
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

func (o VMwareCbtProtectedDiskDetailsResponseOutput) SeedBlobUri() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.SeedBlobUri }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) SeedManagedDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.SeedManagedDiskId }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) TargetBlobUri() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) string { return v.TargetBlobUri }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectedDiskDetailsResponseOutput) TargetDiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtProtectedDiskDetailsResponse) *string { return v.TargetDiskName }).(pulumi.StringPtrOutput)
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
	InstanceType                         string         `pulumi:"instanceType"`
	KeyVaultId                           string         `pulumi:"keyVaultId"`
	KeyVaultUri                          string         `pulumi:"keyVaultUri"`
	RoleSizeToNicCountMap                map[string]int `pulumi:"roleSizeToNicCountMap"`
	ServiceBusConnectionStringSecretName string         `pulumi:"serviceBusConnectionStringSecretName"`
	StorageAccountId                     string         `pulumi:"storageAccountId"`
	StorageAccountSasSecretName          string         `pulumi:"storageAccountSasSecretName"`
	TargetLocation                       string         `pulumi:"targetLocation"`
}

type VMwareDetailsResponse struct {
	AgentCount                         *string                                                  `pulumi:"agentCount"`
	AgentExpiryDate                    *string                                                  `pulumi:"agentExpiryDate"`
	AgentVersion                       *string                                                  `pulumi:"agentVersion"`
	AgentVersionDetails                *VersionDetailsResponse                                  `pulumi:"agentVersionDetails"`
	AvailableMemoryInBytes             *float64                                                 `pulumi:"availableMemoryInBytes"`
	AvailableSpaceInBytes              *float64                                                 `pulumi:"availableSpaceInBytes"`
	CpuLoad                            *string                                                  `pulumi:"cpuLoad"`
	CpuLoadStatus                      *string                                                  `pulumi:"cpuLoadStatus"`
	CsServiceStatus                    *string                                                  `pulumi:"csServiceStatus"`
	DatabaseServerLoad                 *string                                                  `pulumi:"databaseServerLoad"`
	DatabaseServerLoadStatus           *string                                                  `pulumi:"databaseServerLoadStatus"`
	HostName                           *string                                                  `pulumi:"hostName"`
	InstanceType                       string                                                   `pulumi:"instanceType"`
	IpAddress                          *string                                                  `pulumi:"ipAddress"`
	LastHeartbeat                      *string                                                  `pulumi:"lastHeartbeat"`
	MasterTargetServers                []MasterTargetServerResponse                             `pulumi:"masterTargetServers"`
	MemoryUsageStatus                  *string                                                  `pulumi:"memoryUsageStatus"`
	ProcessServerCount                 *string                                                  `pulumi:"processServerCount"`
	ProcessServers                     []ProcessServerResponse                                  `pulumi:"processServers"`
	ProtectedServers                   *string                                                  `pulumi:"protectedServers"`
	PsTemplateVersion                  *string                                                  `pulumi:"psTemplateVersion"`
	ReplicationPairCount               *string                                                  `pulumi:"replicationPairCount"`
	RunAsAccounts                      []RunAsAccountResponse                                   `pulumi:"runAsAccounts"`
	SpaceUsageStatus                   *string                                                  `pulumi:"spaceUsageStatus"`
	SslCertExpiryDate                  *string                                                  `pulumi:"sslCertExpiryDate"`
	SslCertExpiryRemainingDays         *int                                                     `pulumi:"sslCertExpiryRemainingDays"`
	SwitchProviderBlockingErrorDetails []InMageFabricSwitchProviderBlockingErrorDetailsResponse `pulumi:"switchProviderBlockingErrorDetails"`
	SystemLoad                         *string                                                  `pulumi:"systemLoad"`
	SystemLoadStatus                   *string                                                  `pulumi:"systemLoadStatus"`
	TotalMemoryInBytes                 *float64                                                 `pulumi:"totalMemoryInBytes"`
	TotalSpaceInBytes                  *float64                                                 `pulumi:"totalSpaceInBytes"`
	VersionStatus                      *string                                                  `pulumi:"versionStatus"`
	WebLoad                            *string                                                  `pulumi:"webLoad"`
	WebLoadStatus                      *string                                                  `pulumi:"webLoadStatus"`
}

type VMwareV2FabricCreationInput struct {
	InstanceType        string  `pulumi:"instanceType"`
	MigrationSolutionId string  `pulumi:"migrationSolutionId"`
	PhysicalSiteId      *string `pulumi:"physicalSiteId"`
	VmwareSiteId        *string `pulumi:"vmwareSiteId"`
}

type VMwareV2FabricSpecificDetailsResponse struct {
	InstanceType        string                         `pulumi:"instanceType"`
	MigrationSolutionId string                         `pulumi:"migrationSolutionId"`
	PhysicalSiteId      string                         `pulumi:"physicalSiteId"`
	ProcessServers      []ProcessServerDetailsResponse `pulumi:"processServers"`
	ServiceContainerId  string                         `pulumi:"serviceContainerId"`
	ServiceEndpoint     string                         `pulumi:"serviceEndpoint"`
	ServiceResourceId   string                         `pulumi:"serviceResourceId"`
	VmwareSiteId        string                         `pulumi:"vmwareSiteId"`
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
	InstanceType string `pulumi:"instanceType"`
}

type VmmToAzureNetworkMappingSettingsResponse struct {
	InstanceType string `pulumi:"instanceType"`
}

type VmmToVmmCreateNetworkMappingInput struct {
	InstanceType string `pulumi:"instanceType"`
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

func init() {
	pulumi.RegisterOutputType(AddRecoveryServicesProviderInputPropertiesOutput{})
	pulumi.RegisterOutputType(AddVCenterRequestPropertiesOutput{})
	pulumi.RegisterOutputType(AddVCenterRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreateNetworkMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreatePolicyInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreatePolicyInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreateProtectionContainerMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreateProtectionContainerMappingInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreateRecoveryPlanInputPropertiesOutput{})
	pulumi.RegisterOutputType(CriticalJobHistoryDetailsResponseOutput{})
	pulumi.RegisterOutputType(CriticalJobHistoryDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(CurrentJobDetailsResponseOutput{})
	pulumi.RegisterOutputType(CurrentScenarioDetailsResponseOutput{})
	pulumi.RegisterOutputType(CurrentScenarioDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(EnableMigrationInputPropertiesOutput{})
	pulumi.RegisterOutputType(EnableProtectionInputPropertiesOutput{})
	pulumi.RegisterOutputType(EnableProtectionInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsResponseOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(FabricCreationInputPropertiesOutput{})
	pulumi.RegisterOutputType(FabricCreationInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(FabricPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HealthErrorResponseOutput{})
	pulumi.RegisterOutputType(HealthErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityProviderDetailsResponseOutput{})
	pulumi.RegisterOutputType(IdentityProviderDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderInputOutput{})
	pulumi.RegisterOutputType(IdentityProviderInputPtrOutput{})
	pulumi.RegisterOutputType(InnerHealthErrorResponseOutput{})
	pulumi.RegisterOutputType(InnerHealthErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(MigrationItemPropertiesResponseOutput{})
	pulumi.RegisterOutputType(NetworkMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PolicyPropertiesResponseOutput{})
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
	pulumi.RegisterOutputType(StorageClassificationMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(StorageMappingInputPropertiesPtrOutput{})
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
	pulumi.RegisterOutputType(VersionDetailsResponseOutput{})
	pulumi.RegisterOutputType(VersionDetailsResponsePtrOutput{})
}
