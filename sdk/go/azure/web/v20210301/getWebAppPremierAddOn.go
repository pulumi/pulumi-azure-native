


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPremierAddOn(ctx *pulumi.Context, args *LookupWebAppPremierAddOnArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPremierAddOnResult, error) {
	var rv LookupWebAppPremierAddOnResult
	err := ctx.Invoke("azure-native:web/v20210301:getWebAppPremierAddOn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppPremierAddOnArgs struct {
	Name              string `pulumi:"name"`
	PremierAddOnName  string `pulumi:"premierAddOnName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppPremierAddOnResult struct {
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

func LookupWebAppPremierAddOnOutput(ctx *pulumi.Context, args LookupWebAppPremierAddOnOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppPremierAddOnResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppPremierAddOnResult, error) {
			args := v.(LookupWebAppPremierAddOnArgs)
			r, err := LookupWebAppPremierAddOn(ctx, &args, opts...)
			var s LookupWebAppPremierAddOnResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppPremierAddOnResultOutput)
}

type LookupWebAppPremierAddOnOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	PremierAddOnName  pulumi.StringInput `pulumi:"premierAddOnName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppPremierAddOnOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPremierAddOnArgs)(nil)).Elem()
}


type LookupWebAppPremierAddOnResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppPremierAddOnResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppPremierAddOnResult)(nil)).Elem()
}

func (o LookupWebAppPremierAddOnResultOutput) ToLookupWebAppPremierAddOnResultOutput() LookupWebAppPremierAddOnResultOutput {
	return o
}

func (o LookupWebAppPremierAddOnResultOutput) ToLookupWebAppPremierAddOnResultOutputWithContext(ctx context.Context) LookupWebAppPremierAddOnResultOutput {
	return o
}

func (o LookupWebAppPremierAddOnResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppPremierAddOnResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPremierAddOnResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupWebAppPremierAddOnResultOutput) MarketplaceOffer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.MarketplaceOffer }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPremierAddOnResultOutput) MarketplacePublisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.MarketplacePublisher }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPremierAddOnResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppPremierAddOnResultOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPremierAddOnResultOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppPremierAddOnResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWebAppPremierAddOnResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWebAppPremierAddOnResultOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppPremierAddOnResult) *string { return v.Vendor }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppPremierAddOnResultOutput{})
}
