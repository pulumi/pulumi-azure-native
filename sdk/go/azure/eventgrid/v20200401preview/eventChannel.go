


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventChannel struct {
	pulumi.CustomResourceState

	Destination                     EventChannelDestinationResponsePtrOutput `pulumi:"destination"`
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrOutput                   `pulumi:"expirationTimeIfNotActivatedUtc"`
	Filter                          EventChannelFilterResponsePtrOutput      `pulumi:"filter"`
	Name                            pulumi.StringOutput                      `pulumi:"name"`
	PartnerTopicFriendlyDescription pulumi.StringPtrOutput                   `pulumi:"partnerTopicFriendlyDescription"`
	PartnerTopicReadinessState      pulumi.StringOutput                      `pulumi:"partnerTopicReadinessState"`
	ProvisioningState               pulumi.StringOutput                      `pulumi:"provisioningState"`
	Source                          EventChannelSourceResponsePtrOutput      `pulumi:"source"`
	SystemData                      SystemDataResponseOutput                 `pulumi:"systemData"`
	Type                            pulumi.StringOutput                      `pulumi:"type"`
}


func NewEventChannel(ctx *pulumi.Context,
	name string, args *EventChannelArgs, opts ...pulumi.ResourceOption) (*EventChannel, error) {
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
			Type: pulumi.String("azure-native:eventgrid:EventChannel"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:EventChannel"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:EventChannel"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:EventChannel"),
		},
	})
	opts = append(opts, aliases)
	var resource EventChannel
	err := ctx.RegisterResource("azure-native:eventgrid/v20200401preview:EventChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventChannelState, opts ...pulumi.ResourceOption) (*EventChannel, error) {
	var resource EventChannel
	err := ctx.ReadResource("azure-native:eventgrid/v20200401preview:EventChannel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventChannelState struct {
}

type EventChannelState struct {
}

func (EventChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventChannelState)(nil)).Elem()
}

type eventChannelArgs struct {
	Destination                     *EventChannelDestination `pulumi:"destination"`
	EventChannelName                *string                  `pulumi:"eventChannelName"`
	ExpirationTimeIfNotActivatedUtc *string                  `pulumi:"expirationTimeIfNotActivatedUtc"`
	Filter                          *EventChannelFilter      `pulumi:"filter"`
	PartnerNamespaceName            string                   `pulumi:"partnerNamespaceName"`
	PartnerTopicFriendlyDescription *string                  `pulumi:"partnerTopicFriendlyDescription"`
	ResourceGroupName               string                   `pulumi:"resourceGroupName"`
	Source                          *EventChannelSource      `pulumi:"source"`
}


type EventChannelArgs struct {
	Destination                     EventChannelDestinationPtrInput
	EventChannelName                pulumi.StringPtrInput
	ExpirationTimeIfNotActivatedUtc pulumi.StringPtrInput
	Filter                          EventChannelFilterPtrInput
	PartnerNamespaceName            pulumi.StringInput
	PartnerTopicFriendlyDescription pulumi.StringPtrInput
	ResourceGroupName               pulumi.StringInput
	Source                          EventChannelSourcePtrInput
}

func (EventChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventChannelArgs)(nil)).Elem()
}

type EventChannelInput interface {
	pulumi.Input

	ToEventChannelOutput() EventChannelOutput
	ToEventChannelOutputWithContext(ctx context.Context) EventChannelOutput
}

func (*EventChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannel)(nil)).Elem()
}

func (i *EventChannel) ToEventChannelOutput() EventChannelOutput {
	return i.ToEventChannelOutputWithContext(context.Background())
}

func (i *EventChannel) ToEventChannelOutputWithContext(ctx context.Context) EventChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelOutput)
}

type EventChannelOutput struct{ *pulumi.OutputState }

func (EventChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventChannel)(nil)).Elem()
}

func (o EventChannelOutput) ToEventChannelOutput() EventChannelOutput {
	return o
}

func (o EventChannelOutput) ToEventChannelOutputWithContext(ctx context.Context) EventChannelOutput {
	return o
}

func (o EventChannelOutput) Destination() EventChannelDestinationResponsePtrOutput {
	return o.ApplyT(func(v *EventChannel) EventChannelDestinationResponsePtrOutput { return v.Destination }).(EventChannelDestinationResponsePtrOutput)
}

func (o EventChannelOutput) ExpirationTimeIfNotActivatedUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventChannel) pulumi.StringPtrOutput { return v.ExpirationTimeIfNotActivatedUtc }).(pulumi.StringPtrOutput)
}

func (o EventChannelOutput) Filter() EventChannelFilterResponsePtrOutput {
	return o.ApplyT(func(v *EventChannel) EventChannelFilterResponsePtrOutput { return v.Filter }).(EventChannelFilterResponsePtrOutput)
}

func (o EventChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventChannel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EventChannelOutput) PartnerTopicFriendlyDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventChannel) pulumi.StringPtrOutput { return v.PartnerTopicFriendlyDescription }).(pulumi.StringPtrOutput)
}

func (o EventChannelOutput) PartnerTopicReadinessState() pulumi.StringOutput {
	return o.ApplyT(func(v *EventChannel) pulumi.StringOutput { return v.PartnerTopicReadinessState }).(pulumi.StringOutput)
}

func (o EventChannelOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EventChannel) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EventChannelOutput) Source() EventChannelSourceResponsePtrOutput {
	return o.ApplyT(func(v *EventChannel) EventChannelSourceResponsePtrOutput { return v.Source }).(EventChannelSourceResponsePtrOutput)
}

func (o EventChannelOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *EventChannel) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o EventChannelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventChannel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventChannelOutput{})
}
