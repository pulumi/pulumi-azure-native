


package v20180712

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnterpriseChannelNodeState string

const (
	EnterpriseChannelNodeStateCreating     = EnterpriseChannelNodeState("Creating")
	EnterpriseChannelNodeStateCreateFailed = EnterpriseChannelNodeState("CreateFailed")
	EnterpriseChannelNodeStateStarted      = EnterpriseChannelNodeState("Started")
	EnterpriseChannelNodeStateStarting     = EnterpriseChannelNodeState("Starting")
	EnterpriseChannelNodeStateStartFailed  = EnterpriseChannelNodeState("StartFailed")
	EnterpriseChannelNodeStateStopped      = EnterpriseChannelNodeState("Stopped")
	EnterpriseChannelNodeStateStopping     = EnterpriseChannelNodeState("Stopping")
	EnterpriseChannelNodeStateStopFailed   = EnterpriseChannelNodeState("StopFailed")
	EnterpriseChannelNodeStateDeleting     = EnterpriseChannelNodeState("Deleting")
	EnterpriseChannelNodeStateDeleteFailed = EnterpriseChannelNodeState("DeleteFailed")
)

func (EnterpriseChannelNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannelNodeState)(nil)).Elem()
}

func (e EnterpriseChannelNodeState) ToEnterpriseChannelNodeStateOutput() EnterpriseChannelNodeStateOutput {
	return pulumi.ToOutput(e).(EnterpriseChannelNodeStateOutput)
}

func (e EnterpriseChannelNodeState) ToEnterpriseChannelNodeStateOutputWithContext(ctx context.Context) EnterpriseChannelNodeStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnterpriseChannelNodeStateOutput)
}

func (e EnterpriseChannelNodeState) ToEnterpriseChannelNodeStatePtrOutput() EnterpriseChannelNodeStatePtrOutput {
	return e.ToEnterpriseChannelNodeStatePtrOutputWithContext(context.Background())
}

func (e EnterpriseChannelNodeState) ToEnterpriseChannelNodeStatePtrOutputWithContext(ctx context.Context) EnterpriseChannelNodeStatePtrOutput {
	return EnterpriseChannelNodeState(e).ToEnterpriseChannelNodeStateOutputWithContext(ctx).ToEnterpriseChannelNodeStatePtrOutputWithContext(ctx)
}

func (e EnterpriseChannelNodeState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnterpriseChannelNodeState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnterpriseChannelNodeState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnterpriseChannelNodeState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnterpriseChannelNodeStateOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelNodeStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannelNodeState)(nil)).Elem()
}

func (o EnterpriseChannelNodeStateOutput) ToEnterpriseChannelNodeStateOutput() EnterpriseChannelNodeStateOutput {
	return o
}

func (o EnterpriseChannelNodeStateOutput) ToEnterpriseChannelNodeStateOutputWithContext(ctx context.Context) EnterpriseChannelNodeStateOutput {
	return o
}

func (o EnterpriseChannelNodeStateOutput) ToEnterpriseChannelNodeStatePtrOutput() EnterpriseChannelNodeStatePtrOutput {
	return o.ToEnterpriseChannelNodeStatePtrOutputWithContext(context.Background())
}

func (o EnterpriseChannelNodeStateOutput) ToEnterpriseChannelNodeStatePtrOutputWithContext(ctx context.Context) EnterpriseChannelNodeStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnterpriseChannelNodeState) *EnterpriseChannelNodeState {
		return &v
	}).(EnterpriseChannelNodeStatePtrOutput)
}

func (o EnterpriseChannelNodeStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnterpriseChannelNodeStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnterpriseChannelNodeState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnterpriseChannelNodeStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnterpriseChannelNodeStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnterpriseChannelNodeState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnterpriseChannelNodeStatePtrOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelNodeStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseChannelNodeState)(nil)).Elem()
}

func (o EnterpriseChannelNodeStatePtrOutput) ToEnterpriseChannelNodeStatePtrOutput() EnterpriseChannelNodeStatePtrOutput {
	return o
}

func (o EnterpriseChannelNodeStatePtrOutput) ToEnterpriseChannelNodeStatePtrOutputWithContext(ctx context.Context) EnterpriseChannelNodeStatePtrOutput {
	return o
}

func (o EnterpriseChannelNodeStatePtrOutput) Elem() EnterpriseChannelNodeStateOutput {
	return o.ApplyT(func(v *EnterpriseChannelNodeState) EnterpriseChannelNodeState {
		if v != nil {
			return *v
		}
		var ret EnterpriseChannelNodeState
		return ret
	}).(EnterpriseChannelNodeStateOutput)
}

func (o EnterpriseChannelNodeStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnterpriseChannelNodeStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnterpriseChannelNodeState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EnterpriseChannelNodeStateInput interface {
	pulumi.Input

	ToEnterpriseChannelNodeStateOutput() EnterpriseChannelNodeStateOutput
	ToEnterpriseChannelNodeStateOutputWithContext(context.Context) EnterpriseChannelNodeStateOutput
}

var enterpriseChannelNodeStatePtrType = reflect.TypeOf((**EnterpriseChannelNodeState)(nil)).Elem()

type EnterpriseChannelNodeStatePtrInput interface {
	pulumi.Input

	ToEnterpriseChannelNodeStatePtrOutput() EnterpriseChannelNodeStatePtrOutput
	ToEnterpriseChannelNodeStatePtrOutputWithContext(context.Context) EnterpriseChannelNodeStatePtrOutput
}

type enterpriseChannelNodeStatePtr string

func EnterpriseChannelNodeStatePtr(v string) EnterpriseChannelNodeStatePtrInput {
	return (*enterpriseChannelNodeStatePtr)(&v)
}

func (*enterpriseChannelNodeStatePtr) ElementType() reflect.Type {
	return enterpriseChannelNodeStatePtrType
}

func (in *enterpriseChannelNodeStatePtr) ToEnterpriseChannelNodeStatePtrOutput() EnterpriseChannelNodeStatePtrOutput {
	return pulumi.ToOutput(in).(EnterpriseChannelNodeStatePtrOutput)
}

func (in *enterpriseChannelNodeStatePtr) ToEnterpriseChannelNodeStatePtrOutputWithContext(ctx context.Context) EnterpriseChannelNodeStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnterpriseChannelNodeStatePtrOutput)
}

type EnterpriseChannelStateEnum string

const (
	EnterpriseChannelStateEnumCreating     = EnterpriseChannelStateEnum("Creating")
	EnterpriseChannelStateEnumCreateFailed = EnterpriseChannelStateEnum("CreateFailed")
	EnterpriseChannelStateEnumStarted      = EnterpriseChannelStateEnum("Started")
	EnterpriseChannelStateEnumStarting     = EnterpriseChannelStateEnum("Starting")
	EnterpriseChannelStateEnumStartFailed  = EnterpriseChannelStateEnum("StartFailed")
	EnterpriseChannelStateEnumStopped      = EnterpriseChannelStateEnum("Stopped")
	EnterpriseChannelStateEnumStopping     = EnterpriseChannelStateEnum("Stopping")
	EnterpriseChannelStateEnumStopFailed   = EnterpriseChannelStateEnum("StopFailed")
	EnterpriseChannelStateEnumDeleting     = EnterpriseChannelStateEnum("Deleting")
	EnterpriseChannelStateEnumDeleteFailed = EnterpriseChannelStateEnum("DeleteFailed")
)

func (EnterpriseChannelStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannelStateEnum)(nil)).Elem()
}

func (e EnterpriseChannelStateEnum) ToEnterpriseChannelStateEnumOutput() EnterpriseChannelStateEnumOutput {
	return pulumi.ToOutput(e).(EnterpriseChannelStateEnumOutput)
}

func (e EnterpriseChannelStateEnum) ToEnterpriseChannelStateEnumOutputWithContext(ctx context.Context) EnterpriseChannelStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnterpriseChannelStateEnumOutput)
}

func (e EnterpriseChannelStateEnum) ToEnterpriseChannelStateEnumPtrOutput() EnterpriseChannelStateEnumPtrOutput {
	return e.ToEnterpriseChannelStateEnumPtrOutputWithContext(context.Background())
}

func (e EnterpriseChannelStateEnum) ToEnterpriseChannelStateEnumPtrOutputWithContext(ctx context.Context) EnterpriseChannelStateEnumPtrOutput {
	return EnterpriseChannelStateEnum(e).ToEnterpriseChannelStateEnumOutputWithContext(ctx).ToEnterpriseChannelStateEnumPtrOutputWithContext(ctx)
}

func (e EnterpriseChannelStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnterpriseChannelStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnterpriseChannelStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnterpriseChannelStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnterpriseChannelStateEnumOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannelStateEnum)(nil)).Elem()
}

func (o EnterpriseChannelStateEnumOutput) ToEnterpriseChannelStateEnumOutput() EnterpriseChannelStateEnumOutput {
	return o
}

func (o EnterpriseChannelStateEnumOutput) ToEnterpriseChannelStateEnumOutputWithContext(ctx context.Context) EnterpriseChannelStateEnumOutput {
	return o
}

func (o EnterpriseChannelStateEnumOutput) ToEnterpriseChannelStateEnumPtrOutput() EnterpriseChannelStateEnumPtrOutput {
	return o.ToEnterpriseChannelStateEnumPtrOutputWithContext(context.Background())
}

func (o EnterpriseChannelStateEnumOutput) ToEnterpriseChannelStateEnumPtrOutputWithContext(ctx context.Context) EnterpriseChannelStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnterpriseChannelStateEnum) *EnterpriseChannelStateEnum {
		return &v
	}).(EnterpriseChannelStateEnumPtrOutput)
}

func (o EnterpriseChannelStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnterpriseChannelStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnterpriseChannelStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnterpriseChannelStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnterpriseChannelStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnterpriseChannelStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnterpriseChannelStateEnumPtrOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseChannelStateEnum)(nil)).Elem()
}

func (o EnterpriseChannelStateEnumPtrOutput) ToEnterpriseChannelStateEnumPtrOutput() EnterpriseChannelStateEnumPtrOutput {
	return o
}

func (o EnterpriseChannelStateEnumPtrOutput) ToEnterpriseChannelStateEnumPtrOutputWithContext(ctx context.Context) EnterpriseChannelStateEnumPtrOutput {
	return o
}

func (o EnterpriseChannelStateEnumPtrOutput) Elem() EnterpriseChannelStateEnumOutput {
	return o.ApplyT(func(v *EnterpriseChannelStateEnum) EnterpriseChannelStateEnum {
		if v != nil {
			return *v
		}
		var ret EnterpriseChannelStateEnum
		return ret
	}).(EnterpriseChannelStateEnumOutput)
}

func (o EnterpriseChannelStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnterpriseChannelStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnterpriseChannelStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EnterpriseChannelStateEnumInput interface {
	pulumi.Input

	ToEnterpriseChannelStateEnumOutput() EnterpriseChannelStateEnumOutput
	ToEnterpriseChannelStateEnumOutputWithContext(context.Context) EnterpriseChannelStateEnumOutput
}

var enterpriseChannelStateEnumPtrType = reflect.TypeOf((**EnterpriseChannelStateEnum)(nil)).Elem()

type EnterpriseChannelStateEnumPtrInput interface {
	pulumi.Input

	ToEnterpriseChannelStateEnumPtrOutput() EnterpriseChannelStateEnumPtrOutput
	ToEnterpriseChannelStateEnumPtrOutputWithContext(context.Context) EnterpriseChannelStateEnumPtrOutput
}

type enterpriseChannelStateEnumPtr string

func EnterpriseChannelStateEnumPtr(v string) EnterpriseChannelStateEnumPtrInput {
	return (*enterpriseChannelStateEnumPtr)(&v)
}

func (*enterpriseChannelStateEnumPtr) ElementType() reflect.Type {
	return enterpriseChannelStateEnumPtrType
}

func (in *enterpriseChannelStateEnumPtr) ToEnterpriseChannelStateEnumPtrOutput() EnterpriseChannelStateEnumPtrOutput {
	return pulumi.ToOutput(in).(EnterpriseChannelStateEnumPtrOutput)
}

func (in *enterpriseChannelStateEnumPtr) ToEnterpriseChannelStateEnumPtrOutputWithContext(ctx context.Context) EnterpriseChannelStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnterpriseChannelStateEnumPtrOutput)
}

type Kind string

const (
	KindSdk      = Kind("sdk")
	KindDesigner = Kind("designer")
	KindBot      = Kind("bot")
	KindFunction = Kind("function")
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

type SkuName string

const (
	SkuNameF0 = SkuName("F0")
	SkuNameS1 = SkuName("S1")
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
	pulumi.RegisterOutputType(EnterpriseChannelNodeStateOutput{})
	pulumi.RegisterOutputType(EnterpriseChannelNodeStatePtrOutput{})
	pulumi.RegisterOutputType(EnterpriseChannelStateEnumOutput{})
	pulumi.RegisterOutputType(EnterpriseChannelStateEnumPtrOutput{})
	pulumi.RegisterOutputType(KindOutput{})
	pulumi.RegisterOutputType(KindPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
