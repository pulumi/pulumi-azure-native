


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	var rv LookupTopicResult
	err := ctx.Invoke("azure-native:servicebus/v20150801:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicArgs struct {
	NamespaceName     string `pulumi:"namespaceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	TopicName         string `pulumi:"topicName"`
}


type LookupTopicResult struct {
	AccessedAt                          string                      `pulumi:"accessedAt"`
	AutoDeleteOnIdle                    *string                     `pulumi:"autoDeleteOnIdle"`
	CountDetails                        MessageCountDetailsResponse `pulumi:"countDetails"`
	CreatedAt                           string                      `pulumi:"createdAt"`
	DefaultMessageTimeToLive            *string                     `pulumi:"defaultMessageTimeToLive"`
	DuplicateDetectionHistoryTimeWindow *string                     `pulumi:"duplicateDetectionHistoryTimeWindow"`
	EnableBatchedOperations             *bool                       `pulumi:"enableBatchedOperations"`
	EnableExpress                       *bool                       `pulumi:"enableExpress"`
	EnablePartitioning                  *bool                       `pulumi:"enablePartitioning"`
	EntityAvailabilityStatus            *string                     `pulumi:"entityAvailabilityStatus"`
	FilteringMessagesBeforePublishing   *bool                       `pulumi:"filteringMessagesBeforePublishing"`
	Id                                  string                      `pulumi:"id"`
	IsAnonymousAccessible               *bool                       `pulumi:"isAnonymousAccessible"`
	IsExpress                           *bool                       `pulumi:"isExpress"`
	Location                            *string                     `pulumi:"location"`
	MaxSizeInMegabytes                  *float64                    `pulumi:"maxSizeInMegabytes"`
	Name                                string                      `pulumi:"name"`
	RequiresDuplicateDetection          *bool                       `pulumi:"requiresDuplicateDetection"`
	SizeInBytes                         float64                     `pulumi:"sizeInBytes"`
	Status                              *string                     `pulumi:"status"`
	SubscriptionCount                   int                         `pulumi:"subscriptionCount"`
	SupportOrdering                     *bool                       `pulumi:"supportOrdering"`
	Type                                string                      `pulumi:"type"`
	UpdatedAt                           string                      `pulumi:"updatedAt"`
}

func LookupTopicOutput(ctx *pulumi.Context, args LookupTopicOutputArgs, opts ...pulumi.InvokeOption) LookupTopicResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTopicResult, error) {
			args := v.(LookupTopicArgs)
			r, err := LookupTopic(ctx, &args, opts...)
			var s LookupTopicResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTopicResultOutput)
}

type LookupTopicOutputArgs struct {
	NamespaceName     pulumi.StringInput `pulumi:"namespaceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	TopicName         pulumi.StringInput `pulumi:"topicName"`
}

func (LookupTopicOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicArgs)(nil)).Elem()
}


type LookupTopicResultOutput struct{ *pulumi.OutputState }

func (LookupTopicResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTopicResult)(nil)).Elem()
}

func (o LookupTopicResultOutput) ToLookupTopicResultOutput() LookupTopicResultOutput {
	return o
}

func (o LookupTopicResultOutput) ToLookupTopicResultOutputWithContext(ctx context.Context) LookupTopicResultOutput {
	return o
}

func (o LookupTopicResultOutput) AccessedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.AccessedAt }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) AutoDeleteOnIdle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.AutoDeleteOnIdle }).(pulumi.StringPtrOutput)
}

func (o LookupTopicResultOutput) CountDetails() MessageCountDetailsResponseOutput {
	return o.ApplyT(func(v LookupTopicResult) MessageCountDetailsResponse { return v.CountDetails }).(MessageCountDetailsResponseOutput)
}

func (o LookupTopicResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) DefaultMessageTimeToLive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.DefaultMessageTimeToLive }).(pulumi.StringPtrOutput)
}

func (o LookupTopicResultOutput) DuplicateDetectionHistoryTimeWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.DuplicateDetectionHistoryTimeWindow }).(pulumi.StringPtrOutput)
}

func (o LookupTopicResultOutput) EnableBatchedOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.EnableBatchedOperations }).(pulumi.BoolPtrOutput)
}

func (o LookupTopicResultOutput) EnableExpress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.EnableExpress }).(pulumi.BoolPtrOutput)
}

func (o LookupTopicResultOutput) EnablePartitioning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.EnablePartitioning }).(pulumi.BoolPtrOutput)
}

func (o LookupTopicResultOutput) EntityAvailabilityStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.EntityAvailabilityStatus }).(pulumi.StringPtrOutput)
}

func (o LookupTopicResultOutput) FilteringMessagesBeforePublishing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.FilteringMessagesBeforePublishing }).(pulumi.BoolPtrOutput)
}

func (o LookupTopicResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) IsAnonymousAccessible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.IsAnonymousAccessible }).(pulumi.BoolPtrOutput)
}

func (o LookupTopicResultOutput) IsExpress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.IsExpress }).(pulumi.BoolPtrOutput)
}

func (o LookupTopicResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupTopicResultOutput) MaxSizeInMegabytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *float64 { return v.MaxSizeInMegabytes }).(pulumi.Float64PtrOutput)
}

func (o LookupTopicResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) RequiresDuplicateDetection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.RequiresDuplicateDetection }).(pulumi.BoolPtrOutput)
}

func (o LookupTopicResultOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v LookupTopicResult) float64 { return v.SizeInBytes }).(pulumi.Float64Output)
}

func (o LookupTopicResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o LookupTopicResultOutput) SubscriptionCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTopicResult) int { return v.SubscriptionCount }).(pulumi.IntOutput)
}

func (o LookupTopicResultOutput) SupportOrdering() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupTopicResult) *bool { return v.SupportOrdering }).(pulumi.BoolPtrOutput)
}

func (o LookupTopicResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupTopicResultOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTopicResult) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTopicResultOutput{})
}
