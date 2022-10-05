


package v20160810

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type A2AEnableProtectionInput struct {
	FabricObjectId            *string                        `pulumi:"fabricObjectId"`
	InstanceType              *string                        `pulumi:"instanceType"`
	RecoveryAvailabilitySetId *string                        `pulumi:"recoveryAvailabilitySetId"`
	RecoveryCloudServiceId    *string                        `pulumi:"recoveryCloudServiceId"`
	RecoveryContainerId       *string                        `pulumi:"recoveryContainerId"`
	RecoveryResourceGroupId   *string                        `pulumi:"recoveryResourceGroupId"`
	VmDisks                   []A2AVmDiskInputDetails        `pulumi:"vmDisks"`
	VmManagedDisks            []A2AVmManagedDiskInputDetails `pulumi:"vmManagedDisks"`
}

type A2APolicyCreationInput struct {
	AppConsistentFrequencyInMinutes   *int                 `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int                 `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      *string              `pulumi:"instanceType"`
	MultiVmSyncStatus                 SetMultiVmSyncStatus `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int                 `pulumi:"recoveryPointHistory"`
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
	DataPendingAtSourceAgentInMB           *float64 `pulumi:"dataPendingAtSourceAgentInMB"`
	DataPendingInStagingStorageAccountInMB *float64 `pulumi:"dataPendingInStagingStorageAccountInMB"`
	DiskCapacityInBytes                    *float64 `pulumi:"diskCapacityInBytes"`
	DiskName                               *string  `pulumi:"diskName"`
	DiskType                               *string  `pulumi:"diskType"`
	DiskUri                                *string  `pulumi:"diskUri"`
	MonitoringJobType                      *string  `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion         *int     `pulumi:"monitoringPercentageCompletion"`
	PrimaryDiskAzureStorageAccountId       *string  `pulumi:"primaryDiskAzureStorageAccountId"`
	PrimaryStagingAzureStorageAccountId    *string  `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryAzureStorageAccountId          *string  `pulumi:"recoveryAzureStorageAccountId"`
	RecoveryDiskUri                        *string  `pulumi:"recoveryDiskUri"`
	ResyncRequired                         *bool    `pulumi:"resyncRequired"`
}

type A2AProtectedManagedDiskDetailsResponse struct {
	DataPendingAtSourceAgentInMB           *float64 `pulumi:"dataPendingAtSourceAgentInMB"`
	DataPendingInStagingStorageAccountInMB *float64 `pulumi:"dataPendingInStagingStorageAccountInMB"`
	DiskCapacityInBytes                    *float64 `pulumi:"diskCapacityInBytes"`
	DiskId                                 *string  `pulumi:"diskId"`
	DiskName                               *string  `pulumi:"diskName"`
	DiskType                               *string  `pulumi:"diskType"`
	MonitoringJobType                      *string  `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion         *int     `pulumi:"monitoringPercentageCompletion"`
	PrimaryStagingAzureStorageAccountId    *string  `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryAzureResourceGroupId           *string  `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryDiskId                         *string  `pulumi:"recoveryDiskId"`
	ResyncRequired                         *bool    `pulumi:"resyncRequired"`
}

type A2AReplicationDetailsResponse struct {
	AgentVersion                       *string                                    `pulumi:"agentVersion"`
	FabricObjectId                     *string                                    `pulumi:"fabricObjectId"`
	InstanceType                       string                                     `pulumi:"instanceType"`
	IsReplicationAgentUpdateRequired   *bool                                      `pulumi:"isReplicationAgentUpdateRequired"`
	LastHeartbeat                      *string                                    `pulumi:"lastHeartbeat"`
	LastRpoCalculatedTime              *string                                    `pulumi:"lastRpoCalculatedTime"`
	LifecycleId                        *string                                    `pulumi:"lifecycleId"`
	ManagementId                       *string                                    `pulumi:"managementId"`
	MonitoringJobType                  *string                                    `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion     *int                                       `pulumi:"monitoringPercentageCompletion"`
	MultiVmGroupId                     *string                                    `pulumi:"multiVmGroupId"`
	MultiVmGroupName                   *string                                    `pulumi:"multiVmGroupName"`
	OsType                             *string                                    `pulumi:"osType"`
	PrimaryFabricLocation              *string                                    `pulumi:"primaryFabricLocation"`
	ProtectedDisks                     []A2AProtectedDiskDetailsResponse          `pulumi:"protectedDisks"`
	ProtectedManagedDisks              []A2AProtectedManagedDiskDetailsResponse   `pulumi:"protectedManagedDisks"`
	RecoveryAvailabilitySet            *string                                    `pulumi:"recoveryAvailabilitySet"`
	RecoveryAzureResourceGroupId       *string                                    `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureVMName                *string                                    `pulumi:"recoveryAzureVMName"`
	RecoveryAzureVMSize                *string                                    `pulumi:"recoveryAzureVMSize"`
	RecoveryCloudService               *string                                    `pulumi:"recoveryCloudService"`
	RecoveryFabricLocation             *string                                    `pulumi:"recoveryFabricLocation"`
	RecoveryFabricObjectId             *string                                    `pulumi:"recoveryFabricObjectId"`
	RpoInSeconds                       *float64                                   `pulumi:"rpoInSeconds"`
	SelectedRecoveryAzureNetworkId     *string                                    `pulumi:"selectedRecoveryAzureNetworkId"`
	TestFailoverRecoveryFabricObjectId *string                                    `pulumi:"testFailoverRecoveryFabricObjectId"`
	VmNics                             []VMNicDetailsResponse                     `pulumi:"vmNics"`
	VmProtectionState                  *string                                    `pulumi:"vmProtectionState"`
	VmProtectionStateDescription       *string                                    `pulumi:"vmProtectionStateDescription"`
	VmSyncedConfigDetails              *AzureToAzureVmSyncedConfigDetailsResponse `pulumi:"vmSyncedConfigDetails"`
}

type A2AVmDiskInputDetails struct {
	DiskUri                             *string `pulumi:"diskUri"`
	PrimaryStagingAzureStorageAccountId *string `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryAzureStorageAccountId       *string `pulumi:"recoveryAzureStorageAccountId"`
}

type A2AVmManagedDiskInputDetails struct {
	DiskId                              *string `pulumi:"diskId"`
	PrimaryStagingAzureStorageAccountId *string `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryResourceGroupId             *string `pulumi:"recoveryResourceGroupId"`
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

type AzureFabricSpecificDetailsResponse struct {
	ContainerIds []string `pulumi:"containerIds"`
	InstanceType string   `pulumi:"instanceType"`
	Location     *string  `pulumi:"location"`
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
	InputEndpoints  []InputEndpointResponse  `pulumi:"inputEndpoints"`
	RoleAssignments []RoleAssignmentResponse `pulumi:"roleAssignments"`
	Tags            map[string]string        `pulumi:"tags"`
}

type AzureVmDiskDetailsResponse struct {
	LunId              *string `pulumi:"lunId"`
	MaxSizeMB          *string `pulumi:"maxSizeMB"`
	TargetDiskLocation *string `pulumi:"targetDiskLocation"`
	TargetDiskName     *string `pulumi:"targetDiskName"`
	VhdId              *string `pulumi:"vhdId"`
	VhdName            *string `pulumi:"vhdName"`
	VhdType            *string `pulumi:"vhdType"`
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
	PolicyId                    *string                                           `pulumi:"policyId"`
	ProviderSpecificInput       *ReplicationProviderSpecificContainerMappingInput `pulumi:"providerSpecificInput"`
	TargetProtectionContainerId *string                                           `pulumi:"targetProtectionContainerId"`
}





type CreateProtectionContainerMappingInputPropertiesInput interface {
	pulumi.Input

	ToCreateProtectionContainerMappingInputPropertiesOutput() CreateProtectionContainerMappingInputPropertiesOutput
	ToCreateProtectionContainerMappingInputPropertiesOutputWithContext(context.Context) CreateProtectionContainerMappingInputPropertiesOutput
}

type CreateProtectionContainerMappingInputPropertiesArgs struct {
	PolicyId                    pulumi.StringPtrInput                                    `pulumi:"policyId"`
	ProviderSpecificInput       ReplicationProviderSpecificContainerMappingInputPtrInput `pulumi:"providerSpecificInput"`
	TargetProtectionContainerId pulumi.StringPtrInput                                    `pulumi:"targetProtectionContainerId"`
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

func (o CreateProtectionContainerMappingInputPropertiesOutput) ProviderSpecificInput() ReplicationProviderSpecificContainerMappingInputPtrOutput {
	return o.ApplyT(func(v CreateProtectionContainerMappingInputProperties) *ReplicationProviderSpecificContainerMappingInput {
		return v.ProviderSpecificInput
	}).(ReplicationProviderSpecificContainerMappingInputPtrOutput)
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

func (o CreateProtectionContainerMappingInputPropertiesPtrOutput) ProviderSpecificInput() ReplicationProviderSpecificContainerMappingInputPtrOutput {
	return o.ApplyT(func(v *CreateProtectionContainerMappingInputProperties) *ReplicationProviderSpecificContainerMappingInput {
		if v == nil {
			return nil
		}
		return v.ProviderSpecificInput
	}).(ReplicationProviderSpecificContainerMappingInputPtrOutput)
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
	FailoverDeploymentModel *FailoverDeploymentModel `pulumi:"failoverDeploymentModel"`
	Groups                  []RecoveryPlanGroup      `pulumi:"groups"`
	PrimaryFabricId         string                   `pulumi:"primaryFabricId"`
	RecoveryFabricId        string                   `pulumi:"recoveryFabricId"`
}





type CreateRecoveryPlanInputPropertiesInput interface {
	pulumi.Input

	ToCreateRecoveryPlanInputPropertiesOutput() CreateRecoveryPlanInputPropertiesOutput
	ToCreateRecoveryPlanInputPropertiesOutputWithContext(context.Context) CreateRecoveryPlanInputPropertiesOutput
}

type CreateRecoveryPlanInputPropertiesArgs struct {
	FailoverDeploymentModel FailoverDeploymentModelPtrInput `pulumi:"failoverDeploymentModel"`
	Groups                  RecoveryPlanGroupArrayInput     `pulumi:"groups"`
	PrimaryFabricId         pulumi.StringInput              `pulumi:"primaryFabricId"`
	RecoveryFabricId        pulumi.StringInput              `pulumi:"recoveryFabricId"`
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

func (o CreateRecoveryPlanInputPropertiesOutput) FailoverDeploymentModel() FailoverDeploymentModelPtrOutput {
	return o.ApplyT(func(v CreateRecoveryPlanInputProperties) *FailoverDeploymentModel { return v.FailoverDeploymentModel }).(FailoverDeploymentModelPtrOutput)
}

func (o CreateRecoveryPlanInputPropertiesOutput) Groups() RecoveryPlanGroupArrayOutput {
	return o.ApplyT(func(v CreateRecoveryPlanInputProperties) []RecoveryPlanGroup { return v.Groups }).(RecoveryPlanGroupArrayOutput)
}

func (o CreateRecoveryPlanInputPropertiesOutput) PrimaryFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v CreateRecoveryPlanInputProperties) string { return v.PrimaryFabricId }).(pulumi.StringOutput)
}

func (o CreateRecoveryPlanInputPropertiesOutput) RecoveryFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v CreateRecoveryPlanInputProperties) string { return v.RecoveryFabricId }).(pulumi.StringOutput)
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
	ChildErrors                  []HealthErrorResponse `pulumi:"childErrors"`
	CreationTimeUtc              *string               `pulumi:"creationTimeUtc"`
	EntityId                     *string               `pulumi:"entityId"`
	ErrorCode                    *string               `pulumi:"errorCode"`
	ErrorLevel                   *string               `pulumi:"errorLevel"`
	ErrorMessage                 *string               `pulumi:"errorMessage"`
	ErrorSource                  *string               `pulumi:"errorSource"`
	ErrorType                    *string               `pulumi:"errorType"`
	PossibleCauses               *string               `pulumi:"possibleCauses"`
	RecommendedAction            *string               `pulumi:"recommendedAction"`
	RecoveryProviderErrorMessage *string               `pulumi:"recoveryProviderErrorMessage"`
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

func (o HealthErrorResponseOutput) ChildErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v HealthErrorResponse) []HealthErrorResponse { return v.ChildErrors }).(HealthErrorResponseArrayOutput)
}

func (o HealthErrorResponseOutput) CreationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.CreationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) EntityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.EntityId }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.ErrorCode }).(pulumi.StringPtrOutput)
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

func (o HealthErrorResponseOutput) PossibleCauses() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.PossibleCauses }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) RecommendedAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.RecommendedAction }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) RecoveryProviderErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.RecoveryProviderErrorMessage }).(pulumi.StringPtrOutput)
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
	DisksToInclude               []string `pulumi:"disksToInclude"`
	EnableRDPOnTargetOption      *string  `pulumi:"enableRDPOnTargetOption"`
	HvHostVmId                   *string  `pulumi:"hvHostVmId"`
	InstanceType                 *string  `pulumi:"instanceType"`
	LogStorageAccountId          *string  `pulumi:"logStorageAccountId"`
	OsType                       *string  `pulumi:"osType"`
	TargetAzureNetworkId         *string  `pulumi:"targetAzureNetworkId"`
	TargetAzureSubnetId          *string  `pulumi:"targetAzureSubnetId"`
	TargetAzureV1ResourceGroupId *string  `pulumi:"targetAzureV1ResourceGroupId"`
	TargetAzureV2ResourceGroupId *string  `pulumi:"targetAzureV2ResourceGroupId"`
	TargetAzureVmName            *string  `pulumi:"targetAzureVmName"`
	TargetStorageAccountId       *string  `pulumi:"targetStorageAccountId"`
	UseManagedDisks              *string  `pulumi:"useManagedDisks"`
	VhdId                        *string  `pulumi:"vhdId"`
	VmName                       *string  `pulumi:"vmName"`
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
	Encryption                                    *string  `pulumi:"encryption"`
	InstanceType                                  *string  `pulumi:"instanceType"`
	OnlineReplicationStartTime                    *string  `pulumi:"onlineReplicationStartTime"`
	RecoveryPointHistoryDuration                  *int     `pulumi:"recoveryPointHistoryDuration"`
	ReplicationInterval                           *int     `pulumi:"replicationInterval"`
	StorageAccounts                               []string `pulumi:"storageAccounts"`
}

type HyperVReplicaAzureReplicationDetailsResponse struct {
	AzureVMDiskDetails               []AzureVmDiskDetailsResponse       `pulumi:"azureVMDiskDetails"`
	EnableRDPOnTargetOption          *string                            `pulumi:"enableRDPOnTargetOption"`
	Encryption                       *string                            `pulumi:"encryption"`
	InitialReplicationDetails        *InitialReplicationDetailsResponse `pulumi:"initialReplicationDetails"`
	InstanceType                     string                             `pulumi:"instanceType"`
	LastReplicatedTime               *string                            `pulumi:"lastReplicatedTime"`
	LicenseType                      *string                            `pulumi:"licenseType"`
	OSDetails                        *OSDetailsResponse                 `pulumi:"oSDetails"`
	RecoveryAvailabilitySetId        *string                            `pulumi:"recoveryAvailabilitySetId"`
	RecoveryAzureLogStorageAccountId *string                            `pulumi:"recoveryAzureLogStorageAccountId"`
	RecoveryAzureResourceGroupId     *string                            `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureStorageAccount      *string                            `pulumi:"recoveryAzureStorageAccount"`
	RecoveryAzureVMName              *string                            `pulumi:"recoveryAzureVMName"`
	RecoveryAzureVMSize              *string                            `pulumi:"recoveryAzureVMSize"`
	SelectedRecoveryAzureNetworkId   *string                            `pulumi:"selectedRecoveryAzureNetworkId"`
	SourceVmCPUCount                 *int                               `pulumi:"sourceVmCPUCount"`
	SourceVmRAMSizeInMB              *int                               `pulumi:"sourceVmRAMSizeInMB"`
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

type InMageAgentDetailsResponse struct {
	AgentUpdateStatus      *string `pulumi:"agentUpdateStatus"`
	AgentVersion           *string `pulumi:"agentVersion"`
	PostUpdateRebootStatus *string `pulumi:"postUpdateRebootStatus"`
}

type InMageAzureV2EnableProtectionInput struct {
	DisksToInclude               []string `pulumi:"disksToInclude"`
	EnableRDPOnTargetOption      *string  `pulumi:"enableRDPOnTargetOption"`
	InstanceType                 *string  `pulumi:"instanceType"`
	LogStorageAccountId          *string  `pulumi:"logStorageAccountId"`
	MasterTargetId               *string  `pulumi:"masterTargetId"`
	MultiVmGroupId               *string  `pulumi:"multiVmGroupId"`
	MultiVmGroupName             *string  `pulumi:"multiVmGroupName"`
	ProcessServerId              *string  `pulumi:"processServerId"`
	RunAsAccountId               *string  `pulumi:"runAsAccountId"`
	StorageAccountId             string   `pulumi:"storageAccountId"`
	TargetAzureNetworkId         *string  `pulumi:"targetAzureNetworkId"`
	TargetAzureSubnetId          *string  `pulumi:"targetAzureSubnetId"`
	TargetAzureV1ResourceGroupId *string  `pulumi:"targetAzureV1ResourceGroupId"`
	TargetAzureV2ResourceGroupId *string  `pulumi:"targetAzureV2ResourceGroupId"`
	TargetAzureVmName            *string  `pulumi:"targetAzureVmName"`
	UseManagedDisks              *string  `pulumi:"useManagedDisks"`
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
	AppConsistentFrequencyInMinutes   *int                 `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int                 `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      *string              `pulumi:"instanceType"`
	MultiVmSyncStatus                 SetMultiVmSyncStatus `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int                 `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int                 `pulumi:"recoveryPointThresholdInMinutes"`
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
	AgentVersion                     *string                                     `pulumi:"agentVersion"`
	AzureVMDiskDetails               []AzureVmDiskDetailsResponse                `pulumi:"azureVMDiskDetails"`
	CompressedDataRateInMB           *float64                                    `pulumi:"compressedDataRateInMB"`
	Datastores                       []string                                    `pulumi:"datastores"`
	DiscoveryType                    *string                                     `pulumi:"discoveryType"`
	DiskResized                      *string                                     `pulumi:"diskResized"`
	EnableRDPOnTargetOption          *string                                     `pulumi:"enableRDPOnTargetOption"`
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
	ProtectedDisks                   []InMageAzureV2ProtectedDiskDetailsResponse `pulumi:"protectedDisks"`
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
	SourceVmCPUCount                 *int                                        `pulumi:"sourceVmCPUCount"`
	SourceVmRAMSizeInMB              *int                                        `pulumi:"sourceVmRAMSizeInMB"`
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
	AppConsistentFrequencyInMinutes *int                 `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    *string              `pulumi:"instanceType"`
	MultiVmSyncStatus               SetMultiVmSyncStatus `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            *int                 `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes *int                 `pulumi:"recoveryPointThresholdInMinutes"`
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
	SourceVmCPUCount             *int                                 `pulumi:"sourceVmCPUCount"`
	SourceVmRAMSizeInMB          *int                                 `pulumi:"sourceVmRAMSizeInMB"`
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

type InputEndpointResponse struct {
	EndpointName *string `pulumi:"endpointName"`
	PrivatePort  *int    `pulumi:"privatePort"`
	Protocol     *string `pulumi:"protocol"`
	PublicPort   *int    `pulumi:"publicPort"`
}

type MasterTargetServerResponse struct {
	AgentVersion     *string                   `pulumi:"agentVersion"`
	DataStores       []DataStoreResponse       `pulumi:"dataStores"`
	DiskCount        *int                      `pulumi:"diskCount"`
	Id               *string                   `pulumi:"id"`
	IpAddress        *string                   `pulumi:"ipAddress"`
	LastHeartbeat    *string                   `pulumi:"lastHeartbeat"`
	Name             *string                   `pulumi:"name"`
	OsType           *string                   `pulumi:"osType"`
	OsVersion        *string                   `pulumi:"osVersion"`
	RetentionVolumes []RetentionVolumeResponse `pulumi:"retentionVolumes"`
	ValidationErrors []HealthErrorResponse     `pulumi:"validationErrors"`
	VersionStatus    *string                   `pulumi:"versionStatus"`
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

type ProcessServerResponse struct {
	AgentVersion               *string                         `pulumi:"agentVersion"`
	AvailableMemoryInBytes     *float64                        `pulumi:"availableMemoryInBytes"`
	AvailableSpaceInBytes      *float64                        `pulumi:"availableSpaceInBytes"`
	CpuLoad                    *string                         `pulumi:"cpuLoad"`
	CpuLoadStatus              *string                         `pulumi:"cpuLoadStatus"`
	FriendlyName               *string                         `pulumi:"friendlyName"`
	HostId                     *string                         `pulumi:"hostId"`
	Id                         *string                         `pulumi:"id"`
	IpAddress                  *string                         `pulumi:"ipAddress"`
	LastHeartbeat              *string                         `pulumi:"lastHeartbeat"`
	MachineCount               *string                         `pulumi:"machineCount"`
	MemoryUsageStatus          *string                         `pulumi:"memoryUsageStatus"`
	MobilityServiceUpdates     []MobilityServiceUpdateResponse `pulumi:"mobilityServiceUpdates"`
	OsType                     *string                         `pulumi:"osType"`
	OsVersion                  *string                         `pulumi:"osVersion"`
	PsServiceStatus            *string                         `pulumi:"psServiceStatus"`
	ReplicationPairCount       *string                         `pulumi:"replicationPairCount"`
	SpaceUsageStatus           *string                         `pulumi:"spaceUsageStatus"`
	SslCertExpiryDate          *string                         `pulumi:"sslCertExpiryDate"`
	SslCertExpiryRemainingDays *int                            `pulumi:"sslCertExpiryRemainingDays"`
	SystemLoad                 *string                         `pulumi:"systemLoad"`
	SystemLoadStatus           *string                         `pulumi:"systemLoadStatus"`
	TotalMemoryInBytes         *float64                        `pulumi:"totalMemoryInBytes"`
	TotalSpaceInBytes          *float64                        `pulumi:"totalSpaceInBytes"`
	VersionStatus              *string                         `pulumi:"versionStatus"`
}

type ProtectionContainerMappingPropertiesResponse struct {
	Health                                *string                                                    `pulumi:"health"`
	HealthErrorDetails                    []HealthErrorResponse                                      `pulumi:"healthErrorDetails"`
	PolicyFriendlyName                    *string                                                    `pulumi:"policyFriendlyName"`
	PolicyId                              *string                                                    `pulumi:"policyId"`
	ProviderSpecificDetails               *ProtectionContainerMappingProviderSpecificDetailsResponse `pulumi:"providerSpecificDetails"`
	SourceFabricFriendlyName              *string                                                    `pulumi:"sourceFabricFriendlyName"`
	SourceProtectionContainerFriendlyName *string                                                    `pulumi:"sourceProtectionContainerFriendlyName"`
	State                                 *string                                                    `pulumi:"state"`
	TargetFabricFriendlyName              *string                                                    `pulumi:"targetFabricFriendlyName"`
	TargetProtectionContainerFriendlyName *string                                                    `pulumi:"targetProtectionContainerFriendlyName"`
	TargetProtectionContainerId           *string                                                    `pulumi:"targetProtectionContainerId"`
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

func (o ProtectionContainerMappingPropertiesResponseOutput) ProviderSpecificDetails() ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput {
	return o.ApplyT(func(v ProtectionContainerMappingPropertiesResponse) *ProtectionContainerMappingProviderSpecificDetailsResponse {
		return v.ProviderSpecificDetails
	}).(ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput)
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

type ProtectionContainerMappingProviderSpecificDetailsResponse struct {
	InstanceType string `pulumi:"instanceType"`
}

type ProtectionContainerMappingProviderSpecificDetailsResponseOutput struct{ *pulumi.OutputState }

func (ProtectionContainerMappingProviderSpecificDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionContainerMappingProviderSpecificDetailsResponse)(nil)).Elem()
}

func (o ProtectionContainerMappingProviderSpecificDetailsResponseOutput) ToProtectionContainerMappingProviderSpecificDetailsResponseOutput() ProtectionContainerMappingProviderSpecificDetailsResponseOutput {
	return o
}

func (o ProtectionContainerMappingProviderSpecificDetailsResponseOutput) ToProtectionContainerMappingProviderSpecificDetailsResponseOutputWithContext(ctx context.Context) ProtectionContainerMappingProviderSpecificDetailsResponseOutput {
	return o
}

func (o ProtectionContainerMappingProviderSpecificDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v ProtectionContainerMappingProviderSpecificDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

type ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainerMappingProviderSpecificDetailsResponse)(nil)).Elem()
}

func (o ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput) ToProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput() ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput {
	return o
}

func (o ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput) ToProtectionContainerMappingProviderSpecificDetailsResponsePtrOutputWithContext(ctx context.Context) ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput {
	return o
}

func (o ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput) Elem() ProtectionContainerMappingProviderSpecificDetailsResponseOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingProviderSpecificDetailsResponse) ProtectionContainerMappingProviderSpecificDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ProtectionContainerMappingProviderSpecificDetailsResponse
		return ret
	}).(ProtectionContainerMappingProviderSpecificDetailsResponseOutput)
}

func (o ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingProviderSpecificDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InstanceType
	}).(pulumi.StringPtrOutput)
}

type RcmAzureMigrationPolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string  `pulumi:"instanceType"`
	MultiVmSyncStatus                 *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int    `pulumi:"recoveryPointThresholdInMinutes"`
}

type RecoveryPlanAction struct {
	ActionName         string                              `pulumi:"actionName"`
	FailoverDirections []PossibleOperationsDirections      `pulumi:"failoverDirections"`
	FailoverTypes      []ReplicationProtectedItemOperation `pulumi:"failoverTypes"`
}





type RecoveryPlanActionInput interface {
	pulumi.Input

	ToRecoveryPlanActionOutput() RecoveryPlanActionOutput
	ToRecoveryPlanActionOutputWithContext(context.Context) RecoveryPlanActionOutput
}

type RecoveryPlanActionArgs struct {
	ActionName         pulumi.StringInput                          `pulumi:"actionName"`
	FailoverDirections PossibleOperationsDirectionsArrayInput      `pulumi:"failoverDirections"`
	FailoverTypes      ReplicationProtectedItemOperationArrayInput `pulumi:"failoverTypes"`
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

func (o RecoveryPlanActionOutput) FailoverDirections() PossibleOperationsDirectionsArrayOutput {
	return o.ApplyT(func(v RecoveryPlanAction) []PossibleOperationsDirections { return v.FailoverDirections }).(PossibleOperationsDirectionsArrayOutput)
}

func (o RecoveryPlanActionOutput) FailoverTypes() ReplicationProtectedItemOperationArrayOutput {
	return o.ApplyT(func(v RecoveryPlanAction) []ReplicationProtectedItemOperation { return v.FailoverTypes }).(ReplicationProtectedItemOperationArrayOutput)
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
	GroupType                 RecoveryPlanGroupType       `pulumi:"groupType"`
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
	GroupType                 RecoveryPlanGroupTypeInput          `pulumi:"groupType"`
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

func (o RecoveryPlanGroupOutput) GroupType() RecoveryPlanGroupTypeOutput {
	return o.ApplyT(func(v RecoveryPlanGroup) RecoveryPlanGroupType { return v.GroupType }).(RecoveryPlanGroupTypeOutput)
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
	AllowedOperations                []string                        `pulumi:"allowedOperations"`
	CurrentScenario                  *CurrentScenarioDetailsResponse `pulumi:"currentScenario"`
	CurrentScenarioStatus            *string                         `pulumi:"currentScenarioStatus"`
	CurrentScenarioStatusDescription *string                         `pulumi:"currentScenarioStatusDescription"`
	FailoverDeploymentModel          *string                         `pulumi:"failoverDeploymentModel"`
	FriendlyName                     *string                         `pulumi:"friendlyName"`
	Groups                           []RecoveryPlanGroupResponse     `pulumi:"groups"`
	LastPlannedFailoverTime          *string                         `pulumi:"lastPlannedFailoverTime"`
	LastTestFailoverTime             *string                         `pulumi:"lastTestFailoverTime"`
	LastUnplannedFailoverTime        *string                         `pulumi:"lastUnplannedFailoverTime"`
	PrimaryFabricFriendlyName        *string                         `pulumi:"primaryFabricFriendlyName"`
	PrimaryFabricId                  *string                         `pulumi:"primaryFabricId"`
	RecoveryFabricFriendlyName       *string                         `pulumi:"recoveryFabricFriendlyName"`
	RecoveryFabricId                 *string                         `pulumi:"recoveryFabricId"`
	ReplicationProviders             []string                        `pulumi:"replicationProviders"`
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

type ReplicationProtectedItemPropertiesResponse struct {
	ActiveLocation                          *string                         `pulumi:"activeLocation"`
	AllowedOperations                       []string                        `pulumi:"allowedOperations"`
	CurrentScenario                         *CurrentScenarioDetailsResponse `pulumi:"currentScenario"`
	FailoverHealth                          *string                         `pulumi:"failoverHealth"`
	FailoverHealthErrors                    []HealthErrorResponse           `pulumi:"failoverHealthErrors"`
	FailoverRecoveryPointId                 *string                         `pulumi:"failoverRecoveryPointId"`
	FriendlyName                            *string                         `pulumi:"friendlyName"`
	LastSuccessfulFailoverTime              *string                         `pulumi:"lastSuccessfulFailoverTime"`
	LastSuccessfulTestFailoverTime          *string                         `pulumi:"lastSuccessfulTestFailoverTime"`
	PolicyFriendlyName                      *string                         `pulumi:"policyFriendlyName"`
	PolicyId                                *string                         `pulumi:"policyId"`
	PrimaryFabricFriendlyName               *string                         `pulumi:"primaryFabricFriendlyName"`
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
	ReplicationHealthErrors                 []HealthErrorResponse           `pulumi:"replicationHealthErrors"`
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

func (o ReplicationProtectedItemPropertiesResponseOutput) FailoverHealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) []HealthErrorResponse {
		return v.FailoverHealthErrors
	}).(HealthErrorResponseArrayOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) FailoverRecoveryPointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.FailoverRecoveryPointId }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
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

func (o ReplicationProtectedItemPropertiesResponseOutput) ReplicationHealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) []HealthErrorResponse {
		return v.ReplicationHealthErrors
	}).(HealthErrorResponseArrayOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) TestFailoverState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.TestFailoverState }).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponseOutput) TestFailoverStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProtectedItemPropertiesResponse) *string { return v.TestFailoverStateDescription }).(pulumi.StringPtrOutput)
}

type ReplicationProviderSpecificContainerMappingInput struct {
	InstanceType *string `pulumi:"instanceType"`
}





type ReplicationProviderSpecificContainerMappingInputInput interface {
	pulumi.Input

	ToReplicationProviderSpecificContainerMappingInputOutput() ReplicationProviderSpecificContainerMappingInputOutput
	ToReplicationProviderSpecificContainerMappingInputOutputWithContext(context.Context) ReplicationProviderSpecificContainerMappingInputOutput
}

type ReplicationProviderSpecificContainerMappingInputArgs struct {
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
}

func (ReplicationProviderSpecificContainerMappingInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationProviderSpecificContainerMappingInput)(nil)).Elem()
}

func (i ReplicationProviderSpecificContainerMappingInputArgs) ToReplicationProviderSpecificContainerMappingInputOutput() ReplicationProviderSpecificContainerMappingInputOutput {
	return i.ToReplicationProviderSpecificContainerMappingInputOutputWithContext(context.Background())
}

func (i ReplicationProviderSpecificContainerMappingInputArgs) ToReplicationProviderSpecificContainerMappingInputOutputWithContext(ctx context.Context) ReplicationProviderSpecificContainerMappingInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProviderSpecificContainerMappingInputOutput)
}

func (i ReplicationProviderSpecificContainerMappingInputArgs) ToReplicationProviderSpecificContainerMappingInputPtrOutput() ReplicationProviderSpecificContainerMappingInputPtrOutput {
	return i.ToReplicationProviderSpecificContainerMappingInputPtrOutputWithContext(context.Background())
}

func (i ReplicationProviderSpecificContainerMappingInputArgs) ToReplicationProviderSpecificContainerMappingInputPtrOutputWithContext(ctx context.Context) ReplicationProviderSpecificContainerMappingInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProviderSpecificContainerMappingInputOutput).ToReplicationProviderSpecificContainerMappingInputPtrOutputWithContext(ctx)
}









type ReplicationProviderSpecificContainerMappingInputPtrInput interface {
	pulumi.Input

	ToReplicationProviderSpecificContainerMappingInputPtrOutput() ReplicationProviderSpecificContainerMappingInputPtrOutput
	ToReplicationProviderSpecificContainerMappingInputPtrOutputWithContext(context.Context) ReplicationProviderSpecificContainerMappingInputPtrOutput
}

type replicationProviderSpecificContainerMappingInputPtrType ReplicationProviderSpecificContainerMappingInputArgs

func ReplicationProviderSpecificContainerMappingInputPtr(v *ReplicationProviderSpecificContainerMappingInputArgs) ReplicationProviderSpecificContainerMappingInputPtrInput {
	return (*replicationProviderSpecificContainerMappingInputPtrType)(v)
}

func (*replicationProviderSpecificContainerMappingInputPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProviderSpecificContainerMappingInput)(nil)).Elem()
}

func (i *replicationProviderSpecificContainerMappingInputPtrType) ToReplicationProviderSpecificContainerMappingInputPtrOutput() ReplicationProviderSpecificContainerMappingInputPtrOutput {
	return i.ToReplicationProviderSpecificContainerMappingInputPtrOutputWithContext(context.Background())
}

func (i *replicationProviderSpecificContainerMappingInputPtrType) ToReplicationProviderSpecificContainerMappingInputPtrOutputWithContext(ctx context.Context) ReplicationProviderSpecificContainerMappingInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProviderSpecificContainerMappingInputPtrOutput)
}

type ReplicationProviderSpecificContainerMappingInputOutput struct{ *pulumi.OutputState }

func (ReplicationProviderSpecificContainerMappingInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationProviderSpecificContainerMappingInput)(nil)).Elem()
}

func (o ReplicationProviderSpecificContainerMappingInputOutput) ToReplicationProviderSpecificContainerMappingInputOutput() ReplicationProviderSpecificContainerMappingInputOutput {
	return o
}

func (o ReplicationProviderSpecificContainerMappingInputOutput) ToReplicationProviderSpecificContainerMappingInputOutputWithContext(ctx context.Context) ReplicationProviderSpecificContainerMappingInputOutput {
	return o
}

func (o ReplicationProviderSpecificContainerMappingInputOutput) ToReplicationProviderSpecificContainerMappingInputPtrOutput() ReplicationProviderSpecificContainerMappingInputPtrOutput {
	return o.ToReplicationProviderSpecificContainerMappingInputPtrOutputWithContext(context.Background())
}

func (o ReplicationProviderSpecificContainerMappingInputOutput) ToReplicationProviderSpecificContainerMappingInputPtrOutputWithContext(ctx context.Context) ReplicationProviderSpecificContainerMappingInputPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationProviderSpecificContainerMappingInput) *ReplicationProviderSpecificContainerMappingInput {
		return &v
	}).(ReplicationProviderSpecificContainerMappingInputPtrOutput)
}

func (o ReplicationProviderSpecificContainerMappingInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReplicationProviderSpecificContainerMappingInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

type ReplicationProviderSpecificContainerMappingInputPtrOutput struct{ *pulumi.OutputState }

func (ReplicationProviderSpecificContainerMappingInputPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProviderSpecificContainerMappingInput)(nil)).Elem()
}

func (o ReplicationProviderSpecificContainerMappingInputPtrOutput) ToReplicationProviderSpecificContainerMappingInputPtrOutput() ReplicationProviderSpecificContainerMappingInputPtrOutput {
	return o
}

func (o ReplicationProviderSpecificContainerMappingInputPtrOutput) ToReplicationProviderSpecificContainerMappingInputPtrOutputWithContext(ctx context.Context) ReplicationProviderSpecificContainerMappingInputPtrOutput {
	return o
}

func (o ReplicationProviderSpecificContainerMappingInputPtrOutput) Elem() ReplicationProviderSpecificContainerMappingInputOutput {
	return o.ApplyT(func(v *ReplicationProviderSpecificContainerMappingInput) ReplicationProviderSpecificContainerMappingInput {
		if v != nil {
			return *v
		}
		var ret ReplicationProviderSpecificContainerMappingInput
		return ret
	}).(ReplicationProviderSpecificContainerMappingInputOutput)
}

func (o ReplicationProviderSpecificContainerMappingInputPtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProviderSpecificContainerMappingInput) *string {
		if v == nil {
			return nil
		}
		return v.InstanceType
	}).(pulumi.StringPtrOutput)
}

type RetentionVolumeResponse struct {
	CapacityInBytes     *float64 `pulumi:"capacityInBytes"`
	FreeSpaceInBytes    *float64 `pulumi:"freeSpaceInBytes"`
	ThresholdPercentage *int     `pulumi:"thresholdPercentage"`
	VolumeName          *string  `pulumi:"volumeName"`
}

type RoleAssignmentResponse struct {
	Id               *string `pulumi:"id"`
	Name             *string `pulumi:"name"`
	PrincipalId      *string `pulumi:"principalId"`
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	Scope            *string `pulumi:"scope"`
}

type RunAsAccountResponse struct {
	AccountId   *string `pulumi:"accountId"`
	AccountName *string `pulumi:"accountName"`
}

type SanEnableProtectionInput struct {
	InstanceType *string `pulumi:"instanceType"`
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
	DiscoveryStatus       *string `pulumi:"discoveryStatus"`
	FabricArmResourceName *string `pulumi:"fabricArmResourceName"`
	FriendlyName          *string `pulumi:"friendlyName"`
	InfrastructureId      *string `pulumi:"infrastructureId"`
	InternalId            *string `pulumi:"internalId"`
	IpAddress             *string `pulumi:"ipAddress"`
	LastHeartbeat         *string `pulumi:"lastHeartbeat"`
	Port                  *string `pulumi:"port"`
	ProcessServerId       *string `pulumi:"processServerId"`
	RunAsAccountId        *string `pulumi:"runAsAccountId"`
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
	IpAddressType             *string `pulumi:"ipAddressType"`
	NicId                     *string `pulumi:"nicId"`
	PrimaryNicStaticIPAddress *string `pulumi:"primaryNicStaticIPAddress"`
	RecoveryNicIpAddressType  *string `pulumi:"recoveryNicIpAddressType"`
	RecoveryVMNetworkId       *string `pulumi:"recoveryVMNetworkId"`
	RecoveryVMSubnetName      *string `pulumi:"recoveryVMSubnetName"`
	ReplicaNicId              *string `pulumi:"replicaNicId"`
	ReplicaNicStaticIPAddress *string `pulumi:"replicaNicStaticIPAddress"`
	SelectionType             *string `pulumi:"selectionType"`
	SourceNicArmId            *string `pulumi:"sourceNicArmId"`
	VMNetworkName             *string `pulumi:"vMNetworkName"`
	VMSubnetName              *string `pulumi:"vMSubnetName"`
}

type VMwareCbtPolicyCreationInput struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      *string `pulumi:"instanceType"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
}

type VMwareDetailsResponse struct {
	AgentCount                 *string                      `pulumi:"agentCount"`
	AgentVersion               *string                      `pulumi:"agentVersion"`
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

type VMwareV2FabricSpecificDetailsResponse struct {
	InstanceType       string  `pulumi:"instanceType"`
	RcmServiceEndpoint *string `pulumi:"rcmServiceEndpoint"`
	SrsServiceEndpoint *string `pulumi:"srsServiceEndpoint"`
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
	RecoveryPointHistory              *int   `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int   `pulumi:"recoveryPointThresholdInMinutes"`
}

func init() {
	pulumi.RegisterOutputType(AddVCenterRequestPropertiesOutput{})
	pulumi.RegisterOutputType(AddVCenterRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreateNetworkMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreateNetworkMappingInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreatePolicyInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreatePolicyInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreateProtectionContainerMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreateProtectionContainerMappingInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreateRecoveryPlanInputPropertiesOutput{})
	pulumi.RegisterOutputType(CurrentScenarioDetailsResponseOutput{})
	pulumi.RegisterOutputType(CurrentScenarioDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(EnableProtectionInputPropertiesOutput{})
	pulumi.RegisterOutputType(EnableProtectionInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsResponseOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(FabricPropertiesResponseOutput{})
	pulumi.RegisterOutputType(HealthErrorResponseOutput{})
	pulumi.RegisterOutputType(HealthErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ProtectionContainerMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ProtectionContainerMappingProviderSpecificDetailsResponseOutput{})
	pulumi.RegisterOutputType(ProtectionContainerMappingProviderSpecificDetailsResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ReplicationProtectedItemPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ReplicationProviderSpecificContainerMappingInputOutput{})
	pulumi.RegisterOutputType(ReplicationProviderSpecificContainerMappingInputPtrOutput{})
	pulumi.RegisterOutputType(StorageClassificationMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(StorageMappingInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VCenterPropertiesResponseOutput{})
}
