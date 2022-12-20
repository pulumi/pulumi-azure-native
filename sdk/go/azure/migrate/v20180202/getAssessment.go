


package v20180202

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAssessment(ctx *pulumi.Context, args *LookupAssessmentArgs, opts ...pulumi.InvokeOption) (*LookupAssessmentResult, error) {
	var rv LookupAssessmentResult
	err := ctx.Invoke("azure-native:migrate/v20180202:getAssessment", args, &rv, opts...)
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
	AzureHybridUseBenefit        string  `pulumi:"azureHybridUseBenefit"`
	AzureLocation                string  `pulumi:"azureLocation"`
	AzureOfferCode               string  `pulumi:"azureOfferCode"`
	AzurePricingTier             string  `pulumi:"azurePricingTier"`
	AzureStorageRedundancy       string  `pulumi:"azureStorageRedundancy"`
	ConfidenceRatingInPercentage float64 `pulumi:"confidenceRatingInPercentage"`
	CreatedTimestamp             string  `pulumi:"createdTimestamp"`
	Currency                     string  `pulumi:"currency"`
	DiscountPercentage           float64 `pulumi:"discountPercentage"`
	ETag                         *string `pulumi:"eTag"`
	Id                           string  `pulumi:"id"`
	MonthlyBandwidthCost         float64 `pulumi:"monthlyBandwidthCost"`
	MonthlyComputeCost           float64 `pulumi:"monthlyComputeCost"`
	MonthlyStorageCost           float64 `pulumi:"monthlyStorageCost"`
	Name                         string  `pulumi:"name"`
	NumberOfMachines             int     `pulumi:"numberOfMachines"`
	Percentile                   string  `pulumi:"percentile"`
	PricesTimestamp              string  `pulumi:"pricesTimestamp"`
	ScalingFactor                float64 `pulumi:"scalingFactor"`
	SizingCriterion              string  `pulumi:"sizingCriterion"`
	Stage                        string  `pulumi:"stage"`
	Status                       string  `pulumi:"status"`
	TimeRange                    string  `pulumi:"timeRange"`
	Type                         string  `pulumi:"type"`
	UpdatedTimestamp             string  `pulumi:"updatedTimestamp"`
}

func LookupAssessmentOutput(ctx *pulumi.Context, args LookupAssessmentOutputArgs, opts ...pulumi.InvokeOption) LookupAssessmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAssessmentResult, error) {
			args := v.(LookupAssessmentArgs)
			r, err := LookupAssessment(ctx, &args, opts...)
			var s LookupAssessmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAssessmentResultOutput)
}

type LookupAssessmentOutputArgs struct {
	AssessmentName    pulumi.StringInput `pulumi:"assessmentName"`
	GroupName         pulumi.StringInput `pulumi:"groupName"`
	ProjectName       pulumi.StringInput `pulumi:"projectName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAssessmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssessmentArgs)(nil)).Elem()
}


type LookupAssessmentResultOutput struct{ *pulumi.OutputState }

func (LookupAssessmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAssessmentResult)(nil)).Elem()
}

func (o LookupAssessmentResultOutput) ToLookupAssessmentResultOutput() LookupAssessmentResultOutput {
	return o
}

func (o LookupAssessmentResultOutput) ToLookupAssessmentResultOutputWithContext(ctx context.Context) LookupAssessmentResultOutput {
	return o
}

func (o LookupAssessmentResultOutput) AzureHybridUseBenefit() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.AzureHybridUseBenefit }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) AzureLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.AzureLocation }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) AzureOfferCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.AzureOfferCode }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) AzurePricingTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.AzurePricingTier }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) AzureStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.AzureStorageRedundancy }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) ConfidenceRatingInPercentage() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAssessmentResult) float64 { return v.ConfidenceRatingInPercentage }).(pulumi.Float64Output)
}

func (o LookupAssessmentResultOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Currency }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) DiscountPercentage() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAssessmentResult) float64 { return v.DiscountPercentage }).(pulumi.Float64Output)
}

func (o LookupAssessmentResultOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAssessmentResult) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o LookupAssessmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) MonthlyBandwidthCost() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAssessmentResult) float64 { return v.MonthlyBandwidthCost }).(pulumi.Float64Output)
}

func (o LookupAssessmentResultOutput) MonthlyComputeCost() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAssessmentResult) float64 { return v.MonthlyComputeCost }).(pulumi.Float64Output)
}

func (o LookupAssessmentResultOutput) MonthlyStorageCost() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAssessmentResult) float64 { return v.MonthlyStorageCost }).(pulumi.Float64Output)
}

func (o LookupAssessmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) NumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAssessmentResult) int { return v.NumberOfMachines }).(pulumi.IntOutput)
}

func (o LookupAssessmentResultOutput) Percentile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Percentile }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) PricesTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.PricesTimestamp }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) ScalingFactor() pulumi.Float64Output {
	return o.ApplyT(func(v LookupAssessmentResult) float64 { return v.ScalingFactor }).(pulumi.Float64Output)
}

func (o LookupAssessmentResultOutput) SizingCriterion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.SizingCriterion }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) Stage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Stage }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) TimeRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.TimeRange }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupAssessmentResultOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAssessmentResult) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAssessmentResultOutput{})
}
