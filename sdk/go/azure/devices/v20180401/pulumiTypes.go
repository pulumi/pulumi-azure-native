


package v20180401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CertificatePropertiesResponse struct {
	Certificate *string `pulumi:"certificate"`
	Created     string  `pulumi:"created"`
	Expiry      string  `pulumi:"expiry"`
	IsVerified  bool    `pulumi:"isVerified"`
	Subject     string  `pulumi:"subject"`
	Thumbprint  string  `pulumi:"thumbprint"`
	Updated     string  `pulumi:"updated"`
}





type CertificatePropertiesResponseInput interface {
	pulumi.Input

	ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput
	ToCertificatePropertiesResponseOutputWithContext(context.Context) CertificatePropertiesResponseOutput
}

type CertificatePropertiesResponseArgs struct {
	Certificate pulumi.StringPtrInput `pulumi:"certificate"`
	Created     pulumi.StringInput    `pulumi:"created"`
	Expiry      pulumi.StringInput    `pulumi:"expiry"`
	IsVerified  pulumi.BoolInput      `pulumi:"isVerified"`
	Subject     pulumi.StringInput    `pulumi:"subject"`
	Thumbprint  pulumi.StringInput    `pulumi:"thumbprint"`
	Updated     pulumi.StringInput    `pulumi:"updated"`
}

func (CertificatePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatePropertiesResponse)(nil)).Elem()
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput {
	return i.ToCertificatePropertiesResponseOutputWithContext(context.Background())
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponseOutputWithContext(ctx context.Context) CertificatePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesResponseOutput)
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return i.ToCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CertificatePropertiesResponseArgs) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesResponseOutput).ToCertificatePropertiesResponsePtrOutputWithContext(ctx)
}









type CertificatePropertiesResponsePtrInput interface {
	pulumi.Input

	ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput
	ToCertificatePropertiesResponsePtrOutputWithContext(context.Context) CertificatePropertiesResponsePtrOutput
}

type certificatePropertiesResponsePtrType CertificatePropertiesResponseArgs

func CertificatePropertiesResponsePtr(v *CertificatePropertiesResponseArgs) CertificatePropertiesResponsePtrInput {
	return (*certificatePropertiesResponsePtrType)(v)
}

func (*certificatePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificatePropertiesResponse)(nil)).Elem()
}

func (i *certificatePropertiesResponsePtrType) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return i.ToCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *certificatePropertiesResponsePtrType) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificatePropertiesResponsePtrOutput)
}

type CertificatePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificatePropertiesResponse)(nil)).Elem()
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutput() CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponseOutputWithContext(ctx context.Context) CertificatePropertiesResponseOutput {
	return o
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return o.ToCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CertificatePropertiesResponseOutput) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CertificatePropertiesResponse) *CertificatePropertiesResponse {
		return &v
	}).(CertificatePropertiesResponsePtrOutput)
}

func (o CertificatePropertiesResponseOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) *string { return v.Certificate }).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponseOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Created }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Expiry() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Expiry }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) IsVerified() pulumi.BoolOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) bool { return v.IsVerified }).(pulumi.BoolOutput)
}

func (o CertificatePropertiesResponseOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Subject }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Thumbprint() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Thumbprint }).(pulumi.StringOutput)
}

func (o CertificatePropertiesResponseOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v CertificatePropertiesResponse) string { return v.Updated }).(pulumi.StringOutput)
}

type CertificatePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CertificatePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertificatePropertiesResponse)(nil)).Elem()
}

func (o CertificatePropertiesResponsePtrOutput) ToCertificatePropertiesResponsePtrOutput() CertificatePropertiesResponsePtrOutput {
	return o
}

func (o CertificatePropertiesResponsePtrOutput) ToCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) CertificatePropertiesResponsePtrOutput {
	return o
}

func (o CertificatePropertiesResponsePtrOutput) Elem() CertificatePropertiesResponseOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) CertificatePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CertificatePropertiesResponse
		return ret
	}).(CertificatePropertiesResponseOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Certificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Certificate
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Created
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Expiry
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) IsVerified() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsVerified
	}).(pulumi.BoolPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Subject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Subject
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Thumbprint
	}).(pulumi.StringPtrOutput)
}

func (o CertificatePropertiesResponsePtrOutput) Updated() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Updated
	}).(pulumi.StringPtrOutput)
}

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

type FallbackRouteProperties struct {
	Condition     *string  `pulumi:"condition"`
	EndpointNames []string `pulumi:"endpointNames"`
	IsEnabled     bool     `pulumi:"isEnabled"`
	Name          *string  `pulumi:"name"`
	Source        string   `pulumi:"source"`
}





type FallbackRoutePropertiesInput interface {
	pulumi.Input

	ToFallbackRoutePropertiesOutput() FallbackRoutePropertiesOutput
	ToFallbackRoutePropertiesOutputWithContext(context.Context) FallbackRoutePropertiesOutput
}

type FallbackRoutePropertiesArgs struct {
	Condition     pulumi.StringPtrInput   `pulumi:"condition"`
	EndpointNames pulumi.StringArrayInput `pulumi:"endpointNames"`
	IsEnabled     pulumi.BoolInput        `pulumi:"isEnabled"`
	Name          pulumi.StringPtrInput   `pulumi:"name"`
	Source        pulumi.StringInput      `pulumi:"source"`
}

func (FallbackRoutePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FallbackRouteProperties)(nil)).Elem()
}

func (i FallbackRoutePropertiesArgs) ToFallbackRoutePropertiesOutput() FallbackRoutePropertiesOutput {
	return i.ToFallbackRoutePropertiesOutputWithContext(context.Background())
}

func (i FallbackRoutePropertiesArgs) ToFallbackRoutePropertiesOutputWithContext(ctx context.Context) FallbackRoutePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePropertiesOutput)
}

func (i FallbackRoutePropertiesArgs) ToFallbackRoutePropertiesPtrOutput() FallbackRoutePropertiesPtrOutput {
	return i.ToFallbackRoutePropertiesPtrOutputWithContext(context.Background())
}

func (i FallbackRoutePropertiesArgs) ToFallbackRoutePropertiesPtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePropertiesOutput).ToFallbackRoutePropertiesPtrOutputWithContext(ctx)
}









type FallbackRoutePropertiesPtrInput interface {
	pulumi.Input

	ToFallbackRoutePropertiesPtrOutput() FallbackRoutePropertiesPtrOutput
	ToFallbackRoutePropertiesPtrOutputWithContext(context.Context) FallbackRoutePropertiesPtrOutput
}

type fallbackRoutePropertiesPtrType FallbackRoutePropertiesArgs

func FallbackRoutePropertiesPtr(v *FallbackRoutePropertiesArgs) FallbackRoutePropertiesPtrInput {
	return (*fallbackRoutePropertiesPtrType)(v)
}

func (*fallbackRoutePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FallbackRouteProperties)(nil)).Elem()
}

func (i *fallbackRoutePropertiesPtrType) ToFallbackRoutePropertiesPtrOutput() FallbackRoutePropertiesPtrOutput {
	return i.ToFallbackRoutePropertiesPtrOutputWithContext(context.Background())
}

func (i *fallbackRoutePropertiesPtrType) ToFallbackRoutePropertiesPtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePropertiesPtrOutput)
}

type FallbackRoutePropertiesOutput struct{ *pulumi.OutputState }

func (FallbackRoutePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FallbackRouteProperties)(nil)).Elem()
}

func (o FallbackRoutePropertiesOutput) ToFallbackRoutePropertiesOutput() FallbackRoutePropertiesOutput {
	return o
}

func (o FallbackRoutePropertiesOutput) ToFallbackRoutePropertiesOutputWithContext(ctx context.Context) FallbackRoutePropertiesOutput {
	return o
}

func (o FallbackRoutePropertiesOutput) ToFallbackRoutePropertiesPtrOutput() FallbackRoutePropertiesPtrOutput {
	return o.ToFallbackRoutePropertiesPtrOutputWithContext(context.Background())
}

func (o FallbackRoutePropertiesOutput) ToFallbackRoutePropertiesPtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FallbackRouteProperties) *FallbackRouteProperties {
		return &v
	}).(FallbackRoutePropertiesPtrOutput)
}

func (o FallbackRoutePropertiesOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FallbackRouteProperties) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FallbackRouteProperties) []string { return v.EndpointNames }).(pulumi.StringArrayOutput)
}

func (o FallbackRoutePropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FallbackRouteProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o FallbackRoutePropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FallbackRouteProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v FallbackRouteProperties) string { return v.Source }).(pulumi.StringOutput)
}

type FallbackRoutePropertiesPtrOutput struct{ *pulumi.OutputState }

func (FallbackRoutePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FallbackRouteProperties)(nil)).Elem()
}

func (o FallbackRoutePropertiesPtrOutput) ToFallbackRoutePropertiesPtrOutput() FallbackRoutePropertiesPtrOutput {
	return o
}

func (o FallbackRoutePropertiesPtrOutput) ToFallbackRoutePropertiesPtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesPtrOutput {
	return o
}

func (o FallbackRoutePropertiesPtrOutput) Elem() FallbackRoutePropertiesOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) FallbackRouteProperties {
		if v != nil {
			return *v
		}
		var ret FallbackRouteProperties
		return ret
	}).(FallbackRoutePropertiesOutput)
}

func (o FallbackRoutePropertiesPtrOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) *string {
		if v == nil {
			return nil
		}
		return v.Condition
	}).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesPtrOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) []string {
		if v == nil {
			return nil
		}
		return v.EndpointNames
	}).(pulumi.StringArrayOutput)
}

func (o FallbackRoutePropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o FallbackRoutePropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesPtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRouteProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Source
	}).(pulumi.StringPtrOutput)
}

type FallbackRoutePropertiesResponse struct {
	Condition     *string  `pulumi:"condition"`
	EndpointNames []string `pulumi:"endpointNames"`
	IsEnabled     bool     `pulumi:"isEnabled"`
	Name          *string  `pulumi:"name"`
	Source        string   `pulumi:"source"`
}





type FallbackRoutePropertiesResponseInput interface {
	pulumi.Input

	ToFallbackRoutePropertiesResponseOutput() FallbackRoutePropertiesResponseOutput
	ToFallbackRoutePropertiesResponseOutputWithContext(context.Context) FallbackRoutePropertiesResponseOutput
}

type FallbackRoutePropertiesResponseArgs struct {
	Condition     pulumi.StringPtrInput   `pulumi:"condition"`
	EndpointNames pulumi.StringArrayInput `pulumi:"endpointNames"`
	IsEnabled     pulumi.BoolInput        `pulumi:"isEnabled"`
	Name          pulumi.StringPtrInput   `pulumi:"name"`
	Source        pulumi.StringInput      `pulumi:"source"`
}

func (FallbackRoutePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FallbackRoutePropertiesResponse)(nil)).Elem()
}

func (i FallbackRoutePropertiesResponseArgs) ToFallbackRoutePropertiesResponseOutput() FallbackRoutePropertiesResponseOutput {
	return i.ToFallbackRoutePropertiesResponseOutputWithContext(context.Background())
}

func (i FallbackRoutePropertiesResponseArgs) ToFallbackRoutePropertiesResponseOutputWithContext(ctx context.Context) FallbackRoutePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePropertiesResponseOutput)
}

func (i FallbackRoutePropertiesResponseArgs) ToFallbackRoutePropertiesResponsePtrOutput() FallbackRoutePropertiesResponsePtrOutput {
	return i.ToFallbackRoutePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i FallbackRoutePropertiesResponseArgs) ToFallbackRoutePropertiesResponsePtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePropertiesResponseOutput).ToFallbackRoutePropertiesResponsePtrOutputWithContext(ctx)
}









type FallbackRoutePropertiesResponsePtrInput interface {
	pulumi.Input

	ToFallbackRoutePropertiesResponsePtrOutput() FallbackRoutePropertiesResponsePtrOutput
	ToFallbackRoutePropertiesResponsePtrOutputWithContext(context.Context) FallbackRoutePropertiesResponsePtrOutput
}

type fallbackRoutePropertiesResponsePtrType FallbackRoutePropertiesResponseArgs

func FallbackRoutePropertiesResponsePtr(v *FallbackRoutePropertiesResponseArgs) FallbackRoutePropertiesResponsePtrInput {
	return (*fallbackRoutePropertiesResponsePtrType)(v)
}

func (*fallbackRoutePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FallbackRoutePropertiesResponse)(nil)).Elem()
}

func (i *fallbackRoutePropertiesResponsePtrType) ToFallbackRoutePropertiesResponsePtrOutput() FallbackRoutePropertiesResponsePtrOutput {
	return i.ToFallbackRoutePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *fallbackRoutePropertiesResponsePtrType) ToFallbackRoutePropertiesResponsePtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FallbackRoutePropertiesResponsePtrOutput)
}

type FallbackRoutePropertiesResponseOutput struct{ *pulumi.OutputState }

func (FallbackRoutePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FallbackRoutePropertiesResponse)(nil)).Elem()
}

func (o FallbackRoutePropertiesResponseOutput) ToFallbackRoutePropertiesResponseOutput() FallbackRoutePropertiesResponseOutput {
	return o
}

func (o FallbackRoutePropertiesResponseOutput) ToFallbackRoutePropertiesResponseOutputWithContext(ctx context.Context) FallbackRoutePropertiesResponseOutput {
	return o
}

func (o FallbackRoutePropertiesResponseOutput) ToFallbackRoutePropertiesResponsePtrOutput() FallbackRoutePropertiesResponsePtrOutput {
	return o.ToFallbackRoutePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o FallbackRoutePropertiesResponseOutput) ToFallbackRoutePropertiesResponsePtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FallbackRoutePropertiesResponse) *FallbackRoutePropertiesResponse {
		return &v
	}).(FallbackRoutePropertiesResponsePtrOutput)
}

func (o FallbackRoutePropertiesResponseOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FallbackRoutePropertiesResponse) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesResponseOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FallbackRoutePropertiesResponse) []string { return v.EndpointNames }).(pulumi.StringArrayOutput)
}

func (o FallbackRoutePropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FallbackRoutePropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o FallbackRoutePropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FallbackRoutePropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v FallbackRoutePropertiesResponse) string { return v.Source }).(pulumi.StringOutput)
}

type FallbackRoutePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (FallbackRoutePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FallbackRoutePropertiesResponse)(nil)).Elem()
}

func (o FallbackRoutePropertiesResponsePtrOutput) ToFallbackRoutePropertiesResponsePtrOutput() FallbackRoutePropertiesResponsePtrOutput {
	return o
}

func (o FallbackRoutePropertiesResponsePtrOutput) ToFallbackRoutePropertiesResponsePtrOutputWithContext(ctx context.Context) FallbackRoutePropertiesResponsePtrOutput {
	return o
}

func (o FallbackRoutePropertiesResponsePtrOutput) Elem() FallbackRoutePropertiesResponseOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) FallbackRoutePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret FallbackRoutePropertiesResponse
		return ret
	}).(FallbackRoutePropertiesResponseOutput)
}

func (o FallbackRoutePropertiesResponsePtrOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Condition
	}).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesResponsePtrOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.EndpointNames
	}).(pulumi.StringArrayOutput)
}

func (o FallbackRoutePropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o FallbackRoutePropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o FallbackRoutePropertiesResponsePtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FallbackRoutePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Source
	}).(pulumi.StringPtrOutput)
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
	Routing                        *RoutingProperties                       `pulumi:"routing"`
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
	Routing                        RoutingPropertiesPtrInput                        `pulumi:"routing"`
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

func (o IotHubPropertiesOutput) Routing() RoutingPropertiesPtrOutput {
	return o.ApplyT(func(v IotHubProperties) *RoutingProperties { return v.Routing }).(RoutingPropertiesPtrOutput)
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

func (o IotHubPropertiesPtrOutput) Routing() RoutingPropertiesPtrOutput {
	return o.ApplyT(func(v *IotHubProperties) *RoutingProperties {
		if v == nil {
			return nil
		}
		return v.Routing
	}).(RoutingPropertiesPtrOutput)
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
	Routing                        *RoutingPropertiesResponse                       `pulumi:"routing"`
	State                          string                                           `pulumi:"state"`
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
	Routing                        RoutingPropertiesResponsePtrInput                        `pulumi:"routing"`
	State                          pulumi.StringInput                                       `pulumi:"state"`
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

func (o IotHubPropertiesResponseOutput) Routing() RoutingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) *RoutingPropertiesResponse { return v.Routing }).(RoutingPropertiesResponsePtrOutput)
}

func (o IotHubPropertiesResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v IotHubPropertiesResponse) string { return v.State }).(pulumi.StringOutput)
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

func (o IotHubPropertiesResponsePtrOutput) Routing() RoutingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) *RoutingPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Routing
	}).(RoutingPropertiesResponsePtrOutput)
}

func (o IotHubPropertiesResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IotHubPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
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
	Capacity *float64 `pulumi:"capacity"`
	Name     string   `pulumi:"name"`
}





type IotHubSkuInfoInput interface {
	pulumi.Input

	ToIotHubSkuInfoOutput() IotHubSkuInfoOutput
	ToIotHubSkuInfoOutputWithContext(context.Context) IotHubSkuInfoOutput
}

type IotHubSkuInfoArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput     `pulumi:"name"`
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

func (o IotHubSkuInfoOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IotHubSkuInfo) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
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
		return v.Capacity
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
	Capacity *float64 `pulumi:"capacity"`
	Name     string   `pulumi:"name"`
	Tier     string   `pulumi:"tier"`
}





type IotHubSkuInfoResponseInput interface {
	pulumi.Input

	ToIotHubSkuInfoResponseOutput() IotHubSkuInfoResponseOutput
	ToIotHubSkuInfoResponseOutputWithContext(context.Context) IotHubSkuInfoResponseOutput
}

type IotHubSkuInfoResponseArgs struct {
	Capacity pulumi.Float64PtrInput `pulumi:"capacity"`
	Name     pulumi.StringInput     `pulumi:"name"`
	Tier     pulumi.StringInput     `pulumi:"tier"`
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

func (o IotHubSkuInfoResponseOutput) Capacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v IotHubSkuInfoResponse) *float64 { return v.Capacity }).(pulumi.Float64PtrOutput)
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
		return v.Capacity
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

type RouteProperties struct {
	Condition     *string  `pulumi:"condition"`
	EndpointNames []string `pulumi:"endpointNames"`
	IsEnabled     bool     `pulumi:"isEnabled"`
	Name          string   `pulumi:"name"`
	Source        string   `pulumi:"source"`
}





type RoutePropertiesInput interface {
	pulumi.Input

	ToRoutePropertiesOutput() RoutePropertiesOutput
	ToRoutePropertiesOutputWithContext(context.Context) RoutePropertiesOutput
}

type RoutePropertiesArgs struct {
	Condition     pulumi.StringPtrInput   `pulumi:"condition"`
	EndpointNames pulumi.StringArrayInput `pulumi:"endpointNames"`
	IsEnabled     pulumi.BoolInput        `pulumi:"isEnabled"`
	Name          pulumi.StringInput      `pulumi:"name"`
	Source        pulumi.StringInput      `pulumi:"source"`
}

func (RoutePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteProperties)(nil)).Elem()
}

func (i RoutePropertiesArgs) ToRoutePropertiesOutput() RoutePropertiesOutput {
	return i.ToRoutePropertiesOutputWithContext(context.Background())
}

func (i RoutePropertiesArgs) ToRoutePropertiesOutputWithContext(ctx context.Context) RoutePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePropertiesOutput)
}





type RoutePropertiesArrayInput interface {
	pulumi.Input

	ToRoutePropertiesArrayOutput() RoutePropertiesArrayOutput
	ToRoutePropertiesArrayOutputWithContext(context.Context) RoutePropertiesArrayOutput
}

type RoutePropertiesArray []RoutePropertiesInput

func (RoutePropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteProperties)(nil)).Elem()
}

func (i RoutePropertiesArray) ToRoutePropertiesArrayOutput() RoutePropertiesArrayOutput {
	return i.ToRoutePropertiesArrayOutputWithContext(context.Background())
}

func (i RoutePropertiesArray) ToRoutePropertiesArrayOutputWithContext(ctx context.Context) RoutePropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePropertiesArrayOutput)
}

type RoutePropertiesOutput struct{ *pulumi.OutputState }

func (RoutePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RouteProperties)(nil)).Elem()
}

func (o RoutePropertiesOutput) ToRoutePropertiesOutput() RoutePropertiesOutput {
	return o
}

func (o RoutePropertiesOutput) ToRoutePropertiesOutputWithContext(ctx context.Context) RoutePropertiesOutput {
	return o
}

func (o RoutePropertiesOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RouteProperties) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o RoutePropertiesOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RouteProperties) []string { return v.EndpointNames }).(pulumi.StringArrayOutput)
}

func (o RoutePropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RouteProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o RoutePropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RouteProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutePropertiesOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v RouteProperties) string { return v.Source }).(pulumi.StringOutput)
}

type RoutePropertiesArrayOutput struct{ *pulumi.OutputState }

func (RoutePropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RouteProperties)(nil)).Elem()
}

func (o RoutePropertiesArrayOutput) ToRoutePropertiesArrayOutput() RoutePropertiesArrayOutput {
	return o
}

func (o RoutePropertiesArrayOutput) ToRoutePropertiesArrayOutputWithContext(ctx context.Context) RoutePropertiesArrayOutput {
	return o
}

func (o RoutePropertiesArrayOutput) Index(i pulumi.IntInput) RoutePropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RouteProperties {
		return vs[0].([]RouteProperties)[vs[1].(int)]
	}).(RoutePropertiesOutput)
}

type RoutePropertiesResponse struct {
	Condition     *string  `pulumi:"condition"`
	EndpointNames []string `pulumi:"endpointNames"`
	IsEnabled     bool     `pulumi:"isEnabled"`
	Name          string   `pulumi:"name"`
	Source        string   `pulumi:"source"`
}





type RoutePropertiesResponseInput interface {
	pulumi.Input

	ToRoutePropertiesResponseOutput() RoutePropertiesResponseOutput
	ToRoutePropertiesResponseOutputWithContext(context.Context) RoutePropertiesResponseOutput
}

type RoutePropertiesResponseArgs struct {
	Condition     pulumi.StringPtrInput   `pulumi:"condition"`
	EndpointNames pulumi.StringArrayInput `pulumi:"endpointNames"`
	IsEnabled     pulumi.BoolInput        `pulumi:"isEnabled"`
	Name          pulumi.StringInput      `pulumi:"name"`
	Source        pulumi.StringInput      `pulumi:"source"`
}

func (RoutePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutePropertiesResponse)(nil)).Elem()
}

func (i RoutePropertiesResponseArgs) ToRoutePropertiesResponseOutput() RoutePropertiesResponseOutput {
	return i.ToRoutePropertiesResponseOutputWithContext(context.Background())
}

func (i RoutePropertiesResponseArgs) ToRoutePropertiesResponseOutputWithContext(ctx context.Context) RoutePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePropertiesResponseOutput)
}





type RoutePropertiesResponseArrayInput interface {
	pulumi.Input

	ToRoutePropertiesResponseArrayOutput() RoutePropertiesResponseArrayOutput
	ToRoutePropertiesResponseArrayOutputWithContext(context.Context) RoutePropertiesResponseArrayOutput
}

type RoutePropertiesResponseArray []RoutePropertiesResponseInput

func (RoutePropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutePropertiesResponse)(nil)).Elem()
}

func (i RoutePropertiesResponseArray) ToRoutePropertiesResponseArrayOutput() RoutePropertiesResponseArrayOutput {
	return i.ToRoutePropertiesResponseArrayOutputWithContext(context.Background())
}

func (i RoutePropertiesResponseArray) ToRoutePropertiesResponseArrayOutputWithContext(ctx context.Context) RoutePropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePropertiesResponseArrayOutput)
}

type RoutePropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutePropertiesResponse)(nil)).Elem()
}

func (o RoutePropertiesResponseOutput) ToRoutePropertiesResponseOutput() RoutePropertiesResponseOutput {
	return o
}

func (o RoutePropertiesResponseOutput) ToRoutePropertiesResponseOutputWithContext(ctx context.Context) RoutePropertiesResponseOutput {
	return o
}

func (o RoutePropertiesResponseOutput) Condition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutePropertiesResponse) *string { return v.Condition }).(pulumi.StringPtrOutput)
}

func (o RoutePropertiesResponseOutput) EndpointNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RoutePropertiesResponse) []string { return v.EndpointNames }).(pulumi.StringArrayOutput)
}

func (o RoutePropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RoutePropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o RoutePropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutePropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutePropertiesResponseOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v RoutePropertiesResponse) string { return v.Source }).(pulumi.StringOutput)
}

type RoutePropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutePropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutePropertiesResponse)(nil)).Elem()
}

func (o RoutePropertiesResponseArrayOutput) ToRoutePropertiesResponseArrayOutput() RoutePropertiesResponseArrayOutput {
	return o
}

func (o RoutePropertiesResponseArrayOutput) ToRoutePropertiesResponseArrayOutputWithContext(ctx context.Context) RoutePropertiesResponseArrayOutput {
	return o
}

func (o RoutePropertiesResponseArrayOutput) Index(i pulumi.IntInput) RoutePropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutePropertiesResponse {
		return vs[0].([]RoutePropertiesResponse)[vs[1].(int)]
	}).(RoutePropertiesResponseOutput)
}

type RoutingEndpoints struct {
	EventHubs         []RoutingEventHubProperties                `pulumi:"eventHubs"`
	ServiceBusQueues  []RoutingServiceBusQueueEndpointProperties `pulumi:"serviceBusQueues"`
	ServiceBusTopics  []RoutingServiceBusTopicEndpointProperties `pulumi:"serviceBusTopics"`
	StorageContainers []RoutingStorageContainerProperties        `pulumi:"storageContainers"`
}





type RoutingEndpointsInput interface {
	pulumi.Input

	ToRoutingEndpointsOutput() RoutingEndpointsOutput
	ToRoutingEndpointsOutputWithContext(context.Context) RoutingEndpointsOutput
}

type RoutingEndpointsArgs struct {
	EventHubs         RoutingEventHubPropertiesArrayInput                `pulumi:"eventHubs"`
	ServiceBusQueues  RoutingServiceBusQueueEndpointPropertiesArrayInput `pulumi:"serviceBusQueues"`
	ServiceBusTopics  RoutingServiceBusTopicEndpointPropertiesArrayInput `pulumi:"serviceBusTopics"`
	StorageContainers RoutingStorageContainerPropertiesArrayInput        `pulumi:"storageContainers"`
}

func (RoutingEndpointsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEndpoints)(nil)).Elem()
}

func (i RoutingEndpointsArgs) ToRoutingEndpointsOutput() RoutingEndpointsOutput {
	return i.ToRoutingEndpointsOutputWithContext(context.Background())
}

func (i RoutingEndpointsArgs) ToRoutingEndpointsOutputWithContext(ctx context.Context) RoutingEndpointsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEndpointsOutput)
}

func (i RoutingEndpointsArgs) ToRoutingEndpointsPtrOutput() RoutingEndpointsPtrOutput {
	return i.ToRoutingEndpointsPtrOutputWithContext(context.Background())
}

func (i RoutingEndpointsArgs) ToRoutingEndpointsPtrOutputWithContext(ctx context.Context) RoutingEndpointsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEndpointsOutput).ToRoutingEndpointsPtrOutputWithContext(ctx)
}









type RoutingEndpointsPtrInput interface {
	pulumi.Input

	ToRoutingEndpointsPtrOutput() RoutingEndpointsPtrOutput
	ToRoutingEndpointsPtrOutputWithContext(context.Context) RoutingEndpointsPtrOutput
}

type routingEndpointsPtrType RoutingEndpointsArgs

func RoutingEndpointsPtr(v *RoutingEndpointsArgs) RoutingEndpointsPtrInput {
	return (*routingEndpointsPtrType)(v)
}

func (*routingEndpointsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingEndpoints)(nil)).Elem()
}

func (i *routingEndpointsPtrType) ToRoutingEndpointsPtrOutput() RoutingEndpointsPtrOutput {
	return i.ToRoutingEndpointsPtrOutputWithContext(context.Background())
}

func (i *routingEndpointsPtrType) ToRoutingEndpointsPtrOutputWithContext(ctx context.Context) RoutingEndpointsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEndpointsPtrOutput)
}

type RoutingEndpointsOutput struct{ *pulumi.OutputState }

func (RoutingEndpointsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEndpoints)(nil)).Elem()
}

func (o RoutingEndpointsOutput) ToRoutingEndpointsOutput() RoutingEndpointsOutput {
	return o
}

func (o RoutingEndpointsOutput) ToRoutingEndpointsOutputWithContext(ctx context.Context) RoutingEndpointsOutput {
	return o
}

func (o RoutingEndpointsOutput) ToRoutingEndpointsPtrOutput() RoutingEndpointsPtrOutput {
	return o.ToRoutingEndpointsPtrOutputWithContext(context.Background())
}

func (o RoutingEndpointsOutput) ToRoutingEndpointsPtrOutputWithContext(ctx context.Context) RoutingEndpointsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingEndpoints) *RoutingEndpoints {
		return &v
	}).(RoutingEndpointsPtrOutput)
}

func (o RoutingEndpointsOutput) EventHubs() RoutingEventHubPropertiesArrayOutput {
	return o.ApplyT(func(v RoutingEndpoints) []RoutingEventHubProperties { return v.EventHubs }).(RoutingEventHubPropertiesArrayOutput)
}

func (o RoutingEndpointsOutput) ServiceBusQueues() RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return o.ApplyT(func(v RoutingEndpoints) []RoutingServiceBusQueueEndpointProperties { return v.ServiceBusQueues }).(RoutingServiceBusQueueEndpointPropertiesArrayOutput)
}

func (o RoutingEndpointsOutput) ServiceBusTopics() RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return o.ApplyT(func(v RoutingEndpoints) []RoutingServiceBusTopicEndpointProperties { return v.ServiceBusTopics }).(RoutingServiceBusTopicEndpointPropertiesArrayOutput)
}

func (o RoutingEndpointsOutput) StorageContainers() RoutingStorageContainerPropertiesArrayOutput {
	return o.ApplyT(func(v RoutingEndpoints) []RoutingStorageContainerProperties { return v.StorageContainers }).(RoutingStorageContainerPropertiesArrayOutput)
}

type RoutingEndpointsPtrOutput struct{ *pulumi.OutputState }

func (RoutingEndpointsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingEndpoints)(nil)).Elem()
}

func (o RoutingEndpointsPtrOutput) ToRoutingEndpointsPtrOutput() RoutingEndpointsPtrOutput {
	return o
}

func (o RoutingEndpointsPtrOutput) ToRoutingEndpointsPtrOutputWithContext(ctx context.Context) RoutingEndpointsPtrOutput {
	return o
}

func (o RoutingEndpointsPtrOutput) Elem() RoutingEndpointsOutput {
	return o.ApplyT(func(v *RoutingEndpoints) RoutingEndpoints {
		if v != nil {
			return *v
		}
		var ret RoutingEndpoints
		return ret
	}).(RoutingEndpointsOutput)
}

func (o RoutingEndpointsPtrOutput) EventHubs() RoutingEventHubPropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingEndpoints) []RoutingEventHubProperties {
		if v == nil {
			return nil
		}
		return v.EventHubs
	}).(RoutingEventHubPropertiesArrayOutput)
}

func (o RoutingEndpointsPtrOutput) ServiceBusQueues() RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingEndpoints) []RoutingServiceBusQueueEndpointProperties {
		if v == nil {
			return nil
		}
		return v.ServiceBusQueues
	}).(RoutingServiceBusQueueEndpointPropertiesArrayOutput)
}

func (o RoutingEndpointsPtrOutput) ServiceBusTopics() RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingEndpoints) []RoutingServiceBusTopicEndpointProperties {
		if v == nil {
			return nil
		}
		return v.ServiceBusTopics
	}).(RoutingServiceBusTopicEndpointPropertiesArrayOutput)
}

func (o RoutingEndpointsPtrOutput) StorageContainers() RoutingStorageContainerPropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingEndpoints) []RoutingStorageContainerProperties {
		if v == nil {
			return nil
		}
		return v.StorageContainers
	}).(RoutingStorageContainerPropertiesArrayOutput)
}

type RoutingEndpointsResponse struct {
	EventHubs         []RoutingEventHubPropertiesResponse                `pulumi:"eventHubs"`
	ServiceBusQueues  []RoutingServiceBusQueueEndpointPropertiesResponse `pulumi:"serviceBusQueues"`
	ServiceBusTopics  []RoutingServiceBusTopicEndpointPropertiesResponse `pulumi:"serviceBusTopics"`
	StorageContainers []RoutingStorageContainerPropertiesResponse        `pulumi:"storageContainers"`
}





type RoutingEndpointsResponseInput interface {
	pulumi.Input

	ToRoutingEndpointsResponseOutput() RoutingEndpointsResponseOutput
	ToRoutingEndpointsResponseOutputWithContext(context.Context) RoutingEndpointsResponseOutput
}

type RoutingEndpointsResponseArgs struct {
	EventHubs         RoutingEventHubPropertiesResponseArrayInput                `pulumi:"eventHubs"`
	ServiceBusQueues  RoutingServiceBusQueueEndpointPropertiesResponseArrayInput `pulumi:"serviceBusQueues"`
	ServiceBusTopics  RoutingServiceBusTopicEndpointPropertiesResponseArrayInput `pulumi:"serviceBusTopics"`
	StorageContainers RoutingStorageContainerPropertiesResponseArrayInput        `pulumi:"storageContainers"`
}

func (RoutingEndpointsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEndpointsResponse)(nil)).Elem()
}

func (i RoutingEndpointsResponseArgs) ToRoutingEndpointsResponseOutput() RoutingEndpointsResponseOutput {
	return i.ToRoutingEndpointsResponseOutputWithContext(context.Background())
}

func (i RoutingEndpointsResponseArgs) ToRoutingEndpointsResponseOutputWithContext(ctx context.Context) RoutingEndpointsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEndpointsResponseOutput)
}

func (i RoutingEndpointsResponseArgs) ToRoutingEndpointsResponsePtrOutput() RoutingEndpointsResponsePtrOutput {
	return i.ToRoutingEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i RoutingEndpointsResponseArgs) ToRoutingEndpointsResponsePtrOutputWithContext(ctx context.Context) RoutingEndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEndpointsResponseOutput).ToRoutingEndpointsResponsePtrOutputWithContext(ctx)
}









type RoutingEndpointsResponsePtrInput interface {
	pulumi.Input

	ToRoutingEndpointsResponsePtrOutput() RoutingEndpointsResponsePtrOutput
	ToRoutingEndpointsResponsePtrOutputWithContext(context.Context) RoutingEndpointsResponsePtrOutput
}

type routingEndpointsResponsePtrType RoutingEndpointsResponseArgs

func RoutingEndpointsResponsePtr(v *RoutingEndpointsResponseArgs) RoutingEndpointsResponsePtrInput {
	return (*routingEndpointsResponsePtrType)(v)
}

func (*routingEndpointsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingEndpointsResponse)(nil)).Elem()
}

func (i *routingEndpointsResponsePtrType) ToRoutingEndpointsResponsePtrOutput() RoutingEndpointsResponsePtrOutput {
	return i.ToRoutingEndpointsResponsePtrOutputWithContext(context.Background())
}

func (i *routingEndpointsResponsePtrType) ToRoutingEndpointsResponsePtrOutputWithContext(ctx context.Context) RoutingEndpointsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEndpointsResponsePtrOutput)
}

type RoutingEndpointsResponseOutput struct{ *pulumi.OutputState }

func (RoutingEndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEndpointsResponse)(nil)).Elem()
}

func (o RoutingEndpointsResponseOutput) ToRoutingEndpointsResponseOutput() RoutingEndpointsResponseOutput {
	return o
}

func (o RoutingEndpointsResponseOutput) ToRoutingEndpointsResponseOutputWithContext(ctx context.Context) RoutingEndpointsResponseOutput {
	return o
}

func (o RoutingEndpointsResponseOutput) ToRoutingEndpointsResponsePtrOutput() RoutingEndpointsResponsePtrOutput {
	return o.ToRoutingEndpointsResponsePtrOutputWithContext(context.Background())
}

func (o RoutingEndpointsResponseOutput) ToRoutingEndpointsResponsePtrOutputWithContext(ctx context.Context) RoutingEndpointsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingEndpointsResponse) *RoutingEndpointsResponse {
		return &v
	}).(RoutingEndpointsResponsePtrOutput)
}

func (o RoutingEndpointsResponseOutput) EventHubs() RoutingEventHubPropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingEndpointsResponse) []RoutingEventHubPropertiesResponse { return v.EventHubs }).(RoutingEventHubPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponseOutput) ServiceBusQueues() RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingEndpointsResponse) []RoutingServiceBusQueueEndpointPropertiesResponse {
		return v.ServiceBusQueues
	}).(RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponseOutput) ServiceBusTopics() RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingEndpointsResponse) []RoutingServiceBusTopicEndpointPropertiesResponse {
		return v.ServiceBusTopics
	}).(RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponseOutput) StorageContainers() RoutingStorageContainerPropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingEndpointsResponse) []RoutingStorageContainerPropertiesResponse {
		return v.StorageContainers
	}).(RoutingStorageContainerPropertiesResponseArrayOutput)
}

type RoutingEndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (RoutingEndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingEndpointsResponse)(nil)).Elem()
}

func (o RoutingEndpointsResponsePtrOutput) ToRoutingEndpointsResponsePtrOutput() RoutingEndpointsResponsePtrOutput {
	return o
}

func (o RoutingEndpointsResponsePtrOutput) ToRoutingEndpointsResponsePtrOutputWithContext(ctx context.Context) RoutingEndpointsResponsePtrOutput {
	return o
}

func (o RoutingEndpointsResponsePtrOutput) Elem() RoutingEndpointsResponseOutput {
	return o.ApplyT(func(v *RoutingEndpointsResponse) RoutingEndpointsResponse {
		if v != nil {
			return *v
		}
		var ret RoutingEndpointsResponse
		return ret
	}).(RoutingEndpointsResponseOutput)
}

func (o RoutingEndpointsResponsePtrOutput) EventHubs() RoutingEventHubPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingEndpointsResponse) []RoutingEventHubPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.EventHubs
	}).(RoutingEventHubPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponsePtrOutput) ServiceBusQueues() RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingEndpointsResponse) []RoutingServiceBusQueueEndpointPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ServiceBusQueues
	}).(RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponsePtrOutput) ServiceBusTopics() RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingEndpointsResponse) []RoutingServiceBusTopicEndpointPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.ServiceBusTopics
	}).(RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput)
}

func (o RoutingEndpointsResponsePtrOutput) StorageContainers() RoutingStorageContainerPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingEndpointsResponse) []RoutingStorageContainerPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.StorageContainers
	}).(RoutingStorageContainerPropertiesResponseArrayOutput)
}

type RoutingEventHubProperties struct {
	ConnectionString string  `pulumi:"connectionString"`
	Name             string  `pulumi:"name"`
	ResourceGroup    *string `pulumi:"resourceGroup"`
	SubscriptionId   *string `pulumi:"subscriptionId"`
}





type RoutingEventHubPropertiesInput interface {
	pulumi.Input

	ToRoutingEventHubPropertiesOutput() RoutingEventHubPropertiesOutput
	ToRoutingEventHubPropertiesOutputWithContext(context.Context) RoutingEventHubPropertiesOutput
}

type RoutingEventHubPropertiesArgs struct {
	ConnectionString pulumi.StringInput    `pulumi:"connectionString"`
	Name             pulumi.StringInput    `pulumi:"name"`
	ResourceGroup    pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId   pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingEventHubPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEventHubProperties)(nil)).Elem()
}

func (i RoutingEventHubPropertiesArgs) ToRoutingEventHubPropertiesOutput() RoutingEventHubPropertiesOutput {
	return i.ToRoutingEventHubPropertiesOutputWithContext(context.Background())
}

func (i RoutingEventHubPropertiesArgs) ToRoutingEventHubPropertiesOutputWithContext(ctx context.Context) RoutingEventHubPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEventHubPropertiesOutput)
}





type RoutingEventHubPropertiesArrayInput interface {
	pulumi.Input

	ToRoutingEventHubPropertiesArrayOutput() RoutingEventHubPropertiesArrayOutput
	ToRoutingEventHubPropertiesArrayOutputWithContext(context.Context) RoutingEventHubPropertiesArrayOutput
}

type RoutingEventHubPropertiesArray []RoutingEventHubPropertiesInput

func (RoutingEventHubPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingEventHubProperties)(nil)).Elem()
}

func (i RoutingEventHubPropertiesArray) ToRoutingEventHubPropertiesArrayOutput() RoutingEventHubPropertiesArrayOutput {
	return i.ToRoutingEventHubPropertiesArrayOutputWithContext(context.Background())
}

func (i RoutingEventHubPropertiesArray) ToRoutingEventHubPropertiesArrayOutputWithContext(ctx context.Context) RoutingEventHubPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEventHubPropertiesArrayOutput)
}

type RoutingEventHubPropertiesOutput struct{ *pulumi.OutputState }

func (RoutingEventHubPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEventHubProperties)(nil)).Elem()
}

func (o RoutingEventHubPropertiesOutput) ToRoutingEventHubPropertiesOutput() RoutingEventHubPropertiesOutput {
	return o
}

func (o RoutingEventHubPropertiesOutput) ToRoutingEventHubPropertiesOutputWithContext(ctx context.Context) RoutingEventHubPropertiesOutput {
	return o
}

func (o RoutingEventHubPropertiesOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o RoutingEventHubPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingEventHubPropertiesOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingEventHubPropertiesArrayOutput struct{ *pulumi.OutputState }

func (RoutingEventHubPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingEventHubProperties)(nil)).Elem()
}

func (o RoutingEventHubPropertiesArrayOutput) ToRoutingEventHubPropertiesArrayOutput() RoutingEventHubPropertiesArrayOutput {
	return o
}

func (o RoutingEventHubPropertiesArrayOutput) ToRoutingEventHubPropertiesArrayOutputWithContext(ctx context.Context) RoutingEventHubPropertiesArrayOutput {
	return o
}

func (o RoutingEventHubPropertiesArrayOutput) Index(i pulumi.IntInput) RoutingEventHubPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingEventHubProperties {
		return vs[0].([]RoutingEventHubProperties)[vs[1].(int)]
	}).(RoutingEventHubPropertiesOutput)
}

type RoutingEventHubPropertiesResponse struct {
	ConnectionString string  `pulumi:"connectionString"`
	Name             string  `pulumi:"name"`
	ResourceGroup    *string `pulumi:"resourceGroup"`
	SubscriptionId   *string `pulumi:"subscriptionId"`
}





type RoutingEventHubPropertiesResponseInput interface {
	pulumi.Input

	ToRoutingEventHubPropertiesResponseOutput() RoutingEventHubPropertiesResponseOutput
	ToRoutingEventHubPropertiesResponseOutputWithContext(context.Context) RoutingEventHubPropertiesResponseOutput
}

type RoutingEventHubPropertiesResponseArgs struct {
	ConnectionString pulumi.StringInput    `pulumi:"connectionString"`
	Name             pulumi.StringInput    `pulumi:"name"`
	ResourceGroup    pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId   pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingEventHubPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEventHubPropertiesResponse)(nil)).Elem()
}

func (i RoutingEventHubPropertiesResponseArgs) ToRoutingEventHubPropertiesResponseOutput() RoutingEventHubPropertiesResponseOutput {
	return i.ToRoutingEventHubPropertiesResponseOutputWithContext(context.Background())
}

func (i RoutingEventHubPropertiesResponseArgs) ToRoutingEventHubPropertiesResponseOutputWithContext(ctx context.Context) RoutingEventHubPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEventHubPropertiesResponseOutput)
}





type RoutingEventHubPropertiesResponseArrayInput interface {
	pulumi.Input

	ToRoutingEventHubPropertiesResponseArrayOutput() RoutingEventHubPropertiesResponseArrayOutput
	ToRoutingEventHubPropertiesResponseArrayOutputWithContext(context.Context) RoutingEventHubPropertiesResponseArrayOutput
}

type RoutingEventHubPropertiesResponseArray []RoutingEventHubPropertiesResponseInput

func (RoutingEventHubPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingEventHubPropertiesResponse)(nil)).Elem()
}

func (i RoutingEventHubPropertiesResponseArray) ToRoutingEventHubPropertiesResponseArrayOutput() RoutingEventHubPropertiesResponseArrayOutput {
	return i.ToRoutingEventHubPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i RoutingEventHubPropertiesResponseArray) ToRoutingEventHubPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingEventHubPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingEventHubPropertiesResponseArrayOutput)
}

type RoutingEventHubPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutingEventHubPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingEventHubPropertiesResponse)(nil)).Elem()
}

func (o RoutingEventHubPropertiesResponseOutput) ToRoutingEventHubPropertiesResponseOutput() RoutingEventHubPropertiesResponseOutput {
	return o
}

func (o RoutingEventHubPropertiesResponseOutput) ToRoutingEventHubPropertiesResponseOutputWithContext(ctx context.Context) RoutingEventHubPropertiesResponseOutput {
	return o
}

func (o RoutingEventHubPropertiesResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o RoutingEventHubPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingEventHubPropertiesResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingEventHubPropertiesResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingEventHubPropertiesResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingEventHubPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingEventHubPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingEventHubPropertiesResponse)(nil)).Elem()
}

func (o RoutingEventHubPropertiesResponseArrayOutput) ToRoutingEventHubPropertiesResponseArrayOutput() RoutingEventHubPropertiesResponseArrayOutput {
	return o
}

func (o RoutingEventHubPropertiesResponseArrayOutput) ToRoutingEventHubPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingEventHubPropertiesResponseArrayOutput {
	return o
}

func (o RoutingEventHubPropertiesResponseArrayOutput) Index(i pulumi.IntInput) RoutingEventHubPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingEventHubPropertiesResponse {
		return vs[0].([]RoutingEventHubPropertiesResponse)[vs[1].(int)]
	}).(RoutingEventHubPropertiesResponseOutput)
}

type RoutingProperties struct {
	Endpoints     *RoutingEndpoints        `pulumi:"endpoints"`
	FallbackRoute *FallbackRouteProperties `pulumi:"fallbackRoute"`
	Routes        []RouteProperties        `pulumi:"routes"`
}





type RoutingPropertiesInput interface {
	pulumi.Input

	ToRoutingPropertiesOutput() RoutingPropertiesOutput
	ToRoutingPropertiesOutputWithContext(context.Context) RoutingPropertiesOutput
}

type RoutingPropertiesArgs struct {
	Endpoints     RoutingEndpointsPtrInput        `pulumi:"endpoints"`
	FallbackRoute FallbackRoutePropertiesPtrInput `pulumi:"fallbackRoute"`
	Routes        RoutePropertiesArrayInput       `pulumi:"routes"`
}

func (RoutingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingProperties)(nil)).Elem()
}

func (i RoutingPropertiesArgs) ToRoutingPropertiesOutput() RoutingPropertiesOutput {
	return i.ToRoutingPropertiesOutputWithContext(context.Background())
}

func (i RoutingPropertiesArgs) ToRoutingPropertiesOutputWithContext(ctx context.Context) RoutingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPropertiesOutput)
}

func (i RoutingPropertiesArgs) ToRoutingPropertiesPtrOutput() RoutingPropertiesPtrOutput {
	return i.ToRoutingPropertiesPtrOutputWithContext(context.Background())
}

func (i RoutingPropertiesArgs) ToRoutingPropertiesPtrOutputWithContext(ctx context.Context) RoutingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPropertiesOutput).ToRoutingPropertiesPtrOutputWithContext(ctx)
}









type RoutingPropertiesPtrInput interface {
	pulumi.Input

	ToRoutingPropertiesPtrOutput() RoutingPropertiesPtrOutput
	ToRoutingPropertiesPtrOutputWithContext(context.Context) RoutingPropertiesPtrOutput
}

type routingPropertiesPtrType RoutingPropertiesArgs

func RoutingPropertiesPtr(v *RoutingPropertiesArgs) RoutingPropertiesPtrInput {
	return (*routingPropertiesPtrType)(v)
}

func (*routingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingProperties)(nil)).Elem()
}

func (i *routingPropertiesPtrType) ToRoutingPropertiesPtrOutput() RoutingPropertiesPtrOutput {
	return i.ToRoutingPropertiesPtrOutputWithContext(context.Background())
}

func (i *routingPropertiesPtrType) ToRoutingPropertiesPtrOutputWithContext(ctx context.Context) RoutingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPropertiesPtrOutput)
}

type RoutingPropertiesOutput struct{ *pulumi.OutputState }

func (RoutingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingProperties)(nil)).Elem()
}

func (o RoutingPropertiesOutput) ToRoutingPropertiesOutput() RoutingPropertiesOutput {
	return o
}

func (o RoutingPropertiesOutput) ToRoutingPropertiesOutputWithContext(ctx context.Context) RoutingPropertiesOutput {
	return o
}

func (o RoutingPropertiesOutput) ToRoutingPropertiesPtrOutput() RoutingPropertiesPtrOutput {
	return o.ToRoutingPropertiesPtrOutputWithContext(context.Background())
}

func (o RoutingPropertiesOutput) ToRoutingPropertiesPtrOutputWithContext(ctx context.Context) RoutingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingProperties) *RoutingProperties {
		return &v
	}).(RoutingPropertiesPtrOutput)
}

func (o RoutingPropertiesOutput) Endpoints() RoutingEndpointsPtrOutput {
	return o.ApplyT(func(v RoutingProperties) *RoutingEndpoints { return v.Endpoints }).(RoutingEndpointsPtrOutput)
}

func (o RoutingPropertiesOutput) FallbackRoute() FallbackRoutePropertiesPtrOutput {
	return o.ApplyT(func(v RoutingProperties) *FallbackRouteProperties { return v.FallbackRoute }).(FallbackRoutePropertiesPtrOutput)
}

func (o RoutingPropertiesOutput) Routes() RoutePropertiesArrayOutput {
	return o.ApplyT(func(v RoutingProperties) []RouteProperties { return v.Routes }).(RoutePropertiesArrayOutput)
}

type RoutingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (RoutingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingProperties)(nil)).Elem()
}

func (o RoutingPropertiesPtrOutput) ToRoutingPropertiesPtrOutput() RoutingPropertiesPtrOutput {
	return o
}

func (o RoutingPropertiesPtrOutput) ToRoutingPropertiesPtrOutputWithContext(ctx context.Context) RoutingPropertiesPtrOutput {
	return o
}

func (o RoutingPropertiesPtrOutput) Elem() RoutingPropertiesOutput {
	return o.ApplyT(func(v *RoutingProperties) RoutingProperties {
		if v != nil {
			return *v
		}
		var ret RoutingProperties
		return ret
	}).(RoutingPropertiesOutput)
}

func (o RoutingPropertiesPtrOutput) Endpoints() RoutingEndpointsPtrOutput {
	return o.ApplyT(func(v *RoutingProperties) *RoutingEndpoints {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(RoutingEndpointsPtrOutput)
}

func (o RoutingPropertiesPtrOutput) FallbackRoute() FallbackRoutePropertiesPtrOutput {
	return o.ApplyT(func(v *RoutingProperties) *FallbackRouteProperties {
		if v == nil {
			return nil
		}
		return v.FallbackRoute
	}).(FallbackRoutePropertiesPtrOutput)
}

func (o RoutingPropertiesPtrOutput) Routes() RoutePropertiesArrayOutput {
	return o.ApplyT(func(v *RoutingProperties) []RouteProperties {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(RoutePropertiesArrayOutput)
}

type RoutingPropertiesResponse struct {
	Endpoints     *RoutingEndpointsResponse        `pulumi:"endpoints"`
	FallbackRoute *FallbackRoutePropertiesResponse `pulumi:"fallbackRoute"`
	Routes        []RoutePropertiesResponse        `pulumi:"routes"`
}





type RoutingPropertiesResponseInput interface {
	pulumi.Input

	ToRoutingPropertiesResponseOutput() RoutingPropertiesResponseOutput
	ToRoutingPropertiesResponseOutputWithContext(context.Context) RoutingPropertiesResponseOutput
}

type RoutingPropertiesResponseArgs struct {
	Endpoints     RoutingEndpointsResponsePtrInput        `pulumi:"endpoints"`
	FallbackRoute FallbackRoutePropertiesResponsePtrInput `pulumi:"fallbackRoute"`
	Routes        RoutePropertiesResponseArrayInput       `pulumi:"routes"`
}

func (RoutingPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingPropertiesResponse)(nil)).Elem()
}

func (i RoutingPropertiesResponseArgs) ToRoutingPropertiesResponseOutput() RoutingPropertiesResponseOutput {
	return i.ToRoutingPropertiesResponseOutputWithContext(context.Background())
}

func (i RoutingPropertiesResponseArgs) ToRoutingPropertiesResponseOutputWithContext(ctx context.Context) RoutingPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPropertiesResponseOutput)
}

func (i RoutingPropertiesResponseArgs) ToRoutingPropertiesResponsePtrOutput() RoutingPropertiesResponsePtrOutput {
	return i.ToRoutingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i RoutingPropertiesResponseArgs) ToRoutingPropertiesResponsePtrOutputWithContext(ctx context.Context) RoutingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPropertiesResponseOutput).ToRoutingPropertiesResponsePtrOutputWithContext(ctx)
}









type RoutingPropertiesResponsePtrInput interface {
	pulumi.Input

	ToRoutingPropertiesResponsePtrOutput() RoutingPropertiesResponsePtrOutput
	ToRoutingPropertiesResponsePtrOutputWithContext(context.Context) RoutingPropertiesResponsePtrOutput
}

type routingPropertiesResponsePtrType RoutingPropertiesResponseArgs

func RoutingPropertiesResponsePtr(v *RoutingPropertiesResponseArgs) RoutingPropertiesResponsePtrInput {
	return (*routingPropertiesResponsePtrType)(v)
}

func (*routingPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingPropertiesResponse)(nil)).Elem()
}

func (i *routingPropertiesResponsePtrType) ToRoutingPropertiesResponsePtrOutput() RoutingPropertiesResponsePtrOutput {
	return i.ToRoutingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *routingPropertiesResponsePtrType) ToRoutingPropertiesResponsePtrOutputWithContext(ctx context.Context) RoutingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingPropertiesResponsePtrOutput)
}

type RoutingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingPropertiesResponse)(nil)).Elem()
}

func (o RoutingPropertiesResponseOutput) ToRoutingPropertiesResponseOutput() RoutingPropertiesResponseOutput {
	return o
}

func (o RoutingPropertiesResponseOutput) ToRoutingPropertiesResponseOutputWithContext(ctx context.Context) RoutingPropertiesResponseOutput {
	return o
}

func (o RoutingPropertiesResponseOutput) ToRoutingPropertiesResponsePtrOutput() RoutingPropertiesResponsePtrOutput {
	return o.ToRoutingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o RoutingPropertiesResponseOutput) ToRoutingPropertiesResponsePtrOutputWithContext(ctx context.Context) RoutingPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RoutingPropertiesResponse) *RoutingPropertiesResponse {
		return &v
	}).(RoutingPropertiesResponsePtrOutput)
}

func (o RoutingPropertiesResponseOutput) Endpoints() RoutingEndpointsResponsePtrOutput {
	return o.ApplyT(func(v RoutingPropertiesResponse) *RoutingEndpointsResponse { return v.Endpoints }).(RoutingEndpointsResponsePtrOutput)
}

func (o RoutingPropertiesResponseOutput) FallbackRoute() FallbackRoutePropertiesResponsePtrOutput {
	return o.ApplyT(func(v RoutingPropertiesResponse) *FallbackRoutePropertiesResponse { return v.FallbackRoute }).(FallbackRoutePropertiesResponsePtrOutput)
}

func (o RoutingPropertiesResponseOutput) Routes() RoutePropertiesResponseArrayOutput {
	return o.ApplyT(func(v RoutingPropertiesResponse) []RoutePropertiesResponse { return v.Routes }).(RoutePropertiesResponseArrayOutput)
}

type RoutingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (RoutingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoutingPropertiesResponse)(nil)).Elem()
}

func (o RoutingPropertiesResponsePtrOutput) ToRoutingPropertiesResponsePtrOutput() RoutingPropertiesResponsePtrOutput {
	return o
}

func (o RoutingPropertiesResponsePtrOutput) ToRoutingPropertiesResponsePtrOutputWithContext(ctx context.Context) RoutingPropertiesResponsePtrOutput {
	return o
}

func (o RoutingPropertiesResponsePtrOutput) Elem() RoutingPropertiesResponseOutput {
	return o.ApplyT(func(v *RoutingPropertiesResponse) RoutingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret RoutingPropertiesResponse
		return ret
	}).(RoutingPropertiesResponseOutput)
}

func (o RoutingPropertiesResponsePtrOutput) Endpoints() RoutingEndpointsResponsePtrOutput {
	return o.ApplyT(func(v *RoutingPropertiesResponse) *RoutingEndpointsResponse {
		if v == nil {
			return nil
		}
		return v.Endpoints
	}).(RoutingEndpointsResponsePtrOutput)
}

func (o RoutingPropertiesResponsePtrOutput) FallbackRoute() FallbackRoutePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *RoutingPropertiesResponse) *FallbackRoutePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.FallbackRoute
	}).(FallbackRoutePropertiesResponsePtrOutput)
}

func (o RoutingPropertiesResponsePtrOutput) Routes() RoutePropertiesResponseArrayOutput {
	return o.ApplyT(func(v *RoutingPropertiesResponse) []RoutePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Routes
	}).(RoutePropertiesResponseArrayOutput)
}

type RoutingServiceBusQueueEndpointProperties struct {
	ConnectionString string  `pulumi:"connectionString"`
	Name             string  `pulumi:"name"`
	ResourceGroup    *string `pulumi:"resourceGroup"`
	SubscriptionId   *string `pulumi:"subscriptionId"`
}





type RoutingServiceBusQueueEndpointPropertiesInput interface {
	pulumi.Input

	ToRoutingServiceBusQueueEndpointPropertiesOutput() RoutingServiceBusQueueEndpointPropertiesOutput
	ToRoutingServiceBusQueueEndpointPropertiesOutputWithContext(context.Context) RoutingServiceBusQueueEndpointPropertiesOutput
}

type RoutingServiceBusQueueEndpointPropertiesArgs struct {
	ConnectionString pulumi.StringInput    `pulumi:"connectionString"`
	Name             pulumi.StringInput    `pulumi:"name"`
	ResourceGroup    pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId   pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingServiceBusQueueEndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusQueueEndpointProperties)(nil)).Elem()
}

func (i RoutingServiceBusQueueEndpointPropertiesArgs) ToRoutingServiceBusQueueEndpointPropertiesOutput() RoutingServiceBusQueueEndpointPropertiesOutput {
	return i.ToRoutingServiceBusQueueEndpointPropertiesOutputWithContext(context.Background())
}

func (i RoutingServiceBusQueueEndpointPropertiesArgs) ToRoutingServiceBusQueueEndpointPropertiesOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusQueueEndpointPropertiesOutput)
}





type RoutingServiceBusQueueEndpointPropertiesArrayInput interface {
	pulumi.Input

	ToRoutingServiceBusQueueEndpointPropertiesArrayOutput() RoutingServiceBusQueueEndpointPropertiesArrayOutput
	ToRoutingServiceBusQueueEndpointPropertiesArrayOutputWithContext(context.Context) RoutingServiceBusQueueEndpointPropertiesArrayOutput
}

type RoutingServiceBusQueueEndpointPropertiesArray []RoutingServiceBusQueueEndpointPropertiesInput

func (RoutingServiceBusQueueEndpointPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusQueueEndpointProperties)(nil)).Elem()
}

func (i RoutingServiceBusQueueEndpointPropertiesArray) ToRoutingServiceBusQueueEndpointPropertiesArrayOutput() RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return i.ToRoutingServiceBusQueueEndpointPropertiesArrayOutputWithContext(context.Background())
}

func (i RoutingServiceBusQueueEndpointPropertiesArray) ToRoutingServiceBusQueueEndpointPropertiesArrayOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusQueueEndpointPropertiesArrayOutput)
}

type RoutingServiceBusQueueEndpointPropertiesOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusQueueEndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusQueueEndpointProperties)(nil)).Elem()
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) ToRoutingServiceBusQueueEndpointPropertiesOutput() RoutingServiceBusQueueEndpointPropertiesOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) ToRoutingServiceBusQueueEndpointPropertiesOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingServiceBusQueueEndpointPropertiesArrayOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusQueueEndpointPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusQueueEndpointProperties)(nil)).Elem()
}

func (o RoutingServiceBusQueueEndpointPropertiesArrayOutput) ToRoutingServiceBusQueueEndpointPropertiesArrayOutput() RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesArrayOutput) ToRoutingServiceBusQueueEndpointPropertiesArrayOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesArrayOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesArrayOutput) Index(i pulumi.IntInput) RoutingServiceBusQueueEndpointPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingServiceBusQueueEndpointProperties {
		return vs[0].([]RoutingServiceBusQueueEndpointProperties)[vs[1].(int)]
	}).(RoutingServiceBusQueueEndpointPropertiesOutput)
}

type RoutingServiceBusQueueEndpointPropertiesResponse struct {
	ConnectionString string  `pulumi:"connectionString"`
	Name             string  `pulumi:"name"`
	ResourceGroup    *string `pulumi:"resourceGroup"`
	SubscriptionId   *string `pulumi:"subscriptionId"`
}





type RoutingServiceBusQueueEndpointPropertiesResponseInput interface {
	pulumi.Input

	ToRoutingServiceBusQueueEndpointPropertiesResponseOutput() RoutingServiceBusQueueEndpointPropertiesResponseOutput
	ToRoutingServiceBusQueueEndpointPropertiesResponseOutputWithContext(context.Context) RoutingServiceBusQueueEndpointPropertiesResponseOutput
}

type RoutingServiceBusQueueEndpointPropertiesResponseArgs struct {
	ConnectionString pulumi.StringInput    `pulumi:"connectionString"`
	Name             pulumi.StringInput    `pulumi:"name"`
	ResourceGroup    pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId   pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingServiceBusQueueEndpointPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusQueueEndpointPropertiesResponse)(nil)).Elem()
}

func (i RoutingServiceBusQueueEndpointPropertiesResponseArgs) ToRoutingServiceBusQueueEndpointPropertiesResponseOutput() RoutingServiceBusQueueEndpointPropertiesResponseOutput {
	return i.ToRoutingServiceBusQueueEndpointPropertiesResponseOutputWithContext(context.Background())
}

func (i RoutingServiceBusQueueEndpointPropertiesResponseArgs) ToRoutingServiceBusQueueEndpointPropertiesResponseOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusQueueEndpointPropertiesResponseOutput)
}





type RoutingServiceBusQueueEndpointPropertiesResponseArrayInput interface {
	pulumi.Input

	ToRoutingServiceBusQueueEndpointPropertiesResponseArrayOutput() RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput
	ToRoutingServiceBusQueueEndpointPropertiesResponseArrayOutputWithContext(context.Context) RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput
}

type RoutingServiceBusQueueEndpointPropertiesResponseArray []RoutingServiceBusQueueEndpointPropertiesResponseInput

func (RoutingServiceBusQueueEndpointPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusQueueEndpointPropertiesResponse)(nil)).Elem()
}

func (i RoutingServiceBusQueueEndpointPropertiesResponseArray) ToRoutingServiceBusQueueEndpointPropertiesResponseArrayOutput() RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput {
	return i.ToRoutingServiceBusQueueEndpointPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i RoutingServiceBusQueueEndpointPropertiesResponseArray) ToRoutingServiceBusQueueEndpointPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput)
}

type RoutingServiceBusQueueEndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusQueueEndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusQueueEndpointPropertiesResponse)(nil)).Elem()
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) ToRoutingServiceBusQueueEndpointPropertiesResponseOutput() RoutingServiceBusQueueEndpointPropertiesResponseOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) ToRoutingServiceBusQueueEndpointPropertiesResponseOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesResponseOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusQueueEndpointPropertiesResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusQueueEndpointPropertiesResponse)(nil)).Elem()
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput) ToRoutingServiceBusQueueEndpointPropertiesResponseArrayOutput() RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput) ToRoutingServiceBusQueueEndpointPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput {
	return o
}

func (o RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput) Index(i pulumi.IntInput) RoutingServiceBusQueueEndpointPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingServiceBusQueueEndpointPropertiesResponse {
		return vs[0].([]RoutingServiceBusQueueEndpointPropertiesResponse)[vs[1].(int)]
	}).(RoutingServiceBusQueueEndpointPropertiesResponseOutput)
}

type RoutingServiceBusTopicEndpointProperties struct {
	ConnectionString string  `pulumi:"connectionString"`
	Name             string  `pulumi:"name"`
	ResourceGroup    *string `pulumi:"resourceGroup"`
	SubscriptionId   *string `pulumi:"subscriptionId"`
}





type RoutingServiceBusTopicEndpointPropertiesInput interface {
	pulumi.Input

	ToRoutingServiceBusTopicEndpointPropertiesOutput() RoutingServiceBusTopicEndpointPropertiesOutput
	ToRoutingServiceBusTopicEndpointPropertiesOutputWithContext(context.Context) RoutingServiceBusTopicEndpointPropertiesOutput
}

type RoutingServiceBusTopicEndpointPropertiesArgs struct {
	ConnectionString pulumi.StringInput    `pulumi:"connectionString"`
	Name             pulumi.StringInput    `pulumi:"name"`
	ResourceGroup    pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId   pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingServiceBusTopicEndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusTopicEndpointProperties)(nil)).Elem()
}

func (i RoutingServiceBusTopicEndpointPropertiesArgs) ToRoutingServiceBusTopicEndpointPropertiesOutput() RoutingServiceBusTopicEndpointPropertiesOutput {
	return i.ToRoutingServiceBusTopicEndpointPropertiesOutputWithContext(context.Background())
}

func (i RoutingServiceBusTopicEndpointPropertiesArgs) ToRoutingServiceBusTopicEndpointPropertiesOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusTopicEndpointPropertiesOutput)
}





type RoutingServiceBusTopicEndpointPropertiesArrayInput interface {
	pulumi.Input

	ToRoutingServiceBusTopicEndpointPropertiesArrayOutput() RoutingServiceBusTopicEndpointPropertiesArrayOutput
	ToRoutingServiceBusTopicEndpointPropertiesArrayOutputWithContext(context.Context) RoutingServiceBusTopicEndpointPropertiesArrayOutput
}

type RoutingServiceBusTopicEndpointPropertiesArray []RoutingServiceBusTopicEndpointPropertiesInput

func (RoutingServiceBusTopicEndpointPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusTopicEndpointProperties)(nil)).Elem()
}

func (i RoutingServiceBusTopicEndpointPropertiesArray) ToRoutingServiceBusTopicEndpointPropertiesArrayOutput() RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return i.ToRoutingServiceBusTopicEndpointPropertiesArrayOutputWithContext(context.Background())
}

func (i RoutingServiceBusTopicEndpointPropertiesArray) ToRoutingServiceBusTopicEndpointPropertiesArrayOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusTopicEndpointPropertiesArrayOutput)
}

type RoutingServiceBusTopicEndpointPropertiesOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusTopicEndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusTopicEndpointProperties)(nil)).Elem()
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) ToRoutingServiceBusTopicEndpointPropertiesOutput() RoutingServiceBusTopicEndpointPropertiesOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) ToRoutingServiceBusTopicEndpointPropertiesOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingServiceBusTopicEndpointPropertiesArrayOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusTopicEndpointPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusTopicEndpointProperties)(nil)).Elem()
}

func (o RoutingServiceBusTopicEndpointPropertiesArrayOutput) ToRoutingServiceBusTopicEndpointPropertiesArrayOutput() RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesArrayOutput) ToRoutingServiceBusTopicEndpointPropertiesArrayOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesArrayOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesArrayOutput) Index(i pulumi.IntInput) RoutingServiceBusTopicEndpointPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingServiceBusTopicEndpointProperties {
		return vs[0].([]RoutingServiceBusTopicEndpointProperties)[vs[1].(int)]
	}).(RoutingServiceBusTopicEndpointPropertiesOutput)
}

type RoutingServiceBusTopicEndpointPropertiesResponse struct {
	ConnectionString string  `pulumi:"connectionString"`
	Name             string  `pulumi:"name"`
	ResourceGroup    *string `pulumi:"resourceGroup"`
	SubscriptionId   *string `pulumi:"subscriptionId"`
}





type RoutingServiceBusTopicEndpointPropertiesResponseInput interface {
	pulumi.Input

	ToRoutingServiceBusTopicEndpointPropertiesResponseOutput() RoutingServiceBusTopicEndpointPropertiesResponseOutput
	ToRoutingServiceBusTopicEndpointPropertiesResponseOutputWithContext(context.Context) RoutingServiceBusTopicEndpointPropertiesResponseOutput
}

type RoutingServiceBusTopicEndpointPropertiesResponseArgs struct {
	ConnectionString pulumi.StringInput    `pulumi:"connectionString"`
	Name             pulumi.StringInput    `pulumi:"name"`
	ResourceGroup    pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId   pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingServiceBusTopicEndpointPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusTopicEndpointPropertiesResponse)(nil)).Elem()
}

func (i RoutingServiceBusTopicEndpointPropertiesResponseArgs) ToRoutingServiceBusTopicEndpointPropertiesResponseOutput() RoutingServiceBusTopicEndpointPropertiesResponseOutput {
	return i.ToRoutingServiceBusTopicEndpointPropertiesResponseOutputWithContext(context.Background())
}

func (i RoutingServiceBusTopicEndpointPropertiesResponseArgs) ToRoutingServiceBusTopicEndpointPropertiesResponseOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusTopicEndpointPropertiesResponseOutput)
}





type RoutingServiceBusTopicEndpointPropertiesResponseArrayInput interface {
	pulumi.Input

	ToRoutingServiceBusTopicEndpointPropertiesResponseArrayOutput() RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput
	ToRoutingServiceBusTopicEndpointPropertiesResponseArrayOutputWithContext(context.Context) RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput
}

type RoutingServiceBusTopicEndpointPropertiesResponseArray []RoutingServiceBusTopicEndpointPropertiesResponseInput

func (RoutingServiceBusTopicEndpointPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusTopicEndpointPropertiesResponse)(nil)).Elem()
}

func (i RoutingServiceBusTopicEndpointPropertiesResponseArray) ToRoutingServiceBusTopicEndpointPropertiesResponseArrayOutput() RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput {
	return i.ToRoutingServiceBusTopicEndpointPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i RoutingServiceBusTopicEndpointPropertiesResponseArray) ToRoutingServiceBusTopicEndpointPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput)
}

type RoutingServiceBusTopicEndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusTopicEndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingServiceBusTopicEndpointPropertiesResponse)(nil)).Elem()
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) ToRoutingServiceBusTopicEndpointPropertiesResponseOutput() RoutingServiceBusTopicEndpointPropertiesResponseOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) ToRoutingServiceBusTopicEndpointPropertiesResponseOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesResponseOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingServiceBusTopicEndpointPropertiesResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingServiceBusTopicEndpointPropertiesResponse)(nil)).Elem()
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput) ToRoutingServiceBusTopicEndpointPropertiesResponseArrayOutput() RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput) ToRoutingServiceBusTopicEndpointPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput {
	return o
}

func (o RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput) Index(i pulumi.IntInput) RoutingServiceBusTopicEndpointPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingServiceBusTopicEndpointPropertiesResponse {
		return vs[0].([]RoutingServiceBusTopicEndpointPropertiesResponse)[vs[1].(int)]
	}).(RoutingServiceBusTopicEndpointPropertiesResponseOutput)
}

type RoutingStorageContainerProperties struct {
	BatchFrequencyInSeconds *int    `pulumi:"batchFrequencyInSeconds"`
	ConnectionString        string  `pulumi:"connectionString"`
	ContainerName           string  `pulumi:"containerName"`
	Encoding                *string `pulumi:"encoding"`
	FileNameFormat          *string `pulumi:"fileNameFormat"`
	MaxChunkSizeInBytes     *int    `pulumi:"maxChunkSizeInBytes"`
	Name                    string  `pulumi:"name"`
	ResourceGroup           *string `pulumi:"resourceGroup"`
	SubscriptionId          *string `pulumi:"subscriptionId"`
}





type RoutingStorageContainerPropertiesInput interface {
	pulumi.Input

	ToRoutingStorageContainerPropertiesOutput() RoutingStorageContainerPropertiesOutput
	ToRoutingStorageContainerPropertiesOutputWithContext(context.Context) RoutingStorageContainerPropertiesOutput
}

type RoutingStorageContainerPropertiesArgs struct {
	BatchFrequencyInSeconds pulumi.IntPtrInput    `pulumi:"batchFrequencyInSeconds"`
	ConnectionString        pulumi.StringInput    `pulumi:"connectionString"`
	ContainerName           pulumi.StringInput    `pulumi:"containerName"`
	Encoding                pulumi.StringPtrInput `pulumi:"encoding"`
	FileNameFormat          pulumi.StringPtrInput `pulumi:"fileNameFormat"`
	MaxChunkSizeInBytes     pulumi.IntPtrInput    `pulumi:"maxChunkSizeInBytes"`
	Name                    pulumi.StringInput    `pulumi:"name"`
	ResourceGroup           pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId          pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingStorageContainerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingStorageContainerProperties)(nil)).Elem()
}

func (i RoutingStorageContainerPropertiesArgs) ToRoutingStorageContainerPropertiesOutput() RoutingStorageContainerPropertiesOutput {
	return i.ToRoutingStorageContainerPropertiesOutputWithContext(context.Background())
}

func (i RoutingStorageContainerPropertiesArgs) ToRoutingStorageContainerPropertiesOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingStorageContainerPropertiesOutput)
}





type RoutingStorageContainerPropertiesArrayInput interface {
	pulumi.Input

	ToRoutingStorageContainerPropertiesArrayOutput() RoutingStorageContainerPropertiesArrayOutput
	ToRoutingStorageContainerPropertiesArrayOutputWithContext(context.Context) RoutingStorageContainerPropertiesArrayOutput
}

type RoutingStorageContainerPropertiesArray []RoutingStorageContainerPropertiesInput

func (RoutingStorageContainerPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingStorageContainerProperties)(nil)).Elem()
}

func (i RoutingStorageContainerPropertiesArray) ToRoutingStorageContainerPropertiesArrayOutput() RoutingStorageContainerPropertiesArrayOutput {
	return i.ToRoutingStorageContainerPropertiesArrayOutputWithContext(context.Background())
}

func (i RoutingStorageContainerPropertiesArray) ToRoutingStorageContainerPropertiesArrayOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingStorageContainerPropertiesArrayOutput)
}

type RoutingStorageContainerPropertiesOutput struct{ *pulumi.OutputState }

func (RoutingStorageContainerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingStorageContainerProperties)(nil)).Elem()
}

func (o RoutingStorageContainerPropertiesOutput) ToRoutingStorageContainerPropertiesOutput() RoutingStorageContainerPropertiesOutput {
	return o
}

func (o RoutingStorageContainerPropertiesOutput) ToRoutingStorageContainerPropertiesOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesOutput {
	return o
}

func (o RoutingStorageContainerPropertiesOutput) BatchFrequencyInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *int { return v.BatchFrequencyInSeconds }).(pulumi.IntPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o RoutingStorageContainerPropertiesOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o RoutingStorageContainerPropertiesOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) FileNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.FileNameFormat }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) MaxChunkSizeInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *int { return v.MaxChunkSizeInBytes }).(pulumi.IntPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingStorageContainerPropertiesOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerProperties) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingStorageContainerPropertiesArrayOutput struct{ *pulumi.OutputState }

func (RoutingStorageContainerPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingStorageContainerProperties)(nil)).Elem()
}

func (o RoutingStorageContainerPropertiesArrayOutput) ToRoutingStorageContainerPropertiesArrayOutput() RoutingStorageContainerPropertiesArrayOutput {
	return o
}

func (o RoutingStorageContainerPropertiesArrayOutput) ToRoutingStorageContainerPropertiesArrayOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesArrayOutput {
	return o
}

func (o RoutingStorageContainerPropertiesArrayOutput) Index(i pulumi.IntInput) RoutingStorageContainerPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingStorageContainerProperties {
		return vs[0].([]RoutingStorageContainerProperties)[vs[1].(int)]
	}).(RoutingStorageContainerPropertiesOutput)
}

type RoutingStorageContainerPropertiesResponse struct {
	BatchFrequencyInSeconds *int    `pulumi:"batchFrequencyInSeconds"`
	ConnectionString        string  `pulumi:"connectionString"`
	ContainerName           string  `pulumi:"containerName"`
	Encoding                *string `pulumi:"encoding"`
	FileNameFormat          *string `pulumi:"fileNameFormat"`
	MaxChunkSizeInBytes     *int    `pulumi:"maxChunkSizeInBytes"`
	Name                    string  `pulumi:"name"`
	ResourceGroup           *string `pulumi:"resourceGroup"`
	SubscriptionId          *string `pulumi:"subscriptionId"`
}





type RoutingStorageContainerPropertiesResponseInput interface {
	pulumi.Input

	ToRoutingStorageContainerPropertiesResponseOutput() RoutingStorageContainerPropertiesResponseOutput
	ToRoutingStorageContainerPropertiesResponseOutputWithContext(context.Context) RoutingStorageContainerPropertiesResponseOutput
}

type RoutingStorageContainerPropertiesResponseArgs struct {
	BatchFrequencyInSeconds pulumi.IntPtrInput    `pulumi:"batchFrequencyInSeconds"`
	ConnectionString        pulumi.StringInput    `pulumi:"connectionString"`
	ContainerName           pulumi.StringInput    `pulumi:"containerName"`
	Encoding                pulumi.StringPtrInput `pulumi:"encoding"`
	FileNameFormat          pulumi.StringPtrInput `pulumi:"fileNameFormat"`
	MaxChunkSizeInBytes     pulumi.IntPtrInput    `pulumi:"maxChunkSizeInBytes"`
	Name                    pulumi.StringInput    `pulumi:"name"`
	ResourceGroup           pulumi.StringPtrInput `pulumi:"resourceGroup"`
	SubscriptionId          pulumi.StringPtrInput `pulumi:"subscriptionId"`
}

func (RoutingStorageContainerPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingStorageContainerPropertiesResponse)(nil)).Elem()
}

func (i RoutingStorageContainerPropertiesResponseArgs) ToRoutingStorageContainerPropertiesResponseOutput() RoutingStorageContainerPropertiesResponseOutput {
	return i.ToRoutingStorageContainerPropertiesResponseOutputWithContext(context.Background())
}

func (i RoutingStorageContainerPropertiesResponseArgs) ToRoutingStorageContainerPropertiesResponseOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingStorageContainerPropertiesResponseOutput)
}





type RoutingStorageContainerPropertiesResponseArrayInput interface {
	pulumi.Input

	ToRoutingStorageContainerPropertiesResponseArrayOutput() RoutingStorageContainerPropertiesResponseArrayOutput
	ToRoutingStorageContainerPropertiesResponseArrayOutputWithContext(context.Context) RoutingStorageContainerPropertiesResponseArrayOutput
}

type RoutingStorageContainerPropertiesResponseArray []RoutingStorageContainerPropertiesResponseInput

func (RoutingStorageContainerPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingStorageContainerPropertiesResponse)(nil)).Elem()
}

func (i RoutingStorageContainerPropertiesResponseArray) ToRoutingStorageContainerPropertiesResponseArrayOutput() RoutingStorageContainerPropertiesResponseArrayOutput {
	return i.ToRoutingStorageContainerPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i RoutingStorageContainerPropertiesResponseArray) ToRoutingStorageContainerPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutingStorageContainerPropertiesResponseArrayOutput)
}

type RoutingStorageContainerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RoutingStorageContainerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RoutingStorageContainerPropertiesResponse)(nil)).Elem()
}

func (o RoutingStorageContainerPropertiesResponseOutput) ToRoutingStorageContainerPropertiesResponseOutput() RoutingStorageContainerPropertiesResponseOutput {
	return o
}

func (o RoutingStorageContainerPropertiesResponseOutput) ToRoutingStorageContainerPropertiesResponseOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesResponseOutput {
	return o
}

func (o RoutingStorageContainerPropertiesResponseOutput) BatchFrequencyInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *int { return v.BatchFrequencyInSeconds }).(pulumi.IntPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) Encoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.Encoding }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) FileNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.FileNameFormat }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) MaxChunkSizeInBytes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *int { return v.MaxChunkSizeInBytes }).(pulumi.IntPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) ResourceGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.ResourceGroup }).(pulumi.StringPtrOutput)
}

func (o RoutingStorageContainerPropertiesResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RoutingStorageContainerPropertiesResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

type RoutingStorageContainerPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (RoutingStorageContainerPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RoutingStorageContainerPropertiesResponse)(nil)).Elem()
}

func (o RoutingStorageContainerPropertiesResponseArrayOutput) ToRoutingStorageContainerPropertiesResponseArrayOutput() RoutingStorageContainerPropertiesResponseArrayOutput {
	return o
}

func (o RoutingStorageContainerPropertiesResponseArrayOutput) ToRoutingStorageContainerPropertiesResponseArrayOutputWithContext(ctx context.Context) RoutingStorageContainerPropertiesResponseArrayOutput {
	return o
}

func (o RoutingStorageContainerPropertiesResponseArrayOutput) Index(i pulumi.IntInput) RoutingStorageContainerPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RoutingStorageContainerPropertiesResponse {
		return vs[0].([]RoutingStorageContainerPropertiesResponse)[vs[1].(int)]
	}).(RoutingStorageContainerPropertiesResponseOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*CertificatePropertiesResponseInput)(nil)).Elem(), CertificatePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertificatePropertiesResponsePtrInput)(nil)).Elem(), CertificatePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudToDevicePropertiesInput)(nil)).Elem(), CloudToDevicePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudToDevicePropertiesPtrInput)(nil)).Elem(), CloudToDevicePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudToDevicePropertiesResponseInput)(nil)).Elem(), CloudToDevicePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudToDevicePropertiesResponsePtrInput)(nil)).Elem(), CloudToDevicePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubPropertiesInput)(nil)).Elem(), EventHubPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubPropertiesMapInput)(nil)).Elem(), EventHubPropertiesMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubPropertiesResponseInput)(nil)).Elem(), EventHubPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubPropertiesResponseMapInput)(nil)).Elem(), EventHubPropertiesResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*FallbackRoutePropertiesInput)(nil)).Elem(), FallbackRoutePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FallbackRoutePropertiesPtrInput)(nil)).Elem(), FallbackRoutePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FallbackRoutePropertiesResponseInput)(nil)).Elem(), FallbackRoutePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FallbackRoutePropertiesResponsePtrInput)(nil)).Elem(), FallbackRoutePropertiesResponseArgs{})
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
	pulumi.RegisterInputType(reflect.TypeOf((*RoutePropertiesInput)(nil)).Elem(), RoutePropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutePropertiesArrayInput)(nil)).Elem(), RoutePropertiesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutePropertiesResponseInput)(nil)).Elem(), RoutePropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutePropertiesResponseArrayInput)(nil)).Elem(), RoutePropertiesResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingEndpointsInput)(nil)).Elem(), RoutingEndpointsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingEndpointsPtrInput)(nil)).Elem(), RoutingEndpointsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingEndpointsResponseInput)(nil)).Elem(), RoutingEndpointsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingEndpointsResponsePtrInput)(nil)).Elem(), RoutingEndpointsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingEventHubPropertiesInput)(nil)).Elem(), RoutingEventHubPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingEventHubPropertiesArrayInput)(nil)).Elem(), RoutingEventHubPropertiesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingEventHubPropertiesResponseInput)(nil)).Elem(), RoutingEventHubPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingEventHubPropertiesResponseArrayInput)(nil)).Elem(), RoutingEventHubPropertiesResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingPropertiesInput)(nil)).Elem(), RoutingPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingPropertiesPtrInput)(nil)).Elem(), RoutingPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingPropertiesResponseInput)(nil)).Elem(), RoutingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingPropertiesResponsePtrInput)(nil)).Elem(), RoutingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingServiceBusQueueEndpointPropertiesInput)(nil)).Elem(), RoutingServiceBusQueueEndpointPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingServiceBusQueueEndpointPropertiesArrayInput)(nil)).Elem(), RoutingServiceBusQueueEndpointPropertiesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingServiceBusQueueEndpointPropertiesResponseInput)(nil)).Elem(), RoutingServiceBusQueueEndpointPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingServiceBusQueueEndpointPropertiesResponseArrayInput)(nil)).Elem(), RoutingServiceBusQueueEndpointPropertiesResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingServiceBusTopicEndpointPropertiesInput)(nil)).Elem(), RoutingServiceBusTopicEndpointPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingServiceBusTopicEndpointPropertiesArrayInput)(nil)).Elem(), RoutingServiceBusTopicEndpointPropertiesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingServiceBusTopicEndpointPropertiesResponseInput)(nil)).Elem(), RoutingServiceBusTopicEndpointPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingServiceBusTopicEndpointPropertiesResponseArrayInput)(nil)).Elem(), RoutingServiceBusTopicEndpointPropertiesResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingStorageContainerPropertiesInput)(nil)).Elem(), RoutingStorageContainerPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingStorageContainerPropertiesArrayInput)(nil)).Elem(), RoutingStorageContainerPropertiesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingStorageContainerPropertiesResponseInput)(nil)).Elem(), RoutingStorageContainerPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoutingStorageContainerPropertiesResponseArrayInput)(nil)).Elem(), RoutingStorageContainerPropertiesResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleInput)(nil)).Elem(), SharedAccessSignatureAuthorizationRuleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleArrayInput)(nil)).Elem(), SharedAccessSignatureAuthorizationRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleResponseInput)(nil)).Elem(), SharedAccessSignatureAuthorizationRuleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SharedAccessSignatureAuthorizationRuleResponseArrayInput)(nil)).Elem(), SharedAccessSignatureAuthorizationRuleResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageEndpointPropertiesInput)(nil)).Elem(), StorageEndpointPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageEndpointPropertiesMapInput)(nil)).Elem(), StorageEndpointPropertiesMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageEndpointPropertiesResponseInput)(nil)).Elem(), StorageEndpointPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*StorageEndpointPropertiesResponseMapInput)(nil)).Elem(), StorageEndpointPropertiesResponseMap{})
	pulumi.RegisterOutputType(CertificatePropertiesResponseOutput{})
	pulumi.RegisterOutputType(CertificatePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesResponseOutput{})
	pulumi.RegisterOutputType(CloudToDevicePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesMapOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EventHubPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(FallbackRoutePropertiesOutput{})
	pulumi.RegisterOutputType(FallbackRoutePropertiesPtrOutput{})
	pulumi.RegisterOutputType(FallbackRoutePropertiesResponseOutput{})
	pulumi.RegisterOutputType(FallbackRoutePropertiesResponsePtrOutput{})
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
	pulumi.RegisterOutputType(RoutePropertiesOutput{})
	pulumi.RegisterOutputType(RoutePropertiesArrayOutput{})
	pulumi.RegisterOutputType(RoutePropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutePropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingEndpointsOutput{})
	pulumi.RegisterOutputType(RoutingEndpointsPtrOutput{})
	pulumi.RegisterOutputType(RoutingEndpointsResponseOutput{})
	pulumi.RegisterOutputType(RoutingEndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(RoutingEventHubPropertiesOutput{})
	pulumi.RegisterOutputType(RoutingEventHubPropertiesArrayOutput{})
	pulumi.RegisterOutputType(RoutingEventHubPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutingEventHubPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingPropertiesOutput{})
	pulumi.RegisterOutputType(RoutingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(RoutingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusQueueEndpointPropertiesOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusQueueEndpointPropertiesArrayOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusQueueEndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusQueueEndpointPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusTopicEndpointPropertiesOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusTopicEndpointPropertiesArrayOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusTopicEndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutingServiceBusTopicEndpointPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(RoutingStorageContainerPropertiesOutput{})
	pulumi.RegisterOutputType(RoutingStorageContainerPropertiesArrayOutput{})
	pulumi.RegisterOutputType(RoutingStorageContainerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RoutingStorageContainerPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleArrayOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleResponseOutput{})
	pulumi.RegisterOutputType(SharedAccessSignatureAuthorizationRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesMapOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageEndpointPropertiesResponseMapOutput{})
}
