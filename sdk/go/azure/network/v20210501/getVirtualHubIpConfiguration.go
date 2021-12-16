


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualHubIpConfiguration(ctx *pulumi.Context, args *LookupVirtualHubIpConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubIpConfigurationResult, error) {
	var rv LookupVirtualHubIpConfigurationResult
	err := ctx.Invoke("azure-native:network/v20210501:getVirtualHubIpConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
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


func (val *LookupVirtualHubIpConfigurationResult) Defaults() *LookupVirtualHubIpConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.PublicIPAddress = tmp.PublicIPAddress.Defaults()

	tmp.Subnet = tmp.Subnet.Defaults()

	return &tmp
}
