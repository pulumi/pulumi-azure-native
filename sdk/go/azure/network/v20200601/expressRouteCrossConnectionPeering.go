


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExpressRouteCrossConnectionPeering struct {
	pulumi.CustomResourceState

	AzureASN                   pulumi.IntOutput                                      `pulumi:"azureASN"`
	Etag                       pulumi.StringOutput                                   `pulumi:"etag"`
	GatewayManagerEtag         pulumi.StringPtrOutput                                `pulumi:"gatewayManagerEtag"`
	Ipv6PeeringConfig          Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             pulumi.StringOutput                                   `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigResponsePtrOutput     `pulumi:"microsoftPeeringConfig"`
	Name                       pulumi.StringPtrOutput                                `pulumi:"name"`
	PeerASN                    pulumi.Float64PtrOutput                               `pulumi:"peerASN"`
	PeeringType                pulumi.StringPtrOutput                                `pulumi:"peeringType"`
	PrimaryAzurePort           pulumi.StringOutput                                   `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   pulumi.StringPtrOutput                                `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          pulumi.StringOutput                                   `pulumi:"provisioningState"`
	SecondaryAzurePort         pulumi.StringOutput                                   `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix pulumi.StringPtrOutput                                `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  pulumi.StringPtrOutput                                `pulumi:"sharedKey"`
	State                      pulumi.StringPtrOutput                                `pulumi:"state"`
	VlanId                     pulumi.IntPtrOutput                                   `pulumi:"vlanId"`
}


func NewExpressRouteCrossConnectionPeering(ctx *pulumi.Context,
	name string, args *ExpressRouteCrossConnectionPeeringArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCrossConnectionPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CrossConnectionName == nil {
		return nil, errors.New("invalid value for required argument 'CrossConnectionName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ExpressRouteCrossConnectionPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220501:ExpressRouteCrossConnectionPeering"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteCrossConnectionPeering
	err := ctx.RegisterResource("azure-native:network/v20200601:ExpressRouteCrossConnectionPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExpressRouteCrossConnectionPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCrossConnectionPeeringState, opts ...pulumi.ResourceOption) (*ExpressRouteCrossConnectionPeering, error) {
	var resource ExpressRouteCrossConnectionPeering
	err := ctx.ReadResource("azure-native:network/v20200601:ExpressRouteCrossConnectionPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type expressRouteCrossConnectionPeeringState struct {
}

type ExpressRouteCrossConnectionPeeringState struct {
}

func (ExpressRouteCrossConnectionPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCrossConnectionPeeringState)(nil)).Elem()
}

type expressRouteCrossConnectionPeeringArgs struct {
	CrossConnectionName        string                                `pulumi:"crossConnectionName"`
	GatewayManagerEtag         *string                               `pulumi:"gatewayManagerEtag"`
	Id                         *string                               `pulumi:"id"`
	Ipv6PeeringConfig          *Ipv6ExpressRouteCircuitPeeringConfig `pulumi:"ipv6PeeringConfig"`
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfig     `pulumi:"microsoftPeeringConfig"`
	Name                       *string                               `pulumi:"name"`
	PeerASN                    *float64                              `pulumi:"peerASN"`
	PeeringName                *string                               `pulumi:"peeringName"`
	PeeringType                *string                               `pulumi:"peeringType"`
	PrimaryPeerAddressPrefix   *string                               `pulumi:"primaryPeerAddressPrefix"`
	ResourceGroupName          string                                `pulumi:"resourceGroupName"`
	SecondaryPeerAddressPrefix *string                               `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  *string                               `pulumi:"sharedKey"`
	State                      *string                               `pulumi:"state"`
	VlanId                     *int                                  `pulumi:"vlanId"`
}


type ExpressRouteCrossConnectionPeeringArgs struct {
	CrossConnectionName        pulumi.StringInput
	GatewayManagerEtag         pulumi.StringPtrInput
	Id                         pulumi.StringPtrInput
	Ipv6PeeringConfig          Ipv6ExpressRouteCircuitPeeringConfigPtrInput
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigPtrInput
	Name                       pulumi.StringPtrInput
	PeerASN                    pulumi.Float64PtrInput
	PeeringName                pulumi.StringPtrInput
	PeeringType                pulumi.StringPtrInput
	PrimaryPeerAddressPrefix   pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	SecondaryPeerAddressPrefix pulumi.StringPtrInput
	SharedKey                  pulumi.StringPtrInput
	State                      pulumi.StringPtrInput
	VlanId                     pulumi.IntPtrInput
}

func (ExpressRouteCrossConnectionPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCrossConnectionPeeringArgs)(nil)).Elem()
}

type ExpressRouteCrossConnectionPeeringInput interface {
	pulumi.Input

	ToExpressRouteCrossConnectionPeeringOutput() ExpressRouteCrossConnectionPeeringOutput
	ToExpressRouteCrossConnectionPeeringOutputWithContext(ctx context.Context) ExpressRouteCrossConnectionPeeringOutput
}

func (*ExpressRouteCrossConnectionPeering) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCrossConnectionPeering)(nil)).Elem()
}

func (i *ExpressRouteCrossConnectionPeering) ToExpressRouteCrossConnectionPeeringOutput() ExpressRouteCrossConnectionPeeringOutput {
	return i.ToExpressRouteCrossConnectionPeeringOutputWithContext(context.Background())
}

func (i *ExpressRouteCrossConnectionPeering) ToExpressRouteCrossConnectionPeeringOutputWithContext(ctx context.Context) ExpressRouteCrossConnectionPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCrossConnectionPeeringOutput)
}

type ExpressRouteCrossConnectionPeeringOutput struct{ *pulumi.OutputState }

func (ExpressRouteCrossConnectionPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExpressRouteCrossConnectionPeering)(nil)).Elem()
}

func (o ExpressRouteCrossConnectionPeeringOutput) ToExpressRouteCrossConnectionPeeringOutput() ExpressRouteCrossConnectionPeeringOutput {
	return o
}

func (o ExpressRouteCrossConnectionPeeringOutput) ToExpressRouteCrossConnectionPeeringOutputWithContext(ctx context.Context) ExpressRouteCrossConnectionPeeringOutput {
	return o
}

func (o ExpressRouteCrossConnectionPeeringOutput) AzureASN() pulumi.IntOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.IntOutput { return v.AzureASN }).(pulumi.IntOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringOutput { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) ExpressRouteCircuitPeeringConfigResponsePtrOutput {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) PeerASN() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.Float64PtrOutput { return v.PeerASN }).(pulumi.Float64PtrOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.PeeringType }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) PrimaryAzurePort() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringOutput { return v.PrimaryAzurePort }).(pulumi.StringOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) SecondaryAzurePort() pulumi.StringOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringOutput { return v.SecondaryAzurePort }).(pulumi.StringOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput {
		return v.SecondaryPeerAddressPrefix
	}).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o ExpressRouteCrossConnectionPeeringOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ExpressRouteCrossConnectionPeering) pulumi.IntPtrOutput { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCrossConnectionPeeringOutput{})
}
