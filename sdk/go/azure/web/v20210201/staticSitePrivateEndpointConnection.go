


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StaticSitePrivateEndpointConnection struct {
	pulumi.CustomResourceState

	IpAddresses                       pulumi.StringArrayOutput                    `pulumi:"ipAddresses"`
	Kind                              pulumi.StringPtrOutput                      `pulumi:"kind"`
	Name                              pulumi.StringOutput                         `pulumi:"name"`
	PrivateEndpoint                   ArmIdWrapperResponsePtrOutput               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                         `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                         `pulumi:"type"`
}


func NewStaticSitePrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *StaticSitePrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*StaticSitePrivateEndpointConnection, error) {
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
			Type: pulumi.String("azure-native:web:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:StaticSitePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:StaticSitePrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource StaticSitePrivateEndpointConnection
	err := ctx.RegisterResource("azure-native:web/v20210201:StaticSitePrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStaticSitePrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StaticSitePrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*StaticSitePrivateEndpointConnection, error) {
	var resource StaticSitePrivateEndpointConnection
	err := ctx.ReadResource("azure-native:web/v20210201:StaticSitePrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type staticSitePrivateEndpointConnectionState struct {
}

type StaticSitePrivateEndpointConnectionState struct {
}

func (StaticSitePrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSitePrivateEndpointConnectionState)(nil)).Elem()
}

type staticSitePrivateEndpointConnectionArgs struct {
	Kind                              *string                     `pulumi:"kind"`
	Name                              string                      `pulumi:"name"`
	PrivateEndpointConnectionName     *string                     `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                      `pulumi:"resourceGroupName"`
}


type StaticSitePrivateEndpointConnectionArgs struct {
	Kind                              pulumi.StringPtrInput
	Name                              pulumi.StringInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkConnectionStatePtrInput
	ResourceGroupName                 pulumi.StringInput
}

func (StaticSitePrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*staticSitePrivateEndpointConnectionArgs)(nil)).Elem()
}

type StaticSitePrivateEndpointConnectionInput interface {
	pulumi.Input

	ToStaticSitePrivateEndpointConnectionOutput() StaticSitePrivateEndpointConnectionOutput
	ToStaticSitePrivateEndpointConnectionOutputWithContext(ctx context.Context) StaticSitePrivateEndpointConnectionOutput
}

func (*StaticSitePrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSitePrivateEndpointConnection)(nil)).Elem()
}

func (i *StaticSitePrivateEndpointConnection) ToStaticSitePrivateEndpointConnectionOutput() StaticSitePrivateEndpointConnectionOutput {
	return i.ToStaticSitePrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i *StaticSitePrivateEndpointConnection) ToStaticSitePrivateEndpointConnectionOutputWithContext(ctx context.Context) StaticSitePrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSitePrivateEndpointConnectionOutput)
}

type StaticSitePrivateEndpointConnectionOutput struct{ *pulumi.OutputState }

func (StaticSitePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSitePrivateEndpointConnection)(nil)).Elem()
}

func (o StaticSitePrivateEndpointConnectionOutput) ToStaticSitePrivateEndpointConnectionOutput() StaticSitePrivateEndpointConnectionOutput {
	return o
}

func (o StaticSitePrivateEndpointConnectionOutput) ToStaticSitePrivateEndpointConnectionOutputWithContext(ctx context.Context) StaticSitePrivateEndpointConnectionOutput {
	return o
}

func (o StaticSitePrivateEndpointConnectionOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o StaticSitePrivateEndpointConnectionOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o StaticSitePrivateEndpointConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o StaticSitePrivateEndpointConnectionOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) ArmIdWrapperResponsePtrOutput { return v.PrivateEndpoint }).(ArmIdWrapperResponsePtrOutput)
}

func (o StaticSitePrivateEndpointConnectionOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) PrivateLinkConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o StaticSitePrivateEndpointConnectionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o StaticSitePrivateEndpointConnectionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *StaticSitePrivateEndpointConnection) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(StaticSitePrivateEndpointConnectionOutput{})
}
