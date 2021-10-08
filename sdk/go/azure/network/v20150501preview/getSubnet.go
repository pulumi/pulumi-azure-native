


package v20150501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	var rv LookupSubnetResult
	err := ctx.Invoke("azure-native:network/v20150501preview:getSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubnetArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	SubnetName         string `pulumi:"subnetName"`
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}


type LookupSubnetResult struct {
	AddressPrefix        string                `pulumi:"addressPrefix"`
	Etag                 *string               `pulumi:"etag"`
	Id                   *string               `pulumi:"id"`
	IpConfigurations     []SubResourceResponse `pulumi:"ipConfigurations"`
	Name                 *string               `pulumi:"name"`
	NetworkSecurityGroup *SubResourceResponse  `pulumi:"networkSecurityGroup"`
	ProvisioningState    *string               `pulumi:"provisioningState"`
	RouteTable           *SubResourceResponse  `pulumi:"routeTable"`
}
