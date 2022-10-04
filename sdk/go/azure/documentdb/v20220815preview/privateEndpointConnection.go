


package v20220815preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	GroupId                           pulumi.StringPtrOutput                                     `pulumi:"groupId"`
	Name                              pulumi.StringOutput                                        `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointPropertyResponsePtrOutput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringPtrOutput                                     `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                                        `pulumi:"type"`
}


func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:documentdb/v20220815preview:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azure-native:documentdb/v20220815preview:PrivateEndpointConnection", name, id, state, &resource, opts...)
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
	AccountName                       string                                     `pulumi:"accountName"`
	GroupId                           *string                                    `pulumi:"groupId"`
	PrivateEndpoint                   *PrivateEndpointProperty                   `pulumi:"privateEndpoint"`
	PrivateEndpointConnectionName     *string                                    `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateProperty `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string                                    `pulumi:"provisioningState"`
	ResourceGroupName                 string                                     `pulumi:"resourceGroupName"`
}


type PrivateEndpointConnectionArgs struct {
	AccountName                       pulumi.StringInput
	GroupId                           pulumi.StringPtrInput
	PrivateEndpoint                   PrivateEndpointPropertyPtrInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyPtrInput
	ProvisioningState                 pulumi.StringPtrInput
	ResourceGroupName                 pulumi.StringInput
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
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return i.ToPrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *PrivateEndpointConnection) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionOutput)
}

type PrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnection)(nil)).Elem()
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutput() PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) ToPrivateEndpointConnectionOutputWithContext(ctx context.Context) PrivateEndpointConnectionOutput {
	return o
}

func (o PrivateEndpointConnectionOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringPtrOutput { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateEndpointPropertyResponsePtrOutput { return v.PrivateEndpoint }).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrivateEndpointConnectionOutput{})
}
