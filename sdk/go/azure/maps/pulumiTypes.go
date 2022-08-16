


package maps

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreatorPropertiesResponse struct {
	ProvisioningState *string `pulumi:"provisioningState"`
}

type CreatorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CreatorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatorPropertiesResponse)(nil)).Elem()
}

func (o CreatorPropertiesResponseOutput) ToCreatorPropertiesResponseOutput() CreatorPropertiesResponseOutput {
	return o
}

func (o CreatorPropertiesResponseOutput) ToCreatorPropertiesResponseOutputWithContext(ctx context.Context) CreatorPropertiesResponseOutput {
	return o
}

func (o CreatorPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreatorPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type MapsAccountPropertiesResponse struct {
	XMsClientId *string `pulumi:"xMsClientId"`
}

type MapsAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MapsAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MapsAccountPropertiesResponse)(nil)).Elem()
}

func (o MapsAccountPropertiesResponseOutput) ToMapsAccountPropertiesResponseOutput() MapsAccountPropertiesResponseOutput {
	return o
}

func (o MapsAccountPropertiesResponseOutput) ToMapsAccountPropertiesResponseOutputWithContext(ctx context.Context) MapsAccountPropertiesResponseOutput {
	return o
}

func (o MapsAccountPropertiesResponseOutput) XMsClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) *string { return v.XMsClientId }).(pulumi.StringPtrOutput)
}

type PrivateAtlasPropertiesResponse struct {
	ProvisioningState *string `pulumi:"provisioningState"`
}

type PrivateAtlasPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateAtlasPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateAtlasPropertiesResponse)(nil)).Elem()
}

func (o PrivateAtlasPropertiesResponseOutput) ToPrivateAtlasPropertiesResponseOutput() PrivateAtlasPropertiesResponseOutput {
	return o
}

func (o PrivateAtlasPropertiesResponseOutput) ToPrivateAtlasPropertiesResponseOutputWithContext(ctx context.Context) PrivateAtlasPropertiesResponseOutput {
	return o
}

func (o PrivateAtlasPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateAtlasPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
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

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
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

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CreatorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateAtlasPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
}
