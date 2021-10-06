


package v20200401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNatGateway(ctx *pulumi.Context, args *LookupNatGatewayArgs, opts ...pulumi.InvokeOption) (*LookupNatGatewayResult, error) {
	var rv LookupNatGatewayResult
	err := ctx.Invoke("azure-native:network/v20200401:getNatGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNatGatewayArgs struct {
	Expand            *string `pulumi:"expand"`
	NatGatewayName    string  `pulumi:"natGatewayName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type LookupNatGatewayResult struct {
	Etag                 string                 `pulumi:"etag"`
	Id                   *string                `pulumi:"id"`
	IdleTimeoutInMinutes *int                   `pulumi:"idleTimeoutInMinutes"`
	Location             *string                `pulumi:"location"`
	Name                 string                 `pulumi:"name"`
	ProvisioningState    string                 `pulumi:"provisioningState"`
	PublicIpAddresses    []SubResourceResponse  `pulumi:"publicIpAddresses"`
	PublicIpPrefixes     []SubResourceResponse  `pulumi:"publicIpPrefixes"`
	ResourceGuid         string                 `pulumi:"resourceGuid"`
	Sku                  *NatGatewaySkuResponse `pulumi:"sku"`
	Subnets              []SubResourceResponse  `pulumi:"subnets"`
	Tags                 map[string]string      `pulumi:"tags"`
	Type                 string                 `pulumi:"type"`
	Zones                []string               `pulumi:"zones"`
}
