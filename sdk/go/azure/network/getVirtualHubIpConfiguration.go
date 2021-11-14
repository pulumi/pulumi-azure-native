


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualHubIpConfiguration(ctx *pulumi.Context, args *LookupVirtualHubIpConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubIpConfigurationResult, error) {
	var rv LookupVirtualHubIpConfigurationResult
	err := ctx.Invoke("azure-native:network:getVirtualHubIpConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubIpConfigurationArgs struct {
	IpConfigName      string `pulumi:"ipConfigName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupVirtualHubIpConfigurationResult struct {
	Etag                      string                   `pulumi:"etag"`
	Id                        *string                  `pulumi:"id"`
	Name                      *string                  `pulumi:"name"`
	PrivateIPAddress          *string                  `pulumi:"privateIPAddress"`
	PrivateIPAllocationMethod *string                  `pulumi:"privateIPAllocationMethod"`
	ProvisioningState         string                   `pulumi:"provisioningState"`
	PublicIPAddress           *PublicIPAddressResponse `pulumi:"publicIPAddress"`
	Subnet                    *SubnetResponse          `pulumi:"subnet"`
	Type                      string                   `pulumi:"type"`
}
