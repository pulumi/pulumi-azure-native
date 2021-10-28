


package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssessmentProperties struct {
	AzureDiskType          string   `pulumi:"azureDiskType"`
	AzureHybridUseBenefit  string   `pulumi:"azureHybridUseBenefit"`
	AzureLocation          string   `pulumi:"azureLocation"`
	AzureOfferCode         string   `pulumi:"azureOfferCode"`
	AzurePricingTier       string   `pulumi:"azurePricingTier"`
	AzureStorageRedundancy string   `pulumi:"azureStorageRedundancy"`
	AzureVmFamilies        []string `pulumi:"azureVmFamilies"`
	Currency               string   `pulumi:"currency"`
	DiscountPercentage     float64  `pulumi:"discountPercentage"`
	Percentile             string   `pulumi:"percentile"`
	ReservedInstance       string   `pulumi:"reservedInstance"`
	ScalingFactor          float64  `pulumi:"scalingFactor"`
	SizingCriterion        string   `pulumi:"sizingCriterion"`
	Stage                  string   `pulumi:"stage"`
	TimeRange              string   `pulumi:"timeRange"`
	VmUptime               VmUptime `pulumi:"vmUptime"`
}





type AssessmentPropertiesInput interface {
	pulumi.Input

	ToAssessmentPropertiesOutput() AssessmentPropertiesOutput
	ToAssessmentPropertiesOutputWithContext(context.Context) AssessmentPropertiesOutput
}

type AssessmentPropertiesArgs struct {
	AzureDiskType          pulumi.StringInput      `pulumi:"azureDiskType"`
	AzureHybridUseBenefit  pulumi.StringInput      `pulumi:"azureHybridUseBenefit"`
	AzureLocation          pulumi.StringInput      `pulumi:"azureLocation"`
	AzureOfferCode         pulumi.StringInput      `pulumi:"azureOfferCode"`
	AzurePricingTier       pulumi.StringInput      `pulumi:"azurePricingTier"`
	AzureStorageRedundancy pulumi.StringInput      `pulumi:"azureStorageRedundancy"`
	AzureVmFamilies        pulumi.StringArrayInput `pulumi:"azureVmFamilies"`
	Currency               pulumi.StringInput      `pulumi:"currency"`
	DiscountPercentage     pulumi.Float64Input     `pulumi:"discountPercentage"`
	Percentile             pulumi.StringInput      `pulumi:"percentile"`
	ReservedInstance       pulumi.StringInput      `pulumi:"reservedInstance"`
	ScalingFactor          pulumi.Float64Input     `pulumi:"scalingFactor"`
	SizingCriterion        pulumi.StringInput      `pulumi:"sizingCriterion"`
	Stage                  pulumi.StringInput      `pulumi:"stage"`
	TimeRange              pulumi.StringInput      `pulumi:"timeRange"`
	VmUptime               VmUptimeInput           `pulumi:"vmUptime"`
}

func (AssessmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentProperties)(nil)).Elem()
}

func (i AssessmentPropertiesArgs) ToAssessmentPropertiesOutput() AssessmentPropertiesOutput {
	return i.ToAssessmentPropertiesOutputWithContext(context.Background())
}

func (i AssessmentPropertiesArgs) ToAssessmentPropertiesOutputWithContext(ctx context.Context) AssessmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentPropertiesOutput)
}

func (i AssessmentPropertiesArgs) ToAssessmentPropertiesPtrOutput() AssessmentPropertiesPtrOutput {
	return i.ToAssessmentPropertiesPtrOutputWithContext(context.Background())
}

func (i AssessmentPropertiesArgs) ToAssessmentPropertiesPtrOutputWithContext(ctx context.Context) AssessmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentPropertiesOutput).ToAssessmentPropertiesPtrOutputWithContext(ctx)
}









type AssessmentPropertiesPtrInput interface {
	pulumi.Input

	ToAssessmentPropertiesPtrOutput() AssessmentPropertiesPtrOutput
	ToAssessmentPropertiesPtrOutputWithContext(context.Context) AssessmentPropertiesPtrOutput
}

type assessmentPropertiesPtrType AssessmentPropertiesArgs

func AssessmentPropertiesPtr(v *AssessmentPropertiesArgs) AssessmentPropertiesPtrInput {
	return (*assessmentPropertiesPtrType)(v)
}

func (*assessmentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentProperties)(nil)).Elem()
}

func (i *assessmentPropertiesPtrType) ToAssessmentPropertiesPtrOutput() AssessmentPropertiesPtrOutput {
	return i.ToAssessmentPropertiesPtrOutputWithContext(context.Background())
}

func (i *assessmentPropertiesPtrType) ToAssessmentPropertiesPtrOutputWithContext(ctx context.Context) AssessmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentPropertiesPtrOutput)
}

type AssessmentPropertiesOutput struct{ *pulumi.OutputState }

func (AssessmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentProperties)(nil)).Elem()
}

func (o AssessmentPropertiesOutput) ToAssessmentPropertiesOutput() AssessmentPropertiesOutput {
	return o
}

func (o AssessmentPropertiesOutput) ToAssessmentPropertiesOutputWithContext(ctx context.Context) AssessmentPropertiesOutput {
	return o
}

func (o AssessmentPropertiesOutput) ToAssessmentPropertiesPtrOutput() AssessmentPropertiesPtrOutput {
	return o.ToAssessmentPropertiesPtrOutputWithContext(context.Background())
}

func (o AssessmentPropertiesOutput) ToAssessmentPropertiesPtrOutputWithContext(ctx context.Context) AssessmentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentProperties) *AssessmentProperties {
		return &v
	}).(AssessmentPropertiesPtrOutput)
}

func (o AssessmentPropertiesOutput) AzureDiskType() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.AzureDiskType }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) AzureHybridUseBenefit() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.AzureHybridUseBenefit }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) AzureLocation() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.AzureLocation }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) AzureOfferCode() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.AzureOfferCode }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) AzurePricingTier() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.AzurePricingTier }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) AzureStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.AzureStorageRedundancy }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) AzureVmFamilies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AssessmentProperties) []string { return v.AzureVmFamilies }).(pulumi.StringArrayOutput)
}

func (o AssessmentPropertiesOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.Currency }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) DiscountPercentage() pulumi.Float64Output {
	return o.ApplyT(func(v AssessmentProperties) float64 { return v.DiscountPercentage }).(pulumi.Float64Output)
}

func (o AssessmentPropertiesOutput) Percentile() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.Percentile }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) ReservedInstance() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.ReservedInstance }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) ScalingFactor() pulumi.Float64Output {
	return o.ApplyT(func(v AssessmentProperties) float64 { return v.ScalingFactor }).(pulumi.Float64Output)
}

func (o AssessmentPropertiesOutput) SizingCriterion() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.SizingCriterion }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) Stage() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.Stage }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) TimeRange() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentProperties) string { return v.TimeRange }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesOutput) VmUptime() VmUptimeOutput {
	return o.ApplyT(func(v AssessmentProperties) VmUptime { return v.VmUptime }).(VmUptimeOutput)
}

type AssessmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AssessmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentProperties)(nil)).Elem()
}

func (o AssessmentPropertiesPtrOutput) ToAssessmentPropertiesPtrOutput() AssessmentPropertiesPtrOutput {
	return o
}

func (o AssessmentPropertiesPtrOutput) ToAssessmentPropertiesPtrOutputWithContext(ctx context.Context) AssessmentPropertiesPtrOutput {
	return o
}

func (o AssessmentPropertiesPtrOutput) Elem() AssessmentPropertiesOutput {
	return o.ApplyT(func(v *AssessmentProperties) AssessmentProperties {
		if v != nil {
			return *v
		}
		var ret AssessmentProperties
		return ret
	}).(AssessmentPropertiesOutput)
}

func (o AssessmentPropertiesPtrOutput) AzureDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AzureDiskType
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) AzureHybridUseBenefit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AzureHybridUseBenefit
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) AzureLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AzureLocation
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) AzureOfferCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AzureOfferCode
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) AzurePricingTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AzurePricingTier
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) AzureStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AzureStorageRedundancy
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) AzureVmFamilies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssessmentProperties) []string {
		if v == nil {
			return nil
		}
		return v.AzureVmFamilies
	}).(pulumi.StringArrayOutput)
}

func (o AssessmentPropertiesPtrOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Currency
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) DiscountPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *float64 {
		if v == nil {
			return nil
		}
		return &v.DiscountPercentage
	}).(pulumi.Float64PtrOutput)
}

func (o AssessmentPropertiesPtrOutput) Percentile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Percentile
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) ReservedInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ReservedInstance
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) ScalingFactor() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *float64 {
		if v == nil {
			return nil
		}
		return &v.ScalingFactor
	}).(pulumi.Float64PtrOutput)
}

func (o AssessmentPropertiesPtrOutput) SizingCriterion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.SizingCriterion
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) Stage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Stage
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) TimeRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *string {
		if v == nil {
			return nil
		}
		return &v.TimeRange
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesPtrOutput) VmUptime() VmUptimePtrOutput {
	return o.ApplyT(func(v *AssessmentProperties) *VmUptime {
		if v == nil {
			return nil
		}
		return &v.VmUptime
	}).(VmUptimePtrOutput)
}

type AssessmentPropertiesResponse struct {
	AzureDiskType                 string           `pulumi:"azureDiskType"`
	AzureHybridUseBenefit         string           `pulumi:"azureHybridUseBenefit"`
	AzureLocation                 string           `pulumi:"azureLocation"`
	AzureOfferCode                string           `pulumi:"azureOfferCode"`
	AzurePricingTier              string           `pulumi:"azurePricingTier"`
	AzureStorageRedundancy        string           `pulumi:"azureStorageRedundancy"`
	AzureVmFamilies               []string         `pulumi:"azureVmFamilies"`
	ConfidenceRatingInPercentage  float64          `pulumi:"confidenceRatingInPercentage"`
	CreatedTimestamp              string           `pulumi:"createdTimestamp"`
	Currency                      string           `pulumi:"currency"`
	DiscountPercentage            float64          `pulumi:"discountPercentage"`
	EaSubscriptionId              string           `pulumi:"eaSubscriptionId"`
	MonthlyBandwidthCost          float64          `pulumi:"monthlyBandwidthCost"`
	MonthlyComputeCost            float64          `pulumi:"monthlyComputeCost"`
	MonthlyPremiumStorageCost     float64          `pulumi:"monthlyPremiumStorageCost"`
	MonthlyStandardSSDStorageCost float64          `pulumi:"monthlyStandardSSDStorageCost"`
	MonthlyStorageCost            float64          `pulumi:"monthlyStorageCost"`
	NumberOfMachines              int              `pulumi:"numberOfMachines"`
	Percentile                    string           `pulumi:"percentile"`
	PerfDataEndTime               string           `pulumi:"perfDataEndTime"`
	PerfDataStartTime             string           `pulumi:"perfDataStartTime"`
	PricesTimestamp               string           `pulumi:"pricesTimestamp"`
	ReservedInstance              string           `pulumi:"reservedInstance"`
	ScalingFactor                 float64          `pulumi:"scalingFactor"`
	SizingCriterion               string           `pulumi:"sizingCriterion"`
	Stage                         string           `pulumi:"stage"`
	Status                        string           `pulumi:"status"`
	TimeRange                     string           `pulumi:"timeRange"`
	UpdatedTimestamp              string           `pulumi:"updatedTimestamp"`
	VmUptime                      VmUptimeResponse `pulumi:"vmUptime"`
}





type AssessmentPropertiesResponseInput interface {
	pulumi.Input

	ToAssessmentPropertiesResponseOutput() AssessmentPropertiesResponseOutput
	ToAssessmentPropertiesResponseOutputWithContext(context.Context) AssessmentPropertiesResponseOutput
}

type AssessmentPropertiesResponseArgs struct {
	AzureDiskType                 pulumi.StringInput      `pulumi:"azureDiskType"`
	AzureHybridUseBenefit         pulumi.StringInput      `pulumi:"azureHybridUseBenefit"`
	AzureLocation                 pulumi.StringInput      `pulumi:"azureLocation"`
	AzureOfferCode                pulumi.StringInput      `pulumi:"azureOfferCode"`
	AzurePricingTier              pulumi.StringInput      `pulumi:"azurePricingTier"`
	AzureStorageRedundancy        pulumi.StringInput      `pulumi:"azureStorageRedundancy"`
	AzureVmFamilies               pulumi.StringArrayInput `pulumi:"azureVmFamilies"`
	ConfidenceRatingInPercentage  pulumi.Float64Input     `pulumi:"confidenceRatingInPercentage"`
	CreatedTimestamp              pulumi.StringInput      `pulumi:"createdTimestamp"`
	Currency                      pulumi.StringInput      `pulumi:"currency"`
	DiscountPercentage            pulumi.Float64Input     `pulumi:"discountPercentage"`
	EaSubscriptionId              pulumi.StringInput      `pulumi:"eaSubscriptionId"`
	MonthlyBandwidthCost          pulumi.Float64Input     `pulumi:"monthlyBandwidthCost"`
	MonthlyComputeCost            pulumi.Float64Input     `pulumi:"monthlyComputeCost"`
	MonthlyPremiumStorageCost     pulumi.Float64Input     `pulumi:"monthlyPremiumStorageCost"`
	MonthlyStandardSSDStorageCost pulumi.Float64Input     `pulumi:"monthlyStandardSSDStorageCost"`
	MonthlyStorageCost            pulumi.Float64Input     `pulumi:"monthlyStorageCost"`
	NumberOfMachines              pulumi.IntInput         `pulumi:"numberOfMachines"`
	Percentile                    pulumi.StringInput      `pulumi:"percentile"`
	PerfDataEndTime               pulumi.StringInput      `pulumi:"perfDataEndTime"`
	PerfDataStartTime             pulumi.StringInput      `pulumi:"perfDataStartTime"`
	PricesTimestamp               pulumi.StringInput      `pulumi:"pricesTimestamp"`
	ReservedInstance              pulumi.StringInput      `pulumi:"reservedInstance"`
	ScalingFactor                 pulumi.Float64Input     `pulumi:"scalingFactor"`
	SizingCriterion               pulumi.StringInput      `pulumi:"sizingCriterion"`
	Stage                         pulumi.StringInput      `pulumi:"stage"`
	Status                        pulumi.StringInput      `pulumi:"status"`
	TimeRange                     pulumi.StringInput      `pulumi:"timeRange"`
	UpdatedTimestamp              pulumi.StringInput      `pulumi:"updatedTimestamp"`
	VmUptime                      VmUptimeResponseInput   `pulumi:"vmUptime"`
}

func (AssessmentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentPropertiesResponse)(nil)).Elem()
}

func (i AssessmentPropertiesResponseArgs) ToAssessmentPropertiesResponseOutput() AssessmentPropertiesResponseOutput {
	return i.ToAssessmentPropertiesResponseOutputWithContext(context.Background())
}

func (i AssessmentPropertiesResponseArgs) ToAssessmentPropertiesResponseOutputWithContext(ctx context.Context) AssessmentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentPropertiesResponseOutput)
}

func (i AssessmentPropertiesResponseArgs) ToAssessmentPropertiesResponsePtrOutput() AssessmentPropertiesResponsePtrOutput {
	return i.ToAssessmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AssessmentPropertiesResponseArgs) ToAssessmentPropertiesResponsePtrOutputWithContext(ctx context.Context) AssessmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentPropertiesResponseOutput).ToAssessmentPropertiesResponsePtrOutputWithContext(ctx)
}









type AssessmentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToAssessmentPropertiesResponsePtrOutput() AssessmentPropertiesResponsePtrOutput
	ToAssessmentPropertiesResponsePtrOutputWithContext(context.Context) AssessmentPropertiesResponsePtrOutput
}

type assessmentPropertiesResponsePtrType AssessmentPropertiesResponseArgs

func AssessmentPropertiesResponsePtr(v *AssessmentPropertiesResponseArgs) AssessmentPropertiesResponsePtrInput {
	return (*assessmentPropertiesResponsePtrType)(v)
}

func (*assessmentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentPropertiesResponse)(nil)).Elem()
}

func (i *assessmentPropertiesResponsePtrType) ToAssessmentPropertiesResponsePtrOutput() AssessmentPropertiesResponsePtrOutput {
	return i.ToAssessmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *assessmentPropertiesResponsePtrType) ToAssessmentPropertiesResponsePtrOutputWithContext(ctx context.Context) AssessmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentPropertiesResponsePtrOutput)
}

type AssessmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AssessmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentPropertiesResponse)(nil)).Elem()
}

func (o AssessmentPropertiesResponseOutput) ToAssessmentPropertiesResponseOutput() AssessmentPropertiesResponseOutput {
	return o
}

func (o AssessmentPropertiesResponseOutput) ToAssessmentPropertiesResponseOutputWithContext(ctx context.Context) AssessmentPropertiesResponseOutput {
	return o
}

func (o AssessmentPropertiesResponseOutput) ToAssessmentPropertiesResponsePtrOutput() AssessmentPropertiesResponsePtrOutput {
	return o.ToAssessmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AssessmentPropertiesResponseOutput) ToAssessmentPropertiesResponsePtrOutputWithContext(ctx context.Context) AssessmentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentPropertiesResponse) *AssessmentPropertiesResponse {
		return &v
	}).(AssessmentPropertiesResponsePtrOutput)
}

func (o AssessmentPropertiesResponseOutput) AzureDiskType() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.AzureDiskType }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) AzureHybridUseBenefit() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.AzureHybridUseBenefit }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) AzureLocation() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.AzureLocation }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) AzureOfferCode() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.AzureOfferCode }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) AzurePricingTier() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.AzurePricingTier }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) AzureStorageRedundancy() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.AzureStorageRedundancy }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) AzureVmFamilies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) []string { return v.AzureVmFamilies }).(pulumi.StringArrayOutput)
}

func (o AssessmentPropertiesResponseOutput) ConfidenceRatingInPercentage() pulumi.Float64Output {
	return o.ApplyT(func(v AssessmentPropertiesResponse) float64 { return v.ConfidenceRatingInPercentage }).(pulumi.Float64Output)
}

func (o AssessmentPropertiesResponseOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) Currency() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.Currency }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) DiscountPercentage() pulumi.Float64Output {
	return o.ApplyT(func(v AssessmentPropertiesResponse) float64 { return v.DiscountPercentage }).(pulumi.Float64Output)
}

func (o AssessmentPropertiesResponseOutput) EaSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.EaSubscriptionId }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) MonthlyBandwidthCost() pulumi.Float64Output {
	return o.ApplyT(func(v AssessmentPropertiesResponse) float64 { return v.MonthlyBandwidthCost }).(pulumi.Float64Output)
}

func (o AssessmentPropertiesResponseOutput) MonthlyComputeCost() pulumi.Float64Output {
	return o.ApplyT(func(v AssessmentPropertiesResponse) float64 { return v.MonthlyComputeCost }).(pulumi.Float64Output)
}

func (o AssessmentPropertiesResponseOutput) MonthlyPremiumStorageCost() pulumi.Float64Output {
	return o.ApplyT(func(v AssessmentPropertiesResponse) float64 { return v.MonthlyPremiumStorageCost }).(pulumi.Float64Output)
}

func (o AssessmentPropertiesResponseOutput) MonthlyStandardSSDStorageCost() pulumi.Float64Output {
	return o.ApplyT(func(v AssessmentPropertiesResponse) float64 { return v.MonthlyStandardSSDStorageCost }).(pulumi.Float64Output)
}

func (o AssessmentPropertiesResponseOutput) MonthlyStorageCost() pulumi.Float64Output {
	return o.ApplyT(func(v AssessmentPropertiesResponse) float64 { return v.MonthlyStorageCost }).(pulumi.Float64Output)
}

func (o AssessmentPropertiesResponseOutput) NumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) int { return v.NumberOfMachines }).(pulumi.IntOutput)
}

func (o AssessmentPropertiesResponseOutput) Percentile() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.Percentile }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) PerfDataEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.PerfDataEndTime }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) PerfDataStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.PerfDataStartTime }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) PricesTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.PricesTimestamp }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) ReservedInstance() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.ReservedInstance }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) ScalingFactor() pulumi.Float64Output {
	return o.ApplyT(func(v AssessmentPropertiesResponse) float64 { return v.ScalingFactor }).(pulumi.Float64Output)
}

func (o AssessmentPropertiesResponseOutput) SizingCriterion() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.SizingCriterion }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) Stage() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.Stage }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) TimeRange() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.TimeRange }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

func (o AssessmentPropertiesResponseOutput) VmUptime() VmUptimeResponseOutput {
	return o.ApplyT(func(v AssessmentPropertiesResponse) VmUptimeResponse { return v.VmUptime }).(VmUptimeResponseOutput)
}

type AssessmentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AssessmentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentPropertiesResponse)(nil)).Elem()
}

func (o AssessmentPropertiesResponsePtrOutput) ToAssessmentPropertiesResponsePtrOutput() AssessmentPropertiesResponsePtrOutput {
	return o
}

func (o AssessmentPropertiesResponsePtrOutput) ToAssessmentPropertiesResponsePtrOutputWithContext(ctx context.Context) AssessmentPropertiesResponsePtrOutput {
	return o
}

func (o AssessmentPropertiesResponsePtrOutput) Elem() AssessmentPropertiesResponseOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) AssessmentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AssessmentPropertiesResponse
		return ret
	}).(AssessmentPropertiesResponseOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) AzureDiskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzureDiskType
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) AzureHybridUseBenefit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzureHybridUseBenefit
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) AzureLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzureLocation
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) AzureOfferCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzureOfferCode
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) AzurePricingTier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzurePricingTier
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) AzureStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AzureStorageRedundancy
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) AzureVmFamilies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.AzureVmFamilies
	}).(pulumi.StringArrayOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) ConfidenceRatingInPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.ConfidenceRatingInPercentage
	}).(pulumi.Float64PtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) CreatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Currency
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) DiscountPercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.DiscountPercentage
	}).(pulumi.Float64PtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) EaSubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EaSubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) MonthlyBandwidthCost() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MonthlyBandwidthCost
	}).(pulumi.Float64PtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) MonthlyComputeCost() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MonthlyComputeCost
	}).(pulumi.Float64PtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) MonthlyPremiumStorageCost() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MonthlyPremiumStorageCost
	}).(pulumi.Float64PtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) MonthlyStandardSSDStorageCost() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MonthlyStandardSSDStorageCost
	}).(pulumi.Float64PtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) MonthlyStorageCost() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.MonthlyStorageCost
	}).(pulumi.Float64PtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) NumberOfMachines() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.NumberOfMachines
	}).(pulumi.IntPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) Percentile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Percentile
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) PerfDataEndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PerfDataEndTime
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) PerfDataStartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PerfDataStartTime
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) PricesTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PricesTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) ReservedInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ReservedInstance
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) ScalingFactor() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.ScalingFactor
	}).(pulumi.Float64PtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) SizingCriterion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SizingCriterion
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) Stage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Stage
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) TimeRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TimeRange
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) UpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o AssessmentPropertiesResponsePtrOutput) VmUptime() VmUptimeResponsePtrOutput {
	return o.ApplyT(func(v *AssessmentPropertiesResponse) *VmUptimeResponse {
		if v == nil {
			return nil
		}
		return &v.VmUptime
	}).(VmUptimeResponsePtrOutput)
}

type CollectorAgentProperties struct {
	SpnDetails *CollectorBodyAgentSpnProperties `pulumi:"spnDetails"`
}





type CollectorAgentPropertiesInput interface {
	pulumi.Input

	ToCollectorAgentPropertiesOutput() CollectorAgentPropertiesOutput
	ToCollectorAgentPropertiesOutputWithContext(context.Context) CollectorAgentPropertiesOutput
}

type CollectorAgentPropertiesArgs struct {
	SpnDetails CollectorBodyAgentSpnPropertiesPtrInput `pulumi:"spnDetails"`
}

func (CollectorAgentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorAgentProperties)(nil)).Elem()
}

func (i CollectorAgentPropertiesArgs) ToCollectorAgentPropertiesOutput() CollectorAgentPropertiesOutput {
	return i.ToCollectorAgentPropertiesOutputWithContext(context.Background())
}

func (i CollectorAgentPropertiesArgs) ToCollectorAgentPropertiesOutputWithContext(ctx context.Context) CollectorAgentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorAgentPropertiesOutput)
}

func (i CollectorAgentPropertiesArgs) ToCollectorAgentPropertiesPtrOutput() CollectorAgentPropertiesPtrOutput {
	return i.ToCollectorAgentPropertiesPtrOutputWithContext(context.Background())
}

func (i CollectorAgentPropertiesArgs) ToCollectorAgentPropertiesPtrOutputWithContext(ctx context.Context) CollectorAgentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorAgentPropertiesOutput).ToCollectorAgentPropertiesPtrOutputWithContext(ctx)
}









type CollectorAgentPropertiesPtrInput interface {
	pulumi.Input

	ToCollectorAgentPropertiesPtrOutput() CollectorAgentPropertiesPtrOutput
	ToCollectorAgentPropertiesPtrOutputWithContext(context.Context) CollectorAgentPropertiesPtrOutput
}

type collectorAgentPropertiesPtrType CollectorAgentPropertiesArgs

func CollectorAgentPropertiesPtr(v *CollectorAgentPropertiesArgs) CollectorAgentPropertiesPtrInput {
	return (*collectorAgentPropertiesPtrType)(v)
}

func (*collectorAgentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorAgentProperties)(nil)).Elem()
}

func (i *collectorAgentPropertiesPtrType) ToCollectorAgentPropertiesPtrOutput() CollectorAgentPropertiesPtrOutput {
	return i.ToCollectorAgentPropertiesPtrOutputWithContext(context.Background())
}

func (i *collectorAgentPropertiesPtrType) ToCollectorAgentPropertiesPtrOutputWithContext(ctx context.Context) CollectorAgentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorAgentPropertiesPtrOutput)
}

type CollectorAgentPropertiesOutput struct{ *pulumi.OutputState }

func (CollectorAgentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorAgentProperties)(nil)).Elem()
}

func (o CollectorAgentPropertiesOutput) ToCollectorAgentPropertiesOutput() CollectorAgentPropertiesOutput {
	return o
}

func (o CollectorAgentPropertiesOutput) ToCollectorAgentPropertiesOutputWithContext(ctx context.Context) CollectorAgentPropertiesOutput {
	return o
}

func (o CollectorAgentPropertiesOutput) ToCollectorAgentPropertiesPtrOutput() CollectorAgentPropertiesPtrOutput {
	return o.ToCollectorAgentPropertiesPtrOutputWithContext(context.Background())
}

func (o CollectorAgentPropertiesOutput) ToCollectorAgentPropertiesPtrOutputWithContext(ctx context.Context) CollectorAgentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CollectorAgentProperties) *CollectorAgentProperties {
		return &v
	}).(CollectorAgentPropertiesPtrOutput)
}

func (o CollectorAgentPropertiesOutput) SpnDetails() CollectorBodyAgentSpnPropertiesPtrOutput {
	return o.ApplyT(func(v CollectorAgentProperties) *CollectorBodyAgentSpnProperties { return v.SpnDetails }).(CollectorBodyAgentSpnPropertiesPtrOutput)
}

type CollectorAgentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CollectorAgentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorAgentProperties)(nil)).Elem()
}

func (o CollectorAgentPropertiesPtrOutput) ToCollectorAgentPropertiesPtrOutput() CollectorAgentPropertiesPtrOutput {
	return o
}

func (o CollectorAgentPropertiesPtrOutput) ToCollectorAgentPropertiesPtrOutputWithContext(ctx context.Context) CollectorAgentPropertiesPtrOutput {
	return o
}

func (o CollectorAgentPropertiesPtrOutput) Elem() CollectorAgentPropertiesOutput {
	return o.ApplyT(func(v *CollectorAgentProperties) CollectorAgentProperties {
		if v != nil {
			return *v
		}
		var ret CollectorAgentProperties
		return ret
	}).(CollectorAgentPropertiesOutput)
}

func (o CollectorAgentPropertiesPtrOutput) SpnDetails() CollectorBodyAgentSpnPropertiesPtrOutput {
	return o.ApplyT(func(v *CollectorAgentProperties) *CollectorBodyAgentSpnProperties {
		if v == nil {
			return nil
		}
		return v.SpnDetails
	}).(CollectorBodyAgentSpnPropertiesPtrOutput)
}

type CollectorAgentPropertiesResponse struct {
	Id               string                                   `pulumi:"id"`
	LastHeartbeatUtc string                                   `pulumi:"lastHeartbeatUtc"`
	SpnDetails       *CollectorBodyAgentSpnPropertiesResponse `pulumi:"spnDetails"`
	Version          string                                   `pulumi:"version"`
}





type CollectorAgentPropertiesResponseInput interface {
	pulumi.Input

	ToCollectorAgentPropertiesResponseOutput() CollectorAgentPropertiesResponseOutput
	ToCollectorAgentPropertiesResponseOutputWithContext(context.Context) CollectorAgentPropertiesResponseOutput
}

type CollectorAgentPropertiesResponseArgs struct {
	Id               pulumi.StringInput                              `pulumi:"id"`
	LastHeartbeatUtc pulumi.StringInput                              `pulumi:"lastHeartbeatUtc"`
	SpnDetails       CollectorBodyAgentSpnPropertiesResponsePtrInput `pulumi:"spnDetails"`
	Version          pulumi.StringInput                              `pulumi:"version"`
}

func (CollectorAgentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorAgentPropertiesResponse)(nil)).Elem()
}

func (i CollectorAgentPropertiesResponseArgs) ToCollectorAgentPropertiesResponseOutput() CollectorAgentPropertiesResponseOutput {
	return i.ToCollectorAgentPropertiesResponseOutputWithContext(context.Background())
}

func (i CollectorAgentPropertiesResponseArgs) ToCollectorAgentPropertiesResponseOutputWithContext(ctx context.Context) CollectorAgentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorAgentPropertiesResponseOutput)
}

func (i CollectorAgentPropertiesResponseArgs) ToCollectorAgentPropertiesResponsePtrOutput() CollectorAgentPropertiesResponsePtrOutput {
	return i.ToCollectorAgentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CollectorAgentPropertiesResponseArgs) ToCollectorAgentPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorAgentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorAgentPropertiesResponseOutput).ToCollectorAgentPropertiesResponsePtrOutputWithContext(ctx)
}









type CollectorAgentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCollectorAgentPropertiesResponsePtrOutput() CollectorAgentPropertiesResponsePtrOutput
	ToCollectorAgentPropertiesResponsePtrOutputWithContext(context.Context) CollectorAgentPropertiesResponsePtrOutput
}

type collectorAgentPropertiesResponsePtrType CollectorAgentPropertiesResponseArgs

func CollectorAgentPropertiesResponsePtr(v *CollectorAgentPropertiesResponseArgs) CollectorAgentPropertiesResponsePtrInput {
	return (*collectorAgentPropertiesResponsePtrType)(v)
}

func (*collectorAgentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorAgentPropertiesResponse)(nil)).Elem()
}

func (i *collectorAgentPropertiesResponsePtrType) ToCollectorAgentPropertiesResponsePtrOutput() CollectorAgentPropertiesResponsePtrOutput {
	return i.ToCollectorAgentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *collectorAgentPropertiesResponsePtrType) ToCollectorAgentPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorAgentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorAgentPropertiesResponsePtrOutput)
}

type CollectorAgentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CollectorAgentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorAgentPropertiesResponse)(nil)).Elem()
}

func (o CollectorAgentPropertiesResponseOutput) ToCollectorAgentPropertiesResponseOutput() CollectorAgentPropertiesResponseOutput {
	return o
}

func (o CollectorAgentPropertiesResponseOutput) ToCollectorAgentPropertiesResponseOutputWithContext(ctx context.Context) CollectorAgentPropertiesResponseOutput {
	return o
}

func (o CollectorAgentPropertiesResponseOutput) ToCollectorAgentPropertiesResponsePtrOutput() CollectorAgentPropertiesResponsePtrOutput {
	return o.ToCollectorAgentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CollectorAgentPropertiesResponseOutput) ToCollectorAgentPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorAgentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CollectorAgentPropertiesResponse) *CollectorAgentPropertiesResponse {
		return &v
	}).(CollectorAgentPropertiesResponsePtrOutput)
}

func (o CollectorAgentPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CollectorAgentPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o CollectorAgentPropertiesResponseOutput) LastHeartbeatUtc() pulumi.StringOutput {
	return o.ApplyT(func(v CollectorAgentPropertiesResponse) string { return v.LastHeartbeatUtc }).(pulumi.StringOutput)
}

func (o CollectorAgentPropertiesResponseOutput) SpnDetails() CollectorBodyAgentSpnPropertiesResponsePtrOutput {
	return o.ApplyT(func(v CollectorAgentPropertiesResponse) *CollectorBodyAgentSpnPropertiesResponse { return v.SpnDetails }).(CollectorBodyAgentSpnPropertiesResponsePtrOutput)
}

func (o CollectorAgentPropertiesResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v CollectorAgentPropertiesResponse) string { return v.Version }).(pulumi.StringOutput)
}

type CollectorAgentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CollectorAgentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorAgentPropertiesResponse)(nil)).Elem()
}

func (o CollectorAgentPropertiesResponsePtrOutput) ToCollectorAgentPropertiesResponsePtrOutput() CollectorAgentPropertiesResponsePtrOutput {
	return o
}

func (o CollectorAgentPropertiesResponsePtrOutput) ToCollectorAgentPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorAgentPropertiesResponsePtrOutput {
	return o
}

func (o CollectorAgentPropertiesResponsePtrOutput) Elem() CollectorAgentPropertiesResponseOutput {
	return o.ApplyT(func(v *CollectorAgentPropertiesResponse) CollectorAgentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CollectorAgentPropertiesResponse
		return ret
	}).(CollectorAgentPropertiesResponseOutput)
}

func (o CollectorAgentPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorAgentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CollectorAgentPropertiesResponsePtrOutput) LastHeartbeatUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorAgentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastHeartbeatUtc
	}).(pulumi.StringPtrOutput)
}

func (o CollectorAgentPropertiesResponsePtrOutput) SpnDetails() CollectorBodyAgentSpnPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CollectorAgentPropertiesResponse) *CollectorBodyAgentSpnPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.SpnDetails
	}).(CollectorBodyAgentSpnPropertiesResponsePtrOutput)
}

func (o CollectorAgentPropertiesResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorAgentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type CollectorBodyAgentSpnProperties struct {
	ApplicationId *string `pulumi:"applicationId"`
	Audience      *string `pulumi:"audience"`
	Authority     *string `pulumi:"authority"`
	ObjectId      *string `pulumi:"objectId"`
	TenantId      *string `pulumi:"tenantId"`
}





type CollectorBodyAgentSpnPropertiesInput interface {
	pulumi.Input

	ToCollectorBodyAgentSpnPropertiesOutput() CollectorBodyAgentSpnPropertiesOutput
	ToCollectorBodyAgentSpnPropertiesOutputWithContext(context.Context) CollectorBodyAgentSpnPropertiesOutput
}

type CollectorBodyAgentSpnPropertiesArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	Audience      pulumi.StringPtrInput `pulumi:"audience"`
	Authority     pulumi.StringPtrInput `pulumi:"authority"`
	ObjectId      pulumi.StringPtrInput `pulumi:"objectId"`
	TenantId      pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (CollectorBodyAgentSpnPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorBodyAgentSpnProperties)(nil)).Elem()
}

func (i CollectorBodyAgentSpnPropertiesArgs) ToCollectorBodyAgentSpnPropertiesOutput() CollectorBodyAgentSpnPropertiesOutput {
	return i.ToCollectorBodyAgentSpnPropertiesOutputWithContext(context.Background())
}

func (i CollectorBodyAgentSpnPropertiesArgs) ToCollectorBodyAgentSpnPropertiesOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorBodyAgentSpnPropertiesOutput)
}

func (i CollectorBodyAgentSpnPropertiesArgs) ToCollectorBodyAgentSpnPropertiesPtrOutput() CollectorBodyAgentSpnPropertiesPtrOutput {
	return i.ToCollectorBodyAgentSpnPropertiesPtrOutputWithContext(context.Background())
}

func (i CollectorBodyAgentSpnPropertiesArgs) ToCollectorBodyAgentSpnPropertiesPtrOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorBodyAgentSpnPropertiesOutput).ToCollectorBodyAgentSpnPropertiesPtrOutputWithContext(ctx)
}









type CollectorBodyAgentSpnPropertiesPtrInput interface {
	pulumi.Input

	ToCollectorBodyAgentSpnPropertiesPtrOutput() CollectorBodyAgentSpnPropertiesPtrOutput
	ToCollectorBodyAgentSpnPropertiesPtrOutputWithContext(context.Context) CollectorBodyAgentSpnPropertiesPtrOutput
}

type collectorBodyAgentSpnPropertiesPtrType CollectorBodyAgentSpnPropertiesArgs

func CollectorBodyAgentSpnPropertiesPtr(v *CollectorBodyAgentSpnPropertiesArgs) CollectorBodyAgentSpnPropertiesPtrInput {
	return (*collectorBodyAgentSpnPropertiesPtrType)(v)
}

func (*collectorBodyAgentSpnPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorBodyAgentSpnProperties)(nil)).Elem()
}

func (i *collectorBodyAgentSpnPropertiesPtrType) ToCollectorBodyAgentSpnPropertiesPtrOutput() CollectorBodyAgentSpnPropertiesPtrOutput {
	return i.ToCollectorBodyAgentSpnPropertiesPtrOutputWithContext(context.Background())
}

func (i *collectorBodyAgentSpnPropertiesPtrType) ToCollectorBodyAgentSpnPropertiesPtrOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorBodyAgentSpnPropertiesPtrOutput)
}

type CollectorBodyAgentSpnPropertiesOutput struct{ *pulumi.OutputState }

func (CollectorBodyAgentSpnPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorBodyAgentSpnProperties)(nil)).Elem()
}

func (o CollectorBodyAgentSpnPropertiesOutput) ToCollectorBodyAgentSpnPropertiesOutput() CollectorBodyAgentSpnPropertiesOutput {
	return o
}

func (o CollectorBodyAgentSpnPropertiesOutput) ToCollectorBodyAgentSpnPropertiesOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesOutput {
	return o
}

func (o CollectorBodyAgentSpnPropertiesOutput) ToCollectorBodyAgentSpnPropertiesPtrOutput() CollectorBodyAgentSpnPropertiesPtrOutput {
	return o.ToCollectorBodyAgentSpnPropertiesPtrOutputWithContext(context.Background())
}

func (o CollectorBodyAgentSpnPropertiesOutput) ToCollectorBodyAgentSpnPropertiesPtrOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CollectorBodyAgentSpnProperties) *CollectorBodyAgentSpnProperties {
		return &v
	}).(CollectorBodyAgentSpnPropertiesPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorBodyAgentSpnProperties) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorBodyAgentSpnProperties) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorBodyAgentSpnProperties) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorBodyAgentSpnProperties) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorBodyAgentSpnProperties) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type CollectorBodyAgentSpnPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CollectorBodyAgentSpnPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorBodyAgentSpnProperties)(nil)).Elem()
}

func (o CollectorBodyAgentSpnPropertiesPtrOutput) ToCollectorBodyAgentSpnPropertiesPtrOutput() CollectorBodyAgentSpnPropertiesPtrOutput {
	return o
}

func (o CollectorBodyAgentSpnPropertiesPtrOutput) ToCollectorBodyAgentSpnPropertiesPtrOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesPtrOutput {
	return o
}

func (o CollectorBodyAgentSpnPropertiesPtrOutput) Elem() CollectorBodyAgentSpnPropertiesOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnProperties) CollectorBodyAgentSpnProperties {
		if v != nil {
			return *v
		}
		var ret CollectorBodyAgentSpnProperties
		return ret
	}).(CollectorBodyAgentSpnPropertiesOutput)
}

func (o CollectorBodyAgentSpnPropertiesPtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesPtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesPtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesPtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnProperties) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type CollectorBodyAgentSpnPropertiesResponse struct {
	ApplicationId *string `pulumi:"applicationId"`
	Audience      *string `pulumi:"audience"`
	Authority     *string `pulumi:"authority"`
	ObjectId      *string `pulumi:"objectId"`
	TenantId      *string `pulumi:"tenantId"`
}





type CollectorBodyAgentSpnPropertiesResponseInput interface {
	pulumi.Input

	ToCollectorBodyAgentSpnPropertiesResponseOutput() CollectorBodyAgentSpnPropertiesResponseOutput
	ToCollectorBodyAgentSpnPropertiesResponseOutputWithContext(context.Context) CollectorBodyAgentSpnPropertiesResponseOutput
}

type CollectorBodyAgentSpnPropertiesResponseArgs struct {
	ApplicationId pulumi.StringPtrInput `pulumi:"applicationId"`
	Audience      pulumi.StringPtrInput `pulumi:"audience"`
	Authority     pulumi.StringPtrInput `pulumi:"authority"`
	ObjectId      pulumi.StringPtrInput `pulumi:"objectId"`
	TenantId      pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (CollectorBodyAgentSpnPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorBodyAgentSpnPropertiesResponse)(nil)).Elem()
}

func (i CollectorBodyAgentSpnPropertiesResponseArgs) ToCollectorBodyAgentSpnPropertiesResponseOutput() CollectorBodyAgentSpnPropertiesResponseOutput {
	return i.ToCollectorBodyAgentSpnPropertiesResponseOutputWithContext(context.Background())
}

func (i CollectorBodyAgentSpnPropertiesResponseArgs) ToCollectorBodyAgentSpnPropertiesResponseOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorBodyAgentSpnPropertiesResponseOutput)
}

func (i CollectorBodyAgentSpnPropertiesResponseArgs) ToCollectorBodyAgentSpnPropertiesResponsePtrOutput() CollectorBodyAgentSpnPropertiesResponsePtrOutput {
	return i.ToCollectorBodyAgentSpnPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CollectorBodyAgentSpnPropertiesResponseArgs) ToCollectorBodyAgentSpnPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorBodyAgentSpnPropertiesResponseOutput).ToCollectorBodyAgentSpnPropertiesResponsePtrOutputWithContext(ctx)
}









type CollectorBodyAgentSpnPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCollectorBodyAgentSpnPropertiesResponsePtrOutput() CollectorBodyAgentSpnPropertiesResponsePtrOutput
	ToCollectorBodyAgentSpnPropertiesResponsePtrOutputWithContext(context.Context) CollectorBodyAgentSpnPropertiesResponsePtrOutput
}

type collectorBodyAgentSpnPropertiesResponsePtrType CollectorBodyAgentSpnPropertiesResponseArgs

func CollectorBodyAgentSpnPropertiesResponsePtr(v *CollectorBodyAgentSpnPropertiesResponseArgs) CollectorBodyAgentSpnPropertiesResponsePtrInput {
	return (*collectorBodyAgentSpnPropertiesResponsePtrType)(v)
}

func (*collectorBodyAgentSpnPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorBodyAgentSpnPropertiesResponse)(nil)).Elem()
}

func (i *collectorBodyAgentSpnPropertiesResponsePtrType) ToCollectorBodyAgentSpnPropertiesResponsePtrOutput() CollectorBodyAgentSpnPropertiesResponsePtrOutput {
	return i.ToCollectorBodyAgentSpnPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *collectorBodyAgentSpnPropertiesResponsePtrType) ToCollectorBodyAgentSpnPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorBodyAgentSpnPropertiesResponsePtrOutput)
}

type CollectorBodyAgentSpnPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CollectorBodyAgentSpnPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorBodyAgentSpnPropertiesResponse)(nil)).Elem()
}

func (o CollectorBodyAgentSpnPropertiesResponseOutput) ToCollectorBodyAgentSpnPropertiesResponseOutput() CollectorBodyAgentSpnPropertiesResponseOutput {
	return o
}

func (o CollectorBodyAgentSpnPropertiesResponseOutput) ToCollectorBodyAgentSpnPropertiesResponseOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesResponseOutput {
	return o
}

func (o CollectorBodyAgentSpnPropertiesResponseOutput) ToCollectorBodyAgentSpnPropertiesResponsePtrOutput() CollectorBodyAgentSpnPropertiesResponsePtrOutput {
	return o.ToCollectorBodyAgentSpnPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CollectorBodyAgentSpnPropertiesResponseOutput) ToCollectorBodyAgentSpnPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CollectorBodyAgentSpnPropertiesResponse) *CollectorBodyAgentSpnPropertiesResponse {
		return &v
	}).(CollectorBodyAgentSpnPropertiesResponsePtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesResponseOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorBodyAgentSpnPropertiesResponse) *string { return v.ApplicationId }).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesResponseOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorBodyAgentSpnPropertiesResponse) *string { return v.Audience }).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesResponseOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorBodyAgentSpnPropertiesResponse) *string { return v.Authority }).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorBodyAgentSpnPropertiesResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorBodyAgentSpnPropertiesResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type CollectorBodyAgentSpnPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CollectorBodyAgentSpnPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorBodyAgentSpnPropertiesResponse)(nil)).Elem()
}

func (o CollectorBodyAgentSpnPropertiesResponsePtrOutput) ToCollectorBodyAgentSpnPropertiesResponsePtrOutput() CollectorBodyAgentSpnPropertiesResponsePtrOutput {
	return o
}

func (o CollectorBodyAgentSpnPropertiesResponsePtrOutput) ToCollectorBodyAgentSpnPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorBodyAgentSpnPropertiesResponsePtrOutput {
	return o
}

func (o CollectorBodyAgentSpnPropertiesResponsePtrOutput) Elem() CollectorBodyAgentSpnPropertiesResponseOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnPropertiesResponse) CollectorBodyAgentSpnPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CollectorBodyAgentSpnPropertiesResponse
		return ret
	}).(CollectorBodyAgentSpnPropertiesResponseOutput)
}

func (o CollectorBodyAgentSpnPropertiesResponsePtrOutput) ApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesResponsePtrOutput) Audience() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Audience
	}).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesResponsePtrOutput) Authority() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Authority
	}).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o CollectorBodyAgentSpnPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorBodyAgentSpnPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type CollectorProperties struct {
	AgentProperties *CollectorAgentProperties `pulumi:"agentProperties"`
	DiscoverySiteId *string                   `pulumi:"discoverySiteId"`
}





type CollectorPropertiesInput interface {
	pulumi.Input

	ToCollectorPropertiesOutput() CollectorPropertiesOutput
	ToCollectorPropertiesOutputWithContext(context.Context) CollectorPropertiesOutput
}

type CollectorPropertiesArgs struct {
	AgentProperties CollectorAgentPropertiesPtrInput `pulumi:"agentProperties"`
	DiscoverySiteId pulumi.StringPtrInput            `pulumi:"discoverySiteId"`
}

func (CollectorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorProperties)(nil)).Elem()
}

func (i CollectorPropertiesArgs) ToCollectorPropertiesOutput() CollectorPropertiesOutput {
	return i.ToCollectorPropertiesOutputWithContext(context.Background())
}

func (i CollectorPropertiesArgs) ToCollectorPropertiesOutputWithContext(ctx context.Context) CollectorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorPropertiesOutput)
}

func (i CollectorPropertiesArgs) ToCollectorPropertiesPtrOutput() CollectorPropertiesPtrOutput {
	return i.ToCollectorPropertiesPtrOutputWithContext(context.Background())
}

func (i CollectorPropertiesArgs) ToCollectorPropertiesPtrOutputWithContext(ctx context.Context) CollectorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorPropertiesOutput).ToCollectorPropertiesPtrOutputWithContext(ctx)
}









type CollectorPropertiesPtrInput interface {
	pulumi.Input

	ToCollectorPropertiesPtrOutput() CollectorPropertiesPtrOutput
	ToCollectorPropertiesPtrOutputWithContext(context.Context) CollectorPropertiesPtrOutput
}

type collectorPropertiesPtrType CollectorPropertiesArgs

func CollectorPropertiesPtr(v *CollectorPropertiesArgs) CollectorPropertiesPtrInput {
	return (*collectorPropertiesPtrType)(v)
}

func (*collectorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorProperties)(nil)).Elem()
}

func (i *collectorPropertiesPtrType) ToCollectorPropertiesPtrOutput() CollectorPropertiesPtrOutput {
	return i.ToCollectorPropertiesPtrOutputWithContext(context.Background())
}

func (i *collectorPropertiesPtrType) ToCollectorPropertiesPtrOutputWithContext(ctx context.Context) CollectorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorPropertiesPtrOutput)
}

type CollectorPropertiesOutput struct{ *pulumi.OutputState }

func (CollectorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorProperties)(nil)).Elem()
}

func (o CollectorPropertiesOutput) ToCollectorPropertiesOutput() CollectorPropertiesOutput {
	return o
}

func (o CollectorPropertiesOutput) ToCollectorPropertiesOutputWithContext(ctx context.Context) CollectorPropertiesOutput {
	return o
}

func (o CollectorPropertiesOutput) ToCollectorPropertiesPtrOutput() CollectorPropertiesPtrOutput {
	return o.ToCollectorPropertiesPtrOutputWithContext(context.Background())
}

func (o CollectorPropertiesOutput) ToCollectorPropertiesPtrOutputWithContext(ctx context.Context) CollectorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CollectorProperties) *CollectorProperties {
		return &v
	}).(CollectorPropertiesPtrOutput)
}

func (o CollectorPropertiesOutput) AgentProperties() CollectorAgentPropertiesPtrOutput {
	return o.ApplyT(func(v CollectorProperties) *CollectorAgentProperties { return v.AgentProperties }).(CollectorAgentPropertiesPtrOutput)
}

func (o CollectorPropertiesOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorProperties) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

type CollectorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CollectorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorProperties)(nil)).Elem()
}

func (o CollectorPropertiesPtrOutput) ToCollectorPropertiesPtrOutput() CollectorPropertiesPtrOutput {
	return o
}

func (o CollectorPropertiesPtrOutput) ToCollectorPropertiesPtrOutputWithContext(ctx context.Context) CollectorPropertiesPtrOutput {
	return o
}

func (o CollectorPropertiesPtrOutput) Elem() CollectorPropertiesOutput {
	return o.ApplyT(func(v *CollectorProperties) CollectorProperties {
		if v != nil {
			return *v
		}
		var ret CollectorProperties
		return ret
	}).(CollectorPropertiesOutput)
}

func (o CollectorPropertiesPtrOutput) AgentProperties() CollectorAgentPropertiesPtrOutput {
	return o.ApplyT(func(v *CollectorProperties) *CollectorAgentProperties {
		if v == nil {
			return nil
		}
		return v.AgentProperties
	}).(CollectorAgentPropertiesPtrOutput)
}

func (o CollectorPropertiesPtrOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorProperties) *string {
		if v == nil {
			return nil
		}
		return v.DiscoverySiteId
	}).(pulumi.StringPtrOutput)
}

type CollectorPropertiesResponse struct {
	AgentProperties  *CollectorAgentPropertiesResponse `pulumi:"agentProperties"`
	CreatedTimestamp string                            `pulumi:"createdTimestamp"`
	DiscoverySiteId  *string                           `pulumi:"discoverySiteId"`
	UpdatedTimestamp string                            `pulumi:"updatedTimestamp"`
}





type CollectorPropertiesResponseInput interface {
	pulumi.Input

	ToCollectorPropertiesResponseOutput() CollectorPropertiesResponseOutput
	ToCollectorPropertiesResponseOutputWithContext(context.Context) CollectorPropertiesResponseOutput
}

type CollectorPropertiesResponseArgs struct {
	AgentProperties  CollectorAgentPropertiesResponsePtrInput `pulumi:"agentProperties"`
	CreatedTimestamp pulumi.StringInput                       `pulumi:"createdTimestamp"`
	DiscoverySiteId  pulumi.StringPtrInput                    `pulumi:"discoverySiteId"`
	UpdatedTimestamp pulumi.StringInput                       `pulumi:"updatedTimestamp"`
}

func (CollectorPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorPropertiesResponse)(nil)).Elem()
}

func (i CollectorPropertiesResponseArgs) ToCollectorPropertiesResponseOutput() CollectorPropertiesResponseOutput {
	return i.ToCollectorPropertiesResponseOutputWithContext(context.Background())
}

func (i CollectorPropertiesResponseArgs) ToCollectorPropertiesResponseOutputWithContext(ctx context.Context) CollectorPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorPropertiesResponseOutput)
}

func (i CollectorPropertiesResponseArgs) ToCollectorPropertiesResponsePtrOutput() CollectorPropertiesResponsePtrOutput {
	return i.ToCollectorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CollectorPropertiesResponseArgs) ToCollectorPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorPropertiesResponseOutput).ToCollectorPropertiesResponsePtrOutputWithContext(ctx)
}









type CollectorPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCollectorPropertiesResponsePtrOutput() CollectorPropertiesResponsePtrOutput
	ToCollectorPropertiesResponsePtrOutputWithContext(context.Context) CollectorPropertiesResponsePtrOutput
}

type collectorPropertiesResponsePtrType CollectorPropertiesResponseArgs

func CollectorPropertiesResponsePtr(v *CollectorPropertiesResponseArgs) CollectorPropertiesResponsePtrInput {
	return (*collectorPropertiesResponsePtrType)(v)
}

func (*collectorPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorPropertiesResponse)(nil)).Elem()
}

func (i *collectorPropertiesResponsePtrType) ToCollectorPropertiesResponsePtrOutput() CollectorPropertiesResponsePtrOutput {
	return i.ToCollectorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *collectorPropertiesResponsePtrType) ToCollectorPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CollectorPropertiesResponsePtrOutput)
}

type CollectorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CollectorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CollectorPropertiesResponse)(nil)).Elem()
}

func (o CollectorPropertiesResponseOutput) ToCollectorPropertiesResponseOutput() CollectorPropertiesResponseOutput {
	return o
}

func (o CollectorPropertiesResponseOutput) ToCollectorPropertiesResponseOutputWithContext(ctx context.Context) CollectorPropertiesResponseOutput {
	return o
}

func (o CollectorPropertiesResponseOutput) ToCollectorPropertiesResponsePtrOutput() CollectorPropertiesResponsePtrOutput {
	return o.ToCollectorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CollectorPropertiesResponseOutput) ToCollectorPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CollectorPropertiesResponse) *CollectorPropertiesResponse {
		return &v
	}).(CollectorPropertiesResponsePtrOutput)
}

func (o CollectorPropertiesResponseOutput) AgentProperties() CollectorAgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v CollectorPropertiesResponse) *CollectorAgentPropertiesResponse { return v.AgentProperties }).(CollectorAgentPropertiesResponsePtrOutput)
}

func (o CollectorPropertiesResponseOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v CollectorPropertiesResponse) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o CollectorPropertiesResponseOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CollectorPropertiesResponse) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

func (o CollectorPropertiesResponseOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v CollectorPropertiesResponse) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

type CollectorPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CollectorPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CollectorPropertiesResponse)(nil)).Elem()
}

func (o CollectorPropertiesResponsePtrOutput) ToCollectorPropertiesResponsePtrOutput() CollectorPropertiesResponsePtrOutput {
	return o
}

func (o CollectorPropertiesResponsePtrOutput) ToCollectorPropertiesResponsePtrOutputWithContext(ctx context.Context) CollectorPropertiesResponsePtrOutput {
	return o
}

func (o CollectorPropertiesResponsePtrOutput) Elem() CollectorPropertiesResponseOutput {
	return o.ApplyT(func(v *CollectorPropertiesResponse) CollectorPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CollectorPropertiesResponse
		return ret
	}).(CollectorPropertiesResponseOutput)
}

func (o CollectorPropertiesResponsePtrOutput) AgentProperties() CollectorAgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CollectorPropertiesResponse) *CollectorAgentPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.AgentProperties
	}).(CollectorAgentPropertiesResponsePtrOutput)
}

func (o CollectorPropertiesResponsePtrOutput) CreatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o CollectorPropertiesResponsePtrOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiscoverySiteId
	}).(pulumi.StringPtrOutput)
}

func (o CollectorPropertiesResponsePtrOutput) UpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CollectorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

type GroupProperties struct {
	GroupType *string `pulumi:"groupType"`
}





type GroupPropertiesInput interface {
	pulumi.Input

	ToGroupPropertiesOutput() GroupPropertiesOutput
	ToGroupPropertiesOutputWithContext(context.Context) GroupPropertiesOutput
}

type GroupPropertiesArgs struct {
	GroupType pulumi.StringPtrInput `pulumi:"groupType"`
}

func (GroupPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupProperties)(nil)).Elem()
}

func (i GroupPropertiesArgs) ToGroupPropertiesOutput() GroupPropertiesOutput {
	return i.ToGroupPropertiesOutputWithContext(context.Background())
}

func (i GroupPropertiesArgs) ToGroupPropertiesOutputWithContext(ctx context.Context) GroupPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPropertiesOutput)
}

func (i GroupPropertiesArgs) ToGroupPropertiesPtrOutput() GroupPropertiesPtrOutput {
	return i.ToGroupPropertiesPtrOutputWithContext(context.Background())
}

func (i GroupPropertiesArgs) ToGroupPropertiesPtrOutputWithContext(ctx context.Context) GroupPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPropertiesOutput).ToGroupPropertiesPtrOutputWithContext(ctx)
}









type GroupPropertiesPtrInput interface {
	pulumi.Input

	ToGroupPropertiesPtrOutput() GroupPropertiesPtrOutput
	ToGroupPropertiesPtrOutputWithContext(context.Context) GroupPropertiesPtrOutput
}

type groupPropertiesPtrType GroupPropertiesArgs

func GroupPropertiesPtr(v *GroupPropertiesArgs) GroupPropertiesPtrInput {
	return (*groupPropertiesPtrType)(v)
}

func (*groupPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupProperties)(nil)).Elem()
}

func (i *groupPropertiesPtrType) ToGroupPropertiesPtrOutput() GroupPropertiesPtrOutput {
	return i.ToGroupPropertiesPtrOutputWithContext(context.Background())
}

func (i *groupPropertiesPtrType) ToGroupPropertiesPtrOutputWithContext(ctx context.Context) GroupPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPropertiesPtrOutput)
}

type GroupPropertiesOutput struct{ *pulumi.OutputState }

func (GroupPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupProperties)(nil)).Elem()
}

func (o GroupPropertiesOutput) ToGroupPropertiesOutput() GroupPropertiesOutput {
	return o
}

func (o GroupPropertiesOutput) ToGroupPropertiesOutputWithContext(ctx context.Context) GroupPropertiesOutput {
	return o
}

func (o GroupPropertiesOutput) ToGroupPropertiesPtrOutput() GroupPropertiesPtrOutput {
	return o.ToGroupPropertiesPtrOutputWithContext(context.Background())
}

func (o GroupPropertiesOutput) ToGroupPropertiesPtrOutputWithContext(ctx context.Context) GroupPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupProperties) *GroupProperties {
		return &v
	}).(GroupPropertiesPtrOutput)
}

func (o GroupPropertiesOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupProperties) *string { return v.GroupType }).(pulumi.StringPtrOutput)
}

type GroupPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GroupPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupProperties)(nil)).Elem()
}

func (o GroupPropertiesPtrOutput) ToGroupPropertiesPtrOutput() GroupPropertiesPtrOutput {
	return o
}

func (o GroupPropertiesPtrOutput) ToGroupPropertiesPtrOutputWithContext(ctx context.Context) GroupPropertiesPtrOutput {
	return o
}

func (o GroupPropertiesPtrOutput) Elem() GroupPropertiesOutput {
	return o.ApplyT(func(v *GroupProperties) GroupProperties {
		if v != nil {
			return *v
		}
		var ret GroupProperties
		return ret
	}).(GroupPropertiesOutput)
}

func (o GroupPropertiesPtrOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupProperties) *string {
		if v == nil {
			return nil
		}
		return v.GroupType
	}).(pulumi.StringPtrOutput)
}

type GroupPropertiesResponse struct {
	AreAssessmentsRunning bool     `pulumi:"areAssessmentsRunning"`
	Assessments           []string `pulumi:"assessments"`
	CreatedTimestamp      string   `pulumi:"createdTimestamp"`
	GroupStatus           string   `pulumi:"groupStatus"`
	GroupType             *string  `pulumi:"groupType"`
	MachineCount          int      `pulumi:"machineCount"`
	UpdatedTimestamp      string   `pulumi:"updatedTimestamp"`
}





type GroupPropertiesResponseInput interface {
	pulumi.Input

	ToGroupPropertiesResponseOutput() GroupPropertiesResponseOutput
	ToGroupPropertiesResponseOutputWithContext(context.Context) GroupPropertiesResponseOutput
}

type GroupPropertiesResponseArgs struct {
	AreAssessmentsRunning pulumi.BoolInput        `pulumi:"areAssessmentsRunning"`
	Assessments           pulumi.StringArrayInput `pulumi:"assessments"`
	CreatedTimestamp      pulumi.StringInput      `pulumi:"createdTimestamp"`
	GroupStatus           pulumi.StringInput      `pulumi:"groupStatus"`
	GroupType             pulumi.StringPtrInput   `pulumi:"groupType"`
	MachineCount          pulumi.IntInput         `pulumi:"machineCount"`
	UpdatedTimestamp      pulumi.StringInput      `pulumi:"updatedTimestamp"`
}

func (GroupPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupPropertiesResponse)(nil)).Elem()
}

func (i GroupPropertiesResponseArgs) ToGroupPropertiesResponseOutput() GroupPropertiesResponseOutput {
	return i.ToGroupPropertiesResponseOutputWithContext(context.Background())
}

func (i GroupPropertiesResponseArgs) ToGroupPropertiesResponseOutputWithContext(ctx context.Context) GroupPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPropertiesResponseOutput)
}

func (i GroupPropertiesResponseArgs) ToGroupPropertiesResponsePtrOutput() GroupPropertiesResponsePtrOutput {
	return i.ToGroupPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i GroupPropertiesResponseArgs) ToGroupPropertiesResponsePtrOutputWithContext(ctx context.Context) GroupPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPropertiesResponseOutput).ToGroupPropertiesResponsePtrOutputWithContext(ctx)
}









type GroupPropertiesResponsePtrInput interface {
	pulumi.Input

	ToGroupPropertiesResponsePtrOutput() GroupPropertiesResponsePtrOutput
	ToGroupPropertiesResponsePtrOutputWithContext(context.Context) GroupPropertiesResponsePtrOutput
}

type groupPropertiesResponsePtrType GroupPropertiesResponseArgs

func GroupPropertiesResponsePtr(v *GroupPropertiesResponseArgs) GroupPropertiesResponsePtrInput {
	return (*groupPropertiesResponsePtrType)(v)
}

func (*groupPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPropertiesResponse)(nil)).Elem()
}

func (i *groupPropertiesResponsePtrType) ToGroupPropertiesResponsePtrOutput() GroupPropertiesResponsePtrOutput {
	return i.ToGroupPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *groupPropertiesResponsePtrType) ToGroupPropertiesResponsePtrOutputWithContext(ctx context.Context) GroupPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPropertiesResponsePtrOutput)
}

type GroupPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GroupPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupPropertiesResponse)(nil)).Elem()
}

func (o GroupPropertiesResponseOutput) ToGroupPropertiesResponseOutput() GroupPropertiesResponseOutput {
	return o
}

func (o GroupPropertiesResponseOutput) ToGroupPropertiesResponseOutputWithContext(ctx context.Context) GroupPropertiesResponseOutput {
	return o
}

func (o GroupPropertiesResponseOutput) ToGroupPropertiesResponsePtrOutput() GroupPropertiesResponsePtrOutput {
	return o.ToGroupPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o GroupPropertiesResponseOutput) ToGroupPropertiesResponsePtrOutputWithContext(ctx context.Context) GroupPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupPropertiesResponse) *GroupPropertiesResponse {
		return &v
	}).(GroupPropertiesResponsePtrOutput)
}

func (o GroupPropertiesResponseOutput) AreAssessmentsRunning() pulumi.BoolOutput {
	return o.ApplyT(func(v GroupPropertiesResponse) bool { return v.AreAssessmentsRunning }).(pulumi.BoolOutput)
}

func (o GroupPropertiesResponseOutput) Assessments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupPropertiesResponse) []string { return v.Assessments }).(pulumi.StringArrayOutput)
}

func (o GroupPropertiesResponseOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v GroupPropertiesResponse) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o GroupPropertiesResponseOutput) GroupStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GroupPropertiesResponse) string { return v.GroupStatus }).(pulumi.StringOutput)
}

func (o GroupPropertiesResponseOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GroupPropertiesResponse) *string { return v.GroupType }).(pulumi.StringPtrOutput)
}

func (o GroupPropertiesResponseOutput) MachineCount() pulumi.IntOutput {
	return o.ApplyT(func(v GroupPropertiesResponse) int { return v.MachineCount }).(pulumi.IntOutput)
}

func (o GroupPropertiesResponseOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v GroupPropertiesResponse) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

type GroupPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (GroupPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPropertiesResponse)(nil)).Elem()
}

func (o GroupPropertiesResponsePtrOutput) ToGroupPropertiesResponsePtrOutput() GroupPropertiesResponsePtrOutput {
	return o
}

func (o GroupPropertiesResponsePtrOutput) ToGroupPropertiesResponsePtrOutputWithContext(ctx context.Context) GroupPropertiesResponsePtrOutput {
	return o
}

func (o GroupPropertiesResponsePtrOutput) Elem() GroupPropertiesResponseOutput {
	return o.ApplyT(func(v *GroupPropertiesResponse) GroupPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret GroupPropertiesResponse
		return ret
	}).(GroupPropertiesResponseOutput)
}

func (o GroupPropertiesResponsePtrOutput) AreAssessmentsRunning() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.AreAssessmentsRunning
	}).(pulumi.BoolPtrOutput)
}

func (o GroupPropertiesResponsePtrOutput) Assessments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Assessments
	}).(pulumi.StringArrayOutput)
}

func (o GroupPropertiesResponsePtrOutput) CreatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o GroupPropertiesResponsePtrOutput) GroupStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GroupStatus
	}).(pulumi.StringPtrOutput)
}

func (o GroupPropertiesResponsePtrOutput) GroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupType
	}).(pulumi.StringPtrOutput)
}

func (o GroupPropertiesResponsePtrOutput) MachineCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GroupPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.MachineCount
	}).(pulumi.IntPtrOutput)
}

func (o GroupPropertiesResponsePtrOutput) UpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

type ImportCollectorProperties struct {
	DiscoverySiteId *string `pulumi:"discoverySiteId"`
}





type ImportCollectorPropertiesInput interface {
	pulumi.Input

	ToImportCollectorPropertiesOutput() ImportCollectorPropertiesOutput
	ToImportCollectorPropertiesOutputWithContext(context.Context) ImportCollectorPropertiesOutput
}

type ImportCollectorPropertiesArgs struct {
	DiscoverySiteId pulumi.StringPtrInput `pulumi:"discoverySiteId"`
}

func (ImportCollectorPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportCollectorProperties)(nil)).Elem()
}

func (i ImportCollectorPropertiesArgs) ToImportCollectorPropertiesOutput() ImportCollectorPropertiesOutput {
	return i.ToImportCollectorPropertiesOutputWithContext(context.Background())
}

func (i ImportCollectorPropertiesArgs) ToImportCollectorPropertiesOutputWithContext(ctx context.Context) ImportCollectorPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportCollectorPropertiesOutput)
}

func (i ImportCollectorPropertiesArgs) ToImportCollectorPropertiesPtrOutput() ImportCollectorPropertiesPtrOutput {
	return i.ToImportCollectorPropertiesPtrOutputWithContext(context.Background())
}

func (i ImportCollectorPropertiesArgs) ToImportCollectorPropertiesPtrOutputWithContext(ctx context.Context) ImportCollectorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportCollectorPropertiesOutput).ToImportCollectorPropertiesPtrOutputWithContext(ctx)
}









type ImportCollectorPropertiesPtrInput interface {
	pulumi.Input

	ToImportCollectorPropertiesPtrOutput() ImportCollectorPropertiesPtrOutput
	ToImportCollectorPropertiesPtrOutputWithContext(context.Context) ImportCollectorPropertiesPtrOutput
}

type importCollectorPropertiesPtrType ImportCollectorPropertiesArgs

func ImportCollectorPropertiesPtr(v *ImportCollectorPropertiesArgs) ImportCollectorPropertiesPtrInput {
	return (*importCollectorPropertiesPtrType)(v)
}

func (*importCollectorPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportCollectorProperties)(nil)).Elem()
}

func (i *importCollectorPropertiesPtrType) ToImportCollectorPropertiesPtrOutput() ImportCollectorPropertiesPtrOutput {
	return i.ToImportCollectorPropertiesPtrOutputWithContext(context.Background())
}

func (i *importCollectorPropertiesPtrType) ToImportCollectorPropertiesPtrOutputWithContext(ctx context.Context) ImportCollectorPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportCollectorPropertiesPtrOutput)
}

type ImportCollectorPropertiesOutput struct{ *pulumi.OutputState }

func (ImportCollectorPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportCollectorProperties)(nil)).Elem()
}

func (o ImportCollectorPropertiesOutput) ToImportCollectorPropertiesOutput() ImportCollectorPropertiesOutput {
	return o
}

func (o ImportCollectorPropertiesOutput) ToImportCollectorPropertiesOutputWithContext(ctx context.Context) ImportCollectorPropertiesOutput {
	return o
}

func (o ImportCollectorPropertiesOutput) ToImportCollectorPropertiesPtrOutput() ImportCollectorPropertiesPtrOutput {
	return o.ToImportCollectorPropertiesPtrOutputWithContext(context.Background())
}

func (o ImportCollectorPropertiesOutput) ToImportCollectorPropertiesPtrOutputWithContext(ctx context.Context) ImportCollectorPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImportCollectorProperties) *ImportCollectorProperties {
		return &v
	}).(ImportCollectorPropertiesPtrOutput)
}

func (o ImportCollectorPropertiesOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportCollectorProperties) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

type ImportCollectorPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ImportCollectorPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportCollectorProperties)(nil)).Elem()
}

func (o ImportCollectorPropertiesPtrOutput) ToImportCollectorPropertiesPtrOutput() ImportCollectorPropertiesPtrOutput {
	return o
}

func (o ImportCollectorPropertiesPtrOutput) ToImportCollectorPropertiesPtrOutputWithContext(ctx context.Context) ImportCollectorPropertiesPtrOutput {
	return o
}

func (o ImportCollectorPropertiesPtrOutput) Elem() ImportCollectorPropertiesOutput {
	return o.ApplyT(func(v *ImportCollectorProperties) ImportCollectorProperties {
		if v != nil {
			return *v
		}
		var ret ImportCollectorProperties
		return ret
	}).(ImportCollectorPropertiesOutput)
}

func (o ImportCollectorPropertiesPtrOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportCollectorProperties) *string {
		if v == nil {
			return nil
		}
		return v.DiscoverySiteId
	}).(pulumi.StringPtrOutput)
}

type ImportCollectorPropertiesResponse struct {
	CreatedTimestamp string  `pulumi:"createdTimestamp"`
	DiscoverySiteId  *string `pulumi:"discoverySiteId"`
	UpdatedTimestamp string  `pulumi:"updatedTimestamp"`
}





type ImportCollectorPropertiesResponseInput interface {
	pulumi.Input

	ToImportCollectorPropertiesResponseOutput() ImportCollectorPropertiesResponseOutput
	ToImportCollectorPropertiesResponseOutputWithContext(context.Context) ImportCollectorPropertiesResponseOutput
}

type ImportCollectorPropertiesResponseArgs struct {
	CreatedTimestamp pulumi.StringInput    `pulumi:"createdTimestamp"`
	DiscoverySiteId  pulumi.StringPtrInput `pulumi:"discoverySiteId"`
	UpdatedTimestamp pulumi.StringInput    `pulumi:"updatedTimestamp"`
}

func (ImportCollectorPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportCollectorPropertiesResponse)(nil)).Elem()
}

func (i ImportCollectorPropertiesResponseArgs) ToImportCollectorPropertiesResponseOutput() ImportCollectorPropertiesResponseOutput {
	return i.ToImportCollectorPropertiesResponseOutputWithContext(context.Background())
}

func (i ImportCollectorPropertiesResponseArgs) ToImportCollectorPropertiesResponseOutputWithContext(ctx context.Context) ImportCollectorPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportCollectorPropertiesResponseOutput)
}

func (i ImportCollectorPropertiesResponseArgs) ToImportCollectorPropertiesResponsePtrOutput() ImportCollectorPropertiesResponsePtrOutput {
	return i.ToImportCollectorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ImportCollectorPropertiesResponseArgs) ToImportCollectorPropertiesResponsePtrOutputWithContext(ctx context.Context) ImportCollectorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportCollectorPropertiesResponseOutput).ToImportCollectorPropertiesResponsePtrOutputWithContext(ctx)
}









type ImportCollectorPropertiesResponsePtrInput interface {
	pulumi.Input

	ToImportCollectorPropertiesResponsePtrOutput() ImportCollectorPropertiesResponsePtrOutput
	ToImportCollectorPropertiesResponsePtrOutputWithContext(context.Context) ImportCollectorPropertiesResponsePtrOutput
}

type importCollectorPropertiesResponsePtrType ImportCollectorPropertiesResponseArgs

func ImportCollectorPropertiesResponsePtr(v *ImportCollectorPropertiesResponseArgs) ImportCollectorPropertiesResponsePtrInput {
	return (*importCollectorPropertiesResponsePtrType)(v)
}

func (*importCollectorPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportCollectorPropertiesResponse)(nil)).Elem()
}

func (i *importCollectorPropertiesResponsePtrType) ToImportCollectorPropertiesResponsePtrOutput() ImportCollectorPropertiesResponsePtrOutput {
	return i.ToImportCollectorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *importCollectorPropertiesResponsePtrType) ToImportCollectorPropertiesResponsePtrOutputWithContext(ctx context.Context) ImportCollectorPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportCollectorPropertiesResponsePtrOutput)
}

type ImportCollectorPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ImportCollectorPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportCollectorPropertiesResponse)(nil)).Elem()
}

func (o ImportCollectorPropertiesResponseOutput) ToImportCollectorPropertiesResponseOutput() ImportCollectorPropertiesResponseOutput {
	return o
}

func (o ImportCollectorPropertiesResponseOutput) ToImportCollectorPropertiesResponseOutputWithContext(ctx context.Context) ImportCollectorPropertiesResponseOutput {
	return o
}

func (o ImportCollectorPropertiesResponseOutput) ToImportCollectorPropertiesResponsePtrOutput() ImportCollectorPropertiesResponsePtrOutput {
	return o.ToImportCollectorPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ImportCollectorPropertiesResponseOutput) ToImportCollectorPropertiesResponsePtrOutputWithContext(ctx context.Context) ImportCollectorPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImportCollectorPropertiesResponse) *ImportCollectorPropertiesResponse {
		return &v
	}).(ImportCollectorPropertiesResponsePtrOutput)
}

func (o ImportCollectorPropertiesResponseOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ImportCollectorPropertiesResponse) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o ImportCollectorPropertiesResponseOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportCollectorPropertiesResponse) *string { return v.DiscoverySiteId }).(pulumi.StringPtrOutput)
}

func (o ImportCollectorPropertiesResponseOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ImportCollectorPropertiesResponse) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

type ImportCollectorPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ImportCollectorPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportCollectorPropertiesResponse)(nil)).Elem()
}

func (o ImportCollectorPropertiesResponsePtrOutput) ToImportCollectorPropertiesResponsePtrOutput() ImportCollectorPropertiesResponsePtrOutput {
	return o
}

func (o ImportCollectorPropertiesResponsePtrOutput) ToImportCollectorPropertiesResponsePtrOutputWithContext(ctx context.Context) ImportCollectorPropertiesResponsePtrOutput {
	return o
}

func (o ImportCollectorPropertiesResponsePtrOutput) Elem() ImportCollectorPropertiesResponseOutput {
	return o.ApplyT(func(v *ImportCollectorPropertiesResponse) ImportCollectorPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ImportCollectorPropertiesResponse
		return ret
	}).(ImportCollectorPropertiesResponseOutput)
}

func (o ImportCollectorPropertiesResponsePtrOutput) CreatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportCollectorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ImportCollectorPropertiesResponsePtrOutput) DiscoverySiteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportCollectorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiscoverySiteId
	}).(pulumi.StringPtrOutput)
}

func (o ImportCollectorPropertiesResponsePtrOutput) UpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportCollectorPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionProperties struct {
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
}





type PrivateEndpointConnectionPropertiesInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput
	ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesOutput
}

type PrivateEndpointConnectionPropertiesArgs struct {
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput `pulumi:"privateLinkServiceConnectionState"`
}

func (PrivateEndpointConnectionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return i.ToPrivateEndpointConnectionPropertiesOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput)
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesArgs) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesOutput).ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesPtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput
	ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesPtrOutput
}

type privateEndpointConnectionPropertiesPtrType PrivateEndpointConnectionPropertiesArgs

func PrivateEndpointConnectionPropertiesPtr(v *PrivateEndpointConnectionPropertiesArgs) PrivateEndpointConnectionPropertiesPtrInput {
	return (*privateEndpointConnectionPropertiesPtrType)(v)
}

func (*privateEndpointConnectionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesPtrType) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesPtrOutput)
}

type PrivateEndpointConnectionPropertiesOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutput() PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionProperties) *PrivateEndpointConnectionProperties {
		return &v
	}).(PrivateEndpointConnectionPropertiesPtrOutput)
}

func (o PrivateEndpointConnectionPropertiesOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionState {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionProperties)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutput() PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) ToPrivateEndpointConnectionPropertiesPtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesPtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) Elem() PrivateEndpointConnectionPropertiesOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) PrivateEndpointConnectionProperties {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionProperties
		return ret
	}).(PrivateEndpointConnectionPropertiesOutput)
}

func (o PrivateEndpointConnectionPropertiesPtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionProperties) *PrivateLinkServiceConnectionState {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponse struct {
	PrivateEndpoint                   ResourceIdResponse                         `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
}





type PrivateEndpointConnectionPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput
	ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponseOutput
}

type PrivateEndpointConnectionPropertiesResponseArgs struct {
	PrivateEndpoint                   ResourceIdResponseInput                           `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                `pulumi:"provisioningState"`
}

func (PrivateEndpointConnectionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionPropertiesResponseArgs) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponseOutput).ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointConnectionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput
	ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput
}

type privateEndpointConnectionPropertiesResponsePtrType PrivateEndpointConnectionPropertiesResponseArgs

func PrivateEndpointConnectionPropertiesResponsePtr(v *PrivateEndpointConnectionPropertiesResponseArgs) PrivateEndpointConnectionPropertiesResponsePtrInput {
	return (*privateEndpointConnectionPropertiesResponsePtrType)(v)
}

func (*privateEndpointConnectionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return i.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointConnectionPropertiesResponsePtrType) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

type PrivateEndpointConnectionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutput() PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointConnectionPropertiesResponse) *PrivateEndpointConnectionPropertiesResponse {
		return &v
	}).(PrivateEndpointConnectionPropertiesResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateEndpoint() ResourceIdResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) ResourceIdResponse { return v.PrivateEndpoint }).(ResourceIdResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointConnectionPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutput() PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ToPrivateEndpointConnectionPropertiesResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointConnectionPropertiesResponsePtrOutput {
	return o
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) Elem() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) PrivateEndpointConnectionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointConnectionPropertiesResponse
		return ret
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateEndpoint() ResourceIdResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *ResourceIdResponse {
		if v == nil {
			return nil
		}
		return &v.PrivateEndpoint
	}).(ResourceIdResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *PrivateLinkServiceConnectionStateResponse {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointConnectionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	ETag       *string                                     `pulumi:"eTag"`
	Id         string                                      `pulumi:"id"`
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	Type       string                                      `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	ETag       pulumi.StringPtrInput                            `pulumi:"eTag"`
	Id         pulumi.StringInput                               `pulumi:"id"`
	Name       pulumi.StringInput                               `pulumi:"name"`
	Properties PrivateEndpointConnectionPropertiesResponseInput `pulumi:"properties"`
	Type       pulumi.StringInput                               `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Properties() PrivateEndpointConnectionPropertiesResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) PrivateEndpointConnectionPropertiesResponse {
		return v.Properties
	}).(PrivateEndpointConnectionPropertiesResponseOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ProjectProperties struct {
	AssessmentSolutionId        *string `pulumi:"assessmentSolutionId"`
	CustomerStorageAccountArmId *string `pulumi:"customerStorageAccountArmId"`
	CustomerWorkspaceId         *string `pulumi:"customerWorkspaceId"`
	CustomerWorkspaceLocation   *string `pulumi:"customerWorkspaceLocation"`
	ProjectStatus               *string `pulumi:"projectStatus"`
	PublicNetworkAccess         *string `pulumi:"publicNetworkAccess"`
}





type ProjectPropertiesInput interface {
	pulumi.Input

	ToProjectPropertiesOutput() ProjectPropertiesOutput
	ToProjectPropertiesOutputWithContext(context.Context) ProjectPropertiesOutput
}

type ProjectPropertiesArgs struct {
	AssessmentSolutionId        pulumi.StringPtrInput `pulumi:"assessmentSolutionId"`
	CustomerStorageAccountArmId pulumi.StringPtrInput `pulumi:"customerStorageAccountArmId"`
	CustomerWorkspaceId         pulumi.StringPtrInput `pulumi:"customerWorkspaceId"`
	CustomerWorkspaceLocation   pulumi.StringPtrInput `pulumi:"customerWorkspaceLocation"`
	ProjectStatus               pulumi.StringPtrInput `pulumi:"projectStatus"`
	PublicNetworkAccess         pulumi.StringPtrInput `pulumi:"publicNetworkAccess"`
}

func (ProjectPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectProperties)(nil)).Elem()
}

func (i ProjectPropertiesArgs) ToProjectPropertiesOutput() ProjectPropertiesOutput {
	return i.ToProjectPropertiesOutputWithContext(context.Background())
}

func (i ProjectPropertiesArgs) ToProjectPropertiesOutputWithContext(ctx context.Context) ProjectPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPropertiesOutput)
}

func (i ProjectPropertiesArgs) ToProjectPropertiesPtrOutput() ProjectPropertiesPtrOutput {
	return i.ToProjectPropertiesPtrOutputWithContext(context.Background())
}

func (i ProjectPropertiesArgs) ToProjectPropertiesPtrOutputWithContext(ctx context.Context) ProjectPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPropertiesOutput).ToProjectPropertiesPtrOutputWithContext(ctx)
}









type ProjectPropertiesPtrInput interface {
	pulumi.Input

	ToProjectPropertiesPtrOutput() ProjectPropertiesPtrOutput
	ToProjectPropertiesPtrOutputWithContext(context.Context) ProjectPropertiesPtrOutput
}

type projectPropertiesPtrType ProjectPropertiesArgs

func ProjectPropertiesPtr(v *ProjectPropertiesArgs) ProjectPropertiesPtrInput {
	return (*projectPropertiesPtrType)(v)
}

func (*projectPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectProperties)(nil)).Elem()
}

func (i *projectPropertiesPtrType) ToProjectPropertiesPtrOutput() ProjectPropertiesPtrOutput {
	return i.ToProjectPropertiesPtrOutputWithContext(context.Background())
}

func (i *projectPropertiesPtrType) ToProjectPropertiesPtrOutputWithContext(ctx context.Context) ProjectPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPropertiesPtrOutput)
}

type ProjectPropertiesOutput struct{ *pulumi.OutputState }

func (ProjectPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectProperties)(nil)).Elem()
}

func (o ProjectPropertiesOutput) ToProjectPropertiesOutput() ProjectPropertiesOutput {
	return o
}

func (o ProjectPropertiesOutput) ToProjectPropertiesOutputWithContext(ctx context.Context) ProjectPropertiesOutput {
	return o
}

func (o ProjectPropertiesOutput) ToProjectPropertiesPtrOutput() ProjectPropertiesPtrOutput {
	return o.ToProjectPropertiesPtrOutputWithContext(context.Background())
}

func (o ProjectPropertiesOutput) ToProjectPropertiesPtrOutputWithContext(ctx context.Context) ProjectPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectProperties) *ProjectProperties {
		return &v
	}).(ProjectPropertiesPtrOutput)
}

func (o ProjectPropertiesOutput) AssessmentSolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectProperties) *string { return v.AssessmentSolutionId }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesOutput) CustomerStorageAccountArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectProperties) *string { return v.CustomerStorageAccountArmId }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesOutput) CustomerWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectProperties) *string { return v.CustomerWorkspaceId }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesOutput) CustomerWorkspaceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectProperties) *string { return v.CustomerWorkspaceLocation }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesOutput) ProjectStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectProperties) *string { return v.ProjectStatus }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

type ProjectPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ProjectPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectProperties)(nil)).Elem()
}

func (o ProjectPropertiesPtrOutput) ToProjectPropertiesPtrOutput() ProjectPropertiesPtrOutput {
	return o
}

func (o ProjectPropertiesPtrOutput) ToProjectPropertiesPtrOutputWithContext(ctx context.Context) ProjectPropertiesPtrOutput {
	return o
}

func (o ProjectPropertiesPtrOutput) Elem() ProjectPropertiesOutput {
	return o.ApplyT(func(v *ProjectProperties) ProjectProperties {
		if v != nil {
			return *v
		}
		var ret ProjectProperties
		return ret
	}).(ProjectPropertiesOutput)
}

func (o ProjectPropertiesPtrOutput) AssessmentSolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.AssessmentSolutionId
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesPtrOutput) CustomerStorageAccountArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.CustomerStorageAccountArmId
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesPtrOutput) CustomerWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.CustomerWorkspaceId
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesPtrOutput) CustomerWorkspaceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.CustomerWorkspaceLocation
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesPtrOutput) ProjectStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProjectStatus
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

type ProjectPropertiesResponse struct {
	AssessmentSolutionId        *string                             `pulumi:"assessmentSolutionId"`
	CreatedTimestamp            string                              `pulumi:"createdTimestamp"`
	CustomerStorageAccountArmId *string                             `pulumi:"customerStorageAccountArmId"`
	CustomerWorkspaceId         *string                             `pulumi:"customerWorkspaceId"`
	CustomerWorkspaceLocation   *string                             `pulumi:"customerWorkspaceLocation"`
	LastAssessmentTimestamp     string                              `pulumi:"lastAssessmentTimestamp"`
	NumberOfAssessments         int                                 `pulumi:"numberOfAssessments"`
	NumberOfGroups              int                                 `pulumi:"numberOfGroups"`
	NumberOfMachines            int                                 `pulumi:"numberOfMachines"`
	PrivateEndpointConnections  []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	ProjectStatus               *string                             `pulumi:"projectStatus"`
	ProvisioningState           string                              `pulumi:"provisioningState"`
	PublicNetworkAccess         *string                             `pulumi:"publicNetworkAccess"`
	ServiceEndpoint             string                              `pulumi:"serviceEndpoint"`
	UpdatedTimestamp            string                              `pulumi:"updatedTimestamp"`
}





type ProjectPropertiesResponseInput interface {
	pulumi.Input

	ToProjectPropertiesResponseOutput() ProjectPropertiesResponseOutput
	ToProjectPropertiesResponseOutputWithContext(context.Context) ProjectPropertiesResponseOutput
}

type ProjectPropertiesResponseArgs struct {
	AssessmentSolutionId        pulumi.StringPtrInput                       `pulumi:"assessmentSolutionId"`
	CreatedTimestamp            pulumi.StringInput                          `pulumi:"createdTimestamp"`
	CustomerStorageAccountArmId pulumi.StringPtrInput                       `pulumi:"customerStorageAccountArmId"`
	CustomerWorkspaceId         pulumi.StringPtrInput                       `pulumi:"customerWorkspaceId"`
	CustomerWorkspaceLocation   pulumi.StringPtrInput                       `pulumi:"customerWorkspaceLocation"`
	LastAssessmentTimestamp     pulumi.StringInput                          `pulumi:"lastAssessmentTimestamp"`
	NumberOfAssessments         pulumi.IntInput                             `pulumi:"numberOfAssessments"`
	NumberOfGroups              pulumi.IntInput                             `pulumi:"numberOfGroups"`
	NumberOfMachines            pulumi.IntInput                             `pulumi:"numberOfMachines"`
	PrivateEndpointConnections  PrivateEndpointConnectionResponseArrayInput `pulumi:"privateEndpointConnections"`
	ProjectStatus               pulumi.StringPtrInput                       `pulumi:"projectStatus"`
	ProvisioningState           pulumi.StringInput                          `pulumi:"provisioningState"`
	PublicNetworkAccess         pulumi.StringPtrInput                       `pulumi:"publicNetworkAccess"`
	ServiceEndpoint             pulumi.StringInput                          `pulumi:"serviceEndpoint"`
	UpdatedTimestamp            pulumi.StringInput                          `pulumi:"updatedTimestamp"`
}

func (ProjectPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectPropertiesResponse)(nil)).Elem()
}

func (i ProjectPropertiesResponseArgs) ToProjectPropertiesResponseOutput() ProjectPropertiesResponseOutput {
	return i.ToProjectPropertiesResponseOutputWithContext(context.Background())
}

func (i ProjectPropertiesResponseArgs) ToProjectPropertiesResponseOutputWithContext(ctx context.Context) ProjectPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPropertiesResponseOutput)
}

func (i ProjectPropertiesResponseArgs) ToProjectPropertiesResponsePtrOutput() ProjectPropertiesResponsePtrOutput {
	return i.ToProjectPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ProjectPropertiesResponseArgs) ToProjectPropertiesResponsePtrOutputWithContext(ctx context.Context) ProjectPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPropertiesResponseOutput).ToProjectPropertiesResponsePtrOutputWithContext(ctx)
}









type ProjectPropertiesResponsePtrInput interface {
	pulumi.Input

	ToProjectPropertiesResponsePtrOutput() ProjectPropertiesResponsePtrOutput
	ToProjectPropertiesResponsePtrOutputWithContext(context.Context) ProjectPropertiesResponsePtrOutput
}

type projectPropertiesResponsePtrType ProjectPropertiesResponseArgs

func ProjectPropertiesResponsePtr(v *ProjectPropertiesResponseArgs) ProjectPropertiesResponsePtrInput {
	return (*projectPropertiesResponsePtrType)(v)
}

func (*projectPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectPropertiesResponse)(nil)).Elem()
}

func (i *projectPropertiesResponsePtrType) ToProjectPropertiesResponsePtrOutput() ProjectPropertiesResponsePtrOutput {
	return i.ToProjectPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *projectPropertiesResponsePtrType) ToProjectPropertiesResponsePtrOutputWithContext(ctx context.Context) ProjectPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectPropertiesResponsePtrOutput)
}

type ProjectPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ProjectPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProjectPropertiesResponse)(nil)).Elem()
}

func (o ProjectPropertiesResponseOutput) ToProjectPropertiesResponseOutput() ProjectPropertiesResponseOutput {
	return o
}

func (o ProjectPropertiesResponseOutput) ToProjectPropertiesResponseOutputWithContext(ctx context.Context) ProjectPropertiesResponseOutput {
	return o
}

func (o ProjectPropertiesResponseOutput) ToProjectPropertiesResponsePtrOutput() ProjectPropertiesResponsePtrOutput {
	return o.ToProjectPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ProjectPropertiesResponseOutput) ToProjectPropertiesResponsePtrOutputWithContext(ctx context.Context) ProjectPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProjectPropertiesResponse) *ProjectPropertiesResponse {
		return &v
	}).(ProjectPropertiesResponsePtrOutput)
}

func (o ProjectPropertiesResponseOutput) AssessmentSolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) *string { return v.AssessmentSolutionId }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponseOutput) CreatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) string { return v.CreatedTimestamp }).(pulumi.StringOutput)
}

func (o ProjectPropertiesResponseOutput) CustomerStorageAccountArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) *string { return v.CustomerStorageAccountArmId }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponseOutput) CustomerWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) *string { return v.CustomerWorkspaceId }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponseOutput) CustomerWorkspaceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) *string { return v.CustomerWorkspaceLocation }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponseOutput) LastAssessmentTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) string { return v.LastAssessmentTimestamp }).(pulumi.StringOutput)
}

func (o ProjectPropertiesResponseOutput) NumberOfAssessments() pulumi.IntOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) int { return v.NumberOfAssessments }).(pulumi.IntOutput)
}

func (o ProjectPropertiesResponseOutput) NumberOfGroups() pulumi.IntOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) int { return v.NumberOfGroups }).(pulumi.IntOutput)
}

func (o ProjectPropertiesResponseOutput) NumberOfMachines() pulumi.IntOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) int { return v.NumberOfMachines }).(pulumi.IntOutput)
}

func (o ProjectPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) []PrivateEndpointConnectionResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o ProjectPropertiesResponseOutput) ProjectStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) *string { return v.ProjectStatus }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ProjectPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponseOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) string { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

func (o ProjectPropertiesResponseOutput) UpdatedTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v ProjectPropertiesResponse) string { return v.UpdatedTimestamp }).(pulumi.StringOutput)
}

type ProjectPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ProjectPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectPropertiesResponse)(nil)).Elem()
}

func (o ProjectPropertiesResponsePtrOutput) ToProjectPropertiesResponsePtrOutput() ProjectPropertiesResponsePtrOutput {
	return o
}

func (o ProjectPropertiesResponsePtrOutput) ToProjectPropertiesResponsePtrOutputWithContext(ctx context.Context) ProjectPropertiesResponsePtrOutput {
	return o
}

func (o ProjectPropertiesResponsePtrOutput) Elem() ProjectPropertiesResponseOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) ProjectPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ProjectPropertiesResponse
		return ret
	}).(ProjectPropertiesResponseOutput)
}

func (o ProjectPropertiesResponsePtrOutput) AssessmentSolutionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AssessmentSolutionId
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) CreatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CreatedTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) CustomerStorageAccountArmId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomerStorageAccountArmId
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) CustomerWorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomerWorkspaceId
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) CustomerWorkspaceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomerWorkspaceLocation
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) LastAssessmentTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastAssessmentTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) NumberOfAssessments() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.NumberOfAssessments
	}).(pulumi.IntPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) NumberOfGroups() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.NumberOfGroups
	}).(pulumi.IntPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) NumberOfMachines() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.NumberOfMachines
	}).(pulumi.IntPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) []PrivateEndpointConnectionResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionResponseArrayOutput)
}

func (o ProjectPropertiesResponsePtrOutput) ProjectStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProjectStatus
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) ServiceEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServiceEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o ProjectPropertiesResponsePtrOutput) UpdatedTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProjectPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpdatedTimestamp
	}).(pulumi.StringPtrOutput)
}

type ResourceIdResponse struct {
	Id string `pulumi:"id"`
}





type ResourceIdResponseInput interface {
	pulumi.Input

	ToResourceIdResponseOutput() ResourceIdResponseOutput
	ToResourceIdResponseOutputWithContext(context.Context) ResourceIdResponseOutput
}

type ResourceIdResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ResourceIdResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdResponse)(nil)).Elem()
}

func (i ResourceIdResponseArgs) ToResourceIdResponseOutput() ResourceIdResponseOutput {
	return i.ToResourceIdResponseOutputWithContext(context.Background())
}

func (i ResourceIdResponseArgs) ToResourceIdResponseOutputWithContext(ctx context.Context) ResourceIdResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdResponseOutput)
}

func (i ResourceIdResponseArgs) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return i.ToResourceIdResponsePtrOutputWithContext(context.Background())
}

func (i ResourceIdResponseArgs) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdResponseOutput).ToResourceIdResponsePtrOutputWithContext(ctx)
}









type ResourceIdResponsePtrInput interface {
	pulumi.Input

	ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput
	ToResourceIdResponsePtrOutputWithContext(context.Context) ResourceIdResponsePtrOutput
}

type resourceIdResponsePtrType ResourceIdResponseArgs

func ResourceIdResponsePtr(v *ResourceIdResponseArgs) ResourceIdResponsePtrInput {
	return (*resourceIdResponsePtrType)(v)
}

func (*resourceIdResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdResponse)(nil)).Elem()
}

func (i *resourceIdResponsePtrType) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return i.ToResourceIdResponsePtrOutputWithContext(context.Background())
}

func (i *resourceIdResponsePtrType) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdResponsePtrOutput)
}

type ResourceIdResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdResponse)(nil)).Elem()
}

func (o ResourceIdResponseOutput) ToResourceIdResponseOutput() ResourceIdResponseOutput {
	return o
}

func (o ResourceIdResponseOutput) ToResourceIdResponseOutputWithContext(ctx context.Context) ResourceIdResponseOutput {
	return o
}

func (o ResourceIdResponseOutput) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return o.ToResourceIdResponsePtrOutputWithContext(context.Background())
}

func (o ResourceIdResponseOutput) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdResponse) *ResourceIdResponse {
		return &v
	}).(ResourceIdResponsePtrOutput)
}

func (o ResourceIdResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ResourceIdResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdResponse)(nil)).Elem()
}

func (o ResourceIdResponsePtrOutput) ToResourceIdResponsePtrOutput() ResourceIdResponsePtrOutput {
	return o
}

func (o ResourceIdResponsePtrOutput) ToResourceIdResponsePtrOutputWithContext(ctx context.Context) ResourceIdResponsePtrOutput {
	return o
}

func (o ResourceIdResponsePtrOutput) Elem() ResourceIdResponseOutput {
	return o.ApplyT(func(v *ResourceIdResponse) ResourceIdResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdResponse
		return ret
	}).(ResourceIdResponseOutput)
}

func (o ResourceIdResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type VmUptime struct {
	DaysPerMonth *float64 `pulumi:"daysPerMonth"`
	HoursPerDay  *float64 `pulumi:"hoursPerDay"`
}





type VmUptimeInput interface {
	pulumi.Input

	ToVmUptimeOutput() VmUptimeOutput
	ToVmUptimeOutputWithContext(context.Context) VmUptimeOutput
}

type VmUptimeArgs struct {
	DaysPerMonth pulumi.Float64PtrInput `pulumi:"daysPerMonth"`
	HoursPerDay  pulumi.Float64PtrInput `pulumi:"hoursPerDay"`
}

func (VmUptimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmUptime)(nil)).Elem()
}

func (i VmUptimeArgs) ToVmUptimeOutput() VmUptimeOutput {
	return i.ToVmUptimeOutputWithContext(context.Background())
}

func (i VmUptimeArgs) ToVmUptimeOutputWithContext(ctx context.Context) VmUptimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmUptimeOutput)
}

func (i VmUptimeArgs) ToVmUptimePtrOutput() VmUptimePtrOutput {
	return i.ToVmUptimePtrOutputWithContext(context.Background())
}

func (i VmUptimeArgs) ToVmUptimePtrOutputWithContext(ctx context.Context) VmUptimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmUptimeOutput).ToVmUptimePtrOutputWithContext(ctx)
}









type VmUptimePtrInput interface {
	pulumi.Input

	ToVmUptimePtrOutput() VmUptimePtrOutput
	ToVmUptimePtrOutputWithContext(context.Context) VmUptimePtrOutput
}

type vmUptimePtrType VmUptimeArgs

func VmUptimePtr(v *VmUptimeArgs) VmUptimePtrInput {
	return (*vmUptimePtrType)(v)
}

func (*vmUptimePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VmUptime)(nil)).Elem()
}

func (i *vmUptimePtrType) ToVmUptimePtrOutput() VmUptimePtrOutput {
	return i.ToVmUptimePtrOutputWithContext(context.Background())
}

func (i *vmUptimePtrType) ToVmUptimePtrOutputWithContext(ctx context.Context) VmUptimePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmUptimePtrOutput)
}

type VmUptimeOutput struct{ *pulumi.OutputState }

func (VmUptimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmUptime)(nil)).Elem()
}

func (o VmUptimeOutput) ToVmUptimeOutput() VmUptimeOutput {
	return o
}

func (o VmUptimeOutput) ToVmUptimeOutputWithContext(ctx context.Context) VmUptimeOutput {
	return o
}

func (o VmUptimeOutput) ToVmUptimePtrOutput() VmUptimePtrOutput {
	return o.ToVmUptimePtrOutputWithContext(context.Background())
}

func (o VmUptimeOutput) ToVmUptimePtrOutputWithContext(ctx context.Context) VmUptimePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VmUptime) *VmUptime {
		return &v
	}).(VmUptimePtrOutput)
}

func (o VmUptimeOutput) DaysPerMonth() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VmUptime) *float64 { return v.DaysPerMonth }).(pulumi.Float64PtrOutput)
}

func (o VmUptimeOutput) HoursPerDay() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VmUptime) *float64 { return v.HoursPerDay }).(pulumi.Float64PtrOutput)
}

type VmUptimePtrOutput struct{ *pulumi.OutputState }

func (VmUptimePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmUptime)(nil)).Elem()
}

func (o VmUptimePtrOutput) ToVmUptimePtrOutput() VmUptimePtrOutput {
	return o
}

func (o VmUptimePtrOutput) ToVmUptimePtrOutputWithContext(ctx context.Context) VmUptimePtrOutput {
	return o
}

func (o VmUptimePtrOutput) Elem() VmUptimeOutput {
	return o.ApplyT(func(v *VmUptime) VmUptime {
		if v != nil {
			return *v
		}
		var ret VmUptime
		return ret
	}).(VmUptimeOutput)
}

func (o VmUptimePtrOutput) DaysPerMonth() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VmUptime) *float64 {
		if v == nil {
			return nil
		}
		return v.DaysPerMonth
	}).(pulumi.Float64PtrOutput)
}

func (o VmUptimePtrOutput) HoursPerDay() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VmUptime) *float64 {
		if v == nil {
			return nil
		}
		return v.HoursPerDay
	}).(pulumi.Float64PtrOutput)
}

type VmUptimeResponse struct {
	DaysPerMonth *float64 `pulumi:"daysPerMonth"`
	HoursPerDay  *float64 `pulumi:"hoursPerDay"`
}





type VmUptimeResponseInput interface {
	pulumi.Input

	ToVmUptimeResponseOutput() VmUptimeResponseOutput
	ToVmUptimeResponseOutputWithContext(context.Context) VmUptimeResponseOutput
}

type VmUptimeResponseArgs struct {
	DaysPerMonth pulumi.Float64PtrInput `pulumi:"daysPerMonth"`
	HoursPerDay  pulumi.Float64PtrInput `pulumi:"hoursPerDay"`
}

func (VmUptimeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmUptimeResponse)(nil)).Elem()
}

func (i VmUptimeResponseArgs) ToVmUptimeResponseOutput() VmUptimeResponseOutput {
	return i.ToVmUptimeResponseOutputWithContext(context.Background())
}

func (i VmUptimeResponseArgs) ToVmUptimeResponseOutputWithContext(ctx context.Context) VmUptimeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmUptimeResponseOutput)
}

func (i VmUptimeResponseArgs) ToVmUptimeResponsePtrOutput() VmUptimeResponsePtrOutput {
	return i.ToVmUptimeResponsePtrOutputWithContext(context.Background())
}

func (i VmUptimeResponseArgs) ToVmUptimeResponsePtrOutputWithContext(ctx context.Context) VmUptimeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmUptimeResponseOutput).ToVmUptimeResponsePtrOutputWithContext(ctx)
}









type VmUptimeResponsePtrInput interface {
	pulumi.Input

	ToVmUptimeResponsePtrOutput() VmUptimeResponsePtrOutput
	ToVmUptimeResponsePtrOutputWithContext(context.Context) VmUptimeResponsePtrOutput
}

type vmUptimeResponsePtrType VmUptimeResponseArgs

func VmUptimeResponsePtr(v *VmUptimeResponseArgs) VmUptimeResponsePtrInput {
	return (*vmUptimeResponsePtrType)(v)
}

func (*vmUptimeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VmUptimeResponse)(nil)).Elem()
}

func (i *vmUptimeResponsePtrType) ToVmUptimeResponsePtrOutput() VmUptimeResponsePtrOutput {
	return i.ToVmUptimeResponsePtrOutputWithContext(context.Background())
}

func (i *vmUptimeResponsePtrType) ToVmUptimeResponsePtrOutputWithContext(ctx context.Context) VmUptimeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmUptimeResponsePtrOutput)
}

type VmUptimeResponseOutput struct{ *pulumi.OutputState }

func (VmUptimeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmUptimeResponse)(nil)).Elem()
}

func (o VmUptimeResponseOutput) ToVmUptimeResponseOutput() VmUptimeResponseOutput {
	return o
}

func (o VmUptimeResponseOutput) ToVmUptimeResponseOutputWithContext(ctx context.Context) VmUptimeResponseOutput {
	return o
}

func (o VmUptimeResponseOutput) ToVmUptimeResponsePtrOutput() VmUptimeResponsePtrOutput {
	return o.ToVmUptimeResponsePtrOutputWithContext(context.Background())
}

func (o VmUptimeResponseOutput) ToVmUptimeResponsePtrOutputWithContext(ctx context.Context) VmUptimeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VmUptimeResponse) *VmUptimeResponse {
		return &v
	}).(VmUptimeResponsePtrOutput)
}

func (o VmUptimeResponseOutput) DaysPerMonth() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VmUptimeResponse) *float64 { return v.DaysPerMonth }).(pulumi.Float64PtrOutput)
}

func (o VmUptimeResponseOutput) HoursPerDay() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v VmUptimeResponse) *float64 { return v.HoursPerDay }).(pulumi.Float64PtrOutput)
}

type VmUptimeResponsePtrOutput struct{ *pulumi.OutputState }

func (VmUptimeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmUptimeResponse)(nil)).Elem()
}

func (o VmUptimeResponsePtrOutput) ToVmUptimeResponsePtrOutput() VmUptimeResponsePtrOutput {
	return o
}

func (o VmUptimeResponsePtrOutput) ToVmUptimeResponsePtrOutputWithContext(ctx context.Context) VmUptimeResponsePtrOutput {
	return o
}

func (o VmUptimeResponsePtrOutput) Elem() VmUptimeResponseOutput {
	return o.ApplyT(func(v *VmUptimeResponse) VmUptimeResponse {
		if v != nil {
			return *v
		}
		var ret VmUptimeResponse
		return ret
	}).(VmUptimeResponseOutput)
}

func (o VmUptimeResponsePtrOutput) DaysPerMonth() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VmUptimeResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.DaysPerMonth
	}).(pulumi.Float64PtrOutput)
}

func (o VmUptimeResponsePtrOutput) HoursPerDay() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VmUptimeResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.HoursPerDay
	}).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AssessmentPropertiesOutput{})
	pulumi.RegisterOutputType(AssessmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AssessmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AssessmentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CollectorAgentPropertiesOutput{})
	pulumi.RegisterOutputType(CollectorAgentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CollectorAgentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CollectorAgentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CollectorBodyAgentSpnPropertiesOutput{})
	pulumi.RegisterOutputType(CollectorBodyAgentSpnPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CollectorBodyAgentSpnPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CollectorBodyAgentSpnPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CollectorPropertiesOutput{})
	pulumi.RegisterOutputType(CollectorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CollectorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CollectorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(GroupPropertiesOutput{})
	pulumi.RegisterOutputType(GroupPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GroupPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GroupPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ImportCollectorPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ProjectPropertiesOutput{})
	pulumi.RegisterOutputType(ProjectPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ProjectPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ProjectPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdResponsePtrOutput{})
	pulumi.RegisterOutputType(VmUptimeOutput{})
	pulumi.RegisterOutputType(VmUptimePtrOutput{})
	pulumi.RegisterOutputType(VmUptimeResponseOutput{})
	pulumi.RegisterOutputType(VmUptimeResponsePtrOutput{})
}
