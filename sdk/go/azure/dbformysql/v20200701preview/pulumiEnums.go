


package v20200701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreateMode string

const (
	CreateModeDefault            = CreateMode("Default")
	CreateModePointInTimeRestore = CreateMode("PointInTimeRestore")
	CreateModeReplica            = CreateMode("Replica")
)

func (CreateMode) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateMode)(nil)).Elem()
}

func (e CreateMode) ToCreateModeOutput() CreateModeOutput {
	return pulumi.ToOutput(e).(CreateModeOutput)
}

func (e CreateMode) ToCreateModeOutputWithContext(ctx context.Context) CreateModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CreateModeOutput)
}

func (e CreateMode) ToCreateModePtrOutput() CreateModePtrOutput {
	return e.ToCreateModePtrOutputWithContext(context.Background())
}

func (e CreateMode) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return CreateMode(e).ToCreateModeOutputWithContext(ctx).ToCreateModePtrOutputWithContext(ctx)
}

func (e CreateMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreateMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CreateMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CreateModeOutput struct{ *pulumi.OutputState }

func (CreateModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateMode)(nil)).Elem()
}

func (o CreateModeOutput) ToCreateModeOutput() CreateModeOutput {
	return o
}

func (o CreateModeOutput) ToCreateModeOutputWithContext(ctx context.Context) CreateModeOutput {
	return o
}

func (o CreateModeOutput) ToCreateModePtrOutput() CreateModePtrOutput {
	return o.ToCreateModePtrOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateMode) *CreateMode {
		return &v
	}).(CreateModePtrOutput)
}

func (o CreateModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreateMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CreateModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreateModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreateMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CreateModePtrOutput struct{ *pulumi.OutputState }

func (CreateModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateMode)(nil)).Elem()
}

func (o CreateModePtrOutput) ToCreateModePtrOutput() CreateModePtrOutput {
	return o
}

func (o CreateModePtrOutput) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return o
}

func (o CreateModePtrOutput) Elem() CreateModeOutput {
	return o.ApplyT(func(v *CreateMode) CreateMode {
		if v != nil {
			return *v
		}
		var ret CreateMode
		return ret
	}).(CreateModeOutput)
}

func (o CreateModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreateModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CreateMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CreateModeInput interface {
	pulumi.Input

	ToCreateModeOutput() CreateModeOutput
	ToCreateModeOutputWithContext(context.Context) CreateModeOutput
}

var createModePtrType = reflect.TypeOf((**CreateMode)(nil)).Elem()

type CreateModePtrInput interface {
	pulumi.Input

	ToCreateModePtrOutput() CreateModePtrOutput
	ToCreateModePtrOutputWithContext(context.Context) CreateModePtrOutput
}

type createModePtr string

func CreateModePtr(v string) CreateModePtrInput {
	return (*createModePtr)(&v)
}

func (*createModePtr) ElementType() reflect.Type {
	return createModePtrType
}

func (in *createModePtr) ToCreateModePtrOutput() CreateModePtrOutput {
	return pulumi.ToOutput(in).(CreateModePtrOutput)
}

func (in *createModePtr) ToCreateModePtrOutputWithContext(ctx context.Context) CreateModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CreateModePtrOutput)
}

type HaEnabledEnum string

const (
	HaEnabledEnumEnabled  = HaEnabledEnum("Enabled")
	HaEnabledEnumDisabled = HaEnabledEnum("Disabled")
)

func (HaEnabledEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*HaEnabledEnum)(nil)).Elem()
}

func (e HaEnabledEnum) ToHaEnabledEnumOutput() HaEnabledEnumOutput {
	return pulumi.ToOutput(e).(HaEnabledEnumOutput)
}

func (e HaEnabledEnum) ToHaEnabledEnumOutputWithContext(ctx context.Context) HaEnabledEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HaEnabledEnumOutput)
}

func (e HaEnabledEnum) ToHaEnabledEnumPtrOutput() HaEnabledEnumPtrOutput {
	return e.ToHaEnabledEnumPtrOutputWithContext(context.Background())
}

func (e HaEnabledEnum) ToHaEnabledEnumPtrOutputWithContext(ctx context.Context) HaEnabledEnumPtrOutput {
	return HaEnabledEnum(e).ToHaEnabledEnumOutputWithContext(ctx).ToHaEnabledEnumPtrOutputWithContext(ctx)
}

func (e HaEnabledEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HaEnabledEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HaEnabledEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HaEnabledEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HaEnabledEnumOutput struct{ *pulumi.OutputState }

func (HaEnabledEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HaEnabledEnum)(nil)).Elem()
}

func (o HaEnabledEnumOutput) ToHaEnabledEnumOutput() HaEnabledEnumOutput {
	return o
}

func (o HaEnabledEnumOutput) ToHaEnabledEnumOutputWithContext(ctx context.Context) HaEnabledEnumOutput {
	return o
}

func (o HaEnabledEnumOutput) ToHaEnabledEnumPtrOutput() HaEnabledEnumPtrOutput {
	return o.ToHaEnabledEnumPtrOutputWithContext(context.Background())
}

func (o HaEnabledEnumOutput) ToHaEnabledEnumPtrOutputWithContext(ctx context.Context) HaEnabledEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HaEnabledEnum) *HaEnabledEnum {
		return &v
	}).(HaEnabledEnumPtrOutput)
}

func (o HaEnabledEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HaEnabledEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HaEnabledEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HaEnabledEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HaEnabledEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HaEnabledEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HaEnabledEnumPtrOutput struct{ *pulumi.OutputState }

func (HaEnabledEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HaEnabledEnum)(nil)).Elem()
}

func (o HaEnabledEnumPtrOutput) ToHaEnabledEnumPtrOutput() HaEnabledEnumPtrOutput {
	return o
}

func (o HaEnabledEnumPtrOutput) ToHaEnabledEnumPtrOutputWithContext(ctx context.Context) HaEnabledEnumPtrOutput {
	return o
}

func (o HaEnabledEnumPtrOutput) Elem() HaEnabledEnumOutput {
	return o.ApplyT(func(v *HaEnabledEnum) HaEnabledEnum {
		if v != nil {
			return *v
		}
		var ret HaEnabledEnum
		return ret
	}).(HaEnabledEnumOutput)
}

func (o HaEnabledEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HaEnabledEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HaEnabledEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HaEnabledEnumInput interface {
	pulumi.Input

	ToHaEnabledEnumOutput() HaEnabledEnumOutput
	ToHaEnabledEnumOutputWithContext(context.Context) HaEnabledEnumOutput
}

var haEnabledEnumPtrType = reflect.TypeOf((**HaEnabledEnum)(nil)).Elem()

type HaEnabledEnumPtrInput interface {
	pulumi.Input

	ToHaEnabledEnumPtrOutput() HaEnabledEnumPtrOutput
	ToHaEnabledEnumPtrOutputWithContext(context.Context) HaEnabledEnumPtrOutput
}

type haEnabledEnumPtr string

func HaEnabledEnumPtr(v string) HaEnabledEnumPtrInput {
	return (*haEnabledEnumPtr)(&v)
}

func (*haEnabledEnumPtr) ElementType() reflect.Type {
	return haEnabledEnumPtrType
}

func (in *haEnabledEnumPtr) ToHaEnabledEnumPtrOutput() HaEnabledEnumPtrOutput {
	return pulumi.ToOutput(in).(HaEnabledEnumPtrOutput)
}

func (in *haEnabledEnumPtr) ToHaEnabledEnumPtrOutputWithContext(ctx context.Context) HaEnabledEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HaEnabledEnumPtrOutput)
}

type InfrastructureEncryptionEnum string

const (
	InfrastructureEncryptionEnumEnabled  = InfrastructureEncryptionEnum("Enabled")
	InfrastructureEncryptionEnumDisabled = InfrastructureEncryptionEnum("Disabled")
)

func (InfrastructureEncryptionEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*InfrastructureEncryptionEnum)(nil)).Elem()
}

func (e InfrastructureEncryptionEnum) ToInfrastructureEncryptionEnumOutput() InfrastructureEncryptionEnumOutput {
	return pulumi.ToOutput(e).(InfrastructureEncryptionEnumOutput)
}

func (e InfrastructureEncryptionEnum) ToInfrastructureEncryptionEnumOutputWithContext(ctx context.Context) InfrastructureEncryptionEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InfrastructureEncryptionEnumOutput)
}

func (e InfrastructureEncryptionEnum) ToInfrastructureEncryptionEnumPtrOutput() InfrastructureEncryptionEnumPtrOutput {
	return e.ToInfrastructureEncryptionEnumPtrOutputWithContext(context.Background())
}

func (e InfrastructureEncryptionEnum) ToInfrastructureEncryptionEnumPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionEnumPtrOutput {
	return InfrastructureEncryptionEnum(e).ToInfrastructureEncryptionEnumOutputWithContext(ctx).ToInfrastructureEncryptionEnumPtrOutputWithContext(ctx)
}

func (e InfrastructureEncryptionEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InfrastructureEncryptionEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InfrastructureEncryptionEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InfrastructureEncryptionEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InfrastructureEncryptionEnumOutput struct{ *pulumi.OutputState }

func (InfrastructureEncryptionEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InfrastructureEncryptionEnum)(nil)).Elem()
}

func (o InfrastructureEncryptionEnumOutput) ToInfrastructureEncryptionEnumOutput() InfrastructureEncryptionEnumOutput {
	return o
}

func (o InfrastructureEncryptionEnumOutput) ToInfrastructureEncryptionEnumOutputWithContext(ctx context.Context) InfrastructureEncryptionEnumOutput {
	return o
}

func (o InfrastructureEncryptionEnumOutput) ToInfrastructureEncryptionEnumPtrOutput() InfrastructureEncryptionEnumPtrOutput {
	return o.ToInfrastructureEncryptionEnumPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionEnumOutput) ToInfrastructureEncryptionEnumPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InfrastructureEncryptionEnum) *InfrastructureEncryptionEnum {
		return &v
	}).(InfrastructureEncryptionEnumPtrOutput)
}

func (o InfrastructureEncryptionEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InfrastructureEncryptionEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InfrastructureEncryptionEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InfrastructureEncryptionEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InfrastructureEncryptionEnumPtrOutput struct{ *pulumi.OutputState }

func (InfrastructureEncryptionEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InfrastructureEncryptionEnum)(nil)).Elem()
}

func (o InfrastructureEncryptionEnumPtrOutput) ToInfrastructureEncryptionEnumPtrOutput() InfrastructureEncryptionEnumPtrOutput {
	return o
}

func (o InfrastructureEncryptionEnumPtrOutput) ToInfrastructureEncryptionEnumPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionEnumPtrOutput {
	return o
}

func (o InfrastructureEncryptionEnumPtrOutput) Elem() InfrastructureEncryptionEnumOutput {
	return o.ApplyT(func(v *InfrastructureEncryptionEnum) InfrastructureEncryptionEnum {
		if v != nil {
			return *v
		}
		var ret InfrastructureEncryptionEnum
		return ret
	}).(InfrastructureEncryptionEnumOutput)
}

func (o InfrastructureEncryptionEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InfrastructureEncryptionEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InfrastructureEncryptionEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InfrastructureEncryptionEnumInput interface {
	pulumi.Input

	ToInfrastructureEncryptionEnumOutput() InfrastructureEncryptionEnumOutput
	ToInfrastructureEncryptionEnumOutputWithContext(context.Context) InfrastructureEncryptionEnumOutput
}

var infrastructureEncryptionEnumPtrType = reflect.TypeOf((**InfrastructureEncryptionEnum)(nil)).Elem()

type InfrastructureEncryptionEnumPtrInput interface {
	pulumi.Input

	ToInfrastructureEncryptionEnumPtrOutput() InfrastructureEncryptionEnumPtrOutput
	ToInfrastructureEncryptionEnumPtrOutputWithContext(context.Context) InfrastructureEncryptionEnumPtrOutput
}

type infrastructureEncryptionEnumPtr string

func InfrastructureEncryptionEnumPtr(v string) InfrastructureEncryptionEnumPtrInput {
	return (*infrastructureEncryptionEnumPtr)(&v)
}

func (*infrastructureEncryptionEnumPtr) ElementType() reflect.Type {
	return infrastructureEncryptionEnumPtrType
}

func (in *infrastructureEncryptionEnumPtr) ToInfrastructureEncryptionEnumPtrOutput() InfrastructureEncryptionEnumPtrOutput {
	return pulumi.ToOutput(in).(InfrastructureEncryptionEnumPtrOutput)
}

func (in *infrastructureEncryptionEnumPtr) ToInfrastructureEncryptionEnumPtrOutputWithContext(ctx context.Context) InfrastructureEncryptionEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InfrastructureEncryptionEnumPtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

type ServerKeyType string

const (
	ServerKeyTypeAzureKeyVault = ServerKeyType("AzureKeyVault")
)

func (ServerKeyType) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerKeyType)(nil)).Elem()
}

func (e ServerKeyType) ToServerKeyTypeOutput() ServerKeyTypeOutput {
	return pulumi.ToOutput(e).(ServerKeyTypeOutput)
}

func (e ServerKeyType) ToServerKeyTypeOutputWithContext(ctx context.Context) ServerKeyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServerKeyTypeOutput)
}

func (e ServerKeyType) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return e.ToServerKeyTypePtrOutputWithContext(context.Background())
}

func (e ServerKeyType) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return ServerKeyType(e).ToServerKeyTypeOutputWithContext(ctx).ToServerKeyTypePtrOutputWithContext(ctx)
}

func (e ServerKeyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerKeyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerKeyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServerKeyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServerKeyTypeOutput struct{ *pulumi.OutputState }

func (ServerKeyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerKeyType)(nil)).Elem()
}

func (o ServerKeyTypeOutput) ToServerKeyTypeOutput() ServerKeyTypeOutput {
	return o
}

func (o ServerKeyTypeOutput) ToServerKeyTypeOutputWithContext(ctx context.Context) ServerKeyTypeOutput {
	return o
}

func (o ServerKeyTypeOutput) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return o.ToServerKeyTypePtrOutputWithContext(context.Background())
}

func (o ServerKeyTypeOutput) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerKeyType) *ServerKeyType {
		return &v
	}).(ServerKeyTypePtrOutput)
}

func (o ServerKeyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServerKeyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerKeyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServerKeyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerKeyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerKeyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServerKeyTypePtrOutput struct{ *pulumi.OutputState }

func (ServerKeyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerKeyType)(nil)).Elem()
}

func (o ServerKeyTypePtrOutput) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return o
}

func (o ServerKeyTypePtrOutput) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return o
}

func (o ServerKeyTypePtrOutput) Elem() ServerKeyTypeOutput {
	return o.ApplyT(func(v *ServerKeyType) ServerKeyType {
		if v != nil {
			return *v
		}
		var ret ServerKeyType
		return ret
	}).(ServerKeyTypeOutput)
}

func (o ServerKeyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerKeyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServerKeyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServerKeyTypeInput interface {
	pulumi.Input

	ToServerKeyTypeOutput() ServerKeyTypeOutput
	ToServerKeyTypeOutputWithContext(context.Context) ServerKeyTypeOutput
}

var serverKeyTypePtrType = reflect.TypeOf((**ServerKeyType)(nil)).Elem()

type ServerKeyTypePtrInput interface {
	pulumi.Input

	ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput
	ToServerKeyTypePtrOutputWithContext(context.Context) ServerKeyTypePtrOutput
}

type serverKeyTypePtr string

func ServerKeyTypePtr(v string) ServerKeyTypePtrInput {
	return (*serverKeyTypePtr)(&v)
}

func (*serverKeyTypePtr) ElementType() reflect.Type {
	return serverKeyTypePtrType
}

func (in *serverKeyTypePtr) ToServerKeyTypePtrOutput() ServerKeyTypePtrOutput {
	return pulumi.ToOutput(in).(ServerKeyTypePtrOutput)
}

func (in *serverKeyTypePtr) ToServerKeyTypePtrOutputWithContext(ctx context.Context) ServerKeyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServerKeyTypePtrOutput)
}

type ServerVersion string

const (
	ServerVersion_5_7 = ServerVersion("5.7")
)

func (ServerVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerVersion)(nil)).Elem()
}

func (e ServerVersion) ToServerVersionOutput() ServerVersionOutput {
	return pulumi.ToOutput(e).(ServerVersionOutput)
}

func (e ServerVersion) ToServerVersionOutputWithContext(ctx context.Context) ServerVersionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ServerVersionOutput)
}

func (e ServerVersion) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return e.ToServerVersionPtrOutputWithContext(context.Background())
}

func (e ServerVersion) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return ServerVersion(e).ToServerVersionOutputWithContext(ctx).ToServerVersionPtrOutputWithContext(ctx)
}

func (e ServerVersion) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerVersion) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ServerVersion) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ServerVersion) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ServerVersionOutput struct{ *pulumi.OutputState }

func (ServerVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerVersion)(nil)).Elem()
}

func (o ServerVersionOutput) ToServerVersionOutput() ServerVersionOutput {
	return o
}

func (o ServerVersionOutput) ToServerVersionOutputWithContext(ctx context.Context) ServerVersionOutput {
	return o
}

func (o ServerVersionOutput) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return o.ToServerVersionPtrOutputWithContext(context.Background())
}

func (o ServerVersionOutput) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerVersion) *ServerVersion {
		return &v
	}).(ServerVersionPtrOutput)
}

func (o ServerVersionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ServerVersionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerVersion) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ServerVersionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerVersionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ServerVersion) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ServerVersionPtrOutput struct{ *pulumi.OutputState }

func (ServerVersionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerVersion)(nil)).Elem()
}

func (o ServerVersionPtrOutput) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return o
}

func (o ServerVersionPtrOutput) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return o
}

func (o ServerVersionPtrOutput) Elem() ServerVersionOutput {
	return o.ApplyT(func(v *ServerVersion) ServerVersion {
		if v != nil {
			return *v
		}
		var ret ServerVersion
		return ret
	}).(ServerVersionOutput)
}

func (o ServerVersionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ServerVersionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ServerVersion) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ServerVersionInput interface {
	pulumi.Input

	ToServerVersionOutput() ServerVersionOutput
	ToServerVersionOutputWithContext(context.Context) ServerVersionOutput
}

var serverVersionPtrType = reflect.TypeOf((**ServerVersion)(nil)).Elem()

type ServerVersionPtrInput interface {
	pulumi.Input

	ToServerVersionPtrOutput() ServerVersionPtrOutput
	ToServerVersionPtrOutputWithContext(context.Context) ServerVersionPtrOutput
}

type serverVersionPtr string

func ServerVersionPtr(v string) ServerVersionPtrInput {
	return (*serverVersionPtr)(&v)
}

func (*serverVersionPtr) ElementType() reflect.Type {
	return serverVersionPtrType
}

func (in *serverVersionPtr) ToServerVersionPtrOutput() ServerVersionPtrOutput {
	return pulumi.ToOutput(in).(ServerVersionPtrOutput)
}

func (in *serverVersionPtr) ToServerVersionPtrOutputWithContext(ctx context.Context) ServerVersionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ServerVersionPtrOutput)
}

type SkuTier string

const (
	SkuTierBurstable       = SkuTier("Burstable")
	SkuTierGeneralPurpose  = SkuTier("GeneralPurpose")
	SkuTierMemoryOptimized = SkuTier("MemoryOptimized")
)

func (SkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (e SkuTier) ToSkuTierOutput() SkuTierOutput {
	return pulumi.ToOutput(e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuTierOutput)
}

func (e SkuTier) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return e.ToSkuTierPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return SkuTier(e).ToSkuTierOutputWithContext(ctx).ToSkuTierPtrOutputWithContext(ctx)
}

func (e SkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuTierOutput struct{ *pulumi.OutputState }

func (SkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuTier)(nil)).Elem()
}

func (o SkuTierOutput) ToSkuTierOutput() SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierOutputWithContext(ctx context.Context) SkuTierOutput {
	return o
}

func (o SkuTierOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o.ToSkuTierPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuTier) *SkuTier {
		return &v
	}).(SkuTierPtrOutput)
}

func (o SkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuTierPtrOutput struct{ *pulumi.OutputState }

func (SkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuTier)(nil)).Elem()
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return o
}

func (o SkuTierPtrOutput) Elem() SkuTierOutput {
	return o.ApplyT(func(v *SkuTier) SkuTier {
		if v != nil {
			return *v
		}
		var ret SkuTier
		return ret
	}).(SkuTierOutput)
}

func (o SkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuTierInput interface {
	pulumi.Input

	ToSkuTierOutput() SkuTierOutput
	ToSkuTierOutputWithContext(context.Context) SkuTierOutput
}

var skuTierPtrType = reflect.TypeOf((**SkuTier)(nil)).Elem()

type SkuTierPtrInput interface {
	pulumi.Input

	ToSkuTierPtrOutput() SkuTierPtrOutput
	ToSkuTierPtrOutputWithContext(context.Context) SkuTierPtrOutput
}

type skuTierPtr string

func SkuTierPtr(v string) SkuTierPtrInput {
	return (*skuTierPtr)(&v)
}

func (*skuTierPtr) ElementType() reflect.Type {
	return skuTierPtrType
}

func (in *skuTierPtr) ToSkuTierPtrOutput() SkuTierPtrOutput {
	return pulumi.ToOutput(in).(SkuTierPtrOutput)
}

func (in *skuTierPtr) ToSkuTierPtrOutputWithContext(ctx context.Context) SkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuTierPtrOutput)
}

type SslEnforcementEnum string

const (
	SslEnforcementEnumEnabled  = SslEnforcementEnum("Enabled")
	SslEnforcementEnumDisabled = SslEnforcementEnum("Disabled")
)

func (SslEnforcementEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SslEnforcementEnum)(nil)).Elem()
}

func (e SslEnforcementEnum) ToSslEnforcementEnumOutput() SslEnforcementEnumOutput {
	return pulumi.ToOutput(e).(SslEnforcementEnumOutput)
}

func (e SslEnforcementEnum) ToSslEnforcementEnumOutputWithContext(ctx context.Context) SslEnforcementEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SslEnforcementEnumOutput)
}

func (e SslEnforcementEnum) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return e.ToSslEnforcementEnumPtrOutputWithContext(context.Background())
}

func (e SslEnforcementEnum) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return SslEnforcementEnum(e).ToSslEnforcementEnumOutputWithContext(ctx).ToSslEnforcementEnumPtrOutputWithContext(ctx)
}

func (e SslEnforcementEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslEnforcementEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SslEnforcementEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SslEnforcementEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SslEnforcementEnumOutput struct{ *pulumi.OutputState }

func (SslEnforcementEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SslEnforcementEnum)(nil)).Elem()
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumOutput() SslEnforcementEnumOutput {
	return o
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumOutputWithContext(ctx context.Context) SslEnforcementEnumOutput {
	return o
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return o.ToSslEnforcementEnumPtrOutputWithContext(context.Background())
}

func (o SslEnforcementEnumOutput) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SslEnforcementEnum) *SslEnforcementEnum {
		return &v
	}).(SslEnforcementEnumPtrOutput)
}

func (o SslEnforcementEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SslEnforcementEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslEnforcementEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SslEnforcementEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslEnforcementEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SslEnforcementEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SslEnforcementEnumPtrOutput struct{ *pulumi.OutputState }

func (SslEnforcementEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SslEnforcementEnum)(nil)).Elem()
}

func (o SslEnforcementEnumPtrOutput) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return o
}

func (o SslEnforcementEnumPtrOutput) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return o
}

func (o SslEnforcementEnumPtrOutput) Elem() SslEnforcementEnumOutput {
	return o.ApplyT(func(v *SslEnforcementEnum) SslEnforcementEnum {
		if v != nil {
			return *v
		}
		var ret SslEnforcementEnum
		return ret
	}).(SslEnforcementEnumOutput)
}

func (o SslEnforcementEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SslEnforcementEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SslEnforcementEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SslEnforcementEnumInput interface {
	pulumi.Input

	ToSslEnforcementEnumOutput() SslEnforcementEnumOutput
	ToSslEnforcementEnumOutputWithContext(context.Context) SslEnforcementEnumOutput
}

var sslEnforcementEnumPtrType = reflect.TypeOf((**SslEnforcementEnum)(nil)).Elem()

type SslEnforcementEnumPtrInput interface {
	pulumi.Input

	ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput
	ToSslEnforcementEnumPtrOutputWithContext(context.Context) SslEnforcementEnumPtrOutput
}

type sslEnforcementEnumPtr string

func SslEnforcementEnumPtr(v string) SslEnforcementEnumPtrInput {
	return (*sslEnforcementEnumPtr)(&v)
}

func (*sslEnforcementEnumPtr) ElementType() reflect.Type {
	return sslEnforcementEnumPtrType
}

func (in *sslEnforcementEnumPtr) ToSslEnforcementEnumPtrOutput() SslEnforcementEnumPtrOutput {
	return pulumi.ToOutput(in).(SslEnforcementEnumPtrOutput)
}

func (in *sslEnforcementEnumPtr) ToSslEnforcementEnumPtrOutputWithContext(ctx context.Context) SslEnforcementEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SslEnforcementEnumPtrOutput)
}

type StorageAutogrow string

const (
	StorageAutogrowEnabled  = StorageAutogrow("Enabled")
	StorageAutogrowDisabled = StorageAutogrow("Disabled")
)

func (StorageAutogrow) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAutogrow)(nil)).Elem()
}

func (e StorageAutogrow) ToStorageAutogrowOutput() StorageAutogrowOutput {
	return pulumi.ToOutput(e).(StorageAutogrowOutput)
}

func (e StorageAutogrow) ToStorageAutogrowOutputWithContext(ctx context.Context) StorageAutogrowOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StorageAutogrowOutput)
}

func (e StorageAutogrow) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return e.ToStorageAutogrowPtrOutputWithContext(context.Background())
}

func (e StorageAutogrow) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return StorageAutogrow(e).ToStorageAutogrowOutputWithContext(ctx).ToStorageAutogrowPtrOutputWithContext(ctx)
}

func (e StorageAutogrow) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAutogrow) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StorageAutogrow) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StorageAutogrow) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StorageAutogrowOutput struct{ *pulumi.OutputState }

func (StorageAutogrowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAutogrow)(nil)).Elem()
}

func (o StorageAutogrowOutput) ToStorageAutogrowOutput() StorageAutogrowOutput {
	return o
}

func (o StorageAutogrowOutput) ToStorageAutogrowOutputWithContext(ctx context.Context) StorageAutogrowOutput {
	return o
}

func (o StorageAutogrowOutput) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return o.ToStorageAutogrowPtrOutputWithContext(context.Background())
}

func (o StorageAutogrowOutput) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAutogrow) *StorageAutogrow {
		return &v
	}).(StorageAutogrowPtrOutput)
}

func (o StorageAutogrowOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StorageAutogrowOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAutogrow) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StorageAutogrowOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAutogrowOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StorageAutogrow) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StorageAutogrowPtrOutput struct{ *pulumi.OutputState }

func (StorageAutogrowPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAutogrow)(nil)).Elem()
}

func (o StorageAutogrowPtrOutput) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return o
}

func (o StorageAutogrowPtrOutput) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return o
}

func (o StorageAutogrowPtrOutput) Elem() StorageAutogrowOutput {
	return o.ApplyT(func(v *StorageAutogrow) StorageAutogrow {
		if v != nil {
			return *v
		}
		var ret StorageAutogrow
		return ret
	}).(StorageAutogrowOutput)
}

func (o StorageAutogrowPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StorageAutogrowPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StorageAutogrow) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StorageAutogrowInput interface {
	pulumi.Input

	ToStorageAutogrowOutput() StorageAutogrowOutput
	ToStorageAutogrowOutputWithContext(context.Context) StorageAutogrowOutput
}

var storageAutogrowPtrType = reflect.TypeOf((**StorageAutogrow)(nil)).Elem()

type StorageAutogrowPtrInput interface {
	pulumi.Input

	ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput
	ToStorageAutogrowPtrOutputWithContext(context.Context) StorageAutogrowPtrOutput
}

type storageAutogrowPtr string

func StorageAutogrowPtr(v string) StorageAutogrowPtrInput {
	return (*storageAutogrowPtr)(&v)
}

func (*storageAutogrowPtr) ElementType() reflect.Type {
	return storageAutogrowPtrType
}

func (in *storageAutogrowPtr) ToStorageAutogrowPtrOutput() StorageAutogrowPtrOutput {
	return pulumi.ToOutput(in).(StorageAutogrowPtrOutput)
}

func (in *storageAutogrowPtr) ToStorageAutogrowPtrOutputWithContext(ctx context.Context) StorageAutogrowPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StorageAutogrowPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CreateModeInput)(nil)).Elem(), CreateMode("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*CreateModePtrInput)(nil)).Elem(), CreateMode("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*HaEnabledEnumInput)(nil)).Elem(), HaEnabledEnum("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*HaEnabledEnumPtrInput)(nil)).Elem(), HaEnabledEnum("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*InfrastructureEncryptionEnumInput)(nil)).Elem(), InfrastructureEncryptionEnum("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*InfrastructureEncryptionEnumPtrInput)(nil)).Elem(), InfrastructureEncryptionEnum("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServerKeyTypeInput)(nil)).Elem(), ServerKeyType("AzureKeyVault"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServerKeyTypePtrInput)(nil)).Elem(), ServerKeyType("AzureKeyVault"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServerVersionInput)(nil)).Elem(), ServerVersion("5.7"))
	pulumi.RegisterInputType(reflect.TypeOf((*ServerVersionPtrInput)(nil)).Elem(), ServerVersion("5.7"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuTierInput)(nil)).Elem(), SkuTier("Burstable"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuTierPtrInput)(nil)).Elem(), SkuTier("Burstable"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslEnforcementEnumInput)(nil)).Elem(), SslEnforcementEnum("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*SslEnforcementEnumPtrInput)(nil)).Elem(), SslEnforcementEnum("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAutogrowInput)(nil)).Elem(), StorageAutogrow("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*StorageAutogrowPtrInput)(nil)).Elem(), StorageAutogrow("Enabled"))
	pulumi.RegisterOutputType(CreateModeOutput{})
	pulumi.RegisterOutputType(CreateModePtrOutput{})
	pulumi.RegisterOutputType(HaEnabledEnumOutput{})
	pulumi.RegisterOutputType(HaEnabledEnumPtrOutput{})
	pulumi.RegisterOutputType(InfrastructureEncryptionEnumOutput{})
	pulumi.RegisterOutputType(InfrastructureEncryptionEnumPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(ServerKeyTypeOutput{})
	pulumi.RegisterOutputType(ServerKeyTypePtrOutput{})
	pulumi.RegisterOutputType(ServerVersionOutput{})
	pulumi.RegisterOutputType(ServerVersionPtrOutput{})
	pulumi.RegisterOutputType(SkuTierOutput{})
	pulumi.RegisterOutputType(SkuTierPtrOutput{})
	pulumi.RegisterOutputType(SslEnforcementEnumOutput{})
	pulumi.RegisterOutputType(SslEnforcementEnumPtrOutput{})
	pulumi.RegisterOutputType(StorageAutogrowOutput{})
	pulumi.RegisterOutputType(StorageAutogrowPtrOutput{})
}
