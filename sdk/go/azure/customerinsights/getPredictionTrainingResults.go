


package customerinsights

import (
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
