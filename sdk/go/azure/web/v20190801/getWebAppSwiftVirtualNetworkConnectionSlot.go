


package v20190801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppSwiftVirtualNetworkConnectionSlot(ctx *pulumi.Context, args *LookupWebAppSwiftVirtualNetworkConnectionSlotArgs, opts ...pulumi.InvokeOption) (*LookupWebAppSwiftVirtualNetworkConnectionSlotResult, error) {
	var rv LookupWebAppSwiftVirtualNetworkConnectionSlotResult
	err := ctx.Invoke("azure-native:web/v20190801:getWebAppSwiftVirtualNetworkConnectionSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebAppSwiftVirtualNetworkConnectionSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type LookupWebAppSwiftVirtualNetworkConnectionSlotResult struct {
	Id               string  `pulumi:"id"`
	Kind             *string `pulumi:"kind"`
	Name             string  `pulumi:"name"`
	SubnetResourceId *string `pulumi:"subnetResourceId"`
	SwiftSupported   *bool   `pulumi:"swiftSupported"`
	Type             string  `pulumi:"type"`
}
