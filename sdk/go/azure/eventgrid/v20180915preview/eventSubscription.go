


package v20180915preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type EventSubscription struct {
	pulumi.CustomResourceState

	DeadLetterDestination StorageBlobDeadLetterDestinationResponsePtrOutput `pulumi:"deadLetterDestination"`
	Destination           pulumi.AnyOutput                                  `pulumi:"destination"`
	EventDeliverySchema   pulumi.StringPtrOutput                            `pulumi:"eventDeliverySchema"`
	ExpirationTimeUtc     pulumi.StringPtrOutput                            `pulumi:"expirationTimeUtc"`
	Filter                EventSubscriptionFilterResponsePtrOutput          `pulumi:"filter"`
	Labels                pulumi.StringArrayOutput                          `pulumi:"labels"`
	Name                  pulumi.StringOutput                               `pulumi:"name"`
	ProvisioningState     pulumi.StringOutput                               `pulumi:"provisioningState"`
	RetryPolicy           RetryPolicyResponsePtrOutput                      `pulumi:"retryPolicy"`
	Topic                 pulumi.StringOutput                               `pulumi:"topic"`
	Type                  pulumi.StringOutput                               `pulumi:"type"`
}


func NewEventSubscription(ctx *pulumi.Context,
	name string, args *EventSubscriptionArgs, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.Filter != nil {
		args.Filter = args.Filter.ToEventSubscriptionFilterPtrOutput().ApplyT(func(v *EventSubscriptionFilter) *EventSubscriptionFilter { return v.Defaults() }).(EventSubscriptionFilterPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170615preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20170915preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180101:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20180501preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190101:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190201preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20190601:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200101preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200601:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:EventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:EventSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource EventSubscription
	err := ctx.RegisterResource("azure-native:eventgrid/v20180915preview:EventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSubscriptionState, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	var resource EventSubscription
	err := ctx.ReadResource("azure-native:eventgrid/v20180915preview:EventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type eventSubscriptionState struct {
}

type EventSubscriptionState struct {
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
	DeadLetterDestination *StorageBlobDeadLetterDestination `pulumi:"deadLetterDestination"`
	Destination           interface{}                       `pulumi:"destination"`
	EventDeliverySchema   *string                           `pulumi:"eventDeliverySchema"`
	EventSubscriptionName *string                           `pulumi:"eventSubscriptionName"`
	ExpirationTimeUtc     *string                           `pulumi:"expirationTimeUtc"`
	Filter                *EventSubscriptionFilter          `pulumi:"filter"`
	Labels                []string                          `pulumi:"labels"`
	RetryPolicy           *RetryPolicy                      `pulumi:"retryPolicy"`
	Scope                 string                            `pulumi:"scope"`
}


type EventSubscriptionArgs struct {
	DeadLetterDestination StorageBlobDeadLetterDestinationPtrInput
	Destination           pulumi.Input
	EventDeliverySchema   pulumi.StringPtrInput
	EventSubscriptionName pulumi.StringPtrInput
	ExpirationTimeUtc     pulumi.StringPtrInput
	Filter                EventSubscriptionFilterPtrInput
	Labels                pulumi.StringArrayInput
	RetryPolicy           RetryPolicyPtrInput
	Scope                 pulumi.StringInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}

type EventSubscriptionInput interface {
	pulumi.Input

	ToEventSubscriptionOutput() EventSubscriptionOutput
	ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput
}

func (*EventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscription)(nil)).Elem()
}

func (i *EventSubscription) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return i.ToEventSubscriptionOutputWithContext(context.Background())
}

func (i *EventSubscription) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionOutput)
}

type EventSubscriptionOutput struct{ *pulumi.OutputState }

func (EventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscription)(nil)).Elem()
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutput() EventSubscriptionOutput {
	return o
}

func (o EventSubscriptionOutput) ToEventSubscriptionOutputWithContext(ctx context.Context) EventSubscriptionOutput {
	return o
}

func (o EventSubscriptionOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v *EventSubscription) StorageBlobDeadLetterDestinationResponsePtrOutput {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o EventSubscriptionOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.AnyOutput { return v.Destination }).(pulumi.AnyOutput)
}

func (o EventSubscriptionOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringPtrOutput { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringPtrOutput { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v *EventSubscription) EventSubscriptionFilterResponsePtrOutput { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o EventSubscriptionOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EventSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o EventSubscriptionOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *EventSubscription) RetryPolicyResponsePtrOutput { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o EventSubscriptionOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

func (o EventSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EventSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EventSubscriptionOutput{})
}
