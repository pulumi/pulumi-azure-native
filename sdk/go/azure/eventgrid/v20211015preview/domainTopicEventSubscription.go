


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DomainTopicEventSubscription struct {
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


func NewDomainTopicEventSubscription(ctx *pulumi.Context,
	name string, args *DomainTopicEventSubscriptionArgs, opts ...pulumi.ResourceOption) (*DomainTopicEventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
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
			Type: pulumi.String("azure-native:eventgrid:DomainTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:DomainTopicEventSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource DomainTopicEventSubscription
	err := ctx.RegisterResource("azure-native:eventgrid/v20211015preview:DomainTopicEventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDomainTopicEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainTopicEventSubscriptionState, opts ...pulumi.ResourceOption) (*DomainTopicEventSubscription, error) {
	var resource DomainTopicEventSubscription
	err := ctx.ReadResource("azure-native:eventgrid/v20211015preview:DomainTopicEventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type domainTopicEventSubscriptionState struct {
}

type DomainTopicEventSubscriptionState struct {
}

func (DomainTopicEventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainTopicEventSubscriptionState)(nil)).Elem()
}

type domainTopicEventSubscriptionArgs struct {
	DeadLetterDestination          *StorageBlobDeadLetterDestination `pulumi:"deadLetterDestination"`
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentity   `pulumi:"deadLetterWithResourceIdentity"`
	DeliveryWithResourceIdentity   *DeliveryWithResourceIdentity     `pulumi:"deliveryWithResourceIdentity"`
	Destination                    interface{}                       `pulumi:"destination"`
	DomainName                     string                            `pulumi:"domainName"`
	EventDeliverySchema            *string                           `pulumi:"eventDeliverySchema"`
	EventSubscriptionName          *string                           `pulumi:"eventSubscriptionName"`
	ExpirationTimeUtc              *string                           `pulumi:"expirationTimeUtc"`
	Filter                         *EventSubscriptionFilter          `pulumi:"filter"`
	Labels                         []string                          `pulumi:"labels"`
	ResourceGroupName              string                            `pulumi:"resourceGroupName"`
	RetryPolicy                    *RetryPolicy                      `pulumi:"retryPolicy"`
	TopicName                      string                            `pulumi:"topicName"`
}


type DomainTopicEventSubscriptionArgs struct {
	DeadLetterDestination          StorageBlobDeadLetterDestinationPtrInput
	DeadLetterWithResourceIdentity DeadLetterWithResourceIdentityPtrInput
	DeliveryWithResourceIdentity   DeliveryWithResourceIdentityPtrInput
	Destination                    pulumi.Input
	DomainName                     pulumi.StringInput
	EventDeliverySchema            pulumi.StringPtrInput
	EventSubscriptionName          pulumi.StringPtrInput
	ExpirationTimeUtc              pulumi.StringPtrInput
	Filter                         EventSubscriptionFilterPtrInput
	Labels                         pulumi.StringArrayInput
	ResourceGroupName              pulumi.StringInput
	RetryPolicy                    RetryPolicyPtrInput
	TopicName                      pulumi.StringInput
}

func (DomainTopicEventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainTopicEventSubscriptionArgs)(nil)).Elem()
}

type DomainTopicEventSubscriptionInput interface {
	pulumi.Input

	ToDomainTopicEventSubscriptionOutput() DomainTopicEventSubscriptionOutput
	ToDomainTopicEventSubscriptionOutputWithContext(ctx context.Context) DomainTopicEventSubscriptionOutput
}

func (*DomainTopicEventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainTopicEventSubscription)(nil)).Elem()
}

func (i *DomainTopicEventSubscription) ToDomainTopicEventSubscriptionOutput() DomainTopicEventSubscriptionOutput {
	return i.ToDomainTopicEventSubscriptionOutputWithContext(context.Background())
}

func (i *DomainTopicEventSubscription) ToDomainTopicEventSubscriptionOutputWithContext(ctx context.Context) DomainTopicEventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainTopicEventSubscriptionOutput)
}

type DomainTopicEventSubscriptionOutput struct{ *pulumi.OutputState }

func (DomainTopicEventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainTopicEventSubscription)(nil)).Elem()
}

func (o DomainTopicEventSubscriptionOutput) ToDomainTopicEventSubscriptionOutput() DomainTopicEventSubscriptionOutput {
	return o
}

func (o DomainTopicEventSubscriptionOutput) ToDomainTopicEventSubscriptionOutputWithContext(ctx context.Context) DomainTopicEventSubscriptionOutput {
	return o
}

func (o DomainTopicEventSubscriptionOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) StorageBlobDeadLetterDestinationResponsePtrOutput {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o DomainTopicEventSubscriptionOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) DeadLetterWithResourceIdentityResponsePtrOutput {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

func (o DomainTopicEventSubscriptionOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) DeliveryWithResourceIdentityResponsePtrOutput {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

func (o DomainTopicEventSubscriptionOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) pulumi.AnyOutput { return v.Destination }).(pulumi.AnyOutput)
}

func (o DomainTopicEventSubscriptionOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) pulumi.StringPtrOutput { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

func (o DomainTopicEventSubscriptionOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) pulumi.StringPtrOutput { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o DomainTopicEventSubscriptionOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) EventSubscriptionFilterResponsePtrOutput { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o DomainTopicEventSubscriptionOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o DomainTopicEventSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DomainTopicEventSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DomainTopicEventSubscriptionOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) RetryPolicyResponsePtrOutput { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o DomainTopicEventSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DomainTopicEventSubscriptionOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

func (o DomainTopicEventSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainTopicEventSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainTopicEventSubscriptionOutput{})
}
