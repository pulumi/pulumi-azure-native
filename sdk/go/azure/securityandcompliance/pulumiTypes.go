


package securityandcompliance

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnection struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput
	ToPrivateEndpointConnectionOutputWithContext(context.Context) PrivateEndpointConnectionOutput
}

type PrivateEndpointConnectionArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnection)(nil)).Elem()
}

func (i PrivateEndpointConnectionArgs) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionArgs) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionOutput)
}





type PrivateEndpointConnectionArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionArrayOutput() PrivateEndpointConnectionArrayOutput
	ToPrivateEndpointConnectionArrayOutputWithContext(context.Context) PrivateEndpointConnectionArrayOutput
}

type PrivateEndpointConnectionArray []PrivateEndpointConnectionInput

func (PrivateEndpointConnectionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnection)(nil)).Elem()
}

func (i PrivateEndpointConnectionArray) ToPrivateEndpointConnectionArrayOutput() PrivateEndpointConnectionArrayOutput {
	return i.ToPrivateEndpointConnectionArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionArray) ToPrivateEndpointConnectionArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionArrayOutput)
}

type PrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v PrivateEndpointConnection) PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateEndpointConnectionArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionArrayOutput) ToPrivateEndpointConnectionArrayOutput() PrivateEndpointConnectionArrayOutput {
	return o
}

func (o PrivateEndpointConnectionArrayOutput) ToPrivateEndpointConnectionArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionArrayOutput {
	return o
}

func (o PrivateEndpointConnectionArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnection {
		return vs[0].([]PrivateEndpointConnection)[vs[1].(int)]
	}).(PrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                             `pulumi:"id"`
	Name                              pulumi.StringInput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseInput                        `pulumi:"systemData"`
	Type                              pulumi.StringInput                             `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
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

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ServiceAccessPolicyEntry struct {
	ObjectId string `pulumi:"objectId"`
}





type ServiceAccessPolicyEntryInput interface {
	pulumi.Input

	ToServiceAccessPolicyEntryOutput() ServiceAccessPolicyEntryOutput
	ToServiceAccessPolicyEntryOutputWithContext(context.Context) ServiceAccessPolicyEntryOutput
}

type ServiceAccessPolicyEntryArgs struct {
	ObjectId pulumi.StringInput `pulumi:"objectId"`
}

func (ServiceAccessPolicyEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccessPolicyEntry)(nil)).Elem()
}

func (i ServiceAccessPolicyEntryArgs) ToServiceAccessPolicyEntryOutput() ServiceAccessPolicyEntryOutput {
	return i.ToServiceAccessPolicyEntryOutputWithContext(context.Background())
}

func (i ServiceAccessPolicyEntryArgs) ToServiceAccessPolicyEntryOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccessPolicyEntryOutput)
}





type ServiceAccessPolicyEntryArrayInput interface {
	pulumi.Input

	ToServiceAccessPolicyEntryArrayOutput() ServiceAccessPolicyEntryArrayOutput
	ToServiceAccessPolicyEntryArrayOutputWithContext(context.Context) ServiceAccessPolicyEntryArrayOutput
}

type ServiceAccessPolicyEntryArray []ServiceAccessPolicyEntryInput

func (ServiceAccessPolicyEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAccessPolicyEntry)(nil)).Elem()
}

func (i ServiceAccessPolicyEntryArray) ToServiceAccessPolicyEntryArrayOutput() ServiceAccessPolicyEntryArrayOutput {
	return i.ToServiceAccessPolicyEntryArrayOutputWithContext(context.Background())
}

func (i ServiceAccessPolicyEntryArray) ToServiceAccessPolicyEntryArrayOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccessPolicyEntryArrayOutput)
}

type ServiceAccessPolicyEntryOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccessPolicyEntry)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryOutput) ToServiceAccessPolicyEntryOutput() ServiceAccessPolicyEntryOutput {
	return o
}

func (o ServiceAccessPolicyEntryOutput) ToServiceAccessPolicyEntryOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryOutput {
	return o
}

func (o ServiceAccessPolicyEntryOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAccessPolicyEntry) string { return v.ObjectId }).(pulumi.StringOutput)
}

type ServiceAccessPolicyEntryArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAccessPolicyEntry)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryArrayOutput) ToServiceAccessPolicyEntryArrayOutput() ServiceAccessPolicyEntryArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryArrayOutput) ToServiceAccessPolicyEntryArrayOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryArrayOutput) Index(i pulumi.IntInput) ServiceAccessPolicyEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceAccessPolicyEntry {
		return vs[0].([]ServiceAccessPolicyEntry)[vs[1].(int)]
	}).(ServiceAccessPolicyEntryOutput)
}

type ServiceAccessPolicyEntryResponse struct {
	ObjectId string `pulumi:"objectId"`
}





type ServiceAccessPolicyEntryResponseInput interface {
	pulumi.Input

	ToServiceAccessPolicyEntryResponseOutput() ServiceAccessPolicyEntryResponseOutput
	ToServiceAccessPolicyEntryResponseOutputWithContext(context.Context) ServiceAccessPolicyEntryResponseOutput
}

type ServiceAccessPolicyEntryResponseArgs struct {
	ObjectId pulumi.StringInput `pulumi:"objectId"`
}

func (ServiceAccessPolicyEntryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccessPolicyEntryResponse)(nil)).Elem()
}

func (i ServiceAccessPolicyEntryResponseArgs) ToServiceAccessPolicyEntryResponseOutput() ServiceAccessPolicyEntryResponseOutput {
	return i.ToServiceAccessPolicyEntryResponseOutputWithContext(context.Background())
}

func (i ServiceAccessPolicyEntryResponseArgs) ToServiceAccessPolicyEntryResponseOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccessPolicyEntryResponseOutput)
}





type ServiceAccessPolicyEntryResponseArrayInput interface {
	pulumi.Input

	ToServiceAccessPolicyEntryResponseArrayOutput() ServiceAccessPolicyEntryResponseArrayOutput
	ToServiceAccessPolicyEntryResponseArrayOutputWithContext(context.Context) ServiceAccessPolicyEntryResponseArrayOutput
}

type ServiceAccessPolicyEntryResponseArray []ServiceAccessPolicyEntryResponseInput

func (ServiceAccessPolicyEntryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAccessPolicyEntryResponse)(nil)).Elem()
}

func (i ServiceAccessPolicyEntryResponseArray) ToServiceAccessPolicyEntryResponseArrayOutput() ServiceAccessPolicyEntryResponseArrayOutput {
	return i.ToServiceAccessPolicyEntryResponseArrayOutputWithContext(context.Background())
}

func (i ServiceAccessPolicyEntryResponseArray) ToServiceAccessPolicyEntryResponseArrayOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAccessPolicyEntryResponseArrayOutput)
}

type ServiceAccessPolicyEntryResponseOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAccessPolicyEntryResponse)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryResponseOutput) ToServiceAccessPolicyEntryResponseOutput() ServiceAccessPolicyEntryResponseOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseOutput) ToServiceAccessPolicyEntryResponseOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryResponseOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceAccessPolicyEntryResponse) string { return v.ObjectId }).(pulumi.StringOutput)
}

type ServiceAccessPolicyEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceAccessPolicyEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceAccessPolicyEntryResponse)(nil)).Elem()
}

func (o ServiceAccessPolicyEntryResponseArrayOutput) ToServiceAccessPolicyEntryResponseArrayOutput() ServiceAccessPolicyEntryResponseArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseArrayOutput) ToServiceAccessPolicyEntryResponseArrayOutputWithContext(ctx context.Context) ServiceAccessPolicyEntryResponseArrayOutput {
	return o
}

func (o ServiceAccessPolicyEntryResponseArrayOutput) Index(i pulumi.IntInput) ServiceAccessPolicyEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceAccessPolicyEntryResponse {
		return vs[0].([]ServiceAccessPolicyEntryResponse)[vs[1].(int)]
	}).(ServiceAccessPolicyEntryResponseOutput)
}

type ServiceAuthenticationConfigurationInfo struct {
	Audience          *string `pulumi:"audience"`
	Authority         *string `pulumi:"authority"`
	SmartProxyEnabled *bool   `pulumi:"smartProxyEnabled"`
}





type ServiceAuthenticationConfigurationInfoInput interface {
	pulumi.Input

	ToServiceAuthenticationConfigurationInfoOutput() ServiceAuthenticationConfigurationInfoOutput
	ToServiceAuthenticationConfigurationInfoOutputWithContext(context.Context) ServiceAuthenticationConfigurationInfoOutput
}

type ServiceAuthenticationConfigurationInfoArgs struct {
	Audience          pulumi.StringPtrInput `pulumi:"audience"`
	Authority         pulumi.StringPtrInput `pulumi:"authority"`
	SmartProxyEnabled pulumi.BoolPtrInput   `pulumi:"smartProxyEnabled"`
}

func (ServiceAuthenticationConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoOutput() ServiceAuthenticationConfigurationInfoOutput {
	return i.ToServiceAuthenticationConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoOutput)
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return i.ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceAuthenticationConfigurationInfoArgs) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoOutput).ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceAuthenticationConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput
	ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Context) ServiceAuthenticationConfigurationInfoPtrOutput
}

type serviceAuthenticationConfigurationInfoPtrType ServiceAuthenticationConfigurationInfoArgs

func ServiceAuthenticationConfigurationInfoPtr(v *ServiceAuthenticationConfigurationInfoArgs) ServiceAuthenticationConfigurationInfoPtrInput {
	return (*serviceAuthenticationConfigurationInfoPtrType)(v)
}

func (*serviceAuthenticationConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (i *serviceAuthenticationConfigurationInfoPtrType) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return i.ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceAuthenticationConfigurationInfoPtrType) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

type ServiceAuthenticationConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoOutput() ServiceAuthenticationConfigurationInfoOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceAuthenticationConfigurationInfoOutput) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceAuthenticationConfigurationInfo) *ServiceAuthenticationConfigurationInfo {
		return &v
	}).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfo) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfo) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfo) *bool { return v.SmartProxyEnabled }).(pulumi.BoolPtrOutput)
}

type ServiceAuthenticationConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfigurationInfo)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) ToServiceAuthenticationConfigurationInfoPtrOutput() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) ToServiceAuthenticationConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoPtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) Elem() ServiceAuthenticationConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) ServiceAuthenticationConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceAuthenticationConfigurationInfo
		return ret
	}).(ServiceAuthenticationConfigurationInfoOutput)
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoPtrOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfo) *bool {
		if v == nil {
			return nil
		}
		return v.SmartProxyEnabled
	}).(pulumi.BoolPtrOutput)
}

type ServiceAuthenticationConfigurationInfoResponse struct {
	Audience          *string `pulumi:"audience"`
	Authority         *string `pulumi:"authority"`
	SmartProxyEnabled *bool   `pulumi:"smartProxyEnabled"`
}





type ServiceAuthenticationConfigurationInfoResponseInput interface {
	pulumi.Input

	ToServiceAuthenticationConfigurationInfoResponseOutput() ServiceAuthenticationConfigurationInfoResponseOutput
	ToServiceAuthenticationConfigurationInfoResponseOutputWithContext(context.Context) ServiceAuthenticationConfigurationInfoResponseOutput
}

type ServiceAuthenticationConfigurationInfoResponseArgs struct {
	Audience          pulumi.StringPtrInput `pulumi:"audience"`
	Authority         pulumi.StringPtrInput `pulumi:"authority"`
	SmartProxyEnabled pulumi.BoolPtrInput   `pulumi:"smartProxyEnabled"`
}

func (ServiceAuthenticationConfigurationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfigurationInfoResponse)(nil)).Elem()
}

func (i ServiceAuthenticationConfigurationInfoResponseArgs) ToServiceAuthenticationConfigurationInfoResponseOutput() ServiceAuthenticationConfigurationInfoResponseOutput {
	return i.ToServiceAuthenticationConfigurationInfoResponseOutputWithContext(context.Background())
}

func (i ServiceAuthenticationConfigurationInfoResponseArgs) ToServiceAuthenticationConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoResponseOutput)
}

func (i ServiceAuthenticationConfigurationInfoResponseArgs) ToServiceAuthenticationConfigurationInfoResponsePtrOutput() ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return i.ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (i ServiceAuthenticationConfigurationInfoResponseArgs) ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoResponseOutput).ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(ctx)
}









type ServiceAuthenticationConfigurationInfoResponsePtrInput interface {
	pulumi.Input

	ToServiceAuthenticationConfigurationInfoResponsePtrOutput() ServiceAuthenticationConfigurationInfoResponsePtrOutput
	ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(context.Context) ServiceAuthenticationConfigurationInfoResponsePtrOutput
}

type serviceAuthenticationConfigurationInfoResponsePtrType ServiceAuthenticationConfigurationInfoResponseArgs

func ServiceAuthenticationConfigurationInfoResponsePtr(v *ServiceAuthenticationConfigurationInfoResponseArgs) ServiceAuthenticationConfigurationInfoResponsePtrInput {
	return (*serviceAuthenticationConfigurationInfoResponsePtrType)(v)
}

func (*serviceAuthenticationConfigurationInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfigurationInfoResponse)(nil)).Elem()
}

func (i *serviceAuthenticationConfigurationInfoResponsePtrType) ToServiceAuthenticationConfigurationInfoResponsePtrOutput() ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return i.ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (i *serviceAuthenticationConfigurationInfoResponsePtrType) ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceAuthenticationConfigurationInfoResponsePtrOutput)
}

type ServiceAuthenticationConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceAuthenticationConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) ToServiceAuthenticationConfigurationInfoResponseOutput() ServiceAuthenticationConfigurationInfoResponseOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) ToServiceAuthenticationConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoResponseOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) ToServiceAuthenticationConfigurationInfoResponsePtrOutput() ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o.ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceAuthenticationConfigurationInfoResponse) *ServiceAuthenticationConfigurationInfoResponse {
		return &v
	}).(ServiceAuthenticationConfigurationInfoResponsePtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfoResponse) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfoResponse) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponseOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceAuthenticationConfigurationInfoResponse) *bool { return v.SmartProxyEnabled }).(pulumi.BoolPtrOutput)
}

type ServiceAuthenticationConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceAuthenticationConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceAuthenticationConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) ToServiceAuthenticationConfigurationInfoResponsePtrOutput() ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) ToServiceAuthenticationConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) Elem() ServiceAuthenticationConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) ServiceAuthenticationConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceAuthenticationConfigurationInfoResponse
		return ret
	}).(ServiceAuthenticationConfigurationInfoResponseOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o ServiceAuthenticationConfigurationInfoResponsePtrOutput) SmartProxyEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceAuthenticationConfigurationInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SmartProxyEnabled
	}).(pulumi.BoolPtrOutput)
}

type ServiceCorsConfigurationInfo struct {
	AllowCredentials *bool    `pulumi:"allowCredentials"`
	Headers          []string `pulumi:"headers"`
	MaxAge           *float64 `pulumi:"maxAge"`
	Methods          []string `pulumi:"methods"`
	Origins          []string `pulumi:"origins"`
}





type ServiceCorsConfigurationInfoInput interface {
	pulumi.Input

	ToServiceCorsConfigurationInfoOutput() ServiceCorsConfigurationInfoOutput
	ToServiceCorsConfigurationInfoOutputWithContext(context.Context) ServiceCorsConfigurationInfoOutput
}

type ServiceCorsConfigurationInfoArgs struct {
	AllowCredentials pulumi.BoolPtrInput     `pulumi:"allowCredentials"`
	Headers          pulumi.StringArrayInput `pulumi:"headers"`
	MaxAge           pulumi.Float64PtrInput  `pulumi:"maxAge"`
	Methods          pulumi.StringArrayInput `pulumi:"methods"`
	Origins          pulumi.StringArrayInput `pulumi:"origins"`
}

func (ServiceCorsConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoOutput() ServiceCorsConfigurationInfoOutput {
	return i.ToServiceCorsConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoOutput)
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return i.ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceCorsConfigurationInfoArgs) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoOutput).ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceCorsConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput
	ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Context) ServiceCorsConfigurationInfoPtrOutput
}

type serviceCorsConfigurationInfoPtrType ServiceCorsConfigurationInfoArgs

func ServiceCorsConfigurationInfoPtr(v *ServiceCorsConfigurationInfoArgs) ServiceCorsConfigurationInfoPtrInput {
	return (*serviceCorsConfigurationInfoPtrType)(v)
}

func (*serviceCorsConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (i *serviceCorsConfigurationInfoPtrType) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return i.ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceCorsConfigurationInfoPtrType) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoPtrOutput)
}

type ServiceCorsConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoOutput() ServiceCorsConfigurationInfoOutput {
	return o
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoOutput {
	return o
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return o.ToServiceCorsConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceCorsConfigurationInfoOutput) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceCorsConfigurationInfo) *ServiceCorsConfigurationInfo {
		return &v
	}).(ServiceCorsConfigurationInfoPtrOutput)
}

func (o ServiceCorsConfigurationInfoOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoOutput) MaxAge() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) *float64 { return v.MaxAge }).(pulumi.Float64PtrOutput)
}

func (o ServiceCorsConfigurationInfoOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) []string { return v.Methods }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfo) []string { return v.Origins }).(pulumi.StringArrayOutput)
}

type ServiceCorsConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfigurationInfo)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoPtrOutput) ToServiceCorsConfigurationInfoPtrOutput() ServiceCorsConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoPtrOutput) ToServiceCorsConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoPtrOutput) Elem() ServiceCorsConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) ServiceCorsConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceCorsConfigurationInfo
		return ret
	}).(ServiceCorsConfigurationInfoOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) *bool {
		if v == nil {
			return nil
		}
		return v.AllowCredentials
	}).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) []string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) MaxAge() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxAge
	}).(pulumi.Float64PtrOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) []string {
		if v == nil {
			return nil
		}
		return v.Methods
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoPtrOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfo) []string {
		if v == nil {
			return nil
		}
		return v.Origins
	}).(pulumi.StringArrayOutput)
}

type ServiceCorsConfigurationInfoResponse struct {
	AllowCredentials *bool    `pulumi:"allowCredentials"`
	Headers          []string `pulumi:"headers"`
	MaxAge           *float64 `pulumi:"maxAge"`
	Methods          []string `pulumi:"methods"`
	Origins          []string `pulumi:"origins"`
}





type ServiceCorsConfigurationInfoResponseInput interface {
	pulumi.Input

	ToServiceCorsConfigurationInfoResponseOutput() ServiceCorsConfigurationInfoResponseOutput
	ToServiceCorsConfigurationInfoResponseOutputWithContext(context.Context) ServiceCorsConfigurationInfoResponseOutput
}

type ServiceCorsConfigurationInfoResponseArgs struct {
	AllowCredentials pulumi.BoolPtrInput     `pulumi:"allowCredentials"`
	Headers          pulumi.StringArrayInput `pulumi:"headers"`
	MaxAge           pulumi.Float64PtrInput  `pulumi:"maxAge"`
	Methods          pulumi.StringArrayInput `pulumi:"methods"`
	Origins          pulumi.StringArrayInput `pulumi:"origins"`
}

func (ServiceCorsConfigurationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfigurationInfoResponse)(nil)).Elem()
}

func (i ServiceCorsConfigurationInfoResponseArgs) ToServiceCorsConfigurationInfoResponseOutput() ServiceCorsConfigurationInfoResponseOutput {
	return i.ToServiceCorsConfigurationInfoResponseOutputWithContext(context.Background())
}

func (i ServiceCorsConfigurationInfoResponseArgs) ToServiceCorsConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoResponseOutput)
}

func (i ServiceCorsConfigurationInfoResponseArgs) ToServiceCorsConfigurationInfoResponsePtrOutput() ServiceCorsConfigurationInfoResponsePtrOutput {
	return i.ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (i ServiceCorsConfigurationInfoResponseArgs) ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoResponseOutput).ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(ctx)
}









type ServiceCorsConfigurationInfoResponsePtrInput interface {
	pulumi.Input

	ToServiceCorsConfigurationInfoResponsePtrOutput() ServiceCorsConfigurationInfoResponsePtrOutput
	ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(context.Context) ServiceCorsConfigurationInfoResponsePtrOutput
}

type serviceCorsConfigurationInfoResponsePtrType ServiceCorsConfigurationInfoResponseArgs

func ServiceCorsConfigurationInfoResponsePtr(v *ServiceCorsConfigurationInfoResponseArgs) ServiceCorsConfigurationInfoResponsePtrInput {
	return (*serviceCorsConfigurationInfoResponsePtrType)(v)
}

func (*serviceCorsConfigurationInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfigurationInfoResponse)(nil)).Elem()
}

func (i *serviceCorsConfigurationInfoResponsePtrType) ToServiceCorsConfigurationInfoResponsePtrOutput() ServiceCorsConfigurationInfoResponsePtrOutput {
	return i.ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (i *serviceCorsConfigurationInfoResponsePtrType) ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCorsConfigurationInfoResponsePtrOutput)
}

type ServiceCorsConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCorsConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoResponseOutput) ToServiceCorsConfigurationInfoResponseOutput() ServiceCorsConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponseOutput) ToServiceCorsConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponseOutput) ToServiceCorsConfigurationInfoResponsePtrOutput() ServiceCorsConfigurationInfoResponsePtrOutput {
	return o.ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (o ServiceCorsConfigurationInfoResponseOutput) ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceCorsConfigurationInfoResponse) *ServiceCorsConfigurationInfoResponse {
		return &v
	}).(ServiceCorsConfigurationInfoResponsePtrOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) *bool { return v.AllowCredentials }).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) []string { return v.Headers }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) MaxAge() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) *float64 { return v.MaxAge }).(pulumi.Float64PtrOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) []string { return v.Methods }).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponseOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServiceCorsConfigurationInfoResponse) []string { return v.Origins }).(pulumi.StringArrayOutput)
}

type ServiceCorsConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceCorsConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCorsConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) ToServiceCorsConfigurationInfoResponsePtrOutput() ServiceCorsConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) ToServiceCorsConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCorsConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Elem() ServiceCorsConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) ServiceCorsConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceCorsConfigurationInfoResponse
		return ret
	}).(ServiceCorsConfigurationInfoResponseOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) AllowCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowCredentials
	}).(pulumi.BoolPtrOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Headers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) MaxAge() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxAge
	}).(pulumi.Float64PtrOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Methods
	}).(pulumi.StringArrayOutput)
}

func (o ServiceCorsConfigurationInfoResponsePtrOutput) Origins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceCorsConfigurationInfoResponse) []string {
		if v == nil {
			return nil
		}
		return v.Origins
	}).(pulumi.StringArrayOutput)
}

type ServiceCosmosDbConfigurationInfo struct {
	KeyVaultKeyUri  *string  `pulumi:"keyVaultKeyUri"`
	OfferThroughput *float64 `pulumi:"offerThroughput"`
}





type ServiceCosmosDbConfigurationInfoInput interface {
	pulumi.Input

	ToServiceCosmosDbConfigurationInfoOutput() ServiceCosmosDbConfigurationInfoOutput
	ToServiceCosmosDbConfigurationInfoOutputWithContext(context.Context) ServiceCosmosDbConfigurationInfoOutput
}

type ServiceCosmosDbConfigurationInfoArgs struct {
	KeyVaultKeyUri  pulumi.StringPtrInput  `pulumi:"keyVaultKeyUri"`
	OfferThroughput pulumi.Float64PtrInput `pulumi:"offerThroughput"`
}

func (ServiceCosmosDbConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoOutput() ServiceCosmosDbConfigurationInfoOutput {
	return i.ToServiceCosmosDbConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoOutput)
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return i.ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceCosmosDbConfigurationInfoArgs) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoOutput).ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceCosmosDbConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput
	ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Context) ServiceCosmosDbConfigurationInfoPtrOutput
}

type serviceCosmosDbConfigurationInfoPtrType ServiceCosmosDbConfigurationInfoArgs

func ServiceCosmosDbConfigurationInfoPtr(v *ServiceCosmosDbConfigurationInfoArgs) ServiceCosmosDbConfigurationInfoPtrInput {
	return (*serviceCosmosDbConfigurationInfoPtrType)(v)
}

func (*serviceCosmosDbConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (i *serviceCosmosDbConfigurationInfoPtrType) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return i.ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceCosmosDbConfigurationInfoPtrType) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

type ServiceCosmosDbConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoOutput() ServiceCosmosDbConfigurationInfoOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceCosmosDbConfigurationInfoOutput) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceCosmosDbConfigurationInfo) *ServiceCosmosDbConfigurationInfo {
		return &v
	}).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfo) *string { return v.KeyVaultKeyUri }).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoOutput) OfferThroughput() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfo) *float64 { return v.OfferThroughput }).(pulumi.Float64PtrOutput)
}

type ServiceCosmosDbConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCosmosDbConfigurationInfo)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) ToServiceCosmosDbConfigurationInfoPtrOutput() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) ToServiceCosmosDbConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoPtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) Elem() ServiceCosmosDbConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfo) ServiceCosmosDbConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceCosmosDbConfigurationInfo
		return ret
	}).(ServiceCosmosDbConfigurationInfoOutput)
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultKeyUri
	}).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoPtrOutput) OfferThroughput() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfo) *float64 {
		if v == nil {
			return nil
		}
		return v.OfferThroughput
	}).(pulumi.Float64PtrOutput)
}

type ServiceCosmosDbConfigurationInfoResponse struct {
	KeyVaultKeyUri  *string  `pulumi:"keyVaultKeyUri"`
	OfferThroughput *float64 `pulumi:"offerThroughput"`
}





type ServiceCosmosDbConfigurationInfoResponseInput interface {
	pulumi.Input

	ToServiceCosmosDbConfigurationInfoResponseOutput() ServiceCosmosDbConfigurationInfoResponseOutput
	ToServiceCosmosDbConfigurationInfoResponseOutputWithContext(context.Context) ServiceCosmosDbConfigurationInfoResponseOutput
}

type ServiceCosmosDbConfigurationInfoResponseArgs struct {
	KeyVaultKeyUri  pulumi.StringPtrInput  `pulumi:"keyVaultKeyUri"`
	OfferThroughput pulumi.Float64PtrInput `pulumi:"offerThroughput"`
}

func (ServiceCosmosDbConfigurationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCosmosDbConfigurationInfoResponse)(nil)).Elem()
}

func (i ServiceCosmosDbConfigurationInfoResponseArgs) ToServiceCosmosDbConfigurationInfoResponseOutput() ServiceCosmosDbConfigurationInfoResponseOutput {
	return i.ToServiceCosmosDbConfigurationInfoResponseOutputWithContext(context.Background())
}

func (i ServiceCosmosDbConfigurationInfoResponseArgs) ToServiceCosmosDbConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoResponseOutput)
}

func (i ServiceCosmosDbConfigurationInfoResponseArgs) ToServiceCosmosDbConfigurationInfoResponsePtrOutput() ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return i.ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (i ServiceCosmosDbConfigurationInfoResponseArgs) ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoResponseOutput).ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(ctx)
}









type ServiceCosmosDbConfigurationInfoResponsePtrInput interface {
	pulumi.Input

	ToServiceCosmosDbConfigurationInfoResponsePtrOutput() ServiceCosmosDbConfigurationInfoResponsePtrOutput
	ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(context.Context) ServiceCosmosDbConfigurationInfoResponsePtrOutput
}

type serviceCosmosDbConfigurationInfoResponsePtrType ServiceCosmosDbConfigurationInfoResponseArgs

func ServiceCosmosDbConfigurationInfoResponsePtr(v *ServiceCosmosDbConfigurationInfoResponseArgs) ServiceCosmosDbConfigurationInfoResponsePtrInput {
	return (*serviceCosmosDbConfigurationInfoResponsePtrType)(v)
}

func (*serviceCosmosDbConfigurationInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCosmosDbConfigurationInfoResponse)(nil)).Elem()
}

func (i *serviceCosmosDbConfigurationInfoResponsePtrType) ToServiceCosmosDbConfigurationInfoResponsePtrOutput() ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return i.ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (i *serviceCosmosDbConfigurationInfoResponsePtrType) ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceCosmosDbConfigurationInfoResponsePtrOutput)
}

type ServiceCosmosDbConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceCosmosDbConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) ToServiceCosmosDbConfigurationInfoResponseOutput() ServiceCosmosDbConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) ToServiceCosmosDbConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoResponseOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) ToServiceCosmosDbConfigurationInfoResponsePtrOutput() ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o.ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceCosmosDbConfigurationInfoResponse) *ServiceCosmosDbConfigurationInfoResponse {
		return &v
	}).(ServiceCosmosDbConfigurationInfoResponsePtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfoResponse) *string { return v.KeyVaultKeyUri }).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoResponseOutput) OfferThroughput() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ServiceCosmosDbConfigurationInfoResponse) *float64 { return v.OfferThroughput }).(pulumi.Float64PtrOutput)
}

type ServiceCosmosDbConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceCosmosDbConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceCosmosDbConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) ToServiceCosmosDbConfigurationInfoResponsePtrOutput() ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) ToServiceCosmosDbConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) Elem() ServiceCosmosDbConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfoResponse) ServiceCosmosDbConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceCosmosDbConfigurationInfoResponse
		return ret
	}).(ServiceCosmosDbConfigurationInfoResponseOutput)
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) KeyVaultKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultKeyUri
	}).(pulumi.StringPtrOutput)
}

func (o ServiceCosmosDbConfigurationInfoResponsePtrOutput) OfferThroughput() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServiceCosmosDbConfigurationInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.OfferThroughput
	}).(pulumi.Float64PtrOutput)
}

type ServiceExportConfigurationInfo struct {
	StorageAccountName *string `pulumi:"storageAccountName"`
}





type ServiceExportConfigurationInfoInput interface {
	pulumi.Input

	ToServiceExportConfigurationInfoOutput() ServiceExportConfigurationInfoOutput
	ToServiceExportConfigurationInfoOutputWithContext(context.Context) ServiceExportConfigurationInfoOutput
}

type ServiceExportConfigurationInfoArgs struct {
	StorageAccountName pulumi.StringPtrInput `pulumi:"storageAccountName"`
}

func (ServiceExportConfigurationInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceExportConfigurationInfo)(nil)).Elem()
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoOutput() ServiceExportConfigurationInfoOutput {
	return i.ToServiceExportConfigurationInfoOutputWithContext(context.Background())
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoOutput)
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return i.ToServiceExportConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i ServiceExportConfigurationInfoArgs) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoOutput).ToServiceExportConfigurationInfoPtrOutputWithContext(ctx)
}









type ServiceExportConfigurationInfoPtrInput interface {
	pulumi.Input

	ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput
	ToServiceExportConfigurationInfoPtrOutputWithContext(context.Context) ServiceExportConfigurationInfoPtrOutput
}

type serviceExportConfigurationInfoPtrType ServiceExportConfigurationInfoArgs

func ServiceExportConfigurationInfoPtr(v *ServiceExportConfigurationInfoArgs) ServiceExportConfigurationInfoPtrInput {
	return (*serviceExportConfigurationInfoPtrType)(v)
}

func (*serviceExportConfigurationInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceExportConfigurationInfo)(nil)).Elem()
}

func (i *serviceExportConfigurationInfoPtrType) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return i.ToServiceExportConfigurationInfoPtrOutputWithContext(context.Background())
}

func (i *serviceExportConfigurationInfoPtrType) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoPtrOutput)
}

type ServiceExportConfigurationInfoOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceExportConfigurationInfo)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoOutput() ServiceExportConfigurationInfoOutput {
	return o
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoOutput {
	return o
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return o.ToServiceExportConfigurationInfoPtrOutputWithContext(context.Background())
}

func (o ServiceExportConfigurationInfoOutput) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceExportConfigurationInfo) *ServiceExportConfigurationInfo {
		return &v
	}).(ServiceExportConfigurationInfoPtrOutput)
}

func (o ServiceExportConfigurationInfoOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceExportConfigurationInfo) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

type ServiceExportConfigurationInfoPtrOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceExportConfigurationInfo)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoPtrOutput) ToServiceExportConfigurationInfoPtrOutput() ServiceExportConfigurationInfoPtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoPtrOutput) ToServiceExportConfigurationInfoPtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoPtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoPtrOutput) Elem() ServiceExportConfigurationInfoOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfo) ServiceExportConfigurationInfo {
		if v != nil {
			return *v
		}
		var ret ServiceExportConfigurationInfo
		return ret
	}).(ServiceExportConfigurationInfoOutput)
}

func (o ServiceExportConfigurationInfoPtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfo) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type ServiceExportConfigurationInfoResponse struct {
	StorageAccountName *string `pulumi:"storageAccountName"`
}





type ServiceExportConfigurationInfoResponseInput interface {
	pulumi.Input

	ToServiceExportConfigurationInfoResponseOutput() ServiceExportConfigurationInfoResponseOutput
	ToServiceExportConfigurationInfoResponseOutputWithContext(context.Context) ServiceExportConfigurationInfoResponseOutput
}

type ServiceExportConfigurationInfoResponseArgs struct {
	StorageAccountName pulumi.StringPtrInput `pulumi:"storageAccountName"`
}

func (ServiceExportConfigurationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceExportConfigurationInfoResponse)(nil)).Elem()
}

func (i ServiceExportConfigurationInfoResponseArgs) ToServiceExportConfigurationInfoResponseOutput() ServiceExportConfigurationInfoResponseOutput {
	return i.ToServiceExportConfigurationInfoResponseOutputWithContext(context.Background())
}

func (i ServiceExportConfigurationInfoResponseArgs) ToServiceExportConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoResponseOutput)
}

func (i ServiceExportConfigurationInfoResponseArgs) ToServiceExportConfigurationInfoResponsePtrOutput() ServiceExportConfigurationInfoResponsePtrOutput {
	return i.ToServiceExportConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (i ServiceExportConfigurationInfoResponseArgs) ToServiceExportConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoResponseOutput).ToServiceExportConfigurationInfoResponsePtrOutputWithContext(ctx)
}









type ServiceExportConfigurationInfoResponsePtrInput interface {
	pulumi.Input

	ToServiceExportConfigurationInfoResponsePtrOutput() ServiceExportConfigurationInfoResponsePtrOutput
	ToServiceExportConfigurationInfoResponsePtrOutputWithContext(context.Context) ServiceExportConfigurationInfoResponsePtrOutput
}

type serviceExportConfigurationInfoResponsePtrType ServiceExportConfigurationInfoResponseArgs

func ServiceExportConfigurationInfoResponsePtr(v *ServiceExportConfigurationInfoResponseArgs) ServiceExportConfigurationInfoResponsePtrInput {
	return (*serviceExportConfigurationInfoResponsePtrType)(v)
}

func (*serviceExportConfigurationInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceExportConfigurationInfoResponse)(nil)).Elem()
}

func (i *serviceExportConfigurationInfoResponsePtrType) ToServiceExportConfigurationInfoResponsePtrOutput() ServiceExportConfigurationInfoResponsePtrOutput {
	return i.ToServiceExportConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (i *serviceExportConfigurationInfoResponsePtrType) ToServiceExportConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceExportConfigurationInfoResponsePtrOutput)
}

type ServiceExportConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceExportConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoResponseOutput) ToServiceExportConfigurationInfoResponseOutput() ServiceExportConfigurationInfoResponseOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponseOutput) ToServiceExportConfigurationInfoResponseOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoResponseOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponseOutput) ToServiceExportConfigurationInfoResponsePtrOutput() ServiceExportConfigurationInfoResponsePtrOutput {
	return o.ToServiceExportConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (o ServiceExportConfigurationInfoResponseOutput) ToServiceExportConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceExportConfigurationInfoResponse) *ServiceExportConfigurationInfoResponse {
		return &v
	}).(ServiceExportConfigurationInfoResponsePtrOutput)
}

func (o ServiceExportConfigurationInfoResponseOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceExportConfigurationInfoResponse) *string { return v.StorageAccountName }).(pulumi.StringPtrOutput)
}

type ServiceExportConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceExportConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceExportConfigurationInfoResponse)(nil)).Elem()
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) ToServiceExportConfigurationInfoResponsePtrOutput() ServiceExportConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) ToServiceExportConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ServiceExportConfigurationInfoResponsePtrOutput {
	return o
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) Elem() ServiceExportConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfoResponse) ServiceExportConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ServiceExportConfigurationInfoResponse
		return ret
	}).(ServiceExportConfigurationInfoResponseOutput)
}

func (o ServiceExportConfigurationInfoResponsePtrOutput) StorageAccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceExportConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountName
	}).(pulumi.StringPtrOutput)
}

type ServicesProperties struct {
	AccessPolicies              []ServiceAccessPolicyEntry              `pulumi:"accessPolicies"`
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfo `pulumi:"authenticationConfiguration"`
	CorsConfiguration           *ServiceCorsConfigurationInfo           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       *ServiceCosmosDbConfigurationInfo       `pulumi:"cosmosDbConfiguration"`
	ExportConfiguration         *ServiceExportConfigurationInfo         `pulumi:"exportConfiguration"`
	PrivateEndpointConnections  []PrivateEndpointConnection             `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess         *string                                 `pulumi:"publicNetworkAccess"`
}





type ServicesPropertiesInput interface {
	pulumi.Input

	ToServicesPropertiesOutput() ServicesPropertiesOutput
	ToServicesPropertiesOutputWithContext(context.Context) ServicesPropertiesOutput
}

type ServicesPropertiesArgs struct {
	AccessPolicies              ServiceAccessPolicyEntryArrayInput             `pulumi:"accessPolicies"`
	AuthenticationConfiguration ServiceAuthenticationConfigurationInfoPtrInput `pulumi:"authenticationConfiguration"`
	CorsConfiguration           ServiceCorsConfigurationInfoPtrInput           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       ServiceCosmosDbConfigurationInfoPtrInput       `pulumi:"cosmosDbConfiguration"`
	ExportConfiguration         ServiceExportConfigurationInfoPtrInput         `pulumi:"exportConfiguration"`
	PrivateEndpointConnections  PrivateEndpointConnectionArrayInput            `pulumi:"privateEndpointConnections"`
	PublicNetworkAccess         pulumi.StringPtrInput                          `pulumi:"publicNetworkAccess"`
}

func (ServicesPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesProperties)(nil)).Elem()
}

func (i ServicesPropertiesArgs) ToServicesPropertiesOutput() ServicesPropertiesOutput {
	return i.ToServicesPropertiesOutputWithContext(context.Background())
}

func (i ServicesPropertiesArgs) ToServicesPropertiesOutputWithContext(ctx context.Context) ServicesPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesOutput)
}

func (i ServicesPropertiesArgs) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return i.ToServicesPropertiesPtrOutputWithContext(context.Background())
}

func (i ServicesPropertiesArgs) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesOutput).ToServicesPropertiesPtrOutputWithContext(ctx)
}









type ServicesPropertiesPtrInput interface {
	pulumi.Input

	ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput
	ToServicesPropertiesPtrOutputWithContext(context.Context) ServicesPropertiesPtrOutput
}

type servicesPropertiesPtrType ServicesPropertiesArgs

func ServicesPropertiesPtr(v *ServicesPropertiesArgs) ServicesPropertiesPtrInput {
	return (*servicesPropertiesPtrType)(v)
}

func (*servicesPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesProperties)(nil)).Elem()
}

func (i *servicesPropertiesPtrType) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return i.ToServicesPropertiesPtrOutputWithContext(context.Background())
}

func (i *servicesPropertiesPtrType) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesPtrOutput)
}

type ServicesPropertiesOutput struct{ *pulumi.OutputState }

func (ServicesPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesProperties)(nil)).Elem()
}

func (o ServicesPropertiesOutput) ToServicesPropertiesOutput() ServicesPropertiesOutput {
	return o
}

func (o ServicesPropertiesOutput) ToServicesPropertiesOutputWithContext(ctx context.Context) ServicesPropertiesOutput {
	return o
}

func (o ServicesPropertiesOutput) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return o.ToServicesPropertiesPtrOutputWithContext(context.Background())
}

func (o ServicesPropertiesOutput) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicesProperties) *ServicesProperties {
		return &v
	}).(ServicesPropertiesPtrOutput)
}

func (o ServicesPropertiesOutput) AccessPolicies() ServiceAccessPolicyEntryArrayOutput {
	return o.ApplyT(func(v ServicesProperties) []ServiceAccessPolicyEntry { return v.AccessPolicies }).(ServiceAccessPolicyEntryArrayOutput)
}

func (o ServicesPropertiesOutput) AuthenticationConfiguration() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceAuthenticationConfigurationInfo {
		return v.AuthenticationConfiguration
	}).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) CorsConfiguration() ServiceCorsConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceCorsConfigurationInfo { return v.CorsConfiguration }).(ServiceCorsConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) CosmosDbConfiguration() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceCosmosDbConfigurationInfo { return v.CosmosDbConfiguration }).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) ExportConfiguration() ServiceExportConfigurationInfoPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *ServiceExportConfigurationInfo { return v.ExportConfiguration }).(ServiceExportConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesOutput) PrivateEndpointConnections() PrivateEndpointConnectionArrayOutput {
	return o.ApplyT(func(v ServicesProperties) []PrivateEndpointConnection { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionArrayOutput)
}

func (o ServicesPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type ServicesPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ServicesPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesProperties)(nil)).Elem()
}

func (o ServicesPropertiesPtrOutput) ToServicesPropertiesPtrOutput() ServicesPropertiesPtrOutput {
	return o
}

func (o ServicesPropertiesPtrOutput) ToServicesPropertiesPtrOutputWithContext(ctx context.Context) ServicesPropertiesPtrOutput {
	return o
}

func (o ServicesPropertiesPtrOutput) Elem() ServicesPropertiesOutput {
	return o.ApplyT(func(v *ServicesProperties) ServicesProperties {
		if v != nil {
			return *v
		}
		var ret ServicesProperties
		return ret
	}).(ServicesPropertiesOutput)
}

func (o ServicesPropertiesPtrOutput) AccessPolicies() ServiceAccessPolicyEntryArrayOutput {
	return o.ApplyT(func(v *ServicesProperties) []ServiceAccessPolicyEntry {
		if v == nil {
			return nil
		}
		return v.AccessPolicies
	}).(ServiceAccessPolicyEntryArrayOutput)
}

func (o ServicesPropertiesPtrOutput) AuthenticationConfiguration() ServiceAuthenticationConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceAuthenticationConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.AuthenticationConfiguration
	}).(ServiceAuthenticationConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) CorsConfiguration() ServiceCorsConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceCorsConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.CorsConfiguration
	}).(ServiceCorsConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) CosmosDbConfiguration() ServiceCosmosDbConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceCosmosDbConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.CosmosDbConfiguration
	}).(ServiceCosmosDbConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) ExportConfiguration() ServiceExportConfigurationInfoPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *ServiceExportConfigurationInfo {
		if v == nil {
			return nil
		}
		return v.ExportConfiguration
	}).(ServiceExportConfigurationInfoPtrOutput)
}

func (o ServicesPropertiesPtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionArrayOutput {
	return o.ApplyT(func(v *ServicesProperties) []PrivateEndpointConnection {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionArrayOutput)
}

func (o ServicesPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type ServicesPropertiesResponse struct {
	AccessPolicies              []ServiceAccessPolicyEntryResponse              `pulumi:"accessPolicies"`
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfoResponse `pulumi:"authenticationConfiguration"`
	CorsConfiguration           *ServiceCorsConfigurationInfoResponse           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       *ServiceCosmosDbConfigurationInfoResponse       `pulumi:"cosmosDbConfiguration"`
	ExportConfiguration         *ServiceExportConfigurationInfoResponse         `pulumi:"exportConfiguration"`
	PrivateEndpointConnections  []PrivateEndpointConnectionResponse             `pulumi:"privateEndpointConnections"`
	ProvisioningState           string                                          `pulumi:"provisioningState"`
	PublicNetworkAccess         *string                                         `pulumi:"publicNetworkAccess"`
}





type ServicesPropertiesResponseInput interface {
	pulumi.Input

	ToServicesPropertiesResponseOutput() ServicesPropertiesResponseOutput
	ToServicesPropertiesResponseOutputWithContext(context.Context) ServicesPropertiesResponseOutput
}

type ServicesPropertiesResponseArgs struct {
	AccessPolicies              ServiceAccessPolicyEntryResponseArrayInput             `pulumi:"accessPolicies"`
	AuthenticationConfiguration ServiceAuthenticationConfigurationInfoResponsePtrInput `pulumi:"authenticationConfiguration"`
	CorsConfiguration           ServiceCorsConfigurationInfoResponsePtrInput           `pulumi:"corsConfiguration"`
	CosmosDbConfiguration       ServiceCosmosDbConfigurationInfoResponsePtrInput       `pulumi:"cosmosDbConfiguration"`
	ExportConfiguration         ServiceExportConfigurationInfoResponsePtrInput         `pulumi:"exportConfiguration"`
	PrivateEndpointConnections  PrivateEndpointConnectionResponseArrayInput            `pulumi:"privateEndpointConnections"`
	ProvisioningState           pulumi.StringInput                                     `pulumi:"provisioningState"`
	PublicNetworkAccess         pulumi.StringPtrInput                                  `pulumi:"publicNetworkAccess"`
}

func (ServicesPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesPropertiesResponse)(nil)).Elem()
}

func (i ServicesPropertiesResponseArgs) ToServicesPropertiesResponseOutput() ServicesPropertiesResponseOutput {
	return i.ToServicesPropertiesResponseOutputWithContext(context.Background())
}

func (i ServicesPropertiesResponseArgs) ToServicesPropertiesResponseOutputWithContext(ctx context.Context) ServicesPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesResponseOutput)
}

func (i ServicesPropertiesResponseArgs) ToServicesPropertiesResponsePtrOutput() ServicesPropertiesResponsePtrOutput {
	return i.ToServicesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ServicesPropertiesResponseArgs) ToServicesPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicesPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesResponseOutput).ToServicesPropertiesResponsePtrOutputWithContext(ctx)
}









type ServicesPropertiesResponsePtrInput interface {
	pulumi.Input

	ToServicesPropertiesResponsePtrOutput() ServicesPropertiesResponsePtrOutput
	ToServicesPropertiesResponsePtrOutputWithContext(context.Context) ServicesPropertiesResponsePtrOutput
}

type servicesPropertiesResponsePtrType ServicesPropertiesResponseArgs

func ServicesPropertiesResponsePtr(v *ServicesPropertiesResponseArgs) ServicesPropertiesResponsePtrInput {
	return (*servicesPropertiesResponsePtrType)(v)
}

func (*servicesPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesPropertiesResponse)(nil)).Elem()
}

func (i *servicesPropertiesResponsePtrType) ToServicesPropertiesResponsePtrOutput() ServicesPropertiesResponsePtrOutput {
	return i.ToServicesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *servicesPropertiesResponsePtrType) ToServicesPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicesPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesPropertiesResponsePtrOutput)
}

type ServicesPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServicesPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesPropertiesResponse)(nil)).Elem()
}

func (o ServicesPropertiesResponseOutput) ToServicesPropertiesResponseOutput() ServicesPropertiesResponseOutput {
	return o
}

func (o ServicesPropertiesResponseOutput) ToServicesPropertiesResponseOutputWithContext(ctx context.Context) ServicesPropertiesResponseOutput {
	return o
}

func (o ServicesPropertiesResponseOutput) ToServicesPropertiesResponsePtrOutput() ServicesPropertiesResponsePtrOutput {
	return o.ToServicesPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ServicesPropertiesResponseOutput) ToServicesPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicesPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicesPropertiesResponse) *ServicesPropertiesResponse {
		return &v
	}).(ServicesPropertiesResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) AccessPolicies() ServiceAccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) []ServiceAccessPolicyEntryResponse { return v.AccessPolicies }).(ServiceAccessPolicyEntryResponseArrayOutput)
}

func (o ServicesPropertiesResponseOutput) AuthenticationConfiguration() ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceAuthenticationConfigurationInfoResponse {
		return v.AuthenticationConfiguration
	}).(ServiceAuthenticationConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) CorsConfiguration() ServiceCorsConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceCorsConfigurationInfoResponse { return v.CorsConfiguration }).(ServiceCorsConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) CosmosDbConfiguration() ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceCosmosDbConfigurationInfoResponse {
		return v.CosmosDbConfiguration
	}).(ServiceCosmosDbConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) ExportConfiguration() ServiceExportConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *ServiceExportConfigurationInfoResponse {
		return v.ExportConfiguration
	}).(ServiceExportConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o ServicesPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServicesPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type ServicesPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ServicesPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesPropertiesResponse)(nil)).Elem()
}

func (o ServicesPropertiesResponsePtrOutput) ToServicesPropertiesResponsePtrOutput() ServicesPropertiesResponsePtrOutput {
	return o
}

func (o ServicesPropertiesResponsePtrOutput) ToServicesPropertiesResponsePtrOutputWithContext(ctx context.Context) ServicesPropertiesResponsePtrOutput {
	return o
}

func (o ServicesPropertiesResponsePtrOutput) Elem() ServicesPropertiesResponseOutput {
	return o.ApplyT(func(v *ServicesPropertiesResponse) ServicesPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ServicesPropertiesResponse
		return ret
	}).(ServicesPropertiesResponseOutput)
}

func (o ServicesPropertiesResponsePtrOutput) AccessPolicies() ServiceAccessPolicyEntryResponseArrayOutput {
	return o.ApplyT(func(v *ServicesPropertiesResponse) []ServiceAccessPolicyEntryResponse {
		if v == nil {
			return nil
		}
		return v.AccessPolicies
	}).(ServiceAccessPolicyEntryResponseArrayOutput)
}

func (o ServicesPropertiesResponsePtrOutput) AuthenticationConfiguration() ServiceAuthenticationConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v *ServicesPropertiesResponse) *ServiceAuthenticationConfigurationInfoResponse {
		if v == nil {
			return nil
		}
		return v.AuthenticationConfiguration
	}).(ServiceAuthenticationConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponsePtrOutput) CorsConfiguration() ServiceCorsConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v *ServicesPropertiesResponse) *ServiceCorsConfigurationInfoResponse {
		if v == nil {
			return nil
		}
		return v.CorsConfiguration
	}).(ServiceCorsConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponsePtrOutput) CosmosDbConfiguration() ServiceCosmosDbConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v *ServicesPropertiesResponse) *ServiceCosmosDbConfigurationInfoResponse {
		if v == nil {
			return nil
		}
		return v.CosmosDbConfiguration
	}).(ServiceCosmosDbConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponsePtrOutput) ExportConfiguration() ServiceExportConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v *ServicesPropertiesResponse) *ServiceExportConfigurationInfoResponse {
		if v == nil {
			return nil
		}
		return v.ExportConfiguration
	}).(ServiceExportConfigurationInfoResponsePtrOutput)
}

func (o ServicesPropertiesResponsePtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *ServicesPropertiesResponse) []PrivateEndpointConnectionResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o ServicesPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ServicesPropertiesResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type ServicesResourceIdentity struct {
	Type *string `pulumi:"type"`
}





type ServicesResourceIdentityInput interface {
	pulumi.Input

	ToServicesResourceIdentityOutput() ServicesResourceIdentityOutput
	ToServicesResourceIdentityOutputWithContext(context.Context) ServicesResourceIdentityOutput
}

type ServicesResourceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ServicesResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesResourceIdentity)(nil)).Elem()
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityOutput() ServicesResourceIdentityOutput {
	return i.ToServicesResourceIdentityOutputWithContext(context.Background())
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityOutputWithContext(ctx context.Context) ServicesResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceIdentityOutput)
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return i.ToServicesResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ServicesResourceIdentityArgs) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceIdentityOutput).ToServicesResourceIdentityPtrOutputWithContext(ctx)
}









type ServicesResourceIdentityPtrInput interface {
	pulumi.Input

	ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput
	ToServicesResourceIdentityPtrOutputWithContext(context.Context) ServicesResourceIdentityPtrOutput
}

type servicesResourceIdentityPtrType ServicesResourceIdentityArgs

func ServicesResourceIdentityPtr(v *ServicesResourceIdentityArgs) ServicesResourceIdentityPtrInput {
	return (*servicesResourceIdentityPtrType)(v)
}

func (*servicesResourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesResourceIdentity)(nil)).Elem()
}

func (i *servicesResourceIdentityPtrType) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return i.ToServicesResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *servicesResourceIdentityPtrType) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceIdentityPtrOutput)
}

type ServicesResourceIdentityOutput struct{ *pulumi.OutputState }

func (ServicesResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesResourceIdentity)(nil)).Elem()
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityOutput() ServicesResourceIdentityOutput {
	return o
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityOutputWithContext(ctx context.Context) ServicesResourceIdentityOutput {
	return o
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return o.ToServicesResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ServicesResourceIdentityOutput) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicesResourceIdentity) *ServicesResourceIdentity {
		return &v
	}).(ServicesResourceIdentityPtrOutput)
}

func (o ServicesResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServicesResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ServicesResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesResourceIdentity)(nil)).Elem()
}

func (o ServicesResourceIdentityPtrOutput) ToServicesResourceIdentityPtrOutput() ServicesResourceIdentityPtrOutput {
	return o
}

func (o ServicesResourceIdentityPtrOutput) ToServicesResourceIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceIdentityPtrOutput {
	return o
}

func (o ServicesResourceIdentityPtrOutput) Elem() ServicesResourceIdentityOutput {
	return o.ApplyT(func(v *ServicesResourceIdentity) ServicesResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ServicesResourceIdentity
		return ret
	}).(ServicesResourceIdentityOutput)
}

func (o ServicesResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ServicesResourceResponseIdentity struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ServicesResourceResponseIdentityInput interface {
	pulumi.Input

	ToServicesResourceResponseIdentityOutput() ServicesResourceResponseIdentityOutput
	ToServicesResourceResponseIdentityOutputWithContext(context.Context) ServicesResourceResponseIdentityOutput
}

type ServicesResourceResponseIdentityArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ServicesResourceResponseIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesResourceResponseIdentity)(nil)).Elem()
}

func (i ServicesResourceResponseIdentityArgs) ToServicesResourceResponseIdentityOutput() ServicesResourceResponseIdentityOutput {
	return i.ToServicesResourceResponseIdentityOutputWithContext(context.Background())
}

func (i ServicesResourceResponseIdentityArgs) ToServicesResourceResponseIdentityOutputWithContext(ctx context.Context) ServicesResourceResponseIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceResponseIdentityOutput)
}

func (i ServicesResourceResponseIdentityArgs) ToServicesResourceResponseIdentityPtrOutput() ServicesResourceResponseIdentityPtrOutput {
	return i.ToServicesResourceResponseIdentityPtrOutputWithContext(context.Background())
}

func (i ServicesResourceResponseIdentityArgs) ToServicesResourceResponseIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceResponseIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceResponseIdentityOutput).ToServicesResourceResponseIdentityPtrOutputWithContext(ctx)
}









type ServicesResourceResponseIdentityPtrInput interface {
	pulumi.Input

	ToServicesResourceResponseIdentityPtrOutput() ServicesResourceResponseIdentityPtrOutput
	ToServicesResourceResponseIdentityPtrOutputWithContext(context.Context) ServicesResourceResponseIdentityPtrOutput
}

type servicesResourceResponseIdentityPtrType ServicesResourceResponseIdentityArgs

func ServicesResourceResponseIdentityPtr(v *ServicesResourceResponseIdentityArgs) ServicesResourceResponseIdentityPtrInput {
	return (*servicesResourceResponseIdentityPtrType)(v)
}

func (*servicesResourceResponseIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesResourceResponseIdentity)(nil)).Elem()
}

func (i *servicesResourceResponseIdentityPtrType) ToServicesResourceResponseIdentityPtrOutput() ServicesResourceResponseIdentityPtrOutput {
	return i.ToServicesResourceResponseIdentityPtrOutputWithContext(context.Background())
}

func (i *servicesResourceResponseIdentityPtrType) ToServicesResourceResponseIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceResponseIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServicesResourceResponseIdentityPtrOutput)
}

type ServicesResourceResponseIdentityOutput struct{ *pulumi.OutputState }

func (ServicesResourceResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServicesResourceResponseIdentity)(nil)).Elem()
}

func (o ServicesResourceResponseIdentityOutput) ToServicesResourceResponseIdentityOutput() ServicesResourceResponseIdentityOutput {
	return o
}

func (o ServicesResourceResponseIdentityOutput) ToServicesResourceResponseIdentityOutputWithContext(ctx context.Context) ServicesResourceResponseIdentityOutput {
	return o
}

func (o ServicesResourceResponseIdentityOutput) ToServicesResourceResponseIdentityPtrOutput() ServicesResourceResponseIdentityPtrOutput {
	return o.ToServicesResourceResponseIdentityPtrOutputWithContext(context.Background())
}

func (o ServicesResourceResponseIdentityOutput) ToServicesResourceResponseIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceResponseIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServicesResourceResponseIdentity) *ServicesResourceResponseIdentity {
		return &v
	}).(ServicesResourceResponseIdentityPtrOutput)
}

func (o ServicesResourceResponseIdentityOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicesResourceResponseIdentity) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ServicesResourceResponseIdentityOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ServicesResourceResponseIdentity) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ServicesResourceResponseIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServicesResourceResponseIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ServicesResourceResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (ServicesResourceResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServicesResourceResponseIdentity)(nil)).Elem()
}

func (o ServicesResourceResponseIdentityPtrOutput) ToServicesResourceResponseIdentityPtrOutput() ServicesResourceResponseIdentityPtrOutput {
	return o
}

func (o ServicesResourceResponseIdentityPtrOutput) ToServicesResourceResponseIdentityPtrOutputWithContext(ctx context.Context) ServicesResourceResponseIdentityPtrOutput {
	return o
}

func (o ServicesResourceResponseIdentityPtrOutput) Elem() ServicesResourceResponseIdentityOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) ServicesResourceResponseIdentity {
		if v != nil {
			return *v
		}
		var ret ServicesResourceResponseIdentity
		return ret
	}).(ServicesResourceResponseIdentityOutput)
}

func (o ServicesResourceResponseIdentityPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ServicesResourceResponseIdentityPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ServicesResourceResponseIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServicesResourceResponseIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryArrayOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryResponseOutput{})
	pulumi.RegisterOutputType(ServiceAccessPolicyEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceAuthenticationConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceCorsConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceCosmosDbConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoPtrOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ServiceExportConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServicesPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ServicesResourceIdentityOutput{})
	pulumi.RegisterOutputType(ServicesResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ServicesResourceResponseIdentityOutput{})
	pulumi.RegisterOutputType(ServicesResourceResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
