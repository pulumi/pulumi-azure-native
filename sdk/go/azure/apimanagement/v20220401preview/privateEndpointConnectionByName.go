


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnectionByName struct {
	pulumi.CustomResourceState

	Name                              pulumi.StringOutput                             `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrOutput                `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponseOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                             `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                             `pulumi:"type"`
}


func NewPrivateEndpointConnectionByName(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionByNameArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByName, error) {
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
			Type: pulumi.String("azure-native:apimanagement:PrivateEndpointConnectionByName"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210401preview:PrivateEndpointConnectionByName"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20210801:PrivateEndpointConnectionByName"),
		},
		{
			Type: pulumi.String("azure-native:apimanagement/v20211201preview:PrivateEndpointConnectionByName"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnectionByName
	err := ctx.RegisterResource("azure-native:apimanagement/v20220401preview:PrivateEndpointConnectionByName", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnectionByName(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionByNameState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnectionByName, error) {
	var resource PrivateEndpointConnectionByName
	err := ctx.ReadResource("azure-native:apimanagement/v20220401preview:PrivateEndpointConnectionByName", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type privateEndpointConnectionByNameState struct {
}

type PrivateEndpointConnectionByNameState struct {
}

func (PrivateEndpointConnectionByNameState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByNameState)(nil)).Elem()
}

type privateEndpointConnectionByNameArgs struct {
	Id                            *string                                     `pulumi:"id"`
	PrivateEndpointConnectionName *string                                     `pulumi:"privateEndpointConnectionName"`
	Properties                    *PrivateEndpointConnectionRequestProperties `pulumi:"properties"`
	ResourceGroupName             string                                      `pulumi:"resourceGroupName"`
	ServiceName                   string                                      `pulumi:"serviceName"`
}


type PrivateEndpointConnectionByNameArgs struct {
	Id                            pulumi.StringPtrInput
	PrivateEndpointConnectionName pulumi.StringPtrInput
	Properties                    PrivateEndpointConnectionRequestPropertiesPtrInput
	ResourceGroupName             pulumi.StringInput
	ServiceName                   pulumi.StringInput
}

func (PrivateEndpointConnectionByNameArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionByNameArgs)(nil)).Elem()
}

type PrivateEndpointConnectionByNameInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionByNameOutput() PrivateEndpointConnectionByNameOutput
	ToPrivateEndpointConnectionByNameOutputWithContext(ctx context.Context) PrivateEndpointConnectionByNameOutput
}

func (*PrivateEndpointConnectionByName) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByName)(nil)).Elem()
}

func (i *PrivateEndpointConnectionByName) ToPrivateEndpointConnectionByNameOutput() PrivateEndpointConnectionByNameOutput {
	return i.ToPrivateEndpointConnectionByNameOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnectionByName) ToPrivateEndpointConnectionByNameOutputWithContext(ctx context.Context) PrivateEndpointConnectionByNameOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionByNameOutput)
}

type PrivateEndpointConnectionByNameOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionByNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionByName)(nil)).Elem()
}

func (o PrivateEndpointConnectionByNameOutput) ToPrivateEndpointConnectionByNameOutput() PrivateEndpointConnectionByNameOutput {
	return o
}

func (o PrivateEndpointConnectionByNameOutput) ToPrivateEndpointConnectionByNameOutputWithContext(ctx context.Context) PrivateEndpointConnectionByNameOutput {
	return o
}

func (o PrivateEndpointConnectionByNameOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByName) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionByNameOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByName) PrivateEndpointResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionByNameOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByName) PrivateLinkServiceConnectionStateResponseOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateEndpointConnectionByNameOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByName) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionByNameOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionByName) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionByNameOutput{})
}
