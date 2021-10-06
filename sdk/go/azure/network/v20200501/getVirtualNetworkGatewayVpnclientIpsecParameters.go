


package v20200501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualNetworkGatewayVpnclientIpsecParameters(ctx *pulumi.Context, args *GetVirtualNetworkGatewayVpnclientIpsecParametersArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayVpnclientIpsecParametersResult, error) {
	var rv GetVirtualNetworkGatewayVpnclientIpsecParametersResult
	err := ctx.Invoke("azure-native:network/v20200501:getVirtualNetworkGatewayVpnclientIpsecParameters", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayVpnclientIpsecParametersArgs struct {
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}


type GetVirtualNetworkGatewayVpnclientIpsecParametersResult struct {
	DhGroup             string `pulumi:"dhGroup"`
	IkeEncryption       string `pulumi:"ikeEncryption"`
	IkeIntegrity        string `pulumi:"ikeIntegrity"`
	IpsecEncryption     string `pulumi:"ipsecEncryption"`
	IpsecIntegrity      string `pulumi:"ipsecIntegrity"`
	PfsGroup            string `pulumi:"pfsGroup"`
	SaDataSizeKilobytes int    `pulumi:"saDataSizeKilobytes"`
	SaLifeTimeSeconds   int    `pulumi:"saLifeTimeSeconds"`
}
