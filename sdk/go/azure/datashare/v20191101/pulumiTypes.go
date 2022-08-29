


package v20191101

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

func (o IdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Identity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type IdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
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

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ScheduledSourceSynchronizationSettingResponse struct {
	Kind                string  `pulumi:"kind"`
	RecurrenceInterval  *string `pulumi:"recurrenceInterval"`
	SynchronizationTime *string `pulumi:"synchronizationTime"`
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

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(ScheduledSourceSynchronizationSettingResponseOutput{})
	pulumi.RegisterOutputType(ScheduledSourceSynchronizationSettingResponseArrayOutput{})
	pulumi.RegisterOutputType(ShareSubscriptionSynchronizationResponseOutput{})
	pulumi.RegisterOutputType(ShareSubscriptionSynchronizationResponseArrayOutput{})
	pulumi.RegisterOutputType(ShareSynchronizationResponseOutput{})
	pulumi.RegisterOutputType(ShareSynchronizationResponseArrayOutput{})
	pulumi.RegisterOutputType(SynchronizationDetailsResponseOutput{})
	pulumi.RegisterOutputType(SynchronizationDetailsResponseArrayOutput{})
}
