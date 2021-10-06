


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEffectiveVirtualNetworkByNetworkGroup(ctx *pulumi.Context, args *ListEffectiveVirtualNetworkByNetworkGroupArgs, opts ...pulumi.InvokeOption) (*ListEffectiveVirtualNetworkByNetworkGroupResult, error) {
	var rv ListEffectiveVirtualNetworkByNetworkGroupResult
	err := ctx.Invoke("azure-native:network:listEffectiveVirtualNetworkByNetworkGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEffectiveVirtualNetworkByNetworkGroupArgs struct {
	NetworkGroupName   string  `pulumi:"networkGroupName"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
}


type ListEffectiveVirtualNetworkByNetworkGroupResult struct {
	SkipToken *string                           `pulumi:"skipToken"`
	Value     []EffectiveVirtualNetworkResponse `pulumi:"value"`
}
