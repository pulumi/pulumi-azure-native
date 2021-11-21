


package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	var rv LookupNetworkResult
	err := ctx.Invoke("azure-native:servicefabricmesh/v20180701preview:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkArgs struct {
	NetworkName       string `pulumi:"networkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNetworkResult struct {
	AddressPrefix     string                 `pulumi:"addressPrefix"`
	Description       *string                `pulumi:"description"`
	Id                string                 `pulumi:"id"`
	IngressConfig     *IngressConfigResponse `pulumi:"ingressConfig"`
	Location          string                 `pulumi:"location"`
	Name              string                 `pulumi:"name"`
	ProvisioningState string                 `pulumi:"provisioningState"`
	Tags              map[string]string      `pulumi:"tags"`
	Type              string                 `pulumi:"type"`
}
