


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEffectiveVirtualNetworkByNetworkManager(ctx *pulumi.Context, args *ListEffectiveVirtualNetworkByNetworkManagerArgs, opts ...pulumi.InvokeOption) (*ListEffectiveVirtualNetworkByNetworkManagerResult, error) {
	var rv ListEffectiveVirtualNetworkByNetworkManagerResult
	err := ctx.Invoke("azure-native:network/v20210201preview:listEffectiveVirtualNetworkByNetworkManager", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEffectiveVirtualNetworkByNetworkManagerArgs struct {
	ConditionalMembers *string `pulumi:"conditionalMembers"`
	NetworkManagerName string  `pulumi:"networkManagerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
	Top                *int    `pulumi:"top"`
}


type ListEffectiveVirtualNetworkByNetworkManagerResult struct {
	SkipToken *string                           `pulumi:"skipToken"`
	Value     []EffectiveVirtualNetworkResponse `pulumi:"value"`
}
