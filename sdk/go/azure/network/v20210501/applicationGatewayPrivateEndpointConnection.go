


package v20210501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApplicationGatewayPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Etag                              pulumi.StringOutput                                `pulumi:"etag"`
	LinkIdentifier                    pulumi.StringOutput                                `pulumi:"linkIdentifier"`
	Name                              pulumi.StringPtrOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponseOutput                      `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                                `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                                `pulumi:"type"`
}


func NewApplicationGatewayPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *ApplicationGatewayPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*ApplicationGatewayPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApplicationGatewayName == nil {
		return nil, errors.New("invalid value for required argument 'ApplicationGatewayName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:network:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200501:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200601:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200701:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20200801:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20201101:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210201:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210301:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20210801:ApplicationGatewayPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:network/v20220101:ApplicationGatewayPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource ApplicationGatewayPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:network/v20210501:ApplicationGatewayPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetApplicationGatewayPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApplicationGatewayPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*ApplicationGatewayPrivateEndpointConnection, error) {
	var resource ApplicationGatewayPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:network/v20210501:ApplicationGatewayPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type applicationGatewayPrivateEndpointConnectionState struct {
}

type ApplicationGatewayPrivateEndpointConnectionState struct {
}

func (ApplicationGatewayPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayPrivateEndpointConnectionState)(nil)).Elem()
}

type applicationGatewayPrivateEndpointConnectionArgs struct {
	ApplicationGatewayName            string                             `pulumi:"applicationGatewayName"`
	ConnectionName                    *string                            `pulumi:"connectionName"`
	Id                                *string                            `pulumi:"id"`
	Name                              *string                            `pulumi:"name"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                             `pulumi:"resourceGroupName"`
}


type ApplicationGatewayPrivateEndpointConnectionArgs struct {
	ApplicationGatewayName            pulumi.StringInput
	ConnectionName                    pulumi.StringPtrInput
	Id                                pulumi.StringPtrInput
	Name                              pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput
	ResourceGroupName                 pulumi.StringInput
}

func (ApplicationGatewayPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*applicationGatewayPrivateEndpointConnectionArgs)(nil)).Elem()
}

type ApplicationGatewayPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToApplicationGatewayPrivateEndpointConnectionOutput() ApplicationGatewayPrivateEndpointConnectionOutput
	ToApplicationGatewayPrivateEndpointConnectionOutputWithContext(ctx context.Context) ApplicationGatewayPrivateEndpointConnectionOutput
}

func (*ApplicationGatewayPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayPrivateEndpointConnection)(nil)).Elem()
}

func (i *ApplicationGatewayPrivateEndpointConnection) ToApplicationGatewayPrivateEndpointConnectionOutput() ApplicationGatewayPrivateEndpointConnectionOutput {
	return i.ToApplicationGatewayPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *ApplicationGatewayPrivateEndpointConnection) ToApplicationGatewayPrivateEndpointConnectionOutputWithContext(ctx context.Context) ApplicationGatewayPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationGatewayPrivateEndpointConnectionOutput)
}

type ApplicationGatewayPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (ApplicationGatewayPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationGatewayPrivateEndpointConnection)(nil)).Elem()
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) ToApplicationGatewayPrivateEndpointConnectionOutput() ApplicationGatewayPrivateEndpointConnectionOutput {
	return o
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) ToApplicationGatewayPrivateEndpointConnectionOutputWithContext(ctx context.Context) ApplicationGatewayPrivateEndpointConnectionOutput {
	return o
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) LinkIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) pulumi.StringOutput { return v.LinkIdentifier }).(pulumi.StringOutput)
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) pulumi.StringPtrOutput { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) PrivateEndpointResponseOutput {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponseOutput)
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) PrivateLinkServiceConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ApplicationGatewayPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ApplicationGatewayPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ApplicationGatewayPrivateEndpointConnectionOutput{})
}
