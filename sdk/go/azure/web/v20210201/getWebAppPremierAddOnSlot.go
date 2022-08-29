


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPremierAddOnSlot(ctx *pulumi.Context, args *LookupWebAppPremierAddOnSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPremierAddOnSlotResult, error) {
	var rv LookupWebAppPremierAddOnSlotResult
	err := ctx.Invoke("azure-native:web/v20210201:getWebAppPremierAddOnSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPremierAddOnSlotArgs struct {
	Name              string `pulumi:"name"`
	PremierAddOnName  string `pulumi:"premierAddOnName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppPremierAddOnSlotResult struct {
	Id                   string            `pulumi:"id"`
	Kind                 *string           `pulumi:"kind"`
	Location             string            `pulumi:"location"`
	MarketplaceOffer     *string           `pulumi:"marketplaceOffer"`
	MarketplacePublisher *string           `pulumi:"marketplacePublisher"`
	Name                 string            `pulumi:"name"`
	Product              *string           `pulumi:"product"`
	Sku                  *string           `pulumi:"sku"`
	Tags                 map[string]string `pulumi:"tags"`
	Type                 string            `pulumi:"type"`
	Vendor               *string           `pulumi:"vendor"`
}

func LookupWebAppPremierAddOnSlotOutput(ctx *pulumi.Context, args LookupWebAppPremierAddOnSlotOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppPremierAddOnSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppPremierAddOnSlotResult, error) {
			args := v.(LookupWebAppPremierAddOnSlotArgs)
			r, err := LookupWebAppPremierAddOnSlot(ctx, &args, opts...)
			var s LookupWebAppPremierAddOnSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppPremierAddOnSlotResultOutput)
}

type LookupWebAppPremierAddOnSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	PremierAddOnName  pulumi.StringInput `pulumi:"premierAddOnName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (LookupWebAppPremierAddOnSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPremierAddOnSlotArgs)(nil)).Elem()
}


type LookupWebAppPremierAddOnSlotResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppPremierAddOnSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPremierAddOnSlotResult)(nil)).Elem()
}

func (o LookupWebAppPremierAddOnSlotResultOutput) ToLookupWebAppPremierAddOnSlotResultOutput() LookupWebAppPremierAddOnSlotResultOutput {
	return o
}

func (o LookupWebAppPremierAddOnSlotResultOutput) ToLookupWebAppPremierAddOnSlotResultOutputWithContext(ctx context.Context) LookupWebAppPremierAddOnSlotResultOutput {
	return o
}

func (o LookupWebAppPremierAddOnSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppPremierAddOnSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPremierAddOnSlotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWebAppPremierAddOnSlotResultOutput) MarketplaceOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.MarketplaceOffer }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPremierAddOnSlotResultOutput) MarketplacePublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.MarketplacePublisher }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPremierAddOnSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppPremierAddOnSlotResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPremierAddOnSlotResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPremierAddOnSlotResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWebAppPremierAddOnSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppPremierAddOnSlotResultOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnSlotResult) *string { return v.Vendor }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppPremierAddOnSlotResultOutput{})
}
