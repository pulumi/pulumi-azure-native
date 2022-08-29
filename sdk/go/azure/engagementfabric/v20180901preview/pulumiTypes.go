


package v20180901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ChannelTypeDescriptionResponse struct {
	ChannelDescription *string  `pulumi:"channelDescription"`
	ChannelFunctions   []string `pulumi:"channelFunctions"`
	ChannelType        *string  `pulumi:"channelType"`
}

type ChannelTypeDescriptionResponseOutput struct{ *pulumi.OutputState }

func (ChannelTypeDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelTypeDescriptionResponse)(nil)).Elem()
}

func (o ChannelTypeDescriptionResponseOutput) ToChannelTypeDescriptionResponseOutput() ChannelTypeDescriptionResponseOutput {
	return o
}

func (o ChannelTypeDescriptionResponseOutput) ToChannelTypeDescriptionResponseOutputWithContext(ctx context.Context) ChannelTypeDescriptionResponseOutput {
	return o
}

func (o ChannelTypeDescriptionResponseOutput) ChannelDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelTypeDescriptionResponse) *string { return v.ChannelDescription }).(pulumi.StringPtrOutput)
}

func (o ChannelTypeDescriptionResponseOutput) ChannelFunctions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ChannelTypeDescriptionResponse) []string { return v.ChannelFunctions }).(pulumi.StringArrayOutput)
}

func (o ChannelTypeDescriptionResponseOutput) ChannelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelTypeDescriptionResponse) *string { return v.ChannelType }).(pulumi.StringPtrOutput)
}

type ChannelTypeDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ChannelTypeDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ChannelTypeDescriptionResponse)(nil)).Elem()
}

func (o ChannelTypeDescriptionResponseArrayOutput) ToChannelTypeDescriptionResponseArrayOutput() ChannelTypeDescriptionResponseArrayOutput {
	return o
}

func (o ChannelTypeDescriptionResponseArrayOutput) ToChannelTypeDescriptionResponseArrayOutputWithContext(ctx context.Context) ChannelTypeDescriptionResponseArrayOutput {
	return o
}

func (o ChannelTypeDescriptionResponseArrayOutput) Index(i pulumi.IntInput) ChannelTypeDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ChannelTypeDescriptionResponse {
		return vs[0].([]ChannelTypeDescriptionResponse)[vs[1].(int)]
	}).(ChannelTypeDescriptionResponseOutput)
}

type KeyDescriptionResponse struct {
	Name  string `pulumi:"name"`
	Rank  string `pulumi:"rank"`
	Value string `pulumi:"value"`
}

type KeyDescriptionResponseOutput struct{ *pulumi.OutputState }

func (KeyDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyDescriptionResponse)(nil)).Elem()
}

func (o KeyDescriptionResponseOutput) ToKeyDescriptionResponseOutput() KeyDescriptionResponseOutput {
	return o
}

func (o KeyDescriptionResponseOutput) ToKeyDescriptionResponseOutputWithContext(ctx context.Context) KeyDescriptionResponseOutput {
	return o
}

func (o KeyDescriptionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KeyDescriptionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o KeyDescriptionResponseOutput) Rank() pulumi.StringOutput {
	return o.ApplyT(func(v KeyDescriptionResponse) string { return v.Rank }).(pulumi.StringOutput)
}

func (o KeyDescriptionResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v KeyDescriptionResponse) string { return v.Value }).(pulumi.StringOutput)
}

type KeyDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyDescriptionResponse)(nil)).Elem()
}

func (o KeyDescriptionResponseArrayOutput) ToKeyDescriptionResponseArrayOutput() KeyDescriptionResponseArrayOutput {
	return o
}

func (o KeyDescriptionResponseArrayOutput) ToKeyDescriptionResponseArrayOutputWithContext(ctx context.Context) KeyDescriptionResponseArrayOutput {
	return o
}

func (o KeyDescriptionResponseArrayOutput) Index(i pulumi.IntInput) KeyDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyDescriptionResponse {
		return vs[0].([]KeyDescriptionResponse)[vs[1].(int)]
	}).(KeyDescriptionResponseOutput)
}

type SKU struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}





type SKUInput interface {
	pulumi.Input

	ToSKUOutput() SKUOutput
	ToSKUOutputWithContext(context.Context) SKUOutput
}

type SKUArgs struct {
	Name pulumi.StringInput    `pulumi:"name"`
	Tier pulumi.StringPtrInput `pulumi:"tier"`
}

func (SKUArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SKU)(nil)).Elem()
}

func (i SKUArgs) ToSKUOutput() SKUOutput {
	return i.ToSKUOutputWithContext(context.Background())
}

func (i SKUArgs) ToSKUOutputWithContext(ctx context.Context) SKUOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SKUOutput)
}

type SKUOutput struct{ *pulumi.OutputState }

func (SKUOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SKU)(nil)).Elem()
}

func (o SKUOutput) ToSKUOutput() SKUOutput {
	return o
}

func (o SKUOutput) ToSKUOutputWithContext(ctx context.Context) SKUOutput {
	return o
}

func (o SKUOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SKU) string { return v.Name }).(pulumi.StringOutput)
}

func (o SKUOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SKU) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SKUResponse struct {
	Name string  `pulumi:"name"`
	Tier *string `pulumi:"tier"`
}

type SKUResponseOutput struct{ *pulumi.OutputState }

func (SKUResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SKUResponse)(nil)).Elem()
}

func (o SKUResponseOutput) ToSKUResponseOutput() SKUResponseOutput {
	return o
}

func (o SKUResponseOutput) ToSKUResponseOutputWithContext(ctx context.Context) SKUResponseOutput {
	return o
}

func (o SKUResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SKUResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SKUResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SKUResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelTypeDescriptionResponseOutput{})
	pulumi.RegisterOutputType(ChannelTypeDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(KeyDescriptionResponseOutput{})
	pulumi.RegisterOutputType(KeyDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SKUOutput{})
	pulumi.RegisterOutputType(SKUResponseOutput{})
}
