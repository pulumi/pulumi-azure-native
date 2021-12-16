


package v20200701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkInterfaceTapConfiguration(ctx *pulumi.Context, args *LookupNetworkInterfaceTapConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceTapConfigurationResult, error) {
	var rv LookupNetworkInterfaceTapConfigurationResult
	err := ctx.Invoke("azure-native:network/v20200701:getNetworkInterfaceTapConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupNetworkInterfaceTapConfigurationArgs struct {
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	TapConfigurationName string `pulumi:"tapConfigurationName"`
}


type LookupNetworkInterfaceTapConfigurationResult struct {
	Etag              string                     `pulumi:"etag"`
	Id                *string                    `pulumi:"id"`
	Name              *string                    `pulumi:"name"`
	ProvisioningState string                     `pulumi:"provisioningState"`
	Type              string                     `pulumi:"type"`
	VirtualNetworkTap *VirtualNetworkTapResponse `pulumi:"virtualNetworkTap"`
}


func (val *LookupNetworkInterfaceTapConfigurationResult) Defaults() *LookupNetworkInterfaceTapConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.VirtualNetworkTap = tmp.VirtualNetworkTap.Defaults()

	return &tmp
}
