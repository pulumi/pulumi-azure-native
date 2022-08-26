


package network

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExpressRouteCircuitConnection struct {
	pulumi.CustomResourceState

	AddressPrefix                  pulumi.StringPtrOutput                       `pulumi:"addressPrefix"`
	AuthorizationKey               pulumi.StringPtrOutput                       `pulumi:"authorizationKey"`
	CircuitConnectionStatus        pulumi.StringOutput                          `pulumi:"circuitConnectionStatus"`
	Etag                           pulumi.StringOutput                          `pulumi:"etag"`
	ExpressRouteCircuitPeering     SubResourceResponsePtrOutput                 `pulumi:"expressRouteCircuitPeering"`
	Ipv6CircuitConnectionConfig    Ipv6CircuitConnectionConfigResponsePtrOutput `pulumi:"ipv6CircuitConnectionConfig"`
	Name                           pulumi.StringPtrOutput                       `pulumi:"name"`
	PeerExpressRouteCircuitPeering SubResourceResponsePtrOutput                 `pulumi:"peerExpressRouteCircuitPeering"`
	ProvisioningState              pulumi.StringOutput                          `pulumi:"provisioningState"`
	Type                           pulumi.StringOutput                          `pulumi:"type"`
}


func NewExpressRouteCircuitConnection(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitConnectionArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CircuitName == nil {
		return nil, errors.New("invalid value for required argument 'CircuitName'")
	}
	if args.PeeringName == nil {
		return nil, errors.New("invalid value for required argument 'PeeringName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network/v20180201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRouteCircuitConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRouteCircuitConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteCircuitConnection
	err := ctx.RegisterResource("azure-native:network:ExpressRouteCircuitConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExpressRouteCircuitConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitConnectionState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitConnection, error) {
	var resource ExpressRouteCircuitConnection
	err := ctx.ReadResource("azure-native:network:ExpressRouteCircuitConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type expressRouteCircuitConnectionState struct {
}

type ExpressRouteCircuitConnectionState struct {
}

func (ExpressRouteCircuitConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitConnectionState)(nil)).Elem()
}

type expressRouteCircuitConnectionArgs struct {
	AddressPrefix                  *string                      `pulumi:"addressPrefix"`
	AuthorizationKey               *string                      `pulumi:"authorizationKey"`
	CircuitName                    string                       `pulumi:"circuitName"`
	ConnectionName                 *string                      `pulumi:"connectionName"`
	ExpressRouteCircuitPeering     *SubResource                 `pulumi:"expressRouteCircuitPeering"`
	Id                             *string                      `pulumi:"id"`
	Ipv6CircuitConnectionConfig    *Ipv6CircuitConnectionConfig `pulumi:"ipv6CircuitConnectionConfig"`
	Name                           *string                      `pulumi:"name"`
	PeerExpressRouteCircuitPeering *SubResource                 `pulumi:"peerExpressRouteCircuitPeering"`
	PeeringName                    string                       `pulumi:"peeringName"`
	ResourceGroupName              string                       `pulumi:"resourceGroupName"`
}


type ExpressRouteCircuitConnectionArgs struct {
	AddressPrefix                  pulumi.StringPtrInput
	AuthorizationKey               pulumi.StringPtrInput
	CircuitName                    pulumi.StringInput
	ConnectionName                 pulumi.StringPtrInput
	ExpressRouteCircuitPeering     SubResourcePtrInput
	Id                             pulumi.StringPtrInput
	Ipv6CircuitConnectionConfig    Ipv6CircuitConnectionConfigPtrInput
	Name                           pulumi.StringPtrInput
	PeerExpressRouteCircuitPeering SubResourcePtrInput
	PeeringName                    pulumi.StringInput
	ResourceGroupName              pulumi.StringInput
}

func (ExpressRouteCircuitConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitConnectionArgs)(nil)).Elem()
}

type ExpressRouteCircuitConnectionInput interface {
	pulumi.Input

	ToExpressRouteCircuitConnectionOutput() ExpressRouteCircuitConnectionOutput
	ToExpressRouteCircuitConnectionOutputWithContext(ctx context.Context) ExpressRouteCircuitConnectionOutput
}

func (*ExpressRouteCircuitConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitConnection)(nil)).Elem()
}

func (i *ExpressRouteCircuitConnection) ToExpressRouteCircuitConnectionOutput() ExpressRouteCircuitConnectionOutput {
	return i.ToExpressRouteCircuitConnectionOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuitConnection) ToExpressRouteCircuitConnectionOutputWithContext(ctx context.Context) ExpressRouteCircuitConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitConnectionOutput)
}

type ExpressRouteCircuitConnectionOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCircuitConnection)(nil)).Elem()
}

func (o ExpressRouteCircuitConnectionOutput) ToExpressRouteCircuitConnectionOutput() ExpressRouteCircuitConnectionOutput {
	return o
}

func (o ExpressRouteCircuitConnectionOutput) ToExpressRouteCircuitConnectionOutputWithContext(ctx context.Context) ExpressRouteCircuitConnectionOutput {
	return o
}

func (o ExpressRouteCircuitConnectionOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitConnection) pulumi.StringPtrOutput { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitConnectionOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitConnection) pulumi.StringPtrOutput { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitConnectionOutput) CircuitConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitConnection) pulumi.StringOutput { return v.CircuitConnectionStatus }).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitConnectionOutput) ExpressRouteCircuitPeering() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitConnection) SubResourceResponsePtrOutput {
		return v.ExpressRouteCircuitPeering
	}).(SubResourceResponsePtrOutput)
}

func (o ExpressRouteCircuitConnectionOutput) Ipv6CircuitConnectionConfig() Ipv6CircuitConnectionConfigResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitConnection) Ipv6CircuitConnectionConfigResponsePtrOutput {
		return v.Ipv6CircuitConnectionConfig
	}).(Ipv6CircuitConnectionConfigResponsePtrOutput)
}

func (o ExpressRouteCircuitConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitConnection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCircuitConnectionOutput) PeerExpressRouteCircuitPeering() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitConnection) SubResourceResponsePtrOutput {
		return v.PeerExpressRouteCircuitPeering
	}).(SubResourceResponsePtrOutput)
}

func (o ExpressRouteCircuitConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ExpressRouteCircuitConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCircuitConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCircuitConnectionOutput{})
}
