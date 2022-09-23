


package v20180401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Subnet struct {
	pulumi.CustomResourceState

	AddressPrefix           pulumi.StringPtrOutput                             `pulumi:"addressPrefix"`
	Etag                    pulumi.StringPtrOutput                             `pulumi:"etag"`
	IpConfigurations        IPConfigurationResponseArrayOutput                 `pulumi:"ipConfigurations"`
	Name                    pulumi.StringPtrOutput                             `pulumi:"name"`
	NetworkSecurityGroup    NetworkSecurityGroupResponsePtrOutput              `pulumi:"networkSecurityGroup"`
	ProvisioningState       pulumi.StringPtrOutput                             `pulumi:"provisioningState"`
	ResourceNavigationLinks ResourceNavigationLinkResponseArrayOutput          `pulumi:"resourceNavigationLinks"`
	RouteTable              RouteTableResponsePtrOutput                        `pulumi:"routeTable"`
	ServiceEndpoints        ServiceEndpointPropertiesFormatResponseArrayOutput `pulumi:"serviceEndpoints"`
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
			Type: pulumi.String("azure-native:network/v20200701:Subnet"),
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
		{
			Type: pulumi.String("azure-native:network/v20210801:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:Subnet"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:Subnet"),
		},
	})
	opts = append(opts, aliases)
	var resource Subnet
	err := ctx.RegisterResource("azure-native:network/v20180401:Subnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetState, opts ...pulumi.ResourceOption) (*Subnet, error) {
	var resource Subnet
	err := ctx.ReadResource("azure-native:network/v20180401:Subnet", name, id, state, &resource, opts...)
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
	AddressPrefix           *string                           `pulumi:"addressPrefix"`
	Id                      *string                           `pulumi:"id"`
	Name                    *string                           `pulumi:"name"`
	NetworkSecurityGroup    *NetworkSecurityGroupType         `pulumi:"networkSecurityGroup"`
	ProvisioningState       *string                           `pulumi:"provisioningState"`
	ResourceGroupName       string                            `pulumi:"resourceGroupName"`
	ResourceNavigationLinks []ResourceNavigationLink          `pulumi:"resourceNavigationLinks"`
	RouteTable              *RouteTableType                   `pulumi:"routeTable"`
	ServiceEndpoints        []ServiceEndpointPropertiesFormat `pulumi:"serviceEndpoints"`
	SubnetName              *string                           `pulumi:"subnetName"`
	VirtualNetworkName      string                            `pulumi:"virtualNetworkName"`
}


type SubnetArgs struct {
	AddressPrefix           pulumi.StringPtrInput
	Id                      pulumi.StringPtrInput
	Name                    pulumi.StringPtrInput
	NetworkSecurityGroup    NetworkSecurityGroupTypePtrInput
	ProvisioningState       pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
	ResourceNavigationLinks ResourceNavigationLinkArrayInput
	RouteTable              RouteTableTypePtrInput
	ServiceEndpoints        ServiceEndpointPropertiesFormatArrayInput
	SubnetName              pulumi.StringPtrInput
	VirtualNetworkName      pulumi.StringInput
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

func (o SubnetOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringPtrOutput { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o SubnetOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o SubnetOutput) IpConfigurations() IPConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) IPConfigurationResponseArrayOutput { return v.IpConfigurations }).(IPConfigurationResponseArrayOutput)
}

func (o SubnetOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SubnetOutput) NetworkSecurityGroup() NetworkSecurityGroupResponsePtrOutput {
	return o.ApplyT(func(v *Subnet) NetworkSecurityGroupResponsePtrOutput { return v.NetworkSecurityGroup }).(NetworkSecurityGroupResponsePtrOutput)
}

func (o SubnetOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Subnet) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o SubnetOutput) ResourceNavigationLinks() ResourceNavigationLinkResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) ResourceNavigationLinkResponseArrayOutput { return v.ResourceNavigationLinks }).(ResourceNavigationLinkResponseArrayOutput)
}

func (o SubnetOutput) RouteTable() RouteTableResponsePtrOutput {
	return o.ApplyT(func(v *Subnet) RouteTableResponsePtrOutput { return v.RouteTable }).(RouteTableResponsePtrOutput)
}

func (o SubnetOutput) ServiceEndpoints() ServiceEndpointPropertiesFormatResponseArrayOutput {
	return o.ApplyT(func(v *Subnet) ServiceEndpointPropertiesFormatResponseArrayOutput { return v.ServiceEndpoints }).(ServiceEndpointPropertiesFormatResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(SubnetOutput{})
}
