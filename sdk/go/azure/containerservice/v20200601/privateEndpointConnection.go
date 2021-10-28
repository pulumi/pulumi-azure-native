


package v20200601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200601:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200701:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200701:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20200901:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20200901:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20201101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20201201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20201201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210201:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210501:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210501:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210701:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210701:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210801:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210801:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:containerservice/v20210901:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerservice/v20210901:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:containerservice/v20200601:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-native:containerservice/v20200601:PrivateEndpointConnection", name, id, state, &resource, opts...)
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
	PrivateEndpoint                   *PrivateEndpoint                  `pulumi:"privateEndpoint"`
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	ResourceName                      string                            `pulumi:"resourceName"`
}


type PrivateEndpointConnectionArgs struct {
	PrivateEndpoint                   PrivateEndpointPtrInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	ResourceName                      pulumi.StringInput
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
