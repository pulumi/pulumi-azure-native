


package v20170801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualNetworkGatewayConnection struct {
	pulumi.CustomResourceState

	AuthorizationKey               pulumi.StringPtrOutput                    `pulumi:"authorizationKey"`
	ConnectionStatus               pulumi.StringOutput                       `pulumi:"connectionStatus"`
	ConnectionType                 pulumi.StringOutput                       `pulumi:"connectionType"`
	EgressBytesTransferred         pulumi.Float64Output                      `pulumi:"egressBytesTransferred"`
	EnableBgp                      pulumi.BoolPtrOutput                      `pulumi:"enableBgp"`
	Etag                           pulumi.StringPtrOutput                    `pulumi:"etag"`
	IngressBytesTransferred        pulumi.Float64Output                      `pulumi:"ingressBytesTransferred"`
	IpsecPolicies                  IpsecPolicyResponseArrayOutput            `pulumi:"ipsecPolicies"`
	LocalNetworkGateway2           LocalNetworkGatewayResponsePtrOutput      `pulumi:"localNetworkGateway2"`
	Location                       pulumi.StringPtrOutput                    `pulumi:"location"`
	Name                           pulumi.StringOutput                       `pulumi:"name"`
	Peer                           SubResourceResponsePtrOutput              `pulumi:"peer"`
	ProvisioningState              pulumi.StringOutput                       `pulumi:"provisioningState"`
	ResourceGuid                   pulumi.StringPtrOutput                    `pulumi:"resourceGuid"`
	RoutingWeight                  pulumi.IntPtrOutput                       `pulumi:"routingWeight"`
	SharedKey                      pulumi.StringPtrOutput                    `pulumi:"sharedKey"`
	Tags                           pulumi.StringMapOutput                    `pulumi:"tags"`
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
			Type: pulumi.String("azure-native:network/v20191101:VirtualNetworkGatewayConnection"),
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
	})
	opts = append(opts, aliases)
	var resource VirtualNetworkGatewayConnection
	err := ctx.RegisterResource("azure-native:network/v20170801:VirtualNetworkGatewayConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualNetworkGatewayConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkGatewayConnectionState, opts ...pulumi.ResourceOption) (*VirtualNetworkGatewayConnection, error) {
	var resource VirtualNetworkGatewayConnection
	err := ctx.ReadResource("azure-native:network/v20170801:VirtualNetworkGatewayConnection", name, id, state, &resource, opts...)
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
	ConnectionType                      string                     `pulumi:"connectionType"`
	EnableBgp                           *bool                      `pulumi:"enableBgp"`
	Etag                                *string                    `pulumi:"etag"`
	Id                                  *string                    `pulumi:"id"`
	IpsecPolicies                       []IpsecPolicy              `pulumi:"ipsecPolicies"`
	LocalNetworkGateway2                *LocalNetworkGatewayType   `pulumi:"localNetworkGateway2"`
	Location                            *string                    `pulumi:"location"`
	Peer                                *SubResource               `pulumi:"peer"`
	ResourceGroupName                   string                     `pulumi:"resourceGroupName"`
	ResourceGuid                        *string                    `pulumi:"resourceGuid"`
	RoutingWeight                       *int                       `pulumi:"routingWeight"`
	SharedKey                           *string                    `pulumi:"sharedKey"`
	Tags                                map[string]string          `pulumi:"tags"`
	UsePolicyBasedTrafficSelectors      *bool                      `pulumi:"usePolicyBasedTrafficSelectors"`
	VirtualNetworkGateway1              VirtualNetworkGatewayType  `pulumi:"virtualNetworkGateway1"`
	VirtualNetworkGateway2              *VirtualNetworkGatewayType `pulumi:"virtualNetworkGateway2"`
	VirtualNetworkGatewayConnectionName *string                    `pulumi:"virtualNetworkGatewayConnectionName"`
}


type VirtualNetworkGatewayConnectionArgs struct {
	AuthorizationKey                    pulumi.StringPtrInput
	ConnectionType                      pulumi.StringInput
	EnableBgp                           pulumi.BoolPtrInput
	Etag                                pulumi.StringPtrInput
	Id                                  pulumi.StringPtrInput
	IpsecPolicies                       IpsecPolicyArrayInput
	LocalNetworkGateway2                LocalNetworkGatewayTypePtrInput
	Location                            pulumi.StringPtrInput
	Peer                                SubResourcePtrInput
	ResourceGroupName                   pulumi.StringInput
	ResourceGuid                        pulumi.StringPtrInput
	RoutingWeight                       pulumi.IntPtrInput
	SharedKey                           pulumi.StringPtrInput
	Tags                                pulumi.StringMapInput
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
	return reflect.TypeOf((*VirtualNetworkGatewayConnection)(nil))
}

func (i *VirtualNetworkGatewayConnection) ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput {
	return i.ToVirtualNetworkGatewayConnectionOutputWithContext(context.Background())
}

func (i *VirtualNetworkGatewayConnection) ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkGatewayConnectionOutput)
}

type VirtualNetworkGatewayConnectionOutput struct{ *pulumi.OutputState }

func (VirtualNetworkGatewayConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkGatewayConnection)(nil))
}

func (o VirtualNetworkGatewayConnectionOutput) ToVirtualNetworkGatewayConnectionOutput() VirtualNetworkGatewayConnectionOutput {
	return o
}

func (o VirtualNetworkGatewayConnectionOutput) ToVirtualNetworkGatewayConnectionOutputWithContext(ctx context.Context) VirtualNetworkGatewayConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualNetworkGatewayConnectionOutput{})
}
