


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Gateway struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput             `pulumi:"name"`
	Properties GatewayPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput            `pulumi:"sku"`
	SystemData SystemDataResponseOutput        `pulumi:"systemData"`
	Type       pulumi.StringOutput             `pulumi:"type"`
}


func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.Properties != nil {
		args.Properties = args.Properties.ToGatewayPropertiesPtrOutput().ApplyT(func(v *GatewayProperties) *GatewayProperties { return v.Defaults() }).(GatewayPropertiesPtrOutput)
	}
	if args.Sku != nil {
		args.Sku = args.Sku.ToSkuPtrOutput().ApplyT(func(v *Sku) *Sku { return v.Defaults() }).(SkuPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:appplatform:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220101preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220301preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20220901preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:appplatform/v20221101preview:Gateway"),
		},
	})
	opts = append(opts, aliases)
	var resource Gateway
	err := ctx.RegisterResource("azure-native:appplatform/v20220501preview:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("azure-native:appplatform/v20220501preview:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gatewayState struct {
}

type GatewayState struct {
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	GatewayName       *string            `pulumi:"gatewayName"`
	Properties        *GatewayProperties `pulumi:"properties"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	ServiceName       string             `pulumi:"serviceName"`
	Sku               *Sku               `pulumi:"sku"`
}


type GatewayArgs struct {
	GatewayName       pulumi.StringPtrInput
	Properties        GatewayPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
	Sku               SkuPtrInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

func (o GatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GatewayOutput) Properties() GatewayPropertiesResponseOutput {
	return o.ApplyT(func(v *Gateway) GatewayPropertiesResponseOutput { return v.Properties }).(GatewayPropertiesResponseOutput)
}

func (o GatewayOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Gateway) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o GatewayOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Gateway) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o GatewayOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GatewayOutput{})
}
