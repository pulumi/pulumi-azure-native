


package v20221031

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties ConnectionPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput           `pulumi:"systemData"`
	Type       pulumi.StringOutput                `pulumi:"type"`
}


func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
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
			Type: pulumi.String("azure-native:digitaltwins:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20201201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20210630preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:digitaltwins/v20220531:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:digitaltwins/v20221031:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-native:digitaltwins/v20221031:PrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionState struct {
}

type PrivateEndpointConnectionState struct {
}

func (PrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName *string              `pulumi:"privateEndpointConnectionName"`
	Properties                    ConnectionProperties `pulumi:"properties"`
	ResourceGroupName             string               `pulumi:"resourceGroupName"`
	ResourceName                  string               `pulumi:"resourceName"`
}


type PrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName pulumi.StringPtrInput
	Properties                    ConnectionPropertiesInput
	ResourceGroupName             pulumi.StringInput
	ResourceName                  pulumi.StringInput
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionArgs)(nil)).Elem()
}

type PrivateEndpointConnectionInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput
	ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput
}

func (*PrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionOutput) Properties() ConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) ConnectionPropertiesResponseOutput { return v.Properties }).(ConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}
