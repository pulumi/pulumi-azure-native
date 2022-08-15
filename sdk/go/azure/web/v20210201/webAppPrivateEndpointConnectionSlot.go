


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppPrivateEndpointConnectionSlot struct {
	pulumi.CustomResourceState

	IpAddresses                       pulumi.StringArrayOutput                    `pulumi:"ipAddresses"`
	Kind                              pulumi.StringPtrOutput                      `pulumi:"kind"`
	Name                              pulumi.StringOutput                         `pulumi:"name"`
	PrivateEndpoint                   ArmIdWrapperResponsePtrOutput               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringOutput                         `pulumi:"provisioningState"`
	Type                              pulumi.StringOutput                         `pulumi:"type"`
}


func NewWebAppPrivateEndpointConnectionSlot(ctx *pulumi.Context,
	name string, args *WebAppPrivateEndpointConnectionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppPrivateEndpointConnectionSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppPrivateEndpointConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppPrivateEndpointConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppPrivateEndpointConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppPrivateEndpointConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppPrivateEndpointConnectionSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppPrivateEndpointConnectionSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppPrivateEndpointConnectionSlot
	err := ctx.RegisterResource("azure-native:web/v20210201:WebAppPrivateEndpointConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppPrivateEndpointConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppPrivateEndpointConnectionSlotState, opts ...pulumi.ResourceOption) (*WebAppPrivateEndpointConnectionSlot, error) {
	var resource WebAppPrivateEndpointConnectionSlot
	err := ctx.ReadResource("azure-native:web/v20210201:WebAppPrivateEndpointConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppPrivateEndpointConnectionSlotState struct {
}

type WebAppPrivateEndpointConnectionSlotState struct {
}

func (WebAppPrivateEndpointConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPrivateEndpointConnectionSlotState)(nil)).Elem()
}

type webAppPrivateEndpointConnectionSlotArgs struct {
	Kind                              *string                     `pulumi:"kind"`
	Name                              string                      `pulumi:"name"`
	PrivateEndpointConnectionName     *string                     `pulumi:"privateEndpointConnectionName"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionState `pulumi:"privateLinkServiceConnectionState"`
	ResourceGroupName                 string                      `pulumi:"resourceGroupName"`
	Slot                              string                      `pulumi:"slot"`
}


type WebAppPrivateEndpointConnectionSlotArgs struct {
	Kind                              pulumi.StringPtrInput
	Name                              pulumi.StringInput
	PrivateEndpointConnectionName     pulumi.StringPtrInput
	PrivateLinkServiceConnectionState PrivateLinkConnectionStatePtrInput
	ResourceGroupName                 pulumi.StringInput
	Slot                              pulumi.StringInput
}

func (WebAppPrivateEndpointConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppPrivateEndpointConnectionSlotArgs)(nil)).Elem()
}

type WebAppPrivateEndpointConnectionSlotInput interface {
	pulumi.Input

	ToWebAppPrivateEndpointConnectionSlotOutput() WebAppPrivateEndpointConnectionSlotOutput
	ToWebAppPrivateEndpointConnectionSlotOutputWithContext(ctx context.Context) WebAppPrivateEndpointConnectionSlotOutput
}

func (*WebAppPrivateEndpointConnectionSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPrivateEndpointConnectionSlot)(nil)).Elem()
}

func (i *WebAppPrivateEndpointConnectionSlot) ToWebAppPrivateEndpointConnectionSlotOutput() WebAppPrivateEndpointConnectionSlotOutput {
	return i.ToWebAppPrivateEndpointConnectionSlotOutputWithContext(context.Background())
}

func (i *WebAppPrivateEndpointConnectionSlot) ToWebAppPrivateEndpointConnectionSlotOutputWithContext(ctx context.Context) WebAppPrivateEndpointConnectionSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppPrivateEndpointConnectionSlotOutput)
}

type WebAppPrivateEndpointConnectionSlotOutput struct{ *pulumi.OutputState }

func (WebAppPrivateEndpointConnectionSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppPrivateEndpointConnectionSlot)(nil)).Elem()
}

func (o WebAppPrivateEndpointConnectionSlotOutput) ToWebAppPrivateEndpointConnectionSlotOutput() WebAppPrivateEndpointConnectionSlotOutput {
	return o
}

func (o WebAppPrivateEndpointConnectionSlotOutput) ToWebAppPrivateEndpointConnectionSlotOutputWithContext(ctx context.Context) WebAppPrivateEndpointConnectionSlotOutput {
	return o
}

func (o WebAppPrivateEndpointConnectionSlotOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnectionSlot) pulumi.StringArrayOutput { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o WebAppPrivateEndpointConnectionSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnectionSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppPrivateEndpointConnectionSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnectionSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppPrivateEndpointConnectionSlotOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnectionSlot) ArmIdWrapperResponsePtrOutput { return v.PrivateEndpoint }).(ArmIdWrapperResponsePtrOutput)
}

func (o WebAppPrivateEndpointConnectionSlotOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnectionSlot) PrivateLinkConnectionStateResponsePtrOutput {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o WebAppPrivateEndpointConnectionSlotOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnectionSlot) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WebAppPrivateEndpointConnectionSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppPrivateEndpointConnectionSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppPrivateEndpointConnectionSlotOutput{})
}
