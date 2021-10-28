


package v20200401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VpnGateway struct {
	pulumi.CustomResourceState

	BgpSettings         BgpSettingsResponsePtrOutput     `pulumi:"bgpSettings"`
	Connections         VpnConnectionResponseArrayOutput `pulumi:"connections"`
	Etag                pulumi.StringOutput              `pulumi:"etag"`
	Location            pulumi.StringOutput              `pulumi:"location"`
	Name                pulumi.StringOutput              `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput              `pulumi:"provisioningState"`
	Tags                pulumi.StringMapOutput           `pulumi:"tags"`
	Type                pulumi.StringOutput              `pulumi:"type"`
	VirtualHub          SubResourceResponsePtrOutput     `pulumi:"virtualHub"`
	VpnGatewayScaleUnit pulumi.IntPtrOutput              `pulumi:"vpnGatewayScaleUnit"`
}


func NewVpnGateway(ctx *pulumi.Context,
	name string, args *VpnGatewayArgs, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VpnGateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210501:VpnGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource VpnGateway
	err := ctx.RegisterResource("azure-native:network/v20200401:VpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnGatewayState, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	var resource VpnGateway
	err := ctx.ReadResource("azure-native:network/v20200401:VpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vpnGatewayState struct {
}

type VpnGatewayState struct {
}

func (VpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayState)(nil)).Elem()
}

type vpnGatewayArgs struct {
	BgpSettings         *BgpSettings        `pulumi:"bgpSettings"`
	Connections         []VpnConnectionType `pulumi:"connections"`
	GatewayName         *string             `pulumi:"gatewayName"`
	Id                  *string             `pulumi:"id"`
	Location            *string             `pulumi:"location"`
	ResourceGroupName   string              `pulumi:"resourceGroupName"`
	Tags                map[string]string   `pulumi:"tags"`
	VirtualHub          *SubResource        `pulumi:"virtualHub"`
	VpnGatewayScaleUnit *int                `pulumi:"vpnGatewayScaleUnit"`
}


type VpnGatewayArgs struct {
	BgpSettings         BgpSettingsPtrInput
	Connections         VpnConnectionTypeArrayInput
	GatewayName         pulumi.StringPtrInput
	Id                  pulumi.StringPtrInput
	Location            pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
	VirtualHub          SubResourcePtrInput
	VpnGatewayScaleUnit pulumi.IntPtrInput
}

func (VpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayArgs)(nil)).Elem()
}

type VpnGatewayInput interface {
	pulumi.Input

	ToVpnGatewayOutput() VpnGatewayOutput
	ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput
}

func (*VpnGateway) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGateway)(nil))
}

func (i *VpnGateway) ToVpnGatewayOutput() VpnGatewayOutput {
	return i.ToVpnGatewayOutputWithContext(context.Background())
}

func (i *VpnGateway) ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnGatewayOutput)
}

type VpnGatewayOutput struct{ *pulumi.OutputState }

func (VpnGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpnGateway)(nil))
}

func (o VpnGatewayOutput) ToVpnGatewayOutput() VpnGatewayOutput {
	return o
}

func (o VpnGatewayOutput) ToVpnGatewayOutputWithContext(ctx context.Context) VpnGatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VpnGatewayOutput{})
}
