


package v20200301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NatGateway struct {
	pulumi.CustomResourceState

	Etag                 pulumi.StringOutput            `pulumi:"etag"`
	IdleTimeoutInMinutes pulumi.IntPtrOutput            `pulumi:"idleTimeoutInMinutes"`
	Location             pulumi.StringPtrOutput         `pulumi:"location"`
	Name                 pulumi.StringOutput            `pulumi:"name"`
	ProvisioningState    pulumi.StringOutput            `pulumi:"provisioningState"`
	PublicIpAddresses    SubResourceResponseArrayOutput `pulumi:"publicIpAddresses"`
	PublicIpPrefixes     SubResourceResponseArrayOutput `pulumi:"publicIpPrefixes"`
	ResourceGuid         pulumi.StringOutput            `pulumi:"resourceGuid"`
	Sku                  NatGatewaySkuResponsePtrOutput `pulumi:"sku"`
	Subnets              SubResourceResponseArrayOutput `pulumi:"subnets"`
	Tags                 pulumi.StringMapOutput         `pulumi:"tags"`
	Type                 pulumi.StringOutput            `pulumi:"type"`
	Zones                pulumi.StringArrayOutput       `pulumi:"zones"`
}


func NewNatGateway(ctx *pulumi.Context,
	name string, args *NatGatewayArgs, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:NatGateway"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:NatGateway"),
		},
	})
	opts = append(opts, aliases)
	var resource NatGateway
	err := ctx.RegisterResource("azure-native:network/v20200301:NatGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetNatGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayState, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	var resource NatGateway
	err := ctx.ReadResource("azure-native:network/v20200301:NatGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type natGatewayState struct {
}

type NatGatewayState struct {
}

func (NatGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayState)(nil)).Elem()
}

type natGatewayArgs struct {
	Id                   *string           `pulumi:"id"`
	IdleTimeoutInMinutes *int              `pulumi:"idleTimeoutInMinutes"`
	Location             *string           `pulumi:"location"`
	NatGatewayName       *string           `pulumi:"natGatewayName"`
	PublicIpAddresses    []SubResource     `pulumi:"publicIpAddresses"`
	PublicIpPrefixes     []SubResource     `pulumi:"publicIpPrefixes"`
	ResourceGroupName    string            `pulumi:"resourceGroupName"`
	Sku                  *NatGatewaySku    `pulumi:"sku"`
	Tags                 map[string]string `pulumi:"tags"`
	Zones                []string          `pulumi:"zones"`
}


type NatGatewayArgs struct {
	Id                   pulumi.StringPtrInput
	IdleTimeoutInMinutes pulumi.IntPtrInput
	Location             pulumi.StringPtrInput
	NatGatewayName       pulumi.StringPtrInput
	PublicIpAddresses    SubResourceArrayInput
	PublicIpPrefixes     SubResourceArrayInput
	ResourceGroupName    pulumi.StringInput
	Sku                  NatGatewaySkuPtrInput
	Tags                 pulumi.StringMapInput
	Zones                pulumi.StringArrayInput
}

func (NatGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayArgs)(nil)).Elem()
}

type NatGatewayInput interface {
	pulumi.Input

	ToNatGatewayOutput() NatGatewayOutput
	ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput
}

func (*NatGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil)).Elem()
}

func (i *NatGateway) ToNatGatewayOutput() NatGatewayOutput {
	return i.ToNatGatewayOutputWithContext(context.Background())
}

func (i *NatGateway) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NatGatewayOutput)
}

type NatGatewayOutput struct{ *pulumi.OutputState }

func (NatGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NatGateway)(nil)).Elem()
}

func (o NatGatewayOutput) ToNatGatewayOutput() NatGatewayOutput {
	return o
}

func (o NatGatewayOutput) ToNatGatewayOutputWithContext(ctx context.Context) NatGatewayOutput {
	return o
}

func (o NatGatewayOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NatGatewayOutput) IdleTimeoutInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.IntPtrOutput { return v.IdleTimeoutInMinutes }).(pulumi.IntPtrOutput)
}

func (o NatGatewayOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o NatGatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NatGatewayOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o NatGatewayOutput) PublicIpAddresses() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NatGateway) SubResourceResponseArrayOutput { return v.PublicIpAddresses }).(SubResourceResponseArrayOutput)
}

func (o NatGatewayOutput) PublicIpPrefixes() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NatGateway) SubResourceResponseArrayOutput { return v.PublicIpPrefixes }).(SubResourceResponseArrayOutput)
}

func (o NatGatewayOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o NatGatewayOutput) Sku() NatGatewaySkuResponsePtrOutput {
	return o.ApplyT(func(v *NatGateway) NatGatewaySkuResponsePtrOutput { return v.Sku }).(NatGatewaySkuResponsePtrOutput)
}

func (o NatGatewayOutput) Subnets() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *NatGateway) SubResourceResponseArrayOutput { return v.Subnets }).(SubResourceResponseArrayOutput)
}

func (o NatGatewayOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o NatGatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o NatGatewayOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NatGateway) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(NatGatewayOutput{})
}
