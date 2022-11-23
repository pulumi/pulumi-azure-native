


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PartnerTopic struct {
	pulumi.CustomResourceState

	ActivationState                 pulumi.StringPtrOutput         `pulumi:"activationState"`
	EventTypeInfo                   EventTypeInfoResponsePtrOutput `pulumi:"eventTypeInfo"`
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrOutput         `pulumi:"expirationTimeIfNotActivatedUtc"`
	Identity                        IdentityInfoResponsePtrOutput  `pulumi:"identity"`
	Location                        pulumi.StringOutput            `pulumi:"location"`
	MessageForActivation            pulumi.StringPtrOutput         `pulumi:"messageForActivation"`
	Name                            pulumi.StringOutput            `pulumi:"name"`
	PartnerRegistrationImmutableId  pulumi.StringPtrOutput         `pulumi:"partnerRegistrationImmutableId"`
	PartnerTopicFriendlyDescription pulumi.StringPtrOutput         `pulumi:"partnerTopicFriendlyDescription"`
	ProvisioningState               pulumi.StringOutput            `pulumi:"provisioningState"`
	Source                          pulumi.StringPtrOutput         `pulumi:"source"`
	SystemData                      SystemDataResponseOutput       `pulumi:"systemData"`
	Tags                            pulumi.StringMapOutput         `pulumi:"tags"`
	Type                            pulumi.StringOutput            `pulumi:"type"`
}


func NewPartnerTopic(ctx *pulumi.Context,
	name string, args *PartnerTopicArgs, opts ...pulumi.ResourceOption) (*PartnerTopic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:PartnerTopic"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:PartnerTopic"),
		},
	})
	opts = append(opts, aliases)
	var resource PartnerTopic
	err := ctx.RegisterResource("azure-native:eventgrid/v20211015preview:PartnerTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPartnerTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PartnerTopicState, opts ...pulumi.ResourceOption) (*PartnerTopic, error) {
	var resource PartnerTopic
	err := ctx.ReadResource("azure-native:eventgrid/v20211015preview:PartnerTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type partnerTopicState struct {
}

type PartnerTopicState struct {
}

func (PartnerTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerTopicState)(nil)).Elem()
}

type partnerTopicArgs struct {
	ActivationState                 *string           `pulumi:"activationState"`
	EventTypeInfo                   *EventTypeInfo    `pulumi:"eventTypeInfo"`
	ExpirationTimeIfNotActivatedUtc *string           `pulumi:"expirationTimeIfNotActivatedUtc"`
	Identity                        *IdentityInfo     `pulumi:"identity"`
	Location                        *string           `pulumi:"location"`
	MessageForActivation            *string           `pulumi:"messageForActivation"`
	PartnerRegistrationImmutableId  *string           `pulumi:"partnerRegistrationImmutableId"`
	PartnerTopicFriendlyDescription *string           `pulumi:"partnerTopicFriendlyDescription"`
	PartnerTopicName                *string           `pulumi:"partnerTopicName"`
	ResourceGroupName               string            `pulumi:"resourceGroupName"`
	Source                          *string           `pulumi:"source"`
	Tags                            map[string]string `pulumi:"tags"`
}


type PartnerTopicArgs struct {
	ActivationState                 pulumi.StringPtrInput
	EventTypeInfo                   EventTypeInfoPtrInput
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrInput
	Identity                        IdentityInfoPtrInput
	Location                        pulumi.StringPtrInput
	MessageForActivation            pulumi.StringPtrInput
	PartnerRegistrationImmutableId  pulumi.StringPtrInput
	PartnerTopicFriendlyDescription pulumi.StringPtrInput
	PartnerTopicName                pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
	Source                          pulumi.StringPtrInput
	Tags                            pulumi.StringMapInput
}

func (PartnerTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*partnerTopicArgs)(nil)).Elem()
}

type PartnerTopicInput interface {
	pulumi.Input

	ToPartnerTopicOutput() PartnerTopicOutput
	ToPartnerTopicOutputWithContext(ctx context.Context) PartnerTopicOutput
}

func (*PartnerTopic) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerTopic)(nil)).Elem()
}

func (i *PartnerTopic) ToPartnerTopicOutput() PartnerTopicOutput {
	return i.ToPartnerTopicOutputWithContext(context.Background())
}

func (i *PartnerTopic) ToPartnerTopicOutputWithContext(ctx context.Context) PartnerTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PartnerTopicOutput)
}

type PartnerTopicOutput struct{ *pulumi.OutputState }

func (PartnerTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerTopic)(nil)).Elem()
}

func (o PartnerTopicOutput) ToPartnerTopicOutput() PartnerTopicOutput {
	return o
}

func (o PartnerTopicOutput) ToPartnerTopicOutputWithContext(ctx context.Context) PartnerTopicOutput {
	return o
}

func (o PartnerTopicOutput) ActivationState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.ActivationState }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicOutput) EventTypeInfo() EventTypeInfoResponsePtrOutput {
	return o.ApplyT(func(v *PartnerTopic) EventTypeInfoResponsePtrOutput { return v.EventTypeInfo }).(EventTypeInfoResponsePtrOutput)
}

func (o PartnerTopicOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicOutput) Identity() IdentityInfoResponsePtrOutput {
	return o.ApplyT(func(v *PartnerTopic) IdentityInfoResponsePtrOutput { return v.Identity }).(IdentityInfoResponsePtrOutput)
}

func (o PartnerTopicOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PartnerTopicOutput) MessageForActivation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.MessageForActivation }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PartnerTopicOutput) PartnerRegistrationImmutableId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.PartnerRegistrationImmutableId }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicOutput) PartnerTopicFriendlyDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.PartnerTopicFriendlyDescription }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PartnerTopicOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

func (o PartnerTopicOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PartnerTopic) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PartnerTopicOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PartnerTopicOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PartnerTopic) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PartnerTopicOutput{})
}
