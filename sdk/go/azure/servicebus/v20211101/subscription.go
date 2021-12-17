


package v20211101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Subscription struct {
	pulumi.CustomResourceState

	AccessedAt                                pulumi.StringOutput                       `pulumi:"accessedAt"`
	AutoDeleteOnIdle                          pulumi.StringPtrOutput                    `pulumi:"autoDeleteOnIdle"`
	ClientAffineProperties                    SBClientAffinePropertiesResponsePtrOutput `pulumi:"clientAffineProperties"`
	CountDetails                              MessageCountDetailsResponseOutput         `pulumi:"countDetails"`
	CreatedAt                                 pulumi.StringOutput                       `pulumi:"createdAt"`
	DeadLetteringOnFilterEvaluationExceptions pulumi.BoolPtrOutput                      `pulumi:"deadLetteringOnFilterEvaluationExceptions"`
	DeadLetteringOnMessageExpiration          pulumi.BoolPtrOutput                      `pulumi:"deadLetteringOnMessageExpiration"`
	DefaultMessageTimeToLive                  pulumi.StringPtrOutput                    `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow       pulumi.StringPtrOutput                    `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations                   pulumi.BoolPtrOutput                      `pulumi:"enableBatchedOperations"`
	ForwardDeadLetteredMessagesTo             pulumi.StringPtrOutput                    `pulumi:"forwardDeadLetteredMessagesTo"`
	ForwardTo                                 pulumi.StringPtrOutput                    `pulumi:"forwardTo"`
	IsClientAffine                            pulumi.BoolPtrOutput                      `pulumi:"isClientAffine"`
	Location                                  pulumi.StringOutput                       `pulumi:"location"`
	LockDuration                              pulumi.StringPtrOutput                    `pulumi:"lockDuration"`
	MaxDeliveryCount                          pulumi.IntPtrOutput                       `pulumi:"maxDeliveryCount"`
	MessageCount                              pulumi.Float64Output                      `pulumi:"messageCount"`
	Name                                      pulumi.StringOutput                       `pulumi:"name"`
	RequiresSession                           pulumi.BoolPtrOutput                      `pulumi:"requiresSession"`
	Status                                    pulumi.StringPtrOutput                    `pulumi:"status"`
	SystemData                                SystemDataResponseOutput                  `pulumi:"systemData"`
	Type                                      pulumi.StringOutput                       `pulumi:"type"`
	UpdatedAt                                 pulumi.StringOutput                       `pulumi:"updatedAt"`
}


func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:servicebus:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20140901:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20150801:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20170401:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20180101preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210101preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-native:servicebus/v20210601preview:Subscription"),
		},
	})
	opts = append(opts, aliases)
	var resource Subscription
	err := ctx.RegisterResource("azure-native:servicebus/v20211101:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure-native:servicebus/v20211101:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subscriptionState struct {
}

type SubscriptionState struct {
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	AutoDeleteOnIdle                          *string                   `pulumi:"autoDeleteOnIdle"`
	ClientAffineProperties                    *SBClientAffineProperties `pulumi:"clientAffineProperties"`
	DeadLetteringOnFilterEvaluationExceptions *bool                     `pulumi:"deadLetteringOnFilterEvaluationExceptions"`
	DeadLetteringOnMessageExpiration          *bool                     `pulumi:"deadLetteringOnMessageExpiration"`
	DefaultMessageTimeToLive                  *string                   `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow       *string                   `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations                   *bool                     `pulumi:"enableBatchedOperations"`
	ForwardDeadLetteredMessagesTo             *string                   `pulumi:"forwardDeadLetteredMessagesTo"`
	ForwardTo                                 *string                   `pulumi:"forwardTo"`
	IsClientAffine                            *bool                     `pulumi:"isClientAffine"`
	LockDuration                              *string                   `pulumi:"lockDuration"`
	MaxDeliveryCount                          *int                      `pulumi:"maxDeliveryCount"`
	NamespaceName                             string                    `pulumi:"namespaceName"`
	RequiresSession                           *bool                     `pulumi:"requiresSession"`
	ResourceGroupName                         string                    `pulumi:"resourceGroupName"`
	Status                                    *EntityStatus             `pulumi:"status"`
	SubscriptionName                          *string                   `pulumi:"subscriptionName"`
	TopicName                                 string                    `pulumi:"topicName"`
}


type SubscriptionArgs struct {
	AutoDeleteOnIdle                          pulumi.StringPtrInput
	ClientAffineProperties                    SBClientAffinePropertiesPtrInput
	DeadLetteringOnFilterEvaluationExceptions pulumi.BoolPtrInput
	DeadLetteringOnMessageExpiration          pulumi.BoolPtrInput
	DefaultMessageTimeToLive                  pulumi.StringPtrInput
	DuplicateDetectionHistoryTimeWindow       pulumi.StringPtrInput
	EnableBatchedOperations                   pulumi.BoolPtrInput
	ForwardDeadLetteredMessagesTo             pulumi.StringPtrInput
	ForwardTo                                 pulumi.StringPtrInput
	IsClientAffine                            pulumi.BoolPtrInput
	LockDuration                              pulumi.StringPtrInput
	MaxDeliveryCount                          pulumi.IntPtrInput
	NamespaceName                             pulumi.StringInput
	RequiresSession                           pulumi.BoolPtrInput
	ResourceGroupName                         pulumi.StringInput
	Status                                    EntityStatusPtrInput
	SubscriptionName                          pulumi.StringPtrInput
	TopicName                                 pulumi.StringInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}

type SubscriptionInput interface {
	pulumi.Input

	ToSubscriptionOutput() SubscriptionOutput
	ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput
}

func (*Subscription) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (i *Subscription) ToSubscriptionOutput() SubscriptionOutput {
	return i.ToSubscriptionOutputWithContext(context.Background())
}

func (i *Subscription) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionOutput)
}

type SubscriptionOutput struct{ *pulumi.OutputState }

func (SubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Subscription)(nil)).Elem()
}

func (o SubscriptionOutput) ToSubscriptionOutput() SubscriptionOutput {
	return o
}

func (o SubscriptionOutput) ToSubscriptionOutputWithContext(ctx context.Context) SubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubscriptionOutput{})
}
