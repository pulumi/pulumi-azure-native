


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

type TableLevelSharingPropertiesResponse struct {
	ExternalTablesToExclude    []string `pulumi:"externalTablesToExclude"`
	ExternalTablesToInclude    []string `pulumi:"externalTablesToInclude"`
	MaterializedViewsToExclude []string `pulumi:"materializedViewsToExclude"`
	MaterializedViewsToInclude []string `pulumi:"materializedViewsToInclude"`
	TablesToExclude            []string `pulumi:"tablesToExclude"`
	TablesToInclude            []string `pulumi:"tablesToInclude"`
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

func init() {
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesOutput{})
	pulumi.RegisterOutputType(TableLevelSharingPropertiesResponseOutput{})
}
