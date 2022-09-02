


package customerinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetPredictionTrainingResults(ctx *pulumi.Context, args *GetPredictionTrainingResultsArgs, opts ...pulumi.InvokeOption) (*GetPredictionTrainingResultsResult, error) {
	var rv GetPredictionTrainingResultsResult
	err := ctx.Invoke("azure-native:customerinsights:getPredictionTrainingResults", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetPredictionTrainingResultsArgs struct {
	HubName           string `pulumi:"hubName"`
	PredictionName    string `pulumi:"predictionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetPredictionTrainingResultsResult struct {
	CanonicalProfiles           []CanonicalProfileDefinitionResponse     `pulumi:"canonicalProfiles"`
	PredictionDistribution      PredictionDistributionDefinitionResponse `pulumi:"predictionDistribution"`
	PrimaryProfileInstanceCount float64                                  `pulumi:"primaryProfileInstanceCount"`
	ScoreName                   string                                   `pulumi:"scoreName"`
	TenantId                    string                                   `pulumi:"tenantId"`
}

func GetPredictionTrainingResultsOutput(ctx *pulumi.Context, args GetPredictionTrainingResultsOutputArgs, opts ...pulumi.InvokeOption) GetPredictionTrainingResultsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetPredictionTrainingResultsResult, error) {
			args := v.(GetPredictionTrainingResultsArgs)
			r, err := GetPredictionTrainingResults(ctx, &args, opts...)
			var s GetPredictionTrainingResultsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetPredictionTrainingResultsResultOutput)
}

type GetPredictionTrainingResultsOutputArgs struct {
	HubName           pulumi.StringInput `pulumi:"hubName"`
	PredictionName    pulumi.StringInput `pulumi:"predictionName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetPredictionTrainingResultsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPredictionTrainingResultsArgs)(nil)).Elem()
}


type GetPredictionTrainingResultsResultOutput struct{ *pulumi.OutputState }

func (GetPredictionTrainingResultsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetPredictionTrainingResultsResult)(nil)).Elem()
}

func (o GetPredictionTrainingResultsResultOutput) ToGetPredictionTrainingResultsResultOutput() GetPredictionTrainingResultsResultOutput {
	return o
}

func (o GetPredictionTrainingResultsResultOutput) ToGetPredictionTrainingResultsResultOutputWithContext(ctx context.Context) GetPredictionTrainingResultsResultOutput {
	return o
}

func (o GetPredictionTrainingResultsResultOutput) CanonicalProfiles() CanonicalProfileDefinitionResponseArrayOutput {
	return o.ApplyT(func(v GetPredictionTrainingResultsResult) []CanonicalProfileDefinitionResponse {
		return v.CanonicalProfiles
	}).(CanonicalProfileDefinitionResponseArrayOutput)
}

func (o GetPredictionTrainingResultsResultOutput) PredictionDistribution() PredictionDistributionDefinitionResponseOutput {
	return o.ApplyT(func(v GetPredictionTrainingResultsResult) PredictionDistributionDefinitionResponse {
		return v.PredictionDistribution
	}).(PredictionDistributionDefinitionResponseOutput)
}

func (o GetPredictionTrainingResultsResultOutput) PrimaryProfileInstanceCount() pulumi.Float64Output {
	return o.ApplyT(func(v GetPredictionTrainingResultsResult) float64 { return v.PrimaryProfileInstanceCount }).(pulumi.Float64Output)
}

func (o GetPredictionTrainingResultsResultOutput) ScoreName() pulumi.StringOutput {
	return o.ApplyT(func(v GetPredictionTrainingResultsResult) string { return v.ScoreName }).(pulumi.StringOutput)
}

func (o GetPredictionTrainingResultsResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetPredictionTrainingResultsResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetPredictionTrainingResultsResultOutput{})
}
