


package v20171111preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssessment(ctx *pulumi.Context, args *LookupAssessmentArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentResult, error) {
	var rv LookupAssessmentResult
	err := ctx.Invoke("azure-native:migrate/v20171111preview:getAssessment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAssessmentArgs struct {
	AssessmentName    string `pulumi:"assessmentName"`
	GroupName         string `pulumi:"groupName"`
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupAssessmentResult struct {
	AzureHybridUseBenefit  string  `pulumi:"azureHybridUseBenefit"`
	AzureLocation          string  `pulumi:"azureLocation"`
	AzureOfferCode         string  `pulumi:"azureOfferCode"`
	AzurePricingTier       string  `pulumi:"azurePricingTier"`
	AzureStorageRedundancy string  `pulumi:"azureStorageRedundancy"`
	CreatedTimestamp       string  `pulumi:"createdTimestamp"`
	Currency               string  `pulumi:"currency"`
	DiscountPercentage     float64 `pulumi:"discountPercentage"`
	ETag                   *string `pulumi:"eTag"`
	Id                     string  `pulumi:"id"`
	MonthlyBandwidthCost   float64 `pulumi:"monthlyBandwidthCost"`
	MonthlyComputeCost     float64 `pulumi:"monthlyComputeCost"`
	MonthlyStorageCost     float64 `pulumi:"monthlyStorageCost"`
	Name                   string  `pulumi:"name"`
	NumberOfMachines       int     `pulumi:"numberOfMachines"`
	Percentile             string  `pulumi:"percentile"`
	PricesTimestamp        string  `pulumi:"pricesTimestamp"`
	ScalingFactor          float64 `pulumi:"scalingFactor"`
	Stage                  string  `pulumi:"stage"`
	Status                 string  `pulumi:"status"`
	TimeRange              string  `pulumi:"timeRange"`
	Type                   string  `pulumi:"type"`
	UpdatedTimestamp       string  `pulumi:"updatedTimestamp"`
}
