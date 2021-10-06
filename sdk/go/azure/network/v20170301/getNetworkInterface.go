


package v20170301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:network/v20170301:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInterfaceArgs struct {
	Expand               *string `pulumi:"expand"`
	NetworkInterfaceName string  `pulumi:"networkInterfaceName"`
	ResourceGroupName    string  `pulumi:"resourceGroupName"`
}


type LookupNetworkInterfaceResult struct {
	DnsSettings                 *NetworkInterfaceDnsSettingsResponse      `pulumi:"dnsSettings"`
	EnableAcceleratedNetworking *bool                                     `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          *bool                                     `pulumi:"enableIPForwarding"`
	Etag                        *string                                   `pulumi:"etag"`
	Id                          *string                                   `pulumi:"id"`
	IpConfigurations            []NetworkInterfaceIPConfigurationResponse `pulumi:"ipConfigurations"`
	Location                    *string                                   `pulumi:"location"`
	MacAddress                  *string                                   `pulumi:"macAddress"`
	Name                        string                                    `pulumi:"name"`
	NetworkSecurityGroup        *NetworkSecurityGroupResponse             `pulumi:"networkSecurityGroup"`
	Primary                     *bool                                     `pulumi:"primary"`
	ProvisioningState           *string                                   `pulumi:"provisioningState"`
	ResourceGuid                *string                                   `pulumi:"resourceGuid"`
	Tags                        map[string]string                         `pulumi:"tags"`
	Type                        string                                    `pulumi:"type"`
	VirtualMachine              *SubResourceResponse                      `pulumi:"virtualMachine"`
}
