


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureSkuName string

const (
	AzureSkuName_Standard_DS13_v2_1TB_PS    = AzureSkuName("Standard_DS13_v2+1TB_PS")
	AzureSkuName_Standard_DS13_v2_2TB_PS    = AzureSkuName("Standard_DS13_v2+2TB_PS")
	AzureSkuName_Standard_DS14_v2_3TB_PS    = AzureSkuName("Standard_DS14_v2+3TB_PS")
	AzureSkuName_Standard_DS14_v2_4TB_PS    = AzureSkuName("Standard_DS14_v2+4TB_PS")
	AzureSkuName_Standard_D13_v2            = AzureSkuName("Standard_D13_v2")
	AzureSkuName_Standard_D14_v2            = AzureSkuName("Standard_D14_v2")
	AzureSkuName_Standard_L8s               = AzureSkuName("Standard_L8s")
	AzureSkuName_Standard_L16s              = AzureSkuName("Standard_L16s")
	AzureSkuName_Standard_L8s_v2            = AzureSkuName("Standard_L8s_v2")
	AzureSkuName_Standard_L16s_v2           = AzureSkuName("Standard_L16s_v2")
	AzureSkuName_Standard_D11_v2            = AzureSkuName("Standard_D11_v2")
	AzureSkuName_Standard_D12_v2            = AzureSkuName("Standard_D12_v2")
	AzureSkuName_Standard_L4s               = AzureSkuName("Standard_L4s")
	AzureSkuName_Dev_No_SLA_Standard_D11_v2 = AzureSkuName("Dev(No SLA)_Standard_D11_v2")
	AzureSkuName_Standard_E64i_v3           = AzureSkuName("Standard_E64i_v3")
	AzureSkuName_Standard_E80ids_v4         = AzureSkuName("Standard_E80ids_v4")
	AzureSkuName_Standard_E2a_v4            = AzureSkuName("Standard_E2a_v4")
	AzureSkuName_Standard_E4a_v4            = AzureSkuName("Standard_E4a_v4")
	AzureSkuName_Standard_E8a_v4            = AzureSkuName("Standard_E8a_v4")
	AzureSkuName_Standard_E16a_v4           = AzureSkuName("Standard_E16a_v4")
	AzureSkuName_Standard_E8as_v4_1TB_PS    = AzureSkuName("Standard_E8as_v4+1TB_PS")
	AzureSkuName_Standard_E8as_v4_2TB_PS    = AzureSkuName("Standard_E8as_v4+2TB_PS")
	AzureSkuName_Standard_E16as_v4_3TB_PS   = AzureSkuName("Standard_E16as_v4+3TB_PS")
	AzureSkuName_Standard_E16as_v4_4TB_PS   = AzureSkuName("Standard_E16as_v4+4TB_PS")
	AzureSkuName_Dev_No_SLA_Standard_E2a_v4 = AzureSkuName("Dev(No SLA)_Standard_E2a_v4")
)

func (AzureSkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuName)(nil)).Elem()
}

func (e AzureSkuName) ToAzureSkuNameOutput() AzureSkuNameOutput {
	return pulumi.ToOutput(e).(AzureSkuNameOutput)
}

func (e AzureSkuName) ToAzureSkuNameOutputWithContext(ctx context.Context) AzureSkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureSkuNameOutput)
}

func (e AzureSkuName) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return e.ToAzureSkuNamePtrOutputWithContext(context.Background())
}

func (e AzureSkuName) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return AzureSkuName(e).ToAzureSkuNameOutputWithContext(ctx).ToAzureSkuNamePtrOutputWithContext(ctx)
}

func (e AzureSkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureSkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureSkuNameOutput struct{ *pulumi.OutputState }

func (AzureSkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuName)(nil)).Elem()
}

func (o AzureSkuNameOutput) ToAzureSkuNameOutput() AzureSkuNameOutput {
	return o
}

func (o AzureSkuNameOutput) ToAzureSkuNameOutputWithContext(ctx context.Context) AzureSkuNameOutput {
	return o
}

func (o AzureSkuNameOutput) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return o.ToAzureSkuNamePtrOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSkuName) *AzureSkuName {
		return &v
	}).(AzureSkuNamePtrOutput)
}

func (o AzureSkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureSkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureSkuNamePtrOutput struct{ *pulumi.OutputState }

func (AzureSkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuName)(nil)).Elem()
}

func (o AzureSkuNamePtrOutput) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return o
}

func (o AzureSkuNamePtrOutput) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return o
}

func (o AzureSkuNamePtrOutput) Elem() AzureSkuNameOutput {
	return o.ApplyT(func(v *AzureSkuName) AzureSkuName {
		if v != nil {
			return *v
		}
		var ret AzureSkuName
		return ret
	}).(AzureSkuNameOutput)
}

func (o AzureSkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureSkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureSkuNameInput interface {
	pulumi.Input

	ToAzureSkuNameOutput() AzureSkuNameOutput
	ToAzureSkuNameOutputWithContext(context.Context) AzureSkuNameOutput
}

var azureSkuNamePtrType = reflect.TypeOf((**AzureSkuName)(nil)).Elem()

type AzureSkuNamePtrInput interface {
	pulumi.Input

	ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput
	ToAzureSkuNamePtrOutputWithContext(context.Context) AzureSkuNamePtrOutput
}

type azureSkuNamePtr string

func AzureSkuNamePtr(v string) AzureSkuNamePtrInput {
	return (*azureSkuNamePtr)(&v)
}

func (*azureSkuNamePtr) ElementType() reflect.Type {
	return azureSkuNamePtrType
}

func (in *azureSkuNamePtr) ToAzureSkuNamePtrOutput() AzureSkuNamePtrOutput {
	return pulumi.ToOutput(in).(AzureSkuNamePtrOutput)
}

func (in *azureSkuNamePtr) ToAzureSkuNamePtrOutputWithContext(ctx context.Context) AzureSkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureSkuNamePtrOutput)
}

type AzureSkuTier string

const (
	AzureSkuTierBasic    = AzureSkuTier("Basic")
	AzureSkuTierStandard = AzureSkuTier("Standard")
)

func (AzureSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuTier)(nil)).Elem()
}

func (e AzureSkuTier) ToAzureSkuTierOutput() AzureSkuTierOutput {
	return pulumi.ToOutput(e).(AzureSkuTierOutput)
}

func (e AzureSkuTier) ToAzureSkuTierOutputWithContext(ctx context.Context) AzureSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AzureSkuTierOutput)
}

func (e AzureSkuTier) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return e.ToAzureSkuTierPtrOutputWithContext(context.Background())
}

func (e AzureSkuTier) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return AzureSkuTier(e).ToAzureSkuTierOutputWithContext(ctx).ToAzureSkuTierPtrOutputWithContext(ctx)
}

func (e AzureSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AzureSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AzureSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AzureSkuTierOutput struct{ *pulumi.OutputState }

func (AzureSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuTier)(nil)).Elem()
}

func (o AzureSkuTierOutput) ToAzureSkuTierOutput() AzureSkuTierOutput {
	return o
}

func (o AzureSkuTierOutput) ToAzureSkuTierOutputWithContext(ctx context.Context) AzureSkuTierOutput {
	return o
}

func (o AzureSkuTierOutput) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return o.ToAzureSkuTierPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureSkuTier) *AzureSkuTier {
		return &v
	}).(AzureSkuTierPtrOutput)
}

func (o AzureSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AzureSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AzureSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AzureSkuTierPtrOutput struct{ *pulumi.OutputState }

func (AzureSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuTier)(nil)).Elem()
}

func (o AzureSkuTierPtrOutput) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return o
}

func (o AzureSkuTierPtrOutput) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return o
}

func (o AzureSkuTierPtrOutput) Elem() AzureSkuTierOutput {
	return o.ApplyT(func(v *AzureSkuTier) AzureSkuTier {
		if v != nil {
			return *v
		}
		var ret AzureSkuTier
		return ret
	}).(AzureSkuTierOutput)
}

func (o AzureSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AzureSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AzureSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AzureSkuTierInput interface {
	pulumi.Input

	ToAzureSkuTierOutput() AzureSkuTierOutput
	ToAzureSkuTierOutputWithContext(context.Context) AzureSkuTierOutput
}

var azureSkuTierPtrType = reflect.TypeOf((**AzureSkuTier)(nil)).Elem()

type AzureSkuTierPtrInput interface {
	pulumi.Input

	ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput
	ToAzureSkuTierPtrOutputWithContext(context.Context) AzureSkuTierPtrOutput
}

type azureSkuTierPtr string

func AzureSkuTierPtr(v string) AzureSkuTierPtrInput {
	return (*azureSkuTierPtr)(&v)
}

func (*azureSkuTierPtr) ElementType() reflect.Type {
	return azureSkuTierPtrType
}

func (in *azureSkuTierPtr) ToAzureSkuTierPtrOutput() AzureSkuTierPtrOutput {
	return pulumi.ToOutput(in).(AzureSkuTierPtrOutput)
}

func (in *azureSkuTierPtr) ToAzureSkuTierPtrOutputWithContext(ctx context.Context) AzureSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AzureSkuTierPtrOutput)
}

type BlobStorageEventType string

const (
	BlobStorageEventType_Microsoft_Storage_BlobCreated = BlobStorageEventType("Microsoft.Storage.BlobCreated")
	BlobStorageEventType_Microsoft_Storage_BlobRenamed = BlobStorageEventType("Microsoft.Storage.BlobRenamed")
)

func (BlobStorageEventType) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStorageEventType)(nil)).Elem()
}

func (e BlobStorageEventType) ToBlobStorageEventTypeOutput() BlobStorageEventTypeOutput {
	return pulumi.ToOutput(e).(BlobStorageEventTypeOutput)
}

func (e BlobStorageEventType) ToBlobStorageEventTypeOutputWithContext(ctx context.Context) BlobStorageEventTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BlobStorageEventTypeOutput)
}

func (e BlobStorageEventType) ToBlobStorageEventTypePtrOutput() BlobStorageEventTypePtrOutput {
	return e.ToBlobStorageEventTypePtrOutputWithContext(context.Background())
}

func (e BlobStorageEventType) ToBlobStorageEventTypePtrOutputWithContext(ctx context.Context) BlobStorageEventTypePtrOutput {
	return BlobStorageEventType(e).ToBlobStorageEventTypeOutputWithContext(ctx).ToBlobStorageEventTypePtrOutputWithContext(ctx)
}

func (e BlobStorageEventType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobStorageEventType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BlobStorageEventType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BlobStorageEventType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BlobStorageEventTypeOutput struct{ *pulumi.OutputState }

func (BlobStorageEventTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStorageEventType)(nil)).Elem()
}

func (o BlobStorageEventTypeOutput) ToBlobStorageEventTypeOutput() BlobStorageEventTypeOutput {
	return o
}

func (o BlobStorageEventTypeOutput) ToBlobStorageEventTypeOutputWithContext(ctx context.Context) BlobStorageEventTypeOutput {
	return o
}

func (o BlobStorageEventTypeOutput) ToBlobStorageEventTypePtrOutput() BlobStorageEventTypePtrOutput {
	return o.ToBlobStorageEventTypePtrOutputWithContext(context.Background())
}

func (o BlobStorageEventTypeOutput) ToBlobStorageEventTypePtrOutputWithContext(ctx context.Context) BlobStorageEventTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BlobStorageEventType) *BlobStorageEventType {
		return &v
	}).(BlobStorageEventTypePtrOutput)
}

func (o BlobStorageEventTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BlobStorageEventTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobStorageEventType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BlobStorageEventTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobStorageEventTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BlobStorageEventType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BlobStorageEventTypePtrOutput struct{ *pulumi.OutputState }

func (BlobStorageEventTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobStorageEventType)(nil)).Elem()
}

func (o BlobStorageEventTypePtrOutput) ToBlobStorageEventTypePtrOutput() BlobStorageEventTypePtrOutput {
	return o
}

func (o BlobStorageEventTypePtrOutput) ToBlobStorageEventTypePtrOutputWithContext(ctx context.Context) BlobStorageEventTypePtrOutput {
	return o
}

func (o BlobStorageEventTypePtrOutput) Elem() BlobStorageEventTypeOutput {
	return o.ApplyT(func(v *BlobStorageEventType) BlobStorageEventType {
		if v != nil {
			return *v
		}
		var ret BlobStorageEventType
		return ret
	}).(BlobStorageEventTypeOutput)
}

func (o BlobStorageEventTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BlobStorageEventTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BlobStorageEventType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BlobStorageEventTypeInput interface {
	pulumi.Input

	ToBlobStorageEventTypeOutput() BlobStorageEventTypeOutput
	ToBlobStorageEventTypeOutputWithContext(context.Context) BlobStorageEventTypeOutput
}

var blobStorageEventTypePtrType = reflect.TypeOf((**BlobStorageEventType)(nil)).Elem()

type BlobStorageEventTypePtrInput interface {
	pulumi.Input

	ToBlobStorageEventTypePtrOutput() BlobStorageEventTypePtrOutput
	ToBlobStorageEventTypePtrOutputWithContext(context.Context) BlobStorageEventTypePtrOutput
}

type blobStorageEventTypePtr string

func BlobStorageEventTypePtr(v string) BlobStorageEventTypePtrInput {
	return (*blobStorageEventTypePtr)(&v)
}

func (*blobStorageEventTypePtr) ElementType() reflect.Type {
	return blobStorageEventTypePtrType
}

func (in *blobStorageEventTypePtr) ToBlobStorageEventTypePtrOutput() BlobStorageEventTypePtrOutput {
	return pulumi.ToOutput(in).(BlobStorageEventTypePtrOutput)
}

func (in *blobStorageEventTypePtr) ToBlobStorageEventTypePtrOutputWithContext(ctx context.Context) BlobStorageEventTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BlobStorageEventTypePtrOutput)
}

type ClusterPrincipalRole string

const (
	ClusterPrincipalRoleAllDatabasesAdmin  = ClusterPrincipalRole("AllDatabasesAdmin")
	ClusterPrincipalRoleAllDatabasesViewer = ClusterPrincipalRole("AllDatabasesViewer")
)

func (ClusterPrincipalRole) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPrincipalRole)(nil)).Elem()
}

func (e ClusterPrincipalRole) ToClusterPrincipalRoleOutput() ClusterPrincipalRoleOutput {
	return pulumi.ToOutput(e).(ClusterPrincipalRoleOutput)
}

func (e ClusterPrincipalRole) ToClusterPrincipalRoleOutputWithContext(ctx context.Context) ClusterPrincipalRoleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterPrincipalRoleOutput)
}

func (e ClusterPrincipalRole) ToClusterPrincipalRolePtrOutput() ClusterPrincipalRolePtrOutput {
	return e.ToClusterPrincipalRolePtrOutputWithContext(context.Background())
}

func (e ClusterPrincipalRole) ToClusterPrincipalRolePtrOutputWithContext(ctx context.Context) ClusterPrincipalRolePtrOutput {
	return ClusterPrincipalRole(e).ToClusterPrincipalRoleOutputWithContext(ctx).ToClusterPrincipalRolePtrOutputWithContext(ctx)
}

func (e ClusterPrincipalRole) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterPrincipalRole) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterPrincipalRole) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterPrincipalRole) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterPrincipalRoleOutput struct{ *pulumi.OutputState }

func (ClusterPrincipalRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterPrincipalRole)(nil)).Elem()
}

func (o ClusterPrincipalRoleOutput) ToClusterPrincipalRoleOutput() ClusterPrincipalRoleOutput {
	return o
}

func (o ClusterPrincipalRoleOutput) ToClusterPrincipalRoleOutputWithContext(ctx context.Context) ClusterPrincipalRoleOutput {
	return o
}

func (o ClusterPrincipalRoleOutput) ToClusterPrincipalRolePtrOutput() ClusterPrincipalRolePtrOutput {
	return o.ToClusterPrincipalRolePtrOutputWithContext(context.Background())
}

func (o ClusterPrincipalRoleOutput) ToClusterPrincipalRolePtrOutputWithContext(ctx context.Context) ClusterPrincipalRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterPrincipalRole) *ClusterPrincipalRole {
		return &v
	}).(ClusterPrincipalRolePtrOutput)
}

func (o ClusterPrincipalRoleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterPrincipalRoleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterPrincipalRole) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterPrincipalRoleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterPrincipalRoleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterPrincipalRole) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterPrincipalRolePtrOutput struct{ *pulumi.OutputState }

func (ClusterPrincipalRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterPrincipalRole)(nil)).Elem()
}

func (o ClusterPrincipalRolePtrOutput) ToClusterPrincipalRolePtrOutput() ClusterPrincipalRolePtrOutput {
	return o
}

func (o ClusterPrincipalRolePtrOutput) ToClusterPrincipalRolePtrOutputWithContext(ctx context.Context) ClusterPrincipalRolePtrOutput {
	return o
}

func (o ClusterPrincipalRolePtrOutput) Elem() ClusterPrincipalRoleOutput {
	return o.ApplyT(func(v *ClusterPrincipalRole) ClusterPrincipalRole {
		if v != nil {
			return *v
		}
		var ret ClusterPrincipalRole
		return ret
	}).(ClusterPrincipalRoleOutput)
}

func (o ClusterPrincipalRolePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterPrincipalRolePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterPrincipalRole) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClusterPrincipalRoleInput interface {
	pulumi.Input

	ToClusterPrincipalRoleOutput() ClusterPrincipalRoleOutput
	ToClusterPrincipalRoleOutputWithContext(context.Context) ClusterPrincipalRoleOutput
}

var clusterPrincipalRolePtrType = reflect.TypeOf((**ClusterPrincipalRole)(nil)).Elem()

type ClusterPrincipalRolePtrInput interface {
	pulumi.Input

	ToClusterPrincipalRolePtrOutput() ClusterPrincipalRolePtrOutput
	ToClusterPrincipalRolePtrOutputWithContext(context.Context) ClusterPrincipalRolePtrOutput
}

type clusterPrincipalRolePtr string

func ClusterPrincipalRolePtr(v string) ClusterPrincipalRolePtrInput {
	return (*clusterPrincipalRolePtr)(&v)
}

func (*clusterPrincipalRolePtr) ElementType() reflect.Type {
	return clusterPrincipalRolePtrType
}

func (in *clusterPrincipalRolePtr) ToClusterPrincipalRolePtrOutput() ClusterPrincipalRolePtrOutput {
	return pulumi.ToOutput(in).(ClusterPrincipalRolePtrOutput)
}

func (in *clusterPrincipalRolePtr) ToClusterPrincipalRolePtrOutputWithContext(ctx context.Context) ClusterPrincipalRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterPrincipalRolePtrOutput)
}

type Compression string

const (
	CompressionNone = Compression("None")
	CompressionGZip = Compression("GZip")
)

func (Compression) ElementType() reflect.Type {
	return reflect.TypeOf((*Compression)(nil)).Elem()
}

func (e Compression) ToCompressionOutput() CompressionOutput {
	return pulumi.ToOutput(e).(CompressionOutput)
}

func (e Compression) ToCompressionOutputWithContext(ctx context.Context) CompressionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CompressionOutput)
}

func (e Compression) ToCompressionPtrOutput() CompressionPtrOutput {
	return e.ToCompressionPtrOutputWithContext(context.Background())
}

func (e Compression) ToCompressionPtrOutputWithContext(ctx context.Context) CompressionPtrOutput {
	return Compression(e).ToCompressionOutputWithContext(ctx).ToCompressionPtrOutputWithContext(ctx)
}

func (e Compression) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Compression) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Compression) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Compression) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CompressionOutput struct{ *pulumi.OutputState }

func (CompressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Compression)(nil)).Elem()
}

func (o CompressionOutput) ToCompressionOutput() CompressionOutput {
	return o
}

func (o CompressionOutput) ToCompressionOutputWithContext(ctx context.Context) CompressionOutput {
	return o
}

func (o CompressionOutput) ToCompressionPtrOutput() CompressionPtrOutput {
	return o.ToCompressionPtrOutputWithContext(context.Background())
}

func (o CompressionOutput) ToCompressionPtrOutputWithContext(ctx context.Context) CompressionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Compression) *Compression {
		return &v
	}).(CompressionPtrOutput)
}

func (o CompressionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CompressionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Compression) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CompressionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CompressionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Compression) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CompressionPtrOutput struct{ *pulumi.OutputState }

func (CompressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Compression)(nil)).Elem()
}

func (o CompressionPtrOutput) ToCompressionPtrOutput() CompressionPtrOutput {
	return o
}

func (o CompressionPtrOutput) ToCompressionPtrOutputWithContext(ctx context.Context) CompressionPtrOutput {
	return o
}

func (o CompressionPtrOutput) Elem() CompressionOutput {
	return o.ApplyT(func(v *Compression) Compression {
		if v != nil {
			return *v
		}
		var ret Compression
		return ret
	}).(CompressionOutput)
}

func (o CompressionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CompressionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Compression) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CompressionInput interface {
	pulumi.Input

	ToCompressionOutput() CompressionOutput
	ToCompressionOutputWithContext(context.Context) CompressionOutput
}

var compressionPtrType = reflect.TypeOf((**Compression)(nil)).Elem()

type CompressionPtrInput interface {
	pulumi.Input

	ToCompressionPtrOutput() CompressionPtrOutput
	ToCompressionPtrOutputWithContext(context.Context) CompressionPtrOutput
}

type compressionPtr string

func CompressionPtr(v string) CompressionPtrInput {
	return (*compressionPtr)(&v)
}

func (*compressionPtr) ElementType() reflect.Type {
	return compressionPtrType
}

func (in *compressionPtr) ToCompressionPtrOutput() CompressionPtrOutput {
	return pulumi.ToOutput(in).(CompressionPtrOutput)
}

func (in *compressionPtr) ToCompressionPtrOutputWithContext(ctx context.Context) CompressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CompressionPtrOutput)
}

type DataConnectionKind string

const (
	DataConnectionKindEventHub  = DataConnectionKind("EventHub")
	DataConnectionKindEventGrid = DataConnectionKind("EventGrid")
	DataConnectionKindIotHub    = DataConnectionKind("IotHub")
)

func (DataConnectionKind) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectionKind)(nil)).Elem()
}

func (e DataConnectionKind) ToDataConnectionKindOutput() DataConnectionKindOutput {
	return pulumi.ToOutput(e).(DataConnectionKindOutput)
}

func (e DataConnectionKind) ToDataConnectionKindOutputWithContext(ctx context.Context) DataConnectionKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataConnectionKindOutput)
}

func (e DataConnectionKind) ToDataConnectionKindPtrOutput() DataConnectionKindPtrOutput {
	return e.ToDataConnectionKindPtrOutputWithContext(context.Background())
}

func (e DataConnectionKind) ToDataConnectionKindPtrOutputWithContext(ctx context.Context) DataConnectionKindPtrOutput {
	return DataConnectionKind(e).ToDataConnectionKindOutputWithContext(ctx).ToDataConnectionKindPtrOutputWithContext(ctx)
}

func (e DataConnectionKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataConnectionKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataConnectionKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataConnectionKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataConnectionKindOutput struct{ *pulumi.OutputState }

func (DataConnectionKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectionKind)(nil)).Elem()
}

func (o DataConnectionKindOutput) ToDataConnectionKindOutput() DataConnectionKindOutput {
	return o
}

func (o DataConnectionKindOutput) ToDataConnectionKindOutputWithContext(ctx context.Context) DataConnectionKindOutput {
	return o
}

func (o DataConnectionKindOutput) ToDataConnectionKindPtrOutput() DataConnectionKindPtrOutput {
	return o.ToDataConnectionKindPtrOutputWithContext(context.Background())
}

func (o DataConnectionKindOutput) ToDataConnectionKindPtrOutputWithContext(ctx context.Context) DataConnectionKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataConnectionKind) *DataConnectionKind {
		return &v
	}).(DataConnectionKindPtrOutput)
}

func (o DataConnectionKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataConnectionKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataConnectionKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataConnectionKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataConnectionKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataConnectionKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataConnectionKindPtrOutput struct{ *pulumi.OutputState }

func (DataConnectionKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectionKind)(nil)).Elem()
}

func (o DataConnectionKindPtrOutput) ToDataConnectionKindPtrOutput() DataConnectionKindPtrOutput {
	return o
}

func (o DataConnectionKindPtrOutput) ToDataConnectionKindPtrOutputWithContext(ctx context.Context) DataConnectionKindPtrOutput {
	return o
}

func (o DataConnectionKindPtrOutput) Elem() DataConnectionKindOutput {
	return o.ApplyT(func(v *DataConnectionKind) DataConnectionKind {
		if v != nil {
			return *v
		}
		var ret DataConnectionKind
		return ret
	}).(DataConnectionKindOutput)
}

func (o DataConnectionKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataConnectionKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataConnectionKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataConnectionKindInput interface {
	pulumi.Input

	ToDataConnectionKindOutput() DataConnectionKindOutput
	ToDataConnectionKindOutputWithContext(context.Context) DataConnectionKindOutput
}

var dataConnectionKindPtrType = reflect.TypeOf((**DataConnectionKind)(nil)).Elem()

type DataConnectionKindPtrInput interface {
	pulumi.Input

	ToDataConnectionKindPtrOutput() DataConnectionKindPtrOutput
	ToDataConnectionKindPtrOutputWithContext(context.Context) DataConnectionKindPtrOutput
}

type dataConnectionKindPtr string

func DataConnectionKindPtr(v string) DataConnectionKindPtrInput {
	return (*dataConnectionKindPtr)(&v)
}

func (*dataConnectionKindPtr) ElementType() reflect.Type {
	return dataConnectionKindPtrType
}

func (in *dataConnectionKindPtr) ToDataConnectionKindPtrOutput() DataConnectionKindPtrOutput {
	return pulumi.ToOutput(in).(DataConnectionKindPtrOutput)
}

func (in *dataConnectionKindPtr) ToDataConnectionKindPtrOutputWithContext(ctx context.Context) DataConnectionKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataConnectionKindPtrOutput)
}

type DatabasePrincipalRole string

const (
	DatabasePrincipalRoleAdmin              = DatabasePrincipalRole("Admin")
	DatabasePrincipalRoleIngestor           = DatabasePrincipalRole("Ingestor")
	DatabasePrincipalRoleMonitor            = DatabasePrincipalRole("Monitor")
	DatabasePrincipalRoleUser               = DatabasePrincipalRole("User")
	DatabasePrincipalRoleUnrestrictedViewer = DatabasePrincipalRole("UnrestrictedViewer")
	DatabasePrincipalRoleViewer             = DatabasePrincipalRole("Viewer")
)

func (DatabasePrincipalRole) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipalRole)(nil)).Elem()
}

func (e DatabasePrincipalRole) ToDatabasePrincipalRoleOutput() DatabasePrincipalRoleOutput {
	return pulumi.ToOutput(e).(DatabasePrincipalRoleOutput)
}

func (e DatabasePrincipalRole) ToDatabasePrincipalRoleOutputWithContext(ctx context.Context) DatabasePrincipalRoleOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DatabasePrincipalRoleOutput)
}

func (e DatabasePrincipalRole) ToDatabasePrincipalRolePtrOutput() DatabasePrincipalRolePtrOutput {
	return e.ToDatabasePrincipalRolePtrOutputWithContext(context.Background())
}

func (e DatabasePrincipalRole) ToDatabasePrincipalRolePtrOutputWithContext(ctx context.Context) DatabasePrincipalRolePtrOutput {
	return DatabasePrincipalRole(e).ToDatabasePrincipalRoleOutputWithContext(ctx).ToDatabasePrincipalRolePtrOutputWithContext(ctx)
}

func (e DatabasePrincipalRole) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabasePrincipalRole) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DatabasePrincipalRole) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DatabasePrincipalRole) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DatabasePrincipalRoleOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipalRole)(nil)).Elem()
}

func (o DatabasePrincipalRoleOutput) ToDatabasePrincipalRoleOutput() DatabasePrincipalRoleOutput {
	return o
}

func (o DatabasePrincipalRoleOutput) ToDatabasePrincipalRoleOutputWithContext(ctx context.Context) DatabasePrincipalRoleOutput {
	return o
}

func (o DatabasePrincipalRoleOutput) ToDatabasePrincipalRolePtrOutput() DatabasePrincipalRolePtrOutput {
	return o.ToDatabasePrincipalRolePtrOutputWithContext(context.Background())
}

func (o DatabasePrincipalRoleOutput) ToDatabasePrincipalRolePtrOutputWithContext(ctx context.Context) DatabasePrincipalRolePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DatabasePrincipalRole) *DatabasePrincipalRole {
		return &v
	}).(DatabasePrincipalRolePtrOutput)
}

func (o DatabasePrincipalRoleOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DatabasePrincipalRoleOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabasePrincipalRole) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DatabasePrincipalRoleOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabasePrincipalRoleOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DatabasePrincipalRole) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DatabasePrincipalRolePtrOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabasePrincipalRole)(nil)).Elem()
}

func (o DatabasePrincipalRolePtrOutput) ToDatabasePrincipalRolePtrOutput() DatabasePrincipalRolePtrOutput {
	return o
}

func (o DatabasePrincipalRolePtrOutput) ToDatabasePrincipalRolePtrOutputWithContext(ctx context.Context) DatabasePrincipalRolePtrOutput {
	return o
}

func (o DatabasePrincipalRolePtrOutput) Elem() DatabasePrincipalRoleOutput {
	return o.ApplyT(func(v *DatabasePrincipalRole) DatabasePrincipalRole {
		if v != nil {
			return *v
		}
		var ret DatabasePrincipalRole
		return ret
	}).(DatabasePrincipalRoleOutput)
}

func (o DatabasePrincipalRolePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DatabasePrincipalRolePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DatabasePrincipalRole) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DatabasePrincipalRoleInput interface {
	pulumi.Input

	ToDatabasePrincipalRoleOutput() DatabasePrincipalRoleOutput
	ToDatabasePrincipalRoleOutputWithContext(context.Context) DatabasePrincipalRoleOutput
}

var databasePrincipalRolePtrType = reflect.TypeOf((**DatabasePrincipalRole)(nil)).Elem()

type DatabasePrincipalRolePtrInput interface {
	pulumi.Input

	ToDatabasePrincipalRolePtrOutput() DatabasePrincipalRolePtrOutput
	ToDatabasePrincipalRolePtrOutputWithContext(context.Context) DatabasePrincipalRolePtrOutput
}

type databasePrincipalRolePtr string

func DatabasePrincipalRolePtr(v string) DatabasePrincipalRolePtrInput {
	return (*databasePrincipalRolePtr)(&v)
}

func (*databasePrincipalRolePtr) ElementType() reflect.Type {
	return databasePrincipalRolePtrType
}

func (in *databasePrincipalRolePtr) ToDatabasePrincipalRolePtrOutput() DatabasePrincipalRolePtrOutput {
	return pulumi.ToOutput(in).(DatabasePrincipalRolePtrOutput)
}

func (in *databasePrincipalRolePtr) ToDatabasePrincipalRolePtrOutputWithContext(ctx context.Context) DatabasePrincipalRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DatabasePrincipalRolePtrOutput)
}

type DefaultPrincipalsModificationKind string

const (
	DefaultPrincipalsModificationKindUnion   = DefaultPrincipalsModificationKind("Union")
	DefaultPrincipalsModificationKindReplace = DefaultPrincipalsModificationKind("Replace")
	DefaultPrincipalsModificationKindNone    = DefaultPrincipalsModificationKind("None")
)

func (DefaultPrincipalsModificationKind) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultPrincipalsModificationKind)(nil)).Elem()
}

func (e DefaultPrincipalsModificationKind) ToDefaultPrincipalsModificationKindOutput() DefaultPrincipalsModificationKindOutput {
	return pulumi.ToOutput(e).(DefaultPrincipalsModificationKindOutput)
}

func (e DefaultPrincipalsModificationKind) ToDefaultPrincipalsModificationKindOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DefaultPrincipalsModificationKindOutput)
}

func (e DefaultPrincipalsModificationKind) ToDefaultPrincipalsModificationKindPtrOutput() DefaultPrincipalsModificationKindPtrOutput {
	return e.ToDefaultPrincipalsModificationKindPtrOutputWithContext(context.Background())
}

func (e DefaultPrincipalsModificationKind) ToDefaultPrincipalsModificationKindPtrOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindPtrOutput {
	return DefaultPrincipalsModificationKind(e).ToDefaultPrincipalsModificationKindOutputWithContext(ctx).ToDefaultPrincipalsModificationKindPtrOutputWithContext(ctx)
}

func (e DefaultPrincipalsModificationKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultPrincipalsModificationKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultPrincipalsModificationKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefaultPrincipalsModificationKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DefaultPrincipalsModificationKindOutput struct{ *pulumi.OutputState }

func (DefaultPrincipalsModificationKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultPrincipalsModificationKind)(nil)).Elem()
}

func (o DefaultPrincipalsModificationKindOutput) ToDefaultPrincipalsModificationKindOutput() DefaultPrincipalsModificationKindOutput {
	return o
}

func (o DefaultPrincipalsModificationKindOutput) ToDefaultPrincipalsModificationKindOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindOutput {
	return o
}

func (o DefaultPrincipalsModificationKindOutput) ToDefaultPrincipalsModificationKindPtrOutput() DefaultPrincipalsModificationKindPtrOutput {
	return o.ToDefaultPrincipalsModificationKindPtrOutputWithContext(context.Background())
}

func (o DefaultPrincipalsModificationKindOutput) ToDefaultPrincipalsModificationKindPtrOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultPrincipalsModificationKind) *DefaultPrincipalsModificationKind {
		return &v
	}).(DefaultPrincipalsModificationKindPtrOutput)
}

func (o DefaultPrincipalsModificationKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DefaultPrincipalsModificationKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultPrincipalsModificationKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DefaultPrincipalsModificationKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultPrincipalsModificationKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultPrincipalsModificationKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DefaultPrincipalsModificationKindPtrOutput struct{ *pulumi.OutputState }

func (DefaultPrincipalsModificationKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultPrincipalsModificationKind)(nil)).Elem()
}

func (o DefaultPrincipalsModificationKindPtrOutput) ToDefaultPrincipalsModificationKindPtrOutput() DefaultPrincipalsModificationKindPtrOutput {
	return o
}

func (o DefaultPrincipalsModificationKindPtrOutput) ToDefaultPrincipalsModificationKindPtrOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindPtrOutput {
	return o
}

func (o DefaultPrincipalsModificationKindPtrOutput) Elem() DefaultPrincipalsModificationKindOutput {
	return o.ApplyT(func(v *DefaultPrincipalsModificationKind) DefaultPrincipalsModificationKind {
		if v != nil {
			return *v
		}
		var ret DefaultPrincipalsModificationKind
		return ret
	}).(DefaultPrincipalsModificationKindOutput)
}

func (o DefaultPrincipalsModificationKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultPrincipalsModificationKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DefaultPrincipalsModificationKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DefaultPrincipalsModificationKindInput interface {
	pulumi.Input

	ToDefaultPrincipalsModificationKindOutput() DefaultPrincipalsModificationKindOutput
	ToDefaultPrincipalsModificationKindOutputWithContext(context.Context) DefaultPrincipalsModificationKindOutput
}

var defaultPrincipalsModificationKindPtrType = reflect.TypeOf((**DefaultPrincipalsModificationKind)(nil)).Elem()

type DefaultPrincipalsModificationKindPtrInput interface {
	pulumi.Input

	ToDefaultPrincipalsModificationKindPtrOutput() DefaultPrincipalsModificationKindPtrOutput
	ToDefaultPrincipalsModificationKindPtrOutputWithContext(context.Context) DefaultPrincipalsModificationKindPtrOutput
}

type defaultPrincipalsModificationKindPtr string

func DefaultPrincipalsModificationKindPtr(v string) DefaultPrincipalsModificationKindPtrInput {
	return (*defaultPrincipalsModificationKindPtr)(&v)
}

func (*defaultPrincipalsModificationKindPtr) ElementType() reflect.Type {
	return defaultPrincipalsModificationKindPtrType
}

func (in *defaultPrincipalsModificationKindPtr) ToDefaultPrincipalsModificationKindPtrOutput() DefaultPrincipalsModificationKindPtrOutput {
	return pulumi.ToOutput(in).(DefaultPrincipalsModificationKindPtrOutput)
}

func (in *defaultPrincipalsModificationKindPtr) ToDefaultPrincipalsModificationKindPtrOutputWithContext(ctx context.Context) DefaultPrincipalsModificationKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DefaultPrincipalsModificationKindPtrOutput)
}

type EngineType string

const (
	EngineTypeV2 = EngineType("V2")
	EngineTypeV3 = EngineType("V3")
)

func (EngineType) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineType)(nil)).Elem()
}

func (e EngineType) ToEngineTypeOutput() EngineTypeOutput {
	return pulumi.ToOutput(e).(EngineTypeOutput)
}

func (e EngineType) ToEngineTypeOutputWithContext(ctx context.Context) EngineTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EngineTypeOutput)
}

func (e EngineType) ToEngineTypePtrOutput() EngineTypePtrOutput {
	return e.ToEngineTypePtrOutputWithContext(context.Background())
}

func (e EngineType) ToEngineTypePtrOutputWithContext(ctx context.Context) EngineTypePtrOutput {
	return EngineType(e).ToEngineTypeOutputWithContext(ctx).ToEngineTypePtrOutputWithContext(ctx)
}

func (e EngineType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EngineType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EngineType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EngineType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EngineTypeOutput struct{ *pulumi.OutputState }

func (EngineTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EngineType)(nil)).Elem()
}

func (o EngineTypeOutput) ToEngineTypeOutput() EngineTypeOutput {
	return o
}

func (o EngineTypeOutput) ToEngineTypeOutputWithContext(ctx context.Context) EngineTypeOutput {
	return o
}

func (o EngineTypeOutput) ToEngineTypePtrOutput() EngineTypePtrOutput {
	return o.ToEngineTypePtrOutputWithContext(context.Background())
}

func (o EngineTypeOutput) ToEngineTypePtrOutputWithContext(ctx context.Context) EngineTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EngineType) *EngineType {
		return &v
	}).(EngineTypePtrOutput)
}

func (o EngineTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EngineTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EngineType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EngineTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EngineTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EngineType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EngineTypePtrOutput struct{ *pulumi.OutputState }

func (EngineTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EngineType)(nil)).Elem()
}

func (o EngineTypePtrOutput) ToEngineTypePtrOutput() EngineTypePtrOutput {
	return o
}

func (o EngineTypePtrOutput) ToEngineTypePtrOutputWithContext(ctx context.Context) EngineTypePtrOutput {
	return o
}

func (o EngineTypePtrOutput) Elem() EngineTypeOutput {
	return o.ApplyT(func(v *EngineType) EngineType {
		if v != nil {
			return *v
		}
		var ret EngineType
		return ret
	}).(EngineTypeOutput)
}

func (o EngineTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EngineTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EngineType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EngineTypeInput interface {
	pulumi.Input

	ToEngineTypeOutput() EngineTypeOutput
	ToEngineTypeOutputWithContext(context.Context) EngineTypeOutput
}

var engineTypePtrType = reflect.TypeOf((**EngineType)(nil)).Elem()

type EngineTypePtrInput interface {
	pulumi.Input

	ToEngineTypePtrOutput() EngineTypePtrOutput
	ToEngineTypePtrOutputWithContext(context.Context) EngineTypePtrOutput
}

type engineTypePtr string

func EngineTypePtr(v string) EngineTypePtrInput {
	return (*engineTypePtr)(&v)
}

func (*engineTypePtr) ElementType() reflect.Type {
	return engineTypePtrType
}

func (in *engineTypePtr) ToEngineTypePtrOutput() EngineTypePtrOutput {
	return pulumi.ToOutput(in).(EngineTypePtrOutput)
}

func (in *engineTypePtr) ToEngineTypePtrOutputWithContext(ctx context.Context) EngineTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EngineTypePtrOutput)
}

type EventGridDataFormat string

const (
	EventGridDataFormatMULTIJSON  = EventGridDataFormat("MULTIJSON")
	EventGridDataFormatJSON       = EventGridDataFormat("JSON")
	EventGridDataFormatCSV        = EventGridDataFormat("CSV")
	EventGridDataFormatTSV        = EventGridDataFormat("TSV")
	EventGridDataFormatSCSV       = EventGridDataFormat("SCSV")
	EventGridDataFormatSOHSV      = EventGridDataFormat("SOHSV")
	EventGridDataFormatPSV        = EventGridDataFormat("PSV")
	EventGridDataFormatTXT        = EventGridDataFormat("TXT")
	EventGridDataFormatRAW        = EventGridDataFormat("RAW")
	EventGridDataFormatSINGLEJSON = EventGridDataFormat("SINGLEJSON")
	EventGridDataFormatAVRO       = EventGridDataFormat("AVRO")
	EventGridDataFormatTSVE       = EventGridDataFormat("TSVE")
	EventGridDataFormatPARQUET    = EventGridDataFormat("PARQUET")
	EventGridDataFormatORC        = EventGridDataFormat("ORC")
	EventGridDataFormatAPACHEAVRO = EventGridDataFormat("APACHEAVRO")
	EventGridDataFormatW3CLOGFILE = EventGridDataFormat("W3CLOGFILE")
)

func (EventGridDataFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridDataFormat)(nil)).Elem()
}

func (e EventGridDataFormat) ToEventGridDataFormatOutput() EventGridDataFormatOutput {
	return pulumi.ToOutput(e).(EventGridDataFormatOutput)
}

func (e EventGridDataFormat) ToEventGridDataFormatOutputWithContext(ctx context.Context) EventGridDataFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventGridDataFormatOutput)
}

func (e EventGridDataFormat) ToEventGridDataFormatPtrOutput() EventGridDataFormatPtrOutput {
	return e.ToEventGridDataFormatPtrOutputWithContext(context.Background())
}

func (e EventGridDataFormat) ToEventGridDataFormatPtrOutputWithContext(ctx context.Context) EventGridDataFormatPtrOutput {
	return EventGridDataFormat(e).ToEventGridDataFormatOutputWithContext(ctx).ToEventGridDataFormatPtrOutputWithContext(ctx)
}

func (e EventGridDataFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventGridDataFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventGridDataFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventGridDataFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventGridDataFormatOutput struct{ *pulumi.OutputState }

func (EventGridDataFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridDataFormat)(nil)).Elem()
}

func (o EventGridDataFormatOutput) ToEventGridDataFormatOutput() EventGridDataFormatOutput {
	return o
}

func (o EventGridDataFormatOutput) ToEventGridDataFormatOutputWithContext(ctx context.Context) EventGridDataFormatOutput {
	return o
}

func (o EventGridDataFormatOutput) ToEventGridDataFormatPtrOutput() EventGridDataFormatPtrOutput {
	return o.ToEventGridDataFormatPtrOutputWithContext(context.Background())
}

func (o EventGridDataFormatOutput) ToEventGridDataFormatPtrOutputWithContext(ctx context.Context) EventGridDataFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventGridDataFormat) *EventGridDataFormat {
		return &v
	}).(EventGridDataFormatPtrOutput)
}

func (o EventGridDataFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventGridDataFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventGridDataFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventGridDataFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventGridDataFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventGridDataFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventGridDataFormatPtrOutput struct{ *pulumi.OutputState }

func (EventGridDataFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGridDataFormat)(nil)).Elem()
}

func (o EventGridDataFormatPtrOutput) ToEventGridDataFormatPtrOutput() EventGridDataFormatPtrOutput {
	return o
}

func (o EventGridDataFormatPtrOutput) ToEventGridDataFormatPtrOutputWithContext(ctx context.Context) EventGridDataFormatPtrOutput {
	return o
}

func (o EventGridDataFormatPtrOutput) Elem() EventGridDataFormatOutput {
	return o.ApplyT(func(v *EventGridDataFormat) EventGridDataFormat {
		if v != nil {
			return *v
		}
		var ret EventGridDataFormat
		return ret
	}).(EventGridDataFormatOutput)
}

func (o EventGridDataFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventGridDataFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventGridDataFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EventGridDataFormatInput interface {
	pulumi.Input

	ToEventGridDataFormatOutput() EventGridDataFormatOutput
	ToEventGridDataFormatOutputWithContext(context.Context) EventGridDataFormatOutput
}

var eventGridDataFormatPtrType = reflect.TypeOf((**EventGridDataFormat)(nil)).Elem()

type EventGridDataFormatPtrInput interface {
	pulumi.Input

	ToEventGridDataFormatPtrOutput() EventGridDataFormatPtrOutput
	ToEventGridDataFormatPtrOutputWithContext(context.Context) EventGridDataFormatPtrOutput
}

type eventGridDataFormatPtr string

func EventGridDataFormatPtr(v string) EventGridDataFormatPtrInput {
	return (*eventGridDataFormatPtr)(&v)
}

func (*eventGridDataFormatPtr) ElementType() reflect.Type {
	return eventGridDataFormatPtrType
}

func (in *eventGridDataFormatPtr) ToEventGridDataFormatPtrOutput() EventGridDataFormatPtrOutput {
	return pulumi.ToOutput(in).(EventGridDataFormatPtrOutput)
}

func (in *eventGridDataFormatPtr) ToEventGridDataFormatPtrOutputWithContext(ctx context.Context) EventGridDataFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventGridDataFormatPtrOutput)
}

type EventHubDataFormat string

const (
	EventHubDataFormatMULTIJSON  = EventHubDataFormat("MULTIJSON")
	EventHubDataFormatJSON       = EventHubDataFormat("JSON")
	EventHubDataFormatCSV        = EventHubDataFormat("CSV")
	EventHubDataFormatTSV        = EventHubDataFormat("TSV")
	EventHubDataFormatSCSV       = EventHubDataFormat("SCSV")
	EventHubDataFormatSOHSV      = EventHubDataFormat("SOHSV")
	EventHubDataFormatPSV        = EventHubDataFormat("PSV")
	EventHubDataFormatTXT        = EventHubDataFormat("TXT")
	EventHubDataFormatRAW        = EventHubDataFormat("RAW")
	EventHubDataFormatSINGLEJSON = EventHubDataFormat("SINGLEJSON")
	EventHubDataFormatAVRO       = EventHubDataFormat("AVRO")
	EventHubDataFormatTSVE       = EventHubDataFormat("TSVE")
	EventHubDataFormatPARQUET    = EventHubDataFormat("PARQUET")
	EventHubDataFormatORC        = EventHubDataFormat("ORC")
	EventHubDataFormatAPACHEAVRO = EventHubDataFormat("APACHEAVRO")
	EventHubDataFormatW3CLOGFILE = EventHubDataFormat("W3CLOGFILE")
)

func (EventHubDataFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubDataFormat)(nil)).Elem()
}

func (e EventHubDataFormat) ToEventHubDataFormatOutput() EventHubDataFormatOutput {
	return pulumi.ToOutput(e).(EventHubDataFormatOutput)
}

func (e EventHubDataFormat) ToEventHubDataFormatOutputWithContext(ctx context.Context) EventHubDataFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventHubDataFormatOutput)
}

func (e EventHubDataFormat) ToEventHubDataFormatPtrOutput() EventHubDataFormatPtrOutput {
	return e.ToEventHubDataFormatPtrOutputWithContext(context.Background())
}

func (e EventHubDataFormat) ToEventHubDataFormatPtrOutputWithContext(ctx context.Context) EventHubDataFormatPtrOutput {
	return EventHubDataFormat(e).ToEventHubDataFormatOutputWithContext(ctx).ToEventHubDataFormatPtrOutputWithContext(ctx)
}

func (e EventHubDataFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventHubDataFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventHubDataFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventHubDataFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventHubDataFormatOutput struct{ *pulumi.OutputState }

func (EventHubDataFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubDataFormat)(nil)).Elem()
}

func (o EventHubDataFormatOutput) ToEventHubDataFormatOutput() EventHubDataFormatOutput {
	return o
}

func (o EventHubDataFormatOutput) ToEventHubDataFormatOutputWithContext(ctx context.Context) EventHubDataFormatOutput {
	return o
}

func (o EventHubDataFormatOutput) ToEventHubDataFormatPtrOutput() EventHubDataFormatPtrOutput {
	return o.ToEventHubDataFormatPtrOutputWithContext(context.Background())
}

func (o EventHubDataFormatOutput) ToEventHubDataFormatPtrOutputWithContext(ctx context.Context) EventHubDataFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventHubDataFormat) *EventHubDataFormat {
		return &v
	}).(EventHubDataFormatPtrOutput)
}

func (o EventHubDataFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventHubDataFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventHubDataFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventHubDataFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventHubDataFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventHubDataFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventHubDataFormatPtrOutput struct{ *pulumi.OutputState }

func (EventHubDataFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventHubDataFormat)(nil)).Elem()
}

func (o EventHubDataFormatPtrOutput) ToEventHubDataFormatPtrOutput() EventHubDataFormatPtrOutput {
	return o
}

func (o EventHubDataFormatPtrOutput) ToEventHubDataFormatPtrOutputWithContext(ctx context.Context) EventHubDataFormatPtrOutput {
	return o
}

func (o EventHubDataFormatPtrOutput) Elem() EventHubDataFormatOutput {
	return o.ApplyT(func(v *EventHubDataFormat) EventHubDataFormat {
		if v != nil {
			return *v
		}
		var ret EventHubDataFormat
		return ret
	}).(EventHubDataFormatOutput)
}

func (o EventHubDataFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventHubDataFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventHubDataFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EventHubDataFormatInput interface {
	pulumi.Input

	ToEventHubDataFormatOutput() EventHubDataFormatOutput
	ToEventHubDataFormatOutputWithContext(context.Context) EventHubDataFormatOutput
}

var eventHubDataFormatPtrType = reflect.TypeOf((**EventHubDataFormat)(nil)).Elem()

type EventHubDataFormatPtrInput interface {
	pulumi.Input

	ToEventHubDataFormatPtrOutput() EventHubDataFormatPtrOutput
	ToEventHubDataFormatPtrOutputWithContext(context.Context) EventHubDataFormatPtrOutput
}

type eventHubDataFormatPtr string

func EventHubDataFormatPtr(v string) EventHubDataFormatPtrInput {
	return (*eventHubDataFormatPtr)(&v)
}

func (*eventHubDataFormatPtr) ElementType() reflect.Type {
	return eventHubDataFormatPtrType
}

func (in *eventHubDataFormatPtr) ToEventHubDataFormatPtrOutput() EventHubDataFormatPtrOutput {
	return pulumi.ToOutput(in).(EventHubDataFormatPtrOutput)
}

func (in *eventHubDataFormatPtr) ToEventHubDataFormatPtrOutputWithContext(ctx context.Context) EventHubDataFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventHubDataFormatPtrOutput)
}

type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned, UserAssigned")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (e IdentityType) ToIdentityTypeOutput() IdentityTypeOutput {
	return pulumi.ToOutput(e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return e.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (e IdentityType) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return IdentityType(e).ToIdentityTypeOutputWithContext(ctx).ToIdentityTypePtrOutputWithContext(ctx)
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityTypeOutput struct{ *pulumi.OutputState }

func (IdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (o IdentityTypeOutput) ToIdentityTypeOutput() IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityType) *IdentityType {
		return &v
	}).(IdentityTypePtrOutput)
}

func (o IdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityType)(nil)).Elem()
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) Elem() IdentityTypeOutput {
	return o.ApplyT(func(v *IdentityType) IdentityType {
		if v != nil {
			return *v
		}
		var ret IdentityType
		return ret
	}).(IdentityTypeOutput)
}

func (o IdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IdentityTypeInput interface {
	pulumi.Input

	ToIdentityTypeOutput() IdentityTypeOutput
	ToIdentityTypeOutputWithContext(context.Context) IdentityTypeOutput
}

var identityTypePtrType = reflect.TypeOf((**IdentityType)(nil)).Elem()

type IdentityTypePtrInput interface {
	pulumi.Input

	ToIdentityTypePtrOutput() IdentityTypePtrOutput
	ToIdentityTypePtrOutputWithContext(context.Context) IdentityTypePtrOutput
}

type identityTypePtr string

func IdentityTypePtr(v string) IdentityTypePtrInput {
	return (*identityTypePtr)(&v)
}

func (*identityTypePtr) ElementType() reflect.Type {
	return identityTypePtrType
}

func (in *identityTypePtr) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityTypePtrOutput)
}

type IotHubDataFormat string

const (
	IotHubDataFormatMULTIJSON  = IotHubDataFormat("MULTIJSON")
	IotHubDataFormatJSON       = IotHubDataFormat("JSON")
	IotHubDataFormatCSV        = IotHubDataFormat("CSV")
	IotHubDataFormatTSV        = IotHubDataFormat("TSV")
	IotHubDataFormatSCSV       = IotHubDataFormat("SCSV")
	IotHubDataFormatSOHSV      = IotHubDataFormat("SOHSV")
	IotHubDataFormatPSV        = IotHubDataFormat("PSV")
	IotHubDataFormatTXT        = IotHubDataFormat("TXT")
	IotHubDataFormatRAW        = IotHubDataFormat("RAW")
	IotHubDataFormatSINGLEJSON = IotHubDataFormat("SINGLEJSON")
	IotHubDataFormatAVRO       = IotHubDataFormat("AVRO")
	IotHubDataFormatTSVE       = IotHubDataFormat("TSVE")
	IotHubDataFormatPARQUET    = IotHubDataFormat("PARQUET")
	IotHubDataFormatORC        = IotHubDataFormat("ORC")
	IotHubDataFormatAPACHEAVRO = IotHubDataFormat("APACHEAVRO")
	IotHubDataFormatW3CLOGFILE = IotHubDataFormat("W3CLOGFILE")
)

func (IotHubDataFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDataFormat)(nil)).Elem()
}

func (e IotHubDataFormat) ToIotHubDataFormatOutput() IotHubDataFormatOutput {
	return pulumi.ToOutput(e).(IotHubDataFormatOutput)
}

func (e IotHubDataFormat) ToIotHubDataFormatOutputWithContext(ctx context.Context) IotHubDataFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IotHubDataFormatOutput)
}

func (e IotHubDataFormat) ToIotHubDataFormatPtrOutput() IotHubDataFormatPtrOutput {
	return e.ToIotHubDataFormatPtrOutputWithContext(context.Background())
}

func (e IotHubDataFormat) ToIotHubDataFormatPtrOutputWithContext(ctx context.Context) IotHubDataFormatPtrOutput {
	return IotHubDataFormat(e).ToIotHubDataFormatOutputWithContext(ctx).ToIotHubDataFormatPtrOutputWithContext(ctx)
}

func (e IotHubDataFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IotHubDataFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IotHubDataFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IotHubDataFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IotHubDataFormatOutput struct{ *pulumi.OutputState }

func (IotHubDataFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubDataFormat)(nil)).Elem()
}

func (o IotHubDataFormatOutput) ToIotHubDataFormatOutput() IotHubDataFormatOutput {
	return o
}

func (o IotHubDataFormatOutput) ToIotHubDataFormatOutputWithContext(ctx context.Context) IotHubDataFormatOutput {
	return o
}

func (o IotHubDataFormatOutput) ToIotHubDataFormatPtrOutput() IotHubDataFormatPtrOutput {
	return o.ToIotHubDataFormatPtrOutputWithContext(context.Background())
}

func (o IotHubDataFormatOutput) ToIotHubDataFormatPtrOutputWithContext(ctx context.Context) IotHubDataFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotHubDataFormat) *IotHubDataFormat {
		return &v
	}).(IotHubDataFormatPtrOutput)
}

func (o IotHubDataFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IotHubDataFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IotHubDataFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IotHubDataFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IotHubDataFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IotHubDataFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IotHubDataFormatPtrOutput struct{ *pulumi.OutputState }

func (IotHubDataFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubDataFormat)(nil)).Elem()
}

func (o IotHubDataFormatPtrOutput) ToIotHubDataFormatPtrOutput() IotHubDataFormatPtrOutput {
	return o
}

func (o IotHubDataFormatPtrOutput) ToIotHubDataFormatPtrOutputWithContext(ctx context.Context) IotHubDataFormatPtrOutput {
	return o
}

func (o IotHubDataFormatPtrOutput) Elem() IotHubDataFormatOutput {
	return o.ApplyT(func(v *IotHubDataFormat) IotHubDataFormat {
		if v != nil {
			return *v
		}
		var ret IotHubDataFormat
		return ret
	}).(IotHubDataFormatOutput)
}

func (o IotHubDataFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IotHubDataFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IotHubDataFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IotHubDataFormatInput interface {
	pulumi.Input

	ToIotHubDataFormatOutput() IotHubDataFormatOutput
	ToIotHubDataFormatOutputWithContext(context.Context) IotHubDataFormatOutput
}

var iotHubDataFormatPtrType = reflect.TypeOf((**IotHubDataFormat)(nil)).Elem()

type IotHubDataFormatPtrInput interface {
	pulumi.Input

	ToIotHubDataFormatPtrOutput() IotHubDataFormatPtrOutput
	ToIotHubDataFormatPtrOutputWithContext(context.Context) IotHubDataFormatPtrOutput
}

type iotHubDataFormatPtr string

func IotHubDataFormatPtr(v string) IotHubDataFormatPtrInput {
	return (*iotHubDataFormatPtr)(&v)
}

func (*iotHubDataFormatPtr) ElementType() reflect.Type {
	return iotHubDataFormatPtrType
}

func (in *iotHubDataFormatPtr) ToIotHubDataFormatPtrOutput() IotHubDataFormatPtrOutput {
	return pulumi.ToOutput(in).(IotHubDataFormatPtrOutput)
}

func (in *iotHubDataFormatPtr) ToIotHubDataFormatPtrOutputWithContext(ctx context.Context) IotHubDataFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IotHubDataFormatPtrOutput)
}

type Kind string

const (
	KindReadWrite         = Kind("ReadWrite")
	KindReadOnlyFollowing = Kind("ReadOnlyFollowing")
)

func (Kind) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (e Kind) ToKindOutput() KindOutput {
	return pulumi.ToOutput(e).(KindOutput)
}

func (e Kind) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KindOutput)
}

func (e Kind) ToKindPtrOutput() KindPtrOutput {
	return e.ToKindPtrOutputWithContext(context.Background())
}

func (e Kind) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return Kind(e).ToKindOutputWithContext(ctx).ToKindPtrOutputWithContext(ctx)
}

func (e Kind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Kind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Kind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KindOutput struct{ *pulumi.OutputState }

func (KindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Kind)(nil)).Elem()
}

func (o KindOutput) ToKindOutput() KindOutput {
	return o
}

func (o KindOutput) ToKindOutputWithContext(ctx context.Context) KindOutput {
	return o
}

func (o KindOutput) ToKindPtrOutput() KindPtrOutput {
	return o.ToKindPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Kind) *Kind {
		return &v
	}).(KindPtrOutput)
}

func (o KindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Kind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KindPtrOutput struct{ *pulumi.OutputState }

func (KindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kind)(nil)).Elem()
}

func (o KindPtrOutput) ToKindPtrOutput() KindPtrOutput {
	return o
}

func (o KindPtrOutput) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return o
}

func (o KindPtrOutput) Elem() KindOutput {
	return o.ApplyT(func(v *Kind) Kind {
		if v != nil {
			return *v
		}
		var ret Kind
		return ret
	}).(KindOutput)
}

func (o KindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Kind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KindInput interface {
	pulumi.Input

	ToKindOutput() KindOutput
	ToKindOutputWithContext(context.Context) KindOutput
}

var kindPtrType = reflect.TypeOf((**Kind)(nil)).Elem()

type KindPtrInput interface {
	pulumi.Input

	ToKindPtrOutput() KindPtrOutput
	ToKindPtrOutputWithContext(context.Context) KindPtrOutput
}

type kindPtr string

func KindPtr(v string) KindPtrInput {
	return (*kindPtr)(&v)
}

func (*kindPtr) ElementType() reflect.Type {
	return kindPtrType
}

func (in *kindPtr) ToKindPtrOutput() KindPtrOutput {
	return pulumi.ToOutput(in).(KindPtrOutput)
}

func (in *kindPtr) ToKindPtrOutputWithContext(ctx context.Context) KindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KindPtrOutput)
}

type PrincipalType string

const (
	PrincipalTypeApp   = PrincipalType("App")
	PrincipalTypeGroup = PrincipalType("Group")
	PrincipalTypeUser  = PrincipalType("User")
)

func (PrincipalType) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalType)(nil)).Elem()
}

func (e PrincipalType) ToPrincipalTypeOutput() PrincipalTypeOutput {
	return pulumi.ToOutput(e).(PrincipalTypeOutput)
}

func (e PrincipalType) ToPrincipalTypeOutputWithContext(ctx context.Context) PrincipalTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrincipalTypeOutput)
}

func (e PrincipalType) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return e.ToPrincipalTypePtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return PrincipalType(e).ToPrincipalTypeOutputWithContext(ctx).ToPrincipalTypePtrOutputWithContext(ctx)
}

func (e PrincipalType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrincipalTypeOutput struct{ *pulumi.OutputState }

func (PrincipalTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalType)(nil)).Elem()
}

func (o PrincipalTypeOutput) ToPrincipalTypeOutput() PrincipalTypeOutput {
	return o
}

func (o PrincipalTypeOutput) ToPrincipalTypeOutputWithContext(ctx context.Context) PrincipalTypeOutput {
	return o
}

func (o PrincipalTypeOutput) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return o.ToPrincipalTypePtrOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrincipalType) *PrincipalType {
		return &v
	}).(PrincipalTypePtrOutput)
}

func (o PrincipalTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrincipalType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrincipalTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrincipalType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrincipalTypePtrOutput struct{ *pulumi.OutputState }

func (PrincipalTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalType)(nil)).Elem()
}

func (o PrincipalTypePtrOutput) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return o
}

func (o PrincipalTypePtrOutput) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return o
}

func (o PrincipalTypePtrOutput) Elem() PrincipalTypeOutput {
	return o.ApplyT(func(v *PrincipalType) PrincipalType {
		if v != nil {
			return *v
		}
		var ret PrincipalType
		return ret
	}).(PrincipalTypeOutput)
}

func (o PrincipalTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrincipalTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrincipalType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrincipalTypeInput interface {
	pulumi.Input

	ToPrincipalTypeOutput() PrincipalTypeOutput
	ToPrincipalTypeOutputWithContext(context.Context) PrincipalTypeOutput
}

var principalTypePtrType = reflect.TypeOf((**PrincipalType)(nil)).Elem()

type PrincipalTypePtrInput interface {
	pulumi.Input

	ToPrincipalTypePtrOutput() PrincipalTypePtrOutput
	ToPrincipalTypePtrOutputWithContext(context.Context) PrincipalTypePtrOutput
}

type principalTypePtr string

func PrincipalTypePtr(v string) PrincipalTypePtrInput {
	return (*principalTypePtr)(&v)
}

func (*principalTypePtr) ElementType() reflect.Type {
	return principalTypePtrType
}

func (in *principalTypePtr) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return pulumi.ToOutput(in).(PrincipalTypePtrOutput)
}

func (in *principalTypePtr) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrincipalTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureSkuNameOutput{})
	pulumi.RegisterOutputType(AzureSkuNamePtrOutput{})
	pulumi.RegisterOutputType(AzureSkuTierOutput{})
	pulumi.RegisterOutputType(AzureSkuTierPtrOutput{})
	pulumi.RegisterOutputType(BlobStorageEventTypeOutput{})
	pulumi.RegisterOutputType(BlobStorageEventTypePtrOutput{})
	pulumi.RegisterOutputType(ClusterPrincipalRoleOutput{})
	pulumi.RegisterOutputType(ClusterPrincipalRolePtrOutput{})
	pulumi.RegisterOutputType(CompressionOutput{})
	pulumi.RegisterOutputType(CompressionPtrOutput{})
	pulumi.RegisterOutputType(DataConnectionKindOutput{})
	pulumi.RegisterOutputType(DataConnectionKindPtrOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalRoleOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalRolePtrOutput{})
	pulumi.RegisterOutputType(DefaultPrincipalsModificationKindOutput{})
	pulumi.RegisterOutputType(DefaultPrincipalsModificationKindPtrOutput{})
	pulumi.RegisterOutputType(EngineTypeOutput{})
	pulumi.RegisterOutputType(EngineTypePtrOutput{})
	pulumi.RegisterOutputType(EventGridDataFormatOutput{})
	pulumi.RegisterOutputType(EventGridDataFormatPtrOutput{})
	pulumi.RegisterOutputType(EventHubDataFormatOutput{})
	pulumi.RegisterOutputType(EventHubDataFormatPtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(IotHubDataFormatOutput{})
	pulumi.RegisterOutputType(IotHubDataFormatPtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
	pulumi.RegisterOutputType(PrincipalTypeOutput{})
	pulumi.RegisterOutputType(PrincipalTypePtrOutput{})
}
