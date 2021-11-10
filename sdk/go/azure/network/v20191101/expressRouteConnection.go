


package v20191101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExpressRouteConnection struct {
	pulumi.CustomResourceState

	AuthorizationKey           pulumi.StringPtrOutput                     `pulumi:"authorizationKey"`
	EnableInternetSecurity     pulumi.BoolPtrOutput                       `pulumi:"enableInternetSecurity"`
	ExpressRouteCircuitPeering ExpressRouteCircuitPeeringIdResponseOutput `pulumi:"expressRouteCircuitPeering"`
	Name                       pulumi.StringOutput                        `pulumi:"name"`
	ProvisioningState          pulumi.StringOutput                        `pulumi:"provisioningState"`
	RoutingWeight              pulumi.IntPtrOutput                        `pulumi:"routingWeight"`
}


func NewExpressRouteConnection(ctx *pulumi.Context,
	name string, args *ExpressRouteConnectionArgs, opts ...pulumi.ResourceOption) (*ExpressRouteConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExpressRouteCircuitPeering == nil {
		return nil, errors.New("invalid value for required argument 'ExpressRouteCircuitPeering'")
	}
	if args.ExpressRouteGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'ExpressRouteGatewayName'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRouteConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteConnection
	err := ctx.RegisterResource("azure-native:network/v20191101:ExpressRouteConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExpressRouteConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteConnectionState, opts ...pulumi.ResourceOption) (*ExpressRouteConnection, error) {
	var resource ExpressRouteConnection
	err := ctx.ReadResource("azure-native:network/v20191101:ExpressRouteConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type expressRouteConnectionState struct {
}

type ExpressRouteConnectionState struct {
}

func (ExpressRouteConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteConnectionState)(nil)).Elem()
}

type expressRouteConnectionArgs struct {
	AuthorizationKey           *string                      `pulumi:"authorizationKey"`
	ConnectionName             *string                      `pulumi:"connectionName"`
	EnableInternetSecurity     *bool                        `pulumi:"enableInternetSecurity"`
	ExpressRouteCircuitPeering ExpressRouteCircuitPeeringId `pulumi:"expressRouteCircuitPeering"`
	ExpressRouteGatewayName    string                       `pulumi:"expressRouteGatewayName"`
	Id                         *string                      `pulumi:"id"`
	Name                       string                       `pulumi:"name"`
	ResourceGroupName          string                       `pulumi:"resourceGroupName"`
	RoutingWeight              *int                         `pulumi:"routingWeight"`
}


type ExpressRouteConnectionArgs struct {
	AuthorizationKey           pulumi.StringPtrInput
	ConnectionName             pulumi.StringPtrInput
	EnableInternetSecurity     pulumi.BoolPtrInput
	ExpressRouteCircuitPeering ExpressRouteCircuitPeeringIdInput
	ExpressRouteGatewayName    pulumi.StringInput
	Id                         pulumi.StringPtrInput
	Name                       pulumi.StringInput
	ResourceGroupName          pulumi.StringInput
	RoutingWeight              pulumi.IntPtrInput
}

func (ExpressRouteConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteConnectionArgs)(nil)).Elem()
}

type ExpressRouteConnectionInput interface {
	pulumi.Input

	ToExpressRouteConnectionOutput() ExpressRouteConnectionOutput
	ToExpressRouteConnectionOutputWithContext(ctx context.Context) ExpressRouteConnectionOutput
}

func (*ExpressRouteConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteConnection)(nil))
}

func (i *ExpressRouteConnection) ToExpressRouteConnectionOutput() ExpressRouteConnectionOutput {
	return i.ToExpressRouteConnectionOutputWithContext(context.Background())
}

func (i *ExpressRouteConnection) ToExpressRouteConnectionOutputWithContext(ctx context.Context) ExpressRouteConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteConnectionOutput)
}

type ExpressRouteConnectionOutput struct{ *pulumi.OutputState }

func (ExpressRouteConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteConnection)(nil))
}

func (o ExpressRouteConnectionOutput) ToExpressRouteConnectionOutput() ExpressRouteConnectionOutput {
	return o
}

func (o ExpressRouteConnectionOutput) ToExpressRouteConnectionOutputWithContext(ctx context.Context) ExpressRouteConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteConnectionOutput{})
}
