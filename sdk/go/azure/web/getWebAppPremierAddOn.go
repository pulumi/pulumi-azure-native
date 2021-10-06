


package web

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPremierAddOn(ctx *pulumi.Context, args *LookupWebAppPremierAddOnArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPremierAddOnResult, error) {
	var rv LookupWebAppPremierAddOnResult
	err := ctx.Invoke("azure-native:web:getWebAppPremierAddOn", args, &rv, opts...)
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
