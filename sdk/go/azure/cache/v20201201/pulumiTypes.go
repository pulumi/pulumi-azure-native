


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionResponse struct {
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
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

type RedisAccessKeysResponse struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}





type RedisAccessKeysResponseInput interface {
	pulumi.Input

	ToRedisAccessKeysResponseOutput() RedisAccessKeysResponseOutput
	ToRedisAccessKeysResponseOutputWithContext(context.Context) RedisAccessKeysResponseOutput
}

type RedisAccessKeysResponseArgs struct {
	PrimaryKey   pulumi.StringInput `pulumi:"primaryKey"`
	SecondaryKey pulumi.StringInput `pulumi:"secondaryKey"`
}

func (RedisAccessKeysResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisAccessKeysResponse)(nil)).Elem()
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponseOutput() RedisAccessKeysResponseOutput {
	return i.ToRedisAccessKeysResponseOutputWithContext(context.Background())
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponseOutputWithContext(ctx context.Context) RedisAccessKeysResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisAccessKeysResponseOutput)
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return i.ToRedisAccessKeysResponsePtrOutputWithContext(context.Background())
}

func (i RedisAccessKeysResponseArgs) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisAccessKeysResponseOutput).ToRedisAccessKeysResponsePtrOutputWithContext(ctx)
}









type RedisAccessKeysResponsePtrInput interface {
	pulumi.Input

	ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput
	ToRedisAccessKeysResponsePtrOutputWithContext(context.Context) RedisAccessKeysResponsePtrOutput
}

type redisAccessKeysResponsePtrType RedisAccessKeysResponseArgs

func RedisAccessKeysResponsePtr(v *RedisAccessKeysResponseArgs) RedisAccessKeysResponsePtrInput {
	return (*redisAccessKeysResponsePtrType)(v)
}

func (*redisAccessKeysResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisAccessKeysResponse)(nil)).Elem()
}

func (i *redisAccessKeysResponsePtrType) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return i.ToRedisAccessKeysResponsePtrOutputWithContext(context.Background())
}

func (i *redisAccessKeysResponsePtrType) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisAccessKeysResponsePtrOutput)
}

type RedisAccessKeysResponseOutput struct{ *pulumi.OutputState }

func (RedisAccessKeysResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisAccessKeysResponse)(nil)).Elem()
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponseOutput() RedisAccessKeysResponseOutput {
	return o
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponseOutputWithContext(ctx context.Context) RedisAccessKeysResponseOutput {
	return o
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return o.ToRedisAccessKeysResponsePtrOutputWithContext(context.Background())
}

func (o RedisAccessKeysResponseOutput) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RedisAccessKeysResponse) *RedisAccessKeysResponse {
		return &v
	}).(RedisAccessKeysResponsePtrOutput)
}

func (o RedisAccessKeysResponseOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o RedisAccessKeysResponseOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v RedisAccessKeysResponse) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

type RedisAccessKeysResponsePtrOutput struct{ *pulumi.OutputState }

func (RedisAccessKeysResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisAccessKeysResponse)(nil)).Elem()
}

func (o RedisAccessKeysResponsePtrOutput) ToRedisAccessKeysResponsePtrOutput() RedisAccessKeysResponsePtrOutput {
	return o
}

func (o RedisAccessKeysResponsePtrOutput) ToRedisAccessKeysResponsePtrOutputWithContext(ctx context.Context) RedisAccessKeysResponsePtrOutput {
	return o
}

func (o RedisAccessKeysResponsePtrOutput) Elem() RedisAccessKeysResponseOutput {
	return o.ApplyT(func(v *RedisAccessKeysResponse) RedisAccessKeysResponse {
		if v != nil {
			return *v
		}
		var ret RedisAccessKeysResponse
		return ret
	}).(RedisAccessKeysResponseOutput)
}

func (o RedisAccessKeysResponsePtrOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisAccessKeysResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o RedisAccessKeysResponsePtrOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisAccessKeysResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SecondaryKey
	}).(pulumi.StringPtrOutput)
}

type RedisCommonPropertiesRedisConfiguration struct {
	AofStorageConnectionString0    *string `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1    *string `pulumi:"aofStorageConnectionString1"`
	MaxfragmentationmemoryReserved *string `pulumi:"maxfragmentationmemoryReserved"`
	MaxmemoryDelta                 *string `pulumi:"maxmemoryDelta"`
	MaxmemoryPolicy                *string `pulumi:"maxmemoryPolicy"`
	MaxmemoryReserved              *string `pulumi:"maxmemoryReserved"`
	RdbBackupEnabled               *string `pulumi:"rdbBackupEnabled"`
	RdbBackupFrequency             *string `pulumi:"rdbBackupFrequency"`
	RdbBackupMaxSnapshotCount      *string `pulumi:"rdbBackupMaxSnapshotCount"`
	RdbStorageConnectionString     *string `pulumi:"rdbStorageConnectionString"`
}





type RedisCommonPropertiesRedisConfigurationInput interface {
	pulumi.Input

	ToRedisCommonPropertiesRedisConfigurationOutput() RedisCommonPropertiesRedisConfigurationOutput
	ToRedisCommonPropertiesRedisConfigurationOutputWithContext(context.Context) RedisCommonPropertiesRedisConfigurationOutput
}

type RedisCommonPropertiesRedisConfigurationArgs struct {
	AofStorageConnectionString0    pulumi.StringPtrInput `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1    pulumi.StringPtrInput `pulumi:"aofStorageConnectionString1"`
	MaxfragmentationmemoryReserved pulumi.StringPtrInput `pulumi:"maxfragmentationmemoryReserved"`
	MaxmemoryDelta                 pulumi.StringPtrInput `pulumi:"maxmemoryDelta"`
	MaxmemoryPolicy                pulumi.StringPtrInput `pulumi:"maxmemoryPolicy"`
	MaxmemoryReserved              pulumi.StringPtrInput `pulumi:"maxmemoryReserved"`
	RdbBackupEnabled               pulumi.StringPtrInput `pulumi:"rdbBackupEnabled"`
	RdbBackupFrequency             pulumi.StringPtrInput `pulumi:"rdbBackupFrequency"`
	RdbBackupMaxSnapshotCount      pulumi.StringPtrInput `pulumi:"rdbBackupMaxSnapshotCount"`
	RdbStorageConnectionString     pulumi.StringPtrInput `pulumi:"rdbStorageConnectionString"`
}

func (RedisCommonPropertiesRedisConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisCommonPropertiesRedisConfiguration)(nil)).Elem()
}

func (i RedisCommonPropertiesRedisConfigurationArgs) ToRedisCommonPropertiesRedisConfigurationOutput() RedisCommonPropertiesRedisConfigurationOutput {
	return i.ToRedisCommonPropertiesRedisConfigurationOutputWithContext(context.Background())
}

func (i RedisCommonPropertiesRedisConfigurationArgs) ToRedisCommonPropertiesRedisConfigurationOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisCommonPropertiesRedisConfigurationOutput)
}

func (i RedisCommonPropertiesRedisConfigurationArgs) ToRedisCommonPropertiesRedisConfigurationPtrOutput() RedisCommonPropertiesRedisConfigurationPtrOutput {
	return i.ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(context.Background())
}

func (i RedisCommonPropertiesRedisConfigurationArgs) ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisCommonPropertiesRedisConfigurationOutput).ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(ctx)
}









type RedisCommonPropertiesRedisConfigurationPtrInput interface {
	pulumi.Input

	ToRedisCommonPropertiesRedisConfigurationPtrOutput() RedisCommonPropertiesRedisConfigurationPtrOutput
	ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(context.Context) RedisCommonPropertiesRedisConfigurationPtrOutput
}

type redisCommonPropertiesRedisConfigurationPtrType RedisCommonPropertiesRedisConfigurationArgs

func RedisCommonPropertiesRedisConfigurationPtr(v *RedisCommonPropertiesRedisConfigurationArgs) RedisCommonPropertiesRedisConfigurationPtrInput {
	return (*redisCommonPropertiesRedisConfigurationPtrType)(v)
}

func (*redisCommonPropertiesRedisConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisCommonPropertiesRedisConfiguration)(nil)).Elem()
}

func (i *redisCommonPropertiesRedisConfigurationPtrType) ToRedisCommonPropertiesRedisConfigurationPtrOutput() RedisCommonPropertiesRedisConfigurationPtrOutput {
	return i.ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(context.Background())
}

func (i *redisCommonPropertiesRedisConfigurationPtrType) ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisCommonPropertiesRedisConfigurationPtrOutput)
}

type RedisCommonPropertiesRedisConfigurationOutput struct{ *pulumi.OutputState }

func (RedisCommonPropertiesRedisConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisCommonPropertiesRedisConfiguration)(nil)).Elem()
}

func (o RedisCommonPropertiesRedisConfigurationOutput) ToRedisCommonPropertiesRedisConfigurationOutput() RedisCommonPropertiesRedisConfigurationOutput {
	return o
}

func (o RedisCommonPropertiesRedisConfigurationOutput) ToRedisCommonPropertiesRedisConfigurationOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationOutput {
	return o
}

func (o RedisCommonPropertiesRedisConfigurationOutput) ToRedisCommonPropertiesRedisConfigurationPtrOutput() RedisCommonPropertiesRedisConfigurationPtrOutput {
	return o.ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(context.Background())
}

func (o RedisCommonPropertiesRedisConfigurationOutput) ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RedisCommonPropertiesRedisConfiguration) *RedisCommonPropertiesRedisConfiguration {
		return &v
	}).(RedisCommonPropertiesRedisConfigurationPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) AofStorageConnectionString0() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.AofStorageConnectionString0 }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) AofStorageConnectionString1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.AofStorageConnectionString1 }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) MaxfragmentationmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.MaxfragmentationmemoryReserved }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) MaxmemoryDelta() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.MaxmemoryDelta }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) MaxmemoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.MaxmemoryPolicy }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) MaxmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.MaxmemoryReserved }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) RdbBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.RdbBackupEnabled }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) RdbBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.RdbBackupFrequency }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) RdbBackupMaxSnapshotCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.RdbBackupMaxSnapshotCount }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationOutput) RdbStorageConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesRedisConfiguration) *string { return v.RdbStorageConnectionString }).(pulumi.StringPtrOutput)
}

type RedisCommonPropertiesRedisConfigurationPtrOutput struct{ *pulumi.OutputState }

func (RedisCommonPropertiesRedisConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisCommonPropertiesRedisConfiguration)(nil)).Elem()
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) ToRedisCommonPropertiesRedisConfigurationPtrOutput() RedisCommonPropertiesRedisConfigurationPtrOutput {
	return o
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) ToRedisCommonPropertiesRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesRedisConfigurationPtrOutput {
	return o
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) Elem() RedisCommonPropertiesRedisConfigurationOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) RedisCommonPropertiesRedisConfiguration {
		if v != nil {
			return *v
		}
		var ret RedisCommonPropertiesRedisConfiguration
		return ret
	}).(RedisCommonPropertiesRedisConfigurationOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) AofStorageConnectionString0() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AofStorageConnectionString0
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) AofStorageConnectionString1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AofStorageConnectionString1
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) MaxfragmentationmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxfragmentationmemoryReserved
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) MaxmemoryDelta() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryDelta
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) MaxmemoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryPolicy
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) MaxmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryReserved
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) RdbBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupEnabled
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) RdbBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupFrequency
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) RdbBackupMaxSnapshotCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupMaxSnapshotCount
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesRedisConfigurationPtrOutput) RdbStorageConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbStorageConnectionString
	}).(pulumi.StringPtrOutput)
}

type RedisCommonPropertiesResponseRedisConfiguration struct {
	AofStorageConnectionString0    *string `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1    *string `pulumi:"aofStorageConnectionString1"`
	Maxclients                     string  `pulumi:"maxclients"`
	MaxfragmentationmemoryReserved *string `pulumi:"maxfragmentationmemoryReserved"`
	MaxmemoryDelta                 *string `pulumi:"maxmemoryDelta"`
	MaxmemoryPolicy                *string `pulumi:"maxmemoryPolicy"`
	MaxmemoryReserved              *string `pulumi:"maxmemoryReserved"`
	RdbBackupEnabled               *string `pulumi:"rdbBackupEnabled"`
	RdbBackupFrequency             *string `pulumi:"rdbBackupFrequency"`
	RdbBackupMaxSnapshotCount      *string `pulumi:"rdbBackupMaxSnapshotCount"`
	RdbStorageConnectionString     *string `pulumi:"rdbStorageConnectionString"`
}





type RedisCommonPropertiesResponseRedisConfigurationInput interface {
	pulumi.Input

	ToRedisCommonPropertiesResponseRedisConfigurationOutput() RedisCommonPropertiesResponseRedisConfigurationOutput
	ToRedisCommonPropertiesResponseRedisConfigurationOutputWithContext(context.Context) RedisCommonPropertiesResponseRedisConfigurationOutput
}

type RedisCommonPropertiesResponseRedisConfigurationArgs struct {
	AofStorageConnectionString0    pulumi.StringPtrInput `pulumi:"aofStorageConnectionString0"`
	AofStorageConnectionString1    pulumi.StringPtrInput `pulumi:"aofStorageConnectionString1"`
	Maxclients                     pulumi.StringInput    `pulumi:"maxclients"`
	MaxfragmentationmemoryReserved pulumi.StringPtrInput `pulumi:"maxfragmentationmemoryReserved"`
	MaxmemoryDelta                 pulumi.StringPtrInput `pulumi:"maxmemoryDelta"`
	MaxmemoryPolicy                pulumi.StringPtrInput `pulumi:"maxmemoryPolicy"`
	MaxmemoryReserved              pulumi.StringPtrInput `pulumi:"maxmemoryReserved"`
	RdbBackupEnabled               pulumi.StringPtrInput `pulumi:"rdbBackupEnabled"`
	RdbBackupFrequency             pulumi.StringPtrInput `pulumi:"rdbBackupFrequency"`
	RdbBackupMaxSnapshotCount      pulumi.StringPtrInput `pulumi:"rdbBackupMaxSnapshotCount"`
	RdbStorageConnectionString     pulumi.StringPtrInput `pulumi:"rdbStorageConnectionString"`
}

func (RedisCommonPropertiesResponseRedisConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisCommonPropertiesResponseRedisConfiguration)(nil)).Elem()
}

func (i RedisCommonPropertiesResponseRedisConfigurationArgs) ToRedisCommonPropertiesResponseRedisConfigurationOutput() RedisCommonPropertiesResponseRedisConfigurationOutput {
	return i.ToRedisCommonPropertiesResponseRedisConfigurationOutputWithContext(context.Background())
}

func (i RedisCommonPropertiesResponseRedisConfigurationArgs) ToRedisCommonPropertiesResponseRedisConfigurationOutputWithContext(ctx context.Context) RedisCommonPropertiesResponseRedisConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisCommonPropertiesResponseRedisConfigurationOutput)
}

func (i RedisCommonPropertiesResponseRedisConfigurationArgs) ToRedisCommonPropertiesResponseRedisConfigurationPtrOutput() RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return i.ToRedisCommonPropertiesResponseRedisConfigurationPtrOutputWithContext(context.Background())
}

func (i RedisCommonPropertiesResponseRedisConfigurationArgs) ToRedisCommonPropertiesResponseRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisCommonPropertiesResponseRedisConfigurationOutput).ToRedisCommonPropertiesResponseRedisConfigurationPtrOutputWithContext(ctx)
}









type RedisCommonPropertiesResponseRedisConfigurationPtrInput interface {
	pulumi.Input

	ToRedisCommonPropertiesResponseRedisConfigurationPtrOutput() RedisCommonPropertiesResponseRedisConfigurationPtrOutput
	ToRedisCommonPropertiesResponseRedisConfigurationPtrOutputWithContext(context.Context) RedisCommonPropertiesResponseRedisConfigurationPtrOutput
}

type redisCommonPropertiesResponseRedisConfigurationPtrType RedisCommonPropertiesResponseRedisConfigurationArgs

func RedisCommonPropertiesResponseRedisConfigurationPtr(v *RedisCommonPropertiesResponseRedisConfigurationArgs) RedisCommonPropertiesResponseRedisConfigurationPtrInput {
	return (*redisCommonPropertiesResponseRedisConfigurationPtrType)(v)
}

func (*redisCommonPropertiesResponseRedisConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisCommonPropertiesResponseRedisConfiguration)(nil)).Elem()
}

func (i *redisCommonPropertiesResponseRedisConfigurationPtrType) ToRedisCommonPropertiesResponseRedisConfigurationPtrOutput() RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return i.ToRedisCommonPropertiesResponseRedisConfigurationPtrOutputWithContext(context.Background())
}

func (i *redisCommonPropertiesResponseRedisConfigurationPtrType) ToRedisCommonPropertiesResponseRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisCommonPropertiesResponseRedisConfigurationPtrOutput)
}

type RedisCommonPropertiesResponseRedisConfigurationOutput struct{ *pulumi.OutputState }

func (RedisCommonPropertiesResponseRedisConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisCommonPropertiesResponseRedisConfiguration)(nil)).Elem()
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) ToRedisCommonPropertiesResponseRedisConfigurationOutput() RedisCommonPropertiesResponseRedisConfigurationOutput {
	return o
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) ToRedisCommonPropertiesResponseRedisConfigurationOutputWithContext(ctx context.Context) RedisCommonPropertiesResponseRedisConfigurationOutput {
	return o
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) ToRedisCommonPropertiesResponseRedisConfigurationPtrOutput() RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return o.ToRedisCommonPropertiesResponseRedisConfigurationPtrOutputWithContext(context.Background())
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) ToRedisCommonPropertiesResponseRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RedisCommonPropertiesResponseRedisConfiguration) *RedisCommonPropertiesResponseRedisConfiguration {
		return &v
	}).(RedisCommonPropertiesResponseRedisConfigurationPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) AofStorageConnectionString0() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.AofStorageConnectionString0 }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) AofStorageConnectionString1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.AofStorageConnectionString1 }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) Maxclients() pulumi.StringOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) string { return v.Maxclients }).(pulumi.StringOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) MaxfragmentationmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string {
		return v.MaxfragmentationmemoryReserved
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) MaxmemoryDelta() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.MaxmemoryDelta }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) MaxmemoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.MaxmemoryPolicy }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) MaxmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.MaxmemoryReserved }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) RdbBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.RdbBackupEnabled }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) RdbBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.RdbBackupFrequency }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) RdbBackupMaxSnapshotCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.RdbBackupMaxSnapshotCount }).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationOutput) RdbStorageConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RedisCommonPropertiesResponseRedisConfiguration) *string { return v.RdbStorageConnectionString }).(pulumi.StringPtrOutput)
}

type RedisCommonPropertiesResponseRedisConfigurationPtrOutput struct{ *pulumi.OutputState }

func (RedisCommonPropertiesResponseRedisConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisCommonPropertiesResponseRedisConfiguration)(nil)).Elem()
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) ToRedisCommonPropertiesResponseRedisConfigurationPtrOutput() RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return o
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) ToRedisCommonPropertiesResponseRedisConfigurationPtrOutputWithContext(ctx context.Context) RedisCommonPropertiesResponseRedisConfigurationPtrOutput {
	return o
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) Elem() RedisCommonPropertiesResponseRedisConfigurationOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) RedisCommonPropertiesResponseRedisConfiguration {
		if v != nil {
			return *v
		}
		var ret RedisCommonPropertiesResponseRedisConfiguration
		return ret
	}).(RedisCommonPropertiesResponseRedisConfigurationOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) AofStorageConnectionString0() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AofStorageConnectionString0
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) AofStorageConnectionString1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AofStorageConnectionString1
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) Maxclients() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.Maxclients
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) MaxfragmentationmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxfragmentationmemoryReserved
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) MaxmemoryDelta() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryDelta
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) MaxmemoryPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryPolicy
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) MaxmemoryReserved() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.MaxmemoryReserved
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) RdbBackupEnabled() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupEnabled
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) RdbBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupFrequency
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) RdbBackupMaxSnapshotCount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbBackupMaxSnapshotCount
	}).(pulumi.StringPtrOutput)
}

func (o RedisCommonPropertiesResponseRedisConfigurationPtrOutput) RdbStorageConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisCommonPropertiesResponseRedisConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RdbStorageConnectionString
	}).(pulumi.StringPtrOutput)
}

type RedisInstanceDetailsResponse struct {
	IsMaster   bool   `pulumi:"isMaster"`
	IsPrimary  bool   `pulumi:"isPrimary"`
	NonSslPort int    `pulumi:"nonSslPort"`
	ShardId    int    `pulumi:"shardId"`
	SslPort    int    `pulumi:"sslPort"`
	Zone       string `pulumi:"zone"`
}





type RedisInstanceDetailsResponseInput interface {
	pulumi.Input

	ToRedisInstanceDetailsResponseOutput() RedisInstanceDetailsResponseOutput
	ToRedisInstanceDetailsResponseOutputWithContext(context.Context) RedisInstanceDetailsResponseOutput
}

type RedisInstanceDetailsResponseArgs struct {
	IsMaster   pulumi.BoolInput   `pulumi:"isMaster"`
	IsPrimary  pulumi.BoolInput   `pulumi:"isPrimary"`
	NonSslPort pulumi.IntInput    `pulumi:"nonSslPort"`
	ShardId    pulumi.IntInput    `pulumi:"shardId"`
	SslPort    pulumi.IntInput    `pulumi:"sslPort"`
	Zone       pulumi.StringInput `pulumi:"zone"`
}

func (RedisInstanceDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisInstanceDetailsResponse)(nil)).Elem()
}

func (i RedisInstanceDetailsResponseArgs) ToRedisInstanceDetailsResponseOutput() RedisInstanceDetailsResponseOutput {
	return i.ToRedisInstanceDetailsResponseOutputWithContext(context.Background())
}

func (i RedisInstanceDetailsResponseArgs) ToRedisInstanceDetailsResponseOutputWithContext(ctx context.Context) RedisInstanceDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisInstanceDetailsResponseOutput)
}





type RedisInstanceDetailsResponseArrayInput interface {
	pulumi.Input

	ToRedisInstanceDetailsResponseArrayOutput() RedisInstanceDetailsResponseArrayOutput
	ToRedisInstanceDetailsResponseArrayOutputWithContext(context.Context) RedisInstanceDetailsResponseArrayOutput
}

type RedisInstanceDetailsResponseArray []RedisInstanceDetailsResponseInput

func (RedisInstanceDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisInstanceDetailsResponse)(nil)).Elem()
}

func (i RedisInstanceDetailsResponseArray) ToRedisInstanceDetailsResponseArrayOutput() RedisInstanceDetailsResponseArrayOutput {
	return i.ToRedisInstanceDetailsResponseArrayOutputWithContext(context.Background())
}

func (i RedisInstanceDetailsResponseArray) ToRedisInstanceDetailsResponseArrayOutputWithContext(ctx context.Context) RedisInstanceDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisInstanceDetailsResponseArrayOutput)
}

type RedisInstanceDetailsResponseOutput struct{ *pulumi.OutputState }

func (RedisInstanceDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisInstanceDetailsResponse)(nil)).Elem()
}

func (o RedisInstanceDetailsResponseOutput) ToRedisInstanceDetailsResponseOutput() RedisInstanceDetailsResponseOutput {
	return o
}

func (o RedisInstanceDetailsResponseOutput) ToRedisInstanceDetailsResponseOutputWithContext(ctx context.Context) RedisInstanceDetailsResponseOutput {
	return o
}

func (o RedisInstanceDetailsResponseOutput) IsMaster() pulumi.BoolOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) bool { return v.IsMaster }).(pulumi.BoolOutput)
}

func (o RedisInstanceDetailsResponseOutput) IsPrimary() pulumi.BoolOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) bool { return v.IsPrimary }).(pulumi.BoolOutput)
}

func (o RedisInstanceDetailsResponseOutput) NonSslPort() pulumi.IntOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) int { return v.NonSslPort }).(pulumi.IntOutput)
}

func (o RedisInstanceDetailsResponseOutput) ShardId() pulumi.IntOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) int { return v.ShardId }).(pulumi.IntOutput)
}

func (o RedisInstanceDetailsResponseOutput) SslPort() pulumi.IntOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) int { return v.SslPort }).(pulumi.IntOutput)
}

func (o RedisInstanceDetailsResponseOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v RedisInstanceDetailsResponse) string { return v.Zone }).(pulumi.StringOutput)
}

type RedisInstanceDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (RedisInstanceDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisInstanceDetailsResponse)(nil)).Elem()
}

func (o RedisInstanceDetailsResponseArrayOutput) ToRedisInstanceDetailsResponseArrayOutput() RedisInstanceDetailsResponseArrayOutput {
	return o
}

func (o RedisInstanceDetailsResponseArrayOutput) ToRedisInstanceDetailsResponseArrayOutputWithContext(ctx context.Context) RedisInstanceDetailsResponseArrayOutput {
	return o
}

func (o RedisInstanceDetailsResponseArrayOutput) Index(i pulumi.IntInput) RedisInstanceDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RedisInstanceDetailsResponse {
		return vs[0].([]RedisInstanceDetailsResponse)[vs[1].(int)]
	}).(RedisInstanceDetailsResponseOutput)
}

type RedisLinkedServerResponse struct {
	Id string `pulumi:"id"`
}





type RedisLinkedServerResponseInput interface {
	pulumi.Input

	ToRedisLinkedServerResponseOutput() RedisLinkedServerResponseOutput
	ToRedisLinkedServerResponseOutputWithContext(context.Context) RedisLinkedServerResponseOutput
}

type RedisLinkedServerResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (RedisLinkedServerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisLinkedServerResponse)(nil)).Elem()
}

func (i RedisLinkedServerResponseArgs) ToRedisLinkedServerResponseOutput() RedisLinkedServerResponseOutput {
	return i.ToRedisLinkedServerResponseOutputWithContext(context.Background())
}

func (i RedisLinkedServerResponseArgs) ToRedisLinkedServerResponseOutputWithContext(ctx context.Context) RedisLinkedServerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisLinkedServerResponseOutput)
}





type RedisLinkedServerResponseArrayInput interface {
	pulumi.Input

	ToRedisLinkedServerResponseArrayOutput() RedisLinkedServerResponseArrayOutput
	ToRedisLinkedServerResponseArrayOutputWithContext(context.Context) RedisLinkedServerResponseArrayOutput
}

type RedisLinkedServerResponseArray []RedisLinkedServerResponseInput

func (RedisLinkedServerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisLinkedServerResponse)(nil)).Elem()
}

func (i RedisLinkedServerResponseArray) ToRedisLinkedServerResponseArrayOutput() RedisLinkedServerResponseArrayOutput {
	return i.ToRedisLinkedServerResponseArrayOutputWithContext(context.Background())
}

func (i RedisLinkedServerResponseArray) ToRedisLinkedServerResponseArrayOutputWithContext(ctx context.Context) RedisLinkedServerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisLinkedServerResponseArrayOutput)
}

type RedisLinkedServerResponseOutput struct{ *pulumi.OutputState }

func (RedisLinkedServerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RedisLinkedServerResponse)(nil)).Elem()
}

func (o RedisLinkedServerResponseOutput) ToRedisLinkedServerResponseOutput() RedisLinkedServerResponseOutput {
	return o
}

func (o RedisLinkedServerResponseOutput) ToRedisLinkedServerResponseOutputWithContext(ctx context.Context) RedisLinkedServerResponseOutput {
	return o
}

func (o RedisLinkedServerResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RedisLinkedServerResponse) string { return v.Id }).(pulumi.StringOutput)
}

type RedisLinkedServerResponseArrayOutput struct{ *pulumi.OutputState }

func (RedisLinkedServerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RedisLinkedServerResponse)(nil)).Elem()
}

func (o RedisLinkedServerResponseArrayOutput) ToRedisLinkedServerResponseArrayOutput() RedisLinkedServerResponseArrayOutput {
	return o
}

func (o RedisLinkedServerResponseArrayOutput) ToRedisLinkedServerResponseArrayOutputWithContext(ctx context.Context) RedisLinkedServerResponseArrayOutput {
	return o
}

func (o RedisLinkedServerResponseArrayOutput) Index(i pulumi.IntInput) RedisLinkedServerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RedisLinkedServerResponse {
		return vs[0].([]RedisLinkedServerResponse)[vs[1].(int)]
	}).(RedisLinkedServerResponseOutput)
}

type ScheduleEntry struct {
	DayOfWeek         DayOfWeek `pulumi:"dayOfWeek"`
	MaintenanceWindow *string   `pulumi:"maintenanceWindow"`
	StartHourUtc      int       `pulumi:"startHourUtc"`
}





type ScheduleEntryInput interface {
	pulumi.Input

	ToScheduleEntryOutput() ScheduleEntryOutput
	ToScheduleEntryOutputWithContext(context.Context) ScheduleEntryOutput
}

type ScheduleEntryArgs struct {
	DayOfWeek         DayOfWeekInput        `pulumi:"dayOfWeek"`
	MaintenanceWindow pulumi.StringPtrInput `pulumi:"maintenanceWindow"`
	StartHourUtc      pulumi.IntInput       `pulumi:"startHourUtc"`
}

func (ScheduleEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntry)(nil)).Elem()
}

func (i ScheduleEntryArgs) ToScheduleEntryOutput() ScheduleEntryOutput {
	return i.ToScheduleEntryOutputWithContext(context.Background())
}

func (i ScheduleEntryArgs) ToScheduleEntryOutputWithContext(ctx context.Context) ScheduleEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryOutput)
}





type ScheduleEntryArrayInput interface {
	pulumi.Input

	ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput
	ToScheduleEntryArrayOutputWithContext(context.Context) ScheduleEntryArrayOutput
}

type ScheduleEntryArray []ScheduleEntryInput

func (ScheduleEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntry)(nil)).Elem()
}

func (i ScheduleEntryArray) ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput {
	return i.ToScheduleEntryArrayOutputWithContext(context.Background())
}

func (i ScheduleEntryArray) ToScheduleEntryArrayOutputWithContext(ctx context.Context) ScheduleEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryArrayOutput)
}

type ScheduleEntryOutput struct{ *pulumi.OutputState }

func (ScheduleEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntry)(nil)).Elem()
}

func (o ScheduleEntryOutput) ToScheduleEntryOutput() ScheduleEntryOutput {
	return o
}

func (o ScheduleEntryOutput) ToScheduleEntryOutputWithContext(ctx context.Context) ScheduleEntryOutput {
	return o
}

func (o ScheduleEntryOutput) DayOfWeek() DayOfWeekOutput {
	return o.ApplyT(func(v ScheduleEntry) DayOfWeek { return v.DayOfWeek }).(DayOfWeekOutput)
}

func (o ScheduleEntryOutput) MaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleEntry) *string { return v.MaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o ScheduleEntryOutput) StartHourUtc() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleEntry) int { return v.StartHourUtc }).(pulumi.IntOutput)
}

type ScheduleEntryArrayOutput struct{ *pulumi.OutputState }

func (ScheduleEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntry)(nil)).Elem()
}

func (o ScheduleEntryArrayOutput) ToScheduleEntryArrayOutput() ScheduleEntryArrayOutput {
	return o
}

func (o ScheduleEntryArrayOutput) ToScheduleEntryArrayOutputWithContext(ctx context.Context) ScheduleEntryArrayOutput {
	return o
}

func (o ScheduleEntryArrayOutput) Index(i pulumi.IntInput) ScheduleEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleEntry {
		return vs[0].([]ScheduleEntry)[vs[1].(int)]
	}).(ScheduleEntryOutput)
}

type ScheduleEntryResponse struct {
	DayOfWeek         string  `pulumi:"dayOfWeek"`
	MaintenanceWindow *string `pulumi:"maintenanceWindow"`
	StartHourUtc      int     `pulumi:"startHourUtc"`
}





type ScheduleEntryResponseInput interface {
	pulumi.Input

	ToScheduleEntryResponseOutput() ScheduleEntryResponseOutput
	ToScheduleEntryResponseOutputWithContext(context.Context) ScheduleEntryResponseOutput
}

type ScheduleEntryResponseArgs struct {
	DayOfWeek         pulumi.StringInput    `pulumi:"dayOfWeek"`
	MaintenanceWindow pulumi.StringPtrInput `pulumi:"maintenanceWindow"`
	StartHourUtc      pulumi.IntInput       `pulumi:"startHourUtc"`
}

func (ScheduleEntryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntryResponse)(nil)).Elem()
}

func (i ScheduleEntryResponseArgs) ToScheduleEntryResponseOutput() ScheduleEntryResponseOutput {
	return i.ToScheduleEntryResponseOutputWithContext(context.Background())
}

func (i ScheduleEntryResponseArgs) ToScheduleEntryResponseOutputWithContext(ctx context.Context) ScheduleEntryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryResponseOutput)
}





type ScheduleEntryResponseArrayInput interface {
	pulumi.Input

	ToScheduleEntryResponseArrayOutput() ScheduleEntryResponseArrayOutput
	ToScheduleEntryResponseArrayOutputWithContext(context.Context) ScheduleEntryResponseArrayOutput
}

type ScheduleEntryResponseArray []ScheduleEntryResponseInput

func (ScheduleEntryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntryResponse)(nil)).Elem()
}

func (i ScheduleEntryResponseArray) ToScheduleEntryResponseArrayOutput() ScheduleEntryResponseArrayOutput {
	return i.ToScheduleEntryResponseArrayOutputWithContext(context.Background())
}

func (i ScheduleEntryResponseArray) ToScheduleEntryResponseArrayOutputWithContext(ctx context.Context) ScheduleEntryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleEntryResponseArrayOutput)
}

type ScheduleEntryResponseOutput struct{ *pulumi.OutputState }

func (ScheduleEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleEntryResponse)(nil)).Elem()
}

func (o ScheduleEntryResponseOutput) ToScheduleEntryResponseOutput() ScheduleEntryResponseOutput {
	return o
}

func (o ScheduleEntryResponseOutput) ToScheduleEntryResponseOutputWithContext(ctx context.Context) ScheduleEntryResponseOutput {
	return o
}

func (o ScheduleEntryResponseOutput) DayOfWeek() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) string { return v.DayOfWeek }).(pulumi.StringOutput)
}

func (o ScheduleEntryResponseOutput) MaintenanceWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) *string { return v.MaintenanceWindow }).(pulumi.StringPtrOutput)
}

func (o ScheduleEntryResponseOutput) StartHourUtc() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleEntryResponse) int { return v.StartHourUtc }).(pulumi.IntOutput)
}

type ScheduleEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (ScheduleEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduleEntryResponse)(nil)).Elem()
}

func (o ScheduleEntryResponseArrayOutput) ToScheduleEntryResponseArrayOutput() ScheduleEntryResponseArrayOutput {
	return o
}

func (o ScheduleEntryResponseArrayOutput) ToScheduleEntryResponseArrayOutputWithContext(ctx context.Context) ScheduleEntryResponseArrayOutput {
	return o
}

func (o ScheduleEntryResponseArrayOutput) Index(i pulumi.IntInput) ScheduleEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduleEntryResponse {
		return vs[0].([]ScheduleEntryResponse)[vs[1].(int)]
	}).(ScheduleEntryResponseOutput)
}

type Sku struct {
	Capacity int    `pulumi:"capacity"`
	Family   string `pulumi:"family"`
	Name     string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntInput    `pulumi:"capacity"`
	Family   pulumi.StringInput `pulumi:"family"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v Sku) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity int    `pulumi:"capacity"`
	Family   string `pulumi:"family"`
	Name     string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntInput    `pulumi:"capacity"`
	Family   pulumi.StringInput `pulumi:"family"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v SkuResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseInput)(nil)).Elem(), PrivateEndpointConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseArrayInput)(nil)).Elem(), PrivateEndpointConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponseInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponsePtrInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStateResponsePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisAccessKeysResponseInput)(nil)).Elem(), RedisAccessKeysResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisAccessKeysResponsePtrInput)(nil)).Elem(), RedisAccessKeysResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisCommonPropertiesRedisConfigurationInput)(nil)).Elem(), RedisCommonPropertiesRedisConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisCommonPropertiesRedisConfigurationPtrInput)(nil)).Elem(), RedisCommonPropertiesRedisConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisCommonPropertiesResponseRedisConfigurationInput)(nil)).Elem(), RedisCommonPropertiesResponseRedisConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisCommonPropertiesResponseRedisConfigurationPtrInput)(nil)).Elem(), RedisCommonPropertiesResponseRedisConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisInstanceDetailsResponseInput)(nil)).Elem(), RedisInstanceDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisInstanceDetailsResponseArrayInput)(nil)).Elem(), RedisInstanceDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisLinkedServerResponseInput)(nil)).Elem(), RedisLinkedServerResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisLinkedServerResponseArrayInput)(nil)).Elem(), RedisLinkedServerResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleEntryInput)(nil)).Elem(), ScheduleEntryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleEntryArrayInput)(nil)).Elem(), ScheduleEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleEntryResponseInput)(nil)).Elem(), ScheduleEntryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleEntryResponseArrayInput)(nil)).Elem(), ScheduleEntryResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(RedisAccessKeysResponseOutput{})
	pulumi.RegisterOutputType(RedisAccessKeysResponsePtrOutput{})
	pulumi.RegisterOutputType(RedisCommonPropertiesRedisConfigurationOutput{})
	pulumi.RegisterOutputType(RedisCommonPropertiesRedisConfigurationPtrOutput{})
	pulumi.RegisterOutputType(RedisCommonPropertiesResponseRedisConfigurationOutput{})
	pulumi.RegisterOutputType(RedisCommonPropertiesResponseRedisConfigurationPtrOutput{})
	pulumi.RegisterOutputType(RedisInstanceDetailsResponseOutput{})
	pulumi.RegisterOutputType(RedisInstanceDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(RedisLinkedServerResponseOutput{})
	pulumi.RegisterOutputType(RedisLinkedServerResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduleEntryOutput{})
	pulumi.RegisterOutputType(ScheduleEntryArrayOutput{})
	pulumi.RegisterOutputType(ScheduleEntryResponseOutput{})
	pulumi.RegisterOutputType(ScheduleEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
