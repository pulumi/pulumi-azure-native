


package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AsymmetricEncryptedSecret struct {
	EncryptionAlgorithm      EncryptionAlgorithm `pulumi:"encryptionAlgorithm"`
	EncryptionCertThumbprint *string             `pulumi:"encryptionCertThumbprint"`
	Value                    string              `pulumi:"value"`
}





type AsymmetricEncryptedSecretInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretOutput() AsymmetricEncryptedSecretOutput
	ToAsymmetricEncryptedSecretOutputWithContext(context.Context) AsymmetricEncryptedSecretOutput
}

type AsymmetricEncryptedSecretArgs struct {
	EncryptionAlgorithm      EncryptionAlgorithmInput `pulumi:"encryptionAlgorithm"`
	EncryptionCertThumbprint pulumi.StringPtrInput    `pulumi:"encryptionCertThumbprint"`
	Value                    pulumi.StringInput       `pulumi:"value"`
}

func (AsymmetricEncryptedSecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecret)(nil)).Elem()
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretOutput() AsymmetricEncryptedSecretOutput {
	return i.ToAsymmetricEncryptedSecretOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretOutput)
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return i.ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretArgs) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretOutput).ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx)
}









type AsymmetricEncryptedSecretPtrInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput
	ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Context) AsymmetricEncryptedSecretPtrOutput
}

type asymmetricEncryptedSecretPtrType AsymmetricEncryptedSecretArgs

func AsymmetricEncryptedSecretPtr(v *AsymmetricEncryptedSecretArgs) AsymmetricEncryptedSecretPtrInput {
	return (*asymmetricEncryptedSecretPtrType)(v)
}

func (*asymmetricEncryptedSecretPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecret)(nil)).Elem()
}

func (i *asymmetricEncryptedSecretPtrType) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return i.ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Background())
}

func (i *asymmetricEncryptedSecretPtrType) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretPtrOutput)
}

type AsymmetricEncryptedSecretOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecret)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretOutput() AsymmetricEncryptedSecretOutput {
	return o
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretOutput {
	return o
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return o.ToAsymmetricEncryptedSecretPtrOutputWithContext(context.Background())
}

func (o AsymmetricEncryptedSecretOutput) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AsymmetricEncryptedSecret) *AsymmetricEncryptedSecret {
		return &v
	}).(AsymmetricEncryptedSecretPtrOutput)
}

func (o AsymmetricEncryptedSecretOutput) EncryptionAlgorithm() EncryptionAlgorithmOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecret) EncryptionAlgorithm { return v.EncryptionAlgorithm }).(EncryptionAlgorithmOutput)
}

func (o AsymmetricEncryptedSecretOutput) EncryptionCertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecret) *string { return v.EncryptionCertThumbprint }).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecret) string { return v.Value }).(pulumi.StringOutput)
}

type AsymmetricEncryptedSecretPtrOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecret)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretPtrOutput) ToAsymmetricEncryptedSecretPtrOutput() AsymmetricEncryptedSecretPtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretPtrOutput) ToAsymmetricEncryptedSecretPtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretPtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretPtrOutput) Elem() AsymmetricEncryptedSecretOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) AsymmetricEncryptedSecret {
		if v != nil {
			return *v
		}
		var ret AsymmetricEncryptedSecret
		return ret
	}).(AsymmetricEncryptedSecretOutput)
}

func (o AsymmetricEncryptedSecretPtrOutput) EncryptionAlgorithm() EncryptionAlgorithmPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) *EncryptionAlgorithm {
		if v == nil {
			return nil
		}
		return &v.EncryptionAlgorithm
	}).(EncryptionAlgorithmPtrOutput)
}

func (o AsymmetricEncryptedSecretPtrOutput) EncryptionCertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionCertThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecret) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type AsymmetricEncryptedSecretResponse struct {
	EncryptionAlgorithm      string  `pulumi:"encryptionAlgorithm"`
	EncryptionCertThumbprint *string `pulumi:"encryptionCertThumbprint"`
	Value                    string  `pulumi:"value"`
}





type AsymmetricEncryptedSecretResponseInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretResponseOutput() AsymmetricEncryptedSecretResponseOutput
	ToAsymmetricEncryptedSecretResponseOutputWithContext(context.Context) AsymmetricEncryptedSecretResponseOutput
}

type AsymmetricEncryptedSecretResponseArgs struct {
	EncryptionAlgorithm      pulumi.StringInput    `pulumi:"encryptionAlgorithm"`
	EncryptionCertThumbprint pulumi.StringPtrInput `pulumi:"encryptionCertThumbprint"`
	Value                    pulumi.StringInput    `pulumi:"value"`
}

func (AsymmetricEncryptedSecretResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponseOutput() AsymmetricEncryptedSecretResponseOutput {
	return i.ToAsymmetricEncryptedSecretResponseOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponseOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretResponseOutput)
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return i.ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Background())
}

func (i AsymmetricEncryptedSecretResponseArgs) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretResponseOutput).ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx)
}









type AsymmetricEncryptedSecretResponsePtrInput interface {
	pulumi.Input

	ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput
	ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Context) AsymmetricEncryptedSecretResponsePtrOutput
}

type asymmetricEncryptedSecretResponsePtrType AsymmetricEncryptedSecretResponseArgs

func AsymmetricEncryptedSecretResponsePtr(v *AsymmetricEncryptedSecretResponseArgs) AsymmetricEncryptedSecretResponsePtrInput {
	return (*asymmetricEncryptedSecretResponsePtrType)(v)
}

func (*asymmetricEncryptedSecretResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (i *asymmetricEncryptedSecretResponsePtrType) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return i.ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Background())
}

func (i *asymmetricEncryptedSecretResponsePtrType) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AsymmetricEncryptedSecretResponsePtrOutput)
}

type AsymmetricEncryptedSecretResponseOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponseOutput() AsymmetricEncryptedSecretResponseOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponseOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponseOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(context.Background())
}

func (o AsymmetricEncryptedSecretResponseOutput) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AsymmetricEncryptedSecretResponse) *AsymmetricEncryptedSecretResponse {
		return &v
	}).(AsymmetricEncryptedSecretResponsePtrOutput)
}

func (o AsymmetricEncryptedSecretResponseOutput) EncryptionAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecretResponse) string { return v.EncryptionAlgorithm }).(pulumi.StringOutput)
}

func (o AsymmetricEncryptedSecretResponseOutput) EncryptionCertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecretResponse) *string { return v.EncryptionCertThumbprint }).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v AsymmetricEncryptedSecretResponse) string { return v.Value }).(pulumi.StringOutput)
}

type AsymmetricEncryptedSecretResponsePtrOutput struct{ *pulumi.OutputState }

func (AsymmetricEncryptedSecretResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AsymmetricEncryptedSecretResponse)(nil)).Elem()
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) ToAsymmetricEncryptedSecretResponsePtrOutput() AsymmetricEncryptedSecretResponsePtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) ToAsymmetricEncryptedSecretResponsePtrOutputWithContext(ctx context.Context) AsymmetricEncryptedSecretResponsePtrOutput {
	return o
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) Elem() AsymmetricEncryptedSecretResponseOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) AsymmetricEncryptedSecretResponse {
		if v != nil {
			return *v
		}
		var ret AsymmetricEncryptedSecretResponse
		return ret
	}).(AsymmetricEncryptedSecretResponseOutput)
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) EncryptionAlgorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EncryptionAlgorithm
	}).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) EncryptionCertThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) *string {
		if v == nil {
			return nil
		}
		return v.EncryptionCertThumbprint
	}).(pulumi.StringPtrOutput)
}

func (o AsymmetricEncryptedSecretResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AsymmetricEncryptedSecretResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type BandwidthSchedule struct {
	Days       []DayOfWeek `pulumi:"days"`
	RateInMbps int         `pulumi:"rateInMbps"`
	Start      Time        `pulumi:"start"`
	Stop       Time        `pulumi:"stop"`
}





type BandwidthScheduleInput interface {
	pulumi.Input

	ToBandwidthScheduleOutput() BandwidthScheduleOutput
	ToBandwidthScheduleOutputWithContext(context.Context) BandwidthScheduleOutput
}

type BandwidthScheduleArgs struct {
	Days       DayOfWeekArrayInput `pulumi:"days"`
	RateInMbps pulumi.IntInput     `pulumi:"rateInMbps"`
	Start      TimeInput           `pulumi:"start"`
	Stop       TimeInput           `pulumi:"stop"`
}

func (BandwidthScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthSchedule)(nil)).Elem()
}

func (i BandwidthScheduleArgs) ToBandwidthScheduleOutput() BandwidthScheduleOutput {
	return i.ToBandwidthScheduleOutputWithContext(context.Background())
}

func (i BandwidthScheduleArgs) ToBandwidthScheduleOutputWithContext(ctx context.Context) BandwidthScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthScheduleOutput)
}





type BandwidthScheduleArrayInput interface {
	pulumi.Input

	ToBandwidthScheduleArrayOutput() BandwidthScheduleArrayOutput
	ToBandwidthScheduleArrayOutputWithContext(context.Context) BandwidthScheduleArrayOutput
}

type BandwidthScheduleArray []BandwidthScheduleInput

func (BandwidthScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BandwidthSchedule)(nil)).Elem()
}

func (i BandwidthScheduleArray) ToBandwidthScheduleArrayOutput() BandwidthScheduleArrayOutput {
	return i.ToBandwidthScheduleArrayOutputWithContext(context.Background())
}

func (i BandwidthScheduleArray) ToBandwidthScheduleArrayOutputWithContext(ctx context.Context) BandwidthScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthScheduleArrayOutput)
}

type BandwidthScheduleOutput struct{ *pulumi.OutputState }

func (BandwidthScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthSchedule)(nil)).Elem()
}

func (o BandwidthScheduleOutput) ToBandwidthScheduleOutput() BandwidthScheduleOutput {
	return o
}

func (o BandwidthScheduleOutput) ToBandwidthScheduleOutputWithContext(ctx context.Context) BandwidthScheduleOutput {
	return o
}

func (o BandwidthScheduleOutput) Days() DayOfWeekArrayOutput {
	return o.ApplyT(func(v BandwidthSchedule) []DayOfWeek { return v.Days }).(DayOfWeekArrayOutput)
}

func (o BandwidthScheduleOutput) RateInMbps() pulumi.IntOutput {
	return o.ApplyT(func(v BandwidthSchedule) int { return v.RateInMbps }).(pulumi.IntOutput)
}

func (o BandwidthScheduleOutput) Start() TimeOutput {
	return o.ApplyT(func(v BandwidthSchedule) Time { return v.Start }).(TimeOutput)
}

func (o BandwidthScheduleOutput) Stop() TimeOutput {
	return o.ApplyT(func(v BandwidthSchedule) Time { return v.Stop }).(TimeOutput)
}

type BandwidthScheduleArrayOutput struct{ *pulumi.OutputState }

func (BandwidthScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BandwidthSchedule)(nil)).Elem()
}

func (o BandwidthScheduleArrayOutput) ToBandwidthScheduleArrayOutput() BandwidthScheduleArrayOutput {
	return o
}

func (o BandwidthScheduleArrayOutput) ToBandwidthScheduleArrayOutputWithContext(ctx context.Context) BandwidthScheduleArrayOutput {
	return o
}

func (o BandwidthScheduleArrayOutput) Index(i pulumi.IntInput) BandwidthScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BandwidthSchedule {
		return vs[0].([]BandwidthSchedule)[vs[1].(int)]
	}).(BandwidthScheduleOutput)
}

type BandwidthScheduleResponse struct {
	Days       []string     `pulumi:"days"`
	RateInMbps int          `pulumi:"rateInMbps"`
	Start      TimeResponse `pulumi:"start"`
	Stop       TimeResponse `pulumi:"stop"`
}





type BandwidthScheduleResponseInput interface {
	pulumi.Input

	ToBandwidthScheduleResponseOutput() BandwidthScheduleResponseOutput
	ToBandwidthScheduleResponseOutputWithContext(context.Context) BandwidthScheduleResponseOutput
}

type BandwidthScheduleResponseArgs struct {
	Days       pulumi.StringArrayInput `pulumi:"days"`
	RateInMbps pulumi.IntInput         `pulumi:"rateInMbps"`
	Start      TimeResponseInput       `pulumi:"start"`
	Stop       TimeResponseInput       `pulumi:"stop"`
}

func (BandwidthScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthScheduleResponse)(nil)).Elem()
}

func (i BandwidthScheduleResponseArgs) ToBandwidthScheduleResponseOutput() BandwidthScheduleResponseOutput {
	return i.ToBandwidthScheduleResponseOutputWithContext(context.Background())
}

func (i BandwidthScheduleResponseArgs) ToBandwidthScheduleResponseOutputWithContext(ctx context.Context) BandwidthScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthScheduleResponseOutput)
}





type BandwidthScheduleResponseArrayInput interface {
	pulumi.Input

	ToBandwidthScheduleResponseArrayOutput() BandwidthScheduleResponseArrayOutput
	ToBandwidthScheduleResponseArrayOutputWithContext(context.Context) BandwidthScheduleResponseArrayOutput
}

type BandwidthScheduleResponseArray []BandwidthScheduleResponseInput

func (BandwidthScheduleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BandwidthScheduleResponse)(nil)).Elem()
}

func (i BandwidthScheduleResponseArray) ToBandwidthScheduleResponseArrayOutput() BandwidthScheduleResponseArrayOutput {
	return i.ToBandwidthScheduleResponseArrayOutputWithContext(context.Background())
}

func (i BandwidthScheduleResponseArray) ToBandwidthScheduleResponseArrayOutputWithContext(ctx context.Context) BandwidthScheduleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthScheduleResponseArrayOutput)
}

type BandwidthScheduleResponseOutput struct{ *pulumi.OutputState }

func (BandwidthScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthScheduleResponse)(nil)).Elem()
}

func (o BandwidthScheduleResponseOutput) ToBandwidthScheduleResponseOutput() BandwidthScheduleResponseOutput {
	return o
}

func (o BandwidthScheduleResponseOutput) ToBandwidthScheduleResponseOutputWithContext(ctx context.Context) BandwidthScheduleResponseOutput {
	return o
}

func (o BandwidthScheduleResponseOutput) Days() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BandwidthScheduleResponse) []string { return v.Days }).(pulumi.StringArrayOutput)
}

func (o BandwidthScheduleResponseOutput) RateInMbps() pulumi.IntOutput {
	return o.ApplyT(func(v BandwidthScheduleResponse) int { return v.RateInMbps }).(pulumi.IntOutput)
}

func (o BandwidthScheduleResponseOutput) Start() TimeResponseOutput {
	return o.ApplyT(func(v BandwidthScheduleResponse) TimeResponse { return v.Start }).(TimeResponseOutput)
}

func (o BandwidthScheduleResponseOutput) Stop() TimeResponseOutput {
	return o.ApplyT(func(v BandwidthScheduleResponse) TimeResponse { return v.Stop }).(TimeResponseOutput)
}

type BandwidthScheduleResponseArrayOutput struct{ *pulumi.OutputState }

func (BandwidthScheduleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BandwidthScheduleResponse)(nil)).Elem()
}

func (o BandwidthScheduleResponseArrayOutput) ToBandwidthScheduleResponseArrayOutput() BandwidthScheduleResponseArrayOutput {
	return o
}

func (o BandwidthScheduleResponseArrayOutput) ToBandwidthScheduleResponseArrayOutputWithContext(ctx context.Context) BandwidthScheduleResponseArrayOutput {
	return o
}

func (o BandwidthScheduleResponseArrayOutput) Index(i pulumi.IntInput) BandwidthScheduleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BandwidthScheduleResponse {
		return vs[0].([]BandwidthScheduleResponse)[vs[1].(int)]
	}).(BandwidthScheduleResponseOutput)
}

type FailoverSetEligibilityResultResponse struct {
	ErrorMessage          *string `pulumi:"errorMessage"`
	IsEligibleForFailover *bool   `pulumi:"isEligibleForFailover"`
}





type FailoverSetEligibilityResultResponseInput interface {
	pulumi.Input

	ToFailoverSetEligibilityResultResponseOutput() FailoverSetEligibilityResultResponseOutput
	ToFailoverSetEligibilityResultResponseOutputWithContext(context.Context) FailoverSetEligibilityResultResponseOutput
}

type FailoverSetEligibilityResultResponseArgs struct {
	ErrorMessage          pulumi.StringPtrInput `pulumi:"errorMessage"`
	IsEligibleForFailover pulumi.BoolPtrInput   `pulumi:"isEligibleForFailover"`
}

func (FailoverSetEligibilityResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverSetEligibilityResultResponse)(nil)).Elem()
}

func (i FailoverSetEligibilityResultResponseArgs) ToFailoverSetEligibilityResultResponseOutput() FailoverSetEligibilityResultResponseOutput {
	return i.ToFailoverSetEligibilityResultResponseOutputWithContext(context.Background())
}

func (i FailoverSetEligibilityResultResponseArgs) ToFailoverSetEligibilityResultResponseOutputWithContext(ctx context.Context) FailoverSetEligibilityResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverSetEligibilityResultResponseOutput)
}

func (i FailoverSetEligibilityResultResponseArgs) ToFailoverSetEligibilityResultResponsePtrOutput() FailoverSetEligibilityResultResponsePtrOutput {
	return i.ToFailoverSetEligibilityResultResponsePtrOutputWithContext(context.Background())
}

func (i FailoverSetEligibilityResultResponseArgs) ToFailoverSetEligibilityResultResponsePtrOutputWithContext(ctx context.Context) FailoverSetEligibilityResultResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverSetEligibilityResultResponseOutput).ToFailoverSetEligibilityResultResponsePtrOutputWithContext(ctx)
}









type FailoverSetEligibilityResultResponsePtrInput interface {
	pulumi.Input

	ToFailoverSetEligibilityResultResponsePtrOutput() FailoverSetEligibilityResultResponsePtrOutput
	ToFailoverSetEligibilityResultResponsePtrOutputWithContext(context.Context) FailoverSetEligibilityResultResponsePtrOutput
}

type failoverSetEligibilityResultResponsePtrType FailoverSetEligibilityResultResponseArgs

func FailoverSetEligibilityResultResponsePtr(v *FailoverSetEligibilityResultResponseArgs) FailoverSetEligibilityResultResponsePtrInput {
	return (*failoverSetEligibilityResultResponsePtrType)(v)
}

func (*failoverSetEligibilityResultResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverSetEligibilityResultResponse)(nil)).Elem()
}

func (i *failoverSetEligibilityResultResponsePtrType) ToFailoverSetEligibilityResultResponsePtrOutput() FailoverSetEligibilityResultResponsePtrOutput {
	return i.ToFailoverSetEligibilityResultResponsePtrOutputWithContext(context.Background())
}

func (i *failoverSetEligibilityResultResponsePtrType) ToFailoverSetEligibilityResultResponsePtrOutputWithContext(ctx context.Context) FailoverSetEligibilityResultResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverSetEligibilityResultResponsePtrOutput)
}

type FailoverSetEligibilityResultResponseOutput struct{ *pulumi.OutputState }

func (FailoverSetEligibilityResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverSetEligibilityResultResponse)(nil)).Elem()
}

func (o FailoverSetEligibilityResultResponseOutput) ToFailoverSetEligibilityResultResponseOutput() FailoverSetEligibilityResultResponseOutput {
	return o
}

func (o FailoverSetEligibilityResultResponseOutput) ToFailoverSetEligibilityResultResponseOutputWithContext(ctx context.Context) FailoverSetEligibilityResultResponseOutput {
	return o
}

func (o FailoverSetEligibilityResultResponseOutput) ToFailoverSetEligibilityResultResponsePtrOutput() FailoverSetEligibilityResultResponsePtrOutput {
	return o.ToFailoverSetEligibilityResultResponsePtrOutputWithContext(context.Background())
}

func (o FailoverSetEligibilityResultResponseOutput) ToFailoverSetEligibilityResultResponsePtrOutputWithContext(ctx context.Context) FailoverSetEligibilityResultResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FailoverSetEligibilityResultResponse) *FailoverSetEligibilityResultResponse {
		return &v
	}).(FailoverSetEligibilityResultResponsePtrOutput)
}

func (o FailoverSetEligibilityResultResponseOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverSetEligibilityResultResponse) *string { return v.ErrorMessage }).(pulumi.StringPtrOutput)
}

func (o FailoverSetEligibilityResultResponseOutput) IsEligibleForFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FailoverSetEligibilityResultResponse) *bool { return v.IsEligibleForFailover }).(pulumi.BoolPtrOutput)
}

type FailoverSetEligibilityResultResponsePtrOutput struct{ *pulumi.OutputState }

func (FailoverSetEligibilityResultResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverSetEligibilityResultResponse)(nil)).Elem()
}

func (o FailoverSetEligibilityResultResponsePtrOutput) ToFailoverSetEligibilityResultResponsePtrOutput() FailoverSetEligibilityResultResponsePtrOutput {
	return o
}

func (o FailoverSetEligibilityResultResponsePtrOutput) ToFailoverSetEligibilityResultResponsePtrOutputWithContext(ctx context.Context) FailoverSetEligibilityResultResponsePtrOutput {
	return o
}

func (o FailoverSetEligibilityResultResponsePtrOutput) Elem() FailoverSetEligibilityResultResponseOutput {
	return o.ApplyT(func(v *FailoverSetEligibilityResultResponse) FailoverSetEligibilityResultResponse {
		if v != nil {
			return *v
		}
		var ret FailoverSetEligibilityResultResponse
		return ret
	}).(FailoverSetEligibilityResultResponseOutput)
}

func (o FailoverSetEligibilityResultResponsePtrOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FailoverSetEligibilityResultResponse) *string {
		if v == nil {
			return nil
		}
		return v.ErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o FailoverSetEligibilityResultResponsePtrOutput) IsEligibleForFailover() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FailoverSetEligibilityResultResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsEligibleForFailover
	}).(pulumi.BoolPtrOutput)
}

type FailoverSetResponse struct {
	EligibilityResult *FailoverSetEligibilityResultResponse     `pulumi:"eligibilityResult"`
	VolumeContainers  []VolumeContainerFailoverMetadataResponse `pulumi:"volumeContainers"`
}





type FailoverSetResponseInput interface {
	pulumi.Input

	ToFailoverSetResponseOutput() FailoverSetResponseOutput
	ToFailoverSetResponseOutputWithContext(context.Context) FailoverSetResponseOutput
}

type FailoverSetResponseArgs struct {
	EligibilityResult FailoverSetEligibilityResultResponsePtrInput      `pulumi:"eligibilityResult"`
	VolumeContainers  VolumeContainerFailoverMetadataResponseArrayInput `pulumi:"volumeContainers"`
}

func (FailoverSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverSetResponse)(nil)).Elem()
}

func (i FailoverSetResponseArgs) ToFailoverSetResponseOutput() FailoverSetResponseOutput {
	return i.ToFailoverSetResponseOutputWithContext(context.Background())
}

func (i FailoverSetResponseArgs) ToFailoverSetResponseOutputWithContext(ctx context.Context) FailoverSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverSetResponseOutput)
}





type FailoverSetResponseArrayInput interface {
	pulumi.Input

	ToFailoverSetResponseArrayOutput() FailoverSetResponseArrayOutput
	ToFailoverSetResponseArrayOutputWithContext(context.Context) FailoverSetResponseArrayOutput
}

type FailoverSetResponseArray []FailoverSetResponseInput

func (FailoverSetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FailoverSetResponse)(nil)).Elem()
}

func (i FailoverSetResponseArray) ToFailoverSetResponseArrayOutput() FailoverSetResponseArrayOutput {
	return i.ToFailoverSetResponseArrayOutputWithContext(context.Background())
}

func (i FailoverSetResponseArray) ToFailoverSetResponseArrayOutputWithContext(ctx context.Context) FailoverSetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverSetResponseArrayOutput)
}

type FailoverSetResponseOutput struct{ *pulumi.OutputState }

func (FailoverSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverSetResponse)(nil)).Elem()
}

func (o FailoverSetResponseOutput) ToFailoverSetResponseOutput() FailoverSetResponseOutput {
	return o
}

func (o FailoverSetResponseOutput) ToFailoverSetResponseOutputWithContext(ctx context.Context) FailoverSetResponseOutput {
	return o
}

func (o FailoverSetResponseOutput) EligibilityResult() FailoverSetEligibilityResultResponsePtrOutput {
	return o.ApplyT(func(v FailoverSetResponse) *FailoverSetEligibilityResultResponse { return v.EligibilityResult }).(FailoverSetEligibilityResultResponsePtrOutput)
}

func (o FailoverSetResponseOutput) VolumeContainers() VolumeContainerFailoverMetadataResponseArrayOutput {
	return o.ApplyT(func(v FailoverSetResponse) []VolumeContainerFailoverMetadataResponse { return v.VolumeContainers }).(VolumeContainerFailoverMetadataResponseArrayOutput)
}

type FailoverSetResponseArrayOutput struct{ *pulumi.OutputState }

func (FailoverSetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FailoverSetResponse)(nil)).Elem()
}

func (o FailoverSetResponseArrayOutput) ToFailoverSetResponseArrayOutput() FailoverSetResponseArrayOutput {
	return o
}

func (o FailoverSetResponseArrayOutput) ToFailoverSetResponseArrayOutputWithContext(ctx context.Context) FailoverSetResponseArrayOutput {
	return o
}

func (o FailoverSetResponseArrayOutput) Index(i pulumi.IntInput) FailoverSetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FailoverSetResponse {
		return vs[0].([]FailoverSetResponse)[vs[1].(int)]
	}).(FailoverSetResponseOutput)
}

type FailoverTargetResponse struct {
	AvailableLocalStorageInBytes  *float64                         `pulumi:"availableLocalStorageInBytes"`
	AvailableTieredStorageInBytes *float64                         `pulumi:"availableTieredStorageInBytes"`
	DataContainersCount           *int                             `pulumi:"dataContainersCount"`
	DeviceId                      *string                          `pulumi:"deviceId"`
	DeviceLocation                *string                          `pulumi:"deviceLocation"`
	DeviceSoftwareVersion         *string                          `pulumi:"deviceSoftwareVersion"`
	DeviceStatus                  *string                          `pulumi:"deviceStatus"`
	EligibilityResult             *TargetEligibilityResultResponse `pulumi:"eligibilityResult"`
	FriendlyDeviceSoftwareVersion *string                          `pulumi:"friendlyDeviceSoftwareVersion"`
	ModelDescription              *string                          `pulumi:"modelDescription"`
	VolumesCount                  *int                             `pulumi:"volumesCount"`
}





type FailoverTargetResponseInput interface {
	pulumi.Input

	ToFailoverTargetResponseOutput() FailoverTargetResponseOutput
	ToFailoverTargetResponseOutputWithContext(context.Context) FailoverTargetResponseOutput
}

type FailoverTargetResponseArgs struct {
	AvailableLocalStorageInBytes  pulumi.Float64PtrInput                  `pulumi:"availableLocalStorageInBytes"`
	AvailableTieredStorageInBytes pulumi.Float64PtrInput                  `pulumi:"availableTieredStorageInBytes"`
	DataContainersCount           pulumi.IntPtrInput                      `pulumi:"dataContainersCount"`
	DeviceId                      pulumi.StringPtrInput                   `pulumi:"deviceId"`
	DeviceLocation                pulumi.StringPtrInput                   `pulumi:"deviceLocation"`
	DeviceSoftwareVersion         pulumi.StringPtrInput                   `pulumi:"deviceSoftwareVersion"`
	DeviceStatus                  pulumi.StringPtrInput                   `pulumi:"deviceStatus"`
	EligibilityResult             TargetEligibilityResultResponsePtrInput `pulumi:"eligibilityResult"`
	FriendlyDeviceSoftwareVersion pulumi.StringPtrInput                   `pulumi:"friendlyDeviceSoftwareVersion"`
	ModelDescription              pulumi.StringPtrInput                   `pulumi:"modelDescription"`
	VolumesCount                  pulumi.IntPtrInput                      `pulumi:"volumesCount"`
}

func (FailoverTargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverTargetResponse)(nil)).Elem()
}

func (i FailoverTargetResponseArgs) ToFailoverTargetResponseOutput() FailoverTargetResponseOutput {
	return i.ToFailoverTargetResponseOutputWithContext(context.Background())
}

func (i FailoverTargetResponseArgs) ToFailoverTargetResponseOutputWithContext(ctx context.Context) FailoverTargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverTargetResponseOutput)
}





type FailoverTargetResponseArrayInput interface {
	pulumi.Input

	ToFailoverTargetResponseArrayOutput() FailoverTargetResponseArrayOutput
	ToFailoverTargetResponseArrayOutputWithContext(context.Context) FailoverTargetResponseArrayOutput
}

type FailoverTargetResponseArray []FailoverTargetResponseInput

func (FailoverTargetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FailoverTargetResponse)(nil)).Elem()
}

func (i FailoverTargetResponseArray) ToFailoverTargetResponseArrayOutput() FailoverTargetResponseArrayOutput {
	return i.ToFailoverTargetResponseArrayOutputWithContext(context.Background())
}

func (i FailoverTargetResponseArray) ToFailoverTargetResponseArrayOutputWithContext(ctx context.Context) FailoverTargetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverTargetResponseArrayOutput)
}

type FailoverTargetResponseOutput struct{ *pulumi.OutputState }

func (FailoverTargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverTargetResponse)(nil)).Elem()
}

func (o FailoverTargetResponseOutput) ToFailoverTargetResponseOutput() FailoverTargetResponseOutput {
	return o
}

func (o FailoverTargetResponseOutput) ToFailoverTargetResponseOutputWithContext(ctx context.Context) FailoverTargetResponseOutput {
	return o
}

func (o FailoverTargetResponseOutput) AvailableLocalStorageInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *float64 { return v.AvailableLocalStorageInBytes }).(pulumi.Float64PtrOutput)
}

func (o FailoverTargetResponseOutput) AvailableTieredStorageInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *float64 { return v.AvailableTieredStorageInBytes }).(pulumi.Float64PtrOutput)
}

func (o FailoverTargetResponseOutput) DataContainersCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *int { return v.DataContainersCount }).(pulumi.IntPtrOutput)
}

func (o FailoverTargetResponseOutput) DeviceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *string { return v.DeviceId }).(pulumi.StringPtrOutput)
}

func (o FailoverTargetResponseOutput) DeviceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *string { return v.DeviceLocation }).(pulumi.StringPtrOutput)
}

func (o FailoverTargetResponseOutput) DeviceSoftwareVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *string { return v.DeviceSoftwareVersion }).(pulumi.StringPtrOutput)
}

func (o FailoverTargetResponseOutput) DeviceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *string { return v.DeviceStatus }).(pulumi.StringPtrOutput)
}

func (o FailoverTargetResponseOutput) EligibilityResult() TargetEligibilityResultResponsePtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *TargetEligibilityResultResponse { return v.EligibilityResult }).(TargetEligibilityResultResponsePtrOutput)
}

func (o FailoverTargetResponseOutput) FriendlyDeviceSoftwareVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *string { return v.FriendlyDeviceSoftwareVersion }).(pulumi.StringPtrOutput)
}

func (o FailoverTargetResponseOutput) ModelDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *string { return v.ModelDescription }).(pulumi.StringPtrOutput)
}

func (o FailoverTargetResponseOutput) VolumesCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverTargetResponse) *int { return v.VolumesCount }).(pulumi.IntPtrOutput)
}

type FailoverTargetResponseArrayOutput struct{ *pulumi.OutputState }

func (FailoverTargetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FailoverTargetResponse)(nil)).Elem()
}

func (o FailoverTargetResponseArrayOutput) ToFailoverTargetResponseArrayOutput() FailoverTargetResponseArrayOutput {
	return o
}

func (o FailoverTargetResponseArrayOutput) ToFailoverTargetResponseArrayOutputWithContext(ctx context.Context) FailoverTargetResponseArrayOutput {
	return o
}

func (o FailoverTargetResponseArrayOutput) Index(i pulumi.IntInput) FailoverTargetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FailoverTargetResponse {
		return vs[0].([]FailoverTargetResponse)[vs[1].(int)]
	}).(FailoverTargetResponseOutput)
}

type ManagerIntrinsicSettings struct {
	Type ManagerType `pulumi:"type"`
}





type ManagerIntrinsicSettingsInput interface {
	pulumi.Input

	ToManagerIntrinsicSettingsOutput() ManagerIntrinsicSettingsOutput
	ToManagerIntrinsicSettingsOutputWithContext(context.Context) ManagerIntrinsicSettingsOutput
}

type ManagerIntrinsicSettingsArgs struct {
	Type ManagerTypeInput `pulumi:"type"`
}

func (ManagerIntrinsicSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerIntrinsicSettings)(nil)).Elem()
}

func (i ManagerIntrinsicSettingsArgs) ToManagerIntrinsicSettingsOutput() ManagerIntrinsicSettingsOutput {
	return i.ToManagerIntrinsicSettingsOutputWithContext(context.Background())
}

func (i ManagerIntrinsicSettingsArgs) ToManagerIntrinsicSettingsOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsOutput)
}

func (i ManagerIntrinsicSettingsArgs) ToManagerIntrinsicSettingsPtrOutput() ManagerIntrinsicSettingsPtrOutput {
	return i.ToManagerIntrinsicSettingsPtrOutputWithContext(context.Background())
}

func (i ManagerIntrinsicSettingsArgs) ToManagerIntrinsicSettingsPtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsOutput).ToManagerIntrinsicSettingsPtrOutputWithContext(ctx)
}









type ManagerIntrinsicSettingsPtrInput interface {
	pulumi.Input

	ToManagerIntrinsicSettingsPtrOutput() ManagerIntrinsicSettingsPtrOutput
	ToManagerIntrinsicSettingsPtrOutputWithContext(context.Context) ManagerIntrinsicSettingsPtrOutput
}

type managerIntrinsicSettingsPtrType ManagerIntrinsicSettingsArgs

func ManagerIntrinsicSettingsPtr(v *ManagerIntrinsicSettingsArgs) ManagerIntrinsicSettingsPtrInput {
	return (*managerIntrinsicSettingsPtrType)(v)
}

func (*managerIntrinsicSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerIntrinsicSettings)(nil)).Elem()
}

func (i *managerIntrinsicSettingsPtrType) ToManagerIntrinsicSettingsPtrOutput() ManagerIntrinsicSettingsPtrOutput {
	return i.ToManagerIntrinsicSettingsPtrOutputWithContext(context.Background())
}

func (i *managerIntrinsicSettingsPtrType) ToManagerIntrinsicSettingsPtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsPtrOutput)
}

type ManagerIntrinsicSettingsOutput struct{ *pulumi.OutputState }

func (ManagerIntrinsicSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerIntrinsicSettings)(nil)).Elem()
}

func (o ManagerIntrinsicSettingsOutput) ToManagerIntrinsicSettingsOutput() ManagerIntrinsicSettingsOutput {
	return o
}

func (o ManagerIntrinsicSettingsOutput) ToManagerIntrinsicSettingsOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsOutput {
	return o
}

func (o ManagerIntrinsicSettingsOutput) ToManagerIntrinsicSettingsPtrOutput() ManagerIntrinsicSettingsPtrOutput {
	return o.ToManagerIntrinsicSettingsPtrOutputWithContext(context.Background())
}

func (o ManagerIntrinsicSettingsOutput) ToManagerIntrinsicSettingsPtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagerIntrinsicSettings) *ManagerIntrinsicSettings {
		return &v
	}).(ManagerIntrinsicSettingsPtrOutput)
}

func (o ManagerIntrinsicSettingsOutput) Type() ManagerTypeOutput {
	return o.ApplyT(func(v ManagerIntrinsicSettings) ManagerType { return v.Type }).(ManagerTypeOutput)
}

type ManagerIntrinsicSettingsPtrOutput struct{ *pulumi.OutputState }

func (ManagerIntrinsicSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerIntrinsicSettings)(nil)).Elem()
}

func (o ManagerIntrinsicSettingsPtrOutput) ToManagerIntrinsicSettingsPtrOutput() ManagerIntrinsicSettingsPtrOutput {
	return o
}

func (o ManagerIntrinsicSettingsPtrOutput) ToManagerIntrinsicSettingsPtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsPtrOutput {
	return o
}

func (o ManagerIntrinsicSettingsPtrOutput) Elem() ManagerIntrinsicSettingsOutput {
	return o.ApplyT(func(v *ManagerIntrinsicSettings) ManagerIntrinsicSettings {
		if v != nil {
			return *v
		}
		var ret ManagerIntrinsicSettings
		return ret
	}).(ManagerIntrinsicSettingsOutput)
}

func (o ManagerIntrinsicSettingsPtrOutput) Type() ManagerTypePtrOutput {
	return o.ApplyT(func(v *ManagerIntrinsicSettings) *ManagerType {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(ManagerTypePtrOutput)
}

type ManagerIntrinsicSettingsResponse struct {
	Type string `pulumi:"type"`
}





type ManagerIntrinsicSettingsResponseInput interface {
	pulumi.Input

	ToManagerIntrinsicSettingsResponseOutput() ManagerIntrinsicSettingsResponseOutput
	ToManagerIntrinsicSettingsResponseOutputWithContext(context.Context) ManagerIntrinsicSettingsResponseOutput
}

type ManagerIntrinsicSettingsResponseArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (ManagerIntrinsicSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerIntrinsicSettingsResponse)(nil)).Elem()
}

func (i ManagerIntrinsicSettingsResponseArgs) ToManagerIntrinsicSettingsResponseOutput() ManagerIntrinsicSettingsResponseOutput {
	return i.ToManagerIntrinsicSettingsResponseOutputWithContext(context.Background())
}

func (i ManagerIntrinsicSettingsResponseArgs) ToManagerIntrinsicSettingsResponseOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsResponseOutput)
}

func (i ManagerIntrinsicSettingsResponseArgs) ToManagerIntrinsicSettingsResponsePtrOutput() ManagerIntrinsicSettingsResponsePtrOutput {
	return i.ToManagerIntrinsicSettingsResponsePtrOutputWithContext(context.Background())
}

func (i ManagerIntrinsicSettingsResponseArgs) ToManagerIntrinsicSettingsResponsePtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsResponseOutput).ToManagerIntrinsicSettingsResponsePtrOutputWithContext(ctx)
}









type ManagerIntrinsicSettingsResponsePtrInput interface {
	pulumi.Input

	ToManagerIntrinsicSettingsResponsePtrOutput() ManagerIntrinsicSettingsResponsePtrOutput
	ToManagerIntrinsicSettingsResponsePtrOutputWithContext(context.Context) ManagerIntrinsicSettingsResponsePtrOutput
}

type managerIntrinsicSettingsResponsePtrType ManagerIntrinsicSettingsResponseArgs

func ManagerIntrinsicSettingsResponsePtr(v *ManagerIntrinsicSettingsResponseArgs) ManagerIntrinsicSettingsResponsePtrInput {
	return (*managerIntrinsicSettingsResponsePtrType)(v)
}

func (*managerIntrinsicSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerIntrinsicSettingsResponse)(nil)).Elem()
}

func (i *managerIntrinsicSettingsResponsePtrType) ToManagerIntrinsicSettingsResponsePtrOutput() ManagerIntrinsicSettingsResponsePtrOutput {
	return i.ToManagerIntrinsicSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *managerIntrinsicSettingsResponsePtrType) ToManagerIntrinsicSettingsResponsePtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerIntrinsicSettingsResponsePtrOutput)
}

type ManagerIntrinsicSettingsResponseOutput struct{ *pulumi.OutputState }

func (ManagerIntrinsicSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerIntrinsicSettingsResponse)(nil)).Elem()
}

func (o ManagerIntrinsicSettingsResponseOutput) ToManagerIntrinsicSettingsResponseOutput() ManagerIntrinsicSettingsResponseOutput {
	return o
}

func (o ManagerIntrinsicSettingsResponseOutput) ToManagerIntrinsicSettingsResponseOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponseOutput {
	return o
}

func (o ManagerIntrinsicSettingsResponseOutput) ToManagerIntrinsicSettingsResponsePtrOutput() ManagerIntrinsicSettingsResponsePtrOutput {
	return o.ToManagerIntrinsicSettingsResponsePtrOutputWithContext(context.Background())
}

func (o ManagerIntrinsicSettingsResponseOutput) ToManagerIntrinsicSettingsResponsePtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagerIntrinsicSettingsResponse) *ManagerIntrinsicSettingsResponse {
		return &v
	}).(ManagerIntrinsicSettingsResponsePtrOutput)
}

func (o ManagerIntrinsicSettingsResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ManagerIntrinsicSettingsResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ManagerIntrinsicSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagerIntrinsicSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerIntrinsicSettingsResponse)(nil)).Elem()
}

func (o ManagerIntrinsicSettingsResponsePtrOutput) ToManagerIntrinsicSettingsResponsePtrOutput() ManagerIntrinsicSettingsResponsePtrOutput {
	return o
}

func (o ManagerIntrinsicSettingsResponsePtrOutput) ToManagerIntrinsicSettingsResponsePtrOutputWithContext(ctx context.Context) ManagerIntrinsicSettingsResponsePtrOutput {
	return o
}

func (o ManagerIntrinsicSettingsResponsePtrOutput) Elem() ManagerIntrinsicSettingsResponseOutput {
	return o.ApplyT(func(v *ManagerIntrinsicSettingsResponse) ManagerIntrinsicSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ManagerIntrinsicSettingsResponse
		return ret
	}).(ManagerIntrinsicSettingsResponseOutput)
}

func (o ManagerIntrinsicSettingsResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagerIntrinsicSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagerSku struct {
	Name ManagerSkuType `pulumi:"name"`
}





type ManagerSkuInput interface {
	pulumi.Input

	ToManagerSkuOutput() ManagerSkuOutput
	ToManagerSkuOutputWithContext(context.Context) ManagerSkuOutput
}

type ManagerSkuArgs struct {
	Name ManagerSkuTypeInput `pulumi:"name"`
}

func (ManagerSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerSku)(nil)).Elem()
}

func (i ManagerSkuArgs) ToManagerSkuOutput() ManagerSkuOutput {
	return i.ToManagerSkuOutputWithContext(context.Background())
}

func (i ManagerSkuArgs) ToManagerSkuOutputWithContext(ctx context.Context) ManagerSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuOutput)
}

func (i ManagerSkuArgs) ToManagerSkuPtrOutput() ManagerSkuPtrOutput {
	return i.ToManagerSkuPtrOutputWithContext(context.Background())
}

func (i ManagerSkuArgs) ToManagerSkuPtrOutputWithContext(ctx context.Context) ManagerSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuOutput).ToManagerSkuPtrOutputWithContext(ctx)
}









type ManagerSkuPtrInput interface {
	pulumi.Input

	ToManagerSkuPtrOutput() ManagerSkuPtrOutput
	ToManagerSkuPtrOutputWithContext(context.Context) ManagerSkuPtrOutput
}

type managerSkuPtrType ManagerSkuArgs

func ManagerSkuPtr(v *ManagerSkuArgs) ManagerSkuPtrInput {
	return (*managerSkuPtrType)(v)
}

func (*managerSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerSku)(nil)).Elem()
}

func (i *managerSkuPtrType) ToManagerSkuPtrOutput() ManagerSkuPtrOutput {
	return i.ToManagerSkuPtrOutputWithContext(context.Background())
}

func (i *managerSkuPtrType) ToManagerSkuPtrOutputWithContext(ctx context.Context) ManagerSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuPtrOutput)
}

type ManagerSkuOutput struct{ *pulumi.OutputState }

func (ManagerSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerSku)(nil)).Elem()
}

func (o ManagerSkuOutput) ToManagerSkuOutput() ManagerSkuOutput {
	return o
}

func (o ManagerSkuOutput) ToManagerSkuOutputWithContext(ctx context.Context) ManagerSkuOutput {
	return o
}

func (o ManagerSkuOutput) ToManagerSkuPtrOutput() ManagerSkuPtrOutput {
	return o.ToManagerSkuPtrOutputWithContext(context.Background())
}

func (o ManagerSkuOutput) ToManagerSkuPtrOutputWithContext(ctx context.Context) ManagerSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagerSku) *ManagerSku {
		return &v
	}).(ManagerSkuPtrOutput)
}

func (o ManagerSkuOutput) Name() ManagerSkuTypeOutput {
	return o.ApplyT(func(v ManagerSku) ManagerSkuType { return v.Name }).(ManagerSkuTypeOutput)
}

type ManagerSkuPtrOutput struct{ *pulumi.OutputState }

func (ManagerSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerSku)(nil)).Elem()
}

func (o ManagerSkuPtrOutput) ToManagerSkuPtrOutput() ManagerSkuPtrOutput {
	return o
}

func (o ManagerSkuPtrOutput) ToManagerSkuPtrOutputWithContext(ctx context.Context) ManagerSkuPtrOutput {
	return o
}

func (o ManagerSkuPtrOutput) Elem() ManagerSkuOutput {
	return o.ApplyT(func(v *ManagerSku) ManagerSku {
		if v != nil {
			return *v
		}
		var ret ManagerSku
		return ret
	}).(ManagerSkuOutput)
}

func (o ManagerSkuPtrOutput) Name() ManagerSkuTypePtrOutput {
	return o.ApplyT(func(v *ManagerSku) *ManagerSkuType {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(ManagerSkuTypePtrOutput)
}

type ManagerSkuResponse struct {
	Name string `pulumi:"name"`
}





type ManagerSkuResponseInput interface {
	pulumi.Input

	ToManagerSkuResponseOutput() ManagerSkuResponseOutput
	ToManagerSkuResponseOutputWithContext(context.Context) ManagerSkuResponseOutput
}

type ManagerSkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (ManagerSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerSkuResponse)(nil)).Elem()
}

func (i ManagerSkuResponseArgs) ToManagerSkuResponseOutput() ManagerSkuResponseOutput {
	return i.ToManagerSkuResponseOutputWithContext(context.Background())
}

func (i ManagerSkuResponseArgs) ToManagerSkuResponseOutputWithContext(ctx context.Context) ManagerSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuResponseOutput)
}

func (i ManagerSkuResponseArgs) ToManagerSkuResponsePtrOutput() ManagerSkuResponsePtrOutput {
	return i.ToManagerSkuResponsePtrOutputWithContext(context.Background())
}

func (i ManagerSkuResponseArgs) ToManagerSkuResponsePtrOutputWithContext(ctx context.Context) ManagerSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuResponseOutput).ToManagerSkuResponsePtrOutputWithContext(ctx)
}









type ManagerSkuResponsePtrInput interface {
	pulumi.Input

	ToManagerSkuResponsePtrOutput() ManagerSkuResponsePtrOutput
	ToManagerSkuResponsePtrOutputWithContext(context.Context) ManagerSkuResponsePtrOutput
}

type managerSkuResponsePtrType ManagerSkuResponseArgs

func ManagerSkuResponsePtr(v *ManagerSkuResponseArgs) ManagerSkuResponsePtrInput {
	return (*managerSkuResponsePtrType)(v)
}

func (*managerSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerSkuResponse)(nil)).Elem()
}

func (i *managerSkuResponsePtrType) ToManagerSkuResponsePtrOutput() ManagerSkuResponsePtrOutput {
	return i.ToManagerSkuResponsePtrOutputWithContext(context.Background())
}

func (i *managerSkuResponsePtrType) ToManagerSkuResponsePtrOutputWithContext(ctx context.Context) ManagerSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagerSkuResponsePtrOutput)
}

type ManagerSkuResponseOutput struct{ *pulumi.OutputState }

func (ManagerSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagerSkuResponse)(nil)).Elem()
}

func (o ManagerSkuResponseOutput) ToManagerSkuResponseOutput() ManagerSkuResponseOutput {
	return o
}

func (o ManagerSkuResponseOutput) ToManagerSkuResponseOutputWithContext(ctx context.Context) ManagerSkuResponseOutput {
	return o
}

func (o ManagerSkuResponseOutput) ToManagerSkuResponsePtrOutput() ManagerSkuResponsePtrOutput {
	return o.ToManagerSkuResponsePtrOutputWithContext(context.Background())
}

func (o ManagerSkuResponseOutput) ToManagerSkuResponsePtrOutputWithContext(ctx context.Context) ManagerSkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagerSkuResponse) *ManagerSkuResponse {
		return &v
	}).(ManagerSkuResponsePtrOutput)
}

func (o ManagerSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ManagerSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ManagerSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagerSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagerSkuResponse)(nil)).Elem()
}

func (o ManagerSkuResponsePtrOutput) ToManagerSkuResponsePtrOutput() ManagerSkuResponsePtrOutput {
	return o
}

func (o ManagerSkuResponsePtrOutput) ToManagerSkuResponsePtrOutputWithContext(ctx context.Context) ManagerSkuResponsePtrOutput {
	return o
}

func (o ManagerSkuResponsePtrOutput) Elem() ManagerSkuResponseOutput {
	return o.ApplyT(func(v *ManagerSkuResponse) ManagerSkuResponse {
		if v != nil {
			return *v
		}
		var ret ManagerSkuResponse
		return ret
	}).(ManagerSkuResponseOutput)
}

func (o ManagerSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagerSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type ScheduleRecurrence struct {
	RecurrenceType  RecurrenceType `pulumi:"recurrenceType"`
	RecurrenceValue int            `pulumi:"recurrenceValue"`
	WeeklyDaysList  []DayOfWeek    `pulumi:"weeklyDaysList"`
}





type ScheduleRecurrenceInput interface {
	pulumi.Input

	ToScheduleRecurrenceOutput() ScheduleRecurrenceOutput
	ToScheduleRecurrenceOutputWithContext(context.Context) ScheduleRecurrenceOutput
}

type ScheduleRecurrenceArgs struct {
	RecurrenceType  RecurrenceTypeInput `pulumi:"recurrenceType"`
	RecurrenceValue pulumi.IntInput     `pulumi:"recurrenceValue"`
	WeeklyDaysList  DayOfWeekArrayInput `pulumi:"weeklyDaysList"`
}

func (ScheduleRecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleRecurrence)(nil)).Elem()
}

func (i ScheduleRecurrenceArgs) ToScheduleRecurrenceOutput() ScheduleRecurrenceOutput {
	return i.ToScheduleRecurrenceOutputWithContext(context.Background())
}

func (i ScheduleRecurrenceArgs) ToScheduleRecurrenceOutputWithContext(ctx context.Context) ScheduleRecurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleRecurrenceOutput)
}

func (i ScheduleRecurrenceArgs) ToScheduleRecurrencePtrOutput() ScheduleRecurrencePtrOutput {
	return i.ToScheduleRecurrencePtrOutputWithContext(context.Background())
}

func (i ScheduleRecurrenceArgs) ToScheduleRecurrencePtrOutputWithContext(ctx context.Context) ScheduleRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleRecurrenceOutput).ToScheduleRecurrencePtrOutputWithContext(ctx)
}









type ScheduleRecurrencePtrInput interface {
	pulumi.Input

	ToScheduleRecurrencePtrOutput() ScheduleRecurrencePtrOutput
	ToScheduleRecurrencePtrOutputWithContext(context.Context) ScheduleRecurrencePtrOutput
}

type scheduleRecurrencePtrType ScheduleRecurrenceArgs

func ScheduleRecurrencePtr(v *ScheduleRecurrenceArgs) ScheduleRecurrencePtrInput {
	return (*scheduleRecurrencePtrType)(v)
}

func (*scheduleRecurrencePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleRecurrence)(nil)).Elem()
}

func (i *scheduleRecurrencePtrType) ToScheduleRecurrencePtrOutput() ScheduleRecurrencePtrOutput {
	return i.ToScheduleRecurrencePtrOutputWithContext(context.Background())
}

func (i *scheduleRecurrencePtrType) ToScheduleRecurrencePtrOutputWithContext(ctx context.Context) ScheduleRecurrencePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleRecurrencePtrOutput)
}

type ScheduleRecurrenceOutput struct{ *pulumi.OutputState }

func (ScheduleRecurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleRecurrence)(nil)).Elem()
}

func (o ScheduleRecurrenceOutput) ToScheduleRecurrenceOutput() ScheduleRecurrenceOutput {
	return o
}

func (o ScheduleRecurrenceOutput) ToScheduleRecurrenceOutputWithContext(ctx context.Context) ScheduleRecurrenceOutput {
	return o
}

func (o ScheduleRecurrenceOutput) ToScheduleRecurrencePtrOutput() ScheduleRecurrencePtrOutput {
	return o.ToScheduleRecurrencePtrOutputWithContext(context.Background())
}

func (o ScheduleRecurrenceOutput) ToScheduleRecurrencePtrOutputWithContext(ctx context.Context) ScheduleRecurrencePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleRecurrence) *ScheduleRecurrence {
		return &v
	}).(ScheduleRecurrencePtrOutput)
}

func (o ScheduleRecurrenceOutput) RecurrenceType() RecurrenceTypeOutput {
	return o.ApplyT(func(v ScheduleRecurrence) RecurrenceType { return v.RecurrenceType }).(RecurrenceTypeOutput)
}

func (o ScheduleRecurrenceOutput) RecurrenceValue() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleRecurrence) int { return v.RecurrenceValue }).(pulumi.IntOutput)
}

func (o ScheduleRecurrenceOutput) WeeklyDaysList() DayOfWeekArrayOutput {
	return o.ApplyT(func(v ScheduleRecurrence) []DayOfWeek { return v.WeeklyDaysList }).(DayOfWeekArrayOutput)
}

type ScheduleRecurrencePtrOutput struct{ *pulumi.OutputState }

func (ScheduleRecurrencePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleRecurrence)(nil)).Elem()
}

func (o ScheduleRecurrencePtrOutput) ToScheduleRecurrencePtrOutput() ScheduleRecurrencePtrOutput {
	return o
}

func (o ScheduleRecurrencePtrOutput) ToScheduleRecurrencePtrOutputWithContext(ctx context.Context) ScheduleRecurrencePtrOutput {
	return o
}

func (o ScheduleRecurrencePtrOutput) Elem() ScheduleRecurrenceOutput {
	return o.ApplyT(func(v *ScheduleRecurrence) ScheduleRecurrence {
		if v != nil {
			return *v
		}
		var ret ScheduleRecurrence
		return ret
	}).(ScheduleRecurrenceOutput)
}

func (o ScheduleRecurrencePtrOutput) RecurrenceType() RecurrenceTypePtrOutput {
	return o.ApplyT(func(v *ScheduleRecurrence) *RecurrenceType {
		if v == nil {
			return nil
		}
		return &v.RecurrenceType
	}).(RecurrenceTypePtrOutput)
}

func (o ScheduleRecurrencePtrOutput) RecurrenceValue() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduleRecurrence) *int {
		if v == nil {
			return nil
		}
		return &v.RecurrenceValue
	}).(pulumi.IntPtrOutput)
}

func (o ScheduleRecurrencePtrOutput) WeeklyDaysList() DayOfWeekArrayOutput {
	return o.ApplyT(func(v *ScheduleRecurrence) []DayOfWeek {
		if v == nil {
			return nil
		}
		return v.WeeklyDaysList
	}).(DayOfWeekArrayOutput)
}

type ScheduleRecurrenceResponse struct {
	RecurrenceType  string   `pulumi:"recurrenceType"`
	RecurrenceValue int      `pulumi:"recurrenceValue"`
	WeeklyDaysList  []string `pulumi:"weeklyDaysList"`
}





type ScheduleRecurrenceResponseInput interface {
	pulumi.Input

	ToScheduleRecurrenceResponseOutput() ScheduleRecurrenceResponseOutput
	ToScheduleRecurrenceResponseOutputWithContext(context.Context) ScheduleRecurrenceResponseOutput
}

type ScheduleRecurrenceResponseArgs struct {
	RecurrenceType  pulumi.StringInput      `pulumi:"recurrenceType"`
	RecurrenceValue pulumi.IntInput         `pulumi:"recurrenceValue"`
	WeeklyDaysList  pulumi.StringArrayInput `pulumi:"weeklyDaysList"`
}

func (ScheduleRecurrenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleRecurrenceResponse)(nil)).Elem()
}

func (i ScheduleRecurrenceResponseArgs) ToScheduleRecurrenceResponseOutput() ScheduleRecurrenceResponseOutput {
	return i.ToScheduleRecurrenceResponseOutputWithContext(context.Background())
}

func (i ScheduleRecurrenceResponseArgs) ToScheduleRecurrenceResponseOutputWithContext(ctx context.Context) ScheduleRecurrenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleRecurrenceResponseOutput)
}

func (i ScheduleRecurrenceResponseArgs) ToScheduleRecurrenceResponsePtrOutput() ScheduleRecurrenceResponsePtrOutput {
	return i.ToScheduleRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (i ScheduleRecurrenceResponseArgs) ToScheduleRecurrenceResponsePtrOutputWithContext(ctx context.Context) ScheduleRecurrenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleRecurrenceResponseOutput).ToScheduleRecurrenceResponsePtrOutputWithContext(ctx)
}









type ScheduleRecurrenceResponsePtrInput interface {
	pulumi.Input

	ToScheduleRecurrenceResponsePtrOutput() ScheduleRecurrenceResponsePtrOutput
	ToScheduleRecurrenceResponsePtrOutputWithContext(context.Context) ScheduleRecurrenceResponsePtrOutput
}

type scheduleRecurrenceResponsePtrType ScheduleRecurrenceResponseArgs

func ScheduleRecurrenceResponsePtr(v *ScheduleRecurrenceResponseArgs) ScheduleRecurrenceResponsePtrInput {
	return (*scheduleRecurrenceResponsePtrType)(v)
}

func (*scheduleRecurrenceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleRecurrenceResponse)(nil)).Elem()
}

func (i *scheduleRecurrenceResponsePtrType) ToScheduleRecurrenceResponsePtrOutput() ScheduleRecurrenceResponsePtrOutput {
	return i.ToScheduleRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (i *scheduleRecurrenceResponsePtrType) ToScheduleRecurrenceResponsePtrOutputWithContext(ctx context.Context) ScheduleRecurrenceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleRecurrenceResponsePtrOutput)
}

type ScheduleRecurrenceResponseOutput struct{ *pulumi.OutputState }

func (ScheduleRecurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleRecurrenceResponse)(nil)).Elem()
}

func (o ScheduleRecurrenceResponseOutput) ToScheduleRecurrenceResponseOutput() ScheduleRecurrenceResponseOutput {
	return o
}

func (o ScheduleRecurrenceResponseOutput) ToScheduleRecurrenceResponseOutputWithContext(ctx context.Context) ScheduleRecurrenceResponseOutput {
	return o
}

func (o ScheduleRecurrenceResponseOutput) ToScheduleRecurrenceResponsePtrOutput() ScheduleRecurrenceResponsePtrOutput {
	return o.ToScheduleRecurrenceResponsePtrOutputWithContext(context.Background())
}

func (o ScheduleRecurrenceResponseOutput) ToScheduleRecurrenceResponsePtrOutputWithContext(ctx context.Context) ScheduleRecurrenceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleRecurrenceResponse) *ScheduleRecurrenceResponse {
		return &v
	}).(ScheduleRecurrenceResponsePtrOutput)
}

func (o ScheduleRecurrenceResponseOutput) RecurrenceType() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleRecurrenceResponse) string { return v.RecurrenceType }).(pulumi.StringOutput)
}

func (o ScheduleRecurrenceResponseOutput) RecurrenceValue() pulumi.IntOutput {
	return o.ApplyT(func(v ScheduleRecurrenceResponse) int { return v.RecurrenceValue }).(pulumi.IntOutput)
}

func (o ScheduleRecurrenceResponseOutput) WeeklyDaysList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScheduleRecurrenceResponse) []string { return v.WeeklyDaysList }).(pulumi.StringArrayOutput)
}

type ScheduleRecurrenceResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduleRecurrenceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleRecurrenceResponse)(nil)).Elem()
}

func (o ScheduleRecurrenceResponsePtrOutput) ToScheduleRecurrenceResponsePtrOutput() ScheduleRecurrenceResponsePtrOutput {
	return o
}

func (o ScheduleRecurrenceResponsePtrOutput) ToScheduleRecurrenceResponsePtrOutputWithContext(ctx context.Context) ScheduleRecurrenceResponsePtrOutput {
	return o
}

func (o ScheduleRecurrenceResponsePtrOutput) Elem() ScheduleRecurrenceResponseOutput {
	return o.ApplyT(func(v *ScheduleRecurrenceResponse) ScheduleRecurrenceResponse {
		if v != nil {
			return *v
		}
		var ret ScheduleRecurrenceResponse
		return ret
	}).(ScheduleRecurrenceResponseOutput)
}

func (o ScheduleRecurrenceResponsePtrOutput) RecurrenceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleRecurrenceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RecurrenceType
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleRecurrenceResponsePtrOutput) RecurrenceValue() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduleRecurrenceResponse) *int {
		if v == nil {
			return nil
		}
		return &v.RecurrenceValue
	}).(pulumi.IntPtrOutput)
}

func (o ScheduleRecurrenceResponsePtrOutput) WeeklyDaysList() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScheduleRecurrenceResponse) []string {
		if v == nil {
			return nil
		}
		return v.WeeklyDaysList
	}).(pulumi.StringArrayOutput)
}

type TargetEligibilityErrorMessageResponse struct {
	Message    *string `pulumi:"message"`
	Resolution *string `pulumi:"resolution"`
	ResultCode *string `pulumi:"resultCode"`
}





type TargetEligibilityErrorMessageResponseInput interface {
	pulumi.Input

	ToTargetEligibilityErrorMessageResponseOutput() TargetEligibilityErrorMessageResponseOutput
	ToTargetEligibilityErrorMessageResponseOutputWithContext(context.Context) TargetEligibilityErrorMessageResponseOutput
}

type TargetEligibilityErrorMessageResponseArgs struct {
	Message    pulumi.StringPtrInput `pulumi:"message"`
	Resolution pulumi.StringPtrInput `pulumi:"resolution"`
	ResultCode pulumi.StringPtrInput `pulumi:"resultCode"`
}

func (TargetEligibilityErrorMessageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetEligibilityErrorMessageResponse)(nil)).Elem()
}

func (i TargetEligibilityErrorMessageResponseArgs) ToTargetEligibilityErrorMessageResponseOutput() TargetEligibilityErrorMessageResponseOutput {
	return i.ToTargetEligibilityErrorMessageResponseOutputWithContext(context.Background())
}

func (i TargetEligibilityErrorMessageResponseArgs) ToTargetEligibilityErrorMessageResponseOutputWithContext(ctx context.Context) TargetEligibilityErrorMessageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetEligibilityErrorMessageResponseOutput)
}





type TargetEligibilityErrorMessageResponseArrayInput interface {
	pulumi.Input

	ToTargetEligibilityErrorMessageResponseArrayOutput() TargetEligibilityErrorMessageResponseArrayOutput
	ToTargetEligibilityErrorMessageResponseArrayOutputWithContext(context.Context) TargetEligibilityErrorMessageResponseArrayOutput
}

type TargetEligibilityErrorMessageResponseArray []TargetEligibilityErrorMessageResponseInput

func (TargetEligibilityErrorMessageResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetEligibilityErrorMessageResponse)(nil)).Elem()
}

func (i TargetEligibilityErrorMessageResponseArray) ToTargetEligibilityErrorMessageResponseArrayOutput() TargetEligibilityErrorMessageResponseArrayOutput {
	return i.ToTargetEligibilityErrorMessageResponseArrayOutputWithContext(context.Background())
}

func (i TargetEligibilityErrorMessageResponseArray) ToTargetEligibilityErrorMessageResponseArrayOutputWithContext(ctx context.Context) TargetEligibilityErrorMessageResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetEligibilityErrorMessageResponseArrayOutput)
}

type TargetEligibilityErrorMessageResponseOutput struct{ *pulumi.OutputState }

func (TargetEligibilityErrorMessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetEligibilityErrorMessageResponse)(nil)).Elem()
}

func (o TargetEligibilityErrorMessageResponseOutput) ToTargetEligibilityErrorMessageResponseOutput() TargetEligibilityErrorMessageResponseOutput {
	return o
}

func (o TargetEligibilityErrorMessageResponseOutput) ToTargetEligibilityErrorMessageResponseOutputWithContext(ctx context.Context) TargetEligibilityErrorMessageResponseOutput {
	return o
}

func (o TargetEligibilityErrorMessageResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetEligibilityErrorMessageResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o TargetEligibilityErrorMessageResponseOutput) Resolution() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetEligibilityErrorMessageResponse) *string { return v.Resolution }).(pulumi.StringPtrOutput)
}

func (o TargetEligibilityErrorMessageResponseOutput) ResultCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetEligibilityErrorMessageResponse) *string { return v.ResultCode }).(pulumi.StringPtrOutput)
}

type TargetEligibilityErrorMessageResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetEligibilityErrorMessageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetEligibilityErrorMessageResponse)(nil)).Elem()
}

func (o TargetEligibilityErrorMessageResponseArrayOutput) ToTargetEligibilityErrorMessageResponseArrayOutput() TargetEligibilityErrorMessageResponseArrayOutput {
	return o
}

func (o TargetEligibilityErrorMessageResponseArrayOutput) ToTargetEligibilityErrorMessageResponseArrayOutputWithContext(ctx context.Context) TargetEligibilityErrorMessageResponseArrayOutput {
	return o
}

func (o TargetEligibilityErrorMessageResponseArrayOutput) Index(i pulumi.IntInput) TargetEligibilityErrorMessageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetEligibilityErrorMessageResponse {
		return vs[0].([]TargetEligibilityErrorMessageResponse)[vs[1].(int)]
	}).(TargetEligibilityErrorMessageResponseOutput)
}

type TargetEligibilityResultResponse struct {
	EligibilityStatus *string                                 `pulumi:"eligibilityStatus"`
	Messages          []TargetEligibilityErrorMessageResponse `pulumi:"messages"`
}





type TargetEligibilityResultResponseInput interface {
	pulumi.Input

	ToTargetEligibilityResultResponseOutput() TargetEligibilityResultResponseOutput
	ToTargetEligibilityResultResponseOutputWithContext(context.Context) TargetEligibilityResultResponseOutput
}

type TargetEligibilityResultResponseArgs struct {
	EligibilityStatus pulumi.StringPtrInput                           `pulumi:"eligibilityStatus"`
	Messages          TargetEligibilityErrorMessageResponseArrayInput `pulumi:"messages"`
}

func (TargetEligibilityResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetEligibilityResultResponse)(nil)).Elem()
}

func (i TargetEligibilityResultResponseArgs) ToTargetEligibilityResultResponseOutput() TargetEligibilityResultResponseOutput {
	return i.ToTargetEligibilityResultResponseOutputWithContext(context.Background())
}

func (i TargetEligibilityResultResponseArgs) ToTargetEligibilityResultResponseOutputWithContext(ctx context.Context) TargetEligibilityResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetEligibilityResultResponseOutput)
}

func (i TargetEligibilityResultResponseArgs) ToTargetEligibilityResultResponsePtrOutput() TargetEligibilityResultResponsePtrOutput {
	return i.ToTargetEligibilityResultResponsePtrOutputWithContext(context.Background())
}

func (i TargetEligibilityResultResponseArgs) ToTargetEligibilityResultResponsePtrOutputWithContext(ctx context.Context) TargetEligibilityResultResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetEligibilityResultResponseOutput).ToTargetEligibilityResultResponsePtrOutputWithContext(ctx)
}









type TargetEligibilityResultResponsePtrInput interface {
	pulumi.Input

	ToTargetEligibilityResultResponsePtrOutput() TargetEligibilityResultResponsePtrOutput
	ToTargetEligibilityResultResponsePtrOutputWithContext(context.Context) TargetEligibilityResultResponsePtrOutput
}

type targetEligibilityResultResponsePtrType TargetEligibilityResultResponseArgs

func TargetEligibilityResultResponsePtr(v *TargetEligibilityResultResponseArgs) TargetEligibilityResultResponsePtrInput {
	return (*targetEligibilityResultResponsePtrType)(v)
}

func (*targetEligibilityResultResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetEligibilityResultResponse)(nil)).Elem()
}

func (i *targetEligibilityResultResponsePtrType) ToTargetEligibilityResultResponsePtrOutput() TargetEligibilityResultResponsePtrOutput {
	return i.ToTargetEligibilityResultResponsePtrOutputWithContext(context.Background())
}

func (i *targetEligibilityResultResponsePtrType) ToTargetEligibilityResultResponsePtrOutputWithContext(ctx context.Context) TargetEligibilityResultResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetEligibilityResultResponsePtrOutput)
}

type TargetEligibilityResultResponseOutput struct{ *pulumi.OutputState }

func (TargetEligibilityResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetEligibilityResultResponse)(nil)).Elem()
}

func (o TargetEligibilityResultResponseOutput) ToTargetEligibilityResultResponseOutput() TargetEligibilityResultResponseOutput {
	return o
}

func (o TargetEligibilityResultResponseOutput) ToTargetEligibilityResultResponseOutputWithContext(ctx context.Context) TargetEligibilityResultResponseOutput {
	return o
}

func (o TargetEligibilityResultResponseOutput) ToTargetEligibilityResultResponsePtrOutput() TargetEligibilityResultResponsePtrOutput {
	return o.ToTargetEligibilityResultResponsePtrOutputWithContext(context.Background())
}

func (o TargetEligibilityResultResponseOutput) ToTargetEligibilityResultResponsePtrOutputWithContext(ctx context.Context) TargetEligibilityResultResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TargetEligibilityResultResponse) *TargetEligibilityResultResponse {
		return &v
	}).(TargetEligibilityResultResponsePtrOutput)
}

func (o TargetEligibilityResultResponseOutput) EligibilityStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetEligibilityResultResponse) *string { return v.EligibilityStatus }).(pulumi.StringPtrOutput)
}

func (o TargetEligibilityResultResponseOutput) Messages() TargetEligibilityErrorMessageResponseArrayOutput {
	return o.ApplyT(func(v TargetEligibilityResultResponse) []TargetEligibilityErrorMessageResponse { return v.Messages }).(TargetEligibilityErrorMessageResponseArrayOutput)
}

type TargetEligibilityResultResponsePtrOutput struct{ *pulumi.OutputState }

func (TargetEligibilityResultResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetEligibilityResultResponse)(nil)).Elem()
}

func (o TargetEligibilityResultResponsePtrOutput) ToTargetEligibilityResultResponsePtrOutput() TargetEligibilityResultResponsePtrOutput {
	return o
}

func (o TargetEligibilityResultResponsePtrOutput) ToTargetEligibilityResultResponsePtrOutputWithContext(ctx context.Context) TargetEligibilityResultResponsePtrOutput {
	return o
}

func (o TargetEligibilityResultResponsePtrOutput) Elem() TargetEligibilityResultResponseOutput {
	return o.ApplyT(func(v *TargetEligibilityResultResponse) TargetEligibilityResultResponse {
		if v != nil {
			return *v
		}
		var ret TargetEligibilityResultResponse
		return ret
	}).(TargetEligibilityResultResponseOutput)
}

func (o TargetEligibilityResultResponsePtrOutput) EligibilityStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetEligibilityResultResponse) *string {
		if v == nil {
			return nil
		}
		return v.EligibilityStatus
	}).(pulumi.StringPtrOutput)
}

func (o TargetEligibilityResultResponsePtrOutput) Messages() TargetEligibilityErrorMessageResponseArrayOutput {
	return o.ApplyT(func(v *TargetEligibilityResultResponse) []TargetEligibilityErrorMessageResponse {
		if v == nil {
			return nil
		}
		return v.Messages
	}).(TargetEligibilityErrorMessageResponseArrayOutput)
}

type Time struct {
	Hours   int `pulumi:"hours"`
	Minutes int `pulumi:"minutes"`
	Seconds int `pulumi:"seconds"`
}





type TimeInput interface {
	pulumi.Input

	ToTimeOutput() TimeOutput
	ToTimeOutputWithContext(context.Context) TimeOutput
}

type TimeArgs struct {
	Hours   pulumi.IntInput `pulumi:"hours"`
	Minutes pulumi.IntInput `pulumi:"minutes"`
	Seconds pulumi.IntInput `pulumi:"seconds"`
}

func (TimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Time)(nil)).Elem()
}

func (i TimeArgs) ToTimeOutput() TimeOutput {
	return i.ToTimeOutputWithContext(context.Background())
}

func (i TimeArgs) ToTimeOutputWithContext(ctx context.Context) TimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeOutput)
}

type TimeOutput struct{ *pulumi.OutputState }

func (TimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Time)(nil)).Elem()
}

func (o TimeOutput) ToTimeOutput() TimeOutput {
	return o
}

func (o TimeOutput) ToTimeOutputWithContext(ctx context.Context) TimeOutput {
	return o
}

func (o TimeOutput) Hours() pulumi.IntOutput {
	return o.ApplyT(func(v Time) int { return v.Hours }).(pulumi.IntOutput)
}

func (o TimeOutput) Minutes() pulumi.IntOutput {
	return o.ApplyT(func(v Time) int { return v.Minutes }).(pulumi.IntOutput)
}

func (o TimeOutput) Seconds() pulumi.IntOutput {
	return o.ApplyT(func(v Time) int { return v.Seconds }).(pulumi.IntOutput)
}

type TimeResponse struct {
	Hours   int `pulumi:"hours"`
	Minutes int `pulumi:"minutes"`
	Seconds int `pulumi:"seconds"`
}





type TimeResponseInput interface {
	pulumi.Input

	ToTimeResponseOutput() TimeResponseOutput
	ToTimeResponseOutputWithContext(context.Context) TimeResponseOutput
}

type TimeResponseArgs struct {
	Hours   pulumi.IntInput `pulumi:"hours"`
	Minutes pulumi.IntInput `pulumi:"minutes"`
	Seconds pulumi.IntInput `pulumi:"seconds"`
}

func (TimeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeResponse)(nil)).Elem()
}

func (i TimeResponseArgs) ToTimeResponseOutput() TimeResponseOutput {
	return i.ToTimeResponseOutputWithContext(context.Background())
}

func (i TimeResponseArgs) ToTimeResponseOutputWithContext(ctx context.Context) TimeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimeResponseOutput)
}

type TimeResponseOutput struct{ *pulumi.OutputState }

func (TimeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeResponse)(nil)).Elem()
}

func (o TimeResponseOutput) ToTimeResponseOutput() TimeResponseOutput {
	return o
}

func (o TimeResponseOutput) ToTimeResponseOutputWithContext(ctx context.Context) TimeResponseOutput {
	return o
}

func (o TimeResponseOutput) Hours() pulumi.IntOutput {
	return o.ApplyT(func(v TimeResponse) int { return v.Hours }).(pulumi.IntOutput)
}

func (o TimeResponseOutput) Minutes() pulumi.IntOutput {
	return o.ApplyT(func(v TimeResponse) int { return v.Minutes }).(pulumi.IntOutput)
}

func (o TimeResponseOutput) Seconds() pulumi.IntOutput {
	return o.ApplyT(func(v TimeResponse) int { return v.Seconds }).(pulumi.IntOutput)
}

type VolumeContainerFailoverMetadataResponse struct {
	VolumeContainerId *string                          `pulumi:"volumeContainerId"`
	Volumes           []VolumeFailoverMetadataResponse `pulumi:"volumes"`
}





type VolumeContainerFailoverMetadataResponseInput interface {
	pulumi.Input

	ToVolumeContainerFailoverMetadataResponseOutput() VolumeContainerFailoverMetadataResponseOutput
	ToVolumeContainerFailoverMetadataResponseOutputWithContext(context.Context) VolumeContainerFailoverMetadataResponseOutput
}

type VolumeContainerFailoverMetadataResponseArgs struct {
	VolumeContainerId pulumi.StringPtrInput                    `pulumi:"volumeContainerId"`
	Volumes           VolumeFailoverMetadataResponseArrayInput `pulumi:"volumes"`
}

func (VolumeContainerFailoverMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeContainerFailoverMetadataResponse)(nil)).Elem()
}

func (i VolumeContainerFailoverMetadataResponseArgs) ToVolumeContainerFailoverMetadataResponseOutput() VolumeContainerFailoverMetadataResponseOutput {
	return i.ToVolumeContainerFailoverMetadataResponseOutputWithContext(context.Background())
}

func (i VolumeContainerFailoverMetadataResponseArgs) ToVolumeContainerFailoverMetadataResponseOutputWithContext(ctx context.Context) VolumeContainerFailoverMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeContainerFailoverMetadataResponseOutput)
}





type VolumeContainerFailoverMetadataResponseArrayInput interface {
	pulumi.Input

	ToVolumeContainerFailoverMetadataResponseArrayOutput() VolumeContainerFailoverMetadataResponseArrayOutput
	ToVolumeContainerFailoverMetadataResponseArrayOutputWithContext(context.Context) VolumeContainerFailoverMetadataResponseArrayOutput
}

type VolumeContainerFailoverMetadataResponseArray []VolumeContainerFailoverMetadataResponseInput

func (VolumeContainerFailoverMetadataResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeContainerFailoverMetadataResponse)(nil)).Elem()
}

func (i VolumeContainerFailoverMetadataResponseArray) ToVolumeContainerFailoverMetadataResponseArrayOutput() VolumeContainerFailoverMetadataResponseArrayOutput {
	return i.ToVolumeContainerFailoverMetadataResponseArrayOutputWithContext(context.Background())
}

func (i VolumeContainerFailoverMetadataResponseArray) ToVolumeContainerFailoverMetadataResponseArrayOutputWithContext(ctx context.Context) VolumeContainerFailoverMetadataResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeContainerFailoverMetadataResponseArrayOutput)
}

type VolumeContainerFailoverMetadataResponseOutput struct{ *pulumi.OutputState }

func (VolumeContainerFailoverMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeContainerFailoverMetadataResponse)(nil)).Elem()
}

func (o VolumeContainerFailoverMetadataResponseOutput) ToVolumeContainerFailoverMetadataResponseOutput() VolumeContainerFailoverMetadataResponseOutput {
	return o
}

func (o VolumeContainerFailoverMetadataResponseOutput) ToVolumeContainerFailoverMetadataResponseOutputWithContext(ctx context.Context) VolumeContainerFailoverMetadataResponseOutput {
	return o
}

func (o VolumeContainerFailoverMetadataResponseOutput) VolumeContainerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeContainerFailoverMetadataResponse) *string { return v.VolumeContainerId }).(pulumi.StringPtrOutput)
}

func (o VolumeContainerFailoverMetadataResponseOutput) Volumes() VolumeFailoverMetadataResponseArrayOutput {
	return o.ApplyT(func(v VolumeContainerFailoverMetadataResponse) []VolumeFailoverMetadataResponse { return v.Volumes }).(VolumeFailoverMetadataResponseArrayOutput)
}

type VolumeContainerFailoverMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (VolumeContainerFailoverMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeContainerFailoverMetadataResponse)(nil)).Elem()
}

func (o VolumeContainerFailoverMetadataResponseArrayOutput) ToVolumeContainerFailoverMetadataResponseArrayOutput() VolumeContainerFailoverMetadataResponseArrayOutput {
	return o
}

func (o VolumeContainerFailoverMetadataResponseArrayOutput) ToVolumeContainerFailoverMetadataResponseArrayOutputWithContext(ctx context.Context) VolumeContainerFailoverMetadataResponseArrayOutput {
	return o
}

func (o VolumeContainerFailoverMetadataResponseArrayOutput) Index(i pulumi.IntInput) VolumeContainerFailoverMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeContainerFailoverMetadataResponse {
		return vs[0].([]VolumeContainerFailoverMetadataResponse)[vs[1].(int)]
	}).(VolumeContainerFailoverMetadataResponseOutput)
}

type VolumeFailoverMetadataResponse struct {
	BackupCreatedDate *string  `pulumi:"backupCreatedDate"`
	BackupElementId   *string  `pulumi:"backupElementId"`
	BackupId          *string  `pulumi:"backupId"`
	BackupPolicyId    *string  `pulumi:"backupPolicyId"`
	SizeInBytes       *float64 `pulumi:"sizeInBytes"`
	VolumeId          *string  `pulumi:"volumeId"`
	VolumeType        *string  `pulumi:"volumeType"`
}





type VolumeFailoverMetadataResponseInput interface {
	pulumi.Input

	ToVolumeFailoverMetadataResponseOutput() VolumeFailoverMetadataResponseOutput
	ToVolumeFailoverMetadataResponseOutputWithContext(context.Context) VolumeFailoverMetadataResponseOutput
}

type VolumeFailoverMetadataResponseArgs struct {
	BackupCreatedDate pulumi.StringPtrInput  `pulumi:"backupCreatedDate"`
	BackupElementId   pulumi.StringPtrInput  `pulumi:"backupElementId"`
	BackupId          pulumi.StringPtrInput  `pulumi:"backupId"`
	BackupPolicyId    pulumi.StringPtrInput  `pulumi:"backupPolicyId"`
	SizeInBytes       pulumi.Float64PtrInput `pulumi:"sizeInBytes"`
	VolumeId          pulumi.StringPtrInput  `pulumi:"volumeId"`
	VolumeType        pulumi.StringPtrInput  `pulumi:"volumeType"`
}

func (VolumeFailoverMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeFailoverMetadataResponse)(nil)).Elem()
}

func (i VolumeFailoverMetadataResponseArgs) ToVolumeFailoverMetadataResponseOutput() VolumeFailoverMetadataResponseOutput {
	return i.ToVolumeFailoverMetadataResponseOutputWithContext(context.Background())
}

func (i VolumeFailoverMetadataResponseArgs) ToVolumeFailoverMetadataResponseOutputWithContext(ctx context.Context) VolumeFailoverMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeFailoverMetadataResponseOutput)
}





type VolumeFailoverMetadataResponseArrayInput interface {
	pulumi.Input

	ToVolumeFailoverMetadataResponseArrayOutput() VolumeFailoverMetadataResponseArrayOutput
	ToVolumeFailoverMetadataResponseArrayOutputWithContext(context.Context) VolumeFailoverMetadataResponseArrayOutput
}

type VolumeFailoverMetadataResponseArray []VolumeFailoverMetadataResponseInput

func (VolumeFailoverMetadataResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeFailoverMetadataResponse)(nil)).Elem()
}

func (i VolumeFailoverMetadataResponseArray) ToVolumeFailoverMetadataResponseArrayOutput() VolumeFailoverMetadataResponseArrayOutput {
	return i.ToVolumeFailoverMetadataResponseArrayOutputWithContext(context.Background())
}

func (i VolumeFailoverMetadataResponseArray) ToVolumeFailoverMetadataResponseArrayOutputWithContext(ctx context.Context) VolumeFailoverMetadataResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeFailoverMetadataResponseArrayOutput)
}

type VolumeFailoverMetadataResponseOutput struct{ *pulumi.OutputState }

func (VolumeFailoverMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeFailoverMetadataResponse)(nil)).Elem()
}

func (o VolumeFailoverMetadataResponseOutput) ToVolumeFailoverMetadataResponseOutput() VolumeFailoverMetadataResponseOutput {
	return o
}

func (o VolumeFailoverMetadataResponseOutput) ToVolumeFailoverMetadataResponseOutputWithContext(ctx context.Context) VolumeFailoverMetadataResponseOutput {
	return o
}

func (o VolumeFailoverMetadataResponseOutput) BackupCreatedDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeFailoverMetadataResponse) *string { return v.BackupCreatedDate }).(pulumi.StringPtrOutput)
}

func (o VolumeFailoverMetadataResponseOutput) BackupElementId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeFailoverMetadataResponse) *string { return v.BackupElementId }).(pulumi.StringPtrOutput)
}

func (o VolumeFailoverMetadataResponseOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeFailoverMetadataResponse) *string { return v.BackupId }).(pulumi.StringPtrOutput)
}

func (o VolumeFailoverMetadataResponseOutput) BackupPolicyId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeFailoverMetadataResponse) *string { return v.BackupPolicyId }).(pulumi.StringPtrOutput)
}

func (o VolumeFailoverMetadataResponseOutput) SizeInBytes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VolumeFailoverMetadataResponse) *float64 { return v.SizeInBytes }).(pulumi.Float64PtrOutput)
}

func (o VolumeFailoverMetadataResponseOutput) VolumeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeFailoverMetadataResponse) *string { return v.VolumeId }).(pulumi.StringPtrOutput)
}

func (o VolumeFailoverMetadataResponseOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeFailoverMetadataResponse) *string { return v.VolumeType }).(pulumi.StringPtrOutput)
}

type VolumeFailoverMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (VolumeFailoverMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumeFailoverMetadataResponse)(nil)).Elem()
}

func (o VolumeFailoverMetadataResponseArrayOutput) ToVolumeFailoverMetadataResponseArrayOutput() VolumeFailoverMetadataResponseArrayOutput {
	return o
}

func (o VolumeFailoverMetadataResponseArrayOutput) ToVolumeFailoverMetadataResponseArrayOutputWithContext(ctx context.Context) VolumeFailoverMetadataResponseArrayOutput {
	return o
}

func (o VolumeFailoverMetadataResponseArrayOutput) Index(i pulumi.IntInput) VolumeFailoverMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumeFailoverMetadataResponse {
		return vs[0].([]VolumeFailoverMetadataResponse)[vs[1].(int)]
	}).(VolumeFailoverMetadataResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretOutput{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretPtrOutput{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretResponseOutput{})
	pulumi.RegisterOutputType(AsymmetricEncryptedSecretResponsePtrOutput{})
	pulumi.RegisterOutputType(BandwidthScheduleOutput{})
	pulumi.RegisterOutputType(BandwidthScheduleArrayOutput{})
	pulumi.RegisterOutputType(BandwidthScheduleResponseOutput{})
	pulumi.RegisterOutputType(BandwidthScheduleResponseArrayOutput{})
	pulumi.RegisterOutputType(FailoverSetEligibilityResultResponseOutput{})
	pulumi.RegisterOutputType(FailoverSetEligibilityResultResponsePtrOutput{})
	pulumi.RegisterOutputType(FailoverSetResponseOutput{})
	pulumi.RegisterOutputType(FailoverSetResponseArrayOutput{})
	pulumi.RegisterOutputType(FailoverTargetResponseOutput{})
	pulumi.RegisterOutputType(FailoverTargetResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagerIntrinsicSettingsOutput{})
	pulumi.RegisterOutputType(ManagerIntrinsicSettingsPtrOutput{})
	pulumi.RegisterOutputType(ManagerIntrinsicSettingsResponseOutput{})
	pulumi.RegisterOutputType(ManagerIntrinsicSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagerSkuOutput{})
	pulumi.RegisterOutputType(ManagerSkuPtrOutput{})
	pulumi.RegisterOutputType(ManagerSkuResponseOutput{})
	pulumi.RegisterOutputType(ManagerSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(ScheduleRecurrenceOutput{})
	pulumi.RegisterOutputType(ScheduleRecurrencePtrOutput{})
	pulumi.RegisterOutputType(ScheduleRecurrenceResponseOutput{})
	pulumi.RegisterOutputType(ScheduleRecurrenceResponsePtrOutput{})
	pulumi.RegisterOutputType(TargetEligibilityErrorMessageResponseOutput{})
	pulumi.RegisterOutputType(TargetEligibilityErrorMessageResponseArrayOutput{})
	pulumi.RegisterOutputType(TargetEligibilityResultResponseOutput{})
	pulumi.RegisterOutputType(TargetEligibilityResultResponsePtrOutput{})
	pulumi.RegisterOutputType(TimeOutput{})
	pulumi.RegisterOutputType(TimeResponseOutput{})
	pulumi.RegisterOutputType(VolumeContainerFailoverMetadataResponseOutput{})
	pulumi.RegisterOutputType(VolumeContainerFailoverMetadataResponseArrayOutput{})
	pulumi.RegisterOutputType(VolumeFailoverMetadataResponseOutput{})
	pulumi.RegisterOutputType(VolumeFailoverMetadataResponseArrayOutput{})
}
