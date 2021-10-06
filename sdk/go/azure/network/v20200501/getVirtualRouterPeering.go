


package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualRouterPeering(ctx *pulumi.Context, args *LookupVirtualRouterPeeringArgs, opts ...pulumi.InvokeOption) (*LookupVirtualRouterPeeringResult, error) {
	var rv LookupVirtualRouterPeeringResult
	err := ctx.Invoke("azure-native:network/v20200501:getVirtualRouterPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualRouterPeeringArgs struct {
	PeeringName       string `pulumi:"peeringName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualRouterName string `pulumi:"virtualRouterName"`
}


type LookupVirtualRouterPeeringResult struct {
	Etag              string   `pulumi:"etag"`
	Id                *string  `pulumi:"id"`
	Name              *string  `pulumi:"name"`
	PeerAsn           *float64 `pulumi:"peerAsn"`
	PeerIp            *string  `pulumi:"peerIp"`
	ProvisioningState string   `pulumi:"provisioningState"`
	Type              string   `pulumi:"type"`
}
