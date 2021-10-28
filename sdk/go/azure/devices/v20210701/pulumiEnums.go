


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessRights string

const (
	AccessRightsRegistryRead                                             = AccessRights("RegistryRead")
	AccessRightsRegistryWrite                                            = AccessRights("RegistryWrite")
	AccessRightsServiceConnect                                           = AccessRights("ServiceConnect")
	AccessRightsDeviceConnect                                            = AccessRights("DeviceConnect")
	AccessRights_RegistryRead_RegistryWrite                              = AccessRights("RegistryRead, RegistryWrite")
	AccessRights_RegistryRead_ServiceConnect                             = AccessRights("RegistryRead, ServiceConnect")
	AccessRights_RegistryRead_DeviceConnect                              = AccessRights("RegistryRead, DeviceConnect")
	AccessRights_RegistryWrite_ServiceConnect                            = AccessRights("RegistryWrite, ServiceConnect")
	AccessRights_RegistryWrite_DeviceConnect                             = AccessRights("RegistryWrite, DeviceConnect")
	AccessRights_ServiceConnect_DeviceConnect                            = AccessRights("ServiceConnect, DeviceConnect")
	AccessRights_RegistryRead_RegistryWrite_ServiceConnect               = AccessRights("RegistryRead, RegistryWrite, ServiceConnect")
	AccessRights_RegistryRead_RegistryWrite_DeviceConnect                = AccessRights("RegistryRead, RegistryWrite, DeviceConnect")
	AccessRights_RegistryRead_ServiceConnect_DeviceConnect               = AccessRights("RegistryRead, ServiceConnect, DeviceConnect")
	AccessRights_RegistryWrite_ServiceConnect_DeviceConnect              = AccessRights("RegistryWrite, ServiceConnect, DeviceConnect")
	AccessRights_RegistryRead_RegistryWrite_ServiceConnect_DeviceConnect = AccessRights("RegistryRead, RegistryWrite, ServiceConnect, DeviceConnect")
)

func (AccessRights) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRights)(nil)).Elem()
}

func (e AccessRights) ToAccessRightsOutput() AccessRightsOutput {
	return pulumi.ToOutput(e).(AccessRightsOutput)
}

func (e AccessRights) ToAccessRightsOutputWithContext(ctx context.Context) AccessRightsOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessRightsOutput)
}

func (e AccessRights) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return e.ToAccessRightsPtrOutputWithContext(context.Background())
}

func (e AccessRights) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return AccessRights(e).ToAccessRightsOutputWithContext(ctx).ToAccessRightsPtrOutputWithContext(ctx)
}

func (e AccessRights) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRights) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessRights) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessRights) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessRightsOutput struct{ *pulumi.OutputState }

func (AccessRightsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessRights)(nil)).Elem()
}

func (o AccessRightsOutput) ToAccessRightsOutput() AccessRightsOutput {
	return o
}

func (o AccessRightsOutput) ToAccessRightsOutputWithContext(ctx context.Context) AccessRightsOutput {
	return o
}

func (o AccessRightsOutput) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return o.ToAccessRightsPtrOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessRights) *AccessRights {
		return &v
	}).(AccessRightsPtrOutput)
}

func (o AccessRightsOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRights) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessRightsOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRightsOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessRights) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessRightsPtrOutput struct{ *pulumi.OutputState }

func (AccessRightsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessRights)(nil)).Elem()
}

func (o AccessRightsPtrOutput) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return o
}

func (o AccessRightsPtrOutput) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return o
}

func (o AccessRightsPtrOutput) Elem() AccessRightsOutput {
	return o.ApplyT(func(v *AccessRights) AccessRights {
		if v != nil {
			return *v
		}
		var ret AccessRights
		return ret
	}).(AccessRightsOutput)
}

func (o AccessRightsPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessRightsPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessRights) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessRightsInput interface {
	pulumi.Input

	ToAccessRightsOutput() AccessRightsOutput
	ToAccessRightsOutputWithContext(context.Context) AccessRightsOutput
}

var accessRightsPtrType = reflect.TypeOf((**AccessRights)(nil)).Elem()

type AccessRightsPtrInput interface {
	pulumi.Input

	ToAccessRightsPtrOutput() AccessRightsPtrOutput
	ToAccessRightsPtrOutputWithContext(context.Context) AccessRightsPtrOutput
}

type accessRightsPtr string

func AccessRightsPtr(v string) AccessRightsPtrInput {
	return (*accessRightsPtr)(&v)
}

func (*accessRightsPtr) ElementType() reflect.Type {
	return accessRightsPtrType
}

func (in *accessRightsPtr) ToAccessRightsPtrOutput() AccessRightsPtrOutput {
	return pulumi.ToOutput(in).(AccessRightsPtrOutput)
}

func (in *accessRightsPtr) ToAccessRightsPtrOutputWithContext(ctx context.Context) AccessRightsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessRightsPtrOutput)
}

type AuthenticationType string

const (
	AuthenticationTypeKeyBased      = AuthenticationType("keyBased")
	AuthenticationTypeIdentityBased = AuthenticationType("identityBased")
)

func (AuthenticationType) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationType)(nil)).Elem()
}

func (e AuthenticationType) ToAuthenticationTypeOutput() AuthenticationTypeOutput {
	return pulumi.ToOutput(e).(AuthenticationTypeOutput)
}

func (e AuthenticationType) ToAuthenticationTypeOutputWithContext(ctx context.Context) AuthenticationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthenticationTypeOutput)
}

func (e AuthenticationType) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return e.ToAuthenticationTypePtrOutputWithContext(context.Background())
}

func (e AuthenticationType) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return AuthenticationType(e).ToAuthenticationTypeOutputWithContext(ctx).ToAuthenticationTypePtrOutputWithContext(ctx)
}

func (e AuthenticationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthenticationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthenticationTypeOutput struct{ *pulumi.OutputState }

func (AuthenticationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationType)(nil)).Elem()
}

func (o AuthenticationTypeOutput) ToAuthenticationTypeOutput() AuthenticationTypeOutput {
	return o
}

func (o AuthenticationTypeOutput) ToAuthenticationTypeOutputWithContext(ctx context.Context) AuthenticationTypeOutput {
	return o
}

func (o AuthenticationTypeOutput) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return o.ToAuthenticationTypePtrOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationType) *AuthenticationType {
		return &v
	}).(AuthenticationTypePtrOutput)
}

func (o AuthenticationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthenticationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthenticationTypePtrOutput struct{ *pulumi.OutputState }

func (AuthenticationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationType)(nil)).Elem()
}

func (o AuthenticationTypePtrOutput) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return o
}

func (o AuthenticationTypePtrOutput) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return o
}

func (o AuthenticationTypePtrOutput) Elem() AuthenticationTypeOutput {
	return o.ApplyT(func(v *AuthenticationType) AuthenticationType {
		if v != nil {
			return *v
		}
		var ret AuthenticationType
		return ret
	}).(AuthenticationTypeOutput)
}

func (o AuthenticationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthenticationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AuthenticationTypeInput interface {
	pulumi.Input

	ToAuthenticationTypeOutput() AuthenticationTypeOutput
	ToAuthenticationTypeOutputWithContext(context.Context) AuthenticationTypeOutput
}

var authenticationTypePtrType = reflect.TypeOf((**AuthenticationType)(nil)).Elem()

type AuthenticationTypePtrInput interface {
	pulumi.Input

	ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput
	ToAuthenticationTypePtrOutputWithContext(context.Context) AuthenticationTypePtrOutput
}

type authenticationTypePtr string

func AuthenticationTypePtr(v string) AuthenticationTypePtrInput {
	return (*authenticationTypePtr)(&v)
}

func (*authenticationTypePtr) ElementType() reflect.Type {
	return authenticationTypePtrType
}

func (in *authenticationTypePtr) ToAuthenticationTypePtrOutput() AuthenticationTypePtrOutput {
	return pulumi.ToOutput(in).(AuthenticationTypePtrOutput)
}

func (in *authenticationTypePtr) ToAuthenticationTypePtrOutputWithContext(ctx context.Context) AuthenticationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthenticationTypePtrOutput)
}

type Capabilities string

const (
	CapabilitiesNone             = Capabilities("None")
	CapabilitiesDeviceManagement = Capabilities("DeviceManagement")
)

func (Capabilities) ElementType() reflect.Type {
	return reflect.TypeOf((*Capabilities)(nil)).Elem()
}

func (e Capabilities) ToCapabilitiesOutput() CapabilitiesOutput {
	return pulumi.ToOutput(e).(CapabilitiesOutput)
}

func (e Capabilities) ToCapabilitiesOutputWithContext(ctx context.Context) CapabilitiesOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CapabilitiesOutput)
}

func (e Capabilities) ToCapabilitiesPtrOutput() CapabilitiesPtrOutput {
	return e.ToCapabilitiesPtrOutputWithContext(context.Background())
}

func (e Capabilities) ToCapabilitiesPtrOutputWithContext(ctx context.Context) CapabilitiesPtrOutput {
	return Capabilities(e).ToCapabilitiesOutputWithContext(ctx).ToCapabilitiesPtrOutputWithContext(ctx)
}

func (e Capabilities) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Capabilities) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Capabilities) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Capabilities) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CapabilitiesOutput struct{ *pulumi.OutputState }

func (CapabilitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Capabilities)(nil)).Elem()
}

func (o CapabilitiesOutput) ToCapabilitiesOutput() CapabilitiesOutput {
	return o
}

func (o CapabilitiesOutput) ToCapabilitiesOutputWithContext(ctx context.Context) CapabilitiesOutput {
	return o
}

func (o CapabilitiesOutput) ToCapabilitiesPtrOutput() CapabilitiesPtrOutput {
	return o.ToCapabilitiesPtrOutputWithContext(context.Background())
}

func (o CapabilitiesOutput) ToCapabilitiesPtrOutputWithContext(ctx context.Context) CapabilitiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Capabilities) *Capabilities {
		return &v
	}).(CapabilitiesPtrOutput)
}

func (o CapabilitiesOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CapabilitiesOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Capabilities) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CapabilitiesOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CapabilitiesOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Capabilities) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CapabilitiesPtrOutput struct{ *pulumi.OutputState }

func (CapabilitiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Capabilities)(nil)).Elem()
}

func (o CapabilitiesPtrOutput) ToCapabilitiesPtrOutput() CapabilitiesPtrOutput {
	return o
}

func (o CapabilitiesPtrOutput) ToCapabilitiesPtrOutputWithContext(ctx context.Context) CapabilitiesPtrOutput {
	return o
}

func (o CapabilitiesPtrOutput) Elem() CapabilitiesOutput {
	return o.ApplyT(func(v *Capabilities) Capabilities {
		if v != nil {
			return *v
		}
		var ret Capabilities
		return ret
	}).(CapabilitiesOutput)
}

func (o CapabilitiesPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CapabilitiesPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Capabilities) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CapabilitiesInput interface {
	pulumi.Input

	ToCapabilitiesOutput() CapabilitiesOutput
	ToCapabilitiesOutputWithContext(context.Context) CapabilitiesOutput
}

var capabilitiesPtrType = reflect.TypeOf((**Capabilities)(nil)).Elem()

type CapabilitiesPtrInput interface {
	pulumi.Input

	ToCapabilitiesPtrOutput() CapabilitiesPtrOutput
	ToCapabilitiesPtrOutputWithContext(context.Context) CapabilitiesPtrOutput
}

type capabilitiesPtr string

func CapabilitiesPtr(v string) CapabilitiesPtrInput {
	return (*capabilitiesPtr)(&v)
}

func (*capabilitiesPtr) ElementType() reflect.Type {
	return capabilitiesPtrType
}

func (in *capabilitiesPtr) ToCapabilitiesPtrOutput() CapabilitiesPtrOutput {
	return pulumi.ToOutput(in).(CapabilitiesPtrOutput)
}

func (in *capabilitiesPtr) ToCapabilitiesPtrOutputWithContext(ctx context.Context) CapabilitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CapabilitiesPtrOutput)
}

type DefaultAction string

const (
	DefaultActionDeny  = DefaultAction("Deny")
	DefaultActionAllow = DefaultAction("Allow")
)

func (DefaultAction) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAction)(nil)).Elem()
}

func (e DefaultAction) ToDefaultActionOutput() DefaultActionOutput {
	return pulumi.ToOutput(e).(DefaultActionOutput)
}

func (e DefaultAction) ToDefaultActionOutputWithContext(ctx context.Context) DefaultActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DefaultActionOutput)
}

func (e DefaultAction) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return e.ToDefaultActionPtrOutputWithContext(context.Background())
}

func (e DefaultAction) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return DefaultAction(e).ToDefaultActionOutputWithContext(ctx).ToDefaultActionPtrOutputWithContext(ctx)
}

func (e DefaultAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefaultAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DefaultActionOutput struct{ *pulumi.OutputState }

func (DefaultActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultAction)(nil)).Elem()
}

func (o DefaultActionOutput) ToDefaultActionOutput() DefaultActionOutput {
	return o
}

func (o DefaultActionOutput) ToDefaultActionOutputWithContext(ctx context.Context) DefaultActionOutput {
	return o
}

func (o DefaultActionOutput) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return o.ToDefaultActionPtrOutputWithContext(context.Background())
}

func (o DefaultActionOutput) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultAction) *DefaultAction {
		return &v
	}).(DefaultActionPtrOutput)
}

func (o DefaultActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DefaultActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DefaultActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DefaultActionPtrOutput struct{ *pulumi.OutputState }

func (DefaultActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultAction)(nil)).Elem()
}

func (o DefaultActionPtrOutput) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return o
}

func (o DefaultActionPtrOutput) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return o
}

func (o DefaultActionPtrOutput) Elem() DefaultActionOutput {
	return o.ApplyT(func(v *DefaultAction) DefaultAction {
		if v != nil {
			return *v
		}
		var ret DefaultAction
		return ret
	}).(DefaultActionOutput)
}

func (o DefaultActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DefaultAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DefaultActionInput interface {
	pulumi.Input

	ToDefaultActionOutput() DefaultActionOutput
	ToDefaultActionOutputWithContext(context.Context) DefaultActionOutput
}

var defaultActionPtrType = reflect.TypeOf((**DefaultAction)(nil)).Elem()

type DefaultActionPtrInput interface {
	pulumi.Input

	ToDefaultActionPtrOutput() DefaultActionPtrOutput
	ToDefaultActionPtrOutputWithContext(context.Context) DefaultActionPtrOutput
}

type defaultActionPtr string

func DefaultActionPtr(v string) DefaultActionPtrInput {
	return (*defaultActionPtr)(&v)
}

func (*defaultActionPtr) ElementType() reflect.Type {
	return defaultActionPtrType
}

func (in *defaultActionPtr) ToDefaultActionPtrOutput() DefaultActionPtrOutput {
	return pulumi.ToOutput(in).(DefaultActionPtrOutput)
}

func (in *defaultActionPtr) ToDefaultActionPtrOutputWithContext(ctx context.Context) DefaultActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DefaultActionPtrOutput)
}

type IotHubSku string

const (
	IotHubSkuF1 = IotHubSku("F1")
	IotHubSkuS1 = IotHubSku("S1")
	IotHubSkuS2 = IotHubSku("S2")
	IotHubSkuS3 = IotHubSku("S3")
	IotHubSkuB1 = IotHubSku("B1")
	IotHubSkuB2 = IotHubSku("B2")
	IotHubSkuB3 = IotHubSku("B3")
)

func (IotHubSku) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSku)(nil)).Elem()
}

func (e IotHubSku) ToIotHubSkuOutput() IotHubSkuOutput {
	return pulumi.ToOutput(e).(IotHubSkuOutput)
}

func (e IotHubSku) ToIotHubSkuOutputWithContext(ctx context.Context) IotHubSkuOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IotHubSkuOutput)
}

func (e IotHubSku) ToIotHubSkuPtrOutput() IotHubSkuPtrOutput {
	return e.ToIotHubSkuPtrOutputWithContext(context.Background())
}

func (e IotHubSku) ToIotHubSkuPtrOutputWithContext(ctx context.Context) IotHubSkuPtrOutput {
	return IotHubSku(e).ToIotHubSkuOutputWithContext(ctx).ToIotHubSkuPtrOutputWithContext(ctx)
}

func (e IotHubSku) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IotHubSku) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IotHubSku) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IotHubSku) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IotHubSkuOutput struct{ *pulumi.OutputState }

func (IotHubSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSku)(nil)).Elem()
}

func (o IotHubSkuOutput) ToIotHubSkuOutput() IotHubSkuOutput {
	return o
}

func (o IotHubSkuOutput) ToIotHubSkuOutputWithContext(ctx context.Context) IotHubSkuOutput {
	return o
}

func (o IotHubSkuOutput) ToIotHubSkuPtrOutput() IotHubSkuPtrOutput {
	return o.ToIotHubSkuPtrOutputWithContext(context.Background())
}

func (o IotHubSkuOutput) ToIotHubSkuPtrOutputWithContext(ctx context.Context) IotHubSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotHubSku) *IotHubSku {
		return &v
	}).(IotHubSkuPtrOutput)
}

func (o IotHubSkuOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IotHubSkuOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IotHubSku) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IotHubSkuOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IotHubSkuOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IotHubSku) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IotHubSkuPtrOutput struct{ *pulumi.OutputState }

func (IotHubSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubSku)(nil)).Elem()
}

func (o IotHubSkuPtrOutput) ToIotHubSkuPtrOutput() IotHubSkuPtrOutput {
	return o
}

func (o IotHubSkuPtrOutput) ToIotHubSkuPtrOutputWithContext(ctx context.Context) IotHubSkuPtrOutput {
	return o
}

func (o IotHubSkuPtrOutput) Elem() IotHubSkuOutput {
	return o.ApplyT(func(v *IotHubSku) IotHubSku {
		if v != nil {
			return *v
		}
		var ret IotHubSku
		return ret
	}).(IotHubSkuOutput)
}

func (o IotHubSkuPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IotHubSkuPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IotHubSku) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IotHubSkuInput interface {
	pulumi.Input

	ToIotHubSkuOutput() IotHubSkuOutput
	ToIotHubSkuOutputWithContext(context.Context) IotHubSkuOutput
}

var iotHubSkuPtrType = reflect.TypeOf((**IotHubSku)(nil)).Elem()

type IotHubSkuPtrInput interface {
	pulumi.Input

	ToIotHubSkuPtrOutput() IotHubSkuPtrOutput
	ToIotHubSkuPtrOutputWithContext(context.Context) IotHubSkuPtrOutput
}

type iotHubSkuPtr string

func IotHubSkuPtr(v string) IotHubSkuPtrInput {
	return (*iotHubSkuPtr)(&v)
}

func (*iotHubSkuPtr) ElementType() reflect.Type {
	return iotHubSkuPtrType
}

func (in *iotHubSkuPtr) ToIotHubSkuPtrOutput() IotHubSkuPtrOutput {
	return pulumi.ToOutput(in).(IotHubSkuPtrOutput)
}

func (in *iotHubSkuPtr) ToIotHubSkuPtrOutputWithContext(ctx context.Context) IotHubSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IotHubSkuPtrOutput)
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

type NetworkRuleIPAction string

const (
	NetworkRuleIPActionAllow = NetworkRuleIPAction("Allow")
)

func (NetworkRuleIPAction) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleIPAction)(nil)).Elem()
}

func (e NetworkRuleIPAction) ToNetworkRuleIPActionOutput() NetworkRuleIPActionOutput {
	return pulumi.ToOutput(e).(NetworkRuleIPActionOutput)
}

func (e NetworkRuleIPAction) ToNetworkRuleIPActionOutputWithContext(ctx context.Context) NetworkRuleIPActionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NetworkRuleIPActionOutput)
}

func (e NetworkRuleIPAction) ToNetworkRuleIPActionPtrOutput() NetworkRuleIPActionPtrOutput {
	return e.ToNetworkRuleIPActionPtrOutputWithContext(context.Background())
}

func (e NetworkRuleIPAction) ToNetworkRuleIPActionPtrOutputWithContext(ctx context.Context) NetworkRuleIPActionPtrOutput {
	return NetworkRuleIPAction(e).ToNetworkRuleIPActionOutputWithContext(ctx).ToNetworkRuleIPActionPtrOutputWithContext(ctx)
}

func (e NetworkRuleIPAction) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkRuleIPAction) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NetworkRuleIPAction) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NetworkRuleIPAction) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NetworkRuleIPActionOutput struct{ *pulumi.OutputState }

func (NetworkRuleIPActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleIPAction)(nil)).Elem()
}

func (o NetworkRuleIPActionOutput) ToNetworkRuleIPActionOutput() NetworkRuleIPActionOutput {
	return o
}

func (o NetworkRuleIPActionOutput) ToNetworkRuleIPActionOutputWithContext(ctx context.Context) NetworkRuleIPActionOutput {
	return o
}

func (o NetworkRuleIPActionOutput) ToNetworkRuleIPActionPtrOutput() NetworkRuleIPActionPtrOutput {
	return o.ToNetworkRuleIPActionPtrOutputWithContext(context.Background())
}

func (o NetworkRuleIPActionOutput) ToNetworkRuleIPActionPtrOutputWithContext(ctx context.Context) NetworkRuleIPActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleIPAction) *NetworkRuleIPAction {
		return &v
	}).(NetworkRuleIPActionPtrOutput)
}

func (o NetworkRuleIPActionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NetworkRuleIPActionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkRuleIPAction) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NetworkRuleIPActionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkRuleIPActionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NetworkRuleIPAction) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NetworkRuleIPActionPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleIPActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleIPAction)(nil)).Elem()
}

func (o NetworkRuleIPActionPtrOutput) ToNetworkRuleIPActionPtrOutput() NetworkRuleIPActionPtrOutput {
	return o
}

func (o NetworkRuleIPActionPtrOutput) ToNetworkRuleIPActionPtrOutputWithContext(ctx context.Context) NetworkRuleIPActionPtrOutput {
	return o
}

func (o NetworkRuleIPActionPtrOutput) Elem() NetworkRuleIPActionOutput {
	return o.ApplyT(func(v *NetworkRuleIPAction) NetworkRuleIPAction {
		if v != nil {
			return *v
		}
		var ret NetworkRuleIPAction
		return ret
	}).(NetworkRuleIPActionOutput)
}

func (o NetworkRuleIPActionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NetworkRuleIPActionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NetworkRuleIPAction) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NetworkRuleIPActionInput interface {
	pulumi.Input

	ToNetworkRuleIPActionOutput() NetworkRuleIPActionOutput
	ToNetworkRuleIPActionOutputWithContext(context.Context) NetworkRuleIPActionOutput
}

var networkRuleIPActionPtrType = reflect.TypeOf((**NetworkRuleIPAction)(nil)).Elem()

type NetworkRuleIPActionPtrInput interface {
	pulumi.Input

	ToNetworkRuleIPActionPtrOutput() NetworkRuleIPActionPtrOutput
	ToNetworkRuleIPActionPtrOutputWithContext(context.Context) NetworkRuleIPActionPtrOutput
}

type networkRuleIPActionPtr string

func NetworkRuleIPActionPtr(v string) NetworkRuleIPActionPtrInput {
	return (*networkRuleIPActionPtr)(&v)
}

func (*networkRuleIPActionPtr) ElementType() reflect.Type {
	return networkRuleIPActionPtrType
}

func (in *networkRuleIPActionPtr) ToNetworkRuleIPActionPtrOutput() NetworkRuleIPActionPtrOutput {
	return pulumi.ToOutput(in).(NetworkRuleIPActionPtrOutput)
}

func (in *networkRuleIPActionPtr) ToNetworkRuleIPActionPtrOutputWithContext(ctx context.Context) NetworkRuleIPActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NetworkRuleIPActionPtrOutput)
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

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned               = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeUserAssigned                 = ResourceIdentityType("UserAssigned")
	ResourceIdentityType_SystemAssigned_UserAssigned = ResourceIdentityType("SystemAssigned, UserAssigned")
	ResourceIdentityTypeNone                         = ResourceIdentityType("None")
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

type RoutingSource string

const (
	RoutingSourceInvalid                     = RoutingSource("Invalid")
	RoutingSourceDeviceMessages              = RoutingSource("DeviceMessages")
	RoutingSourceTwinChangeEvents            = RoutingSource("TwinChangeEvents")
	RoutingSourceDeviceLifecycleEvents       = RoutingSource("DeviceLifecycleEvents")
	RoutingSourceDeviceJobLifecycleEvents    = RoutingSource("DeviceJobLifecycleEvents")
	RoutingSourceDeviceConnectionStateEvents = RoutingSource("DeviceConnectionStateEvents")
)

func (RoutingSource) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingSource)(nil)).Elem()
}

func (e RoutingSource) ToRoutingSourceOutput() RoutingSourceOutput {
	return pulumi.ToOutput(e).(RoutingSourceOutput)
}

func (e RoutingSource) ToRoutingSourceOutputWithContext(ctx context.Context) RoutingSourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RoutingSourceOutput)
}

func (e RoutingSource) ToRoutingSourcePtrOutput() RoutingSourcePtrOutput {
	return e.ToRoutingSourcePtrOutputWithContext(context.Background())
}

func (e RoutingSource) ToRoutingSourcePtrOutputWithContext(ctx context.Context) RoutingSourcePtrOutput {
	return RoutingSource(e).ToRoutingSourceOutputWithContext(ctx).ToRoutingSourcePtrOutputWithContext(ctx)
}

func (e RoutingSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoutingSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RoutingSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RoutingSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RoutingSourceOutput struct{ *pulumi.OutputState }

func (RoutingSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingSource)(nil)).Elem()
}

func (o RoutingSourceOutput) ToRoutingSourceOutput() RoutingSourceOutput {
	return o
}

func (o RoutingSourceOutput) ToRoutingSourceOutputWithContext(ctx context.Context) RoutingSourceOutput {
	return o
}

func (o RoutingSourceOutput) ToRoutingSourcePtrOutput() RoutingSourcePtrOutput {
	return o.ToRoutingSourcePtrOutputWithContext(context.Background())
}

func (o RoutingSourceOutput) ToRoutingSourcePtrOutputWithContext(ctx context.Context) RoutingSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingSource) *RoutingSource {
		return &v
	}).(RoutingSourcePtrOutput)
}

func (o RoutingSourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RoutingSourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoutingSource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RoutingSourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoutingSourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RoutingSource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RoutingSourcePtrOutput struct{ *pulumi.OutputState }

func (RoutingSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingSource)(nil)).Elem()
}

func (o RoutingSourcePtrOutput) ToRoutingSourcePtrOutput() RoutingSourcePtrOutput {
	return o
}

func (o RoutingSourcePtrOutput) ToRoutingSourcePtrOutputWithContext(ctx context.Context) RoutingSourcePtrOutput {
	return o
}

func (o RoutingSourcePtrOutput) Elem() RoutingSourceOutput {
	return o.ApplyT(func(v *RoutingSource) RoutingSource {
		if v != nil {
			return *v
		}
		var ret RoutingSource
		return ret
	}).(RoutingSourceOutput)
}

func (o RoutingSourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RoutingSourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RoutingSource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RoutingSourceInput interface {
	pulumi.Input

	ToRoutingSourceOutput() RoutingSourceOutput
	ToRoutingSourceOutputWithContext(context.Context) RoutingSourceOutput
}

var routingSourcePtrType = reflect.TypeOf((**RoutingSource)(nil)).Elem()

type RoutingSourcePtrInput interface {
	pulumi.Input

	ToRoutingSourcePtrOutput() RoutingSourcePtrOutput
	ToRoutingSourcePtrOutputWithContext(context.Context) RoutingSourcePtrOutput
}

type routingSourcePtr string

func RoutingSourcePtr(v string) RoutingSourcePtrInput {
	return (*routingSourcePtr)(&v)
}

func (*routingSourcePtr) ElementType() reflect.Type {
	return routingSourcePtrType
}

func (in *routingSourcePtr) ToRoutingSourcePtrOutput() RoutingSourcePtrOutput {
	return pulumi.ToOutput(in).(RoutingSourcePtrOutput)
}

func (in *routingSourcePtr) ToRoutingSourcePtrOutputWithContext(ctx context.Context) RoutingSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RoutingSourcePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRightsInput)(nil)).Elem(), AccessRights("RegistryRead"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessRightsPtrInput)(nil)).Elem(), AccessRights("RegistryRead"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationTypeInput)(nil)).Elem(), AuthenticationType("keyBased"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationTypePtrInput)(nil)).Elem(), AuthenticationType("keyBased"))
	pulumi.RegisterInputType(reflect.TypeOf((*CapabilitiesInput)(nil)).Elem(), Capabilities("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*CapabilitiesPtrInput)(nil)).Elem(), Capabilities("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultActionInput)(nil)).Elem(), DefaultAction("Deny"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultActionPtrInput)(nil)).Elem(), DefaultAction("Deny"))
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubSkuInput)(nil)).Elem(), IotHubSku("F1"))
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubSkuPtrInput)(nil)).Elem(), IotHubSku("F1"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterActionTypeInput)(nil)).Elem(), IpFilterActionType("Accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterActionTypePtrInput)(nil)).Elem(), IpFilterActionType("Accept"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRuleIPActionInput)(nil)).Elem(), NetworkRuleIPAction("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRuleIPActionPtrInput)(nil)).Elem(), NetworkRuleIPAction("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatusInput)(nil)).Elem(), PrivateLinkServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatusPtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicNetworkAccessInput)(nil)).Elem(), PublicNetworkAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicNetworkAccessPtrInput)(nil)).Elem(), PublicNetworkAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingSourceInput)(nil)).Elem(), RoutingSource("Invalid"))
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingSourcePtrInput)(nil)).Elem(), RoutingSource("Invalid"))
	pulumi.RegisterOutputType(AccessRightsOutput{})
	pulumi.RegisterOutputType(AccessRightsPtrOutput{})
	pulumi.RegisterOutputType(AuthenticationTypeOutput{})
	pulumi.RegisterOutputType(AuthenticationTypePtrOutput{})
	pulumi.RegisterOutputType(CapabilitiesOutput{})
	pulumi.RegisterOutputType(CapabilitiesPtrOutput{})
	pulumi.RegisterOutputType(DefaultActionOutput{})
	pulumi.RegisterOutputType(DefaultActionPtrOutput{})
	pulumi.RegisterOutputType(IotHubSkuOutput{})
	pulumi.RegisterOutputType(IotHubSkuPtrOutput{})
	pulumi.RegisterOutputType(IpFilterActionTypeOutput{})
	pulumi.RegisterOutputType(IpFilterActionTypePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleIPActionOutput{})
	pulumi.RegisterOutputType(NetworkRuleIPActionPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(RoutingSourceOutput{})
	pulumi.RegisterOutputType(RoutingSourcePtrOutput{})
}
