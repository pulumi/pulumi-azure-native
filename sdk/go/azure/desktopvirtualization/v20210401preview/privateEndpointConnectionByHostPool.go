


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionByHostPool struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseOutput                        `pulumi:"systemData"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnectionByHostPool(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionByHostPoolArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByHostPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostPoolName == nil {
		return nil, errors.New("invalid value for required argument 'HostPoolName'")
	}
	if args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("invalid value for required argument 'PrivateLinkServiceConnectionState'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:desktopvirtualization:PrivateEndpointConnectionByHostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20210903preview:PrivateEndpointConnectionByHostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220210preview:PrivateEndpointConnectionByHostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20220401preview:PrivateEndpointConnectionByHostPool"),
		},
		{
			Type: pulumi.String("azure-native:desktopvirtualization/v20221014preview:PrivateEndpointConnectionByHostPool"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionByHostPool
	err := ctx.RegisterResource("azure-native:desktopvirtualization/v20210401preview:PrivateEndpointConnectionByHostPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionByHostPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionByHostPoolState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByHostPool, error) {
	var resource PrivateEndpointConnectionByHostPool
	err := ctx.ReadResource("azure-native:desktopvirtualization/v20210401preview:PrivateEndpointConnectionByHostPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionByHostPoolState struct {
}

type PrivateEndpointConnectionByHostPoolState struct {
}

func (PrivateEndpointConnectionByHostPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByHostPoolState)(nil)).Elem()
}

type privateEndpointConnectionByHostPoolArgs struct {
	HostPoolName                      string                            `pulumi:"hostPoolName"`
	PrivateEndpointConnectionName     *string                           `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                            `pulumi:"resourceGroupName"`
}


type PrivateEndpointConnectionByHostPoolArgs struct {
	HostPoolName                      pulumi.StringInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	ResourceGroupName                 pulumi.StringInput
}

func (PrivateEndpointConnectionByHostPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByHostPoolArgs)(nil)).Elem()
}

type PrivateEndpointConnectionByHostPoolInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionByHostPoolOutput() PrivateEndpointConnectionByHostPoolOutput
	ToPrivateEndpointConnectionByHostPoolOutputWithContext(ctx context.Context) PrivateEndpointConnectionByHostPoolOutput
}

func (*PrivateEndpointConnectionByHostPool) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByHostPool)(nil)).Elem()
}

func (i *PrivateEndpointConnectionByHostPool) ToPrivateEndpointConnectionByHostPoolOutput() PrivateEndpointConnectionByHostPoolOutput {
	return i.ToPrivateEndpointConnectionByHostPoolOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionByHostPool) ToPrivateEndpointConnectionByHostPoolOutputWithContext(ctx context.Context) PrivateEndpointConnectionByHostPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionByHostPoolOutput)
}

type PrivateEndpointConnectionByHostPoolOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionByHostPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByHostPool)(nil)).Elem()
}

func (o PrivateEndpointConnectionByHostPoolOutput) ToPrivateEndpointConnectionByHostPoolOutput() PrivateEndpointConnectionByHostPoolOutput {
	return o
}

func (o PrivateEndpointConnectionByHostPoolOutput) ToPrivateEndpointConnectionByHostPoolOutputWithContext(ctx context.Context) PrivateEndpointConnectionByHostPoolOutput {
	return o
}

func (o PrivateEndpointConnectionByHostPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionByHostPoolOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) PrivateEndpointResponsePtrOutput {
		return v.PrivateEndpoint
	}).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionByHostPoolOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionByHostPoolOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionByHostPoolOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrivateEndpointConnectionByHostPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByHostPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionByHostPoolOutput{})
}
