


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdditionalWorkspaceDataType string

const (
	AdditionalWorkspaceDataTypeAlerts    = AdditionalWorkspaceDataType("Alerts")
	AdditionalWorkspaceDataTypeRawEvents = AdditionalWorkspaceDataType("RawEvents")
)

func (AdditionalWorkspaceDataType) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspaceDataType)(nil)).Elem()
}

func (e AdditionalWorkspaceDataType) ToAdditionalWorkspaceDataTypeOutput() AdditionalWorkspaceDataTypeOutput {
	return pulumi.ToOutput(e).(AdditionalWorkspaceDataTypeOutput)
}

func (e AdditionalWorkspaceDataType) ToAdditionalWorkspaceDataTypeOutputWithContext(ctx context.Context) AdditionalWorkspaceDataTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AdditionalWorkspaceDataTypeOutput)
}

func (e AdditionalWorkspaceDataType) ToAdditionalWorkspaceDataTypePtrOutput() AdditionalWorkspaceDataTypePtrOutput {
	return e.ToAdditionalWorkspaceDataTypePtrOutputWithContext(context.Background())
}

func (e AdditionalWorkspaceDataType) ToAdditionalWorkspaceDataTypePtrOutputWithContext(ctx context.Context) AdditionalWorkspaceDataTypePtrOutput {
	return AdditionalWorkspaceDataType(e).ToAdditionalWorkspaceDataTypeOutputWithContext(ctx).ToAdditionalWorkspaceDataTypePtrOutputWithContext(ctx)
}

func (e AdditionalWorkspaceDataType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AdditionalWorkspaceDataType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AdditionalWorkspaceDataType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AdditionalWorkspaceDataType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AdditionalWorkspaceDataTypeOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspaceDataTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspaceDataType)(nil)).Elem()
}

func (o AdditionalWorkspaceDataTypeOutput) ToAdditionalWorkspaceDataTypeOutput() AdditionalWorkspaceDataTypeOutput {
	return o
}

func (o AdditionalWorkspaceDataTypeOutput) ToAdditionalWorkspaceDataTypeOutputWithContext(ctx context.Context) AdditionalWorkspaceDataTypeOutput {
	return o
}

func (o AdditionalWorkspaceDataTypeOutput) ToAdditionalWorkspaceDataTypePtrOutput() AdditionalWorkspaceDataTypePtrOutput {
	return o.ToAdditionalWorkspaceDataTypePtrOutputWithContext(context.Background())
}

func (o AdditionalWorkspaceDataTypeOutput) ToAdditionalWorkspaceDataTypePtrOutputWithContext(ctx context.Context) AdditionalWorkspaceDataTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdditionalWorkspaceDataType) *AdditionalWorkspaceDataType {
		return &v
	}).(AdditionalWorkspaceDataTypePtrOutput)
}

func (o AdditionalWorkspaceDataTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AdditionalWorkspaceDataTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AdditionalWorkspaceDataType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AdditionalWorkspaceDataTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AdditionalWorkspaceDataTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AdditionalWorkspaceDataType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AdditionalWorkspaceDataTypePtrOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspaceDataTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalWorkspaceDataType)(nil)).Elem()
}

func (o AdditionalWorkspaceDataTypePtrOutput) ToAdditionalWorkspaceDataTypePtrOutput() AdditionalWorkspaceDataTypePtrOutput {
	return o
}

func (o AdditionalWorkspaceDataTypePtrOutput) ToAdditionalWorkspaceDataTypePtrOutputWithContext(ctx context.Context) AdditionalWorkspaceDataTypePtrOutput {
	return o
}

func (o AdditionalWorkspaceDataTypePtrOutput) Elem() AdditionalWorkspaceDataTypeOutput {
	return o.ApplyT(func(v *AdditionalWorkspaceDataType) AdditionalWorkspaceDataType {
		if v != nil {
			return *v
		}
		var ret AdditionalWorkspaceDataType
		return ret
	}).(AdditionalWorkspaceDataTypeOutput)
}

func (o AdditionalWorkspaceDataTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AdditionalWorkspaceDataTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AdditionalWorkspaceDataType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AdditionalWorkspaceDataTypeInput interface {
	pulumi.Input

	ToAdditionalWorkspaceDataTypeOutput() AdditionalWorkspaceDataTypeOutput
	ToAdditionalWorkspaceDataTypeOutputWithContext(context.Context) AdditionalWorkspaceDataTypeOutput
}

var additionalWorkspaceDataTypePtrType = reflect.TypeOf((**AdditionalWorkspaceDataType)(nil)).Elem()

type AdditionalWorkspaceDataTypePtrInput interface {
	pulumi.Input

	ToAdditionalWorkspaceDataTypePtrOutput() AdditionalWorkspaceDataTypePtrOutput
	ToAdditionalWorkspaceDataTypePtrOutputWithContext(context.Context) AdditionalWorkspaceDataTypePtrOutput
}

type additionalWorkspaceDataTypePtr string

func AdditionalWorkspaceDataTypePtr(v string) AdditionalWorkspaceDataTypePtrInput {
	return (*additionalWorkspaceDataTypePtr)(&v)
}

func (*additionalWorkspaceDataTypePtr) ElementType() reflect.Type {
	return additionalWorkspaceDataTypePtrType
}

func (in *additionalWorkspaceDataTypePtr) ToAdditionalWorkspaceDataTypePtrOutput() AdditionalWorkspaceDataTypePtrOutput {
	return pulumi.ToOutput(in).(AdditionalWorkspaceDataTypePtrOutput)
}

func (in *additionalWorkspaceDataTypePtr) ToAdditionalWorkspaceDataTypePtrOutputWithContext(ctx context.Context) AdditionalWorkspaceDataTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AdditionalWorkspaceDataTypePtrOutput)
}

type AdditionalWorkspaceType string

const (
	AdditionalWorkspaceTypeSentinel = AdditionalWorkspaceType("Sentinel")
)

func (AdditionalWorkspaceType) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspaceType)(nil)).Elem()
}

func (e AdditionalWorkspaceType) ToAdditionalWorkspaceTypeOutput() AdditionalWorkspaceTypeOutput {
	return pulumi.ToOutput(e).(AdditionalWorkspaceTypeOutput)
}

func (e AdditionalWorkspaceType) ToAdditionalWorkspaceTypeOutputWithContext(ctx context.Context) AdditionalWorkspaceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AdditionalWorkspaceTypeOutput)
}

func (e AdditionalWorkspaceType) ToAdditionalWorkspaceTypePtrOutput() AdditionalWorkspaceTypePtrOutput {
	return e.ToAdditionalWorkspaceTypePtrOutputWithContext(context.Background())
}

func (e AdditionalWorkspaceType) ToAdditionalWorkspaceTypePtrOutputWithContext(ctx context.Context) AdditionalWorkspaceTypePtrOutput {
	return AdditionalWorkspaceType(e).ToAdditionalWorkspaceTypeOutputWithContext(ctx).ToAdditionalWorkspaceTypePtrOutputWithContext(ctx)
}

func (e AdditionalWorkspaceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AdditionalWorkspaceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AdditionalWorkspaceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AdditionalWorkspaceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AdditionalWorkspaceTypeOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspaceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalWorkspaceType)(nil)).Elem()
}

func (o AdditionalWorkspaceTypeOutput) ToAdditionalWorkspaceTypeOutput() AdditionalWorkspaceTypeOutput {
	return o
}

func (o AdditionalWorkspaceTypeOutput) ToAdditionalWorkspaceTypeOutputWithContext(ctx context.Context) AdditionalWorkspaceTypeOutput {
	return o
}

func (o AdditionalWorkspaceTypeOutput) ToAdditionalWorkspaceTypePtrOutput() AdditionalWorkspaceTypePtrOutput {
	return o.ToAdditionalWorkspaceTypePtrOutputWithContext(context.Background())
}

func (o AdditionalWorkspaceTypeOutput) ToAdditionalWorkspaceTypePtrOutputWithContext(ctx context.Context) AdditionalWorkspaceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdditionalWorkspaceType) *AdditionalWorkspaceType {
		return &v
	}).(AdditionalWorkspaceTypePtrOutput)
}

func (o AdditionalWorkspaceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AdditionalWorkspaceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AdditionalWorkspaceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AdditionalWorkspaceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AdditionalWorkspaceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AdditionalWorkspaceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AdditionalWorkspaceTypePtrOutput struct{ *pulumi.OutputState }

func (AdditionalWorkspaceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalWorkspaceType)(nil)).Elem()
}

func (o AdditionalWorkspaceTypePtrOutput) ToAdditionalWorkspaceTypePtrOutput() AdditionalWorkspaceTypePtrOutput {
	return o
}

func (o AdditionalWorkspaceTypePtrOutput) ToAdditionalWorkspaceTypePtrOutputWithContext(ctx context.Context) AdditionalWorkspaceTypePtrOutput {
	return o
}

func (o AdditionalWorkspaceTypePtrOutput) Elem() AdditionalWorkspaceTypeOutput {
	return o.ApplyT(func(v *AdditionalWorkspaceType) AdditionalWorkspaceType {
		if v != nil {
			return *v
		}
		var ret AdditionalWorkspaceType
		return ret
	}).(AdditionalWorkspaceTypeOutput)
}

func (o AdditionalWorkspaceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AdditionalWorkspaceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AdditionalWorkspaceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AdditionalWorkspaceTypeInput interface {
	pulumi.Input

	ToAdditionalWorkspaceTypeOutput() AdditionalWorkspaceTypeOutput
	ToAdditionalWorkspaceTypeOutputWithContext(context.Context) AdditionalWorkspaceTypeOutput
}

var additionalWorkspaceTypePtrType = reflect.TypeOf((**AdditionalWorkspaceType)(nil)).Elem()

type AdditionalWorkspaceTypePtrInput interface {
	pulumi.Input

	ToAdditionalWorkspaceTypePtrOutput() AdditionalWorkspaceTypePtrOutput
	ToAdditionalWorkspaceTypePtrOutputWithContext(context.Context) AdditionalWorkspaceTypePtrOutput
}

type additionalWorkspaceTypePtr string

func AdditionalWorkspaceTypePtr(v string) AdditionalWorkspaceTypePtrInput {
	return (*additionalWorkspaceTypePtr)(&v)
}

func (*additionalWorkspaceTypePtr) ElementType() reflect.Type {
	return additionalWorkspaceTypePtrType
}

func (in *additionalWorkspaceTypePtr) ToAdditionalWorkspaceTypePtrOutput() AdditionalWorkspaceTypePtrOutput {
	return pulumi.ToOutput(in).(AdditionalWorkspaceTypePtrOutput)
}

func (in *additionalWorkspaceTypePtr) ToAdditionalWorkspaceTypePtrOutputWithContext(ctx context.Context) AdditionalWorkspaceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AdditionalWorkspaceTypePtrOutput)
}

type DataSource string

const (
	// Devices twin data
	DataSourceTwinData = DataSource("TwinData")
)

func (DataSource) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSource)(nil)).Elem()
}

func (e DataSource) ToDataSourceOutput() DataSourceOutput {
	return pulumi.ToOutput(e).(DataSourceOutput)
}

func (e DataSource) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DataSourceOutput)
}

func (e DataSource) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return e.ToDataSourcePtrOutputWithContext(context.Background())
}

func (e DataSource) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return DataSource(e).ToDataSourceOutputWithContext(ctx).ToDataSourcePtrOutputWithContext(ctx)
}

func (e DataSource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DataSource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DataSource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DataSourceOutput struct{ *pulumi.OutputState }

func (DataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataSource)(nil)).Elem()
}

func (o DataSourceOutput) ToDataSourceOutput() DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourceOutputWithContext(ctx context.Context) DataSourceOutput {
	return o
}

func (o DataSourceOutput) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return o.ToDataSourcePtrOutputWithContext(context.Background())
}

func (o DataSourceOutput) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataSource) *DataSource {
		return &v
	}).(DataSourcePtrOutput)
}

func (o DataSourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DataSourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataSource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DataSourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataSourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DataSource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DataSourcePtrOutput struct{ *pulumi.OutputState }

func (DataSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataSource)(nil)).Elem()
}

func (o DataSourcePtrOutput) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return o
}

func (o DataSourcePtrOutput) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return o
}

func (o DataSourcePtrOutput) Elem() DataSourceOutput {
	return o.ApplyT(func(v *DataSource) DataSource {
		if v != nil {
			return *v
		}
		var ret DataSource
		return ret
	}).(DataSourceOutput)
}

func (o DataSourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DataSourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DataSource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DataSourceInput interface {
	pulumi.Input

	ToDataSourceOutput() DataSourceOutput
	ToDataSourceOutputWithContext(context.Context) DataSourceOutput
}

var dataSourcePtrType = reflect.TypeOf((**DataSource)(nil)).Elem()

type DataSourcePtrInput interface {
	pulumi.Input

	ToDataSourcePtrOutput() DataSourcePtrOutput
	ToDataSourcePtrOutputWithContext(context.Context) DataSourcePtrOutput
}

type dataSourcePtr string

func DataSourcePtr(v string) DataSourcePtrInput {
	return (*dataSourcePtr)(&v)
}

func (*dataSourcePtr) ElementType() reflect.Type {
	return dataSourcePtrType
}

func (in *dataSourcePtr) ToDataSourcePtrOutput() DataSourcePtrOutput {
	return pulumi.ToOutput(in).(DataSourcePtrOutput)
}

func (in *dataSourcePtr) ToDataSourcePtrOutputWithContext(ctx context.Context) DataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DataSourcePtrOutput)
}

type ExportData string

const (
	// Agent raw events
	ExportDataRawEvents = ExportData("RawEvents")
)

func (ExportData) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportData)(nil)).Elem()
}

func (e ExportData) ToExportDataOutput() ExportDataOutput {
	return pulumi.ToOutput(e).(ExportDataOutput)
}

func (e ExportData) ToExportDataOutputWithContext(ctx context.Context) ExportDataOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExportDataOutput)
}

func (e ExportData) ToExportDataPtrOutput() ExportDataPtrOutput {
	return e.ToExportDataPtrOutputWithContext(context.Background())
}

func (e ExportData) ToExportDataPtrOutputWithContext(ctx context.Context) ExportDataPtrOutput {
	return ExportData(e).ToExportDataOutputWithContext(ctx).ToExportDataPtrOutputWithContext(ctx)
}

func (e ExportData) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExportData) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExportData) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExportData) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExportDataOutput struct{ *pulumi.OutputState }

func (ExportDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportData)(nil)).Elem()
}

func (o ExportDataOutput) ToExportDataOutput() ExportDataOutput {
	return o
}

func (o ExportDataOutput) ToExportDataOutputWithContext(ctx context.Context) ExportDataOutput {
	return o
}

func (o ExportDataOutput) ToExportDataPtrOutput() ExportDataPtrOutput {
	return o.ToExportDataPtrOutputWithContext(context.Background())
}

func (o ExportDataOutput) ToExportDataPtrOutputWithContext(ctx context.Context) ExportDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportData) *ExportData {
		return &v
	}).(ExportDataPtrOutput)
}

func (o ExportDataOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExportDataOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExportData) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExportDataOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExportDataOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExportData) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExportDataPtrOutput struct{ *pulumi.OutputState }

func (ExportDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportData)(nil)).Elem()
}

func (o ExportDataPtrOutput) ToExportDataPtrOutput() ExportDataPtrOutput {
	return o
}

func (o ExportDataPtrOutput) ToExportDataPtrOutputWithContext(ctx context.Context) ExportDataPtrOutput {
	return o
}

func (o ExportDataPtrOutput) Elem() ExportDataOutput {
	return o.ApplyT(func(v *ExportData) ExportData {
		if v != nil {
			return *v
		}
		var ret ExportData
		return ret
	}).(ExportDataOutput)
}

func (o ExportDataPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExportDataPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExportData) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExportDataInput interface {
	pulumi.Input

	ToExportDataOutput() ExportDataOutput
	ToExportDataOutputWithContext(context.Context) ExportDataOutput
}

var exportDataPtrType = reflect.TypeOf((**ExportData)(nil)).Elem()

type ExportDataPtrInput interface {
	pulumi.Input

	ToExportDataPtrOutput() ExportDataPtrOutput
	ToExportDataPtrOutputWithContext(context.Context) ExportDataPtrOutput
}

type exportDataPtr string

func ExportDataPtr(v string) ExportDataPtrInput {
	return (*exportDataPtr)(&v)
}

func (*exportDataPtr) ElementType() reflect.Type {
	return exportDataPtrType
}

func (in *exportDataPtr) ToExportDataPtrOutput() ExportDataPtrOutput {
	return pulumi.ToOutput(in).(ExportDataPtrOutput)
}

func (in *exportDataPtr) ToExportDataPtrOutputWithContext(ctx context.Context) ExportDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExportDataPtrOutput)
}

type RecommendationConfigStatus string

const (
	RecommendationConfigStatusDisabled = RecommendationConfigStatus("Disabled")
	RecommendationConfigStatusEnabled  = RecommendationConfigStatus("Enabled")
)

func (RecommendationConfigStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationConfigStatus)(nil)).Elem()
}

func (e RecommendationConfigStatus) ToRecommendationConfigStatusOutput() RecommendationConfigStatusOutput {
	return pulumi.ToOutput(e).(RecommendationConfigStatusOutput)
}

func (e RecommendationConfigStatus) ToRecommendationConfigStatusOutputWithContext(ctx context.Context) RecommendationConfigStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecommendationConfigStatusOutput)
}

func (e RecommendationConfigStatus) ToRecommendationConfigStatusPtrOutput() RecommendationConfigStatusPtrOutput {
	return e.ToRecommendationConfigStatusPtrOutputWithContext(context.Background())
}

func (e RecommendationConfigStatus) ToRecommendationConfigStatusPtrOutputWithContext(ctx context.Context) RecommendationConfigStatusPtrOutput {
	return RecommendationConfigStatus(e).ToRecommendationConfigStatusOutputWithContext(ctx).ToRecommendationConfigStatusPtrOutputWithContext(ctx)
}

func (e RecommendationConfigStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecommendationConfigStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecommendationConfigStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecommendationConfigStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecommendationConfigStatusOutput struct{ *pulumi.OutputState }

func (RecommendationConfigStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationConfigStatus)(nil)).Elem()
}

func (o RecommendationConfigStatusOutput) ToRecommendationConfigStatusOutput() RecommendationConfigStatusOutput {
	return o
}

func (o RecommendationConfigStatusOutput) ToRecommendationConfigStatusOutputWithContext(ctx context.Context) RecommendationConfigStatusOutput {
	return o
}

func (o RecommendationConfigStatusOutput) ToRecommendationConfigStatusPtrOutput() RecommendationConfigStatusPtrOutput {
	return o.ToRecommendationConfigStatusPtrOutputWithContext(context.Background())
}

func (o RecommendationConfigStatusOutput) ToRecommendationConfigStatusPtrOutputWithContext(ctx context.Context) RecommendationConfigStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecommendationConfigStatus) *RecommendationConfigStatus {
		return &v
	}).(RecommendationConfigStatusPtrOutput)
}

func (o RecommendationConfigStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecommendationConfigStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecommendationConfigStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecommendationConfigStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecommendationConfigStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecommendationConfigStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecommendationConfigStatusPtrOutput struct{ *pulumi.OutputState }

func (RecommendationConfigStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecommendationConfigStatus)(nil)).Elem()
}

func (o RecommendationConfigStatusPtrOutput) ToRecommendationConfigStatusPtrOutput() RecommendationConfigStatusPtrOutput {
	return o
}

func (o RecommendationConfigStatusPtrOutput) ToRecommendationConfigStatusPtrOutputWithContext(ctx context.Context) RecommendationConfigStatusPtrOutput {
	return o
}

func (o RecommendationConfigStatusPtrOutput) Elem() RecommendationConfigStatusOutput {
	return o.ApplyT(func(v *RecommendationConfigStatus) RecommendationConfigStatus {
		if v != nil {
			return *v
		}
		var ret RecommendationConfigStatus
		return ret
	}).(RecommendationConfigStatusOutput)
}

func (o RecommendationConfigStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecommendationConfigStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecommendationConfigStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RecommendationConfigStatusInput interface {
	pulumi.Input

	ToRecommendationConfigStatusOutput() RecommendationConfigStatusOutput
	ToRecommendationConfigStatusOutputWithContext(context.Context) RecommendationConfigStatusOutput
}

var recommendationConfigStatusPtrType = reflect.TypeOf((**RecommendationConfigStatus)(nil)).Elem()

type RecommendationConfigStatusPtrInput interface {
	pulumi.Input

	ToRecommendationConfigStatusPtrOutput() RecommendationConfigStatusPtrOutput
	ToRecommendationConfigStatusPtrOutputWithContext(context.Context) RecommendationConfigStatusPtrOutput
}

type recommendationConfigStatusPtr string

func RecommendationConfigStatusPtr(v string) RecommendationConfigStatusPtrInput {
	return (*recommendationConfigStatusPtr)(&v)
}

func (*recommendationConfigStatusPtr) ElementType() reflect.Type {
	return recommendationConfigStatusPtrType
}

func (in *recommendationConfigStatusPtr) ToRecommendationConfigStatusPtrOutput() RecommendationConfigStatusPtrOutput {
	return pulumi.ToOutput(in).(RecommendationConfigStatusPtrOutput)
}

func (in *recommendationConfigStatusPtr) ToRecommendationConfigStatusPtrOutputWithContext(ctx context.Context) RecommendationConfigStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecommendationConfigStatusPtrOutput)
}

type RecommendationType string

const (
	// Authentication schema used for pull an edge module from an ACR repository does not use Service Principal Authentication.
	RecommendationType_IoT_ACRAuthentication = RecommendationType("IoT_ACRAuthentication")
	// IoT agent message size capacity is currently underutilized, causing an increase in the number of sent messages. Adjust message intervals for better utilization.
	RecommendationType_IoT_AgentSendsUnutilizedMessages = RecommendationType("IoT_AgentSendsUnutilizedMessages")
	// Identified security related system configuration issues.
	RecommendationType_IoT_Baseline = RecommendationType("IoT_Baseline")
	// You can optimize Edge Hub memory usage by turning off protocol heads for any protocols not used by Edge modules in your solution.
	RecommendationType_IoT_EdgeHubMemOptimize = RecommendationType("IoT_EdgeHubMemOptimize")
	// Logging is disabled for this edge module.
	RecommendationType_IoT_EdgeLoggingOptions = RecommendationType("IoT_EdgeLoggingOptions")
	// A minority within a device security group has inconsistent Edge Module settings with the rest of their group.
	RecommendationType_IoT_InconsistentModuleSettings = RecommendationType("IoT_InconsistentModuleSettings")
	// Install the Azure Security of Things Agent.
	RecommendationType_IoT_InstallAgent = RecommendationType("IoT_InstallAgent")
	// IP Filter Configuration should have rules defined for allowed traffic and should deny all other traffic by default.
	RecommendationType_IoT_IPFilter_DenyAll = RecommendationType("IoT_IPFilter_DenyAll")
	// An Allow IP Filter rules source IP range is too large. Overly permissive rules might expose your IoT hub to malicious intenders.
	RecommendationType_IoT_IPFilter_PermissiveRule = RecommendationType("IoT_IPFilter_PermissiveRule")
	// A listening endpoint was found on the device.
	RecommendationType_IoT_OpenPorts = RecommendationType("IoT_OpenPorts")
	// An Allowed firewall policy was found (INPUT/OUTPUT). The policy should Deny all traffic by default and define rules to allow necessary communication to/from the device.
	RecommendationType_IoT_PermissiveFirewallPolicy = RecommendationType("IoT_PermissiveFirewallPolicy")
	// A rule in the firewall has been found that contains a permissive pattern for a wide range of IP addresses or Ports.
	RecommendationType_IoT_PermissiveInputFirewallRules = RecommendationType("IoT_PermissiveInputFirewallRules")
	// A rule in the firewall has been found that contains a permissive pattern for a wide range of IP addresses or Ports.
	RecommendationType_IoT_PermissiveOutputFirewallRules = RecommendationType("IoT_PermissiveOutputFirewallRules")
	// Edge module is configured to run in privileged mode, with extensive Linux capabilities or with host-level network access (send/receive data to host machine).
	RecommendationType_IoT_PrivilegedDockerOptions = RecommendationType("IoT_PrivilegedDockerOptions")
	// Same authentication credentials to the IoT Hub used by multiple devices. This could indicate an illegitimate device impersonating a legitimate device. It also exposes the risk of device impersonation by an attacker.
	RecommendationType_IoT_SharedCredentials = RecommendationType("IoT_SharedCredentials")
	// Insecure TLS configurations detected. Immediate upgrade recommended.
	RecommendationType_IoT_VulnerableTLSCipherSuite = RecommendationType("IoT_VulnerableTLSCipherSuite")
)

func (RecommendationType) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationType)(nil)).Elem()
}

func (e RecommendationType) ToRecommendationTypeOutput() RecommendationTypeOutput {
	return pulumi.ToOutput(e).(RecommendationTypeOutput)
}

func (e RecommendationType) ToRecommendationTypeOutputWithContext(ctx context.Context) RecommendationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecommendationTypeOutput)
}

func (e RecommendationType) ToRecommendationTypePtrOutput() RecommendationTypePtrOutput {
	return e.ToRecommendationTypePtrOutputWithContext(context.Background())
}

func (e RecommendationType) ToRecommendationTypePtrOutputWithContext(ctx context.Context) RecommendationTypePtrOutput {
	return RecommendationType(e).ToRecommendationTypeOutputWithContext(ctx).ToRecommendationTypePtrOutputWithContext(ctx)
}

func (e RecommendationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecommendationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecommendationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecommendationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecommendationTypeOutput struct{ *pulumi.OutputState }

func (RecommendationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecommendationType)(nil)).Elem()
}

func (o RecommendationTypeOutput) ToRecommendationTypeOutput() RecommendationTypeOutput {
	return o
}

func (o RecommendationTypeOutput) ToRecommendationTypeOutputWithContext(ctx context.Context) RecommendationTypeOutput {
	return o
}

func (o RecommendationTypeOutput) ToRecommendationTypePtrOutput() RecommendationTypePtrOutput {
	return o.ToRecommendationTypePtrOutputWithContext(context.Background())
}

func (o RecommendationTypeOutput) ToRecommendationTypePtrOutputWithContext(ctx context.Context) RecommendationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecommendationType) *RecommendationType {
		return &v
	}).(RecommendationTypePtrOutput)
}

func (o RecommendationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecommendationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecommendationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecommendationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecommendationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecommendationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecommendationTypePtrOutput struct{ *pulumi.OutputState }

func (RecommendationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecommendationType)(nil)).Elem()
}

func (o RecommendationTypePtrOutput) ToRecommendationTypePtrOutput() RecommendationTypePtrOutput {
	return o
}

func (o RecommendationTypePtrOutput) ToRecommendationTypePtrOutputWithContext(ctx context.Context) RecommendationTypePtrOutput {
	return o
}

func (o RecommendationTypePtrOutput) Elem() RecommendationTypeOutput {
	return o.ApplyT(func(v *RecommendationType) RecommendationType {
		if v != nil {
			return *v
		}
		var ret RecommendationType
		return ret
	}).(RecommendationTypeOutput)
}

func (o RecommendationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecommendationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecommendationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RecommendationTypeInput interface {
	pulumi.Input

	ToRecommendationTypeOutput() RecommendationTypeOutput
	ToRecommendationTypeOutputWithContext(context.Context) RecommendationTypeOutput
}

var recommendationTypePtrType = reflect.TypeOf((**RecommendationType)(nil)).Elem()

type RecommendationTypePtrInput interface {
	pulumi.Input

	ToRecommendationTypePtrOutput() RecommendationTypePtrOutput
	ToRecommendationTypePtrOutputWithContext(context.Context) RecommendationTypePtrOutput
}

type recommendationTypePtr string

func RecommendationTypePtr(v string) RecommendationTypePtrInput {
	return (*recommendationTypePtr)(&v)
}

func (*recommendationTypePtr) ElementType() reflect.Type {
	return recommendationTypePtrType
}

func (in *recommendationTypePtr) ToRecommendationTypePtrOutput() RecommendationTypePtrOutput {
	return pulumi.ToOutput(in).(RecommendationTypePtrOutput)
}

func (in *recommendationTypePtr) ToRecommendationTypePtrOutputWithContext(ctx context.Context) RecommendationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecommendationTypePtrOutput)
}

type SecuritySolutionStatus string

const (
	SecuritySolutionStatusEnabled  = SecuritySolutionStatus("Enabled")
	SecuritySolutionStatusDisabled = SecuritySolutionStatus("Disabled")
)

func (SecuritySolutionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*SecuritySolutionStatus)(nil)).Elem()
}

func (e SecuritySolutionStatus) ToSecuritySolutionStatusOutput() SecuritySolutionStatusOutput {
	return pulumi.ToOutput(e).(SecuritySolutionStatusOutput)
}

func (e SecuritySolutionStatus) ToSecuritySolutionStatusOutputWithContext(ctx context.Context) SecuritySolutionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SecuritySolutionStatusOutput)
}

func (e SecuritySolutionStatus) ToSecuritySolutionStatusPtrOutput() SecuritySolutionStatusPtrOutput {
	return e.ToSecuritySolutionStatusPtrOutputWithContext(context.Background())
}

func (e SecuritySolutionStatus) ToSecuritySolutionStatusPtrOutputWithContext(ctx context.Context) SecuritySolutionStatusPtrOutput {
	return SecuritySolutionStatus(e).ToSecuritySolutionStatusOutputWithContext(ctx).ToSecuritySolutionStatusPtrOutputWithContext(ctx)
}

func (e SecuritySolutionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecuritySolutionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SecuritySolutionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SecuritySolutionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SecuritySolutionStatusOutput struct{ *pulumi.OutputState }

func (SecuritySolutionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecuritySolutionStatus)(nil)).Elem()
}

func (o SecuritySolutionStatusOutput) ToSecuritySolutionStatusOutput() SecuritySolutionStatusOutput {
	return o
}

func (o SecuritySolutionStatusOutput) ToSecuritySolutionStatusOutputWithContext(ctx context.Context) SecuritySolutionStatusOutput {
	return o
}

func (o SecuritySolutionStatusOutput) ToSecuritySolutionStatusPtrOutput() SecuritySolutionStatusPtrOutput {
	return o.ToSecuritySolutionStatusPtrOutputWithContext(context.Background())
}

func (o SecuritySolutionStatusOutput) ToSecuritySolutionStatusPtrOutputWithContext(ctx context.Context) SecuritySolutionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SecuritySolutionStatus) *SecuritySolutionStatus {
		return &v
	}).(SecuritySolutionStatusPtrOutput)
}

func (o SecuritySolutionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SecuritySolutionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecuritySolutionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SecuritySolutionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecuritySolutionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SecuritySolutionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SecuritySolutionStatusPtrOutput struct{ *pulumi.OutputState }

func (SecuritySolutionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecuritySolutionStatus)(nil)).Elem()
}

func (o SecuritySolutionStatusPtrOutput) ToSecuritySolutionStatusPtrOutput() SecuritySolutionStatusPtrOutput {
	return o
}

func (o SecuritySolutionStatusPtrOutput) ToSecuritySolutionStatusPtrOutputWithContext(ctx context.Context) SecuritySolutionStatusPtrOutput {
	return o
}

func (o SecuritySolutionStatusPtrOutput) Elem() SecuritySolutionStatusOutput {
	return o.ApplyT(func(v *SecuritySolutionStatus) SecuritySolutionStatus {
		if v != nil {
			return *v
		}
		var ret SecuritySolutionStatus
		return ret
	}).(SecuritySolutionStatusOutput)
}

func (o SecuritySolutionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SecuritySolutionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SecuritySolutionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SecuritySolutionStatusInput interface {
	pulumi.Input

	ToSecuritySolutionStatusOutput() SecuritySolutionStatusOutput
	ToSecuritySolutionStatusOutputWithContext(context.Context) SecuritySolutionStatusOutput
}

var securitySolutionStatusPtrType = reflect.TypeOf((**SecuritySolutionStatus)(nil)).Elem()

type SecuritySolutionStatusPtrInput interface {
	pulumi.Input

	ToSecuritySolutionStatusPtrOutput() SecuritySolutionStatusPtrOutput
	ToSecuritySolutionStatusPtrOutputWithContext(context.Context) SecuritySolutionStatusPtrOutput
}

type securitySolutionStatusPtr string

func SecuritySolutionStatusPtr(v string) SecuritySolutionStatusPtrInput {
	return (*securitySolutionStatusPtr)(&v)
}

func (*securitySolutionStatusPtr) ElementType() reflect.Type {
	return securitySolutionStatusPtrType
}

func (in *securitySolutionStatusPtr) ToSecuritySolutionStatusPtrOutput() SecuritySolutionStatusPtrOutput {
	return pulumi.ToOutput(in).(SecuritySolutionStatusPtrOutput)
}

func (in *securitySolutionStatusPtr) ToSecuritySolutionStatusPtrOutputWithContext(ctx context.Context) SecuritySolutionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SecuritySolutionStatusPtrOutput)
}

type UnmaskedIpLoggingStatus string

const (
	// Unmasked IP logging is disabled
	UnmaskedIpLoggingStatusDisabled = UnmaskedIpLoggingStatus("Disabled")
	// Unmasked IP logging is enabled
	UnmaskedIpLoggingStatusEnabled = UnmaskedIpLoggingStatus("Enabled")
)

func (UnmaskedIpLoggingStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*UnmaskedIpLoggingStatus)(nil)).Elem()
}

func (e UnmaskedIpLoggingStatus) ToUnmaskedIpLoggingStatusOutput() UnmaskedIpLoggingStatusOutput {
	return pulumi.ToOutput(e).(UnmaskedIpLoggingStatusOutput)
}

func (e UnmaskedIpLoggingStatus) ToUnmaskedIpLoggingStatusOutputWithContext(ctx context.Context) UnmaskedIpLoggingStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UnmaskedIpLoggingStatusOutput)
}

func (e UnmaskedIpLoggingStatus) ToUnmaskedIpLoggingStatusPtrOutput() UnmaskedIpLoggingStatusPtrOutput {
	return e.ToUnmaskedIpLoggingStatusPtrOutputWithContext(context.Background())
}

func (e UnmaskedIpLoggingStatus) ToUnmaskedIpLoggingStatusPtrOutputWithContext(ctx context.Context) UnmaskedIpLoggingStatusPtrOutput {
	return UnmaskedIpLoggingStatus(e).ToUnmaskedIpLoggingStatusOutputWithContext(ctx).ToUnmaskedIpLoggingStatusPtrOutputWithContext(ctx)
}

func (e UnmaskedIpLoggingStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UnmaskedIpLoggingStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UnmaskedIpLoggingStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UnmaskedIpLoggingStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UnmaskedIpLoggingStatusOutput struct{ *pulumi.OutputState }

func (UnmaskedIpLoggingStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnmaskedIpLoggingStatus)(nil)).Elem()
}

func (o UnmaskedIpLoggingStatusOutput) ToUnmaskedIpLoggingStatusOutput() UnmaskedIpLoggingStatusOutput {
	return o
}

func (o UnmaskedIpLoggingStatusOutput) ToUnmaskedIpLoggingStatusOutputWithContext(ctx context.Context) UnmaskedIpLoggingStatusOutput {
	return o
}

func (o UnmaskedIpLoggingStatusOutput) ToUnmaskedIpLoggingStatusPtrOutput() UnmaskedIpLoggingStatusPtrOutput {
	return o.ToUnmaskedIpLoggingStatusPtrOutputWithContext(context.Background())
}

func (o UnmaskedIpLoggingStatusOutput) ToUnmaskedIpLoggingStatusPtrOutputWithContext(ctx context.Context) UnmaskedIpLoggingStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UnmaskedIpLoggingStatus) *UnmaskedIpLoggingStatus {
		return &v
	}).(UnmaskedIpLoggingStatusPtrOutput)
}

func (o UnmaskedIpLoggingStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UnmaskedIpLoggingStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UnmaskedIpLoggingStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UnmaskedIpLoggingStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UnmaskedIpLoggingStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UnmaskedIpLoggingStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UnmaskedIpLoggingStatusPtrOutput struct{ *pulumi.OutputState }

func (UnmaskedIpLoggingStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnmaskedIpLoggingStatus)(nil)).Elem()
}

func (o UnmaskedIpLoggingStatusPtrOutput) ToUnmaskedIpLoggingStatusPtrOutput() UnmaskedIpLoggingStatusPtrOutput {
	return o
}

func (o UnmaskedIpLoggingStatusPtrOutput) ToUnmaskedIpLoggingStatusPtrOutputWithContext(ctx context.Context) UnmaskedIpLoggingStatusPtrOutput {
	return o
}

func (o UnmaskedIpLoggingStatusPtrOutput) Elem() UnmaskedIpLoggingStatusOutput {
	return o.ApplyT(func(v *UnmaskedIpLoggingStatus) UnmaskedIpLoggingStatus {
		if v != nil {
			return *v
		}
		var ret UnmaskedIpLoggingStatus
		return ret
	}).(UnmaskedIpLoggingStatusOutput)
}

func (o UnmaskedIpLoggingStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UnmaskedIpLoggingStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UnmaskedIpLoggingStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UnmaskedIpLoggingStatusInput interface {
	pulumi.Input

	ToUnmaskedIpLoggingStatusOutput() UnmaskedIpLoggingStatusOutput
	ToUnmaskedIpLoggingStatusOutputWithContext(context.Context) UnmaskedIpLoggingStatusOutput
}

var unmaskedIpLoggingStatusPtrType = reflect.TypeOf((**UnmaskedIpLoggingStatus)(nil)).Elem()

type UnmaskedIpLoggingStatusPtrInput interface {
	pulumi.Input

	ToUnmaskedIpLoggingStatusPtrOutput() UnmaskedIpLoggingStatusPtrOutput
	ToUnmaskedIpLoggingStatusPtrOutputWithContext(context.Context) UnmaskedIpLoggingStatusPtrOutput
}

type unmaskedIpLoggingStatusPtr string

func UnmaskedIpLoggingStatusPtr(v string) UnmaskedIpLoggingStatusPtrInput {
	return (*unmaskedIpLoggingStatusPtr)(&v)
}

func (*unmaskedIpLoggingStatusPtr) ElementType() reflect.Type {
	return unmaskedIpLoggingStatusPtrType
}

func (in *unmaskedIpLoggingStatusPtr) ToUnmaskedIpLoggingStatusPtrOutput() UnmaskedIpLoggingStatusPtrOutput {
	return pulumi.ToOutput(in).(UnmaskedIpLoggingStatusPtrOutput)
}

func (in *unmaskedIpLoggingStatusPtr) ToUnmaskedIpLoggingStatusPtrOutputWithContext(ctx context.Context) UnmaskedIpLoggingStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UnmaskedIpLoggingStatusPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdditionalWorkspaceDataTypeOutput{})
	pulumi.RegisterOutputType(AdditionalWorkspaceDataTypePtrOutput{})
	pulumi.RegisterOutputType(AdditionalWorkspaceTypeOutput{})
	pulumi.RegisterOutputType(AdditionalWorkspaceTypePtrOutput{})
	pulumi.RegisterOutputType(DataSourceOutput{})
	pulumi.RegisterOutputType(DataSourcePtrOutput{})
	pulumi.RegisterOutputType(ExportDataOutput{})
	pulumi.RegisterOutputType(ExportDataPtrOutput{})
	pulumi.RegisterOutputType(RecommendationConfigStatusOutput{})
	pulumi.RegisterOutputType(RecommendationConfigStatusPtrOutput{})
	pulumi.RegisterOutputType(RecommendationTypeOutput{})
	pulumi.RegisterOutputType(RecommendationTypePtrOutput{})
	pulumi.RegisterOutputType(SecuritySolutionStatusOutput{})
	pulumi.RegisterOutputType(SecuritySolutionStatusPtrOutput{})
	pulumi.RegisterOutputType(UnmaskedIpLoggingStatusOutput{})
	pulumi.RegisterOutputType(UnmaskedIpLoggingStatusPtrOutput{})
}
