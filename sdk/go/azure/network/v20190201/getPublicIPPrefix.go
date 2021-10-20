


package v20190201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPublicIPPrefix(ctx *pulumi.Context, args *LookupPublicIPPrefixArgs, opts ...pulumi.InvokeOption) (*LookupPublicIPPrefixResult, error) {
	var rv LookupPublicIPPrefixResult
	err := ctx.Invoke("azure-native:network/v20190201:getPublicIPPrefix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPublicIPPrefixArgs struct {
	Expand             *string `pulumi:"expand"`
	PublicIpPrefixName string  `pulumi:"publicIpPrefixName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupPublicIPPrefixResult struct {
	Etag                                *string                             `pulumi:"etag"`
	Id                                  *string                             `pulumi:"id"`
	IpPrefix                            *string                             `pulumi:"ipPrefix"`
	IpTags                              []IpTagResponse                     `pulumi:"ipTags"`
	LoadBalancerFrontendIpConfiguration SubResourceResponse                 `pulumi:"loadBalancerFrontendIpConfiguration"`
	Location                            *string                             `pulumi:"location"`
	Name                                string                              `pulumi:"name"`
	PrefixLength                        *int                                `pulumi:"prefixLength"`
	ProvisioningState                   *string                             `pulumi:"provisioningState"`
	PublicIPAddressVersion              *string                             `pulumi:"publicIPAddressVersion"`
	PublicIPAddresses                   []ReferencedPublicIpAddressResponse `pulumi:"publicIPAddresses"`
	ResourceGuid                        *string                             `pulumi:"resourceGuid"`
	Sku                                 *PublicIPPrefixSkuResponse          `pulumi:"sku"`
	Tags                                map[string]string                   `pulumi:"tags"`
	Type                                string                              `pulumi:"type"`
	Zones                               []string                            `pulumi:"zones"`
}
