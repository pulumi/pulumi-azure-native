


package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubscription(ctx *pulumi.Context, args *LookupSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionResult, error) {
	var rv LookupSubscriptionResult
	err := ctx.Invoke("azure-native:servicebus:getSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SubscriptionName  string `pulumi:"subscriptionName"`
	TopicName         string `pulumi:"topicName"`
}


type LookupSubscriptionResult struct {
	AccessedAt                                string                      `pulumi:"accessedAt"`
	AutoDeleteOnIdle                          *string                     `pulumi:"autoDeleteOnIdle"`
	CountDetails                              MessageCountDetailsResponse `pulumi:"countDetails"`
	CreatedAt                                 string                      `pulumi:"createdAt"`
	DeadLetteringOnFilterEvaluationExceptions *bool                       `pulumi:"deadLetteringOnFilterEvaluationExceptions"`
	DeadLetteringOnMessageExpiration          *bool                       `pulumi:"deadLetteringOnMessageExpiration"`
	DefaultMessageTimeToLive                  *string                     `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow       *string                     `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations                   *bool                       `pulumi:"enableBatchedOperations"`
	ForwardDeadLetteredMessagesTo             *string                     `pulumi:"forwardDeadLetteredMessagesTo"`
	ForwardTo                                 *string                     `pulumi:"forwardTo"`
	Id                                        string                      `pulumi:"id"`
	LockDuration                              *string                     `pulumi:"lockDuration"`
	MaxDeliveryCount                          *int                        `pulumi:"maxDeliveryCount"`
	MessageCount                              float64                     `pulumi:"messageCount"`
	Name                                      string                      `pulumi:"name"`
	RequiresSession                           *bool                       `pulumi:"requiresSession"`
	Status                                    *string                     `pulumi:"status"`
	Type                                      string                      `pulumi:"type"`
	UpdatedAt                                 string                      `pulumi:"updatedAt"`
}

func LookupSubscriptionOutput(ctx *pulumi.Context, args LookupSubscriptionOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionResult, error) {
			args := v.(LookupSubscriptionArgs)
			r, err := LookupSubscription(ctx, &args, opts...)
			var s LookupSubscriptionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionResultOutput)
}

type LookupSubscriptionOutputArgs struct {
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SubscriptionName  pulumi.StringInput `pulumi:"subscriptionName"`
	TopicName         pulumi.StringInput `pulumi:"topicName"`
}

func (LookupSubscriptionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionArgs)(nil)).Elem()
}


type LookupSubscriptionResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionResult)(nil)).Elem()
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutput() LookupSubscriptionResultOutput {
	return o
}

func (o LookupSubscriptionResultOutput) ToLookupSubscriptionResultOutputWithContext(ctx context.Context) LookupSubscriptionResultOutput {
	return o
}

func (o LookupSubscriptionResultOutput) AccessedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.AccessedAt }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) AutoDeleteOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.AutoDeleteOnIdle }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) CountDetails() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) MessageCountDetailsResponse { return v.CountDetails }).(MessageCountDetailsResponseOutput)
}

func (o LookupSubscriptionResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) DeadLetteringOnFilterEvaluationExceptions() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *bool { return v.DeadLetteringOnFilterEvaluationExceptions }).(pulumi.BoolPtrOutput)
}

func (o LookupSubscriptionResultOutput) DeadLetteringOnMessageExpiration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *bool { return v.DeadLetteringOnMessageExpiration }).(pulumi.BoolPtrOutput)
}

func (o LookupSubscriptionResultOutput) DefaultMessageTimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.DefaultMessageTimeToLive }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) DuplicateDetectionHistoryTimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.DuplicateDetectionHistoryTimeWindow }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) EnableBatchedOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *bool { return v.EnableBatchedOperations }).(pulumi.BoolPtrOutput)
}

func (o LookupSubscriptionResultOutput) ForwardDeadLetteredMessagesTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.ForwardDeadLetteredMessagesTo }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) ForwardTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.ForwardTo }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) LockDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.LockDuration }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o LookupSubscriptionResultOutput) MessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupSubscriptionResult) float64 { return v.MessageCount }).(pulumi.Float64Output)
}

func (o LookupSubscriptionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) RequiresSession() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *bool { return v.RequiresSession }).(pulumi.BoolPtrOutput)
}

func (o LookupSubscriptionResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSubscriptionResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionResultOutput{})
}
