


package v20210201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListActiveConnectivityConfiguration(ctx *pulumi.Context, args *ListActiveConnectivityConfigurationArgs, opts ...pulumi.InvokeOption) (*ListActiveConnectivityConfigurationResult, error) {
	var rv ListActiveConnectivityConfigurationResult
	err := ctx.Invoke("azure-native:network/v20210201preview:listActiveConnectivityConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListActiveConnectivityConfigurationArgs struct {
	NetworkManagerName string   `pulumi:"networkManagerName"`
	Regions            []string `pulumi:"regions"`
	ResourceGroupName  string   `pulumi:"resourceGroupName"`
	SkipToken          *string  `pulumi:"skipToken"`
}


type ListActiveConnectivityConfigurationResult struct {
	SkipToken *string                                   `pulumi:"skipToken"`
	Value     []ActiveConnectivityConfigurationResponse `pulumi:"value"`
}
