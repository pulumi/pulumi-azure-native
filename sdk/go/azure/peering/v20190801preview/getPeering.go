


package v20190801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPeering(ctx *pulumi.Context, args *LookupPeeringArgs, opts ...pulumi.InvokeOption) (*LookupPeeringResult, error) {
	var rv LookupPeeringResult
	err := ctx.Invoke("azure-native:peering/v20190801preview:getPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeeringArgs struct {
	PeeringName       string `pulumi:"peeringName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPeeringResult struct {
	Direct            *PeeringPropertiesDirectResponse   `pulumi:"direct"`
	Exchange          *PeeringPropertiesExchangeResponse `pulumi:"exchange"`
	Id                string                             `pulumi:"id"`
	Kind              string                             `pulumi:"kind"`
	Location          string                             `pulumi:"location"`
	Name              string                             `pulumi:"name"`
	PeeringLocation   *string                            `pulumi:"peeringLocation"`
	ProvisioningState string                             `pulumi:"provisioningState"`
	Sku               PeeringSkuResponse                 `pulumi:"sku"`
	Tags              map[string]string                  `pulumi:"tags"`
	Type              string                             `pulumi:"type"`
}
