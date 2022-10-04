


package eventgrid

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DomainEventSubscription struct {
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


func NewDomainEventSubscription(ctx *pulumi.Context,
	name string, args *DomainEventSubscriptionArgs, opts ...pulumi.ResourceOption) (*DomainEventSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DomainName == nil {
		return nil, errors.New("invalid value for required argument 'DomainName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
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
			Type: pulumi.String("azure-native:eventgrid/v20211015preview:DomainEventSubscription"),
		},
		{
			Type: pulumi.String("azure-native:eventgrid/v20220615:DomainEventSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource DomainEventSubscription
	err := ctx.RegisterResource("azure-native:eventgrid:DomainEventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDomainEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainEventSubscriptionState, opts ...pulumi.ResourceOption) (*DomainEventSubscription, error) {
	var resource DomainEventSubscription
	err := ctx.ReadResource("azure-native:eventgrid:DomainEventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type domainEventSubscriptionState struct {
}

type DomainEventSubscriptionState struct {
}

func (DomainEventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainEventSubscriptionState)(nil)).Elem()
}

type domainEventSubscriptionArgs struct {
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
}


type DomainEventSubscriptionArgs struct {
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
}

func (DomainEventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainEventSubscriptionArgs)(nil)).Elem()
}

type DomainEventSubscriptionInput interface {
	pulumi.Input

	ToDomainEventSubscriptionOutput() DomainEventSubscriptionOutput
	ToDomainEventSubscriptionOutputWithContext(ctx context.Context) DomainEventSubscriptionOutput
}

func (*DomainEventSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainEventSubscription)(nil)).Elem()
}

func (i *DomainEventSubscription) ToDomainEventSubscriptionOutput() DomainEventSubscriptionOutput {
	return i.ToDomainEventSubscriptionOutputWithContext(context.Background())
}

func (i *DomainEventSubscription) ToDomainEventSubscriptionOutputWithContext(ctx context.Context) DomainEventSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainEventSubscriptionOutput)
}

type DomainEventSubscriptionOutput struct{ *pulumi.OutputState }

func (DomainEventSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DomainEventSubscription)(nil)).Elem()
}

func (o DomainEventSubscriptionOutput) ToDomainEventSubscriptionOutput() DomainEventSubscriptionOutput {
	return o
}

func (o DomainEventSubscriptionOutput) ToDomainEventSubscriptionOutputWithContext(ctx context.Context) DomainEventSubscriptionOutput {
	return o
}

func (o DomainEventSubscriptionOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) StorageBlobDeadLetterDestinationResponsePtrOutput {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o DomainEventSubscriptionOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) DeadLetterWithResourceIdentityResponsePtrOutput {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

func (o DomainEventSubscriptionOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) DeliveryWithResourceIdentityResponsePtrOutput {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

func (o DomainEventSubscriptionOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.AnyOutput { return v.Destination }).(pulumi.AnyOutput)
}

func (o DomainEventSubscriptionOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringPtrOutput { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

func (o DomainEventSubscriptionOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringPtrOutput { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o DomainEventSubscriptionOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) EventSubscriptionFilterResponsePtrOutput { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o DomainEventSubscriptionOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringArrayOutput { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o DomainEventSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DomainEventSubscriptionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o DomainEventSubscriptionOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v *DomainEventSubscription) RetryPolicyResponsePtrOutput { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o DomainEventSubscriptionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *DomainEventSubscription) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o DomainEventSubscriptionOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

func (o DomainEventSubscriptionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DomainEventSubscription) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(DomainEventSubscriptionOutput{})
}
