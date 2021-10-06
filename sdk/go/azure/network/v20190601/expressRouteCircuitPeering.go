


package v20190601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExpressRouteCircuitPeering struct {
	pulumi.CustomResourceState

	AzureASN                   pulumi.IntPtrOutput                                   `pulumi:"azureASN"`
	Connections                ExpressRouteCircuitConnectionResponseArrayOutput      `pulumi:"connections"`
	Etag                       pulumi.StringOutput                                   `pulumi:"etag"`
	ExpressRouteConnection     ExpressRouteConnectionIdResponsePtrOutput             `pulumi:"expressRouteConnection"`
	GatewayManagerEtag         pulumi.StringPtrOutput                                `pulumi:"gatewayManagerEtag"`
	Ipv6PeeringConfig          Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             pulumi.StringPtrOutput                                `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigResponsePtrOutput     `pulumi:"microsoftPeeringConfig"`
	Name                       pulumi.StringPtrOutput                                `pulumi:"name"`
	PeerASN                    pulumi.Float64PtrOutput                               `pulumi:"peerASN"`
	PeeredConnections          PeerExpressRouteCircuitConnectionResponseArrayOutput  `pulumi:"peeredConnections"`
	PeeringType                pulumi.StringPtrOutput                                `pulumi:"peeringType"`
	PrimaryAzurePort           pulumi.StringPtrOutput                                `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   pulumi.StringPtrOutput                                `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          pulumi.StringPtrOutput                                `pulumi:"provisioningState"`
	RouteFilter                SubResourceResponsePtrOutput                          `pulumi:"routeFilter"`
	SecondaryAzurePort         pulumi.StringPtrOutput                                `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix pulumi.StringPtrOutput                                `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  pulumi.StringPtrOutput                                `pulumi:"sharedKey"`
	State                      pulumi.StringPtrOutput                                `pulumi:"state"`
	Stats                      ExpressRouteCircuitStatsResponsePtrOutput             `pulumi:"stats"`
	Type                       pulumi.StringOutput                                   `pulumi:"type"`
	VlanId                     pulumi.IntPtrOutput                                   `pulumi:"vlanId"`
}


func NewExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitPeeringArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CircuitName == nil {
		return nil, errors.New("invalid value for required argument 'CircuitName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150501preview:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150615:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160330:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20161201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170301:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171001:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200801:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20201101:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210201:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteCircuitPeering"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20210301:ExpressRouteCircuitPeering"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteCircuitPeering
	err := ctx.RegisterResource("azure-native:network/v20190601:ExpressRouteCircuitPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExpressRouteCircuitPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitPeeringState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitPeering, error) {
	var resource ExpressRouteCircuitPeering
	err := ctx.ReadResource("azure-native:network/v20190601:ExpressRouteCircuitPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type expressRouteCircuitPeeringState struct {
}

type ExpressRouteCircuitPeeringState struct {
}

func (ExpressRouteCircuitPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringState)(nil)).Elem()
}

type expressRouteCircuitPeeringArgs struct {
	AzureASN                   *int                                  `pulumi:"azureASN"`
	CircuitName                string                                `pulumi:"circuitName"`
	Connections                []ExpressRouteCircuitConnectionType   `pulumi:"connections"`
	GatewayManagerEtag         *string                               `pulumi:"gatewayManagerEtag"`
	Id                         *string                               `pulumi:"id"`
	Ipv6PeeringConfig          *Ipv6ExpressRouteCircuitPeeringConfig `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             *string                               `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfig     `pulumi:"microsoftPeeringConfig"`
	Name                       *string                               `pulumi:"name"`
	PeerASN                    *float64                              `pulumi:"peerASN"`
	PeeringName                *string                               `pulumi:"peeringName"`
	PeeringType                *string                               `pulumi:"peeringType"`
	PrimaryAzurePort           *string                               `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   *string                               `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          *string                               `pulumi:"provisioningState"`
	ResourceGroupName          string                                `pulumi:"resourceGroupName"`
	RouteFilter                *SubResource                          `pulumi:"routeFilter"`
	SecondaryAzurePort         *string                               `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix *string                               `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  *string                               `pulumi:"sharedKey"`
	State                      *string                               `pulumi:"state"`
	Stats                      *ExpressRouteCircuitStats             `pulumi:"stats"`
	VlanId                     *int                                  `pulumi:"vlanId"`
}


type ExpressRouteCircuitPeeringArgs struct {
	AzureASN                   pulumi.IntPtrInput
	CircuitName                pulumi.StringInput
	Connections                ExpressRouteCircuitConnectionTypeArrayInput
	GatewayManagerEtag         pulumi.StringPtrInput
	Id                         pulumi.StringPtrInput
	Ipv6PeeringConfig          Ipv6ExpressRouteCircuitPeeringConfigPtrInput
	LastModifiedBy             pulumi.StringPtrInput
	MicrosoftPeeringConfig     ExpressRouteCircuitPeeringConfigPtrInput
	Name                       pulumi.StringPtrInput
	PeerASN                    pulumi.Float64PtrInput
	PeeringName                pulumi.StringPtrInput
	PeeringType                pulumi.StringPtrInput
	PrimaryAzurePort           pulumi.StringPtrInput
	PrimaryPeerAddressPrefix   pulumi.StringPtrInput
	ProvisioningState          pulumi.StringPtrInput
	ResourceGroupName          pulumi.StringInput
	RouteFilter                SubResourcePtrInput
	SecondaryAzurePort         pulumi.StringPtrInput
	SecondaryPeerAddressPrefix pulumi.StringPtrInput
	SharedKey                  pulumi.StringPtrInput
	State                      pulumi.StringPtrInput
	Stats                      ExpressRouteCircuitStatsPtrInput
	VlanId                     pulumi.IntPtrInput
}

func (ExpressRouteCircuitPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitPeeringArgs)(nil)).Elem()
}

type ExpressRouteCircuitPeeringInput interface {
	pulumi.Input

	ToExpressRouteCircuitPeeringOutput() ExpressRouteCircuitPeeringOutput
	ToExpressRouteCircuitPeeringOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringOutput
}

func (*ExpressRouteCircuitPeering) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeering)(nil))
}

func (i *ExpressRouteCircuitPeering) ToExpressRouteCircuitPeeringOutput() ExpressRouteCircuitPeeringOutput {
	return i.ToExpressRouteCircuitPeeringOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuitPeering) ToExpressRouteCircuitPeeringOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitPeeringOutput)
}

type ExpressRouteCircuitPeeringOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitPeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitPeering)(nil))
}

func (o ExpressRouteCircuitPeeringOutput) ToExpressRouteCircuitPeeringOutput() ExpressRouteCircuitPeeringOutput {
	return o
}

func (o ExpressRouteCircuitPeeringOutput) ToExpressRouteCircuitPeeringOutputWithContext(ctx context.Context) ExpressRouteCircuitPeeringOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCircuitPeeringOutput{})
}
