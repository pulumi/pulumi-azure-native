


package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDomainEventSubscription(ctx *pulumi.Context, args *LookupDomainEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupDomainEventSubscriptionResult, error) {
	var rv LookupDomainEventSubscriptionResult
	err := ctx.Invoke("azure-native:eventgrid/v20220615:getDomainEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDomainEventSubscriptionArgs struct {
	DomainName            string `pulumi:"domainName"`
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupDomainEventSubscriptionResult struct {
	DeadLetterDestination          *StorageBlobDeadLetterDestinationResponse `pulumi:"deadLetterDestination"`
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentityResponse   `pulumi:"deadLetterWithResourceIdentity"`
	DeliveryWithResourceIdentity   *DeliveryWithResourceIdentityResponse     `pulumi:"deliveryWithResourceIdentity"`
	Destination                    interface{}                               `pulumi:"destination"`
	EventDeliverySchema            *string                                   `pulumi:"eventDeliverySchema"`
	ExpirationTimeUtc              *string                                   `pulumi:"expirationTimeUtc"`
	Filter                         *EventSubscriptionFilterResponse          `pulumi:"filter"`
	Id                             string                                    `pulumi:"id"`
	Labels                         []string                                  `pulumi:"labels"`
	Name                           string                                    `pulumi:"name"`
	ProvisioningState              string                                    `pulumi:"provisioningState"`
	RetryPolicy                    *RetryPolicyResponse                      `pulumi:"retryPolicy"`
	SystemData                     SystemDataResponse                        `pulumi:"systemData"`
	Topic                          string                                    `pulumi:"topic"`
	Type                           string                                    `pulumi:"type"`
}


func (val *LookupDomainEventSubscriptionResult) Defaults() *LookupDomainEventSubscriptionResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EventDeliverySchema) {
		eventDeliverySchema_ := "EventGridSchema"
		tmp.EventDeliverySchema = &eventDeliverySchema_
	}
	tmp.Filter = tmp.Filter.Defaults()

	tmp.RetryPolicy = tmp.RetryPolicy.Defaults()

	return &tmp
}

func LookupDomainEventSubscriptionOutput(ctx *pulumi.Context, args LookupDomainEventSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupDomainEventSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainEventSubscriptionResult, error) {
			args := v.(LookupDomainEventSubscriptionArgs)
			r, err := LookupDomainEventSubscription(ctx, &args, opts...)
			var s LookupDomainEventSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainEventSubscriptionResultOutput)
}

type LookupDomainEventSubscriptionOutputArgs struct {
	DomainName            pulumi.StringInput `pulumi:"domainName"`
	EventSubscriptionName pulumi.StringInput `pulumi:"eventSubscriptionName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDomainEventSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainEventSubscriptionArgs)(nil)).Elem()
}


type LookupDomainEventSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupDomainEventSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainEventSubscriptionResult)(nil)).Elem()
}

func (o LookupDomainEventSubscriptionResultOutput) ToLookupDomainEventSubscriptionResultOutput() LookupDomainEventSubscriptionResultOutput {
	return o
}

func (o LookupDomainEventSubscriptionResultOutput) ToLookupDomainEventSubscriptionResultOutputWithContext(ctx context.Context) LookupDomainEventSubscriptionResultOutput {
	return o
}

func (o LookupDomainEventSubscriptionResultOutput) DeadLetterDestination() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *StorageBlobDeadLetterDestinationResponse {
		return v.DeadLetterDestination
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) DeadLetterWithResourceIdentity() DeadLetterWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *DeadLetterWithResourceIdentityResponse {
		return v.DeadLetterWithResourceIdentity
	}).(DeadLetterWithResourceIdentityResponsePtrOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) DeliveryWithResourceIdentity() DeliveryWithResourceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *DeliveryWithResourceIdentityResponse {
		return v.DeliveryWithResourceIdentity
	}).(DeliveryWithResourceIdentityResponsePtrOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) Destination() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) interface{} { return v.Destination }).(pulumi.AnyOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) EventDeliverySchema() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *string { return v.EventDeliverySchema }).(pulumi.StringPtrOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *string { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) Filter() EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *EventSubscriptionFilterResponse { return v.Filter }).(EventSubscriptionFilterResponsePtrOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) RetryPolicy() RetryPolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) *RetryPolicyResponse { return v.RetryPolicy }).(RetryPolicyResponsePtrOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) string { return v.Topic }).(pulumi.StringOutput)
}

func (o LookupDomainEventSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainEventSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainEventSubscriptionResultOutput{})
}
