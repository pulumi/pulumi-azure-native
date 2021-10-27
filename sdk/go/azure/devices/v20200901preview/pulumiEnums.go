


package v20200901preview

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

type IpFilterActionType string

const (
	IpFilterActionTypeAccept = IpFilterActionType("Accept")
	IpFilterActionTypeReject = IpFilterActionType("Reject")
)

func (IpFilterActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterActionType)(nil)).Elem()
}

func (e IpFilterActionType) ToIpFilterActionTypeOutput() IpFilterActionTypeOutput {
	return pulumi.ToOutput(e).(IpFilterActionTypeOutput)
}

func (e IpFilterActionType) ToIpFilterActionTypeOutputWithContext(ctx context.Context) IpFilterActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpFilterActionTypeOutput)
}

func (e IpFilterActionType) ToIpFilterActionTypePtrOutput() IpFilterActionTypePtrOutput {
	return e.ToIpFilterActionTypePtrOutputWithContext(context.Background())
}

func (e IpFilterActionType) ToIpFilterActionTypePtrOutputWithContext(ctx context.Context) IpFilterActionTypePtrOutput {
	return IpFilterActionType(e).ToIpFilterActionTypeOutputWithContext(ctx).ToIpFilterActionTypePtrOutputWithContext(ctx)
}

func (e IpFilterActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpFilterActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpFilterActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpFilterActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpFilterActionTypeOutput struct{ *pulumi.OutputState }

func (IpFilterActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterActionType)(nil)).Elem()
}

func (o IpFilterActionTypeOutput) ToIpFilterActionTypeOutput() IpFilterActionTypeOutput {
	return o
}

func (o IpFilterActionTypeOutput) ToIpFilterActionTypeOutputWithContext(ctx context.Context) IpFilterActionTypeOutput {
	return o
}

func (o IpFilterActionTypeOutput) ToIpFilterActionTypePtrOutput() IpFilterActionTypePtrOutput {
	return o.ToIpFilterActionTypePtrOutputWithContext(context.Background())
}

func (o IpFilterActionTypeOutput) ToIpFilterActionTypePtrOutputWithContext(ctx context.Context) IpFilterActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpFilterActionType) *IpFilterActionType {
		return &v
	}).(IpFilterActionTypePtrOutput)
}

func (o IpFilterActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpFilterActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpFilterActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpFilterActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpFilterActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpFilterActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpFilterActionTypePtrOutput struct{ *pulumi.OutputState }

func (IpFilterActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpFilterActionType)(nil)).Elem()
}

func (o IpFilterActionTypePtrOutput) ToIpFilterActionTypePtrOutput() IpFilterActionTypePtrOutput {
	return o
}

func (o IpFilterActionTypePtrOutput) ToIpFilterActionTypePtrOutputWithContext(ctx context.Context) IpFilterActionTypePtrOutput {
	return o
}

func (o IpFilterActionTypePtrOutput) Elem() IpFilterActionTypeOutput {
	return o.ApplyT(func(v *IpFilterActionType) IpFilterActionType {
		if v != nil {
			return *v
		}
		var ret IpFilterActionType
		return ret
	}).(IpFilterActionTypeOutput)
}

func (o IpFilterActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpFilterActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpFilterActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IpFilterActionTypeInput interface {
	pulumi.Input

	ToIpFilterActionTypeOutput() IpFilterActionTypeOutput
	ToIpFilterActionTypeOutputWithContext(context.Context) IpFilterActionTypeOutput
}

var ipFilterActionTypePtrType = reflect.TypeOf((**IpFilterActionType)(nil)).Elem()

type IpFilterActionTypePtrInput interface {
	pulumi.Input

	ToIpFilterActionTypePtrOutput() IpFilterActionTypePtrOutput
	ToIpFilterActionTypePtrOutputWithContext(context.Context) IpFilterActionTypePtrOutput
}

type ipFilterActionTypePtr string

func IpFilterActionTypePtr(v string) IpFilterActionTypePtrInput {
	return (*ipFilterActionTypePtr)(&v)
}

func (*ipFilterActionTypePtr) ElementType() reflect.Type {
	return ipFilterActionTypePtrType
}

func (in *ipFilterActionTypePtr) ToIpFilterActionTypePtrOutput() IpFilterActionTypePtrOutput {
	return pulumi.ToOutput(in).(IpFilterActionTypePtrOutput)
}

func (in *ipFilterActionTypePtr) ToIpFilterActionTypePtrOutputWithContext(ctx context.Context) IpFilterActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpFilterActionTypePtrOutput)
}

type IpFilterTargetType string

const (
	IpFilterTargetTypeAll        = IpFilterTargetType("all")
	IpFilterTargetTypeServiceApi = IpFilterTargetType("serviceApi")
	IpFilterTargetTypeDeviceApi  = IpFilterTargetType("deviceApi")
)

func (IpFilterTargetType) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterTargetType)(nil)).Elem()
}

func (e IpFilterTargetType) ToIpFilterTargetTypeOutput() IpFilterTargetTypeOutput {
	return pulumi.ToOutput(e).(IpFilterTargetTypeOutput)
}

func (e IpFilterTargetType) ToIpFilterTargetTypeOutputWithContext(ctx context.Context) IpFilterTargetTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpFilterTargetTypeOutput)
}

func (e IpFilterTargetType) ToIpFilterTargetTypePtrOutput() IpFilterTargetTypePtrOutput {
	return e.ToIpFilterTargetTypePtrOutputWithContext(context.Background())
}

func (e IpFilterTargetType) ToIpFilterTargetTypePtrOutputWithContext(ctx context.Context) IpFilterTargetTypePtrOutput {
	return IpFilterTargetType(e).ToIpFilterTargetTypeOutputWithContext(ctx).ToIpFilterTargetTypePtrOutputWithContext(ctx)
}

func (e IpFilterTargetType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpFilterTargetType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpFilterTargetType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpFilterTargetType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpFilterTargetTypeOutput struct{ *pulumi.OutputState }

func (IpFilterTargetTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterTargetType)(nil)).Elem()
}

func (o IpFilterTargetTypeOutput) ToIpFilterTargetTypeOutput() IpFilterTargetTypeOutput {
	return o
}

func (o IpFilterTargetTypeOutput) ToIpFilterTargetTypeOutputWithContext(ctx context.Context) IpFilterTargetTypeOutput {
	return o
}

func (o IpFilterTargetTypeOutput) ToIpFilterTargetTypePtrOutput() IpFilterTargetTypePtrOutput {
	return o.ToIpFilterTargetTypePtrOutputWithContext(context.Background())
}

func (o IpFilterTargetTypeOutput) ToIpFilterTargetTypePtrOutputWithContext(ctx context.Context) IpFilterTargetTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpFilterTargetType) *IpFilterTargetType {
		return &v
	}).(IpFilterTargetTypePtrOutput)
}

func (o IpFilterTargetTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpFilterTargetTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpFilterTargetType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpFilterTargetTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpFilterTargetTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpFilterTargetType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpFilterTargetTypePtrOutput struct{ *pulumi.OutputState }

func (IpFilterTargetTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpFilterTargetType)(nil)).Elem()
}

func (o IpFilterTargetTypePtrOutput) ToIpFilterTargetTypePtrOutput() IpFilterTargetTypePtrOutput {
	return o
}

func (o IpFilterTargetTypePtrOutput) ToIpFilterTargetTypePtrOutputWithContext(ctx context.Context) IpFilterTargetTypePtrOutput {
	return o
}

func (o IpFilterTargetTypePtrOutput) Elem() IpFilterTargetTypeOutput {
	return o.ApplyT(func(v *IpFilterTargetType) IpFilterTargetType {
		if v != nil {
			return *v
		}
		var ret IpFilterTargetType
		return ret
	}).(IpFilterTargetTypeOutput)
}

func (o IpFilterTargetTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpFilterTargetTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpFilterTargetType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IpFilterTargetTypeInput interface {
	pulumi.Input

	ToIpFilterTargetTypeOutput() IpFilterTargetTypeOutput
	ToIpFilterTargetTypeOutputWithContext(context.Context) IpFilterTargetTypeOutput
}

var ipFilterTargetTypePtrType = reflect.TypeOf((**IpFilterTargetType)(nil)).Elem()

type IpFilterTargetTypePtrInput interface {
	pulumi.Input

	ToIpFilterTargetTypePtrOutput() IpFilterTargetTypePtrOutput
	ToIpFilterTargetTypePtrOutputWithContext(context.Context) IpFilterTargetTypePtrOutput
}

type ipFilterTargetTypePtr string

func IpFilterTargetTypePtr(v string) IpFilterTargetTypePtrInput {
	return (*ipFilterTargetTypePtr)(&v)
}

func (*ipFilterTargetTypePtr) ElementType() reflect.Type {
	return ipFilterTargetTypePtrType
}

func (in *ipFilterTargetTypePtr) ToIpFilterTargetTypePtrOutput() IpFilterTargetTypePtrOutput {
	return pulumi.ToOutput(in).(IpFilterTargetTypePtrOutput)
}

func (in *ipFilterTargetTypePtr) ToIpFilterTargetTypePtrOutputWithContext(ctx context.Context) IpFilterTargetTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpFilterTargetTypePtrOutput)
}

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusPending      = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusApproved     = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected     = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusDisconnected = PrivateLinkServiceConnectionStatus("Disconnected")
)

func (PrivateLinkServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateLinkServiceConnectionStatusOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateLinkServiceConnectionStatusOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return e.ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return PrivateLinkServiceConnectionStatus(e).ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx).ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateLinkServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateLinkServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStatus) *PrivateLinkServiceConnectionStatus {
		return &v
	}).(PrivateLinkServiceConnectionStatusPtrOutput)
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) Elem() PrivateLinkServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatus) PrivateLinkServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatus
		return ret
	}).(PrivateLinkServiceConnectionStatusOutput)
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateLinkServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateLinkServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput
	ToPrivateLinkServiceConnectionStatusOutputWithContext(context.Context) PrivateLinkServiceConnectionStatusOutput
}

var privateLinkServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateLinkServiceConnectionStatus)(nil)).Elem()

type PrivateLinkServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput
	ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatusPtrOutput
}

type privateLinkServiceConnectionStatusPtr string

func PrivateLinkServiceConnectionStatusPtr(v string) PrivateLinkServiceConnectionStatusPtrInput {
	return (*privateLinkServiceConnectionStatusPtr)(&v)
}

func (*privateLinkServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateLinkServiceConnectionStatusPtrType
}

func (in *privateLinkServiceConnectionStatusPtr) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateLinkServiceConnectionStatusPtrOutput)
}

func (in *privateLinkServiceConnectionStatusPtr) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateLinkServiceConnectionStatusPtrOutput)
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

func (PublicNetworkAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return e.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return PublicNetworkAccess(e).ToPublicNetworkAccessOutputWithContext(ctx).ToPublicNetworkAccessPtrOutputWithContext(ctx)
}

func (e PublicNetworkAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccess) *PublicNetworkAccess {
		return &v
	}).(PublicNetworkAccessPtrOutput)
}

func (o PublicNetworkAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessPtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) Elem() PublicNetworkAccessOutput {
	return o.ApplyT(func(v *PublicNetworkAccess) PublicNetworkAccess {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccess
		return ret
	}).(PublicNetworkAccessOutput)
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicNetworkAccessInput interface {
	pulumi.Input

	ToPublicNetworkAccessOutput() PublicNetworkAccessOutput
	ToPublicNetworkAccessOutputWithContext(context.Context) PublicNetworkAccessOutput
}

var publicNetworkAccessPtrType = reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()

type PublicNetworkAccessPtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput
	ToPublicNetworkAccessPtrOutputWithContext(context.Context) PublicNetworkAccessPtrOutput
}

type publicNetworkAccessPtr string

func PublicNetworkAccessPtr(v string) PublicNetworkAccessPtrInput {
	return (*publicNetworkAccessPtr)(&v)
}

func (*publicNetworkAccessPtr) ElementType() reflect.Type {
	return publicNetworkAccessPtrType
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessPtrOutput)
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessPtrOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterActionTypeInput)(nil)).Elem(), IpFilterActionType("Accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterActionTypePtrInput)(nil)).Elem(), IpFilterActionType("Accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterTargetTypeInput)(nil)).Elem(), IpFilterTargetType("all"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterTargetTypePtrInput)(nil)).Elem(), IpFilterTargetType("all"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatusInput)(nil)).Elem(), PrivateLinkServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatusPtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicNetworkAccessInput)(nil)).Elem(), PublicNetworkAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicNetworkAccessPtrInput)(nil)).Elem(), PublicNetworkAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*StateInput)(nil)).Elem(), State("Activating"))
	pulumi.RegisterInputType(reflect.TypeOf((*StatePtrInput)(nil)).Elem(), State("Activating"))
	pulumi.RegisterOutputType(AccessRightsDescriptionOutput{})
	pulumi.RegisterOutputType(AccessRightsDescriptionPtrOutput{})
	pulumi.RegisterOutputType(AllocationPolicyOutput{})
	pulumi.RegisterOutputType(AllocationPolicyPtrOutput{})
	pulumi.RegisterOutputType(IotDpsSkuOutput{})
	pulumi.RegisterOutputType(IotDpsSkuPtrOutput{})
	pulumi.RegisterOutputType(IpFilterActionTypeOutput{})
	pulumi.RegisterOutputType(IpFilterActionTypePtrOutput{})
	pulumi.RegisterOutputType(IpFilterTargetTypeOutput{})
	pulumi.RegisterOutputType(IpFilterTargetTypePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(StateOutput{})
	pulumi.RegisterOutputType(StatePtrOutput{})
}
