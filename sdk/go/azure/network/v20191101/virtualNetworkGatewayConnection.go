


package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkGatewayConnection struct {
	pulumi.CustomResourceState

	AuthorizationKey               pulumi.StringPtrOutput                    `pulumi:"authorizationKey"`
	ConnectionProtocol             pulumi.StringPtrOutput                    `pulumi:"connectionProtocol"`
	ConnectionStatus               pulumi.StringOutput                       `pulumi:"connectionStatus"`
	ConnectionType                 pulumi.StringOutput                       `pulumi:"connectionType"`
	EgressBytesTransferred         pulumi.Float64Output                      `pulumi:"egressBytesTransferred"`
	EnableBgp                      pulumi.BoolPtrOutput                      `pulumi:"enableBgp"`
	Etag                           pulumi.StringOutput                       `pulumi:"etag"`
	ExpressRouteGatewayBypass      pulumi.BoolPtrOutput                      `pulumi:"expressRouteGatewayBypass"`
	IngressBytesTransferred        pulumi.Float64Output                      `pulumi:"ingressBytesTransferred"`
	IpsecPolicies                  IpsecPolicyResponseArrayOutput            `pulumi:"ipsecPolicies"`
	LocalNetworkGateway2           LocalNetworkGatewayResponsePtrOutput      `pulumi:"localNetworkGateway2"`
	Location                       pulumi.StringPtrOutput                    `pulumi:"location"`
	Name                           pulumi.StringOutput                       `pulumi:"name"`
	Peer                           SubResourceResponsePtrOutput              `pulumi:"peer"`
	ProvisioningState              pulumi.StringOutput                       `pulumi:"provisioningState"`
	ResourceGuid                   pulumi.StringOutput                       `pulumi:"resourceGuid"`
	RoutingWeight                  pulumi.IntPtrOutput                       `pulumi:"routingWeight"`
	SharedKey                      pulumi.StringPtrOutput                    `pulumi:"sharedKey"`
	Tags                           pulumi.StringMapOutput                    `pulumi:"tags"`
	TrafficSelectorPolicies        TrafficSelectorPolicyResponseArrayOutput  `pulumi:"trafficSelectorPolicies"`
	TunnelConnectionStatus         TunnelConnectionHealthResponseArrayOutput `pulumi:"tunnelConnectionStatus"`
	Type                           pulumi.StringOutput                       `pulumi:"type"`
	UsePolicyBasedTrafficSelectors pulumi.BoolPtrOutput                      `pulumi:"usePolicyBasedTrafficSelectors"`
	VirtualNetworkGateway1         VirtualNetworkGatewayResponseOutput       `pulumi:"virtualNetworkGateway1"`
	VirtualNetworkGateway2         VirtualNetworkGatewayResponsePtrOutput    `pulumi:"virtualNetworkGateway2"`
}


func NewVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, args *VirtualNetworkGatewayConnectionArgs, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionType == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualNetworkGateway1 == nil {
		return nil, errors.New("invalid value for required argument 'VirtualNetworkGateway1'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:VirtualNetworkGatewayConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:VirtualNetworkGatewayConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkGatewayConnection
	err := ctx.RegisterResource("azure-native:network/v20191101:VirtualNetworkGatewayConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayConnectionState, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayConnection, error) {
	var resource VirtualNetworkGatewayConnection
	err := ctx.ReadResource("azure-native:network/v20191101:VirtualNetworkGatewayConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualNetworkGatewayConnectionState struct {
}

type VirtualNetworkGatewayConnectionState struct {
}

func (VirtualNetworkGatewayConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayConnectionState)(nil)).Elem()
}

type virtualNetworkGatewayConnectionArgs struct {
	AuthorizationKey                    *string                    `pulumi:"authorizationKey"`
	ConnectionProtocol                  *string                    `pulumi:"connectionProtocol"`
	ConnectionType                      string                     `pulumi:"connectionType"`
	EnableBgp                           *bool                      `pulumi:"enableBgp"`
	ExpressRouteGatewayBypass           *bool                      `pulumi:"expressRouteGatewayBypass"`
	Id                                  *string                    `pulumi:"id"`
	IpsecPolicies                       []IpsecPolicy              `pulumi:"ipsecPolicies"`
	LocalNetworkGateway2                *LocalNetworkGatewayType   `pulumi:"localNetworkGateway2"`
	Location                            *string                    `pulumi:"location"`
	Peer                                *SubResource               `pulumi:"peer"`
	ResourceGroupName                   string                     `pulumi:"resourceGroupName"`
	RoutingWeight                       *int                       `pulumi:"routingWeight"`
	SharedKey                           *string                    `pulumi:"sharedKey"`
	Tags                                map[string]string          `pulumi:"tags"`
	TrafficSelectorPolicies             []TrafficSelectorPolicy    `pulumi:"trafficSelectorPolicies"`
	UsePolicyBasedTrafficSelectors      *bool                      `pulumi:"usePolicyBasedTrafficSelectors"`
	VirtualNetworkGateway1              VirtualNetworkGatewayType  `pulumi:"virtualNetworkGateway1"`
	VirtualNetworkGateway2              *VirtualNetworkGatewayType `pulumi:"virtualNetworkGateway2"`
	VirtualNetworkGatewayConnectionName *string                    `pulumi:"virtualNetworkGatewayConnectionName"`
}


type VirtualNetworkGatewayConnectionArgs struct {
	AuthorizationKey                    pulumi.StringPtrInput
	ConnectionProtocol                  pulumi.StringPtrInput
	ConnectionType                      pulumi.StringInput
	EnableBgp                           pulumi.BoolPtrInput
	ExpressRouteGatewayBypass           pulumi.BoolPtrInput
	Id                                  pulumi.StringPtrInput
	IpsecPolicies                       IpsecPolicyArrayInput
	LocalNetworkGateway2                LocalNetworkGatewayTypePtrInput
	Location                            pulumi.StringPtrInput
	Peer                                SubResourcePtrInput
	ResourceGroupName                   pulumi.StringInput
	RoutingWeight                       pulumi.IntPtrInput
	SharedKey                           pulumi.StringPtrInput
	Tags                                pulumi.StringMapInput
	TrafficSelectorPolicies             TrafficSelectorPolicyArrayInput
	UsePolicyBasedTrafficSelectors      pulumi.BoolPtrInput
	VirtualNetworkGateway1              VirtualNetworkGatewayTypeInput
	VirtualNetworkGateway2              VirtualNetworkGatewayTypePtrInput
	VirtualNetworkGatewayConnectionName pulumi.StringPtrInput
}

func (VirtualNetworkGatewayConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkGatewayConnectionArgs)(nil)).Elem()
}

type VirtualNetworkGatewayConnectionInput interface {
	pulumi.Input

	ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput
	ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput
}

func (*VirtualNetworkGatewayConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayConnection)(nil)).Elem()
}

func (i *VirtualNetworkGatewayConnection) ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput {
	return i.ToVirtualNetworkGatewayConnectionOutputWithContext(context.Background())
}

func (i *VirtualNetworkGatewayConnection) ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayConnectionOutput)
}

type VirtualNetworkGatewayConnectionOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkGatewayConnection)(nil)).Elem()
}

func (o VirtualNetworkGatewayConnectionOutput) ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionOutput) ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) ConnectionProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.ConnectionProtocol }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringOutput { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) ConnectionType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringOutput { return v.ConnectionType }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.Float64Output { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o VirtualNetworkGatewayConnectionOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.BoolPtrOutput { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) ExpressRouteGatewayBypass() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.BoolPtrOutput { return v.ExpressRouteGatewayBypass }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.Float64Output { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o VirtualNetworkGatewayConnectionOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) IpsecPolicyResponseArrayOutput { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) LocalNetworkGateway2() LocalNetworkGatewayResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) LocalNetworkGatewayResponsePtrOutput {
		return v.LocalNetworkGateway2
	}).(LocalNetworkGatewayResponsePtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Peer() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) SubResourceResponsePtrOutput { return v.Peer }).(SubResourceResponsePtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringOutput { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.IntPtrOutput { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringPtrOutput { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) TrafficSelectorPolicies() TrafficSelectorPolicyResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) TrafficSelectorPolicyResponseArrayOutput {
		return v.TrafficSelectorPolicies
	}).(TrafficSelectorPolicyResponseArrayOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) TunnelConnectionStatus() TunnelConnectionHealthResponseArrayOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) TunnelConnectionHealthResponseArrayOutput {
		return v.TunnelConnectionStatus
	}).(TunnelConnectionHealthResponseArrayOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) pulumi.BoolPtrOutput { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) VirtualNetworkGateway1() VirtualNetworkGatewayResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) VirtualNetworkGatewayResponseOutput {
		return v.VirtualNetworkGateway1
	}).(VirtualNetworkGatewayResponseOutput)
}

func (o VirtualNetworkGatewayConnectionOutput) VirtualNetworkGateway2() VirtualNetworkGatewayResponsePtrOutput {
	return o.ApplyT(func(v *VirtualNetworkGatewayConnection) VirtualNetworkGatewayResponsePtrOutput {
		return v.VirtualNetworkGateway2
	}).(VirtualNetworkGatewayResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionOutput{})
}
