


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegisteredAsn(ctx *pulumi.Context, args *LookupRegisteredAsnArgs, opts ...pulumi.InvokeOption) (*LookupRegisteredAsnResult, error) {
	var rv LookupRegisteredAsnResult
	err := ctx.Invoke("azure-native:peering/v20210101:getRegisteredAsn", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRegisteredAsnArgs struct {
	PeeringName       string `pulumi:"peeringName"`
	RegisteredAsnName string `pulumi:"registeredAsnName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRegisteredAsnResult struct {
	Asn                     *int   `pulumi:"asn"`
	Id                      string `pulumi:"id"`
	Name                    string `pulumi:"name"`
	PeeringServicePrefixKey string `pulumi:"peeringServicePrefixKey"`
	ProvisioningState       string `pulumi:"provisioningState"`
	Type                    string `pulumi:"type"`
}
