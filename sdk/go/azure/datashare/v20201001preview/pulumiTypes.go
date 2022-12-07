


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ADLSGen2StorageAccountPath struct {
	ConsumerPath  *string `pulumi:"consumerPath"`
	ContainerName string  `pulumi:"containerName"`
	ProviderPath  *string `pulumi:"providerPath"`
}





type ADLSGen2StorageAccountPathInput interface {
	pulumi.Input

	ToADLSGen2StorageAccountPathOutput() ADLSGen2StorageAccountPathOutput
	ToADLSGen2StorageAccountPathOutputWithContext(context.Context) ADLSGen2StorageAccountPathOutput
}

type ADLSGen2StorageAccountPathArgs struct {
	ConsumerPath  pulumi.StringPtrInput `pulumi:"consumerPath"`
	ContainerName pulumi.StringInput    `pulumi:"containerName"`
	ProviderPath  pulumi.StringPtrInput `pulumi:"providerPath"`
}

func (ADLSGen2StorageAccountPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ADLSGen2StorageAccountPath)(nil)).Elem()
}

func (i ADLSGen2StorageAccountPathArgs) ToADLSGen2StorageAccountPathOutput() ADLSGen2StorageAccountPathOutput {
	return i.ToADLSGen2StorageAccountPathOutputWithContext(context.Background())
}

func (i ADLSGen2StorageAccountPathArgs) ToADLSGen2StorageAccountPathOutputWithContext(ctx context.Context) ADLSGen2StorageAccountPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2StorageAccountPathOutput)
}





type ADLSGen2StorageAccountPathArrayInput interface {
	pulumi.Input

	ToADLSGen2StorageAccountPathArrayOutput() ADLSGen2StorageAccountPathArrayOutput
	ToADLSGen2StorageAccountPathArrayOutputWithContext(context.Context) ADLSGen2StorageAccountPathArrayOutput
}

type ADLSGen2StorageAccountPathArray []ADLSGen2StorageAccountPathInput

func (ADLSGen2StorageAccountPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ADLSGen2StorageAccountPath)(nil)).Elem()
}

func (i ADLSGen2StorageAccountPathArray) ToADLSGen2StorageAccountPathArrayOutput() ADLSGen2StorageAccountPathArrayOutput {
	return i.ToADLSGen2StorageAccountPathArrayOutputWithContext(context.Background())
}

func (i ADLSGen2StorageAccountPathArray) ToADLSGen2StorageAccountPathArrayOutputWithContext(ctx context.Context) ADLSGen2StorageAccountPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ADLSGen2StorageAccountPathArrayOutput)
}

type ADLSGen2StorageAccountPathOutput struct{ *pulumi.OutputState }

func (ADLSGen2StorageAccountPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ADLSGen2StorageAccountPath)(nil)).Elem()
}

func (o ADLSGen2StorageAccountPathOutput) ToADLSGen2StorageAccountPathOutput() ADLSGen2StorageAccountPathOutput {
	return o
}

func (o ADLSGen2StorageAccountPathOutput) ToADLSGen2StorageAccountPathOutputWithContext(ctx context.Context) ADLSGen2StorageAccountPathOutput {
	return o
}

func (o ADLSGen2StorageAccountPathOutput) ConsumerPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ADLSGen2StorageAccountPath) *string { return v.ConsumerPath }).(pulumi.StringPtrOutput)
}

func (o ADLSGen2StorageAccountPathOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v ADLSGen2StorageAccountPath) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o ADLSGen2StorageAccountPathOutput) ProviderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ADLSGen2StorageAccountPath) *string { return v.ProviderPath }).(pulumi.StringPtrOutput)
}

type ADLSGen2StorageAccountPathArrayOutput struct{ *pulumi.OutputState }

func (ADLSGen2StorageAccountPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ADLSGen2StorageAccountPath)(nil)).Elem()
}

func (o ADLSGen2StorageAccountPathArrayOutput) ToADLSGen2StorageAccountPathArrayOutput() ADLSGen2StorageAccountPathArrayOutput {
	return o
}

func (o ADLSGen2StorageAccountPathArrayOutput) ToADLSGen2StorageAccountPathArrayOutputWithContext(ctx context.Context) ADLSGen2StorageAccountPathArrayOutput {
	return o
}

func (o ADLSGen2StorageAccountPathArrayOutput) Index(i pulumi.IntInput) ADLSGen2StorageAccountPathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ADLSGen2StorageAccountPath {
		return vs[0].([]ADLSGen2StorageAccountPath)[vs[1].(int)]
	}).(ADLSGen2StorageAccountPathOutput)
}

type ADLSGen2StorageAccountPathResponse struct {
	ConsumerPath  *string `pulumi:"consumerPath"`
	ContainerName string  `pulumi:"containerName"`
	ProviderPath  *string `pulumi:"providerPath"`
}

type ADLSGen2StorageAccountPathResponseOutput struct{ *pulumi.OutputState }

func (ADLSGen2StorageAccountPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ADLSGen2StorageAccountPathResponse)(nil)).Elem()
}

func (o ADLSGen2StorageAccountPathResponseOutput) ToADLSGen2StorageAccountPathResponseOutput() ADLSGen2StorageAccountPathResponseOutput {
	return o
}

func (o ADLSGen2StorageAccountPathResponseOutput) ToADLSGen2StorageAccountPathResponseOutputWithContext(ctx context.Context) ADLSGen2StorageAccountPathResponseOutput {
	return o
}

func (o ADLSGen2StorageAccountPathResponseOutput) ConsumerPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ADLSGen2StorageAccountPathResponse) *string { return v.ConsumerPath }).(pulumi.StringPtrOutput)
}

func (o ADLSGen2StorageAccountPathResponseOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v ADLSGen2StorageAccountPathResponse) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o ADLSGen2StorageAccountPathResponseOutput) ProviderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ADLSGen2StorageAccountPathResponse) *string { return v.ProviderPath }).(pulumi.StringPtrOutput)
}

type ADLSGen2StorageAccountPathResponseArrayOutput struct{ *pulumi.OutputState }

func (ADLSGen2StorageAccountPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ADLSGen2StorageAccountPathResponse)(nil)).Elem()
}

func (o ADLSGen2StorageAccountPathResponseArrayOutput) ToADLSGen2StorageAccountPathResponseArrayOutput() ADLSGen2StorageAccountPathResponseArrayOutput {
	return o
}

func (o ADLSGen2StorageAccountPathResponseArrayOutput) ToADLSGen2StorageAccountPathResponseArrayOutputWithContext(ctx context.Context) ADLSGen2StorageAccountPathResponseArrayOutput {
	return o
}

func (o ADLSGen2StorageAccountPathResponseArrayOutput) Index(i pulumi.IntInput) ADLSGen2StorageAccountPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ADLSGen2StorageAccountPathResponse {
		return vs[0].([]ADLSGen2StorageAccountPathResponse)[vs[1].(int)]
	}).(ADLSGen2StorageAccountPathResponseOutput)
}

type BlobStorageAccountPath struct {
	ConsumerPath  *string `pulumi:"consumerPath"`
	ContainerName string  `pulumi:"containerName"`
	ProviderPath  *string `pulumi:"providerPath"`
}





type BlobStorageAccountPathInput interface {
	pulumi.Input

	ToBlobStorageAccountPathOutput() BlobStorageAccountPathOutput
	ToBlobStorageAccountPathOutputWithContext(context.Context) BlobStorageAccountPathOutput
}

type BlobStorageAccountPathArgs struct {
	ConsumerPath  pulumi.StringPtrInput `pulumi:"consumerPath"`
	ContainerName pulumi.StringInput    `pulumi:"containerName"`
	ProviderPath  pulumi.StringPtrInput `pulumi:"providerPath"`
}

func (BlobStorageAccountPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStorageAccountPath)(nil)).Elem()
}

func (i BlobStorageAccountPathArgs) ToBlobStorageAccountPathOutput() BlobStorageAccountPathOutput {
	return i.ToBlobStorageAccountPathOutputWithContext(context.Background())
}

func (i BlobStorageAccountPathArgs) ToBlobStorageAccountPathOutputWithContext(ctx context.Context) BlobStorageAccountPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobStorageAccountPathOutput)
}





type BlobStorageAccountPathArrayInput interface {
	pulumi.Input

	ToBlobStorageAccountPathArrayOutput() BlobStorageAccountPathArrayOutput
	ToBlobStorageAccountPathArrayOutputWithContext(context.Context) BlobStorageAccountPathArrayOutput
}

type BlobStorageAccountPathArray []BlobStorageAccountPathInput

func (BlobStorageAccountPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BlobStorageAccountPath)(nil)).Elem()
}

func (i BlobStorageAccountPathArray) ToBlobStorageAccountPathArrayOutput() BlobStorageAccountPathArrayOutput {
	return i.ToBlobStorageAccountPathArrayOutputWithContext(context.Background())
}

func (i BlobStorageAccountPathArray) ToBlobStorageAccountPathArrayOutputWithContext(ctx context.Context) BlobStorageAccountPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobStorageAccountPathArrayOutput)
}

type BlobStorageAccountPathOutput struct{ *pulumi.OutputState }

func (BlobStorageAccountPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStorageAccountPath)(nil)).Elem()
}

func (o BlobStorageAccountPathOutput) ToBlobStorageAccountPathOutput() BlobStorageAccountPathOutput {
	return o
}

func (o BlobStorageAccountPathOutput) ToBlobStorageAccountPathOutputWithContext(ctx context.Context) BlobStorageAccountPathOutput {
	return o
}

func (o BlobStorageAccountPathOutput) ConsumerPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStorageAccountPath) *string { return v.ConsumerPath }).(pulumi.StringPtrOutput)
}

func (o BlobStorageAccountPathOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v BlobStorageAccountPath) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o BlobStorageAccountPathOutput) ProviderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStorageAccountPath) *string { return v.ProviderPath }).(pulumi.StringPtrOutput)
}

type BlobStorageAccountPathArrayOutput struct{ *pulumi.OutputState }

func (BlobStorageAccountPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BlobStorageAccountPath)(nil)).Elem()
}

func (o BlobStorageAccountPathArrayOutput) ToBlobStorageAccountPathArrayOutput() BlobStorageAccountPathArrayOutput {
	return o
}

func (o BlobStorageAccountPathArrayOutput) ToBlobStorageAccountPathArrayOutputWithContext(ctx context.Context) BlobStorageAccountPathArrayOutput {
	return o
}

func (o BlobStorageAccountPathArrayOutput) Index(i pulumi.IntInput) BlobStorageAccountPathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BlobStorageAccountPath {
		return vs[0].([]BlobStorageAccountPath)[vs[1].(int)]
	}).(BlobStorageAccountPathOutput)
}

type BlobStorageAccountPathResponse struct {
	ConsumerPath  *string `pulumi:"consumerPath"`
	ContainerName string  `pulumi:"containerName"`
	ProviderPath  *string `pulumi:"providerPath"`
}

type BlobStorageAccountPathResponseOutput struct{ *pulumi.OutputState }

func (BlobStorageAccountPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobStorageAccountPathResponse)(nil)).Elem()
}

func (o BlobStorageAccountPathResponseOutput) ToBlobStorageAccountPathResponseOutput() BlobStorageAccountPathResponseOutput {
	return o
}

func (o BlobStorageAccountPathResponseOutput) ToBlobStorageAccountPathResponseOutputWithContext(ctx context.Context) BlobStorageAccountPathResponseOutput {
	return o
}

func (o BlobStorageAccountPathResponseOutput) ConsumerPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStorageAccountPathResponse) *string { return v.ConsumerPath }).(pulumi.StringPtrOutput)
}

func (o BlobStorageAccountPathResponseOutput) ContainerName() pulumi.StringOutput {
	return o.ApplyT(func(v BlobStorageAccountPathResponse) string { return v.ContainerName }).(pulumi.StringOutput)
}

func (o BlobStorageAccountPathResponseOutput) ProviderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BlobStorageAccountPathResponse) *string { return v.ProviderPath }).(pulumi.StringPtrOutput)
}

type BlobStorageAccountPathResponseArrayOutput struct{ *pulumi.OutputState }

func (BlobStorageAccountPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BlobStorageAccountPathResponse)(nil)).Elem()
}

func (o BlobStorageAccountPathResponseArrayOutput) ToBlobStorageAccountPathResponseArrayOutput() BlobStorageAccountPathResponseArrayOutput {
	return o
}

func (o BlobStorageAccountPathResponseArrayOutput) ToBlobStorageAccountPathResponseArrayOutputWithContext(ctx context.Context) BlobStorageAccountPathResponseArrayOutput {
	return o
}

func (o BlobStorageAccountPathResponseArrayOutput) Index(i pulumi.IntInput) BlobStorageAccountPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BlobStorageAccountPathResponse {
		return vs[0].([]BlobStorageAccountPathResponse)[vs[1].(int)]
	}).(BlobStorageAccountPathResponseOutput)
}

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

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
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

func init() {
	pulumi.RegisterOutputType(ADLSGen2StorageAccountPathOutput{})
	pulumi.RegisterOutputType(ADLSGen2StorageAccountPathArrayOutput{})
	pulumi.RegisterOutputType(ADLSGen2StorageAccountPathResponseOutput{})
	pulumi.RegisterOutputType(ADLSGen2StorageAccountPathResponseArrayOutput{})
	pulumi.RegisterOutputType(BlobStorageAccountPathOutput{})
	pulumi.RegisterOutputType(BlobStorageAccountPathArrayOutput{})
	pulumi.RegisterOutputType(BlobStorageAccountPathResponseOutput{})
	pulumi.RegisterOutputType(BlobStorageAccountPathResponseArrayOutput{})
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
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
