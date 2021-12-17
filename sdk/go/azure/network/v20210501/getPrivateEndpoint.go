


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateEndpoint(ctx *pulumi.Context, args *LookupPrivateEndpointArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointResult, error) {
	var rv LookupPrivateEndpointResult
	err := ctx.Invoke("azure-native:network/v20210501:getPrivateEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupPrivateEndpointArgs struct {
	Expand              *string `pulumi:"expand"`
	PrivateEndpointName string  `pulumi:"privateEndpointName"`
	ResourceGroupName   string  `pulumi:"resourceGroupName"`
}


type LookupPrivateEndpointResult struct {
	ApplicationSecurityGroups           []ApplicationSecurityGroupResponse        `pulumi:"applicationSecurityGroups"`
	CustomDnsConfigs                    []CustomDnsConfigPropertiesFormatResponse `pulumi:"customDnsConfigs"`
	CustomNetworkInterfaceName          *string                                   `pulumi:"customNetworkInterfaceName"`
	Etag                                string                                    `pulumi:"etag"`
	ExtendedLocation                    *ExtendedLocationResponse                 `pulumi:"extendedLocation"`
	Id                                  *string                                   `pulumi:"id"`
	IpConfigurations                    []PrivateEndpointIPConfigurationResponse  `pulumi:"ipConfigurations"`
	Location                            *string                                   `pulumi:"location"`
	ManualPrivateLinkServiceConnections []PrivateLinkServiceConnectionResponse    `pulumi:"manualPrivateLinkServiceConnections"`
	Name                                string                                    `pulumi:"name"`
	NetworkInterfaces                   []NetworkInterfaceResponse                `pulumi:"networkInterfaces"`
	PrivateLinkServiceConnections       []PrivateLinkServiceConnectionResponse    `pulumi:"privateLinkServiceConnections"`
	ProvisioningState                   string                                    `pulumi:"provisioningState"`
	Subnet                              *SubnetResponse                           `pulumi:"subnet"`
	Tags                                map[string]string                         `pulumi:"tags"`
	Type                                string                                    `pulumi:"type"`
}


func (val *LookupPrivateEndpointResult) Defaults() *LookupPrivateEndpointResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Subnet = tmp.Subnet.Defaults()

	return &tmp
}
