


package v20171201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointPropertyResponse struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyResponseInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput
	ToPrivateEndpointPropertyResponseOutputWithContext(context.Context) PrivateEndpointPropertyResponseOutput
}

type PrivateEndpointPropertyResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return i.ToPrivateEndpointPropertyResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput)
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput).ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput
	ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Context) PrivateEndpointPropertyResponsePtrOutput
}

type privateEndpointPropertyResponsePtrType PrivateEndpointPropertyResponseArgs

func PrivateEndpointPropertyResponsePtr(v *PrivateEndpointPropertyResponseArgs) PrivateEndpointPropertyResponsePtrInput {
	return (*privateEndpointPropertyResponsePtrType)(v)
}

func (*privateEndpointPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponsePtrOutput)
}

type PrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointPropertyResponse) *PrivateEndpointPropertyResponse {
		return &v
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointPropertyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointPropertyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) Elem() PrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) PrivateEndpointPropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointPropertyResponse
		return ret
	}).(PrivateEndpointPropertyResponseOutput)
}

func (o PrivateEndpointPropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentity struct {
	Type *string `pulumi:"type"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ResourceIdentityResponseInput interface {
	pulumi.Input

	ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput
	ToResourceIdentityResponseOutputWithContext(context.Context) ResourceIdentityResponseOutput
}

type ResourceIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ResourceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return i.ToResourceIdentityResponseOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput)
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdentityResponseArgs) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponseOutput).ToResourceIdentityResponsePtrOutputWithContext(ctx)
}









type ResourceIdentityResponsePtrInput interface {
	pulumi.Input

	ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput
	ToResourceIdentityResponsePtrOutputWithContext(context.Context) ResourceIdentityResponsePtrOutput
}

type resourceIdentityResponsePtrType ResourceIdentityResponseArgs

func ResourceIdentityResponsePtr(v *ResourceIdentityResponseArgs) ResourceIdentityResponsePtrInput {
	return (*resourceIdentityResponsePtrType)(v)
}

func (*resourceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return i.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdentityResponsePtrType) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityResponsePtrOutput)
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o.ToResourceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityResponse) *ResourceIdentityResponse {
		return &v
	}).(ResourceIdentityResponsePtrOutput)
}

func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ServerPrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                         `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ServerPrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                                   `pulumi:"provisioningState"`
}





type ServerPrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToServerPrivateEndpointConnectionPropertiesResponseOutput() ServerPrivateEndpointConnectionPropertiesResponseOutput
	ToServerPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) ServerPrivateEndpointConnectionPropertiesResponseOutput
}

type ServerPrivateEndpointConnectionPropertiesResponseArgs struct {
	PrivateEndpoint                   PrivateEndpointPropertyResponsePtrInput                         `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState ServerPrivateLinkServiceConnectionStatePropertyResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                              `pulumi:"provisioningState"`
}

func (ServerPrivateEndpointConnectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i ServerPrivateEndpointConnectionPropertiesResponseArgs) ToServerPrivateEndpointConnectionPropertiesResponseOutput() ServerPrivateEndpointConnectionPropertiesResponseOutput {
	return i.ToServerPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Background())
}

func (i ServerPrivateEndpointConnectionPropertiesResponseArgs) ToServerPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPrivateEndpointConnectionPropertiesResponseOutput)
}

type ServerPrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServerPrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o ServerPrivateEndpointConnectionPropertiesResponseOutput) ToServerPrivateEndpointConnectionPropertiesResponseOutput() ServerPrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o ServerPrivateEndpointConnectionPropertiesResponseOutput) ToServerPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o ServerPrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v ServerPrivateEndpointConnectionPropertiesResponse) *PrivateEndpointPropertyResponse {
		return v.PrivateEndpoint
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o ServerPrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v ServerPrivateEndpointConnectionPropertiesResponse) *ServerPrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o ServerPrivateEndpointConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPrivateEndpointConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type ServerPrivateEndpointConnectionResponse struct {
	Id         string                                            `pulumi:"id"`
	Properties ServerPrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
}





type ServerPrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToServerPrivateEndpointConnectionResponseOutput() ServerPrivateEndpointConnectionResponseOutput
	ToServerPrivateEndpointConnectionResponseOutputWithContext(context.Context) ServerPrivateEndpointConnectionResponseOutput
}

type ServerPrivateEndpointConnectionResponseArgs struct {
	Id         pulumi.StringInput                                     `pulumi:"id"`
	Properties ServerPrivateEndpointConnectionPropertiesResponseInput `pulumi:"properties"`
}

func (ServerPrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i ServerPrivateEndpointConnectionResponseArgs) ToServerPrivateEndpointConnectionResponseOutput() ServerPrivateEndpointConnectionResponseOutput {
	return i.ToServerPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i ServerPrivateEndpointConnectionResponseArgs) ToServerPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPrivateEndpointConnectionResponseOutput)
}





type ServerPrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToServerPrivateEndpointConnectionResponseArrayOutput() ServerPrivateEndpointConnectionResponseArrayOutput
	ToServerPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) ServerPrivateEndpointConnectionResponseArrayOutput
}

type ServerPrivateEndpointConnectionResponseArray []ServerPrivateEndpointConnectionResponseInput

func (ServerPrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i ServerPrivateEndpointConnectionResponseArray) ToServerPrivateEndpointConnectionResponseArrayOutput() ServerPrivateEndpointConnectionResponseArrayOutput {
	return i.ToServerPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i ServerPrivateEndpointConnectionResponseArray) ToServerPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPrivateEndpointConnectionResponseArrayOutput)
}

type ServerPrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (ServerPrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o ServerPrivateEndpointConnectionResponseOutput) ToServerPrivateEndpointConnectionResponseOutput() ServerPrivateEndpointConnectionResponseOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseOutput) ToServerPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionResponseOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ServerPrivateEndpointConnectionResponseOutput) Properties() ServerPrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v ServerPrivateEndpointConnectionResponse) ServerPrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(ServerPrivateEndpointConnectionPropertiesResponseOutput)
}

type ServerPrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (ServerPrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o ServerPrivateEndpointConnectionResponseArrayOutput) ToServerPrivateEndpointConnectionResponseArrayOutput() ServerPrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseArrayOutput) ToServerPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) ServerPrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o ServerPrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) ServerPrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerPrivateEndpointConnectionResponse {
		return vs[0].([]ServerPrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(ServerPrivateEndpointConnectionResponseOutput)
}

type ServerPrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}





type ServerPrivateLinkServiceConnectionStatePropertyResponseInput interface {
	pulumi.Input

	ToServerPrivateLinkServiceConnectionStatePropertyResponseOutput() ServerPrivateLinkServiceConnectionStatePropertyResponseOutput
	ToServerPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Context) ServerPrivateLinkServiceConnectionStatePropertyResponseOutput
}

type ServerPrivateLinkServiceConnectionStatePropertyResponseArgs struct {
	ActionsRequired pulumi.StringInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput `pulumi:"description"`
	Status          pulumi.StringInput `pulumi:"status"`
}

func (ServerPrivateLinkServiceConnectionStatePropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i ServerPrivateLinkServiceConnectionStatePropertyResponseArgs) ToServerPrivateLinkServiceConnectionStatePropertyResponseOutput() ServerPrivateLinkServiceConnectionStatePropertyResponseOutput {
	return i.ToServerPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Background())
}

func (i ServerPrivateLinkServiceConnectionStatePropertyResponseArgs) ToServerPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) ServerPrivateLinkServiceConnectionStatePropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (i ServerPrivateLinkServiceConnectionStatePropertyResponseArgs) ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i ServerPrivateLinkServiceConnectionStatePropertyResponseArgs) ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPrivateLinkServiceConnectionStatePropertyResponseOutput).ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx)
}









type ServerPrivateLinkServiceConnectionStatePropertyResponsePtrInput interface {
	pulumi.Input

	ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput
	ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Context) ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput
}

type serverPrivateLinkServiceConnectionStatePropertyResponsePtrType ServerPrivateLinkServiceConnectionStatePropertyResponseArgs

func ServerPrivateLinkServiceConnectionStatePropertyResponsePtr(v *ServerPrivateLinkServiceConnectionStatePropertyResponseArgs) ServerPrivateLinkServiceConnectionStatePropertyResponsePtrInput {
	return (*serverPrivateLinkServiceConnectionStatePropertyResponsePtrType)(v)
}

func (*serverPrivateLinkServiceConnectionStatePropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i *serverPrivateLinkServiceConnectionStatePropertyResponsePtrType) ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i *serverPrivateLinkServiceConnectionStatePropertyResponsePtrType) ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

type ServerPrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (ServerPrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponseOutput) ToServerPrivateLinkServiceConnectionStatePropertyResponseOutput() ServerPrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponseOutput) ToServerPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) ServerPrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponseOutput) ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponseOutput) ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerPrivateLinkServiceConnectionStatePropertyResponse) *ServerPrivateLinkServiceConnectionStatePropertyResponse {
		return &v
	}).(ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPrivateLinkServiceConnectionStatePropertyResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPrivateLinkServiceConnectionStatePropertyResponse) string { return v.Status }).(pulumi.StringOutput)
}

type ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerPrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() ServerPrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *ServerPrivateLinkServiceConnectionStatePropertyResponse) ServerPrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret ServerPrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(ServerPrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerPrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerPrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerPrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type ServerPropertiesForDefaultCreate struct {
	AdministratorLogin         string              `pulumi:"administratorLogin"`
	AdministratorLoginPassword string              `pulumi:"administratorLoginPassword"`
	CreateMode                 string              `pulumi:"createMode"`
	InfrastructureEncryption   *string             `pulumi:"infrastructureEncryption"`
	MinimalTlsVersion          *string             `pulumi:"minimalTlsVersion"`
	PublicNetworkAccess        *string             `pulumi:"publicNetworkAccess"`
	SslEnforcement             *SslEnforcementEnum `pulumi:"sslEnforcement"`
	StorageProfile             *StorageProfile     `pulumi:"storageProfile"`
	Version                    *string             `pulumi:"version"`
}





type ServerPropertiesForDefaultCreateInput interface {
	pulumi.Input

	ToServerPropertiesForDefaultCreateOutput() ServerPropertiesForDefaultCreateOutput
	ToServerPropertiesForDefaultCreateOutputWithContext(context.Context) ServerPropertiesForDefaultCreateOutput
}

type ServerPropertiesForDefaultCreateArgs struct {
	AdministratorLogin         pulumi.StringInput         `pulumi:"administratorLogin"`
	AdministratorLoginPassword pulumi.StringInput         `pulumi:"administratorLoginPassword"`
	CreateMode                 pulumi.StringInput         `pulumi:"createMode"`
	InfrastructureEncryption   pulumi.StringPtrInput      `pulumi:"infrastructureEncryption"`
	MinimalTlsVersion          pulumi.StringPtrInput      `pulumi:"minimalTlsVersion"`
	PublicNetworkAccess        pulumi.StringPtrInput      `pulumi:"publicNetworkAccess"`
	SslEnforcement             SslEnforcementEnumPtrInput `pulumi:"sslEnforcement"`
	StorageProfile             StorageProfilePtrInput     `pulumi:"storageProfile"`
	Version                    pulumi.StringPtrInput      `pulumi:"version"`
}

func (ServerPropertiesForDefaultCreateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForDefaultCreate)(nil)).Elem()
}

func (i ServerPropertiesForDefaultCreateArgs) ToServerPropertiesForDefaultCreateOutput() ServerPropertiesForDefaultCreateOutput {
	return i.ToServerPropertiesForDefaultCreateOutputWithContext(context.Background())
}

func (i ServerPropertiesForDefaultCreateArgs) ToServerPropertiesForDefaultCreateOutputWithContext(ctx context.Context) ServerPropertiesForDefaultCreateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesForDefaultCreateOutput)
}

type ServerPropertiesForDefaultCreateOutput struct{ *pulumi.OutputState }

func (ServerPropertiesForDefaultCreateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForDefaultCreate)(nil)).Elem()
}

func (o ServerPropertiesForDefaultCreateOutput) ToServerPropertiesForDefaultCreateOutput() ServerPropertiesForDefaultCreateOutput {
	return o
}

func (o ServerPropertiesForDefaultCreateOutput) ToServerPropertiesForDefaultCreateOutputWithContext(ctx context.Context) ServerPropertiesForDefaultCreateOutput {
	return o
}

func (o ServerPropertiesForDefaultCreateOutput) AdministratorLogin() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForDefaultCreate) string { return v.AdministratorLogin }).(pulumi.StringOutput)
}

func (o ServerPropertiesForDefaultCreateOutput) AdministratorLoginPassword() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForDefaultCreate) string { return v.AdministratorLoginPassword }).(pulumi.StringOutput)
}

func (o ServerPropertiesForDefaultCreateOutput) CreateMode() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForDefaultCreate) string { return v.CreateMode }).(pulumi.StringOutput)
}

func (o ServerPropertiesForDefaultCreateOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForDefaultCreate) *string { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForDefaultCreateOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForDefaultCreate) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForDefaultCreateOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForDefaultCreate) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForDefaultCreateOutput) SslEnforcement() SslEnforcementEnumPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForDefaultCreate) *SslEnforcementEnum { return v.SslEnforcement }).(SslEnforcementEnumPtrOutput)
}

func (o ServerPropertiesForDefaultCreateOutput) StorageProfile() StorageProfilePtrOutput {
	return o.ApplyT(func(v ServerPropertiesForDefaultCreate) *StorageProfile { return v.StorageProfile }).(StorageProfilePtrOutput)
}

func (o ServerPropertiesForDefaultCreateOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForDefaultCreate) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ServerPropertiesForGeoRestore struct {
	CreateMode               string              `pulumi:"createMode"`
	InfrastructureEncryption *string             `pulumi:"infrastructureEncryption"`
	MinimalTlsVersion        *string             `pulumi:"minimalTlsVersion"`
	PublicNetworkAccess      *string             `pulumi:"publicNetworkAccess"`
	SourceServerId           string              `pulumi:"sourceServerId"`
	SslEnforcement           *SslEnforcementEnum `pulumi:"sslEnforcement"`
	StorageProfile           *StorageProfile     `pulumi:"storageProfile"`
	Version                  *string             `pulumi:"version"`
}





type ServerPropertiesForGeoRestoreInput interface {
	pulumi.Input

	ToServerPropertiesForGeoRestoreOutput() ServerPropertiesForGeoRestoreOutput
	ToServerPropertiesForGeoRestoreOutputWithContext(context.Context) ServerPropertiesForGeoRestoreOutput
}

type ServerPropertiesForGeoRestoreArgs struct {
	CreateMode               pulumi.StringInput         `pulumi:"createMode"`
	InfrastructureEncryption pulumi.StringPtrInput      `pulumi:"infrastructureEncryption"`
	MinimalTlsVersion        pulumi.StringPtrInput      `pulumi:"minimalTlsVersion"`
	PublicNetworkAccess      pulumi.StringPtrInput      `pulumi:"publicNetworkAccess"`
	SourceServerId           pulumi.StringInput         `pulumi:"sourceServerId"`
	SslEnforcement           SslEnforcementEnumPtrInput `pulumi:"sslEnforcement"`
	StorageProfile           StorageProfilePtrInput     `pulumi:"storageProfile"`
	Version                  pulumi.StringPtrInput      `pulumi:"version"`
}

func (ServerPropertiesForGeoRestoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForGeoRestore)(nil)).Elem()
}

func (i ServerPropertiesForGeoRestoreArgs) ToServerPropertiesForGeoRestoreOutput() ServerPropertiesForGeoRestoreOutput {
	return i.ToServerPropertiesForGeoRestoreOutputWithContext(context.Background())
}

func (i ServerPropertiesForGeoRestoreArgs) ToServerPropertiesForGeoRestoreOutputWithContext(ctx context.Context) ServerPropertiesForGeoRestoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesForGeoRestoreOutput)
}

type ServerPropertiesForGeoRestoreOutput struct{ *pulumi.OutputState }

func (ServerPropertiesForGeoRestoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForGeoRestore)(nil)).Elem()
}

func (o ServerPropertiesForGeoRestoreOutput) ToServerPropertiesForGeoRestoreOutput() ServerPropertiesForGeoRestoreOutput {
	return o
}

func (o ServerPropertiesForGeoRestoreOutput) ToServerPropertiesForGeoRestoreOutputWithContext(ctx context.Context) ServerPropertiesForGeoRestoreOutput {
	return o
}

func (o ServerPropertiesForGeoRestoreOutput) CreateMode() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForGeoRestore) string { return v.CreateMode }).(pulumi.StringOutput)
}

func (o ServerPropertiesForGeoRestoreOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForGeoRestore) *string { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForGeoRestoreOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForGeoRestore) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForGeoRestoreOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForGeoRestore) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForGeoRestoreOutput) SourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForGeoRestore) string { return v.SourceServerId }).(pulumi.StringOutput)
}

func (o ServerPropertiesForGeoRestoreOutput) SslEnforcement() SslEnforcementEnumPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForGeoRestore) *SslEnforcementEnum { return v.SslEnforcement }).(SslEnforcementEnumPtrOutput)
}

func (o ServerPropertiesForGeoRestoreOutput) StorageProfile() StorageProfilePtrOutput {
	return o.ApplyT(func(v ServerPropertiesForGeoRestore) *StorageProfile { return v.StorageProfile }).(StorageProfilePtrOutput)
}

func (o ServerPropertiesForGeoRestoreOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForGeoRestore) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ServerPropertiesForReplica struct {
	CreateMode               string              `pulumi:"createMode"`
	InfrastructureEncryption *string             `pulumi:"infrastructureEncryption"`
	MinimalTlsVersion        *string             `pulumi:"minimalTlsVersion"`
	PublicNetworkAccess      *string             `pulumi:"publicNetworkAccess"`
	SourceServerId           string              `pulumi:"sourceServerId"`
	SslEnforcement           *SslEnforcementEnum `pulumi:"sslEnforcement"`
	StorageProfile           *StorageProfile     `pulumi:"storageProfile"`
	Version                  *string             `pulumi:"version"`
}





type ServerPropertiesForReplicaInput interface {
	pulumi.Input

	ToServerPropertiesForReplicaOutput() ServerPropertiesForReplicaOutput
	ToServerPropertiesForReplicaOutputWithContext(context.Context) ServerPropertiesForReplicaOutput
}

type ServerPropertiesForReplicaArgs struct {
	CreateMode               pulumi.StringInput         `pulumi:"createMode"`
	InfrastructureEncryption pulumi.StringPtrInput      `pulumi:"infrastructureEncryption"`
	MinimalTlsVersion        pulumi.StringPtrInput      `pulumi:"minimalTlsVersion"`
	PublicNetworkAccess      pulumi.StringPtrInput      `pulumi:"publicNetworkAccess"`
	SourceServerId           pulumi.StringInput         `pulumi:"sourceServerId"`
	SslEnforcement           SslEnforcementEnumPtrInput `pulumi:"sslEnforcement"`
	StorageProfile           StorageProfilePtrInput     `pulumi:"storageProfile"`
	Version                  pulumi.StringPtrInput      `pulumi:"version"`
}

func (ServerPropertiesForReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForReplica)(nil)).Elem()
}

func (i ServerPropertiesForReplicaArgs) ToServerPropertiesForReplicaOutput() ServerPropertiesForReplicaOutput {
	return i.ToServerPropertiesForReplicaOutputWithContext(context.Background())
}

func (i ServerPropertiesForReplicaArgs) ToServerPropertiesForReplicaOutputWithContext(ctx context.Context) ServerPropertiesForReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesForReplicaOutput)
}

type ServerPropertiesForReplicaOutput struct{ *pulumi.OutputState }

func (ServerPropertiesForReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForReplica)(nil)).Elem()
}

func (o ServerPropertiesForReplicaOutput) ToServerPropertiesForReplicaOutput() ServerPropertiesForReplicaOutput {
	return o
}

func (o ServerPropertiesForReplicaOutput) ToServerPropertiesForReplicaOutputWithContext(ctx context.Context) ServerPropertiesForReplicaOutput {
	return o
}

func (o ServerPropertiesForReplicaOutput) CreateMode() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForReplica) string { return v.CreateMode }).(pulumi.StringOutput)
}

func (o ServerPropertiesForReplicaOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForReplica) *string { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForReplicaOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForReplica) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForReplicaOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForReplica) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForReplicaOutput) SourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForReplica) string { return v.SourceServerId }).(pulumi.StringOutput)
}

func (o ServerPropertiesForReplicaOutput) SslEnforcement() SslEnforcementEnumPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForReplica) *SslEnforcementEnum { return v.SslEnforcement }).(SslEnforcementEnumPtrOutput)
}

func (o ServerPropertiesForReplicaOutput) StorageProfile() StorageProfilePtrOutput {
	return o.ApplyT(func(v ServerPropertiesForReplica) *StorageProfile { return v.StorageProfile }).(StorageProfilePtrOutput)
}

func (o ServerPropertiesForReplicaOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForReplica) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ServerPropertiesForRestore struct {
	CreateMode               string              `pulumi:"createMode"`
	InfrastructureEncryption *string             `pulumi:"infrastructureEncryption"`
	MinimalTlsVersion        *string             `pulumi:"minimalTlsVersion"`
	PublicNetworkAccess      *string             `pulumi:"publicNetworkAccess"`
	RestorePointInTime       string              `pulumi:"restorePointInTime"`
	SourceServerId           string              `pulumi:"sourceServerId"`
	SslEnforcement           *SslEnforcementEnum `pulumi:"sslEnforcement"`
	StorageProfile           *StorageProfile     `pulumi:"storageProfile"`
	Version                  *string             `pulumi:"version"`
}





type ServerPropertiesForRestoreInput interface {
	pulumi.Input

	ToServerPropertiesForRestoreOutput() ServerPropertiesForRestoreOutput
	ToServerPropertiesForRestoreOutputWithContext(context.Context) ServerPropertiesForRestoreOutput
}

type ServerPropertiesForRestoreArgs struct {
	CreateMode               pulumi.StringInput         `pulumi:"createMode"`
	InfrastructureEncryption pulumi.StringPtrInput      `pulumi:"infrastructureEncryption"`
	MinimalTlsVersion        pulumi.StringPtrInput      `pulumi:"minimalTlsVersion"`
	PublicNetworkAccess      pulumi.StringPtrInput      `pulumi:"publicNetworkAccess"`
	RestorePointInTime       pulumi.StringInput         `pulumi:"restorePointInTime"`
	SourceServerId           pulumi.StringInput         `pulumi:"sourceServerId"`
	SslEnforcement           SslEnforcementEnumPtrInput `pulumi:"sslEnforcement"`
	StorageProfile           StorageProfilePtrInput     `pulumi:"storageProfile"`
	Version                  pulumi.StringPtrInput      `pulumi:"version"`
}

func (ServerPropertiesForRestoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForRestore)(nil)).Elem()
}

func (i ServerPropertiesForRestoreArgs) ToServerPropertiesForRestoreOutput() ServerPropertiesForRestoreOutput {
	return i.ToServerPropertiesForRestoreOutputWithContext(context.Background())
}

func (i ServerPropertiesForRestoreArgs) ToServerPropertiesForRestoreOutputWithContext(ctx context.Context) ServerPropertiesForRestoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerPropertiesForRestoreOutput)
}

type ServerPropertiesForRestoreOutput struct{ *pulumi.OutputState }

func (ServerPropertiesForRestoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerPropertiesForRestore)(nil)).Elem()
}

func (o ServerPropertiesForRestoreOutput) ToServerPropertiesForRestoreOutput() ServerPropertiesForRestoreOutput {
	return o
}

func (o ServerPropertiesForRestoreOutput) ToServerPropertiesForRestoreOutputWithContext(ctx context.Context) ServerPropertiesForRestoreOutput {
	return o
}

func (o ServerPropertiesForRestoreOutput) CreateMode() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) string { return v.CreateMode }).(pulumi.StringOutput)
}

func (o ServerPropertiesForRestoreOutput) InfrastructureEncryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) *string { return v.InfrastructureEncryption }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForRestoreOutput) MinimalTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) *string { return v.MinimalTlsVersion }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForRestoreOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ServerPropertiesForRestoreOutput) RestorePointInTime() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) string { return v.RestorePointInTime }).(pulumi.StringOutput)
}

func (o ServerPropertiesForRestoreOutput) SourceServerId() pulumi.StringOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) string { return v.SourceServerId }).(pulumi.StringOutput)
}

func (o ServerPropertiesForRestoreOutput) SslEnforcement() SslEnforcementEnumPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) *SslEnforcementEnum { return v.SslEnforcement }).(SslEnforcementEnumPtrOutput)
}

func (o ServerPropertiesForRestoreOutput) StorageProfile() StorageProfilePtrOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) *StorageProfile { return v.StorageProfile }).(StorageProfilePtrOutput)
}

func (o ServerPropertiesForRestoreOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerPropertiesForRestore) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
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
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
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

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
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

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
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
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
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

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type StorageProfile struct {
	BackupRetentionDays *int    `pulumi:"backupRetentionDays"`
	GeoRedundantBackup  *string `pulumi:"geoRedundantBackup"`
	StorageAutogrow     *string `pulumi:"storageAutogrow"`
	StorageMB           *int    `pulumi:"storageMB"`
}





type StorageProfileInput interface {
	pulumi.Input

	ToStorageProfileOutput() StorageProfileOutput
	ToStorageProfileOutputWithContext(context.Context) StorageProfileOutput
}

type StorageProfileArgs struct {
	BackupRetentionDays pulumi.IntPtrInput    `pulumi:"backupRetentionDays"`
	GeoRedundantBackup  pulumi.StringPtrInput `pulumi:"geoRedundantBackup"`
	StorageAutogrow     pulumi.StringPtrInput `pulumi:"storageAutogrow"`
	StorageMB           pulumi.IntPtrInput    `pulumi:"storageMB"`
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

func (o StorageProfileOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfile) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o StorageProfileOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.GeoRedundantBackup }).(pulumi.StringPtrOutput)
}

func (o StorageProfileOutput) StorageAutogrow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfile) *string { return v.StorageAutogrow }).(pulumi.StringPtrOutput)
}

func (o StorageProfileOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfile) *int { return v.StorageMB }).(pulumi.IntPtrOutput)
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

func (o StorageProfilePtrOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *int {
		if v == nil {
			return nil
		}
		return v.BackupRetentionDays
	}).(pulumi.IntPtrOutput)
}

func (o StorageProfilePtrOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.GeoRedundantBackup
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfilePtrOutput) StorageAutogrow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *string {
		if v == nil {
			return nil
		}
		return v.StorageAutogrow
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfilePtrOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfile) *int {
		if v == nil {
			return nil
		}
		return v.StorageMB
	}).(pulumi.IntPtrOutput)
}

type StorageProfileResponse struct {
	BackupRetentionDays *int    `pulumi:"backupRetentionDays"`
	GeoRedundantBackup  *string `pulumi:"geoRedundantBackup"`
	StorageAutogrow     *string `pulumi:"storageAutogrow"`
	StorageMB           *int    `pulumi:"storageMB"`
}





type StorageProfileResponseInput interface {
	pulumi.Input

	ToStorageProfileResponseOutput() StorageProfileResponseOutput
	ToStorageProfileResponseOutputWithContext(context.Context) StorageProfileResponseOutput
}

type StorageProfileResponseArgs struct {
	BackupRetentionDays pulumi.IntPtrInput    `pulumi:"backupRetentionDays"`
	GeoRedundantBackup  pulumi.StringPtrInput `pulumi:"geoRedundantBackup"`
	StorageAutogrow     pulumi.StringPtrInput `pulumi:"storageAutogrow"`
	StorageMB           pulumi.IntPtrInput    `pulumi:"storageMB"`
}

func (StorageProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageProfileResponse)(nil)).Elem()
}

func (i StorageProfileResponseArgs) ToStorageProfileResponseOutput() StorageProfileResponseOutput {
	return i.ToStorageProfileResponseOutputWithContext(context.Background())
}

func (i StorageProfileResponseArgs) ToStorageProfileResponseOutputWithContext(ctx context.Context) StorageProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponseOutput)
}

func (i StorageProfileResponseArgs) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return i.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (i StorageProfileResponseArgs) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponseOutput).ToStorageProfileResponsePtrOutputWithContext(ctx)
}









type StorageProfileResponsePtrInput interface {
	pulumi.Input

	ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput
	ToStorageProfileResponsePtrOutputWithContext(context.Context) StorageProfileResponsePtrOutput
}

type storageProfileResponsePtrType StorageProfileResponseArgs

func StorageProfileResponsePtr(v *StorageProfileResponseArgs) StorageProfileResponsePtrInput {
	return (*storageProfileResponsePtrType)(v)
}

func (*storageProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageProfileResponse)(nil)).Elem()
}

func (i *storageProfileResponsePtrType) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return i.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (i *storageProfileResponsePtrType) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageProfileResponsePtrOutput)
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

func (o StorageProfileResponseOutput) ToStorageProfileResponsePtrOutput() StorageProfileResponsePtrOutput {
	return o.ToStorageProfileResponsePtrOutputWithContext(context.Background())
}

func (o StorageProfileResponseOutput) ToStorageProfileResponsePtrOutputWithContext(ctx context.Context) StorageProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageProfileResponse) *StorageProfileResponse {
		return &v
	}).(StorageProfileResponsePtrOutput)
}

func (o StorageProfileResponseOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *int { return v.BackupRetentionDays }).(pulumi.IntPtrOutput)
}

func (o StorageProfileResponseOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.GeoRedundantBackup }).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponseOutput) StorageAutogrow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *string { return v.StorageAutogrow }).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponseOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StorageProfileResponse) *int { return v.StorageMB }).(pulumi.IntPtrOutput)
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

func (o StorageProfileResponsePtrOutput) BackupRetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.BackupRetentionDays
	}).(pulumi.IntPtrOutput)
}

func (o StorageProfileResponsePtrOutput) GeoRedundantBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.GeoRedundantBackup
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponsePtrOutput) StorageAutogrow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAutogrow
	}).(pulumi.StringPtrOutput)
}

func (o StorageProfileResponsePtrOutput) StorageMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *StorageProfileResponse) *int {
		if v == nil {
			return nil
		}
		return v.StorageMB
	}).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerPrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServerPrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(ServerPrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(ServerPrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(ServerPrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerPropertiesForDefaultCreateOutput{})
	pulumi.RegisterOutputType(ServerPropertiesForGeoRestoreOutput{})
	pulumi.RegisterOutputType(ServerPropertiesForReplicaOutput{})
	pulumi.RegisterOutputType(ServerPropertiesForRestoreOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileOutput{})
	pulumi.RegisterOutputType(StorageProfilePtrOutput{})
	pulumi.RegisterOutputType(StorageProfileResponseOutput{})
	pulumi.RegisterOutputType(StorageProfileResponsePtrOutput{})
}
