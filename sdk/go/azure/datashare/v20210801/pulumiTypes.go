


package v20210801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Identity struct {
	Type *string `pulumi:"type"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Identity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type IdentityResponseInput interface {
	pulumi.Input

	ToIdentityResponseOutput() IdentityResponseOutput
	ToIdentityResponseOutputWithContext(context.Context) IdentityResponseOutput
}

type IdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (IdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (i IdentityResponseArgs) ToIdentityResponseOutput() IdentityResponseOutput {
	return i.ToIdentityResponseOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput)
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i IdentityResponseArgs) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponseOutput).ToIdentityResponsePtrOutputWithContext(ctx)
}









type IdentityResponsePtrInput interface {
	pulumi.Input

	ToIdentityResponsePtrOutput() IdentityResponsePtrOutput
	ToIdentityResponsePtrOutputWithContext(context.Context) IdentityResponsePtrOutput
}

type identityResponsePtrType IdentityResponseArgs

func IdentityResponsePtr(v *IdentityResponseArgs) IdentityResponsePtrInput {
	return (*identityResponsePtrType)(v)
}

func (*identityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return i.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *identityResponsePtrType) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityResponsePtrOutput)
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o.ToIdentityResponsePtrOutputWithContext(context.Background())
}

func (o IdentityResponseOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityResponse) *IdentityResponse {
		return &v
	}).(IdentityResponsePtrOutput)
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ScheduledSourceSynchronizationSettingResponse struct {
	Kind                string  `pulumi:"kind"`
	RecurrenceInterval  *string `pulumi:"recurrenceInterval"`
	SynchronizationTime *string `pulumi:"synchronizationTime"`
}





type ScheduledSourceSynchronizationSettingResponseInput interface {
	pulumi.Input

	ToScheduledSourceSynchronizationSettingResponseOutput() ScheduledSourceSynchronizationSettingResponseOutput
	ToScheduledSourceSynchronizationSettingResponseOutputWithContext(context.Context) ScheduledSourceSynchronizationSettingResponseOutput
}

type ScheduledSourceSynchronizationSettingResponseArgs struct {
	Kind                pulumi.StringInput    `pulumi:"kind"`
	RecurrenceInterval  pulumi.StringPtrInput `pulumi:"recurrenceInterval"`
	SynchronizationTime pulumi.StringPtrInput `pulumi:"synchronizationTime"`
}

func (ScheduledSourceSynchronizationSettingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledSourceSynchronizationSettingResponse)(nil)).Elem()
}

func (i ScheduledSourceSynchronizationSettingResponseArgs) ToScheduledSourceSynchronizationSettingResponseOutput() ScheduledSourceSynchronizationSettingResponseOutput {
	return i.ToScheduledSourceSynchronizationSettingResponseOutputWithContext(context.Background())
}

func (i ScheduledSourceSynchronizationSettingResponseArgs) ToScheduledSourceSynchronizationSettingResponseOutputWithContext(ctx context.Context) ScheduledSourceSynchronizationSettingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledSourceSynchronizationSettingResponseOutput)
}





type ScheduledSourceSynchronizationSettingResponseArrayInput interface {
	pulumi.Input

	ToScheduledSourceSynchronizationSettingResponseArrayOutput() ScheduledSourceSynchronizationSettingResponseArrayOutput
	ToScheduledSourceSynchronizationSettingResponseArrayOutputWithContext(context.Context) ScheduledSourceSynchronizationSettingResponseArrayOutput
}

type ScheduledSourceSynchronizationSettingResponseArray []ScheduledSourceSynchronizationSettingResponseInput

func (ScheduledSourceSynchronizationSettingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduledSourceSynchronizationSettingResponse)(nil)).Elem()
}

func (i ScheduledSourceSynchronizationSettingResponseArray) ToScheduledSourceSynchronizationSettingResponseArrayOutput() ScheduledSourceSynchronizationSettingResponseArrayOutput {
	return i.ToScheduledSourceSynchronizationSettingResponseArrayOutputWithContext(context.Background())
}

func (i ScheduledSourceSynchronizationSettingResponseArray) ToScheduledSourceSynchronizationSettingResponseArrayOutputWithContext(ctx context.Context) ScheduledSourceSynchronizationSettingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduledSourceSynchronizationSettingResponseArrayOutput)
}

type ScheduledSourceSynchronizationSettingResponseOutput struct{ *pulumi.OutputState }

func (ScheduledSourceSynchronizationSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduledSourceSynchronizationSettingResponse)(nil)).Elem()
}

func (o ScheduledSourceSynchronizationSettingResponseOutput) ToScheduledSourceSynchronizationSettingResponseOutput() ScheduledSourceSynchronizationSettingResponseOutput {
	return o
}

func (o ScheduledSourceSynchronizationSettingResponseOutput) ToScheduledSourceSynchronizationSettingResponseOutputWithContext(ctx context.Context) ScheduledSourceSynchronizationSettingResponseOutput {
	return o
}

func (o ScheduledSourceSynchronizationSettingResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduledSourceSynchronizationSettingResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o ScheduledSourceSynchronizationSettingResponseOutput) RecurrenceInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduledSourceSynchronizationSettingResponse) *string { return v.RecurrenceInterval }).(pulumi.StringPtrOutput)
}

func (o ScheduledSourceSynchronizationSettingResponseOutput) SynchronizationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduledSourceSynchronizationSettingResponse) *string { return v.SynchronizationTime }).(pulumi.StringPtrOutput)
}

type ScheduledSourceSynchronizationSettingResponseArrayOutput struct{ *pulumi.OutputState }

func (ScheduledSourceSynchronizationSettingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScheduledSourceSynchronizationSettingResponse)(nil)).Elem()
}

func (o ScheduledSourceSynchronizationSettingResponseArrayOutput) ToScheduledSourceSynchronizationSettingResponseArrayOutput() ScheduledSourceSynchronizationSettingResponseArrayOutput {
	return o
}

func (o ScheduledSourceSynchronizationSettingResponseArrayOutput) ToScheduledSourceSynchronizationSettingResponseArrayOutputWithContext(ctx context.Context) ScheduledSourceSynchronizationSettingResponseArrayOutput {
	return o
}

func (o ScheduledSourceSynchronizationSettingResponseArrayOutput) Index(i pulumi.IntInput) ScheduledSourceSynchronizationSettingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScheduledSourceSynchronizationSettingResponse {
		return vs[0].([]ScheduledSourceSynchronizationSettingResponse)[vs[1].(int)]
	}).(ScheduledSourceSynchronizationSettingResponseOutput)
}

type ShareSubscriptionSynchronizationResponse struct {
	DurationMs          int    `pulumi:"durationMs"`
	EndTime             string `pulumi:"endTime"`
	Message             string `pulumi:"message"`
	StartTime           string `pulumi:"startTime"`
	Status              string `pulumi:"status"`
	SynchronizationId   string `pulumi:"synchronizationId"`
	SynchronizationMode string `pulumi:"synchronizationMode"`
}





type ShareSubscriptionSynchronizationResponseInput interface {
	pulumi.Input

	ToShareSubscriptionSynchronizationResponseOutput() ShareSubscriptionSynchronizationResponseOutput
	ToShareSubscriptionSynchronizationResponseOutputWithContext(context.Context) ShareSubscriptionSynchronizationResponseOutput
}

type ShareSubscriptionSynchronizationResponseArgs struct {
	DurationMs          pulumi.IntInput    `pulumi:"durationMs"`
	EndTime             pulumi.StringInput `pulumi:"endTime"`
	Message             pulumi.StringInput `pulumi:"message"`
	StartTime           pulumi.StringInput `pulumi:"startTime"`
	Status              pulumi.StringInput `pulumi:"status"`
	SynchronizationId   pulumi.StringInput `pulumi:"synchronizationId"`
	SynchronizationMode pulumi.StringInput `pulumi:"synchronizationMode"`
}

func (ShareSubscriptionSynchronizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareSubscriptionSynchronizationResponse)(nil)).Elem()
}

func (i ShareSubscriptionSynchronizationResponseArgs) ToShareSubscriptionSynchronizationResponseOutput() ShareSubscriptionSynchronizationResponseOutput {
	return i.ToShareSubscriptionSynchronizationResponseOutputWithContext(context.Background())
}

func (i ShareSubscriptionSynchronizationResponseArgs) ToShareSubscriptionSynchronizationResponseOutputWithContext(ctx context.Context) ShareSubscriptionSynchronizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareSubscriptionSynchronizationResponseOutput)
}





type ShareSubscriptionSynchronizationResponseArrayInput interface {
	pulumi.Input

	ToShareSubscriptionSynchronizationResponseArrayOutput() ShareSubscriptionSynchronizationResponseArrayOutput
	ToShareSubscriptionSynchronizationResponseArrayOutputWithContext(context.Context) ShareSubscriptionSynchronizationResponseArrayOutput
}

type ShareSubscriptionSynchronizationResponseArray []ShareSubscriptionSynchronizationResponseInput

func (ShareSubscriptionSynchronizationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareSubscriptionSynchronizationResponse)(nil)).Elem()
}

func (i ShareSubscriptionSynchronizationResponseArray) ToShareSubscriptionSynchronizationResponseArrayOutput() ShareSubscriptionSynchronizationResponseArrayOutput {
	return i.ToShareSubscriptionSynchronizationResponseArrayOutputWithContext(context.Background())
}

func (i ShareSubscriptionSynchronizationResponseArray) ToShareSubscriptionSynchronizationResponseArrayOutputWithContext(ctx context.Context) ShareSubscriptionSynchronizationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareSubscriptionSynchronizationResponseArrayOutput)
}

type ShareSubscriptionSynchronizationResponseOutput struct{ *pulumi.OutputState }

func (ShareSubscriptionSynchronizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareSubscriptionSynchronizationResponse)(nil)).Elem()
}

func (o ShareSubscriptionSynchronizationResponseOutput) ToShareSubscriptionSynchronizationResponseOutput() ShareSubscriptionSynchronizationResponseOutput {
	return o
}

func (o ShareSubscriptionSynchronizationResponseOutput) ToShareSubscriptionSynchronizationResponseOutputWithContext(ctx context.Context) ShareSubscriptionSynchronizationResponseOutput {
	return o
}

func (o ShareSubscriptionSynchronizationResponseOutput) DurationMs() pulumi.IntOutput {
	return o.ApplyT(func(v ShareSubscriptionSynchronizationResponse) int { return v.DurationMs }).(pulumi.IntOutput)
}

func (o ShareSubscriptionSynchronizationResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v ShareSubscriptionSynchronizationResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o ShareSubscriptionSynchronizationResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ShareSubscriptionSynchronizationResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o ShareSubscriptionSynchronizationResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v ShareSubscriptionSynchronizationResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o ShareSubscriptionSynchronizationResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ShareSubscriptionSynchronizationResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o ShareSubscriptionSynchronizationResponseOutput) SynchronizationId() pulumi.StringOutput {
	return o.ApplyT(func(v ShareSubscriptionSynchronizationResponse) string { return v.SynchronizationId }).(pulumi.StringOutput)
}

func (o ShareSubscriptionSynchronizationResponseOutput) SynchronizationMode() pulumi.StringOutput {
	return o.ApplyT(func(v ShareSubscriptionSynchronizationResponse) string { return v.SynchronizationMode }).(pulumi.StringOutput)
}

type ShareSubscriptionSynchronizationResponseArrayOutput struct{ *pulumi.OutputState }

func (ShareSubscriptionSynchronizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareSubscriptionSynchronizationResponse)(nil)).Elem()
}

func (o ShareSubscriptionSynchronizationResponseArrayOutput) ToShareSubscriptionSynchronizationResponseArrayOutput() ShareSubscriptionSynchronizationResponseArrayOutput {
	return o
}

func (o ShareSubscriptionSynchronizationResponseArrayOutput) ToShareSubscriptionSynchronizationResponseArrayOutputWithContext(ctx context.Context) ShareSubscriptionSynchronizationResponseArrayOutput {
	return o
}

func (o ShareSubscriptionSynchronizationResponseArrayOutput) Index(i pulumi.IntInput) ShareSubscriptionSynchronizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ShareSubscriptionSynchronizationResponse {
		return vs[0].([]ShareSubscriptionSynchronizationResponse)[vs[1].(int)]
	}).(ShareSubscriptionSynchronizationResponseOutput)
}

type ShareSynchronizationResponse struct {
	ConsumerEmail       *string `pulumi:"consumerEmail"`
	ConsumerName        *string `pulumi:"consumerName"`
	ConsumerTenantName  *string `pulumi:"consumerTenantName"`
	DurationMs          *int    `pulumi:"durationMs"`
	EndTime             *string `pulumi:"endTime"`
	Message             *string `pulumi:"message"`
	StartTime           *string `pulumi:"startTime"`
	Status              *string `pulumi:"status"`
	SynchronizationId   *string `pulumi:"synchronizationId"`
	SynchronizationMode string  `pulumi:"synchronizationMode"`
}





type ShareSynchronizationResponseInput interface {
	pulumi.Input

	ToShareSynchronizationResponseOutput() ShareSynchronizationResponseOutput
	ToShareSynchronizationResponseOutputWithContext(context.Context) ShareSynchronizationResponseOutput
}

type ShareSynchronizationResponseArgs struct {
	ConsumerEmail       pulumi.StringPtrInput `pulumi:"consumerEmail"`
	ConsumerName        pulumi.StringPtrInput `pulumi:"consumerName"`
	ConsumerTenantName  pulumi.StringPtrInput `pulumi:"consumerTenantName"`
	DurationMs          pulumi.IntPtrInput    `pulumi:"durationMs"`
	EndTime             pulumi.StringPtrInput `pulumi:"endTime"`
	Message             pulumi.StringPtrInput `pulumi:"message"`
	StartTime           pulumi.StringPtrInput `pulumi:"startTime"`
	Status              pulumi.StringPtrInput `pulumi:"status"`
	SynchronizationId   pulumi.StringPtrInput `pulumi:"synchronizationId"`
	SynchronizationMode pulumi.StringInput    `pulumi:"synchronizationMode"`
}

func (ShareSynchronizationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareSynchronizationResponse)(nil)).Elem()
}

func (i ShareSynchronizationResponseArgs) ToShareSynchronizationResponseOutput() ShareSynchronizationResponseOutput {
	return i.ToShareSynchronizationResponseOutputWithContext(context.Background())
}

func (i ShareSynchronizationResponseArgs) ToShareSynchronizationResponseOutputWithContext(ctx context.Context) ShareSynchronizationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareSynchronizationResponseOutput)
}





type ShareSynchronizationResponseArrayInput interface {
	pulumi.Input

	ToShareSynchronizationResponseArrayOutput() ShareSynchronizationResponseArrayOutput
	ToShareSynchronizationResponseArrayOutputWithContext(context.Context) ShareSynchronizationResponseArrayOutput
}

type ShareSynchronizationResponseArray []ShareSynchronizationResponseInput

func (ShareSynchronizationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareSynchronizationResponse)(nil)).Elem()
}

func (i ShareSynchronizationResponseArray) ToShareSynchronizationResponseArrayOutput() ShareSynchronizationResponseArrayOutput {
	return i.ToShareSynchronizationResponseArrayOutputWithContext(context.Background())
}

func (i ShareSynchronizationResponseArray) ToShareSynchronizationResponseArrayOutputWithContext(ctx context.Context) ShareSynchronizationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ShareSynchronizationResponseArrayOutput)
}

type ShareSynchronizationResponseOutput struct{ *pulumi.OutputState }

func (ShareSynchronizationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ShareSynchronizationResponse)(nil)).Elem()
}

func (o ShareSynchronizationResponseOutput) ToShareSynchronizationResponseOutput() ShareSynchronizationResponseOutput {
	return o
}

func (o ShareSynchronizationResponseOutput) ToShareSynchronizationResponseOutputWithContext(ctx context.Context) ShareSynchronizationResponseOutput {
	return o
}

func (o ShareSynchronizationResponseOutput) ConsumerEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShareSynchronizationResponse) *string { return v.ConsumerEmail }).(pulumi.StringPtrOutput)
}

func (o ShareSynchronizationResponseOutput) ConsumerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShareSynchronizationResponse) *string { return v.ConsumerName }).(pulumi.StringPtrOutput)
}

func (o ShareSynchronizationResponseOutput) ConsumerTenantName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShareSynchronizationResponse) *string { return v.ConsumerTenantName }).(pulumi.StringPtrOutput)
}

func (o ShareSynchronizationResponseOutput) DurationMs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ShareSynchronizationResponse) *int { return v.DurationMs }).(pulumi.IntPtrOutput)
}

func (o ShareSynchronizationResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShareSynchronizationResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o ShareSynchronizationResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShareSynchronizationResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ShareSynchronizationResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShareSynchronizationResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o ShareSynchronizationResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShareSynchronizationResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ShareSynchronizationResponseOutput) SynchronizationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ShareSynchronizationResponse) *string { return v.SynchronizationId }).(pulumi.StringPtrOutput)
}

func (o ShareSynchronizationResponseOutput) SynchronizationMode() pulumi.StringOutput {
	return o.ApplyT(func(v ShareSynchronizationResponse) string { return v.SynchronizationMode }).(pulumi.StringOutput)
}

type ShareSynchronizationResponseArrayOutput struct{ *pulumi.OutputState }

func (ShareSynchronizationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ShareSynchronizationResponse)(nil)).Elem()
}

func (o ShareSynchronizationResponseArrayOutput) ToShareSynchronizationResponseArrayOutput() ShareSynchronizationResponseArrayOutput {
	return o
}

func (o ShareSynchronizationResponseArrayOutput) ToShareSynchronizationResponseArrayOutputWithContext(ctx context.Context) ShareSynchronizationResponseArrayOutput {
	return o
}

func (o ShareSynchronizationResponseArrayOutput) Index(i pulumi.IntInput) ShareSynchronizationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ShareSynchronizationResponse {
		return vs[0].([]ShareSynchronizationResponse)[vs[1].(int)]
	}).(ShareSynchronizationResponseOutput)
}

type SynchronizationDetailsResponse struct {
	DataSetId    string  `pulumi:"dataSetId"`
	DataSetType  string  `pulumi:"dataSetType"`
	DurationMs   int     `pulumi:"durationMs"`
	EndTime      string  `pulumi:"endTime"`
	FilesRead    float64 `pulumi:"filesRead"`
	FilesWritten float64 `pulumi:"filesWritten"`
	Message      string  `pulumi:"message"`
	Name         string  `pulumi:"name"`
	RowsCopied   float64 `pulumi:"rowsCopied"`
	RowsRead     float64 `pulumi:"rowsRead"`
	SizeRead     float64 `pulumi:"sizeRead"`
	SizeWritten  float64 `pulumi:"sizeWritten"`
	StartTime    string  `pulumi:"startTime"`
	Status       string  `pulumi:"status"`
	VCore        float64 `pulumi:"vCore"`
}





type SynchronizationDetailsResponseInput interface {
	pulumi.Input

	ToSynchronizationDetailsResponseOutput() SynchronizationDetailsResponseOutput
	ToSynchronizationDetailsResponseOutputWithContext(context.Context) SynchronizationDetailsResponseOutput
}

type SynchronizationDetailsResponseArgs struct {
	DataSetId    pulumi.StringInput  `pulumi:"dataSetId"`
	DataSetType  pulumi.StringInput  `pulumi:"dataSetType"`
	DurationMs   pulumi.IntInput     `pulumi:"durationMs"`
	EndTime      pulumi.StringInput  `pulumi:"endTime"`
	FilesRead    pulumi.Float64Input `pulumi:"filesRead"`
	FilesWritten pulumi.Float64Input `pulumi:"filesWritten"`
	Message      pulumi.StringInput  `pulumi:"message"`
	Name         pulumi.StringInput  `pulumi:"name"`
	RowsCopied   pulumi.Float64Input `pulumi:"rowsCopied"`
	RowsRead     pulumi.Float64Input `pulumi:"rowsRead"`
	SizeRead     pulumi.Float64Input `pulumi:"sizeRead"`
	SizeWritten  pulumi.Float64Input `pulumi:"sizeWritten"`
	StartTime    pulumi.StringInput  `pulumi:"startTime"`
	Status       pulumi.StringInput  `pulumi:"status"`
	VCore        pulumi.Float64Input `pulumi:"vCore"`
}

func (SynchronizationDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SynchronizationDetailsResponse)(nil)).Elem()
}

func (i SynchronizationDetailsResponseArgs) ToSynchronizationDetailsResponseOutput() SynchronizationDetailsResponseOutput {
	return i.ToSynchronizationDetailsResponseOutputWithContext(context.Background())
}

func (i SynchronizationDetailsResponseArgs) ToSynchronizationDetailsResponseOutputWithContext(ctx context.Context) SynchronizationDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationDetailsResponseOutput)
}





type SynchronizationDetailsResponseArrayInput interface {
	pulumi.Input

	ToSynchronizationDetailsResponseArrayOutput() SynchronizationDetailsResponseArrayOutput
	ToSynchronizationDetailsResponseArrayOutputWithContext(context.Context) SynchronizationDetailsResponseArrayOutput
}

type SynchronizationDetailsResponseArray []SynchronizationDetailsResponseInput

func (SynchronizationDetailsResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SynchronizationDetailsResponse)(nil)).Elem()
}

func (i SynchronizationDetailsResponseArray) ToSynchronizationDetailsResponseArrayOutput() SynchronizationDetailsResponseArrayOutput {
	return i.ToSynchronizationDetailsResponseArrayOutputWithContext(context.Background())
}

func (i SynchronizationDetailsResponseArray) ToSynchronizationDetailsResponseArrayOutputWithContext(ctx context.Context) SynchronizationDetailsResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SynchronizationDetailsResponseArrayOutput)
}

type SynchronizationDetailsResponseOutput struct{ *pulumi.OutputState }

func (SynchronizationDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SynchronizationDetailsResponse)(nil)).Elem()
}

func (o SynchronizationDetailsResponseOutput) ToSynchronizationDetailsResponseOutput() SynchronizationDetailsResponseOutput {
	return o
}

func (o SynchronizationDetailsResponseOutput) ToSynchronizationDetailsResponseOutputWithContext(ctx context.Context) SynchronizationDetailsResponseOutput {
	return o
}

func (o SynchronizationDetailsResponseOutput) DataSetId() pulumi.StringOutput {
	return o.ApplyT(func(v SynchronizationDetailsResponse) string { return v.DataSetId }).(pulumi.StringOutput)
}

func (o SynchronizationDetailsResponseOutput) DataSetType() pulumi.StringOutput {
	return o.ApplyT(func(v SynchronizationDetailsResponse) string { return v.DataSetType }).(pulumi.StringOutput)
}

func (o SynchronizationDetailsResponseOutput) DurationMs() pulumi.IntOutput {
	return o.ApplyT(func(v SynchronizationDetailsResponse) int { return v.DurationMs }).(pulumi.IntOutput)
}

func (o SynchronizationDetailsResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v SynchronizationDetailsResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o SynchronizationDetailsResponseOutput) FilesRead() pulumi.Float64Output {
	return o.ApplyT(func(v SynchronizationDetailsResponse) float64 { return v.FilesRead }).(pulumi.Float64Output)
}

func (o SynchronizationDetailsResponseOutput) FilesWritten() pulumi.Float64Output {
	return o.ApplyT(func(v SynchronizationDetailsResponse) float64 { return v.FilesWritten }).(pulumi.Float64Output)
}

func (o SynchronizationDetailsResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v SynchronizationDetailsResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o SynchronizationDetailsResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SynchronizationDetailsResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SynchronizationDetailsResponseOutput) RowsCopied() pulumi.Float64Output {
	return o.ApplyT(func(v SynchronizationDetailsResponse) float64 { return v.RowsCopied }).(pulumi.Float64Output)
}

func (o SynchronizationDetailsResponseOutput) RowsRead() pulumi.Float64Output {
	return o.ApplyT(func(v SynchronizationDetailsResponse) float64 { return v.RowsRead }).(pulumi.Float64Output)
}

func (o SynchronizationDetailsResponseOutput) SizeRead() pulumi.Float64Output {
	return o.ApplyT(func(v SynchronizationDetailsResponse) float64 { return v.SizeRead }).(pulumi.Float64Output)
}

func (o SynchronizationDetailsResponseOutput) SizeWritten() pulumi.Float64Output {
	return o.ApplyT(func(v SynchronizationDetailsResponse) float64 { return v.SizeWritten }).(pulumi.Float64Output)
}

func (o SynchronizationDetailsResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v SynchronizationDetailsResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o SynchronizationDetailsResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v SynchronizationDetailsResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o SynchronizationDetailsResponseOutput) VCore() pulumi.Float64Output {
	return o.ApplyT(func(v SynchronizationDetailsResponse) float64 { return v.VCore }).(pulumi.Float64Output)
}

type SynchronizationDetailsResponseArrayOutput struct{ *pulumi.OutputState }

func (SynchronizationDetailsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SynchronizationDetailsResponse)(nil)).Elem()
}

func (o SynchronizationDetailsResponseArrayOutput) ToSynchronizationDetailsResponseArrayOutput() SynchronizationDetailsResponseArrayOutput {
	return o
}

func (o SynchronizationDetailsResponseArrayOutput) ToSynchronizationDetailsResponseArrayOutputWithContext(ctx context.Context) SynchronizationDetailsResponseArrayOutput {
	return o
}

func (o SynchronizationDetailsResponseArrayOutput) Index(i pulumi.IntInput) SynchronizationDetailsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SynchronizationDetailsResponse {
		return vs[0].([]SynchronizationDetailsResponse)[vs[1].(int)]
	}).(SynchronizationDetailsResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type TableLevelSharingProperties struct {
	ExternalTablesToExclude    []string `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    []string `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude []string `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude []string `pulumi:"materializedViewsToInclude"`
	TablesToExclude            []string `pulumi:"tablesToExclude"`
	TablesToInclude            []string `pulumi:"tablesToInclude"`
}





type TableLevelSharingPropertiesInput interface {
	pulumi.Input

	ToTableLevelSharingPropertiesOutput() TableLevelSharingPropertiesOutput
	ToTableLevelSharingPropertiesOutputWithContext(context.Context) TableLevelSharingPropertiesOutput
}

type TableLevelSharingPropertiesArgs struct {
	ExternalTablesToExclude    pulumi.StringArrayInput `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    pulumi.StringArrayInput `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude pulumi.StringArrayInput `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude pulumi.StringArrayInput `pulumi:"materializedViewsToInclude"`
	TablesToExclude            pulumi.StringArrayInput `pulumi:"tablesToExclude"`
	TablesToInclude            pulumi.StringArrayInput `pulumi:"tablesToInclude"`
}

func (TableLevelSharingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableLevelSharingProperties)(nil)).Elem()
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesOutput() TableLevelSharingPropertiesOutput {
	return i.ToTableLevelSharingPropertiesOutputWithContext(context.Background())
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesOutputWithContext(ctx context.Context) TableLevelSharingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesOutput)
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return i.ToTableLevelSharingPropertiesPtrOutputWithContext(context.Background())
}

func (i TableLevelSharingPropertiesArgs) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesOutput).ToTableLevelSharingPropertiesPtrOutputWithContext(ctx)
}









type TableLevelSharingPropertiesPtrInput interface {
	pulumi.Input

	ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput
	ToTableLevelSharingPropertiesPtrOutputWithContext(context.Context) TableLevelSharingPropertiesPtrOutput
}

type tableLevelSharingPropertiesPtrType TableLevelSharingPropertiesArgs

func TableLevelSharingPropertiesPtr(v *TableLevelSharingPropertiesArgs) TableLevelSharingPropertiesPtrInput {
	return (*tableLevelSharingPropertiesPtrType)(v)
}

func (*tableLevelSharingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableLevelSharingProperties)(nil)).Elem()
}

func (i *tableLevelSharingPropertiesPtrType) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return i.ToTableLevelSharingPropertiesPtrOutputWithContext(context.Background())
}

func (i *tableLevelSharingPropertiesPtrType) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesPtrOutput)
}

type TableLevelSharingPropertiesOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableLevelSharingProperties)(nil)).Elem()
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesOutput() TableLevelSharingPropertiesOutput {
	return o
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesOutputWithContext(ctx context.Context) TableLevelSharingPropertiesOutput {
	return o
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return o.ToTableLevelSharingPropertiesPtrOutputWithContext(context.Background())
}

func (o TableLevelSharingPropertiesOutput) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableLevelSharingProperties) *TableLevelSharingProperties {
		return &v
	}).(TableLevelSharingPropertiesPtrOutput)
}

func (o TableLevelSharingPropertiesOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.ExternalTablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.ExternalTablesToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.MaterializedViewsToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.MaterializedViewsToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.TablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingProperties) []string { return v.TablesToInclude }).(pulumi.StringArrayOutput)
}

type TableLevelSharingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableLevelSharingProperties)(nil)).Elem()
}

func (o TableLevelSharingPropertiesPtrOutput) ToTableLevelSharingPropertiesPtrOutput() TableLevelSharingPropertiesPtrOutput {
	return o
}

func (o TableLevelSharingPropertiesPtrOutput) ToTableLevelSharingPropertiesPtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesPtrOutput {
	return o
}

func (o TableLevelSharingPropertiesPtrOutput) Elem() TableLevelSharingPropertiesOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) TableLevelSharingProperties {
		if v != nil {
			return *v
		}
		var ret TableLevelSharingProperties
		return ret
	}).(TableLevelSharingPropertiesOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.TablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesPtrOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingProperties) []string {
		if v == nil {
			return nil
		}
		return v.TablesToInclude
	}).(pulumi.StringArrayOutput)
}

type TableLevelSharingPropertiesResponse struct {
	ExternalTablesToExclude    []string `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    []string `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude []string `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude []string `pulumi:"materializedViewsToInclude"`
	TablesToExclude            []string `pulumi:"tablesToExclude"`
	TablesToInclude            []string `pulumi:"tablesToInclude"`
}





type TableLevelSharingPropertiesResponseInput interface {
	pulumi.Input

	ToTableLevelSharingPropertiesResponseOutput() TableLevelSharingPropertiesResponseOutput
	ToTableLevelSharingPropertiesResponseOutputWithContext(context.Context) TableLevelSharingPropertiesResponseOutput
}

type TableLevelSharingPropertiesResponseArgs struct {
	ExternalTablesToExclude    pulumi.StringArrayInput `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    pulumi.StringArrayInput `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude pulumi.StringArrayInput `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude pulumi.StringArrayInput `pulumi:"materializedViewsToInclude"`
	TablesToExclude            pulumi.StringArrayInput `pulumi:"tablesToExclude"`
	TablesToInclude            pulumi.StringArrayInput `pulumi:"tablesToInclude"`
}

func (TableLevelSharingPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableLevelSharingPropertiesResponse)(nil)).Elem()
}

func (i TableLevelSharingPropertiesResponseArgs) ToTableLevelSharingPropertiesResponseOutput() TableLevelSharingPropertiesResponseOutput {
	return i.ToTableLevelSharingPropertiesResponseOutputWithContext(context.Background())
}

func (i TableLevelSharingPropertiesResponseArgs) ToTableLevelSharingPropertiesResponseOutputWithContext(ctx context.Context) TableLevelSharingPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesResponseOutput)
}

func (i TableLevelSharingPropertiesResponseArgs) ToTableLevelSharingPropertiesResponsePtrOutput() TableLevelSharingPropertiesResponsePtrOutput {
	return i.ToTableLevelSharingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i TableLevelSharingPropertiesResponseArgs) ToTableLevelSharingPropertiesResponsePtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesResponseOutput).ToTableLevelSharingPropertiesResponsePtrOutputWithContext(ctx)
}









type TableLevelSharingPropertiesResponsePtrInput interface {
	pulumi.Input

	ToTableLevelSharingPropertiesResponsePtrOutput() TableLevelSharingPropertiesResponsePtrOutput
	ToTableLevelSharingPropertiesResponsePtrOutputWithContext(context.Context) TableLevelSharingPropertiesResponsePtrOutput
}

type tableLevelSharingPropertiesResponsePtrType TableLevelSharingPropertiesResponseArgs

func TableLevelSharingPropertiesResponsePtr(v *TableLevelSharingPropertiesResponseArgs) TableLevelSharingPropertiesResponsePtrInput {
	return (*tableLevelSharingPropertiesResponsePtrType)(v)
}

func (*tableLevelSharingPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableLevelSharingPropertiesResponse)(nil)).Elem()
}

func (i *tableLevelSharingPropertiesResponsePtrType) ToTableLevelSharingPropertiesResponsePtrOutput() TableLevelSharingPropertiesResponsePtrOutput {
	return i.ToTableLevelSharingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *tableLevelSharingPropertiesResponsePtrType) ToTableLevelSharingPropertiesResponsePtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableLevelSharingPropertiesResponsePtrOutput)
}

type TableLevelSharingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableLevelSharingPropertiesResponse)(nil)).Elem()
}

func (o TableLevelSharingPropertiesResponseOutput) ToTableLevelSharingPropertiesResponseOutput() TableLevelSharingPropertiesResponseOutput {
	return o
}

func (o TableLevelSharingPropertiesResponseOutput) ToTableLevelSharingPropertiesResponseOutputWithContext(ctx context.Context) TableLevelSharingPropertiesResponseOutput {
	return o
}

func (o TableLevelSharingPropertiesResponseOutput) ToTableLevelSharingPropertiesResponsePtrOutput() TableLevelSharingPropertiesResponsePtrOutput {
	return o.ToTableLevelSharingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o TableLevelSharingPropertiesResponseOutput) ToTableLevelSharingPropertiesResponsePtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableLevelSharingPropertiesResponse) *TableLevelSharingPropertiesResponse {
		return &v
	}).(TableLevelSharingPropertiesResponsePtrOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.ExternalTablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.ExternalTablesToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.MaterializedViewsToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.MaterializedViewsToInclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.TablesToExclude }).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponseOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TableLevelSharingPropertiesResponse) []string { return v.TablesToInclude }).(pulumi.StringArrayOutput)
}

type TableLevelSharingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TableLevelSharingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableLevelSharingPropertiesResponse)(nil)).Elem()
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ToTableLevelSharingPropertiesResponsePtrOutput() TableLevelSharingPropertiesResponsePtrOutput {
	return o
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ToTableLevelSharingPropertiesResponsePtrOutputWithContext(ctx context.Context) TableLevelSharingPropertiesResponsePtrOutput {
	return o
}

func (o TableLevelSharingPropertiesResponsePtrOutput) Elem() TableLevelSharingPropertiesResponseOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) TableLevelSharingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TableLevelSharingPropertiesResponse
		return ret
	}).(TableLevelSharingPropertiesResponseOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ExternalTablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) ExternalTablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExternalTablesToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) MaterializedViewsToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) MaterializedViewsToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.MaterializedViewsToInclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) TablesToExclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.TablesToExclude
	}).(pulumi.StringArrayOutput)
}

func (o TableLevelSharingPropertiesResponsePtrOutput) TablesToInclude() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TableLevelSharingPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.TablesToInclude
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityPtrInput)(nil)).Elem(), IdentityArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponseInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityResponsePtrInput)(nil)).Elem(), IdentityResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledSourceSynchronizationSettingResponseInput)(nil)).Elem(), ScheduledSourceSynchronizationSettingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduledSourceSynchronizationSettingResponseArrayInput)(nil)).Elem(), ScheduledSourceSynchronizationSettingResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareSubscriptionSynchronizationResponseInput)(nil)).Elem(), ShareSubscriptionSynchronizationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareSubscriptionSynchronizationResponseArrayInput)(nil)).Elem(), ShareSubscriptionSynchronizationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareSynchronizationResponseInput)(nil)).Elem(), ShareSynchronizationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ShareSynchronizationResponseArrayInput)(nil)).Elem(), ShareSynchronizationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationDetailsResponseInput)(nil)).Elem(), SynchronizationDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SynchronizationDetailsResponseArrayInput)(nil)).Elem(), SynchronizationDetailsResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableLevelSharingPropertiesInput)(nil)).Elem(), TableLevelSharingPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableLevelSharingPropertiesPtrInput)(nil)).Elem(), TableLevelSharingPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableLevelSharingPropertiesResponseInput)(nil)).Elem(), TableLevelSharingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TableLevelSharingPropertiesResponsePtrInput)(nil)).Elem(), TableLevelSharingPropertiesResponseArgs{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ScheduledSourceSynchronizationSettingResponseOutput{})
	pulumi.RegisterOutputType(ScheduledSourceSynchronizationSettingResponseArrayOutput{})
	pulumi.RegisterOutputType(ShareSubscriptionSynchronizationResponseOutput{})
	pulumi.RegisterOutputType(ShareSubscriptionSynchronizationResponseArrayOutput{})
	pulumi.RegisterOutputType(ShareSynchronizationResponseOutput{})
	pulumi.RegisterOutputType(ShareSynchronizationResponseArrayOutput{})
	pulumi.RegisterOutputType(SynchronizationDetailsResponseOutput{})
	pulumi.RegisterOutputType(SynchronizationDetailsResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesResponsePtrOutput{})
}
