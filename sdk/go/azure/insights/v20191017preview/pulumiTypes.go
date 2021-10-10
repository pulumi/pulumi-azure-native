


package v20191017preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionResponse struct {
	Id                                string                                             `pulumi:"id"`
	Name                              string                                             `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                             `pulumi:"provisioningState"`
	Type                              string                                             `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                                        `pulumi:"id"`
	Name                              pulumi.StringInput                                        `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointPropertyResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                        `pulumi:"provisioningState"`
	Type                              pulumi.StringInput                                        `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointPropertyResponse { return v.PrivateEndpoint }).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
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

type PrivateEndpointProperty struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput
	ToPrivateEndpointPropertyOutputWithContext(context.Context) PrivateEndpointPropertyOutput
}

type PrivateEndpointPropertyArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return i.ToPrivateEndpointPropertyOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput)
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput).ToPrivateEndpointPropertyPtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput
	ToPrivateEndpointPropertyPtrOutputWithContext(context.Context) PrivateEndpointPropertyPtrOutput
}

type privateEndpointPropertyPtrType PrivateEndpointPropertyArgs

func PrivateEndpointPropertyPtr(v *PrivateEndpointPropertyArgs) PrivateEndpointPropertyPtrInput {
	return (*privateEndpointPropertyPtrType)(v)
}

func (*privateEndpointPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyPtrOutput)
}

type PrivateEndpointPropertyOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointProperty) *PrivateEndpointProperty {
		return &v
	}).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointPropertyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointProperty) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) Elem() PrivateEndpointPropertyOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) PrivateEndpointProperty {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointProperty
		return ret
	}).(PrivateEndpointPropertyOutput)
}

func (o PrivateEndpointPropertyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

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

type PrivateLinkServiceConnectionStateProperty struct {
	Description string `pulumi:"description"`
	Status      string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput
	ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyOutput
}

type PrivateLinkServiceConnectionStatePropertyArgs struct {
	Description pulumi.StringInput `pulumi:"description"`
	Status      pulumi.StringInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput).ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput
	ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput
}

type privateLinkServiceConnectionStatePropertyPtrType PrivateLinkServiceConnectionStatePropertyArgs

func PrivateLinkServiceConnectionStatePropertyPtr(v *PrivateLinkServiceConnectionStatePropertyArgs) PrivateLinkServiceConnectionStatePropertyPtrInput {
	return (*privateLinkServiceConnectionStatePropertyPtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateProperty) *PrivateLinkServiceConnectionStateProperty {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) PrivateLinkServiceConnectionStateProperty {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateProperty
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput
	ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput
}

type PrivateLinkServiceConnectionStatePropertyResponseArgs struct {
	ActionsRequired pulumi.StringInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput `pulumi:"description"`
	Status          pulumi.StringInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput).ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
}

type privateLinkServiceConnectionStatePropertyResponsePtrType PrivateLinkServiceConnectionStatePropertyResponseArgs

func PrivateLinkServiceConnectionStatePropertyResponsePtr(v *PrivateLinkServiceConnectionStatePropertyResponseArgs) PrivateLinkServiceConnectionStatePropertyResponsePtrInput {
	return (*privateLinkServiceConnectionStatePropertyResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStatePropertyResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) PrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type WorkbookTemplateGallery struct {
	Category     *string `pulumi:"category"`
	Name         *string `pulumi:"name"`
	Order        *int    `pulumi:"order"`
	ResourceType *string `pulumi:"resourceType"`
	Type         *string `pulumi:"type"`
}





type WorkbookTemplateGalleryInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryOutput() WorkbookTemplateGalleryOutput
	ToWorkbookTemplateGalleryOutputWithContext(context.Context) WorkbookTemplateGalleryOutput
}

type WorkbookTemplateGalleryArgs struct {
	Category     pulumi.StringPtrInput `pulumi:"category"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Order        pulumi.IntPtrInput    `pulumi:"order"`
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (WorkbookTemplateGalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGallery)(nil)).Elem()
}

func (i WorkbookTemplateGalleryArgs) ToWorkbookTemplateGalleryOutput() WorkbookTemplateGalleryOutput {
	return i.ToWorkbookTemplateGalleryOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryArgs) ToWorkbookTemplateGalleryOutputWithContext(ctx context.Context) WorkbookTemplateGalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryOutput)
}





type WorkbookTemplateGalleryArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryArrayOutput() WorkbookTemplateGalleryArrayOutput
	ToWorkbookTemplateGalleryArrayOutputWithContext(context.Context) WorkbookTemplateGalleryArrayOutput
}

type WorkbookTemplateGalleryArray []WorkbookTemplateGalleryInput

func (WorkbookTemplateGalleryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGallery)(nil)).Elem()
}

func (i WorkbookTemplateGalleryArray) ToWorkbookTemplateGalleryArrayOutput() WorkbookTemplateGalleryArrayOutput {
	return i.ToWorkbookTemplateGalleryArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryArray) ToWorkbookTemplateGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryArrayOutput)
}

type WorkbookTemplateGalleryOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGallery)(nil)).Elem()
}

func (o WorkbookTemplateGalleryOutput) ToWorkbookTemplateGalleryOutput() WorkbookTemplateGalleryOutput {
	return o
}

func (o WorkbookTemplateGalleryOutput) ToWorkbookTemplateGalleryOutputWithContext(ctx context.Context) WorkbookTemplateGalleryOutput {
	return o
}

func (o WorkbookTemplateGalleryOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGallery) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkbookTemplateGalleryArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGallery)(nil)).Elem()
}

func (o WorkbookTemplateGalleryArrayOutput) ToWorkbookTemplateGalleryArrayOutput() WorkbookTemplateGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryArrayOutput) ToWorkbookTemplateGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateGalleryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateGallery {
		return vs[0].([]WorkbookTemplateGallery)[vs[1].(int)]
	}).(WorkbookTemplateGalleryOutput)
}

type WorkbookTemplateGalleryResponse struct {
	Category     *string `pulumi:"category"`
	Name         *string `pulumi:"name"`
	Order        *int    `pulumi:"order"`
	ResourceType *string `pulumi:"resourceType"`
	Type         *string `pulumi:"type"`
}





type WorkbookTemplateGalleryResponseInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryResponseOutput() WorkbookTemplateGalleryResponseOutput
	ToWorkbookTemplateGalleryResponseOutputWithContext(context.Context) WorkbookTemplateGalleryResponseOutput
}

type WorkbookTemplateGalleryResponseArgs struct {
	Category     pulumi.StringPtrInput `pulumi:"category"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Order        pulumi.IntPtrInput    `pulumi:"order"`
	ResourceType pulumi.StringPtrInput `pulumi:"resourceType"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (WorkbookTemplateGalleryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateGalleryResponseArgs) ToWorkbookTemplateGalleryResponseOutput() WorkbookTemplateGalleryResponseOutput {
	return i.ToWorkbookTemplateGalleryResponseOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryResponseArgs) ToWorkbookTemplateGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryResponseOutput)
}





type WorkbookTemplateGalleryResponseArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateGalleryResponseArrayOutput() WorkbookTemplateGalleryResponseArrayOutput
	ToWorkbookTemplateGalleryResponseArrayOutputWithContext(context.Context) WorkbookTemplateGalleryResponseArrayOutput
}

type WorkbookTemplateGalleryResponseArray []WorkbookTemplateGalleryResponseInput

func (WorkbookTemplateGalleryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateGalleryResponseArray) ToWorkbookTemplateGalleryResponseArrayOutput() WorkbookTemplateGalleryResponseArrayOutput {
	return i.ToWorkbookTemplateGalleryResponseArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateGalleryResponseArray) ToWorkbookTemplateGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateGalleryResponseArrayOutput)
}

type WorkbookTemplateGalleryResponseOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateGalleryResponseOutput) ToWorkbookTemplateGalleryResponseOutput() WorkbookTemplateGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseOutput) ToWorkbookTemplateGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.ResourceType }).(pulumi.StringPtrOutput)
}

func (o WorkbookTemplateGalleryResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkbookTemplateGalleryResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type WorkbookTemplateGalleryResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateGalleryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateGalleryResponseArrayOutput) ToWorkbookTemplateGalleryResponseArrayOutput() WorkbookTemplateGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseArrayOutput) ToWorkbookTemplateGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateGalleryResponseArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateGalleryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateGalleryResponse {
		return vs[0].([]WorkbookTemplateGalleryResponse)[vs[1].(int)]
	}).(WorkbookTemplateGalleryResponseOutput)
}

type WorkbookTemplateLocalizedGallery struct {
	Galleries    []WorkbookTemplateGallery `pulumi:"galleries"`
	TemplateData interface{}               `pulumi:"templateData"`
}





type WorkbookTemplateLocalizedGalleryInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryOutput() WorkbookTemplateLocalizedGalleryOutput
	ToWorkbookTemplateLocalizedGalleryOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryOutput
}

type WorkbookTemplateLocalizedGalleryArgs struct {
	Galleries    WorkbookTemplateGalleryArrayInput `pulumi:"galleries"`
	TemplateData pulumi.Input                      `pulumi:"templateData"`
}

func (WorkbookTemplateLocalizedGalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryArgs) ToWorkbookTemplateLocalizedGalleryOutput() WorkbookTemplateLocalizedGalleryOutput {
	return i.ToWorkbookTemplateLocalizedGalleryOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryArgs) ToWorkbookTemplateLocalizedGalleryOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryOutput)
}





type WorkbookTemplateLocalizedGalleryArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryArrayOutput() WorkbookTemplateLocalizedGalleryArrayOutput
	ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryArrayOutput
}

type WorkbookTemplateLocalizedGalleryArray []WorkbookTemplateLocalizedGalleryInput

func (WorkbookTemplateLocalizedGalleryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryArray) ToWorkbookTemplateLocalizedGalleryArrayOutput() WorkbookTemplateLocalizedGalleryArrayOutput {
	return i.ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryArray) ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryArrayOutput)
}

type WorkbookTemplateLocalizedGalleryOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryOutput) ToWorkbookTemplateLocalizedGalleryOutput() WorkbookTemplateLocalizedGalleryOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryOutput) ToWorkbookTemplateLocalizedGalleryOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryOutput) Galleries() WorkbookTemplateGalleryArrayOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGallery) []WorkbookTemplateGallery { return v.Galleries }).(WorkbookTemplateGalleryArrayOutput)
}

func (o WorkbookTemplateLocalizedGalleryOutput) TemplateData() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGallery) interface{} { return v.TemplateData }).(pulumi.AnyOutput)
}

type WorkbookTemplateLocalizedGalleryArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGallery)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryArrayOutput) ToWorkbookTemplateLocalizedGalleryArrayOutput() WorkbookTemplateLocalizedGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayOutput) ToWorkbookTemplateLocalizedGalleryArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateLocalizedGalleryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGallery {
		return vs[0].([]WorkbookTemplateLocalizedGallery)[vs[1].(int)]
	}).(WorkbookTemplateLocalizedGalleryOutput)
}

type WorkbookTemplateLocalizedGalleryResponse struct {
	Galleries    []WorkbookTemplateGalleryResponse `pulumi:"galleries"`
	TemplateData interface{}                       `pulumi:"templateData"`
}





type WorkbookTemplateLocalizedGalleryResponseInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryResponseOutput() WorkbookTemplateLocalizedGalleryResponseOutput
	ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryResponseOutput
}

type WorkbookTemplateLocalizedGalleryResponseArgs struct {
	Galleries    WorkbookTemplateGalleryResponseArrayInput `pulumi:"galleries"`
	TemplateData pulumi.Input                              `pulumi:"templateData"`
}

func (WorkbookTemplateLocalizedGalleryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryResponseArgs) ToWorkbookTemplateLocalizedGalleryResponseOutput() WorkbookTemplateLocalizedGalleryResponseOutput {
	return i.ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryResponseArgs) ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryResponseOutput)
}





type WorkbookTemplateLocalizedGalleryResponseArrayInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryResponseArrayOutput() WorkbookTemplateLocalizedGalleryResponseArrayOutput
	ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryResponseArrayOutput
}

type WorkbookTemplateLocalizedGalleryResponseArray []WorkbookTemplateLocalizedGalleryResponseInput

func (WorkbookTemplateLocalizedGalleryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryResponseArray) ToWorkbookTemplateLocalizedGalleryResponseArrayOutput() WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return i.ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryResponseArray) ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryResponseArrayOutput)
}

type WorkbookTemplateLocalizedGalleryResponseOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) ToWorkbookTemplateLocalizedGalleryResponseOutput() WorkbookTemplateLocalizedGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) ToWorkbookTemplateLocalizedGalleryResponseOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) Galleries() WorkbookTemplateGalleryResponseArrayOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGalleryResponse) []WorkbookTemplateGalleryResponse { return v.Galleries }).(WorkbookTemplateGalleryResponseArrayOutput)
}

func (o WorkbookTemplateLocalizedGalleryResponseOutput) TemplateData() pulumi.AnyOutput {
	return o.ApplyT(func(v WorkbookTemplateLocalizedGalleryResponse) interface{} { return v.TemplateData }).(pulumi.AnyOutput)
}

type WorkbookTemplateLocalizedGalleryResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkbookTemplateLocalizedGalleryResponse)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayOutput() WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayOutput) Index(i pulumi.IntInput) WorkbookTemplateLocalizedGalleryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGalleryResponse {
		return vs[0].([]WorkbookTemplateLocalizedGalleryResponse)[vs[1].(int)]
	}).(WorkbookTemplateLocalizedGalleryResponseOutput)
}

type WorkbookTemplateLocalizedGalleryArrayMap map[string]WorkbookTemplateLocalizedGalleryArrayInput

func (WorkbookTemplateLocalizedGalleryArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryArray)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryArrayMap) ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return i.ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryArrayMap) ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryArrayMapOutput)
}

type WorkbookTemplateLocalizedGalleryArrayMapOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkbookTemplateLocalizedGalleryArray)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryArrayMapOutput) MapIndex(k pulumi.StringInput) WorkbookTemplateLocalizedGalleryArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGalleryArray {
		return vs[0].(map[string]WorkbookTemplateLocalizedGalleryArray)[vs[1].(string)]
	}).(WorkbookTemplateLocalizedGalleryArrayOutput)
}





type WorkbookTemplateLocalizedGalleryArrayMapInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryArrayMapOutput() WorkbookTemplateLocalizedGalleryArrayMapOutput
	ToWorkbookTemplateLocalizedGalleryArrayMapOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryArrayMapOutput
}

type WorkbookTemplateLocalizedGalleryResponseArrayMap map[string]WorkbookTemplateLocalizedGalleryResponseArrayInput

func (WorkbookTemplateLocalizedGalleryResponseArrayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponseArray)(nil)).Elem()
}

func (i WorkbookTemplateLocalizedGalleryResponseArrayMap) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutput() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return i.ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(context.Background())
}

func (i WorkbookTemplateLocalizedGalleryResponseArrayMap) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkbookTemplateLocalizedGalleryResponseArrayMapOutput)
}

type WorkbookTemplateLocalizedGalleryResponseArrayMapOutput struct{ *pulumi.OutputState }

func (WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]WorkbookTemplateLocalizedGalleryResponseArray)(nil)).Elem()
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutput() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(ctx context.Context) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput {
	return o
}

func (o WorkbookTemplateLocalizedGalleryResponseArrayMapOutput) MapIndex(k pulumi.StringInput) WorkbookTemplateLocalizedGalleryResponseArrayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) WorkbookTemplateLocalizedGalleryResponseArray {
		return vs[0].(map[string]WorkbookTemplateLocalizedGalleryResponseArray)[vs[1].(string)]
	}).(WorkbookTemplateLocalizedGalleryResponseArrayOutput)
}





type WorkbookTemplateLocalizedGalleryResponseArrayMapInput interface {
	pulumi.Input

	ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutput() WorkbookTemplateLocalizedGalleryResponseArrayMapOutput
	ToWorkbookTemplateLocalizedGalleryResponseArrayMapOutputWithContext(context.Context) WorkbookTemplateLocalizedGalleryResponseArrayMapOutput
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseInput)(nil)).Elem(), PrivateEndpointConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionResponseArrayInput)(nil)).Elem(), PrivateEndpointConnectionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointPropertyInput)(nil)).Elem(), PrivateEndpointPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointPropertyPtrInput)(nil)).Elem(), PrivateEndpointPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointPropertyResponseInput)(nil)).Elem(), PrivateEndpointPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointPropertyResponsePtrInput)(nil)).Elem(), PrivateEndpointPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyInput)(nil)).Elem(), PrivateLinkServiceConnectionStatePropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyPtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStatePropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponseInput)(nil)).Elem(), PrivateLinkServiceConnectionStatePropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponsePtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStatePropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateGalleryInput)(nil)).Elem(), WorkbookTemplateGalleryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateGalleryArrayInput)(nil)).Elem(), WorkbookTemplateGalleryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateGalleryResponseInput)(nil)).Elem(), WorkbookTemplateGalleryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateGalleryResponseArrayInput)(nil)).Elem(), WorkbookTemplateGalleryResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryArrayInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponseInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponseArrayInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryArrayMapInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryArrayMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkbookTemplateLocalizedGalleryResponseArrayMapInput)(nil)).Elem(), WorkbookTemplateLocalizedGalleryResponseArrayMap{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryResponseOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateGalleryResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryResponseOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryArrayMapOutput{})
	pulumi.RegisterOutputType(WorkbookTemplateLocalizedGalleryResponseArrayMapOutput{})
}
