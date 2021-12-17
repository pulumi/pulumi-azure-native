


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Identity struct {
	Type IdentityType `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type IdentityTypeInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() IdentityTypeOutput {
	return o.ApplyT(func(v Identity) IdentityType { return v.Type }).(IdentityTypeOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() IdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *IdentityType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(IdentityTypePtrOutput)
}

type IdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type IpRule struct {
	Value *string `pulumi:"value"`
}





type IpRuleInput interface {
	pulumi.Input

	ToIpRuleOutput() IpRuleOutput
	ToIpRuleOutputWithContext(context.Context) IpRuleOutput
}

type IpRuleArgs struct {
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (IpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRule)(nil)).Elem()
}

func (i IpRuleArgs) ToIpRuleOutput() IpRuleOutput {
	return i.ToIpRuleOutputWithContext(context.Background())
}

func (i IpRuleArgs) ToIpRuleOutputWithContext(ctx context.Context) IpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleOutput)
}





type IpRuleArrayInput interface {
	pulumi.Input

	ToIpRuleArrayOutput() IpRuleArrayOutput
	ToIpRuleArrayOutputWithContext(context.Context) IpRuleArrayOutput
}

type IpRuleArray []IpRuleInput

func (IpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRule)(nil)).Elem()
}

func (i IpRuleArray) ToIpRuleArrayOutput() IpRuleArrayOutput {
	return i.ToIpRuleArrayOutputWithContext(context.Background())
}

func (i IpRuleArray) ToIpRuleArrayOutputWithContext(ctx context.Context) IpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpRuleArrayOutput)
}

type IpRuleOutput struct{ *pulumi.OutputState }

func (IpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRule)(nil)).Elem()
}

func (o IpRuleOutput) ToIpRuleOutput() IpRuleOutput {
	return o
}

func (o IpRuleOutput) ToIpRuleOutputWithContext(ctx context.Context) IpRuleOutput {
	return o
}

func (o IpRuleOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpRule) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type IpRuleArrayOutput struct{ *pulumi.OutputState }

func (IpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRule)(nil)).Elem()
}

func (o IpRuleArrayOutput) ToIpRuleArrayOutput() IpRuleArrayOutput {
	return o
}

func (o IpRuleArrayOutput) ToIpRuleArrayOutputWithContext(ctx context.Context) IpRuleArrayOutput {
	return o
}

func (o IpRuleArrayOutput) Index(i pulumi.IntInput) IpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpRule {
		return vs[0].([]IpRule)[vs[1].(int)]
	}).(IpRuleOutput)
}

type IpRuleResponse struct {
	Value *string `pulumi:"value"`
}

type IpRuleResponseOutput struct{ *pulumi.OutputState }

func (IpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpRuleResponse)(nil)).Elem()
}

func (o IpRuleResponseOutput) ToIpRuleResponseOutput() IpRuleResponseOutput {
	return o
}

func (o IpRuleResponseOutput) ToIpRuleResponseOutputWithContext(ctx context.Context) IpRuleResponseOutput {
	return o
}

func (o IpRuleResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpRuleResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type IpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpRuleResponse)(nil)).Elem()
}

func (o IpRuleResponseArrayOutput) ToIpRuleResponseArrayOutput() IpRuleResponseArrayOutput {
	return o
}

func (o IpRuleResponseArrayOutput) ToIpRuleResponseArrayOutputWithContext(ctx context.Context) IpRuleResponseArrayOutput {
	return o
}

func (o IpRuleResponseArrayOutput) Index(i pulumi.IntInput) IpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpRuleResponse {
		return vs[0].([]IpRuleResponse)[vs[1].(int)]
	}).(IpRuleResponseOutput)
}

type NetworkRuleSet struct {
	EndpointAccess *EndpointAccess `pulumi:"endpointAccess"`
	IpRules        []IpRule        `pulumi:"ipRules"`
}


func (val *NetworkRuleSet) Defaults() *NetworkRuleSet {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EndpointAccess) {
		endpointAccess_ := EndpointAccess("Public")
		tmp.EndpointAccess = &endpointAccess_
	}
	return &tmp
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	EndpointAccess EndpointAccessPtrInput `pulumi:"endpointAccess"`
	IpRules        IpRuleArrayInput       `pulumi:"ipRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) EndpointAccess() EndpointAccessPtrOutput {
	return o.ApplyT(func(v NetworkRuleSet) *EndpointAccess { return v.EndpointAccess }).(EndpointAccessPtrOutput)
}

func (o NetworkRuleSetOutput) IpRules() IpRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IpRule { return v.IpRules }).(IpRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) EndpointAccess() EndpointAccessPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *EndpointAccess {
		if v == nil {
			return nil
		}
		return v.EndpointAccess
	}).(EndpointAccessPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IpRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IpRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IpRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	EndpointAccess *string          `pulumi:"endpointAccess"`
	IpRules        []IpRuleResponse `pulumi:"ipRules"`
}


func (val *NetworkRuleSetResponse) Defaults() *NetworkRuleSetResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EndpointAccess) {
		endpointAccess_ := "Public"
		tmp.EndpointAccess = &endpointAccess_
	}
	return &tmp
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) EndpointAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) *string { return v.EndpointAccess }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IpRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IpRuleResponse { return v.IpRules }).(IpRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) EndpointAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndpointAccess
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) IpRules() IpRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []IpRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IpRuleResponseArrayOutput)
}

type PrivateEndpointConnectionProperties struct {
	PrivateEndpoint                   *PrivateEndpointConnectionPropertiesPrivateEndpoint                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}


func (val *PrivateEndpointConnectionProperties) Defaults() *PrivateEndpointConnectionProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PrivateLinkServiceConnectionState = tmp.PrivateLinkServiceConnectionState.Defaults()

	return &tmp
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	PrivateEndpoint                   PrivateEndpointConnectionPropertiesPrivateEndpointPtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
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

func (o PrivateEndpointConnectionPropertiesOutput) PrivateEndpoint() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionPropertiesPrivateEndpoint {
		return v.PrivateEndpoint
	}).(PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
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

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateEndpoint() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateEndpointConnectionPropertiesPrivateEndpoint {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointConnectionPropertiesPrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPrivateEndpointOutput() PrivateEndpointConnectionPropertiesPrivateEndpointOutput
	ToPrivateEndpointConnectionPropertiesPrivateEndpointOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointOutput
}

type PrivateEndpointConnectionPropertiesPrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointConnectionPropertiesPrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesPrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesPrivateEndpointOutput() PrivateEndpointConnectionPropertiesPrivateEndpointOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesPrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateEndpointOutput)
}

func (i PrivateEndpointConnectionPropertiesPrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesPrivateEndpointArgs) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateEndpointOutput).ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput
	ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput
}

type privateEndpointConnectionPropertiesPrivateEndpointPtrType PrivateEndpointConnectionPropertiesPrivateEndpointArgs

func PrivateEndpointConnectionPropertiesPrivateEndpointPtr(v *PrivateEndpointConnectionPropertiesPrivateEndpointArgs) PrivateEndpointConnectionPropertiesPrivateEndpointPtrInput {
	return (*privateEndpointConnectionPropertiesPrivateEndpointPtrType)(v)
}

func (*privateEndpointConnectionPropertiesPrivateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesPrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPrivateEndpointPtrType) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPrivateEndpointPtrType) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointOutput() PrivateEndpointConnectionPropertiesPrivateEndpointOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesPrivateEndpoint) *PrivateEndpointConnectionPropertiesPrivateEndpoint {
		return &v
	}).(PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesPrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesPrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput) ToPrivateEndpointConnectionPropertiesPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput) Elem() PrivateEndpointConnectionPropertiesPrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateEndpoint) PrivateEndpointConnectionPropertiesPrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesPrivateEndpoint
		return ret
	}).(PrivateEndpointConnectionPropertiesPrivateEndpointOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState struct {
	ActionsRequired *string                             `pulumi:"actionsRequired"`
	Description     *string                             `pulumi:"description"`
	Status          *PrivateLinkServiceConnectionStatus `pulumi:"status"`
}


func (val *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) Defaults() *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ActionsRequired) {
		actionsRequired_ := "None"
		tmp.ActionsRequired = &actionsRequired_
	}
	return &tmp
}





type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput
	ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput
}

type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput                      `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput                      `pulumi:"description"`
	Status          PrivateLinkServiceConnectionStatusPtrInput `pulumi:"status"`
}

func (PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput)
}

func (i PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput).ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput
	ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput
}

type privateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrType PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs

func PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtr(v *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrInput {
	return (*privateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrType)(v)
}

func (*privateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrType) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrType) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState {
		return &v
	}).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput) Status() PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionStatus {
		return v.Status
	}).(PrivateLinkServiceConnectionStatusPtrOutput)
}

type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ToPrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState
		return ret
	}).(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput) Status() PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionStatus {
		if v == nil {
			return nil
		}
		return v.Status
	}).(PrivateLinkServiceConnectionStatusPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}


func (val *PrivateEndpointConnectionPropertiesResponse) Defaults() *PrivateEndpointConnectionPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PrivateLinkServiceConnectionState = tmp.PrivateLinkServiceConnectionState.Defaults()

	return &tmp
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint {
		return v.PrivateEndpoint
	}).(PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) Elem() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) PrivateEndpointConnectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponse
		return ret
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateEndpoint() PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePrivateEndpoint struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput() PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponsePrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponsePrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput) Elem() PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint) PrivateEndpointConnectionPropertiesResponsePrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponsePrivateEndpoint
		return ret
	}).(PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}


func (val *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) Defaults() *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ActionsRequired) {
		actionsRequired_ := "None"
		tmp.ActionsRequired = &actionsRequired_
	}
	return &tmp
}

type PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState
		return ret
	}).(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id         *string                                      `pulumi:"id"`
	Name       string                                       `pulumi:"name"`
	Properties *PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                       `pulumi:"type"`
}


func (val *PrivateEndpointConnectionResponse) Defaults() *PrivateEndpointConnectionResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
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

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
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

type QueryKeyResponse struct {
	Key  string `pulumi:"key"`
	Name string `pulumi:"name"`
}

type Sku struct {
	Name *SkuName `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name SkuNamePtrInput `pulumi:"name"`
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

func (o SkuOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v Sku) *SkuName { return v.Name }).(SkuNamePtrOutput)
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

func (o SkuPtrOutput) Name() SkuNamePtrOutput {
	return o.ApplyT(func(v *Sku) *SkuName {
		if v == nil {
			return nil
		}
		return v.Name
	}).(SkuNamePtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
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

func (o SkuResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
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

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IpRuleOutput{})
	pulumi.RegisterOutputType(IpRuleArrayOutput{})
	pulumi.RegisterOutputType(IpRuleResponseOutput{})
	pulumi.RegisterOutputType(IpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
