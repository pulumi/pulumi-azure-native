


package servicefabricmesh

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetwork(ctx *pulumi.Context, args *LookupNetworkArgs, opts ...pulumi.InvokeOption) (*LookupNetworkResult, error) {
	var rv LookupNetworkResult
	err := ctx.Invoke("azure-native:servicefabricmesh:getNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkArgs struct {
	NetworkResourceName string `pulumi:"networkResourceName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupNetworkResult struct {
	Id         string                            `pulumi:"id"`
	Location   string                            `pulumi:"location"`
	Name       string                            `pulumi:"name"`
	Properties NetworkResourcePropertiesResponse `pulumi:"properties"`
	Tags       map[string]string                 `pulumi:"tags"`
	Type       string                            `pulumi:"type"`
}
