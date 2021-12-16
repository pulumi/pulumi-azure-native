


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DigitalTwinsEndpoint struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput `pulumi:"name"`
	Properties pulumi.AnyOutput    `pulumi:"properties"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewDigitalTwinsEndpoint(ctx *pulumi.Context,
	name string, args *DigitalTwinsEndpointArgs, opts ...pulumi.ResourceOption) (*DigitalTwinsEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:digitaltwins:DigitalTwinsEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20200301preview:DigitalTwinsEndpoint"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20201031:DigitalTwinsEndpoint"),
		},
	})
	opts = append(opts, aliases)
	var resource DigitalTwinsEndpoint
	err := ctx.RegisterResource("azure-native:digitaltwins/v20201201:DigitalTwinsEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDigitalTwinsEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DigitalTwinsEndpointState, opts ...pulumi.ResourceOption) (*DigitalTwinsEndpoint, error) {
	var resource DigitalTwinsEndpoint
	err := ctx.ReadResource("azure-native:digitaltwins/v20201201:DigitalTwinsEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type digitalTwinsEndpointState struct {
}

type DigitalTwinsEndpointState struct {
}

func (DigitalTwinsEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*digitalTwinsEndpointState)(nil)).Elem()
}

type digitalTwinsEndpointArgs struct {
	EndpointName      *string     `pulumi:"endpointName"`
	Properties        interface{} `pulumi:"properties"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	ResourceName      string      `pulumi:"resourceName"`
}


type DigitalTwinsEndpointArgs struct {
	EndpointName      pulumi.StringPtrInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringInput
}

func (DigitalTwinsEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*digitalTwinsEndpointArgs)(nil)).Elem()
}

type DigitalTwinsEndpointInput interface {
	pulumi.Input

	ToDigitalTwinsEndpointOutput() DigitalTwinsEndpointOutput
	ToDigitalTwinsEndpointOutputWithContext(ctx context.Context) DigitalTwinsEndpointOutput
}

func (*DigitalTwinsEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsEndpoint)(nil)).Elem()
}

func (i *DigitalTwinsEndpoint) ToDigitalTwinsEndpointOutput() DigitalTwinsEndpointOutput {
	return i.ToDigitalTwinsEndpointOutputWithContext(context.Background())
}

func (i *DigitalTwinsEndpoint) ToDigitalTwinsEndpointOutputWithContext(ctx context.Context) DigitalTwinsEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DigitalTwinsEndpointOutput)
}

type DigitalTwinsEndpointOutput struct{ *pulumi.OutputState }

func (DigitalTwinsEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DigitalTwinsEndpoint)(nil)).Elem()
}

func (o DigitalTwinsEndpointOutput) ToDigitalTwinsEndpointOutput() DigitalTwinsEndpointOutput {
	return o
}

func (o DigitalTwinsEndpointOutput) ToDigitalTwinsEndpointOutputWithContext(ctx context.Context) DigitalTwinsEndpointOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DigitalTwinsEndpointOutput{})
}
