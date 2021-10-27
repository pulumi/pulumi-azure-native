


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AuthenticationMethod string

const (
	AuthenticationMethodToken = AuthenticationMethod("Token")
)

func (AuthenticationMethod) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationMethod)(nil)).Elem()
}

func (e AuthenticationMethod) ToAuthenticationMethodOutput() AuthenticationMethodOutput {
	return pulumi.ToOutput(e).(AuthenticationMethodOutput)
}

func (e AuthenticationMethod) ToAuthenticationMethodOutputWithContext(ctx context.Context) AuthenticationMethodOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AuthenticationMethodOutput)
}

func (e AuthenticationMethod) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return e.ToAuthenticationMethodPtrOutputWithContext(context.Background())
}

func (e AuthenticationMethod) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return AuthenticationMethod(e).ToAuthenticationMethodOutputWithContext(ctx).ToAuthenticationMethodPtrOutputWithContext(ctx)
}

func (e AuthenticationMethod) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationMethod) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AuthenticationMethod) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AuthenticationMethod) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AuthenticationMethodOutput struct{ *pulumi.OutputState }

func (AuthenticationMethodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthenticationMethod)(nil)).Elem()
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodOutput() AuthenticationMethodOutput {
	return o
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodOutputWithContext(ctx context.Context) AuthenticationMethodOutput {
	return o
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return o.ToAuthenticationMethodPtrOutputWithContext(context.Background())
}

func (o AuthenticationMethodOutput) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthenticationMethod) *AuthenticationMethod {
		return &v
	}).(AuthenticationMethodPtrOutput)
}

func (o AuthenticationMethodOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AuthenticationMethodOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationMethod) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AuthenticationMethodOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationMethodOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AuthenticationMethod) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AuthenticationMethodPtrOutput struct{ *pulumi.OutputState }

func (AuthenticationMethodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthenticationMethod)(nil)).Elem()
}

func (o AuthenticationMethodPtrOutput) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return o
}

func (o AuthenticationMethodPtrOutput) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return o
}

func (o AuthenticationMethodPtrOutput) Elem() AuthenticationMethodOutput {
	return o.ApplyT(func(v *AuthenticationMethod) AuthenticationMethod {
		if v != nil {
			return *v
		}
		var ret AuthenticationMethod
		return ret
	}).(AuthenticationMethodOutput)
}

func (o AuthenticationMethodPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AuthenticationMethodPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AuthenticationMethod) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AuthenticationMethodInput interface {
	pulumi.Input

	ToAuthenticationMethodOutput() AuthenticationMethodOutput
	ToAuthenticationMethodOutputWithContext(context.Context) AuthenticationMethodOutput
}

var authenticationMethodPtrType = reflect.TypeOf((**AuthenticationMethod)(nil)).Elem()

type AuthenticationMethodPtrInput interface {
	pulumi.Input

	ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput
	ToAuthenticationMethodPtrOutputWithContext(context.Context) AuthenticationMethodPtrOutput
}

type authenticationMethodPtr string

func AuthenticationMethodPtr(v string) AuthenticationMethodPtrInput {
	return (*authenticationMethodPtr)(&v)
}

func (*authenticationMethodPtr) ElementType() reflect.Type {
	return authenticationMethodPtrType
}

func (in *authenticationMethodPtr) ToAuthenticationMethodPtrOutput() AuthenticationMethodPtrOutput {
	return pulumi.ToOutput(in).(AuthenticationMethodPtrOutput)
}

func (in *authenticationMethodPtr) ToAuthenticationMethodPtrOutputWithContext(ctx context.Context) AuthenticationMethodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AuthenticationMethodPtrOutput)
}

type ConnectivityStatus string

const (
	ConnectivityStatusConnecting = ConnectivityStatus("Connecting")
	ConnectivityStatusConnected  = ConnectivityStatus("Connected")
	ConnectivityStatusOffline    = ConnectivityStatus("Offline")
	ConnectivityStatusExpired    = ConnectivityStatus("Expired")
)

func (ConnectivityStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityStatus)(nil)).Elem()
}

func (e ConnectivityStatus) ToConnectivityStatusOutput() ConnectivityStatusOutput {
	return pulumi.ToOutput(e).(ConnectivityStatusOutput)
}

func (e ConnectivityStatus) ToConnectivityStatusOutputWithContext(ctx context.Context) ConnectivityStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ConnectivityStatusOutput)
}

func (e ConnectivityStatus) ToConnectivityStatusPtrOutput() ConnectivityStatusPtrOutput {
	return e.ToConnectivityStatusPtrOutputWithContext(context.Background())
}

func (e ConnectivityStatus) ToConnectivityStatusPtrOutputWithContext(ctx context.Context) ConnectivityStatusPtrOutput {
	return ConnectivityStatus(e).ToConnectivityStatusOutputWithContext(ctx).ToConnectivityStatusPtrOutputWithContext(ctx)
}

func (e ConnectivityStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectivityStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ConnectivityStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ConnectivityStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ConnectivityStatusOutput struct{ *pulumi.OutputState }

func (ConnectivityStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityStatus)(nil)).Elem()
}

func (o ConnectivityStatusOutput) ToConnectivityStatusOutput() ConnectivityStatusOutput {
	return o
}

func (o ConnectivityStatusOutput) ToConnectivityStatusOutputWithContext(ctx context.Context) ConnectivityStatusOutput {
	return o
}

func (o ConnectivityStatusOutput) ToConnectivityStatusPtrOutput() ConnectivityStatusPtrOutput {
	return o.ToConnectivityStatusPtrOutputWithContext(context.Background())
}

func (o ConnectivityStatusOutput) ToConnectivityStatusPtrOutputWithContext(ctx context.Context) ConnectivityStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectivityStatus) *ConnectivityStatus {
		return &v
	}).(ConnectivityStatusPtrOutput)
}

func (o ConnectivityStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ConnectivityStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectivityStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ConnectivityStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectivityStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ConnectivityStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ConnectivityStatusPtrOutput struct{ *pulumi.OutputState }

func (ConnectivityStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectivityStatus)(nil)).Elem()
}

func (o ConnectivityStatusPtrOutput) ToConnectivityStatusPtrOutput() ConnectivityStatusPtrOutput {
	return o
}

func (o ConnectivityStatusPtrOutput) ToConnectivityStatusPtrOutputWithContext(ctx context.Context) ConnectivityStatusPtrOutput {
	return o
}

func (o ConnectivityStatusPtrOutput) Elem() ConnectivityStatusOutput {
	return o.ApplyT(func(v *ConnectivityStatus) ConnectivityStatus {
		if v != nil {
			return *v
		}
		var ret ConnectivityStatus
		return ret
	}).(ConnectivityStatusOutput)
}

func (o ConnectivityStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ConnectivityStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ConnectivityStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ConnectivityStatusInput interface {
	pulumi.Input

	ToConnectivityStatusOutput() ConnectivityStatusOutput
	ToConnectivityStatusOutputWithContext(context.Context) ConnectivityStatusOutput
}

var connectivityStatusPtrType = reflect.TypeOf((**ConnectivityStatus)(nil)).Elem()

type ConnectivityStatusPtrInput interface {
	pulumi.Input

	ToConnectivityStatusPtrOutput() ConnectivityStatusPtrOutput
	ToConnectivityStatusPtrOutputWithContext(context.Context) ConnectivityStatusPtrOutput
}

type connectivityStatusPtr string

func ConnectivityStatusPtr(v string) ConnectivityStatusPtrInput {
	return (*connectivityStatusPtr)(&v)
}

func (*connectivityStatusPtr) ElementType() reflect.Type {
	return connectivityStatusPtrType
}

func (in *connectivityStatusPtr) ToConnectivityStatusPtrOutput() ConnectivityStatusPtrOutput {
	return pulumi.ToOutput(in).(ConnectivityStatusPtrOutput)
}

func (in *connectivityStatusPtr) ToConnectivityStatusPtrOutputWithContext(ctx context.Context) ConnectivityStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ConnectivityStatusPtrOutput)
}

type ProvisioningState string

const (
	ProvisioningStateSucceeded    = ProvisioningState("Succeeded")
	ProvisioningStateFailed       = ProvisioningState("Failed")
	ProvisioningStateCanceled     = ProvisioningState("Canceled")
	ProvisioningStateProvisioning = ProvisioningState("Provisioning")
	ProvisioningStateUpdating     = ProvisioningState("Updating")
	ProvisioningStateDeleting     = ProvisioningState("Deleting")
	ProvisioningStateAccepted     = ProvisioningState("Accepted")
)

func (ProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (e ProvisioningState) ToProvisioningStateOutput() ProvisioningStateOutput {
	return pulumi.ToOutput(e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ProvisioningStateOutput)
}

func (e ProvisioningState) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return e.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return ProvisioningState(e).ToProvisioningStateOutputWithContext(ctx).ToProvisioningStatePtrOutputWithContext(ctx)
}

func (e ProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ProvisioningStateOutput struct{ *pulumi.OutputState }

func (ProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStateOutput) ToProvisioningStateOutput() ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStateOutputWithContext(ctx context.Context) ProvisioningStateOutput {
	return o
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o.ToProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProvisioningState) *ProvisioningState {
		return &v
	}).(ProvisioningStatePtrOutput)
}

func (o ProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProvisioningState)(nil)).Elem()
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return o
}

func (o ProvisioningStatePtrOutput) Elem() ProvisioningStateOutput {
	return o.ApplyT(func(v *ProvisioningState) ProvisioningState {
		if v != nil {
			return *v
		}
		var ret ProvisioningState
		return ret
	}).(ProvisioningStateOutput)
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ProvisioningStateInput interface {
	pulumi.Input

	ToProvisioningStateOutput() ProvisioningStateOutput
	ToProvisioningStateOutputWithContext(context.Context) ProvisioningStateOutput
}

var provisioningStatePtrType = reflect.TypeOf((**ProvisioningState)(nil)).Elem()

type ProvisioningStatePtrInput interface {
	pulumi.Input

	ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput
	ToProvisioningStatePtrOutputWithContext(context.Context) ProvisioningStatePtrOutput
}

type provisioningStatePtr string

func ProvisioningStatePtr(v string) ProvisioningStatePtrInput {
	return (*provisioningStatePtr)(&v)
}

func (*provisioningStatePtr) ElementType() reflect.Type {
	return provisioningStatePtrType
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutput() ProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ProvisioningStatePtrOutput)
}

func (in *provisioningStatePtr) ToProvisioningStatePtrOutputWithContext(ctx context.Context) ProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ProvisioningStatePtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationMethodInput)(nil)).Elem(), AuthenticationMethod("Token"))
	pulumi.RegisterInputType(reflect.TypeOf((*AuthenticationMethodPtrInput)(nil)).Elem(), AuthenticationMethod("Token"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectivityStatusInput)(nil)).Elem(), ConnectivityStatus("Connecting"))
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectivityStatusPtrInput)(nil)).Elem(), ConnectivityStatus("Connecting"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisioningStateInput)(nil)).Elem(), ProvisioningState("Succeeded"))
	pulumi.RegisterInputType(reflect.TypeOf((*ProvisioningStatePtrInput)(nil)).Elem(), ProvisioningState("Succeeded"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("None"))
	pulumi.RegisterOutputType(AuthenticationMethodOutput{})
	pulumi.RegisterOutputType(AuthenticationMethodPtrOutput{})
	pulumi.RegisterOutputType(ConnectivityStatusOutput{})
	pulumi.RegisterOutputType(ConnectivityStatusPtrOutput{})
	pulumi.RegisterOutputType(ProvisioningStateOutput{})
	pulumi.RegisterOutputType(ProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
