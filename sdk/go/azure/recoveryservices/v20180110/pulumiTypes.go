


package v20180110

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





type A2AContainerMappingInputInput interface {
	pulumi.Input

	ToA2AContainerMappingInputOutput() A2AContainerMappingInputOutput
	ToA2AContainerMappingInputOutputWithContext(context.Context) A2AContainerMappingInputOutput
}

type A2AContainerMappingInputArgs struct {
	AgentAutoUpdateStatus  pulumi.StringPtrInput `pulumi:"agentAutoUpdateStatus"`
	AutomationAccountArmId pulumi.StringPtrInput `pulumi:"automationAccountArmId"`
	InstanceType           pulumi.StringPtrInput `pulumi:"instanceType"`
}

func (A2AContainerMappingInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AContainerMappingInput)(nil)).Elem()
}

func (i A2AContainerMappingInputArgs) ToA2AContainerMappingInputOutput() A2AContainerMappingInputOutput {
	return i.ToA2AContainerMappingInputOutputWithContext(context.Background())
}

func (i A2AContainerMappingInputArgs) ToA2AContainerMappingInputOutputWithContext(ctx context.Context) A2AContainerMappingInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AContainerMappingInputOutput)
}

type A2AContainerMappingInputOutput struct{ *pulumi.OutputState }

func (A2AContainerMappingInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AContainerMappingInput)(nil)).Elem()
}

func (o A2AContainerMappingInputOutput) ToA2AContainerMappingInputOutput() A2AContainerMappingInputOutput {
	return o
}

func (o A2AContainerMappingInputOutput) ToA2AContainerMappingInputOutputWithContext(ctx context.Context) A2AContainerMappingInputOutput {
	return o
}

func (o A2AContainerMappingInputOutput) AgentAutoUpdateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AContainerMappingInput) *string { return v.AgentAutoUpdateStatus }).(pulumi.StringPtrOutput)
}

func (o A2AContainerMappingInputOutput) AutomationAccountArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AContainerMappingInput) *string { return v.AutomationAccountArmId }).(pulumi.StringPtrOutput)
}

func (o A2AContainerMappingInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AContainerMappingInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

type A2AEnableProtectionInput struct {
	DiskEncryptionInfo               *DiskEncryptionInfo            `pulumi:"diskEncryptionInfo"`
	FabricObjectId                   *string                        `pulumi:"fabricObjectId"`
	InstanceType                     *string                        `pulumi:"instanceType"`
	MultiVmGroupName                 *string                        `pulumi:"multiVmGroupName"`
	RecoveryAvailabilitySetId        *string                        `pulumi:"recoveryAvailabilitySetId"`
	RecoveryBootDiagStorageAccountId *string                        `pulumi:"recoveryBootDiagStorageAccountId"`
	RecoveryCloudServiceId           *string                        `pulumi:"recoveryCloudServiceId"`
	RecoveryContainerId              *string                        `pulumi:"recoveryContainerId"`
	RecoveryResourceGroupId          *string                        `pulumi:"recoveryResourceGroupId"`
	VmDisks                          []A2AVmDiskInputDetails        `pulumi:"vmDisks"`
	VmManagedDisks                   []A2AVmManagedDiskInputDetails `pulumi:"vmManagedDisks"`
}





type A2AEnableProtectionInputInput interface {
	pulumi.Input

	ToA2AEnableProtectionInputOutput() A2AEnableProtectionInputOutput
	ToA2AEnableProtectionInputOutputWithContext(context.Context) A2AEnableProtectionInputOutput
}

type A2AEnableProtectionInputArgs struct {
	DiskEncryptionInfo               DiskEncryptionInfoPtrInput             `pulumi:"diskEncryptionInfo"`
	FabricObjectId                   pulumi.StringPtrInput                  `pulumi:"fabricObjectId"`
	InstanceType                     pulumi.StringPtrInput                  `pulumi:"instanceType"`
	MultiVmGroupName                 pulumi.StringPtrInput                  `pulumi:"multiVmGroupName"`
	RecoveryAvailabilitySetId        pulumi.StringPtrInput                  `pulumi:"recoveryAvailabilitySetId"`
	RecoveryBootDiagStorageAccountId pulumi.StringPtrInput                  `pulumi:"recoveryBootDiagStorageAccountId"`
	RecoveryCloudServiceId           pulumi.StringPtrInput                  `pulumi:"recoveryCloudServiceId"`
	RecoveryContainerId              pulumi.StringPtrInput                  `pulumi:"recoveryContainerId"`
	RecoveryResourceGroupId          pulumi.StringPtrInput                  `pulumi:"recoveryResourceGroupId"`
	VmDisks                          A2AVmDiskInputDetailsArrayInput        `pulumi:"vmDisks"`
	VmManagedDisks                   A2AVmManagedDiskInputDetailsArrayInput `pulumi:"vmManagedDisks"`
}

func (A2AEnableProtectionInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AEnableProtectionInput)(nil)).Elem()
}

func (i A2AEnableProtectionInputArgs) ToA2AEnableProtectionInputOutput() A2AEnableProtectionInputOutput {
	return i.ToA2AEnableProtectionInputOutputWithContext(context.Background())
}

func (i A2AEnableProtectionInputArgs) ToA2AEnableProtectionInputOutputWithContext(ctx context.Context) A2AEnableProtectionInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AEnableProtectionInputOutput)
}

type A2AEnableProtectionInputOutput struct{ *pulumi.OutputState }

func (A2AEnableProtectionInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AEnableProtectionInput)(nil)).Elem()
}

func (o A2AEnableProtectionInputOutput) ToA2AEnableProtectionInputOutput() A2AEnableProtectionInputOutput {
	return o
}

func (o A2AEnableProtectionInputOutput) ToA2AEnableProtectionInputOutputWithContext(ctx context.Context) A2AEnableProtectionInputOutput {
	return o
}

func (o A2AEnableProtectionInputOutput) DiskEncryptionInfo() DiskEncryptionInfoPtrOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) *DiskEncryptionInfo { return v.DiskEncryptionInfo }).(DiskEncryptionInfoPtrOutput)
}

func (o A2AEnableProtectionInputOutput) FabricObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) *string { return v.FabricObjectId }).(pulumi.StringPtrOutput)
}

func (o A2AEnableProtectionInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o A2AEnableProtectionInputOutput) MultiVmGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) *string { return v.MultiVmGroupName }).(pulumi.StringPtrOutput)
}

func (o A2AEnableProtectionInputOutput) RecoveryAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) *string { return v.RecoveryAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o A2AEnableProtectionInputOutput) RecoveryBootDiagStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) *string { return v.RecoveryBootDiagStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o A2AEnableProtectionInputOutput) RecoveryCloudServiceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) *string { return v.RecoveryCloudServiceId }).(pulumi.StringPtrOutput)
}

func (o A2AEnableProtectionInputOutput) RecoveryContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) *string { return v.RecoveryContainerId }).(pulumi.StringPtrOutput)
}

func (o A2AEnableProtectionInputOutput) RecoveryResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) *string { return v.RecoveryResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o A2AEnableProtectionInputOutput) VmDisks() A2AVmDiskInputDetailsArrayOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) []A2AVmDiskInputDetails { return v.VmDisks }).(A2AVmDiskInputDetailsArrayOutput)
}

func (o A2AEnableProtectionInputOutput) VmManagedDisks() A2AVmManagedDiskInputDetailsArrayOutput {
	return o.ApplyT(func(v A2AEnableProtectionInput) []A2AVmManagedDiskInputDetails { return v.VmManagedDisks }).(A2AVmManagedDiskInputDetailsArrayOutput)
}

type A2APolicyCreationInput struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      *string `pulumi:"instanceType"`
	MultiVmSyncStatus                 string  `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
}





type A2APolicyCreationInputInput interface {
	pulumi.Input

	ToA2APolicyCreationInputOutput() A2APolicyCreationInputOutput
	ToA2APolicyCreationInputOutputWithContext(context.Context) A2APolicyCreationInputOutput
}

type A2APolicyCreationInputArgs struct {
	AppConsistentFrequencyInMinutes   pulumi.IntPtrInput    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes pulumi.IntPtrInput    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      pulumi.StringPtrInput `pulumi:"instanceType"`
	MultiVmSyncStatus                 pulumi.StringInput    `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              pulumi.IntPtrInput    `pulumi:"recoveryPointHistory"`
}

func (A2APolicyCreationInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*A2APolicyCreationInput)(nil)).Elem()
}

func (i A2APolicyCreationInputArgs) ToA2APolicyCreationInputOutput() A2APolicyCreationInputOutput {
	return i.ToA2APolicyCreationInputOutputWithContext(context.Background())
}

func (i A2APolicyCreationInputArgs) ToA2APolicyCreationInputOutputWithContext(ctx context.Context) A2APolicyCreationInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2APolicyCreationInputOutput)
}

type A2APolicyCreationInputOutput struct{ *pulumi.OutputState }

func (A2APolicyCreationInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*A2APolicyCreationInput)(nil)).Elem()
}

func (o A2APolicyCreationInputOutput) ToA2APolicyCreationInputOutput() A2APolicyCreationInputOutput {
	return o
}

func (o A2APolicyCreationInputOutput) ToA2APolicyCreationInputOutputWithContext(ctx context.Context) A2APolicyCreationInputOutput {
	return o
}

func (o A2APolicyCreationInputOutput) AppConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v A2APolicyCreationInput) *int { return v.AppConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o A2APolicyCreationInputOutput) CrashConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v A2APolicyCreationInput) *int { return v.CrashConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o A2APolicyCreationInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2APolicyCreationInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o A2APolicyCreationInputOutput) MultiVmSyncStatus() pulumi.StringOutput {
	return o.ApplyT(func(v A2APolicyCreationInput) string { return v.MultiVmSyncStatus }).(pulumi.StringOutput)
}

func (o A2APolicyCreationInputOutput) RecoveryPointHistory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v A2APolicyCreationInput) *int { return v.RecoveryPointHistory }).(pulumi.IntPtrOutput)
}

type A2APolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string  `pulumi:"instanceType"`
	MultiVmSyncStatus                 *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int    `pulumi:"recoveryPointThresholdInMinutes"`
}





type A2APolicyDetailsResponseInput interface {
	pulumi.Input

	ToA2APolicyDetailsResponseOutput() A2APolicyDetailsResponseOutput
	ToA2APolicyDetailsResponseOutputWithContext(context.Context) A2APolicyDetailsResponseOutput
}

type A2APolicyDetailsResponseArgs struct {
	AppConsistentFrequencyInMinutes   pulumi.IntPtrInput    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes pulumi.IntPtrInput    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      pulumi.StringInput    `pulumi:"instanceType"`
	MultiVmSyncStatus                 pulumi.StringPtrInput `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              pulumi.IntPtrInput    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   pulumi.IntPtrInput    `pulumi:"recoveryPointThresholdInMinutes"`
}

func (A2APolicyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*A2APolicyDetailsResponse)(nil)).Elem()
}

func (i A2APolicyDetailsResponseArgs) ToA2APolicyDetailsResponseOutput() A2APolicyDetailsResponseOutput {
	return i.ToA2APolicyDetailsResponseOutputWithContext(context.Background())
}

func (i A2APolicyDetailsResponseArgs) ToA2APolicyDetailsResponseOutputWithContext(ctx context.Context) A2APolicyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2APolicyDetailsResponseOutput)
}

type A2APolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (A2APolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*A2APolicyDetailsResponse)(nil)).Elem()
}

func (o A2APolicyDetailsResponseOutput) ToA2APolicyDetailsResponseOutput() A2APolicyDetailsResponseOutput {
	return o
}

func (o A2APolicyDetailsResponseOutput) ToA2APolicyDetailsResponseOutputWithContext(ctx context.Context) A2APolicyDetailsResponseOutput {
	return o
}

func (o A2APolicyDetailsResponseOutput) AppConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v A2APolicyDetailsResponse) *int { return v.AppConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o A2APolicyDetailsResponseOutput) CrashConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v A2APolicyDetailsResponse) *int { return v.CrashConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o A2APolicyDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v A2APolicyDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o A2APolicyDetailsResponseOutput) MultiVmSyncStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2APolicyDetailsResponse) *string { return v.MultiVmSyncStatus }).(pulumi.StringPtrOutput)
}

func (o A2APolicyDetailsResponseOutput) RecoveryPointHistory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v A2APolicyDetailsResponse) *int { return v.RecoveryPointHistory }).(pulumi.IntPtrOutput)
}

func (o A2APolicyDetailsResponseOutput) RecoveryPointThresholdInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v A2APolicyDetailsResponse) *int { return v.RecoveryPointThresholdInMinutes }).(pulumi.IntPtrOutput)
}

type A2AProtectedDiskDetailsResponse struct {
	DataPendingAtSourceAgentInMB           *float64 `pulumi:"dataPendingAtSourceAgentInMB"`
	DataPendingInStagingStorageAccountInMB *float64 `pulumi:"dataPendingInStagingStorageAccountInMB"`
	DekKeyVaultArmId                       *string  `pulumi:"dekKeyVaultArmId"`
	DiskCapacityInBytes                    *float64 `pulumi:"diskCapacityInBytes"`
	DiskName                               *string  `pulumi:"diskName"`
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





type A2AProtectedDiskDetailsResponseInput interface {
	pulumi.Input

	ToA2AProtectedDiskDetailsResponseOutput() A2AProtectedDiskDetailsResponseOutput
	ToA2AProtectedDiskDetailsResponseOutputWithContext(context.Context) A2AProtectedDiskDetailsResponseOutput
}

type A2AProtectedDiskDetailsResponseArgs struct {
	DataPendingAtSourceAgentInMB           pulumi.Float64PtrInput `pulumi:"dataPendingAtSourceAgentInMB"`
	DataPendingInStagingStorageAccountInMB pulumi.Float64PtrInput `pulumi:"dataPendingInStagingStorageAccountInMB"`
	DekKeyVaultArmId                       pulumi.StringPtrInput  `pulumi:"dekKeyVaultArmId"`
	DiskCapacityInBytes                    pulumi.Float64PtrInput `pulumi:"diskCapacityInBytes"`
	DiskName                               pulumi.StringPtrInput  `pulumi:"diskName"`
	DiskType                               pulumi.StringPtrInput  `pulumi:"diskType"`
	DiskUri                                pulumi.StringPtrInput  `pulumi:"diskUri"`
	IsDiskEncrypted                        pulumi.BoolPtrInput    `pulumi:"isDiskEncrypted"`
	IsDiskKeyEncrypted                     pulumi.BoolPtrInput    `pulumi:"isDiskKeyEncrypted"`
	KekKeyVaultArmId                       pulumi.StringPtrInput  `pulumi:"kekKeyVaultArmId"`
	KeyIdentifier                          pulumi.StringPtrInput  `pulumi:"keyIdentifier"`
	MonitoringJobType                      pulumi.StringPtrInput  `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion         pulumi.IntPtrInput     `pulumi:"monitoringPercentageCompletion"`
	PrimaryDiskAzureStorageAccountId       pulumi.StringPtrInput  `pulumi:"primaryDiskAzureStorageAccountId"`
	PrimaryStagingAzureStorageAccountId    pulumi.StringPtrInput  `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryAzureStorageAccountId          pulumi.StringPtrInput  `pulumi:"recoveryAzureStorageAccountId"`
	RecoveryDiskUri                        pulumi.StringPtrInput  `pulumi:"recoveryDiskUri"`
	ResyncRequired                         pulumi.BoolPtrInput    `pulumi:"resyncRequired"`
	SecretIdentifier                       pulumi.StringPtrInput  `pulumi:"secretIdentifier"`
}

func (A2AProtectedDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AProtectedDiskDetailsResponse)(nil)).Elem()
}

func (i A2AProtectedDiskDetailsResponseArgs) ToA2AProtectedDiskDetailsResponseOutput() A2AProtectedDiskDetailsResponseOutput {
	return i.ToA2AProtectedDiskDetailsResponseOutputWithContext(context.Background())
}

func (i A2AProtectedDiskDetailsResponseArgs) ToA2AProtectedDiskDetailsResponseOutputWithContext(ctx context.Context) A2AProtectedDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AProtectedDiskDetailsResponseOutput)
}





type A2AProtectedDiskDetailsResponseArrayInput interface {
	pulumi.Input

	ToA2AProtectedDiskDetailsResponseArrayOutput() A2AProtectedDiskDetailsResponseArrayOutput
	ToA2AProtectedDiskDetailsResponseArrayOutputWithContext(context.Context) A2AProtectedDiskDetailsResponseArrayOutput
}

type A2AProtectedDiskDetailsResponseArray []A2AProtectedDiskDetailsResponseInput

func (A2AProtectedDiskDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]A2AProtectedDiskDetailsResponse)(nil)).Elem()
}

func (i A2AProtectedDiskDetailsResponseArray) ToA2AProtectedDiskDetailsResponseArrayOutput() A2AProtectedDiskDetailsResponseArrayOutput {
	return i.ToA2AProtectedDiskDetailsResponseArrayOutputWithContext(context.Background())
}

func (i A2AProtectedDiskDetailsResponseArray) ToA2AProtectedDiskDetailsResponseArrayOutputWithContext(ctx context.Context) A2AProtectedDiskDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AProtectedDiskDetailsResponseArrayOutput)
}

type A2AProtectedDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (A2AProtectedDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AProtectedDiskDetailsResponse)(nil)).Elem()
}

func (o A2AProtectedDiskDetailsResponseOutput) ToA2AProtectedDiskDetailsResponseOutput() A2AProtectedDiskDetailsResponseOutput {
	return o
}

func (o A2AProtectedDiskDetailsResponseOutput) ToA2AProtectedDiskDetailsResponseOutputWithContext(ctx context.Context) A2AProtectedDiskDetailsResponseOutput {
	return o
}

func (o A2AProtectedDiskDetailsResponseOutput) DataPendingAtSourceAgentInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *float64 { return v.DataPendingAtSourceAgentInMB }).(pulumi.Float64PtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) DataPendingInStagingStorageAccountInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *float64 { return v.DataPendingInStagingStorageAccountInMB }).(pulumi.Float64PtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) DekKeyVaultArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.DekKeyVaultArmId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) DiskCapacityInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *float64 { return v.DiskCapacityInBytes }).(pulumi.Float64PtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) DiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.DiskName }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) DiskUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.DiskUri }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) IsDiskEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *bool { return v.IsDiskEncrypted }).(pulumi.BoolPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) IsDiskKeyEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *bool { return v.IsDiskKeyEncrypted }).(pulumi.BoolPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) KekKeyVaultArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.KekKeyVaultArmId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) MonitoringJobType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.MonitoringJobType }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) MonitoringPercentageCompletion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *int { return v.MonitoringPercentageCompletion }).(pulumi.IntPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) PrimaryDiskAzureStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.PrimaryDiskAzureStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) PrimaryStagingAzureStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.PrimaryStagingAzureStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) RecoveryAzureStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.RecoveryAzureStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) RecoveryDiskUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.RecoveryDiskUri }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) ResyncRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *bool { return v.ResyncRequired }).(pulumi.BoolPtrOutput)
}

func (o A2AProtectedDiskDetailsResponseOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedDiskDetailsResponse) *string { return v.SecretIdentifier }).(pulumi.StringPtrOutput)
}

type A2AProtectedDiskDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (A2AProtectedDiskDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]A2AProtectedDiskDetailsResponse)(nil)).Elem()
}

func (o A2AProtectedDiskDetailsResponseArrayOutput) ToA2AProtectedDiskDetailsResponseArrayOutput() A2AProtectedDiskDetailsResponseArrayOutput {
	return o
}

func (o A2AProtectedDiskDetailsResponseArrayOutput) ToA2AProtectedDiskDetailsResponseArrayOutputWithContext(ctx context.Context) A2AProtectedDiskDetailsResponseArrayOutput {
	return o
}

func (o A2AProtectedDiskDetailsResponseArrayOutput) Index(i pulumi.IntInput) A2AProtectedDiskDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) A2AProtectedDiskDetailsResponse {
		return vs[0].([]A2AProtectedDiskDetailsResponse)[vs[1].(int)]
	}).(A2AProtectedDiskDetailsResponseOutput)
}

type A2AProtectedManagedDiskDetailsResponse struct {
	DataPendingAtSourceAgentInMB           *float64 `pulumi:"dataPendingAtSourceAgentInMB"`
	DataPendingInStagingStorageAccountInMB *float64 `pulumi:"dataPendingInStagingStorageAccountInMB"`
	DekKeyVaultArmId                       *string  `pulumi:"dekKeyVaultArmId"`
	DiskCapacityInBytes                    *float64 `pulumi:"diskCapacityInBytes"`
	DiskId                                 *string  `pulumi:"diskId"`
	DiskName                               *string  `pulumi:"diskName"`
	DiskType                               *string  `pulumi:"diskType"`
	IsDiskEncrypted                        *bool    `pulumi:"isDiskEncrypted"`
	IsDiskKeyEncrypted                     *bool    `pulumi:"isDiskKeyEncrypted"`
	KekKeyVaultArmId                       *string  `pulumi:"kekKeyVaultArmId"`
	KeyIdentifier                          *string  `pulumi:"keyIdentifier"`
	MonitoringJobType                      *string  `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion         *int     `pulumi:"monitoringPercentageCompletion"`
	PrimaryStagingAzureStorageAccountId    *string  `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryReplicaDiskAccountType         *string  `pulumi:"recoveryReplicaDiskAccountType"`
	RecoveryReplicaDiskId                  *string  `pulumi:"recoveryReplicaDiskId"`
	RecoveryResourceGroupId                *string  `pulumi:"recoveryResourceGroupId"`
	RecoveryTargetDiskAccountType          *string  `pulumi:"recoveryTargetDiskAccountType"`
	RecoveryTargetDiskId                   *string  `pulumi:"recoveryTargetDiskId"`
	ResyncRequired                         *bool    `pulumi:"resyncRequired"`
	SecretIdentifier                       *string  `pulumi:"secretIdentifier"`
}





type A2AProtectedManagedDiskDetailsResponseInput interface {
	pulumi.Input

	ToA2AProtectedManagedDiskDetailsResponseOutput() A2AProtectedManagedDiskDetailsResponseOutput
	ToA2AProtectedManagedDiskDetailsResponseOutputWithContext(context.Context) A2AProtectedManagedDiskDetailsResponseOutput
}

type A2AProtectedManagedDiskDetailsResponseArgs struct {
	DataPendingAtSourceAgentInMB           pulumi.Float64PtrInput `pulumi:"dataPendingAtSourceAgentInMB"`
	DataPendingInStagingStorageAccountInMB pulumi.Float64PtrInput `pulumi:"dataPendingInStagingStorageAccountInMB"`
	DekKeyVaultArmId                       pulumi.StringPtrInput  `pulumi:"dekKeyVaultArmId"`
	DiskCapacityInBytes                    pulumi.Float64PtrInput `pulumi:"diskCapacityInBytes"`
	DiskId                                 pulumi.StringPtrInput  `pulumi:"diskId"`
	DiskName                               pulumi.StringPtrInput  `pulumi:"diskName"`
	DiskType                               pulumi.StringPtrInput  `pulumi:"diskType"`
	IsDiskEncrypted                        pulumi.BoolPtrInput    `pulumi:"isDiskEncrypted"`
	IsDiskKeyEncrypted                     pulumi.BoolPtrInput    `pulumi:"isDiskKeyEncrypted"`
	KekKeyVaultArmId                       pulumi.StringPtrInput  `pulumi:"kekKeyVaultArmId"`
	KeyIdentifier                          pulumi.StringPtrInput  `pulumi:"keyIdentifier"`
	MonitoringJobType                      pulumi.StringPtrInput  `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion         pulumi.IntPtrInput     `pulumi:"monitoringPercentageCompletion"`
	PrimaryStagingAzureStorageAccountId    pulumi.StringPtrInput  `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryReplicaDiskAccountType         pulumi.StringPtrInput  `pulumi:"recoveryReplicaDiskAccountType"`
	RecoveryReplicaDiskId                  pulumi.StringPtrInput  `pulumi:"recoveryReplicaDiskId"`
	RecoveryResourceGroupId                pulumi.StringPtrInput  `pulumi:"recoveryResourceGroupId"`
	RecoveryTargetDiskAccountType          pulumi.StringPtrInput  `pulumi:"recoveryTargetDiskAccountType"`
	RecoveryTargetDiskId                   pulumi.StringPtrInput  `pulumi:"recoveryTargetDiskId"`
	ResyncRequired                         pulumi.BoolPtrInput    `pulumi:"resyncRequired"`
	SecretIdentifier                       pulumi.StringPtrInput  `pulumi:"secretIdentifier"`
}

func (A2AProtectedManagedDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AProtectedManagedDiskDetailsResponse)(nil)).Elem()
}

func (i A2AProtectedManagedDiskDetailsResponseArgs) ToA2AProtectedManagedDiskDetailsResponseOutput() A2AProtectedManagedDiskDetailsResponseOutput {
	return i.ToA2AProtectedManagedDiskDetailsResponseOutputWithContext(context.Background())
}

func (i A2AProtectedManagedDiskDetailsResponseArgs) ToA2AProtectedManagedDiskDetailsResponseOutputWithContext(ctx context.Context) A2AProtectedManagedDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AProtectedManagedDiskDetailsResponseOutput)
}





type A2AProtectedManagedDiskDetailsResponseArrayInput interface {
	pulumi.Input

	ToA2AProtectedManagedDiskDetailsResponseArrayOutput() A2AProtectedManagedDiskDetailsResponseArrayOutput
	ToA2AProtectedManagedDiskDetailsResponseArrayOutputWithContext(context.Context) A2AProtectedManagedDiskDetailsResponseArrayOutput
}

type A2AProtectedManagedDiskDetailsResponseArray []A2AProtectedManagedDiskDetailsResponseInput

func (A2AProtectedManagedDiskDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]A2AProtectedManagedDiskDetailsResponse)(nil)).Elem()
}

func (i A2AProtectedManagedDiskDetailsResponseArray) ToA2AProtectedManagedDiskDetailsResponseArrayOutput() A2AProtectedManagedDiskDetailsResponseArrayOutput {
	return i.ToA2AProtectedManagedDiskDetailsResponseArrayOutputWithContext(context.Background())
}

func (i A2AProtectedManagedDiskDetailsResponseArray) ToA2AProtectedManagedDiskDetailsResponseArrayOutputWithContext(ctx context.Context) A2AProtectedManagedDiskDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AProtectedManagedDiskDetailsResponseArrayOutput)
}

type A2AProtectedManagedDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (A2AProtectedManagedDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AProtectedManagedDiskDetailsResponse)(nil)).Elem()
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) ToA2AProtectedManagedDiskDetailsResponseOutput() A2AProtectedManagedDiskDetailsResponseOutput {
	return o
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) ToA2AProtectedManagedDiskDetailsResponseOutputWithContext(ctx context.Context) A2AProtectedManagedDiskDetailsResponseOutput {
	return o
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) DataPendingAtSourceAgentInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *float64 { return v.DataPendingAtSourceAgentInMB }).(pulumi.Float64PtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) DataPendingInStagingStorageAccountInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *float64 {
		return v.DataPendingInStagingStorageAccountInMB
	}).(pulumi.Float64PtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) DekKeyVaultArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.DekKeyVaultArmId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) DiskCapacityInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *float64 { return v.DiskCapacityInBytes }).(pulumi.Float64PtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) DiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.DiskId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) DiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.DiskName }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) DiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.DiskType }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) IsDiskEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *bool { return v.IsDiskEncrypted }).(pulumi.BoolPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) IsDiskKeyEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *bool { return v.IsDiskKeyEncrypted }).(pulumi.BoolPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) KekKeyVaultArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.KekKeyVaultArmId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) MonitoringJobType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.MonitoringJobType }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) MonitoringPercentageCompletion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *int { return v.MonitoringPercentageCompletion }).(pulumi.IntPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) PrimaryStagingAzureStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.PrimaryStagingAzureStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) RecoveryReplicaDiskAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.RecoveryReplicaDiskAccountType }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) RecoveryReplicaDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.RecoveryReplicaDiskId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) RecoveryResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.RecoveryResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) RecoveryTargetDiskAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.RecoveryTargetDiskAccountType }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) RecoveryTargetDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.RecoveryTargetDiskId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) ResyncRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *bool { return v.ResyncRequired }).(pulumi.BoolPtrOutput)
}

func (o A2AProtectedManagedDiskDetailsResponseOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectedManagedDiskDetailsResponse) *string { return v.SecretIdentifier }).(pulumi.StringPtrOutput)
}

type A2AProtectedManagedDiskDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (A2AProtectedManagedDiskDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]A2AProtectedManagedDiskDetailsResponse)(nil)).Elem()
}

func (o A2AProtectedManagedDiskDetailsResponseArrayOutput) ToA2AProtectedManagedDiskDetailsResponseArrayOutput() A2AProtectedManagedDiskDetailsResponseArrayOutput {
	return o
}

func (o A2AProtectedManagedDiskDetailsResponseArrayOutput) ToA2AProtectedManagedDiskDetailsResponseArrayOutputWithContext(ctx context.Context) A2AProtectedManagedDiskDetailsResponseArrayOutput {
	return o
}

func (o A2AProtectedManagedDiskDetailsResponseArrayOutput) Index(i pulumi.IntInput) A2AProtectedManagedDiskDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) A2AProtectedManagedDiskDetailsResponse {
		return vs[0].([]A2AProtectedManagedDiskDetailsResponse)[vs[1].(int)]
	}).(A2AProtectedManagedDiskDetailsResponseOutput)
}

type A2AProtectionContainerMappingDetailsResponse struct {
	AgentAutoUpdateStatus  *string `pulumi:"agentAutoUpdateStatus"`
	AutomationAccountArmId *string `pulumi:"automationAccountArmId"`
	InstanceType           string  `pulumi:"instanceType"`
	JobScheduleName        *string `pulumi:"jobScheduleName"`
	ScheduleName           *string `pulumi:"scheduleName"`
}





type A2AProtectionContainerMappingDetailsResponseInput interface {
	pulumi.Input

	ToA2AProtectionContainerMappingDetailsResponseOutput() A2AProtectionContainerMappingDetailsResponseOutput
	ToA2AProtectionContainerMappingDetailsResponseOutputWithContext(context.Context) A2AProtectionContainerMappingDetailsResponseOutput
}

type A2AProtectionContainerMappingDetailsResponseArgs struct {
	AgentAutoUpdateStatus  pulumi.StringPtrInput `pulumi:"agentAutoUpdateStatus"`
	AutomationAccountArmId pulumi.StringPtrInput `pulumi:"automationAccountArmId"`
	InstanceType           pulumi.StringInput    `pulumi:"instanceType"`
	JobScheduleName        pulumi.StringPtrInput `pulumi:"jobScheduleName"`
	ScheduleName           pulumi.StringPtrInput `pulumi:"scheduleName"`
}

func (A2AProtectionContainerMappingDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AProtectionContainerMappingDetailsResponse)(nil)).Elem()
}

func (i A2AProtectionContainerMappingDetailsResponseArgs) ToA2AProtectionContainerMappingDetailsResponseOutput() A2AProtectionContainerMappingDetailsResponseOutput {
	return i.ToA2AProtectionContainerMappingDetailsResponseOutputWithContext(context.Background())
}

func (i A2AProtectionContainerMappingDetailsResponseArgs) ToA2AProtectionContainerMappingDetailsResponseOutputWithContext(ctx context.Context) A2AProtectionContainerMappingDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AProtectionContainerMappingDetailsResponseOutput)
}

type A2AProtectionContainerMappingDetailsResponseOutput struct{ *pulumi.OutputState }

func (A2AProtectionContainerMappingDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AProtectionContainerMappingDetailsResponse)(nil)).Elem()
}

func (o A2AProtectionContainerMappingDetailsResponseOutput) ToA2AProtectionContainerMappingDetailsResponseOutput() A2AProtectionContainerMappingDetailsResponseOutput {
	return o
}

func (o A2AProtectionContainerMappingDetailsResponseOutput) ToA2AProtectionContainerMappingDetailsResponseOutputWithContext(ctx context.Context) A2AProtectionContainerMappingDetailsResponseOutput {
	return o
}

func (o A2AProtectionContainerMappingDetailsResponseOutput) AgentAutoUpdateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectionContainerMappingDetailsResponse) *string { return v.AgentAutoUpdateStatus }).(pulumi.StringPtrOutput)
}

func (o A2AProtectionContainerMappingDetailsResponseOutput) AutomationAccountArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectionContainerMappingDetailsResponse) *string { return v.AutomationAccountArmId }).(pulumi.StringPtrOutput)
}

func (o A2AProtectionContainerMappingDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v A2AProtectionContainerMappingDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o A2AProtectionContainerMappingDetailsResponseOutput) JobScheduleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectionContainerMappingDetailsResponse) *string { return v.JobScheduleName }).(pulumi.StringPtrOutput)
}

func (o A2AProtectionContainerMappingDetailsResponseOutput) ScheduleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AProtectionContainerMappingDetailsResponse) *string { return v.ScheduleName }).(pulumi.StringPtrOutput)
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
	MultiVmGroupCreateOption           *string                                    `pulumi:"multiVmGroupCreateOption"`
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
	RecoveryBootDiagStorageAccountId   *string                                    `pulumi:"recoveryBootDiagStorageAccountId"`
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





type A2AReplicationDetailsResponseInput interface {
	pulumi.Input

	ToA2AReplicationDetailsResponseOutput() A2AReplicationDetailsResponseOutput
	ToA2AReplicationDetailsResponseOutputWithContext(context.Context) A2AReplicationDetailsResponseOutput
}

type A2AReplicationDetailsResponseArgs struct {
	AgentVersion                       pulumi.StringPtrInput                             `pulumi:"agentVersion"`
	FabricObjectId                     pulumi.StringPtrInput                             `pulumi:"fabricObjectId"`
	InstanceType                       pulumi.StringInput                                `pulumi:"instanceType"`
	IsReplicationAgentUpdateRequired   pulumi.BoolPtrInput                               `pulumi:"isReplicationAgentUpdateRequired"`
	LastHeartbeat                      pulumi.StringPtrInput                             `pulumi:"lastHeartbeat"`
	LastRpoCalculatedTime              pulumi.StringPtrInput                             `pulumi:"lastRpoCalculatedTime"`
	LifecycleId                        pulumi.StringPtrInput                             `pulumi:"lifecycleId"`
	ManagementId                       pulumi.StringPtrInput                             `pulumi:"managementId"`
	MonitoringJobType                  pulumi.StringPtrInput                             `pulumi:"monitoringJobType"`
	MonitoringPercentageCompletion     pulumi.IntPtrInput                                `pulumi:"monitoringPercentageCompletion"`
	MultiVmGroupCreateOption           pulumi.StringPtrInput                             `pulumi:"multiVmGroupCreateOption"`
	MultiVmGroupId                     pulumi.StringPtrInput                             `pulumi:"multiVmGroupId"`
	MultiVmGroupName                   pulumi.StringPtrInput                             `pulumi:"multiVmGroupName"`
	OsType                             pulumi.StringPtrInput                             `pulumi:"osType"`
	PrimaryFabricLocation              pulumi.StringPtrInput                             `pulumi:"primaryFabricLocation"`
	ProtectedDisks                     A2AProtectedDiskDetailsResponseArrayInput         `pulumi:"protectedDisks"`
	ProtectedManagedDisks              A2AProtectedManagedDiskDetailsResponseArrayInput  `pulumi:"protectedManagedDisks"`
	RecoveryAvailabilitySet            pulumi.StringPtrInput                             `pulumi:"recoveryAvailabilitySet"`
	RecoveryAzureResourceGroupId       pulumi.StringPtrInput                             `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureVMName                pulumi.StringPtrInput                             `pulumi:"recoveryAzureVMName"`
	RecoveryAzureVMSize                pulumi.StringPtrInput                             `pulumi:"recoveryAzureVMSize"`
	RecoveryBootDiagStorageAccountId   pulumi.StringPtrInput                             `pulumi:"recoveryBootDiagStorageAccountId"`
	RecoveryCloudService               pulumi.StringPtrInput                             `pulumi:"recoveryCloudService"`
	RecoveryFabricLocation             pulumi.StringPtrInput                             `pulumi:"recoveryFabricLocation"`
	RecoveryFabricObjectId             pulumi.StringPtrInput                             `pulumi:"recoveryFabricObjectId"`
	RpoInSeconds                       pulumi.Float64PtrInput                            `pulumi:"rpoInSeconds"`
	SelectedRecoveryAzureNetworkId     pulumi.StringPtrInput                             `pulumi:"selectedRecoveryAzureNetworkId"`
	TestFailoverRecoveryFabricObjectId pulumi.StringPtrInput                             `pulumi:"testFailoverRecoveryFabricObjectId"`
	VmNics                             VMNicDetailsResponseArrayInput                    `pulumi:"vmNics"`
	VmProtectionState                  pulumi.StringPtrInput                             `pulumi:"vmProtectionState"`
	VmProtectionStateDescription       pulumi.StringPtrInput                             `pulumi:"vmProtectionStateDescription"`
	VmSyncedConfigDetails              AzureToAzureVmSyncedConfigDetailsResponsePtrInput `pulumi:"vmSyncedConfigDetails"`
}

func (A2AReplicationDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AReplicationDetailsResponse)(nil)).Elem()
}

func (i A2AReplicationDetailsResponseArgs) ToA2AReplicationDetailsResponseOutput() A2AReplicationDetailsResponseOutput {
	return i.ToA2AReplicationDetailsResponseOutputWithContext(context.Background())
}

func (i A2AReplicationDetailsResponseArgs) ToA2AReplicationDetailsResponseOutputWithContext(ctx context.Context) A2AReplicationDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AReplicationDetailsResponseOutput)
}

type A2AReplicationDetailsResponseOutput struct{ *pulumi.OutputState }

func (A2AReplicationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AReplicationDetailsResponse)(nil)).Elem()
}

func (o A2AReplicationDetailsResponseOutput) ToA2AReplicationDetailsResponseOutput() A2AReplicationDetailsResponseOutput {
	return o
}

func (o A2AReplicationDetailsResponseOutput) ToA2AReplicationDetailsResponseOutputWithContext(ctx context.Context) A2AReplicationDetailsResponseOutput {
	return o
}

func (o A2AReplicationDetailsResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) FabricObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.FabricObjectId }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o A2AReplicationDetailsResponseOutput) IsReplicationAgentUpdateRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *bool { return v.IsReplicationAgentUpdateRequired }).(pulumi.BoolPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) LastHeartbeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.LastHeartbeat }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) LastRpoCalculatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.LastRpoCalculatedTime }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) LifecycleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.LifecycleId }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) ManagementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.ManagementId }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) MonitoringJobType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.MonitoringJobType }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) MonitoringPercentageCompletion() pulumi.IntPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *int { return v.MonitoringPercentageCompletion }).(pulumi.IntPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) MultiVmGroupCreateOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.MultiVmGroupCreateOption }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) MultiVmGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.MultiVmGroupId }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) MultiVmGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.MultiVmGroupName }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) PrimaryFabricLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.PrimaryFabricLocation }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) ProtectedDisks() A2AProtectedDiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) []A2AProtectedDiskDetailsResponse { return v.ProtectedDisks }).(A2AProtectedDiskDetailsResponseArrayOutput)
}

func (o A2AReplicationDetailsResponseOutput) ProtectedManagedDisks() A2AProtectedManagedDiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) []A2AProtectedManagedDiskDetailsResponse {
		return v.ProtectedManagedDisks
	}).(A2AProtectedManagedDiskDetailsResponseArrayOutput)
}

func (o A2AReplicationDetailsResponseOutput) RecoveryAvailabilitySet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.RecoveryAvailabilitySet }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) RecoveryAzureResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.RecoveryAzureResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) RecoveryAzureVMName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.RecoveryAzureVMName }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) RecoveryAzureVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.RecoveryAzureVMSize }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) RecoveryBootDiagStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.RecoveryBootDiagStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) RecoveryCloudService() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.RecoveryCloudService }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) RecoveryFabricLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.RecoveryFabricLocation }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) RecoveryFabricObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.RecoveryFabricObjectId }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) RpoInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *float64 { return v.RpoInSeconds }).(pulumi.Float64PtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) SelectedRecoveryAzureNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.SelectedRecoveryAzureNetworkId }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) TestFailoverRecoveryFabricObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.TestFailoverRecoveryFabricObjectId }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) VmNics() VMNicDetailsResponseArrayOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) []VMNicDetailsResponse { return v.VmNics }).(VMNicDetailsResponseArrayOutput)
}

func (o A2AReplicationDetailsResponseOutput) VmProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.VmProtectionState }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) VmProtectionStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *string { return v.VmProtectionStateDescription }).(pulumi.StringPtrOutput)
}

func (o A2AReplicationDetailsResponseOutput) VmSyncedConfigDetails() AzureToAzureVmSyncedConfigDetailsResponsePtrOutput {
	return o.ApplyT(func(v A2AReplicationDetailsResponse) *AzureToAzureVmSyncedConfigDetailsResponse {
		return v.VmSyncedConfigDetails
	}).(AzureToAzureVmSyncedConfigDetailsResponsePtrOutput)
}

type A2AVmDiskInputDetails struct {
	DiskUri                             *string `pulumi:"diskUri"`
	PrimaryStagingAzureStorageAccountId *string `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryAzureStorageAccountId       *string `pulumi:"recoveryAzureStorageAccountId"`
}





type A2AVmDiskInputDetailsInput interface {
	pulumi.Input

	ToA2AVmDiskInputDetailsOutput() A2AVmDiskInputDetailsOutput
	ToA2AVmDiskInputDetailsOutputWithContext(context.Context) A2AVmDiskInputDetailsOutput
}

type A2AVmDiskInputDetailsArgs struct {
	DiskUri                             pulumi.StringPtrInput `pulumi:"diskUri"`
	PrimaryStagingAzureStorageAccountId pulumi.StringPtrInput `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryAzureStorageAccountId       pulumi.StringPtrInput `pulumi:"recoveryAzureStorageAccountId"`
}

func (A2AVmDiskInputDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AVmDiskInputDetails)(nil)).Elem()
}

func (i A2AVmDiskInputDetailsArgs) ToA2AVmDiskInputDetailsOutput() A2AVmDiskInputDetailsOutput {
	return i.ToA2AVmDiskInputDetailsOutputWithContext(context.Background())
}

func (i A2AVmDiskInputDetailsArgs) ToA2AVmDiskInputDetailsOutputWithContext(ctx context.Context) A2AVmDiskInputDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AVmDiskInputDetailsOutput)
}





type A2AVmDiskInputDetailsArrayInput interface {
	pulumi.Input

	ToA2AVmDiskInputDetailsArrayOutput() A2AVmDiskInputDetailsArrayOutput
	ToA2AVmDiskInputDetailsArrayOutputWithContext(context.Context) A2AVmDiskInputDetailsArrayOutput
}

type A2AVmDiskInputDetailsArray []A2AVmDiskInputDetailsInput

func (A2AVmDiskInputDetailsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]A2AVmDiskInputDetails)(nil)).Elem()
}

func (i A2AVmDiskInputDetailsArray) ToA2AVmDiskInputDetailsArrayOutput() A2AVmDiskInputDetailsArrayOutput {
	return i.ToA2AVmDiskInputDetailsArrayOutputWithContext(context.Background())
}

func (i A2AVmDiskInputDetailsArray) ToA2AVmDiskInputDetailsArrayOutputWithContext(ctx context.Context) A2AVmDiskInputDetailsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AVmDiskInputDetailsArrayOutput)
}

type A2AVmDiskInputDetailsOutput struct{ *pulumi.OutputState }

func (A2AVmDiskInputDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AVmDiskInputDetails)(nil)).Elem()
}

func (o A2AVmDiskInputDetailsOutput) ToA2AVmDiskInputDetailsOutput() A2AVmDiskInputDetailsOutput {
	return o
}

func (o A2AVmDiskInputDetailsOutput) ToA2AVmDiskInputDetailsOutputWithContext(ctx context.Context) A2AVmDiskInputDetailsOutput {
	return o
}

func (o A2AVmDiskInputDetailsOutput) DiskUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AVmDiskInputDetails) *string { return v.DiskUri }).(pulumi.StringPtrOutput)
}

func (o A2AVmDiskInputDetailsOutput) PrimaryStagingAzureStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AVmDiskInputDetails) *string { return v.PrimaryStagingAzureStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o A2AVmDiskInputDetailsOutput) RecoveryAzureStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AVmDiskInputDetails) *string { return v.RecoveryAzureStorageAccountId }).(pulumi.StringPtrOutput)
}

type A2AVmDiskInputDetailsArrayOutput struct{ *pulumi.OutputState }

func (A2AVmDiskInputDetailsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]A2AVmDiskInputDetails)(nil)).Elem()
}

func (o A2AVmDiskInputDetailsArrayOutput) ToA2AVmDiskInputDetailsArrayOutput() A2AVmDiskInputDetailsArrayOutput {
	return o
}

func (o A2AVmDiskInputDetailsArrayOutput) ToA2AVmDiskInputDetailsArrayOutputWithContext(ctx context.Context) A2AVmDiskInputDetailsArrayOutput {
	return o
}

func (o A2AVmDiskInputDetailsArrayOutput) Index(i pulumi.IntInput) A2AVmDiskInputDetailsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) A2AVmDiskInputDetails {
		return vs[0].([]A2AVmDiskInputDetails)[vs[1].(int)]
	}).(A2AVmDiskInputDetailsOutput)
}

type A2AVmManagedDiskInputDetails struct {
	DiskId                              *string `pulumi:"diskId"`
	PrimaryStagingAzureStorageAccountId *string `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryReplicaDiskAccountType      *string `pulumi:"recoveryReplicaDiskAccountType"`
	RecoveryResourceGroupId             *string `pulumi:"recoveryResourceGroupId"`
	RecoveryTargetDiskAccountType       *string `pulumi:"recoveryTargetDiskAccountType"`
}





type A2AVmManagedDiskInputDetailsInput interface {
	pulumi.Input

	ToA2AVmManagedDiskInputDetailsOutput() A2AVmManagedDiskInputDetailsOutput
	ToA2AVmManagedDiskInputDetailsOutputWithContext(context.Context) A2AVmManagedDiskInputDetailsOutput
}

type A2AVmManagedDiskInputDetailsArgs struct {
	DiskId                              pulumi.StringPtrInput `pulumi:"diskId"`
	PrimaryStagingAzureStorageAccountId pulumi.StringPtrInput `pulumi:"primaryStagingAzureStorageAccountId"`
	RecoveryReplicaDiskAccountType      pulumi.StringPtrInput `pulumi:"recoveryReplicaDiskAccountType"`
	RecoveryResourceGroupId             pulumi.StringPtrInput `pulumi:"recoveryResourceGroupId"`
	RecoveryTargetDiskAccountType       pulumi.StringPtrInput `pulumi:"recoveryTargetDiskAccountType"`
}

func (A2AVmManagedDiskInputDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AVmManagedDiskInputDetails)(nil)).Elem()
}

func (i A2AVmManagedDiskInputDetailsArgs) ToA2AVmManagedDiskInputDetailsOutput() A2AVmManagedDiskInputDetailsOutput {
	return i.ToA2AVmManagedDiskInputDetailsOutputWithContext(context.Background())
}

func (i A2AVmManagedDiskInputDetailsArgs) ToA2AVmManagedDiskInputDetailsOutputWithContext(ctx context.Context) A2AVmManagedDiskInputDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AVmManagedDiskInputDetailsOutput)
}





type A2AVmManagedDiskInputDetailsArrayInput interface {
	pulumi.Input

	ToA2AVmManagedDiskInputDetailsArrayOutput() A2AVmManagedDiskInputDetailsArrayOutput
	ToA2AVmManagedDiskInputDetailsArrayOutputWithContext(context.Context) A2AVmManagedDiskInputDetailsArrayOutput
}

type A2AVmManagedDiskInputDetailsArray []A2AVmManagedDiskInputDetailsInput

func (A2AVmManagedDiskInputDetailsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]A2AVmManagedDiskInputDetails)(nil)).Elem()
}

func (i A2AVmManagedDiskInputDetailsArray) ToA2AVmManagedDiskInputDetailsArrayOutput() A2AVmManagedDiskInputDetailsArrayOutput {
	return i.ToA2AVmManagedDiskInputDetailsArrayOutputWithContext(context.Background())
}

func (i A2AVmManagedDiskInputDetailsArray) ToA2AVmManagedDiskInputDetailsArrayOutputWithContext(ctx context.Context) A2AVmManagedDiskInputDetailsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(A2AVmManagedDiskInputDetailsArrayOutput)
}

type A2AVmManagedDiskInputDetailsOutput struct{ *pulumi.OutputState }

func (A2AVmManagedDiskInputDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*A2AVmManagedDiskInputDetails)(nil)).Elem()
}

func (o A2AVmManagedDiskInputDetailsOutput) ToA2AVmManagedDiskInputDetailsOutput() A2AVmManagedDiskInputDetailsOutput {
	return o
}

func (o A2AVmManagedDiskInputDetailsOutput) ToA2AVmManagedDiskInputDetailsOutputWithContext(ctx context.Context) A2AVmManagedDiskInputDetailsOutput {
	return o
}

func (o A2AVmManagedDiskInputDetailsOutput) DiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AVmManagedDiskInputDetails) *string { return v.DiskId }).(pulumi.StringPtrOutput)
}

func (o A2AVmManagedDiskInputDetailsOutput) PrimaryStagingAzureStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AVmManagedDiskInputDetails) *string { return v.PrimaryStagingAzureStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o A2AVmManagedDiskInputDetailsOutput) RecoveryReplicaDiskAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AVmManagedDiskInputDetails) *string { return v.RecoveryReplicaDiskAccountType }).(pulumi.StringPtrOutput)
}

func (o A2AVmManagedDiskInputDetailsOutput) RecoveryResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AVmManagedDiskInputDetails) *string { return v.RecoveryResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o A2AVmManagedDiskInputDetailsOutput) RecoveryTargetDiskAccountType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v A2AVmManagedDiskInputDetails) *string { return v.RecoveryTargetDiskAccountType }).(pulumi.StringPtrOutput)
}

type A2AVmManagedDiskInputDetailsArrayOutput struct{ *pulumi.OutputState }

func (A2AVmManagedDiskInputDetailsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]A2AVmManagedDiskInputDetails)(nil)).Elem()
}

func (o A2AVmManagedDiskInputDetailsArrayOutput) ToA2AVmManagedDiskInputDetailsArrayOutput() A2AVmManagedDiskInputDetailsArrayOutput {
	return o
}

func (o A2AVmManagedDiskInputDetailsArrayOutput) ToA2AVmManagedDiskInputDetailsArrayOutputWithContext(ctx context.Context) A2AVmManagedDiskInputDetailsArrayOutput {
	return o
}

func (o A2AVmManagedDiskInputDetailsArrayOutput) Index(i pulumi.IntInput) A2AVmManagedDiskInputDetailsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) A2AVmManagedDiskInputDetails {
		return vs[0].([]A2AVmManagedDiskInputDetails)[vs[1].(int)]
	}).(A2AVmManagedDiskInputDetailsOutput)
}

type AddRecoveryServicesProviderInputProperties struct {
	AuthenticationIdentityInput IdentityProviderInput `pulumi:"authenticationIdentityInput"`
	MachineName                 string                `pulumi:"machineName"`
	ResourceAccessIdentityInput IdentityProviderInput `pulumi:"resourceAccessIdentityInput"`
}





type AddRecoveryServicesProviderInputPropertiesInput interface {
	pulumi.Input

	ToAddRecoveryServicesProviderInputPropertiesOutput() AddRecoveryServicesProviderInputPropertiesOutput
	ToAddRecoveryServicesProviderInputPropertiesOutputWithContext(context.Context) AddRecoveryServicesProviderInputPropertiesOutput
}

type AddRecoveryServicesProviderInputPropertiesArgs struct {
	AuthenticationIdentityInput IdentityProviderInputInput `pulumi:"authenticationIdentityInput"`
	MachineName                 pulumi.StringInput         `pulumi:"machineName"`
	ResourceAccessIdentityInput IdentityProviderInputInput `pulumi:"resourceAccessIdentityInput"`
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

func (i AddRecoveryServicesProviderInputPropertiesArgs) ToAddRecoveryServicesProviderInputPropertiesPtrOutput() AddRecoveryServicesProviderInputPropertiesPtrOutput {
	return i.ToAddRecoveryServicesProviderInputPropertiesPtrOutputWithContext(context.Background())
}

func (i AddRecoveryServicesProviderInputPropertiesArgs) ToAddRecoveryServicesProviderInputPropertiesPtrOutputWithContext(ctx context.Context) AddRecoveryServicesProviderInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddRecoveryServicesProviderInputPropertiesOutput).ToAddRecoveryServicesProviderInputPropertiesPtrOutputWithContext(ctx)
}









type AddRecoveryServicesProviderInputPropertiesPtrInput interface {
	pulumi.Input

	ToAddRecoveryServicesProviderInputPropertiesPtrOutput() AddRecoveryServicesProviderInputPropertiesPtrOutput
	ToAddRecoveryServicesProviderInputPropertiesPtrOutputWithContext(context.Context) AddRecoveryServicesProviderInputPropertiesPtrOutput
}

type addRecoveryServicesProviderInputPropertiesPtrType AddRecoveryServicesProviderInputPropertiesArgs

func AddRecoveryServicesProviderInputPropertiesPtr(v *AddRecoveryServicesProviderInputPropertiesArgs) AddRecoveryServicesProviderInputPropertiesPtrInput {
	return (*addRecoveryServicesProviderInputPropertiesPtrType)(v)
}

func (*addRecoveryServicesProviderInputPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AddRecoveryServicesProviderInputProperties)(nil)).Elem()
}

func (i *addRecoveryServicesProviderInputPropertiesPtrType) ToAddRecoveryServicesProviderInputPropertiesPtrOutput() AddRecoveryServicesProviderInputPropertiesPtrOutput {
	return i.ToAddRecoveryServicesProviderInputPropertiesPtrOutputWithContext(context.Background())
}

func (i *addRecoveryServicesProviderInputPropertiesPtrType) ToAddRecoveryServicesProviderInputPropertiesPtrOutputWithContext(ctx context.Context) AddRecoveryServicesProviderInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddRecoveryServicesProviderInputPropertiesPtrOutput)
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

func (o AddRecoveryServicesProviderInputPropertiesOutput) ToAddRecoveryServicesProviderInputPropertiesPtrOutput() AddRecoveryServicesProviderInputPropertiesPtrOutput {
	return o.ToAddRecoveryServicesProviderInputPropertiesPtrOutputWithContext(context.Background())
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) ToAddRecoveryServicesProviderInputPropertiesPtrOutputWithContext(ctx context.Context) AddRecoveryServicesProviderInputPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AddRecoveryServicesProviderInputProperties) *AddRecoveryServicesProviderInputProperties {
		return &v
	}).(AddRecoveryServicesProviderInputPropertiesPtrOutput)
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) AuthenticationIdentityInput() IdentityProviderInputOutput {
	return o.ApplyT(func(v AddRecoveryServicesProviderInputProperties) IdentityProviderInput {
		return v.AuthenticationIdentityInput
	}).(IdentityProviderInputOutput)
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) MachineName() pulumi.StringOutput {
	return o.ApplyT(func(v AddRecoveryServicesProviderInputProperties) string { return v.MachineName }).(pulumi.StringOutput)
}

func (o AddRecoveryServicesProviderInputPropertiesOutput) ResourceAccessIdentityInput() IdentityProviderInputOutput {
	return o.ApplyT(func(v AddRecoveryServicesProviderInputProperties) IdentityProviderInput {
		return v.ResourceAccessIdentityInput
	}).(IdentityProviderInputOutput)
}

type AddRecoveryServicesProviderInputPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AddRecoveryServicesProviderInputPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AddRecoveryServicesProviderInputProperties)(nil)).Elem()
}

func (o AddRecoveryServicesProviderInputPropertiesPtrOutput) ToAddRecoveryServicesProviderInputPropertiesPtrOutput() AddRecoveryServicesProviderInputPropertiesPtrOutput {
	return o
}

func (o AddRecoveryServicesProviderInputPropertiesPtrOutput) ToAddRecoveryServicesProviderInputPropertiesPtrOutputWithContext(ctx context.Context) AddRecoveryServicesProviderInputPropertiesPtrOutput {
	return o
}

func (o AddRecoveryServicesProviderInputPropertiesPtrOutput) Elem() AddRecoveryServicesProviderInputPropertiesOutput {
	return o.ApplyT(func(v *AddRecoveryServicesProviderInputProperties) AddRecoveryServicesProviderInputProperties {
		if v != nil {
			return *v
		}
		var ret AddRecoveryServicesProviderInputProperties
		return ret
	}).(AddRecoveryServicesProviderInputPropertiesOutput)
}

func (o AddRecoveryServicesProviderInputPropertiesPtrOutput) AuthenticationIdentityInput() IdentityProviderInputPtrOutput {
	return o.ApplyT(func(v *AddRecoveryServicesProviderInputProperties) *IdentityProviderInput {
		if v == nil {
			return nil
		}
		return &v.AuthenticationIdentityInput
	}).(IdentityProviderInputPtrOutput)
}

func (o AddRecoveryServicesProviderInputPropertiesPtrOutput) MachineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AddRecoveryServicesProviderInputProperties) *string {
		if v == nil {
			return nil
		}
		return &v.MachineName
	}).(pulumi.StringPtrOutput)
}

func (o AddRecoveryServicesProviderInputPropertiesPtrOutput) ResourceAccessIdentityInput() IdentityProviderInputPtrOutput {
	return o.ApplyT(func(v *AddRecoveryServicesProviderInputProperties) *IdentityProviderInput {
		if v == nil {
			return nil
		}
		return &v.ResourceAccessIdentityInput
	}).(IdentityProviderInputPtrOutput)
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

type AzureFabricCreationInput struct {
	InstanceType *string `pulumi:"instanceType"`
	Location     *string `pulumi:"location"`
}





type AzureFabricCreationInputInput interface {
	pulumi.Input

	ToAzureFabricCreationInputOutput() AzureFabricCreationInputOutput
	ToAzureFabricCreationInputOutputWithContext(context.Context) AzureFabricCreationInputOutput
}

type AzureFabricCreationInputArgs struct {
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	Location     pulumi.StringPtrInput `pulumi:"location"`
}

func (AzureFabricCreationInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFabricCreationInput)(nil)).Elem()
}

func (i AzureFabricCreationInputArgs) ToAzureFabricCreationInputOutput() AzureFabricCreationInputOutput {
	return i.ToAzureFabricCreationInputOutputWithContext(context.Background())
}

func (i AzureFabricCreationInputArgs) ToAzureFabricCreationInputOutputWithContext(ctx context.Context) AzureFabricCreationInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFabricCreationInputOutput)
}

type AzureFabricCreationInputOutput struct{ *pulumi.OutputState }

func (AzureFabricCreationInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFabricCreationInput)(nil)).Elem()
}

func (o AzureFabricCreationInputOutput) ToAzureFabricCreationInputOutput() AzureFabricCreationInputOutput {
	return o
}

func (o AzureFabricCreationInputOutput) ToAzureFabricCreationInputOutputWithContext(ctx context.Context) AzureFabricCreationInputOutput {
	return o
}

func (o AzureFabricCreationInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFabricCreationInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o AzureFabricCreationInputOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFabricCreationInput) *string { return v.Location }).(pulumi.StringPtrOutput)
}

type AzureFabricSpecificDetailsResponse struct {
	ContainerIds []string `pulumi:"containerIds"`
	InstanceType string   `pulumi:"instanceType"`
	Location     *string  `pulumi:"location"`
}





type AzureFabricSpecificDetailsResponseInput interface {
	pulumi.Input

	ToAzureFabricSpecificDetailsResponseOutput() AzureFabricSpecificDetailsResponseOutput
	ToAzureFabricSpecificDetailsResponseOutputWithContext(context.Context) AzureFabricSpecificDetailsResponseOutput
}

type AzureFabricSpecificDetailsResponseArgs struct {
	ContainerIds pulumi.StringArrayInput `pulumi:"containerIds"`
	InstanceType pulumi.StringInput      `pulumi:"instanceType"`
	Location     pulumi.StringPtrInput   `pulumi:"location"`
}

func (AzureFabricSpecificDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFabricSpecificDetailsResponse)(nil)).Elem()
}

func (i AzureFabricSpecificDetailsResponseArgs) ToAzureFabricSpecificDetailsResponseOutput() AzureFabricSpecificDetailsResponseOutput {
	return i.ToAzureFabricSpecificDetailsResponseOutputWithContext(context.Background())
}

func (i AzureFabricSpecificDetailsResponseArgs) ToAzureFabricSpecificDetailsResponseOutputWithContext(ctx context.Context) AzureFabricSpecificDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureFabricSpecificDetailsResponseOutput)
}

type AzureFabricSpecificDetailsResponseOutput struct{ *pulumi.OutputState }

func (AzureFabricSpecificDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureFabricSpecificDetailsResponse)(nil)).Elem()
}

func (o AzureFabricSpecificDetailsResponseOutput) ToAzureFabricSpecificDetailsResponseOutput() AzureFabricSpecificDetailsResponseOutput {
	return o
}

func (o AzureFabricSpecificDetailsResponseOutput) ToAzureFabricSpecificDetailsResponseOutputWithContext(ctx context.Context) AzureFabricSpecificDetailsResponseOutput {
	return o
}

func (o AzureFabricSpecificDetailsResponseOutput) ContainerIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureFabricSpecificDetailsResponse) []string { return v.ContainerIds }).(pulumi.StringArrayOutput)
}

func (o AzureFabricSpecificDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureFabricSpecificDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o AzureFabricSpecificDetailsResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureFabricSpecificDetailsResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

type AzureToAzureCreateNetworkMappingInput struct {
	InstanceType     *string `pulumi:"instanceType"`
	PrimaryNetworkId *string `pulumi:"primaryNetworkId"`
}





type AzureToAzureCreateNetworkMappingInputInput interface {
	pulumi.Input

	ToAzureToAzureCreateNetworkMappingInputOutput() AzureToAzureCreateNetworkMappingInputOutput
	ToAzureToAzureCreateNetworkMappingInputOutputWithContext(context.Context) AzureToAzureCreateNetworkMappingInputOutput
}

type AzureToAzureCreateNetworkMappingInputArgs struct {
	InstanceType     pulumi.StringPtrInput `pulumi:"instanceType"`
	PrimaryNetworkId pulumi.StringPtrInput `pulumi:"primaryNetworkId"`
}

func (AzureToAzureCreateNetworkMappingInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureToAzureCreateNetworkMappingInput)(nil)).Elem()
}

func (i AzureToAzureCreateNetworkMappingInputArgs) ToAzureToAzureCreateNetworkMappingInputOutput() AzureToAzureCreateNetworkMappingInputOutput {
	return i.ToAzureToAzureCreateNetworkMappingInputOutputWithContext(context.Background())
}

func (i AzureToAzureCreateNetworkMappingInputArgs) ToAzureToAzureCreateNetworkMappingInputOutputWithContext(ctx context.Context) AzureToAzureCreateNetworkMappingInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureToAzureCreateNetworkMappingInputOutput)
}

type AzureToAzureCreateNetworkMappingInputOutput struct{ *pulumi.OutputState }

func (AzureToAzureCreateNetworkMappingInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureToAzureCreateNetworkMappingInput)(nil)).Elem()
}

func (o AzureToAzureCreateNetworkMappingInputOutput) ToAzureToAzureCreateNetworkMappingInputOutput() AzureToAzureCreateNetworkMappingInputOutput {
	return o
}

func (o AzureToAzureCreateNetworkMappingInputOutput) ToAzureToAzureCreateNetworkMappingInputOutputWithContext(ctx context.Context) AzureToAzureCreateNetworkMappingInputOutput {
	return o
}

func (o AzureToAzureCreateNetworkMappingInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureToAzureCreateNetworkMappingInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o AzureToAzureCreateNetworkMappingInputOutput) PrimaryNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureToAzureCreateNetworkMappingInput) *string { return v.PrimaryNetworkId }).(pulumi.StringPtrOutput)
}

type AzureToAzureNetworkMappingSettingsResponse struct {
	InstanceType           string  `pulumi:"instanceType"`
	PrimaryFabricLocation  *string `pulumi:"primaryFabricLocation"`
	RecoveryFabricLocation *string `pulumi:"recoveryFabricLocation"`
}





type AzureToAzureNetworkMappingSettingsResponseInput interface {
	pulumi.Input

	ToAzureToAzureNetworkMappingSettingsResponseOutput() AzureToAzureNetworkMappingSettingsResponseOutput
	ToAzureToAzureNetworkMappingSettingsResponseOutputWithContext(context.Context) AzureToAzureNetworkMappingSettingsResponseOutput
}

type AzureToAzureNetworkMappingSettingsResponseArgs struct {
	InstanceType           pulumi.StringInput    `pulumi:"instanceType"`
	PrimaryFabricLocation  pulumi.StringPtrInput `pulumi:"primaryFabricLocation"`
	RecoveryFabricLocation pulumi.StringPtrInput `pulumi:"recoveryFabricLocation"`
}

func (AzureToAzureNetworkMappingSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureToAzureNetworkMappingSettingsResponse)(nil)).Elem()
}

func (i AzureToAzureNetworkMappingSettingsResponseArgs) ToAzureToAzureNetworkMappingSettingsResponseOutput() AzureToAzureNetworkMappingSettingsResponseOutput {
	return i.ToAzureToAzureNetworkMappingSettingsResponseOutputWithContext(context.Background())
}

func (i AzureToAzureNetworkMappingSettingsResponseArgs) ToAzureToAzureNetworkMappingSettingsResponseOutputWithContext(ctx context.Context) AzureToAzureNetworkMappingSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureToAzureNetworkMappingSettingsResponseOutput)
}

type AzureToAzureNetworkMappingSettingsResponseOutput struct{ *pulumi.OutputState }

func (AzureToAzureNetworkMappingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureToAzureNetworkMappingSettingsResponse)(nil)).Elem()
}

func (o AzureToAzureNetworkMappingSettingsResponseOutput) ToAzureToAzureNetworkMappingSettingsResponseOutput() AzureToAzureNetworkMappingSettingsResponseOutput {
	return o
}

func (o AzureToAzureNetworkMappingSettingsResponseOutput) ToAzureToAzureNetworkMappingSettingsResponseOutputWithContext(ctx context.Context) AzureToAzureNetworkMappingSettingsResponseOutput {
	return o
}

func (o AzureToAzureNetworkMappingSettingsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v AzureToAzureNetworkMappingSettingsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o AzureToAzureNetworkMappingSettingsResponseOutput) PrimaryFabricLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureToAzureNetworkMappingSettingsResponse) *string { return v.PrimaryFabricLocation }).(pulumi.StringPtrOutput)
}

func (o AzureToAzureNetworkMappingSettingsResponseOutput) RecoveryFabricLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureToAzureNetworkMappingSettingsResponse) *string { return v.RecoveryFabricLocation }).(pulumi.StringPtrOutput)
}

type AzureToAzureVmSyncedConfigDetailsResponse struct {
	InputEndpoints  []InputEndpointResponse  `pulumi:"inputEndpoints"`
	RoleAssignments []RoleAssignmentResponse `pulumi:"roleAssignments"`
	Tags            map[string]string        `pulumi:"tags"`
}





type AzureToAzureVmSyncedConfigDetailsResponseInput interface {
	pulumi.Input

	ToAzureToAzureVmSyncedConfigDetailsResponseOutput() AzureToAzureVmSyncedConfigDetailsResponseOutput
	ToAzureToAzureVmSyncedConfigDetailsResponseOutputWithContext(context.Context) AzureToAzureVmSyncedConfigDetailsResponseOutput
}

type AzureToAzureVmSyncedConfigDetailsResponseArgs struct {
	InputEndpoints  InputEndpointResponseArrayInput  `pulumi:"inputEndpoints"`
	RoleAssignments RoleAssignmentResponseArrayInput `pulumi:"roleAssignments"`
	Tags            pulumi.StringMapInput            `pulumi:"tags"`
}

func (AzureToAzureVmSyncedConfigDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureToAzureVmSyncedConfigDetailsResponse)(nil)).Elem()
}

func (i AzureToAzureVmSyncedConfigDetailsResponseArgs) ToAzureToAzureVmSyncedConfigDetailsResponseOutput() AzureToAzureVmSyncedConfigDetailsResponseOutput {
	return i.ToAzureToAzureVmSyncedConfigDetailsResponseOutputWithContext(context.Background())
}

func (i AzureToAzureVmSyncedConfigDetailsResponseArgs) ToAzureToAzureVmSyncedConfigDetailsResponseOutputWithContext(ctx context.Context) AzureToAzureVmSyncedConfigDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureToAzureVmSyncedConfigDetailsResponseOutput)
}

func (i AzureToAzureVmSyncedConfigDetailsResponseArgs) ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutput() AzureToAzureVmSyncedConfigDetailsResponsePtrOutput {
	return i.ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutputWithContext(context.Background())
}

func (i AzureToAzureVmSyncedConfigDetailsResponseArgs) ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutputWithContext(ctx context.Context) AzureToAzureVmSyncedConfigDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureToAzureVmSyncedConfigDetailsResponseOutput).ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutputWithContext(ctx)
}









type AzureToAzureVmSyncedConfigDetailsResponsePtrInput interface {
	pulumi.Input

	ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutput() AzureToAzureVmSyncedConfigDetailsResponsePtrOutput
	ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutputWithContext(context.Context) AzureToAzureVmSyncedConfigDetailsResponsePtrOutput
}

type azureToAzureVmSyncedConfigDetailsResponsePtrType AzureToAzureVmSyncedConfigDetailsResponseArgs

func AzureToAzureVmSyncedConfigDetailsResponsePtr(v *AzureToAzureVmSyncedConfigDetailsResponseArgs) AzureToAzureVmSyncedConfigDetailsResponsePtrInput {
	return (*azureToAzureVmSyncedConfigDetailsResponsePtrType)(v)
}

func (*azureToAzureVmSyncedConfigDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureToAzureVmSyncedConfigDetailsResponse)(nil)).Elem()
}

func (i *azureToAzureVmSyncedConfigDetailsResponsePtrType) ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutput() AzureToAzureVmSyncedConfigDetailsResponsePtrOutput {
	return i.ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *azureToAzureVmSyncedConfigDetailsResponsePtrType) ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutputWithContext(ctx context.Context) AzureToAzureVmSyncedConfigDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureToAzureVmSyncedConfigDetailsResponsePtrOutput)
}

type AzureToAzureVmSyncedConfigDetailsResponseOutput struct{ *pulumi.OutputState }

func (AzureToAzureVmSyncedConfigDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureToAzureVmSyncedConfigDetailsResponse)(nil)).Elem()
}

func (o AzureToAzureVmSyncedConfigDetailsResponseOutput) ToAzureToAzureVmSyncedConfigDetailsResponseOutput() AzureToAzureVmSyncedConfigDetailsResponseOutput {
	return o
}

func (o AzureToAzureVmSyncedConfigDetailsResponseOutput) ToAzureToAzureVmSyncedConfigDetailsResponseOutputWithContext(ctx context.Context) AzureToAzureVmSyncedConfigDetailsResponseOutput {
	return o
}

func (o AzureToAzureVmSyncedConfigDetailsResponseOutput) ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutput() AzureToAzureVmSyncedConfigDetailsResponsePtrOutput {
	return o.ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutputWithContext(context.Background())
}

func (o AzureToAzureVmSyncedConfigDetailsResponseOutput) ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutputWithContext(ctx context.Context) AzureToAzureVmSyncedConfigDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureToAzureVmSyncedConfigDetailsResponse) *AzureToAzureVmSyncedConfigDetailsResponse {
		return &v
	}).(AzureToAzureVmSyncedConfigDetailsResponsePtrOutput)
}

func (o AzureToAzureVmSyncedConfigDetailsResponseOutput) InputEndpoints() InputEndpointResponseArrayOutput {
	return o.ApplyT(func(v AzureToAzureVmSyncedConfigDetailsResponse) []InputEndpointResponse { return v.InputEndpoints }).(InputEndpointResponseArrayOutput)
}

func (o AzureToAzureVmSyncedConfigDetailsResponseOutput) RoleAssignments() RoleAssignmentResponseArrayOutput {
	return o.ApplyT(func(v AzureToAzureVmSyncedConfigDetailsResponse) []RoleAssignmentResponse { return v.RoleAssignments }).(RoleAssignmentResponseArrayOutput)
}

func (o AzureToAzureVmSyncedConfigDetailsResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v AzureToAzureVmSyncedConfigDetailsResponse) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

type AzureToAzureVmSyncedConfigDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureToAzureVmSyncedConfigDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureToAzureVmSyncedConfigDetailsResponse)(nil)).Elem()
}

func (o AzureToAzureVmSyncedConfigDetailsResponsePtrOutput) ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutput() AzureToAzureVmSyncedConfigDetailsResponsePtrOutput {
	return o
}

func (o AzureToAzureVmSyncedConfigDetailsResponsePtrOutput) ToAzureToAzureVmSyncedConfigDetailsResponsePtrOutputWithContext(ctx context.Context) AzureToAzureVmSyncedConfigDetailsResponsePtrOutput {
	return o
}

func (o AzureToAzureVmSyncedConfigDetailsResponsePtrOutput) Elem() AzureToAzureVmSyncedConfigDetailsResponseOutput {
	return o.ApplyT(func(v *AzureToAzureVmSyncedConfigDetailsResponse) AzureToAzureVmSyncedConfigDetailsResponse {
		if v != nil {
			return *v
		}
		var ret AzureToAzureVmSyncedConfigDetailsResponse
		return ret
	}).(AzureToAzureVmSyncedConfigDetailsResponseOutput)
}

func (o AzureToAzureVmSyncedConfigDetailsResponsePtrOutput) InputEndpoints() InputEndpointResponseArrayOutput {
	return o.ApplyT(func(v *AzureToAzureVmSyncedConfigDetailsResponse) []InputEndpointResponse {
		if v == nil {
			return nil
		}
		return v.InputEndpoints
	}).(InputEndpointResponseArrayOutput)
}

func (o AzureToAzureVmSyncedConfigDetailsResponsePtrOutput) RoleAssignments() RoleAssignmentResponseArrayOutput {
	return o.ApplyT(func(v *AzureToAzureVmSyncedConfigDetailsResponse) []RoleAssignmentResponse {
		if v == nil {
			return nil
		}
		return v.RoleAssignments
	}).(RoleAssignmentResponseArrayOutput)
}

func (o AzureToAzureVmSyncedConfigDetailsResponsePtrOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AzureToAzureVmSyncedConfigDetailsResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringMapOutput)
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





type AzureVmDiskDetailsResponseInput interface {
	pulumi.Input

	ToAzureVmDiskDetailsResponseOutput() AzureVmDiskDetailsResponseOutput
	ToAzureVmDiskDetailsResponseOutputWithContext(context.Context) AzureVmDiskDetailsResponseOutput
}

type AzureVmDiskDetailsResponseArgs struct {
	LunId              pulumi.StringPtrInput `pulumi:"lunId"`
	MaxSizeMB          pulumi.StringPtrInput `pulumi:"maxSizeMB"`
	TargetDiskLocation pulumi.StringPtrInput `pulumi:"targetDiskLocation"`
	TargetDiskName     pulumi.StringPtrInput `pulumi:"targetDiskName"`
	VhdId              pulumi.StringPtrInput `pulumi:"vhdId"`
	VhdName            pulumi.StringPtrInput `pulumi:"vhdName"`
	VhdType            pulumi.StringPtrInput `pulumi:"vhdType"`
}

func (AzureVmDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmDiskDetailsResponse)(nil)).Elem()
}

func (i AzureVmDiskDetailsResponseArgs) ToAzureVmDiskDetailsResponseOutput() AzureVmDiskDetailsResponseOutput {
	return i.ToAzureVmDiskDetailsResponseOutputWithContext(context.Background())
}

func (i AzureVmDiskDetailsResponseArgs) ToAzureVmDiskDetailsResponseOutputWithContext(ctx context.Context) AzureVmDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmDiskDetailsResponseOutput)
}





type AzureVmDiskDetailsResponseArrayInput interface {
	pulumi.Input

	ToAzureVmDiskDetailsResponseArrayOutput() AzureVmDiskDetailsResponseArrayOutput
	ToAzureVmDiskDetailsResponseArrayOutputWithContext(context.Context) AzureVmDiskDetailsResponseArrayOutput
}

type AzureVmDiskDetailsResponseArray []AzureVmDiskDetailsResponseInput

func (AzureVmDiskDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureVmDiskDetailsResponse)(nil)).Elem()
}

func (i AzureVmDiskDetailsResponseArray) ToAzureVmDiskDetailsResponseArrayOutput() AzureVmDiskDetailsResponseArrayOutput {
	return i.ToAzureVmDiskDetailsResponseArrayOutputWithContext(context.Background())
}

func (i AzureVmDiskDetailsResponseArray) ToAzureVmDiskDetailsResponseArrayOutputWithContext(ctx context.Context) AzureVmDiskDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureVmDiskDetailsResponseArrayOutput)
}

type AzureVmDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (AzureVmDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureVmDiskDetailsResponse)(nil)).Elem()
}

func (o AzureVmDiskDetailsResponseOutput) ToAzureVmDiskDetailsResponseOutput() AzureVmDiskDetailsResponseOutput {
	return o
}

func (o AzureVmDiskDetailsResponseOutput) ToAzureVmDiskDetailsResponseOutputWithContext(ctx context.Context) AzureVmDiskDetailsResponseOutput {
	return o
}

func (o AzureVmDiskDetailsResponseOutput) LunId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmDiskDetailsResponse) *string { return v.LunId }).(pulumi.StringPtrOutput)
}

func (o AzureVmDiskDetailsResponseOutput) MaxSizeMB() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmDiskDetailsResponse) *string { return v.MaxSizeMB }).(pulumi.StringPtrOutput)
}

func (o AzureVmDiskDetailsResponseOutput) TargetDiskLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmDiskDetailsResponse) *string { return v.TargetDiskLocation }).(pulumi.StringPtrOutput)
}

func (o AzureVmDiskDetailsResponseOutput) TargetDiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmDiskDetailsResponse) *string { return v.TargetDiskName }).(pulumi.StringPtrOutput)
}

func (o AzureVmDiskDetailsResponseOutput) VhdId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmDiskDetailsResponse) *string { return v.VhdId }).(pulumi.StringPtrOutput)
}

func (o AzureVmDiskDetailsResponseOutput) VhdName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmDiskDetailsResponse) *string { return v.VhdName }).(pulumi.StringPtrOutput)
}

func (o AzureVmDiskDetailsResponseOutput) VhdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureVmDiskDetailsResponse) *string { return v.VhdType }).(pulumi.StringPtrOutput)
}

type AzureVmDiskDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureVmDiskDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureVmDiskDetailsResponse)(nil)).Elem()
}

func (o AzureVmDiskDetailsResponseArrayOutput) ToAzureVmDiskDetailsResponseArrayOutput() AzureVmDiskDetailsResponseArrayOutput {
	return o
}

func (o AzureVmDiskDetailsResponseArrayOutput) ToAzureVmDiskDetailsResponseArrayOutputWithContext(ctx context.Context) AzureVmDiskDetailsResponseArrayOutput {
	return o
}

func (o AzureVmDiskDetailsResponseArrayOutput) Index(i pulumi.IntInput) AzureVmDiskDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureVmDiskDetailsResponse {
		return vs[0].([]AzureVmDiskDetailsResponse)[vs[1].(int)]
	}).(AzureVmDiskDetailsResponseOutput)
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
	FailoverDeploymentModel *string             `pulumi:"failoverDeploymentModel"`
	Groups                  []RecoveryPlanGroup `pulumi:"groups"`
	PrimaryFabricId         string              `pulumi:"primaryFabricId"`
	RecoveryFabricId        string              `pulumi:"recoveryFabricId"`
}





type CreateRecoveryPlanInputPropertiesInput interface {
	pulumi.Input

	ToCreateRecoveryPlanInputPropertiesOutput() CreateRecoveryPlanInputPropertiesOutput
	ToCreateRecoveryPlanInputPropertiesOutputWithContext(context.Context) CreateRecoveryPlanInputPropertiesOutput
}

type CreateRecoveryPlanInputPropertiesArgs struct {
	FailoverDeploymentModel pulumi.StringPtrInput       `pulumi:"failoverDeploymentModel"`
	Groups                  RecoveryPlanGroupArrayInput `pulumi:"groups"`
	PrimaryFabricId         pulumi.StringInput          `pulumi:"primaryFabricId"`
	RecoveryFabricId        pulumi.StringInput          `pulumi:"recoveryFabricId"`
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

func (i CreateRecoveryPlanInputPropertiesArgs) ToCreateRecoveryPlanInputPropertiesPtrOutput() CreateRecoveryPlanInputPropertiesPtrOutput {
	return i.ToCreateRecoveryPlanInputPropertiesPtrOutputWithContext(context.Background())
}

func (i CreateRecoveryPlanInputPropertiesArgs) ToCreateRecoveryPlanInputPropertiesPtrOutputWithContext(ctx context.Context) CreateRecoveryPlanInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateRecoveryPlanInputPropertiesOutput).ToCreateRecoveryPlanInputPropertiesPtrOutputWithContext(ctx)
}









type CreateRecoveryPlanInputPropertiesPtrInput interface {
	pulumi.Input

	ToCreateRecoveryPlanInputPropertiesPtrOutput() CreateRecoveryPlanInputPropertiesPtrOutput
	ToCreateRecoveryPlanInputPropertiesPtrOutputWithContext(context.Context) CreateRecoveryPlanInputPropertiesPtrOutput
}

type createRecoveryPlanInputPropertiesPtrType CreateRecoveryPlanInputPropertiesArgs

func CreateRecoveryPlanInputPropertiesPtr(v *CreateRecoveryPlanInputPropertiesArgs) CreateRecoveryPlanInputPropertiesPtrInput {
	return (*createRecoveryPlanInputPropertiesPtrType)(v)
}

func (*createRecoveryPlanInputPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateRecoveryPlanInputProperties)(nil)).Elem()
}

func (i *createRecoveryPlanInputPropertiesPtrType) ToCreateRecoveryPlanInputPropertiesPtrOutput() CreateRecoveryPlanInputPropertiesPtrOutput {
	return i.ToCreateRecoveryPlanInputPropertiesPtrOutputWithContext(context.Background())
}

func (i *createRecoveryPlanInputPropertiesPtrType) ToCreateRecoveryPlanInputPropertiesPtrOutputWithContext(ctx context.Context) CreateRecoveryPlanInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateRecoveryPlanInputPropertiesPtrOutput)
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

func (o CreateRecoveryPlanInputPropertiesOutput) ToCreateRecoveryPlanInputPropertiesPtrOutput() CreateRecoveryPlanInputPropertiesPtrOutput {
	return o.ToCreateRecoveryPlanInputPropertiesPtrOutputWithContext(context.Background())
}

func (o CreateRecoveryPlanInputPropertiesOutput) ToCreateRecoveryPlanInputPropertiesPtrOutputWithContext(ctx context.Context) CreateRecoveryPlanInputPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateRecoveryPlanInputProperties) *CreateRecoveryPlanInputProperties {
		return &v
	}).(CreateRecoveryPlanInputPropertiesPtrOutput)
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

func (o CreateRecoveryPlanInputPropertiesOutput) RecoveryFabricId() pulumi.StringOutput {
	return o.ApplyT(func(v CreateRecoveryPlanInputProperties) string { return v.RecoveryFabricId }).(pulumi.StringOutput)
}

type CreateRecoveryPlanInputPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CreateRecoveryPlanInputPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateRecoveryPlanInputProperties)(nil)).Elem()
}

func (o CreateRecoveryPlanInputPropertiesPtrOutput) ToCreateRecoveryPlanInputPropertiesPtrOutput() CreateRecoveryPlanInputPropertiesPtrOutput {
	return o
}

func (o CreateRecoveryPlanInputPropertiesPtrOutput) ToCreateRecoveryPlanInputPropertiesPtrOutputWithContext(ctx context.Context) CreateRecoveryPlanInputPropertiesPtrOutput {
	return o
}

func (o CreateRecoveryPlanInputPropertiesPtrOutput) Elem() CreateRecoveryPlanInputPropertiesOutput {
	return o.ApplyT(func(v *CreateRecoveryPlanInputProperties) CreateRecoveryPlanInputProperties {
		if v != nil {
			return *v
		}
		var ret CreateRecoveryPlanInputProperties
		return ret
	}).(CreateRecoveryPlanInputPropertiesOutput)
}

func (o CreateRecoveryPlanInputPropertiesPtrOutput) FailoverDeploymentModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateRecoveryPlanInputProperties) *string {
		if v == nil {
			return nil
		}
		return v.FailoverDeploymentModel
	}).(pulumi.StringPtrOutput)
}

func (o CreateRecoveryPlanInputPropertiesPtrOutput) Groups() RecoveryPlanGroupArrayOutput {
	return o.ApplyT(func(v *CreateRecoveryPlanInputProperties) []RecoveryPlanGroup {
		if v == nil {
			return nil
		}
		return v.Groups
	}).(RecoveryPlanGroupArrayOutput)
}

func (o CreateRecoveryPlanInputPropertiesPtrOutput) PrimaryFabricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateRecoveryPlanInputProperties) *string {
		if v == nil {
			return nil
		}
		return &v.PrimaryFabricId
	}).(pulumi.StringPtrOutput)
}

func (o CreateRecoveryPlanInputPropertiesPtrOutput) RecoveryFabricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateRecoveryPlanInputProperties) *string {
		if v == nil {
			return nil
		}
		return &v.RecoveryFabricId
	}).(pulumi.StringPtrOutput)
}

type CurrentJobDetailsResponse struct {
	JobId     *string `pulumi:"jobId"`
	JobName   *string `pulumi:"jobName"`
	StartTime *string `pulumi:"startTime"`
}





type CurrentJobDetailsResponseInput interface {
	pulumi.Input

	ToCurrentJobDetailsResponseOutput() CurrentJobDetailsResponseOutput
	ToCurrentJobDetailsResponseOutputWithContext(context.Context) CurrentJobDetailsResponseOutput
}

type CurrentJobDetailsResponseArgs struct {
	JobId     pulumi.StringPtrInput `pulumi:"jobId"`
	JobName   pulumi.StringPtrInput `pulumi:"jobName"`
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
}

func (CurrentJobDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CurrentJobDetailsResponse)(nil)).Elem()
}

func (i CurrentJobDetailsResponseArgs) ToCurrentJobDetailsResponseOutput() CurrentJobDetailsResponseOutput {
	return i.ToCurrentJobDetailsResponseOutputWithContext(context.Background())
}

func (i CurrentJobDetailsResponseArgs) ToCurrentJobDetailsResponseOutputWithContext(ctx context.Context) CurrentJobDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CurrentJobDetailsResponseOutput)
}

func (i CurrentJobDetailsResponseArgs) ToCurrentJobDetailsResponsePtrOutput() CurrentJobDetailsResponsePtrOutput {
	return i.ToCurrentJobDetailsResponsePtrOutputWithContext(context.Background())
}

func (i CurrentJobDetailsResponseArgs) ToCurrentJobDetailsResponsePtrOutputWithContext(ctx context.Context) CurrentJobDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CurrentJobDetailsResponseOutput).ToCurrentJobDetailsResponsePtrOutputWithContext(ctx)
}









type CurrentJobDetailsResponsePtrInput interface {
	pulumi.Input

	ToCurrentJobDetailsResponsePtrOutput() CurrentJobDetailsResponsePtrOutput
	ToCurrentJobDetailsResponsePtrOutputWithContext(context.Context) CurrentJobDetailsResponsePtrOutput
}

type currentJobDetailsResponsePtrType CurrentJobDetailsResponseArgs

func CurrentJobDetailsResponsePtr(v *CurrentJobDetailsResponseArgs) CurrentJobDetailsResponsePtrInput {
	return (*currentJobDetailsResponsePtrType)(v)
}

func (*currentJobDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CurrentJobDetailsResponse)(nil)).Elem()
}

func (i *currentJobDetailsResponsePtrType) ToCurrentJobDetailsResponsePtrOutput() CurrentJobDetailsResponsePtrOutput {
	return i.ToCurrentJobDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *currentJobDetailsResponsePtrType) ToCurrentJobDetailsResponsePtrOutputWithContext(ctx context.Context) CurrentJobDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CurrentJobDetailsResponsePtrOutput)
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

func (o CurrentJobDetailsResponseOutput) ToCurrentJobDetailsResponsePtrOutput() CurrentJobDetailsResponsePtrOutput {
	return o.ToCurrentJobDetailsResponsePtrOutputWithContext(context.Background())
}

func (o CurrentJobDetailsResponseOutput) ToCurrentJobDetailsResponsePtrOutputWithContext(ctx context.Context) CurrentJobDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CurrentJobDetailsResponse) *CurrentJobDetailsResponse {
		return &v
	}).(CurrentJobDetailsResponsePtrOutput)
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

type CurrentJobDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (CurrentJobDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CurrentJobDetailsResponse)(nil)).Elem()
}

func (o CurrentJobDetailsResponsePtrOutput) ToCurrentJobDetailsResponsePtrOutput() CurrentJobDetailsResponsePtrOutput {
	return o
}

func (o CurrentJobDetailsResponsePtrOutput) ToCurrentJobDetailsResponsePtrOutputWithContext(ctx context.Context) CurrentJobDetailsResponsePtrOutput {
	return o
}

func (o CurrentJobDetailsResponsePtrOutput) Elem() CurrentJobDetailsResponseOutput {
	return o.ApplyT(func(v *CurrentJobDetailsResponse) CurrentJobDetailsResponse {
		if v != nil {
			return *v
		}
		var ret CurrentJobDetailsResponse
		return ret
	}).(CurrentJobDetailsResponseOutput)
}

func (o CurrentJobDetailsResponsePtrOutput) JobId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CurrentJobDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.JobId
	}).(pulumi.StringPtrOutput)
}

func (o CurrentJobDetailsResponsePtrOutput) JobName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CurrentJobDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.JobName
	}).(pulumi.StringPtrOutput)
}

func (o CurrentJobDetailsResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CurrentJobDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type CurrentScenarioDetailsResponse struct {
	JobId        *string `pulumi:"jobId"`
	ScenarioName *string `pulumi:"scenarioName"`
	StartTime    *string `pulumi:"startTime"`
}





type CurrentScenarioDetailsResponseInput interface {
	pulumi.Input

	ToCurrentScenarioDetailsResponseOutput() CurrentScenarioDetailsResponseOutput
	ToCurrentScenarioDetailsResponseOutputWithContext(context.Context) CurrentScenarioDetailsResponseOutput
}

type CurrentScenarioDetailsResponseArgs struct {
	JobId        pulumi.StringPtrInput `pulumi:"jobId"`
	ScenarioName pulumi.StringPtrInput `pulumi:"scenarioName"`
	StartTime    pulumi.StringPtrInput `pulumi:"startTime"`
}

func (CurrentScenarioDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CurrentScenarioDetailsResponse)(nil)).Elem()
}

func (i CurrentScenarioDetailsResponseArgs) ToCurrentScenarioDetailsResponseOutput() CurrentScenarioDetailsResponseOutput {
	return i.ToCurrentScenarioDetailsResponseOutputWithContext(context.Background())
}

func (i CurrentScenarioDetailsResponseArgs) ToCurrentScenarioDetailsResponseOutputWithContext(ctx context.Context) CurrentScenarioDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CurrentScenarioDetailsResponseOutput)
}

func (i CurrentScenarioDetailsResponseArgs) ToCurrentScenarioDetailsResponsePtrOutput() CurrentScenarioDetailsResponsePtrOutput {
	return i.ToCurrentScenarioDetailsResponsePtrOutputWithContext(context.Background())
}

func (i CurrentScenarioDetailsResponseArgs) ToCurrentScenarioDetailsResponsePtrOutputWithContext(ctx context.Context) CurrentScenarioDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CurrentScenarioDetailsResponseOutput).ToCurrentScenarioDetailsResponsePtrOutputWithContext(ctx)
}









type CurrentScenarioDetailsResponsePtrInput interface {
	pulumi.Input

	ToCurrentScenarioDetailsResponsePtrOutput() CurrentScenarioDetailsResponsePtrOutput
	ToCurrentScenarioDetailsResponsePtrOutputWithContext(context.Context) CurrentScenarioDetailsResponsePtrOutput
}

type currentScenarioDetailsResponsePtrType CurrentScenarioDetailsResponseArgs

func CurrentScenarioDetailsResponsePtr(v *CurrentScenarioDetailsResponseArgs) CurrentScenarioDetailsResponsePtrInput {
	return (*currentScenarioDetailsResponsePtrType)(v)
}

func (*currentScenarioDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CurrentScenarioDetailsResponse)(nil)).Elem()
}

func (i *currentScenarioDetailsResponsePtrType) ToCurrentScenarioDetailsResponsePtrOutput() CurrentScenarioDetailsResponsePtrOutput {
	return i.ToCurrentScenarioDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *currentScenarioDetailsResponsePtrType) ToCurrentScenarioDetailsResponsePtrOutputWithContext(ctx context.Context) CurrentScenarioDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CurrentScenarioDetailsResponsePtrOutput)
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

func (o CurrentScenarioDetailsResponseOutput) ToCurrentScenarioDetailsResponsePtrOutput() CurrentScenarioDetailsResponsePtrOutput {
	return o.ToCurrentScenarioDetailsResponsePtrOutputWithContext(context.Background())
}

func (o CurrentScenarioDetailsResponseOutput) ToCurrentScenarioDetailsResponsePtrOutputWithContext(ctx context.Context) CurrentScenarioDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CurrentScenarioDetailsResponse) *CurrentScenarioDetailsResponse {
		return &v
	}).(CurrentScenarioDetailsResponsePtrOutput)
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





type DataStoreResponseInput interface {
	pulumi.Input

	ToDataStoreResponseOutput() DataStoreResponseOutput
	ToDataStoreResponseOutputWithContext(context.Context) DataStoreResponseOutput
}

type DataStoreResponseArgs struct {
	Capacity     pulumi.StringPtrInput `pulumi:"capacity"`
	FreeSpace    pulumi.StringPtrInput `pulumi:"freeSpace"`
	SymbolicName pulumi.StringPtrInput `pulumi:"symbolicName"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
	Uuid         pulumi.StringPtrInput `pulumi:"uuid"`
}

func (DataStoreResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreResponse)(nil)).Elem()
}

func (i DataStoreResponseArgs) ToDataStoreResponseOutput() DataStoreResponseOutput {
	return i.ToDataStoreResponseOutputWithContext(context.Background())
}

func (i DataStoreResponseArgs) ToDataStoreResponseOutputWithContext(ctx context.Context) DataStoreResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreResponseOutput)
}





type DataStoreResponseArrayInput interface {
	pulumi.Input

	ToDataStoreResponseArrayOutput() DataStoreResponseArrayOutput
	ToDataStoreResponseArrayOutputWithContext(context.Context) DataStoreResponseArrayOutput
}

type DataStoreResponseArray []DataStoreResponseInput

func (DataStoreResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreResponse)(nil)).Elem()
}

func (i DataStoreResponseArray) ToDataStoreResponseArrayOutput() DataStoreResponseArrayOutput {
	return i.ToDataStoreResponseArrayOutputWithContext(context.Background())
}

func (i DataStoreResponseArray) ToDataStoreResponseArrayOutputWithContext(ctx context.Context) DataStoreResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataStoreResponseArrayOutput)
}

type DataStoreResponseOutput struct{ *pulumi.OutputState }

func (DataStoreResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataStoreResponse)(nil)).Elem()
}

func (o DataStoreResponseOutput) ToDataStoreResponseOutput() DataStoreResponseOutput {
	return o
}

func (o DataStoreResponseOutput) ToDataStoreResponseOutputWithContext(ctx context.Context) DataStoreResponseOutput {
	return o
}

func (o DataStoreResponseOutput) Capacity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataStoreResponse) *string { return v.Capacity }).(pulumi.StringPtrOutput)
}

func (o DataStoreResponseOutput) FreeSpace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataStoreResponse) *string { return v.FreeSpace }).(pulumi.StringPtrOutput)
}

func (o DataStoreResponseOutput) SymbolicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataStoreResponse) *string { return v.SymbolicName }).(pulumi.StringPtrOutput)
}

func (o DataStoreResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataStoreResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o DataStoreResponseOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataStoreResponse) *string { return v.Uuid }).(pulumi.StringPtrOutput)
}

type DataStoreResponseArrayOutput struct{ *pulumi.OutputState }

func (DataStoreResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataStoreResponse)(nil)).Elem()
}

func (o DataStoreResponseArrayOutput) ToDataStoreResponseArrayOutput() DataStoreResponseArrayOutput {
	return o
}

func (o DataStoreResponseArrayOutput) ToDataStoreResponseArrayOutputWithContext(ctx context.Context) DataStoreResponseArrayOutput {
	return o
}

func (o DataStoreResponseArrayOutput) Index(i pulumi.IntInput) DataStoreResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataStoreResponse {
		return vs[0].([]DataStoreResponse)[vs[1].(int)]
	}).(DataStoreResponseOutput)
}

type DiskDetailsResponse struct {
	MaxSizeMB *float64 `pulumi:"maxSizeMB"`
	VhdId     *string  `pulumi:"vhdId"`
	VhdName   *string  `pulumi:"vhdName"`
	VhdType   *string  `pulumi:"vhdType"`
}





type DiskDetailsResponseInput interface {
	pulumi.Input

	ToDiskDetailsResponseOutput() DiskDetailsResponseOutput
	ToDiskDetailsResponseOutputWithContext(context.Context) DiskDetailsResponseOutput
}

type DiskDetailsResponseArgs struct {
	MaxSizeMB pulumi.Float64PtrInput `pulumi:"maxSizeMB"`
	VhdId     pulumi.StringPtrInput  `pulumi:"vhdId"`
	VhdName   pulumi.StringPtrInput  `pulumi:"vhdName"`
	VhdType   pulumi.StringPtrInput  `pulumi:"vhdType"`
}

func (DiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskDetailsResponse)(nil)).Elem()
}

func (i DiskDetailsResponseArgs) ToDiskDetailsResponseOutput() DiskDetailsResponseOutput {
	return i.ToDiskDetailsResponseOutputWithContext(context.Background())
}

func (i DiskDetailsResponseArgs) ToDiskDetailsResponseOutputWithContext(ctx context.Context) DiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskDetailsResponseOutput)
}





type DiskDetailsResponseArrayInput interface {
	pulumi.Input

	ToDiskDetailsResponseArrayOutput() DiskDetailsResponseArrayOutput
	ToDiskDetailsResponseArrayOutputWithContext(context.Context) DiskDetailsResponseArrayOutput
}

type DiskDetailsResponseArray []DiskDetailsResponseInput

func (DiskDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskDetailsResponse)(nil)).Elem()
}

func (i DiskDetailsResponseArray) ToDiskDetailsResponseArrayOutput() DiskDetailsResponseArrayOutput {
	return i.ToDiskDetailsResponseArrayOutputWithContext(context.Background())
}

func (i DiskDetailsResponseArray) ToDiskDetailsResponseArrayOutputWithContext(ctx context.Context) DiskDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskDetailsResponseArrayOutput)
}

type DiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (DiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskDetailsResponse)(nil)).Elem()
}

func (o DiskDetailsResponseOutput) ToDiskDetailsResponseOutput() DiskDetailsResponseOutput {
	return o
}

func (o DiskDetailsResponseOutput) ToDiskDetailsResponseOutputWithContext(ctx context.Context) DiskDetailsResponseOutput {
	return o
}

func (o DiskDetailsResponseOutput) MaxSizeMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *float64 { return v.MaxSizeMB }).(pulumi.Float64PtrOutput)
}

func (o DiskDetailsResponseOutput) VhdId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *string { return v.VhdId }).(pulumi.StringPtrOutput)
}

func (o DiskDetailsResponseOutput) VhdName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *string { return v.VhdName }).(pulumi.StringPtrOutput)
}

func (o DiskDetailsResponseOutput) VhdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskDetailsResponse) *string { return v.VhdType }).(pulumi.StringPtrOutput)
}

type DiskDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (DiskDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DiskDetailsResponse)(nil)).Elem()
}

func (o DiskDetailsResponseArrayOutput) ToDiskDetailsResponseArrayOutput() DiskDetailsResponseArrayOutput {
	return o
}

func (o DiskDetailsResponseArrayOutput) ToDiskDetailsResponseArrayOutputWithContext(ctx context.Context) DiskDetailsResponseArrayOutput {
	return o
}

func (o DiskDetailsResponseArrayOutput) Index(i pulumi.IntInput) DiskDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DiskDetailsResponse {
		return vs[0].([]DiskDetailsResponse)[vs[1].(int)]
	}).(DiskDetailsResponseOutput)
}

type DiskEncryptionInfo struct {
	DiskEncryptionKeyInfo *DiskEncryptionKeyInfo `pulumi:"diskEncryptionKeyInfo"`
	KeyEncryptionKeyInfo  *KeyEncryptionKeyInfo  `pulumi:"keyEncryptionKeyInfo"`
}





type DiskEncryptionInfoInput interface {
	pulumi.Input

	ToDiskEncryptionInfoOutput() DiskEncryptionInfoOutput
	ToDiskEncryptionInfoOutputWithContext(context.Context) DiskEncryptionInfoOutput
}

type DiskEncryptionInfoArgs struct {
	DiskEncryptionKeyInfo DiskEncryptionKeyInfoPtrInput `pulumi:"diskEncryptionKeyInfo"`
	KeyEncryptionKeyInfo  KeyEncryptionKeyInfoPtrInput  `pulumi:"keyEncryptionKeyInfo"`
}

func (DiskEncryptionInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionInfo)(nil)).Elem()
}

func (i DiskEncryptionInfoArgs) ToDiskEncryptionInfoOutput() DiskEncryptionInfoOutput {
	return i.ToDiskEncryptionInfoOutputWithContext(context.Background())
}

func (i DiskEncryptionInfoArgs) ToDiskEncryptionInfoOutputWithContext(ctx context.Context) DiskEncryptionInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionInfoOutput)
}

func (i DiskEncryptionInfoArgs) ToDiskEncryptionInfoPtrOutput() DiskEncryptionInfoPtrOutput {
	return i.ToDiskEncryptionInfoPtrOutputWithContext(context.Background())
}

func (i DiskEncryptionInfoArgs) ToDiskEncryptionInfoPtrOutputWithContext(ctx context.Context) DiskEncryptionInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionInfoOutput).ToDiskEncryptionInfoPtrOutputWithContext(ctx)
}









type DiskEncryptionInfoPtrInput interface {
	pulumi.Input

	ToDiskEncryptionInfoPtrOutput() DiskEncryptionInfoPtrOutput
	ToDiskEncryptionInfoPtrOutputWithContext(context.Context) DiskEncryptionInfoPtrOutput
}

type diskEncryptionInfoPtrType DiskEncryptionInfoArgs

func DiskEncryptionInfoPtr(v *DiskEncryptionInfoArgs) DiskEncryptionInfoPtrInput {
	return (*diskEncryptionInfoPtrType)(v)
}

func (*diskEncryptionInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionInfo)(nil)).Elem()
}

func (i *diskEncryptionInfoPtrType) ToDiskEncryptionInfoPtrOutput() DiskEncryptionInfoPtrOutput {
	return i.ToDiskEncryptionInfoPtrOutputWithContext(context.Background())
}

func (i *diskEncryptionInfoPtrType) ToDiskEncryptionInfoPtrOutputWithContext(ctx context.Context) DiskEncryptionInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionInfoPtrOutput)
}

type DiskEncryptionInfoOutput struct{ *pulumi.OutputState }

func (DiskEncryptionInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionInfo)(nil)).Elem()
}

func (o DiskEncryptionInfoOutput) ToDiskEncryptionInfoOutput() DiskEncryptionInfoOutput {
	return o
}

func (o DiskEncryptionInfoOutput) ToDiskEncryptionInfoOutputWithContext(ctx context.Context) DiskEncryptionInfoOutput {
	return o
}

func (o DiskEncryptionInfoOutput) ToDiskEncryptionInfoPtrOutput() DiskEncryptionInfoPtrOutput {
	return o.ToDiskEncryptionInfoPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionInfoOutput) ToDiskEncryptionInfoPtrOutputWithContext(ctx context.Context) DiskEncryptionInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskEncryptionInfo) *DiskEncryptionInfo {
		return &v
	}).(DiskEncryptionInfoPtrOutput)
}

func (o DiskEncryptionInfoOutput) DiskEncryptionKeyInfo() DiskEncryptionKeyInfoPtrOutput {
	return o.ApplyT(func(v DiskEncryptionInfo) *DiskEncryptionKeyInfo { return v.DiskEncryptionKeyInfo }).(DiskEncryptionKeyInfoPtrOutput)
}

func (o DiskEncryptionInfoOutput) KeyEncryptionKeyInfo() KeyEncryptionKeyInfoPtrOutput {
	return o.ApplyT(func(v DiskEncryptionInfo) *KeyEncryptionKeyInfo { return v.KeyEncryptionKeyInfo }).(KeyEncryptionKeyInfoPtrOutput)
}

type DiskEncryptionInfoPtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionInfo)(nil)).Elem()
}

func (o DiskEncryptionInfoPtrOutput) ToDiskEncryptionInfoPtrOutput() DiskEncryptionInfoPtrOutput {
	return o
}

func (o DiskEncryptionInfoPtrOutput) ToDiskEncryptionInfoPtrOutputWithContext(ctx context.Context) DiskEncryptionInfoPtrOutput {
	return o
}

func (o DiskEncryptionInfoPtrOutput) Elem() DiskEncryptionInfoOutput {
	return o.ApplyT(func(v *DiskEncryptionInfo) DiskEncryptionInfo {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionInfo
		return ret
	}).(DiskEncryptionInfoOutput)
}

func (o DiskEncryptionInfoPtrOutput) DiskEncryptionKeyInfo() DiskEncryptionKeyInfoPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionInfo) *DiskEncryptionKeyInfo {
		if v == nil {
			return nil
		}
		return v.DiskEncryptionKeyInfo
	}).(DiskEncryptionKeyInfoPtrOutput)
}

func (o DiskEncryptionInfoPtrOutput) KeyEncryptionKeyInfo() KeyEncryptionKeyInfoPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionInfo) *KeyEncryptionKeyInfo {
		if v == nil {
			return nil
		}
		return v.KeyEncryptionKeyInfo
	}).(KeyEncryptionKeyInfoPtrOutput)
}

type DiskEncryptionKeyInfo struct {
	KeyVaultResourceArmId *string `pulumi:"keyVaultResourceArmId"`
	SecretIdentifier      *string `pulumi:"secretIdentifier"`
}





type DiskEncryptionKeyInfoInput interface {
	pulumi.Input

	ToDiskEncryptionKeyInfoOutput() DiskEncryptionKeyInfoOutput
	ToDiskEncryptionKeyInfoOutputWithContext(context.Context) DiskEncryptionKeyInfoOutput
}

type DiskEncryptionKeyInfoArgs struct {
	KeyVaultResourceArmId pulumi.StringPtrInput `pulumi:"keyVaultResourceArmId"`
	SecretIdentifier      pulumi.StringPtrInput `pulumi:"secretIdentifier"`
}

func (DiskEncryptionKeyInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionKeyInfo)(nil)).Elem()
}

func (i DiskEncryptionKeyInfoArgs) ToDiskEncryptionKeyInfoOutput() DiskEncryptionKeyInfoOutput {
	return i.ToDiskEncryptionKeyInfoOutputWithContext(context.Background())
}

func (i DiskEncryptionKeyInfoArgs) ToDiskEncryptionKeyInfoOutputWithContext(ctx context.Context) DiskEncryptionKeyInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionKeyInfoOutput)
}

func (i DiskEncryptionKeyInfoArgs) ToDiskEncryptionKeyInfoPtrOutput() DiskEncryptionKeyInfoPtrOutput {
	return i.ToDiskEncryptionKeyInfoPtrOutputWithContext(context.Background())
}

func (i DiskEncryptionKeyInfoArgs) ToDiskEncryptionKeyInfoPtrOutputWithContext(ctx context.Context) DiskEncryptionKeyInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionKeyInfoOutput).ToDiskEncryptionKeyInfoPtrOutputWithContext(ctx)
}









type DiskEncryptionKeyInfoPtrInput interface {
	pulumi.Input

	ToDiskEncryptionKeyInfoPtrOutput() DiskEncryptionKeyInfoPtrOutput
	ToDiskEncryptionKeyInfoPtrOutputWithContext(context.Context) DiskEncryptionKeyInfoPtrOutput
}

type diskEncryptionKeyInfoPtrType DiskEncryptionKeyInfoArgs

func DiskEncryptionKeyInfoPtr(v *DiskEncryptionKeyInfoArgs) DiskEncryptionKeyInfoPtrInput {
	return (*diskEncryptionKeyInfoPtrType)(v)
}

func (*diskEncryptionKeyInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionKeyInfo)(nil)).Elem()
}

func (i *diskEncryptionKeyInfoPtrType) ToDiskEncryptionKeyInfoPtrOutput() DiskEncryptionKeyInfoPtrOutput {
	return i.ToDiskEncryptionKeyInfoPtrOutputWithContext(context.Background())
}

func (i *diskEncryptionKeyInfoPtrType) ToDiskEncryptionKeyInfoPtrOutputWithContext(ctx context.Context) DiskEncryptionKeyInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskEncryptionKeyInfoPtrOutput)
}

type DiskEncryptionKeyInfoOutput struct{ *pulumi.OutputState }

func (DiskEncryptionKeyInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskEncryptionKeyInfo)(nil)).Elem()
}

func (o DiskEncryptionKeyInfoOutput) ToDiskEncryptionKeyInfoOutput() DiskEncryptionKeyInfoOutput {
	return o
}

func (o DiskEncryptionKeyInfoOutput) ToDiskEncryptionKeyInfoOutputWithContext(ctx context.Context) DiskEncryptionKeyInfoOutput {
	return o
}

func (o DiskEncryptionKeyInfoOutput) ToDiskEncryptionKeyInfoPtrOutput() DiskEncryptionKeyInfoPtrOutput {
	return o.ToDiskEncryptionKeyInfoPtrOutputWithContext(context.Background())
}

func (o DiskEncryptionKeyInfoOutput) ToDiskEncryptionKeyInfoPtrOutputWithContext(ctx context.Context) DiskEncryptionKeyInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskEncryptionKeyInfo) *DiskEncryptionKeyInfo {
		return &v
	}).(DiskEncryptionKeyInfoPtrOutput)
}

func (o DiskEncryptionKeyInfoOutput) KeyVaultResourceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionKeyInfo) *string { return v.KeyVaultResourceArmId }).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionKeyInfoOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiskEncryptionKeyInfo) *string { return v.SecretIdentifier }).(pulumi.StringPtrOutput)
}

type DiskEncryptionKeyInfoPtrOutput struct{ *pulumi.OutputState }

func (DiskEncryptionKeyInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskEncryptionKeyInfo)(nil)).Elem()
}

func (o DiskEncryptionKeyInfoPtrOutput) ToDiskEncryptionKeyInfoPtrOutput() DiskEncryptionKeyInfoPtrOutput {
	return o
}

func (o DiskEncryptionKeyInfoPtrOutput) ToDiskEncryptionKeyInfoPtrOutputWithContext(ctx context.Context) DiskEncryptionKeyInfoPtrOutput {
	return o
}

func (o DiskEncryptionKeyInfoPtrOutput) Elem() DiskEncryptionKeyInfoOutput {
	return o.ApplyT(func(v *DiskEncryptionKeyInfo) DiskEncryptionKeyInfo {
		if v != nil {
			return *v
		}
		var ret DiskEncryptionKeyInfo
		return ret
	}).(DiskEncryptionKeyInfoOutput)
}

func (o DiskEncryptionKeyInfoPtrOutput) KeyVaultResourceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionKeyInfo) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultResourceArmId
	}).(pulumi.StringPtrOutput)
}

func (o DiskEncryptionKeyInfoPtrOutput) SecretIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DiskEncryptionKeyInfo) *string {
		if v == nil {
			return nil
		}
		return v.SecretIdentifier
	}).(pulumi.StringPtrOutput)
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

func (i EnableMigrationInputPropertiesArgs) ToEnableMigrationInputPropertiesPtrOutput() EnableMigrationInputPropertiesPtrOutput {
	return i.ToEnableMigrationInputPropertiesPtrOutputWithContext(context.Background())
}

func (i EnableMigrationInputPropertiesArgs) ToEnableMigrationInputPropertiesPtrOutputWithContext(ctx context.Context) EnableMigrationInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnableMigrationInputPropertiesOutput).ToEnableMigrationInputPropertiesPtrOutputWithContext(ctx)
}









type EnableMigrationInputPropertiesPtrInput interface {
	pulumi.Input

	ToEnableMigrationInputPropertiesPtrOutput() EnableMigrationInputPropertiesPtrOutput
	ToEnableMigrationInputPropertiesPtrOutputWithContext(context.Context) EnableMigrationInputPropertiesPtrOutput
}

type enableMigrationInputPropertiesPtrType EnableMigrationInputPropertiesArgs

func EnableMigrationInputPropertiesPtr(v *EnableMigrationInputPropertiesArgs) EnableMigrationInputPropertiesPtrInput {
	return (*enableMigrationInputPropertiesPtrType)(v)
}

func (*enableMigrationInputPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnableMigrationInputProperties)(nil)).Elem()
}

func (i *enableMigrationInputPropertiesPtrType) ToEnableMigrationInputPropertiesPtrOutput() EnableMigrationInputPropertiesPtrOutput {
	return i.ToEnableMigrationInputPropertiesPtrOutputWithContext(context.Background())
}

func (i *enableMigrationInputPropertiesPtrType) ToEnableMigrationInputPropertiesPtrOutputWithContext(ctx context.Context) EnableMigrationInputPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnableMigrationInputPropertiesPtrOutput)
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

func (o EnableMigrationInputPropertiesOutput) ToEnableMigrationInputPropertiesPtrOutput() EnableMigrationInputPropertiesPtrOutput {
	return o.ToEnableMigrationInputPropertiesPtrOutputWithContext(context.Background())
}

func (o EnableMigrationInputPropertiesOutput) ToEnableMigrationInputPropertiesPtrOutputWithContext(ctx context.Context) EnableMigrationInputPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnableMigrationInputProperties) *EnableMigrationInputProperties {
		return &v
	}).(EnableMigrationInputPropertiesPtrOutput)
}

func (o EnableMigrationInputPropertiesOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v EnableMigrationInputProperties) string { return v.PolicyId }).(pulumi.StringOutput)
}

func (o EnableMigrationInputPropertiesOutput) ProviderSpecificDetails() VMwareCbtEnableMigrationInputOutput {
	return o.ApplyT(func(v EnableMigrationInputProperties) VMwareCbtEnableMigrationInput { return v.ProviderSpecificDetails }).(VMwareCbtEnableMigrationInputOutput)
}

type EnableMigrationInputPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EnableMigrationInputPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnableMigrationInputProperties)(nil)).Elem()
}

func (o EnableMigrationInputPropertiesPtrOutput) ToEnableMigrationInputPropertiesPtrOutput() EnableMigrationInputPropertiesPtrOutput {
	return o
}

func (o EnableMigrationInputPropertiesPtrOutput) ToEnableMigrationInputPropertiesPtrOutputWithContext(ctx context.Context) EnableMigrationInputPropertiesPtrOutput {
	return o
}

func (o EnableMigrationInputPropertiesPtrOutput) Elem() EnableMigrationInputPropertiesOutput {
	return o.ApplyT(func(v *EnableMigrationInputProperties) EnableMigrationInputProperties {
		if v != nil {
			return *v
		}
		var ret EnableMigrationInputProperties
		return ret
	}).(EnableMigrationInputPropertiesOutput)
}

func (o EnableMigrationInputPropertiesPtrOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnableMigrationInputProperties) *string {
		if v == nil {
			return nil
		}
		return &v.PolicyId
	}).(pulumi.StringPtrOutput)
}

func (o EnableMigrationInputPropertiesPtrOutput) ProviderSpecificDetails() VMwareCbtEnableMigrationInputPtrOutput {
	return o.ApplyT(func(v *EnableMigrationInputProperties) *VMwareCbtEnableMigrationInput {
		if v == nil {
			return nil
		}
		return &v.ProviderSpecificDetails
	}).(VMwareCbtEnableMigrationInputPtrOutput)
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





type EncryptionDetailsResponseInput interface {
	pulumi.Input

	ToEncryptionDetailsResponseOutput() EncryptionDetailsResponseOutput
	ToEncryptionDetailsResponseOutputWithContext(context.Context) EncryptionDetailsResponseOutput
}

type EncryptionDetailsResponseArgs struct {
	KekCertExpiryDate pulumi.StringPtrInput `pulumi:"kekCertExpiryDate"`
	KekCertThumbprint pulumi.StringPtrInput `pulumi:"kekCertThumbprint"`
	KekState          pulumi.StringPtrInput `pulumi:"kekState"`
}

func (EncryptionDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionDetailsResponse)(nil)).Elem()
}

func (i EncryptionDetailsResponseArgs) ToEncryptionDetailsResponseOutput() EncryptionDetailsResponseOutput {
	return i.ToEncryptionDetailsResponseOutputWithContext(context.Background())
}

func (i EncryptionDetailsResponseArgs) ToEncryptionDetailsResponseOutputWithContext(ctx context.Context) EncryptionDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionDetailsResponseOutput)
}

func (i EncryptionDetailsResponseArgs) ToEncryptionDetailsResponsePtrOutput() EncryptionDetailsResponsePtrOutput {
	return i.ToEncryptionDetailsResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionDetailsResponseArgs) ToEncryptionDetailsResponsePtrOutputWithContext(ctx context.Context) EncryptionDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionDetailsResponseOutput).ToEncryptionDetailsResponsePtrOutputWithContext(ctx)
}









type EncryptionDetailsResponsePtrInput interface {
	pulumi.Input

	ToEncryptionDetailsResponsePtrOutput() EncryptionDetailsResponsePtrOutput
	ToEncryptionDetailsResponsePtrOutputWithContext(context.Context) EncryptionDetailsResponsePtrOutput
}

type encryptionDetailsResponsePtrType EncryptionDetailsResponseArgs

func EncryptionDetailsResponsePtr(v *EncryptionDetailsResponseArgs) EncryptionDetailsResponsePtrInput {
	return (*encryptionDetailsResponsePtrType)(v)
}

func (*encryptionDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionDetailsResponse)(nil)).Elem()
}

func (i *encryptionDetailsResponsePtrType) ToEncryptionDetailsResponsePtrOutput() EncryptionDetailsResponsePtrOutput {
	return i.ToEncryptionDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionDetailsResponsePtrType) ToEncryptionDetailsResponsePtrOutputWithContext(ctx context.Context) EncryptionDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionDetailsResponsePtrOutput)
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

func (o EncryptionDetailsResponseOutput) ToEncryptionDetailsResponsePtrOutput() EncryptionDetailsResponsePtrOutput {
	return o.ToEncryptionDetailsResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionDetailsResponseOutput) ToEncryptionDetailsResponsePtrOutputWithContext(ctx context.Context) EncryptionDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionDetailsResponse) *EncryptionDetailsResponse {
		return &v
	}).(EncryptionDetailsResponsePtrOutput)
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





type FabricPropertiesResponseInput interface {
	pulumi.Input

	ToFabricPropertiesResponseOutput() FabricPropertiesResponseOutput
	ToFabricPropertiesResponseOutputWithContext(context.Context) FabricPropertiesResponseOutput
}

type FabricPropertiesResponseArgs struct {
	BcdrState                 pulumi.StringPtrInput             `pulumi:"bcdrState"`
	CustomDetails             pulumi.Input                      `pulumi:"customDetails"`
	EncryptionDetails         EncryptionDetailsResponsePtrInput `pulumi:"encryptionDetails"`
	FriendlyName              pulumi.StringPtrInput             `pulumi:"friendlyName"`
	Health                    pulumi.StringPtrInput             `pulumi:"health"`
	HealthErrorDetails        HealthErrorResponseArrayInput     `pulumi:"healthErrorDetails"`
	InternalIdentifier        pulumi.StringPtrInput             `pulumi:"internalIdentifier"`
	RolloverEncryptionDetails EncryptionDetailsResponsePtrInput `pulumi:"rolloverEncryptionDetails"`
}

func (FabricPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FabricPropertiesResponse)(nil)).Elem()
}

func (i FabricPropertiesResponseArgs) ToFabricPropertiesResponseOutput() FabricPropertiesResponseOutput {
	return i.ToFabricPropertiesResponseOutputWithContext(context.Background())
}

func (i FabricPropertiesResponseArgs) ToFabricPropertiesResponseOutputWithContext(ctx context.Context) FabricPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricPropertiesResponseOutput)
}

func (i FabricPropertiesResponseArgs) ToFabricPropertiesResponsePtrOutput() FabricPropertiesResponsePtrOutput {
	return i.ToFabricPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i FabricPropertiesResponseArgs) ToFabricPropertiesResponsePtrOutputWithContext(ctx context.Context) FabricPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricPropertiesResponseOutput).ToFabricPropertiesResponsePtrOutputWithContext(ctx)
}









type FabricPropertiesResponsePtrInput interface {
	pulumi.Input

	ToFabricPropertiesResponsePtrOutput() FabricPropertiesResponsePtrOutput
	ToFabricPropertiesResponsePtrOutputWithContext(context.Context) FabricPropertiesResponsePtrOutput
}

type fabricPropertiesResponsePtrType FabricPropertiesResponseArgs

func FabricPropertiesResponsePtr(v *FabricPropertiesResponseArgs) FabricPropertiesResponsePtrInput {
	return (*fabricPropertiesResponsePtrType)(v)
}

func (*fabricPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FabricPropertiesResponse)(nil)).Elem()
}

func (i *fabricPropertiesResponsePtrType) ToFabricPropertiesResponsePtrOutput() FabricPropertiesResponsePtrOutput {
	return i.ToFabricPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *fabricPropertiesResponsePtrType) ToFabricPropertiesResponsePtrOutputWithContext(ctx context.Context) FabricPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FabricPropertiesResponsePtrOutput)
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

func (o FabricPropertiesResponseOutput) ToFabricPropertiesResponsePtrOutput() FabricPropertiesResponsePtrOutput {
	return o.ToFabricPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o FabricPropertiesResponseOutput) ToFabricPropertiesResponsePtrOutputWithContext(ctx context.Context) FabricPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FabricPropertiesResponse) *FabricPropertiesResponse {
		return &v
	}).(FabricPropertiesResponsePtrOutput)
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

type FabricPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (FabricPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FabricPropertiesResponse)(nil)).Elem()
}

func (o FabricPropertiesResponsePtrOutput) ToFabricPropertiesResponsePtrOutput() FabricPropertiesResponsePtrOutput {
	return o
}

func (o FabricPropertiesResponsePtrOutput) ToFabricPropertiesResponsePtrOutputWithContext(ctx context.Context) FabricPropertiesResponsePtrOutput {
	return o
}

func (o FabricPropertiesResponsePtrOutput) Elem() FabricPropertiesResponseOutput {
	return o.ApplyT(func(v *FabricPropertiesResponse) FabricPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret FabricPropertiesResponse
		return ret
	}).(FabricPropertiesResponseOutput)
}

func (o FabricPropertiesResponsePtrOutput) BcdrState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FabricPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.BcdrState
	}).(pulumi.StringPtrOutput)
}

func (o FabricPropertiesResponsePtrOutput) CustomDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *FabricPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.CustomDetails
	}).(pulumi.AnyOutput)
}

func (o FabricPropertiesResponsePtrOutput) EncryptionDetails() EncryptionDetailsResponsePtrOutput {
	return o.ApplyT(func(v *FabricPropertiesResponse) *EncryptionDetailsResponse {
		if v == nil {
			return nil
		}
		return v.EncryptionDetails
	}).(EncryptionDetailsResponsePtrOutput)
}

func (o FabricPropertiesResponsePtrOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FabricPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o FabricPropertiesResponsePtrOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FabricPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Health
	}).(pulumi.StringPtrOutput)
}

func (o FabricPropertiesResponsePtrOutput) HealthErrorDetails() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v *FabricPropertiesResponse) []HealthErrorResponse {
		if v == nil {
			return nil
		}
		return v.HealthErrorDetails
	}).(HealthErrorResponseArrayOutput)
}

func (o FabricPropertiesResponsePtrOutput) InternalIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FabricPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.InternalIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o FabricPropertiesResponsePtrOutput) RolloverEncryptionDetails() EncryptionDetailsResponsePtrOutput {
	return o.ApplyT(func(v *FabricPropertiesResponse) *EncryptionDetailsResponse {
		if v == nil {
			return nil
		}
		return v.RolloverEncryptionDetails
	}).(EncryptionDetailsResponsePtrOutput)
}

type HealthErrorResponse struct {
	CreationTimeUtc              *string                    `pulumi:"creationTimeUtc"`
	EntityId                     *string                    `pulumi:"entityId"`
	ErrorCategory                *string                    `pulumi:"errorCategory"`
	ErrorCode                    *string                    `pulumi:"errorCode"`
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





type HealthErrorResponseInput interface {
	pulumi.Input

	ToHealthErrorResponseOutput() HealthErrorResponseOutput
	ToHealthErrorResponseOutputWithContext(context.Context) HealthErrorResponseOutput
}

type HealthErrorResponseArgs struct {
	CreationTimeUtc              pulumi.StringPtrInput              `pulumi:"creationTimeUtc"`
	EntityId                     pulumi.StringPtrInput              `pulumi:"entityId"`
	ErrorCategory                pulumi.StringPtrInput              `pulumi:"errorCategory"`
	ErrorCode                    pulumi.StringPtrInput              `pulumi:"errorCode"`
	ErrorLevel                   pulumi.StringPtrInput              `pulumi:"errorLevel"`
	ErrorMessage                 pulumi.StringPtrInput              `pulumi:"errorMessage"`
	ErrorSource                  pulumi.StringPtrInput              `pulumi:"errorSource"`
	ErrorType                    pulumi.StringPtrInput              `pulumi:"errorType"`
	InnerHealthErrors            InnerHealthErrorResponseArrayInput `pulumi:"innerHealthErrors"`
	PossibleCauses               pulumi.StringPtrInput              `pulumi:"possibleCauses"`
	RecommendedAction            pulumi.StringPtrInput              `pulumi:"recommendedAction"`
	RecoveryProviderErrorMessage pulumi.StringPtrInput              `pulumi:"recoveryProviderErrorMessage"`
	SummaryMessage               pulumi.StringPtrInput              `pulumi:"summaryMessage"`
}

func (HealthErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthErrorResponse)(nil)).Elem()
}

func (i HealthErrorResponseArgs) ToHealthErrorResponseOutput() HealthErrorResponseOutput {
	return i.ToHealthErrorResponseOutputWithContext(context.Background())
}

func (i HealthErrorResponseArgs) ToHealthErrorResponseOutputWithContext(ctx context.Context) HealthErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthErrorResponseOutput)
}





type HealthErrorResponseArrayInput interface {
	pulumi.Input

	ToHealthErrorResponseArrayOutput() HealthErrorResponseArrayOutput
	ToHealthErrorResponseArrayOutputWithContext(context.Context) HealthErrorResponseArrayOutput
}

type HealthErrorResponseArray []HealthErrorResponseInput

func (HealthErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthErrorResponse)(nil)).Elem()
}

func (i HealthErrorResponseArray) ToHealthErrorResponseArrayOutput() HealthErrorResponseArrayOutput {
	return i.ToHealthErrorResponseArrayOutputWithContext(context.Background())
}

func (i HealthErrorResponseArray) ToHealthErrorResponseArrayOutputWithContext(ctx context.Context) HealthErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthErrorResponseArrayOutput)
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

func (o HealthErrorResponseOutput) EntityId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.EntityId }).(pulumi.StringPtrOutput)
}

func (o HealthErrorResponseOutput) ErrorCategory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthErrorResponse) *string { return v.ErrorCategory }).(pulumi.StringPtrOutput)
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
	DisksToInclude               []string `pulumi:"disksToInclude"`
	EnableRdpOnTargetOption      *string  `pulumi:"enableRdpOnTargetOption"`
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





type HyperVReplicaAzureEnableProtectionInputInput interface {
	pulumi.Input

	ToHyperVReplicaAzureEnableProtectionInputOutput() HyperVReplicaAzureEnableProtectionInputOutput
	ToHyperVReplicaAzureEnableProtectionInputOutputWithContext(context.Context) HyperVReplicaAzureEnableProtectionInputOutput
}

type HyperVReplicaAzureEnableProtectionInputArgs struct {
	DisksToInclude               pulumi.StringArrayInput `pulumi:"disksToInclude"`
	EnableRdpOnTargetOption      pulumi.StringPtrInput   `pulumi:"enableRdpOnTargetOption"`
	HvHostVmId                   pulumi.StringPtrInput   `pulumi:"hvHostVmId"`
	InstanceType                 pulumi.StringPtrInput   `pulumi:"instanceType"`
	LogStorageAccountId          pulumi.StringPtrInput   `pulumi:"logStorageAccountId"`
	OsType                       pulumi.StringPtrInput   `pulumi:"osType"`
	TargetAzureNetworkId         pulumi.StringPtrInput   `pulumi:"targetAzureNetworkId"`
	TargetAzureSubnetId          pulumi.StringPtrInput   `pulumi:"targetAzureSubnetId"`
	TargetAzureV1ResourceGroupId pulumi.StringPtrInput   `pulumi:"targetAzureV1ResourceGroupId"`
	TargetAzureV2ResourceGroupId pulumi.StringPtrInput   `pulumi:"targetAzureV2ResourceGroupId"`
	TargetAzureVmName            pulumi.StringPtrInput   `pulumi:"targetAzureVmName"`
	TargetStorageAccountId       pulumi.StringPtrInput   `pulumi:"targetStorageAccountId"`
	UseManagedDisks              pulumi.StringPtrInput   `pulumi:"useManagedDisks"`
	VhdId                        pulumi.StringPtrInput   `pulumi:"vhdId"`
	VmName                       pulumi.StringPtrInput   `pulumi:"vmName"`
}

func (HyperVReplicaAzureEnableProtectionInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaAzureEnableProtectionInput)(nil)).Elem()
}

func (i HyperVReplicaAzureEnableProtectionInputArgs) ToHyperVReplicaAzureEnableProtectionInputOutput() HyperVReplicaAzureEnableProtectionInputOutput {
	return i.ToHyperVReplicaAzureEnableProtectionInputOutputWithContext(context.Background())
}

func (i HyperVReplicaAzureEnableProtectionInputArgs) ToHyperVReplicaAzureEnableProtectionInputOutputWithContext(ctx context.Context) HyperVReplicaAzureEnableProtectionInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaAzureEnableProtectionInputOutput)
}

type HyperVReplicaAzureEnableProtectionInputOutput struct{ *pulumi.OutputState }

func (HyperVReplicaAzureEnableProtectionInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaAzureEnableProtectionInput)(nil)).Elem()
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) ToHyperVReplicaAzureEnableProtectionInputOutput() HyperVReplicaAzureEnableProtectionInputOutput {
	return o
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) ToHyperVReplicaAzureEnableProtectionInputOutputWithContext(ctx context.Context) HyperVReplicaAzureEnableProtectionInputOutput {
	return o
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) DisksToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) []string { return v.DisksToInclude }).(pulumi.StringArrayOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) EnableRdpOnTargetOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.EnableRdpOnTargetOption }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) HvHostVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.HvHostVmId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) LogStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.LogStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) TargetAzureNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.TargetAzureNetworkId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) TargetAzureSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.TargetAzureSubnetId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) TargetAzureV1ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.TargetAzureV1ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) TargetAzureV2ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.TargetAzureV2ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) TargetAzureVmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.TargetAzureVmName }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) TargetStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.TargetStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) UseManagedDisks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.UseManagedDisks }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) VhdId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.VhdId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureEnableProtectionInputOutput) VmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureEnableProtectionInput) *string { return v.VmName }).(pulumi.StringPtrOutput)
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





type HyperVReplicaAzurePolicyDetailsResponseInput interface {
	pulumi.Input

	ToHyperVReplicaAzurePolicyDetailsResponseOutput() HyperVReplicaAzurePolicyDetailsResponseOutput
	ToHyperVReplicaAzurePolicyDetailsResponseOutputWithContext(context.Context) HyperVReplicaAzurePolicyDetailsResponseOutput
}

type HyperVReplicaAzurePolicyDetailsResponseArgs struct {
	ActiveStorageAccountId                        pulumi.StringPtrInput `pulumi:"activeStorageAccountId"`
	ApplicationConsistentSnapshotFrequencyInHours pulumi.IntPtrInput    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Encryption                                    pulumi.StringPtrInput `pulumi:"encryption"`
	InstanceType                                  pulumi.StringInput    `pulumi:"instanceType"`
	OnlineReplicationStartTime                    pulumi.StringPtrInput `pulumi:"onlineReplicationStartTime"`
	RecoveryPointHistoryDurationInHours           pulumi.IntPtrInput    `pulumi:"recoveryPointHistoryDurationInHours"`
	ReplicationInterval                           pulumi.IntPtrInput    `pulumi:"replicationInterval"`
}

func (HyperVReplicaAzurePolicyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaAzurePolicyDetailsResponse)(nil)).Elem()
}

func (i HyperVReplicaAzurePolicyDetailsResponseArgs) ToHyperVReplicaAzurePolicyDetailsResponseOutput() HyperVReplicaAzurePolicyDetailsResponseOutput {
	return i.ToHyperVReplicaAzurePolicyDetailsResponseOutputWithContext(context.Background())
}

func (i HyperVReplicaAzurePolicyDetailsResponseArgs) ToHyperVReplicaAzurePolicyDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaAzurePolicyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaAzurePolicyDetailsResponseOutput)
}

type HyperVReplicaAzurePolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (HyperVReplicaAzurePolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaAzurePolicyDetailsResponse)(nil)).Elem()
}

func (o HyperVReplicaAzurePolicyDetailsResponseOutput) ToHyperVReplicaAzurePolicyDetailsResponseOutput() HyperVReplicaAzurePolicyDetailsResponseOutput {
	return o
}

func (o HyperVReplicaAzurePolicyDetailsResponseOutput) ToHyperVReplicaAzurePolicyDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaAzurePolicyDetailsResponseOutput {
	return o
}

func (o HyperVReplicaAzurePolicyDetailsResponseOutput) ActiveStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyDetailsResponse) *string { return v.ActiveStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzurePolicyDetailsResponseOutput) ApplicationConsistentSnapshotFrequencyInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyDetailsResponse) *int {
		return v.ApplicationConsistentSnapshotFrequencyInHours
	}).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaAzurePolicyDetailsResponseOutput) Encryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyDetailsResponse) *string { return v.Encryption }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzurePolicyDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o HyperVReplicaAzurePolicyDetailsResponseOutput) OnlineReplicationStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyDetailsResponse) *string { return v.OnlineReplicationStartTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzurePolicyDetailsResponseOutput) RecoveryPointHistoryDurationInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyDetailsResponse) *int { return v.RecoveryPointHistoryDurationInHours }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaAzurePolicyDetailsResponseOutput) ReplicationInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyDetailsResponse) *int { return v.ReplicationInterval }).(pulumi.IntPtrOutput)
}

type HyperVReplicaAzurePolicyInput struct {
	ApplicationConsistentSnapshotFrequencyInHours *int     `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	InstanceType                                  *string  `pulumi:"instanceType"`
	OnlineReplicationStartTime                    *string  `pulumi:"onlineReplicationStartTime"`
	RecoveryPointHistoryDuration                  *int     `pulumi:"recoveryPointHistoryDuration"`
	ReplicationInterval                           *int     `pulumi:"replicationInterval"`
	StorageAccounts                               []string `pulumi:"storageAccounts"`
}





type HyperVReplicaAzurePolicyInputInput interface {
	pulumi.Input

	ToHyperVReplicaAzurePolicyInputOutput() HyperVReplicaAzurePolicyInputOutput
	ToHyperVReplicaAzurePolicyInputOutputWithContext(context.Context) HyperVReplicaAzurePolicyInputOutput
}

type HyperVReplicaAzurePolicyInputArgs struct {
	ApplicationConsistentSnapshotFrequencyInHours pulumi.IntPtrInput      `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	InstanceType                                  pulumi.StringPtrInput   `pulumi:"instanceType"`
	OnlineReplicationStartTime                    pulumi.StringPtrInput   `pulumi:"onlineReplicationStartTime"`
	RecoveryPointHistoryDuration                  pulumi.IntPtrInput      `pulumi:"recoveryPointHistoryDuration"`
	ReplicationInterval                           pulumi.IntPtrInput      `pulumi:"replicationInterval"`
	StorageAccounts                               pulumi.StringArrayInput `pulumi:"storageAccounts"`
}

func (HyperVReplicaAzurePolicyInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaAzurePolicyInput)(nil)).Elem()
}

func (i HyperVReplicaAzurePolicyInputArgs) ToHyperVReplicaAzurePolicyInputOutput() HyperVReplicaAzurePolicyInputOutput {
	return i.ToHyperVReplicaAzurePolicyInputOutputWithContext(context.Background())
}

func (i HyperVReplicaAzurePolicyInputArgs) ToHyperVReplicaAzurePolicyInputOutputWithContext(ctx context.Context) HyperVReplicaAzurePolicyInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaAzurePolicyInputOutput)
}

type HyperVReplicaAzurePolicyInputOutput struct{ *pulumi.OutputState }

func (HyperVReplicaAzurePolicyInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaAzurePolicyInput)(nil)).Elem()
}

func (o HyperVReplicaAzurePolicyInputOutput) ToHyperVReplicaAzurePolicyInputOutput() HyperVReplicaAzurePolicyInputOutput {
	return o
}

func (o HyperVReplicaAzurePolicyInputOutput) ToHyperVReplicaAzurePolicyInputOutputWithContext(ctx context.Context) HyperVReplicaAzurePolicyInputOutput {
	return o
}

func (o HyperVReplicaAzurePolicyInputOutput) ApplicationConsistentSnapshotFrequencyInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyInput) *int { return v.ApplicationConsistentSnapshotFrequencyInHours }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaAzurePolicyInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzurePolicyInputOutput) OnlineReplicationStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyInput) *string { return v.OnlineReplicationStartTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzurePolicyInputOutput) RecoveryPointHistoryDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyInput) *int { return v.RecoveryPointHistoryDuration }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaAzurePolicyInputOutput) ReplicationInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyInput) *int { return v.ReplicationInterval }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaAzurePolicyInputOutput) StorageAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v HyperVReplicaAzurePolicyInput) []string { return v.StorageAccounts }).(pulumi.StringArrayOutput)
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
	UseManagedDisks                  *string                            `pulumi:"useManagedDisks"`
	VmId                             *string                            `pulumi:"vmId"`
	VmNics                           []VMNicDetailsResponse             `pulumi:"vmNics"`
	VmProtectionState                *string                            `pulumi:"vmProtectionState"`
	VmProtectionStateDescription     *string                            `pulumi:"vmProtectionStateDescription"`
}





type HyperVReplicaAzureReplicationDetailsResponseInput interface {
	pulumi.Input

	ToHyperVReplicaAzureReplicationDetailsResponseOutput() HyperVReplicaAzureReplicationDetailsResponseOutput
	ToHyperVReplicaAzureReplicationDetailsResponseOutputWithContext(context.Context) HyperVReplicaAzureReplicationDetailsResponseOutput
}

type HyperVReplicaAzureReplicationDetailsResponseArgs struct {
	AzureVmDiskDetails               AzureVmDiskDetailsResponseArrayInput      `pulumi:"azureVmDiskDetails"`
	EnableRdpOnTargetOption          pulumi.StringPtrInput                     `pulumi:"enableRdpOnTargetOption"`
	Encryption                       pulumi.StringPtrInput                     `pulumi:"encryption"`
	InitialReplicationDetails        InitialReplicationDetailsResponsePtrInput `pulumi:"initialReplicationDetails"`
	InstanceType                     pulumi.StringInput                        `pulumi:"instanceType"`
	LastReplicatedTime               pulumi.StringPtrInput                     `pulumi:"lastReplicatedTime"`
	LastRpoCalculatedTime            pulumi.StringPtrInput                     `pulumi:"lastRpoCalculatedTime"`
	LicenseType                      pulumi.StringPtrInput                     `pulumi:"licenseType"`
	OSDetails                        OSDetailsResponsePtrInput                 `pulumi:"oSDetails"`
	RecoveryAvailabilitySetId        pulumi.StringPtrInput                     `pulumi:"recoveryAvailabilitySetId"`
	RecoveryAzureLogStorageAccountId pulumi.StringPtrInput                     `pulumi:"recoveryAzureLogStorageAccountId"`
	RecoveryAzureResourceGroupId     pulumi.StringPtrInput                     `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureStorageAccount      pulumi.StringPtrInput                     `pulumi:"recoveryAzureStorageAccount"`
	RecoveryAzureVMSize              pulumi.StringPtrInput                     `pulumi:"recoveryAzureVMSize"`
	RecoveryAzureVmName              pulumi.StringPtrInput                     `pulumi:"recoveryAzureVmName"`
	RpoInSeconds                     pulumi.Float64PtrInput                    `pulumi:"rpoInSeconds"`
	SelectedRecoveryAzureNetworkId   pulumi.StringPtrInput                     `pulumi:"selectedRecoveryAzureNetworkId"`
	SelectedSourceNicId              pulumi.StringPtrInput                     `pulumi:"selectedSourceNicId"`
	SourceVmCpuCount                 pulumi.IntPtrInput                        `pulumi:"sourceVmCpuCount"`
	SourceVmRamSizeInMB              pulumi.IntPtrInput                        `pulumi:"sourceVmRamSizeInMB"`
	UseManagedDisks                  pulumi.StringPtrInput                     `pulumi:"useManagedDisks"`
	VmId                             pulumi.StringPtrInput                     `pulumi:"vmId"`
	VmNics                           VMNicDetailsResponseArrayInput            `pulumi:"vmNics"`
	VmProtectionState                pulumi.StringPtrInput                     `pulumi:"vmProtectionState"`
	VmProtectionStateDescription     pulumi.StringPtrInput                     `pulumi:"vmProtectionStateDescription"`
}

func (HyperVReplicaAzureReplicationDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaAzureReplicationDetailsResponse)(nil)).Elem()
}

func (i HyperVReplicaAzureReplicationDetailsResponseArgs) ToHyperVReplicaAzureReplicationDetailsResponseOutput() HyperVReplicaAzureReplicationDetailsResponseOutput {
	return i.ToHyperVReplicaAzureReplicationDetailsResponseOutputWithContext(context.Background())
}

func (i HyperVReplicaAzureReplicationDetailsResponseArgs) ToHyperVReplicaAzureReplicationDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaAzureReplicationDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaAzureReplicationDetailsResponseOutput)
}

type HyperVReplicaAzureReplicationDetailsResponseOutput struct{ *pulumi.OutputState }

func (HyperVReplicaAzureReplicationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaAzureReplicationDetailsResponse)(nil)).Elem()
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) ToHyperVReplicaAzureReplicationDetailsResponseOutput() HyperVReplicaAzureReplicationDetailsResponseOutput {
	return o
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) ToHyperVReplicaAzureReplicationDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaAzureReplicationDetailsResponseOutput {
	return o
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) AzureVmDiskDetails() AzureVmDiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) []AzureVmDiskDetailsResponse {
		return v.AzureVmDiskDetails
	}).(AzureVmDiskDetailsResponseArrayOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) EnableRdpOnTargetOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.EnableRdpOnTargetOption }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) Encryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.Encryption }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) InitialReplicationDetails() InitialReplicationDetailsResponsePtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *InitialReplicationDetailsResponse {
		return v.InitialReplicationDetails
	}).(InitialReplicationDetailsResponsePtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) LastReplicatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.LastReplicatedTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) LastRpoCalculatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.LastRpoCalculatedTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) OSDetails() OSDetailsResponsePtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *OSDetailsResponse { return v.OSDetails }).(OSDetailsResponsePtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) RecoveryAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.RecoveryAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) RecoveryAzureLogStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string {
		return v.RecoveryAzureLogStorageAccountId
	}).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) RecoveryAzureResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.RecoveryAzureResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) RecoveryAzureStorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.RecoveryAzureStorageAccount }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) RecoveryAzureVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.RecoveryAzureVMSize }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) RecoveryAzureVmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.RecoveryAzureVmName }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) RpoInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *float64 { return v.RpoInSeconds }).(pulumi.Float64PtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) SelectedRecoveryAzureNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.SelectedRecoveryAzureNetworkId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) SelectedSourceNicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.SelectedSourceNicId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) SourceVmCpuCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *int { return v.SourceVmCpuCount }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) SourceVmRamSizeInMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *int { return v.SourceVmRamSizeInMB }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) UseManagedDisks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.UseManagedDisks }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) VmNics() VMNicDetailsResponseArrayOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) []VMNicDetailsResponse { return v.VmNics }).(VMNicDetailsResponseArrayOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) VmProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.VmProtectionState }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaAzureReplicationDetailsResponseOutput) VmProtectionStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaAzureReplicationDetailsResponse) *string { return v.VmProtectionStateDescription }).(pulumi.StringPtrOutput)
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





type HyperVReplicaBasePolicyDetailsResponseInput interface {
	pulumi.Input

	ToHyperVReplicaBasePolicyDetailsResponseOutput() HyperVReplicaBasePolicyDetailsResponseOutput
	ToHyperVReplicaBasePolicyDetailsResponseOutputWithContext(context.Context) HyperVReplicaBasePolicyDetailsResponseOutput
}

type HyperVReplicaBasePolicyDetailsResponseArgs struct {
	AllowedAuthenticationType                     pulumi.IntPtrInput    `pulumi:"allowedAuthenticationType"`
	ApplicationConsistentSnapshotFrequencyInHours pulumi.IntPtrInput    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Compression                                   pulumi.StringPtrInput `pulumi:"compression"`
	InitialReplicationMethod                      pulumi.StringPtrInput `pulumi:"initialReplicationMethod"`
	InstanceType                                  pulumi.StringInput    `pulumi:"instanceType"`
	OfflineReplicationExportPath                  pulumi.StringPtrInput `pulumi:"offlineReplicationExportPath"`
	OfflineReplicationImportPath                  pulumi.StringPtrInput `pulumi:"offlineReplicationImportPath"`
	OnlineReplicationStartTime                    pulumi.StringPtrInput `pulumi:"onlineReplicationStartTime"`
	RecoveryPoints                                pulumi.IntPtrInput    `pulumi:"recoveryPoints"`
	ReplicaDeletionOption                         pulumi.StringPtrInput `pulumi:"replicaDeletionOption"`
	ReplicationPort                               pulumi.IntPtrInput    `pulumi:"replicationPort"`
}

func (HyperVReplicaBasePolicyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaBasePolicyDetailsResponse)(nil)).Elem()
}

func (i HyperVReplicaBasePolicyDetailsResponseArgs) ToHyperVReplicaBasePolicyDetailsResponseOutput() HyperVReplicaBasePolicyDetailsResponseOutput {
	return i.ToHyperVReplicaBasePolicyDetailsResponseOutputWithContext(context.Background())
}

func (i HyperVReplicaBasePolicyDetailsResponseArgs) ToHyperVReplicaBasePolicyDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaBasePolicyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaBasePolicyDetailsResponseOutput)
}

type HyperVReplicaBasePolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (HyperVReplicaBasePolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaBasePolicyDetailsResponse)(nil)).Elem()
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) ToHyperVReplicaBasePolicyDetailsResponseOutput() HyperVReplicaBasePolicyDetailsResponseOutput {
	return o
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) ToHyperVReplicaBasePolicyDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaBasePolicyDetailsResponseOutput {
	return o
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) AllowedAuthenticationType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) *int { return v.AllowedAuthenticationType }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) ApplicationConsistentSnapshotFrequencyInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) *int {
		return v.ApplicationConsistentSnapshotFrequencyInHours
	}).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) *string { return v.Compression }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) InitialReplicationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) *string { return v.InitialReplicationMethod }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) OfflineReplicationExportPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) *string { return v.OfflineReplicationExportPath }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) OfflineReplicationImportPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) *string { return v.OfflineReplicationImportPath }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) OnlineReplicationStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) *string { return v.OnlineReplicationStartTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) RecoveryPoints() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) *int { return v.RecoveryPoints }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) ReplicaDeletionOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) *string { return v.ReplicaDeletionOption }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBasePolicyDetailsResponseOutput) ReplicationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBasePolicyDetailsResponse) *int { return v.ReplicationPort }).(pulumi.IntPtrOutput)
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





type HyperVReplicaBaseReplicationDetailsResponseInput interface {
	pulumi.Input

	ToHyperVReplicaBaseReplicationDetailsResponseOutput() HyperVReplicaBaseReplicationDetailsResponseOutput
	ToHyperVReplicaBaseReplicationDetailsResponseOutputWithContext(context.Context) HyperVReplicaBaseReplicationDetailsResponseOutput
}

type HyperVReplicaBaseReplicationDetailsResponseArgs struct {
	InitialReplicationDetails    InitialReplicationDetailsResponsePtrInput `pulumi:"initialReplicationDetails"`
	InstanceType                 pulumi.StringInput                        `pulumi:"instanceType"`
	LastReplicatedTime           pulumi.StringPtrInput                     `pulumi:"lastReplicatedTime"`
	VMDiskDetails                DiskDetailsResponseArrayInput             `pulumi:"vMDiskDetails"`
	VmId                         pulumi.StringPtrInput                     `pulumi:"vmId"`
	VmNics                       VMNicDetailsResponseArrayInput            `pulumi:"vmNics"`
	VmProtectionState            pulumi.StringPtrInput                     `pulumi:"vmProtectionState"`
	VmProtectionStateDescription pulumi.StringPtrInput                     `pulumi:"vmProtectionStateDescription"`
}

func (HyperVReplicaBaseReplicationDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaBaseReplicationDetailsResponse)(nil)).Elem()
}

func (i HyperVReplicaBaseReplicationDetailsResponseArgs) ToHyperVReplicaBaseReplicationDetailsResponseOutput() HyperVReplicaBaseReplicationDetailsResponseOutput {
	return i.ToHyperVReplicaBaseReplicationDetailsResponseOutputWithContext(context.Background())
}

func (i HyperVReplicaBaseReplicationDetailsResponseArgs) ToHyperVReplicaBaseReplicationDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaBaseReplicationDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaBaseReplicationDetailsResponseOutput)
}

type HyperVReplicaBaseReplicationDetailsResponseOutput struct{ *pulumi.OutputState }

func (HyperVReplicaBaseReplicationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaBaseReplicationDetailsResponse)(nil)).Elem()
}

func (o HyperVReplicaBaseReplicationDetailsResponseOutput) ToHyperVReplicaBaseReplicationDetailsResponseOutput() HyperVReplicaBaseReplicationDetailsResponseOutput {
	return o
}

func (o HyperVReplicaBaseReplicationDetailsResponseOutput) ToHyperVReplicaBaseReplicationDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaBaseReplicationDetailsResponseOutput {
	return o
}

func (o HyperVReplicaBaseReplicationDetailsResponseOutput) InitialReplicationDetails() InitialReplicationDetailsResponsePtrOutput {
	return o.ApplyT(func(v HyperVReplicaBaseReplicationDetailsResponse) *InitialReplicationDetailsResponse {
		return v.InitialReplicationDetails
	}).(InitialReplicationDetailsResponsePtrOutput)
}

func (o HyperVReplicaBaseReplicationDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v HyperVReplicaBaseReplicationDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o HyperVReplicaBaseReplicationDetailsResponseOutput) LastReplicatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBaseReplicationDetailsResponse) *string { return v.LastReplicatedTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBaseReplicationDetailsResponseOutput) VMDiskDetails() DiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v HyperVReplicaBaseReplicationDetailsResponse) []DiskDetailsResponse { return v.VMDiskDetails }).(DiskDetailsResponseArrayOutput)
}

func (o HyperVReplicaBaseReplicationDetailsResponseOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBaseReplicationDetailsResponse) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBaseReplicationDetailsResponseOutput) VmNics() VMNicDetailsResponseArrayOutput {
	return o.ApplyT(func(v HyperVReplicaBaseReplicationDetailsResponse) []VMNicDetailsResponse { return v.VmNics }).(VMNicDetailsResponseArrayOutput)
}

func (o HyperVReplicaBaseReplicationDetailsResponseOutput) VmProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBaseReplicationDetailsResponse) *string { return v.VmProtectionState }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBaseReplicationDetailsResponseOutput) VmProtectionStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBaseReplicationDetailsResponse) *string { return v.VmProtectionStateDescription }).(pulumi.StringPtrOutput)
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





type HyperVReplicaBluePolicyDetailsResponseInput interface {
	pulumi.Input

	ToHyperVReplicaBluePolicyDetailsResponseOutput() HyperVReplicaBluePolicyDetailsResponseOutput
	ToHyperVReplicaBluePolicyDetailsResponseOutputWithContext(context.Context) HyperVReplicaBluePolicyDetailsResponseOutput
}

type HyperVReplicaBluePolicyDetailsResponseArgs struct {
	AllowedAuthenticationType                     pulumi.IntPtrInput    `pulumi:"allowedAuthenticationType"`
	ApplicationConsistentSnapshotFrequencyInHours pulumi.IntPtrInput    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Compression                                   pulumi.StringPtrInput `pulumi:"compression"`
	InitialReplicationMethod                      pulumi.StringPtrInput `pulumi:"initialReplicationMethod"`
	InstanceType                                  pulumi.StringInput    `pulumi:"instanceType"`
	OfflineReplicationExportPath                  pulumi.StringPtrInput `pulumi:"offlineReplicationExportPath"`
	OfflineReplicationImportPath                  pulumi.StringPtrInput `pulumi:"offlineReplicationImportPath"`
	OnlineReplicationStartTime                    pulumi.StringPtrInput `pulumi:"onlineReplicationStartTime"`
	RecoveryPoints                                pulumi.IntPtrInput    `pulumi:"recoveryPoints"`
	ReplicaDeletionOption                         pulumi.StringPtrInput `pulumi:"replicaDeletionOption"`
	ReplicationFrequencyInSeconds                 pulumi.IntPtrInput    `pulumi:"replicationFrequencyInSeconds"`
	ReplicationPort                               pulumi.IntPtrInput    `pulumi:"replicationPort"`
}

func (HyperVReplicaBluePolicyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaBluePolicyDetailsResponse)(nil)).Elem()
}

func (i HyperVReplicaBluePolicyDetailsResponseArgs) ToHyperVReplicaBluePolicyDetailsResponseOutput() HyperVReplicaBluePolicyDetailsResponseOutput {
	return i.ToHyperVReplicaBluePolicyDetailsResponseOutputWithContext(context.Background())
}

func (i HyperVReplicaBluePolicyDetailsResponseArgs) ToHyperVReplicaBluePolicyDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaBluePolicyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaBluePolicyDetailsResponseOutput)
}

type HyperVReplicaBluePolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (HyperVReplicaBluePolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaBluePolicyDetailsResponse)(nil)).Elem()
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) ToHyperVReplicaBluePolicyDetailsResponseOutput() HyperVReplicaBluePolicyDetailsResponseOutput {
	return o
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) ToHyperVReplicaBluePolicyDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaBluePolicyDetailsResponseOutput {
	return o
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) AllowedAuthenticationType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *int { return v.AllowedAuthenticationType }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) ApplicationConsistentSnapshotFrequencyInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *int {
		return v.ApplicationConsistentSnapshotFrequencyInHours
	}).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *string { return v.Compression }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) InitialReplicationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *string { return v.InitialReplicationMethod }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) OfflineReplicationExportPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *string { return v.OfflineReplicationExportPath }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) OfflineReplicationImportPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *string { return v.OfflineReplicationImportPath }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) OnlineReplicationStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *string { return v.OnlineReplicationStartTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) RecoveryPoints() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *int { return v.RecoveryPoints }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) ReplicaDeletionOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *string { return v.ReplicaDeletionOption }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) ReplicationFrequencyInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *int { return v.ReplicationFrequencyInSeconds }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBluePolicyDetailsResponseOutput) ReplicationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyDetailsResponse) *int { return v.ReplicationPort }).(pulumi.IntPtrOutput)
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





type HyperVReplicaBluePolicyInputInput interface {
	pulumi.Input

	ToHyperVReplicaBluePolicyInputOutput() HyperVReplicaBluePolicyInputOutput
	ToHyperVReplicaBluePolicyInputOutputWithContext(context.Context) HyperVReplicaBluePolicyInputOutput
}

type HyperVReplicaBluePolicyInputArgs struct {
	AllowedAuthenticationType                     pulumi.IntPtrInput    `pulumi:"allowedAuthenticationType"`
	ApplicationConsistentSnapshotFrequencyInHours pulumi.IntPtrInput    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Compression                                   pulumi.StringPtrInput `pulumi:"compression"`
	InitialReplicationMethod                      pulumi.StringPtrInput `pulumi:"initialReplicationMethod"`
	InstanceType                                  pulumi.StringPtrInput `pulumi:"instanceType"`
	OfflineReplicationExportPath                  pulumi.StringPtrInput `pulumi:"offlineReplicationExportPath"`
	OfflineReplicationImportPath                  pulumi.StringPtrInput `pulumi:"offlineReplicationImportPath"`
	OnlineReplicationStartTime                    pulumi.StringPtrInput `pulumi:"onlineReplicationStartTime"`
	RecoveryPoints                                pulumi.IntPtrInput    `pulumi:"recoveryPoints"`
	ReplicaDeletion                               pulumi.StringPtrInput `pulumi:"replicaDeletion"`
	ReplicationFrequencyInSeconds                 pulumi.IntPtrInput    `pulumi:"replicationFrequencyInSeconds"`
	ReplicationPort                               pulumi.IntPtrInput    `pulumi:"replicationPort"`
}

func (HyperVReplicaBluePolicyInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaBluePolicyInput)(nil)).Elem()
}

func (i HyperVReplicaBluePolicyInputArgs) ToHyperVReplicaBluePolicyInputOutput() HyperVReplicaBluePolicyInputOutput {
	return i.ToHyperVReplicaBluePolicyInputOutputWithContext(context.Background())
}

func (i HyperVReplicaBluePolicyInputArgs) ToHyperVReplicaBluePolicyInputOutputWithContext(ctx context.Context) HyperVReplicaBluePolicyInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaBluePolicyInputOutput)
}

type HyperVReplicaBluePolicyInputOutput struct{ *pulumi.OutputState }

func (HyperVReplicaBluePolicyInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaBluePolicyInput)(nil)).Elem()
}

func (o HyperVReplicaBluePolicyInputOutput) ToHyperVReplicaBluePolicyInputOutput() HyperVReplicaBluePolicyInputOutput {
	return o
}

func (o HyperVReplicaBluePolicyInputOutput) ToHyperVReplicaBluePolicyInputOutputWithContext(ctx context.Context) HyperVReplicaBluePolicyInputOutput {
	return o
}

func (o HyperVReplicaBluePolicyInputOutput) AllowedAuthenticationType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *int { return v.AllowedAuthenticationType }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) ApplicationConsistentSnapshotFrequencyInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *int { return v.ApplicationConsistentSnapshotFrequencyInHours }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *string { return v.Compression }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) InitialReplicationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *string { return v.InitialReplicationMethod }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) OfflineReplicationExportPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *string { return v.OfflineReplicationExportPath }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) OfflineReplicationImportPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *string { return v.OfflineReplicationImportPath }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) OnlineReplicationStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *string { return v.OnlineReplicationStartTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) RecoveryPoints() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *int { return v.RecoveryPoints }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) ReplicaDeletion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *string { return v.ReplicaDeletion }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) ReplicationFrequencyInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *int { return v.ReplicationFrequencyInSeconds }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaBluePolicyInputOutput) ReplicationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBluePolicyInput) *int { return v.ReplicationPort }).(pulumi.IntPtrOutput)
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





type HyperVReplicaBlueReplicationDetailsResponseInput interface {
	pulumi.Input

	ToHyperVReplicaBlueReplicationDetailsResponseOutput() HyperVReplicaBlueReplicationDetailsResponseOutput
	ToHyperVReplicaBlueReplicationDetailsResponseOutputWithContext(context.Context) HyperVReplicaBlueReplicationDetailsResponseOutput
}

type HyperVReplicaBlueReplicationDetailsResponseArgs struct {
	InitialReplicationDetails    InitialReplicationDetailsResponsePtrInput `pulumi:"initialReplicationDetails"`
	InstanceType                 pulumi.StringInput                        `pulumi:"instanceType"`
	LastReplicatedTime           pulumi.StringPtrInput                     `pulumi:"lastReplicatedTime"`
	VMDiskDetails                DiskDetailsResponseArrayInput             `pulumi:"vMDiskDetails"`
	VmId                         pulumi.StringPtrInput                     `pulumi:"vmId"`
	VmNics                       VMNicDetailsResponseArrayInput            `pulumi:"vmNics"`
	VmProtectionState            pulumi.StringPtrInput                     `pulumi:"vmProtectionState"`
	VmProtectionStateDescription pulumi.StringPtrInput                     `pulumi:"vmProtectionStateDescription"`
}

func (HyperVReplicaBlueReplicationDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaBlueReplicationDetailsResponse)(nil)).Elem()
}

func (i HyperVReplicaBlueReplicationDetailsResponseArgs) ToHyperVReplicaBlueReplicationDetailsResponseOutput() HyperVReplicaBlueReplicationDetailsResponseOutput {
	return i.ToHyperVReplicaBlueReplicationDetailsResponseOutputWithContext(context.Background())
}

func (i HyperVReplicaBlueReplicationDetailsResponseArgs) ToHyperVReplicaBlueReplicationDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaBlueReplicationDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaBlueReplicationDetailsResponseOutput)
}

type HyperVReplicaBlueReplicationDetailsResponseOutput struct{ *pulumi.OutputState }

func (HyperVReplicaBlueReplicationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaBlueReplicationDetailsResponse)(nil)).Elem()
}

func (o HyperVReplicaBlueReplicationDetailsResponseOutput) ToHyperVReplicaBlueReplicationDetailsResponseOutput() HyperVReplicaBlueReplicationDetailsResponseOutput {
	return o
}

func (o HyperVReplicaBlueReplicationDetailsResponseOutput) ToHyperVReplicaBlueReplicationDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaBlueReplicationDetailsResponseOutput {
	return o
}

func (o HyperVReplicaBlueReplicationDetailsResponseOutput) InitialReplicationDetails() InitialReplicationDetailsResponsePtrOutput {
	return o.ApplyT(func(v HyperVReplicaBlueReplicationDetailsResponse) *InitialReplicationDetailsResponse {
		return v.InitialReplicationDetails
	}).(InitialReplicationDetailsResponsePtrOutput)
}

func (o HyperVReplicaBlueReplicationDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v HyperVReplicaBlueReplicationDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o HyperVReplicaBlueReplicationDetailsResponseOutput) LastReplicatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBlueReplicationDetailsResponse) *string { return v.LastReplicatedTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBlueReplicationDetailsResponseOutput) VMDiskDetails() DiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v HyperVReplicaBlueReplicationDetailsResponse) []DiskDetailsResponse { return v.VMDiskDetails }).(DiskDetailsResponseArrayOutput)
}

func (o HyperVReplicaBlueReplicationDetailsResponseOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBlueReplicationDetailsResponse) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBlueReplicationDetailsResponseOutput) VmNics() VMNicDetailsResponseArrayOutput {
	return o.ApplyT(func(v HyperVReplicaBlueReplicationDetailsResponse) []VMNicDetailsResponse { return v.VmNics }).(VMNicDetailsResponseArrayOutput)
}

func (o HyperVReplicaBlueReplicationDetailsResponseOutput) VmProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBlueReplicationDetailsResponse) *string { return v.VmProtectionState }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaBlueReplicationDetailsResponseOutput) VmProtectionStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaBlueReplicationDetailsResponse) *string { return v.VmProtectionStateDescription }).(pulumi.StringPtrOutput)
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





type HyperVReplicaPolicyDetailsResponseInput interface {
	pulumi.Input

	ToHyperVReplicaPolicyDetailsResponseOutput() HyperVReplicaPolicyDetailsResponseOutput
	ToHyperVReplicaPolicyDetailsResponseOutputWithContext(context.Context) HyperVReplicaPolicyDetailsResponseOutput
}

type HyperVReplicaPolicyDetailsResponseArgs struct {
	AllowedAuthenticationType                     pulumi.IntPtrInput    `pulumi:"allowedAuthenticationType"`
	ApplicationConsistentSnapshotFrequencyInHours pulumi.IntPtrInput    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Compression                                   pulumi.StringPtrInput `pulumi:"compression"`
	InitialReplicationMethod                      pulumi.StringPtrInput `pulumi:"initialReplicationMethod"`
	InstanceType                                  pulumi.StringInput    `pulumi:"instanceType"`
	OfflineReplicationExportPath                  pulumi.StringPtrInput `pulumi:"offlineReplicationExportPath"`
	OfflineReplicationImportPath                  pulumi.StringPtrInput `pulumi:"offlineReplicationImportPath"`
	OnlineReplicationStartTime                    pulumi.StringPtrInput `pulumi:"onlineReplicationStartTime"`
	RecoveryPoints                                pulumi.IntPtrInput    `pulumi:"recoveryPoints"`
	ReplicaDeletionOption                         pulumi.StringPtrInput `pulumi:"replicaDeletionOption"`
	ReplicationPort                               pulumi.IntPtrInput    `pulumi:"replicationPort"`
}

func (HyperVReplicaPolicyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaPolicyDetailsResponse)(nil)).Elem()
}

func (i HyperVReplicaPolicyDetailsResponseArgs) ToHyperVReplicaPolicyDetailsResponseOutput() HyperVReplicaPolicyDetailsResponseOutput {
	return i.ToHyperVReplicaPolicyDetailsResponseOutputWithContext(context.Background())
}

func (i HyperVReplicaPolicyDetailsResponseArgs) ToHyperVReplicaPolicyDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaPolicyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaPolicyDetailsResponseOutput)
}

type HyperVReplicaPolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (HyperVReplicaPolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaPolicyDetailsResponse)(nil)).Elem()
}

func (o HyperVReplicaPolicyDetailsResponseOutput) ToHyperVReplicaPolicyDetailsResponseOutput() HyperVReplicaPolicyDetailsResponseOutput {
	return o
}

func (o HyperVReplicaPolicyDetailsResponseOutput) ToHyperVReplicaPolicyDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaPolicyDetailsResponseOutput {
	return o
}

func (o HyperVReplicaPolicyDetailsResponseOutput) AllowedAuthenticationType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) *int { return v.AllowedAuthenticationType }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaPolicyDetailsResponseOutput) ApplicationConsistentSnapshotFrequencyInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) *int {
		return v.ApplicationConsistentSnapshotFrequencyInHours
	}).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaPolicyDetailsResponseOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) *string { return v.Compression }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyDetailsResponseOutput) InitialReplicationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) *string { return v.InitialReplicationMethod }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o HyperVReplicaPolicyDetailsResponseOutput) OfflineReplicationExportPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) *string { return v.OfflineReplicationExportPath }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyDetailsResponseOutput) OfflineReplicationImportPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) *string { return v.OfflineReplicationImportPath }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyDetailsResponseOutput) OnlineReplicationStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) *string { return v.OnlineReplicationStartTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyDetailsResponseOutput) RecoveryPoints() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) *int { return v.RecoveryPoints }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaPolicyDetailsResponseOutput) ReplicaDeletionOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) *string { return v.ReplicaDeletionOption }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyDetailsResponseOutput) ReplicationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyDetailsResponse) *int { return v.ReplicationPort }).(pulumi.IntPtrOutput)
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





type HyperVReplicaPolicyInputInput interface {
	pulumi.Input

	ToHyperVReplicaPolicyInputOutput() HyperVReplicaPolicyInputOutput
	ToHyperVReplicaPolicyInputOutputWithContext(context.Context) HyperVReplicaPolicyInputOutput
}

type HyperVReplicaPolicyInputArgs struct {
	AllowedAuthenticationType                     pulumi.IntPtrInput    `pulumi:"allowedAuthenticationType"`
	ApplicationConsistentSnapshotFrequencyInHours pulumi.IntPtrInput    `pulumi:"applicationConsistentSnapshotFrequencyInHours"`
	Compression                                   pulumi.StringPtrInput `pulumi:"compression"`
	InitialReplicationMethod                      pulumi.StringPtrInput `pulumi:"initialReplicationMethod"`
	InstanceType                                  pulumi.StringPtrInput `pulumi:"instanceType"`
	OfflineReplicationExportPath                  pulumi.StringPtrInput `pulumi:"offlineReplicationExportPath"`
	OfflineReplicationImportPath                  pulumi.StringPtrInput `pulumi:"offlineReplicationImportPath"`
	OnlineReplicationStartTime                    pulumi.StringPtrInput `pulumi:"onlineReplicationStartTime"`
	RecoveryPoints                                pulumi.IntPtrInput    `pulumi:"recoveryPoints"`
	ReplicaDeletion                               pulumi.StringPtrInput `pulumi:"replicaDeletion"`
	ReplicationPort                               pulumi.IntPtrInput    `pulumi:"replicationPort"`
}

func (HyperVReplicaPolicyInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaPolicyInput)(nil)).Elem()
}

func (i HyperVReplicaPolicyInputArgs) ToHyperVReplicaPolicyInputOutput() HyperVReplicaPolicyInputOutput {
	return i.ToHyperVReplicaPolicyInputOutputWithContext(context.Background())
}

func (i HyperVReplicaPolicyInputArgs) ToHyperVReplicaPolicyInputOutputWithContext(ctx context.Context) HyperVReplicaPolicyInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaPolicyInputOutput)
}

type HyperVReplicaPolicyInputOutput struct{ *pulumi.OutputState }

func (HyperVReplicaPolicyInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaPolicyInput)(nil)).Elem()
}

func (o HyperVReplicaPolicyInputOutput) ToHyperVReplicaPolicyInputOutput() HyperVReplicaPolicyInputOutput {
	return o
}

func (o HyperVReplicaPolicyInputOutput) ToHyperVReplicaPolicyInputOutputWithContext(ctx context.Context) HyperVReplicaPolicyInputOutput {
	return o
}

func (o HyperVReplicaPolicyInputOutput) AllowedAuthenticationType() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *int { return v.AllowedAuthenticationType }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaPolicyInputOutput) ApplicationConsistentSnapshotFrequencyInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *int { return v.ApplicationConsistentSnapshotFrequencyInHours }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaPolicyInputOutput) Compression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *string { return v.Compression }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyInputOutput) InitialReplicationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *string { return v.InitialReplicationMethod }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyInputOutput) OfflineReplicationExportPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *string { return v.OfflineReplicationExportPath }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyInputOutput) OfflineReplicationImportPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *string { return v.OfflineReplicationImportPath }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyInputOutput) OnlineReplicationStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *string { return v.OnlineReplicationStartTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyInputOutput) RecoveryPoints() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *int { return v.RecoveryPoints }).(pulumi.IntPtrOutput)
}

func (o HyperVReplicaPolicyInputOutput) ReplicaDeletion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *string { return v.ReplicaDeletion }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaPolicyInputOutput) ReplicationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HyperVReplicaPolicyInput) *int { return v.ReplicationPort }).(pulumi.IntPtrOutput)
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





type HyperVReplicaReplicationDetailsResponseInput interface {
	pulumi.Input

	ToHyperVReplicaReplicationDetailsResponseOutput() HyperVReplicaReplicationDetailsResponseOutput
	ToHyperVReplicaReplicationDetailsResponseOutputWithContext(context.Context) HyperVReplicaReplicationDetailsResponseOutput
}

type HyperVReplicaReplicationDetailsResponseArgs struct {
	InitialReplicationDetails    InitialReplicationDetailsResponsePtrInput `pulumi:"initialReplicationDetails"`
	InstanceType                 pulumi.StringInput                        `pulumi:"instanceType"`
	LastReplicatedTime           pulumi.StringPtrInput                     `pulumi:"lastReplicatedTime"`
	VMDiskDetails                DiskDetailsResponseArrayInput             `pulumi:"vMDiskDetails"`
	VmId                         pulumi.StringPtrInput                     `pulumi:"vmId"`
	VmNics                       VMNicDetailsResponseArrayInput            `pulumi:"vmNics"`
	VmProtectionState            pulumi.StringPtrInput                     `pulumi:"vmProtectionState"`
	VmProtectionStateDescription pulumi.StringPtrInput                     `pulumi:"vmProtectionStateDescription"`
}

func (HyperVReplicaReplicationDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaReplicationDetailsResponse)(nil)).Elem()
}

func (i HyperVReplicaReplicationDetailsResponseArgs) ToHyperVReplicaReplicationDetailsResponseOutput() HyperVReplicaReplicationDetailsResponseOutput {
	return i.ToHyperVReplicaReplicationDetailsResponseOutputWithContext(context.Background())
}

func (i HyperVReplicaReplicationDetailsResponseArgs) ToHyperVReplicaReplicationDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaReplicationDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVReplicaReplicationDetailsResponseOutput)
}

type HyperVReplicaReplicationDetailsResponseOutput struct{ *pulumi.OutputState }

func (HyperVReplicaReplicationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVReplicaReplicationDetailsResponse)(nil)).Elem()
}

func (o HyperVReplicaReplicationDetailsResponseOutput) ToHyperVReplicaReplicationDetailsResponseOutput() HyperVReplicaReplicationDetailsResponseOutput {
	return o
}

func (o HyperVReplicaReplicationDetailsResponseOutput) ToHyperVReplicaReplicationDetailsResponseOutputWithContext(ctx context.Context) HyperVReplicaReplicationDetailsResponseOutput {
	return o
}

func (o HyperVReplicaReplicationDetailsResponseOutput) InitialReplicationDetails() InitialReplicationDetailsResponsePtrOutput {
	return o.ApplyT(func(v HyperVReplicaReplicationDetailsResponse) *InitialReplicationDetailsResponse {
		return v.InitialReplicationDetails
	}).(InitialReplicationDetailsResponsePtrOutput)
}

func (o HyperVReplicaReplicationDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v HyperVReplicaReplicationDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o HyperVReplicaReplicationDetailsResponseOutput) LastReplicatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaReplicationDetailsResponse) *string { return v.LastReplicatedTime }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaReplicationDetailsResponseOutput) VMDiskDetails() DiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v HyperVReplicaReplicationDetailsResponse) []DiskDetailsResponse { return v.VMDiskDetails }).(DiskDetailsResponseArrayOutput)
}

func (o HyperVReplicaReplicationDetailsResponseOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaReplicationDetailsResponse) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaReplicationDetailsResponseOutput) VmNics() VMNicDetailsResponseArrayOutput {
	return o.ApplyT(func(v HyperVReplicaReplicationDetailsResponse) []VMNicDetailsResponse { return v.VmNics }).(VMNicDetailsResponseArrayOutput)
}

func (o HyperVReplicaReplicationDetailsResponseOutput) VmProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaReplicationDetailsResponse) *string { return v.VmProtectionState }).(pulumi.StringPtrOutput)
}

func (o HyperVReplicaReplicationDetailsResponseOutput) VmProtectionStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HyperVReplicaReplicationDetailsResponse) *string { return v.VmProtectionStateDescription }).(pulumi.StringPtrOutput)
}

type HyperVSiteDetailsResponse struct {
	InstanceType string `pulumi:"instanceType"`
}





type HyperVSiteDetailsResponseInput interface {
	pulumi.Input

	ToHyperVSiteDetailsResponseOutput() HyperVSiteDetailsResponseOutput
	ToHyperVSiteDetailsResponseOutputWithContext(context.Context) HyperVSiteDetailsResponseOutput
}

type HyperVSiteDetailsResponseArgs struct {
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
}

func (HyperVSiteDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVSiteDetailsResponse)(nil)).Elem()
}

func (i HyperVSiteDetailsResponseArgs) ToHyperVSiteDetailsResponseOutput() HyperVSiteDetailsResponseOutput {
	return i.ToHyperVSiteDetailsResponseOutputWithContext(context.Background())
}

func (i HyperVSiteDetailsResponseArgs) ToHyperVSiteDetailsResponseOutputWithContext(ctx context.Context) HyperVSiteDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HyperVSiteDetailsResponseOutput)
}

type HyperVSiteDetailsResponseOutput struct{ *pulumi.OutputState }

func (HyperVSiteDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HyperVSiteDetailsResponse)(nil)).Elem()
}

func (o HyperVSiteDetailsResponseOutput) ToHyperVSiteDetailsResponseOutput() HyperVSiteDetailsResponseOutput {
	return o
}

func (o HyperVSiteDetailsResponseOutput) ToHyperVSiteDetailsResponseOutputWithContext(ctx context.Context) HyperVSiteDetailsResponseOutput {
	return o
}

func (o HyperVSiteDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v HyperVSiteDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

type IdentityProviderDetailsResponse struct {
	AadAuthority  *string `pulumi:"aadAuthority"`
	ApplicationId *string `pulumi:"applicationId"`
	Audience      *string `pulumi:"audience"`
	ObjectId      *string `pulumi:"objectId"`
	TenantId      *string `pulumi:"tenantId"`
}





type IdentityProviderDetailsResponseInput interface {
	pulumi.Input

	ToIdentityProviderDetailsResponseOutput() IdentityProviderDetailsResponseOutput
	ToIdentityProviderDetailsResponseOutputWithContext(context.Context) IdentityProviderDetailsResponseOutput
}

type IdentityProviderDetailsResponseArgs struct {
	AadAuthority  pulumi.StringPtrInput `pulumi:"aadAuthority"`
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	Audience      pulumi.StringPtrInput `pulumi:"audience"`
	ObjectId      pulumi.StringPtrInput `pulumi:"objectId"`
	TenantId      pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (IdentityProviderDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProviderDetailsResponse)(nil)).Elem()
}

func (i IdentityProviderDetailsResponseArgs) ToIdentityProviderDetailsResponseOutput() IdentityProviderDetailsResponseOutput {
	return i.ToIdentityProviderDetailsResponseOutputWithContext(context.Background())
}

func (i IdentityProviderDetailsResponseArgs) ToIdentityProviderDetailsResponseOutputWithContext(ctx context.Context) IdentityProviderDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderDetailsResponseOutput)
}

func (i IdentityProviderDetailsResponseArgs) ToIdentityProviderDetailsResponsePtrOutput() IdentityProviderDetailsResponsePtrOutput {
	return i.ToIdentityProviderDetailsResponsePtrOutputWithContext(context.Background())
}

func (i IdentityProviderDetailsResponseArgs) ToIdentityProviderDetailsResponsePtrOutputWithContext(ctx context.Context) IdentityProviderDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderDetailsResponseOutput).ToIdentityProviderDetailsResponsePtrOutputWithContext(ctx)
}









type IdentityProviderDetailsResponsePtrInput interface {
	pulumi.Input

	ToIdentityProviderDetailsResponsePtrOutput() IdentityProviderDetailsResponsePtrOutput
	ToIdentityProviderDetailsResponsePtrOutputWithContext(context.Context) IdentityProviderDetailsResponsePtrOutput
}

type identityProviderDetailsResponsePtrType IdentityProviderDetailsResponseArgs

func IdentityProviderDetailsResponsePtr(v *IdentityProviderDetailsResponseArgs) IdentityProviderDetailsResponsePtrInput {
	return (*identityProviderDetailsResponsePtrType)(v)
}

func (*identityProviderDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProviderDetailsResponse)(nil)).Elem()
}

func (i *identityProviderDetailsResponsePtrType) ToIdentityProviderDetailsResponsePtrOutput() IdentityProviderDetailsResponsePtrOutput {
	return i.ToIdentityProviderDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *identityProviderDetailsResponsePtrType) ToIdentityProviderDetailsResponsePtrOutputWithContext(ctx context.Context) IdentityProviderDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityProviderDetailsResponsePtrOutput)
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

func (o IdentityProviderDetailsResponseOutput) ToIdentityProviderDetailsResponsePtrOutput() IdentityProviderDetailsResponsePtrOutput {
	return o.ToIdentityProviderDetailsResponsePtrOutputWithContext(context.Background())
}

func (o IdentityProviderDetailsResponseOutput) ToIdentityProviderDetailsResponsePtrOutputWithContext(ctx context.Context) IdentityProviderDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProviderDetailsResponse) *IdentityProviderDetailsResponse {
		return &v
	}).(IdentityProviderDetailsResponsePtrOutput)
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





type InMageAgentDetailsResponseInput interface {
	pulumi.Input

	ToInMageAgentDetailsResponseOutput() InMageAgentDetailsResponseOutput
	ToInMageAgentDetailsResponseOutputWithContext(context.Context) InMageAgentDetailsResponseOutput
}

type InMageAgentDetailsResponseArgs struct {
	AgentExpiryDate        pulumi.StringPtrInput `pulumi:"agentExpiryDate"`
	AgentUpdateStatus      pulumi.StringPtrInput `pulumi:"agentUpdateStatus"`
	AgentVersion           pulumi.StringPtrInput `pulumi:"agentVersion"`
	PostUpdateRebootStatus pulumi.StringPtrInput `pulumi:"postUpdateRebootStatus"`
}

func (InMageAgentDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAgentDetailsResponse)(nil)).Elem()
}

func (i InMageAgentDetailsResponseArgs) ToInMageAgentDetailsResponseOutput() InMageAgentDetailsResponseOutput {
	return i.ToInMageAgentDetailsResponseOutputWithContext(context.Background())
}

func (i InMageAgentDetailsResponseArgs) ToInMageAgentDetailsResponseOutputWithContext(ctx context.Context) InMageAgentDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageAgentDetailsResponseOutput)
}

func (i InMageAgentDetailsResponseArgs) ToInMageAgentDetailsResponsePtrOutput() InMageAgentDetailsResponsePtrOutput {
	return i.ToInMageAgentDetailsResponsePtrOutputWithContext(context.Background())
}

func (i InMageAgentDetailsResponseArgs) ToInMageAgentDetailsResponsePtrOutputWithContext(ctx context.Context) InMageAgentDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageAgentDetailsResponseOutput).ToInMageAgentDetailsResponsePtrOutputWithContext(ctx)
}









type InMageAgentDetailsResponsePtrInput interface {
	pulumi.Input

	ToInMageAgentDetailsResponsePtrOutput() InMageAgentDetailsResponsePtrOutput
	ToInMageAgentDetailsResponsePtrOutputWithContext(context.Context) InMageAgentDetailsResponsePtrOutput
}

type inMageAgentDetailsResponsePtrType InMageAgentDetailsResponseArgs

func InMageAgentDetailsResponsePtr(v *InMageAgentDetailsResponseArgs) InMageAgentDetailsResponsePtrInput {
	return (*inMageAgentDetailsResponsePtrType)(v)
}

func (*inMageAgentDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InMageAgentDetailsResponse)(nil)).Elem()
}

func (i *inMageAgentDetailsResponsePtrType) ToInMageAgentDetailsResponsePtrOutput() InMageAgentDetailsResponsePtrOutput {
	return i.ToInMageAgentDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *inMageAgentDetailsResponsePtrType) ToInMageAgentDetailsResponsePtrOutputWithContext(ctx context.Context) InMageAgentDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageAgentDetailsResponsePtrOutput)
}

type InMageAgentDetailsResponseOutput struct{ *pulumi.OutputState }

func (InMageAgentDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAgentDetailsResponse)(nil)).Elem()
}

func (o InMageAgentDetailsResponseOutput) ToInMageAgentDetailsResponseOutput() InMageAgentDetailsResponseOutput {
	return o
}

func (o InMageAgentDetailsResponseOutput) ToInMageAgentDetailsResponseOutputWithContext(ctx context.Context) InMageAgentDetailsResponseOutput {
	return o
}

func (o InMageAgentDetailsResponseOutput) ToInMageAgentDetailsResponsePtrOutput() InMageAgentDetailsResponsePtrOutput {
	return o.ToInMageAgentDetailsResponsePtrOutputWithContext(context.Background())
}

func (o InMageAgentDetailsResponseOutput) ToInMageAgentDetailsResponsePtrOutputWithContext(ctx context.Context) InMageAgentDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InMageAgentDetailsResponse) *InMageAgentDetailsResponse {
		return &v
	}).(InMageAgentDetailsResponsePtrOutput)
}

func (o InMageAgentDetailsResponseOutput) AgentExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAgentDetailsResponse) *string { return v.AgentExpiryDate }).(pulumi.StringPtrOutput)
}

func (o InMageAgentDetailsResponseOutput) AgentUpdateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAgentDetailsResponse) *string { return v.AgentUpdateStatus }).(pulumi.StringPtrOutput)
}

func (o InMageAgentDetailsResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAgentDetailsResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o InMageAgentDetailsResponseOutput) PostUpdateRebootStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAgentDetailsResponse) *string { return v.PostUpdateRebootStatus }).(pulumi.StringPtrOutput)
}

type InMageAgentDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (InMageAgentDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InMageAgentDetailsResponse)(nil)).Elem()
}

func (o InMageAgentDetailsResponsePtrOutput) ToInMageAgentDetailsResponsePtrOutput() InMageAgentDetailsResponsePtrOutput {
	return o
}

func (o InMageAgentDetailsResponsePtrOutput) ToInMageAgentDetailsResponsePtrOutputWithContext(ctx context.Context) InMageAgentDetailsResponsePtrOutput {
	return o
}

func (o InMageAgentDetailsResponsePtrOutput) Elem() InMageAgentDetailsResponseOutput {
	return o.ApplyT(func(v *InMageAgentDetailsResponse) InMageAgentDetailsResponse {
		if v != nil {
			return *v
		}
		var ret InMageAgentDetailsResponse
		return ret
	}).(InMageAgentDetailsResponseOutput)
}

func (o InMageAgentDetailsResponsePtrOutput) AgentExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InMageAgentDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgentExpiryDate
	}).(pulumi.StringPtrOutput)
}

func (o InMageAgentDetailsResponsePtrOutput) AgentUpdateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InMageAgentDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgentUpdateStatus
	}).(pulumi.StringPtrOutput)
}

func (o InMageAgentDetailsResponsePtrOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InMageAgentDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AgentVersion
	}).(pulumi.StringPtrOutput)
}

func (o InMageAgentDetailsResponsePtrOutput) PostUpdateRebootStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InMageAgentDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PostUpdateRebootStatus
	}).(pulumi.StringPtrOutput)
}

type InMageAzureV2EnableProtectionInput struct {
	DisksToInclude               []string `pulumi:"disksToInclude"`
	EnableRdpOnTargetOption      *string  `pulumi:"enableRdpOnTargetOption"`
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





type InMageAzureV2EnableProtectionInputInput interface {
	pulumi.Input

	ToInMageAzureV2EnableProtectionInputOutput() InMageAzureV2EnableProtectionInputOutput
	ToInMageAzureV2EnableProtectionInputOutputWithContext(context.Context) InMageAzureV2EnableProtectionInputOutput
}

type InMageAzureV2EnableProtectionInputArgs struct {
	DisksToInclude               pulumi.StringArrayInput `pulumi:"disksToInclude"`
	EnableRdpOnTargetOption      pulumi.StringPtrInput   `pulumi:"enableRdpOnTargetOption"`
	InstanceType                 pulumi.StringPtrInput   `pulumi:"instanceType"`
	LogStorageAccountId          pulumi.StringPtrInput   `pulumi:"logStorageAccountId"`
	MasterTargetId               pulumi.StringPtrInput   `pulumi:"masterTargetId"`
	MultiVmGroupId               pulumi.StringPtrInput   `pulumi:"multiVmGroupId"`
	MultiVmGroupName             pulumi.StringPtrInput   `pulumi:"multiVmGroupName"`
	ProcessServerId              pulumi.StringPtrInput   `pulumi:"processServerId"`
	RunAsAccountId               pulumi.StringPtrInput   `pulumi:"runAsAccountId"`
	StorageAccountId             pulumi.StringInput      `pulumi:"storageAccountId"`
	TargetAzureNetworkId         pulumi.StringPtrInput   `pulumi:"targetAzureNetworkId"`
	TargetAzureSubnetId          pulumi.StringPtrInput   `pulumi:"targetAzureSubnetId"`
	TargetAzureV1ResourceGroupId pulumi.StringPtrInput   `pulumi:"targetAzureV1ResourceGroupId"`
	TargetAzureV2ResourceGroupId pulumi.StringPtrInput   `pulumi:"targetAzureV2ResourceGroupId"`
	TargetAzureVmName            pulumi.StringPtrInput   `pulumi:"targetAzureVmName"`
	UseManagedDisks              pulumi.StringPtrInput   `pulumi:"useManagedDisks"`
}

func (InMageAzureV2EnableProtectionInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAzureV2EnableProtectionInput)(nil)).Elem()
}

func (i InMageAzureV2EnableProtectionInputArgs) ToInMageAzureV2EnableProtectionInputOutput() InMageAzureV2EnableProtectionInputOutput {
	return i.ToInMageAzureV2EnableProtectionInputOutputWithContext(context.Background())
}

func (i InMageAzureV2EnableProtectionInputArgs) ToInMageAzureV2EnableProtectionInputOutputWithContext(ctx context.Context) InMageAzureV2EnableProtectionInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageAzureV2EnableProtectionInputOutput)
}

type InMageAzureV2EnableProtectionInputOutput struct{ *pulumi.OutputState }

func (InMageAzureV2EnableProtectionInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAzureV2EnableProtectionInput)(nil)).Elem()
}

func (o InMageAzureV2EnableProtectionInputOutput) ToInMageAzureV2EnableProtectionInputOutput() InMageAzureV2EnableProtectionInputOutput {
	return o
}

func (o InMageAzureV2EnableProtectionInputOutput) ToInMageAzureV2EnableProtectionInputOutputWithContext(ctx context.Context) InMageAzureV2EnableProtectionInputOutput {
	return o
}

func (o InMageAzureV2EnableProtectionInputOutput) DisksToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) []string { return v.DisksToInclude }).(pulumi.StringArrayOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) EnableRdpOnTargetOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.EnableRdpOnTargetOption }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) LogStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.LogStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) MasterTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.MasterTargetId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) MultiVmGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.MultiVmGroupId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) MultiVmGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.MultiVmGroupName }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) ProcessServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.ProcessServerId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) RunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.RunAsAccountId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) TargetAzureNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.TargetAzureNetworkId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) TargetAzureSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.TargetAzureSubnetId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) TargetAzureV1ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.TargetAzureV1ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) TargetAzureV2ResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.TargetAzureV2ResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) TargetAzureVmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.TargetAzureVmName }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2EnableProtectionInputOutput) UseManagedDisks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2EnableProtectionInput) *string { return v.UseManagedDisks }).(pulumi.StringPtrOutput)
}

type InMageAzureV2PolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string  `pulumi:"instanceType"`
	MultiVmSyncStatus                 *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int    `pulumi:"recoveryPointThresholdInMinutes"`
}





type InMageAzureV2PolicyDetailsResponseInput interface {
	pulumi.Input

	ToInMageAzureV2PolicyDetailsResponseOutput() InMageAzureV2PolicyDetailsResponseOutput
	ToInMageAzureV2PolicyDetailsResponseOutputWithContext(context.Context) InMageAzureV2PolicyDetailsResponseOutput
}

type InMageAzureV2PolicyDetailsResponseArgs struct {
	AppConsistentFrequencyInMinutes   pulumi.IntPtrInput    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes pulumi.IntPtrInput    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      pulumi.StringInput    `pulumi:"instanceType"`
	MultiVmSyncStatus                 pulumi.StringPtrInput `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              pulumi.IntPtrInput    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   pulumi.IntPtrInput    `pulumi:"recoveryPointThresholdInMinutes"`
}

func (InMageAzureV2PolicyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAzureV2PolicyDetailsResponse)(nil)).Elem()
}

func (i InMageAzureV2PolicyDetailsResponseArgs) ToInMageAzureV2PolicyDetailsResponseOutput() InMageAzureV2PolicyDetailsResponseOutput {
	return i.ToInMageAzureV2PolicyDetailsResponseOutputWithContext(context.Background())
}

func (i InMageAzureV2PolicyDetailsResponseArgs) ToInMageAzureV2PolicyDetailsResponseOutputWithContext(ctx context.Context) InMageAzureV2PolicyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageAzureV2PolicyDetailsResponseOutput)
}

type InMageAzureV2PolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (InMageAzureV2PolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAzureV2PolicyDetailsResponse)(nil)).Elem()
}

func (o InMageAzureV2PolicyDetailsResponseOutput) ToInMageAzureV2PolicyDetailsResponseOutput() InMageAzureV2PolicyDetailsResponseOutput {
	return o
}

func (o InMageAzureV2PolicyDetailsResponseOutput) ToInMageAzureV2PolicyDetailsResponseOutputWithContext(ctx context.Context) InMageAzureV2PolicyDetailsResponseOutput {
	return o
}

func (o InMageAzureV2PolicyDetailsResponseOutput) AppConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyDetailsResponse) *int { return v.AppConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o InMageAzureV2PolicyDetailsResponseOutput) CrashConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyDetailsResponse) *int { return v.CrashConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o InMageAzureV2PolicyDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o InMageAzureV2PolicyDetailsResponseOutput) MultiVmSyncStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyDetailsResponse) *string { return v.MultiVmSyncStatus }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2PolicyDetailsResponseOutput) RecoveryPointHistory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyDetailsResponse) *int { return v.RecoveryPointHistory }).(pulumi.IntPtrOutput)
}

func (o InMageAzureV2PolicyDetailsResponseOutput) RecoveryPointThresholdInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyDetailsResponse) *int { return v.RecoveryPointThresholdInMinutes }).(pulumi.IntPtrOutput)
}

type InMageAzureV2PolicyInput struct {
	AppConsistentFrequencyInMinutes   *int    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      *string `pulumi:"instanceType"`
	MultiVmSyncStatus                 string  `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   *int    `pulumi:"recoveryPointThresholdInMinutes"`
}





type InMageAzureV2PolicyInputInput interface {
	pulumi.Input

	ToInMageAzureV2PolicyInputOutput() InMageAzureV2PolicyInputOutput
	ToInMageAzureV2PolicyInputOutputWithContext(context.Context) InMageAzureV2PolicyInputOutput
}

type InMageAzureV2PolicyInputArgs struct {
	AppConsistentFrequencyInMinutes   pulumi.IntPtrInput    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes pulumi.IntPtrInput    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      pulumi.StringPtrInput `pulumi:"instanceType"`
	MultiVmSyncStatus                 pulumi.StringInput    `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              pulumi.IntPtrInput    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   pulumi.IntPtrInput    `pulumi:"recoveryPointThresholdInMinutes"`
}

func (InMageAzureV2PolicyInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAzureV2PolicyInput)(nil)).Elem()
}

func (i InMageAzureV2PolicyInputArgs) ToInMageAzureV2PolicyInputOutput() InMageAzureV2PolicyInputOutput {
	return i.ToInMageAzureV2PolicyInputOutputWithContext(context.Background())
}

func (i InMageAzureV2PolicyInputArgs) ToInMageAzureV2PolicyInputOutputWithContext(ctx context.Context) InMageAzureV2PolicyInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageAzureV2PolicyInputOutput)
}

type InMageAzureV2PolicyInputOutput struct{ *pulumi.OutputState }

func (InMageAzureV2PolicyInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAzureV2PolicyInput)(nil)).Elem()
}

func (o InMageAzureV2PolicyInputOutput) ToInMageAzureV2PolicyInputOutput() InMageAzureV2PolicyInputOutput {
	return o
}

func (o InMageAzureV2PolicyInputOutput) ToInMageAzureV2PolicyInputOutputWithContext(ctx context.Context) InMageAzureV2PolicyInputOutput {
	return o
}

func (o InMageAzureV2PolicyInputOutput) AppConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyInput) *int { return v.AppConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o InMageAzureV2PolicyInputOutput) CrashConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyInput) *int { return v.CrashConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o InMageAzureV2PolicyInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2PolicyInputOutput) MultiVmSyncStatus() pulumi.StringOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyInput) string { return v.MultiVmSyncStatus }).(pulumi.StringOutput)
}

func (o InMageAzureV2PolicyInputOutput) RecoveryPointHistory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyInput) *int { return v.RecoveryPointHistory }).(pulumi.IntPtrOutput)
}

func (o InMageAzureV2PolicyInputOutput) RecoveryPointThresholdInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2PolicyInput) *int { return v.RecoveryPointThresholdInMinutes }).(pulumi.IntPtrOutput)
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





type InMageAzureV2ProtectedDiskDetailsResponseInput interface {
	pulumi.Input

	ToInMageAzureV2ProtectedDiskDetailsResponseOutput() InMageAzureV2ProtectedDiskDetailsResponseOutput
	ToInMageAzureV2ProtectedDiskDetailsResponseOutputWithContext(context.Context) InMageAzureV2ProtectedDiskDetailsResponseOutput
}

type InMageAzureV2ProtectedDiskDetailsResponseArgs struct {
	DiskCapacityInBytes       pulumi.Float64PtrInput `pulumi:"diskCapacityInBytes"`
	DiskId                    pulumi.StringPtrInput  `pulumi:"diskId"`
	DiskName                  pulumi.StringPtrInput  `pulumi:"diskName"`
	DiskResized               pulumi.StringPtrInput  `pulumi:"diskResized"`
	FileSystemCapacityInBytes pulumi.Float64PtrInput `pulumi:"fileSystemCapacityInBytes"`
	HealthErrorCode           pulumi.StringPtrInput  `pulumi:"healthErrorCode"`
	LastRpoCalculatedTime     pulumi.StringPtrInput  `pulumi:"lastRpoCalculatedTime"`
	ProtectionStage           pulumi.StringPtrInput  `pulumi:"protectionStage"`
	PsDataInMegaBytes         pulumi.Float64PtrInput `pulumi:"psDataInMegaBytes"`
	ResyncDurationInSeconds   pulumi.Float64PtrInput `pulumi:"resyncDurationInSeconds"`
	ResyncProgressPercentage  pulumi.IntPtrInput     `pulumi:"resyncProgressPercentage"`
	ResyncRequired            pulumi.StringPtrInput  `pulumi:"resyncRequired"`
	RpoInSeconds              pulumi.Float64PtrInput `pulumi:"rpoInSeconds"`
	SourceDataInMegaBytes     pulumi.Float64PtrInput `pulumi:"sourceDataInMegaBytes"`
	TargetDataInMegaBytes     pulumi.Float64PtrInput `pulumi:"targetDataInMegaBytes"`
}

func (InMageAzureV2ProtectedDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAzureV2ProtectedDiskDetailsResponse)(nil)).Elem()
}

func (i InMageAzureV2ProtectedDiskDetailsResponseArgs) ToInMageAzureV2ProtectedDiskDetailsResponseOutput() InMageAzureV2ProtectedDiskDetailsResponseOutput {
	return i.ToInMageAzureV2ProtectedDiskDetailsResponseOutputWithContext(context.Background())
}

func (i InMageAzureV2ProtectedDiskDetailsResponseArgs) ToInMageAzureV2ProtectedDiskDetailsResponseOutputWithContext(ctx context.Context) InMageAzureV2ProtectedDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageAzureV2ProtectedDiskDetailsResponseOutput)
}





type InMageAzureV2ProtectedDiskDetailsResponseArrayInput interface {
	pulumi.Input

	ToInMageAzureV2ProtectedDiskDetailsResponseArrayOutput() InMageAzureV2ProtectedDiskDetailsResponseArrayOutput
	ToInMageAzureV2ProtectedDiskDetailsResponseArrayOutputWithContext(context.Context) InMageAzureV2ProtectedDiskDetailsResponseArrayOutput
}

type InMageAzureV2ProtectedDiskDetailsResponseArray []InMageAzureV2ProtectedDiskDetailsResponseInput

func (InMageAzureV2ProtectedDiskDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InMageAzureV2ProtectedDiskDetailsResponse)(nil)).Elem()
}

func (i InMageAzureV2ProtectedDiskDetailsResponseArray) ToInMageAzureV2ProtectedDiskDetailsResponseArrayOutput() InMageAzureV2ProtectedDiskDetailsResponseArrayOutput {
	return i.ToInMageAzureV2ProtectedDiskDetailsResponseArrayOutputWithContext(context.Background())
}

func (i InMageAzureV2ProtectedDiskDetailsResponseArray) ToInMageAzureV2ProtectedDiskDetailsResponseArrayOutputWithContext(ctx context.Context) InMageAzureV2ProtectedDiskDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageAzureV2ProtectedDiskDetailsResponseArrayOutput)
}

type InMageAzureV2ProtectedDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (InMageAzureV2ProtectedDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAzureV2ProtectedDiskDetailsResponse)(nil)).Elem()
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) ToInMageAzureV2ProtectedDiskDetailsResponseOutput() InMageAzureV2ProtectedDiskDetailsResponseOutput {
	return o
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) ToInMageAzureV2ProtectedDiskDetailsResponseOutputWithContext(ctx context.Context) InMageAzureV2ProtectedDiskDetailsResponseOutput {
	return o
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) DiskCapacityInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *float64 { return v.DiskCapacityInBytes }).(pulumi.Float64PtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) DiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *string { return v.DiskId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) DiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *string { return v.DiskName }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) DiskResized() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *string { return v.DiskResized }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) FileSystemCapacityInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *float64 { return v.FileSystemCapacityInBytes }).(pulumi.Float64PtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) HealthErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *string { return v.HealthErrorCode }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) LastRpoCalculatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *string { return v.LastRpoCalculatedTime }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) ProtectionStage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *string { return v.ProtectionStage }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) PsDataInMegaBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *float64 { return v.PsDataInMegaBytes }).(pulumi.Float64PtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) ResyncDurationInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *float64 { return v.ResyncDurationInSeconds }).(pulumi.Float64PtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) ResyncProgressPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *int { return v.ResyncProgressPercentage }).(pulumi.IntPtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) ResyncRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *string { return v.ResyncRequired }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) RpoInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *float64 { return v.RpoInSeconds }).(pulumi.Float64PtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) SourceDataInMegaBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *float64 { return v.SourceDataInMegaBytes }).(pulumi.Float64PtrOutput)
}

func (o InMageAzureV2ProtectedDiskDetailsResponseOutput) TargetDataInMegaBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageAzureV2ProtectedDiskDetailsResponse) *float64 { return v.TargetDataInMegaBytes }).(pulumi.Float64PtrOutput)
}

type InMageAzureV2ProtectedDiskDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (InMageAzureV2ProtectedDiskDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InMageAzureV2ProtectedDiskDetailsResponse)(nil)).Elem()
}

func (o InMageAzureV2ProtectedDiskDetailsResponseArrayOutput) ToInMageAzureV2ProtectedDiskDetailsResponseArrayOutput() InMageAzureV2ProtectedDiskDetailsResponseArrayOutput {
	return o
}

func (o InMageAzureV2ProtectedDiskDetailsResponseArrayOutput) ToInMageAzureV2ProtectedDiskDetailsResponseArrayOutputWithContext(ctx context.Context) InMageAzureV2ProtectedDiskDetailsResponseArrayOutput {
	return o
}

func (o InMageAzureV2ProtectedDiskDetailsResponseArrayOutput) Index(i pulumi.IntInput) InMageAzureV2ProtectedDiskDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InMageAzureV2ProtectedDiskDetailsResponse {
		return vs[0].([]InMageAzureV2ProtectedDiskDetailsResponse)[vs[1].(int)]
	}).(InMageAzureV2ProtectedDiskDetailsResponseOutput)
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
	SelectedSourceNicId              *string                                     `pulumi:"selectedSourceNicId"`
	SourceVmCpuCount                 *int                                        `pulumi:"sourceVmCpuCount"`
	SourceVmRamSizeInMB              *int                                        `pulumi:"sourceVmRamSizeInMB"`
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





type InMageAzureV2ReplicationDetailsResponseInput interface {
	pulumi.Input

	ToInMageAzureV2ReplicationDetailsResponseOutput() InMageAzureV2ReplicationDetailsResponseOutput
	ToInMageAzureV2ReplicationDetailsResponseOutputWithContext(context.Context) InMageAzureV2ReplicationDetailsResponseOutput
}

type InMageAzureV2ReplicationDetailsResponseArgs struct {
	AgentExpiryDate                  pulumi.StringPtrInput                               `pulumi:"agentExpiryDate"`
	AgentVersion                     pulumi.StringPtrInput                               `pulumi:"agentVersion"`
	AzureVMDiskDetails               AzureVmDiskDetailsResponseArrayInput                `pulumi:"azureVMDiskDetails"`
	CompressedDataRateInMB           pulumi.Float64PtrInput                              `pulumi:"compressedDataRateInMB"`
	Datastores                       pulumi.StringArrayInput                             `pulumi:"datastores"`
	DiscoveryType                    pulumi.StringPtrInput                               `pulumi:"discoveryType"`
	DiskResized                      pulumi.StringPtrInput                               `pulumi:"diskResized"`
	EnableRdpOnTargetOption          pulumi.StringPtrInput                               `pulumi:"enableRdpOnTargetOption"`
	InfrastructureVmId               pulumi.StringPtrInput                               `pulumi:"infrastructureVmId"`
	InstanceType                     pulumi.StringInput                                  `pulumi:"instanceType"`
	IpAddress                        pulumi.StringPtrInput                               `pulumi:"ipAddress"`
	IsAgentUpdateRequired            pulumi.StringPtrInput                               `pulumi:"isAgentUpdateRequired"`
	IsRebootAfterUpdateRequired      pulumi.StringPtrInput                               `pulumi:"isRebootAfterUpdateRequired"`
	LastHeartbeat                    pulumi.StringPtrInput                               `pulumi:"lastHeartbeat"`
	LastRpoCalculatedTime            pulumi.StringPtrInput                               `pulumi:"lastRpoCalculatedTime"`
	LastUpdateReceivedTime           pulumi.StringPtrInput                               `pulumi:"lastUpdateReceivedTime"`
	LicenseType                      pulumi.StringPtrInput                               `pulumi:"licenseType"`
	MasterTargetId                   pulumi.StringPtrInput                               `pulumi:"masterTargetId"`
	MultiVmGroupId                   pulumi.StringPtrInput                               `pulumi:"multiVmGroupId"`
	MultiVmGroupName                 pulumi.StringPtrInput                               `pulumi:"multiVmGroupName"`
	MultiVmSyncStatus                pulumi.StringPtrInput                               `pulumi:"multiVmSyncStatus"`
	OsDiskId                         pulumi.StringPtrInput                               `pulumi:"osDiskId"`
	OsType                           pulumi.StringPtrInput                               `pulumi:"osType"`
	OsVersion                        pulumi.StringPtrInput                               `pulumi:"osVersion"`
	ProcessServerId                  pulumi.StringPtrInput                               `pulumi:"processServerId"`
	ProtectedDisks                   InMageAzureV2ProtectedDiskDetailsResponseArrayInput `pulumi:"protectedDisks"`
	ProtectionStage                  pulumi.StringPtrInput                               `pulumi:"protectionStage"`
	RecoveryAvailabilitySetId        pulumi.StringPtrInput                               `pulumi:"recoveryAvailabilitySetId"`
	RecoveryAzureLogStorageAccountId pulumi.StringPtrInput                               `pulumi:"recoveryAzureLogStorageAccountId"`
	RecoveryAzureResourceGroupId     pulumi.StringPtrInput                               `pulumi:"recoveryAzureResourceGroupId"`
	RecoveryAzureStorageAccount      pulumi.StringPtrInput                               `pulumi:"recoveryAzureStorageAccount"`
	RecoveryAzureVMName              pulumi.StringPtrInput                               `pulumi:"recoveryAzureVMName"`
	RecoveryAzureVMSize              pulumi.StringPtrInput                               `pulumi:"recoveryAzureVMSize"`
	ReplicaId                        pulumi.StringPtrInput                               `pulumi:"replicaId"`
	ResyncProgressPercentage         pulumi.IntPtrInput                                  `pulumi:"resyncProgressPercentage"`
	RpoInSeconds                     pulumi.Float64PtrInput                              `pulumi:"rpoInSeconds"`
	SelectedRecoveryAzureNetworkId   pulumi.StringPtrInput                               `pulumi:"selectedRecoveryAzureNetworkId"`
	SelectedSourceNicId              pulumi.StringPtrInput                               `pulumi:"selectedSourceNicId"`
	SourceVmCpuCount                 pulumi.IntPtrInput                                  `pulumi:"sourceVmCpuCount"`
	SourceVmRamSizeInMB              pulumi.IntPtrInput                                  `pulumi:"sourceVmRamSizeInMB"`
	TargetVmId                       pulumi.StringPtrInput                               `pulumi:"targetVmId"`
	UncompressedDataRateInMB         pulumi.Float64PtrInput                              `pulumi:"uncompressedDataRateInMB"`
	UseManagedDisks                  pulumi.StringPtrInput                               `pulumi:"useManagedDisks"`
	VCenterInfrastructureId          pulumi.StringPtrInput                               `pulumi:"vCenterInfrastructureId"`
	ValidationErrors                 HealthErrorResponseArrayInput                       `pulumi:"validationErrors"`
	VhdName                          pulumi.StringPtrInput                               `pulumi:"vhdName"`
	VmId                             pulumi.StringPtrInput                               `pulumi:"vmId"`
	VmNics                           VMNicDetailsResponseArrayInput                      `pulumi:"vmNics"`
	VmProtectionState                pulumi.StringPtrInput                               `pulumi:"vmProtectionState"`
	VmProtectionStateDescription     pulumi.StringPtrInput                               `pulumi:"vmProtectionStateDescription"`
}

func (InMageAzureV2ReplicationDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAzureV2ReplicationDetailsResponse)(nil)).Elem()
}

func (i InMageAzureV2ReplicationDetailsResponseArgs) ToInMageAzureV2ReplicationDetailsResponseOutput() InMageAzureV2ReplicationDetailsResponseOutput {
	return i.ToInMageAzureV2ReplicationDetailsResponseOutputWithContext(context.Background())
}

func (i InMageAzureV2ReplicationDetailsResponseArgs) ToInMageAzureV2ReplicationDetailsResponseOutputWithContext(ctx context.Context) InMageAzureV2ReplicationDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageAzureV2ReplicationDetailsResponseOutput)
}

type InMageAzureV2ReplicationDetailsResponseOutput struct{ *pulumi.OutputState }

func (InMageAzureV2ReplicationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageAzureV2ReplicationDetailsResponse)(nil)).Elem()
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) ToInMageAzureV2ReplicationDetailsResponseOutput() InMageAzureV2ReplicationDetailsResponseOutput {
	return o
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) ToInMageAzureV2ReplicationDetailsResponseOutputWithContext(ctx context.Context) InMageAzureV2ReplicationDetailsResponseOutput {
	return o
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) AgentExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.AgentExpiryDate }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) AzureVMDiskDetails() AzureVmDiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) []AzureVmDiskDetailsResponse {
		return v.AzureVMDiskDetails
	}).(AzureVmDiskDetailsResponseArrayOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) CompressedDataRateInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *float64 { return v.CompressedDataRateInMB }).(pulumi.Float64PtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) Datastores() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) []string { return v.Datastores }).(pulumi.StringArrayOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) DiscoveryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.DiscoveryType }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) DiskResized() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.DiskResized }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) EnableRdpOnTargetOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.EnableRdpOnTargetOption }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) InfrastructureVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.InfrastructureVmId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) IsAgentUpdateRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.IsAgentUpdateRequired }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) IsRebootAfterUpdateRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.IsRebootAfterUpdateRequired }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) LastHeartbeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.LastHeartbeat }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) LastRpoCalculatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.LastRpoCalculatedTime }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) LastUpdateReceivedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.LastUpdateReceivedTime }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.LicenseType }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) MasterTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.MasterTargetId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) MultiVmGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.MultiVmGroupId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) MultiVmGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.MultiVmGroupName }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) MultiVmSyncStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.MultiVmSyncStatus }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) OsDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.OsDiskId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.OsVersion }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) ProcessServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.ProcessServerId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) ProtectedDisks() InMageAzureV2ProtectedDiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) []InMageAzureV2ProtectedDiskDetailsResponse {
		return v.ProtectedDisks
	}).(InMageAzureV2ProtectedDiskDetailsResponseArrayOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) ProtectionStage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.ProtectionStage }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) RecoveryAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.RecoveryAvailabilitySetId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) RecoveryAzureLogStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.RecoveryAzureLogStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) RecoveryAzureResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.RecoveryAzureResourceGroupId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) RecoveryAzureStorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.RecoveryAzureStorageAccount }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) RecoveryAzureVMName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.RecoveryAzureVMName }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) RecoveryAzureVMSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.RecoveryAzureVMSize }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) ReplicaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.ReplicaId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) ResyncProgressPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *int { return v.ResyncProgressPercentage }).(pulumi.IntPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) RpoInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *float64 { return v.RpoInSeconds }).(pulumi.Float64PtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) SelectedRecoveryAzureNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.SelectedRecoveryAzureNetworkId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) SelectedSourceNicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.SelectedSourceNicId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) SourceVmCpuCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *int { return v.SourceVmCpuCount }).(pulumi.IntPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) SourceVmRamSizeInMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *int { return v.SourceVmRamSizeInMB }).(pulumi.IntPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) TargetVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.TargetVmId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) UncompressedDataRateInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *float64 { return v.UncompressedDataRateInMB }).(pulumi.Float64PtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) UseManagedDisks() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.UseManagedDisks }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) VCenterInfrastructureId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.VCenterInfrastructureId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) ValidationErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) []HealthErrorResponse { return v.ValidationErrors }).(HealthErrorResponseArrayOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) VhdName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.VhdName }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) VmNics() VMNicDetailsResponseArrayOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) []VMNicDetailsResponse { return v.VmNics }).(VMNicDetailsResponseArrayOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) VmProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.VmProtectionState }).(pulumi.StringPtrOutput)
}

func (o InMageAzureV2ReplicationDetailsResponseOutput) VmProtectionStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageAzureV2ReplicationDetailsResponse) *string { return v.VmProtectionStateDescription }).(pulumi.StringPtrOutput)
}

type InMageBasePolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes *int    `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    string  `pulumi:"instanceType"`
	MultiVmSyncStatus               *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes *int    `pulumi:"recoveryPointThresholdInMinutes"`
}





type InMageBasePolicyDetailsResponseInput interface {
	pulumi.Input

	ToInMageBasePolicyDetailsResponseOutput() InMageBasePolicyDetailsResponseOutput
	ToInMageBasePolicyDetailsResponseOutputWithContext(context.Context) InMageBasePolicyDetailsResponseOutput
}

type InMageBasePolicyDetailsResponseArgs struct {
	AppConsistentFrequencyInMinutes pulumi.IntPtrInput    `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    pulumi.StringInput    `pulumi:"instanceType"`
	MultiVmSyncStatus               pulumi.StringPtrInput `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            pulumi.IntPtrInput    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes pulumi.IntPtrInput    `pulumi:"recoveryPointThresholdInMinutes"`
}

func (InMageBasePolicyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageBasePolicyDetailsResponse)(nil)).Elem()
}

func (i InMageBasePolicyDetailsResponseArgs) ToInMageBasePolicyDetailsResponseOutput() InMageBasePolicyDetailsResponseOutput {
	return i.ToInMageBasePolicyDetailsResponseOutputWithContext(context.Background())
}

func (i InMageBasePolicyDetailsResponseArgs) ToInMageBasePolicyDetailsResponseOutputWithContext(ctx context.Context) InMageBasePolicyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageBasePolicyDetailsResponseOutput)
}

type InMageBasePolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (InMageBasePolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageBasePolicyDetailsResponse)(nil)).Elem()
}

func (o InMageBasePolicyDetailsResponseOutput) ToInMageBasePolicyDetailsResponseOutput() InMageBasePolicyDetailsResponseOutput {
	return o
}

func (o InMageBasePolicyDetailsResponseOutput) ToInMageBasePolicyDetailsResponseOutputWithContext(ctx context.Context) InMageBasePolicyDetailsResponseOutput {
	return o
}

func (o InMageBasePolicyDetailsResponseOutput) AppConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageBasePolicyDetailsResponse) *int { return v.AppConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o InMageBasePolicyDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v InMageBasePolicyDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o InMageBasePolicyDetailsResponseOutput) MultiVmSyncStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageBasePolicyDetailsResponse) *string { return v.MultiVmSyncStatus }).(pulumi.StringPtrOutput)
}

func (o InMageBasePolicyDetailsResponseOutput) RecoveryPointHistory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageBasePolicyDetailsResponse) *int { return v.RecoveryPointHistory }).(pulumi.IntPtrOutput)
}

func (o InMageBasePolicyDetailsResponseOutput) RecoveryPointThresholdInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageBasePolicyDetailsResponse) *int { return v.RecoveryPointThresholdInMinutes }).(pulumi.IntPtrOutput)
}

type InMageDiskExclusionInput struct {
	DiskSignatureOptions []InMageDiskSignatureExclusionOptions `pulumi:"diskSignatureOptions"`
	VolumeOptions        []InMageVolumeExclusionOptions        `pulumi:"volumeOptions"`
}





type InMageDiskExclusionInputInput interface {
	pulumi.Input

	ToInMageDiskExclusionInputOutput() InMageDiskExclusionInputOutput
	ToInMageDiskExclusionInputOutputWithContext(context.Context) InMageDiskExclusionInputOutput
}

type InMageDiskExclusionInputArgs struct {
	DiskSignatureOptions InMageDiskSignatureExclusionOptionsArrayInput `pulumi:"diskSignatureOptions"`
	VolumeOptions        InMageVolumeExclusionOptionsArrayInput        `pulumi:"volumeOptions"`
}

func (InMageDiskExclusionInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageDiskExclusionInput)(nil)).Elem()
}

func (i InMageDiskExclusionInputArgs) ToInMageDiskExclusionInputOutput() InMageDiskExclusionInputOutput {
	return i.ToInMageDiskExclusionInputOutputWithContext(context.Background())
}

func (i InMageDiskExclusionInputArgs) ToInMageDiskExclusionInputOutputWithContext(ctx context.Context) InMageDiskExclusionInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageDiskExclusionInputOutput)
}

func (i InMageDiskExclusionInputArgs) ToInMageDiskExclusionInputPtrOutput() InMageDiskExclusionInputPtrOutput {
	return i.ToInMageDiskExclusionInputPtrOutputWithContext(context.Background())
}

func (i InMageDiskExclusionInputArgs) ToInMageDiskExclusionInputPtrOutputWithContext(ctx context.Context) InMageDiskExclusionInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageDiskExclusionInputOutput).ToInMageDiskExclusionInputPtrOutputWithContext(ctx)
}









type InMageDiskExclusionInputPtrInput interface {
	pulumi.Input

	ToInMageDiskExclusionInputPtrOutput() InMageDiskExclusionInputPtrOutput
	ToInMageDiskExclusionInputPtrOutputWithContext(context.Context) InMageDiskExclusionInputPtrOutput
}

type inMageDiskExclusionInputPtrType InMageDiskExclusionInputArgs

func InMageDiskExclusionInputPtr(v *InMageDiskExclusionInputArgs) InMageDiskExclusionInputPtrInput {
	return (*inMageDiskExclusionInputPtrType)(v)
}

func (*inMageDiskExclusionInputPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InMageDiskExclusionInput)(nil)).Elem()
}

func (i *inMageDiskExclusionInputPtrType) ToInMageDiskExclusionInputPtrOutput() InMageDiskExclusionInputPtrOutput {
	return i.ToInMageDiskExclusionInputPtrOutputWithContext(context.Background())
}

func (i *inMageDiskExclusionInputPtrType) ToInMageDiskExclusionInputPtrOutputWithContext(ctx context.Context) InMageDiskExclusionInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageDiskExclusionInputPtrOutput)
}

type InMageDiskExclusionInputOutput struct{ *pulumi.OutputState }

func (InMageDiskExclusionInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageDiskExclusionInput)(nil)).Elem()
}

func (o InMageDiskExclusionInputOutput) ToInMageDiskExclusionInputOutput() InMageDiskExclusionInputOutput {
	return o
}

func (o InMageDiskExclusionInputOutput) ToInMageDiskExclusionInputOutputWithContext(ctx context.Context) InMageDiskExclusionInputOutput {
	return o
}

func (o InMageDiskExclusionInputOutput) ToInMageDiskExclusionInputPtrOutput() InMageDiskExclusionInputPtrOutput {
	return o.ToInMageDiskExclusionInputPtrOutputWithContext(context.Background())
}

func (o InMageDiskExclusionInputOutput) ToInMageDiskExclusionInputPtrOutputWithContext(ctx context.Context) InMageDiskExclusionInputPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InMageDiskExclusionInput) *InMageDiskExclusionInput {
		return &v
	}).(InMageDiskExclusionInputPtrOutput)
}

func (o InMageDiskExclusionInputOutput) DiskSignatureOptions() InMageDiskSignatureExclusionOptionsArrayOutput {
	return o.ApplyT(func(v InMageDiskExclusionInput) []InMageDiskSignatureExclusionOptions { return v.DiskSignatureOptions }).(InMageDiskSignatureExclusionOptionsArrayOutput)
}

func (o InMageDiskExclusionInputOutput) VolumeOptions() InMageVolumeExclusionOptionsArrayOutput {
	return o.ApplyT(func(v InMageDiskExclusionInput) []InMageVolumeExclusionOptions { return v.VolumeOptions }).(InMageVolumeExclusionOptionsArrayOutput)
}

type InMageDiskExclusionInputPtrOutput struct{ *pulumi.OutputState }

func (InMageDiskExclusionInputPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InMageDiskExclusionInput)(nil)).Elem()
}

func (o InMageDiskExclusionInputPtrOutput) ToInMageDiskExclusionInputPtrOutput() InMageDiskExclusionInputPtrOutput {
	return o
}

func (o InMageDiskExclusionInputPtrOutput) ToInMageDiskExclusionInputPtrOutputWithContext(ctx context.Context) InMageDiskExclusionInputPtrOutput {
	return o
}

func (o InMageDiskExclusionInputPtrOutput) Elem() InMageDiskExclusionInputOutput {
	return o.ApplyT(func(v *InMageDiskExclusionInput) InMageDiskExclusionInput {
		if v != nil {
			return *v
		}
		var ret InMageDiskExclusionInput
		return ret
	}).(InMageDiskExclusionInputOutput)
}

func (o InMageDiskExclusionInputPtrOutput) DiskSignatureOptions() InMageDiskSignatureExclusionOptionsArrayOutput {
	return o.ApplyT(func(v *InMageDiskExclusionInput) []InMageDiskSignatureExclusionOptions {
		if v == nil {
			return nil
		}
		return v.DiskSignatureOptions
	}).(InMageDiskSignatureExclusionOptionsArrayOutput)
}

func (o InMageDiskExclusionInputPtrOutput) VolumeOptions() InMageVolumeExclusionOptionsArrayOutput {
	return o.ApplyT(func(v *InMageDiskExclusionInput) []InMageVolumeExclusionOptions {
		if v == nil {
			return nil
		}
		return v.VolumeOptions
	}).(InMageVolumeExclusionOptionsArrayOutput)
}

type InMageDiskSignatureExclusionOptions struct {
	DiskSignature *string `pulumi:"diskSignature"`
}





type InMageDiskSignatureExclusionOptionsInput interface {
	pulumi.Input

	ToInMageDiskSignatureExclusionOptionsOutput() InMageDiskSignatureExclusionOptionsOutput
	ToInMageDiskSignatureExclusionOptionsOutputWithContext(context.Context) InMageDiskSignatureExclusionOptionsOutput
}

type InMageDiskSignatureExclusionOptionsArgs struct {
	DiskSignature pulumi.StringPtrInput `pulumi:"diskSignature"`
}

func (InMageDiskSignatureExclusionOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageDiskSignatureExclusionOptions)(nil)).Elem()
}

func (i InMageDiskSignatureExclusionOptionsArgs) ToInMageDiskSignatureExclusionOptionsOutput() InMageDiskSignatureExclusionOptionsOutput {
	return i.ToInMageDiskSignatureExclusionOptionsOutputWithContext(context.Background())
}

func (i InMageDiskSignatureExclusionOptionsArgs) ToInMageDiskSignatureExclusionOptionsOutputWithContext(ctx context.Context) InMageDiskSignatureExclusionOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageDiskSignatureExclusionOptionsOutput)
}





type InMageDiskSignatureExclusionOptionsArrayInput interface {
	pulumi.Input

	ToInMageDiskSignatureExclusionOptionsArrayOutput() InMageDiskSignatureExclusionOptionsArrayOutput
	ToInMageDiskSignatureExclusionOptionsArrayOutputWithContext(context.Context) InMageDiskSignatureExclusionOptionsArrayOutput
}

type InMageDiskSignatureExclusionOptionsArray []InMageDiskSignatureExclusionOptionsInput

func (InMageDiskSignatureExclusionOptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InMageDiskSignatureExclusionOptions)(nil)).Elem()
}

func (i InMageDiskSignatureExclusionOptionsArray) ToInMageDiskSignatureExclusionOptionsArrayOutput() InMageDiskSignatureExclusionOptionsArrayOutput {
	return i.ToInMageDiskSignatureExclusionOptionsArrayOutputWithContext(context.Background())
}

func (i InMageDiskSignatureExclusionOptionsArray) ToInMageDiskSignatureExclusionOptionsArrayOutputWithContext(ctx context.Context) InMageDiskSignatureExclusionOptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageDiskSignatureExclusionOptionsArrayOutput)
}

type InMageDiskSignatureExclusionOptionsOutput struct{ *pulumi.OutputState }

func (InMageDiskSignatureExclusionOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageDiskSignatureExclusionOptions)(nil)).Elem()
}

func (o InMageDiskSignatureExclusionOptionsOutput) ToInMageDiskSignatureExclusionOptionsOutput() InMageDiskSignatureExclusionOptionsOutput {
	return o
}

func (o InMageDiskSignatureExclusionOptionsOutput) ToInMageDiskSignatureExclusionOptionsOutputWithContext(ctx context.Context) InMageDiskSignatureExclusionOptionsOutput {
	return o
}

func (o InMageDiskSignatureExclusionOptionsOutput) DiskSignature() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageDiskSignatureExclusionOptions) *string { return v.DiskSignature }).(pulumi.StringPtrOutput)
}

type InMageDiskSignatureExclusionOptionsArrayOutput struct{ *pulumi.OutputState }

func (InMageDiskSignatureExclusionOptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InMageDiskSignatureExclusionOptions)(nil)).Elem()
}

func (o InMageDiskSignatureExclusionOptionsArrayOutput) ToInMageDiskSignatureExclusionOptionsArrayOutput() InMageDiskSignatureExclusionOptionsArrayOutput {
	return o
}

func (o InMageDiskSignatureExclusionOptionsArrayOutput) ToInMageDiskSignatureExclusionOptionsArrayOutputWithContext(ctx context.Context) InMageDiskSignatureExclusionOptionsArrayOutput {
	return o
}

func (o InMageDiskSignatureExclusionOptionsArrayOutput) Index(i pulumi.IntInput) InMageDiskSignatureExclusionOptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InMageDiskSignatureExclusionOptions {
		return vs[0].([]InMageDiskSignatureExclusionOptions)[vs[1].(int)]
	}).(InMageDiskSignatureExclusionOptionsOutput)
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





type InMageEnableProtectionInputInput interface {
	pulumi.Input

	ToInMageEnableProtectionInputOutput() InMageEnableProtectionInputOutput
	ToInMageEnableProtectionInputOutputWithContext(context.Context) InMageEnableProtectionInputOutput
}

type InMageEnableProtectionInputArgs struct {
	DatastoreName      pulumi.StringPtrInput            `pulumi:"datastoreName"`
	DiskExclusionInput InMageDiskExclusionInputPtrInput `pulumi:"diskExclusionInput"`
	DisksToInclude     pulumi.StringArrayInput          `pulumi:"disksToInclude"`
	InstanceType       pulumi.StringPtrInput            `pulumi:"instanceType"`
	MasterTargetId     pulumi.StringInput               `pulumi:"masterTargetId"`
	MultiVmGroupId     pulumi.StringInput               `pulumi:"multiVmGroupId"`
	MultiVmGroupName   pulumi.StringInput               `pulumi:"multiVmGroupName"`
	ProcessServerId    pulumi.StringInput               `pulumi:"processServerId"`
	RetentionDrive     pulumi.StringInput               `pulumi:"retentionDrive"`
	RunAsAccountId     pulumi.StringPtrInput            `pulumi:"runAsAccountId"`
	VmFriendlyName     pulumi.StringPtrInput            `pulumi:"vmFriendlyName"`
}

func (InMageEnableProtectionInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageEnableProtectionInput)(nil)).Elem()
}

func (i InMageEnableProtectionInputArgs) ToInMageEnableProtectionInputOutput() InMageEnableProtectionInputOutput {
	return i.ToInMageEnableProtectionInputOutputWithContext(context.Background())
}

func (i InMageEnableProtectionInputArgs) ToInMageEnableProtectionInputOutputWithContext(ctx context.Context) InMageEnableProtectionInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageEnableProtectionInputOutput)
}

type InMageEnableProtectionInputOutput struct{ *pulumi.OutputState }

func (InMageEnableProtectionInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageEnableProtectionInput)(nil)).Elem()
}

func (o InMageEnableProtectionInputOutput) ToInMageEnableProtectionInputOutput() InMageEnableProtectionInputOutput {
	return o
}

func (o InMageEnableProtectionInputOutput) ToInMageEnableProtectionInputOutputWithContext(ctx context.Context) InMageEnableProtectionInputOutput {
	return o
}

func (o InMageEnableProtectionInputOutput) DatastoreName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) *string { return v.DatastoreName }).(pulumi.StringPtrOutput)
}

func (o InMageEnableProtectionInputOutput) DiskExclusionInput() InMageDiskExclusionInputPtrOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) *InMageDiskExclusionInput { return v.DiskExclusionInput }).(InMageDiskExclusionInputPtrOutput)
}

func (o InMageEnableProtectionInputOutput) DisksToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) []string { return v.DisksToInclude }).(pulumi.StringArrayOutput)
}

func (o InMageEnableProtectionInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o InMageEnableProtectionInputOutput) MasterTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) string { return v.MasterTargetId }).(pulumi.StringOutput)
}

func (o InMageEnableProtectionInputOutput) MultiVmGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) string { return v.MultiVmGroupId }).(pulumi.StringOutput)
}

func (o InMageEnableProtectionInputOutput) MultiVmGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) string { return v.MultiVmGroupName }).(pulumi.StringOutput)
}

func (o InMageEnableProtectionInputOutput) ProcessServerId() pulumi.StringOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) string { return v.ProcessServerId }).(pulumi.StringOutput)
}

func (o InMageEnableProtectionInputOutput) RetentionDrive() pulumi.StringOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) string { return v.RetentionDrive }).(pulumi.StringOutput)
}

func (o InMageEnableProtectionInputOutput) RunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) *string { return v.RunAsAccountId }).(pulumi.StringPtrOutput)
}

func (o InMageEnableProtectionInputOutput) VmFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageEnableProtectionInput) *string { return v.VmFriendlyName }).(pulumi.StringPtrOutput)
}

type InMagePolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes *int    `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    string  `pulumi:"instanceType"`
	MultiVmSyncStatus               *string `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes *int    `pulumi:"recoveryPointThresholdInMinutes"`
}





type InMagePolicyDetailsResponseInput interface {
	pulumi.Input

	ToInMagePolicyDetailsResponseOutput() InMagePolicyDetailsResponseOutput
	ToInMagePolicyDetailsResponseOutputWithContext(context.Context) InMagePolicyDetailsResponseOutput
}

type InMagePolicyDetailsResponseArgs struct {
	AppConsistentFrequencyInMinutes pulumi.IntPtrInput    `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    pulumi.StringInput    `pulumi:"instanceType"`
	MultiVmSyncStatus               pulumi.StringPtrInput `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            pulumi.IntPtrInput    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes pulumi.IntPtrInput    `pulumi:"recoveryPointThresholdInMinutes"`
}

func (InMagePolicyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMagePolicyDetailsResponse)(nil)).Elem()
}

func (i InMagePolicyDetailsResponseArgs) ToInMagePolicyDetailsResponseOutput() InMagePolicyDetailsResponseOutput {
	return i.ToInMagePolicyDetailsResponseOutputWithContext(context.Background())
}

func (i InMagePolicyDetailsResponseArgs) ToInMagePolicyDetailsResponseOutputWithContext(ctx context.Context) InMagePolicyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMagePolicyDetailsResponseOutput)
}

type InMagePolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (InMagePolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMagePolicyDetailsResponse)(nil)).Elem()
}

func (o InMagePolicyDetailsResponseOutput) ToInMagePolicyDetailsResponseOutput() InMagePolicyDetailsResponseOutput {
	return o
}

func (o InMagePolicyDetailsResponseOutput) ToInMagePolicyDetailsResponseOutputWithContext(ctx context.Context) InMagePolicyDetailsResponseOutput {
	return o
}

func (o InMagePolicyDetailsResponseOutput) AppConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMagePolicyDetailsResponse) *int { return v.AppConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o InMagePolicyDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v InMagePolicyDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o InMagePolicyDetailsResponseOutput) MultiVmSyncStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMagePolicyDetailsResponse) *string { return v.MultiVmSyncStatus }).(pulumi.StringPtrOutput)
}

func (o InMagePolicyDetailsResponseOutput) RecoveryPointHistory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMagePolicyDetailsResponse) *int { return v.RecoveryPointHistory }).(pulumi.IntPtrOutput)
}

func (o InMagePolicyDetailsResponseOutput) RecoveryPointThresholdInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMagePolicyDetailsResponse) *int { return v.RecoveryPointThresholdInMinutes }).(pulumi.IntPtrOutput)
}

type InMagePolicyInput struct {
	AppConsistentFrequencyInMinutes *int    `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    *string `pulumi:"instanceType"`
	MultiVmSyncStatus               string  `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            *int    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes *int    `pulumi:"recoveryPointThresholdInMinutes"`
}





type InMagePolicyInputInput interface {
	pulumi.Input

	ToInMagePolicyInputOutput() InMagePolicyInputOutput
	ToInMagePolicyInputOutputWithContext(context.Context) InMagePolicyInputOutput
}

type InMagePolicyInputArgs struct {
	AppConsistentFrequencyInMinutes pulumi.IntPtrInput    `pulumi:"appConsistentFrequencyInMinutes"`
	InstanceType                    pulumi.StringPtrInput `pulumi:"instanceType"`
	MultiVmSyncStatus               pulumi.StringInput    `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory            pulumi.IntPtrInput    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes pulumi.IntPtrInput    `pulumi:"recoveryPointThresholdInMinutes"`
}

func (InMagePolicyInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMagePolicyInput)(nil)).Elem()
}

func (i InMagePolicyInputArgs) ToInMagePolicyInputOutput() InMagePolicyInputOutput {
	return i.ToInMagePolicyInputOutputWithContext(context.Background())
}

func (i InMagePolicyInputArgs) ToInMagePolicyInputOutputWithContext(ctx context.Context) InMagePolicyInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMagePolicyInputOutput)
}

type InMagePolicyInputOutput struct{ *pulumi.OutputState }

func (InMagePolicyInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMagePolicyInput)(nil)).Elem()
}

func (o InMagePolicyInputOutput) ToInMagePolicyInputOutput() InMagePolicyInputOutput {
	return o
}

func (o InMagePolicyInputOutput) ToInMagePolicyInputOutputWithContext(ctx context.Context) InMagePolicyInputOutput {
	return o
}

func (o InMagePolicyInputOutput) AppConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMagePolicyInput) *int { return v.AppConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o InMagePolicyInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMagePolicyInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o InMagePolicyInputOutput) MultiVmSyncStatus() pulumi.StringOutput {
	return o.ApplyT(func(v InMagePolicyInput) string { return v.MultiVmSyncStatus }).(pulumi.StringOutput)
}

func (o InMagePolicyInputOutput) RecoveryPointHistory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMagePolicyInput) *int { return v.RecoveryPointHistory }).(pulumi.IntPtrOutput)
}

func (o InMagePolicyInputOutput) RecoveryPointThresholdInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMagePolicyInput) *int { return v.RecoveryPointThresholdInMinutes }).(pulumi.IntPtrOutput)
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





type InMageProtectedDiskDetailsResponseInput interface {
	pulumi.Input

	ToInMageProtectedDiskDetailsResponseOutput() InMageProtectedDiskDetailsResponseOutput
	ToInMageProtectedDiskDetailsResponseOutputWithContext(context.Context) InMageProtectedDiskDetailsResponseOutput
}

type InMageProtectedDiskDetailsResponseArgs struct {
	DiskCapacityInBytes       pulumi.Float64PtrInput `pulumi:"diskCapacityInBytes"`
	DiskId                    pulumi.StringPtrInput  `pulumi:"diskId"`
	DiskName                  pulumi.StringPtrInput  `pulumi:"diskName"`
	DiskResized               pulumi.StringPtrInput  `pulumi:"diskResized"`
	FileSystemCapacityInBytes pulumi.Float64PtrInput `pulumi:"fileSystemCapacityInBytes"`
	HealthErrorCode           pulumi.StringPtrInput  `pulumi:"healthErrorCode"`
	LastRpoCalculatedTime     pulumi.StringPtrInput  `pulumi:"lastRpoCalculatedTime"`
	ProtectionStage           pulumi.StringPtrInput  `pulumi:"protectionStage"`
	PsDataInMB                pulumi.Float64PtrInput `pulumi:"psDataInMB"`
	ResyncDurationInSeconds   pulumi.Float64PtrInput `pulumi:"resyncDurationInSeconds"`
	ResyncProgressPercentage  pulumi.IntPtrInput     `pulumi:"resyncProgressPercentage"`
	ResyncRequired            pulumi.StringPtrInput  `pulumi:"resyncRequired"`
	RpoInSeconds              pulumi.Float64PtrInput `pulumi:"rpoInSeconds"`
	SourceDataInMB            pulumi.Float64PtrInput `pulumi:"sourceDataInMB"`
	TargetDataInMB            pulumi.Float64PtrInput `pulumi:"targetDataInMB"`
}

func (InMageProtectedDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageProtectedDiskDetailsResponse)(nil)).Elem()
}

func (i InMageProtectedDiskDetailsResponseArgs) ToInMageProtectedDiskDetailsResponseOutput() InMageProtectedDiskDetailsResponseOutput {
	return i.ToInMageProtectedDiskDetailsResponseOutputWithContext(context.Background())
}

func (i InMageProtectedDiskDetailsResponseArgs) ToInMageProtectedDiskDetailsResponseOutputWithContext(ctx context.Context) InMageProtectedDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageProtectedDiskDetailsResponseOutput)
}





type InMageProtectedDiskDetailsResponseArrayInput interface {
	pulumi.Input

	ToInMageProtectedDiskDetailsResponseArrayOutput() InMageProtectedDiskDetailsResponseArrayOutput
	ToInMageProtectedDiskDetailsResponseArrayOutputWithContext(context.Context) InMageProtectedDiskDetailsResponseArrayOutput
}

type InMageProtectedDiskDetailsResponseArray []InMageProtectedDiskDetailsResponseInput

func (InMageProtectedDiskDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InMageProtectedDiskDetailsResponse)(nil)).Elem()
}

func (i InMageProtectedDiskDetailsResponseArray) ToInMageProtectedDiskDetailsResponseArrayOutput() InMageProtectedDiskDetailsResponseArrayOutput {
	return i.ToInMageProtectedDiskDetailsResponseArrayOutputWithContext(context.Background())
}

func (i InMageProtectedDiskDetailsResponseArray) ToInMageProtectedDiskDetailsResponseArrayOutputWithContext(ctx context.Context) InMageProtectedDiskDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageProtectedDiskDetailsResponseArrayOutput)
}

type InMageProtectedDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (InMageProtectedDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageProtectedDiskDetailsResponse)(nil)).Elem()
}

func (o InMageProtectedDiskDetailsResponseOutput) ToInMageProtectedDiskDetailsResponseOutput() InMageProtectedDiskDetailsResponseOutput {
	return o
}

func (o InMageProtectedDiskDetailsResponseOutput) ToInMageProtectedDiskDetailsResponseOutputWithContext(ctx context.Context) InMageProtectedDiskDetailsResponseOutput {
	return o
}

func (o InMageProtectedDiskDetailsResponseOutput) DiskCapacityInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *float64 { return v.DiskCapacityInBytes }).(pulumi.Float64PtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) DiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *string { return v.DiskId }).(pulumi.StringPtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) DiskName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *string { return v.DiskName }).(pulumi.StringPtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) DiskResized() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *string { return v.DiskResized }).(pulumi.StringPtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) FileSystemCapacityInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *float64 { return v.FileSystemCapacityInBytes }).(pulumi.Float64PtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) HealthErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *string { return v.HealthErrorCode }).(pulumi.StringPtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) LastRpoCalculatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *string { return v.LastRpoCalculatedTime }).(pulumi.StringPtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) ProtectionStage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *string { return v.ProtectionStage }).(pulumi.StringPtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) PsDataInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *float64 { return v.PsDataInMB }).(pulumi.Float64PtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) ResyncDurationInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *float64 { return v.ResyncDurationInSeconds }).(pulumi.Float64PtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) ResyncProgressPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *int { return v.ResyncProgressPercentage }).(pulumi.IntPtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) ResyncRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *string { return v.ResyncRequired }).(pulumi.StringPtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) RpoInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *float64 { return v.RpoInSeconds }).(pulumi.Float64PtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) SourceDataInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *float64 { return v.SourceDataInMB }).(pulumi.Float64PtrOutput)
}

func (o InMageProtectedDiskDetailsResponseOutput) TargetDataInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageProtectedDiskDetailsResponse) *float64 { return v.TargetDataInMB }).(pulumi.Float64PtrOutput)
}

type InMageProtectedDiskDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (InMageProtectedDiskDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InMageProtectedDiskDetailsResponse)(nil)).Elem()
}

func (o InMageProtectedDiskDetailsResponseArrayOutput) ToInMageProtectedDiskDetailsResponseArrayOutput() InMageProtectedDiskDetailsResponseArrayOutput {
	return o
}

func (o InMageProtectedDiskDetailsResponseArrayOutput) ToInMageProtectedDiskDetailsResponseArrayOutputWithContext(ctx context.Context) InMageProtectedDiskDetailsResponseArrayOutput {
	return o
}

func (o InMageProtectedDiskDetailsResponseArrayOutput) Index(i pulumi.IntInput) InMageProtectedDiskDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InMageProtectedDiskDetailsResponse {
		return vs[0].([]InMageProtectedDiskDetailsResponse)[vs[1].(int)]
	}).(InMageProtectedDiskDetailsResponseOutput)
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





type InMageReplicationDetailsResponseInput interface {
	pulumi.Input

	ToInMageReplicationDetailsResponseOutput() InMageReplicationDetailsResponseOutput
	ToInMageReplicationDetailsResponseOutputWithContext(context.Context) InMageReplicationDetailsResponseOutput
}

type InMageReplicationDetailsResponseArgs struct {
	ActiveSiteType               pulumi.StringPtrInput                        `pulumi:"activeSiteType"`
	AgentDetails                 InMageAgentDetailsResponsePtrInput           `pulumi:"agentDetails"`
	AzureStorageAccountId        pulumi.StringPtrInput                        `pulumi:"azureStorageAccountId"`
	CompressedDataRateInMB       pulumi.Float64PtrInput                       `pulumi:"compressedDataRateInMB"`
	ConsistencyPoints            pulumi.StringMapInput                        `pulumi:"consistencyPoints"`
	Datastores                   pulumi.StringArrayInput                      `pulumi:"datastores"`
	DiscoveryType                pulumi.StringPtrInput                        `pulumi:"discoveryType"`
	DiskResized                  pulumi.StringPtrInput                        `pulumi:"diskResized"`
	InfrastructureVmId           pulumi.StringPtrInput                        `pulumi:"infrastructureVmId"`
	InstanceType                 pulumi.StringInput                           `pulumi:"instanceType"`
	IpAddress                    pulumi.StringPtrInput                        `pulumi:"ipAddress"`
	LastHeartbeat                pulumi.StringPtrInput                        `pulumi:"lastHeartbeat"`
	LastRpoCalculatedTime        pulumi.StringPtrInput                        `pulumi:"lastRpoCalculatedTime"`
	LastUpdateReceivedTime       pulumi.StringPtrInput                        `pulumi:"lastUpdateReceivedTime"`
	MasterTargetId               pulumi.StringPtrInput                        `pulumi:"masterTargetId"`
	MultiVmGroupId               pulumi.StringPtrInput                        `pulumi:"multiVmGroupId"`
	MultiVmGroupName             pulumi.StringPtrInput                        `pulumi:"multiVmGroupName"`
	MultiVmSyncStatus            pulumi.StringPtrInput                        `pulumi:"multiVmSyncStatus"`
	OsDetails                    OSDiskDetailsResponsePtrInput                `pulumi:"osDetails"`
	OsVersion                    pulumi.StringPtrInput                        `pulumi:"osVersion"`
	ProcessServerId              pulumi.StringPtrInput                        `pulumi:"processServerId"`
	ProtectedDisks               InMageProtectedDiskDetailsResponseArrayInput `pulumi:"protectedDisks"`
	ProtectionStage              pulumi.StringPtrInput                        `pulumi:"protectionStage"`
	RebootAfterUpdateStatus      pulumi.StringPtrInput                        `pulumi:"rebootAfterUpdateStatus"`
	ReplicaId                    pulumi.StringPtrInput                        `pulumi:"replicaId"`
	ResyncDetails                InitialReplicationDetailsResponsePtrInput    `pulumi:"resyncDetails"`
	RetentionWindowEnd           pulumi.StringPtrInput                        `pulumi:"retentionWindowEnd"`
	RetentionWindowStart         pulumi.StringPtrInput                        `pulumi:"retentionWindowStart"`
	RpoInSeconds                 pulumi.Float64PtrInput                       `pulumi:"rpoInSeconds"`
	SourceVmCpuCount             pulumi.IntPtrInput                           `pulumi:"sourceVmCpuCount"`
	SourceVmRamSizeInMB          pulumi.IntPtrInput                           `pulumi:"sourceVmRamSizeInMB"`
	UncompressedDataRateInMB     pulumi.Float64PtrInput                       `pulumi:"uncompressedDataRateInMB"`
	VCenterInfrastructureId      pulumi.StringPtrInput                        `pulumi:"vCenterInfrastructureId"`
	ValidationErrors             HealthErrorResponseArrayInput                `pulumi:"validationErrors"`
	VmId                         pulumi.StringPtrInput                        `pulumi:"vmId"`
	VmNics                       VMNicDetailsResponseArrayInput               `pulumi:"vmNics"`
	VmProtectionState            pulumi.StringPtrInput                        `pulumi:"vmProtectionState"`
	VmProtectionStateDescription pulumi.StringPtrInput                        `pulumi:"vmProtectionStateDescription"`
}

func (InMageReplicationDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageReplicationDetailsResponse)(nil)).Elem()
}

func (i InMageReplicationDetailsResponseArgs) ToInMageReplicationDetailsResponseOutput() InMageReplicationDetailsResponseOutput {
	return i.ToInMageReplicationDetailsResponseOutputWithContext(context.Background())
}

func (i InMageReplicationDetailsResponseArgs) ToInMageReplicationDetailsResponseOutputWithContext(ctx context.Context) InMageReplicationDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageReplicationDetailsResponseOutput)
}

type InMageReplicationDetailsResponseOutput struct{ *pulumi.OutputState }

func (InMageReplicationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageReplicationDetailsResponse)(nil)).Elem()
}

func (o InMageReplicationDetailsResponseOutput) ToInMageReplicationDetailsResponseOutput() InMageReplicationDetailsResponseOutput {
	return o
}

func (o InMageReplicationDetailsResponseOutput) ToInMageReplicationDetailsResponseOutputWithContext(ctx context.Context) InMageReplicationDetailsResponseOutput {
	return o
}

func (o InMageReplicationDetailsResponseOutput) ActiveSiteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.ActiveSiteType }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) AgentDetails() InMageAgentDetailsResponsePtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *InMageAgentDetailsResponse { return v.AgentDetails }).(InMageAgentDetailsResponsePtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) AzureStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.AzureStorageAccountId }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) CompressedDataRateInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *float64 { return v.CompressedDataRateInMB }).(pulumi.Float64PtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) ConsistencyPoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) map[string]string { return v.ConsistencyPoints }).(pulumi.StringMapOutput)
}

func (o InMageReplicationDetailsResponseOutput) Datastores() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) []string { return v.Datastores }).(pulumi.StringArrayOutput)
}

func (o InMageReplicationDetailsResponseOutput) DiscoveryType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.DiscoveryType }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) DiskResized() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.DiskResized }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) InfrastructureVmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.InfrastructureVmId }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o InMageReplicationDetailsResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) LastHeartbeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.LastHeartbeat }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) LastRpoCalculatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.LastRpoCalculatedTime }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) LastUpdateReceivedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.LastUpdateReceivedTime }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) MasterTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.MasterTargetId }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) MultiVmGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.MultiVmGroupId }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) MultiVmGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.MultiVmGroupName }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) MultiVmSyncStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.MultiVmSyncStatus }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) OsDetails() OSDiskDetailsResponsePtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *OSDiskDetailsResponse { return v.OsDetails }).(OSDiskDetailsResponsePtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.OsVersion }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) ProcessServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.ProcessServerId }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) ProtectedDisks() InMageProtectedDiskDetailsResponseArrayOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) []InMageProtectedDiskDetailsResponse { return v.ProtectedDisks }).(InMageProtectedDiskDetailsResponseArrayOutput)
}

func (o InMageReplicationDetailsResponseOutput) ProtectionStage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.ProtectionStage }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) RebootAfterUpdateStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.RebootAfterUpdateStatus }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) ReplicaId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.ReplicaId }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) ResyncDetails() InitialReplicationDetailsResponsePtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *InitialReplicationDetailsResponse { return v.ResyncDetails }).(InitialReplicationDetailsResponsePtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) RetentionWindowEnd() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.RetentionWindowEnd }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) RetentionWindowStart() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.RetentionWindowStart }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) RpoInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *float64 { return v.RpoInSeconds }).(pulumi.Float64PtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) SourceVmCpuCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *int { return v.SourceVmCpuCount }).(pulumi.IntPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) SourceVmRamSizeInMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *int { return v.SourceVmRamSizeInMB }).(pulumi.IntPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) UncompressedDataRateInMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *float64 { return v.UncompressedDataRateInMB }).(pulumi.Float64PtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) VCenterInfrastructureId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.VCenterInfrastructureId }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) ValidationErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) []HealthErrorResponse { return v.ValidationErrors }).(HealthErrorResponseArrayOutput)
}

func (o InMageReplicationDetailsResponseOutput) VmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.VmId }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) VmNics() VMNicDetailsResponseArrayOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) []VMNicDetailsResponse { return v.VmNics }).(VMNicDetailsResponseArrayOutput)
}

func (o InMageReplicationDetailsResponseOutput) VmProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.VmProtectionState }).(pulumi.StringPtrOutput)
}

func (o InMageReplicationDetailsResponseOutput) VmProtectionStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageReplicationDetailsResponse) *string { return v.VmProtectionStateDescription }).(pulumi.StringPtrOutput)
}

type InMageVolumeExclusionOptions struct {
	OnlyExcludeIfSingleVolume *string `pulumi:"onlyExcludeIfSingleVolume"`
	VolumeLabel               *string `pulumi:"volumeLabel"`
}





type InMageVolumeExclusionOptionsInput interface {
	pulumi.Input

	ToInMageVolumeExclusionOptionsOutput() InMageVolumeExclusionOptionsOutput
	ToInMageVolumeExclusionOptionsOutputWithContext(context.Context) InMageVolumeExclusionOptionsOutput
}

type InMageVolumeExclusionOptionsArgs struct {
	OnlyExcludeIfSingleVolume pulumi.StringPtrInput `pulumi:"onlyExcludeIfSingleVolume"`
	VolumeLabel               pulumi.StringPtrInput `pulumi:"volumeLabel"`
}

func (InMageVolumeExclusionOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageVolumeExclusionOptions)(nil)).Elem()
}

func (i InMageVolumeExclusionOptionsArgs) ToInMageVolumeExclusionOptionsOutput() InMageVolumeExclusionOptionsOutput {
	return i.ToInMageVolumeExclusionOptionsOutputWithContext(context.Background())
}

func (i InMageVolumeExclusionOptionsArgs) ToInMageVolumeExclusionOptionsOutputWithContext(ctx context.Context) InMageVolumeExclusionOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageVolumeExclusionOptionsOutput)
}





type InMageVolumeExclusionOptionsArrayInput interface {
	pulumi.Input

	ToInMageVolumeExclusionOptionsArrayOutput() InMageVolumeExclusionOptionsArrayOutput
	ToInMageVolumeExclusionOptionsArrayOutputWithContext(context.Context) InMageVolumeExclusionOptionsArrayOutput
}

type InMageVolumeExclusionOptionsArray []InMageVolumeExclusionOptionsInput

func (InMageVolumeExclusionOptionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InMageVolumeExclusionOptions)(nil)).Elem()
}

func (i InMageVolumeExclusionOptionsArray) ToInMageVolumeExclusionOptionsArrayOutput() InMageVolumeExclusionOptionsArrayOutput {
	return i.ToInMageVolumeExclusionOptionsArrayOutputWithContext(context.Background())
}

func (i InMageVolumeExclusionOptionsArray) ToInMageVolumeExclusionOptionsArrayOutputWithContext(ctx context.Context) InMageVolumeExclusionOptionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InMageVolumeExclusionOptionsArrayOutput)
}

type InMageVolumeExclusionOptionsOutput struct{ *pulumi.OutputState }

func (InMageVolumeExclusionOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InMageVolumeExclusionOptions)(nil)).Elem()
}

func (o InMageVolumeExclusionOptionsOutput) ToInMageVolumeExclusionOptionsOutput() InMageVolumeExclusionOptionsOutput {
	return o
}

func (o InMageVolumeExclusionOptionsOutput) ToInMageVolumeExclusionOptionsOutputWithContext(ctx context.Context) InMageVolumeExclusionOptionsOutput {
	return o
}

func (o InMageVolumeExclusionOptionsOutput) OnlyExcludeIfSingleVolume() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageVolumeExclusionOptions) *string { return v.OnlyExcludeIfSingleVolume }).(pulumi.StringPtrOutput)
}

func (o InMageVolumeExclusionOptionsOutput) VolumeLabel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InMageVolumeExclusionOptions) *string { return v.VolumeLabel }).(pulumi.StringPtrOutput)
}

type InMageVolumeExclusionOptionsArrayOutput struct{ *pulumi.OutputState }

func (InMageVolumeExclusionOptionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InMageVolumeExclusionOptions)(nil)).Elem()
}

func (o InMageVolumeExclusionOptionsArrayOutput) ToInMageVolumeExclusionOptionsArrayOutput() InMageVolumeExclusionOptionsArrayOutput {
	return o
}

func (o InMageVolumeExclusionOptionsArrayOutput) ToInMageVolumeExclusionOptionsArrayOutputWithContext(ctx context.Context) InMageVolumeExclusionOptionsArrayOutput {
	return o
}

func (o InMageVolumeExclusionOptionsArrayOutput) Index(i pulumi.IntInput) InMageVolumeExclusionOptionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InMageVolumeExclusionOptions {
		return vs[0].([]InMageVolumeExclusionOptions)[vs[1].(int)]
	}).(InMageVolumeExclusionOptionsOutput)
}

type InitialReplicationDetailsResponse struct {
	InitialReplicationProgressPercentage *string `pulumi:"initialReplicationProgressPercentage"`
	InitialReplicationType               *string `pulumi:"initialReplicationType"`
}





type InitialReplicationDetailsResponseInput interface {
	pulumi.Input

	ToInitialReplicationDetailsResponseOutput() InitialReplicationDetailsResponseOutput
	ToInitialReplicationDetailsResponseOutputWithContext(context.Context) InitialReplicationDetailsResponseOutput
}

type InitialReplicationDetailsResponseArgs struct {
	InitialReplicationProgressPercentage pulumi.StringPtrInput `pulumi:"initialReplicationProgressPercentage"`
	InitialReplicationType               pulumi.StringPtrInput `pulumi:"initialReplicationType"`
}

func (InitialReplicationDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InitialReplicationDetailsResponse)(nil)).Elem()
}

func (i InitialReplicationDetailsResponseArgs) ToInitialReplicationDetailsResponseOutput() InitialReplicationDetailsResponseOutput {
	return i.ToInitialReplicationDetailsResponseOutputWithContext(context.Background())
}

func (i InitialReplicationDetailsResponseArgs) ToInitialReplicationDetailsResponseOutputWithContext(ctx context.Context) InitialReplicationDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InitialReplicationDetailsResponseOutput)
}

func (i InitialReplicationDetailsResponseArgs) ToInitialReplicationDetailsResponsePtrOutput() InitialReplicationDetailsResponsePtrOutput {
	return i.ToInitialReplicationDetailsResponsePtrOutputWithContext(context.Background())
}

func (i InitialReplicationDetailsResponseArgs) ToInitialReplicationDetailsResponsePtrOutputWithContext(ctx context.Context) InitialReplicationDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InitialReplicationDetailsResponseOutput).ToInitialReplicationDetailsResponsePtrOutputWithContext(ctx)
}









type InitialReplicationDetailsResponsePtrInput interface {
	pulumi.Input

	ToInitialReplicationDetailsResponsePtrOutput() InitialReplicationDetailsResponsePtrOutput
	ToInitialReplicationDetailsResponsePtrOutputWithContext(context.Context) InitialReplicationDetailsResponsePtrOutput
}

type initialReplicationDetailsResponsePtrType InitialReplicationDetailsResponseArgs

func InitialReplicationDetailsResponsePtr(v *InitialReplicationDetailsResponseArgs) InitialReplicationDetailsResponsePtrInput {
	return (*initialReplicationDetailsResponsePtrType)(v)
}

func (*initialReplicationDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InitialReplicationDetailsResponse)(nil)).Elem()
}

func (i *initialReplicationDetailsResponsePtrType) ToInitialReplicationDetailsResponsePtrOutput() InitialReplicationDetailsResponsePtrOutput {
	return i.ToInitialReplicationDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *initialReplicationDetailsResponsePtrType) ToInitialReplicationDetailsResponsePtrOutputWithContext(ctx context.Context) InitialReplicationDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InitialReplicationDetailsResponsePtrOutput)
}

type InitialReplicationDetailsResponseOutput struct{ *pulumi.OutputState }

func (InitialReplicationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InitialReplicationDetailsResponse)(nil)).Elem()
}

func (o InitialReplicationDetailsResponseOutput) ToInitialReplicationDetailsResponseOutput() InitialReplicationDetailsResponseOutput {
	return o
}

func (o InitialReplicationDetailsResponseOutput) ToInitialReplicationDetailsResponseOutputWithContext(ctx context.Context) InitialReplicationDetailsResponseOutput {
	return o
}

func (o InitialReplicationDetailsResponseOutput) ToInitialReplicationDetailsResponsePtrOutput() InitialReplicationDetailsResponsePtrOutput {
	return o.ToInitialReplicationDetailsResponsePtrOutputWithContext(context.Background())
}

func (o InitialReplicationDetailsResponseOutput) ToInitialReplicationDetailsResponsePtrOutputWithContext(ctx context.Context) InitialReplicationDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InitialReplicationDetailsResponse) *InitialReplicationDetailsResponse {
		return &v
	}).(InitialReplicationDetailsResponsePtrOutput)
}

func (o InitialReplicationDetailsResponseOutput) InitialReplicationProgressPercentage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InitialReplicationDetailsResponse) *string { return v.InitialReplicationProgressPercentage }).(pulumi.StringPtrOutput)
}

func (o InitialReplicationDetailsResponseOutput) InitialReplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InitialReplicationDetailsResponse) *string { return v.InitialReplicationType }).(pulumi.StringPtrOutput)
}

type InitialReplicationDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (InitialReplicationDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InitialReplicationDetailsResponse)(nil)).Elem()
}

func (o InitialReplicationDetailsResponsePtrOutput) ToInitialReplicationDetailsResponsePtrOutput() InitialReplicationDetailsResponsePtrOutput {
	return o
}

func (o InitialReplicationDetailsResponsePtrOutput) ToInitialReplicationDetailsResponsePtrOutputWithContext(ctx context.Context) InitialReplicationDetailsResponsePtrOutput {
	return o
}

func (o InitialReplicationDetailsResponsePtrOutput) Elem() InitialReplicationDetailsResponseOutput {
	return o.ApplyT(func(v *InitialReplicationDetailsResponse) InitialReplicationDetailsResponse {
		if v != nil {
			return *v
		}
		var ret InitialReplicationDetailsResponse
		return ret
	}).(InitialReplicationDetailsResponseOutput)
}

func (o InitialReplicationDetailsResponsePtrOutput) InitialReplicationProgressPercentage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InitialReplicationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InitialReplicationProgressPercentage
	}).(pulumi.StringPtrOutput)
}

func (o InitialReplicationDetailsResponsePtrOutput) InitialReplicationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InitialReplicationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.InitialReplicationType
	}).(pulumi.StringPtrOutput)
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





type InnerHealthErrorResponseInput interface {
	pulumi.Input

	ToInnerHealthErrorResponseOutput() InnerHealthErrorResponseOutput
	ToInnerHealthErrorResponseOutputWithContext(context.Context) InnerHealthErrorResponseOutput
}

type InnerHealthErrorResponseArgs struct {
	CreationTimeUtc              pulumi.StringPtrInput `pulumi:"creationTimeUtc"`
	EntityId                     pulumi.StringPtrInput `pulumi:"entityId"`
	ErrorCategory                pulumi.StringPtrInput `pulumi:"errorCategory"`
	ErrorCode                    pulumi.StringPtrInput `pulumi:"errorCode"`
	ErrorLevel                   pulumi.StringPtrInput `pulumi:"errorLevel"`
	ErrorMessage                 pulumi.StringPtrInput `pulumi:"errorMessage"`
	ErrorSource                  pulumi.StringPtrInput `pulumi:"errorSource"`
	ErrorType                    pulumi.StringPtrInput `pulumi:"errorType"`
	PossibleCauses               pulumi.StringPtrInput `pulumi:"possibleCauses"`
	RecommendedAction            pulumi.StringPtrInput `pulumi:"recommendedAction"`
	RecoveryProviderErrorMessage pulumi.StringPtrInput `pulumi:"recoveryProviderErrorMessage"`
	SummaryMessage               pulumi.StringPtrInput `pulumi:"summaryMessage"`
}

func (InnerHealthErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InnerHealthErrorResponse)(nil)).Elem()
}

func (i InnerHealthErrorResponseArgs) ToInnerHealthErrorResponseOutput() InnerHealthErrorResponseOutput {
	return i.ToInnerHealthErrorResponseOutputWithContext(context.Background())
}

func (i InnerHealthErrorResponseArgs) ToInnerHealthErrorResponseOutputWithContext(ctx context.Context) InnerHealthErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InnerHealthErrorResponseOutput)
}





type InnerHealthErrorResponseArrayInput interface {
	pulumi.Input

	ToInnerHealthErrorResponseArrayOutput() InnerHealthErrorResponseArrayOutput
	ToInnerHealthErrorResponseArrayOutputWithContext(context.Context) InnerHealthErrorResponseArrayOutput
}

type InnerHealthErrorResponseArray []InnerHealthErrorResponseInput

func (InnerHealthErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InnerHealthErrorResponse)(nil)).Elem()
}

func (i InnerHealthErrorResponseArray) ToInnerHealthErrorResponseArrayOutput() InnerHealthErrorResponseArrayOutput {
	return i.ToInnerHealthErrorResponseArrayOutputWithContext(context.Background())
}

func (i InnerHealthErrorResponseArray) ToInnerHealthErrorResponseArrayOutputWithContext(ctx context.Context) InnerHealthErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InnerHealthErrorResponseArrayOutput)
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





type InputEndpointResponseInput interface {
	pulumi.Input

	ToInputEndpointResponseOutput() InputEndpointResponseOutput
	ToInputEndpointResponseOutputWithContext(context.Context) InputEndpointResponseOutput
}

type InputEndpointResponseArgs struct {
	EndpointName pulumi.StringPtrInput `pulumi:"endpointName"`
	PrivatePort  pulumi.IntPtrInput    `pulumi:"privatePort"`
	Protocol     pulumi.StringPtrInput `pulumi:"protocol"`
	PublicPort   pulumi.IntPtrInput    `pulumi:"publicPort"`
}

func (InputEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InputEndpointResponse)(nil)).Elem()
}

func (i InputEndpointResponseArgs) ToInputEndpointResponseOutput() InputEndpointResponseOutput {
	return i.ToInputEndpointResponseOutputWithContext(context.Background())
}

func (i InputEndpointResponseArgs) ToInputEndpointResponseOutputWithContext(ctx context.Context) InputEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputEndpointResponseOutput)
}





type InputEndpointResponseArrayInput interface {
	pulumi.Input

	ToInputEndpointResponseArrayOutput() InputEndpointResponseArrayOutput
	ToInputEndpointResponseArrayOutputWithContext(context.Context) InputEndpointResponseArrayOutput
}

type InputEndpointResponseArray []InputEndpointResponseInput

func (InputEndpointResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputEndpointResponse)(nil)).Elem()
}

func (i InputEndpointResponseArray) ToInputEndpointResponseArrayOutput() InputEndpointResponseArrayOutput {
	return i.ToInputEndpointResponseArrayOutputWithContext(context.Background())
}

func (i InputEndpointResponseArray) ToInputEndpointResponseArrayOutputWithContext(ctx context.Context) InputEndpointResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InputEndpointResponseArrayOutput)
}

type InputEndpointResponseOutput struct{ *pulumi.OutputState }

func (InputEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputEndpointResponse)(nil)).Elem()
}

func (o InputEndpointResponseOutput) ToInputEndpointResponseOutput() InputEndpointResponseOutput {
	return o
}

func (o InputEndpointResponseOutput) ToInputEndpointResponseOutputWithContext(ctx context.Context) InputEndpointResponseOutput {
	return o
}

func (o InputEndpointResponseOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputEndpointResponse) *string { return v.EndpointName }).(pulumi.StringPtrOutput)
}

func (o InputEndpointResponseOutput) PrivatePort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InputEndpointResponse) *int { return v.PrivatePort }).(pulumi.IntPtrOutput)
}

func (o InputEndpointResponseOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InputEndpointResponse) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o InputEndpointResponseOutput) PublicPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v InputEndpointResponse) *int { return v.PublicPort }).(pulumi.IntPtrOutput)
}

type InputEndpointResponseArrayOutput struct{ *pulumi.OutputState }

func (InputEndpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InputEndpointResponse)(nil)).Elem()
}

func (o InputEndpointResponseArrayOutput) ToInputEndpointResponseArrayOutput() InputEndpointResponseArrayOutput {
	return o
}

func (o InputEndpointResponseArrayOutput) ToInputEndpointResponseArrayOutputWithContext(ctx context.Context) InputEndpointResponseArrayOutput {
	return o
}

func (o InputEndpointResponseArrayOutput) Index(i pulumi.IntInput) InputEndpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InputEndpointResponse {
		return vs[0].([]InputEndpointResponse)[vs[1].(int)]
	}).(InputEndpointResponseOutput)
}

type KeyEncryptionKeyInfo struct {
	KeyIdentifier         *string `pulumi:"keyIdentifier"`
	KeyVaultResourceArmId *string `pulumi:"keyVaultResourceArmId"`
}





type KeyEncryptionKeyInfoInput interface {
	pulumi.Input

	ToKeyEncryptionKeyInfoOutput() KeyEncryptionKeyInfoOutput
	ToKeyEncryptionKeyInfoOutputWithContext(context.Context) KeyEncryptionKeyInfoOutput
}

type KeyEncryptionKeyInfoArgs struct {
	KeyIdentifier         pulumi.StringPtrInput `pulumi:"keyIdentifier"`
	KeyVaultResourceArmId pulumi.StringPtrInput `pulumi:"keyVaultResourceArmId"`
}

func (KeyEncryptionKeyInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyEncryptionKeyInfo)(nil)).Elem()
}

func (i KeyEncryptionKeyInfoArgs) ToKeyEncryptionKeyInfoOutput() KeyEncryptionKeyInfoOutput {
	return i.ToKeyEncryptionKeyInfoOutputWithContext(context.Background())
}

func (i KeyEncryptionKeyInfoArgs) ToKeyEncryptionKeyInfoOutputWithContext(ctx context.Context) KeyEncryptionKeyInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyEncryptionKeyInfoOutput)
}

func (i KeyEncryptionKeyInfoArgs) ToKeyEncryptionKeyInfoPtrOutput() KeyEncryptionKeyInfoPtrOutput {
	return i.ToKeyEncryptionKeyInfoPtrOutputWithContext(context.Background())
}

func (i KeyEncryptionKeyInfoArgs) ToKeyEncryptionKeyInfoPtrOutputWithContext(ctx context.Context) KeyEncryptionKeyInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyEncryptionKeyInfoOutput).ToKeyEncryptionKeyInfoPtrOutputWithContext(ctx)
}









type KeyEncryptionKeyInfoPtrInput interface {
	pulumi.Input

	ToKeyEncryptionKeyInfoPtrOutput() KeyEncryptionKeyInfoPtrOutput
	ToKeyEncryptionKeyInfoPtrOutputWithContext(context.Context) KeyEncryptionKeyInfoPtrOutput
}

type keyEncryptionKeyInfoPtrType KeyEncryptionKeyInfoArgs

func KeyEncryptionKeyInfoPtr(v *KeyEncryptionKeyInfoArgs) KeyEncryptionKeyInfoPtrInput {
	return (*keyEncryptionKeyInfoPtrType)(v)
}

func (*keyEncryptionKeyInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyEncryptionKeyInfo)(nil)).Elem()
}

func (i *keyEncryptionKeyInfoPtrType) ToKeyEncryptionKeyInfoPtrOutput() KeyEncryptionKeyInfoPtrOutput {
	return i.ToKeyEncryptionKeyInfoPtrOutputWithContext(context.Background())
}

func (i *keyEncryptionKeyInfoPtrType) ToKeyEncryptionKeyInfoPtrOutputWithContext(ctx context.Context) KeyEncryptionKeyInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyEncryptionKeyInfoPtrOutput)
}

type KeyEncryptionKeyInfoOutput struct{ *pulumi.OutputState }

func (KeyEncryptionKeyInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyEncryptionKeyInfo)(nil)).Elem()
}

func (o KeyEncryptionKeyInfoOutput) ToKeyEncryptionKeyInfoOutput() KeyEncryptionKeyInfoOutput {
	return o
}

func (o KeyEncryptionKeyInfoOutput) ToKeyEncryptionKeyInfoOutputWithContext(ctx context.Context) KeyEncryptionKeyInfoOutput {
	return o
}

func (o KeyEncryptionKeyInfoOutput) ToKeyEncryptionKeyInfoPtrOutput() KeyEncryptionKeyInfoPtrOutput {
	return o.ToKeyEncryptionKeyInfoPtrOutputWithContext(context.Background())
}

func (o KeyEncryptionKeyInfoOutput) ToKeyEncryptionKeyInfoPtrOutputWithContext(ctx context.Context) KeyEncryptionKeyInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyEncryptionKeyInfo) *KeyEncryptionKeyInfo {
		return &v
	}).(KeyEncryptionKeyInfoPtrOutput)
}

func (o KeyEncryptionKeyInfoOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyEncryptionKeyInfo) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
}

func (o KeyEncryptionKeyInfoOutput) KeyVaultResourceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyEncryptionKeyInfo) *string { return v.KeyVaultResourceArmId }).(pulumi.StringPtrOutput)
}

type KeyEncryptionKeyInfoPtrOutput struct{ *pulumi.OutputState }

func (KeyEncryptionKeyInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyEncryptionKeyInfo)(nil)).Elem()
}

func (o KeyEncryptionKeyInfoPtrOutput) ToKeyEncryptionKeyInfoPtrOutput() KeyEncryptionKeyInfoPtrOutput {
	return o
}

func (o KeyEncryptionKeyInfoPtrOutput) ToKeyEncryptionKeyInfoPtrOutputWithContext(ctx context.Context) KeyEncryptionKeyInfoPtrOutput {
	return o
}

func (o KeyEncryptionKeyInfoPtrOutput) Elem() KeyEncryptionKeyInfoOutput {
	return o.ApplyT(func(v *KeyEncryptionKeyInfo) KeyEncryptionKeyInfo {
		if v != nil {
			return *v
		}
		var ret KeyEncryptionKeyInfo
		return ret
	}).(KeyEncryptionKeyInfoOutput)
}

func (o KeyEncryptionKeyInfoPtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyEncryptionKeyInfo) *string {
		if v == nil {
			return nil
		}
		return v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o KeyEncryptionKeyInfoPtrOutput) KeyVaultResourceArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyEncryptionKeyInfo) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultResourceArmId
	}).(pulumi.StringPtrOutput)
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





type MasterTargetServerResponseInput interface {
	pulumi.Input

	ToMasterTargetServerResponseOutput() MasterTargetServerResponseOutput
	ToMasterTargetServerResponseOutputWithContext(context.Context) MasterTargetServerResponseOutput
}

type MasterTargetServerResponseArgs struct {
	AgentExpiryDate         pulumi.StringPtrInput             `pulumi:"agentExpiryDate"`
	AgentVersion            pulumi.StringPtrInput             `pulumi:"agentVersion"`
	AgentVersionDetails     VersionDetailsResponsePtrInput    `pulumi:"agentVersionDetails"`
	DataStores              DataStoreResponseArrayInput       `pulumi:"dataStores"`
	DiskCount               pulumi.IntPtrInput                `pulumi:"diskCount"`
	HealthErrors            HealthErrorResponseArrayInput     `pulumi:"healthErrors"`
	Id                      pulumi.StringPtrInput             `pulumi:"id"`
	IpAddress               pulumi.StringPtrInput             `pulumi:"ipAddress"`
	LastHeartbeat           pulumi.StringPtrInput             `pulumi:"lastHeartbeat"`
	MarsAgentExpiryDate     pulumi.StringPtrInput             `pulumi:"marsAgentExpiryDate"`
	MarsAgentVersion        pulumi.StringPtrInput             `pulumi:"marsAgentVersion"`
	MarsAgentVersionDetails VersionDetailsResponsePtrInput    `pulumi:"marsAgentVersionDetails"`
	Name                    pulumi.StringPtrInput             `pulumi:"name"`
	OsType                  pulumi.StringPtrInput             `pulumi:"osType"`
	OsVersion               pulumi.StringPtrInput             `pulumi:"osVersion"`
	RetentionVolumes        RetentionVolumeResponseArrayInput `pulumi:"retentionVolumes"`
	ValidationErrors        HealthErrorResponseArrayInput     `pulumi:"validationErrors"`
	VersionStatus           pulumi.StringPtrInput             `pulumi:"versionStatus"`
}

func (MasterTargetServerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterTargetServerResponse)(nil)).Elem()
}

func (i MasterTargetServerResponseArgs) ToMasterTargetServerResponseOutput() MasterTargetServerResponseOutput {
	return i.ToMasterTargetServerResponseOutputWithContext(context.Background())
}

func (i MasterTargetServerResponseArgs) ToMasterTargetServerResponseOutputWithContext(ctx context.Context) MasterTargetServerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterTargetServerResponseOutput)
}





type MasterTargetServerResponseArrayInput interface {
	pulumi.Input

	ToMasterTargetServerResponseArrayOutput() MasterTargetServerResponseArrayOutput
	ToMasterTargetServerResponseArrayOutputWithContext(context.Context) MasterTargetServerResponseArrayOutput
}

type MasterTargetServerResponseArray []MasterTargetServerResponseInput

func (MasterTargetServerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MasterTargetServerResponse)(nil)).Elem()
}

func (i MasterTargetServerResponseArray) ToMasterTargetServerResponseArrayOutput() MasterTargetServerResponseArrayOutput {
	return i.ToMasterTargetServerResponseArrayOutputWithContext(context.Background())
}

func (i MasterTargetServerResponseArray) ToMasterTargetServerResponseArrayOutputWithContext(ctx context.Context) MasterTargetServerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MasterTargetServerResponseArrayOutput)
}

type MasterTargetServerResponseOutput struct{ *pulumi.OutputState }

func (MasterTargetServerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MasterTargetServerResponse)(nil)).Elem()
}

func (o MasterTargetServerResponseOutput) ToMasterTargetServerResponseOutput() MasterTargetServerResponseOutput {
	return o
}

func (o MasterTargetServerResponseOutput) ToMasterTargetServerResponseOutputWithContext(ctx context.Context) MasterTargetServerResponseOutput {
	return o
}

func (o MasterTargetServerResponseOutput) AgentExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.AgentExpiryDate }).(pulumi.StringPtrOutput)
}

func (o MasterTargetServerResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o MasterTargetServerResponseOutput) AgentVersionDetails() VersionDetailsResponsePtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *VersionDetailsResponse { return v.AgentVersionDetails }).(VersionDetailsResponsePtrOutput)
}

func (o MasterTargetServerResponseOutput) DataStores() DataStoreResponseArrayOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) []DataStoreResponse { return v.DataStores }).(DataStoreResponseArrayOutput)
}

func (o MasterTargetServerResponseOutput) DiskCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *int { return v.DiskCount }).(pulumi.IntPtrOutput)
}

func (o MasterTargetServerResponseOutput) HealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) []HealthErrorResponse { return v.HealthErrors }).(HealthErrorResponseArrayOutput)
}

func (o MasterTargetServerResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o MasterTargetServerResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o MasterTargetServerResponseOutput) LastHeartbeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.LastHeartbeat }).(pulumi.StringPtrOutput)
}

func (o MasterTargetServerResponseOutput) MarsAgentExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.MarsAgentExpiryDate }).(pulumi.StringPtrOutput)
}

func (o MasterTargetServerResponseOutput) MarsAgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.MarsAgentVersion }).(pulumi.StringPtrOutput)
}

func (o MasterTargetServerResponseOutput) MarsAgentVersionDetails() VersionDetailsResponsePtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *VersionDetailsResponse { return v.MarsAgentVersionDetails }).(VersionDetailsResponsePtrOutput)
}

func (o MasterTargetServerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MasterTargetServerResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o MasterTargetServerResponseOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.OsVersion }).(pulumi.StringPtrOutput)
}

func (o MasterTargetServerResponseOutput) RetentionVolumes() RetentionVolumeResponseArrayOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) []RetentionVolumeResponse { return v.RetentionVolumes }).(RetentionVolumeResponseArrayOutput)
}

func (o MasterTargetServerResponseOutput) ValidationErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) []HealthErrorResponse { return v.ValidationErrors }).(HealthErrorResponseArrayOutput)
}

func (o MasterTargetServerResponseOutput) VersionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MasterTargetServerResponse) *string { return v.VersionStatus }).(pulumi.StringPtrOutput)
}

type MasterTargetServerResponseArrayOutput struct{ *pulumi.OutputState }

func (MasterTargetServerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MasterTargetServerResponse)(nil)).Elem()
}

func (o MasterTargetServerResponseArrayOutput) ToMasterTargetServerResponseArrayOutput() MasterTargetServerResponseArrayOutput {
	return o
}

func (o MasterTargetServerResponseArrayOutput) ToMasterTargetServerResponseArrayOutputWithContext(ctx context.Context) MasterTargetServerResponseArrayOutput {
	return o
}

func (o MasterTargetServerResponseArrayOutput) Index(i pulumi.IntInput) MasterTargetServerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MasterTargetServerResponse {
		return vs[0].([]MasterTargetServerResponse)[vs[1].(int)]
	}).(MasterTargetServerResponseOutput)
}

type MigrationItemPropertiesResponse struct {
	AllowedOperations           []string                           `pulumi:"allowedOperations"`
	CurrentJob                  CurrentJobDetailsResponse          `pulumi:"currentJob"`
	EventCorrelationId          string                             `pulumi:"eventCorrelationId"`
	Health                      string                             `pulumi:"health"`
	HealthErrors                []HealthErrorResponse              `pulumi:"healthErrors"`
	LastTestMigrationStatus     string                             `pulumi:"lastTestMigrationStatus"`
	LastTestMigrationTime       string                             `pulumi:"lastTestMigrationTime"`
	MachineName                 string                             `pulumi:"machineName"`
	MigrationState              string                             `pulumi:"migrationState"`
	MigrationStateDescription   string                             `pulumi:"migrationStateDescription"`
	PolicyFriendlyName          string                             `pulumi:"policyFriendlyName"`
	PolicyId                    string                             `pulumi:"policyId"`
	ProviderSpecificDetails     *VMwareCbtMigrationDetailsResponse `pulumi:"providerSpecificDetails"`
	TestMigrateState            string                             `pulumi:"testMigrateState"`
	TestMigrateStateDescription string                             `pulumi:"testMigrateStateDescription"`
}





type MigrationItemPropertiesResponseInput interface {
	pulumi.Input

	ToMigrationItemPropertiesResponseOutput() MigrationItemPropertiesResponseOutput
	ToMigrationItemPropertiesResponseOutputWithContext(context.Context) MigrationItemPropertiesResponseOutput
}

type MigrationItemPropertiesResponseArgs struct {
	AllowedOperations           pulumi.StringArrayInput                   `pulumi:"allowedOperations"`
	CurrentJob                  CurrentJobDetailsResponseInput            `pulumi:"currentJob"`
	EventCorrelationId          pulumi.StringInput                        `pulumi:"eventCorrelationId"`
	Health                      pulumi.StringInput                        `pulumi:"health"`
	HealthErrors                HealthErrorResponseArrayInput             `pulumi:"healthErrors"`
	LastTestMigrationStatus     pulumi.StringInput                        `pulumi:"lastTestMigrationStatus"`
	LastTestMigrationTime       pulumi.StringInput                        `pulumi:"lastTestMigrationTime"`
	MachineName                 pulumi.StringInput                        `pulumi:"machineName"`
	MigrationState              pulumi.StringInput                        `pulumi:"migrationState"`
	MigrationStateDescription   pulumi.StringInput                        `pulumi:"migrationStateDescription"`
	PolicyFriendlyName          pulumi.StringInput                        `pulumi:"policyFriendlyName"`
	PolicyId                    pulumi.StringInput                        `pulumi:"policyId"`
	ProviderSpecificDetails     VMwareCbtMigrationDetailsResponsePtrInput `pulumi:"providerSpecificDetails"`
	TestMigrateState            pulumi.StringInput                        `pulumi:"testMigrateState"`
	TestMigrateStateDescription pulumi.StringInput                        `pulumi:"testMigrateStateDescription"`
}

func (MigrationItemPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MigrationItemPropertiesResponse)(nil)).Elem()
}

func (i MigrationItemPropertiesResponseArgs) ToMigrationItemPropertiesResponseOutput() MigrationItemPropertiesResponseOutput {
	return i.ToMigrationItemPropertiesResponseOutputWithContext(context.Background())
}

func (i MigrationItemPropertiesResponseArgs) ToMigrationItemPropertiesResponseOutputWithContext(ctx context.Context) MigrationItemPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationItemPropertiesResponseOutput)
}

func (i MigrationItemPropertiesResponseArgs) ToMigrationItemPropertiesResponsePtrOutput() MigrationItemPropertiesResponsePtrOutput {
	return i.ToMigrationItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MigrationItemPropertiesResponseArgs) ToMigrationItemPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrationItemPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationItemPropertiesResponseOutput).ToMigrationItemPropertiesResponsePtrOutputWithContext(ctx)
}









type MigrationItemPropertiesResponsePtrInput interface {
	pulumi.Input

	ToMigrationItemPropertiesResponsePtrOutput() MigrationItemPropertiesResponsePtrOutput
	ToMigrationItemPropertiesResponsePtrOutputWithContext(context.Context) MigrationItemPropertiesResponsePtrOutput
}

type migrationItemPropertiesResponsePtrType MigrationItemPropertiesResponseArgs

func MigrationItemPropertiesResponsePtr(v *MigrationItemPropertiesResponseArgs) MigrationItemPropertiesResponsePtrInput {
	return (*migrationItemPropertiesResponsePtrType)(v)
}

func (*migrationItemPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationItemPropertiesResponse)(nil)).Elem()
}

func (i *migrationItemPropertiesResponsePtrType) ToMigrationItemPropertiesResponsePtrOutput() MigrationItemPropertiesResponsePtrOutput {
	return i.ToMigrationItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *migrationItemPropertiesResponsePtrType) ToMigrationItemPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrationItemPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MigrationItemPropertiesResponsePtrOutput)
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

func (o MigrationItemPropertiesResponseOutput) ToMigrationItemPropertiesResponsePtrOutput() MigrationItemPropertiesResponsePtrOutput {
	return o.ToMigrationItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MigrationItemPropertiesResponseOutput) ToMigrationItemPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrationItemPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MigrationItemPropertiesResponse) *MigrationItemPropertiesResponse {
		return &v
	}).(MigrationItemPropertiesResponsePtrOutput)
}

func (o MigrationItemPropertiesResponseOutput) AllowedOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) []string { return v.AllowedOperations }).(pulumi.StringArrayOutput)
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

func (o MigrationItemPropertiesResponseOutput) TestMigrateState() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.TestMigrateState }).(pulumi.StringOutput)
}

func (o MigrationItemPropertiesResponseOutput) TestMigrateStateDescription() pulumi.StringOutput {
	return o.ApplyT(func(v MigrationItemPropertiesResponse) string { return v.TestMigrateStateDescription }).(pulumi.StringOutput)
}

type MigrationItemPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MigrationItemPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MigrationItemPropertiesResponse)(nil)).Elem()
}

func (o MigrationItemPropertiesResponsePtrOutput) ToMigrationItemPropertiesResponsePtrOutput() MigrationItemPropertiesResponsePtrOutput {
	return o
}

func (o MigrationItemPropertiesResponsePtrOutput) ToMigrationItemPropertiesResponsePtrOutputWithContext(ctx context.Context) MigrationItemPropertiesResponsePtrOutput {
	return o
}

func (o MigrationItemPropertiesResponsePtrOutput) Elem() MigrationItemPropertiesResponseOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) MigrationItemPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MigrationItemPropertiesResponse
		return ret
	}).(MigrationItemPropertiesResponseOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) AllowedOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOperations
	}).(pulumi.StringArrayOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) CurrentJob() CurrentJobDetailsResponsePtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *CurrentJobDetailsResponse {
		if v == nil {
			return nil
		}
		return &v.CurrentJob
	}).(CurrentJobDetailsResponsePtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) EventCorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EventCorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Health
	}).(pulumi.StringPtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) HealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) []HealthErrorResponse {
		if v == nil {
			return nil
		}
		return v.HealthErrors
	}).(HealthErrorResponseArrayOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) LastTestMigrationStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastTestMigrationStatus
	}).(pulumi.StringPtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) LastTestMigrationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastTestMigrationTime
	}).(pulumi.StringPtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) MachineName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MachineName
	}).(pulumi.StringPtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) MigrationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MigrationState
	}).(pulumi.StringPtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) MigrationStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MigrationStateDescription
	}).(pulumi.StringPtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) PolicyFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PolicyFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PolicyId
	}).(pulumi.StringPtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) ProviderSpecificDetails() VMwareCbtMigrationDetailsResponsePtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *VMwareCbtMigrationDetailsResponse {
		if v == nil {
			return nil
		}
		return v.ProviderSpecificDetails
	}).(VMwareCbtMigrationDetailsResponsePtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) TestMigrateState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TestMigrateState
	}).(pulumi.StringPtrOutput)
}

func (o MigrationItemPropertiesResponsePtrOutput) TestMigrateStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MigrationItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TestMigrateStateDescription
	}).(pulumi.StringPtrOutput)
}

type MobilityServiceUpdateResponse struct {
	OsType       *string `pulumi:"osType"`
	RebootStatus *string `pulumi:"rebootStatus"`
	Version      *string `pulumi:"version"`
}





type MobilityServiceUpdateResponseInput interface {
	pulumi.Input

	ToMobilityServiceUpdateResponseOutput() MobilityServiceUpdateResponseOutput
	ToMobilityServiceUpdateResponseOutputWithContext(context.Context) MobilityServiceUpdateResponseOutput
}

type MobilityServiceUpdateResponseArgs struct {
	OsType       pulumi.StringPtrInput `pulumi:"osType"`
	RebootStatus pulumi.StringPtrInput `pulumi:"rebootStatus"`
	Version      pulumi.StringPtrInput `pulumi:"version"`
}

func (MobilityServiceUpdateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MobilityServiceUpdateResponse)(nil)).Elem()
}

func (i MobilityServiceUpdateResponseArgs) ToMobilityServiceUpdateResponseOutput() MobilityServiceUpdateResponseOutput {
	return i.ToMobilityServiceUpdateResponseOutputWithContext(context.Background())
}

func (i MobilityServiceUpdateResponseArgs) ToMobilityServiceUpdateResponseOutputWithContext(ctx context.Context) MobilityServiceUpdateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MobilityServiceUpdateResponseOutput)
}





type MobilityServiceUpdateResponseArrayInput interface {
	pulumi.Input

	ToMobilityServiceUpdateResponseArrayOutput() MobilityServiceUpdateResponseArrayOutput
	ToMobilityServiceUpdateResponseArrayOutputWithContext(context.Context) MobilityServiceUpdateResponseArrayOutput
}

type MobilityServiceUpdateResponseArray []MobilityServiceUpdateResponseInput

func (MobilityServiceUpdateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MobilityServiceUpdateResponse)(nil)).Elem()
}

func (i MobilityServiceUpdateResponseArray) ToMobilityServiceUpdateResponseArrayOutput() MobilityServiceUpdateResponseArrayOutput {
	return i.ToMobilityServiceUpdateResponseArrayOutputWithContext(context.Background())
}

func (i MobilityServiceUpdateResponseArray) ToMobilityServiceUpdateResponseArrayOutputWithContext(ctx context.Context) MobilityServiceUpdateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MobilityServiceUpdateResponseArrayOutput)
}

type MobilityServiceUpdateResponseOutput struct{ *pulumi.OutputState }

func (MobilityServiceUpdateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MobilityServiceUpdateResponse)(nil)).Elem()
}

func (o MobilityServiceUpdateResponseOutput) ToMobilityServiceUpdateResponseOutput() MobilityServiceUpdateResponseOutput {
	return o
}

func (o MobilityServiceUpdateResponseOutput) ToMobilityServiceUpdateResponseOutputWithContext(ctx context.Context) MobilityServiceUpdateResponseOutput {
	return o
}

func (o MobilityServiceUpdateResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MobilityServiceUpdateResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o MobilityServiceUpdateResponseOutput) RebootStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MobilityServiceUpdateResponse) *string { return v.RebootStatus }).(pulumi.StringPtrOutput)
}

func (o MobilityServiceUpdateResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MobilityServiceUpdateResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type MobilityServiceUpdateResponseArrayOutput struct{ *pulumi.OutputState }

func (MobilityServiceUpdateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MobilityServiceUpdateResponse)(nil)).Elem()
}

func (o MobilityServiceUpdateResponseArrayOutput) ToMobilityServiceUpdateResponseArrayOutput() MobilityServiceUpdateResponseArrayOutput {
	return o
}

func (o MobilityServiceUpdateResponseArrayOutput) ToMobilityServiceUpdateResponseArrayOutputWithContext(ctx context.Context) MobilityServiceUpdateResponseArrayOutput {
	return o
}

func (o MobilityServiceUpdateResponseArrayOutput) Index(i pulumi.IntInput) MobilityServiceUpdateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MobilityServiceUpdateResponse {
		return vs[0].([]MobilityServiceUpdateResponse)[vs[1].(int)]
	}).(MobilityServiceUpdateResponseOutput)
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





type NetworkMappingPropertiesResponseInput interface {
	pulumi.Input

	ToNetworkMappingPropertiesResponseOutput() NetworkMappingPropertiesResponseOutput
	ToNetworkMappingPropertiesResponseOutputWithContext(context.Context) NetworkMappingPropertiesResponseOutput
}

type NetworkMappingPropertiesResponseArgs struct {
	FabricSpecificSettings      pulumi.Input          `pulumi:"fabricSpecificSettings"`
	PrimaryFabricFriendlyName   pulumi.StringPtrInput `pulumi:"primaryFabricFriendlyName"`
	PrimaryNetworkFriendlyName  pulumi.StringPtrInput `pulumi:"primaryNetworkFriendlyName"`
	PrimaryNetworkId            pulumi.StringPtrInput `pulumi:"primaryNetworkId"`
	RecoveryFabricArmId         pulumi.StringPtrInput `pulumi:"recoveryFabricArmId"`
	RecoveryFabricFriendlyName  pulumi.StringPtrInput `pulumi:"recoveryFabricFriendlyName"`
	RecoveryNetworkFriendlyName pulumi.StringPtrInput `pulumi:"recoveryNetworkFriendlyName"`
	RecoveryNetworkId           pulumi.StringPtrInput `pulumi:"recoveryNetworkId"`
	State                       pulumi.StringPtrInput `pulumi:"state"`
}

func (NetworkMappingPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkMappingPropertiesResponse)(nil)).Elem()
}

func (i NetworkMappingPropertiesResponseArgs) ToNetworkMappingPropertiesResponseOutput() NetworkMappingPropertiesResponseOutput {
	return i.ToNetworkMappingPropertiesResponseOutputWithContext(context.Background())
}

func (i NetworkMappingPropertiesResponseArgs) ToNetworkMappingPropertiesResponseOutputWithContext(ctx context.Context) NetworkMappingPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMappingPropertiesResponseOutput)
}

func (i NetworkMappingPropertiesResponseArgs) ToNetworkMappingPropertiesResponsePtrOutput() NetworkMappingPropertiesResponsePtrOutput {
	return i.ToNetworkMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i NetworkMappingPropertiesResponseArgs) ToNetworkMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) NetworkMappingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMappingPropertiesResponseOutput).ToNetworkMappingPropertiesResponsePtrOutputWithContext(ctx)
}









type NetworkMappingPropertiesResponsePtrInput interface {
	pulumi.Input

	ToNetworkMappingPropertiesResponsePtrOutput() NetworkMappingPropertiesResponsePtrOutput
	ToNetworkMappingPropertiesResponsePtrOutputWithContext(context.Context) NetworkMappingPropertiesResponsePtrOutput
}

type networkMappingPropertiesResponsePtrType NetworkMappingPropertiesResponseArgs

func NetworkMappingPropertiesResponsePtr(v *NetworkMappingPropertiesResponseArgs) NetworkMappingPropertiesResponsePtrInput {
	return (*networkMappingPropertiesResponsePtrType)(v)
}

func (*networkMappingPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkMappingPropertiesResponse)(nil)).Elem()
}

func (i *networkMappingPropertiesResponsePtrType) ToNetworkMappingPropertiesResponsePtrOutput() NetworkMappingPropertiesResponsePtrOutput {
	return i.ToNetworkMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *networkMappingPropertiesResponsePtrType) ToNetworkMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) NetworkMappingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkMappingPropertiesResponsePtrOutput)
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

func (o NetworkMappingPropertiesResponseOutput) ToNetworkMappingPropertiesResponsePtrOutput() NetworkMappingPropertiesResponsePtrOutput {
	return o.ToNetworkMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o NetworkMappingPropertiesResponseOutput) ToNetworkMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) NetworkMappingPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkMappingPropertiesResponse) *NetworkMappingPropertiesResponse {
		return &v
	}).(NetworkMappingPropertiesResponsePtrOutput)
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

type NetworkMappingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkMappingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkMappingPropertiesResponse)(nil)).Elem()
}

func (o NetworkMappingPropertiesResponsePtrOutput) ToNetworkMappingPropertiesResponsePtrOutput() NetworkMappingPropertiesResponsePtrOutput {
	return o
}

func (o NetworkMappingPropertiesResponsePtrOutput) ToNetworkMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) NetworkMappingPropertiesResponsePtrOutput {
	return o
}

func (o NetworkMappingPropertiesResponsePtrOutput) Elem() NetworkMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *NetworkMappingPropertiesResponse) NetworkMappingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret NetworkMappingPropertiesResponse
		return ret
	}).(NetworkMappingPropertiesResponseOutput)
}

func (o NetworkMappingPropertiesResponsePtrOutput) FabricSpecificSettings() pulumi.AnyOutput {
	return o.ApplyT(func(v *NetworkMappingPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.FabricSpecificSettings
	}).(pulumi.AnyOutput)
}

func (o NetworkMappingPropertiesResponsePtrOutput) PrimaryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryFabricFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponsePtrOutput) PrimaryNetworkFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryNetworkFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponsePtrOutput) PrimaryNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryNetworkId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponsePtrOutput) RecoveryFabricArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryFabricArmId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponsePtrOutput) RecoveryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryFabricFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponsePtrOutput) RecoveryNetworkFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryNetworkFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponsePtrOutput) RecoveryNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryNetworkId
	}).(pulumi.StringPtrOutput)
}

func (o NetworkMappingPropertiesResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type OSDetailsResponse struct {
	OSMajorVersion *string `pulumi:"oSMajorVersion"`
	OSMinorVersion *string `pulumi:"oSMinorVersion"`
	OSVersion      *string `pulumi:"oSVersion"`
	OsEdition      *string `pulumi:"osEdition"`
	OsType         *string `pulumi:"osType"`
	ProductType    *string `pulumi:"productType"`
}





type OSDetailsResponseInput interface {
	pulumi.Input

	ToOSDetailsResponseOutput() OSDetailsResponseOutput
	ToOSDetailsResponseOutputWithContext(context.Context) OSDetailsResponseOutput
}

type OSDetailsResponseArgs struct {
	OSMajorVersion pulumi.StringPtrInput `pulumi:"oSMajorVersion"`
	OSMinorVersion pulumi.StringPtrInput `pulumi:"oSMinorVersion"`
	OSVersion      pulumi.StringPtrInput `pulumi:"oSVersion"`
	OsEdition      pulumi.StringPtrInput `pulumi:"osEdition"`
	OsType         pulumi.StringPtrInput `pulumi:"osType"`
	ProductType    pulumi.StringPtrInput `pulumi:"productType"`
}

func (OSDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDetailsResponse)(nil)).Elem()
}

func (i OSDetailsResponseArgs) ToOSDetailsResponseOutput() OSDetailsResponseOutput {
	return i.ToOSDetailsResponseOutputWithContext(context.Background())
}

func (i OSDetailsResponseArgs) ToOSDetailsResponseOutputWithContext(ctx context.Context) OSDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDetailsResponseOutput)
}

func (i OSDetailsResponseArgs) ToOSDetailsResponsePtrOutput() OSDetailsResponsePtrOutput {
	return i.ToOSDetailsResponsePtrOutputWithContext(context.Background())
}

func (i OSDetailsResponseArgs) ToOSDetailsResponsePtrOutputWithContext(ctx context.Context) OSDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDetailsResponseOutput).ToOSDetailsResponsePtrOutputWithContext(ctx)
}









type OSDetailsResponsePtrInput interface {
	pulumi.Input

	ToOSDetailsResponsePtrOutput() OSDetailsResponsePtrOutput
	ToOSDetailsResponsePtrOutputWithContext(context.Context) OSDetailsResponsePtrOutput
}

type osdetailsResponsePtrType OSDetailsResponseArgs

func OSDetailsResponsePtr(v *OSDetailsResponseArgs) OSDetailsResponsePtrInput {
	return (*osdetailsResponsePtrType)(v)
}

func (*osdetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDetailsResponse)(nil)).Elem()
}

func (i *osdetailsResponsePtrType) ToOSDetailsResponsePtrOutput() OSDetailsResponsePtrOutput {
	return i.ToOSDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *osdetailsResponsePtrType) ToOSDetailsResponsePtrOutputWithContext(ctx context.Context) OSDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDetailsResponsePtrOutput)
}

type OSDetailsResponseOutput struct{ *pulumi.OutputState }

func (OSDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDetailsResponse)(nil)).Elem()
}

func (o OSDetailsResponseOutput) ToOSDetailsResponseOutput() OSDetailsResponseOutput {
	return o
}

func (o OSDetailsResponseOutput) ToOSDetailsResponseOutputWithContext(ctx context.Context) OSDetailsResponseOutput {
	return o
}

func (o OSDetailsResponseOutput) ToOSDetailsResponsePtrOutput() OSDetailsResponsePtrOutput {
	return o.ToOSDetailsResponsePtrOutputWithContext(context.Background())
}

func (o OSDetailsResponseOutput) ToOSDetailsResponsePtrOutputWithContext(ctx context.Context) OSDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSDetailsResponse) *OSDetailsResponse {
		return &v
	}).(OSDetailsResponsePtrOutput)
}

func (o OSDetailsResponseOutput) OSMajorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDetailsResponse) *string { return v.OSMajorVersion }).(pulumi.StringPtrOutput)
}

func (o OSDetailsResponseOutput) OSMinorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDetailsResponse) *string { return v.OSMinorVersion }).(pulumi.StringPtrOutput)
}

func (o OSDetailsResponseOutput) OSVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDetailsResponse) *string { return v.OSVersion }).(pulumi.StringPtrOutput)
}

func (o OSDetailsResponseOutput) OsEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDetailsResponse) *string { return v.OsEdition }).(pulumi.StringPtrOutput)
}

func (o OSDetailsResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDetailsResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OSDetailsResponseOutput) ProductType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDetailsResponse) *string { return v.ProductType }).(pulumi.StringPtrOutput)
}

type OSDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (OSDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDetailsResponse)(nil)).Elem()
}

func (o OSDetailsResponsePtrOutput) ToOSDetailsResponsePtrOutput() OSDetailsResponsePtrOutput {
	return o
}

func (o OSDetailsResponsePtrOutput) ToOSDetailsResponsePtrOutputWithContext(ctx context.Context) OSDetailsResponsePtrOutput {
	return o
}

func (o OSDetailsResponsePtrOutput) Elem() OSDetailsResponseOutput {
	return o.ApplyT(func(v *OSDetailsResponse) OSDetailsResponse {
		if v != nil {
			return *v
		}
		var ret OSDetailsResponse
		return ret
	}).(OSDetailsResponseOutput)
}

func (o OSDetailsResponsePtrOutput) OSMajorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.OSMajorVersion
	}).(pulumi.StringPtrOutput)
}

func (o OSDetailsResponsePtrOutput) OSMinorVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.OSMinorVersion
	}).(pulumi.StringPtrOutput)
}

func (o OSDetailsResponsePtrOutput) OSVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.OSVersion
	}).(pulumi.StringPtrOutput)
}

func (o OSDetailsResponsePtrOutput) OsEdition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsEdition
	}).(pulumi.StringPtrOutput)
}

func (o OSDetailsResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o OSDetailsResponsePtrOutput) ProductType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProductType
	}).(pulumi.StringPtrOutput)
}

type OSDiskDetailsResponse struct {
	OsType  *string `pulumi:"osType"`
	OsVhdId *string `pulumi:"osVhdId"`
	VhdName *string `pulumi:"vhdName"`
}





type OSDiskDetailsResponseInput interface {
	pulumi.Input

	ToOSDiskDetailsResponseOutput() OSDiskDetailsResponseOutput
	ToOSDiskDetailsResponseOutputWithContext(context.Context) OSDiskDetailsResponseOutput
}

type OSDiskDetailsResponseArgs struct {
	OsType  pulumi.StringPtrInput `pulumi:"osType"`
	OsVhdId pulumi.StringPtrInput `pulumi:"osVhdId"`
	VhdName pulumi.StringPtrInput `pulumi:"vhdName"`
}

func (OSDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskDetailsResponse)(nil)).Elem()
}

func (i OSDiskDetailsResponseArgs) ToOSDiskDetailsResponseOutput() OSDiskDetailsResponseOutput {
	return i.ToOSDiskDetailsResponseOutputWithContext(context.Background())
}

func (i OSDiskDetailsResponseArgs) ToOSDiskDetailsResponseOutputWithContext(ctx context.Context) OSDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskDetailsResponseOutput)
}

func (i OSDiskDetailsResponseArgs) ToOSDiskDetailsResponsePtrOutput() OSDiskDetailsResponsePtrOutput {
	return i.ToOSDiskDetailsResponsePtrOutputWithContext(context.Background())
}

func (i OSDiskDetailsResponseArgs) ToOSDiskDetailsResponsePtrOutputWithContext(ctx context.Context) OSDiskDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskDetailsResponseOutput).ToOSDiskDetailsResponsePtrOutputWithContext(ctx)
}









type OSDiskDetailsResponsePtrInput interface {
	pulumi.Input

	ToOSDiskDetailsResponsePtrOutput() OSDiskDetailsResponsePtrOutput
	ToOSDiskDetailsResponsePtrOutputWithContext(context.Context) OSDiskDetailsResponsePtrOutput
}

type osdiskDetailsResponsePtrType OSDiskDetailsResponseArgs

func OSDiskDetailsResponsePtr(v *OSDiskDetailsResponseArgs) OSDiskDetailsResponsePtrInput {
	return (*osdiskDetailsResponsePtrType)(v)
}

func (*osdiskDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskDetailsResponse)(nil)).Elem()
}

func (i *osdiskDetailsResponsePtrType) ToOSDiskDetailsResponsePtrOutput() OSDiskDetailsResponsePtrOutput {
	return i.ToOSDiskDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *osdiskDetailsResponsePtrType) ToOSDiskDetailsResponsePtrOutputWithContext(ctx context.Context) OSDiskDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OSDiskDetailsResponsePtrOutput)
}

type OSDiskDetailsResponseOutput struct{ *pulumi.OutputState }

func (OSDiskDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OSDiskDetailsResponse)(nil)).Elem()
}

func (o OSDiskDetailsResponseOutput) ToOSDiskDetailsResponseOutput() OSDiskDetailsResponseOutput {
	return o
}

func (o OSDiskDetailsResponseOutput) ToOSDiskDetailsResponseOutputWithContext(ctx context.Context) OSDiskDetailsResponseOutput {
	return o
}

func (o OSDiskDetailsResponseOutput) ToOSDiskDetailsResponsePtrOutput() OSDiskDetailsResponsePtrOutput {
	return o.ToOSDiskDetailsResponsePtrOutputWithContext(context.Background())
}

func (o OSDiskDetailsResponseOutput) ToOSDiskDetailsResponsePtrOutputWithContext(ctx context.Context) OSDiskDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OSDiskDetailsResponse) *OSDiskDetailsResponse {
		return &v
	}).(OSDiskDetailsResponsePtrOutput)
}

func (o OSDiskDetailsResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskDetailsResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o OSDiskDetailsResponseOutput) OsVhdId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskDetailsResponse) *string { return v.OsVhdId }).(pulumi.StringPtrOutput)
}

func (o OSDiskDetailsResponseOutput) VhdName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OSDiskDetailsResponse) *string { return v.VhdName }).(pulumi.StringPtrOutput)
}

type OSDiskDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (OSDiskDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OSDiskDetailsResponse)(nil)).Elem()
}

func (o OSDiskDetailsResponsePtrOutput) ToOSDiskDetailsResponsePtrOutput() OSDiskDetailsResponsePtrOutput {
	return o
}

func (o OSDiskDetailsResponsePtrOutput) ToOSDiskDetailsResponsePtrOutputWithContext(ctx context.Context) OSDiskDetailsResponsePtrOutput {
	return o
}

func (o OSDiskDetailsResponsePtrOutput) Elem() OSDiskDetailsResponseOutput {
	return o.ApplyT(func(v *OSDiskDetailsResponse) OSDiskDetailsResponse {
		if v != nil {
			return *v
		}
		var ret OSDiskDetailsResponse
		return ret
	}).(OSDiskDetailsResponseOutput)
}

func (o OSDiskDetailsResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsType
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskDetailsResponsePtrOutput) OsVhdId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.OsVhdId
	}).(pulumi.StringPtrOutput)
}

func (o OSDiskDetailsResponsePtrOutput) VhdName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OSDiskDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.VhdName
	}).(pulumi.StringPtrOutput)
}

type PolicyPropertiesResponse struct {
	FriendlyName            *string     `pulumi:"friendlyName"`
	ProviderSpecificDetails interface{} `pulumi:"providerSpecificDetails"`
}





type PolicyPropertiesResponseInput interface {
	pulumi.Input

	ToPolicyPropertiesResponseOutput() PolicyPropertiesResponseOutput
	ToPolicyPropertiesResponseOutputWithContext(context.Context) PolicyPropertiesResponseOutput
}

type PolicyPropertiesResponseArgs struct {
	FriendlyName            pulumi.StringPtrInput `pulumi:"friendlyName"`
	ProviderSpecificDetails pulumi.Input          `pulumi:"providerSpecificDetails"`
}

func (PolicyPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyPropertiesResponse)(nil)).Elem()
}

func (i PolicyPropertiesResponseArgs) ToPolicyPropertiesResponseOutput() PolicyPropertiesResponseOutput {
	return i.ToPolicyPropertiesResponseOutputWithContext(context.Background())
}

func (i PolicyPropertiesResponseArgs) ToPolicyPropertiesResponseOutputWithContext(ctx context.Context) PolicyPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPropertiesResponseOutput)
}

func (i PolicyPropertiesResponseArgs) ToPolicyPropertiesResponsePtrOutput() PolicyPropertiesResponsePtrOutput {
	return i.ToPolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PolicyPropertiesResponseArgs) ToPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) PolicyPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPropertiesResponseOutput).ToPolicyPropertiesResponsePtrOutputWithContext(ctx)
}









type PolicyPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPolicyPropertiesResponsePtrOutput() PolicyPropertiesResponsePtrOutput
	ToPolicyPropertiesResponsePtrOutputWithContext(context.Context) PolicyPropertiesResponsePtrOutput
}

type policyPropertiesResponsePtrType PolicyPropertiesResponseArgs

func PolicyPropertiesResponsePtr(v *PolicyPropertiesResponseArgs) PolicyPropertiesResponsePtrInput {
	return (*policyPropertiesResponsePtrType)(v)
}

func (*policyPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyPropertiesResponse)(nil)).Elem()
}

func (i *policyPropertiesResponsePtrType) ToPolicyPropertiesResponsePtrOutput() PolicyPropertiesResponsePtrOutput {
	return i.ToPolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *policyPropertiesResponsePtrType) ToPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) PolicyPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPropertiesResponsePtrOutput)
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

func (o PolicyPropertiesResponseOutput) ToPolicyPropertiesResponsePtrOutput() PolicyPropertiesResponsePtrOutput {
	return o.ToPolicyPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PolicyPropertiesResponseOutput) ToPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) PolicyPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyPropertiesResponse) *PolicyPropertiesResponse {
		return &v
	}).(PolicyPropertiesResponsePtrOutput)
}

func (o PolicyPropertiesResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PolicyPropertiesResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o PolicyPropertiesResponseOutput) ProviderSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v PolicyPropertiesResponse) interface{} { return v.ProviderSpecificDetails }).(pulumi.AnyOutput)
}

type PolicyPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PolicyPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyPropertiesResponse)(nil)).Elem()
}

func (o PolicyPropertiesResponsePtrOutput) ToPolicyPropertiesResponsePtrOutput() PolicyPropertiesResponsePtrOutput {
	return o
}

func (o PolicyPropertiesResponsePtrOutput) ToPolicyPropertiesResponsePtrOutputWithContext(ctx context.Context) PolicyPropertiesResponsePtrOutput {
	return o
}

func (o PolicyPropertiesResponsePtrOutput) Elem() PolicyPropertiesResponseOutput {
	return o.ApplyT(func(v *PolicyPropertiesResponse) PolicyPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PolicyPropertiesResponse
		return ret
	}).(PolicyPropertiesResponseOutput)
}

func (o PolicyPropertiesResponsePtrOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PolicyPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o PolicyPropertiesResponsePtrOutput) ProviderSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *PolicyPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.ProviderSpecificDetails
	}).(pulumi.AnyOutput)
}

type ProcessServerResponse struct {
	AgentExpiryDate            *string                         `pulumi:"agentExpiryDate"`
	AgentVersion               *string                         `pulumi:"agentVersion"`
	AgentVersionDetails        *VersionDetailsResponse         `pulumi:"agentVersionDetails"`
	AvailableMemoryInBytes     *float64                        `pulumi:"availableMemoryInBytes"`
	AvailableSpaceInBytes      *float64                        `pulumi:"availableSpaceInBytes"`
	CpuLoad                    *string                         `pulumi:"cpuLoad"`
	CpuLoadStatus              *string                         `pulumi:"cpuLoadStatus"`
	FriendlyName               *string                         `pulumi:"friendlyName"`
	HealthErrors               []HealthErrorResponse           `pulumi:"healthErrors"`
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





type ProcessServerResponseInput interface {
	pulumi.Input

	ToProcessServerResponseOutput() ProcessServerResponseOutput
	ToProcessServerResponseOutputWithContext(context.Context) ProcessServerResponseOutput
}

type ProcessServerResponseArgs struct {
	AgentExpiryDate            pulumi.StringPtrInput                   `pulumi:"agentExpiryDate"`
	AgentVersion               pulumi.StringPtrInput                   `pulumi:"agentVersion"`
	AgentVersionDetails        VersionDetailsResponsePtrInput          `pulumi:"agentVersionDetails"`
	AvailableMemoryInBytes     pulumi.Float64PtrInput                  `pulumi:"availableMemoryInBytes"`
	AvailableSpaceInBytes      pulumi.Float64PtrInput                  `pulumi:"availableSpaceInBytes"`
	CpuLoad                    pulumi.StringPtrInput                   `pulumi:"cpuLoad"`
	CpuLoadStatus              pulumi.StringPtrInput                   `pulumi:"cpuLoadStatus"`
	FriendlyName               pulumi.StringPtrInput                   `pulumi:"friendlyName"`
	HealthErrors               HealthErrorResponseArrayInput           `pulumi:"healthErrors"`
	HostId                     pulumi.StringPtrInput                   `pulumi:"hostId"`
	Id                         pulumi.StringPtrInput                   `pulumi:"id"`
	IpAddress                  pulumi.StringPtrInput                   `pulumi:"ipAddress"`
	LastHeartbeat              pulumi.StringPtrInput                   `pulumi:"lastHeartbeat"`
	MachineCount               pulumi.StringPtrInput                   `pulumi:"machineCount"`
	MemoryUsageStatus          pulumi.StringPtrInput                   `pulumi:"memoryUsageStatus"`
	MobilityServiceUpdates     MobilityServiceUpdateResponseArrayInput `pulumi:"mobilityServiceUpdates"`
	OsType                     pulumi.StringPtrInput                   `pulumi:"osType"`
	OsVersion                  pulumi.StringPtrInput                   `pulumi:"osVersion"`
	PsServiceStatus            pulumi.StringPtrInput                   `pulumi:"psServiceStatus"`
	ReplicationPairCount       pulumi.StringPtrInput                   `pulumi:"replicationPairCount"`
	SpaceUsageStatus           pulumi.StringPtrInput                   `pulumi:"spaceUsageStatus"`
	SslCertExpiryDate          pulumi.StringPtrInput                   `pulumi:"sslCertExpiryDate"`
	SslCertExpiryRemainingDays pulumi.IntPtrInput                      `pulumi:"sslCertExpiryRemainingDays"`
	SystemLoad                 pulumi.StringPtrInput                   `pulumi:"systemLoad"`
	SystemLoadStatus           pulumi.StringPtrInput                   `pulumi:"systemLoadStatus"`
	TotalMemoryInBytes         pulumi.Float64PtrInput                  `pulumi:"totalMemoryInBytes"`
	TotalSpaceInBytes          pulumi.Float64PtrInput                  `pulumi:"totalSpaceInBytes"`
	VersionStatus              pulumi.StringPtrInput                   `pulumi:"versionStatus"`
}

func (ProcessServerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProcessServerResponse)(nil)).Elem()
}

func (i ProcessServerResponseArgs) ToProcessServerResponseOutput() ProcessServerResponseOutput {
	return i.ToProcessServerResponseOutputWithContext(context.Background())
}

func (i ProcessServerResponseArgs) ToProcessServerResponseOutputWithContext(ctx context.Context) ProcessServerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProcessServerResponseOutput)
}





type ProcessServerResponseArrayInput interface {
	pulumi.Input

	ToProcessServerResponseArrayOutput() ProcessServerResponseArrayOutput
	ToProcessServerResponseArrayOutputWithContext(context.Context) ProcessServerResponseArrayOutput
}

type ProcessServerResponseArray []ProcessServerResponseInput

func (ProcessServerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProcessServerResponse)(nil)).Elem()
}

func (i ProcessServerResponseArray) ToProcessServerResponseArrayOutput() ProcessServerResponseArrayOutput {
	return i.ToProcessServerResponseArrayOutputWithContext(context.Background())
}

func (i ProcessServerResponseArray) ToProcessServerResponseArrayOutputWithContext(ctx context.Context) ProcessServerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProcessServerResponseArrayOutput)
}

type ProcessServerResponseOutput struct{ *pulumi.OutputState }

func (ProcessServerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProcessServerResponse)(nil)).Elem()
}

func (o ProcessServerResponseOutput) ToProcessServerResponseOutput() ProcessServerResponseOutput {
	return o
}

func (o ProcessServerResponseOutput) ToProcessServerResponseOutputWithContext(ctx context.Context) ProcessServerResponseOutput {
	return o
}

func (o ProcessServerResponseOutput) AgentExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.AgentExpiryDate }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) AgentVersionDetails() VersionDetailsResponsePtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *VersionDetailsResponse { return v.AgentVersionDetails }).(VersionDetailsResponsePtrOutput)
}

func (o ProcessServerResponseOutput) AvailableMemoryInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *float64 { return v.AvailableMemoryInBytes }).(pulumi.Float64PtrOutput)
}

func (o ProcessServerResponseOutput) AvailableSpaceInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *float64 { return v.AvailableSpaceInBytes }).(pulumi.Float64PtrOutput)
}

func (o ProcessServerResponseOutput) CpuLoad() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.CpuLoad }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) CpuLoadStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.CpuLoadStatus }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) HealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v ProcessServerResponse) []HealthErrorResponse { return v.HealthErrors }).(HealthErrorResponseArrayOutput)
}

func (o ProcessServerResponseOutput) HostId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.HostId }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) LastHeartbeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.LastHeartbeat }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) MachineCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.MachineCount }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) MemoryUsageStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.MemoryUsageStatus }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) MobilityServiceUpdates() MobilityServiceUpdateResponseArrayOutput {
	return o.ApplyT(func(v ProcessServerResponse) []MobilityServiceUpdateResponse { return v.MobilityServiceUpdates }).(MobilityServiceUpdateResponseArrayOutput)
}

func (o ProcessServerResponseOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.OsType }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) OsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.OsVersion }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) PsServiceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.PsServiceStatus }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) ReplicationPairCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.ReplicationPairCount }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) SpaceUsageStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.SpaceUsageStatus }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) SslCertExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.SslCertExpiryDate }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) SslCertExpiryRemainingDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *int { return v.SslCertExpiryRemainingDays }).(pulumi.IntPtrOutput)
}

func (o ProcessServerResponseOutput) SystemLoad() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.SystemLoad }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) SystemLoadStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.SystemLoadStatus }).(pulumi.StringPtrOutput)
}

func (o ProcessServerResponseOutput) TotalMemoryInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *float64 { return v.TotalMemoryInBytes }).(pulumi.Float64PtrOutput)
}

func (o ProcessServerResponseOutput) TotalSpaceInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *float64 { return v.TotalSpaceInBytes }).(pulumi.Float64PtrOutput)
}

func (o ProcessServerResponseOutput) VersionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProcessServerResponse) *string { return v.VersionStatus }).(pulumi.StringPtrOutput)
}

type ProcessServerResponseArrayOutput struct{ *pulumi.OutputState }

func (ProcessServerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ProcessServerResponse)(nil)).Elem()
}

func (o ProcessServerResponseArrayOutput) ToProcessServerResponseArrayOutput() ProcessServerResponseArrayOutput {
	return o
}

func (o ProcessServerResponseArrayOutput) ToProcessServerResponseArrayOutputWithContext(ctx context.Context) ProcessServerResponseArrayOutput {
	return o
}

func (o ProcessServerResponseArrayOutput) Index(i pulumi.IntInput) ProcessServerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ProcessServerResponse {
		return vs[0].([]ProcessServerResponse)[vs[1].(int)]
	}).(ProcessServerResponseOutput)
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





type ProtectionContainerMappingPropertiesResponseInput interface {
	pulumi.Input

	ToProtectionContainerMappingPropertiesResponseOutput() ProtectionContainerMappingPropertiesResponseOutput
	ToProtectionContainerMappingPropertiesResponseOutputWithContext(context.Context) ProtectionContainerMappingPropertiesResponseOutput
}

type ProtectionContainerMappingPropertiesResponseArgs struct {
	Health                                pulumi.StringPtrInput         `pulumi:"health"`
	HealthErrorDetails                    HealthErrorResponseArrayInput `pulumi:"healthErrorDetails"`
	PolicyFriendlyName                    pulumi.StringPtrInput         `pulumi:"policyFriendlyName"`
	PolicyId                              pulumi.StringPtrInput         `pulumi:"policyId"`
	ProviderSpecificDetails               pulumi.Input                  `pulumi:"providerSpecificDetails"`
	SourceFabricFriendlyName              pulumi.StringPtrInput         `pulumi:"sourceFabricFriendlyName"`
	SourceProtectionContainerFriendlyName pulumi.StringPtrInput         `pulumi:"sourceProtectionContainerFriendlyName"`
	State                                 pulumi.StringPtrInput         `pulumi:"state"`
	TargetFabricFriendlyName              pulumi.StringPtrInput         `pulumi:"targetFabricFriendlyName"`
	TargetProtectionContainerFriendlyName pulumi.StringPtrInput         `pulumi:"targetProtectionContainerFriendlyName"`
	TargetProtectionContainerId           pulumi.StringPtrInput         `pulumi:"targetProtectionContainerId"`
}

func (ProtectionContainerMappingPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProtectionContainerMappingPropertiesResponse)(nil)).Elem()
}

func (i ProtectionContainerMappingPropertiesResponseArgs) ToProtectionContainerMappingPropertiesResponseOutput() ProtectionContainerMappingPropertiesResponseOutput {
	return i.ToProtectionContainerMappingPropertiesResponseOutputWithContext(context.Background())
}

func (i ProtectionContainerMappingPropertiesResponseArgs) ToProtectionContainerMappingPropertiesResponseOutputWithContext(ctx context.Context) ProtectionContainerMappingPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerMappingPropertiesResponseOutput)
}

func (i ProtectionContainerMappingPropertiesResponseArgs) ToProtectionContainerMappingPropertiesResponsePtrOutput() ProtectionContainerMappingPropertiesResponsePtrOutput {
	return i.ToProtectionContainerMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ProtectionContainerMappingPropertiesResponseArgs) ToProtectionContainerMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) ProtectionContainerMappingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerMappingPropertiesResponseOutput).ToProtectionContainerMappingPropertiesResponsePtrOutputWithContext(ctx)
}









type ProtectionContainerMappingPropertiesResponsePtrInput interface {
	pulumi.Input

	ToProtectionContainerMappingPropertiesResponsePtrOutput() ProtectionContainerMappingPropertiesResponsePtrOutput
	ToProtectionContainerMappingPropertiesResponsePtrOutputWithContext(context.Context) ProtectionContainerMappingPropertiesResponsePtrOutput
}

type protectionContainerMappingPropertiesResponsePtrType ProtectionContainerMappingPropertiesResponseArgs

func ProtectionContainerMappingPropertiesResponsePtr(v *ProtectionContainerMappingPropertiesResponseArgs) ProtectionContainerMappingPropertiesResponsePtrInput {
	return (*protectionContainerMappingPropertiesResponsePtrType)(v)
}

func (*protectionContainerMappingPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainerMappingPropertiesResponse)(nil)).Elem()
}

func (i *protectionContainerMappingPropertiesResponsePtrType) ToProtectionContainerMappingPropertiesResponsePtrOutput() ProtectionContainerMappingPropertiesResponsePtrOutput {
	return i.ToProtectionContainerMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *protectionContainerMappingPropertiesResponsePtrType) ToProtectionContainerMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) ProtectionContainerMappingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProtectionContainerMappingPropertiesResponsePtrOutput)
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

func (o ProtectionContainerMappingPropertiesResponseOutput) ToProtectionContainerMappingPropertiesResponsePtrOutput() ProtectionContainerMappingPropertiesResponsePtrOutput {
	return o.ToProtectionContainerMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ProtectionContainerMappingPropertiesResponseOutput) ToProtectionContainerMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) ProtectionContainerMappingPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProtectionContainerMappingPropertiesResponse) *ProtectionContainerMappingPropertiesResponse {
		return &v
	}).(ProtectionContainerMappingPropertiesResponsePtrOutput)
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

type ProtectionContainerMappingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ProtectionContainerMappingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProtectionContainerMappingPropertiesResponse)(nil)).Elem()
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) ToProtectionContainerMappingPropertiesResponsePtrOutput() ProtectionContainerMappingPropertiesResponsePtrOutput {
	return o
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) ToProtectionContainerMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) ProtectionContainerMappingPropertiesResponsePtrOutput {
	return o
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) Elem() ProtectionContainerMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) ProtectionContainerMappingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ProtectionContainerMappingPropertiesResponse
		return ret
	}).(ProtectionContainerMappingPropertiesResponseOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) Health() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Health
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) HealthErrorDetails() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) []HealthErrorResponse {
		if v == nil {
			return nil
		}
		return v.HealthErrorDetails
	}).(HealthErrorResponseArrayOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) PolicyFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyId
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) ProviderSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.ProviderSpecificDetails
	}).(pulumi.AnyOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) SourceFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceFabricFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) SourceProtectionContainerFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceProtectionContainerFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) TargetFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetFabricFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) TargetProtectionContainerFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetProtectionContainerFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ProtectionContainerMappingPropertiesResponsePtrOutput) TargetProtectionContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProtectionContainerMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetProtectionContainerId
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





type RcmAzureMigrationPolicyDetailsResponseInput interface {
	pulumi.Input

	ToRcmAzureMigrationPolicyDetailsResponseOutput() RcmAzureMigrationPolicyDetailsResponseOutput
	ToRcmAzureMigrationPolicyDetailsResponseOutputWithContext(context.Context) RcmAzureMigrationPolicyDetailsResponseOutput
}

type RcmAzureMigrationPolicyDetailsResponseArgs struct {
	AppConsistentFrequencyInMinutes   pulumi.IntPtrInput    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes pulumi.IntPtrInput    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      pulumi.StringInput    `pulumi:"instanceType"`
	MultiVmSyncStatus                 pulumi.StringPtrInput `pulumi:"multiVmSyncStatus"`
	RecoveryPointHistory              pulumi.IntPtrInput    `pulumi:"recoveryPointHistory"`
	RecoveryPointThresholdInMinutes   pulumi.IntPtrInput    `pulumi:"recoveryPointThresholdInMinutes"`
}

func (RcmAzureMigrationPolicyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RcmAzureMigrationPolicyDetailsResponse)(nil)).Elem()
}

func (i RcmAzureMigrationPolicyDetailsResponseArgs) ToRcmAzureMigrationPolicyDetailsResponseOutput() RcmAzureMigrationPolicyDetailsResponseOutput {
	return i.ToRcmAzureMigrationPolicyDetailsResponseOutputWithContext(context.Background())
}

func (i RcmAzureMigrationPolicyDetailsResponseArgs) ToRcmAzureMigrationPolicyDetailsResponseOutputWithContext(ctx context.Context) RcmAzureMigrationPolicyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RcmAzureMigrationPolicyDetailsResponseOutput)
}

type RcmAzureMigrationPolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (RcmAzureMigrationPolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RcmAzureMigrationPolicyDetailsResponse)(nil)).Elem()
}

func (o RcmAzureMigrationPolicyDetailsResponseOutput) ToRcmAzureMigrationPolicyDetailsResponseOutput() RcmAzureMigrationPolicyDetailsResponseOutput {
	return o
}

func (o RcmAzureMigrationPolicyDetailsResponseOutput) ToRcmAzureMigrationPolicyDetailsResponseOutputWithContext(ctx context.Context) RcmAzureMigrationPolicyDetailsResponseOutput {
	return o
}

func (o RcmAzureMigrationPolicyDetailsResponseOutput) AppConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RcmAzureMigrationPolicyDetailsResponse) *int { return v.AppConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o RcmAzureMigrationPolicyDetailsResponseOutput) CrashConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RcmAzureMigrationPolicyDetailsResponse) *int { return v.CrashConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o RcmAzureMigrationPolicyDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v RcmAzureMigrationPolicyDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o RcmAzureMigrationPolicyDetailsResponseOutput) MultiVmSyncStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RcmAzureMigrationPolicyDetailsResponse) *string { return v.MultiVmSyncStatus }).(pulumi.StringPtrOutput)
}

func (o RcmAzureMigrationPolicyDetailsResponseOutput) RecoveryPointHistory() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RcmAzureMigrationPolicyDetailsResponse) *int { return v.RecoveryPointHistory }).(pulumi.IntPtrOutput)
}

func (o RcmAzureMigrationPolicyDetailsResponseOutput) RecoveryPointThresholdInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RcmAzureMigrationPolicyDetailsResponse) *int { return v.RecoveryPointThresholdInMinutes }).(pulumi.IntPtrOutput)
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





type RecoveryPlanActionResponseInput interface {
	pulumi.Input

	ToRecoveryPlanActionResponseOutput() RecoveryPlanActionResponseOutput
	ToRecoveryPlanActionResponseOutputWithContext(context.Context) RecoveryPlanActionResponseOutput
}

type RecoveryPlanActionResponseArgs struct {
	ActionName         pulumi.StringInput      `pulumi:"actionName"`
	CustomDetails      pulumi.Input            `pulumi:"customDetails"`
	FailoverDirections pulumi.StringArrayInput `pulumi:"failoverDirections"`
	FailoverTypes      pulumi.StringArrayInput `pulumi:"failoverTypes"`
}

func (RecoveryPlanActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanActionResponse)(nil)).Elem()
}

func (i RecoveryPlanActionResponseArgs) ToRecoveryPlanActionResponseOutput() RecoveryPlanActionResponseOutput {
	return i.ToRecoveryPlanActionResponseOutputWithContext(context.Background())
}

func (i RecoveryPlanActionResponseArgs) ToRecoveryPlanActionResponseOutputWithContext(ctx context.Context) RecoveryPlanActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanActionResponseOutput)
}





type RecoveryPlanActionResponseArrayInput interface {
	pulumi.Input

	ToRecoveryPlanActionResponseArrayOutput() RecoveryPlanActionResponseArrayOutput
	ToRecoveryPlanActionResponseArrayOutputWithContext(context.Context) RecoveryPlanActionResponseArrayOutput
}

type RecoveryPlanActionResponseArray []RecoveryPlanActionResponseInput

func (RecoveryPlanActionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanActionResponse)(nil)).Elem()
}

func (i RecoveryPlanActionResponseArray) ToRecoveryPlanActionResponseArrayOutput() RecoveryPlanActionResponseArrayOutput {
	return i.ToRecoveryPlanActionResponseArrayOutputWithContext(context.Background())
}

func (i RecoveryPlanActionResponseArray) ToRecoveryPlanActionResponseArrayOutputWithContext(ctx context.Context) RecoveryPlanActionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanActionResponseArrayOutput)
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





type RecoveryPlanAutomationRunbookActionDetailsResponseInput interface {
	pulumi.Input

	ToRecoveryPlanAutomationRunbookActionDetailsResponseOutput() RecoveryPlanAutomationRunbookActionDetailsResponseOutput
	ToRecoveryPlanAutomationRunbookActionDetailsResponseOutputWithContext(context.Context) RecoveryPlanAutomationRunbookActionDetailsResponseOutput
}

type RecoveryPlanAutomationRunbookActionDetailsResponseArgs struct {
	FabricLocation pulumi.StringInput    `pulumi:"fabricLocation"`
	InstanceType   pulumi.StringInput    `pulumi:"instanceType"`
	RunbookId      pulumi.StringPtrInput `pulumi:"runbookId"`
	Timeout        pulumi.StringPtrInput `pulumi:"timeout"`
}

func (RecoveryPlanAutomationRunbookActionDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanAutomationRunbookActionDetailsResponse)(nil)).Elem()
}

func (i RecoveryPlanAutomationRunbookActionDetailsResponseArgs) ToRecoveryPlanAutomationRunbookActionDetailsResponseOutput() RecoveryPlanAutomationRunbookActionDetailsResponseOutput {
	return i.ToRecoveryPlanAutomationRunbookActionDetailsResponseOutputWithContext(context.Background())
}

func (i RecoveryPlanAutomationRunbookActionDetailsResponseArgs) ToRecoveryPlanAutomationRunbookActionDetailsResponseOutputWithContext(ctx context.Context) RecoveryPlanAutomationRunbookActionDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanAutomationRunbookActionDetailsResponseOutput)
}

type RecoveryPlanAutomationRunbookActionDetailsResponseOutput struct{ *pulumi.OutputState }

func (RecoveryPlanAutomationRunbookActionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanAutomationRunbookActionDetailsResponse)(nil)).Elem()
}

func (o RecoveryPlanAutomationRunbookActionDetailsResponseOutput) ToRecoveryPlanAutomationRunbookActionDetailsResponseOutput() RecoveryPlanAutomationRunbookActionDetailsResponseOutput {
	return o
}

func (o RecoveryPlanAutomationRunbookActionDetailsResponseOutput) ToRecoveryPlanAutomationRunbookActionDetailsResponseOutputWithContext(ctx context.Context) RecoveryPlanAutomationRunbookActionDetailsResponseOutput {
	return o
}

func (o RecoveryPlanAutomationRunbookActionDetailsResponseOutput) FabricLocation() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanAutomationRunbookActionDetailsResponse) string { return v.FabricLocation }).(pulumi.StringOutput)
}

func (o RecoveryPlanAutomationRunbookActionDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanAutomationRunbookActionDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o RecoveryPlanAutomationRunbookActionDetailsResponseOutput) RunbookId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanAutomationRunbookActionDetailsResponse) *string { return v.RunbookId }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanAutomationRunbookActionDetailsResponseOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanAutomationRunbookActionDetailsResponse) *string { return v.Timeout }).(pulumi.StringPtrOutput)
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





type RecoveryPlanGroupResponseInput interface {
	pulumi.Input

	ToRecoveryPlanGroupResponseOutput() RecoveryPlanGroupResponseOutput
	ToRecoveryPlanGroupResponseOutputWithContext(context.Context) RecoveryPlanGroupResponseOutput
}

type RecoveryPlanGroupResponseArgs struct {
	EndGroupActions           RecoveryPlanActionResponseArrayInput        `pulumi:"endGroupActions"`
	GroupType                 pulumi.StringInput                          `pulumi:"groupType"`
	ReplicationProtectedItems RecoveryPlanProtectedItemResponseArrayInput `pulumi:"replicationProtectedItems"`
	StartGroupActions         RecoveryPlanActionResponseArrayInput        `pulumi:"startGroupActions"`
}

func (RecoveryPlanGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanGroupResponse)(nil)).Elem()
}

func (i RecoveryPlanGroupResponseArgs) ToRecoveryPlanGroupResponseOutput() RecoveryPlanGroupResponseOutput {
	return i.ToRecoveryPlanGroupResponseOutputWithContext(context.Background())
}

func (i RecoveryPlanGroupResponseArgs) ToRecoveryPlanGroupResponseOutputWithContext(ctx context.Context) RecoveryPlanGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanGroupResponseOutput)
}





type RecoveryPlanGroupResponseArrayInput interface {
	pulumi.Input

	ToRecoveryPlanGroupResponseArrayOutput() RecoveryPlanGroupResponseArrayOutput
	ToRecoveryPlanGroupResponseArrayOutputWithContext(context.Context) RecoveryPlanGroupResponseArrayOutput
}

type RecoveryPlanGroupResponseArray []RecoveryPlanGroupResponseInput

func (RecoveryPlanGroupResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanGroupResponse)(nil)).Elem()
}

func (i RecoveryPlanGroupResponseArray) ToRecoveryPlanGroupResponseArrayOutput() RecoveryPlanGroupResponseArrayOutput {
	return i.ToRecoveryPlanGroupResponseArrayOutputWithContext(context.Background())
}

func (i RecoveryPlanGroupResponseArray) ToRecoveryPlanGroupResponseArrayOutputWithContext(ctx context.Context) RecoveryPlanGroupResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanGroupResponseArrayOutput)
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





type RecoveryPlanManualActionDetailsResponseInput interface {
	pulumi.Input

	ToRecoveryPlanManualActionDetailsResponseOutput() RecoveryPlanManualActionDetailsResponseOutput
	ToRecoveryPlanManualActionDetailsResponseOutputWithContext(context.Context) RecoveryPlanManualActionDetailsResponseOutput
}

type RecoveryPlanManualActionDetailsResponseArgs struct {
	Description  pulumi.StringPtrInput `pulumi:"description"`
	InstanceType pulumi.StringInput    `pulumi:"instanceType"`
}

func (RecoveryPlanManualActionDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanManualActionDetailsResponse)(nil)).Elem()
}

func (i RecoveryPlanManualActionDetailsResponseArgs) ToRecoveryPlanManualActionDetailsResponseOutput() RecoveryPlanManualActionDetailsResponseOutput {
	return i.ToRecoveryPlanManualActionDetailsResponseOutputWithContext(context.Background())
}

func (i RecoveryPlanManualActionDetailsResponseArgs) ToRecoveryPlanManualActionDetailsResponseOutputWithContext(ctx context.Context) RecoveryPlanManualActionDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanManualActionDetailsResponseOutput)
}

type RecoveryPlanManualActionDetailsResponseOutput struct{ *pulumi.OutputState }

func (RecoveryPlanManualActionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanManualActionDetailsResponse)(nil)).Elem()
}

func (o RecoveryPlanManualActionDetailsResponseOutput) ToRecoveryPlanManualActionDetailsResponseOutput() RecoveryPlanManualActionDetailsResponseOutput {
	return o
}

func (o RecoveryPlanManualActionDetailsResponseOutput) ToRecoveryPlanManualActionDetailsResponseOutputWithContext(ctx context.Context) RecoveryPlanManualActionDetailsResponseOutput {
	return o
}

func (o RecoveryPlanManualActionDetailsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanManualActionDetailsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanManualActionDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanManualActionDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
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





type RecoveryPlanPropertiesResponseInput interface {
	pulumi.Input

	ToRecoveryPlanPropertiesResponseOutput() RecoveryPlanPropertiesResponseOutput
	ToRecoveryPlanPropertiesResponseOutputWithContext(context.Context) RecoveryPlanPropertiesResponseOutput
}

type RecoveryPlanPropertiesResponseArgs struct {
	AllowedOperations                pulumi.StringArrayInput                `pulumi:"allowedOperations"`
	CurrentScenario                  CurrentScenarioDetailsResponsePtrInput `pulumi:"currentScenario"`
	CurrentScenarioStatus            pulumi.StringPtrInput                  `pulumi:"currentScenarioStatus"`
	CurrentScenarioStatusDescription pulumi.StringPtrInput                  `pulumi:"currentScenarioStatusDescription"`
	FailoverDeploymentModel          pulumi.StringPtrInput                  `pulumi:"failoverDeploymentModel"`
	FriendlyName                     pulumi.StringPtrInput                  `pulumi:"friendlyName"`
	Groups                           RecoveryPlanGroupResponseArrayInput    `pulumi:"groups"`
	LastPlannedFailoverTime          pulumi.StringPtrInput                  `pulumi:"lastPlannedFailoverTime"`
	LastTestFailoverTime             pulumi.StringPtrInput                  `pulumi:"lastTestFailoverTime"`
	LastUnplannedFailoverTime        pulumi.StringPtrInput                  `pulumi:"lastUnplannedFailoverTime"`
	PrimaryFabricFriendlyName        pulumi.StringPtrInput                  `pulumi:"primaryFabricFriendlyName"`
	PrimaryFabricId                  pulumi.StringPtrInput                  `pulumi:"primaryFabricId"`
	RecoveryFabricFriendlyName       pulumi.StringPtrInput                  `pulumi:"recoveryFabricFriendlyName"`
	RecoveryFabricId                 pulumi.StringPtrInput                  `pulumi:"recoveryFabricId"`
	ReplicationProviders             pulumi.StringArrayInput                `pulumi:"replicationProviders"`
}

func (RecoveryPlanPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanPropertiesResponse)(nil)).Elem()
}

func (i RecoveryPlanPropertiesResponseArgs) ToRecoveryPlanPropertiesResponseOutput() RecoveryPlanPropertiesResponseOutput {
	return i.ToRecoveryPlanPropertiesResponseOutputWithContext(context.Background())
}

func (i RecoveryPlanPropertiesResponseArgs) ToRecoveryPlanPropertiesResponseOutputWithContext(ctx context.Context) RecoveryPlanPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanPropertiesResponseOutput)
}

func (i RecoveryPlanPropertiesResponseArgs) ToRecoveryPlanPropertiesResponsePtrOutput() RecoveryPlanPropertiesResponsePtrOutput {
	return i.ToRecoveryPlanPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i RecoveryPlanPropertiesResponseArgs) ToRecoveryPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) RecoveryPlanPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanPropertiesResponseOutput).ToRecoveryPlanPropertiesResponsePtrOutputWithContext(ctx)
}









type RecoveryPlanPropertiesResponsePtrInput interface {
	pulumi.Input

	ToRecoveryPlanPropertiesResponsePtrOutput() RecoveryPlanPropertiesResponsePtrOutput
	ToRecoveryPlanPropertiesResponsePtrOutputWithContext(context.Context) RecoveryPlanPropertiesResponsePtrOutput
}

type recoveryPlanPropertiesResponsePtrType RecoveryPlanPropertiesResponseArgs

func RecoveryPlanPropertiesResponsePtr(v *RecoveryPlanPropertiesResponseArgs) RecoveryPlanPropertiesResponsePtrInput {
	return (*recoveryPlanPropertiesResponsePtrType)(v)
}

func (*recoveryPlanPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RecoveryPlanPropertiesResponse)(nil)).Elem()
}

func (i *recoveryPlanPropertiesResponsePtrType) ToRecoveryPlanPropertiesResponsePtrOutput() RecoveryPlanPropertiesResponsePtrOutput {
	return i.ToRecoveryPlanPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *recoveryPlanPropertiesResponsePtrType) ToRecoveryPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) RecoveryPlanPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanPropertiesResponsePtrOutput)
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

func (o RecoveryPlanPropertiesResponseOutput) ToRecoveryPlanPropertiesResponsePtrOutput() RecoveryPlanPropertiesResponsePtrOutput {
	return o.ToRecoveryPlanPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o RecoveryPlanPropertiesResponseOutput) ToRecoveryPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) RecoveryPlanPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecoveryPlanPropertiesResponse) *RecoveryPlanPropertiesResponse {
		return &v
	}).(RecoveryPlanPropertiesResponsePtrOutput)
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

type RecoveryPlanPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (RecoveryPlanPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecoveryPlanPropertiesResponse)(nil)).Elem()
}

func (o RecoveryPlanPropertiesResponsePtrOutput) ToRecoveryPlanPropertiesResponsePtrOutput() RecoveryPlanPropertiesResponsePtrOutput {
	return o
}

func (o RecoveryPlanPropertiesResponsePtrOutput) ToRecoveryPlanPropertiesResponsePtrOutputWithContext(ctx context.Context) RecoveryPlanPropertiesResponsePtrOutput {
	return o
}

func (o RecoveryPlanPropertiesResponsePtrOutput) Elem() RecoveryPlanPropertiesResponseOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) RecoveryPlanPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret RecoveryPlanPropertiesResponse
		return ret
	}).(RecoveryPlanPropertiesResponseOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) AllowedOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOperations
	}).(pulumi.StringArrayOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) CurrentScenario() CurrentScenarioDetailsResponsePtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *CurrentScenarioDetailsResponse {
		if v == nil {
			return nil
		}
		return v.CurrentScenario
	}).(CurrentScenarioDetailsResponsePtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) CurrentScenarioStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CurrentScenarioStatus
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) CurrentScenarioStatusDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CurrentScenarioStatusDescription
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) FailoverDeploymentModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FailoverDeploymentModel
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) Groups() RecoveryPlanGroupResponseArrayOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) []RecoveryPlanGroupResponse {
		if v == nil {
			return nil
		}
		return v.Groups
	}).(RecoveryPlanGroupResponseArrayOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) LastPlannedFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastPlannedFailoverTime
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) LastTestFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastTestFailoverTime
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) LastUnplannedFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastUnplannedFailoverTime
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) PrimaryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryFabricFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) PrimaryFabricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryFabricId
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) RecoveryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryFabricFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) RecoveryFabricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryFabricId
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryPlanPropertiesResponsePtrOutput) ReplicationProviders() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RecoveryPlanPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ReplicationProviders
	}).(pulumi.StringArrayOutput)
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





type RecoveryPlanProtectedItemResponseInput interface {
	pulumi.Input

	ToRecoveryPlanProtectedItemResponseOutput() RecoveryPlanProtectedItemResponseOutput
	ToRecoveryPlanProtectedItemResponseOutputWithContext(context.Context) RecoveryPlanProtectedItemResponseOutput
}

type RecoveryPlanProtectedItemResponseArgs struct {
	Id               pulumi.StringPtrInput `pulumi:"id"`
	VirtualMachineId pulumi.StringPtrInput `pulumi:"virtualMachineId"`
}

func (RecoveryPlanProtectedItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanProtectedItemResponse)(nil)).Elem()
}

func (i RecoveryPlanProtectedItemResponseArgs) ToRecoveryPlanProtectedItemResponseOutput() RecoveryPlanProtectedItemResponseOutput {
	return i.ToRecoveryPlanProtectedItemResponseOutputWithContext(context.Background())
}

func (i RecoveryPlanProtectedItemResponseArgs) ToRecoveryPlanProtectedItemResponseOutputWithContext(ctx context.Context) RecoveryPlanProtectedItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanProtectedItemResponseOutput)
}





type RecoveryPlanProtectedItemResponseArrayInput interface {
	pulumi.Input

	ToRecoveryPlanProtectedItemResponseArrayOutput() RecoveryPlanProtectedItemResponseArrayOutput
	ToRecoveryPlanProtectedItemResponseArrayOutputWithContext(context.Context) RecoveryPlanProtectedItemResponseArrayOutput
}

type RecoveryPlanProtectedItemResponseArray []RecoveryPlanProtectedItemResponseInput

func (RecoveryPlanProtectedItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RecoveryPlanProtectedItemResponse)(nil)).Elem()
}

func (i RecoveryPlanProtectedItemResponseArray) ToRecoveryPlanProtectedItemResponseArrayOutput() RecoveryPlanProtectedItemResponseArrayOutput {
	return i.ToRecoveryPlanProtectedItemResponseArrayOutputWithContext(context.Background())
}

func (i RecoveryPlanProtectedItemResponseArray) ToRecoveryPlanProtectedItemResponseArrayOutputWithContext(ctx context.Context) RecoveryPlanProtectedItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanProtectedItemResponseArrayOutput)
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





type RecoveryPlanScriptActionDetailsResponseInput interface {
	pulumi.Input

	ToRecoveryPlanScriptActionDetailsResponseOutput() RecoveryPlanScriptActionDetailsResponseOutput
	ToRecoveryPlanScriptActionDetailsResponseOutputWithContext(context.Context) RecoveryPlanScriptActionDetailsResponseOutput
}

type RecoveryPlanScriptActionDetailsResponseArgs struct {
	FabricLocation pulumi.StringInput    `pulumi:"fabricLocation"`
	InstanceType   pulumi.StringInput    `pulumi:"instanceType"`
	Path           pulumi.StringInput    `pulumi:"path"`
	Timeout        pulumi.StringPtrInput `pulumi:"timeout"`
}

func (RecoveryPlanScriptActionDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanScriptActionDetailsResponse)(nil)).Elem()
}

func (i RecoveryPlanScriptActionDetailsResponseArgs) ToRecoveryPlanScriptActionDetailsResponseOutput() RecoveryPlanScriptActionDetailsResponseOutput {
	return i.ToRecoveryPlanScriptActionDetailsResponseOutputWithContext(context.Background())
}

func (i RecoveryPlanScriptActionDetailsResponseArgs) ToRecoveryPlanScriptActionDetailsResponseOutputWithContext(ctx context.Context) RecoveryPlanScriptActionDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryPlanScriptActionDetailsResponseOutput)
}

type RecoveryPlanScriptActionDetailsResponseOutput struct{ *pulumi.OutputState }

func (RecoveryPlanScriptActionDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryPlanScriptActionDetailsResponse)(nil)).Elem()
}

func (o RecoveryPlanScriptActionDetailsResponseOutput) ToRecoveryPlanScriptActionDetailsResponseOutput() RecoveryPlanScriptActionDetailsResponseOutput {
	return o
}

func (o RecoveryPlanScriptActionDetailsResponseOutput) ToRecoveryPlanScriptActionDetailsResponseOutputWithContext(ctx context.Context) RecoveryPlanScriptActionDetailsResponseOutput {
	return o
}

func (o RecoveryPlanScriptActionDetailsResponseOutput) FabricLocation() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanScriptActionDetailsResponse) string { return v.FabricLocation }).(pulumi.StringOutput)
}

func (o RecoveryPlanScriptActionDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanScriptActionDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o RecoveryPlanScriptActionDetailsResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v RecoveryPlanScriptActionDetailsResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o RecoveryPlanScriptActionDetailsResponseOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RecoveryPlanScriptActionDetailsResponse) *string { return v.Timeout }).(pulumi.StringPtrOutput)
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





type RecoveryServicesProviderPropertiesResponseInput interface {
	pulumi.Input

	ToRecoveryServicesProviderPropertiesResponseOutput() RecoveryServicesProviderPropertiesResponseOutput
	ToRecoveryServicesProviderPropertiesResponseOutputWithContext(context.Context) RecoveryServicesProviderPropertiesResponseOutput
}

type RecoveryServicesProviderPropertiesResponseArgs struct {
	AllowedScenarios              pulumi.StringArrayInput                 `pulumi:"allowedScenarios"`
	AuthenticationIdentityDetails IdentityProviderDetailsResponsePtrInput `pulumi:"authenticationIdentityDetails"`
	ConnectionStatus              pulumi.StringPtrInput                   `pulumi:"connectionStatus"`
	DraIdentifier                 pulumi.StringPtrInput                   `pulumi:"draIdentifier"`
	FabricFriendlyName            pulumi.StringPtrInput                   `pulumi:"fabricFriendlyName"`
	FabricType                    pulumi.StringPtrInput                   `pulumi:"fabricType"`
	FriendlyName                  pulumi.StringPtrInput                   `pulumi:"friendlyName"`
	HealthErrorDetails            HealthErrorResponseArrayInput           `pulumi:"healthErrorDetails"`
	LastHeartBeat                 pulumi.StringPtrInput                   `pulumi:"lastHeartBeat"`
	ProtectedItemCount            pulumi.IntPtrInput                      `pulumi:"protectedItemCount"`
	ProviderVersion               pulumi.StringPtrInput                   `pulumi:"providerVersion"`
	ProviderVersionDetails        VersionDetailsResponsePtrInput          `pulumi:"providerVersionDetails"`
	ProviderVersionExpiryDate     pulumi.StringPtrInput                   `pulumi:"providerVersionExpiryDate"`
	ProviderVersionState          pulumi.StringPtrInput                   `pulumi:"providerVersionState"`
	ResourceAccessIdentityDetails IdentityProviderDetailsResponsePtrInput `pulumi:"resourceAccessIdentityDetails"`
	ServerVersion                 pulumi.StringPtrInput                   `pulumi:"serverVersion"`
}

func (RecoveryServicesProviderPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RecoveryServicesProviderPropertiesResponse)(nil)).Elem()
}

func (i RecoveryServicesProviderPropertiesResponseArgs) ToRecoveryServicesProviderPropertiesResponseOutput() RecoveryServicesProviderPropertiesResponseOutput {
	return i.ToRecoveryServicesProviderPropertiesResponseOutputWithContext(context.Background())
}

func (i RecoveryServicesProviderPropertiesResponseArgs) ToRecoveryServicesProviderPropertiesResponseOutputWithContext(ctx context.Context) RecoveryServicesProviderPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryServicesProviderPropertiesResponseOutput)
}

func (i RecoveryServicesProviderPropertiesResponseArgs) ToRecoveryServicesProviderPropertiesResponsePtrOutput() RecoveryServicesProviderPropertiesResponsePtrOutput {
	return i.ToRecoveryServicesProviderPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i RecoveryServicesProviderPropertiesResponseArgs) ToRecoveryServicesProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) RecoveryServicesProviderPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryServicesProviderPropertiesResponseOutput).ToRecoveryServicesProviderPropertiesResponsePtrOutputWithContext(ctx)
}









type RecoveryServicesProviderPropertiesResponsePtrInput interface {
	pulumi.Input

	ToRecoveryServicesProviderPropertiesResponsePtrOutput() RecoveryServicesProviderPropertiesResponsePtrOutput
	ToRecoveryServicesProviderPropertiesResponsePtrOutputWithContext(context.Context) RecoveryServicesProviderPropertiesResponsePtrOutput
}

type recoveryServicesProviderPropertiesResponsePtrType RecoveryServicesProviderPropertiesResponseArgs

func RecoveryServicesProviderPropertiesResponsePtr(v *RecoveryServicesProviderPropertiesResponseArgs) RecoveryServicesProviderPropertiesResponsePtrInput {
	return (*recoveryServicesProviderPropertiesResponsePtrType)(v)
}

func (*recoveryServicesProviderPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RecoveryServicesProviderPropertiesResponse)(nil)).Elem()
}

func (i *recoveryServicesProviderPropertiesResponsePtrType) ToRecoveryServicesProviderPropertiesResponsePtrOutput() RecoveryServicesProviderPropertiesResponsePtrOutput {
	return i.ToRecoveryServicesProviderPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *recoveryServicesProviderPropertiesResponsePtrType) ToRecoveryServicesProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) RecoveryServicesProviderPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RecoveryServicesProviderPropertiesResponsePtrOutput)
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

func (o RecoveryServicesProviderPropertiesResponseOutput) ToRecoveryServicesProviderPropertiesResponsePtrOutput() RecoveryServicesProviderPropertiesResponsePtrOutput {
	return o.ToRecoveryServicesProviderPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o RecoveryServicesProviderPropertiesResponseOutput) ToRecoveryServicesProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) RecoveryServicesProviderPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecoveryServicesProviderPropertiesResponse) *RecoveryServicesProviderPropertiesResponse {
		return &v
	}).(RecoveryServicesProviderPropertiesResponsePtrOutput)
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

type RecoveryServicesProviderPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (RecoveryServicesProviderPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecoveryServicesProviderPropertiesResponse)(nil)).Elem()
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) ToRecoveryServicesProviderPropertiesResponsePtrOutput() RecoveryServicesProviderPropertiesResponsePtrOutput {
	return o
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) ToRecoveryServicesProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) RecoveryServicesProviderPropertiesResponsePtrOutput {
	return o
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) Elem() RecoveryServicesProviderPropertiesResponseOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) RecoveryServicesProviderPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret RecoveryServicesProviderPropertiesResponse
		return ret
	}).(RecoveryServicesProviderPropertiesResponseOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) AllowedScenarios() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedScenarios
	}).(pulumi.StringArrayOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) AuthenticationIdentityDetails() IdentityProviderDetailsResponsePtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *IdentityProviderDetailsResponse {
		if v == nil {
			return nil
		}
		return v.AuthenticationIdentityDetails
	}).(IdentityProviderDetailsResponsePtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) ConnectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConnectionStatus
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) DraIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DraIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) FabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FabricFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) FabricType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FabricType
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) HealthErrorDetails() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) []HealthErrorResponse {
		if v == nil {
			return nil
		}
		return v.HealthErrorDetails
	}).(HealthErrorResponseArrayOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) LastHeartBeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastHeartBeat
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) ProtectedItemCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.ProtectedItemCount
	}).(pulumi.IntPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) ProviderVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProviderVersion
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) ProviderVersionDetails() VersionDetailsResponsePtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *VersionDetailsResponse {
		if v == nil {
			return nil
		}
		return v.ProviderVersionDetails
	}).(VersionDetailsResponsePtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) ProviderVersionExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProviderVersionExpiryDate
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) ProviderVersionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProviderVersionState
	}).(pulumi.StringPtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) ResourceAccessIdentityDetails() IdentityProviderDetailsResponsePtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *IdentityProviderDetailsResponse {
		if v == nil {
			return nil
		}
		return v.ResourceAccessIdentityDetails
	}).(IdentityProviderDetailsResponsePtrOutput)
}

func (o RecoveryServicesProviderPropertiesResponsePtrOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RecoveryServicesProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerVersion
	}).(pulumi.StringPtrOutput)
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





type ReplicationProtectedItemPropertiesResponseInput interface {
	pulumi.Input

	ToReplicationProtectedItemPropertiesResponseOutput() ReplicationProtectedItemPropertiesResponseOutput
	ToReplicationProtectedItemPropertiesResponseOutputWithContext(context.Context) ReplicationProtectedItemPropertiesResponseOutput
}

type ReplicationProtectedItemPropertiesResponseArgs struct {
	ActiveLocation                          pulumi.StringPtrInput                  `pulumi:"activeLocation"`
	AllowedOperations                       pulumi.StringArrayInput                `pulumi:"allowedOperations"`
	CurrentScenario                         CurrentScenarioDetailsResponsePtrInput `pulumi:"currentScenario"`
	FailoverHealth                          pulumi.StringPtrInput                  `pulumi:"failoverHealth"`
	FailoverRecoveryPointId                 pulumi.StringPtrInput                  `pulumi:"failoverRecoveryPointId"`
	FriendlyName                            pulumi.StringPtrInput                  `pulumi:"friendlyName"`
	HealthErrors                            HealthErrorResponseArrayInput          `pulumi:"healthErrors"`
	LastSuccessfulFailoverTime              pulumi.StringPtrInput                  `pulumi:"lastSuccessfulFailoverTime"`
	LastSuccessfulTestFailoverTime          pulumi.StringPtrInput                  `pulumi:"lastSuccessfulTestFailoverTime"`
	PolicyFriendlyName                      pulumi.StringPtrInput                  `pulumi:"policyFriendlyName"`
	PolicyId                                pulumi.StringPtrInput                  `pulumi:"policyId"`
	PrimaryFabricFriendlyName               pulumi.StringPtrInput                  `pulumi:"primaryFabricFriendlyName"`
	PrimaryFabricProvider                   pulumi.StringPtrInput                  `pulumi:"primaryFabricProvider"`
	PrimaryProtectionContainerFriendlyName  pulumi.StringPtrInput                  `pulumi:"primaryProtectionContainerFriendlyName"`
	ProtectableItemId                       pulumi.StringPtrInput                  `pulumi:"protectableItemId"`
	ProtectedItemType                       pulumi.StringPtrInput                  `pulumi:"protectedItemType"`
	ProtectionState                         pulumi.StringPtrInput                  `pulumi:"protectionState"`
	ProtectionStateDescription              pulumi.StringPtrInput                  `pulumi:"protectionStateDescription"`
	ProviderSpecificDetails                 pulumi.Input                           `pulumi:"providerSpecificDetails"`
	RecoveryContainerId                     pulumi.StringPtrInput                  `pulumi:"recoveryContainerId"`
	RecoveryFabricFriendlyName              pulumi.StringPtrInput                  `pulumi:"recoveryFabricFriendlyName"`
	RecoveryFabricId                        pulumi.StringPtrInput                  `pulumi:"recoveryFabricId"`
	RecoveryProtectionContainerFriendlyName pulumi.StringPtrInput                  `pulumi:"recoveryProtectionContainerFriendlyName"`
	RecoveryServicesProviderId              pulumi.StringPtrInput                  `pulumi:"recoveryServicesProviderId"`
	ReplicationHealth                       pulumi.StringPtrInput                  `pulumi:"replicationHealth"`
	TestFailoverState                       pulumi.StringPtrInput                  `pulumi:"testFailoverState"`
	TestFailoverStateDescription            pulumi.StringPtrInput                  `pulumi:"testFailoverStateDescription"`
}

func (ReplicationProtectedItemPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReplicationProtectedItemPropertiesResponse)(nil)).Elem()
}

func (i ReplicationProtectedItemPropertiesResponseArgs) ToReplicationProtectedItemPropertiesResponseOutput() ReplicationProtectedItemPropertiesResponseOutput {
	return i.ToReplicationProtectedItemPropertiesResponseOutputWithContext(context.Background())
}

func (i ReplicationProtectedItemPropertiesResponseArgs) ToReplicationProtectedItemPropertiesResponseOutputWithContext(ctx context.Context) ReplicationProtectedItemPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProtectedItemPropertiesResponseOutput)
}

func (i ReplicationProtectedItemPropertiesResponseArgs) ToReplicationProtectedItemPropertiesResponsePtrOutput() ReplicationProtectedItemPropertiesResponsePtrOutput {
	return i.ToReplicationProtectedItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ReplicationProtectedItemPropertiesResponseArgs) ToReplicationProtectedItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ReplicationProtectedItemPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProtectedItemPropertiesResponseOutput).ToReplicationProtectedItemPropertiesResponsePtrOutputWithContext(ctx)
}









type ReplicationProtectedItemPropertiesResponsePtrInput interface {
	pulumi.Input

	ToReplicationProtectedItemPropertiesResponsePtrOutput() ReplicationProtectedItemPropertiesResponsePtrOutput
	ToReplicationProtectedItemPropertiesResponsePtrOutputWithContext(context.Context) ReplicationProtectedItemPropertiesResponsePtrOutput
}

type replicationProtectedItemPropertiesResponsePtrType ReplicationProtectedItemPropertiesResponseArgs

func ReplicationProtectedItemPropertiesResponsePtr(v *ReplicationProtectedItemPropertiesResponseArgs) ReplicationProtectedItemPropertiesResponsePtrInput {
	return (*replicationProtectedItemPropertiesResponsePtrType)(v)
}

func (*replicationProtectedItemPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProtectedItemPropertiesResponse)(nil)).Elem()
}

func (i *replicationProtectedItemPropertiesResponsePtrType) ToReplicationProtectedItemPropertiesResponsePtrOutput() ReplicationProtectedItemPropertiesResponsePtrOutput {
	return i.ToReplicationProtectedItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *replicationProtectedItemPropertiesResponsePtrType) ToReplicationProtectedItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ReplicationProtectedItemPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReplicationProtectedItemPropertiesResponsePtrOutput)
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

func (o ReplicationProtectedItemPropertiesResponseOutput) ToReplicationProtectedItemPropertiesResponsePtrOutput() ReplicationProtectedItemPropertiesResponsePtrOutput {
	return o.ToReplicationProtectedItemPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ReplicationProtectedItemPropertiesResponseOutput) ToReplicationProtectedItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ReplicationProtectedItemPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReplicationProtectedItemPropertiesResponse) *ReplicationProtectedItemPropertiesResponse {
		return &v
	}).(ReplicationProtectedItemPropertiesResponsePtrOutput)
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

type ReplicationProtectedItemPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ReplicationProtectedItemPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReplicationProtectedItemPropertiesResponse)(nil)).Elem()
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) ToReplicationProtectedItemPropertiesResponsePtrOutput() ReplicationProtectedItemPropertiesResponsePtrOutput {
	return o
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) ToReplicationProtectedItemPropertiesResponsePtrOutputWithContext(ctx context.Context) ReplicationProtectedItemPropertiesResponsePtrOutput {
	return o
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) Elem() ReplicationProtectedItemPropertiesResponseOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) ReplicationProtectedItemPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ReplicationProtectedItemPropertiesResponse
		return ret
	}).(ReplicationProtectedItemPropertiesResponseOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) ActiveLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActiveLocation
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) AllowedOperations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOperations
	}).(pulumi.StringArrayOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) CurrentScenario() CurrentScenarioDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *CurrentScenarioDetailsResponse {
		if v == nil {
			return nil
		}
		return v.CurrentScenario
	}).(CurrentScenarioDetailsResponsePtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) FailoverHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FailoverHealth
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) FailoverRecoveryPointId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FailoverRecoveryPointId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) HealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) []HealthErrorResponse {
		if v == nil {
			return nil
		}
		return v.HealthErrors
	}).(HealthErrorResponseArrayOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) LastSuccessfulFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastSuccessfulFailoverTime
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) LastSuccessfulTestFailoverTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastSuccessfulTestFailoverTime
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) PolicyFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PolicyId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) PrimaryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryFabricFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) PrimaryFabricProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryFabricProvider
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) PrimaryProtectionContainerFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrimaryProtectionContainerFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) ProtectableItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProtectableItemId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) ProtectedItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProtectedItemType
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProtectionState
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) ProtectionStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProtectionStateDescription
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) ProviderSpecificDetails() pulumi.AnyOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.ProviderSpecificDetails
	}).(pulumi.AnyOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) RecoveryContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryContainerId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) RecoveryFabricFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryFabricFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) RecoveryFabricId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryFabricId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) RecoveryProtectionContainerFriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryProtectionContainerFriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) RecoveryServicesProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RecoveryServicesProviderId
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) ReplicationHealth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReplicationHealth
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) TestFailoverState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TestFailoverState
	}).(pulumi.StringPtrOutput)
}

func (o ReplicationProtectedItemPropertiesResponsePtrOutput) TestFailoverStateDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReplicationProtectedItemPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TestFailoverStateDescription
	}).(pulumi.StringPtrOutput)
}

type RetentionVolumeResponse struct {
	CapacityInBytes     *float64 `pulumi:"capacityInBytes"`
	FreeSpaceInBytes    *float64 `pulumi:"freeSpaceInBytes"`
	ThresholdPercentage *int     `pulumi:"thresholdPercentage"`
	VolumeName          *string  `pulumi:"volumeName"`
}





type RetentionVolumeResponseInput interface {
	pulumi.Input

	ToRetentionVolumeResponseOutput() RetentionVolumeResponseOutput
	ToRetentionVolumeResponseOutputWithContext(context.Context) RetentionVolumeResponseOutput
}

type RetentionVolumeResponseArgs struct {
	CapacityInBytes     pulumi.Float64PtrInput `pulumi:"capacityInBytes"`
	FreeSpaceInBytes    pulumi.Float64PtrInput `pulumi:"freeSpaceInBytes"`
	ThresholdPercentage pulumi.IntPtrInput     `pulumi:"thresholdPercentage"`
	VolumeName          pulumi.StringPtrInput  `pulumi:"volumeName"`
}

func (RetentionVolumeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionVolumeResponse)(nil)).Elem()
}

func (i RetentionVolumeResponseArgs) ToRetentionVolumeResponseOutput() RetentionVolumeResponseOutput {
	return i.ToRetentionVolumeResponseOutputWithContext(context.Background())
}

func (i RetentionVolumeResponseArgs) ToRetentionVolumeResponseOutputWithContext(ctx context.Context) RetentionVolumeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionVolumeResponseOutput)
}





type RetentionVolumeResponseArrayInput interface {
	pulumi.Input

	ToRetentionVolumeResponseArrayOutput() RetentionVolumeResponseArrayOutput
	ToRetentionVolumeResponseArrayOutputWithContext(context.Context) RetentionVolumeResponseArrayOutput
}

type RetentionVolumeResponseArray []RetentionVolumeResponseInput

func (RetentionVolumeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RetentionVolumeResponse)(nil)).Elem()
}

func (i RetentionVolumeResponseArray) ToRetentionVolumeResponseArrayOutput() RetentionVolumeResponseArrayOutput {
	return i.ToRetentionVolumeResponseArrayOutputWithContext(context.Background())
}

func (i RetentionVolumeResponseArray) ToRetentionVolumeResponseArrayOutputWithContext(ctx context.Context) RetentionVolumeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionVolumeResponseArrayOutput)
}

type RetentionVolumeResponseOutput struct{ *pulumi.OutputState }

func (RetentionVolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionVolumeResponse)(nil)).Elem()
}

func (o RetentionVolumeResponseOutput) ToRetentionVolumeResponseOutput() RetentionVolumeResponseOutput {
	return o
}

func (o RetentionVolumeResponseOutput) ToRetentionVolumeResponseOutputWithContext(ctx context.Context) RetentionVolumeResponseOutput {
	return o
}

func (o RetentionVolumeResponseOutput) CapacityInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RetentionVolumeResponse) *float64 { return v.CapacityInBytes }).(pulumi.Float64PtrOutput)
}

func (o RetentionVolumeResponseOutput) FreeSpaceInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RetentionVolumeResponse) *float64 { return v.FreeSpaceInBytes }).(pulumi.Float64PtrOutput)
}

func (o RetentionVolumeResponseOutput) ThresholdPercentage() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionVolumeResponse) *int { return v.ThresholdPercentage }).(pulumi.IntPtrOutput)
}

func (o RetentionVolumeResponseOutput) VolumeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionVolumeResponse) *string { return v.VolumeName }).(pulumi.StringPtrOutput)
}

type RetentionVolumeResponseArrayOutput struct{ *pulumi.OutputState }

func (RetentionVolumeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RetentionVolumeResponse)(nil)).Elem()
}

func (o RetentionVolumeResponseArrayOutput) ToRetentionVolumeResponseArrayOutput() RetentionVolumeResponseArrayOutput {
	return o
}

func (o RetentionVolumeResponseArrayOutput) ToRetentionVolumeResponseArrayOutputWithContext(ctx context.Context) RetentionVolumeResponseArrayOutput {
	return o
}

func (o RetentionVolumeResponseArrayOutput) Index(i pulumi.IntInput) RetentionVolumeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RetentionVolumeResponse {
		return vs[0].([]RetentionVolumeResponse)[vs[1].(int)]
	}).(RetentionVolumeResponseOutput)
}

type RoleAssignmentResponse struct {
	Id               *string `pulumi:"id"`
	Name             *string `pulumi:"name"`
	PrincipalId      *string `pulumi:"principalId"`
	RoleDefinitionId *string `pulumi:"roleDefinitionId"`
	Scope            *string `pulumi:"scope"`
}





type RoleAssignmentResponseInput interface {
	pulumi.Input

	ToRoleAssignmentResponseOutput() RoleAssignmentResponseOutput
	ToRoleAssignmentResponseOutputWithContext(context.Context) RoleAssignmentResponseOutput
}

type RoleAssignmentResponseArgs struct {
	Id               pulumi.StringPtrInput `pulumi:"id"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	PrincipalId      pulumi.StringPtrInput `pulumi:"principalId"`
	RoleDefinitionId pulumi.StringPtrInput `pulumi:"roleDefinitionId"`
	Scope            pulumi.StringPtrInput `pulumi:"scope"`
}

func (RoleAssignmentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssignmentResponse)(nil)).Elem()
}

func (i RoleAssignmentResponseArgs) ToRoleAssignmentResponseOutput() RoleAssignmentResponseOutput {
	return i.ToRoleAssignmentResponseOutputWithContext(context.Background())
}

func (i RoleAssignmentResponseArgs) ToRoleAssignmentResponseOutputWithContext(ctx context.Context) RoleAssignmentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentResponseOutput)
}





type RoleAssignmentResponseArrayInput interface {
	pulumi.Input

	ToRoleAssignmentResponseArrayOutput() RoleAssignmentResponseArrayOutput
	ToRoleAssignmentResponseArrayOutputWithContext(context.Context) RoleAssignmentResponseArrayOutput
}

type RoleAssignmentResponseArray []RoleAssignmentResponseInput

func (RoleAssignmentResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoleAssignmentResponse)(nil)).Elem()
}

func (i RoleAssignmentResponseArray) ToRoleAssignmentResponseArrayOutput() RoleAssignmentResponseArrayOutput {
	return i.ToRoleAssignmentResponseArrayOutputWithContext(context.Background())
}

func (i RoleAssignmentResponseArray) ToRoleAssignmentResponseArrayOutputWithContext(ctx context.Context) RoleAssignmentResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleAssignmentResponseArrayOutput)
}

type RoleAssignmentResponseOutput struct{ *pulumi.OutputState }

func (RoleAssignmentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoleAssignmentResponse)(nil)).Elem()
}

func (o RoleAssignmentResponseOutput) ToRoleAssignmentResponseOutput() RoleAssignmentResponseOutput {
	return o
}

func (o RoleAssignmentResponseOutput) ToRoleAssignmentResponseOutputWithContext(ctx context.Context) RoleAssignmentResponseOutput {
	return o
}

func (o RoleAssignmentResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleAssignmentResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleAssignmentResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleAssignmentResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentResponseOutput) RoleDefinitionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleAssignmentResponse) *string { return v.RoleDefinitionId }).(pulumi.StringPtrOutput)
}

func (o RoleAssignmentResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoleAssignmentResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

type RoleAssignmentResponseArrayOutput struct{ *pulumi.OutputState }

func (RoleAssignmentResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoleAssignmentResponse)(nil)).Elem()
}

func (o RoleAssignmentResponseArrayOutput) ToRoleAssignmentResponseArrayOutput() RoleAssignmentResponseArrayOutput {
	return o
}

func (o RoleAssignmentResponseArrayOutput) ToRoleAssignmentResponseArrayOutputWithContext(ctx context.Context) RoleAssignmentResponseArrayOutput {
	return o
}

func (o RoleAssignmentResponseArrayOutput) Index(i pulumi.IntInput) RoleAssignmentResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoleAssignmentResponse {
		return vs[0].([]RoleAssignmentResponse)[vs[1].(int)]
	}).(RoleAssignmentResponseOutput)
}

type RunAsAccountResponse struct {
	AccountId   *string `pulumi:"accountId"`
	AccountName *string `pulumi:"accountName"`
}





type RunAsAccountResponseInput interface {
	pulumi.Input

	ToRunAsAccountResponseOutput() RunAsAccountResponseOutput
	ToRunAsAccountResponseOutputWithContext(context.Context) RunAsAccountResponseOutput
}

type RunAsAccountResponseArgs struct {
	AccountId   pulumi.StringPtrInput `pulumi:"accountId"`
	AccountName pulumi.StringPtrInput `pulumi:"accountName"`
}

func (RunAsAccountResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsAccountResponse)(nil)).Elem()
}

func (i RunAsAccountResponseArgs) ToRunAsAccountResponseOutput() RunAsAccountResponseOutput {
	return i.ToRunAsAccountResponseOutputWithContext(context.Background())
}

func (i RunAsAccountResponseArgs) ToRunAsAccountResponseOutputWithContext(ctx context.Context) RunAsAccountResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsAccountResponseOutput)
}





type RunAsAccountResponseArrayInput interface {
	pulumi.Input

	ToRunAsAccountResponseArrayOutput() RunAsAccountResponseArrayOutput
	ToRunAsAccountResponseArrayOutputWithContext(context.Context) RunAsAccountResponseArrayOutput
}

type RunAsAccountResponseArray []RunAsAccountResponseInput

func (RunAsAccountResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RunAsAccountResponse)(nil)).Elem()
}

func (i RunAsAccountResponseArray) ToRunAsAccountResponseArrayOutput() RunAsAccountResponseArrayOutput {
	return i.ToRunAsAccountResponseArrayOutputWithContext(context.Background())
}

func (i RunAsAccountResponseArray) ToRunAsAccountResponseArrayOutputWithContext(ctx context.Context) RunAsAccountResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsAccountResponseArrayOutput)
}

type RunAsAccountResponseOutput struct{ *pulumi.OutputState }

func (RunAsAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsAccountResponse)(nil)).Elem()
}

func (o RunAsAccountResponseOutput) ToRunAsAccountResponseOutput() RunAsAccountResponseOutput {
	return o
}

func (o RunAsAccountResponseOutput) ToRunAsAccountResponseOutputWithContext(ctx context.Context) RunAsAccountResponseOutput {
	return o
}

func (o RunAsAccountResponseOutput) AccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunAsAccountResponse) *string { return v.AccountId }).(pulumi.StringPtrOutput)
}

func (o RunAsAccountResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunAsAccountResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

type RunAsAccountResponseArrayOutput struct{ *pulumi.OutputState }

func (RunAsAccountResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RunAsAccountResponse)(nil)).Elem()
}

func (o RunAsAccountResponseArrayOutput) ToRunAsAccountResponseArrayOutput() RunAsAccountResponseArrayOutput {
	return o
}

func (o RunAsAccountResponseArrayOutput) ToRunAsAccountResponseArrayOutputWithContext(ctx context.Context) RunAsAccountResponseArrayOutput {
	return o
}

func (o RunAsAccountResponseArrayOutput) Index(i pulumi.IntInput) RunAsAccountResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RunAsAccountResponse {
		return vs[0].([]RunAsAccountResponse)[vs[1].(int)]
	}).(RunAsAccountResponseOutput)
}

type SanEnableProtectionInput struct {
	InstanceType *string `pulumi:"instanceType"`
}





type SanEnableProtectionInputInput interface {
	pulumi.Input

	ToSanEnableProtectionInputOutput() SanEnableProtectionInputOutput
	ToSanEnableProtectionInputOutputWithContext(context.Context) SanEnableProtectionInputOutput
}

type SanEnableProtectionInputArgs struct {
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
}

func (SanEnableProtectionInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SanEnableProtectionInput)(nil)).Elem()
}

func (i SanEnableProtectionInputArgs) ToSanEnableProtectionInputOutput() SanEnableProtectionInputOutput {
	return i.ToSanEnableProtectionInputOutputWithContext(context.Background())
}

func (i SanEnableProtectionInputArgs) ToSanEnableProtectionInputOutputWithContext(ctx context.Context) SanEnableProtectionInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SanEnableProtectionInputOutput)
}

type SanEnableProtectionInputOutput struct{ *pulumi.OutputState }

func (SanEnableProtectionInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SanEnableProtectionInput)(nil)).Elem()
}

func (o SanEnableProtectionInputOutput) ToSanEnableProtectionInputOutput() SanEnableProtectionInputOutput {
	return o
}

func (o SanEnableProtectionInputOutput) ToSanEnableProtectionInputOutputWithContext(ctx context.Context) SanEnableProtectionInputOutput {
	return o
}

func (o SanEnableProtectionInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SanEnableProtectionInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

type StorageClassificationMappingPropertiesResponse struct {
	TargetStorageClassificationId *string `pulumi:"targetStorageClassificationId"`
}





type StorageClassificationMappingPropertiesResponseInput interface {
	pulumi.Input

	ToStorageClassificationMappingPropertiesResponseOutput() StorageClassificationMappingPropertiesResponseOutput
	ToStorageClassificationMappingPropertiesResponseOutputWithContext(context.Context) StorageClassificationMappingPropertiesResponseOutput
}

type StorageClassificationMappingPropertiesResponseArgs struct {
	TargetStorageClassificationId pulumi.StringPtrInput `pulumi:"targetStorageClassificationId"`
}

func (StorageClassificationMappingPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageClassificationMappingPropertiesResponse)(nil)).Elem()
}

func (i StorageClassificationMappingPropertiesResponseArgs) ToStorageClassificationMappingPropertiesResponseOutput() StorageClassificationMappingPropertiesResponseOutput {
	return i.ToStorageClassificationMappingPropertiesResponseOutputWithContext(context.Background())
}

func (i StorageClassificationMappingPropertiesResponseArgs) ToStorageClassificationMappingPropertiesResponseOutputWithContext(ctx context.Context) StorageClassificationMappingPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassificationMappingPropertiesResponseOutput)
}

func (i StorageClassificationMappingPropertiesResponseArgs) ToStorageClassificationMappingPropertiesResponsePtrOutput() StorageClassificationMappingPropertiesResponsePtrOutput {
	return i.ToStorageClassificationMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i StorageClassificationMappingPropertiesResponseArgs) ToStorageClassificationMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageClassificationMappingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassificationMappingPropertiesResponseOutput).ToStorageClassificationMappingPropertiesResponsePtrOutputWithContext(ctx)
}









type StorageClassificationMappingPropertiesResponsePtrInput interface {
	pulumi.Input

	ToStorageClassificationMappingPropertiesResponsePtrOutput() StorageClassificationMappingPropertiesResponsePtrOutput
	ToStorageClassificationMappingPropertiesResponsePtrOutputWithContext(context.Context) StorageClassificationMappingPropertiesResponsePtrOutput
}

type storageClassificationMappingPropertiesResponsePtrType StorageClassificationMappingPropertiesResponseArgs

func StorageClassificationMappingPropertiesResponsePtr(v *StorageClassificationMappingPropertiesResponseArgs) StorageClassificationMappingPropertiesResponsePtrInput {
	return (*storageClassificationMappingPropertiesResponsePtrType)(v)
}

func (*storageClassificationMappingPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageClassificationMappingPropertiesResponse)(nil)).Elem()
}

func (i *storageClassificationMappingPropertiesResponsePtrType) ToStorageClassificationMappingPropertiesResponsePtrOutput() StorageClassificationMappingPropertiesResponsePtrOutput {
	return i.ToStorageClassificationMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *storageClassificationMappingPropertiesResponsePtrType) ToStorageClassificationMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageClassificationMappingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageClassificationMappingPropertiesResponsePtrOutput)
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

func (o StorageClassificationMappingPropertiesResponseOutput) ToStorageClassificationMappingPropertiesResponsePtrOutput() StorageClassificationMappingPropertiesResponsePtrOutput {
	return o.ToStorageClassificationMappingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o StorageClassificationMappingPropertiesResponseOutput) ToStorageClassificationMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageClassificationMappingPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageClassificationMappingPropertiesResponse) *StorageClassificationMappingPropertiesResponse {
		return &v
	}).(StorageClassificationMappingPropertiesResponsePtrOutput)
}

func (o StorageClassificationMappingPropertiesResponseOutput) TargetStorageClassificationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageClassificationMappingPropertiesResponse) *string { return v.TargetStorageClassificationId }).(pulumi.StringPtrOutput)
}

type StorageClassificationMappingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageClassificationMappingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageClassificationMappingPropertiesResponse)(nil)).Elem()
}

func (o StorageClassificationMappingPropertiesResponsePtrOutput) ToStorageClassificationMappingPropertiesResponsePtrOutput() StorageClassificationMappingPropertiesResponsePtrOutput {
	return o
}

func (o StorageClassificationMappingPropertiesResponsePtrOutput) ToStorageClassificationMappingPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageClassificationMappingPropertiesResponsePtrOutput {
	return o
}

func (o StorageClassificationMappingPropertiesResponsePtrOutput) Elem() StorageClassificationMappingPropertiesResponseOutput {
	return o.ApplyT(func(v *StorageClassificationMappingPropertiesResponse) StorageClassificationMappingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StorageClassificationMappingPropertiesResponse
		return ret
	}).(StorageClassificationMappingPropertiesResponseOutput)
}

func (o StorageClassificationMappingPropertiesResponsePtrOutput) TargetStorageClassificationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageClassificationMappingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetStorageClassificationId
	}).(pulumi.StringPtrOutput)
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





type VCenterPropertiesResponseInput interface {
	pulumi.Input

	ToVCenterPropertiesResponseOutput() VCenterPropertiesResponseOutput
	ToVCenterPropertiesResponseOutputWithContext(context.Context) VCenterPropertiesResponseOutput
}

type VCenterPropertiesResponseArgs struct {
	DiscoveryStatus       pulumi.StringPtrInput         `pulumi:"discoveryStatus"`
	FabricArmResourceName pulumi.StringPtrInput         `pulumi:"fabricArmResourceName"`
	FriendlyName          pulumi.StringPtrInput         `pulumi:"friendlyName"`
	HealthErrors          HealthErrorResponseArrayInput `pulumi:"healthErrors"`
	InfrastructureId      pulumi.StringPtrInput         `pulumi:"infrastructureId"`
	InternalId            pulumi.StringPtrInput         `pulumi:"internalId"`
	IpAddress             pulumi.StringPtrInput         `pulumi:"ipAddress"`
	LastHeartbeat         pulumi.StringPtrInput         `pulumi:"lastHeartbeat"`
	Port                  pulumi.StringPtrInput         `pulumi:"port"`
	ProcessServerId       pulumi.StringPtrInput         `pulumi:"processServerId"`
	RunAsAccountId        pulumi.StringPtrInput         `pulumi:"runAsAccountId"`
}

func (VCenterPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VCenterPropertiesResponse)(nil)).Elem()
}

func (i VCenterPropertiesResponseArgs) ToVCenterPropertiesResponseOutput() VCenterPropertiesResponseOutput {
	return i.ToVCenterPropertiesResponseOutputWithContext(context.Background())
}

func (i VCenterPropertiesResponseArgs) ToVCenterPropertiesResponseOutputWithContext(ctx context.Context) VCenterPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VCenterPropertiesResponseOutput)
}

func (i VCenterPropertiesResponseArgs) ToVCenterPropertiesResponsePtrOutput() VCenterPropertiesResponsePtrOutput {
	return i.ToVCenterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i VCenterPropertiesResponseArgs) ToVCenterPropertiesResponsePtrOutputWithContext(ctx context.Context) VCenterPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VCenterPropertiesResponseOutput).ToVCenterPropertiesResponsePtrOutputWithContext(ctx)
}









type VCenterPropertiesResponsePtrInput interface {
	pulumi.Input

	ToVCenterPropertiesResponsePtrOutput() VCenterPropertiesResponsePtrOutput
	ToVCenterPropertiesResponsePtrOutputWithContext(context.Context) VCenterPropertiesResponsePtrOutput
}

type vcenterPropertiesResponsePtrType VCenterPropertiesResponseArgs

func VCenterPropertiesResponsePtr(v *VCenterPropertiesResponseArgs) VCenterPropertiesResponsePtrInput {
	return (*vcenterPropertiesResponsePtrType)(v)
}

func (*vcenterPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VCenterPropertiesResponse)(nil)).Elem()
}

func (i *vcenterPropertiesResponsePtrType) ToVCenterPropertiesResponsePtrOutput() VCenterPropertiesResponsePtrOutput {
	return i.ToVCenterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *vcenterPropertiesResponsePtrType) ToVCenterPropertiesResponsePtrOutputWithContext(ctx context.Context) VCenterPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VCenterPropertiesResponsePtrOutput)
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

func (o VCenterPropertiesResponseOutput) ToVCenterPropertiesResponsePtrOutput() VCenterPropertiesResponsePtrOutput {
	return o.ToVCenterPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o VCenterPropertiesResponseOutput) ToVCenterPropertiesResponsePtrOutputWithContext(ctx context.Context) VCenterPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VCenterPropertiesResponse) *VCenterPropertiesResponse {
		return &v
	}).(VCenterPropertiesResponsePtrOutput)
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

type VCenterPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VCenterPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VCenterPropertiesResponse)(nil)).Elem()
}

func (o VCenterPropertiesResponsePtrOutput) ToVCenterPropertiesResponsePtrOutput() VCenterPropertiesResponsePtrOutput {
	return o
}

func (o VCenterPropertiesResponsePtrOutput) ToVCenterPropertiesResponsePtrOutputWithContext(ctx context.Context) VCenterPropertiesResponsePtrOutput {
	return o
}

func (o VCenterPropertiesResponsePtrOutput) Elem() VCenterPropertiesResponseOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) VCenterPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VCenterPropertiesResponse
		return ret
	}).(VCenterPropertiesResponseOutput)
}

func (o VCenterPropertiesResponsePtrOutput) DiscoveryStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiscoveryStatus
	}).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponsePtrOutput) FabricArmResourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FabricArmResourceName
	}).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponsePtrOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FriendlyName
	}).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponsePtrOutput) HealthErrors() HealthErrorResponseArrayOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) []HealthErrorResponse {
		if v == nil {
			return nil
		}
		return v.HealthErrors
	}).(HealthErrorResponseArrayOutput)
}

func (o VCenterPropertiesResponsePtrOutput) InfrastructureId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.InfrastructureId
	}).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponsePtrOutput) InternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.InternalId
	}).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponsePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponsePtrOutput) LastHeartbeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastHeartbeat
	}).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponsePtrOutput) Port() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponsePtrOutput) ProcessServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProcessServerId
	}).(pulumi.StringPtrOutput)
}

func (o VCenterPropertiesResponsePtrOutput) RunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VCenterPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RunAsAccountId
	}).(pulumi.StringPtrOutput)
}

type VMNicDetailsResponse struct {
	EnableAcceleratedNetworkingOnRecovery *bool   `pulumi:"enableAcceleratedNetworkingOnRecovery"`
	IpAddressType                         *string `pulumi:"ipAddressType"`
	NicId                                 *string `pulumi:"nicId"`
	PrimaryNicStaticIPAddress             *string `pulumi:"primaryNicStaticIPAddress"`
	RecoveryNicIpAddressType              *string `pulumi:"recoveryNicIpAddressType"`
	RecoveryVMNetworkId                   *string `pulumi:"recoveryVMNetworkId"`
	RecoveryVMSubnetName                  *string `pulumi:"recoveryVMSubnetName"`
	ReplicaNicId                          *string `pulumi:"replicaNicId"`
	ReplicaNicStaticIPAddress             *string `pulumi:"replicaNicStaticIPAddress"`
	SelectionType                         *string `pulumi:"selectionType"`
	SourceNicArmId                        *string `pulumi:"sourceNicArmId"`
	VMNetworkName                         *string `pulumi:"vMNetworkName"`
	VMSubnetName                          *string `pulumi:"vMSubnetName"`
}





type VMNicDetailsResponseInput interface {
	pulumi.Input

	ToVMNicDetailsResponseOutput() VMNicDetailsResponseOutput
	ToVMNicDetailsResponseOutputWithContext(context.Context) VMNicDetailsResponseOutput
}

type VMNicDetailsResponseArgs struct {
	EnableAcceleratedNetworkingOnRecovery pulumi.BoolPtrInput   `pulumi:"enableAcceleratedNetworkingOnRecovery"`
	IpAddressType                         pulumi.StringPtrInput `pulumi:"ipAddressType"`
	NicId                                 pulumi.StringPtrInput `pulumi:"nicId"`
	PrimaryNicStaticIPAddress             pulumi.StringPtrInput `pulumi:"primaryNicStaticIPAddress"`
	RecoveryNicIpAddressType              pulumi.StringPtrInput `pulumi:"recoveryNicIpAddressType"`
	RecoveryVMNetworkId                   pulumi.StringPtrInput `pulumi:"recoveryVMNetworkId"`
	RecoveryVMSubnetName                  pulumi.StringPtrInput `pulumi:"recoveryVMSubnetName"`
	ReplicaNicId                          pulumi.StringPtrInput `pulumi:"replicaNicId"`
	ReplicaNicStaticIPAddress             pulumi.StringPtrInput `pulumi:"replicaNicStaticIPAddress"`
	SelectionType                         pulumi.StringPtrInput `pulumi:"selectionType"`
	SourceNicArmId                        pulumi.StringPtrInput `pulumi:"sourceNicArmId"`
	VMNetworkName                         pulumi.StringPtrInput `pulumi:"vMNetworkName"`
	VMSubnetName                          pulumi.StringPtrInput `pulumi:"vMSubnetName"`
}

func (VMNicDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMNicDetailsResponse)(nil)).Elem()
}

func (i VMNicDetailsResponseArgs) ToVMNicDetailsResponseOutput() VMNicDetailsResponseOutput {
	return i.ToVMNicDetailsResponseOutputWithContext(context.Background())
}

func (i VMNicDetailsResponseArgs) ToVMNicDetailsResponseOutputWithContext(ctx context.Context) VMNicDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMNicDetailsResponseOutput)
}





type VMNicDetailsResponseArrayInput interface {
	pulumi.Input

	ToVMNicDetailsResponseArrayOutput() VMNicDetailsResponseArrayOutput
	ToVMNicDetailsResponseArrayOutputWithContext(context.Context) VMNicDetailsResponseArrayOutput
}

type VMNicDetailsResponseArray []VMNicDetailsResponseInput

func (VMNicDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMNicDetailsResponse)(nil)).Elem()
}

func (i VMNicDetailsResponseArray) ToVMNicDetailsResponseArrayOutput() VMNicDetailsResponseArrayOutput {
	return i.ToVMNicDetailsResponseArrayOutputWithContext(context.Background())
}

func (i VMNicDetailsResponseArray) ToVMNicDetailsResponseArrayOutputWithContext(ctx context.Context) VMNicDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMNicDetailsResponseArrayOutput)
}

type VMNicDetailsResponseOutput struct{ *pulumi.OutputState }

func (VMNicDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMNicDetailsResponse)(nil)).Elem()
}

func (o VMNicDetailsResponseOutput) ToVMNicDetailsResponseOutput() VMNicDetailsResponseOutput {
	return o
}

func (o VMNicDetailsResponseOutput) ToVMNicDetailsResponseOutputWithContext(ctx context.Context) VMNicDetailsResponseOutput {
	return o
}

func (o VMNicDetailsResponseOutput) EnableAcceleratedNetworkingOnRecovery() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *bool { return v.EnableAcceleratedNetworkingOnRecovery }).(pulumi.BoolPtrOutput)
}

func (o VMNicDetailsResponseOutput) IpAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.IpAddressType }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) NicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.NicId }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) PrimaryNicStaticIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.PrimaryNicStaticIPAddress }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) RecoveryNicIpAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.RecoveryNicIpAddressType }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) RecoveryVMNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.RecoveryVMNetworkId }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) RecoveryVMSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.RecoveryVMSubnetName }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) ReplicaNicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.ReplicaNicId }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) ReplicaNicStaticIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.ReplicaNicStaticIPAddress }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) SelectionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.SelectionType }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) SourceNicArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.SourceNicArmId }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) VMNetworkName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.VMNetworkName }).(pulumi.StringPtrOutput)
}

func (o VMNicDetailsResponseOutput) VMSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMNicDetailsResponse) *string { return v.VMSubnetName }).(pulumi.StringPtrOutput)
}

type VMNicDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (VMNicDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMNicDetailsResponse)(nil)).Elem()
}

func (o VMNicDetailsResponseArrayOutput) ToVMNicDetailsResponseArrayOutput() VMNicDetailsResponseArrayOutput {
	return o
}

func (o VMNicDetailsResponseArrayOutput) ToVMNicDetailsResponseArrayOutputWithContext(ctx context.Context) VMNicDetailsResponseArrayOutput {
	return o
}

func (o VMNicDetailsResponseArrayOutput) Index(i pulumi.IntInput) VMNicDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMNicDetailsResponse {
		return vs[0].([]VMNicDetailsResponse)[vs[1].(int)]
	}).(VMNicDetailsResponseOutput)
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





type VMwareCbtContainerMappingInputInput interface {
	pulumi.Input

	ToVMwareCbtContainerMappingInputOutput() VMwareCbtContainerMappingInputOutput
	ToVMwareCbtContainerMappingInputOutputWithContext(context.Context) VMwareCbtContainerMappingInputOutput
}

type VMwareCbtContainerMappingInputArgs struct {
	InstanceType                         pulumi.StringPtrInput `pulumi:"instanceType"`
	KeyVaultId                           pulumi.StringInput    `pulumi:"keyVaultId"`
	KeyVaultUri                          pulumi.StringInput    `pulumi:"keyVaultUri"`
	ServiceBusConnectionStringSecretName pulumi.StringInput    `pulumi:"serviceBusConnectionStringSecretName"`
	StorageAccountId                     pulumi.StringInput    `pulumi:"storageAccountId"`
	StorageAccountSasSecretName          pulumi.StringInput    `pulumi:"storageAccountSasSecretName"`
	TargetLocation                       pulumi.StringInput    `pulumi:"targetLocation"`
}

func (VMwareCbtContainerMappingInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtContainerMappingInput)(nil)).Elem()
}

func (i VMwareCbtContainerMappingInputArgs) ToVMwareCbtContainerMappingInputOutput() VMwareCbtContainerMappingInputOutput {
	return i.ToVMwareCbtContainerMappingInputOutputWithContext(context.Background())
}

func (i VMwareCbtContainerMappingInputArgs) ToVMwareCbtContainerMappingInputOutputWithContext(ctx context.Context) VMwareCbtContainerMappingInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtContainerMappingInputOutput)
}

type VMwareCbtContainerMappingInputOutput struct{ *pulumi.OutputState }

func (VMwareCbtContainerMappingInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtContainerMappingInput)(nil)).Elem()
}

func (o VMwareCbtContainerMappingInputOutput) ToVMwareCbtContainerMappingInputOutput() VMwareCbtContainerMappingInputOutput {
	return o
}

func (o VMwareCbtContainerMappingInputOutput) ToVMwareCbtContainerMappingInputOutputWithContext(ctx context.Context) VMwareCbtContainerMappingInputOutput {
	return o
}

func (o VMwareCbtContainerMappingInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtContainerMappingInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtContainerMappingInputOutput) KeyVaultId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtContainerMappingInput) string { return v.KeyVaultId }).(pulumi.StringOutput)
}

func (o VMwareCbtContainerMappingInputOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtContainerMappingInput) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o VMwareCbtContainerMappingInputOutput) ServiceBusConnectionStringSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtContainerMappingInput) string { return v.ServiceBusConnectionStringSecretName }).(pulumi.StringOutput)
}

func (o VMwareCbtContainerMappingInputOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtContainerMappingInput) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtContainerMappingInputOutput) StorageAccountSasSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtContainerMappingInput) string { return v.StorageAccountSasSecretName }).(pulumi.StringOutput)
}

func (o VMwareCbtContainerMappingInputOutput) TargetLocation() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtContainerMappingInput) string { return v.TargetLocation }).(pulumi.StringOutput)
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
	SnapshotRunAsAccountId                string               `pulumi:"snapshotRunAsAccountId"`
	TargetAvailabilitySetId               *string              `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                *string              `pulumi:"targetAvailabilityZone"`
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
	PerformAutoResync                     pulumi.StringPtrInput        `pulumi:"performAutoResync"`
	SnapshotRunAsAccountId                pulumi.StringInput           `pulumi:"snapshotRunAsAccountId"`
	TargetAvailabilitySetId               pulumi.StringPtrInput        `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                pulumi.StringPtrInput        `pulumi:"targetAvailabilityZone"`
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

func (i VMwareCbtEnableMigrationInputArgs) ToVMwareCbtEnableMigrationInputPtrOutput() VMwareCbtEnableMigrationInputPtrOutput {
	return i.ToVMwareCbtEnableMigrationInputPtrOutputWithContext(context.Background())
}

func (i VMwareCbtEnableMigrationInputArgs) ToVMwareCbtEnableMigrationInputPtrOutputWithContext(ctx context.Context) VMwareCbtEnableMigrationInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtEnableMigrationInputOutput).ToVMwareCbtEnableMigrationInputPtrOutputWithContext(ctx)
}









type VMwareCbtEnableMigrationInputPtrInput interface {
	pulumi.Input

	ToVMwareCbtEnableMigrationInputPtrOutput() VMwareCbtEnableMigrationInputPtrOutput
	ToVMwareCbtEnableMigrationInputPtrOutputWithContext(context.Context) VMwareCbtEnableMigrationInputPtrOutput
}

type vmwareCbtEnableMigrationInputPtrType VMwareCbtEnableMigrationInputArgs

func VMwareCbtEnableMigrationInputPtr(v *VMwareCbtEnableMigrationInputArgs) VMwareCbtEnableMigrationInputPtrInput {
	return (*vmwareCbtEnableMigrationInputPtrType)(v)
}

func (*vmwareCbtEnableMigrationInputPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VMwareCbtEnableMigrationInput)(nil)).Elem()
}

func (i *vmwareCbtEnableMigrationInputPtrType) ToVMwareCbtEnableMigrationInputPtrOutput() VMwareCbtEnableMigrationInputPtrOutput {
	return i.ToVMwareCbtEnableMigrationInputPtrOutputWithContext(context.Background())
}

func (i *vmwareCbtEnableMigrationInputPtrType) ToVMwareCbtEnableMigrationInputPtrOutputWithContext(ctx context.Context) VMwareCbtEnableMigrationInputPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtEnableMigrationInputPtrOutput)
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

func (o VMwareCbtEnableMigrationInputOutput) ToVMwareCbtEnableMigrationInputPtrOutput() VMwareCbtEnableMigrationInputPtrOutput {
	return o.ToVMwareCbtEnableMigrationInputPtrOutputWithContext(context.Background())
}

func (o VMwareCbtEnableMigrationInputOutput) ToVMwareCbtEnableMigrationInputPtrOutputWithContext(ctx context.Context) VMwareCbtEnableMigrationInputPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VMwareCbtEnableMigrationInput) *VMwareCbtEnableMigrationInput {
		return &v
	}).(VMwareCbtEnableMigrationInputPtrOutput)
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

func (o VMwareCbtEnableMigrationInputOutput) SnapshotRunAsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtEnableMigrationInput) string { return v.SnapshotRunAsAccountId }).(pulumi.StringOutput)
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

type VMwareCbtEnableMigrationInputPtrOutput struct{ *pulumi.OutputState }

func (VMwareCbtEnableMigrationInputPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMwareCbtEnableMigrationInput)(nil)).Elem()
}

func (o VMwareCbtEnableMigrationInputPtrOutput) ToVMwareCbtEnableMigrationInputPtrOutput() VMwareCbtEnableMigrationInputPtrOutput {
	return o
}

func (o VMwareCbtEnableMigrationInputPtrOutput) ToVMwareCbtEnableMigrationInputPtrOutputWithContext(ctx context.Context) VMwareCbtEnableMigrationInputPtrOutput {
	return o
}

func (o VMwareCbtEnableMigrationInputPtrOutput) Elem() VMwareCbtEnableMigrationInputOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) VMwareCbtEnableMigrationInput {
		if v != nil {
			return *v
		}
		var ret VMwareCbtEnableMigrationInput
		return ret
	}).(VMwareCbtEnableMigrationInputOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) DataMoverRunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return &v.DataMoverRunAsAccountId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) DisksToInclude() VMwareCbtDiskInputArrayOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) []VMwareCbtDiskInput {
		if v == nil {
			return nil
		}
		return v.DisksToInclude
	}).(VMwareCbtDiskInputArrayOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return &v.InstanceType
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return v.LicenseType
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) PerformAutoResync() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return v.PerformAutoResync
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) SnapshotRunAsAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return &v.SnapshotRunAsAccountId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) TargetAvailabilitySetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return v.TargetAvailabilitySetId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) TargetAvailabilityZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return v.TargetAvailabilityZone
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) TargetBootDiagnosticsStorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return v.TargetBootDiagnosticsStorageAccountId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) TargetNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return &v.TargetNetworkId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) TargetResourceGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return &v.TargetResourceGroupId
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) TargetSubnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return v.TargetSubnetName
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) TargetVmName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return v.TargetVmName
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) TargetVmSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return v.TargetVmSize
	}).(pulumi.StringPtrOutput)
}

func (o VMwareCbtEnableMigrationInputPtrOutput) VmwareMachineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtEnableMigrationInput) *string {
		if v == nil {
			return nil
		}
		return &v.VmwareMachineId
	}).(pulumi.StringPtrOutput)
}

type VMwareCbtMigrationDetailsResponse struct {
	DataMoverRunAsAccountId               string                                  `pulumi:"dataMoverRunAsAccountId"`
	FirmwareType                          string                                  `pulumi:"firmwareType"`
	InitialSeedingProgressPercentage      int                                     `pulumi:"initialSeedingProgressPercentage"`
	InstanceType                          string                                  `pulumi:"instanceType"`
	LastRecoveryPointId                   string                                  `pulumi:"lastRecoveryPointId"`
	LastRecoveryPointReceived             string                                  `pulumi:"lastRecoveryPointReceived"`
	LicenseType                           *string                                 `pulumi:"licenseType"`
	MigrationProgressPercentage           int                                     `pulumi:"migrationProgressPercentage"`
	MigrationRecoveryPointId              string                                  `pulumi:"migrationRecoveryPointId"`
	OsType                                string                                  `pulumi:"osType"`
	PerformAutoResync                     *string                                 `pulumi:"performAutoResync"`
	ProtectedDisks                        []VMwareCbtProtectedDiskDetailsResponse `pulumi:"protectedDisks"`
	ResyncProgressPercentage              int                                     `pulumi:"resyncProgressPercentage"`
	ResyncRequired                        string                                  `pulumi:"resyncRequired"`
	ResyncState                           string                                  `pulumi:"resyncState"`
	SnapshotRunAsAccountId                string                                  `pulumi:"snapshotRunAsAccountId"`
	TargetAvailabilitySetId               *string                                 `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                *string                                 `pulumi:"targetAvailabilityZone"`
	TargetBootDiagnosticsStorageAccountId *string                                 `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetGeneration                      string                                  `pulumi:"targetGeneration"`
	TargetLocation                        string                                  `pulumi:"targetLocation"`
	TargetNetworkId                       *string                                 `pulumi:"targetNetworkId"`
	TargetResourceGroupId                 *string                                 `pulumi:"targetResourceGroupId"`
	TargetVmName                          *string                                 `pulumi:"targetVmName"`
	TargetVmSize                          *string                                 `pulumi:"targetVmSize"`
	VmNics                                []VMwareCbtNicDetailsResponse           `pulumi:"vmNics"`
	VmwareMachineId                       string                                  `pulumi:"vmwareMachineId"`
}





type VMwareCbtMigrationDetailsResponseInput interface {
	pulumi.Input

	ToVMwareCbtMigrationDetailsResponseOutput() VMwareCbtMigrationDetailsResponseOutput
	ToVMwareCbtMigrationDetailsResponseOutputWithContext(context.Context) VMwareCbtMigrationDetailsResponseOutput
}

type VMwareCbtMigrationDetailsResponseArgs struct {
	DataMoverRunAsAccountId               pulumi.StringInput                              `pulumi:"dataMoverRunAsAccountId"`
	FirmwareType                          pulumi.StringInput                              `pulumi:"firmwareType"`
	InitialSeedingProgressPercentage      pulumi.IntInput                                 `pulumi:"initialSeedingProgressPercentage"`
	InstanceType                          pulumi.StringInput                              `pulumi:"instanceType"`
	LastRecoveryPointId                   pulumi.StringInput                              `pulumi:"lastRecoveryPointId"`
	LastRecoveryPointReceived             pulumi.StringInput                              `pulumi:"lastRecoveryPointReceived"`
	LicenseType                           pulumi.StringPtrInput                           `pulumi:"licenseType"`
	MigrationProgressPercentage           pulumi.IntInput                                 `pulumi:"migrationProgressPercentage"`
	MigrationRecoveryPointId              pulumi.StringInput                              `pulumi:"migrationRecoveryPointId"`
	OsType                                pulumi.StringInput                              `pulumi:"osType"`
	PerformAutoResync                     pulumi.StringPtrInput                           `pulumi:"performAutoResync"`
	ProtectedDisks                        VMwareCbtProtectedDiskDetailsResponseArrayInput `pulumi:"protectedDisks"`
	ResyncProgressPercentage              pulumi.IntInput                                 `pulumi:"resyncProgressPercentage"`
	ResyncRequired                        pulumi.StringInput                              `pulumi:"resyncRequired"`
	ResyncState                           pulumi.StringInput                              `pulumi:"resyncState"`
	SnapshotRunAsAccountId                pulumi.StringInput                              `pulumi:"snapshotRunAsAccountId"`
	TargetAvailabilitySetId               pulumi.StringPtrInput                           `pulumi:"targetAvailabilitySetId"`
	TargetAvailabilityZone                pulumi.StringPtrInput                           `pulumi:"targetAvailabilityZone"`
	TargetBootDiagnosticsStorageAccountId pulumi.StringPtrInput                           `pulumi:"targetBootDiagnosticsStorageAccountId"`
	TargetGeneration                      pulumi.StringInput                              `pulumi:"targetGeneration"`
	TargetLocation                        pulumi.StringInput                              `pulumi:"targetLocation"`
	TargetNetworkId                       pulumi.StringPtrInput                           `pulumi:"targetNetworkId"`
	TargetResourceGroupId                 pulumi.StringPtrInput                           `pulumi:"targetResourceGroupId"`
	TargetVmName                          pulumi.StringPtrInput                           `pulumi:"targetVmName"`
	TargetVmSize                          pulumi.StringPtrInput                           `pulumi:"targetVmSize"`
	VmNics                                VMwareCbtNicDetailsResponseArrayInput           `pulumi:"vmNics"`
	VmwareMachineId                       pulumi.StringInput                              `pulumi:"vmwareMachineId"`
}

func (VMwareCbtMigrationDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtMigrationDetailsResponse)(nil)).Elem()
}

func (i VMwareCbtMigrationDetailsResponseArgs) ToVMwareCbtMigrationDetailsResponseOutput() VMwareCbtMigrationDetailsResponseOutput {
	return i.ToVMwareCbtMigrationDetailsResponseOutputWithContext(context.Background())
}

func (i VMwareCbtMigrationDetailsResponseArgs) ToVMwareCbtMigrationDetailsResponseOutputWithContext(ctx context.Context) VMwareCbtMigrationDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtMigrationDetailsResponseOutput)
}

func (i VMwareCbtMigrationDetailsResponseArgs) ToVMwareCbtMigrationDetailsResponsePtrOutput() VMwareCbtMigrationDetailsResponsePtrOutput {
	return i.ToVMwareCbtMigrationDetailsResponsePtrOutputWithContext(context.Background())
}

func (i VMwareCbtMigrationDetailsResponseArgs) ToVMwareCbtMigrationDetailsResponsePtrOutputWithContext(ctx context.Context) VMwareCbtMigrationDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtMigrationDetailsResponseOutput).ToVMwareCbtMigrationDetailsResponsePtrOutputWithContext(ctx)
}









type VMwareCbtMigrationDetailsResponsePtrInput interface {
	pulumi.Input

	ToVMwareCbtMigrationDetailsResponsePtrOutput() VMwareCbtMigrationDetailsResponsePtrOutput
	ToVMwareCbtMigrationDetailsResponsePtrOutputWithContext(context.Context) VMwareCbtMigrationDetailsResponsePtrOutput
}

type vmwareCbtMigrationDetailsResponsePtrType VMwareCbtMigrationDetailsResponseArgs

func VMwareCbtMigrationDetailsResponsePtr(v *VMwareCbtMigrationDetailsResponseArgs) VMwareCbtMigrationDetailsResponsePtrInput {
	return (*vmwareCbtMigrationDetailsResponsePtrType)(v)
}

func (*vmwareCbtMigrationDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VMwareCbtMigrationDetailsResponse)(nil)).Elem()
}

func (i *vmwareCbtMigrationDetailsResponsePtrType) ToVMwareCbtMigrationDetailsResponsePtrOutput() VMwareCbtMigrationDetailsResponsePtrOutput {
	return i.ToVMwareCbtMigrationDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *vmwareCbtMigrationDetailsResponsePtrType) ToVMwareCbtMigrationDetailsResponsePtrOutputWithContext(ctx context.Context) VMwareCbtMigrationDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtMigrationDetailsResponsePtrOutput)
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

func (o VMwareCbtMigrationDetailsResponseOutput) ToVMwareCbtMigrationDetailsResponsePtrOutput() VMwareCbtMigrationDetailsResponsePtrOutput {
	return o.ToVMwareCbtMigrationDetailsResponsePtrOutputWithContext(context.Background())
}

func (o VMwareCbtMigrationDetailsResponseOutput) ToVMwareCbtMigrationDetailsResponsePtrOutputWithContext(ctx context.Context) VMwareCbtMigrationDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VMwareCbtMigrationDetailsResponse) *VMwareCbtMigrationDetailsResponse {
		return &v
	}).(VMwareCbtMigrationDetailsResponsePtrOutput)
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

func (o VMwareCbtMigrationDetailsResponseOutput) ResyncProgressPercentage() pulumi.IntOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) int { return v.ResyncProgressPercentage }).(pulumi.IntOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) ResyncRequired() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.ResyncRequired }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) ResyncState() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.ResyncState }).(pulumi.StringOutput)
}

func (o VMwareCbtMigrationDetailsResponseOutput) SnapshotRunAsAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.SnapshotRunAsAccountId }).(pulumi.StringOutput)
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

func (o VMwareCbtMigrationDetailsResponseOutput) TargetGeneration() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtMigrationDetailsResponse) string { return v.TargetGeneration }).(pulumi.StringOutput)
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

func (o VMwareCbtMigrationDetailsResponsePtrOutput) ResyncState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareCbtMigrationDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResyncState
	}).(pulumi.StringPtrOutput)
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





type VMwareCbtNicDetailsResponseInput interface {
	pulumi.Input

	ToVMwareCbtNicDetailsResponseOutput() VMwareCbtNicDetailsResponseOutput
	ToVMwareCbtNicDetailsResponseOutputWithContext(context.Context) VMwareCbtNicDetailsResponseOutput
}

type VMwareCbtNicDetailsResponseArgs struct {
	IsPrimaryNic           pulumi.StringPtrInput `pulumi:"isPrimaryNic"`
	IsSelectedForMigration pulumi.StringPtrInput `pulumi:"isSelectedForMigration"`
	NicId                  pulumi.StringInput    `pulumi:"nicId"`
	SourceIPAddress        pulumi.StringInput    `pulumi:"sourceIPAddress"`
	SourceIPAddressType    pulumi.StringInput    `pulumi:"sourceIPAddressType"`
	SourceNetworkId        pulumi.StringInput    `pulumi:"sourceNetworkId"`
	TargetIPAddress        pulumi.StringPtrInput `pulumi:"targetIPAddress"`
	TargetIPAddressType    pulumi.StringPtrInput `pulumi:"targetIPAddressType"`
	TargetSubnetName       pulumi.StringPtrInput `pulumi:"targetSubnetName"`
}

func (VMwareCbtNicDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtNicDetailsResponse)(nil)).Elem()
}

func (i VMwareCbtNicDetailsResponseArgs) ToVMwareCbtNicDetailsResponseOutput() VMwareCbtNicDetailsResponseOutput {
	return i.ToVMwareCbtNicDetailsResponseOutputWithContext(context.Background())
}

func (i VMwareCbtNicDetailsResponseArgs) ToVMwareCbtNicDetailsResponseOutputWithContext(ctx context.Context) VMwareCbtNicDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtNicDetailsResponseOutput)
}





type VMwareCbtNicDetailsResponseArrayInput interface {
	pulumi.Input

	ToVMwareCbtNicDetailsResponseArrayOutput() VMwareCbtNicDetailsResponseArrayOutput
	ToVMwareCbtNicDetailsResponseArrayOutputWithContext(context.Context) VMwareCbtNicDetailsResponseArrayOutput
}

type VMwareCbtNicDetailsResponseArray []VMwareCbtNicDetailsResponseInput

func (VMwareCbtNicDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMwareCbtNicDetailsResponse)(nil)).Elem()
}

func (i VMwareCbtNicDetailsResponseArray) ToVMwareCbtNicDetailsResponseArrayOutput() VMwareCbtNicDetailsResponseArrayOutput {
	return i.ToVMwareCbtNicDetailsResponseArrayOutputWithContext(context.Background())
}

func (i VMwareCbtNicDetailsResponseArray) ToVMwareCbtNicDetailsResponseArrayOutputWithContext(ctx context.Context) VMwareCbtNicDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtNicDetailsResponseArrayOutput)
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





type VMwareCbtPolicyCreationInputInput interface {
	pulumi.Input

	ToVMwareCbtPolicyCreationInputOutput() VMwareCbtPolicyCreationInputOutput
	ToVMwareCbtPolicyCreationInputOutputWithContext(context.Context) VMwareCbtPolicyCreationInputOutput
}

type VMwareCbtPolicyCreationInputArgs struct {
	AppConsistentFrequencyInMinutes   pulumi.IntPtrInput    `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes pulumi.IntPtrInput    `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      pulumi.StringPtrInput `pulumi:"instanceType"`
	RecoveryPointHistoryInMinutes     pulumi.IntPtrInput    `pulumi:"recoveryPointHistoryInMinutes"`
}

func (VMwareCbtPolicyCreationInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtPolicyCreationInput)(nil)).Elem()
}

func (i VMwareCbtPolicyCreationInputArgs) ToVMwareCbtPolicyCreationInputOutput() VMwareCbtPolicyCreationInputOutput {
	return i.ToVMwareCbtPolicyCreationInputOutputWithContext(context.Background())
}

func (i VMwareCbtPolicyCreationInputArgs) ToVMwareCbtPolicyCreationInputOutputWithContext(ctx context.Context) VMwareCbtPolicyCreationInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtPolicyCreationInputOutput)
}

type VMwareCbtPolicyCreationInputOutput struct{ *pulumi.OutputState }

func (VMwareCbtPolicyCreationInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtPolicyCreationInput)(nil)).Elem()
}

func (o VMwareCbtPolicyCreationInputOutput) ToVMwareCbtPolicyCreationInputOutput() VMwareCbtPolicyCreationInputOutput {
	return o
}

func (o VMwareCbtPolicyCreationInputOutput) ToVMwareCbtPolicyCreationInputOutputWithContext(ctx context.Context) VMwareCbtPolicyCreationInputOutput {
	return o
}

func (o VMwareCbtPolicyCreationInputOutput) AppConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VMwareCbtPolicyCreationInput) *int { return v.AppConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o VMwareCbtPolicyCreationInputOutput) CrashConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VMwareCbtPolicyCreationInput) *int { return v.CrashConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o VMwareCbtPolicyCreationInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareCbtPolicyCreationInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o VMwareCbtPolicyCreationInputOutput) RecoveryPointHistoryInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VMwareCbtPolicyCreationInput) *int { return v.RecoveryPointHistoryInMinutes }).(pulumi.IntPtrOutput)
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
	SeedManagedDiskId              string  `pulumi:"seedManagedDiskId"`
	TargetManagedDiskId            string  `pulumi:"targetManagedDiskId"`
}





type VMwareCbtProtectedDiskDetailsResponseInput interface {
	pulumi.Input

	ToVMwareCbtProtectedDiskDetailsResponseOutput() VMwareCbtProtectedDiskDetailsResponseOutput
	ToVMwareCbtProtectedDiskDetailsResponseOutputWithContext(context.Context) VMwareCbtProtectedDiskDetailsResponseOutput
}

type VMwareCbtProtectedDiskDetailsResponseArgs struct {
	CapacityInBytes                pulumi.Float64Input   `pulumi:"capacityInBytes"`
	DiskEncryptionSetId            pulumi.StringInput    `pulumi:"diskEncryptionSetId"`
	DiskId                         pulumi.StringInput    `pulumi:"diskId"`
	DiskName                       pulumi.StringInput    `pulumi:"diskName"`
	DiskPath                       pulumi.StringInput    `pulumi:"diskPath"`
	DiskType                       pulumi.StringPtrInput `pulumi:"diskType"`
	IsOSDisk                       pulumi.StringInput    `pulumi:"isOSDisk"`
	LogStorageAccountId            pulumi.StringInput    `pulumi:"logStorageAccountId"`
	LogStorageAccountSasSecretName pulumi.StringInput    `pulumi:"logStorageAccountSasSecretName"`
	SeedManagedDiskId              pulumi.StringInput    `pulumi:"seedManagedDiskId"`
	TargetManagedDiskId            pulumi.StringInput    `pulumi:"targetManagedDiskId"`
}

func (VMwareCbtProtectedDiskDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtProtectedDiskDetailsResponse)(nil)).Elem()
}

func (i VMwareCbtProtectedDiskDetailsResponseArgs) ToVMwareCbtProtectedDiskDetailsResponseOutput() VMwareCbtProtectedDiskDetailsResponseOutput {
	return i.ToVMwareCbtProtectedDiskDetailsResponseOutputWithContext(context.Background())
}

func (i VMwareCbtProtectedDiskDetailsResponseArgs) ToVMwareCbtProtectedDiskDetailsResponseOutputWithContext(ctx context.Context) VMwareCbtProtectedDiskDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtProtectedDiskDetailsResponseOutput)
}





type VMwareCbtProtectedDiskDetailsResponseArrayInput interface {
	pulumi.Input

	ToVMwareCbtProtectedDiskDetailsResponseArrayOutput() VMwareCbtProtectedDiskDetailsResponseArrayOutput
	ToVMwareCbtProtectedDiskDetailsResponseArrayOutputWithContext(context.Context) VMwareCbtProtectedDiskDetailsResponseArrayOutput
}

type VMwareCbtProtectedDiskDetailsResponseArray []VMwareCbtProtectedDiskDetailsResponseInput

func (VMwareCbtProtectedDiskDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMwareCbtProtectedDiskDetailsResponse)(nil)).Elem()
}

func (i VMwareCbtProtectedDiskDetailsResponseArray) ToVMwareCbtProtectedDiskDetailsResponseArrayOutput() VMwareCbtProtectedDiskDetailsResponseArrayOutput {
	return i.ToVMwareCbtProtectedDiskDetailsResponseArrayOutputWithContext(context.Background())
}

func (i VMwareCbtProtectedDiskDetailsResponseArray) ToVMwareCbtProtectedDiskDetailsResponseArrayOutputWithContext(ctx context.Context) VMwareCbtProtectedDiskDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtProtectedDiskDetailsResponseArrayOutput)
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





type VMwareCbtProtectionContainerMappingDetailsResponseInput interface {
	pulumi.Input

	ToVMwareCbtProtectionContainerMappingDetailsResponseOutput() VMwareCbtProtectionContainerMappingDetailsResponseOutput
	ToVMwareCbtProtectionContainerMappingDetailsResponseOutputWithContext(context.Context) VMwareCbtProtectionContainerMappingDetailsResponseOutput
}

type VMwareCbtProtectionContainerMappingDetailsResponseArgs struct {
	InstanceType                         pulumi.StringInput `pulumi:"instanceType"`
	KeyVaultId                           pulumi.StringInput `pulumi:"keyVaultId"`
	KeyVaultUri                          pulumi.StringInput `pulumi:"keyVaultUri"`
	ServiceBusConnectionStringSecretName pulumi.StringInput `pulumi:"serviceBusConnectionStringSecretName"`
	StorageAccountId                     pulumi.StringInput `pulumi:"storageAccountId"`
	StorageAccountSasSecretName          pulumi.StringInput `pulumi:"storageAccountSasSecretName"`
	TargetLocation                       pulumi.StringInput `pulumi:"targetLocation"`
}

func (VMwareCbtProtectionContainerMappingDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtProtectionContainerMappingDetailsResponse)(nil)).Elem()
}

func (i VMwareCbtProtectionContainerMappingDetailsResponseArgs) ToVMwareCbtProtectionContainerMappingDetailsResponseOutput() VMwareCbtProtectionContainerMappingDetailsResponseOutput {
	return i.ToVMwareCbtProtectionContainerMappingDetailsResponseOutputWithContext(context.Background())
}

func (i VMwareCbtProtectionContainerMappingDetailsResponseArgs) ToVMwareCbtProtectionContainerMappingDetailsResponseOutputWithContext(ctx context.Context) VMwareCbtProtectionContainerMappingDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareCbtProtectionContainerMappingDetailsResponseOutput)
}

type VMwareCbtProtectionContainerMappingDetailsResponseOutput struct{ *pulumi.OutputState }

func (VMwareCbtProtectionContainerMappingDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareCbtProtectionContainerMappingDetailsResponse)(nil)).Elem()
}

func (o VMwareCbtProtectionContainerMappingDetailsResponseOutput) ToVMwareCbtProtectionContainerMappingDetailsResponseOutput() VMwareCbtProtectionContainerMappingDetailsResponseOutput {
	return o
}

func (o VMwareCbtProtectionContainerMappingDetailsResponseOutput) ToVMwareCbtProtectionContainerMappingDetailsResponseOutputWithContext(ctx context.Context) VMwareCbtProtectionContainerMappingDetailsResponseOutput {
	return o
}

func (o VMwareCbtProtectionContainerMappingDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectionContainerMappingDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectionContainerMappingDetailsResponseOutput) KeyVaultId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectionContainerMappingDetailsResponse) string { return v.KeyVaultId }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectionContainerMappingDetailsResponseOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectionContainerMappingDetailsResponse) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectionContainerMappingDetailsResponseOutput) ServiceBusConnectionStringSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectionContainerMappingDetailsResponse) string {
		return v.ServiceBusConnectionStringSecretName
	}).(pulumi.StringOutput)
}

func (o VMwareCbtProtectionContainerMappingDetailsResponseOutput) StorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectionContainerMappingDetailsResponse) string { return v.StorageAccountId }).(pulumi.StringOutput)
}

func (o VMwareCbtProtectionContainerMappingDetailsResponseOutput) StorageAccountSasSecretName() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectionContainerMappingDetailsResponse) string {
		return v.StorageAccountSasSecretName
	}).(pulumi.StringOutput)
}

func (o VMwareCbtProtectionContainerMappingDetailsResponseOutput) TargetLocation() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareCbtProtectionContainerMappingDetailsResponse) string { return v.TargetLocation }).(pulumi.StringOutput)
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





type VMwareDetailsResponseInput interface {
	pulumi.Input

	ToVMwareDetailsResponseOutput() VMwareDetailsResponseOutput
	ToVMwareDetailsResponseOutputWithContext(context.Context) VMwareDetailsResponseOutput
}

type VMwareDetailsResponseArgs struct {
	AgentCount                 pulumi.StringPtrInput                `pulumi:"agentCount"`
	AgentExpiryDate            pulumi.StringPtrInput                `pulumi:"agentExpiryDate"`
	AgentVersion               pulumi.StringPtrInput                `pulumi:"agentVersion"`
	AgentVersionDetails        VersionDetailsResponsePtrInput       `pulumi:"agentVersionDetails"`
	AvailableMemoryInBytes     pulumi.Float64PtrInput               `pulumi:"availableMemoryInBytes"`
	AvailableSpaceInBytes      pulumi.Float64PtrInput               `pulumi:"availableSpaceInBytes"`
	CpuLoad                    pulumi.StringPtrInput                `pulumi:"cpuLoad"`
	CpuLoadStatus              pulumi.StringPtrInput                `pulumi:"cpuLoadStatus"`
	CsServiceStatus            pulumi.StringPtrInput                `pulumi:"csServiceStatus"`
	DatabaseServerLoad         pulumi.StringPtrInput                `pulumi:"databaseServerLoad"`
	DatabaseServerLoadStatus   pulumi.StringPtrInput                `pulumi:"databaseServerLoadStatus"`
	HostName                   pulumi.StringPtrInput                `pulumi:"hostName"`
	InstanceType               pulumi.StringInput                   `pulumi:"instanceType"`
	IpAddress                  pulumi.StringPtrInput                `pulumi:"ipAddress"`
	LastHeartbeat              pulumi.StringPtrInput                `pulumi:"lastHeartbeat"`
	MasterTargetServers        MasterTargetServerResponseArrayInput `pulumi:"masterTargetServers"`
	MemoryUsageStatus          pulumi.StringPtrInput                `pulumi:"memoryUsageStatus"`
	ProcessServerCount         pulumi.StringPtrInput                `pulumi:"processServerCount"`
	ProcessServers             ProcessServerResponseArrayInput      `pulumi:"processServers"`
	ProtectedServers           pulumi.StringPtrInput                `pulumi:"protectedServers"`
	PsTemplateVersion          pulumi.StringPtrInput                `pulumi:"psTemplateVersion"`
	ReplicationPairCount       pulumi.StringPtrInput                `pulumi:"replicationPairCount"`
	RunAsAccounts              RunAsAccountResponseArrayInput       `pulumi:"runAsAccounts"`
	SpaceUsageStatus           pulumi.StringPtrInput                `pulumi:"spaceUsageStatus"`
	SslCertExpiryDate          pulumi.StringPtrInput                `pulumi:"sslCertExpiryDate"`
	SslCertExpiryRemainingDays pulumi.IntPtrInput                   `pulumi:"sslCertExpiryRemainingDays"`
	SystemLoad                 pulumi.StringPtrInput                `pulumi:"systemLoad"`
	SystemLoadStatus           pulumi.StringPtrInput                `pulumi:"systemLoadStatus"`
	TotalMemoryInBytes         pulumi.Float64PtrInput               `pulumi:"totalMemoryInBytes"`
	TotalSpaceInBytes          pulumi.Float64PtrInput               `pulumi:"totalSpaceInBytes"`
	VersionStatus              pulumi.StringPtrInput                `pulumi:"versionStatus"`
	WebLoad                    pulumi.StringPtrInput                `pulumi:"webLoad"`
	WebLoadStatus              pulumi.StringPtrInput                `pulumi:"webLoadStatus"`
}

func (VMwareDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareDetailsResponse)(nil)).Elem()
}

func (i VMwareDetailsResponseArgs) ToVMwareDetailsResponseOutput() VMwareDetailsResponseOutput {
	return i.ToVMwareDetailsResponseOutputWithContext(context.Background())
}

func (i VMwareDetailsResponseArgs) ToVMwareDetailsResponseOutputWithContext(ctx context.Context) VMwareDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareDetailsResponseOutput)
}

type VMwareDetailsResponseOutput struct{ *pulumi.OutputState }

func (VMwareDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareDetailsResponse)(nil)).Elem()
}

func (o VMwareDetailsResponseOutput) ToVMwareDetailsResponseOutput() VMwareDetailsResponseOutput {
	return o
}

func (o VMwareDetailsResponseOutput) ToVMwareDetailsResponseOutputWithContext(ctx context.Context) VMwareDetailsResponseOutput {
	return o
}

func (o VMwareDetailsResponseOutput) AgentCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.AgentCount }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) AgentExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.AgentExpiryDate }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) AgentVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.AgentVersion }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) AgentVersionDetails() VersionDetailsResponsePtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *VersionDetailsResponse { return v.AgentVersionDetails }).(VersionDetailsResponsePtrOutput)
}

func (o VMwareDetailsResponseOutput) AvailableMemoryInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *float64 { return v.AvailableMemoryInBytes }).(pulumi.Float64PtrOutput)
}

func (o VMwareDetailsResponseOutput) AvailableSpaceInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *float64 { return v.AvailableSpaceInBytes }).(pulumi.Float64PtrOutput)
}

func (o VMwareDetailsResponseOutput) CpuLoad() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.CpuLoad }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) CpuLoadStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.CpuLoadStatus }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) CsServiceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.CsServiceStatus }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) DatabaseServerLoad() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.DatabaseServerLoad }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) DatabaseServerLoadStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.DatabaseServerLoadStatus }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.HostName }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o VMwareDetailsResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) LastHeartbeat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.LastHeartbeat }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) MasterTargetServers() MasterTargetServerResponseArrayOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) []MasterTargetServerResponse { return v.MasterTargetServers }).(MasterTargetServerResponseArrayOutput)
}

func (o VMwareDetailsResponseOutput) MemoryUsageStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.MemoryUsageStatus }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) ProcessServerCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.ProcessServerCount }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) ProcessServers() ProcessServerResponseArrayOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) []ProcessServerResponse { return v.ProcessServers }).(ProcessServerResponseArrayOutput)
}

func (o VMwareDetailsResponseOutput) ProtectedServers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.ProtectedServers }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) PsTemplateVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.PsTemplateVersion }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) ReplicationPairCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.ReplicationPairCount }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) RunAsAccounts() RunAsAccountResponseArrayOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) []RunAsAccountResponse { return v.RunAsAccounts }).(RunAsAccountResponseArrayOutput)
}

func (o VMwareDetailsResponseOutput) SpaceUsageStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.SpaceUsageStatus }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) SslCertExpiryDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.SslCertExpiryDate }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) SslCertExpiryRemainingDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *int { return v.SslCertExpiryRemainingDays }).(pulumi.IntPtrOutput)
}

func (o VMwareDetailsResponseOutput) SystemLoad() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.SystemLoad }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) SystemLoadStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.SystemLoadStatus }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) TotalMemoryInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *float64 { return v.TotalMemoryInBytes }).(pulumi.Float64PtrOutput)
}

func (o VMwareDetailsResponseOutput) TotalSpaceInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *float64 { return v.TotalSpaceInBytes }).(pulumi.Float64PtrOutput)
}

func (o VMwareDetailsResponseOutput) VersionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.VersionStatus }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) WebLoad() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.WebLoad }).(pulumi.StringPtrOutput)
}

func (o VMwareDetailsResponseOutput) WebLoadStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareDetailsResponse) *string { return v.WebLoadStatus }).(pulumi.StringPtrOutput)
}

type VMwareV2FabricCreationInput struct {
	InstanceType        *string `pulumi:"instanceType"`
	MigrationSolutionId string  `pulumi:"migrationSolutionId"`
	PhysicalSiteId      *string `pulumi:"physicalSiteId"`
	VmwareSiteId        *string `pulumi:"vmwareSiteId"`
}





type VMwareV2FabricCreationInputInput interface {
	pulumi.Input

	ToVMwareV2FabricCreationInputOutput() VMwareV2FabricCreationInputOutput
	ToVMwareV2FabricCreationInputOutputWithContext(context.Context) VMwareV2FabricCreationInputOutput
}

type VMwareV2FabricCreationInputArgs struct {
	InstanceType        pulumi.StringPtrInput `pulumi:"instanceType"`
	MigrationSolutionId pulumi.StringInput    `pulumi:"migrationSolutionId"`
	PhysicalSiteId      pulumi.StringPtrInput `pulumi:"physicalSiteId"`
	VmwareSiteId        pulumi.StringPtrInput `pulumi:"vmwareSiteId"`
}

func (VMwareV2FabricCreationInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareV2FabricCreationInput)(nil)).Elem()
}

func (i VMwareV2FabricCreationInputArgs) ToVMwareV2FabricCreationInputOutput() VMwareV2FabricCreationInputOutput {
	return i.ToVMwareV2FabricCreationInputOutputWithContext(context.Background())
}

func (i VMwareV2FabricCreationInputArgs) ToVMwareV2FabricCreationInputOutputWithContext(ctx context.Context) VMwareV2FabricCreationInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareV2FabricCreationInputOutput)
}

type VMwareV2FabricCreationInputOutput struct{ *pulumi.OutputState }

func (VMwareV2FabricCreationInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareV2FabricCreationInput)(nil)).Elem()
}

func (o VMwareV2FabricCreationInputOutput) ToVMwareV2FabricCreationInputOutput() VMwareV2FabricCreationInputOutput {
	return o
}

func (o VMwareV2FabricCreationInputOutput) ToVMwareV2FabricCreationInputOutputWithContext(ctx context.Context) VMwareV2FabricCreationInputOutput {
	return o
}

func (o VMwareV2FabricCreationInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareV2FabricCreationInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

func (o VMwareV2FabricCreationInputOutput) MigrationSolutionId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareV2FabricCreationInput) string { return v.MigrationSolutionId }).(pulumi.StringOutput)
}

func (o VMwareV2FabricCreationInputOutput) PhysicalSiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareV2FabricCreationInput) *string { return v.PhysicalSiteId }).(pulumi.StringPtrOutput)
}

func (o VMwareV2FabricCreationInputOutput) VmwareSiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMwareV2FabricCreationInput) *string { return v.VmwareSiteId }).(pulumi.StringPtrOutput)
}

type VMwareV2FabricSpecificDetailsResponse struct {
	InstanceType        string `pulumi:"instanceType"`
	MigrationSolutionId string `pulumi:"migrationSolutionId"`
	PhysicalSiteId      string `pulumi:"physicalSiteId"`
	ServiceEndpoint     string `pulumi:"serviceEndpoint"`
	ServiceResourceId   string `pulumi:"serviceResourceId"`
	VmwareSiteId        string `pulumi:"vmwareSiteId"`
}





type VMwareV2FabricSpecificDetailsResponseInput interface {
	pulumi.Input

	ToVMwareV2FabricSpecificDetailsResponseOutput() VMwareV2FabricSpecificDetailsResponseOutput
	ToVMwareV2FabricSpecificDetailsResponseOutputWithContext(context.Context) VMwareV2FabricSpecificDetailsResponseOutput
}

type VMwareV2FabricSpecificDetailsResponseArgs struct {
	InstanceType        pulumi.StringInput `pulumi:"instanceType"`
	MigrationSolutionId pulumi.StringInput `pulumi:"migrationSolutionId"`
	PhysicalSiteId      pulumi.StringInput `pulumi:"physicalSiteId"`
	ServiceEndpoint     pulumi.StringInput `pulumi:"serviceEndpoint"`
	ServiceResourceId   pulumi.StringInput `pulumi:"serviceResourceId"`
	VmwareSiteId        pulumi.StringInput `pulumi:"vmwareSiteId"`
}

func (VMwareV2FabricSpecificDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareV2FabricSpecificDetailsResponse)(nil)).Elem()
}

func (i VMwareV2FabricSpecificDetailsResponseArgs) ToVMwareV2FabricSpecificDetailsResponseOutput() VMwareV2FabricSpecificDetailsResponseOutput {
	return i.ToVMwareV2FabricSpecificDetailsResponseOutputWithContext(context.Background())
}

func (i VMwareV2FabricSpecificDetailsResponseArgs) ToVMwareV2FabricSpecificDetailsResponseOutputWithContext(ctx context.Context) VMwareV2FabricSpecificDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareV2FabricSpecificDetailsResponseOutput)
}

type VMwareV2FabricSpecificDetailsResponseOutput struct{ *pulumi.OutputState }

func (VMwareV2FabricSpecificDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMwareV2FabricSpecificDetailsResponse)(nil)).Elem()
}

func (o VMwareV2FabricSpecificDetailsResponseOutput) ToVMwareV2FabricSpecificDetailsResponseOutput() VMwareV2FabricSpecificDetailsResponseOutput {
	return o
}

func (o VMwareV2FabricSpecificDetailsResponseOutput) ToVMwareV2FabricSpecificDetailsResponseOutputWithContext(ctx context.Context) VMwareV2FabricSpecificDetailsResponseOutput {
	return o
}

func (o VMwareV2FabricSpecificDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareV2FabricSpecificDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o VMwareV2FabricSpecificDetailsResponseOutput) MigrationSolutionId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareV2FabricSpecificDetailsResponse) string { return v.MigrationSolutionId }).(pulumi.StringOutput)
}

func (o VMwareV2FabricSpecificDetailsResponseOutput) PhysicalSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareV2FabricSpecificDetailsResponse) string { return v.PhysicalSiteId }).(pulumi.StringOutput)
}

func (o VMwareV2FabricSpecificDetailsResponseOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareV2FabricSpecificDetailsResponse) string { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

func (o VMwareV2FabricSpecificDetailsResponseOutput) ServiceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareV2FabricSpecificDetailsResponse) string { return v.ServiceResourceId }).(pulumi.StringOutput)
}

func (o VMwareV2FabricSpecificDetailsResponseOutput) VmwareSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v VMwareV2FabricSpecificDetailsResponse) string { return v.VmwareSiteId }).(pulumi.StringOutput)
}

type VersionDetailsResponse struct {
	ExpiryDate *string `pulumi:"expiryDate"`
	Status     *string `pulumi:"status"`
	Version    *string `pulumi:"version"`
}





type VersionDetailsResponseInput interface {
	pulumi.Input

	ToVersionDetailsResponseOutput() VersionDetailsResponseOutput
	ToVersionDetailsResponseOutputWithContext(context.Context) VersionDetailsResponseOutput
}

type VersionDetailsResponseArgs struct {
	ExpiryDate pulumi.StringPtrInput `pulumi:"expiryDate"`
	Status     pulumi.StringPtrInput `pulumi:"status"`
	Version    pulumi.StringPtrInput `pulumi:"version"`
}

func (VersionDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VersionDetailsResponse)(nil)).Elem()
}

func (i VersionDetailsResponseArgs) ToVersionDetailsResponseOutput() VersionDetailsResponseOutput {
	return i.ToVersionDetailsResponseOutputWithContext(context.Background())
}

func (i VersionDetailsResponseArgs) ToVersionDetailsResponseOutputWithContext(ctx context.Context) VersionDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VersionDetailsResponseOutput)
}

func (i VersionDetailsResponseArgs) ToVersionDetailsResponsePtrOutput() VersionDetailsResponsePtrOutput {
	return i.ToVersionDetailsResponsePtrOutputWithContext(context.Background())
}

func (i VersionDetailsResponseArgs) ToVersionDetailsResponsePtrOutputWithContext(ctx context.Context) VersionDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VersionDetailsResponseOutput).ToVersionDetailsResponsePtrOutputWithContext(ctx)
}









type VersionDetailsResponsePtrInput interface {
	pulumi.Input

	ToVersionDetailsResponsePtrOutput() VersionDetailsResponsePtrOutput
	ToVersionDetailsResponsePtrOutputWithContext(context.Context) VersionDetailsResponsePtrOutput
}

type versionDetailsResponsePtrType VersionDetailsResponseArgs

func VersionDetailsResponsePtr(v *VersionDetailsResponseArgs) VersionDetailsResponsePtrInput {
	return (*versionDetailsResponsePtrType)(v)
}

func (*versionDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VersionDetailsResponse)(nil)).Elem()
}

func (i *versionDetailsResponsePtrType) ToVersionDetailsResponsePtrOutput() VersionDetailsResponsePtrOutput {
	return i.ToVersionDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *versionDetailsResponsePtrType) ToVersionDetailsResponsePtrOutputWithContext(ctx context.Context) VersionDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VersionDetailsResponsePtrOutput)
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

func (o VersionDetailsResponseOutput) ToVersionDetailsResponsePtrOutput() VersionDetailsResponsePtrOutput {
	return o.ToVersionDetailsResponsePtrOutputWithContext(context.Background())
}

func (o VersionDetailsResponseOutput) ToVersionDetailsResponsePtrOutputWithContext(ctx context.Context) VersionDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VersionDetailsResponse) *VersionDetailsResponse {
		return &v
	}).(VersionDetailsResponsePtrOutput)
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





type VmmDetailsResponseInput interface {
	pulumi.Input

	ToVmmDetailsResponseOutput() VmmDetailsResponseOutput
	ToVmmDetailsResponseOutputWithContext(context.Context) VmmDetailsResponseOutput
}

type VmmDetailsResponseArgs struct {
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
}

func (VmmDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmmDetailsResponse)(nil)).Elem()
}

func (i VmmDetailsResponseArgs) ToVmmDetailsResponseOutput() VmmDetailsResponseOutput {
	return i.ToVmmDetailsResponseOutputWithContext(context.Background())
}

func (i VmmDetailsResponseArgs) ToVmmDetailsResponseOutputWithContext(ctx context.Context) VmmDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmmDetailsResponseOutput)
}

type VmmDetailsResponseOutput struct{ *pulumi.OutputState }

func (VmmDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmmDetailsResponse)(nil)).Elem()
}

func (o VmmDetailsResponseOutput) ToVmmDetailsResponseOutput() VmmDetailsResponseOutput {
	return o
}

func (o VmmDetailsResponseOutput) ToVmmDetailsResponseOutputWithContext(ctx context.Context) VmmDetailsResponseOutput {
	return o
}

func (o VmmDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v VmmDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

type VmmToAzureCreateNetworkMappingInput struct {
	InstanceType *string `pulumi:"instanceType"`
}





type VmmToAzureCreateNetworkMappingInputInput interface {
	pulumi.Input

	ToVmmToAzureCreateNetworkMappingInputOutput() VmmToAzureCreateNetworkMappingInputOutput
	ToVmmToAzureCreateNetworkMappingInputOutputWithContext(context.Context) VmmToAzureCreateNetworkMappingInputOutput
}

type VmmToAzureCreateNetworkMappingInputArgs struct {
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
}

func (VmmToAzureCreateNetworkMappingInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmmToAzureCreateNetworkMappingInput)(nil)).Elem()
}

func (i VmmToAzureCreateNetworkMappingInputArgs) ToVmmToAzureCreateNetworkMappingInputOutput() VmmToAzureCreateNetworkMappingInputOutput {
	return i.ToVmmToAzureCreateNetworkMappingInputOutputWithContext(context.Background())
}

func (i VmmToAzureCreateNetworkMappingInputArgs) ToVmmToAzureCreateNetworkMappingInputOutputWithContext(ctx context.Context) VmmToAzureCreateNetworkMappingInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmmToAzureCreateNetworkMappingInputOutput)
}

type VmmToAzureCreateNetworkMappingInputOutput struct{ *pulumi.OutputState }

func (VmmToAzureCreateNetworkMappingInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmmToAzureCreateNetworkMappingInput)(nil)).Elem()
}

func (o VmmToAzureCreateNetworkMappingInputOutput) ToVmmToAzureCreateNetworkMappingInputOutput() VmmToAzureCreateNetworkMappingInputOutput {
	return o
}

func (o VmmToAzureCreateNetworkMappingInputOutput) ToVmmToAzureCreateNetworkMappingInputOutputWithContext(ctx context.Context) VmmToAzureCreateNetworkMappingInputOutput {
	return o
}

func (o VmmToAzureCreateNetworkMappingInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmmToAzureCreateNetworkMappingInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

type VmmToAzureNetworkMappingSettingsResponse struct {
	InstanceType string `pulumi:"instanceType"`
}





type VmmToAzureNetworkMappingSettingsResponseInput interface {
	pulumi.Input

	ToVmmToAzureNetworkMappingSettingsResponseOutput() VmmToAzureNetworkMappingSettingsResponseOutput
	ToVmmToAzureNetworkMappingSettingsResponseOutputWithContext(context.Context) VmmToAzureNetworkMappingSettingsResponseOutput
}

type VmmToAzureNetworkMappingSettingsResponseArgs struct {
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
}

func (VmmToAzureNetworkMappingSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmmToAzureNetworkMappingSettingsResponse)(nil)).Elem()
}

func (i VmmToAzureNetworkMappingSettingsResponseArgs) ToVmmToAzureNetworkMappingSettingsResponseOutput() VmmToAzureNetworkMappingSettingsResponseOutput {
	return i.ToVmmToAzureNetworkMappingSettingsResponseOutputWithContext(context.Background())
}

func (i VmmToAzureNetworkMappingSettingsResponseArgs) ToVmmToAzureNetworkMappingSettingsResponseOutputWithContext(ctx context.Context) VmmToAzureNetworkMappingSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmmToAzureNetworkMappingSettingsResponseOutput)
}

type VmmToAzureNetworkMappingSettingsResponseOutput struct{ *pulumi.OutputState }

func (VmmToAzureNetworkMappingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmmToAzureNetworkMappingSettingsResponse)(nil)).Elem()
}

func (o VmmToAzureNetworkMappingSettingsResponseOutput) ToVmmToAzureNetworkMappingSettingsResponseOutput() VmmToAzureNetworkMappingSettingsResponseOutput {
	return o
}

func (o VmmToAzureNetworkMappingSettingsResponseOutput) ToVmmToAzureNetworkMappingSettingsResponseOutputWithContext(ctx context.Context) VmmToAzureNetworkMappingSettingsResponseOutput {
	return o
}

func (o VmmToAzureNetworkMappingSettingsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v VmmToAzureNetworkMappingSettingsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

type VmmToVmmCreateNetworkMappingInput struct {
	InstanceType *string `pulumi:"instanceType"`
}





type VmmToVmmCreateNetworkMappingInputInput interface {
	pulumi.Input

	ToVmmToVmmCreateNetworkMappingInputOutput() VmmToVmmCreateNetworkMappingInputOutput
	ToVmmToVmmCreateNetworkMappingInputOutputWithContext(context.Context) VmmToVmmCreateNetworkMappingInputOutput
}

type VmmToVmmCreateNetworkMappingInputArgs struct {
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
}

func (VmmToVmmCreateNetworkMappingInputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmmToVmmCreateNetworkMappingInput)(nil)).Elem()
}

func (i VmmToVmmCreateNetworkMappingInputArgs) ToVmmToVmmCreateNetworkMappingInputOutput() VmmToVmmCreateNetworkMappingInputOutput {
	return i.ToVmmToVmmCreateNetworkMappingInputOutputWithContext(context.Background())
}

func (i VmmToVmmCreateNetworkMappingInputArgs) ToVmmToVmmCreateNetworkMappingInputOutputWithContext(ctx context.Context) VmmToVmmCreateNetworkMappingInputOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmmToVmmCreateNetworkMappingInputOutput)
}

type VmmToVmmCreateNetworkMappingInputOutput struct{ *pulumi.OutputState }

func (VmmToVmmCreateNetworkMappingInputOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmmToVmmCreateNetworkMappingInput)(nil)).Elem()
}

func (o VmmToVmmCreateNetworkMappingInputOutput) ToVmmToVmmCreateNetworkMappingInputOutput() VmmToVmmCreateNetworkMappingInputOutput {
	return o
}

func (o VmmToVmmCreateNetworkMappingInputOutput) ToVmmToVmmCreateNetworkMappingInputOutputWithContext(ctx context.Context) VmmToVmmCreateNetworkMappingInputOutput {
	return o
}

func (o VmmToVmmCreateNetworkMappingInputOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VmmToVmmCreateNetworkMappingInput) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

type VmmToVmmNetworkMappingSettingsResponse struct {
	InstanceType string `pulumi:"instanceType"`
}





type VmmToVmmNetworkMappingSettingsResponseInput interface {
	pulumi.Input

	ToVmmToVmmNetworkMappingSettingsResponseOutput() VmmToVmmNetworkMappingSettingsResponseOutput
	ToVmmToVmmNetworkMappingSettingsResponseOutputWithContext(context.Context) VmmToVmmNetworkMappingSettingsResponseOutput
}

type VmmToVmmNetworkMappingSettingsResponseArgs struct {
	InstanceType pulumi.StringInput `pulumi:"instanceType"`
}

func (VmmToVmmNetworkMappingSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmmToVmmNetworkMappingSettingsResponse)(nil)).Elem()
}

func (i VmmToVmmNetworkMappingSettingsResponseArgs) ToVmmToVmmNetworkMappingSettingsResponseOutput() VmmToVmmNetworkMappingSettingsResponseOutput {
	return i.ToVmmToVmmNetworkMappingSettingsResponseOutputWithContext(context.Background())
}

func (i VmmToVmmNetworkMappingSettingsResponseArgs) ToVmmToVmmNetworkMappingSettingsResponseOutputWithContext(ctx context.Context) VmmToVmmNetworkMappingSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmmToVmmNetworkMappingSettingsResponseOutput)
}

type VmmToVmmNetworkMappingSettingsResponseOutput struct{ *pulumi.OutputState }

func (VmmToVmmNetworkMappingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmmToVmmNetworkMappingSettingsResponse)(nil)).Elem()
}

func (o VmmToVmmNetworkMappingSettingsResponseOutput) ToVmmToVmmNetworkMappingSettingsResponseOutput() VmmToVmmNetworkMappingSettingsResponseOutput {
	return o
}

func (o VmmToVmmNetworkMappingSettingsResponseOutput) ToVmmToVmmNetworkMappingSettingsResponseOutputWithContext(ctx context.Context) VmmToVmmNetworkMappingSettingsResponseOutput {
	return o
}

func (o VmmToVmmNetworkMappingSettingsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v VmmToVmmNetworkMappingSettingsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

type VmwareCbtPolicyDetailsResponse struct {
	AppConsistentFrequencyInMinutes   *int   `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes *int   `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      string `pulumi:"instanceType"`
	RecoveryPointHistoryInMinutes     *int   `pulumi:"recoveryPointHistoryInMinutes"`
}





type VmwareCbtPolicyDetailsResponseInput interface {
	pulumi.Input

	ToVmwareCbtPolicyDetailsResponseOutput() VmwareCbtPolicyDetailsResponseOutput
	ToVmwareCbtPolicyDetailsResponseOutputWithContext(context.Context) VmwareCbtPolicyDetailsResponseOutput
}

type VmwareCbtPolicyDetailsResponseArgs struct {
	AppConsistentFrequencyInMinutes   pulumi.IntPtrInput `pulumi:"appConsistentFrequencyInMinutes"`
	CrashConsistentFrequencyInMinutes pulumi.IntPtrInput `pulumi:"crashConsistentFrequencyInMinutes"`
	InstanceType                      pulumi.StringInput `pulumi:"instanceType"`
	RecoveryPointHistoryInMinutes     pulumi.IntPtrInput `pulumi:"recoveryPointHistoryInMinutes"`
}

func (VmwareCbtPolicyDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmwareCbtPolicyDetailsResponse)(nil)).Elem()
}

func (i VmwareCbtPolicyDetailsResponseArgs) ToVmwareCbtPolicyDetailsResponseOutput() VmwareCbtPolicyDetailsResponseOutput {
	return i.ToVmwareCbtPolicyDetailsResponseOutputWithContext(context.Background())
}

func (i VmwareCbtPolicyDetailsResponseArgs) ToVmwareCbtPolicyDetailsResponseOutputWithContext(ctx context.Context) VmwareCbtPolicyDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmwareCbtPolicyDetailsResponseOutput)
}

type VmwareCbtPolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (VmwareCbtPolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmwareCbtPolicyDetailsResponse)(nil)).Elem()
}

func (o VmwareCbtPolicyDetailsResponseOutput) ToVmwareCbtPolicyDetailsResponseOutput() VmwareCbtPolicyDetailsResponseOutput {
	return o
}

func (o VmwareCbtPolicyDetailsResponseOutput) ToVmwareCbtPolicyDetailsResponseOutputWithContext(ctx context.Context) VmwareCbtPolicyDetailsResponseOutput {
	return o
}

func (o VmwareCbtPolicyDetailsResponseOutput) AppConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VmwareCbtPolicyDetailsResponse) *int { return v.AppConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o VmwareCbtPolicyDetailsResponseOutput) CrashConsistentFrequencyInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VmwareCbtPolicyDetailsResponse) *int { return v.CrashConsistentFrequencyInMinutes }).(pulumi.IntPtrOutput)
}

func (o VmwareCbtPolicyDetailsResponseOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v VmwareCbtPolicyDetailsResponse) string { return v.InstanceType }).(pulumi.StringOutput)
}

func (o VmwareCbtPolicyDetailsResponseOutput) RecoveryPointHistoryInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VmwareCbtPolicyDetailsResponse) *int { return v.RecoveryPointHistoryInMinutes }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*A2AContainerMappingInputInput)(nil)).Elem(), A2AContainerMappingInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AEnableProtectionInputInput)(nil)).Elem(), A2AEnableProtectionInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2APolicyCreationInputInput)(nil)).Elem(), A2APolicyCreationInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2APolicyDetailsResponseInput)(nil)).Elem(), A2APolicyDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AProtectedDiskDetailsResponseInput)(nil)).Elem(), A2AProtectedDiskDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AProtectedDiskDetailsResponseArrayInput)(nil)).Elem(), A2AProtectedDiskDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AProtectedManagedDiskDetailsResponseInput)(nil)).Elem(), A2AProtectedManagedDiskDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AProtectedManagedDiskDetailsResponseArrayInput)(nil)).Elem(), A2AProtectedManagedDiskDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AProtectionContainerMappingDetailsResponseInput)(nil)).Elem(), A2AProtectionContainerMappingDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AReplicationDetailsResponseInput)(nil)).Elem(), A2AReplicationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AVmDiskInputDetailsInput)(nil)).Elem(), A2AVmDiskInputDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AVmDiskInputDetailsArrayInput)(nil)).Elem(), A2AVmDiskInputDetailsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AVmManagedDiskInputDetailsInput)(nil)).Elem(), A2AVmManagedDiskInputDetailsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*A2AVmManagedDiskInputDetailsArrayInput)(nil)).Elem(), A2AVmManagedDiskInputDetailsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddRecoveryServicesProviderInputPropertiesInput)(nil)).Elem(), AddRecoveryServicesProviderInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddRecoveryServicesProviderInputPropertiesPtrInput)(nil)).Elem(), AddRecoveryServicesProviderInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddVCenterRequestPropertiesInput)(nil)).Elem(), AddVCenterRequestPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddVCenterRequestPropertiesPtrInput)(nil)).Elem(), AddVCenterRequestPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFabricCreationInputInput)(nil)).Elem(), AzureFabricCreationInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureFabricSpecificDetailsResponseInput)(nil)).Elem(), AzureFabricSpecificDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureToAzureCreateNetworkMappingInputInput)(nil)).Elem(), AzureToAzureCreateNetworkMappingInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureToAzureNetworkMappingSettingsResponseInput)(nil)).Elem(), AzureToAzureNetworkMappingSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureToAzureVmSyncedConfigDetailsResponseInput)(nil)).Elem(), AzureToAzureVmSyncedConfigDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureToAzureVmSyncedConfigDetailsResponsePtrInput)(nil)).Elem(), AzureToAzureVmSyncedConfigDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureVmDiskDetailsResponseInput)(nil)).Elem(), AzureVmDiskDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureVmDiskDetailsResponseArrayInput)(nil)).Elem(), AzureVmDiskDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateNetworkMappingInputPropertiesInput)(nil)).Elem(), CreateNetworkMappingInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateNetworkMappingInputPropertiesPtrInput)(nil)).Elem(), CreateNetworkMappingInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreatePolicyInputPropertiesInput)(nil)).Elem(), CreatePolicyInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreatePolicyInputPropertiesPtrInput)(nil)).Elem(), CreatePolicyInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateProtectionContainerMappingInputPropertiesInput)(nil)).Elem(), CreateProtectionContainerMappingInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateProtectionContainerMappingInputPropertiesPtrInput)(nil)).Elem(), CreateProtectionContainerMappingInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateRecoveryPlanInputPropertiesInput)(nil)).Elem(), CreateRecoveryPlanInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CreateRecoveryPlanInputPropertiesPtrInput)(nil)).Elem(), CreateRecoveryPlanInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CurrentJobDetailsResponseInput)(nil)).Elem(), CurrentJobDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CurrentJobDetailsResponsePtrInput)(nil)).Elem(), CurrentJobDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CurrentScenarioDetailsResponseInput)(nil)).Elem(), CurrentScenarioDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CurrentScenarioDetailsResponsePtrInput)(nil)).Elem(), CurrentScenarioDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataStoreResponseInput)(nil)).Elem(), DataStoreResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataStoreResponseArrayInput)(nil)).Elem(), DataStoreResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskDetailsResponseInput)(nil)).Elem(), DiskDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskDetailsResponseArrayInput)(nil)).Elem(), DiskDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskEncryptionInfoInput)(nil)).Elem(), DiskEncryptionInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskEncryptionInfoPtrInput)(nil)).Elem(), DiskEncryptionInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskEncryptionKeyInfoInput)(nil)).Elem(), DiskEncryptionKeyInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiskEncryptionKeyInfoPtrInput)(nil)).Elem(), DiskEncryptionKeyInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnableMigrationInputPropertiesInput)(nil)).Elem(), EnableMigrationInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnableMigrationInputPropertiesPtrInput)(nil)).Elem(), EnableMigrationInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnableProtectionInputPropertiesInput)(nil)).Elem(), EnableProtectionInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnableProtectionInputPropertiesPtrInput)(nil)).Elem(), EnableProtectionInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionDetailsResponseInput)(nil)).Elem(), EncryptionDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EncryptionDetailsResponsePtrInput)(nil)).Elem(), EncryptionDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FabricCreationInputPropertiesInput)(nil)).Elem(), FabricCreationInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FabricCreationInputPropertiesPtrInput)(nil)).Elem(), FabricCreationInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FabricPropertiesResponseInput)(nil)).Elem(), FabricPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FabricPropertiesResponsePtrInput)(nil)).Elem(), FabricPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthErrorResponseInput)(nil)).Elem(), HealthErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthErrorResponseArrayInput)(nil)).Elem(), HealthErrorResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaAzureEnableProtectionInputInput)(nil)).Elem(), HyperVReplicaAzureEnableProtectionInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaAzurePolicyDetailsResponseInput)(nil)).Elem(), HyperVReplicaAzurePolicyDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaAzurePolicyInputInput)(nil)).Elem(), HyperVReplicaAzurePolicyInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaAzureReplicationDetailsResponseInput)(nil)).Elem(), HyperVReplicaAzureReplicationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaBasePolicyDetailsResponseInput)(nil)).Elem(), HyperVReplicaBasePolicyDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaBaseReplicationDetailsResponseInput)(nil)).Elem(), HyperVReplicaBaseReplicationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaBluePolicyDetailsResponseInput)(nil)).Elem(), HyperVReplicaBluePolicyDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaBluePolicyInputInput)(nil)).Elem(), HyperVReplicaBluePolicyInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaBlueReplicationDetailsResponseInput)(nil)).Elem(), HyperVReplicaBlueReplicationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaPolicyDetailsResponseInput)(nil)).Elem(), HyperVReplicaPolicyDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaPolicyInputInput)(nil)).Elem(), HyperVReplicaPolicyInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVReplicaReplicationDetailsResponseInput)(nil)).Elem(), HyperVReplicaReplicationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HyperVSiteDetailsResponseInput)(nil)).Elem(), HyperVSiteDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderDetailsResponseInput)(nil)).Elem(), IdentityProviderDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderDetailsResponsePtrInput)(nil)).Elem(), IdentityProviderDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderInputInput)(nil)).Elem(), IdentityProviderInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityProviderInputPtrInput)(nil)).Elem(), IdentityProviderInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageAgentDetailsResponseInput)(nil)).Elem(), InMageAgentDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageAgentDetailsResponsePtrInput)(nil)).Elem(), InMageAgentDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageAzureV2EnableProtectionInputInput)(nil)).Elem(), InMageAzureV2EnableProtectionInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageAzureV2PolicyDetailsResponseInput)(nil)).Elem(), InMageAzureV2PolicyDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageAzureV2PolicyInputInput)(nil)).Elem(), InMageAzureV2PolicyInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageAzureV2ProtectedDiskDetailsResponseInput)(nil)).Elem(), InMageAzureV2ProtectedDiskDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageAzureV2ProtectedDiskDetailsResponseArrayInput)(nil)).Elem(), InMageAzureV2ProtectedDiskDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageAzureV2ReplicationDetailsResponseInput)(nil)).Elem(), InMageAzureV2ReplicationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageBasePolicyDetailsResponseInput)(nil)).Elem(), InMageBasePolicyDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageDiskExclusionInputInput)(nil)).Elem(), InMageDiskExclusionInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageDiskExclusionInputPtrInput)(nil)).Elem(), InMageDiskExclusionInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageDiskSignatureExclusionOptionsInput)(nil)).Elem(), InMageDiskSignatureExclusionOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageDiskSignatureExclusionOptionsArrayInput)(nil)).Elem(), InMageDiskSignatureExclusionOptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageEnableProtectionInputInput)(nil)).Elem(), InMageEnableProtectionInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMagePolicyDetailsResponseInput)(nil)).Elem(), InMagePolicyDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMagePolicyInputInput)(nil)).Elem(), InMagePolicyInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageProtectedDiskDetailsResponseInput)(nil)).Elem(), InMageProtectedDiskDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageProtectedDiskDetailsResponseArrayInput)(nil)).Elem(), InMageProtectedDiskDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageReplicationDetailsResponseInput)(nil)).Elem(), InMageReplicationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageVolumeExclusionOptionsInput)(nil)).Elem(), InMageVolumeExclusionOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InMageVolumeExclusionOptionsArrayInput)(nil)).Elem(), InMageVolumeExclusionOptionsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InitialReplicationDetailsResponseInput)(nil)).Elem(), InitialReplicationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InitialReplicationDetailsResponsePtrInput)(nil)).Elem(), InitialReplicationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InnerHealthErrorResponseInput)(nil)).Elem(), InnerHealthErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InnerHealthErrorResponseArrayInput)(nil)).Elem(), InnerHealthErrorResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputEndpointResponseInput)(nil)).Elem(), InputEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InputEndpointResponseArrayInput)(nil)).Elem(), InputEndpointResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyEncryptionKeyInfoInput)(nil)).Elem(), KeyEncryptionKeyInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyEncryptionKeyInfoPtrInput)(nil)).Elem(), KeyEncryptionKeyInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MasterTargetServerResponseInput)(nil)).Elem(), MasterTargetServerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MasterTargetServerResponseArrayInput)(nil)).Elem(), MasterTargetServerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrationItemPropertiesResponseInput)(nil)).Elem(), MigrationItemPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MigrationItemPropertiesResponsePtrInput)(nil)).Elem(), MigrationItemPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MobilityServiceUpdateResponseInput)(nil)).Elem(), MobilityServiceUpdateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MobilityServiceUpdateResponseArrayInput)(nil)).Elem(), MobilityServiceUpdateResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkMappingPropertiesResponseInput)(nil)).Elem(), NetworkMappingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkMappingPropertiesResponsePtrInput)(nil)).Elem(), NetworkMappingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OSDetailsResponseInput)(nil)).Elem(), OSDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OSDetailsResponsePtrInput)(nil)).Elem(), OSDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OSDiskDetailsResponseInput)(nil)).Elem(), OSDiskDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OSDiskDetailsResponsePtrInput)(nil)).Elem(), OSDiskDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyPropertiesResponseInput)(nil)).Elem(), PolicyPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyPropertiesResponsePtrInput)(nil)).Elem(), PolicyPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProcessServerResponseInput)(nil)).Elem(), ProcessServerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProcessServerResponseArrayInput)(nil)).Elem(), ProcessServerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProtectionContainerMappingPropertiesResponseInput)(nil)).Elem(), ProtectionContainerMappingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProtectionContainerMappingPropertiesResponsePtrInput)(nil)).Elem(), ProtectionContainerMappingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RcmAzureMigrationPolicyDetailsResponseInput)(nil)).Elem(), RcmAzureMigrationPolicyDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanActionInput)(nil)).Elem(), RecoveryPlanActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanActionArrayInput)(nil)).Elem(), RecoveryPlanActionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanActionResponseInput)(nil)).Elem(), RecoveryPlanActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanActionResponseArrayInput)(nil)).Elem(), RecoveryPlanActionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanAutomationRunbookActionDetailsResponseInput)(nil)).Elem(), RecoveryPlanAutomationRunbookActionDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanGroupInput)(nil)).Elem(), RecoveryPlanGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanGroupArrayInput)(nil)).Elem(), RecoveryPlanGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanGroupResponseInput)(nil)).Elem(), RecoveryPlanGroupResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanGroupResponseArrayInput)(nil)).Elem(), RecoveryPlanGroupResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanManualActionDetailsResponseInput)(nil)).Elem(), RecoveryPlanManualActionDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanPropertiesResponseInput)(nil)).Elem(), RecoveryPlanPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanPropertiesResponsePtrInput)(nil)).Elem(), RecoveryPlanPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanProtectedItemInput)(nil)).Elem(), RecoveryPlanProtectedItemArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanProtectedItemArrayInput)(nil)).Elem(), RecoveryPlanProtectedItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanProtectedItemResponseInput)(nil)).Elem(), RecoveryPlanProtectedItemResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanProtectedItemResponseArrayInput)(nil)).Elem(), RecoveryPlanProtectedItemResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryPlanScriptActionDetailsResponseInput)(nil)).Elem(), RecoveryPlanScriptActionDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryServicesProviderPropertiesResponseInput)(nil)).Elem(), RecoveryServicesProviderPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RecoveryServicesProviderPropertiesResponsePtrInput)(nil)).Elem(), RecoveryServicesProviderPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationProtectedItemPropertiesResponseInput)(nil)).Elem(), ReplicationProtectedItemPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReplicationProtectedItemPropertiesResponsePtrInput)(nil)).Elem(), ReplicationProtectedItemPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RetentionVolumeResponseInput)(nil)).Elem(), RetentionVolumeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RetentionVolumeResponseArrayInput)(nil)).Elem(), RetentionVolumeResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAssignmentResponseInput)(nil)).Elem(), RoleAssignmentResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleAssignmentResponseArrayInput)(nil)).Elem(), RoleAssignmentResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunAsAccountResponseInput)(nil)).Elem(), RunAsAccountResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunAsAccountResponseArrayInput)(nil)).Elem(), RunAsAccountResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SanEnableProtectionInputInput)(nil)).Elem(), SanEnableProtectionInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageClassificationMappingPropertiesResponseInput)(nil)).Elem(), StorageClassificationMappingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageClassificationMappingPropertiesResponsePtrInput)(nil)).Elem(), StorageClassificationMappingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageMappingInputPropertiesInput)(nil)).Elem(), StorageMappingInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageMappingInputPropertiesPtrInput)(nil)).Elem(), StorageMappingInputPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VCenterPropertiesResponseInput)(nil)).Elem(), VCenterPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VCenterPropertiesResponsePtrInput)(nil)).Elem(), VCenterPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMNicDetailsResponseInput)(nil)).Elem(), VMNicDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMNicDetailsResponseArrayInput)(nil)).Elem(), VMNicDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtContainerMappingInputInput)(nil)).Elem(), VMwareCbtContainerMappingInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtDiskInputInput)(nil)).Elem(), VMwareCbtDiskInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtDiskInputArrayInput)(nil)).Elem(), VMwareCbtDiskInputArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtEnableMigrationInputInput)(nil)).Elem(), VMwareCbtEnableMigrationInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtEnableMigrationInputPtrInput)(nil)).Elem(), VMwareCbtEnableMigrationInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtMigrationDetailsResponseInput)(nil)).Elem(), VMwareCbtMigrationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtMigrationDetailsResponsePtrInput)(nil)).Elem(), VMwareCbtMigrationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtNicDetailsResponseInput)(nil)).Elem(), VMwareCbtNicDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtNicDetailsResponseArrayInput)(nil)).Elem(), VMwareCbtNicDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtPolicyCreationInputInput)(nil)).Elem(), VMwareCbtPolicyCreationInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtProtectedDiskDetailsResponseInput)(nil)).Elem(), VMwareCbtProtectedDiskDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtProtectedDiskDetailsResponseArrayInput)(nil)).Elem(), VMwareCbtProtectedDiskDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareCbtProtectionContainerMappingDetailsResponseInput)(nil)).Elem(), VMwareCbtProtectionContainerMappingDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareDetailsResponseInput)(nil)).Elem(), VMwareDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareV2FabricCreationInputInput)(nil)).Elem(), VMwareV2FabricCreationInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareV2FabricSpecificDetailsResponseInput)(nil)).Elem(), VMwareV2FabricSpecificDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VersionDetailsResponseInput)(nil)).Elem(), VersionDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VersionDetailsResponsePtrInput)(nil)).Elem(), VersionDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmmDetailsResponseInput)(nil)).Elem(), VmmDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmmToAzureCreateNetworkMappingInputInput)(nil)).Elem(), VmmToAzureCreateNetworkMappingInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmmToAzureNetworkMappingSettingsResponseInput)(nil)).Elem(), VmmToAzureNetworkMappingSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmmToVmmCreateNetworkMappingInputInput)(nil)).Elem(), VmmToVmmCreateNetworkMappingInputArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmmToVmmNetworkMappingSettingsResponseInput)(nil)).Elem(), VmmToVmmNetworkMappingSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmwareCbtPolicyDetailsResponseInput)(nil)).Elem(), VmwareCbtPolicyDetailsResponseArgs{})
	pulumi.RegisterOutputType(A2AContainerMappingInputOutput{})
	pulumi.RegisterOutputType(A2AEnableProtectionInputOutput{})
	pulumi.RegisterOutputType(A2APolicyCreationInputOutput{})
	pulumi.RegisterOutputType(A2APolicyDetailsResponseOutput{})
	pulumi.RegisterOutputType(A2AProtectedDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(A2AProtectedDiskDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(A2AProtectedManagedDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(A2AProtectedManagedDiskDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(A2AProtectionContainerMappingDetailsResponseOutput{})
	pulumi.RegisterOutputType(A2AReplicationDetailsResponseOutput{})
	pulumi.RegisterOutputType(A2AVmDiskInputDetailsOutput{})
	pulumi.RegisterOutputType(A2AVmDiskInputDetailsArrayOutput{})
	pulumi.RegisterOutputType(A2AVmManagedDiskInputDetailsOutput{})
	pulumi.RegisterOutputType(A2AVmManagedDiskInputDetailsArrayOutput{})
	pulumi.RegisterOutputType(AddRecoveryServicesProviderInputPropertiesOutput{})
	pulumi.RegisterOutputType(AddRecoveryServicesProviderInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AddVCenterRequestPropertiesOutput{})
	pulumi.RegisterOutputType(AddVCenterRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AzureFabricCreationInputOutput{})
	pulumi.RegisterOutputType(AzureFabricSpecificDetailsResponseOutput{})
	pulumi.RegisterOutputType(AzureToAzureCreateNetworkMappingInputOutput{})
	pulumi.RegisterOutputType(AzureToAzureNetworkMappingSettingsResponseOutput{})
	pulumi.RegisterOutputType(AzureToAzureVmSyncedConfigDetailsResponseOutput{})
	pulumi.RegisterOutputType(AzureToAzureVmSyncedConfigDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureVmDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(AzureVmDiskDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(CreateNetworkMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreateNetworkMappingInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreatePolicyInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreatePolicyInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreateProtectionContainerMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreateProtectionContainerMappingInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CreateRecoveryPlanInputPropertiesOutput{})
	pulumi.RegisterOutputType(CreateRecoveryPlanInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CurrentJobDetailsResponseOutput{})
	pulumi.RegisterOutputType(CurrentJobDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(CurrentScenarioDetailsResponseOutput{})
	pulumi.RegisterOutputType(CurrentScenarioDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(DataStoreResponseOutput{})
	pulumi.RegisterOutputType(DataStoreResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(DiskDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(DiskEncryptionInfoOutput{})
	pulumi.RegisterOutputType(DiskEncryptionInfoPtrOutput{})
	pulumi.RegisterOutputType(DiskEncryptionKeyInfoOutput{})
	pulumi.RegisterOutputType(DiskEncryptionKeyInfoPtrOutput{})
	pulumi.RegisterOutputType(EnableMigrationInputPropertiesOutput{})
	pulumi.RegisterOutputType(EnableMigrationInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EnableProtectionInputPropertiesOutput{})
	pulumi.RegisterOutputType(EnableProtectionInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsResponseOutput{})
	pulumi.RegisterOutputType(EncryptionDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(FabricCreationInputPropertiesOutput{})
	pulumi.RegisterOutputType(FabricCreationInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(FabricPropertiesResponseOutput{})
	pulumi.RegisterOutputType(FabricPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(HealthErrorResponseOutput{})
	pulumi.RegisterOutputType(HealthErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(HyperVReplicaAzureEnableProtectionInputOutput{})
	pulumi.RegisterOutputType(HyperVReplicaAzurePolicyDetailsResponseOutput{})
	pulumi.RegisterOutputType(HyperVReplicaAzurePolicyInputOutput{})
	pulumi.RegisterOutputType(HyperVReplicaAzureReplicationDetailsResponseOutput{})
	pulumi.RegisterOutputType(HyperVReplicaBasePolicyDetailsResponseOutput{})
	pulumi.RegisterOutputType(HyperVReplicaBaseReplicationDetailsResponseOutput{})
	pulumi.RegisterOutputType(HyperVReplicaBluePolicyDetailsResponseOutput{})
	pulumi.RegisterOutputType(HyperVReplicaBluePolicyInputOutput{})
	pulumi.RegisterOutputType(HyperVReplicaBlueReplicationDetailsResponseOutput{})
	pulumi.RegisterOutputType(HyperVReplicaPolicyDetailsResponseOutput{})
	pulumi.RegisterOutputType(HyperVReplicaPolicyInputOutput{})
	pulumi.RegisterOutputType(HyperVReplicaReplicationDetailsResponseOutput{})
	pulumi.RegisterOutputType(HyperVSiteDetailsResponseOutput{})
	pulumi.RegisterOutputType(IdentityProviderDetailsResponseOutput{})
	pulumi.RegisterOutputType(IdentityProviderDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityProviderInputOutput{})
	pulumi.RegisterOutputType(IdentityProviderInputPtrOutput{})
	pulumi.RegisterOutputType(InMageAgentDetailsResponseOutput{})
	pulumi.RegisterOutputType(InMageAgentDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(InMageAzureV2EnableProtectionInputOutput{})
	pulumi.RegisterOutputType(InMageAzureV2PolicyDetailsResponseOutput{})
	pulumi.RegisterOutputType(InMageAzureV2PolicyInputOutput{})
	pulumi.RegisterOutputType(InMageAzureV2ProtectedDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(InMageAzureV2ProtectedDiskDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(InMageAzureV2ReplicationDetailsResponseOutput{})
	pulumi.RegisterOutputType(InMageBasePolicyDetailsResponseOutput{})
	pulumi.RegisterOutputType(InMageDiskExclusionInputOutput{})
	pulumi.RegisterOutputType(InMageDiskExclusionInputPtrOutput{})
	pulumi.RegisterOutputType(InMageDiskSignatureExclusionOptionsOutput{})
	pulumi.RegisterOutputType(InMageDiskSignatureExclusionOptionsArrayOutput{})
	pulumi.RegisterOutputType(InMageEnableProtectionInputOutput{})
	pulumi.RegisterOutputType(InMagePolicyDetailsResponseOutput{})
	pulumi.RegisterOutputType(InMagePolicyInputOutput{})
	pulumi.RegisterOutputType(InMageProtectedDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(InMageProtectedDiskDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(InMageReplicationDetailsResponseOutput{})
	pulumi.RegisterOutputType(InMageVolumeExclusionOptionsOutput{})
	pulumi.RegisterOutputType(InMageVolumeExclusionOptionsArrayOutput{})
	pulumi.RegisterOutputType(InitialReplicationDetailsResponseOutput{})
	pulumi.RegisterOutputType(InitialReplicationDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(InnerHealthErrorResponseOutput{})
	pulumi.RegisterOutputType(InnerHealthErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(InputEndpointResponseOutput{})
	pulumi.RegisterOutputType(InputEndpointResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyEncryptionKeyInfoOutput{})
	pulumi.RegisterOutputType(KeyEncryptionKeyInfoPtrOutput{})
	pulumi.RegisterOutputType(MasterTargetServerResponseOutput{})
	pulumi.RegisterOutputType(MasterTargetServerResponseArrayOutput{})
	pulumi.RegisterOutputType(MigrationItemPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MigrationItemPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MobilityServiceUpdateResponseOutput{})
	pulumi.RegisterOutputType(MobilityServiceUpdateResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(NetworkMappingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(OSDetailsResponseOutput{})
	pulumi.RegisterOutputType(OSDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(OSDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(OSDiskDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(PolicyPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PolicyPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ProcessServerResponseOutput{})
	pulumi.RegisterOutputType(ProcessServerResponseArrayOutput{})
	pulumi.RegisterOutputType(ProtectionContainerMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ProtectionContainerMappingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RcmAzureMigrationPolicyDetailsResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanActionOutput{})
	pulumi.RegisterOutputType(RecoveryPlanActionArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanActionResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanActionResponseArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanAutomationRunbookActionDetailsResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanGroupOutput{})
	pulumi.RegisterOutputType(RecoveryPlanGroupArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanGroupResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanGroupResponseArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanManualActionDetailsResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RecoveryPlanProtectedItemOutput{})
	pulumi.RegisterOutputType(RecoveryPlanProtectedItemArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanProtectedItemResponseOutput{})
	pulumi.RegisterOutputType(RecoveryPlanProtectedItemResponseArrayOutput{})
	pulumi.RegisterOutputType(RecoveryPlanScriptActionDetailsResponseOutput{})
	pulumi.RegisterOutputType(RecoveryServicesProviderPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RecoveryServicesProviderPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ReplicationProtectedItemPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ReplicationProtectedItemPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RetentionVolumeResponseOutput{})
	pulumi.RegisterOutputType(RetentionVolumeResponseArrayOutput{})
	pulumi.RegisterOutputType(RoleAssignmentResponseOutput{})
	pulumi.RegisterOutputType(RoleAssignmentResponseArrayOutput{})
	pulumi.RegisterOutputType(RunAsAccountResponseOutput{})
	pulumi.RegisterOutputType(RunAsAccountResponseArrayOutput{})
	pulumi.RegisterOutputType(SanEnableProtectionInputOutput{})
	pulumi.RegisterOutputType(StorageClassificationMappingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageClassificationMappingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageMappingInputPropertiesOutput{})
	pulumi.RegisterOutputType(StorageMappingInputPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VCenterPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VCenterPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(VMNicDetailsResponseOutput{})
	pulumi.RegisterOutputType(VMNicDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(VMwareCbtContainerMappingInputOutput{})
	pulumi.RegisterOutputType(VMwareCbtDiskInputOutput{})
	pulumi.RegisterOutputType(VMwareCbtDiskInputArrayOutput{})
	pulumi.RegisterOutputType(VMwareCbtEnableMigrationInputOutput{})
	pulumi.RegisterOutputType(VMwareCbtEnableMigrationInputPtrOutput{})
	pulumi.RegisterOutputType(VMwareCbtMigrationDetailsResponseOutput{})
	pulumi.RegisterOutputType(VMwareCbtMigrationDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(VMwareCbtNicDetailsResponseOutput{})
	pulumi.RegisterOutputType(VMwareCbtNicDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(VMwareCbtPolicyCreationInputOutput{})
	pulumi.RegisterOutputType(VMwareCbtProtectedDiskDetailsResponseOutput{})
	pulumi.RegisterOutputType(VMwareCbtProtectedDiskDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(VMwareCbtProtectionContainerMappingDetailsResponseOutput{})
	pulumi.RegisterOutputType(VMwareDetailsResponseOutput{})
	pulumi.RegisterOutputType(VMwareV2FabricCreationInputOutput{})
	pulumi.RegisterOutputType(VMwareV2FabricSpecificDetailsResponseOutput{})
	pulumi.RegisterOutputType(VersionDetailsResponseOutput{})
	pulumi.RegisterOutputType(VersionDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(VmmDetailsResponseOutput{})
	pulumi.RegisterOutputType(VmmToAzureCreateNetworkMappingInputOutput{})
	pulumi.RegisterOutputType(VmmToAzureNetworkMappingSettingsResponseOutput{})
	pulumi.RegisterOutputType(VmmToVmmCreateNetworkMappingInputOutput{})
	pulumi.RegisterOutputType(VmmToVmmNetworkMappingSettingsResponseOutput{})
	pulumi.RegisterOutputType(VmwareCbtPolicyDetailsResponseOutput{})
}
