


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GatewayRouteConfig struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties GatewayRouteConfigPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                   `pulumi:"systemData"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewGatewayRouteConfig(ctx *pulumi.Context,
	name string, args *GatewayRouteConfigArgs, opts ...pulumi.ResourceOption) (*GatewayRouteConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayName == nil {
		return nil, errors.New("invalid value for required argument 'GatewayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:GatewayRouteConfig"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:GatewayRouteConfig"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220501preview:GatewayRouteConfig"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:GatewayRouteConfig"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:GatewayRouteConfig"),
		},
	})
	opts = append(opts, aliases)
	var resource GatewayRouteConfig
	err := ctx.RegisterResource("azure-native:appplatform/v20220101preview:GatewayRouteConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGatewayRouteConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayRouteConfigState, opts ...pulumi.ResourceOption) (*GatewayRouteConfig, error) {
	var resource GatewayRouteConfig
	err := ctx.ReadResource("azure-native:appplatform/v20220101preview:GatewayRouteConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gatewayRouteConfigState struct {
}

type GatewayRouteConfigState struct {
}

func (GatewayRouteConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteConfigState)(nil)).Elem()
}

type gatewayRouteConfigArgs struct {
	GatewayName       string                        `pulumi:"gatewayName"`
	Properties        *GatewayRouteConfigProperties `pulumi:"properties"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	RouteConfigName   *string                       `pulumi:"routeConfigName"`
	ServiceName       string                        `pulumi:"serviceName"`
}


type GatewayRouteConfigArgs struct {
	GatewayName       pulumi.StringInput
	Properties        GatewayRouteConfigPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	RouteConfigName   pulumi.StringPtrInput
	ServiceName       pulumi.StringInput
}

func (GatewayRouteConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayRouteConfigArgs)(nil)).Elem()
}

type GatewayRouteConfigInput interface {
	pulumi.Input

	ToGatewayRouteConfigOutput() GatewayRouteConfigOutput
	ToGatewayRouteConfigOutputWithContext(ctx context.Context) GatewayRouteConfigOutput
}

func (*GatewayRouteConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRouteConfig)(nil)).Elem()
}

func (i *GatewayRouteConfig) ToGatewayRouteConfigOutput() GatewayRouteConfigOutput {
	return i.ToGatewayRouteConfigOutputWithContext(context.Background())
}

func (i *GatewayRouteConfig) ToGatewayRouteConfigOutputWithContext(ctx context.Context) GatewayRouteConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayRouteConfigOutput)
}

type GatewayRouteConfigOutput struct{ *pulumi.OutputState }

func (GatewayRouteConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GatewayRouteConfig)(nil)).Elem()
}

func (o GatewayRouteConfigOutput) ToGatewayRouteConfigOutput() GatewayRouteConfigOutput {
	return o
}

func (o GatewayRouteConfigOutput) ToGatewayRouteConfigOutputWithContext(ctx context.Context) GatewayRouteConfigOutput {
	return o
}

func (o GatewayRouteConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRouteConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GatewayRouteConfigOutput) Properties() GatewayRouteConfigPropertiesResponseOutput {
	return o.ApplyT(func(v *GatewayRouteConfig) GatewayRouteConfigPropertiesResponseOutput { return v.Properties }).(GatewayRouteConfigPropertiesResponseOutput)
}

func (o GatewayRouteConfigOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *GatewayRouteConfig) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GatewayRouteConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GatewayRouteConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayRouteConfigOutput{})
}
