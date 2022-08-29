


package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupQueue(ctx *pulumi.Context, args *LookupQueueArgs, opts ...pulumi.InvokeOption) (*LookupQueueResult, error) {
	var rv LookupQueueResult
	err := ctx.Invoke("azure-native:servicebus:getQueue", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupQueueArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	QueueName         string `pulumi:"queueName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupQueueResult struct {
	AccessedAt                          string                      `pulumi:"accessedAt"`
	AutoDeleteOnIdle                    *string                     `pulumi:"autoDeleteOnIdle"`
	CountDetails                        MessageCountDetailsResponse `pulumi:"countDetails"`
	CreatedAt                           string                      `pulumi:"createdAt"`
	DeadLetteringOnMessageExpiration    *bool                       `pulumi:"deadLetteringOnMessageExpiration"`
	DefaultMessageTimeToLive            *string                     `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow *string                     `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations             *bool                       `pulumi:"enableBatchedOperations"`
	EnableExpress                       *bool                       `pulumi:"enableExpress"`
	EnablePartitioning                  *bool                       `pulumi:"enablePartitioning"`
	ForwardDeadLetteredMessagesTo       *string                     `pulumi:"forwardDeadLetteredMessagesTo"`
	ForwardTo                           *string                     `pulumi:"forwardTo"`
	Id                                  string                      `pulumi:"id"`
	LockDuration                        *string                     `pulumi:"lockDuration"`
	MaxDeliveryCount                    *int                        `pulumi:"maxDeliveryCount"`
	MaxSizeInMegabytes                  *int                        `pulumi:"maxSizeInMegabytes"`
	MessageCount                        float64                     `pulumi:"messageCount"`
	Name                                string                      `pulumi:"name"`
	RequiresDuplicateDetection          *bool                       `pulumi:"requiresDuplicateDetection"`
	RequiresSession                     *bool                       `pulumi:"requiresSession"`
	SizeInBytes                         float64                     `pulumi:"sizeInBytes"`
	Status                              *string                     `pulumi:"status"`
	Type                                string                      `pulumi:"type"`
	UpdatedAt                           string                      `pulumi:"updatedAt"`
}

func LookupQueueOutput(ctx *pulumi.Context, args LookupQueueOutputArgs, opts ...pulumi.InvokeOption) LookupQueueResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupQueueResult, error) {
			args := v.(LookupQueueArgs)
			r, err := LookupQueue(ctx, &args, opts...)
			var s LookupQueueResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupQueueResultOutput)
}

type LookupQueueOutputArgs struct {
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	QueueName         pulumi.StringInput `pulumi:"queueName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupQueueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueArgs)(nil)).Elem()
}


type LookupQueueResultOutput struct{ *pulumi.OutputState }

func (LookupQueueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupQueueResult)(nil)).Elem()
}

func (o LookupQueueResultOutput) ToLookupQueueResultOutput() LookupQueueResultOutput {
	return o
}

func (o LookupQueueResultOutput) ToLookupQueueResultOutputWithContext(ctx context.Context) LookupQueueResultOutput {
	return o
}

func (o LookupQueueResultOutput) AccessedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.AccessedAt }).(pulumi.StringOutput)
}

func (o LookupQueueResultOutput) AutoDeleteOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.AutoDeleteOnIdle }).(pulumi.StringPtrOutput)
}

func (o LookupQueueResultOutput) CountDetails() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v LookupQueueResult) MessageCountDetailsResponse { return v.CountDetails }).(MessageCountDetailsResponseOutput)
}

func (o LookupQueueResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupQueueResultOutput) DeadLetteringOnMessageExpiration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *bool { return v.DeadLetteringOnMessageExpiration }).(pulumi.BoolPtrOutput)
}

func (o LookupQueueResultOutput) DefaultMessageTimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.DefaultMessageTimeToLive }).(pulumi.StringPtrOutput)
}

func (o LookupQueueResultOutput) DuplicateDetectionHistoryTimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.DuplicateDetectionHistoryTimeWindow }).(pulumi.StringPtrOutput)
}

func (o LookupQueueResultOutput) EnableBatchedOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *bool { return v.EnableBatchedOperations }).(pulumi.BoolPtrOutput)
}

func (o LookupQueueResultOutput) EnableExpress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *bool { return v.EnableExpress }).(pulumi.BoolPtrOutput)
}

func (o LookupQueueResultOutput) EnablePartitioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *bool { return v.EnablePartitioning }).(pulumi.BoolPtrOutput)
}

func (o LookupQueueResultOutput) ForwardDeadLetteredMessagesTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.ForwardDeadLetteredMessagesTo }).(pulumi.StringPtrOutput)
}

func (o LookupQueueResultOutput) ForwardTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.ForwardTo }).(pulumi.StringPtrOutput)
}

func (o LookupQueueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupQueueResultOutput) LockDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.LockDuration }).(pulumi.StringPtrOutput)
}

func (o LookupQueueResultOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o LookupQueueResultOutput) MaxSizeInMegabytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *int { return v.MaxSizeInMegabytes }).(pulumi.IntPtrOutput)
}

func (o LookupQueueResultOutput) MessageCount() pulumi.Float64Output {
	return o.ApplyT(func(v LookupQueueResult) float64 { return v.MessageCount }).(pulumi.Float64Output)
}

func (o LookupQueueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupQueueResultOutput) RequiresDuplicateDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *bool { return v.RequiresDuplicateDetection }).(pulumi.BoolPtrOutput)
}

func (o LookupQueueResultOutput) RequiresSession() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *bool { return v.RequiresSession }).(pulumi.BoolPtrOutput)
}

func (o LookupQueueResultOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupQueueResult) float64 { return v.SizeInBytes }).(pulumi.Float64Output)
}

func (o LookupQueueResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupQueueResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupQueueResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupQueueResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupQueueResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupQueueResultOutput{})
}
