


package v20210501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkPeering struct {
	pulumi.CustomResourceState

	AllowForwardedTraffic            pulumi.BoolPtrOutput                          `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit              pulumi.BoolPtrOutput                          `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess        pulumi.BoolPtrOutput                          `pulumi:"allowVirtualNetworkAccess"`
	DoNotVerifyRemoteGateways        pulumi.BoolPtrOutput                          `pulumi:"doNotVerifyRemoteGateways"`
	Etag                             pulumi.StringOutput                           `pulumi:"etag"`
	Name                             pulumi.StringPtrOutput                        `pulumi:"name"`
	PeeringState                     pulumi.StringPtrOutput                        `pulumi:"peeringState"`
	PeeringSyncLevel                 pulumi.StringPtrOutput                        `pulumi:"peeringSyncLevel"`
	ProvisioningState                pulumi.StringOutput                           `pulumi:"provisioningState"`
	RemoteAddressSpace               AddressSpaceResponsePtrOutput                 `pulumi:"remoteAddressSpace"`
	RemoteBgpCommunities             VirtualNetworkBgpCommunitiesResponsePtrOutput `pulumi:"remoteBgpCommunities"`
	RemoteVirtualNetwork             SubResourceResponsePtrOutput                  `pulumi:"remoteVirtualNetwork"`
	RemoteVirtualNetworkAddressSpace AddressSpaceResponsePtrOutput                 `pulumi:"remoteVirtualNetworkAddressSpace"`
	RemoteVirtualNetworkEncryption   VirtualNetworkEncryptionResponseOutput        `pulumi:"remoteVirtualNetworkEncryption"`
	ResourceGuid                     pulumi.StringOutput                           `pulumi:"resourceGuid"`
	Type                             pulumi.StringPtrOutput                        `pulumi:"type"`
	UseRemoteGateways                pulumi.BoolPtrOutput                          `pulumi:"useRemoteGateways"`
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
			Type: pulumi.String("azure-nextgen:network/v20210501:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160901:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20161201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170301:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetworkPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:VirtualNetworkPeering"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkPeering
	err := ctx.RegisterResource("azure-native:network/v20210501:VirtualNetworkPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkPeeringState, opts ...pulumi.ResourceOption) (*VirtualNetworkPeering, error) {
	var resource VirtualNetworkPeering
	err := ctx.ReadResource("azure-native:network/v20210501:VirtualNetworkPeering", name, id, state, &resource, opts...)
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
	AllowForwardedTraffic            *bool                         `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit              *bool                         `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess        *bool                         `pulumi:"allowVirtualNetworkAccess"`
	DoNotVerifyRemoteGateways        *bool                         `pulumi:"doNotVerifyRemoteGateways"`
	Id                               *string                       `pulumi:"id"`
	Name                             *string                       `pulumi:"name"`
	PeeringState                     *string                       `pulumi:"peeringState"`
	PeeringSyncLevel                 *string                       `pulumi:"peeringSyncLevel"`
	RemoteAddressSpace               *AddressSpace                 `pulumi:"remoteAddressSpace"`
	RemoteBgpCommunities             *VirtualNetworkBgpCommunities `pulumi:"remoteBgpCommunities"`
	RemoteVirtualNetwork             *SubResource                  `pulumi:"remoteVirtualNetwork"`
	RemoteVirtualNetworkAddressSpace *AddressSpace                 `pulumi:"remoteVirtualNetworkAddressSpace"`
	ResourceGroupName                string                        `pulumi:"resourceGroupName"`
	SyncRemoteAddressSpace           *string                       `pulumi:"syncRemoteAddressSpace"`
	Type                             *string                       `pulumi:"type"`
	UseRemoteGateways                *bool                         `pulumi:"useRemoteGateways"`
	VirtualNetworkName               string                        `pulumi:"virtualNetworkName"`
	VirtualNetworkPeeringName        *string                       `pulumi:"virtualNetworkPeeringName"`
}


type VirtualNetworkPeeringArgs struct {
	AllowForwardedTraffic            pulumi.BoolPtrInput
	AllowGatewayTransit              pulumi.BoolPtrInput
	AllowVirtualNetworkAccess        pulumi.BoolPtrInput
	DoNotVerifyRemoteGateways        pulumi.BoolPtrInput
	Id                               pulumi.StringPtrInput
	Name                             pulumi.StringPtrInput
	PeeringState                     pulumi.StringPtrInput
	PeeringSyncLevel                 pulumi.StringPtrInput
	RemoteAddressSpace               AddressSpacePtrInput
	RemoteBgpCommunities             VirtualNetworkBgpCommunitiesPtrInput
	RemoteVirtualNetwork             SubResourcePtrInput
	RemoteVirtualNetworkAddressSpace AddressSpacePtrInput
	ResourceGroupName                pulumi.StringInput
	SyncRemoteAddressSpace           pulumi.StringPtrInput
	Type                             pulumi.StringPtrInput
	UseRemoteGateways                pulumi.BoolPtrInput
	VirtualNetworkName               pulumi.StringInput
	VirtualNetworkPeeringName        pulumi.StringPtrInput
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
	return reflect.TypeOf((*VirtualNetworkPeering)(nil))
}

func (i *VirtualNetworkPeering) ToVirtualNetworkPeeringOutput() VirtualNetworkPeeringOutput {
	return i.ToVirtualNetworkPeeringOutputWithContext(context.Background())
}

func (i *VirtualNetworkPeering) ToVirtualNetworkPeeringOutputWithContext(ctx context.Context) VirtualNetworkPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkPeeringOutput)
}

type VirtualNetworkPeeringOutput struct{ *pulumi.OutputState }

func (VirtualNetworkPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkPeering)(nil))
}

func (o VirtualNetworkPeeringOutput) ToVirtualNetworkPeeringOutput() VirtualNetworkPeeringOutput {
	return o
}

func (o VirtualNetworkPeeringOutput) ToVirtualNetworkPeeringOutputWithContext(ctx context.Context) VirtualNetworkPeeringOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualNetworkPeeringInput)(nil)).Elem(), &VirtualNetworkPeering{})
	pulumi.RegisterOutputType(VirtualNetworkPeeringOutput{})
}
