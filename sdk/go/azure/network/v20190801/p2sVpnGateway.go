


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type P2sVpnGateway struct {
	pulumi.CustomResourceState

	Etag                        pulumi.StringOutput                           `pulumi:"etag"`
	Location                    pulumi.StringOutput                           `pulumi:"location"`
	Name                        pulumi.StringOutput                           `pulumi:"name"`
	P2SConnectionConfigurations P2SConnectionConfigurationResponseArrayOutput `pulumi:"p2SConnectionConfigurations"`
	ProvisioningState           pulumi.StringOutput                           `pulumi:"provisioningState"`
	Tags                        pulumi.StringMapOutput                        `pulumi:"tags"`
	Type                        pulumi.StringOutput                           `pulumi:"type"`
	VirtualHub                  SubResourceResponsePtrOutput                  `pulumi:"virtualHub"`
	VpnClientConnectionHealth   VpnClientConnectionHealthResponseOutput       `pulumi:"vpnClientConnectionHealth"`
	VpnGatewayScaleUnit         pulumi.IntPtrOutput                           `pulumi:"vpnGatewayScaleUnit"`
	VpnServerConfiguration      SubResourceResponsePtrOutput                  `pulumi:"vpnServerConfiguration"`
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
			Type: pulumi.String("azure-nextgen:network/v20190801:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:P2sVpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:P2sVpnGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource P2sVpnGateway
	err := ctx.RegisterResource("azure-native:network/v20190801:P2sVpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetP2sVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *P2sVpnGatewayState, opts ...pulumi.ResourceOption) (*P2sVpnGateway, error) {
	var resource P2sVpnGateway
	err := ctx.ReadResource("azure-native:network/v20190801:P2sVpnGateway", name, id, state, &resource, opts...)
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
	GatewayName                 *string                      `pulumi:"gatewayName"`
	Id                          *string                      `pulumi:"id"`
	Location                    *string                      `pulumi:"location"`
	P2SConnectionConfigurations []P2SConnectionConfiguration `pulumi:"p2SConnectionConfigurations"`
	ResourceGroupName           string                       `pulumi:"resourceGroupName"`
	Tags                        map[string]string            `pulumi:"tags"`
	VirtualHub                  *SubResource                 `pulumi:"virtualHub"`
	VpnGatewayScaleUnit         *int                         `pulumi:"vpnGatewayScaleUnit"`
	VpnServerConfiguration      *SubResource                 `pulumi:"vpnServerConfiguration"`
}


type P2sVpnGatewayArgs struct {
	GatewayName                 pulumi.StringPtrInput
	Id                          pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	P2SConnectionConfigurations P2SConnectionConfigurationArrayInput
	ResourceGroupName           pulumi.StringInput
	Tags                        pulumi.StringMapInput
	VirtualHub                  SubResourcePtrInput
	VpnGatewayScaleUnit         pulumi.IntPtrInput
	VpnServerConfiguration      SubResourcePtrInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*P2sVpnGatewayInput)(nil)).Elem(), &P2sVpnGateway{})
	pulumi.RegisterOutputType(P2sVpnGatewayOutput{})
}
