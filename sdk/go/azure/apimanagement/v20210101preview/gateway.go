


package v20210101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Gateway struct {
	pulumi.CustomResourceState

	Description  pulumi.StringPtrOutput                        `pulumi:"description"`
	LocationData ResourceLocationDataContractResponsePtrOutput `pulumi:"locationData"`
	Name         pulumi.StringOutput                           `pulumi:"name"`
	Type         pulumi.StringOutput                           `pulumi:"type"`
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210101preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement:Gateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201:Gateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20191201preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20200601preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20201201:Gateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20201201:Gateway"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:Gateway"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20210401preview:Gateway"),
		},
	})
	opts = append(opts, aliases)
	var resource Gateway
	err := ctx.RegisterResource("azure-native:apimanagement/v20210101preview:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("azure-native:apimanagement/v20210101preview:Gateway", name, id, state, &resource, opts...)
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
	Description       *string                       `pulumi:"description"`
	GatewayId         *string                       `pulumi:"gatewayId"`
	LocationData      *ResourceLocationDataContract `pulumi:"locationData"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	ServiceName       string                        `pulumi:"serviceName"`
}


type GatewayArgs struct {
	Description       pulumi.StringPtrInput
	GatewayId         pulumi.StringPtrInput
	LocationData      ResourceLocationDataContractPtrInput
	ResourceGroupName pulumi.StringInput
	ServiceName       pulumi.StringInput
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
	return reflect.TypeOf((*Gateway)(nil))
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Gateway)(nil))
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GatewayOutput{})
}
