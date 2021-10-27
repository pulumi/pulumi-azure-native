


package v20171115

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessRightsDescription string

const (
	AccessRightsDescriptionServiceConfig           = AccessRightsDescription("ServiceConfig")
	AccessRightsDescriptionEnrollmentRead          = AccessRightsDescription("EnrollmentRead")
	AccessRightsDescriptionEnrollmentWrite         = AccessRightsDescription("EnrollmentWrite")
	AccessRightsDescriptionDeviceConnect           = AccessRightsDescription("DeviceConnect")
	AccessRightsDescriptionRegistrationStatusRead  = AccessRightsDescription("RegistrationStatusRead")
	AccessRightsDescriptionRegistrationStatusWrite = AccessRightsDescription("RegistrationStatusWrite")
)

func (AccessRightsDescription) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRightsDescription)(nil)).Elem()
}

func (e AccessRightsDescription) ToAccessRightsDescriptionOutput() AccessRightsDescriptionOutput {
	return pulumi.ToOutput(e).(AccessRightsDescriptionOutput)
}

func (e AccessRightsDescription) ToAccessRightsDescriptionOutputWithContext(ctx context.Context) AccessRightsDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessRightsDescriptionOutput)
}

func (e AccessRightsDescription) ToAccessRightsDescriptionPtrOutput() AccessRightsDescriptionPtrOutput {
	return e.ToAccessRightsDescriptionPtrOutputWithContext(context.Background())
}

func (e AccessRightsDescription) ToAccessRightsDescriptionPtrOutputWithContext(ctx context.Context) AccessRightsDescriptionPtrOutput {
	return AccessRightsDescription(e).ToAccessRightsDescriptionOutputWithContext(ctx).ToAccessRightsDescriptionPtrOutputWithContext(ctx)
}

func (e AccessRightsDescription) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRightsDescription) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRightsDescription) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessRightsDescription) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessRightsDescriptionOutput struct{ *pulumi.OutputState }

func (AccessRightsDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRightsDescription)(nil)).Elem()
}

func (o AccessRightsDescriptionOutput) ToAccessRightsDescriptionOutput() AccessRightsDescriptionOutput {
	return o
}

func (o AccessRightsDescriptionOutput) ToAccessRightsDescriptionOutputWithContext(ctx context.Context) AccessRightsDescriptionOutput {
	return o
}

func (o AccessRightsDescriptionOutput) ToAccessRightsDescriptionPtrOutput() AccessRightsDescriptionPtrOutput {
	return o.ToAccessRightsDescriptionPtrOutputWithContext(context.Background())
}

func (o AccessRightsDescriptionOutput) ToAccessRightsDescriptionPtrOutputWithContext(ctx context.Context) AccessRightsDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessRightsDescription) *AccessRightsDescription {
		return &v
	}).(AccessRightsDescriptionPtrOutput)
}

func (o AccessRightsDescriptionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessRightsDescriptionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRightsDescription) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessRightsDescriptionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRightsDescriptionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRightsDescription) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessRightsDescriptionPtrOutput struct{ *pulumi.OutputState }

func (AccessRightsDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRightsDescription)(nil)).Elem()
}

func (o AccessRightsDescriptionPtrOutput) ToAccessRightsDescriptionPtrOutput() AccessRightsDescriptionPtrOutput {
	return o
}

func (o AccessRightsDescriptionPtrOutput) ToAccessRightsDescriptionPtrOutputWithContext(ctx context.Context) AccessRightsDescriptionPtrOutput {
	return o
}

func (o AccessRightsDescriptionPtrOutput) Elem() AccessRightsDescriptionOutput {
	return o.ApplyT(func(v *AccessRightsDescription) AccessRightsDescription {
		if v != nil {
			return *v
		}
		var ret AccessRightsDescription
		return ret
	}).(AccessRightsDescriptionOutput)
}

func (o AccessRightsDescriptionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRightsDescriptionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessRightsDescription) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessRightsDescriptionInput interface {
	pulumi.Input

	ToAccessRightsDescriptionOutput() AccessRightsDescriptionOutput
	ToAccessRightsDescriptionOutputWithContext(context.Context) AccessRightsDescriptionOutput
}

var accessRightsDescriptionPtrType = reflect.TypeOf((**AccessRightsDescription)(nil)).Elem()

type AccessRightsDescriptionPtrInput interface {
	pulumi.Input

	ToAccessRightsDescriptionPtrOutput() AccessRightsDescriptionPtrOutput
	ToAccessRightsDescriptionPtrOutputWithContext(context.Context) AccessRightsDescriptionPtrOutput
}

type accessRightsDescriptionPtr string

func AccessRightsDescriptionPtr(v string) AccessRightsDescriptionPtrInput {
	return (*accessRightsDescriptionPtr)(&v)
}

func (*accessRightsDescriptionPtr) ElementType() reflect.Type {
	return accessRightsDescriptionPtrType
}

func (in *accessRightsDescriptionPtr) ToAccessRightsDescriptionPtrOutput() AccessRightsDescriptionPtrOutput {
	return pulumi.ToOutput(in).(AccessRightsDescriptionPtrOutput)
}

func (in *accessRightsDescriptionPtr) ToAccessRightsDescriptionPtrOutputWithContext(ctx context.Context) AccessRightsDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessRightsDescriptionPtrOutput)
}

type AllocationPolicy string

const (
	AllocationPolicyHashed     = AllocationPolicy("Hashed")
	AllocationPolicyGeoLatency = AllocationPolicy("GeoLatency")
	AllocationPolicyStatic     = AllocationPolicy("Static")
)

func (AllocationPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*AllocationPolicy)(nil)).Elem()
}

func (e AllocationPolicy) ToAllocationPolicyOutput() AllocationPolicyOutput {
	return pulumi.ToOutput(e).(AllocationPolicyOutput)
}

func (e AllocationPolicy) ToAllocationPolicyOutputWithContext(ctx context.Context) AllocationPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AllocationPolicyOutput)
}

func (e AllocationPolicy) ToAllocationPolicyPtrOutput() AllocationPolicyPtrOutput {
	return e.ToAllocationPolicyPtrOutputWithContext(context.Background())
}

func (e AllocationPolicy) ToAllocationPolicyPtrOutputWithContext(ctx context.Context) AllocationPolicyPtrOutput {
	return AllocationPolicy(e).ToAllocationPolicyOutputWithContext(ctx).ToAllocationPolicyPtrOutputWithContext(ctx)
}

func (e AllocationPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AllocationPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AllocationPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AllocationPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AllocationPolicyOutput struct{ *pulumi.OutputState }

func (AllocationPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllocationPolicy)(nil)).Elem()
}

func (o AllocationPolicyOutput) ToAllocationPolicyOutput() AllocationPolicyOutput {
	return o
}

func (o AllocationPolicyOutput) ToAllocationPolicyOutputWithContext(ctx context.Context) AllocationPolicyOutput {
	return o
}

func (o AllocationPolicyOutput) ToAllocationPolicyPtrOutput() AllocationPolicyPtrOutput {
	return o.ToAllocationPolicyPtrOutputWithContext(context.Background())
}

func (o AllocationPolicyOutput) ToAllocationPolicyPtrOutputWithContext(ctx context.Context) AllocationPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AllocationPolicy) *AllocationPolicy {
		return &v
	}).(AllocationPolicyPtrOutput)
}

func (o AllocationPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AllocationPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AllocationPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AllocationPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AllocationPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AllocationPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AllocationPolicyPtrOutput struct{ *pulumi.OutputState }

func (AllocationPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AllocationPolicy)(nil)).Elem()
}

func (o AllocationPolicyPtrOutput) ToAllocationPolicyPtrOutput() AllocationPolicyPtrOutput {
	return o
}

func (o AllocationPolicyPtrOutput) ToAllocationPolicyPtrOutputWithContext(ctx context.Context) AllocationPolicyPtrOutput {
	return o
}

func (o AllocationPolicyPtrOutput) Elem() AllocationPolicyOutput {
	return o.ApplyT(func(v *AllocationPolicy) AllocationPolicy {
		if v != nil {
			return *v
		}
		var ret AllocationPolicy
		return ret
	}).(AllocationPolicyOutput)
}

func (o AllocationPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AllocationPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AllocationPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AllocationPolicyInput interface {
	pulumi.Input

	ToAllocationPolicyOutput() AllocationPolicyOutput
	ToAllocationPolicyOutputWithContext(context.Context) AllocationPolicyOutput
}

var allocationPolicyPtrType = reflect.TypeOf((**AllocationPolicy)(nil)).Elem()

type AllocationPolicyPtrInput interface {
	pulumi.Input

	ToAllocationPolicyPtrOutput() AllocationPolicyPtrOutput
	ToAllocationPolicyPtrOutputWithContext(context.Context) AllocationPolicyPtrOutput
}

type allocationPolicyPtr string

func AllocationPolicyPtr(v string) AllocationPolicyPtrInput {
	return (*allocationPolicyPtr)(&v)
}

func (*allocationPolicyPtr) ElementType() reflect.Type {
	return allocationPolicyPtrType
}

func (in *allocationPolicyPtr) ToAllocationPolicyPtrOutput() AllocationPolicyPtrOutput {
	return pulumi.ToOutput(in).(AllocationPolicyPtrOutput)
}

func (in *allocationPolicyPtr) ToAllocationPolicyPtrOutputWithContext(ctx context.Context) AllocationPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AllocationPolicyPtrOutput)
}

type IotDpsSku string

const (
	IotDpsSkuS1 = IotDpsSku("S1")
)

func (IotDpsSku) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsSku)(nil)).Elem()
}

func (e IotDpsSku) ToIotDpsSkuOutput() IotDpsSkuOutput {
	return pulumi.ToOutput(e).(IotDpsSkuOutput)
}

func (e IotDpsSku) ToIotDpsSkuOutputWithContext(ctx context.Context) IotDpsSkuOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IotDpsSkuOutput)
}

func (e IotDpsSku) ToIotDpsSkuPtrOutput() IotDpsSkuPtrOutput {
	return e.ToIotDpsSkuPtrOutputWithContext(context.Background())
}

func (e IotDpsSku) ToIotDpsSkuPtrOutputWithContext(ctx context.Context) IotDpsSkuPtrOutput {
	return IotDpsSku(e).ToIotDpsSkuOutputWithContext(ctx).ToIotDpsSkuPtrOutputWithContext(ctx)
}

func (e IotDpsSku) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IotDpsSku) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IotDpsSku) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IotDpsSku) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IotDpsSkuOutput struct{ *pulumi.OutputState }

func (IotDpsSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsSku)(nil)).Elem()
}

func (o IotDpsSkuOutput) ToIotDpsSkuOutput() IotDpsSkuOutput {
	return o
}

func (o IotDpsSkuOutput) ToIotDpsSkuOutputWithContext(ctx context.Context) IotDpsSkuOutput {
	return o
}

func (o IotDpsSkuOutput) ToIotDpsSkuPtrOutput() IotDpsSkuPtrOutput {
	return o.ToIotDpsSkuPtrOutputWithContext(context.Background())
}

func (o IotDpsSkuOutput) ToIotDpsSkuPtrOutputWithContext(ctx context.Context) IotDpsSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotDpsSku) *IotDpsSku {
		return &v
	}).(IotDpsSkuPtrOutput)
}

func (o IotDpsSkuOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IotDpsSkuOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IotDpsSku) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IotDpsSkuOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IotDpsSkuOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IotDpsSku) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IotDpsSkuPtrOutput struct{ *pulumi.OutputState }

func (IotDpsSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotDpsSku)(nil)).Elem()
}

func (o IotDpsSkuPtrOutput) ToIotDpsSkuPtrOutput() IotDpsSkuPtrOutput {
	return o
}

func (o IotDpsSkuPtrOutput) ToIotDpsSkuPtrOutputWithContext(ctx context.Context) IotDpsSkuPtrOutput {
	return o
}

func (o IotDpsSkuPtrOutput) Elem() IotDpsSkuOutput {
	return o.ApplyT(func(v *IotDpsSku) IotDpsSku {
		if v != nil {
			return *v
		}
		var ret IotDpsSku
		return ret
	}).(IotDpsSkuOutput)
}

func (o IotDpsSkuPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IotDpsSkuPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IotDpsSku) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IotDpsSkuInput interface {
	pulumi.Input

	ToIotDpsSkuOutput() IotDpsSkuOutput
	ToIotDpsSkuOutputWithContext(context.Context) IotDpsSkuOutput
}

var iotDpsSkuPtrType = reflect.TypeOf((**IotDpsSku)(nil)).Elem()

type IotDpsSkuPtrInput interface {
	pulumi.Input

	ToIotDpsSkuPtrOutput() IotDpsSkuPtrOutput
	ToIotDpsSkuPtrOutputWithContext(context.Context) IotDpsSkuPtrOutput
}

type iotDpsSkuPtr string

func IotDpsSkuPtr(v string) IotDpsSkuPtrInput {
	return (*iotDpsSkuPtr)(&v)
}

func (*iotDpsSkuPtr) ElementType() reflect.Type {
	return iotDpsSkuPtrType
}

func (in *iotDpsSkuPtr) ToIotDpsSkuPtrOutput() IotDpsSkuPtrOutput {
	return pulumi.ToOutput(in).(IotDpsSkuPtrOutput)
}

func (in *iotDpsSkuPtr) ToIotDpsSkuPtrOutputWithContext(ctx context.Context) IotDpsSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IotDpsSkuPtrOutput)
}

type State string

const (
	StateActivating       = State("Activating")
	StateActive           = State("Active")
	StateDeleting         = State("Deleting")
	StateDeleted          = State("Deleted")
	StateActivationFailed = State("ActivationFailed")
	StateDeletionFailed   = State("DeletionFailed")
	StateTransitioning    = State("Transitioning")
	StateSuspending       = State("Suspending")
	StateSuspended        = State("Suspended")
	StateResuming         = State("Resuming")
	StateFailingOver      = State("FailingOver")
	StateFailoverFailed   = State("FailoverFailed")
)

func (State) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (e State) ToStateOutput() StateOutput {
	return pulumi.ToOutput(e).(StateOutput)
}

func (e State) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StateOutput)
}

func (e State) ToStatePtrOutput() StatePtrOutput {
	return e.ToStatePtrOutputWithContext(context.Background())
}

func (e State) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return State(e).ToStateOutputWithContext(ctx).ToStatePtrOutputWithContext(ctx)
}

func (e State) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e State) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e State) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StateOutput struct{ *pulumi.OutputState }

func (StateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*State)(nil)).Elem()
}

func (o StateOutput) ToStateOutput() StateOutput {
	return o
}

func (o StateOutput) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return o
}

func (o StateOutput) ToStatePtrOutput() StatePtrOutput {
	return o.ToStatePtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v State) *State {
		return &v
	}).(StatePtrOutput)
}

func (o StateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e State) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatePtrOutput struct{ *pulumi.OutputState }

func (StatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**State)(nil)).Elem()
}

func (o StatePtrOutput) ToStatePtrOutput() StatePtrOutput {
	return o
}

func (o StatePtrOutput) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return o
}

func (o StatePtrOutput) Elem() StateOutput {
	return o.ApplyT(func(v *State) State {
		if v != nil {
			return *v
		}
		var ret State
		return ret
	}).(StateOutput)
}

func (o StatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *State) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StateInput interface {
	pulumi.Input

	ToStateOutput() StateOutput
	ToStateOutputWithContext(context.Context) StateOutput
}

var statePtrType = reflect.TypeOf((**State)(nil)).Elem()

type StatePtrInput interface {
	pulumi.Input

	ToStatePtrOutput() StatePtrOutput
	ToStatePtrOutputWithContext(context.Context) StatePtrOutput
}

type statePtr string

func StatePtr(v string) StatePtrInput {
	return (*statePtr)(&v)
}

func (*statePtr) ElementType() reflect.Type {
	return statePtrType
}

func (in *statePtr) ToStatePtrOutput() StatePtrOutput {
	return pulumi.ToOutput(in).(StatePtrOutput)
}

func (in *statePtr) ToStatePtrOutputWithContext(ctx context.Context) StatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRightsDescriptionInput)(nil)).Elem(), AccessRightsDescription("ServiceConfig"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRightsDescriptionPtrInput)(nil)).Elem(), AccessRightsDescription("ServiceConfig"))
	pulumi.RegisterInputType(reflect.TypeOf((*AllocationPolicyInput)(nil)).Elem(), AllocationPolicy("Hashed"))
	pulumi.RegisterInputType(reflect.TypeOf((*AllocationPolicyPtrInput)(nil)).Elem(), AllocationPolicy("Hashed"))
	pulumi.RegisterInputType(reflect.TypeOf((*IotDpsSkuInput)(nil)).Elem(), IotDpsSku("S1"))
	pulumi.RegisterInputType(reflect.TypeOf((*IotDpsSkuPtrInput)(nil)).Elem(), IotDpsSku("S1"))
	pulumi.RegisterInputType(reflect.TypeOf((*StateInput)(nil)).Elem(), State("Activating"))
	pulumi.RegisterInputType(reflect.TypeOf((*StatePtrInput)(nil)).Elem(), State("Activating"))
	pulumi.RegisterOutputType(AccessRightsDescriptionOutput{})
	pulumi.RegisterOutputType(AccessRightsDescriptionPtrOutput{})
	pulumi.RegisterOutputType(AllocationPolicyOutput{})
	pulumi.RegisterOutputType(AllocationPolicyPtrOutput{})
	pulumi.RegisterOutputType(IotDpsSkuOutput{})
	pulumi.RegisterOutputType(IotDpsSkuPtrOutput{})
	pulumi.RegisterOutputType(StateOutput{})
	pulumi.RegisterOutputType(StatePtrOutput{})
}
