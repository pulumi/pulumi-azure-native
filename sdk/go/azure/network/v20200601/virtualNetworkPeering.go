


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkPeering struct {
	pulumi.CustomResourceState

	AllowForwardedTraffic     pulumi.BoolPtrOutput                          `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       pulumi.BoolPtrOutput                          `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess pulumi.BoolPtrOutput                          `pulumi:"allowVirtualNetworkAccess"`
	Etag                      pulumi.StringOutput                           `pulumi:"etag"`
	Name                      pulumi.StringPtrOutput                        `pulumi:"name"`
	PeeringState              pulumi.StringPtrOutput                        `pulumi:"peeringState"`
	ProvisioningState         pulumi.StringOutput                           `pulumi:"provisioningState"`
	RemoteAddressSpace        AddressSpaceResponsePtrOutput                 `pulumi:"remoteAddressSpace"`
	RemoteBgpCommunities      VirtualNetworkBgpCommunitiesResponsePtrOutput `pulumi:"remoteBgpCommunities"`
	RemoteVirtualNetwork      SubResourceResponsePtrOutput                  `pulumi:"remoteVirtualNetwork"`
	UseRemoteGateways         pulumi.BoolPtrOutput                          `pulumi:"useRemoteGateways"`
}


func NewVirtualNetworkPeering(ctx *pulumi.Context,
	name string, args *VirtualNetworkPeeringArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkPeering, error) {
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
			Type: pulumi.String("azure-native:network:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualNetworkPeering"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkPeering
	err := ctx.RegisterResource("azure-native:network/v20200601:VirtualNetworkPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkPeeringState, opts ...pulumi.ResourceOption) (*VirtualNetworkPeering, error) {
	var resource VirtualNetworkPeering
	err := ctx.ReadResource("azure-native:network/v20200601:VirtualNetworkPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkPeeringState struct {
}

type VirtualNetworkPeeringState struct {
}

func (VirtualNetworkPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkPeeringState)(nil)).Elem()
}

type virtualNetworkPeeringArgs struct {
	AllowForwardedTraffic     *bool                         `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool                         `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool                         `pulumi:"allowVirtualNetworkAccess"`
	Id                        *string                       `pulumi:"id"`
	Name                      *string                       `pulumi:"name"`
	PeeringState              *string                       `pulumi:"peeringState"`
	RemoteAddressSpace        *AddressSpace                 `pulumi:"remoteAddressSpace"`
	RemoteBgpCommunities      *VirtualNetworkBgpCommunities `pulumi:"remoteBgpCommunities"`
	RemoteVirtualNetwork      *SubResource                  `pulumi:"remoteVirtualNetwork"`
	ResourceGroupName         string                        `pulumi:"resourceGroupName"`
	UseRemoteGateways         *bool                         `pulumi:"useRemoteGateways"`
	VirtualNetworkName        string                        `pulumi:"virtualNetworkName"`
	VirtualNetworkPeeringName *string                       `pulumi:"virtualNetworkPeeringName"`
}


type VirtualNetworkPeeringArgs struct {
	AllowForwardedTraffic     pulumi.BoolPtrInput
	AllowGatewayTransit       pulumi.BoolPtrInput
	AllowVirtualNetworkAccess pulumi.BoolPtrInput
	Id                        pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	PeeringState              pulumi.StringPtrInput
	RemoteAddressSpace        AddressSpacePtrInput
	RemoteBgpCommunities      VirtualNetworkBgpCommunitiesPtrInput
	RemoteVirtualNetwork      SubResourcePtrInput
	ResourceGroupName         pulumi.StringInput
	UseRemoteGateways         pulumi.BoolPtrInput
	VirtualNetworkName        pulumi.StringInput
	VirtualNetworkPeeringName pulumi.StringPtrInput
}

func (VirtualNetworkPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkPeeringArgs)(nil)).Elem()
}

type VirtualNetworkPeeringInput interface {
	pulumi.Input

	ToVirtualNetworkPeeringOutput() VirtualNetworkPeeringOutput
	ToVirtualNetworkPeeringOutputWithContext(ctx context.Context) VirtualNetworkPeeringOutput
}

func (*VirtualNetworkPeering) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeering)(nil)).Elem()
}

func (i *VirtualNetworkPeering) ToVirtualNetworkPeeringOutput() VirtualNetworkPeeringOutput {
	return i.ToVirtualNetworkPeeringOutputWithContext(context.Background())
}

func (i *VirtualNetworkPeering) ToVirtualNetworkPeeringOutputWithContext(ctx context.Context) VirtualNetworkPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringOutput)
}

type VirtualNetworkPeeringOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkPeering)(nil)).Elem()
}

func (o VirtualNetworkPeeringOutput) ToVirtualNetworkPeeringOutput() VirtualNetworkPeeringOutput {
	return o
}

func (o VirtualNetworkPeeringOutput) ToVirtualNetworkPeeringOutputWithContext(ctx context.Context) VirtualNetworkPeeringOutput {
	return o
}

func (o VirtualNetworkPeeringOutput) AllowForwardedTraffic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) pulumi.BoolPtrOutput { return v.AllowForwardedTraffic }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringOutput) AllowGatewayTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) pulumi.BoolPtrOutput { return v.AllowGatewayTransit }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringOutput) AllowVirtualNetworkAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) pulumi.BoolPtrOutput { return v.AllowVirtualNetworkAccess }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkPeeringOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualNetworkPeeringOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringOutput) PeeringState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) pulumi.StringPtrOutput { return v.PeeringState }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkPeeringOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkPeeringOutput) RemoteAddressSpace() AddressSpaceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) AddressSpaceResponsePtrOutput { return v.RemoteAddressSpace }).(AddressSpaceResponsePtrOutput)
}

func (o VirtualNetworkPeeringOutput) RemoteBgpCommunities() VirtualNetworkBgpCommunitiesResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) VirtualNetworkBgpCommunitiesResponsePtrOutput {
		return v.RemoteBgpCommunities
	}).(VirtualNetworkBgpCommunitiesResponsePtrOutput)
}

func (o VirtualNetworkPeeringOutput) RemoteVirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) SubResourceResponsePtrOutput { return v.RemoteVirtualNetwork }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkPeeringOutput) UseRemoteGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkPeering) pulumi.BoolPtrOutput { return v.UseRemoteGateways }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkPeeringOutput{})
}
