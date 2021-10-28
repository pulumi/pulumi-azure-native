


package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualHub struct {
	pulumi.CustomResourceState

	AddressPrefix             pulumi.StringPtrOutput                         `pulumi:"addressPrefix"`
	Etag                      pulumi.StringOutput                            `pulumi:"etag"`
	ExpressRouteGateway       SubResourceResponsePtrOutput                   `pulumi:"expressRouteGateway"`
	Location                  pulumi.StringOutput                            `pulumi:"location"`
	Name                      pulumi.StringOutput                            `pulumi:"name"`
	P2SVpnGateway             SubResourceResponsePtrOutput                   `pulumi:"p2SVpnGateway"`
	ProvisioningState         pulumi.StringOutput                            `pulumi:"provisioningState"`
	RouteTable                VirtualHubRouteTableResponsePtrOutput          `pulumi:"routeTable"`
	Tags                      pulumi.StringMapOutput                         `pulumi:"tags"`
	Type                      pulumi.StringOutput                            `pulumi:"type"`
	VirtualNetworkConnections HubVirtualNetworkConnectionResponseArrayOutput `pulumi:"virtualNetworkConnections"`
	VirtualWan                SubResourceResponsePtrOutput                   `pulumi:"virtualWan"`
	VpnGateway                SubResourceResponsePtrOutput                   `pulumi:"vpnGateway"`
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
			Type: pulumi.String("azure-nextgen:network/v20190401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:VirtualHub"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualHub
	err := ctx.RegisterResource("azure-native:network/v20190401:VirtualHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubState, opts ...pulumi.ResourceOption) (*VirtualHub, error) {
	var resource VirtualHub
	err := ctx.ReadResource("azure-native:network/v20190401:VirtualHub", name, id, state, &resource, opts...)
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
	AddressPrefix             *string                       `pulumi:"addressPrefix"`
	ExpressRouteGateway       *SubResource                  `pulumi:"expressRouteGateway"`
	Id                        *string                       `pulumi:"id"`
	Location                  *string                       `pulumi:"location"`
	P2SVpnGateway             *SubResource                  `pulumi:"p2SVpnGateway"`
	ResourceGroupName         string                        `pulumi:"resourceGroupName"`
	RouteTable                *VirtualHubRouteTable         `pulumi:"routeTable"`
	Tags                      map[string]string             `pulumi:"tags"`
	VirtualHubName            *string                       `pulumi:"virtualHubName"`
	VirtualNetworkConnections []HubVirtualNetworkConnection `pulumi:"virtualNetworkConnections"`
	VirtualWan                *SubResource                  `pulumi:"virtualWan"`
	VpnGateway                *SubResource                  `pulumi:"vpnGateway"`
}


type VirtualHubArgs struct {
	AddressPrefix             pulumi.StringPtrInput
	ExpressRouteGateway       SubResourcePtrInput
	Id                        pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	P2SVpnGateway             SubResourcePtrInput
	ResourceGroupName         pulumi.StringInput
	RouteTable                VirtualHubRouteTablePtrInput
	Tags                      pulumi.StringMapInput
	VirtualHubName            pulumi.StringPtrInput
	VirtualNetworkConnections HubVirtualNetworkConnectionArrayInput
	VirtualWan                SubResourcePtrInput
	VpnGateway                SubResourcePtrInput
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
	return reflect.TypeOf((*VirtualHub)(nil))
}

func (i *VirtualHub) ToVirtualHubOutput() VirtualHubOutput {
	return i.ToVirtualHubOutputWithContext(context.Background())
}

func (i *VirtualHub) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubOutput)
}

type VirtualHubOutput struct{ *pulumi.OutputState }

func (VirtualHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHub)(nil))
}

func (o VirtualHubOutput) ToVirtualHubOutput() VirtualHubOutput {
	return o
}

func (o VirtualHubOutput) ToVirtualHubOutputWithContext(ctx context.Context) VirtualHubOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualHubInput)(nil)).Elem(), &VirtualHub{})
	pulumi.RegisterOutputType(VirtualHubOutput{})
}
