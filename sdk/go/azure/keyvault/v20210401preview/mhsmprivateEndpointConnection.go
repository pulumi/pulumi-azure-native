


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MHSMPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Etag                              pulumi.StringPtrOutput                                 `pulumi:"etag"`
	Location                          pulumi.StringPtrOutput                                 `pulumi:"location"`
	Name                              pulumi.StringOutput                                    `pulumi:"name"`
	PrivateEndpoint                   MHSMPrivateEndpointResponsePtrOutput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState MHSMPrivateLinkServiceConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                                    `pulumi:"provisioningState"`
	Sku                               ManagedHsmSkuResponsePtrOutput                         `pulumi:"sku"`
	SystemData                        SystemDataResponseOutput                               `pulumi:"systemData"`
	Tags                              pulumi.StringMapOutput                                 `pulumi:"tags"`
	Type                              pulumi.StringOutput                                    `pulumi:"type"`
}


func NewMHSMPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *MHSMPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*MHSMPrivateEndpointConnection, error) {
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
			Type: pulumi.String("azure-native:keyvault:MHSMPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:keyvault/v20210601preview:MHSMPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource MHSMPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:keyvault/v20210401preview:MHSMPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMHSMPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MHSMPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*MHSMPrivateEndpointConnection, error) {
	var resource MHSMPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:keyvault/v20210401preview:MHSMPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mhsmprivateEndpointConnectionState struct {
}

type MHSMPrivateEndpointConnectionState struct {
}

func (MHSMPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*mhsmprivateEndpointConnectionState)(nil)).Elem()
}

type mhsmprivateEndpointConnectionArgs struct {
	Etag                              *string                                `pulumi:"etag"`
	Location                          *string                                `pulumi:"location"`
	Name                              string                                 `pulumi:"name"`
	PrivateEndpointConnectionName     *string                                `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState *MHSMPrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                                 `pulumi:"resourceGroupName"`
	Sku                               *ManagedHsmSku                         `pulumi:"sku"`
	Tags                              map[string]string                      `pulumi:"tags"`
}


type MHSMPrivateEndpointConnectionArgs struct {
	Etag                              pulumi.StringPtrInput
	Location                          pulumi.StringPtrInput
	Name                              pulumi.StringInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState MHSMPrivateLinkServiceConnectionStatePtrInput
	ResourceGroupName                 pulumi.StringInput
	Sku                               ManagedHsmSkuPtrInput
	Tags                              pulumi.StringMapInput
}

func (MHSMPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mhsmprivateEndpointConnectionArgs)(nil)).Elem()
}

type MHSMPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToMHSMPrivateEndpointConnectionOutput() MHSMPrivateEndpointConnectionOutput
	ToMHSMPrivateEndpointConnectionOutputWithContext(ctx context.Context) MHSMPrivateEndpointConnectionOutput
}

func (*MHSMPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMPrivateEndpointConnection)(nil))
}

func (i *MHSMPrivateEndpointConnection) ToMHSMPrivateEndpointConnectionOutput() MHSMPrivateEndpointConnectionOutput {
	return i.ToMHSMPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *MHSMPrivateEndpointConnection) ToMHSMPrivateEndpointConnectionOutputWithContext(ctx context.Context) MHSMPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MHSMPrivateEndpointConnectionOutput)
}

type MHSMPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (MHSMPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MHSMPrivateEndpointConnection)(nil))
}

func (o MHSMPrivateEndpointConnectionOutput) ToMHSMPrivateEndpointConnectionOutput() MHSMPrivateEndpointConnectionOutput {
	return o
}

func (o MHSMPrivateEndpointConnectionOutput) ToMHSMPrivateEndpointConnectionOutputWithContext(ctx context.Context) MHSMPrivateEndpointConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MHSMPrivateEndpointConnectionOutput{})
}
