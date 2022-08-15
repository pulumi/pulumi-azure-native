


package v20200605preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Checkpoint struct {
	CheckpointID       *string `pulumi:"checkpointID"`
	Description        *string `pulumi:"description"`
	Name               *string `pulumi:"name"`
	ParentCheckpointID *string `pulumi:"parentCheckpointID"`
}





type CheckpointInput interface {
	pulumi.Input

	ToCheckpointOutput() CheckpointOutput
	ToCheckpointOutputWithContext(context.Context) CheckpointOutput
}

type CheckpointArgs struct {
	CheckpointID       pulumi.StringPtrInput `pulumi:"checkpointID"`
	Description        pulumi.StringPtrInput `pulumi:"description"`
	Name               pulumi.StringPtrInput `pulumi:"name"`
	ParentCheckpointID pulumi.StringPtrInput `pulumi:"parentCheckpointID"`
}

func (CheckpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Checkpoint)(nil)).Elem()
}

func (i CheckpointArgs) ToCheckpointOutput() CheckpointOutput {
	return i.ToCheckpointOutputWithContext(context.Background())
}

func (i CheckpointArgs) ToCheckpointOutputWithContext(ctx context.Context) CheckpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckpointOutput)
}





type CheckpointArrayInput interface {
	pulumi.Input

	ToCheckpointArrayOutput() CheckpointArrayOutput
	ToCheckpointArrayOutputWithContext(context.Context) CheckpointArrayOutput
}

type CheckpointArray []CheckpointInput

func (CheckpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Checkpoint)(nil)).Elem()
}

func (i CheckpointArray) ToCheckpointArrayOutput() CheckpointArrayOutput {
	return i.ToCheckpointArrayOutputWithContext(context.Background())
}

func (i CheckpointArray) ToCheckpointArrayOutputWithContext(ctx context.Context) CheckpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckpointArrayOutput)
}

type CheckpointOutput struct{ *pulumi.OutputState }

func (CheckpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Checkpoint)(nil)).Elem()
}

func (o CheckpointOutput) ToCheckpointOutput() CheckpointOutput {
	return o
}

func (o CheckpointOutput) ToCheckpointOutputWithContext(ctx context.Context) CheckpointOutput {
	return o
}

func (o CheckpointOutput) CheckpointID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Checkpoint) *string { return v.CheckpointID }).(pulumi.StringPtrOutput)
}

func (o CheckpointOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Checkpoint) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CheckpointOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Checkpoint) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CheckpointOutput) ParentCheckpointID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Checkpoint) *string { return v.ParentCheckpointID }).(pulumi.StringPtrOutput)
}

type CheckpointArrayOutput struct{ *pulumi.OutputState }

func (CheckpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Checkpoint)(nil)).Elem()
}

func (o CheckpointArrayOutput) ToCheckpointArrayOutput() CheckpointArrayOutput {
	return o
}

func (o CheckpointArrayOutput) ToCheckpointArrayOutputWithContext(ctx context.Context) CheckpointArrayOutput {
	return o
}

func (o CheckpointArrayOutput) Index(i pulumi.IntInput) CheckpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Checkpoint {
		return vs[0].([]Checkpoint)[vs[1].(int)]
	}).(CheckpointOutput)
}

type CheckpointResponse struct {
	CheckpointID       *string `pulumi:"checkpointID"`
	Description        *string `pulumi:"description"`
	Name               *string `pulumi:"name"`
	ParentCheckpointID *string `pulumi:"parentCheckpointID"`
}

type CheckpointResponseOutput struct{ *pulumi.OutputState }

func (CheckpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CheckpointResponse)(nil)).Elem()
}

func (o CheckpointResponseOutput) ToCheckpointResponseOutput() CheckpointResponseOutput {
	return o
}

func (o CheckpointResponseOutput) ToCheckpointResponseOutputWithContext(ctx context.Context) CheckpointResponseOutput {
	return o
}

func (o CheckpointResponseOutput) CheckpointID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CheckpointResponse) *string { return v.CheckpointID }).(pulumi.StringPtrOutput)
}

func (o CheckpointResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CheckpointResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CheckpointResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CheckpointResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CheckpointResponseOutput) ParentCheckpointID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CheckpointResponse) *string { return v.ParentCheckpointID }).(pulumi.StringPtrOutput)
}

type CheckpointResponseArrayOutput struct{ *pulumi.OutputState }

func (CheckpointResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CheckpointResponse)(nil)).Elem()
}

func (o CheckpointResponseArrayOutput) ToCheckpointResponseArrayOutput() CheckpointResponseArrayOutput {
	return o
}

func (o CheckpointResponseArrayOutput) ToCheckpointResponseArrayOutputWithContext(ctx context.Context) CheckpointResponseArrayOutput {
	return o
}

func (o CheckpointResponseArrayOutput) Index(i pulumi.IntInput) CheckpointResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CheckpointResponse {
		return vs[0].([]CheckpointResponse)[vs[1].(int)]
	}).(CheckpointResponseOutput)
}

type CloudCapacityResponse struct {
	CpuCount *float64 `pulumi:"cpuCount"`
	MemoryMB *float64 `pulumi:"memoryMB"`
	VmCount  *float64 `pulumi:"vmCount"`
}

type CloudCapacityResponseOutput struct{ *pulumi.OutputState }

func (CloudCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudCapacityResponse)(nil)).Elem()
}

func (o CloudCapacityResponseOutput) ToCloudCapacityResponseOutput() CloudCapacityResponseOutput {
	return o
}

func (o CloudCapacityResponseOutput) ToCloudCapacityResponseOutputWithContext(ctx context.Context) CloudCapacityResponseOutput {
	return o
}

func (o CloudCapacityResponseOutput) CpuCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CloudCapacityResponse) *float64 { return v.CpuCount }).(pulumi.Float64PtrOutput)
}

func (o CloudCapacityResponseOutput) MemoryMB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CloudCapacityResponse) *float64 { return v.MemoryMB }).(pulumi.Float64PtrOutput)
}

func (o CloudCapacityResponseOutput) VmCount() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CloudCapacityResponse) *float64 { return v.VmCount }).(pulumi.Float64PtrOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
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

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Type }).(pulumi.StringPtrOutput)
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
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type HardwareProfile struct {
	CpuCount             *int    `pulumi:"cpuCount"`
	DynamicMemoryEnabled *string `pulumi:"dynamicMemoryEnabled"`
	DynamicMemoryMaxMB   *int    `pulumi:"dynamicMemoryMaxMB"`
	DynamicMemoryMinMB   *int    `pulumi:"dynamicMemoryMinMB"`
	IsHighlyAvailable    *string `pulumi:"isHighlyAvailable"`
	LimitCpuForMigration *string `pulumi:"limitCpuForMigration"`
	MemoryMB             *int    `pulumi:"memoryMB"`
}





type HardwareProfileInput interface {
	pulumi.Input

	ToHardwareProfileOutput() HardwareProfileOutput
	ToHardwareProfileOutputWithContext(context.Context) HardwareProfileOutput
}

type HardwareProfileArgs struct {
	CpuCount             pulumi.IntPtrInput    `pulumi:"cpuCount"`
	DynamicMemoryEnabled pulumi.StringPtrInput `pulumi:"dynamicMemoryEnabled"`
	DynamicMemoryMaxMB   pulumi.IntPtrInput    `pulumi:"dynamicMemoryMaxMB"`
	DynamicMemoryMinMB   pulumi.IntPtrInput    `pulumi:"dynamicMemoryMinMB"`
	IsHighlyAvailable    pulumi.StringPtrInput `pulumi:"isHighlyAvailable"`
	LimitCpuForMigration pulumi.StringPtrInput `pulumi:"limitCpuForMigration"`
	MemoryMB             pulumi.IntPtrInput    `pulumi:"memoryMB"`
}

func (HardwareProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (i HardwareProfileArgs) ToHardwareProfileOutput() HardwareProfileOutput {
	return i.ToHardwareProfileOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput)
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i HardwareProfileArgs) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfileOutput).ToHardwareProfilePtrOutputWithContext(ctx)
}









type HardwareProfilePtrInput interface {
	pulumi.Input

	ToHardwareProfilePtrOutput() HardwareProfilePtrOutput
	ToHardwareProfilePtrOutputWithContext(context.Context) HardwareProfilePtrOutput
}

type hardwareProfilePtrType HardwareProfileArgs

func HardwareProfilePtr(v *HardwareProfileArgs) HardwareProfilePtrInput {
	return (*hardwareProfilePtrType)(v)
}

func (*hardwareProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return i.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (i *hardwareProfilePtrType) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HardwareProfilePtrOutput)
}

type HardwareProfileOutput struct{ *pulumi.OutputState }

func (HardwareProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfile)(nil)).Elem()
}

func (o HardwareProfileOutput) ToHardwareProfileOutput() HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfileOutputWithContext(ctx context.Context) HardwareProfileOutput {
	return o
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o.ToHardwareProfilePtrOutputWithContext(context.Background())
}

func (o HardwareProfileOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HardwareProfile) *HardwareProfile {
		return &v
	}).(HardwareProfilePtrOutput)
}

func (o HardwareProfileOutput) CpuCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *int { return v.CpuCount }).(pulumi.IntPtrOutput)
}

func (o HardwareProfileOutput) DynamicMemoryEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *string { return v.DynamicMemoryEnabled }).(pulumi.StringPtrOutput)
}

func (o HardwareProfileOutput) DynamicMemoryMaxMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *int { return v.DynamicMemoryMaxMB }).(pulumi.IntPtrOutput)
}

func (o HardwareProfileOutput) DynamicMemoryMinMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *int { return v.DynamicMemoryMinMB }).(pulumi.IntPtrOutput)
}

func (o HardwareProfileOutput) IsHighlyAvailable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *string { return v.IsHighlyAvailable }).(pulumi.StringPtrOutput)
}

func (o HardwareProfileOutput) LimitCpuForMigration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *string { return v.LimitCpuForMigration }).(pulumi.StringPtrOutput)
}

func (o HardwareProfileOutput) MemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfile) *int { return v.MemoryMB }).(pulumi.IntPtrOutput)
}

type HardwareProfilePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfile)(nil)).Elem()
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutput() HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) ToHardwareProfilePtrOutputWithContext(ctx context.Context) HardwareProfilePtrOutput {
	return o
}

func (o HardwareProfilePtrOutput) Elem() HardwareProfileOutput {
	return o.ApplyT(func(v *HardwareProfile) HardwareProfile {
		if v != nil {
			return *v
		}
		var ret HardwareProfile
		return ret
	}).(HardwareProfileOutput)
}

func (o HardwareProfilePtrOutput) CpuCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.CpuCount
	}).(pulumi.IntPtrOutput)
}

func (o HardwareProfilePtrOutput) DynamicMemoryEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.DynamicMemoryEnabled
	}).(pulumi.StringPtrOutput)
}

func (o HardwareProfilePtrOutput) DynamicMemoryMaxMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.DynamicMemoryMaxMB
	}).(pulumi.IntPtrOutput)
}

func (o HardwareProfilePtrOutput) DynamicMemoryMinMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.DynamicMemoryMinMB
	}).(pulumi.IntPtrOutput)
}

func (o HardwareProfilePtrOutput) IsHighlyAvailable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.IsHighlyAvailable
	}).(pulumi.StringPtrOutput)
}

func (o HardwareProfilePtrOutput) LimitCpuForMigration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *string {
		if v == nil {
			return nil
		}
		return v.LimitCpuForMigration
	}).(pulumi.StringPtrOutput)
}

func (o HardwareProfilePtrOutput) MemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfile) *int {
		if v == nil {
			return nil
		}
		return v.MemoryMB
	}).(pulumi.IntPtrOutput)
}

type HardwareProfileResponse struct {
	CpuCount             *int    `pulumi:"cpuCount"`
	DynamicMemoryEnabled *string `pulumi:"dynamicMemoryEnabled"`
	DynamicMemoryMaxMB   *int    `pulumi:"dynamicMemoryMaxMB"`
	DynamicMemoryMinMB   *int    `pulumi:"dynamicMemoryMinMB"`
	IsHighlyAvailable    *string `pulumi:"isHighlyAvailable"`
	LimitCpuForMigration *string `pulumi:"limitCpuForMigration"`
	MemoryMB             *int    `pulumi:"memoryMB"`
}

type HardwareProfileResponseOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutput() HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) ToHardwareProfileResponseOutputWithContext(ctx context.Context) HardwareProfileResponseOutput {
	return o
}

func (o HardwareProfileResponseOutput) CpuCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *int { return v.CpuCount }).(pulumi.IntPtrOutput)
}

func (o HardwareProfileResponseOutput) DynamicMemoryEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *string { return v.DynamicMemoryEnabled }).(pulumi.StringPtrOutput)
}

func (o HardwareProfileResponseOutput) DynamicMemoryMaxMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *int { return v.DynamicMemoryMaxMB }).(pulumi.IntPtrOutput)
}

func (o HardwareProfileResponseOutput) DynamicMemoryMinMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *int { return v.DynamicMemoryMinMB }).(pulumi.IntPtrOutput)
}

func (o HardwareProfileResponseOutput) IsHighlyAvailable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *string { return v.IsHighlyAvailable }).(pulumi.StringPtrOutput)
}

func (o HardwareProfileResponseOutput) LimitCpuForMigration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *string { return v.LimitCpuForMigration }).(pulumi.StringPtrOutput)
}

func (o HardwareProfileResponseOutput) MemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v HardwareProfileResponse) *int { return v.MemoryMB }).(pulumi.IntPtrOutput)
}

type HardwareProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HardwareProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HardwareProfileResponse)(nil)).Elem()
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutput() HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) ToHardwareProfileResponsePtrOutputWithContext(ctx context.Context) HardwareProfileResponsePtrOutput {
	return o
}

func (o HardwareProfileResponsePtrOutput) Elem() HardwareProfileResponseOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) HardwareProfileResponse {
		if v != nil {
			return *v
		}
		var ret HardwareProfileResponse
		return ret
	}).(HardwareProfileResponseOutput)
}

func (o HardwareProfileResponsePtrOutput) CpuCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.CpuCount
	}).(pulumi.IntPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) DynamicMemoryEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DynamicMemoryEnabled
	}).(pulumi.StringPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) DynamicMemoryMaxMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.DynamicMemoryMaxMB
	}).(pulumi.IntPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) DynamicMemoryMinMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.DynamicMemoryMinMB
	}).(pulumi.IntPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) IsHighlyAvailable() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.IsHighlyAvailable
	}).(pulumi.StringPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) LimitCpuForMigration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.LimitCpuForMigration
	}).(pulumi.StringPtrOutput)
}

func (o HardwareProfileResponsePtrOutput) MemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *HardwareProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.MemoryMB
	}).(pulumi.IntPtrOutput)
}

type NetworkInterfaces struct {
	Ipv4AddressType  *string `pulumi:"ipv4AddressType"`
	Ipv6AddressType  *string `pulumi:"ipv6AddressType"`
	MacAddress       *string `pulumi:"macAddress"`
	MacAddressType   *string `pulumi:"macAddressType"`
	Name             *string `pulumi:"name"`
	NicId            *string `pulumi:"nicId"`
	VirtualNetworkId *string `pulumi:"virtualNetworkId"`
}





type NetworkInterfacesInput interface {
	pulumi.Input

	ToNetworkInterfacesOutput() NetworkInterfacesOutput
	ToNetworkInterfacesOutputWithContext(context.Context) NetworkInterfacesOutput
}

type NetworkInterfacesArgs struct {
	Ipv4AddressType  pulumi.StringPtrInput `pulumi:"ipv4AddressType"`
	Ipv6AddressType  pulumi.StringPtrInput `pulumi:"ipv6AddressType"`
	MacAddress       pulumi.StringPtrInput `pulumi:"macAddress"`
	MacAddressType   pulumi.StringPtrInput `pulumi:"macAddressType"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	NicId            pulumi.StringPtrInput `pulumi:"nicId"`
	VirtualNetworkId pulumi.StringPtrInput `pulumi:"virtualNetworkId"`
}

func (NetworkInterfacesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaces)(nil)).Elem()
}

func (i NetworkInterfacesArgs) ToNetworkInterfacesOutput() NetworkInterfacesOutput {
	return i.ToNetworkInterfacesOutputWithContext(context.Background())
}

func (i NetworkInterfacesArgs) ToNetworkInterfacesOutputWithContext(ctx context.Context) NetworkInterfacesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfacesOutput)
}





type NetworkInterfacesArrayInput interface {
	pulumi.Input

	ToNetworkInterfacesArrayOutput() NetworkInterfacesArrayOutput
	ToNetworkInterfacesArrayOutputWithContext(context.Context) NetworkInterfacesArrayOutput
}

type NetworkInterfacesArray []NetworkInterfacesInput

func (NetworkInterfacesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaces)(nil)).Elem()
}

func (i NetworkInterfacesArray) ToNetworkInterfacesArrayOutput() NetworkInterfacesArrayOutput {
	return i.ToNetworkInterfacesArrayOutputWithContext(context.Background())
}

func (i NetworkInterfacesArray) ToNetworkInterfacesArrayOutputWithContext(ctx context.Context) NetworkInterfacesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfacesArrayOutput)
}

type NetworkInterfacesOutput struct{ *pulumi.OutputState }

func (NetworkInterfacesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaces)(nil)).Elem()
}

func (o NetworkInterfacesOutput) ToNetworkInterfacesOutput() NetworkInterfacesOutput {
	return o
}

func (o NetworkInterfacesOutput) ToNetworkInterfacesOutputWithContext(ctx context.Context) NetworkInterfacesOutput {
	return o
}

func (o NetworkInterfacesOutput) Ipv4AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaces) *string { return v.Ipv4AddressType }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesOutput) Ipv6AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaces) *string { return v.Ipv6AddressType }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaces) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesOutput) MacAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaces) *string { return v.MacAddressType }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaces) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesOutput) NicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaces) *string { return v.NicId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesOutput) VirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaces) *string { return v.VirtualNetworkId }).(pulumi.StringPtrOutput)
}

type NetworkInterfacesArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfacesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaces)(nil)).Elem()
}

func (o NetworkInterfacesArrayOutput) ToNetworkInterfacesArrayOutput() NetworkInterfacesArrayOutput {
	return o
}

func (o NetworkInterfacesArrayOutput) ToNetworkInterfacesArrayOutputWithContext(ctx context.Context) NetworkInterfacesArrayOutput {
	return o
}

func (o NetworkInterfacesArrayOutput) Index(i pulumi.IntInput) NetworkInterfacesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaces {
		return vs[0].([]NetworkInterfaces)[vs[1].(int)]
	}).(NetworkInterfacesOutput)
}

type NetworkInterfacesResponse struct {
	DisplayName      string   `pulumi:"displayName"`
	Ipv4AddressType  *string  `pulumi:"ipv4AddressType"`
	Ipv4Addresses    []string `pulumi:"ipv4Addresses"`
	Ipv6AddressType  *string  `pulumi:"ipv6AddressType"`
	Ipv6Addresses    []string `pulumi:"ipv6Addresses"`
	MacAddress       *string  `pulumi:"macAddress"`
	MacAddressType   *string  `pulumi:"macAddressType"`
	Name             *string  `pulumi:"name"`
	NetworkName      string   `pulumi:"networkName"`
	NicId            *string  `pulumi:"nicId"`
	VirtualNetworkId *string  `pulumi:"virtualNetworkId"`
}

type NetworkInterfacesResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfacesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfacesResponse)(nil)).Elem()
}

func (o NetworkInterfacesResponseOutput) ToNetworkInterfacesResponseOutput() NetworkInterfacesResponseOutput {
	return o
}

func (o NetworkInterfacesResponseOutput) ToNetworkInterfacesResponseOutputWithContext(ctx context.Context) NetworkInterfacesResponseOutput {
	return o
}

func (o NetworkInterfacesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o NetworkInterfacesResponseOutput) Ipv4AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) *string { return v.Ipv4AddressType }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesResponseOutput) Ipv4Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) []string { return v.Ipv4Addresses }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfacesResponseOutput) Ipv6AddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) *string { return v.Ipv6AddressType }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesResponseOutput) Ipv6Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) []string { return v.Ipv6Addresses }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfacesResponseOutput) MacAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) *string { return v.MacAddress }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesResponseOutput) MacAddressType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) *string { return v.MacAddressType }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesResponseOutput) NetworkName() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) string { return v.NetworkName }).(pulumi.StringOutput)
}

func (o NetworkInterfacesResponseOutput) NicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) *string { return v.NicId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesResponseOutput) VirtualNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResponse) *string { return v.VirtualNetworkId }).(pulumi.StringPtrOutput)
}

type NetworkInterfacesResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfacesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfacesResponse)(nil)).Elem()
}

func (o NetworkInterfacesResponseArrayOutput) ToNetworkInterfacesResponseArrayOutput() NetworkInterfacesResponseArrayOutput {
	return o
}

func (o NetworkInterfacesResponseArrayOutput) ToNetworkInterfacesResponseArrayOutputWithContext(ctx context.Context) NetworkInterfacesResponseArrayOutput {
	return o
}

func (o NetworkInterfacesResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfacesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfacesResponse {
		return vs[0].([]NetworkInterfacesResponse)[vs[1].(int)]
	}).(NetworkInterfacesResponseOutput)
}

type NetworkProfile struct {
	NetworkInterfaces []NetworkInterfaces `pulumi:"networkInterfaces"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	NetworkInterfaces NetworkInterfacesArrayInput `pulumi:"networkInterfaces"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) NetworkInterfaces() NetworkInterfacesArrayOutput {
	return o.ApplyT(func(v NetworkProfile) []NetworkInterfaces { return v.NetworkInterfaces }).(NetworkInterfacesArrayOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) NetworkInterfaces() NetworkInterfacesArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) []NetworkInterfaces {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfacesArrayOutput)
}

type NetworkProfileResponse struct {
	NetworkInterfaces []NetworkInterfacesResponse `pulumi:"networkInterfaces"`
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) NetworkInterfaces() NetworkInterfacesResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []NetworkInterfacesResponse { return v.NetworkInterfaces }).(NetworkInterfacesResponseArrayOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) NetworkInterfaces() NetworkInterfacesResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []NetworkInterfacesResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfacesResponseArrayOutput)
}

type OsProfile struct {
	AdminPassword *string `pulumi:"adminPassword"`
	ComputerName  *string `pulumi:"computerName"`
}





type OsProfileInput interface {
	pulumi.Input

	ToOsProfileOutput() OsProfileOutput
	ToOsProfileOutputWithContext(context.Context) OsProfileOutput
}

type OsProfileArgs struct {
	AdminPassword pulumi.StringPtrInput `pulumi:"adminPassword"`
	ComputerName  pulumi.StringPtrInput `pulumi:"computerName"`
}

func (OsProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfile)(nil)).Elem()
}

func (i OsProfileArgs) ToOsProfileOutput() OsProfileOutput {
	return i.ToOsProfileOutputWithContext(context.Background())
}

func (i OsProfileArgs) ToOsProfileOutputWithContext(ctx context.Context) OsProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileOutput)
}

func (i OsProfileArgs) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return i.ToOsProfilePtrOutputWithContext(context.Background())
}

func (i OsProfileArgs) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfileOutput).ToOsProfilePtrOutputWithContext(ctx)
}









type OsProfilePtrInput interface {
	pulumi.Input

	ToOsProfilePtrOutput() OsProfilePtrOutput
	ToOsProfilePtrOutputWithContext(context.Context) OsProfilePtrOutput
}

type osProfilePtrType OsProfileArgs

func OsProfilePtr(v *OsProfileArgs) OsProfilePtrInput {
	return (*osProfilePtrType)(v)
}

func (*osProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfile)(nil)).Elem()
}

func (i *osProfilePtrType) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return i.ToOsProfilePtrOutputWithContext(context.Background())
}

func (i *osProfilePtrType) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsProfilePtrOutput)
}

type OsProfileOutput struct{ *pulumi.OutputState }

func (OsProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfile)(nil)).Elem()
}

func (o OsProfileOutput) ToOsProfileOutput() OsProfileOutput {
	return o
}

func (o OsProfileOutput) ToOsProfileOutputWithContext(ctx context.Context) OsProfileOutput {
	return o
}

func (o OsProfileOutput) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return o.ToOsProfilePtrOutputWithContext(context.Background())
}

func (o OsProfileOutput) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OsProfile) *OsProfile {
		return &v
	}).(OsProfilePtrOutput)
}

func (o OsProfileOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfile) *string { return v.AdminPassword }).(pulumi.StringPtrOutput)
}

func (o OsProfileOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfile) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

type OsProfilePtrOutput struct{ *pulumi.OutputState }

func (OsProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfile)(nil)).Elem()
}

func (o OsProfilePtrOutput) ToOsProfilePtrOutput() OsProfilePtrOutput {
	return o
}

func (o OsProfilePtrOutput) ToOsProfilePtrOutputWithContext(ctx context.Context) OsProfilePtrOutput {
	return o
}

func (o OsProfilePtrOutput) Elem() OsProfileOutput {
	return o.ApplyT(func(v *OsProfile) OsProfile {
		if v != nil {
			return *v
		}
		var ret OsProfile
		return ret
	}).(OsProfileOutput)
}

func (o OsProfilePtrOutput) AdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfile) *string {
		if v == nil {
			return nil
		}
		return v.AdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o OsProfilePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfile) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

type OsProfileResponse struct {
	ComputerName *string `pulumi:"computerName"`
	OsName       string  `pulumi:"osName"`
	OsType       string  `pulumi:"osType"`
}

type OsProfileResponseOutput struct{ *pulumi.OutputState }

func (OsProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsProfileResponse)(nil)).Elem()
}

func (o OsProfileResponseOutput) ToOsProfileResponseOutput() OsProfileResponseOutput {
	return o
}

func (o OsProfileResponseOutput) ToOsProfileResponseOutputWithContext(ctx context.Context) OsProfileResponseOutput {
	return o
}

func (o OsProfileResponseOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OsProfileResponse) *string { return v.ComputerName }).(pulumi.StringPtrOutput)
}

func (o OsProfileResponseOutput) OsName() pulumi.StringOutput {
	return o.ApplyT(func(v OsProfileResponse) string { return v.OsName }).(pulumi.StringOutput)
}

func (o OsProfileResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v OsProfileResponse) string { return v.OsType }).(pulumi.StringOutput)
}

type OsProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (OsProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OsProfileResponse)(nil)).Elem()
}

func (o OsProfileResponsePtrOutput) ToOsProfileResponsePtrOutput() OsProfileResponsePtrOutput {
	return o
}

func (o OsProfileResponsePtrOutput) ToOsProfileResponsePtrOutputWithContext(ctx context.Context) OsProfileResponsePtrOutput {
	return o
}

func (o OsProfileResponsePtrOutput) Elem() OsProfileResponseOutput {
	return o.ApplyT(func(v *OsProfileResponse) OsProfileResponse {
		if v != nil {
			return *v
		}
		var ret OsProfileResponse
		return ret
	}).(OsProfileResponseOutput)
}

func (o OsProfileResponsePtrOutput) ComputerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ComputerName
	}).(pulumi.StringPtrOutput)
}

func (o OsProfileResponsePtrOutput) OsName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsName
	}).(pulumi.StringPtrOutput)
}

func (o OsProfileResponsePtrOutput) OsType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OsProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OsType
	}).(pulumi.StringPtrOutput)
}

type StorageProfile struct {
	Disks []VirtualDisk `pulumi:"disks"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	Disks VirtualDiskArrayInput `pulumi:"disks"`
}

func (StorageProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (i StorageProfileArgs) ToStorageProfileOutput() StorageProfileOutput {
	return i.ToStorageProfileOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput)
}

func (i StorageProfileArgs) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i StorageProfileArgs) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileOutput).ToStorageProfilePtrOutputWithContext(ctx)
}









type StorageProfilePtrInput interface {
	pulumi.Input

	ToStorageProfilePtrOutput() StorageProfilePtrOutput
	ToStorageProfilePtrOutputWithContext(context.Context) StorageProfilePtrOutput
}

type storageProfilePtrType StorageProfileArgs

func StorageProfilePtr(v *StorageProfileArgs) StorageProfilePtrInput {
	return (*storageProfilePtrType)(v)
}

func (*storageProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return i.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (i *storageProfilePtrType) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfilePtrOutput)
}

type StorageProfileOutput struct{ *pulumi.OutputState }

func (StorageProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfile)(nil)).Elem()
}

func (o StorageProfileOutput) ToStorageProfileOutput() StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfileOutputWithContext(ctx context.Context) StorageProfileOutput {
	return o
}

func (o StorageProfileOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o.ToStorageProfilePtrOutputWithContext(context.Background())
}

func (o StorageProfileOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfile) *StorageProfile {
		return &v
	}).(StorageProfilePtrOutput)
}

func (o StorageProfileOutput) Disks() VirtualDiskArrayOutput {
	return o.ApplyT(func(v StorageProfile) []VirtualDisk { return v.Disks }).(VirtualDiskArrayOutput)
}

type StorageProfilePtrOutput struct{ *pulumi.OutputState }

func (StorageProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfile)(nil)).Elem()
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutput() StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) ToStorageProfilePtrOutputWithContext(ctx context.Context) StorageProfilePtrOutput {
	return o
}

func (o StorageProfilePtrOutput) Elem() StorageProfileOutput {
	return o.ApplyT(func(v *StorageProfile) StorageProfile {
		if v != nil {
			return *v
		}
		var ret StorageProfile
		return ret
	}).(StorageProfileOutput)
}

func (o StorageProfilePtrOutput) Disks() VirtualDiskArrayOutput {
	return o.ApplyT(func(v *StorageProfile) []VirtualDisk {
		if v == nil {
			return nil
		}
		return v.Disks
	}).(VirtualDiskArrayOutput)
}

type StorageProfileResponse struct {
	Disks []VirtualDiskResponse `pulumi:"disks"`
}

type StorageProfileResponseOutput struct{ *pulumi.OutputState }

func (StorageProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return o
}

func (o StorageProfileResponseOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v StorageProfileResponse) []VirtualDiskResponse { return v.Disks }).(VirtualDiskResponseArrayOutput)
}

type StorageProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o
}

func (o StorageProfileResponsePtrOutput) Elem() StorageProfileResponseOutput {
	return o.ApplyT(func(v *StorageProfileResponse) StorageProfileResponse {
		if v != nil {
			return *v
		}
		var ret StorageProfileResponse
		return ret
	}).(StorageProfileResponseOutput)
}

func (o StorageProfileResponsePtrOutput) Disks() VirtualDiskResponseArrayOutput {
	return o.ApplyT(func(v *StorageProfileResponse) []VirtualDiskResponse {
		if v == nil {
			return nil
		}
		return v.Disks
	}).(VirtualDiskResponseArrayOutput)
}

type StorageQoSPolicyDetails struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}





type StorageQoSPolicyDetailsInput interface {
	pulumi.Input

	ToStorageQoSPolicyDetailsOutput() StorageQoSPolicyDetailsOutput
	ToStorageQoSPolicyDetailsOutputWithContext(context.Context) StorageQoSPolicyDetailsOutput
}

type StorageQoSPolicyDetailsArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (StorageQoSPolicyDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQoSPolicyDetails)(nil)).Elem()
}

func (i StorageQoSPolicyDetailsArgs) ToStorageQoSPolicyDetailsOutput() StorageQoSPolicyDetailsOutput {
	return i.ToStorageQoSPolicyDetailsOutputWithContext(context.Background())
}

func (i StorageQoSPolicyDetailsArgs) ToStorageQoSPolicyDetailsOutputWithContext(ctx context.Context) StorageQoSPolicyDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQoSPolicyDetailsOutput)
}

func (i StorageQoSPolicyDetailsArgs) ToStorageQoSPolicyDetailsPtrOutput() StorageQoSPolicyDetailsPtrOutput {
	return i.ToStorageQoSPolicyDetailsPtrOutputWithContext(context.Background())
}

func (i StorageQoSPolicyDetailsArgs) ToStorageQoSPolicyDetailsPtrOutputWithContext(ctx context.Context) StorageQoSPolicyDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQoSPolicyDetailsOutput).ToStorageQoSPolicyDetailsPtrOutputWithContext(ctx)
}









type StorageQoSPolicyDetailsPtrInput interface {
	pulumi.Input

	ToStorageQoSPolicyDetailsPtrOutput() StorageQoSPolicyDetailsPtrOutput
	ToStorageQoSPolicyDetailsPtrOutputWithContext(context.Context) StorageQoSPolicyDetailsPtrOutput
}

type storageQoSPolicyDetailsPtrType StorageQoSPolicyDetailsArgs

func StorageQoSPolicyDetailsPtr(v *StorageQoSPolicyDetailsArgs) StorageQoSPolicyDetailsPtrInput {
	return (*storageQoSPolicyDetailsPtrType)(v)
}

func (*storageQoSPolicyDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageQoSPolicyDetails)(nil)).Elem()
}

func (i *storageQoSPolicyDetailsPtrType) ToStorageQoSPolicyDetailsPtrOutput() StorageQoSPolicyDetailsPtrOutput {
	return i.ToStorageQoSPolicyDetailsPtrOutputWithContext(context.Background())
}

func (i *storageQoSPolicyDetailsPtrType) ToStorageQoSPolicyDetailsPtrOutputWithContext(ctx context.Context) StorageQoSPolicyDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQoSPolicyDetailsPtrOutput)
}

type StorageQoSPolicyDetailsOutput struct{ *pulumi.OutputState }

func (StorageQoSPolicyDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQoSPolicyDetails)(nil)).Elem()
}

func (o StorageQoSPolicyDetailsOutput) ToStorageQoSPolicyDetailsOutput() StorageQoSPolicyDetailsOutput {
	return o
}

func (o StorageQoSPolicyDetailsOutput) ToStorageQoSPolicyDetailsOutputWithContext(ctx context.Context) StorageQoSPolicyDetailsOutput {
	return o
}

func (o StorageQoSPolicyDetailsOutput) ToStorageQoSPolicyDetailsPtrOutput() StorageQoSPolicyDetailsPtrOutput {
	return o.ToStorageQoSPolicyDetailsPtrOutputWithContext(context.Background())
}

func (o StorageQoSPolicyDetailsOutput) ToStorageQoSPolicyDetailsPtrOutputWithContext(ctx context.Context) StorageQoSPolicyDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageQoSPolicyDetails) *StorageQoSPolicyDetails {
		return &v
	}).(StorageQoSPolicyDetailsPtrOutput)
}

func (o StorageQoSPolicyDetailsOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQoSPolicyDetails) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StorageQoSPolicyDetailsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQoSPolicyDetails) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type StorageQoSPolicyDetailsPtrOutput struct{ *pulumi.OutputState }

func (StorageQoSPolicyDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageQoSPolicyDetails)(nil)).Elem()
}

func (o StorageQoSPolicyDetailsPtrOutput) ToStorageQoSPolicyDetailsPtrOutput() StorageQoSPolicyDetailsPtrOutput {
	return o
}

func (o StorageQoSPolicyDetailsPtrOutput) ToStorageQoSPolicyDetailsPtrOutputWithContext(ctx context.Context) StorageQoSPolicyDetailsPtrOutput {
	return o
}

func (o StorageQoSPolicyDetailsPtrOutput) Elem() StorageQoSPolicyDetailsOutput {
	return o.ApplyT(func(v *StorageQoSPolicyDetails) StorageQoSPolicyDetails {
		if v != nil {
			return *v
		}
		var ret StorageQoSPolicyDetails
		return ret
	}).(StorageQoSPolicyDetailsOutput)
}

func (o StorageQoSPolicyDetailsPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQoSPolicyDetails) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o StorageQoSPolicyDetailsPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQoSPolicyDetails) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type StorageQoSPolicyDetailsResponse struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}

type StorageQoSPolicyDetailsResponseOutput struct{ *pulumi.OutputState }

func (StorageQoSPolicyDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQoSPolicyDetailsResponse)(nil)).Elem()
}

func (o StorageQoSPolicyDetailsResponseOutput) ToStorageQoSPolicyDetailsResponseOutput() StorageQoSPolicyDetailsResponseOutput {
	return o
}

func (o StorageQoSPolicyDetailsResponseOutput) ToStorageQoSPolicyDetailsResponseOutputWithContext(ctx context.Context) StorageQoSPolicyDetailsResponseOutput {
	return o
}

func (o StorageQoSPolicyDetailsResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQoSPolicyDetailsResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StorageQoSPolicyDetailsResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQoSPolicyDetailsResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type StorageQoSPolicyDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageQoSPolicyDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageQoSPolicyDetailsResponse)(nil)).Elem()
}

func (o StorageQoSPolicyDetailsResponsePtrOutput) ToStorageQoSPolicyDetailsResponsePtrOutput() StorageQoSPolicyDetailsResponsePtrOutput {
	return o
}

func (o StorageQoSPolicyDetailsResponsePtrOutput) ToStorageQoSPolicyDetailsResponsePtrOutputWithContext(ctx context.Context) StorageQoSPolicyDetailsResponsePtrOutput {
	return o
}

func (o StorageQoSPolicyDetailsResponsePtrOutput) Elem() StorageQoSPolicyDetailsResponseOutput {
	return o.ApplyT(func(v *StorageQoSPolicyDetailsResponse) StorageQoSPolicyDetailsResponse {
		if v != nil {
			return *v
		}
		var ret StorageQoSPolicyDetailsResponse
		return ret
	}).(StorageQoSPolicyDetailsResponseOutput)
}

func (o StorageQoSPolicyDetailsResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQoSPolicyDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o StorageQoSPolicyDetailsResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageQoSPolicyDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type StorageQoSPolicyResponse struct {
	BandwidthLimit *float64 `pulumi:"bandwidthLimit"`
	Id             *string  `pulumi:"id"`
	IopsMaximum    *float64 `pulumi:"iopsMaximum"`
	IopsMinimum    *float64 `pulumi:"iopsMinimum"`
	Name           *string  `pulumi:"name"`
	PolicyId       *string  `pulumi:"policyId"`
}

type StorageQoSPolicyResponseOutput struct{ *pulumi.OutputState }

func (StorageQoSPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQoSPolicyResponse)(nil)).Elem()
}

func (o StorageQoSPolicyResponseOutput) ToStorageQoSPolicyResponseOutput() StorageQoSPolicyResponseOutput {
	return o
}

func (o StorageQoSPolicyResponseOutput) ToStorageQoSPolicyResponseOutputWithContext(ctx context.Context) StorageQoSPolicyResponseOutput {
	return o
}

func (o StorageQoSPolicyResponseOutput) BandwidthLimit() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StorageQoSPolicyResponse) *float64 { return v.BandwidthLimit }).(pulumi.Float64PtrOutput)
}

func (o StorageQoSPolicyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQoSPolicyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o StorageQoSPolicyResponseOutput) IopsMaximum() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StorageQoSPolicyResponse) *float64 { return v.IopsMaximum }).(pulumi.Float64PtrOutput)
}

func (o StorageQoSPolicyResponseOutput) IopsMinimum() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StorageQoSPolicyResponse) *float64 { return v.IopsMinimum }).(pulumi.Float64PtrOutput)
}

func (o StorageQoSPolicyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQoSPolicyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StorageQoSPolicyResponseOutput) PolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQoSPolicyResponse) *string { return v.PolicyId }).(pulumi.StringPtrOutput)
}

type StorageQoSPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (StorageQoSPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StorageQoSPolicyResponse)(nil)).Elem()
}

func (o StorageQoSPolicyResponseArrayOutput) ToStorageQoSPolicyResponseArrayOutput() StorageQoSPolicyResponseArrayOutput {
	return o
}

func (o StorageQoSPolicyResponseArrayOutput) ToStorageQoSPolicyResponseArrayOutputWithContext(ctx context.Context) StorageQoSPolicyResponseArrayOutput {
	return o
}

func (o StorageQoSPolicyResponseArrayOutput) Index(i pulumi.IntInput) StorageQoSPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StorageQoSPolicyResponse {
		return vs[0].([]StorageQoSPolicyResponse)[vs[1].(int)]
	}).(StorageQoSPolicyResponseOutput)
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

type VMMServerPropertiesCredentials struct {
	Password *string `pulumi:"password"`
	Username *string `pulumi:"username"`
}





type VMMServerPropertiesCredentialsInput interface {
	pulumi.Input

	ToVMMServerPropertiesCredentialsOutput() VMMServerPropertiesCredentialsOutput
	ToVMMServerPropertiesCredentialsOutputWithContext(context.Context) VMMServerPropertiesCredentialsOutput
}

type VMMServerPropertiesCredentialsArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Username pulumi.StringPtrInput `pulumi:"username"`
}

func (VMMServerPropertiesCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMMServerPropertiesCredentials)(nil)).Elem()
}

func (i VMMServerPropertiesCredentialsArgs) ToVMMServerPropertiesCredentialsOutput() VMMServerPropertiesCredentialsOutput {
	return i.ToVMMServerPropertiesCredentialsOutputWithContext(context.Background())
}

func (i VMMServerPropertiesCredentialsArgs) ToVMMServerPropertiesCredentialsOutputWithContext(ctx context.Context) VMMServerPropertiesCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMMServerPropertiesCredentialsOutput)
}

func (i VMMServerPropertiesCredentialsArgs) ToVMMServerPropertiesCredentialsPtrOutput() VMMServerPropertiesCredentialsPtrOutput {
	return i.ToVMMServerPropertiesCredentialsPtrOutputWithContext(context.Background())
}

func (i VMMServerPropertiesCredentialsArgs) ToVMMServerPropertiesCredentialsPtrOutputWithContext(ctx context.Context) VMMServerPropertiesCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMMServerPropertiesCredentialsOutput).ToVMMServerPropertiesCredentialsPtrOutputWithContext(ctx)
}









type VMMServerPropertiesCredentialsPtrInput interface {
	pulumi.Input

	ToVMMServerPropertiesCredentialsPtrOutput() VMMServerPropertiesCredentialsPtrOutput
	ToVMMServerPropertiesCredentialsPtrOutputWithContext(context.Context) VMMServerPropertiesCredentialsPtrOutput
}

type vmmserverPropertiesCredentialsPtrType VMMServerPropertiesCredentialsArgs

func VMMServerPropertiesCredentialsPtr(v *VMMServerPropertiesCredentialsArgs) VMMServerPropertiesCredentialsPtrInput {
	return (*vmmserverPropertiesCredentialsPtrType)(v)
}

func (*vmmserverPropertiesCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VMMServerPropertiesCredentials)(nil)).Elem()
}

func (i *vmmserverPropertiesCredentialsPtrType) ToVMMServerPropertiesCredentialsPtrOutput() VMMServerPropertiesCredentialsPtrOutput {
	return i.ToVMMServerPropertiesCredentialsPtrOutputWithContext(context.Background())
}

func (i *vmmserverPropertiesCredentialsPtrType) ToVMMServerPropertiesCredentialsPtrOutputWithContext(ctx context.Context) VMMServerPropertiesCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMMServerPropertiesCredentialsPtrOutput)
}

type VMMServerPropertiesCredentialsOutput struct{ *pulumi.OutputState }

func (VMMServerPropertiesCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMMServerPropertiesCredentials)(nil)).Elem()
}

func (o VMMServerPropertiesCredentialsOutput) ToVMMServerPropertiesCredentialsOutput() VMMServerPropertiesCredentialsOutput {
	return o
}

func (o VMMServerPropertiesCredentialsOutput) ToVMMServerPropertiesCredentialsOutputWithContext(ctx context.Context) VMMServerPropertiesCredentialsOutput {
	return o
}

func (o VMMServerPropertiesCredentialsOutput) ToVMMServerPropertiesCredentialsPtrOutput() VMMServerPropertiesCredentialsPtrOutput {
	return o.ToVMMServerPropertiesCredentialsPtrOutputWithContext(context.Background())
}

func (o VMMServerPropertiesCredentialsOutput) ToVMMServerPropertiesCredentialsPtrOutputWithContext(ctx context.Context) VMMServerPropertiesCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VMMServerPropertiesCredentials) *VMMServerPropertiesCredentials {
		return &v
	}).(VMMServerPropertiesCredentialsPtrOutput)
}

func (o VMMServerPropertiesCredentialsOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMMServerPropertiesCredentials) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o VMMServerPropertiesCredentialsOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMMServerPropertiesCredentials) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type VMMServerPropertiesCredentialsPtrOutput struct{ *pulumi.OutputState }

func (VMMServerPropertiesCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMMServerPropertiesCredentials)(nil)).Elem()
}

func (o VMMServerPropertiesCredentialsPtrOutput) ToVMMServerPropertiesCredentialsPtrOutput() VMMServerPropertiesCredentialsPtrOutput {
	return o
}

func (o VMMServerPropertiesCredentialsPtrOutput) ToVMMServerPropertiesCredentialsPtrOutputWithContext(ctx context.Context) VMMServerPropertiesCredentialsPtrOutput {
	return o
}

func (o VMMServerPropertiesCredentialsPtrOutput) Elem() VMMServerPropertiesCredentialsOutput {
	return o.ApplyT(func(v *VMMServerPropertiesCredentials) VMMServerPropertiesCredentials {
		if v != nil {
			return *v
		}
		var ret VMMServerPropertiesCredentials
		return ret
	}).(VMMServerPropertiesCredentialsOutput)
}

func (o VMMServerPropertiesCredentialsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMMServerPropertiesCredentials) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o VMMServerPropertiesCredentialsPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMMServerPropertiesCredentials) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type VMMServerPropertiesResponseCredentials struct {
	Username *string `pulumi:"username"`
}

type VMMServerPropertiesResponseCredentialsOutput struct{ *pulumi.OutputState }

func (VMMServerPropertiesResponseCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMMServerPropertiesResponseCredentials)(nil)).Elem()
}

func (o VMMServerPropertiesResponseCredentialsOutput) ToVMMServerPropertiesResponseCredentialsOutput() VMMServerPropertiesResponseCredentialsOutput {
	return o
}

func (o VMMServerPropertiesResponseCredentialsOutput) ToVMMServerPropertiesResponseCredentialsOutputWithContext(ctx context.Context) VMMServerPropertiesResponseCredentialsOutput {
	return o
}

func (o VMMServerPropertiesResponseCredentialsOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VMMServerPropertiesResponseCredentials) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type VMMServerPropertiesResponseCredentialsPtrOutput struct{ *pulumi.OutputState }

func (VMMServerPropertiesResponseCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMMServerPropertiesResponseCredentials)(nil)).Elem()
}

func (o VMMServerPropertiesResponseCredentialsPtrOutput) ToVMMServerPropertiesResponseCredentialsPtrOutput() VMMServerPropertiesResponseCredentialsPtrOutput {
	return o
}

func (o VMMServerPropertiesResponseCredentialsPtrOutput) ToVMMServerPropertiesResponseCredentialsPtrOutputWithContext(ctx context.Context) VMMServerPropertiesResponseCredentialsPtrOutput {
	return o
}

func (o VMMServerPropertiesResponseCredentialsPtrOutput) Elem() VMMServerPropertiesResponseCredentialsOutput {
	return o.ApplyT(func(v *VMMServerPropertiesResponseCredentials) VMMServerPropertiesResponseCredentials {
		if v != nil {
			return *v
		}
		var ret VMMServerPropertiesResponseCredentials
		return ret
	}).(VMMServerPropertiesResponseCredentialsOutput)
}

func (o VMMServerPropertiesResponseCredentialsPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMMServerPropertiesResponseCredentials) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type VirtualDisk struct {
	Bus              *int                     `pulumi:"bus"`
	BusType          *string                  `pulumi:"busType"`
	CreateDiffDisk   *string                  `pulumi:"createDiffDisk"`
	DiskId           *string                  `pulumi:"diskId"`
	DiskSizeGB       *int                     `pulumi:"diskSizeGB"`
	Lun              *int                     `pulumi:"lun"`
	Name             *string                  `pulumi:"name"`
	StorageQoSPolicy *StorageQoSPolicyDetails `pulumi:"storageQoSPolicy"`
	TemplateDiskId   *string                  `pulumi:"templateDiskId"`
	VhdType          *string                  `pulumi:"vhdType"`
}





type VirtualDiskInput interface {
	pulumi.Input

	ToVirtualDiskOutput() VirtualDiskOutput
	ToVirtualDiskOutputWithContext(context.Context) VirtualDiskOutput
}

type VirtualDiskArgs struct {
	Bus              pulumi.IntPtrInput              `pulumi:"bus"`
	BusType          pulumi.StringPtrInput           `pulumi:"busType"`
	CreateDiffDisk   pulumi.StringPtrInput           `pulumi:"createDiffDisk"`
	DiskId           pulumi.StringPtrInput           `pulumi:"diskId"`
	DiskSizeGB       pulumi.IntPtrInput              `pulumi:"diskSizeGB"`
	Lun              pulumi.IntPtrInput              `pulumi:"lun"`
	Name             pulumi.StringPtrInput           `pulumi:"name"`
	StorageQoSPolicy StorageQoSPolicyDetailsPtrInput `pulumi:"storageQoSPolicy"`
	TemplateDiskId   pulumi.StringPtrInput           `pulumi:"templateDiskId"`
	VhdType          pulumi.StringPtrInput           `pulumi:"vhdType"`
}

func (VirtualDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDisk)(nil)).Elem()
}

func (i VirtualDiskArgs) ToVirtualDiskOutput() VirtualDiskOutput {
	return i.ToVirtualDiskOutputWithContext(context.Background())
}

func (i VirtualDiskArgs) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskOutput)
}





type VirtualDiskArrayInput interface {
	pulumi.Input

	ToVirtualDiskArrayOutput() VirtualDiskArrayOutput
	ToVirtualDiskArrayOutputWithContext(context.Context) VirtualDiskArrayOutput
}

type VirtualDiskArray []VirtualDiskInput

func (VirtualDiskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDisk)(nil)).Elem()
}

func (i VirtualDiskArray) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return i.ToVirtualDiskArrayOutputWithContext(context.Background())
}

func (i VirtualDiskArray) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDiskArrayOutput)
}

type VirtualDiskOutput struct{ *pulumi.OutputState }

func (VirtualDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskOutput) ToVirtualDiskOutput() VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) ToVirtualDiskOutputWithContext(ctx context.Context) VirtualDiskOutput {
	return o
}

func (o VirtualDiskOutput) Bus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *int { return v.Bus }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskOutput) BusType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.BusType }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskOutput) CreateDiffDisk() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.CreateDiffDisk }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskOutput) DiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.DiskId }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *int { return v.Lun }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskOutput) StorageQoSPolicy() StorageQoSPolicyDetailsPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *StorageQoSPolicyDetails { return v.StorageQoSPolicy }).(StorageQoSPolicyDetailsPtrOutput)
}

func (o VirtualDiskOutput) TemplateDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.TemplateDiskId }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskOutput) VhdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDisk) *string { return v.VhdType }).(pulumi.StringPtrOutput)
}

type VirtualDiskArrayOutput struct{ *pulumi.OutputState }

func (VirtualDiskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDisk)(nil)).Elem()
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutput() VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) ToVirtualDiskArrayOutputWithContext(ctx context.Context) VirtualDiskArrayOutput {
	return o
}

func (o VirtualDiskArrayOutput) Index(i pulumi.IntInput) VirtualDiskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDisk {
		return vs[0].([]VirtualDisk)[vs[1].(int)]
	}).(VirtualDiskOutput)
}

type VirtualDiskResponse struct {
	Bus              *int                             `pulumi:"bus"`
	BusType          *string                          `pulumi:"busType"`
	CreateDiffDisk   *string                          `pulumi:"createDiffDisk"`
	DiskId           *string                          `pulumi:"diskId"`
	DiskSizeGB       *int                             `pulumi:"diskSizeGB"`
	DisplayName      string                           `pulumi:"displayName"`
	Lun              *int                             `pulumi:"lun"`
	MaxDiskSizeGB    int                              `pulumi:"maxDiskSizeGB"`
	Name             *string                          `pulumi:"name"`
	StorageQoSPolicy *StorageQoSPolicyDetailsResponse `pulumi:"storageQoSPolicy"`
	TemplateDiskId   *string                          `pulumi:"templateDiskId"`
	VhdFormatType    string                           `pulumi:"vhdFormatType"`
	VhdType          *string                          `pulumi:"vhdType"`
	VolumeType       string                           `pulumi:"volumeType"`
}

type VirtualDiskResponseOutput struct{ *pulumi.OutputState }

func (VirtualDiskResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDiskResponse)(nil)).Elem()
}

func (o VirtualDiskResponseOutput) ToVirtualDiskResponseOutput() VirtualDiskResponseOutput {
	return o
}

func (o VirtualDiskResponseOutput) ToVirtualDiskResponseOutputWithContext(ctx context.Context) VirtualDiskResponseOutput {
	return o
}

func (o VirtualDiskResponseOutput) Bus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *int { return v.Bus }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskResponseOutput) BusType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.BusType }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) CreateDiffDisk() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.CreateDiffDisk }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) DiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.DiskId }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) DiskSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *int { return v.DiskSizeGB }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o VirtualDiskResponseOutput) Lun() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *int { return v.Lun }).(pulumi.IntPtrOutput)
}

func (o VirtualDiskResponseOutput) MaxDiskSizeGB() pulumi.IntOutput {
	return o.ApplyT(func(v VirtualDiskResponse) int { return v.MaxDiskSizeGB }).(pulumi.IntOutput)
}

func (o VirtualDiskResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) StorageQoSPolicy() StorageQoSPolicyDetailsResponsePtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *StorageQoSPolicyDetailsResponse { return v.StorageQoSPolicy }).(StorageQoSPolicyDetailsResponsePtrOutput)
}

func (o VirtualDiskResponseOutput) TemplateDiskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.TemplateDiskId }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) VhdFormatType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskResponse) string { return v.VhdFormatType }).(pulumi.StringOutput)
}

func (o VirtualDiskResponseOutput) VhdType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDiskResponse) *string { return v.VhdType }).(pulumi.StringPtrOutput)
}

func (o VirtualDiskResponseOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualDiskResponse) string { return v.VolumeType }).(pulumi.StringOutput)
}

type VirtualDiskResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualDiskResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDiskResponse)(nil)).Elem()
}

func (o VirtualDiskResponseArrayOutput) ToVirtualDiskResponseArrayOutput() VirtualDiskResponseArrayOutput {
	return o
}

func (o VirtualDiskResponseArrayOutput) ToVirtualDiskResponseArrayOutputWithContext(ctx context.Context) VirtualDiskResponseArrayOutput {
	return o
}

func (o VirtualDiskResponseArrayOutput) Index(i pulumi.IntInput) VirtualDiskResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDiskResponse {
		return vs[0].([]VirtualDiskResponse)[vs[1].(int)]
	}).(VirtualDiskResponseOutput)
}

type VirtualMachinePropertiesAvailabilitySets struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}





type VirtualMachinePropertiesAvailabilitySetsInput interface {
	pulumi.Input

	ToVirtualMachinePropertiesAvailabilitySetsOutput() VirtualMachinePropertiesAvailabilitySetsOutput
	ToVirtualMachinePropertiesAvailabilitySetsOutputWithContext(context.Context) VirtualMachinePropertiesAvailabilitySetsOutput
}

type VirtualMachinePropertiesAvailabilitySetsArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (VirtualMachinePropertiesAvailabilitySetsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePropertiesAvailabilitySets)(nil)).Elem()
}

func (i VirtualMachinePropertiesAvailabilitySetsArgs) ToVirtualMachinePropertiesAvailabilitySetsOutput() VirtualMachinePropertiesAvailabilitySetsOutput {
	return i.ToVirtualMachinePropertiesAvailabilitySetsOutputWithContext(context.Background())
}

func (i VirtualMachinePropertiesAvailabilitySetsArgs) ToVirtualMachinePropertiesAvailabilitySetsOutputWithContext(ctx context.Context) VirtualMachinePropertiesAvailabilitySetsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePropertiesAvailabilitySetsOutput)
}





type VirtualMachinePropertiesAvailabilitySetsArrayInput interface {
	pulumi.Input

	ToVirtualMachinePropertiesAvailabilitySetsArrayOutput() VirtualMachinePropertiesAvailabilitySetsArrayOutput
	ToVirtualMachinePropertiesAvailabilitySetsArrayOutputWithContext(context.Context) VirtualMachinePropertiesAvailabilitySetsArrayOutput
}

type VirtualMachinePropertiesAvailabilitySetsArray []VirtualMachinePropertiesAvailabilitySetsInput

func (VirtualMachinePropertiesAvailabilitySetsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachinePropertiesAvailabilitySets)(nil)).Elem()
}

func (i VirtualMachinePropertiesAvailabilitySetsArray) ToVirtualMachinePropertiesAvailabilitySetsArrayOutput() VirtualMachinePropertiesAvailabilitySetsArrayOutput {
	return i.ToVirtualMachinePropertiesAvailabilitySetsArrayOutputWithContext(context.Background())
}

func (i VirtualMachinePropertiesAvailabilitySetsArray) ToVirtualMachinePropertiesAvailabilitySetsArrayOutputWithContext(ctx context.Context) VirtualMachinePropertiesAvailabilitySetsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachinePropertiesAvailabilitySetsArrayOutput)
}

type VirtualMachinePropertiesAvailabilitySetsOutput struct{ *pulumi.OutputState }

func (VirtualMachinePropertiesAvailabilitySetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePropertiesAvailabilitySets)(nil)).Elem()
}

func (o VirtualMachinePropertiesAvailabilitySetsOutput) ToVirtualMachinePropertiesAvailabilitySetsOutput() VirtualMachinePropertiesAvailabilitySetsOutput {
	return o
}

func (o VirtualMachinePropertiesAvailabilitySetsOutput) ToVirtualMachinePropertiesAvailabilitySetsOutputWithContext(ctx context.Context) VirtualMachinePropertiesAvailabilitySetsOutput {
	return o
}

func (o VirtualMachinePropertiesAvailabilitySetsOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachinePropertiesAvailabilitySets) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePropertiesAvailabilitySetsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachinePropertiesAvailabilitySets) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VirtualMachinePropertiesAvailabilitySetsArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachinePropertiesAvailabilitySetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachinePropertiesAvailabilitySets)(nil)).Elem()
}

func (o VirtualMachinePropertiesAvailabilitySetsArrayOutput) ToVirtualMachinePropertiesAvailabilitySetsArrayOutput() VirtualMachinePropertiesAvailabilitySetsArrayOutput {
	return o
}

func (o VirtualMachinePropertiesAvailabilitySetsArrayOutput) ToVirtualMachinePropertiesAvailabilitySetsArrayOutputWithContext(ctx context.Context) VirtualMachinePropertiesAvailabilitySetsArrayOutput {
	return o
}

func (o VirtualMachinePropertiesAvailabilitySetsArrayOutput) Index(i pulumi.IntInput) VirtualMachinePropertiesAvailabilitySetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachinePropertiesAvailabilitySets {
		return vs[0].([]VirtualMachinePropertiesAvailabilitySets)[vs[1].(int)]
	}).(VirtualMachinePropertiesAvailabilitySetsOutput)
}

type VirtualMachinePropertiesResponseAvailabilitySets struct {
	Id   *string `pulumi:"id"`
	Name *string `pulumi:"name"`
}

type VirtualMachinePropertiesResponseAvailabilitySetsOutput struct{ *pulumi.OutputState }

func (VirtualMachinePropertiesResponseAvailabilitySetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachinePropertiesResponseAvailabilitySets)(nil)).Elem()
}

func (o VirtualMachinePropertiesResponseAvailabilitySetsOutput) ToVirtualMachinePropertiesResponseAvailabilitySetsOutput() VirtualMachinePropertiesResponseAvailabilitySetsOutput {
	return o
}

func (o VirtualMachinePropertiesResponseAvailabilitySetsOutput) ToVirtualMachinePropertiesResponseAvailabilitySetsOutputWithContext(ctx context.Context) VirtualMachinePropertiesResponseAvailabilitySetsOutput {
	return o
}

func (o VirtualMachinePropertiesResponseAvailabilitySetsOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachinePropertiesResponseAvailabilitySets) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualMachinePropertiesResponseAvailabilitySetsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualMachinePropertiesResponseAvailabilitySets) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput struct{ *pulumi.OutputState }

func (VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualMachinePropertiesResponseAvailabilitySets)(nil)).Elem()
}

func (o VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput) ToVirtualMachinePropertiesResponseAvailabilitySetsArrayOutput() VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput {
	return o
}

func (o VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput) ToVirtualMachinePropertiesResponseAvailabilitySetsArrayOutputWithContext(ctx context.Context) VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput {
	return o
}

func (o VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput) Index(i pulumi.IntInput) VirtualMachinePropertiesResponseAvailabilitySetsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualMachinePropertiesResponseAvailabilitySets {
		return vs[0].([]VirtualMachinePropertiesResponseAvailabilitySets)[vs[1].(int)]
	}).(VirtualMachinePropertiesResponseAvailabilitySetsOutput)
}

func init() {
	pulumi.RegisterOutputType(CheckpointOutput{})
	pulumi.RegisterOutputType(CheckpointArrayOutput{})
	pulumi.RegisterOutputType(CheckpointResponseOutput{})
	pulumi.RegisterOutputType(CheckpointResponseArrayOutput{})
	pulumi.RegisterOutputType(CloudCapacityResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileOutput{})
	pulumi.RegisterOutputType(HardwareProfilePtrOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponseOutput{})
	pulumi.RegisterOutputType(HardwareProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfacesOutput{})
	pulumi.RegisterOutputType(NetworkInterfacesArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfacesResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfacesResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(OsProfileOutput{})
	pulumi.RegisterOutputType(OsProfilePtrOutput{})
	pulumi.RegisterOutputType(OsProfileResponseOutput{})
	pulumi.RegisterOutputType(OsProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageQoSPolicyDetailsOutput{})
	pulumi.RegisterOutputType(StorageQoSPolicyDetailsPtrOutput{})
	pulumi.RegisterOutputType(StorageQoSPolicyDetailsResponseOutput{})
	pulumi.RegisterOutputType(StorageQoSPolicyDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageQoSPolicyResponseOutput{})
	pulumi.RegisterOutputType(StorageQoSPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VMMServerPropertiesCredentialsOutput{})
	pulumi.RegisterOutputType(VMMServerPropertiesCredentialsPtrOutput{})
	pulumi.RegisterOutputType(VMMServerPropertiesResponseCredentialsOutput{})
	pulumi.RegisterOutputType(VMMServerPropertiesResponseCredentialsPtrOutput{})
	pulumi.RegisterOutputType(VirtualDiskOutput{})
	pulumi.RegisterOutputType(VirtualDiskArrayOutput{})
	pulumi.RegisterOutputType(VirtualDiskResponseOutput{})
	pulumi.RegisterOutputType(VirtualDiskResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachinePropertiesAvailabilitySetsOutput{})
	pulumi.RegisterOutputType(VirtualMachinePropertiesAvailabilitySetsArrayOutput{})
	pulumi.RegisterOutputType(VirtualMachinePropertiesResponseAvailabilitySetsOutput{})
	pulumi.RegisterOutputType(VirtualMachinePropertiesResponseAvailabilitySetsArrayOutput{})
}
