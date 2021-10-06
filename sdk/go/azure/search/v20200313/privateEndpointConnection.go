


package v20200313

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
	if args.SearchServiceName == nil {
		return nil, errors.New("invalid value for required argument 'SearchServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:search/v20200313:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:search:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20191001preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:search/v20191001preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200801:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:search/v20200801:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20200801preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:search/v20200801preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:search/v20210401preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:search/v20210401preview:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:search/v20200313:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-native:search/v20200313:PrivateEndpointConnection", name, id, state, &resource, opts...)
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
	SearchServiceName             string                               `pulumi:"searchServiceName"`
}


type PrivateEndpointConnectionArgs struct {
	PrivateEndpointConnectionName pulumi.StringPtrInput
	Properties                    PrivateEndpointConnectionPropertiesPtrInput
	ResourceGroupName             pulumi.StringInput
	SearchServiceName             pulumi.StringInput
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
	return reflect.TypeOf((*PrivateEndpointConnection)(nil))
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnection)(nil))
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}
