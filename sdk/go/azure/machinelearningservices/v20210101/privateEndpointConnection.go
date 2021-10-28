


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Identity                          IdentityResponsePtrOutput                       `pulumi:"identity"`
	Location                          pulumi.StringPtrOutput                          `pulumi:"location"`
	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	Sku                               SkuResponsePtrOutput                            `pulumi:"sku"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Tags                              pulumi.StringMapOutput                          `pulumi:"tags"`
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
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200101:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200218preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200218preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200501preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200515preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200601:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200801:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200901preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210301preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210701:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210701:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20210101:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-native:machinelearningservices/v20210101:PrivateEndpointConnection", name, id, state, &resource, opts...)
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
	Identity                          *Identity                         `pulumi:"identity"`
	Location                          *string                           `pulumi:"location"`
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
	Sku                               *Sku                              `pulumi:"sku"`
	Tags                              map[string]string                 `pulumi:"tags"`
	WorkspaceName                     string                            `pulumi:"workspaceName"`
}


type PrivateEndpointConnectionArgs struct {
	Identity                          IdentityPtrInput
	Location                          pulumi.StringPtrInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
	Sku                               SkuPtrInput
	Tags                              pulumi.StringMapInput
	WorkspaceName                     pulumi.StringInput
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
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionInput)(nil)).Elem(), &PrivateEndpointConnection{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}
