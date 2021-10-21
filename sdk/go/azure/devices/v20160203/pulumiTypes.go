


package v20160203

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CloudToDeviceProperties struct {
	DefaultTtlAsIso8601 *string             `pulumi:"defaultTtlAsIso8601"`
	Feedback            *FeedbackProperties `pulumi:"feedback"`
	MaxDeliveryCount    *int                `pulumi:"maxDeliveryCount"`
}





type CloudToDevicePropertiesInput interface {
	pulumi.Input

	ToCloudToDevicePropertiesOutput() CloudToDevicePropertiesOutput
	ToCloudToDevicePropertiesOutputWithContext(context.Context) CloudToDevicePropertiesOutput
}

type CloudToDevicePropertiesArgs struct {
	DefaultTtlAsIso8601 pulumi.StringPtrInput      `pulumi:"defaultTtlAsIso8601"`
	Feedback            FeedbackPropertiesPtrInput `pulumi:"feedback"`
	MaxDeliveryCount    pulumi.IntPtrInput         `pulumi:"maxDeliveryCount"`
}

func (CloudToDevicePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudToDeviceProperties)(nil)).Elem()
}

func (i CloudToDevicePropertiesArgs) ToCloudToDevicePropertiesOutput() CloudToDevicePropertiesOutput {
	return i.ToCloudToDevicePropertiesOutputWithContext(context.Background())
}

func (i CloudToDevicePropertiesArgs) ToCloudToDevicePropertiesOutputWithContext(ctx context.Context) CloudToDevicePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudToDevicePropertiesOutput)
}

func (i CloudToDevicePropertiesArgs) ToCloudToDevicePropertiesPtrOutput() CloudToDevicePropertiesPtrOutput {
	return i.ToCloudToDevicePropertiesPtrOutputWithContext(context.Background())
}

func (i CloudToDevicePropertiesArgs) ToCloudToDevicePropertiesPtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudToDevicePropertiesOutput).ToCloudToDevicePropertiesPtrOutputWithContext(ctx)
}









type CloudToDevicePropertiesPtrInput interface {
	pulumi.Input

	ToCloudToDevicePropertiesPtrOutput() CloudToDevicePropertiesPtrOutput
	ToCloudToDevicePropertiesPtrOutputWithContext(context.Context) CloudToDevicePropertiesPtrOutput
}

type cloudToDevicePropertiesPtrType CloudToDevicePropertiesArgs

func CloudToDevicePropertiesPtr(v *CloudToDevicePropertiesArgs) CloudToDevicePropertiesPtrInput {
	return (*cloudToDevicePropertiesPtrType)(v)
}

func (*cloudToDevicePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudToDeviceProperties)(nil)).Elem()
}

func (i *cloudToDevicePropertiesPtrType) ToCloudToDevicePropertiesPtrOutput() CloudToDevicePropertiesPtrOutput {
	return i.ToCloudToDevicePropertiesPtrOutputWithContext(context.Background())
}

func (i *cloudToDevicePropertiesPtrType) ToCloudToDevicePropertiesPtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudToDevicePropertiesPtrOutput)
}

type CloudToDevicePropertiesOutput struct{ *pulumi.OutputState }

func (CloudToDevicePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudToDeviceProperties)(nil)).Elem()
}

func (o CloudToDevicePropertiesOutput) ToCloudToDevicePropertiesOutput() CloudToDevicePropertiesOutput {
	return o
}

func (o CloudToDevicePropertiesOutput) ToCloudToDevicePropertiesOutputWithContext(ctx context.Context) CloudToDevicePropertiesOutput {
	return o
}

func (o CloudToDevicePropertiesOutput) ToCloudToDevicePropertiesPtrOutput() CloudToDevicePropertiesPtrOutput {
	return o.ToCloudToDevicePropertiesPtrOutputWithContext(context.Background())
}

func (o CloudToDevicePropertiesOutput) ToCloudToDevicePropertiesPtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudToDeviceProperties) *CloudToDeviceProperties {
		return &v
	}).(CloudToDevicePropertiesPtrOutput)
}

func (o CloudToDevicePropertiesOutput) DefaultTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudToDeviceProperties) *string { return v.DefaultTtlAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o CloudToDevicePropertiesOutput) Feedback() FeedbackPropertiesPtrOutput {
	return o.ApplyT(func(v CloudToDeviceProperties) *FeedbackProperties { return v.Feedback }).(FeedbackPropertiesPtrOutput)
}

func (o CloudToDevicePropertiesOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CloudToDeviceProperties) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

type CloudToDevicePropertiesPtrOutput struct{ *pulumi.OutputState }

func (CloudToDevicePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudToDeviceProperties)(nil)).Elem()
}

func (o CloudToDevicePropertiesPtrOutput) ToCloudToDevicePropertiesPtrOutput() CloudToDevicePropertiesPtrOutput {
	return o
}

func (o CloudToDevicePropertiesPtrOutput) ToCloudToDevicePropertiesPtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesPtrOutput {
	return o
}

func (o CloudToDevicePropertiesPtrOutput) Elem() CloudToDevicePropertiesOutput {
	return o.ApplyT(func(v *CloudToDeviceProperties) CloudToDeviceProperties {
		if v != nil {
			return *v
		}
		var ret CloudToDeviceProperties
		return ret
	}).(CloudToDevicePropertiesOutput)
}

func (o CloudToDevicePropertiesPtrOutput) DefaultTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudToDeviceProperties) *string {
		if v == nil {
			return nil
		}
		return v.DefaultTtlAsIso8601
	}).(pulumi.StringPtrOutput)
}

func (o CloudToDevicePropertiesPtrOutput) Feedback() FeedbackPropertiesPtrOutput {
	return o.ApplyT(func(v *CloudToDeviceProperties) *FeedbackProperties {
		if v == nil {
			return nil
		}
		return v.Feedback
	}).(FeedbackPropertiesPtrOutput)
}

func (o CloudToDevicePropertiesPtrOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudToDeviceProperties) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryCount
	}).(pulumi.IntPtrOutput)
}

type CloudToDevicePropertiesResponse struct {
	DefaultTtlAsIso8601 *string                     `pulumi:"defaultTtlAsIso8601"`
	Feedback            *FeedbackPropertiesResponse `pulumi:"feedback"`
	MaxDeliveryCount    *int                        `pulumi:"maxDeliveryCount"`
}





type CloudToDevicePropertiesResponseInput interface {
	pulumi.Input

	ToCloudToDevicePropertiesResponseOutput() CloudToDevicePropertiesResponseOutput
	ToCloudToDevicePropertiesResponseOutputWithContext(context.Context) CloudToDevicePropertiesResponseOutput
}

type CloudToDevicePropertiesResponseArgs struct {
	DefaultTtlAsIso8601 pulumi.StringPtrInput              `pulumi:"defaultTtlAsIso8601"`
	Feedback            FeedbackPropertiesResponsePtrInput `pulumi:"feedback"`
	MaxDeliveryCount    pulumi.IntPtrInput                 `pulumi:"maxDeliveryCount"`
}

func (CloudToDevicePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudToDevicePropertiesResponse)(nil)).Elem()
}

func (i CloudToDevicePropertiesResponseArgs) ToCloudToDevicePropertiesResponseOutput() CloudToDevicePropertiesResponseOutput {
	return i.ToCloudToDevicePropertiesResponseOutputWithContext(context.Background())
}

func (i CloudToDevicePropertiesResponseArgs) ToCloudToDevicePropertiesResponseOutputWithContext(ctx context.Context) CloudToDevicePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudToDevicePropertiesResponseOutput)
}

func (i CloudToDevicePropertiesResponseArgs) ToCloudToDevicePropertiesResponsePtrOutput() CloudToDevicePropertiesResponsePtrOutput {
	return i.ToCloudToDevicePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CloudToDevicePropertiesResponseArgs) ToCloudToDevicePropertiesResponsePtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudToDevicePropertiesResponseOutput).ToCloudToDevicePropertiesResponsePtrOutputWithContext(ctx)
}









type CloudToDevicePropertiesResponsePtrInput interface {
	pulumi.Input

	ToCloudToDevicePropertiesResponsePtrOutput() CloudToDevicePropertiesResponsePtrOutput
	ToCloudToDevicePropertiesResponsePtrOutputWithContext(context.Context) CloudToDevicePropertiesResponsePtrOutput
}

type cloudToDevicePropertiesResponsePtrType CloudToDevicePropertiesResponseArgs

func CloudToDevicePropertiesResponsePtr(v *CloudToDevicePropertiesResponseArgs) CloudToDevicePropertiesResponsePtrInput {
	return (*cloudToDevicePropertiesResponsePtrType)(v)
}

func (*cloudToDevicePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudToDevicePropertiesResponse)(nil)).Elem()
}

func (i *cloudToDevicePropertiesResponsePtrType) ToCloudToDevicePropertiesResponsePtrOutput() CloudToDevicePropertiesResponsePtrOutput {
	return i.ToCloudToDevicePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *cloudToDevicePropertiesResponsePtrType) ToCloudToDevicePropertiesResponsePtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudToDevicePropertiesResponsePtrOutput)
}

type CloudToDevicePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CloudToDevicePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloudToDevicePropertiesResponse)(nil)).Elem()
}

func (o CloudToDevicePropertiesResponseOutput) ToCloudToDevicePropertiesResponseOutput() CloudToDevicePropertiesResponseOutput {
	return o
}

func (o CloudToDevicePropertiesResponseOutput) ToCloudToDevicePropertiesResponseOutputWithContext(ctx context.Context) CloudToDevicePropertiesResponseOutput {
	return o
}

func (o CloudToDevicePropertiesResponseOutput) ToCloudToDevicePropertiesResponsePtrOutput() CloudToDevicePropertiesResponsePtrOutput {
	return o.ToCloudToDevicePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CloudToDevicePropertiesResponseOutput) ToCloudToDevicePropertiesResponsePtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloudToDevicePropertiesResponse) *CloudToDevicePropertiesResponse {
		return &v
	}).(CloudToDevicePropertiesResponsePtrOutput)
}

func (o CloudToDevicePropertiesResponseOutput) DefaultTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloudToDevicePropertiesResponse) *string { return v.DefaultTtlAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o CloudToDevicePropertiesResponseOutput) Feedback() FeedbackPropertiesResponsePtrOutput {
	return o.ApplyT(func(v CloudToDevicePropertiesResponse) *FeedbackPropertiesResponse { return v.Feedback }).(FeedbackPropertiesResponsePtrOutput)
}

func (o CloudToDevicePropertiesResponseOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CloudToDevicePropertiesResponse) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

type CloudToDevicePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CloudToDevicePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudToDevicePropertiesResponse)(nil)).Elem()
}

func (o CloudToDevicePropertiesResponsePtrOutput) ToCloudToDevicePropertiesResponsePtrOutput() CloudToDevicePropertiesResponsePtrOutput {
	return o
}

func (o CloudToDevicePropertiesResponsePtrOutput) ToCloudToDevicePropertiesResponsePtrOutputWithContext(ctx context.Context) CloudToDevicePropertiesResponsePtrOutput {
	return o
}

func (o CloudToDevicePropertiesResponsePtrOutput) Elem() CloudToDevicePropertiesResponseOutput {
	return o.ApplyT(func(v *CloudToDevicePropertiesResponse) CloudToDevicePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CloudToDevicePropertiesResponse
		return ret
	}).(CloudToDevicePropertiesResponseOutput)
}

func (o CloudToDevicePropertiesResponsePtrOutput) DefaultTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudToDevicePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultTtlAsIso8601
	}).(pulumi.StringPtrOutput)
}

func (o CloudToDevicePropertiesResponsePtrOutput) Feedback() FeedbackPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CloudToDevicePropertiesResponse) *FeedbackPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Feedback
	}).(FeedbackPropertiesResponsePtrOutput)
}

func (o CloudToDevicePropertiesResponsePtrOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CloudToDevicePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryCount
	}).(pulumi.IntPtrOutput)
}

type EventHubProperties struct {
	PartitionCount      *int     `pulumi:"partitionCount"`
	RetentionTimeInDays *float64 `pulumi:"retentionTimeInDays"`
}





type EventHubPropertiesInput interface {
	pulumi.Input

	ToEventHubPropertiesOutput() EventHubPropertiesOutput
	ToEventHubPropertiesOutputWithContext(context.Context) EventHubPropertiesOutput
}

type EventHubPropertiesArgs struct {
	PartitionCount      pulumi.IntPtrInput     `pulumi:"partitionCount"`
	RetentionTimeInDays pulumi.Float64PtrInput `pulumi:"retentionTimeInDays"`
}

func (EventHubPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubProperties)(nil)).Elem()
}

func (i EventHubPropertiesArgs) ToEventHubPropertiesOutput() EventHubPropertiesOutput {
	return i.ToEventHubPropertiesOutputWithContext(context.Background())
}

func (i EventHubPropertiesArgs) ToEventHubPropertiesOutputWithContext(ctx context.Context) EventHubPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubPropertiesOutput)
}





type EventHubPropertiesMapInput interface {
	pulumi.Input

	ToEventHubPropertiesMapOutput() EventHubPropertiesMapOutput
	ToEventHubPropertiesMapOutputWithContext(context.Context) EventHubPropertiesMapOutput
}

type EventHubPropertiesMap map[string]EventHubPropertiesInput

func (EventHubPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventHubProperties)(nil)).Elem()
}

func (i EventHubPropertiesMap) ToEventHubPropertiesMapOutput() EventHubPropertiesMapOutput {
	return i.ToEventHubPropertiesMapOutputWithContext(context.Background())
}

func (i EventHubPropertiesMap) ToEventHubPropertiesMapOutputWithContext(ctx context.Context) EventHubPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubPropertiesMapOutput)
}

type EventHubPropertiesOutput struct{ *pulumi.OutputState }

func (EventHubPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubProperties)(nil)).Elem()
}

func (o EventHubPropertiesOutput) ToEventHubPropertiesOutput() EventHubPropertiesOutput {
	return o
}

func (o EventHubPropertiesOutput) ToEventHubPropertiesOutputWithContext(ctx context.Context) EventHubPropertiesOutput {
	return o
}

func (o EventHubPropertiesOutput) PartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EventHubProperties) *int { return v.PartitionCount }).(pulumi.IntPtrOutput)
}

func (o EventHubPropertiesOutput) RetentionTimeInDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EventHubProperties) *float64 { return v.RetentionTimeInDays }).(pulumi.Float64PtrOutput)
}

type EventHubPropertiesMapOutput struct{ *pulumi.OutputState }

func (EventHubPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventHubProperties)(nil)).Elem()
}

func (o EventHubPropertiesMapOutput) ToEventHubPropertiesMapOutput() EventHubPropertiesMapOutput {
	return o
}

func (o EventHubPropertiesMapOutput) ToEventHubPropertiesMapOutputWithContext(ctx context.Context) EventHubPropertiesMapOutput {
	return o
}

func (o EventHubPropertiesMapOutput) MapIndex(k pulumi.StringInput) EventHubPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventHubProperties {
		return vs[0].(map[string]EventHubProperties)[vs[1].(string)]
	}).(EventHubPropertiesOutput)
}

type EventHubPropertiesResponse struct {
	Endpoint            string   `pulumi:"endpoint"`
	PartitionCount      *int     `pulumi:"partitionCount"`
	PartitionIds        []string `pulumi:"partitionIds"`
	Path                string   `pulumi:"path"`
	RetentionTimeInDays *float64 `pulumi:"retentionTimeInDays"`
}





type EventHubPropertiesResponseInput interface {
	pulumi.Input

	ToEventHubPropertiesResponseOutput() EventHubPropertiesResponseOutput
	ToEventHubPropertiesResponseOutputWithContext(context.Context) EventHubPropertiesResponseOutput
}

type EventHubPropertiesResponseArgs struct {
	Endpoint            pulumi.StringInput      `pulumi:"endpoint"`
	PartitionCount      pulumi.IntPtrInput      `pulumi:"partitionCount"`
	PartitionIds        pulumi.StringArrayInput `pulumi:"partitionIds"`
	Path                pulumi.StringInput      `pulumi:"path"`
	RetentionTimeInDays pulumi.Float64PtrInput  `pulumi:"retentionTimeInDays"`
}

func (EventHubPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubPropertiesResponse)(nil)).Elem()
}

func (i EventHubPropertiesResponseArgs) ToEventHubPropertiesResponseOutput() EventHubPropertiesResponseOutput {
	return i.ToEventHubPropertiesResponseOutputWithContext(context.Background())
}

func (i EventHubPropertiesResponseArgs) ToEventHubPropertiesResponseOutputWithContext(ctx context.Context) EventHubPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubPropertiesResponseOutput)
}





type EventHubPropertiesResponseMapInput interface {
	pulumi.Input

	ToEventHubPropertiesResponseMapOutput() EventHubPropertiesResponseMapOutput
	ToEventHubPropertiesResponseMapOutputWithContext(context.Context) EventHubPropertiesResponseMapOutput
}

type EventHubPropertiesResponseMap map[string]EventHubPropertiesResponseInput

func (EventHubPropertiesResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventHubPropertiesResponse)(nil)).Elem()
}

func (i EventHubPropertiesResponseMap) ToEventHubPropertiesResponseMapOutput() EventHubPropertiesResponseMapOutput {
	return i.ToEventHubPropertiesResponseMapOutputWithContext(context.Background())
}

func (i EventHubPropertiesResponseMap) ToEventHubPropertiesResponseMapOutputWithContext(ctx context.Context) EventHubPropertiesResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubPropertiesResponseMapOutput)
}

type EventHubPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EventHubPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubPropertiesResponse)(nil)).Elem()
}

func (o EventHubPropertiesResponseOutput) ToEventHubPropertiesResponseOutput() EventHubPropertiesResponseOutput {
	return o
}

func (o EventHubPropertiesResponseOutput) ToEventHubPropertiesResponseOutputWithContext(ctx context.Context) EventHubPropertiesResponseOutput {
	return o
}

func (o EventHubPropertiesResponseOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubPropertiesResponse) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o EventHubPropertiesResponseOutput) PartitionCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EventHubPropertiesResponse) *int { return v.PartitionCount }).(pulumi.IntPtrOutput)
}

func (o EventHubPropertiesResponseOutput) PartitionIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventHubPropertiesResponse) []string { return v.PartitionIds }).(pulumi.StringArrayOutput)
}

func (o EventHubPropertiesResponseOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubPropertiesResponse) string { return v.Path }).(pulumi.StringOutput)
}

func (o EventHubPropertiesResponseOutput) RetentionTimeInDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v EventHubPropertiesResponse) *float64 { return v.RetentionTimeInDays }).(pulumi.Float64PtrOutput)
}

type EventHubPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (EventHubPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]EventHubPropertiesResponse)(nil)).Elem()
}

func (o EventHubPropertiesResponseMapOutput) ToEventHubPropertiesResponseMapOutput() EventHubPropertiesResponseMapOutput {
	return o
}

func (o EventHubPropertiesResponseMapOutput) ToEventHubPropertiesResponseMapOutputWithContext(ctx context.Context) EventHubPropertiesResponseMapOutput {
	return o
}

func (o EventHubPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) EventHubPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) EventHubPropertiesResponse {
		return vs[0].(map[string]EventHubPropertiesResponse)[vs[1].(string)]
	}).(EventHubPropertiesResponseOutput)
}

type FeedbackProperties struct {
	LockDurationAsIso8601 *string `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      *int    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          *string `pulumi:"ttlAsIso8601"`
}





type FeedbackPropertiesInput interface {
	pulumi.Input

	ToFeedbackPropertiesOutput() FeedbackPropertiesOutput
	ToFeedbackPropertiesOutputWithContext(context.Context) FeedbackPropertiesOutput
}

type FeedbackPropertiesArgs struct {
	LockDurationAsIso8601 pulumi.StringPtrInput `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      pulumi.IntPtrInput    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          pulumi.StringPtrInput `pulumi:"ttlAsIso8601"`
}

func (FeedbackPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FeedbackProperties)(nil)).Elem()
}

func (i FeedbackPropertiesArgs) ToFeedbackPropertiesOutput() FeedbackPropertiesOutput {
	return i.ToFeedbackPropertiesOutputWithContext(context.Background())
}

func (i FeedbackPropertiesArgs) ToFeedbackPropertiesOutputWithContext(ctx context.Context) FeedbackPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedbackPropertiesOutput)
}

func (i FeedbackPropertiesArgs) ToFeedbackPropertiesPtrOutput() FeedbackPropertiesPtrOutput {
	return i.ToFeedbackPropertiesPtrOutputWithContext(context.Background())
}

func (i FeedbackPropertiesArgs) ToFeedbackPropertiesPtrOutputWithContext(ctx context.Context) FeedbackPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedbackPropertiesOutput).ToFeedbackPropertiesPtrOutputWithContext(ctx)
}









type FeedbackPropertiesPtrInput interface {
	pulumi.Input

	ToFeedbackPropertiesPtrOutput() FeedbackPropertiesPtrOutput
	ToFeedbackPropertiesPtrOutputWithContext(context.Context) FeedbackPropertiesPtrOutput
}

type feedbackPropertiesPtrType FeedbackPropertiesArgs

func FeedbackPropertiesPtr(v *FeedbackPropertiesArgs) FeedbackPropertiesPtrInput {
	return (*feedbackPropertiesPtrType)(v)
}

func (*feedbackPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FeedbackProperties)(nil)).Elem()
}

func (i *feedbackPropertiesPtrType) ToFeedbackPropertiesPtrOutput() FeedbackPropertiesPtrOutput {
	return i.ToFeedbackPropertiesPtrOutputWithContext(context.Background())
}

func (i *feedbackPropertiesPtrType) ToFeedbackPropertiesPtrOutputWithContext(ctx context.Context) FeedbackPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedbackPropertiesPtrOutput)
}

type FeedbackPropertiesOutput struct{ *pulumi.OutputState }

func (FeedbackPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FeedbackProperties)(nil)).Elem()
}

func (o FeedbackPropertiesOutput) ToFeedbackPropertiesOutput() FeedbackPropertiesOutput {
	return o
}

func (o FeedbackPropertiesOutput) ToFeedbackPropertiesOutputWithContext(ctx context.Context) FeedbackPropertiesOutput {
	return o
}

func (o FeedbackPropertiesOutput) ToFeedbackPropertiesPtrOutput() FeedbackPropertiesPtrOutput {
	return o.ToFeedbackPropertiesPtrOutputWithContext(context.Background())
}

func (o FeedbackPropertiesOutput) ToFeedbackPropertiesPtrOutputWithContext(ctx context.Context) FeedbackPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FeedbackProperties) *FeedbackProperties {
		return &v
	}).(FeedbackPropertiesPtrOutput)
}

func (o FeedbackPropertiesOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FeedbackProperties) *string { return v.LockDurationAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o FeedbackPropertiesOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FeedbackProperties) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o FeedbackPropertiesOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FeedbackProperties) *string { return v.TtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type FeedbackPropertiesPtrOutput struct{ *pulumi.OutputState }

func (FeedbackPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeedbackProperties)(nil)).Elem()
}

func (o FeedbackPropertiesPtrOutput) ToFeedbackPropertiesPtrOutput() FeedbackPropertiesPtrOutput {
	return o
}

func (o FeedbackPropertiesPtrOutput) ToFeedbackPropertiesPtrOutputWithContext(ctx context.Context) FeedbackPropertiesPtrOutput {
	return o
}

func (o FeedbackPropertiesPtrOutput) Elem() FeedbackPropertiesOutput {
	return o.ApplyT(func(v *FeedbackProperties) FeedbackProperties {
		if v != nil {
			return *v
		}
		var ret FeedbackProperties
		return ret
	}).(FeedbackPropertiesOutput)
}

func (o FeedbackPropertiesPtrOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeedbackProperties) *string {
		if v == nil {
			return nil
		}
		return v.LockDurationAsIso8601
	}).(pulumi.StringPtrOutput)
}

func (o FeedbackPropertiesPtrOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FeedbackProperties) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryCount
	}).(pulumi.IntPtrOutput)
}

func (o FeedbackPropertiesPtrOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeedbackProperties) *string {
		if v == nil {
			return nil
		}
		return v.TtlAsIso8601
	}).(pulumi.StringPtrOutput)
}

type FeedbackPropertiesResponse struct {
	LockDurationAsIso8601 *string `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      *int    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          *string `pulumi:"ttlAsIso8601"`
}





type FeedbackPropertiesResponseInput interface {
	pulumi.Input

	ToFeedbackPropertiesResponseOutput() FeedbackPropertiesResponseOutput
	ToFeedbackPropertiesResponseOutputWithContext(context.Context) FeedbackPropertiesResponseOutput
}

type FeedbackPropertiesResponseArgs struct {
	LockDurationAsIso8601 pulumi.StringPtrInput `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      pulumi.IntPtrInput    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          pulumi.StringPtrInput `pulumi:"ttlAsIso8601"`
}

func (FeedbackPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FeedbackPropertiesResponse)(nil)).Elem()
}

func (i FeedbackPropertiesResponseArgs) ToFeedbackPropertiesResponseOutput() FeedbackPropertiesResponseOutput {
	return i.ToFeedbackPropertiesResponseOutputWithContext(context.Background())
}

func (i FeedbackPropertiesResponseArgs) ToFeedbackPropertiesResponseOutputWithContext(ctx context.Context) FeedbackPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedbackPropertiesResponseOutput)
}

func (i FeedbackPropertiesResponseArgs) ToFeedbackPropertiesResponsePtrOutput() FeedbackPropertiesResponsePtrOutput {
	return i.ToFeedbackPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i FeedbackPropertiesResponseArgs) ToFeedbackPropertiesResponsePtrOutputWithContext(ctx context.Context) FeedbackPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedbackPropertiesResponseOutput).ToFeedbackPropertiesResponsePtrOutputWithContext(ctx)
}









type FeedbackPropertiesResponsePtrInput interface {
	pulumi.Input

	ToFeedbackPropertiesResponsePtrOutput() FeedbackPropertiesResponsePtrOutput
	ToFeedbackPropertiesResponsePtrOutputWithContext(context.Context) FeedbackPropertiesResponsePtrOutput
}

type feedbackPropertiesResponsePtrType FeedbackPropertiesResponseArgs

func FeedbackPropertiesResponsePtr(v *FeedbackPropertiesResponseArgs) FeedbackPropertiesResponsePtrInput {
	return (*feedbackPropertiesResponsePtrType)(v)
}

func (*feedbackPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FeedbackPropertiesResponse)(nil)).Elem()
}

func (i *feedbackPropertiesResponsePtrType) ToFeedbackPropertiesResponsePtrOutput() FeedbackPropertiesResponsePtrOutput {
	return i.ToFeedbackPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *feedbackPropertiesResponsePtrType) ToFeedbackPropertiesResponsePtrOutputWithContext(ctx context.Context) FeedbackPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeedbackPropertiesResponsePtrOutput)
}

type FeedbackPropertiesResponseOutput struct{ *pulumi.OutputState }

func (FeedbackPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FeedbackPropertiesResponse)(nil)).Elem()
}

func (o FeedbackPropertiesResponseOutput) ToFeedbackPropertiesResponseOutput() FeedbackPropertiesResponseOutput {
	return o
}

func (o FeedbackPropertiesResponseOutput) ToFeedbackPropertiesResponseOutputWithContext(ctx context.Context) FeedbackPropertiesResponseOutput {
	return o
}

func (o FeedbackPropertiesResponseOutput) ToFeedbackPropertiesResponsePtrOutput() FeedbackPropertiesResponsePtrOutput {
	return o.ToFeedbackPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o FeedbackPropertiesResponseOutput) ToFeedbackPropertiesResponsePtrOutputWithContext(ctx context.Context) FeedbackPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FeedbackPropertiesResponse) *FeedbackPropertiesResponse {
		return &v
	}).(FeedbackPropertiesResponsePtrOutput)
}

func (o FeedbackPropertiesResponseOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FeedbackPropertiesResponse) *string { return v.LockDurationAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o FeedbackPropertiesResponseOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FeedbackPropertiesResponse) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o FeedbackPropertiesResponseOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FeedbackPropertiesResponse) *string { return v.TtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type FeedbackPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (FeedbackPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FeedbackPropertiesResponse)(nil)).Elem()
}

func (o FeedbackPropertiesResponsePtrOutput) ToFeedbackPropertiesResponsePtrOutput() FeedbackPropertiesResponsePtrOutput {
	return o
}

func (o FeedbackPropertiesResponsePtrOutput) ToFeedbackPropertiesResponsePtrOutputWithContext(ctx context.Context) FeedbackPropertiesResponsePtrOutput {
	return o
}

func (o FeedbackPropertiesResponsePtrOutput) Elem() FeedbackPropertiesResponseOutput {
	return o.ApplyT(func(v *FeedbackPropertiesResponse) FeedbackPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret FeedbackPropertiesResponse
		return ret
	}).(FeedbackPropertiesResponseOutput)
}

func (o FeedbackPropertiesResponsePtrOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeedbackPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LockDurationAsIso8601
	}).(pulumi.StringPtrOutput)
}

func (o FeedbackPropertiesResponsePtrOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FeedbackPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryCount
	}).(pulumi.IntPtrOutput)
}

func (o FeedbackPropertiesResponsePtrOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FeedbackPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TtlAsIso8601
	}).(pulumi.StringPtrOutput)
}

type IotHubProperties struct {
	AuthorizationPolicies          []SharedAccessSignatureAuthorizationRule `pulumi:"authorizationPolicies"`
	CloudToDevice                  *CloudToDeviceProperties                 `pulumi:"cloudToDevice"`
	Comments                       *string                                  `pulumi:"comments"`
	EnableFileUploadNotifications  *bool                                    `pulumi:"enableFileUploadNotifications"`
	EventHubEndpoints              map[string]EventHubProperties            `pulumi:"eventHubEndpoints"`
	Features                       *string                                  `pulumi:"features"`
	IpFilterRules                  []IpFilterRule                           `pulumi:"ipFilterRules"`
	MessagingEndpoints             map[string]MessagingEndpointProperties   `pulumi:"messagingEndpoints"`
	OperationsMonitoringProperties *OperationsMonitoringProperties          `pulumi:"operationsMonitoringProperties"`
	StorageEndpoints               map[string]StorageEndpointProperties     `pulumi:"storageEndpoints"`
}





type IotHubPropertiesInput interface {
	pulumi.Input

	ToIotHubPropertiesOutput() IotHubPropertiesOutput
	ToIotHubPropertiesOutputWithContext(context.Context) IotHubPropertiesOutput
}

type IotHubPropertiesArgs struct {
	AuthorizationPolicies          SharedAccessSignatureAuthorizationRuleArrayInput `pulumi:"authorizationPolicies"`
	CloudToDevice                  CloudToDevicePropertiesPtrInput                  `pulumi:"cloudToDevice"`
	Comments                       pulumi.StringPtrInput                            `pulumi:"comments"`
	EnableFileUploadNotifications  pulumi.BoolPtrInput                              `pulumi:"enableFileUploadNotifications"`
	EventHubEndpoints              EventHubPropertiesMapInput                       `pulumi:"eventHubEndpoints"`
	Features                       pulumi.StringPtrInput                            `pulumi:"features"`
	IpFilterRules                  IpFilterRuleArrayInput                           `pulumi:"ipFilterRules"`
	MessagingEndpoints             MessagingEndpointPropertiesMapInput              `pulumi:"messagingEndpoints"`
	OperationsMonitoringProperties OperationsMonitoringPropertiesPtrInput           `pulumi:"operationsMonitoringProperties"`
	StorageEndpoints               StorageEndpointPropertiesMapInput                `pulumi:"storageEndpoints"`
}

func (IotHubPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubProperties)(nil)).Elem()
}

func (i IotHubPropertiesArgs) ToIotHubPropertiesOutput() IotHubPropertiesOutput {
	return i.ToIotHubPropertiesOutputWithContext(context.Background())
}

func (i IotHubPropertiesArgs) ToIotHubPropertiesOutputWithContext(ctx context.Context) IotHubPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesOutput)
}

func (i IotHubPropertiesArgs) ToIotHubPropertiesPtrOutput() IotHubPropertiesPtrOutput {
	return i.ToIotHubPropertiesPtrOutputWithContext(context.Background())
}

func (i IotHubPropertiesArgs) ToIotHubPropertiesPtrOutputWithContext(ctx context.Context) IotHubPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesOutput).ToIotHubPropertiesPtrOutputWithContext(ctx)
}









type IotHubPropertiesPtrInput interface {
	pulumi.Input

	ToIotHubPropertiesPtrOutput() IotHubPropertiesPtrOutput
	ToIotHubPropertiesPtrOutputWithContext(context.Context) IotHubPropertiesPtrOutput
}

type iotHubPropertiesPtrType IotHubPropertiesArgs

func IotHubPropertiesPtr(v *IotHubPropertiesArgs) IotHubPropertiesPtrInput {
	return (*iotHubPropertiesPtrType)(v)
}

func (*iotHubPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubProperties)(nil)).Elem()
}

func (i *iotHubPropertiesPtrType) ToIotHubPropertiesPtrOutput() IotHubPropertiesPtrOutput {
	return i.ToIotHubPropertiesPtrOutputWithContext(context.Background())
}

func (i *iotHubPropertiesPtrType) ToIotHubPropertiesPtrOutputWithContext(ctx context.Context) IotHubPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesPtrOutput)
}

type IotHubPropertiesOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubProperties)(nil)).Elem()
}

func (o IotHubPropertiesOutput) ToIotHubPropertiesOutput() IotHubPropertiesOutput {
	return o
}

func (o IotHubPropertiesOutput) ToIotHubPropertiesOutputWithContext(ctx context.Context) IotHubPropertiesOutput {
	return o
}

func (o IotHubPropertiesOutput) ToIotHubPropertiesPtrOutput() IotHubPropertiesPtrOutput {
	return o.ToIotHubPropertiesPtrOutputWithContext(context.Background())
}

func (o IotHubPropertiesOutput) ToIotHubPropertiesPtrOutputWithContext(ctx context.Context) IotHubPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotHubProperties) *IotHubProperties {
		return &v
	}).(IotHubPropertiesPtrOutput)
}

func (o IotHubPropertiesOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleArrayOutput {
	return o.ApplyT(func(v IotHubProperties) []SharedAccessSignatureAuthorizationRule { return v.AuthorizationPolicies }).(SharedAccessSignatureAuthorizationRuleArrayOutput)
}

func (o IotHubPropertiesOutput) CloudToDevice() CloudToDevicePropertiesPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *CloudToDeviceProperties { return v.CloudToDevice }).(CloudToDevicePropertiesPtrOutput)
}

func (o IotHubPropertiesOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesOutput) EnableFileUploadNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *bool { return v.EnableFileUploadNotifications }).(pulumi.BoolPtrOutput)
}

func (o IotHubPropertiesOutput) EventHubEndpoints() EventHubPropertiesMapOutput {
	return o.ApplyT(func(v IotHubProperties) map[string]EventHubProperties { return v.EventHubEndpoints }).(EventHubPropertiesMapOutput)
}

func (o IotHubPropertiesOutput) Features() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *string { return v.Features }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesOutput) IpFilterRules() IpFilterRuleArrayOutput {
	return o.ApplyT(func(v IotHubProperties) []IpFilterRule { return v.IpFilterRules }).(IpFilterRuleArrayOutput)
}

func (o IotHubPropertiesOutput) MessagingEndpoints() MessagingEndpointPropertiesMapOutput {
	return o.ApplyT(func(v IotHubProperties) map[string]MessagingEndpointProperties { return v.MessagingEndpoints }).(MessagingEndpointPropertiesMapOutput)
}

func (o IotHubPropertiesOutput) OperationsMonitoringProperties() OperationsMonitoringPropertiesPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *OperationsMonitoringProperties { return v.OperationsMonitoringProperties }).(OperationsMonitoringPropertiesPtrOutput)
}

func (o IotHubPropertiesOutput) StorageEndpoints() StorageEndpointPropertiesMapOutput {
	return o.ApplyT(func(v IotHubProperties) map[string]StorageEndpointProperties { return v.StorageEndpoints }).(StorageEndpointPropertiesMapOutput)
}

type IotHubPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubProperties)(nil)).Elem()
}

func (o IotHubPropertiesPtrOutput) ToIotHubPropertiesPtrOutput() IotHubPropertiesPtrOutput {
	return o
}

func (o IotHubPropertiesPtrOutput) ToIotHubPropertiesPtrOutputWithContext(ctx context.Context) IotHubPropertiesPtrOutput {
	return o
}

func (o IotHubPropertiesPtrOutput) Elem() IotHubPropertiesOutput {
	return o.ApplyT(func(v *IotHubProperties) IotHubProperties {
		if v != nil {
			return *v
		}
		var ret IotHubProperties
		return ret
	}).(IotHubPropertiesOutput)
}

func (o IotHubPropertiesPtrOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleArrayOutput {
	return o.ApplyT(func(v *IotHubProperties) []SharedAccessSignatureAuthorizationRule {
		if v == nil {
			return nil
		}
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleArrayOutput)
}

func (o IotHubPropertiesPtrOutput) CloudToDevice() CloudToDevicePropertiesPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *CloudToDeviceProperties {
		if v == nil {
			return nil
		}
		return v.CloudToDevice
	}).(CloudToDevicePropertiesPtrOutput)
}

func (o IotHubPropertiesPtrOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *string {
		if v == nil {
			return nil
		}
		return v.Comments
	}).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesPtrOutput) EnableFileUploadNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableFileUploadNotifications
	}).(pulumi.BoolPtrOutput)
}

func (o IotHubPropertiesPtrOutput) EventHubEndpoints() EventHubPropertiesMapOutput {
	return o.ApplyT(func(v *IotHubProperties) map[string]EventHubProperties {
		if v == nil {
			return nil
		}
		return v.EventHubEndpoints
	}).(EventHubPropertiesMapOutput)
}

func (o IotHubPropertiesPtrOutput) Features() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *string {
		if v == nil {
			return nil
		}
		return v.Features
	}).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesPtrOutput) IpFilterRules() IpFilterRuleArrayOutput {
	return o.ApplyT(func(v *IotHubProperties) []IpFilterRule {
		if v == nil {
			return nil
		}
		return v.IpFilterRules
	}).(IpFilterRuleArrayOutput)
}

func (o IotHubPropertiesPtrOutput) MessagingEndpoints() MessagingEndpointPropertiesMapOutput {
	return o.ApplyT(func(v *IotHubProperties) map[string]MessagingEndpointProperties {
		if v == nil {
			return nil
		}
		return v.MessagingEndpoints
	}).(MessagingEndpointPropertiesMapOutput)
}

func (o IotHubPropertiesPtrOutput) OperationsMonitoringProperties() OperationsMonitoringPropertiesPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *OperationsMonitoringProperties {
		if v == nil {
			return nil
		}
		return v.OperationsMonitoringProperties
	}).(OperationsMonitoringPropertiesPtrOutput)
}

func (o IotHubPropertiesPtrOutput) StorageEndpoints() StorageEndpointPropertiesMapOutput {
	return o.ApplyT(func(v *IotHubProperties) map[string]StorageEndpointProperties {
		if v == nil {
			return nil
		}
		return v.StorageEndpoints
	}).(StorageEndpointPropertiesMapOutput)
}

type IotHubPropertiesResponse struct {
	AuthorizationPolicies          []SharedAccessSignatureAuthorizationRuleResponse `pulumi:"authorizationPolicies"`
	CloudToDevice                  *CloudToDevicePropertiesResponse                 `pulumi:"cloudToDevice"`
	Comments                       *string                                          `pulumi:"comments"`
	EnableFileUploadNotifications  *bool                                            `pulumi:"enableFileUploadNotifications"`
	EventHubEndpoints              map[string]EventHubPropertiesResponse            `pulumi:"eventHubEndpoints"`
	Features                       *string                                          `pulumi:"features"`
	HostName                       string                                           `pulumi:"hostName"`
	IpFilterRules                  []IpFilterRuleResponse                           `pulumi:"ipFilterRules"`
	MessagingEndpoints             map[string]MessagingEndpointPropertiesResponse   `pulumi:"messagingEndpoints"`
	OperationsMonitoringProperties *OperationsMonitoringPropertiesResponse          `pulumi:"operationsMonitoringProperties"`
	ProvisioningState              string                                           `pulumi:"provisioningState"`
	StorageEndpoints               map[string]StorageEndpointPropertiesResponse     `pulumi:"storageEndpoints"`
}





type IotHubPropertiesResponseInput interface {
	pulumi.Input

	ToIotHubPropertiesResponseOutput() IotHubPropertiesResponseOutput
	ToIotHubPropertiesResponseOutputWithContext(context.Context) IotHubPropertiesResponseOutput
}

type IotHubPropertiesResponseArgs struct {
	AuthorizationPolicies          SharedAccessSignatureAuthorizationRuleResponseArrayInput `pulumi:"authorizationPolicies"`
	CloudToDevice                  CloudToDevicePropertiesResponsePtrInput                  `pulumi:"cloudToDevice"`
	Comments                       pulumi.StringPtrInput                                    `pulumi:"comments"`
	EnableFileUploadNotifications  pulumi.BoolPtrInput                                      `pulumi:"enableFileUploadNotifications"`
	EventHubEndpoints              EventHubPropertiesResponseMapInput                       `pulumi:"eventHubEndpoints"`
	Features                       pulumi.StringPtrInput                                    `pulumi:"features"`
	HostName                       pulumi.StringInput                                       `pulumi:"hostName"`
	IpFilterRules                  IpFilterRuleResponseArrayInput                           `pulumi:"ipFilterRules"`
	MessagingEndpoints             MessagingEndpointPropertiesResponseMapInput              `pulumi:"messagingEndpoints"`
	OperationsMonitoringProperties OperationsMonitoringPropertiesResponsePtrInput           `pulumi:"operationsMonitoringProperties"`
	ProvisioningState              pulumi.StringInput                                       `pulumi:"provisioningState"`
	StorageEndpoints               StorageEndpointPropertiesResponseMapInput                `pulumi:"storageEndpoints"`
}

func (IotHubPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubPropertiesResponse)(nil)).Elem()
}

func (i IotHubPropertiesResponseArgs) ToIotHubPropertiesResponseOutput() IotHubPropertiesResponseOutput {
	return i.ToIotHubPropertiesResponseOutputWithContext(context.Background())
}

func (i IotHubPropertiesResponseArgs) ToIotHubPropertiesResponseOutputWithContext(ctx context.Context) IotHubPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesResponseOutput)
}

func (i IotHubPropertiesResponseArgs) ToIotHubPropertiesResponsePtrOutput() IotHubPropertiesResponsePtrOutput {
	return i.ToIotHubPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i IotHubPropertiesResponseArgs) ToIotHubPropertiesResponsePtrOutputWithContext(ctx context.Context) IotHubPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesResponseOutput).ToIotHubPropertiesResponsePtrOutputWithContext(ctx)
}









type IotHubPropertiesResponsePtrInput interface {
	pulumi.Input

	ToIotHubPropertiesResponsePtrOutput() IotHubPropertiesResponsePtrOutput
	ToIotHubPropertiesResponsePtrOutputWithContext(context.Context) IotHubPropertiesResponsePtrOutput
}

type iotHubPropertiesResponsePtrType IotHubPropertiesResponseArgs

func IotHubPropertiesResponsePtr(v *IotHubPropertiesResponseArgs) IotHubPropertiesResponsePtrInput {
	return (*iotHubPropertiesResponsePtrType)(v)
}

func (*iotHubPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubPropertiesResponse)(nil)).Elem()
}

func (i *iotHubPropertiesResponsePtrType) ToIotHubPropertiesResponsePtrOutput() IotHubPropertiesResponsePtrOutput {
	return i.ToIotHubPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *iotHubPropertiesResponsePtrType) ToIotHubPropertiesResponsePtrOutputWithContext(ctx context.Context) IotHubPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubPropertiesResponsePtrOutput)
}

type IotHubPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubPropertiesResponse)(nil)).Elem()
}

func (o IotHubPropertiesResponseOutput) ToIotHubPropertiesResponseOutput() IotHubPropertiesResponseOutput {
	return o
}

func (o IotHubPropertiesResponseOutput) ToIotHubPropertiesResponseOutputWithContext(ctx context.Context) IotHubPropertiesResponseOutput {
	return o
}

func (o IotHubPropertiesResponseOutput) ToIotHubPropertiesResponsePtrOutput() IotHubPropertiesResponsePtrOutput {
	return o.ToIotHubPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o IotHubPropertiesResponseOutput) ToIotHubPropertiesResponsePtrOutputWithContext(ctx context.Context) IotHubPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotHubPropertiesResponse) *IotHubPropertiesResponse {
		return &v
	}).(IotHubPropertiesResponsePtrOutput)
}

func (o IotHubPropertiesResponseOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) []SharedAccessSignatureAuthorizationRuleResponse {
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleResponseArrayOutput)
}

func (o IotHubPropertiesResponseOutput) CloudToDevice() CloudToDevicePropertiesResponsePtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *CloudToDevicePropertiesResponse { return v.CloudToDevice }).(CloudToDevicePropertiesResponsePtrOutput)
}

func (o IotHubPropertiesResponseOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *string { return v.Comments }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesResponseOutput) EnableFileUploadNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *bool { return v.EnableFileUploadNotifications }).(pulumi.BoolPtrOutput)
}

func (o IotHubPropertiesResponseOutput) EventHubEndpoints() EventHubPropertiesResponseMapOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) map[string]EventHubPropertiesResponse { return v.EventHubEndpoints }).(EventHubPropertiesResponseMapOutput)
}

func (o IotHubPropertiesResponseOutput) Features() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *string { return v.Features }).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesResponseOutput) HostName() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) string { return v.HostName }).(pulumi.StringOutput)
}

func (o IotHubPropertiesResponseOutput) IpFilterRules() IpFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) []IpFilterRuleResponse { return v.IpFilterRules }).(IpFilterRuleResponseArrayOutput)
}

func (o IotHubPropertiesResponseOutput) MessagingEndpoints() MessagingEndpointPropertiesResponseMapOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) map[string]MessagingEndpointPropertiesResponse {
		return v.MessagingEndpoints
	}).(MessagingEndpointPropertiesResponseMapOutput)
}

func (o IotHubPropertiesResponseOutput) OperationsMonitoringProperties() OperationsMonitoringPropertiesResponsePtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *OperationsMonitoringPropertiesResponse {
		return v.OperationsMonitoringProperties
	}).(OperationsMonitoringPropertiesResponsePtrOutput)
}

func (o IotHubPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o IotHubPropertiesResponseOutput) StorageEndpoints() StorageEndpointPropertiesResponseMapOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) map[string]StorageEndpointPropertiesResponse {
		return v.StorageEndpoints
	}).(StorageEndpointPropertiesResponseMapOutput)
}

type IotHubPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (IotHubPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubPropertiesResponse)(nil)).Elem()
}

func (o IotHubPropertiesResponsePtrOutput) ToIotHubPropertiesResponsePtrOutput() IotHubPropertiesResponsePtrOutput {
	return o
}

func (o IotHubPropertiesResponsePtrOutput) ToIotHubPropertiesResponsePtrOutputWithContext(ctx context.Context) IotHubPropertiesResponsePtrOutput {
	return o
}

func (o IotHubPropertiesResponsePtrOutput) Elem() IotHubPropertiesResponseOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) IotHubPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret IotHubPropertiesResponse
		return ret
	}).(IotHubPropertiesResponseOutput)
}

func (o IotHubPropertiesResponsePtrOutput) AuthorizationPolicies() SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) []SharedAccessSignatureAuthorizationRuleResponse {
		if v == nil {
			return nil
		}
		return v.AuthorizationPolicies
	}).(SharedAccessSignatureAuthorizationRuleResponseArrayOutput)
}

func (o IotHubPropertiesResponsePtrOutput) CloudToDevice() CloudToDevicePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) *CloudToDevicePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.CloudToDevice
	}).(CloudToDevicePropertiesResponsePtrOutput)
}

func (o IotHubPropertiesResponsePtrOutput) Comments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Comments
	}).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesResponsePtrOutput) EnableFileUploadNotifications() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableFileUploadNotifications
	}).(pulumi.BoolPtrOutput)
}

func (o IotHubPropertiesResponsePtrOutput) EventHubEndpoints() EventHubPropertiesResponseMapOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) map[string]EventHubPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.EventHubEndpoints
	}).(EventHubPropertiesResponseMapOutput)
}

func (o IotHubPropertiesResponsePtrOutput) Features() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Features
	}).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesResponsePtrOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HostName
	}).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesResponsePtrOutput) IpFilterRules() IpFilterRuleResponseArrayOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) []IpFilterRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpFilterRules
	}).(IpFilterRuleResponseArrayOutput)
}

func (o IotHubPropertiesResponsePtrOutput) MessagingEndpoints() MessagingEndpointPropertiesResponseMapOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) map[string]MessagingEndpointPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.MessagingEndpoints
	}).(MessagingEndpointPropertiesResponseMapOutput)
}

func (o IotHubPropertiesResponsePtrOutput) OperationsMonitoringProperties() OperationsMonitoringPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) *OperationsMonitoringPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.OperationsMonitoringProperties
	}).(OperationsMonitoringPropertiesResponsePtrOutput)
}

func (o IotHubPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o IotHubPropertiesResponsePtrOutput) StorageEndpoints() StorageEndpointPropertiesResponseMapOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) map[string]StorageEndpointPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.StorageEndpoints
	}).(StorageEndpointPropertiesResponseMapOutput)
}

type IotHubSkuInfo struct {
	Capacity float64 `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
}





type IotHubSkuInfoInput interface {
	pulumi.Input

	ToIotHubSkuInfoOutput() IotHubSkuInfoOutput
	ToIotHubSkuInfoOutputWithContext(context.Context) IotHubSkuInfoOutput
}

type IotHubSkuInfoArgs struct {
	Capacity pulumi.Float64Input `pulumi:"capacity"`
	Name     pulumi.StringInput  `pulumi:"name"`
}

func (IotHubSkuInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSkuInfo)(nil)).Elem()
}

func (i IotHubSkuInfoArgs) ToIotHubSkuInfoOutput() IotHubSkuInfoOutput {
	return i.ToIotHubSkuInfoOutputWithContext(context.Background())
}

func (i IotHubSkuInfoArgs) ToIotHubSkuInfoOutputWithContext(ctx context.Context) IotHubSkuInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSkuInfoOutput)
}

func (i IotHubSkuInfoArgs) ToIotHubSkuInfoPtrOutput() IotHubSkuInfoPtrOutput {
	return i.ToIotHubSkuInfoPtrOutputWithContext(context.Background())
}

func (i IotHubSkuInfoArgs) ToIotHubSkuInfoPtrOutputWithContext(ctx context.Context) IotHubSkuInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSkuInfoOutput).ToIotHubSkuInfoPtrOutputWithContext(ctx)
}









type IotHubSkuInfoPtrInput interface {
	pulumi.Input

	ToIotHubSkuInfoPtrOutput() IotHubSkuInfoPtrOutput
	ToIotHubSkuInfoPtrOutputWithContext(context.Context) IotHubSkuInfoPtrOutput
}

type iotHubSkuInfoPtrType IotHubSkuInfoArgs

func IotHubSkuInfoPtr(v *IotHubSkuInfoArgs) IotHubSkuInfoPtrInput {
	return (*iotHubSkuInfoPtrType)(v)
}

func (*iotHubSkuInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubSkuInfo)(nil)).Elem()
}

func (i *iotHubSkuInfoPtrType) ToIotHubSkuInfoPtrOutput() IotHubSkuInfoPtrOutput {
	return i.ToIotHubSkuInfoPtrOutputWithContext(context.Background())
}

func (i *iotHubSkuInfoPtrType) ToIotHubSkuInfoPtrOutputWithContext(ctx context.Context) IotHubSkuInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSkuInfoPtrOutput)
}

type IotHubSkuInfoOutput struct{ *pulumi.OutputState }

func (IotHubSkuInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSkuInfo)(nil)).Elem()
}

func (o IotHubSkuInfoOutput) ToIotHubSkuInfoOutput() IotHubSkuInfoOutput {
	return o
}

func (o IotHubSkuInfoOutput) ToIotHubSkuInfoOutputWithContext(ctx context.Context) IotHubSkuInfoOutput {
	return o
}

func (o IotHubSkuInfoOutput) ToIotHubSkuInfoPtrOutput() IotHubSkuInfoPtrOutput {
	return o.ToIotHubSkuInfoPtrOutputWithContext(context.Background())
}

func (o IotHubSkuInfoOutput) ToIotHubSkuInfoPtrOutputWithContext(ctx context.Context) IotHubSkuInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotHubSkuInfo) *IotHubSkuInfo {
		return &v
	}).(IotHubSkuInfoPtrOutput)
}

func (o IotHubSkuInfoOutput) Capacity() pulumi.Float64Output {
	return o.ApplyT(func(v IotHubSkuInfo) float64 { return v.Capacity }).(pulumi.Float64Output)
}

func (o IotHubSkuInfoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubSkuInfo) string { return v.Name }).(pulumi.StringOutput)
}

type IotHubSkuInfoPtrOutput struct{ *pulumi.OutputState }

func (IotHubSkuInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubSkuInfo)(nil)).Elem()
}

func (o IotHubSkuInfoPtrOutput) ToIotHubSkuInfoPtrOutput() IotHubSkuInfoPtrOutput {
	return o
}

func (o IotHubSkuInfoPtrOutput) ToIotHubSkuInfoPtrOutputWithContext(ctx context.Context) IotHubSkuInfoPtrOutput {
	return o
}

func (o IotHubSkuInfoPtrOutput) Elem() IotHubSkuInfoOutput {
	return o.ApplyT(func(v *IotHubSkuInfo) IotHubSkuInfo {
		if v != nil {
			return *v
		}
		var ret IotHubSkuInfo
		return ret
	}).(IotHubSkuInfoOutput)
}

func (o IotHubSkuInfoPtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *IotHubSkuInfo) *float64 {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o IotHubSkuInfoPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubSkuInfo) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type IotHubSkuInfoResponse struct {
	Capacity float64 `pulumi:"capacity"`
	Name     string  `pulumi:"name"`
	Tier     string  `pulumi:"tier"`
}





type IotHubSkuInfoResponseInput interface {
	pulumi.Input

	ToIotHubSkuInfoResponseOutput() IotHubSkuInfoResponseOutput
	ToIotHubSkuInfoResponseOutputWithContext(context.Context) IotHubSkuInfoResponseOutput
}

type IotHubSkuInfoResponseArgs struct {
	Capacity pulumi.Float64Input `pulumi:"capacity"`
	Name     pulumi.StringInput  `pulumi:"name"`
	Tier     pulumi.StringInput  `pulumi:"tier"`
}

func (IotHubSkuInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSkuInfoResponse)(nil)).Elem()
}

func (i IotHubSkuInfoResponseArgs) ToIotHubSkuInfoResponseOutput() IotHubSkuInfoResponseOutput {
	return i.ToIotHubSkuInfoResponseOutputWithContext(context.Background())
}

func (i IotHubSkuInfoResponseArgs) ToIotHubSkuInfoResponseOutputWithContext(ctx context.Context) IotHubSkuInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSkuInfoResponseOutput)
}

func (i IotHubSkuInfoResponseArgs) ToIotHubSkuInfoResponsePtrOutput() IotHubSkuInfoResponsePtrOutput {
	return i.ToIotHubSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (i IotHubSkuInfoResponseArgs) ToIotHubSkuInfoResponsePtrOutputWithContext(ctx context.Context) IotHubSkuInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSkuInfoResponseOutput).ToIotHubSkuInfoResponsePtrOutputWithContext(ctx)
}









type IotHubSkuInfoResponsePtrInput interface {
	pulumi.Input

	ToIotHubSkuInfoResponsePtrOutput() IotHubSkuInfoResponsePtrOutput
	ToIotHubSkuInfoResponsePtrOutputWithContext(context.Context) IotHubSkuInfoResponsePtrOutput
}

type iotHubSkuInfoResponsePtrType IotHubSkuInfoResponseArgs

func IotHubSkuInfoResponsePtr(v *IotHubSkuInfoResponseArgs) IotHubSkuInfoResponsePtrInput {
	return (*iotHubSkuInfoResponsePtrType)(v)
}

func (*iotHubSkuInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubSkuInfoResponse)(nil)).Elem()
}

func (i *iotHubSkuInfoResponsePtrType) ToIotHubSkuInfoResponsePtrOutput() IotHubSkuInfoResponsePtrOutput {
	return i.ToIotHubSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (i *iotHubSkuInfoResponsePtrType) ToIotHubSkuInfoResponsePtrOutputWithContext(ctx context.Context) IotHubSkuInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotHubSkuInfoResponsePtrOutput)
}

type IotHubSkuInfoResponseOutput struct{ *pulumi.OutputState }

func (IotHubSkuInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotHubSkuInfoResponse)(nil)).Elem()
}

func (o IotHubSkuInfoResponseOutput) ToIotHubSkuInfoResponseOutput() IotHubSkuInfoResponseOutput {
	return o
}

func (o IotHubSkuInfoResponseOutput) ToIotHubSkuInfoResponseOutputWithContext(ctx context.Context) IotHubSkuInfoResponseOutput {
	return o
}

func (o IotHubSkuInfoResponseOutput) ToIotHubSkuInfoResponsePtrOutput() IotHubSkuInfoResponsePtrOutput {
	return o.ToIotHubSkuInfoResponsePtrOutputWithContext(context.Background())
}

func (o IotHubSkuInfoResponseOutput) ToIotHubSkuInfoResponsePtrOutputWithContext(ctx context.Context) IotHubSkuInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IotHubSkuInfoResponse) *IotHubSkuInfoResponse {
		return &v
	}).(IotHubSkuInfoResponsePtrOutput)
}

func (o IotHubSkuInfoResponseOutput) Capacity() pulumi.Float64Output {
	return o.ApplyT(func(v IotHubSkuInfoResponse) float64 { return v.Capacity }).(pulumi.Float64Output)
}

func (o IotHubSkuInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubSkuInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o IotHubSkuInfoResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubSkuInfoResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type IotHubSkuInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (IotHubSkuInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IotHubSkuInfoResponse)(nil)).Elem()
}

func (o IotHubSkuInfoResponsePtrOutput) ToIotHubSkuInfoResponsePtrOutput() IotHubSkuInfoResponsePtrOutput {
	return o
}

func (o IotHubSkuInfoResponsePtrOutput) ToIotHubSkuInfoResponsePtrOutputWithContext(ctx context.Context) IotHubSkuInfoResponsePtrOutput {
	return o
}

func (o IotHubSkuInfoResponsePtrOutput) Elem() IotHubSkuInfoResponseOutput {
	return o.ApplyT(func(v *IotHubSkuInfoResponse) IotHubSkuInfoResponse {
		if v != nil {
			return *v
		}
		var ret IotHubSkuInfoResponse
		return ret
	}).(IotHubSkuInfoResponseOutput)
}

func (o IotHubSkuInfoResponsePtrOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *IotHubSkuInfoResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.Float64PtrOutput)
}

func (o IotHubSkuInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubSkuInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o IotHubSkuInfoResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubSkuInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type IpFilterRule struct {
	Action     IpFilterActionType `pulumi:"action"`
	FilterName string             `pulumi:"filterName"`
	IpMask     string             `pulumi:"ipMask"`
}





type IpFilterRuleInput interface {
	pulumi.Input

	ToIpFilterRuleOutput() IpFilterRuleOutput
	ToIpFilterRuleOutputWithContext(context.Context) IpFilterRuleOutput
}

type IpFilterRuleArgs struct {
	Action     IpFilterActionTypeInput `pulumi:"action"`
	FilterName pulumi.StringInput      `pulumi:"filterName"`
	IpMask     pulumi.StringInput      `pulumi:"ipMask"`
}

func (IpFilterRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterRule)(nil)).Elem()
}

func (i IpFilterRuleArgs) ToIpFilterRuleOutput() IpFilterRuleOutput {
	return i.ToIpFilterRuleOutputWithContext(context.Background())
}

func (i IpFilterRuleArgs) ToIpFilterRuleOutputWithContext(ctx context.Context) IpFilterRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpFilterRuleOutput)
}





type IpFilterRuleArrayInput interface {
	pulumi.Input

	ToIpFilterRuleArrayOutput() IpFilterRuleArrayOutput
	ToIpFilterRuleArrayOutputWithContext(context.Context) IpFilterRuleArrayOutput
}

type IpFilterRuleArray []IpFilterRuleInput

func (IpFilterRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpFilterRule)(nil)).Elem()
}

func (i IpFilterRuleArray) ToIpFilterRuleArrayOutput() IpFilterRuleArrayOutput {
	return i.ToIpFilterRuleArrayOutputWithContext(context.Background())
}

func (i IpFilterRuleArray) ToIpFilterRuleArrayOutputWithContext(ctx context.Context) IpFilterRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpFilterRuleArrayOutput)
}

type IpFilterRuleOutput struct{ *pulumi.OutputState }

func (IpFilterRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterRule)(nil)).Elem()
}

func (o IpFilterRuleOutput) ToIpFilterRuleOutput() IpFilterRuleOutput {
	return o
}

func (o IpFilterRuleOutput) ToIpFilterRuleOutputWithContext(ctx context.Context) IpFilterRuleOutput {
	return o
}

func (o IpFilterRuleOutput) Action() IpFilterActionTypeOutput {
	return o.ApplyT(func(v IpFilterRule) IpFilterActionType { return v.Action }).(IpFilterActionTypeOutput)
}

func (o IpFilterRuleOutput) FilterName() pulumi.StringOutput {
	return o.ApplyT(func(v IpFilterRule) string { return v.FilterName }).(pulumi.StringOutput)
}

func (o IpFilterRuleOutput) IpMask() pulumi.StringOutput {
	return o.ApplyT(func(v IpFilterRule) string { return v.IpMask }).(pulumi.StringOutput)
}

type IpFilterRuleArrayOutput struct{ *pulumi.OutputState }

func (IpFilterRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpFilterRule)(nil)).Elem()
}

func (o IpFilterRuleArrayOutput) ToIpFilterRuleArrayOutput() IpFilterRuleArrayOutput {
	return o
}

func (o IpFilterRuleArrayOutput) ToIpFilterRuleArrayOutputWithContext(ctx context.Context) IpFilterRuleArrayOutput {
	return o
}

func (o IpFilterRuleArrayOutput) Index(i pulumi.IntInput) IpFilterRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpFilterRule {
		return vs[0].([]IpFilterRule)[vs[1].(int)]
	}).(IpFilterRuleOutput)
}

type IpFilterRuleResponse struct {
	Action     string `pulumi:"action"`
	FilterName string `pulumi:"filterName"`
	IpMask     string `pulumi:"ipMask"`
}





type IpFilterRuleResponseInput interface {
	pulumi.Input

	ToIpFilterRuleResponseOutput() IpFilterRuleResponseOutput
	ToIpFilterRuleResponseOutputWithContext(context.Context) IpFilterRuleResponseOutput
}

type IpFilterRuleResponseArgs struct {
	Action     pulumi.StringInput `pulumi:"action"`
	FilterName pulumi.StringInput `pulumi:"filterName"`
	IpMask     pulumi.StringInput `pulumi:"ipMask"`
}

func (IpFilterRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterRuleResponse)(nil)).Elem()
}

func (i IpFilterRuleResponseArgs) ToIpFilterRuleResponseOutput() IpFilterRuleResponseOutput {
	return i.ToIpFilterRuleResponseOutputWithContext(context.Background())
}

func (i IpFilterRuleResponseArgs) ToIpFilterRuleResponseOutputWithContext(ctx context.Context) IpFilterRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpFilterRuleResponseOutput)
}





type IpFilterRuleResponseArrayInput interface {
	pulumi.Input

	ToIpFilterRuleResponseArrayOutput() IpFilterRuleResponseArrayOutput
	ToIpFilterRuleResponseArrayOutputWithContext(context.Context) IpFilterRuleResponseArrayOutput
}

type IpFilterRuleResponseArray []IpFilterRuleResponseInput

func (IpFilterRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpFilterRuleResponse)(nil)).Elem()
}

func (i IpFilterRuleResponseArray) ToIpFilterRuleResponseArrayOutput() IpFilterRuleResponseArrayOutput {
	return i.ToIpFilterRuleResponseArrayOutputWithContext(context.Background())
}

func (i IpFilterRuleResponseArray) ToIpFilterRuleResponseArrayOutputWithContext(ctx context.Context) IpFilterRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpFilterRuleResponseArrayOutput)
}

type IpFilterRuleResponseOutput struct{ *pulumi.OutputState }

func (IpFilterRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpFilterRuleResponse)(nil)).Elem()
}

func (o IpFilterRuleResponseOutput) ToIpFilterRuleResponseOutput() IpFilterRuleResponseOutput {
	return o
}

func (o IpFilterRuleResponseOutput) ToIpFilterRuleResponseOutputWithContext(ctx context.Context) IpFilterRuleResponseOutput {
	return o
}

func (o IpFilterRuleResponseOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v IpFilterRuleResponse) string { return v.Action }).(pulumi.StringOutput)
}

func (o IpFilterRuleResponseOutput) FilterName() pulumi.StringOutput {
	return o.ApplyT(func(v IpFilterRuleResponse) string { return v.FilterName }).(pulumi.StringOutput)
}

func (o IpFilterRuleResponseOutput) IpMask() pulumi.StringOutput {
	return o.ApplyT(func(v IpFilterRuleResponse) string { return v.IpMask }).(pulumi.StringOutput)
}

type IpFilterRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IpFilterRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpFilterRuleResponse)(nil)).Elem()
}

func (o IpFilterRuleResponseArrayOutput) ToIpFilterRuleResponseArrayOutput() IpFilterRuleResponseArrayOutput {
	return o
}

func (o IpFilterRuleResponseArrayOutput) ToIpFilterRuleResponseArrayOutputWithContext(ctx context.Context) IpFilterRuleResponseArrayOutput {
	return o
}

func (o IpFilterRuleResponseArrayOutput) Index(i pulumi.IntInput) IpFilterRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpFilterRuleResponse {
		return vs[0].([]IpFilterRuleResponse)[vs[1].(int)]
	}).(IpFilterRuleResponseOutput)
}

type MessagingEndpointProperties struct {
	LockDurationAsIso8601 *string `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      *int    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          *string `pulumi:"ttlAsIso8601"`
}





type MessagingEndpointPropertiesInput interface {
	pulumi.Input

	ToMessagingEndpointPropertiesOutput() MessagingEndpointPropertiesOutput
	ToMessagingEndpointPropertiesOutputWithContext(context.Context) MessagingEndpointPropertiesOutput
}

type MessagingEndpointPropertiesArgs struct {
	LockDurationAsIso8601 pulumi.StringPtrInput `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      pulumi.IntPtrInput    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          pulumi.StringPtrInput `pulumi:"ttlAsIso8601"`
}

func (MessagingEndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MessagingEndpointProperties)(nil)).Elem()
}

func (i MessagingEndpointPropertiesArgs) ToMessagingEndpointPropertiesOutput() MessagingEndpointPropertiesOutput {
	return i.ToMessagingEndpointPropertiesOutputWithContext(context.Background())
}

func (i MessagingEndpointPropertiesArgs) ToMessagingEndpointPropertiesOutputWithContext(ctx context.Context) MessagingEndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessagingEndpointPropertiesOutput)
}





type MessagingEndpointPropertiesMapInput interface {
	pulumi.Input

	ToMessagingEndpointPropertiesMapOutput() MessagingEndpointPropertiesMapOutput
	ToMessagingEndpointPropertiesMapOutputWithContext(context.Context) MessagingEndpointPropertiesMapOutput
}

type MessagingEndpointPropertiesMap map[string]MessagingEndpointPropertiesInput

func (MessagingEndpointPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MessagingEndpointProperties)(nil)).Elem()
}

func (i MessagingEndpointPropertiesMap) ToMessagingEndpointPropertiesMapOutput() MessagingEndpointPropertiesMapOutput {
	return i.ToMessagingEndpointPropertiesMapOutputWithContext(context.Background())
}

func (i MessagingEndpointPropertiesMap) ToMessagingEndpointPropertiesMapOutputWithContext(ctx context.Context) MessagingEndpointPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessagingEndpointPropertiesMapOutput)
}

type MessagingEndpointPropertiesOutput struct{ *pulumi.OutputState }

func (MessagingEndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessagingEndpointProperties)(nil)).Elem()
}

func (o MessagingEndpointPropertiesOutput) ToMessagingEndpointPropertiesOutput() MessagingEndpointPropertiesOutput {
	return o
}

func (o MessagingEndpointPropertiesOutput) ToMessagingEndpointPropertiesOutputWithContext(ctx context.Context) MessagingEndpointPropertiesOutput {
	return o
}

func (o MessagingEndpointPropertiesOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MessagingEndpointProperties) *string { return v.LockDurationAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o MessagingEndpointPropertiesOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MessagingEndpointProperties) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o MessagingEndpointPropertiesOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MessagingEndpointProperties) *string { return v.TtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type MessagingEndpointPropertiesMapOutput struct{ *pulumi.OutputState }

func (MessagingEndpointPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MessagingEndpointProperties)(nil)).Elem()
}

func (o MessagingEndpointPropertiesMapOutput) ToMessagingEndpointPropertiesMapOutput() MessagingEndpointPropertiesMapOutput {
	return o
}

func (o MessagingEndpointPropertiesMapOutput) ToMessagingEndpointPropertiesMapOutputWithContext(ctx context.Context) MessagingEndpointPropertiesMapOutput {
	return o
}

func (o MessagingEndpointPropertiesMapOutput) MapIndex(k pulumi.StringInput) MessagingEndpointPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MessagingEndpointProperties {
		return vs[0].(map[string]MessagingEndpointProperties)[vs[1].(string)]
	}).(MessagingEndpointPropertiesOutput)
}

type MessagingEndpointPropertiesResponse struct {
	LockDurationAsIso8601 *string `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      *int    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          *string `pulumi:"ttlAsIso8601"`
}





type MessagingEndpointPropertiesResponseInput interface {
	pulumi.Input

	ToMessagingEndpointPropertiesResponseOutput() MessagingEndpointPropertiesResponseOutput
	ToMessagingEndpointPropertiesResponseOutputWithContext(context.Context) MessagingEndpointPropertiesResponseOutput
}

type MessagingEndpointPropertiesResponseArgs struct {
	LockDurationAsIso8601 pulumi.StringPtrInput `pulumi:"lockDurationAsIso8601"`
	MaxDeliveryCount      pulumi.IntPtrInput    `pulumi:"maxDeliveryCount"`
	TtlAsIso8601          pulumi.StringPtrInput `pulumi:"ttlAsIso8601"`
}

func (MessagingEndpointPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MessagingEndpointPropertiesResponse)(nil)).Elem()
}

func (i MessagingEndpointPropertiesResponseArgs) ToMessagingEndpointPropertiesResponseOutput() MessagingEndpointPropertiesResponseOutput {
	return i.ToMessagingEndpointPropertiesResponseOutputWithContext(context.Background())
}

func (i MessagingEndpointPropertiesResponseArgs) ToMessagingEndpointPropertiesResponseOutputWithContext(ctx context.Context) MessagingEndpointPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessagingEndpointPropertiesResponseOutput)
}





type MessagingEndpointPropertiesResponseMapInput interface {
	pulumi.Input

	ToMessagingEndpointPropertiesResponseMapOutput() MessagingEndpointPropertiesResponseMapOutput
	ToMessagingEndpointPropertiesResponseMapOutputWithContext(context.Context) MessagingEndpointPropertiesResponseMapOutput
}

type MessagingEndpointPropertiesResponseMap map[string]MessagingEndpointPropertiesResponseInput

func (MessagingEndpointPropertiesResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MessagingEndpointPropertiesResponse)(nil)).Elem()
}

func (i MessagingEndpointPropertiesResponseMap) ToMessagingEndpointPropertiesResponseMapOutput() MessagingEndpointPropertiesResponseMapOutput {
	return i.ToMessagingEndpointPropertiesResponseMapOutputWithContext(context.Background())
}

func (i MessagingEndpointPropertiesResponseMap) ToMessagingEndpointPropertiesResponseMapOutputWithContext(ctx context.Context) MessagingEndpointPropertiesResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MessagingEndpointPropertiesResponseMapOutput)
}

type MessagingEndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MessagingEndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MessagingEndpointPropertiesResponse)(nil)).Elem()
}

func (o MessagingEndpointPropertiesResponseOutput) ToMessagingEndpointPropertiesResponseOutput() MessagingEndpointPropertiesResponseOutput {
	return o
}

func (o MessagingEndpointPropertiesResponseOutput) ToMessagingEndpointPropertiesResponseOutputWithContext(ctx context.Context) MessagingEndpointPropertiesResponseOutput {
	return o
}

func (o MessagingEndpointPropertiesResponseOutput) LockDurationAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MessagingEndpointPropertiesResponse) *string { return v.LockDurationAsIso8601 }).(pulumi.StringPtrOutput)
}

func (o MessagingEndpointPropertiesResponseOutput) MaxDeliveryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MessagingEndpointPropertiesResponse) *int { return v.MaxDeliveryCount }).(pulumi.IntPtrOutput)
}

func (o MessagingEndpointPropertiesResponseOutput) TtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MessagingEndpointPropertiesResponse) *string { return v.TtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type MessagingEndpointPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (MessagingEndpointPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]MessagingEndpointPropertiesResponse)(nil)).Elem()
}

func (o MessagingEndpointPropertiesResponseMapOutput) ToMessagingEndpointPropertiesResponseMapOutput() MessagingEndpointPropertiesResponseMapOutput {
	return o
}

func (o MessagingEndpointPropertiesResponseMapOutput) ToMessagingEndpointPropertiesResponseMapOutputWithContext(ctx context.Context) MessagingEndpointPropertiesResponseMapOutput {
	return o
}

func (o MessagingEndpointPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) MessagingEndpointPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) MessagingEndpointPropertiesResponse {
		return vs[0].(map[string]MessagingEndpointPropertiesResponse)[vs[1].(string)]
	}).(MessagingEndpointPropertiesResponseOutput)
}

type OperationsMonitoringProperties struct {
	Events map[string]string `pulumi:"events"`
}





type OperationsMonitoringPropertiesInput interface {
	pulumi.Input

	ToOperationsMonitoringPropertiesOutput() OperationsMonitoringPropertiesOutput
	ToOperationsMonitoringPropertiesOutputWithContext(context.Context) OperationsMonitoringPropertiesOutput
}

type OperationsMonitoringPropertiesArgs struct {
	Events pulumi.StringMapInput `pulumi:"events"`
}

func (OperationsMonitoringPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationsMonitoringProperties)(nil)).Elem()
}

func (i OperationsMonitoringPropertiesArgs) ToOperationsMonitoringPropertiesOutput() OperationsMonitoringPropertiesOutput {
	return i.ToOperationsMonitoringPropertiesOutputWithContext(context.Background())
}

func (i OperationsMonitoringPropertiesArgs) ToOperationsMonitoringPropertiesOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationsMonitoringPropertiesOutput)
}

func (i OperationsMonitoringPropertiesArgs) ToOperationsMonitoringPropertiesPtrOutput() OperationsMonitoringPropertiesPtrOutput {
	return i.ToOperationsMonitoringPropertiesPtrOutputWithContext(context.Background())
}

func (i OperationsMonitoringPropertiesArgs) ToOperationsMonitoringPropertiesPtrOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationsMonitoringPropertiesOutput).ToOperationsMonitoringPropertiesPtrOutputWithContext(ctx)
}









type OperationsMonitoringPropertiesPtrInput interface {
	pulumi.Input

	ToOperationsMonitoringPropertiesPtrOutput() OperationsMonitoringPropertiesPtrOutput
	ToOperationsMonitoringPropertiesPtrOutputWithContext(context.Context) OperationsMonitoringPropertiesPtrOutput
}

type operationsMonitoringPropertiesPtrType OperationsMonitoringPropertiesArgs

func OperationsMonitoringPropertiesPtr(v *OperationsMonitoringPropertiesArgs) OperationsMonitoringPropertiesPtrInput {
	return (*operationsMonitoringPropertiesPtrType)(v)
}

func (*operationsMonitoringPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationsMonitoringProperties)(nil)).Elem()
}

func (i *operationsMonitoringPropertiesPtrType) ToOperationsMonitoringPropertiesPtrOutput() OperationsMonitoringPropertiesPtrOutput {
	return i.ToOperationsMonitoringPropertiesPtrOutputWithContext(context.Background())
}

func (i *operationsMonitoringPropertiesPtrType) ToOperationsMonitoringPropertiesPtrOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationsMonitoringPropertiesPtrOutput)
}

type OperationsMonitoringPropertiesOutput struct{ *pulumi.OutputState }

func (OperationsMonitoringPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationsMonitoringProperties)(nil)).Elem()
}

func (o OperationsMonitoringPropertiesOutput) ToOperationsMonitoringPropertiesOutput() OperationsMonitoringPropertiesOutput {
	return o
}

func (o OperationsMonitoringPropertiesOutput) ToOperationsMonitoringPropertiesOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesOutput {
	return o
}

func (o OperationsMonitoringPropertiesOutput) ToOperationsMonitoringPropertiesPtrOutput() OperationsMonitoringPropertiesPtrOutput {
	return o.ToOperationsMonitoringPropertiesPtrOutputWithContext(context.Background())
}

func (o OperationsMonitoringPropertiesOutput) ToOperationsMonitoringPropertiesPtrOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperationsMonitoringProperties) *OperationsMonitoringProperties {
		return &v
	}).(OperationsMonitoringPropertiesPtrOutput)
}

func (o OperationsMonitoringPropertiesOutput) Events() pulumi.StringMapOutput {
	return o.ApplyT(func(v OperationsMonitoringProperties) map[string]string { return v.Events }).(pulumi.StringMapOutput)
}

type OperationsMonitoringPropertiesPtrOutput struct{ *pulumi.OutputState }

func (OperationsMonitoringPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationsMonitoringProperties)(nil)).Elem()
}

func (o OperationsMonitoringPropertiesPtrOutput) ToOperationsMonitoringPropertiesPtrOutput() OperationsMonitoringPropertiesPtrOutput {
	return o
}

func (o OperationsMonitoringPropertiesPtrOutput) ToOperationsMonitoringPropertiesPtrOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesPtrOutput {
	return o
}

func (o OperationsMonitoringPropertiesPtrOutput) Elem() OperationsMonitoringPropertiesOutput {
	return o.ApplyT(func(v *OperationsMonitoringProperties) OperationsMonitoringProperties {
		if v != nil {
			return *v
		}
		var ret OperationsMonitoringProperties
		return ret
	}).(OperationsMonitoringPropertiesOutput)
}

func (o OperationsMonitoringPropertiesPtrOutput) Events() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OperationsMonitoringProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Events
	}).(pulumi.StringMapOutput)
}

type OperationsMonitoringPropertiesResponse struct {
	Events map[string]string `pulumi:"events"`
}





type OperationsMonitoringPropertiesResponseInput interface {
	pulumi.Input

	ToOperationsMonitoringPropertiesResponseOutput() OperationsMonitoringPropertiesResponseOutput
	ToOperationsMonitoringPropertiesResponseOutputWithContext(context.Context) OperationsMonitoringPropertiesResponseOutput
}

type OperationsMonitoringPropertiesResponseArgs struct {
	Events pulumi.StringMapInput `pulumi:"events"`
}

func (OperationsMonitoringPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationsMonitoringPropertiesResponse)(nil)).Elem()
}

func (i OperationsMonitoringPropertiesResponseArgs) ToOperationsMonitoringPropertiesResponseOutput() OperationsMonitoringPropertiesResponseOutput {
	return i.ToOperationsMonitoringPropertiesResponseOutputWithContext(context.Background())
}

func (i OperationsMonitoringPropertiesResponseArgs) ToOperationsMonitoringPropertiesResponseOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationsMonitoringPropertiesResponseOutput)
}

func (i OperationsMonitoringPropertiesResponseArgs) ToOperationsMonitoringPropertiesResponsePtrOutput() OperationsMonitoringPropertiesResponsePtrOutput {
	return i.ToOperationsMonitoringPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i OperationsMonitoringPropertiesResponseArgs) ToOperationsMonitoringPropertiesResponsePtrOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationsMonitoringPropertiesResponseOutput).ToOperationsMonitoringPropertiesResponsePtrOutputWithContext(ctx)
}









type OperationsMonitoringPropertiesResponsePtrInput interface {
	pulumi.Input

	ToOperationsMonitoringPropertiesResponsePtrOutput() OperationsMonitoringPropertiesResponsePtrOutput
	ToOperationsMonitoringPropertiesResponsePtrOutputWithContext(context.Context) OperationsMonitoringPropertiesResponsePtrOutput
}

type operationsMonitoringPropertiesResponsePtrType OperationsMonitoringPropertiesResponseArgs

func OperationsMonitoringPropertiesResponsePtr(v *OperationsMonitoringPropertiesResponseArgs) OperationsMonitoringPropertiesResponsePtrInput {
	return (*operationsMonitoringPropertiesResponsePtrType)(v)
}

func (*operationsMonitoringPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationsMonitoringPropertiesResponse)(nil)).Elem()
}

func (i *operationsMonitoringPropertiesResponsePtrType) ToOperationsMonitoringPropertiesResponsePtrOutput() OperationsMonitoringPropertiesResponsePtrOutput {
	return i.ToOperationsMonitoringPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *operationsMonitoringPropertiesResponsePtrType) ToOperationsMonitoringPropertiesResponsePtrOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OperationsMonitoringPropertiesResponsePtrOutput)
}

type OperationsMonitoringPropertiesResponseOutput struct{ *pulumi.OutputState }

func (OperationsMonitoringPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperationsMonitoringPropertiesResponse)(nil)).Elem()
}

func (o OperationsMonitoringPropertiesResponseOutput) ToOperationsMonitoringPropertiesResponseOutput() OperationsMonitoringPropertiesResponseOutput {
	return o
}

func (o OperationsMonitoringPropertiesResponseOutput) ToOperationsMonitoringPropertiesResponseOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesResponseOutput {
	return o
}

func (o OperationsMonitoringPropertiesResponseOutput) ToOperationsMonitoringPropertiesResponsePtrOutput() OperationsMonitoringPropertiesResponsePtrOutput {
	return o.ToOperationsMonitoringPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o OperationsMonitoringPropertiesResponseOutput) ToOperationsMonitoringPropertiesResponsePtrOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperationsMonitoringPropertiesResponse) *OperationsMonitoringPropertiesResponse {
		return &v
	}).(OperationsMonitoringPropertiesResponsePtrOutput)
}

func (o OperationsMonitoringPropertiesResponseOutput) Events() pulumi.StringMapOutput {
	return o.ApplyT(func(v OperationsMonitoringPropertiesResponse) map[string]string { return v.Events }).(pulumi.StringMapOutput)
}

type OperationsMonitoringPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (OperationsMonitoringPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperationsMonitoringPropertiesResponse)(nil)).Elem()
}

func (o OperationsMonitoringPropertiesResponsePtrOutput) ToOperationsMonitoringPropertiesResponsePtrOutput() OperationsMonitoringPropertiesResponsePtrOutput {
	return o
}

func (o OperationsMonitoringPropertiesResponsePtrOutput) ToOperationsMonitoringPropertiesResponsePtrOutputWithContext(ctx context.Context) OperationsMonitoringPropertiesResponsePtrOutput {
	return o
}

func (o OperationsMonitoringPropertiesResponsePtrOutput) Elem() OperationsMonitoringPropertiesResponseOutput {
	return o.ApplyT(func(v *OperationsMonitoringPropertiesResponse) OperationsMonitoringPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret OperationsMonitoringPropertiesResponse
		return ret
	}).(OperationsMonitoringPropertiesResponseOutput)
}

func (o OperationsMonitoringPropertiesResponsePtrOutput) Events() pulumi.StringMapOutput {
	return o.ApplyT(func(v *OperationsMonitoringPropertiesResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Events
	}).(pulumi.StringMapOutput)
}

type SharedAccessSignatureAuthorizationRule struct {
	KeyName      string       `pulumi:"keyName"`
	PrimaryKey   *string      `pulumi:"primaryKey"`
	Rights       AccessRights `pulumi:"rights"`
	SecondaryKey *string      `pulumi:"secondaryKey"`
}





type SharedAccessSignatureAuthorizationRuleInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleOutput() SharedAccessSignatureAuthorizationRuleOutput
	ToSharedAccessSignatureAuthorizationRuleOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleOutput
}

type SharedAccessSignatureAuthorizationRuleArgs struct {
	KeyName      pulumi.StringInput    `pulumi:"keyName"`
	PrimaryKey   pulumi.StringPtrInput `pulumi:"primaryKey"`
	Rights       AccessRightsInput     `pulumi:"rights"`
	SecondaryKey pulumi.StringPtrInput `pulumi:"secondaryKey"`
}

func (SharedAccessSignatureAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRule)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleArgs) ToSharedAccessSignatureAuthorizationRuleOutput() SharedAccessSignatureAuthorizationRuleOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleArgs) ToSharedAccessSignatureAuthorizationRuleOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleOutput)
}





type SharedAccessSignatureAuthorizationRuleArrayInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleArrayOutput() SharedAccessSignatureAuthorizationRuleArrayOutput
	ToSharedAccessSignatureAuthorizationRuleArrayOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleArrayOutput
}

type SharedAccessSignatureAuthorizationRuleArray []SharedAccessSignatureAuthorizationRuleInput

func (SharedAccessSignatureAuthorizationRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRule)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleArray) ToSharedAccessSignatureAuthorizationRuleArrayOutput() SharedAccessSignatureAuthorizationRuleArrayOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleArrayOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleArray) ToSharedAccessSignatureAuthorizationRuleArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleArrayOutput)
}

type SharedAccessSignatureAuthorizationRuleOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRule)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleOutput) ToSharedAccessSignatureAuthorizationRuleOutput() SharedAccessSignatureAuthorizationRuleOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleOutput) ToSharedAccessSignatureAuthorizationRuleOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRule) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRule) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o SharedAccessSignatureAuthorizationRuleOutput) Rights() AccessRightsOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRule) AccessRights { return v.Rights }).(AccessRightsOutput)
}

func (o SharedAccessSignatureAuthorizationRuleOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRule) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

type SharedAccessSignatureAuthorizationRuleArrayOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRule)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleArrayOutput) ToSharedAccessSignatureAuthorizationRuleArrayOutput() SharedAccessSignatureAuthorizationRuleArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleArrayOutput) ToSharedAccessSignatureAuthorizationRuleArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleArrayOutput) Index(i pulumi.IntInput) SharedAccessSignatureAuthorizationRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedAccessSignatureAuthorizationRule {
		return vs[0].([]SharedAccessSignatureAuthorizationRule)[vs[1].(int)]
	}).(SharedAccessSignatureAuthorizationRuleOutput)
}

type SharedAccessSignatureAuthorizationRuleResponse struct {
	KeyName      string  `pulumi:"keyName"`
	PrimaryKey   *string `pulumi:"primaryKey"`
	Rights       string  `pulumi:"rights"`
	SecondaryKey *string `pulumi:"secondaryKey"`
}





type SharedAccessSignatureAuthorizationRuleResponseInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleResponseOutput() SharedAccessSignatureAuthorizationRuleResponseOutput
	ToSharedAccessSignatureAuthorizationRuleResponseOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleResponseOutput
}

type SharedAccessSignatureAuthorizationRuleResponseArgs struct {
	KeyName      pulumi.StringInput    `pulumi:"keyName"`
	PrimaryKey   pulumi.StringPtrInput `pulumi:"primaryKey"`
	Rights       pulumi.StringInput    `pulumi:"rights"`
	SecondaryKey pulumi.StringPtrInput `pulumi:"secondaryKey"`
}

func (SharedAccessSignatureAuthorizationRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleResponse)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleResponseArgs) ToSharedAccessSignatureAuthorizationRuleResponseOutput() SharedAccessSignatureAuthorizationRuleResponseOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleResponseOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleResponseArgs) ToSharedAccessSignatureAuthorizationRuleResponseOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleResponseOutput)
}





type SharedAccessSignatureAuthorizationRuleResponseArrayInput interface {
	pulumi.Input

	ToSharedAccessSignatureAuthorizationRuleResponseArrayOutput() SharedAccessSignatureAuthorizationRuleResponseArrayOutput
	ToSharedAccessSignatureAuthorizationRuleResponseArrayOutputWithContext(context.Context) SharedAccessSignatureAuthorizationRuleResponseArrayOutput
}

type SharedAccessSignatureAuthorizationRuleResponseArray []SharedAccessSignatureAuthorizationRuleResponseInput

func (SharedAccessSignatureAuthorizationRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRuleResponse)(nil)).Elem()
}

func (i SharedAccessSignatureAuthorizationRuleResponseArray) ToSharedAccessSignatureAuthorizationRuleResponseArrayOutput() SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return i.ToSharedAccessSignatureAuthorizationRuleResponseArrayOutputWithContext(context.Background())
}

func (i SharedAccessSignatureAuthorizationRuleResponseArray) ToSharedAccessSignatureAuthorizationRuleResponseArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SharedAccessSignatureAuthorizationRuleResponseArrayOutput)
}

type SharedAccessSignatureAuthorizationRuleResponseOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleResponse)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) ToSharedAccessSignatureAuthorizationRuleResponseOutput() SharedAccessSignatureAuthorizationRuleResponseOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) ToSharedAccessSignatureAuthorizationRuleResponseOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleResponseOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleResponse) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) Rights() pulumi.StringOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleResponse) string { return v.Rights }).(pulumi.StringOutput)
}

func (o SharedAccessSignatureAuthorizationRuleResponseOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SharedAccessSignatureAuthorizationRuleResponse) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

type SharedAccessSignatureAuthorizationRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (SharedAccessSignatureAuthorizationRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SharedAccessSignatureAuthorizationRuleResponse)(nil)).Elem()
}

func (o SharedAccessSignatureAuthorizationRuleResponseArrayOutput) ToSharedAccessSignatureAuthorizationRuleResponseArrayOutput() SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleResponseArrayOutput) ToSharedAccessSignatureAuthorizationRuleResponseArrayOutputWithContext(ctx context.Context) SharedAccessSignatureAuthorizationRuleResponseArrayOutput {
	return o
}

func (o SharedAccessSignatureAuthorizationRuleResponseArrayOutput) Index(i pulumi.IntInput) SharedAccessSignatureAuthorizationRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SharedAccessSignatureAuthorizationRuleResponse {
		return vs[0].([]SharedAccessSignatureAuthorizationRuleResponse)[vs[1].(int)]
	}).(SharedAccessSignatureAuthorizationRuleResponseOutput)
}

type StorageEndpointProperties struct {
	ConnectionString string  `pulumi:"connectionString"`
	ContainerName    string  `pulumi:"containerName"`
	SasTtlAsIso8601  *string `pulumi:"sasTtlAsIso8601"`
}





type StorageEndpointPropertiesInput interface {
	pulumi.Input

	ToStorageEndpointPropertiesOutput() StorageEndpointPropertiesOutput
	ToStorageEndpointPropertiesOutputWithContext(context.Context) StorageEndpointPropertiesOutput
}

type StorageEndpointPropertiesArgs struct {
	ConnectionString pulumi.StringInput    `pulumi:"connectionString"`
	ContainerName    pulumi.StringInput    `pulumi:"containerName"`
	SasTtlAsIso8601  pulumi.StringPtrInput `pulumi:"sasTtlAsIso8601"`
}

func (StorageEndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageEndpointProperties)(nil)).Elem()
}

func (i StorageEndpointPropertiesArgs) ToStorageEndpointPropertiesOutput() StorageEndpointPropertiesOutput {
	return i.ToStorageEndpointPropertiesOutputWithContext(context.Background())
}

func (i StorageEndpointPropertiesArgs) ToStorageEndpointPropertiesOutputWithContext(ctx context.Context) StorageEndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageEndpointPropertiesOutput)
}





type StorageEndpointPropertiesMapInput interface {
	pulumi.Input

	ToStorageEndpointPropertiesMapOutput() StorageEndpointPropertiesMapOutput
	ToStorageEndpointPropertiesMapOutputWithContext(context.Context) StorageEndpointPropertiesMapOutput
}

type StorageEndpointPropertiesMap map[string]StorageEndpointPropertiesInput

func (StorageEndpointPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StorageEndpointProperties)(nil)).Elem()
}

func (i StorageEndpointPropertiesMap) ToStorageEndpointPropertiesMapOutput() StorageEndpointPropertiesMapOutput {
	return i.ToStorageEndpointPropertiesMapOutputWithContext(context.Background())
}

func (i StorageEndpointPropertiesMap) ToStorageEndpointPropertiesMapOutputWithContext(ctx context.Context) StorageEndpointPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageEndpointPropertiesMapOutput)
}

type StorageEndpointPropertiesOutput struct{ *pulumi.OutputState }

func (StorageEndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageEndpointProperties)(nil)).Elem()
}

func (o StorageEndpointPropertiesOutput) ToStorageEndpointPropertiesOutput() StorageEndpointPropertiesOutput {
	return o
}

func (o StorageEndpointPropertiesOutput) ToStorageEndpointPropertiesOutputWithContext(ctx context.Context) StorageEndpointPropertiesOutput {
	return o
}

func (o StorageEndpointPropertiesOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v StorageEndpointProperties) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o StorageEndpointPropertiesOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v StorageEndpointProperties) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o StorageEndpointPropertiesOutput) SasTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageEndpointProperties) *string { return v.SasTtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type StorageEndpointPropertiesMapOutput struct{ *pulumi.OutputState }

func (StorageEndpointPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StorageEndpointProperties)(nil)).Elem()
}

func (o StorageEndpointPropertiesMapOutput) ToStorageEndpointPropertiesMapOutput() StorageEndpointPropertiesMapOutput {
	return o
}

func (o StorageEndpointPropertiesMapOutput) ToStorageEndpointPropertiesMapOutputWithContext(ctx context.Context) StorageEndpointPropertiesMapOutput {
	return o
}

func (o StorageEndpointPropertiesMapOutput) MapIndex(k pulumi.StringInput) StorageEndpointPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StorageEndpointProperties {
		return vs[0].(map[string]StorageEndpointProperties)[vs[1].(string)]
	}).(StorageEndpointPropertiesOutput)
}

type StorageEndpointPropertiesResponse struct {
	ConnectionString string  `pulumi:"connectionString"`
	ContainerName    string  `pulumi:"containerName"`
	SasTtlAsIso8601  *string `pulumi:"sasTtlAsIso8601"`
}





type StorageEndpointPropertiesResponseInput interface {
	pulumi.Input

	ToStorageEndpointPropertiesResponseOutput() StorageEndpointPropertiesResponseOutput
	ToStorageEndpointPropertiesResponseOutputWithContext(context.Context) StorageEndpointPropertiesResponseOutput
}

type StorageEndpointPropertiesResponseArgs struct {
	ConnectionString pulumi.StringInput    `pulumi:"connectionString"`
	ContainerName    pulumi.StringInput    `pulumi:"containerName"`
	SasTtlAsIso8601  pulumi.StringPtrInput `pulumi:"sasTtlAsIso8601"`
}

func (StorageEndpointPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageEndpointPropertiesResponse)(nil)).Elem()
}

func (i StorageEndpointPropertiesResponseArgs) ToStorageEndpointPropertiesResponseOutput() StorageEndpointPropertiesResponseOutput {
	return i.ToStorageEndpointPropertiesResponseOutputWithContext(context.Background())
}

func (i StorageEndpointPropertiesResponseArgs) ToStorageEndpointPropertiesResponseOutputWithContext(ctx context.Context) StorageEndpointPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageEndpointPropertiesResponseOutput)
}





type StorageEndpointPropertiesResponseMapInput interface {
	pulumi.Input

	ToStorageEndpointPropertiesResponseMapOutput() StorageEndpointPropertiesResponseMapOutput
	ToStorageEndpointPropertiesResponseMapOutputWithContext(context.Context) StorageEndpointPropertiesResponseMapOutput
}

type StorageEndpointPropertiesResponseMap map[string]StorageEndpointPropertiesResponseInput

func (StorageEndpointPropertiesResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StorageEndpointPropertiesResponse)(nil)).Elem()
}

func (i StorageEndpointPropertiesResponseMap) ToStorageEndpointPropertiesResponseMapOutput() StorageEndpointPropertiesResponseMapOutput {
	return i.ToStorageEndpointPropertiesResponseMapOutputWithContext(context.Background())
}

func (i StorageEndpointPropertiesResponseMap) ToStorageEndpointPropertiesResponseMapOutputWithContext(ctx context.Context) StorageEndpointPropertiesResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageEndpointPropertiesResponseMapOutput)
}

type StorageEndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageEndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageEndpointPropertiesResponse)(nil)).Elem()
}

func (o StorageEndpointPropertiesResponseOutput) ToStorageEndpointPropertiesResponseOutput() StorageEndpointPropertiesResponseOutput {
	return o
}

func (o StorageEndpointPropertiesResponseOutput) ToStorageEndpointPropertiesResponseOutputWithContext(ctx context.Context) StorageEndpointPropertiesResponseOutput {
	return o
}

func (o StorageEndpointPropertiesResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v StorageEndpointPropertiesResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o StorageEndpointPropertiesResponseOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v StorageEndpointPropertiesResponse) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o StorageEndpointPropertiesResponseOutput) SasTtlAsIso8601() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageEndpointPropertiesResponse) *string { return v.SasTtlAsIso8601 }).(pulumi.StringPtrOutput)
}

type StorageEndpointPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (StorageEndpointPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StorageEndpointPropertiesResponse)(nil)).Elem()
}

func (o StorageEndpointPropertiesResponseMapOutput) ToStorageEndpointPropertiesResponseMapOutput() StorageEndpointPropertiesResponseMapOutput {
	return o
}

func (o StorageEndpointPropertiesResponseMapOutput) ToStorageEndpointPropertiesResponseMapOutputWithContext(ctx context.Context) StorageEndpointPropertiesResponseMapOutput {
	return o
}

func (o StorageEndpointPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) StorageEndpointPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StorageEndpointPropertiesResponse {
		return vs[0].(map[string]StorageEndpointPropertiesResponse)[vs[1].(string)]
	}).(StorageEndpointPropertiesResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudToDevicePropertiesInput)(nil)).Elem(), CloudToDevicePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudToDevicePropertiesPtrInput)(nil)).Elem(), CloudToDevicePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudToDevicePropertiesResponseInput)(nil)).Elem(), CloudToDevicePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudToDevicePropertiesResponsePtrInput)(nil)).Elem(), CloudToDevicePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubPropertiesInput)(nil)).Elem(), EventHubPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubPropertiesMapInput)(nil)).Elem(), EventHubPropertiesMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubPropertiesResponseInput)(nil)).Elem(), EventHubPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubPropertiesResponseMapInput)(nil)).Elem(), EventHubPropertiesResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeedbackPropertiesInput)(nil)).Elem(), FeedbackPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeedbackPropertiesPtrInput)(nil)).Elem(), FeedbackPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeedbackPropertiesResponseInput)(nil)).Elem(), FeedbackPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeedbackPropertiesResponsePtrInput)(nil)).Elem(), FeedbackPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubPropertiesInput)(nil)).Elem(), IotHubPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubPropertiesPtrInput)(nil)).Elem(), IotHubPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubPropertiesResponseInput)(nil)).Elem(), IotHubPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubPropertiesResponsePtrInput)(nil)).Elem(), IotHubPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubSkuInfoInput)(nil)).Elem(), IotHubSkuInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubSkuInfoPtrInput)(nil)).Elem(), IotHubSkuInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubSkuInfoResponseInput)(nil)).Elem(), IotHubSkuInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IotHubSkuInfoResponsePtrInput)(nil)).Elem(), IotHubSkuInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterRuleInput)(nil)).Elem(), IpFilterRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterRuleArrayInput)(nil)).Elem(), IpFilterRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterRuleResponseInput)(nil)).Elem(), IpFilterRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IpFilterRuleResponseArrayInput)(nil)).Elem(), IpFilterRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MessagingEndpointPropertiesInput)(nil)).Elem(), MessagingEndpointPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MessagingEndpointPropertiesMapInput)(nil)).Elem(), MessagingEndpointPropertiesMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*MessagingEndpointPropertiesResponseInput)(nil)).Elem(), MessagingEndpointPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MessagingEndpointPropertiesResponseMapInput)(nil)).Elem(), MessagingEndpointPropertiesResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*OperationsMonitoringPropertiesInput)(nil)).Elem(), OperationsMonitoringPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OperationsMonitoringPropertiesPtrInput)(nil)).Elem(), OperationsMonitoringPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OperationsMonitoringPropertiesResponseInput)(nil)).Elem(), OperationsMonitoringPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OperationsMonitoringPropertiesResponsePtrInput)(nil)).Elem(), OperationsMonitoringPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleInput)(nil)).Elem(), SharedAccessSignatureAuthorizationRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleArrayInput)(nil)).Elem(), SharedAccessSignatureAuthorizationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleResponseInput)(nil)).Elem(), SharedAccessSignatureAuthorizationRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleResponseArrayInput)(nil)).Elem(), SharedAccessSignatureAuthorizationRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageEndpointPropertiesInput)(nil)).Elem(), StorageEndpointPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageEndpointPropertiesMapInput)(nil)).Elem(), StorageEndpointPropertiesMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageEndpointPropertiesResponseInput)(nil)).Elem(), StorageEndpointPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageEndpointPropertiesResponseMapInput)(nil)).Elem(), StorageEndpointPropertiesResponseMap{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesResponseOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesMapOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(FeedbackPropertiesOutput{})
	pulumi.RegisterOutputType(FeedbackPropertiesPtrOutput{})
	pulumi.RegisterOutputType(FeedbackPropertiesResponseOutput{})
	pulumi.RegisterOutputType(FeedbackPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IotHubPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IotHubSkuInfoOutput{})
	pulumi.RegisterOutputType(IotHubSkuInfoPtrOutput{})
	pulumi.RegisterOutputType(IotHubSkuInfoResponseOutput{})
	pulumi.RegisterOutputType(IotHubSkuInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(IpFilterRuleOutput{})
	pulumi.RegisterOutputType(IpFilterRuleArrayOutput{})
	pulumi.RegisterOutputType(IpFilterRuleResponseOutput{})
	pulumi.RegisterOutputType(IpFilterRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(MessagingEndpointPropertiesOutput{})
	pulumi.RegisterOutputType(MessagingEndpointPropertiesMapOutput{})
	pulumi.RegisterOutputType(MessagingEndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MessagingEndpointPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(OperationsMonitoringPropertiesOutput{})
	pulumi.RegisterOutputType(OperationsMonitoringPropertiesPtrOutput{})
	pulumi.RegisterOutputType(OperationsMonitoringPropertiesResponseOutput{})
	pulumi.RegisterOutputType(OperationsMonitoringPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleResponseOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesMapOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesResponseMapOutput{})
}
