


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListEffectiveConnectivityConfiguration(ctx *pulumi.Context, args *ListEffectiveConnectivityConfigurationArgs, opts ...pulumi.InvokeOption) (*ListEffectiveConnectivityConfigurationResult, error) {
	var rv ListEffectiveConnectivityConfigurationResult
	err := ctx.Invoke("azure-native:network/v20210201preview:listEffectiveConnectivityConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListEffectiveConnectivityConfigurationArgs struct {
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SkipToken          *string `pulumi:"skipToken"`
	VirtualNetworkName string  `pulumi:"virtualNetworkName"`
}


type ListEffectiveConnectivityConfigurationResult struct {
	SkipToken *string                                      `pulumi:"skipToken"`
	Value     []EffectiveConnectivityConfigurationResponse `pulumi:"value"`
}
