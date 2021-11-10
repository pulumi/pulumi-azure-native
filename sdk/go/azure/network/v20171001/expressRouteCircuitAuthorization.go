


package v20171001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExpressRouteCircuitAuthorization struct {
	pulumi.CustomResourceState

	AuthorizationKey       pulumi.StringPtrOutput `pulumi:"authorizationKey"`
	AuthorizationUseStatus pulumi.StringPtrOutput `pulumi:"authorizationUseStatus"`
	Etag                   pulumi.StringOutput    `pulumi:"etag"`
	Name                   pulumi.StringPtrOutput `pulumi:"name"`
	ProvisioningState      pulumi.StringPtrOutput `pulumi:"provisioningState"`
}


func NewExpressRouteCircuitAuthorization(ctx *pulumi.Context,
	name string, args *ExpressRouteCircuitAuthorizationArgs, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitAuthorization, error) {
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
			Type: pulumi.String("azure-native:network:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150501preview:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20150615:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160330:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160601:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20160901:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20161201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170301:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170601:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170801:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20170901:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20171101:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180101:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180401:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180601:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180701:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20180801:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181001:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181101:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20181201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190401:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190601:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190701:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190801:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ExpressRouteCircuitAuthorization"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:ExpressRouteCircuitAuthorization"),
		},
	})
	opts = append(opts, aliases)
	var resource ExpressRouteCircuitAuthorization
	err := ctx.RegisterResource("azure-native:network/v20171001:ExpressRouteCircuitAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetExpressRouteCircuitAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExpressRouteCircuitAuthorizationState, opts ...pulumi.ResourceOption) (*ExpressRouteCircuitAuthorization, error) {
	var resource ExpressRouteCircuitAuthorization
	err := ctx.ReadResource("azure-native:network/v20171001:ExpressRouteCircuitAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type expressRouteCircuitAuthorizationState struct {
}

type ExpressRouteCircuitAuthorizationState struct {
}

func (ExpressRouteCircuitAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitAuthorizationState)(nil)).Elem()
}

type expressRouteCircuitAuthorizationArgs struct {
	AuthorizationKey       *string `pulumi:"authorizationKey"`
	AuthorizationName      *string `pulumi:"authorizationName"`
	AuthorizationUseStatus *string `pulumi:"authorizationUseStatus"`
	CircuitName            string  `pulumi:"circuitName"`
	Id                     *string `pulumi:"id"`
	Name                   *string `pulumi:"name"`
	ProvisioningState      *string `pulumi:"provisioningState"`
	ResourceGroupName      string  `pulumi:"resourceGroupName"`
}


type ExpressRouteCircuitAuthorizationArgs struct {
	AuthorizationKey       pulumi.StringPtrInput
	AuthorizationName      pulumi.StringPtrInput
	AuthorizationUseStatus pulumi.StringPtrInput
	CircuitName            pulumi.StringInput
	Id                     pulumi.StringPtrInput
	Name                   pulumi.StringPtrInput
	ProvisioningState      pulumi.StringPtrInput
	ResourceGroupName      pulumi.StringInput
}

func (ExpressRouteCircuitAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*expressRouteCircuitAuthorizationArgs)(nil)).Elem()
}

type ExpressRouteCircuitAuthorizationInput interface {
	pulumi.Input

	ToExpressRouteCircuitAuthorizationOutput() ExpressRouteCircuitAuthorizationOutput
	ToExpressRouteCircuitAuthorizationOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationOutput
}

func (*ExpressRouteCircuitAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorization)(nil))
}

func (i *ExpressRouteCircuitAuthorization) ToExpressRouteCircuitAuthorizationOutput() ExpressRouteCircuitAuthorizationOutput {
	return i.ToExpressRouteCircuitAuthorizationOutputWithContext(context.Background())
}

func (i *ExpressRouteCircuitAuthorization) ToExpressRouteCircuitAuthorizationOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExpressRouteCircuitAuthorizationOutput)
}

type ExpressRouteCircuitAuthorizationOutput struct{ *pulumi.OutputState }

func (ExpressRouteCircuitAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressRouteCircuitAuthorization)(nil))
}

func (o ExpressRouteCircuitAuthorizationOutput) ToExpressRouteCircuitAuthorizationOutput() ExpressRouteCircuitAuthorizationOutput {
	return o
}

func (o ExpressRouteCircuitAuthorizationOutput) ToExpressRouteCircuitAuthorizationOutputWithContext(ctx context.Context) ExpressRouteCircuitAuthorizationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ExpressRouteCircuitAuthorizationOutput{})
}
