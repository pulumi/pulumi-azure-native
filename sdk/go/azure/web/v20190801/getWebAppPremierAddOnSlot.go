


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppPremierAddOnSlot(ctx *pulumi.Context, args *LookupWebAppPremierAddOnSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppPremierAddOnSlotResult, error) {
	var rv LookupWebAppPremierAddOnSlotResult
	err := ctx.Invoke("azure-native:web/v20190801:getWebAppPremierAddOnSlot", args, &rv, opts...)
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
