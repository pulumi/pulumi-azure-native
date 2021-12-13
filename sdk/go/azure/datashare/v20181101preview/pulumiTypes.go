


package v20181101preview

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

type ShareSubscriptionSynchronizationResponse struct {
	DurationMs          int    `pulumi:"durationMs"`
	EndTime             string `pulumi:"endTime"`
	Message             string `pulumi:"message"`
	StartTime           string `pulumi:"startTime"`
	Status              string `pulumi:"status"`
	SynchronizationId   string `pulumi:"synchronizationId"`
	SynchronizationMode string `pulumi:"synchronizationMode"`
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

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
}
