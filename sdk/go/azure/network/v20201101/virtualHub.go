


package v20201101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualHub struct {
	pulumi.CustomResourceState

	AddressPrefix              pulumi.StringPtrOutput                    `pulumi:"addressPrefix"`
	AllowBranchToBranchTraffic pulumi.BoolPtrOutput                      `pulumi:"allowBranchToBranchTraffic"`
	AzureFirewall              SubResourceResponsePtrOutput              `pulumi:"azureFirewall"`
	BgpConnections             SubResourceResponseArrayOutput            `pulumi:"bgpConnections"`
	Etag                       pulumi.StringOutput                       `pulumi:"etag"`
	ExpressRouteGateway        SubResourceResponsePtrOutput              `pulumi:"expressRouteGateway"`
	IpConfigurations           SubResourceResponseArrayOutput            `pulumi:"ipConfigurations"`
	Location                   pulumi.StringOutput                       `pulumi:"location"`
	Name                       pulumi.StringOutput                       `pulumi:"name"`
	P2SVpnGateway              SubResourceResponsePtrOutput              `pulumi:"p2SVpnGateway"`
	ProvisioningState          pulumi.StringOutput                       `pulumi:"provisioningState"`
	RouteTable                 VirtualHubRouteTableResponsePtrOutput     `pulumi:"routeTable"`
	RoutingState               pulumi.StringOutput                       `pulumi:"routingState"`
	SecurityPartnerProvider    SubResourceResponsePtrOutput              `pulumi:"securityPartnerProvider"`
	SecurityProviderName       pulumi.StringPtrOutput                    `pulumi:"securityProviderName"`
	Sku                        pulumi.StringPtrOutput                    `pulumi:"sku"`
	Tags                       pulumi.StringMapOutput                    `pulumi:"tags"`
	Type                       pulumi.StringOutput                       `pulumi:"type"`
	VirtualHubRouteTableV2s    VirtualHubRouteTableV2ResponseArrayOutput `pulumi:"virtualHubRouteTableV2s"`
	VirtualRouterAsn           pulumi.Float64PtrOutput                   `pulumi:"virtualRouterAsn"`
	VirtualRouterIps           pulumi.StringArrayOutput                  `pulumi:"virtualRouterIps"`
	VirtualWan                 SubResourceResponsePtrOutput              `pulumi:"virtualWan"`
	VpnGateway                 SubResourceResponsePtrOutput              `pulumi:"vpnGateway"`
}


func NewVirtualHub(ctx *pulumi.Context,
	name string, args *VirtualHubArgs, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualHub"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualHub
	err := ctx.RegisterResource("azure-native:network/v20201101:VirtualHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubState, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	var resource VirtualHub
	err := ctx.ReadResource("azure-native:network/v20201101:VirtualHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualHubState struct {
}

type VirtualHubState struct {
}

func (VirtualHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubState)(nil)).Elem()
}

type virtualHubArgs struct {
	AddressPrefix              *string                      `pulumi:"addressPrefix"`
	AllowBranchToBranchTraffic *bool                        `pulumi:"allowBranchToBranchTraffic"`
	AzureFirewall              *SubResource                 `pulumi:"azureFirewall"`
	ExpressRouteGateway        *SubResource                 `pulumi:"expressRouteGateway"`
	Id                         *string                      `pulumi:"id"`
	Location                   *string                      `pulumi:"location"`
	P2SVpnGateway              *SubResource                 `pulumi:"p2SVpnGateway"`
	ResourceGroupName          string                       `pulumi:"resourceGroupName"`
	RouteTable                 *VirtualHubRouteTable        `pulumi:"routeTable"`
	SecurityPartnerProvider    *SubResource                 `pulumi:"securityPartnerProvider"`
	SecurityProviderName       *string                      `pulumi:"securityProviderName"`
	Sku                        *string                      `pulumi:"sku"`
	Tags                       map[string]string            `pulumi:"tags"`
	VirtualHubName             *string                      `pulumi:"virtualHubName"`
	VirtualHubRouteTableV2s    []VirtualHubRouteTableV2Type `pulumi:"virtualHubRouteTableV2s"`
	VirtualRouterAsn           *float64                     `pulumi:"virtualRouterAsn"`
	VirtualRouterIps           []string                     `pulumi:"virtualRouterIps"`
	VirtualWan                 *SubResource                 `pulumi:"virtualWan"`
	VpnGateway                 *SubResource                 `pulumi:"vpnGateway"`
}


type VirtualHubArgs struct {
	AddressPrefix              pulumi.StringPtrInput
	AllowBranchToBranchTraffic pulumi.BoolPtrInput
	AzureFirewall              SubResourcePtrInput
	ExpressRouteGateway        SubResourcePtrInput
	Id                         pulumi.StringPtrInput
	Location                   pulumi.StringPtrInput
	P2SVpnGateway              SubResourcePtrInput
	ResourceGroupName          pulumi.StringInput
	RouteTable                 VirtualHubRouteTablePtrInput
	SecurityPartnerProvider    SubResourcePtrInput
	SecurityProviderName       pulumi.StringPtrInput
	Sku                        pulumi.StringPtrInput
	Tags                       pulumi.StringMapInput
	VirtualHubName             pulumi.StringPtrInput
	VirtualHubRouteTableV2s    VirtualHubRouteTableV2TypeArrayInput
	VirtualRouterAsn           pulumi.Float64PtrInput
	VirtualRouterIps           pulumi.StringArrayInput
	VirtualWan                 SubResourcePtrInput
	VpnGateway                 SubResourcePtrInput
}

func (VirtualHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubArgs)(nil)).Elem()
}

type VirtualHubInput interface {
	pulumi.Input

	ToVirtualHubOutput() VirtualHubOutput
	ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput
}

func (*VirtualHub) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHub)(nil)).Elem()
}

func (i *VirtualHub) ToVirtualHubOutput() VirtualHubOutput {
	return i.ToVirtualHubOutputWithContext(context.Background())
}

func (i *VirtualHub) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubOutput)
}

type VirtualHubOutput struct{ *pulumi.OutputState }

func (VirtualHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualHub)(nil)).Elem()
}

func (o VirtualHubOutput) ToVirtualHubOutput() VirtualHubOutput {
	return o
}

func (o VirtualHubOutput) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return o
}

func (o VirtualHubOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringPtrOutput { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o VirtualHubOutput) AllowBranchToBranchTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.BoolPtrOutput { return v.AllowBranchToBranchTraffic }).(pulumi.BoolPtrOutput)
}

func (o VirtualHubOutput) AzureFirewall() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.AzureFirewall }).(SubResourceResponsePtrOutput)
}

func (o VirtualHubOutput) BgpConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponseArrayOutput { return v.BgpConnections }).(SubResourceResponseArrayOutput)
}

func (o VirtualHubOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) ExpressRouteGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.ExpressRouteGateway }).(SubResourceResponsePtrOutput)
}

func (o VirtualHubOutput) IpConfigurations() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponseArrayOutput { return v.IpConfigurations }).(SubResourceResponseArrayOutput)
}

func (o VirtualHubOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) P2SVpnGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.P2SVpnGateway }).(SubResourceResponsePtrOutput)
}

func (o VirtualHubOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) RouteTable() VirtualHubRouteTableResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) VirtualHubRouteTableResponsePtrOutput { return v.RouteTable }).(VirtualHubRouteTableResponsePtrOutput)
}

func (o VirtualHubOutput) RoutingState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.RoutingState }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) SecurityPartnerProvider() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.SecurityPartnerProvider }).(SubResourceResponsePtrOutput)
}

func (o VirtualHubOutput) SecurityProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringPtrOutput { return v.SecurityProviderName }).(pulumi.StringPtrOutput)
}

func (o VirtualHubOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringPtrOutput { return v.Sku }).(pulumi.StringPtrOutput)
}

func (o VirtualHubOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualHubOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualHubOutput) VirtualHubRouteTableV2s() VirtualHubRouteTableV2ResponseArrayOutput {
	return o.ApplyT(func(v *VirtualHub) VirtualHubRouteTableV2ResponseArrayOutput { return v.VirtualHubRouteTableV2s }).(VirtualHubRouteTableV2ResponseArrayOutput)
}

func (o VirtualHubOutput) VirtualRouterAsn() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.Float64PtrOutput { return v.VirtualRouterAsn }).(pulumi.Float64PtrOutput)
}

func (o VirtualHubOutput) VirtualRouterIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualHub) pulumi.StringArrayOutput { return v.VirtualRouterIps }).(pulumi.StringArrayOutput)
}

func (o VirtualHubOutput) VirtualWan() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.VirtualWan }).(SubResourceResponsePtrOutput)
}

func (o VirtualHubOutput) VpnGateway() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualHub) SubResourceResponsePtrOutput { return v.VpnGateway }).(SubResourceResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualHubOutput{})
}
