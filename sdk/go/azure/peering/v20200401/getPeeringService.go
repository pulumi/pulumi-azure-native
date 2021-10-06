


package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPeeringService(ctx *pulumi.Context, args *LookupPeeringServiceArgs, opts ...pulumi.InvokeOption) (*LookupPeeringServiceResult, error) {
	var rv LookupPeeringServiceResult
	err := ctx.Invoke("azure-native:peering/v20200401:getPeeringService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPeeringServiceArgs struct {
	PeeringServiceName string `pulumi:"peeringServiceName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
}


type LookupPeeringServiceResult struct {
	Id                     string                     `pulumi:"id"`
	Location               string                     `pulumi:"location"`
	Name                   string                     `pulumi:"name"`
	PeeringServiceLocation *string                    `pulumi:"peeringServiceLocation"`
	PeeringServiceProvider *string                    `pulumi:"peeringServiceProvider"`
	ProvisioningState      string                     `pulumi:"provisioningState"`
	Sku                    *PeeringServiceSkuResponse `pulumi:"sku"`
	Tags                   map[string]string          `pulumi:"tags"`
	Type                   string                     `pulumi:"type"`
}
