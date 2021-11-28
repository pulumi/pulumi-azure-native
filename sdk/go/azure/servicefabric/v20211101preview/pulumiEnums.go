


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Access string

const (
	AccessAllow = Access("allow")
	AccessDeny  = Access("deny")
)

func (Access) ElementType() reflect.Type {
	return reflect.TypeOf((*Access)(nil)).Elem()
}

func (e Access) ToAccessOutput() AccessOutput {
	return pulumi.ToOutput(e).(AccessOutput)
}

func (e Access) ToAccessOutputWithContext(ctx context.Context) AccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessOutput)
}

func (e Access) ToAccessPtrOutput() AccessPtrOutput {
	return e.ToAccessPtrOutputWithContext(context.Background())
}

func (e Access) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return Access(e).ToAccessOutputWithContext(ctx).ToAccessPtrOutputWithContext(ctx)
}

func (e Access) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Access) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Access) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Access) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessOutput struct{ *pulumi.OutputState }

func (AccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Access)(nil)).Elem()
}

func (o AccessOutput) ToAccessOutput() AccessOutput {
	return o
}

func (o AccessOutput) ToAccessOutputWithContext(ctx context.Context) AccessOutput {
	return o
}

func (o AccessOutput) ToAccessPtrOutput() AccessPtrOutput {
	return o.ToAccessPtrOutputWithContext(context.Background())
}

func (o AccessOutput) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Access) *Access {
		return &v
	}).(AccessPtrOutput)
}

func (o AccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Access) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Access) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessPtrOutput struct{ *pulumi.OutputState }

func (AccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Access)(nil)).Elem()
}

func (o AccessPtrOutput) ToAccessPtrOutput() AccessPtrOutput {
	return o
}

func (o AccessPtrOutput) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return o
}

func (o AccessPtrOutput) Elem() AccessOutput {
	return o.ApplyT(func(v *Access) Access {
		if v != nil {
			return *v
		}
		var ret Access
		return ret
	}).(AccessOutput)
}

func (o AccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Access) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessInput interface {
	pulumi.Input

	ToAccessOutput() AccessOutput
	ToAccessOutputWithContext(context.Context) AccessOutput
}

var accessPtrType = reflect.TypeOf((**Access)(nil)).Elem()

type AccessPtrInput interface {
	pulumi.Input

	ToAccessPtrOutput() AccessPtrOutput
	ToAccessPtrOutputWithContext(context.Context) AccessPtrOutput
}

type accessPtr string

func AccessPtr(v string) AccessPtrInput {
	return (*accessPtr)(&v)
}

func (*accessPtr) ElementType() reflect.Type {
	return accessPtrType
}

func (in *accessPtr) ToAccessPtrOutput() AccessPtrOutput {
	return pulumi.ToOutput(in).(AccessPtrOutput)
}

func (in *accessPtr) ToAccessPtrOutputWithContext(ctx context.Context) AccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessPtrOutput)
}

type ClusterUpgradeCadence string

const (
	// Cluster upgrade starts immediately after a new version is rolled out. Recommended for Test/Dev clusters.
	ClusterUpgradeCadenceWave0 = ClusterUpgradeCadence("Wave0")
	// Cluster upgrade starts 7 days after a new version is rolled out. Recommended for Pre-prod clusters.
	ClusterUpgradeCadenceWave1 = ClusterUpgradeCadence("Wave1")
	// Cluster upgrade starts 14 days after a new version is rolled out. Recommended for Production clusters.
	ClusterUpgradeCadenceWave2 = ClusterUpgradeCadence("Wave2")
)

func (ClusterUpgradeCadence) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeCadence)(nil)).Elem()
}

func (e ClusterUpgradeCadence) ToClusterUpgradeCadenceOutput() ClusterUpgradeCadenceOutput {
	return pulumi.ToOutput(e).(ClusterUpgradeCadenceOutput)
}

func (e ClusterUpgradeCadence) ToClusterUpgradeCadenceOutputWithContext(ctx context.Context) ClusterUpgradeCadenceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterUpgradeCadenceOutput)
}

func (e ClusterUpgradeCadence) ToClusterUpgradeCadencePtrOutput() ClusterUpgradeCadencePtrOutput {
	return e.ToClusterUpgradeCadencePtrOutputWithContext(context.Background())
}

func (e ClusterUpgradeCadence) ToClusterUpgradeCadencePtrOutputWithContext(ctx context.Context) ClusterUpgradeCadencePtrOutput {
	return ClusterUpgradeCadence(e).ToClusterUpgradeCadenceOutputWithContext(ctx).ToClusterUpgradeCadencePtrOutputWithContext(ctx)
}

func (e ClusterUpgradeCadence) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterUpgradeCadence) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterUpgradeCadence) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterUpgradeCadence) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterUpgradeCadenceOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeCadenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeCadence)(nil)).Elem()
}

func (o ClusterUpgradeCadenceOutput) ToClusterUpgradeCadenceOutput() ClusterUpgradeCadenceOutput {
	return o
}

func (o ClusterUpgradeCadenceOutput) ToClusterUpgradeCadenceOutputWithContext(ctx context.Context) ClusterUpgradeCadenceOutput {
	return o
}

func (o ClusterUpgradeCadenceOutput) ToClusterUpgradeCadencePtrOutput() ClusterUpgradeCadencePtrOutput {
	return o.ToClusterUpgradeCadencePtrOutputWithContext(context.Background())
}

func (o ClusterUpgradeCadenceOutput) ToClusterUpgradeCadencePtrOutputWithContext(ctx context.Context) ClusterUpgradeCadencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterUpgradeCadence) *ClusterUpgradeCadence {
		return &v
	}).(ClusterUpgradeCadencePtrOutput)
}

func (o ClusterUpgradeCadenceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterUpgradeCadenceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterUpgradeCadence) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterUpgradeCadenceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterUpgradeCadenceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterUpgradeCadence) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterUpgradeCadencePtrOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeCadencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradeCadence)(nil)).Elem()
}

func (o ClusterUpgradeCadencePtrOutput) ToClusterUpgradeCadencePtrOutput() ClusterUpgradeCadencePtrOutput {
	return o
}

func (o ClusterUpgradeCadencePtrOutput) ToClusterUpgradeCadencePtrOutputWithContext(ctx context.Context) ClusterUpgradeCadencePtrOutput {
	return o
}

func (o ClusterUpgradeCadencePtrOutput) Elem() ClusterUpgradeCadenceOutput {
	return o.ApplyT(func(v *ClusterUpgradeCadence) ClusterUpgradeCadence {
		if v != nil {
			return *v
		}
		var ret ClusterUpgradeCadence
		return ret
	}).(ClusterUpgradeCadenceOutput)
}

func (o ClusterUpgradeCadencePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterUpgradeCadencePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterUpgradeCadence) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClusterUpgradeCadenceInput interface {
	pulumi.Input

	ToClusterUpgradeCadenceOutput() ClusterUpgradeCadenceOutput
	ToClusterUpgradeCadenceOutputWithContext(context.Context) ClusterUpgradeCadenceOutput
}

var clusterUpgradeCadencePtrType = reflect.TypeOf((**ClusterUpgradeCadence)(nil)).Elem()

type ClusterUpgradeCadencePtrInput interface {
	pulumi.Input

	ToClusterUpgradeCadencePtrOutput() ClusterUpgradeCadencePtrOutput
	ToClusterUpgradeCadencePtrOutputWithContext(context.Context) ClusterUpgradeCadencePtrOutput
}

type clusterUpgradeCadencePtr string

func ClusterUpgradeCadencePtr(v string) ClusterUpgradeCadencePtrInput {
	return (*clusterUpgradeCadencePtr)(&v)
}

func (*clusterUpgradeCadencePtr) ElementType() reflect.Type {
	return clusterUpgradeCadencePtrType
}

func (in *clusterUpgradeCadencePtr) ToClusterUpgradeCadencePtrOutput() ClusterUpgradeCadencePtrOutput {
	return pulumi.ToOutput(in).(ClusterUpgradeCadencePtrOutput)
}

func (in *clusterUpgradeCadencePtr) ToClusterUpgradeCadencePtrOutputWithContext(ctx context.Context) ClusterUpgradeCadencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterUpgradeCadencePtrOutput)
}

type ClusterUpgradeMode string

const (
	// The cluster will be automatically upgraded to the latest Service Fabric runtime version, **clusterUpgradeCadence** will determine when the upgrade starts after the new version becomes available.
	ClusterUpgradeModeAutomatic = ClusterUpgradeMode("Automatic")
	// The cluster will not be automatically upgraded to the latest Service Fabric runtime version. The cluster is upgraded by setting the **clusterCodeVersion** property in the cluster resource.
	ClusterUpgradeModeManual = ClusterUpgradeMode("Manual")
)

func (ClusterUpgradeMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeMode)(nil)).Elem()
}

func (e ClusterUpgradeMode) ToClusterUpgradeModeOutput() ClusterUpgradeModeOutput {
	return pulumi.ToOutput(e).(ClusterUpgradeModeOutput)
}

func (e ClusterUpgradeMode) ToClusterUpgradeModeOutputWithContext(ctx context.Context) ClusterUpgradeModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ClusterUpgradeModeOutput)
}

func (e ClusterUpgradeMode) ToClusterUpgradeModePtrOutput() ClusterUpgradeModePtrOutput {
	return e.ToClusterUpgradeModePtrOutputWithContext(context.Background())
}

func (e ClusterUpgradeMode) ToClusterUpgradeModePtrOutputWithContext(ctx context.Context) ClusterUpgradeModePtrOutput {
	return ClusterUpgradeMode(e).ToClusterUpgradeModeOutputWithContext(ctx).ToClusterUpgradeModePtrOutputWithContext(ctx)
}

func (e ClusterUpgradeMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterUpgradeMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ClusterUpgradeMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ClusterUpgradeMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ClusterUpgradeModeOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterUpgradeMode)(nil)).Elem()
}

func (o ClusterUpgradeModeOutput) ToClusterUpgradeModeOutput() ClusterUpgradeModeOutput {
	return o
}

func (o ClusterUpgradeModeOutput) ToClusterUpgradeModeOutputWithContext(ctx context.Context) ClusterUpgradeModeOutput {
	return o
}

func (o ClusterUpgradeModeOutput) ToClusterUpgradeModePtrOutput() ClusterUpgradeModePtrOutput {
	return o.ToClusterUpgradeModePtrOutputWithContext(context.Background())
}

func (o ClusterUpgradeModeOutput) ToClusterUpgradeModePtrOutputWithContext(ctx context.Context) ClusterUpgradeModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterUpgradeMode) *ClusterUpgradeMode {
		return &v
	}).(ClusterUpgradeModePtrOutput)
}

func (o ClusterUpgradeModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ClusterUpgradeModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterUpgradeMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ClusterUpgradeModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterUpgradeModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ClusterUpgradeMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ClusterUpgradeModePtrOutput struct{ *pulumi.OutputState }

func (ClusterUpgradeModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterUpgradeMode)(nil)).Elem()
}

func (o ClusterUpgradeModePtrOutput) ToClusterUpgradeModePtrOutput() ClusterUpgradeModePtrOutput {
	return o
}

func (o ClusterUpgradeModePtrOutput) ToClusterUpgradeModePtrOutputWithContext(ctx context.Context) ClusterUpgradeModePtrOutput {
	return o
}

func (o ClusterUpgradeModePtrOutput) Elem() ClusterUpgradeModeOutput {
	return o.ApplyT(func(v *ClusterUpgradeMode) ClusterUpgradeMode {
		if v != nil {
			return *v
		}
		var ret ClusterUpgradeMode
		return ret
	}).(ClusterUpgradeModeOutput)
}

func (o ClusterUpgradeModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ClusterUpgradeModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ClusterUpgradeMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ClusterUpgradeModeInput interface {
	pulumi.Input

	ToClusterUpgradeModeOutput() ClusterUpgradeModeOutput
	ToClusterUpgradeModeOutputWithContext(context.Context) ClusterUpgradeModeOutput
}

var clusterUpgradeModePtrType = reflect.TypeOf((**ClusterUpgradeMode)(nil)).Elem()

type ClusterUpgradeModePtrInput interface {
	pulumi.Input

	ToClusterUpgradeModePtrOutput() ClusterUpgradeModePtrOutput
	ToClusterUpgradeModePtrOutputWithContext(context.Context) ClusterUpgradeModePtrOutput
}

type clusterUpgradeModePtr string

func ClusterUpgradeModePtr(v string) ClusterUpgradeModePtrInput {
	return (*clusterUpgradeModePtr)(&v)
}

func (*clusterUpgradeModePtr) ElementType() reflect.Type {
	return clusterUpgradeModePtrType
}

func (in *clusterUpgradeModePtr) ToClusterUpgradeModePtrOutput() ClusterUpgradeModePtrOutput {
	return pulumi.ToOutput(in).(ClusterUpgradeModePtrOutput)
}

func (in *clusterUpgradeModePtr) ToClusterUpgradeModePtrOutputWithContext(ctx context.Context) ClusterUpgradeModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ClusterUpgradeModePtrOutput)
}

type Direction string

const (
	DirectionInbound  = Direction("inbound")
	DirectionOutbound = Direction("outbound")
)

func (Direction) ElementType() reflect.Type {
	return reflect.TypeOf((*Direction)(nil)).Elem()
}

func (e Direction) ToDirectionOutput() DirectionOutput {
	return pulumi.ToOutput(e).(DirectionOutput)
}

func (e Direction) ToDirectionOutputWithContext(ctx context.Context) DirectionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DirectionOutput)
}

func (e Direction) ToDirectionPtrOutput() DirectionPtrOutput {
	return e.ToDirectionPtrOutputWithContext(context.Background())
}

func (e Direction) ToDirectionPtrOutputWithContext(ctx context.Context) DirectionPtrOutput {
	return Direction(e).ToDirectionOutputWithContext(ctx).ToDirectionPtrOutputWithContext(ctx)
}

func (e Direction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Direction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Direction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Direction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DirectionOutput struct{ *pulumi.OutputState }

func (DirectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Direction)(nil)).Elem()
}

func (o DirectionOutput) ToDirectionOutput() DirectionOutput {
	return o
}

func (o DirectionOutput) ToDirectionOutputWithContext(ctx context.Context) DirectionOutput {
	return o
}

func (o DirectionOutput) ToDirectionPtrOutput() DirectionPtrOutput {
	return o.ToDirectionPtrOutputWithContext(context.Background())
}

func (o DirectionOutput) ToDirectionPtrOutputWithContext(ctx context.Context) DirectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Direction) *Direction {
		return &v
	}).(DirectionPtrOutput)
}

func (o DirectionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DirectionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Direction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DirectionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DirectionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Direction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DirectionPtrOutput struct{ *pulumi.OutputState }

func (DirectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Direction)(nil)).Elem()
}

func (o DirectionPtrOutput) ToDirectionPtrOutput() DirectionPtrOutput {
	return o
}

func (o DirectionPtrOutput) ToDirectionPtrOutputWithContext(ctx context.Context) DirectionPtrOutput {
	return o
}

func (o DirectionPtrOutput) Elem() DirectionOutput {
	return o.ApplyT(func(v *Direction) Direction {
		if v != nil {
			return *v
		}
		var ret Direction
		return ret
	}).(DirectionOutput)
}

func (o DirectionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DirectionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Direction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DirectionInput interface {
	pulumi.Input

	ToDirectionOutput() DirectionOutput
	ToDirectionOutputWithContext(context.Context) DirectionOutput
}

var directionPtrType = reflect.TypeOf((**Direction)(nil)).Elem()

type DirectionPtrInput interface {
	pulumi.Input

	ToDirectionPtrOutput() DirectionPtrOutput
	ToDirectionPtrOutputWithContext(context.Context) DirectionPtrOutput
}

type directionPtr string

func DirectionPtr(v string) DirectionPtrInput {
	return (*directionPtr)(&v)
}

func (*directionPtr) ElementType() reflect.Type {
	return directionPtrType
}

func (in *directionPtr) ToDirectionPtrOutput() DirectionPtrOutput {
	return pulumi.ToOutput(in).(DirectionPtrOutput)
}

func (in *directionPtr) ToDirectionPtrOutputWithContext(ctx context.Context) DirectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DirectionPtrOutput)
}

type DiskType string

const (
	// Standard HDD locally redundant storage. Best for backup, non-critical, and infrequent access.
	DiskType_Standard_LRS = DiskType("Standard_LRS")
	// Standard SSD locally redundant storage. Best for web servers, lightly used enterprise applications and dev/test.
	DiskType_StandardSSD_LRS = DiskType("StandardSSD_LRS")
	// Premium SSD locally redundant storage. Best for production and performance sensitive workloads.
	DiskType_Premium_LRS = DiskType("Premium_LRS")
)

func (DiskType) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskType)(nil)).Elem()
}

func (e DiskType) ToDiskTypeOutput() DiskTypeOutput {
	return pulumi.ToOutput(e).(DiskTypeOutput)
}

func (e DiskType) ToDiskTypeOutputWithContext(ctx context.Context) DiskTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DiskTypeOutput)
}

func (e DiskType) ToDiskTypePtrOutput() DiskTypePtrOutput {
	return e.ToDiskTypePtrOutputWithContext(context.Background())
}

func (e DiskType) ToDiskTypePtrOutputWithContext(ctx context.Context) DiskTypePtrOutput {
	return DiskType(e).ToDiskTypeOutputWithContext(ctx).ToDiskTypePtrOutputWithContext(ctx)
}

func (e DiskType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DiskType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DiskType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DiskTypeOutput struct{ *pulumi.OutputState }

func (DiskTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskType)(nil)).Elem()
}

func (o DiskTypeOutput) ToDiskTypeOutput() DiskTypeOutput {
	return o
}

func (o DiskTypeOutput) ToDiskTypeOutputWithContext(ctx context.Context) DiskTypeOutput {
	return o
}

func (o DiskTypeOutput) ToDiskTypePtrOutput() DiskTypePtrOutput {
	return o.ToDiskTypePtrOutputWithContext(context.Background())
}

func (o DiskTypeOutput) ToDiskTypePtrOutputWithContext(ctx context.Context) DiskTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiskType) *DiskType {
		return &v
	}).(DiskTypePtrOutput)
}

func (o DiskTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DiskTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DiskTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DiskType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DiskTypePtrOutput struct{ *pulumi.OutputState }

func (DiskTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiskType)(nil)).Elem()
}

func (o DiskTypePtrOutput) ToDiskTypePtrOutput() DiskTypePtrOutput {
	return o
}

func (o DiskTypePtrOutput) ToDiskTypePtrOutputWithContext(ctx context.Context) DiskTypePtrOutput {
	return o
}

func (o DiskTypePtrOutput) Elem() DiskTypeOutput {
	return o.ApplyT(func(v *DiskType) DiskType {
		if v != nil {
			return *v
		}
		var ret DiskType
		return ret
	}).(DiskTypeOutput)
}

func (o DiskTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DiskTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DiskType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DiskTypeInput interface {
	pulumi.Input

	ToDiskTypeOutput() DiskTypeOutput
	ToDiskTypeOutputWithContext(context.Context) DiskTypeOutput
}

var diskTypePtrType = reflect.TypeOf((**DiskType)(nil)).Elem()

type DiskTypePtrInput interface {
	pulumi.Input

	ToDiskTypePtrOutput() DiskTypePtrOutput
	ToDiskTypePtrOutputWithContext(context.Context) DiskTypePtrOutput
}

type diskTypePtr string

func DiskTypePtr(v string) DiskTypePtrInput {
	return (*diskTypePtr)(&v)
}

func (*diskTypePtr) ElementType() reflect.Type {
	return diskTypePtrType
}

func (in *diskTypePtr) ToDiskTypePtrOutput() DiskTypePtrOutput {
	return pulumi.ToOutput(in).(DiskTypePtrOutput)
}

func (in *diskTypePtr) ToDiskTypePtrOutputWithContext(ctx context.Context) DiskTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DiskTypePtrOutput)
}

type FailureAction string

const (
	// Indicates that a rollback of the upgrade will be performed by Service Fabric if the upgrade fails.
	FailureActionRollback = FailureAction("Rollback")
	// Indicates that a manual repair will need to be performed by the administrator if the upgrade fails. Service Fabric will not proceed to the next upgrade domain automatically.
	FailureActionManual = FailureAction("Manual")
)

func (FailureAction) ElementType() reflect.Type {
	return reflect.TypeOf((*FailureAction)(nil)).Elem()
}

func (e FailureAction) ToFailureActionOutput() FailureActionOutput {
	return pulumi.ToOutput(e).(FailureActionOutput)
}

func (e FailureAction) ToFailureActionOutputWithContext(ctx context.Context) FailureActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FailureActionOutput)
}

func (e FailureAction) ToFailureActionPtrOutput() FailureActionPtrOutput {
	return e.ToFailureActionPtrOutputWithContext(context.Background())
}

func (e FailureAction) ToFailureActionPtrOutputWithContext(ctx context.Context) FailureActionPtrOutput {
	return FailureAction(e).ToFailureActionOutputWithContext(ctx).ToFailureActionPtrOutputWithContext(ctx)
}

func (e FailureAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FailureAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FailureAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FailureAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FailureActionOutput struct{ *pulumi.OutputState }

func (FailureActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailureAction)(nil)).Elem()
}

func (o FailureActionOutput) ToFailureActionOutput() FailureActionOutput {
	return o
}

func (o FailureActionOutput) ToFailureActionOutputWithContext(ctx context.Context) FailureActionOutput {
	return o
}

func (o FailureActionOutput) ToFailureActionPtrOutput() FailureActionPtrOutput {
	return o.ToFailureActionPtrOutputWithContext(context.Background())
}

func (o FailureActionOutput) ToFailureActionPtrOutputWithContext(ctx context.Context) FailureActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailureAction) *FailureAction {
		return &v
	}).(FailureActionPtrOutput)
}

func (o FailureActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FailureActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FailureAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FailureActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FailureActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FailureAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FailureActionPtrOutput struct{ *pulumi.OutputState }

func (FailureActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailureAction)(nil)).Elem()
}

func (o FailureActionPtrOutput) ToFailureActionPtrOutput() FailureActionPtrOutput {
	return o
}

func (o FailureActionPtrOutput) ToFailureActionPtrOutputWithContext(ctx context.Context) FailureActionPtrOutput {
	return o
}

func (o FailureActionPtrOutput) Elem() FailureActionOutput {
	return o.ApplyT(func(v *FailureAction) FailureAction {
		if v != nil {
			return *v
		}
		var ret FailureAction
		return ret
	}).(FailureActionOutput)
}

func (o FailureActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FailureActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FailureAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FailureActionInput interface {
	pulumi.Input

	ToFailureActionOutput() FailureActionOutput
	ToFailureActionOutputWithContext(context.Context) FailureActionOutput
}

var failureActionPtrType = reflect.TypeOf((**FailureAction)(nil)).Elem()

type FailureActionPtrInput interface {
	pulumi.Input

	ToFailureActionPtrOutput() FailureActionPtrOutput
	ToFailureActionPtrOutputWithContext(context.Context) FailureActionPtrOutput
}

type failureActionPtr string

func FailureActionPtr(v string) FailureActionPtrInput {
	return (*failureActionPtr)(&v)
}

func (*failureActionPtr) ElementType() reflect.Type {
	return failureActionPtrType
}

func (in *failureActionPtr) ToFailureActionPtrOutput() FailureActionPtrOutput {
	return pulumi.ToOutput(in).(FailureActionPtrOutput)
}

func (in *failureActionPtr) ToFailureActionPtrOutputWithContext(ctx context.Context) FailureActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FailureActionPtrOutput)
}

type IPAddressType string

const (
	// IPv4 address type.
	IPAddressTypeIPv4 = IPAddressType("IPv4")
	// IPv6 address type.
	IPAddressTypeIPv6 = IPAddressType("IPv6")
)

func (IPAddressType) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAddressType)(nil)).Elem()
}

func (e IPAddressType) ToIPAddressTypeOutput() IPAddressTypeOutput {
	return pulumi.ToOutput(e).(IPAddressTypeOutput)
}

func (e IPAddressType) ToIPAddressTypeOutputWithContext(ctx context.Context) IPAddressTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IPAddressTypeOutput)
}

func (e IPAddressType) ToIPAddressTypePtrOutput() IPAddressTypePtrOutput {
	return e.ToIPAddressTypePtrOutputWithContext(context.Background())
}

func (e IPAddressType) ToIPAddressTypePtrOutputWithContext(ctx context.Context) IPAddressTypePtrOutput {
	return IPAddressType(e).ToIPAddressTypeOutputWithContext(ctx).ToIPAddressTypePtrOutputWithContext(ctx)
}

func (e IPAddressType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAddressType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IPAddressType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IPAddressType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IPAddressTypeOutput struct{ *pulumi.OutputState }

func (IPAddressTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPAddressType)(nil)).Elem()
}

func (o IPAddressTypeOutput) ToIPAddressTypeOutput() IPAddressTypeOutput {
	return o
}

func (o IPAddressTypeOutput) ToIPAddressTypeOutputWithContext(ctx context.Context) IPAddressTypeOutput {
	return o
}

func (o IPAddressTypeOutput) ToIPAddressTypePtrOutput() IPAddressTypePtrOutput {
	return o.ToIPAddressTypePtrOutputWithContext(context.Background())
}

func (o IPAddressTypeOutput) ToIPAddressTypePtrOutputWithContext(ctx context.Context) IPAddressTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IPAddressType) *IPAddressType {
		return &v
	}).(IPAddressTypePtrOutput)
}

func (o IPAddressTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IPAddressTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPAddressType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IPAddressTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPAddressTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IPAddressType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IPAddressTypePtrOutput struct{ *pulumi.OutputState }

func (IPAddressTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IPAddressType)(nil)).Elem()
}

func (o IPAddressTypePtrOutput) ToIPAddressTypePtrOutput() IPAddressTypePtrOutput {
	return o
}

func (o IPAddressTypePtrOutput) ToIPAddressTypePtrOutputWithContext(ctx context.Context) IPAddressTypePtrOutput {
	return o
}

func (o IPAddressTypePtrOutput) Elem() IPAddressTypeOutput {
	return o.ApplyT(func(v *IPAddressType) IPAddressType {
		if v != nil {
			return *v
		}
		var ret IPAddressType
		return ret
	}).(IPAddressTypeOutput)
}

func (o IPAddressTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IPAddressTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IPAddressType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IPAddressTypeInput interface {
	pulumi.Input

	ToIPAddressTypeOutput() IPAddressTypeOutput
	ToIPAddressTypeOutputWithContext(context.Context) IPAddressTypeOutput
}

var ipaddressTypePtrType = reflect.TypeOf((**IPAddressType)(nil)).Elem()

type IPAddressTypePtrInput interface {
	pulumi.Input

	ToIPAddressTypePtrOutput() IPAddressTypePtrOutput
	ToIPAddressTypePtrOutputWithContext(context.Context) IPAddressTypePtrOutput
}

type ipaddressTypePtr string

func IPAddressTypePtr(v string) IPAddressTypePtrInput {
	return (*ipaddressTypePtr)(&v)
}

func (*ipaddressTypePtr) ElementType() reflect.Type {
	return ipaddressTypePtrType
}

func (in *ipaddressTypePtr) ToIPAddressTypePtrOutput() IPAddressTypePtrOutput {
	return pulumi.ToOutput(in).(IPAddressTypePtrOutput)
}

func (in *ipaddressTypePtr) ToIPAddressTypePtrOutputWithContext(ctx context.Context) IPAddressTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IPAddressTypePtrOutput)
}

type ManagedIdentityType string

const (
	// Indicates that no identity is associated with the resource.
	ManagedIdentityTypeNone = ManagedIdentityType("None")
	// Indicates that system assigned identity is associated with the resource.
	ManagedIdentityTypeSystemAssigned = ManagedIdentityType("SystemAssigned")
	// Indicates that user assigned identity is associated with the resource.
	ManagedIdentityTypeUserAssigned = ManagedIdentityType("UserAssigned")
	// Indicates that both system assigned and user assigned identity are associated with the resource.
	ManagedIdentityType_SystemAssigned_UserAssigned = ManagedIdentityType("SystemAssigned, UserAssigned")
)

func (ManagedIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityType)(nil)).Elem()
}

func (e ManagedIdentityType) ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput {
	return pulumi.ToOutput(e).(ManagedIdentityTypeOutput)
}

func (e ManagedIdentityType) ToManagedIdentityTypeOutputWithContext(ctx context.Context) ManagedIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ManagedIdentityTypeOutput)
}

func (e ManagedIdentityType) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return e.ToManagedIdentityTypePtrOutputWithContext(context.Background())
}

func (e ManagedIdentityType) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return ManagedIdentityType(e).ToManagedIdentityTypeOutputWithContext(ctx).ToManagedIdentityTypePtrOutputWithContext(ctx)
}

func (e ManagedIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ManagedIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ManagedIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ManagedIdentityTypeOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedIdentityType)(nil)).Elem()
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput {
	return o
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypeOutputWithContext(ctx context.Context) ManagedIdentityTypeOutput {
	return o
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return o.ToManagedIdentityTypePtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedIdentityType) *ManagedIdentityType {
		return &v
	}).(ManagedIdentityTypePtrOutput)
}

func (o ManagedIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ManagedIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ManagedIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ManagedIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ManagedIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedIdentityType)(nil)).Elem()
}

func (o ManagedIdentityTypePtrOutput) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return o
}

func (o ManagedIdentityTypePtrOutput) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return o
}

func (o ManagedIdentityTypePtrOutput) Elem() ManagedIdentityTypeOutput {
	return o.ApplyT(func(v *ManagedIdentityType) ManagedIdentityType {
		if v != nil {
			return *v
		}
		var ret ManagedIdentityType
		return ret
	}).(ManagedIdentityTypeOutput)
}

func (o ManagedIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ManagedIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ManagedIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ManagedIdentityTypeInput interface {
	pulumi.Input

	ToManagedIdentityTypeOutput() ManagedIdentityTypeOutput
	ToManagedIdentityTypeOutputWithContext(context.Context) ManagedIdentityTypeOutput
}

var managedIdentityTypePtrType = reflect.TypeOf((**ManagedIdentityType)(nil)).Elem()

type ManagedIdentityTypePtrInput interface {
	pulumi.Input

	ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput
	ToManagedIdentityTypePtrOutputWithContext(context.Context) ManagedIdentityTypePtrOutput
}

type managedIdentityTypePtr string

func ManagedIdentityTypePtr(v string) ManagedIdentityTypePtrInput {
	return (*managedIdentityTypePtr)(&v)
}

func (*managedIdentityTypePtr) ElementType() reflect.Type {
	return managedIdentityTypePtrType
}

func (in *managedIdentityTypePtr) ToManagedIdentityTypePtrOutput() ManagedIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ManagedIdentityTypePtrOutput)
}

func (in *managedIdentityTypePtr) ToManagedIdentityTypePtrOutputWithContext(ctx context.Context) ManagedIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ManagedIdentityTypePtrOutput)
}

type MoveCost string

const (
	// Zero move cost. This value is zero.
	MoveCostZero = MoveCost("Zero")
	// Specifies the move cost of the service as Low. The value is 1.
	MoveCostLow = MoveCost("Low")
	// Specifies the move cost of the service as Medium. The value is 2.
	MoveCostMedium = MoveCost("Medium")
	// Specifies the move cost of the service as High. The value is 3.
	MoveCostHigh = MoveCost("High")
)

func (MoveCost) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCost)(nil)).Elem()
}

func (e MoveCost) ToMoveCostOutput() MoveCostOutput {
	return pulumi.ToOutput(e).(MoveCostOutput)
}

func (e MoveCost) ToMoveCostOutputWithContext(ctx context.Context) MoveCostOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MoveCostOutput)
}

func (e MoveCost) ToMoveCostPtrOutput() MoveCostPtrOutput {
	return e.ToMoveCostPtrOutputWithContext(context.Background())
}

func (e MoveCost) ToMoveCostPtrOutputWithContext(ctx context.Context) MoveCostPtrOutput {
	return MoveCost(e).ToMoveCostOutputWithContext(ctx).ToMoveCostPtrOutputWithContext(ctx)
}

func (e MoveCost) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MoveCost) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MoveCost) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MoveCost) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MoveCostOutput struct{ *pulumi.OutputState }

func (MoveCostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MoveCost)(nil)).Elem()
}

func (o MoveCostOutput) ToMoveCostOutput() MoveCostOutput {
	return o
}

func (o MoveCostOutput) ToMoveCostOutputWithContext(ctx context.Context) MoveCostOutput {
	return o
}

func (o MoveCostOutput) ToMoveCostPtrOutput() MoveCostPtrOutput {
	return o.ToMoveCostPtrOutputWithContext(context.Background())
}

func (o MoveCostOutput) ToMoveCostPtrOutputWithContext(ctx context.Context) MoveCostPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MoveCost) *MoveCost {
		return &v
	}).(MoveCostPtrOutput)
}

func (o MoveCostOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MoveCostOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MoveCost) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MoveCostOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MoveCostOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MoveCost) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MoveCostPtrOutput struct{ *pulumi.OutputState }

func (MoveCostPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MoveCost)(nil)).Elem()
}

func (o MoveCostPtrOutput) ToMoveCostPtrOutput() MoveCostPtrOutput {
	return o
}

func (o MoveCostPtrOutput) ToMoveCostPtrOutputWithContext(ctx context.Context) MoveCostPtrOutput {
	return o
}

func (o MoveCostPtrOutput) Elem() MoveCostOutput {
	return o.ApplyT(func(v *MoveCost) MoveCost {
		if v != nil {
			return *v
		}
		var ret MoveCost
		return ret
	}).(MoveCostOutput)
}

func (o MoveCostPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MoveCostPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MoveCost) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MoveCostInput interface {
	pulumi.Input

	ToMoveCostOutput() MoveCostOutput
	ToMoveCostOutputWithContext(context.Context) MoveCostOutput
}

var moveCostPtrType = reflect.TypeOf((**MoveCost)(nil)).Elem()

type MoveCostPtrInput interface {
	pulumi.Input

	ToMoveCostPtrOutput() MoveCostPtrOutput
	ToMoveCostPtrOutputWithContext(context.Context) MoveCostPtrOutput
}

type moveCostPtr string

func MoveCostPtr(v string) MoveCostPtrInput {
	return (*moveCostPtr)(&v)
}

func (*moveCostPtr) ElementType() reflect.Type {
	return moveCostPtrType
}

func (in *moveCostPtr) ToMoveCostPtrOutput() MoveCostPtrOutput {
	return pulumi.ToOutput(in).(MoveCostPtrOutput)
}

func (in *moveCostPtr) ToMoveCostPtrOutputWithContext(ctx context.Context) MoveCostPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MoveCostPtrOutput)
}

type NsgProtocol string

const (
	NsgProtocolHttp  = NsgProtocol("http")
	NsgProtocolHttps = NsgProtocol("https")
	NsgProtocolTcp   = NsgProtocol("tcp")
	NsgProtocolUdp   = NsgProtocol("udp")
	NsgProtocolIcmp  = NsgProtocol("icmp")
	NsgProtocolAh    = NsgProtocol("ah")
	NsgProtocolEsp   = NsgProtocol("esp")
)

func (NsgProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgProtocol)(nil)).Elem()
}

func (e NsgProtocol) ToNsgProtocolOutput() NsgProtocolOutput {
	return pulumi.ToOutput(e).(NsgProtocolOutput)
}

func (e NsgProtocol) ToNsgProtocolOutputWithContext(ctx context.Context) NsgProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NsgProtocolOutput)
}

func (e NsgProtocol) ToNsgProtocolPtrOutput() NsgProtocolPtrOutput {
	return e.ToNsgProtocolPtrOutputWithContext(context.Background())
}

func (e NsgProtocol) ToNsgProtocolPtrOutputWithContext(ctx context.Context) NsgProtocolPtrOutput {
	return NsgProtocol(e).ToNsgProtocolOutputWithContext(ctx).ToNsgProtocolPtrOutputWithContext(ctx)
}

func (e NsgProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NsgProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NsgProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NsgProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NsgProtocolOutput struct{ *pulumi.OutputState }

func (NsgProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NsgProtocol)(nil)).Elem()
}

func (o NsgProtocolOutput) ToNsgProtocolOutput() NsgProtocolOutput {
	return o
}

func (o NsgProtocolOutput) ToNsgProtocolOutputWithContext(ctx context.Context) NsgProtocolOutput {
	return o
}

func (o NsgProtocolOutput) ToNsgProtocolPtrOutput() NsgProtocolPtrOutput {
	return o.ToNsgProtocolPtrOutputWithContext(context.Background())
}

func (o NsgProtocolOutput) ToNsgProtocolPtrOutputWithContext(ctx context.Context) NsgProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NsgProtocol) *NsgProtocol {
		return &v
	}).(NsgProtocolPtrOutput)
}

func (o NsgProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NsgProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NsgProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NsgProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NsgProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NsgProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NsgProtocolPtrOutput struct{ *pulumi.OutputState }

func (NsgProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsgProtocol)(nil)).Elem()
}

func (o NsgProtocolPtrOutput) ToNsgProtocolPtrOutput() NsgProtocolPtrOutput {
	return o
}

func (o NsgProtocolPtrOutput) ToNsgProtocolPtrOutputWithContext(ctx context.Context) NsgProtocolPtrOutput {
	return o
}

func (o NsgProtocolPtrOutput) Elem() NsgProtocolOutput {
	return o.ApplyT(func(v *NsgProtocol) NsgProtocol {
		if v != nil {
			return *v
		}
		var ret NsgProtocol
		return ret
	}).(NsgProtocolOutput)
}

func (o NsgProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NsgProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NsgProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NsgProtocolInput interface {
	pulumi.Input

	ToNsgProtocolOutput() NsgProtocolOutput
	ToNsgProtocolOutputWithContext(context.Context) NsgProtocolOutput
}

var nsgProtocolPtrType = reflect.TypeOf((**NsgProtocol)(nil)).Elem()

type NsgProtocolPtrInput interface {
	pulumi.Input

	ToNsgProtocolPtrOutput() NsgProtocolPtrOutput
	ToNsgProtocolPtrOutputWithContext(context.Context) NsgProtocolPtrOutput
}

type nsgProtocolPtr string

func NsgProtocolPtr(v string) NsgProtocolPtrInput {
	return (*nsgProtocolPtr)(&v)
}

func (*nsgProtocolPtr) ElementType() reflect.Type {
	return nsgProtocolPtrType
}

func (in *nsgProtocolPtr) ToNsgProtocolPtrOutput() NsgProtocolPtrOutput {
	return pulumi.ToOutput(in).(NsgProtocolPtrOutput)
}

func (in *nsgProtocolPtr) ToNsgProtocolPtrOutputWithContext(ctx context.Context) NsgProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NsgProtocolPtrOutput)
}

type PartitionScheme string

const (
	// Indicates that the partition is based on string names, and is a SingletonPartitionScheme object, The value is 0.
	PartitionSchemeSingleton = PartitionScheme("Singleton")
	// Indicates that the partition is based on Int64 key ranges, and is a UniformInt64RangePartitionScheme object. The value is 1.
	PartitionSchemeUniformInt64Range = PartitionScheme("UniformInt64Range")
	// Indicates that the partition is based on string names, and is a NamedPartitionScheme object. The value is 2.
	PartitionSchemeNamed = PartitionScheme("Named")
)

func (PartitionScheme) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionScheme)(nil)).Elem()
}

func (e PartitionScheme) ToPartitionSchemeOutput() PartitionSchemeOutput {
	return pulumi.ToOutput(e).(PartitionSchemeOutput)
}

func (e PartitionScheme) ToPartitionSchemeOutputWithContext(ctx context.Context) PartitionSchemeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PartitionSchemeOutput)
}

func (e PartitionScheme) ToPartitionSchemePtrOutput() PartitionSchemePtrOutput {
	return e.ToPartitionSchemePtrOutputWithContext(context.Background())
}

func (e PartitionScheme) ToPartitionSchemePtrOutputWithContext(ctx context.Context) PartitionSchemePtrOutput {
	return PartitionScheme(e).ToPartitionSchemeOutputWithContext(ctx).ToPartitionSchemePtrOutputWithContext(ctx)
}

func (e PartitionScheme) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartitionScheme) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartitionScheme) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PartitionScheme) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PartitionSchemeOutput struct{ *pulumi.OutputState }

func (PartitionSchemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartitionScheme)(nil)).Elem()
}

func (o PartitionSchemeOutput) ToPartitionSchemeOutput() PartitionSchemeOutput {
	return o
}

func (o PartitionSchemeOutput) ToPartitionSchemeOutputWithContext(ctx context.Context) PartitionSchemeOutput {
	return o
}

func (o PartitionSchemeOutput) ToPartitionSchemePtrOutput() PartitionSchemePtrOutput {
	return o.ToPartitionSchemePtrOutputWithContext(context.Background())
}

func (o PartitionSchemeOutput) ToPartitionSchemePtrOutputWithContext(ctx context.Context) PartitionSchemePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartitionScheme) *PartitionScheme {
		return &v
	}).(PartitionSchemePtrOutput)
}

func (o PartitionSchemeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PartitionSchemeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartitionScheme) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PartitionSchemeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartitionSchemeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartitionScheme) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PartitionSchemePtrOutput struct{ *pulumi.OutputState }

func (PartitionSchemePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartitionScheme)(nil)).Elem()
}

func (o PartitionSchemePtrOutput) ToPartitionSchemePtrOutput() PartitionSchemePtrOutput {
	return o
}

func (o PartitionSchemePtrOutput) ToPartitionSchemePtrOutputWithContext(ctx context.Context) PartitionSchemePtrOutput {
	return o
}

func (o PartitionSchemePtrOutput) Elem() PartitionSchemeOutput {
	return o.ApplyT(func(v *PartitionScheme) PartitionScheme {
		if v != nil {
			return *v
		}
		var ret PartitionScheme
		return ret
	}).(PartitionSchemeOutput)
}

func (o PartitionSchemePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartitionSchemePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PartitionScheme) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PartitionSchemeInput interface {
	pulumi.Input

	ToPartitionSchemeOutput() PartitionSchemeOutput
	ToPartitionSchemeOutputWithContext(context.Context) PartitionSchemeOutput
}

var partitionSchemePtrType = reflect.TypeOf((**PartitionScheme)(nil)).Elem()

type PartitionSchemePtrInput interface {
	pulumi.Input

	ToPartitionSchemePtrOutput() PartitionSchemePtrOutput
	ToPartitionSchemePtrOutputWithContext(context.Context) PartitionSchemePtrOutput
}

type partitionSchemePtr string

func PartitionSchemePtr(v string) PartitionSchemePtrInput {
	return (*partitionSchemePtr)(&v)
}

func (*partitionSchemePtr) ElementType() reflect.Type {
	return partitionSchemePtrType
}

func (in *partitionSchemePtr) ToPartitionSchemePtrOutput() PartitionSchemePtrOutput {
	return pulumi.ToOutput(in).(PartitionSchemePtrOutput)
}

func (in *partitionSchemePtr) ToPartitionSchemePtrOutputWithContext(ctx context.Context) PartitionSchemePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PartitionSchemePtrOutput)
}

type PrivateEndpointNetworkPolicies string

const (
	PrivateEndpointNetworkPoliciesEnabled  = PrivateEndpointNetworkPolicies("enabled")
	PrivateEndpointNetworkPoliciesDisabled = PrivateEndpointNetworkPolicies("disabled")
)

func (PrivateEndpointNetworkPolicies) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointNetworkPolicies)(nil)).Elem()
}

func (e PrivateEndpointNetworkPolicies) ToPrivateEndpointNetworkPoliciesOutput() PrivateEndpointNetworkPoliciesOutput {
	return pulumi.ToOutput(e).(PrivateEndpointNetworkPoliciesOutput)
}

func (e PrivateEndpointNetworkPolicies) ToPrivateEndpointNetworkPoliciesOutputWithContext(ctx context.Context) PrivateEndpointNetworkPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateEndpointNetworkPoliciesOutput)
}

func (e PrivateEndpointNetworkPolicies) ToPrivateEndpointNetworkPoliciesPtrOutput() PrivateEndpointNetworkPoliciesPtrOutput {
	return e.ToPrivateEndpointNetworkPoliciesPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointNetworkPolicies) ToPrivateEndpointNetworkPoliciesPtrOutputWithContext(ctx context.Context) PrivateEndpointNetworkPoliciesPtrOutput {
	return PrivateEndpointNetworkPolicies(e).ToPrivateEndpointNetworkPoliciesOutputWithContext(ctx).ToPrivateEndpointNetworkPoliciesPtrOutputWithContext(ctx)
}

func (e PrivateEndpointNetworkPolicies) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointNetworkPolicies) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateEndpointNetworkPolicies) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateEndpointNetworkPolicies) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateEndpointNetworkPoliciesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointNetworkPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointNetworkPolicies)(nil)).Elem()
}

func (o PrivateEndpointNetworkPoliciesOutput) ToPrivateEndpointNetworkPoliciesOutput() PrivateEndpointNetworkPoliciesOutput {
	return o
}

func (o PrivateEndpointNetworkPoliciesOutput) ToPrivateEndpointNetworkPoliciesOutputWithContext(ctx context.Context) PrivateEndpointNetworkPoliciesOutput {
	return o
}

func (o PrivateEndpointNetworkPoliciesOutput) ToPrivateEndpointNetworkPoliciesPtrOutput() PrivateEndpointNetworkPoliciesPtrOutput {
	return o.ToPrivateEndpointNetworkPoliciesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointNetworkPoliciesOutput) ToPrivateEndpointNetworkPoliciesPtrOutputWithContext(ctx context.Context) PrivateEndpointNetworkPoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointNetworkPolicies) *PrivateEndpointNetworkPolicies {
		return &v
	}).(PrivateEndpointNetworkPoliciesPtrOutput)
}

func (o PrivateEndpointNetworkPoliciesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateEndpointNetworkPoliciesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointNetworkPolicies) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateEndpointNetworkPoliciesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointNetworkPoliciesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateEndpointNetworkPolicies) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointNetworkPoliciesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointNetworkPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointNetworkPolicies)(nil)).Elem()
}

func (o PrivateEndpointNetworkPoliciesPtrOutput) ToPrivateEndpointNetworkPoliciesPtrOutput() PrivateEndpointNetworkPoliciesPtrOutput {
	return o
}

func (o PrivateEndpointNetworkPoliciesPtrOutput) ToPrivateEndpointNetworkPoliciesPtrOutputWithContext(ctx context.Context) PrivateEndpointNetworkPoliciesPtrOutput {
	return o
}

func (o PrivateEndpointNetworkPoliciesPtrOutput) Elem() PrivateEndpointNetworkPoliciesOutput {
	return o.ApplyT(func(v *PrivateEndpointNetworkPolicies) PrivateEndpointNetworkPolicies {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointNetworkPolicies
		return ret
	}).(PrivateEndpointNetworkPoliciesOutput)
}

func (o PrivateEndpointNetworkPoliciesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointNetworkPoliciesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateEndpointNetworkPolicies) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateEndpointNetworkPoliciesInput interface {
	pulumi.Input

	ToPrivateEndpointNetworkPoliciesOutput() PrivateEndpointNetworkPoliciesOutput
	ToPrivateEndpointNetworkPoliciesOutputWithContext(context.Context) PrivateEndpointNetworkPoliciesOutput
}

var privateEndpointNetworkPoliciesPtrType = reflect.TypeOf((**PrivateEndpointNetworkPolicies)(nil)).Elem()

type PrivateEndpointNetworkPoliciesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointNetworkPoliciesPtrOutput() PrivateEndpointNetworkPoliciesPtrOutput
	ToPrivateEndpointNetworkPoliciesPtrOutputWithContext(context.Context) PrivateEndpointNetworkPoliciesPtrOutput
}

type privateEndpointNetworkPoliciesPtr string

func PrivateEndpointNetworkPoliciesPtr(v string) PrivateEndpointNetworkPoliciesPtrInput {
	return (*privateEndpointNetworkPoliciesPtr)(&v)
}

func (*privateEndpointNetworkPoliciesPtr) ElementType() reflect.Type {
	return privateEndpointNetworkPoliciesPtrType
}

func (in *privateEndpointNetworkPoliciesPtr) ToPrivateEndpointNetworkPoliciesPtrOutput() PrivateEndpointNetworkPoliciesPtrOutput {
	return pulumi.ToOutput(in).(PrivateEndpointNetworkPoliciesPtrOutput)
}

func (in *privateEndpointNetworkPoliciesPtr) ToPrivateEndpointNetworkPoliciesPtrOutputWithContext(ctx context.Context) PrivateEndpointNetworkPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateEndpointNetworkPoliciesPtrOutput)
}

type PrivateLinkServiceNetworkPolicies string

const (
	PrivateLinkServiceNetworkPoliciesEnabled  = PrivateLinkServiceNetworkPolicies("enabled")
	PrivateLinkServiceNetworkPoliciesDisabled = PrivateLinkServiceNetworkPolicies("disabled")
)

func (PrivateLinkServiceNetworkPolicies) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceNetworkPolicies)(nil)).Elem()
}

func (e PrivateLinkServiceNetworkPolicies) ToPrivateLinkServiceNetworkPoliciesOutput() PrivateLinkServiceNetworkPoliciesOutput {
	return pulumi.ToOutput(e).(PrivateLinkServiceNetworkPoliciesOutput)
}

func (e PrivateLinkServiceNetworkPolicies) ToPrivateLinkServiceNetworkPoliciesOutputWithContext(ctx context.Context) PrivateLinkServiceNetworkPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateLinkServiceNetworkPoliciesOutput)
}

func (e PrivateLinkServiceNetworkPolicies) ToPrivateLinkServiceNetworkPoliciesPtrOutput() PrivateLinkServiceNetworkPoliciesPtrOutput {
	return e.ToPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceNetworkPolicies) ToPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(ctx context.Context) PrivateLinkServiceNetworkPoliciesPtrOutput {
	return PrivateLinkServiceNetworkPolicies(e).ToPrivateLinkServiceNetworkPoliciesOutputWithContext(ctx).ToPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(ctx)
}

func (e PrivateLinkServiceNetworkPolicies) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceNetworkPolicies) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceNetworkPolicies) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceNetworkPolicies) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateLinkServiceNetworkPoliciesOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceNetworkPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceNetworkPolicies)(nil)).Elem()
}

func (o PrivateLinkServiceNetworkPoliciesOutput) ToPrivateLinkServiceNetworkPoliciesOutput() PrivateLinkServiceNetworkPoliciesOutput {
	return o
}

func (o PrivateLinkServiceNetworkPoliciesOutput) ToPrivateLinkServiceNetworkPoliciesOutputWithContext(ctx context.Context) PrivateLinkServiceNetworkPoliciesOutput {
	return o
}

func (o PrivateLinkServiceNetworkPoliciesOutput) ToPrivateLinkServiceNetworkPoliciesPtrOutput() PrivateLinkServiceNetworkPoliciesPtrOutput {
	return o.ToPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceNetworkPoliciesOutput) ToPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(ctx context.Context) PrivateLinkServiceNetworkPoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceNetworkPolicies) *PrivateLinkServiceNetworkPolicies {
		return &v
	}).(PrivateLinkServiceNetworkPoliciesPtrOutput)
}

func (o PrivateLinkServiceNetworkPoliciesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateLinkServiceNetworkPoliciesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceNetworkPolicies) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateLinkServiceNetworkPoliciesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceNetworkPoliciesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceNetworkPolicies) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceNetworkPoliciesPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceNetworkPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceNetworkPolicies)(nil)).Elem()
}

func (o PrivateLinkServiceNetworkPoliciesPtrOutput) ToPrivateLinkServiceNetworkPoliciesPtrOutput() PrivateLinkServiceNetworkPoliciesPtrOutput {
	return o
}

func (o PrivateLinkServiceNetworkPoliciesPtrOutput) ToPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(ctx context.Context) PrivateLinkServiceNetworkPoliciesPtrOutput {
	return o
}

func (o PrivateLinkServiceNetworkPoliciesPtrOutput) Elem() PrivateLinkServiceNetworkPoliciesOutput {
	return o.ApplyT(func(v *PrivateLinkServiceNetworkPolicies) PrivateLinkServiceNetworkPolicies {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceNetworkPolicies
		return ret
	}).(PrivateLinkServiceNetworkPoliciesOutput)
}

func (o PrivateLinkServiceNetworkPoliciesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceNetworkPoliciesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateLinkServiceNetworkPolicies) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateLinkServiceNetworkPoliciesInput interface {
	pulumi.Input

	ToPrivateLinkServiceNetworkPoliciesOutput() PrivateLinkServiceNetworkPoliciesOutput
	ToPrivateLinkServiceNetworkPoliciesOutputWithContext(context.Context) PrivateLinkServiceNetworkPoliciesOutput
}

var privateLinkServiceNetworkPoliciesPtrType = reflect.TypeOf((**PrivateLinkServiceNetworkPolicies)(nil)).Elem()

type PrivateLinkServiceNetworkPoliciesPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceNetworkPoliciesPtrOutput() PrivateLinkServiceNetworkPoliciesPtrOutput
	ToPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(context.Context) PrivateLinkServiceNetworkPoliciesPtrOutput
}

type privateLinkServiceNetworkPoliciesPtr string

func PrivateLinkServiceNetworkPoliciesPtr(v string) PrivateLinkServiceNetworkPoliciesPtrInput {
	return (*privateLinkServiceNetworkPoliciesPtr)(&v)
}

func (*privateLinkServiceNetworkPoliciesPtr) ElementType() reflect.Type {
	return privateLinkServiceNetworkPoliciesPtrType
}

func (in *privateLinkServiceNetworkPoliciesPtr) ToPrivateLinkServiceNetworkPoliciesPtrOutput() PrivateLinkServiceNetworkPoliciesPtrOutput {
	return pulumi.ToOutput(in).(PrivateLinkServiceNetworkPoliciesPtrOutput)
}

func (in *privateLinkServiceNetworkPoliciesPtr) ToPrivateLinkServiceNetworkPoliciesPtrOutputWithContext(ctx context.Context) PrivateLinkServiceNetworkPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateLinkServiceNetworkPoliciesPtrOutput)
}

type ProbeProtocol string

const (
	ProbeProtocolTcp   = ProbeProtocol("tcp")
	ProbeProtocolHttp  = ProbeProtocol("http")
	ProbeProtocolHttps = ProbeProtocol("https")
)

func (ProbeProtocol) ElementType() reflect.Type {
	return reflect.TypeOf((*ProbeProtocol)(nil)).Elem()
}

func (e ProbeProtocol) ToProbeProtocolOutput() ProbeProtocolOutput {
	return pulumi.ToOutput(e).(ProbeProtocolOutput)
}

func (e ProbeProtocol) ToProbeProtocolOutputWithContext(ctx context.Context) ProbeProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProbeProtocolOutput)
}

func (e ProbeProtocol) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return e.ToProbeProtocolPtrOutputWithContext(context.Background())
}

func (e ProbeProtocol) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return ProbeProtocol(e).ToProbeProtocolOutputWithContext(ctx).ToProbeProtocolPtrOutputWithContext(ctx)
}

func (e ProbeProtocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProbeProtocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProbeProtocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProbeProtocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProbeProtocolOutput struct{ *pulumi.OutputState }

func (ProbeProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProbeProtocol)(nil)).Elem()
}

func (o ProbeProtocolOutput) ToProbeProtocolOutput() ProbeProtocolOutput {
	return o
}

func (o ProbeProtocolOutput) ToProbeProtocolOutputWithContext(ctx context.Context) ProbeProtocolOutput {
	return o
}

func (o ProbeProtocolOutput) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return o.ToProbeProtocolPtrOutputWithContext(context.Background())
}

func (o ProbeProtocolOutput) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProbeProtocol) *ProbeProtocol {
		return &v
	}).(ProbeProtocolPtrOutput)
}

func (o ProbeProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProbeProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProbeProtocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProbeProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProbeProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProbeProtocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProbeProtocolPtrOutput struct{ *pulumi.OutputState }

func (ProbeProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProbeProtocol)(nil)).Elem()
}

func (o ProbeProtocolPtrOutput) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return o
}

func (o ProbeProtocolPtrOutput) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return o
}

func (o ProbeProtocolPtrOutput) Elem() ProbeProtocolOutput {
	return o.ApplyT(func(v *ProbeProtocol) ProbeProtocol {
		if v != nil {
			return *v
		}
		var ret ProbeProtocol
		return ret
	}).(ProbeProtocolOutput)
}

func (o ProbeProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProbeProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProbeProtocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProbeProtocolInput interface {
	pulumi.Input

	ToProbeProtocolOutput() ProbeProtocolOutput
	ToProbeProtocolOutputWithContext(context.Context) ProbeProtocolOutput
}

var probeProtocolPtrType = reflect.TypeOf((**ProbeProtocol)(nil)).Elem()

type ProbeProtocolPtrInput interface {
	pulumi.Input

	ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput
	ToProbeProtocolPtrOutputWithContext(context.Context) ProbeProtocolPtrOutput
}

type probeProtocolPtr string

func ProbeProtocolPtr(v string) ProbeProtocolPtrInput {
	return (*probeProtocolPtr)(&v)
}

func (*probeProtocolPtr) ElementType() reflect.Type {
	return probeProtocolPtrType
}

func (in *probeProtocolPtr) ToProbeProtocolPtrOutput() ProbeProtocolPtrOutput {
	return pulumi.ToOutput(in).(ProbeProtocolPtrOutput)
}

func (in *probeProtocolPtr) ToProbeProtocolPtrOutputWithContext(ctx context.Context) ProbeProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProbeProtocolPtrOutput)
}

type Protocol string

const (
	ProtocolTcp = Protocol("tcp")
	ProtocolUdp = Protocol("udp")
)

func (Protocol) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (e Protocol) ToProtocolOutput() ProtocolOutput {
	return pulumi.ToOutput(e).(ProtocolOutput)
}

func (e Protocol) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProtocolOutput)
}

func (e Protocol) ToProtocolPtrOutput() ProtocolPtrOutput {
	return e.ToProtocolPtrOutputWithContext(context.Background())
}

func (e Protocol) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return Protocol(e).ToProtocolOutputWithContext(ctx).ToProtocolPtrOutputWithContext(ctx)
}

func (e Protocol) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Protocol) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Protocol) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProtocolOutput struct{ *pulumi.OutputState }

func (ProtocolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Protocol)(nil)).Elem()
}

func (o ProtocolOutput) ToProtocolOutput() ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolOutputWithContext(ctx context.Context) ProtocolOutput {
	return o
}

func (o ProtocolOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o.ToProtocolPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Protocol) *Protocol {
		return &v
	}).(ProtocolPtrOutput)
}

func (o ProtocolOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProtocolOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Protocol) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProtocolPtrOutput struct{ *pulumi.OutputState }

func (ProtocolPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Protocol)(nil)).Elem()
}

func (o ProtocolPtrOutput) ToProtocolPtrOutput() ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return o
}

func (o ProtocolPtrOutput) Elem() ProtocolOutput {
	return o.ApplyT(func(v *Protocol) Protocol {
		if v != nil {
			return *v
		}
		var ret Protocol
		return ret
	}).(ProtocolOutput)
}

func (o ProtocolPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProtocolPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Protocol) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProtocolInput interface {
	pulumi.Input

	ToProtocolOutput() ProtocolOutput
	ToProtocolOutputWithContext(context.Context) ProtocolOutput
}

var protocolPtrType = reflect.TypeOf((**Protocol)(nil)).Elem()

type ProtocolPtrInput interface {
	pulumi.Input

	ToProtocolPtrOutput() ProtocolPtrOutput
	ToProtocolPtrOutputWithContext(context.Context) ProtocolPtrOutput
}

type protocolPtr string

func ProtocolPtr(v string) ProtocolPtrInput {
	return (*protocolPtr)(&v)
}

func (*protocolPtr) ElementType() reflect.Type {
	return protocolPtrType
}

func (in *protocolPtr) ToProtocolPtrOutput() ProtocolPtrOutput {
	return pulumi.ToOutput(in).(ProtocolPtrOutput)
}

func (in *protocolPtr) ToProtocolPtrOutputWithContext(ctx context.Context) ProtocolPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProtocolPtrOutput)
}

type RollingUpgradeMode string

const (
	// The upgrade will stop after completing each upgrade domain and automatically monitor health before proceeding. The value is 0.
	RollingUpgradeModeMonitored = RollingUpgradeMode("Monitored")
	// The upgrade will proceed automatically without performing any health monitoring. The value is 1.
	RollingUpgradeModeUnmonitoredAuto = RollingUpgradeMode("UnmonitoredAuto")
)

func (RollingUpgradeMode) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradeMode)(nil)).Elem()
}

func (e RollingUpgradeMode) ToRollingUpgradeModeOutput() RollingUpgradeModeOutput {
	return pulumi.ToOutput(e).(RollingUpgradeModeOutput)
}

func (e RollingUpgradeMode) ToRollingUpgradeModeOutputWithContext(ctx context.Context) RollingUpgradeModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RollingUpgradeModeOutput)
}

func (e RollingUpgradeMode) ToRollingUpgradeModePtrOutput() RollingUpgradeModePtrOutput {
	return e.ToRollingUpgradeModePtrOutputWithContext(context.Background())
}

func (e RollingUpgradeMode) ToRollingUpgradeModePtrOutputWithContext(ctx context.Context) RollingUpgradeModePtrOutput {
	return RollingUpgradeMode(e).ToRollingUpgradeModeOutputWithContext(ctx).ToRollingUpgradeModePtrOutputWithContext(ctx)
}

func (e RollingUpgradeMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RollingUpgradeMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RollingUpgradeMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RollingUpgradeMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RollingUpgradeModeOutput struct{ *pulumi.OutputState }

func (RollingUpgradeModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RollingUpgradeMode)(nil)).Elem()
}

func (o RollingUpgradeModeOutput) ToRollingUpgradeModeOutput() RollingUpgradeModeOutput {
	return o
}

func (o RollingUpgradeModeOutput) ToRollingUpgradeModeOutputWithContext(ctx context.Context) RollingUpgradeModeOutput {
	return o
}

func (o RollingUpgradeModeOutput) ToRollingUpgradeModePtrOutput() RollingUpgradeModePtrOutput {
	return o.ToRollingUpgradeModePtrOutputWithContext(context.Background())
}

func (o RollingUpgradeModeOutput) ToRollingUpgradeModePtrOutputWithContext(ctx context.Context) RollingUpgradeModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RollingUpgradeMode) *RollingUpgradeMode {
		return &v
	}).(RollingUpgradeModePtrOutput)
}

func (o RollingUpgradeModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RollingUpgradeModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RollingUpgradeMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RollingUpgradeModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RollingUpgradeModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RollingUpgradeMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RollingUpgradeModePtrOutput struct{ *pulumi.OutputState }

func (RollingUpgradeModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RollingUpgradeMode)(nil)).Elem()
}

func (o RollingUpgradeModePtrOutput) ToRollingUpgradeModePtrOutput() RollingUpgradeModePtrOutput {
	return o
}

func (o RollingUpgradeModePtrOutput) ToRollingUpgradeModePtrOutputWithContext(ctx context.Context) RollingUpgradeModePtrOutput {
	return o
}

func (o RollingUpgradeModePtrOutput) Elem() RollingUpgradeModeOutput {
	return o.ApplyT(func(v *RollingUpgradeMode) RollingUpgradeMode {
		if v != nil {
			return *v
		}
		var ret RollingUpgradeMode
		return ret
	}).(RollingUpgradeModeOutput)
}

func (o RollingUpgradeModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RollingUpgradeModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RollingUpgradeMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RollingUpgradeModeInput interface {
	pulumi.Input

	ToRollingUpgradeModeOutput() RollingUpgradeModeOutput
	ToRollingUpgradeModeOutputWithContext(context.Context) RollingUpgradeModeOutput
}

var rollingUpgradeModePtrType = reflect.TypeOf((**RollingUpgradeMode)(nil)).Elem()

type RollingUpgradeModePtrInput interface {
	pulumi.Input

	ToRollingUpgradeModePtrOutput() RollingUpgradeModePtrOutput
	ToRollingUpgradeModePtrOutputWithContext(context.Context) RollingUpgradeModePtrOutput
}

type rollingUpgradeModePtr string

func RollingUpgradeModePtr(v string) RollingUpgradeModePtrInput {
	return (*rollingUpgradeModePtr)(&v)
}

func (*rollingUpgradeModePtr) ElementType() reflect.Type {
	return rollingUpgradeModePtrType
}

func (in *rollingUpgradeModePtr) ToRollingUpgradeModePtrOutput() RollingUpgradeModePtrOutput {
	return pulumi.ToOutput(in).(RollingUpgradeModePtrOutput)
}

func (in *rollingUpgradeModePtr) ToRollingUpgradeModePtrOutputWithContext(ctx context.Context) RollingUpgradeModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RollingUpgradeModePtrOutput)
}

type ServiceCorrelationScheme string

const (
	// Aligned affinity ensures that the primaries of the partitions of the affinitized services are collocated on the same nodes. This is the default and is the same as selecting the Affinity scheme. The value is 0.
	ServiceCorrelationSchemeAlignedAffinity = ServiceCorrelationScheme("AlignedAffinity")
	// Non-Aligned affinity guarantees that all replicas of each service will be placed on the same nodes. Unlike Aligned Affinity, this does not guarantee that replicas of particular role will be collocated. The value is 1.
	ServiceCorrelationSchemeNonAlignedAffinity = ServiceCorrelationScheme("NonAlignedAffinity")
)

func (ServiceCorrelationScheme) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorrelationScheme)(nil)).Elem()
}

func (e ServiceCorrelationScheme) ToServiceCorrelationSchemeOutput() ServiceCorrelationSchemeOutput {
	return pulumi.ToOutput(e).(ServiceCorrelationSchemeOutput)
}

func (e ServiceCorrelationScheme) ToServiceCorrelationSchemeOutputWithContext(ctx context.Context) ServiceCorrelationSchemeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceCorrelationSchemeOutput)
}

func (e ServiceCorrelationScheme) ToServiceCorrelationSchemePtrOutput() ServiceCorrelationSchemePtrOutput {
	return e.ToServiceCorrelationSchemePtrOutputWithContext(context.Background())
}

func (e ServiceCorrelationScheme) ToServiceCorrelationSchemePtrOutputWithContext(ctx context.Context) ServiceCorrelationSchemePtrOutput {
	return ServiceCorrelationScheme(e).ToServiceCorrelationSchemeOutputWithContext(ctx).ToServiceCorrelationSchemePtrOutputWithContext(ctx)
}

func (e ServiceCorrelationScheme) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceCorrelationScheme) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceCorrelationScheme) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceCorrelationScheme) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceCorrelationSchemeOutput struct{ *pulumi.OutputState }

func (ServiceCorrelationSchemeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorrelationScheme)(nil)).Elem()
}

func (o ServiceCorrelationSchemeOutput) ToServiceCorrelationSchemeOutput() ServiceCorrelationSchemeOutput {
	return o
}

func (o ServiceCorrelationSchemeOutput) ToServiceCorrelationSchemeOutputWithContext(ctx context.Context) ServiceCorrelationSchemeOutput {
	return o
}

func (o ServiceCorrelationSchemeOutput) ToServiceCorrelationSchemePtrOutput() ServiceCorrelationSchemePtrOutput {
	return o.ToServiceCorrelationSchemePtrOutputWithContext(context.Background())
}

func (o ServiceCorrelationSchemeOutput) ToServiceCorrelationSchemePtrOutputWithContext(ctx context.Context) ServiceCorrelationSchemePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceCorrelationScheme) *ServiceCorrelationScheme {
		return &v
	}).(ServiceCorrelationSchemePtrOutput)
}

func (o ServiceCorrelationSchemeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceCorrelationSchemeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceCorrelationScheme) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceCorrelationSchemeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceCorrelationSchemeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceCorrelationScheme) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceCorrelationSchemePtrOutput struct{ *pulumi.OutputState }

func (ServiceCorrelationSchemePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorrelationScheme)(nil)).Elem()
}

func (o ServiceCorrelationSchemePtrOutput) ToServiceCorrelationSchemePtrOutput() ServiceCorrelationSchemePtrOutput {
	return o
}

func (o ServiceCorrelationSchemePtrOutput) ToServiceCorrelationSchemePtrOutputWithContext(ctx context.Context) ServiceCorrelationSchemePtrOutput {
	return o
}

func (o ServiceCorrelationSchemePtrOutput) Elem() ServiceCorrelationSchemeOutput {
	return o.ApplyT(func(v *ServiceCorrelationScheme) ServiceCorrelationScheme {
		if v != nil {
			return *v
		}
		var ret ServiceCorrelationScheme
		return ret
	}).(ServiceCorrelationSchemeOutput)
}

func (o ServiceCorrelationSchemePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceCorrelationSchemePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceCorrelationScheme) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServiceCorrelationSchemeInput interface {
	pulumi.Input

	ToServiceCorrelationSchemeOutput() ServiceCorrelationSchemeOutput
	ToServiceCorrelationSchemeOutputWithContext(context.Context) ServiceCorrelationSchemeOutput
}

var serviceCorrelationSchemePtrType = reflect.TypeOf((**ServiceCorrelationScheme)(nil)).Elem()

type ServiceCorrelationSchemePtrInput interface {
	pulumi.Input

	ToServiceCorrelationSchemePtrOutput() ServiceCorrelationSchemePtrOutput
	ToServiceCorrelationSchemePtrOutputWithContext(context.Context) ServiceCorrelationSchemePtrOutput
}

type serviceCorrelationSchemePtr string

func ServiceCorrelationSchemePtr(v string) ServiceCorrelationSchemePtrInput {
	return (*serviceCorrelationSchemePtr)(&v)
}

func (*serviceCorrelationSchemePtr) ElementType() reflect.Type {
	return serviceCorrelationSchemePtrType
}

func (in *serviceCorrelationSchemePtr) ToServiceCorrelationSchemePtrOutput() ServiceCorrelationSchemePtrOutput {
	return pulumi.ToOutput(in).(ServiceCorrelationSchemePtrOutput)
}

func (in *serviceCorrelationSchemePtr) ToServiceCorrelationSchemePtrOutputWithContext(ctx context.Context) ServiceCorrelationSchemePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceCorrelationSchemePtrOutput)
}

type ServiceKind string

const (
	// Does not use Service Fabric to make its state highly available or reliable. The value is 0.
	ServiceKindStateless = ServiceKind("Stateless")
	// Uses Service Fabric to make its state or part of its state highly available and reliable. The value is 1.
	ServiceKindStateful = ServiceKind("Stateful")
)

func (ServiceKind) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceKind)(nil)).Elem()
}

func (e ServiceKind) ToServiceKindOutput() ServiceKindOutput {
	return pulumi.ToOutput(e).(ServiceKindOutput)
}

func (e ServiceKind) ToServiceKindOutputWithContext(ctx context.Context) ServiceKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceKindOutput)
}

func (e ServiceKind) ToServiceKindPtrOutput() ServiceKindPtrOutput {
	return e.ToServiceKindPtrOutputWithContext(context.Background())
}

func (e ServiceKind) ToServiceKindPtrOutputWithContext(ctx context.Context) ServiceKindPtrOutput {
	return ServiceKind(e).ToServiceKindOutputWithContext(ctx).ToServiceKindPtrOutputWithContext(ctx)
}

func (e ServiceKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceKindOutput struct{ *pulumi.OutputState }

func (ServiceKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceKind)(nil)).Elem()
}

func (o ServiceKindOutput) ToServiceKindOutput() ServiceKindOutput {
	return o
}

func (o ServiceKindOutput) ToServiceKindOutputWithContext(ctx context.Context) ServiceKindOutput {
	return o
}

func (o ServiceKindOutput) ToServiceKindPtrOutput() ServiceKindPtrOutput {
	return o.ToServiceKindPtrOutputWithContext(context.Background())
}

func (o ServiceKindOutput) ToServiceKindPtrOutputWithContext(ctx context.Context) ServiceKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceKind) *ServiceKind {
		return &v
	}).(ServiceKindPtrOutput)
}

func (o ServiceKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceKindPtrOutput struct{ *pulumi.OutputState }

func (ServiceKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceKind)(nil)).Elem()
}

func (o ServiceKindPtrOutput) ToServiceKindPtrOutput() ServiceKindPtrOutput {
	return o
}

func (o ServiceKindPtrOutput) ToServiceKindPtrOutputWithContext(ctx context.Context) ServiceKindPtrOutput {
	return o
}

func (o ServiceKindPtrOutput) Elem() ServiceKindOutput {
	return o.ApplyT(func(v *ServiceKind) ServiceKind {
		if v != nil {
			return *v
		}
		var ret ServiceKind
		return ret
	}).(ServiceKindOutput)
}

func (o ServiceKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServiceKindInput interface {
	pulumi.Input

	ToServiceKindOutput() ServiceKindOutput
	ToServiceKindOutputWithContext(context.Context) ServiceKindOutput
}

var serviceKindPtrType = reflect.TypeOf((**ServiceKind)(nil)).Elem()

type ServiceKindPtrInput interface {
	pulumi.Input

	ToServiceKindPtrOutput() ServiceKindPtrOutput
	ToServiceKindPtrOutputWithContext(context.Context) ServiceKindPtrOutput
}

type serviceKindPtr string

func ServiceKindPtr(v string) ServiceKindPtrInput {
	return (*serviceKindPtr)(&v)
}

func (*serviceKindPtr) ElementType() reflect.Type {
	return serviceKindPtrType
}

func (in *serviceKindPtr) ToServiceKindPtrOutput() ServiceKindPtrOutput {
	return pulumi.ToOutput(in).(ServiceKindPtrOutput)
}

func (in *serviceKindPtr) ToServiceKindPtrOutputWithContext(ctx context.Context) ServiceKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceKindPtrOutput)
}

type ServiceLoadMetricWeight string

const (
	// Disables resource balancing for this metric. This value is zero.
	ServiceLoadMetricWeightZero = ServiceLoadMetricWeight("Zero")
	// Specifies the metric weight of the service load as Low. The value is 1.
	ServiceLoadMetricWeightLow = ServiceLoadMetricWeight("Low")
	// Specifies the metric weight of the service load as Medium. The value is 2.
	ServiceLoadMetricWeightMedium = ServiceLoadMetricWeight("Medium")
	// Specifies the metric weight of the service load as High. The value is 3.
	ServiceLoadMetricWeightHigh = ServiceLoadMetricWeight("High")
)

func (ServiceLoadMetricWeight) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLoadMetricWeight)(nil)).Elem()
}

func (e ServiceLoadMetricWeight) ToServiceLoadMetricWeightOutput() ServiceLoadMetricWeightOutput {
	return pulumi.ToOutput(e).(ServiceLoadMetricWeightOutput)
}

func (e ServiceLoadMetricWeight) ToServiceLoadMetricWeightOutputWithContext(ctx context.Context) ServiceLoadMetricWeightOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceLoadMetricWeightOutput)
}

func (e ServiceLoadMetricWeight) ToServiceLoadMetricWeightPtrOutput() ServiceLoadMetricWeightPtrOutput {
	return e.ToServiceLoadMetricWeightPtrOutputWithContext(context.Background())
}

func (e ServiceLoadMetricWeight) ToServiceLoadMetricWeightPtrOutputWithContext(ctx context.Context) ServiceLoadMetricWeightPtrOutput {
	return ServiceLoadMetricWeight(e).ToServiceLoadMetricWeightOutputWithContext(ctx).ToServiceLoadMetricWeightPtrOutputWithContext(ctx)
}

func (e ServiceLoadMetricWeight) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceLoadMetricWeight) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceLoadMetricWeight) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceLoadMetricWeight) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceLoadMetricWeightOutput struct{ *pulumi.OutputState }

func (ServiceLoadMetricWeightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceLoadMetricWeight)(nil)).Elem()
}

func (o ServiceLoadMetricWeightOutput) ToServiceLoadMetricWeightOutput() ServiceLoadMetricWeightOutput {
	return o
}

func (o ServiceLoadMetricWeightOutput) ToServiceLoadMetricWeightOutputWithContext(ctx context.Context) ServiceLoadMetricWeightOutput {
	return o
}

func (o ServiceLoadMetricWeightOutput) ToServiceLoadMetricWeightPtrOutput() ServiceLoadMetricWeightPtrOutput {
	return o.ToServiceLoadMetricWeightPtrOutputWithContext(context.Background())
}

func (o ServiceLoadMetricWeightOutput) ToServiceLoadMetricWeightPtrOutputWithContext(ctx context.Context) ServiceLoadMetricWeightPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceLoadMetricWeight) *ServiceLoadMetricWeight {
		return &v
	}).(ServiceLoadMetricWeightPtrOutput)
}

func (o ServiceLoadMetricWeightOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceLoadMetricWeightOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceLoadMetricWeight) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceLoadMetricWeightOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceLoadMetricWeightOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceLoadMetricWeight) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceLoadMetricWeightPtrOutput struct{ *pulumi.OutputState }

func (ServiceLoadMetricWeightPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceLoadMetricWeight)(nil)).Elem()
}

func (o ServiceLoadMetricWeightPtrOutput) ToServiceLoadMetricWeightPtrOutput() ServiceLoadMetricWeightPtrOutput {
	return o
}

func (o ServiceLoadMetricWeightPtrOutput) ToServiceLoadMetricWeightPtrOutputWithContext(ctx context.Context) ServiceLoadMetricWeightPtrOutput {
	return o
}

func (o ServiceLoadMetricWeightPtrOutput) Elem() ServiceLoadMetricWeightOutput {
	return o.ApplyT(func(v *ServiceLoadMetricWeight) ServiceLoadMetricWeight {
		if v != nil {
			return *v
		}
		var ret ServiceLoadMetricWeight
		return ret
	}).(ServiceLoadMetricWeightOutput)
}

func (o ServiceLoadMetricWeightPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceLoadMetricWeightPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceLoadMetricWeight) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServiceLoadMetricWeightInput interface {
	pulumi.Input

	ToServiceLoadMetricWeightOutput() ServiceLoadMetricWeightOutput
	ToServiceLoadMetricWeightOutputWithContext(context.Context) ServiceLoadMetricWeightOutput
}

var serviceLoadMetricWeightPtrType = reflect.TypeOf((**ServiceLoadMetricWeight)(nil)).Elem()

type ServiceLoadMetricWeightPtrInput interface {
	pulumi.Input

	ToServiceLoadMetricWeightPtrOutput() ServiceLoadMetricWeightPtrOutput
	ToServiceLoadMetricWeightPtrOutputWithContext(context.Context) ServiceLoadMetricWeightPtrOutput
}

type serviceLoadMetricWeightPtr string

func ServiceLoadMetricWeightPtr(v string) ServiceLoadMetricWeightPtrInput {
	return (*serviceLoadMetricWeightPtr)(&v)
}

func (*serviceLoadMetricWeightPtr) ElementType() reflect.Type {
	return serviceLoadMetricWeightPtrType
}

func (in *serviceLoadMetricWeightPtr) ToServiceLoadMetricWeightPtrOutput() ServiceLoadMetricWeightPtrOutput {
	return pulumi.ToOutput(in).(ServiceLoadMetricWeightPtrOutput)
}

func (in *serviceLoadMetricWeightPtr) ToServiceLoadMetricWeightPtrOutputWithContext(ctx context.Context) ServiceLoadMetricWeightPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceLoadMetricWeightPtrOutput)
}

type ServicePackageActivationMode string

const (
	// Indicates the application package activation mode will use shared process.
	ServicePackageActivationModeSharedProcess = ServicePackageActivationMode("SharedProcess")
	// Indicates the application package activation mode will use exclusive process.
	ServicePackageActivationModeExclusiveProcess = ServicePackageActivationMode("ExclusiveProcess")
)

func (ServicePackageActivationMode) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePackageActivationMode)(nil)).Elem()
}

func (e ServicePackageActivationMode) ToServicePackageActivationModeOutput() ServicePackageActivationModeOutput {
	return pulumi.ToOutput(e).(ServicePackageActivationModeOutput)
}

func (e ServicePackageActivationMode) ToServicePackageActivationModeOutputWithContext(ctx context.Context) ServicePackageActivationModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServicePackageActivationModeOutput)
}

func (e ServicePackageActivationMode) ToServicePackageActivationModePtrOutput() ServicePackageActivationModePtrOutput {
	return e.ToServicePackageActivationModePtrOutputWithContext(context.Background())
}

func (e ServicePackageActivationMode) ToServicePackageActivationModePtrOutputWithContext(ctx context.Context) ServicePackageActivationModePtrOutput {
	return ServicePackageActivationMode(e).ToServicePackageActivationModeOutputWithContext(ctx).ToServicePackageActivationModePtrOutputWithContext(ctx)
}

func (e ServicePackageActivationMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServicePackageActivationMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServicePackageActivationMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServicePackageActivationMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServicePackageActivationModeOutput struct{ *pulumi.OutputState }

func (ServicePackageActivationModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePackageActivationMode)(nil)).Elem()
}

func (o ServicePackageActivationModeOutput) ToServicePackageActivationModeOutput() ServicePackageActivationModeOutput {
	return o
}

func (o ServicePackageActivationModeOutput) ToServicePackageActivationModeOutputWithContext(ctx context.Context) ServicePackageActivationModeOutput {
	return o
}

func (o ServicePackageActivationModeOutput) ToServicePackageActivationModePtrOutput() ServicePackageActivationModePtrOutput {
	return o.ToServicePackageActivationModePtrOutputWithContext(context.Background())
}

func (o ServicePackageActivationModeOutput) ToServicePackageActivationModePtrOutputWithContext(ctx context.Context) ServicePackageActivationModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePackageActivationMode) *ServicePackageActivationMode {
		return &v
	}).(ServicePackageActivationModePtrOutput)
}

func (o ServicePackageActivationModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServicePackageActivationModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServicePackageActivationMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServicePackageActivationModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServicePackageActivationModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServicePackageActivationMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServicePackageActivationModePtrOutput struct{ *pulumi.OutputState }

func (ServicePackageActivationModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePackageActivationMode)(nil)).Elem()
}

func (o ServicePackageActivationModePtrOutput) ToServicePackageActivationModePtrOutput() ServicePackageActivationModePtrOutput {
	return o
}

func (o ServicePackageActivationModePtrOutput) ToServicePackageActivationModePtrOutputWithContext(ctx context.Context) ServicePackageActivationModePtrOutput {
	return o
}

func (o ServicePackageActivationModePtrOutput) Elem() ServicePackageActivationModeOutput {
	return o.ApplyT(func(v *ServicePackageActivationMode) ServicePackageActivationMode {
		if v != nil {
			return *v
		}
		var ret ServicePackageActivationMode
		return ret
	}).(ServicePackageActivationModeOutput)
}

func (o ServicePackageActivationModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServicePackageActivationModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServicePackageActivationMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServicePackageActivationModeInput interface {
	pulumi.Input

	ToServicePackageActivationModeOutput() ServicePackageActivationModeOutput
	ToServicePackageActivationModeOutputWithContext(context.Context) ServicePackageActivationModeOutput
}

var servicePackageActivationModePtrType = reflect.TypeOf((**ServicePackageActivationMode)(nil)).Elem()

type ServicePackageActivationModePtrInput interface {
	pulumi.Input

	ToServicePackageActivationModePtrOutput() ServicePackageActivationModePtrOutput
	ToServicePackageActivationModePtrOutputWithContext(context.Context) ServicePackageActivationModePtrOutput
}

type servicePackageActivationModePtr string

func ServicePackageActivationModePtr(v string) ServicePackageActivationModePtrInput {
	return (*servicePackageActivationModePtr)(&v)
}

func (*servicePackageActivationModePtr) ElementType() reflect.Type {
	return servicePackageActivationModePtrType
}

func (in *servicePackageActivationModePtr) ToServicePackageActivationModePtrOutput() ServicePackageActivationModePtrOutput {
	return pulumi.ToOutput(in).(ServicePackageActivationModePtrOutput)
}

func (in *servicePackageActivationModePtr) ToServicePackageActivationModePtrOutputWithContext(ctx context.Context) ServicePackageActivationModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServicePackageActivationModePtrOutput)
}

type ServicePlacementPolicyType string

const (
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementInvalidDomainPolicyDescription, which indicates that a particular fault or upgrade domain cannot be used for placement of this service. The value is 0.
	ServicePlacementPolicyTypeInvalidDomain = ServicePlacementPolicyType("InvalidDomain")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription indicating that the replicas of the service must be placed in a specific domain. The value is 1.
	ServicePlacementPolicyTypeRequiredDomain = ServicePlacementPolicyType("RequiredDomain")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementPreferPrimaryDomainPolicyDescription, which indicates that if possible the Primary replica for the partitions of the service should be located in a particular domain as an optimization. The value is 2.
	ServicePlacementPolicyTypePreferredPrimaryDomain = ServicePlacementPolicyType("PreferredPrimaryDomain")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementRequireDomainDistributionPolicyDescription, indicating that the system will disallow placement of any two replicas from the same partition in the same domain at any time. The value is 3.
	ServicePlacementPolicyTypeRequiredDomainDistribution = ServicePlacementPolicyType("RequiredDomainDistribution")
	// Indicates that the ServicePlacementPolicyDescription is of type ServicePlacementNonPartiallyPlaceServicePolicyDescription, which indicates that if possible all replicas of a particular partition of the service should be placed atomically. The value is 4.
	ServicePlacementPolicyTypeNonPartiallyPlaceService = ServicePlacementPolicyType("NonPartiallyPlaceService")
)

func (ServicePlacementPolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementPolicyType)(nil)).Elem()
}

func (e ServicePlacementPolicyType) ToServicePlacementPolicyTypeOutput() ServicePlacementPolicyTypeOutput {
	return pulumi.ToOutput(e).(ServicePlacementPolicyTypeOutput)
}

func (e ServicePlacementPolicyType) ToServicePlacementPolicyTypeOutputWithContext(ctx context.Context) ServicePlacementPolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServicePlacementPolicyTypeOutput)
}

func (e ServicePlacementPolicyType) ToServicePlacementPolicyTypePtrOutput() ServicePlacementPolicyTypePtrOutput {
	return e.ToServicePlacementPolicyTypePtrOutputWithContext(context.Background())
}

func (e ServicePlacementPolicyType) ToServicePlacementPolicyTypePtrOutputWithContext(ctx context.Context) ServicePlacementPolicyTypePtrOutput {
	return ServicePlacementPolicyType(e).ToServicePlacementPolicyTypeOutputWithContext(ctx).ToServicePlacementPolicyTypePtrOutputWithContext(ctx)
}

func (e ServicePlacementPolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServicePlacementPolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServicePlacementPolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServicePlacementPolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServicePlacementPolicyTypeOutput struct{ *pulumi.OutputState }

func (ServicePlacementPolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicePlacementPolicyType)(nil)).Elem()
}

func (o ServicePlacementPolicyTypeOutput) ToServicePlacementPolicyTypeOutput() ServicePlacementPolicyTypeOutput {
	return o
}

func (o ServicePlacementPolicyTypeOutput) ToServicePlacementPolicyTypeOutputWithContext(ctx context.Context) ServicePlacementPolicyTypeOutput {
	return o
}

func (o ServicePlacementPolicyTypeOutput) ToServicePlacementPolicyTypePtrOutput() ServicePlacementPolicyTypePtrOutput {
	return o.ToServicePlacementPolicyTypePtrOutputWithContext(context.Background())
}

func (o ServicePlacementPolicyTypeOutput) ToServicePlacementPolicyTypePtrOutputWithContext(ctx context.Context) ServicePlacementPolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicePlacementPolicyType) *ServicePlacementPolicyType {
		return &v
	}).(ServicePlacementPolicyTypePtrOutput)
}

func (o ServicePlacementPolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServicePlacementPolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServicePlacementPolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServicePlacementPolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServicePlacementPolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServicePlacementPolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServicePlacementPolicyTypePtrOutput struct{ *pulumi.OutputState }

func (ServicePlacementPolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicePlacementPolicyType)(nil)).Elem()
}

func (o ServicePlacementPolicyTypePtrOutput) ToServicePlacementPolicyTypePtrOutput() ServicePlacementPolicyTypePtrOutput {
	return o
}

func (o ServicePlacementPolicyTypePtrOutput) ToServicePlacementPolicyTypePtrOutputWithContext(ctx context.Context) ServicePlacementPolicyTypePtrOutput {
	return o
}

func (o ServicePlacementPolicyTypePtrOutput) Elem() ServicePlacementPolicyTypeOutput {
	return o.ApplyT(func(v *ServicePlacementPolicyType) ServicePlacementPolicyType {
		if v != nil {
			return *v
		}
		var ret ServicePlacementPolicyType
		return ret
	}).(ServicePlacementPolicyTypeOutput)
}

func (o ServicePlacementPolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServicePlacementPolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServicePlacementPolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServicePlacementPolicyTypeInput interface {
	pulumi.Input

	ToServicePlacementPolicyTypeOutput() ServicePlacementPolicyTypeOutput
	ToServicePlacementPolicyTypeOutputWithContext(context.Context) ServicePlacementPolicyTypeOutput
}

var servicePlacementPolicyTypePtrType = reflect.TypeOf((**ServicePlacementPolicyType)(nil)).Elem()

type ServicePlacementPolicyTypePtrInput interface {
	pulumi.Input

	ToServicePlacementPolicyTypePtrOutput() ServicePlacementPolicyTypePtrOutput
	ToServicePlacementPolicyTypePtrOutputWithContext(context.Context) ServicePlacementPolicyTypePtrOutput
}

type servicePlacementPolicyTypePtr string

func ServicePlacementPolicyTypePtr(v string) ServicePlacementPolicyTypePtrInput {
	return (*servicePlacementPolicyTypePtr)(&v)
}

func (*servicePlacementPolicyTypePtr) ElementType() reflect.Type {
	return servicePlacementPolicyTypePtrType
}

func (in *servicePlacementPolicyTypePtr) ToServicePlacementPolicyTypePtrOutput() ServicePlacementPolicyTypePtrOutput {
	return pulumi.ToOutput(in).(ServicePlacementPolicyTypePtrOutput)
}

func (in *servicePlacementPolicyTypePtr) ToServicePlacementPolicyTypePtrOutputWithContext(ctx context.Context) ServicePlacementPolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServicePlacementPolicyTypePtrOutput)
}

type ServiceScalingMechanismKind string

const (
	// Represents a scaling mechanism for adding or removing instances of stateless service partition. The value is 0.
	ServiceScalingMechanismKindScalePartitionInstanceCount = ServiceScalingMechanismKind("ScalePartitionInstanceCount")
	// Represents a scaling mechanism for adding or removing named partitions of a stateless service. The value is 1.
	ServiceScalingMechanismKindAddRemoveIncrementalNamedPartition = ServiceScalingMechanismKind("AddRemoveIncrementalNamedPartition")
)

func (ServiceScalingMechanismKind) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceScalingMechanismKind)(nil)).Elem()
}

func (e ServiceScalingMechanismKind) ToServiceScalingMechanismKindOutput() ServiceScalingMechanismKindOutput {
	return pulumi.ToOutput(e).(ServiceScalingMechanismKindOutput)
}

func (e ServiceScalingMechanismKind) ToServiceScalingMechanismKindOutputWithContext(ctx context.Context) ServiceScalingMechanismKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceScalingMechanismKindOutput)
}

func (e ServiceScalingMechanismKind) ToServiceScalingMechanismKindPtrOutput() ServiceScalingMechanismKindPtrOutput {
	return e.ToServiceScalingMechanismKindPtrOutputWithContext(context.Background())
}

func (e ServiceScalingMechanismKind) ToServiceScalingMechanismKindPtrOutputWithContext(ctx context.Context) ServiceScalingMechanismKindPtrOutput {
	return ServiceScalingMechanismKind(e).ToServiceScalingMechanismKindOutputWithContext(ctx).ToServiceScalingMechanismKindPtrOutputWithContext(ctx)
}

func (e ServiceScalingMechanismKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceScalingMechanismKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceScalingMechanismKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceScalingMechanismKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceScalingMechanismKindOutput struct{ *pulumi.OutputState }

func (ServiceScalingMechanismKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceScalingMechanismKind)(nil)).Elem()
}

func (o ServiceScalingMechanismKindOutput) ToServiceScalingMechanismKindOutput() ServiceScalingMechanismKindOutput {
	return o
}

func (o ServiceScalingMechanismKindOutput) ToServiceScalingMechanismKindOutputWithContext(ctx context.Context) ServiceScalingMechanismKindOutput {
	return o
}

func (o ServiceScalingMechanismKindOutput) ToServiceScalingMechanismKindPtrOutput() ServiceScalingMechanismKindPtrOutput {
	return o.ToServiceScalingMechanismKindPtrOutputWithContext(context.Background())
}

func (o ServiceScalingMechanismKindOutput) ToServiceScalingMechanismKindPtrOutputWithContext(ctx context.Context) ServiceScalingMechanismKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceScalingMechanismKind) *ServiceScalingMechanismKind {
		return &v
	}).(ServiceScalingMechanismKindPtrOutput)
}

func (o ServiceScalingMechanismKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceScalingMechanismKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceScalingMechanismKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceScalingMechanismKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceScalingMechanismKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceScalingMechanismKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceScalingMechanismKindPtrOutput struct{ *pulumi.OutputState }

func (ServiceScalingMechanismKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceScalingMechanismKind)(nil)).Elem()
}

func (o ServiceScalingMechanismKindPtrOutput) ToServiceScalingMechanismKindPtrOutput() ServiceScalingMechanismKindPtrOutput {
	return o
}

func (o ServiceScalingMechanismKindPtrOutput) ToServiceScalingMechanismKindPtrOutputWithContext(ctx context.Context) ServiceScalingMechanismKindPtrOutput {
	return o
}

func (o ServiceScalingMechanismKindPtrOutput) Elem() ServiceScalingMechanismKindOutput {
	return o.ApplyT(func(v *ServiceScalingMechanismKind) ServiceScalingMechanismKind {
		if v != nil {
			return *v
		}
		var ret ServiceScalingMechanismKind
		return ret
	}).(ServiceScalingMechanismKindOutput)
}

func (o ServiceScalingMechanismKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceScalingMechanismKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceScalingMechanismKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServiceScalingMechanismKindInput interface {
	pulumi.Input

	ToServiceScalingMechanismKindOutput() ServiceScalingMechanismKindOutput
	ToServiceScalingMechanismKindOutputWithContext(context.Context) ServiceScalingMechanismKindOutput
}

var serviceScalingMechanismKindPtrType = reflect.TypeOf((**ServiceScalingMechanismKind)(nil)).Elem()

type ServiceScalingMechanismKindPtrInput interface {
	pulumi.Input

	ToServiceScalingMechanismKindPtrOutput() ServiceScalingMechanismKindPtrOutput
	ToServiceScalingMechanismKindPtrOutputWithContext(context.Context) ServiceScalingMechanismKindPtrOutput
}

type serviceScalingMechanismKindPtr string

func ServiceScalingMechanismKindPtr(v string) ServiceScalingMechanismKindPtrInput {
	return (*serviceScalingMechanismKindPtr)(&v)
}

func (*serviceScalingMechanismKindPtr) ElementType() reflect.Type {
	return serviceScalingMechanismKindPtrType
}

func (in *serviceScalingMechanismKindPtr) ToServiceScalingMechanismKindPtrOutput() ServiceScalingMechanismKindPtrOutput {
	return pulumi.ToOutput(in).(ServiceScalingMechanismKindPtrOutput)
}

func (in *serviceScalingMechanismKindPtr) ToServiceScalingMechanismKindPtrOutputWithContext(ctx context.Context) ServiceScalingMechanismKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceScalingMechanismKindPtrOutput)
}

type ServiceScalingTriggerKind string

const (
	// Represents a scaling trigger related to an average load of a metric/resource of a partition. The value is 0.
	ServiceScalingTriggerKindAveragePartitionLoad = ServiceScalingTriggerKind("AveragePartitionLoad")
	// Represents a scaling policy related to an average load of a metric/resource of a service. The value is 1.
	ServiceScalingTriggerKindAverageServiceLoad = ServiceScalingTriggerKind("AverageServiceLoad")
)

func (ServiceScalingTriggerKind) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceScalingTriggerKind)(nil)).Elem()
}

func (e ServiceScalingTriggerKind) ToServiceScalingTriggerKindOutput() ServiceScalingTriggerKindOutput {
	return pulumi.ToOutput(e).(ServiceScalingTriggerKindOutput)
}

func (e ServiceScalingTriggerKind) ToServiceScalingTriggerKindOutputWithContext(ctx context.Context) ServiceScalingTriggerKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServiceScalingTriggerKindOutput)
}

func (e ServiceScalingTriggerKind) ToServiceScalingTriggerKindPtrOutput() ServiceScalingTriggerKindPtrOutput {
	return e.ToServiceScalingTriggerKindPtrOutputWithContext(context.Background())
}

func (e ServiceScalingTriggerKind) ToServiceScalingTriggerKindPtrOutputWithContext(ctx context.Context) ServiceScalingTriggerKindPtrOutput {
	return ServiceScalingTriggerKind(e).ToServiceScalingTriggerKindOutputWithContext(ctx).ToServiceScalingTriggerKindPtrOutputWithContext(ctx)
}

func (e ServiceScalingTriggerKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceScalingTriggerKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServiceScalingTriggerKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServiceScalingTriggerKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServiceScalingTriggerKindOutput struct{ *pulumi.OutputState }

func (ServiceScalingTriggerKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceScalingTriggerKind)(nil)).Elem()
}

func (o ServiceScalingTriggerKindOutput) ToServiceScalingTriggerKindOutput() ServiceScalingTriggerKindOutput {
	return o
}

func (o ServiceScalingTriggerKindOutput) ToServiceScalingTriggerKindOutputWithContext(ctx context.Context) ServiceScalingTriggerKindOutput {
	return o
}

func (o ServiceScalingTriggerKindOutput) ToServiceScalingTriggerKindPtrOutput() ServiceScalingTriggerKindPtrOutput {
	return o.ToServiceScalingTriggerKindPtrOutputWithContext(context.Background())
}

func (o ServiceScalingTriggerKindOutput) ToServiceScalingTriggerKindPtrOutputWithContext(ctx context.Context) ServiceScalingTriggerKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceScalingTriggerKind) *ServiceScalingTriggerKind {
		return &v
	}).(ServiceScalingTriggerKindPtrOutput)
}

func (o ServiceScalingTriggerKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServiceScalingTriggerKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceScalingTriggerKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServiceScalingTriggerKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceScalingTriggerKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServiceScalingTriggerKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServiceScalingTriggerKindPtrOutput struct{ *pulumi.OutputState }

func (ServiceScalingTriggerKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceScalingTriggerKind)(nil)).Elem()
}

func (o ServiceScalingTriggerKindPtrOutput) ToServiceScalingTriggerKindPtrOutput() ServiceScalingTriggerKindPtrOutput {
	return o
}

func (o ServiceScalingTriggerKindPtrOutput) ToServiceScalingTriggerKindPtrOutputWithContext(ctx context.Context) ServiceScalingTriggerKindPtrOutput {
	return o
}

func (o ServiceScalingTriggerKindPtrOutput) Elem() ServiceScalingTriggerKindOutput {
	return o.ApplyT(func(v *ServiceScalingTriggerKind) ServiceScalingTriggerKind {
		if v != nil {
			return *v
		}
		var ret ServiceScalingTriggerKind
		return ret
	}).(ServiceScalingTriggerKindOutput)
}

func (o ServiceScalingTriggerKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServiceScalingTriggerKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServiceScalingTriggerKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServiceScalingTriggerKindInput interface {
	pulumi.Input

	ToServiceScalingTriggerKindOutput() ServiceScalingTriggerKindOutput
	ToServiceScalingTriggerKindOutputWithContext(context.Context) ServiceScalingTriggerKindOutput
}

var serviceScalingTriggerKindPtrType = reflect.TypeOf((**ServiceScalingTriggerKind)(nil)).Elem()

type ServiceScalingTriggerKindPtrInput interface {
	pulumi.Input

	ToServiceScalingTriggerKindPtrOutput() ServiceScalingTriggerKindPtrOutput
	ToServiceScalingTriggerKindPtrOutputWithContext(context.Context) ServiceScalingTriggerKindPtrOutput
}

type serviceScalingTriggerKindPtr string

func ServiceScalingTriggerKindPtr(v string) ServiceScalingTriggerKindPtrInput {
	return (*serviceScalingTriggerKindPtr)(&v)
}

func (*serviceScalingTriggerKindPtr) ElementType() reflect.Type {
	return serviceScalingTriggerKindPtrType
}

func (in *serviceScalingTriggerKindPtr) ToServiceScalingTriggerKindPtrOutput() ServiceScalingTriggerKindPtrOutput {
	return pulumi.ToOutput(in).(ServiceScalingTriggerKindPtrOutput)
}

func (in *serviceScalingTriggerKindPtr) ToServiceScalingTriggerKindPtrOutputWithContext(ctx context.Context) ServiceScalingTriggerKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServiceScalingTriggerKindPtrOutput)
}

type SkuName string

const (
	// Basic requires a minimum of 3 nodes and allows only 1 node type.
	SkuNameBasic = SkuName("Basic")
	// Requires a minimum of 5 nodes and allows 1 or more node type.
	SkuNameStandard = SkuName("Standard")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AccessOutput{})
	pulumi.RegisterOutputType(AccessPtrOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeCadenceOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeCadencePtrOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeModeOutput{})
	pulumi.RegisterOutputType(ClusterUpgradeModePtrOutput{})
	pulumi.RegisterOutputType(DirectionOutput{})
	pulumi.RegisterOutputType(DirectionPtrOutput{})
	pulumi.RegisterOutputType(DiskTypeOutput{})
	pulumi.RegisterOutputType(DiskTypePtrOutput{})
	pulumi.RegisterOutputType(FailureActionOutput{})
	pulumi.RegisterOutputType(FailureActionPtrOutput{})
	pulumi.RegisterOutputType(IPAddressTypeOutput{})
	pulumi.RegisterOutputType(IPAddressTypePtrOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypeOutput{})
	pulumi.RegisterOutputType(ManagedIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(MoveCostOutput{})
	pulumi.RegisterOutputType(MoveCostPtrOutput{})
	pulumi.RegisterOutputType(NsgProtocolOutput{})
	pulumi.RegisterOutputType(NsgProtocolPtrOutput{})
	pulumi.RegisterOutputType(PartitionSchemeOutput{})
	pulumi.RegisterOutputType(PartitionSchemePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointNetworkPoliciesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointNetworkPoliciesPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceNetworkPoliciesOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceNetworkPoliciesPtrOutput{})
	pulumi.RegisterOutputType(ProbeProtocolOutput{})
	pulumi.RegisterOutputType(ProbeProtocolPtrOutput{})
	pulumi.RegisterOutputType(ProtocolOutput{})
	pulumi.RegisterOutputType(ProtocolPtrOutput{})
	pulumi.RegisterOutputType(RollingUpgradeModeOutput{})
	pulumi.RegisterOutputType(RollingUpgradeModePtrOutput{})
	pulumi.RegisterOutputType(ServiceCorrelationSchemeOutput{})
	pulumi.RegisterOutputType(ServiceCorrelationSchemePtrOutput{})
	pulumi.RegisterOutputType(ServiceKindOutput{})
	pulumi.RegisterOutputType(ServiceKindPtrOutput{})
	pulumi.RegisterOutputType(ServiceLoadMetricWeightOutput{})
	pulumi.RegisterOutputType(ServiceLoadMetricWeightPtrOutput{})
	pulumi.RegisterOutputType(ServicePackageActivationModeOutput{})
	pulumi.RegisterOutputType(ServicePackageActivationModePtrOutput{})
	pulumi.RegisterOutputType(ServicePlacementPolicyTypeOutput{})
	pulumi.RegisterOutputType(ServicePlacementPolicyTypePtrOutput{})
	pulumi.RegisterOutputType(ServiceScalingMechanismKindOutput{})
	pulumi.RegisterOutputType(ServiceScalingMechanismKindPtrOutput{})
	pulumi.RegisterOutputType(ServiceScalingTriggerKindOutput{})
	pulumi.RegisterOutputType(ServiceScalingTriggerKindPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
