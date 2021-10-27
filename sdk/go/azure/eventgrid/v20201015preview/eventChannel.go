


package v20201015preview

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
			Type: pulumi.String("azure-nextgen:eventgrid/v20201015preview:EventChannel"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid:EventChannel"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid:EventChannel"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:EventChannel"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20200401preview:EventChannel"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:EventChannel"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventgrid/v20210601preview:EventChannel"),
		},
	})
	opts = append(opts, aliases)
	var resource EventChannel
	err := ctx.RegisterResource("azure-native:eventgrid/v20201015preview:EventChannel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventChannelState, opts ...pulumi.ResourceOption) (*EventChannel, error) {
	var resource EventChannel
	err := ctx.ReadResource("azure-native:eventgrid/v20201015preview:EventChannel", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*EventChannel)(nil))
}

func (i *EventChannel) ToEventChannelOutput() EventChannelOutput {
	return i.ToEventChannelOutputWithContext(context.Background())
}

func (i *EventChannel) ToEventChannelOutputWithContext(ctx context.Context) EventChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventChannelOutput)
}

type EventChannelOutput struct{ *pulumi.OutputState }

func (EventChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventChannel)(nil))
}

func (o EventChannelOutput) ToEventChannelOutput() EventChannelOutput {
	return o
}

func (o EventChannelOutput) ToEventChannelOutputWithContext(ctx context.Context) EventChannelOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EventChannelInput)(nil)).Elem(), &EventChannel{})
	pulumi.RegisterOutputType(EventChannelOutput{})
}
