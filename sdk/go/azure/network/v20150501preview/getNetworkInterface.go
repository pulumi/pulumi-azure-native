


package v20150501preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:network/v20150501preview:getNetworkInterface", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNetworkInterfaceArgs struct {
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupNetworkInterfaceResult struct {
	DnsSettings          *NetworkInterfaceDnsSettingsResponse      `pulumi:"dnsSettings"`
	EnableIPForwarding   *bool                                     `pulumi:"enableIPForwarding"`
	Etag                 *string                                   `pulumi:"etag"`
	Id                   string                                    `pulumi:"id"`
	IpConfigurations     []NetworkInterfaceIpConfigurationResponse `pulumi:"ipConfigurations"`
	Location             string                                    `pulumi:"location"`
	MacAddress           *string                                   `pulumi:"macAddress"`
	Name                 string                                    `pulumi:"name"`
	NetworkSecurityGroup *SubResourceResponse                      `pulumi:"networkSecurityGroup"`
	Primary              *bool                                     `pulumi:"primary"`
	ProvisioningState    *string                                   `pulumi:"provisioningState"`
	ResourceGuid         *string                                   `pulumi:"resourceGuid"`
	Tags                 map[string]string                         `pulumi:"tags"`
	Type                 string                                    `pulumi:"type"`
	VirtualMachine       *SubResourceResponse                      `pulumi:"virtualMachine"`
}
