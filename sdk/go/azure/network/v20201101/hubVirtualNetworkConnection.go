


package v20201101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HubVirtualNetworkConnection struct {
	pulumi.CustomResourceState

	AllowHubToRemoteVnetTransit         pulumi.BoolPtrOutput                  `pulumi:"allowHubToRemoteVnetTransit"`
	AllowRemoteVnetToUseHubVnetGateways pulumi.BoolPtrOutput                  `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	EnableInternetSecurity              pulumi.BoolPtrOutput                  `pulumi:"enableInternetSecurity"`
	Etag                                pulumi.StringOutput                   `pulumi:"etag"`
	Name                                pulumi.StringPtrOutput                `pulumi:"name"`
	ProvisioningState                   pulumi.StringOutput                   `pulumi:"provisioningState"`
	RemoteVirtualNetwork                SubResourceResponsePtrOutput          `pulumi:"remoteVirtualNetwork"`
	RoutingConfiguration                RoutingConfigurationResponsePtrOutput `pulumi:"routingConfiguration"`
}


func NewHubVirtualNetworkConnection(ctx *pulumi.Context,
	name string, args *HubVirtualNetworkConnectionArgs, opts ...pulumi.ResourceOption) (*HubVirtualNetworkConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualHubName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualHubName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:HubVirtualNetworkConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:HubVirtualNetworkConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource HubVirtualNetworkConnection
	err := ctx.RegisterResource("azure-native:network/v20201101:HubVirtualNetworkConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHubVirtualNetworkConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HubVirtualNetworkConnectionState, opts ...pulumi.ResourceOption) (*HubVirtualNetworkConnection, error) {
	var resource HubVirtualNetworkConnection
	err := ctx.ReadResource("azure-native:network/v20201101:HubVirtualNetworkConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hubVirtualNetworkConnectionState struct {
}

type HubVirtualNetworkConnectionState struct {
}

func (HubVirtualNetworkConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*hubVirtualNetworkConnectionState)(nil)).Elem()
}

type hubVirtualNetworkConnectionArgs struct {
	AllowHubToRemoteVnetTransit         *bool                 `pulumi:"allowHubToRemoteVnetTransit"`
	AllowRemoteVnetToUseHubVnetGateways *bool                 `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	ConnectionName                      *string               `pulumi:"connectionName"`
	EnableInternetSecurity              *bool                 `pulumi:"enableInternetSecurity"`
	Id                                  *string               `pulumi:"id"`
	Name                                *string               `pulumi:"name"`
	RemoteVirtualNetwork                *SubResource          `pulumi:"remoteVirtualNetwork"`
	ResourceGroupName                   string                `pulumi:"resourceGroupName"`
	RoutingConfiguration                *RoutingConfiguration `pulumi:"routingConfiguration"`
	VirtualHubName                      string                `pulumi:"virtualHubName"`
}


type HubVirtualNetworkConnectionArgs struct {
	AllowHubToRemoteVnetTransit         pulumi.BoolPtrInput
	AllowRemoteVnetToUseHubVnetGateways pulumi.BoolPtrInput
	ConnectionName                      pulumi.StringPtrInput
	EnableInternetSecurity              pulumi.BoolPtrInput
	Id                                  pulumi.StringPtrInput
	Name                                pulumi.StringPtrInput
	RemoteVirtualNetwork                SubResourcePtrInput
	ResourceGroupName                   pulumi.StringInput
	RoutingConfiguration                RoutingConfigurationPtrInput
	VirtualHubName                      pulumi.StringInput
}

func (HubVirtualNetworkConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hubVirtualNetworkConnectionArgs)(nil)).Elem()
}

type HubVirtualNetworkConnectionInput interface {
	pulumi.Input

	ToHubVirtualNetworkConnectionOutput() HubVirtualNetworkConnectionOutput
	ToHubVirtualNetworkConnectionOutputWithContext(ctx context.Context) HubVirtualNetworkConnectionOutput
}

func (*HubVirtualNetworkConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**HubVirtualNetworkConnection)(nil)).Elem()
}

func (i *HubVirtualNetworkConnection) ToHubVirtualNetworkConnectionOutput() HubVirtualNetworkConnectionOutput {
	return i.ToHubVirtualNetworkConnectionOutputWithContext(context.Background())
}

func (i *HubVirtualNetworkConnection) ToHubVirtualNetworkConnectionOutputWithContext(ctx context.Context) HubVirtualNetworkConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HubVirtualNetworkConnectionOutput)
}

type HubVirtualNetworkConnectionOutput struct{ *pulumi.OutputState }

func (HubVirtualNetworkConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HubVirtualNetworkConnection)(nil)).Elem()
}

func (o HubVirtualNetworkConnectionOutput) ToHubVirtualNetworkConnectionOutput() HubVirtualNetworkConnectionOutput {
	return o
}

func (o HubVirtualNetworkConnectionOutput) ToHubVirtualNetworkConnectionOutputWithContext(ctx context.Context) HubVirtualNetworkConnectionOutput {
	return o
}

func (o HubVirtualNetworkConnectionOutput) AllowHubToRemoteVnetTransit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.BoolPtrOutput { return v.AllowHubToRemoteVnetTransit }).(pulumi.BoolPtrOutput)
}

func (o HubVirtualNetworkConnectionOutput) AllowRemoteVnetToUseHubVnetGateways() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.BoolPtrOutput {
		return v.AllowRemoteVnetToUseHubVnetGateways
	}).(pulumi.BoolPtrOutput)
}

func (o HubVirtualNetworkConnectionOutput) EnableInternetSecurity() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.BoolPtrOutput { return v.EnableInternetSecurity }).(pulumi.BoolPtrOutput)
}

func (o HubVirtualNetworkConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o HubVirtualNetworkConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HubVirtualNetworkConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HubVirtualNetworkConnectionOutput) RemoteVirtualNetwork() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) SubResourceResponsePtrOutput { return v.RemoteVirtualNetwork }).(SubResourceResponsePtrOutput)
}

func (o HubVirtualNetworkConnectionOutput) RoutingConfiguration() RoutingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *HubVirtualNetworkConnection) RoutingConfigurationResponsePtrOutput {
		return v.RoutingConfiguration
	}).(RoutingConfigurationResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(HubVirtualNetworkConnectionOutput{})
}
