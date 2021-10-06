


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubnet(ctx *pulumi.Context, args *LookupSubnetArgs, opts ...pulumi.InvokeOption) (*LookupSubnetResult, error) {
	var rv LookupSubnetResult
	err := ctx.Invoke("azure-native:network/v20210301:getSubnet", args, &rv, opts...)
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
	AddressPrefix                      *string                                     `pulumi:"addressPrefix"`
	AddressPrefixes                    []string                                    `pulumi:"addressPrefixes"`
	ApplicationGatewayIpConfigurations []ApplicationGatewayIPConfigurationResponse `pulumi:"applicationGatewayIpConfigurations"`
	Delegations                        []DelegationResponse                        `pulumi:"delegations"`
	Etag                               string                                      `pulumi:"etag"`
	Id                                 *string                                     `pulumi:"id"`
	IpAllocations                      []SubResourceResponse                       `pulumi:"ipAllocations"`
	IpConfigurationProfiles            []IPConfigurationProfileResponse            `pulumi:"ipConfigurationProfiles"`
	IpConfigurations                   []IPConfigurationResponse                   `pulumi:"ipConfigurations"`
	Name                               *string                                     `pulumi:"name"`
	NatGateway                         *SubResourceResponse                        `pulumi:"natGateway"`
	NetworkSecurityGroup               *NetworkSecurityGroupResponse               `pulumi:"networkSecurityGroup"`
	PrivateEndpointNetworkPolicies     *string                                     `pulumi:"privateEndpointNetworkPolicies"`
	PrivateEndpoints                   []PrivateEndpointResponse                   `pulumi:"privateEndpoints"`
	PrivateLinkServiceNetworkPolicies  *string                                     `pulumi:"privateLinkServiceNetworkPolicies"`
	ProvisioningState                  string                                      `pulumi:"provisioningState"`
	Purpose                            string                                      `pulumi:"purpose"`
	ResourceNavigationLinks            []ResourceNavigationLinkResponse            `pulumi:"resourceNavigationLinks"`
	RouteTable                         *RouteTableResponse                         `pulumi:"routeTable"`
	ServiceAssociationLinks            []ServiceAssociationLinkResponse            `pulumi:"serviceAssociationLinks"`
	ServiceEndpointPolicies            []ServiceEndpointPolicyResponse             `pulumi:"serviceEndpointPolicies"`
	ServiceEndpoints                   []ServiceEndpointPropertiesFormatResponse   `pulumi:"serviceEndpoints"`
	Type                               *string                                     `pulumi:"type"`
}
