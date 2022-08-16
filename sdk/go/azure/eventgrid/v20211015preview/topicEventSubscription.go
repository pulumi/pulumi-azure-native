


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TopicEventSubscription struct {
	pulumi.CustomResourceState

	DeadLetterDestination          StorageBlobDeadLetterDestinationResponsePtrOutput `pulumi:"deadLetterDestination"`
	DeadLetterWithResourceIdentity DeadLetterWithResourceIdentityResponsePtrOutput   `pulumi:"deadLetterWithResourceIdentity"`
	DeliveryWithResourceIdentity   DeliveryWithResourceIdentityResponsePtrOutput     `pulumi:"deliveryWithResourceIdentity"`
	Destination                    pulumi.AnyOutput                                  `pulumi:"destination"`
	EventDeliverySchema            pulumi.StringPtrOutput                            `pulumi:"eventDeliverySchema"`
	ExpirationTimeUtc              pulumi.StringPtrOutput                            `pulumi:"expirationTimeUtc"`
	Filter                         EventSubscriptionFilterResponsePtrOutput          `pulumi:"filter"`
	Labels                         pulumi.StringArrayOutput                          `pulumi:"labels"`
	Name                           pulumi.StringOutput                               `pulumi:"name"`
	ProvisioningState              pulumi.StringOutput                               `pulumi:"provisioningState"`
	RetryPolicy                    RetryPolicyResponsePtrOutput                      `pulumi:"retryPolicy"`
	SystemData                     SystemDataResponseOutput                          `pulumi:"systemData"`
	Topic                          pulumi.StringOutput                               `pulumi:"topic"`
	Type                           pulumi.StringOutput                               `pulumi:"type"`
}


func NewTopicEventSubscription(ctx *pulumi.Context,
	name string, args *TopicEventSubscriptionArgs, opts ...pulumi.ResourceOption) (*TopicEventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	if isZero(args.EventDeliverySchema) {
		args.EventDeliverySchema = pulumi.StringPtr("EventGridSchema")
	}
	if args.Filter != nil {
		args.Filter = args.Filter.ToEventSubscriptionFilterPtrOutput().ApplyT(func(v *EventSubscriptionFilter) *EventSubscriptionFilter { return v.Defaults() }).(EventSubscriptionFilterPtrOutput)
	}
	if args.RetryPolicy != nil {
		args.RetryPolicy = args.RetryPolicy.ToRetryPolicyPtrOutput().ApplyT(func(v *RetryPolicy) *RetryPolicy { return v.Defaults() }).(RetryPolicyPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:eventgrid:TopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:TopicEventSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource TopicEventSubscription
	err := ctx.RegisterResource("azure-native:eventgrid/v20211015preview:TopicEventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTopicEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicEventSubscriptionState, opts ...pulumi.ResourceOption) (*TopicEventSubscription, error) {
	var resource TopicEventSubscription
	err := ctx.ReadResource("azure-native:eventgrid/v20211015preview:TopicEventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type topicEventSubscriptionState struct {
}

type TopicEventSubscriptionState struct {
}

func (TopicEventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicEventSubscriptionState)(nil)).Elem()
}

type topicEventSubscriptionArgs struct {
	DeadLetterDestination          *StorageBlobDeadLetterDestination `pulumi:"deadLetterDestination"`
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentity   `pulumi:"deadLetterWithResourceIdentity"`
	DeliveryWithResourceIdentity   *DeliveryWithResourceIdentity     `pulumi:"deliveryWithResourceIdentity"`
	Destination                    interface{}                       `pulumi:"destination"`
	EventDeliverySchema            *string                           `pulumi:"eventDeliverySchema"`
	EventSubscriptionName          *string                           `pulumi:"eventSubscriptionName"`
	ExpirationTimeUtc              *string                           `pulumi:"expirationTimeUtc"`
	Filter                         *EventSubscriptionFilter          `pulumi:"filter"`
	Labels                         []string                          `pulumi:"labels"`
	ResourceGroupName              string                            `pulumi:"resourceGroupName"`
	RetryPolicy                    *RetryPolicy                      `pulumi:"retryPolicy"`
	TopicName                      string                            `pulumi:"topicName"`
}


type TopicEventSubscriptionArgs struct {
	DeadLetterDestination          StorageBlobDeadLetterDestinationPtrInput
	DeadLetterWithResourceIdentity DeadLetterWithResourceIdentityPtrInput
	DeliveryWithResourceIdentity   DeliveryWithResourceIdentityPtrInput
	Destination                    pulumi.Input
	EventDeliverySchema            pulumi.StringPtrInput
	EventSubscriptionName          pulumi.StringPtrInput
	ExpirationTimeUtc              pulumi.StringPtrInput
	Filter                         EventSubscriptionFilterPtrInput
	Labels                         pulumi.StringArrayInput
	ResourceGroupName              pulumi.StringInput
	RetryPolicy                    RetryPolicyPtrInput
	TopicName                      pulumi.StringInput
}

func (TopicEventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicEventSubscriptionArgs)(nil)).Elem()
}

type TopicEventSubscriptionInput interface {
	pulumi.Input

	ToTopicEventSubscriptionOutput() TopicEventSubscriptionOutput
	ToTopicEventSubscriptionOutputWithContext(ctx context.Context) TopicEventSubscriptionOutput
}

func (*TopicEventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicEventSubscription)(nil)).Elem()
}

func (i *TopicEventSubscription) ToTopicEventSubscriptionOutput() TopicEventSubscriptionOutput {
	return i.ToTopicEventSubscriptionOutputWithContext(context.Background())
}

func (i *TopicEventSubscription) ToTopicEventSubscriptionOutputWithContext(ctx context.Context) TopicEventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicEventSubscriptionOutput)
}

type TopicEventSubscriptionOutput struct{ *pulumi.OutputState }

func (TopicEventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TopicEventSubscription)(nil)).Elem()
}

func (o TopicEventSubscriptionOutput) ToTopicEventSubscriptionOutput() TopicEventSubscriptionOutput {
	return o
}

func (o TopicEventSubscriptionOutput) ToTopicEventSubscriptionOutputWithContext(ctx context.Context) TopicEventSubscriptionOutput {
	return o
}

func (o TopicEventSubscriptionOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) StorageBlobDeadLetterDestinationResponsePtrOutput {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o TopicEventSubscriptionOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) DeadLetterWithResourceIdentityResponsePtrOutput {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

func (o TopicEventSubscriptionOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) DeliveryWithResourceIdentityResponsePtrOutput {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

func (o TopicEventSubscriptionOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.AnyOutput { return v.Destination }).(pulumi.AnyOutput)
}

func (o TopicEventSubscriptionOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringPtrOutput { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

func (o TopicEventSubscriptionOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringPtrOutput { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o TopicEventSubscriptionOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) EventSubscriptionFilterResponsePtrOutput { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o TopicEventSubscriptionOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o TopicEventSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TopicEventSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o TopicEventSubscriptionOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *TopicEventSubscription) RetryPolicyResponsePtrOutput { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o TopicEventSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *TopicEventSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o TopicEventSubscriptionOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

func (o TopicEventSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *TopicEventSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TopicEventSubscriptionOutput{})
}
