


package v20220310

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                               `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                          `pulumi:"systemData"`
	Type       pulumi.StringOutput                               `pulumi:"type"`
}


func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScopeName == nil {
		return nil, errors.New("invalid value for required argument 'ScopeName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:hybridcompute:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20200815preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210128preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210325preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210422preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210517preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210520:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20210610preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20211210preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220510preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:hybridcompute/v20220811preview:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:hybridcompute/v20220310:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-native:hybridcompute/v20220310:PrivateEndpointConnection", name, id, state, &resource, opts...)
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
	PrivateEndpointConnectionName *string                              `pulumi:"privateEndpointConnectionName"`
	Properties                    *PrivateEndpointConnectionProperties `pulumi:"properties"`
	ResourceGroupName             string                               `pulumi:"resourceGroupName"`
	ScopeName                     string                               `pulumi:"scopeName"`
}


type PrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName pulumi.StringPtrInput
	Properties                    PrivateEndpointConnectionPropertiesPtrInput
	ResourceGroupName             pulumi.StringInput
	ScopeName                     pulumi.StringInput
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

func (o PrivateEndpointConnectionOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateEndpointConnectionPropertiesResponseOutput {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
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
