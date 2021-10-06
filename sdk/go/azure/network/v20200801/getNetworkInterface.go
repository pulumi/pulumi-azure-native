


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNetworkInterface(ctx *pulumi.Context, args *LookupNetworkInterfaceArgs, opts ...pulumi.InvokeOption) (*LookupNetworkInterfaceResult, error) {
	var rv LookupNetworkInterfaceResult
	err := ctx.Invoke("azure-native:network/v20200801:getNetworkInterface", args, &rv, opts...)
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
	DnsSettings                 *NetworkInterfaceDnsSettingsResponse       `pulumi:"dnsSettings"`
	DscpConfiguration           SubResourceResponse                        `pulumi:"dscpConfiguration"`
	EnableAcceleratedNetworking *bool                                      `pulumi:"enableAcceleratedNetworking"`
	EnableIPForwarding          *bool                                      `pulumi:"enableIPForwarding"`
	Etag                        string                                     `pulumi:"etag"`
	ExtendedLocation            *ExtendedLocationResponse                  `pulumi:"extendedLocation"`
	HostedWorkloads             []string                                   `pulumi:"hostedWorkloads"`
	Id                          *string                                    `pulumi:"id"`
	IpConfigurations            []NetworkInterfaceIPConfigurationResponse  `pulumi:"ipConfigurations"`
	Location                    *string                                    `pulumi:"location"`
	MacAddress                  string                                     `pulumi:"macAddress"`
	MigrationPhase              *string                                    `pulumi:"migrationPhase"`
	Name                        string                                     `pulumi:"name"`
	NetworkSecurityGroup        *NetworkSecurityGroupResponse              `pulumi:"networkSecurityGroup"`
	NicType                     *string                                    `pulumi:"nicType"`
	Primary                     bool                                       `pulumi:"primary"`
	PrivateEndpoint             PrivateEndpointResponse                    `pulumi:"privateEndpoint"`
	PrivateLinkService          *PrivateLinkServiceResponse                `pulumi:"privateLinkService"`
	ProvisioningState           string                                     `pulumi:"provisioningState"`
	ResourceGuid                string                                     `pulumi:"resourceGuid"`
	Tags                        map[string]string                          `pulumi:"tags"`
	TapConfigurations           []NetworkInterfaceTapConfigurationResponse `pulumi:"tapConfigurations"`
	Type                        string                                     `pulumi:"type"`
	VirtualMachine              SubResourceResponse                        `pulumi:"virtualMachine"`
}
