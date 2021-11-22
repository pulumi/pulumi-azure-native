


package digitaltwins

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConnectionPropertiesPrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     string  `pulumi:"description"`
	Status          string  `pulumi:"status"`
}





type ConnectionPropertiesPrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToConnectionPropertiesPrivateLinkServiceConnectionStateOutput() ConnectionPropertiesPrivateLinkServiceConnectionStateOutput
	ToConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStateOutput
}

type ConnectionPropertiesPrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput    `pulumi:"description"`
	Status          pulumi.StringInput    `pulumi:"status"`
}

func (ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesPrivateLinkServiceConnectionStateOutput() ConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return i.ToConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesPrivateLinkServiceConnectionStateOutput)
}

func (i ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return i.ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesPrivateLinkServiceConnectionStateOutput).ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type ConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput
	ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput
}

type connectionPropertiesPrivateLinkServiceConnectionStatePtrType ConnectionPropertiesPrivateLinkServiceConnectionStateArgs

func ConnectionPropertiesPrivateLinkServiceConnectionStatePtr(v *ConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput {
	return (*connectionPropertiesPrivateLinkServiceConnectionStatePtrType)(v)
}

func (*connectionPropertiesPrivateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *connectionPropertiesPrivateLinkServiceConnectionStatePtrType) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return i.ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *connectionPropertiesPrivateLinkServiceConnectionStatePtrType) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

type ConnectionPropertiesPrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStateOutput() ConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionPropertiesPrivateLinkServiceConnectionState) *ConnectionPropertiesPrivateLinkServiceConnectionState {
		return &v
	}).(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionPropertiesPrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesPrivateLinkServiceConnectionState) string { return v.Description }).(pulumi.StringOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesPrivateLinkServiceConnectionState) string { return v.Status }).(pulumi.StringOutput)
}

type ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ToConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Elem() ConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *ConnectionPropertiesPrivateLinkServiceConnectionState) ConnectionPropertiesPrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret ConnectionPropertiesPrivateLinkServiceConnectionState
		return ret
	}).(ConnectionPropertiesPrivateLinkServiceConnectionStateOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type ConnectionPropertiesResponsePrivateEndpoint struct {
	Id string `pulumi:"id"`
}





type ConnectionPropertiesResponsePrivateEndpointInput interface {
	pulumi.Input

	ToConnectionPropertiesResponsePrivateEndpointOutput() ConnectionPropertiesResponsePrivateEndpointOutput
	ToConnectionPropertiesResponsePrivateEndpointOutputWithContext(context.Context) ConnectionPropertiesResponsePrivateEndpointOutput
}

type ConnectionPropertiesResponsePrivateEndpointArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ConnectionPropertiesResponsePrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (i ConnectionPropertiesResponsePrivateEndpointArgs) ToConnectionPropertiesResponsePrivateEndpointOutput() ConnectionPropertiesResponsePrivateEndpointOutput {
	return i.ToConnectionPropertiesResponsePrivateEndpointOutputWithContext(context.Background())
}

func (i ConnectionPropertiesResponsePrivateEndpointArgs) ToConnectionPropertiesResponsePrivateEndpointOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesResponsePrivateEndpointOutput)
}

func (i ConnectionPropertiesResponsePrivateEndpointArgs) ToConnectionPropertiesResponsePrivateEndpointPtrOutput() ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return i.ToConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(context.Background())
}

func (i ConnectionPropertiesResponsePrivateEndpointArgs) ToConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesResponsePrivateEndpointOutput).ToConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx)
}









type ConnectionPropertiesResponsePrivateEndpointPtrInput interface {
	pulumi.Input

	ToConnectionPropertiesResponsePrivateEndpointPtrOutput() ConnectionPropertiesResponsePrivateEndpointPtrOutput
	ToConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(context.Context) ConnectionPropertiesResponsePrivateEndpointPtrOutput
}

type connectionPropertiesResponsePrivateEndpointPtrType ConnectionPropertiesResponsePrivateEndpointArgs

func ConnectionPropertiesResponsePrivateEndpointPtr(v *ConnectionPropertiesResponsePrivateEndpointArgs) ConnectionPropertiesResponsePrivateEndpointPtrInput {
	return (*connectionPropertiesResponsePrivateEndpointPtrType)(v)
}

func (*connectionPropertiesResponsePrivateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (i *connectionPropertiesResponsePrivateEndpointPtrType) ToConnectionPropertiesResponsePrivateEndpointPtrOutput() ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return i.ToConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *connectionPropertiesResponsePrivateEndpointPtrType) ToConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

type ConnectionPropertiesResponsePrivateEndpointOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesResponsePrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (o ConnectionPropertiesResponsePrivateEndpointOutput) ToConnectionPropertiesResponsePrivateEndpointOutput() ConnectionPropertiesResponsePrivateEndpointOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateEndpointOutput) ToConnectionPropertiesResponsePrivateEndpointOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateEndpointOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateEndpointOutput) ToConnectionPropertiesResponsePrivateEndpointPtrOutput() ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ToConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(context.Background())
}

func (o ConnectionPropertiesResponsePrivateEndpointOutput) ToConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionPropertiesResponsePrivateEndpoint) *ConnectionPropertiesResponsePrivateEndpoint {
		return &v
	}).(ConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

func (o ConnectionPropertiesResponsePrivateEndpointOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponsePrivateEndpoint) string { return v.Id }).(pulumi.StringOutput)
}

type ConnectionPropertiesResponsePrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesResponsePrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (o ConnectionPropertiesResponsePrivateEndpointPtrOutput) ToConnectionPropertiesResponsePrivateEndpointPtrOutput() ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateEndpointPtrOutput) ToConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateEndpointPtrOutput) Elem() ConnectionPropertiesResponsePrivateEndpointOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateEndpoint) ConnectionPropertiesResponsePrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret ConnectionPropertiesResponsePrivateEndpoint
		return ret
	}).(ConnectionPropertiesResponsePrivateEndpointOutput)
}

func (o ConnectionPropertiesResponsePrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ConnectionPropertiesResponsePrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     string  `pulumi:"description"`
	Status          string  `pulumi:"status"`
}





type ConnectionPropertiesResponsePrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput() ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput
	ToConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutputWithContext(context.Context) ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput
}

type ConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput    `pulumi:"description"`
	Status          pulumi.StringInput    `pulumi:"status"`
}

func (ConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i ConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput() ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return i.ToConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i ConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput)
}

func (i ConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return i.ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i ConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput).ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput
	ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput
}

type connectionPropertiesResponsePrivateLinkServiceConnectionStatePtrType ConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs

func ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtr(v *ConnectionPropertiesResponsePrivateLinkServiceConnectionStateArgs) ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrInput {
	return (*connectionPropertiesResponsePrivateLinkServiceConnectionStatePtrType)(v)
}

func (*connectionPropertiesResponsePrivateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *connectionPropertiesResponsePrivateLinkServiceConnectionStatePtrType) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return i.ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *connectionPropertiesResponsePrivateLinkServiceConnectionStatePtrType) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

type ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput() ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionPropertiesResponsePrivateLinkServiceConnectionState) *ConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		return &v
	}).(ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponsePrivateLinkServiceConnectionState) string { return v.Description }).(pulumi.StringOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionPropertiesResponsePrivateLinkServiceConnectionState) string { return v.Status }).(pulumi.StringOutput)
}

type ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ToConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Elem() ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateLinkServiceConnectionState) ConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret ConnectionPropertiesResponsePrivateLinkServiceConnectionState
		return ret
	}).(ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type DigitalTwinsIdentity struct {
	Type *string `pulumi:"type"`
}





type DigitalTwinsIdentityInput interface {
	pulumi.Input

	ToDigitalTwinsIdentityOutput() DigitalTwinsIdentityOutput
	ToDigitalTwinsIdentityOutputWithContext(context.Context) DigitalTwinsIdentityOutput
}

type DigitalTwinsIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (DigitalTwinsIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsIdentity)(nil)).Elem()
}

func (i DigitalTwinsIdentityArgs) ToDigitalTwinsIdentityOutput() DigitalTwinsIdentityOutput {
	return i.ToDigitalTwinsIdentityOutputWithContext(context.Background())
}

func (i DigitalTwinsIdentityArgs) ToDigitalTwinsIdentityOutputWithContext(ctx context.Context) DigitalTwinsIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsIdentityOutput)
}

func (i DigitalTwinsIdentityArgs) ToDigitalTwinsIdentityPtrOutput() DigitalTwinsIdentityPtrOutput {
	return i.ToDigitalTwinsIdentityPtrOutputWithContext(context.Background())
}

func (i DigitalTwinsIdentityArgs) ToDigitalTwinsIdentityPtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsIdentityOutput).ToDigitalTwinsIdentityPtrOutputWithContext(ctx)
}









type DigitalTwinsIdentityPtrInput interface {
	pulumi.Input

	ToDigitalTwinsIdentityPtrOutput() DigitalTwinsIdentityPtrOutput
	ToDigitalTwinsIdentityPtrOutputWithContext(context.Context) DigitalTwinsIdentityPtrOutput
}

type digitalTwinsIdentityPtrType DigitalTwinsIdentityArgs

func DigitalTwinsIdentityPtr(v *DigitalTwinsIdentityArgs) DigitalTwinsIdentityPtrInput {
	return (*digitalTwinsIdentityPtrType)(v)
}

func (*digitalTwinsIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsIdentity)(nil)).Elem()
}

func (i *digitalTwinsIdentityPtrType) ToDigitalTwinsIdentityPtrOutput() DigitalTwinsIdentityPtrOutput {
	return i.ToDigitalTwinsIdentityPtrOutputWithContext(context.Background())
}

func (i *digitalTwinsIdentityPtrType) ToDigitalTwinsIdentityPtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsIdentityPtrOutput)
}

type DigitalTwinsIdentityOutput struct{ *pulumi.OutputState }

func (DigitalTwinsIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsIdentity)(nil)).Elem()
}

func (o DigitalTwinsIdentityOutput) ToDigitalTwinsIdentityOutput() DigitalTwinsIdentityOutput {
	return o
}

func (o DigitalTwinsIdentityOutput) ToDigitalTwinsIdentityOutputWithContext(ctx context.Context) DigitalTwinsIdentityOutput {
	return o
}

func (o DigitalTwinsIdentityOutput) ToDigitalTwinsIdentityPtrOutput() DigitalTwinsIdentityPtrOutput {
	return o.ToDigitalTwinsIdentityPtrOutputWithContext(context.Background())
}

func (o DigitalTwinsIdentityOutput) ToDigitalTwinsIdentityPtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DigitalTwinsIdentity) *DigitalTwinsIdentity {
		return &v
	}).(DigitalTwinsIdentityPtrOutput)
}

func (o DigitalTwinsIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DigitalTwinsIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type DigitalTwinsIdentityPtrOutput struct{ *pulumi.OutputState }

func (DigitalTwinsIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsIdentity)(nil)).Elem()
}

func (o DigitalTwinsIdentityPtrOutput) ToDigitalTwinsIdentityPtrOutput() DigitalTwinsIdentityPtrOutput {
	return o
}

func (o DigitalTwinsIdentityPtrOutput) ToDigitalTwinsIdentityPtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityPtrOutput {
	return o
}

func (o DigitalTwinsIdentityPtrOutput) Elem() DigitalTwinsIdentityOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentity) DigitalTwinsIdentity {
		if v != nil {
			return *v
		}
		var ret DigitalTwinsIdentity
		return ret
	}).(DigitalTwinsIdentityOutput)
}

func (o DigitalTwinsIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type DigitalTwinsIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type DigitalTwinsIdentityResponseInput interface {
	pulumi.Input

	ToDigitalTwinsIdentityResponseOutput() DigitalTwinsIdentityResponseOutput
	ToDigitalTwinsIdentityResponseOutputWithContext(context.Context) DigitalTwinsIdentityResponseOutput
}

type DigitalTwinsIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (DigitalTwinsIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsIdentityResponse)(nil)).Elem()
}

func (i DigitalTwinsIdentityResponseArgs) ToDigitalTwinsIdentityResponseOutput() DigitalTwinsIdentityResponseOutput {
	return i.ToDigitalTwinsIdentityResponseOutputWithContext(context.Background())
}

func (i DigitalTwinsIdentityResponseArgs) ToDigitalTwinsIdentityResponseOutputWithContext(ctx context.Context) DigitalTwinsIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsIdentityResponseOutput)
}

func (i DigitalTwinsIdentityResponseArgs) ToDigitalTwinsIdentityResponsePtrOutput() DigitalTwinsIdentityResponsePtrOutput {
	return i.ToDigitalTwinsIdentityResponsePtrOutputWithContext(context.Background())
}

func (i DigitalTwinsIdentityResponseArgs) ToDigitalTwinsIdentityResponsePtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsIdentityResponseOutput).ToDigitalTwinsIdentityResponsePtrOutputWithContext(ctx)
}









type DigitalTwinsIdentityResponsePtrInput interface {
	pulumi.Input

	ToDigitalTwinsIdentityResponsePtrOutput() DigitalTwinsIdentityResponsePtrOutput
	ToDigitalTwinsIdentityResponsePtrOutputWithContext(context.Context) DigitalTwinsIdentityResponsePtrOutput
}

type digitalTwinsIdentityResponsePtrType DigitalTwinsIdentityResponseArgs

func DigitalTwinsIdentityResponsePtr(v *DigitalTwinsIdentityResponseArgs) DigitalTwinsIdentityResponsePtrInput {
	return (*digitalTwinsIdentityResponsePtrType)(v)
}

func (*digitalTwinsIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsIdentityResponse)(nil)).Elem()
}

func (i *digitalTwinsIdentityResponsePtrType) ToDigitalTwinsIdentityResponsePtrOutput() DigitalTwinsIdentityResponsePtrOutput {
	return i.ToDigitalTwinsIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *digitalTwinsIdentityResponsePtrType) ToDigitalTwinsIdentityResponsePtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsIdentityResponsePtrOutput)
}

type DigitalTwinsIdentityResponseOutput struct{ *pulumi.OutputState }

func (DigitalTwinsIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DigitalTwinsIdentityResponse)(nil)).Elem()
}

func (o DigitalTwinsIdentityResponseOutput) ToDigitalTwinsIdentityResponseOutput() DigitalTwinsIdentityResponseOutput {
	return o
}

func (o DigitalTwinsIdentityResponseOutput) ToDigitalTwinsIdentityResponseOutputWithContext(ctx context.Context) DigitalTwinsIdentityResponseOutput {
	return o
}

func (o DigitalTwinsIdentityResponseOutput) ToDigitalTwinsIdentityResponsePtrOutput() DigitalTwinsIdentityResponsePtrOutput {
	return o.ToDigitalTwinsIdentityResponsePtrOutputWithContext(context.Background())
}

func (o DigitalTwinsIdentityResponseOutput) ToDigitalTwinsIdentityResponsePtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DigitalTwinsIdentityResponse) *DigitalTwinsIdentityResponse {
		return &v
	}).(DigitalTwinsIdentityResponsePtrOutput)
}

func (o DigitalTwinsIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v DigitalTwinsIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o DigitalTwinsIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v DigitalTwinsIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o DigitalTwinsIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DigitalTwinsIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type DigitalTwinsIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (DigitalTwinsIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsIdentityResponse)(nil)).Elem()
}

func (o DigitalTwinsIdentityResponsePtrOutput) ToDigitalTwinsIdentityResponsePtrOutput() DigitalTwinsIdentityResponsePtrOutput {
	return o
}

func (o DigitalTwinsIdentityResponsePtrOutput) ToDigitalTwinsIdentityResponsePtrOutputWithContext(ctx context.Context) DigitalTwinsIdentityResponsePtrOutput {
	return o
}

func (o DigitalTwinsIdentityResponsePtrOutput) Elem() DigitalTwinsIdentityResponseOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentityResponse) DigitalTwinsIdentityResponse {
		if v != nil {
			return *v
		}
		var ret DigitalTwinsIdentityResponse
		return ret
	}).(DigitalTwinsIdentityResponseOutput)
}

func (o DigitalTwinsIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o DigitalTwinsIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o DigitalTwinsIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DigitalTwinsIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type EventGrid struct {
	AccessKey1         string  `pulumi:"accessKey1"`
	AccessKey2         *string `pulumi:"accessKey2"`
	AuthenticationType *string `pulumi:"authenticationType"`
	DeadLetterSecret   *string `pulumi:"deadLetterSecret"`
	DeadLetterUri      *string `pulumi:"deadLetterUri"`
	EndpointType       string  `pulumi:"endpointType"`
	TopicEndpoint      string  `pulumi:"topicEndpoint"`
}





type EventGridInput interface {
	pulumi.Input

	ToEventGridOutput() EventGridOutput
	ToEventGridOutputWithContext(context.Context) EventGridOutput
}

type EventGridArgs struct {
	AccessKey1         pulumi.StringInput    `pulumi:"accessKey1"`
	AccessKey2         pulumi.StringPtrInput `pulumi:"accessKey2"`
	AuthenticationType pulumi.StringPtrInput `pulumi:"authenticationType"`
	DeadLetterSecret   pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	DeadLetterUri      pulumi.StringPtrInput `pulumi:"deadLetterUri"`
	EndpointType       pulumi.StringInput    `pulumi:"endpointType"`
	TopicEndpoint      pulumi.StringInput    `pulumi:"topicEndpoint"`
}

func (EventGridArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGrid)(nil)).Elem()
}

func (i EventGridArgs) ToEventGridOutput() EventGridOutput {
	return i.ToEventGridOutputWithContext(context.Background())
}

func (i EventGridArgs) ToEventGridOutputWithContext(ctx context.Context) EventGridOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridOutput)
}

type EventGridOutput struct{ *pulumi.OutputState }

func (EventGridOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGrid)(nil)).Elem()
}

func (o EventGridOutput) ToEventGridOutput() EventGridOutput {
	return o
}

func (o EventGridOutput) ToEventGridOutputWithContext(ctx context.Context) EventGridOutput {
	return o
}

func (o EventGridOutput) AccessKey1() pulumi.StringOutput {
	return o.ApplyT(func(v EventGrid) string { return v.AccessKey1 }).(pulumi.StringOutput)
}

func (o EventGridOutput) AccessKey2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGrid) *string { return v.AccessKey2 }).(pulumi.StringPtrOutput)
}

func (o EventGridOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGrid) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o EventGridOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGrid) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o EventGridOutput) DeadLetterUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGrid) *string { return v.DeadLetterUri }).(pulumi.StringPtrOutput)
}

func (o EventGridOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventGrid) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventGridOutput) TopicEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v EventGrid) string { return v.TopicEndpoint }).(pulumi.StringOutput)
}

type EventGridResponse struct {
	AccessKey1         string  `pulumi:"accessKey1"`
	AccessKey2         *string `pulumi:"accessKey2"`
	AuthenticationType *string `pulumi:"authenticationType"`
	CreatedTime        string  `pulumi:"createdTime"`
	DeadLetterSecret   *string `pulumi:"deadLetterSecret"`
	DeadLetterUri      *string `pulumi:"deadLetterUri"`
	EndpointType       string  `pulumi:"endpointType"`
	ProvisioningState  string  `pulumi:"provisioningState"`
	TopicEndpoint      string  `pulumi:"topicEndpoint"`
}





type EventGridResponseInput interface {
	pulumi.Input

	ToEventGridResponseOutput() EventGridResponseOutput
	ToEventGridResponseOutputWithContext(context.Context) EventGridResponseOutput
}

type EventGridResponseArgs struct {
	AccessKey1         pulumi.StringInput    `pulumi:"accessKey1"`
	AccessKey2         pulumi.StringPtrInput `pulumi:"accessKey2"`
	AuthenticationType pulumi.StringPtrInput `pulumi:"authenticationType"`
	CreatedTime        pulumi.StringInput    `pulumi:"createdTime"`
	DeadLetterSecret   pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	DeadLetterUri      pulumi.StringPtrInput `pulumi:"deadLetterUri"`
	EndpointType       pulumi.StringInput    `pulumi:"endpointType"`
	ProvisioningState  pulumi.StringInput    `pulumi:"provisioningState"`
	TopicEndpoint      pulumi.StringInput    `pulumi:"topicEndpoint"`
}

func (EventGridResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridResponse)(nil)).Elem()
}

func (i EventGridResponseArgs) ToEventGridResponseOutput() EventGridResponseOutput {
	return i.ToEventGridResponseOutputWithContext(context.Background())
}

func (i EventGridResponseArgs) ToEventGridResponseOutputWithContext(ctx context.Context) EventGridResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGridResponseOutput)
}

type EventGridResponseOutput struct{ *pulumi.OutputState }

func (EventGridResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGridResponse)(nil)).Elem()
}

func (o EventGridResponseOutput) ToEventGridResponseOutput() EventGridResponseOutput {
	return o
}

func (o EventGridResponseOutput) ToEventGridResponseOutputWithContext(ctx context.Context) EventGridResponseOutput {
	return o
}

func (o EventGridResponseOutput) AccessKey1() pulumi.StringOutput {
	return o.ApplyT(func(v EventGridResponse) string { return v.AccessKey1 }).(pulumi.StringOutput)
}

func (o EventGridResponseOutput) AccessKey2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGridResponse) *string { return v.AccessKey2 }).(pulumi.StringPtrOutput)
}

func (o EventGridResponseOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGridResponse) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o EventGridResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v EventGridResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o EventGridResponseOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGridResponse) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o EventGridResponseOutput) DeadLetterUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGridResponse) *string { return v.DeadLetterUri }).(pulumi.StringPtrOutput)
}

func (o EventGridResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventGridResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventGridResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v EventGridResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EventGridResponseOutput) TopicEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v EventGridResponse) string { return v.TopicEndpoint }).(pulumi.StringOutput)
}

type EventHub struct {
	AuthenticationType           *string `pulumi:"authenticationType"`
	ConnectionStringPrimaryKey   *string `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey *string `pulumi:"connectionStringSecondaryKey"`
	DeadLetterSecret             *string `pulumi:"deadLetterSecret"`
	DeadLetterUri                *string `pulumi:"deadLetterUri"`
	EndpointType                 string  `pulumi:"endpointType"`
	EndpointUri                  *string `pulumi:"endpointUri"`
	EntityPath                   *string `pulumi:"entityPath"`
}





type EventHubInput interface {
	pulumi.Input

	ToEventHubOutput() EventHubOutput
	ToEventHubOutputWithContext(context.Context) EventHubOutput
}

type EventHubArgs struct {
	AuthenticationType           pulumi.StringPtrInput `pulumi:"authenticationType"`
	ConnectionStringPrimaryKey   pulumi.StringPtrInput `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey pulumi.StringPtrInput `pulumi:"connectionStringSecondaryKey"`
	DeadLetterSecret             pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	DeadLetterUri                pulumi.StringPtrInput `pulumi:"deadLetterUri"`
	EndpointType                 pulumi.StringInput    `pulumi:"endpointType"`
	EndpointUri                  pulumi.StringPtrInput `pulumi:"endpointUri"`
	EntityPath                   pulumi.StringPtrInput `pulumi:"entityPath"`
}

func (EventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHub)(nil)).Elem()
}

func (i EventHubArgs) ToEventHubOutput() EventHubOutput {
	return i.ToEventHubOutputWithContext(context.Background())
}

func (i EventHubArgs) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubOutput)
}

type EventHubOutput struct{ *pulumi.OutputState }

func (EventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHub)(nil)).Elem()
}

func (o EventHubOutput) ToEventHubOutput() EventHubOutput {
	return o
}

func (o EventHubOutput) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return o
}

func (o EventHubOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHub) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o EventHubOutput) ConnectionStringPrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHub) *string { return v.ConnectionStringPrimaryKey }).(pulumi.StringPtrOutput)
}

func (o EventHubOutput) ConnectionStringSecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHub) *string { return v.ConnectionStringSecondaryKey }).(pulumi.StringPtrOutput)
}

func (o EventHubOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHub) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o EventHubOutput) DeadLetterUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHub) *string { return v.DeadLetterUri }).(pulumi.StringPtrOutput)
}

func (o EventHubOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventHub) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventHubOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHub) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o EventHubOutput) EntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHub) *string { return v.EntityPath }).(pulumi.StringPtrOutput)
}

type EventHubResponse struct {
	AuthenticationType           *string `pulumi:"authenticationType"`
	ConnectionStringPrimaryKey   *string `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey *string `pulumi:"connectionStringSecondaryKey"`
	CreatedTime                  string  `pulumi:"createdTime"`
	DeadLetterSecret             *string `pulumi:"deadLetterSecret"`
	DeadLetterUri                *string `pulumi:"deadLetterUri"`
	EndpointType                 string  `pulumi:"endpointType"`
	EndpointUri                  *string `pulumi:"endpointUri"`
	EntityPath                   *string `pulumi:"entityPath"`
	ProvisioningState            string  `pulumi:"provisioningState"`
}





type EventHubResponseInput interface {
	pulumi.Input

	ToEventHubResponseOutput() EventHubResponseOutput
	ToEventHubResponseOutputWithContext(context.Context) EventHubResponseOutput
}

type EventHubResponseArgs struct {
	AuthenticationType           pulumi.StringPtrInput `pulumi:"authenticationType"`
	ConnectionStringPrimaryKey   pulumi.StringPtrInput `pulumi:"connectionStringPrimaryKey"`
	ConnectionStringSecondaryKey pulumi.StringPtrInput `pulumi:"connectionStringSecondaryKey"`
	CreatedTime                  pulumi.StringInput    `pulumi:"createdTime"`
	DeadLetterSecret             pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	DeadLetterUri                pulumi.StringPtrInput `pulumi:"deadLetterUri"`
	EndpointType                 pulumi.StringInput    `pulumi:"endpointType"`
	EndpointUri                  pulumi.StringPtrInput `pulumi:"endpointUri"`
	EntityPath                   pulumi.StringPtrInput `pulumi:"entityPath"`
	ProvisioningState            pulumi.StringInput    `pulumi:"provisioningState"`
}

func (EventHubResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubResponse)(nil)).Elem()
}

func (i EventHubResponseArgs) ToEventHubResponseOutput() EventHubResponseOutput {
	return i.ToEventHubResponseOutputWithContext(context.Background())
}

func (i EventHubResponseArgs) ToEventHubResponseOutputWithContext(ctx context.Context) EventHubResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubResponseOutput)
}

type EventHubResponseOutput struct{ *pulumi.OutputState }

func (EventHubResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubResponse)(nil)).Elem()
}

func (o EventHubResponseOutput) ToEventHubResponseOutput() EventHubResponseOutput {
	return o
}

func (o EventHubResponseOutput) ToEventHubResponseOutputWithContext(ctx context.Context) EventHubResponseOutput {
	return o
}

func (o EventHubResponseOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubResponse) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o EventHubResponseOutput) ConnectionStringPrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubResponse) *string { return v.ConnectionStringPrimaryKey }).(pulumi.StringPtrOutput)
}

func (o EventHubResponseOutput) ConnectionStringSecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubResponse) *string { return v.ConnectionStringSecondaryKey }).(pulumi.StringPtrOutput)
}

func (o EventHubResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o EventHubResponseOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubResponse) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o EventHubResponseOutput) DeadLetterUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubResponse) *string { return v.DeadLetterUri }).(pulumi.StringPtrOutput)
}

func (o EventHubResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventHubResponseOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubResponse) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o EventHubResponseOutput) EntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubResponse) *string { return v.EntityPath }).(pulumi.StringPtrOutput)
}

func (o EventHubResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionType struct {
	Properties PrivateEndpointConnectionProperties `pulumi:"properties"`
}





type PrivateEndpointConnectionTypeInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput
	ToPrivateEndpointConnectionTypeOutputWithContext(context.Context) PrivateEndpointConnectionTypeOutput
}

type PrivateEndpointConnectionTypeArgs struct {
	Properties PrivateEndpointConnectionPropertiesInput `pulumi:"properties"`
}

func (PrivateEndpointConnectionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return i.ToPrivateEndpointConnectionTypeOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArgs) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeOutput)
}





type PrivateEndpointConnectionTypeArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput
	ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Context) PrivateEndpointConnectionTypeArrayOutput
}

type PrivateEndpointConnectionTypeArray []PrivateEndpointConnectionTypeInput

func (PrivateEndpointConnectionTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return i.ToPrivateEndpointConnectionTypeArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionTypeArray) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionTypeArrayOutput)
}

type PrivateEndpointConnectionTypeOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutput() PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) ToPrivateEndpointConnectionTypeOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeOutput {
	return o
}

func (o PrivateEndpointConnectionTypeOutput) Properties() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionType) PrivateEndpointConnectionProperties { return v.Properties }).(PrivateEndpointConnectionPropertiesOutput)
}

type PrivateEndpointConnectionTypeArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionType)(nil)).Elem()
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutput() PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) ToPrivateEndpointConnectionTypeArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionTypeArrayOutput {
	return o
}

func (o PrivateEndpointConnectionTypeArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionType {
		return vs[0].([]PrivateEndpointConnectionType)[vs[1].(int)]
	}).(PrivateEndpointConnectionTypeOutput)
}

type PrivateEndpointConnectionProperties struct {
	GroupIds                          []string                                               `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState *ConnectionPropertiesPrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	GroupIds                          pulumi.StringArrayInput                                       `pulumi:"groupIds"`
	PrivateLinkServiceConnectionState ConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return i.ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput)
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput).ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput
	ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPtrOutput
}

type privateEndpointConnectionPropertiesPtrType PrivateEndpointConnectionPropertiesArgs

func PrivateEndpointConnectionPropertiesPtr(v *PrivateEndpointConnectionPropertiesArgs) PrivateEndpointConnectionPropertiesPtrInput {
	return (*privateEndpointConnectionPropertiesPtrType)(v)
}

func (*privateEndpointConnectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPtrOutput)
}

type PrivateEndpointConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionProperties {
		return &v
	}).(PrivateEndpointConnectionPropertiesPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *ConnectionPropertiesPrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) Elem() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) PrivateEndpointConnectionProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionProperties
		return ret
	}).(PrivateEndpointConnectionPropertiesOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateLinkServiceConnectionState() ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *ConnectionPropertiesPrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionResponseProperties `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id         pulumi.StringInput                               `pulumi:"id"`
	Name       pulumi.StringInput                               `pulumi:"name"`
	Properties PrivateEndpointConnectionResponsePropertiesInput `pulumi:"properties"`
	Type       pulumi.StringInput                               `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionResponsePropertiesOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointConnectionResponseProperties {
		return v.Properties
	}).(PrivateEndpointConnectionResponsePropertiesOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointConnectionResponseProperties struct {
	GroupIds                          []string                                                       `pulumi:"groupIds"`
	PrivateEndpoint                   *ConnectionPropertiesResponsePrivateEndpoint                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ConnectionPropertiesResponsePrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                                         `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionResponsePropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponsePropertiesOutput() PrivateEndpointConnectionResponsePropertiesOutput
	ToPrivateEndpointConnectionResponsePropertiesOutputWithContext(context.Context) PrivateEndpointConnectionResponsePropertiesOutput
}

type PrivateEndpointConnectionResponsePropertiesArgs struct {
	GroupIds                          pulumi.StringArrayInput                                               `pulumi:"groupIds"`
	PrivateEndpoint                   ConnectionPropertiesResponsePrivateEndpointPtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                                    `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponseProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponsePropertiesArgs) ToPrivateEndpointConnectionResponsePropertiesOutput() PrivateEndpointConnectionResponsePropertiesOutput {
	return i.ToPrivateEndpointConnectionResponsePropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponsePropertiesArgs) ToPrivateEndpointConnectionResponsePropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponsePropertiesOutput)
}

func (i PrivateEndpointConnectionResponsePropertiesArgs) ToPrivateEndpointConnectionResponsePropertiesPtrOutput() PrivateEndpointConnectionResponsePropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponsePropertiesArgs) ToPrivateEndpointConnectionResponsePropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponsePropertiesOutput).ToPrivateEndpointConnectionResponsePropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionResponsePropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponsePropertiesPtrOutput() PrivateEndpointConnectionResponsePropertiesPtrOutput
	ToPrivateEndpointConnectionResponsePropertiesPtrOutputWithContext(context.Context) PrivateEndpointConnectionResponsePropertiesPtrOutput
}

type privateEndpointConnectionResponsePropertiesPtrType PrivateEndpointConnectionResponsePropertiesArgs

func PrivateEndpointConnectionResponsePropertiesPtr(v *PrivateEndpointConnectionResponsePropertiesArgs) PrivateEndpointConnectionResponsePropertiesPtrInput {
	return (*privateEndpointConnectionResponsePropertiesPtrType)(v)
}

func (*privateEndpointConnectionResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionResponseProperties)(nil)).Elem()
}

func (i *privateEndpointConnectionResponsePropertiesPtrType) ToPrivateEndpointConnectionResponsePropertiesPtrOutput() PrivateEndpointConnectionResponsePropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionResponsePropertiesPtrType) ToPrivateEndpointConnectionResponsePropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponsePropertiesPtrOutput)
}

type PrivateEndpointConnectionResponsePropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponseProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) ToPrivateEndpointConnectionResponsePropertiesOutput() PrivateEndpointConnectionResponsePropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) ToPrivateEndpointConnectionResponsePropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) ToPrivateEndpointConnectionResponsePropertiesPtrOutput() PrivateEndpointConnectionResponsePropertiesPtrOutput {
	return o.ToPrivateEndpointConnectionResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) ToPrivateEndpointConnectionResponsePropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionResponseProperties) *PrivateEndpointConnectionResponseProperties {
		return &v
	}).(PrivateEndpointConnectionResponsePropertiesPtrOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponseProperties) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) PrivateEndpoint() ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponseProperties) *ConnectionPropertiesResponsePrivateEndpoint {
		return v.PrivateEndpoint
	}).(ConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) PrivateLinkServiceConnectionState() ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponseProperties) *ConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponseProperties) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionResponseProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponsePropertiesPtrOutput) ToPrivateEndpointConnectionResponsePropertiesPtrOutput() PrivateEndpointConnectionResponsePropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionResponsePropertiesPtrOutput) ToPrivateEndpointConnectionResponsePropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponsePropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionResponsePropertiesPtrOutput) Elem() PrivateEndpointConnectionResponsePropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionResponseProperties) PrivateEndpointConnectionResponseProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionResponseProperties
		return ret
	}).(PrivateEndpointConnectionResponsePropertiesOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesPtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionResponseProperties) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesPtrOutput) PrivateEndpoint() ConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionResponseProperties) *ConnectionPropertiesResponsePrivateEndpoint {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(ConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesPtrOutput) PrivateLinkServiceConnectionState() ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionResponseProperties) *ConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionResponsePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionResponseProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type ServiceBus struct {
	AuthenticationType        *string `pulumi:"authenticationType"`
	DeadLetterSecret          *string `pulumi:"deadLetterSecret"`
	DeadLetterUri             *string `pulumi:"deadLetterUri"`
	EndpointType              string  `pulumi:"endpointType"`
	EndpointUri               *string `pulumi:"endpointUri"`
	EntityPath                *string `pulumi:"entityPath"`
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
}





type ServiceBusInput interface {
	pulumi.Input

	ToServiceBusOutput() ServiceBusOutput
	ToServiceBusOutputWithContext(context.Context) ServiceBusOutput
}

type ServiceBusArgs struct {
	AuthenticationType        pulumi.StringPtrInput `pulumi:"authenticationType"`
	DeadLetterSecret          pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	DeadLetterUri             pulumi.StringPtrInput `pulumi:"deadLetterUri"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	EndpointUri               pulumi.StringPtrInput `pulumi:"endpointUri"`
	EntityPath                pulumi.StringPtrInput `pulumi:"entityPath"`
	PrimaryConnectionString   pulumi.StringPtrInput `pulumi:"primaryConnectionString"`
	SecondaryConnectionString pulumi.StringPtrInput `pulumi:"secondaryConnectionString"`
}

func (ServiceBusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBus)(nil)).Elem()
}

func (i ServiceBusArgs) ToServiceBusOutput() ServiceBusOutput {
	return i.ToServiceBusOutputWithContext(context.Background())
}

func (i ServiceBusArgs) ToServiceBusOutputWithContext(ctx context.Context) ServiceBusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusOutput)
}

type ServiceBusOutput struct{ *pulumi.OutputState }

func (ServiceBusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBus)(nil)).Elem()
}

func (o ServiceBusOutput) ToServiceBusOutput() ServiceBusOutput {
	return o
}

func (o ServiceBusOutput) ToServiceBusOutputWithContext(ctx context.Context) ServiceBusOutput {
	return o
}

func (o ServiceBusOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBus) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o ServiceBusOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBus) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o ServiceBusOutput) DeadLetterUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBus) *string { return v.DeadLetterUri }).(pulumi.StringPtrOutput)
}

func (o ServiceBusOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBus) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o ServiceBusOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBus) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o ServiceBusOutput) EntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBus) *string { return v.EntityPath }).(pulumi.StringPtrOutput)
}

func (o ServiceBusOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBus) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ServiceBusOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBus) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

type ServiceBusResponse struct {
	AuthenticationType        *string `pulumi:"authenticationType"`
	CreatedTime               string  `pulumi:"createdTime"`
	DeadLetterSecret          *string `pulumi:"deadLetterSecret"`
	DeadLetterUri             *string `pulumi:"deadLetterUri"`
	EndpointType              string  `pulumi:"endpointType"`
	EndpointUri               *string `pulumi:"endpointUri"`
	EntityPath                *string `pulumi:"entityPath"`
	PrimaryConnectionString   *string `pulumi:"primaryConnectionString"`
	ProvisioningState         string  `pulumi:"provisioningState"`
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
}





type ServiceBusResponseInput interface {
	pulumi.Input

	ToServiceBusResponseOutput() ServiceBusResponseOutput
	ToServiceBusResponseOutputWithContext(context.Context) ServiceBusResponseOutput
}

type ServiceBusResponseArgs struct {
	AuthenticationType        pulumi.StringPtrInput `pulumi:"authenticationType"`
	CreatedTime               pulumi.StringInput    `pulumi:"createdTime"`
	DeadLetterSecret          pulumi.StringPtrInput `pulumi:"deadLetterSecret"`
	DeadLetterUri             pulumi.StringPtrInput `pulumi:"deadLetterUri"`
	EndpointType              pulumi.StringInput    `pulumi:"endpointType"`
	EndpointUri               pulumi.StringPtrInput `pulumi:"endpointUri"`
	EntityPath                pulumi.StringPtrInput `pulumi:"entityPath"`
	PrimaryConnectionString   pulumi.StringPtrInput `pulumi:"primaryConnectionString"`
	ProvisioningState         pulumi.StringInput    `pulumi:"provisioningState"`
	SecondaryConnectionString pulumi.StringPtrInput `pulumi:"secondaryConnectionString"`
}

func (ServiceBusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusResponse)(nil)).Elem()
}

func (i ServiceBusResponseArgs) ToServiceBusResponseOutput() ServiceBusResponseOutput {
	return i.ToServiceBusResponseOutputWithContext(context.Background())
}

func (i ServiceBusResponseArgs) ToServiceBusResponseOutputWithContext(ctx context.Context) ServiceBusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceBusResponseOutput)
}

type ServiceBusResponseOutput struct{ *pulumi.OutputState }

func (ServiceBusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceBusResponse)(nil)).Elem()
}

func (o ServiceBusResponseOutput) ToServiceBusResponseOutput() ServiceBusResponseOutput {
	return o
}

func (o ServiceBusResponseOutput) ToServiceBusResponseOutputWithContext(ctx context.Context) ServiceBusResponseOutput {
	return o
}

func (o ServiceBusResponseOutput) AuthenticationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusResponse) *string { return v.AuthenticationType }).(pulumi.StringPtrOutput)
}

func (o ServiceBusResponseOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusResponse) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o ServiceBusResponseOutput) DeadLetterSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusResponse) *string { return v.DeadLetterSecret }).(pulumi.StringPtrOutput)
}

func (o ServiceBusResponseOutput) DeadLetterUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusResponse) *string { return v.DeadLetterUri }).(pulumi.StringPtrOutput)
}

func (o ServiceBusResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o ServiceBusResponseOutput) EndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusResponse) *string { return v.EndpointUri }).(pulumi.StringPtrOutput)
}

func (o ServiceBusResponseOutput) EntityPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusResponse) *string { return v.EntityPath }).(pulumi.StringPtrOutput)
}

func (o ServiceBusResponseOutput) PrimaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusResponse) *string { return v.PrimaryConnectionString }).(pulumi.StringPtrOutput)
}

func (o ServiceBusResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceBusResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceBusResponseOutput) SecondaryConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceBusResponse) *string { return v.SecondaryConnectionString }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ConnectionPropertiesPrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponsePrivateEndpointOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponsePrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityPtrOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityResponseOutput{})
	pulumi.RegisterOutputType(DigitalTwinsIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(EventGridOutput{})
	pulumi.RegisterOutputType(EventGridResponseOutput{})
	pulumi.RegisterOutputType(EventHubOutput{})
	pulumi.RegisterOutputType(EventHubResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionTypeArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponsePropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ServiceBusOutput{})
	pulumi.RegisterOutputType(ServiceBusResponseOutput{})
}
