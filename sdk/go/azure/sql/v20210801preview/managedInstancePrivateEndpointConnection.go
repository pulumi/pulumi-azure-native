


package v20210801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagedInstancePrivateEndpointConnection struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                                                       `pulumi:"name"`
	PrivateEndpoint                   ManagedInstancePrivateEndpointPropertyResponsePtrOutput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                                                       `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                                                       `pulumi:"type"`
}


func NewManagedInstancePrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *ManagedInstancePrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*ManagedInstancePrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedInstanceName == nil {
		return nil, errors.New("invalid value for required argument 'ManagedInstanceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20211101preview:ManagedInstancePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20220201preview:ManagedInstancePrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagedInstancePrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:sql/v20210801preview:ManagedInstancePrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagedInstancePrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedInstancePrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*ManagedInstancePrivateEndpointConnection, error) {
	var resource ManagedInstancePrivateEndpointConnection
	err := ctx.ReadResource("azure-native:sql/v20210801preview:ManagedInstancePrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managedInstancePrivateEndpointConnectionState struct {
}

type ManagedInstancePrivateEndpointConnectionState struct {
}

func (ManagedInstancePrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstancePrivateEndpointConnectionState)(nil)).Elem()
}

type managedInstancePrivateEndpointConnectionArgs struct {
	ManagedInstanceName               string                                                    `pulumi:"managedInstanceName"`
	PrivateEndpoint                   *ManagedInstancePrivateEndpointProperty                   `pulumi:"privateEndpoint"`
	PrivateEndpointConnectionName     *string                                                   `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState *ManagedInstancePrivateLinkServiceConnectionStateProperty `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                                                    `pulumi:"resourceGroupName"`
}


type ManagedInstancePrivateEndpointConnectionArgs struct {
	ManagedInstanceName               pulumi.StringInput
	PrivateEndpoint                   ManagedInstancePrivateEndpointPropertyPtrInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState ManagedInstancePrivateLinkServiceConnectionStatePropertyPtrInput
	ResourceGroupName                 pulumi.StringInput
}

func (ManagedInstancePrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedInstancePrivateEndpointConnectionArgs)(nil)).Elem()
}

type ManagedInstancePrivateEndpointConnectionInput interface {
	pulumi.Input

	ToManagedInstancePrivateEndpointConnectionOutput() ManagedInstancePrivateEndpointConnectionOutput
	ToManagedInstancePrivateEndpointConnectionOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointConnectionOutput
}

func (*ManagedInstancePrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateEndpointConnection)(nil)).Elem()
}

func (i *ManagedInstancePrivateEndpointConnection) ToManagedInstancePrivateEndpointConnectionOutput() ManagedInstancePrivateEndpointConnectionOutput {
	return i.ToManagedInstancePrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *ManagedInstancePrivateEndpointConnection) ToManagedInstancePrivateEndpointConnectionOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedInstancePrivateEndpointConnectionOutput)
}

type ManagedInstancePrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (ManagedInstancePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedInstancePrivateEndpointConnection)(nil)).Elem()
}

func (o ManagedInstancePrivateEndpointConnectionOutput) ToManagedInstancePrivateEndpointConnectionOutput() ManagedInstancePrivateEndpointConnectionOutput {
	return o
}

func (o ManagedInstancePrivateEndpointConnectionOutput) ToManagedInstancePrivateEndpointConnectionOutputWithContext(ctx context.Context) ManagedInstancePrivateEndpointConnectionOutput {
	return o
}

func (o ManagedInstancePrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagedInstancePrivateEndpointConnectionOutput) PrivateEndpoint() ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointConnection) ManagedInstancePrivateEndpointPropertyResponsePtrOutput {
		return v.PrivateEndpoint
	}).(ManagedInstancePrivateEndpointPropertyResponsePtrOutput)
}

func (o ManagedInstancePrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointConnection) ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(ManagedInstancePrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o ManagedInstancePrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ManagedInstancePrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagedInstancePrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagedInstancePrivateEndpointConnectionOutput{})
}
