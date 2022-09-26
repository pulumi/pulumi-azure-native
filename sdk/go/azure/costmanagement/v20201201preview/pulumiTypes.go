


package v20201201preview

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
}
