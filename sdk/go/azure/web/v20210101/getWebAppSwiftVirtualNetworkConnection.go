


package v20210101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSwiftVirtualNetworkConnection(ctx *pulumi.Context, args *LookupWebAppSwiftVirtualNetworkConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSwiftVirtualNetworkConnectionResult, error) {
	var rv LookupWebAppSwiftVirtualNetworkConnectionResult
	err := ctx.Invoke("azure-native:web/v20210101:getWebAppSwiftVirtualNetworkConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSwiftVirtualNetworkConnectionArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppSwiftVirtualNetworkConnectionResult struct {
	Id               string  `pulumi:"id"`
	Kind             *string `pulumi:"kind"`
	Name             string  `pulumi:"name"`
	SubnetResourceId *string `pulumi:"subnetResourceId"`
	SwiftSupported   *bool   `pulumi:"swiftSupported"`
	Type             string  `pulumi:"type"`
}
