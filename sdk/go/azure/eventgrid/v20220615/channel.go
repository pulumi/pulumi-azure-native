


package v20220615

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Channel struct {
	pulumi.CustomResourceState

	ChannelType                     pulumi.StringPtrOutput            `pulumi:"channelType"`
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrOutput            `pulumi:"expirationTimeIfNotActivatedUtc"`
	MessageForActivation            pulumi.StringPtrOutput            `pulumi:"messageForActivation"`
	Name                            pulumi.StringOutput               `pulumi:"name"`
	PartnerTopicInfo                PartnerTopicInfoResponsePtrOutput `pulumi:"partnerTopicInfo"`
	ProvisioningState               pulumi.StringPtrOutput            `pulumi:"provisioningState"`
	ReadinessState                  pulumi.StringPtrOutput            `pulumi:"readinessState"`
	SystemData                      SystemDataResponseOutput          `pulumi:"systemData"`
	Type                            pulumi.StringOutput               `pulumi:"type"`
}


func NewChannel(ctx *pulumi.Context,
	name string, args *ChannelArgs, opts ...pulumi.ResourceOption) (*Channel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PartnerNamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'PartnerNamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:Channel"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:Channel"),
		},
	})
	opts = append(opts, aliases)
	var resource Channel
	err := ctx.RegisterResource("azure-native:eventgrid/v20220615:Channel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChannelState, opts ...pulumi.ResourceOption) (*Channel, error) {
	var resource Channel
	err := ctx.ReadResource("azure-native:eventgrid/v20220615:Channel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type channelState struct {
}

type ChannelState struct {
}

func (ChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*channelState)(nil)).Elem()
}

type channelArgs struct {
	ChannelName                     *string           `pulumi:"channelName"`
	ChannelType                     *string           `pulumi:"channelType"`
	ExpirationTimeIfNotActivatedUtc *string           `pulumi:"expirationTimeIfNotActivatedUtc"`
	MessageForActivation            *string           `pulumi:"messageForActivation"`
	PartnerNamespaceName            string            `pulumi:"partnerNamespaceName"`
	PartnerTopicInfo                *PartnerTopicInfo `pulumi:"partnerTopicInfo"`
	ProvisioningState               *string           `pulumi:"provisioningState"`
	ReadinessState                  *string           `pulumi:"readinessState"`
	ResourceGroupName               string            `pulumi:"resourceGroupName"`
}


type ChannelArgs struct {
	ChannelName                     pulumi.StringPtrInput
	ChannelType                     pulumi.StringPtrInput
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrInput
	MessageForActivation            pulumi.StringPtrInput
	PartnerNamespaceName            pulumi.StringInput
	PartnerTopicInfo                PartnerTopicInfoPtrInput
	ProvisioningState               pulumi.StringPtrInput
	ReadinessState                  pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
}

func (ChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*channelArgs)(nil)).Elem()
}

type ChannelInput interface {
	pulumi.Input

	ToChannelOutput() ChannelOutput
	ToChannelOutputWithContext(ctx context.Context) ChannelOutput
}

func (*Channel) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (i *Channel) ToChannelOutput() ChannelOutput {
	return i.ToChannelOutputWithContext(context.Background())
}

func (i *Channel) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChannelOutput)
}

type ChannelOutput struct{ *pulumi.OutputState }

func (ChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Channel)(nil)).Elem()
}

func (o ChannelOutput) ToChannelOutput() ChannelOutput {
	return o
}

func (o ChannelOutput) ToChannelOutputWithContext(ctx context.Context) ChannelOutput {
	return o
}

func (o ChannelOutput) ChannelType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.ChannelType }).(pulumi.StringPtrOutput)
}

func (o ChannelOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

func (o ChannelOutput) MessageForActivation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.MessageForActivation }).(pulumi.StringPtrOutput)
}

func (o ChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ChannelOutput) PartnerTopicInfo() PartnerTopicInfoResponsePtrOutput {
	return o.ApplyT(func(v *Channel) PartnerTopicInfoResponsePtrOutput { return v.PartnerTopicInfo }).(PartnerTopicInfoResponsePtrOutput)
}

func (o ChannelOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ChannelOutput) ReadinessState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringPtrOutput { return v.ReadinessState }).(pulumi.StringPtrOutput)
}

func (o ChannelOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Channel) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ChannelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Channel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ChannelOutput{})
}
