


package v20170701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureRecoveryServiceVaultProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType *string `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureRecoveryServiceVaultProtectionIntentInput interface {
	pulumi.Input

	ToAzureRecoveryServiceVaultProtectionIntentOutput() AzureRecoveryServiceVaultProtectionIntentOutput
	ToAzureRecoveryServiceVaultProtectionIntentOutputWithContext(context.Context) AzureRecoveryServiceVaultProtectionIntentOutput
}

type AzureRecoveryServiceVaultProtectionIntentArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringPtrInput `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureRecoveryServiceVaultProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRecoveryServiceVaultProtectionIntent)(nil)).Elem()
}

func (i AzureRecoveryServiceVaultProtectionIntentArgs) ToAzureRecoveryServiceVaultProtectionIntentOutput() AzureRecoveryServiceVaultProtectionIntentOutput {
	return i.ToAzureRecoveryServiceVaultProtectionIntentOutputWithContext(context.Background())
}

func (i AzureRecoveryServiceVaultProtectionIntentArgs) ToAzureRecoveryServiceVaultProtectionIntentOutputWithContext(ctx context.Context) AzureRecoveryServiceVaultProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRecoveryServiceVaultProtectionIntentOutput)
}

type AzureRecoveryServiceVaultProtectionIntentOutput struct{ *pulumi.OutputState }

func (AzureRecoveryServiceVaultProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRecoveryServiceVaultProtectionIntent)(nil)).Elem()
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) ToAzureRecoveryServiceVaultProtectionIntentOutput() AzureRecoveryServiceVaultProtectionIntentOutput {
	return o
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) ToAzureRecoveryServiceVaultProtectionIntentOutputWithContext(ctx context.Context) AzureRecoveryServiceVaultProtectionIntentOutput {
	return o
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) ProtectionIntentItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.ProtectionIntentItemType }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntent) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureRecoveryServiceVaultProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType *string `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureRecoveryServiceVaultProtectionIntentResponseInput interface {
	pulumi.Input

	ToAzureRecoveryServiceVaultProtectionIntentResponseOutput() AzureRecoveryServiceVaultProtectionIntentResponseOutput
	ToAzureRecoveryServiceVaultProtectionIntentResponseOutputWithContext(context.Context) AzureRecoveryServiceVaultProtectionIntentResponseOutput
}

type AzureRecoveryServiceVaultProtectionIntentResponseArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringPtrInput `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureRecoveryServiceVaultProtectionIntentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRecoveryServiceVaultProtectionIntentResponse)(nil)).Elem()
}

func (i AzureRecoveryServiceVaultProtectionIntentResponseArgs) ToAzureRecoveryServiceVaultProtectionIntentResponseOutput() AzureRecoveryServiceVaultProtectionIntentResponseOutput {
	return i.ToAzureRecoveryServiceVaultProtectionIntentResponseOutputWithContext(context.Background())
}

func (i AzureRecoveryServiceVaultProtectionIntentResponseArgs) ToAzureRecoveryServiceVaultProtectionIntentResponseOutputWithContext(ctx context.Context) AzureRecoveryServiceVaultProtectionIntentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureRecoveryServiceVaultProtectionIntentResponseOutput)
}

type AzureRecoveryServiceVaultProtectionIntentResponseOutput struct{ *pulumi.OutputState }

func (AzureRecoveryServiceVaultProtectionIntentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureRecoveryServiceVaultProtectionIntentResponse)(nil)).Elem()
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) ToAzureRecoveryServiceVaultProtectionIntentResponseOutput() AzureRecoveryServiceVaultProtectionIntentResponseOutput {
	return o
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) ToAzureRecoveryServiceVaultProtectionIntentResponseOutputWithContext(ctx context.Context) AzureRecoveryServiceVaultProtectionIntentResponseOutput {
	return o
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) ProtectionIntentItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.ProtectionIntentItemType }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureRecoveryServiceVaultProtectionIntentResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureRecoveryServiceVaultProtectionIntentResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureResourceProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	FriendlyName             *string `pulumi:"friendlyName"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType *string `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureResourceProtectionIntentInput interface {
	pulumi.Input

	ToAzureResourceProtectionIntentOutput() AzureResourceProtectionIntentOutput
	ToAzureResourceProtectionIntentOutputWithContext(context.Context) AzureResourceProtectionIntentOutput
}

type AzureResourceProtectionIntentArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	FriendlyName             pulumi.StringPtrInput `pulumi:"friendlyName"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringPtrInput `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureResourceProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceProtectionIntent)(nil)).Elem()
}

func (i AzureResourceProtectionIntentArgs) ToAzureResourceProtectionIntentOutput() AzureResourceProtectionIntentOutput {
	return i.ToAzureResourceProtectionIntentOutputWithContext(context.Background())
}

func (i AzureResourceProtectionIntentArgs) ToAzureResourceProtectionIntentOutputWithContext(ctx context.Context) AzureResourceProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureResourceProtectionIntentOutput)
}

type AzureResourceProtectionIntentOutput struct{ *pulumi.OutputState }

func (AzureResourceProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceProtectionIntent)(nil)).Elem()
}

func (o AzureResourceProtectionIntentOutput) ToAzureResourceProtectionIntentOutput() AzureResourceProtectionIntentOutput {
	return o
}

func (o AzureResourceProtectionIntentOutput) ToAzureResourceProtectionIntentOutputWithContext(ctx context.Context) AzureResourceProtectionIntentOutput {
	return o
}

func (o AzureResourceProtectionIntentOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) ProtectionIntentItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.ProtectionIntentItemType }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntent) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureResourceProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	FriendlyName             *string `pulumi:"friendlyName"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType *string `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureResourceProtectionIntentResponseInput interface {
	pulumi.Input

	ToAzureResourceProtectionIntentResponseOutput() AzureResourceProtectionIntentResponseOutput
	ToAzureResourceProtectionIntentResponseOutputWithContext(context.Context) AzureResourceProtectionIntentResponseOutput
}

type AzureResourceProtectionIntentResponseArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	FriendlyName             pulumi.StringPtrInput `pulumi:"friendlyName"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringPtrInput `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureResourceProtectionIntentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceProtectionIntentResponse)(nil)).Elem()
}

func (i AzureResourceProtectionIntentResponseArgs) ToAzureResourceProtectionIntentResponseOutput() AzureResourceProtectionIntentResponseOutput {
	return i.ToAzureResourceProtectionIntentResponseOutputWithContext(context.Background())
}

func (i AzureResourceProtectionIntentResponseArgs) ToAzureResourceProtectionIntentResponseOutputWithContext(ctx context.Context) AzureResourceProtectionIntentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureResourceProtectionIntentResponseOutput)
}

type AzureResourceProtectionIntentResponseOutput struct{ *pulumi.OutputState }

func (AzureResourceProtectionIntentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceProtectionIntentResponse)(nil)).Elem()
}

func (o AzureResourceProtectionIntentResponseOutput) ToAzureResourceProtectionIntentResponseOutput() AzureResourceProtectionIntentResponseOutput {
	return o
}

func (o AzureResourceProtectionIntentResponseOutput) ToAzureResourceProtectionIntentResponseOutputWithContext(ctx context.Context) AzureResourceProtectionIntentResponseOutput {
	return o
}

func (o AzureResourceProtectionIntentResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) ProtectionIntentItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.ProtectionIntentItemType }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureResourceProtectionIntentResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureResourceProtectionIntentResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureWorkloadAutoProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType *string `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureWorkloadAutoProtectionIntentInput interface {
	pulumi.Input

	ToAzureWorkloadAutoProtectionIntentOutput() AzureWorkloadAutoProtectionIntentOutput
	ToAzureWorkloadAutoProtectionIntentOutputWithContext(context.Context) AzureWorkloadAutoProtectionIntentOutput
}

type AzureWorkloadAutoProtectionIntentArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringPtrInput `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureWorkloadAutoProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadAutoProtectionIntent)(nil)).Elem()
}

func (i AzureWorkloadAutoProtectionIntentArgs) ToAzureWorkloadAutoProtectionIntentOutput() AzureWorkloadAutoProtectionIntentOutput {
	return i.ToAzureWorkloadAutoProtectionIntentOutputWithContext(context.Background())
}

func (i AzureWorkloadAutoProtectionIntentArgs) ToAzureWorkloadAutoProtectionIntentOutputWithContext(ctx context.Context) AzureWorkloadAutoProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadAutoProtectionIntentOutput)
}

type AzureWorkloadAutoProtectionIntentOutput struct{ *pulumi.OutputState }

func (AzureWorkloadAutoProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadAutoProtectionIntent)(nil)).Elem()
}

func (o AzureWorkloadAutoProtectionIntentOutput) ToAzureWorkloadAutoProtectionIntentOutput() AzureWorkloadAutoProtectionIntentOutput {
	return o
}

func (o AzureWorkloadAutoProtectionIntentOutput) ToAzureWorkloadAutoProtectionIntentOutputWithContext(ctx context.Context) AzureWorkloadAutoProtectionIntentOutput {
	return o
}

func (o AzureWorkloadAutoProtectionIntentOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentOutput) ProtectionIntentItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.ProtectionIntentItemType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntent) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureWorkloadAutoProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType *string `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
}





type AzureWorkloadAutoProtectionIntentResponseInput interface {
	pulumi.Input

	ToAzureWorkloadAutoProtectionIntentResponseOutput() AzureWorkloadAutoProtectionIntentResponseOutput
	ToAzureWorkloadAutoProtectionIntentResponseOutputWithContext(context.Context) AzureWorkloadAutoProtectionIntentResponseOutput
}

type AzureWorkloadAutoProtectionIntentResponseArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringPtrInput `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
}

func (AzureWorkloadAutoProtectionIntentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadAutoProtectionIntentResponse)(nil)).Elem()
}

func (i AzureWorkloadAutoProtectionIntentResponseArgs) ToAzureWorkloadAutoProtectionIntentResponseOutput() AzureWorkloadAutoProtectionIntentResponseOutput {
	return i.ToAzureWorkloadAutoProtectionIntentResponseOutputWithContext(context.Background())
}

func (i AzureWorkloadAutoProtectionIntentResponseArgs) ToAzureWorkloadAutoProtectionIntentResponseOutputWithContext(ctx context.Context) AzureWorkloadAutoProtectionIntentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadAutoProtectionIntentResponseOutput)
}

type AzureWorkloadAutoProtectionIntentResponseOutput struct{ *pulumi.OutputState }

func (AzureWorkloadAutoProtectionIntentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadAutoProtectionIntentResponse)(nil)).Elem()
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) ToAzureWorkloadAutoProtectionIntentResponseOutput() AzureWorkloadAutoProtectionIntentResponseOutput {
	return o
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) ToAzureWorkloadAutoProtectionIntentResponseOutputWithContext(ctx context.Context) AzureWorkloadAutoProtectionIntentResponseOutput {
	return o
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) ProtectionIntentItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.ProtectionIntentItemType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadAutoProtectionIntentResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadAutoProtectionIntentResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

type AzureWorkloadSQLAutoProtectionIntent struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType *string `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
	WorkloadItemType         *string `pulumi:"workloadItemType"`
}





type AzureWorkloadSQLAutoProtectionIntentInput interface {
	pulumi.Input

	ToAzureWorkloadSQLAutoProtectionIntentOutput() AzureWorkloadSQLAutoProtectionIntentOutput
	ToAzureWorkloadSQLAutoProtectionIntentOutputWithContext(context.Context) AzureWorkloadSQLAutoProtectionIntentOutput
}

type AzureWorkloadSQLAutoProtectionIntentArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringPtrInput `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
	WorkloadItemType         pulumi.StringPtrInput `pulumi:"workloadItemType"`
}

func (AzureWorkloadSQLAutoProtectionIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadSQLAutoProtectionIntent)(nil)).Elem()
}

func (i AzureWorkloadSQLAutoProtectionIntentArgs) ToAzureWorkloadSQLAutoProtectionIntentOutput() AzureWorkloadSQLAutoProtectionIntentOutput {
	return i.ToAzureWorkloadSQLAutoProtectionIntentOutputWithContext(context.Background())
}

func (i AzureWorkloadSQLAutoProtectionIntentArgs) ToAzureWorkloadSQLAutoProtectionIntentOutputWithContext(ctx context.Context) AzureWorkloadSQLAutoProtectionIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadSQLAutoProtectionIntentOutput)
}

type AzureWorkloadSQLAutoProtectionIntentOutput struct{ *pulumi.OutputState }

func (AzureWorkloadSQLAutoProtectionIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadSQLAutoProtectionIntent)(nil)).Elem()
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) ToAzureWorkloadSQLAutoProtectionIntentOutput() AzureWorkloadSQLAutoProtectionIntentOutput {
	return o
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) ToAzureWorkloadSQLAutoProtectionIntentOutputWithContext(ctx context.Context) AzureWorkloadSQLAutoProtectionIntentOutput {
	return o
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) ProtectionIntentItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.ProtectionIntentItemType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentOutput) WorkloadItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntent) *string { return v.WorkloadItemType }).(pulumi.StringPtrOutput)
}

type AzureWorkloadSQLAutoProtectionIntentResponse struct {
	BackupManagementType     *string `pulumi:"backupManagementType"`
	ItemId                   *string `pulumi:"itemId"`
	PolicyId                 *string `pulumi:"policyId"`
	ProtectionIntentItemType *string `pulumi:"protectionIntentItemType"`
	ProtectionState          *string `pulumi:"protectionState"`
	SourceResourceId         *string `pulumi:"sourceResourceId"`
	WorkloadItemType         *string `pulumi:"workloadItemType"`
}





type AzureWorkloadSQLAutoProtectionIntentResponseInput interface {
	pulumi.Input

	ToAzureWorkloadSQLAutoProtectionIntentResponseOutput() AzureWorkloadSQLAutoProtectionIntentResponseOutput
	ToAzureWorkloadSQLAutoProtectionIntentResponseOutputWithContext(context.Context) AzureWorkloadSQLAutoProtectionIntentResponseOutput
}

type AzureWorkloadSQLAutoProtectionIntentResponseArgs struct {
	BackupManagementType     pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ItemId                   pulumi.StringPtrInput `pulumi:"itemId"`
	PolicyId                 pulumi.StringPtrInput `pulumi:"policyId"`
	ProtectionIntentItemType pulumi.StringPtrInput `pulumi:"protectionIntentItemType"`
	ProtectionState          pulumi.StringPtrInput `pulumi:"protectionState"`
	SourceResourceId         pulumi.StringPtrInput `pulumi:"sourceResourceId"`
	WorkloadItemType         pulumi.StringPtrInput `pulumi:"workloadItemType"`
}

func (AzureWorkloadSQLAutoProtectionIntentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadSQLAutoProtectionIntentResponse)(nil)).Elem()
}

func (i AzureWorkloadSQLAutoProtectionIntentResponseArgs) ToAzureWorkloadSQLAutoProtectionIntentResponseOutput() AzureWorkloadSQLAutoProtectionIntentResponseOutput {
	return i.ToAzureWorkloadSQLAutoProtectionIntentResponseOutputWithContext(context.Background())
}

func (i AzureWorkloadSQLAutoProtectionIntentResponseArgs) ToAzureWorkloadSQLAutoProtectionIntentResponseOutputWithContext(ctx context.Context) AzureWorkloadSQLAutoProtectionIntentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureWorkloadSQLAutoProtectionIntentResponseOutput)
}

type AzureWorkloadSQLAutoProtectionIntentResponseOutput struct{ *pulumi.OutputState }

func (AzureWorkloadSQLAutoProtectionIntentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureWorkloadSQLAutoProtectionIntentResponse)(nil)).Elem()
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) ToAzureWorkloadSQLAutoProtectionIntentResponseOutput() AzureWorkloadSQLAutoProtectionIntentResponseOutput {
	return o
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) ToAzureWorkloadSQLAutoProtectionIntentResponseOutputWithContext(ctx context.Context) AzureWorkloadSQLAutoProtectionIntentResponseOutput {
	return o
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) ItemId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.ItemId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) ProtectionIntentItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.ProtectionIntentItemType }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) ProtectionState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.ProtectionState }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) SourceResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.SourceResourceId }).(pulumi.StringPtrOutput)
}

func (o AzureWorkloadSQLAutoProtectionIntentResponseOutput) WorkloadItemType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureWorkloadSQLAutoProtectionIntentResponse) *string { return v.WorkloadItemType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureRecoveryServiceVaultProtectionIntentInput)(nil)).Elem(), AzureRecoveryServiceVaultProtectionIntentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureRecoveryServiceVaultProtectionIntentResponseInput)(nil)).Elem(), AzureRecoveryServiceVaultProtectionIntentResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureResourceProtectionIntentInput)(nil)).Elem(), AzureResourceProtectionIntentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureResourceProtectionIntentResponseInput)(nil)).Elem(), AzureResourceProtectionIntentResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureWorkloadAutoProtectionIntentInput)(nil)).Elem(), AzureWorkloadAutoProtectionIntentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureWorkloadAutoProtectionIntentResponseInput)(nil)).Elem(), AzureWorkloadAutoProtectionIntentResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureWorkloadSQLAutoProtectionIntentInput)(nil)).Elem(), AzureWorkloadSQLAutoProtectionIntentArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureWorkloadSQLAutoProtectionIntentResponseInput)(nil)).Elem(), AzureWorkloadSQLAutoProtectionIntentResponseArgs{})
	pulumi.RegisterOutputType(AzureRecoveryServiceVaultProtectionIntentOutput{})
	pulumi.RegisterOutputType(AzureRecoveryServiceVaultProtectionIntentResponseOutput{})
	pulumi.RegisterOutputType(AzureResourceProtectionIntentOutput{})
	pulumi.RegisterOutputType(AzureResourceProtectionIntentResponseOutput{})
	pulumi.RegisterOutputType(AzureWorkloadAutoProtectionIntentOutput{})
	pulumi.RegisterOutputType(AzureWorkloadAutoProtectionIntentResponseOutput{})
	pulumi.RegisterOutputType(AzureWorkloadSQLAutoProtectionIntentOutput{})
	pulumi.RegisterOutputType(AzureWorkloadSQLAutoProtectionIntentResponseOutput{})
}
