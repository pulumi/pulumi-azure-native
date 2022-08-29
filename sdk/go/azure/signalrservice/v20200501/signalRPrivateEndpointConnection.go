


package v20200501

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SignalRPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                                `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                                `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                                `pulumi:"type"`
}


func NewSignalRPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *SignalRPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*SignalRPrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ResourceName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:signalrservice:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20200701preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210401preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210601preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20210901preview:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20211001:SignalRPrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:signalrservice/v20220201:SignalRPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource SignalRPrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:signalrservice/v20200501:SignalRPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSignalRPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*SignalRPrivateEndpointConnection, error) {
	var resource SignalRPrivateEndpointConnection
	err := ctx.ReadResource("azure-native:signalrservice/v20200501:SignalRPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type signalRPrivateEndpointConnectionState struct {
}

type SignalRPrivateEndpointConnectionState struct {
}

func (SignalRPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRPrivateEndpointConnectionState)(nil)).Elem()
}

type signalRPrivateEndpointConnectionArgs struct {
	PrivateEndpoint                   *PrivateEndpoint                   `pulumi:"privateEndpoint"`
	PrivateEndpointConnectionName     *string                            `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                             `pulumi:"resourceGroupName"`
	ResourceName                      string                             `pulumi:"resourceName"`
}


type SignalRPrivateEndpointConnectionArgs struct {
	PrivateEndpoint                   PrivateEndpointPtrInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput
	ResourceGroupName                 pulumi.StringInput
	ResourceName                      pulumi.StringInput
}

func (SignalRPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRPrivateEndpointConnectionArgs)(nil)).Elem()
}

type SignalRPrivateEndpointConnectionInput interface {
	pulumi.Input

	ToSignalRPrivateEndpointConnectionOutput() SignalRPrivateEndpointConnectionOutput
	ToSignalRPrivateEndpointConnectionOutputWithContext(ctx context.Context) SignalRPrivateEndpointConnectionOutput
}

func (*SignalRPrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRPrivateEndpointConnection)(nil)).Elem()
}

func (i *SignalRPrivateEndpointConnection) ToSignalRPrivateEndpointConnectionOutput() SignalRPrivateEndpointConnectionOutput {
	return i.ToSignalRPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *SignalRPrivateEndpointConnection) ToSignalRPrivateEndpointConnectionOutputWithContext(ctx context.Context) SignalRPrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SignalRPrivateEndpointConnectionOutput)
}

type SignalRPrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (SignalRPrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRPrivateEndpointConnection)(nil)).Elem()
}

func (o SignalRPrivateEndpointConnectionOutput) ToSignalRPrivateEndpointConnectionOutput() SignalRPrivateEndpointConnectionOutput {
	return o
}

func (o SignalRPrivateEndpointConnectionOutput) ToSignalRPrivateEndpointConnectionOutputWithContext(ctx context.Context) SignalRPrivateEndpointConnectionOutput {
	return o
}

func (o SignalRPrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SignalRPrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o SignalRPrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) PrivateLinkServiceConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o SignalRPrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SignalRPrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SignalRPrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SignalRPrivateEndpointConnectionOutput{})
}
