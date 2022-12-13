


package customerinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKpi(ctx *pulumi.Context, args *LookupKpiArgs, opts ...pulumi.InvokeOption) (*LookupKpiResult, error) {
	var rv LookupKpiResult
	err := ctx.Invoke("azure-native:customerinsights:getKpi", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupKpiArgs struct {
	HubName           string `pulumi:"hubName"`
	KpiName           string `pulumi:"kpiName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupKpiResult struct {
	Aliases                     []KpiAliasResponse                       `pulumi:"aliases"`
	CalculationWindow           string                                   `pulumi:"calculationWindow"`
	CalculationWindowFieldName  *string                                  `pulumi:"calculationWindowFieldName"`
	Description                 map[string]string                        `pulumi:"description"`
	DisplayName                 map[string]string                        `pulumi:"displayName"`
	EntityType                  string                                   `pulumi:"entityType"`
	EntityTypeName              string                                   `pulumi:"entityTypeName"`
	Expression                  string                                   `pulumi:"expression"`
	Extracts                    []KpiExtractResponse                     `pulumi:"extracts"`
	Filter                      *string                                  `pulumi:"filter"`
	Function                    string                                   `pulumi:"function"`
	GroupBy                     []string                                 `pulumi:"groupBy"`
	GroupByMetadata             []KpiGroupByMetadataResponse             `pulumi:"groupByMetadata"`
	Id                          string                                   `pulumi:"id"`
	KpiName                     string                                   `pulumi:"kpiName"`
	Name                        string                                   `pulumi:"name"`
	ParticipantProfilesMetadata []KpiParticipantProfilesMetadataResponse `pulumi:"participantProfilesMetadata"`
	ProvisioningState           string                                   `pulumi:"provisioningState"`
	TenantId                    string                                   `pulumi:"tenantId"`
	ThresHolds                  *KpiThresholdsResponse                   `pulumi:"thresHolds"`
	Type                        string                                   `pulumi:"type"`
	Unit                        *string                                  `pulumi:"unit"`
}

func LookupKpiOutput(ctx *pulumi.Context, args LookupKpiOutputArgs, opts ...pulumi.InvokeOption) LookupKpiResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupKpiResult, error) {
			args := v.(LookupKpiArgs)
			r, err := LookupKpi(ctx, &args, opts...)
			var s LookupKpiResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupKpiResultOutput)
}

type LookupKpiOutputArgs struct {
	HubName           pulumi.StringInput `pulumi:"hubName"`
	KpiName           pulumi.StringInput `pulumi:"kpiName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupKpiOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKpiArgs)(nil)).Elem()
}


type LookupKpiResultOutput struct{ *pulumi.OutputState }

func (LookupKpiResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKpiResult)(nil)).Elem()
}

func (o LookupKpiResultOutput) ToLookupKpiResultOutput() LookupKpiResultOutput {
	return o
}

func (o LookupKpiResultOutput) ToLookupKpiResultOutputWithContext(ctx context.Context) LookupKpiResultOutput {
	return o
}

func (o LookupKpiResultOutput) Aliases() KpiAliasResponseArrayOutput {
	return o.ApplyT(func(v LookupKpiResult) []KpiAliasResponse { return v.Aliases }).(KpiAliasResponseArrayOutput)
}

func (o LookupKpiResultOutput) CalculationWindow() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.CalculationWindow }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) CalculationWindowFieldName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKpiResult) *string { return v.CalculationWindowFieldName }).(pulumi.StringPtrOutput)
}

func (o LookupKpiResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKpiResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

func (o LookupKpiResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupKpiResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o LookupKpiResultOutput) EntityType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.EntityType }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) EntityTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.EntityTypeName }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.Expression }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) Extracts() KpiExtractResponseArrayOutput {
	return o.ApplyT(func(v LookupKpiResult) []KpiExtractResponse { return v.Extracts }).(KpiExtractResponseArrayOutput)
}

func (o LookupKpiResultOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKpiResult) *string { return v.Filter }).(pulumi.StringPtrOutput)
}

func (o LookupKpiResultOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.Function }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) GroupBy() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupKpiResult) []string { return v.GroupBy }).(pulumi.StringArrayOutput)
}

func (o LookupKpiResultOutput) GroupByMetadata() KpiGroupByMetadataResponseArrayOutput {
	return o.ApplyT(func(v LookupKpiResult) []KpiGroupByMetadataResponse { return v.GroupByMetadata }).(KpiGroupByMetadataResponseArrayOutput)
}

func (o LookupKpiResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) KpiName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.KpiName }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) ParticipantProfilesMetadata() KpiParticipantProfilesMetadataResponseArrayOutput {
	return o.ApplyT(func(v LookupKpiResult) []KpiParticipantProfilesMetadataResponse { return v.ParticipantProfilesMetadata }).(KpiParticipantProfilesMetadataResponseArrayOutput)
}

func (o LookupKpiResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) ThresHolds() KpiThresholdsResponsePtrOutput {
	return o.ApplyT(func(v LookupKpiResult) *KpiThresholdsResponse { return v.ThresHolds }).(KpiThresholdsResponsePtrOutput)
}

func (o LookupKpiResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKpiResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupKpiResultOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupKpiResult) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKpiResultOutput{})
}
