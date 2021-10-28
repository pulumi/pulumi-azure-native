


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExportDefinition struct {
	DataSet    *QueryDataset    `pulumi:"dataSet"`
	TimePeriod *QueryTimePeriod `pulumi:"timePeriod"`
	Timeframe  string           `pulumi:"timeframe"`
	Type       string           `pulumi:"type"`
}





type ExportDefinitionInput interface {
	pulumi.Input

	ToExportDefinitionOutput() ExportDefinitionOutput
	ToExportDefinitionOutputWithContext(context.Context) ExportDefinitionOutput
}

type ExportDefinitionArgs struct {
	DataSet    QueryDatasetPtrInput    `pulumi:"dataSet"`
	TimePeriod QueryTimePeriodPtrInput `pulumi:"timePeriod"`
	Timeframe  pulumi.StringInput      `pulumi:"timeframe"`
	Type       pulumi.StringInput      `pulumi:"type"`
}

func (ExportDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDefinition)(nil)).Elem()
}

func (i ExportDefinitionArgs) ToExportDefinitionOutput() ExportDefinitionOutput {
	return i.ToExportDefinitionOutputWithContext(context.Background())
}

func (i ExportDefinitionArgs) ToExportDefinitionOutputWithContext(ctx context.Context) ExportDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDefinitionOutput)
}

func (i ExportDefinitionArgs) ToExportDefinitionPtrOutput() ExportDefinitionPtrOutput {
	return i.ToExportDefinitionPtrOutputWithContext(context.Background())
}

func (i ExportDefinitionArgs) ToExportDefinitionPtrOutputWithContext(ctx context.Context) ExportDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDefinitionOutput).ToExportDefinitionPtrOutputWithContext(ctx)
}









type ExportDefinitionPtrInput interface {
	pulumi.Input

	ToExportDefinitionPtrOutput() ExportDefinitionPtrOutput
	ToExportDefinitionPtrOutputWithContext(context.Context) ExportDefinitionPtrOutput
}

type exportDefinitionPtrType ExportDefinitionArgs

func ExportDefinitionPtr(v *ExportDefinitionArgs) ExportDefinitionPtrInput {
	return (*exportDefinitionPtrType)(v)
}

func (*exportDefinitionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDefinition)(nil)).Elem()
}

func (i *exportDefinitionPtrType) ToExportDefinitionPtrOutput() ExportDefinitionPtrOutput {
	return i.ToExportDefinitionPtrOutputWithContext(context.Background())
}

func (i *exportDefinitionPtrType) ToExportDefinitionPtrOutputWithContext(ctx context.Context) ExportDefinitionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDefinitionPtrOutput)
}

type ExportDefinitionOutput struct{ *pulumi.OutputState }

func (ExportDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDefinition)(nil)).Elem()
}

func (o ExportDefinitionOutput) ToExportDefinitionOutput() ExportDefinitionOutput {
	return o
}

func (o ExportDefinitionOutput) ToExportDefinitionOutputWithContext(ctx context.Context) ExportDefinitionOutput {
	return o
}

func (o ExportDefinitionOutput) ToExportDefinitionPtrOutput() ExportDefinitionPtrOutput {
	return o.ToExportDefinitionPtrOutputWithContext(context.Background())
}

func (o ExportDefinitionOutput) ToExportDefinitionPtrOutputWithContext(ctx context.Context) ExportDefinitionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDefinition) *ExportDefinition {
		return &v
	}).(ExportDefinitionPtrOutput)
}

func (o ExportDefinitionOutput) DataSet() QueryDatasetPtrOutput {
	return o.ApplyT(func(v ExportDefinition) *QueryDataset { return v.DataSet }).(QueryDatasetPtrOutput)
}

func (o ExportDefinitionOutput) TimePeriod() QueryTimePeriodPtrOutput {
	return o.ApplyT(func(v ExportDefinition) *QueryTimePeriod { return v.TimePeriod }).(QueryTimePeriodPtrOutput)
}

func (o ExportDefinitionOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDefinition) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ExportDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type ExportDefinitionPtrOutput struct{ *pulumi.OutputState }

func (ExportDefinitionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDefinition)(nil)).Elem()
}

func (o ExportDefinitionPtrOutput) ToExportDefinitionPtrOutput() ExportDefinitionPtrOutput {
	return o
}

func (o ExportDefinitionPtrOutput) ToExportDefinitionPtrOutputWithContext(ctx context.Context) ExportDefinitionPtrOutput {
	return o
}

func (o ExportDefinitionPtrOutput) Elem() ExportDefinitionOutput {
	return o.ApplyT(func(v *ExportDefinition) ExportDefinition {
		if v != nil {
			return *v
		}
		var ret ExportDefinition
		return ret
	}).(ExportDefinitionOutput)
}

func (o ExportDefinitionPtrOutput) DataSet() QueryDatasetPtrOutput {
	return o.ApplyT(func(v *ExportDefinition) *QueryDataset {
		if v == nil {
			return nil
		}
		return v.DataSet
	}).(QueryDatasetPtrOutput)
}

func (o ExportDefinitionPtrOutput) TimePeriod() QueryTimePeriodPtrOutput {
	return o.ApplyT(func(v *ExportDefinition) *QueryTimePeriod {
		if v == nil {
			return nil
		}
		return v.TimePeriod
	}).(QueryTimePeriodPtrOutput)
}

func (o ExportDefinitionPtrOutput) Timeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDefinition) *string {
		if v == nil {
			return nil
		}
		return &v.Timeframe
	}).(pulumi.StringPtrOutput)
}

func (o ExportDefinitionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDefinition) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ExportDefinitionResponse struct {
	DataSet    *QueryDatasetResponse    `pulumi:"dataSet"`
	TimePeriod *QueryTimePeriodResponse `pulumi:"timePeriod"`
	Timeframe  string                   `pulumi:"timeframe"`
	Type       string                   `pulumi:"type"`
}





type ExportDefinitionResponseInput interface {
	pulumi.Input

	ToExportDefinitionResponseOutput() ExportDefinitionResponseOutput
	ToExportDefinitionResponseOutputWithContext(context.Context) ExportDefinitionResponseOutput
}

type ExportDefinitionResponseArgs struct {
	DataSet    QueryDatasetResponsePtrInput    `pulumi:"dataSet"`
	TimePeriod QueryTimePeriodResponsePtrInput `pulumi:"timePeriod"`
	Timeframe  pulumi.StringInput              `pulumi:"timeframe"`
	Type       pulumi.StringInput              `pulumi:"type"`
}

func (ExportDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDefinitionResponse)(nil)).Elem()
}

func (i ExportDefinitionResponseArgs) ToExportDefinitionResponseOutput() ExportDefinitionResponseOutput {
	return i.ToExportDefinitionResponseOutputWithContext(context.Background())
}

func (i ExportDefinitionResponseArgs) ToExportDefinitionResponseOutputWithContext(ctx context.Context) ExportDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDefinitionResponseOutput)
}

func (i ExportDefinitionResponseArgs) ToExportDefinitionResponsePtrOutput() ExportDefinitionResponsePtrOutput {
	return i.ToExportDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i ExportDefinitionResponseArgs) ToExportDefinitionResponsePtrOutputWithContext(ctx context.Context) ExportDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDefinitionResponseOutput).ToExportDefinitionResponsePtrOutputWithContext(ctx)
}









type ExportDefinitionResponsePtrInput interface {
	pulumi.Input

	ToExportDefinitionResponsePtrOutput() ExportDefinitionResponsePtrOutput
	ToExportDefinitionResponsePtrOutputWithContext(context.Context) ExportDefinitionResponsePtrOutput
}

type exportDefinitionResponsePtrType ExportDefinitionResponseArgs

func ExportDefinitionResponsePtr(v *ExportDefinitionResponseArgs) ExportDefinitionResponsePtrInput {
	return (*exportDefinitionResponsePtrType)(v)
}

func (*exportDefinitionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDefinitionResponse)(nil)).Elem()
}

func (i *exportDefinitionResponsePtrType) ToExportDefinitionResponsePtrOutput() ExportDefinitionResponsePtrOutput {
	return i.ToExportDefinitionResponsePtrOutputWithContext(context.Background())
}

func (i *exportDefinitionResponsePtrType) ToExportDefinitionResponsePtrOutputWithContext(ctx context.Context) ExportDefinitionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDefinitionResponsePtrOutput)
}

type ExportDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ExportDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDefinitionResponse)(nil)).Elem()
}

func (o ExportDefinitionResponseOutput) ToExportDefinitionResponseOutput() ExportDefinitionResponseOutput {
	return o
}

func (o ExportDefinitionResponseOutput) ToExportDefinitionResponseOutputWithContext(ctx context.Context) ExportDefinitionResponseOutput {
	return o
}

func (o ExportDefinitionResponseOutput) ToExportDefinitionResponsePtrOutput() ExportDefinitionResponsePtrOutput {
	return o.ToExportDefinitionResponsePtrOutputWithContext(context.Background())
}

func (o ExportDefinitionResponseOutput) ToExportDefinitionResponsePtrOutputWithContext(ctx context.Context) ExportDefinitionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDefinitionResponse) *ExportDefinitionResponse {
		return &v
	}).(ExportDefinitionResponsePtrOutput)
}

func (o ExportDefinitionResponseOutput) DataSet() QueryDatasetResponsePtrOutput {
	return o.ApplyT(func(v ExportDefinitionResponse) *QueryDatasetResponse { return v.DataSet }).(QueryDatasetResponsePtrOutput)
}

func (o ExportDefinitionResponseOutput) TimePeriod() QueryTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v ExportDefinitionResponse) *QueryTimePeriodResponse { return v.TimePeriod }).(QueryTimePeriodResponsePtrOutput)
}

func (o ExportDefinitionResponseOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDefinitionResponse) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ExportDefinitionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDefinitionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ExportDefinitionResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportDefinitionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDefinitionResponse)(nil)).Elem()
}

func (o ExportDefinitionResponsePtrOutput) ToExportDefinitionResponsePtrOutput() ExportDefinitionResponsePtrOutput {
	return o
}

func (o ExportDefinitionResponsePtrOutput) ToExportDefinitionResponsePtrOutputWithContext(ctx context.Context) ExportDefinitionResponsePtrOutput {
	return o
}

func (o ExportDefinitionResponsePtrOutput) Elem() ExportDefinitionResponseOutput {
	return o.ApplyT(func(v *ExportDefinitionResponse) ExportDefinitionResponse {
		if v != nil {
			return *v
		}
		var ret ExportDefinitionResponse
		return ret
	}).(ExportDefinitionResponseOutput)
}

func (o ExportDefinitionResponsePtrOutput) DataSet() QueryDatasetResponsePtrOutput {
	return o.ApplyT(func(v *ExportDefinitionResponse) *QueryDatasetResponse {
		if v == nil {
			return nil
		}
		return v.DataSet
	}).(QueryDatasetResponsePtrOutput)
}

func (o ExportDefinitionResponsePtrOutput) TimePeriod() QueryTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v *ExportDefinitionResponse) *QueryTimePeriodResponse {
		if v == nil {
			return nil
		}
		return v.TimePeriod
	}).(QueryTimePeriodResponsePtrOutput)
}

func (o ExportDefinitionResponsePtrOutput) Timeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Timeframe
	}).(pulumi.StringPtrOutput)
}

func (o ExportDefinitionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDefinitionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type ExportDeliveryDestination struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
}





type ExportDeliveryDestinationInput interface {
	pulumi.Input

	ToExportDeliveryDestinationOutput() ExportDeliveryDestinationOutput
	ToExportDeliveryDestinationOutputWithContext(context.Context) ExportDeliveryDestinationOutput
}

type ExportDeliveryDestinationArgs struct {
	Container      pulumi.StringInput    `pulumi:"container"`
	ResourceId     pulumi.StringInput    `pulumi:"resourceId"`
	RootFolderPath pulumi.StringPtrInput `pulumi:"rootFolderPath"`
}

func (ExportDeliveryDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryDestination)(nil)).Elem()
}

func (i ExportDeliveryDestinationArgs) ToExportDeliveryDestinationOutput() ExportDeliveryDestinationOutput {
	return i.ToExportDeliveryDestinationOutputWithContext(context.Background())
}

func (i ExportDeliveryDestinationArgs) ToExportDeliveryDestinationOutputWithContext(ctx context.Context) ExportDeliveryDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationOutput)
}

func (i ExportDeliveryDestinationArgs) ToExportDeliveryDestinationPtrOutput() ExportDeliveryDestinationPtrOutput {
	return i.ToExportDeliveryDestinationPtrOutputWithContext(context.Background())
}

func (i ExportDeliveryDestinationArgs) ToExportDeliveryDestinationPtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationOutput).ToExportDeliveryDestinationPtrOutputWithContext(ctx)
}









type ExportDeliveryDestinationPtrInput interface {
	pulumi.Input

	ToExportDeliveryDestinationPtrOutput() ExportDeliveryDestinationPtrOutput
	ToExportDeliveryDestinationPtrOutputWithContext(context.Context) ExportDeliveryDestinationPtrOutput
}

type exportDeliveryDestinationPtrType ExportDeliveryDestinationArgs

func ExportDeliveryDestinationPtr(v *ExportDeliveryDestinationArgs) ExportDeliveryDestinationPtrInput {
	return (*exportDeliveryDestinationPtrType)(v)
}

func (*exportDeliveryDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryDestination)(nil)).Elem()
}

func (i *exportDeliveryDestinationPtrType) ToExportDeliveryDestinationPtrOutput() ExportDeliveryDestinationPtrOutput {
	return i.ToExportDeliveryDestinationPtrOutputWithContext(context.Background())
}

func (i *exportDeliveryDestinationPtrType) ToExportDeliveryDestinationPtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationPtrOutput)
}

type ExportDeliveryDestinationOutput struct{ *pulumi.OutputState }

func (ExportDeliveryDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryDestination)(nil)).Elem()
}

func (o ExportDeliveryDestinationOutput) ToExportDeliveryDestinationOutput() ExportDeliveryDestinationOutput {
	return o
}

func (o ExportDeliveryDestinationOutput) ToExportDeliveryDestinationOutputWithContext(ctx context.Context) ExportDeliveryDestinationOutput {
	return o
}

func (o ExportDeliveryDestinationOutput) ToExportDeliveryDestinationPtrOutput() ExportDeliveryDestinationPtrOutput {
	return o.ToExportDeliveryDestinationPtrOutputWithContext(context.Background())
}

func (o ExportDeliveryDestinationOutput) ToExportDeliveryDestinationPtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDeliveryDestination) *ExportDeliveryDestination {
		return &v
	}).(ExportDeliveryDestinationPtrOutput)
}

func (o ExportDeliveryDestinationOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) string { return v.Container }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ExportDeliveryDestinationPtrOutput struct{ *pulumi.OutputState }

func (ExportDeliveryDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryDestination)(nil)).Elem()
}

func (o ExportDeliveryDestinationPtrOutput) ToExportDeliveryDestinationPtrOutput() ExportDeliveryDestinationPtrOutput {
	return o
}

func (o ExportDeliveryDestinationPtrOutput) ToExportDeliveryDestinationPtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationPtrOutput {
	return o
}

func (o ExportDeliveryDestinationPtrOutput) Elem() ExportDeliveryDestinationOutput {
	return o.ApplyT(func(v *ExportDeliveryDestination) ExportDeliveryDestination {
		if v != nil {
			return *v
		}
		var ret ExportDeliveryDestination
		return ret
	}).(ExportDeliveryDestinationOutput)
}

func (o ExportDeliveryDestinationPtrOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestination) *string {
		if v == nil {
			return nil
		}
		return &v.Container
	}).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestination) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationPtrOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestination) *string {
		if v == nil {
			return nil
		}
		return v.RootFolderPath
	}).(pulumi.StringPtrOutput)
}

type ExportDeliveryDestinationResponse struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
}





type ExportDeliveryDestinationResponseInput interface {
	pulumi.Input

	ToExportDeliveryDestinationResponseOutput() ExportDeliveryDestinationResponseOutput
	ToExportDeliveryDestinationResponseOutputWithContext(context.Context) ExportDeliveryDestinationResponseOutput
}

type ExportDeliveryDestinationResponseArgs struct {
	Container      pulumi.StringInput    `pulumi:"container"`
	ResourceId     pulumi.StringInput    `pulumi:"resourceId"`
	RootFolderPath pulumi.StringPtrInput `pulumi:"rootFolderPath"`
}

func (ExportDeliveryDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryDestinationResponse)(nil)).Elem()
}

func (i ExportDeliveryDestinationResponseArgs) ToExportDeliveryDestinationResponseOutput() ExportDeliveryDestinationResponseOutput {
	return i.ToExportDeliveryDestinationResponseOutputWithContext(context.Background())
}

func (i ExportDeliveryDestinationResponseArgs) ToExportDeliveryDestinationResponseOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationResponseOutput)
}

func (i ExportDeliveryDestinationResponseArgs) ToExportDeliveryDestinationResponsePtrOutput() ExportDeliveryDestinationResponsePtrOutput {
	return i.ToExportDeliveryDestinationResponsePtrOutputWithContext(context.Background())
}

func (i ExportDeliveryDestinationResponseArgs) ToExportDeliveryDestinationResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationResponseOutput).ToExportDeliveryDestinationResponsePtrOutputWithContext(ctx)
}









type ExportDeliveryDestinationResponsePtrInput interface {
	pulumi.Input

	ToExportDeliveryDestinationResponsePtrOutput() ExportDeliveryDestinationResponsePtrOutput
	ToExportDeliveryDestinationResponsePtrOutputWithContext(context.Context) ExportDeliveryDestinationResponsePtrOutput
}

type exportDeliveryDestinationResponsePtrType ExportDeliveryDestinationResponseArgs

func ExportDeliveryDestinationResponsePtr(v *ExportDeliveryDestinationResponseArgs) ExportDeliveryDestinationResponsePtrInput {
	return (*exportDeliveryDestinationResponsePtrType)(v)
}

func (*exportDeliveryDestinationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryDestinationResponse)(nil)).Elem()
}

func (i *exportDeliveryDestinationResponsePtrType) ToExportDeliveryDestinationResponsePtrOutput() ExportDeliveryDestinationResponsePtrOutput {
	return i.ToExportDeliveryDestinationResponsePtrOutputWithContext(context.Background())
}

func (i *exportDeliveryDestinationResponsePtrType) ToExportDeliveryDestinationResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryDestinationResponsePtrOutput)
}

type ExportDeliveryDestinationResponseOutput struct{ *pulumi.OutputState }

func (ExportDeliveryDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryDestinationResponse)(nil)).Elem()
}

func (o ExportDeliveryDestinationResponseOutput) ToExportDeliveryDestinationResponseOutput() ExportDeliveryDestinationResponseOutput {
	return o
}

func (o ExportDeliveryDestinationResponseOutput) ToExportDeliveryDestinationResponseOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponseOutput {
	return o
}

func (o ExportDeliveryDestinationResponseOutput) ToExportDeliveryDestinationResponsePtrOutput() ExportDeliveryDestinationResponsePtrOutput {
	return o.ToExportDeliveryDestinationResponsePtrOutputWithContext(context.Background())
}

func (o ExportDeliveryDestinationResponseOutput) ToExportDeliveryDestinationResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDeliveryDestinationResponse) *ExportDeliveryDestinationResponse {
		return &v
	}).(ExportDeliveryDestinationResponsePtrOutput)
}

func (o ExportDeliveryDestinationResponseOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) string { return v.Container }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationResponseOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ExportDeliveryDestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportDeliveryDestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryDestinationResponse)(nil)).Elem()
}

func (o ExportDeliveryDestinationResponsePtrOutput) ToExportDeliveryDestinationResponsePtrOutput() ExportDeliveryDestinationResponsePtrOutput {
	return o
}

func (o ExportDeliveryDestinationResponsePtrOutput) ToExportDeliveryDestinationResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryDestinationResponsePtrOutput {
	return o
}

func (o ExportDeliveryDestinationResponsePtrOutput) Elem() ExportDeliveryDestinationResponseOutput {
	return o.ApplyT(func(v *ExportDeliveryDestinationResponse) ExportDeliveryDestinationResponse {
		if v != nil {
			return *v
		}
		var ret ExportDeliveryDestinationResponse
		return ret
	}).(ExportDeliveryDestinationResponseOutput)
}

func (o ExportDeliveryDestinationResponsePtrOutput) Container() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Container
	}).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationResponsePtrOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RootFolderPath
	}).(pulumi.StringPtrOutput)
}

type ExportDeliveryInfo struct {
	Destination ExportDeliveryDestination `pulumi:"destination"`
}





type ExportDeliveryInfoInput interface {
	pulumi.Input

	ToExportDeliveryInfoOutput() ExportDeliveryInfoOutput
	ToExportDeliveryInfoOutputWithContext(context.Context) ExportDeliveryInfoOutput
}

type ExportDeliveryInfoArgs struct {
	Destination ExportDeliveryDestinationInput `pulumi:"destination"`
}

func (ExportDeliveryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryInfo)(nil)).Elem()
}

func (i ExportDeliveryInfoArgs) ToExportDeliveryInfoOutput() ExportDeliveryInfoOutput {
	return i.ToExportDeliveryInfoOutputWithContext(context.Background())
}

func (i ExportDeliveryInfoArgs) ToExportDeliveryInfoOutputWithContext(ctx context.Context) ExportDeliveryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoOutput)
}

func (i ExportDeliveryInfoArgs) ToExportDeliveryInfoPtrOutput() ExportDeliveryInfoPtrOutput {
	return i.ToExportDeliveryInfoPtrOutputWithContext(context.Background())
}

func (i ExportDeliveryInfoArgs) ToExportDeliveryInfoPtrOutputWithContext(ctx context.Context) ExportDeliveryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoOutput).ToExportDeliveryInfoPtrOutputWithContext(ctx)
}









type ExportDeliveryInfoPtrInput interface {
	pulumi.Input

	ToExportDeliveryInfoPtrOutput() ExportDeliveryInfoPtrOutput
	ToExportDeliveryInfoPtrOutputWithContext(context.Context) ExportDeliveryInfoPtrOutput
}

type exportDeliveryInfoPtrType ExportDeliveryInfoArgs

func ExportDeliveryInfoPtr(v *ExportDeliveryInfoArgs) ExportDeliveryInfoPtrInput {
	return (*exportDeliveryInfoPtrType)(v)
}

func (*exportDeliveryInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryInfo)(nil)).Elem()
}

func (i *exportDeliveryInfoPtrType) ToExportDeliveryInfoPtrOutput() ExportDeliveryInfoPtrOutput {
	return i.ToExportDeliveryInfoPtrOutputWithContext(context.Background())
}

func (i *exportDeliveryInfoPtrType) ToExportDeliveryInfoPtrOutputWithContext(ctx context.Context) ExportDeliveryInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoPtrOutput)
}

type ExportDeliveryInfoOutput struct{ *pulumi.OutputState }

func (ExportDeliveryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryInfo)(nil)).Elem()
}

func (o ExportDeliveryInfoOutput) ToExportDeliveryInfoOutput() ExportDeliveryInfoOutput {
	return o
}

func (o ExportDeliveryInfoOutput) ToExportDeliveryInfoOutputWithContext(ctx context.Context) ExportDeliveryInfoOutput {
	return o
}

func (o ExportDeliveryInfoOutput) ToExportDeliveryInfoPtrOutput() ExportDeliveryInfoPtrOutput {
	return o.ToExportDeliveryInfoPtrOutputWithContext(context.Background())
}

func (o ExportDeliveryInfoOutput) ToExportDeliveryInfoPtrOutputWithContext(ctx context.Context) ExportDeliveryInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDeliveryInfo) *ExportDeliveryInfo {
		return &v
	}).(ExportDeliveryInfoPtrOutput)
}

func (o ExportDeliveryInfoOutput) Destination() ExportDeliveryDestinationOutput {
	return o.ApplyT(func(v ExportDeliveryInfo) ExportDeliveryDestination { return v.Destination }).(ExportDeliveryDestinationOutput)
}

type ExportDeliveryInfoPtrOutput struct{ *pulumi.OutputState }

func (ExportDeliveryInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryInfo)(nil)).Elem()
}

func (o ExportDeliveryInfoPtrOutput) ToExportDeliveryInfoPtrOutput() ExportDeliveryInfoPtrOutput {
	return o
}

func (o ExportDeliveryInfoPtrOutput) ToExportDeliveryInfoPtrOutputWithContext(ctx context.Context) ExportDeliveryInfoPtrOutput {
	return o
}

func (o ExportDeliveryInfoPtrOutput) Elem() ExportDeliveryInfoOutput {
	return o.ApplyT(func(v *ExportDeliveryInfo) ExportDeliveryInfo {
		if v != nil {
			return *v
		}
		var ret ExportDeliveryInfo
		return ret
	}).(ExportDeliveryInfoOutput)
}

func (o ExportDeliveryInfoPtrOutput) Destination() ExportDeliveryDestinationPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryInfo) *ExportDeliveryDestination {
		if v == nil {
			return nil
		}
		return &v.Destination
	}).(ExportDeliveryDestinationPtrOutput)
}

type ExportDeliveryInfoResponse struct {
	Destination ExportDeliveryDestinationResponse `pulumi:"destination"`
}





type ExportDeliveryInfoResponseInput interface {
	pulumi.Input

	ToExportDeliveryInfoResponseOutput() ExportDeliveryInfoResponseOutput
	ToExportDeliveryInfoResponseOutputWithContext(context.Context) ExportDeliveryInfoResponseOutput
}

type ExportDeliveryInfoResponseArgs struct {
	Destination ExportDeliveryDestinationResponseInput `pulumi:"destination"`
}

func (ExportDeliveryInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryInfoResponse)(nil)).Elem()
}

func (i ExportDeliveryInfoResponseArgs) ToExportDeliveryInfoResponseOutput() ExportDeliveryInfoResponseOutput {
	return i.ToExportDeliveryInfoResponseOutputWithContext(context.Background())
}

func (i ExportDeliveryInfoResponseArgs) ToExportDeliveryInfoResponseOutputWithContext(ctx context.Context) ExportDeliveryInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoResponseOutput)
}

func (i ExportDeliveryInfoResponseArgs) ToExportDeliveryInfoResponsePtrOutput() ExportDeliveryInfoResponsePtrOutput {
	return i.ToExportDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (i ExportDeliveryInfoResponseArgs) ToExportDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoResponseOutput).ToExportDeliveryInfoResponsePtrOutputWithContext(ctx)
}









type ExportDeliveryInfoResponsePtrInput interface {
	pulumi.Input

	ToExportDeliveryInfoResponsePtrOutput() ExportDeliveryInfoResponsePtrOutput
	ToExportDeliveryInfoResponsePtrOutputWithContext(context.Context) ExportDeliveryInfoResponsePtrOutput
}

type exportDeliveryInfoResponsePtrType ExportDeliveryInfoResponseArgs

func ExportDeliveryInfoResponsePtr(v *ExportDeliveryInfoResponseArgs) ExportDeliveryInfoResponsePtrInput {
	return (*exportDeliveryInfoResponsePtrType)(v)
}

func (*exportDeliveryInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryInfoResponse)(nil)).Elem()
}

func (i *exportDeliveryInfoResponsePtrType) ToExportDeliveryInfoResponsePtrOutput() ExportDeliveryInfoResponsePtrOutput {
	return i.ToExportDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (i *exportDeliveryInfoResponsePtrType) ToExportDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDeliveryInfoResponsePtrOutput)
}

type ExportDeliveryInfoResponseOutput struct{ *pulumi.OutputState }

func (ExportDeliveryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDeliveryInfoResponse)(nil)).Elem()
}

func (o ExportDeliveryInfoResponseOutput) ToExportDeliveryInfoResponseOutput() ExportDeliveryInfoResponseOutput {
	return o
}

func (o ExportDeliveryInfoResponseOutput) ToExportDeliveryInfoResponseOutputWithContext(ctx context.Context) ExportDeliveryInfoResponseOutput {
	return o
}

func (o ExportDeliveryInfoResponseOutput) ToExportDeliveryInfoResponsePtrOutput() ExportDeliveryInfoResponsePtrOutput {
	return o.ToExportDeliveryInfoResponsePtrOutputWithContext(context.Background())
}

func (o ExportDeliveryInfoResponseOutput) ToExportDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDeliveryInfoResponse) *ExportDeliveryInfoResponse {
		return &v
	}).(ExportDeliveryInfoResponsePtrOutput)
}

func (o ExportDeliveryInfoResponseOutput) Destination() ExportDeliveryDestinationResponseOutput {
	return o.ApplyT(func(v ExportDeliveryInfoResponse) ExportDeliveryDestinationResponse { return v.Destination }).(ExportDeliveryDestinationResponseOutput)
}

type ExportDeliveryInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportDeliveryInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDeliveryInfoResponse)(nil)).Elem()
}

func (o ExportDeliveryInfoResponsePtrOutput) ToExportDeliveryInfoResponsePtrOutput() ExportDeliveryInfoResponsePtrOutput {
	return o
}

func (o ExportDeliveryInfoResponsePtrOutput) ToExportDeliveryInfoResponsePtrOutputWithContext(ctx context.Context) ExportDeliveryInfoResponsePtrOutput {
	return o
}

func (o ExportDeliveryInfoResponsePtrOutput) Elem() ExportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v *ExportDeliveryInfoResponse) ExportDeliveryInfoResponse {
		if v != nil {
			return *v
		}
		var ret ExportDeliveryInfoResponse
		return ret
	}).(ExportDeliveryInfoResponseOutput)
}

func (o ExportDeliveryInfoResponsePtrOutput) Destination() ExportDeliveryDestinationResponsePtrOutput {
	return o.ApplyT(func(v *ExportDeliveryInfoResponse) *ExportDeliveryDestinationResponse {
		if v == nil {
			return nil
		}
		return &v.Destination
	}).(ExportDeliveryDestinationResponsePtrOutput)
}

type ExportRecurrencePeriod struct {
	From string  `pulumi:"from"`
	To   *string `pulumi:"to"`
}





type ExportRecurrencePeriodInput interface {
	pulumi.Input

	ToExportRecurrencePeriodOutput() ExportRecurrencePeriodOutput
	ToExportRecurrencePeriodOutputWithContext(context.Context) ExportRecurrencePeriodOutput
}

type ExportRecurrencePeriodArgs struct {
	From pulumi.StringInput    `pulumi:"from"`
	To   pulumi.StringPtrInput `pulumi:"to"`
}

func (ExportRecurrencePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportRecurrencePeriod)(nil)).Elem()
}

func (i ExportRecurrencePeriodArgs) ToExportRecurrencePeriodOutput() ExportRecurrencePeriodOutput {
	return i.ToExportRecurrencePeriodOutputWithContext(context.Background())
}

func (i ExportRecurrencePeriodArgs) ToExportRecurrencePeriodOutputWithContext(ctx context.Context) ExportRecurrencePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodOutput)
}

func (i ExportRecurrencePeriodArgs) ToExportRecurrencePeriodPtrOutput() ExportRecurrencePeriodPtrOutput {
	return i.ToExportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (i ExportRecurrencePeriodArgs) ToExportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodOutput).ToExportRecurrencePeriodPtrOutputWithContext(ctx)
}









type ExportRecurrencePeriodPtrInput interface {
	pulumi.Input

	ToExportRecurrencePeriodPtrOutput() ExportRecurrencePeriodPtrOutput
	ToExportRecurrencePeriodPtrOutputWithContext(context.Context) ExportRecurrencePeriodPtrOutput
}

type exportRecurrencePeriodPtrType ExportRecurrencePeriodArgs

func ExportRecurrencePeriodPtr(v *ExportRecurrencePeriodArgs) ExportRecurrencePeriodPtrInput {
	return (*exportRecurrencePeriodPtrType)(v)
}

func (*exportRecurrencePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportRecurrencePeriod)(nil)).Elem()
}

func (i *exportRecurrencePeriodPtrType) ToExportRecurrencePeriodPtrOutput() ExportRecurrencePeriodPtrOutput {
	return i.ToExportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (i *exportRecurrencePeriodPtrType) ToExportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodPtrOutput)
}

type ExportRecurrencePeriodOutput struct{ *pulumi.OutputState }

func (ExportRecurrencePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportRecurrencePeriod)(nil)).Elem()
}

func (o ExportRecurrencePeriodOutput) ToExportRecurrencePeriodOutput() ExportRecurrencePeriodOutput {
	return o
}

func (o ExportRecurrencePeriodOutput) ToExportRecurrencePeriodOutputWithContext(ctx context.Context) ExportRecurrencePeriodOutput {
	return o
}

func (o ExportRecurrencePeriodOutput) ToExportRecurrencePeriodPtrOutput() ExportRecurrencePeriodPtrOutput {
	return o.ToExportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (o ExportRecurrencePeriodOutput) ToExportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportRecurrencePeriod) *ExportRecurrencePeriod {
		return &v
	}).(ExportRecurrencePeriodPtrOutput)
}

func (o ExportRecurrencePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ExportRecurrencePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o ExportRecurrencePeriodOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportRecurrencePeriod) *string { return v.To }).(pulumi.StringPtrOutput)
}

type ExportRecurrencePeriodPtrOutput struct{ *pulumi.OutputState }

func (ExportRecurrencePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportRecurrencePeriod)(nil)).Elem()
}

func (o ExportRecurrencePeriodPtrOutput) ToExportRecurrencePeriodPtrOutput() ExportRecurrencePeriodPtrOutput {
	return o
}

func (o ExportRecurrencePeriodPtrOutput) ToExportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodPtrOutput {
	return o
}

func (o ExportRecurrencePeriodPtrOutput) Elem() ExportRecurrencePeriodOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriod) ExportRecurrencePeriod {
		if v != nil {
			return *v
		}
		var ret ExportRecurrencePeriod
		return ret
	}).(ExportRecurrencePeriodOutput)
}

func (o ExportRecurrencePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ExportRecurrencePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriod) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type ExportRecurrencePeriodResponse struct {
	From string  `pulumi:"from"`
	To   *string `pulumi:"to"`
}





type ExportRecurrencePeriodResponseInput interface {
	pulumi.Input

	ToExportRecurrencePeriodResponseOutput() ExportRecurrencePeriodResponseOutput
	ToExportRecurrencePeriodResponseOutputWithContext(context.Context) ExportRecurrencePeriodResponseOutput
}

type ExportRecurrencePeriodResponseArgs struct {
	From pulumi.StringInput    `pulumi:"from"`
	To   pulumi.StringPtrInput `pulumi:"to"`
}

func (ExportRecurrencePeriodResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportRecurrencePeriodResponse)(nil)).Elem()
}

func (i ExportRecurrencePeriodResponseArgs) ToExportRecurrencePeriodResponseOutput() ExportRecurrencePeriodResponseOutput {
	return i.ToExportRecurrencePeriodResponseOutputWithContext(context.Background())
}

func (i ExportRecurrencePeriodResponseArgs) ToExportRecurrencePeriodResponseOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodResponseOutput)
}

func (i ExportRecurrencePeriodResponseArgs) ToExportRecurrencePeriodResponsePtrOutput() ExportRecurrencePeriodResponsePtrOutput {
	return i.ToExportRecurrencePeriodResponsePtrOutputWithContext(context.Background())
}

func (i ExportRecurrencePeriodResponseArgs) ToExportRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodResponseOutput).ToExportRecurrencePeriodResponsePtrOutputWithContext(ctx)
}









type ExportRecurrencePeriodResponsePtrInput interface {
	pulumi.Input

	ToExportRecurrencePeriodResponsePtrOutput() ExportRecurrencePeriodResponsePtrOutput
	ToExportRecurrencePeriodResponsePtrOutputWithContext(context.Context) ExportRecurrencePeriodResponsePtrOutput
}

type exportRecurrencePeriodResponsePtrType ExportRecurrencePeriodResponseArgs

func ExportRecurrencePeriodResponsePtr(v *ExportRecurrencePeriodResponseArgs) ExportRecurrencePeriodResponsePtrInput {
	return (*exportRecurrencePeriodResponsePtrType)(v)
}

func (*exportRecurrencePeriodResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportRecurrencePeriodResponse)(nil)).Elem()
}

func (i *exportRecurrencePeriodResponsePtrType) ToExportRecurrencePeriodResponsePtrOutput() ExportRecurrencePeriodResponsePtrOutput {
	return i.ToExportRecurrencePeriodResponsePtrOutputWithContext(context.Background())
}

func (i *exportRecurrencePeriodResponsePtrType) ToExportRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportRecurrencePeriodResponsePtrOutput)
}

type ExportRecurrencePeriodResponseOutput struct{ *pulumi.OutputState }

func (ExportRecurrencePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportRecurrencePeriodResponse)(nil)).Elem()
}

func (o ExportRecurrencePeriodResponseOutput) ToExportRecurrencePeriodResponseOutput() ExportRecurrencePeriodResponseOutput {
	return o
}

func (o ExportRecurrencePeriodResponseOutput) ToExportRecurrencePeriodResponseOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponseOutput {
	return o
}

func (o ExportRecurrencePeriodResponseOutput) ToExportRecurrencePeriodResponsePtrOutput() ExportRecurrencePeriodResponsePtrOutput {
	return o.ToExportRecurrencePeriodResponsePtrOutputWithContext(context.Background())
}

func (o ExportRecurrencePeriodResponseOutput) ToExportRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportRecurrencePeriodResponse) *ExportRecurrencePeriodResponse {
		return &v
	}).(ExportRecurrencePeriodResponsePtrOutput)
}

func (o ExportRecurrencePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ExportRecurrencePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o ExportRecurrencePeriodResponseOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportRecurrencePeriodResponse) *string { return v.To }).(pulumi.StringPtrOutput)
}

type ExportRecurrencePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportRecurrencePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportRecurrencePeriodResponse)(nil)).Elem()
}

func (o ExportRecurrencePeriodResponsePtrOutput) ToExportRecurrencePeriodResponsePtrOutput() ExportRecurrencePeriodResponsePtrOutput {
	return o
}

func (o ExportRecurrencePeriodResponsePtrOutput) ToExportRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ExportRecurrencePeriodResponsePtrOutput {
	return o
}

func (o ExportRecurrencePeriodResponsePtrOutput) Elem() ExportRecurrencePeriodResponseOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriodResponse) ExportRecurrencePeriodResponse {
		if v != nil {
			return *v
		}
		var ret ExportRecurrencePeriodResponse
		return ret
	}).(ExportRecurrencePeriodResponseOutput)
}

func (o ExportRecurrencePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ExportRecurrencePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportRecurrencePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type ExportSchedule struct {
	Recurrence       string                  `pulumi:"recurrence"`
	RecurrencePeriod *ExportRecurrencePeriod `pulumi:"recurrencePeriod"`
	Status           *string                 `pulumi:"status"`
}





type ExportScheduleInput interface {
	pulumi.Input

	ToExportScheduleOutput() ExportScheduleOutput
	ToExportScheduleOutputWithContext(context.Context) ExportScheduleOutput
}

type ExportScheduleArgs struct {
	Recurrence       pulumi.StringInput             `pulumi:"recurrence"`
	RecurrencePeriod ExportRecurrencePeriodPtrInput `pulumi:"recurrencePeriod"`
	Status           pulumi.StringPtrInput          `pulumi:"status"`
}

func (ExportScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportSchedule)(nil)).Elem()
}

func (i ExportScheduleArgs) ToExportScheduleOutput() ExportScheduleOutput {
	return i.ToExportScheduleOutputWithContext(context.Background())
}

func (i ExportScheduleArgs) ToExportScheduleOutputWithContext(ctx context.Context) ExportScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportScheduleOutput)
}

func (i ExportScheduleArgs) ToExportSchedulePtrOutput() ExportSchedulePtrOutput {
	return i.ToExportSchedulePtrOutputWithContext(context.Background())
}

func (i ExportScheduleArgs) ToExportSchedulePtrOutputWithContext(ctx context.Context) ExportSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportScheduleOutput).ToExportSchedulePtrOutputWithContext(ctx)
}









type ExportSchedulePtrInput interface {
	pulumi.Input

	ToExportSchedulePtrOutput() ExportSchedulePtrOutput
	ToExportSchedulePtrOutputWithContext(context.Context) ExportSchedulePtrOutput
}

type exportSchedulePtrType ExportScheduleArgs

func ExportSchedulePtr(v *ExportScheduleArgs) ExportSchedulePtrInput {
	return (*exportSchedulePtrType)(v)
}

func (*exportSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportSchedule)(nil)).Elem()
}

func (i *exportSchedulePtrType) ToExportSchedulePtrOutput() ExportSchedulePtrOutput {
	return i.ToExportSchedulePtrOutputWithContext(context.Background())
}

func (i *exportSchedulePtrType) ToExportSchedulePtrOutputWithContext(ctx context.Context) ExportSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportSchedulePtrOutput)
}

type ExportScheduleOutput struct{ *pulumi.OutputState }

func (ExportScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportSchedule)(nil)).Elem()
}

func (o ExportScheduleOutput) ToExportScheduleOutput() ExportScheduleOutput {
	return o
}

func (o ExportScheduleOutput) ToExportScheduleOutputWithContext(ctx context.Context) ExportScheduleOutput {
	return o
}

func (o ExportScheduleOutput) ToExportSchedulePtrOutput() ExportSchedulePtrOutput {
	return o.ToExportSchedulePtrOutputWithContext(context.Background())
}

func (o ExportScheduleOutput) ToExportSchedulePtrOutputWithContext(ctx context.Context) ExportSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportSchedule) *ExportSchedule {
		return &v
	}).(ExportSchedulePtrOutput)
}

func (o ExportScheduleOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v ExportSchedule) string { return v.Recurrence }).(pulumi.StringOutput)
}

func (o ExportScheduleOutput) RecurrencePeriod() ExportRecurrencePeriodPtrOutput {
	return o.ApplyT(func(v ExportSchedule) *ExportRecurrencePeriod { return v.RecurrencePeriod }).(ExportRecurrencePeriodPtrOutput)
}

func (o ExportScheduleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportSchedule) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ExportSchedulePtrOutput struct{ *pulumi.OutputState }

func (ExportSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportSchedule)(nil)).Elem()
}

func (o ExportSchedulePtrOutput) ToExportSchedulePtrOutput() ExportSchedulePtrOutput {
	return o
}

func (o ExportSchedulePtrOutput) ToExportSchedulePtrOutputWithContext(ctx context.Context) ExportSchedulePtrOutput {
	return o
}

func (o ExportSchedulePtrOutput) Elem() ExportScheduleOutput {
	return o.ApplyT(func(v *ExportSchedule) ExportSchedule {
		if v != nil {
			return *v
		}
		var ret ExportSchedule
		return ret
	}).(ExportScheduleOutput)
}

func (o ExportSchedulePtrOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportSchedule) *string {
		if v == nil {
			return nil
		}
		return &v.Recurrence
	}).(pulumi.StringPtrOutput)
}

func (o ExportSchedulePtrOutput) RecurrencePeriod() ExportRecurrencePeriodPtrOutput {
	return o.ApplyT(func(v *ExportSchedule) *ExportRecurrencePeriod {
		if v == nil {
			return nil
		}
		return v.RecurrencePeriod
	}).(ExportRecurrencePeriodPtrOutput)
}

func (o ExportSchedulePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportSchedule) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ExportScheduleResponse struct {
	Recurrence       string                          `pulumi:"recurrence"`
	RecurrencePeriod *ExportRecurrencePeriodResponse `pulumi:"recurrencePeriod"`
	Status           *string                         `pulumi:"status"`
}





type ExportScheduleResponseInput interface {
	pulumi.Input

	ToExportScheduleResponseOutput() ExportScheduleResponseOutput
	ToExportScheduleResponseOutputWithContext(context.Context) ExportScheduleResponseOutput
}

type ExportScheduleResponseArgs struct {
	Recurrence       pulumi.StringInput                     `pulumi:"recurrence"`
	RecurrencePeriod ExportRecurrencePeriodResponsePtrInput `pulumi:"recurrencePeriod"`
	Status           pulumi.StringPtrInput                  `pulumi:"status"`
}

func (ExportScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportScheduleResponse)(nil)).Elem()
}

func (i ExportScheduleResponseArgs) ToExportScheduleResponseOutput() ExportScheduleResponseOutput {
	return i.ToExportScheduleResponseOutputWithContext(context.Background())
}

func (i ExportScheduleResponseArgs) ToExportScheduleResponseOutputWithContext(ctx context.Context) ExportScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportScheduleResponseOutput)
}

func (i ExportScheduleResponseArgs) ToExportScheduleResponsePtrOutput() ExportScheduleResponsePtrOutput {
	return i.ToExportScheduleResponsePtrOutputWithContext(context.Background())
}

func (i ExportScheduleResponseArgs) ToExportScheduleResponsePtrOutputWithContext(ctx context.Context) ExportScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportScheduleResponseOutput).ToExportScheduleResponsePtrOutputWithContext(ctx)
}









type ExportScheduleResponsePtrInput interface {
	pulumi.Input

	ToExportScheduleResponsePtrOutput() ExportScheduleResponsePtrOutput
	ToExportScheduleResponsePtrOutputWithContext(context.Context) ExportScheduleResponsePtrOutput
}

type exportScheduleResponsePtrType ExportScheduleResponseArgs

func ExportScheduleResponsePtr(v *ExportScheduleResponseArgs) ExportScheduleResponsePtrInput {
	return (*exportScheduleResponsePtrType)(v)
}

func (*exportScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportScheduleResponse)(nil)).Elem()
}

func (i *exportScheduleResponsePtrType) ToExportScheduleResponsePtrOutput() ExportScheduleResponsePtrOutput {
	return i.ToExportScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *exportScheduleResponsePtrType) ToExportScheduleResponsePtrOutputWithContext(ctx context.Context) ExportScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportScheduleResponsePtrOutput)
}

type ExportScheduleResponseOutput struct{ *pulumi.OutputState }

func (ExportScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportScheduleResponse)(nil)).Elem()
}

func (o ExportScheduleResponseOutput) ToExportScheduleResponseOutput() ExportScheduleResponseOutput {
	return o
}

func (o ExportScheduleResponseOutput) ToExportScheduleResponseOutputWithContext(ctx context.Context) ExportScheduleResponseOutput {
	return o
}

func (o ExportScheduleResponseOutput) ToExportScheduleResponsePtrOutput() ExportScheduleResponsePtrOutput {
	return o.ToExportScheduleResponsePtrOutputWithContext(context.Background())
}

func (o ExportScheduleResponseOutput) ToExportScheduleResponsePtrOutputWithContext(ctx context.Context) ExportScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportScheduleResponse) *ExportScheduleResponse {
		return &v
	}).(ExportScheduleResponsePtrOutput)
}

func (o ExportScheduleResponseOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v ExportScheduleResponse) string { return v.Recurrence }).(pulumi.StringOutput)
}

func (o ExportScheduleResponseOutput) RecurrencePeriod() ExportRecurrencePeriodResponsePtrOutput {
	return o.ApplyT(func(v ExportScheduleResponse) *ExportRecurrencePeriodResponse { return v.RecurrencePeriod }).(ExportRecurrencePeriodResponsePtrOutput)
}

func (o ExportScheduleResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportScheduleResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ExportScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportScheduleResponse)(nil)).Elem()
}

func (o ExportScheduleResponsePtrOutput) ToExportScheduleResponsePtrOutput() ExportScheduleResponsePtrOutput {
	return o
}

func (o ExportScheduleResponsePtrOutput) ToExportScheduleResponsePtrOutputWithContext(ctx context.Context) ExportScheduleResponsePtrOutput {
	return o
}

func (o ExportScheduleResponsePtrOutput) Elem() ExportScheduleResponseOutput {
	return o.ApplyT(func(v *ExportScheduleResponse) ExportScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ExportScheduleResponse
		return ret
	}).(ExportScheduleResponseOutput)
}

func (o ExportScheduleResponsePtrOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Recurrence
	}).(pulumi.StringPtrOutput)
}

func (o ExportScheduleResponsePtrOutput) RecurrencePeriod() ExportRecurrencePeriodResponsePtrOutput {
	return o.ApplyT(func(v *ExportScheduleResponse) *ExportRecurrencePeriodResponse {
		if v == nil {
			return nil
		}
		return v.RecurrencePeriod
	}).(ExportRecurrencePeriodResponsePtrOutput)
}

func (o ExportScheduleResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type KpiProperties struct {
	Enabled *bool   `pulumi:"enabled"`
	Id      *string `pulumi:"id"`
	Type    *string `pulumi:"type"`
}





type KpiPropertiesInput interface {
	pulumi.Input

	ToKpiPropertiesOutput() KpiPropertiesOutput
	ToKpiPropertiesOutputWithContext(context.Context) KpiPropertiesOutput
}

type KpiPropertiesArgs struct {
	Enabled pulumi.BoolPtrInput   `pulumi:"enabled"`
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Type    pulumi.StringPtrInput `pulumi:"type"`
}

func (KpiPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiProperties)(nil)).Elem()
}

func (i KpiPropertiesArgs) ToKpiPropertiesOutput() KpiPropertiesOutput {
	return i.ToKpiPropertiesOutputWithContext(context.Background())
}

func (i KpiPropertiesArgs) ToKpiPropertiesOutputWithContext(ctx context.Context) KpiPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiPropertiesOutput)
}





type KpiPropertiesArrayInput interface {
	pulumi.Input

	ToKpiPropertiesArrayOutput() KpiPropertiesArrayOutput
	ToKpiPropertiesArrayOutputWithContext(context.Context) KpiPropertiesArrayOutput
}

type KpiPropertiesArray []KpiPropertiesInput

func (KpiPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiProperties)(nil)).Elem()
}

func (i KpiPropertiesArray) ToKpiPropertiesArrayOutput() KpiPropertiesArrayOutput {
	return i.ToKpiPropertiesArrayOutputWithContext(context.Background())
}

func (i KpiPropertiesArray) ToKpiPropertiesArrayOutputWithContext(ctx context.Context) KpiPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiPropertiesArrayOutput)
}

type KpiPropertiesOutput struct{ *pulumi.OutputState }

func (KpiPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiProperties)(nil)).Elem()
}

func (o KpiPropertiesOutput) ToKpiPropertiesOutput() KpiPropertiesOutput {
	return o
}

func (o KpiPropertiesOutput) ToKpiPropertiesOutputWithContext(ctx context.Context) KpiPropertiesOutput {
	return o
}

func (o KpiPropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KpiProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o KpiPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KpiProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KpiPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KpiProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type KpiPropertiesArrayOutput struct{ *pulumi.OutputState }

func (KpiPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiProperties)(nil)).Elem()
}

func (o KpiPropertiesArrayOutput) ToKpiPropertiesArrayOutput() KpiPropertiesArrayOutput {
	return o
}

func (o KpiPropertiesArrayOutput) ToKpiPropertiesArrayOutputWithContext(ctx context.Context) KpiPropertiesArrayOutput {
	return o
}

func (o KpiPropertiesArrayOutput) Index(i pulumi.IntInput) KpiPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KpiProperties {
		return vs[0].([]KpiProperties)[vs[1].(int)]
	}).(KpiPropertiesOutput)
}

type KpiPropertiesResponse struct {
	Enabled *bool   `pulumi:"enabled"`
	Id      *string `pulumi:"id"`
	Type    *string `pulumi:"type"`
}





type KpiPropertiesResponseInput interface {
	pulumi.Input

	ToKpiPropertiesResponseOutput() KpiPropertiesResponseOutput
	ToKpiPropertiesResponseOutputWithContext(context.Context) KpiPropertiesResponseOutput
}

type KpiPropertiesResponseArgs struct {
	Enabled pulumi.BoolPtrInput   `pulumi:"enabled"`
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Type    pulumi.StringPtrInput `pulumi:"type"`
}

func (KpiPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiPropertiesResponse)(nil)).Elem()
}

func (i KpiPropertiesResponseArgs) ToKpiPropertiesResponseOutput() KpiPropertiesResponseOutput {
	return i.ToKpiPropertiesResponseOutputWithContext(context.Background())
}

func (i KpiPropertiesResponseArgs) ToKpiPropertiesResponseOutputWithContext(ctx context.Context) KpiPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiPropertiesResponseOutput)
}





type KpiPropertiesResponseArrayInput interface {
	pulumi.Input

	ToKpiPropertiesResponseArrayOutput() KpiPropertiesResponseArrayOutput
	ToKpiPropertiesResponseArrayOutputWithContext(context.Context) KpiPropertiesResponseArrayOutput
}

type KpiPropertiesResponseArray []KpiPropertiesResponseInput

func (KpiPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiPropertiesResponse)(nil)).Elem()
}

func (i KpiPropertiesResponseArray) ToKpiPropertiesResponseArrayOutput() KpiPropertiesResponseArrayOutput {
	return i.ToKpiPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i KpiPropertiesResponseArray) ToKpiPropertiesResponseArrayOutputWithContext(ctx context.Context) KpiPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KpiPropertiesResponseArrayOutput)
}

type KpiPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KpiPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiPropertiesResponse)(nil)).Elem()
}

func (o KpiPropertiesResponseOutput) ToKpiPropertiesResponseOutput() KpiPropertiesResponseOutput {
	return o
}

func (o KpiPropertiesResponseOutput) ToKpiPropertiesResponseOutputWithContext(ctx context.Context) KpiPropertiesResponseOutput {
	return o
}

func (o KpiPropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KpiPropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o KpiPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KpiPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KpiPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KpiPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type KpiPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (KpiPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KpiPropertiesResponse)(nil)).Elem()
}

func (o KpiPropertiesResponseArrayOutput) ToKpiPropertiesResponseArrayOutput() KpiPropertiesResponseArrayOutput {
	return o
}

func (o KpiPropertiesResponseArrayOutput) ToKpiPropertiesResponseArrayOutputWithContext(ctx context.Context) KpiPropertiesResponseArrayOutput {
	return o
}

func (o KpiPropertiesResponseArrayOutput) Index(i pulumi.IntInput) KpiPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KpiPropertiesResponse {
		return vs[0].([]KpiPropertiesResponse)[vs[1].(int)]
	}).(KpiPropertiesResponseOutput)
}

type PivotProperties struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type PivotPropertiesInput interface {
	pulumi.Input

	ToPivotPropertiesOutput() PivotPropertiesOutput
	ToPivotPropertiesOutputWithContext(context.Context) PivotPropertiesOutput
}

type PivotPropertiesArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PivotPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PivotProperties)(nil)).Elem()
}

func (i PivotPropertiesArgs) ToPivotPropertiesOutput() PivotPropertiesOutput {
	return i.ToPivotPropertiesOutputWithContext(context.Background())
}

func (i PivotPropertiesArgs) ToPivotPropertiesOutputWithContext(ctx context.Context) PivotPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PivotPropertiesOutput)
}





type PivotPropertiesArrayInput interface {
	pulumi.Input

	ToPivotPropertiesArrayOutput() PivotPropertiesArrayOutput
	ToPivotPropertiesArrayOutputWithContext(context.Context) PivotPropertiesArrayOutput
}

type PivotPropertiesArray []PivotPropertiesInput

func (PivotPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PivotProperties)(nil)).Elem()
}

func (i PivotPropertiesArray) ToPivotPropertiesArrayOutput() PivotPropertiesArrayOutput {
	return i.ToPivotPropertiesArrayOutputWithContext(context.Background())
}

func (i PivotPropertiesArray) ToPivotPropertiesArrayOutputWithContext(ctx context.Context) PivotPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PivotPropertiesArrayOutput)
}

type PivotPropertiesOutput struct{ *pulumi.OutputState }

func (PivotPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PivotProperties)(nil)).Elem()
}

func (o PivotPropertiesOutput) ToPivotPropertiesOutput() PivotPropertiesOutput {
	return o
}

func (o PivotPropertiesOutput) ToPivotPropertiesOutputWithContext(ctx context.Context) PivotPropertiesOutput {
	return o
}

func (o PivotPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PivotProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PivotPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PivotProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PivotPropertiesArrayOutput struct{ *pulumi.OutputState }

func (PivotPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PivotProperties)(nil)).Elem()
}

func (o PivotPropertiesArrayOutput) ToPivotPropertiesArrayOutput() PivotPropertiesArrayOutput {
	return o
}

func (o PivotPropertiesArrayOutput) ToPivotPropertiesArrayOutputWithContext(ctx context.Context) PivotPropertiesArrayOutput {
	return o
}

func (o PivotPropertiesArrayOutput) Index(i pulumi.IntInput) PivotPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PivotProperties {
		return vs[0].([]PivotProperties)[vs[1].(int)]
	}).(PivotPropertiesOutput)
}

type PivotPropertiesResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type PivotPropertiesResponseInput interface {
	pulumi.Input

	ToPivotPropertiesResponseOutput() PivotPropertiesResponseOutput
	ToPivotPropertiesResponseOutputWithContext(context.Context) PivotPropertiesResponseOutput
}

type PivotPropertiesResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PivotPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PivotPropertiesResponse)(nil)).Elem()
}

func (i PivotPropertiesResponseArgs) ToPivotPropertiesResponseOutput() PivotPropertiesResponseOutput {
	return i.ToPivotPropertiesResponseOutputWithContext(context.Background())
}

func (i PivotPropertiesResponseArgs) ToPivotPropertiesResponseOutputWithContext(ctx context.Context) PivotPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PivotPropertiesResponseOutput)
}





type PivotPropertiesResponseArrayInput interface {
	pulumi.Input

	ToPivotPropertiesResponseArrayOutput() PivotPropertiesResponseArrayOutput
	ToPivotPropertiesResponseArrayOutputWithContext(context.Context) PivotPropertiesResponseArrayOutput
}

type PivotPropertiesResponseArray []PivotPropertiesResponseInput

func (PivotPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PivotPropertiesResponse)(nil)).Elem()
}

func (i PivotPropertiesResponseArray) ToPivotPropertiesResponseArrayOutput() PivotPropertiesResponseArrayOutput {
	return i.ToPivotPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i PivotPropertiesResponseArray) ToPivotPropertiesResponseArrayOutputWithContext(ctx context.Context) PivotPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PivotPropertiesResponseArrayOutput)
}

type PivotPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PivotPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PivotPropertiesResponse)(nil)).Elem()
}

func (o PivotPropertiesResponseOutput) ToPivotPropertiesResponseOutput() PivotPropertiesResponseOutput {
	return o
}

func (o PivotPropertiesResponseOutput) ToPivotPropertiesResponseOutputWithContext(ctx context.Context) PivotPropertiesResponseOutput {
	return o
}

func (o PivotPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PivotPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PivotPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PivotPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PivotPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (PivotPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PivotPropertiesResponse)(nil)).Elem()
}

func (o PivotPropertiesResponseArrayOutput) ToPivotPropertiesResponseArrayOutput() PivotPropertiesResponseArrayOutput {
	return o
}

func (o PivotPropertiesResponseArrayOutput) ToPivotPropertiesResponseArrayOutputWithContext(ctx context.Context) PivotPropertiesResponseArrayOutput {
	return o
}

func (o PivotPropertiesResponseArrayOutput) Index(i pulumi.IntInput) PivotPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PivotPropertiesResponse {
		return vs[0].([]PivotPropertiesResponse)[vs[1].(int)]
	}).(PivotPropertiesResponseOutput)
}

type QueryAggregation struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}





type QueryAggregationInput interface {
	pulumi.Input

	ToQueryAggregationOutput() QueryAggregationOutput
	ToQueryAggregationOutputWithContext(context.Context) QueryAggregationOutput
}

type QueryAggregationArgs struct {
	Function pulumi.StringInput `pulumi:"function"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (QueryAggregationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryAggregation)(nil)).Elem()
}

func (i QueryAggregationArgs) ToQueryAggregationOutput() QueryAggregationOutput {
	return i.ToQueryAggregationOutputWithContext(context.Background())
}

func (i QueryAggregationArgs) ToQueryAggregationOutputWithContext(ctx context.Context) QueryAggregationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryAggregationOutput)
}





type QueryAggregationMapInput interface {
	pulumi.Input

	ToQueryAggregationMapOutput() QueryAggregationMapOutput
	ToQueryAggregationMapOutputWithContext(context.Context) QueryAggregationMapOutput
}

type QueryAggregationMap map[string]QueryAggregationInput

func (QueryAggregationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QueryAggregation)(nil)).Elem()
}

func (i QueryAggregationMap) ToQueryAggregationMapOutput() QueryAggregationMapOutput {
	return i.ToQueryAggregationMapOutputWithContext(context.Background())
}

func (i QueryAggregationMap) ToQueryAggregationMapOutputWithContext(ctx context.Context) QueryAggregationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryAggregationMapOutput)
}

type QueryAggregationOutput struct{ *pulumi.OutputState }

func (QueryAggregationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryAggregation)(nil)).Elem()
}

func (o QueryAggregationOutput) ToQueryAggregationOutput() QueryAggregationOutput {
	return o
}

func (o QueryAggregationOutput) ToQueryAggregationOutputWithContext(ctx context.Context) QueryAggregationOutput {
	return o
}

func (o QueryAggregationOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v QueryAggregation) string { return v.Function }).(pulumi.StringOutput)
}

func (o QueryAggregationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryAggregation) string { return v.Name }).(pulumi.StringOutput)
}

type QueryAggregationMapOutput struct{ *pulumi.OutputState }

func (QueryAggregationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QueryAggregation)(nil)).Elem()
}

func (o QueryAggregationMapOutput) ToQueryAggregationMapOutput() QueryAggregationMapOutput {
	return o
}

func (o QueryAggregationMapOutput) ToQueryAggregationMapOutputWithContext(ctx context.Context) QueryAggregationMapOutput {
	return o
}

func (o QueryAggregationMapOutput) MapIndex(k pulumi.StringInput) QueryAggregationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QueryAggregation {
		return vs[0].(map[string]QueryAggregation)[vs[1].(string)]
	}).(QueryAggregationOutput)
}

type QueryAggregationResponse struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}





type QueryAggregationResponseInput interface {
	pulumi.Input

	ToQueryAggregationResponseOutput() QueryAggregationResponseOutput
	ToQueryAggregationResponseOutputWithContext(context.Context) QueryAggregationResponseOutput
}

type QueryAggregationResponseArgs struct {
	Function pulumi.StringInput `pulumi:"function"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (QueryAggregationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryAggregationResponse)(nil)).Elem()
}

func (i QueryAggregationResponseArgs) ToQueryAggregationResponseOutput() QueryAggregationResponseOutput {
	return i.ToQueryAggregationResponseOutputWithContext(context.Background())
}

func (i QueryAggregationResponseArgs) ToQueryAggregationResponseOutputWithContext(ctx context.Context) QueryAggregationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryAggregationResponseOutput)
}





type QueryAggregationResponseMapInput interface {
	pulumi.Input

	ToQueryAggregationResponseMapOutput() QueryAggregationResponseMapOutput
	ToQueryAggregationResponseMapOutputWithContext(context.Context) QueryAggregationResponseMapOutput
}

type QueryAggregationResponseMap map[string]QueryAggregationResponseInput

func (QueryAggregationResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QueryAggregationResponse)(nil)).Elem()
}

func (i QueryAggregationResponseMap) ToQueryAggregationResponseMapOutput() QueryAggregationResponseMapOutput {
	return i.ToQueryAggregationResponseMapOutputWithContext(context.Background())
}

func (i QueryAggregationResponseMap) ToQueryAggregationResponseMapOutputWithContext(ctx context.Context) QueryAggregationResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryAggregationResponseMapOutput)
}

type QueryAggregationResponseOutput struct{ *pulumi.OutputState }

func (QueryAggregationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryAggregationResponse)(nil)).Elem()
}

func (o QueryAggregationResponseOutput) ToQueryAggregationResponseOutput() QueryAggregationResponseOutput {
	return o
}

func (o QueryAggregationResponseOutput) ToQueryAggregationResponseOutputWithContext(ctx context.Context) QueryAggregationResponseOutput {
	return o
}

func (o QueryAggregationResponseOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v QueryAggregationResponse) string { return v.Function }).(pulumi.StringOutput)
}

func (o QueryAggregationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryAggregationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type QueryAggregationResponseMapOutput struct{ *pulumi.OutputState }

func (QueryAggregationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QueryAggregationResponse)(nil)).Elem()
}

func (o QueryAggregationResponseMapOutput) ToQueryAggregationResponseMapOutput() QueryAggregationResponseMapOutput {
	return o
}

func (o QueryAggregationResponseMapOutput) ToQueryAggregationResponseMapOutputWithContext(ctx context.Context) QueryAggregationResponseMapOutput {
	return o
}

func (o QueryAggregationResponseMapOutput) MapIndex(k pulumi.StringInput) QueryAggregationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QueryAggregationResponse {
		return vs[0].(map[string]QueryAggregationResponse)[vs[1].(string)]
	}).(QueryAggregationResponseOutput)
}

type QueryComparisonExpression struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type QueryComparisonExpressionInput interface {
	pulumi.Input

	ToQueryComparisonExpressionOutput() QueryComparisonExpressionOutput
	ToQueryComparisonExpressionOutputWithContext(context.Context) QueryComparisonExpressionOutput
}

type QueryComparisonExpressionArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (QueryComparisonExpressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryComparisonExpression)(nil)).Elem()
}

func (i QueryComparisonExpressionArgs) ToQueryComparisonExpressionOutput() QueryComparisonExpressionOutput {
	return i.ToQueryComparisonExpressionOutputWithContext(context.Background())
}

func (i QueryComparisonExpressionArgs) ToQueryComparisonExpressionOutputWithContext(ctx context.Context) QueryComparisonExpressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionOutput)
}

func (i QueryComparisonExpressionArgs) ToQueryComparisonExpressionPtrOutput() QueryComparisonExpressionPtrOutput {
	return i.ToQueryComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i QueryComparisonExpressionArgs) ToQueryComparisonExpressionPtrOutputWithContext(ctx context.Context) QueryComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionOutput).ToQueryComparisonExpressionPtrOutputWithContext(ctx)
}









type QueryComparisonExpressionPtrInput interface {
	pulumi.Input

	ToQueryComparisonExpressionPtrOutput() QueryComparisonExpressionPtrOutput
	ToQueryComparisonExpressionPtrOutputWithContext(context.Context) QueryComparisonExpressionPtrOutput
}

type queryComparisonExpressionPtrType QueryComparisonExpressionArgs

func QueryComparisonExpressionPtr(v *QueryComparisonExpressionArgs) QueryComparisonExpressionPtrInput {
	return (*queryComparisonExpressionPtrType)(v)
}

func (*queryComparisonExpressionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryComparisonExpression)(nil)).Elem()
}

func (i *queryComparisonExpressionPtrType) ToQueryComparisonExpressionPtrOutput() QueryComparisonExpressionPtrOutput {
	return i.ToQueryComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i *queryComparisonExpressionPtrType) ToQueryComparisonExpressionPtrOutputWithContext(ctx context.Context) QueryComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionPtrOutput)
}

type QueryComparisonExpressionOutput struct{ *pulumi.OutputState }

func (QueryComparisonExpressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryComparisonExpression)(nil)).Elem()
}

func (o QueryComparisonExpressionOutput) ToQueryComparisonExpressionOutput() QueryComparisonExpressionOutput {
	return o
}

func (o QueryComparisonExpressionOutput) ToQueryComparisonExpressionOutputWithContext(ctx context.Context) QueryComparisonExpressionOutput {
	return o
}

func (o QueryComparisonExpressionOutput) ToQueryComparisonExpressionPtrOutput() QueryComparisonExpressionPtrOutput {
	return o.ToQueryComparisonExpressionPtrOutputWithContext(context.Background())
}

func (o QueryComparisonExpressionOutput) ToQueryComparisonExpressionPtrOutputWithContext(ctx context.Context) QueryComparisonExpressionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryComparisonExpression) *QueryComparisonExpression {
		return &v
	}).(QueryComparisonExpressionPtrOutput)
}

func (o QueryComparisonExpressionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryComparisonExpression) string { return v.Name }).(pulumi.StringOutput)
}

func (o QueryComparisonExpressionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v QueryComparisonExpression) string { return v.Operator }).(pulumi.StringOutput)
}

func (o QueryComparisonExpressionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryComparisonExpression) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type QueryComparisonExpressionPtrOutput struct{ *pulumi.OutputState }

func (QueryComparisonExpressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryComparisonExpression)(nil)).Elem()
}

func (o QueryComparisonExpressionPtrOutput) ToQueryComparisonExpressionPtrOutput() QueryComparisonExpressionPtrOutput {
	return o
}

func (o QueryComparisonExpressionPtrOutput) ToQueryComparisonExpressionPtrOutputWithContext(ctx context.Context) QueryComparisonExpressionPtrOutput {
	return o
}

func (o QueryComparisonExpressionPtrOutput) Elem() QueryComparisonExpressionOutput {
	return o.ApplyT(func(v *QueryComparisonExpression) QueryComparisonExpression {
		if v != nil {
			return *v
		}
		var ret QueryComparisonExpression
		return ret
	}).(QueryComparisonExpressionOutput)
}

func (o QueryComparisonExpressionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o QueryComparisonExpressionPtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o QueryComparisonExpressionPtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueryComparisonExpression) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type QueryComparisonExpressionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type QueryComparisonExpressionResponseInput interface {
	pulumi.Input

	ToQueryComparisonExpressionResponseOutput() QueryComparisonExpressionResponseOutput
	ToQueryComparisonExpressionResponseOutputWithContext(context.Context) QueryComparisonExpressionResponseOutput
}

type QueryComparisonExpressionResponseArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (QueryComparisonExpressionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryComparisonExpressionResponse)(nil)).Elem()
}

func (i QueryComparisonExpressionResponseArgs) ToQueryComparisonExpressionResponseOutput() QueryComparisonExpressionResponseOutput {
	return i.ToQueryComparisonExpressionResponseOutputWithContext(context.Background())
}

func (i QueryComparisonExpressionResponseArgs) ToQueryComparisonExpressionResponseOutputWithContext(ctx context.Context) QueryComparisonExpressionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionResponseOutput)
}

func (i QueryComparisonExpressionResponseArgs) ToQueryComparisonExpressionResponsePtrOutput() QueryComparisonExpressionResponsePtrOutput {
	return i.ToQueryComparisonExpressionResponsePtrOutputWithContext(context.Background())
}

func (i QueryComparisonExpressionResponseArgs) ToQueryComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) QueryComparisonExpressionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionResponseOutput).ToQueryComparisonExpressionResponsePtrOutputWithContext(ctx)
}









type QueryComparisonExpressionResponsePtrInput interface {
	pulumi.Input

	ToQueryComparisonExpressionResponsePtrOutput() QueryComparisonExpressionResponsePtrOutput
	ToQueryComparisonExpressionResponsePtrOutputWithContext(context.Context) QueryComparisonExpressionResponsePtrOutput
}

type queryComparisonExpressionResponsePtrType QueryComparisonExpressionResponseArgs

func QueryComparisonExpressionResponsePtr(v *QueryComparisonExpressionResponseArgs) QueryComparisonExpressionResponsePtrInput {
	return (*queryComparisonExpressionResponsePtrType)(v)
}

func (*queryComparisonExpressionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryComparisonExpressionResponse)(nil)).Elem()
}

func (i *queryComparisonExpressionResponsePtrType) ToQueryComparisonExpressionResponsePtrOutput() QueryComparisonExpressionResponsePtrOutput {
	return i.ToQueryComparisonExpressionResponsePtrOutputWithContext(context.Background())
}

func (i *queryComparisonExpressionResponsePtrType) ToQueryComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) QueryComparisonExpressionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryComparisonExpressionResponsePtrOutput)
}

type QueryComparisonExpressionResponseOutput struct{ *pulumi.OutputState }

func (QueryComparisonExpressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryComparisonExpressionResponse)(nil)).Elem()
}

func (o QueryComparisonExpressionResponseOutput) ToQueryComparisonExpressionResponseOutput() QueryComparisonExpressionResponseOutput {
	return o
}

func (o QueryComparisonExpressionResponseOutput) ToQueryComparisonExpressionResponseOutputWithContext(ctx context.Context) QueryComparisonExpressionResponseOutput {
	return o
}

func (o QueryComparisonExpressionResponseOutput) ToQueryComparisonExpressionResponsePtrOutput() QueryComparisonExpressionResponsePtrOutput {
	return o.ToQueryComparisonExpressionResponsePtrOutputWithContext(context.Background())
}

func (o QueryComparisonExpressionResponseOutput) ToQueryComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) QueryComparisonExpressionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryComparisonExpressionResponse) *QueryComparisonExpressionResponse {
		return &v
	}).(QueryComparisonExpressionResponsePtrOutput)
}

func (o QueryComparisonExpressionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryComparisonExpressionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o QueryComparisonExpressionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v QueryComparisonExpressionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o QueryComparisonExpressionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryComparisonExpressionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type QueryComparisonExpressionResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryComparisonExpressionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryComparisonExpressionResponse)(nil)).Elem()
}

func (o QueryComparisonExpressionResponsePtrOutput) ToQueryComparisonExpressionResponsePtrOutput() QueryComparisonExpressionResponsePtrOutput {
	return o
}

func (o QueryComparisonExpressionResponsePtrOutput) ToQueryComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) QueryComparisonExpressionResponsePtrOutput {
	return o
}

func (o QueryComparisonExpressionResponsePtrOutput) Elem() QueryComparisonExpressionResponseOutput {
	return o.ApplyT(func(v *QueryComparisonExpressionResponse) QueryComparisonExpressionResponse {
		if v != nil {
			return *v
		}
		var ret QueryComparisonExpressionResponse
		return ret
	}).(QueryComparisonExpressionResponseOutput)
}

func (o QueryComparisonExpressionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o QueryComparisonExpressionResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o QueryComparisonExpressionResponsePtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueryComparisonExpressionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type QueryDataset struct {
	Aggregation   map[string]QueryAggregation `pulumi:"aggregation"`
	Configuration *QueryDatasetConfiguration  `pulumi:"configuration"`
	Filter        *QueryFilter                `pulumi:"filter"`
	Granularity   *string                     `pulumi:"granularity"`
	Grouping      []QueryGrouping             `pulumi:"grouping"`
}





type QueryDatasetInput interface {
	pulumi.Input

	ToQueryDatasetOutput() QueryDatasetOutput
	ToQueryDatasetOutputWithContext(context.Context) QueryDatasetOutput
}

type QueryDatasetArgs struct {
	Aggregation   QueryAggregationMapInput          `pulumi:"aggregation"`
	Configuration QueryDatasetConfigurationPtrInput `pulumi:"configuration"`
	Filter        QueryFilterPtrInput               `pulumi:"filter"`
	Granularity   pulumi.StringPtrInput             `pulumi:"granularity"`
	Grouping      QueryGroupingArrayInput           `pulumi:"grouping"`
}

func (QueryDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDataset)(nil)).Elem()
}

func (i QueryDatasetArgs) ToQueryDatasetOutput() QueryDatasetOutput {
	return i.ToQueryDatasetOutputWithContext(context.Background())
}

func (i QueryDatasetArgs) ToQueryDatasetOutputWithContext(ctx context.Context) QueryDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetOutput)
}

func (i QueryDatasetArgs) ToQueryDatasetPtrOutput() QueryDatasetPtrOutput {
	return i.ToQueryDatasetPtrOutputWithContext(context.Background())
}

func (i QueryDatasetArgs) ToQueryDatasetPtrOutputWithContext(ctx context.Context) QueryDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetOutput).ToQueryDatasetPtrOutputWithContext(ctx)
}









type QueryDatasetPtrInput interface {
	pulumi.Input

	ToQueryDatasetPtrOutput() QueryDatasetPtrOutput
	ToQueryDatasetPtrOutputWithContext(context.Context) QueryDatasetPtrOutput
}

type queryDatasetPtrType QueryDatasetArgs

func QueryDatasetPtr(v *QueryDatasetArgs) QueryDatasetPtrInput {
	return (*queryDatasetPtrType)(v)
}

func (*queryDatasetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDataset)(nil)).Elem()
}

func (i *queryDatasetPtrType) ToQueryDatasetPtrOutput() QueryDatasetPtrOutput {
	return i.ToQueryDatasetPtrOutputWithContext(context.Background())
}

func (i *queryDatasetPtrType) ToQueryDatasetPtrOutputWithContext(ctx context.Context) QueryDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetPtrOutput)
}

type QueryDatasetOutput struct{ *pulumi.OutputState }

func (QueryDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDataset)(nil)).Elem()
}

func (o QueryDatasetOutput) ToQueryDatasetOutput() QueryDatasetOutput {
	return o
}

func (o QueryDatasetOutput) ToQueryDatasetOutputWithContext(ctx context.Context) QueryDatasetOutput {
	return o
}

func (o QueryDatasetOutput) ToQueryDatasetPtrOutput() QueryDatasetPtrOutput {
	return o.ToQueryDatasetPtrOutputWithContext(context.Background())
}

func (o QueryDatasetOutput) ToQueryDatasetPtrOutputWithContext(ctx context.Context) QueryDatasetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDataset) *QueryDataset {
		return &v
	}).(QueryDatasetPtrOutput)
}

func (o QueryDatasetOutput) Aggregation() QueryAggregationMapOutput {
	return o.ApplyT(func(v QueryDataset) map[string]QueryAggregation { return v.Aggregation }).(QueryAggregationMapOutput)
}

func (o QueryDatasetOutput) Configuration() QueryDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v QueryDataset) *QueryDatasetConfiguration { return v.Configuration }).(QueryDatasetConfigurationPtrOutput)
}

func (o QueryDatasetOutput) Filter() QueryFilterPtrOutput {
	return o.ApplyT(func(v QueryDataset) *QueryFilter { return v.Filter }).(QueryFilterPtrOutput)
}

func (o QueryDatasetOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueryDataset) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o QueryDatasetOutput) Grouping() QueryGroupingArrayOutput {
	return o.ApplyT(func(v QueryDataset) []QueryGrouping { return v.Grouping }).(QueryGroupingArrayOutput)
}

type QueryDatasetPtrOutput struct{ *pulumi.OutputState }

func (QueryDatasetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDataset)(nil)).Elem()
}

func (o QueryDatasetPtrOutput) ToQueryDatasetPtrOutput() QueryDatasetPtrOutput {
	return o
}

func (o QueryDatasetPtrOutput) ToQueryDatasetPtrOutputWithContext(ctx context.Context) QueryDatasetPtrOutput {
	return o
}

func (o QueryDatasetPtrOutput) Elem() QueryDatasetOutput {
	return o.ApplyT(func(v *QueryDataset) QueryDataset {
		if v != nil {
			return *v
		}
		var ret QueryDataset
		return ret
	}).(QueryDatasetOutput)
}

func (o QueryDatasetPtrOutput) Aggregation() QueryAggregationMapOutput {
	return o.ApplyT(func(v *QueryDataset) map[string]QueryAggregation {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(QueryAggregationMapOutput)
}

func (o QueryDatasetPtrOutput) Configuration() QueryDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v *QueryDataset) *QueryDatasetConfiguration {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(QueryDatasetConfigurationPtrOutput)
}

func (o QueryDatasetPtrOutput) Filter() QueryFilterPtrOutput {
	return o.ApplyT(func(v *QueryDataset) *QueryFilter {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(QueryFilterPtrOutput)
}

func (o QueryDatasetPtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryDataset) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o QueryDatasetPtrOutput) Grouping() QueryGroupingArrayOutput {
	return o.ApplyT(func(v *QueryDataset) []QueryGrouping {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(QueryGroupingArrayOutput)
}

type QueryDatasetConfiguration struct {
	Columns []string `pulumi:"columns"`
}





type QueryDatasetConfigurationInput interface {
	pulumi.Input

	ToQueryDatasetConfigurationOutput() QueryDatasetConfigurationOutput
	ToQueryDatasetConfigurationOutputWithContext(context.Context) QueryDatasetConfigurationOutput
}

type QueryDatasetConfigurationArgs struct {
	Columns pulumi.StringArrayInput `pulumi:"columns"`
}

func (QueryDatasetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetConfiguration)(nil)).Elem()
}

func (i QueryDatasetConfigurationArgs) ToQueryDatasetConfigurationOutput() QueryDatasetConfigurationOutput {
	return i.ToQueryDatasetConfigurationOutputWithContext(context.Background())
}

func (i QueryDatasetConfigurationArgs) ToQueryDatasetConfigurationOutputWithContext(ctx context.Context) QueryDatasetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationOutput)
}

func (i QueryDatasetConfigurationArgs) ToQueryDatasetConfigurationPtrOutput() QueryDatasetConfigurationPtrOutput {
	return i.ToQueryDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i QueryDatasetConfigurationArgs) ToQueryDatasetConfigurationPtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationOutput).ToQueryDatasetConfigurationPtrOutputWithContext(ctx)
}









type QueryDatasetConfigurationPtrInput interface {
	pulumi.Input

	ToQueryDatasetConfigurationPtrOutput() QueryDatasetConfigurationPtrOutput
	ToQueryDatasetConfigurationPtrOutputWithContext(context.Context) QueryDatasetConfigurationPtrOutput
}

type queryDatasetConfigurationPtrType QueryDatasetConfigurationArgs

func QueryDatasetConfigurationPtr(v *QueryDatasetConfigurationArgs) QueryDatasetConfigurationPtrInput {
	return (*queryDatasetConfigurationPtrType)(v)
}

func (*queryDatasetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetConfiguration)(nil)).Elem()
}

func (i *queryDatasetConfigurationPtrType) ToQueryDatasetConfigurationPtrOutput() QueryDatasetConfigurationPtrOutput {
	return i.ToQueryDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i *queryDatasetConfigurationPtrType) ToQueryDatasetConfigurationPtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationPtrOutput)
}

type QueryDatasetConfigurationOutput struct{ *pulumi.OutputState }

func (QueryDatasetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetConfiguration)(nil)).Elem()
}

func (o QueryDatasetConfigurationOutput) ToQueryDatasetConfigurationOutput() QueryDatasetConfigurationOutput {
	return o
}

func (o QueryDatasetConfigurationOutput) ToQueryDatasetConfigurationOutputWithContext(ctx context.Context) QueryDatasetConfigurationOutput {
	return o
}

func (o QueryDatasetConfigurationOutput) ToQueryDatasetConfigurationPtrOutput() QueryDatasetConfigurationPtrOutput {
	return o.ToQueryDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (o QueryDatasetConfigurationOutput) ToQueryDatasetConfigurationPtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDatasetConfiguration) *QueryDatasetConfiguration {
		return &v
	}).(QueryDatasetConfigurationPtrOutput)
}

func (o QueryDatasetConfigurationOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryDatasetConfiguration) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type QueryDatasetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (QueryDatasetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetConfiguration)(nil)).Elem()
}

func (o QueryDatasetConfigurationPtrOutput) ToQueryDatasetConfigurationPtrOutput() QueryDatasetConfigurationPtrOutput {
	return o
}

func (o QueryDatasetConfigurationPtrOutput) ToQueryDatasetConfigurationPtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationPtrOutput {
	return o
}

func (o QueryDatasetConfigurationPtrOutput) Elem() QueryDatasetConfigurationOutput {
	return o.ApplyT(func(v *QueryDatasetConfiguration) QueryDatasetConfiguration {
		if v != nil {
			return *v
		}
		var ret QueryDatasetConfiguration
		return ret
	}).(QueryDatasetConfigurationOutput)
}

func (o QueryDatasetConfigurationPtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueryDatasetConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type QueryDatasetConfigurationResponse struct {
	Columns []string `pulumi:"columns"`
}





type QueryDatasetConfigurationResponseInput interface {
	pulumi.Input

	ToQueryDatasetConfigurationResponseOutput() QueryDatasetConfigurationResponseOutput
	ToQueryDatasetConfigurationResponseOutputWithContext(context.Context) QueryDatasetConfigurationResponseOutput
}

type QueryDatasetConfigurationResponseArgs struct {
	Columns pulumi.StringArrayInput `pulumi:"columns"`
}

func (QueryDatasetConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetConfigurationResponse)(nil)).Elem()
}

func (i QueryDatasetConfigurationResponseArgs) ToQueryDatasetConfigurationResponseOutput() QueryDatasetConfigurationResponseOutput {
	return i.ToQueryDatasetConfigurationResponseOutputWithContext(context.Background())
}

func (i QueryDatasetConfigurationResponseArgs) ToQueryDatasetConfigurationResponseOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationResponseOutput)
}

func (i QueryDatasetConfigurationResponseArgs) ToQueryDatasetConfigurationResponsePtrOutput() QueryDatasetConfigurationResponsePtrOutput {
	return i.ToQueryDatasetConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i QueryDatasetConfigurationResponseArgs) ToQueryDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationResponseOutput).ToQueryDatasetConfigurationResponsePtrOutputWithContext(ctx)
}









type QueryDatasetConfigurationResponsePtrInput interface {
	pulumi.Input

	ToQueryDatasetConfigurationResponsePtrOutput() QueryDatasetConfigurationResponsePtrOutput
	ToQueryDatasetConfigurationResponsePtrOutputWithContext(context.Context) QueryDatasetConfigurationResponsePtrOutput
}

type queryDatasetConfigurationResponsePtrType QueryDatasetConfigurationResponseArgs

func QueryDatasetConfigurationResponsePtr(v *QueryDatasetConfigurationResponseArgs) QueryDatasetConfigurationResponsePtrInput {
	return (*queryDatasetConfigurationResponsePtrType)(v)
}

func (*queryDatasetConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetConfigurationResponse)(nil)).Elem()
}

func (i *queryDatasetConfigurationResponsePtrType) ToQueryDatasetConfigurationResponsePtrOutput() QueryDatasetConfigurationResponsePtrOutput {
	return i.ToQueryDatasetConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *queryDatasetConfigurationResponsePtrType) ToQueryDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetConfigurationResponsePtrOutput)
}

type QueryDatasetConfigurationResponseOutput struct{ *pulumi.OutputState }

func (QueryDatasetConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetConfigurationResponse)(nil)).Elem()
}

func (o QueryDatasetConfigurationResponseOutput) ToQueryDatasetConfigurationResponseOutput() QueryDatasetConfigurationResponseOutput {
	return o
}

func (o QueryDatasetConfigurationResponseOutput) ToQueryDatasetConfigurationResponseOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponseOutput {
	return o
}

func (o QueryDatasetConfigurationResponseOutput) ToQueryDatasetConfigurationResponsePtrOutput() QueryDatasetConfigurationResponsePtrOutput {
	return o.ToQueryDatasetConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o QueryDatasetConfigurationResponseOutput) ToQueryDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDatasetConfigurationResponse) *QueryDatasetConfigurationResponse {
		return &v
	}).(QueryDatasetConfigurationResponsePtrOutput)
}

func (o QueryDatasetConfigurationResponseOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v QueryDatasetConfigurationResponse) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type QueryDatasetConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryDatasetConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetConfigurationResponse)(nil)).Elem()
}

func (o QueryDatasetConfigurationResponsePtrOutput) ToQueryDatasetConfigurationResponsePtrOutput() QueryDatasetConfigurationResponsePtrOutput {
	return o
}

func (o QueryDatasetConfigurationResponsePtrOutput) ToQueryDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) QueryDatasetConfigurationResponsePtrOutput {
	return o
}

func (o QueryDatasetConfigurationResponsePtrOutput) Elem() QueryDatasetConfigurationResponseOutput {
	return o.ApplyT(func(v *QueryDatasetConfigurationResponse) QueryDatasetConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret QueryDatasetConfigurationResponse
		return ret
	}).(QueryDatasetConfigurationResponseOutput)
}

func (o QueryDatasetConfigurationResponsePtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *QueryDatasetConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type QueryDatasetResponse struct {
	Aggregation   map[string]QueryAggregationResponse `pulumi:"aggregation"`
	Configuration *QueryDatasetConfigurationResponse  `pulumi:"configuration"`
	Filter        *QueryFilterResponse                `pulumi:"filter"`
	Granularity   *string                             `pulumi:"granularity"`
	Grouping      []QueryGroupingResponse             `pulumi:"grouping"`
}





type QueryDatasetResponseInput interface {
	pulumi.Input

	ToQueryDatasetResponseOutput() QueryDatasetResponseOutput
	ToQueryDatasetResponseOutputWithContext(context.Context) QueryDatasetResponseOutput
}

type QueryDatasetResponseArgs struct {
	Aggregation   QueryAggregationResponseMapInput          `pulumi:"aggregation"`
	Configuration QueryDatasetConfigurationResponsePtrInput `pulumi:"configuration"`
	Filter        QueryFilterResponsePtrInput               `pulumi:"filter"`
	Granularity   pulumi.StringPtrInput                     `pulumi:"granularity"`
	Grouping      QueryGroupingResponseArrayInput           `pulumi:"grouping"`
}

func (QueryDatasetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetResponse)(nil)).Elem()
}

func (i QueryDatasetResponseArgs) ToQueryDatasetResponseOutput() QueryDatasetResponseOutput {
	return i.ToQueryDatasetResponseOutputWithContext(context.Background())
}

func (i QueryDatasetResponseArgs) ToQueryDatasetResponseOutputWithContext(ctx context.Context) QueryDatasetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetResponseOutput)
}

func (i QueryDatasetResponseArgs) ToQueryDatasetResponsePtrOutput() QueryDatasetResponsePtrOutput {
	return i.ToQueryDatasetResponsePtrOutputWithContext(context.Background())
}

func (i QueryDatasetResponseArgs) ToQueryDatasetResponsePtrOutputWithContext(ctx context.Context) QueryDatasetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetResponseOutput).ToQueryDatasetResponsePtrOutputWithContext(ctx)
}









type QueryDatasetResponsePtrInput interface {
	pulumi.Input

	ToQueryDatasetResponsePtrOutput() QueryDatasetResponsePtrOutput
	ToQueryDatasetResponsePtrOutputWithContext(context.Context) QueryDatasetResponsePtrOutput
}

type queryDatasetResponsePtrType QueryDatasetResponseArgs

func QueryDatasetResponsePtr(v *QueryDatasetResponseArgs) QueryDatasetResponsePtrInput {
	return (*queryDatasetResponsePtrType)(v)
}

func (*queryDatasetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetResponse)(nil)).Elem()
}

func (i *queryDatasetResponsePtrType) ToQueryDatasetResponsePtrOutput() QueryDatasetResponsePtrOutput {
	return i.ToQueryDatasetResponsePtrOutputWithContext(context.Background())
}

func (i *queryDatasetResponsePtrType) ToQueryDatasetResponsePtrOutputWithContext(ctx context.Context) QueryDatasetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryDatasetResponsePtrOutput)
}

type QueryDatasetResponseOutput struct{ *pulumi.OutputState }

func (QueryDatasetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryDatasetResponse)(nil)).Elem()
}

func (o QueryDatasetResponseOutput) ToQueryDatasetResponseOutput() QueryDatasetResponseOutput {
	return o
}

func (o QueryDatasetResponseOutput) ToQueryDatasetResponseOutputWithContext(ctx context.Context) QueryDatasetResponseOutput {
	return o
}

func (o QueryDatasetResponseOutput) ToQueryDatasetResponsePtrOutput() QueryDatasetResponsePtrOutput {
	return o.ToQueryDatasetResponsePtrOutputWithContext(context.Background())
}

func (o QueryDatasetResponseOutput) ToQueryDatasetResponsePtrOutputWithContext(ctx context.Context) QueryDatasetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryDatasetResponse) *QueryDatasetResponse {
		return &v
	}).(QueryDatasetResponsePtrOutput)
}

func (o QueryDatasetResponseOutput) Aggregation() QueryAggregationResponseMapOutput {
	return o.ApplyT(func(v QueryDatasetResponse) map[string]QueryAggregationResponse { return v.Aggregation }).(QueryAggregationResponseMapOutput)
}

func (o QueryDatasetResponseOutput) Configuration() QueryDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v QueryDatasetResponse) *QueryDatasetConfigurationResponse { return v.Configuration }).(QueryDatasetConfigurationResponsePtrOutput)
}

func (o QueryDatasetResponseOutput) Filter() QueryFilterResponsePtrOutput {
	return o.ApplyT(func(v QueryDatasetResponse) *QueryFilterResponse { return v.Filter }).(QueryFilterResponsePtrOutput)
}

func (o QueryDatasetResponseOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueryDatasetResponse) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o QueryDatasetResponseOutput) Grouping() QueryGroupingResponseArrayOutput {
	return o.ApplyT(func(v QueryDatasetResponse) []QueryGroupingResponse { return v.Grouping }).(QueryGroupingResponseArrayOutput)
}

type QueryDatasetResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryDatasetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryDatasetResponse)(nil)).Elem()
}

func (o QueryDatasetResponsePtrOutput) ToQueryDatasetResponsePtrOutput() QueryDatasetResponsePtrOutput {
	return o
}

func (o QueryDatasetResponsePtrOutput) ToQueryDatasetResponsePtrOutputWithContext(ctx context.Context) QueryDatasetResponsePtrOutput {
	return o
}

func (o QueryDatasetResponsePtrOutput) Elem() QueryDatasetResponseOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) QueryDatasetResponse {
		if v != nil {
			return *v
		}
		var ret QueryDatasetResponse
		return ret
	}).(QueryDatasetResponseOutput)
}

func (o QueryDatasetResponsePtrOutput) Aggregation() QueryAggregationResponseMapOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) map[string]QueryAggregationResponse {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(QueryAggregationResponseMapOutput)
}

func (o QueryDatasetResponsePtrOutput) Configuration() QueryDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) *QueryDatasetConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(QueryDatasetConfigurationResponsePtrOutput)
}

func (o QueryDatasetResponsePtrOutput) Filter() QueryFilterResponsePtrOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) *QueryFilterResponse {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(QueryFilterResponsePtrOutput)
}

func (o QueryDatasetResponsePtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o QueryDatasetResponsePtrOutput) Grouping() QueryGroupingResponseArrayOutput {
	return o.ApplyT(func(v *QueryDatasetResponse) []QueryGroupingResponse {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(QueryGroupingResponseArrayOutput)
}

type QueryFilter struct {
	And        []QueryFilter              `pulumi:"and"`
	Dimensions *QueryComparisonExpression `pulumi:"dimensions"`
	Or         []QueryFilter              `pulumi:"or"`
	Tags       *QueryComparisonExpression `pulumi:"tags"`
}





type QueryFilterInput interface {
	pulumi.Input

	ToQueryFilterOutput() QueryFilterOutput
	ToQueryFilterOutputWithContext(context.Context) QueryFilterOutput
}

type QueryFilterArgs struct {
	And        QueryFilterArrayInput             `pulumi:"and"`
	Dimensions QueryComparisonExpressionPtrInput `pulumi:"dimensions"`
	Or         QueryFilterArrayInput             `pulumi:"or"`
	Tags       QueryComparisonExpressionPtrInput `pulumi:"tags"`
}

func (QueryFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryFilter)(nil)).Elem()
}

func (i QueryFilterArgs) ToQueryFilterOutput() QueryFilterOutput {
	return i.ToQueryFilterOutputWithContext(context.Background())
}

func (i QueryFilterArgs) ToQueryFilterOutputWithContext(ctx context.Context) QueryFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterOutput)
}

func (i QueryFilterArgs) ToQueryFilterPtrOutput() QueryFilterPtrOutput {
	return i.ToQueryFilterPtrOutputWithContext(context.Background())
}

func (i QueryFilterArgs) ToQueryFilterPtrOutputWithContext(ctx context.Context) QueryFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterOutput).ToQueryFilterPtrOutputWithContext(ctx)
}









type QueryFilterPtrInput interface {
	pulumi.Input

	ToQueryFilterPtrOutput() QueryFilterPtrOutput
	ToQueryFilterPtrOutputWithContext(context.Context) QueryFilterPtrOutput
}

type queryFilterPtrType QueryFilterArgs

func QueryFilterPtr(v *QueryFilterArgs) QueryFilterPtrInput {
	return (*queryFilterPtrType)(v)
}

func (*queryFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryFilter)(nil)).Elem()
}

func (i *queryFilterPtrType) ToQueryFilterPtrOutput() QueryFilterPtrOutput {
	return i.ToQueryFilterPtrOutputWithContext(context.Background())
}

func (i *queryFilterPtrType) ToQueryFilterPtrOutputWithContext(ctx context.Context) QueryFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterPtrOutput)
}





type QueryFilterArrayInput interface {
	pulumi.Input

	ToQueryFilterArrayOutput() QueryFilterArrayOutput
	ToQueryFilterArrayOutputWithContext(context.Context) QueryFilterArrayOutput
}

type QueryFilterArray []QueryFilterInput

func (QueryFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryFilter)(nil)).Elem()
}

func (i QueryFilterArray) ToQueryFilterArrayOutput() QueryFilterArrayOutput {
	return i.ToQueryFilterArrayOutputWithContext(context.Background())
}

func (i QueryFilterArray) ToQueryFilterArrayOutputWithContext(ctx context.Context) QueryFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterArrayOutput)
}

type QueryFilterOutput struct{ *pulumi.OutputState }

func (QueryFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryFilter)(nil)).Elem()
}

func (o QueryFilterOutput) ToQueryFilterOutput() QueryFilterOutput {
	return o
}

func (o QueryFilterOutput) ToQueryFilterOutputWithContext(ctx context.Context) QueryFilterOutput {
	return o
}

func (o QueryFilterOutput) ToQueryFilterPtrOutput() QueryFilterPtrOutput {
	return o.ToQueryFilterPtrOutputWithContext(context.Background())
}

func (o QueryFilterOutput) ToQueryFilterPtrOutputWithContext(ctx context.Context) QueryFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryFilter) *QueryFilter {
		return &v
	}).(QueryFilterPtrOutput)
}

func (o QueryFilterOutput) And() QueryFilterArrayOutput {
	return o.ApplyT(func(v QueryFilter) []QueryFilter { return v.And }).(QueryFilterArrayOutput)
}

func (o QueryFilterOutput) Dimensions() QueryComparisonExpressionPtrOutput {
	return o.ApplyT(func(v QueryFilter) *QueryComparisonExpression { return v.Dimensions }).(QueryComparisonExpressionPtrOutput)
}

func (o QueryFilterOutput) Or() QueryFilterArrayOutput {
	return o.ApplyT(func(v QueryFilter) []QueryFilter { return v.Or }).(QueryFilterArrayOutput)
}

func (o QueryFilterOutput) Tags() QueryComparisonExpressionPtrOutput {
	return o.ApplyT(func(v QueryFilter) *QueryComparisonExpression { return v.Tags }).(QueryComparisonExpressionPtrOutput)
}

type QueryFilterPtrOutput struct{ *pulumi.OutputState }

func (QueryFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryFilter)(nil)).Elem()
}

func (o QueryFilterPtrOutput) ToQueryFilterPtrOutput() QueryFilterPtrOutput {
	return o
}

func (o QueryFilterPtrOutput) ToQueryFilterPtrOutputWithContext(ctx context.Context) QueryFilterPtrOutput {
	return o
}

func (o QueryFilterPtrOutput) Elem() QueryFilterOutput {
	return o.ApplyT(func(v *QueryFilter) QueryFilter {
		if v != nil {
			return *v
		}
		var ret QueryFilter
		return ret
	}).(QueryFilterOutput)
}

func (o QueryFilterPtrOutput) And() QueryFilterArrayOutput {
	return o.ApplyT(func(v *QueryFilter) []QueryFilter {
		if v == nil {
			return nil
		}
		return v.And
	}).(QueryFilterArrayOutput)
}

func (o QueryFilterPtrOutput) Dimensions() QueryComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *QueryFilter) *QueryComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Dimensions
	}).(QueryComparisonExpressionPtrOutput)
}

func (o QueryFilterPtrOutput) Or() QueryFilterArrayOutput {
	return o.ApplyT(func(v *QueryFilter) []QueryFilter {
		if v == nil {
			return nil
		}
		return v.Or
	}).(QueryFilterArrayOutput)
}

func (o QueryFilterPtrOutput) Tags() QueryComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *QueryFilter) *QueryComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(QueryComparisonExpressionPtrOutput)
}

type QueryFilterArrayOutput struct{ *pulumi.OutputState }

func (QueryFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryFilter)(nil)).Elem()
}

func (o QueryFilterArrayOutput) ToQueryFilterArrayOutput() QueryFilterArrayOutput {
	return o
}

func (o QueryFilterArrayOutput) ToQueryFilterArrayOutputWithContext(ctx context.Context) QueryFilterArrayOutput {
	return o
}

func (o QueryFilterArrayOutput) Index(i pulumi.IntInput) QueryFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueryFilter {
		return vs[0].([]QueryFilter)[vs[1].(int)]
	}).(QueryFilterOutput)
}

type QueryFilterResponse struct {
	And        []QueryFilterResponse              `pulumi:"and"`
	Dimensions *QueryComparisonExpressionResponse `pulumi:"dimensions"`
	Or         []QueryFilterResponse              `pulumi:"or"`
	Tags       *QueryComparisonExpressionResponse `pulumi:"tags"`
}





type QueryFilterResponseInput interface {
	pulumi.Input

	ToQueryFilterResponseOutput() QueryFilterResponseOutput
	ToQueryFilterResponseOutputWithContext(context.Context) QueryFilterResponseOutput
}

type QueryFilterResponseArgs struct {
	And        QueryFilterResponseArrayInput             `pulumi:"and"`
	Dimensions QueryComparisonExpressionResponsePtrInput `pulumi:"dimensions"`
	Or         QueryFilterResponseArrayInput             `pulumi:"or"`
	Tags       QueryComparisonExpressionResponsePtrInput `pulumi:"tags"`
}

func (QueryFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryFilterResponse)(nil)).Elem()
}

func (i QueryFilterResponseArgs) ToQueryFilterResponseOutput() QueryFilterResponseOutput {
	return i.ToQueryFilterResponseOutputWithContext(context.Background())
}

func (i QueryFilterResponseArgs) ToQueryFilterResponseOutputWithContext(ctx context.Context) QueryFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterResponseOutput)
}

func (i QueryFilterResponseArgs) ToQueryFilterResponsePtrOutput() QueryFilterResponsePtrOutput {
	return i.ToQueryFilterResponsePtrOutputWithContext(context.Background())
}

func (i QueryFilterResponseArgs) ToQueryFilterResponsePtrOutputWithContext(ctx context.Context) QueryFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterResponseOutput).ToQueryFilterResponsePtrOutputWithContext(ctx)
}









type QueryFilterResponsePtrInput interface {
	pulumi.Input

	ToQueryFilterResponsePtrOutput() QueryFilterResponsePtrOutput
	ToQueryFilterResponsePtrOutputWithContext(context.Context) QueryFilterResponsePtrOutput
}

type queryFilterResponsePtrType QueryFilterResponseArgs

func QueryFilterResponsePtr(v *QueryFilterResponseArgs) QueryFilterResponsePtrInput {
	return (*queryFilterResponsePtrType)(v)
}

func (*queryFilterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryFilterResponse)(nil)).Elem()
}

func (i *queryFilterResponsePtrType) ToQueryFilterResponsePtrOutput() QueryFilterResponsePtrOutput {
	return i.ToQueryFilterResponsePtrOutputWithContext(context.Background())
}

func (i *queryFilterResponsePtrType) ToQueryFilterResponsePtrOutputWithContext(ctx context.Context) QueryFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterResponsePtrOutput)
}





type QueryFilterResponseArrayInput interface {
	pulumi.Input

	ToQueryFilterResponseArrayOutput() QueryFilterResponseArrayOutput
	ToQueryFilterResponseArrayOutputWithContext(context.Context) QueryFilterResponseArrayOutput
}

type QueryFilterResponseArray []QueryFilterResponseInput

func (QueryFilterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryFilterResponse)(nil)).Elem()
}

func (i QueryFilterResponseArray) ToQueryFilterResponseArrayOutput() QueryFilterResponseArrayOutput {
	return i.ToQueryFilterResponseArrayOutputWithContext(context.Background())
}

func (i QueryFilterResponseArray) ToQueryFilterResponseArrayOutputWithContext(ctx context.Context) QueryFilterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryFilterResponseArrayOutput)
}

type QueryFilterResponseOutput struct{ *pulumi.OutputState }

func (QueryFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryFilterResponse)(nil)).Elem()
}

func (o QueryFilterResponseOutput) ToQueryFilterResponseOutput() QueryFilterResponseOutput {
	return o
}

func (o QueryFilterResponseOutput) ToQueryFilterResponseOutputWithContext(ctx context.Context) QueryFilterResponseOutput {
	return o
}

func (o QueryFilterResponseOutput) ToQueryFilterResponsePtrOutput() QueryFilterResponsePtrOutput {
	return o.ToQueryFilterResponsePtrOutputWithContext(context.Background())
}

func (o QueryFilterResponseOutput) ToQueryFilterResponsePtrOutputWithContext(ctx context.Context) QueryFilterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryFilterResponse) *QueryFilterResponse {
		return &v
	}).(QueryFilterResponsePtrOutput)
}

func (o QueryFilterResponseOutput) And() QueryFilterResponseArrayOutput {
	return o.ApplyT(func(v QueryFilterResponse) []QueryFilterResponse { return v.And }).(QueryFilterResponseArrayOutput)
}

func (o QueryFilterResponseOutput) Dimensions() QueryComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v QueryFilterResponse) *QueryComparisonExpressionResponse { return v.Dimensions }).(QueryComparisonExpressionResponsePtrOutput)
}

func (o QueryFilterResponseOutput) Or() QueryFilterResponseArrayOutput {
	return o.ApplyT(func(v QueryFilterResponse) []QueryFilterResponse { return v.Or }).(QueryFilterResponseArrayOutput)
}

func (o QueryFilterResponseOutput) Tags() QueryComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v QueryFilterResponse) *QueryComparisonExpressionResponse { return v.Tags }).(QueryComparisonExpressionResponsePtrOutput)
}

type QueryFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryFilterResponse)(nil)).Elem()
}

func (o QueryFilterResponsePtrOutput) ToQueryFilterResponsePtrOutput() QueryFilterResponsePtrOutput {
	return o
}

func (o QueryFilterResponsePtrOutput) ToQueryFilterResponsePtrOutputWithContext(ctx context.Context) QueryFilterResponsePtrOutput {
	return o
}

func (o QueryFilterResponsePtrOutput) Elem() QueryFilterResponseOutput {
	return o.ApplyT(func(v *QueryFilterResponse) QueryFilterResponse {
		if v != nil {
			return *v
		}
		var ret QueryFilterResponse
		return ret
	}).(QueryFilterResponseOutput)
}

func (o QueryFilterResponsePtrOutput) And() QueryFilterResponseArrayOutput {
	return o.ApplyT(func(v *QueryFilterResponse) []QueryFilterResponse {
		if v == nil {
			return nil
		}
		return v.And
	}).(QueryFilterResponseArrayOutput)
}

func (o QueryFilterResponsePtrOutput) Dimensions() QueryComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *QueryFilterResponse) *QueryComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Dimensions
	}).(QueryComparisonExpressionResponsePtrOutput)
}

func (o QueryFilterResponsePtrOutput) Or() QueryFilterResponseArrayOutput {
	return o.ApplyT(func(v *QueryFilterResponse) []QueryFilterResponse {
		if v == nil {
			return nil
		}
		return v.Or
	}).(QueryFilterResponseArrayOutput)
}

func (o QueryFilterResponsePtrOutput) Tags() QueryComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *QueryFilterResponse) *QueryComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(QueryComparisonExpressionResponsePtrOutput)
}

type QueryFilterResponseArrayOutput struct{ *pulumi.OutputState }

func (QueryFilterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryFilterResponse)(nil)).Elem()
}

func (o QueryFilterResponseArrayOutput) ToQueryFilterResponseArrayOutput() QueryFilterResponseArrayOutput {
	return o
}

func (o QueryFilterResponseArrayOutput) ToQueryFilterResponseArrayOutputWithContext(ctx context.Context) QueryFilterResponseArrayOutput {
	return o
}

func (o QueryFilterResponseArrayOutput) Index(i pulumi.IntInput) QueryFilterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueryFilterResponse {
		return vs[0].([]QueryFilterResponse)[vs[1].(int)]
	}).(QueryFilterResponseOutput)
}

type QueryGrouping struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type QueryGroupingInput interface {
	pulumi.Input

	ToQueryGroupingOutput() QueryGroupingOutput
	ToQueryGroupingOutputWithContext(context.Context) QueryGroupingOutput
}

type QueryGroupingArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (QueryGroupingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryGrouping)(nil)).Elem()
}

func (i QueryGroupingArgs) ToQueryGroupingOutput() QueryGroupingOutput {
	return i.ToQueryGroupingOutputWithContext(context.Background())
}

func (i QueryGroupingArgs) ToQueryGroupingOutputWithContext(ctx context.Context) QueryGroupingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryGroupingOutput)
}





type QueryGroupingArrayInput interface {
	pulumi.Input

	ToQueryGroupingArrayOutput() QueryGroupingArrayOutput
	ToQueryGroupingArrayOutputWithContext(context.Context) QueryGroupingArrayOutput
}

type QueryGroupingArray []QueryGroupingInput

func (QueryGroupingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryGrouping)(nil)).Elem()
}

func (i QueryGroupingArray) ToQueryGroupingArrayOutput() QueryGroupingArrayOutput {
	return i.ToQueryGroupingArrayOutputWithContext(context.Background())
}

func (i QueryGroupingArray) ToQueryGroupingArrayOutputWithContext(ctx context.Context) QueryGroupingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryGroupingArrayOutput)
}

type QueryGroupingOutput struct{ *pulumi.OutputState }

func (QueryGroupingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryGrouping)(nil)).Elem()
}

func (o QueryGroupingOutput) ToQueryGroupingOutput() QueryGroupingOutput {
	return o
}

func (o QueryGroupingOutput) ToQueryGroupingOutputWithContext(ctx context.Context) QueryGroupingOutput {
	return o
}

func (o QueryGroupingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryGrouping) string { return v.Name }).(pulumi.StringOutput)
}

func (o QueryGroupingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v QueryGrouping) string { return v.Type }).(pulumi.StringOutput)
}

type QueryGroupingArrayOutput struct{ *pulumi.OutputState }

func (QueryGroupingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryGrouping)(nil)).Elem()
}

func (o QueryGroupingArrayOutput) ToQueryGroupingArrayOutput() QueryGroupingArrayOutput {
	return o
}

func (o QueryGroupingArrayOutput) ToQueryGroupingArrayOutputWithContext(ctx context.Context) QueryGroupingArrayOutput {
	return o
}

func (o QueryGroupingArrayOutput) Index(i pulumi.IntInput) QueryGroupingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueryGrouping {
		return vs[0].([]QueryGrouping)[vs[1].(int)]
	}).(QueryGroupingOutput)
}

type QueryGroupingResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type QueryGroupingResponseInput interface {
	pulumi.Input

	ToQueryGroupingResponseOutput() QueryGroupingResponseOutput
	ToQueryGroupingResponseOutputWithContext(context.Context) QueryGroupingResponseOutput
}

type QueryGroupingResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (QueryGroupingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryGroupingResponse)(nil)).Elem()
}

func (i QueryGroupingResponseArgs) ToQueryGroupingResponseOutput() QueryGroupingResponseOutput {
	return i.ToQueryGroupingResponseOutputWithContext(context.Background())
}

func (i QueryGroupingResponseArgs) ToQueryGroupingResponseOutputWithContext(ctx context.Context) QueryGroupingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryGroupingResponseOutput)
}





type QueryGroupingResponseArrayInput interface {
	pulumi.Input

	ToQueryGroupingResponseArrayOutput() QueryGroupingResponseArrayOutput
	ToQueryGroupingResponseArrayOutputWithContext(context.Context) QueryGroupingResponseArrayOutput
}

type QueryGroupingResponseArray []QueryGroupingResponseInput

func (QueryGroupingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryGroupingResponse)(nil)).Elem()
}

func (i QueryGroupingResponseArray) ToQueryGroupingResponseArrayOutput() QueryGroupingResponseArrayOutput {
	return i.ToQueryGroupingResponseArrayOutputWithContext(context.Background())
}

func (i QueryGroupingResponseArray) ToQueryGroupingResponseArrayOutputWithContext(ctx context.Context) QueryGroupingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryGroupingResponseArrayOutput)
}

type QueryGroupingResponseOutput struct{ *pulumi.OutputState }

func (QueryGroupingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryGroupingResponse)(nil)).Elem()
}

func (o QueryGroupingResponseOutput) ToQueryGroupingResponseOutput() QueryGroupingResponseOutput {
	return o
}

func (o QueryGroupingResponseOutput) ToQueryGroupingResponseOutputWithContext(ctx context.Context) QueryGroupingResponseOutput {
	return o
}

func (o QueryGroupingResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v QueryGroupingResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o QueryGroupingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v QueryGroupingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type QueryGroupingResponseArrayOutput struct{ *pulumi.OutputState }

func (QueryGroupingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QueryGroupingResponse)(nil)).Elem()
}

func (o QueryGroupingResponseArrayOutput) ToQueryGroupingResponseArrayOutput() QueryGroupingResponseArrayOutput {
	return o
}

func (o QueryGroupingResponseArrayOutput) ToQueryGroupingResponseArrayOutputWithContext(ctx context.Context) QueryGroupingResponseArrayOutput {
	return o
}

func (o QueryGroupingResponseArrayOutput) Index(i pulumi.IntInput) QueryGroupingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QueryGroupingResponse {
		return vs[0].([]QueryGroupingResponse)[vs[1].(int)]
	}).(QueryGroupingResponseOutput)
}

type QueryTimePeriod struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}





type QueryTimePeriodInput interface {
	pulumi.Input

	ToQueryTimePeriodOutput() QueryTimePeriodOutput
	ToQueryTimePeriodOutputWithContext(context.Context) QueryTimePeriodOutput
}

type QueryTimePeriodArgs struct {
	From pulumi.StringInput `pulumi:"from"`
	To   pulumi.StringInput `pulumi:"to"`
}

func (QueryTimePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryTimePeriod)(nil)).Elem()
}

func (i QueryTimePeriodArgs) ToQueryTimePeriodOutput() QueryTimePeriodOutput {
	return i.ToQueryTimePeriodOutputWithContext(context.Background())
}

func (i QueryTimePeriodArgs) ToQueryTimePeriodOutputWithContext(ctx context.Context) QueryTimePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodOutput)
}

func (i QueryTimePeriodArgs) ToQueryTimePeriodPtrOutput() QueryTimePeriodPtrOutput {
	return i.ToQueryTimePeriodPtrOutputWithContext(context.Background())
}

func (i QueryTimePeriodArgs) ToQueryTimePeriodPtrOutputWithContext(ctx context.Context) QueryTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodOutput).ToQueryTimePeriodPtrOutputWithContext(ctx)
}









type QueryTimePeriodPtrInput interface {
	pulumi.Input

	ToQueryTimePeriodPtrOutput() QueryTimePeriodPtrOutput
	ToQueryTimePeriodPtrOutputWithContext(context.Context) QueryTimePeriodPtrOutput
}

type queryTimePeriodPtrType QueryTimePeriodArgs

func QueryTimePeriodPtr(v *QueryTimePeriodArgs) QueryTimePeriodPtrInput {
	return (*queryTimePeriodPtrType)(v)
}

func (*queryTimePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryTimePeriod)(nil)).Elem()
}

func (i *queryTimePeriodPtrType) ToQueryTimePeriodPtrOutput() QueryTimePeriodPtrOutput {
	return i.ToQueryTimePeriodPtrOutputWithContext(context.Background())
}

func (i *queryTimePeriodPtrType) ToQueryTimePeriodPtrOutputWithContext(ctx context.Context) QueryTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodPtrOutput)
}

type QueryTimePeriodOutput struct{ *pulumi.OutputState }

func (QueryTimePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryTimePeriod)(nil)).Elem()
}

func (o QueryTimePeriodOutput) ToQueryTimePeriodOutput() QueryTimePeriodOutput {
	return o
}

func (o QueryTimePeriodOutput) ToQueryTimePeriodOutputWithContext(ctx context.Context) QueryTimePeriodOutput {
	return o
}

func (o QueryTimePeriodOutput) ToQueryTimePeriodPtrOutput() QueryTimePeriodPtrOutput {
	return o.ToQueryTimePeriodPtrOutputWithContext(context.Background())
}

func (o QueryTimePeriodOutput) ToQueryTimePeriodPtrOutputWithContext(ctx context.Context) QueryTimePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryTimePeriod) *QueryTimePeriod {
		return &v
	}).(QueryTimePeriodPtrOutput)
}

func (o QueryTimePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v QueryTimePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o QueryTimePeriodOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v QueryTimePeriod) string { return v.To }).(pulumi.StringOutput)
}

type QueryTimePeriodPtrOutput struct{ *pulumi.OutputState }

func (QueryTimePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryTimePeriod)(nil)).Elem()
}

func (o QueryTimePeriodPtrOutput) ToQueryTimePeriodPtrOutput() QueryTimePeriodPtrOutput {
	return o
}

func (o QueryTimePeriodPtrOutput) ToQueryTimePeriodPtrOutputWithContext(ctx context.Context) QueryTimePeriodPtrOutput {
	return o
}

func (o QueryTimePeriodPtrOutput) Elem() QueryTimePeriodOutput {
	return o.ApplyT(func(v *QueryTimePeriod) QueryTimePeriod {
		if v != nil {
			return *v
		}
		var ret QueryTimePeriod
		return ret
	}).(QueryTimePeriodOutput)
}

func (o QueryTimePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o QueryTimePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

type QueryTimePeriodResponse struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}





type QueryTimePeriodResponseInput interface {
	pulumi.Input

	ToQueryTimePeriodResponseOutput() QueryTimePeriodResponseOutput
	ToQueryTimePeriodResponseOutputWithContext(context.Context) QueryTimePeriodResponseOutput
}

type QueryTimePeriodResponseArgs struct {
	From pulumi.StringInput `pulumi:"from"`
	To   pulumi.StringInput `pulumi:"to"`
}

func (QueryTimePeriodResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryTimePeriodResponse)(nil)).Elem()
}

func (i QueryTimePeriodResponseArgs) ToQueryTimePeriodResponseOutput() QueryTimePeriodResponseOutput {
	return i.ToQueryTimePeriodResponseOutputWithContext(context.Background())
}

func (i QueryTimePeriodResponseArgs) ToQueryTimePeriodResponseOutputWithContext(ctx context.Context) QueryTimePeriodResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodResponseOutput)
}

func (i QueryTimePeriodResponseArgs) ToQueryTimePeriodResponsePtrOutput() QueryTimePeriodResponsePtrOutput {
	return i.ToQueryTimePeriodResponsePtrOutputWithContext(context.Background())
}

func (i QueryTimePeriodResponseArgs) ToQueryTimePeriodResponsePtrOutputWithContext(ctx context.Context) QueryTimePeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodResponseOutput).ToQueryTimePeriodResponsePtrOutputWithContext(ctx)
}









type QueryTimePeriodResponsePtrInput interface {
	pulumi.Input

	ToQueryTimePeriodResponsePtrOutput() QueryTimePeriodResponsePtrOutput
	ToQueryTimePeriodResponsePtrOutputWithContext(context.Context) QueryTimePeriodResponsePtrOutput
}

type queryTimePeriodResponsePtrType QueryTimePeriodResponseArgs

func QueryTimePeriodResponsePtr(v *QueryTimePeriodResponseArgs) QueryTimePeriodResponsePtrInput {
	return (*queryTimePeriodResponsePtrType)(v)
}

func (*queryTimePeriodResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryTimePeriodResponse)(nil)).Elem()
}

func (i *queryTimePeriodResponsePtrType) ToQueryTimePeriodResponsePtrOutput() QueryTimePeriodResponsePtrOutput {
	return i.ToQueryTimePeriodResponsePtrOutputWithContext(context.Background())
}

func (i *queryTimePeriodResponsePtrType) ToQueryTimePeriodResponsePtrOutputWithContext(ctx context.Context) QueryTimePeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueryTimePeriodResponsePtrOutput)
}

type QueryTimePeriodResponseOutput struct{ *pulumi.OutputState }

func (QueryTimePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryTimePeriodResponse)(nil)).Elem()
}

func (o QueryTimePeriodResponseOutput) ToQueryTimePeriodResponseOutput() QueryTimePeriodResponseOutput {
	return o
}

func (o QueryTimePeriodResponseOutput) ToQueryTimePeriodResponseOutputWithContext(ctx context.Context) QueryTimePeriodResponseOutput {
	return o
}

func (o QueryTimePeriodResponseOutput) ToQueryTimePeriodResponsePtrOutput() QueryTimePeriodResponsePtrOutput {
	return o.ToQueryTimePeriodResponsePtrOutputWithContext(context.Background())
}

func (o QueryTimePeriodResponseOutput) ToQueryTimePeriodResponsePtrOutputWithContext(ctx context.Context) QueryTimePeriodResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryTimePeriodResponse) *QueryTimePeriodResponse {
		return &v
	}).(QueryTimePeriodResponsePtrOutput)
}

func (o QueryTimePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v QueryTimePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o QueryTimePeriodResponseOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v QueryTimePeriodResponse) string { return v.To }).(pulumi.StringOutput)
}

type QueryTimePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (QueryTimePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryTimePeriodResponse)(nil)).Elem()
}

func (o QueryTimePeriodResponsePtrOutput) ToQueryTimePeriodResponsePtrOutput() QueryTimePeriodResponsePtrOutput {
	return o
}

func (o QueryTimePeriodResponsePtrOutput) ToQueryTimePeriodResponsePtrOutputWithContext(ctx context.Context) QueryTimePeriodResponsePtrOutput {
	return o
}

func (o QueryTimePeriodResponsePtrOutput) Elem() QueryTimePeriodResponseOutput {
	return o.ApplyT(func(v *QueryTimePeriodResponse) QueryTimePeriodResponse {
		if v != nil {
			return *v
		}
		var ret QueryTimePeriodResponse
		return ret
	}).(QueryTimePeriodResponseOutput)
}

func (o QueryTimePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o QueryTimePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueryTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

type ReportConfigAggregation struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}





type ReportConfigAggregationInput interface {
	pulumi.Input

	ToReportConfigAggregationOutput() ReportConfigAggregationOutput
	ToReportConfigAggregationOutputWithContext(context.Context) ReportConfigAggregationOutput
}

type ReportConfigAggregationArgs struct {
	Function pulumi.StringInput `pulumi:"function"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (ReportConfigAggregationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigAggregation)(nil)).Elem()
}

func (i ReportConfigAggregationArgs) ToReportConfigAggregationOutput() ReportConfigAggregationOutput {
	return i.ToReportConfigAggregationOutputWithContext(context.Background())
}

func (i ReportConfigAggregationArgs) ToReportConfigAggregationOutputWithContext(ctx context.Context) ReportConfigAggregationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigAggregationOutput)
}





type ReportConfigAggregationMapInput interface {
	pulumi.Input

	ToReportConfigAggregationMapOutput() ReportConfigAggregationMapOutput
	ToReportConfigAggregationMapOutputWithContext(context.Context) ReportConfigAggregationMapOutput
}

type ReportConfigAggregationMap map[string]ReportConfigAggregationInput

func (ReportConfigAggregationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportConfigAggregation)(nil)).Elem()
}

func (i ReportConfigAggregationMap) ToReportConfigAggregationMapOutput() ReportConfigAggregationMapOutput {
	return i.ToReportConfigAggregationMapOutputWithContext(context.Background())
}

func (i ReportConfigAggregationMap) ToReportConfigAggregationMapOutputWithContext(ctx context.Context) ReportConfigAggregationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigAggregationMapOutput)
}

type ReportConfigAggregationOutput struct{ *pulumi.OutputState }

func (ReportConfigAggregationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigAggregation)(nil)).Elem()
}

func (o ReportConfigAggregationOutput) ToReportConfigAggregationOutput() ReportConfigAggregationOutput {
	return o
}

func (o ReportConfigAggregationOutput) ToReportConfigAggregationOutputWithContext(ctx context.Context) ReportConfigAggregationOutput {
	return o
}

func (o ReportConfigAggregationOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigAggregation) string { return v.Function }).(pulumi.StringOutput)
}

func (o ReportConfigAggregationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigAggregation) string { return v.Name }).(pulumi.StringOutput)
}

type ReportConfigAggregationMapOutput struct{ *pulumi.OutputState }

func (ReportConfigAggregationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportConfigAggregation)(nil)).Elem()
}

func (o ReportConfigAggregationMapOutput) ToReportConfigAggregationMapOutput() ReportConfigAggregationMapOutput {
	return o
}

func (o ReportConfigAggregationMapOutput) ToReportConfigAggregationMapOutputWithContext(ctx context.Context) ReportConfigAggregationMapOutput {
	return o
}

func (o ReportConfigAggregationMapOutput) MapIndex(k pulumi.StringInput) ReportConfigAggregationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReportConfigAggregation {
		return vs[0].(map[string]ReportConfigAggregation)[vs[1].(string)]
	}).(ReportConfigAggregationOutput)
}

type ReportConfigAggregationResponse struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}





type ReportConfigAggregationResponseInput interface {
	pulumi.Input

	ToReportConfigAggregationResponseOutput() ReportConfigAggregationResponseOutput
	ToReportConfigAggregationResponseOutputWithContext(context.Context) ReportConfigAggregationResponseOutput
}

type ReportConfigAggregationResponseArgs struct {
	Function pulumi.StringInput `pulumi:"function"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (ReportConfigAggregationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigAggregationResponse)(nil)).Elem()
}

func (i ReportConfigAggregationResponseArgs) ToReportConfigAggregationResponseOutput() ReportConfigAggregationResponseOutput {
	return i.ToReportConfigAggregationResponseOutputWithContext(context.Background())
}

func (i ReportConfigAggregationResponseArgs) ToReportConfigAggregationResponseOutputWithContext(ctx context.Context) ReportConfigAggregationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigAggregationResponseOutput)
}





type ReportConfigAggregationResponseMapInput interface {
	pulumi.Input

	ToReportConfigAggregationResponseMapOutput() ReportConfigAggregationResponseMapOutput
	ToReportConfigAggregationResponseMapOutputWithContext(context.Context) ReportConfigAggregationResponseMapOutput
}

type ReportConfigAggregationResponseMap map[string]ReportConfigAggregationResponseInput

func (ReportConfigAggregationResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportConfigAggregationResponse)(nil)).Elem()
}

func (i ReportConfigAggregationResponseMap) ToReportConfigAggregationResponseMapOutput() ReportConfigAggregationResponseMapOutput {
	return i.ToReportConfigAggregationResponseMapOutputWithContext(context.Background())
}

func (i ReportConfigAggregationResponseMap) ToReportConfigAggregationResponseMapOutputWithContext(ctx context.Context) ReportConfigAggregationResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigAggregationResponseMapOutput)
}

type ReportConfigAggregationResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigAggregationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigAggregationResponse)(nil)).Elem()
}

func (o ReportConfigAggregationResponseOutput) ToReportConfigAggregationResponseOutput() ReportConfigAggregationResponseOutput {
	return o
}

func (o ReportConfigAggregationResponseOutput) ToReportConfigAggregationResponseOutputWithContext(ctx context.Context) ReportConfigAggregationResponseOutput {
	return o
}

func (o ReportConfigAggregationResponseOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigAggregationResponse) string { return v.Function }).(pulumi.StringOutput)
}

func (o ReportConfigAggregationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigAggregationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ReportConfigAggregationResponseMapOutput struct{ *pulumi.OutputState }

func (ReportConfigAggregationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportConfigAggregationResponse)(nil)).Elem()
}

func (o ReportConfigAggregationResponseMapOutput) ToReportConfigAggregationResponseMapOutput() ReportConfigAggregationResponseMapOutput {
	return o
}

func (o ReportConfigAggregationResponseMapOutput) ToReportConfigAggregationResponseMapOutputWithContext(ctx context.Context) ReportConfigAggregationResponseMapOutput {
	return o
}

func (o ReportConfigAggregationResponseMapOutput) MapIndex(k pulumi.StringInput) ReportConfigAggregationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReportConfigAggregationResponse {
		return vs[0].(map[string]ReportConfigAggregationResponse)[vs[1].(string)]
	}).(ReportConfigAggregationResponseOutput)
}

type ReportConfigComparisonExpression struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type ReportConfigComparisonExpressionInput interface {
	pulumi.Input

	ToReportConfigComparisonExpressionOutput() ReportConfigComparisonExpressionOutput
	ToReportConfigComparisonExpressionOutputWithContext(context.Context) ReportConfigComparisonExpressionOutput
}

type ReportConfigComparisonExpressionArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (ReportConfigComparisonExpressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigComparisonExpression)(nil)).Elem()
}

func (i ReportConfigComparisonExpressionArgs) ToReportConfigComparisonExpressionOutput() ReportConfigComparisonExpressionOutput {
	return i.ToReportConfigComparisonExpressionOutputWithContext(context.Background())
}

func (i ReportConfigComparisonExpressionArgs) ToReportConfigComparisonExpressionOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigComparisonExpressionOutput)
}

func (i ReportConfigComparisonExpressionArgs) ToReportConfigComparisonExpressionPtrOutput() ReportConfigComparisonExpressionPtrOutput {
	return i.ToReportConfigComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i ReportConfigComparisonExpressionArgs) ToReportConfigComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigComparisonExpressionOutput).ToReportConfigComparisonExpressionPtrOutputWithContext(ctx)
}









type ReportConfigComparisonExpressionPtrInput interface {
	pulumi.Input

	ToReportConfigComparisonExpressionPtrOutput() ReportConfigComparisonExpressionPtrOutput
	ToReportConfigComparisonExpressionPtrOutputWithContext(context.Context) ReportConfigComparisonExpressionPtrOutput
}

type reportConfigComparisonExpressionPtrType ReportConfigComparisonExpressionArgs

func ReportConfigComparisonExpressionPtr(v *ReportConfigComparisonExpressionArgs) ReportConfigComparisonExpressionPtrInput {
	return (*reportConfigComparisonExpressionPtrType)(v)
}

func (*reportConfigComparisonExpressionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigComparisonExpression)(nil)).Elem()
}

func (i *reportConfigComparisonExpressionPtrType) ToReportConfigComparisonExpressionPtrOutput() ReportConfigComparisonExpressionPtrOutput {
	return i.ToReportConfigComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i *reportConfigComparisonExpressionPtrType) ToReportConfigComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigComparisonExpressionPtrOutput)
}

type ReportConfigComparisonExpressionOutput struct{ *pulumi.OutputState }

func (ReportConfigComparisonExpressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigComparisonExpression)(nil)).Elem()
}

func (o ReportConfigComparisonExpressionOutput) ToReportConfigComparisonExpressionOutput() ReportConfigComparisonExpressionOutput {
	return o
}

func (o ReportConfigComparisonExpressionOutput) ToReportConfigComparisonExpressionOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionOutput {
	return o
}

func (o ReportConfigComparisonExpressionOutput) ToReportConfigComparisonExpressionPtrOutput() ReportConfigComparisonExpressionPtrOutput {
	return o.ToReportConfigComparisonExpressionPtrOutputWithContext(context.Background())
}

func (o ReportConfigComparisonExpressionOutput) ToReportConfigComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigComparisonExpression) *ReportConfigComparisonExpression {
		return &v
	}).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigComparisonExpressionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpression) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportConfigComparisonExpressionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpression) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ReportConfigComparisonExpressionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpression) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ReportConfigComparisonExpressionPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigComparisonExpressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigComparisonExpression)(nil)).Elem()
}

func (o ReportConfigComparisonExpressionPtrOutput) ToReportConfigComparisonExpressionPtrOutput() ReportConfigComparisonExpressionPtrOutput {
	return o
}

func (o ReportConfigComparisonExpressionPtrOutput) ToReportConfigComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionPtrOutput {
	return o
}

func (o ReportConfigComparisonExpressionPtrOutput) Elem() ReportConfigComparisonExpressionOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpression) ReportConfigComparisonExpression {
		if v != nil {
			return *v
		}
		var ret ReportConfigComparisonExpression
		return ret
	}).(ReportConfigComparisonExpressionOutput)
}

func (o ReportConfigComparisonExpressionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigComparisonExpressionPtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigComparisonExpressionPtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpression) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type ReportConfigComparisonExpressionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type ReportConfigComparisonExpressionResponseInput interface {
	pulumi.Input

	ToReportConfigComparisonExpressionResponseOutput() ReportConfigComparisonExpressionResponseOutput
	ToReportConfigComparisonExpressionResponseOutputWithContext(context.Context) ReportConfigComparisonExpressionResponseOutput
}

type ReportConfigComparisonExpressionResponseArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (ReportConfigComparisonExpressionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigComparisonExpressionResponse)(nil)).Elem()
}

func (i ReportConfigComparisonExpressionResponseArgs) ToReportConfigComparisonExpressionResponseOutput() ReportConfigComparisonExpressionResponseOutput {
	return i.ToReportConfigComparisonExpressionResponseOutputWithContext(context.Background())
}

func (i ReportConfigComparisonExpressionResponseArgs) ToReportConfigComparisonExpressionResponseOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigComparisonExpressionResponseOutput)
}

func (i ReportConfigComparisonExpressionResponseArgs) ToReportConfigComparisonExpressionResponsePtrOutput() ReportConfigComparisonExpressionResponsePtrOutput {
	return i.ToReportConfigComparisonExpressionResponsePtrOutputWithContext(context.Background())
}

func (i ReportConfigComparisonExpressionResponseArgs) ToReportConfigComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigComparisonExpressionResponseOutput).ToReportConfigComparisonExpressionResponsePtrOutputWithContext(ctx)
}









type ReportConfigComparisonExpressionResponsePtrInput interface {
	pulumi.Input

	ToReportConfigComparisonExpressionResponsePtrOutput() ReportConfigComparisonExpressionResponsePtrOutput
	ToReportConfigComparisonExpressionResponsePtrOutputWithContext(context.Context) ReportConfigComparisonExpressionResponsePtrOutput
}

type reportConfigComparisonExpressionResponsePtrType ReportConfigComparisonExpressionResponseArgs

func ReportConfigComparisonExpressionResponsePtr(v *ReportConfigComparisonExpressionResponseArgs) ReportConfigComparisonExpressionResponsePtrInput {
	return (*reportConfigComparisonExpressionResponsePtrType)(v)
}

func (*reportConfigComparisonExpressionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigComparisonExpressionResponse)(nil)).Elem()
}

func (i *reportConfigComparisonExpressionResponsePtrType) ToReportConfigComparisonExpressionResponsePtrOutput() ReportConfigComparisonExpressionResponsePtrOutput {
	return i.ToReportConfigComparisonExpressionResponsePtrOutputWithContext(context.Background())
}

func (i *reportConfigComparisonExpressionResponsePtrType) ToReportConfigComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigComparisonExpressionResponsePtrOutput)
}

type ReportConfigComparisonExpressionResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigComparisonExpressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigComparisonExpressionResponse)(nil)).Elem()
}

func (o ReportConfigComparisonExpressionResponseOutput) ToReportConfigComparisonExpressionResponseOutput() ReportConfigComparisonExpressionResponseOutput {
	return o
}

func (o ReportConfigComparisonExpressionResponseOutput) ToReportConfigComparisonExpressionResponseOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionResponseOutput {
	return o
}

func (o ReportConfigComparisonExpressionResponseOutput) ToReportConfigComparisonExpressionResponsePtrOutput() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ToReportConfigComparisonExpressionResponsePtrOutputWithContext(context.Background())
}

func (o ReportConfigComparisonExpressionResponseOutput) ToReportConfigComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigComparisonExpressionResponse) *ReportConfigComparisonExpressionResponse {
		return &v
	}).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigComparisonExpressionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpressionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportConfigComparisonExpressionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpressionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ReportConfigComparisonExpressionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportConfigComparisonExpressionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ReportConfigComparisonExpressionResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigComparisonExpressionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigComparisonExpressionResponse)(nil)).Elem()
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) ToReportConfigComparisonExpressionResponsePtrOutput() ReportConfigComparisonExpressionResponsePtrOutput {
	return o
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) ToReportConfigComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) ReportConfigComparisonExpressionResponsePtrOutput {
	return o
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) Elem() ReportConfigComparisonExpressionResponseOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpressionResponse) ReportConfigComparisonExpressionResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigComparisonExpressionResponse
		return ret
	}).(ReportConfigComparisonExpressionResponseOutput)
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigComparisonExpressionResponsePtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportConfigComparisonExpressionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type ReportConfigDataset struct {
	Aggregation   map[string]ReportConfigAggregation `pulumi:"aggregation"`
	Configuration *ReportConfigDatasetConfiguration  `pulumi:"configuration"`
	Filter        *ReportConfigFilter                `pulumi:"filter"`
	Granularity   *string                            `pulumi:"granularity"`
	Grouping      []ReportConfigGrouping             `pulumi:"grouping"`
	Sorting       []ReportConfigSorting              `pulumi:"sorting"`
}





type ReportConfigDatasetInput interface {
	pulumi.Input

	ToReportConfigDatasetOutput() ReportConfigDatasetOutput
	ToReportConfigDatasetOutputWithContext(context.Context) ReportConfigDatasetOutput
}

type ReportConfigDatasetArgs struct {
	Aggregation   ReportConfigAggregationMapInput          `pulumi:"aggregation"`
	Configuration ReportConfigDatasetConfigurationPtrInput `pulumi:"configuration"`
	Filter        ReportConfigFilterPtrInput               `pulumi:"filter"`
	Granularity   pulumi.StringPtrInput                    `pulumi:"granularity"`
	Grouping      ReportConfigGroupingArrayInput           `pulumi:"grouping"`
	Sorting       ReportConfigSortingArrayInput            `pulumi:"sorting"`
}

func (ReportConfigDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDataset)(nil)).Elem()
}

func (i ReportConfigDatasetArgs) ToReportConfigDatasetOutput() ReportConfigDatasetOutput {
	return i.ToReportConfigDatasetOutputWithContext(context.Background())
}

func (i ReportConfigDatasetArgs) ToReportConfigDatasetOutputWithContext(ctx context.Context) ReportConfigDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetOutput)
}

func (i ReportConfigDatasetArgs) ToReportConfigDatasetPtrOutput() ReportConfigDatasetPtrOutput {
	return i.ToReportConfigDatasetPtrOutputWithContext(context.Background())
}

func (i ReportConfigDatasetArgs) ToReportConfigDatasetPtrOutputWithContext(ctx context.Context) ReportConfigDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetOutput).ToReportConfigDatasetPtrOutputWithContext(ctx)
}









type ReportConfigDatasetPtrInput interface {
	pulumi.Input

	ToReportConfigDatasetPtrOutput() ReportConfigDatasetPtrOutput
	ToReportConfigDatasetPtrOutputWithContext(context.Context) ReportConfigDatasetPtrOutput
}

type reportConfigDatasetPtrType ReportConfigDatasetArgs

func ReportConfigDatasetPtr(v *ReportConfigDatasetArgs) ReportConfigDatasetPtrInput {
	return (*reportConfigDatasetPtrType)(v)
}

func (*reportConfigDatasetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDataset)(nil)).Elem()
}

func (i *reportConfigDatasetPtrType) ToReportConfigDatasetPtrOutput() ReportConfigDatasetPtrOutput {
	return i.ToReportConfigDatasetPtrOutputWithContext(context.Background())
}

func (i *reportConfigDatasetPtrType) ToReportConfigDatasetPtrOutputWithContext(ctx context.Context) ReportConfigDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetPtrOutput)
}

type ReportConfigDatasetOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDataset)(nil)).Elem()
}

func (o ReportConfigDatasetOutput) ToReportConfigDatasetOutput() ReportConfigDatasetOutput {
	return o
}

func (o ReportConfigDatasetOutput) ToReportConfigDatasetOutputWithContext(ctx context.Context) ReportConfigDatasetOutput {
	return o
}

func (o ReportConfigDatasetOutput) ToReportConfigDatasetPtrOutput() ReportConfigDatasetPtrOutput {
	return o.ToReportConfigDatasetPtrOutputWithContext(context.Background())
}

func (o ReportConfigDatasetOutput) ToReportConfigDatasetPtrOutputWithContext(ctx context.Context) ReportConfigDatasetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigDataset) *ReportConfigDataset {
		return &v
	}).(ReportConfigDatasetPtrOutput)
}

func (o ReportConfigDatasetOutput) Aggregation() ReportConfigAggregationMapOutput {
	return o.ApplyT(func(v ReportConfigDataset) map[string]ReportConfigAggregation { return v.Aggregation }).(ReportConfigAggregationMapOutput)
}

func (o ReportConfigDatasetOutput) Configuration() ReportConfigDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v ReportConfigDataset) *ReportConfigDatasetConfiguration { return v.Configuration }).(ReportConfigDatasetConfigurationPtrOutput)
}

func (o ReportConfigDatasetOutput) Filter() ReportConfigFilterPtrOutput {
	return o.ApplyT(func(v ReportConfigDataset) *ReportConfigFilter { return v.Filter }).(ReportConfigFilterPtrOutput)
}

func (o ReportConfigDatasetOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigDataset) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o ReportConfigDatasetOutput) Grouping() ReportConfigGroupingArrayOutput {
	return o.ApplyT(func(v ReportConfigDataset) []ReportConfigGrouping { return v.Grouping }).(ReportConfigGroupingArrayOutput)
}

func (o ReportConfigDatasetOutput) Sorting() ReportConfigSortingArrayOutput {
	return o.ApplyT(func(v ReportConfigDataset) []ReportConfigSorting { return v.Sorting }).(ReportConfigSortingArrayOutput)
}

type ReportConfigDatasetPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDataset)(nil)).Elem()
}

func (o ReportConfigDatasetPtrOutput) ToReportConfigDatasetPtrOutput() ReportConfigDatasetPtrOutput {
	return o
}

func (o ReportConfigDatasetPtrOutput) ToReportConfigDatasetPtrOutputWithContext(ctx context.Context) ReportConfigDatasetPtrOutput {
	return o
}

func (o ReportConfigDatasetPtrOutput) Elem() ReportConfigDatasetOutput {
	return o.ApplyT(func(v *ReportConfigDataset) ReportConfigDataset {
		if v != nil {
			return *v
		}
		var ret ReportConfigDataset
		return ret
	}).(ReportConfigDatasetOutput)
}

func (o ReportConfigDatasetPtrOutput) Aggregation() ReportConfigAggregationMapOutput {
	return o.ApplyT(func(v *ReportConfigDataset) map[string]ReportConfigAggregation {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(ReportConfigAggregationMapOutput)
}

func (o ReportConfigDatasetPtrOutput) Configuration() ReportConfigDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v *ReportConfigDataset) *ReportConfigDatasetConfiguration {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ReportConfigDatasetConfigurationPtrOutput)
}

func (o ReportConfigDatasetPtrOutput) Filter() ReportConfigFilterPtrOutput {
	return o.ApplyT(func(v *ReportConfigDataset) *ReportConfigFilter {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(ReportConfigFilterPtrOutput)
}

func (o ReportConfigDatasetPtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigDataset) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigDatasetPtrOutput) Grouping() ReportConfigGroupingArrayOutput {
	return o.ApplyT(func(v *ReportConfigDataset) []ReportConfigGrouping {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(ReportConfigGroupingArrayOutput)
}

func (o ReportConfigDatasetPtrOutput) Sorting() ReportConfigSortingArrayOutput {
	return o.ApplyT(func(v *ReportConfigDataset) []ReportConfigSorting {
		if v == nil {
			return nil
		}
		return v.Sorting
	}).(ReportConfigSortingArrayOutput)
}

type ReportConfigDatasetConfiguration struct {
	Columns []string `pulumi:"columns"`
}





type ReportConfigDatasetConfigurationInput interface {
	pulumi.Input

	ToReportConfigDatasetConfigurationOutput() ReportConfigDatasetConfigurationOutput
	ToReportConfigDatasetConfigurationOutputWithContext(context.Context) ReportConfigDatasetConfigurationOutput
}

type ReportConfigDatasetConfigurationArgs struct {
	Columns pulumi.StringArrayInput `pulumi:"columns"`
}

func (ReportConfigDatasetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDatasetConfiguration)(nil)).Elem()
}

func (i ReportConfigDatasetConfigurationArgs) ToReportConfigDatasetConfigurationOutput() ReportConfigDatasetConfigurationOutput {
	return i.ToReportConfigDatasetConfigurationOutputWithContext(context.Background())
}

func (i ReportConfigDatasetConfigurationArgs) ToReportConfigDatasetConfigurationOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetConfigurationOutput)
}

func (i ReportConfigDatasetConfigurationArgs) ToReportConfigDatasetConfigurationPtrOutput() ReportConfigDatasetConfigurationPtrOutput {
	return i.ToReportConfigDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i ReportConfigDatasetConfigurationArgs) ToReportConfigDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetConfigurationOutput).ToReportConfigDatasetConfigurationPtrOutputWithContext(ctx)
}









type ReportConfigDatasetConfigurationPtrInput interface {
	pulumi.Input

	ToReportConfigDatasetConfigurationPtrOutput() ReportConfigDatasetConfigurationPtrOutput
	ToReportConfigDatasetConfigurationPtrOutputWithContext(context.Context) ReportConfigDatasetConfigurationPtrOutput
}

type reportConfigDatasetConfigurationPtrType ReportConfigDatasetConfigurationArgs

func ReportConfigDatasetConfigurationPtr(v *ReportConfigDatasetConfigurationArgs) ReportConfigDatasetConfigurationPtrInput {
	return (*reportConfigDatasetConfigurationPtrType)(v)
}

func (*reportConfigDatasetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDatasetConfiguration)(nil)).Elem()
}

func (i *reportConfigDatasetConfigurationPtrType) ToReportConfigDatasetConfigurationPtrOutput() ReportConfigDatasetConfigurationPtrOutput {
	return i.ToReportConfigDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i *reportConfigDatasetConfigurationPtrType) ToReportConfigDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetConfigurationPtrOutput)
}

type ReportConfigDatasetConfigurationOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDatasetConfiguration)(nil)).Elem()
}

func (o ReportConfigDatasetConfigurationOutput) ToReportConfigDatasetConfigurationOutput() ReportConfigDatasetConfigurationOutput {
	return o
}

func (o ReportConfigDatasetConfigurationOutput) ToReportConfigDatasetConfigurationOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationOutput {
	return o
}

func (o ReportConfigDatasetConfigurationOutput) ToReportConfigDatasetConfigurationPtrOutput() ReportConfigDatasetConfigurationPtrOutput {
	return o.ToReportConfigDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (o ReportConfigDatasetConfigurationOutput) ToReportConfigDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigDatasetConfiguration) *ReportConfigDatasetConfiguration {
		return &v
	}).(ReportConfigDatasetConfigurationPtrOutput)
}

func (o ReportConfigDatasetConfigurationOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportConfigDatasetConfiguration) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type ReportConfigDatasetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDatasetConfiguration)(nil)).Elem()
}

func (o ReportConfigDatasetConfigurationPtrOutput) ToReportConfigDatasetConfigurationPtrOutput() ReportConfigDatasetConfigurationPtrOutput {
	return o
}

func (o ReportConfigDatasetConfigurationPtrOutput) ToReportConfigDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationPtrOutput {
	return o
}

func (o ReportConfigDatasetConfigurationPtrOutput) Elem() ReportConfigDatasetConfigurationOutput {
	return o.ApplyT(func(v *ReportConfigDatasetConfiguration) ReportConfigDatasetConfiguration {
		if v != nil {
			return *v
		}
		var ret ReportConfigDatasetConfiguration
		return ret
	}).(ReportConfigDatasetConfigurationOutput)
}

func (o ReportConfigDatasetConfigurationPtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportConfigDatasetConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type ReportConfigDatasetConfigurationResponse struct {
	Columns []string `pulumi:"columns"`
}





type ReportConfigDatasetConfigurationResponseInput interface {
	pulumi.Input

	ToReportConfigDatasetConfigurationResponseOutput() ReportConfigDatasetConfigurationResponseOutput
	ToReportConfigDatasetConfigurationResponseOutputWithContext(context.Context) ReportConfigDatasetConfigurationResponseOutput
}

type ReportConfigDatasetConfigurationResponseArgs struct {
	Columns pulumi.StringArrayInput `pulumi:"columns"`
}

func (ReportConfigDatasetConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDatasetConfigurationResponse)(nil)).Elem()
}

func (i ReportConfigDatasetConfigurationResponseArgs) ToReportConfigDatasetConfigurationResponseOutput() ReportConfigDatasetConfigurationResponseOutput {
	return i.ToReportConfigDatasetConfigurationResponseOutputWithContext(context.Background())
}

func (i ReportConfigDatasetConfigurationResponseArgs) ToReportConfigDatasetConfigurationResponseOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetConfigurationResponseOutput)
}

func (i ReportConfigDatasetConfigurationResponseArgs) ToReportConfigDatasetConfigurationResponsePtrOutput() ReportConfigDatasetConfigurationResponsePtrOutput {
	return i.ToReportConfigDatasetConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i ReportConfigDatasetConfigurationResponseArgs) ToReportConfigDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetConfigurationResponseOutput).ToReportConfigDatasetConfigurationResponsePtrOutputWithContext(ctx)
}









type ReportConfigDatasetConfigurationResponsePtrInput interface {
	pulumi.Input

	ToReportConfigDatasetConfigurationResponsePtrOutput() ReportConfigDatasetConfigurationResponsePtrOutput
	ToReportConfigDatasetConfigurationResponsePtrOutputWithContext(context.Context) ReportConfigDatasetConfigurationResponsePtrOutput
}

type reportConfigDatasetConfigurationResponsePtrType ReportConfigDatasetConfigurationResponseArgs

func ReportConfigDatasetConfigurationResponsePtr(v *ReportConfigDatasetConfigurationResponseArgs) ReportConfigDatasetConfigurationResponsePtrInput {
	return (*reportConfigDatasetConfigurationResponsePtrType)(v)
}

func (*reportConfigDatasetConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDatasetConfigurationResponse)(nil)).Elem()
}

func (i *reportConfigDatasetConfigurationResponsePtrType) ToReportConfigDatasetConfigurationResponsePtrOutput() ReportConfigDatasetConfigurationResponsePtrOutput {
	return i.ToReportConfigDatasetConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *reportConfigDatasetConfigurationResponsePtrType) ToReportConfigDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetConfigurationResponsePtrOutput)
}

type ReportConfigDatasetConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDatasetConfigurationResponse)(nil)).Elem()
}

func (o ReportConfigDatasetConfigurationResponseOutput) ToReportConfigDatasetConfigurationResponseOutput() ReportConfigDatasetConfigurationResponseOutput {
	return o
}

func (o ReportConfigDatasetConfigurationResponseOutput) ToReportConfigDatasetConfigurationResponseOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationResponseOutput {
	return o
}

func (o ReportConfigDatasetConfigurationResponseOutput) ToReportConfigDatasetConfigurationResponsePtrOutput() ReportConfigDatasetConfigurationResponsePtrOutput {
	return o.ToReportConfigDatasetConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o ReportConfigDatasetConfigurationResponseOutput) ToReportConfigDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigDatasetConfigurationResponse) *ReportConfigDatasetConfigurationResponse {
		return &v
	}).(ReportConfigDatasetConfigurationResponsePtrOutput)
}

func (o ReportConfigDatasetConfigurationResponseOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportConfigDatasetConfigurationResponse) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type ReportConfigDatasetConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDatasetConfigurationResponse)(nil)).Elem()
}

func (o ReportConfigDatasetConfigurationResponsePtrOutput) ToReportConfigDatasetConfigurationResponsePtrOutput() ReportConfigDatasetConfigurationResponsePtrOutput {
	return o
}

func (o ReportConfigDatasetConfigurationResponsePtrOutput) ToReportConfigDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) ReportConfigDatasetConfigurationResponsePtrOutput {
	return o
}

func (o ReportConfigDatasetConfigurationResponsePtrOutput) Elem() ReportConfigDatasetConfigurationResponseOutput {
	return o.ApplyT(func(v *ReportConfigDatasetConfigurationResponse) ReportConfigDatasetConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigDatasetConfigurationResponse
		return ret
	}).(ReportConfigDatasetConfigurationResponseOutput)
}

func (o ReportConfigDatasetConfigurationResponsePtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportConfigDatasetConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type ReportConfigDatasetResponse struct {
	Aggregation   map[string]ReportConfigAggregationResponse `pulumi:"aggregation"`
	Configuration *ReportConfigDatasetConfigurationResponse  `pulumi:"configuration"`
	Filter        *ReportConfigFilterResponse                `pulumi:"filter"`
	Granularity   *string                                    `pulumi:"granularity"`
	Grouping      []ReportConfigGroupingResponse             `pulumi:"grouping"`
	Sorting       []ReportConfigSortingResponse              `pulumi:"sorting"`
}





type ReportConfigDatasetResponseInput interface {
	pulumi.Input

	ToReportConfigDatasetResponseOutput() ReportConfigDatasetResponseOutput
	ToReportConfigDatasetResponseOutputWithContext(context.Context) ReportConfigDatasetResponseOutput
}

type ReportConfigDatasetResponseArgs struct {
	Aggregation   ReportConfigAggregationResponseMapInput          `pulumi:"aggregation"`
	Configuration ReportConfigDatasetConfigurationResponsePtrInput `pulumi:"configuration"`
	Filter        ReportConfigFilterResponsePtrInput               `pulumi:"filter"`
	Granularity   pulumi.StringPtrInput                            `pulumi:"granularity"`
	Grouping      ReportConfigGroupingResponseArrayInput           `pulumi:"grouping"`
	Sorting       ReportConfigSortingResponseArrayInput            `pulumi:"sorting"`
}

func (ReportConfigDatasetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDatasetResponse)(nil)).Elem()
}

func (i ReportConfigDatasetResponseArgs) ToReportConfigDatasetResponseOutput() ReportConfigDatasetResponseOutput {
	return i.ToReportConfigDatasetResponseOutputWithContext(context.Background())
}

func (i ReportConfigDatasetResponseArgs) ToReportConfigDatasetResponseOutputWithContext(ctx context.Context) ReportConfigDatasetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetResponseOutput)
}

func (i ReportConfigDatasetResponseArgs) ToReportConfigDatasetResponsePtrOutput() ReportConfigDatasetResponsePtrOutput {
	return i.ToReportConfigDatasetResponsePtrOutputWithContext(context.Background())
}

func (i ReportConfigDatasetResponseArgs) ToReportConfigDatasetResponsePtrOutputWithContext(ctx context.Context) ReportConfigDatasetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetResponseOutput).ToReportConfigDatasetResponsePtrOutputWithContext(ctx)
}









type ReportConfigDatasetResponsePtrInput interface {
	pulumi.Input

	ToReportConfigDatasetResponsePtrOutput() ReportConfigDatasetResponsePtrOutput
	ToReportConfigDatasetResponsePtrOutputWithContext(context.Context) ReportConfigDatasetResponsePtrOutput
}

type reportConfigDatasetResponsePtrType ReportConfigDatasetResponseArgs

func ReportConfigDatasetResponsePtr(v *ReportConfigDatasetResponseArgs) ReportConfigDatasetResponsePtrInput {
	return (*reportConfigDatasetResponsePtrType)(v)
}

func (*reportConfigDatasetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDatasetResponse)(nil)).Elem()
}

func (i *reportConfigDatasetResponsePtrType) ToReportConfigDatasetResponsePtrOutput() ReportConfigDatasetResponsePtrOutput {
	return i.ToReportConfigDatasetResponsePtrOutputWithContext(context.Background())
}

func (i *reportConfigDatasetResponsePtrType) ToReportConfigDatasetResponsePtrOutputWithContext(ctx context.Context) ReportConfigDatasetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigDatasetResponsePtrOutput)
}

type ReportConfigDatasetResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigDatasetResponse)(nil)).Elem()
}

func (o ReportConfigDatasetResponseOutput) ToReportConfigDatasetResponseOutput() ReportConfigDatasetResponseOutput {
	return o
}

func (o ReportConfigDatasetResponseOutput) ToReportConfigDatasetResponseOutputWithContext(ctx context.Context) ReportConfigDatasetResponseOutput {
	return o
}

func (o ReportConfigDatasetResponseOutput) ToReportConfigDatasetResponsePtrOutput() ReportConfigDatasetResponsePtrOutput {
	return o.ToReportConfigDatasetResponsePtrOutputWithContext(context.Background())
}

func (o ReportConfigDatasetResponseOutput) ToReportConfigDatasetResponsePtrOutputWithContext(ctx context.Context) ReportConfigDatasetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigDatasetResponse) *ReportConfigDatasetResponse {
		return &v
	}).(ReportConfigDatasetResponsePtrOutput)
}

func (o ReportConfigDatasetResponseOutput) Aggregation() ReportConfigAggregationResponseMapOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) map[string]ReportConfigAggregationResponse { return v.Aggregation }).(ReportConfigAggregationResponseMapOutput)
}

func (o ReportConfigDatasetResponseOutput) Configuration() ReportConfigDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) *ReportConfigDatasetConfigurationResponse { return v.Configuration }).(ReportConfigDatasetConfigurationResponsePtrOutput)
}

func (o ReportConfigDatasetResponseOutput) Filter() ReportConfigFilterResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) *ReportConfigFilterResponse { return v.Filter }).(ReportConfigFilterResponsePtrOutput)
}

func (o ReportConfigDatasetResponseOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o ReportConfigDatasetResponseOutput) Grouping() ReportConfigGroupingResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) []ReportConfigGroupingResponse { return v.Grouping }).(ReportConfigGroupingResponseArrayOutput)
}

func (o ReportConfigDatasetResponseOutput) Sorting() ReportConfigSortingResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigDatasetResponse) []ReportConfigSortingResponse { return v.Sorting }).(ReportConfigSortingResponseArrayOutput)
}

type ReportConfigDatasetResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigDatasetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigDatasetResponse)(nil)).Elem()
}

func (o ReportConfigDatasetResponsePtrOutput) ToReportConfigDatasetResponsePtrOutput() ReportConfigDatasetResponsePtrOutput {
	return o
}

func (o ReportConfigDatasetResponsePtrOutput) ToReportConfigDatasetResponsePtrOutputWithContext(ctx context.Context) ReportConfigDatasetResponsePtrOutput {
	return o
}

func (o ReportConfigDatasetResponsePtrOutput) Elem() ReportConfigDatasetResponseOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) ReportConfigDatasetResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigDatasetResponse
		return ret
	}).(ReportConfigDatasetResponseOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Aggregation() ReportConfigAggregationResponseMapOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) map[string]ReportConfigAggregationResponse {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(ReportConfigAggregationResponseMapOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Configuration() ReportConfigDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) *ReportConfigDatasetConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ReportConfigDatasetConfigurationResponsePtrOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Filter() ReportConfigFilterResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) *ReportConfigFilterResponse {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(ReportConfigFilterResponsePtrOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Grouping() ReportConfigGroupingResponseArrayOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) []ReportConfigGroupingResponse {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(ReportConfigGroupingResponseArrayOutput)
}

func (o ReportConfigDatasetResponsePtrOutput) Sorting() ReportConfigSortingResponseArrayOutput {
	return o.ApplyT(func(v *ReportConfigDatasetResponse) []ReportConfigSortingResponse {
		if v == nil {
			return nil
		}
		return v.Sorting
	}).(ReportConfigSortingResponseArrayOutput)
}

type ReportConfigFilter struct {
	And        []ReportConfigFilter              `pulumi:"and"`
	Dimensions *ReportConfigComparisonExpression `pulumi:"dimensions"`
	Or         []ReportConfigFilter              `pulumi:"or"`
	TagKey     *ReportConfigComparisonExpression `pulumi:"tagKey"`
	TagValue   *ReportConfigComparisonExpression `pulumi:"tagValue"`
	Tags       *ReportConfigComparisonExpression `pulumi:"tags"`
}





type ReportConfigFilterInput interface {
	pulumi.Input

	ToReportConfigFilterOutput() ReportConfigFilterOutput
	ToReportConfigFilterOutputWithContext(context.Context) ReportConfigFilterOutput
}

type ReportConfigFilterArgs struct {
	And        ReportConfigFilterArrayInput             `pulumi:"and"`
	Dimensions ReportConfigComparisonExpressionPtrInput `pulumi:"dimensions"`
	Or         ReportConfigFilterArrayInput             `pulumi:"or"`
	TagKey     ReportConfigComparisonExpressionPtrInput `pulumi:"tagKey"`
	TagValue   ReportConfigComparisonExpressionPtrInput `pulumi:"tagValue"`
	Tags       ReportConfigComparisonExpressionPtrInput `pulumi:"tags"`
}

func (ReportConfigFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigFilter)(nil)).Elem()
}

func (i ReportConfigFilterArgs) ToReportConfigFilterOutput() ReportConfigFilterOutput {
	return i.ToReportConfigFilterOutputWithContext(context.Background())
}

func (i ReportConfigFilterArgs) ToReportConfigFilterOutputWithContext(ctx context.Context) ReportConfigFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterOutput)
}

func (i ReportConfigFilterArgs) ToReportConfigFilterPtrOutput() ReportConfigFilterPtrOutput {
	return i.ToReportConfigFilterPtrOutputWithContext(context.Background())
}

func (i ReportConfigFilterArgs) ToReportConfigFilterPtrOutputWithContext(ctx context.Context) ReportConfigFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterOutput).ToReportConfigFilterPtrOutputWithContext(ctx)
}









type ReportConfigFilterPtrInput interface {
	pulumi.Input

	ToReportConfigFilterPtrOutput() ReportConfigFilterPtrOutput
	ToReportConfigFilterPtrOutputWithContext(context.Context) ReportConfigFilterPtrOutput
}

type reportConfigFilterPtrType ReportConfigFilterArgs

func ReportConfigFilterPtr(v *ReportConfigFilterArgs) ReportConfigFilterPtrInput {
	return (*reportConfigFilterPtrType)(v)
}

func (*reportConfigFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigFilter)(nil)).Elem()
}

func (i *reportConfigFilterPtrType) ToReportConfigFilterPtrOutput() ReportConfigFilterPtrOutput {
	return i.ToReportConfigFilterPtrOutputWithContext(context.Background())
}

func (i *reportConfigFilterPtrType) ToReportConfigFilterPtrOutputWithContext(ctx context.Context) ReportConfigFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterPtrOutput)
}





type ReportConfigFilterArrayInput interface {
	pulumi.Input

	ToReportConfigFilterArrayOutput() ReportConfigFilterArrayOutput
	ToReportConfigFilterArrayOutputWithContext(context.Context) ReportConfigFilterArrayOutput
}

type ReportConfigFilterArray []ReportConfigFilterInput

func (ReportConfigFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigFilter)(nil)).Elem()
}

func (i ReportConfigFilterArray) ToReportConfigFilterArrayOutput() ReportConfigFilterArrayOutput {
	return i.ToReportConfigFilterArrayOutputWithContext(context.Background())
}

func (i ReportConfigFilterArray) ToReportConfigFilterArrayOutputWithContext(ctx context.Context) ReportConfigFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterArrayOutput)
}

type ReportConfigFilterOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigFilter)(nil)).Elem()
}

func (o ReportConfigFilterOutput) ToReportConfigFilterOutput() ReportConfigFilterOutput {
	return o
}

func (o ReportConfigFilterOutput) ToReportConfigFilterOutputWithContext(ctx context.Context) ReportConfigFilterOutput {
	return o
}

func (o ReportConfigFilterOutput) ToReportConfigFilterPtrOutput() ReportConfigFilterPtrOutput {
	return o.ToReportConfigFilterPtrOutputWithContext(context.Background())
}

func (o ReportConfigFilterOutput) ToReportConfigFilterPtrOutputWithContext(ctx context.Context) ReportConfigFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigFilter) *ReportConfigFilter {
		return &v
	}).(ReportConfigFilterPtrOutput)
}

func (o ReportConfigFilterOutput) And() ReportConfigFilterArrayOutput {
	return o.ApplyT(func(v ReportConfigFilter) []ReportConfigFilter { return v.And }).(ReportConfigFilterArrayOutput)
}

func (o ReportConfigFilterOutput) Dimensions() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportConfigFilter) *ReportConfigComparisonExpression { return v.Dimensions }).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigFilterOutput) Or() ReportConfigFilterArrayOutput {
	return o.ApplyT(func(v ReportConfigFilter) []ReportConfigFilter { return v.Or }).(ReportConfigFilterArrayOutput)
}

func (o ReportConfigFilterOutput) TagKey() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportConfigFilter) *ReportConfigComparisonExpression { return v.TagKey }).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigFilterOutput) TagValue() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportConfigFilter) *ReportConfigComparisonExpression { return v.TagValue }).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigFilterOutput) Tags() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportConfigFilter) *ReportConfigComparisonExpression { return v.Tags }).(ReportConfigComparisonExpressionPtrOutput)
}

type ReportConfigFilterPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigFilter)(nil)).Elem()
}

func (o ReportConfigFilterPtrOutput) ToReportConfigFilterPtrOutput() ReportConfigFilterPtrOutput {
	return o
}

func (o ReportConfigFilterPtrOutput) ToReportConfigFilterPtrOutputWithContext(ctx context.Context) ReportConfigFilterPtrOutput {
	return o
}

func (o ReportConfigFilterPtrOutput) Elem() ReportConfigFilterOutput {
	return o.ApplyT(func(v *ReportConfigFilter) ReportConfigFilter {
		if v != nil {
			return *v
		}
		var ret ReportConfigFilter
		return ret
	}).(ReportConfigFilterOutput)
}

func (o ReportConfigFilterPtrOutput) And() ReportConfigFilterArrayOutput {
	return o.ApplyT(func(v *ReportConfigFilter) []ReportConfigFilter {
		if v == nil {
			return nil
		}
		return v.And
	}).(ReportConfigFilterArrayOutput)
}

func (o ReportConfigFilterPtrOutput) Dimensions() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportConfigFilter) *ReportConfigComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Dimensions
	}).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigFilterPtrOutput) Or() ReportConfigFilterArrayOutput {
	return o.ApplyT(func(v *ReportConfigFilter) []ReportConfigFilter {
		if v == nil {
			return nil
		}
		return v.Or
	}).(ReportConfigFilterArrayOutput)
}

func (o ReportConfigFilterPtrOutput) TagKey() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportConfigFilter) *ReportConfigComparisonExpression {
		if v == nil {
			return nil
		}
		return v.TagKey
	}).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigFilterPtrOutput) TagValue() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportConfigFilter) *ReportConfigComparisonExpression {
		if v == nil {
			return nil
		}
		return v.TagValue
	}).(ReportConfigComparisonExpressionPtrOutput)
}

func (o ReportConfigFilterPtrOutput) Tags() ReportConfigComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportConfigFilter) *ReportConfigComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(ReportConfigComparisonExpressionPtrOutput)
}

type ReportConfigFilterArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigFilter)(nil)).Elem()
}

func (o ReportConfigFilterArrayOutput) ToReportConfigFilterArrayOutput() ReportConfigFilterArrayOutput {
	return o
}

func (o ReportConfigFilterArrayOutput) ToReportConfigFilterArrayOutputWithContext(ctx context.Context) ReportConfigFilterArrayOutput {
	return o
}

func (o ReportConfigFilterArrayOutput) Index(i pulumi.IntInput) ReportConfigFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigFilter {
		return vs[0].([]ReportConfigFilter)[vs[1].(int)]
	}).(ReportConfigFilterOutput)
}

type ReportConfigFilterResponse struct {
	And        []ReportConfigFilterResponse              `pulumi:"and"`
	Dimensions *ReportConfigComparisonExpressionResponse `pulumi:"dimensions"`
	Or         []ReportConfigFilterResponse              `pulumi:"or"`
	TagKey     *ReportConfigComparisonExpressionResponse `pulumi:"tagKey"`
	TagValue   *ReportConfigComparisonExpressionResponse `pulumi:"tagValue"`
	Tags       *ReportConfigComparisonExpressionResponse `pulumi:"tags"`
}





type ReportConfigFilterResponseInput interface {
	pulumi.Input

	ToReportConfigFilterResponseOutput() ReportConfigFilterResponseOutput
	ToReportConfigFilterResponseOutputWithContext(context.Context) ReportConfigFilterResponseOutput
}

type ReportConfigFilterResponseArgs struct {
	And        ReportConfigFilterResponseArrayInput             `pulumi:"and"`
	Dimensions ReportConfigComparisonExpressionResponsePtrInput `pulumi:"dimensions"`
	Or         ReportConfigFilterResponseArrayInput             `pulumi:"or"`
	TagKey     ReportConfigComparisonExpressionResponsePtrInput `pulumi:"tagKey"`
	TagValue   ReportConfigComparisonExpressionResponsePtrInput `pulumi:"tagValue"`
	Tags       ReportConfigComparisonExpressionResponsePtrInput `pulumi:"tags"`
}

func (ReportConfigFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigFilterResponse)(nil)).Elem()
}

func (i ReportConfigFilterResponseArgs) ToReportConfigFilterResponseOutput() ReportConfigFilterResponseOutput {
	return i.ToReportConfigFilterResponseOutputWithContext(context.Background())
}

func (i ReportConfigFilterResponseArgs) ToReportConfigFilterResponseOutputWithContext(ctx context.Context) ReportConfigFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterResponseOutput)
}

func (i ReportConfigFilterResponseArgs) ToReportConfigFilterResponsePtrOutput() ReportConfigFilterResponsePtrOutput {
	return i.ToReportConfigFilterResponsePtrOutputWithContext(context.Background())
}

func (i ReportConfigFilterResponseArgs) ToReportConfigFilterResponsePtrOutputWithContext(ctx context.Context) ReportConfigFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterResponseOutput).ToReportConfigFilterResponsePtrOutputWithContext(ctx)
}









type ReportConfigFilterResponsePtrInput interface {
	pulumi.Input

	ToReportConfigFilterResponsePtrOutput() ReportConfigFilterResponsePtrOutput
	ToReportConfigFilterResponsePtrOutputWithContext(context.Context) ReportConfigFilterResponsePtrOutput
}

type reportConfigFilterResponsePtrType ReportConfigFilterResponseArgs

func ReportConfigFilterResponsePtr(v *ReportConfigFilterResponseArgs) ReportConfigFilterResponsePtrInput {
	return (*reportConfigFilterResponsePtrType)(v)
}

func (*reportConfigFilterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigFilterResponse)(nil)).Elem()
}

func (i *reportConfigFilterResponsePtrType) ToReportConfigFilterResponsePtrOutput() ReportConfigFilterResponsePtrOutput {
	return i.ToReportConfigFilterResponsePtrOutputWithContext(context.Background())
}

func (i *reportConfigFilterResponsePtrType) ToReportConfigFilterResponsePtrOutputWithContext(ctx context.Context) ReportConfigFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterResponsePtrOutput)
}





type ReportConfigFilterResponseArrayInput interface {
	pulumi.Input

	ToReportConfigFilterResponseArrayOutput() ReportConfigFilterResponseArrayOutput
	ToReportConfigFilterResponseArrayOutputWithContext(context.Context) ReportConfigFilterResponseArrayOutput
}

type ReportConfigFilterResponseArray []ReportConfigFilterResponseInput

func (ReportConfigFilterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigFilterResponse)(nil)).Elem()
}

func (i ReportConfigFilterResponseArray) ToReportConfigFilterResponseArrayOutput() ReportConfigFilterResponseArrayOutput {
	return i.ToReportConfigFilterResponseArrayOutputWithContext(context.Background())
}

func (i ReportConfigFilterResponseArray) ToReportConfigFilterResponseArrayOutputWithContext(ctx context.Context) ReportConfigFilterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigFilterResponseArrayOutput)
}

type ReportConfigFilterResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigFilterResponse)(nil)).Elem()
}

func (o ReportConfigFilterResponseOutput) ToReportConfigFilterResponseOutput() ReportConfigFilterResponseOutput {
	return o
}

func (o ReportConfigFilterResponseOutput) ToReportConfigFilterResponseOutputWithContext(ctx context.Context) ReportConfigFilterResponseOutput {
	return o
}

func (o ReportConfigFilterResponseOutput) ToReportConfigFilterResponsePtrOutput() ReportConfigFilterResponsePtrOutput {
	return o.ToReportConfigFilterResponsePtrOutputWithContext(context.Background())
}

func (o ReportConfigFilterResponseOutput) ToReportConfigFilterResponsePtrOutputWithContext(ctx context.Context) ReportConfigFilterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigFilterResponse) *ReportConfigFilterResponse {
		return &v
	}).(ReportConfigFilterResponsePtrOutput)
}

func (o ReportConfigFilterResponseOutput) And() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) []ReportConfigFilterResponse { return v.And }).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponseOutput) Dimensions() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse { return v.Dimensions }).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponseOutput) Or() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) []ReportConfigFilterResponse { return v.Or }).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponseOutput) TagKey() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse { return v.TagKey }).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponseOutput) TagValue() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse { return v.TagValue }).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponseOutput) Tags() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse { return v.Tags }).(ReportConfigComparisonExpressionResponsePtrOutput)
}

type ReportConfigFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigFilterResponse)(nil)).Elem()
}

func (o ReportConfigFilterResponsePtrOutput) ToReportConfigFilterResponsePtrOutput() ReportConfigFilterResponsePtrOutput {
	return o
}

func (o ReportConfigFilterResponsePtrOutput) ToReportConfigFilterResponsePtrOutputWithContext(ctx context.Context) ReportConfigFilterResponsePtrOutput {
	return o
}

func (o ReportConfigFilterResponsePtrOutput) Elem() ReportConfigFilterResponseOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) ReportConfigFilterResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigFilterResponse
		return ret
	}).(ReportConfigFilterResponseOutput)
}

func (o ReportConfigFilterResponsePtrOutput) And() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) []ReportConfigFilterResponse {
		if v == nil {
			return nil
		}
		return v.And
	}).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponsePtrOutput) Dimensions() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Dimensions
	}).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponsePtrOutput) Or() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) []ReportConfigFilterResponse {
		if v == nil {
			return nil
		}
		return v.Or
	}).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponsePtrOutput) TagKey() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.TagKey
	}).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponsePtrOutput) TagValue() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.TagValue
	}).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponsePtrOutput) Tags() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(ReportConfigComparisonExpressionResponsePtrOutput)
}

type ReportConfigFilterResponseArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigFilterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigFilterResponse)(nil)).Elem()
}

func (o ReportConfigFilterResponseArrayOutput) ToReportConfigFilterResponseArrayOutput() ReportConfigFilterResponseArrayOutput {
	return o
}

func (o ReportConfigFilterResponseArrayOutput) ToReportConfigFilterResponseArrayOutputWithContext(ctx context.Context) ReportConfigFilterResponseArrayOutput {
	return o
}

func (o ReportConfigFilterResponseArrayOutput) Index(i pulumi.IntInput) ReportConfigFilterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigFilterResponse {
		return vs[0].([]ReportConfigFilterResponse)[vs[1].(int)]
	}).(ReportConfigFilterResponseOutput)
}

type ReportConfigGrouping struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type ReportConfigGroupingInput interface {
	pulumi.Input

	ToReportConfigGroupingOutput() ReportConfigGroupingOutput
	ToReportConfigGroupingOutputWithContext(context.Context) ReportConfigGroupingOutput
}

type ReportConfigGroupingArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ReportConfigGroupingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigGrouping)(nil)).Elem()
}

func (i ReportConfigGroupingArgs) ToReportConfigGroupingOutput() ReportConfigGroupingOutput {
	return i.ToReportConfigGroupingOutputWithContext(context.Background())
}

func (i ReportConfigGroupingArgs) ToReportConfigGroupingOutputWithContext(ctx context.Context) ReportConfigGroupingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigGroupingOutput)
}





type ReportConfigGroupingArrayInput interface {
	pulumi.Input

	ToReportConfigGroupingArrayOutput() ReportConfigGroupingArrayOutput
	ToReportConfigGroupingArrayOutputWithContext(context.Context) ReportConfigGroupingArrayOutput
}

type ReportConfigGroupingArray []ReportConfigGroupingInput

func (ReportConfigGroupingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigGrouping)(nil)).Elem()
}

func (i ReportConfigGroupingArray) ToReportConfigGroupingArrayOutput() ReportConfigGroupingArrayOutput {
	return i.ToReportConfigGroupingArrayOutputWithContext(context.Background())
}

func (i ReportConfigGroupingArray) ToReportConfigGroupingArrayOutputWithContext(ctx context.Context) ReportConfigGroupingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigGroupingArrayOutput)
}

type ReportConfigGroupingOutput struct{ *pulumi.OutputState }

func (ReportConfigGroupingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigGrouping)(nil)).Elem()
}

func (o ReportConfigGroupingOutput) ToReportConfigGroupingOutput() ReportConfigGroupingOutput {
	return o
}

func (o ReportConfigGroupingOutput) ToReportConfigGroupingOutputWithContext(ctx context.Context) ReportConfigGroupingOutput {
	return o
}

func (o ReportConfigGroupingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGrouping) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportConfigGroupingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGrouping) string { return v.Type }).(pulumi.StringOutput)
}

type ReportConfigGroupingArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigGroupingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigGrouping)(nil)).Elem()
}

func (o ReportConfigGroupingArrayOutput) ToReportConfigGroupingArrayOutput() ReportConfigGroupingArrayOutput {
	return o
}

func (o ReportConfigGroupingArrayOutput) ToReportConfigGroupingArrayOutputWithContext(ctx context.Context) ReportConfigGroupingArrayOutput {
	return o
}

func (o ReportConfigGroupingArrayOutput) Index(i pulumi.IntInput) ReportConfigGroupingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigGrouping {
		return vs[0].([]ReportConfigGrouping)[vs[1].(int)]
	}).(ReportConfigGroupingOutput)
}

type ReportConfigGroupingResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type ReportConfigGroupingResponseInput interface {
	pulumi.Input

	ToReportConfigGroupingResponseOutput() ReportConfigGroupingResponseOutput
	ToReportConfigGroupingResponseOutputWithContext(context.Context) ReportConfigGroupingResponseOutput
}

type ReportConfigGroupingResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ReportConfigGroupingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigGroupingResponse)(nil)).Elem()
}

func (i ReportConfigGroupingResponseArgs) ToReportConfigGroupingResponseOutput() ReportConfigGroupingResponseOutput {
	return i.ToReportConfigGroupingResponseOutputWithContext(context.Background())
}

func (i ReportConfigGroupingResponseArgs) ToReportConfigGroupingResponseOutputWithContext(ctx context.Context) ReportConfigGroupingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigGroupingResponseOutput)
}





type ReportConfigGroupingResponseArrayInput interface {
	pulumi.Input

	ToReportConfigGroupingResponseArrayOutput() ReportConfigGroupingResponseArrayOutput
	ToReportConfigGroupingResponseArrayOutputWithContext(context.Context) ReportConfigGroupingResponseArrayOutput
}

type ReportConfigGroupingResponseArray []ReportConfigGroupingResponseInput

func (ReportConfigGroupingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigGroupingResponse)(nil)).Elem()
}

func (i ReportConfigGroupingResponseArray) ToReportConfigGroupingResponseArrayOutput() ReportConfigGroupingResponseArrayOutput {
	return i.ToReportConfigGroupingResponseArrayOutputWithContext(context.Background())
}

func (i ReportConfigGroupingResponseArray) ToReportConfigGroupingResponseArrayOutputWithContext(ctx context.Context) ReportConfigGroupingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigGroupingResponseArrayOutput)
}

type ReportConfigGroupingResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigGroupingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigGroupingResponse)(nil)).Elem()
}

func (o ReportConfigGroupingResponseOutput) ToReportConfigGroupingResponseOutput() ReportConfigGroupingResponseOutput {
	return o
}

func (o ReportConfigGroupingResponseOutput) ToReportConfigGroupingResponseOutputWithContext(ctx context.Context) ReportConfigGroupingResponseOutput {
	return o
}

func (o ReportConfigGroupingResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGroupingResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportConfigGroupingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigGroupingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ReportConfigGroupingResponseArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigGroupingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigGroupingResponse)(nil)).Elem()
}

func (o ReportConfigGroupingResponseArrayOutput) ToReportConfigGroupingResponseArrayOutput() ReportConfigGroupingResponseArrayOutput {
	return o
}

func (o ReportConfigGroupingResponseArrayOutput) ToReportConfigGroupingResponseArrayOutputWithContext(ctx context.Context) ReportConfigGroupingResponseArrayOutput {
	return o
}

func (o ReportConfigGroupingResponseArrayOutput) Index(i pulumi.IntInput) ReportConfigGroupingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigGroupingResponse {
		return vs[0].([]ReportConfigGroupingResponse)[vs[1].(int)]
	}).(ReportConfigGroupingResponseOutput)
}

type ReportConfigSorting struct {
	Direction *string `pulumi:"direction"`
	Name      string  `pulumi:"name"`
}





type ReportConfigSortingInput interface {
	pulumi.Input

	ToReportConfigSortingOutput() ReportConfigSortingOutput
	ToReportConfigSortingOutputWithContext(context.Context) ReportConfigSortingOutput
}

type ReportConfigSortingArgs struct {
	Direction pulumi.StringPtrInput `pulumi:"direction"`
	Name      pulumi.StringInput    `pulumi:"name"`
}

func (ReportConfigSortingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigSorting)(nil)).Elem()
}

func (i ReportConfigSortingArgs) ToReportConfigSortingOutput() ReportConfigSortingOutput {
	return i.ToReportConfigSortingOutputWithContext(context.Background())
}

func (i ReportConfigSortingArgs) ToReportConfigSortingOutputWithContext(ctx context.Context) ReportConfigSortingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigSortingOutput)
}





type ReportConfigSortingArrayInput interface {
	pulumi.Input

	ToReportConfigSortingArrayOutput() ReportConfigSortingArrayOutput
	ToReportConfigSortingArrayOutputWithContext(context.Context) ReportConfigSortingArrayOutput
}

type ReportConfigSortingArray []ReportConfigSortingInput

func (ReportConfigSortingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigSorting)(nil)).Elem()
}

func (i ReportConfigSortingArray) ToReportConfigSortingArrayOutput() ReportConfigSortingArrayOutput {
	return i.ToReportConfigSortingArrayOutputWithContext(context.Background())
}

func (i ReportConfigSortingArray) ToReportConfigSortingArrayOutputWithContext(ctx context.Context) ReportConfigSortingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigSortingArrayOutput)
}

type ReportConfigSortingOutput struct{ *pulumi.OutputState }

func (ReportConfigSortingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigSorting)(nil)).Elem()
}

func (o ReportConfigSortingOutput) ToReportConfigSortingOutput() ReportConfigSortingOutput {
	return o
}

func (o ReportConfigSortingOutput) ToReportConfigSortingOutputWithContext(ctx context.Context) ReportConfigSortingOutput {
	return o
}

func (o ReportConfigSortingOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigSorting) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o ReportConfigSortingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigSorting) string { return v.Name }).(pulumi.StringOutput)
}

type ReportConfigSortingArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigSortingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigSorting)(nil)).Elem()
}

func (o ReportConfigSortingArrayOutput) ToReportConfigSortingArrayOutput() ReportConfigSortingArrayOutput {
	return o
}

func (o ReportConfigSortingArrayOutput) ToReportConfigSortingArrayOutputWithContext(ctx context.Context) ReportConfigSortingArrayOutput {
	return o
}

func (o ReportConfigSortingArrayOutput) Index(i pulumi.IntInput) ReportConfigSortingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigSorting {
		return vs[0].([]ReportConfigSorting)[vs[1].(int)]
	}).(ReportConfigSortingOutput)
}

type ReportConfigSortingResponse struct {
	Direction *string `pulumi:"direction"`
	Name      string  `pulumi:"name"`
}





type ReportConfigSortingResponseInput interface {
	pulumi.Input

	ToReportConfigSortingResponseOutput() ReportConfigSortingResponseOutput
	ToReportConfigSortingResponseOutputWithContext(context.Context) ReportConfigSortingResponseOutput
}

type ReportConfigSortingResponseArgs struct {
	Direction pulumi.StringPtrInput `pulumi:"direction"`
	Name      pulumi.StringInput    `pulumi:"name"`
}

func (ReportConfigSortingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigSortingResponse)(nil)).Elem()
}

func (i ReportConfigSortingResponseArgs) ToReportConfigSortingResponseOutput() ReportConfigSortingResponseOutput {
	return i.ToReportConfigSortingResponseOutputWithContext(context.Background())
}

func (i ReportConfigSortingResponseArgs) ToReportConfigSortingResponseOutputWithContext(ctx context.Context) ReportConfigSortingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigSortingResponseOutput)
}





type ReportConfigSortingResponseArrayInput interface {
	pulumi.Input

	ToReportConfigSortingResponseArrayOutput() ReportConfigSortingResponseArrayOutput
	ToReportConfigSortingResponseArrayOutputWithContext(context.Context) ReportConfigSortingResponseArrayOutput
}

type ReportConfigSortingResponseArray []ReportConfigSortingResponseInput

func (ReportConfigSortingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigSortingResponse)(nil)).Elem()
}

func (i ReportConfigSortingResponseArray) ToReportConfigSortingResponseArrayOutput() ReportConfigSortingResponseArrayOutput {
	return i.ToReportConfigSortingResponseArrayOutputWithContext(context.Background())
}

func (i ReportConfigSortingResponseArray) ToReportConfigSortingResponseArrayOutputWithContext(ctx context.Context) ReportConfigSortingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigSortingResponseArrayOutput)
}

type ReportConfigSortingResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigSortingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigSortingResponse)(nil)).Elem()
}

func (o ReportConfigSortingResponseOutput) ToReportConfigSortingResponseOutput() ReportConfigSortingResponseOutput {
	return o
}

func (o ReportConfigSortingResponseOutput) ToReportConfigSortingResponseOutputWithContext(ctx context.Context) ReportConfigSortingResponseOutput {
	return o
}

func (o ReportConfigSortingResponseOutput) Direction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportConfigSortingResponse) *string { return v.Direction }).(pulumi.StringPtrOutput)
}

func (o ReportConfigSortingResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigSortingResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ReportConfigSortingResponseArrayOutput struct{ *pulumi.OutputState }

func (ReportConfigSortingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportConfigSortingResponse)(nil)).Elem()
}

func (o ReportConfigSortingResponseArrayOutput) ToReportConfigSortingResponseArrayOutput() ReportConfigSortingResponseArrayOutput {
	return o
}

func (o ReportConfigSortingResponseArrayOutput) ToReportConfigSortingResponseArrayOutputWithContext(ctx context.Context) ReportConfigSortingResponseArrayOutput {
	return o
}

func (o ReportConfigSortingResponseArrayOutput) Index(i pulumi.IntInput) ReportConfigSortingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportConfigSortingResponse {
		return vs[0].([]ReportConfigSortingResponse)[vs[1].(int)]
	}).(ReportConfigSortingResponseOutput)
}

type ReportConfigTimePeriod struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}





type ReportConfigTimePeriodInput interface {
	pulumi.Input

	ToReportConfigTimePeriodOutput() ReportConfigTimePeriodOutput
	ToReportConfigTimePeriodOutputWithContext(context.Context) ReportConfigTimePeriodOutput
}

type ReportConfigTimePeriodArgs struct {
	From pulumi.StringInput `pulumi:"from"`
	To   pulumi.StringInput `pulumi:"to"`
}

func (ReportConfigTimePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigTimePeriod)(nil)).Elem()
}

func (i ReportConfigTimePeriodArgs) ToReportConfigTimePeriodOutput() ReportConfigTimePeriodOutput {
	return i.ToReportConfigTimePeriodOutputWithContext(context.Background())
}

func (i ReportConfigTimePeriodArgs) ToReportConfigTimePeriodOutputWithContext(ctx context.Context) ReportConfigTimePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigTimePeriodOutput)
}

func (i ReportConfigTimePeriodArgs) ToReportConfigTimePeriodPtrOutput() ReportConfigTimePeriodPtrOutput {
	return i.ToReportConfigTimePeriodPtrOutputWithContext(context.Background())
}

func (i ReportConfigTimePeriodArgs) ToReportConfigTimePeriodPtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigTimePeriodOutput).ToReportConfigTimePeriodPtrOutputWithContext(ctx)
}









type ReportConfigTimePeriodPtrInput interface {
	pulumi.Input

	ToReportConfigTimePeriodPtrOutput() ReportConfigTimePeriodPtrOutput
	ToReportConfigTimePeriodPtrOutputWithContext(context.Context) ReportConfigTimePeriodPtrOutput
}

type reportConfigTimePeriodPtrType ReportConfigTimePeriodArgs

func ReportConfigTimePeriodPtr(v *ReportConfigTimePeriodArgs) ReportConfigTimePeriodPtrInput {
	return (*reportConfigTimePeriodPtrType)(v)
}

func (*reportConfigTimePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigTimePeriod)(nil)).Elem()
}

func (i *reportConfigTimePeriodPtrType) ToReportConfigTimePeriodPtrOutput() ReportConfigTimePeriodPtrOutput {
	return i.ToReportConfigTimePeriodPtrOutputWithContext(context.Background())
}

func (i *reportConfigTimePeriodPtrType) ToReportConfigTimePeriodPtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigTimePeriodPtrOutput)
}

type ReportConfigTimePeriodOutput struct{ *pulumi.OutputState }

func (ReportConfigTimePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigTimePeriod)(nil)).Elem()
}

func (o ReportConfigTimePeriodOutput) ToReportConfigTimePeriodOutput() ReportConfigTimePeriodOutput {
	return o
}

func (o ReportConfigTimePeriodOutput) ToReportConfigTimePeriodOutputWithContext(ctx context.Context) ReportConfigTimePeriodOutput {
	return o
}

func (o ReportConfigTimePeriodOutput) ToReportConfigTimePeriodPtrOutput() ReportConfigTimePeriodPtrOutput {
	return o.ToReportConfigTimePeriodPtrOutputWithContext(context.Background())
}

func (o ReportConfigTimePeriodOutput) ToReportConfigTimePeriodPtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigTimePeriod) *ReportConfigTimePeriod {
		return &v
	}).(ReportConfigTimePeriodPtrOutput)
}

func (o ReportConfigTimePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigTimePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportConfigTimePeriodOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigTimePeriod) string { return v.To }).(pulumi.StringOutput)
}

type ReportConfigTimePeriodPtrOutput struct{ *pulumi.OutputState }

func (ReportConfigTimePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigTimePeriod)(nil)).Elem()
}

func (o ReportConfigTimePeriodPtrOutput) ToReportConfigTimePeriodPtrOutput() ReportConfigTimePeriodPtrOutput {
	return o
}

func (o ReportConfigTimePeriodPtrOutput) ToReportConfigTimePeriodPtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodPtrOutput {
	return o
}

func (o ReportConfigTimePeriodPtrOutput) Elem() ReportConfigTimePeriodOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriod) ReportConfigTimePeriod {
		if v != nil {
			return *v
		}
		var ret ReportConfigTimePeriod
		return ret
	}).(ReportConfigTimePeriodOutput)
}

func (o ReportConfigTimePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigTimePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

type ReportConfigTimePeriodResponse struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}





type ReportConfigTimePeriodResponseInput interface {
	pulumi.Input

	ToReportConfigTimePeriodResponseOutput() ReportConfigTimePeriodResponseOutput
	ToReportConfigTimePeriodResponseOutputWithContext(context.Context) ReportConfigTimePeriodResponseOutput
}

type ReportConfigTimePeriodResponseArgs struct {
	From pulumi.StringInput `pulumi:"from"`
	To   pulumi.StringInput `pulumi:"to"`
}

func (ReportConfigTimePeriodResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigTimePeriodResponse)(nil)).Elem()
}

func (i ReportConfigTimePeriodResponseArgs) ToReportConfigTimePeriodResponseOutput() ReportConfigTimePeriodResponseOutput {
	return i.ToReportConfigTimePeriodResponseOutputWithContext(context.Background())
}

func (i ReportConfigTimePeriodResponseArgs) ToReportConfigTimePeriodResponseOutputWithContext(ctx context.Context) ReportConfigTimePeriodResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigTimePeriodResponseOutput)
}

func (i ReportConfigTimePeriodResponseArgs) ToReportConfigTimePeriodResponsePtrOutput() ReportConfigTimePeriodResponsePtrOutput {
	return i.ToReportConfigTimePeriodResponsePtrOutputWithContext(context.Background())
}

func (i ReportConfigTimePeriodResponseArgs) ToReportConfigTimePeriodResponsePtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigTimePeriodResponseOutput).ToReportConfigTimePeriodResponsePtrOutputWithContext(ctx)
}









type ReportConfigTimePeriodResponsePtrInput interface {
	pulumi.Input

	ToReportConfigTimePeriodResponsePtrOutput() ReportConfigTimePeriodResponsePtrOutput
	ToReportConfigTimePeriodResponsePtrOutputWithContext(context.Context) ReportConfigTimePeriodResponsePtrOutput
}

type reportConfigTimePeriodResponsePtrType ReportConfigTimePeriodResponseArgs

func ReportConfigTimePeriodResponsePtr(v *ReportConfigTimePeriodResponseArgs) ReportConfigTimePeriodResponsePtrInput {
	return (*reportConfigTimePeriodResponsePtrType)(v)
}

func (*reportConfigTimePeriodResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigTimePeriodResponse)(nil)).Elem()
}

func (i *reportConfigTimePeriodResponsePtrType) ToReportConfigTimePeriodResponsePtrOutput() ReportConfigTimePeriodResponsePtrOutput {
	return i.ToReportConfigTimePeriodResponsePtrOutputWithContext(context.Background())
}

func (i *reportConfigTimePeriodResponsePtrType) ToReportConfigTimePeriodResponsePtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportConfigTimePeriodResponsePtrOutput)
}

type ReportConfigTimePeriodResponseOutput struct{ *pulumi.OutputState }

func (ReportConfigTimePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigTimePeriodResponse)(nil)).Elem()
}

func (o ReportConfigTimePeriodResponseOutput) ToReportConfigTimePeriodResponseOutput() ReportConfigTimePeriodResponseOutput {
	return o
}

func (o ReportConfigTimePeriodResponseOutput) ToReportConfigTimePeriodResponseOutputWithContext(ctx context.Context) ReportConfigTimePeriodResponseOutput {
	return o
}

func (o ReportConfigTimePeriodResponseOutput) ToReportConfigTimePeriodResponsePtrOutput() ReportConfigTimePeriodResponsePtrOutput {
	return o.ToReportConfigTimePeriodResponsePtrOutputWithContext(context.Background())
}

func (o ReportConfigTimePeriodResponseOutput) ToReportConfigTimePeriodResponsePtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigTimePeriodResponse) *ReportConfigTimePeriodResponse {
		return &v
	}).(ReportConfigTimePeriodResponsePtrOutput)
}

func (o ReportConfigTimePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigTimePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportConfigTimePeriodResponseOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v ReportConfigTimePeriodResponse) string { return v.To }).(pulumi.StringOutput)
}

type ReportConfigTimePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigTimePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigTimePeriodResponse)(nil)).Elem()
}

func (o ReportConfigTimePeriodResponsePtrOutput) ToReportConfigTimePeriodResponsePtrOutput() ReportConfigTimePeriodResponsePtrOutput {
	return o
}

func (o ReportConfigTimePeriodResponsePtrOutput) ToReportConfigTimePeriodResponsePtrOutputWithContext(ctx context.Context) ReportConfigTimePeriodResponsePtrOutput {
	return o
}

func (o ReportConfigTimePeriodResponsePtrOutput) Elem() ReportConfigTimePeriodResponseOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriodResponse) ReportConfigTimePeriodResponse {
		if v != nil {
			return *v
		}
		var ret ReportConfigTimePeriodResponse
		return ret
	}).(ReportConfigTimePeriodResponseOutput)
}

func (o ReportConfigTimePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportConfigTimePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportConfigTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

type SettingsPropertiesCache struct {
	Channel    string  `pulumi:"channel"`
	Id         string  `pulumi:"id"`
	Name       string  `pulumi:"name"`
	Parent     *string `pulumi:"parent"`
	Status     *string `pulumi:"status"`
	Subchannel string  `pulumi:"subchannel"`
}





type SettingsPropertiesCacheInput interface {
	pulumi.Input

	ToSettingsPropertiesCacheOutput() SettingsPropertiesCacheOutput
	ToSettingsPropertiesCacheOutputWithContext(context.Context) SettingsPropertiesCacheOutput
}

type SettingsPropertiesCacheArgs struct {
	Channel    pulumi.StringInput    `pulumi:"channel"`
	Id         pulumi.StringInput    `pulumi:"id"`
	Name       pulumi.StringInput    `pulumi:"name"`
	Parent     pulumi.StringPtrInput `pulumi:"parent"`
	Status     pulumi.StringPtrInput `pulumi:"status"`
	Subchannel pulumi.StringInput    `pulumi:"subchannel"`
}

func (SettingsPropertiesCacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsPropertiesCache)(nil)).Elem()
}

func (i SettingsPropertiesCacheArgs) ToSettingsPropertiesCacheOutput() SettingsPropertiesCacheOutput {
	return i.ToSettingsPropertiesCacheOutputWithContext(context.Background())
}

func (i SettingsPropertiesCacheArgs) ToSettingsPropertiesCacheOutputWithContext(ctx context.Context) SettingsPropertiesCacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsPropertiesCacheOutput)
}





type SettingsPropertiesCacheArrayInput interface {
	pulumi.Input

	ToSettingsPropertiesCacheArrayOutput() SettingsPropertiesCacheArrayOutput
	ToSettingsPropertiesCacheArrayOutputWithContext(context.Context) SettingsPropertiesCacheArrayOutput
}

type SettingsPropertiesCacheArray []SettingsPropertiesCacheInput

func (SettingsPropertiesCacheArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsPropertiesCache)(nil)).Elem()
}

func (i SettingsPropertiesCacheArray) ToSettingsPropertiesCacheArrayOutput() SettingsPropertiesCacheArrayOutput {
	return i.ToSettingsPropertiesCacheArrayOutputWithContext(context.Background())
}

func (i SettingsPropertiesCacheArray) ToSettingsPropertiesCacheArrayOutputWithContext(ctx context.Context) SettingsPropertiesCacheArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsPropertiesCacheArrayOutput)
}

type SettingsPropertiesCacheOutput struct{ *pulumi.OutputState }

func (SettingsPropertiesCacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsPropertiesCache)(nil)).Elem()
}

func (o SettingsPropertiesCacheOutput) ToSettingsPropertiesCacheOutput() SettingsPropertiesCacheOutput {
	return o
}

func (o SettingsPropertiesCacheOutput) ToSettingsPropertiesCacheOutputWithContext(ctx context.Context) SettingsPropertiesCacheOutput {
	return o
}

func (o SettingsPropertiesCacheOutput) Channel() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsPropertiesCache) string { return v.Channel }).(pulumi.StringOutput)
}

func (o SettingsPropertiesCacheOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsPropertiesCache) string { return v.Id }).(pulumi.StringOutput)
}

func (o SettingsPropertiesCacheOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsPropertiesCache) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsPropertiesCacheOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingsPropertiesCache) *string { return v.Parent }).(pulumi.StringPtrOutput)
}

func (o SettingsPropertiesCacheOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingsPropertiesCache) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o SettingsPropertiesCacheOutput) Subchannel() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsPropertiesCache) string { return v.Subchannel }).(pulumi.StringOutput)
}

type SettingsPropertiesCacheArrayOutput struct{ *pulumi.OutputState }

func (SettingsPropertiesCacheArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsPropertiesCache)(nil)).Elem()
}

func (o SettingsPropertiesCacheArrayOutput) ToSettingsPropertiesCacheArrayOutput() SettingsPropertiesCacheArrayOutput {
	return o
}

func (o SettingsPropertiesCacheArrayOutput) ToSettingsPropertiesCacheArrayOutputWithContext(ctx context.Context) SettingsPropertiesCacheArrayOutput {
	return o
}

func (o SettingsPropertiesCacheArrayOutput) Index(i pulumi.IntInput) SettingsPropertiesCacheOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsPropertiesCache {
		return vs[0].([]SettingsPropertiesCache)[vs[1].(int)]
	}).(SettingsPropertiesCacheOutput)
}

type SettingsPropertiesResponseCache struct {
	Channel    string  `pulumi:"channel"`
	Id         string  `pulumi:"id"`
	Name       string  `pulumi:"name"`
	Parent     *string `pulumi:"parent"`
	Status     *string `pulumi:"status"`
	Subchannel string  `pulumi:"subchannel"`
}





type SettingsPropertiesResponseCacheInput interface {
	pulumi.Input

	ToSettingsPropertiesResponseCacheOutput() SettingsPropertiesResponseCacheOutput
	ToSettingsPropertiesResponseCacheOutputWithContext(context.Context) SettingsPropertiesResponseCacheOutput
}

type SettingsPropertiesResponseCacheArgs struct {
	Channel    pulumi.StringInput    `pulumi:"channel"`
	Id         pulumi.StringInput    `pulumi:"id"`
	Name       pulumi.StringInput    `pulumi:"name"`
	Parent     pulumi.StringPtrInput `pulumi:"parent"`
	Status     pulumi.StringPtrInput `pulumi:"status"`
	Subchannel pulumi.StringInput    `pulumi:"subchannel"`
}

func (SettingsPropertiesResponseCacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsPropertiesResponseCache)(nil)).Elem()
}

func (i SettingsPropertiesResponseCacheArgs) ToSettingsPropertiesResponseCacheOutput() SettingsPropertiesResponseCacheOutput {
	return i.ToSettingsPropertiesResponseCacheOutputWithContext(context.Background())
}

func (i SettingsPropertiesResponseCacheArgs) ToSettingsPropertiesResponseCacheOutputWithContext(ctx context.Context) SettingsPropertiesResponseCacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsPropertiesResponseCacheOutput)
}





type SettingsPropertiesResponseCacheArrayInput interface {
	pulumi.Input

	ToSettingsPropertiesResponseCacheArrayOutput() SettingsPropertiesResponseCacheArrayOutput
	ToSettingsPropertiesResponseCacheArrayOutputWithContext(context.Context) SettingsPropertiesResponseCacheArrayOutput
}

type SettingsPropertiesResponseCacheArray []SettingsPropertiesResponseCacheInput

func (SettingsPropertiesResponseCacheArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsPropertiesResponseCache)(nil)).Elem()
}

func (i SettingsPropertiesResponseCacheArray) ToSettingsPropertiesResponseCacheArrayOutput() SettingsPropertiesResponseCacheArrayOutput {
	return i.ToSettingsPropertiesResponseCacheArrayOutputWithContext(context.Background())
}

func (i SettingsPropertiesResponseCacheArray) ToSettingsPropertiesResponseCacheArrayOutputWithContext(ctx context.Context) SettingsPropertiesResponseCacheArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingsPropertiesResponseCacheArrayOutput)
}

type SettingsPropertiesResponseCacheOutput struct{ *pulumi.OutputState }

func (SettingsPropertiesResponseCacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingsPropertiesResponseCache)(nil)).Elem()
}

func (o SettingsPropertiesResponseCacheOutput) ToSettingsPropertiesResponseCacheOutput() SettingsPropertiesResponseCacheOutput {
	return o
}

func (o SettingsPropertiesResponseCacheOutput) ToSettingsPropertiesResponseCacheOutputWithContext(ctx context.Context) SettingsPropertiesResponseCacheOutput {
	return o
}

func (o SettingsPropertiesResponseCacheOutput) Channel() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsPropertiesResponseCache) string { return v.Channel }).(pulumi.StringOutput)
}

func (o SettingsPropertiesResponseCacheOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsPropertiesResponseCache) string { return v.Id }).(pulumi.StringOutput)
}

func (o SettingsPropertiesResponseCacheOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsPropertiesResponseCache) string { return v.Name }).(pulumi.StringOutput)
}

func (o SettingsPropertiesResponseCacheOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingsPropertiesResponseCache) *string { return v.Parent }).(pulumi.StringPtrOutput)
}

func (o SettingsPropertiesResponseCacheOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingsPropertiesResponseCache) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o SettingsPropertiesResponseCacheOutput) Subchannel() pulumi.StringOutput {
	return o.ApplyT(func(v SettingsPropertiesResponseCache) string { return v.Subchannel }).(pulumi.StringOutput)
}

type SettingsPropertiesResponseCacheArrayOutput struct{ *pulumi.OutputState }

func (SettingsPropertiesResponseCacheArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingsPropertiesResponseCache)(nil)).Elem()
}

func (o SettingsPropertiesResponseCacheArrayOutput) ToSettingsPropertiesResponseCacheArrayOutput() SettingsPropertiesResponseCacheArrayOutput {
	return o
}

func (o SettingsPropertiesResponseCacheArrayOutput) ToSettingsPropertiesResponseCacheArrayOutputWithContext(ctx context.Context) SettingsPropertiesResponseCacheArrayOutput {
	return o
}

func (o SettingsPropertiesResponseCacheArrayOutput) Index(i pulumi.IntInput) SettingsPropertiesResponseCacheOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingsPropertiesResponseCache {
		return vs[0].([]SettingsPropertiesResponseCache)[vs[1].(int)]
	}).(SettingsPropertiesResponseCacheOutput)
}

func init() {
	pulumi.RegisterOutputType(ExportDefinitionOutput{})
	pulumi.RegisterOutputType(ExportDefinitionPtrOutput{})
	pulumi.RegisterOutputType(ExportDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ExportDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportDeliveryDestinationOutput{})
	pulumi.RegisterOutputType(ExportDeliveryDestinationPtrOutput{})
	pulumi.RegisterOutputType(ExportDeliveryDestinationResponseOutput{})
	pulumi.RegisterOutputType(ExportDeliveryDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoPtrOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoResponseOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodPtrOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodResponseOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportScheduleOutput{})
	pulumi.RegisterOutputType(ExportSchedulePtrOutput{})
	pulumi.RegisterOutputType(ExportScheduleResponseOutput{})
	pulumi.RegisterOutputType(ExportScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(KpiPropertiesOutput{})
	pulumi.RegisterOutputType(KpiPropertiesArrayOutput{})
	pulumi.RegisterOutputType(KpiPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KpiPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(PivotPropertiesOutput{})
	pulumi.RegisterOutputType(PivotPropertiesArrayOutput{})
	pulumi.RegisterOutputType(PivotPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PivotPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(QueryAggregationOutput{})
	pulumi.RegisterOutputType(QueryAggregationMapOutput{})
	pulumi.RegisterOutputType(QueryAggregationResponseOutput{})
	pulumi.RegisterOutputType(QueryAggregationResponseMapOutput{})
	pulumi.RegisterOutputType(QueryComparisonExpressionOutput{})
	pulumi.RegisterOutputType(QueryComparisonExpressionPtrOutput{})
	pulumi.RegisterOutputType(QueryComparisonExpressionResponseOutput{})
	pulumi.RegisterOutputType(QueryComparisonExpressionResponsePtrOutput{})
	pulumi.RegisterOutputType(QueryDatasetOutput{})
	pulumi.RegisterOutputType(QueryDatasetPtrOutput{})
	pulumi.RegisterOutputType(QueryDatasetConfigurationOutput{})
	pulumi.RegisterOutputType(QueryDatasetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(QueryDatasetConfigurationResponseOutput{})
	pulumi.RegisterOutputType(QueryDatasetConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(QueryDatasetResponseOutput{})
	pulumi.RegisterOutputType(QueryDatasetResponsePtrOutput{})
	pulumi.RegisterOutputType(QueryFilterOutput{})
	pulumi.RegisterOutputType(QueryFilterPtrOutput{})
	pulumi.RegisterOutputType(QueryFilterArrayOutput{})
	pulumi.RegisterOutputType(QueryFilterResponseOutput{})
	pulumi.RegisterOutputType(QueryFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(QueryFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(QueryGroupingOutput{})
	pulumi.RegisterOutputType(QueryGroupingArrayOutput{})
	pulumi.RegisterOutputType(QueryGroupingResponseOutput{})
	pulumi.RegisterOutputType(QueryGroupingResponseArrayOutput{})
	pulumi.RegisterOutputType(QueryTimePeriodOutput{})
	pulumi.RegisterOutputType(QueryTimePeriodPtrOutput{})
	pulumi.RegisterOutputType(QueryTimePeriodResponseOutput{})
	pulumi.RegisterOutputType(QueryTimePeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigAggregationOutput{})
	pulumi.RegisterOutputType(ReportConfigAggregationMapOutput{})
	pulumi.RegisterOutputType(ReportConfigAggregationResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigAggregationResponseMapOutput{})
	pulumi.RegisterOutputType(ReportConfigComparisonExpressionOutput{})
	pulumi.RegisterOutputType(ReportConfigComparisonExpressionPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigComparisonExpressionResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigComparisonExpressionResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetConfigurationOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigDatasetResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigGroupingOutput{})
	pulumi.RegisterOutputType(ReportConfigGroupingArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigGroupingResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigGroupingResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigSortingOutput{})
	pulumi.RegisterOutputType(ReportConfigSortingArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigSortingResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigSortingResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodPtrOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodResponseOutput{})
	pulumi.RegisterOutputType(ReportConfigTimePeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(SettingsPropertiesCacheOutput{})
	pulumi.RegisterOutputType(SettingsPropertiesCacheArrayOutput{})
	pulumi.RegisterOutputType(SettingsPropertiesResponseCacheOutput{})
	pulumi.RegisterOutputType(SettingsPropertiesResponseCacheArrayOutput{})
}
