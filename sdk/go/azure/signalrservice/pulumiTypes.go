


package signalrservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkACL struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
}





type NetworkACLInput interface {
	pulumi.Input

	ToNetworkACLOutput() NetworkACLOutput
	ToNetworkACLOutputWithContext(context.Context) NetworkACLOutput
}

type NetworkACLArgs struct {
	Allow pulumi.StringArrayInput `pulumi:"allow"`
	Deny  pulumi.StringArrayInput `pulumi:"deny"`
}

func (NetworkACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkACL)(nil)).Elem()
}

func (i NetworkACLArgs) ToNetworkACLOutput() NetworkACLOutput {
	return i.ToNetworkACLOutputWithContext(context.Background())
}

func (i NetworkACLArgs) ToNetworkACLOutputWithContext(ctx context.Context) NetworkACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLOutput)
}

func (i NetworkACLArgs) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return i.ToNetworkACLPtrOutputWithContext(context.Background())
}

func (i NetworkACLArgs) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLOutput).ToNetworkACLPtrOutputWithContext(ctx)
}









type NetworkACLPtrInput interface {
	pulumi.Input

	ToNetworkACLPtrOutput() NetworkACLPtrOutput
	ToNetworkACLPtrOutputWithContext(context.Context) NetworkACLPtrOutput
}

type networkACLPtrType NetworkACLArgs

func NetworkACLPtr(v *NetworkACLArgs) NetworkACLPtrInput {
	return (*networkACLPtrType)(v)
}

func (*networkACLPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkACL)(nil)).Elem()
}

func (i *networkACLPtrType) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return i.ToNetworkACLPtrOutputWithContext(context.Background())
}

func (i *networkACLPtrType) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkACLPtrOutput)
}

type NetworkACLOutput struct{ *pulumi.OutputState }

func (NetworkACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkACL)(nil)).Elem()
}

func (o NetworkACLOutput) ToNetworkACLOutput() NetworkACLOutput {
	return o
}

func (o NetworkACLOutput) ToNetworkACLOutputWithContext(ctx context.Context) NetworkACLOutput {
	return o
}

func (o NetworkACLOutput) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return o.ToNetworkACLPtrOutputWithContext(context.Background())
}

func (o NetworkACLOutput) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkACL) *NetworkACL {
		return &v
	}).(NetworkACLPtrOutput)
}

func (o NetworkACLOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACL) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o NetworkACLOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACL) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

type NetworkACLPtrOutput struct{ *pulumi.OutputState }

func (NetworkACLPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkACL)(nil)).Elem()
}

func (o NetworkACLPtrOutput) ToNetworkACLPtrOutput() NetworkACLPtrOutput {
	return o
}

func (o NetworkACLPtrOutput) ToNetworkACLPtrOutputWithContext(ctx context.Context) NetworkACLPtrOutput {
	return o
}

func (o NetworkACLPtrOutput) Elem() NetworkACLOutput {
	return o.ApplyT(func(v *NetworkACL) NetworkACL {
		if v != nil {
			return *v
		}
		var ret NetworkACL
		return ret
	}).(NetworkACLOutput)
}

func (o NetworkACLPtrOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACL) []string {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(pulumi.StringArrayOutput)
}

func (o NetworkACLPtrOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACL) []string {
		if v == nil {
			return nil
		}
		return v.Deny
	}).(pulumi.StringArrayOutput)
}

type NetworkACLResponse struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
}

type NetworkACLResponseOutput struct{ *pulumi.OutputState }

func (NetworkACLResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkACLResponse)(nil)).Elem()
}

func (o NetworkACLResponseOutput) ToNetworkACLResponseOutput() NetworkACLResponseOutput {
	return o
}

func (o NetworkACLResponseOutput) ToNetworkACLResponseOutputWithContext(ctx context.Context) NetworkACLResponseOutput {
	return o
}

func (o NetworkACLResponseOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACLResponse) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o NetworkACLResponseOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkACLResponse) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

type NetworkACLResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkACLResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkACLResponse)(nil)).Elem()
}

func (o NetworkACLResponsePtrOutput) ToNetworkACLResponsePtrOutput() NetworkACLResponsePtrOutput {
	return o
}

func (o NetworkACLResponsePtrOutput) ToNetworkACLResponsePtrOutputWithContext(ctx context.Context) NetworkACLResponsePtrOutput {
	return o
}

func (o NetworkACLResponsePtrOutput) Elem() NetworkACLResponseOutput {
	return o.ApplyT(func(v *NetworkACLResponse) NetworkACLResponse {
		if v != nil {
			return *v
		}
		var ret NetworkACLResponse
		return ret
	}).(NetworkACLResponseOutput)
}

func (o NetworkACLResponsePtrOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACLResponse) []string {
		if v == nil {
			return nil
		}
		return v.Allow
	}).(pulumi.StringArrayOutput)
}

func (o NetworkACLResponsePtrOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkACLResponse) []string {
		if v == nil {
			return nil
		}
		return v.Deny
	}).(pulumi.StringArrayOutput)
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointACL struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
	Name  string   `pulumi:"name"`
}





type PrivateEndpointACLInput interface {
	pulumi.Input

	ToPrivateEndpointACLOutput() PrivateEndpointACLOutput
	ToPrivateEndpointACLOutputWithContext(context.Context) PrivateEndpointACLOutput
}

type PrivateEndpointACLArgs struct {
	Allow pulumi.StringArrayInput `pulumi:"allow"`
	Deny  pulumi.StringArrayInput `pulumi:"deny"`
	Name  pulumi.StringInput      `pulumi:"name"`
}

func (PrivateEndpointACLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointACL)(nil)).Elem()
}

func (i PrivateEndpointACLArgs) ToPrivateEndpointACLOutput() PrivateEndpointACLOutput {
	return i.ToPrivateEndpointACLOutputWithContext(context.Background())
}

func (i PrivateEndpointACLArgs) ToPrivateEndpointACLOutputWithContext(ctx context.Context) PrivateEndpointACLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointACLOutput)
}





type PrivateEndpointACLArrayInput interface {
	pulumi.Input

	ToPrivateEndpointACLArrayOutput() PrivateEndpointACLArrayOutput
	ToPrivateEndpointACLArrayOutputWithContext(context.Context) PrivateEndpointACLArrayOutput
}

type PrivateEndpointACLArray []PrivateEndpointACLInput

func (PrivateEndpointACLArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointACL)(nil)).Elem()
}

func (i PrivateEndpointACLArray) ToPrivateEndpointACLArrayOutput() PrivateEndpointACLArrayOutput {
	return i.ToPrivateEndpointACLArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointACLArray) ToPrivateEndpointACLArrayOutputWithContext(ctx context.Context) PrivateEndpointACLArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointACLArrayOutput)
}

type PrivateEndpointACLOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointACL)(nil)).Elem()
}

func (o PrivateEndpointACLOutput) ToPrivateEndpointACLOutput() PrivateEndpointACLOutput {
	return o
}

func (o PrivateEndpointACLOutput) ToPrivateEndpointACLOutputWithContext(ctx context.Context) PrivateEndpointACLOutput {
	return o
}

func (o PrivateEndpointACLOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACL) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACL) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointACL) string { return v.Name }).(pulumi.StringOutput)
}

type PrivateEndpointACLArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointACL)(nil)).Elem()
}

func (o PrivateEndpointACLArrayOutput) ToPrivateEndpointACLArrayOutput() PrivateEndpointACLArrayOutput {
	return o
}

func (o PrivateEndpointACLArrayOutput) ToPrivateEndpointACLArrayOutputWithContext(ctx context.Context) PrivateEndpointACLArrayOutput {
	return o
}

func (o PrivateEndpointACLArrayOutput) Index(i pulumi.IntInput) PrivateEndpointACLOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointACL {
		return vs[0].([]PrivateEndpointACL)[vs[1].(int)]
	}).(PrivateEndpointACLOutput)
}

type PrivateEndpointACLResponse struct {
	Allow []string `pulumi:"allow"`
	Deny  []string `pulumi:"deny"`
	Name  string   `pulumi:"name"`
}

type PrivateEndpointACLResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointACLResponse)(nil)).Elem()
}

func (o PrivateEndpointACLResponseOutput) ToPrivateEndpointACLResponseOutput() PrivateEndpointACLResponseOutput {
	return o
}

func (o PrivateEndpointACLResponseOutput) ToPrivateEndpointACLResponseOutputWithContext(ctx context.Context) PrivateEndpointACLResponseOutput {
	return o
}

func (o PrivateEndpointACLResponseOutput) Allow() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACLResponse) []string { return v.Allow }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLResponseOutput) Deny() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointACLResponse) []string { return v.Deny }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointACLResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointACLResponse) string { return v.Name }).(pulumi.StringOutput)
}

type PrivateEndpointACLResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointACLResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointACLResponse)(nil)).Elem()
}

func (o PrivateEndpointACLResponseArrayOutput) ToPrivateEndpointACLResponseArrayOutput() PrivateEndpointACLResponseArrayOutput {
	return o
}

func (o PrivateEndpointACLResponseArrayOutput) ToPrivateEndpointACLResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointACLResponseArrayOutput {
	return o
}

func (o PrivateEndpointACLResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointACLResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointACLResponse {
		return vs[0].([]PrivateEndpointACLResponse)[vs[1].(int)]
	}).(PrivateEndpointACLResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                     `pulumi:"id"`
	Name                              string                                     `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	Type                              string                                     `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
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
	Id *string `pulumi:"id"`
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

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
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
		return v.Id
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

type ResourceReference struct {
	Id *string `pulumi:"id"`
}





type ResourceReferenceInput interface {
	pulumi.Input

	ToResourceReferenceOutput() ResourceReferenceOutput
	ToResourceReferenceOutputWithContext(context.Context) ResourceReferenceOutput
}

type ResourceReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ResourceReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (i ResourceReferenceArgs) ToResourceReferenceOutput() ResourceReferenceOutput {
	return i.ToResourceReferenceOutputWithContext(context.Background())
}

func (i ResourceReferenceArgs) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceReferenceOutput)
}

type ResourceReferenceOutput struct{ *pulumi.OutputState }

func (ResourceReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReference)(nil)).Elem()
}

func (o ResourceReferenceOutput) ToResourceReferenceOutput() ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) ToResourceReferenceOutputWithContext(ctx context.Context) ResourceReferenceOutput {
	return o
}

func (o ResourceReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceReferenceResponse struct {
	Id *string `pulumi:"id"`
}

type ResourceReferenceResponseOutput struct{ *pulumi.OutputState }

func (ResourceReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceReferenceResponse)(nil)).Elem()
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutput() ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) ToResourceReferenceResponseOutputWithContext(ctx context.Context) ResourceReferenceResponseOutput {
	return o
}

func (o ResourceReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ResourceSku struct {
	Capacity *int    `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     *string `pulumi:"tier"`
}





type ResourceSkuInput interface {
	pulumi.Input

	ToResourceSkuOutput() ResourceSkuOutput
	ToResourceSkuOutputWithContext(context.Context) ResourceSkuOutput
}

type ResourceSkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (ResourceSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (i ResourceSkuArgs) ToResourceSkuOutput() ResourceSkuOutput {
	return i.ToResourceSkuOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput)
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i ResourceSkuArgs) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuOutput).ToResourceSkuPtrOutputWithContext(ctx)
}









type ResourceSkuPtrInput interface {
	pulumi.Input

	ToResourceSkuPtrOutput() ResourceSkuPtrOutput
	ToResourceSkuPtrOutputWithContext(context.Context) ResourceSkuPtrOutput
}

type resourceSkuPtrType ResourceSkuArgs

func ResourceSkuPtr(v *ResourceSkuArgs) ResourceSkuPtrInput {
	return (*resourceSkuPtrType)(v)
}

func (*resourceSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return i.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (i *resourceSkuPtrType) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceSkuPtrOutput)
}

type ResourceSkuOutput struct{ *pulumi.OutputState }

func (ResourceSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSku)(nil)).Elem()
}

func (o ResourceSkuOutput) ToResourceSkuOutput() ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuOutputWithContext(ctx context.Context) ResourceSkuOutput {
	return o
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o.ToResourceSkuPtrOutputWithContext(context.Background())
}

func (o ResourceSkuOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceSku) *ResourceSku {
		return &v
	}).(ResourceSkuPtrOutput)
}

func (o ResourceSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSku) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuPtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSku)(nil)).Elem()
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutput() ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) ToResourceSkuPtrOutputWithContext(ctx context.Context) ResourceSkuPtrOutput {
	return o
}

func (o ResourceSkuPtrOutput) Elem() ResourceSkuOutput {
	return o.ApplyT(func(v *ResourceSku) ResourceSku {
		if v != nil {
			return *v
		}
		var ret ResourceSku
		return ret
	}).(ResourceSkuOutput)
}

func (o ResourceSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ResourceSkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   string  `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     string  `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}

type ResourceSkuResponseOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutput() ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) ToResourceSkuResponseOutputWithContext(ctx context.Context) ResourceSkuResponseOutput {
	return o
}

func (o ResourceSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponseOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Family }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceSkuResponse) string { return v.Size }).(pulumi.StringOutput)
}

func (o ResourceSkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceSkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type ResourceSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceSkuResponse)(nil)).Elem()
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutput() ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) ToResourceSkuResponsePtrOutputWithContext(ctx context.Context) ResourceSkuResponsePtrOutput {
	return o
}

func (o ResourceSkuResponsePtrOutput) Elem() ResourceSkuResponseOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) ResourceSkuResponse {
		if v != nil {
			return *v
		}
		var ret ResourceSkuResponse
		return ret
	}).(ResourceSkuResponseOutput)
}

func (o ResourceSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Family
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Size
	}).(pulumi.StringPtrOutput)
}

func (o ResourceSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceSkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type ServerlessUpstreamSettings struct {
	Templates []UpstreamTemplate `pulumi:"templates"`
}





type ServerlessUpstreamSettingsInput interface {
	pulumi.Input

	ToServerlessUpstreamSettingsOutput() ServerlessUpstreamSettingsOutput
	ToServerlessUpstreamSettingsOutputWithContext(context.Context) ServerlessUpstreamSettingsOutput
}

type ServerlessUpstreamSettingsArgs struct {
	Templates UpstreamTemplateArrayInput `pulumi:"templates"`
}

func (ServerlessUpstreamSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerlessUpstreamSettings)(nil)).Elem()
}

func (i ServerlessUpstreamSettingsArgs) ToServerlessUpstreamSettingsOutput() ServerlessUpstreamSettingsOutput {
	return i.ToServerlessUpstreamSettingsOutputWithContext(context.Background())
}

func (i ServerlessUpstreamSettingsArgs) ToServerlessUpstreamSettingsOutputWithContext(ctx context.Context) ServerlessUpstreamSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessUpstreamSettingsOutput)
}

func (i ServerlessUpstreamSettingsArgs) ToServerlessUpstreamSettingsPtrOutput() ServerlessUpstreamSettingsPtrOutput {
	return i.ToServerlessUpstreamSettingsPtrOutputWithContext(context.Background())
}

func (i ServerlessUpstreamSettingsArgs) ToServerlessUpstreamSettingsPtrOutputWithContext(ctx context.Context) ServerlessUpstreamSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessUpstreamSettingsOutput).ToServerlessUpstreamSettingsPtrOutputWithContext(ctx)
}









type ServerlessUpstreamSettingsPtrInput interface {
	pulumi.Input

	ToServerlessUpstreamSettingsPtrOutput() ServerlessUpstreamSettingsPtrOutput
	ToServerlessUpstreamSettingsPtrOutputWithContext(context.Context) ServerlessUpstreamSettingsPtrOutput
}

type serverlessUpstreamSettingsPtrType ServerlessUpstreamSettingsArgs

func ServerlessUpstreamSettingsPtr(v *ServerlessUpstreamSettingsArgs) ServerlessUpstreamSettingsPtrInput {
	return (*serverlessUpstreamSettingsPtrType)(v)
}

func (*serverlessUpstreamSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessUpstreamSettings)(nil)).Elem()
}

func (i *serverlessUpstreamSettingsPtrType) ToServerlessUpstreamSettingsPtrOutput() ServerlessUpstreamSettingsPtrOutput {
	return i.ToServerlessUpstreamSettingsPtrOutputWithContext(context.Background())
}

func (i *serverlessUpstreamSettingsPtrType) ToServerlessUpstreamSettingsPtrOutputWithContext(ctx context.Context) ServerlessUpstreamSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerlessUpstreamSettingsPtrOutput)
}

type ServerlessUpstreamSettingsOutput struct{ *pulumi.OutputState }

func (ServerlessUpstreamSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerlessUpstreamSettings)(nil)).Elem()
}

func (o ServerlessUpstreamSettingsOutput) ToServerlessUpstreamSettingsOutput() ServerlessUpstreamSettingsOutput {
	return o
}

func (o ServerlessUpstreamSettingsOutput) ToServerlessUpstreamSettingsOutputWithContext(ctx context.Context) ServerlessUpstreamSettingsOutput {
	return o
}

func (o ServerlessUpstreamSettingsOutput) ToServerlessUpstreamSettingsPtrOutput() ServerlessUpstreamSettingsPtrOutput {
	return o.ToServerlessUpstreamSettingsPtrOutputWithContext(context.Background())
}

func (o ServerlessUpstreamSettingsOutput) ToServerlessUpstreamSettingsPtrOutputWithContext(ctx context.Context) ServerlessUpstreamSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerlessUpstreamSettings) *ServerlessUpstreamSettings {
		return &v
	}).(ServerlessUpstreamSettingsPtrOutput)
}

func (o ServerlessUpstreamSettingsOutput) Templates() UpstreamTemplateArrayOutput {
	return o.ApplyT(func(v ServerlessUpstreamSettings) []UpstreamTemplate { return v.Templates }).(UpstreamTemplateArrayOutput)
}

type ServerlessUpstreamSettingsPtrOutput struct{ *pulumi.OutputState }

func (ServerlessUpstreamSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessUpstreamSettings)(nil)).Elem()
}

func (o ServerlessUpstreamSettingsPtrOutput) ToServerlessUpstreamSettingsPtrOutput() ServerlessUpstreamSettingsPtrOutput {
	return o
}

func (o ServerlessUpstreamSettingsPtrOutput) ToServerlessUpstreamSettingsPtrOutputWithContext(ctx context.Context) ServerlessUpstreamSettingsPtrOutput {
	return o
}

func (o ServerlessUpstreamSettingsPtrOutput) Elem() ServerlessUpstreamSettingsOutput {
	return o.ApplyT(func(v *ServerlessUpstreamSettings) ServerlessUpstreamSettings {
		if v != nil {
			return *v
		}
		var ret ServerlessUpstreamSettings
		return ret
	}).(ServerlessUpstreamSettingsOutput)
}

func (o ServerlessUpstreamSettingsPtrOutput) Templates() UpstreamTemplateArrayOutput {
	return o.ApplyT(func(v *ServerlessUpstreamSettings) []UpstreamTemplate {
		if v == nil {
			return nil
		}
		return v.Templates
	}).(UpstreamTemplateArrayOutput)
}

type ServerlessUpstreamSettingsResponse struct {
	Templates []UpstreamTemplateResponse `pulumi:"templates"`
}

type ServerlessUpstreamSettingsResponseOutput struct{ *pulumi.OutputState }

func (ServerlessUpstreamSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerlessUpstreamSettingsResponse)(nil)).Elem()
}

func (o ServerlessUpstreamSettingsResponseOutput) ToServerlessUpstreamSettingsResponseOutput() ServerlessUpstreamSettingsResponseOutput {
	return o
}

func (o ServerlessUpstreamSettingsResponseOutput) ToServerlessUpstreamSettingsResponseOutputWithContext(ctx context.Context) ServerlessUpstreamSettingsResponseOutput {
	return o
}

func (o ServerlessUpstreamSettingsResponseOutput) Templates() UpstreamTemplateResponseArrayOutput {
	return o.ApplyT(func(v ServerlessUpstreamSettingsResponse) []UpstreamTemplateResponse { return v.Templates }).(UpstreamTemplateResponseArrayOutput)
}

type ServerlessUpstreamSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerlessUpstreamSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerlessUpstreamSettingsResponse)(nil)).Elem()
}

func (o ServerlessUpstreamSettingsResponsePtrOutput) ToServerlessUpstreamSettingsResponsePtrOutput() ServerlessUpstreamSettingsResponsePtrOutput {
	return o
}

func (o ServerlessUpstreamSettingsResponsePtrOutput) ToServerlessUpstreamSettingsResponsePtrOutputWithContext(ctx context.Context) ServerlessUpstreamSettingsResponsePtrOutput {
	return o
}

func (o ServerlessUpstreamSettingsResponsePtrOutput) Elem() ServerlessUpstreamSettingsResponseOutput {
	return o.ApplyT(func(v *ServerlessUpstreamSettingsResponse) ServerlessUpstreamSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ServerlessUpstreamSettingsResponse
		return ret
	}).(ServerlessUpstreamSettingsResponseOutput)
}

func (o ServerlessUpstreamSettingsResponsePtrOutput) Templates() UpstreamTemplateResponseArrayOutput {
	return o.ApplyT(func(v *ServerlessUpstreamSettingsResponse) []UpstreamTemplateResponse {
		if v == nil {
			return nil
		}
		return v.Templates
	}).(UpstreamTemplateResponseArrayOutput)
}

type SignalRCorsSettings struct {
	AllowedOrigins []string `pulumi:"allowedOrigins"`
}





type SignalRCorsSettingsInput interface {
	pulumi.Input

	ToSignalRCorsSettingsOutput() SignalRCorsSettingsOutput
	ToSignalRCorsSettingsOutputWithContext(context.Context) SignalRCorsSettingsOutput
}

type SignalRCorsSettingsArgs struct {
	AllowedOrigins pulumi.StringArrayInput `pulumi:"allowedOrigins"`
}

func (SignalRCorsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRCorsSettings)(nil)).Elem()
}

func (i SignalRCorsSettingsArgs) ToSignalRCorsSettingsOutput() SignalRCorsSettingsOutput {
	return i.ToSignalRCorsSettingsOutputWithContext(context.Background())
}

func (i SignalRCorsSettingsArgs) ToSignalRCorsSettingsOutputWithContext(ctx context.Context) SignalRCorsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCorsSettingsOutput)
}

func (i SignalRCorsSettingsArgs) ToSignalRCorsSettingsPtrOutput() SignalRCorsSettingsPtrOutput {
	return i.ToSignalRCorsSettingsPtrOutputWithContext(context.Background())
}

func (i SignalRCorsSettingsArgs) ToSignalRCorsSettingsPtrOutputWithContext(ctx context.Context) SignalRCorsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCorsSettingsOutput).ToSignalRCorsSettingsPtrOutputWithContext(ctx)
}









type SignalRCorsSettingsPtrInput interface {
	pulumi.Input

	ToSignalRCorsSettingsPtrOutput() SignalRCorsSettingsPtrOutput
	ToSignalRCorsSettingsPtrOutputWithContext(context.Context) SignalRCorsSettingsPtrOutput
}

type signalRCorsSettingsPtrType SignalRCorsSettingsArgs

func SignalRCorsSettingsPtr(v *SignalRCorsSettingsArgs) SignalRCorsSettingsPtrInput {
	return (*signalRCorsSettingsPtrType)(v)
}

func (*signalRCorsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCorsSettings)(nil)).Elem()
}

func (i *signalRCorsSettingsPtrType) ToSignalRCorsSettingsPtrOutput() SignalRCorsSettingsPtrOutput {
	return i.ToSignalRCorsSettingsPtrOutputWithContext(context.Background())
}

func (i *signalRCorsSettingsPtrType) ToSignalRCorsSettingsPtrOutputWithContext(ctx context.Context) SignalRCorsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRCorsSettingsPtrOutput)
}

type SignalRCorsSettingsOutput struct{ *pulumi.OutputState }

func (SignalRCorsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRCorsSettings)(nil)).Elem()
}

func (o SignalRCorsSettingsOutput) ToSignalRCorsSettingsOutput() SignalRCorsSettingsOutput {
	return o
}

func (o SignalRCorsSettingsOutput) ToSignalRCorsSettingsOutputWithContext(ctx context.Context) SignalRCorsSettingsOutput {
	return o
}

func (o SignalRCorsSettingsOutput) ToSignalRCorsSettingsPtrOutput() SignalRCorsSettingsPtrOutput {
	return o.ToSignalRCorsSettingsPtrOutputWithContext(context.Background())
}

func (o SignalRCorsSettingsOutput) ToSignalRCorsSettingsPtrOutputWithContext(ctx context.Context) SignalRCorsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignalRCorsSettings) *SignalRCorsSettings {
		return &v
	}).(SignalRCorsSettingsPtrOutput)
}

func (o SignalRCorsSettingsOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SignalRCorsSettings) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

type SignalRCorsSettingsPtrOutput struct{ *pulumi.OutputState }

func (SignalRCorsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCorsSettings)(nil)).Elem()
}

func (o SignalRCorsSettingsPtrOutput) ToSignalRCorsSettingsPtrOutput() SignalRCorsSettingsPtrOutput {
	return o
}

func (o SignalRCorsSettingsPtrOutput) ToSignalRCorsSettingsPtrOutputWithContext(ctx context.Context) SignalRCorsSettingsPtrOutput {
	return o
}

func (o SignalRCorsSettingsPtrOutput) Elem() SignalRCorsSettingsOutput {
	return o.ApplyT(func(v *SignalRCorsSettings) SignalRCorsSettings {
		if v != nil {
			return *v
		}
		var ret SignalRCorsSettings
		return ret
	}).(SignalRCorsSettingsOutput)
}

func (o SignalRCorsSettingsPtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SignalRCorsSettings) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

type SignalRCorsSettingsResponse struct {
	AllowedOrigins []string `pulumi:"allowedOrigins"`
}

type SignalRCorsSettingsResponseOutput struct{ *pulumi.OutputState }

func (SignalRCorsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRCorsSettingsResponse)(nil)).Elem()
}

func (o SignalRCorsSettingsResponseOutput) ToSignalRCorsSettingsResponseOutput() SignalRCorsSettingsResponseOutput {
	return o
}

func (o SignalRCorsSettingsResponseOutput) ToSignalRCorsSettingsResponseOutputWithContext(ctx context.Context) SignalRCorsSettingsResponseOutput {
	return o
}

func (o SignalRCorsSettingsResponseOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SignalRCorsSettingsResponse) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

type SignalRCorsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SignalRCorsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRCorsSettingsResponse)(nil)).Elem()
}

func (o SignalRCorsSettingsResponsePtrOutput) ToSignalRCorsSettingsResponsePtrOutput() SignalRCorsSettingsResponsePtrOutput {
	return o
}

func (o SignalRCorsSettingsResponsePtrOutput) ToSignalRCorsSettingsResponsePtrOutputWithContext(ctx context.Context) SignalRCorsSettingsResponsePtrOutput {
	return o
}

func (o SignalRCorsSettingsResponsePtrOutput) Elem() SignalRCorsSettingsResponseOutput {
	return o.ApplyT(func(v *SignalRCorsSettingsResponse) SignalRCorsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret SignalRCorsSettingsResponse
		return ret
	}).(SignalRCorsSettingsResponseOutput)
}

func (o SignalRCorsSettingsResponsePtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SignalRCorsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

type SignalRFeature struct {
	Flag       string            `pulumi:"flag"`
	Properties map[string]string `pulumi:"properties"`
	Value      string            `pulumi:"value"`
}





type SignalRFeatureInput interface {
	pulumi.Input

	ToSignalRFeatureOutput() SignalRFeatureOutput
	ToSignalRFeatureOutputWithContext(context.Context) SignalRFeatureOutput
}

type SignalRFeatureArgs struct {
	Flag       pulumi.StringInput    `pulumi:"flag"`
	Properties pulumi.StringMapInput `pulumi:"properties"`
	Value      pulumi.StringInput    `pulumi:"value"`
}

func (SignalRFeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRFeature)(nil)).Elem()
}

func (i SignalRFeatureArgs) ToSignalRFeatureOutput() SignalRFeatureOutput {
	return i.ToSignalRFeatureOutputWithContext(context.Background())
}

func (i SignalRFeatureArgs) ToSignalRFeatureOutputWithContext(ctx context.Context) SignalRFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRFeatureOutput)
}





type SignalRFeatureArrayInput interface {
	pulumi.Input

	ToSignalRFeatureArrayOutput() SignalRFeatureArrayOutput
	ToSignalRFeatureArrayOutputWithContext(context.Context) SignalRFeatureArrayOutput
}

type SignalRFeatureArray []SignalRFeatureInput

func (SignalRFeatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SignalRFeature)(nil)).Elem()
}

func (i SignalRFeatureArray) ToSignalRFeatureArrayOutput() SignalRFeatureArrayOutput {
	return i.ToSignalRFeatureArrayOutputWithContext(context.Background())
}

func (i SignalRFeatureArray) ToSignalRFeatureArrayOutputWithContext(ctx context.Context) SignalRFeatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRFeatureArrayOutput)
}

type SignalRFeatureOutput struct{ *pulumi.OutputState }

func (SignalRFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRFeature)(nil)).Elem()
}

func (o SignalRFeatureOutput) ToSignalRFeatureOutput() SignalRFeatureOutput {
	return o
}

func (o SignalRFeatureOutput) ToSignalRFeatureOutputWithContext(ctx context.Context) SignalRFeatureOutput {
	return o
}

func (o SignalRFeatureOutput) Flag() pulumi.StringOutput {
	return o.ApplyT(func(v SignalRFeature) string { return v.Flag }).(pulumi.StringOutput)
}

func (o SignalRFeatureOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v SignalRFeature) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o SignalRFeatureOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SignalRFeature) string { return v.Value }).(pulumi.StringOutput)
}

type SignalRFeatureArrayOutput struct{ *pulumi.OutputState }

func (SignalRFeatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SignalRFeature)(nil)).Elem()
}

func (o SignalRFeatureArrayOutput) ToSignalRFeatureArrayOutput() SignalRFeatureArrayOutput {
	return o
}

func (o SignalRFeatureArrayOutput) ToSignalRFeatureArrayOutputWithContext(ctx context.Context) SignalRFeatureArrayOutput {
	return o
}

func (o SignalRFeatureArrayOutput) Index(i pulumi.IntInput) SignalRFeatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SignalRFeature {
		return vs[0].([]SignalRFeature)[vs[1].(int)]
	}).(SignalRFeatureOutput)
}

type SignalRFeatureResponse struct {
	Flag       string            `pulumi:"flag"`
	Properties map[string]string `pulumi:"properties"`
	Value      string            `pulumi:"value"`
}

type SignalRFeatureResponseOutput struct{ *pulumi.OutputState }

func (SignalRFeatureResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRFeatureResponse)(nil)).Elem()
}

func (o SignalRFeatureResponseOutput) ToSignalRFeatureResponseOutput() SignalRFeatureResponseOutput {
	return o
}

func (o SignalRFeatureResponseOutput) ToSignalRFeatureResponseOutputWithContext(ctx context.Context) SignalRFeatureResponseOutput {
	return o
}

func (o SignalRFeatureResponseOutput) Flag() pulumi.StringOutput {
	return o.ApplyT(func(v SignalRFeatureResponse) string { return v.Flag }).(pulumi.StringOutput)
}

func (o SignalRFeatureResponseOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v SignalRFeatureResponse) map[string]string { return v.Properties }).(pulumi.StringMapOutput)
}

func (o SignalRFeatureResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v SignalRFeatureResponse) string { return v.Value }).(pulumi.StringOutput)
}

type SignalRFeatureResponseArrayOutput struct{ *pulumi.OutputState }

func (SignalRFeatureResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SignalRFeatureResponse)(nil)).Elem()
}

func (o SignalRFeatureResponseArrayOutput) ToSignalRFeatureResponseArrayOutput() SignalRFeatureResponseArrayOutput {
	return o
}

func (o SignalRFeatureResponseArrayOutput) ToSignalRFeatureResponseArrayOutputWithContext(ctx context.Context) SignalRFeatureResponseArrayOutput {
	return o
}

func (o SignalRFeatureResponseArrayOutput) Index(i pulumi.IntInput) SignalRFeatureResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SignalRFeatureResponse {
		return vs[0].([]SignalRFeatureResponse)[vs[1].(int)]
	}).(SignalRFeatureResponseOutput)
}

type SignalRNetworkACLs struct {
	DefaultAction    *string              `pulumi:"defaultAction"`
	PrivateEndpoints []PrivateEndpointACL `pulumi:"privateEndpoints"`
	PublicNetwork    *NetworkACL          `pulumi:"publicNetwork"`
}


func (val *SignalRNetworkACLs) Defaults() *SignalRNetworkACLs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		defaultAction_ := "Deny"
		tmp.DefaultAction = &defaultAction_
	}
	return &tmp
}





type SignalRNetworkACLsInput interface {
	pulumi.Input

	ToSignalRNetworkACLsOutput() SignalRNetworkACLsOutput
	ToSignalRNetworkACLsOutputWithContext(context.Context) SignalRNetworkACLsOutput
}

type SignalRNetworkACLsArgs struct {
	DefaultAction    pulumi.StringPtrInput        `pulumi:"defaultAction"`
	PrivateEndpoints PrivateEndpointACLArrayInput `pulumi:"privateEndpoints"`
	PublicNetwork    NetworkACLPtrInput           `pulumi:"publicNetwork"`
}


func (val *SignalRNetworkACLsArgs) Defaults() *SignalRNetworkACLsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = pulumi.StringPtr("Deny")
	}
	return &tmp
}
func (SignalRNetworkACLsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRNetworkACLs)(nil)).Elem()
}

func (i SignalRNetworkACLsArgs) ToSignalRNetworkACLsOutput() SignalRNetworkACLsOutput {
	return i.ToSignalRNetworkACLsOutputWithContext(context.Background())
}

func (i SignalRNetworkACLsArgs) ToSignalRNetworkACLsOutputWithContext(ctx context.Context) SignalRNetworkACLsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRNetworkACLsOutput)
}

func (i SignalRNetworkACLsArgs) ToSignalRNetworkACLsPtrOutput() SignalRNetworkACLsPtrOutput {
	return i.ToSignalRNetworkACLsPtrOutputWithContext(context.Background())
}

func (i SignalRNetworkACLsArgs) ToSignalRNetworkACLsPtrOutputWithContext(ctx context.Context) SignalRNetworkACLsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRNetworkACLsOutput).ToSignalRNetworkACLsPtrOutputWithContext(ctx)
}









type SignalRNetworkACLsPtrInput interface {
	pulumi.Input

	ToSignalRNetworkACLsPtrOutput() SignalRNetworkACLsPtrOutput
	ToSignalRNetworkACLsPtrOutputWithContext(context.Context) SignalRNetworkACLsPtrOutput
}

type signalRNetworkACLsPtrType SignalRNetworkACLsArgs

func SignalRNetworkACLsPtr(v *SignalRNetworkACLsArgs) SignalRNetworkACLsPtrInput {
	return (*signalRNetworkACLsPtrType)(v)
}

func (*signalRNetworkACLsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRNetworkACLs)(nil)).Elem()
}

func (i *signalRNetworkACLsPtrType) ToSignalRNetworkACLsPtrOutput() SignalRNetworkACLsPtrOutput {
	return i.ToSignalRNetworkACLsPtrOutputWithContext(context.Background())
}

func (i *signalRNetworkACLsPtrType) ToSignalRNetworkACLsPtrOutputWithContext(ctx context.Context) SignalRNetworkACLsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRNetworkACLsPtrOutput)
}

type SignalRNetworkACLsOutput struct{ *pulumi.OutputState }

func (SignalRNetworkACLsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRNetworkACLs)(nil)).Elem()
}

func (o SignalRNetworkACLsOutput) ToSignalRNetworkACLsOutput() SignalRNetworkACLsOutput {
	return o
}

func (o SignalRNetworkACLsOutput) ToSignalRNetworkACLsOutputWithContext(ctx context.Context) SignalRNetworkACLsOutput {
	return o
}

func (o SignalRNetworkACLsOutput) ToSignalRNetworkACLsPtrOutput() SignalRNetworkACLsPtrOutput {
	return o.ToSignalRNetworkACLsPtrOutputWithContext(context.Background())
}

func (o SignalRNetworkACLsOutput) ToSignalRNetworkACLsPtrOutputWithContext(ctx context.Context) SignalRNetworkACLsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignalRNetworkACLs) *SignalRNetworkACLs {
		return &v
	}).(SignalRNetworkACLsPtrOutput)
}

func (o SignalRNetworkACLsOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SignalRNetworkACLs) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o SignalRNetworkACLsOutput) PrivateEndpoints() PrivateEndpointACLArrayOutput {
	return o.ApplyT(func(v SignalRNetworkACLs) []PrivateEndpointACL { return v.PrivateEndpoints }).(PrivateEndpointACLArrayOutput)
}

func (o SignalRNetworkACLsOutput) PublicNetwork() NetworkACLPtrOutput {
	return o.ApplyT(func(v SignalRNetworkACLs) *NetworkACL { return v.PublicNetwork }).(NetworkACLPtrOutput)
}

type SignalRNetworkACLsPtrOutput struct{ *pulumi.OutputState }

func (SignalRNetworkACLsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRNetworkACLs)(nil)).Elem()
}

func (o SignalRNetworkACLsPtrOutput) ToSignalRNetworkACLsPtrOutput() SignalRNetworkACLsPtrOutput {
	return o
}

func (o SignalRNetworkACLsPtrOutput) ToSignalRNetworkACLsPtrOutputWithContext(ctx context.Context) SignalRNetworkACLsPtrOutput {
	return o
}

func (o SignalRNetworkACLsPtrOutput) Elem() SignalRNetworkACLsOutput {
	return o.ApplyT(func(v *SignalRNetworkACLs) SignalRNetworkACLs {
		if v != nil {
			return *v
		}
		var ret SignalRNetworkACLs
		return ret
	}).(SignalRNetworkACLsOutput)
}

func (o SignalRNetworkACLsPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalRNetworkACLs) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o SignalRNetworkACLsPtrOutput) PrivateEndpoints() PrivateEndpointACLArrayOutput {
	return o.ApplyT(func(v *SignalRNetworkACLs) []PrivateEndpointACL {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoints
	}).(PrivateEndpointACLArrayOutput)
}

func (o SignalRNetworkACLsPtrOutput) PublicNetwork() NetworkACLPtrOutput {
	return o.ApplyT(func(v *SignalRNetworkACLs) *NetworkACL {
		if v == nil {
			return nil
		}
		return v.PublicNetwork
	}).(NetworkACLPtrOutput)
}

type SignalRNetworkACLsResponse struct {
	DefaultAction    *string                      `pulumi:"defaultAction"`
	PrivateEndpoints []PrivateEndpointACLResponse `pulumi:"privateEndpoints"`
	PublicNetwork    *NetworkACLResponse          `pulumi:"publicNetwork"`
}


func (val *SignalRNetworkACLsResponse) Defaults() *SignalRNetworkACLsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		defaultAction_ := "Deny"
		tmp.DefaultAction = &defaultAction_
	}
	return &tmp
}

type SignalRNetworkACLsResponseOutput struct{ *pulumi.OutputState }

func (SignalRNetworkACLsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRNetworkACLsResponse)(nil)).Elem()
}

func (o SignalRNetworkACLsResponseOutput) ToSignalRNetworkACLsResponseOutput() SignalRNetworkACLsResponseOutput {
	return o
}

func (o SignalRNetworkACLsResponseOutput) ToSignalRNetworkACLsResponseOutputWithContext(ctx context.Context) SignalRNetworkACLsResponseOutput {
	return o
}

func (o SignalRNetworkACLsResponseOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SignalRNetworkACLsResponse) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o SignalRNetworkACLsResponseOutput) PrivateEndpoints() PrivateEndpointACLResponseArrayOutput {
	return o.ApplyT(func(v SignalRNetworkACLsResponse) []PrivateEndpointACLResponse { return v.PrivateEndpoints }).(PrivateEndpointACLResponseArrayOutput)
}

func (o SignalRNetworkACLsResponseOutput) PublicNetwork() NetworkACLResponsePtrOutput {
	return o.ApplyT(func(v SignalRNetworkACLsResponse) *NetworkACLResponse { return v.PublicNetwork }).(NetworkACLResponsePtrOutput)
}

type SignalRNetworkACLsResponsePtrOutput struct{ *pulumi.OutputState }

func (SignalRNetworkACLsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRNetworkACLsResponse)(nil)).Elem()
}

func (o SignalRNetworkACLsResponsePtrOutput) ToSignalRNetworkACLsResponsePtrOutput() SignalRNetworkACLsResponsePtrOutput {
	return o
}

func (o SignalRNetworkACLsResponsePtrOutput) ToSignalRNetworkACLsResponsePtrOutputWithContext(ctx context.Context) SignalRNetworkACLsResponsePtrOutput {
	return o
}

func (o SignalRNetworkACLsResponsePtrOutput) Elem() SignalRNetworkACLsResponseOutput {
	return o.ApplyT(func(v *SignalRNetworkACLsResponse) SignalRNetworkACLsResponse {
		if v != nil {
			return *v
		}
		var ret SignalRNetworkACLsResponse
		return ret
	}).(SignalRNetworkACLsResponseOutput)
}

func (o SignalRNetworkACLsResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SignalRNetworkACLsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o SignalRNetworkACLsResponsePtrOutput) PrivateEndpoints() PrivateEndpointACLResponseArrayOutput {
	return o.ApplyT(func(v *SignalRNetworkACLsResponse) []PrivateEndpointACLResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoints
	}).(PrivateEndpointACLResponseArrayOutput)
}

func (o SignalRNetworkACLsResponsePtrOutput) PublicNetwork() NetworkACLResponsePtrOutput {
	return o.ApplyT(func(v *SignalRNetworkACLsResponse) *NetworkACLResponse {
		if v == nil {
			return nil
		}
		return v.PublicNetwork
	}).(NetworkACLResponsePtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

type UpstreamTemplate struct {
	CategoryPattern *string `pulumi:"categoryPattern"`
	EventPattern    *string `pulumi:"eventPattern"`
	HubPattern      *string `pulumi:"hubPattern"`
	UrlTemplate     string  `pulumi:"urlTemplate"`
}





type UpstreamTemplateInput interface {
	pulumi.Input

	ToUpstreamTemplateOutput() UpstreamTemplateOutput
	ToUpstreamTemplateOutputWithContext(context.Context) UpstreamTemplateOutput
}

type UpstreamTemplateArgs struct {
	CategoryPattern pulumi.StringPtrInput `pulumi:"categoryPattern"`
	EventPattern    pulumi.StringPtrInput `pulumi:"eventPattern"`
	HubPattern      pulumi.StringPtrInput `pulumi:"hubPattern"`
	UrlTemplate     pulumi.StringInput    `pulumi:"urlTemplate"`
}

func (UpstreamTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamTemplate)(nil)).Elem()
}

func (i UpstreamTemplateArgs) ToUpstreamTemplateOutput() UpstreamTemplateOutput {
	return i.ToUpstreamTemplateOutputWithContext(context.Background())
}

func (i UpstreamTemplateArgs) ToUpstreamTemplateOutputWithContext(ctx context.Context) UpstreamTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamTemplateOutput)
}





type UpstreamTemplateArrayInput interface {
	pulumi.Input

	ToUpstreamTemplateArrayOutput() UpstreamTemplateArrayOutput
	ToUpstreamTemplateArrayOutputWithContext(context.Context) UpstreamTemplateArrayOutput
}

type UpstreamTemplateArray []UpstreamTemplateInput

func (UpstreamTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpstreamTemplate)(nil)).Elem()
}

func (i UpstreamTemplateArray) ToUpstreamTemplateArrayOutput() UpstreamTemplateArrayOutput {
	return i.ToUpstreamTemplateArrayOutputWithContext(context.Background())
}

func (i UpstreamTemplateArray) ToUpstreamTemplateArrayOutputWithContext(ctx context.Context) UpstreamTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpstreamTemplateArrayOutput)
}

type UpstreamTemplateOutput struct{ *pulumi.OutputState }

func (UpstreamTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamTemplate)(nil)).Elem()
}

func (o UpstreamTemplateOutput) ToUpstreamTemplateOutput() UpstreamTemplateOutput {
	return o
}

func (o UpstreamTemplateOutput) ToUpstreamTemplateOutputWithContext(ctx context.Context) UpstreamTemplateOutput {
	return o
}

func (o UpstreamTemplateOutput) CategoryPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpstreamTemplate) *string { return v.CategoryPattern }).(pulumi.StringPtrOutput)
}

func (o UpstreamTemplateOutput) EventPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpstreamTemplate) *string { return v.EventPattern }).(pulumi.StringPtrOutput)
}

func (o UpstreamTemplateOutput) HubPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpstreamTemplate) *string { return v.HubPattern }).(pulumi.StringPtrOutput)
}

func (o UpstreamTemplateOutput) UrlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v UpstreamTemplate) string { return v.UrlTemplate }).(pulumi.StringOutput)
}

type UpstreamTemplateArrayOutput struct{ *pulumi.OutputState }

func (UpstreamTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpstreamTemplate)(nil)).Elem()
}

func (o UpstreamTemplateArrayOutput) ToUpstreamTemplateArrayOutput() UpstreamTemplateArrayOutput {
	return o
}

func (o UpstreamTemplateArrayOutput) ToUpstreamTemplateArrayOutputWithContext(ctx context.Context) UpstreamTemplateArrayOutput {
	return o
}

func (o UpstreamTemplateArrayOutput) Index(i pulumi.IntInput) UpstreamTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UpstreamTemplate {
		return vs[0].([]UpstreamTemplate)[vs[1].(int)]
	}).(UpstreamTemplateOutput)
}

type UpstreamTemplateResponse struct {
	CategoryPattern *string `pulumi:"categoryPattern"`
	EventPattern    *string `pulumi:"eventPattern"`
	HubPattern      *string `pulumi:"hubPattern"`
	UrlTemplate     string  `pulumi:"urlTemplate"`
}

type UpstreamTemplateResponseOutput struct{ *pulumi.OutputState }

func (UpstreamTemplateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpstreamTemplateResponse)(nil)).Elem()
}

func (o UpstreamTemplateResponseOutput) ToUpstreamTemplateResponseOutput() UpstreamTemplateResponseOutput {
	return o
}

func (o UpstreamTemplateResponseOutput) ToUpstreamTemplateResponseOutputWithContext(ctx context.Context) UpstreamTemplateResponseOutput {
	return o
}

func (o UpstreamTemplateResponseOutput) CategoryPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpstreamTemplateResponse) *string { return v.CategoryPattern }).(pulumi.StringPtrOutput)
}

func (o UpstreamTemplateResponseOutput) EventPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpstreamTemplateResponse) *string { return v.EventPattern }).(pulumi.StringPtrOutput)
}

func (o UpstreamTemplateResponseOutput) HubPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpstreamTemplateResponse) *string { return v.HubPattern }).(pulumi.StringPtrOutput)
}

func (o UpstreamTemplateResponseOutput) UrlTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v UpstreamTemplateResponse) string { return v.UrlTemplate }).(pulumi.StringOutput)
}

type UpstreamTemplateResponseArrayOutput struct{ *pulumi.OutputState }

func (UpstreamTemplateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UpstreamTemplateResponse)(nil)).Elem()
}

func (o UpstreamTemplateResponseArrayOutput) ToUpstreamTemplateResponseArrayOutput() UpstreamTemplateResponseArrayOutput {
	return o
}

func (o UpstreamTemplateResponseArrayOutput) ToUpstreamTemplateResponseArrayOutputWithContext(ctx context.Context) UpstreamTemplateResponseArrayOutput {
	return o
}

func (o UpstreamTemplateResponseArrayOutput) Index(i pulumi.IntInput) UpstreamTemplateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UpstreamTemplateResponse {
		return vs[0].([]UpstreamTemplateResponse)[vs[1].(int)]
	}).(UpstreamTemplateResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkACLOutput{})
	pulumi.RegisterOutputType(NetworkACLPtrOutput{})
	pulumi.RegisterOutputType(NetworkACLResponseOutput{})
	pulumi.RegisterOutputType(NetworkACLResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointACLResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceReferenceOutput{})
	pulumi.RegisterOutputType(ResourceReferenceResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuOutput{})
	pulumi.RegisterOutputType(ResourceSkuPtrOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponseOutput{})
	pulumi.RegisterOutputType(ResourceSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerlessUpstreamSettingsOutput{})
	pulumi.RegisterOutputType(ServerlessUpstreamSettingsPtrOutput{})
	pulumi.RegisterOutputType(ServerlessUpstreamSettingsResponseOutput{})
	pulumi.RegisterOutputType(ServerlessUpstreamSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SignalRCorsSettingsOutput{})
	pulumi.RegisterOutputType(SignalRCorsSettingsPtrOutput{})
	pulumi.RegisterOutputType(SignalRCorsSettingsResponseOutput{})
	pulumi.RegisterOutputType(SignalRCorsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SignalRFeatureOutput{})
	pulumi.RegisterOutputType(SignalRFeatureArrayOutput{})
	pulumi.RegisterOutputType(SignalRFeatureResponseOutput{})
	pulumi.RegisterOutputType(SignalRFeatureResponseArrayOutput{})
	pulumi.RegisterOutputType(SignalRNetworkACLsOutput{})
	pulumi.RegisterOutputType(SignalRNetworkACLsPtrOutput{})
	pulumi.RegisterOutputType(SignalRNetworkACLsResponseOutput{})
	pulumi.RegisterOutputType(SignalRNetworkACLsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(UpstreamTemplateOutput{})
	pulumi.RegisterOutputType(UpstreamTemplateArrayOutput{})
	pulumi.RegisterOutputType(UpstreamTemplateResponseOutput{})
	pulumi.RegisterOutputType(UpstreamTemplateResponseArrayOutput{})
}
