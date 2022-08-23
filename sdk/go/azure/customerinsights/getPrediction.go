


package customerinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrediction(ctx *pulumi.Context, args *LookupPredictionArgs, opts ...pulumi.InvokeOption) (*LookupPredictionResult, error) {
	var rv LookupPredictionResult
	err := ctx.Invoke("azure-native:customerinsights:getPrediction", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPredictionArgs struct {
	HubName           string `pulumi:"hubName"`
	PredictionName    string `pulumi:"predictionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPredictionResult struct {
	AutoAnalyze               bool                                      `pulumi:"autoAnalyze"`
	Description               map[string]string                         `pulumi:"description"`
	DisplayName               map[string]string                         `pulumi:"displayName"`
	Grades                    []PredictionResponseGrades                `pulumi:"grades"`
	Id                        string                                    `pulumi:"id"`
	InvolvedInteractionTypes  []string                                  `pulumi:"involvedInteractionTypes"`
	InvolvedKpiTypes          []string                                  `pulumi:"involvedKpiTypes"`
	InvolvedRelationships     []string                                  `pulumi:"involvedRelationships"`
	Mappings                  PredictionResponseMappings                `pulumi:"mappings"`
	Name                      string                                    `pulumi:"name"`
	NegativeOutcomeExpression string                                    `pulumi:"negativeOutcomeExpression"`
	PositiveOutcomeExpression string                                    `pulumi:"positiveOutcomeExpression"`
	PredictionName            *string                                   `pulumi:"predictionName"`
	PrimaryProfileType        string                                    `pulumi:"primaryProfileType"`
	ProvisioningState         string                                    `pulumi:"provisioningState"`
	ScopeExpression           string                                    `pulumi:"scopeExpression"`
	ScoreLabel                string                                    `pulumi:"scoreLabel"`
	SystemGeneratedEntities   PredictionResponseSystemGeneratedEntities `pulumi:"systemGeneratedEntities"`
	TenantId                  string                                    `pulumi:"tenantId"`
	Type                      string                                    `pulumi:"type"`
}

func LookupPredictionOutput(ctx *pulumi.Context, args LookupPredictionOutputArgs, opts ...pulumi.InvokeOption) LookupPredictionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPredictionResult, error) {
			args := v.(LookupPredictionArgs)
			r, err := LookupPrediction(ctx, &args, opts...)
			var s LookupPredictionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPredictionResultOutput)
}

type LookupPredictionOutputArgs struct {
	HubName           pulumi.StringInput `pulumi:"hubName"`
	PredictionName    pulumi.StringInput `pulumi:"predictionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPredictionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPredictionArgs)(nil)).Elem()
}


type LookupPredictionResultOutput struct{ *pulumi.OutputState }

func (LookupPredictionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPredictionResult)(nil)).Elem()
}

func (o LookupPredictionResultOutput) ToLookupPredictionResultOutput() LookupPredictionResultOutput {
	return o
}

func (o LookupPredictionResultOutput) ToLookupPredictionResultOutputWithContext(ctx context.Context) LookupPredictionResultOutput {
	return o
}

func (o LookupPredictionResultOutput) AutoAnalyze() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupPredictionResult) bool { return v.AutoAnalyze }).(pulumi.BoolOutput)
}

func (o LookupPredictionResultOutput) Description() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPredictionResult) map[string]string { return v.Description }).(pulumi.StringMapOutput)
}

func (o LookupPredictionResultOutput) DisplayName() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPredictionResult) map[string]string { return v.DisplayName }).(pulumi.StringMapOutput)
}

func (o LookupPredictionResultOutput) Grades() PredictionResponseGradesArrayOutput {
	return o.ApplyT(func(v LookupPredictionResult) []PredictionResponseGrades { return v.Grades }).(PredictionResponseGradesArrayOutput)
}

func (o LookupPredictionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPredictionResultOutput) InvolvedInteractionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPredictionResult) []string { return v.InvolvedInteractionTypes }).(pulumi.StringArrayOutput)
}

func (o LookupPredictionResultOutput) InvolvedKpiTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPredictionResult) []string { return v.InvolvedKpiTypes }).(pulumi.StringArrayOutput)
}

func (o LookupPredictionResultOutput) InvolvedRelationships() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPredictionResult) []string { return v.InvolvedRelationships }).(pulumi.StringArrayOutput)
}

func (o LookupPredictionResultOutput) Mappings() PredictionResponseMappingsOutput {
	return o.ApplyT(func(v LookupPredictionResult) PredictionResponseMappings { return v.Mappings }).(PredictionResponseMappingsOutput)
}

func (o LookupPredictionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPredictionResultOutput) NegativeOutcomeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.NegativeOutcomeExpression }).(pulumi.StringOutput)
}

func (o LookupPredictionResultOutput) PositiveOutcomeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.PositiveOutcomeExpression }).(pulumi.StringOutput)
}

func (o LookupPredictionResultOutput) PredictionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPredictionResult) *string { return v.PredictionName }).(pulumi.StringPtrOutput)
}

func (o LookupPredictionResultOutput) PrimaryProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.PrimaryProfileType }).(pulumi.StringOutput)
}

func (o LookupPredictionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupPredictionResultOutput) ScopeExpression() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.ScopeExpression }).(pulumi.StringOutput)
}

func (o LookupPredictionResultOutput) ScoreLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.ScoreLabel }).(pulumi.StringOutput)
}

func (o LookupPredictionResultOutput) SystemGeneratedEntities() PredictionResponseSystemGeneratedEntitiesOutput {
	return o.ApplyT(func(v LookupPredictionResult) PredictionResponseSystemGeneratedEntities {
		return v.SystemGeneratedEntities
	}).(PredictionResponseSystemGeneratedEntitiesOutput)
}

func (o LookupPredictionResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupPredictionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPredictionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPredictionResultOutput{})
}
