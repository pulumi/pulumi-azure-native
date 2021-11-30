


package v20190201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type P2sVpnGateway struct {
	pulumi.CustomResourceState

	CustomRoutes              AddressSpaceResponsePtrOutput           `pulumi:"customRoutes"`
	Etag                      pulumi.StringOutput                     `pulumi:"etag"`
	Location                  pulumi.StringOutput                     `pulumi:"location"`
	Name                      pulumi.StringOutput                     `pulumi:"name"`
	P2SVpnServerConfiguration SubResourceResponsePtrOutput            `pulumi:"p2SVpnServerConfiguration"`
	ProvisioningState         pulumi.StringOutput                     `pulumi:"provisioningState"`
	Tags                      pulumi.StringMapOutput                  `pulumi:"tags"`
	Type                      pulumi.StringOutput                     `pulumi:"type"`
	VirtualHub                SubResourceResponsePtrOutput            `pulumi:"virtualHub"`
	VpnClientAddressPool      AddressSpaceResponsePtrOutput           `pulumi:"vpnClientAddressPool"`
	VpnClientConnectionHealth VpnClientConnectionHealthResponseOutput `pulumi:"vpnClientConnectionHealth"`
	VpnGatewayScaleUnit       pulumi.IntPtrOutput                     `pulumi:"vpnGatewayScaleUnit"`
}


func NewP2sVpnGateway(ctx *pulumi.Context,
	name string, args *P2sVpnGatewayArgs, opts ...pulumi.ResourceOption) (*P2sVpnGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:P2sVpnGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource P2sVpnGateway
	err := ctx.RegisterResource("azure-native:network/v20190201:P2sVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetP2sVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *P2sVpnGatewayState, opts ...pulumi.ResourceOption) (*P2sVpnGateway, error) {
	var resource P2sVpnGateway
	err := ctx.ReadResource("azure-native:network/v20190201:P2sVpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type p2sVpnGatewayState struct {
}

type P2sVpnGatewayState struct {
}

func (P2sVpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*p2sVpnGatewayState)(nil)).Elem()
}

type p2sVpnGatewayArgs struct {
	CustomRoutes              *AddressSpace     `pulumi:"customRoutes"`
	GatewayName               *string           `pulumi:"gatewayName"`
	Id                        *string           `pulumi:"id"`
	Location                  *string           `pulumi:"location"`
	P2SVpnServerConfiguration *SubResource      `pulumi:"p2SVpnServerConfiguration"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	Tags                      map[string]string `pulumi:"tags"`
	VirtualHub                *SubResource      `pulumi:"virtualHub"`
	VpnClientAddressPool      *AddressSpace     `pulumi:"vpnClientAddressPool"`
	VpnGatewayScaleUnit       *int              `pulumi:"vpnGatewayScaleUnit"`
}


type P2sVpnGatewayArgs struct {
	CustomRoutes              AddressSpacePtrInput
	GatewayName               pulumi.StringPtrInput
	Id                        pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	P2SVpnServerConfiguration SubResourcePtrInput
	ResourceGroupName         pulumi.StringInput
	Tags                      pulumi.StringMapInput
	VirtualHub                SubResourcePtrInput
	VpnClientAddressPool      AddressSpacePtrInput
	VpnGatewayScaleUnit       pulumi.IntPtrInput
}

func (P2sVpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*p2sVpnGatewayArgs)(nil)).Elem()
}

type P2sVpnGatewayInput interface {
	pulumi.Input

	ToP2sVpnGatewayOutput() P2sVpnGatewayOutput
	ToP2sVpnGatewayOutputWithContext(ctx context.Context) P2sVpnGatewayOutput
}

func (*P2sVpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*P2sVpnGateway)(nil))
}

func (i *P2sVpnGateway) ToP2sVpnGatewayOutput() P2sVpnGatewayOutput {
	return i.ToP2sVpnGatewayOutputWithContext(context.Background())
}

func (i *P2sVpnGateway) ToP2sVpnGatewayOutputWithContext(ctx context.Context) P2sVpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(P2sVpnGatewayOutput)
}

type P2sVpnGatewayOutput struct{ *pulumi.OutputState }

func (P2sVpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*P2sVpnGateway)(nil))
}

func (o P2sVpnGatewayOutput) ToP2sVpnGatewayOutput() P2sVpnGatewayOutput {
	return o
}

func (o P2sVpnGatewayOutput) ToP2sVpnGatewayOutputWithContext(ctx context.Context) P2sVpnGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(P2sVpnGatewayOutput{})
}
