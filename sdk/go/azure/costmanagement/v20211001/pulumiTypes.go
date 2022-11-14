


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CommonExportPropertiesResponse struct {
	Definition          ExportDefinitionResponse           `pulumi:"definition"`
	DeliveryInfo        ExportDeliveryInfoResponse         `pulumi:"deliveryInfo"`
	Format              *string                            `pulumi:"format"`
	NextRunTimeEstimate string                             `pulumi:"nextRunTimeEstimate"`
	PartitionData       *bool                              `pulumi:"partitionData"`
	RunHistory          *ExportExecutionListResultResponse `pulumi:"runHistory"`
}

type CommonExportPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CommonExportPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CommonExportPropertiesResponse)(nil)).Elem()
}

func (o CommonExportPropertiesResponseOutput) ToCommonExportPropertiesResponseOutput() CommonExportPropertiesResponseOutput {
	return o
}

func (o CommonExportPropertiesResponseOutput) ToCommonExportPropertiesResponseOutputWithContext(ctx context.Context) CommonExportPropertiesResponseOutput {
	return o
}

func (o CommonExportPropertiesResponseOutput) Definition() ExportDefinitionResponseOutput {
	return o.ApplyT(func(v CommonExportPropertiesResponse) ExportDefinitionResponse { return v.Definition }).(ExportDefinitionResponseOutput)
}

func (o CommonExportPropertiesResponseOutput) DeliveryInfo() ExportDeliveryInfoResponseOutput {
	return o.ApplyT(func(v CommonExportPropertiesResponse) ExportDeliveryInfoResponse { return v.DeliveryInfo }).(ExportDeliveryInfoResponseOutput)
}

func (o CommonExportPropertiesResponseOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CommonExportPropertiesResponse) *string { return v.Format }).(pulumi.StringPtrOutput)
}

func (o CommonExportPropertiesResponseOutput) NextRunTimeEstimate() pulumi.StringOutput {
	return o.ApplyT(func(v CommonExportPropertiesResponse) string { return v.NextRunTimeEstimate }).(pulumi.StringOutput)
}

func (o CommonExportPropertiesResponseOutput) PartitionData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CommonExportPropertiesResponse) *bool { return v.PartitionData }).(pulumi.BoolPtrOutput)
}

func (o CommonExportPropertiesResponseOutput) RunHistory() ExportExecutionListResultResponsePtrOutput {
	return o.ApplyT(func(v CommonExportPropertiesResponse) *ExportExecutionListResultResponse { return v.RunHistory }).(ExportExecutionListResultResponsePtrOutput)
}

type CommonExportPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CommonExportPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CommonExportPropertiesResponse)(nil)).Elem()
}

func (o CommonExportPropertiesResponsePtrOutput) ToCommonExportPropertiesResponsePtrOutput() CommonExportPropertiesResponsePtrOutput {
	return o
}

func (o CommonExportPropertiesResponsePtrOutput) ToCommonExportPropertiesResponsePtrOutputWithContext(ctx context.Context) CommonExportPropertiesResponsePtrOutput {
	return o
}

func (o CommonExportPropertiesResponsePtrOutput) Elem() CommonExportPropertiesResponseOutput {
	return o.ApplyT(func(v *CommonExportPropertiesResponse) CommonExportPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CommonExportPropertiesResponse
		return ret
	}).(CommonExportPropertiesResponseOutput)
}

func (o CommonExportPropertiesResponsePtrOutput) Definition() ExportDefinitionResponsePtrOutput {
	return o.ApplyT(func(v *CommonExportPropertiesResponse) *ExportDefinitionResponse {
		if v == nil {
			return nil
		}
		return &v.Definition
	}).(ExportDefinitionResponsePtrOutput)
}

func (o CommonExportPropertiesResponsePtrOutput) DeliveryInfo() ExportDeliveryInfoResponsePtrOutput {
	return o.ApplyT(func(v *CommonExportPropertiesResponse) *ExportDeliveryInfoResponse {
		if v == nil {
			return nil
		}
		return &v.DeliveryInfo
	}).(ExportDeliveryInfoResponsePtrOutput)
}

func (o CommonExportPropertiesResponsePtrOutput) Format() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommonExportPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Format
	}).(pulumi.StringPtrOutput)
}

func (o CommonExportPropertiesResponsePtrOutput) NextRunTimeEstimate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CommonExportPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.NextRunTimeEstimate
	}).(pulumi.StringPtrOutput)
}

func (o CommonExportPropertiesResponsePtrOutput) PartitionData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CommonExportPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PartitionData
	}).(pulumi.BoolPtrOutput)
}

func (o CommonExportPropertiesResponsePtrOutput) RunHistory() ExportExecutionListResultResponsePtrOutput {
	return o.ApplyT(func(v *CommonExportPropertiesResponse) *ExportExecutionListResultResponse {
		if v == nil {
			return nil
		}
		return v.RunHistory
	}).(ExportExecutionListResultResponsePtrOutput)
}

type ErrorDetailsResponse struct {
	Code    string `pulumi:"code"`
	Message string `pulumi:"message"`
}

type ErrorDetailsResponseOutput struct{ *pulumi.OutputState }

func (ErrorDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetailsResponse)(nil)).Elem()
}

func (o ErrorDetailsResponseOutput) ToErrorDetailsResponseOutput() ErrorDetailsResponseOutput {
	return o
}

func (o ErrorDetailsResponseOutput) ToErrorDetailsResponseOutputWithContext(ctx context.Context) ErrorDetailsResponseOutput {
	return o
}

func (o ErrorDetailsResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailsResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o ErrorDetailsResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetailsResponse) string { return v.Message }).(pulumi.StringOutput)
}

type ErrorDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorDetailsResponse)(nil)).Elem()
}

func (o ErrorDetailsResponsePtrOutput) ToErrorDetailsResponsePtrOutput() ErrorDetailsResponsePtrOutput {
	return o
}

func (o ErrorDetailsResponsePtrOutput) ToErrorDetailsResponsePtrOutputWithContext(ctx context.Context) ErrorDetailsResponsePtrOutput {
	return o
}

func (o ErrorDetailsResponsePtrOutput) Elem() ErrorDetailsResponseOutput {
	return o.ApplyT(func(v *ErrorDetailsResponse) ErrorDetailsResponse {
		if v != nil {
			return *v
		}
		var ret ErrorDetailsResponse
		return ret
	}).(ErrorDetailsResponseOutput)
}

func (o ErrorDetailsResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorDetailsResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

type ExportDataset struct {
	Configuration *ExportDatasetConfiguration `pulumi:"configuration"`
	Granularity   *string                     `pulumi:"granularity"`
}





type ExportDatasetInput interface {
	pulumi.Input

	ToExportDatasetOutput() ExportDatasetOutput
	ToExportDatasetOutputWithContext(context.Context) ExportDatasetOutput
}

type ExportDatasetArgs struct {
	Configuration ExportDatasetConfigurationPtrInput `pulumi:"configuration"`
	Granularity   pulumi.StringPtrInput              `pulumi:"granularity"`
}

func (ExportDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDataset)(nil)).Elem()
}

func (i ExportDatasetArgs) ToExportDatasetOutput() ExportDatasetOutput {
	return i.ToExportDatasetOutputWithContext(context.Background())
}

func (i ExportDatasetArgs) ToExportDatasetOutputWithContext(ctx context.Context) ExportDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDatasetOutput)
}

func (i ExportDatasetArgs) ToExportDatasetPtrOutput() ExportDatasetPtrOutput {
	return i.ToExportDatasetPtrOutputWithContext(context.Background())
}

func (i ExportDatasetArgs) ToExportDatasetPtrOutputWithContext(ctx context.Context) ExportDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDatasetOutput).ToExportDatasetPtrOutputWithContext(ctx)
}









type ExportDatasetPtrInput interface {
	pulumi.Input

	ToExportDatasetPtrOutput() ExportDatasetPtrOutput
	ToExportDatasetPtrOutputWithContext(context.Context) ExportDatasetPtrOutput
}

type exportDatasetPtrType ExportDatasetArgs

func ExportDatasetPtr(v *ExportDatasetArgs) ExportDatasetPtrInput {
	return (*exportDatasetPtrType)(v)
}

func (*exportDatasetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDataset)(nil)).Elem()
}

func (i *exportDatasetPtrType) ToExportDatasetPtrOutput() ExportDatasetPtrOutput {
	return i.ToExportDatasetPtrOutputWithContext(context.Background())
}

func (i *exportDatasetPtrType) ToExportDatasetPtrOutputWithContext(ctx context.Context) ExportDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDatasetPtrOutput)
}

type ExportDatasetOutput struct{ *pulumi.OutputState }

func (ExportDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDataset)(nil)).Elem()
}

func (o ExportDatasetOutput) ToExportDatasetOutput() ExportDatasetOutput {
	return o
}

func (o ExportDatasetOutput) ToExportDatasetOutputWithContext(ctx context.Context) ExportDatasetOutput {
	return o
}

func (o ExportDatasetOutput) ToExportDatasetPtrOutput() ExportDatasetPtrOutput {
	return o.ToExportDatasetPtrOutputWithContext(context.Background())
}

func (o ExportDatasetOutput) ToExportDatasetPtrOutputWithContext(ctx context.Context) ExportDatasetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDataset) *ExportDataset {
		return &v
	}).(ExportDatasetPtrOutput)
}

func (o ExportDatasetOutput) Configuration() ExportDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v ExportDataset) *ExportDatasetConfiguration { return v.Configuration }).(ExportDatasetConfigurationPtrOutput)
}

func (o ExportDatasetOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDataset) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

type ExportDatasetPtrOutput struct{ *pulumi.OutputState }

func (ExportDatasetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDataset)(nil)).Elem()
}

func (o ExportDatasetPtrOutput) ToExportDatasetPtrOutput() ExportDatasetPtrOutput {
	return o
}

func (o ExportDatasetPtrOutput) ToExportDatasetPtrOutputWithContext(ctx context.Context) ExportDatasetPtrOutput {
	return o
}

func (o ExportDatasetPtrOutput) Elem() ExportDatasetOutput {
	return o.ApplyT(func(v *ExportDataset) ExportDataset {
		if v != nil {
			return *v
		}
		var ret ExportDataset
		return ret
	}).(ExportDatasetOutput)
}

func (o ExportDatasetPtrOutput) Configuration() ExportDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v *ExportDataset) *ExportDatasetConfiguration {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ExportDatasetConfigurationPtrOutput)
}

func (o ExportDatasetPtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDataset) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

type ExportDatasetConfiguration struct {
	Columns []string `pulumi:"columns"`
}





type ExportDatasetConfigurationInput interface {
	pulumi.Input

	ToExportDatasetConfigurationOutput() ExportDatasetConfigurationOutput
	ToExportDatasetConfigurationOutputWithContext(context.Context) ExportDatasetConfigurationOutput
}

type ExportDatasetConfigurationArgs struct {
	Columns pulumi.StringArrayInput `pulumi:"columns"`
}

func (ExportDatasetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDatasetConfiguration)(nil)).Elem()
}

func (i ExportDatasetConfigurationArgs) ToExportDatasetConfigurationOutput() ExportDatasetConfigurationOutput {
	return i.ToExportDatasetConfigurationOutputWithContext(context.Background())
}

func (i ExportDatasetConfigurationArgs) ToExportDatasetConfigurationOutputWithContext(ctx context.Context) ExportDatasetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDatasetConfigurationOutput)
}

func (i ExportDatasetConfigurationArgs) ToExportDatasetConfigurationPtrOutput() ExportDatasetConfigurationPtrOutput {
	return i.ToExportDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i ExportDatasetConfigurationArgs) ToExportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ExportDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDatasetConfigurationOutput).ToExportDatasetConfigurationPtrOutputWithContext(ctx)
}









type ExportDatasetConfigurationPtrInput interface {
	pulumi.Input

	ToExportDatasetConfigurationPtrOutput() ExportDatasetConfigurationPtrOutput
	ToExportDatasetConfigurationPtrOutputWithContext(context.Context) ExportDatasetConfigurationPtrOutput
}

type exportDatasetConfigurationPtrType ExportDatasetConfigurationArgs

func ExportDatasetConfigurationPtr(v *ExportDatasetConfigurationArgs) ExportDatasetConfigurationPtrInput {
	return (*exportDatasetConfigurationPtrType)(v)
}

func (*exportDatasetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDatasetConfiguration)(nil)).Elem()
}

func (i *exportDatasetConfigurationPtrType) ToExportDatasetConfigurationPtrOutput() ExportDatasetConfigurationPtrOutput {
	return i.ToExportDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i *exportDatasetConfigurationPtrType) ToExportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ExportDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportDatasetConfigurationPtrOutput)
}

type ExportDatasetConfigurationOutput struct{ *pulumi.OutputState }

func (ExportDatasetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDatasetConfiguration)(nil)).Elem()
}

func (o ExportDatasetConfigurationOutput) ToExportDatasetConfigurationOutput() ExportDatasetConfigurationOutput {
	return o
}

func (o ExportDatasetConfigurationOutput) ToExportDatasetConfigurationOutputWithContext(ctx context.Context) ExportDatasetConfigurationOutput {
	return o
}

func (o ExportDatasetConfigurationOutput) ToExportDatasetConfigurationPtrOutput() ExportDatasetConfigurationPtrOutput {
	return o.ToExportDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (o ExportDatasetConfigurationOutput) ToExportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ExportDatasetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportDatasetConfiguration) *ExportDatasetConfiguration {
		return &v
	}).(ExportDatasetConfigurationPtrOutput)
}

func (o ExportDatasetConfigurationOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExportDatasetConfiguration) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type ExportDatasetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ExportDatasetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDatasetConfiguration)(nil)).Elem()
}

func (o ExportDatasetConfigurationPtrOutput) ToExportDatasetConfigurationPtrOutput() ExportDatasetConfigurationPtrOutput {
	return o
}

func (o ExportDatasetConfigurationPtrOutput) ToExportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ExportDatasetConfigurationPtrOutput {
	return o
}

func (o ExportDatasetConfigurationPtrOutput) Elem() ExportDatasetConfigurationOutput {
	return o.ApplyT(func(v *ExportDatasetConfiguration) ExportDatasetConfiguration {
		if v != nil {
			return *v
		}
		var ret ExportDatasetConfiguration
		return ret
	}).(ExportDatasetConfigurationOutput)
}

func (o ExportDatasetConfigurationPtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExportDatasetConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type ExportDatasetConfigurationResponse struct {
	Columns []string `pulumi:"columns"`
}

type ExportDatasetConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ExportDatasetConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDatasetConfigurationResponse)(nil)).Elem()
}

func (o ExportDatasetConfigurationResponseOutput) ToExportDatasetConfigurationResponseOutput() ExportDatasetConfigurationResponseOutput {
	return o
}

func (o ExportDatasetConfigurationResponseOutput) ToExportDatasetConfigurationResponseOutputWithContext(ctx context.Context) ExportDatasetConfigurationResponseOutput {
	return o
}

func (o ExportDatasetConfigurationResponseOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ExportDatasetConfigurationResponse) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type ExportDatasetConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportDatasetConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDatasetConfigurationResponse)(nil)).Elem()
}

func (o ExportDatasetConfigurationResponsePtrOutput) ToExportDatasetConfigurationResponsePtrOutput() ExportDatasetConfigurationResponsePtrOutput {
	return o
}

func (o ExportDatasetConfigurationResponsePtrOutput) ToExportDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) ExportDatasetConfigurationResponsePtrOutput {
	return o
}

func (o ExportDatasetConfigurationResponsePtrOutput) Elem() ExportDatasetConfigurationResponseOutput {
	return o.ApplyT(func(v *ExportDatasetConfigurationResponse) ExportDatasetConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ExportDatasetConfigurationResponse
		return ret
	}).(ExportDatasetConfigurationResponseOutput)
}

func (o ExportDatasetConfigurationResponsePtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExportDatasetConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type ExportDatasetResponse struct {
	Configuration *ExportDatasetConfigurationResponse `pulumi:"configuration"`
	Granularity   *string                             `pulumi:"granularity"`
}

type ExportDatasetResponseOutput struct{ *pulumi.OutputState }

func (ExportDatasetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportDatasetResponse)(nil)).Elem()
}

func (o ExportDatasetResponseOutput) ToExportDatasetResponseOutput() ExportDatasetResponseOutput {
	return o
}

func (o ExportDatasetResponseOutput) ToExportDatasetResponseOutputWithContext(ctx context.Context) ExportDatasetResponseOutput {
	return o
}

func (o ExportDatasetResponseOutput) Configuration() ExportDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v ExportDatasetResponse) *ExportDatasetConfigurationResponse { return v.Configuration }).(ExportDatasetConfigurationResponsePtrOutput)
}

func (o ExportDatasetResponseOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDatasetResponse) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

type ExportDatasetResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportDatasetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportDatasetResponse)(nil)).Elem()
}

func (o ExportDatasetResponsePtrOutput) ToExportDatasetResponsePtrOutput() ExportDatasetResponsePtrOutput {
	return o
}

func (o ExportDatasetResponsePtrOutput) ToExportDatasetResponsePtrOutputWithContext(ctx context.Context) ExportDatasetResponsePtrOutput {
	return o
}

func (o ExportDatasetResponsePtrOutput) Elem() ExportDatasetResponseOutput {
	return o.ApplyT(func(v *ExportDatasetResponse) ExportDatasetResponse {
		if v != nil {
			return *v
		}
		var ret ExportDatasetResponse
		return ret
	}).(ExportDatasetResponseOutput)
}

func (o ExportDatasetResponsePtrOutput) Configuration() ExportDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ExportDatasetResponse) *ExportDatasetConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ExportDatasetConfigurationResponsePtrOutput)
}

func (o ExportDatasetResponsePtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDatasetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

type ExportDefinition struct {
	DataSet    *ExportDataset    `pulumi:"dataSet"`
	TimePeriod *ExportTimePeriod `pulumi:"timePeriod"`
	Timeframe  string            `pulumi:"timeframe"`
	Type       string            `pulumi:"type"`
}





type ExportDefinitionInput interface {
	pulumi.Input

	ToExportDefinitionOutput() ExportDefinitionOutput
	ToExportDefinitionOutputWithContext(context.Context) ExportDefinitionOutput
}

type ExportDefinitionArgs struct {
	DataSet    ExportDatasetPtrInput    `pulumi:"dataSet"`
	TimePeriod ExportTimePeriodPtrInput `pulumi:"timePeriod"`
	Timeframe  pulumi.StringInput       `pulumi:"timeframe"`
	Type       pulumi.StringInput       `pulumi:"type"`
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

func (o ExportDefinitionOutput) DataSet() ExportDatasetPtrOutput {
	return o.ApplyT(func(v ExportDefinition) *ExportDataset { return v.DataSet }).(ExportDatasetPtrOutput)
}

func (o ExportDefinitionOutput) TimePeriod() ExportTimePeriodPtrOutput {
	return o.ApplyT(func(v ExportDefinition) *ExportTimePeriod { return v.TimePeriod }).(ExportTimePeriodPtrOutput)
}

func (o ExportDefinitionOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDefinition) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ExportDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type ExportDefinitionResponse struct {
	DataSet    *ExportDatasetResponse    `pulumi:"dataSet"`
	TimePeriod *ExportTimePeriodResponse `pulumi:"timePeriod"`
	Timeframe  string                    `pulumi:"timeframe"`
	Type       string                    `pulumi:"type"`
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

func (o ExportDefinitionResponseOutput) DataSet() ExportDatasetResponsePtrOutput {
	return o.ApplyT(func(v ExportDefinitionResponse) *ExportDatasetResponse { return v.DataSet }).(ExportDatasetResponsePtrOutput)
}

func (o ExportDefinitionResponseOutput) TimePeriod() ExportTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v ExportDefinitionResponse) *ExportTimePeriodResponse { return v.TimePeriod }).(ExportTimePeriodResponsePtrOutput)
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

func (o ExportDefinitionResponsePtrOutput) DataSet() ExportDatasetResponsePtrOutput {
	return o.ApplyT(func(v *ExportDefinitionResponse) *ExportDatasetResponse {
		if v == nil {
			return nil
		}
		return v.DataSet
	}).(ExportDatasetResponsePtrOutput)
}

func (o ExportDefinitionResponsePtrOutput) TimePeriod() ExportTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v *ExportDefinitionResponse) *ExportTimePeriodResponse {
		if v == nil {
			return nil
		}
		return v.TimePeriod
	}).(ExportTimePeriodResponsePtrOutput)
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
	ResourceId     *string `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
	SasToken       *string `pulumi:"sasToken"`
	StorageAccount *string `pulumi:"storageAccount"`
}





type ExportDeliveryDestinationInput interface {
	pulumi.Input

	ToExportDeliveryDestinationOutput() ExportDeliveryDestinationOutput
	ToExportDeliveryDestinationOutputWithContext(context.Context) ExportDeliveryDestinationOutput
}

type ExportDeliveryDestinationArgs struct {
	Container      pulumi.StringInput    `pulumi:"container"`
	ResourceId     pulumi.StringPtrInput `pulumi:"resourceId"`
	RootFolderPath pulumi.StringPtrInput `pulumi:"rootFolderPath"`
	SasToken       pulumi.StringPtrInput `pulumi:"sasToken"`
	StorageAccount pulumi.StringPtrInput `pulumi:"storageAccount"`
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

func (o ExportDeliveryDestinationOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) string { return v.Container }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) *string { return v.SasToken }).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) *string { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

type ExportDeliveryDestinationResponse struct {
	Container      string  `pulumi:"container"`
	ResourceId     *string `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
	SasToken       *string `pulumi:"sasToken"`
	StorageAccount *string `pulumi:"storageAccount"`
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

func (o ExportDeliveryDestinationResponseOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) string { return v.Container }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationResponseOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationResponseOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) *string { return v.SasToken }).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationResponseOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestinationResponse) *string { return v.StorageAccount }).(pulumi.StringPtrOutput)
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
		return v.ResourceId
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

func (o ExportDeliveryDestinationResponsePtrOutput) SasToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasToken
	}).(pulumi.StringPtrOutput)
}

func (o ExportDeliveryDestinationResponsePtrOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportDeliveryDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccount
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

func (o ExportDeliveryInfoOutput) Destination() ExportDeliveryDestinationOutput {
	return o.ApplyT(func(v ExportDeliveryInfo) ExportDeliveryDestination { return v.Destination }).(ExportDeliveryDestinationOutput)
}

type ExportDeliveryInfoResponse struct {
	Destination ExportDeliveryDestinationResponse `pulumi:"destination"`
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

type ExportExecutionListResultResponse struct {
	Value []ExportExecutionResponse `pulumi:"value"`
}

type ExportExecutionListResultResponseOutput struct{ *pulumi.OutputState }

func (ExportExecutionListResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportExecutionListResultResponse)(nil)).Elem()
}

func (o ExportExecutionListResultResponseOutput) ToExportExecutionListResultResponseOutput() ExportExecutionListResultResponseOutput {
	return o
}

func (o ExportExecutionListResultResponseOutput) ToExportExecutionListResultResponseOutputWithContext(ctx context.Context) ExportExecutionListResultResponseOutput {
	return o
}

func (o ExportExecutionListResultResponseOutput) Value() ExportExecutionResponseArrayOutput {
	return o.ApplyT(func(v ExportExecutionListResultResponse) []ExportExecutionResponse { return v.Value }).(ExportExecutionResponseArrayOutput)
}

type ExportExecutionListResultResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportExecutionListResultResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportExecutionListResultResponse)(nil)).Elem()
}

func (o ExportExecutionListResultResponsePtrOutput) ToExportExecutionListResultResponsePtrOutput() ExportExecutionListResultResponsePtrOutput {
	return o
}

func (o ExportExecutionListResultResponsePtrOutput) ToExportExecutionListResultResponsePtrOutputWithContext(ctx context.Context) ExportExecutionListResultResponsePtrOutput {
	return o
}

func (o ExportExecutionListResultResponsePtrOutput) Elem() ExportExecutionListResultResponseOutput {
	return o.ApplyT(func(v *ExportExecutionListResultResponse) ExportExecutionListResultResponse {
		if v != nil {
			return *v
		}
		var ret ExportExecutionListResultResponse
		return ret
	}).(ExportExecutionListResultResponseOutput)
}

func (o ExportExecutionListResultResponsePtrOutput) Value() ExportExecutionResponseArrayOutput {
	return o.ApplyT(func(v *ExportExecutionListResultResponse) []ExportExecutionResponse {
		if v == nil {
			return nil
		}
		return v.Value
	}).(ExportExecutionResponseArrayOutput)
}

type ExportExecutionResponse struct {
	ETag                *string                         `pulumi:"eTag"`
	Error               *ErrorDetailsResponse           `pulumi:"error"`
	ExecutionType       *string                         `pulumi:"executionType"`
	FileName            *string                         `pulumi:"fileName"`
	Id                  string                          `pulumi:"id"`
	Name                string                          `pulumi:"name"`
	ProcessingEndTime   *string                         `pulumi:"processingEndTime"`
	ProcessingStartTime *string                         `pulumi:"processingStartTime"`
	RunSettings         *CommonExportPropertiesResponse `pulumi:"runSettings"`
	Status              *string                         `pulumi:"status"`
	SubmittedBy         *string                         `pulumi:"submittedBy"`
	SubmittedTime       *string                         `pulumi:"submittedTime"`
	Type                string                          `pulumi:"type"`
}

type ExportExecutionResponseOutput struct{ *pulumi.OutputState }

func (ExportExecutionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportExecutionResponse)(nil)).Elem()
}

func (o ExportExecutionResponseOutput) ToExportExecutionResponseOutput() ExportExecutionResponseOutput {
	return o
}

func (o ExportExecutionResponseOutput) ToExportExecutionResponseOutputWithContext(ctx context.Context) ExportExecutionResponseOutput {
	return o
}

func (o ExportExecutionResponseOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportExecutionResponse) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o ExportExecutionResponseOutput) Error() ErrorDetailsResponsePtrOutput {
	return o.ApplyT(func(v ExportExecutionResponse) *ErrorDetailsResponse { return v.Error }).(ErrorDetailsResponsePtrOutput)
}

func (o ExportExecutionResponseOutput) ExecutionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportExecutionResponse) *string { return v.ExecutionType }).(pulumi.StringPtrOutput)
}

func (o ExportExecutionResponseOutput) FileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportExecutionResponse) *string { return v.FileName }).(pulumi.StringPtrOutput)
}

func (o ExportExecutionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ExportExecutionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ExportExecutionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ExportExecutionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ExportExecutionResponseOutput) ProcessingEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportExecutionResponse) *string { return v.ProcessingEndTime }).(pulumi.StringPtrOutput)
}

func (o ExportExecutionResponseOutput) ProcessingStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportExecutionResponse) *string { return v.ProcessingStartTime }).(pulumi.StringPtrOutput)
}

func (o ExportExecutionResponseOutput) RunSettings() CommonExportPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ExportExecutionResponse) *CommonExportPropertiesResponse { return v.RunSettings }).(CommonExportPropertiesResponsePtrOutput)
}

func (o ExportExecutionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportExecutionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ExportExecutionResponseOutput) SubmittedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportExecutionResponse) *string { return v.SubmittedBy }).(pulumi.StringPtrOutput)
}

func (o ExportExecutionResponseOutput) SubmittedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportExecutionResponse) *string { return v.SubmittedTime }).(pulumi.StringPtrOutput)
}

func (o ExportExecutionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ExportExecutionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ExportExecutionResponseArrayOutput struct{ *pulumi.OutputState }

func (ExportExecutionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExportExecutionResponse)(nil)).Elem()
}

func (o ExportExecutionResponseArrayOutput) ToExportExecutionResponseArrayOutput() ExportExecutionResponseArrayOutput {
	return o
}

func (o ExportExecutionResponseArrayOutput) ToExportExecutionResponseArrayOutputWithContext(ctx context.Context) ExportExecutionResponseArrayOutput {
	return o
}

func (o ExportExecutionResponseArrayOutput) Index(i pulumi.IntInput) ExportExecutionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExportExecutionResponse {
		return vs[0].([]ExportExecutionResponse)[vs[1].(int)]
	}).(ExportExecutionResponseOutput)
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
	Recurrence       *string                 `pulumi:"recurrence"`
	RecurrencePeriod *ExportRecurrencePeriod `pulumi:"recurrencePeriod"`
	Status           *string                 `pulumi:"status"`
}





type ExportScheduleInput interface {
	pulumi.Input

	ToExportScheduleOutput() ExportScheduleOutput
	ToExportScheduleOutputWithContext(context.Context) ExportScheduleOutput
}

type ExportScheduleArgs struct {
	Recurrence       pulumi.StringPtrInput          `pulumi:"recurrence"`
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

func (o ExportScheduleOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportSchedule) *string { return v.Recurrence }).(pulumi.StringPtrOutput)
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
		return v.Recurrence
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
	Recurrence       *string                         `pulumi:"recurrence"`
	RecurrencePeriod *ExportRecurrencePeriodResponse `pulumi:"recurrencePeriod"`
	Status           *string                         `pulumi:"status"`
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

func (o ExportScheduleResponseOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportScheduleResponse) *string { return v.Recurrence }).(pulumi.StringPtrOutput)
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
		return v.Recurrence
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

type ExportTimePeriod struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}





type ExportTimePeriodInput interface {
	pulumi.Input

	ToExportTimePeriodOutput() ExportTimePeriodOutput
	ToExportTimePeriodOutputWithContext(context.Context) ExportTimePeriodOutput
}

type ExportTimePeriodArgs struct {
	From pulumi.StringInput `pulumi:"from"`
	To   pulumi.StringInput `pulumi:"to"`
}

func (ExportTimePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportTimePeriod)(nil)).Elem()
}

func (i ExportTimePeriodArgs) ToExportTimePeriodOutput() ExportTimePeriodOutput {
	return i.ToExportTimePeriodOutputWithContext(context.Background())
}

func (i ExportTimePeriodArgs) ToExportTimePeriodOutputWithContext(ctx context.Context) ExportTimePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportTimePeriodOutput)
}

func (i ExportTimePeriodArgs) ToExportTimePeriodPtrOutput() ExportTimePeriodPtrOutput {
	return i.ToExportTimePeriodPtrOutputWithContext(context.Background())
}

func (i ExportTimePeriodArgs) ToExportTimePeriodPtrOutputWithContext(ctx context.Context) ExportTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportTimePeriodOutput).ToExportTimePeriodPtrOutputWithContext(ctx)
}









type ExportTimePeriodPtrInput interface {
	pulumi.Input

	ToExportTimePeriodPtrOutput() ExportTimePeriodPtrOutput
	ToExportTimePeriodPtrOutputWithContext(context.Context) ExportTimePeriodPtrOutput
}

type exportTimePeriodPtrType ExportTimePeriodArgs

func ExportTimePeriodPtr(v *ExportTimePeriodArgs) ExportTimePeriodPtrInput {
	return (*exportTimePeriodPtrType)(v)
}

func (*exportTimePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportTimePeriod)(nil)).Elem()
}

func (i *exportTimePeriodPtrType) ToExportTimePeriodPtrOutput() ExportTimePeriodPtrOutput {
	return i.ToExportTimePeriodPtrOutputWithContext(context.Background())
}

func (i *exportTimePeriodPtrType) ToExportTimePeriodPtrOutputWithContext(ctx context.Context) ExportTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportTimePeriodPtrOutput)
}

type ExportTimePeriodOutput struct{ *pulumi.OutputState }

func (ExportTimePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportTimePeriod)(nil)).Elem()
}

func (o ExportTimePeriodOutput) ToExportTimePeriodOutput() ExportTimePeriodOutput {
	return o
}

func (o ExportTimePeriodOutput) ToExportTimePeriodOutputWithContext(ctx context.Context) ExportTimePeriodOutput {
	return o
}

func (o ExportTimePeriodOutput) ToExportTimePeriodPtrOutput() ExportTimePeriodPtrOutput {
	return o.ToExportTimePeriodPtrOutputWithContext(context.Background())
}

func (o ExportTimePeriodOutput) ToExportTimePeriodPtrOutputWithContext(ctx context.Context) ExportTimePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportTimePeriod) *ExportTimePeriod {
		return &v
	}).(ExportTimePeriodPtrOutput)
}

func (o ExportTimePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ExportTimePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o ExportTimePeriodOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v ExportTimePeriod) string { return v.To }).(pulumi.StringOutput)
}

type ExportTimePeriodPtrOutput struct{ *pulumi.OutputState }

func (ExportTimePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportTimePeriod)(nil)).Elem()
}

func (o ExportTimePeriodPtrOutput) ToExportTimePeriodPtrOutput() ExportTimePeriodPtrOutput {
	return o
}

func (o ExportTimePeriodPtrOutput) ToExportTimePeriodPtrOutputWithContext(ctx context.Context) ExportTimePeriodPtrOutput {
	return o
}

func (o ExportTimePeriodPtrOutput) Elem() ExportTimePeriodOutput {
	return o.ApplyT(func(v *ExportTimePeriod) ExportTimePeriod {
		if v != nil {
			return *v
		}
		var ret ExportTimePeriod
		return ret
	}).(ExportTimePeriodOutput)
}

func (o ExportTimePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ExportTimePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

type ExportTimePeriodResponse struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}

type ExportTimePeriodResponseOutput struct{ *pulumi.OutputState }

func (ExportTimePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportTimePeriodResponse)(nil)).Elem()
}

func (o ExportTimePeriodResponseOutput) ToExportTimePeriodResponseOutput() ExportTimePeriodResponseOutput {
	return o
}

func (o ExportTimePeriodResponseOutput) ToExportTimePeriodResponseOutputWithContext(ctx context.Context) ExportTimePeriodResponseOutput {
	return o
}

func (o ExportTimePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ExportTimePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o ExportTimePeriodResponseOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v ExportTimePeriodResponse) string { return v.To }).(pulumi.StringOutput)
}

type ExportTimePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportTimePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportTimePeriodResponse)(nil)).Elem()
}

func (o ExportTimePeriodResponsePtrOutput) ToExportTimePeriodResponsePtrOutput() ExportTimePeriodResponsePtrOutput {
	return o
}

func (o ExportTimePeriodResponsePtrOutput) ToExportTimePeriodResponsePtrOutputWithContext(ctx context.Context) ExportTimePeriodResponsePtrOutput {
	return o
}

func (o ExportTimePeriodResponsePtrOutput) Elem() ExportTimePeriodResponseOutput {
	return o.ApplyT(func(v *ExportTimePeriodResponse) ExportTimePeriodResponse {
		if v != nil {
			return *v
		}
		var ret ExportTimePeriodResponse
		return ret
	}).(ExportTimePeriodResponseOutput)
}

func (o ExportTimePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ExportTimePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.To
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
	Tags       *ReportConfigComparisonExpressionResponse `pulumi:"tags"`
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

func (o ReportConfigFilterResponseOutput) And() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) []ReportConfigFilterResponse { return v.And }).(ReportConfigFilterResponseArrayOutput)
}

func (o ReportConfigFilterResponseOutput) Dimensions() ReportConfigComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) *ReportConfigComparisonExpressionResponse { return v.Dimensions }).(ReportConfigComparisonExpressionResponsePtrOutput)
}

func (o ReportConfigFilterResponseOutput) Or() ReportConfigFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportConfigFilterResponse) []ReportConfigFilterResponse { return v.Or }).(ReportConfigFilterResponseArrayOutput)
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

func init() {
	pulumi.RegisterOutputType(CommonExportPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CommonExportPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorDetailsResponseOutput{})
	pulumi.RegisterOutputType(ErrorDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportDatasetOutput{})
	pulumi.RegisterOutputType(ExportDatasetPtrOutput{})
	pulumi.RegisterOutputType(ExportDatasetConfigurationOutput{})
	pulumi.RegisterOutputType(ExportDatasetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ExportDatasetConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ExportDatasetConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportDatasetResponseOutput{})
	pulumi.RegisterOutputType(ExportDatasetResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportDefinitionOutput{})
	pulumi.RegisterOutputType(ExportDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ExportDefinitionResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportDeliveryDestinationOutput{})
	pulumi.RegisterOutputType(ExportDeliveryDestinationResponseOutput{})
	pulumi.RegisterOutputType(ExportDeliveryDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoResponseOutput{})
	pulumi.RegisterOutputType(ExportDeliveryInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportExecutionListResultResponseOutput{})
	pulumi.RegisterOutputType(ExportExecutionListResultResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportExecutionResponseOutput{})
	pulumi.RegisterOutputType(ExportExecutionResponseArrayOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodPtrOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodResponseOutput{})
	pulumi.RegisterOutputType(ExportRecurrencePeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportScheduleOutput{})
	pulumi.RegisterOutputType(ExportSchedulePtrOutput{})
	pulumi.RegisterOutputType(ExportScheduleResponseOutput{})
	pulumi.RegisterOutputType(ExportScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportTimePeriodOutput{})
	pulumi.RegisterOutputType(ExportTimePeriodPtrOutput{})
	pulumi.RegisterOutputType(ExportTimePeriodResponseOutput{})
	pulumi.RegisterOutputType(ExportTimePeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(KpiPropertiesOutput{})
	pulumi.RegisterOutputType(KpiPropertiesArrayOutput{})
	pulumi.RegisterOutputType(KpiPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KpiPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(PivotPropertiesOutput{})
	pulumi.RegisterOutputType(PivotPropertiesArrayOutput{})
	pulumi.RegisterOutputType(PivotPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PivotPropertiesResponseArrayOutput{})
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
}
