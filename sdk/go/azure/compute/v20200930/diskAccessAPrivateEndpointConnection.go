


package v20200930

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DiskAccessAPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewDiskAccessAPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *DiskAccessAPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*DiskAccessAPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DiskAccessName == nil {
		return nil, errors.New("invalid value for required argument 'DiskAccessName'")
	}
	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/v20200930:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20201201:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:DiskAccessAPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20210401:DiskAccessAPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource DiskAccessAPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:compute/v20200930:DiskAccessAPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDiskAccessAPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskAccessAPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*DiskAccessAPrivateEndpointConnection, error) {
	var resource DiskAccessAPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:compute/v20200930:DiskAccessAPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type diskAccessAPrivateEndpointConnectionState struct {
}

type DiskAccessAPrivateEndpointConnectionState struct {
}

func (DiskAccessAPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessAPrivateEndpointConnectionState)(nil)).Elem()
}

type diskAccessAPrivateEndpointConnectionArgs struct {
	DiskAccessName                    string                            `pulumi:"diskAccessName"`
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
}


type DiskAccessAPrivateEndpointConnectionArgs struct {
	DiskAccessName                    pulumi.StringInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
}

func (DiskAccessAPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessAPrivateEndpointConnectionArgs)(nil)).Elem()
}

type DiskAccessAPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToDiskAccessAPrivateEndpointConnectionOutput() DiskAccessAPrivateEndpointConnectionOutput
	ToDiskAccessAPrivateEndpointConnectionOutputWithContext(ctx context.Context) DiskAccessAPrivateEndpointConnectionOutput
}

func (*DiskAccessAPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskAccessAPrivateEndpointConnection)(nil))
}

func (i *DiskAccessAPrivateEndpointConnection) ToDiskAccessAPrivateEndpointConnectionOutput() DiskAccessAPrivateEndpointConnectionOutput {
	return i.ToDiskAccessAPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *DiskAccessAPrivateEndpointConnection) ToDiskAccessAPrivateEndpointConnectionOutputWithContext(ctx context.Context) DiskAccessAPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiskAccessAPrivateEndpointConnectionOutput)
}

type DiskAccessAPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (DiskAccessAPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiskAccessAPrivateEndpointConnection)(nil))
}

func (o DiskAccessAPrivateEndpointConnectionOutput) ToDiskAccessAPrivateEndpointConnectionOutput() DiskAccessAPrivateEndpointConnectionOutput {
	return o
}

func (o DiskAccessAPrivateEndpointConnectionOutput) ToDiskAccessAPrivateEndpointConnectionOutputWithContext(ctx context.Context) DiskAccessAPrivateEndpointConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DiskAccessAPrivateEndpointConnectionInput)(nil)).Elem(), &DiskAccessAPrivateEndpointConnection{})
	pulumi.RegisterOutputType(DiskAccessAPrivateEndpointConnectionOutput{})
}
