


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppServiceEnvironmentPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	IpAddresses                       pulumi.StringArrayOutput                    `pulumi:"ipAddresses"`
	Kind                              pulumi.StringPtrOutput                      `pulumi:"kind"`
	Name                              pulumi.StringOutput                         `pulumi:"name"`
	PrivateEndpoint                   ArmIdWrapperResponsePtrOutput               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                         `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                         `pulumi:"type"`
}


func NewAppServiceEnvironmentPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *AppServiceEnvironmentPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*AppServiceEnvironmentPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:web/v20210101:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20201201:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210115:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:AppServiceEnvironmentPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:web/v20210201:AppServiceEnvironmentPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource AppServiceEnvironmentPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:web/v20210101:AppServiceEnvironmentPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAppServiceEnvironmentPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceEnvironmentPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*AppServiceEnvironmentPrivateEndpointConnection, error) {
	var resource AppServiceEnvironmentPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:web/v20210101:AppServiceEnvironmentPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type appServiceEnvironmentPrivateEndpointConnectionState struct {
}

type AppServiceEnvironmentPrivateEndpointConnectionState struct {
}

func (AppServiceEnvironmentPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentPrivateEndpointConnectionState)(nil)).Elem()
}

type appServiceEnvironmentPrivateEndpointConnectionArgs struct {
	Kind                              *string                     `pulumi:"kind"`
	Name                              string                      `pulumi:"name"`
	PrivateEndpointConnectionName     *string                     `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                      `pulumi:"resourceGroupName"`
}


type AppServiceEnvironmentPrivateEndpointConnectionArgs struct {
	Kind                              pulumi.StringPtrInput
	Name                              pulumi.StringInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkConnectionStatePtrInput
	ResourceGroupName                 pulumi.StringInput
}

func (AppServiceEnvironmentPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceEnvironmentPrivateEndpointConnectionArgs)(nil)).Elem()
}

type AppServiceEnvironmentPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToAppServiceEnvironmentPrivateEndpointConnectionOutput() AppServiceEnvironmentPrivateEndpointConnectionOutput
	ToAppServiceEnvironmentPrivateEndpointConnectionOutputWithContext(ctx context.Context) AppServiceEnvironmentPrivateEndpointConnectionOutput
}

func (*AppServiceEnvironmentPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceEnvironmentPrivateEndpointConnection)(nil))
}

func (i *AppServiceEnvironmentPrivateEndpointConnection) ToAppServiceEnvironmentPrivateEndpointConnectionOutput() AppServiceEnvironmentPrivateEndpointConnectionOutput {
	return i.ToAppServiceEnvironmentPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *AppServiceEnvironmentPrivateEndpointConnection) ToAppServiceEnvironmentPrivateEndpointConnectionOutputWithContext(ctx context.Context) AppServiceEnvironmentPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppServiceEnvironmentPrivateEndpointConnectionOutput)
}

type AppServiceEnvironmentPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (AppServiceEnvironmentPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppServiceEnvironmentPrivateEndpointConnection)(nil))
}

func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) ToAppServiceEnvironmentPrivateEndpointConnectionOutput() AppServiceEnvironmentPrivateEndpointConnectionOutput {
	return o
}

func (o AppServiceEnvironmentPrivateEndpointConnectionOutput) ToAppServiceEnvironmentPrivateEndpointConnectionOutputWithContext(ctx context.Context) AppServiceEnvironmentPrivateEndpointConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppServiceEnvironmentPrivateEndpointConnectionInput)(nil)).Elem(), &AppServiceEnvironmentPrivateEndpointConnection{})
	pulumi.RegisterOutputType(AppServiceEnvironmentPrivateEndpointConnectionOutput{})
}
