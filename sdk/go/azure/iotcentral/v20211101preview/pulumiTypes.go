


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppSkuInfo struct {
	Name string `pulumi:"name"`
}





type AppSkuInfoInput interface {
	pulumi.Input

	ToAppSkuInfoOutput() AppSkuInfoOutput
	ToAppSkuInfoOutputWithContext(context.Context) AppSkuInfoOutput
}

type AppSkuInfoArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (AppSkuInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfo)(nil)).Elem()
}

func (i AppSkuInfoArgs) ToAppSkuInfoOutput() AppSkuInfoOutput {
	return i.ToAppSkuInfoOutputWithContext(context.Background())
}

func (i AppSkuInfoArgs) ToAppSkuInfoOutputWithContext(ctx context.Context) AppSkuInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppSkuInfoOutput)
}

type AppSkuInfoOutput struct{ *pulumi.OutputState }

func (AppSkuInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfo)(nil)).Elem()
}

func (o AppSkuInfoOutput) ToAppSkuInfoOutput() AppSkuInfoOutput {
	return o
}

func (o AppSkuInfoOutput) ToAppSkuInfoOutputWithContext(ctx context.Context) AppSkuInfoOutput {
	return o
}

func (o AppSkuInfoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppSkuInfo) string { return v.Name }).(pulumi.StringOutput)
}

type AppSkuInfoResponse struct {
	Name string `pulumi:"name"`
}

type AppSkuInfoResponseOutput struct{ *pulumi.OutputState }

func (AppSkuInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSkuInfoResponse)(nil)).Elem()
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponseOutput() AppSkuInfoResponseOutput {
	return o
}

func (o AppSkuInfoResponseOutput) ToAppSkuInfoResponseOutputWithContext(ctx context.Context) AppSkuInfoResponseOutput {
	return o
}

func (o AppSkuInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AppSkuInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

type NetworkRuleSetIpRule struct {
	FilterName *string `pulumi:"filterName"`
	IpMask     *string `pulumi:"ipMask"`
}





type NetworkRuleSetIpRuleInput interface {
	pulumi.Input

	ToNetworkRuleSetIpRuleOutput() NetworkRuleSetIpRuleOutput
	ToNetworkRuleSetIpRuleOutputWithContext(context.Context) NetworkRuleSetIpRuleOutput
}

type NetworkRuleSetIpRuleArgs struct {
	FilterName pulumi.StringPtrInput `pulumi:"filterName"`
	IpMask     pulumi.StringPtrInput `pulumi:"ipMask"`
}

func (NetworkRuleSetIpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetIpRule)(nil)).Elem()
}

func (i NetworkRuleSetIpRuleArgs) ToNetworkRuleSetIpRuleOutput() NetworkRuleSetIpRuleOutput {
	return i.ToNetworkRuleSetIpRuleOutputWithContext(context.Background())
}

func (i NetworkRuleSetIpRuleArgs) ToNetworkRuleSetIpRuleOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetIpRuleOutput)
}





type NetworkRuleSetIpRuleArrayInput interface {
	pulumi.Input

	ToNetworkRuleSetIpRuleArrayOutput() NetworkRuleSetIpRuleArrayOutput
	ToNetworkRuleSetIpRuleArrayOutputWithContext(context.Context) NetworkRuleSetIpRuleArrayOutput
}

type NetworkRuleSetIpRuleArray []NetworkRuleSetIpRuleInput

func (NetworkRuleSetIpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRuleSetIpRule)(nil)).Elem()
}

func (i NetworkRuleSetIpRuleArray) ToNetworkRuleSetIpRuleArrayOutput() NetworkRuleSetIpRuleArrayOutput {
	return i.ToNetworkRuleSetIpRuleArrayOutputWithContext(context.Background())
}

func (i NetworkRuleSetIpRuleArray) ToNetworkRuleSetIpRuleArrayOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetIpRuleArrayOutput)
}

type NetworkRuleSetIpRuleOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetIpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetIpRule)(nil)).Elem()
}

func (o NetworkRuleSetIpRuleOutput) ToNetworkRuleSetIpRuleOutput() NetworkRuleSetIpRuleOutput {
	return o
}

func (o NetworkRuleSetIpRuleOutput) ToNetworkRuleSetIpRuleOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleOutput {
	return o
}

func (o NetworkRuleSetIpRuleOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRule) *string { return v.FilterName }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetIpRuleOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRule) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type NetworkRuleSetIpRuleArrayOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetIpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRuleSetIpRule)(nil)).Elem()
}

func (o NetworkRuleSetIpRuleArrayOutput) ToNetworkRuleSetIpRuleArrayOutput() NetworkRuleSetIpRuleArrayOutput {
	return o
}

func (o NetworkRuleSetIpRuleArrayOutput) ToNetworkRuleSetIpRuleArrayOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleArrayOutput {
	return o
}

func (o NetworkRuleSetIpRuleArrayOutput) Index(i pulumi.IntInput) NetworkRuleSetIpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkRuleSetIpRule {
		return vs[0].([]NetworkRuleSetIpRule)[vs[1].(int)]
	}).(NetworkRuleSetIpRuleOutput)
}

type NetworkRuleSetIpRuleResponse struct {
	Action     string  `pulumi:"action"`
	FilterName *string `pulumi:"filterName"`
	IpMask     *string `pulumi:"ipMask"`
}

type NetworkRuleSetIpRuleResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetIpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetIpRuleResponse)(nil)).Elem()
}

func (o NetworkRuleSetIpRuleResponseOutput) ToNetworkRuleSetIpRuleResponseOutput() NetworkRuleSetIpRuleResponseOutput {
	return o
}

func (o NetworkRuleSetIpRuleResponseOutput) ToNetworkRuleSetIpRuleResponseOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleResponseOutput {
	return o
}

func (o NetworkRuleSetIpRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o NetworkRuleSetIpRuleResponseOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRuleResponse) *string { return v.FilterName }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetIpRuleResponseOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetIpRuleResponse) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type NetworkRuleSetIpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetIpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRuleSetIpRuleResponse)(nil)).Elem()
}

func (o NetworkRuleSetIpRuleResponseArrayOutput) ToNetworkRuleSetIpRuleResponseArrayOutput() NetworkRuleSetIpRuleResponseArrayOutput {
	return o
}

func (o NetworkRuleSetIpRuleResponseArrayOutput) ToNetworkRuleSetIpRuleResponseArrayOutputWithContext(ctx context.Context) NetworkRuleSetIpRuleResponseArrayOutput {
	return o
}

func (o NetworkRuleSetIpRuleResponseArrayOutput) Index(i pulumi.IntInput) NetworkRuleSetIpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkRuleSetIpRuleResponse {
		return vs[0].([]NetworkRuleSetIpRuleResponse)[vs[1].(int)]
	}).(NetworkRuleSetIpRuleResponseOutput)
}

type NetworkRuleSets struct {
	ApplyToDevices    *bool                  `pulumi:"applyToDevices"`
	ApplyToIoTCentral *bool                  `pulumi:"applyToIoTCentral"`
	DefaultAction     *string                `pulumi:"defaultAction"`
	IpRules           []NetworkRuleSetIpRule `pulumi:"ipRules"`
}


func (val *NetworkRuleSets) Defaults() *NetworkRuleSets {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ApplyToDevices) {
		applyToDevices_ := false
		tmp.ApplyToDevices = &applyToDevices_
	}
	if isZero(tmp.ApplyToIoTCentral) {
		applyToIoTCentral_ := false
		tmp.ApplyToIoTCentral = &applyToIoTCentral_
	}
	return &tmp
}





type NetworkRuleSetsInput interface {
	pulumi.Input

	ToNetworkRuleSetsOutput() NetworkRuleSetsOutput
	ToNetworkRuleSetsOutputWithContext(context.Context) NetworkRuleSetsOutput
}

type NetworkRuleSetsArgs struct {
	ApplyToDevices    pulumi.BoolPtrInput            `pulumi:"applyToDevices"`
	ApplyToIoTCentral pulumi.BoolPtrInput            `pulumi:"applyToIoTCentral"`
	DefaultAction     pulumi.StringPtrInput          `pulumi:"defaultAction"`
	IpRules           NetworkRuleSetIpRuleArrayInput `pulumi:"ipRules"`
}


func (val *NetworkRuleSetsArgs) Defaults() *NetworkRuleSetsArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ApplyToDevices) {
		tmp.ApplyToDevices = pulumi.BoolPtr(false)
	}
	if isZero(tmp.ApplyToIoTCentral) {
		tmp.ApplyToIoTCentral = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (NetworkRuleSetsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSets)(nil)).Elem()
}

func (i NetworkRuleSetsArgs) ToNetworkRuleSetsOutput() NetworkRuleSetsOutput {
	return i.ToNetworkRuleSetsOutputWithContext(context.Background())
}

func (i NetworkRuleSetsArgs) ToNetworkRuleSetsOutputWithContext(ctx context.Context) NetworkRuleSetsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetsOutput)
}

func (i NetworkRuleSetsArgs) ToNetworkRuleSetsPtrOutput() NetworkRuleSetsPtrOutput {
	return i.ToNetworkRuleSetsPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetsArgs) ToNetworkRuleSetsPtrOutputWithContext(ctx context.Context) NetworkRuleSetsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetsOutput).ToNetworkRuleSetsPtrOutputWithContext(ctx)
}









type NetworkRuleSetsPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetsPtrOutput() NetworkRuleSetsPtrOutput
	ToNetworkRuleSetsPtrOutputWithContext(context.Context) NetworkRuleSetsPtrOutput
}

type networkRuleSetsPtrType NetworkRuleSetsArgs

func NetworkRuleSetsPtr(v *NetworkRuleSetsArgs) NetworkRuleSetsPtrInput {
	return (*networkRuleSetsPtrType)(v)
}

func (*networkRuleSetsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSets)(nil)).Elem()
}

func (i *networkRuleSetsPtrType) ToNetworkRuleSetsPtrOutput() NetworkRuleSetsPtrOutput {
	return i.ToNetworkRuleSetsPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetsPtrType) ToNetworkRuleSetsPtrOutputWithContext(ctx context.Context) NetworkRuleSetsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetsPtrOutput)
}

type NetworkRuleSetsOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSets)(nil)).Elem()
}

func (o NetworkRuleSetsOutput) ToNetworkRuleSetsOutput() NetworkRuleSetsOutput {
	return o
}

func (o NetworkRuleSetsOutput) ToNetworkRuleSetsOutputWithContext(ctx context.Context) NetworkRuleSetsOutput {
	return o
}

func (o NetworkRuleSetsOutput) ToNetworkRuleSetsPtrOutput() NetworkRuleSetsPtrOutput {
	return o.ToNetworkRuleSetsPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetsOutput) ToNetworkRuleSetsPtrOutputWithContext(ctx context.Context) NetworkRuleSetsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSets) *NetworkRuleSets {
		return &v
	}).(NetworkRuleSetsPtrOutput)
}

func (o NetworkRuleSetsOutput) ApplyToDevices() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkRuleSets) *bool { return v.ApplyToDevices }).(pulumi.BoolPtrOutput)
}

func (o NetworkRuleSetsOutput) ApplyToIoTCentral() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkRuleSets) *bool { return v.ApplyToIoTCentral }).(pulumi.BoolPtrOutput)
}

func (o NetworkRuleSetsOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSets) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetsOutput) IpRules() NetworkRuleSetIpRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSets) []NetworkRuleSetIpRule { return v.IpRules }).(NetworkRuleSetIpRuleArrayOutput)
}

type NetworkRuleSetsPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSets)(nil)).Elem()
}

func (o NetworkRuleSetsPtrOutput) ToNetworkRuleSetsPtrOutput() NetworkRuleSetsPtrOutput {
	return o
}

func (o NetworkRuleSetsPtrOutput) ToNetworkRuleSetsPtrOutputWithContext(ctx context.Context) NetworkRuleSetsPtrOutput {
	return o
}

func (o NetworkRuleSetsPtrOutput) Elem() NetworkRuleSetsOutput {
	return o.ApplyT(func(v *NetworkRuleSets) NetworkRuleSets {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSets
		return ret
	}).(NetworkRuleSetsOutput)
}

func (o NetworkRuleSetsPtrOutput) ApplyToDevices() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSets) *bool {
		if v == nil {
			return nil
		}
		return v.ApplyToDevices
	}).(pulumi.BoolPtrOutput)
}

func (o NetworkRuleSetsPtrOutput) ApplyToIoTCentral() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSets) *bool {
		if v == nil {
			return nil
		}
		return v.ApplyToIoTCentral
	}).(pulumi.BoolPtrOutput)
}

func (o NetworkRuleSetsPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSets) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetsPtrOutput) IpRules() NetworkRuleSetIpRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSets) []NetworkRuleSetIpRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(NetworkRuleSetIpRuleArrayOutput)
}

type NetworkRuleSetsResponse struct {
	ApplyToDevices    *bool                          `pulumi:"applyToDevices"`
	ApplyToIoTCentral *bool                          `pulumi:"applyToIoTCentral"`
	DefaultAction     *string                        `pulumi:"defaultAction"`
	IpRules           []NetworkRuleSetIpRuleResponse `pulumi:"ipRules"`
}


func (val *NetworkRuleSetsResponse) Defaults() *NetworkRuleSetsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ApplyToDevices) {
		applyToDevices_ := false
		tmp.ApplyToDevices = &applyToDevices_
	}
	if isZero(tmp.ApplyToIoTCentral) {
		applyToIoTCentral_ := false
		tmp.ApplyToIoTCentral = &applyToIoTCentral_
	}
	return &tmp
}

type NetworkRuleSetsResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetsResponse)(nil)).Elem()
}

func (o NetworkRuleSetsResponseOutput) ToNetworkRuleSetsResponseOutput() NetworkRuleSetsResponseOutput {
	return o
}

func (o NetworkRuleSetsResponseOutput) ToNetworkRuleSetsResponseOutputWithContext(ctx context.Context) NetworkRuleSetsResponseOutput {
	return o
}

func (o NetworkRuleSetsResponseOutput) ApplyToDevices() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetsResponse) *bool { return v.ApplyToDevices }).(pulumi.BoolPtrOutput)
}

func (o NetworkRuleSetsResponseOutput) ApplyToIoTCentral() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetsResponse) *bool { return v.ApplyToIoTCentral }).(pulumi.BoolPtrOutput)
}

func (o NetworkRuleSetsResponseOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRuleSetsResponse) *string { return v.DefaultAction }).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetsResponseOutput) IpRules() NetworkRuleSetIpRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetsResponse) []NetworkRuleSetIpRuleResponse { return v.IpRules }).(NetworkRuleSetIpRuleResponseArrayOutput)
}

type NetworkRuleSetsResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetsResponse)(nil)).Elem()
}

func (o NetworkRuleSetsResponsePtrOutput) ToNetworkRuleSetsResponsePtrOutput() NetworkRuleSetsResponsePtrOutput {
	return o
}

func (o NetworkRuleSetsResponsePtrOutput) ToNetworkRuleSetsResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetsResponsePtrOutput {
	return o
}

func (o NetworkRuleSetsResponsePtrOutput) Elem() NetworkRuleSetsResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetsResponse) NetworkRuleSetsResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetsResponse
		return ret
	}).(NetworkRuleSetsResponseOutput)
}

func (o NetworkRuleSetsResponsePtrOutput) ApplyToDevices() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ApplyToDevices
	}).(pulumi.BoolPtrOutput)
}

func (o NetworkRuleSetsResponsePtrOutput) ApplyToIoTCentral() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ApplyToIoTCentral
	}).(pulumi.BoolPtrOutput)
}

func (o NetworkRuleSetsResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetsResponsePtrOutput) IpRules() NetworkRuleSetIpRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetsResponse) []NetworkRuleSetIpRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(NetworkRuleSetIpRuleResponseArrayOutput)
}

type PrivateEndpointConnectionResponse struct {
	GroupIds                          []string                                  `pulumi:"groupIds"`
	Id                                string                                    `pulumi:"id"`
	Name                              string                                    `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                  `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                    `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                        `pulumi:"systemData"`
	Type                              string                                    `pulumi:"type"`
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

func (o PrivateEndpointConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
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

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
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

type SystemAssignedServiceIdentity struct {
	Type string `pulumi:"type"`
}





type SystemAssignedServiceIdentityInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput
	ToSystemAssignedServiceIdentityOutputWithContext(context.Context) SystemAssignedServiceIdentityOutput
}

type SystemAssignedServiceIdentityArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (SystemAssignedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentity)(nil)).Elem()
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput {
	return i.ToSystemAssignedServiceIdentityOutputWithContext(context.Background())
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityOutput)
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return i.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i SystemAssignedServiceIdentityArgs) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityOutput).ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx)
}









type SystemAssignedServiceIdentityPtrInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput
	ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Context) SystemAssignedServiceIdentityPtrOutput
}

type systemAssignedServiceIdentityPtrType SystemAssignedServiceIdentityArgs

func SystemAssignedServiceIdentityPtr(v *SystemAssignedServiceIdentityArgs) SystemAssignedServiceIdentityPtrInput {
	return (*systemAssignedServiceIdentityPtrType)(v)
}

func (*systemAssignedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentity)(nil)).Elem()
}

func (i *systemAssignedServiceIdentityPtrType) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return i.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *systemAssignedServiceIdentityPtrType) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemAssignedServiceIdentityPtrOutput)
}

type SystemAssignedServiceIdentityOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentity)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityOutput() SystemAssignedServiceIdentityOutput {
	return o
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityOutput {
	return o
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return o.ToSystemAssignedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityOutput) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemAssignedServiceIdentity) *SystemAssignedServiceIdentity {
		return &v
	}).(SystemAssignedServiceIdentityPtrOutput)
}

func (o SystemAssignedServiceIdentityOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentity) string { return v.Type }).(pulumi.StringOutput)
}

type SystemAssignedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentity)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityPtrOutput) ToSystemAssignedServiceIdentityPtrOutput() SystemAssignedServiceIdentityPtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityPtrOutput) ToSystemAssignedServiceIdentityPtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityPtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityPtrOutput) Elem() SystemAssignedServiceIdentityOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentity) SystemAssignedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentity
		return ret
	}).(SystemAssignedServiceIdentityOutput)
}

func (o SystemAssignedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type SystemAssignedServiceIdentityResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}

type SystemAssignedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityResponse)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityResponseOutput) ToSystemAssignedServiceIdentityResponseOutput() SystemAssignedServiceIdentityResponseOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponseOutput) ToSystemAssignedServiceIdentityResponseOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponseOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o SystemAssignedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o SystemAssignedServiceIdentityResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SystemAssignedServiceIdentityResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SystemAssignedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentityResponse)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) ToSystemAssignedServiceIdentityResponsePtrOutput() SystemAssignedServiceIdentityResponsePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) ToSystemAssignedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityResponsePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) Elem() SystemAssignedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) SystemAssignedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentityResponse
		return ret
	}).(SystemAssignedServiceIdentityResponseOutput)
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o SystemAssignedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
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

func init() {
	pulumi.RegisterOutputType(AppSkuInfoOutput{})
	pulumi.RegisterOutputType(AppSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetIpRuleOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetIpRuleArrayOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetIpRuleResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetIpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetsOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetsPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetsResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetsResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
