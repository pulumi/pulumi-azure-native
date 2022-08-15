


package v20220615

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SystemTopicEventSubscription struct {
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


func NewSystemTopicEventSubscription(ctx *pulumi.Context,
	name string, args *SystemTopicEventSubscriptionArgs, opts ...pulumi.ResourceOption) (*SystemTopicEventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SystemTopicName == nil {
		return nil, errors.New("invalid value for required argument 'SystemTopicName'")
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
			Type: pulumi.String("azure-native:eventgrid:SystemTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20200401preview:SystemTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20201015preview:SystemTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20210601preview:SystemTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:SystemTopicEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20211201:SystemTopicEventSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource SystemTopicEventSubscription
	err := ctx.RegisterResource("azure-native:eventgrid/v20220615:SystemTopicEventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSystemTopicEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SystemTopicEventSubscriptionState, opts ...pulumi.ResourceOption) (*SystemTopicEventSubscription, error) {
	var resource SystemTopicEventSubscription
	err := ctx.ReadResource("azure-native:eventgrid/v20220615:SystemTopicEventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type systemTopicEventSubscriptionState struct {
}

type SystemTopicEventSubscriptionState struct {
}

func (SystemTopicEventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicEventSubscriptionState)(nil)).Elem()
}

type systemTopicEventSubscriptionArgs struct {
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
	SystemTopicName                string                            `pulumi:"systemTopicName"`
}


type SystemTopicEventSubscriptionArgs struct {
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
	SystemTopicName                pulumi.StringInput
}

func (SystemTopicEventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*systemTopicEventSubscriptionArgs)(nil)).Elem()
}

type SystemTopicEventSubscriptionInput interface {
	pulumi.Input

	ToSystemTopicEventSubscriptionOutput() SystemTopicEventSubscriptionOutput
	ToSystemTopicEventSubscriptionOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionOutput
}

func (*SystemTopicEventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemTopicEventSubscription)(nil)).Elem()
}

func (i *SystemTopicEventSubscription) ToSystemTopicEventSubscriptionOutput() SystemTopicEventSubscriptionOutput {
	return i.ToSystemTopicEventSubscriptionOutputWithContext(context.Background())
}

func (i *SystemTopicEventSubscription) ToSystemTopicEventSubscriptionOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemTopicEventSubscriptionOutput)
}

type SystemTopicEventSubscriptionOutput struct{ *pulumi.OutputState }

func (SystemTopicEventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemTopicEventSubscription)(nil)).Elem()
}

func (o SystemTopicEventSubscriptionOutput) ToSystemTopicEventSubscriptionOutput() SystemTopicEventSubscriptionOutput {
	return o
}

func (o SystemTopicEventSubscriptionOutput) ToSystemTopicEventSubscriptionOutputWithContext(ctx context.Context) SystemTopicEventSubscriptionOutput {
	return o
}

func (o SystemTopicEventSubscriptionOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) StorageBlobDeadLetterDestinationResponsePtrOutput {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o SystemTopicEventSubscriptionOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) DeadLetterWithResourceIdentityResponsePtrOutput {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

func (o SystemTopicEventSubscriptionOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) DeliveryWithResourceIdentityResponsePtrOutput {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

func (o SystemTopicEventSubscriptionOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) pulumi.AnyOutput { return v.Destination }).(pulumi.AnyOutput)
}

func (o SystemTopicEventSubscriptionOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) pulumi.StringPtrOutput { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

func (o SystemTopicEventSubscriptionOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) pulumi.StringPtrOutput { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o SystemTopicEventSubscriptionOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) EventSubscriptionFilterResponsePtrOutput { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o SystemTopicEventSubscriptionOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o SystemTopicEventSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SystemTopicEventSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o SystemTopicEventSubscriptionOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) RetryPolicyResponsePtrOutput { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o SystemTopicEventSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o SystemTopicEventSubscriptionOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

func (o SystemTopicEventSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SystemTopicEventSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SystemTopicEventSubscriptionOutput{})
}
