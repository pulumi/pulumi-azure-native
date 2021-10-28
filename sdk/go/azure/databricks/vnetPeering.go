


package databricks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VNetPeering struct {
	pulumi.CustomResourceState

	AllowForwardedTraffic     pulumi.BoolPtrOutput                                                           `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       pulumi.BoolPtrOutput                                                           `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess pulumi.BoolPtrOutput                                                           `pulumi:"allowVirtualNetworkAccess"`
	DatabricksAddressSpace    AddressSpaceResponsePtrOutput                                                  `pulumi:"databricksAddressSpace"`
	DatabricksVirtualNetwork  VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetworkPtrOutput `pulumi:"databricksVirtualNetwork"`
	Name                      pulumi.StringOutput                                                            `pulumi:"name"`
	PeeringState              pulumi.StringOutput                                                            `pulumi:"peeringState"`
	ProvisioningState         pulumi.StringOutput                                                            `pulumi:"provisioningState"`
	RemoteAddressSpace        AddressSpaceResponsePtrOutput                                                  `pulumi:"remoteAddressSpace"`
	RemoteVirtualNetwork      VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetworkOutput        `pulumi:"remoteVirtualNetwork"`
	Type                      pulumi.StringOutput                                                            `pulumi:"type"`
	UseRemoteGateways         pulumi.BoolPtrOutput                                                           `pulumi:"useRemoteGateways"`
}


func NewVNetPeering(ctx *pulumi.Context,
	name string, args *VNetPeeringArgs, opts ...pulumi.ResourceOption) (*VNetPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RemoteVirtualNetwork == nil {
		return nil, errors.New("invalid value for required argument 'RemoteVirtualNetwork'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:databricks:vNetPeering"),
		},
		{
			Type: pulumi.String("azure-native:databricks/v20180401:vNetPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:databricks/v20180401:vNetPeering"),
		},
		{
			Type: pulumi.String("azure-native:databricks/v20210401preview:vNetPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:databricks/v20210401preview:vNetPeering"),
		},
	})
	opts = append(opts, aliases)
	var resource VNetPeering
	err := ctx.RegisterResource("azure-native:databricks:vNetPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVNetPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VNetPeeringState, opts ...pulumi.ResourceOption) (*VNetPeering, error) {
	var resource VNetPeering
	err := ctx.ReadResource("azure-native:databricks:vNetPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type vnetPeeringState struct {
}

type VNetPeeringState struct {
}

func (VNetPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*vnetPeeringState)(nil)).Elem()
}

type vnetPeeringArgs struct {
	AllowForwardedTraffic     *bool                                                          `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool                                                          `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool                                                          `pulumi:"allowVirtualNetworkAccess"`
	DatabricksAddressSpace    *AddressSpace                                                  `pulumi:"databricksAddressSpace"`
	DatabricksVirtualNetwork  *VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetwork `pulumi:"databricksVirtualNetwork"`
	PeeringName               *string                                                        `pulumi:"peeringName"`
	RemoteAddressSpace        *AddressSpace                                                  `pulumi:"remoteAddressSpace"`
	RemoteVirtualNetwork      VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetwork      `pulumi:"remoteVirtualNetwork"`
	ResourceGroupName         string                                                         `pulumi:"resourceGroupName"`
	UseRemoteGateways         *bool                                                          `pulumi:"useRemoteGateways"`
	WorkspaceName             string                                                         `pulumi:"workspaceName"`
}


type VNetPeeringArgs struct {
	AllowForwardedTraffic     pulumi.BoolPtrInput
	AllowGatewayTransit       pulumi.BoolPtrInput
	AllowVirtualNetworkAccess pulumi.BoolPtrInput
	DatabricksAddressSpace    AddressSpacePtrInput
	DatabricksVirtualNetwork  VirtualNetworkPeeringPropertiesFormatDatabricksVirtualNetworkPtrInput
	PeeringName               pulumi.StringPtrInput
	RemoteAddressSpace        AddressSpacePtrInput
	RemoteVirtualNetwork      VirtualNetworkPeeringPropertiesFormatRemoteVirtualNetworkInput
	ResourceGroupName         pulumi.StringInput
	UseRemoteGateways         pulumi.BoolPtrInput
	WorkspaceName             pulumi.StringInput
}

func (VNetPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vnetPeeringArgs)(nil)).Elem()
}

type VNetPeeringInput interface {
	pulumi.Input

	ToVNetPeeringOutput() VNetPeeringOutput
	ToVNetPeeringOutputWithContext(ctx context.Context) VNetPeeringOutput
}

func (*VNetPeering) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetPeering)(nil))
}

func (i *VNetPeering) ToVNetPeeringOutput() VNetPeeringOutput {
	return i.ToVNetPeeringOutputWithContext(context.Background())
}

func (i *VNetPeering) ToVNetPeeringOutputWithContext(ctx context.Context) VNetPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VNetPeeringOutput)
}

type VNetPeeringOutput struct{ *pulumi.OutputState }

func (VNetPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VNetPeering)(nil))
}

func (o VNetPeeringOutput) ToVNetPeeringOutput() VNetPeeringOutput {
	return o
}

func (o VNetPeeringOutput) ToVNetPeeringOutputWithContext(ctx context.Context) VNetPeeringOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VNetPeeringOutput{})
}
