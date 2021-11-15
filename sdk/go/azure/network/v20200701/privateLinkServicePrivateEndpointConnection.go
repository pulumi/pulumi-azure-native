


package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateLinkServicePrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Etag                              pulumi.StringOutput                                `pulumi:"etag"`
	LinkIdentifier                    pulumi.StringOutput                                `pulumi:"linkIdentifier"`
	Name                              pulumi.StringPtrOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponseOutput                      `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                                `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                                `pulumi:"type"`
}


func NewPrivateLinkServicePrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateLinkServicePrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateLinkServicePrivateEndpointConnection, error) {
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
			Type: pulumi.String("azure-native:network:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20190901:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191101:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20191201:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200301:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200401:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:PrivateLinkServicePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210501:PrivateLinkServicePrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateLinkServicePrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:network/v20200701:PrivateLinkServicePrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateLinkServicePrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateLinkServicePrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateLinkServicePrivateEndpointConnection, error) {
	var resource PrivateLinkServicePrivateEndpointConnection
	err := ctx.ReadResource("azure-native:network/v20200701:PrivateLinkServicePrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateLinkServicePrivateEndpointConnectionState struct {
}

type PrivateLinkServicePrivateEndpointConnectionState struct {
}

func (PrivateLinkServicePrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicePrivateEndpointConnectionState)(nil)).Elem()
}

type privateLinkServicePrivateEndpointConnectionArgs struct {
	Id                                *string                            `pulumi:"id"`
	Name                              *string                            `pulumi:"name"`
	PeConnectionName                  *string                            `pulumi:"peConnectionName"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                             `pulumi:"resourceGroupName"`
	ServiceName                       string                             `pulumi:"serviceName"`
}


type PrivateLinkServicePrivateEndpointConnectionArgs struct {
	Id                                pulumi.StringPtrInput
	Name                              pulumi.StringPtrInput
	PeConnectionName                  pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput
	ResourceGroupName                 pulumi.StringInput
	ServiceName                       pulumi.StringInput
}

func (PrivateLinkServicePrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateLinkServicePrivateEndpointConnectionArgs)(nil)).Elem()
}

type PrivateLinkServicePrivateEndpointConnectionInput interface {
	pulumi.Input

	ToPrivateLinkServicePrivateEndpointConnectionOutput() PrivateLinkServicePrivateEndpointConnectionOutput
	ToPrivateLinkServicePrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateLinkServicePrivateEndpointConnectionOutput
}

func (*PrivateLinkServicePrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServicePrivateEndpointConnection)(nil))
}

func (i *PrivateLinkServicePrivateEndpointConnection) ToPrivateLinkServicePrivateEndpointConnectionOutput() PrivateLinkServicePrivateEndpointConnectionOutput {
	return i.ToPrivateLinkServicePrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *PrivateLinkServicePrivateEndpointConnection) ToPrivateLinkServicePrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateLinkServicePrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServicePrivateEndpointConnectionOutput)
}

type PrivateLinkServicePrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateLinkServicePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServicePrivateEndpointConnection)(nil))
}

func (o PrivateLinkServicePrivateEndpointConnectionOutput) ToPrivateLinkServicePrivateEndpointConnectionOutput() PrivateLinkServicePrivateEndpointConnectionOutput {
	return o
}

func (o PrivateLinkServicePrivateEndpointConnectionOutput) ToPrivateLinkServicePrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateLinkServicePrivateEndpointConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PrivateLinkServicePrivateEndpointConnectionOutput{})
}
