


package engagementfabric

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

type KeyDescriptionResponse struct {
	Name  string `pulumi:"name"`
	Rank  string `pulumi:"rank"`
	Value string `pulumi:"value"`
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
	pulumi.RegisterOutputType(SKUOutput{})
	pulumi.RegisterOutputType(SKUResponseOutput{})
}
