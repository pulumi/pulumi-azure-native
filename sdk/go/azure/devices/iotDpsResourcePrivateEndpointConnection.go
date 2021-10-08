


package devices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type IotDpsResourcePrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                               `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                               `pulumi:"type"`
}


func NewIotDpsResourcePrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *IotDpsResourcePrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*IotDpsResourcePrivateEndpointConnection, error) {
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
			Type: pulumi.String("azure-nextgen:devices:IotDpsResourcePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200301:IotDpsResourcePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200301:IotDpsResourcePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:devices/v20200901preview:IotDpsResourcePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200901preview:IotDpsResourcePrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource IotDpsResourcePrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:devices:IotDpsResourcePrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetIotDpsResourcePrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotDpsResourcePrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*IotDpsResourcePrivateEndpointConnection, error) {
	var resource IotDpsResourcePrivateEndpointConnection
	err := ctx.ReadResource("azure-native:devices:IotDpsResourcePrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type iotDpsResourcePrivateEndpointConnectionState struct {
}

type IotDpsResourcePrivateEndpointConnectionState struct {
}

func (IotDpsResourcePrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDpsResourcePrivateEndpointConnectionState)(nil)).Elem()
}

type iotDpsResourcePrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName *string                             `pulumi:"privateEndpointConnectionName"`
	Properties                    PrivateEndpointConnectionProperties `pulumi:"properties"`
	ResourceGroupName             string                              `pulumi:"resourceGroupName"`
	ResourceName                  string                              `pulumi:"resourceName"`
}


type IotDpsResourcePrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName pulumi.StringPtrInput
	Properties                    PrivateEndpointConnectionPropertiesInput
	ResourceGroupName             pulumi.StringInput
	ResourceName                  pulumi.StringInput
}

func (IotDpsResourcePrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDpsResourcePrivateEndpointConnectionArgs)(nil)).Elem()
}

type IotDpsResourcePrivateEndpointConnectionInput interface {
	pulumi.Input

	ToIotDpsResourcePrivateEndpointConnectionOutput() IotDpsResourcePrivateEndpointConnectionOutput
	ToIotDpsResourcePrivateEndpointConnectionOutputWithContext(ctx context.Context) IotDpsResourcePrivateEndpointConnectionOutput
}

func (*IotDpsResourcePrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsResourcePrivateEndpointConnection)(nil))
}

func (i *IotDpsResourcePrivateEndpointConnection) ToIotDpsResourcePrivateEndpointConnectionOutput() IotDpsResourcePrivateEndpointConnectionOutput {
	return i.ToIotDpsResourcePrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *IotDpsResourcePrivateEndpointConnection) ToIotDpsResourcePrivateEndpointConnectionOutputWithContext(ctx context.Context) IotDpsResourcePrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsResourcePrivateEndpointConnectionOutput)
}

type IotDpsResourcePrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (IotDpsResourcePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsResourcePrivateEndpointConnection)(nil))
}

func (o IotDpsResourcePrivateEndpointConnectionOutput) ToIotDpsResourcePrivateEndpointConnectionOutput() IotDpsResourcePrivateEndpointConnectionOutput {
	return o
}

func (o IotDpsResourcePrivateEndpointConnectionOutput) ToIotDpsResourcePrivateEndpointConnectionOutputWithContext(ctx context.Context) IotDpsResourcePrivateEndpointConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IotDpsResourcePrivateEndpointConnectionOutput{})
}
