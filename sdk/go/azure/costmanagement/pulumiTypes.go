


package costmanagement

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

type ConnectorCollectionErrorInfoResponse struct {
	ErrorCode         string `pulumi:"errorCode"`
	ErrorInnerMessage string `pulumi:"errorInnerMessage"`
	ErrorMessage      string `pulumi:"errorMessage"`
	ErrorStartTime    string `pulumi:"errorStartTime"`
}

type ConnectorCollectionErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (ConnectorCollectionErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorCollectionErrorInfoResponse)(nil)).Elem()
}

func (o ConnectorCollectionErrorInfoResponseOutput) ToConnectorCollectionErrorInfoResponseOutput() ConnectorCollectionErrorInfoResponseOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponseOutput) ToConnectorCollectionErrorInfoResponseOutputWithContext(ctx context.Context) ConnectorCollectionErrorInfoResponseOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorCode() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorCode }).(pulumi.StringOutput)
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorInnerMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorInnerMessage }).(pulumi.StringOutput)
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o ConnectorCollectionErrorInfoResponseOutput) ErrorStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionErrorInfoResponse) string { return v.ErrorStartTime }).(pulumi.StringOutput)
}

type ConnectorCollectionErrorInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectorCollectionErrorInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectorCollectionErrorInfoResponse)(nil)).Elem()
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ToConnectorCollectionErrorInfoResponsePtrOutput() ConnectorCollectionErrorInfoResponsePtrOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ToConnectorCollectionErrorInfoResponsePtrOutputWithContext(ctx context.Context) ConnectorCollectionErrorInfoResponsePtrOutput {
	return o
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) Elem() ConnectorCollectionErrorInfoResponseOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) ConnectorCollectionErrorInfoResponse {
		if v != nil {
			return *v
		}
		var ret ConnectorCollectionErrorInfoResponse
		return ret
	}).(ConnectorCollectionErrorInfoResponseOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorCode
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorInnerMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorInnerMessage
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o ConnectorCollectionErrorInfoResponsePtrOutput) ErrorStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectorCollectionErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ErrorStartTime
	}).(pulumi.StringPtrOutput)
}

type ConnectorCollectionInfoResponse struct {
	Error             *ConnectorCollectionErrorInfoResponse `pulumi:"error"`
	LastChecked       string                                `pulumi:"lastChecked"`
	LastUpdated       string                                `pulumi:"lastUpdated"`
	SourceLastUpdated string                                `pulumi:"sourceLastUpdated"`
}

type ConnectorCollectionInfoResponseOutput struct{ *pulumi.OutputState }

func (ConnectorCollectionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectorCollectionInfoResponse)(nil)).Elem()
}

func (o ConnectorCollectionInfoResponseOutput) ToConnectorCollectionInfoResponseOutput() ConnectorCollectionInfoResponseOutput {
	return o
}

func (o ConnectorCollectionInfoResponseOutput) ToConnectorCollectionInfoResponseOutputWithContext(ctx context.Context) ConnectorCollectionInfoResponseOutput {
	return o
}

func (o ConnectorCollectionInfoResponseOutput) Error() ConnectorCollectionErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) *ConnectorCollectionErrorInfoResponse { return v.Error }).(ConnectorCollectionErrorInfoResponsePtrOutput)
}

func (o ConnectorCollectionInfoResponseOutput) LastChecked() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) string { return v.LastChecked }).(pulumi.StringOutput)
}

func (o ConnectorCollectionInfoResponseOutput) LastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) string { return v.LastUpdated }).(pulumi.StringOutput)
}

func (o ConnectorCollectionInfoResponseOutput) SourceLastUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectorCollectionInfoResponse) string { return v.SourceLastUpdated }).(pulumi.StringOutput)
}

type CostAllocationProportion struct {
	Name       string  `pulumi:"name"`
	Percentage float64 `pulumi:"percentage"`
}





type CostAllocationProportionInput interface {
	pulumi.Input

	ToCostAllocationProportionOutput() CostAllocationProportionOutput
	ToCostAllocationProportionOutputWithContext(context.Context) CostAllocationProportionOutput
}

type CostAllocationProportionArgs struct {
	Name       pulumi.StringInput  `pulumi:"name"`
	Percentage pulumi.Float64Input `pulumi:"percentage"`
}

func (CostAllocationProportionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationProportion)(nil)).Elem()
}

func (i CostAllocationProportionArgs) ToCostAllocationProportionOutput() CostAllocationProportionOutput {
	return i.ToCostAllocationProportionOutputWithContext(context.Background())
}

func (i CostAllocationProportionArgs) ToCostAllocationProportionOutputWithContext(ctx context.Context) CostAllocationProportionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationProportionOutput)
}





type CostAllocationProportionArrayInput interface {
	pulumi.Input

	ToCostAllocationProportionArrayOutput() CostAllocationProportionArrayOutput
	ToCostAllocationProportionArrayOutputWithContext(context.Context) CostAllocationProportionArrayOutput
}

type CostAllocationProportionArray []CostAllocationProportionInput

func (CostAllocationProportionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CostAllocationProportion)(nil)).Elem()
}

func (i CostAllocationProportionArray) ToCostAllocationProportionArrayOutput() CostAllocationProportionArrayOutput {
	return i.ToCostAllocationProportionArrayOutputWithContext(context.Background())
}

func (i CostAllocationProportionArray) ToCostAllocationProportionArrayOutputWithContext(ctx context.Context) CostAllocationProportionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationProportionArrayOutput)
}

type CostAllocationProportionOutput struct{ *pulumi.OutputState }

func (CostAllocationProportionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationProportion)(nil)).Elem()
}

func (o CostAllocationProportionOutput) ToCostAllocationProportionOutput() CostAllocationProportionOutput {
	return o
}

func (o CostAllocationProportionOutput) ToCostAllocationProportionOutputWithContext(ctx context.Context) CostAllocationProportionOutput {
	return o
}

func (o CostAllocationProportionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationProportion) string { return v.Name }).(pulumi.StringOutput)
}

func (o CostAllocationProportionOutput) Percentage() pulumi.Float64Output {
	return o.ApplyT(func(v CostAllocationProportion) float64 { return v.Percentage }).(pulumi.Float64Output)
}

type CostAllocationProportionArrayOutput struct{ *pulumi.OutputState }

func (CostAllocationProportionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CostAllocationProportion)(nil)).Elem()
}

func (o CostAllocationProportionArrayOutput) ToCostAllocationProportionArrayOutput() CostAllocationProportionArrayOutput {
	return o
}

func (o CostAllocationProportionArrayOutput) ToCostAllocationProportionArrayOutputWithContext(ctx context.Context) CostAllocationProportionArrayOutput {
	return o
}

func (o CostAllocationProportionArrayOutput) Index(i pulumi.IntInput) CostAllocationProportionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CostAllocationProportion {
		return vs[0].([]CostAllocationProportion)[vs[1].(int)]
	}).(CostAllocationProportionOutput)
}

type CostAllocationProportionResponse struct {
	Name       string  `pulumi:"name"`
	Percentage float64 `pulumi:"percentage"`
}

type CostAllocationProportionResponseOutput struct{ *pulumi.OutputState }

func (CostAllocationProportionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationProportionResponse)(nil)).Elem()
}

func (o CostAllocationProportionResponseOutput) ToCostAllocationProportionResponseOutput() CostAllocationProportionResponseOutput {
	return o
}

func (o CostAllocationProportionResponseOutput) ToCostAllocationProportionResponseOutputWithContext(ctx context.Context) CostAllocationProportionResponseOutput {
	return o
}

func (o CostAllocationProportionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationProportionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CostAllocationProportionResponseOutput) Percentage() pulumi.Float64Output {
	return o.ApplyT(func(v CostAllocationProportionResponse) float64 { return v.Percentage }).(pulumi.Float64Output)
}

type CostAllocationProportionResponseArrayOutput struct{ *pulumi.OutputState }

func (CostAllocationProportionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CostAllocationProportionResponse)(nil)).Elem()
}

func (o CostAllocationProportionResponseArrayOutput) ToCostAllocationProportionResponseArrayOutput() CostAllocationProportionResponseArrayOutput {
	return o
}

func (o CostAllocationProportionResponseArrayOutput) ToCostAllocationProportionResponseArrayOutputWithContext(ctx context.Context) CostAllocationProportionResponseArrayOutput {
	return o
}

func (o CostAllocationProportionResponseArrayOutput) Index(i pulumi.IntInput) CostAllocationProportionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CostAllocationProportionResponse {
		return vs[0].([]CostAllocationProportionResponse)[vs[1].(int)]
	}).(CostAllocationProportionResponseOutput)
}

type CostAllocationRuleDetails struct {
	SourceResources []SourceCostAllocationResource `pulumi:"sourceResources"`
	TargetResources []TargetCostAllocationResource `pulumi:"targetResources"`
}





type CostAllocationRuleDetailsInput interface {
	pulumi.Input

	ToCostAllocationRuleDetailsOutput() CostAllocationRuleDetailsOutput
	ToCostAllocationRuleDetailsOutputWithContext(context.Context) CostAllocationRuleDetailsOutput
}

type CostAllocationRuleDetailsArgs struct {
	SourceResources SourceCostAllocationResourceArrayInput `pulumi:"sourceResources"`
	TargetResources TargetCostAllocationResourceArrayInput `pulumi:"targetResources"`
}

func (CostAllocationRuleDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRuleDetails)(nil)).Elem()
}

func (i CostAllocationRuleDetailsArgs) ToCostAllocationRuleDetailsOutput() CostAllocationRuleDetailsOutput {
	return i.ToCostAllocationRuleDetailsOutputWithContext(context.Background())
}

func (i CostAllocationRuleDetailsArgs) ToCostAllocationRuleDetailsOutputWithContext(ctx context.Context) CostAllocationRuleDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRuleDetailsOutput)
}

func (i CostAllocationRuleDetailsArgs) ToCostAllocationRuleDetailsPtrOutput() CostAllocationRuleDetailsPtrOutput {
	return i.ToCostAllocationRuleDetailsPtrOutputWithContext(context.Background())
}

func (i CostAllocationRuleDetailsArgs) ToCostAllocationRuleDetailsPtrOutputWithContext(ctx context.Context) CostAllocationRuleDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRuleDetailsOutput).ToCostAllocationRuleDetailsPtrOutputWithContext(ctx)
}









type CostAllocationRuleDetailsPtrInput interface {
	pulumi.Input

	ToCostAllocationRuleDetailsPtrOutput() CostAllocationRuleDetailsPtrOutput
	ToCostAllocationRuleDetailsPtrOutputWithContext(context.Context) CostAllocationRuleDetailsPtrOutput
}

type costAllocationRuleDetailsPtrType CostAllocationRuleDetailsArgs

func CostAllocationRuleDetailsPtr(v *CostAllocationRuleDetailsArgs) CostAllocationRuleDetailsPtrInput {
	return (*costAllocationRuleDetailsPtrType)(v)
}

func (*costAllocationRuleDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRuleDetails)(nil)).Elem()
}

func (i *costAllocationRuleDetailsPtrType) ToCostAllocationRuleDetailsPtrOutput() CostAllocationRuleDetailsPtrOutput {
	return i.ToCostAllocationRuleDetailsPtrOutputWithContext(context.Background())
}

func (i *costAllocationRuleDetailsPtrType) ToCostAllocationRuleDetailsPtrOutputWithContext(ctx context.Context) CostAllocationRuleDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRuleDetailsPtrOutput)
}

type CostAllocationRuleDetailsOutput struct{ *pulumi.OutputState }

func (CostAllocationRuleDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRuleDetails)(nil)).Elem()
}

func (o CostAllocationRuleDetailsOutput) ToCostAllocationRuleDetailsOutput() CostAllocationRuleDetailsOutput {
	return o
}

func (o CostAllocationRuleDetailsOutput) ToCostAllocationRuleDetailsOutputWithContext(ctx context.Context) CostAllocationRuleDetailsOutput {
	return o
}

func (o CostAllocationRuleDetailsOutput) ToCostAllocationRuleDetailsPtrOutput() CostAllocationRuleDetailsPtrOutput {
	return o.ToCostAllocationRuleDetailsPtrOutputWithContext(context.Background())
}

func (o CostAllocationRuleDetailsOutput) ToCostAllocationRuleDetailsPtrOutputWithContext(ctx context.Context) CostAllocationRuleDetailsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CostAllocationRuleDetails) *CostAllocationRuleDetails {
		return &v
	}).(CostAllocationRuleDetailsPtrOutput)
}

func (o CostAllocationRuleDetailsOutput) SourceResources() SourceCostAllocationResourceArrayOutput {
	return o.ApplyT(func(v CostAllocationRuleDetails) []SourceCostAllocationResource { return v.SourceResources }).(SourceCostAllocationResourceArrayOutput)
}

func (o CostAllocationRuleDetailsOutput) TargetResources() TargetCostAllocationResourceArrayOutput {
	return o.ApplyT(func(v CostAllocationRuleDetails) []TargetCostAllocationResource { return v.TargetResources }).(TargetCostAllocationResourceArrayOutput)
}

type CostAllocationRuleDetailsPtrOutput struct{ *pulumi.OutputState }

func (CostAllocationRuleDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRuleDetails)(nil)).Elem()
}

func (o CostAllocationRuleDetailsPtrOutput) ToCostAllocationRuleDetailsPtrOutput() CostAllocationRuleDetailsPtrOutput {
	return o
}

func (o CostAllocationRuleDetailsPtrOutput) ToCostAllocationRuleDetailsPtrOutputWithContext(ctx context.Context) CostAllocationRuleDetailsPtrOutput {
	return o
}

func (o CostAllocationRuleDetailsPtrOutput) Elem() CostAllocationRuleDetailsOutput {
	return o.ApplyT(func(v *CostAllocationRuleDetails) CostAllocationRuleDetails {
		if v != nil {
			return *v
		}
		var ret CostAllocationRuleDetails
		return ret
	}).(CostAllocationRuleDetailsOutput)
}

func (o CostAllocationRuleDetailsPtrOutput) SourceResources() SourceCostAllocationResourceArrayOutput {
	return o.ApplyT(func(v *CostAllocationRuleDetails) []SourceCostAllocationResource {
		if v == nil {
			return nil
		}
		return v.SourceResources
	}).(SourceCostAllocationResourceArrayOutput)
}

func (o CostAllocationRuleDetailsPtrOutput) TargetResources() TargetCostAllocationResourceArrayOutput {
	return o.ApplyT(func(v *CostAllocationRuleDetails) []TargetCostAllocationResource {
		if v == nil {
			return nil
		}
		return v.TargetResources
	}).(TargetCostAllocationResourceArrayOutput)
}

type CostAllocationRuleDetailsResponse struct {
	SourceResources []SourceCostAllocationResourceResponse `pulumi:"sourceResources"`
	TargetResources []TargetCostAllocationResourceResponse `pulumi:"targetResources"`
}

type CostAllocationRuleDetailsResponseOutput struct{ *pulumi.OutputState }

func (CostAllocationRuleDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRuleDetailsResponse)(nil)).Elem()
}

func (o CostAllocationRuleDetailsResponseOutput) ToCostAllocationRuleDetailsResponseOutput() CostAllocationRuleDetailsResponseOutput {
	return o
}

func (o CostAllocationRuleDetailsResponseOutput) ToCostAllocationRuleDetailsResponseOutputWithContext(ctx context.Context) CostAllocationRuleDetailsResponseOutput {
	return o
}

func (o CostAllocationRuleDetailsResponseOutput) SourceResources() SourceCostAllocationResourceResponseArrayOutput {
	return o.ApplyT(func(v CostAllocationRuleDetailsResponse) []SourceCostAllocationResourceResponse {
		return v.SourceResources
	}).(SourceCostAllocationResourceResponseArrayOutput)
}

func (o CostAllocationRuleDetailsResponseOutput) TargetResources() TargetCostAllocationResourceResponseArrayOutput {
	return o.ApplyT(func(v CostAllocationRuleDetailsResponse) []TargetCostAllocationResourceResponse {
		return v.TargetResources
	}).(TargetCostAllocationResourceResponseArrayOutput)
}

type CostAllocationRuleProperties struct {
	Description *string                   `pulumi:"description"`
	Details     CostAllocationRuleDetails `pulumi:"details"`
	Status      string                    `pulumi:"status"`
}





type CostAllocationRulePropertiesInput interface {
	pulumi.Input

	ToCostAllocationRulePropertiesOutput() CostAllocationRulePropertiesOutput
	ToCostAllocationRulePropertiesOutputWithContext(context.Context) CostAllocationRulePropertiesOutput
}

type CostAllocationRulePropertiesArgs struct {
	Description pulumi.StringPtrInput          `pulumi:"description"`
	Details     CostAllocationRuleDetailsInput `pulumi:"details"`
	Status      pulumi.StringInput             `pulumi:"status"`
}

func (CostAllocationRulePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRuleProperties)(nil)).Elem()
}

func (i CostAllocationRulePropertiesArgs) ToCostAllocationRulePropertiesOutput() CostAllocationRulePropertiesOutput {
	return i.ToCostAllocationRulePropertiesOutputWithContext(context.Background())
}

func (i CostAllocationRulePropertiesArgs) ToCostAllocationRulePropertiesOutputWithContext(ctx context.Context) CostAllocationRulePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRulePropertiesOutput)
}

func (i CostAllocationRulePropertiesArgs) ToCostAllocationRulePropertiesPtrOutput() CostAllocationRulePropertiesPtrOutput {
	return i.ToCostAllocationRulePropertiesPtrOutputWithContext(context.Background())
}

func (i CostAllocationRulePropertiesArgs) ToCostAllocationRulePropertiesPtrOutputWithContext(ctx context.Context) CostAllocationRulePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRulePropertiesOutput).ToCostAllocationRulePropertiesPtrOutputWithContext(ctx)
}









type CostAllocationRulePropertiesPtrInput interface {
	pulumi.Input

	ToCostAllocationRulePropertiesPtrOutput() CostAllocationRulePropertiesPtrOutput
	ToCostAllocationRulePropertiesPtrOutputWithContext(context.Context) CostAllocationRulePropertiesPtrOutput
}

type costAllocationRulePropertiesPtrType CostAllocationRulePropertiesArgs

func CostAllocationRulePropertiesPtr(v *CostAllocationRulePropertiesArgs) CostAllocationRulePropertiesPtrInput {
	return (*costAllocationRulePropertiesPtrType)(v)
}

func (*costAllocationRulePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRuleProperties)(nil)).Elem()
}

func (i *costAllocationRulePropertiesPtrType) ToCostAllocationRulePropertiesPtrOutput() CostAllocationRulePropertiesPtrOutput {
	return i.ToCostAllocationRulePropertiesPtrOutputWithContext(context.Background())
}

func (i *costAllocationRulePropertiesPtrType) ToCostAllocationRulePropertiesPtrOutputWithContext(ctx context.Context) CostAllocationRulePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRulePropertiesPtrOutput)
}

type CostAllocationRulePropertiesOutput struct{ *pulumi.OutputState }

func (CostAllocationRulePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRuleProperties)(nil)).Elem()
}

func (o CostAllocationRulePropertiesOutput) ToCostAllocationRulePropertiesOutput() CostAllocationRulePropertiesOutput {
	return o
}

func (o CostAllocationRulePropertiesOutput) ToCostAllocationRulePropertiesOutputWithContext(ctx context.Context) CostAllocationRulePropertiesOutput {
	return o
}

func (o CostAllocationRulePropertiesOutput) ToCostAllocationRulePropertiesPtrOutput() CostAllocationRulePropertiesPtrOutput {
	return o.ToCostAllocationRulePropertiesPtrOutputWithContext(context.Background())
}

func (o CostAllocationRulePropertiesOutput) ToCostAllocationRulePropertiesPtrOutputWithContext(ctx context.Context) CostAllocationRulePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CostAllocationRuleProperties) *CostAllocationRuleProperties {
		return &v
	}).(CostAllocationRulePropertiesPtrOutput)
}

func (o CostAllocationRulePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CostAllocationRuleProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CostAllocationRulePropertiesOutput) Details() CostAllocationRuleDetailsOutput {
	return o.ApplyT(func(v CostAllocationRuleProperties) CostAllocationRuleDetails { return v.Details }).(CostAllocationRuleDetailsOutput)
}

func (o CostAllocationRulePropertiesOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationRuleProperties) string { return v.Status }).(pulumi.StringOutput)
}

type CostAllocationRulePropertiesPtrOutput struct{ *pulumi.OutputState }

func (CostAllocationRulePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRuleProperties)(nil)).Elem()
}

func (o CostAllocationRulePropertiesPtrOutput) ToCostAllocationRulePropertiesPtrOutput() CostAllocationRulePropertiesPtrOutput {
	return o
}

func (o CostAllocationRulePropertiesPtrOutput) ToCostAllocationRulePropertiesPtrOutputWithContext(ctx context.Context) CostAllocationRulePropertiesPtrOutput {
	return o
}

func (o CostAllocationRulePropertiesPtrOutput) Elem() CostAllocationRulePropertiesOutput {
	return o.ApplyT(func(v *CostAllocationRuleProperties) CostAllocationRuleProperties {
		if v != nil {
			return *v
		}
		var ret CostAllocationRuleProperties
		return ret
	}).(CostAllocationRulePropertiesOutput)
}

func (o CostAllocationRulePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CostAllocationRuleProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o CostAllocationRulePropertiesPtrOutput) Details() CostAllocationRuleDetailsPtrOutput {
	return o.ApplyT(func(v *CostAllocationRuleProperties) *CostAllocationRuleDetails {
		if v == nil {
			return nil
		}
		return &v.Details
	}).(CostAllocationRuleDetailsPtrOutput)
}

func (o CostAllocationRulePropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CostAllocationRuleProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type CostAllocationRulePropertiesResponse struct {
	CreatedDate string                            `pulumi:"createdDate"`
	Description *string                           `pulumi:"description"`
	Details     CostAllocationRuleDetailsResponse `pulumi:"details"`
	Status      string                            `pulumi:"status"`
	UpdatedDate string                            `pulumi:"updatedDate"`
}

type CostAllocationRulePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CostAllocationRulePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CostAllocationRulePropertiesResponse)(nil)).Elem()
}

func (o CostAllocationRulePropertiesResponseOutput) ToCostAllocationRulePropertiesResponseOutput() CostAllocationRulePropertiesResponseOutput {
	return o
}

func (o CostAllocationRulePropertiesResponseOutput) ToCostAllocationRulePropertiesResponseOutputWithContext(ctx context.Context) CostAllocationRulePropertiesResponseOutput {
	return o
}

func (o CostAllocationRulePropertiesResponseOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationRulePropertiesResponse) string { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o CostAllocationRulePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CostAllocationRulePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CostAllocationRulePropertiesResponseOutput) Details() CostAllocationRuleDetailsResponseOutput {
	return o.ApplyT(func(v CostAllocationRulePropertiesResponse) CostAllocationRuleDetailsResponse { return v.Details }).(CostAllocationRuleDetailsResponseOutput)
}

func (o CostAllocationRulePropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationRulePropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o CostAllocationRulePropertiesResponseOutput) UpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v CostAllocationRulePropertiesResponse) string { return v.UpdatedDate }).(pulumi.StringOutput)
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

func (o ExportDeliveryDestinationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ExportDeliveryDestinationOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportDeliveryDestination) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ExportDeliveryDestinationResponse struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
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

type FileDestination struct {
	FileFormats []string `pulumi:"fileFormats"`
}





type FileDestinationInput interface {
	pulumi.Input

	ToFileDestinationOutput() FileDestinationOutput
	ToFileDestinationOutputWithContext(context.Context) FileDestinationOutput
}

type FileDestinationArgs struct {
	FileFormats pulumi.StringArrayInput `pulumi:"fileFormats"`
}

func (FileDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileDestination)(nil)).Elem()
}

func (i FileDestinationArgs) ToFileDestinationOutput() FileDestinationOutput {
	return i.ToFileDestinationOutputWithContext(context.Background())
}

func (i FileDestinationArgs) ToFileDestinationOutputWithContext(ctx context.Context) FileDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileDestinationOutput)
}

func (i FileDestinationArgs) ToFileDestinationPtrOutput() FileDestinationPtrOutput {
	return i.ToFileDestinationPtrOutputWithContext(context.Background())
}

func (i FileDestinationArgs) ToFileDestinationPtrOutputWithContext(ctx context.Context) FileDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileDestinationOutput).ToFileDestinationPtrOutputWithContext(ctx)
}









type FileDestinationPtrInput interface {
	pulumi.Input

	ToFileDestinationPtrOutput() FileDestinationPtrOutput
	ToFileDestinationPtrOutputWithContext(context.Context) FileDestinationPtrOutput
}

type fileDestinationPtrType FileDestinationArgs

func FileDestinationPtr(v *FileDestinationArgs) FileDestinationPtrInput {
	return (*fileDestinationPtrType)(v)
}

func (*fileDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileDestination)(nil)).Elem()
}

func (i *fileDestinationPtrType) ToFileDestinationPtrOutput() FileDestinationPtrOutput {
	return i.ToFileDestinationPtrOutputWithContext(context.Background())
}

func (i *fileDestinationPtrType) ToFileDestinationPtrOutputWithContext(ctx context.Context) FileDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileDestinationPtrOutput)
}

type FileDestinationOutput struct{ *pulumi.OutputState }

func (FileDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileDestination)(nil)).Elem()
}

func (o FileDestinationOutput) ToFileDestinationOutput() FileDestinationOutput {
	return o
}

func (o FileDestinationOutput) ToFileDestinationOutputWithContext(ctx context.Context) FileDestinationOutput {
	return o
}

func (o FileDestinationOutput) ToFileDestinationPtrOutput() FileDestinationPtrOutput {
	return o.ToFileDestinationPtrOutputWithContext(context.Background())
}

func (o FileDestinationOutput) ToFileDestinationPtrOutputWithContext(ctx context.Context) FileDestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileDestination) *FileDestination {
		return &v
	}).(FileDestinationPtrOutput)
}

func (o FileDestinationOutput) FileFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FileDestination) []string { return v.FileFormats }).(pulumi.StringArrayOutput)
}

type FileDestinationPtrOutput struct{ *pulumi.OutputState }

func (FileDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileDestination)(nil)).Elem()
}

func (o FileDestinationPtrOutput) ToFileDestinationPtrOutput() FileDestinationPtrOutput {
	return o
}

func (o FileDestinationPtrOutput) ToFileDestinationPtrOutputWithContext(ctx context.Context) FileDestinationPtrOutput {
	return o
}

func (o FileDestinationPtrOutput) Elem() FileDestinationOutput {
	return o.ApplyT(func(v *FileDestination) FileDestination {
		if v != nil {
			return *v
		}
		var ret FileDestination
		return ret
	}).(FileDestinationOutput)
}

func (o FileDestinationPtrOutput) FileFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FileDestination) []string {
		if v == nil {
			return nil
		}
		return v.FileFormats
	}).(pulumi.StringArrayOutput)
}

type FileDestinationResponse struct {
	FileFormats []string `pulumi:"fileFormats"`
}

type FileDestinationResponseOutput struct{ *pulumi.OutputState }

func (FileDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileDestinationResponse)(nil)).Elem()
}

func (o FileDestinationResponseOutput) ToFileDestinationResponseOutput() FileDestinationResponseOutput {
	return o
}

func (o FileDestinationResponseOutput) ToFileDestinationResponseOutputWithContext(ctx context.Context) FileDestinationResponseOutput {
	return o
}

func (o FileDestinationResponseOutput) FileFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FileDestinationResponse) []string { return v.FileFormats }).(pulumi.StringArrayOutput)
}

type FileDestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (FileDestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileDestinationResponse)(nil)).Elem()
}

func (o FileDestinationResponsePtrOutput) ToFileDestinationResponsePtrOutput() FileDestinationResponsePtrOutput {
	return o
}

func (o FileDestinationResponsePtrOutput) ToFileDestinationResponsePtrOutputWithContext(ctx context.Context) FileDestinationResponsePtrOutput {
	return o
}

func (o FileDestinationResponsePtrOutput) Elem() FileDestinationResponseOutput {
	return o.ApplyT(func(v *FileDestinationResponse) FileDestinationResponse {
		if v != nil {
			return *v
		}
		var ret FileDestinationResponse
		return ret
	}).(FileDestinationResponseOutput)
}

func (o FileDestinationResponsePtrOutput) FileFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FileDestinationResponse) []string {
		if v == nil {
			return nil
		}
		return v.FileFormats
	}).(pulumi.StringArrayOutput)
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

type NotificationProperties struct {
	Message *string  `pulumi:"message"`
	Subject string   `pulumi:"subject"`
	To      []string `pulumi:"to"`
}





type NotificationPropertiesInput interface {
	pulumi.Input

	ToNotificationPropertiesOutput() NotificationPropertiesOutput
	ToNotificationPropertiesOutputWithContext(context.Context) NotificationPropertiesOutput
}

type NotificationPropertiesArgs struct {
	Message pulumi.StringPtrInput   `pulumi:"message"`
	Subject pulumi.StringInput      `pulumi:"subject"`
	To      pulumi.StringArrayInput `pulumi:"to"`
}

func (NotificationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationProperties)(nil)).Elem()
}

func (i NotificationPropertiesArgs) ToNotificationPropertiesOutput() NotificationPropertiesOutput {
	return i.ToNotificationPropertiesOutputWithContext(context.Background())
}

func (i NotificationPropertiesArgs) ToNotificationPropertiesOutputWithContext(ctx context.Context) NotificationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPropertiesOutput)
}

type NotificationPropertiesOutput struct{ *pulumi.OutputState }

func (NotificationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationProperties)(nil)).Elem()
}

func (o NotificationPropertiesOutput) ToNotificationPropertiesOutput() NotificationPropertiesOutput {
	return o
}

func (o NotificationPropertiesOutput) ToNotificationPropertiesOutputWithContext(ctx context.Context) NotificationPropertiesOutput {
	return o
}

func (o NotificationPropertiesOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationProperties) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o NotificationPropertiesOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationProperties) string { return v.Subject }).(pulumi.StringOutput)
}

func (o NotificationPropertiesOutput) To() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationProperties) []string { return v.To }).(pulumi.StringArrayOutput)
}

type NotificationPropertiesResponse struct {
	Message *string  `pulumi:"message"`
	Subject string   `pulumi:"subject"`
	To      []string `pulumi:"to"`
}

type NotificationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (NotificationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPropertiesResponse)(nil)).Elem()
}

func (o NotificationPropertiesResponseOutput) ToNotificationPropertiesResponseOutput() NotificationPropertiesResponseOutput {
	return o
}

func (o NotificationPropertiesResponseOutput) ToNotificationPropertiesResponseOutputWithContext(ctx context.Context) NotificationPropertiesResponseOutput {
	return o
}

func (o NotificationPropertiesResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationPropertiesResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o NotificationPropertiesResponseOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationPropertiesResponse) string { return v.Subject }).(pulumi.StringOutput)
}

func (o NotificationPropertiesResponseOutput) To() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationPropertiesResponse) []string { return v.To }).(pulumi.StringArrayOutput)
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

type ReportAggregation struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}





type ReportAggregationInput interface {
	pulumi.Input

	ToReportAggregationOutput() ReportAggregationOutput
	ToReportAggregationOutputWithContext(context.Context) ReportAggregationOutput
}

type ReportAggregationArgs struct {
	Function pulumi.StringInput `pulumi:"function"`
	Name     pulumi.StringInput `pulumi:"name"`
}

func (ReportAggregationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportAggregation)(nil)).Elem()
}

func (i ReportAggregationArgs) ToReportAggregationOutput() ReportAggregationOutput {
	return i.ToReportAggregationOutputWithContext(context.Background())
}

func (i ReportAggregationArgs) ToReportAggregationOutputWithContext(ctx context.Context) ReportAggregationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportAggregationOutput)
}





type ReportAggregationMapInput interface {
	pulumi.Input

	ToReportAggregationMapOutput() ReportAggregationMapOutput
	ToReportAggregationMapOutputWithContext(context.Context) ReportAggregationMapOutput
}

type ReportAggregationMap map[string]ReportAggregationInput

func (ReportAggregationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportAggregation)(nil)).Elem()
}

func (i ReportAggregationMap) ToReportAggregationMapOutput() ReportAggregationMapOutput {
	return i.ToReportAggregationMapOutputWithContext(context.Background())
}

func (i ReportAggregationMap) ToReportAggregationMapOutputWithContext(ctx context.Context) ReportAggregationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportAggregationMapOutput)
}

type ReportAggregationOutput struct{ *pulumi.OutputState }

func (ReportAggregationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportAggregation)(nil)).Elem()
}

func (o ReportAggregationOutput) ToReportAggregationOutput() ReportAggregationOutput {
	return o
}

func (o ReportAggregationOutput) ToReportAggregationOutputWithContext(ctx context.Context) ReportAggregationOutput {
	return o
}

func (o ReportAggregationOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v ReportAggregation) string { return v.Function }).(pulumi.StringOutput)
}

func (o ReportAggregationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportAggregation) string { return v.Name }).(pulumi.StringOutput)
}

type ReportAggregationMapOutput struct{ *pulumi.OutputState }

func (ReportAggregationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportAggregation)(nil)).Elem()
}

func (o ReportAggregationMapOutput) ToReportAggregationMapOutput() ReportAggregationMapOutput {
	return o
}

func (o ReportAggregationMapOutput) ToReportAggregationMapOutputWithContext(ctx context.Context) ReportAggregationMapOutput {
	return o
}

func (o ReportAggregationMapOutput) MapIndex(k pulumi.StringInput) ReportAggregationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReportAggregation {
		return vs[0].(map[string]ReportAggregation)[vs[1].(string)]
	}).(ReportAggregationOutput)
}

type ReportAggregationResponse struct {
	Function string `pulumi:"function"`
	Name     string `pulumi:"name"`
}

type ReportAggregationResponseOutput struct{ *pulumi.OutputState }

func (ReportAggregationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportAggregationResponse)(nil)).Elem()
}

func (o ReportAggregationResponseOutput) ToReportAggregationResponseOutput() ReportAggregationResponseOutput {
	return o
}

func (o ReportAggregationResponseOutput) ToReportAggregationResponseOutputWithContext(ctx context.Context) ReportAggregationResponseOutput {
	return o
}

func (o ReportAggregationResponseOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v ReportAggregationResponse) string { return v.Function }).(pulumi.StringOutput)
}

func (o ReportAggregationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportAggregationResponse) string { return v.Name }).(pulumi.StringOutput)
}

type ReportAggregationResponseMapOutput struct{ *pulumi.OutputState }

func (ReportAggregationResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ReportAggregationResponse)(nil)).Elem()
}

func (o ReportAggregationResponseMapOutput) ToReportAggregationResponseMapOutput() ReportAggregationResponseMapOutput {
	return o
}

func (o ReportAggregationResponseMapOutput) ToReportAggregationResponseMapOutputWithContext(ctx context.Context) ReportAggregationResponseMapOutput {
	return o
}

func (o ReportAggregationResponseMapOutput) MapIndex(k pulumi.StringInput) ReportAggregationResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ReportAggregationResponse {
		return vs[0].(map[string]ReportAggregationResponse)[vs[1].(string)]
	}).(ReportAggregationResponseOutput)
}

type ReportComparisonExpression struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type ReportComparisonExpressionInput interface {
	pulumi.Input

	ToReportComparisonExpressionOutput() ReportComparisonExpressionOutput
	ToReportComparisonExpressionOutputWithContext(context.Context) ReportComparisonExpressionOutput
}

type ReportComparisonExpressionArgs struct {
	Name     pulumi.StringInput      `pulumi:"name"`
	Operator pulumi.StringInput      `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (ReportComparisonExpressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportComparisonExpression)(nil)).Elem()
}

func (i ReportComparisonExpressionArgs) ToReportComparisonExpressionOutput() ReportComparisonExpressionOutput {
	return i.ToReportComparisonExpressionOutputWithContext(context.Background())
}

func (i ReportComparisonExpressionArgs) ToReportComparisonExpressionOutputWithContext(ctx context.Context) ReportComparisonExpressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportComparisonExpressionOutput)
}

func (i ReportComparisonExpressionArgs) ToReportComparisonExpressionPtrOutput() ReportComparisonExpressionPtrOutput {
	return i.ToReportComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i ReportComparisonExpressionArgs) ToReportComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportComparisonExpressionOutput).ToReportComparisonExpressionPtrOutputWithContext(ctx)
}









type ReportComparisonExpressionPtrInput interface {
	pulumi.Input

	ToReportComparisonExpressionPtrOutput() ReportComparisonExpressionPtrOutput
	ToReportComparisonExpressionPtrOutputWithContext(context.Context) ReportComparisonExpressionPtrOutput
}

type reportComparisonExpressionPtrType ReportComparisonExpressionArgs

func ReportComparisonExpressionPtr(v *ReportComparisonExpressionArgs) ReportComparisonExpressionPtrInput {
	return (*reportComparisonExpressionPtrType)(v)
}

func (*reportComparisonExpressionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportComparisonExpression)(nil)).Elem()
}

func (i *reportComparisonExpressionPtrType) ToReportComparisonExpressionPtrOutput() ReportComparisonExpressionPtrOutput {
	return i.ToReportComparisonExpressionPtrOutputWithContext(context.Background())
}

func (i *reportComparisonExpressionPtrType) ToReportComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportComparisonExpressionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportComparisonExpressionPtrOutput)
}

type ReportComparisonExpressionOutput struct{ *pulumi.OutputState }

func (ReportComparisonExpressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportComparisonExpression)(nil)).Elem()
}

func (o ReportComparisonExpressionOutput) ToReportComparisonExpressionOutput() ReportComparisonExpressionOutput {
	return o
}

func (o ReportComparisonExpressionOutput) ToReportComparisonExpressionOutputWithContext(ctx context.Context) ReportComparisonExpressionOutput {
	return o
}

func (o ReportComparisonExpressionOutput) ToReportComparisonExpressionPtrOutput() ReportComparisonExpressionPtrOutput {
	return o.ToReportComparisonExpressionPtrOutputWithContext(context.Background())
}

func (o ReportComparisonExpressionOutput) ToReportComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportComparisonExpressionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportComparisonExpression) *ReportComparisonExpression {
		return &v
	}).(ReportComparisonExpressionPtrOutput)
}

func (o ReportComparisonExpressionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportComparisonExpression) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportComparisonExpressionOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ReportComparisonExpression) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ReportComparisonExpressionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportComparisonExpression) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ReportComparisonExpressionPtrOutput struct{ *pulumi.OutputState }

func (ReportComparisonExpressionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportComparisonExpression)(nil)).Elem()
}

func (o ReportComparisonExpressionPtrOutput) ToReportComparisonExpressionPtrOutput() ReportComparisonExpressionPtrOutput {
	return o
}

func (o ReportComparisonExpressionPtrOutput) ToReportComparisonExpressionPtrOutputWithContext(ctx context.Context) ReportComparisonExpressionPtrOutput {
	return o
}

func (o ReportComparisonExpressionPtrOutput) Elem() ReportComparisonExpressionOutput {
	return o.ApplyT(func(v *ReportComparisonExpression) ReportComparisonExpression {
		if v != nil {
			return *v
		}
		var ret ReportComparisonExpression
		return ret
	}).(ReportComparisonExpressionOutput)
}

func (o ReportComparisonExpressionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ReportComparisonExpressionPtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportComparisonExpression) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ReportComparisonExpressionPtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportComparisonExpression) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type ReportComparisonExpressionResponse struct {
	Name     string   `pulumi:"name"`
	Operator string   `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type ReportComparisonExpressionResponseOutput struct{ *pulumi.OutputState }

func (ReportComparisonExpressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportComparisonExpressionResponse)(nil)).Elem()
}

func (o ReportComparisonExpressionResponseOutput) ToReportComparisonExpressionResponseOutput() ReportComparisonExpressionResponseOutput {
	return o
}

func (o ReportComparisonExpressionResponseOutput) ToReportComparisonExpressionResponseOutputWithContext(ctx context.Context) ReportComparisonExpressionResponseOutput {
	return o
}

func (o ReportComparisonExpressionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportComparisonExpressionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportComparisonExpressionResponseOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v ReportComparisonExpressionResponse) string { return v.Operator }).(pulumi.StringOutput)
}

func (o ReportComparisonExpressionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportComparisonExpressionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ReportComparisonExpressionResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportComparisonExpressionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportComparisonExpressionResponse)(nil)).Elem()
}

func (o ReportComparisonExpressionResponsePtrOutput) ToReportComparisonExpressionResponsePtrOutput() ReportComparisonExpressionResponsePtrOutput {
	return o
}

func (o ReportComparisonExpressionResponsePtrOutput) ToReportComparisonExpressionResponsePtrOutputWithContext(ctx context.Context) ReportComparisonExpressionResponsePtrOutput {
	return o
}

func (o ReportComparisonExpressionResponsePtrOutput) Elem() ReportComparisonExpressionResponseOutput {
	return o.ApplyT(func(v *ReportComparisonExpressionResponse) ReportComparisonExpressionResponse {
		if v != nil {
			return *v
		}
		var ret ReportComparisonExpressionResponse
		return ret
	}).(ReportComparisonExpressionResponseOutput)
}

func (o ReportComparisonExpressionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ReportComparisonExpressionResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportComparisonExpressionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ReportComparisonExpressionResponsePtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportComparisonExpressionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
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

type ReportDataset struct {
	Aggregation   map[string]ReportAggregation `pulumi:"aggregation"`
	Configuration *ReportDatasetConfiguration  `pulumi:"configuration"`
	Filter        *ReportFilter                `pulumi:"filter"`
	Granularity   *string                      `pulumi:"granularity"`
	Grouping      []ReportGrouping             `pulumi:"grouping"`
}





type ReportDatasetInput interface {
	pulumi.Input

	ToReportDatasetOutput() ReportDatasetOutput
	ToReportDatasetOutputWithContext(context.Context) ReportDatasetOutput
}

type ReportDatasetArgs struct {
	Aggregation   ReportAggregationMapInput          `pulumi:"aggregation"`
	Configuration ReportDatasetConfigurationPtrInput `pulumi:"configuration"`
	Filter        ReportFilterPtrInput               `pulumi:"filter"`
	Granularity   pulumi.StringPtrInput              `pulumi:"granularity"`
	Grouping      ReportGroupingArrayInput           `pulumi:"grouping"`
}

func (ReportDatasetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDataset)(nil)).Elem()
}

func (i ReportDatasetArgs) ToReportDatasetOutput() ReportDatasetOutput {
	return i.ToReportDatasetOutputWithContext(context.Background())
}

func (i ReportDatasetArgs) ToReportDatasetOutputWithContext(ctx context.Context) ReportDatasetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetOutput)
}

func (i ReportDatasetArgs) ToReportDatasetPtrOutput() ReportDatasetPtrOutput {
	return i.ToReportDatasetPtrOutputWithContext(context.Background())
}

func (i ReportDatasetArgs) ToReportDatasetPtrOutputWithContext(ctx context.Context) ReportDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetOutput).ToReportDatasetPtrOutputWithContext(ctx)
}









type ReportDatasetPtrInput interface {
	pulumi.Input

	ToReportDatasetPtrOutput() ReportDatasetPtrOutput
	ToReportDatasetPtrOutputWithContext(context.Context) ReportDatasetPtrOutput
}

type reportDatasetPtrType ReportDatasetArgs

func ReportDatasetPtr(v *ReportDatasetArgs) ReportDatasetPtrInput {
	return (*reportDatasetPtrType)(v)
}

func (*reportDatasetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDataset)(nil)).Elem()
}

func (i *reportDatasetPtrType) ToReportDatasetPtrOutput() ReportDatasetPtrOutput {
	return i.ToReportDatasetPtrOutputWithContext(context.Background())
}

func (i *reportDatasetPtrType) ToReportDatasetPtrOutputWithContext(ctx context.Context) ReportDatasetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetPtrOutput)
}

type ReportDatasetOutput struct{ *pulumi.OutputState }

func (ReportDatasetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDataset)(nil)).Elem()
}

func (o ReportDatasetOutput) ToReportDatasetOutput() ReportDatasetOutput {
	return o
}

func (o ReportDatasetOutput) ToReportDatasetOutputWithContext(ctx context.Context) ReportDatasetOutput {
	return o
}

func (o ReportDatasetOutput) ToReportDatasetPtrOutput() ReportDatasetPtrOutput {
	return o.ToReportDatasetPtrOutputWithContext(context.Background())
}

func (o ReportDatasetOutput) ToReportDatasetPtrOutputWithContext(ctx context.Context) ReportDatasetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportDataset) *ReportDataset {
		return &v
	}).(ReportDatasetPtrOutput)
}

func (o ReportDatasetOutput) Aggregation() ReportAggregationMapOutput {
	return o.ApplyT(func(v ReportDataset) map[string]ReportAggregation { return v.Aggregation }).(ReportAggregationMapOutput)
}

func (o ReportDatasetOutput) Configuration() ReportDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v ReportDataset) *ReportDatasetConfiguration { return v.Configuration }).(ReportDatasetConfigurationPtrOutput)
}

func (o ReportDatasetOutput) Filter() ReportFilterPtrOutput {
	return o.ApplyT(func(v ReportDataset) *ReportFilter { return v.Filter }).(ReportFilterPtrOutput)
}

func (o ReportDatasetOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportDataset) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o ReportDatasetOutput) Grouping() ReportGroupingArrayOutput {
	return o.ApplyT(func(v ReportDataset) []ReportGrouping { return v.Grouping }).(ReportGroupingArrayOutput)
}

type ReportDatasetPtrOutput struct{ *pulumi.OutputState }

func (ReportDatasetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDataset)(nil)).Elem()
}

func (o ReportDatasetPtrOutput) ToReportDatasetPtrOutput() ReportDatasetPtrOutput {
	return o
}

func (o ReportDatasetPtrOutput) ToReportDatasetPtrOutputWithContext(ctx context.Context) ReportDatasetPtrOutput {
	return o
}

func (o ReportDatasetPtrOutput) Elem() ReportDatasetOutput {
	return o.ApplyT(func(v *ReportDataset) ReportDataset {
		if v != nil {
			return *v
		}
		var ret ReportDataset
		return ret
	}).(ReportDatasetOutput)
}

func (o ReportDatasetPtrOutput) Aggregation() ReportAggregationMapOutput {
	return o.ApplyT(func(v *ReportDataset) map[string]ReportAggregation {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(ReportAggregationMapOutput)
}

func (o ReportDatasetPtrOutput) Configuration() ReportDatasetConfigurationPtrOutput {
	return o.ApplyT(func(v *ReportDataset) *ReportDatasetConfiguration {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ReportDatasetConfigurationPtrOutput)
}

func (o ReportDatasetPtrOutput) Filter() ReportFilterPtrOutput {
	return o.ApplyT(func(v *ReportDataset) *ReportFilter {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(ReportFilterPtrOutput)
}

func (o ReportDatasetPtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportDataset) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o ReportDatasetPtrOutput) Grouping() ReportGroupingArrayOutput {
	return o.ApplyT(func(v *ReportDataset) []ReportGrouping {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(ReportGroupingArrayOutput)
}

type ReportDatasetConfiguration struct {
	Columns []string `pulumi:"columns"`
}





type ReportDatasetConfigurationInput interface {
	pulumi.Input

	ToReportDatasetConfigurationOutput() ReportDatasetConfigurationOutput
	ToReportDatasetConfigurationOutputWithContext(context.Context) ReportDatasetConfigurationOutput
}

type ReportDatasetConfigurationArgs struct {
	Columns pulumi.StringArrayInput `pulumi:"columns"`
}

func (ReportDatasetConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDatasetConfiguration)(nil)).Elem()
}

func (i ReportDatasetConfigurationArgs) ToReportDatasetConfigurationOutput() ReportDatasetConfigurationOutput {
	return i.ToReportDatasetConfigurationOutputWithContext(context.Background())
}

func (i ReportDatasetConfigurationArgs) ToReportDatasetConfigurationOutputWithContext(ctx context.Context) ReportDatasetConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetConfigurationOutput)
}

func (i ReportDatasetConfigurationArgs) ToReportDatasetConfigurationPtrOutput() ReportDatasetConfigurationPtrOutput {
	return i.ToReportDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i ReportDatasetConfigurationArgs) ToReportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetConfigurationOutput).ToReportDatasetConfigurationPtrOutputWithContext(ctx)
}









type ReportDatasetConfigurationPtrInput interface {
	pulumi.Input

	ToReportDatasetConfigurationPtrOutput() ReportDatasetConfigurationPtrOutput
	ToReportDatasetConfigurationPtrOutputWithContext(context.Context) ReportDatasetConfigurationPtrOutput
}

type reportDatasetConfigurationPtrType ReportDatasetConfigurationArgs

func ReportDatasetConfigurationPtr(v *ReportDatasetConfigurationArgs) ReportDatasetConfigurationPtrInput {
	return (*reportDatasetConfigurationPtrType)(v)
}

func (*reportDatasetConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDatasetConfiguration)(nil)).Elem()
}

func (i *reportDatasetConfigurationPtrType) ToReportDatasetConfigurationPtrOutput() ReportDatasetConfigurationPtrOutput {
	return i.ToReportDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (i *reportDatasetConfigurationPtrType) ToReportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportDatasetConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDatasetConfigurationPtrOutput)
}

type ReportDatasetConfigurationOutput struct{ *pulumi.OutputState }

func (ReportDatasetConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDatasetConfiguration)(nil)).Elem()
}

func (o ReportDatasetConfigurationOutput) ToReportDatasetConfigurationOutput() ReportDatasetConfigurationOutput {
	return o
}

func (o ReportDatasetConfigurationOutput) ToReportDatasetConfigurationOutputWithContext(ctx context.Context) ReportDatasetConfigurationOutput {
	return o
}

func (o ReportDatasetConfigurationOutput) ToReportDatasetConfigurationPtrOutput() ReportDatasetConfigurationPtrOutput {
	return o.ToReportDatasetConfigurationPtrOutputWithContext(context.Background())
}

func (o ReportDatasetConfigurationOutput) ToReportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportDatasetConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportDatasetConfiguration) *ReportDatasetConfiguration {
		return &v
	}).(ReportDatasetConfigurationPtrOutput)
}

func (o ReportDatasetConfigurationOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportDatasetConfiguration) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type ReportDatasetConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ReportDatasetConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDatasetConfiguration)(nil)).Elem()
}

func (o ReportDatasetConfigurationPtrOutput) ToReportDatasetConfigurationPtrOutput() ReportDatasetConfigurationPtrOutput {
	return o
}

func (o ReportDatasetConfigurationPtrOutput) ToReportDatasetConfigurationPtrOutputWithContext(ctx context.Context) ReportDatasetConfigurationPtrOutput {
	return o
}

func (o ReportDatasetConfigurationPtrOutput) Elem() ReportDatasetConfigurationOutput {
	return o.ApplyT(func(v *ReportDatasetConfiguration) ReportDatasetConfiguration {
		if v != nil {
			return *v
		}
		var ret ReportDatasetConfiguration
		return ret
	}).(ReportDatasetConfigurationOutput)
}

func (o ReportDatasetConfigurationPtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportDatasetConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type ReportDatasetConfigurationResponse struct {
	Columns []string `pulumi:"columns"`
}

type ReportDatasetConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ReportDatasetConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDatasetConfigurationResponse)(nil)).Elem()
}

func (o ReportDatasetConfigurationResponseOutput) ToReportDatasetConfigurationResponseOutput() ReportDatasetConfigurationResponseOutput {
	return o
}

func (o ReportDatasetConfigurationResponseOutput) ToReportDatasetConfigurationResponseOutputWithContext(ctx context.Context) ReportDatasetConfigurationResponseOutput {
	return o
}

func (o ReportDatasetConfigurationResponseOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ReportDatasetConfigurationResponse) []string { return v.Columns }).(pulumi.StringArrayOutput)
}

type ReportDatasetConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportDatasetConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDatasetConfigurationResponse)(nil)).Elem()
}

func (o ReportDatasetConfigurationResponsePtrOutput) ToReportDatasetConfigurationResponsePtrOutput() ReportDatasetConfigurationResponsePtrOutput {
	return o
}

func (o ReportDatasetConfigurationResponsePtrOutput) ToReportDatasetConfigurationResponsePtrOutputWithContext(ctx context.Context) ReportDatasetConfigurationResponsePtrOutput {
	return o
}

func (o ReportDatasetConfigurationResponsePtrOutput) Elem() ReportDatasetConfigurationResponseOutput {
	return o.ApplyT(func(v *ReportDatasetConfigurationResponse) ReportDatasetConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ReportDatasetConfigurationResponse
		return ret
	}).(ReportDatasetConfigurationResponseOutput)
}

func (o ReportDatasetConfigurationResponsePtrOutput) Columns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReportDatasetConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(pulumi.StringArrayOutput)
}

type ReportDatasetResponse struct {
	Aggregation   map[string]ReportAggregationResponse `pulumi:"aggregation"`
	Configuration *ReportDatasetConfigurationResponse  `pulumi:"configuration"`
	Filter        *ReportFilterResponse                `pulumi:"filter"`
	Granularity   *string                              `pulumi:"granularity"`
	Grouping      []ReportGroupingResponse             `pulumi:"grouping"`
}

type ReportDatasetResponseOutput struct{ *pulumi.OutputState }

func (ReportDatasetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDatasetResponse)(nil)).Elem()
}

func (o ReportDatasetResponseOutput) ToReportDatasetResponseOutput() ReportDatasetResponseOutput {
	return o
}

func (o ReportDatasetResponseOutput) ToReportDatasetResponseOutputWithContext(ctx context.Context) ReportDatasetResponseOutput {
	return o
}

func (o ReportDatasetResponseOutput) Aggregation() ReportAggregationResponseMapOutput {
	return o.ApplyT(func(v ReportDatasetResponse) map[string]ReportAggregationResponse { return v.Aggregation }).(ReportAggregationResponseMapOutput)
}

func (o ReportDatasetResponseOutput) Configuration() ReportDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v ReportDatasetResponse) *ReportDatasetConfigurationResponse { return v.Configuration }).(ReportDatasetConfigurationResponsePtrOutput)
}

func (o ReportDatasetResponseOutput) Filter() ReportFilterResponsePtrOutput {
	return o.ApplyT(func(v ReportDatasetResponse) *ReportFilterResponse { return v.Filter }).(ReportFilterResponsePtrOutput)
}

func (o ReportDatasetResponseOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportDatasetResponse) *string { return v.Granularity }).(pulumi.StringPtrOutput)
}

func (o ReportDatasetResponseOutput) Grouping() ReportGroupingResponseArrayOutput {
	return o.ApplyT(func(v ReportDatasetResponse) []ReportGroupingResponse { return v.Grouping }).(ReportGroupingResponseArrayOutput)
}

type ReportDatasetResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportDatasetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportDatasetResponse)(nil)).Elem()
}

func (o ReportDatasetResponsePtrOutput) ToReportDatasetResponsePtrOutput() ReportDatasetResponsePtrOutput {
	return o
}

func (o ReportDatasetResponsePtrOutput) ToReportDatasetResponsePtrOutputWithContext(ctx context.Context) ReportDatasetResponsePtrOutput {
	return o
}

func (o ReportDatasetResponsePtrOutput) Elem() ReportDatasetResponseOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) ReportDatasetResponse {
		if v != nil {
			return *v
		}
		var ret ReportDatasetResponse
		return ret
	}).(ReportDatasetResponseOutput)
}

func (o ReportDatasetResponsePtrOutput) Aggregation() ReportAggregationResponseMapOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) map[string]ReportAggregationResponse {
		if v == nil {
			return nil
		}
		return v.Aggregation
	}).(ReportAggregationResponseMapOutput)
}

func (o ReportDatasetResponsePtrOutput) Configuration() ReportDatasetConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) *ReportDatasetConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ReportDatasetConfigurationResponsePtrOutput)
}

func (o ReportDatasetResponsePtrOutput) Filter() ReportFilterResponsePtrOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) *ReportFilterResponse {
		if v == nil {
			return nil
		}
		return v.Filter
	}).(ReportFilterResponsePtrOutput)
}

func (o ReportDatasetResponsePtrOutput) Granularity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Granularity
	}).(pulumi.StringPtrOutput)
}

func (o ReportDatasetResponsePtrOutput) Grouping() ReportGroupingResponseArrayOutput {
	return o.ApplyT(func(v *ReportDatasetResponse) []ReportGroupingResponse {
		if v == nil {
			return nil
		}
		return v.Grouping
	}).(ReportGroupingResponseArrayOutput)
}

type ReportDefinition struct {
	Dataset    *ReportDataset    `pulumi:"dataset"`
	TimePeriod *ReportTimePeriod `pulumi:"timePeriod"`
	Timeframe  string            `pulumi:"timeframe"`
	Type       string            `pulumi:"type"`
}





type ReportDefinitionInput interface {
	pulumi.Input

	ToReportDefinitionOutput() ReportDefinitionOutput
	ToReportDefinitionOutputWithContext(context.Context) ReportDefinitionOutput
}

type ReportDefinitionArgs struct {
	Dataset    ReportDatasetPtrInput    `pulumi:"dataset"`
	TimePeriod ReportTimePeriodPtrInput `pulumi:"timePeriod"`
	Timeframe  pulumi.StringInput       `pulumi:"timeframe"`
	Type       pulumi.StringInput       `pulumi:"type"`
}

func (ReportDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDefinition)(nil)).Elem()
}

func (i ReportDefinitionArgs) ToReportDefinitionOutput() ReportDefinitionOutput {
	return i.ToReportDefinitionOutputWithContext(context.Background())
}

func (i ReportDefinitionArgs) ToReportDefinitionOutputWithContext(ctx context.Context) ReportDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDefinitionOutput)
}

type ReportDefinitionOutput struct{ *pulumi.OutputState }

func (ReportDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDefinition)(nil)).Elem()
}

func (o ReportDefinitionOutput) ToReportDefinitionOutput() ReportDefinitionOutput {
	return o
}

func (o ReportDefinitionOutput) ToReportDefinitionOutputWithContext(ctx context.Context) ReportDefinitionOutput {
	return o
}

func (o ReportDefinitionOutput) Dataset() ReportDatasetPtrOutput {
	return o.ApplyT(func(v ReportDefinition) *ReportDataset { return v.Dataset }).(ReportDatasetPtrOutput)
}

func (o ReportDefinitionOutput) TimePeriod() ReportTimePeriodPtrOutput {
	return o.ApplyT(func(v ReportDefinition) *ReportTimePeriod { return v.TimePeriod }).(ReportTimePeriodPtrOutput)
}

func (o ReportDefinitionOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDefinition) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ReportDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type ReportDefinitionResponse struct {
	Dataset    *ReportDatasetResponse    `pulumi:"dataset"`
	TimePeriod *ReportTimePeriodResponse `pulumi:"timePeriod"`
	Timeframe  string                    `pulumi:"timeframe"`
	Type       string                    `pulumi:"type"`
}

type ReportDefinitionResponseOutput struct{ *pulumi.OutputState }

func (ReportDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDefinitionResponse)(nil)).Elem()
}

func (o ReportDefinitionResponseOutput) ToReportDefinitionResponseOutput() ReportDefinitionResponseOutput {
	return o
}

func (o ReportDefinitionResponseOutput) ToReportDefinitionResponseOutputWithContext(ctx context.Context) ReportDefinitionResponseOutput {
	return o
}

func (o ReportDefinitionResponseOutput) Dataset() ReportDatasetResponsePtrOutput {
	return o.ApplyT(func(v ReportDefinitionResponse) *ReportDatasetResponse { return v.Dataset }).(ReportDatasetResponsePtrOutput)
}

func (o ReportDefinitionResponseOutput) TimePeriod() ReportTimePeriodResponsePtrOutput {
	return o.ApplyT(func(v ReportDefinitionResponse) *ReportTimePeriodResponse { return v.TimePeriod }).(ReportTimePeriodResponsePtrOutput)
}

func (o ReportDefinitionResponseOutput) Timeframe() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDefinitionResponse) string { return v.Timeframe }).(pulumi.StringOutput)
}

func (o ReportDefinitionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDefinitionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ReportDeliveryDestination struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
}





type ReportDeliveryDestinationInput interface {
	pulumi.Input

	ToReportDeliveryDestinationOutput() ReportDeliveryDestinationOutput
	ToReportDeliveryDestinationOutputWithContext(context.Context) ReportDeliveryDestinationOutput
}

type ReportDeliveryDestinationArgs struct {
	Container      pulumi.StringInput    `pulumi:"container"`
	ResourceId     pulumi.StringInput    `pulumi:"resourceId"`
	RootFolderPath pulumi.StringPtrInput `pulumi:"rootFolderPath"`
}

func (ReportDeliveryDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryDestination)(nil)).Elem()
}

func (i ReportDeliveryDestinationArgs) ToReportDeliveryDestinationOutput() ReportDeliveryDestinationOutput {
	return i.ToReportDeliveryDestinationOutputWithContext(context.Background())
}

func (i ReportDeliveryDestinationArgs) ToReportDeliveryDestinationOutputWithContext(ctx context.Context) ReportDeliveryDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDeliveryDestinationOutput)
}

type ReportDeliveryDestinationOutput struct{ *pulumi.OutputState }

func (ReportDeliveryDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryDestination)(nil)).Elem()
}

func (o ReportDeliveryDestinationOutput) ToReportDeliveryDestinationOutput() ReportDeliveryDestinationOutput {
	return o
}

func (o ReportDeliveryDestinationOutput) ToReportDeliveryDestinationOutputWithContext(ctx context.Context) ReportDeliveryDestinationOutput {
	return o
}

func (o ReportDeliveryDestinationOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDeliveryDestination) string { return v.Container }).(pulumi.StringOutput)
}

func (o ReportDeliveryDestinationOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDeliveryDestination) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ReportDeliveryDestinationOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportDeliveryDestination) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ReportDeliveryDestinationResponse struct {
	Container      string  `pulumi:"container"`
	ResourceId     string  `pulumi:"resourceId"`
	RootFolderPath *string `pulumi:"rootFolderPath"`
}

type ReportDeliveryDestinationResponseOutput struct{ *pulumi.OutputState }

func (ReportDeliveryDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryDestinationResponse)(nil)).Elem()
}

func (o ReportDeliveryDestinationResponseOutput) ToReportDeliveryDestinationResponseOutput() ReportDeliveryDestinationResponseOutput {
	return o
}

func (o ReportDeliveryDestinationResponseOutput) ToReportDeliveryDestinationResponseOutputWithContext(ctx context.Context) ReportDeliveryDestinationResponseOutput {
	return o
}

func (o ReportDeliveryDestinationResponseOutput) Container() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDeliveryDestinationResponse) string { return v.Container }).(pulumi.StringOutput)
}

func (o ReportDeliveryDestinationResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v ReportDeliveryDestinationResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

func (o ReportDeliveryDestinationResponseOutput) RootFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportDeliveryDestinationResponse) *string { return v.RootFolderPath }).(pulumi.StringPtrOutput)
}

type ReportDeliveryInfo struct {
	Destination ReportDeliveryDestination `pulumi:"destination"`
}





type ReportDeliveryInfoInput interface {
	pulumi.Input

	ToReportDeliveryInfoOutput() ReportDeliveryInfoOutput
	ToReportDeliveryInfoOutputWithContext(context.Context) ReportDeliveryInfoOutput
}

type ReportDeliveryInfoArgs struct {
	Destination ReportDeliveryDestinationInput `pulumi:"destination"`
}

func (ReportDeliveryInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryInfo)(nil)).Elem()
}

func (i ReportDeliveryInfoArgs) ToReportDeliveryInfoOutput() ReportDeliveryInfoOutput {
	return i.ToReportDeliveryInfoOutputWithContext(context.Background())
}

func (i ReportDeliveryInfoArgs) ToReportDeliveryInfoOutputWithContext(ctx context.Context) ReportDeliveryInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportDeliveryInfoOutput)
}

type ReportDeliveryInfoOutput struct{ *pulumi.OutputState }

func (ReportDeliveryInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryInfo)(nil)).Elem()
}

func (o ReportDeliveryInfoOutput) ToReportDeliveryInfoOutput() ReportDeliveryInfoOutput {
	return o
}

func (o ReportDeliveryInfoOutput) ToReportDeliveryInfoOutputWithContext(ctx context.Context) ReportDeliveryInfoOutput {
	return o
}

func (o ReportDeliveryInfoOutput) Destination() ReportDeliveryDestinationOutput {
	return o.ApplyT(func(v ReportDeliveryInfo) ReportDeliveryDestination { return v.Destination }).(ReportDeliveryDestinationOutput)
}

type ReportDeliveryInfoResponse struct {
	Destination ReportDeliveryDestinationResponse `pulumi:"destination"`
}

type ReportDeliveryInfoResponseOutput struct{ *pulumi.OutputState }

func (ReportDeliveryInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportDeliveryInfoResponse)(nil)).Elem()
}

func (o ReportDeliveryInfoResponseOutput) ToReportDeliveryInfoResponseOutput() ReportDeliveryInfoResponseOutput {
	return o
}

func (o ReportDeliveryInfoResponseOutput) ToReportDeliveryInfoResponseOutputWithContext(ctx context.Context) ReportDeliveryInfoResponseOutput {
	return o
}

func (o ReportDeliveryInfoResponseOutput) Destination() ReportDeliveryDestinationResponseOutput {
	return o.ApplyT(func(v ReportDeliveryInfoResponse) ReportDeliveryDestinationResponse { return v.Destination }).(ReportDeliveryDestinationResponseOutput)
}

type ReportFilter struct {
	And       []ReportFilter              `pulumi:"and"`
	Dimension *ReportComparisonExpression `pulumi:"dimension"`
	Not       *ReportFilter               `pulumi:"not"`
	Or        []ReportFilter              `pulumi:"or"`
	Tag       *ReportComparisonExpression `pulumi:"tag"`
}





type ReportFilterInput interface {
	pulumi.Input

	ToReportFilterOutput() ReportFilterOutput
	ToReportFilterOutputWithContext(context.Context) ReportFilterOutput
}

type ReportFilterArgs struct {
	And       ReportFilterArrayInput             `pulumi:"and"`
	Dimension ReportComparisonExpressionPtrInput `pulumi:"dimension"`
	Not       ReportFilterPtrInput               `pulumi:"not"`
	Or        ReportFilterArrayInput             `pulumi:"or"`
	Tag       ReportComparisonExpressionPtrInput `pulumi:"tag"`
}

func (ReportFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportFilter)(nil)).Elem()
}

func (i ReportFilterArgs) ToReportFilterOutput() ReportFilterOutput {
	return i.ToReportFilterOutputWithContext(context.Background())
}

func (i ReportFilterArgs) ToReportFilterOutputWithContext(ctx context.Context) ReportFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportFilterOutput)
}

func (i ReportFilterArgs) ToReportFilterPtrOutput() ReportFilterPtrOutput {
	return i.ToReportFilterPtrOutputWithContext(context.Background())
}

func (i ReportFilterArgs) ToReportFilterPtrOutputWithContext(ctx context.Context) ReportFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportFilterOutput).ToReportFilterPtrOutputWithContext(ctx)
}









type ReportFilterPtrInput interface {
	pulumi.Input

	ToReportFilterPtrOutput() ReportFilterPtrOutput
	ToReportFilterPtrOutputWithContext(context.Context) ReportFilterPtrOutput
}

type reportFilterPtrType ReportFilterArgs

func ReportFilterPtr(v *ReportFilterArgs) ReportFilterPtrInput {
	return (*reportFilterPtrType)(v)
}

func (*reportFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportFilter)(nil)).Elem()
}

func (i *reportFilterPtrType) ToReportFilterPtrOutput() ReportFilterPtrOutput {
	return i.ToReportFilterPtrOutputWithContext(context.Background())
}

func (i *reportFilterPtrType) ToReportFilterPtrOutputWithContext(ctx context.Context) ReportFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportFilterPtrOutput)
}





type ReportFilterArrayInput interface {
	pulumi.Input

	ToReportFilterArrayOutput() ReportFilterArrayOutput
	ToReportFilterArrayOutputWithContext(context.Context) ReportFilterArrayOutput
}

type ReportFilterArray []ReportFilterInput

func (ReportFilterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportFilter)(nil)).Elem()
}

func (i ReportFilterArray) ToReportFilterArrayOutput() ReportFilterArrayOutput {
	return i.ToReportFilterArrayOutputWithContext(context.Background())
}

func (i ReportFilterArray) ToReportFilterArrayOutputWithContext(ctx context.Context) ReportFilterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportFilterArrayOutput)
}

type ReportFilterOutput struct{ *pulumi.OutputState }

func (ReportFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportFilter)(nil)).Elem()
}

func (o ReportFilterOutput) ToReportFilterOutput() ReportFilterOutput {
	return o
}

func (o ReportFilterOutput) ToReportFilterOutputWithContext(ctx context.Context) ReportFilterOutput {
	return o
}

func (o ReportFilterOutput) ToReportFilterPtrOutput() ReportFilterPtrOutput {
	return o.ToReportFilterPtrOutputWithContext(context.Background())
}

func (o ReportFilterOutput) ToReportFilterPtrOutputWithContext(ctx context.Context) ReportFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportFilter) *ReportFilter {
		return &v
	}).(ReportFilterPtrOutput)
}

func (o ReportFilterOutput) And() ReportFilterArrayOutput {
	return o.ApplyT(func(v ReportFilter) []ReportFilter { return v.And }).(ReportFilterArrayOutput)
}

func (o ReportFilterOutput) Dimension() ReportComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportFilter) *ReportComparisonExpression { return v.Dimension }).(ReportComparisonExpressionPtrOutput)
}

func (o ReportFilterOutput) Not() ReportFilterPtrOutput {
	return o.ApplyT(func(v ReportFilter) *ReportFilter { return v.Not }).(ReportFilterPtrOutput)
}

func (o ReportFilterOutput) Or() ReportFilterArrayOutput {
	return o.ApplyT(func(v ReportFilter) []ReportFilter { return v.Or }).(ReportFilterArrayOutput)
}

func (o ReportFilterOutput) Tag() ReportComparisonExpressionPtrOutput {
	return o.ApplyT(func(v ReportFilter) *ReportComparisonExpression { return v.Tag }).(ReportComparisonExpressionPtrOutput)
}

type ReportFilterPtrOutput struct{ *pulumi.OutputState }

func (ReportFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportFilter)(nil)).Elem()
}

func (o ReportFilterPtrOutput) ToReportFilterPtrOutput() ReportFilterPtrOutput {
	return o
}

func (o ReportFilterPtrOutput) ToReportFilterPtrOutputWithContext(ctx context.Context) ReportFilterPtrOutput {
	return o
}

func (o ReportFilterPtrOutput) Elem() ReportFilterOutput {
	return o.ApplyT(func(v *ReportFilter) ReportFilter {
		if v != nil {
			return *v
		}
		var ret ReportFilter
		return ret
	}).(ReportFilterOutput)
}

func (o ReportFilterPtrOutput) And() ReportFilterArrayOutput {
	return o.ApplyT(func(v *ReportFilter) []ReportFilter {
		if v == nil {
			return nil
		}
		return v.And
	}).(ReportFilterArrayOutput)
}

func (o ReportFilterPtrOutput) Dimension() ReportComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportFilter) *ReportComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Dimension
	}).(ReportComparisonExpressionPtrOutput)
}

func (o ReportFilterPtrOutput) Not() ReportFilterPtrOutput {
	return o.ApplyT(func(v *ReportFilter) *ReportFilter {
		if v == nil {
			return nil
		}
		return v.Not
	}).(ReportFilterPtrOutput)
}

func (o ReportFilterPtrOutput) Or() ReportFilterArrayOutput {
	return o.ApplyT(func(v *ReportFilter) []ReportFilter {
		if v == nil {
			return nil
		}
		return v.Or
	}).(ReportFilterArrayOutput)
}

func (o ReportFilterPtrOutput) Tag() ReportComparisonExpressionPtrOutput {
	return o.ApplyT(func(v *ReportFilter) *ReportComparisonExpression {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(ReportComparisonExpressionPtrOutput)
}

type ReportFilterArrayOutput struct{ *pulumi.OutputState }

func (ReportFilterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportFilter)(nil)).Elem()
}

func (o ReportFilterArrayOutput) ToReportFilterArrayOutput() ReportFilterArrayOutput {
	return o
}

func (o ReportFilterArrayOutput) ToReportFilterArrayOutputWithContext(ctx context.Context) ReportFilterArrayOutput {
	return o
}

func (o ReportFilterArrayOutput) Index(i pulumi.IntInput) ReportFilterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportFilter {
		return vs[0].([]ReportFilter)[vs[1].(int)]
	}).(ReportFilterOutput)
}

type ReportFilterResponse struct {
	And       []ReportFilterResponse              `pulumi:"and"`
	Dimension *ReportComparisonExpressionResponse `pulumi:"dimension"`
	Not       *ReportFilterResponse               `pulumi:"not"`
	Or        []ReportFilterResponse              `pulumi:"or"`
	Tag       *ReportComparisonExpressionResponse `pulumi:"tag"`
}

type ReportFilterResponseOutput struct{ *pulumi.OutputState }

func (ReportFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportFilterResponse)(nil)).Elem()
}

func (o ReportFilterResponseOutput) ToReportFilterResponseOutput() ReportFilterResponseOutput {
	return o
}

func (o ReportFilterResponseOutput) ToReportFilterResponseOutputWithContext(ctx context.Context) ReportFilterResponseOutput {
	return o
}

func (o ReportFilterResponseOutput) And() ReportFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportFilterResponse) []ReportFilterResponse { return v.And }).(ReportFilterResponseArrayOutput)
}

func (o ReportFilterResponseOutput) Dimension() ReportComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportFilterResponse) *ReportComparisonExpressionResponse { return v.Dimension }).(ReportComparisonExpressionResponsePtrOutput)
}

func (o ReportFilterResponseOutput) Not() ReportFilterResponsePtrOutput {
	return o.ApplyT(func(v ReportFilterResponse) *ReportFilterResponse { return v.Not }).(ReportFilterResponsePtrOutput)
}

func (o ReportFilterResponseOutput) Or() ReportFilterResponseArrayOutput {
	return o.ApplyT(func(v ReportFilterResponse) []ReportFilterResponse { return v.Or }).(ReportFilterResponseArrayOutput)
}

func (o ReportFilterResponseOutput) Tag() ReportComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v ReportFilterResponse) *ReportComparisonExpressionResponse { return v.Tag }).(ReportComparisonExpressionResponsePtrOutput)
}

type ReportFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportFilterResponse)(nil)).Elem()
}

func (o ReportFilterResponsePtrOutput) ToReportFilterResponsePtrOutput() ReportFilterResponsePtrOutput {
	return o
}

func (o ReportFilterResponsePtrOutput) ToReportFilterResponsePtrOutputWithContext(ctx context.Context) ReportFilterResponsePtrOutput {
	return o
}

func (o ReportFilterResponsePtrOutput) Elem() ReportFilterResponseOutput {
	return o.ApplyT(func(v *ReportFilterResponse) ReportFilterResponse {
		if v != nil {
			return *v
		}
		var ret ReportFilterResponse
		return ret
	}).(ReportFilterResponseOutput)
}

func (o ReportFilterResponsePtrOutput) And() ReportFilterResponseArrayOutput {
	return o.ApplyT(func(v *ReportFilterResponse) []ReportFilterResponse {
		if v == nil {
			return nil
		}
		return v.And
	}).(ReportFilterResponseArrayOutput)
}

func (o ReportFilterResponsePtrOutput) Dimension() ReportComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportFilterResponse) *ReportComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Dimension
	}).(ReportComparisonExpressionResponsePtrOutput)
}

func (o ReportFilterResponsePtrOutput) Not() ReportFilterResponsePtrOutput {
	return o.ApplyT(func(v *ReportFilterResponse) *ReportFilterResponse {
		if v == nil {
			return nil
		}
		return v.Not
	}).(ReportFilterResponsePtrOutput)
}

func (o ReportFilterResponsePtrOutput) Or() ReportFilterResponseArrayOutput {
	return o.ApplyT(func(v *ReportFilterResponse) []ReportFilterResponse {
		if v == nil {
			return nil
		}
		return v.Or
	}).(ReportFilterResponseArrayOutput)
}

func (o ReportFilterResponsePtrOutput) Tag() ReportComparisonExpressionResponsePtrOutput {
	return o.ApplyT(func(v *ReportFilterResponse) *ReportComparisonExpressionResponse {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(ReportComparisonExpressionResponsePtrOutput)
}

type ReportFilterResponseArrayOutput struct{ *pulumi.OutputState }

func (ReportFilterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportFilterResponse)(nil)).Elem()
}

func (o ReportFilterResponseArrayOutput) ToReportFilterResponseArrayOutput() ReportFilterResponseArrayOutput {
	return o
}

func (o ReportFilterResponseArrayOutput) ToReportFilterResponseArrayOutputWithContext(ctx context.Context) ReportFilterResponseArrayOutput {
	return o
}

func (o ReportFilterResponseArrayOutput) Index(i pulumi.IntInput) ReportFilterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportFilterResponse {
		return vs[0].([]ReportFilterResponse)[vs[1].(int)]
	}).(ReportFilterResponseOutput)
}

type ReportGrouping struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}





type ReportGroupingInput interface {
	pulumi.Input

	ToReportGroupingOutput() ReportGroupingOutput
	ToReportGroupingOutputWithContext(context.Context) ReportGroupingOutput
}

type ReportGroupingArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Type pulumi.StringInput `pulumi:"type"`
}

func (ReportGroupingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportGrouping)(nil)).Elem()
}

func (i ReportGroupingArgs) ToReportGroupingOutput() ReportGroupingOutput {
	return i.ToReportGroupingOutputWithContext(context.Background())
}

func (i ReportGroupingArgs) ToReportGroupingOutputWithContext(ctx context.Context) ReportGroupingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportGroupingOutput)
}





type ReportGroupingArrayInput interface {
	pulumi.Input

	ToReportGroupingArrayOutput() ReportGroupingArrayOutput
	ToReportGroupingArrayOutputWithContext(context.Context) ReportGroupingArrayOutput
}

type ReportGroupingArray []ReportGroupingInput

func (ReportGroupingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportGrouping)(nil)).Elem()
}

func (i ReportGroupingArray) ToReportGroupingArrayOutput() ReportGroupingArrayOutput {
	return i.ToReportGroupingArrayOutputWithContext(context.Background())
}

func (i ReportGroupingArray) ToReportGroupingArrayOutputWithContext(ctx context.Context) ReportGroupingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportGroupingArrayOutput)
}

type ReportGroupingOutput struct{ *pulumi.OutputState }

func (ReportGroupingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportGrouping)(nil)).Elem()
}

func (o ReportGroupingOutput) ToReportGroupingOutput() ReportGroupingOutput {
	return o
}

func (o ReportGroupingOutput) ToReportGroupingOutputWithContext(ctx context.Context) ReportGroupingOutput {
	return o
}

func (o ReportGroupingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportGrouping) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportGroupingOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportGrouping) string { return v.Type }).(pulumi.StringOutput)
}

type ReportGroupingArrayOutput struct{ *pulumi.OutputState }

func (ReportGroupingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportGrouping)(nil)).Elem()
}

func (o ReportGroupingArrayOutput) ToReportGroupingArrayOutput() ReportGroupingArrayOutput {
	return o
}

func (o ReportGroupingArrayOutput) ToReportGroupingArrayOutputWithContext(ctx context.Context) ReportGroupingArrayOutput {
	return o
}

func (o ReportGroupingArrayOutput) Index(i pulumi.IntInput) ReportGroupingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportGrouping {
		return vs[0].([]ReportGrouping)[vs[1].(int)]
	}).(ReportGroupingOutput)
}

type ReportGroupingResponse struct {
	Name string `pulumi:"name"`
	Type string `pulumi:"type"`
}

type ReportGroupingResponseOutput struct{ *pulumi.OutputState }

func (ReportGroupingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportGroupingResponse)(nil)).Elem()
}

func (o ReportGroupingResponseOutput) ToReportGroupingResponseOutput() ReportGroupingResponseOutput {
	return o
}

func (o ReportGroupingResponseOutput) ToReportGroupingResponseOutputWithContext(ctx context.Context) ReportGroupingResponseOutput {
	return o
}

func (o ReportGroupingResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ReportGroupingResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ReportGroupingResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ReportGroupingResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ReportGroupingResponseArrayOutput struct{ *pulumi.OutputState }

func (ReportGroupingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReportGroupingResponse)(nil)).Elem()
}

func (o ReportGroupingResponseArrayOutput) ToReportGroupingResponseArrayOutput() ReportGroupingResponseArrayOutput {
	return o
}

func (o ReportGroupingResponseArrayOutput) ToReportGroupingResponseArrayOutputWithContext(ctx context.Context) ReportGroupingResponseArrayOutput {
	return o
}

func (o ReportGroupingResponseArrayOutput) Index(i pulumi.IntInput) ReportGroupingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReportGroupingResponse {
		return vs[0].([]ReportGroupingResponse)[vs[1].(int)]
	}).(ReportGroupingResponseOutput)
}

type ReportRecurrencePeriod struct {
	From string  `pulumi:"from"`
	To   *string `pulumi:"to"`
}





type ReportRecurrencePeriodInput interface {
	pulumi.Input

	ToReportRecurrencePeriodOutput() ReportRecurrencePeriodOutput
	ToReportRecurrencePeriodOutputWithContext(context.Context) ReportRecurrencePeriodOutput
}

type ReportRecurrencePeriodArgs struct {
	From pulumi.StringInput    `pulumi:"from"`
	To   pulumi.StringPtrInput `pulumi:"to"`
}

func (ReportRecurrencePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportRecurrencePeriod)(nil)).Elem()
}

func (i ReportRecurrencePeriodArgs) ToReportRecurrencePeriodOutput() ReportRecurrencePeriodOutput {
	return i.ToReportRecurrencePeriodOutputWithContext(context.Background())
}

func (i ReportRecurrencePeriodArgs) ToReportRecurrencePeriodOutputWithContext(ctx context.Context) ReportRecurrencePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportRecurrencePeriodOutput)
}

func (i ReportRecurrencePeriodArgs) ToReportRecurrencePeriodPtrOutput() ReportRecurrencePeriodPtrOutput {
	return i.ToReportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (i ReportRecurrencePeriodArgs) ToReportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportRecurrencePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportRecurrencePeriodOutput).ToReportRecurrencePeriodPtrOutputWithContext(ctx)
}









type ReportRecurrencePeriodPtrInput interface {
	pulumi.Input

	ToReportRecurrencePeriodPtrOutput() ReportRecurrencePeriodPtrOutput
	ToReportRecurrencePeriodPtrOutputWithContext(context.Context) ReportRecurrencePeriodPtrOutput
}

type reportRecurrencePeriodPtrType ReportRecurrencePeriodArgs

func ReportRecurrencePeriodPtr(v *ReportRecurrencePeriodArgs) ReportRecurrencePeriodPtrInput {
	return (*reportRecurrencePeriodPtrType)(v)
}

func (*reportRecurrencePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportRecurrencePeriod)(nil)).Elem()
}

func (i *reportRecurrencePeriodPtrType) ToReportRecurrencePeriodPtrOutput() ReportRecurrencePeriodPtrOutput {
	return i.ToReportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (i *reportRecurrencePeriodPtrType) ToReportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportRecurrencePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportRecurrencePeriodPtrOutput)
}

type ReportRecurrencePeriodOutput struct{ *pulumi.OutputState }

func (ReportRecurrencePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportRecurrencePeriod)(nil)).Elem()
}

func (o ReportRecurrencePeriodOutput) ToReportRecurrencePeriodOutput() ReportRecurrencePeriodOutput {
	return o
}

func (o ReportRecurrencePeriodOutput) ToReportRecurrencePeriodOutputWithContext(ctx context.Context) ReportRecurrencePeriodOutput {
	return o
}

func (o ReportRecurrencePeriodOutput) ToReportRecurrencePeriodPtrOutput() ReportRecurrencePeriodPtrOutput {
	return o.ToReportRecurrencePeriodPtrOutputWithContext(context.Background())
}

func (o ReportRecurrencePeriodOutput) ToReportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportRecurrencePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportRecurrencePeriod) *ReportRecurrencePeriod {
		return &v
	}).(ReportRecurrencePeriodPtrOutput)
}

func (o ReportRecurrencePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportRecurrencePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportRecurrencePeriodOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportRecurrencePeriod) *string { return v.To }).(pulumi.StringPtrOutput)
}

type ReportRecurrencePeriodPtrOutput struct{ *pulumi.OutputState }

func (ReportRecurrencePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportRecurrencePeriod)(nil)).Elem()
}

func (o ReportRecurrencePeriodPtrOutput) ToReportRecurrencePeriodPtrOutput() ReportRecurrencePeriodPtrOutput {
	return o
}

func (o ReportRecurrencePeriodPtrOutput) ToReportRecurrencePeriodPtrOutputWithContext(ctx context.Context) ReportRecurrencePeriodPtrOutput {
	return o
}

func (o ReportRecurrencePeriodPtrOutput) Elem() ReportRecurrencePeriodOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriod) ReportRecurrencePeriod {
		if v != nil {
			return *v
		}
		var ret ReportRecurrencePeriod
		return ret
	}).(ReportRecurrencePeriodOutput)
}

func (o ReportRecurrencePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportRecurrencePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriod) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type ReportRecurrencePeriodResponse struct {
	From string  `pulumi:"from"`
	To   *string `pulumi:"to"`
}

type ReportRecurrencePeriodResponseOutput struct{ *pulumi.OutputState }

func (ReportRecurrencePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportRecurrencePeriodResponse)(nil)).Elem()
}

func (o ReportRecurrencePeriodResponseOutput) ToReportRecurrencePeriodResponseOutput() ReportRecurrencePeriodResponseOutput {
	return o
}

func (o ReportRecurrencePeriodResponseOutput) ToReportRecurrencePeriodResponseOutputWithContext(ctx context.Context) ReportRecurrencePeriodResponseOutput {
	return o
}

func (o ReportRecurrencePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportRecurrencePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportRecurrencePeriodResponseOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportRecurrencePeriodResponse) *string { return v.To }).(pulumi.StringPtrOutput)
}

type ReportRecurrencePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportRecurrencePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportRecurrencePeriodResponse)(nil)).Elem()
}

func (o ReportRecurrencePeriodResponsePtrOutput) ToReportRecurrencePeriodResponsePtrOutput() ReportRecurrencePeriodResponsePtrOutput {
	return o
}

func (o ReportRecurrencePeriodResponsePtrOutput) ToReportRecurrencePeriodResponsePtrOutputWithContext(ctx context.Context) ReportRecurrencePeriodResponsePtrOutput {
	return o
}

func (o ReportRecurrencePeriodResponsePtrOutput) Elem() ReportRecurrencePeriodResponseOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriodResponse) ReportRecurrencePeriodResponse {
		if v != nil {
			return *v
		}
		var ret ReportRecurrencePeriodResponse
		return ret
	}).(ReportRecurrencePeriodResponseOutput)
}

func (o ReportRecurrencePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportRecurrencePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportRecurrencePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return v.To
	}).(pulumi.StringPtrOutput)
}

type ReportSchedule struct {
	Recurrence       string                  `pulumi:"recurrence"`
	RecurrencePeriod *ReportRecurrencePeriod `pulumi:"recurrencePeriod"`
	Status           *string                 `pulumi:"status"`
}





type ReportScheduleInput interface {
	pulumi.Input

	ToReportScheduleOutput() ReportScheduleOutput
	ToReportScheduleOutputWithContext(context.Context) ReportScheduleOutput
}

type ReportScheduleArgs struct {
	Recurrence       pulumi.StringInput             `pulumi:"recurrence"`
	RecurrencePeriod ReportRecurrencePeriodPtrInput `pulumi:"recurrencePeriod"`
	Status           pulumi.StringPtrInput          `pulumi:"status"`
}

func (ReportScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportSchedule)(nil)).Elem()
}

func (i ReportScheduleArgs) ToReportScheduleOutput() ReportScheduleOutput {
	return i.ToReportScheduleOutputWithContext(context.Background())
}

func (i ReportScheduleArgs) ToReportScheduleOutputWithContext(ctx context.Context) ReportScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportScheduleOutput)
}

func (i ReportScheduleArgs) ToReportSchedulePtrOutput() ReportSchedulePtrOutput {
	return i.ToReportSchedulePtrOutputWithContext(context.Background())
}

func (i ReportScheduleArgs) ToReportSchedulePtrOutputWithContext(ctx context.Context) ReportSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportScheduleOutput).ToReportSchedulePtrOutputWithContext(ctx)
}









type ReportSchedulePtrInput interface {
	pulumi.Input

	ToReportSchedulePtrOutput() ReportSchedulePtrOutput
	ToReportSchedulePtrOutputWithContext(context.Context) ReportSchedulePtrOutput
}

type reportSchedulePtrType ReportScheduleArgs

func ReportSchedulePtr(v *ReportScheduleArgs) ReportSchedulePtrInput {
	return (*reportSchedulePtrType)(v)
}

func (*reportSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportSchedule)(nil)).Elem()
}

func (i *reportSchedulePtrType) ToReportSchedulePtrOutput() ReportSchedulePtrOutput {
	return i.ToReportSchedulePtrOutputWithContext(context.Background())
}

func (i *reportSchedulePtrType) ToReportSchedulePtrOutputWithContext(ctx context.Context) ReportSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportSchedulePtrOutput)
}

type ReportScheduleOutput struct{ *pulumi.OutputState }

func (ReportScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportSchedule)(nil)).Elem()
}

func (o ReportScheduleOutput) ToReportScheduleOutput() ReportScheduleOutput {
	return o
}

func (o ReportScheduleOutput) ToReportScheduleOutputWithContext(ctx context.Context) ReportScheduleOutput {
	return o
}

func (o ReportScheduleOutput) ToReportSchedulePtrOutput() ReportSchedulePtrOutput {
	return o.ToReportSchedulePtrOutputWithContext(context.Background())
}

func (o ReportScheduleOutput) ToReportSchedulePtrOutputWithContext(ctx context.Context) ReportSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportSchedule) *ReportSchedule {
		return &v
	}).(ReportSchedulePtrOutput)
}

func (o ReportScheduleOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v ReportSchedule) string { return v.Recurrence }).(pulumi.StringOutput)
}

func (o ReportScheduleOutput) RecurrencePeriod() ReportRecurrencePeriodPtrOutput {
	return o.ApplyT(func(v ReportSchedule) *ReportRecurrencePeriod { return v.RecurrencePeriod }).(ReportRecurrencePeriodPtrOutput)
}

func (o ReportScheduleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportSchedule) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ReportSchedulePtrOutput struct{ *pulumi.OutputState }

func (ReportSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportSchedule)(nil)).Elem()
}

func (o ReportSchedulePtrOutput) ToReportSchedulePtrOutput() ReportSchedulePtrOutput {
	return o
}

func (o ReportSchedulePtrOutput) ToReportSchedulePtrOutputWithContext(ctx context.Context) ReportSchedulePtrOutput {
	return o
}

func (o ReportSchedulePtrOutput) Elem() ReportScheduleOutput {
	return o.ApplyT(func(v *ReportSchedule) ReportSchedule {
		if v != nil {
			return *v
		}
		var ret ReportSchedule
		return ret
	}).(ReportScheduleOutput)
}

func (o ReportSchedulePtrOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportSchedule) *string {
		if v == nil {
			return nil
		}
		return &v.Recurrence
	}).(pulumi.StringPtrOutput)
}

func (o ReportSchedulePtrOutput) RecurrencePeriod() ReportRecurrencePeriodPtrOutput {
	return o.ApplyT(func(v *ReportSchedule) *ReportRecurrencePeriod {
		if v == nil {
			return nil
		}
		return v.RecurrencePeriod
	}).(ReportRecurrencePeriodPtrOutput)
}

func (o ReportSchedulePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportSchedule) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ReportScheduleResponse struct {
	Recurrence       string                          `pulumi:"recurrence"`
	RecurrencePeriod *ReportRecurrencePeriodResponse `pulumi:"recurrencePeriod"`
	Status           *string                         `pulumi:"status"`
}

type ReportScheduleResponseOutput struct{ *pulumi.OutputState }

func (ReportScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportScheduleResponse)(nil)).Elem()
}

func (o ReportScheduleResponseOutput) ToReportScheduleResponseOutput() ReportScheduleResponseOutput {
	return o
}

func (o ReportScheduleResponseOutput) ToReportScheduleResponseOutputWithContext(ctx context.Context) ReportScheduleResponseOutput {
	return o
}

func (o ReportScheduleResponseOutput) Recurrence() pulumi.StringOutput {
	return o.ApplyT(func(v ReportScheduleResponse) string { return v.Recurrence }).(pulumi.StringOutput)
}

func (o ReportScheduleResponseOutput) RecurrencePeriod() ReportRecurrencePeriodResponsePtrOutput {
	return o.ApplyT(func(v ReportScheduleResponse) *ReportRecurrencePeriodResponse { return v.RecurrencePeriod }).(ReportRecurrencePeriodResponsePtrOutput)
}

func (o ReportScheduleResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReportScheduleResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ReportScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportScheduleResponse)(nil)).Elem()
}

func (o ReportScheduleResponsePtrOutput) ToReportScheduleResponsePtrOutput() ReportScheduleResponsePtrOutput {
	return o
}

func (o ReportScheduleResponsePtrOutput) ToReportScheduleResponsePtrOutputWithContext(ctx context.Context) ReportScheduleResponsePtrOutput {
	return o
}

func (o ReportScheduleResponsePtrOutput) Elem() ReportScheduleResponseOutput {
	return o.ApplyT(func(v *ReportScheduleResponse) ReportScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ReportScheduleResponse
		return ret
	}).(ReportScheduleResponseOutput)
}

func (o ReportScheduleResponsePtrOutput) Recurrence() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Recurrence
	}).(pulumi.StringPtrOutput)
}

func (o ReportScheduleResponsePtrOutput) RecurrencePeriod() ReportRecurrencePeriodResponsePtrOutput {
	return o.ApplyT(func(v *ReportScheduleResponse) *ReportRecurrencePeriodResponse {
		if v == nil {
			return nil
		}
		return v.RecurrencePeriod
	}).(ReportRecurrencePeriodResponsePtrOutput)
}

func (o ReportScheduleResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ReportTimePeriod struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}





type ReportTimePeriodInput interface {
	pulumi.Input

	ToReportTimePeriodOutput() ReportTimePeriodOutput
	ToReportTimePeriodOutputWithContext(context.Context) ReportTimePeriodOutput
}

type ReportTimePeriodArgs struct {
	From pulumi.StringInput `pulumi:"from"`
	To   pulumi.StringInput `pulumi:"to"`
}

func (ReportTimePeriodArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportTimePeriod)(nil)).Elem()
}

func (i ReportTimePeriodArgs) ToReportTimePeriodOutput() ReportTimePeriodOutput {
	return i.ToReportTimePeriodOutputWithContext(context.Background())
}

func (i ReportTimePeriodArgs) ToReportTimePeriodOutputWithContext(ctx context.Context) ReportTimePeriodOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportTimePeriodOutput)
}

func (i ReportTimePeriodArgs) ToReportTimePeriodPtrOutput() ReportTimePeriodPtrOutput {
	return i.ToReportTimePeriodPtrOutputWithContext(context.Background())
}

func (i ReportTimePeriodArgs) ToReportTimePeriodPtrOutputWithContext(ctx context.Context) ReportTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportTimePeriodOutput).ToReportTimePeriodPtrOutputWithContext(ctx)
}









type ReportTimePeriodPtrInput interface {
	pulumi.Input

	ToReportTimePeriodPtrOutput() ReportTimePeriodPtrOutput
	ToReportTimePeriodPtrOutputWithContext(context.Context) ReportTimePeriodPtrOutput
}

type reportTimePeriodPtrType ReportTimePeriodArgs

func ReportTimePeriodPtr(v *ReportTimePeriodArgs) ReportTimePeriodPtrInput {
	return (*reportTimePeriodPtrType)(v)
}

func (*reportTimePeriodPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportTimePeriod)(nil)).Elem()
}

func (i *reportTimePeriodPtrType) ToReportTimePeriodPtrOutput() ReportTimePeriodPtrOutput {
	return i.ToReportTimePeriodPtrOutputWithContext(context.Background())
}

func (i *reportTimePeriodPtrType) ToReportTimePeriodPtrOutputWithContext(ctx context.Context) ReportTimePeriodPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReportTimePeriodPtrOutput)
}

type ReportTimePeriodOutput struct{ *pulumi.OutputState }

func (ReportTimePeriodOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportTimePeriod)(nil)).Elem()
}

func (o ReportTimePeriodOutput) ToReportTimePeriodOutput() ReportTimePeriodOutput {
	return o
}

func (o ReportTimePeriodOutput) ToReportTimePeriodOutputWithContext(ctx context.Context) ReportTimePeriodOutput {
	return o
}

func (o ReportTimePeriodOutput) ToReportTimePeriodPtrOutput() ReportTimePeriodPtrOutput {
	return o.ToReportTimePeriodPtrOutputWithContext(context.Background())
}

func (o ReportTimePeriodOutput) ToReportTimePeriodPtrOutputWithContext(ctx context.Context) ReportTimePeriodPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportTimePeriod) *ReportTimePeriod {
		return &v
	}).(ReportTimePeriodPtrOutput)
}

func (o ReportTimePeriodOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportTimePeriod) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportTimePeriodOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v ReportTimePeriod) string { return v.To }).(pulumi.StringOutput)
}

type ReportTimePeriodPtrOutput struct{ *pulumi.OutputState }

func (ReportTimePeriodPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportTimePeriod)(nil)).Elem()
}

func (o ReportTimePeriodPtrOutput) ToReportTimePeriodPtrOutput() ReportTimePeriodPtrOutput {
	return o
}

func (o ReportTimePeriodPtrOutput) ToReportTimePeriodPtrOutputWithContext(ctx context.Context) ReportTimePeriodPtrOutput {
	return o
}

func (o ReportTimePeriodPtrOutput) Elem() ReportTimePeriodOutput {
	return o.ApplyT(func(v *ReportTimePeriod) ReportTimePeriod {
		if v != nil {
			return *v
		}
		var ret ReportTimePeriod
		return ret
	}).(ReportTimePeriodOutput)
}

func (o ReportTimePeriodPtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportTimePeriodPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportTimePeriod) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

type ReportTimePeriodResponse struct {
	From string `pulumi:"from"`
	To   string `pulumi:"to"`
}

type ReportTimePeriodResponseOutput struct{ *pulumi.OutputState }

func (ReportTimePeriodResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportTimePeriodResponse)(nil)).Elem()
}

func (o ReportTimePeriodResponseOutput) ToReportTimePeriodResponseOutput() ReportTimePeriodResponseOutput {
	return o
}

func (o ReportTimePeriodResponseOutput) ToReportTimePeriodResponseOutputWithContext(ctx context.Context) ReportTimePeriodResponseOutput {
	return o
}

func (o ReportTimePeriodResponseOutput) From() pulumi.StringOutput {
	return o.ApplyT(func(v ReportTimePeriodResponse) string { return v.From }).(pulumi.StringOutput)
}

func (o ReportTimePeriodResponseOutput) To() pulumi.StringOutput {
	return o.ApplyT(func(v ReportTimePeriodResponse) string { return v.To }).(pulumi.StringOutput)
}

type ReportTimePeriodResponsePtrOutput struct{ *pulumi.OutputState }

func (ReportTimePeriodResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportTimePeriodResponse)(nil)).Elem()
}

func (o ReportTimePeriodResponsePtrOutput) ToReportTimePeriodResponsePtrOutput() ReportTimePeriodResponsePtrOutput {
	return o
}

func (o ReportTimePeriodResponsePtrOutput) ToReportTimePeriodResponsePtrOutputWithContext(ctx context.Context) ReportTimePeriodResponsePtrOutput {
	return o
}

func (o ReportTimePeriodResponsePtrOutput) Elem() ReportTimePeriodResponseOutput {
	return o.ApplyT(func(v *ReportTimePeriodResponse) ReportTimePeriodResponse {
		if v != nil {
			return *v
		}
		var ret ReportTimePeriodResponse
		return ret
	}).(ReportTimePeriodResponseOutput)
}

func (o ReportTimePeriodResponsePtrOutput) From() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.From
	}).(pulumi.StringPtrOutput)
}

func (o ReportTimePeriodResponsePtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReportTimePeriodResponse) *string {
		if v == nil {
			return nil
		}
		return &v.To
	}).(pulumi.StringPtrOutput)
}

type ScheduleProperties struct {
	DayOfMonth   *int     `pulumi:"dayOfMonth"`
	DaysOfWeek   []string `pulumi:"daysOfWeek"`
	EndDate      string   `pulumi:"endDate"`
	Frequency    string   `pulumi:"frequency"`
	HourOfDay    *int     `pulumi:"hourOfDay"`
	StartDate    string   `pulumi:"startDate"`
	WeeksOfMonth []string `pulumi:"weeksOfMonth"`
}





type SchedulePropertiesInput interface {
	pulumi.Input

	ToSchedulePropertiesOutput() SchedulePropertiesOutput
	ToSchedulePropertiesOutputWithContext(context.Context) SchedulePropertiesOutput
}

type SchedulePropertiesArgs struct {
	DayOfMonth   pulumi.IntPtrInput      `pulumi:"dayOfMonth"`
	DaysOfWeek   pulumi.StringArrayInput `pulumi:"daysOfWeek"`
	EndDate      pulumi.StringInput      `pulumi:"endDate"`
	Frequency    pulumi.StringInput      `pulumi:"frequency"`
	HourOfDay    pulumi.IntPtrInput      `pulumi:"hourOfDay"`
	StartDate    pulumi.StringInput      `pulumi:"startDate"`
	WeeksOfMonth pulumi.StringArrayInput `pulumi:"weeksOfMonth"`
}

func (SchedulePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleProperties)(nil)).Elem()
}

func (i SchedulePropertiesArgs) ToSchedulePropertiesOutput() SchedulePropertiesOutput {
	return i.ToSchedulePropertiesOutputWithContext(context.Background())
}

func (i SchedulePropertiesArgs) ToSchedulePropertiesOutputWithContext(ctx context.Context) SchedulePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulePropertiesOutput)
}

type SchedulePropertiesOutput struct{ *pulumi.OutputState }

func (SchedulePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleProperties)(nil)).Elem()
}

func (o SchedulePropertiesOutput) ToSchedulePropertiesOutput() SchedulePropertiesOutput {
	return o
}

func (o SchedulePropertiesOutput) ToSchedulePropertiesOutputWithContext(ctx context.Context) SchedulePropertiesOutput {
	return o
}

func (o SchedulePropertiesOutput) DayOfMonth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *int { return v.DayOfMonth }).(pulumi.IntPtrOutput)
}

func (o SchedulePropertiesOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScheduleProperties) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o SchedulePropertiesOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleProperties) string { return v.EndDate }).(pulumi.StringOutput)
}

func (o SchedulePropertiesOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleProperties) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o SchedulePropertiesOutput) HourOfDay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *int { return v.HourOfDay }).(pulumi.IntPtrOutput)
}

func (o SchedulePropertiesOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleProperties) string { return v.StartDate }).(pulumi.StringOutput)
}

func (o SchedulePropertiesOutput) WeeksOfMonth() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScheduleProperties) []string { return v.WeeksOfMonth }).(pulumi.StringArrayOutput)
}

type SchedulePropertiesResponse struct {
	DayOfMonth   *int     `pulumi:"dayOfMonth"`
	DaysOfWeek   []string `pulumi:"daysOfWeek"`
	EndDate      string   `pulumi:"endDate"`
	Frequency    string   `pulumi:"frequency"`
	HourOfDay    *int     `pulumi:"hourOfDay"`
	StartDate    string   `pulumi:"startDate"`
	WeeksOfMonth []string `pulumi:"weeksOfMonth"`
}

type SchedulePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SchedulePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchedulePropertiesResponse)(nil)).Elem()
}

func (o SchedulePropertiesResponseOutput) ToSchedulePropertiesResponseOutput() SchedulePropertiesResponseOutput {
	return o
}

func (o SchedulePropertiesResponseOutput) ToSchedulePropertiesResponseOutputWithContext(ctx context.Context) SchedulePropertiesResponseOutput {
	return o
}

func (o SchedulePropertiesResponseOutput) DayOfMonth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *int { return v.DayOfMonth }).(pulumi.IntPtrOutput)
}

func (o SchedulePropertiesResponseOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o SchedulePropertiesResponseOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) string { return v.EndDate }).(pulumi.StringOutput)
}

func (o SchedulePropertiesResponseOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o SchedulePropertiesResponseOutput) HourOfDay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *int { return v.HourOfDay }).(pulumi.IntPtrOutput)
}

func (o SchedulePropertiesResponseOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) string { return v.StartDate }).(pulumi.StringOutput)
}

func (o SchedulePropertiesResponseOutput) WeeksOfMonth() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) []string { return v.WeeksOfMonth }).(pulumi.StringArrayOutput)
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

type SourceCostAllocationResource struct {
	Name         string   `pulumi:"name"`
	ResourceType string   `pulumi:"resourceType"`
	Values       []string `pulumi:"values"`
}





type SourceCostAllocationResourceInput interface {
	pulumi.Input

	ToSourceCostAllocationResourceOutput() SourceCostAllocationResourceOutput
	ToSourceCostAllocationResourceOutputWithContext(context.Context) SourceCostAllocationResourceOutput
}

type SourceCostAllocationResourceArgs struct {
	Name         pulumi.StringInput      `pulumi:"name"`
	ResourceType pulumi.StringInput      `pulumi:"resourceType"`
	Values       pulumi.StringArrayInput `pulumi:"values"`
}

func (SourceCostAllocationResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCostAllocationResource)(nil)).Elem()
}

func (i SourceCostAllocationResourceArgs) ToSourceCostAllocationResourceOutput() SourceCostAllocationResourceOutput {
	return i.ToSourceCostAllocationResourceOutputWithContext(context.Background())
}

func (i SourceCostAllocationResourceArgs) ToSourceCostAllocationResourceOutputWithContext(ctx context.Context) SourceCostAllocationResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCostAllocationResourceOutput)
}





type SourceCostAllocationResourceArrayInput interface {
	pulumi.Input

	ToSourceCostAllocationResourceArrayOutput() SourceCostAllocationResourceArrayOutput
	ToSourceCostAllocationResourceArrayOutputWithContext(context.Context) SourceCostAllocationResourceArrayOutput
}

type SourceCostAllocationResourceArray []SourceCostAllocationResourceInput

func (SourceCostAllocationResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceCostAllocationResource)(nil)).Elem()
}

func (i SourceCostAllocationResourceArray) ToSourceCostAllocationResourceArrayOutput() SourceCostAllocationResourceArrayOutput {
	return i.ToSourceCostAllocationResourceArrayOutputWithContext(context.Background())
}

func (i SourceCostAllocationResourceArray) ToSourceCostAllocationResourceArrayOutputWithContext(ctx context.Context) SourceCostAllocationResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceCostAllocationResourceArrayOutput)
}

type SourceCostAllocationResourceOutput struct{ *pulumi.OutputState }

func (SourceCostAllocationResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCostAllocationResource)(nil)).Elem()
}

func (o SourceCostAllocationResourceOutput) ToSourceCostAllocationResourceOutput() SourceCostAllocationResourceOutput {
	return o
}

func (o SourceCostAllocationResourceOutput) ToSourceCostAllocationResourceOutputWithContext(ctx context.Context) SourceCostAllocationResourceOutput {
	return o
}

func (o SourceCostAllocationResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SourceCostAllocationResource) string { return v.Name }).(pulumi.StringOutput)
}

func (o SourceCostAllocationResourceOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SourceCostAllocationResource) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SourceCostAllocationResourceOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SourceCostAllocationResource) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type SourceCostAllocationResourceArrayOutput struct{ *pulumi.OutputState }

func (SourceCostAllocationResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceCostAllocationResource)(nil)).Elem()
}

func (o SourceCostAllocationResourceArrayOutput) ToSourceCostAllocationResourceArrayOutput() SourceCostAllocationResourceArrayOutput {
	return o
}

func (o SourceCostAllocationResourceArrayOutput) ToSourceCostAllocationResourceArrayOutputWithContext(ctx context.Context) SourceCostAllocationResourceArrayOutput {
	return o
}

func (o SourceCostAllocationResourceArrayOutput) Index(i pulumi.IntInput) SourceCostAllocationResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SourceCostAllocationResource {
		return vs[0].([]SourceCostAllocationResource)[vs[1].(int)]
	}).(SourceCostAllocationResourceOutput)
}

type SourceCostAllocationResourceResponse struct {
	Name         string   `pulumi:"name"`
	ResourceType string   `pulumi:"resourceType"`
	Values       []string `pulumi:"values"`
}

type SourceCostAllocationResourceResponseOutput struct{ *pulumi.OutputState }

func (SourceCostAllocationResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceCostAllocationResourceResponse)(nil)).Elem()
}

func (o SourceCostAllocationResourceResponseOutput) ToSourceCostAllocationResourceResponseOutput() SourceCostAllocationResourceResponseOutput {
	return o
}

func (o SourceCostAllocationResourceResponseOutput) ToSourceCostAllocationResourceResponseOutputWithContext(ctx context.Context) SourceCostAllocationResourceResponseOutput {
	return o
}

func (o SourceCostAllocationResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SourceCostAllocationResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SourceCostAllocationResourceResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v SourceCostAllocationResourceResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o SourceCostAllocationResourceResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SourceCostAllocationResourceResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type SourceCostAllocationResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SourceCostAllocationResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SourceCostAllocationResourceResponse)(nil)).Elem()
}

func (o SourceCostAllocationResourceResponseArrayOutput) ToSourceCostAllocationResourceResponseArrayOutput() SourceCostAllocationResourceResponseArrayOutput {
	return o
}

func (o SourceCostAllocationResourceResponseArrayOutput) ToSourceCostAllocationResourceResponseArrayOutputWithContext(ctx context.Context) SourceCostAllocationResourceResponseArrayOutput {
	return o
}

func (o SourceCostAllocationResourceResponseArrayOutput) Index(i pulumi.IntInput) SourceCostAllocationResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SourceCostAllocationResourceResponse {
		return vs[0].([]SourceCostAllocationResourceResponse)[vs[1].(int)]
	}).(SourceCostAllocationResourceResponseOutput)
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

type TargetCostAllocationResource struct {
	Name         string                     `pulumi:"name"`
	PolicyType   string                     `pulumi:"policyType"`
	ResourceType string                     `pulumi:"resourceType"`
	Values       []CostAllocationProportion `pulumi:"values"`
}





type TargetCostAllocationResourceInput interface {
	pulumi.Input

	ToTargetCostAllocationResourceOutput() TargetCostAllocationResourceOutput
	ToTargetCostAllocationResourceOutputWithContext(context.Context) TargetCostAllocationResourceOutput
}

type TargetCostAllocationResourceArgs struct {
	Name         pulumi.StringInput                 `pulumi:"name"`
	PolicyType   pulumi.StringInput                 `pulumi:"policyType"`
	ResourceType pulumi.StringInput                 `pulumi:"resourceType"`
	Values       CostAllocationProportionArrayInput `pulumi:"values"`
}

func (TargetCostAllocationResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetCostAllocationResource)(nil)).Elem()
}

func (i TargetCostAllocationResourceArgs) ToTargetCostAllocationResourceOutput() TargetCostAllocationResourceOutput {
	return i.ToTargetCostAllocationResourceOutputWithContext(context.Background())
}

func (i TargetCostAllocationResourceArgs) ToTargetCostAllocationResourceOutputWithContext(ctx context.Context) TargetCostAllocationResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetCostAllocationResourceOutput)
}





type TargetCostAllocationResourceArrayInput interface {
	pulumi.Input

	ToTargetCostAllocationResourceArrayOutput() TargetCostAllocationResourceArrayOutput
	ToTargetCostAllocationResourceArrayOutputWithContext(context.Context) TargetCostAllocationResourceArrayOutput
}

type TargetCostAllocationResourceArray []TargetCostAllocationResourceInput

func (TargetCostAllocationResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetCostAllocationResource)(nil)).Elem()
}

func (i TargetCostAllocationResourceArray) ToTargetCostAllocationResourceArrayOutput() TargetCostAllocationResourceArrayOutput {
	return i.ToTargetCostAllocationResourceArrayOutputWithContext(context.Background())
}

func (i TargetCostAllocationResourceArray) ToTargetCostAllocationResourceArrayOutputWithContext(ctx context.Context) TargetCostAllocationResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetCostAllocationResourceArrayOutput)
}

type TargetCostAllocationResourceOutput struct{ *pulumi.OutputState }

func (TargetCostAllocationResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetCostAllocationResource)(nil)).Elem()
}

func (o TargetCostAllocationResourceOutput) ToTargetCostAllocationResourceOutput() TargetCostAllocationResourceOutput {
	return o
}

func (o TargetCostAllocationResourceOutput) ToTargetCostAllocationResourceOutputWithContext(ctx context.Context) TargetCostAllocationResourceOutput {
	return o
}

func (o TargetCostAllocationResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResource) string { return v.Name }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResource) string { return v.PolicyType }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResource) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceOutput) Values() CostAllocationProportionArrayOutput {
	return o.ApplyT(func(v TargetCostAllocationResource) []CostAllocationProportion { return v.Values }).(CostAllocationProportionArrayOutput)
}

type TargetCostAllocationResourceArrayOutput struct{ *pulumi.OutputState }

func (TargetCostAllocationResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetCostAllocationResource)(nil)).Elem()
}

func (o TargetCostAllocationResourceArrayOutput) ToTargetCostAllocationResourceArrayOutput() TargetCostAllocationResourceArrayOutput {
	return o
}

func (o TargetCostAllocationResourceArrayOutput) ToTargetCostAllocationResourceArrayOutputWithContext(ctx context.Context) TargetCostAllocationResourceArrayOutput {
	return o
}

func (o TargetCostAllocationResourceArrayOutput) Index(i pulumi.IntInput) TargetCostAllocationResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetCostAllocationResource {
		return vs[0].([]TargetCostAllocationResource)[vs[1].(int)]
	}).(TargetCostAllocationResourceOutput)
}

type TargetCostAllocationResourceResponse struct {
	Name         string                             `pulumi:"name"`
	PolicyType   string                             `pulumi:"policyType"`
	ResourceType string                             `pulumi:"resourceType"`
	Values       []CostAllocationProportionResponse `pulumi:"values"`
}

type TargetCostAllocationResourceResponseOutput struct{ *pulumi.OutputState }

func (TargetCostAllocationResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetCostAllocationResourceResponse)(nil)).Elem()
}

func (o TargetCostAllocationResourceResponseOutput) ToTargetCostAllocationResourceResponseOutput() TargetCostAllocationResourceResponseOutput {
	return o
}

func (o TargetCostAllocationResourceResponseOutput) ToTargetCostAllocationResourceResponseOutputWithContext(ctx context.Context) TargetCostAllocationResourceResponseOutput {
	return o
}

func (o TargetCostAllocationResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceResponseOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResourceResponse) string { return v.PolicyType }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v TargetCostAllocationResourceResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o TargetCostAllocationResourceResponseOutput) Values() CostAllocationProportionResponseArrayOutput {
	return o.ApplyT(func(v TargetCostAllocationResourceResponse) []CostAllocationProportionResponse { return v.Values }).(CostAllocationProportionResponseArrayOutput)
}

type TargetCostAllocationResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (TargetCostAllocationResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TargetCostAllocationResourceResponse)(nil)).Elem()
}

func (o TargetCostAllocationResourceResponseArrayOutput) ToTargetCostAllocationResourceResponseArrayOutput() TargetCostAllocationResourceResponseArrayOutput {
	return o
}

func (o TargetCostAllocationResourceResponseArrayOutput) ToTargetCostAllocationResourceResponseArrayOutputWithContext(ctx context.Context) TargetCostAllocationResourceResponseArrayOutput {
	return o
}

func (o TargetCostAllocationResourceResponseArrayOutput) Index(i pulumi.IntInput) TargetCostAllocationResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TargetCostAllocationResourceResponse {
		return vs[0].([]TargetCostAllocationResourceResponse)[vs[1].(int)]
	}).(TargetCostAllocationResourceResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(CommonExportPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CommonExportPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectorCollectionErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(ConnectorCollectionErrorInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectorCollectionInfoResponseOutput{})
	pulumi.RegisterOutputType(CostAllocationProportionOutput{})
	pulumi.RegisterOutputType(CostAllocationProportionArrayOutput{})
	pulumi.RegisterOutputType(CostAllocationProportionResponseOutput{})
	pulumi.RegisterOutputType(CostAllocationProportionResponseArrayOutput{})
	pulumi.RegisterOutputType(CostAllocationRuleDetailsOutput{})
	pulumi.RegisterOutputType(CostAllocationRuleDetailsPtrOutput{})
	pulumi.RegisterOutputType(CostAllocationRuleDetailsResponseOutput{})
	pulumi.RegisterOutputType(CostAllocationRulePropertiesOutput{})
	pulumi.RegisterOutputType(CostAllocationRulePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CostAllocationRulePropertiesResponseOutput{})
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
	pulumi.RegisterOutputType(FileDestinationOutput{})
	pulumi.RegisterOutputType(FileDestinationPtrOutput{})
	pulumi.RegisterOutputType(FileDestinationResponseOutput{})
	pulumi.RegisterOutputType(FileDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(KpiPropertiesOutput{})
	pulumi.RegisterOutputType(KpiPropertiesArrayOutput{})
	pulumi.RegisterOutputType(KpiPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KpiPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(NotificationPropertiesOutput{})
	pulumi.RegisterOutputType(NotificationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PivotPropertiesOutput{})
	pulumi.RegisterOutputType(PivotPropertiesArrayOutput{})
	pulumi.RegisterOutputType(PivotPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PivotPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportAggregationOutput{})
	pulumi.RegisterOutputType(ReportAggregationMapOutput{})
	pulumi.RegisterOutputType(ReportAggregationResponseOutput{})
	pulumi.RegisterOutputType(ReportAggregationResponseMapOutput{})
	pulumi.RegisterOutputType(ReportComparisonExpressionOutput{})
	pulumi.RegisterOutputType(ReportComparisonExpressionPtrOutput{})
	pulumi.RegisterOutputType(ReportComparisonExpressionResponseOutput{})
	pulumi.RegisterOutputType(ReportComparisonExpressionResponsePtrOutput{})
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
	pulumi.RegisterOutputType(ReportDatasetOutput{})
	pulumi.RegisterOutputType(ReportDatasetPtrOutput{})
	pulumi.RegisterOutputType(ReportDatasetConfigurationOutput{})
	pulumi.RegisterOutputType(ReportDatasetConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ReportDatasetConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ReportDatasetConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportDatasetResponseOutput{})
	pulumi.RegisterOutputType(ReportDatasetResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportDefinitionOutput{})
	pulumi.RegisterOutputType(ReportDefinitionResponseOutput{})
	pulumi.RegisterOutputType(ReportDeliveryDestinationOutput{})
	pulumi.RegisterOutputType(ReportDeliveryDestinationResponseOutput{})
	pulumi.RegisterOutputType(ReportDeliveryInfoOutput{})
	pulumi.RegisterOutputType(ReportDeliveryInfoResponseOutput{})
	pulumi.RegisterOutputType(ReportFilterOutput{})
	pulumi.RegisterOutputType(ReportFilterPtrOutput{})
	pulumi.RegisterOutputType(ReportFilterArrayOutput{})
	pulumi.RegisterOutputType(ReportFilterResponseOutput{})
	pulumi.RegisterOutputType(ReportFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportFilterResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportGroupingOutput{})
	pulumi.RegisterOutputType(ReportGroupingArrayOutput{})
	pulumi.RegisterOutputType(ReportGroupingResponseOutput{})
	pulumi.RegisterOutputType(ReportGroupingResponseArrayOutput{})
	pulumi.RegisterOutputType(ReportRecurrencePeriodOutput{})
	pulumi.RegisterOutputType(ReportRecurrencePeriodPtrOutput{})
	pulumi.RegisterOutputType(ReportRecurrencePeriodResponseOutput{})
	pulumi.RegisterOutputType(ReportRecurrencePeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportScheduleOutput{})
	pulumi.RegisterOutputType(ReportSchedulePtrOutput{})
	pulumi.RegisterOutputType(ReportScheduleResponseOutput{})
	pulumi.RegisterOutputType(ReportScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(ReportTimePeriodOutput{})
	pulumi.RegisterOutputType(ReportTimePeriodPtrOutput{})
	pulumi.RegisterOutputType(ReportTimePeriodResponseOutput{})
	pulumi.RegisterOutputType(ReportTimePeriodResponsePtrOutput{})
	pulumi.RegisterOutputType(SchedulePropertiesOutput{})
	pulumi.RegisterOutputType(SchedulePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SettingsPropertiesCacheOutput{})
	pulumi.RegisterOutputType(SettingsPropertiesCacheArrayOutput{})
	pulumi.RegisterOutputType(SettingsPropertiesResponseCacheOutput{})
	pulumi.RegisterOutputType(SettingsPropertiesResponseCacheArrayOutput{})
	pulumi.RegisterOutputType(SourceCostAllocationResourceOutput{})
	pulumi.RegisterOutputType(SourceCostAllocationResourceArrayOutput{})
	pulumi.RegisterOutputType(SourceCostAllocationResourceResponseOutput{})
	pulumi.RegisterOutputType(SourceCostAllocationResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TargetCostAllocationResourceOutput{})
	pulumi.RegisterOutputType(TargetCostAllocationResourceArrayOutput{})
	pulumi.RegisterOutputType(TargetCostAllocationResourceResponseOutput{})
	pulumi.RegisterOutputType(TargetCostAllocationResourceResponseArrayOutput{})
}
