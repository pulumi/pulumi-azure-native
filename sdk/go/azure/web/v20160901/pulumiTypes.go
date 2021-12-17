


package v20160901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Capability struct {
	Name   *string `pulumi:"name"`
	Reason *string `pulumi:"reason"`
	Value  *string `pulumi:"value"`
}





type CapabilityInput interface {
	pulumi.Input

	ToCapabilityOutput() CapabilityOutput
	ToCapabilityOutputWithContext(context.Context) CapabilityOutput
}

type CapabilityArgs struct {
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Reason pulumi.StringPtrInput `pulumi:"reason"`
	Value  pulumi.StringPtrInput `pulumi:"value"`
}

func (CapabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Capability)(nil)).Elem()
}

func (i CapabilityArgs) ToCapabilityOutput() CapabilityOutput {
	return i.ToCapabilityOutputWithContext(context.Background())
}

func (i CapabilityArgs) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityOutput)
}





type CapabilityArrayInput interface {
	pulumi.Input

	ToCapabilityArrayOutput() CapabilityArrayOutput
	ToCapabilityArrayOutputWithContext(context.Context) CapabilityArrayOutput
}

type CapabilityArray []CapabilityInput

func (CapabilityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Capability)(nil)).Elem()
}

func (i CapabilityArray) ToCapabilityArrayOutput() CapabilityArrayOutput {
	return i.ToCapabilityArrayOutputWithContext(context.Background())
}

func (i CapabilityArray) ToCapabilityArrayOutputWithContext(ctx context.Context) CapabilityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityArrayOutput)
}

type CapabilityOutput struct{ *pulumi.OutputState }

func (CapabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Capability)(nil)).Elem()
}

func (o CapabilityOutput) ToCapabilityOutput() CapabilityOutput {
	return o
}

func (o CapabilityOutput) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return o
}

func (o CapabilityOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CapabilityOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

func (o CapabilityOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type CapabilityArrayOutput struct{ *pulumi.OutputState }

func (CapabilityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Capability)(nil)).Elem()
}

func (o CapabilityArrayOutput) ToCapabilityArrayOutput() CapabilityArrayOutput {
	return o
}

func (o CapabilityArrayOutput) ToCapabilityArrayOutputWithContext(ctx context.Context) CapabilityArrayOutput {
	return o
}

func (o CapabilityArrayOutput) Index(i pulumi.IntInput) CapabilityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Capability {
		return vs[0].([]Capability)[vs[1].(int)]
	}).(CapabilityOutput)
}

type CapabilityResponse struct {
	Name   *string `pulumi:"name"`
	Reason *string `pulumi:"reason"`
	Value  *string `pulumi:"value"`
}

type CapabilityResponseOutput struct{ *pulumi.OutputState }

func (CapabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapabilityResponse)(nil)).Elem()
}

func (o CapabilityResponseOutput) ToCapabilityResponseOutput() CapabilityResponseOutput {
	return o
}

func (o CapabilityResponseOutput) ToCapabilityResponseOutputWithContext(ctx context.Context) CapabilityResponseOutput {
	return o
}

func (o CapabilityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o CapabilityResponseOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

func (o CapabilityResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type CapabilityResponseArrayOutput struct{ *pulumi.OutputState }

func (CapabilityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CapabilityResponse)(nil)).Elem()
}

func (o CapabilityResponseArrayOutput) ToCapabilityResponseArrayOutput() CapabilityResponseArrayOutput {
	return o
}

func (o CapabilityResponseArrayOutput) ToCapabilityResponseArrayOutputWithContext(ctx context.Context) CapabilityResponseArrayOutput {
	return o
}

func (o CapabilityResponseArrayOutput) Index(i pulumi.IntInput) CapabilityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CapabilityResponse {
		return vs[0].([]CapabilityResponse)[vs[1].(int)]
	}).(CapabilityResponseOutput)
}

type HostingEnvironmentProfile struct {
	Id *string `pulumi:"id"`
}





type HostingEnvironmentProfileInput interface {
	pulumi.Input

	ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput
	ToHostingEnvironmentProfileOutputWithContext(context.Context) HostingEnvironmentProfileOutput
}

type HostingEnvironmentProfileArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (HostingEnvironmentProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfile)(nil)).Elem()
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput {
	return i.ToHostingEnvironmentProfileOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfileOutputWithContext(ctx context.Context) HostingEnvironmentProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileOutput)
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return i.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileOutput).ToHostingEnvironmentProfilePtrOutputWithContext(ctx)
}









type HostingEnvironmentProfilePtrInput interface {
	pulumi.Input

	ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput
	ToHostingEnvironmentProfilePtrOutputWithContext(context.Context) HostingEnvironmentProfilePtrOutput
}

type hostingEnvironmentProfilePtrType HostingEnvironmentProfileArgs

func HostingEnvironmentProfilePtr(v *HostingEnvironmentProfileArgs) HostingEnvironmentProfilePtrInput {
	return (*hostingEnvironmentProfilePtrType)(v)
}

func (*hostingEnvironmentProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfile)(nil)).Elem()
}

func (i *hostingEnvironmentProfilePtrType) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return i.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i *hostingEnvironmentProfilePtrType) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfilePtrOutput)
}

type HostingEnvironmentProfileOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfile)(nil)).Elem()
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput {
	return o
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfileOutputWithContext(ctx context.Context) HostingEnvironmentProfileOutput {
	return o
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return o.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostingEnvironmentProfile) *HostingEnvironmentProfile {
		return &v
	}).(HostingEnvironmentProfilePtrOutput)
}

func (o HostingEnvironmentProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type HostingEnvironmentProfilePtrOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfile)(nil)).Elem()
}

func (o HostingEnvironmentProfilePtrOutput) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return o
}

func (o HostingEnvironmentProfilePtrOutput) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return o
}

func (o HostingEnvironmentProfilePtrOutput) Elem() HostingEnvironmentProfileOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfile) HostingEnvironmentProfile {
		if v != nil {
			return *v
		}
		var ret HostingEnvironmentProfile
		return ret
	}).(HostingEnvironmentProfileOutput)
}

func (o HostingEnvironmentProfilePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfile) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type HostingEnvironmentProfileResponse struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type HostingEnvironmentProfileResponseOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponseOutput() HostingEnvironmentProfileResponseOutput {
	return o
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponseOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponseOutput {
	return o
}

func (o HostingEnvironmentProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) string { return v.Type }).(pulumi.StringOutput)
}

type HostingEnvironmentProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (o HostingEnvironmentProfileResponsePtrOutput) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return o
}

func (o HostingEnvironmentProfileResponsePtrOutput) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return o
}

func (o HostingEnvironmentProfileResponsePtrOutput) Elem() HostingEnvironmentProfileResponseOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) HostingEnvironmentProfileResponse {
		if v != nil {
			return *v
		}
		var ret HostingEnvironmentProfileResponse
		return ret
	}).(HostingEnvironmentProfileResponseOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type NameValuePair struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type NameValuePairInput interface {
	pulumi.Input

	ToNameValuePairOutput() NameValuePairOutput
	ToNameValuePairOutputWithContext(context.Context) NameValuePairOutput
}

type NameValuePairArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (NameValuePairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePair)(nil)).Elem()
}

func (i NameValuePairArgs) ToNameValuePairOutput() NameValuePairOutput {
	return i.ToNameValuePairOutputWithContext(context.Background())
}

func (i NameValuePairArgs) ToNameValuePairOutputWithContext(ctx context.Context) NameValuePairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameValuePairOutput)
}





type NameValuePairArrayInput interface {
	pulumi.Input

	ToNameValuePairArrayOutput() NameValuePairArrayOutput
	ToNameValuePairArrayOutputWithContext(context.Context) NameValuePairArrayOutput
}

type NameValuePairArray []NameValuePairInput

func (NameValuePairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePair)(nil)).Elem()
}

func (i NameValuePairArray) ToNameValuePairArrayOutput() NameValuePairArrayOutput {
	return i.ToNameValuePairArrayOutputWithContext(context.Background())
}

func (i NameValuePairArray) ToNameValuePairArrayOutputWithContext(ctx context.Context) NameValuePairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameValuePairArrayOutput)
}

type NameValuePairOutput struct{ *pulumi.OutputState }

func (NameValuePairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePair)(nil)).Elem()
}

func (o NameValuePairOutput) ToNameValuePairOutput() NameValuePairOutput {
	return o
}

func (o NameValuePairOutput) ToNameValuePairOutputWithContext(ctx context.Context) NameValuePairOutput {
	return o
}

func (o NameValuePairOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePair) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NameValuePairOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePair) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type NameValuePairArrayOutput struct{ *pulumi.OutputState }

func (NameValuePairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePair)(nil)).Elem()
}

func (o NameValuePairArrayOutput) ToNameValuePairArrayOutput() NameValuePairArrayOutput {
	return o
}

func (o NameValuePairArrayOutput) ToNameValuePairArrayOutputWithContext(ctx context.Context) NameValuePairArrayOutput {
	return o
}

func (o NameValuePairArrayOutput) Index(i pulumi.IntInput) NameValuePairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NameValuePair {
		return vs[0].([]NameValuePair)[vs[1].(int)]
	}).(NameValuePairOutput)
}

type NameValuePairResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type NameValuePairResponseOutput struct{ *pulumi.OutputState }

func (NameValuePairResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePairResponse)(nil)).Elem()
}

func (o NameValuePairResponseOutput) ToNameValuePairResponseOutput() NameValuePairResponseOutput {
	return o
}

func (o NameValuePairResponseOutput) ToNameValuePairResponseOutputWithContext(ctx context.Context) NameValuePairResponseOutput {
	return o
}

func (o NameValuePairResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePairResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NameValuePairResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePairResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type NameValuePairResponseArrayOutput struct{ *pulumi.OutputState }

func (NameValuePairResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePairResponse)(nil)).Elem()
}

func (o NameValuePairResponseArrayOutput) ToNameValuePairResponseArrayOutput() NameValuePairResponseArrayOutput {
	return o
}

func (o NameValuePairResponseArrayOutput) ToNameValuePairResponseArrayOutputWithContext(ctx context.Context) NameValuePairResponseArrayOutput {
	return o
}

func (o NameValuePairResponseArrayOutput) Index(i pulumi.IntInput) NameValuePairResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NameValuePairResponse {
		return vs[0].([]NameValuePairResponse)[vs[1].(int)]
	}).(NameValuePairResponseOutput)
}

type NetworkAccessControlEntry struct {
	Action       *AccessControlEntryAction `pulumi:"action"`
	Description  *string                   `pulumi:"description"`
	Order        *int                      `pulumi:"order"`
	RemoteSubnet *string                   `pulumi:"remoteSubnet"`
}





type NetworkAccessControlEntryInput interface {
	pulumi.Input

	ToNetworkAccessControlEntryOutput() NetworkAccessControlEntryOutput
	ToNetworkAccessControlEntryOutputWithContext(context.Context) NetworkAccessControlEntryOutput
}

type NetworkAccessControlEntryArgs struct {
	Action       AccessControlEntryActionPtrInput `pulumi:"action"`
	Description  pulumi.StringPtrInput            `pulumi:"description"`
	Order        pulumi.IntPtrInput               `pulumi:"order"`
	RemoteSubnet pulumi.StringPtrInput            `pulumi:"remoteSubnet"`
}

func (NetworkAccessControlEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlEntry)(nil)).Elem()
}

func (i NetworkAccessControlEntryArgs) ToNetworkAccessControlEntryOutput() NetworkAccessControlEntryOutput {
	return i.ToNetworkAccessControlEntryOutputWithContext(context.Background())
}

func (i NetworkAccessControlEntryArgs) ToNetworkAccessControlEntryOutputWithContext(ctx context.Context) NetworkAccessControlEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlEntryOutput)
}





type NetworkAccessControlEntryArrayInput interface {
	pulumi.Input

	ToNetworkAccessControlEntryArrayOutput() NetworkAccessControlEntryArrayOutput
	ToNetworkAccessControlEntryArrayOutputWithContext(context.Context) NetworkAccessControlEntryArrayOutput
}

type NetworkAccessControlEntryArray []NetworkAccessControlEntryInput

func (NetworkAccessControlEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkAccessControlEntry)(nil)).Elem()
}

func (i NetworkAccessControlEntryArray) ToNetworkAccessControlEntryArrayOutput() NetworkAccessControlEntryArrayOutput {
	return i.ToNetworkAccessControlEntryArrayOutputWithContext(context.Background())
}

func (i NetworkAccessControlEntryArray) ToNetworkAccessControlEntryArrayOutputWithContext(ctx context.Context) NetworkAccessControlEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlEntryArrayOutput)
}

type NetworkAccessControlEntryOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlEntry)(nil)).Elem()
}

func (o NetworkAccessControlEntryOutput) ToNetworkAccessControlEntryOutput() NetworkAccessControlEntryOutput {
	return o
}

func (o NetworkAccessControlEntryOutput) ToNetworkAccessControlEntryOutputWithContext(ctx context.Context) NetworkAccessControlEntryOutput {
	return o
}

func (o NetworkAccessControlEntryOutput) Action() AccessControlEntryActionPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *AccessControlEntryAction { return v.Action }).(AccessControlEntryActionPtrOutput)
}

func (o NetworkAccessControlEntryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkAccessControlEntryOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o NetworkAccessControlEntryOutput) RemoteSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *string { return v.RemoteSubnet }).(pulumi.StringPtrOutput)
}

type NetworkAccessControlEntryArrayOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkAccessControlEntry)(nil)).Elem()
}

func (o NetworkAccessControlEntryArrayOutput) ToNetworkAccessControlEntryArrayOutput() NetworkAccessControlEntryArrayOutput {
	return o
}

func (o NetworkAccessControlEntryArrayOutput) ToNetworkAccessControlEntryArrayOutputWithContext(ctx context.Context) NetworkAccessControlEntryArrayOutput {
	return o
}

func (o NetworkAccessControlEntryArrayOutput) Index(i pulumi.IntInput) NetworkAccessControlEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkAccessControlEntry {
		return vs[0].([]NetworkAccessControlEntry)[vs[1].(int)]
	}).(NetworkAccessControlEntryOutput)
}

type NetworkAccessControlEntryResponse struct {
	Action       *string `pulumi:"action"`
	Description  *string `pulumi:"description"`
	Order        *int    `pulumi:"order"`
	RemoteSubnet *string `pulumi:"remoteSubnet"`
}

type NetworkAccessControlEntryResponseOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlEntryResponse)(nil)).Elem()
}

func (o NetworkAccessControlEntryResponseOutput) ToNetworkAccessControlEntryResponseOutput() NetworkAccessControlEntryResponseOutput {
	return o
}

func (o NetworkAccessControlEntryResponseOutput) ToNetworkAccessControlEntryResponseOutputWithContext(ctx context.Context) NetworkAccessControlEntryResponseOutput {
	return o
}

func (o NetworkAccessControlEntryResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NetworkAccessControlEntryResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkAccessControlEntryResponseOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o NetworkAccessControlEntryResponseOutput) RemoteSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *string { return v.RemoteSubnet }).(pulumi.StringPtrOutput)
}

type NetworkAccessControlEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkAccessControlEntryResponse)(nil)).Elem()
}

func (o NetworkAccessControlEntryResponseArrayOutput) ToNetworkAccessControlEntryResponseArrayOutput() NetworkAccessControlEntryResponseArrayOutput {
	return o
}

func (o NetworkAccessControlEntryResponseArrayOutput) ToNetworkAccessControlEntryResponseArrayOutputWithContext(ctx context.Context) NetworkAccessControlEntryResponseArrayOutput {
	return o
}

func (o NetworkAccessControlEntryResponseArrayOutput) Index(i pulumi.IntInput) NetworkAccessControlEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkAccessControlEntryResponse {
		return vs[0].([]NetworkAccessControlEntryResponse)[vs[1].(int)]
	}).(NetworkAccessControlEntryResponseOutput)
}

type SkuCapacity struct {
	Default   *int    `pulumi:"default"`
	Maximum   *int    `pulumi:"maximum"`
	Minimum   *int    `pulumi:"minimum"`
	ScaleType *string `pulumi:"scaleType"`
}





type SkuCapacityInput interface {
	pulumi.Input

	ToSkuCapacityOutput() SkuCapacityOutput
	ToSkuCapacityOutputWithContext(context.Context) SkuCapacityOutput
}

type SkuCapacityArgs struct {
	Default   pulumi.IntPtrInput    `pulumi:"default"`
	Maximum   pulumi.IntPtrInput    `pulumi:"maximum"`
	Minimum   pulumi.IntPtrInput    `pulumi:"minimum"`
	ScaleType pulumi.StringPtrInput `pulumi:"scaleType"`
}

func (SkuCapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapacity)(nil)).Elem()
}

func (i SkuCapacityArgs) ToSkuCapacityOutput() SkuCapacityOutput {
	return i.ToSkuCapacityOutputWithContext(context.Background())
}

func (i SkuCapacityArgs) ToSkuCapacityOutputWithContext(ctx context.Context) SkuCapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityOutput)
}

func (i SkuCapacityArgs) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return i.ToSkuCapacityPtrOutputWithContext(context.Background())
}

func (i SkuCapacityArgs) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityOutput).ToSkuCapacityPtrOutputWithContext(ctx)
}









type SkuCapacityPtrInput interface {
	pulumi.Input

	ToSkuCapacityPtrOutput() SkuCapacityPtrOutput
	ToSkuCapacityPtrOutputWithContext(context.Context) SkuCapacityPtrOutput
}

type skuCapacityPtrType SkuCapacityArgs

func SkuCapacityPtr(v *SkuCapacityArgs) SkuCapacityPtrInput {
	return (*skuCapacityPtrType)(v)
}

func (*skuCapacityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuCapacity)(nil)).Elem()
}

func (i *skuCapacityPtrType) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return i.ToSkuCapacityPtrOutputWithContext(context.Background())
}

func (i *skuCapacityPtrType) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityPtrOutput)
}

type SkuCapacityOutput struct{ *pulumi.OutputState }

func (SkuCapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapacity)(nil)).Elem()
}

func (o SkuCapacityOutput) ToSkuCapacityOutput() SkuCapacityOutput {
	return o
}

func (o SkuCapacityOutput) ToSkuCapacityOutputWithContext(ctx context.Context) SkuCapacityOutput {
	return o
}

func (o SkuCapacityOutput) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return o.ToSkuCapacityPtrOutputWithContext(context.Background())
}

func (o SkuCapacityOutput) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuCapacity) *SkuCapacity {
		return &v
	}).(SkuCapacityPtrOutput)
}

func (o SkuCapacityOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.Default }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.Maximum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.Minimum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *string { return v.ScaleType }).(pulumi.StringPtrOutput)
}

type SkuCapacityPtrOutput struct{ *pulumi.OutputState }

func (SkuCapacityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuCapacity)(nil)).Elem()
}

func (o SkuCapacityPtrOutput) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return o
}

func (o SkuCapacityPtrOutput) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return o
}

func (o SkuCapacityPtrOutput) Elem() SkuCapacityOutput {
	return o.ApplyT(func(v *SkuCapacity) SkuCapacity {
		if v != nil {
			return *v
		}
		var ret SkuCapacity
		return ret
	}).(SkuCapacityOutput)
}

func (o SkuCapacityPtrOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.Default
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *string {
		if v == nil {
			return nil
		}
		return v.ScaleType
	}).(pulumi.StringPtrOutput)
}

type SkuCapacityResponse struct {
	Default   *int    `pulumi:"default"`
	Maximum   *int    `pulumi:"maximum"`
	Minimum   *int    `pulumi:"minimum"`
	ScaleType *string `pulumi:"scaleType"`
}

type SkuCapacityResponseOutput struct{ *pulumi.OutputState }

func (SkuCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapacityResponse)(nil)).Elem()
}

func (o SkuCapacityResponseOutput) ToSkuCapacityResponseOutput() SkuCapacityResponseOutput {
	return o
}

func (o SkuCapacityResponseOutput) ToSkuCapacityResponseOutputWithContext(ctx context.Context) SkuCapacityResponseOutput {
	return o
}

func (o SkuCapacityResponseOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.Default }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.Maximum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.Minimum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *string { return v.ScaleType }).(pulumi.StringPtrOutput)
}

type SkuCapacityResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuCapacityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuCapacityResponse)(nil)).Elem()
}

func (o SkuCapacityResponsePtrOutput) ToSkuCapacityResponsePtrOutput() SkuCapacityResponsePtrOutput {
	return o
}

func (o SkuCapacityResponsePtrOutput) ToSkuCapacityResponsePtrOutputWithContext(ctx context.Context) SkuCapacityResponsePtrOutput {
	return o
}

func (o SkuCapacityResponsePtrOutput) Elem() SkuCapacityResponseOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) SkuCapacityResponse {
		if v != nil {
			return *v
		}
		var ret SkuCapacityResponse
		return ret
	}).(SkuCapacityResponseOutput)
}

func (o SkuCapacityResponsePtrOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.Default
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScaleType
	}).(pulumi.StringPtrOutput)
}

type SkuDescription struct {
	Capabilities []Capability `pulumi:"capabilities"`
	Capacity     *int         `pulumi:"capacity"`
	Family       *string      `pulumi:"family"`
	Locations    []string     `pulumi:"locations"`
	Name         *string      `pulumi:"name"`
	Size         *string      `pulumi:"size"`
	SkuCapacity  *SkuCapacity `pulumi:"skuCapacity"`
	Tier         *string      `pulumi:"tier"`
}





type SkuDescriptionInput interface {
	pulumi.Input

	ToSkuDescriptionOutput() SkuDescriptionOutput
	ToSkuDescriptionOutputWithContext(context.Context) SkuDescriptionOutput
}

type SkuDescriptionArgs struct {
	Capabilities CapabilityArrayInput    `pulumi:"capabilities"`
	Capacity     pulumi.IntPtrInput      `pulumi:"capacity"`
	Family       pulumi.StringPtrInput   `pulumi:"family"`
	Locations    pulumi.StringArrayInput `pulumi:"locations"`
	Name         pulumi.StringPtrInput   `pulumi:"name"`
	Size         pulumi.StringPtrInput   `pulumi:"size"`
	SkuCapacity  SkuCapacityPtrInput     `pulumi:"skuCapacity"`
	Tier         pulumi.StringPtrInput   `pulumi:"tier"`
}

func (SkuDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescription)(nil)).Elem()
}

func (i SkuDescriptionArgs) ToSkuDescriptionOutput() SkuDescriptionOutput {
	return i.ToSkuDescriptionOutputWithContext(context.Background())
}

func (i SkuDescriptionArgs) ToSkuDescriptionOutputWithContext(ctx context.Context) SkuDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionOutput)
}

func (i SkuDescriptionArgs) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return i.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (i SkuDescriptionArgs) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionOutput).ToSkuDescriptionPtrOutputWithContext(ctx)
}









type SkuDescriptionPtrInput interface {
	pulumi.Input

	ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput
	ToSkuDescriptionPtrOutputWithContext(context.Context) SkuDescriptionPtrOutput
}

type skuDescriptionPtrType SkuDescriptionArgs

func SkuDescriptionPtr(v *SkuDescriptionArgs) SkuDescriptionPtrInput {
	return (*skuDescriptionPtrType)(v)
}

func (*skuDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescription)(nil)).Elem()
}

func (i *skuDescriptionPtrType) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return i.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (i *skuDescriptionPtrType) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionPtrOutput)
}

type SkuDescriptionOutput struct{ *pulumi.OutputState }

func (SkuDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescription)(nil)).Elem()
}

func (o SkuDescriptionOutput) ToSkuDescriptionOutput() SkuDescriptionOutput {
	return o
}

func (o SkuDescriptionOutput) ToSkuDescriptionOutputWithContext(ctx context.Context) SkuDescriptionOutput {
	return o
}

func (o SkuDescriptionOutput) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return o.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (o SkuDescriptionOutput) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuDescription) *SkuDescription {
		return &v
	}).(SkuDescriptionPtrOutput)
}

func (o SkuDescriptionOutput) Capabilities() CapabilityArrayOutput {
	return o.ApplyT(func(v SkuDescription) []Capability { return v.Capabilities }).(CapabilityArrayOutput)
}

func (o SkuDescriptionOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuDescription) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SkuDescription) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) SkuCapacity() SkuCapacityPtrOutput {
	return o.ApplyT(func(v SkuDescription) *SkuCapacity { return v.SkuCapacity }).(SkuCapacityPtrOutput)
}

func (o SkuDescriptionOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuDescriptionPtrOutput struct{ *pulumi.OutputState }

func (SkuDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescription)(nil)).Elem()
}

func (o SkuDescriptionPtrOutput) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return o
}

func (o SkuDescriptionPtrOutput) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return o
}

func (o SkuDescriptionPtrOutput) Elem() SkuDescriptionOutput {
	return o.ApplyT(func(v *SkuDescription) SkuDescription {
		if v != nil {
			return *v
		}
		var ret SkuDescription
		return ret
	}).(SkuDescriptionOutput)
}

func (o SkuDescriptionPtrOutput) Capabilities() CapabilityArrayOutput {
	return o.ApplyT(func(v *SkuDescription) []Capability {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(CapabilityArrayOutput)
}

func (o SkuDescriptionPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SkuDescription) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) SkuCapacity() SkuCapacityPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *SkuCapacity {
		if v == nil {
			return nil
		}
		return v.SkuCapacity
	}).(SkuCapacityPtrOutput)
}

func (o SkuDescriptionPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuDescriptionResponse struct {
	Capabilities []CapabilityResponse `pulumi:"capabilities"`
	Capacity     *int                 `pulumi:"capacity"`
	Family       *string              `pulumi:"family"`
	Locations    []string             `pulumi:"locations"`
	Name         *string              `pulumi:"name"`
	Size         *string              `pulumi:"size"`
	SkuCapacity  *SkuCapacityResponse `pulumi:"skuCapacity"`
	Tier         *string              `pulumi:"tier"`
}

type SkuDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SkuDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescriptionResponse)(nil)).Elem()
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponseOutput() SkuDescriptionResponseOutput {
	return o
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponseOutputWithContext(ctx context.Context) SkuDescriptionResponseOutput {
	return o
}

func (o SkuDescriptionResponseOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) []CapabilityResponse { return v.Capabilities }).(CapabilityResponseArrayOutput)
}

func (o SkuDescriptionResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) SkuCapacity() SkuCapacityResponsePtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *SkuCapacityResponse { return v.SkuCapacity }).(SkuCapacityResponsePtrOutput)
}

func (o SkuDescriptionResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescriptionResponse)(nil)).Elem()
}

func (o SkuDescriptionResponsePtrOutput) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return o
}

func (o SkuDescriptionResponsePtrOutput) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return o
}

func (o SkuDescriptionResponsePtrOutput) Elem() SkuDescriptionResponseOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) SkuDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret SkuDescriptionResponse
		return ret
	}).(SkuDescriptionResponseOutput)
}

func (o SkuDescriptionResponsePtrOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) []CapabilityResponse {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(CapabilityResponseArrayOutput)
}

func (o SkuDescriptionResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) SkuCapacity() SkuCapacityResponsePtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *SkuCapacityResponse {
		if v == nil {
			return nil
		}
		return v.SkuCapacity
	}).(SkuCapacityResponsePtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type StampCapacityResponse struct {
	AvailableCapacity              *float64 `pulumi:"availableCapacity"`
	ComputeMode                    *string  `pulumi:"computeMode"`
	ExcludeFromCapacityAllocation  *bool    `pulumi:"excludeFromCapacityAllocation"`
	IsApplicableForAllComputeModes *bool    `pulumi:"isApplicableForAllComputeModes"`
	Name                           *string  `pulumi:"name"`
	SiteMode                       *string  `pulumi:"siteMode"`
	TotalCapacity                  *float64 `pulumi:"totalCapacity"`
	Unit                           *string  `pulumi:"unit"`
	WorkerSize                     *string  `pulumi:"workerSize"`
	WorkerSizeId                   *int     `pulumi:"workerSizeId"`
}

type StampCapacityResponseOutput struct{ *pulumi.OutputState }

func (StampCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StampCapacityResponse)(nil)).Elem()
}

func (o StampCapacityResponseOutput) ToStampCapacityResponseOutput() StampCapacityResponseOutput {
	return o
}

func (o StampCapacityResponseOutput) ToStampCapacityResponseOutputWithContext(ctx context.Context) StampCapacityResponseOutput {
	return o
}

func (o StampCapacityResponseOutput) AvailableCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *float64 { return v.AvailableCapacity }).(pulumi.Float64PtrOutput)
}

func (o StampCapacityResponseOutput) ComputeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.ComputeMode }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) ExcludeFromCapacityAllocation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *bool { return v.ExcludeFromCapacityAllocation }).(pulumi.BoolPtrOutput)
}

func (o StampCapacityResponseOutput) IsApplicableForAllComputeModes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *bool { return v.IsApplicableForAllComputeModes }).(pulumi.BoolPtrOutput)
}

func (o StampCapacityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) SiteMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.SiteMode }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) TotalCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *float64 { return v.TotalCapacity }).(pulumi.Float64PtrOutput)
}

func (o StampCapacityResponseOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) WorkerSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.WorkerSize }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) WorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *int { return v.WorkerSizeId }).(pulumi.IntPtrOutput)
}

type StampCapacityResponseArrayOutput struct{ *pulumi.OutputState }

func (StampCapacityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StampCapacityResponse)(nil)).Elem()
}

func (o StampCapacityResponseArrayOutput) ToStampCapacityResponseArrayOutput() StampCapacityResponseArrayOutput {
	return o
}

func (o StampCapacityResponseArrayOutput) ToStampCapacityResponseArrayOutputWithContext(ctx context.Context) StampCapacityResponseArrayOutput {
	return o
}

func (o StampCapacityResponseArrayOutput) Index(i pulumi.IntInput) StampCapacityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StampCapacityResponse {
		return vs[0].([]StampCapacityResponse)[vs[1].(int)]
	}).(StampCapacityResponseOutput)
}

type VirtualIPMappingResponse struct {
	InUse             *bool   `pulumi:"inUse"`
	InternalHttpPort  *int    `pulumi:"internalHttpPort"`
	InternalHttpsPort *int    `pulumi:"internalHttpsPort"`
	VirtualIP         *string `pulumi:"virtualIP"`
}

type VirtualIPMappingResponseOutput struct{ *pulumi.OutputState }

func (VirtualIPMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualIPMappingResponse)(nil)).Elem()
}

func (o VirtualIPMappingResponseOutput) ToVirtualIPMappingResponseOutput() VirtualIPMappingResponseOutput {
	return o
}

func (o VirtualIPMappingResponseOutput) ToVirtualIPMappingResponseOutputWithContext(ctx context.Context) VirtualIPMappingResponseOutput {
	return o
}

func (o VirtualIPMappingResponseOutput) InUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *bool { return v.InUse }).(pulumi.BoolPtrOutput)
}

func (o VirtualIPMappingResponseOutput) InternalHttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *int { return v.InternalHttpPort }).(pulumi.IntPtrOutput)
}

func (o VirtualIPMappingResponseOutput) InternalHttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *int { return v.InternalHttpsPort }).(pulumi.IntPtrOutput)
}

func (o VirtualIPMappingResponseOutput) VirtualIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *string { return v.VirtualIP }).(pulumi.StringPtrOutput)
}

type VirtualIPMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualIPMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualIPMappingResponse)(nil)).Elem()
}

func (o VirtualIPMappingResponseArrayOutput) ToVirtualIPMappingResponseArrayOutput() VirtualIPMappingResponseArrayOutput {
	return o
}

func (o VirtualIPMappingResponseArrayOutput) ToVirtualIPMappingResponseArrayOutputWithContext(ctx context.Context) VirtualIPMappingResponseArrayOutput {
	return o
}

func (o VirtualIPMappingResponseArrayOutput) Index(i pulumi.IntInput) VirtualIPMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualIPMappingResponse {
		return vs[0].([]VirtualIPMappingResponse)[vs[1].(int)]
	}).(VirtualIPMappingResponseOutput)
}

type VirtualNetworkProfile struct {
	Id     *string `pulumi:"id"`
	Subnet *string `pulumi:"subnet"`
}





type VirtualNetworkProfileInput interface {
	pulumi.Input

	ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput
	ToVirtualNetworkProfileOutputWithContext(context.Context) VirtualNetworkProfileOutput
}

type VirtualNetworkProfileArgs struct {
	Id     pulumi.StringPtrInput `pulumi:"id"`
	Subnet pulumi.StringPtrInput `pulumi:"subnet"`
}

func (VirtualNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return i.ToVirtualNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileOutput)
}

type VirtualNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfileResponse struct {
	Id     *string `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Subnet *string `pulumi:"subnet"`
	Type   string  `pulumi:"type"`
}

type VirtualNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutput() VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutputWithContext(ctx context.Context) VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkProfileResponseOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) string { return v.Type }).(pulumi.StringOutput)
}

type WorkerPool struct {
	ComputeMode  *ComputeModeOptions `pulumi:"computeMode"`
	WorkerCount  *int                `pulumi:"workerCount"`
	WorkerSize   *string             `pulumi:"workerSize"`
	WorkerSizeId *int                `pulumi:"workerSizeId"`
}





type WorkerPoolInput interface {
	pulumi.Input

	ToWorkerPoolOutput() WorkerPoolOutput
	ToWorkerPoolOutputWithContext(context.Context) WorkerPoolOutput
}

type WorkerPoolArgs struct {
	ComputeMode  ComputeModeOptionsPtrInput `pulumi:"computeMode"`
	WorkerCount  pulumi.IntPtrInput         `pulumi:"workerCount"`
	WorkerSize   pulumi.StringPtrInput      `pulumi:"workerSize"`
	WorkerSizeId pulumi.IntPtrInput         `pulumi:"workerSizeId"`
}

func (WorkerPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPool)(nil)).Elem()
}

func (i WorkerPoolArgs) ToWorkerPoolOutput() WorkerPoolOutput {
	return i.ToWorkerPoolOutputWithContext(context.Background())
}

func (i WorkerPoolArgs) ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolOutput)
}





type WorkerPoolArrayInput interface {
	pulumi.Input

	ToWorkerPoolArrayOutput() WorkerPoolArrayOutput
	ToWorkerPoolArrayOutputWithContext(context.Context) WorkerPoolArrayOutput
}

type WorkerPoolArray []WorkerPoolInput

func (WorkerPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPool)(nil)).Elem()
}

func (i WorkerPoolArray) ToWorkerPoolArrayOutput() WorkerPoolArrayOutput {
	return i.ToWorkerPoolArrayOutputWithContext(context.Background())
}

func (i WorkerPoolArray) ToWorkerPoolArrayOutputWithContext(ctx context.Context) WorkerPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolArrayOutput)
}

type WorkerPoolOutput struct{ *pulumi.OutputState }

func (WorkerPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPool)(nil)).Elem()
}

func (o WorkerPoolOutput) ToWorkerPoolOutput() WorkerPoolOutput {
	return o
}

func (o WorkerPoolOutput) ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput {
	return o
}

func (o WorkerPoolOutput) ComputeMode() ComputeModeOptionsPtrOutput {
	return o.ApplyT(func(v WorkerPool) *ComputeModeOptions { return v.ComputeMode }).(ComputeModeOptionsPtrOutput)
}

func (o WorkerPoolOutput) WorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPool) *int { return v.WorkerCount }).(pulumi.IntPtrOutput)
}

func (o WorkerPoolOutput) WorkerSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPool) *string { return v.WorkerSize }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolOutput) WorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPool) *int { return v.WorkerSizeId }).(pulumi.IntPtrOutput)
}

type WorkerPoolArrayOutput struct{ *pulumi.OutputState }

func (WorkerPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPool)(nil)).Elem()
}

func (o WorkerPoolArrayOutput) ToWorkerPoolArrayOutput() WorkerPoolArrayOutput {
	return o
}

func (o WorkerPoolArrayOutput) ToWorkerPoolArrayOutputWithContext(ctx context.Context) WorkerPoolArrayOutput {
	return o
}

func (o WorkerPoolArrayOutput) Index(i pulumi.IntInput) WorkerPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkerPool {
		return vs[0].([]WorkerPool)[vs[1].(int)]
	}).(WorkerPoolOutput)
}

type WorkerPoolResponse struct {
	ComputeMode   *string  `pulumi:"computeMode"`
	InstanceNames []string `pulumi:"instanceNames"`
	WorkerCount   *int     `pulumi:"workerCount"`
	WorkerSize    *string  `pulumi:"workerSize"`
	WorkerSizeId  *int     `pulumi:"workerSizeId"`
}

type WorkerPoolResponseOutput struct{ *pulumi.OutputState }

func (WorkerPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPoolResponse)(nil)).Elem()
}

func (o WorkerPoolResponseOutput) ToWorkerPoolResponseOutput() WorkerPoolResponseOutput {
	return o
}

func (o WorkerPoolResponseOutput) ToWorkerPoolResponseOutputWithContext(ctx context.Context) WorkerPoolResponseOutput {
	return o
}

func (o WorkerPoolResponseOutput) ComputeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *string { return v.ComputeMode }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolResponseOutput) InstanceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkerPoolResponse) []string { return v.InstanceNames }).(pulumi.StringArrayOutput)
}

func (o WorkerPoolResponseOutput) WorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *int { return v.WorkerCount }).(pulumi.IntPtrOutput)
}

func (o WorkerPoolResponseOutput) WorkerSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *string { return v.WorkerSize }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolResponseOutput) WorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *int { return v.WorkerSizeId }).(pulumi.IntPtrOutput)
}

type WorkerPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkerPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPoolResponse)(nil)).Elem()
}

func (o WorkerPoolResponseArrayOutput) ToWorkerPoolResponseArrayOutput() WorkerPoolResponseArrayOutput {
	return o
}

func (o WorkerPoolResponseArrayOutput) ToWorkerPoolResponseArrayOutputWithContext(ctx context.Context) WorkerPoolResponseArrayOutput {
	return o
}

func (o WorkerPoolResponseArrayOutput) Index(i pulumi.IntInput) WorkerPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkerPoolResponse {
		return vs[0].([]WorkerPoolResponse)[vs[1].(int)]
	}).(WorkerPoolResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CapabilityOutput{})
	pulumi.RegisterOutputType(CapabilityArrayOutput{})
	pulumi.RegisterOutputType(CapabilityResponseOutput{})
	pulumi.RegisterOutputType(CapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfilePtrOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileResponseOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(NameValuePairOutput{})
	pulumi.RegisterOutputType(NameValuePairArrayOutput{})
	pulumi.RegisterOutputType(NameValuePairResponseOutput{})
	pulumi.RegisterOutputType(NameValuePairResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryArrayOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryResponseOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuCapacityOutput{})
	pulumi.RegisterOutputType(SkuCapacityPtrOutput{})
	pulumi.RegisterOutputType(SkuCapacityResponseOutput{})
	pulumi.RegisterOutputType(SkuCapacityResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuDescriptionOutput{})
	pulumi.RegisterOutputType(SkuDescriptionPtrOutput{})
	pulumi.RegisterOutputType(SkuDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SkuDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(StampCapacityResponseOutput{})
	pulumi.RegisterOutputType(StampCapacityResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualIPMappingResponseOutput{})
	pulumi.RegisterOutputType(VirtualIPMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(WorkerPoolOutput{})
	pulumi.RegisterOutputType(WorkerPoolArrayOutput{})
	pulumi.RegisterOutputType(WorkerPoolResponseOutput{})
	pulumi.RegisterOutputType(WorkerPoolResponseArrayOutput{})
}
