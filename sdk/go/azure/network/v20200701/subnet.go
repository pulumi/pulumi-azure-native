


package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Subnet struct {
	pulumi.CustomResourceState

	AddressPrefix                     pulumi.StringPtrOutput                             `pulumi:"addressPrefix"`
	AddressPrefixes                   pulumi.StringArrayOutput                           `pulumi:"addressPrefixes"`
	Delegations                       DelegationResponseArrayOutput                      `pulumi:"delegations"`
	Etag                              pulumi.StringOutput                                `pulumi:"etag"`
	IpAllocations                     SubResourceResponseArrayOutput                     `pulumi:"ipAllocations"`
	IpConfigurationProfiles           IPConfigurationProfileResponseArrayOutput          `pulumi:"ipConfigurationProfiles"`
	IpConfigurations                  IPConfigurationResponseArrayOutput                 `pulumi:"ipConfigurations"`
	Name                              pulumi.StringPtrOutput                             `pulumi:"name"`
	NatGateway                        SubResourceResponsePtrOutput                       `pulumi:"natGateway"`
	NetworkSecurityGroup              NetworkSecurityGroupResponsePtrOutput              `pulumi:"networkSecurityGroup"`
	PrivateEndpointNetworkPolicies    pulumi.StringPtrOutput                             `pulumi:"privateEndpointNetworkPolicies"`
	PrivateEndpoints                  PrivateEndpointResponseArrayOutput                 `pulumi:"privateEndpoints"`
	PrivateLinkServiceNetworkPolicies pulumi.StringPtrOutput                             `pulumi:"privateLinkServiceNetworkPolicies"`
	ProvisioningState                 pulumi.StringOutput                                `pulumi:"provisioningState"`
	Purpose                           pulumi.StringOutput                                `pulumi:"purpose"`
	ResourceNavigationLinks           ResourceNavigationLinkResponseArrayOutput          `pulumi:"resourceNavigationLinks"`
	RouteTable                        RouteTableResponsePtrOutput                        `pulumi:"routeTable"`
	ServiceAssociationLinks           ServiceAssociationLinkResponseArrayOutput          `pulumi:"serviceAssociationLinks"`
	ServiceEndpointPolicies           ServiceEndpointPolicyResponseArrayOutput           `pulumi:"serviceEndpointPolicies"`
	ServiceEndpoints                  ServiceEndpointPropertiesFormatResponseArrayOutput `pulumi:"serviceEndpoints"`
	Type                              pulumi.StringPtrOutput                             `pulumi:"type"`
}


func NewSubnet(ctx *pulumi.Context,
	name string, args *SubnetArgs, opts ...pulumi.ResourceOption) (*Subnet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetworkName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetworkName'")
	}
	if isZero(args.PrivateEndpointNetworkPolicies) {
		args.PrivateEndpointNetworkPolicies = pulumi.StringPtr("Enabled")
	}
	if isZero(args.PrivateLinkServiceNetworkPolicies) {
		args.PrivateLinkServiceNetworkPolicies = pulumi.StringPtr("Enabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:Subnet"),
		},
	})
	opts = append(opts, aliases)
	var resource Subnet
	err := ctx.RegisterResource("azure-native:network/v20200701:Subnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetState, opts ...pulumi.ResourceOption) (*Subnet, error) {
	var resource Subnet
	err := ctx.ReadResource("azure-native:network/v20200701:Subnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subnetState struct {
}

type SubnetState struct {
}

func (SubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetState)(nil)).Elem()
}

type subnetArgs struct {
	AddressPrefix                     *string                           `pulumi:"addressPrefix"`
	AddressPrefixes                   []string                          `pulumi:"addressPrefixes"`
	Delegations                       []Delegation                      `pulumi:"delegations"`
	Id                                *string                           `pulumi:"id"`
	IpAllocations                     []SubResource                     `pulumi:"ipAllocations"`
	Name                              *string                           `pulumi:"name"`
	NatGateway                        *SubResource                      `pulumi:"natGateway"`
	NetworkSecurityGroup              *NetworkSecurityGroupType         `pulumi:"networkSecurityGroup"`
	PrivateEndpointNetworkPolicies    *string                           `pulumi:"privateEndpointNetworkPolicies"`
	PrivateLinkServiceNetworkPolicies *string                           `pulumi:"privateLinkServiceNetworkPolicies"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	RouteTable                        *RouteTableType                   `pulumi:"routeTable"`
	ServiceEndpointPolicies           []ServiceEndpointPolicyType       `pulumi:"serviceEndpointPolicies"`
	ServiceEndpoints                  []ServiceEndpointPropertiesFormat `pulumi:"serviceEndpoints"`
	SubnetName                        *string                           `pulumi:"subnetName"`
	Type                              *string                           `pulumi:"type"`
	VirtualNetworkName                string                            `pulumi:"virtualNetworkName"`
}


type SubnetArgs struct {
	AddressPrefix                     pulumi.StringPtrInput
	AddressPrefixes                   pulumi.StringArrayInput
	Delegations                       DelegationArrayInput
	Id                                pulumi.StringPtrInput
	IpAllocations                     SubResourceArrayInput
	Name                              pulumi.StringPtrInput
	NatGateway                        SubResourcePtrInput
	NetworkSecurityGroup              NetworkSecurityGroupTypePtrInput
	PrivateEndpointNetworkPolicies    pulumi.StringPtrInput
	PrivateLinkServiceNetworkPolicies pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
	RouteTable                        RouteTableTypePtrInput
	ServiceEndpointPolicies           ServiceEndpointPolicyTypeArrayInput
	ServiceEndpoints                  ServiceEndpointPropertiesFormatArrayInput
	SubnetName                        pulumi.StringPtrInput
	Type                              pulumi.StringPtrInput
	VirtualNetworkName                pulumi.StringInput
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetArgs)(nil)).Elem()
}

type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(ctx context.Context) SubnetOutput
}

func (*Subnet) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (i *Subnet) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i *Subnet) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}

type SubnetOutput struct{ *pulumi.OutputState }

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subnet)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubnetOutput{})
}
