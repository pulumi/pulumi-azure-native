


package v20170801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	var rv LookupSubnetResult
	err := ctx.Invoke("azure-native:network/v20170801:getSubnet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubnetArgs struct {
	Expand             *string `pulumi:"expand"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
	SubnetName         string  `pulumi:"subnetName"`
	VirtualNetworkName string  `pulumi:"virtualNetworkName"`
}


type LookupSubnetResult struct {
	AddressPrefix           *string                                   `pulumi:"addressPrefix"`
	Etag                    *string                                   `pulumi:"etag"`
	Id                      *string                                   `pulumi:"id"`
	IpConfigurations        []IPConfigurationResponse                 `pulumi:"ipConfigurations"`
	Name                    *string                                   `pulumi:"name"`
	NetworkSecurityGroup    *NetworkSecurityGroupResponse             `pulumi:"networkSecurityGroup"`
	ProvisioningState       *string                                   `pulumi:"provisioningState"`
	ResourceNavigationLinks []ResourceNavigationLinkResponse          `pulumi:"resourceNavigationLinks"`
	RouteTable              *RouteTableResponse                       `pulumi:"routeTable"`
	ServiceEndpoints        []ServiceEndpointPropertiesFormatResponse `pulumi:"serviceEndpoints"`
}
