


package v20200401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkTap struct {
	pulumi.CustomResourceState

	DestinationLoadBalancerFrontEndIPConfiguration FrontendIPConfigurationResponsePtrOutput            `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	DestinationNetworkInterfaceIPConfiguration     NetworkInterfaceIPConfigurationResponsePtrOutput    `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	DestinationPort                                pulumi.IntPtrOutput                                 `pulumi:"destinationPort"`
	Etag                                           pulumi.StringOutput                                 `pulumi:"etag"`
	Location                                       pulumi.StringPtrOutput                              `pulumi:"location"`
	Name                                           pulumi.StringOutput                                 `pulumi:"name"`
	NetworkInterfaceTapConfigurations              NetworkInterfaceTapConfigurationResponseArrayOutput `pulumi:"networkInterfaceTapConfigurations"`
	ProvisioningState                              pulumi.StringOutput                                 `pulumi:"provisioningState"`
	ResourceGuid                                   pulumi.StringOutput                                 `pulumi:"resourceGuid"`
	Tags                                           pulumi.StringMapOutput                              `pulumi:"tags"`
	Type                                           pulumi.StringOutput                                 `pulumi:"type"`
}


func NewVirtualNetworkTap(ctx *pulumi.Context,
	name string, args *VirtualNetworkTapArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkTap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:VirtualNetworkTap"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220701:VirtualNetworkTap"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkTap
	err := ctx.RegisterResource("azure-native:network/v20200401:VirtualNetworkTap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkTap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkTapState, opts ...pulumi.ResourceOption) (*VirtualNetworkTap, error) {
	var resource VirtualNetworkTap
	err := ctx.ReadResource("azure-native:network/v20200401:VirtualNetworkTap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkTapState struct {
}

type VirtualNetworkTapState struct {
}

func (VirtualNetworkTapState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkTapState)(nil)).Elem()
}

type virtualNetworkTapArgs struct {
	DestinationLoadBalancerFrontEndIPConfiguration *FrontendIPConfiguration         `pulumi:"destinationLoadBalancerFrontEndIPConfiguration"`
	DestinationNetworkInterfaceIPConfiguration     *NetworkInterfaceIPConfiguration `pulumi:"destinationNetworkInterfaceIPConfiguration"`
	DestinationPort                                *int                             `pulumi:"destinationPort"`
	Id                                             *string                          `pulumi:"id"`
	Location                                       *string                          `pulumi:"location"`
	ResourceGroupName                              string                           `pulumi:"resourceGroupName"`
	Tags                                           map[string]string                `pulumi:"tags"`
	TapName                                        *string                          `pulumi:"tapName"`
}


type VirtualNetworkTapArgs struct {
	DestinationLoadBalancerFrontEndIPConfiguration FrontendIPConfigurationPtrInput
	DestinationNetworkInterfaceIPConfiguration     NetworkInterfaceIPConfigurationPtrInput
	DestinationPort                                pulumi.IntPtrInput
	Id                                             pulumi.StringPtrInput
	Location                                       pulumi.StringPtrInput
	ResourceGroupName                              pulumi.StringInput
	Tags                                           pulumi.StringMapInput
	TapName                                        pulumi.StringPtrInput
}

func (VirtualNetworkTapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkTapArgs)(nil)).Elem()
}

type VirtualNetworkTapInput interface {
	pulumi.Input

	ToVirtualNetworkTapOutput() VirtualNetworkTapOutput
	ToVirtualNetworkTapOutputWithContext(ctx context.Context) VirtualNetworkTapOutput
}

func (*VirtualNetworkTap) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkTap)(nil)).Elem()
}

func (i *VirtualNetworkTap) ToVirtualNetworkTapOutput() VirtualNetworkTapOutput {
	return i.ToVirtualNetworkTapOutputWithContext(context.Background())
}

func (i *VirtualNetworkTap) ToVirtualNetworkTapOutputWithContext(ctx context.Context) VirtualNetworkTapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkTapOutput)
}

type VirtualNetworkTapOutput struct{ *pulumi.OutputState }

func (VirtualNetworkTapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkTap)(nil)).Elem()
}

func (o VirtualNetworkTapOutput) ToVirtualNetworkTapOutput() VirtualNetworkTapOutput {
	return o
}

func (o VirtualNetworkTapOutput) ToVirtualNetworkTapOutputWithContext(ctx context.Context) VirtualNetworkTapOutput {
	return o
}

func (o VirtualNetworkTapOutput) DestinationLoadBalancerFrontEndIPConfiguration() FrontendIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) FrontendIPConfigurationResponsePtrOutput {
		return v.DestinationLoadBalancerFrontEndIPConfiguration
	}).(FrontendIPConfigurationResponsePtrOutput)
}

func (o VirtualNetworkTapOutput) DestinationNetworkInterfaceIPConfiguration() NetworkInterfaceIPConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) NetworkInterfaceIPConfigurationResponsePtrOutput {
		return v.DestinationNetworkInterfaceIPConfiguration
	}).(NetworkInterfaceIPConfigurationResponsePtrOutput)
}

func (o VirtualNetworkTapOutput) DestinationPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.IntPtrOutput { return v.DestinationPort }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkTapOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkTapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapOutput) NetworkInterfaceTapConfigurations() NetworkInterfaceTapConfigurationResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) NetworkInterfaceTapConfigurationResponseArrayOutput {
		return v.NetworkInterfaceTapConfigurations
	}).(NetworkInterfaceTapConfigurationResponseArrayOutput)
}

func (o VirtualNetworkTapOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o VirtualNetworkTapOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkTapOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkTap) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkTapOutput{})
}
