


package hardwaresecuritymodules

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiEntityReference struct {
	Id *string `pulumi:"id"`
}





type ApiEntityReferenceInput interface {
	pulumi.Input

	ToApiEntityReferenceOutput() ApiEntityReferenceOutput
	ToApiEntityReferenceOutputWithContext(context.Context) ApiEntityReferenceOutput
}

type ApiEntityReferenceArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ApiEntityReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReference)(nil)).Elem()
}

func (i ApiEntityReferenceArgs) ToApiEntityReferenceOutput() ApiEntityReferenceOutput {
	return i.ToApiEntityReferenceOutputWithContext(context.Background())
}

func (i ApiEntityReferenceArgs) ToApiEntityReferenceOutputWithContext(ctx context.Context) ApiEntityReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceOutput)
}

func (i ApiEntityReferenceArgs) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return i.ToApiEntityReferencePtrOutputWithContext(context.Background())
}

func (i ApiEntityReferenceArgs) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceOutput).ToApiEntityReferencePtrOutputWithContext(ctx)
}









type ApiEntityReferencePtrInput interface {
	pulumi.Input

	ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput
	ToApiEntityReferencePtrOutputWithContext(context.Context) ApiEntityReferencePtrOutput
}

type apiEntityReferencePtrType ApiEntityReferenceArgs

func ApiEntityReferencePtr(v *ApiEntityReferenceArgs) ApiEntityReferencePtrInput {
	return (*apiEntityReferencePtrType)(v)
}

func (*apiEntityReferencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityReference)(nil)).Elem()
}

func (i *apiEntityReferencePtrType) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return i.ToApiEntityReferencePtrOutputWithContext(context.Background())
}

func (i *apiEntityReferencePtrType) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferencePtrOutput)
}

type ApiEntityReferenceOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReference)(nil)).Elem()
}

func (o ApiEntityReferenceOutput) ToApiEntityReferenceOutput() ApiEntityReferenceOutput {
	return o
}

func (o ApiEntityReferenceOutput) ToApiEntityReferenceOutputWithContext(ctx context.Context) ApiEntityReferenceOutput {
	return o
}

func (o ApiEntityReferenceOutput) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return o.ToApiEntityReferencePtrOutputWithContext(context.Background())
}

func (o ApiEntityReferenceOutput) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiEntityReference) *ApiEntityReference {
		return &v
	}).(ApiEntityReferencePtrOutput)
}

func (o ApiEntityReferenceOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityReference) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiEntityReferencePtrOutput struct{ *pulumi.OutputState }

func (ApiEntityReferencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityReference)(nil)).Elem()
}

func (o ApiEntityReferencePtrOutput) ToApiEntityReferencePtrOutput() ApiEntityReferencePtrOutput {
	return o
}

func (o ApiEntityReferencePtrOutput) ToApiEntityReferencePtrOutputWithContext(ctx context.Context) ApiEntityReferencePtrOutput {
	return o
}

func (o ApiEntityReferencePtrOutput) Elem() ApiEntityReferenceOutput {
	return o.ApplyT(func(v *ApiEntityReference) ApiEntityReference {
		if v != nil {
			return *v
		}
		var ret ApiEntityReference
		return ret
	}).(ApiEntityReferenceOutput)
}

func (o ApiEntityReferencePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityReference) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ApiEntityReferenceResponse struct {
	Id *string `pulumi:"id"`
}





type ApiEntityReferenceResponseInput interface {
	pulumi.Input

	ToApiEntityReferenceResponseOutput() ApiEntityReferenceResponseOutput
	ToApiEntityReferenceResponseOutputWithContext(context.Context) ApiEntityReferenceResponseOutput
}

type ApiEntityReferenceResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ApiEntityReferenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReferenceResponse)(nil)).Elem()
}

func (i ApiEntityReferenceResponseArgs) ToApiEntityReferenceResponseOutput() ApiEntityReferenceResponseOutput {
	return i.ToApiEntityReferenceResponseOutputWithContext(context.Background())
}

func (i ApiEntityReferenceResponseArgs) ToApiEntityReferenceResponseOutputWithContext(ctx context.Context) ApiEntityReferenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceResponseOutput)
}

func (i ApiEntityReferenceResponseArgs) ToApiEntityReferenceResponsePtrOutput() ApiEntityReferenceResponsePtrOutput {
	return i.ToApiEntityReferenceResponsePtrOutputWithContext(context.Background())
}

func (i ApiEntityReferenceResponseArgs) ToApiEntityReferenceResponsePtrOutputWithContext(ctx context.Context) ApiEntityReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceResponseOutput).ToApiEntityReferenceResponsePtrOutputWithContext(ctx)
}









type ApiEntityReferenceResponsePtrInput interface {
	pulumi.Input

	ToApiEntityReferenceResponsePtrOutput() ApiEntityReferenceResponsePtrOutput
	ToApiEntityReferenceResponsePtrOutputWithContext(context.Context) ApiEntityReferenceResponsePtrOutput
}

type apiEntityReferenceResponsePtrType ApiEntityReferenceResponseArgs

func ApiEntityReferenceResponsePtr(v *ApiEntityReferenceResponseArgs) ApiEntityReferenceResponsePtrInput {
	return (*apiEntityReferenceResponsePtrType)(v)
}

func (*apiEntityReferenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityReferenceResponse)(nil)).Elem()
}

func (i *apiEntityReferenceResponsePtrType) ToApiEntityReferenceResponsePtrOutput() ApiEntityReferenceResponsePtrOutput {
	return i.ToApiEntityReferenceResponsePtrOutputWithContext(context.Background())
}

func (i *apiEntityReferenceResponsePtrType) ToApiEntityReferenceResponsePtrOutputWithContext(ctx context.Context) ApiEntityReferenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiEntityReferenceResponsePtrOutput)
}

type ApiEntityReferenceResponseOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiEntityReferenceResponse)(nil)).Elem()
}

func (o ApiEntityReferenceResponseOutput) ToApiEntityReferenceResponseOutput() ApiEntityReferenceResponseOutput {
	return o
}

func (o ApiEntityReferenceResponseOutput) ToApiEntityReferenceResponseOutputWithContext(ctx context.Context) ApiEntityReferenceResponseOutput {
	return o
}

func (o ApiEntityReferenceResponseOutput) ToApiEntityReferenceResponsePtrOutput() ApiEntityReferenceResponsePtrOutput {
	return o.ToApiEntityReferenceResponsePtrOutputWithContext(context.Background())
}

func (o ApiEntityReferenceResponseOutput) ToApiEntityReferenceResponsePtrOutputWithContext(ctx context.Context) ApiEntityReferenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiEntityReferenceResponse) *ApiEntityReferenceResponse {
		return &v
	}).(ApiEntityReferenceResponsePtrOutput)
}

func (o ApiEntityReferenceResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiEntityReferenceResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiEntityReferenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiEntityReferenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiEntityReferenceResponse)(nil)).Elem()
}

func (o ApiEntityReferenceResponsePtrOutput) ToApiEntityReferenceResponsePtrOutput() ApiEntityReferenceResponsePtrOutput {
	return o
}

func (o ApiEntityReferenceResponsePtrOutput) ToApiEntityReferenceResponsePtrOutputWithContext(ctx context.Context) ApiEntityReferenceResponsePtrOutput {
	return o
}

func (o ApiEntityReferenceResponsePtrOutput) Elem() ApiEntityReferenceResponseOutput {
	return o.ApplyT(func(v *ApiEntityReferenceResponse) ApiEntityReferenceResponse {
		if v != nil {
			return *v
		}
		var ret ApiEntityReferenceResponse
		return ret
	}).(ApiEntityReferenceResponseOutput)
}

func (o ApiEntityReferenceResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiEntityReferenceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type NetworkInterface struct {
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
}





type NetworkInterfaceInput interface {
	pulumi.Input

	ToNetworkInterfaceOutput() NetworkInterfaceOutput
	ToNetworkInterfaceOutputWithContext(context.Context) NetworkInterfaceOutput
}

type NetworkInterfaceArgs struct {
	PrivateIpAddress pulumi.StringPtrInput `pulumi:"privateIpAddress"`
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceArgs) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return i.ToNetworkInterfaceOutputWithContext(context.Background())
}

func (i NetworkInterfaceArgs) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceOutput)
}





type NetworkInterfaceArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput
	ToNetworkInterfaceArrayOutputWithContext(context.Context) NetworkInterfaceArrayOutput
}

type NetworkInterfaceArray []NetworkInterfaceInput

func (NetworkInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterface)(nil)).Elem()
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return i.ToNetworkInterfaceArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceArray) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceArrayOutput)
}

type NetworkInterfaceOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutput() NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) ToNetworkInterfaceOutputWithContext(ctx context.Context) NetworkInterfaceOutput {
	return o
}

func (o NetworkInterfaceOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterface) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterface)(nil)).Elem()
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutput() NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) ToNetworkInterfaceArrayOutputWithContext(ctx context.Context) NetworkInterfaceArrayOutput {
	return o
}

func (o NetworkInterfaceArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterface {
		return vs[0].([]NetworkInterface)[vs[1].(int)]
	}).(NetworkInterfaceOutput)
}

type NetworkInterfaceResponse struct {
	Id               string  `pulumi:"id"`
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
}





type NetworkInterfaceResponseInput interface {
	pulumi.Input

	ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput
	ToNetworkInterfaceResponseOutputWithContext(context.Context) NetworkInterfaceResponseOutput
}

type NetworkInterfaceResponseArgs struct {
	Id               pulumi.StringInput    `pulumi:"id"`
	PrivateIpAddress pulumi.StringPtrInput `pulumi:"privateIpAddress"`
}

func (NetworkInterfaceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResponse)(nil)).Elem()
}

func (i NetworkInterfaceResponseArgs) ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput {
	return i.ToNetworkInterfaceResponseOutputWithContext(context.Background())
}

func (i NetworkInterfaceResponseArgs) ToNetworkInterfaceResponseOutputWithContext(ctx context.Context) NetworkInterfaceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceResponseOutput)
}





type NetworkInterfaceResponseArrayInput interface {
	pulumi.Input

	ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput
	ToNetworkInterfaceResponseArrayOutputWithContext(context.Context) NetworkInterfaceResponseArrayOutput
}

type NetworkInterfaceResponseArray []NetworkInterfaceResponseInput

func (NetworkInterfaceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceResponse)(nil)).Elem()
}

func (i NetworkInterfaceResponseArray) ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput {
	return i.ToNetworkInterfaceResponseArrayOutputWithContext(context.Background())
}

func (i NetworkInterfaceResponseArray) ToNetworkInterfaceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkInterfaceResponseArrayOutput)
}

type NetworkInterfaceResponseOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutput() NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) ToNetworkInterfaceResponseOutputWithContext(ctx context.Context) NetworkInterfaceResponseOutput {
	return o
}

func (o NetworkInterfaceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o NetworkInterfaceResponseOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfaceResponse) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

type NetworkInterfaceResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkInterfaceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkInterfaceResponse)(nil)).Elem()
}

func (o NetworkInterfaceResponseArrayOutput) ToNetworkInterfaceResponseArrayOutput() NetworkInterfaceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceResponseArrayOutput) ToNetworkInterfaceResponseArrayOutputWithContext(ctx context.Context) NetworkInterfaceResponseArrayOutput {
	return o
}

func (o NetworkInterfaceResponseArrayOutput) Index(i pulumi.IntInput) NetworkInterfaceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkInterfaceResponse {
		return vs[0].([]NetworkInterfaceResponse)[vs[1].(int)]
	}).(NetworkInterfaceResponseOutput)
}

type NetworkProfile struct {
	NetworkInterfaces []NetworkInterface  `pulumi:"networkInterfaces"`
	Subnet            *ApiEntityReference `pulumi:"subnet"`
}





type NetworkProfileInput interface {
	pulumi.Input

	ToNetworkProfileOutput() NetworkProfileOutput
	ToNetworkProfileOutputWithContext(context.Context) NetworkProfileOutput
}

type NetworkProfileArgs struct {
	NetworkInterfaces NetworkInterfaceArrayInput `pulumi:"networkInterfaces"`
	Subnet            ApiEntityReferencePtrInput `pulumi:"subnet"`
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (i NetworkProfileArgs) ToNetworkProfileOutput() NetworkProfileOutput {
	return i.ToNetworkProfileOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput)
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i NetworkProfileArgs) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileOutput).ToNetworkProfilePtrOutputWithContext(ctx)
}









type NetworkProfilePtrInput interface {
	pulumi.Input

	ToNetworkProfilePtrOutput() NetworkProfilePtrOutput
	ToNetworkProfilePtrOutputWithContext(context.Context) NetworkProfilePtrOutput
}

type networkProfilePtrType NetworkProfileArgs

func NetworkProfilePtr(v *NetworkProfileArgs) NetworkProfilePtrInput {
	return (*networkProfilePtrType)(v)
}

func (*networkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return i.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *networkProfilePtrType) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfilePtrOutput)
}

type NetworkProfileOutput struct{ *pulumi.OutputState }

func (NetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfile)(nil)).Elem()
}

func (o NetworkProfileOutput) ToNetworkProfileOutput() NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfileOutputWithContext(ctx context.Context) NetworkProfileOutput {
	return o
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o.ToNetworkProfilePtrOutputWithContext(context.Background())
}

func (o NetworkProfileOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfile) *NetworkProfile {
		return &v
	}).(NetworkProfilePtrOutput)
}

func (o NetworkProfileOutput) NetworkInterfaces() NetworkInterfaceArrayOutput {
	return o.ApplyT(func(v NetworkProfile) []NetworkInterface { return v.NetworkInterfaces }).(NetworkInterfaceArrayOutput)
}

func (o NetworkProfileOutput) Subnet() ApiEntityReferencePtrOutput {
	return o.ApplyT(func(v NetworkProfile) *ApiEntityReference { return v.Subnet }).(ApiEntityReferencePtrOutput)
}

type NetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfile)(nil)).Elem()
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutput() NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) ToNetworkProfilePtrOutputWithContext(ctx context.Context) NetworkProfilePtrOutput {
	return o
}

func (o NetworkProfilePtrOutput) Elem() NetworkProfileOutput {
	return o.ApplyT(func(v *NetworkProfile) NetworkProfile {
		if v != nil {
			return *v
		}
		var ret NetworkProfile
		return ret
	}).(NetworkProfileOutput)
}

func (o NetworkProfilePtrOutput) NetworkInterfaces() NetworkInterfaceArrayOutput {
	return o.ApplyT(func(v *NetworkProfile) []NetworkInterface {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceArrayOutput)
}

func (o NetworkProfilePtrOutput) Subnet() ApiEntityReferencePtrOutput {
	return o.ApplyT(func(v *NetworkProfile) *ApiEntityReference {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(ApiEntityReferencePtrOutput)
}

type NetworkProfileResponse struct {
	NetworkInterfaces []NetworkInterfaceResponse  `pulumi:"networkInterfaces"`
	Subnet            *ApiEntityReferenceResponse `pulumi:"subnet"`
}





type NetworkProfileResponseInput interface {
	pulumi.Input

	ToNetworkProfileResponseOutput() NetworkProfileResponseOutput
	ToNetworkProfileResponseOutputWithContext(context.Context) NetworkProfileResponseOutput
}

type NetworkProfileResponseArgs struct {
	NetworkInterfaces NetworkInterfaceResponseArrayInput `pulumi:"networkInterfaces"`
	Subnet            ApiEntityReferenceResponsePtrInput `pulumi:"subnet"`
}

func (NetworkProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (i NetworkProfileResponseArgs) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return i.ToNetworkProfileResponseOutputWithContext(context.Background())
}

func (i NetworkProfileResponseArgs) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileResponseOutput)
}

func (i NetworkProfileResponseArgs) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return i.ToNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i NetworkProfileResponseArgs) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileResponseOutput).ToNetworkProfileResponsePtrOutputWithContext(ctx)
}









type NetworkProfileResponsePtrInput interface {
	pulumi.Input

	ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput
	ToNetworkProfileResponsePtrOutputWithContext(context.Context) NetworkProfileResponsePtrOutput
}

type networkProfileResponsePtrType NetworkProfileResponseArgs

func NetworkProfileResponsePtr(v *NetworkProfileResponseArgs) NetworkProfileResponsePtrInput {
	return (*networkProfileResponsePtrType)(v)
}

func (*networkProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (i *networkProfileResponsePtrType) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return i.ToNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i *networkProfileResponsePtrType) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkProfileResponsePtrOutput)
}

type NetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutput() NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponseOutputWithContext(ctx context.Context) NetworkProfileResponseOutput {
	return o
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o.ToNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (o NetworkProfileResponseOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkProfileResponse) *NetworkProfileResponse {
		return &v
	}).(NetworkProfileResponsePtrOutput)
}

func (o NetworkProfileResponseOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v NetworkProfileResponse) []NetworkInterfaceResponse { return v.NetworkInterfaces }).(NetworkInterfaceResponseArrayOutput)
}

func (o NetworkProfileResponseOutput) Subnet() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v NetworkProfileResponse) *ApiEntityReferenceResponse { return v.Subnet }).(ApiEntityReferenceResponsePtrOutput)
}

type NetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkProfileResponse)(nil)).Elem()
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutput() NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) ToNetworkProfileResponsePtrOutputWithContext(ctx context.Context) NetworkProfileResponsePtrOutput {
	return o
}

func (o NetworkProfileResponsePtrOutput) Elem() NetworkProfileResponseOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) NetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret NetworkProfileResponse
		return ret
	}).(NetworkProfileResponseOutput)
}

func (o NetworkProfileResponsePtrOutput) NetworkInterfaces() NetworkInterfaceResponseArrayOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) []NetworkInterfaceResponse {
		if v == nil {
			return nil
		}
		return v.NetworkInterfaces
	}).(NetworkInterfaceResponseArrayOutput)
}

func (o NetworkProfileResponsePtrOutput) Subnet() ApiEntityReferenceResponsePtrOutput {
	return o.ApplyT(func(v *NetworkProfileResponse) *ApiEntityReferenceResponse {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(ApiEntityReferenceResponsePtrOutput)
}

type Sku struct {
	Name *string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
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

func (o SkuOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Name }).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name *string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
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
	pulumi.RegisterOutputType(ApiEntityReferenceOutput{})
	pulumi.RegisterOutputType(ApiEntityReferencePtrOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceResponseOutput{})
	pulumi.RegisterOutputType(ApiEntityReferenceResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceArrayOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseOutput{})
	pulumi.RegisterOutputType(NetworkInterfaceResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkProfileOutput{})
	pulumi.RegisterOutputType(NetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(NetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
