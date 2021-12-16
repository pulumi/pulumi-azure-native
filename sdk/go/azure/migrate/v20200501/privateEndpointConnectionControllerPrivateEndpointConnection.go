


package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionControllerPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	ETag       pulumi.StringOutput                               `pulumi:"eTag"`
	Name       pulumi.StringOutput                               `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                          `pulumi:"systemData"`
	Type       pulumi.StringOutput                               `pulumi:"type"`
}


func NewPrivateEndpointConnectionControllerPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionControllerPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionControllerPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MigrateProjectName == nil {
		return nil, errors.New("invalid value for required argument 'MigrateProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource PrivateEndpointConnectionControllerPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:migrate/v20200501:PrivateEndpointConnectionControllerPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionControllerPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionControllerPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionControllerPrivateEndpointConnection, error) {
	var resource PrivateEndpointConnectionControllerPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:migrate/v20200501:PrivateEndpointConnectionControllerPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionControllerPrivateEndpointConnectionState struct {
}

type PrivateEndpointConnectionControllerPrivateEndpointConnectionState struct {
}

func (PrivateEndpointConnectionControllerPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionControllerPrivateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionControllerPrivateEndpointConnectionArgs struct {
	ETag               *string                               `pulumi:"eTag"`
	MigrateProjectName string                                `pulumi:"migrateProjectName"`
	PeConnectionName   *string                               `pulumi:"peConnectionName"`
	Properties         *ConnectionStateRequestBodyProperties `pulumi:"properties"`
	ResourceGroupName  string                                `pulumi:"resourceGroupName"`
}


type PrivateEndpointConnectionControllerPrivateEndpointConnectionArgs struct {
	ETag               pulumi.StringPtrInput
	MigrateProjectName pulumi.StringInput
	PeConnectionName   pulumi.StringPtrInput
	Properties         ConnectionStateRequestBodyPropertiesPtrInput
	ResourceGroupName  pulumi.StringInput
}

func (PrivateEndpointConnectionControllerPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionControllerPrivateEndpointConnectionArgs)(nil)).Elem()
}

type PrivateEndpointConnectionControllerPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionControllerPrivateEndpointConnectionOutput() PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput
	ToPrivateEndpointConnectionControllerPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput
}

func (*PrivateEndpointConnectionControllerPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionControllerPrivateEndpointConnection)(nil)).Elem()
}

func (i *PrivateEndpointConnectionControllerPrivateEndpointConnection) ToPrivateEndpointConnectionControllerPrivateEndpointConnectionOutput() PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionControllerPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionControllerPrivateEndpointConnection) ToPrivateEndpointConnectionControllerPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionControllerPrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput) ToPrivateEndpointConnectionControllerPrivateEndpointConnectionOutput() PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput) ToPrivateEndpointConnectionControllerPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionControllerPrivateEndpointConnectionOutput{})
}
